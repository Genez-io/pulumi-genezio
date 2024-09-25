package resources

import (
	"context"
	_ "embed"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	ca "github.com/Genez-io/pulumi-genezio/provider/cloud_adapters"
	"github.com/Genez-io/pulumi-genezio/provider/domain"
	"github.com/Genez-io/pulumi-genezio/provider/requests"
	"github.com/Genez-io/pulumi-genezio/provider/utils"
	p "github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi-go-provider/infer"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
)

type Frontend struct{}

type FrontendArgs struct {
	Project       domain.Project                `pulumi:"project"`
	Path          resource.Archive              `pulumi:"path"`
	Subdomain     *string                       `pulumi:"subdomain,optional"`
	Publish       string                        `pulumi:"publish"`
	BuildCommands *[]string                     `pulumi:"buildCommands,optional"`
	Environment   *[]domain.EnvironmentVariable `pulumi:"environment,optional"`
}

type FrontendState struct {
	FrontendArgs

	URL string `pulumi:"url"`
}

//go:embed documentation/project.md
var frontendDocumentation string

func (r *Frontend) Annotate(a infer.Annotator) {
	a.Describe(&r, frontendDocumentation)
}

func (r *FrontendArgs) Annotate(a infer.Annotator) {
	a.Describe(&r.Project, `The project to which the frontend will be deployed.`)
	a.Describe(&r.Path, `The path to the frontend files.`)
	a.Describe(&r.Subdomain, `The subdomain of the frontend.`)
	a.Describe(&r.Publish, `The folder in the path that contains the files to be published.`)
	a.Describe(&r.BuildCommands, `The commands to run before deploying the frontend.`)
	a.Describe(&r.Environment, `The environment variables that will be set for the frontend.`)
}

func (r *FrontendState) Annotate(a infer.Annotator) {
	a.Describe(&r.URL, `The URL of the frontend.`)
}

