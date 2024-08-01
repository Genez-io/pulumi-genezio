package resources

import (
	"fmt"
	"log"
	"strings"

	"github.com/Genez-io/pulumi-genezio/provider/domain"
	"github.com/Genez-io/pulumi-genezio/provider/requests"
	p "github.com/pulumi/pulumi-go-provider"
)

type Project struct{}

type ProjectArgs struct {
	Name          string `pulumi:"name"`
	Region        string `pulumi:"region"`
	Stage         *string `pulumi:"stage,optional"`
	CloudProvider string `pulumi:"cloudProvider"`
	EnvironmentVariables map[string]string `pulumi:"environmentVariables,optional"`
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

	if olds.CloudProvider != news.CloudProvider {
		diff["cloudProvider"] = p.PropertyDiff{Kind: p.DeleteReplace}
	}

	if olds.Stage == nil {
		if news.Stage != nil && *news.Stage != "prod" {
			diff["stage"] = p.PropertyDiff{Kind: p.DeleteReplace}
		}
	} else {
		if news.Stage != nil {
			if *olds.Stage != *news.Stage {
				diff["stage"] = p.PropertyDiff{Kind: p.DeleteReplace}
			}
		} else {
			if *olds.Stage != "prod" {
				diff["stage"] = p.PropertyDiff{Kind: p.DeleteReplace}
			}
		}
	}

	if len(olds.EnvironmentVariables) != len(news.EnvironmentVariables) {
		diff["environmentVariables"] = p.PropertyDiff{Kind: p.Update}
	} else {
		for key, value := range news.EnvironmentVariables {
			if oldValue, ok := olds.EnvironmentVariables[key]; !ok || oldValue != value {
				diff["environmentVariables"] = p.PropertyDiff{Kind: p.Update}
				break
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
	
	projectDetails,err := requests.GetProjectDetails(ctx, inputs.Name)
	if err != nil {
		if strings.Contains(err.Error(), "405 Method Not Allowed") {
			return id, inputs, ProjectState{}, nil
		}
		return id, inputs, state, err
	}

	stage := "prod"
	if state.Stage != nil {
		stage = *state.Stage
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
	state.Region = inputs.Region
	state.Name = inputs.Name
	state.Stage = inputs.Stage

	return id, inputs, state, nil
}

func (*Project) Create(ctx p.Context, name string, input ProjectArgs, preview bool) (string, ProjectState, error) {
	state := ProjectState{ProjectArgs: input}
	if preview {
		return name, state, nil
	}


	stage := "prod"
	if input.Stage != nil {
		stage = *input.Stage
	}
	
	createProjectResponse,err := requests.CreateProject(ctx, domain.CreateProjectRequest{
		ProjectName: input.Name,
		Region: input.Region,
		Stage: stage,
		CloudProvider: input.CloudProvider,
	})
	if err != nil {
		return name, state, fmt.Errorf("error creating project: %v", err)
	}

	// Set environment variables
	var environmentVariablesData []domain.EnvironmentVariable
	for key, value := range input.EnvironmentVariables {
		environmentVariablesData = append(environmentVariablesData, domain.EnvironmentVariable{
			Name: key,
			Value: value,
		})
	}
	if len(environmentVariablesData) > 0{
		err := requests.SetEnvironmentVariables(ctx, createProjectResponse.ProjectID, createProjectResponse.ProjectEnvID, domain.SetEnvironmentVariablesRequest{
			EnvironmentVariables: environmentVariablesData,
		})
			if err != nil {
				log.Println("Error setting environment variables", err)
				return "", ProjectState{}, err
			}
		}

	state.ProjectId = createProjectResponse.ProjectID
	state.ProjectEnvId = createProjectResponse.ProjectEnvID

	return name, state, nil
}

func (*Project) Update(ctx p.Context, id string, olds ProjectState, news ProjectArgs, preview bool) (ProjectState, error) {
	
	state := ProjectState{ProjectArgs: news}
	if preview {
		return state, nil
	}
	
	state.ProjectId = olds.ProjectId
	state.ProjectEnvId = olds.ProjectEnvId
	
	// Set environment variables
	var environmentVariablesData []domain.EnvironmentVariable
	for key, value := range news.EnvironmentVariables {
		environmentVariablesData = append(environmentVariablesData, domain.EnvironmentVariable{
			Name: key,
			Value: value,
		})
	}


	if len(environmentVariablesData) > 0{
		err := requests.SetEnvironmentVariables(ctx, state.ProjectId, state.ProjectEnvId, domain.SetEnvironmentVariablesRequest{
			EnvironmentVariables: environmentVariablesData,
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
		if strings.Contains(err.Error(), "405 Method Not Allowed") {
			return nil
		}
		log.Println("Error deleting project", err)
		return err
	}

	return nil
}
