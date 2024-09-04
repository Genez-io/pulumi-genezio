package resources

import (
	_ "embed"
	"fmt"
	"log"
	"strings"

	"github.com/Genez-io/pulumi-genezio/provider/domain"
	"github.com/Genez-io/pulumi-genezio/provider/requests"
	p "github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi-go-provider/infer"
)

type Project struct{}

type ProjectArgs struct {
	Name          string                        `pulumi:"name"`
	Region        string                        `pulumi:"region"`
	CloudProvider *string                       `pulumi:"cloudProvider,optional"`
	Environment   *[]domain.EnvironmentVariable `pulumi:"environment,optional"`
}

type ProjectState struct {
	ProjectArgs

	ProjectId    string `pulumi:"projectId"`
	ProjectEnvId string `pulumi:"projectEnvId"`
}

func (*Project) Diff(ctx p.Context, id string, olds ProjectState, news ProjectArgs) (p.DiffResponse, error) {
	diff := map[string]p.PropertyDiff{}

	if olds.Name != news.Name {
		diff["name"] = p.PropertyDiff{Kind: p.DeleteReplace}
	}

	if olds.Region != news.Region {
		diff["region"] = p.PropertyDiff{Kind: p.DeleteReplace}
	}

	if olds.CloudProvider == nil {
		if news.CloudProvider != nil && *news.CloudProvider != "genezio-cloud" {
			diff["cloudProvider"] = p.PropertyDiff{Kind: p.DeleteReplace}
		}
	} else {
		if news.CloudProvider != nil {
			if *olds.CloudProvider != *news.CloudProvider {
				diff["cloudProvider"] = p.PropertyDiff{Kind: p.DeleteReplace}
			}
		} else {
			if *olds.CloudProvider != "genezio-cloud" {
				diff["cloudProvider"] = p.PropertyDiff{Kind: p.DeleteReplace}
			}
		}
	}

	if olds.CloudProvider != news.CloudProvider {
		diff["cloudProvider"] = p.PropertyDiff{Kind: p.DeleteReplace}
	}

	if olds.Environment == nil {
		if news.Environment != nil {
			diff["environment"] = p.PropertyDiff{Kind: p.Update}
		}
	} else {
		if news.Environment != nil {
			if len(*olds.Environment) != len(*news.Environment) {
				diff["environment"] = p.PropertyDiff{Kind: p.Update}
			} else {
				for i, envVar := range *news.Environment {
					if (*olds.Environment)[i].Name != envVar.Name || (*olds.Environment)[i].Value != envVar.Value {
						diff["environment"] = p.PropertyDiff{Kind: p.Update}
						break
					}
				}
			}
		}
	}

	return p.DiffResponse{
		DeleteBeforeReplace: true,
		HasChanges:          len(diff) > 0,
		DetailedDiff:        diff,
	}, nil
}

func (*Project) Read(ctx p.Context, id string, inputs ProjectArgs, state ProjectState) (string, ProjectArgs, ProjectState, error) {

	if state.Name == "" {
		return id, inputs, state, nil
	}

	projectDetails, err := requests.GetProjectDetails(ctx, state.Name)
	if err != nil {
		if strings.Contains(err.Error(), "record not found") {
			return id, inputs, ProjectState{}, nil
		}
		return id, inputs, state, err
	}

	stage := "prod"

	contextStage := infer.GetConfig[*domain.Config](ctx).Stage
	if contextStage != nil {
		stage = *contextStage
	}

	var currentProjectEnv *domain.ProjectEnvDetails
	for _, projectEnv := range projectDetails.Project.ProjectEnvs {
		if projectEnv.Name == stage {
			currentProjectEnv = &projectEnv
			break
		}
	}

	if currentProjectEnv == nil {
		return id, inputs, ProjectState{}, nil
	}

	state.ProjectId = projectDetails.Project.Id
	state.ProjectEnvId = currentProjectEnv.Id
	state.CloudProvider = inputs.CloudProvider
	state.Region = projectDetails.Project.Region
	state.Name = projectDetails.Project.Name

	return id, inputs, state, nil
}

