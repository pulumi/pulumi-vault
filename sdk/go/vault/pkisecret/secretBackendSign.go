// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package pkisecret

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-vault/sdk/v6/go/vault/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## Example Usage
//
// <!--Start PulumiCodeChooser -->
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-vault/sdk/v6/go/vault/pkiSecret"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := pkiSecret.NewSecretBackendSign(ctx, "test", &pkiSecret.SecretBackendSignArgs{
//				Backend: pulumi.Any(vault_mount.Pki.Path),
//				Csr: pulumi.String(`-----BEGIN CERTIFICATE REQUEST-----
//
// MIIEqDCCApACAQAwYzELMAkGA1UEBhMCQVUxEzARBgNVBAgMClNvbWUtU3RhdGUx
// ITAfBgNVBAoMGEludGVybmV0IFdpZGdpdHMgUHR5IEx0ZDEcMBoGA1UEAwwTY2Vy
// dC50ZXN0Lm15LmRvbWFpbjCCAiIwDQYJKoZIhvcNAQEBBQADggIPADCCAgoCggIB
// AJupYCQ8UVCWII1Zof1c6YcSSaM9hEaDU78cfKP5RoSeH10BvrWRfT+mzCONVpNP
// CW9Iabtvk6hm0ot6ilnndEyVJbc0g7hdDLBX5BM25D+DGZGJRKUz1V+uBrWmXtIt
// Vonj7JTDTe7ViH0GDsB7CvqXFGXO2a2cDYBchLkL6vQiFPshxvUsLtwxuy/qdYgy
// X6ya+AUoZcoQGy1XxNjfH6cPtWSWQGEp1oPR6vL9hU3laTZb3C+VV4jZem+he8/0
// V+qV6fLG92WTXm2hmf8nrtUqqJ+C7mW/RJod+TviviBadIX0OHXW7k5HVsZood01
// te8vMRUNJNiZfa9EMIK5oncbQn0LcM3Wo9VrjpL7jREb/4HCS2gswYGv7hzk9cCS
// kVY4rDucchKbApuI3kfzmO7GFOF5eiSkYZpY/czNn7VVM3WCu6dpOX4+3rhgrZQw
// kY14L930DaLVRUgve/zKVP2D2GHdEOs+MbV7s96UgigT9pXly/yHPj+1sSYqmnaD
// 5b7jSeJusmzO/nrwXVGLsnezR87VzHl9Ux9g5s6zh+R+PrZuVxYsLvoUpaasH47O
// gIcBzSb/6pSGZKAUizmYsHsR1k88dAvsQ+FsUDaNokdi9VndEB4QPmiFmjyLV+0I
// 1TFoXop4sW11NPz1YCq+IxnYrEaIN3PyhY0GvBJDFY1/AgMBAAGgADANBgkqhkiG
// 9w0BAQsFAAOCAgEActuqnqS8Y9UF7e08w7tR3FPzGecWreuvxILrlFEZJxiLPFqL
// It7uJvtypCVQvz6UQzKdBYO7tMpRaWViB8DrWzXNZjLMrg+QHcpveg8C0Ett4scG
// fnvLk6fTDFYrnGvwHTqiHos5i0y3bFLyS1BGwSpdLAykGtvC+VM8mRyw/Y7CPcKN
// 77kebY/9xduW1g2uxWLr0x90RuQDv9psPojT+59tRLGSp5Kt0IeD3QtnAZEFE4aN
// vt+Pd69eg3BgZ8ZeDgoqAw3yppvOkpAFiE5pw2qPZaM4SRphl4d2Lek2zNIMyZqv
// do5zh356HOgXtDaSg0POnRGrN/Ua+LMCRTg6GEPUnx9uQb/zt8Zu0hIexDGyykp1
// OGqtWlv/Nc8UYuS38v0BeB6bMPeoqQUjkqs8nHlAEFn0KlgYdtDC+7SdQx6wS4te
// dBKRNDfC4lS3jYJgs55jHqonZgkpSi3bamlxpfpW0ukGBcmq91wRe4bOw/4uD/vf
// UwqMWOdCYcU3mdYNjTWy22ORW3SGFQxMBwpUEURCSoeqWr6aJeQ7KAYkx1PrB5T8
// OTEc13lWf+B0PU9UJuGTsmpIuImPDVd0EVDayr3mT5dDbqTVDbe8ppf2IswABmf0
// o3DybUeUmknYjl109rdSf+76nuREICHatxXgN3xCMFuBaN4WLO+ksd6Y1Ys=
// -----END CERTIFICATE REQUEST-----
// `),
//
//				CommonName: pulumi.String("test.my.domain"),
//			}, pulumi.DependsOn([]pulumi.Resource{
//				vault_pki_secret_backend_role.Admin,
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
type SecretBackendSign struct {
	pulumi.CustomResourceState

	// List of alternative names
	AltNames pulumi.StringArrayOutput `pulumi:"altNames"`
	// If set to `true`, certs will be renewed if the expiration is within `minSecondsRemaining`. Default `false`
	AutoRenew pulumi.BoolPtrOutput `pulumi:"autoRenew"`
	// The PKI secret backend the resource belongs to.
	Backend pulumi.StringOutput `pulumi:"backend"`
	// The CA chain
	CaChains pulumi.StringArrayOutput `pulumi:"caChains"`
	// The certificate
	Certificate pulumi.StringOutput `pulumi:"certificate"`
	// CN of certificate to create
	CommonName pulumi.StringOutput `pulumi:"commonName"`
	// The CSR
	Csr pulumi.StringOutput `pulumi:"csr"`
	// Flag to exclude CN from SANs
	ExcludeCnFromSans pulumi.BoolPtrOutput `pulumi:"excludeCnFromSans"`
	// The expiration date of the certificate in unix epoch format
	Expiration pulumi.IntOutput `pulumi:"expiration"`
	// The format of data
	Format pulumi.StringPtrOutput `pulumi:"format"`
	// List of alternative IPs
	IpSans pulumi.StringArrayOutput `pulumi:"ipSans"`
	// Specifies the default issuer of this request. Can
	// be the value `default`, a name, or an issuer ID. Use ACLs to prevent access to
	// the `/pki/issuer/:issuer_ref/{issue,sign}/:name` paths to prevent users
	// overriding the role's `issuerRef` value.
	IssuerRef pulumi.StringPtrOutput `pulumi:"issuerRef"`
	// The issuing CA
	IssuingCa pulumi.StringOutput `pulumi:"issuingCa"`
	// Generate a new certificate when the expiration is within this number of seconds, default is 604800 (7 days)
	MinSecondsRemaining pulumi.IntPtrOutput `pulumi:"minSecondsRemaining"`
	// Name of the role to create the certificate against
	Name pulumi.StringOutput `pulumi:"name"`
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
	// *Available only for Vault Enterprise*.
	Namespace pulumi.StringPtrOutput `pulumi:"namespace"`
	// List of other SANs
	OtherSans pulumi.StringArrayOutput `pulumi:"otherSans"`
	// `true` if the current time (during refresh) is after the start of the early renewal window declared by `minSecondsRemaining`, and `false` otherwise; if `autoRenew` is set to `true` then the provider will plan to replace the certificate once renewal is pending.
	RenewPending pulumi.BoolOutput `pulumi:"renewPending"`
	// The certificate's serial number, hex formatted.
	SerialNumber pulumi.StringOutput `pulumi:"serialNumber"`
	// Time to live
	Ttl pulumi.StringPtrOutput `pulumi:"ttl"`
	// List of alternative URIs
	UriSans pulumi.StringArrayOutput `pulumi:"uriSans"`
}

