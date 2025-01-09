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
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-vault/sdk/v6/go/vault/pkisecret"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := pkisecret.NewSecretBackendCert(ctx, "app", &pkisecret.SecretBackendCertArgs{
//				Backend:    pulumi.Any(intermediate.Path),
//				Name:       pulumi.Any(test.Name),
//				CommonName: pulumi.String("app.my.domain"),
//			}, pulumi.DependsOn([]pulumi.Resource{
//				admin,
//			}))
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
type SecretBackendCert struct {
	pulumi.CustomResourceState

	// List of alternative names
	AltNames pulumi.StringArrayOutput `pulumi:"altNames"`
	// If set to `true`, certs will be renewed if the expiration is within `minSecondsRemaining`. Default `false`
	AutoRenew pulumi.BoolPtrOutput `pulumi:"autoRenew"`
	// The PKI secret backend the resource belongs to.
	Backend pulumi.StringOutput `pulumi:"backend"`
	// The CA chain
	CaChain pulumi.StringOutput `pulumi:"caChain"`
	// The certificate
	Certificate pulumi.StringOutput `pulumi:"certificate"`
	// CN of certificate to create
	CommonName pulumi.StringOutput `pulumi:"commonName"`
	// Flag to exclude CN from SANs
	ExcludeCnFromSans pulumi.BoolPtrOutput `pulumi:"excludeCnFromSans"`
	// The expiration date of the certificate in unix epoch format
	Expiration pulumi.IntOutput `pulumi:"expiration"`
	// The format of data
	Format pulumi.StringPtrOutput `pulumi:"format"`
	// List of alternative IPs
	IpSans pulumi.StringArrayOutput `pulumi:"ipSans"`
	// Specifies the default issuer of this request.
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
	// The private key
	PrivateKey pulumi.StringOutput `pulumi:"privateKey"`
	// The private key format
	PrivateKeyFormat pulumi.StringPtrOutput `pulumi:"privateKeyFormat"`
	// The private key type
	PrivateKeyType pulumi.StringOutput `pulumi:"privateKeyType"`
	// `true` if the current time (during refresh) is after the start of the early renewal window declared by `minSecondsRemaining`, and `false` otherwise; if `autoRenew` is set to `true` then the provider will plan to replace the certificate once renewal is pending.
	RenewPending pulumi.BoolOutput `pulumi:"renewPending"`
	// If set to `true`, the certificate will be revoked on resource destruction.
	Revoke pulumi.BoolPtrOutput `pulumi:"revoke"`
	// The serial number
	SerialNumber pulumi.StringOutput `pulumi:"serialNumber"`
	// Time to live
	Ttl pulumi.StringPtrOutput `pulumi:"ttl"`
	// List of alternative URIs
	UriSans pulumi.StringArrayOutput `pulumi:"uriSans"`
	// List of Subject User IDs
	UserIds pulumi.StringArrayOutput `pulumi:"userIds"`
}

