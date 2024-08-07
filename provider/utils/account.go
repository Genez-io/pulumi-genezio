package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/Genez-io/pulumi-genezio/provider/constants"
	"github.com/Genez-io/pulumi-genezio/provider/domain"
	p "github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi-go-provider/infer"
)

func GetAuthToken(ctx p.Context) (string, error) {
	token := infer.GetConfig[*domain.Config](ctx).AuthToken
	if token == "" {
		return "", fmt.Errorf("no authentification token provided")
	}

	return token, nil
}

func IsLoggedIn(ctx p.Context) (string, error) {
	_, authToken, err := GetUser(ctx)
	if err != nil {
		return "", err
	}

	return authToken, nil

}

func GetUser(ctx p.Context) (string, string, error) {

	authToken, err := GetAuthToken(ctx)
	if err != nil {
		return "", "", err
	}

	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/users/user", constants.API_URL), nil)
	if err != nil {
		return "", "", err
	}

	req.Header.Set("Accept-Version", "genezio-cli/2.2.0")
	req.Header.Set("Authorization", "Bearer "+authToken)

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return "", "", err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", "", err
	}

	if resp.StatusCode != http.StatusOK {
		return "", "", fmt.Errorf("error: %s", body)
	}

	var user domain.UserPayload
	err = json.Unmarshal(body, &user)
	if err != nil {
		return "", "", err
	}

	return user.ID, authToken, nil
}
