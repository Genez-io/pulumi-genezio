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

func GetFrontendPresignedUrl(ctx p.Context, subdomain string, projectName string, stage string) (domain.FrontendPresignedUrlResponse, error) {

	region := "us-east-1"

	if projectName == "" {
		return domain.FrontendPresignedUrlResponse{}, fmt.Errorf("projectName is required")
	}

	type request struct{
		SubdomainName string `json:"subdomainName"`
		ProjectName string `json:"projectName"`
		Region string `json:"region"`
		Stage string `json:"stage"`
	}

	data := request{
		SubdomainName: subdomain,
		ProjectName: projectName,
		Region: region,
		Stage: stage,
	} 

	jsonMarshal, err := json.Marshal(data)
	if err != nil {
		return domain.FrontendPresignedUrlResponse{}, err
	}

	req, err := http.NewRequest(http.MethodGet,fmt.Sprintf("%s/core/frontend-deployment-url",constants.API_URL),bytes.NewBuffer(jsonMarshal))
	if err != nil {
		return domain.FrontendPresignedUrlResponse{}, err
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
		return domain.FrontendPresignedUrlResponse{}, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return domain.FrontendPresignedUrlResponse{}, err
	}

	if resp.StatusCode != http.StatusOK {
		return domain.FrontendPresignedUrlResponse{}, fmt.Errorf("failed to get presigned url: %s", string(body))
	}

	var dataResponse domain.FrontendPresignedUrlResponse
	err = json.Unmarshal(body, &dataResponse)
	if err != nil {
		return domain.FrontendPresignedUrlResponse{}, err
	}


	return dataResponse, nil

}