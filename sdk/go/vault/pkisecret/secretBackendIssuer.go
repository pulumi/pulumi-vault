// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package pkisecret

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-vault/sdk/v5/go/vault/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
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
//				Path:                   pulumi.String("pki"),
//				Type:                   pulumi.String("pki"),
//				DefaultLeaseTtlSeconds: pulumi.Int(3600),
//				MaxLeaseTtlSeconds:     pulumi.Int(86400),
//			})
//			if err != nil {
//				return err
//			}
//			root, err := pkiSecret.NewSecretBackendRootCert(ctx, "root", &pkiSecret.SecretBackendRootCertArgs{
//				Backend:    pki.Path,
//				Type:       pulumi.String("internal"),
//				CommonName: pulumi.String("test"),
//				Ttl:        pulumi.String("86400"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = pkiSecret.NewSecretBackendIssuer(ctx, "example", &pkiSecret.SecretBackendIssuerArgs{
//				Backend:    root.Backend,
//				IssuerRef:  root.IssuerId,
//				IssuerName: pulumi.String("example-issuer"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// ## Import
//
// PKI secret backend issuer can be imported using the `id`, e.g.
//
// ```sh
//
//	$ pulumi import vault:pkiSecret/secretBackendIssuer:SecretBackendIssuer example pki/issuer/bf9b0d48-d0dd-652c-30be-77d04fc7e94d
//
// ```
type SecretBackendIssuer struct {
	pulumi.CustomResourceState

	// The path the PKI secret backend is mounted at, with no
	// leading or trailing `/`s.
	Backend pulumi.StringOutput `pulumi:"backend"`
	// Specifies the URL values for the CRL
	// Distribution Points field.
	CrlDistributionPoints pulumi.StringArrayOutput `pulumi:"crlDistributionPoints"`
	// Specifies that the AIA URL values should
	// be templated.
	EnableAiaUrlTemplating pulumi.BoolPtrOutput `pulumi:"enableAiaUrlTemplating"`
	// ID of the issuer.
	IssuerId pulumi.StringOutput `pulumi:"issuerId"`
	// Name of the issuer.
	IssuerName pulumi.StringPtrOutput `pulumi:"issuerName"`
	// Reference to an existing issuer.
	IssuerRef pulumi.StringOutput `pulumi:"issuerRef"`
	// Specifies the URL values for the Issuing
	// Certificate field.
	IssuingCertificates pulumi.StringArrayOutput `pulumi:"issuingCertificates"`
	// Behavior of a leaf's NotAfter field during
	// issuance.
	LeafNotAfterBehavior pulumi.StringOutput `pulumi:"leafNotAfterBehavior"`
	// Chain of issuer references to build this issuer's
	// computed CAChain field from, when non-empty.
	ManualChains pulumi.StringArrayOutput `pulumi:"manualChains"`
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
	// *Available only for Vault Enterprise*.
	Namespace pulumi.StringPtrOutput `pulumi:"namespace"`
	// Specifies the URL values for the OCSP Servers field.
	OcspServers pulumi.StringArrayOutput `pulumi:"ocspServers"`
	// Which signature algorithm to use
	// when building CRLs.
	RevocationSignatureAlgorithm pulumi.StringOutput `pulumi:"revocationSignatureAlgorithm"`
	// Allowed usages for this issuer.
	Usage pulumi.StringOutput `pulumi:"usage"`
}

