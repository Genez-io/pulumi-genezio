package resources

import (
	"fmt"
	"log"
	"strings"

	"github.com/Genez-io/pulumi-genezio/provider/domain"
	"github.com/Genez-io/pulumi-genezio/provider/requests"
	"github.com/Genez-io/pulumi-genezio/provider/utils"
	p "github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi-go-provider/infer"
)

type Authentication struct{}

type AuthenticationArgs struct {
	Project      domain.Project                  `pulumi:"project"`
	DatabaseType string                          `pulumi:"databaseType"`
	DatabaseUrl  string                          `pulumi:"databaseUrl"`
	Provider     *domain.AuthenticationProviders `pulumi:"provider,optional"`
}

type AuthenticationState struct {
	AuthenticationArgs

	Token  string `pulumi:"token"`
	Region string `pulumi:"region"`
}

func (*Authentication) Diff(ctx p.Context, id string, olds AuthenticationState, news AuthenticationArgs) (p.DiffResponse, error) {
	diff := map[string]p.PropertyDiff{}

	areProjectsIdentical := utils.CompareProjects(olds.Project, news.Project)
	if !areProjectsIdentical {
		diff["project"] = p.PropertyDiff{Kind: p.DeleteReplace}
	}

	if olds.DatabaseType != news.DatabaseType {
		diff["databaseType"] = p.PropertyDiff{Kind: p.DeleteReplace}
	}

	if olds.DatabaseUrl != news.DatabaseUrl {
		diff["databaseUrl"] = p.PropertyDiff{Kind: p.DeleteReplace}
	}

	if olds.Provider == nil {
		if news.Provider != nil {
			diff["provider"] = p.PropertyDiff{Kind: p.Update}
		}
	} else {
		if news.Provider == nil {
			diff["provider"] = p.PropertyDiff{Kind: p.Update}
		} else {
			areProvidersIdentical := utils.CompareAuthProviders(*olds.Provider, *news.Provider)
			if !areProvidersIdentical {
				diff["provider"] = p.PropertyDiff{Kind: p.Update}
			}
		}
	}

	return p.DiffResponse{
		DeleteBeforeReplace: true,
		HasChanges:          len(diff) > 0,
		DetailedDiff:        diff,
	}, nil
}

func (*Authentication) Read(ctx p.Context, id string, inputs AuthenticationArgs, state AuthenticationState) (string, AuthenticationArgs, AuthenticationState, error) {

	stage := "prod"
	contextStage := infer.GetConfig[*domain.Config](ctx).Stage
	if contextStage != nil {
		stage = *contextStage
	}

	projectDetails, err := requests.GetProjectDetails(ctx, state.Project.Name)
	if err != nil {
		log.Println("Error getting project details: ", err)
		return id, inputs, state, err
	}

	var currentProjectEnv *domain.ProjectEnvDetails
	for _, projectEnv := range projectDetails.Project.ProjectEnvs {
		if projectEnv.Name == stage {
			currentProjectEnv = &projectEnv
			break
		}
	}

	if currentProjectEnv == nil {
		return id, inputs, state, fmt.Errorf("project environment not found")
	}

	getAuthenticationResponse, err := requests.GetAuthentication(ctx, currentProjectEnv.Id)
	if err != nil {
		log.Println("Error getting authentication", err)
		return id, inputs, state, err
	}

	state.DatabaseType = getAuthenticationResponse.DatabaseType
	state.DatabaseUrl = getAuthenticationResponse.DatabaseUrl
	state.Token = getAuthenticationResponse.Token
	state.Region = getAuthenticationResponse.Region
	state.Project.Name = projectDetails.Project.Name
	state.Project.Region = projectDetails.Project.Region

	authProvidersResponse, err := requests.GetAuthProviders(ctx, currentProjectEnv.Id)
	if err != nil {
		log.Println("Error getting auth providers", err)
		return id, inputs, state, err
	}

	for _, provider := range authProvidersResponse.AuthProviders {
		switch provider.Name {
		case "email":
			if provider.Enabled {

				*state.Provider.Email = true
			} else {
				state.Provider.Email = nil
			}
		case "web3":
			if provider.Enabled {
				*state.Provider.Web3 = true
			} else {
				state.Provider.Web3 = nil
			}
		case "google":
			if provider.Enabled && provider.Config != nil {
				state.Provider.Google = &domain.GoogleProvider{
					ID:     provider.Config["GNZ_AUTH_GOOGLE_ID"],
					Secret: provider.Config["GNZ_AUTH_GOOGLE_SECRET"],
				}
			} else {
				state.Provider.Google = nil
			}
		}
	}

	return id, inputs, state, nil
}

