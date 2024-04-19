// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package pkisecret

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-vault/sdk/v6/go/vault/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-vault/sdk/v6/go/vault"
//	"github.com/pulumi/pulumi-vault/sdk/v6/go/vault/pkiSecret"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			pki, err := vault.NewMount(ctx, "pki", &vault.MountArgs{
//				Path:        pulumi.String("pki"),
//				Type:        pulumi.String("pki"),
//				Description: pulumi.String("PKI secret engine mount"),
//			})
//			if err != nil {
//				return err
//			}
//			key, err := pkiSecret.NewSecretBackendKey(ctx, "key", &pkiSecret.SecretBackendKeyArgs{
//				Backend: pki.Path,
//				Type:    pulumi.String("internal"),
//				KeyName: pulumi.String("example"),
//				KeyType: pulumi.String("rsa"),
//				KeyBits: pulumi.Int(4096),
//			})
//			if err != nil {
//				return err
//			}
//			_ = key.KeyId.ApplyT(func(keyId string) (pkisecret.GetBackendKeyResult, error) {
//				return pkiSecret.GetBackendKeyOutput(ctx, pkisecret.GetBackendKeyOutputArgs{
//					Backend: vault_mount.Key.Path,
//					KeyRef:  keyId,
//				}, nil), nil
//			}).(pkisecret.GetBackendKeyResultOutput)
//			return nil
//		})
//	}
//
// ```
func GetBackendKey(ctx *pulumi.Context, args *GetBackendKeyArgs, opts ...pulumi.InvokeOption) (*GetBackendKeyResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetBackendKeyResult
	err := ctx.Invoke("vault:pkiSecret/getBackendKey:getBackendKey", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getBackendKey.
type GetBackendKeyArgs struct {
	// The path to the PKI secret backend to
	// read the key from, with no leading or trailing `/`s.
	Backend string `pulumi:"backend"`
	// Reference to an existing key.
	KeyRef string `pulumi:"keyRef"`
	// The namespace of the target resource.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
	// *Available only for Vault Enterprise*.
	Namespace *string `pulumi:"namespace"`
}

// A collection of values returned by getBackendKey.
type GetBackendKeyResult struct {
	Backend string `pulumi:"backend"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// ID of the key.
	KeyId string `pulumi:"keyId"`
	// Name of the key.
	KeyName string `pulumi:"keyName"`
	KeyRef  string `pulumi:"keyRef"`
	// Type of the key.
	KeyType   string  `pulumi:"keyType"`
	Namespace *string `pulumi:"namespace"`
}

func GetBackendKeyOutput(ctx *pulumi.Context, args GetBackendKeyOutputArgs, opts ...pulumi.InvokeOption) GetBackendKeyResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetBackendKeyResult, error) {
			args := v.(GetBackendKeyArgs)
			r, err := GetBackendKey(ctx, &args, opts...)
			var s GetBackendKeyResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetBackendKeyResultOutput)
}

// A collection of arguments for invoking getBackendKey.
type GetBackendKeyOutputArgs struct {
	// The path to the PKI secret backend to
	// read the key from, with no leading or trailing `/`s.
	Backend pulumi.StringInput `pulumi:"backend"`
	// Reference to an existing key.
	KeyRef pulumi.StringInput `pulumi:"keyRef"`
	// The namespace of the target resource.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
	// *Available only for Vault Enterprise*.
	Namespace pulumi.StringPtrInput `pulumi:"namespace"`
}

func (GetBackendKeyOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetBackendKeyArgs)(nil)).Elem()
}

// A collection of values returned by getBackendKey.
type GetBackendKeyResultOutput struct{ *pulumi.OutputState }

func (GetBackendKeyResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetBackendKeyResult)(nil)).Elem()
}

func (o GetBackendKeyResultOutput) ToGetBackendKeyResultOutput() GetBackendKeyResultOutput {
	return o
}

func (o GetBackendKeyResultOutput) ToGetBackendKeyResultOutputWithContext(ctx context.Context) GetBackendKeyResultOutput {
	return o
}

func (o GetBackendKeyResultOutput) Backend() pulumi.StringOutput {
	return o.ApplyT(func(v GetBackendKeyResult) string { return v.Backend }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetBackendKeyResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetBackendKeyResult) string { return v.Id }).(pulumi.StringOutput)
}

// ID of the key.
func (o GetBackendKeyResultOutput) KeyId() pulumi.StringOutput {
	return o.ApplyT(func(v GetBackendKeyResult) string { return v.KeyId }).(pulumi.StringOutput)
}

// Name of the key.
func (o GetBackendKeyResultOutput) KeyName() pulumi.StringOutput {
	return o.ApplyT(func(v GetBackendKeyResult) string { return v.KeyName }).(pulumi.StringOutput)
}

func (o GetBackendKeyResultOutput) KeyRef() pulumi.StringOutput {
	return o.ApplyT(func(v GetBackendKeyResult) string { return v.KeyRef }).(pulumi.StringOutput)
}

// Type of the key.
func (o GetBackendKeyResultOutput) KeyType() pulumi.StringOutput {
	return o.ApplyT(func(v GetBackendKeyResult) string { return v.KeyType }).(pulumi.StringOutput)
}

func (o GetBackendKeyResultOutput) Namespace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetBackendKeyResult) *string { return v.Namespace }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetBackendKeyResultOutput{})
}