// NewSecretBackendIssuer registers a new resource with the given unique name, arguments, and options.
func NewSecretBackendIssuer(ctx *pulumi.Context,
	name string, args *SecretBackendIssuerArgs, opts ...pulumi.ResourceOption) (*SecretBackendIssuer, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Backend == nil {
		return nil, errors.New("invalid value for required argument 'Backend'")
	}
	if args.IssuerRef == nil {
		return nil, errors.New("invalid value for required argument 'IssuerRef'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource SecretBackendIssuer
	err := ctx.RegisterResource("vault:pkiSecret/secretBackendIssuer:SecretBackendIssuer", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSecretBackendIssuer gets an existing SecretBackendIssuer resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSecretBackendIssuer(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SecretBackendIssuerState, opts ...pulumi.ResourceOption) (*SecretBackendIssuer, error) {
	var resource SecretBackendIssuer
	err := ctx.ReadResource("vault:pkiSecret/secretBackendIssuer:SecretBackendIssuer", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SecretBackendIssuer resources.
type secretBackendIssuerState struct {
	// The path the PKI secret backend is mounted at, with no
	// leading or trailing `/`s.
	Backend *string `pulumi:"backend"`
	// Specifies the URL values for the CRL
	// Distribution Points field.
	CrlDistributionPoints []string `pulumi:"crlDistributionPoints"`
	// Specifies that the AIA URL values should
	// be templated.
	EnableAiaUrlTemplating *bool `pulumi:"enableAiaUrlTemplating"`
	// ID of the issuer.
	IssuerId *string `pulumi:"issuerId"`
	// Name of the issuer.
	IssuerName *string `pulumi:"issuerName"`
	// Reference to an existing issuer.
	IssuerRef *string `pulumi:"issuerRef"`
	// Specifies the URL values for the Issuing
	// Certificate field.
	IssuingCertificates []string `pulumi:"issuingCertificates"`
	// Behavior of a leaf's NotAfter field during
	// issuance.
	LeafNotAfterBehavior *string `pulumi:"leafNotAfterBehavior"`
	// Chain of issuer references to build this issuer's
	// computed CAChain field from, when non-empty.
	ManualChains []string `pulumi:"manualChains"`
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
	// *Available only for Vault Enterprise*.
	Namespace *string `pulumi:"namespace"`
	// Specifies the URL values for the OCSP Servers field.
	OcspServers []string `pulumi:"ocspServers"`
	// Which signature algorithm to use
	// when building CRLs.
	RevocationSignatureAlgorithm *string `pulumi:"revocationSignatureAlgorithm"`
	// Allowed usages for this issuer.
	Usage *string `pulumi:"usage"`
}

type SecretBackendIssuerState struct {
	// The path the PKI secret backend is mounted at, with no
	// leading or trailing `/`s.
	Backend pulumi.StringPtrInput
	// Specifies the URL values for the CRL
	// Distribution Points field.
	CrlDistributionPoints pulumi.StringArrayInput
	// Specifies that the AIA URL values should
	// be templated.
	EnableAiaUrlTemplating pulumi.BoolPtrInput
	// ID of the issuer.
	IssuerId pulumi.StringPtrInput
	// Name of the issuer.
	IssuerName pulumi.StringPtrInput
	// Reference to an existing issuer.
	IssuerRef pulumi.StringPtrInput
	// Specifies the URL values for the Issuing
	// Certificate field.
	IssuingCertificates pulumi.StringArrayInput
	// Behavior of a leaf's NotAfter field during
	// issuance.
	LeafNotAfterBehavior pulumi.StringPtrInput
	// Chain of issuer references to build this issuer's
	// computed CAChain field from, when non-empty.
	ManualChains pulumi.StringArrayInput
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
	// *Available only for Vault Enterprise*.
	Namespace pulumi.StringPtrInput
	// Specifies the URL values for the OCSP Servers field.
	OcspServers pulumi.StringArrayInput
	// Which signature algorithm to use
	// when building CRLs.
	RevocationSignatureAlgorithm pulumi.StringPtrInput
	// Allowed usages for this issuer.
	Usage pulumi.StringPtrInput
}

func (SecretBackendIssuerState) ElementType() reflect.Type {
	return reflect.TypeOf((*secretBackendIssuerState)(nil)).Elem()
}

type secretBackendIssuerArgs struct {
	// The path the PKI secret backend is mounted at, with no
	// leading or trailing `/`s.
	Backend string `pulumi:"backend"`
	// Specifies the URL values for the CRL
	// Distribution Points field.
	CrlDistributionPoints []string `pulumi:"crlDistributionPoints"`
	// Specifies that the AIA URL values should
	// be templated.
	EnableAiaUrlTemplating *bool `pulumi:"enableAiaUrlTemplating"`
	// Name of the issuer.
	IssuerName *string `pulumi:"issuerName"`
	// Reference to an existing issuer.
	IssuerRef string `pulumi:"issuerRef"`
	// Specifies the URL values for the Issuing
	// Certificate field.
	IssuingCertificates []string `pulumi:"issuingCertificates"`
	// Behavior of a leaf's NotAfter field during
	// issuance.
	LeafNotAfterBehavior *string `pulumi:"leafNotAfterBehavior"`
	// Chain of issuer references to build this issuer's
	// computed CAChain field from, when non-empty.
	ManualChains []string `pulumi:"manualChains"`
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
	// *Available only for Vault Enterprise*.
	Namespace *string `pulumi:"namespace"`
	// Specifies the URL values for the OCSP Servers field.
	OcspServers []string `pulumi:"ocspServers"`
	// Which signature algorithm to use
	// when building CRLs.
	RevocationSignatureAlgorithm *string `pulumi:"revocationSignatureAlgorithm"`
	// Allowed usages for this issuer.
	Usage *string `pulumi:"usage"`
}

// The set of arguments for constructing a SecretBackendIssuer resource.
type SecretBackendIssuerArgs struct {
	// The path the PKI secret backend is mounted at, with no
	// leading or trailing `/`s.
	Backend pulumi.StringInput
	// Specifies the URL values for the CRL
	// Distribution Points field.
	CrlDistributionPoints pulumi.StringArrayInput
	// Specifies that the AIA URL values should
	// be templated.
	EnableAiaUrlTemplating pulumi.BoolPtrInput
	// Name of the issuer.
	IssuerName pulumi.StringPtrInput
	// Reference to an existing issuer.
	IssuerRef pulumi.StringInput
	// Specifies the URL values for the Issuing
	// Certificate field.
	IssuingCertificates pulumi.StringArrayInput
	// Behavior of a leaf's NotAfter field during
	// issuance.
	LeafNotAfterBehavior pulumi.StringPtrInput
	// Chain of issuer references to build this issuer's
	// computed CAChain field from, when non-empty.
	ManualChains pulumi.StringArrayInput
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
	// *Available only for Vault Enterprise*.
	Namespace pulumi.StringPtrInput
	// Specifies the URL values for the OCSP Servers field.
	OcspServers pulumi.StringArrayInput
	// Which signature algorithm to use
	// when building CRLs.
	RevocationSignatureAlgorithm pulumi.StringPtrInput
	// Allowed usages for this issuer.
	Usage pulumi.StringPtrInput
}

func (SecretBackendIssuerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*secretBackendIssuerArgs)(nil)).Elem()
}

type SecretBackendIssuerInput interface {
	pulumi.Input

	ToSecretBackendIssuerOutput() SecretBackendIssuerOutput
	ToSecretBackendIssuerOutputWithContext(ctx context.Context) SecretBackendIssuerOutput
}

func (*SecretBackendIssuer) ElementType() reflect.Type {
	return reflect.TypeOf((**SecretBackendIssuer)(nil)).Elem()
}

func (i *SecretBackendIssuer) ToSecretBackendIssuerOutput() SecretBackendIssuerOutput {
	return i.ToSecretBackendIssuerOutputWithContext(context.Background())
}

func (i *SecretBackendIssuer) ToSecretBackendIssuerOutputWithContext(ctx context.Context) SecretBackendIssuerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecretBackendIssuerOutput)
}

// SecretBackendIssuerArrayInput is an input type that accepts SecretBackendIssuerArray and SecretBackendIssuerArrayOutput values.
// You can construct a concrete instance of `SecretBackendIssuerArrayInput` via:
//
//	SecretBackendIssuerArray{ SecretBackendIssuerArgs{...} }
type SecretBackendIssuerArrayInput interface {
	pulumi.Input

	ToSecretBackendIssuerArrayOutput() SecretBackendIssuerArrayOutput
	ToSecretBackendIssuerArrayOutputWithContext(context.Context) SecretBackendIssuerArrayOutput
}

type SecretBackendIssuerArray []SecretBackendIssuerInput

func (SecretBackendIssuerArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SecretBackendIssuer)(nil)).Elem()
}

