// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package generic

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-vault/sdk/v7/go/vault/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## Example Usage
//
// ### Generic secret
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-vault/sdk/v7/go/vault/generic"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := generic.LookupSecret(ctx, &generic.LookupSecretArgs{
//				Path: "secret/rundeck_auth",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// ### KV
//
// For this example, consider `example` as a path for a KV engine.
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-vault/sdk/v7/go/vault/generic"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func notImplemented(message string) pulumi.AnyOutput {
//		panic(message)
//	}
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := generic.LookupSecret(ctx, &generic.LookupSecretArgs{
//				Path: "example/creds",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_ = notImplemented("The template_file data resource is not yet supported.")
//			return nil
//		})
//	}
//
// ```
//
// ## Required Vault Capabilities
//
// Use of this resource requires the `read` capability on the given path.
func LookupSecret(ctx *pulumi.Context, args *LookupSecretArgs, opts ...pulumi.InvokeOption) (*LookupSecretResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupSecretResult
	err := ctx.Invoke("vault:generic/getSecret:getSecret", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSecret.
type LookupSecretArgs struct {
	// The namespace of the target resource.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
	// *Available only for Vault Enterprise*.
	Namespace *string `pulumi:"namespace"`
	// The full logical path from which to request data.
	// To read data from the "generic" secret backend mounted in Vault by
	// default, this should be prefixed with `secret/`. Reading from other backends
	// with this data source is possible; consult each backend's documentation
	// to see which endpoints support the `GET` method.
	Path string `pulumi:"path"`
	// The version of the secret to read. This is used by the
	// Vault KV secrets engine - version 2 to indicate which version of the secret
	// to read.
	Version *int `pulumi:"version"`
	// If set to true, stores `leaseStartTime` in the TF state.
	// Note that storing the `leaseStartTime` in the TF state will cause a persistent drift
	// on every `pulumi preview` and will require a `pulumi up`.
	WithLeaseStartTime *bool `pulumi:"withLeaseStartTime"`
}

// A collection of values returned by getSecret.
type LookupSecretResult struct {
	// A mapping whose keys are the top-level data keys returned from
	// Vault and whose values are the corresponding values. This map can only
	// represent string data, so any non-string values returned from Vault are
	// serialized as JSON.
	Data map[string]string `pulumi:"data"`
	// A string containing the full data payload retrieved from
	// Vault, serialized in JSON format.
	DataJson string `pulumi:"dataJson"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The duration of the secret lease, in seconds relative
	// to the time the data was requested. Once this time has passed any plan
	// generated with this data may fail to apply.
	LeaseDuration int `pulumi:"leaseDuration"`
	// The lease identifier assigned by Vault, if any.
	LeaseId            string  `pulumi:"leaseId"`
	LeaseRenewable     bool    `pulumi:"leaseRenewable"`
	LeaseStartTime     string  `pulumi:"leaseStartTime"`
	Namespace          *string `pulumi:"namespace"`
	Path               string  `pulumi:"path"`
	Version            *int    `pulumi:"version"`
	WithLeaseStartTime *bool   `pulumi:"withLeaseStartTime"`
}

func LookupSecretOutput(ctx *pulumi.Context, args LookupSecretOutputArgs, opts ...pulumi.InvokeOption) LookupSecretResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupSecretResultOutput, error) {
			args := v.(LookupSecretArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("vault:generic/getSecret:getSecret", args, LookupSecretResultOutput{}, options).(LookupSecretResultOutput), nil
		}).(LookupSecretResultOutput)
}

// A collection of arguments for invoking getSecret.
type LookupSecretOutputArgs struct {
	// The namespace of the target resource.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
	// *Available only for Vault Enterprise*.
	Namespace pulumi.StringPtrInput `pulumi:"namespace"`
	// The full logical path from which to request data.
	// To read data from the "generic" secret backend mounted in Vault by
	// default, this should be prefixed with `secret/`. Reading from other backends
	// with this data source is possible; consult each backend's documentation
	// to see which endpoints support the `GET` method.
	Path pulumi.StringInput `pulumi:"path"`
	// The version of the secret to read. This is used by the
	// Vault KV secrets engine - version 2 to indicate which version of the secret
	// to read.
	Version pulumi.IntPtrInput `pulumi:"version"`
	// If set to true, stores `leaseStartTime` in the TF state.
	// Note that storing the `leaseStartTime` in the TF state will cause a persistent drift
	// on every `pulumi preview` and will require a `pulumi up`.
	WithLeaseStartTime pulumi.BoolPtrInput `pulumi:"withLeaseStartTime"`
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
func (o LookupSecretResultOutput) Data() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupSecretResult) map[string]string { return v.Data }).(pulumi.StringMapOutput)
}

// A string containing the full data payload retrieved from
// Vault, serialized in JSON format.
func (o LookupSecretResultOutput) DataJson() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSecretResult) string { return v.DataJson }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupSecretResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSecretResult) string { return v.Id }).(pulumi.StringOutput)
}

// The duration of the secret lease, in seconds relative
// to the time the data was requested. Once this time has passed any plan
// generated with this data may fail to apply.
func (o LookupSecretResultOutput) LeaseDuration() pulumi.IntOutput {
	return o.ApplyT(func(v LookupSecretResult) int { return v.LeaseDuration }).(pulumi.IntOutput)
}

// The lease identifier assigned by Vault, if any.
func (o LookupSecretResultOutput) LeaseId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSecretResult) string { return v.LeaseId }).(pulumi.StringOutput)
}

func (o LookupSecretResultOutput) LeaseRenewable() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupSecretResult) bool { return v.LeaseRenewable }).(pulumi.BoolOutput)
}

func (o LookupSecretResultOutput) LeaseStartTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSecretResult) string { return v.LeaseStartTime }).(pulumi.StringOutput)
}

func (o LookupSecretResultOutput) Namespace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSecretResult) *string { return v.Namespace }).(pulumi.StringPtrOutput)
}

func (o LookupSecretResultOutput) Path() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSecretResult) string { return v.Path }).(pulumi.StringOutput)
}

func (o LookupSecretResultOutput) Version() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupSecretResult) *int { return v.Version }).(pulumi.IntPtrOutput)
}

func (o LookupSecretResultOutput) WithLeaseStartTime() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupSecretResult) *bool { return v.WithLeaseStartTime }).(pulumi.BoolPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSecretResultOutput{})
}
