// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package pkisecret

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-vault/sdk/v6/go/vault/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Allows setting the duration for which the generated CRL should be marked valid. If the CRL is disabled, it will return a signed but zero-length CRL for any request. If enabled, it will re-build the CRL.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-vault/sdk/v6/go/vault"
//	"github.com/pulumi/pulumi-vault/sdk/v6/go/vault/pkisecret"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			pki, err := vault.NewMount(ctx, "pki", &vault.MountArgs{
//				Path:                   pulumi.String("%s"),
//				Type:                   pulumi.String("pki"),
//				DefaultLeaseTtlSeconds: pulumi.Int(3600),
//				MaxLeaseTtlSeconds:     pulumi.Int(86400),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = pkisecret.NewSecretBackendCrlConfig(ctx, "crl_config", &pkisecret.SecretBackendCrlConfigArgs{
//				Backend: pki.Path,
//				Expiry:  pulumi.String("72h"),
//				Disable: pulumi.Bool(false),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
type SecretBackendCrlConfig struct {
	pulumi.CustomResourceState

	// Enables periodic rebuilding of the CRL upon expiry. **Vault 1.12+**
	AutoRebuild pulumi.BoolPtrOutput `pulumi:"autoRebuild"`
	// Grace period before CRL expiry to attempt rebuild of CRL. **Vault 1.12+**
	AutoRebuildGracePeriod pulumi.StringOutput `pulumi:"autoRebuildGracePeriod"`
	// The path the PKI secret backend is mounted at, with no leading or trailing `/`s.
	Backend pulumi.StringOutput `pulumi:"backend"`
	// Enable cross-cluster revocation request queues. **Vault 1.13+**
	CrossClusterRevocation pulumi.BoolOutput `pulumi:"crossClusterRevocation"`
	// Interval to check for new revocations on, to regenerate the delta CRL.
	DeltaRebuildInterval pulumi.StringOutput `pulumi:"deltaRebuildInterval"`
	// Disables or enables CRL building.
	Disable pulumi.BoolPtrOutput `pulumi:"disable"`
	// Enables building of delta CRLs with up-to-date revocation information,
	// augmenting the last complete CRL.  **Vault 1.12+**
	EnableDelta pulumi.BoolPtrOutput `pulumi:"enableDelta"`
	// Specifies the time until expiration.
	Expiry pulumi.StringPtrOutput `pulumi:"expiry"`
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
	// *Available only for Vault Enterprise*.
	Namespace pulumi.StringPtrOutput `pulumi:"namespace"`
	// Disables the OCSP responder in Vault. **Vault 1.12+**
	OcspDisable pulumi.BoolPtrOutput `pulumi:"ocspDisable"`
	// The amount of time an OCSP response can be cached for, useful for OCSP stapling
	// refresh durations. **Vault 1.12+**
	OcspExpiry pulumi.StringOutput `pulumi:"ocspExpiry"`
	// Enables unified CRL and OCSP building. **Vault 1.13+**
	UnifiedCrl pulumi.BoolOutput `pulumi:"unifiedCrl"`
	// Enables serving the unified CRL and OCSP on the existing, previously
	// cluster-local paths. **Vault 1.13+**
	UnifiedCrlOnExistingPaths pulumi.BoolOutput `pulumi:"unifiedCrlOnExistingPaths"`
}