// NewSecretBackendCert registers a new resource with the given unique name, arguments, and options.
func NewSecretBackendCert(ctx *pulumi.Context,
	name string, args *SecretBackendCertArgs, opts ...pulumi.ResourceOption) (*SecretBackendCert, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Backend == nil {
		return nil, errors.New("invalid value for required argument 'Backend'")
	}
	if args.CommonName == nil {
		return nil, errors.New("invalid value for required argument 'CommonName'")
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"privateKey",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource SecretBackendCert
	err := ctx.RegisterResource("vault:pkiSecret/secretBackendCert:SecretBackendCert", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSecretBackendCert gets an existing SecretBackendCert resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSecretBackendCert(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SecretBackendCertState, opts ...pulumi.ResourceOption) (*SecretBackendCert, error) {
	var resource SecretBackendCert
	err := ctx.ReadResource("vault:pkiSecret/secretBackendCert:SecretBackendCert", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SecretBackendCert resources.
type secretBackendCertState struct {
	// List of alternative names
	AltNames []string `pulumi:"altNames"`
	// If set to `true`, certs will be renewed if the expiration is within `minSecondsRemaining`. Default `false`
	AutoRenew *bool `pulumi:"autoRenew"`
	// The PKI secret backend the resource belongs to.
	Backend *string `pulumi:"backend"`
	// The CA chain
	CaChain *string `pulumi:"caChain"`
	// The certificate
	Certificate *string `pulumi:"certificate"`
	// CN of certificate to create
	CommonName *string `pulumi:"commonName"`
	// Flag to exclude CN from SANs
	ExcludeCnFromSans *bool `pulumi:"excludeCnFromSans"`
	// The expiration date of the certificate in unix epoch format
	Expiration *int `pulumi:"expiration"`
	// The format of data
	Format *string `pulumi:"format"`
	// List of alternative IPs
	IpSans []string `pulumi:"ipSans"`
	// Specifies the default issuer of this request.
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
	// The private key
	PrivateKey *string `pulumi:"privateKey"`
	// The private key format
	PrivateKeyFormat *string `pulumi:"privateKeyFormat"`
	// The private key type
	PrivateKeyType *string `pulumi:"privateKeyType"`
	// `true` if the current time (during refresh) is after the start of the early renewal window declared by `minSecondsRemaining`, and `false` otherwise; if `autoRenew` is set to `true` then the provider will plan to replace the certificate once renewal is pending.
	RenewPending *bool `pulumi:"renewPending"`
	// If set to `true`, the certificate will be revoked on resource destruction.
	Revoke *bool `pulumi:"revoke"`
	// The serial number
	SerialNumber *string `pulumi:"serialNumber"`
	// Time to live
	Ttl *string `pulumi:"ttl"`
	// List of alternative URIs
	UriSans []string `pulumi:"uriSans"`
	// List of Subject User IDs
	UserIds []string `pulumi:"userIds"`
}

type SecretBackendCertState struct {
	// List of alternative names
	AltNames pulumi.StringArrayInput
	// If set to `true`, certs will be renewed if the expiration is within `minSecondsRemaining`. Default `false`
	AutoRenew pulumi.BoolPtrInput
	// The PKI secret backend the resource belongs to.
	Backend pulumi.StringPtrInput
	// The CA chain
	CaChain pulumi.StringPtrInput
	// The certificate
	Certificate pulumi.StringPtrInput
	// CN of certificate to create
	CommonName pulumi.StringPtrInput
	// Flag to exclude CN from SANs
	ExcludeCnFromSans pulumi.BoolPtrInput
	// The expiration date of the certificate in unix epoch format
	Expiration pulumi.IntPtrInput
	// The format of data
	Format pulumi.StringPtrInput
	// List of alternative IPs
	IpSans pulumi.StringArrayInput
	// Specifies the default issuer of this request.
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
	// The private key
	PrivateKey pulumi.StringPtrInput
	// The private key format
	PrivateKeyFormat pulumi.StringPtrInput
	// The private key type
	PrivateKeyType pulumi.StringPtrInput
	// `true` if the current time (during refresh) is after the start of the early renewal window declared by `minSecondsRemaining`, and `false` otherwise; if `autoRenew` is set to `true` then the provider will plan to replace the certificate once renewal is pending.
	RenewPending pulumi.BoolPtrInput
	// If set to `true`, the certificate will be revoked on resource destruction.
	Revoke pulumi.BoolPtrInput
	// The serial number
	SerialNumber pulumi.StringPtrInput
	// Time to live
	Ttl pulumi.StringPtrInput
	// List of alternative URIs
	UriSans pulumi.StringArrayInput
	// List of Subject User IDs
	UserIds pulumi.StringArrayInput
}

func (SecretBackendCertState) ElementType() reflect.Type {
	return reflect.TypeOf((*secretBackendCertState)(nil)).Elem()
}

type secretBackendCertArgs struct {
	// List of alternative names
	AltNames []string `pulumi:"altNames"`
	// If set to `true`, certs will be renewed if the expiration is within `minSecondsRemaining`. Default `false`
	AutoRenew *bool `pulumi:"autoRenew"`
	// The PKI secret backend the resource belongs to.
	Backend string `pulumi:"backend"`
	// CN of certificate to create
	CommonName string `pulumi:"commonName"`
	// Flag to exclude CN from SANs
	ExcludeCnFromSans *bool `pulumi:"excludeCnFromSans"`
	// The format of data
	Format *string `pulumi:"format"`
	// List of alternative IPs
	IpSans []string `pulumi:"ipSans"`
	// Specifies the default issuer of this request.
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
	// The private key format
	PrivateKeyFormat *string `pulumi:"privateKeyFormat"`
	// If set to `true`, the certificate will be revoked on resource destruction.
	Revoke *bool `pulumi:"revoke"`
	// Time to live
	Ttl *string `pulumi:"ttl"`
	// List of alternative URIs
	UriSans []string `pulumi:"uriSans"`
	// List of Subject User IDs
	UserIds []string `pulumi:"userIds"`
}

// The set of arguments for constructing a SecretBackendCert resource.
type SecretBackendCertArgs struct {
	// List of alternative names
	AltNames pulumi.StringArrayInput
	// If set to `true`, certs will be renewed if the expiration is within `minSecondsRemaining`. Default `false`
	AutoRenew pulumi.BoolPtrInput
	// The PKI secret backend the resource belongs to.
	Backend pulumi.StringInput
	// CN of certificate to create
	CommonName pulumi.StringInput
	// Flag to exclude CN from SANs
	ExcludeCnFromSans pulumi.BoolPtrInput
	// The format of data
	Format pulumi.StringPtrInput
	// List of alternative IPs
	IpSans pulumi.StringArrayInput
	// Specifies the default issuer of this request.
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
	// The private key format
	PrivateKeyFormat pulumi.StringPtrInput
	// If set to `true`, the certificate will be revoked on resource destruction.
	Revoke pulumi.BoolPtrInput
	// Time to live
	Ttl pulumi.StringPtrInput
	// List of alternative URIs
	UriSans pulumi.StringArrayInput
	// List of Subject User IDs
	UserIds pulumi.StringArrayInput
}

func (SecretBackendCertArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*secretBackendCertArgs)(nil)).Elem()
}

type SecretBackendCertInput interface {
	pulumi.Input

	ToSecretBackendCertOutput() SecretBackendCertOutput
	ToSecretBackendCertOutputWithContext(ctx context.Context) SecretBackendCertOutput
}

func (*SecretBackendCert) ElementType() reflect.Type {
	return reflect.TypeOf((**SecretBackendCert)(nil)).Elem()
}

func (i *SecretBackendCert) ToSecretBackendCertOutput() SecretBackendCertOutput {
	return i.ToSecretBackendCertOutputWithContext(context.Background())
}

func (i *SecretBackendCert) ToSecretBackendCertOutputWithContext(ctx context.Context) SecretBackendCertOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecretBackendCertOutput)
}

// SecretBackendCertArrayInput is an input type that accepts SecretBackendCertArray and SecretBackendCertArrayOutput values.
// You can construct a concrete instance of `SecretBackendCertArrayInput` via:
//
//	SecretBackendCertArray{ SecretBackendCertArgs{...} }
type SecretBackendCertArrayInput interface {
	pulumi.Input

	ToSecretBackendCertArrayOutput() SecretBackendCertArrayOutput
	ToSecretBackendCertArrayOutputWithContext(context.Context) SecretBackendCertArrayOutput
}

type SecretBackendCertArray []SecretBackendCertInput

func (SecretBackendCertArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SecretBackendCert)(nil)).Elem()
}

