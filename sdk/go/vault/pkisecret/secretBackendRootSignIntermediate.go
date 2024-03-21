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

// Creates PKI certificate.
//
// ## Example Usage
//
// <!--Start PulumiCodeChooser -->
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-vault/sdk/v5/go/vault/pkiSecret"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := pkiSecret.NewSecretBackendRootSignIntermediate(ctx, "root", &pkiSecret.SecretBackendRootSignIntermediateArgs{
//				Backend:           pulumi.Any(vault_mount.Root.Path),
//				Csr:               pulumi.Any(vault_pki_secret_backend_intermediate_cert_request.Intermediate.Csr),
//				CommonName:        pulumi.String("Intermediate CA"),
//				ExcludeCnFromSans: pulumi.Bool(true),
//				Ou:                pulumi.String("My OU"),
//				Organization:      pulumi.String("My organization"),
//			}, pulumi.DependsOn([]pulumi.Resource{
//				vault_pki_secret_backend_intermediate_cert_request.Intermediate,
//			}))
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// <!--End PulumiCodeChooser -->
//
// ## Deprecations
//
// * `serial` - Use `serialNumber` instead.
type SecretBackendRootSignIntermediate struct {
	pulumi.CustomResourceState

	// List of alternative names
	AltNames pulumi.StringArrayOutput `pulumi:"altNames"`
	// The PKI secret backend the resource belongs to.
	Backend pulumi.StringOutput `pulumi:"backend"`
	// A list of the issuing and intermediate CA certificates in the `format` specified.
	CaChains pulumi.StringArrayOutput `pulumi:"caChains"`
	// The intermediate CA certificate in the `format` specified.
	Certificate pulumi.StringOutput `pulumi:"certificate"`
	// The concatenation of the intermediate CA and the issuing CA certificates (PEM encoded).
	// Requires the `format` to be set to any of: pem, pem_bundle. The value will be empty for all other formats.
	CertificateBundle pulumi.StringOutput `pulumi:"certificateBundle"`
	// CN of intermediate to create
	CommonName pulumi.StringOutput `pulumi:"commonName"`
	// The country
	Country pulumi.StringPtrOutput `pulumi:"country"`
	// The CSR
	Csr pulumi.StringOutput `pulumi:"csr"`
	// Flag to exclude CN from SANs
	ExcludeCnFromSans pulumi.BoolPtrOutput `pulumi:"excludeCnFromSans"`
	// The format of data
	Format pulumi.StringPtrOutput `pulumi:"format"`
	// List of alternative IPs
	IpSans pulumi.StringArrayOutput `pulumi:"ipSans"`
	// Specifies the default issuer of this request. May
	// be the value `default`, a name, or an issuer ID. Use ACLs to prevent access to
	// the `/pki/issuer/:issuer_ref/{issue,sign}/:name` paths to prevent users
	// overriding the role's `issuerRef` value.
	IssuerRef pulumi.StringPtrOutput `pulumi:"issuerRef"`
	// The issuing CA certificate in the `format` specified.
	IssuingCa pulumi.StringOutput `pulumi:"issuingCa"`
	// The locality
	Locality pulumi.StringPtrOutput `pulumi:"locality"`
	// The maximum path length to encode in the generated certificate
	MaxPathLength pulumi.IntPtrOutput `pulumi:"maxPathLength"`
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
	// *Available only for Vault Enterprise*.
	Namespace pulumi.StringPtrOutput `pulumi:"namespace"`
	// The organization
	Organization pulumi.StringPtrOutput `pulumi:"organization"`
	// List of other SANs
	OtherSans pulumi.StringArrayOutput `pulumi:"otherSans"`
	// The organization unit
	Ou pulumi.StringPtrOutput `pulumi:"ou"`
	// List of domains for which certificates are allowed to be issued
	PermittedDnsDomains pulumi.StringArrayOutput `pulumi:"permittedDnsDomains"`
	// The postal code
	PostalCode pulumi.StringPtrOutput `pulumi:"postalCode"`
	// The province
	Province pulumi.StringPtrOutput `pulumi:"province"`
	// If set to `true`, the certificate will be revoked on resource destruction.
	Revoke pulumi.BoolPtrOutput `pulumi:"revoke"`
	// The serial number.
	//
	// Deprecated: Use serialNumber instead
	Serial pulumi.StringOutput `pulumi:"serial"`
	// The certificate's serial number, hex formatted.
	SerialNumber pulumi.StringOutput `pulumi:"serialNumber"`
	// The street address
	StreetAddress pulumi.StringPtrOutput `pulumi:"streetAddress"`
	// Time to live
	Ttl pulumi.StringPtrOutput `pulumi:"ttl"`
	// List of alternative URIs
	UriSans pulumi.StringArrayOutput `pulumi:"uriSans"`
	// Preserve CSR values
	UseCsrValues pulumi.BoolPtrOutput `pulumi:"useCsrValues"`
}

