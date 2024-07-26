package requests

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/Genez-io/pulumi-genezio/provider/constants"
	"github.com/Genez-io/pulumi-genezio/provider/domain"
	p "github.com/pulumi/pulumi-go-provider"
)

func ListDatabases(ctx p.Context) ([]domain.DatabaseDetails, error) {

	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/databases", constants.API_URL), nil)
	if err != nil {
		return []domain.DatabaseDetails{}, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+ctx.Value("authToken").(string))
	req.Header.Set("Accept-Version", "genezio-webapp/0.3.0")

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return []domain.DatabaseDetails{}, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return []domain.DatabaseDetails{}, fmt.Errorf("error: %s", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return []domain.DatabaseDetails{}, err
	}

	var data domain.GetDatabasesResponse
	err = json.Unmarshal(body, &data)
	if err != nil {
		return []domain.DatabaseDetails{}, err
	}

	return data.Databases, nil


}