func (i SecretBackendCertArray) ToSecretBackendCertArrayOutput() SecretBackendCertArrayOutput {
	return i.ToSecretBackendCertArrayOutputWithContext(context.Background())
}

func (i SecretBackendCertArray) ToSecretBackendCertArrayOutputWithContext(ctx context.Context) SecretBackendCertArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecretBackendCertArrayOutput)
}

// SecretBackendCertMapInput is an input type that accepts SecretBackendCertMap and SecretBackendCertMapOutput values.
// You can construct a concrete instance of `SecretBackendCertMapInput` via:
//
//	SecretBackendCertMap{ "key": SecretBackendCertArgs{...} }
type SecretBackendCertMapInput interface {
	pulumi.Input

	ToSecretBackendCertMapOutput() SecretBackendCertMapOutput
	ToSecretBackendCertMapOutputWithContext(context.Context) SecretBackendCertMapOutput
}

type SecretBackendCertMap map[string]SecretBackendCertInput

func (SecretBackendCertMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SecretBackendCert)(nil)).Elem()
}

func (i SecretBackendCertMap) ToSecretBackendCertMapOutput() SecretBackendCertMapOutput {
	return i.ToSecretBackendCertMapOutputWithContext(context.Background())
}

func (i SecretBackendCertMap) ToSecretBackendCertMapOutputWithContext(ctx context.Context) SecretBackendCertMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecretBackendCertMapOutput)
}

