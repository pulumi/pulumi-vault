// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package terraformcloud

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-vault/sdk/v6/go/vault/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-vault/sdk/v6/go/vault/terraformcloud"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := terraformcloud.NewSecretBackend(ctx, "test", &terraformcloud.SecretBackendArgs{
//				Backend:     pulumi.String("terraform"),
//				Description: pulumi.String("Manages the Terraform Cloud backend"),
//				Token:       pulumi.String("V0idfhi2iksSDU234ucdbi2nidsi..."),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// ## Import
//
// Terraform Cloud secret backends can be imported using the `backend`, e.g.
//
// ```sh
// $ pulumi import vault:terraformcloud/secretBackend:SecretBackend example terraform
// ```
type SecretBackend struct {
	pulumi.CustomResourceState

	Address  pulumi.StringPtrOutput `pulumi:"address"`
	Backend  pulumi.StringPtrOutput `pulumi:"backend"`
	BasePath pulumi.StringPtrOutput `pulumi:"basePath"`
	// The default TTL for credentials issued by this backend.
	DefaultLeaseTtlSeconds pulumi.IntPtrOutput `pulumi:"defaultLeaseTtlSeconds"`
	// A human-friendly description for this backend.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// If set, opts out of mount migration on path updates.
	// See here for more info on [Mount Migration](https://www.vaultproject.io/docs/concepts/mount-migration)
	DisableRemount pulumi.BoolPtrOutput `pulumi:"disableRemount"`
	// The maximum TTL that can be requested
	// for credentials issued by this backend.
	MaxLeaseTtlSeconds pulumi.IntPtrOutput `pulumi:"maxLeaseTtlSeconds"`
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured namespace.
	// *Available only for Vault Enterprise*.
	Namespace pulumi.StringPtrOutput `pulumi:"namespace"`
	Token     pulumi.StringPtrOutput `pulumi:"token"`
}