// NewSecretBackendRootSignIntermediate registers a new resource with the given unique name, arguments, and options.
func NewSecretBackendRootSignIntermediate(ctx *pulumi.Context,
	name string, args *SecretBackendRootSignIntermediateArgs, opts ...pulumi.ResourceOption) (*SecretBackendRootSignIntermediate, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Backend == nil {
		return nil, errors.New("invalid value for required argument 'Backend'")
	}
	if args.CommonName == nil {
		return nil, errors.New("invalid value for required argument 'CommonName'")
	}
	if args.Csr == nil {
		return nil, errors.New("invalid value for required argument 'Csr'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource SecretBackendRootSignIntermediate
	err := ctx.RegisterResource("vault:pkiSecret/secretBackendRootSignIntermediate:SecretBackendRootSignIntermediate", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSecretBackendRootSignIntermediate gets an existing SecretBackendRootSignIntermediate resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSecretBackendRootSignIntermediate(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SecretBackendRootSignIntermediateState, opts ...pulumi.ResourceOption) (*SecretBackendRootSignIntermediate, error) {
	var resource SecretBackendRootSignIntermediate
	err := ctx.ReadResource("vault:pkiSecret/secretBackendRootSignIntermediate:SecretBackendRootSignIntermediate", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SecretBackendRootSignIntermediate resources.
type secretBackendRootSignIntermediateState struct {
	// List of alternative names
	AltNames []string `pulumi:"altNames"`
	// The PKI secret backend the resource belongs to.
	Backend *string `pulumi:"backend"`
	// A list of the issuing and intermediate CA certificates in the `format` specified.
	CaChains []string `pulumi:"caChains"`
	// The intermediate CA certificate in the `format` specified.
	Certificate *string `pulumi:"certificate"`
	// The concatenation of the intermediate CA and the issuing CA certificates (PEM encoded).
	// Requires the `format` to be set to any of: pem, pem_bundle. The value will be empty for all other formats.
	CertificateBundle *string `pulumi:"certificateBundle"`
	// CN of intermediate to create
	CommonName *string `pulumi:"commonName"`
	// The country
	Country *string `pulumi:"country"`
	// The CSR
	Csr *string `pulumi:"csr"`
	// Flag to exclude CN from SANs
	ExcludeCnFromSans *bool `pulumi:"excludeCnFromSans"`
	// The format of data
	Format *string `pulumi:"format"`
	// List of alternative IPs
	IpSans []string `pulumi:"ipSans"`
	// Specifies the default issuer of this request. May
	// be the value `default`, a name, or an issuer ID. Use ACLs to prevent access to
	// the `/pki/issuer/:issuer_ref/{issue,sign}/:name` paths to prevent users
	// overriding the role's `issuerRef` value.
	IssuerRef *string `pulumi:"issuerRef"`
	// The issuing CA certificate in the `format` specified.
	IssuingCa *string `pulumi:"issuingCa"`
	// The locality
	Locality *string `pulumi:"locality"`
	// The maximum path length to encode in the generated certificate
	MaxPathLength *int `pulumi:"maxPathLength"`
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
	// *Available only for Vault Enterprise*.
	Namespace *string `pulumi:"namespace"`
	// The organization
	Organization *string `pulumi:"organization"`
	// List of other SANs
	OtherSans []string `pulumi:"otherSans"`
	// The organization unit
	Ou *string `pulumi:"ou"`
	// List of domains for which certificates are allowed to be issued
	PermittedDnsDomains []string `pulumi:"permittedDnsDomains"`
	// The postal code
	PostalCode *string `pulumi:"postalCode"`
	// The province
	Province *string `pulumi:"province"`
	// If set to `true`, the certificate will be revoked on resource destruction.
	Revoke *bool `pulumi:"revoke"`
	// The serial number.
	//
	// Deprecated: Use serialNumber instead
	Serial *string `pulumi:"serial"`
	// The certificate's serial number, hex formatted.
	SerialNumber *string `pulumi:"serialNumber"`
	// The street address
	StreetAddress *string `pulumi:"streetAddress"`
	// Time to live
	Ttl *string `pulumi:"ttl"`
	// List of alternative URIs
	UriSans []string `pulumi:"uriSans"`
	// Preserve CSR values
	UseCsrValues *bool `pulumi:"useCsrValues"`
}

type SecretBackendRootSignIntermediateState struct {
	// List of alternative names
	AltNames pulumi.StringArrayInput
	// The PKI secret backend the resource belongs to.
	Backend pulumi.StringPtrInput
	// A list of the issuing and intermediate CA certificates in the `format` specified.
	CaChains pulumi.StringArrayInput
	// The intermediate CA certificate in the `format` specified.
	Certificate pulumi.StringPtrInput
	// The concatenation of the intermediate CA and the issuing CA certificates (PEM encoded).
	// Requires the `format` to be set to any of: pem, pem_bundle. The value will be empty for all other formats.
	CertificateBundle pulumi.StringPtrInput
	// CN of intermediate to create
	CommonName pulumi.StringPtrInput
	// The country
	Country pulumi.StringPtrInput
	// The CSR
	Csr pulumi.StringPtrInput
	// Flag to exclude CN from SANs
	ExcludeCnFromSans pulumi.BoolPtrInput
	// The format of data
	Format pulumi.StringPtrInput
	// List of alternative IPs
	IpSans pulumi.StringArrayInput
	// Specifies the default issuer of this request. May
	// be the value `default`, a name, or an issuer ID. Use ACLs to prevent access to
	// the `/pki/issuer/:issuer_ref/{issue,sign}/:name` paths to prevent users
	// overriding the role's `issuerRef` value.
	IssuerRef pulumi.StringPtrInput
	// The issuing CA certificate in the `format` specified.
	IssuingCa pulumi.StringPtrInput
	// The locality
	Locality pulumi.StringPtrInput
	// The maximum path length to encode in the generated certificate
	MaxPathLength pulumi.IntPtrInput
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
	// *Available only for Vault Enterprise*.
	Namespace pulumi.StringPtrInput
	// The organization
	Organization pulumi.StringPtrInput
	// List of other SANs
	OtherSans pulumi.StringArrayInput
	// The organization unit
	Ou pulumi.StringPtrInput
	// List of domains for which certificates are allowed to be issued
	PermittedDnsDomains pulumi.StringArrayInput
	// The postal code
	PostalCode pulumi.StringPtrInput
	// The province
	Province pulumi.StringPtrInput
	// If set to `true`, the certificate will be revoked on resource destruction.
	Revoke pulumi.BoolPtrInput
	// The serial number.
	//
	// Deprecated: Use serialNumber instead
	Serial pulumi.StringPtrInput
	// The certificate's serial number, hex formatted.
	SerialNumber pulumi.StringPtrInput
	// The street address
	StreetAddress pulumi.StringPtrInput
	// Time to live
	Ttl pulumi.StringPtrInput
	// List of alternative URIs
	UriSans pulumi.StringArrayInput
	// Preserve CSR values
	UseCsrValues pulumi.BoolPtrInput
}

func (SecretBackendRootSignIntermediateState) ElementType() reflect.Type {
	return reflect.TypeOf((*secretBackendRootSignIntermediateState)(nil)).Elem()
}

type secretBackendRootSignIntermediateArgs struct {
	// List of alternative names
	AltNames []string `pulumi:"altNames"`
	// The PKI secret backend the resource belongs to.
	Backend string `pulumi:"backend"`
	// CN of intermediate to create
	CommonName string `pulumi:"commonName"`
	// The country
	Country *string `pulumi:"country"`
	// The CSR
	Csr string `pulumi:"csr"`
	// Flag to exclude CN from SANs
	ExcludeCnFromSans *bool `pulumi:"excludeCnFromSans"`
	// The format of data
	Format *string `pulumi:"format"`
	// List of alternative IPs
	IpSans []string `pulumi:"ipSans"`
	// Specifies the default issuer of this request. May
	// be the value `default`, a name, or an issuer ID. Use ACLs to prevent access to
	// the `/pki/issuer/:issuer_ref/{issue,sign}/:name` paths to prevent users
	// overriding the role's `issuerRef` value.
	IssuerRef *string `pulumi:"issuerRef"`
	// The locality
	Locality *string `pulumi:"locality"`
	// The maximum path length to encode in the generated certificate
	MaxPathLength *int `pulumi:"maxPathLength"`
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
	// *Available only for Vault Enterprise*.
	Namespace *string `pulumi:"namespace"`
	// The organization
	Organization *string `pulumi:"organization"`
	// List of other SANs
	OtherSans []string `pulumi:"otherSans"`
	// The organization unit
	Ou *string `pulumi:"ou"`
	// List of domains for which certificates are allowed to be issued
	PermittedDnsDomains []string `pulumi:"permittedDnsDomains"`
	// The postal code
	PostalCode *string `pulumi:"postalCode"`
	// The province
	Province *string `pulumi:"province"`
	// If set to `true`, the certificate will be revoked on resource destruction.
	Revoke *bool `pulumi:"revoke"`
	// The street address
	StreetAddress *string `pulumi:"streetAddress"`
	// Time to live
	Ttl *string `pulumi:"ttl"`
	// List of alternative URIs
	UriSans []string `pulumi:"uriSans"`
	// Preserve CSR values
	UseCsrValues *bool `pulumi:"useCsrValues"`
}

// The set of arguments for constructing a SecretBackendRootSignIntermediate resource.
type SecretBackendRootSignIntermediateArgs struct {
	// List of alternative names
	AltNames pulumi.StringArrayInput
	// The PKI secret backend the resource belongs to.
	Backend pulumi.StringInput
	// CN of intermediate to create
	CommonName pulumi.StringInput
	// The country
	Country pulumi.StringPtrInput
	// The CSR
	Csr pulumi.StringInput
	// Flag to exclude CN from SANs
	ExcludeCnFromSans pulumi.BoolPtrInput
	// The format of data
	Format pulumi.StringPtrInput
	// List of alternative IPs
	IpSans pulumi.StringArrayInput
	// Specifies the default issuer of this request. May
	// be the value `default`, a name, or an issuer ID. Use ACLs to prevent access to
	// the `/pki/issuer/:issuer_ref/{issue,sign}/:name` paths to prevent users
	// overriding the role's `issuerRef` value.
	IssuerRef pulumi.StringPtrInput
	// The locality
	Locality pulumi.StringPtrInput
	// The maximum path length to encode in the generated certificate
	MaxPathLength pulumi.IntPtrInput
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
	// *Available only for Vault Enterprise*.
	Namespace pulumi.StringPtrInput
	// The organization
	Organization pulumi.StringPtrInput
	// List of other SANs
	OtherSans pulumi.StringArrayInput
	// The organization unit
	Ou pulumi.StringPtrInput
	// List of domains for which certificates are allowed to be issued
	PermittedDnsDomains pulumi.StringArrayInput
	// The postal code
	PostalCode pulumi.StringPtrInput
	// The province
	Province pulumi.StringPtrInput
	// If set to `true`, the certificate will be revoked on resource destruction.
	Revoke pulumi.BoolPtrInput
	// The street address
	StreetAddress pulumi.StringPtrInput
	// Time to live
	Ttl pulumi.StringPtrInput
	// List of alternative URIs
	UriSans pulumi.StringArrayInput
	// Preserve CSR values
	UseCsrValues pulumi.BoolPtrInput
}

func (SecretBackendRootSignIntermediateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*secretBackendRootSignIntermediateArgs)(nil)).Elem()
}

type SecretBackendRootSignIntermediateInput interface {
	pulumi.Input

	ToSecretBackendRootSignIntermediateOutput() SecretBackendRootSignIntermediateOutput
	ToSecretBackendRootSignIntermediateOutputWithContext(ctx context.Context) SecretBackendRootSignIntermediateOutput
}

func (*SecretBackendRootSignIntermediate) ElementType() reflect.Type {
	return reflect.TypeOf((**SecretBackendRootSignIntermediate)(nil)).Elem()
}

func (i *SecretBackendRootSignIntermediate) ToSecretBackendRootSignIntermediateOutput() SecretBackendRootSignIntermediateOutput {
	return i.ToSecretBackendRootSignIntermediateOutputWithContext(context.Background())
}

func (i *SecretBackendRootSignIntermediate) ToSecretBackendRootSignIntermediateOutputWithContext(ctx context.Context) SecretBackendRootSignIntermediateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecretBackendRootSignIntermediateOutput)
}

// SecretBackendRootSignIntermediateArrayInput is an input type that accepts SecretBackendRootSignIntermediateArray and SecretBackendRootSignIntermediateArrayOutput values.
// You can construct a concrete instance of `SecretBackendRootSignIntermediateArrayInput` via:
//
//	SecretBackendRootSignIntermediateArray{ SecretBackendRootSignIntermediateArgs{...} }
type SecretBackendRootSignIntermediateArrayInput interface {
	pulumi.Input

	ToSecretBackendRootSignIntermediateArrayOutput() SecretBackendRootSignIntermediateArrayOutput
	ToSecretBackendRootSignIntermediateArrayOutputWithContext(context.Context) SecretBackendRootSignIntermediateArrayOutput
}

type SecretBackendRootSignIntermediateArray []SecretBackendRootSignIntermediateInput

func (SecretBackendRootSignIntermediateArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SecretBackendRootSignIntermediate)(nil)).Elem()
}

