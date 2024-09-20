// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vault

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-vault/sdk/v6/go/vault/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## Example Usage
//
// ### Child namespaces
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-vault/sdk/v6/go/vault"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := vault.GetNamespaces(ctx, nil, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetNamespaces(ctx *pulumi.Context, args *GetNamespacesArgs, opts ...pulumi.InvokeOption) (*GetNamespacesResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetNamespacesResult
	err := ctx.Invoke("vault:index/getNamespaces:getNamespaces", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getNamespaces.
type GetNamespacesArgs struct {
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
	Namespace *string `pulumi:"namespace"`
}

// A collection of values returned by getNamespaces.
type GetNamespacesResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id        string  `pulumi:"id"`
	Namespace *string `pulumi:"namespace"`
	// Set of the paths of direct child namespaces.
	Paths []string `pulumi:"paths"`
}

func GetNamespacesOutput(ctx *pulumi.Context, args GetNamespacesOutputArgs, opts ...pulumi.InvokeOption) GetNamespacesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetNamespacesResultOutput, error) {
			args := v.(GetNamespacesArgs)
			opts = internal.PkgInvokeDefaultOpts(opts)
			var rv GetNamespacesResult
			secret, err := ctx.InvokePackageRaw("vault:index/getNamespaces:getNamespaces", args, &rv, "", opts...)
			if err != nil {
				return GetNamespacesResultOutput{}, err
			}

			output := pulumi.ToOutput(rv).(GetNamespacesResultOutput)
			if secret {
				return pulumi.ToSecret(output).(GetNamespacesResultOutput), nil
			}
			return output, nil
		}).(GetNamespacesResultOutput)
}

// A collection of arguments for invoking getNamespaces.
type GetNamespacesOutputArgs struct {
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
	Namespace pulumi.StringPtrInput `pulumi:"namespace"`
}

func (GetNamespacesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetNamespacesArgs)(nil)).Elem()
}

// A collection of values returned by getNamespaces.
type GetNamespacesResultOutput struct{ *pulumi.OutputState }

func (GetNamespacesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetNamespacesResult)(nil)).Elem()
}

func (o GetNamespacesResultOutput) ToGetNamespacesResultOutput() GetNamespacesResultOutput {
	return o
}

func (o GetNamespacesResultOutput) ToGetNamespacesResultOutputWithContext(ctx context.Context) GetNamespacesResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o GetNamespacesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetNamespacesResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetNamespacesResultOutput) Namespace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetNamespacesResult) *string { return v.Namespace }).(pulumi.StringPtrOutput)
}

// Set of the paths of direct child namespaces.
func (o GetNamespacesResultOutput) Paths() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetNamespacesResult) []string { return v.Paths }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetNamespacesResultOutput{})
}
