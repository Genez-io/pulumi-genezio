package resources

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	ca "github.com/Genez-io/pulumi-genezio/provider/cloud_adapters"
	"github.com/Genez-io/pulumi-genezio/provider/domain"
	"github.com/Genez-io/pulumi-genezio/provider/requests"
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

	ID  string `pulumi:"functionId"`
	URL string `pulumi:"url"`
}

func (*ServerlessFunction) Diff(ctx p.Context, id string, olds ServerlessFunctionState, news ServerlessFunctionArgs) (p.DiffResponse, error) {
	diff := map[string]p.PropertyDiff{}

	areProjectsIdentical := utils.CompareProjects(olds.Project, news.Project)
	if !areProjectsIdentical {
		diff["project"] = p.PropertyDiff{Kind: p.DeleteReplace}
	}

	if olds.BackendPath == nil {
		if news.BackendPath != nil {
			diff["backendPath"] = p.PropertyDiff{Kind: p.Update}
		}
	} else {
		if news.BackendPath == nil {
			diff["backendPath"] = p.PropertyDiff{Kind: p.Update}
		} else {
			if *olds.BackendPath != *news.BackendPath {
				diff["backendPath"] = p.PropertyDiff{Kind: p.Update}
			}
		}
	}

	if olds.Language == nil {
		if news.Language != nil && *news.Language != "js" {
			diff["language"] = p.PropertyDiff{Kind: p.Update}
		}
	} else {
		if news.Language != nil {
			if *olds.Language != *news.Language {
				diff["language"] = p.PropertyDiff{Kind: p.Update}
			}
		} else {
			if *olds.Language != "js" {
				diff["language"] = p.PropertyDiff{Kind: p.Update}
			}
		}
	}

	if olds.Path.Hash != news.Path.Hash {
		diff["path"] = p.PropertyDiff{Kind: p.Update}
	}

	if olds.Name != news.Name {
		diff["name"] = p.PropertyDiff{Kind: p.DeleteReplace}
	}

	if olds.Entry != news.Entry {
		diff["entry"] = p.PropertyDiff{Kind: p.Update}
	}

	if olds.Handler != news.Handler {
		diff["handler"] = p.PropertyDiff{Kind: p.Update}
	}

	return p.DiffResponse{
		DeleteBeforeReplace: true,
		HasChanges:          len(diff) > 0,
		DetailedDiff:        diff,
	}, nil

}

