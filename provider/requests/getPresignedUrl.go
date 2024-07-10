package requests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func GetPresignedUrl(
	region string,
	archiveName string,
	projectName string,
	deployUnitName string,
	authToken string,
) (string, error) {

	if region == "" || archiveName=="" || projectName=="" || deployUnitName=="" {
		return "",nil
	}
	type request struct {
		Region string `json:"region"`
		Filename string `json:"filename"`
		ProjectName string `json:"projectName"`
		ClassName string `json:"className"`
	}


	jsonReq := request{
		Region: region,
		Filename: archiveName,
		ProjectName: projectName,
		ClassName: deployUnitName,
	}

	type responseData struct {
		PresignedURL string `json:"presignedUrl"`
	}

	jsonMarshal, err := json.Marshal(jsonReq)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("GET", "https://dev.api.genez.io/core/deployment-url", bytes.NewBuffer(jsonMarshal))
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", "application/json")
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
		return "", err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	fmt.Println(string(body))

	var data responseData
	err = json.Unmarshal(body, &data)
	if err != nil {
		return "", err
	}

	return data.PresignedURL, nil

}