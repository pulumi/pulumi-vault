// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package pkisecret

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type SecretBackendRootCert struct {
	pulumi.CustomResourceState

	// List of alternative names
	AltNames pulumi.StringArrayOutput `pulumi:"altNames"`
	// The PKI secret backend the resource belongs to.
	Backend pulumi.StringOutput `pulumi:"backend"`
	// The certificate
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
	// The issuing CA
	IssuingCa pulumi.StringOutput `pulumi:"issuingCa"`
	// The number of bits to use
	KeyBits pulumi.IntPtrOutput `pulumi:"keyBits"`
	// The desired key type
	KeyType pulumi.StringPtrOutput `pulumi:"keyType"`
	// The locality
	Locality pulumi.StringPtrOutput `pulumi:"locality"`
	// The maximum path length to encode in the generated certificate
	MaxPathLength pulumi.IntPtrOutput `pulumi:"maxPathLength"`
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
	// The serial
	Serial pulumi.StringOutput `pulumi:"serial"`
	// The street address
	StreetAddress pulumi.StringPtrOutput `pulumi:"streetAddress"`
	// Time to live
	Ttl pulumi.StringPtrOutput `pulumi:"ttl"`
	// Type of intermediate to create. Must be either \"exported\" or \"internal\"
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
	// The certificate
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
	// The issuing CA
	IssuingCa *string `pulumi:"issuingCa"`
	// The number of bits to use
	KeyBits *int `pulumi:"keyBits"`
	// The desired key type
	KeyType *string `pulumi:"keyType"`
	// The locality
	Locality *string `pulumi:"locality"`
	// The maximum path length to encode in the generated certificate
	MaxPathLength *int `pulumi:"maxPathLength"`
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
	// The serial
	Serial *string `pulumi:"serial"`
	// The street address
	StreetAddress *string `pulumi:"streetAddress"`
	// Time to live
	Ttl *string `pulumi:"ttl"`
	// Type of intermediate to create. Must be either \"exported\" or \"internal\"
	Type *string `pulumi:"type"`
	// List of alternative URIs
	UriSans []string `pulumi:"uriSans"`
}

type SecretBackendRootCertState struct {
	// List of alternative names
	AltNames pulumi.StringArrayInput
	// The PKI secret backend the resource belongs to.
	Backend pulumi.StringPtrInput
	// The certificate
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
	// The issuing CA
	IssuingCa pulumi.StringPtrInput
	// The number of bits to use
	KeyBits pulumi.IntPtrInput
	// The desired key type
	KeyType pulumi.StringPtrInput
	// The locality
	Locality pulumi.StringPtrInput
	// The maximum path length to encode in the generated certificate
	MaxPathLength pulumi.IntPtrInput
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
	// The serial
	Serial pulumi.StringPtrInput
	// The street address
	StreetAddress pulumi.StringPtrInput
	// Time to live
	Ttl pulumi.StringPtrInput
	// Type of intermediate to create. Must be either \"exported\" or \"internal\"
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
	// The number of bits to use
	KeyBits *int `pulumi:"keyBits"`
	// The desired key type
	KeyType *string `pulumi:"keyType"`
	// The locality
	Locality *string `pulumi:"locality"`
	// The maximum path length to encode in the generated certificate
	MaxPathLength *int `pulumi:"maxPathLength"`
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
	// Type of intermediate to create. Must be either \"exported\" or \"internal\"
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
	// The number of bits to use
	KeyBits pulumi.IntPtrInput
	// The desired key type
	KeyType pulumi.StringPtrInput
	// The locality
	Locality pulumi.StringPtrInput
	// The maximum path length to encode in the generated certificate
	MaxPathLength pulumi.IntPtrInput
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
	// Type of intermediate to create. Must be either \"exported\" or \"internal\"
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
	return reflect.TypeOf((*SecretBackendRootCert)(nil))
}

func (i *SecretBackendRootCert) ToSecretBackendRootCertOutput() SecretBackendRootCertOutput {
	return i.ToSecretBackendRootCertOutputWithContext(context.Background())
}

func (i *SecretBackendRootCert) ToSecretBackendRootCertOutputWithContext(ctx context.Context) SecretBackendRootCertOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecretBackendRootCertOutput)
}

func (i *SecretBackendRootCert) ToSecretBackendRootCertPtrOutput() SecretBackendRootCertPtrOutput {
	return i.ToSecretBackendRootCertPtrOutputWithContext(context.Background())
}

func (i *SecretBackendRootCert) ToSecretBackendRootCertPtrOutputWithContext(ctx context.Context) SecretBackendRootCertPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecretBackendRootCertPtrOutput)
}

type SecretBackendRootCertPtrInput interface {
	pulumi.Input

	ToSecretBackendRootCertPtrOutput() SecretBackendRootCertPtrOutput
	ToSecretBackendRootCertPtrOutputWithContext(ctx context.Context) SecretBackendRootCertPtrOutput
}

type secretBackendRootCertPtrType SecretBackendRootCertArgs

func (*secretBackendRootCertPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SecretBackendRootCert)(nil))
}

func (i *secretBackendRootCertPtrType) ToSecretBackendRootCertPtrOutput() SecretBackendRootCertPtrOutput {
	return i.ToSecretBackendRootCertPtrOutputWithContext(context.Background())
}