func (*ServerlessFunction) Update(ctx p.Context, id string, olds ServerlessFunctionState, news ServerlessFunctionArgs, preview bool) (ServerlessFunctionState, error) {
	// TODO Investigate why is this needed - .Sig is used to recognize the asset type when unmarshalling the resource object
	// This should not be hardcoded here, but rather automatically generated by the pulumi SDK
	// Code: https://github.com/pulumi/pulumi/blob/master/sdk/go/common/resource/sig/sig.go
	// Documentation: https://pulumi-developer-docs.readthedocs.io/en/latest/architecture/deployment-schema.html#dabf18193072939515e22adb298388d-required
	news.Path.Sig = resource.ArchiveSig
	olds.Path.Sig = resource.ArchiveSig

	state := ServerlessFunctionState{ServerlessFunctionArgs: news, ID: olds.ID, URL: olds.URL}
	if preview {
		return state, nil
	}

	stage := "prod"
	contextStage := infer.GetConfig[*domain.Config](ctx).Stage
	if contextStage != nil {
		stage = *contextStage
	}

	var absolueBackendPath string
	backendPath, err := os.Getwd()
	if err != nil {
		fmt.Printf("An error occurred while trying to get the current working directory %v", err)
		return ServerlessFunctionState{}, err
	}
	absolueBackendPath = backendPath

	if news.BackendPath != nil {
		absolueBackendPath = filepath.Join(absolueBackendPath, *news.BackendPath)
	}

	relFunctionPath, err := filepath.Rel(absolueBackendPath, news.Path.Path)
	if err != nil {
		fmt.Printf("An error occurred while trying to get the relative path %v", err)
		return ServerlessFunctionState{}, err
	}

	language := "js"
	if news.Language != nil {
		language = *news.Language
	}

	if news.Project.Name == "" {
		return ServerlessFunctionState{}, fmt.Errorf("project name is required")
	}

	// Create the project and stage if they don't exist
	var currentProjectEnv *domain.ProjectEnvDetails
	projectDetails, err := requests.GetProjectDetails(ctx, news.Project.Name)
	if err != nil {
		return ServerlessFunctionState{}, fmt.Errorf("error getting project details: %v", err)
	}
	for _, projectEnv := range projectDetails.Project.ProjectEnvs {
		if projectEnv.Name == stage {
			currentProjectEnv = &projectEnv
			break
		}
	}
	if currentProjectEnv == nil {
		return ServerlessFunctionState{}, fmt.Errorf("project env not found")
	}

	functionConfiguration := domain.FunctionConfiguration{
		Name:     news.Name,
		Path:     relFunctionPath,
		Language: language,
		Handler:  news.Handler,
		Entry:    "index.mjs",
		Type:     "aws",
	}
	archivePath, err := utils.CreateTemporaryFolder(nil, nil)
	if err != nil {
		fmt.Printf("An error occurred while trying to create a temporary folder %v\n", err)
		return ServerlessFunctionState{}, err
	}

	bundleFunctionScript := fmt.Sprintf("genezio bundleFunction --functionName %s --handler %s --entry %s --functionPath %s --backendPath %s --output %s", news.Name, news.Handler, news.Entry, relFunctionPath, absolueBackendPath, archivePath)
	err = utils.RunScriptsInDirectory(absolueBackendPath, []string{bundleFunctionScript}, nil)
	if err != nil {
		fmt.Printf("An error occurred while trying to bundle the function %v", err)
		return ServerlessFunctionState{}, err
	}

	cloudAdapter := ca.NewGenezioCloudAdapter()

	response, err := cloudAdapter.DeployFunction(ctx, news.Project.Name, news.Project.Region, functionConfiguration, archivePath, stage)
	if err != nil {
		fmt.Printf("An error occurred while trying to deploy the function %v", err)
		return ServerlessFunctionState{}, err
	}

	state.ID = response.Id
	state.URL = response.CloudURL

	fmt.Printf("the temporary folder is genezio-%d\n", os.Getpid())
	// err = utils.DeleteTemporaryFolder()
	// if err != nil {
	// 	log.Println("Error deleting temporary folder", err)
	// 	return state, err
	// }

	return state, nil

}

func (*ServerlessFunction) Read(ctx p.Context, id string, inputs ServerlessFunctionArgs, state ServerlessFunctionState) (string, ServerlessFunctionArgs, ServerlessFunctionState, error) {

	// TODO Investigate why is this needed - .Sig is used to recognize the asset type when unmarshalling the resource object
	// This should not be hardcoded here, but rather automatically generated by the pulumi SDK
	// Code: https://github.com/pulumi/pulumi/blob/master/sdk/go/common/resource/sig/sig.go
	// Documentation: https://pulumi-developer-docs.readthedocs.io/en/latest/architecture/deployment-schema.html#dabf18193072939515e22adb298388d-required
	inputs.Path.Sig = resource.ArchiveSig
	state.Path.Sig = resource.ArchiveSig

	stage := "prod"
	contextStage := infer.GetConfig[*domain.Config](ctx).Stage
	if contextStage != nil {
		stage = *contextStage
	}

	if state.Project.Name == "" {
		return id, inputs, state, nil
	}

	var currentProjectEnv *domain.ProjectEnvDetails
	projectDetails, err := requests.GetProjectDetails(ctx, state.Project.Name)
	if err != nil {
		if strings.Contains(err.Error(), "record not found") {
			state.Project.Name = ""
			return id, inputs, state, nil

		} else {
			return id, inputs, state, fmt.Errorf("error getting project details: %v", err)
		}
	} else {
		for _, projectEnv := range projectDetails.Project.ProjectEnvs {
			if projectEnv.Name == stage {
				currentProjectEnv = &projectEnv
				break
			}
		}
		if currentProjectEnv == nil {
			state.Project.Name = ""
			return id, inputs, state, nil
		}
	}

	getFunctionResponse, err := requests.GetFunction(ctx, state.ID)
	if err != nil {
		if strings.Contains(err.Error(), "record not found") {
			return id, inputs, state, nil
		} else {
			return id, inputs, state, err
		}
	}

	state.URL = getFunctionResponse.Function.CloudURL
	state.Name = getFunctionResponse.Function.Name
	state.Project.Name = projectDetails.Project.Name
	state.Project.Region = projectDetails.Project.Region

	return id, inputs, state, nil

}

