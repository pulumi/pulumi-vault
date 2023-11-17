// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package aws

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-vault/sdk/v5/go/vault/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configures the periodic tidying operation of the blacklisted role tag entries.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-vault/sdk/v5/go/vault"
//	"github.com/pulumi/pulumi-vault/sdk/v5/go/vault/aws"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			exampleAuthBackend, err := vault.NewAuthBackend(ctx, "exampleAuthBackend", &vault.AuthBackendArgs{
//				Type: pulumi.String("aws"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = aws.NewAuthBackendRoletagBlacklist(ctx, "exampleAuthBackendRoletagBlacklist", &aws.AuthBackendRoletagBlacklistArgs{
//				Backend:      exampleAuthBackend.Path,
//				SafetyBuffer: pulumi.Int(360),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
type AuthBackendRoletagBlacklist struct {
	pulumi.CustomResourceState

	// The path the AWS auth backend being configured was
	// mounted at.
	Backend pulumi.StringOutput `pulumi:"backend"`
	// If set to true, disables the periodic
	// tidying of the roletag blacklist entries. Defaults to false.
	DisablePeriodicTidy pulumi.BoolPtrOutput `pulumi:"disablePeriodicTidy"`
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
	// *Available only for Vault Enterprise*.
	Namespace pulumi.StringPtrOutput `pulumi:"namespace"`
	// The amount of extra time that must have passed
	// beyond the roletag expiration, before it is removed from the backend storage.
	// Defaults to 259,200 seconds, or 72 hours.
	SafetyBuffer pulumi.IntPtrOutput `pulumi:"safetyBuffer"`
}

// NewAuthBackendRoletagBlacklist registers a new resource with the given unique name, arguments, and options.
func NewAuthBackendRoletagBlacklist(ctx *pulumi.Context,
	name string, args *AuthBackendRoletagBlacklistArgs, opts ...pulumi.ResourceOption) (*AuthBackendRoletagBlacklist, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Backend == nil {
		return nil, errors.New("invalid value for required argument 'Backend'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource AuthBackendRoletagBlacklist
	err := ctx.RegisterResource("vault:aws/authBackendRoletagBlacklist:AuthBackendRoletagBlacklist", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAuthBackendRoletagBlacklist gets an existing AuthBackendRoletagBlacklist resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAuthBackendRoletagBlacklist(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AuthBackendRoletagBlacklistState, opts ...pulumi.ResourceOption) (*AuthBackendRoletagBlacklist, error) {
	var resource AuthBackendRoletagBlacklist
	err := ctx.ReadResource("vault:aws/authBackendRoletagBlacklist:AuthBackendRoletagBlacklist", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AuthBackendRoletagBlacklist resources.
type authBackendRoletagBlacklistState struct {
	// The path the AWS auth backend being configured was
	// mounted at.
	Backend *string `pulumi:"backend"`
	// If set to true, disables the periodic
	// tidying of the roletag blacklist entries. Defaults to false.
	DisablePeriodicTidy *bool `pulumi:"disablePeriodicTidy"`
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
	// *Available only for Vault Enterprise*.
	Namespace *string `pulumi:"namespace"`
	// The amount of extra time that must have passed
	// beyond the roletag expiration, before it is removed from the backend storage.
	// Defaults to 259,200 seconds, or 72 hours.
	SafetyBuffer *int `pulumi:"safetyBuffer"`
}

type AuthBackendRoletagBlacklistState struct {
	// The path the AWS auth backend being configured was
	// mounted at.
	Backend pulumi.StringPtrInput
	// If set to true, disables the periodic
	// tidying of the roletag blacklist entries. Defaults to false.
	DisablePeriodicTidy pulumi.BoolPtrInput
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
	// *Available only for Vault Enterprise*.
	Namespace pulumi.StringPtrInput
	// The amount of extra time that must have passed
	// beyond the roletag expiration, before it is removed from the backend storage.
	// Defaults to 259,200 seconds, or 72 hours.
	SafetyBuffer pulumi.IntPtrInput
}

func (AuthBackendRoletagBlacklistState) ElementType() reflect.Type {
	return reflect.TypeOf((*authBackendRoletagBlacklistState)(nil)).Elem()
}

type authBackendRoletagBlacklistArgs struct {
	// The path the AWS auth backend being configured was
	// mounted at.
	Backend string `pulumi:"backend"`
	// If set to true, disables the periodic
	// tidying of the roletag blacklist entries. Defaults to false.
	DisablePeriodicTidy *bool `pulumi:"disablePeriodicTidy"`
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
	// *Available only for Vault Enterprise*.
	Namespace *string `pulumi:"namespace"`
	// The amount of extra time that must have passed
	// beyond the roletag expiration, before it is removed from the backend storage.
	// Defaults to 259,200 seconds, or 72 hours.
	SafetyBuffer *int `pulumi:"safetyBuffer"`
}

// The set of arguments for constructing a AuthBackendRoletagBlacklist resource.
type AuthBackendRoletagBlacklistArgs struct {
	// The path the AWS auth backend being configured was
	// mounted at.
	Backend pulumi.StringInput
	// If set to true, disables the periodic
	// tidying of the roletag blacklist entries. Defaults to false.
	DisablePeriodicTidy pulumi.BoolPtrInput
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
	// *Available only for Vault Enterprise*.
	Namespace pulumi.StringPtrInput
	// The amount of extra time that must have passed
	// beyond the roletag expiration, before it is removed from the backend storage.
	// Defaults to 259,200 seconds, or 72 hours.
	SafetyBuffer pulumi.IntPtrInput
}

func (AuthBackendRoletagBlacklistArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*authBackendRoletagBlacklistArgs)(nil)).Elem()
}

type AuthBackendRoletagBlacklistInput interface {
	pulumi.Input

	ToAuthBackendRoletagBlacklistOutput() AuthBackendRoletagBlacklistOutput
	ToAuthBackendRoletagBlacklistOutputWithContext(ctx context.Context) AuthBackendRoletagBlacklistOutput
}

func (*AuthBackendRoletagBlacklist) ElementType() reflect.Type {
	return reflect.TypeOf((**AuthBackendRoletagBlacklist)(nil)).Elem()
}

func (i *AuthBackendRoletagBlacklist) ToAuthBackendRoletagBlacklistOutput() AuthBackendRoletagBlacklistOutput {
	return i.ToAuthBackendRoletagBlacklistOutputWithContext(context.Background())
}

func (i *AuthBackendRoletagBlacklist) ToAuthBackendRoletagBlacklistOutputWithContext(ctx context.Context) AuthBackendRoletagBlacklistOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AuthBackendRoletagBlacklistOutput)
}

// AuthBackendRoletagBlacklistArrayInput is an input type that accepts AuthBackendRoletagBlacklistArray and AuthBackendRoletagBlacklistArrayOutput values.
// You can construct a concrete instance of `AuthBackendRoletagBlacklistArrayInput` via:
//
//	AuthBackendRoletagBlacklistArray{ AuthBackendRoletagBlacklistArgs{...} }
type AuthBackendRoletagBlacklistArrayInput interface {
	pulumi.Input

	ToAuthBackendRoletagBlacklistArrayOutput() AuthBackendRoletagBlacklistArrayOutput
	ToAuthBackendRoletagBlacklistArrayOutputWithContext(context.Context) AuthBackendRoletagBlacklistArrayOutput
}

type AuthBackendRoletagBlacklistArray []AuthBackendRoletagBlacklistInput

func (AuthBackendRoletagBlacklistArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AuthBackendRoletagBlacklist)(nil)).Elem()
}