func (*Authentication) Update(ctx p.Context, id string, olds AuthenticationState, news AuthenticationArgs, preview bool) (AuthenticationState, error) {

	state := AuthenticationState{AuthenticationArgs: news, Token: olds.Token, Region: olds.Region}
	if preview {
		return state, nil
	}

	if news.Provider != nil {
		stage := "prod"
		contextStage := infer.GetConfig[*domain.Config](ctx).Stage
		if contextStage != nil {
			stage = *contextStage
		}

		projectDetails, err := requests.GetProjectDetails(ctx, news.Project.Name)
		if err != nil {
			log.Println("Error getting project details: ", err)
			return state, err
		}

		var currentProjectEnv *domain.ProjectEnvDetails
		for _, projectEnv := range projectDetails.Project.ProjectEnvs {
			if projectEnv.Name == stage {
				currentProjectEnv = &projectEnv
				break
			}
		}

		if currentProjectEnv == nil {
			return state, fmt.Errorf("project environment not found")
		}

		authProvidersResponse, err := requests.GetAuthProviders(ctx, currentProjectEnv.Id)
		if err != nil {
			log.Println("Error getting auth providers", err)
			return state, err
		}

		var providersDetails []domain.AuthProviderDetails

		for _, provider := range authProvidersResponse.AuthProviders {
			switch provider.Name {
			case "email":
				if news.Provider.Email != nil {
					providersDetails = append(providersDetails, domain.AuthProviderDetails{
						Id:      provider.Id,
						Name:    provider.Name,
						Enabled: *news.Provider.Email,
					})
				} else {
					providersDetails = append(providersDetails, domain.AuthProviderDetails{
						Id:      provider.Id,
						Name:    provider.Name,
						Enabled: false,
					})
				}
			case "web3":
				if news.Provider.Web3 != nil {
					providersDetails = append(providersDetails, domain.AuthProviderDetails{
						Id:      provider.Id,
						Name:    provider.Name,
						Enabled: *news.Provider.Web3,
					})
				} else {
					providersDetails = append(providersDetails, domain.AuthProviderDetails{
						Id:      provider.Id,
						Name:    provider.Name,
						Enabled: false,
					})
				}
			case "google":
				if news.Provider.Google != nil {
					providersDetails = append(providersDetails, domain.AuthProviderDetails{
						Id:      provider.Id,
						Name:    provider.Name,
						Enabled: true,
						Config: map[string]string{
							"GNZ_AUTH_GOOGLE_ID":     news.Provider.Google.ID,
							"GNZ_AUTH_GOOGLE_SECRET": news.Provider.Google.Secret,
						},
					})
				} else {
					providersDetails = append(providersDetails, domain.AuthProviderDetails{
						Id:      provider.Id,
						Name:    provider.Name,
						Enabled: false,
					})
				}

			}
		}

		if len(providersDetails) > 0 {
			_, err = requests.SetAuthProviders(ctx, currentProjectEnv.Id, domain.SetAuthProvidersRequest{
				AuthProviders: providersDetails,
			})
			if err != nil {
				log.Println("Error setting auth providers", err)
				return state, err
			}
		}

	}

	state.Token = olds.Token
	state.Region = olds.Region

	return state, nil
}

