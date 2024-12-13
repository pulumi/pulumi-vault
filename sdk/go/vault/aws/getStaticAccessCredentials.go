// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package aws

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-vault/sdk/v6/go/vault/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetStaticAccessCredentials(ctx *pulumi.Context, args *GetStaticAccessCredentialsArgs, opts ...pulumi.InvokeOption) (*GetStaticAccessCredentialsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetStaticAccessCredentialsResult
	err := ctx.Invoke("vault:aws/getStaticAccessCredentials:getStaticAccessCredentials", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getStaticAccessCredentials.
type GetStaticAccessCredentialsArgs struct {
	Backend   string  `pulumi:"backend"`
	Name      string  `pulumi:"name"`
	Namespace *string `pulumi:"namespace"`
}

// A collection of values returned by getStaticAccessCredentials.
type GetStaticAccessCredentialsResult struct {
	AccessKey string `pulumi:"accessKey"`
	Backend   string `pulumi:"backend"`
	// The provider-assigned unique ID for this managed resource.
	Id        string  `pulumi:"id"`
	Name      string  `pulumi:"name"`
	Namespace *string `pulumi:"namespace"`
	SecretKey string  `pulumi:"secretKey"`
}

func GetStaticAccessCredentialsOutput(ctx *pulumi.Context, args GetStaticAccessCredentialsOutputArgs, opts ...pulumi.InvokeOption) GetStaticAccessCredentialsResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (GetStaticAccessCredentialsResultOutput, error) {
			args := v.(GetStaticAccessCredentialsArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("vault:aws/getStaticAccessCredentials:getStaticAccessCredentials", args, GetStaticAccessCredentialsResultOutput{}, options).(GetStaticAccessCredentialsResultOutput), nil
		}).(GetStaticAccessCredentialsResultOutput)
}

// A collection of arguments for invoking getStaticAccessCredentials.
type GetStaticAccessCredentialsOutputArgs struct {
	Backend   pulumi.StringInput    `pulumi:"backend"`
	Name      pulumi.StringInput    `pulumi:"name"`
	Namespace pulumi.StringPtrInput `pulumi:"namespace"`
}

func (GetStaticAccessCredentialsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetStaticAccessCredentialsArgs)(nil)).Elem()
}

// A collection of values returned by getStaticAccessCredentials.
type GetStaticAccessCredentialsResultOutput struct{ *pulumi.OutputState }

func (GetStaticAccessCredentialsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetStaticAccessCredentialsResult)(nil)).Elem()
}

func (o GetStaticAccessCredentialsResultOutput) ToGetStaticAccessCredentialsResultOutput() GetStaticAccessCredentialsResultOutput {
	return o
}

func (o GetStaticAccessCredentialsResultOutput) ToGetStaticAccessCredentialsResultOutputWithContext(ctx context.Context) GetStaticAccessCredentialsResultOutput {
	return o
}

func (o GetStaticAccessCredentialsResultOutput) AccessKey() pulumi.StringOutput {
	return o.ApplyT(func(v GetStaticAccessCredentialsResult) string { return v.AccessKey }).(pulumi.StringOutput)
}

func (o GetStaticAccessCredentialsResultOutput) Backend() pulumi.StringOutput {
	return o.ApplyT(func(v GetStaticAccessCredentialsResult) string { return v.Backend }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetStaticAccessCredentialsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetStaticAccessCredentialsResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetStaticAccessCredentialsResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GetStaticAccessCredentialsResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o GetStaticAccessCredentialsResultOutput) Namespace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetStaticAccessCredentialsResult) *string { return v.Namespace }).(pulumi.StringPtrOutput)
}

func (o GetStaticAccessCredentialsResultOutput) SecretKey() pulumi.StringOutput {
	return o.ApplyT(func(v GetStaticAccessCredentialsResult) string { return v.SecretKey }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetStaticAccessCredentialsResultOutput{})
}
