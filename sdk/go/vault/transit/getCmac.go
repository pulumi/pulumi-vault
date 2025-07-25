// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package transit

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-vault/sdk/v7/go/vault/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This is a data source which can be used to generate a CMAC using a Vault Transit key.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-vault/sdk/v7/go/vault/transit"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := transit.GetCmac(ctx, &transit.GetCmacArgs{
//				Path:  "transit",
//				Name:  "test",
//				Input: pulumi.StringRef("aGVsbG8gd29ybGQ="),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetCmac(ctx *pulumi.Context, args *GetCmacArgs, opts ...pulumi.InvokeOption) (*GetCmacResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetCmacResult
	err := ctx.Invoke("vault:transit/getCmac:getCmac", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getCmac.
type GetCmacArgs struct {
	BatchInputs []map[string]string `pulumi:"batchInputs"`
	// The results returned from Vault if using `batchInput`
	BatchResults []map[string]string `pulumi:"batchResults"`
	// The CMAC returned from Vault if using `input`
	Cmac         *string `pulumi:"cmac"`
	Input        *string `pulumi:"input"`
	KeyVersion   *int    `pulumi:"keyVersion"`
	MacLength    *int    `pulumi:"macLength"`
	Name         string  `pulumi:"name"`
	Namespace    *string `pulumi:"namespace"`
	Path         string  `pulumi:"path"`
	UrlMacLength *int    `pulumi:"urlMacLength"`
}

// A collection of values returned by getCmac.
type GetCmacResult struct {
	BatchInputs []map[string]string `pulumi:"batchInputs"`
	// The results returned from Vault if using `batchInput`
	BatchResults []map[string]string `pulumi:"batchResults"`
	// The CMAC returned from Vault if using `input`
	Cmac string `pulumi:"cmac"`
	// The provider-assigned unique ID for this managed resource.
	Id           string  `pulumi:"id"`
	Input        *string `pulumi:"input"`
	KeyVersion   *int    `pulumi:"keyVersion"`
	MacLength    *int    `pulumi:"macLength"`
	Name         string  `pulumi:"name"`
	Namespace    *string `pulumi:"namespace"`
	Path         string  `pulumi:"path"`
	UrlMacLength *int    `pulumi:"urlMacLength"`
}

func GetCmacOutput(ctx *pulumi.Context, args GetCmacOutputArgs, opts ...pulumi.InvokeOption) GetCmacResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (GetCmacResultOutput, error) {
			args := v.(GetCmacArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("vault:transit/getCmac:getCmac", args, GetCmacResultOutput{}, options).(GetCmacResultOutput), nil
		}).(GetCmacResultOutput)
}

// A collection of arguments for invoking getCmac.
type GetCmacOutputArgs struct {
	BatchInputs pulumi.StringMapArrayInput `pulumi:"batchInputs"`
	// The results returned from Vault if using `batchInput`
	BatchResults pulumi.StringMapArrayInput `pulumi:"batchResults"`
	// The CMAC returned from Vault if using `input`
	Cmac         pulumi.StringPtrInput `pulumi:"cmac"`
	Input        pulumi.StringPtrInput `pulumi:"input"`
	KeyVersion   pulumi.IntPtrInput    `pulumi:"keyVersion"`
	MacLength    pulumi.IntPtrInput    `pulumi:"macLength"`
	Name         pulumi.StringInput    `pulumi:"name"`
	Namespace    pulumi.StringPtrInput `pulumi:"namespace"`
	Path         pulumi.StringInput    `pulumi:"path"`
	UrlMacLength pulumi.IntPtrInput    `pulumi:"urlMacLength"`
}

func (GetCmacOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetCmacArgs)(nil)).Elem()
}

// A collection of values returned by getCmac.
type GetCmacResultOutput struct{ *pulumi.OutputState }

func (GetCmacResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetCmacResult)(nil)).Elem()
}

func (o GetCmacResultOutput) ToGetCmacResultOutput() GetCmacResultOutput {
	return o
}

func (o GetCmacResultOutput) ToGetCmacResultOutputWithContext(ctx context.Context) GetCmacResultOutput {
	return o
}

func (o GetCmacResultOutput) BatchInputs() pulumi.StringMapArrayOutput {
	return o.ApplyT(func(v GetCmacResult) []map[string]string { return v.BatchInputs }).(pulumi.StringMapArrayOutput)
}

// The results returned from Vault if using `batchInput`
func (o GetCmacResultOutput) BatchResults() pulumi.StringMapArrayOutput {
	return o.ApplyT(func(v GetCmacResult) []map[string]string { return v.BatchResults }).(pulumi.StringMapArrayOutput)
}

// The CMAC returned from Vault if using `input`
func (o GetCmacResultOutput) Cmac() pulumi.StringOutput {
	return o.ApplyT(func(v GetCmacResult) string { return v.Cmac }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetCmacResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetCmacResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetCmacResultOutput) Input() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetCmacResult) *string { return v.Input }).(pulumi.StringPtrOutput)
}

func (o GetCmacResultOutput) KeyVersion() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GetCmacResult) *int { return v.KeyVersion }).(pulumi.IntPtrOutput)
}

func (o GetCmacResultOutput) MacLength() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GetCmacResult) *int { return v.MacLength }).(pulumi.IntPtrOutput)
}

func (o GetCmacResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GetCmacResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o GetCmacResultOutput) Namespace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetCmacResult) *string { return v.Namespace }).(pulumi.StringPtrOutput)
}

func (o GetCmacResultOutput) Path() pulumi.StringOutput {
	return o.ApplyT(func(v GetCmacResult) string { return v.Path }).(pulumi.StringOutput)
}

func (o GetCmacResultOutput) UrlMacLength() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GetCmacResult) *int { return v.UrlMacLength }).(pulumi.IntPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetCmacResultOutput{})
}
