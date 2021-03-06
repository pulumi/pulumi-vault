// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package pkisecret

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

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
	// The issuing CA
	IssuingCa pulumi.StringOutput `pulumi:"issuingCa"`
	// Generate a new certificate when the expiration is within this number of seconds, default is 604800 (7 days)
	MinSecondsRemaining pulumi.IntPtrOutput `pulumi:"minSecondsRemaining"`
	// Name of the role to create the certificate against
	Name pulumi.StringOutput `pulumi:"name"`
	// List of other SANs
	OtherSans pulumi.StringArrayOutput `pulumi:"otherSans"`
	// The private key
	PrivateKey pulumi.StringOutput `pulumi:"privateKey"`
	// The private key format
	PrivateKeyFormat pulumi.StringPtrOutput `pulumi:"privateKeyFormat"`
	// The private key type
	PrivateKeyType pulumi.StringOutput `pulumi:"privateKeyType"`
	// The serial number
	SerialNumber pulumi.StringOutput `pulumi:"serialNumber"`
	// Time to live
	Ttl pulumi.StringPtrOutput `pulumi:"ttl"`
	// List of alternative URIs
	UriSans pulumi.StringArrayOutput `pulumi:"uriSans"`
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
	// The issuing CA
	IssuingCa *string `pulumi:"issuingCa"`
	// Generate a new certificate when the expiration is within this number of seconds, default is 604800 (7 days)
	MinSecondsRemaining *int `pulumi:"minSecondsRemaining"`
	// Name of the role to create the certificate against
	Name *string `pulumi:"name"`
	// List of other SANs
	OtherSans []string `pulumi:"otherSans"`
	// The private key
	PrivateKey *string `pulumi:"privateKey"`
	// The private key format
	PrivateKeyFormat *string `pulumi:"privateKeyFormat"`
	// The private key type
	PrivateKeyType *string `pulumi:"privateKeyType"`
	// The serial number
	SerialNumber *string `pulumi:"serialNumber"`
	// Time to live
	Ttl *string `pulumi:"ttl"`
	// List of alternative URIs
	UriSans []string `pulumi:"uriSans"`
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
	// The issuing CA
	IssuingCa pulumi.StringPtrInput
	// Generate a new certificate when the expiration is within this number of seconds, default is 604800 (7 days)
	MinSecondsRemaining pulumi.IntPtrInput
	// Name of the role to create the certificate against
	Name pulumi.StringPtrInput
	// List of other SANs
	OtherSans pulumi.StringArrayInput
	// The private key
	PrivateKey pulumi.StringPtrInput
	// The private key format
	PrivateKeyFormat pulumi.StringPtrInput
	// The private key type
	PrivateKeyType pulumi.StringPtrInput
	// The serial number
	SerialNumber pulumi.StringPtrInput
	// Time to live
	Ttl pulumi.StringPtrInput
	// List of alternative URIs
	UriSans pulumi.StringArrayInput
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
	// Generate a new certificate when the expiration is within this number of seconds, default is 604800 (7 days)
	MinSecondsRemaining *int `pulumi:"minSecondsRemaining"`
	// Name of the role to create the certificate against
	Name *string `pulumi:"name"`
	// List of other SANs
	OtherSans []string `pulumi:"otherSans"`
	// The private key format
	PrivateKeyFormat *string `pulumi:"privateKeyFormat"`
	// Time to live
	Ttl *string `pulumi:"ttl"`
	// List of alternative URIs
	UriSans []string `pulumi:"uriSans"`
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
	// Generate a new certificate when the expiration is within this number of seconds, default is 604800 (7 days)
	MinSecondsRemaining pulumi.IntPtrInput
	// Name of the role to create the certificate against
	Name pulumi.StringPtrInput
	// List of other SANs
	OtherSans pulumi.StringArrayInput
	// The private key format
	PrivateKeyFormat pulumi.StringPtrInput
	// Time to live
	Ttl pulumi.StringPtrInput
	// List of alternative URIs
	UriSans pulumi.StringArrayInput
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
	return reflect.TypeOf((*SecretBackendCert)(nil))
}

func (i *SecretBackendCert) ToSecretBackendCertOutput() SecretBackendCertOutput {
	return i.ToSecretBackendCertOutputWithContext(context.Background())
}

func (i *SecretBackendCert) ToSecretBackendCertOutputWithContext(ctx context.Context) SecretBackendCertOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecretBackendCertOutput)
}

func (i *SecretBackendCert) ToSecretBackendCertPtrOutput() SecretBackendCertPtrOutput {
	return i.ToSecretBackendCertPtrOutputWithContext(context.Background())
}

func (i *SecretBackendCert) ToSecretBackendCertPtrOutputWithContext(ctx context.Context) SecretBackendCertPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecretBackendCertPtrOutput)
}

type SecretBackendCertPtrInput interface {
	pulumi.Input

	ToSecretBackendCertPtrOutput() SecretBackendCertPtrOutput
	ToSecretBackendCertPtrOutputWithContext(ctx context.Context) SecretBackendCertPtrOutput
}

type secretBackendCertPtrType SecretBackendCertArgs

