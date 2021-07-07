// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package rabbitmq

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## Import
//
// RabbitMQ secret backends can be imported using the `path`, e.g.
//
// ```sh
//  $ pulumi import vault:rabbitMq/secretBackend:SecretBackend rabbitmq rabbitmq
// ```
type SecretBackend struct {
	pulumi.CustomResourceState

	// Specifies the RabbitMQ connection URI.
	ConnectionUri pulumi.StringOutput `pulumi:"connectionUri"`
	// The default TTL for credentials
	// issued by this backend.
	DefaultLeaseTtlSeconds pulumi.IntOutput `pulumi:"defaultLeaseTtlSeconds"`
	// A human-friendly description for this backend.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The maximum TTL that can be requested
	// for credentials issued by this backend.
	MaxLeaseTtlSeconds pulumi.IntOutput `pulumi:"maxLeaseTtlSeconds"`
	// Specifies the RabbitMQ management administrator password.
	Password pulumi.StringOutput `pulumi:"password"`
	// The unique path this backend should be mounted at. Must
	// not begin or end with a `/`. Defaults to `rabbitmq`.
	Path pulumi.StringPtrOutput `pulumi:"path"`
	// Specifies the RabbitMQ management administrator username.
	Username pulumi.StringOutput `pulumi:"username"`
	// Specifies whether to verify connection URI, username, and password.
	// Defaults to `true`.
	VerifyConnection pulumi.BoolPtrOutput `pulumi:"verifyConnection"`
}

// NewSecretBackend registers a new resource with the given unique name, arguments, and options.
func NewSecretBackend(ctx *pulumi.Context,
	name string, args *SecretBackendArgs, opts ...pulumi.ResourceOption) (*SecretBackend, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ConnectionUri == nil {
		return nil, errors.New("invalid value for required argument 'ConnectionUri'")
	}
	if args.Password == nil {
		return nil, errors.New("invalid value for required argument 'Password'")
	}
	if args.Username == nil {
		return nil, errors.New("invalid value for required argument 'Username'")
	}
	var resource SecretBackend
	err := ctx.RegisterResource("vault:rabbitMq/secretBackend:SecretBackend", name, args, &resource, opts...)
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
	err := ctx.ReadResource("vault:rabbitMq/secretBackend:SecretBackend", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SecretBackend resources.
type secretBackendState struct {
	// Specifies the RabbitMQ connection URI.
	ConnectionUri *string `pulumi:"connectionUri"`
	// The default TTL for credentials
	// issued by this backend.
	DefaultLeaseTtlSeconds *int `pulumi:"defaultLeaseTtlSeconds"`
	// A human-friendly description for this backend.
	Description *string `pulumi:"description"`
	// The maximum TTL that can be requested
	// for credentials issued by this backend.
	MaxLeaseTtlSeconds *int `pulumi:"maxLeaseTtlSeconds"`
	// Specifies the RabbitMQ management administrator password.
	Password *string `pulumi:"password"`
	// The unique path this backend should be mounted at. Must
	// not begin or end with a `/`. Defaults to `rabbitmq`.
	Path *string `pulumi:"path"`
	// Specifies the RabbitMQ management administrator username.
	Username *string `pulumi:"username"`
	// Specifies whether to verify connection URI, username, and password.
	// Defaults to `true`.
	VerifyConnection *bool `pulumi:"verifyConnection"`
}

type SecretBackendState struct {
	// Specifies the RabbitMQ connection URI.
	ConnectionUri pulumi.StringPtrInput
	// The default TTL for credentials
	// issued by this backend.
	DefaultLeaseTtlSeconds pulumi.IntPtrInput
	// A human-friendly description for this backend.
	Description pulumi.StringPtrInput
	// The maximum TTL that can be requested
	// for credentials issued by this backend.
	MaxLeaseTtlSeconds pulumi.IntPtrInput
	// Specifies the RabbitMQ management administrator password.
	Password pulumi.StringPtrInput
	// The unique path this backend should be mounted at. Must
	// not begin or end with a `/`. Defaults to `rabbitmq`.
	Path pulumi.StringPtrInput
	// Specifies the RabbitMQ management administrator username.
	Username pulumi.StringPtrInput
	// Specifies whether to verify connection URI, username, and password.
	// Defaults to `true`.
	VerifyConnection pulumi.BoolPtrInput
}

func (SecretBackendState) ElementType() reflect.Type {
	return reflect.TypeOf((*secretBackendState)(nil)).Elem()
}

type secretBackendArgs struct {
	// Specifies the RabbitMQ connection URI.
	ConnectionUri string `pulumi:"connectionUri"`
	// The default TTL for credentials
	// issued by this backend.
	DefaultLeaseTtlSeconds *int `pulumi:"defaultLeaseTtlSeconds"`
	// A human-friendly description for this backend.
	Description *string `pulumi:"description"`
	// The maximum TTL that can be requested
	// for credentials issued by this backend.
	MaxLeaseTtlSeconds *int `pulumi:"maxLeaseTtlSeconds"`
	// Specifies the RabbitMQ management administrator password.
	Password string `pulumi:"password"`
	// The unique path this backend should be mounted at. Must
	// not begin or end with a `/`. Defaults to `rabbitmq`.
	Path *string `pulumi:"path"`
	// Specifies the RabbitMQ management administrator username.
	Username string `pulumi:"username"`
	// Specifies whether to verify connection URI, username, and password.
	// Defaults to `true`.
	VerifyConnection *bool `pulumi:"verifyConnection"`
}

// The set of arguments for constructing a SecretBackend resource.
type SecretBackendArgs struct {
	// Specifies the RabbitMQ connection URI.
	ConnectionUri pulumi.StringInput
	// The default TTL for credentials
	// issued by this backend.
	DefaultLeaseTtlSeconds pulumi.IntPtrInput
	// A human-friendly description for this backend.
	Description pulumi.StringPtrInput
	// The maximum TTL that can be requested
	// for credentials issued by this backend.
	MaxLeaseTtlSeconds pulumi.IntPtrInput
	// Specifies the RabbitMQ management administrator password.
	Password pulumi.StringInput
	// The unique path this backend should be mounted at. Must
	// not begin or end with a `/`. Defaults to `rabbitmq`.
	Path pulumi.StringPtrInput
	// Specifies the RabbitMQ management administrator username.
	Username pulumi.StringInput
	// Specifies whether to verify connection URI, username, and password.
	// Defaults to `true`.
	VerifyConnection pulumi.BoolPtrInput
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
	return reflect.TypeOf((*SecretBackend)(nil))
}

func (i *SecretBackend) ToSecretBackendOutput() SecretBackendOutput {
	return i.ToSecretBackendOutputWithContext(context.Background())
}

func (i *SecretBackend) ToSecretBackendOutputWithContext(ctx context.Context) SecretBackendOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecretBackendOutput)
}