// NewSecretBackendSign registers a new resource with the given unique name, arguments, and options.
func NewSecretBackendSign(ctx *pulumi.Context,
	name string, args *SecretBackendSignArgs, opts ...pulumi.ResourceOption) (*SecretBackendSign, error) {
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
	var resource SecretBackendSign
	err := ctx.RegisterResource("vault:pkiSecret/secretBackendSign:SecretBackendSign", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSecretBackendSign gets an existing SecretBackendSign resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSecretBackendSign(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SecretBackendSignState, opts ...pulumi.ResourceOption) (*SecretBackendSign, error) {
	var resource SecretBackendSign
	err := ctx.ReadResource("vault:pkiSecret/secretBackendSign:SecretBackendSign", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SecretBackendSign resources.
type secretBackendSignState struct {
	// List of alternative names
	AltNames []string `pulumi:"altNames"`
	// If set to `true`, certs will be renewed if the expiration is within `minSecondsRemaining`. Default `false`
	AutoRenew *bool `pulumi:"autoRenew"`
	// The PKI secret backend the resource belongs to.
	Backend *string `pulumi:"backend"`
	// The CA chain
	CaChains []string `pulumi:"caChains"`
	// The certificate
	Certificate *string `pulumi:"certificate"`
	// CN of certificate to create
	CommonName *string `pulumi:"commonName"`
	// The CSR
	Csr *string `pulumi:"csr"`
	// Flag to exclude CN from SANs
	ExcludeCnFromSans *bool `pulumi:"excludeCnFromSans"`
	// The expiration date of the certificate in unix epoch format
	Expiration *int `pulumi:"expiration"`
	// The format of data
	Format *string `pulumi:"format"`
	// List of alternative IPs
	IpSans []string `pulumi:"ipSans"`
	// Specifies the default issuer of this request. Can
	// be the value `default`, a name, or an issuer ID. Use ACLs to prevent access to
	// the `/pki/issuer/:issuer_ref/{issue,sign}/:name` paths to prevent users
	// overriding the role's `issuerRef` value.
	IssuerRef *string `pulumi:"issuerRef"`
	// The issuing CA
	IssuingCa *string `pulumi:"issuingCa"`
	// Generate a new certificate when the expiration is within this number of seconds, default is 604800 (7 days)
	MinSecondsRemaining *int `pulumi:"minSecondsRemaining"`
	// Name of the role to create the certificate against
	Name *string `pulumi:"name"`
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
	// *Available only for Vault Enterprise*.
	Namespace *string `pulumi:"namespace"`
	// List of other SANs
	OtherSans []string `pulumi:"otherSans"`
	// `true` if the current time (during refresh) is after the start of the early renewal window declared by `minSecondsRemaining`, and `false` otherwise; if `autoRenew` is set to `true` then the provider will plan to replace the certificate once renewal is pending.
	RenewPending *bool `pulumi:"renewPending"`
	// The certificate's serial number, hex formatted.
	SerialNumber *string `pulumi:"serialNumber"`
	// Time to live
	Ttl *string `pulumi:"ttl"`
	// List of alternative URIs
	UriSans []string `pulumi:"uriSans"`
}

type SecretBackendSignState struct {
	// List of alternative names
	AltNames pulumi.StringArrayInput
	// If set to `true`, certs will be renewed if the expiration is within `minSecondsRemaining`. Default `false`
	AutoRenew pulumi.BoolPtrInput
	// The PKI secret backend the resource belongs to.
	Backend pulumi.StringPtrInput
	// The CA chain
	CaChains pulumi.StringArrayInput
	// The certificate
	Certificate pulumi.StringPtrInput
	// CN of certificate to create
	CommonName pulumi.StringPtrInput
	// The CSR
	Csr pulumi.StringPtrInput
	// Flag to exclude CN from SANs
	ExcludeCnFromSans pulumi.BoolPtrInput
	// The expiration date of the certificate in unix epoch format
	Expiration pulumi.IntPtrInput
	// The format of data
	Format pulumi.StringPtrInput
	// List of alternative IPs
	IpSans pulumi.StringArrayInput
	// Specifies the default issuer of this request. Can
	// be the value `default`, a name, or an issuer ID. Use ACLs to prevent access to
	// the `/pki/issuer/:issuer_ref/{issue,sign}/:name` paths to prevent users
	// overriding the role's `issuerRef` value.
	IssuerRef pulumi.StringPtrInput
	// The issuing CA
	IssuingCa pulumi.StringPtrInput
	// Generate a new certificate when the expiration is within this number of seconds, default is 604800 (7 days)
	MinSecondsRemaining pulumi.IntPtrInput
	// Name of the role to create the certificate against
	Name pulumi.StringPtrInput
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
	// *Available only for Vault Enterprise*.
	Namespace pulumi.StringPtrInput
	// List of other SANs
	OtherSans pulumi.StringArrayInput
	// `true` if the current time (during refresh) is after the start of the early renewal window declared by `minSecondsRemaining`, and `false` otherwise; if `autoRenew` is set to `true` then the provider will plan to replace the certificate once renewal is pending.
	RenewPending pulumi.BoolPtrInput
	// The certificate's serial number, hex formatted.
	SerialNumber pulumi.StringPtrInput
	// Time to live
	Ttl pulumi.StringPtrInput
	// List of alternative URIs
	UriSans pulumi.StringArrayInput
}

func (SecretBackendSignState) ElementType() reflect.Type {
	return reflect.TypeOf((*secretBackendSignState)(nil)).Elem()
}

type secretBackendSignArgs struct {
	// List of alternative names
	AltNames []string `pulumi:"altNames"`
	// If set to `true`, certs will be renewed if the expiration is within `minSecondsRemaining`. Default `false`
	AutoRenew *bool `pulumi:"autoRenew"`
	// The PKI secret backend the resource belongs to.
	Backend string `pulumi:"backend"`
	// CN of certificate to create
	CommonName string `pulumi:"commonName"`
	// The CSR
	Csr string `pulumi:"csr"`
	// Flag to exclude CN from SANs
	ExcludeCnFromSans *bool `pulumi:"excludeCnFromSans"`
	// The format of data
	Format *string `pulumi:"format"`
	// List of alternative IPs
	IpSans []string `pulumi:"ipSans"`
	// Specifies the default issuer of this request. Can
	// be the value `default`, a name, or an issuer ID. Use ACLs to prevent access to
	// the `/pki/issuer/:issuer_ref/{issue,sign}/:name` paths to prevent users
	// overriding the role's `issuerRef` value.
	IssuerRef *string `pulumi:"issuerRef"`
	// Generate a new certificate when the expiration is within this number of seconds, default is 604800 (7 days)
	MinSecondsRemaining *int `pulumi:"minSecondsRemaining"`
	// Name of the role to create the certificate against
	Name *string `pulumi:"name"`
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
	// *Available only for Vault Enterprise*.
	Namespace *string `pulumi:"namespace"`
	// List of other SANs
	OtherSans []string `pulumi:"otherSans"`
	// Time to live
	Ttl *string `pulumi:"ttl"`
	// List of alternative URIs
	UriSans []string `pulumi:"uriSans"`
}

// The set of arguments for constructing a SecretBackendSign resource.
type SecretBackendSignArgs struct {
	// List of alternative names
	AltNames pulumi.StringArrayInput
	// If set to `true`, certs will be renewed if the expiration is within `minSecondsRemaining`. Default `false`
	AutoRenew pulumi.BoolPtrInput
	// The PKI secret backend the resource belongs to.
	Backend pulumi.StringInput
	// CN of certificate to create
	CommonName pulumi.StringInput
	// The CSR
	Csr pulumi.StringInput
	// Flag to exclude CN from SANs
	ExcludeCnFromSans pulumi.BoolPtrInput
	// The format of data
	Format pulumi.StringPtrInput
	// List of alternative IPs
	IpSans pulumi.StringArrayInput
	// Specifies the default issuer of this request. Can
	// be the value `default`, a name, or an issuer ID. Use ACLs to prevent access to
	// the `/pki/issuer/:issuer_ref/{issue,sign}/:name` paths to prevent users
	// overriding the role's `issuerRef` value.
	IssuerRef pulumi.StringPtrInput
	// Generate a new certificate when the expiration is within this number of seconds, default is 604800 (7 days)
	MinSecondsRemaining pulumi.IntPtrInput
	// Name of the role to create the certificate against
	Name pulumi.StringPtrInput
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
	// *Available only for Vault Enterprise*.
	Namespace pulumi.StringPtrInput
	// List of other SANs
	OtherSans pulumi.StringArrayInput
	// Time to live
	Ttl pulumi.StringPtrInput
	// List of alternative URIs
	UriSans pulumi.StringArrayInput
}

func (SecretBackendSignArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*secretBackendSignArgs)(nil)).Elem()
}

type SecretBackendSignInput interface {
	pulumi.Input

	ToSecretBackendSignOutput() SecretBackendSignOutput
	ToSecretBackendSignOutputWithContext(ctx context.Context) SecretBackendSignOutput
}

func (*SecretBackendSign) ElementType() reflect.Type {
	return reflect.TypeOf((**SecretBackendSign)(nil)).Elem()
}

func (i *SecretBackendSign) ToSecretBackendSignOutput() SecretBackendSignOutput {
	return i.ToSecretBackendSignOutputWithContext(context.Background())
}

func (i *SecretBackendSign) ToSecretBackendSignOutputWithContext(ctx context.Context) SecretBackendSignOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecretBackendSignOutput)
}

