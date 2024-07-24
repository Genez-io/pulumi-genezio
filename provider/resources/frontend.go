package resources

type Frontend struct{}

type FrontendArgs struct {
	ProjectName string `pulumi:"projectName"`
	Region 	string `pulumi:"region"`
	AuthToken string `pulumi:"authToken"`
	Path string `pulumi:"path"`
	Subdomain *string `pulumi:"subdomain,optional"`
	Publish string `pulumi:"publish"`
}

