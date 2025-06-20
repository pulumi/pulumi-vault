// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package kv

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-vault/sdk/v7/go/vault/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Writes a KV-V2 secret to a given path in Vault.
//
// For more information on Vault's KV-V2 secret backend
// [see here](https://www.vaultproject.io/docs/secrets/kv/kv-v2).
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"encoding/json"
//
//	"github.com/pulumi/pulumi-vault/sdk/v7/go/vault"
//	"github.com/pulumi/pulumi-vault/sdk/v7/go/vault/kv"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			kvv2, err := vault.NewMount(ctx, "kvv2", &vault.MountArgs{
//				Path: pulumi.String("kvv2"),
//				Type: pulumi.String("kv"),
//				Options: pulumi.StringMap{
//					"version": pulumi.String("2"),
//				},
//				Description: pulumi.String("KV Version 2 secret engine mount"),
//			})
//			if err != nil {
//				return err
//			}
//			tmpJSON0, err := json.Marshal(map[string]interface{}{
//				"zip": "zap",
//				"foo": "bar",
//			})
//			if err != nil {
//				return err
//			}
//			json0 := string(tmpJSON0)
//			_, err = kv.NewSecretV2(ctx, "example", &kv.SecretV2Args{
//				Mount:             kvv2.Path,
//				Name:              pulumi.String("secret"),
//				Cas:               pulumi.Int(1),
//				DeleteAllVersions: pulumi.Bool(true),
//				DataJson:          pulumi.String(json0),
//				CustomMetadata: &kv.SecretV2CustomMetadataArgs{
//					MaxVersions: pulumi.Int(5),
//					Data: pulumi.StringMap{
//						"foo": pulumi.String("vault@example.com"),
//						"bar": pulumi.String("12345"),
//					},
//				},
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
// ## Required Vault Capabilities
//
// Use of this resource requires the `create` or `update` capability
// (depending on whether the resource already exists) on the given path,
// the `delete` capability if the resource is removed from configuration,
// and the `read` capability for drift detection (by default).
//
// ### Custom Metadata Configuration Options
//
// * `maxVersions` - (Optional) The number of versions to keep per key.
//
// * `casRequired` - (Optional) If true, all keys will require the cas
// parameter to be set on all write requests.
//
// * `deleteVersionAfter` - (Optional) If set, specifies the length of time before
// a version is deleted. Accepts duration in integer seconds.
//
// * `data` - (Optional) A string to string map describing the secret.
//
// ## Ephemeral Attributes Reference
//
// The following write-only attributes are supported:
//
//   - `dataJsonWo` - (Optional) JSON-encoded secret data to write to Vault. Can be updated.
//     **Note**: This property is write-only and will not be read from the API.
//
// ## Import
//
// KV-V2 secrets can be imported using the `path`, e.g.
//
// ```sh
// $ pulumi import vault:kv/secretV2:SecretV2 example kvv2/data/secret
// ```
type SecretV2 struct {
	pulumi.CustomResourceState

	// This flag is required if `casRequired` is set to true
	// on either the secret or the engine's config. In order for a
	// write operation to be successful, cas must be set to the current version
	// of the secret.
	Cas pulumi.IntPtrOutput `pulumi:"cas"`
	// A nested block that allows configuring metadata for the
	// KV secret. Refer to the
	// Configuration Options for more info.
	CustomMetadata SecretV2CustomMetadataOutput `pulumi:"customMetadata"`
	// **Deprecated. Please use new ephemeral resource `kv.SecretV2` to read back
	// secret data from Vault**. A mapping whose keys are the top-level data keys returned from
	// Vault and whose values are the corresponding values. This map can only represent string data,
	// so any non-string values returned from Vault are serialized as JSON.
	//
	// Deprecated: Deprecated. Will no longer be set on a read.
	Data pulumi.StringMapOutput `pulumi:"data"`
	// JSON-encoded string that will be
	// written as the secret data at the given path.
	DataJson pulumi.StringPtrOutput `pulumi:"dataJson"`
	// The version of the `dataJsonWo`. For more info see updating write-only attributes.
	DataJsonWoVersion pulumi.IntPtrOutput `pulumi:"dataJsonWoVersion"`
	// If set to true, permanently deletes all
	// versions for the specified key.
	DeleteAllVersions pulumi.BoolPtrOutput `pulumi:"deleteAllVersions"`
	// If set to true, disables reading secret from Vault;
	// note: drift won't be detected.
	DisableRead pulumi.BoolPtrOutput `pulumi:"disableRead"`
	// Metadata associated with this secret read from Vault.
	Metadata pulumi.StringMapOutput `pulumi:"metadata"`
	// Path where KV-V2 engine is mounted.
	Mount pulumi.StringOutput `pulumi:"mount"`
	// Full name of the secret. For a nested secret
	// the name is the nested path excluding the mount and data
	// prefix. For example, for a secret at `kvv2/data/foo/bar/baz`
	// the name is `foo/bar/baz`.
	Name pulumi.StringOutput `pulumi:"name"`
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
	// *Available only for Vault Enterprise*.
	Namespace pulumi.StringPtrOutput `pulumi:"namespace"`
	// An object that holds option settings.
	Options pulumi.StringMapOutput `pulumi:"options"`
	// Full path where the KV-V2 secret will be written.
	Path pulumi.StringOutput `pulumi:"path"`
}

