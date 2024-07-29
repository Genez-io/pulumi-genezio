package domain

type DeploymentStatus string

type CollaborationRole string

type Provider string

const (
	GenezioCloud     Provider = "genezio-cloud"
	GenezioUnikernel Provider = "genezio-unikernel"
	GenezioAws       Provider = "genezio-aws"
	GenezioCluster   Provider = "genezio-cluster"
)

type FunctionDetails struct {
	Id          string           `json:"id"`
	Name        string           `json:"name"`
	ProjectName string           `json:"projectName"`
	Status      DeploymentStatus `json:"status"`
	CloudURL    string           `json:"cloudURL"`
	CreatedAt   int64            `json:"createdAt"`
	UpdatedAt   int64            `json:"updatedAt"`
}

type ProjectEnvDetails struct {
	Id        string            `json:"id"`
	Name      string            `json:"name"`
	Classes   []string          `json:"classes"` // TODO - This is incomplete
	Functions []FunctionDetails `json:"functions"`
}

type ProjectDetails struct {
	Id          string              `json:"id"`
	Name        string              `json:"name"`
	Region      string              `json:"region"`
	CreatedAt   int64               `json:"createdAt"`
	UpdatedAt   int64               `json:"updatedAt"`
	ProjectEnvs []ProjectEnvDetails `json:"projectEnvs"`
	Stack       []string            `json:"stack"`
}

type ProjectDetailsResponse struct {
	Status  string         `json:"status"`
	Project ProjectDetails `json:"project"`
}

type DeleteProjectResponse struct {
	Status string `json:"status"`
}
