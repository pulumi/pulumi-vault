// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package rabbitmq

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type SecretBackendRoleVhost struct {
	Configure string `pulumi:"configure"`
	Host      string `pulumi:"host"`
	Read      string `pulumi:"read"`
	Write     string `pulumi:"write"`
}

// SecretBackendRoleVhostInput is an input type that accepts SecretBackendRoleVhostArgs and SecretBackendRoleVhostOutput values.
// You can construct a concrete instance of `SecretBackendRoleVhostInput` via:
//
//          SecretBackendRoleVhostArgs{...}
type SecretBackendRoleVhostInput interface {
	pulumi.Input

	ToSecretBackendRoleVhostOutput() SecretBackendRoleVhostOutput
	ToSecretBackendRoleVhostOutputWithContext(context.Context) SecretBackendRoleVhostOutput
}

type SecretBackendRoleVhostArgs struct {
	Configure pulumi.StringInput `pulumi:"configure"`
	Host      pulumi.StringInput `pulumi:"host"`
	Read      pulumi.StringInput `pulumi:"read"`
	Write     pulumi.StringInput `pulumi:"write"`
}

func (SecretBackendRoleVhostArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SecretBackendRoleVhost)(nil)).Elem()
}

func (i SecretBackendRoleVhostArgs) ToSecretBackendRoleVhostOutput() SecretBackendRoleVhostOutput {
	return i.ToSecretBackendRoleVhostOutputWithContext(context.Background())
}

func (i SecretBackendRoleVhostArgs) ToSecretBackendRoleVhostOutputWithContext(ctx context.Context) SecretBackendRoleVhostOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecretBackendRoleVhostOutput)
}

// SecretBackendRoleVhostArrayInput is an input type that accepts SecretBackendRoleVhostArray and SecretBackendRoleVhostArrayOutput values.
// You can construct a concrete instance of `SecretBackendRoleVhostArrayInput` via:
//
//          SecretBackendRoleVhostArray{ SecretBackendRoleVhostArgs{...} }
type SecretBackendRoleVhostArrayInput interface {
	pulumi.Input

	ToSecretBackendRoleVhostArrayOutput() SecretBackendRoleVhostArrayOutput
	ToSecretBackendRoleVhostArrayOutputWithContext(context.Context) SecretBackendRoleVhostArrayOutput
}

type SecretBackendRoleVhostArray []SecretBackendRoleVhostInput

func (SecretBackendRoleVhostArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SecretBackendRoleVhost)(nil)).Elem()
}

func (i SecretBackendRoleVhostArray) ToSecretBackendRoleVhostArrayOutput() SecretBackendRoleVhostArrayOutput {
	return i.ToSecretBackendRoleVhostArrayOutputWithContext(context.Background())
}

func (i SecretBackendRoleVhostArray) ToSecretBackendRoleVhostArrayOutputWithContext(ctx context.Context) SecretBackendRoleVhostArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecretBackendRoleVhostArrayOutput)
}

type SecretBackendRoleVhostOutput struct{ *pulumi.OutputState }

func (SecretBackendRoleVhostOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SecretBackendRoleVhost)(nil)).Elem()
}

func (o SecretBackendRoleVhostOutput) ToSecretBackendRoleVhostOutput() SecretBackendRoleVhostOutput {
	return o
}

func (o SecretBackendRoleVhostOutput) ToSecretBackendRoleVhostOutputWithContext(ctx context.Context) SecretBackendRoleVhostOutput {
	return o
}

func (o SecretBackendRoleVhostOutput) Configure() pulumi.StringOutput {
	return o.ApplyT(func(v SecretBackendRoleVhost) string { return v.Configure }).(pulumi.StringOutput)
}

func (o SecretBackendRoleVhostOutput) Host() pulumi.StringOutput {
	return o.ApplyT(func(v SecretBackendRoleVhost) string { return v.Host }).(pulumi.StringOutput)
}

func (o SecretBackendRoleVhostOutput) Read() pulumi.StringOutput {
	return o.ApplyT(func(v SecretBackendRoleVhost) string { return v.Read }).(pulumi.StringOutput)
}

func (o SecretBackendRoleVhostOutput) Write() pulumi.StringOutput {
	return o.ApplyT(func(v SecretBackendRoleVhost) string { return v.Write }).(pulumi.StringOutput)
}

type SecretBackendRoleVhostArrayOutput struct{ *pulumi.OutputState }

func (SecretBackendRoleVhostArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SecretBackendRoleVhost)(nil)).Elem()
}

func (o SecretBackendRoleVhostArrayOutput) ToSecretBackendRoleVhostArrayOutput() SecretBackendRoleVhostArrayOutput {
	return o
}

func (o SecretBackendRoleVhostArrayOutput) ToSecretBackendRoleVhostArrayOutputWithContext(ctx context.Context) SecretBackendRoleVhostArrayOutput {
	return o
}

func (o SecretBackendRoleVhostArrayOutput) Index(i pulumi.IntInput) SecretBackendRoleVhostOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SecretBackendRoleVhost {
		return vs[0].([]SecretBackendRoleVhost)[vs[1].(int)]
	}).(SecretBackendRoleVhostOutput)
}

type SecretBackendRoleVhostTopic struct {
	Host string `pulumi:"host"`
	// Specifies a map of virtual hosts to permissions.
	Vhosts []SecretBackendRoleVhostTopicVhost `pulumi:"vhosts"`
}

// SecretBackendRoleVhostTopicInput is an input type that accepts SecretBackendRoleVhostTopicArgs and SecretBackendRoleVhostTopicOutput values.
// You can construct a concrete instance of `SecretBackendRoleVhostTopicInput` via:
//
//          SecretBackendRoleVhostTopicArgs{...}
type SecretBackendRoleVhostTopicInput interface {
	pulumi.Input

	ToSecretBackendRoleVhostTopicOutput() SecretBackendRoleVhostTopicOutput
	ToSecretBackendRoleVhostTopicOutputWithContext(context.Context) SecretBackendRoleVhostTopicOutput
}

type SecretBackendRoleVhostTopicArgs struct {
	Host pulumi.StringInput `pulumi:"host"`
	// Specifies a map of virtual hosts to permissions.
	Vhosts SecretBackendRoleVhostTopicVhostArrayInput `pulumi:"vhosts"`
}

func (SecretBackendRoleVhostTopicArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SecretBackendRoleVhostTopic)(nil)).Elem()
}

func (i SecretBackendRoleVhostTopicArgs) ToSecretBackendRoleVhostTopicOutput() SecretBackendRoleVhostTopicOutput {
	return i.ToSecretBackendRoleVhostTopicOutputWithContext(context.Background())
}

func (i SecretBackendRoleVhostTopicArgs) ToSecretBackendRoleVhostTopicOutputWithContext(ctx context.Context) SecretBackendRoleVhostTopicOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecretBackendRoleVhostTopicOutput)
}

// SecretBackendRoleVhostTopicArrayInput is an input type that accepts SecretBackendRoleVhostTopicArray and SecretBackendRoleVhostTopicArrayOutput values.
// You can construct a concrete instance of `SecretBackendRoleVhostTopicArrayInput` via:
//
//          SecretBackendRoleVhostTopicArray{ SecretBackendRoleVhostTopicArgs{...} }
type SecretBackendRoleVhostTopicArrayInput interface {
	pulumi.Input

	ToSecretBackendRoleVhostTopicArrayOutput() SecretBackendRoleVhostTopicArrayOutput
	ToSecretBackendRoleVhostTopicArrayOutputWithContext(context.Context) SecretBackendRoleVhostTopicArrayOutput
}

type SecretBackendRoleVhostTopicArray []SecretBackendRoleVhostTopicInput

func (SecretBackendRoleVhostTopicArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SecretBackendRoleVhostTopic)(nil)).Elem()
}

func (i SecretBackendRoleVhostTopicArray) ToSecretBackendRoleVhostTopicArrayOutput() SecretBackendRoleVhostTopicArrayOutput {
	return i.ToSecretBackendRoleVhostTopicArrayOutputWithContext(context.Background())
}

func (i SecretBackendRoleVhostTopicArray) ToSecretBackendRoleVhostTopicArrayOutputWithContext(ctx context.Context) SecretBackendRoleVhostTopicArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecretBackendRoleVhostTopicArrayOutput)
}

type SecretBackendRoleVhostTopicOutput struct{ *pulumi.OutputState }

func (SecretBackendRoleVhostTopicOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SecretBackendRoleVhostTopic)(nil)).Elem()
}

func (o SecretBackendRoleVhostTopicOutput) ToSecretBackendRoleVhostTopicOutput() SecretBackendRoleVhostTopicOutput {
	return o
}

func (o SecretBackendRoleVhostTopicOutput) ToSecretBackendRoleVhostTopicOutputWithContext(ctx context.Context) SecretBackendRoleVhostTopicOutput {
	return o
}

