package domain

type FrontendConfiguration struct {
	Path string `json:"path"`
	Subdomain string `json:"subdomain"`
	Publish string `json:"publish"`
}

type CreateFrontendProjectRequest struct {
	GenezioDomain string `json:"genezioDomain"`
	ProjectName string `json:"projectName"`
	Region string `json:"region"`
	Stage string `json:"stage"`
}

type CreateFrontendProjectResponse struct {
	Domain string `json:"domain"`
}

type GetFrontendPresignedUrlRequest struct {
	SubdomainName string `json:"subdomainName"`
	ProjectName string `json:"projectName"`
	Region string `json:"region"`
	Stage string `json:"stage"`
}

type FrontendPresignedUrlResponse struct {
	UserID string `json:"userId"`
	PresignedURL string `json:"presignedURL"`
	Domain string `json:"domain"`
}