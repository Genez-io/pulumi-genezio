package resources

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	ca "github.com/Genez-io/pulumi-genezio/provider/cloud_adapters"
	"github.com/Genez-io/pulumi-genezio/provider/domain"
	fhp "github.com/Genez-io/pulumi-genezio/provider/function_handler_provider"
	"github.com/Genez-io/pulumi-genezio/provider/utils"
	p "github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi-go-provider/infer"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
)

type ServerlessFunction struct{}

type ServerlessFunctionArgs struct {
	Project     domain.Project   `pulumi:"project"`
	BackendPath *string          `pulumi:"backendPath,optional"`
	Language    *string          `pulumi:"language,optional"`
	Path        resource.Archive `pulumi:"path"`
	Name        string           `pulumi:"name"`
	Entry       string           `pulumi:"entry"`
	Handler     string           `pulumi:"handler"`
}

type ServerlessFunctionState struct {
	ServerlessFunctionArgs

	ID           string `pulumi:"functionId"`
	URL          string `pulumi:"url"`
	ProjectId    string `pulumi:"projectId"`
	ProjectEnvId string `pulumi:"projectEnvId"`
}

func (*ServerlessFunction) Diff(ctx p.Context, id string, olds ServerlessFunctionState, news ServerlessFunctionArgs) (p.DiffResponse, error) {
	diff := map[string]p.PropertyDiff{}

	areProjectsIdentical := utils.CompareProjects(olds.Project, news.Project)
	if !areProjectsIdentical {
		diff["project"] = p.PropertyDiff{Kind: p.DeleteReplace}
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

	if olds.Path.Hash != news.Path.Hash {
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
		HasChanges:          len(diff) > 0,
		DetailedDiff:        diff,
	}, nil

}

func (*ServerlessFunction) Create(ctx p.Context, name string, input ServerlessFunctionArgs, preview bool) (string, ServerlessFunctionState, error) {

	// // TODO Will need to investigate further why this is needed, For now this is needed for the FileArchive to work
	// // More info here https://pulumi-developer-docs.readthedocs.io/en/latest/architecture/deployment-schema.html#dabf18193072939515e22adb298388d-required
	input.Path.Sig = "0def7320c3a5731c473e5ecbe6d01bc7"

	state := ServerlessFunctionState{ServerlessFunctionArgs: input}
	if preview {
		return name, state, nil
	}

	cloudProvider := "genezio-cloud"

	stage := "prod"
	contextStage := infer.GetConfig[*domain.Config](ctx).Stage
	if contextStage != nil {
		stage = *contextStage
	}

	var absolueBackendPath string
	backendPath, err := os.Getwd()
	if err != nil {
		fmt.Printf("An error occurred while trying to get the current working directory %v", err)
		return "", ServerlessFunctionState{}, err
	}
	absolueBackendPath = backendPath

	if input.BackendPath != nil {
		absolueBackendPath = filepath.Join(absolueBackendPath, *input.BackendPath)
	}

	relFunctionPath, err := filepath.Rel(absolueBackendPath, input.Path.Path)
	if err != nil {
		fmt.Printf("An error occurred while trying to get the relative path %v", err)
		return "", ServerlessFunctionState{}, err
	}

	language := "js"
	if input.Language != nil {
		language = *input.Language
	}

	projectConfiguration := domain.ProjectConfiguration{
		Name:   input.Project.Name,
		Region: input.Project.Region,
		Options: domain.Options{
			NodeRuntime:  "nodejs20.x",
			Architecture: "arm64",
		},
		CloudProvider: cloudProvider,
		Workspace: domain.Workspace{
			Backend: absolueBackendPath,
		},
		AstSummary: domain.AstSummary{
			Version: "2",
			Classes: []string{},
		},
		Classes: []string{},
		Functions: []domain.FunctionConfiguration{
			{
				Name:     input.Name,
				Path:     relFunctionPath,
				Language: language,
				Handler:  input.Handler,
				Entry:    input.Entry,
				Type:     "aws",
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

	response, err := cloudAdapter.Deploy(ctx, cloudInputs, projectConfiguration, ca.CloudAdapterOptions{Stage: &stage}, nil)
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
