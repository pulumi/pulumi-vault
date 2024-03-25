// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package transit

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-vault/sdk/v6/go/vault/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configure the cache for the Transit Secret Backend in Vault.
//
// ## Example Usage
//
// <!--Start PulumiCodeChooser -->
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-vault/sdk/v6/go/vault"
//	"github.com/pulumi/pulumi-vault/sdk/v6/go/vault/transit"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			transit, err := vault.NewMount(ctx, "transit", &vault.MountArgs{
//				Path:                   pulumi.String("transit"),
//				Type:                   pulumi.String("transit"),
//				Description:            pulumi.String("Example description"),
//				DefaultLeaseTtlSeconds: pulumi.Int(3600),
//				MaxLeaseTtlSeconds:     pulumi.Int(86400),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = transit.NewSecretCacheConfig(ctx, "cfg", &transit.SecretCacheConfigArgs{
//				Backend: transit.Path,
//				Size:    pulumi.Int(500),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// <!--End PulumiCodeChooser -->
type SecretCacheConfig struct {
	pulumi.CustomResourceState

	// The path the transit secret backend is mounted at, with no leading or trailing `/`s.
	Backend pulumi.StringOutput `pulumi:"backend"`
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
	// *Available only for Vault Enterprise*.
	Namespace pulumi.StringPtrOutput `pulumi:"namespace"`
	// The number of cache entries. 0 means unlimited.
	Size pulumi.IntOutput `pulumi:"size"`
}

// NewSecretCacheConfig registers a new resource with the given unique name, arguments, and options.
func NewSecretCacheConfig(ctx *pulumi.Context,
	name string, args *SecretCacheConfigArgs, opts ...pulumi.ResourceOption) (*SecretCacheConfig, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Backend == nil {
		return nil, errors.New("invalid value for required argument 'Backend'")
	}
	if args.Size == nil {
		return nil, errors.New("invalid value for required argument 'Size'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource SecretCacheConfig
	err := ctx.RegisterResource("vault:transit/secretCacheConfig:SecretCacheConfig", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSecretCacheConfig gets an existing SecretCacheConfig resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSecretCacheConfig(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SecretCacheConfigState, opts ...pulumi.ResourceOption) (*SecretCacheConfig, error) {
	var resource SecretCacheConfig
	err := ctx.ReadResource("vault:transit/secretCacheConfig:SecretCacheConfig", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SecretCacheConfig resources.
type secretCacheConfigState struct {
	// The path the transit secret backend is mounted at, with no leading or trailing `/`s.
	Backend *string `pulumi:"backend"`
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
	// *Available only for Vault Enterprise*.
	Namespace *string `pulumi:"namespace"`
	// The number of cache entries. 0 means unlimited.
	Size *int `pulumi:"size"`
}

type SecretCacheConfigState struct {
	// The path the transit secret backend is mounted at, with no leading or trailing `/`s.
	Backend pulumi.StringPtrInput
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
	// *Available only for Vault Enterprise*.
	Namespace pulumi.StringPtrInput
	// The number of cache entries. 0 means unlimited.
	Size pulumi.IntPtrInput
}

func (SecretCacheConfigState) ElementType() reflect.Type {
	return reflect.TypeOf((*secretCacheConfigState)(nil)).Elem()
}

type secretCacheConfigArgs struct {
	// The path the transit secret backend is mounted at, with no leading or trailing `/`s.
	Backend string `pulumi:"backend"`
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
	// *Available only for Vault Enterprise*.
	Namespace *string `pulumi:"namespace"`
	// The number of cache entries. 0 means unlimited.
	Size int `pulumi:"size"`
}

// The set of arguments for constructing a SecretCacheConfig resource.
type SecretCacheConfigArgs struct {
	// The path the transit secret backend is mounted at, with no leading or trailing `/`s.
	Backend pulumi.StringInput
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
	// *Available only for Vault Enterprise*.
	Namespace pulumi.StringPtrInput
	// The number of cache entries. 0 means unlimited.
	Size pulumi.IntInput
}

func (SecretCacheConfigArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*secretCacheConfigArgs)(nil)).Elem()
}

type SecretCacheConfigInput interface {
	pulumi.Input

	ToSecretCacheConfigOutput() SecretCacheConfigOutput
	ToSecretCacheConfigOutputWithContext(ctx context.Context) SecretCacheConfigOutput
}

func (*SecretCacheConfig) ElementType() reflect.Type {
	return reflect.TypeOf((**SecretCacheConfig)(nil)).Elem()
}

func (i *SecretCacheConfig) ToSecretCacheConfigOutput() SecretCacheConfigOutput {
	return i.ToSecretCacheConfigOutputWithContext(context.Background())
}

func (i *SecretCacheConfig) ToSecretCacheConfigOutputWithContext(ctx context.Context) SecretCacheConfigOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecretCacheConfigOutput)
}

// SecretCacheConfigArrayInput is an input type that accepts SecretCacheConfigArray and SecretCacheConfigArrayOutput values.
// You can construct a concrete instance of `SecretCacheConfigArrayInput` via:
//
//	SecretCacheConfigArray{ SecretCacheConfigArgs{...} }
type SecretCacheConfigArrayInput interface {
	pulumi.Input

	ToSecretCacheConfigArrayOutput() SecretCacheConfigArrayOutput
	ToSecretCacheConfigArrayOutputWithContext(context.Context) SecretCacheConfigArrayOutput
}

type SecretCacheConfigArray []SecretCacheConfigInput

func (SecretCacheConfigArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SecretCacheConfig)(nil)).Elem()
}

