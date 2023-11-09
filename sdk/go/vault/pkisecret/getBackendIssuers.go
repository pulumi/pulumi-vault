// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package pkisecret

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-vault/sdk/v5/go/vault/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-vault/sdk/v5/go/vault"
//	"github.com/pulumi/pulumi-vault/sdk/v5/go/vault/pkiSecret"
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
//				IssuerName: pulumi.String("example"),
//			})
//			if err != nil {
//				return err
//			}
//			_ = pkiSecret.GetBackendIssuersOutput(ctx, pkisecret.GetBackendIssuersOutputArgs{
//				Backend: root.Backend,
//			}, nil)
//			return nil
//		})
//	}
//
// ```
func GetBackendIssuers(ctx *pulumi.Context, args *GetBackendIssuersArgs, opts ...pulumi.InvokeOption) (*GetBackendIssuersResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetBackendIssuersResult
	err := ctx.Invoke("vault:pkiSecret/getBackendIssuers:getBackendIssuers", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getBackendIssuers.
type GetBackendIssuersArgs struct {
	// The path to the PKI secret backend to
	// read the issuers from, with no leading or trailing `/`s.
	Backend string `pulumi:"backend"`
	// The namespace of the target resource.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
	// *Available only for Vault Enterprise*.
	Namespace *string `pulumi:"namespace"`
}

// A collection of values returned by getBackendIssuers.
type GetBackendIssuersResult struct {
	Backend string `pulumi:"backend"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Map of issuer strings read from Vault.
	KeyInfo map[string]interface{} `pulumi:"keyInfo"`
	// JSON-encoded issuer data read from Vault.
	KeyInfoJson string `pulumi:"keyInfoJson"`
	// Keys used by issuers under the backend path.
	Keys      []string `pulumi:"keys"`
	Namespace *string  `pulumi:"namespace"`
}

func GetBackendIssuersOutput(ctx *pulumi.Context, args GetBackendIssuersOutputArgs, opts ...pulumi.InvokeOption) GetBackendIssuersResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetBackendIssuersResult, error) {
			args := v.(GetBackendIssuersArgs)
			r, err := GetBackendIssuers(ctx, &args, opts...)
			var s GetBackendIssuersResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetBackendIssuersResultOutput)
}

// A collection of arguments for invoking getBackendIssuers.
type GetBackendIssuersOutputArgs struct {
	// The path to the PKI secret backend to
	// read the issuers from, with no leading or trailing `/`s.
	Backend pulumi.StringInput `pulumi:"backend"`
	// The namespace of the target resource.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
	// *Available only for Vault Enterprise*.
	Namespace pulumi.StringPtrInput `pulumi:"namespace"`
}

func (GetBackendIssuersOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetBackendIssuersArgs)(nil)).Elem()
}

// A collection of values returned by getBackendIssuers.
type GetBackendIssuersResultOutput struct{ *pulumi.OutputState }

func (GetBackendIssuersResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetBackendIssuersResult)(nil)).Elem()
}

func (o GetBackendIssuersResultOutput) ToGetBackendIssuersResultOutput() GetBackendIssuersResultOutput {
	return o
}

func (o GetBackendIssuersResultOutput) ToGetBackendIssuersResultOutputWithContext(ctx context.Context) GetBackendIssuersResultOutput {
	return o
}

func (o GetBackendIssuersResultOutput) ToOutput(ctx context.Context) pulumix.Output[GetBackendIssuersResult] {
	return pulumix.Output[GetBackendIssuersResult]{
		OutputState: o.OutputState,
	}
}

func (o GetBackendIssuersResultOutput) Backend() pulumi.StringOutput {
	return o.ApplyT(func(v GetBackendIssuersResult) string { return v.Backend }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetBackendIssuersResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetBackendIssuersResult) string { return v.Id }).(pulumi.StringOutput)
}

// Map of issuer strings read from Vault.
func (o GetBackendIssuersResultOutput) KeyInfo() pulumi.MapOutput {
	return o.ApplyT(func(v GetBackendIssuersResult) map[string]interface{} { return v.KeyInfo }).(pulumi.MapOutput)
}

// JSON-encoded issuer data read from Vault.
func (o GetBackendIssuersResultOutput) KeyInfoJson() pulumi.StringOutput {
	return o.ApplyT(func(v GetBackendIssuersResult) string { return v.KeyInfoJson }).(pulumi.StringOutput)
}

// Keys used by issuers under the backend path.
func (o GetBackendIssuersResultOutput) Keys() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetBackendIssuersResult) []string { return v.Keys }).(pulumi.StringArrayOutput)
}

func (o GetBackendIssuersResultOutput) Namespace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetBackendIssuersResult) *string { return v.Namespace }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetBackendIssuersResultOutput{})
}
