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

func CreateDatabase(
	ctx p.Context,
	DbType string,
	region string,
	name string,
) (domain.CreateDatabaseResponse, error) {
	
	type request struct {
		DbType string `json:"type"`
		Region string `json:"region"`
		Name string `json:"name"`
	}

	data := request{
		DbType: DbType,
		Region: region,
		Name: name,
	}

	jsonData,err := json.Marshal(data)
	if err != nil {
		return domain.CreateDatabaseResponse{}, err
	}

	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/databases",constants.API_URL), bytes.NewBuffer(jsonData))
	if err != nil {
		return domain.CreateDatabaseResponse{}, err
	}


	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+ctx.Value("authToken").(string))
	req.Header.Set("Accept-Version", "genezio-webapp/0.3.0")

	client := &http.Client{
	}


	resp, err := client.Do(req)
	if err != nil {
		return domain.CreateDatabaseResponse{}, err
	}


	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return domain.CreateDatabaseResponse{}, err
	}


	if resp.StatusCode != http.StatusOK {
		return domain.CreateDatabaseResponse{}, fmt.Errorf("error: %s and response %v", string(body), resp)
	}


	var dataResponse domain.CreateDatabaseResponse
	err = json.Unmarshal(body, &dataResponse)
	if err != nil {
		return domain.CreateDatabaseResponse{}, err
	}

	return dataResponse, nil
}