package resources

import (
	"fmt"
	"os"

	ca "github.com/Genez-io/pulumi-genezio/provider/cloud_adapters"
	"github.com/Genez-io/pulumi-genezio/provider/domain"
	fhp "github.com/Genez-io/pulumi-genezio/provider/function_handler_provider"
	"github.com/Genez-io/pulumi-genezio/provider/requests"
	p "github.com/pulumi/pulumi-go-provider"
)



type ServerlessFunction struct{}

type ServerlessFunctionArgs struct {
	ProjectName  string `pulumi:"projectName"`
	Region	   string `pulumi:"region"`
	Stage *string `pulumi:"stage,optional"`
	CloudProvider *string `pulumi:"cloudProvider,optional"`
	BackendPath *string `pulumi:"backendPath,optional"`
	Language *string `pulumi:"language,optional"`
	Path string `pulumi:"path"` 
	Name string `pulumi:"name"`
	Entry string `pulumi:"entry"`
	Handler string `pulumi:"handler"`
	FolderHash *string `pulumi:"folderHash,optional"`
	EnvironmentVariables map[string]string `pulumi:"environmentVariables,optional"`
}

type ServerlessFunctionState struct {
	ServerlessFunctionArgs

	ID string `pulumi:"functionId"`
	URL string `pulumi:"url"`
	ProjectId string `pulumi:"projectId"`
	ProjectEnvId string `pulumi:"projectEnvId"`
}


func (*ServerlessFunction) Create(ctx p.Context, name string, input ServerlessFunctionArgs, preview bool) (string, ServerlessFunctionState, error) {
	
	state := ServerlessFunctionState{ServerlessFunctionArgs: input}
	if preview {
		return name, state, nil
	}

	cloudProvider := "genezio-cloud"
	if input.CloudProvider != nil {
		cloudProvider = *input.CloudProvider
	}

	backendPath,err := os.Getwd()
	if err != nil {
		fmt.Printf("An error occurred while trying to get the current working directory %v", err)
		return "", ServerlessFunctionState{}, err
	}

	if input.BackendPath != nil {
		backendPath = *input.BackendPath
	}

	language := "js"
	if input.Language != nil {
		language = *input.Language
	}



	projectConfiguration := domain.ProjectConfiguration{
		Name: input.ProjectName,
		Region: input.Region,
		Options: domain.Options{
			NodeRuntime: "nodejs20.x",
			Architecture: "arm64",
		},
		CloudProvider: cloudProvider,
		Workspace: domain.Workspace{
			Backend: backendPath,
		},
		AstSummary: domain.AstSummary{
			Version: "2",
			Classes: []string{},
		},
		Classes: []string{},
		Functions: []domain.FunctionConfiguration{
			{
				Name: input.Name,
				Path: input.Path,
				Language: language,
				Handler: input.Handler,
				Entry: input.Entry,
				Type: "aws",
			},
		},
	}

	cloudInput, err := fhp.FunctionToCloudInput(projectConfiguration.Functions[0], backendPath)
	if err != nil {
		fmt.Printf("An error occurred while trying to convert the function to cloud input %v", err)
		return "", ServerlessFunctionState{}, err
	}
	cloudInputs := []domain.GenezioCloudInput{cloudInput}

	cloudAdapter := ca.NewGenezioCloudAdapter()

	response, err := cloudAdapter.Deploy(ctx, cloudInputs, projectConfiguration, ca.CloudAdapterOptions{Stage: input.Stage}, nil)
	if err != nil {
		fmt.Printf("An error occurred while trying to deploy the function %v", err)
		return "", ServerlessFunctionState{}, err
	}

	var environmentVariablesData []domain.EnvironmentVariable
	for key, value := range input.EnvironmentVariables {
		environmentVariablesData = append(environmentVariablesData, domain.EnvironmentVariable{
			Name: key,
			Value: value,
		})
	}

	if len(environmentVariablesData) > 0{
	responseEnv := requests.SetEnvironmentVariables(ctx, response.ProjectID, response.ProjectEnvID, domain.SetEnvironmentVariablesRequest{
		EnvironmentVariables: environmentVariablesData,
	})
		if responseEnv != nil {
			fmt.Printf("An error occurred while trying to set environment variables %v", responseEnv)
			return "", ServerlessFunctionState{}, responseEnv
		}
	}



	state.ID = response.Functions[0].ID	
	state.URL = response.Functions[0].CloudUrl
	state.ProjectId = response.ProjectID
	state.ProjectEnvId = response.ProjectEnvID

	return name, state, nil
}