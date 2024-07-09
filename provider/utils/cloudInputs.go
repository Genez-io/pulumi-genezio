package utils

type GenezioCloudInput struct {
	Type string `pulumi:"type"`
	Name string `pulumi:"name"`
	ArchivePath string `pulumi:"archivePath"`
	EntryFile string `pulumi:"entryFile"`
	UnzippedBundleSize int64 `pulumi:"unzippedBundleSize"`
}