type SecretBackendCertOutput struct{ *pulumi.OutputState }

func (SecretBackendCertOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SecretBackendCert)(nil)).Elem()
}

func (o SecretBackendCertOutput) ToSecretBackendCertOutput() SecretBackendCertOutput {
	return o
}

func (o SecretBackendCertOutput) ToSecretBackendCertOutputWithContext(ctx context.Context) SecretBackendCertOutput {
	return o
}

// List of alternative names
func (o SecretBackendCertOutput) AltNames() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *SecretBackendCert) pulumi.StringArrayOutput { return v.AltNames }).(pulumi.StringArrayOutput)
}

// If set to `true`, certs will be renewed if the expiration is within `minSecondsRemaining`. Default `false`
func (o SecretBackendCertOutput) AutoRenew() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SecretBackendCert) pulumi.BoolPtrOutput { return v.AutoRenew }).(pulumi.BoolPtrOutput)
}

// The PKI secret backend the resource belongs to.
func (o SecretBackendCertOutput) Backend() pulumi.StringOutput {
	return o.ApplyT(func(v *SecretBackendCert) pulumi.StringOutput { return v.Backend }).(pulumi.StringOutput)
}

// The CA chain
func (o SecretBackendCertOutput) CaChain() pulumi.StringOutput {
	return o.ApplyT(func(v *SecretBackendCert) pulumi.StringOutput { return v.CaChain }).(pulumi.StringOutput)
}

// The certificate
func (o SecretBackendCertOutput) Certificate() pulumi.StringOutput {
	return o.ApplyT(func(v *SecretBackendCert) pulumi.StringOutput { return v.Certificate }).(pulumi.StringOutput)
}

// CN of certificate to create
func (o SecretBackendCertOutput) CommonName() pulumi.StringOutput {
	return o.ApplyT(func(v *SecretBackendCert) pulumi.StringOutput { return v.CommonName }).(pulumi.StringOutput)
}

// Flag to exclude CN from SANs
func (o SecretBackendCertOutput) ExcludeCnFromSans() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SecretBackendCert) pulumi.BoolPtrOutput { return v.ExcludeCnFromSans }).(pulumi.BoolPtrOutput)
}

// The expiration date of the certificate in unix epoch format
func (o SecretBackendCertOutput) Expiration() pulumi.IntOutput {
	return o.ApplyT(func(v *SecretBackendCert) pulumi.IntOutput { return v.Expiration }).(pulumi.IntOutput)
}

// The format of data
func (o SecretBackendCertOutput) Format() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecretBackendCert) pulumi.StringPtrOutput { return v.Format }).(pulumi.StringPtrOutput)
}

// List of alternative IPs
func (o SecretBackendCertOutput) IpSans() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *SecretBackendCert) pulumi.StringArrayOutput { return v.IpSans }).(pulumi.StringArrayOutput)
}

// Specifies the default issuer of this request.
func (o SecretBackendCertOutput) IssuerRef() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecretBackendCert) pulumi.StringPtrOutput { return v.IssuerRef }).(pulumi.StringPtrOutput)
}

// The issuing CA
func (o SecretBackendCertOutput) IssuingCa() pulumi.StringOutput {
	return o.ApplyT(func(v *SecretBackendCert) pulumi.StringOutput { return v.IssuingCa }).(pulumi.StringOutput)
}

// Generate a new certificate when the expiration is within this number of seconds, default is 604800 (7 days)
func (o SecretBackendCertOutput) MinSecondsRemaining() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SecretBackendCert) pulumi.IntPtrOutput { return v.MinSecondsRemaining }).(pulumi.IntPtrOutput)
}

// Name of the role to create the certificate against
func (o SecretBackendCertOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SecretBackendCert) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The namespace to provision the resource in.
// The value should not contain leading or trailing forward slashes.
// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
// *Available only for Vault Enterprise*.
func (o SecretBackendCertOutput) Namespace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecretBackendCert) pulumi.StringPtrOutput { return v.Namespace }).(pulumi.StringPtrOutput)
}

// List of other SANs
func (o SecretBackendCertOutput) OtherSans() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *SecretBackendCert) pulumi.StringArrayOutput { return v.OtherSans }).(pulumi.StringArrayOutput)
}

