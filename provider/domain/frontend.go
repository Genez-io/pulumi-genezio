package domain

type FrontendConfiguration struct {
	Path      string `json:"path"`
	Subdomain string `json:"subdomain"`
	Publish   string `json:"publish"`
}

type CreateFrontendProjectRequest struct {
	GenezioDomain string `json:"genezioDomain"`
	ProjectName   string `json:"projectName"`
	Region        string `json:"region"`
	Stage         string `json:"stage"`
}

type CreateFrontendProjectResponse struct {
	Domain string `json:"domain"`
}

type GetFrontendPresignedUrlRequest struct {
	SubdomainName string `json:"subdomainName"`
	ProjectName   string `json:"projectName"`
	Region        string `json:"region"`
	Stage         string `json:"stage"`
}

type FrontendPresignedUrlResponse struct {
	UserID       string `json:"userId"`
	PresignedURL string `json:"presignedURL"`
	Domain       string `json:"domain"`
}

type FrontendDetail struct {
	GenezioDomain      string `json:"genezioDomain"`
	FullDomain         string `json:"fullDomain"`
	CustomDomain       string `json:"customDomain"`
	CreatedAt          int64  `json:"createdAt"`
	UpdatedAt          int64  `json:"updatedAt"`
	CNameName          string `json:"cnameName"`
	CNameValue         string `json:"cnameValue"`
	DistributionDomain string `json:"distributionDomain"`
	Status             string `json:"status"`
	FailureReason      string `json:"failReason"`
}

type GetFrontendByEnvIdResponse struct {
	Status     string           `json:"status"`
	TotalCount int              `json:"totalCount"`
	List       []FrontendDetail `json:"list"`
}
