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
//	"github.com/pulumi/pulumi-vault/sdk/v5/go/vault/pkiSecret"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := pkiSecret.NewSecretBackendRootCert(ctx, "test", &pkiSecret.SecretBackendRootCertArgs{
//				Backend:           pulumi.Any(vault_mount.Pki.Path),
//				Type:              pulumi.String("internal"),
//				CommonName:        pulumi.String("Root CA"),
//				Ttl:               pulumi.String("315360000"),
//				Format:            pulumi.String("pem"),
//				PrivateKeyFormat:  pulumi.String("der"),
//				KeyType:           pulumi.String("rsa"),
//				KeyBits:           pulumi.Int(4096),
//				ExcludeCnFromSans: pulumi.Bool(true),
//				Ou:                pulumi.String("My OU"),
//				Organization:      pulumi.String("My organization"),
//			}, pulumi.DependsOn([]pulumi.Resource{
//				vault_mount.Pki,
//			}))
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
type SecretBackendRootCert struct {
	pulumi.CustomResourceState

	// List of alternative names
	AltNames pulumi.StringArrayOutput `pulumi:"altNames"`
	// The PKI secret backend the resource belongs to.
	Backend pulumi.StringOutput `pulumi:"backend"`
	// The certificate.
	Certificate pulumi.StringOutput `pulumi:"certificate"`
	// CN of intermediate to create
	CommonName pulumi.StringOutput `pulumi:"commonName"`
	// The country
	Country pulumi.StringPtrOutput `pulumi:"country"`
	// Flag to exclude CN from SANs
	ExcludeCnFromSans pulumi.BoolPtrOutput `pulumi:"excludeCnFromSans"`
	// The format of data
	Format pulumi.StringPtrOutput `pulumi:"format"`
	// List of alternative IPs
	IpSans pulumi.StringArrayOutput `pulumi:"ipSans"`
	// The ID of the generated issuer.
	IssuerId pulumi.StringOutput `pulumi:"issuerId"`
	// Provides a name to the specified issuer. The name must be unique
	// across all issuers and not be the reserved value `default`
	IssuerName pulumi.StringOutput `pulumi:"issuerName"`
	// The issuing CA certificate.
	IssuingCa pulumi.StringOutput `pulumi:"issuingCa"`
	// The number of bits to use
	KeyBits pulumi.IntPtrOutput `pulumi:"keyBits"`
	// The ID of the generated key.
	KeyId pulumi.StringOutput `pulumi:"keyId"`
	// When a new key is created with this request, optionally specifies
	// the name for this. The global ref `default` may not be used as a name.
	KeyName pulumi.StringOutput `pulumi:"keyName"`
	// Specifies the key (either default, by name, or by identifier) to use
	// for generating this request. Only suitable for `type=existing` requests.
	KeyRef pulumi.StringOutput `pulumi:"keyRef"`
	// The desired key type
	KeyType pulumi.StringPtrOutput `pulumi:"keyType"`
	// The locality
	Locality pulumi.StringPtrOutput `pulumi:"locality"`
	// The ID of the previously configured managed key. This field is
	// required if `type` is `kms` and it conflicts with `managedKeyName`
	ManagedKeyId pulumi.StringOutput `pulumi:"managedKeyId"`
	// The name of the previously configured managed key. This field is
	// required if `type` is `kms`  and it conflicts with `managedKeyId`
	ManagedKeyName pulumi.StringOutput `pulumi:"managedKeyName"`
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
	// The private key format
	PrivateKeyFormat pulumi.StringPtrOutput `pulumi:"privateKeyFormat"`
	// The province
	Province pulumi.StringPtrOutput `pulumi:"province"`
	// Deprecated, use `serialNumber` instead.
	//
	// Deprecated: Use serial_number instead
	Serial pulumi.StringOutput `pulumi:"serial"`
	// The certificate's serial number, hex formatted.
	SerialNumber pulumi.StringOutput `pulumi:"serialNumber"`
	// The street address
	StreetAddress pulumi.StringPtrOutput `pulumi:"streetAddress"`
	// Time to live
	Ttl pulumi.StringPtrOutput `pulumi:"ttl"`
	// Type of intermediate to create. Must be either \"exported\", \"internal\"
	// or \"kms\"
	Type pulumi.StringOutput `pulumi:"type"`
	// List of alternative URIs
	UriSans pulumi.StringArrayOutput `pulumi:"uriSans"`
}

