// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package aws

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configures the periodic tidying operation of the whitelisted identity entries.
//
// For more information, see the
// [Vault docs](https://www.vaultproject.io/api-docs/auth/aws#configure-identity-whitelist-tidy-operation).
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-vault/sdk/v5/go/vault"
// 	"github.com/pulumi/pulumi-vault/sdk/v5/go/vault/aws"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		exampleAuthBackend, err := vault.NewAuthBackend(ctx, "exampleAuthBackend", &vault.AuthBackendArgs{
// 			Type: pulumi.String("aws"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = aws.NewAuthBackendIdentityWhitelist(ctx, "exampleAuthBackendIdentityWhitelist", &aws.AuthBackendIdentityWhitelistArgs{
// 			Backend:      exampleAuthBackend.Path,
// 			SafetyBuffer: pulumi.Int(3600),
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
// AWS auth backend identity whitelists can be imported using `auth/`, the `backend` path, and `/config/tidy/identity-whitelist` e.g.
//
// ```sh
//  $ pulumi import vault:aws/authBackendIdentityWhitelist:AuthBackendIdentityWhitelist example auth/aws/config/tidy/identity-whitelist
// ```
type AuthBackendIdentityWhitelist struct {
	pulumi.CustomResourceState

	// The path of the AWS backend being configured.
	Backend pulumi.StringPtrOutput `pulumi:"backend"`
	// If set to true, disables the periodic
	// tidying of the identity-whitelist entries.
	DisablePeriodicTidy pulumi.BoolPtrOutput `pulumi:"disablePeriodicTidy"`
	// The amount of extra time, in minutes, that must
	// have passed beyond the roletag expiration, before it is removed from the
	// backend storage.
	SafetyBuffer pulumi.IntPtrOutput `pulumi:"safetyBuffer"`
}

// NewAuthBackendIdentityWhitelist registers a new resource with the given unique name, arguments, and options.
func NewAuthBackendIdentityWhitelist(ctx *pulumi.Context,
	name string, args *AuthBackendIdentityWhitelistArgs, opts ...pulumi.ResourceOption) (*AuthBackendIdentityWhitelist, error) {
	if args == nil {
		args = &AuthBackendIdentityWhitelistArgs{}
	}

	var resource AuthBackendIdentityWhitelist
	err := ctx.RegisterResource("vault:aws/authBackendIdentityWhitelist:AuthBackendIdentityWhitelist", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAuthBackendIdentityWhitelist gets an existing AuthBackendIdentityWhitelist resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAuthBackendIdentityWhitelist(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AuthBackendIdentityWhitelistState, opts ...pulumi.ResourceOption) (*AuthBackendIdentityWhitelist, error) {
	var resource AuthBackendIdentityWhitelist
	err := ctx.ReadResource("vault:aws/authBackendIdentityWhitelist:AuthBackendIdentityWhitelist", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AuthBackendIdentityWhitelist resources.
type authBackendIdentityWhitelistState struct {
	// The path of the AWS backend being configured.
	Backend *string `pulumi:"backend"`
	// If set to true, disables the periodic
	// tidying of the identity-whitelist entries.
	DisablePeriodicTidy *bool `pulumi:"disablePeriodicTidy"`
	// The amount of extra time, in minutes, that must
	// have passed beyond the roletag expiration, before it is removed from the
	// backend storage.
	SafetyBuffer *int `pulumi:"safetyBuffer"`
}

type AuthBackendIdentityWhitelistState struct {
	// The path of the AWS backend being configured.
	Backend pulumi.StringPtrInput
	// If set to true, disables the periodic
	// tidying of the identity-whitelist entries.
	DisablePeriodicTidy pulumi.BoolPtrInput
	// The amount of extra time, in minutes, that must
	// have passed beyond the roletag expiration, before it is removed from the
	// backend storage.
	SafetyBuffer pulumi.IntPtrInput
}

func (AuthBackendIdentityWhitelistState) ElementType() reflect.Type {
	return reflect.TypeOf((*authBackendIdentityWhitelistState)(nil)).Elem()
}

type authBackendIdentityWhitelistArgs struct {
	// The path of the AWS backend being configured.
	Backend *string `pulumi:"backend"`
	// If set to true, disables the periodic
	// tidying of the identity-whitelist entries.
	DisablePeriodicTidy *bool `pulumi:"disablePeriodicTidy"`
	// The amount of extra time, in minutes, that must
	// have passed beyond the roletag expiration, before it is removed from the
	// backend storage.
	SafetyBuffer *int `pulumi:"safetyBuffer"`
}

// The set of arguments for constructing a AuthBackendIdentityWhitelist resource.
type AuthBackendIdentityWhitelistArgs struct {
	// The path of the AWS backend being configured.
	Backend pulumi.StringPtrInput
	// If set to true, disables the periodic
	// tidying of the identity-whitelist entries.
	DisablePeriodicTidy pulumi.BoolPtrInput
	// The amount of extra time, in minutes, that must
	// have passed beyond the roletag expiration, before it is removed from the
	// backend storage.
	SafetyBuffer pulumi.IntPtrInput
}

func (AuthBackendIdentityWhitelistArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*authBackendIdentityWhitelistArgs)(nil)).Elem()
}

type AuthBackendIdentityWhitelistInput interface {
	pulumi.Input

	ToAuthBackendIdentityWhitelistOutput() AuthBackendIdentityWhitelistOutput
	ToAuthBackendIdentityWhitelistOutputWithContext(ctx context.Context) AuthBackendIdentityWhitelistOutput
}

func (*AuthBackendIdentityWhitelist) ElementType() reflect.Type {
	return reflect.TypeOf((**AuthBackendIdentityWhitelist)(nil)).Elem()
}

func (i *AuthBackendIdentityWhitelist) ToAuthBackendIdentityWhitelistOutput() AuthBackendIdentityWhitelistOutput {
	return i.ToAuthBackendIdentityWhitelistOutputWithContext(context.Background())
}

func (i *AuthBackendIdentityWhitelist) ToAuthBackendIdentityWhitelistOutputWithContext(ctx context.Context) AuthBackendIdentityWhitelistOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AuthBackendIdentityWhitelistOutput)
}

