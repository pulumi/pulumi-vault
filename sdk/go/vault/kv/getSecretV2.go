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
func LookupSecretV2(ctx *pulumi.Context, args *LookupSecretV2Args, opts ...pulumi.InvokeOption) (*LookupSecretV2Result, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupSecretV2Result
	err := ctx.Invoke("vault:kv/getSecretV2:getSecretV2", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSecretV2.
type LookupSecretV2Args struct {
	// Path where KV-V2 engine is mounted.
	Mount string `pulumi:"mount"`
	// Full name of the secret. For a nested secret
	// the name is the nested path excluding the mount and data
	// prefix. For example, for a secret at `kvv2/data/foo/bar/baz`
	// the name is `foo/bar/baz`.
	Name string `pulumi:"name"`
	// The namespace of the target resource.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
	// *Available only for Vault Enterprise*.
	Namespace *string `pulumi:"namespace"`
	// Version of the secret to retrieve.
	Version *int `pulumi:"version"`
}

// A collection of values returned by getSecretV2.
type LookupSecretV2Result struct {
	// Time at which secret was created.
	CreatedTime string `pulumi:"createdTime"`
	// Custom metadata for the secret.
	CustomMetadata map[string]interface{} `pulumi:"customMetadata"`
	// A mapping whose keys are the top-level data keys returned from
	// Vault and whose values are the corresponding values. This map can only
	// represent string data, so any non-string values returned from Vault are
	// serialized as JSON.
	Data map[string]interface{} `pulumi:"data"`
	// JSON-encoded string that that is
	// read as the secret data at the given path.
	DataJson string `pulumi:"dataJson"`
	// Deletion time for the secret.
	DeletionTime string `pulumi:"deletionTime"`
	// Indicates whether the secret has been destroyed.
	Destroyed bool `pulumi:"destroyed"`
	// The provider-assigned unique ID for this managed resource.
	Id        string  `pulumi:"id"`
	Mount     string  `pulumi:"mount"`
	Name      string  `pulumi:"name"`
	Namespace *string `pulumi:"namespace"`
	// Full path where the KVV2 secret is written.
	Path string `pulumi:"path"`
	// Version of the secret.
	Version *int `pulumi:"version"`
}

func LookupSecretV2Output(ctx *pulumi.Context, args LookupSecretV2OutputArgs, opts ...pulumi.InvokeOption) LookupSecretV2ResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSecretV2Result, error) {
			args := v.(LookupSecretV2Args)
			r, err := LookupSecretV2(ctx, &args, opts...)
			var s LookupSecretV2Result
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSecretV2ResultOutput)
}

// A collection of arguments for invoking getSecretV2.
type LookupSecretV2OutputArgs struct {
	// Path where KV-V2 engine is mounted.
	Mount pulumi.StringInput `pulumi:"mount"`
	// Full name of the secret. For a nested secret
	// the name is the nested path excluding the mount and data
	// prefix. For example, for a secret at `kvv2/data/foo/bar/baz`
	// the name is `foo/bar/baz`.
	Name pulumi.StringInput `pulumi:"name"`
	// The namespace of the target resource.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
	// *Available only for Vault Enterprise*.
	Namespace pulumi.StringPtrInput `pulumi:"namespace"`
	// Version of the secret to retrieve.
	Version pulumi.IntPtrInput `pulumi:"version"`
}

func (LookupSecretV2OutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSecretV2Args)(nil)).Elem()
}

// A collection of values returned by getSecretV2.
type LookupSecretV2ResultOutput struct{ *pulumi.OutputState }

func (LookupSecretV2ResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSecretV2Result)(nil)).Elem()
}

func (o LookupSecretV2ResultOutput) ToLookupSecretV2ResultOutput() LookupSecretV2ResultOutput {
	return o
}

func (o LookupSecretV2ResultOutput) ToLookupSecretV2ResultOutputWithContext(ctx context.Context) LookupSecretV2ResultOutput {
	return o
}

// Time at which secret was created.
func (o LookupSecretV2ResultOutput) CreatedTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSecretV2Result) string { return v.CreatedTime }).(pulumi.StringOutput)
}

// Custom metadata for the secret.
func (o LookupSecretV2ResultOutput) CustomMetadata() pulumi.MapOutput {
	return o.ApplyT(func(v LookupSecretV2Result) map[string]interface{} { return v.CustomMetadata }).(pulumi.MapOutput)
}

// A mapping whose keys are the top-level data keys returned from
// Vault and whose values are the corresponding values. This map can only
// represent string data, so any non-string values returned from Vault are
// serialized as JSON.
func (o LookupSecretV2ResultOutput) Data() pulumi.MapOutput {
	return o.ApplyT(func(v LookupSecretV2Result) map[string]interface{} { return v.Data }).(pulumi.MapOutput)
}

// JSON-encoded string that that is
// read as the secret data at the given path.
func (o LookupSecretV2ResultOutput) DataJson() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSecretV2Result) string { return v.DataJson }).(pulumi.StringOutput)
}

// Deletion time for the secret.
func (o LookupSecretV2ResultOutput) DeletionTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSecretV2Result) string { return v.DeletionTime }).(pulumi.StringOutput)
}

// Indicates whether the secret has been destroyed.
func (o LookupSecretV2ResultOutput) Destroyed() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupSecretV2Result) bool { return v.Destroyed }).(pulumi.BoolOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupSecretV2ResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSecretV2Result) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupSecretV2ResultOutput) Mount() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSecretV2Result) string { return v.Mount }).(pulumi.StringOutput)
}

func (o LookupSecretV2ResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSecretV2Result) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupSecretV2ResultOutput) Namespace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSecretV2Result) *string { return v.Namespace }).(pulumi.StringPtrOutput)
}

// Full path where the KVV2 secret is written.
func (o LookupSecretV2ResultOutput) Path() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSecretV2Result) string { return v.Path }).(pulumi.StringOutput)
}

// Version of the secret.
func (o LookupSecretV2ResultOutput) Version() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupSecretV2Result) *int { return v.Version }).(pulumi.IntPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSecretV2ResultOutput{})
}