// NewSecretBackendRootCert registers a new resource with the given unique name, arguments, and options.
func NewSecretBackendRootCert(ctx *pulumi.Context,
	name string, args *SecretBackendRootCertArgs, opts ...pulumi.ResourceOption) (*SecretBackendRootCert, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Backend == nil {
		return nil, errors.New("invalid value for required argument 'Backend'")
	}
	if args.CommonName == nil {
		return nil, errors.New("invalid value for required argument 'CommonName'")
	}
	if args.Type == nil {
		return nil, errors.New("invalid value for required argument 'Type'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource SecretBackendRootCert
	err := ctx.RegisterResource("vault:pkiSecret/secretBackendRootCert:SecretBackendRootCert", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSecretBackendRootCert gets an existing SecretBackendRootCert resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSecretBackendRootCert(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SecretBackendRootCertState, opts ...pulumi.ResourceOption) (*SecretBackendRootCert, error) {
	var resource SecretBackendRootCert
	err := ctx.ReadResource("vault:pkiSecret/secretBackendRootCert:SecretBackendRootCert", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SecretBackendRootCert resources.
type secretBackendRootCertState struct {
	// List of alternative names
	AltNames []string `pulumi:"altNames"`
	// The PKI secret backend the resource belongs to.
	Backend *string `pulumi:"backend"`
	// The certificate.
	Certificate *string `pulumi:"certificate"`
	// CN of intermediate to create
	CommonName *string `pulumi:"commonName"`
	// The country
	Country *string `pulumi:"country"`
	// Flag to exclude CN from SANs
	ExcludeCnFromSans *bool `pulumi:"excludeCnFromSans"`
	// The format of data
	Format *string `pulumi:"format"`
	// List of alternative IPs
	IpSans []string `pulumi:"ipSans"`
	// The ID of the generated issuer.
	IssuerId *string `pulumi:"issuerId"`
	// Provides a name to the specified issuer. The name must be unique
	// across all issuers and not be the reserved value `default`
	IssuerName *string `pulumi:"issuerName"`
	// The issuing CA certificate.
	IssuingCa *string `pulumi:"issuingCa"`
	// The number of bits to use
	KeyBits *int `pulumi:"keyBits"`
	// The ID of the generated key.
	KeyId *string `pulumi:"keyId"`
	// When a new key is created with this request, optionally specifies
	// the name for this. The global ref `default` may not be used as a name.
	KeyName *string `pulumi:"keyName"`
	// Specifies the key (either default, by name, or by identifier) to use
	// for generating this request. Only suitable for `type=existing` requests.
	KeyRef *string `pulumi:"keyRef"`
	// The desired key type
	KeyType *string `pulumi:"keyType"`
	// The locality
	Locality *string `pulumi:"locality"`
	// The ID of the previously configured managed key. This field is
	// required if `type` is `kms` and it conflicts with `managedKeyName`
	ManagedKeyId *string `pulumi:"managedKeyId"`
	// The name of the previously configured managed key. This field is
	// required if `type` is `kms`  and it conflicts with `managedKeyId`
	ManagedKeyName *string `pulumi:"managedKeyName"`
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
	// The private key format
	PrivateKeyFormat *string `pulumi:"privateKeyFormat"`
	// The province
	Province *string `pulumi:"province"`
	// Deprecated, use `serialNumber` instead.
	//
	// Deprecated: Use serial_number instead
	Serial *string `pulumi:"serial"`
	// The certificate's serial number, hex formatted.
	SerialNumber *string `pulumi:"serialNumber"`
	// The street address
	StreetAddress *string `pulumi:"streetAddress"`
	// Time to live
	Ttl *string `pulumi:"ttl"`
	// Type of intermediate to create. Must be either \"exported\", \"internal\"
	// or \"kms\"
	Type *string `pulumi:"type"`
	// List of alternative URIs
	UriSans []string `pulumi:"uriSans"`
}

type SecretBackendRootCertState struct {
	// List of alternative names
	AltNames pulumi.StringArrayInput
	// The PKI secret backend the resource belongs to.
	Backend pulumi.StringPtrInput
	// The certificate.
	Certificate pulumi.StringPtrInput
	// CN of intermediate to create
	CommonName pulumi.StringPtrInput
	// The country
	Country pulumi.StringPtrInput
	// Flag to exclude CN from SANs
	ExcludeCnFromSans pulumi.BoolPtrInput
	// The format of data
	Format pulumi.StringPtrInput
	// List of alternative IPs
	IpSans pulumi.StringArrayInput
	// The ID of the generated issuer.
	IssuerId pulumi.StringPtrInput
	// Provides a name to the specified issuer. The name must be unique
	// across all issuers and not be the reserved value `default`
	IssuerName pulumi.StringPtrInput
	// The issuing CA certificate.
	IssuingCa pulumi.StringPtrInput
	// The number of bits to use
	KeyBits pulumi.IntPtrInput
	// The ID of the generated key.
	KeyId pulumi.StringPtrInput
	// When a new key is created with this request, optionally specifies
	// the name for this. The global ref `default` may not be used as a name.
	KeyName pulumi.StringPtrInput
	// Specifies the key (either default, by name, or by identifier) to use
	// for generating this request. Only suitable for `type=existing` requests.
	KeyRef pulumi.StringPtrInput
	// The desired key type
	KeyType pulumi.StringPtrInput
	// The locality
	Locality pulumi.StringPtrInput
	// The ID of the previously configured managed key. This field is
	// required if `type` is `kms` and it conflicts with `managedKeyName`
	ManagedKeyId pulumi.StringPtrInput
	// The name of the previously configured managed key. This field is
	// required if `type` is `kms`  and it conflicts with `managedKeyId`
	ManagedKeyName pulumi.StringPtrInput
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
	// The private key format
	PrivateKeyFormat pulumi.StringPtrInput
	// The province
	Province pulumi.StringPtrInput
	// Deprecated, use `serialNumber` instead.
	//
	// Deprecated: Use serial_number instead
	Serial pulumi.StringPtrInput
	// The certificate's serial number, hex formatted.
	SerialNumber pulumi.StringPtrInput
	// The street address
	StreetAddress pulumi.StringPtrInput
	// Time to live
	Ttl pulumi.StringPtrInput
	// Type of intermediate to create. Must be either \"exported\", \"internal\"
	// or \"kms\"
	Type pulumi.StringPtrInput
	// List of alternative URIs
	UriSans pulumi.StringArrayInput
}

func (SecretBackendRootCertState) ElementType() reflect.Type {
	return reflect.TypeOf((*secretBackendRootCertState)(nil)).Elem()
}

type secretBackendRootCertArgs struct {
	// List of alternative names
	AltNames []string `pulumi:"altNames"`
	// The PKI secret backend the resource belongs to.
	Backend string `pulumi:"backend"`
	// CN of intermediate to create
	CommonName string `pulumi:"commonName"`
	// The country
	Country *string `pulumi:"country"`
	// Flag to exclude CN from SANs
	ExcludeCnFromSans *bool `pulumi:"excludeCnFromSans"`
	// The format of data
	Format *string `pulumi:"format"`
	// List of alternative IPs
	IpSans []string `pulumi:"ipSans"`
	// Provides a name to the specified issuer. The name must be unique
	// across all issuers and not be the reserved value `default`
	IssuerName *string `pulumi:"issuerName"`
	// The number of bits to use
	KeyBits *int `pulumi:"keyBits"`
	// When a new key is created with this request, optionally specifies
	// the name for this. The global ref `default` may not be used as a name.
	KeyName *string `pulumi:"keyName"`
	// Specifies the key (either default, by name, or by identifier) to use
	// for generating this request. Only suitable for `type=existing` requests.
	KeyRef *string `pulumi:"keyRef"`
	// The desired key type
	KeyType *string `pulumi:"keyType"`
	// The locality
	Locality *string `pulumi:"locality"`
	// The ID of the previously configured managed key. This field is
	// required if `type` is `kms` and it conflicts with `managedKeyName`
	ManagedKeyId *string `pulumi:"managedKeyId"`
	// The name of the previously configured managed key. This field is
	// required if `type` is `kms`  and it conflicts with `managedKeyId`
	ManagedKeyName *string `pulumi:"managedKeyName"`
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
	// The private key format
	PrivateKeyFormat *string `pulumi:"privateKeyFormat"`
	// The province
	Province *string `pulumi:"province"`
	// The street address
	StreetAddress *string `pulumi:"streetAddress"`
	// Time to live
	Ttl *string `pulumi:"ttl"`
	// Type of intermediate to create. Must be either \"exported\", \"internal\"
	// or \"kms\"
	Type string `pulumi:"type"`
	// List of alternative URIs
	UriSans []string `pulumi:"uriSans"`
}

// The set of arguments for constructing a SecretBackendRootCert resource.
type SecretBackendRootCertArgs struct {
	// List of alternative names
	AltNames pulumi.StringArrayInput
	// The PKI secret backend the resource belongs to.
	Backend pulumi.StringInput
	// CN of intermediate to create
	CommonName pulumi.StringInput
	// The country
	Country pulumi.StringPtrInput
	// Flag to exclude CN from SANs
	ExcludeCnFromSans pulumi.BoolPtrInput
	// The format of data
	Format pulumi.StringPtrInput
	// List of alternative IPs
	IpSans pulumi.StringArrayInput
	// Provides a name to the specified issuer. The name must be unique
	// across all issuers and not be the reserved value `default`
	IssuerName pulumi.StringPtrInput
	// The number of bits to use
	KeyBits pulumi.IntPtrInput
	// When a new key is created with this request, optionally specifies
	// the name for this. The global ref `default` may not be used as a name.
	KeyName pulumi.StringPtrInput
	// Specifies the key (either default, by name, or by identifier) to use
	// for generating this request. Only suitable for `type=existing` requests.
	KeyRef pulumi.StringPtrInput
	// The desired key type
	KeyType pulumi.StringPtrInput
	// The locality
	Locality pulumi.StringPtrInput
	// The ID of the previously configured managed key. This field is
	// required if `type` is `kms` and it conflicts with `managedKeyName`
	ManagedKeyId pulumi.StringPtrInput
	// The name of the previously configured managed key. This field is
	// required if `type` is `kms`  and it conflicts with `managedKeyId`
	ManagedKeyName pulumi.StringPtrInput
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
	// The private key format
	PrivateKeyFormat pulumi.StringPtrInput
	// The province
	Province pulumi.StringPtrInput
	// The street address
	StreetAddress pulumi.StringPtrInput
	// Time to live
	Ttl pulumi.StringPtrInput
	// Type of intermediate to create. Must be either \"exported\", \"internal\"
	// or \"kms\"
	Type pulumi.StringInput
	// List of alternative URIs
	UriSans pulumi.StringArrayInput
}

func (SecretBackendRootCertArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*secretBackendRootCertArgs)(nil)).Elem()
}

type SecretBackendRootCertInput interface {
	pulumi.Input

	ToSecretBackendRootCertOutput() SecretBackendRootCertOutput
	ToSecretBackendRootCertOutputWithContext(ctx context.Context) SecretBackendRootCertOutput
}

func (*SecretBackendRootCert) ElementType() reflect.Type {
	return reflect.TypeOf((**SecretBackendRootCert)(nil)).Elem()
}

func (i *SecretBackendRootCert) ToSecretBackendRootCertOutput() SecretBackendRootCertOutput {
	return i.ToSecretBackendRootCertOutputWithContext(context.Background())
}

func (i *SecretBackendRootCert) ToSecretBackendRootCertOutputWithContext(ctx context.Context) SecretBackendRootCertOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecretBackendRootCertOutput)
}

// SecretBackendRootCertArrayInput is an input type that accepts SecretBackendRootCertArray and SecretBackendRootCertArrayOutput values.
// You can construct a concrete instance of `SecretBackendRootCertArrayInput` via:
//
//	SecretBackendRootCertArray{ SecretBackendRootCertArgs{...} }
type SecretBackendRootCertArrayInput interface {
	pulumi.Input

	ToSecretBackendRootCertArrayOutput() SecretBackendRootCertArrayOutput
	ToSecretBackendRootCertArrayOutputWithContext(context.Context) SecretBackendRootCertArrayOutput
}

type SecretBackendRootCertArray []SecretBackendRootCertInput

func (SecretBackendRootCertArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SecretBackendRootCert)(nil)).Elem()
}

