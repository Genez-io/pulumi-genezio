package utils

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
			fmt.Println("Uploading to S3 1")
			presignedUrl,err := requests.GetPresignedUrl(projectConfiguration.Region, "genezioDeploy.zip",projectConfiguration.Name, element.Name, authToken)
			if err != nil {
				return domain.GenezioCloudOutput{}, err
			}
			fmt.Printf("Uploading to S3 2 %s", presignedUrl)

			err = requests.UploadContentToS3(&presignedUrl, element.ArchivePath, nil)
			if err != nil {
				return domain.GenezioCloudOutput{}, err
			}
			fmt.Println("Uploading to S3 3")
		} 
	
		fmt.Println("Uploading to S3 4")
		response, err:= requests.DeployRequest(projectConfiguration, input, stage, nil, authToken) 
		if err != nil {
			return domain.GenezioCloudOutput{}, err
		}
		fmt.Println("Uploading to S3 5")

		fmt.Println(response)

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

