package domain

type Workspace struct {
	Backend string `pulumi:"backend"`
}

type AstSummary struct {
	Version string   `pulumi:"version"`
	Classes []string `pulumi:"classes"`
}

type ClassDetails struct {
	ID          string                 `json:"id"`
	Name        string                 `json:"name"`
	ProjectName string                 `json:"projectName"`
	Status      DeploymentStatus       `json:"status"`
	Ast         map[string]interface{} `json:"ast"`
	CloudUrl    string                 `json:"cloudUrl"`
	CreatedAt   int64                  `json:"createdAt"`
	UpdatedAt   int64                  `json:"updatedAt"`
}
