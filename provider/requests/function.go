package requests

import (
	"context"
	"net/http"

	"github.com/Genez-io/pulumi-genezio/provider/domain"
)

func CreateFunction(ctx context.Context, request domain.CreateFunctionRequest) (domain.GetFunctionResponse, error) {

	var response domain.GetFunctionResponse
	err := MakeRequest(ctx, http.MethodPost, "functions", request, &response)

	return response, err
}

func GetFunction(ctx context.Context, id string) (domain.GetFunctionResponse, error) {
	var response domain.GetFunctionResponse
	err := MakeRequest(ctx, http.MethodGet, "functions/"+id, nil, &response)
	return response, err
}

func DeleteFunction(ctx context.Context, id string) (domain.DeleteFunctionResponse, error) {
	var response domain.DeleteFunctionResponse
	err := MakeRequest(ctx, http.MethodDelete, "functions/"+id, nil, &response)
	return response, err
}