// AuthBackendIdentityWhitelistArrayInput is an input type that accepts AuthBackendIdentityWhitelistArray and AuthBackendIdentityWhitelistArrayOutput values.
// You can construct a concrete instance of `AuthBackendIdentityWhitelistArrayInput` via:
//
//          AuthBackendIdentityWhitelistArray{ AuthBackendIdentityWhitelistArgs{...} }
type AuthBackendIdentityWhitelistArrayInput interface {
	pulumi.Input

	ToAuthBackendIdentityWhitelistArrayOutput() AuthBackendIdentityWhitelistArrayOutput
	ToAuthBackendIdentityWhitelistArrayOutputWithContext(context.Context) AuthBackendIdentityWhitelistArrayOutput
}

type AuthBackendIdentityWhitelistArray []AuthBackendIdentityWhitelistInput

func (AuthBackendIdentityWhitelistArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AuthBackendIdentityWhitelist)(nil)).Elem()
}

func (i AuthBackendIdentityWhitelistArray) ToAuthBackendIdentityWhitelistArrayOutput() AuthBackendIdentityWhitelistArrayOutput {
	return i.ToAuthBackendIdentityWhitelistArrayOutputWithContext(context.Background())
}

func (i AuthBackendIdentityWhitelistArray) ToAuthBackendIdentityWhitelistArrayOutputWithContext(ctx context.Context) AuthBackendIdentityWhitelistArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AuthBackendIdentityWhitelistArrayOutput)
}

// AuthBackendIdentityWhitelistMapInput is an input type that accepts AuthBackendIdentityWhitelistMap and AuthBackendIdentityWhitelistMapOutput values.
// You can construct a concrete instance of `AuthBackendIdentityWhitelistMapInput` via:
//
//          AuthBackendIdentityWhitelistMap{ "key": AuthBackendIdentityWhitelistArgs{...} }
type AuthBackendIdentityWhitelistMapInput interface {
	pulumi.Input

	ToAuthBackendIdentityWhitelistMapOutput() AuthBackendIdentityWhitelistMapOutput
	ToAuthBackendIdentityWhitelistMapOutputWithContext(context.Context) AuthBackendIdentityWhitelistMapOutput
}