func (o SecretBackendRoleVhostTopicOutput) Host() pulumi.StringOutput {
	return o.ApplyT(func(v SecretBackendRoleVhostTopic) string { return v.Host }).(pulumi.StringOutput)
}

// Specifies a map of virtual hosts to permissions.
func (o SecretBackendRoleVhostTopicOutput) Vhosts() SecretBackendRoleVhostTopicVhostArrayOutput {
	return o.ApplyT(func(v SecretBackendRoleVhostTopic) []SecretBackendRoleVhostTopicVhost { return v.Vhosts }).(SecretBackendRoleVhostTopicVhostArrayOutput)
}

type SecretBackendRoleVhostTopicArrayOutput struct{ *pulumi.OutputState }

func (SecretBackendRoleVhostTopicArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SecretBackendRoleVhostTopic)(nil)).Elem()
}

func (o SecretBackendRoleVhostTopicArrayOutput) ToSecretBackendRoleVhostTopicArrayOutput() SecretBackendRoleVhostTopicArrayOutput {
	return o
}

func (o SecretBackendRoleVhostTopicArrayOutput) ToSecretBackendRoleVhostTopicArrayOutputWithContext(ctx context.Context) SecretBackendRoleVhostTopicArrayOutput {
	return o
}

func (o SecretBackendRoleVhostTopicArrayOutput) Index(i pulumi.IntInput) SecretBackendRoleVhostTopicOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SecretBackendRoleVhostTopic {
		return vs[0].([]SecretBackendRoleVhostTopic)[vs[1].(int)]
	}).(SecretBackendRoleVhostTopicOutput)
}

type SecretBackendRoleVhostTopicVhost struct {
	Read  string `pulumi:"read"`
	Topic string `pulumi:"topic"`
	Write string `pulumi:"write"`
}

// SecretBackendRoleVhostTopicVhostInput is an input type that accepts SecretBackendRoleVhostTopicVhostArgs and SecretBackendRoleVhostTopicVhostOutput values.
// You can construct a concrete instance of `SecretBackendRoleVhostTopicVhostInput` via:
//
//          SecretBackendRoleVhostTopicVhostArgs{...}
type SecretBackendRoleVhostTopicVhostInput interface {
	pulumi.Input

	ToSecretBackendRoleVhostTopicVhostOutput() SecretBackendRoleVhostTopicVhostOutput
	ToSecretBackendRoleVhostTopicVhostOutputWithContext(context.Context) SecretBackendRoleVhostTopicVhostOutput
}

type SecretBackendRoleVhostTopicVhostArgs struct {
	Read  pulumi.StringInput `pulumi:"read"`
	Topic pulumi.StringInput `pulumi:"topic"`
	Write pulumi.StringInput `pulumi:"write"`
}

func (SecretBackendRoleVhostTopicVhostArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SecretBackendRoleVhostTopicVhost)(nil)).Elem()
}

func (i SecretBackendRoleVhostTopicVhostArgs) ToSecretBackendRoleVhostTopicVhostOutput() SecretBackendRoleVhostTopicVhostOutput {
	return i.ToSecretBackendRoleVhostTopicVhostOutputWithContext(context.Background())
}

func (i SecretBackendRoleVhostTopicVhostArgs) ToSecretBackendRoleVhostTopicVhostOutputWithContext(ctx context.Context) SecretBackendRoleVhostTopicVhostOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecretBackendRoleVhostTopicVhostOutput)
}

// SecretBackendRoleVhostTopicVhostArrayInput is an input type that accepts SecretBackendRoleVhostTopicVhostArray and SecretBackendRoleVhostTopicVhostArrayOutput values.
// You can construct a concrete instance of `SecretBackendRoleVhostTopicVhostArrayInput` via:
//
//          SecretBackendRoleVhostTopicVhostArray{ SecretBackendRoleVhostTopicVhostArgs{...} }
type SecretBackendRoleVhostTopicVhostArrayInput interface {
	pulumi.Input

	ToSecretBackendRoleVhostTopicVhostArrayOutput() SecretBackendRoleVhostTopicVhostArrayOutput
	ToSecretBackendRoleVhostTopicVhostArrayOutputWithContext(context.Context) SecretBackendRoleVhostTopicVhostArrayOutput
}

type SecretBackendRoleVhostTopicVhostArray []SecretBackendRoleVhostTopicVhostInput

func (SecretBackendRoleVhostTopicVhostArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SecretBackendRoleVhostTopicVhost)(nil)).Elem()
}

func (i SecretBackendRoleVhostTopicVhostArray) ToSecretBackendRoleVhostTopicVhostArrayOutput() SecretBackendRoleVhostTopicVhostArrayOutput {
	return i.ToSecretBackendRoleVhostTopicVhostArrayOutputWithContext(context.Background())
}

func (i SecretBackendRoleVhostTopicVhostArray) ToSecretBackendRoleVhostTopicVhostArrayOutputWithContext(ctx context.Context) SecretBackendRoleVhostTopicVhostArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecretBackendRoleVhostTopicVhostArrayOutput)
}

type SecretBackendRoleVhostTopicVhostOutput struct{ *pulumi.OutputState }

func (SecretBackendRoleVhostTopicVhostOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SecretBackendRoleVhostTopicVhost)(nil)).Elem()
}

func (o SecretBackendRoleVhostTopicVhostOutput) ToSecretBackendRoleVhostTopicVhostOutput() SecretBackendRoleVhostTopicVhostOutput {
	return o
}

func (o SecretBackendRoleVhostTopicVhostOutput) ToSecretBackendRoleVhostTopicVhostOutputWithContext(ctx context.Context) SecretBackendRoleVhostTopicVhostOutput {
	return o
}

func (o SecretBackendRoleVhostTopicVhostOutput) Read() pulumi.StringOutput {
	return o.ApplyT(func(v SecretBackendRoleVhostTopicVhost) string { return v.Read }).(pulumi.StringOutput)
}

func (o SecretBackendRoleVhostTopicVhostOutput) Topic() pulumi.StringOutput {
	return o.ApplyT(func(v SecretBackendRoleVhostTopicVhost) string { return v.Topic }).(pulumi.StringOutput)
}

func (o SecretBackendRoleVhostTopicVhostOutput) Write() pulumi.StringOutput {
	return o.ApplyT(func(v SecretBackendRoleVhostTopicVhost) string { return v.Write }).(pulumi.StringOutput)
}

type SecretBackendRoleVhostTopicVhostArrayOutput struct{ *pulumi.OutputState }

func (SecretBackendRoleVhostTopicVhostArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SecretBackendRoleVhostTopicVhost)(nil)).Elem()
}

func (o SecretBackendRoleVhostTopicVhostArrayOutput) ToSecretBackendRoleVhostTopicVhostArrayOutput() SecretBackendRoleVhostTopicVhostArrayOutput {
	return o
}

func (o SecretBackendRoleVhostTopicVhostArrayOutput) ToSecretBackendRoleVhostTopicVhostArrayOutputWithContext(ctx context.Context) SecretBackendRoleVhostTopicVhostArrayOutput {
	return o
}

func (o SecretBackendRoleVhostTopicVhostArrayOutput) Index(i pulumi.IntInput) SecretBackendRoleVhostTopicVhostOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SecretBackendRoleVhostTopicVhost {
		return vs[0].([]SecretBackendRoleVhostTopicVhost)[vs[1].(int)]
	}).(SecretBackendRoleVhostTopicVhostOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SecretBackendRoleVhostInput)(nil)).Elem(), SecretBackendRoleVhostArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SecretBackendRoleVhostArrayInput)(nil)).Elem(), SecretBackendRoleVhostArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SecretBackendRoleVhostTopicInput)(nil)).Elem(), SecretBackendRoleVhostTopicArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SecretBackendRoleVhostTopicArrayInput)(nil)).Elem(), SecretBackendRoleVhostTopicArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SecretBackendRoleVhostTopicVhostInput)(nil)).Elem(), SecretBackendRoleVhostTopicVhostArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SecretBackendRoleVhostTopicVhostArrayInput)(nil)).Elem(), SecretBackendRoleVhostTopicVhostArray{})
	pulumi.RegisterOutputType(SecretBackendRoleVhostOutput{})
	pulumi.RegisterOutputType(SecretBackendRoleVhostArrayOutput{})
	pulumi.RegisterOutputType(SecretBackendRoleVhostTopicOutput{})
	pulumi.RegisterOutputType(SecretBackendRoleVhostTopicArrayOutput{})
	pulumi.RegisterOutputType(SecretBackendRoleVhostTopicVhostOutput{})
	pulumi.RegisterOutputType(SecretBackendRoleVhostTopicVhostArrayOutput{})
}