func (*Project) Create(ctx p.Context, name string, input ProjectArgs, preview bool) (string, ProjectState, error) {
	state := ProjectState{ProjectArgs: input}
	if preview {
		return name, state, nil
	}

	stage := "prod"

	cloudProvider := "genezio-cloud"
	if input.CloudProvider != nil {
		cloudProvider = *input.CloudProvider
	}

	// Check if the project and stage exists exists
	var currentProjectEnv *domain.ProjectEnvDetails
	var createProjectResponse *domain.CreateProjectResponse
	projectDetails, err := requests.GetProjectDetails(ctx, input.Name)
	if err != nil {
		if strings.Contains(err.Error(), "record not found") {
			response, err := requests.CreateProject(ctx, domain.CreateProjectRequest{
				ProjectName:   input.Name,
				Region:        input.Region,
				Stage:         stage,
				CloudProvider: cloudProvider,
			})
			if err != nil {
				return name, state, fmt.Errorf("error creating project: %v", err)
			}
			createProjectResponse = &response

		} else {
			return name, state, fmt.Errorf("error getting project details: %v", err)
		}
	} else {
		for _, projectEnv := range projectDetails.Project.ProjectEnvs {
			if projectEnv.Name == stage {
				currentProjectEnv = &projectEnv
				break
			}
		}
		if currentProjectEnv == nil {
			response, err := requests.CreateProject(ctx, domain.CreateProjectRequest{
				ProjectName:   input.Name,
				Region:        input.Region,
				Stage:         stage,
				CloudProvider: cloudProvider,
			})
			if err != nil {
				return name, state, fmt.Errorf("error creating project: %v", err)
			}
			createProjectResponse = &response
		} else {
			createProjectResponse = &domain.CreateProjectResponse{
				ProjectID:    projectDetails.Project.Id,
				ProjectEnvID: currentProjectEnv.Id,
			}
		}
	}

	// Set environment variables
	if input.Environment != nil && len(*input.Environment) > 0 {
		err := requests.SetEnvironmentVariables(ctx, createProjectResponse.ProjectID, createProjectResponse.ProjectEnvID, domain.SetEnvironmentVariablesRequest{
			EnvironmentVariables: *input.Environment,
		})
		if err != nil {
			log.Println("Error setting environment variables", err)
			return name, state, err
		}
	}

	state.ProjectId = createProjectResponse.ProjectID
	state.ProjectEnvId = createProjectResponse.ProjectEnvID

	return name, state, nil
}

func (*Project) Update(ctx p.Context, id string, olds ProjectState, news ProjectArgs, preview bool) (ProjectState, error) {

	state := ProjectState{
		ProjectArgs:  news,
		ProjectId:    olds.ProjectId,
		ProjectEnvId: olds.ProjectEnvId,
	}
	if preview {
		return state, nil
	}

	if news.Environment != nil && len(*news.Environment) > 0 {
		err := requests.SetEnvironmentVariables(ctx, state.ProjectId, state.ProjectEnvId, domain.SetEnvironmentVariablesRequest{
			EnvironmentVariables: *news.Environment,
		})
		if err != nil {
			log.Println("Error setting environment variables", err)
			return ProjectState{}, err
		}
	}

	return state, nil
}

func (*Project) Delete(ctx p.Context, id string, state ProjectState) error {
	_, err := requests.DeleteProject(ctx, state.ProjectId)
	if err != nil {
		if strings.Contains(err.Error(), "record not found") || strings.Contains(err.Error(), "405 Method Not Allowed") {
			return nil
		}
		log.Println("Error deleting project", err)
		return err
	}

	return nil
}

//go:embed documentation/project.md
var resourceDoc string

func (p *Project) Annotate(a infer.Annotator) {
	a.Describe(&p, resourceDoc)
}

func (p *ProjectArgs) Annotate(a infer.Annotator) {
	a.Describe(&p.Name, `The name of the project to be deployed. This is a required field.
	If you already have a project deployed with this name, then it will be updated with the new values. 
	Changing the name will result in a new project being created. 
	All the projects you deploy have a unique name. 
	If you try to deploy a project with the same name as an existing project but with a different region, 
	it will throw an error as you can't have two projects with the same name and different regions.`)
	a.Describe(&p.Region, `The region in which the project will be deployed.This is a required field. 
	You can only deploy a project in one region and you can't have two projects with the same name and different regions.

	Right now the only supported regions are:
	- us-east-1
	- eu-central-1
	`)
	a.Describe(&p.CloudProvider, `The cloud provider on which the project will be deployed. This is an optional field. The default value is "genezio-cloud".
	Currenly the only supported cloud providers are:
	- genezio-cloud

	More cloud providers will be supported in the future.`)

	a.Describe(&p.EnvironmentVariables, `The environment variables that will be set for the project. This is an optional field.
	This variable will allow you to set environment variables for the project. These environment variables will be available to all the functions and classes in the project.
	`)
}

func (p *ProjectState) Annotate(a infer.Annotator) {
	a.Describe(&p.ProjectId, `The ID of the project that was created. This is a unique identifier for the project.`)
	a.Describe(&p.ProjectEnvId, `The ID of the environment that was created. 
	Each project in the genezio platform can have environments such as prod, dev and staging. This is a unique identifier for the environment.`)
}
