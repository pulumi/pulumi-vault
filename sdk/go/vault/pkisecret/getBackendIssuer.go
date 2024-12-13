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
//				IssuerName: pulumi.String("example"),
//			})
//			if err != nil {
//				return err
//			}
//			_ = root.IssuerId.ApplyT(func(issuerId string) (pkisecret.GetBackendIssuerResult, error) {
//				return pkisecret.GetBackendIssuerResult(interface{}(pkiSecret.GetBackendIssuerOutput(ctx, pkisecret.GetBackendIssuerOutputArgs{
//					Backend:   root.Path,
//					IssuerRef: issuerId,
//				}, nil))), nil
//			}).(pkisecret.GetBackendIssuerResultOutput)
//			return nil
//		})
//	}
//
// ```
func GetBackendIssuer(ctx *pulumi.Context, args *GetBackendIssuerArgs, opts ...pulumi.InvokeOption) (*GetBackendIssuerResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetBackendIssuerResult
	err := ctx.Invoke("vault:pkiSecret/getBackendIssuer:getBackendIssuer", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getBackendIssuer.
type GetBackendIssuerArgs struct {
	// The path to the PKI secret backend to
	// read the issuer from, with no leading or trailing `/`s.
	Backend string `pulumi:"backend"`
	// Reference to an existing issuer.
	IssuerRef string `pulumi:"issuerRef"`
	// The namespace of the target resource.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
	// *Available only for Vault Enterprise*.
	Namespace *string `pulumi:"namespace"`
}

// A collection of values returned by getBackendIssuer.
type GetBackendIssuerResult struct {
	Backend string `pulumi:"backend"`
	// The CA chain as a list of format specific certificates.
	CaChains []string `pulumi:"caChains"`
	// Certificate associated with this issuer.
	Certificate string `pulumi:"certificate"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// ID of the issuer.
	IssuerId string `pulumi:"issuerId"`
	// Name of the issuer.
	IssuerName string `pulumi:"issuerName"`
	IssuerRef  string `pulumi:"issuerRef"`
	// ID of the key used by the issuer.
	KeyId string `pulumi:"keyId"`
	// Behavior of a leaf's NotAfter field during issuance.
	LeafNotAfterBehavior string `pulumi:"leafNotAfterBehavior"`
	// Chain of issuer references to build this issuer's computed
	// CAChain field from, when non-empty.
	ManualChains []string `pulumi:"manualChains"`
	Namespace    *string  `pulumi:"namespace"`
	// Allowed usages for this issuer.
	Usage string `pulumi:"usage"`
}

func GetBackendIssuerOutput(ctx *pulumi.Context, args GetBackendIssuerOutputArgs, opts ...pulumi.InvokeOption) GetBackendIssuerResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (GetBackendIssuerResultOutput, error) {
			args := v.(GetBackendIssuerArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("vault:pkiSecret/getBackendIssuer:getBackendIssuer", args, GetBackendIssuerResultOutput{}, options).(GetBackendIssuerResultOutput), nil
		}).(GetBackendIssuerResultOutput)
}

// A collection of arguments for invoking getBackendIssuer.
type GetBackendIssuerOutputArgs struct {
	// The path to the PKI secret backend to
	// read the issuer from, with no leading or trailing `/`s.
	Backend pulumi.StringInput `pulumi:"backend"`
	// Reference to an existing issuer.
	IssuerRef pulumi.StringInput `pulumi:"issuerRef"`
	// The namespace of the target resource.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
	// *Available only for Vault Enterprise*.
	Namespace pulumi.StringPtrInput `pulumi:"namespace"`
}

func (GetBackendIssuerOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetBackendIssuerArgs)(nil)).Elem()
}

// A collection of values returned by getBackendIssuer.
type GetBackendIssuerResultOutput struct{ *pulumi.OutputState }

func (GetBackendIssuerResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetBackendIssuerResult)(nil)).Elem()
}

func (o GetBackendIssuerResultOutput) ToGetBackendIssuerResultOutput() GetBackendIssuerResultOutput {
	return o
}

func (o GetBackendIssuerResultOutput) ToGetBackendIssuerResultOutputWithContext(ctx context.Context) GetBackendIssuerResultOutput {
	return o
}

func (o GetBackendIssuerResultOutput) Backend() pulumi.StringOutput {
	return o.ApplyT(func(v GetBackendIssuerResult) string { return v.Backend }).(pulumi.StringOutput)
}

// The CA chain as a list of format specific certificates.
func (o GetBackendIssuerResultOutput) CaChains() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetBackendIssuerResult) []string { return v.CaChains }).(pulumi.StringArrayOutput)
}

// Certificate associated with this issuer.
func (o GetBackendIssuerResultOutput) Certificate() pulumi.StringOutput {
	return o.ApplyT(func(v GetBackendIssuerResult) string { return v.Certificate }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetBackendIssuerResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetBackendIssuerResult) string { return v.Id }).(pulumi.StringOutput)
}

// ID of the issuer.
func (o GetBackendIssuerResultOutput) IssuerId() pulumi.StringOutput {
	return o.ApplyT(func(v GetBackendIssuerResult) string { return v.IssuerId }).(pulumi.StringOutput)
}

// Name of the issuer.
func (o GetBackendIssuerResultOutput) IssuerName() pulumi.StringOutput {
	return o.ApplyT(func(v GetBackendIssuerResult) string { return v.IssuerName }).(pulumi.StringOutput)
}

func (o GetBackendIssuerResultOutput) IssuerRef() pulumi.StringOutput {
	return o.ApplyT(func(v GetBackendIssuerResult) string { return v.IssuerRef }).(pulumi.StringOutput)
}

// ID of the key used by the issuer.
func (o GetBackendIssuerResultOutput) KeyId() pulumi.StringOutput {
	return o.ApplyT(func(v GetBackendIssuerResult) string { return v.KeyId }).(pulumi.StringOutput)
}

// Behavior of a leaf's NotAfter field during issuance.
func (o GetBackendIssuerResultOutput) LeafNotAfterBehavior() pulumi.StringOutput {
	return o.ApplyT(func(v GetBackendIssuerResult) string { return v.LeafNotAfterBehavior }).(pulumi.StringOutput)
}

// Chain of issuer references to build this issuer's computed
// CAChain field from, when non-empty.
func (o GetBackendIssuerResultOutput) ManualChains() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetBackendIssuerResult) []string { return v.ManualChains }).(pulumi.StringArrayOutput)
}

func (o GetBackendIssuerResultOutput) Namespace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetBackendIssuerResult) *string { return v.Namespace }).(pulumi.StringPtrOutput)
}

// Allowed usages for this issuer.
func (o GetBackendIssuerResultOutput) Usage() pulumi.StringOutput {
	return o.ApplyT(func(v GetBackendIssuerResult) string { return v.Usage }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetBackendIssuerResultOutput{})
}