func (i SecretBackendRootCertArray) ToSecretBackendRootCertArrayOutput() SecretBackendRootCertArrayOutput {
	return i.ToSecretBackendRootCertArrayOutputWithContext(context.Background())
}

func (i SecretBackendRootCertArray) ToSecretBackendRootCertArrayOutputWithContext(ctx context.Context) SecretBackendRootCertArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecretBackendRootCertArrayOutput)
}

// SecretBackendRootCertMapInput is an input type that accepts SecretBackendRootCertMap and SecretBackendRootCertMapOutput values.
// You can construct a concrete instance of `SecretBackendRootCertMapInput` via:
//
//	SecretBackendRootCertMap{ "key": SecretBackendRootCertArgs{...} }
type SecretBackendRootCertMapInput interface {
	pulumi.Input

	ToSecretBackendRootCertMapOutput() SecretBackendRootCertMapOutput
	ToSecretBackendRootCertMapOutputWithContext(context.Context) SecretBackendRootCertMapOutput
}

type SecretBackendRootCertMap map[string]SecretBackendRootCertInput

func (SecretBackendRootCertMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SecretBackendRootCert)(nil)).Elem()
}

func (i SecretBackendRootCertMap) ToSecretBackendRootCertMapOutput() SecretBackendRootCertMapOutput {
	return i.ToSecretBackendRootCertMapOutputWithContext(context.Background())
}

