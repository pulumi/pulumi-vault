// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package okta

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a resource for managing an
// [Okta auth backend within Vault](https://www.vaultproject.io/docs/auth/okta.html).
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-vault/sdk/v5/go/vault/okta"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := okta.NewAuthBackend(ctx, "example", &okta.AuthBackendArgs{
// 			Description: pulumi.String("Demonstration of the Terraform Okta auth backend"),
// 			Groups: okta.AuthBackendGroupArray{
// 				&okta.AuthBackendGroupArgs{
// 					GroupName: pulumi.String("foo"),
// 					Policies: pulumi.StringArray{
// 						pulumi.String("one"),
// 						pulumi.String("two"),
// 					},
// 				},
// 			},
// 			Organization: pulumi.String("example"),
// 			Token:        pulumi.String("something that should be kept secret"),
// 			Users: okta.AuthBackendUserArray{
// 				&okta.AuthBackendUserArgs{
// 					Groups: pulumi.StringArray{
// 						pulumi.String("foo"),
// 					},
// 					Username: pulumi.String("bar"),
// 				},
// 			},
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
// Okta authentication backends can be imported using its `path`, e.g.
//
// ```sh
//  $ pulumi import vault:okta/authBackend:AuthBackend example okta
// ```
type AuthBackend struct {
	pulumi.CustomResourceState

	// The mount accessor related to the auth mount. It is useful for integration with [Identity Secrets Engine](https://www.vaultproject.io/docs/secrets/identity/index.html).
	Accessor pulumi.StringOutput `pulumi:"accessor"`
	// The Okta url. Examples: oktapreview.com, okta.com
	BaseUrl pulumi.StringPtrOutput `pulumi:"baseUrl"`
	// When true, requests by Okta for a MFA check will be bypassed. This also disallows certain status checks on the account, such as whether the password is expired.
	BypassOktaMfa pulumi.BoolPtrOutput `pulumi:"bypassOktaMfa"`
	// The description of the auth backend
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Associate Okta groups with policies within Vault.
	// See below for more details.
	Groups AuthBackendGroupTypeArrayOutput `pulumi:"groups"`
	// Maximum duration after which authentication will be expired
	// [See the documentation for info on valid duration formats](https://golang.org/pkg/time/#ParseDuration).
	MaxTtl pulumi.StringPtrOutput `pulumi:"maxTtl"`
	// The Okta organization. This will be the first part of the url `https://XXX.okta.com`
	Organization pulumi.StringOutput `pulumi:"organization"`
	// Path to mount the Okta auth backend
	Path pulumi.StringPtrOutput `pulumi:"path"`
	// The Okta API token. This is required to query Okta for user group membership.
	// If this is not supplied only locally configured groups will be enabled.
	Token pulumi.StringPtrOutput `pulumi:"token"`
	// Duration after which authentication will be expired.
	// [See the documentation for info on valid duration formats](https://golang.org/pkg/time/#ParseDuration).
	Ttl pulumi.StringPtrOutput `pulumi:"ttl"`
	// Associate Okta users with groups or policies within Vault.
	// See below for more details.
	Users AuthBackendUserTypeArrayOutput `pulumi:"users"`
}