func (i AuthBackendRoletagBlacklistArray) ToAuthBackendRoletagBlacklistArrayOutput() AuthBackendRoletagBlacklistArrayOutput {
	return i.ToAuthBackendRoletagBlacklistArrayOutputWithContext(context.Background())
}

func (i AuthBackendRoletagBlacklistArray) ToAuthBackendRoletagBlacklistArrayOutputWithContext(ctx context.Context) AuthBackendRoletagBlacklistArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AuthBackendRoletagBlacklistArrayOutput)
}

// AuthBackendRoletagBlacklistMapInput is an input type that accepts AuthBackendRoletagBlacklistMap and AuthBackendRoletagBlacklistMapOutput values.
// You can construct a concrete instance of `AuthBackendRoletagBlacklistMapInput` via:
//
//	AuthBackendRoletagBlacklistMap{ "key": AuthBackendRoletagBlacklistArgs{...} }
type AuthBackendRoletagBlacklistMapInput interface {
	pulumi.Input

	ToAuthBackendRoletagBlacklistMapOutput() AuthBackendRoletagBlacklistMapOutput
	ToAuthBackendRoletagBlacklistMapOutputWithContext(context.Context) AuthBackendRoletagBlacklistMapOutput
}

type AuthBackendRoletagBlacklistMap map[string]AuthBackendRoletagBlacklistInput

func (AuthBackendRoletagBlacklistMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AuthBackendRoletagBlacklist)(nil)).Elem()
}

func (i AuthBackendRoletagBlacklistMap) ToAuthBackendRoletagBlacklistMapOutput() AuthBackendRoletagBlacklistMapOutput {
	return i.ToAuthBackendRoletagBlacklistMapOutputWithContext(context.Background())
}

func (i AuthBackendRoletagBlacklistMap) ToAuthBackendRoletagBlacklistMapOutputWithContext(ctx context.Context) AuthBackendRoletagBlacklistMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AuthBackendRoletagBlacklistMapOutput)
}

