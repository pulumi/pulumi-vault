// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package kmip

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages KMIP Secret backends in a Vault server. This feature requires
// Vault Enterprise. See the [Vault documentation](https://www.vaultproject.io/docs/secrets/kmip)
// for more information.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-vault/sdk/v5/go/vault/kmip"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := kmip.NewSecretBackend(ctx, "default", &kmip.SecretBackendArgs{
// 			DefaultTlsClientKeyBits: pulumi.Int(4096),
// 			DefaultTlsClientKeyType: pulumi.String("rsa"),
// 			DefaultTlsClientTtl:     pulumi.Int(86400),
// 			Description:             pulumi.String("Vault KMIP backend"),
// 			ListenAddrs: pulumi.StringArray{
// 				pulumi.String("127.0.0.1:5696"),
// 				pulumi.String("127.0.0.1:8080"),
// 			},
// 			Path:         pulumi.String("kmip"),
// 			TlsCaKeyBits: pulumi.Int(4096),
// 			TlsCaKeyType: pulumi.String("rsa"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## Import
//
// KMIP Secret backend can be imported using the `path`, e.g.
//
// ```sh
//  $ pulumi import vault:kmip/secretBackend:SecretBackend default kmip
// ```
type SecretBackend struct {
	pulumi.CustomResourceState

	// Client certificate key bits, valid values depend on key type.
	DefaultTlsClientKeyBits pulumi.IntOutput `pulumi:"defaultTlsClientKeyBits"`
	// Client certificate key type, `rsa` or `ec`.
	DefaultTlsClientKeyType pulumi.StringOutput `pulumi:"defaultTlsClientKeyType"`
	// Client certificate TTL in seconds
	DefaultTlsClientTtl pulumi.IntOutput `pulumi:"defaultTlsClientTtl"`
	// A human-friendly description for this backend.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// If set, opts out of mount migration on path updates.
	// See here for more info on [Mount Migration](https://www.vaultproject.io/docs/concepts/mount-migration)
	DisableRemount pulumi.BoolPtrOutput `pulumi:"disableRemount"`
	// Addresses the KMIP server should listen on (`host:port`).
	ListenAddrs pulumi.StringArrayOutput `pulumi:"listenAddrs"`
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
	// *Available only for Vault Enterprise*.
	Namespace pulumi.StringPtrOutput `pulumi:"namespace"`
	// The unique path this backend should be mounted at. Must
	// not begin or end with a `/`. Defaults to `kmip`.
	Path pulumi.StringOutput `pulumi:"path"`
	// Hostnames to include in the server's TLS certificate as SAN DNS names. The first will be used as the common name (CN).
	ServerHostnames pulumi.StringArrayOutput `pulumi:"serverHostnames"`
	// IPs to include in the server's TLS certificate as SAN IP addresses.
	ServerIps pulumi.StringArrayOutput `pulumi:"serverIps"`
	// CA key bits, valid values depend on key type.
	TlsCaKeyBits pulumi.IntOutput `pulumi:"tlsCaKeyBits"`
	// CA key type, rsa or ec.
	TlsCaKeyType pulumi.StringOutput `pulumi:"tlsCaKeyType"`
	// Minimum TLS version to accept.
	TlsMinVersion pulumi.StringOutput `pulumi:"tlsMinVersion"`
}