// NewSecretBackendCrlConfig registers a new resource with the given unique name, arguments, and options.
func NewSecretBackendCrlConfig(ctx *pulumi.Context,
	name string, args *SecretBackendCrlConfigArgs, opts ...pulumi.ResourceOption) (*SecretBackendCrlConfig, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Backend == nil {
		return nil, errors.New("invalid value for required argument 'Backend'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource SecretBackendCrlConfig
	err := ctx.RegisterResource("vault:pkiSecret/secretBackendCrlConfig:SecretBackendCrlConfig", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSecretBackendCrlConfig gets an existing SecretBackendCrlConfig resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSecretBackendCrlConfig(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SecretBackendCrlConfigState, opts ...pulumi.ResourceOption) (*SecretBackendCrlConfig, error) {
	var resource SecretBackendCrlConfig
	err := ctx.ReadResource("vault:pkiSecret/secretBackendCrlConfig:SecretBackendCrlConfig", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SecretBackendCrlConfig resources.
type secretBackendCrlConfigState struct {
	// Enables periodic rebuilding of the CRL upon expiry. **Vault 1.12+**
	AutoRebuild *bool `pulumi:"autoRebuild"`
	// Grace period before CRL expiry to attempt rebuild of CRL. **Vault 1.12+**
	AutoRebuildGracePeriod *string `pulumi:"autoRebuildGracePeriod"`
	// The path the PKI secret backend is mounted at, with no leading or trailing `/`s.
	Backend *string `pulumi:"backend"`
	// Enable cross-cluster revocation request queues. **Vault 1.13+**
	CrossClusterRevocation *bool `pulumi:"crossClusterRevocation"`
	// Interval to check for new revocations on, to regenerate the delta CRL.
	DeltaRebuildInterval *string `pulumi:"deltaRebuildInterval"`
	// Disables or enables CRL building.
	Disable *bool `pulumi:"disable"`
	// Enables building of delta CRLs with up-to-date revocation information,
	// augmenting the last complete CRL.  **Vault 1.12+**
	EnableDelta *bool `pulumi:"enableDelta"`
	// Specifies the time until expiration.
	Expiry *string `pulumi:"expiry"`
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
	// *Available only for Vault Enterprise*.
	Namespace *string `pulumi:"namespace"`
	// Disables the OCSP responder in Vault. **Vault 1.12+**
	OcspDisable *bool `pulumi:"ocspDisable"`
	// The amount of time an OCSP response can be cached for, useful for OCSP stapling
	// refresh durations. **Vault 1.12+**
	OcspExpiry *string `pulumi:"ocspExpiry"`
	// Enables unified CRL and OCSP building. **Vault 1.13+**
	UnifiedCrl *bool `pulumi:"unifiedCrl"`
	// Enables serving the unified CRL and OCSP on the existing, previously
	// cluster-local paths. **Vault 1.13+**
	UnifiedCrlOnExistingPaths *bool `pulumi:"unifiedCrlOnExistingPaths"`
}

type SecretBackendCrlConfigState struct {
	// Enables periodic rebuilding of the CRL upon expiry. **Vault 1.12+**
	AutoRebuild pulumi.BoolPtrInput
	// Grace period before CRL expiry to attempt rebuild of CRL. **Vault 1.12+**
	AutoRebuildGracePeriod pulumi.StringPtrInput
	// The path the PKI secret backend is mounted at, with no leading or trailing `/`s.
	Backend pulumi.StringPtrInput
	// Enable cross-cluster revocation request queues. **Vault 1.13+**
	CrossClusterRevocation pulumi.BoolPtrInput
	// Interval to check for new revocations on, to regenerate the delta CRL.
	DeltaRebuildInterval pulumi.StringPtrInput
	// Disables or enables CRL building.
	Disable pulumi.BoolPtrInput
	// Enables building of delta CRLs with up-to-date revocation information,
	// augmenting the last complete CRL.  **Vault 1.12+**
	EnableDelta pulumi.BoolPtrInput
	// Specifies the time until expiration.
	Expiry pulumi.StringPtrInput
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
	// *Available only for Vault Enterprise*.
	Namespace pulumi.StringPtrInput
	// Disables the OCSP responder in Vault. **Vault 1.12+**
	OcspDisable pulumi.BoolPtrInput
	// The amount of time an OCSP response can be cached for, useful for OCSP stapling
	// refresh durations. **Vault 1.12+**
	OcspExpiry pulumi.StringPtrInput
	// Enables unified CRL and OCSP building. **Vault 1.13+**
	UnifiedCrl pulumi.BoolPtrInput
	// Enables serving the unified CRL and OCSP on the existing, previously
	// cluster-local paths. **Vault 1.13+**
	UnifiedCrlOnExistingPaths pulumi.BoolPtrInput
}

func (SecretBackendCrlConfigState) ElementType() reflect.Type {
	return reflect.TypeOf((*secretBackendCrlConfigState)(nil)).Elem()
}

type secretBackendCrlConfigArgs struct {
	// Enables periodic rebuilding of the CRL upon expiry. **Vault 1.12+**
	AutoRebuild *bool `pulumi:"autoRebuild"`
	// Grace period before CRL expiry to attempt rebuild of CRL. **Vault 1.12+**
	AutoRebuildGracePeriod *string `pulumi:"autoRebuildGracePeriod"`
	// The path the PKI secret backend is mounted at, with no leading or trailing `/`s.
	Backend string `pulumi:"backend"`
	// Enable cross-cluster revocation request queues. **Vault 1.13+**
	CrossClusterRevocation *bool `pulumi:"crossClusterRevocation"`
	// Interval to check for new revocations on, to regenerate the delta CRL.
	DeltaRebuildInterval *string `pulumi:"deltaRebuildInterval"`
	// Disables or enables CRL building.
	Disable *bool `pulumi:"disable"`
	// Enables building of delta CRLs with up-to-date revocation information,
	// augmenting the last complete CRL.  **Vault 1.12+**
	EnableDelta *bool `pulumi:"enableDelta"`
	// Specifies the time until expiration.
	Expiry *string `pulumi:"expiry"`
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
	// *Available only for Vault Enterprise*.
	Namespace *string `pulumi:"namespace"`
	// Disables the OCSP responder in Vault. **Vault 1.12+**
	OcspDisable *bool `pulumi:"ocspDisable"`
	// The amount of time an OCSP response can be cached for, useful for OCSP stapling
	// refresh durations. **Vault 1.12+**
	OcspExpiry *string `pulumi:"ocspExpiry"`
	// Enables unified CRL and OCSP building. **Vault 1.13+**
	UnifiedCrl *bool `pulumi:"unifiedCrl"`
	// Enables serving the unified CRL and OCSP on the existing, previously
	// cluster-local paths. **Vault 1.13+**
	UnifiedCrlOnExistingPaths *bool `pulumi:"unifiedCrlOnExistingPaths"`
}

// The set of arguments for constructing a SecretBackendCrlConfig resource.
type SecretBackendCrlConfigArgs struct {
	// Enables periodic rebuilding of the CRL upon expiry. **Vault 1.12+**
	AutoRebuild pulumi.BoolPtrInput
	// Grace period before CRL expiry to attempt rebuild of CRL. **Vault 1.12+**
	AutoRebuildGracePeriod pulumi.StringPtrInput
	// The path the PKI secret backend is mounted at, with no leading or trailing `/`s.
	Backend pulumi.StringInput
	// Enable cross-cluster revocation request queues. **Vault 1.13+**
	CrossClusterRevocation pulumi.BoolPtrInput
	// Interval to check for new revocations on, to regenerate the delta CRL.
	DeltaRebuildInterval pulumi.StringPtrInput
	// Disables or enables CRL building.
	Disable pulumi.BoolPtrInput
	// Enables building of delta CRLs with up-to-date revocation information,
	// augmenting the last complete CRL.  **Vault 1.12+**
	EnableDelta pulumi.BoolPtrInput
	// Specifies the time until expiration.
	Expiry pulumi.StringPtrInput
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
	// *Available only for Vault Enterprise*.
	Namespace pulumi.StringPtrInput
	// Disables the OCSP responder in Vault. **Vault 1.12+**
	OcspDisable pulumi.BoolPtrInput
	// The amount of time an OCSP response can be cached for, useful for OCSP stapling
	// refresh durations. **Vault 1.12+**
	OcspExpiry pulumi.StringPtrInput
	// Enables unified CRL and OCSP building. **Vault 1.13+**
	UnifiedCrl pulumi.BoolPtrInput
	// Enables serving the unified CRL and OCSP on the existing, previously
	// cluster-local paths. **Vault 1.13+**
	UnifiedCrlOnExistingPaths pulumi.BoolPtrInput
}

func (SecretBackendCrlConfigArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*secretBackendCrlConfigArgs)(nil)).Elem()
}

type SecretBackendCrlConfigInput interface {
	pulumi.Input

	ToSecretBackendCrlConfigOutput() SecretBackendCrlConfigOutput
	ToSecretBackendCrlConfigOutputWithContext(ctx context.Context) SecretBackendCrlConfigOutput
}

func (*SecretBackendCrlConfig) ElementType() reflect.Type {
	return reflect.TypeOf((**SecretBackendCrlConfig)(nil)).Elem()
}

func (i *SecretBackendCrlConfig) ToSecretBackendCrlConfigOutput() SecretBackendCrlConfigOutput {
	return i.ToSecretBackendCrlConfigOutputWithContext(context.Background())
}

func (i *SecretBackendCrlConfig) ToSecretBackendCrlConfigOutputWithContext(ctx context.Context) SecretBackendCrlConfigOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecretBackendCrlConfigOutput)
}

// SecretBackendCrlConfigArrayInput is an input type that accepts SecretBackendCrlConfigArray and SecretBackendCrlConfigArrayOutput values.
// You can construct a concrete instance of `SecretBackendCrlConfigArrayInput` via:
//
//	SecretBackendCrlConfigArray{ SecretBackendCrlConfigArgs{...} }
type SecretBackendCrlConfigArrayInput interface {
	pulumi.Input

	ToSecretBackendCrlConfigArrayOutput() SecretBackendCrlConfigArrayOutput
	ToSecretBackendCrlConfigArrayOutputWithContext(context.Context) SecretBackendCrlConfigArrayOutput
}

type SecretBackendCrlConfigArray []SecretBackendCrlConfigInput

func (SecretBackendCrlConfigArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SecretBackendCrlConfig)(nil)).Elem()
}