func (*secretBackendCertPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SecretBackendCert)(nil))
}

func (i *secretBackendCertPtrType) ToSecretBackendCertPtrOutput() SecretBackendCertPtrOutput {
	return i.ToSecretBackendCertPtrOutputWithContext(context.Background())
}

func (i *secretBackendCertPtrType) ToSecretBackendCertPtrOutputWithContext(ctx context.Context) SecretBackendCertPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecretBackendCertPtrOutput)
}

// SecretBackendCertArrayInput is an input type that accepts SecretBackendCertArray and SecretBackendCertArrayOutput values.
// You can construct a concrete instance of `SecretBackendCertArrayInput` via:
//
//          SecretBackendCertArray{ SecretBackendCertArgs{...} }
type SecretBackendCertArrayInput interface {
	pulumi.Input

	ToSecretBackendCertArrayOutput() SecretBackendCertArrayOutput
	ToSecretBackendCertArrayOutputWithContext(context.Context) SecretBackendCertArrayOutput
}

type SecretBackendCertArray []SecretBackendCertInput

func (SecretBackendCertArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*SecretBackendCert)(nil))
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
//          SecretBackendCertMap{ "key": SecretBackendCertArgs{...} }
type SecretBackendCertMapInput interface {
	pulumi.Input

	ToSecretBackendCertMapOutput() SecretBackendCertMapOutput
	ToSecretBackendCertMapOutputWithContext(context.Context) SecretBackendCertMapOutput
}

type SecretBackendCertMap map[string]SecretBackendCertInput

func (SecretBackendCertMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*SecretBackendCert)(nil))
}

func (i SecretBackendCertMap) ToSecretBackendCertMapOutput() SecretBackendCertMapOutput {
	return i.ToSecretBackendCertMapOutputWithContext(context.Background())
}

func (i SecretBackendCertMap) ToSecretBackendCertMapOutputWithContext(ctx context.Context) SecretBackendCertMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecretBackendCertMapOutput)
}

type SecretBackendCertOutput struct {
	*pulumi.OutputState
}

func (SecretBackendCertOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SecretBackendCert)(nil))
}

func (o SecretBackendCertOutput) ToSecretBackendCertOutput() SecretBackendCertOutput {
	return o
}

func (o SecretBackendCertOutput) ToSecretBackendCertOutputWithContext(ctx context.Context) SecretBackendCertOutput {
	return o
}

func (o SecretBackendCertOutput) ToSecretBackendCertPtrOutput() SecretBackendCertPtrOutput {
	return o.ToSecretBackendCertPtrOutputWithContext(context.Background())
}

func (o SecretBackendCertOutput) ToSecretBackendCertPtrOutputWithContext(ctx context.Context) SecretBackendCertPtrOutput {
	return o.ApplyT(func(v SecretBackendCert) *SecretBackendCert {
		return &v
	}).(SecretBackendCertPtrOutput)
}

type SecretBackendCertPtrOutput struct {
	*pulumi.OutputState
}

func (SecretBackendCertPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SecretBackendCert)(nil))
}

func (o SecretBackendCertPtrOutput) ToSecretBackendCertPtrOutput() SecretBackendCertPtrOutput {
	return o
}

func (o SecretBackendCertPtrOutput) ToSecretBackendCertPtrOutputWithContext(ctx context.Context) SecretBackendCertPtrOutput {
	return o
}

type SecretBackendCertArrayOutput struct{ *pulumi.OutputState }

func (SecretBackendCertArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SecretBackendCert)(nil))
}

func (o SecretBackendCertArrayOutput) ToSecretBackendCertArrayOutput() SecretBackendCertArrayOutput {
	return o
}

func (o SecretBackendCertArrayOutput) ToSecretBackendCertArrayOutputWithContext(ctx context.Context) SecretBackendCertArrayOutput {
	return o
}

func (o SecretBackendCertArrayOutput) Index(i pulumi.IntInput) SecretBackendCertOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SecretBackendCert {
		return vs[0].([]SecretBackendCert)[vs[1].(int)]
	}).(SecretBackendCertOutput)
}

type SecretBackendCertMapOutput struct{ *pulumi.OutputState }

func (SecretBackendCertMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]SecretBackendCert)(nil))
}

func (o SecretBackendCertMapOutput) ToSecretBackendCertMapOutput() SecretBackendCertMapOutput {
	return o
}

func (o SecretBackendCertMapOutput) ToSecretBackendCertMapOutputWithContext(ctx context.Context) SecretBackendCertMapOutput {
	return o
}

func (o SecretBackendCertMapOutput) MapIndex(k pulumi.StringInput) SecretBackendCertOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) SecretBackendCert {
		return vs[0].(map[string]SecretBackendCert)[vs[1].(string)]
	}).(SecretBackendCertOutput)
}

func init() {
	pulumi.RegisterOutputType(SecretBackendCertOutput{})
	pulumi.RegisterOutputType(SecretBackendCertPtrOutput{})
	pulumi.RegisterOutputType(SecretBackendCertArrayOutput{})
	pulumi.RegisterOutputType(SecretBackendCertMapOutput{})
}
