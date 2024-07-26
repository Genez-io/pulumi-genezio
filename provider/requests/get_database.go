package requests

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/Genez-io/pulumi-genezio/provider/constants"
	"github.com/Genez-io/pulumi-genezio/provider/domain"
)

func GetDatabase(id string, authToken string) (domain.GetDatabaseResponse, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/databases/%s", constants.API_URL, id), nil)
	if err != nil {
		return domain.GetDatabaseResponse{}, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+authToken)
	req.Header.Set("Accept-Version", "genezio-webapp/0.3.0")

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return domain.GetDatabaseResponse{}, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return domain.GetDatabaseResponse{}, fmt.Errorf("error: %s", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return domain.GetDatabaseResponse{}, err
	}

	var data domain.GetDatabaseResponse = domain.GetDatabaseResponse{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		return domain.GetDatabaseResponse{}, err
	}

	defer resp.Body.Close()

	return data, nil

}
