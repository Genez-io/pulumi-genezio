package resources

import (
	"fmt"

	ca "github.com/Genez-io/pulumi-genezio/provider/cloud_adapters"
	"github.com/Genez-io/pulumi-genezio/provider/domain"
	fhp "github.com/Genez-io/pulumi-genezio/provider/function_handler_provider"
	"github.com/Genez-io/pulumi-genezio/provider/requests"
	"github.com/Genez-io/pulumi-genezio/provider/utils"
	p "github.com/pulumi/pulumi-go-provider"
)



type ServerlessFunction struct{}

type ServerlessFunctionArgs struct {
	Path string `pulumi:"path"` 
	ProjectName  string `pulumi:"projectName"`
	Name string `pulumi:"name"`
	Region	   string `pulumi:"region"`
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
	authToken, err := utils.IsLoggedIn(ctx)
	if err != nil {
		return name, ServerlessFunctionState{}, err
	}
	ctx = p.CtxWithValue(ctx, "authToken", authToken)

	
	state := ServerlessFunctionState{ServerlessFunctionArgs: input}
	if preview {
		return name, state, nil
	}


	backendPath := "."

	projectConfiguration := domain.ProjectConfiguration{
		Name: input.ProjectName,
		Region: input.Region,
		Options: domain.Options{
			NodeRuntime: "nodejs20.x",
			Architecture: "arm64",
		},
		CloudProvider: "genezio-cloud",
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
				Language: "ts",
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

	response, err := cloudAdapter.Deploy(ctx, cloudInputs, projectConfiguration, ca.CloudAdapterOptions{Stage: nil}, nil)
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
	responseEnv := requests.SetEnvironmentVariables(ctx, response.ProjectID, response.ProjectEnvID, environmentVariablesData)
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