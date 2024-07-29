package requests

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/Genez-io/pulumi-genezio/provider/constants"
	"github.com/Genez-io/pulumi-genezio/provider/domain"
	"github.com/Genez-io/pulumi-genezio/provider/utils"
	p "github.com/pulumi/pulumi-go-provider"
)

func GetDatabaseConnectionUrl(
	ctx p.Context,
	databaseId string,
) (string, error) {

	if databaseId == "" {
		return "", fmt.Errorf("databaseId is required")
	}

	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/databases/%s", constants.API_URL, databaseId), nil)
	if err != nil {
		return "", err
	}

	authToken, err := utils.GetAuthToken(ctx)
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+authToken)
	req.Header.Set("Accept-Version", "genezio-webapp/0.3.0")

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("error: %s", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var data domain.GetDatabaseConnectionUrlResponse
	err = json.Unmarshal(body, &data)
	if err != nil {
		return "", err
	}

	return data.ConnectionUrl,nil

}
