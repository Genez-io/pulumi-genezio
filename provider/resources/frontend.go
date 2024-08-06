package resources

import (
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
	Project   domain.Project   `pulumi:"project"`
	Path      string           `pulumi:"path"`
	Subdomain *string          `pulumi:"subdomain,optional"`
	Publish   resource.Archive `pulumi:"publish"`
}

type FrontendState struct {
	FrontendArgs

	URL string `pulumi:"url"`
}

func (*Frontend) Diff(ctx p.Context, id string, olds FrontendState, news FrontendArgs) (p.DiffResponse, error) {
	diff := map[string]p.PropertyDiff{}

	areProjectsIdentical := utils.CompareProjects(olds.Project, news.Project)
	if !areProjectsIdentical {
		diff["project"] = p.PropertyDiff{Kind: p.DeleteReplace}
	}

	if olds.Path != news.Path {
		diff["path"] = p.PropertyDiff{Kind: p.DeleteReplace}
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

	if olds.Publish.Hash != news.Publish.Hash {
		diff["publish"] = p.PropertyDiff{Kind: p.DeleteReplace}
	}

	return p.DiffResponse{
		DeleteBeforeReplace: true,
		HasChanges:          len(diff) > 0,
		DetailedDiff:        diff,
	}, nil
}

func (*Frontend) Read(ctx p.Context, id string, inputs FrontendArgs, state FrontendState) (string, FrontendArgs, FrontendState, error) {

	projectDetails, err := requests.GetProjectDetails(ctx, inputs.Project.Name)
	if err != nil {
		if strings.Contains(err.Error(), "record not found") {
			return id, inputs, FrontendState{}, nil
		}
		return "", FrontendArgs{}, FrontendState{}, fmt.Errorf("error getting project details: %v", err)
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
		return id, inputs, FrontendState{}, nil
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
			return id, inputs, FrontendState{}, nil
		}
	} else {
		subdomain = *inputs.Subdomain
	}
	for _, frontend := range frontends.List {
		if frontend.GenezioDomain == subdomain {
			state.Project = inputs.Project
			state.Path = inputs.Path
			state.Subdomain = &subdomain
			state.Publish = inputs.Publish
			return id, inputs, state, nil
		}
	}
	return id, inputs, FrontendState{}, nil
}

func (*Frontend) Create(ctx p.Context, name string, input FrontendArgs, preview bool) (string, FrontendState, error) {

	// TODO Will need to investigate further why this is needed, For now this is needed for the FileArchive to work
	// More info here https://pulumi-developer-docs.readthedocs.io/en/latest/architecture/deployment-schema.html#dabf18193072939515e22adb298388d-required
	input.Publish.Sig = "0def7320c3a5731c473e5ecbe6d01bc7"

	state := FrontendState{FrontendArgs: input}
	if preview {
		return name, state, nil
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

	var absolueFrontendPath string
	frontendPath, err := os.Getwd()
	if err != nil {
		fmt.Printf("An error occurred while trying to get the current working directory %v", err)
		return "", FrontendState{}, err
	}
	absolueFrontendPath = frontendPath

	absolueFrontendPath = filepath.Join(absolueFrontendPath, input.Path)

	relPublishPath, err := filepath.Rel(absolueFrontendPath, input.Publish.Path)
	if err != nil {
		return "", FrontendState{}, err
	}

	if _, err := os.Stat(input.Publish.Path); os.IsNotExist(err) {
		return "", FrontendState{}, fmt.Errorf("publish folder does not exist")
	}

	dir, err := os.ReadDir(input.Publish.Path)
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
		Path:      absolueFrontendPath,
		Subdomain: *input.Subdomain,
		Publish:   relPublishPath,
	}

	response, err := cloudAdapter.DeployFrontend(ctx, input.Project.Name, input.Project.Region, frontendConfiguration, stage)
	if err != nil {
		fmt.Printf("An error occurred while trying to deploy the frontend %v\n", err)
		return "", FrontendState{}, err
	}

	state.URL = response

	err = utils.DeleteTemporaryFolder()
	if err != nil {
		log.Println("Error deleting temporary folder", err)
		return "", state, err
	}

	return name, state, nil
}

func (*Frontend) Delete(ctx p.Context, id string, state FrontendState) error {
	projectDetails, err := requests.GetProjectDetails(ctx, state.Project.Name)
	if err != nil {
		if strings.Contains(err.Error(), "405 Method Not Allowed") {
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
