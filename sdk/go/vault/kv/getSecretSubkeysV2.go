// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package kv

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"encoding/json"
//
//	"github.com/pulumi/pulumi-vault/sdk/v5/go/vault"
//	"github.com/pulumi/pulumi-vault/sdk/v5/go/vault/kv"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			kvv2, err := vault.NewMount(ctx, "kvv2", &vault.MountArgs{
//				Path: pulumi.String("kvv2"),
//				Type: pulumi.String("kv"),
//				Options: pulumi.AnyMap{
//					"version": pulumi.Any("2"),
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
//			awsSecret, err := kv.NewSecretV2(ctx, "awsSecret", &kv.SecretV2Args{
//				Mount:    kvv2.Path,
//				DataJson: pulumi.String(json0),
//			})
//			if err != nil {
//				return err
//			}
//			_ = kv.GetSecretSubkeysV2Output(ctx, kv.GetSecretSubkeysV2OutputArgs{
//				Mount: kvv2.Path,
//				Name:  awsSecret.Name,
//			}, nil)
//			return nil
//		})
//	}
//
// ```
// ## Required Vault Capabilities
//
// Use of this resource requires the `read` capability on the given path.
func GetSecretSubkeysV2(ctx *pulumi.Context, args *GetSecretSubkeysV2Args, opts ...pulumi.InvokeOption) (*GetSecretSubkeysV2Result, error) {
	var rv GetSecretSubkeysV2Result
	err := ctx.Invoke("vault:kv/getSecretSubkeysV2:getSecretSubkeysV2", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSecretSubkeysV2.
type GetSecretSubkeysV2Args struct {
	// Specifies the deepest nesting level to provide in the output.
	// If non-zero, keys that reside at the specified depth value will be
	// artificially treated as leaves and will thus be `null` even if further
	// underlying sub-keys exist.
	Depth *int `pulumi:"depth"`
	// Path where KV-V2 engine is mounted.
	Mount string `pulumi:"mount"`
	// Full name of the secret. For a nested secret
	// the name is the nested path excluding the mount and data
	// prefix. For example, for a secret at `kvv2/data/foo/bar/baz`
	// the name is `foo/bar/baz`.
	Name string `pulumi:"name"`
	// The namespace of the target resource.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
	// *Available only for Vault Enterprise*.
	Namespace *string `pulumi:"namespace"`
	// Specifies the version to return. If not
	// set the latest version is returned.
	Version *int `pulumi:"version"`
}

// A collection of values returned by getSecretSubkeysV2.
type GetSecretSubkeysV2Result struct {
	// Subkeys for the KV-V2 secret stored as a serialized map of strings.
	Data map[string]interface{} `pulumi:"data"`
	// Subkeys for the KV-V2 secret read from Vault.
	DataJson string `pulumi:"dataJson"`
	Depth    *int   `pulumi:"depth"`
	// The provider-assigned unique ID for this managed resource.
	Id        string  `pulumi:"id"`
	Mount     string  `pulumi:"mount"`
	Name      string  `pulumi:"name"`
	Namespace *string `pulumi:"namespace"`
	// Full path where the KV-V2 secrets are listed.
	Path    string `pulumi:"path"`
	Version *int   `pulumi:"version"`
}

func GetSecretSubkeysV2Output(ctx *pulumi.Context, args GetSecretSubkeysV2OutputArgs, opts ...pulumi.InvokeOption) GetSecretSubkeysV2ResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetSecretSubkeysV2Result, error) {
			args := v.(GetSecretSubkeysV2Args)
			r, err := GetSecretSubkeysV2(ctx, &args, opts...)
			var s GetSecretSubkeysV2Result
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetSecretSubkeysV2ResultOutput)
}

// A collection of arguments for invoking getSecretSubkeysV2.
type GetSecretSubkeysV2OutputArgs struct {
	// Specifies the deepest nesting level to provide in the output.
	// If non-zero, keys that reside at the specified depth value will be
	// artificially treated as leaves and will thus be `null` even if further
	// underlying sub-keys exist.
	Depth pulumi.IntPtrInput `pulumi:"depth"`
	// Path where KV-V2 engine is mounted.
	Mount pulumi.StringInput `pulumi:"mount"`
	// Full name of the secret. For a nested secret
	// the name is the nested path excluding the mount and data
	// prefix. For example, for a secret at `kvv2/data/foo/bar/baz`
	// the name is `foo/bar/baz`.
	Name pulumi.StringInput `pulumi:"name"`
	// The namespace of the target resource.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
	// *Available only for Vault Enterprise*.
	Namespace pulumi.StringPtrInput `pulumi:"namespace"`
	// Specifies the version to return. If not
	// set the latest version is returned.
	Version pulumi.IntPtrInput `pulumi:"version"`
}

func (GetSecretSubkeysV2OutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSecretSubkeysV2Args)(nil)).Elem()
}

// A collection of values returned by getSecretSubkeysV2.
type GetSecretSubkeysV2ResultOutput struct{ *pulumi.OutputState }

func (GetSecretSubkeysV2ResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSecretSubkeysV2Result)(nil)).Elem()
}

func (o GetSecretSubkeysV2ResultOutput) ToGetSecretSubkeysV2ResultOutput() GetSecretSubkeysV2ResultOutput {
	return o
}

func (o GetSecretSubkeysV2ResultOutput) ToGetSecretSubkeysV2ResultOutputWithContext(ctx context.Context) GetSecretSubkeysV2ResultOutput {
	return o
}

// Subkeys for the KV-V2 secret stored as a serialized map of strings.
func (o GetSecretSubkeysV2ResultOutput) Data() pulumi.MapOutput {
	return o.ApplyT(func(v GetSecretSubkeysV2Result) map[string]interface{} { return v.Data }).(pulumi.MapOutput)
}

// Subkeys for the KV-V2 secret read from Vault.
func (o GetSecretSubkeysV2ResultOutput) DataJson() pulumi.StringOutput {
	return o.ApplyT(func(v GetSecretSubkeysV2Result) string { return v.DataJson }).(pulumi.StringOutput)
}

func (o GetSecretSubkeysV2ResultOutput) Depth() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GetSecretSubkeysV2Result) *int { return v.Depth }).(pulumi.IntPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetSecretSubkeysV2ResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetSecretSubkeysV2Result) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetSecretSubkeysV2ResultOutput) Mount() pulumi.StringOutput {
	return o.ApplyT(func(v GetSecretSubkeysV2Result) string { return v.Mount }).(pulumi.StringOutput)
}

func (o GetSecretSubkeysV2ResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GetSecretSubkeysV2Result) string { return v.Name }).(pulumi.StringOutput)
}

func (o GetSecretSubkeysV2ResultOutput) Namespace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetSecretSubkeysV2Result) *string { return v.Namespace }).(pulumi.StringPtrOutput)
}

// Full path where the KV-V2 secrets are listed.
func (o GetSecretSubkeysV2ResultOutput) Path() pulumi.StringOutput {
	return o.ApplyT(func(v GetSecretSubkeysV2Result) string { return v.Path }).(pulumi.StringOutput)
}

func (o GetSecretSubkeysV2ResultOutput) Version() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GetSecretSubkeysV2Result) *int { return v.Version }).(pulumi.IntPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetSecretSubkeysV2ResultOutput{})
}