package requests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/Genez-io/pulumi-genezio/provider/constants"
	"github.com/Genez-io/pulumi-genezio/provider/domain"
)

func CreateProject(cloudProvider string, region string, authToken string, name string, stage string) (domain.CreateProjectResponse, error) {

	if cloudProvider == "" {
		cloudProvider = "genezio-cloud"
	}

	if region == "" {
		region = "eu-central-1"
	}

	if name == "" {
		return domain.CreateProjectResponse{}, fmt.Errorf("name is required")
	}

	if stage == "" {
		stage = "prod"
	}

	type request struct {
		ProjectName string `json:"projectName"`
		CloudProvider string `json:"cloudProvider"`
		Region string `json:"region"`
		Stage string `json:"stage"`
	}

	data := request{
		ProjectName: name,
		CloudProvider: cloudProvider,
		Region: region,
		Stage: stage,
	}

	jsonMarshal, err := json.Marshal(data)
	if err != nil {
		return domain.CreateProjectResponse{}, err
	}

	req, err := http.NewRequest(http.MethodPut, fmt.Sprintf("%s/core/deployment", constants.API_URL), bytes.NewBuffer(jsonMarshal))
	if err != nil {
		return domain.CreateProjectResponse{}, err
	}

	req.Header.Set("Authorization", "Bearer "+authToken)
	req.Header.Set("Accept-Version", "genezio-webapp/0.3.0")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{
		Transport: &http.Transport{
			MaxIdleConns: 100,
			MaxIdleConnsPerHost: 100,
		},
	}

	resp, err := client.Do(req)
	if err != nil {
		return domain.CreateProjectResponse{}, err
	}

	
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return domain.CreateProjectResponse{}, err
	}

	if resp.StatusCode != http.StatusOK {
		return domain.CreateProjectResponse{}, fmt.Errorf("error: %s", string(body))
	}

	var dataResponse domain.DeployCodeResponse
	err = json.Unmarshal(body, &dataResponse)
	if err != nil {
		return domain.CreateProjectResponse{}, err
	}

	projectEnvId:= dataResponse.ProjectEnvID
	projectId:= dataResponse.ProjectID

	return domain.CreateProjectResponse{
		ProjectEnvID: projectEnvId,
		ProjectID: projectId,
	}, nil

}