package requests

import (
	"net/http"

	"github.com/Genez-io/pulumi-genezio/provider/domain"
	p "github.com/pulumi/pulumi-go-provider"
)

func SetAuthentication(ctx p.Context, envId string, request domain.SetAuthenticationRequest) (domain.SetAuthenticationResponse, error) {
	var response domain.SetAuthenticationResponse
	err := MakeRequest(ctx, http.MethodPut, "core/auth/"+envId, request, &response)
	return response, err
}

func GetAuthentication(ctx p.Context, envId string) (domain.GetAuthenticationResponse, error) {
	var response domain.GetAuthenticationResponse
	err := MakeRequest(ctx, http.MethodGet, "core/auth/"+envId, nil, &response)
	return response, err
}

func GetAuthProviders(ctx p.Context, envId string) (domain.GetAuthProvidersResponse, error) {
	var response domain.GetAuthProvidersResponse
	err := MakeRequest(ctx, http.MethodGet, "core/auth/providers/"+envId, nil, &response)
	return response, err
}

func SetAuthProviders(ctx p.Context, envId string, request domain.SetAuthProvidersRequest) (domain.SetAuthProvidersResponse, error) {
	var response domain.SetAuthProvidersResponse
	err := MakeRequest(ctx, http.MethodPut, "core/auth/providers/"+envId, request, &response)
	return response, err
}
