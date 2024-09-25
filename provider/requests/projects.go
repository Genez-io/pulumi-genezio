package requests

import (
	"context"
	"net/http"

	"github.com/Genez-io/pulumi-genezio/provider/domain"
)

func DeleteProject(ctx context.Context, id string) (domain.DeleteProjectResponse, error) {
	var response domain.DeleteProjectResponse
	err := MakeRequest(ctx, http.MethodDelete, "projects/"+id, nil, &response)
	return response, err
}

func GetProjectDetails(ctx context.Context, name string) (domain.ProjectDetailsResponse, error) {
	var response domain.ProjectDetailsResponse
	err := MakeRequest(ctx, http.MethodGet, "projects/name/"+name, nil, &response)

	return response, err
}
