package requests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/Genez-io/pulumi-genezio/provider/constants"
	"github.com/Genez-io/pulumi-genezio/provider/domain"

	p "github.com/pulumi/pulumi-go-provider"
)

func SetEnvironmentVariables(
	ctx p.Context,
	projectId string,
	projectEnvId string,
	environmentVariablesData []domain.EnvironmentVariable,
) error {
	if projectId == "" || projectEnvId == "" {
		return fmt.Errorf("projectId and ProjectEnvId is required")
	}

	type request struct {
		EnvironmentVariables []domain.EnvironmentVariable `json:"environmentVariables"`
	}

	data := request{
		EnvironmentVariables: environmentVariablesData,
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/projects/%s/%s/environment-variables", constants.API_URL, projectId, projectEnvId), bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}

	req.Header.Set("Authorization", "Bearer "+ctx.Value("authToken").(string))
	req.Header.Set("Accept-Version", "genezio-cli/2.2.0")

	client := &http.Client{
	}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("error: %s", string(body))
	}

	return nil

}