func (i *SecretBackend) ToSecretBackendPtrOutput() SecretBackendPtrOutput {
	return i.ToSecretBackendPtrOutputWithContext(context.Background())
}

func (i *SecretBackend) ToSecretBackendPtrOutputWithContext(ctx context.Context) SecretBackendPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecretBackendPtrOutput)
}

type SecretBackendPtrInput interface {
	pulumi.Input

	ToSecretBackendPtrOutput() SecretBackendPtrOutput
	ToSecretBackendPtrOutputWithContext(ctx context.Context) SecretBackendPtrOutput
}

type secretBackendPtrType SecretBackendArgs

func (*secretBackendPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SecretBackend)(nil))
}

func (i *secretBackendPtrType) ToSecretBackendPtrOutput() SecretBackendPtrOutput {
	return i.ToSecretBackendPtrOutputWithContext(context.Background())
}

func (i *secretBackendPtrType) ToSecretBackendPtrOutputWithContext(ctx context.Context) SecretBackendPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecretBackendPtrOutput)
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
	return reflect.TypeOf(([]*SecretBackend)(nil))
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
	return reflect.TypeOf((map[string]*SecretBackend)(nil))
}

func (i SecretBackendMap) ToSecretBackendMapOutput() SecretBackendMapOutput {
	return i.ToSecretBackendMapOutputWithContext(context.Background())
}

func (i SecretBackendMap) ToSecretBackendMapOutputWithContext(ctx context.Context) SecretBackendMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecretBackendMapOutput)
}

type SecretBackendOutput struct {
	*pulumi.OutputState
}

func (SecretBackendOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SecretBackend)(nil))
}

func (o SecretBackendOutput) ToSecretBackendOutput() SecretBackendOutput {
	return o
}

func (o SecretBackendOutput) ToSecretBackendOutputWithContext(ctx context.Context) SecretBackendOutput {
	return o
}

func (o SecretBackendOutput) ToSecretBackendPtrOutput() SecretBackendPtrOutput {
	return o.ToSecretBackendPtrOutputWithContext(context.Background())
}

func (o SecretBackendOutput) ToSecretBackendPtrOutputWithContext(ctx context.Context) SecretBackendPtrOutput {
	return o.ApplyT(func(v SecretBackend) *SecretBackend {
		return &v
	}).(SecretBackendPtrOutput)
}

type SecretBackendPtrOutput struct {
	*pulumi.OutputState
}

func (SecretBackendPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SecretBackend)(nil))
}

func (o SecretBackendPtrOutput) ToSecretBackendPtrOutput() SecretBackendPtrOutput {
	return o
}

func (o SecretBackendPtrOutput) ToSecretBackendPtrOutputWithContext(ctx context.Context) SecretBackendPtrOutput {
	return o
}

type SecretBackendArrayOutput struct{ *pulumi.OutputState }

func (SecretBackendArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SecretBackend)(nil))
}

func (o SecretBackendArrayOutput) ToSecretBackendArrayOutput() SecretBackendArrayOutput {
	return o
}

func (o SecretBackendArrayOutput) ToSecretBackendArrayOutputWithContext(ctx context.Context) SecretBackendArrayOutput {
	return o
}

func (o SecretBackendArrayOutput) Index(i pulumi.IntInput) SecretBackendOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SecretBackend {
		return vs[0].([]SecretBackend)[vs[1].(int)]
	}).(SecretBackendOutput)
}

type SecretBackendMapOutput struct{ *pulumi.OutputState }

func (SecretBackendMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]SecretBackend)(nil))
}

func (o SecretBackendMapOutput) ToSecretBackendMapOutput() SecretBackendMapOutput {
	return o
}

func (o SecretBackendMapOutput) ToSecretBackendMapOutputWithContext(ctx context.Context) SecretBackendMapOutput {
	return o
}

func (o SecretBackendMapOutput) MapIndex(k pulumi.StringInput) SecretBackendOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) SecretBackend {
		return vs[0].(map[string]SecretBackend)[vs[1].(string)]
	}).(SecretBackendOutput)
}

func init() {
	pulumi.RegisterOutputType(SecretBackendOutput{})
	pulumi.RegisterOutputType(SecretBackendPtrOutput{})
	pulumi.RegisterOutputType(SecretBackendArrayOutput{})
	pulumi.RegisterOutputType(SecretBackendMapOutput{})
}
