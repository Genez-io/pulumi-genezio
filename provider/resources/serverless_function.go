package resources

import (
	"fmt"
	"log"
	"os"

	ca "github.com/Genez-io/pulumi-genezio/provider/cloud_adapters"
	"github.com/Genez-io/pulumi-genezio/provider/domain"
	fhp "github.com/Genez-io/pulumi-genezio/provider/function_handler_provider"
	"github.com/Genez-io/pulumi-genezio/provider/utils"
	p "github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
)



type ServerlessFunction struct{}

type ServerlessFunctionArgs struct {
	ProjectName  string `pulumi:"projectName"`
	Region	   string `pulumi:"region"`
	Stage *string `pulumi:"stage,optional"`
	CloudProvider *string `pulumi:"cloudProvider,optional"`
	BackendPath *string `pulumi:"backendPath,optional"`
	Language *string `pulumi:"language,optional"`
	PathAsset resource.Archive `pulumi:"pathAsset"` 
	Name string `pulumi:"name"`
	Entry string `pulumi:"entry"`
	Handler string `pulumi:"handler"`
}

type ServerlessFunctionState struct {
	ServerlessFunctionArgs

	ID string `pulumi:"functionId"`
	URL string `pulumi:"url"`
	ProjectId string `pulumi:"projectId"`
	ProjectEnvId string `pulumi:"projectEnvId"`
}

func (*ServerlessFunction) Diff(ctx p.Context, id string, olds ServerlessFunctionState, news ServerlessFunctionArgs) (p.DiffResponse, error){
	diff := map[string]p.PropertyDiff{}

	if olds.ProjectName != news.ProjectName {
		diff["projectName"] = p.PropertyDiff{Kind: p.DeleteReplace}
	}

	if olds.Region != news.Region {
		diff["region"] = p.PropertyDiff{Kind: p.DeleteReplace}
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


	if olds.BackendPath == nil {
		if news.BackendPath != nil {
			diff["backendPath"] = p.PropertyDiff{Kind: p.DeleteReplace}
		}
	} else {
		if news.BackendPath == nil {
			diff["backendPath"] = p.PropertyDiff{Kind: p.DeleteReplace}
		} else {
			if *olds.BackendPath != *news.BackendPath {
				diff["backendPath"] = p.PropertyDiff{Kind: p.DeleteReplace}
			}
		}
	}

	if olds.Language == nil {
		if news.Language != nil && *news.Language != "js" {
			diff["language"] = p.PropertyDiff{Kind: p.DeleteReplace}
		}
	} else {
		if news.Language != nil {
			if *olds.Language != *news.Language {
				diff["language"] = p.PropertyDiff{Kind: p.DeleteReplace}
			}
		} else {
			if *olds.Language != "js" {
				diff["language"] = p.PropertyDiff{Kind: p.DeleteReplace}
			}
		}
	}

	if olds.PathAsset.Hash != news.PathAsset.Hash {
		diff["path"] = p.PropertyDiff{Kind: p.DeleteReplace}
	}

	if olds.Name != news.Name {
		diff["name"] = p.PropertyDiff{Kind: p.DeleteReplace}
	}

	if olds.Entry != news.Entry {
		diff["entry"] = p.PropertyDiff{Kind: p.DeleteReplace}
	}

	if olds.Handler != news.Handler {
		diff["handler"] = p.PropertyDiff{Kind: p.DeleteReplace}
	}

	return p.DiffResponse{
		DeleteBeforeReplace: true,
		HasChanges: len(diff) > 0,
		DetailedDiff: diff,
	}, nil


	
}


func (*ServerlessFunction) Create(ctx p.Context, name string, input ServerlessFunctionArgs, preview bool) (string, ServerlessFunctionState, error) {
	
	// TODO Will need to investigate further why this is needed, For now this is needed for the FileArchive to work
	// More info here https://pulumi-developer-docs.readthedocs.io/en/latest/architecture/deployment-schema.html#dabf18193072939515e22adb298388d-required
	input.PathAsset.Sig = "0def7320c3a5731c473e5ecbe6d01bc7"

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
				Path: input.PathAsset.Path,
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


	state.ID = response.Functions[0].ID	
	state.URL = response.Functions[0].CloudUrl
	state.ProjectId = response.ProjectID
	state.ProjectEnvId = response.ProjectEnvID

	err = utils.DeleteTemporaryFolder()
	if err != nil {
		log.Println("Error deleting temporary folder", err)
		return "", state, err
	}

	return name, state, nil
}