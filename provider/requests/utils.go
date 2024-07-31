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

type ResponseStatus string

const (
	Success ResponseStatus = "ok"
	Failure ResponseStatus = "error"
)

func MakeRequest(ctx p.Context, method string, endpoint string, body interface{}, response interface{}) error {
	data, err := json.Marshal(body)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(method, fmt.Sprintf("%s/%s", constants.API_URL, endpoint), bytes.NewBuffer(data))
	if err != nil {
		return err
	}

	authToken, err := utils.GetAuthToken(ctx)
	if err != nil {
		return err
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
		return err
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("error: %s and response %v", string(bodyBytes), resp)
	}


	if response != nil {
		err = json.Unmarshal(bodyBytes, &response)
		if err != nil {
			return err
		}
	}
	

	return nil
}