func (i SecretBackendIssuerArray) ToSecretBackendIssuerArrayOutput() SecretBackendIssuerArrayOutput {
	return i.ToSecretBackendIssuerArrayOutputWithContext(context.Background())
}

func (i SecretBackendIssuerArray) ToSecretBackendIssuerArrayOutputWithContext(ctx context.Context) SecretBackendIssuerArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecretBackendIssuerArrayOutput)
}

// SecretBackendIssuerMapInput is an input type that accepts SecretBackendIssuerMap and SecretBackendIssuerMapOutput values.
// You can construct a concrete instance of `SecretBackendIssuerMapInput` via:
//
//	SecretBackendIssuerMap{ "key": SecretBackendIssuerArgs{...} }
type SecretBackendIssuerMapInput interface {
	pulumi.Input

	ToSecretBackendIssuerMapOutput() SecretBackendIssuerMapOutput
	ToSecretBackendIssuerMapOutputWithContext(context.Context) SecretBackendIssuerMapOutput
}

type SecretBackendIssuerMap map[string]SecretBackendIssuerInput

func (SecretBackendIssuerMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SecretBackendIssuer)(nil)).Elem()
}

func (i SecretBackendIssuerMap) ToSecretBackendIssuerMapOutput() SecretBackendIssuerMapOutput {
	return i.ToSecretBackendIssuerMapOutputWithContext(context.Background())
}