// NewAuthBackend registers a new resource with the given unique name, arguments, and options.
func NewAuthBackend(ctx *pulumi.Context,
	name string, args *AuthBackendArgs, opts ...pulumi.ResourceOption) (*AuthBackend, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Organization == nil {
		return nil, errors.New("invalid value for required argument 'Organization'")
	}
	var resource AuthBackend
	err := ctx.RegisterResource("vault:okta/authBackend:AuthBackend", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAuthBackend gets an existing AuthBackend resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAuthBackend(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AuthBackendState, opts ...pulumi.ResourceOption) (*AuthBackend, error) {
	var resource AuthBackend
	err := ctx.ReadResource("vault:okta/authBackend:AuthBackend", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AuthBackend resources.
type authBackendState struct {
	// The mount accessor related to the auth mount. It is useful for integration with [Identity Secrets Engine](https://www.vaultproject.io/docs/secrets/identity/index.html).
	Accessor *string `pulumi:"accessor"`
	// The Okta url. Examples: oktapreview.com, okta.com
	BaseUrl *string `pulumi:"baseUrl"`
	// When true, requests by Okta for a MFA check will be bypassed. This also disallows certain status checks on the account, such as whether the password is expired.
	BypassOktaMfa *bool `pulumi:"bypassOktaMfa"`
	// The description of the auth backend
	Description *string `pulumi:"description"`
	// Associate Okta groups with policies within Vault.
	// See below for more details.
	Groups []AuthBackendGroupType `pulumi:"groups"`
	// Maximum duration after which authentication will be expired
	// [See the documentation for info on valid duration formats](https://golang.org/pkg/time/#ParseDuration).
	MaxTtl *string `pulumi:"maxTtl"`
	// The Okta organization. This will be the first part of the url `https://XXX.okta.com`
	Organization *string `pulumi:"organization"`
	// Path to mount the Okta auth backend
	Path *string `pulumi:"path"`
	// The Okta API token. This is required to query Okta for user group membership.
	// If this is not supplied only locally configured groups will be enabled.
	Token *string `pulumi:"token"`
	// Duration after which authentication will be expired.
	// [See the documentation for info on valid duration formats](https://golang.org/pkg/time/#ParseDuration).
	Ttl *string `pulumi:"ttl"`
	// Associate Okta users with groups or policies within Vault.
	// See below for more details.
	Users []AuthBackendUserType `pulumi:"users"`
}

type AuthBackendState struct {
	// The mount accessor related to the auth mount. It is useful for integration with [Identity Secrets Engine](https://www.vaultproject.io/docs/secrets/identity/index.html).
	Accessor pulumi.StringPtrInput
	// The Okta url. Examples: oktapreview.com, okta.com
	BaseUrl pulumi.StringPtrInput
	// When true, requests by Okta for a MFA check will be bypassed. This also disallows certain status checks on the account, such as whether the password is expired.
	BypassOktaMfa pulumi.BoolPtrInput
	// The description of the auth backend
	Description pulumi.StringPtrInput
	// Associate Okta groups with policies within Vault.
	// See below for more details.
	Groups AuthBackendGroupTypeArrayInput
	// Maximum duration after which authentication will be expired
	// [See the documentation for info on valid duration formats](https://golang.org/pkg/time/#ParseDuration).
	MaxTtl pulumi.StringPtrInput
	// The Okta organization. This will be the first part of the url `https://XXX.okta.com`
	Organization pulumi.StringPtrInput
	// Path to mount the Okta auth backend
	Path pulumi.StringPtrInput
	// The Okta API token. This is required to query Okta for user group membership.
	// If this is not supplied only locally configured groups will be enabled.
	Token pulumi.StringPtrInput
	// Duration after which authentication will be expired.
	// [See the documentation for info on valid duration formats](https://golang.org/pkg/time/#ParseDuration).
	Ttl pulumi.StringPtrInput
	// Associate Okta users with groups or policies within Vault.
	// See below for more details.
	Users AuthBackendUserTypeArrayInput
}

func (AuthBackendState) ElementType() reflect.Type {
	return reflect.TypeOf((*authBackendState)(nil)).Elem()
}

type authBackendArgs struct {
	// The Okta url. Examples: oktapreview.com, okta.com
	BaseUrl *string `pulumi:"baseUrl"`
	// When true, requests by Okta for a MFA check will be bypassed. This also disallows certain status checks on the account, such as whether the password is expired.
	BypassOktaMfa *bool `pulumi:"bypassOktaMfa"`
	// The description of the auth backend
	Description *string `pulumi:"description"`
	// Associate Okta groups with policies within Vault.
	// See below for more details.
	Groups []AuthBackendGroupType `pulumi:"groups"`
	// Maximum duration after which authentication will be expired
	// [See the documentation for info on valid duration formats](https://golang.org/pkg/time/#ParseDuration).
	MaxTtl *string `pulumi:"maxTtl"`
	// The Okta organization. This will be the first part of the url `https://XXX.okta.com`
	Organization string `pulumi:"organization"`
	// Path to mount the Okta auth backend
	Path *string `pulumi:"path"`
	// The Okta API token. This is required to query Okta for user group membership.
	// If this is not supplied only locally configured groups will be enabled.
	Token *string `pulumi:"token"`
	// Duration after which authentication will be expired.
	// [See the documentation for info on valid duration formats](https://golang.org/pkg/time/#ParseDuration).
	Ttl *string `pulumi:"ttl"`
	// Associate Okta users with groups or policies within Vault.
	// See below for more details.
	Users []AuthBackendUserType `pulumi:"users"`
}

// The set of arguments for constructing a AuthBackend resource.
type AuthBackendArgs struct {
	// The Okta url. Examples: oktapreview.com, okta.com
	BaseUrl pulumi.StringPtrInput
	// When true, requests by Okta for a MFA check will be bypassed. This also disallows certain status checks on the account, such as whether the password is expired.
	BypassOktaMfa pulumi.BoolPtrInput
	// The description of the auth backend
	Description pulumi.StringPtrInput
	// Associate Okta groups with policies within Vault.
	// See below for more details.
	Groups AuthBackendGroupTypeArrayInput
	// Maximum duration after which authentication will be expired
	// [See the documentation for info on valid duration formats](https://golang.org/pkg/time/#ParseDuration).
	MaxTtl pulumi.StringPtrInput
	// The Okta organization. This will be the first part of the url `https://XXX.okta.com`
	Organization pulumi.StringInput
	// Path to mount the Okta auth backend
	Path pulumi.StringPtrInput
	// The Okta API token. This is required to query Okta for user group membership.
	// If this is not supplied only locally configured groups will be enabled.
	Token pulumi.StringPtrInput
	// Duration after which authentication will be expired.
	// [See the documentation for info on valid duration formats](https://golang.org/pkg/time/#ParseDuration).
	Ttl pulumi.StringPtrInput
	// Associate Okta users with groups or policies within Vault.
	// See below for more details.
	Users AuthBackendUserTypeArrayInput
}

func (AuthBackendArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*authBackendArgs)(nil)).Elem()
}

type AuthBackendInput interface {
	pulumi.Input

	ToAuthBackendOutput() AuthBackendOutput
	ToAuthBackendOutputWithContext(ctx context.Context) AuthBackendOutput
}

func (*AuthBackend) ElementType() reflect.Type {
	return reflect.TypeOf((**AuthBackend)(nil)).Elem()
}

func (i *AuthBackend) ToAuthBackendOutput() AuthBackendOutput {
	return i.ToAuthBackendOutputWithContext(context.Background())
}

func (i *AuthBackend) ToAuthBackendOutputWithContext(ctx context.Context) AuthBackendOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AuthBackendOutput)
}

