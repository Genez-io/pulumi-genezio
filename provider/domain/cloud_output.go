package domain

type GenezioCloudOutput struct {
	ProjectID string `pulumi:"projectID"`
	ProjectEnvID string `pulumi:"projectEnvID"`
	Classes []string `pulumi:"classes"`
	Functions []DeployCodeFunctionResponse `pulumi:"functions"`
}

type DeployCodeFunctionResponse struct {
	CloudUrl string `pulumi:"cloudUrl"`
	ID string `pulumi:"functionID"`
	Name string `pulumi:"name"`
}

type DeployCodeResponse struct {
	Status string `pulumi:"status"`
	ProjectID string `pulumi:"projectID"`
	ProjectEnvID string `pulumi:"projectEnvID"`
	Classes []string `pulumi:"classes"`
	Functions []DeployCodeFunctionResponse `pulumi:"functions"`
}

type CreateProjectResponse struct {
	ProjectID string `pulumi:"projectID"`
	ProjectEnvID string `pulumi:"projectEnvID"`
}