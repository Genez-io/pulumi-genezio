package requests

import (
	"net/http"

	"github.com/Genez-io/pulumi-genezio/provider/domain"
)

func CreateDatabase(request domain.CreateDatabaseRequest, authToken string) (domain.CreateDatabaseResponse, error) {
	var response domain.CreateDatabaseResponse
	err := MakeRequest(http.MethodPost, "databases", request, &response, authToken)

	return response, err
}

func ListDatabases(authToken string) ([]domain.DatabaseDetails, error) {
	var response domain.GetDatabaseResponse
	err := MakeRequest(http.MethodGet, "databases", nil, &response, authToken)

	return response.Databases, err
}

func GetDatabaseConnectionUrl(databaseId string, authToken string) (string, error) {
	var response domain.GetDatabaseConnectionUrlResponse
	err := MakeRequest(http.MethodGet, "databases/"+databaseId, nil, &response, authToken)

	return response.ConnectionUrl, err
}