func (i SecretBackendCrlConfigArray) ToSecretBackendCrlConfigArrayOutput() SecretBackendCrlConfigArrayOutput {
	return i.ToSecretBackendCrlConfigArrayOutputWithContext(context.Background())
}

func (i SecretBackendCrlConfigArray) ToSecretBackendCrlConfigArrayOutputWithContext(ctx context.Context) SecretBackendCrlConfigArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecretBackendCrlConfigArrayOutput)
}

// SecretBackendCrlConfigMapInput is an input type that accepts SecretBackendCrlConfigMap and SecretBackendCrlConfigMapOutput values.
// You can construct a concrete instance of `SecretBackendCrlConfigMapInput` via:
//
//	SecretBackendCrlConfigMap{ "key": SecretBackendCrlConfigArgs{...} }
type SecretBackendCrlConfigMapInput interface {
	pulumi.Input

	ToSecretBackendCrlConfigMapOutput() SecretBackendCrlConfigMapOutput
	ToSecretBackendCrlConfigMapOutputWithContext(context.Context) SecretBackendCrlConfigMapOutput
}

type SecretBackendCrlConfigMap map[string]SecretBackendCrlConfigInput

func (SecretBackendCrlConfigMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SecretBackendCrlConfig)(nil)).Elem()
}