// NewSecretV2 registers a new resource with the given unique name, arguments, and options.
func NewSecretV2(ctx *pulumi.Context,
	name string, args *SecretV2Args, opts ...pulumi.ResourceOption) (*SecretV2, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Mount == nil {
		return nil, errors.New("invalid value for required argument 'Mount'")
	}
	if args.DataJson != nil {
		args.DataJson = pulumi.ToSecret(args.DataJson).(pulumi.StringPtrInput)
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"data",
		"dataJson",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource SecretV2
	err := ctx.RegisterResource("vault:kv/secretV2:SecretV2", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSecretV2 gets an existing SecretV2 resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSecretV2(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SecretV2State, opts ...pulumi.ResourceOption) (*SecretV2, error) {
	var resource SecretV2
	err := ctx.ReadResource("vault:kv/secretV2:SecretV2", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SecretV2 resources.
type secretV2State struct {
	// This flag is required if `casRequired` is set to true
	// on either the secret or the engine's config. In order for a
	// write operation to be successful, cas must be set to the current version
	// of the secret.
	Cas *int `pulumi:"cas"`
	// A nested block that allows configuring metadata for the
	// KV secret. Refer to the
	// Configuration Options for more info.
	CustomMetadata *SecretV2CustomMetadata `pulumi:"customMetadata"`
	// **Deprecated. Please use new ephemeral resource `kv.SecretV2` to read back
	// secret data from Vault**. A mapping whose keys are the top-level data keys returned from
	// Vault and whose values are the corresponding values. This map can only represent string data,
	// so any non-string values returned from Vault are serialized as JSON.
	//
	// Deprecated: Deprecated. Will no longer be set on a read.
	Data map[string]string `pulumi:"data"`
	// JSON-encoded string that will be
	// written as the secret data at the given path.
	DataJson *string `pulumi:"dataJson"`
	// The version of the `dataJsonWo`. For more info see updating write-only attributes.
	DataJsonWoVersion *int `pulumi:"dataJsonWoVersion"`
	// If set to true, permanently deletes all
	// versions for the specified key.
	DeleteAllVersions *bool `pulumi:"deleteAllVersions"`
	// If set to true, disables reading secret from Vault;
	// note: drift won't be detected.
	DisableRead *bool `pulumi:"disableRead"`
	// Metadata associated with this secret read from Vault.
	Metadata map[string]string `pulumi:"metadata"`
	// Path where KV-V2 engine is mounted.
	Mount *string `pulumi:"mount"`
	// Full name of the secret. For a nested secret
	// the name is the nested path excluding the mount and data
	// prefix. For example, for a secret at `kvv2/data/foo/bar/baz`
	// the name is `foo/bar/baz`.
	Name *string `pulumi:"name"`
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
	// *Available only for Vault Enterprise*.
	Namespace *string `pulumi:"namespace"`
	// An object that holds option settings.
	Options map[string]string `pulumi:"options"`
	// Full path where the KV-V2 secret will be written.
	Path *string `pulumi:"path"`
}

type SecretV2State struct {
	// This flag is required if `casRequired` is set to true
	// on either the secret or the engine's config. In order for a
	// write operation to be successful, cas must be set to the current version
	// of the secret.
	Cas pulumi.IntPtrInput
	// A nested block that allows configuring metadata for the
	// KV secret. Refer to the
	// Configuration Options for more info.
	CustomMetadata SecretV2CustomMetadataPtrInput
	// **Deprecated. Please use new ephemeral resource `kv.SecretV2` to read back
	// secret data from Vault**. A mapping whose keys are the top-level data keys returned from
	// Vault and whose values are the corresponding values. This map can only represent string data,
	// so any non-string values returned from Vault are serialized as JSON.
	//
	// Deprecated: Deprecated. Will no longer be set on a read.
	Data pulumi.StringMapInput
	// JSON-encoded string that will be
	// written as the secret data at the given path.
	DataJson pulumi.StringPtrInput
	// The version of the `dataJsonWo`. For more info see updating write-only attributes.
	DataJsonWoVersion pulumi.IntPtrInput
	// If set to true, permanently deletes all
	// versions for the specified key.
	DeleteAllVersions pulumi.BoolPtrInput
	// If set to true, disables reading secret from Vault;
	// note: drift won't be detected.
	DisableRead pulumi.BoolPtrInput
	// Metadata associated with this secret read from Vault.
	Metadata pulumi.StringMapInput
	// Path where KV-V2 engine is mounted.
	Mount pulumi.StringPtrInput
	// Full name of the secret. For a nested secret
	// the name is the nested path excluding the mount and data
	// prefix. For example, for a secret at `kvv2/data/foo/bar/baz`
	// the name is `foo/bar/baz`.
	Name pulumi.StringPtrInput
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
	// *Available only for Vault Enterprise*.
	Namespace pulumi.StringPtrInput
	// An object that holds option settings.
	Options pulumi.StringMapInput
	// Full path where the KV-V2 secret will be written.
	Path pulumi.StringPtrInput
}

func (SecretV2State) ElementType() reflect.Type {
	return reflect.TypeOf((*secretV2State)(nil)).Elem()
}

type secretV2Args struct {
	// This flag is required if `casRequired` is set to true
	// on either the secret or the engine's config. In order for a
	// write operation to be successful, cas must be set to the current version
	// of the secret.
	Cas *int `pulumi:"cas"`
	// A nested block that allows configuring metadata for the
	// KV secret. Refer to the
	// Configuration Options for more info.
	CustomMetadata *SecretV2CustomMetadata `pulumi:"customMetadata"`
	// JSON-encoded string that will be
	// written as the secret data at the given path.
	DataJson *string `pulumi:"dataJson"`
	// The version of the `dataJsonWo`. For more info see updating write-only attributes.
	DataJsonWoVersion *int `pulumi:"dataJsonWoVersion"`
	// If set to true, permanently deletes all
	// versions for the specified key.
	DeleteAllVersions *bool `pulumi:"deleteAllVersions"`
	// If set to true, disables reading secret from Vault;
	// note: drift won't be detected.
	DisableRead *bool `pulumi:"disableRead"`
	// Path where KV-V2 engine is mounted.
	Mount string `pulumi:"mount"`
	// Full name of the secret. For a nested secret
	// the name is the nested path excluding the mount and data
	// prefix. For example, for a secret at `kvv2/data/foo/bar/baz`
	// the name is `foo/bar/baz`.
	Name *string `pulumi:"name"`
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
	// *Available only for Vault Enterprise*.
	Namespace *string `pulumi:"namespace"`
	// An object that holds option settings.
	Options map[string]string `pulumi:"options"`
}

// The set of arguments for constructing a SecretV2 resource.
type SecretV2Args struct {
	// This flag is required if `casRequired` is set to true
	// on either the secret or the engine's config. In order for a
	// write operation to be successful, cas must be set to the current version
	// of the secret.
	Cas pulumi.IntPtrInput
	// A nested block that allows configuring metadata for the
	// KV secret. Refer to the
	// Configuration Options for more info.
	CustomMetadata SecretV2CustomMetadataPtrInput
	// JSON-encoded string that will be
	// written as the secret data at the given path.
	DataJson pulumi.StringPtrInput
	// The version of the `dataJsonWo`. For more info see updating write-only attributes.
	DataJsonWoVersion pulumi.IntPtrInput
	// If set to true, permanently deletes all
	// versions for the specified key.
	DeleteAllVersions pulumi.BoolPtrInput
	// If set to true, disables reading secret from Vault;
	// note: drift won't be detected.
	DisableRead pulumi.BoolPtrInput
	// Path where KV-V2 engine is mounted.
	Mount pulumi.StringInput
	// Full name of the secret. For a nested secret
	// the name is the nested path excluding the mount and data
	// prefix. For example, for a secret at `kvv2/data/foo/bar/baz`
	// the name is `foo/bar/baz`.
	Name pulumi.StringPtrInput
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
	// *Available only for Vault Enterprise*.
	Namespace pulumi.StringPtrInput
	// An object that holds option settings.
	Options pulumi.StringMapInput
}

func (SecretV2Args) ElementType() reflect.Type {
	return reflect.TypeOf((*secretV2Args)(nil)).Elem()
}

type SecretV2Input interface {
	pulumi.Input

	ToSecretV2Output() SecretV2Output
	ToSecretV2OutputWithContext(ctx context.Context) SecretV2Output
}

func (*SecretV2) ElementType() reflect.Type {
	return reflect.TypeOf((**SecretV2)(nil)).Elem()
}

func (i *SecretV2) ToSecretV2Output() SecretV2Output {
	return i.ToSecretV2OutputWithContext(context.Background())
}

func (i *SecretV2) ToSecretV2OutputWithContext(ctx context.Context) SecretV2Output {
	return pulumi.ToOutputWithContext(ctx, i).(SecretV2Output)
}

// SecretV2ArrayInput is an input type that accepts SecretV2Array and SecretV2ArrayOutput values.
// You can construct a concrete instance of `SecretV2ArrayInput` via:
//
//	SecretV2Array{ SecretV2Args{...} }
type SecretV2ArrayInput interface {
	pulumi.Input

	ToSecretV2ArrayOutput() SecretV2ArrayOutput
	ToSecretV2ArrayOutputWithContext(context.Context) SecretV2ArrayOutput
}

type SecretV2Array []SecretV2Input

func (SecretV2Array) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SecretV2)(nil)).Elem()
}

func (i SecretV2Array) ToSecretV2ArrayOutput() SecretV2ArrayOutput {
	return i.ToSecretV2ArrayOutputWithContext(context.Background())
}

func (i SecretV2Array) ToSecretV2ArrayOutputWithContext(ctx context.Context) SecretV2ArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecretV2ArrayOutput)
}

