// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package github

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a GitHub Auth mount in a Vault server. See the [Vault
// documentation](https://www.vaultproject.io/docs/auth/github/) for more
// information.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-vault/sdk/v5/go/vault/github"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := github.NewAuthBackend(ctx, "example", &github.AuthBackendArgs{
// 			Organization: pulumi.String("myorg"),
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
// GitHub authentication mounts can be imported using the `path`, e.g.
//
// ```sh
//  $ pulumi import vault:github/authBackend:AuthBackend example github
// ```
type AuthBackend struct {
	pulumi.CustomResourceState

	// The mount accessor related to the auth mount. It is useful for integration with [Identity Secrets Engine](https://www.vaultproject.io/docs/secrets/identity/index.html).
	Accessor pulumi.StringOutput `pulumi:"accessor"`
	// The API endpoint to use. Useful if you
	// are running GitHub Enterprise or an API-compatible authentication server.
	BaseUrl pulumi.StringPtrOutput `pulumi:"baseUrl"`
	// Specifies the description of the mount.
	// This overrides the current stored value, if any.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The organization configured users must be part of.
	Organization pulumi.StringOutput `pulumi:"organization"`
	// The ID of the organization users must be part of.
	// Vault will attempt to fetch and set this value if it is not provided. (Vault 1.10+)
	OrganizationId pulumi.IntOutput `pulumi:"organizationId"`
	// Path where the auth backend is mounted. Defaults to `auth/github`
	// if not specified.
	Path pulumi.StringPtrOutput `pulumi:"path"`
	// (Optional) List of CIDR blocks; if set, specifies blocks of IP
	// addresses which can authenticate successfully, and ties the resulting token to these blocks
	// as well.
	TokenBoundCidrs pulumi.StringArrayOutput `pulumi:"tokenBoundCidrs"`
	// (Optional) If set, will encode an
	// [explicit max TTL](https://www.vaultproject.io/docs/concepts/tokens.html#token-time-to-live-periodic-tokens-and-explicit-max-ttls)
	// onto the token in number of seconds. This is a hard cap even if `tokenTtl` and
	// `tokenMaxTtl` would otherwise allow a renewal.
	TokenExplicitMaxTtl pulumi.IntPtrOutput `pulumi:"tokenExplicitMaxTtl"`
	// (Optional) The maximum lifetime for generated tokens in number of seconds.
	// Its current value will be referenced at renewal time.
	TokenMaxTtl pulumi.IntPtrOutput `pulumi:"tokenMaxTtl"`
	// (Optional) If set, the default policy will not be set on
	// generated tokens; otherwise it will be added to the policies set in token_policies.
	TokenNoDefaultPolicy pulumi.BoolPtrOutput `pulumi:"tokenNoDefaultPolicy"`
	// (Optional) The
	// [period](https://www.vaultproject.io/docs/concepts/tokens.html#token-time-to-live-periodic-tokens-and-explicit-max-ttls),
	// if any, in number of seconds to set on the token.
	TokenNumUses pulumi.IntPtrOutput `pulumi:"tokenNumUses"`
	// (Optional) If set, indicates that the
	// token generated using this role should never expire. The token should be renewed within the
	// duration specified by this value. At each renewal, the token's TTL will be set to the
	// value of this field. Specified in seconds.
	TokenPeriod pulumi.IntPtrOutput `pulumi:"tokenPeriod"`
	// (Optional) List of policies to encode onto generated tokens. Depending
	// on the auth method, this list may be supplemented by user/group/other values.
	TokenPolicies pulumi.StringArrayOutput `pulumi:"tokenPolicies"`
	// (Optional) The incremental lifetime for generated tokens in number of seconds.
	// Its current value will be referenced at renewal time.
	TokenTtl pulumi.IntPtrOutput `pulumi:"tokenTtl"`
	// Specifies the type of tokens that should be returned by
	// the mount. Valid values are "default-service", "default-batch", "service", "batch".
	TokenType pulumi.StringPtrOutput `pulumi:"tokenType"`
	// Extra configuration block. Structure is documented below.
	Tune AuthBackendTuneOutput `pulumi:"tune"`
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
	err := ctx.RegisterResource("vault:github/authBackend:AuthBackend", name, args, &resource, opts...)
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
	err := ctx.ReadResource("vault:github/authBackend:AuthBackend", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AuthBackend resources.
type authBackendState struct {
	// The mount accessor related to the auth mount. It is useful for integration with [Identity Secrets Engine](https://www.vaultproject.io/docs/secrets/identity/index.html).
	Accessor *string `pulumi:"accessor"`
	// The API endpoint to use. Useful if you
	// are running GitHub Enterprise or an API-compatible authentication server.
	BaseUrl *string `pulumi:"baseUrl"`
	// Specifies the description of the mount.
	// This overrides the current stored value, if any.
	Description *string `pulumi:"description"`
	// The organization configured users must be part of.
	Organization *string `pulumi:"organization"`
	// The ID of the organization users must be part of.
	// Vault will attempt to fetch and set this value if it is not provided. (Vault 1.10+)
	OrganizationId *int `pulumi:"organizationId"`
	// Path where the auth backend is mounted. Defaults to `auth/github`
	// if not specified.
	Path *string `pulumi:"path"`
	// (Optional) List of CIDR blocks; if set, specifies blocks of IP
	// addresses which can authenticate successfully, and ties the resulting token to these blocks
	// as well.
	TokenBoundCidrs []string `pulumi:"tokenBoundCidrs"`
	// (Optional) If set, will encode an
	// [explicit max TTL](https://www.vaultproject.io/docs/concepts/tokens.html#token-time-to-live-periodic-tokens-and-explicit-max-ttls)
	// onto the token in number of seconds. This is a hard cap even if `tokenTtl` and
	// `tokenMaxTtl` would otherwise allow a renewal.
	TokenExplicitMaxTtl *int `pulumi:"tokenExplicitMaxTtl"`
	// (Optional) The maximum lifetime for generated tokens in number of seconds.
	// Its current value will be referenced at renewal time.
	TokenMaxTtl *int `pulumi:"tokenMaxTtl"`
	// (Optional) If set, the default policy will not be set on
	// generated tokens; otherwise it will be added to the policies set in token_policies.
	TokenNoDefaultPolicy *bool `pulumi:"tokenNoDefaultPolicy"`
	// (Optional) The
	// [period](https://www.vaultproject.io/docs/concepts/tokens.html#token-time-to-live-periodic-tokens-and-explicit-max-ttls),
	// if any, in number of seconds to set on the token.
	TokenNumUses *int `pulumi:"tokenNumUses"`
	// (Optional) If set, indicates that the
	// token generated using this role should never expire. The token should be renewed within the
	// duration specified by this value. At each renewal, the token's TTL will be set to the
	// value of this field. Specified in seconds.
	TokenPeriod *int `pulumi:"tokenPeriod"`
	// (Optional) List of policies to encode onto generated tokens. Depending
	// on the auth method, this list may be supplemented by user/group/other values.
	TokenPolicies []string `pulumi:"tokenPolicies"`
	// (Optional) The incremental lifetime for generated tokens in number of seconds.
	// Its current value will be referenced at renewal time.
	TokenTtl *int `pulumi:"tokenTtl"`
	// Specifies the type of tokens that should be returned by
	// the mount. Valid values are "default-service", "default-batch", "service", "batch".
	TokenType *string `pulumi:"tokenType"`
	// Extra configuration block. Structure is documented below.
	Tune *AuthBackendTune `pulumi:"tune"`
}

type AuthBackendState struct {
	// The mount accessor related to the auth mount. It is useful for integration with [Identity Secrets Engine](https://www.vaultproject.io/docs/secrets/identity/index.html).
	Accessor pulumi.StringPtrInput
	// The API endpoint to use. Useful if you
	// are running GitHub Enterprise or an API-compatible authentication server.
	BaseUrl pulumi.StringPtrInput
	// Specifies the description of the mount.
	// This overrides the current stored value, if any.
	Description pulumi.StringPtrInput
	// The organization configured users must be part of.
	Organization pulumi.StringPtrInput
	// The ID of the organization users must be part of.
	// Vault will attempt to fetch and set this value if it is not provided. (Vault 1.10+)
	OrganizationId pulumi.IntPtrInput
	// Path where the auth backend is mounted. Defaults to `auth/github`
	// if not specified.
	Path pulumi.StringPtrInput
	// (Optional) List of CIDR blocks; if set, specifies blocks of IP
	// addresses which can authenticate successfully, and ties the resulting token to these blocks
	// as well.
	TokenBoundCidrs pulumi.StringArrayInput
	// (Optional) If set, will encode an
	// [explicit max TTL](https://www.vaultproject.io/docs/concepts/tokens.html#token-time-to-live-periodic-tokens-and-explicit-max-ttls)
	// onto the token in number of seconds. This is a hard cap even if `tokenTtl` and
	// `tokenMaxTtl` would otherwise allow a renewal.
	TokenExplicitMaxTtl pulumi.IntPtrInput
	// (Optional) The maximum lifetime for generated tokens in number of seconds.
	// Its current value will be referenced at renewal time.
	TokenMaxTtl pulumi.IntPtrInput
	// (Optional) If set, the default policy will not be set on
	// generated tokens; otherwise it will be added to the policies set in token_policies.
	TokenNoDefaultPolicy pulumi.BoolPtrInput
	// (Optional) The
	// [period](https://www.vaultproject.io/docs/concepts/tokens.html#token-time-to-live-periodic-tokens-and-explicit-max-ttls),
	// if any, in number of seconds to set on the token.
	TokenNumUses pulumi.IntPtrInput
	// (Optional) If set, indicates that the
	// token generated using this role should never expire. The token should be renewed within the
	// duration specified by this value. At each renewal, the token's TTL will be set to the
	// value of this field. Specified in seconds.
	TokenPeriod pulumi.IntPtrInput
	// (Optional) List of policies to encode onto generated tokens. Depending
	// on the auth method, this list may be supplemented by user/group/other values.
	TokenPolicies pulumi.StringArrayInput
	// (Optional) The incremental lifetime for generated tokens in number of seconds.
	// Its current value will be referenced at renewal time.
	TokenTtl pulumi.IntPtrInput
	// Specifies the type of tokens that should be returned by
	// the mount. Valid values are "default-service", "default-batch", "service", "batch".
	TokenType pulumi.StringPtrInput
	// Extra configuration block. Structure is documented below.
	Tune AuthBackendTunePtrInput
}

func (AuthBackendState) ElementType() reflect.Type {
	return reflect.TypeOf((*authBackendState)(nil)).Elem()
}

type authBackendArgs struct {
	// The API endpoint to use. Useful if you
	// are running GitHub Enterprise or an API-compatible authentication server.
	BaseUrl *string `pulumi:"baseUrl"`
	// Specifies the description of the mount.
	// This overrides the current stored value, if any.
	Description *string `pulumi:"description"`
	// The organization configured users must be part of.
	Organization string `pulumi:"organization"`
	// The ID of the organization users must be part of.
	// Vault will attempt to fetch and set this value if it is not provided. (Vault 1.10+)
	OrganizationId *int `pulumi:"organizationId"`
	// Path where the auth backend is mounted. Defaults to `auth/github`
	// if not specified.
	Path *string `pulumi:"path"`
	// (Optional) List of CIDR blocks; if set, specifies blocks of IP
	// addresses which can authenticate successfully, and ties the resulting token to these blocks
	// as well.
	TokenBoundCidrs []string `pulumi:"tokenBoundCidrs"`
	// (Optional) If set, will encode an
	// [explicit max TTL](https://www.vaultproject.io/docs/concepts/tokens.html#token-time-to-live-periodic-tokens-and-explicit-max-ttls)
	// onto the token in number of seconds. This is a hard cap even if `tokenTtl` and
	// `tokenMaxTtl` would otherwise allow a renewal.
	TokenExplicitMaxTtl *int `pulumi:"tokenExplicitMaxTtl"`
	// (Optional) The maximum lifetime for generated tokens in number of seconds.
	// Its current value will be referenced at renewal time.
	TokenMaxTtl *int `pulumi:"tokenMaxTtl"`
	// (Optional) If set, the default policy will not be set on
	// generated tokens; otherwise it will be added to the policies set in token_policies.
	TokenNoDefaultPolicy *bool `pulumi:"tokenNoDefaultPolicy"`
	// (Optional) The
	// [period](https://www.vaultproject.io/docs/concepts/tokens.html#token-time-to-live-periodic-tokens-and-explicit-max-ttls),
	// if any, in number of seconds to set on the token.
	TokenNumUses *int `pulumi:"tokenNumUses"`
	// (Optional) If set, indicates that the
	// token generated using this role should never expire. The token should be renewed within the
	// duration specified by this value. At each renewal, the token's TTL will be set to the
	// value of this field. Specified in seconds.
	TokenPeriod *int `pulumi:"tokenPeriod"`
	// (Optional) List of policies to encode onto generated tokens. Depending
	// on the auth method, this list may be supplemented by user/group/other values.
	TokenPolicies []string `pulumi:"tokenPolicies"`
	// (Optional) The incremental lifetime for generated tokens in number of seconds.
	// Its current value will be referenced at renewal time.
	TokenTtl *int `pulumi:"tokenTtl"`
	// Specifies the type of tokens that should be returned by
	// the mount. Valid values are "default-service", "default-batch", "service", "batch".
	TokenType *string `pulumi:"tokenType"`
	// Extra configuration block. Structure is documented below.
	Tune *AuthBackendTune `pulumi:"tune"`
}

// The set of arguments for constructing a AuthBackend resource.
type AuthBackendArgs struct {
	// The API endpoint to use. Useful if you
	// are running GitHub Enterprise or an API-compatible authentication server.
	BaseUrl pulumi.StringPtrInput
	// Specifies the description of the mount.
	// This overrides the current stored value, if any.
	Description pulumi.StringPtrInput
	// The organization configured users must be part of.
	Organization pulumi.StringInput
	// The ID of the organization users must be part of.
	// Vault will attempt to fetch and set this value if it is not provided. (Vault 1.10+)
	OrganizationId pulumi.IntPtrInput
	// Path where the auth backend is mounted. Defaults to `auth/github`
	// if not specified.
	Path pulumi.StringPtrInput
	// (Optional) List of CIDR blocks; if set, specifies blocks of IP
	// addresses which can authenticate successfully, and ties the resulting token to these blocks
	// as well.
	TokenBoundCidrs pulumi.StringArrayInput
	// (Optional) If set, will encode an
	// [explicit max TTL](https://www.vaultproject.io/docs/concepts/tokens.html#token-time-to-live-periodic-tokens-and-explicit-max-ttls)
	// onto the token in number of seconds. This is a hard cap even if `tokenTtl` and
	// `tokenMaxTtl` would otherwise allow a renewal.
	TokenExplicitMaxTtl pulumi.IntPtrInput
	// (Optional) The maximum lifetime for generated tokens in number of seconds.
	// Its current value will be referenced at renewal time.
	TokenMaxTtl pulumi.IntPtrInput
	// (Optional) If set, the default policy will not be set on
	// generated tokens; otherwise it will be added to the policies set in token_policies.
	TokenNoDefaultPolicy pulumi.BoolPtrInput
	// (Optional) The
	// [period](https://www.vaultproject.io/docs/concepts/tokens.html#token-time-to-live-periodic-tokens-and-explicit-max-ttls),
	// if any, in number of seconds to set on the token.
	TokenNumUses pulumi.IntPtrInput
	// (Optional) If set, indicates that the
	// token generated using this role should never expire. The token should be renewed within the
	// duration specified by this value. At each renewal, the token's TTL will be set to the
	// value of this field. Specified in seconds.
	TokenPeriod pulumi.IntPtrInput
	// (Optional) List of policies to encode onto generated tokens. Depending
	// on the auth method, this list may be supplemented by user/group/other values.
	TokenPolicies pulumi.StringArrayInput
	// (Optional) The incremental lifetime for generated tokens in number of seconds.
	// Its current value will be referenced at renewal time.
	TokenTtl pulumi.IntPtrInput
	// Specifies the type of tokens that should be returned by
	// the mount. Valid values are "default-service", "default-batch", "service", "batch".
	TokenType pulumi.StringPtrInput
	// Extra configuration block. Structure is documented below.
	Tune AuthBackendTunePtrInput
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