func (i SecretBackendIssuerMap) ToSecretBackendIssuerMapOutputWithContext(ctx context.Context) SecretBackendIssuerMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecretBackendIssuerMapOutput)
}

type SecretBackendIssuerOutput struct{ *pulumi.OutputState }

func (SecretBackendIssuerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SecretBackendIssuer)(nil)).Elem()
}

func (o SecretBackendIssuerOutput) ToSecretBackendIssuerOutput() SecretBackendIssuerOutput {
	return o
}

func (o SecretBackendIssuerOutput) ToSecretBackendIssuerOutputWithContext(ctx context.Context) SecretBackendIssuerOutput {
	return o
}

// The path the PKI secret backend is mounted at, with no
// leading or trailing `/`s.
func (o SecretBackendIssuerOutput) Backend() pulumi.StringOutput {
	return o.ApplyT(func(v *SecretBackendIssuer) pulumi.StringOutput { return v.Backend }).(pulumi.StringOutput)
}

// Specifies the URL values for the CRL
// Distribution Points field.
func (o SecretBackendIssuerOutput) CrlDistributionPoints() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *SecretBackendIssuer) pulumi.StringArrayOutput { return v.CrlDistributionPoints }).(pulumi.StringArrayOutput)
}

// Specifies that the AIA URL values should
// be templated.
func (o SecretBackendIssuerOutput) EnableAiaUrlTemplating() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SecretBackendIssuer) pulumi.BoolPtrOutput { return v.EnableAiaUrlTemplating }).(pulumi.BoolPtrOutput)
}

// ID of the issuer.
func (o SecretBackendIssuerOutput) IssuerId() pulumi.StringOutput {
	return o.ApplyT(func(v *SecretBackendIssuer) pulumi.StringOutput { return v.IssuerId }).(pulumi.StringOutput)
}

// Name of the issuer.
func (o SecretBackendIssuerOutput) IssuerName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecretBackendIssuer) pulumi.StringPtrOutput { return v.IssuerName }).(pulumi.StringPtrOutput)
}

// Reference to an existing issuer.
func (o SecretBackendIssuerOutput) IssuerRef() pulumi.StringOutput {
	return o.ApplyT(func(v *SecretBackendIssuer) pulumi.StringOutput { return v.IssuerRef }).(pulumi.StringOutput)
}

// Specifies the URL values for the Issuing
// Certificate field.
func (o SecretBackendIssuerOutput) IssuingCertificates() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *SecretBackendIssuer) pulumi.StringArrayOutput { return v.IssuingCertificates }).(pulumi.StringArrayOutput)
}

