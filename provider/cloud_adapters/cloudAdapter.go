package cloud_adapters

import (
	"fmt"
	"path/filepath"

	"github.com/Genez-io/pulumi-genezio/provider/domain"
	"github.com/Genez-io/pulumi-genezio/provider/requests"
	"github.com/Genez-io/pulumi-genezio/provider/utils"

	p "github.com/pulumi/pulumi-go-provider"
)

<<<<<<< HEAD
type CloudAdapter interface{
	Deploy(ctx p.Context, input []domain.GenezioCloudInput, projectConfiguration domain.ProjectConfiguration, cloudAdapterOptions CloudAdapterOptions, stack *string) (domain.GenezioCloudOutput, error)
	DeployFrontend(ctx p.Context,projectName string, projectRegion string, frontend domain.FrontendConfiguration,stage string) (string, error)
=======
type CloudAdapter interface {
	Deploy(input []domain.GenezioCloudInput, projectConfiguration domain.ProjectConfiguration, cloudAdapterOptions CloudAdapterOptions, stack *string, authToken string) (domain.GenezioCloudOutput, error)
	DeployFrontend(projectName string, projectRegion string, frontend domain.FrontendConfiguration, stage string, authToken string) (string, error)
>>>>>>> cf969b7 (Fix linting warnings and remove logs)
}

type genezioCloudAdapter struct {
}

func NewGenezioCloudAdapter() CloudAdapter {
	return &genezioCloudAdapter{}
}

<<<<<<< HEAD
func (g *genezioCloudAdapter) Deploy(ctx p.Context, input []domain.GenezioCloudInput, projectConfiguration domain.ProjectConfiguration, cloudAdapterOptions CloudAdapterOptions, stack *string) (domain. GenezioCloudOutput, error) {
=======
func (g *genezioCloudAdapter) Deploy(input []domain.GenezioCloudInput, projectConfiguration domain.ProjectConfiguration, cloudAdapterOptions CloudAdapterOptions, stack *string, authToken string) (domain.GenezioCloudOutput, error) {
>>>>>>> cf969b7 (Fix linting warnings and remove logs)

	stage := ""
	if cloudAdapterOptions.Stage != nil {
		stage = *cloudAdapterOptions.Stage
	}
<<<<<<< HEAD

		for _, element := range input {
			presignedUrl,err := requests.GetPresignedUrl(ctx, projectConfiguration.Region, "genezioDeploy.zip",projectConfiguration.Name, element.Name)
			if err != nil {
				fmt.Printf("An error occurred while trying to get the presigned url %v\n", err)
				return domain.GenezioCloudOutput{}, err
			}



			err = requests.UploadContentToS3(&presignedUrl, element.ArchivePath, nil)
			if err != nil {
				fmt.Printf("An error occurred while trying to upload the content to S3 %v\n", err)
				return domain.GenezioCloudOutput{}, err
			}


		}

		response, err:= requests.DeployRequest(ctx,projectConfiguration, input, stage, nil)
=======

	for _, element := range input {
		presignedUrl, err := requests.GetPresignedUrl(projectConfiguration.Region, "genezioDeploy.zip", projectConfiguration.Name, element.Name, authToken)
>>>>>>> cf969b7 (Fix linting warnings and remove logs)
		if err != nil {
			fmt.Printf("An error occurred while trying to get the presigned url %v\n", err)
			return domain.GenezioCloudOutput{}, err
		}

		err = requests.UploadContentToS3(&presignedUrl, element.ArchivePath, nil)
		if err != nil {
			fmt.Printf("An error occurred while trying to upload the content to S3 %v\n", err)
			return domain.GenezioCloudOutput{}, err
		}

	}

	response, err := requests.DeployRequest(projectConfiguration, input, stage, nil, authToken)
	if err != nil {
		fmt.Printf("An error occurred while trying to deploy the request %v\n", err)
		return domain.GenezioCloudOutput{}, err
	}

	return domain.GenezioCloudOutput{
		ProjectID:    response.ProjectID,
		ProjectEnvID: response.ProjectEnvID,
		Classes:      response.Classes,
		Functions:    response.Functions,
	}, nil
}

func (g *genezioCloudAdapter) DeployFrontend(ctx p.Context, projectName string, projectRegion string, frontend domain.FrontendConfiguration, stage string) (string, error) {

	var finalStageName string
	if stage != "" && stage != "prod" {
		finalStageName = fmt.Sprintf("-%s", stage)
	} else {
		finalStageName = ""
	}

	finalSubdomain := fmt.Sprintf("%s%s", frontend.Subdomain, finalStageName)

	temporaryFolder, err := utils.CreateTemporaryFolder(nil, nil)
	if err != nil {
		fmt.Printf("An error occurred while trying to create a temporary folder %v\n", err)
		return "", err
	}

	archivePath := filepath.Join(temporaryFolder, fmt.Sprintf("%s.zip", finalSubdomain))
	if frontend.Publish == "" {
		frontend.Publish = "."
	}
	frontendPath := filepath.Join(frontend.Path, frontend.Publish)

	exclussionList := []string{".git", ".github"}
	err = utils.ZipDirectoryToDestinationPath(frontendPath, finalSubdomain, archivePath, exclussionList)
	if err != nil {
		fmt.Printf("An error occurred while trying to zip the directory %v\n", err)
		return "", err
	}

	presignedUrl, err := requests.GetFrontendPresignedUrl(ctx, finalSubdomain,projectName,stage)
	if err != nil {
		fmt.Printf("An error occurred while trying to get the presigned url %v\n", err)
		return "", err
	}

	err = requests.UploadContentToS3(&presignedUrl.PresignedURL, archivePath, &presignedUrl.UserID)
	if err != nil {
		fmt.Printf("An error occurred while trying to upload the content to S3 %v\n", err)
		return "", err
	}

	finalDomain, err := requests.CreateFrontendProject(ctx, finalSubdomain,projectName,projectRegion,stage)
	if err != nil {
		fmt.Printf("An error occurred while trying to create the frontend project %v\n", err)
		return "", err
	}

	return finalDomain, nil
}

type CloudAdapterOptions struct {
	Stage *string `pulumi:"stage"`
}
