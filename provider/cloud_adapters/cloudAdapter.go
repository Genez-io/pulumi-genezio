package cloud_adapters

import (
	"fmt"
	"path/filepath"

	"github.com/Genez-io/pulumi-genezio/provider/domain"
	"github.com/Genez-io/pulumi-genezio/provider/requests"
	"github.com/Genez-io/pulumi-genezio/provider/utils"

	p "github.com/pulumi/pulumi-go-provider"
)

type CloudAdapter interface {
	Deploy(ctx p.Context, input []domain.GenezioCloudInput, projectConfiguration domain.ProjectConfiguration, cloudAdapterOptions CloudAdapterOptions, stack *string) (domain.GenezioCloudOutput, error)
	DeployFrontend(ctx p.Context, projectName string, projectRegion string, frontend domain.FrontendConfiguration, stage string) (string, error)
}

type genezioCloudAdapter struct {
}

func NewGenezioCloudAdapter() CloudAdapter {
	return &genezioCloudAdapter{}
}

func (g *genezioCloudAdapter) Deploy(ctx p.Context, input []domain.GenezioCloudInput, projectConfiguration domain.ProjectConfiguration, cloudAdapterOptions CloudAdapterOptions, stack *string) (domain.GenezioCloudOutput, error) {

	stage := ""
	if cloudAdapterOptions.Stage != nil {
		stage = *cloudAdapterOptions.Stage
	}

	
		for _, element := range input {
			presignedUrlResponse,err := requests.GetPresignedUrl(ctx, domain.GetPresignedUrlRequest{
				ProjectName: projectConfiguration.Name,
				Region: projectConfiguration.Region,
				Filename: "genezioDeploy.zip",
				ClassName: element.Name,
			})
			if err != nil {
				fmt.Printf("An error occurred while trying to get the presigned url %v\n", err)
				return domain.GenezioCloudOutput{}, err
			}



			err = requests.UploadContentToS3(&presignedUrlResponse.PresignedUrl, element.ArchivePath, nil)
			if err != nil {
				fmt.Printf("An error occurred while trying to upload the content to S3 %v\n", err)
				return domain.GenezioCloudOutput{}, err
			}

			
		} 
	
		mappedFunctions := []domain.MappedFunction{}

		for _, fun := range projectConfiguration.Functions {
			entryFile := ""
			for _, input := range input {
				if input.Name == fun.Name {
					entryFile = input.EntryFile
					break
				}
			}

			mappedFunctions = append(mappedFunctions, domain.MappedFunction{
				Name: fun.Name,
				Language: fun.Language,
				EntryFile: entryFile,
			})
		}

		response, err:= requests.DeployRequest(ctx, domain.DeployRequest{
			Options: projectConfiguration.Options,
			Classes: projectConfiguration.Classes,
			Functions: mappedFunctions,
			ProjectName: projectConfiguration.Name,
			Region: projectConfiguration.Region,
			CloudProvider: projectConfiguration.CloudProvider,
			Stage: stage,
			Stack: nil,
		}) 
		if err != nil {
			fmt.Printf("An error occurred while trying to upload the content to S3 %v\n", err)
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

	presignedUrl, err := requests.GetFrontendPresignedUrl(ctx, domain.GetFrontendPresignedUrlRequest{
		SubdomainName: finalSubdomain,
		ProjectName: projectName,
		Region: projectRegion,
		Stage: stage,
	})
	if err != nil {
		fmt.Printf("An error occurred while trying to get the presigned url %v\n", err)
		return "", err
	}

	err = requests.UploadContentToS3(&presignedUrl.PresignedURL, archivePath, &presignedUrl.UserID)
	if err != nil {
		fmt.Printf("An error occurred while trying to upload the content to S3 %v\n", err)
		return "", err
	}


	createFrontendResponse, err := requests.CreateFrontendProject(ctx, domain.CreateFrontendProjectRequest{
		ProjectName: projectName,
		Region: projectRegion,
		GenezioDomain: finalSubdomain,
		Stage: stage,
	})
	if err != nil {
		fmt.Printf("An error occurred while trying to create the frontend project %v\n", err)
		return "", err
	}
	
	fmt.Printf("Frontend deployed successfully at %s\n", createFrontendResponse.Domain)
	return createFrontendResponse.Domain, nil
}

type CloudAdapterOptions struct {
	Stage *string `pulumi:"stage"`
}