// Behavior of a leaf's NotAfter field during
// issuance.
func (o SecretBackendIssuerOutput) LeafNotAfterBehavior() pulumi.StringOutput {
	return o.ApplyT(func(v *SecretBackendIssuer) pulumi.StringOutput { return v.LeafNotAfterBehavior }).(pulumi.StringOutput)
}

// Chain of issuer references to build this issuer's
// computed CAChain field from, when non-empty.
func (o SecretBackendIssuerOutput) ManualChains() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *SecretBackendIssuer) pulumi.StringArrayOutput { return v.ManualChains }).(pulumi.StringArrayOutput)
}

// The namespace to provision the resource in.
// The value should not contain leading or trailing forward slashes.
// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
// *Available only for Vault Enterprise*.
func (o SecretBackendIssuerOutput) Namespace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecretBackendIssuer) pulumi.StringPtrOutput { return v.Namespace }).(pulumi.StringPtrOutput)
}

// Specifies the URL values for the OCSP Servers field.
func (o SecretBackendIssuerOutput) OcspServers() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *SecretBackendIssuer) pulumi.StringArrayOutput { return v.OcspServers }).(pulumi.StringArrayOutput)
}

// Which signature algorithm to use
// when building CRLs.
func (o SecretBackendIssuerOutput) RevocationSignatureAlgorithm() pulumi.StringOutput {
	return o.ApplyT(func(v *SecretBackendIssuer) pulumi.StringOutput { return v.RevocationSignatureAlgorithm }).(pulumi.StringOutput)
}

// Allowed usages for this issuer.
func (o SecretBackendIssuerOutput) Usage() pulumi.StringOutput {
	return o.ApplyT(func(v *SecretBackendIssuer) pulumi.StringOutput { return v.Usage }).(pulumi.StringOutput)
}

type SecretBackendIssuerArrayOutput struct{ *pulumi.OutputState }

func (SecretBackendIssuerArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SecretBackendIssuer)(nil)).Elem()
}

func (o SecretBackendIssuerArrayOutput) ToSecretBackendIssuerArrayOutput() SecretBackendIssuerArrayOutput {
	return o
}

func (o SecretBackendIssuerArrayOutput) ToSecretBackendIssuerArrayOutputWithContext(ctx context.Context) SecretBackendIssuerArrayOutput {
	return o
}

func (o SecretBackendIssuerArrayOutput) Index(i pulumi.IntInput) SecretBackendIssuerOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SecretBackendIssuer {
		return vs[0].([]*SecretBackendIssuer)[vs[1].(int)]
	}).(SecretBackendIssuerOutput)
}

type SecretBackendIssuerMapOutput struct{ *pulumi.OutputState }

func (SecretBackendIssuerMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SecretBackendIssuer)(nil)).Elem()
}

func (o SecretBackendIssuerMapOutput) ToSecretBackendIssuerMapOutput() SecretBackendIssuerMapOutput {
	return o
}

func (o SecretBackendIssuerMapOutput) ToSecretBackendIssuerMapOutputWithContext(ctx context.Context) SecretBackendIssuerMapOutput {
	return o
}

func (o SecretBackendIssuerMapOutput) MapIndex(k pulumi.StringInput) SecretBackendIssuerOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SecretBackendIssuer {
		return vs[0].(map[string]*SecretBackendIssuer)[vs[1].(string)]
	}).(SecretBackendIssuerOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SecretBackendIssuerInput)(nil)).Elem(), &SecretBackendIssuer{})
	pulumi.RegisterInputType(reflect.TypeOf((*SecretBackendIssuerArrayInput)(nil)).Elem(), SecretBackendIssuerArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SecretBackendIssuerMapInput)(nil)).Elem(), SecretBackendIssuerMap{})
	pulumi.RegisterOutputType(SecretBackendIssuerOutput{})
	pulumi.RegisterOutputType(SecretBackendIssuerArrayOutput{})
	pulumi.RegisterOutputType(SecretBackendIssuerMapOutput{})
}