type AuthBackendRoletagBlacklistOutput struct{ *pulumi.OutputState }

func (AuthBackendRoletagBlacklistOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AuthBackendRoletagBlacklist)(nil)).Elem()
}

func (o AuthBackendRoletagBlacklistOutput) ToAuthBackendRoletagBlacklistOutput() AuthBackendRoletagBlacklistOutput {
	return o
}

func (o AuthBackendRoletagBlacklistOutput) ToAuthBackendRoletagBlacklistOutputWithContext(ctx context.Context) AuthBackendRoletagBlacklistOutput {
	return o
}

// The path the AWS auth backend being configured was
// mounted at.
func (o AuthBackendRoletagBlacklistOutput) Backend() pulumi.StringOutput {
	return o.ApplyT(func(v *AuthBackendRoletagBlacklist) pulumi.StringOutput { return v.Backend }).(pulumi.StringOutput)
}

// If set to true, disables the periodic
// tidying of the roletag blacklist entries. Defaults to false.
func (o AuthBackendRoletagBlacklistOutput) DisablePeriodicTidy() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AuthBackendRoletagBlacklist) pulumi.BoolPtrOutput { return v.DisablePeriodicTidy }).(pulumi.BoolPtrOutput)
}

// The namespace to provision the resource in.
// The value should not contain leading or trailing forward slashes.
// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
// *Available only for Vault Enterprise*.
func (o AuthBackendRoletagBlacklistOutput) Namespace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AuthBackendRoletagBlacklist) pulumi.StringPtrOutput { return v.Namespace }).(pulumi.StringPtrOutput)
}

// The amount of extra time that must have passed
// beyond the roletag expiration, before it is removed from the backend storage.
// Defaults to 259,200 seconds, or 72 hours.
func (o AuthBackendRoletagBlacklistOutput) SafetyBuffer() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *AuthBackendRoletagBlacklist) pulumi.IntPtrOutput { return v.SafetyBuffer }).(pulumi.IntPtrOutput)
}

type AuthBackendRoletagBlacklistArrayOutput struct{ *pulumi.OutputState }

func (AuthBackendRoletagBlacklistArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AuthBackendRoletagBlacklist)(nil)).Elem()
}

func (o AuthBackendRoletagBlacklistArrayOutput) ToAuthBackendRoletagBlacklistArrayOutput() AuthBackendRoletagBlacklistArrayOutput {
	return o
}

func (o AuthBackendRoletagBlacklistArrayOutput) ToAuthBackendRoletagBlacklistArrayOutputWithContext(ctx context.Context) AuthBackendRoletagBlacklistArrayOutput {
	return o
}

func (o AuthBackendRoletagBlacklistArrayOutput) Index(i pulumi.IntInput) AuthBackendRoletagBlacklistOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *AuthBackendRoletagBlacklist {
		return vs[0].([]*AuthBackendRoletagBlacklist)[vs[1].(int)]
	}).(AuthBackendRoletagBlacklistOutput)
}

type AuthBackendRoletagBlacklistMapOutput struct{ *pulumi.OutputState }

func (AuthBackendRoletagBlacklistMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AuthBackendRoletagBlacklist)(nil)).Elem()
}

func (o AuthBackendRoletagBlacklistMapOutput) ToAuthBackendRoletagBlacklistMapOutput() AuthBackendRoletagBlacklistMapOutput {
	return o
}

func (o AuthBackendRoletagBlacklistMapOutput) ToAuthBackendRoletagBlacklistMapOutputWithContext(ctx context.Context) AuthBackendRoletagBlacklistMapOutput {
	return o
}

func (o AuthBackendRoletagBlacklistMapOutput) MapIndex(k pulumi.StringInput) AuthBackendRoletagBlacklistOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *AuthBackendRoletagBlacklist {
		return vs[0].(map[string]*AuthBackendRoletagBlacklist)[vs[1].(string)]
	}).(AuthBackendRoletagBlacklistOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AuthBackendRoletagBlacklistInput)(nil)).Elem(), &AuthBackendRoletagBlacklist{})
	pulumi.RegisterInputType(reflect.TypeOf((*AuthBackendRoletagBlacklistArrayInput)(nil)).Elem(), AuthBackendRoletagBlacklistArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AuthBackendRoletagBlacklistMapInput)(nil)).Elem(), AuthBackendRoletagBlacklistMap{})
	pulumi.RegisterOutputType(AuthBackendRoletagBlacklistOutput{})
	pulumi.RegisterOutputType(AuthBackendRoletagBlacklistArrayOutput{})
	pulumi.RegisterOutputType(AuthBackendRoletagBlacklistMapOutput{})
}
