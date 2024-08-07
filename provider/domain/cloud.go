package domain

type GenezioCloudInput struct {
	Type               string `pulumi:"type"`
	Name               string `pulumi:"name"`
	ArchivePath        string `pulumi:"archivePath"`
	EntryFile          string `pulumi:"entryFile"`
	UnzippedBundleSize int64  `pulumi:"unzippedBundleSize"`
}

type GenezioCloudOutput struct {
	ProjectID    string                       `pulumi:"projectID"`
	ProjectEnvID string                       `pulumi:"projectEnvID"`
	Classes      []ClassDetails               `pulumi:"classes"`
	Functions    []DeployCodeFunctionResponse `pulumi:"functions"`
}