// NewSecretBackend registers a new resource with the given unique name, arguments, and options.
func NewSecretBackend(ctx *pulumi.Context,
	name string, args *SecretBackendArgs, opts ...pulumi.ResourceOption) (*SecretBackend, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Path == nil {
		return nil, errors.New("invalid value for required argument 'Path'")
	}
	var resource SecretBackend
	err := ctx.RegisterResource("vault:kmip/secretBackend:SecretBackend", name, args, &resource, opts...)
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
	err := ctx.ReadResource("vault:kmip/secretBackend:SecretBackend", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SecretBackend resources.
type secretBackendState struct {
	// Client certificate key bits, valid values depend on key type.
	DefaultTlsClientKeyBits *int `pulumi:"defaultTlsClientKeyBits"`
	// Client certificate key type, `rsa` or `ec`.
	DefaultTlsClientKeyType *string `pulumi:"defaultTlsClientKeyType"`
	// Client certificate TTL in seconds
	DefaultTlsClientTtl *int `pulumi:"defaultTlsClientTtl"`
	// A human-friendly description for this backend.
	Description *string `pulumi:"description"`
	// If set, opts out of mount migration on path updates.
	// See here for more info on [Mount Migration](https://www.vaultproject.io/docs/concepts/mount-migration)
	DisableRemount *bool `pulumi:"disableRemount"`
	// Addresses the KMIP server should listen on (`host:port`).
	ListenAddrs []string `pulumi:"listenAddrs"`
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
	// *Available only for Vault Enterprise*.
	Namespace *string `pulumi:"namespace"`
	// The unique path this backend should be mounted at. Must
	// not begin or end with a `/`. Defaults to `kmip`.
	Path *string `pulumi:"path"`
	// Hostnames to include in the server's TLS certificate as SAN DNS names. The first will be used as the common name (CN).
	ServerHostnames []string `pulumi:"serverHostnames"`
	// IPs to include in the server's TLS certificate as SAN IP addresses.
	ServerIps []string `pulumi:"serverIps"`
	// CA key bits, valid values depend on key type.
	TlsCaKeyBits *int `pulumi:"tlsCaKeyBits"`
	// CA key type, rsa or ec.
	TlsCaKeyType *string `pulumi:"tlsCaKeyType"`
	// Minimum TLS version to accept.
	TlsMinVersion *string `pulumi:"tlsMinVersion"`
}

type SecretBackendState struct {
	// Client certificate key bits, valid values depend on key type.
	DefaultTlsClientKeyBits pulumi.IntPtrInput
	// Client certificate key type, `rsa` or `ec`.
	DefaultTlsClientKeyType pulumi.StringPtrInput
	// Client certificate TTL in seconds
	DefaultTlsClientTtl pulumi.IntPtrInput
	// A human-friendly description for this backend.
	Description pulumi.StringPtrInput
	// If set, opts out of mount migration on path updates.
	// See here for more info on [Mount Migration](https://www.vaultproject.io/docs/concepts/mount-migration)
	DisableRemount pulumi.BoolPtrInput
	// Addresses the KMIP server should listen on (`host:port`).
	ListenAddrs pulumi.StringArrayInput
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
	// *Available only for Vault Enterprise*.
	Namespace pulumi.StringPtrInput
	// The unique path this backend should be mounted at. Must
	// not begin or end with a `/`. Defaults to `kmip`.
	Path pulumi.StringPtrInput
	// Hostnames to include in the server's TLS certificate as SAN DNS names. The first will be used as the common name (CN).
	ServerHostnames pulumi.StringArrayInput
	// IPs to include in the server's TLS certificate as SAN IP addresses.
	ServerIps pulumi.StringArrayInput
	// CA key bits, valid values depend on key type.
	TlsCaKeyBits pulumi.IntPtrInput
	// CA key type, rsa or ec.
	TlsCaKeyType pulumi.StringPtrInput
	// Minimum TLS version to accept.
	TlsMinVersion pulumi.StringPtrInput
}

func (SecretBackendState) ElementType() reflect.Type {
	return reflect.TypeOf((*secretBackendState)(nil)).Elem()
}

type secretBackendArgs struct {
	// Client certificate key bits, valid values depend on key type.
	DefaultTlsClientKeyBits *int `pulumi:"defaultTlsClientKeyBits"`
	// Client certificate key type, `rsa` or `ec`.
	DefaultTlsClientKeyType *string `pulumi:"defaultTlsClientKeyType"`
	// Client certificate TTL in seconds
	DefaultTlsClientTtl *int `pulumi:"defaultTlsClientTtl"`
	// A human-friendly description for this backend.
	Description *string `pulumi:"description"`
	// If set, opts out of mount migration on path updates.
	// See here for more info on [Mount Migration](https://www.vaultproject.io/docs/concepts/mount-migration)
	DisableRemount *bool `pulumi:"disableRemount"`
	// Addresses the KMIP server should listen on (`host:port`).
	ListenAddrs []string `pulumi:"listenAddrs"`
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
	// *Available only for Vault Enterprise*.
	Namespace *string `pulumi:"namespace"`
	// The unique path this backend should be mounted at. Must
	// not begin or end with a `/`. Defaults to `kmip`.
	Path string `pulumi:"path"`
	// Hostnames to include in the server's TLS certificate as SAN DNS names. The first will be used as the common name (CN).
	ServerHostnames []string `pulumi:"serverHostnames"`
	// IPs to include in the server's TLS certificate as SAN IP addresses.
	ServerIps []string `pulumi:"serverIps"`
	// CA key bits, valid values depend on key type.
	TlsCaKeyBits *int `pulumi:"tlsCaKeyBits"`
	// CA key type, rsa or ec.
	TlsCaKeyType *string `pulumi:"tlsCaKeyType"`
	// Minimum TLS version to accept.
	TlsMinVersion *string `pulumi:"tlsMinVersion"`
}

// The set of arguments for constructing a SecretBackend resource.
type SecretBackendArgs struct {
	// Client certificate key bits, valid values depend on key type.
	DefaultTlsClientKeyBits pulumi.IntPtrInput
	// Client certificate key type, `rsa` or `ec`.
	DefaultTlsClientKeyType pulumi.StringPtrInput
	// Client certificate TTL in seconds
	DefaultTlsClientTtl pulumi.IntPtrInput
	// A human-friendly description for this backend.
	Description pulumi.StringPtrInput
	// If set, opts out of mount migration on path updates.
	// See here for more info on [Mount Migration](https://www.vaultproject.io/docs/concepts/mount-migration)
	DisableRemount pulumi.BoolPtrInput
	// Addresses the KMIP server should listen on (`host:port`).
	ListenAddrs pulumi.StringArrayInput
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
	// *Available only for Vault Enterprise*.
	Namespace pulumi.StringPtrInput
	// The unique path this backend should be mounted at. Must
	// not begin or end with a `/`. Defaults to `kmip`.
	Path pulumi.StringInput
	// Hostnames to include in the server's TLS certificate as SAN DNS names. The first will be used as the common name (CN).
	ServerHostnames pulumi.StringArrayInput
	// IPs to include in the server's TLS certificate as SAN IP addresses.
	ServerIps pulumi.StringArrayInput
	// CA key bits, valid values depend on key type.
	TlsCaKeyBits pulumi.IntPtrInput
	// CA key type, rsa or ec.
	TlsCaKeyType pulumi.StringPtrInput
	// Minimum TLS version to accept.
	TlsMinVersion pulumi.StringPtrInput
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

// Client certificate key bits, valid values depend on key type.
func (o SecretBackendOutput) DefaultTlsClientKeyBits() pulumi.IntOutput {
	return o.ApplyT(func(v *SecretBackend) pulumi.IntOutput { return v.DefaultTlsClientKeyBits }).(pulumi.IntOutput)
}

// Client certificate key type, `rsa` or `ec`.
func (o SecretBackendOutput) DefaultTlsClientKeyType() pulumi.StringOutput {
	return o.ApplyT(func(v *SecretBackend) pulumi.StringOutput { return v.DefaultTlsClientKeyType }).(pulumi.StringOutput)
}

// Client certificate TTL in seconds
func (o SecretBackendOutput) DefaultTlsClientTtl() pulumi.IntOutput {
	return o.ApplyT(func(v *SecretBackend) pulumi.IntOutput { return v.DefaultTlsClientTtl }).(pulumi.IntOutput)
}

// A human-friendly description for this backend.
func (o SecretBackendOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecretBackend) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// If set, opts out of mount migration on path updates.
// See here for more info on [Mount Migration](https://www.vaultproject.io/docs/concepts/mount-migration)
func (o SecretBackendOutput) DisableRemount() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SecretBackend) pulumi.BoolPtrOutput { return v.DisableRemount }).(pulumi.BoolPtrOutput)
}

