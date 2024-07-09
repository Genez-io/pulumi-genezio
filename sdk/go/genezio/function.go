// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package genezio

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"internal"
)

type Function struct {
	pulumi.CustomResourceState

	Entry       pulumi.StringOutput `pulumi:"entry"`
	FunctionId  pulumi.StringOutput `pulumi:"functionId"`
	Handler     pulumi.StringOutput `pulumi:"handler"`
	Name        pulumi.StringOutput `pulumi:"name"`
	Path        pulumi.StringOutput `pulumi:"path"`
	ProjectName pulumi.StringOutput `pulumi:"projectName"`
	Region      pulumi.StringOutput `pulumi:"region"`
	Url         pulumi.StringOutput `pulumi:"url"`
}

// NewFunction registers a new resource with the given unique name, arguments, and options.
func NewFunction(ctx *pulumi.Context,
	name string, args *FunctionArgs, opts ...pulumi.ResourceOption) (*Function, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Entry == nil {
		return nil, errors.New("invalid value for required argument 'Entry'")
	}
	if args.Handler == nil {
		return nil, errors.New("invalid value for required argument 'Handler'")
	}
	if args.Name == nil {
		return nil, errors.New("invalid value for required argument 'Name'")
	}
	if args.Path == nil {
		return nil, errors.New("invalid value for required argument 'Path'")
	}
	if args.ProjectName == nil {
		return nil, errors.New("invalid value for required argument 'ProjectName'")
	}
	if args.Region == nil {
		return nil, errors.New("invalid value for required argument 'Region'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Function
	err := ctx.RegisterResource("genezio:index:Function", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFunction gets an existing Function resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFunction(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FunctionState, opts ...pulumi.ResourceOption) (*Function, error) {
	var resource Function
	err := ctx.ReadResource("genezio:index:Function", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Function resources.
type functionState struct {
}

type FunctionState struct {
}

func (FunctionState) ElementType() reflect.Type {
	return reflect.TypeOf((*functionState)(nil)).Elem()
}

type functionArgs struct {
	Entry       string `pulumi:"entry"`
	Handler     string `pulumi:"handler"`
	Name        string `pulumi:"name"`
	Path        string `pulumi:"path"`
	ProjectName string `pulumi:"projectName"`
	Region      string `pulumi:"region"`
}

// The set of arguments for constructing a Function resource.
type FunctionArgs struct {
	Entry       pulumi.StringInput
	Handler     pulumi.StringInput
	Name        pulumi.StringInput
	Path        pulumi.StringInput
	ProjectName pulumi.StringInput
	Region      pulumi.StringInput
}

func (FunctionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*functionArgs)(nil)).Elem()
}

type FunctionInput interface {
	pulumi.Input

	ToFunctionOutput() FunctionOutput
	ToFunctionOutputWithContext(ctx context.Context) FunctionOutput
}

func (*Function) ElementType() reflect.Type {
	return reflect.TypeOf((**Function)(nil)).Elem()
}

func (i *Function) ToFunctionOutput() FunctionOutput {
	return i.ToFunctionOutputWithContext(context.Background())
}

func (i *Function) ToFunctionOutputWithContext(ctx context.Context) FunctionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FunctionOutput)
}

type FunctionOutput struct{ *pulumi.OutputState }

func (FunctionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Function)(nil)).Elem()
}

func (o FunctionOutput) ToFunctionOutput() FunctionOutput {
	return o
}

func (o FunctionOutput) ToFunctionOutputWithContext(ctx context.Context) FunctionOutput {
	return o
}

func (o FunctionOutput) Entry() pulumi.StringOutput {
	return o.ApplyT(func(v *Function) pulumi.StringOutput { return v.Entry }).(pulumi.StringOutput)
}

func (o FunctionOutput) FunctionId() pulumi.StringOutput {
	return o.ApplyT(func(v *Function) pulumi.StringOutput { return v.FunctionId }).(pulumi.StringOutput)
}

func (o FunctionOutput) Handler() pulumi.StringOutput {
	return o.ApplyT(func(v *Function) pulumi.StringOutput { return v.Handler }).(pulumi.StringOutput)
}

func (o FunctionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Function) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o FunctionOutput) Path() pulumi.StringOutput {
	return o.ApplyT(func(v *Function) pulumi.StringOutput { return v.Path }).(pulumi.StringOutput)
}

func (o FunctionOutput) ProjectName() pulumi.StringOutput {
	return o.ApplyT(func(v *Function) pulumi.StringOutput { return v.ProjectName }).(pulumi.StringOutput)
}

func (o FunctionOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *Function) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

func (o FunctionOutput) Url() pulumi.StringOutput {
	return o.ApplyT(func(v *Function) pulumi.StringOutput { return v.Url }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FunctionInput)(nil)).Elem(), &Function{})
	pulumi.RegisterOutputType(FunctionOutput{})
}