// The private key
func (o SecretBackendCertOutput) PrivateKey() pulumi.StringOutput {
	return o.ApplyT(func(v *SecretBackendCert) pulumi.StringOutput { return v.PrivateKey }).(pulumi.StringOutput)
}

// The private key format
func (o SecretBackendCertOutput) PrivateKeyFormat() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecretBackendCert) pulumi.StringPtrOutput { return v.PrivateKeyFormat }).(pulumi.StringPtrOutput)
}

// The private key type
func (o SecretBackendCertOutput) PrivateKeyType() pulumi.StringOutput {
	return o.ApplyT(func(v *SecretBackendCert) pulumi.StringOutput { return v.PrivateKeyType }).(pulumi.StringOutput)
}

// `true` if the current time (during refresh) is after the start of the early renewal window declared by `minSecondsRemaining`, and `false` otherwise; if `autoRenew` is set to `true` then the provider will plan to replace the certificate once renewal is pending.
func (o SecretBackendCertOutput) RenewPending() pulumi.BoolOutput {
	return o.ApplyT(func(v *SecretBackendCert) pulumi.BoolOutput { return v.RenewPending }).(pulumi.BoolOutput)
}

// If set to `true`, the certificate will be revoked on resource destruction.
func (o SecretBackendCertOutput) Revoke() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SecretBackendCert) pulumi.BoolPtrOutput { return v.Revoke }).(pulumi.BoolPtrOutput)
}

// The serial number
func (o SecretBackendCertOutput) SerialNumber() pulumi.StringOutput {
	return o.ApplyT(func(v *SecretBackendCert) pulumi.StringOutput { return v.SerialNumber }).(pulumi.StringOutput)
}

// Time to live
func (o SecretBackendCertOutput) Ttl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecretBackendCert) pulumi.StringPtrOutput { return v.Ttl }).(pulumi.StringPtrOutput)
}

// List of alternative URIs
func (o SecretBackendCertOutput) UriSans() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *SecretBackendCert) pulumi.StringArrayOutput { return v.UriSans }).(pulumi.StringArrayOutput)
}

// List of Subject User IDs
func (o SecretBackendCertOutput) UserIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *SecretBackendCert) pulumi.StringArrayOutput { return v.UserIds }).(pulumi.StringArrayOutput)
}

type SecretBackendCertArrayOutput struct{ *pulumi.OutputState }

func (SecretBackendCertArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SecretBackendCert)(nil)).Elem()
}

func (o SecretBackendCertArrayOutput) ToSecretBackendCertArrayOutput() SecretBackendCertArrayOutput {
	return o
}

func (o SecretBackendCertArrayOutput) ToSecretBackendCertArrayOutputWithContext(ctx context.Context) SecretBackendCertArrayOutput {
	return o
}

func (o SecretBackendCertArrayOutput) Index(i pulumi.IntInput) SecretBackendCertOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SecretBackendCert {
		return vs[0].([]*SecretBackendCert)[vs[1].(int)]
	}).(SecretBackendCertOutput)
}

type SecretBackendCertMapOutput struct{ *pulumi.OutputState }

func (SecretBackendCertMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SecretBackendCert)(nil)).Elem()
}

func (o SecretBackendCertMapOutput) ToSecretBackendCertMapOutput() SecretBackendCertMapOutput {
	return o
}

func (o SecretBackendCertMapOutput) ToSecretBackendCertMapOutputWithContext(ctx context.Context) SecretBackendCertMapOutput {
	return o
}

func (o SecretBackendCertMapOutput) MapIndex(k pulumi.StringInput) SecretBackendCertOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SecretBackendCert {
		return vs[0].(map[string]*SecretBackendCert)[vs[1].(string)]
	}).(SecretBackendCertOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SecretBackendCertInput)(nil)).Elem(), &SecretBackendCert{})
	pulumi.RegisterInputType(reflect.TypeOf((*SecretBackendCertArrayInput)(nil)).Elem(), SecretBackendCertArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SecretBackendCertMapInput)(nil)).Elem(), SecretBackendCertMap{})
	pulumi.RegisterOutputType(SecretBackendCertOutput{})
	pulumi.RegisterOutputType(SecretBackendCertArrayOutput{})
	pulumi.RegisterOutputType(SecretBackendCertMapOutput{})
}