func (i SecretBackendRootCertMap) ToSecretBackendRootCertMapOutputWithContext(ctx context.Context) SecretBackendRootCertMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecretBackendRootCertMapOutput)
}

type SecretBackendRootCertOutput struct{ *pulumi.OutputState }

func (SecretBackendRootCertOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SecretBackendRootCert)(nil)).Elem()
}

func (o SecretBackendRootCertOutput) ToSecretBackendRootCertOutput() SecretBackendRootCertOutput {
	return o
}

func (o SecretBackendRootCertOutput) ToSecretBackendRootCertOutputWithContext(ctx context.Context) SecretBackendRootCertOutput {
	return o
}

// List of alternative names
func (o SecretBackendRootCertOutput) AltNames() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *SecretBackendRootCert) pulumi.StringArrayOutput { return v.AltNames }).(pulumi.StringArrayOutput)
}

// The PKI secret backend the resource belongs to.
func (o SecretBackendRootCertOutput) Backend() pulumi.StringOutput {
	return o.ApplyT(func(v *SecretBackendRootCert) pulumi.StringOutput { return v.Backend }).(pulumi.StringOutput)
}

// The certificate.
func (o SecretBackendRootCertOutput) Certificate() pulumi.StringOutput {
	return o.ApplyT(func(v *SecretBackendRootCert) pulumi.StringOutput { return v.Certificate }).(pulumi.StringOutput)
}

