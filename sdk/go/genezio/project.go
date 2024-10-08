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
type Project struct {
	pulumi.CustomResourceState

	// The cloud provider on which the project will be deployed.
	//
	//     Supported cloud providers are:
	//     - genezio-cloud
	CloudProvider pulumi.StringPtrOutput `pulumi:"cloudProvider"`
	// The backend environment variables that will be securely stored for the project.
	Environment domain.EnvironmentVariableArrayOutput `pulumi:"environment"`
	// The name of the project to be deployed.
	Name pulumi.StringOutput `pulumi:"name"`
	// The environment ID.
	ProjectEnvId pulumi.StringOutput `pulumi:"projectEnvId"`
	// The project ID.
	ProjectId pulumi.StringOutput `pulumi:"projectId"`
	// The region in which the project will be deployed.
	//
	//     Supported regions are:
	//     - us-east-1
	//     - eu-central-1
	Region pulumi.StringOutput `pulumi:"region"`
}

// NewProject registers a new resource with the given unique name, arguments, and options.
func NewProject(ctx *pulumi.Context,
	name string, args *ProjectArgs, opts ...pulumi.ResourceOption) (*Project, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Name == nil {
		return nil, errors.New("invalid value for required argument 'Name'")
	}
	if args.Region == nil {
		return nil, errors.New("invalid value for required argument 'Region'")
	}
	if args.CloudProvider == nil {
		args.CloudProvider = pulumi.StringPtr("genezio-cloud")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Project
	err := ctx.RegisterResource("genezio:index:Project", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetProject gets an existing Project resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetProject(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ProjectState, opts ...pulumi.ResourceOption) (*Project, error) {
	var resource Project
	err := ctx.ReadResource("genezio:index:Project", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Project resources.
type projectState struct {
}

type ProjectState struct {
}

func (ProjectState) ElementType() reflect.Type {
	return reflect.TypeOf((*projectState)(nil)).Elem()
}

type projectArgs struct {
	// The cloud provider on which the project will be deployed.
	//
	//     Supported cloud providers are:
	//     - genezio-cloud
	CloudProvider *string `pulumi:"cloudProvider"`
	// The backend environment variables that will be securely stored for the project.
	Environment []domain.EnvironmentVariable `pulumi:"environment"`
	// The name of the project to be deployed.
	Name string `pulumi:"name"`
	// The region in which the project will be deployed.
	//
	//     Supported regions are:
	//     - us-east-1
	//     - eu-central-1
	Region string `pulumi:"region"`
}

// The set of arguments for constructing a Project resource.
type ProjectArgs struct {
	// The cloud provider on which the project will be deployed.
	//
	//     Supported cloud providers are:
	//     - genezio-cloud
	CloudProvider pulumi.StringPtrInput
	// The backend environment variables that will be securely stored for the project.
	Environment domain.EnvironmentVariableArrayInput
	// The name of the project to be deployed.
	Name pulumi.StringInput
	// The region in which the project will be deployed.
	//
	//     Supported regions are:
	//     - us-east-1
	//     - eu-central-1
	Region pulumi.StringInput
}

func (ProjectArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*projectArgs)(nil)).Elem()
}

type ProjectInput interface {
	pulumi.Input

	ToProjectOutput() ProjectOutput
	ToProjectOutputWithContext(ctx context.Context) ProjectOutput
}

func (*Project) ElementType() reflect.Type {
	return reflect.TypeOf((**Project)(nil)).Elem()
}

func (i *Project) ToProjectOutput() ProjectOutput {
	return i.ToProjectOutputWithContext(context.Background())
}

func (i *Project) ToProjectOutputWithContext(ctx context.Context) ProjectOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProjectOutput)
}

type ProjectOutput struct{ *pulumi.OutputState }

func (ProjectOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Project)(nil)).Elem()
}

func (o ProjectOutput) ToProjectOutput() ProjectOutput {
	return o
}

func (o ProjectOutput) ToProjectOutputWithContext(ctx context.Context) ProjectOutput {
	return o
}

// The cloud provider on which the project will be deployed.
//
//	Supported cloud providers are:
//	- genezio-cloud
func (o ProjectOutput) CloudProvider() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Project) pulumi.StringPtrOutput { return v.CloudProvider }).(pulumi.StringPtrOutput)
}

// The backend environment variables that will be securely stored for the project.
func (o ProjectOutput) Environment() domain.EnvironmentVariableArrayOutput {
	return o.ApplyT(func(v *Project) domain.EnvironmentVariableArrayOutput { return v.Environment }).(domain.EnvironmentVariableArrayOutput)
}

// The name of the project to be deployed.
func (o ProjectOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Project) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The environment ID.
func (o ProjectOutput) ProjectEnvId() pulumi.StringOutput {
	return o.ApplyT(func(v *Project) pulumi.StringOutput { return v.ProjectEnvId }).(pulumi.StringOutput)
}

// The project ID.
func (o ProjectOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v *Project) pulumi.StringOutput { return v.ProjectId }).(pulumi.StringOutput)
}

// The region in which the project will be deployed.
//
//	Supported regions are:
//	- us-east-1
//	- eu-central-1
func (o ProjectOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *Project) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ProjectInput)(nil)).Elem(), &Project{})
	pulumi.RegisterOutputType(ProjectOutput{})
}