func (*ServerlessFunction) Create(ctx p.Context, name string, input ServerlessFunctionArgs, preview bool) (string, ServerlessFunctionState, error) {
	// TODO Investigate why is this needed - .Sig is used to recognize the asset type when unmarshalling the resource object
	// This should not be hardcoded here, but rather automatically generated by the pulumi SDK
	// Code: https://github.com/pulumi/pulumi/blob/master/sdk/go/common/resource/sig/sig.go
	// Documentation: https://pulumi-developer-docs.readthedocs.io/en/latest/architecture/deployment-schema.html#dabf18193072939515e22adb298388d-required
	input.Path.Sig = resource.ArchiveSig

	state := ServerlessFunctionState{ServerlessFunctionArgs: input}
	if preview {
		return name, state, nil
	}

	if input.Project.Name == "" {
		return "", ServerlessFunctionState{}, fmt.Errorf("project name cannot be empty")
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

	// Create the project and stage if they don't exist
	var currentProjectEnv *domain.ProjectEnvDetails
	projectDetails, err := requests.GetProjectDetails(ctx, input.Project.Name)
	if err != nil {
		if strings.Contains(err.Error(), "record not found") {
			_, err := requests.CreateProject(ctx, domain.CreateProjectRequest{
				ProjectName:   input.Project.Name,
				Region:        input.Project.Region,
				Stage:         stage,
				CloudProvider: cloudProvider,
			})
			if err != nil {
				return name, state, fmt.Errorf("error creating project: %v", err)
			}

		} else {
			return "", ServerlessFunctionState{}, fmt.Errorf("error getting project details: %v", err)
		}
	} else {
		for _, projectEnv := range projectDetails.Project.ProjectEnvs {
			if projectEnv.Name == stage {
				currentProjectEnv = &projectEnv
				break
			}
		}
		if currentProjectEnv == nil {
			_, err := requests.CreateProject(ctx, domain.CreateProjectRequest{
				ProjectName:   input.Project.Name,
				Region:        input.Project.Region,
				Stage:         stage,
				CloudProvider: cloudProvider,
			})
			if err != nil {
				return name, state, fmt.Errorf("error creating project: %v", err)
			}
		}
	}

	functionConfiguration := domain.FunctionConfiguration{
		Name:     input.Name,
		Path:     relFunctionPath,
		Language: language,
		Handler:  input.Handler,
		Entry:    "index.mjs",
		Type:     "aws",
	}

	archivePath, err := utils.CreateTemporaryFolder(nil, nil)
	if err != nil {
		fmt.Printf("An error occurred while trying to create a temporary folder %v\n", err)
		return "", ServerlessFunctionState{}, err
	}

	bundleFunctionScript := fmt.Sprintf("genezio bundleFunction --functionName %s --handler %s --entry %s --functionPath %s --backendPath %s --output %s", input.Name, input.Handler, input.Entry, relFunctionPath, absolueBackendPath, archivePath)
	err = utils.RunScriptsInDirectory(absolueBackendPath, []string{bundleFunctionScript}, nil)
	if err != nil {
		fmt.Printf("An error occurred while trying to bundle the function %v", err)
		return "", ServerlessFunctionState{}, err
	}

	cloudAdapter := ca.NewGenezioCloudAdapter()

	response, err := cloudAdapter.DeployFunction(ctx, input.Project.Name, input.Project.Region, functionConfiguration, archivePath, stage)
	if err != nil {
		fmt.Printf("An error occurred while trying to deploy the function %v", err)
		return "", ServerlessFunctionState{}, err
	}

	state.ID = response.Id
	state.URL = response.CloudURL

	err = utils.DeleteTemporaryFolder()
	if err != nil {
		log.Println("Error deleting temporary folder", err)
		return "", state, err
	}

	return name, state, nil
}

func (*ServerlessFunction) Delete(ctx p.Context, id string, state ServerlessFunctionState) error {
	_, err := requests.DeleteFunction(ctx, state.ID)
	if err != nil {
		if strings.Contains(err.Error(), "record not found") {
			log.Println("Function is already deleted")
			return nil
		}
		log.Println("Error deleting function", err.Error())
		return err
	}
	return nil
}