// CN of intermediate to create
func (o SecretBackendRootCertOutput) CommonName() pulumi.StringOutput {
	return o.ApplyT(func(v *SecretBackendRootCert) pulumi.StringOutput { return v.CommonName }).(pulumi.StringOutput)
}

// The country
func (o SecretBackendRootCertOutput) Country() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecretBackendRootCert) pulumi.StringPtrOutput { return v.Country }).(pulumi.StringPtrOutput)
}

// Flag to exclude CN from SANs
func (o SecretBackendRootCertOutput) ExcludeCnFromSans() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SecretBackendRootCert) pulumi.BoolPtrOutput { return v.ExcludeCnFromSans }).(pulumi.BoolPtrOutput)
}

// The format of data
func (o SecretBackendRootCertOutput) Format() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecretBackendRootCert) pulumi.StringPtrOutput { return v.Format }).(pulumi.StringPtrOutput)
}

// List of alternative IPs
func (o SecretBackendRootCertOutput) IpSans() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *SecretBackendRootCert) pulumi.StringArrayOutput { return v.IpSans }).(pulumi.StringArrayOutput)
}

// The ID of the generated issuer.
func (o SecretBackendRootCertOutput) IssuerId() pulumi.StringOutput {
	return o.ApplyT(func(v *SecretBackendRootCert) pulumi.StringOutput { return v.IssuerId }).(pulumi.StringOutput)
}