func (i *secretBackendRootCertPtrType) ToSecretBackendRootCertPtrOutputWithContext(ctx context.Context) SecretBackendRootCertPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecretBackendRootCertPtrOutput)
}

// SecretBackendRootCertArrayInput is an input type that accepts SecretBackendRootCertArray and SecretBackendRootCertArrayOutput values.
// You can construct a concrete instance of `SecretBackendRootCertArrayInput` via:
//
//          SecretBackendRootCertArray{ SecretBackendRootCertArgs{...} }
type SecretBackendRootCertArrayInput interface {
	pulumi.Input

	ToSecretBackendRootCertArrayOutput() SecretBackendRootCertArrayOutput
	ToSecretBackendRootCertArrayOutputWithContext(context.Context) SecretBackendRootCertArrayOutput
}

type SecretBackendRootCertArray []SecretBackendRootCertInput

func (SecretBackendRootCertArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*SecretBackendRootCert)(nil))
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
//          SecretBackendRootCertMap{ "key": SecretBackendRootCertArgs{...} }
type SecretBackendRootCertMapInput interface {
	pulumi.Input

	ToSecretBackendRootCertMapOutput() SecretBackendRootCertMapOutput
	ToSecretBackendRootCertMapOutputWithContext(context.Context) SecretBackendRootCertMapOutput
}

type SecretBackendRootCertMap map[string]SecretBackendRootCertInput

func (SecretBackendRootCertMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*SecretBackendRootCert)(nil))
}

func (i SecretBackendRootCertMap) ToSecretBackendRootCertMapOutput() SecretBackendRootCertMapOutput {
	return i.ToSecretBackendRootCertMapOutputWithContext(context.Background())
}

func (i SecretBackendRootCertMap) ToSecretBackendRootCertMapOutputWithContext(ctx context.Context) SecretBackendRootCertMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecretBackendRootCertMapOutput)
}

type SecretBackendRootCertOutput struct {
	*pulumi.OutputState
}

func (SecretBackendRootCertOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SecretBackendRootCert)(nil))
}

func (o SecretBackendRootCertOutput) ToSecretBackendRootCertOutput() SecretBackendRootCertOutput {
	return o
}

func (o SecretBackendRootCertOutput) ToSecretBackendRootCertOutputWithContext(ctx context.Context) SecretBackendRootCertOutput {
	return o
}

func (o SecretBackendRootCertOutput) ToSecretBackendRootCertPtrOutput() SecretBackendRootCertPtrOutput {
	return o.ToSecretBackendRootCertPtrOutputWithContext(context.Background())
}

func (o SecretBackendRootCertOutput) ToSecretBackendRootCertPtrOutputWithContext(ctx context.Context) SecretBackendRootCertPtrOutput {
	return o.ApplyT(func(v SecretBackendRootCert) *SecretBackendRootCert {
		return &v
	}).(SecretBackendRootCertPtrOutput)
}

type SecretBackendRootCertPtrOutput struct {
	*pulumi.OutputState
}

func (SecretBackendRootCertPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SecretBackendRootCert)(nil))
}

func (o SecretBackendRootCertPtrOutput) ToSecretBackendRootCertPtrOutput() SecretBackendRootCertPtrOutput {
	return o
}

func (o SecretBackendRootCertPtrOutput) ToSecretBackendRootCertPtrOutputWithContext(ctx context.Context) SecretBackendRootCertPtrOutput {
	return o
}

type SecretBackendRootCertArrayOutput struct{ *pulumi.OutputState }

func (SecretBackendRootCertArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SecretBackendRootCert)(nil))
}

func (o SecretBackendRootCertArrayOutput) ToSecretBackendRootCertArrayOutput() SecretBackendRootCertArrayOutput {
	return o
}

func (o SecretBackendRootCertArrayOutput) ToSecretBackendRootCertArrayOutputWithContext(ctx context.Context) SecretBackendRootCertArrayOutput {
	return o
}

func (o SecretBackendRootCertArrayOutput) Index(i pulumi.IntInput) SecretBackendRootCertOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SecretBackendRootCert {
		return vs[0].([]SecretBackendRootCert)[vs[1].(int)]
	}).(SecretBackendRootCertOutput)
}

type SecretBackendRootCertMapOutput struct{ *pulumi.OutputState }

func (SecretBackendRootCertMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]SecretBackendRootCert)(nil))
}

func (o SecretBackendRootCertMapOutput) ToSecretBackendRootCertMapOutput() SecretBackendRootCertMapOutput {
	return o
}

func (o SecretBackendRootCertMapOutput) ToSecretBackendRootCertMapOutputWithContext(ctx context.Context) SecretBackendRootCertMapOutput {
	return o
}

func (o SecretBackendRootCertMapOutput) MapIndex(k pulumi.StringInput) SecretBackendRootCertOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) SecretBackendRootCert {
		return vs[0].(map[string]SecretBackendRootCert)[vs[1].(string)]
	}).(SecretBackendRootCertOutput)
}

func init() {
	pulumi.RegisterOutputType(SecretBackendRootCertOutput{})
	pulumi.RegisterOutputType(SecretBackendRootCertPtrOutput{})
	pulumi.RegisterOutputType(SecretBackendRootCertArrayOutput{})
	pulumi.RegisterOutputType(SecretBackendRootCertMapOutput{})
}