func (*Authentication) Create(ctx p.Context, name string, input AuthenticationArgs, preview bool) (string, AuthenticationState, error) {

	state := AuthenticationState{AuthenticationArgs: input}
	if preview {
		return name, state, nil
	}

	stage := "prod"
	contextStage := infer.GetConfig[*domain.Config](ctx).Stage
	if contextStage != nil {
		stage = *contextStage
	}

	projectDetails, err := requests.GetProjectDetails(ctx, input.Project.Name)
	if err != nil {
		log.Println("Error getting project details: ", err)
		return "", state, err
	}

	var currentProjectEnv *domain.ProjectEnvDetails
	for _, projectEnv := range projectDetails.Project.ProjectEnvs {
		if projectEnv.Name == stage {
			currentProjectEnv = &projectEnv
			break
		}
	}

	if currentProjectEnv == nil {
		return name, state, fmt.Errorf("project environment not found")
	}

	createAuthenticationResponse, err := requests.SetAuthentication(ctx, currentProjectEnv.Id, domain.SetAuthenticationRequest{
		Enabled:      true,
		DatabaseType: input.DatabaseType,
		DatabaseUrl:  input.DatabaseUrl,
	})
	if err != nil {
		log.Println("Error creating authentication", err)
		return name, state, err
	}

	state.Token = createAuthenticationResponse.Token
	state.Region = createAuthenticationResponse.Region

	authProvidersResponse, err := requests.GetAuthProviders(ctx, currentProjectEnv.Id)
	if err != nil {
		log.Println("Error getting auth providers", err)
		return name, state, err
	}

	var providersDetails []domain.AuthProviderDetails

	if input.Provider != nil {
		for _, provider := range authProvidersResponse.AuthProviders {
			switch provider.Name {
			case "email":
				if input.Provider.Email != nil && *input.Provider.Email {
					providersDetails = append(providersDetails, domain.AuthProviderDetails{
						Id:      provider.Id,
						Name:    provider.Name,
						Enabled: true,
					})
				}
			case "web3":
				if input.Provider.Web3 != nil && *input.Provider.Web3 {
					providersDetails = append(providersDetails, domain.AuthProviderDetails{
						Id:      provider.Id,
						Name:    provider.Name,
						Enabled: true,
					})
				}
			case "google":
				if input.Provider.Google != nil {
					providersDetails = append(providersDetails, domain.AuthProviderDetails{
						Id:      provider.Id,
						Name:    provider.Name,
						Enabled: true,
						Config: map[string]string{
							"GNZ_AUTH_GOOGLE_ID":     input.Provider.Google.ID,
							"GNZ_AUTH_GOOGLE_SECRET": input.Provider.Google.Secret,
						},
					})
				}

			}
		}
	}

	if len(providersDetails) > 0 {
		_, err = requests.SetAuthProviders(ctx, currentProjectEnv.Id, domain.SetAuthProvidersRequest{
			AuthProviders: providersDetails,
		})
		if err != nil {
			log.Println("Error setting auth providers", err)
			return name, state, err
		}
	}

	return name, state, nil

}

func (*Authentication) Delete(ctx p.Context, id string, state AuthenticationState) error {
	stage := "prod"
	contextStage := infer.GetConfig[*domain.Config](ctx).Stage
	if contextStage != nil {
		stage = *contextStage
	}

	projectDetails, err := requests.GetProjectDetails(ctx, state.Project.Name)
	if err != nil {
		log.Println("Error getting project details: ", err)
		return err
	}

	var currentProjectEnv *domain.ProjectEnvDetails
	for _, projectEnv := range projectDetails.Project.ProjectEnvs {
		if projectEnv.Name == stage {
			currentProjectEnv = &projectEnv
			break
		}
	}

	if currentProjectEnv == nil {
		return fmt.Errorf("project environment not found")
	}

	_, err = requests.SetAuthentication(ctx, currentProjectEnv.Id, domain.SetAuthenticationRequest{
		Enabled: false,
	})
	if err != nil {
		if strings.Contains(err.Error(), "project integration not found") {
			log.Println("Authentication is already deleted")
			return nil
		}
		log.Println("Error deleting authentication", err)
		return err
	}

	return nil
}