// SecretBackendSignArrayInput is an input type that accepts SecretBackendSignArray and SecretBackendSignArrayOutput values.
// You can construct a concrete instance of `SecretBackendSignArrayInput` via:
//
//	SecretBackendSignArray{ SecretBackendSignArgs{...} }
type SecretBackendSignArrayInput interface {
	pulumi.Input

	ToSecretBackendSignArrayOutput() SecretBackendSignArrayOutput
	ToSecretBackendSignArrayOutputWithContext(context.Context) SecretBackendSignArrayOutput
}

type SecretBackendSignArray []SecretBackendSignInput

func (SecretBackendSignArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SecretBackendSign)(nil)).Elem()
}

func (i SecretBackendSignArray) ToSecretBackendSignArrayOutput() SecretBackendSignArrayOutput {
	return i.ToSecretBackendSignArrayOutputWithContext(context.Background())
}

func (i SecretBackendSignArray) ToSecretBackendSignArrayOutputWithContext(ctx context.Context) SecretBackendSignArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecretBackendSignArrayOutput)
}

// SecretBackendSignMapInput is an input type that accepts SecretBackendSignMap and SecretBackendSignMapOutput values.
// You can construct a concrete instance of `SecretBackendSignMapInput` via:
//
//	SecretBackendSignMap{ "key": SecretBackendSignArgs{...} }
type SecretBackendSignMapInput interface {
	pulumi.Input

	ToSecretBackendSignMapOutput() SecretBackendSignMapOutput
	ToSecretBackendSignMapOutputWithContext(context.Context) SecretBackendSignMapOutput
}

