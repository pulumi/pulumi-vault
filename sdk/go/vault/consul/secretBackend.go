// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package consul

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## Import
//
// Consul secret backends can be imported using the `path`, e.g.
//
// ```sh
//  $ pulumi import vault:consul/secretBackend:SecretBackend example consul
// ```
type SecretBackend struct {
	pulumi.CustomResourceState

	// Specifies the address of the Consul instance, provided as "host:port" like "127.0.0.1:8500".
	Address pulumi.StringOutput `pulumi:"address"`
	// CA certificate to use when verifying Consul server certificate, must be x509 PEM encoded.
	CaCert pulumi.StringPtrOutput `pulumi:"caCert"`
	// Client certificate used for Consul's TLS communication, must be x509 PEM encoded and if this is set you need to also set client_key.
	ClientCert pulumi.StringPtrOutput `pulumi:"clientCert"`
	// Client key used for Consul's TLS communication, must be x509 PEM encoded and if this is set you need to also set client_cert.
	ClientKey pulumi.StringPtrOutput `pulumi:"clientKey"`
	// The default TTL for credentials issued by this backend.
	DefaultLeaseTtlSeconds pulumi.IntPtrOutput `pulumi:"defaultLeaseTtlSeconds"`
	// A human-friendly description for this backend.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Specifies if the secret backend is local only.
	Local pulumi.BoolPtrOutput `pulumi:"local"`
	// The maximum TTL that can be requested
	// for credentials issued by this backend.
	MaxLeaseTtlSeconds pulumi.IntPtrOutput `pulumi:"maxLeaseTtlSeconds"`
	// The unique location this backend should be mounted at. Must not begin or end with a `/`. Defaults to `consul`.
	Path pulumi.StringPtrOutput `pulumi:"path"`
	// Specifies the URL scheme to use. Defaults to `http`.
	Scheme pulumi.StringPtrOutput `pulumi:"scheme"`
	// The Consul management token this backend should use to issue new tokens.
	Token pulumi.StringOutput `pulumi:"token"`
}

