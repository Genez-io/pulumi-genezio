package utils

type Options struct {
	NodeRuntime string `pulumi:"nodeRuntime"`
	Architecture string `pulumi:"architecture"`
  }

  type Workspace struct {
	Backend string `pulumi:"backend"`
  }

  type FunctionConfiguration struct {
	Name string `pulumi:"name"`
	Path string `pulumi:"path"`
	Language string `pulumi:"language"`
	Handler string `pulumi:"handler"`
	Entry string `pulumi:"entry"`
	Type string `pulumi:"type"`
  }

  type AstSummary struct {
	Version string `pulumi:"version"`
	Classes []string `pulumi:"classes"`
  }

  type ProjectConfiguration struct {
	Name string `pulumi:"name"`
	Region string `pulumi:"region"`
	Options Options `pulumi:"options"`
	CloudProvider string `pulumi:"cloudProvider"`
	Workspace Workspace `pulumi:"workspace"`
	AstSummary AstSummary `pulumi:"astSummary"`
	Classes []string `pulumi:"classes"`
	Functions []FunctionConfiguration `pulumi:"functions"`
  }
