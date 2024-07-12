package cloud_adapters

import (
	"fmt"

	"github.com/Genez-io/pulumi-genezio/provider/domain"
	"github.com/Genez-io/pulumi-genezio/provider/requests"
)

type CloudAdapter interface{
	Deploy(input []domain.GenezioCloudInput, projectConfiguration domain.ProjectConfiguration, cloudAdapterOptions CloudAdapterOptions, stack *string, authToken string) (domain.GenezioCloudOutput, error)
}

type genezioCloudAdapter struct {

}

func NewGenezioCloudAdapter() CloudAdapter {
	return &genezioCloudAdapter{}
}

func (g *genezioCloudAdapter) Deploy(input []domain.GenezioCloudInput, projectConfiguration domain.ProjectConfiguration, cloudAdapterOptions CloudAdapterOptions, stack *string, authToken string) (domain. GenezioCloudOutput, error) {

	stage := ""
	if cloudAdapterOptions.Stage != nil {
		stage = *cloudAdapterOptions.Stage
	}
	
		for _, element := range input {
			presignedUrl,err := requests.GetPresignedUrl(projectConfiguration.Region, "genezioDeploy.zip",projectConfiguration.Name, element.Name, authToken)
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
	
		response, err:= requests.DeployRequest(projectConfiguration, input, stage, nil, authToken) 
		if err != nil {
			fmt.Printf("An error occurred while trying to deploy the request %v\n", err)
			return domain.GenezioCloudOutput{}, err
		}

	return domain.GenezioCloudOutput{
		ProjectID: response.ProjectID,
		ProjectEnvID: response.ProjectEnvID,
		Classes: response.Classes,
		Functions: response.Functions,
	}, nil
}

type CloudAdapterOptions struct {
	Stage *string `pulumi:"stage"`
}

