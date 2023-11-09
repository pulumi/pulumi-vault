// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vault

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-vault/sdk/v5/go/vault/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// ## Import
//
// Auth methods can be imported using the `path`, e.g.
//
// ```sh
//
//	$ pulumi import vault:index/authBackend:AuthBackend example github
//
// ```
type AuthBackend struct {
	pulumi.CustomResourceState

	// The accessor for this auth method
	Accessor pulumi.StringOutput `pulumi:"accessor"`
	// A description of the auth method.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// If set, opts out of mount migration on path updates.
	// See here for more info on [Mount Migration](https://www.vaultproject.io/docs/concepts/mount-migration)
	DisableRemount pulumi.BoolPtrOutput `pulumi:"disableRemount"`
	// Specifies if the auth method is local only.
	Local pulumi.BoolPtrOutput `pulumi:"local"`
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
	// *Available only for Vault Enterprise*.
	Namespace pulumi.StringPtrOutput `pulumi:"namespace"`
	// The path to mount the auth method — this defaults to the name of the type.
	Path pulumi.StringOutput `pulumi:"path"`
	// Extra configuration block. Structure is documented below.
	//
	// The `tune` block is used to tune the auth backend:
	Tune AuthBackendTuneOutput `pulumi:"tune"`
	// The name of the auth method type.
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
	opts = internal.PkgResourceDefaultOpts(opts)
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
	// A description of the auth method.
	Description *string `pulumi:"description"`
	// If set, opts out of mount migration on path updates.
	// See here for more info on [Mount Migration](https://www.vaultproject.io/docs/concepts/mount-migration)
	DisableRemount *bool `pulumi:"disableRemount"`
	// Specifies if the auth method is local only.
	Local *bool `pulumi:"local"`
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
	// *Available only for Vault Enterprise*.
	Namespace *string `pulumi:"namespace"`
	// The path to mount the auth method — this defaults to the name of the type.
	Path *string `pulumi:"path"`
	// Extra configuration block. Structure is documented below.
	//
	// The `tune` block is used to tune the auth backend:
	Tune *AuthBackendTune `pulumi:"tune"`
	// The name of the auth method type.
	Type *string `pulumi:"type"`
}

type AuthBackendState struct {
	// The accessor for this auth method
	Accessor pulumi.StringPtrInput
	// A description of the auth method.
	Description pulumi.StringPtrInput
	// If set, opts out of mount migration on path updates.
	// See here for more info on [Mount Migration](https://www.vaultproject.io/docs/concepts/mount-migration)
	DisableRemount pulumi.BoolPtrInput
	// Specifies if the auth method is local only.
	Local pulumi.BoolPtrInput
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
	// *Available only for Vault Enterprise*.
	Namespace pulumi.StringPtrInput
	// The path to mount the auth method — this defaults to the name of the type.
	Path pulumi.StringPtrInput
	// Extra configuration block. Structure is documented below.
	//
	// The `tune` block is used to tune the auth backend:
	Tune AuthBackendTunePtrInput
	// The name of the auth method type.
	Type pulumi.StringPtrInput
}

func (AuthBackendState) ElementType() reflect.Type {
	return reflect.TypeOf((*authBackendState)(nil)).Elem()
}

type authBackendArgs struct {
	// A description of the auth method.
	Description *string `pulumi:"description"`
	// If set, opts out of mount migration on path updates.
	// See here for more info on [Mount Migration](https://www.vaultproject.io/docs/concepts/mount-migration)
	DisableRemount *bool `pulumi:"disableRemount"`
	// Specifies if the auth method is local only.
	Local *bool `pulumi:"local"`
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
	// *Available only for Vault Enterprise*.
	Namespace *string `pulumi:"namespace"`
	// The path to mount the auth method — this defaults to the name of the type.
	Path *string `pulumi:"path"`
	// Extra configuration block. Structure is documented below.
	//
	// The `tune` block is used to tune the auth backend:
	Tune *AuthBackendTune `pulumi:"tune"`
	// The name of the auth method type.
	Type string `pulumi:"type"`
}