func (i SecretBackendCrlConfigMap) ToSecretBackendCrlConfigMapOutput() SecretBackendCrlConfigMapOutput {
	return i.ToSecretBackendCrlConfigMapOutputWithContext(context.Background())
}

func (i SecretBackendCrlConfigMap) ToSecretBackendCrlConfigMapOutputWithContext(ctx context.Context) SecretBackendCrlConfigMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecretBackendCrlConfigMapOutput)
}

type SecretBackendCrlConfigOutput struct{ *pulumi.OutputState }

func (SecretBackendCrlConfigOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SecretBackendCrlConfig)(nil)).Elem()
}

func (o SecretBackendCrlConfigOutput) ToSecretBackendCrlConfigOutput() SecretBackendCrlConfigOutput {
	return o
}

func (o SecretBackendCrlConfigOutput) ToSecretBackendCrlConfigOutputWithContext(ctx context.Context) SecretBackendCrlConfigOutput {
	return o
}

// Enables periodic rebuilding of the CRL upon expiry. **Vault 1.12+**
func (o SecretBackendCrlConfigOutput) AutoRebuild() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SecretBackendCrlConfig) pulumi.BoolPtrOutput { return v.AutoRebuild }).(pulumi.BoolPtrOutput)
}

// Grace period before CRL expiry to attempt rebuild of CRL. **Vault 1.12+**
func (o SecretBackendCrlConfigOutput) AutoRebuildGracePeriod() pulumi.StringOutput {
	return o.ApplyT(func(v *SecretBackendCrlConfig) pulumi.StringOutput { return v.AutoRebuildGracePeriod }).(pulumi.StringOutput)
}

// The path the PKI secret backend is mounted at, with no leading or trailing `/`s.
func (o SecretBackendCrlConfigOutput) Backend() pulumi.StringOutput {
	return o.ApplyT(func(v *SecretBackendCrlConfig) pulumi.StringOutput { return v.Backend }).(pulumi.StringOutput)
}

// Enable cross-cluster revocation request queues. **Vault 1.13+**
func (o SecretBackendCrlConfigOutput) CrossClusterRevocation() pulumi.BoolOutput {
	return o.ApplyT(func(v *SecretBackendCrlConfig) pulumi.BoolOutput { return v.CrossClusterRevocation }).(pulumi.BoolOutput)
}

// Interval to check for new revocations on, to regenerate the delta CRL.
func (o SecretBackendCrlConfigOutput) DeltaRebuildInterval() pulumi.StringOutput {
	return o.ApplyT(func(v *SecretBackendCrlConfig) pulumi.StringOutput { return v.DeltaRebuildInterval }).(pulumi.StringOutput)
}

// Disables or enables CRL building.
func (o SecretBackendCrlConfigOutput) Disable() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SecretBackendCrlConfig) pulumi.BoolPtrOutput { return v.Disable }).(pulumi.BoolPtrOutput)
}