func (i SecretBackendRootSignIntermediateArray) ToSecretBackendRootSignIntermediateArrayOutput() SecretBackendRootSignIntermediateArrayOutput {
	return i.ToSecretBackendRootSignIntermediateArrayOutputWithContext(context.Background())
}

func (i SecretBackendRootSignIntermediateArray) ToSecretBackendRootSignIntermediateArrayOutputWithContext(ctx context.Context) SecretBackendRootSignIntermediateArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecretBackendRootSignIntermediateArrayOutput)
}

// SecretBackendRootSignIntermediateMapInput is an input type that accepts SecretBackendRootSignIntermediateMap and SecretBackendRootSignIntermediateMapOutput values.
// You can construct a concrete instance of `SecretBackendRootSignIntermediateMapInput` via:
//
//	SecretBackendRootSignIntermediateMap{ "key": SecretBackendRootSignIntermediateArgs{...} }
type SecretBackendRootSignIntermediateMapInput interface {
	pulumi.Input

	ToSecretBackendRootSignIntermediateMapOutput() SecretBackendRootSignIntermediateMapOutput
	ToSecretBackendRootSignIntermediateMapOutputWithContext(context.Context) SecretBackendRootSignIntermediateMapOutput
}

type SecretBackendRootSignIntermediateMap map[string]SecretBackendRootSignIntermediateInput

func (SecretBackendRootSignIntermediateMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SecretBackendRootSignIntermediate)(nil)).Elem()
}

func (i SecretBackendRootSignIntermediateMap) ToSecretBackendRootSignIntermediateMapOutput() SecretBackendRootSignIntermediateMapOutput {
	return i.ToSecretBackendRootSignIntermediateMapOutputWithContext(context.Background())
}

func (i SecretBackendRootSignIntermediateMap) ToSecretBackendRootSignIntermediateMapOutputWithContext(ctx context.Context) SecretBackendRootSignIntermediateMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecretBackendRootSignIntermediateMapOutput)
}

type SecretBackendRootSignIntermediateOutput struct{ *pulumi.OutputState }

func (SecretBackendRootSignIntermediateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SecretBackendRootSignIntermediate)(nil)).Elem()
}

func (o SecretBackendRootSignIntermediateOutput) ToSecretBackendRootSignIntermediateOutput() SecretBackendRootSignIntermediateOutput {
	return o
}

func (o SecretBackendRootSignIntermediateOutput) ToSecretBackendRootSignIntermediateOutputWithContext(ctx context.Context) SecretBackendRootSignIntermediateOutput {
	return o
}

// List of alternative names
func (o SecretBackendRootSignIntermediateOutput) AltNames() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *SecretBackendRootSignIntermediate) pulumi.StringArrayOutput { return v.AltNames }).(pulumi.StringArrayOutput)
}

