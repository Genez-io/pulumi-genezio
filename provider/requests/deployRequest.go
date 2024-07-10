package requests

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"

	"github.com/Genez-io/pulumi-genezio/provider/domain"
)


func DeployRequest(
	projectConfiguration domain.ProjectConfiguration,
	genezioDeployInput []domain.GenezioCloudInput,
	stage string,
	stack []string,
	authToken string,
) (domain.DeployCodeResponse, error) {
	

	type functionsMap struct {
		Name string `json:"name"`
		Language string `json:"language"`
		EntryFile string `json:"entryFile"`
	}

	type request struct {
		Options domain.Options `json:"options"`
		Classes []string `json:"classes"`
		Functions []functionsMap `json:"functions"`
		ProjectName string `json:"projectName"`
		Region string `json:"region"`
		CloudProvider string `json:"cloudProvider"`
		Stage string `json:"stage"`
		Stack []string `json:"stack"`
	}

	functionsMappping := []functionsMap{}

	for _, fun := range projectConfiguration.Functions {
		entryFile := ""
		for _, input := range genezioDeployInput {
			if input.Name == fun.Name {
				entryFile = input.EntryFile
				break
			}
		}

		functionsMappping = append(functionsMappping, functionsMap{
			Name: fun.Name,
			Language: fun.Language,
			EntryFile: entryFile,
		})
	}
	
	data := request{
		Options: projectConfiguration.Options,
		Classes: projectConfiguration.Classes,
		Functions: functionsMappping,
		ProjectName: projectConfiguration.Name,
		Region: projectConfiguration.Region,
		CloudProvider: projectConfiguration.CloudProvider,
		Stage: stage,
		Stack: stack,
	}

	jsonMarshal,err := json.Marshal(data)
	if err != nil {
		return domain.DeployCodeResponse{}, err
	}

	req, err := http.NewRequest("PUT", "https://dev.api.genez.io/core/deployment", bytes.NewBuffer(jsonMarshal))
	if err != nil {
		return domain.DeployCodeResponse{}, err
	}

	req.Header.Set("Authorization", "Bearer "+authToken)
	req.Header.Set("Accept-Version", "genezio-webapp/0.3.0")

	client := &http.Client{
		Transport: &http.Transport{
			MaxIdleConns: 100,
			MaxIdleConnsPerHost: 100,
		},
	}

	resp, err := client.Do(req)
	if err != nil {
		return domain.DeployCodeResponse{}, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return domain.DeployCodeResponse{}, err
	}

	var dataResponse domain.DeployCodeResponse
	err = json.Unmarshal(body, &dataResponse)
	if err != nil {
		return domain.DeployCodeResponse{}, err
	}

	return dataResponse, nil

}