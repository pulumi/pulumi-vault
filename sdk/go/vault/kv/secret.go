// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package kv

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-vault/sdk/v6/go/vault/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Writes a KV-V1 secret to a given path in Vault.
//
// For more information on Vault's KV-V1 secret backend
// [see here](https://www.vaultproject.io/docs/secrets/kv/kv-v1).
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"encoding/json"
//	"fmt"
//
//	"github.com/pulumi/pulumi-vault/sdk/v6/go/vault"
//	"github.com/pulumi/pulumi-vault/sdk/v6/go/vault/kv"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			kvv1, err := vault.NewMount(ctx, "kvv1", &vault.MountArgs{
//				Path: pulumi.String("kvv1"),
//				Type: pulumi.String("kv"),
//				Options: pulumi.StringMap{
//					"version": pulumi.String("1"),
//				},
//				Description: pulumi.String("KV Version 1 secret engine mount"),
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
//			_, err = kv.NewSecret(ctx, "secret", &kv.SecretArgs{
//				Path: kvv1.Path.ApplyT(func(path string) (string, error) {
//					return fmt.Sprintf("%v/secret", path), nil
//				}).(pulumi.StringOutput),
//				DataJson: pulumi.String(json0),
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
// ## Import
//
// KV-V1 secrets can be imported using the `path`, e.g.
//
// ```sh
// $ pulumi import vault:kv/secret:Secret secret kvv1/secret
// ```
type Secret struct {
	pulumi.CustomResourceState

	// A mapping whose keys are the top-level data keys returned from
	// Vault and whose values are the corresponding values. This map can only
	// represent string data, so any non-string values returned from Vault are
	// serialized as JSON.
	Data pulumi.StringMapOutput `pulumi:"data"`
	// JSON-encoded string that will be
	// written as the secret data at the given path.
	DataJson pulumi.StringOutput `pulumi:"dataJson"`
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
	// *Available only for Vault Enterprise*.
	Namespace pulumi.StringPtrOutput `pulumi:"namespace"`
	// Full path of the KV-V1 secret.
	Path pulumi.StringOutput `pulumi:"path"`
}

// NewSecret registers a new resource with the given unique name, arguments, and options.
func NewSecret(ctx *pulumi.Context,
	name string, args *SecretArgs, opts ...pulumi.ResourceOption) (*Secret, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DataJson == nil {
		return nil, errors.New("invalid value for required argument 'DataJson'")
	}
	if args.Path == nil {
		return nil, errors.New("invalid value for required argument 'Path'")
	}
	if args.DataJson != nil {
		args.DataJson = pulumi.ToSecret(args.DataJson).(pulumi.StringInput)
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"data",
		"dataJson",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Secret
	err := ctx.RegisterResource("vault:kv/secret:Secret", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSecret gets an existing Secret resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSecret(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SecretState, opts ...pulumi.ResourceOption) (*Secret, error) {
	var resource Secret
	err := ctx.ReadResource("vault:kv/secret:Secret", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Secret resources.
type secretState struct {
	// A mapping whose keys are the top-level data keys returned from
	// Vault and whose values are the corresponding values. This map can only
	// represent string data, so any non-string values returned from Vault are
	// serialized as JSON.
	Data map[string]string `pulumi:"data"`
	// JSON-encoded string that will be
	// written as the secret data at the given path.
	DataJson *string `pulumi:"dataJson"`
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
	// *Available only for Vault Enterprise*.
	Namespace *string `pulumi:"namespace"`
	// Full path of the KV-V1 secret.
	Path *string `pulumi:"path"`
}

type SecretState struct {
	// A mapping whose keys are the top-level data keys returned from
	// Vault and whose values are the corresponding values. This map can only
	// represent string data, so any non-string values returned from Vault are
	// serialized as JSON.
	Data pulumi.StringMapInput
	// JSON-encoded string that will be
	// written as the secret data at the given path.
	DataJson pulumi.StringPtrInput
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
	// *Available only for Vault Enterprise*.
	Namespace pulumi.StringPtrInput
	// Full path of the KV-V1 secret.
	Path pulumi.StringPtrInput
}

func (SecretState) ElementType() reflect.Type {
	return reflect.TypeOf((*secretState)(nil)).Elem()
}

type secretArgs struct {
	// JSON-encoded string that will be
	// written as the secret data at the given path.
	DataJson string `pulumi:"dataJson"`
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
	// *Available only for Vault Enterprise*.
	Namespace *string `pulumi:"namespace"`
	// Full path of the KV-V1 secret.
	Path string `pulumi:"path"`
}

// The set of arguments for constructing a Secret resource.
type SecretArgs struct {
	// JSON-encoded string that will be
	// written as the secret data at the given path.
	DataJson pulumi.StringInput
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
	// *Available only for Vault Enterprise*.
	Namespace pulumi.StringPtrInput
	// Full path of the KV-V1 secret.
	Path pulumi.StringInput
}

func (SecretArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*secretArgs)(nil)).Elem()
}

type SecretInput interface {
	pulumi.Input

	ToSecretOutput() SecretOutput
	ToSecretOutputWithContext(ctx context.Context) SecretOutput
}

func (*Secret) ElementType() reflect.Type {
	return reflect.TypeOf((**Secret)(nil)).Elem()
}

func (i *Secret) ToSecretOutput() SecretOutput {
	return i.ToSecretOutputWithContext(context.Background())
}

func (i *Secret) ToSecretOutputWithContext(ctx context.Context) SecretOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecretOutput)
}

// SecretArrayInput is an input type that accepts SecretArray and SecretArrayOutput values.
// You can construct a concrete instance of `SecretArrayInput` via:
//
//	SecretArray{ SecretArgs{...} }
type SecretArrayInput interface {
	pulumi.Input

	ToSecretArrayOutput() SecretArrayOutput
	ToSecretArrayOutputWithContext(context.Context) SecretArrayOutput
}

type SecretArray []SecretInput

func (SecretArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Secret)(nil)).Elem()
}

