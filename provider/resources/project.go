package resources

import (
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
	Name                 string                        `pulumi:"name"`
	Region               string                        `pulumi:"region"`
	CloudProvider        *string                       `pulumi:"cloudProvider,optional"`
	EnvironmentVariables *[]domain.EnvironmentVariable `pulumi:"environmentVariables,optional"`
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

	if olds.EnvironmentVariables == nil {
		if news.EnvironmentVariables != nil {
			diff["environmentVariables"] = p.PropertyDiff{Kind: p.Update}
		}
	} else {
		if news.EnvironmentVariables != nil {
			if len(*olds.EnvironmentVariables) != len(*news.EnvironmentVariables) {
				diff["environmentVariables"] = p.PropertyDiff{Kind: p.Update}
			} else {
				for i, envVar := range *news.EnvironmentVariables {
					if (*olds.EnvironmentVariables)[i].Name != envVar.Name || (*olds.EnvironmentVariables)[i].Value != envVar.Value {
						diff["environmentVariables"] = p.PropertyDiff{Kind: p.Update}
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
	if input.EnvironmentVariables != nil && len(*input.EnvironmentVariables) > 0 {
		err := requests.SetEnvironmentVariables(ctx, createProjectResponse.ProjectID, createProjectResponse.ProjectEnvID, domain.SetEnvironmentVariablesRequest{
			EnvironmentVariables: *input.EnvironmentVariables,
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

	if news.EnvironmentVariables != nil && len(*news.EnvironmentVariables) > 0 {
		err := requests.SetEnvironmentVariables(ctx, state.ProjectId, state.ProjectEnvId, domain.SetEnvironmentVariablesRequest{
			EnvironmentVariables: *news.EnvironmentVariables,
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
		if strings.Contains(err.Error(), "record not found") {
			return nil
		}
		log.Println("Error deleting project", err)
		return err
	}

	return nil
}
