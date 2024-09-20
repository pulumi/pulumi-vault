// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package kv

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
//			})
//			if err != nil {
//				return err
//			}
//			json0 := string(tmpJSON0)
//			_, err = kv.NewSecretV2(ctx, "aws_secret", &kv.SecretV2Args{
//				Mount:    kvv2.Path,
//				Name:     pulumi.String("aws_secret"),
//				DataJson: pulumi.String(json0),
//			})
//			if err != nil {
//				return err
//			}
//			tmpJSON1, err := json.Marshal(map[string]interface{}{
//				"foo": "bar",
//			})
//			if err != nil {
//				return err
//			}
//			json1 := string(tmpJSON1)
//			azureSecret, err := kv.NewSecretV2(ctx, "azure_secret", &kv.SecretV2Args{
//				Mount:    kvv2.Path,
//				Name:     pulumi.String("azure_secret"),
//				DataJson: pulumi.String(json1),
//			})
//			if err != nil {
//				return err
//			}
//			tmpJSON2, err := json.Marshal(map[string]interface{}{
//				"password": "test",
//			})
//			if err != nil {
//				return err
//			}
//			json2 := string(tmpJSON2)
//			_, err = kv.NewSecretV2(ctx, "nested_secret", &kv.SecretV2Args{
//				Mount: kvv2.Path,
//				Name: azureSecret.Name.ApplyT(func(name string) (string, error) {
//					return fmt.Sprintf("%v/dev", name), nil
//				}).(pulumi.StringOutput),
//				DataJson: pulumi.String(json2),
//			})
//			if err != nil {
//				return err
//			}
//			_ = kv.GetSecretsListV2Output(ctx, kv.GetSecretsListV2OutputArgs{
//				Mount: kvv2.Path,
//			}, nil)
//			_ = kvv2.Path.ApplyT(func(path string) (kv.GetSecretsListV2Result, error) {
//				return kv.GetSecretsListV2Result(interface{}(kv.GetSecretsListV2Output(ctx, kv.GetSecretsListV2OutputArgs{
//					Mount: path,
//					Name:  test2.Name,
//				}, nil))), nil
//			}).(kv.GetSecretsListV2ResultOutput)
//			return nil
//		})
//	}
//
// ```
//
// ## Required Vault Capabilities
//
// Use of this resource requires the `read` capability on the given path.
func GetSecretsListV2(ctx *pulumi.Context, args *GetSecretsListV2Args, opts ...pulumi.InvokeOption) (*GetSecretsListV2Result, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetSecretsListV2Result
	err := ctx.Invoke("vault:kv/getSecretsListV2:getSecretsListV2", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSecretsListV2.
type GetSecretsListV2Args struct {
	// Path where KV-V2 engine is mounted.
	Mount string `pulumi:"mount"`
	// Full name of the secret. For a nested secret
	// the name is the nested path excluding the mount and data
	// prefix. For example, for a secret at `kvv2/data/foo/bar/baz`
	// the name is `foo/bar/baz`.
	Name *string `pulumi:"name"`
	// The namespace of the target resource.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
	// *Available only for Vault Enterprise*.
	Namespace *string `pulumi:"namespace"`
}

// A collection of values returned by getSecretsListV2.
type GetSecretsListV2Result struct {
	// The provider-assigned unique ID for this managed resource.
	Id    string  `pulumi:"id"`
	Mount string  `pulumi:"mount"`
	Name  *string `pulumi:"name"`
	// List of all secret names listed under the given path.
	Names     []string `pulumi:"names"`
	Namespace *string  `pulumi:"namespace"`
	// Full path where the KV-V2 secrets are listed.
	Path string `pulumi:"path"`
}

func GetSecretsListV2Output(ctx *pulumi.Context, args GetSecretsListV2OutputArgs, opts ...pulumi.InvokeOption) GetSecretsListV2ResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetSecretsListV2ResultOutput, error) {
			args := v.(GetSecretsListV2Args)
			opts = internal.PkgInvokeDefaultOpts(opts)
			var rv GetSecretsListV2Result
			secret, err := ctx.InvokePackageRaw("vault:kv/getSecretsListV2:getSecretsListV2", args, &rv, "", opts...)
			if err != nil {
				return GetSecretsListV2ResultOutput{}, err
			}

			output := pulumi.ToOutput(rv).(GetSecretsListV2ResultOutput)
			if secret {
				return pulumi.ToSecret(output).(GetSecretsListV2ResultOutput), nil
			}
			return output, nil
		}).(GetSecretsListV2ResultOutput)
}

// A collection of arguments for invoking getSecretsListV2.
type GetSecretsListV2OutputArgs struct {
	// Path where KV-V2 engine is mounted.
	Mount pulumi.StringInput `pulumi:"mount"`
	// Full name of the secret. For a nested secret
	// the name is the nested path excluding the mount and data
	// prefix. For example, for a secret at `kvv2/data/foo/bar/baz`
	// the name is `foo/bar/baz`.
	Name pulumi.StringPtrInput `pulumi:"name"`
	// The namespace of the target resource.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
	// *Available only for Vault Enterprise*.
	Namespace pulumi.StringPtrInput `pulumi:"namespace"`
}

func (GetSecretsListV2OutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSecretsListV2Args)(nil)).Elem()
}

// A collection of values returned by getSecretsListV2.
type GetSecretsListV2ResultOutput struct{ *pulumi.OutputState }

func (GetSecretsListV2ResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSecretsListV2Result)(nil)).Elem()
}

func (o GetSecretsListV2ResultOutput) ToGetSecretsListV2ResultOutput() GetSecretsListV2ResultOutput {
	return o
}

func (o GetSecretsListV2ResultOutput) ToGetSecretsListV2ResultOutputWithContext(ctx context.Context) GetSecretsListV2ResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o GetSecretsListV2ResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetSecretsListV2Result) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetSecretsListV2ResultOutput) Mount() pulumi.StringOutput {
	return o.ApplyT(func(v GetSecretsListV2Result) string { return v.Mount }).(pulumi.StringOutput)
}

func (o GetSecretsListV2ResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetSecretsListV2Result) *string { return v.Name }).(pulumi.StringPtrOutput)
}

// List of all secret names listed under the given path.
func (o GetSecretsListV2ResultOutput) Names() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetSecretsListV2Result) []string { return v.Names }).(pulumi.StringArrayOutput)
}

func (o GetSecretsListV2ResultOutput) Namespace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetSecretsListV2Result) *string { return v.Namespace }).(pulumi.StringPtrOutput)
}

// Full path where the KV-V2 secrets are listed.
func (o GetSecretsListV2ResultOutput) Path() pulumi.StringOutput {
	return o.ApplyT(func(v GetSecretsListV2Result) string { return v.Path }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetSecretsListV2ResultOutput{})
}