type SecretBackendSignMap map[string]SecretBackendSignInput

func (SecretBackendSignMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SecretBackendSign)(nil)).Elem()
}

func (i SecretBackendSignMap) ToSecretBackendSignMapOutput() SecretBackendSignMapOutput {
	return i.ToSecretBackendSignMapOutputWithContext(context.Background())
}

func (i SecretBackendSignMap) ToSecretBackendSignMapOutputWithContext(ctx context.Context) SecretBackendSignMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecretBackendSignMapOutput)
}

type SecretBackendSignOutput struct{ *pulumi.OutputState }

func (SecretBackendSignOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SecretBackendSign)(nil)).Elem()
}

func (o SecretBackendSignOutput) ToSecretBackendSignOutput() SecretBackendSignOutput {
	return o
}

func (o SecretBackendSignOutput) ToSecretBackendSignOutputWithContext(ctx context.Context) SecretBackendSignOutput {
	return o
}

// List of alternative names
func (o SecretBackendSignOutput) AltNames() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *SecretBackendSign) pulumi.StringArrayOutput { return v.AltNames }).(pulumi.StringArrayOutput)
}

// If set to `true`, certs will be renewed if the expiration is within `minSecondsRemaining`. Default `false`
func (o SecretBackendSignOutput) AutoRenew() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SecretBackendSign) pulumi.BoolPtrOutput { return v.AutoRenew }).(pulumi.BoolPtrOutput)
}