// NewSecretBackend registers a new resource with the given unique name, arguments, and options.
func NewSecretBackend(ctx *pulumi.Context,
	name string, args *SecretBackendArgs, opts ...pulumi.ResourceOption) (*SecretBackend, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Address == nil {
		return nil, errors.New("invalid value for required argument 'Address'")
	}
	if args.Token == nil {
		return nil, errors.New("invalid value for required argument 'Token'")
	}
	var resource SecretBackend
	err := ctx.RegisterResource("vault:consul/secretBackend:SecretBackend", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSecretBackend gets an existing SecretBackend resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSecretBackend(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SecretBackendState, opts ...pulumi.ResourceOption) (*SecretBackend, error) {
	var resource SecretBackend
	err := ctx.ReadResource("vault:consul/secretBackend:SecretBackend", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SecretBackend resources.
type secretBackendState struct {
	// Specifies the address of the Consul instance, provided as "host:port" like "127.0.0.1:8500".
	Address *string `pulumi:"address"`
	// CA certificate to use when verifying Consul server certificate, must be x509 PEM encoded.
	CaCert *string `pulumi:"caCert"`
	// Client certificate used for Consul's TLS communication, must be x509 PEM encoded and if this is set you need to also set client_key.
	ClientCert *string `pulumi:"clientCert"`
	// Client key used for Consul's TLS communication, must be x509 PEM encoded and if this is set you need to also set client_cert.
	ClientKey *string `pulumi:"clientKey"`
	// The default TTL for credentials issued by this backend.
	DefaultLeaseTtlSeconds *int `pulumi:"defaultLeaseTtlSeconds"`
	// A human-friendly description for this backend.
	Description *string `pulumi:"description"`
	// Specifies if the secret backend is local only.
	Local *bool `pulumi:"local"`
	// The maximum TTL that can be requested
	// for credentials issued by this backend.
	MaxLeaseTtlSeconds *int `pulumi:"maxLeaseTtlSeconds"`
	// The unique location this backend should be mounted at. Must not begin or end with a `/`. Defaults to `consul`.
	Path *string `pulumi:"path"`
	// Specifies the URL scheme to use. Defaults to `http`.
	Scheme *string `pulumi:"scheme"`
	// The Consul management token this backend should use to issue new tokens.
	Token *string `pulumi:"token"`
}

type SecretBackendState struct {
	// Specifies the address of the Consul instance, provided as "host:port" like "127.0.0.1:8500".
	Address pulumi.StringPtrInput
	// CA certificate to use when verifying Consul server certificate, must be x509 PEM encoded.
	CaCert pulumi.StringPtrInput
	// Client certificate used for Consul's TLS communication, must be x509 PEM encoded and if this is set you need to also set client_key.
	ClientCert pulumi.StringPtrInput
	// Client key used for Consul's TLS communication, must be x509 PEM encoded and if this is set you need to also set client_cert.
	ClientKey pulumi.StringPtrInput
	// The default TTL for credentials issued by this backend.
	DefaultLeaseTtlSeconds pulumi.IntPtrInput
	// A human-friendly description for this backend.
	Description pulumi.StringPtrInput
	// Specifies if the secret backend is local only.
	Local pulumi.BoolPtrInput
	// The maximum TTL that can be requested
	// for credentials issued by this backend.
	MaxLeaseTtlSeconds pulumi.IntPtrInput
	// The unique location this backend should be mounted at. Must not begin or end with a `/`. Defaults to `consul`.
	Path pulumi.StringPtrInput
	// Specifies the URL scheme to use. Defaults to `http`.
	Scheme pulumi.StringPtrInput
	// The Consul management token this backend should use to issue new tokens.
	Token pulumi.StringPtrInput
}

func (SecretBackendState) ElementType() reflect.Type {
	return reflect.TypeOf((*secretBackendState)(nil)).Elem()
}

type secretBackendArgs struct {
	// Specifies the address of the Consul instance, provided as "host:port" like "127.0.0.1:8500".
	Address string `pulumi:"address"`
	// CA certificate to use when verifying Consul server certificate, must be x509 PEM encoded.
	CaCert *string `pulumi:"caCert"`
	// Client certificate used for Consul's TLS communication, must be x509 PEM encoded and if this is set you need to also set client_key.
	ClientCert *string `pulumi:"clientCert"`
	// Client key used for Consul's TLS communication, must be x509 PEM encoded and if this is set you need to also set client_cert.
	ClientKey *string `pulumi:"clientKey"`
	// The default TTL for credentials issued by this backend.
	DefaultLeaseTtlSeconds *int `pulumi:"defaultLeaseTtlSeconds"`
	// A human-friendly description for this backend.
	Description *string `pulumi:"description"`
	// Specifies if the secret backend is local only.
	Local *bool `pulumi:"local"`
	// The maximum TTL that can be requested
	// for credentials issued by this backend.
	MaxLeaseTtlSeconds *int `pulumi:"maxLeaseTtlSeconds"`
	// The unique location this backend should be mounted at. Must not begin or end with a `/`. Defaults to `consul`.
	Path *string `pulumi:"path"`
	// Specifies the URL scheme to use. Defaults to `http`.
	Scheme *string `pulumi:"scheme"`
	// The Consul management token this backend should use to issue new tokens.
	Token string `pulumi:"token"`
}

// The set of arguments for constructing a SecretBackend resource.
type SecretBackendArgs struct {
	// Specifies the address of the Consul instance, provided as "host:port" like "127.0.0.1:8500".
	Address pulumi.StringInput
	// CA certificate to use when verifying Consul server certificate, must be x509 PEM encoded.
	CaCert pulumi.StringPtrInput
	// Client certificate used for Consul's TLS communication, must be x509 PEM encoded and if this is set you need to also set client_key.
	ClientCert pulumi.StringPtrInput
	// Client key used for Consul's TLS communication, must be x509 PEM encoded and if this is set you need to also set client_cert.
	ClientKey pulumi.StringPtrInput
	// The default TTL for credentials issued by this backend.
	DefaultLeaseTtlSeconds pulumi.IntPtrInput
	// A human-friendly description for this backend.
	Description pulumi.StringPtrInput
	// Specifies if the secret backend is local only.
	Local pulumi.BoolPtrInput
	// The maximum TTL that can be requested
	// for credentials issued by this backend.
	MaxLeaseTtlSeconds pulumi.IntPtrInput
	// The unique location this backend should be mounted at. Must not begin or end with a `/`. Defaults to `consul`.
	Path pulumi.StringPtrInput
	// Specifies the URL scheme to use. Defaults to `http`.
	Scheme pulumi.StringPtrInput
	// The Consul management token this backend should use to issue new tokens.
	Token pulumi.StringInput
}

func (SecretBackendArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*secretBackendArgs)(nil)).Elem()
}

type SecretBackendInput interface {
	pulumi.Input

	ToSecretBackendOutput() SecretBackendOutput
	ToSecretBackendOutputWithContext(ctx context.Context) SecretBackendOutput
}

func (*SecretBackend) ElementType() reflect.Type {
	return reflect.TypeOf((**SecretBackend)(nil)).Elem()
}

func (i *SecretBackend) ToSecretBackendOutput() SecretBackendOutput {
	return i.ToSecretBackendOutputWithContext(context.Background())
}

func (i *SecretBackend) ToSecretBackendOutputWithContext(ctx context.Context) SecretBackendOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecretBackendOutput)
}

