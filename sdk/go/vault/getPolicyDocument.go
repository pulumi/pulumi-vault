// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vault

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-vault/sdk/v6/go/vault/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This is a data source which can be used to construct a HCL representation of an Vault policy document, for use with resources which expect policy documents, such as the `Policy` resource.
//
// ## Example Usage
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
//			example, err := vault.GetPolicyDocument(ctx, &vault.GetPolicyDocumentArgs{
//				Rules: []vault.GetPolicyDocumentRule{
//					{
//						Path: "secret/*",
//						Capabilities: []string{
//							"create",
//							"read",
//							"update",
//							"delete",
//							"list",
//						},
//						Description: pulumi.StringRef("allow all on secrets"),
//					},
//				},
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = vault.NewPolicy(ctx, "example", &vault.PolicyArgs{
//				Name:   pulumi.String("example_policy"),
//				Policy: pulumi.String(example.Hcl),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetPolicyDocument(ctx *pulumi.Context, args *GetPolicyDocumentArgs, opts ...pulumi.InvokeOption) (*GetPolicyDocumentResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetPolicyDocumentResult
	err := ctx.Invoke("vault:index/getPolicyDocument:getPolicyDocument", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getPolicyDocument.
type GetPolicyDocumentArgs struct {
	Namespace *string                 `pulumi:"namespace"`
	Rules     []GetPolicyDocumentRule `pulumi:"rules"`
}

// A collection of values returned by getPolicyDocument.
type GetPolicyDocumentResult struct {
	// The above arguments serialized as a standard Vault HCL policy document.
	Hcl string `pulumi:"hcl"`
	// The provider-assigned unique ID for this managed resource.
	Id        string                  `pulumi:"id"`
	Namespace *string                 `pulumi:"namespace"`
	Rules     []GetPolicyDocumentRule `pulumi:"rules"`
}

func GetPolicyDocumentOutput(ctx *pulumi.Context, args GetPolicyDocumentOutputArgs, opts ...pulumi.InvokeOption) GetPolicyDocumentResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (GetPolicyDocumentResultOutput, error) {
			args := v.(GetPolicyDocumentArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("vault:index/getPolicyDocument:getPolicyDocument", args, GetPolicyDocumentResultOutput{}, options).(GetPolicyDocumentResultOutput), nil
		}).(GetPolicyDocumentResultOutput)
}

// A collection of arguments for invoking getPolicyDocument.
type GetPolicyDocumentOutputArgs struct {
	Namespace pulumi.StringPtrInput           `pulumi:"namespace"`
	Rules     GetPolicyDocumentRuleArrayInput `pulumi:"rules"`
}

func (GetPolicyDocumentOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetPolicyDocumentArgs)(nil)).Elem()
}

// A collection of values returned by getPolicyDocument.
type GetPolicyDocumentResultOutput struct{ *pulumi.OutputState }

func (GetPolicyDocumentResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetPolicyDocumentResult)(nil)).Elem()
}

func (o GetPolicyDocumentResultOutput) ToGetPolicyDocumentResultOutput() GetPolicyDocumentResultOutput {
	return o
}

func (o GetPolicyDocumentResultOutput) ToGetPolicyDocumentResultOutputWithContext(ctx context.Context) GetPolicyDocumentResultOutput {
	return o
}

// The above arguments serialized as a standard Vault HCL policy document.
func (o GetPolicyDocumentResultOutput) Hcl() pulumi.StringOutput {
	return o.ApplyT(func(v GetPolicyDocumentResult) string { return v.Hcl }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetPolicyDocumentResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetPolicyDocumentResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetPolicyDocumentResultOutput) Namespace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetPolicyDocumentResult) *string { return v.Namespace }).(pulumi.StringPtrOutput)
}

func (o GetPolicyDocumentResultOutput) Rules() GetPolicyDocumentRuleArrayOutput {
	return o.ApplyT(func(v GetPolicyDocumentResult) []GetPolicyDocumentRule { return v.Rules }).(GetPolicyDocumentRuleArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetPolicyDocumentResultOutput{})
}