func (*Frontend) Diff(ctx context.Context, id string, olds FrontendState, news FrontendArgs) (p.DiffResponse, error) {
	diff := map[string]p.PropertyDiff{}

	areProjectsIdentical := utils.CompareProjects(olds.Project, news.Project)
	if !areProjectsIdentical {
		diff["project"] = p.PropertyDiff{Kind: p.DeleteReplace}
	}

	if olds.Publish != news.Publish {
		diff["path"] = p.PropertyDiff{Kind: p.Update}
	}

	if olds.Subdomain == nil {
		if news.Subdomain != nil {
			diff["subdomain"] = p.PropertyDiff{Kind: p.DeleteReplace}
		}
	} else {
		if news.Subdomain != nil && *olds.Subdomain != *news.Subdomain {
			diff["subdomain"] = p.PropertyDiff{Kind: p.DeleteReplace}
		}
	}

	if olds.Path.Hash != news.Path.Hash {
		diff["publish"] = p.PropertyDiff{Kind: p.Update}
	}

	if olds.BuildCommands == nil {
		if news.BuildCommands != nil {
			diff["buildCommands"] = p.PropertyDiff{Kind: p.Update}
		}
	} else {
		if news.BuildCommands != nil {
			if len(*olds.BuildCommands) != len(*news.BuildCommands) {
				diff["buildCommands"] = p.PropertyDiff{Kind: p.Update}
			} else {
				for i, buildCommand := range *news.BuildCommands {
					if (*olds.BuildCommands)[i] != buildCommand {
						diff["buildCommands"] = p.PropertyDiff{Kind: p.Update}
						break
					}
				}
			}
		} else {
			diff["buildCommands"] = p.PropertyDiff{Kind: p.Update}
		}
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

func (*Frontend) Read(ctx context.Context, id string, inputs FrontendArgs, state FrontendState) (string, FrontendArgs, FrontendState, error) {

	// TODO Investigate why is this needed - .Sig is used to recognize the asset type when unmarshalling the resource object
	// This should not be hardcoded here, but rather automatically generated by the pulumi SDK
	// Code: https://github.com/pulumi/pulumi/blob/master/sdk/go/common/resource/sig/sig.go
	// Documentation: https://pulumi-developer-docs.readthedocs.io/en/latest/architecture/deployment-schema.html#dabf18193072939515e22adb298388d-required
	inputs.Path.Sig = resource.ArchiveSig
	state.Path.Sig = resource.ArchiveSig

	if state.Project.Name == "" {
		return id, inputs, state, nil
	}

	projectDetails, err := requests.GetProjectDetails(ctx, state.Project.Name)
	if err != nil {
		if strings.Contains(err.Error(), "record not found") {
			state.Project.Name = ""
			return id, inputs, state, nil
		}
		return "", FrontendArgs{}, FrontendState{}, fmt.Errorf("error getting project details: %v", err)
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
		state.Project.Name = ""
		return id, inputs, state, nil
	}

	frontends, err := requests.GetFrontendsByEnvId(ctx, currentProjectEnv.Id)
	if err != nil {
		return "", FrontendArgs{}, FrontendState{}, fmt.Errorf("error getting frontends: %v", err)
	}

	subdomain := ""
	if inputs.Subdomain == nil {
		if state.Subdomain != nil {
			subdomain = *state.Subdomain
		} else {
			state.Project.Name = ""
			return id, inputs, state, nil
		}
	} else {
		subdomain = *inputs.Subdomain
	}
	for _, frontend := range frontends.List {
		if frontend.GenezioDomain == subdomain {
			state.Project.Name = projectDetails.Project.Name
			state.Project.Region = projectDetails.Project.Region
			state.Subdomain = &subdomain
			return id, inputs, state, nil
		}
	}
	return id, inputs, FrontendState{}, nil
}

func (*Frontend) Update(ctx context.Context, id string, olds FrontendState, news FrontendArgs, preview bool) (FrontendState, error) {

	news.Path.Sig = resource.ArchiveSig
	olds.Path.Sig = resource.ArchiveSig

	state := FrontendState{FrontendArgs: news, URL: olds.URL}

	if preview {
		return state, nil
	}

	if news.BuildCommands != nil {
		err := utils.RunScriptsInDirectory(news.Path.Path, *news.BuildCommands, news.Environment)
		if err != nil {
			return FrontendState{}, err
		}
	}

	stage := "prod"

	contextStage := infer.GetConfig[*domain.Config](ctx).Stage
	if contextStage != nil {
		stage = *contextStage
	}

	if news.Subdomain != nil {
		match, err := regexp.MatchString("^[a-z0-9-]+$", *news.Subdomain)
		if err != nil {
			return FrontendState{}, err
		}
		if !match {
			return FrontendState{}, fmt.Errorf("invalid subdomain format")
		}
	}

	absolutePublishPath := filepath.Join(news.Path.Path, news.Publish)

	if _, err := os.Stat(absolutePublishPath); os.IsNotExist(err) {
		return FrontendState{}, fmt.Errorf("publish folder does not exist")
	}

	dir, err := os.ReadDir(absolutePublishPath)
	if err != nil {
		return FrontendState{}, err
	}

	if len(dir) == 0 {
		return FrontendState{}, fmt.Errorf("publish directory is empty")
	}

	cloudAdapter := ca.NewGenezioCloudAdapter()

	frontendConfiguration := domain.FrontendConfiguration{
		Path:      news.Path.Path,
		Subdomain: *olds.Subdomain,
		Publish:   news.Publish,
	}

	response, err := cloudAdapter.DeployFrontend(ctx, news.Project.Name, news.Project.Region, frontendConfiguration, stage)
	if err != nil {
		fmt.Printf("An error occurred while trying to deploy the frontend %v\n", err)
		return FrontendState{}, err
	}

	state.URL = response

	return state, nil
}

func (*Frontend) Create(ctx context.Context, name string, input FrontendArgs, preview bool) (string, FrontendState, error) {
	// TODO Investigate why is this needed - .Sig is used to recognize the asset type when unmarshalling the resource object
	// This should not be hardcoded here, but rather automatically generated by the pulumi SDK
	// Code: https://github.com/pulumi/pulumi/blob/master/sdk/go/common/resource/sig/sig.go
	// Documentation: https://pulumi-developer-docs.readthedocs.io/en/latest/architecture/deployment-schema.html#dabf18193072939515e22adb298388d-required
	input.Path.Sig = resource.ArchiveSig

	state := FrontendState{FrontendArgs: input}
	if preview {
		return name, state, nil
	}

	if input.BuildCommands != nil {
		err := utils.RunScriptsInDirectory(input.Path.Path, *input.BuildCommands, input.Environment)
		if err != nil {
			return "", FrontendState{}, err
		}
	}

	stage := "prod"

	contextStage := infer.GetConfig[*domain.Config](ctx).Stage
	if contextStage != nil {
		stage = *contextStage
	}

	if input.Subdomain != nil {
		match, err := regexp.MatchString("^[a-z0-9-]+$", *input.Subdomain)
		if err != nil {
			return "", FrontendState{}, err
		}
		if !match {
			return "", FrontendState{}, fmt.Errorf("invalid subdomain format")
		}
	}

	absolutePublishPath := filepath.Join(input.Path.Path, input.Publish)

	if _, err := os.Stat(absolutePublishPath); os.IsNotExist(err) {
		return "", FrontendState{}, fmt.Errorf("publish folder does not exist")
	}

	dir, err := os.ReadDir(absolutePublishPath)
	if err != nil {
		return "", FrontendState{}, err
	}

	if len(dir) == 0 {
		return "", FrontendState{}, fmt.Errorf("publish directory is empty")
	}

	if input.Subdomain == nil {
		fmt.Println("No subdomain provided, we will generate a random subdomain")
		randomSubdomain := utils.GenerateRandomSubdomain()
		input.Subdomain = &randomSubdomain
		state.Subdomain = &randomSubdomain
	}

	cloudAdapter := ca.NewGenezioCloudAdapter()

	frontendConfiguration := domain.FrontendConfiguration{
		Path:      input.Path.Path,
		Subdomain: *input.Subdomain,
		Publish:   input.Publish,
	}

	response, err := cloudAdapter.DeployFrontend(ctx, input.Project.Name, input.Project.Region, frontendConfiguration, stage)
	if err != nil {
		fmt.Printf("An error occurred while trying to deploy the frontend %v\n", err)
		return "", FrontendState{}, err
	}

	state.URL = response

	return name, state, nil
}

func (*Frontend) Delete(ctx context.Context, id string, state FrontendState) error {

	if state.Project.Name == "" {
		return nil
	}

	projectDetails, err := requests.GetProjectDetails(ctx, state.Project.Name)
	if err != nil {
		if strings.Contains(err.Error(), "record not found") {
			return nil
		}
		log.Println("Error getting project details", err)
		return err
	}

	stage := "prod"

	var currentProjectEnv *domain.ProjectEnvDetails
	for _, projectEnv := range projectDetails.Project.ProjectEnvs {
		if projectEnv.Name == stage {
			currentProjectEnv = &projectEnv
			break
		}
	}

	if currentProjectEnv == nil {
		return nil
	}

	frontends, err := requests.GetFrontendsByEnvId(ctx, currentProjectEnv.Id)
	if err != nil {
		return fmt.Errorf("error getting frontends: %v", err)
	}

	for _, frontend := range frontends.List {
		if frontend.GenezioDomain == *state.Subdomain {
			err := requests.DeleteFrontend(ctx, frontend.GenezioDomain)
			if err != nil {
				if strings.Contains(err.Error(), "record not found") {
					return nil
				}
				log.Println("Error deleting frontend", err)
				return err
			}
		}
	}

	return nil
}