// Provides a name to the specified issuer. The name must be unique
// across all issuers and not be the reserved value `default`
func (o SecretBackendRootCertOutput) IssuerName() pulumi.StringOutput {
	return o.ApplyT(func(v *SecretBackendRootCert) pulumi.StringOutput { return v.IssuerName }).(pulumi.StringOutput)
}

// The issuing CA certificate.
func (o SecretBackendRootCertOutput) IssuingCa() pulumi.StringOutput {
	return o.ApplyT(func(v *SecretBackendRootCert) pulumi.StringOutput { return v.IssuingCa }).(pulumi.StringOutput)
}

// The number of bits to use
func (o SecretBackendRootCertOutput) KeyBits() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SecretBackendRootCert) pulumi.IntPtrOutput { return v.KeyBits }).(pulumi.IntPtrOutput)
}

// The ID of the generated key.
func (o SecretBackendRootCertOutput) KeyId() pulumi.StringOutput {
	return o.ApplyT(func(v *SecretBackendRootCert) pulumi.StringOutput { return v.KeyId }).(pulumi.StringOutput)
}

// When a new key is created with this request, optionally specifies
// the name for this. The global ref `default` may not be used as a name.
func (o SecretBackendRootCertOutput) KeyName() pulumi.StringOutput {
	return o.ApplyT(func(v *SecretBackendRootCert) pulumi.StringOutput { return v.KeyName }).(pulumi.StringOutput)
}

// Specifies the key (either default, by name, or by identifier) to use
// for generating this request. Only suitable for `type=existing` requests.
func (o SecretBackendRootCertOutput) KeyRef() pulumi.StringOutput {
	return o.ApplyT(func(v *SecretBackendRootCert) pulumi.StringOutput { return v.KeyRef }).(pulumi.StringOutput)
}

// The desired key type
func (o SecretBackendRootCertOutput) KeyType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecretBackendRootCert) pulumi.StringPtrOutput { return v.KeyType }).(pulumi.StringPtrOutput)
}

// The locality
func (o SecretBackendRootCertOutput) Locality() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecretBackendRootCert) pulumi.StringPtrOutput { return v.Locality }).(pulumi.StringPtrOutput)
}

// The ID of the previously configured managed key. This field is
// required if `type` is `kms` and it conflicts with `managedKeyName`
func (o SecretBackendRootCertOutput) ManagedKeyId() pulumi.StringOutput {
	return o.ApplyT(func(v *SecretBackendRootCert) pulumi.StringOutput { return v.ManagedKeyId }).(pulumi.StringOutput)
}

// The name of the previously configured managed key. This field is
// required if `type` is `kms`  and it conflicts with `managedKeyId`
func (o SecretBackendRootCertOutput) ManagedKeyName() pulumi.StringOutput {
	return o.ApplyT(func(v *SecretBackendRootCert) pulumi.StringOutput { return v.ManagedKeyName }).(pulumi.StringOutput)
}

// The maximum path length to encode in the generated certificate
func (o SecretBackendRootCertOutput) MaxPathLength() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SecretBackendRootCert) pulumi.IntPtrOutput { return v.MaxPathLength }).(pulumi.IntPtrOutput)
}

// The namespace to provision the resource in.
// The value should not contain leading or trailing forward slashes.
// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
// *Available only for Vault Enterprise*.
func (o SecretBackendRootCertOutput) Namespace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecretBackendRootCert) pulumi.StringPtrOutput { return v.Namespace }).(pulumi.StringPtrOutput)
}

// The organization
func (o SecretBackendRootCertOutput) Organization() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecretBackendRootCert) pulumi.StringPtrOutput { return v.Organization }).(pulumi.StringPtrOutput)
}

// List of other SANs
func (o SecretBackendRootCertOutput) OtherSans() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *SecretBackendRootCert) pulumi.StringArrayOutput { return v.OtherSans }).(pulumi.StringArrayOutput)
}

// The organization unit
func (o SecretBackendRootCertOutput) Ou() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecretBackendRootCert) pulumi.StringPtrOutput { return v.Ou }).(pulumi.StringPtrOutput)
}

// List of domains for which certificates are allowed to be issued
func (o SecretBackendRootCertOutput) PermittedDnsDomains() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *SecretBackendRootCert) pulumi.StringArrayOutput { return v.PermittedDnsDomains }).(pulumi.StringArrayOutput)
}