// SecretV2MapInput is an input type that accepts SecretV2Map and SecretV2MapOutput values.
// You can construct a concrete instance of `SecretV2MapInput` via:
//
//	SecretV2Map{ "key": SecretV2Args{...} }
type SecretV2MapInput interface {
	pulumi.Input

	ToSecretV2MapOutput() SecretV2MapOutput
	ToSecretV2MapOutputWithContext(context.Context) SecretV2MapOutput
}

type SecretV2Map map[string]SecretV2Input

func (SecretV2Map) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SecretV2)(nil)).Elem()
}

func (i SecretV2Map) ToSecretV2MapOutput() SecretV2MapOutput {
	return i.ToSecretV2MapOutputWithContext(context.Background())
}

func (i SecretV2Map) ToSecretV2MapOutputWithContext(ctx context.Context) SecretV2MapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecretV2MapOutput)
}

type SecretV2Output struct{ *pulumi.OutputState }

func (SecretV2Output) ElementType() reflect.Type {
	return reflect.TypeOf((**SecretV2)(nil)).Elem()
}

func (o SecretV2Output) ToSecretV2Output() SecretV2Output {
	return o
}

func (o SecretV2Output) ToSecretV2OutputWithContext(ctx context.Context) SecretV2Output {
	return o
}

