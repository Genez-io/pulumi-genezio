package requests

import (
	"context"
	"net/http"

	"github.com/Genez-io/pulumi-genezio/provider/domain"
)

func CreateDatabase(ctx context.Context, request domain.CreateDatabaseRequest) (domain.CreateDatabaseResponse, error) {
	var response domain.CreateDatabaseResponse
	err := MakeRequest(ctx, http.MethodPost, "databases", request, &response)

	return response, err
}

func DeleteDatabase(ctx context.Context, databaseId string) error {
	return MakeRequest(ctx, http.MethodDelete, "databases/"+databaseId, nil, nil)
}

func ListDatabases(ctx context.Context) ([]domain.DatabaseDetails, error) {
	var response domain.GetDatabaseResponse
	err := MakeRequest(ctx, http.MethodGet, "databases", nil, &response)

	return response.Databases, err
}

func GetDatabaseConnectionUrl(ctx context.Context, databaseId string) (string, error) {
	var response domain.GetDatabaseConnectionUrlResponse
	err := MakeRequest(ctx, http.MethodGet, "databases/"+databaseId, nil, &response)

	return response.ConnectionUrl, err
}

func LinkDatabaseToProject(ctx context.Context, request domain.LinkDatabaseToProjectRequest) (domain.LinkDatabaseToProjectResponse, error) {
	var response domain.LinkDatabaseToProjectResponse
	err := MakeRequest(ctx, http.MethodPost, "projects/"+request.ProjectId+"/"+request.StageId+"/databases/"+request.DatabaseId, nil, &response)

	return response, err
}
