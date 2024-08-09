package requests

import (
	"net/http"

	"github.com/Genez-io/pulumi-genezio/provider/domain"
	p "github.com/pulumi/pulumi-go-provider"
)

func CreateFunction(ctx p.Context, request domain.CreateFunctionRequest) (domain.GetFunctionResponse, error) {

	var response domain.GetFunctionResponse
	err := MakeRequest(ctx, http.MethodPost, "functions", request, &response)

	return response, err
}

func GetFunction(ctx p.Context, id string) (domain.GetFunctionResponse, error) {
	var response domain.GetFunctionResponse
	err := MakeRequest(ctx, http.MethodGet, "functions/"+id, nil, &response)
	return response, err
}

func DeleteFunction(ctx p.Context, id string) (domain.DeleteFunctionResponse, error) {
	var response domain.DeleteFunctionResponse
	err := MakeRequest(ctx, http.MethodDelete, "functions/"+id, nil, &response)
	return response, err
}