// The PKI secret backend the resource belongs to.
func (o SecretBackendSignOutput) Backend() pulumi.StringOutput {
	return o.ApplyT(func(v *SecretBackendSign) pulumi.StringOutput { return v.Backend }).(pulumi.StringOutput)
}

// The CA chain
func (o SecretBackendSignOutput) CaChains() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *SecretBackendSign) pulumi.StringArrayOutput { return v.CaChains }).(pulumi.StringArrayOutput)
}

// The certificate
func (o SecretBackendSignOutput) Certificate() pulumi.StringOutput {
	return o.ApplyT(func(v *SecretBackendSign) pulumi.StringOutput { return v.Certificate }).(pulumi.StringOutput)
}

// CN of certificate to create
func (o SecretBackendSignOutput) CommonName() pulumi.StringOutput {
	return o.ApplyT(func(v *SecretBackendSign) pulumi.StringOutput { return v.CommonName }).(pulumi.StringOutput)
}

// The CSR
func (o SecretBackendSignOutput) Csr() pulumi.StringOutput {
	return o.ApplyT(func(v *SecretBackendSign) pulumi.StringOutput { return v.Csr }).(pulumi.StringOutput)
}

// Flag to exclude CN from SANs
func (o SecretBackendSignOutput) ExcludeCnFromSans() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SecretBackendSign) pulumi.BoolPtrOutput { return v.ExcludeCnFromSans }).(pulumi.BoolPtrOutput)
}

