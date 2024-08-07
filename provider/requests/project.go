package requests

import (
	"fmt"
	"net/http"

	"github.com/Genez-io/pulumi-genezio/provider/domain"
	p "github.com/pulumi/pulumi-go-provider"
)

func CreateProject(ctx p.Context, request domain.CreateProjectRequest) (domain.CreateProjectResponse, error) {

	var response domain.DeployCodeResponse
	err := MakeRequest(ctx, http.MethodPut, "core/deployment", request, &response)

	return domain.CreateProjectResponse{
		ProjectEnvID: response.ProjectEnvID,
		ProjectID:    response.ProjectID,
	}, err

}

func DeployRequest(ctx p.Context, request domain.DeployRequest) (domain.DeployCodeResponse, error) {
	var response domain.DeployCodeResponse
	err := MakeRequest(ctx, http.MethodPut, "core/deployment", request, &response)

	return response, err
}

func GetPresignedUrl(
	ctx p.Context,
	request domain.GetPresignedUrlRequest,
) (domain.GetPresignedUrlResponse, error) {

	if request.Region == "" || request.Filename == "" || request.ProjectName == "" || request.ClassName == "" {
		return domain.GetPresignedUrlResponse{}, fmt.Errorf("invalid request to get presigned url")
	}

	var response domain.GetPresignedUrlResponse
	err := MakeRequest(ctx, http.MethodGet, "core/deployment-url", request, &response)

	return response, err
}
