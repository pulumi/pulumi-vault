// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package okta

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AuthBackendGroupType struct {
	// Name of the group within the Okta
	GroupName string `pulumi:"groupName"`
	// List of Vault policies to associate with this user
	Policies []string `pulumi:"policies"`
}

// AuthBackendGroupTypeInput is an input type that accepts AuthBackendGroupTypeArgs and AuthBackendGroupTypeOutput values.
// You can construct a concrete instance of `AuthBackendGroupTypeInput` via:
//
//          AuthBackendGroupTypeArgs{...}
type AuthBackendGroupTypeInput interface {
	pulumi.Input

	ToAuthBackendGroupTypeOutput() AuthBackendGroupTypeOutput
	ToAuthBackendGroupTypeOutputWithContext(context.Context) AuthBackendGroupTypeOutput
}

type AuthBackendGroupTypeArgs struct {
	// Name of the group within the Okta
	GroupName pulumi.StringInput `pulumi:"groupName"`
	// List of Vault policies to associate with this user
	Policies pulumi.StringArrayInput `pulumi:"policies"`
}

func (AuthBackendGroupTypeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AuthBackendGroupType)(nil)).Elem()
}

func (i AuthBackendGroupTypeArgs) ToAuthBackendGroupTypeOutput() AuthBackendGroupTypeOutput {
	return i.ToAuthBackendGroupTypeOutputWithContext(context.Background())
}

func (i AuthBackendGroupTypeArgs) ToAuthBackendGroupTypeOutputWithContext(ctx context.Context) AuthBackendGroupTypeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AuthBackendGroupTypeOutput)
}

// AuthBackendGroupTypeArrayInput is an input type that accepts AuthBackendGroupTypeArray and AuthBackendGroupTypeArrayOutput values.
// You can construct a concrete instance of `AuthBackendGroupTypeArrayInput` via:
//
//          AuthBackendGroupTypeArray{ AuthBackendGroupTypeArgs{...} }
type AuthBackendGroupTypeArrayInput interface {
	pulumi.Input

	ToAuthBackendGroupTypeArrayOutput() AuthBackendGroupTypeArrayOutput
	ToAuthBackendGroupTypeArrayOutputWithContext(context.Context) AuthBackendGroupTypeArrayOutput
}

type AuthBackendGroupTypeArray []AuthBackendGroupTypeInput

func (AuthBackendGroupTypeArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AuthBackendGroupType)(nil)).Elem()
}

func (i AuthBackendGroupTypeArray) ToAuthBackendGroupTypeArrayOutput() AuthBackendGroupTypeArrayOutput {
	return i.ToAuthBackendGroupTypeArrayOutputWithContext(context.Background())
}

func (i AuthBackendGroupTypeArray) ToAuthBackendGroupTypeArrayOutputWithContext(ctx context.Context) AuthBackendGroupTypeArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AuthBackendGroupTypeArrayOutput)
}

type AuthBackendGroupTypeOutput struct{ *pulumi.OutputState }

func (AuthBackendGroupTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AuthBackendGroupType)(nil)).Elem()
}

func (o AuthBackendGroupTypeOutput) ToAuthBackendGroupTypeOutput() AuthBackendGroupTypeOutput {
	return o
}

func (o AuthBackendGroupTypeOutput) ToAuthBackendGroupTypeOutputWithContext(ctx context.Context) AuthBackendGroupTypeOutput {
	return o
}

// Name of the group within the Okta
func (o AuthBackendGroupTypeOutput) GroupName() pulumi.StringOutput {
	return o.ApplyT(func(v AuthBackendGroupType) string { return v.GroupName }).(pulumi.StringOutput)
}

// List of Vault policies to associate with this user
func (o AuthBackendGroupTypeOutput) Policies() pulumi.StringArrayOutput {
	return o.ApplyT(func(v AuthBackendGroupType) []string { return v.Policies }).(pulumi.StringArrayOutput)
}

type AuthBackendGroupTypeArrayOutput struct{ *pulumi.OutputState }

func (AuthBackendGroupTypeArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AuthBackendGroupType)(nil)).Elem()
}

func (o AuthBackendGroupTypeArrayOutput) ToAuthBackendGroupTypeArrayOutput() AuthBackendGroupTypeArrayOutput {
	return o
}

func (o AuthBackendGroupTypeArrayOutput) ToAuthBackendGroupTypeArrayOutputWithContext(ctx context.Context) AuthBackendGroupTypeArrayOutput {
	return o
}

func (o AuthBackendGroupTypeArrayOutput) Index(i pulumi.IntInput) AuthBackendGroupTypeOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) AuthBackendGroupType {
		return vs[0].([]AuthBackendGroupType)[vs[1].(int)]
	}).(AuthBackendGroupTypeOutput)
}

type AuthBackendUserType struct {
	// List of Okta groups to associate with this user
	Groups []string `pulumi:"groups"`
	// List of Vault policies to associate with this user
	Policies []string `pulumi:"policies"`
	// Name of the user within Okta
	Username string `pulumi:"username"`
}

