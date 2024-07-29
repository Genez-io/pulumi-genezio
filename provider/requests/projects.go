package requests

import (
	"net/http"

	"github.com/Genez-io/pulumi-genezio/provider/domain"
	p "github.com/pulumi/pulumi-go-provider"
)

func DeleteProject(ctx p.Context, id string) (domain.DeleteProjectResponse, error) {
	var response domain.DeleteProjectResponse
	err := MakeRequest(ctx, http.MethodGet, "projects/"+id, nil, &response)
	return response, err
}

func GetProjectDetails(ctx p.Context, name string) (domain.ProjectDetailsResponse, error) {
	var response domain.ProjectDetailsResponse
	err := MakeRequest(ctx, http.MethodGet, "projects/name/"+name, nil, &response)

	return response, err
}