// Addresses the KMIP server should listen on (`host:port`).
func (o SecretBackendOutput) ListenAddrs() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *SecretBackend) pulumi.StringArrayOutput { return v.ListenAddrs }).(pulumi.StringArrayOutput)
}

// The namespace to provision the resource in.
// The value should not contain leading or trailing forward slashes.
// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
// *Available only for Vault Enterprise*.
func (o SecretBackendOutput) Namespace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecretBackend) pulumi.StringPtrOutput { return v.Namespace }).(pulumi.StringPtrOutput)
}

// The unique path this backend should be mounted at. Must
// not begin or end with a `/`. Defaults to `kmip`.
func (o SecretBackendOutput) Path() pulumi.StringOutput {
	return o.ApplyT(func(v *SecretBackend) pulumi.StringOutput { return v.Path }).(pulumi.StringOutput)
}

// Hostnames to include in the server's TLS certificate as SAN DNS names. The first will be used as the common name (CN).
func (o SecretBackendOutput) ServerHostnames() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *SecretBackend) pulumi.StringArrayOutput { return v.ServerHostnames }).(pulumi.StringArrayOutput)
}

// IPs to include in the server's TLS certificate as SAN IP addresses.
func (o SecretBackendOutput) ServerIps() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *SecretBackend) pulumi.StringArrayOutput { return v.ServerIps }).(pulumi.StringArrayOutput)
}

// CA key bits, valid values depend on key type.
func (o SecretBackendOutput) TlsCaKeyBits() pulumi.IntOutput {
	return o.ApplyT(func(v *SecretBackend) pulumi.IntOutput { return v.TlsCaKeyBits }).(pulumi.IntOutput)
}

// CA key type, rsa or ec.
func (o SecretBackendOutput) TlsCaKeyType() pulumi.StringOutput {
	return o.ApplyT(func(v *SecretBackend) pulumi.StringOutput { return v.TlsCaKeyType }).(pulumi.StringOutput)
}

// Minimum TLS version to accept.
func (o SecretBackendOutput) TlsMinVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *SecretBackend) pulumi.StringOutput { return v.TlsMinVersion }).(pulumi.StringOutput)
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
