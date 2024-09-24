package requests

import (
	"context"
	"net/http"

	"github.com/Genez-io/pulumi-genezio/provider/domain"
)

func CreateFrontendProject(ctx context.Context, request domain.CreateFrontendProjectRequest) (domain.CreateFrontendProjectResponse, error) {

	var response domain.CreateFrontendProjectResponse
	err := MakeRequest(ctx, http.MethodPut, "frontend", request, &response)

	return response, err
}

func GetFrontendPresignedUrl(ctx context.Context, request domain.GetFrontendPresignedUrlRequest) (domain.FrontendPresignedUrlResponse, error) {
	request.Region = "us-east-1"

	var response domain.FrontendPresignedUrlResponse
	err := MakeRequest(ctx, http.MethodGet, "core/frontend-deployment-url", request, &response)

	return response, err
}

func GetFrontendsByEnvId(ctx context.Context, envId string) (domain.GetFrontendByEnvIdResponse, error) {
	var response domain.GetFrontendByEnvIdResponse
	err := MakeRequest(ctx, http.MethodGet, "frontend/"+envId, nil, &response)

	return response, err
}

func DeleteFrontend(ctx context.Context, subdomain string) error {
	return MakeRequest(ctx, http.MethodDelete, "frontend/"+subdomain, nil, nil)
}
