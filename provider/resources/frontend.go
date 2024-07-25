package resources

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"

	ca "github.com/Genez-io/pulumi-genezio/provider/cloud_adapters"
	"github.com/Genez-io/pulumi-genezio/provider/domain"
	"github.com/Genez-io/pulumi-genezio/provider/utils"
	p "github.com/pulumi/pulumi-go-provider"
)

type Frontend struct{}

type FrontendArgs struct {
	ProjectName string `pulumi:"projectName"`
	Region 	string `pulumi:"region"`
	AuthToken string `pulumi:"authToken"`
	Path string `pulumi:"path"`
	Subdomain *string `pulumi:"subdomain,optional"`
	Publish string `pulumi:"publish"`
}

type FrontendState struct {
	FrontendArgs

	URL string `pulumi:"url"`
}



func (*Frontend) Create(ctx p.Context, name string, input FrontendArgs, preview bool) (string, FrontendState, error) {

	state := FrontendState{FrontendArgs: input}
	if preview {
		return name, state, nil
	}

	stage := ""

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
	}

	cloudAdapter := ca.NewGenezioCloudAdapter()

	frontendConfiguration := domain.FrontendConfiguration{
		Path: input.Path,
		Subdomain: *input.Subdomain,
		Publish: input.Publish,
	}

	response, err := cloudAdapter.DeployFrontend(input.ProjectName,input.Region, frontendConfiguration, stage, input.AuthToken)
	if err != nil {
		fmt.Printf("An error occurred while trying to deploy the frontend %v\n", err)
		return "", FrontendState{}, err
	}

	state.URL = response
	return name, state, nil
}