// The PKI secret backend the resource belongs to.
func (o SecretBackendRootSignIntermediateOutput) Backend() pulumi.StringOutput {
	return o.ApplyT(func(v *SecretBackendRootSignIntermediate) pulumi.StringOutput { return v.Backend }).(pulumi.StringOutput)
}

// A list of the issuing and intermediate CA certificates in the `format` specified.
func (o SecretBackendRootSignIntermediateOutput) CaChains() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *SecretBackendRootSignIntermediate) pulumi.StringArrayOutput { return v.CaChains }).(pulumi.StringArrayOutput)
}

// The intermediate CA certificate in the `format` specified.
func (o SecretBackendRootSignIntermediateOutput) Certificate() pulumi.StringOutput {
	return o.ApplyT(func(v *SecretBackendRootSignIntermediate) pulumi.StringOutput { return v.Certificate }).(pulumi.StringOutput)
}

// The concatenation of the intermediate CA and the issuing CA certificates (PEM encoded).
// Requires the `format` to be set to any of: pem, pem_bundle. The value will be empty for all other formats.
func (o SecretBackendRootSignIntermediateOutput) CertificateBundle() pulumi.StringOutput {
	return o.ApplyT(func(v *SecretBackendRootSignIntermediate) pulumi.StringOutput { return v.CertificateBundle }).(pulumi.StringOutput)
}