type AuthBackendIdentityWhitelistMap map[string]AuthBackendIdentityWhitelistInput

func (AuthBackendIdentityWhitelistMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AuthBackendIdentityWhitelist)(nil)).Elem()
}

func (i AuthBackendIdentityWhitelistMap) ToAuthBackendIdentityWhitelistMapOutput() AuthBackendIdentityWhitelistMapOutput {
	return i.ToAuthBackendIdentityWhitelistMapOutputWithContext(context.Background())
}

func (i AuthBackendIdentityWhitelistMap) ToAuthBackendIdentityWhitelistMapOutputWithContext(ctx context.Context) AuthBackendIdentityWhitelistMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AuthBackendIdentityWhitelistMapOutput)
}

type AuthBackendIdentityWhitelistOutput struct{ *pulumi.OutputState }

func (AuthBackendIdentityWhitelistOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AuthBackendIdentityWhitelist)(nil)).Elem()
}

func (o AuthBackendIdentityWhitelistOutput) ToAuthBackendIdentityWhitelistOutput() AuthBackendIdentityWhitelistOutput {
	return o
}

func (o AuthBackendIdentityWhitelistOutput) ToAuthBackendIdentityWhitelistOutputWithContext(ctx context.Context) AuthBackendIdentityWhitelistOutput {
	return o
}

type AuthBackendIdentityWhitelistArrayOutput struct{ *pulumi.OutputState }

func (AuthBackendIdentityWhitelistArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AuthBackendIdentityWhitelist)(nil)).Elem()
}

func (o AuthBackendIdentityWhitelistArrayOutput) ToAuthBackendIdentityWhitelistArrayOutput() AuthBackendIdentityWhitelistArrayOutput {
	return o
}

func (o AuthBackendIdentityWhitelistArrayOutput) ToAuthBackendIdentityWhitelistArrayOutputWithContext(ctx context.Context) AuthBackendIdentityWhitelistArrayOutput {
	return o
}

func (o AuthBackendIdentityWhitelistArrayOutput) Index(i pulumi.IntInput) AuthBackendIdentityWhitelistOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *AuthBackendIdentityWhitelist {
		return vs[0].([]*AuthBackendIdentityWhitelist)[vs[1].(int)]
	}).(AuthBackendIdentityWhitelistOutput)
}

type AuthBackendIdentityWhitelistMapOutput struct{ *pulumi.OutputState }

func (AuthBackendIdentityWhitelistMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AuthBackendIdentityWhitelist)(nil)).Elem()
}

func (o AuthBackendIdentityWhitelistMapOutput) ToAuthBackendIdentityWhitelistMapOutput() AuthBackendIdentityWhitelistMapOutput {
	return o
}

func (o AuthBackendIdentityWhitelistMapOutput) ToAuthBackendIdentityWhitelistMapOutputWithContext(ctx context.Context) AuthBackendIdentityWhitelistMapOutput {
	return o
}

func (o AuthBackendIdentityWhitelistMapOutput) MapIndex(k pulumi.StringInput) AuthBackendIdentityWhitelistOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *AuthBackendIdentityWhitelist {
		return vs[0].(map[string]*AuthBackendIdentityWhitelist)[vs[1].(string)]
	}).(AuthBackendIdentityWhitelistOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AuthBackendIdentityWhitelistInput)(nil)).Elem(), &AuthBackendIdentityWhitelist{})
	pulumi.RegisterInputType(reflect.TypeOf((*AuthBackendIdentityWhitelistArrayInput)(nil)).Elem(), AuthBackendIdentityWhitelistArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AuthBackendIdentityWhitelistMapInput)(nil)).Elem(), AuthBackendIdentityWhitelistMap{})
	pulumi.RegisterOutputType(AuthBackendIdentityWhitelistOutput{})
	pulumi.RegisterOutputType(AuthBackendIdentityWhitelistArrayOutput{})
	pulumi.RegisterOutputType(AuthBackendIdentityWhitelistMapOutput{})
}