func (i SecretCacheConfigArray) ToSecretCacheConfigArrayOutput() SecretCacheConfigArrayOutput {
	return i.ToSecretCacheConfigArrayOutputWithContext(context.Background())
}

func (i SecretCacheConfigArray) ToSecretCacheConfigArrayOutputWithContext(ctx context.Context) SecretCacheConfigArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecretCacheConfigArrayOutput)
}

// SecretCacheConfigMapInput is an input type that accepts SecretCacheConfigMap and SecretCacheConfigMapOutput values.
// You can construct a concrete instance of `SecretCacheConfigMapInput` via:
//
//	SecretCacheConfigMap{ "key": SecretCacheConfigArgs{...} }
type SecretCacheConfigMapInput interface {
	pulumi.Input

	ToSecretCacheConfigMapOutput() SecretCacheConfigMapOutput
	ToSecretCacheConfigMapOutputWithContext(context.Context) SecretCacheConfigMapOutput
}

type SecretCacheConfigMap map[string]SecretCacheConfigInput

func (SecretCacheConfigMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SecretCacheConfig)(nil)).Elem()
}

func (i SecretCacheConfigMap) ToSecretCacheConfigMapOutput() SecretCacheConfigMapOutput {
	return i.ToSecretCacheConfigMapOutputWithContext(context.Background())
}

func (i SecretCacheConfigMap) ToSecretCacheConfigMapOutputWithContext(ctx context.Context) SecretCacheConfigMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecretCacheConfigMapOutput)
}

type SecretCacheConfigOutput struct{ *pulumi.OutputState }

func (SecretCacheConfigOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SecretCacheConfig)(nil)).Elem()
}

func (o SecretCacheConfigOutput) ToSecretCacheConfigOutput() SecretCacheConfigOutput {
	return o
}

func (o SecretCacheConfigOutput) ToSecretCacheConfigOutputWithContext(ctx context.Context) SecretCacheConfigOutput {
	return o
}

// The path the transit secret backend is mounted at, with no leading or trailing `/`s.
func (o SecretCacheConfigOutput) Backend() pulumi.StringOutput {
	return o.ApplyT(func(v *SecretCacheConfig) pulumi.StringOutput { return v.Backend }).(pulumi.StringOutput)
}

// The namespace to provision the resource in.
// The value should not contain leading or trailing forward slashes.
// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
// *Available only for Vault Enterprise*.
func (o SecretCacheConfigOutput) Namespace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecretCacheConfig) pulumi.StringPtrOutput { return v.Namespace }).(pulumi.StringPtrOutput)
}

// The number of cache entries. 0 means unlimited.
func (o SecretCacheConfigOutput) Size() pulumi.IntOutput {
	return o.ApplyT(func(v *SecretCacheConfig) pulumi.IntOutput { return v.Size }).(pulumi.IntOutput)
}

type SecretCacheConfigArrayOutput struct{ *pulumi.OutputState }

func (SecretCacheConfigArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SecretCacheConfig)(nil)).Elem()
}

func (o SecretCacheConfigArrayOutput) ToSecretCacheConfigArrayOutput() SecretCacheConfigArrayOutput {
	return o
}

func (o SecretCacheConfigArrayOutput) ToSecretCacheConfigArrayOutputWithContext(ctx context.Context) SecretCacheConfigArrayOutput {
	return o
}

func (o SecretCacheConfigArrayOutput) Index(i pulumi.IntInput) SecretCacheConfigOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SecretCacheConfig {
		return vs[0].([]*SecretCacheConfig)[vs[1].(int)]
	}).(SecretCacheConfigOutput)
}

type SecretCacheConfigMapOutput struct{ *pulumi.OutputState }

func (SecretCacheConfigMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SecretCacheConfig)(nil)).Elem()
}

func (o SecretCacheConfigMapOutput) ToSecretCacheConfigMapOutput() SecretCacheConfigMapOutput {
	return o
}

func (o SecretCacheConfigMapOutput) ToSecretCacheConfigMapOutputWithContext(ctx context.Context) SecretCacheConfigMapOutput {
	return o
}

func (o SecretCacheConfigMapOutput) MapIndex(k pulumi.StringInput) SecretCacheConfigOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SecretCacheConfig {
		return vs[0].(map[string]*SecretCacheConfig)[vs[1].(string)]
	}).(SecretCacheConfigOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SecretCacheConfigInput)(nil)).Elem(), &SecretCacheConfig{})
	pulumi.RegisterInputType(reflect.TypeOf((*SecretCacheConfigArrayInput)(nil)).Elem(), SecretCacheConfigArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SecretCacheConfigMapInput)(nil)).Elem(), SecretCacheConfigMap{})
	pulumi.RegisterOutputType(SecretCacheConfigOutput{})
	pulumi.RegisterOutputType(SecretCacheConfigArrayOutput{})
	pulumi.RegisterOutputType(SecretCacheConfigMapOutput{})
}