// CN of intermediate to create
func (o SecretBackendRootSignIntermediateOutput) CommonName() pulumi.StringOutput {
	return o.ApplyT(func(v *SecretBackendRootSignIntermediate) pulumi.StringOutput { return v.CommonName }).(pulumi.StringOutput)
}

// The country
func (o SecretBackendRootSignIntermediateOutput) Country() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecretBackendRootSignIntermediate) pulumi.StringPtrOutput { return v.Country }).(pulumi.StringPtrOutput)
}

// The CSR
func (o SecretBackendRootSignIntermediateOutput) Csr() pulumi.StringOutput {
	return o.ApplyT(func(v *SecretBackendRootSignIntermediate) pulumi.StringOutput { return v.Csr }).(pulumi.StringOutput)
}

// Flag to exclude CN from SANs
func (o SecretBackendRootSignIntermediateOutput) ExcludeCnFromSans() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SecretBackendRootSignIntermediate) pulumi.BoolPtrOutput { return v.ExcludeCnFromSans }).(pulumi.BoolPtrOutput)
}

// The format of data
func (o SecretBackendRootSignIntermediateOutput) Format() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecretBackendRootSignIntermediate) pulumi.StringPtrOutput { return v.Format }).(pulumi.StringPtrOutput)
}

// List of alternative IPs
func (o SecretBackendRootSignIntermediateOutput) IpSans() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *SecretBackendRootSignIntermediate) pulumi.StringArrayOutput { return v.IpSans }).(pulumi.StringArrayOutput)
}