// This flag is required if `casRequired` is set to true
// on either the secret or the engine's config. In order for a
// write operation to be successful, cas must be set to the current version
// of the secret.
func (o SecretV2Output) Cas() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SecretV2) pulumi.IntPtrOutput { return v.Cas }).(pulumi.IntPtrOutput)
}

// A nested block that allows configuring metadata for the
// KV secret. Refer to the
// Configuration Options for more info.
func (o SecretV2Output) CustomMetadata() SecretV2CustomMetadataOutput {
	return o.ApplyT(func(v *SecretV2) SecretV2CustomMetadataOutput { return v.CustomMetadata }).(SecretV2CustomMetadataOutput)
}

// **Deprecated. Please use new ephemeral resource `kv.SecretV2` to read back
// secret data from Vault**. A mapping whose keys are the top-level data keys returned from
// Vault and whose values are the corresponding values. This map can only represent string data,
// so any non-string values returned from Vault are serialized as JSON.
//
// Deprecated: Deprecated. Will no longer be set on a read.
func (o SecretV2Output) Data() pulumi.StringMapOutput {
	return o.ApplyT(func(v *SecretV2) pulumi.StringMapOutput { return v.Data }).(pulumi.StringMapOutput)
}

// JSON-encoded string that will be
// written as the secret data at the given path.
func (o SecretV2Output) DataJson() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecretV2) pulumi.StringPtrOutput { return v.DataJson }).(pulumi.StringPtrOutput)
}

