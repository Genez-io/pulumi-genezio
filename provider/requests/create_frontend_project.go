package requests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/Genez-io/pulumi-genezio/provider/constants"
	"github.com/Genez-io/pulumi-genezio/provider/utils"
	p "github.com/pulumi/pulumi-go-provider"
)

func CreateFrontendProject(ctx p.Context, genezioDomain string, projectName string, region string, stage string) (string, error){

	type request struct{
		GenezioDomain string `json:"genezioDomain"`
		ProjectName string `json:"projectName"`
		Region string `json:"region"`
		Stage string `json:"stage"`
	}

	data := request{
		GenezioDomain: genezioDomain,
		ProjectName: projectName,
		Region: region,
		Stage: stage,
	}

	jsonMarshal, err := json.Marshal(data)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest(http.MethodPut, fmt.Sprintf("%s/frontend", constants.API_URL), bytes.NewBuffer(jsonMarshal))
	if err != nil {
		return "", err
	}

	authToken, err := utils.GetAuthToken(ctx)
	if err != nil {
		return "", err
	}


	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept-Version", "genezio-cli/2.2.0")
	req.Header.Set("Authorization", "Bearer "+authToken)

	client := &http.Client{
		Transport: &http.Transport{
			MaxIdleConns: 100,
			MaxIdleConnsPerHost: 100,
		},
	}

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("failed to create project: %s", string(body))
	}

	type response struct {
		Domain string `json:"domain"`
	}

	var dataResponse response
	err = json.Unmarshal(body, &dataResponse)
	if err != nil {
		return "", err
	}

	return dataResponse.Domain, nil

}