// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package config

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AuthLogins struct {
	Method     *string           `pulumi:"method"`
	Namespace  *string           `pulumi:"namespace"`
	Parameters map[string]string `pulumi:"parameters"`
	Path       string            `pulumi:"path"`
}

// AuthLoginsInput is an input type that accepts AuthLoginsArgs and AuthLoginsOutput values.
// You can construct a concrete instance of `AuthLoginsInput` via:
//
//          AuthLoginsArgs{...}
type AuthLoginsInput interface {
	pulumi.Input

	ToAuthLoginsOutput() AuthLoginsOutput
	ToAuthLoginsOutputWithContext(context.Context) AuthLoginsOutput
}

type AuthLoginsArgs struct {
	Method     pulumi.StringPtrInput `pulumi:"method"`
	Namespace  pulumi.StringPtrInput `pulumi:"namespace"`
	Parameters pulumi.StringMapInput `pulumi:"parameters"`
	Path       pulumi.StringInput    `pulumi:"path"`
}

func (AuthLoginsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AuthLogins)(nil)).Elem()
}

func (i AuthLoginsArgs) ToAuthLoginsOutput() AuthLoginsOutput {
	return i.ToAuthLoginsOutputWithContext(context.Background())
}

func (i AuthLoginsArgs) ToAuthLoginsOutputWithContext(ctx context.Context) AuthLoginsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AuthLoginsOutput)
}

// AuthLoginsArrayInput is an input type that accepts AuthLoginsArray and AuthLoginsArrayOutput values.
// You can construct a concrete instance of `AuthLoginsArrayInput` via:
//
//          AuthLoginsArray{ AuthLoginsArgs{...} }
type AuthLoginsArrayInput interface {
	pulumi.Input

	ToAuthLoginsArrayOutput() AuthLoginsArrayOutput
	ToAuthLoginsArrayOutputWithContext(context.Context) AuthLoginsArrayOutput
}

type AuthLoginsArray []AuthLoginsInput

func (AuthLoginsArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AuthLogins)(nil)).Elem()
}

func (i AuthLoginsArray) ToAuthLoginsArrayOutput() AuthLoginsArrayOutput {
	return i.ToAuthLoginsArrayOutputWithContext(context.Background())
}

func (i AuthLoginsArray) ToAuthLoginsArrayOutputWithContext(ctx context.Context) AuthLoginsArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AuthLoginsArrayOutput)
}

type AuthLoginsOutput struct{ *pulumi.OutputState }

func (AuthLoginsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AuthLogins)(nil)).Elem()
}

func (o AuthLoginsOutput) ToAuthLoginsOutput() AuthLoginsOutput {
	return o
}

func (o AuthLoginsOutput) ToAuthLoginsOutputWithContext(ctx context.Context) AuthLoginsOutput {
	return o
}

func (o AuthLoginsOutput) Method() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AuthLogins) *string { return v.Method }).(pulumi.StringPtrOutput)
}

func (o AuthLoginsOutput) Namespace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AuthLogins) *string { return v.Namespace }).(pulumi.StringPtrOutput)
}

func (o AuthLoginsOutput) Parameters() pulumi.StringMapOutput {
	return o.ApplyT(func(v AuthLogins) map[string]string { return v.Parameters }).(pulumi.StringMapOutput)
}

func (o AuthLoginsOutput) Path() pulumi.StringOutput {
	return o.ApplyT(func(v AuthLogins) string { return v.Path }).(pulumi.StringOutput)
}

type AuthLoginsArrayOutput struct{ *pulumi.OutputState }

func (AuthLoginsArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AuthLogins)(nil)).Elem()
}

func (o AuthLoginsArrayOutput) ToAuthLoginsArrayOutput() AuthLoginsArrayOutput {
	return o
}

func (o AuthLoginsArrayOutput) ToAuthLoginsArrayOutputWithContext(ctx context.Context) AuthLoginsArrayOutput {
	return o
}

func (o AuthLoginsArrayOutput) Index(i pulumi.IntInput) AuthLoginsOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) AuthLogins {
		return vs[0].([]AuthLogins)[vs[1].(int)]
	}).(AuthLoginsOutput)
}

type ClientAuths struct {
	CertFile string `pulumi:"certFile"`
	KeyFile  string `pulumi:"keyFile"`
}

// ClientAuthsInput is an input type that accepts ClientAuthsArgs and ClientAuthsOutput values.
// You can construct a concrete instance of `ClientAuthsInput` via:
//
//          ClientAuthsArgs{...}
type ClientAuthsInput interface {
	pulumi.Input

	ToClientAuthsOutput() ClientAuthsOutput
	ToClientAuthsOutputWithContext(context.Context) ClientAuthsOutput
}

type ClientAuthsArgs struct {
	CertFile pulumi.StringInput `pulumi:"certFile"`
	KeyFile  pulumi.StringInput `pulumi:"keyFile"`
}

func (ClientAuthsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ClientAuths)(nil)).Elem()
}

func (i ClientAuthsArgs) ToClientAuthsOutput() ClientAuthsOutput {
	return i.ToClientAuthsOutputWithContext(context.Background())
}

func (i ClientAuthsArgs) ToClientAuthsOutputWithContext(ctx context.Context) ClientAuthsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClientAuthsOutput)
}

// ClientAuthsArrayInput is an input type that accepts ClientAuthsArray and ClientAuthsArrayOutput values.
// You can construct a concrete instance of `ClientAuthsArrayInput` via:
//
//          ClientAuthsArray{ ClientAuthsArgs{...} }
type ClientAuthsArrayInput interface {
	pulumi.Input

	ToClientAuthsArrayOutput() ClientAuthsArrayOutput
	ToClientAuthsArrayOutputWithContext(context.Context) ClientAuthsArrayOutput
}

