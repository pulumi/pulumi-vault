// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package transit

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-vault/sdk/v6/go/vault/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This is a data source which can be used to encrypt plaintext using a Vault Transit key.
func GetEncrypt(ctx *pulumi.Context, args *GetEncryptArgs, opts ...pulumi.InvokeOption) (*GetEncryptResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetEncryptResult
	err := ctx.Invoke("vault:transit/getEncrypt:getEncrypt", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getEncrypt.
type GetEncryptArgs struct {
	Backend    string  `pulumi:"backend"`
	Context    *string `pulumi:"context"`
	Key        string  `pulumi:"key"`
	KeyVersion *int    `pulumi:"keyVersion"`
	Namespace  *string `pulumi:"namespace"`
	Plaintext  string  `pulumi:"plaintext"`
}

// A collection of values returned by getEncrypt.
type GetEncryptResult struct {
	Backend string `pulumi:"backend"`
	// Encrypted ciphertext returned from Vault
	Ciphertext string  `pulumi:"ciphertext"`
	Context    *string `pulumi:"context"`
	// The provider-assigned unique ID for this managed resource.
	Id         string  `pulumi:"id"`
	Key        string  `pulumi:"key"`
	KeyVersion *int    `pulumi:"keyVersion"`
	Namespace  *string `pulumi:"namespace"`
	Plaintext  string  `pulumi:"plaintext"`
}

func GetEncryptOutput(ctx *pulumi.Context, args GetEncryptOutputArgs, opts ...pulumi.InvokeOption) GetEncryptResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetEncryptResult, error) {
			args := v.(GetEncryptArgs)
			r, err := GetEncrypt(ctx, &args, opts...)
			var s GetEncryptResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetEncryptResultOutput)
}

// A collection of arguments for invoking getEncrypt.
type GetEncryptOutputArgs struct {
	Backend    pulumi.StringInput    `pulumi:"backend"`
	Context    pulumi.StringPtrInput `pulumi:"context"`
	Key        pulumi.StringInput    `pulumi:"key"`
	KeyVersion pulumi.IntPtrInput    `pulumi:"keyVersion"`
	Namespace  pulumi.StringPtrInput `pulumi:"namespace"`
	Plaintext  pulumi.StringInput    `pulumi:"plaintext"`
}

func (GetEncryptOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetEncryptArgs)(nil)).Elem()
}

// A collection of values returned by getEncrypt.
type GetEncryptResultOutput struct{ *pulumi.OutputState }

func (GetEncryptResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetEncryptResult)(nil)).Elem()
}

func (o GetEncryptResultOutput) ToGetEncryptResultOutput() GetEncryptResultOutput {
	return o
}

func (o GetEncryptResultOutput) ToGetEncryptResultOutputWithContext(ctx context.Context) GetEncryptResultOutput {
	return o
}

func (o GetEncryptResultOutput) Backend() pulumi.StringOutput {
	return o.ApplyT(func(v GetEncryptResult) string { return v.Backend }).(pulumi.StringOutput)
}

// Encrypted ciphertext returned from Vault
func (o GetEncryptResultOutput) Ciphertext() pulumi.StringOutput {
	return o.ApplyT(func(v GetEncryptResult) string { return v.Ciphertext }).(pulumi.StringOutput)
}

func (o GetEncryptResultOutput) Context() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetEncryptResult) *string { return v.Context }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetEncryptResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetEncryptResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetEncryptResultOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v GetEncryptResult) string { return v.Key }).(pulumi.StringOutput)
}

func (o GetEncryptResultOutput) KeyVersion() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GetEncryptResult) *int { return v.KeyVersion }).(pulumi.IntPtrOutput)
}

func (o GetEncryptResultOutput) Namespace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetEncryptResult) *string { return v.Namespace }).(pulumi.StringPtrOutput)
}

func (o GetEncryptResultOutput) Plaintext() pulumi.StringOutput {
	return o.ApplyT(func(v GetEncryptResult) string { return v.Plaintext }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetEncryptResultOutput{})
}
