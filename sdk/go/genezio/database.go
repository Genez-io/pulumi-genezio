// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package genezio

import (
	"context"
	"reflect"

	"domain"
	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"internal"
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
type Database struct {
	pulumi.CustomResourceState

	// The database ID.
	DatabaseId pulumi.StringOutput `pulumi:"databaseId"`
	// The name of the database to be deployed.
	Name pulumi.StringOutput `pulumi:"name"`
	// A database can be used in a project by linking it.
	// 	Linking the database will expose a connection URL as an environment variable for convenience.
	// 	The same database can be linked to multiple projects.
	Project domain.ProjectPtrOutput `pulumi:"project"`
	// The region in which the database will be deployed.
	Region pulumi.StringPtrOutput `pulumi:"region"`
	// The type of the database to be deployed.
	//
	//     Supported types are:
	//     - postgres-neon
	Type pulumi.StringPtrOutput `pulumi:"type"`
	// The URL of the database.
	Url pulumi.StringOutput `pulumi:"url"`
}

// NewDatabase registers a new resource with the given unique name, arguments, and options.
func NewDatabase(ctx *pulumi.Context,
	name string, args *DatabaseArgs, opts ...pulumi.ResourceOption) (*Database, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Name == nil {
		return nil, errors.New("invalid value for required argument 'Name'")
	}
	if args.Region == nil {
		args.Region = pulumi.StringPtr("us-east-1")
	}
	if args.Type == nil {
		args.Type = pulumi.StringPtr("postgres-neon")
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"url",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Database
	err := ctx.RegisterResource("genezio:index:Database", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDatabase gets an existing Database resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDatabase(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DatabaseState, opts ...pulumi.ResourceOption) (*Database, error) {
	var resource Database
	err := ctx.ReadResource("genezio:index:Database", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Database resources.
type databaseState struct {
}

type DatabaseState struct {
}

func (DatabaseState) ElementType() reflect.Type {
	return reflect.TypeOf((*databaseState)(nil)).Elem()
}

type databaseArgs struct {
	// The name of the database to be deployed.
	Name string `pulumi:"name"`
	// A database can be used in a project by linking it.
	// 	Linking the database will expose a connection URL as an environment variable for convenience.
	// 	The same database can be linked to multiple projects.
	Project *domain.Project `pulumi:"project"`
	// The region in which the database will be deployed.
	Region *string `pulumi:"region"`
	// The type of the database to be deployed.
	//
	//     Supported types are:
	//     - postgres-neon
	Type *string `pulumi:"type"`
}

// The set of arguments for constructing a Database resource.
type DatabaseArgs struct {
	// The name of the database to be deployed.
	Name pulumi.StringInput
	// A database can be used in a project by linking it.
	// 	Linking the database will expose a connection URL as an environment variable for convenience.
	// 	The same database can be linked to multiple projects.
	Project domain.ProjectPtrInput
	// The region in which the database will be deployed.
	Region pulumi.StringPtrInput
	// The type of the database to be deployed.
	//
	//     Supported types are:
	//     - postgres-neon
	Type pulumi.StringPtrInput
}

func (DatabaseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*databaseArgs)(nil)).Elem()
}

type DatabaseInput interface {
	pulumi.Input

	ToDatabaseOutput() DatabaseOutput
	ToDatabaseOutputWithContext(ctx context.Context) DatabaseOutput
}

func (*Database) ElementType() reflect.Type {
	return reflect.TypeOf((**Database)(nil)).Elem()
}

func (i *Database) ToDatabaseOutput() DatabaseOutput {
	return i.ToDatabaseOutputWithContext(context.Background())
}

func (i *Database) ToDatabaseOutputWithContext(ctx context.Context) DatabaseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatabaseOutput)
}

type DatabaseOutput struct{ *pulumi.OutputState }

func (DatabaseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Database)(nil)).Elem()
}

func (o DatabaseOutput) ToDatabaseOutput() DatabaseOutput {
	return o
}

func (o DatabaseOutput) ToDatabaseOutputWithContext(ctx context.Context) DatabaseOutput {
	return o
}

// The database ID.
func (o DatabaseOutput) DatabaseId() pulumi.StringOutput {
	return o.ApplyT(func(v *Database) pulumi.StringOutput { return v.DatabaseId }).(pulumi.StringOutput)
}

// The name of the database to be deployed.
func (o DatabaseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Database) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// A database can be used in a project by linking it.
//
//	Linking the database will expose a connection URL as an environment variable for convenience.
//	The same database can be linked to multiple projects.
func (o DatabaseOutput) Project() domain.ProjectPtrOutput {
	return o.ApplyT(func(v *Database) domain.ProjectPtrOutput { return v.Project }).(domain.ProjectPtrOutput)
}

// The region in which the database will be deployed.
func (o DatabaseOutput) Region() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Database) pulumi.StringPtrOutput { return v.Region }).(pulumi.StringPtrOutput)
}

// The type of the database to be deployed.
//
//	Supported types are:
//	- postgres-neon
func (o DatabaseOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Database) pulumi.StringPtrOutput { return v.Type }).(pulumi.StringPtrOutput)
}

// The URL of the database.
func (o DatabaseOutput) Url() pulumi.StringOutput {
	return o.ApplyT(func(v *Database) pulumi.StringOutput { return v.Url }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DatabaseInput)(nil)).Elem(), &Database{})
	pulumi.RegisterOutputType(DatabaseOutput{})
}