// The expiration date of the certificate in unix epoch format
func (o SecretBackendSignOutput) Expiration() pulumi.IntOutput {
	return o.ApplyT(func(v *SecretBackendSign) pulumi.IntOutput { return v.Expiration }).(pulumi.IntOutput)
}

// The format of data
func (o SecretBackendSignOutput) Format() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecretBackendSign) pulumi.StringPtrOutput { return v.Format }).(pulumi.StringPtrOutput)
}

// List of alternative IPs
func (o SecretBackendSignOutput) IpSans() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *SecretBackendSign) pulumi.StringArrayOutput { return v.IpSans }).(pulumi.StringArrayOutput)
}

// Specifies the default issuer of this request. Can
// be the value `default`, a name, or an issuer ID. Use ACLs to prevent access to
// the `/pki/issuer/:issuer_ref/{issue,sign}/:name` paths to prevent users
// overriding the role's `issuerRef` value.
func (o SecretBackendSignOutput) IssuerRef() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecretBackendSign) pulumi.StringPtrOutput { return v.IssuerRef }).(pulumi.StringPtrOutput)
}

// The issuing CA
func (o SecretBackendSignOutput) IssuingCa() pulumi.StringOutput {
	return o.ApplyT(func(v *SecretBackendSign) pulumi.StringOutput { return v.IssuingCa }).(pulumi.StringOutput)
}

// Generate a new certificate when the expiration is within this number of seconds, default is 604800 (7 days)
func (o SecretBackendSignOutput) MinSecondsRemaining() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SecretBackendSign) pulumi.IntPtrOutput { return v.MinSecondsRemaining }).(pulumi.IntPtrOutput)
}