// NewSecretBackend registers a new resource with the given unique name, arguments, and options.
func NewSecretBackend(ctx *pulumi.Context,
	name string, args *SecretBackendArgs, opts ...pulumi.ResourceOption) (*SecretBackend, error) {
	if args == nil {
		args = &SecretBackendArgs{}
	}

	if args.Token != nil {
		args.Token = pulumi.ToSecret(args.Token).(pulumi.StringPtrInput)
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"token",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource SecretBackend
	err := ctx.RegisterResource("vault:terraformcloud/secretBackend:SecretBackend", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSecretBackend gets an existing SecretBackend resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSecretBackend(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SecretBackendState, opts ...pulumi.ResourceOption) (*SecretBackend, error) {
	var resource SecretBackend
	err := ctx.ReadResource("vault:terraformcloud/secretBackend:SecretBackend", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SecretBackend resources.
type secretBackendState struct {
	Address  *string `pulumi:"address"`
	Backend  *string `pulumi:"backend"`
	BasePath *string `pulumi:"basePath"`
	// The default TTL for credentials issued by this backend.
	DefaultLeaseTtlSeconds *int `pulumi:"defaultLeaseTtlSeconds"`
	// A human-friendly description for this backend.
	Description *string `pulumi:"description"`
	// If set, opts out of mount migration on path updates.
	// See here for more info on [Mount Migration](https://www.vaultproject.io/docs/concepts/mount-migration)
	DisableRemount *bool `pulumi:"disableRemount"`
	// The maximum TTL that can be requested
	// for credentials issued by this backend.
	MaxLeaseTtlSeconds *int `pulumi:"maxLeaseTtlSeconds"`
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured namespace.
	// *Available only for Vault Enterprise*.
	Namespace *string `pulumi:"namespace"`
	Token     *string `pulumi:"token"`
}

type SecretBackendState struct {
	Address  pulumi.StringPtrInput
	Backend  pulumi.StringPtrInput
	BasePath pulumi.StringPtrInput
	// The default TTL for credentials issued by this backend.
	DefaultLeaseTtlSeconds pulumi.IntPtrInput
	// A human-friendly description for this backend.
	Description pulumi.StringPtrInput
	// If set, opts out of mount migration on path updates.
	// See here for more info on [Mount Migration](https://www.vaultproject.io/docs/concepts/mount-migration)
	DisableRemount pulumi.BoolPtrInput
	// The maximum TTL that can be requested
	// for credentials issued by this backend.
	MaxLeaseTtlSeconds pulumi.IntPtrInput
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured namespace.
	// *Available only for Vault Enterprise*.
	Namespace pulumi.StringPtrInput
	Token     pulumi.StringPtrInput
}

func (SecretBackendState) ElementType() reflect.Type {
	return reflect.TypeOf((*secretBackendState)(nil)).Elem()
}

type secretBackendArgs struct {
	Address  *string `pulumi:"address"`
	Backend  *string `pulumi:"backend"`
	BasePath *string `pulumi:"basePath"`
	// The default TTL for credentials issued by this backend.
	DefaultLeaseTtlSeconds *int `pulumi:"defaultLeaseTtlSeconds"`
	// A human-friendly description for this backend.
	Description *string `pulumi:"description"`
	// If set, opts out of mount migration on path updates.
	// See here for more info on [Mount Migration](https://www.vaultproject.io/docs/concepts/mount-migration)
	DisableRemount *bool `pulumi:"disableRemount"`
	// The maximum TTL that can be requested
	// for credentials issued by this backend.
	MaxLeaseTtlSeconds *int `pulumi:"maxLeaseTtlSeconds"`
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured namespace.
	// *Available only for Vault Enterprise*.
	Namespace *string `pulumi:"namespace"`
	Token     *string `pulumi:"token"`
}

// The set of arguments for constructing a SecretBackend resource.
type SecretBackendArgs struct {
	Address  pulumi.StringPtrInput
	Backend  pulumi.StringPtrInput
	BasePath pulumi.StringPtrInput
	// The default TTL for credentials issued by this backend.
	DefaultLeaseTtlSeconds pulumi.IntPtrInput
	// A human-friendly description for this backend.
	Description pulumi.StringPtrInput
	// If set, opts out of mount migration on path updates.
	// See here for more info on [Mount Migration](https://www.vaultproject.io/docs/concepts/mount-migration)
	DisableRemount pulumi.BoolPtrInput
	// The maximum TTL that can be requested
	// for credentials issued by this backend.
	MaxLeaseTtlSeconds pulumi.IntPtrInput
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured namespace.
	// *Available only for Vault Enterprise*.
	Namespace pulumi.StringPtrInput
	Token     pulumi.StringPtrInput
}

func (SecretBackendArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*secretBackendArgs)(nil)).Elem()
}

type SecretBackendInput interface {
	pulumi.Input

	ToSecretBackendOutput() SecretBackendOutput
	ToSecretBackendOutputWithContext(ctx context.Context) SecretBackendOutput
}

func (*SecretBackend) ElementType() reflect.Type {
	return reflect.TypeOf((**SecretBackend)(nil)).Elem()
}

func (i *SecretBackend) ToSecretBackendOutput() SecretBackendOutput {
	return i.ToSecretBackendOutputWithContext(context.Background())
}

func (i *SecretBackend) ToSecretBackendOutputWithContext(ctx context.Context) SecretBackendOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecretBackendOutput)
}

// SecretBackendArrayInput is an input type that accepts SecretBackendArray and SecretBackendArrayOutput values.
// You can construct a concrete instance of `SecretBackendArrayInput` via:
//
//	SecretBackendArray{ SecretBackendArgs{...} }
type SecretBackendArrayInput interface {
	pulumi.Input

	ToSecretBackendArrayOutput() SecretBackendArrayOutput
	ToSecretBackendArrayOutputWithContext(context.Context) SecretBackendArrayOutput
}

type SecretBackendArray []SecretBackendInput

func (SecretBackendArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SecretBackend)(nil)).Elem()
}

func (i SecretBackendArray) ToSecretBackendArrayOutput() SecretBackendArrayOutput {
	return i.ToSecretBackendArrayOutputWithContext(context.Background())
}

func (i SecretBackendArray) ToSecretBackendArrayOutputWithContext(ctx context.Context) SecretBackendArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecretBackendArrayOutput)
}

// SecretBackendMapInput is an input type that accepts SecretBackendMap and SecretBackendMapOutput values.
// You can construct a concrete instance of `SecretBackendMapInput` via:
//
//	SecretBackendMap{ "key": SecretBackendArgs{...} }
type SecretBackendMapInput interface {
	pulumi.Input

	ToSecretBackendMapOutput() SecretBackendMapOutput
	ToSecretBackendMapOutputWithContext(context.Context) SecretBackendMapOutput
}

type SecretBackendMap map[string]SecretBackendInput

func (SecretBackendMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SecretBackend)(nil)).Elem()
}

