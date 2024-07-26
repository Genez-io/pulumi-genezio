package requests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/Genez-io/pulumi-genezio/provider/constants"
	p "github.com/pulumi/pulumi-go-provider"
)

func GetPresignedUrl(
	ctx p.Context,
	region string,
	archiveName string,
	projectName string,
	deployUnitName string,
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

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/core/deployment-url",constants.API_URL), bytes.NewBuffer(jsonMarshal))
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+ctx.Value("authToken").(string))
	req.Header.Set("Accept-Version", "genezio-cli/2.2.0")

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

	if resp.StatusCode != 200 {
		return "", fmt.Errorf("error: %s", string(body))
	}

	var data responseData
	err = json.Unmarshal(body, &data)
	if err != nil {
		return "", err
	}

	return data.PresignedURL, nil

}