package requests

import (
	"fmt"
	"net/http"

	"github.com/Genez-io/pulumi-genezio/provider/domain"

	p "github.com/pulumi/pulumi-go-provider"
)

func SetEnvironmentVariables(
	ctx p.Context,
	projectId string,
	projectEnvId string,
	request domain.SetEnvironmentVariablesRequest,
) error {
	err := MakeRequest(ctx, http.MethodPost, fmt.Sprintf("projects/%s/%s/environment-variables", projectId, projectEnvId), request, nil)
	return err
}
