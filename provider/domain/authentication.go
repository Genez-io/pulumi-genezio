package domain

type GoogleProvider struct {
	ID     string `pulumi:"id"`
	Secret string `pulumi:"secret"`
}

type AuthenticationProviders struct {
	Email  *bool           `pulumi:"email,optional"`
	Web3   *bool           `pulumi:"web3,optional"`
	Google *GoogleProvider `pulumi:"google,optional"`
}

type SetAuthenticationRequest struct {
	Enabled      bool   `json:"enabled"`
	DatabaseType string `json:"databaseType"`
	DatabaseUrl  string `json:"databaseUrl"`
}

type SetAuthenticationResponse struct {
	Enabled      bool   `json:"enabled"`
	DatabaseType string `json:"databaseType"`
	DatabaseUrl  string `json:"databaseUrl"`
	Region       string `json:"region"`
	Token        string `json:"token"`
}

type AuthProviderDetails struct {
	Id      string            `json:"id"`
	Name    string            `json:"name"`
	Enabled bool              `json:"enabled"`
	Config  map[string]string `json:"config"`
}

type GetAuthProvidersResponse struct {
	Status        string                `json:"status"`
	AuthProviders []AuthProviderDetails `json:"authProviders"`
}

type SetAuthProvidersRequest struct {
	AuthProviders []AuthProviderDetails `json:"authProviders"`
}

type SetAuthProvidersResponse struct {
	Status        string                `json:"status"`
	AuthProviders []AuthProviderDetails `json:"authProviders"`
}

type GetAuthenticationResponse struct {
	Enabled      bool   `json:"enabled"`
	DatabaseUrl  string `json:"databaseUrl"`
	DatabaseType string `json:"databaseType"`
	Token        string `json:"token"`
	Region       string `json:"region"`
}