// The set of arguments for constructing a AuthBackend resource.
type AuthBackendArgs struct {
	// A description of the auth method.
	Description pulumi.StringPtrInput
	// If set, opts out of mount migration on path updates.
	// See here for more info on [Mount Migration](https://www.vaultproject.io/docs/concepts/mount-migration)
	DisableRemount pulumi.BoolPtrInput
	// Specifies if the auth method is local only.
	Local pulumi.BoolPtrInput
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
	// *Available only for Vault Enterprise*.
	Namespace pulumi.StringPtrInput
	// The path to mount the auth method — this defaults to the name of the type.
	Path pulumi.StringPtrInput
	// Extra configuration block. Structure is documented below.
	//
	// The `tune` block is used to tune the auth backend:
	Tune AuthBackendTunePtrInput
	// The name of the auth method type.
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

func (i *AuthBackend) ToOutput(ctx context.Context) pulumix.Output[*AuthBackend] {
	return pulumix.Output[*AuthBackend]{
		OutputState: i.ToAuthBackendOutputWithContext(ctx).OutputState,
	}
}

// AuthBackendArrayInput is an input type that accepts AuthBackendArray and AuthBackendArrayOutput values.
// You can construct a concrete instance of `AuthBackendArrayInput` via:
//
//	AuthBackendArray{ AuthBackendArgs{...} }
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

func (i AuthBackendArray) ToOutput(ctx context.Context) pulumix.Output[[]*AuthBackend] {
	return pulumix.Output[[]*AuthBackend]{
		OutputState: i.ToAuthBackendArrayOutputWithContext(ctx).OutputState,
	}
}

// AuthBackendMapInput is an input type that accepts AuthBackendMap and AuthBackendMapOutput values.
// You can construct a concrete instance of `AuthBackendMapInput` via:
//
//	AuthBackendMap{ "key": AuthBackendArgs{...} }
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

func (i AuthBackendMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*AuthBackend] {
	return pulumix.Output[map[string]*AuthBackend]{
		OutputState: i.ToAuthBackendMapOutputWithContext(ctx).OutputState,
	}
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

func (o AuthBackendOutput) ToOutput(ctx context.Context) pulumix.Output[*AuthBackend] {
	return pulumix.Output[*AuthBackend]{
		OutputState: o.OutputState,
	}
}

// The accessor for this auth method
func (o AuthBackendOutput) Accessor() pulumi.StringOutput {
	return o.ApplyT(func(v *AuthBackend) pulumi.StringOutput { return v.Accessor }).(pulumi.StringOutput)
}

// A description of the auth method.
func (o AuthBackendOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AuthBackend) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// If set, opts out of mount migration on path updates.
// See here for more info on [Mount Migration](https://www.vaultproject.io/docs/concepts/mount-migration)
func (o AuthBackendOutput) DisableRemount() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AuthBackend) pulumi.BoolPtrOutput { return v.DisableRemount }).(pulumi.BoolPtrOutput)
}

// Specifies if the auth method is local only.
func (o AuthBackendOutput) Local() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AuthBackend) pulumi.BoolPtrOutput { return v.Local }).(pulumi.BoolPtrOutput)
}

// The namespace to provision the resource in.
// The value should not contain leading or trailing forward slashes.
// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
// *Available only for Vault Enterprise*.
func (o AuthBackendOutput) Namespace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AuthBackend) pulumi.StringPtrOutput { return v.Namespace }).(pulumi.StringPtrOutput)
}

// The path to mount the auth method — this defaults to the name of the type.
func (o AuthBackendOutput) Path() pulumi.StringOutput {
	return o.ApplyT(func(v *AuthBackend) pulumi.StringOutput { return v.Path }).(pulumi.StringOutput)
}

// Extra configuration block. Structure is documented below.
//
// The `tune` block is used to tune the auth backend:
func (o AuthBackendOutput) Tune() AuthBackendTuneOutput {
	return o.ApplyT(func(v *AuthBackend) AuthBackendTuneOutput { return v.Tune }).(AuthBackendTuneOutput)
}

// The name of the auth method type.
func (o AuthBackendOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *AuthBackend) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
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

func (o AuthBackendArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*AuthBackend] {
	return pulumix.Output[[]*AuthBackend]{
		OutputState: o.OutputState,
	}
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

func (o AuthBackendMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*AuthBackend] {
	return pulumix.Output[map[string]*AuthBackend]{
		OutputState: o.OutputState,
	}
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
