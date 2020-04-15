// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package gcp

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

type SecretRolesetBinding struct {
	// Resource or resource path for which IAM policy information will be bound. The resource path may be specified in a few different [formats](https://www.vaultproject.io/docs/secrets/gcp/index.html#roleset-bindings).
	Resource string `pulumi:"resource"`
	// List of [GCP IAM roles](https://cloud.google.com/iam/docs/understanding-roles) for the resource.
	Roles []string `pulumi:"roles"`
}

// SecretRolesetBindingInput is an input type that accepts SecretRolesetBindingArgs and SecretRolesetBindingOutput values.
// You can construct a concrete instance of `SecretRolesetBindingInput` via:
//
// 		 SecretRolesetBindingArgs{...}
//
type SecretRolesetBindingInput interface {
	pulumi.Input

	ToSecretRolesetBindingOutput() SecretRolesetBindingOutput
	ToSecretRolesetBindingOutputWithContext(context.Context) SecretRolesetBindingOutput
}

type SecretRolesetBindingArgs struct {
	// Resource or resource path for which IAM policy information will be bound. The resource path may be specified in a few different [formats](https://www.vaultproject.io/docs/secrets/gcp/index.html#roleset-bindings).
	Resource pulumi.StringInput `pulumi:"resource"`
	// List of [GCP IAM roles](https://cloud.google.com/iam/docs/understanding-roles) for the resource.
	Roles pulumi.StringArrayInput `pulumi:"roles"`
}

func (SecretRolesetBindingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SecretRolesetBinding)(nil)).Elem()
}

func (i SecretRolesetBindingArgs) ToSecretRolesetBindingOutput() SecretRolesetBindingOutput {
	return i.ToSecretRolesetBindingOutputWithContext(context.Background())
}

func (i SecretRolesetBindingArgs) ToSecretRolesetBindingOutputWithContext(ctx context.Context) SecretRolesetBindingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecretRolesetBindingOutput)
}

// SecretRolesetBindingArrayInput is an input type that accepts SecretRolesetBindingArray and SecretRolesetBindingArrayOutput values.
// You can construct a concrete instance of `SecretRolesetBindingArrayInput` via:
//
// 		 SecretRolesetBindingArray{ SecretRolesetBindingArgs{...} }
//
type SecretRolesetBindingArrayInput interface {
	pulumi.Input

	ToSecretRolesetBindingArrayOutput() SecretRolesetBindingArrayOutput
	ToSecretRolesetBindingArrayOutputWithContext(context.Context) SecretRolesetBindingArrayOutput
}

type SecretRolesetBindingArray []SecretRolesetBindingInput

func (SecretRolesetBindingArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SecretRolesetBinding)(nil)).Elem()
}

func (i SecretRolesetBindingArray) ToSecretRolesetBindingArrayOutput() SecretRolesetBindingArrayOutput {
	return i.ToSecretRolesetBindingArrayOutputWithContext(context.Background())
}

func (i SecretRolesetBindingArray) ToSecretRolesetBindingArrayOutputWithContext(ctx context.Context) SecretRolesetBindingArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecretRolesetBindingArrayOutput)
}

type SecretRolesetBindingOutput struct{ *pulumi.OutputState }

func (SecretRolesetBindingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SecretRolesetBinding)(nil)).Elem()
}

func (o SecretRolesetBindingOutput) ToSecretRolesetBindingOutput() SecretRolesetBindingOutput {
	return o
}

func (o SecretRolesetBindingOutput) ToSecretRolesetBindingOutputWithContext(ctx context.Context) SecretRolesetBindingOutput {
	return o
}

// Resource or resource path for which IAM policy information will be bound. The resource path may be specified in a few different [formats](https://www.vaultproject.io/docs/secrets/gcp/index.html#roleset-bindings).
func (o SecretRolesetBindingOutput) Resource() pulumi.StringOutput {
	return o.ApplyT(func(v SecretRolesetBinding) string { return v.Resource }).(pulumi.StringOutput)
}

// List of [GCP IAM roles](https://cloud.google.com/iam/docs/understanding-roles) for the resource.
func (o SecretRolesetBindingOutput) Roles() pulumi.StringArrayOutput {
	return o.ApplyT(func(v SecretRolesetBinding) []string { return v.Roles }).(pulumi.StringArrayOutput)
}

type SecretRolesetBindingArrayOutput struct{ *pulumi.OutputState }

func (SecretRolesetBindingArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SecretRolesetBinding)(nil)).Elem()
}

func (o SecretRolesetBindingArrayOutput) ToSecretRolesetBindingArrayOutput() SecretRolesetBindingArrayOutput {
	return o
}

func (o SecretRolesetBindingArrayOutput) ToSecretRolesetBindingArrayOutputWithContext(ctx context.Context) SecretRolesetBindingArrayOutput {
	return o
}

func (o SecretRolesetBindingArrayOutput) Index(i pulumi.IntInput) SecretRolesetBindingOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SecretRolesetBinding {
		return vs[0].([]SecretRolesetBinding)[vs[1].(int)]
	}).(SecretRolesetBindingOutput)
}

func init() {
	pulumi.RegisterOutputType(SecretRolesetBindingOutput{})
	pulumi.RegisterOutputType(SecretRolesetBindingArrayOutput{})
}
