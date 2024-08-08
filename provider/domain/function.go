package domain

type FunctionDetails struct {
	Id          string           `json:"id"`
	Name        string           `json:"name"`
	ProjectName string           `json:"projectName"`
	Status      DeploymentStatus `json:"status"`
	CloudURL    string           `json:"cloudURL"`
	CreatedAt   int64            `json:"createdAt"`
	UpdatedAt   int64            `json:"updatedAt"`
}

type FunctionConfiguration struct {
	Name     string `pulumi:"name"`
	Path     string `pulumi:"path"`
	Language string `pulumi:"language"`
	Handler  string `pulumi:"handler"`
	Entry    string `pulumi:"entry"`
	Type     string `pulumi:"type"`
}

type MappedFunction struct {
	Name      string `json:"name"`
	Language  string `json:"language"`
	EntryFile string `json:"entryFile"`
}
