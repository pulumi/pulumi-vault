// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ldap

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-vault/sdk/v6/go/vault/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetDynamicCredentials(ctx *pulumi.Context, args *GetDynamicCredentialsArgs, opts ...pulumi.InvokeOption) (*GetDynamicCredentialsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetDynamicCredentialsResult
	err := ctx.Invoke("vault:ldap/getDynamicCredentials:getDynamicCredentials", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getDynamicCredentials.
type GetDynamicCredentialsArgs struct {
	Mount     string  `pulumi:"mount"`
	Namespace *string `pulumi:"namespace"`
	RoleName  string  `pulumi:"roleName"`
}

// A collection of values returned by getDynamicCredentials.
type GetDynamicCredentialsResult struct {
	DistinguishedNames []string `pulumi:"distinguishedNames"`
	// The provider-assigned unique ID for this managed resource.
	Id             string  `pulumi:"id"`
	LeaseDuration  int     `pulumi:"leaseDuration"`
	LeaseId        string  `pulumi:"leaseId"`
	LeaseRenewable bool    `pulumi:"leaseRenewable"`
	Mount          string  `pulumi:"mount"`
	Namespace      *string `pulumi:"namespace"`
	Password       string  `pulumi:"password"`
	RoleName       string  `pulumi:"roleName"`
	Username       string  `pulumi:"username"`
}

func GetDynamicCredentialsOutput(ctx *pulumi.Context, args GetDynamicCredentialsOutputArgs, opts ...pulumi.InvokeOption) GetDynamicCredentialsResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (GetDynamicCredentialsResultOutput, error) {
			args := v.(GetDynamicCredentialsArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("vault:ldap/getDynamicCredentials:getDynamicCredentials", args, GetDynamicCredentialsResultOutput{}, options).(GetDynamicCredentialsResultOutput), nil
		}).(GetDynamicCredentialsResultOutput)
}

// A collection of arguments for invoking getDynamicCredentials.
type GetDynamicCredentialsOutputArgs struct {
	Mount     pulumi.StringInput    `pulumi:"mount"`
	Namespace pulumi.StringPtrInput `pulumi:"namespace"`
	RoleName  pulumi.StringInput    `pulumi:"roleName"`
}

func (GetDynamicCredentialsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetDynamicCredentialsArgs)(nil)).Elem()
}

// A collection of values returned by getDynamicCredentials.
type GetDynamicCredentialsResultOutput struct{ *pulumi.OutputState }

func (GetDynamicCredentialsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetDynamicCredentialsResult)(nil)).Elem()
}

func (o GetDynamicCredentialsResultOutput) ToGetDynamicCredentialsResultOutput() GetDynamicCredentialsResultOutput {
	return o
}

func (o GetDynamicCredentialsResultOutput) ToGetDynamicCredentialsResultOutputWithContext(ctx context.Context) GetDynamicCredentialsResultOutput {
	return o
}

func (o GetDynamicCredentialsResultOutput) DistinguishedNames() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetDynamicCredentialsResult) []string { return v.DistinguishedNames }).(pulumi.StringArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetDynamicCredentialsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetDynamicCredentialsResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetDynamicCredentialsResultOutput) LeaseDuration() pulumi.IntOutput {
	return o.ApplyT(func(v GetDynamicCredentialsResult) int { return v.LeaseDuration }).(pulumi.IntOutput)
}

func (o GetDynamicCredentialsResultOutput) LeaseId() pulumi.StringOutput {
	return o.ApplyT(func(v GetDynamicCredentialsResult) string { return v.LeaseId }).(pulumi.StringOutput)
}

func (o GetDynamicCredentialsResultOutput) LeaseRenewable() pulumi.BoolOutput {
	return o.ApplyT(func(v GetDynamicCredentialsResult) bool { return v.LeaseRenewable }).(pulumi.BoolOutput)
}

func (o GetDynamicCredentialsResultOutput) Mount() pulumi.StringOutput {
	return o.ApplyT(func(v GetDynamicCredentialsResult) string { return v.Mount }).(pulumi.StringOutput)
}

func (o GetDynamicCredentialsResultOutput) Namespace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetDynamicCredentialsResult) *string { return v.Namespace }).(pulumi.StringPtrOutput)
}

func (o GetDynamicCredentialsResultOutput) Password() pulumi.StringOutput {
	return o.ApplyT(func(v GetDynamicCredentialsResult) string { return v.Password }).(pulumi.StringOutput)
}

func (o GetDynamicCredentialsResultOutput) RoleName() pulumi.StringOutput {
	return o.ApplyT(func(v GetDynamicCredentialsResult) string { return v.RoleName }).(pulumi.StringOutput)
}

func (o GetDynamicCredentialsResultOutput) Username() pulumi.StringOutput {
	return o.ApplyT(func(v GetDynamicCredentialsResult) string { return v.Username }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetDynamicCredentialsResultOutput{})
}