// AuthBackendArrayInput is an input type that accepts AuthBackendArray and AuthBackendArrayOutput values.
// You can construct a concrete instance of `AuthBackendArrayInput` via:
//
//          AuthBackendArray{ AuthBackendArgs{...} }
type AuthBackendArrayInput interface {
	pulumi.Input

	ToAuthBackendArrayOutput() AuthBackendArrayOutput
	ToAuthBackendArrayOutputWithContext(context.Context) AuthBackendArrayOutput
}

type AuthBackendArray []AuthBackendInput

func (AuthBackendArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AuthBackend)(nil)).Elem()
}

func (i AuthBackendArray) ToAuthBackendArrayOutput() AuthBackendArrayOutput {
	return i.ToAuthBackendArrayOutputWithContext(context.Background())
}

func (i AuthBackendArray) ToAuthBackendArrayOutputWithContext(ctx context.Context) AuthBackendArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AuthBackendArrayOutput)
}

// AuthBackendMapInput is an input type that accepts AuthBackendMap and AuthBackendMapOutput values.
// You can construct a concrete instance of `AuthBackendMapInput` via:
//
//          AuthBackendMap{ "key": AuthBackendArgs{...} }
type AuthBackendMapInput interface {
	pulumi.Input

	ToAuthBackendMapOutput() AuthBackendMapOutput
	ToAuthBackendMapOutputWithContext(context.Context) AuthBackendMapOutput
}

type AuthBackendMap map[string]AuthBackendInput

func (AuthBackendMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AuthBackend)(nil)).Elem()
}

func (i AuthBackendMap) ToAuthBackendMapOutput() AuthBackendMapOutput {
	return i.ToAuthBackendMapOutputWithContext(context.Background())
}

func (i AuthBackendMap) ToAuthBackendMapOutputWithContext(ctx context.Context) AuthBackendMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AuthBackendMapOutput)
}

type AuthBackendOutput struct{ *pulumi.OutputState }

func (AuthBackendOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AuthBackend)(nil)).Elem()
}

func (o AuthBackendOutput) ToAuthBackendOutput() AuthBackendOutput {
	return o
}

func (o AuthBackendOutput) ToAuthBackendOutputWithContext(ctx context.Context) AuthBackendOutput {
	return o
}

type AuthBackendArrayOutput struct{ *pulumi.OutputState }

func (AuthBackendArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AuthBackend)(nil)).Elem()
}

func (o AuthBackendArrayOutput) ToAuthBackendArrayOutput() AuthBackendArrayOutput {
	return o
}

func (o AuthBackendArrayOutput) ToAuthBackendArrayOutputWithContext(ctx context.Context) AuthBackendArrayOutput {
	return o
}

func (o AuthBackendArrayOutput) Index(i pulumi.IntInput) AuthBackendOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *AuthBackend {
		return vs[0].([]*AuthBackend)[vs[1].(int)]
	}).(AuthBackendOutput)
}

type AuthBackendMapOutput struct{ *pulumi.OutputState }

func (AuthBackendMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AuthBackend)(nil)).Elem()
}

func (o AuthBackendMapOutput) ToAuthBackendMapOutput() AuthBackendMapOutput {
	return o
}

func (o AuthBackendMapOutput) ToAuthBackendMapOutputWithContext(ctx context.Context) AuthBackendMapOutput {
	return o
}

func (o AuthBackendMapOutput) MapIndex(k pulumi.StringInput) AuthBackendOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *AuthBackend {
		return vs[0].(map[string]*AuthBackend)[vs[1].(string)]
	}).(AuthBackendOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AuthBackendInput)(nil)).Elem(), &AuthBackend{})
	pulumi.RegisterInputType(reflect.TypeOf((*AuthBackendArrayInput)(nil)).Elem(), AuthBackendArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AuthBackendMapInput)(nil)).Elem(), AuthBackendMap{})
	pulumi.RegisterOutputType(AuthBackendOutput{})
	pulumi.RegisterOutputType(AuthBackendArrayOutput{})
	pulumi.RegisterOutputType(AuthBackendMapOutput{})
}
