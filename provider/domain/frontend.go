package domain

type FrontendPresignedUrlResponse struct {
	UserID string `json:"userId"`
	PresignedURL string `json:"presignedURL"`
	Domain string `json:"domain"`
}

type FrontendConfiguration struct {
	Path string `json:"path"`
	Subdomain string `json:"subdomain"`
	Publish string `json:"publish"`
}