// SecretBackendArrayInput is an input type that accepts SecretBackendArray and SecretBackendArrayOutput values.
// You can construct a concrete instance of `SecretBackendArrayInput` via:
//
//          SecretBackendArray{ SecretBackendArgs{...} }
type SecretBackendArrayInput interface {
	pulumi.Input

	ToSecretBackendArrayOutput() SecretBackendArrayOutput
	ToSecretBackendArrayOutputWithContext(context.Context) SecretBackendArrayOutput
}

type SecretBackendArray []SecretBackendInput

func (SecretBackendArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SecretBackend)(nil)).Elem()
}

func (i SecretBackendArray) ToSecretBackendArrayOutput() SecretBackendArrayOutput {
	return i.ToSecretBackendArrayOutputWithContext(context.Background())
}

func (i SecretBackendArray) ToSecretBackendArrayOutputWithContext(ctx context.Context) SecretBackendArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecretBackendArrayOutput)
}

// SecretBackendMapInput is an input type that accepts SecretBackendMap and SecretBackendMapOutput values.
// You can construct a concrete instance of `SecretBackendMapInput` via:
//
//          SecretBackendMap{ "key": SecretBackendArgs{...} }
type SecretBackendMapInput interface {
	pulumi.Input

	ToSecretBackendMapOutput() SecretBackendMapOutput
	ToSecretBackendMapOutputWithContext(context.Context) SecretBackendMapOutput
}

type SecretBackendMap map[string]SecretBackendInput

func (SecretBackendMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SecretBackend)(nil)).Elem()
}

func (i SecretBackendMap) ToSecretBackendMapOutput() SecretBackendMapOutput {
	return i.ToSecretBackendMapOutputWithContext(context.Background())
}

func (i SecretBackendMap) ToSecretBackendMapOutputWithContext(ctx context.Context) SecretBackendMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecretBackendMapOutput)
}

type SecretBackendOutput struct{ *pulumi.OutputState }

func (SecretBackendOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SecretBackend)(nil)).Elem()
}

func (o SecretBackendOutput) ToSecretBackendOutput() SecretBackendOutput {
	return o
}

func (o SecretBackendOutput) ToSecretBackendOutputWithContext(ctx context.Context) SecretBackendOutput {
	return o
}

type SecretBackendArrayOutput struct{ *pulumi.OutputState }

func (SecretBackendArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SecretBackend)(nil)).Elem()
}

func (o SecretBackendArrayOutput) ToSecretBackendArrayOutput() SecretBackendArrayOutput {
	return o
}

func (o SecretBackendArrayOutput) ToSecretBackendArrayOutputWithContext(ctx context.Context) SecretBackendArrayOutput {
	return o
}

func (o SecretBackendArrayOutput) Index(i pulumi.IntInput) SecretBackendOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SecretBackend {
		return vs[0].([]*SecretBackend)[vs[1].(int)]
	}).(SecretBackendOutput)
}

type SecretBackendMapOutput struct{ *pulumi.OutputState }

func (SecretBackendMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SecretBackend)(nil)).Elem()
}

func (o SecretBackendMapOutput) ToSecretBackendMapOutput() SecretBackendMapOutput {
	return o
}

func (o SecretBackendMapOutput) ToSecretBackendMapOutputWithContext(ctx context.Context) SecretBackendMapOutput {
	return o
}

func (o SecretBackendMapOutput) MapIndex(k pulumi.StringInput) SecretBackendOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SecretBackend {
		return vs[0].(map[string]*SecretBackend)[vs[1].(string)]
	}).(SecretBackendOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SecretBackendInput)(nil)).Elem(), &SecretBackend{})
	pulumi.RegisterInputType(reflect.TypeOf((*SecretBackendArrayInput)(nil)).Elem(), SecretBackendArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SecretBackendMapInput)(nil)).Elem(), SecretBackendMap{})
	pulumi.RegisterOutputType(SecretBackendOutput{})
	pulumi.RegisterOutputType(SecretBackendArrayOutput{})
	pulumi.RegisterOutputType(SecretBackendMapOutput{})
}
