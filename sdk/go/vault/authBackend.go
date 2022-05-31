// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vault

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## Import
//
// Auth methods can be imported using the `path`, e.g.
//
// ```sh
//  $ pulumi import vault:index/authBackend:AuthBackend example github
// ```
type AuthBackend struct {
	pulumi.CustomResourceState

	// The accessor for this auth method
	Accessor pulumi.StringOutput `pulumi:"accessor"`
	// A description of the auth method
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Specifies if the auth method is local only.
	Local pulumi.BoolPtrOutput `pulumi:"local"`
	// The path to mount the auth method — this defaults to the name of the type
	Path pulumi.StringOutput `pulumi:"path"`
	// Extra configuration block. Structure is documented below.
	Tune AuthBackendTuneOutput `pulumi:"tune"`
	// The name of the auth method type
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewAuthBackend registers a new resource with the given unique name, arguments, and options.
func NewAuthBackend(ctx *pulumi.Context,
	name string, args *AuthBackendArgs, opts ...pulumi.ResourceOption) (*AuthBackend, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Type == nil {
		return nil, errors.New("invalid value for required argument 'Type'")
	}
	var resource AuthBackend
	err := ctx.RegisterResource("vault:index/authBackend:AuthBackend", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAuthBackend gets an existing AuthBackend resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAuthBackend(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AuthBackendState, opts ...pulumi.ResourceOption) (*AuthBackend, error) {
	var resource AuthBackend
	err := ctx.ReadResource("vault:index/authBackend:AuthBackend", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AuthBackend resources.
type authBackendState struct {
	// The accessor for this auth method
	Accessor *string `pulumi:"accessor"`
	// A description of the auth method
	Description *string `pulumi:"description"`
	// Specifies if the auth method is local only.
	Local *bool `pulumi:"local"`
	// The path to mount the auth method — this defaults to the name of the type
	Path *string `pulumi:"path"`
	// Extra configuration block. Structure is documented below.
	Tune *AuthBackendTune `pulumi:"tune"`
	// The name of the auth method type
	Type *string `pulumi:"type"`
}

type AuthBackendState struct {
	// The accessor for this auth method
	Accessor pulumi.StringPtrInput
	// A description of the auth method
	Description pulumi.StringPtrInput
	// Specifies if the auth method is local only.
	Local pulumi.BoolPtrInput
	// The path to mount the auth method — this defaults to the name of the type
	Path pulumi.StringPtrInput
	// Extra configuration block. Structure is documented below.
	Tune AuthBackendTunePtrInput
	// The name of the auth method type
	Type pulumi.StringPtrInput
}

func (AuthBackendState) ElementType() reflect.Type {
	return reflect.TypeOf((*authBackendState)(nil)).Elem()
}

type authBackendArgs struct {
	// A description of the auth method
	Description *string `pulumi:"description"`
	// Specifies if the auth method is local only.
	Local *bool `pulumi:"local"`
	// The path to mount the auth method — this defaults to the name of the type
	Path *string `pulumi:"path"`
	// Extra configuration block. Structure is documented below.
	Tune *AuthBackendTune `pulumi:"tune"`
	// The name of the auth method type
	Type string `pulumi:"type"`
}

// The set of arguments for constructing a AuthBackend resource.
type AuthBackendArgs struct {
	// A description of the auth method
	Description pulumi.StringPtrInput
	// Specifies if the auth method is local only.
	Local pulumi.BoolPtrInput
	// The path to mount the auth method — this defaults to the name of the type
	Path pulumi.StringPtrInput
	// Extra configuration block. Structure is documented below.
	Tune AuthBackendTunePtrInput
	// The name of the auth method type
	Type pulumi.StringInput
}

func (AuthBackendArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*authBackendArgs)(nil)).Elem()
}

type AuthBackendInput interface {
	pulumi.Input

	ToAuthBackendOutput() AuthBackendOutput
	ToAuthBackendOutputWithContext(ctx context.Context) AuthBackendOutput
}

func (*AuthBackend) ElementType() reflect.Type {
	return reflect.TypeOf((**AuthBackend)(nil)).Elem()
}

func (i *AuthBackend) ToAuthBackendOutput() AuthBackendOutput {
	return i.ToAuthBackendOutputWithContext(context.Background())
}

func (i *AuthBackend) ToAuthBackendOutputWithContext(ctx context.Context) AuthBackendOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AuthBackendOutput)
}

// AuthBackendArrayInput is an input type that accepts AuthBackendArray and AuthBackendArrayOutput values.
// You can construct a concrete instance of `AuthBackendArrayInput` via:
//
//          AuthBackendArray{ AuthBackendArgs{...} }
type AuthBackendArrayInput interface {
	pulumi.Input

	ToAuthBackendArrayOutput() AuthBackendArrayOutput
	ToAuthBackendArrayOutputWithContext(context.Context) AuthBackendArrayOutput
}

type AuthBackendArray []AuthBackendInput

func (AuthBackendArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AuthBackend)(nil)).Elem()
}

