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
//			root, err := pkiSecret.NewSecretBackendRootCert(ctx, "root", &pkiSecret.SecretBackendRootCertArgs{
//				Backend:    pki.Path,
//				Type:       pulumi.String("internal"),
//				CommonName: pulumi.String("example"),
//				Ttl:        pulumi.String("86400"),
//				KeyName:    pulumi.String("example"),
//			})
//			if err != nil {
//				return err
//			}
//			_ = pkiSecret.GetBackendKeysOutput(ctx, pkisecret.GetBackendKeysOutputArgs{
//				Backend: root.Backend,
//			}, nil)
//			return nil
//		})
//	}
//
// ```
func GetBackendKeys(ctx *pulumi.Context, args *GetBackendKeysArgs, opts ...pulumi.InvokeOption) (*GetBackendKeysResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetBackendKeysResult
	err := ctx.Invoke("vault:pkiSecret/getBackendKeys:getBackendKeys", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getBackendKeys.
type GetBackendKeysArgs struct {
	// The path to the PKI secret backend to
	// read the keys from, with no leading or trailing `/`s.
	Backend string `pulumi:"backend"`
	// The namespace of the target resource.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
	// *Available only for Vault Enterprise*.
	Namespace *string `pulumi:"namespace"`
}

// A collection of values returned by getBackendKeys.
type GetBackendKeysResult struct {
	Backend string `pulumi:"backend"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Map of key strings read from Vault.
	KeyInfo map[string]string `pulumi:"keyInfo"`
	// JSON-encoded key data read from Vault.
	KeyInfoJson string `pulumi:"keyInfoJson"`
	// Keys used under the backend path.
	Keys      []string `pulumi:"keys"`
	Namespace *string  `pulumi:"namespace"`
}

func GetBackendKeysOutput(ctx *pulumi.Context, args GetBackendKeysOutputArgs, opts ...pulumi.InvokeOption) GetBackendKeysResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetBackendKeysResult, error) {
			args := v.(GetBackendKeysArgs)
			r, err := GetBackendKeys(ctx, &args, opts...)
			var s GetBackendKeysResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetBackendKeysResultOutput)
}

// A collection of arguments for invoking getBackendKeys.
type GetBackendKeysOutputArgs struct {
	// The path to the PKI secret backend to
	// read the keys from, with no leading or trailing `/`s.
	Backend pulumi.StringInput `pulumi:"backend"`
	// The namespace of the target resource.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
	// *Available only for Vault Enterprise*.
	Namespace pulumi.StringPtrInput `pulumi:"namespace"`
}

func (GetBackendKeysOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetBackendKeysArgs)(nil)).Elem()
}

// A collection of values returned by getBackendKeys.
type GetBackendKeysResultOutput struct{ *pulumi.OutputState }

func (GetBackendKeysResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetBackendKeysResult)(nil)).Elem()
}

func (o GetBackendKeysResultOutput) ToGetBackendKeysResultOutput() GetBackendKeysResultOutput {
	return o
}

func (o GetBackendKeysResultOutput) ToGetBackendKeysResultOutputWithContext(ctx context.Context) GetBackendKeysResultOutput {
	return o
}

func (o GetBackendKeysResultOutput) Backend() pulumi.StringOutput {
	return o.ApplyT(func(v GetBackendKeysResult) string { return v.Backend }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetBackendKeysResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetBackendKeysResult) string { return v.Id }).(pulumi.StringOutput)
}

// Map of key strings read from Vault.
func (o GetBackendKeysResultOutput) KeyInfo() pulumi.StringMapOutput {
	return o.ApplyT(func(v GetBackendKeysResult) map[string]string { return v.KeyInfo }).(pulumi.StringMapOutput)
}

// JSON-encoded key data read from Vault.
func (o GetBackendKeysResultOutput) KeyInfoJson() pulumi.StringOutput {
	return o.ApplyT(func(v GetBackendKeysResult) string { return v.KeyInfoJson }).(pulumi.StringOutput)
}

// Keys used under the backend path.
func (o GetBackendKeysResultOutput) Keys() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetBackendKeysResult) []string { return v.Keys }).(pulumi.StringArrayOutput)
}

func (o GetBackendKeysResultOutput) Namespace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetBackendKeysResult) *string { return v.Namespace }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetBackendKeysResultOutput{})
}
