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

type ProjectEnvDetails struct {
	Id        string            `json:"id"`
	Name      string            `json:"name"`
	Classes   []ClassDetails    `json:"classes"` // TODO - This is incomplete
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
type Options struct {
	NodeRuntime  string `pulumi:"nodeRuntime"`
	Architecture string `pulumi:"architecture"`
}

type ProjectConfiguration struct {
	Name          string                  `pulumi:"name"`
	Region        string                  `pulumi:"region"`
	Options       Options                 `pulumi:"options"`
	CloudProvider string                  `pulumi:"cloudProvider"`
	Workspace     Workspace               `pulumi:"workspace"`
	AstSummary    AstSummary              `pulumi:"astSummary"`
	Classes       []ClassDetails          `pulumi:"classes"`
	Functions     []FunctionConfiguration `pulumi:"functions"`
}

type DeployCodeFunctionResponse struct {
	CloudUrl string `pulumi:"cloudUrl"`
	ID       string `pulumi:"functionID"`
	Name     string `pulumi:"name"`
}

type DeployCodeResponse struct {
	Status       string                       `pulumi:"status"`
	ProjectID    string                       `pulumi:"projectID"`
	ProjectEnvID string                       `pulumi:"projectEnvID"`
	Classes      []ClassDetails               `pulumi:"classes"`
	Functions    []DeployCodeFunctionResponse `pulumi:"functions"`
}

type CreateProjectResponse struct {
	ProjectID    string `pulumi:"projectID"`
	ProjectEnvID string `pulumi:"projectEnvID"`
}

type CreateProjectRequest struct {
	ProjectName   string `json:"projectName"`
	CloudProvider string `json:"cloudProvider"`
	Region        string `json:"region"`
	Stage         string `json:"stage"`
}

type DeployRequest struct {
	Options       Options          `json:"options"`
	Classes       []ClassDetails   `json:"classes"`
	Functions     []MappedFunction `json:"functions"`
	ProjectName   string           `json:"projectName"`
	Region        string           `json:"region"`
	CloudProvider string           `json:"cloudProvider"`
	Stage         string           `json:"stage"`
	Stack         []string         `json:"stack"`
}

type GetPresignedUrlRequest struct {
	Region      string `json:"region"`
	Filename    string `json:"filename"`
	ProjectName string `json:"projectName"`
	ClassName   string `json:"className"`
}

type GetPresignedUrlResponse struct {
	PresignedUrl string `json:"presignedUrl"`
}

type Project struct {
	Name   string `pulumi:"name" json:"name"`
	Region string `pulumi:"region" json:"region"`
}