type ClientAuthsArray []ClientAuthsInput

func (ClientAuthsArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ClientAuths)(nil)).Elem()
}

func (i ClientAuthsArray) ToClientAuthsArrayOutput() ClientAuthsArrayOutput {
	return i.ToClientAuthsArrayOutputWithContext(context.Background())
}

func (i ClientAuthsArray) ToClientAuthsArrayOutputWithContext(ctx context.Context) ClientAuthsArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClientAuthsArrayOutput)
}

type ClientAuthsOutput struct{ *pulumi.OutputState }

func (ClientAuthsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ClientAuths)(nil)).Elem()
}

func (o ClientAuthsOutput) ToClientAuthsOutput() ClientAuthsOutput {
	return o
}

func (o ClientAuthsOutput) ToClientAuthsOutputWithContext(ctx context.Context) ClientAuthsOutput {
	return o
}

func (o ClientAuthsOutput) CertFile() pulumi.StringOutput {
	return o.ApplyT(func(v ClientAuths) string { return v.CertFile }).(pulumi.StringOutput)
}

func (o ClientAuthsOutput) KeyFile() pulumi.StringOutput {
	return o.ApplyT(func(v ClientAuths) string { return v.KeyFile }).(pulumi.StringOutput)
}

type ClientAuthsArrayOutput struct{ *pulumi.OutputState }

func (ClientAuthsArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ClientAuths)(nil)).Elem()
}

func (o ClientAuthsArrayOutput) ToClientAuthsArrayOutput() ClientAuthsArrayOutput {
	return o
}

func (o ClientAuthsArrayOutput) ToClientAuthsArrayOutputWithContext(ctx context.Context) ClientAuthsArrayOutput {
	return o
}

func (o ClientAuthsArrayOutput) Index(i pulumi.IntInput) ClientAuthsOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ClientAuths {
		return vs[0].([]ClientAuths)[vs[1].(int)]
	}).(ClientAuthsOutput)
}

type Headers struct {
	Name  string `pulumi:"name"`
	Value string `pulumi:"value"`
}

// HeadersInput is an input type that accepts HeadersArgs and HeadersOutput values.
// You can construct a concrete instance of `HeadersInput` via:
//
//          HeadersArgs{...}
type HeadersInput interface {
	pulumi.Input

	ToHeadersOutput() HeadersOutput
	ToHeadersOutputWithContext(context.Context) HeadersOutput
}

type HeadersArgs struct {
	Name  pulumi.StringInput `pulumi:"name"`
	Value pulumi.StringInput `pulumi:"value"`
}

func (HeadersArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Headers)(nil)).Elem()
}

func (i HeadersArgs) ToHeadersOutput() HeadersOutput {
	return i.ToHeadersOutputWithContext(context.Background())
}

func (i HeadersArgs) ToHeadersOutputWithContext(ctx context.Context) HeadersOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HeadersOutput)
}

// HeadersArrayInput is an input type that accepts HeadersArray and HeadersArrayOutput values.
// You can construct a concrete instance of `HeadersArrayInput` via:
//
//          HeadersArray{ HeadersArgs{...} }
type HeadersArrayInput interface {
	pulumi.Input

	ToHeadersArrayOutput() HeadersArrayOutput
	ToHeadersArrayOutputWithContext(context.Context) HeadersArrayOutput
}

type HeadersArray []HeadersInput

func (HeadersArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Headers)(nil)).Elem()
}

func (i HeadersArray) ToHeadersArrayOutput() HeadersArrayOutput {
	return i.ToHeadersArrayOutputWithContext(context.Background())
}

func (i HeadersArray) ToHeadersArrayOutputWithContext(ctx context.Context) HeadersArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HeadersArrayOutput)
}

type HeadersOutput struct{ *pulumi.OutputState }

func (HeadersOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Headers)(nil)).Elem()
}

func (o HeadersOutput) ToHeadersOutput() HeadersOutput {
	return o
}

func (o HeadersOutput) ToHeadersOutputWithContext(ctx context.Context) HeadersOutput {
	return o
}

func (o HeadersOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v Headers) string { return v.Name }).(pulumi.StringOutput)
}

func (o HeadersOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v Headers) string { return v.Value }).(pulumi.StringOutput)
}

type HeadersArrayOutput struct{ *pulumi.OutputState }

func (HeadersArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Headers)(nil)).Elem()
}

func (o HeadersArrayOutput) ToHeadersArrayOutput() HeadersArrayOutput {
	return o
}

func (o HeadersArrayOutput) ToHeadersArrayOutputWithContext(ctx context.Context) HeadersArrayOutput {
	return o
}

func (o HeadersArrayOutput) Index(i pulumi.IntInput) HeadersOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Headers {
		return vs[0].([]Headers)[vs[1].(int)]
	}).(HeadersOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AuthLoginsInput)(nil)).Elem(), AuthLoginsArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*AuthLoginsArrayInput)(nil)).Elem(), AuthLoginsArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ClientAuthsInput)(nil)).Elem(), ClientAuthsArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ClientAuthsArrayInput)(nil)).Elem(), ClientAuthsArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*HeadersInput)(nil)).Elem(), HeadersArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*HeadersArrayInput)(nil)).Elem(), HeadersArray{})
	pulumi.RegisterOutputType(AuthLoginsOutput{})
	pulumi.RegisterOutputType(AuthLoginsArrayOutput{})
	pulumi.RegisterOutputType(ClientAuthsOutput{})
	pulumi.RegisterOutputType(ClientAuthsArrayOutput{})
	pulumi.RegisterOutputType(HeadersOutput{})
	pulumi.RegisterOutputType(HeadersArrayOutput{})
}
