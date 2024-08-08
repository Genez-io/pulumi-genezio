package domain

type EnvironmentVariable struct {
	Name  string `pulumi:"name" json:"name"`
	Value string `pulumi:"value" json:"value" provider:"secret"`
}

type SetEnvironmentVariablesRequest struct {
	EnvironmentVariables []EnvironmentVariable `json:"environmentVariables"`
}
