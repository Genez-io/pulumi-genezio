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
)

type Frontend struct{}

type FrontendArgs struct {
	ProjectName string `pulumi:"projectName"`
	Region 	string `pulumi:"region"`
	Stage 	*string `pulumi:"stage,optional"`
	Path string `pulumi:"path"`
	Subdomain *string `pulumi:"subdomain,optional"`
	Publish string `pulumi:"publish"`
}

type FrontendState struct {
	FrontendArgs

	URL string `pulumi:"url"`
}

func (*Frontend) Diff(ctx p.Context, id string, olds FrontendState, news FrontendArgs) (p.DiffResponse, error) {
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

	if olds.Publish != news.Publish {
		diff["publish"] = p.PropertyDiff{Kind: p.DeleteReplace}
	}

	return p.DiffResponse{
		DeleteBeforeReplace: true,
		HasChanges: 		len(diff) > 0,
		DetailedDiff: 		diff,
	}, nil
}

func (*Frontend) Read(ctx p.Context, id string, inputs FrontendArgs, state FrontendState) (string, FrontendArgs ,FrontendState, error) {

	projectDetails,err := requests.GetProjectDetails(ctx, inputs.ProjectName)
	if err != nil {
		if strings.Contains(err.Error(), "record not found") {
			return id, inputs, FrontendState{}, nil
		}
		return "", FrontendArgs{}, FrontendState{}, fmt.Errorf("error getting project details: %v", err)
	}

	stage := "prod"
	if inputs.Stage != nil {
		stage = *inputs.Stage
	}

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
			state.ProjectName = inputs.ProjectName
			state.Region = inputs.Region
			state.Stage = inputs.Stage
			state.Path = inputs.Path
			state.Subdomain = &subdomain
			state.Publish = inputs.Publish
			return id, inputs, state, nil
		}
	}
	return id, inputs, FrontendState{}, nil 
}


func (*Frontend) Create(ctx p.Context, name string, input FrontendArgs, preview bool) (string, FrontendState, error) {

	state := FrontendState{FrontendArgs: input}
	if preview {
		return name, state, nil
	}

	stage := ""
	if input.Stage != nil {
		stage = *input.Stage
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

	frontendPath := filepath.Join(input.Path, input.Publish)

	if _, err := os.Stat(frontendPath); os.IsNotExist(err) {
		return "", FrontendState{}, fmt.Errorf("publish folder does not exist")
	}

	dir, err := os.ReadDir(frontendPath)
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
		Path: input.Path,
		Subdomain: *input.Subdomain,
		Publish: input.Publish,
	}

	response, err := cloudAdapter.DeployFrontend(ctx, input.ProjectName,input.Region, frontendConfiguration, stage,)
	if err != nil {
		fmt.Printf("An error occurred while trying to deploy the frontend %v\n", err)
		return "", FrontendState{}, err
	}

	state.URL = response
	return name, state, nil
}

func (*Frontend) Delete(ctx p.Context, id string, state FrontendState) error {
	projectDetails, err := requests.GetProjectDetails(ctx, state.ProjectName)
	if err != nil {
		if strings.Contains(err.Error(), "record not found") {
			return  nil
		}
		log.Println("Error getting project details", err)
		return err
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
				log.Println("Error deleting frontend", err)
				return err
			}
		}
	}

	return nil
}