// The postal code
func (o SecretBackendRootCertOutput) PostalCode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecretBackendRootCert) pulumi.StringPtrOutput { return v.PostalCode }).(pulumi.StringPtrOutput)
}

// The private key format
func (o SecretBackendRootCertOutput) PrivateKeyFormat() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecretBackendRootCert) pulumi.StringPtrOutput { return v.PrivateKeyFormat }).(pulumi.StringPtrOutput)
}

// The province
func (o SecretBackendRootCertOutput) Province() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecretBackendRootCert) pulumi.StringPtrOutput { return v.Province }).(pulumi.StringPtrOutput)
}

// Deprecated, use `serialNumber` instead.
//
// Deprecated: Use serial_number instead
func (o SecretBackendRootCertOutput) Serial() pulumi.StringOutput {
	return o.ApplyT(func(v *SecretBackendRootCert) pulumi.StringOutput { return v.Serial }).(pulumi.StringOutput)
}

// The certificate's serial number, hex formatted.
func (o SecretBackendRootCertOutput) SerialNumber() pulumi.StringOutput {
	return o.ApplyT(func(v *SecretBackendRootCert) pulumi.StringOutput { return v.SerialNumber }).(pulumi.StringOutput)
}

// The street address
func (o SecretBackendRootCertOutput) StreetAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecretBackendRootCert) pulumi.StringPtrOutput { return v.StreetAddress }).(pulumi.StringPtrOutput)
}

// Time to live
func (o SecretBackendRootCertOutput) Ttl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecretBackendRootCert) pulumi.StringPtrOutput { return v.Ttl }).(pulumi.StringPtrOutput)
}

// Type of intermediate to create. Must be either \"exported\", \"internal\"
// or \"kms\"
func (o SecretBackendRootCertOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *SecretBackendRootCert) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// List of alternative URIs
func (o SecretBackendRootCertOutput) UriSans() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *SecretBackendRootCert) pulumi.StringArrayOutput { return v.UriSans }).(pulumi.StringArrayOutput)
}

type SecretBackendRootCertArrayOutput struct{ *pulumi.OutputState }

func (SecretBackendRootCertArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SecretBackendRootCert)(nil)).Elem()
}

func (o SecretBackendRootCertArrayOutput) ToSecretBackendRootCertArrayOutput() SecretBackendRootCertArrayOutput {
	return o
}

func (o SecretBackendRootCertArrayOutput) ToSecretBackendRootCertArrayOutputWithContext(ctx context.Context) SecretBackendRootCertArrayOutput {
	return o
}

func (o SecretBackendRootCertArrayOutput) Index(i pulumi.IntInput) SecretBackendRootCertOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SecretBackendRootCert {
		return vs[0].([]*SecretBackendRootCert)[vs[1].(int)]
	}).(SecretBackendRootCertOutput)
}

type SecretBackendRootCertMapOutput struct{ *pulumi.OutputState }

func (SecretBackendRootCertMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SecretBackendRootCert)(nil)).Elem()
}

func (o SecretBackendRootCertMapOutput) ToSecretBackendRootCertMapOutput() SecretBackendRootCertMapOutput {
	return o
}

func (o SecretBackendRootCertMapOutput) ToSecretBackendRootCertMapOutputWithContext(ctx context.Context) SecretBackendRootCertMapOutput {
	return o
}

func (o SecretBackendRootCertMapOutput) MapIndex(k pulumi.StringInput) SecretBackendRootCertOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SecretBackendRootCert {
		return vs[0].(map[string]*SecretBackendRootCert)[vs[1].(string)]
	}).(SecretBackendRootCertOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SecretBackendRootCertInput)(nil)).Elem(), &SecretBackendRootCert{})
	pulumi.RegisterInputType(reflect.TypeOf((*SecretBackendRootCertArrayInput)(nil)).Elem(), SecretBackendRootCertArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SecretBackendRootCertMapInput)(nil)).Elem(), SecretBackendRootCertMap{})
	pulumi.RegisterOutputType(SecretBackendRootCertOutput{})
	pulumi.RegisterOutputType(SecretBackendRootCertArrayOutput{})
	pulumi.RegisterOutputType(SecretBackendRootCertMapOutput{})
}