// AuthBackendUserTypeInput is an input type that accepts AuthBackendUserTypeArgs and AuthBackendUserTypeOutput values.
// You can construct a concrete instance of `AuthBackendUserTypeInput` via:
//
//          AuthBackendUserTypeArgs{...}
type AuthBackendUserTypeInput interface {
	pulumi.Input

	ToAuthBackendUserTypeOutput() AuthBackendUserTypeOutput
	ToAuthBackendUserTypeOutputWithContext(context.Context) AuthBackendUserTypeOutput
}

type AuthBackendUserTypeArgs struct {
	// List of Okta groups to associate with this user
	Groups pulumi.StringArrayInput `pulumi:"groups"`
	// List of Vault policies to associate with this user
	Policies pulumi.StringArrayInput `pulumi:"policies"`
	// Name of the user within Okta
	Username pulumi.StringInput `pulumi:"username"`
}

func (AuthBackendUserTypeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AuthBackendUserType)(nil)).Elem()
}

func (i AuthBackendUserTypeArgs) ToAuthBackendUserTypeOutput() AuthBackendUserTypeOutput {
	return i.ToAuthBackendUserTypeOutputWithContext(context.Background())
}

func (i AuthBackendUserTypeArgs) ToAuthBackendUserTypeOutputWithContext(ctx context.Context) AuthBackendUserTypeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AuthBackendUserTypeOutput)
}

// AuthBackendUserTypeArrayInput is an input type that accepts AuthBackendUserTypeArray and AuthBackendUserTypeArrayOutput values.
// You can construct a concrete instance of `AuthBackendUserTypeArrayInput` via:
//
//          AuthBackendUserTypeArray{ AuthBackendUserTypeArgs{...} }
type AuthBackendUserTypeArrayInput interface {
	pulumi.Input

	ToAuthBackendUserTypeArrayOutput() AuthBackendUserTypeArrayOutput
	ToAuthBackendUserTypeArrayOutputWithContext(context.Context) AuthBackendUserTypeArrayOutput
}

type AuthBackendUserTypeArray []AuthBackendUserTypeInput

func (AuthBackendUserTypeArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AuthBackendUserType)(nil)).Elem()
}

func (i AuthBackendUserTypeArray) ToAuthBackendUserTypeArrayOutput() AuthBackendUserTypeArrayOutput {
	return i.ToAuthBackendUserTypeArrayOutputWithContext(context.Background())
}

func (i AuthBackendUserTypeArray) ToAuthBackendUserTypeArrayOutputWithContext(ctx context.Context) AuthBackendUserTypeArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AuthBackendUserTypeArrayOutput)
}

type AuthBackendUserTypeOutput struct{ *pulumi.OutputState }

func (AuthBackendUserTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AuthBackendUserType)(nil)).Elem()
}

func (o AuthBackendUserTypeOutput) ToAuthBackendUserTypeOutput() AuthBackendUserTypeOutput {
	return o
}

func (o AuthBackendUserTypeOutput) ToAuthBackendUserTypeOutputWithContext(ctx context.Context) AuthBackendUserTypeOutput {
	return o
}

// List of Okta groups to associate with this user
func (o AuthBackendUserTypeOutput) Groups() pulumi.StringArrayOutput {
	return o.ApplyT(func(v AuthBackendUserType) []string { return v.Groups }).(pulumi.StringArrayOutput)
}

// List of Vault policies to associate with this user
func (o AuthBackendUserTypeOutput) Policies() pulumi.StringArrayOutput {
	return o.ApplyT(func(v AuthBackendUserType) []string { return v.Policies }).(pulumi.StringArrayOutput)
}

// Name of the user within Okta
func (o AuthBackendUserTypeOutput) Username() pulumi.StringOutput {
	return o.ApplyT(func(v AuthBackendUserType) string { return v.Username }).(pulumi.StringOutput)
}

type AuthBackendUserTypeArrayOutput struct{ *pulumi.OutputState }

func (AuthBackendUserTypeArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AuthBackendUserType)(nil)).Elem()
}

func (o AuthBackendUserTypeArrayOutput) ToAuthBackendUserTypeArrayOutput() AuthBackendUserTypeArrayOutput {
	return o
}

func (o AuthBackendUserTypeArrayOutput) ToAuthBackendUserTypeArrayOutputWithContext(ctx context.Context) AuthBackendUserTypeArrayOutput {
	return o
}

func (o AuthBackendUserTypeArrayOutput) Index(i pulumi.IntInput) AuthBackendUserTypeOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) AuthBackendUserType {
		return vs[0].([]AuthBackendUserType)[vs[1].(int)]
	}).(AuthBackendUserTypeOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AuthBackendGroupTypeInput)(nil)).Elem(), AuthBackendGroupTypeArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*AuthBackendGroupTypeArrayInput)(nil)).Elem(), AuthBackendGroupTypeArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AuthBackendUserTypeInput)(nil)).Elem(), AuthBackendUserTypeArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*AuthBackendUserTypeArrayInput)(nil)).Elem(), AuthBackendUserTypeArray{})
	pulumi.RegisterOutputType(AuthBackendGroupTypeOutput{})
	pulumi.RegisterOutputType(AuthBackendGroupTypeArrayOutput{})
	pulumi.RegisterOutputType(AuthBackendUserTypeOutput{})
	pulumi.RegisterOutputType(AuthBackendUserTypeArrayOutput{})
}