func (i AuthBackendArray) ToAuthBackendArrayOutput() AuthBackendArrayOutput {
	return i.ToAuthBackendArrayOutputWithContext(context.Background())
}

func (i AuthBackendArray) ToAuthBackendArrayOutputWithContext(ctx context.Context) AuthBackendArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AuthBackendArrayOutput)
}

// AuthBackendMapInput is an input type that accepts AuthBackendMap and AuthBackendMapOutput values.
// You can construct a concrete instance of `AuthBackendMapInput` via:
//
//          AuthBackendMap{ "key": AuthBackendArgs{...} }
type AuthBackendMapInput interface {
	pulumi.Input

	ToAuthBackendMapOutput() AuthBackendMapOutput
	ToAuthBackendMapOutputWithContext(context.Context) AuthBackendMapOutput
}

type AuthBackendMap map[string]AuthBackendInput

func (AuthBackendMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AuthBackend)(nil)).Elem()
}

func (i AuthBackendMap) ToAuthBackendMapOutput() AuthBackendMapOutput {
	return i.ToAuthBackendMapOutputWithContext(context.Background())
}

func (i AuthBackendMap) ToAuthBackendMapOutputWithContext(ctx context.Context) AuthBackendMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AuthBackendMapOutput)
}

type AuthBackendOutput struct{ *pulumi.OutputState }

func (AuthBackendOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AuthBackend)(nil)).Elem()
}

func (o AuthBackendOutput) ToAuthBackendOutput() AuthBackendOutput {
	return o
}

func (o AuthBackendOutput) ToAuthBackendOutputWithContext(ctx context.Context) AuthBackendOutput {
	return o
}

type AuthBackendArrayOutput struct{ *pulumi.OutputState }

func (AuthBackendArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AuthBackend)(nil)).Elem()
}

func (o AuthBackendArrayOutput) ToAuthBackendArrayOutput() AuthBackendArrayOutput {
	return o
}

func (o AuthBackendArrayOutput) ToAuthBackendArrayOutputWithContext(ctx context.Context) AuthBackendArrayOutput {
	return o
}

func (o AuthBackendArrayOutput) Index(i pulumi.IntInput) AuthBackendOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *AuthBackend {
		return vs[0].([]*AuthBackend)[vs[1].(int)]
	}).(AuthBackendOutput)
}

type AuthBackendMapOutput struct{ *pulumi.OutputState }

func (AuthBackendMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AuthBackend)(nil)).Elem()
}

func (o AuthBackendMapOutput) ToAuthBackendMapOutput() AuthBackendMapOutput {
	return o
}

func (o AuthBackendMapOutput) ToAuthBackendMapOutputWithContext(ctx context.Context) AuthBackendMapOutput {
	return o
}

func (o AuthBackendMapOutput) MapIndex(k pulumi.StringInput) AuthBackendOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *AuthBackend {
		return vs[0].(map[string]*AuthBackend)[vs[1].(string)]
	}).(AuthBackendOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AuthBackendInput)(nil)).Elem(), &AuthBackend{})
	pulumi.RegisterInputType(reflect.TypeOf((*AuthBackendArrayInput)(nil)).Elem(), AuthBackendArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AuthBackendMapInput)(nil)).Elem(), AuthBackendMap{})
	pulumi.RegisterOutputType(AuthBackendOutput{})
	pulumi.RegisterOutputType(AuthBackendArrayOutput{})
	pulumi.RegisterOutputType(AuthBackendMapOutput{})
}