// Enables building of delta CRLs with up-to-date revocation information,
// augmenting the last complete CRL.  **Vault 1.12+**
func (o SecretBackendCrlConfigOutput) EnableDelta() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SecretBackendCrlConfig) pulumi.BoolPtrOutput { return v.EnableDelta }).(pulumi.BoolPtrOutput)
}

// Specifies the time until expiration.
func (o SecretBackendCrlConfigOutput) Expiry() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecretBackendCrlConfig) pulumi.StringPtrOutput { return v.Expiry }).(pulumi.StringPtrOutput)
}

// The namespace to provision the resource in.
// The value should not contain leading or trailing forward slashes.
// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
// *Available only for Vault Enterprise*.
func (o SecretBackendCrlConfigOutput) Namespace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecretBackendCrlConfig) pulumi.StringPtrOutput { return v.Namespace }).(pulumi.StringPtrOutput)
}

// Disables the OCSP responder in Vault. **Vault 1.12+**
func (o SecretBackendCrlConfigOutput) OcspDisable() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SecretBackendCrlConfig) pulumi.BoolPtrOutput { return v.OcspDisable }).(pulumi.BoolPtrOutput)
}

// The amount of time an OCSP response can be cached for, useful for OCSP stapling
// refresh durations. **Vault 1.12+**
func (o SecretBackendCrlConfigOutput) OcspExpiry() pulumi.StringOutput {
	return o.ApplyT(func(v *SecretBackendCrlConfig) pulumi.StringOutput { return v.OcspExpiry }).(pulumi.StringOutput)
}

// Enables unified CRL and OCSP building. **Vault 1.13+**
func (o SecretBackendCrlConfigOutput) UnifiedCrl() pulumi.BoolOutput {
	return o.ApplyT(func(v *SecretBackendCrlConfig) pulumi.BoolOutput { return v.UnifiedCrl }).(pulumi.BoolOutput)
}

// Enables serving the unified CRL and OCSP on the existing, previously
// cluster-local paths. **Vault 1.13+**
func (o SecretBackendCrlConfigOutput) UnifiedCrlOnExistingPaths() pulumi.BoolOutput {
	return o.ApplyT(func(v *SecretBackendCrlConfig) pulumi.BoolOutput { return v.UnifiedCrlOnExistingPaths }).(pulumi.BoolOutput)
}

type SecretBackendCrlConfigArrayOutput struct{ *pulumi.OutputState }

func (SecretBackendCrlConfigArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SecretBackendCrlConfig)(nil)).Elem()
}

func (o SecretBackendCrlConfigArrayOutput) ToSecretBackendCrlConfigArrayOutput() SecretBackendCrlConfigArrayOutput {
	return o
}

func (o SecretBackendCrlConfigArrayOutput) ToSecretBackendCrlConfigArrayOutputWithContext(ctx context.Context) SecretBackendCrlConfigArrayOutput {
	return o
}

func (o SecretBackendCrlConfigArrayOutput) Index(i pulumi.IntInput) SecretBackendCrlConfigOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SecretBackendCrlConfig {
		return vs[0].([]*SecretBackendCrlConfig)[vs[1].(int)]
	}).(SecretBackendCrlConfigOutput)
}

type SecretBackendCrlConfigMapOutput struct{ *pulumi.OutputState }

func (SecretBackendCrlConfigMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SecretBackendCrlConfig)(nil)).Elem()
}

func (o SecretBackendCrlConfigMapOutput) ToSecretBackendCrlConfigMapOutput() SecretBackendCrlConfigMapOutput {
	return o
}

func (o SecretBackendCrlConfigMapOutput) ToSecretBackendCrlConfigMapOutputWithContext(ctx context.Context) SecretBackendCrlConfigMapOutput {
	return o
}

func (o SecretBackendCrlConfigMapOutput) MapIndex(k pulumi.StringInput) SecretBackendCrlConfigOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SecretBackendCrlConfig {
		return vs[0].(map[string]*SecretBackendCrlConfig)[vs[1].(string)]
	}).(SecretBackendCrlConfigOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SecretBackendCrlConfigInput)(nil)).Elem(), &SecretBackendCrlConfig{})
	pulumi.RegisterInputType(reflect.TypeOf((*SecretBackendCrlConfigArrayInput)(nil)).Elem(), SecretBackendCrlConfigArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SecretBackendCrlConfigMapInput)(nil)).Elem(), SecretBackendCrlConfigMap{})
	pulumi.RegisterOutputType(SecretBackendCrlConfigOutput{})
	pulumi.RegisterOutputType(SecretBackendCrlConfigArrayOutput{})
	pulumi.RegisterOutputType(SecretBackendCrlConfigMapOutput{})
}
