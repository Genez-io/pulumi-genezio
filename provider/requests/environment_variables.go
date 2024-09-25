package requests

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Genez-io/pulumi-genezio/provider/domain"
)

func SetEnvironmentVariables(
	ctx context.Context,
	projectId string,
	projectEnvId string,
	request domain.SetEnvironmentVariablesRequest,
) error {
	err := MakeRequest(ctx, http.MethodPost, fmt.Sprintf("projects/%s/%s/environment-variables", projectId, projectEnvId), request, nil)
	return err
}
