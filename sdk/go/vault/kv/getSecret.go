// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package kv

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-vault/sdk/v5/go/vault/internal"
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
//	"github.com/pulumi/pulumi-vault/sdk/v5/go/vault"
//	"github.com/pulumi/pulumi-vault/sdk/v5/go/vault/kv"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			kvv1, err := vault.NewMount(ctx, "kvv1", &vault.MountArgs{
//				Path: pulumi.String("kvv1"),
//				Type: pulumi.String("kv"),
//				Options: pulumi.Map{
//					"version": pulumi.Any("1"),
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
//			secret, err := kv.NewSecret(ctx, "secret", &kv.SecretArgs{
//				Path: kvv1.Path.ApplyT(func(path string) (string, error) {
//					return fmt.Sprintf("%v/secret", path), nil
//				}).(pulumi.StringOutput),
//				DataJson: pulumi.String(json0),
//			})
//			if err != nil {
//				return err
//			}
//			_ = kv.LookupSecretOutput(ctx, kv.GetSecretOutputArgs{
//				Path: secret.Path,
//			}, nil)
//			return nil
//		})
//	}
//
// ```
// ## Required Vault Capabilities
//
// Use of this resource requires the `read` capability on the given path.
func LookupSecret(ctx *pulumi.Context, args *LookupSecretArgs, opts ...pulumi.InvokeOption) (*LookupSecretResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupSecretResult
	err := ctx.Invoke("vault:kv/getSecret:getSecret", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSecret.
type LookupSecretArgs struct {
	// The namespace of the target resource.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
	// *Available only for Vault Enterprise*.
	Namespace *string `pulumi:"namespace"`
	// Full path of the KV-V1 secret.
	Path string `pulumi:"path"`
}

// A collection of values returned by getSecret.
type LookupSecretResult struct {
	// A mapping whose keys are the top-level data keys returned from
	// Vault and whose values are the corresponding values. This map can only
	// represent string data, so any non-string values returned from Vault are
	// serialized as JSON.
	Data map[string]interface{} `pulumi:"data"`
	// JSON-encoded string that that is
	// read as the secret data at the given path.
	DataJson string `pulumi:"dataJson"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The duration of the secret lease, in seconds. Once
	// this time has passed any plan generated with this data may fail to apply.
	LeaseDuration int `pulumi:"leaseDuration"`
	// The lease identifier assigned by Vault, if any.
	LeaseId string `pulumi:"leaseId"`
	// True if the duration of this lease can be extended
	// through renewal.
	LeaseRenewable bool    `pulumi:"leaseRenewable"`
	Namespace      *string `pulumi:"namespace"`
	Path           string  `pulumi:"path"`
}

func LookupSecretOutput(ctx *pulumi.Context, args LookupSecretOutputArgs, opts ...pulumi.InvokeOption) LookupSecretResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSecretResult, error) {
			args := v.(LookupSecretArgs)
			r, err := LookupSecret(ctx, &args, opts...)
			var s LookupSecretResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSecretResultOutput)
}

// A collection of arguments for invoking getSecret.
type LookupSecretOutputArgs struct {
	// The namespace of the target resource.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
	// *Available only for Vault Enterprise*.
	Namespace pulumi.StringPtrInput `pulumi:"namespace"`
	// Full path of the KV-V1 secret.
	Path pulumi.StringInput `pulumi:"path"`
}

func (LookupSecretOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSecretArgs)(nil)).Elem()
}

// A collection of values returned by getSecret.
type LookupSecretResultOutput struct{ *pulumi.OutputState }

func (LookupSecretResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSecretResult)(nil)).Elem()
}

func (o LookupSecretResultOutput) ToLookupSecretResultOutput() LookupSecretResultOutput {
	return o
}

func (o LookupSecretResultOutput) ToLookupSecretResultOutputWithContext(ctx context.Context) LookupSecretResultOutput {
	return o
}

// A mapping whose keys are the top-level data keys returned from
// Vault and whose values are the corresponding values. This map can only
// represent string data, so any non-string values returned from Vault are
// serialized as JSON.
func (o LookupSecretResultOutput) Data() pulumi.MapOutput {
	return o.ApplyT(func(v LookupSecretResult) map[string]interface{} { return v.Data }).(pulumi.MapOutput)
}

// JSON-encoded string that that is
// read as the secret data at the given path.
func (o LookupSecretResultOutput) DataJson() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSecretResult) string { return v.DataJson }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupSecretResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSecretResult) string { return v.Id }).(pulumi.StringOutput)
}

// The duration of the secret lease, in seconds. Once
// this time has passed any plan generated with this data may fail to apply.
func (o LookupSecretResultOutput) LeaseDuration() pulumi.IntOutput {
	return o.ApplyT(func(v LookupSecretResult) int { return v.LeaseDuration }).(pulumi.IntOutput)
}

// The lease identifier assigned by Vault, if any.
func (o LookupSecretResultOutput) LeaseId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSecretResult) string { return v.LeaseId }).(pulumi.StringOutput)
}

// True if the duration of this lease can be extended
// through renewal.
func (o LookupSecretResultOutput) LeaseRenewable() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupSecretResult) bool { return v.LeaseRenewable }).(pulumi.BoolOutput)
}

func (o LookupSecretResultOutput) Namespace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSecretResult) *string { return v.Namespace }).(pulumi.StringPtrOutput)
}

func (o LookupSecretResultOutput) Path() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSecretResult) string { return v.Path }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSecretResultOutput{})
}
