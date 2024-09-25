// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package genezio

import (
	"context"
	"reflect"

	"errors"
	"example.com/pulumi-genezio/sdk/go/genezio/domain"
	"example.com/pulumi-genezio/sdk/go/genezio/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// A project resource that will be deployed on the Genezio platform.The project resource is used to group resources together and manage them as a single unit.
//
// The project resource will deploy an empty project on the Genezio platform.
//
// It is recommended to create a Project Resource as the first step in your deployment workflow. The output from this resource can then be utilized to provision and configure other resources within the project, ensuring they are properly associated and managed under a unified project.
//
// ## Example Usage
//
// ### Basic Usage
//
// ### Environment Variables
//
// ## Pulumi Output Reference
//
// Once the project is created, the `projectId` and `projectUrl` are available as outputs.
type Frontend struct {
	pulumi.CustomResourceState

	// The commands to run before deploying the frontend.
	BuildCommands pulumi.StringArrayOutput `pulumi:"buildCommands"`
	// The environment variables that will be set for the frontend.
	Environment domain.EnvironmentVariableArrayOutput `pulumi:"environment"`
	// The path to the frontend files.
	Path pulumi.ArchiveOutput `pulumi:"path"`
	// The project to which the frontend will be deployed.
	Project domain.ProjectOutput `pulumi:"project"`
	// The folder in the path that contains the files to be published.
	Publish pulumi.StringOutput `pulumi:"publish"`
	// The subdomain of the frontend.
	Subdomain pulumi.StringPtrOutput `pulumi:"subdomain"`
	// The URL of the frontend.
	Url pulumi.StringOutput `pulumi:"url"`
}

// NewFrontend registers a new resource with the given unique name, arguments, and options.
func NewFrontend(ctx *pulumi.Context,
	name string, args *FrontendArgs, opts ...pulumi.ResourceOption) (*Frontend, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Path == nil {
		return nil, errors.New("invalid value for required argument 'Path'")
	}
	if args.Project == nil {
		return nil, errors.New("invalid value for required argument 'Project'")
	}
	if args.Publish == nil {
		return nil, errors.New("invalid value for required argument 'Publish'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Frontend
	err := ctx.RegisterResource("genezio:index:Frontend", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFrontend gets an existing Frontend resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFrontend(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FrontendState, opts ...pulumi.ResourceOption) (*Frontend, error) {
	var resource Frontend
	err := ctx.ReadResource("genezio:index:Frontend", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Frontend resources.
type frontendState struct {
}

type FrontendState struct {
}

func (FrontendState) ElementType() reflect.Type {
	return reflect.TypeOf((*frontendState)(nil)).Elem()
}

type frontendArgs struct {
	// The commands to run before deploying the frontend.
	BuildCommands []string `pulumi:"buildCommands"`
	// The environment variables that will be set for the frontend.
	Environment []domain.EnvironmentVariable `pulumi:"environment"`
	// The path to the frontend files.
	Path pulumi.Archive `pulumi:"path"`
	// The project to which the frontend will be deployed.
	Project domain.Project `pulumi:"project"`
	// The folder in the path that contains the files to be published.
	Publish string `pulumi:"publish"`
	// The subdomain of the frontend.
	Subdomain *string `pulumi:"subdomain"`
}

// The set of arguments for constructing a Frontend resource.
type FrontendArgs struct {
	// The commands to run before deploying the frontend.
	BuildCommands pulumi.StringArrayInput
	// The environment variables that will be set for the frontend.
	Environment domain.EnvironmentVariableArrayInput
	// The path to the frontend files.
	Path pulumi.ArchiveInput
	// The project to which the frontend will be deployed.
	Project domain.ProjectInput
	// The folder in the path that contains the files to be published.
	Publish pulumi.StringInput
	// The subdomain of the frontend.
	Subdomain pulumi.StringPtrInput
}

func (FrontendArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*frontendArgs)(nil)).Elem()
}

type FrontendInput interface {
	pulumi.Input

	ToFrontendOutput() FrontendOutput
	ToFrontendOutputWithContext(ctx context.Context) FrontendOutput
}

func (*Frontend) ElementType() reflect.Type {
	return reflect.TypeOf((**Frontend)(nil)).Elem()
}

func (i *Frontend) ToFrontendOutput() FrontendOutput {
	return i.ToFrontendOutputWithContext(context.Background())
}

func (i *Frontend) ToFrontendOutputWithContext(ctx context.Context) FrontendOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FrontendOutput)
}

type FrontendOutput struct{ *pulumi.OutputState }

func (FrontendOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Frontend)(nil)).Elem()
}

func (o FrontendOutput) ToFrontendOutput() FrontendOutput {
	return o
}

func (o FrontendOutput) ToFrontendOutputWithContext(ctx context.Context) FrontendOutput {
	return o
}

// The commands to run before deploying the frontend.
func (o FrontendOutput) BuildCommands() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Frontend) pulumi.StringArrayOutput { return v.BuildCommands }).(pulumi.StringArrayOutput)
}

// The environment variables that will be set for the frontend.
func (o FrontendOutput) Environment() domain.EnvironmentVariableArrayOutput {
	return o.ApplyT(func(v *Frontend) domain.EnvironmentVariableArrayOutput { return v.Environment }).(domain.EnvironmentVariableArrayOutput)
}

// The path to the frontend files.
func (o FrontendOutput) Path() pulumi.ArchiveOutput {
	return o.ApplyT(func(v *Frontend) pulumi.ArchiveOutput { return v.Path }).(pulumi.ArchiveOutput)
}

// The project to which the frontend will be deployed.
func (o FrontendOutput) Project() domain.ProjectOutput {
	return o.ApplyT(func(v *Frontend) domain.ProjectOutput { return v.Project }).(domain.ProjectOutput)
}

// The folder in the path that contains the files to be published.
func (o FrontendOutput) Publish() pulumi.StringOutput {
	return o.ApplyT(func(v *Frontend) pulumi.StringOutput { return v.Publish }).(pulumi.StringOutput)
}

// The subdomain of the frontend.
func (o FrontendOutput) Subdomain() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Frontend) pulumi.StringPtrOutput { return v.Subdomain }).(pulumi.StringPtrOutput)
}

// The URL of the frontend.
func (o FrontendOutput) Url() pulumi.StringOutput {
	return o.ApplyT(func(v *Frontend) pulumi.StringOutput { return v.Url }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FrontendInput)(nil)).Elem(), &Frontend{})
	pulumi.RegisterOutputType(FrontendOutput{})
}
