// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ldap

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-vault/sdk/v7/go/vault/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetStaticCredentials(ctx *pulumi.Context, args *GetStaticCredentialsArgs, opts ...pulumi.InvokeOption) (*GetStaticCredentialsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetStaticCredentialsResult
	err := ctx.Invoke("vault:ldap/getStaticCredentials:getStaticCredentials", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getStaticCredentials.
type GetStaticCredentialsArgs struct {
	Mount     string  `pulumi:"mount"`
	Namespace *string `pulumi:"namespace"`
	RoleName  string  `pulumi:"roleName"`
}

// A collection of values returned by getStaticCredentials.
type GetStaticCredentialsResult struct {
	Dn string `pulumi:"dn"`
	// The provider-assigned unique ID for this managed resource.
	Id                string  `pulumi:"id"`
	LastPassword      string  `pulumi:"lastPassword"`
	LastVaultRotation string  `pulumi:"lastVaultRotation"`
	Mount             string  `pulumi:"mount"`
	Namespace         *string `pulumi:"namespace"`
	Password          string  `pulumi:"password"`
	RoleName          string  `pulumi:"roleName"`
	RotationPeriod    int     `pulumi:"rotationPeriod"`
	Ttl               int     `pulumi:"ttl"`
	Username          string  `pulumi:"username"`
}

func GetStaticCredentialsOutput(ctx *pulumi.Context, args GetStaticCredentialsOutputArgs, opts ...pulumi.InvokeOption) GetStaticCredentialsResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (GetStaticCredentialsResultOutput, error) {
			args := v.(GetStaticCredentialsArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("vault:ldap/getStaticCredentials:getStaticCredentials", args, GetStaticCredentialsResultOutput{}, options).(GetStaticCredentialsResultOutput), nil
		}).(GetStaticCredentialsResultOutput)
}

// A collection of arguments for invoking getStaticCredentials.
type GetStaticCredentialsOutputArgs struct {
	Mount     pulumi.StringInput    `pulumi:"mount"`
	Namespace pulumi.StringPtrInput `pulumi:"namespace"`
	RoleName  pulumi.StringInput    `pulumi:"roleName"`
}

func (GetStaticCredentialsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetStaticCredentialsArgs)(nil)).Elem()
}

// A collection of values returned by getStaticCredentials.
type GetStaticCredentialsResultOutput struct{ *pulumi.OutputState }

func (GetStaticCredentialsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetStaticCredentialsResult)(nil)).Elem()
}

func (o GetStaticCredentialsResultOutput) ToGetStaticCredentialsResultOutput() GetStaticCredentialsResultOutput {
	return o
}

func (o GetStaticCredentialsResultOutput) ToGetStaticCredentialsResultOutputWithContext(ctx context.Context) GetStaticCredentialsResultOutput {
	return o
}

func (o GetStaticCredentialsResultOutput) Dn() pulumi.StringOutput {
	return o.ApplyT(func(v GetStaticCredentialsResult) string { return v.Dn }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetStaticCredentialsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetStaticCredentialsResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetStaticCredentialsResultOutput) LastPassword() pulumi.StringOutput {
	return o.ApplyT(func(v GetStaticCredentialsResult) string { return v.LastPassword }).(pulumi.StringOutput)
}

func (o GetStaticCredentialsResultOutput) LastVaultRotation() pulumi.StringOutput {
	return o.ApplyT(func(v GetStaticCredentialsResult) string { return v.LastVaultRotation }).(pulumi.StringOutput)
}

func (o GetStaticCredentialsResultOutput) Mount() pulumi.StringOutput {
	return o.ApplyT(func(v GetStaticCredentialsResult) string { return v.Mount }).(pulumi.StringOutput)
}

func (o GetStaticCredentialsResultOutput) Namespace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetStaticCredentialsResult) *string { return v.Namespace }).(pulumi.StringPtrOutput)
}

func (o GetStaticCredentialsResultOutput) Password() pulumi.StringOutput {
	return o.ApplyT(func(v GetStaticCredentialsResult) string { return v.Password }).(pulumi.StringOutput)
}

func (o GetStaticCredentialsResultOutput) RoleName() pulumi.StringOutput {
	return o.ApplyT(func(v GetStaticCredentialsResult) string { return v.RoleName }).(pulumi.StringOutput)
}

func (o GetStaticCredentialsResultOutput) RotationPeriod() pulumi.IntOutput {
	return o.ApplyT(func(v GetStaticCredentialsResult) int { return v.RotationPeriod }).(pulumi.IntOutput)
}

func (o GetStaticCredentialsResultOutput) Ttl() pulumi.IntOutput {
	return o.ApplyT(func(v GetStaticCredentialsResult) int { return v.Ttl }).(pulumi.IntOutput)
}

func (o GetStaticCredentialsResultOutput) Username() pulumi.StringOutput {
	return o.ApplyT(func(v GetStaticCredentialsResult) string { return v.Username }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetStaticCredentialsResultOutput{})
}
