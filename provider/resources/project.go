package resources

import (
	"fmt"
	"log"

	"github.com/Genez-io/pulumi-genezio/provider/domain"
	"github.com/Genez-io/pulumi-genezio/provider/requests"
	p "github.com/pulumi/pulumi-go-provider"
)

type Project struct{}

type ProjectArgs struct {
	Name          string `pulumi:"name"`
	Region        string `pulumi:"region"`
	Stage         string `pulumi:"stage"`
	CloudProvider string `pulumi:"cloudProvider"`
}

type ProjectState struct {
	ProjectArgs

	ProjectId    string `pulumi:"projectId"`
	ProjectEnvId string `pulumi:"projectEnvId"`
}

func (*Project) Create(ctx p.Context, name string, input ProjectArgs, preview bool) (string, ProjectState, error) {
	state := ProjectState{ProjectArgs: input}
	if preview {
		return name, state, nil
	}


	
	createProjectResponse,err := requests.CreateProject(ctx, domain.CreateProjectRequest{
		ProjectName: input.Name,
		Region: input.Region,
		Stage: input.Stage,
		CloudProvider: input.CloudProvider,
	})
	if err != nil {
		return name, state, fmt.Errorf("error creating project: %v", err)
	}

	state.ProjectId = createProjectResponse.ProjectID
	state.ProjectEnvId = createProjectResponse.ProjectEnvID

	return name, state, nil
}

func (*Project) Delete(ctx p.Context, id string, state ProjectState) error {
	_, err := requests.DeleteProject(ctx, state.ProjectId)
	if err != nil {
		log.Println("Error deleting project", err)
		return err
	}

	return nil
}