func (i SecretArray) ToSecretArrayOutput() SecretArrayOutput {
	return i.ToSecretArrayOutputWithContext(context.Background())
}

func (i SecretArray) ToSecretArrayOutputWithContext(ctx context.Context) SecretArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecretArrayOutput)
}

// SecretMapInput is an input type that accepts SecretMap and SecretMapOutput values.
// You can construct a concrete instance of `SecretMapInput` via:
//
//	SecretMap{ "key": SecretArgs{...} }
type SecretMapInput interface {
	pulumi.Input

	ToSecretMapOutput() SecretMapOutput
	ToSecretMapOutputWithContext(context.Context) SecretMapOutput
}

type SecretMap map[string]SecretInput

func (SecretMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Secret)(nil)).Elem()
}

func (i SecretMap) ToSecretMapOutput() SecretMapOutput {
	return i.ToSecretMapOutputWithContext(context.Background())
}

func (i SecretMap) ToSecretMapOutputWithContext(ctx context.Context) SecretMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecretMapOutput)
}

type SecretOutput struct{ *pulumi.OutputState }

func (SecretOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Secret)(nil)).Elem()
}

func (o SecretOutput) ToSecretOutput() SecretOutput {
	return o
}

func (o SecretOutput) ToSecretOutputWithContext(ctx context.Context) SecretOutput {
	return o
}

// A mapping whose keys are the top-level data keys returned from
// Vault and whose values are the corresponding values. This map can only
// represent string data, so any non-string values returned from Vault are
// serialized as JSON.
func (o SecretOutput) Data() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Secret) pulumi.StringMapOutput { return v.Data }).(pulumi.StringMapOutput)
}

// JSON-encoded string that will be
// written as the secret data at the given path.
func (o SecretOutput) DataJson() pulumi.StringOutput {
	return o.ApplyT(func(v *Secret) pulumi.StringOutput { return v.DataJson }).(pulumi.StringOutput)
}

// The namespace to provision the resource in.
// The value should not contain leading or trailing forward slashes.
// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
// *Available only for Vault Enterprise*.
func (o SecretOutput) Namespace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Secret) pulumi.StringPtrOutput { return v.Namespace }).(pulumi.StringPtrOutput)
}

// Full path of the KV-V1 secret.
func (o SecretOutput) Path() pulumi.StringOutput {
	return o.ApplyT(func(v *Secret) pulumi.StringOutput { return v.Path }).(pulumi.StringOutput)
}

type SecretArrayOutput struct{ *pulumi.OutputState }

func (SecretArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Secret)(nil)).Elem()
}

func (o SecretArrayOutput) ToSecretArrayOutput() SecretArrayOutput {
	return o
}

func (o SecretArrayOutput) ToSecretArrayOutputWithContext(ctx context.Context) SecretArrayOutput {
	return o
}

func (o SecretArrayOutput) Index(i pulumi.IntInput) SecretOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Secret {
		return vs[0].([]*Secret)[vs[1].(int)]
	}).(SecretOutput)
}

type SecretMapOutput struct{ *pulumi.OutputState }

func (SecretMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Secret)(nil)).Elem()
}

func (o SecretMapOutput) ToSecretMapOutput() SecretMapOutput {
	return o
}

func (o SecretMapOutput) ToSecretMapOutputWithContext(ctx context.Context) SecretMapOutput {
	return o
}

func (o SecretMapOutput) MapIndex(k pulumi.StringInput) SecretOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Secret {
		return vs[0].(map[string]*Secret)[vs[1].(string)]
	}).(SecretOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SecretInput)(nil)).Elem(), &Secret{})
	pulumi.RegisterInputType(reflect.TypeOf((*SecretArrayInput)(nil)).Elem(), SecretArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SecretMapInput)(nil)).Elem(), SecretMap{})
	pulumi.RegisterOutputType(SecretOutput{})
	pulumi.RegisterOutputType(SecretArrayOutput{})
	pulumi.RegisterOutputType(SecretMapOutput{})
}