// The version of the `dataJsonWo`. For more info see updating write-only attributes.
func (o SecretV2Output) DataJsonWoVersion() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SecretV2) pulumi.IntPtrOutput { return v.DataJsonWoVersion }).(pulumi.IntPtrOutput)
}

// If set to true, permanently deletes all
// versions for the specified key.
func (o SecretV2Output) DeleteAllVersions() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SecretV2) pulumi.BoolPtrOutput { return v.DeleteAllVersions }).(pulumi.BoolPtrOutput)
}

// If set to true, disables reading secret from Vault;
// note: drift won't be detected.
func (o SecretV2Output) DisableRead() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SecretV2) pulumi.BoolPtrOutput { return v.DisableRead }).(pulumi.BoolPtrOutput)
}

// Metadata associated with this secret read from Vault.
func (o SecretV2Output) Metadata() pulumi.StringMapOutput {
	return o.ApplyT(func(v *SecretV2) pulumi.StringMapOutput { return v.Metadata }).(pulumi.StringMapOutput)
}

// Path where KV-V2 engine is mounted.
func (o SecretV2Output) Mount() pulumi.StringOutput {
	return o.ApplyT(func(v *SecretV2) pulumi.StringOutput { return v.Mount }).(pulumi.StringOutput)
}

// Full name of the secret. For a nested secret
// the name is the nested path excluding the mount and data
// prefix. For example, for a secret at `kvv2/data/foo/bar/baz`
// the name is `foo/bar/baz`.
func (o SecretV2Output) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SecretV2) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The namespace to provision the resource in.
// The value should not contain leading or trailing forward slashes.
// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
// *Available only for Vault Enterprise*.
func (o SecretV2Output) Namespace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecretV2) pulumi.StringPtrOutput { return v.Namespace }).(pulumi.StringPtrOutput)
}

// An object that holds option settings.
func (o SecretV2Output) Options() pulumi.StringMapOutput {
	return o.ApplyT(func(v *SecretV2) pulumi.StringMapOutput { return v.Options }).(pulumi.StringMapOutput)
}

// Full path where the KV-V2 secret will be written.
func (o SecretV2Output) Path() pulumi.StringOutput {
	return o.ApplyT(func(v *SecretV2) pulumi.StringOutput { return v.Path }).(pulumi.StringOutput)
}

type SecretV2ArrayOutput struct{ *pulumi.OutputState }

func (SecretV2ArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SecretV2)(nil)).Elem()
}

func (o SecretV2ArrayOutput) ToSecretV2ArrayOutput() SecretV2ArrayOutput {
	return o
}

func (o SecretV2ArrayOutput) ToSecretV2ArrayOutputWithContext(ctx context.Context) SecretV2ArrayOutput {
	return o
}

func (o SecretV2ArrayOutput) Index(i pulumi.IntInput) SecretV2Output {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SecretV2 {
		return vs[0].([]*SecretV2)[vs[1].(int)]
	}).(SecretV2Output)
}

type SecretV2MapOutput struct{ *pulumi.OutputState }

func (SecretV2MapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SecretV2)(nil)).Elem()
}

func (o SecretV2MapOutput) ToSecretV2MapOutput() SecretV2MapOutput {
	return o
}

func (o SecretV2MapOutput) ToSecretV2MapOutputWithContext(ctx context.Context) SecretV2MapOutput {
	return o
}

func (o SecretV2MapOutput) MapIndex(k pulumi.StringInput) SecretV2Output {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SecretV2 {
		return vs[0].(map[string]*SecretV2)[vs[1].(string)]
	}).(SecretV2Output)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SecretV2Input)(nil)).Elem(), &SecretV2{})
	pulumi.RegisterInputType(reflect.TypeOf((*SecretV2ArrayInput)(nil)).Elem(), SecretV2Array{})
	pulumi.RegisterInputType(reflect.TypeOf((*SecretV2MapInput)(nil)).Elem(), SecretV2Map{})
	pulumi.RegisterOutputType(SecretV2Output{})
	pulumi.RegisterOutputType(SecretV2ArrayOutput{})
	pulumi.RegisterOutputType(SecretV2MapOutput{})
}