func (i SecretBackendMap) ToSecretBackendMapOutput() SecretBackendMapOutput {
	return i.ToSecretBackendMapOutputWithContext(context.Background())
}

func (i SecretBackendMap) ToSecretBackendMapOutputWithContext(ctx context.Context) SecretBackendMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecretBackendMapOutput)
}

type SecretBackendOutput struct{ *pulumi.OutputState }

func (SecretBackendOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SecretBackend)(nil)).Elem()
}

func (o SecretBackendOutput) ToSecretBackendOutput() SecretBackendOutput {
	return o
}

func (o SecretBackendOutput) ToSecretBackendOutputWithContext(ctx context.Context) SecretBackendOutput {
	return o
}

func (o SecretBackendOutput) Address() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecretBackend) pulumi.StringPtrOutput { return v.Address }).(pulumi.StringPtrOutput)
}

func (o SecretBackendOutput) Backend() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecretBackend) pulumi.StringPtrOutput { return v.Backend }).(pulumi.StringPtrOutput)
}

func (o SecretBackendOutput) BasePath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecretBackend) pulumi.StringPtrOutput { return v.BasePath }).(pulumi.StringPtrOutput)
}

// The default TTL for credentials issued by this backend.
func (o SecretBackendOutput) DefaultLeaseTtlSeconds() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SecretBackend) pulumi.IntPtrOutput { return v.DefaultLeaseTtlSeconds }).(pulumi.IntPtrOutput)
}

// A human-friendly description for this backend.
func (o SecretBackendOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecretBackend) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// If set, opts out of mount migration on path updates.
// See here for more info on [Mount Migration](https://www.vaultproject.io/docs/concepts/mount-migration)
func (o SecretBackendOutput) DisableRemount() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SecretBackend) pulumi.BoolPtrOutput { return v.DisableRemount }).(pulumi.BoolPtrOutput)
}

// The maximum TTL that can be requested
// for credentials issued by this backend.
func (o SecretBackendOutput) MaxLeaseTtlSeconds() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SecretBackend) pulumi.IntPtrOutput { return v.MaxLeaseTtlSeconds }).(pulumi.IntPtrOutput)
}

// The namespace to provision the resource in.
// The value should not contain leading or trailing forward slashes.
// The `namespace` is always relative to the provider's configured namespace.
// *Available only for Vault Enterprise*.
func (o SecretBackendOutput) Namespace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecretBackend) pulumi.StringPtrOutput { return v.Namespace }).(pulumi.StringPtrOutput)
}

func (o SecretBackendOutput) Token() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecretBackend) pulumi.StringPtrOutput { return v.Token }).(pulumi.StringPtrOutput)
}

type SecretBackendArrayOutput struct{ *pulumi.OutputState }

func (SecretBackendArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SecretBackend)(nil)).Elem()
}

func (o SecretBackendArrayOutput) ToSecretBackendArrayOutput() SecretBackendArrayOutput {
	return o
}

func (o SecretBackendArrayOutput) ToSecretBackendArrayOutputWithContext(ctx context.Context) SecretBackendArrayOutput {
	return o
}

func (o SecretBackendArrayOutput) Index(i pulumi.IntInput) SecretBackendOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SecretBackend {
		return vs[0].([]*SecretBackend)[vs[1].(int)]
	}).(SecretBackendOutput)
}

type SecretBackendMapOutput struct{ *pulumi.OutputState }

func (SecretBackendMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SecretBackend)(nil)).Elem()
}

func (o SecretBackendMapOutput) ToSecretBackendMapOutput() SecretBackendMapOutput {
	return o
}

func (o SecretBackendMapOutput) ToSecretBackendMapOutputWithContext(ctx context.Context) SecretBackendMapOutput {
	return o
}

func (o SecretBackendMapOutput) MapIndex(k pulumi.StringInput) SecretBackendOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SecretBackend {
		return vs[0].(map[string]*SecretBackend)[vs[1].(string)]
	}).(SecretBackendOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SecretBackendInput)(nil)).Elem(), &SecretBackend{})
	pulumi.RegisterInputType(reflect.TypeOf((*SecretBackendArrayInput)(nil)).Elem(), SecretBackendArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SecretBackendMapInput)(nil)).Elem(), SecretBackendMap{})
	pulumi.RegisterOutputType(SecretBackendOutput{})
	pulumi.RegisterOutputType(SecretBackendArrayOutput{})
	pulumi.RegisterOutputType(SecretBackendMapOutput{})
}