// Specifies the default issuer of this request. May
// be the value `default`, a name, or an issuer ID. Use ACLs to prevent access to
// the `/pki/issuer/:issuer_ref/{issue,sign}/:name` paths to prevent users
// overriding the role's `issuerRef` value.
func (o SecretBackendRootSignIntermediateOutput) IssuerRef() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecretBackendRootSignIntermediate) pulumi.StringPtrOutput { return v.IssuerRef }).(pulumi.StringPtrOutput)
}

// The issuing CA certificate in the `format` specified.
func (o SecretBackendRootSignIntermediateOutput) IssuingCa() pulumi.StringOutput {
	return o.ApplyT(func(v *SecretBackendRootSignIntermediate) pulumi.StringOutput { return v.IssuingCa }).(pulumi.StringOutput)
}

// The locality
func (o SecretBackendRootSignIntermediateOutput) Locality() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecretBackendRootSignIntermediate) pulumi.StringPtrOutput { return v.Locality }).(pulumi.StringPtrOutput)
}

// The maximum path length to encode in the generated certificate
func (o SecretBackendRootSignIntermediateOutput) MaxPathLength() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SecretBackendRootSignIntermediate) pulumi.IntPtrOutput { return v.MaxPathLength }).(pulumi.IntPtrOutput)
}

// The namespace to provision the resource in.
// The value should not contain leading or trailing forward slashes.
// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
// *Available only for Vault Enterprise*.
func (o SecretBackendRootSignIntermediateOutput) Namespace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecretBackendRootSignIntermediate) pulumi.StringPtrOutput { return v.Namespace }).(pulumi.StringPtrOutput)
}

// The organization
func (o SecretBackendRootSignIntermediateOutput) Organization() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecretBackendRootSignIntermediate) pulumi.StringPtrOutput { return v.Organization }).(pulumi.StringPtrOutput)
}

// List of other SANs
func (o SecretBackendRootSignIntermediateOutput) OtherSans() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *SecretBackendRootSignIntermediate) pulumi.StringArrayOutput { return v.OtherSans }).(pulumi.StringArrayOutput)
}

// The organization unit
func (o SecretBackendRootSignIntermediateOutput) Ou() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecretBackendRootSignIntermediate) pulumi.StringPtrOutput { return v.Ou }).(pulumi.StringPtrOutput)
}

// List of domains for which certificates are allowed to be issued
func (o SecretBackendRootSignIntermediateOutput) PermittedDnsDomains() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *SecretBackendRootSignIntermediate) pulumi.StringArrayOutput { return v.PermittedDnsDomains }).(pulumi.StringArrayOutput)
}

// The postal code
func (o SecretBackendRootSignIntermediateOutput) PostalCode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecretBackendRootSignIntermediate) pulumi.StringPtrOutput { return v.PostalCode }).(pulumi.StringPtrOutput)
}

// The province
func (o SecretBackendRootSignIntermediateOutput) Province() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecretBackendRootSignIntermediate) pulumi.StringPtrOutput { return v.Province }).(pulumi.StringPtrOutput)
}

// If set to `true`, the certificate will be revoked on resource destruction.
func (o SecretBackendRootSignIntermediateOutput) Revoke() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SecretBackendRootSignIntermediate) pulumi.BoolPtrOutput { return v.Revoke }).(pulumi.BoolPtrOutput)
}

// The serial number.
//
// Deprecated: Use serialNumber instead
func (o SecretBackendRootSignIntermediateOutput) Serial() pulumi.StringOutput {
	return o.ApplyT(func(v *SecretBackendRootSignIntermediate) pulumi.StringOutput { return v.Serial }).(pulumi.StringOutput)
}