// Name of the role to create the certificate against
func (o SecretBackendSignOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SecretBackendSign) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The namespace to provision the resource in.
// The value should not contain leading or trailing forward slashes.
// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
// *Available only for Vault Enterprise*.
func (o SecretBackendSignOutput) Namespace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecretBackendSign) pulumi.StringPtrOutput { return v.Namespace }).(pulumi.StringPtrOutput)
}

// List of other SANs
func (o SecretBackendSignOutput) OtherSans() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *SecretBackendSign) pulumi.StringArrayOutput { return v.OtherSans }).(pulumi.StringArrayOutput)
}

// `true` if the current time (during refresh) is after the start of the early renewal window declared by `minSecondsRemaining`, and `false` otherwise; if `autoRenew` is set to `true` then the provider will plan to replace the certificate once renewal is pending.
func (o SecretBackendSignOutput) RenewPending() pulumi.BoolOutput {
	return o.ApplyT(func(v *SecretBackendSign) pulumi.BoolOutput { return v.RenewPending }).(pulumi.BoolOutput)
}

// The certificate's serial number, hex formatted.
func (o SecretBackendSignOutput) SerialNumber() pulumi.StringOutput {
	return o.ApplyT(func(v *SecretBackendSign) pulumi.StringOutput { return v.SerialNumber }).(pulumi.StringOutput)
}

// Time to live
func (o SecretBackendSignOutput) Ttl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecretBackendSign) pulumi.StringPtrOutput { return v.Ttl }).(pulumi.StringPtrOutput)
}

// List of alternative URIs
func (o SecretBackendSignOutput) UriSans() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *SecretBackendSign) pulumi.StringArrayOutput { return v.UriSans }).(pulumi.StringArrayOutput)
}

type SecretBackendSignArrayOutput struct{ *pulumi.OutputState }

func (SecretBackendSignArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SecretBackendSign)(nil)).Elem()
}

func (o SecretBackendSignArrayOutput) ToSecretBackendSignArrayOutput() SecretBackendSignArrayOutput {
	return o
}

func (o SecretBackendSignArrayOutput) ToSecretBackendSignArrayOutputWithContext(ctx context.Context) SecretBackendSignArrayOutput {
	return o
}

func (o SecretBackendSignArrayOutput) Index(i pulumi.IntInput) SecretBackendSignOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SecretBackendSign {
		return vs[0].([]*SecretBackendSign)[vs[1].(int)]
	}).(SecretBackendSignOutput)
}

type SecretBackendSignMapOutput struct{ *pulumi.OutputState }

func (SecretBackendSignMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SecretBackendSign)(nil)).Elem()
}

func (o SecretBackendSignMapOutput) ToSecretBackendSignMapOutput() SecretBackendSignMapOutput {
	return o
}

func (o SecretBackendSignMapOutput) ToSecretBackendSignMapOutputWithContext(ctx context.Context) SecretBackendSignMapOutput {
	return o
}

func (o SecretBackendSignMapOutput) MapIndex(k pulumi.StringInput) SecretBackendSignOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SecretBackendSign {
		return vs[0].(map[string]*SecretBackendSign)[vs[1].(string)]
	}).(SecretBackendSignOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SecretBackendSignInput)(nil)).Elem(), &SecretBackendSign{})
	pulumi.RegisterInputType(reflect.TypeOf((*SecretBackendSignArrayInput)(nil)).Elem(), SecretBackendSignArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SecretBackendSignMapInput)(nil)).Elem(), SecretBackendSignMap{})
	pulumi.RegisterOutputType(SecretBackendSignOutput{})
	pulumi.RegisterOutputType(SecretBackendSignArrayOutput{})
	pulumi.RegisterOutputType(SecretBackendSignMapOutput{})
}