// The certificate's serial number, hex formatted.
func (o SecretBackendRootSignIntermediateOutput) SerialNumber() pulumi.StringOutput {
	return o.ApplyT(func(v *SecretBackendRootSignIntermediate) pulumi.StringOutput { return v.SerialNumber }).(pulumi.StringOutput)
}

// The street address
func (o SecretBackendRootSignIntermediateOutput) StreetAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecretBackendRootSignIntermediate) pulumi.StringPtrOutput { return v.StreetAddress }).(pulumi.StringPtrOutput)
}

// Time to live
func (o SecretBackendRootSignIntermediateOutput) Ttl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecretBackendRootSignIntermediate) pulumi.StringPtrOutput { return v.Ttl }).(pulumi.StringPtrOutput)
}

// List of alternative URIs
func (o SecretBackendRootSignIntermediateOutput) UriSans() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *SecretBackendRootSignIntermediate) pulumi.StringArrayOutput { return v.UriSans }).(pulumi.StringArrayOutput)
}

// Preserve CSR values
func (o SecretBackendRootSignIntermediateOutput) UseCsrValues() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SecretBackendRootSignIntermediate) pulumi.BoolPtrOutput { return v.UseCsrValues }).(pulumi.BoolPtrOutput)
}

type SecretBackendRootSignIntermediateArrayOutput struct{ *pulumi.OutputState }

func (SecretBackendRootSignIntermediateArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SecretBackendRootSignIntermediate)(nil)).Elem()
}

func (o SecretBackendRootSignIntermediateArrayOutput) ToSecretBackendRootSignIntermediateArrayOutput() SecretBackendRootSignIntermediateArrayOutput {
	return o
}

func (o SecretBackendRootSignIntermediateArrayOutput) ToSecretBackendRootSignIntermediateArrayOutputWithContext(ctx context.Context) SecretBackendRootSignIntermediateArrayOutput {
	return o
}

func (o SecretBackendRootSignIntermediateArrayOutput) Index(i pulumi.IntInput) SecretBackendRootSignIntermediateOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SecretBackendRootSignIntermediate {
		return vs[0].([]*SecretBackendRootSignIntermediate)[vs[1].(int)]
	}).(SecretBackendRootSignIntermediateOutput)
}

type SecretBackendRootSignIntermediateMapOutput struct{ *pulumi.OutputState }

func (SecretBackendRootSignIntermediateMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SecretBackendRootSignIntermediate)(nil)).Elem()
}

func (o SecretBackendRootSignIntermediateMapOutput) ToSecretBackendRootSignIntermediateMapOutput() SecretBackendRootSignIntermediateMapOutput {
	return o
}

func (o SecretBackendRootSignIntermediateMapOutput) ToSecretBackendRootSignIntermediateMapOutputWithContext(ctx context.Context) SecretBackendRootSignIntermediateMapOutput {
	return o
}

func (o SecretBackendRootSignIntermediateMapOutput) MapIndex(k pulumi.StringInput) SecretBackendRootSignIntermediateOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SecretBackendRootSignIntermediate {
		return vs[0].(map[string]*SecretBackendRootSignIntermediate)[vs[1].(string)]
	}).(SecretBackendRootSignIntermediateOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SecretBackendRootSignIntermediateInput)(nil)).Elem(), &SecretBackendRootSignIntermediate{})
	pulumi.RegisterInputType(reflect.TypeOf((*SecretBackendRootSignIntermediateArrayInput)(nil)).Elem(), SecretBackendRootSignIntermediateArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SecretBackendRootSignIntermediateMapInput)(nil)).Elem(), SecretBackendRootSignIntermediateMap{})
	pulumi.RegisterOutputType(SecretBackendRootSignIntermediateOutput{})
	pulumi.RegisterOutputType(SecretBackendRootSignIntermediateArrayOutput{})
	pulumi.RegisterOutputType(SecretBackendRootSignIntermediateMapOutput{})
}
