// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package azure

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-vault/sdk/v5/go/vault/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages an Azure auth backend role in a Vault server. Roles constrain the
// instances or principals that can perform the login operation against the
// backend. See the [Vault
// documentation](https://www.vaultproject.io/docs/auth/azure.html) for more
// information.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-vault/sdk/v5/go/vault"
//	"github.com/pulumi/pulumi-vault/sdk/v5/go/vault/azure"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			azure, err := vault.NewAuthBackend(ctx, "azure", &vault.AuthBackendArgs{
//				Type: pulumi.String("azure"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = azure.NewAuthBackendRole(ctx, "example", &azure.AuthBackendRoleArgs{
//				Backend: azure.Path,
//				Role:    pulumi.String("test-role"),
//				BoundSubscriptionIds: pulumi.StringArray{
//					pulumi.String("11111111-2222-3333-4444-555555555555"),
//				},
//				BoundResourceGroups: pulumi.StringArray{
//					pulumi.String("123456789012"),
//				},
//				TokenTtl:    pulumi.Int(60),
//				TokenMaxTtl: pulumi.Int(120),
//				TokenPolicies: pulumi.StringArray{
//					pulumi.String("default"),
//					pulumi.String("dev"),
//					pulumi.String("prod"),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// ## Import
//
// Azure auth backend roles can be imported using `auth/`, the `backend` path, `/role/`, and the `role` name e.g.
//
// ```sh
//
//	$ pulumi import vault:azure/authBackendRole:AuthBackendRole example auth/azure/role/test-role
//
// ```
type AuthBackendRole struct {
	pulumi.CustomResourceState

	// Unique name of the auth backend to configure.
	Backend pulumi.StringPtrOutput `pulumi:"backend"`
	// If set, defines a constraint on the groups
	// that can perform the login operation that they should be using the group
	// ID specified by this field.
	BoundGroupIds pulumi.StringArrayOutput `pulumi:"boundGroupIds"`
	// If set, defines a constraint on the virtual machines
	// that can perform the login operation that the location in their identity
	// document must match the one specified by this field.
	BoundLocations pulumi.StringArrayOutput `pulumi:"boundLocations"`
	// If set, defines a constraint on the virtual
	// machines that can perform the login operation that they be associated with
	// the resource group that matches the value specified by this field.
	BoundResourceGroups pulumi.StringArrayOutput `pulumi:"boundResourceGroups"`
	// If set, defines a constraint on the virtual
	// machines that can perform the login operation that they must match the scale set
	// specified by this field.
	BoundScaleSets pulumi.StringArrayOutput `pulumi:"boundScaleSets"`
	// If set, defines a constraint on the
	// service principals that can perform the login operation that they should be possess
	// the ids specified by this field.
	BoundServicePrincipalIds pulumi.StringArrayOutput `pulumi:"boundServicePrincipalIds"`
	// If set, defines a constraint on the subscriptions
	// that can perform the login operation to ones which  matches the value specified by this
	// field.
	BoundSubscriptionIds pulumi.StringArrayOutput `pulumi:"boundSubscriptionIds"`
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
	// *Available only for Vault Enterprise*.
	Namespace pulumi.StringPtrOutput `pulumi:"namespace"`
	// The name of the role.
	Role pulumi.StringOutput `pulumi:"role"`
	// List of CIDR blocks; if set, specifies blocks of IP
	// addresses which can authenticate successfully, and ties the resulting token to these blocks
	// as well.
	TokenBoundCidrs pulumi.StringArrayOutput `pulumi:"tokenBoundCidrs"`
	// If set, will encode an
	// [explicit max TTL](https://www.vaultproject.io/docs/concepts/tokens.html#token-time-to-live-periodic-tokens-and-explicit-max-ttls)
	// onto the token in number of seconds. This is a hard cap even if `tokenTtl` and
	// `tokenMaxTtl` would otherwise allow a renewal.
	TokenExplicitMaxTtl pulumi.IntPtrOutput `pulumi:"tokenExplicitMaxTtl"`
	// The maximum lifetime for generated tokens in number of seconds.
	// Its current value will be referenced at renewal time.
	TokenMaxTtl pulumi.IntPtrOutput `pulumi:"tokenMaxTtl"`
	// If set, the default policy will not be set on
	// generated tokens; otherwise it will be added to the policies set in token_policies.
	TokenNoDefaultPolicy pulumi.BoolPtrOutput `pulumi:"tokenNoDefaultPolicy"`
	// The [maximum number](https://www.vaultproject.io/api-docs/azure#token_num_uses)
	// of times a generated token may be used (within its lifetime); 0 means unlimited.
	TokenNumUses pulumi.IntPtrOutput `pulumi:"tokenNumUses"`
	// If set, indicates that the
	// token generated using this role should never expire. The token should be renewed within the
	// duration specified by this value. At each renewal, the token's TTL will be set to the
	// value of this field. Specified in seconds.
	TokenPeriod pulumi.IntPtrOutput `pulumi:"tokenPeriod"`
	// List of policies to encode onto generated tokens. Depending
	// on the auth method, this list may be supplemented by user/group/other values.
	TokenPolicies pulumi.StringArrayOutput `pulumi:"tokenPolicies"`
	// The incremental lifetime for generated tokens in number of seconds.
	// Its current value will be referenced at renewal time.
	TokenTtl pulumi.IntPtrOutput `pulumi:"tokenTtl"`
	// The type of token that should be generated. Can be `service`,
	// `batch`, or `default` to use the mount's tuned default (which unless changed will be
	// `service` tokens). For token store roles, there are two additional possibilities:
	// `default-service` and `default-batch` which specify the type to return unless the client
	// requests a different type at generation time.
	TokenType pulumi.StringPtrOutput `pulumi:"tokenType"`
}

// NewAuthBackendRole registers a new resource with the given unique name, arguments, and options.
func NewAuthBackendRole(ctx *pulumi.Context,
	name string, args *AuthBackendRoleArgs, opts ...pulumi.ResourceOption) (*AuthBackendRole, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Role == nil {
		return nil, errors.New("invalid value for required argument 'Role'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource AuthBackendRole
	err := ctx.RegisterResource("vault:azure/authBackendRole:AuthBackendRole", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAuthBackendRole gets an existing AuthBackendRole resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAuthBackendRole(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AuthBackendRoleState, opts ...pulumi.ResourceOption) (*AuthBackendRole, error) {
	var resource AuthBackendRole
	err := ctx.ReadResource("vault:azure/authBackendRole:AuthBackendRole", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AuthBackendRole resources.
type authBackendRoleState struct {
	// Unique name of the auth backend to configure.
	Backend *string `pulumi:"backend"`
	// If set, defines a constraint on the groups
	// that can perform the login operation that they should be using the group
	// ID specified by this field.
	BoundGroupIds []string `pulumi:"boundGroupIds"`
	// If set, defines a constraint on the virtual machines
	// that can perform the login operation that the location in their identity
	// document must match the one specified by this field.
	BoundLocations []string `pulumi:"boundLocations"`
	// If set, defines a constraint on the virtual
	// machines that can perform the login operation that they be associated with
	// the resource group that matches the value specified by this field.
	BoundResourceGroups []string `pulumi:"boundResourceGroups"`
	// If set, defines a constraint on the virtual
	// machines that can perform the login operation that they must match the scale set
	// specified by this field.
	BoundScaleSets []string `pulumi:"boundScaleSets"`
	// If set, defines a constraint on the
	// service principals that can perform the login operation that they should be possess
	// the ids specified by this field.
	BoundServicePrincipalIds []string `pulumi:"boundServicePrincipalIds"`
	// If set, defines a constraint on the subscriptions
	// that can perform the login operation to ones which  matches the value specified by this
	// field.
	BoundSubscriptionIds []string `pulumi:"boundSubscriptionIds"`
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
	// *Available only for Vault Enterprise*.
	Namespace *string `pulumi:"namespace"`
	// The name of the role.
	Role *string `pulumi:"role"`
	// List of CIDR blocks; if set, specifies blocks of IP
	// addresses which can authenticate successfully, and ties the resulting token to these blocks
	// as well.
	TokenBoundCidrs []string `pulumi:"tokenBoundCidrs"`
	// If set, will encode an
	// [explicit max TTL](https://www.vaultproject.io/docs/concepts/tokens.html#token-time-to-live-periodic-tokens-and-explicit-max-ttls)
	// onto the token in number of seconds. This is a hard cap even if `tokenTtl` and
	// `tokenMaxTtl` would otherwise allow a renewal.
	TokenExplicitMaxTtl *int `pulumi:"tokenExplicitMaxTtl"`
	// The maximum lifetime for generated tokens in number of seconds.
	// Its current value will be referenced at renewal time.
	TokenMaxTtl *int `pulumi:"tokenMaxTtl"`
	// If set, the default policy will not be set on
	// generated tokens; otherwise it will be added to the policies set in token_policies.
	TokenNoDefaultPolicy *bool `pulumi:"tokenNoDefaultPolicy"`
	// The [maximum number](https://www.vaultproject.io/api-docs/azure#token_num_uses)
	// of times a generated token may be used (within its lifetime); 0 means unlimited.
	TokenNumUses *int `pulumi:"tokenNumUses"`
	// If set, indicates that the
	// token generated using this role should never expire. The token should be renewed within the
	// duration specified by this value. At each renewal, the token's TTL will be set to the
	// value of this field. Specified in seconds.
	TokenPeriod *int `pulumi:"tokenPeriod"`
	// List of policies to encode onto generated tokens. Depending
	// on the auth method, this list may be supplemented by user/group/other values.
	TokenPolicies []string `pulumi:"tokenPolicies"`
	// The incremental lifetime for generated tokens in number of seconds.
	// Its current value will be referenced at renewal time.
	TokenTtl *int `pulumi:"tokenTtl"`
	// The type of token that should be generated. Can be `service`,
	// `batch`, or `default` to use the mount's tuned default (which unless changed will be
	// `service` tokens). For token store roles, there are two additional possibilities:
	// `default-service` and `default-batch` which specify the type to return unless the client
	// requests a different type at generation time.
	TokenType *string `pulumi:"tokenType"`
}

type AuthBackendRoleState struct {
	// Unique name of the auth backend to configure.
	Backend pulumi.StringPtrInput
	// If set, defines a constraint on the groups
	// that can perform the login operation that they should be using the group
	// ID specified by this field.
	BoundGroupIds pulumi.StringArrayInput
	// If set, defines a constraint on the virtual machines
	// that can perform the login operation that the location in their identity
	// document must match the one specified by this field.
	BoundLocations pulumi.StringArrayInput
	// If set, defines a constraint on the virtual
	// machines that can perform the login operation that they be associated with
	// the resource group that matches the value specified by this field.
	BoundResourceGroups pulumi.StringArrayInput
	// If set, defines a constraint on the virtual
	// machines that can perform the login operation that they must match the scale set
	// specified by this field.
	BoundScaleSets pulumi.StringArrayInput
	// If set, defines a constraint on the
	// service principals that can perform the login operation that they should be possess
	// the ids specified by this field.
	BoundServicePrincipalIds pulumi.StringArrayInput
	// If set, defines a constraint on the subscriptions
	// that can perform the login operation to ones which  matches the value specified by this
	// field.
	BoundSubscriptionIds pulumi.StringArrayInput
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
	// *Available only for Vault Enterprise*.
	Namespace pulumi.StringPtrInput
	// The name of the role.
	Role pulumi.StringPtrInput
	// List of CIDR blocks; if set, specifies blocks of IP
	// addresses which can authenticate successfully, and ties the resulting token to these blocks
	// as well.
	TokenBoundCidrs pulumi.StringArrayInput
	// If set, will encode an
	// [explicit max TTL](https://www.vaultproject.io/docs/concepts/tokens.html#token-time-to-live-periodic-tokens-and-explicit-max-ttls)
	// onto the token in number of seconds. This is a hard cap even if `tokenTtl` and
	// `tokenMaxTtl` would otherwise allow a renewal.
	TokenExplicitMaxTtl pulumi.IntPtrInput
	// The maximum lifetime for generated tokens in number of seconds.
	// Its current value will be referenced at renewal time.
	TokenMaxTtl pulumi.IntPtrInput
	// If set, the default policy will not be set on
	// generated tokens; otherwise it will be added to the policies set in token_policies.
	TokenNoDefaultPolicy pulumi.BoolPtrInput
	// The [maximum number](https://www.vaultproject.io/api-docs/azure#token_num_uses)
	// of times a generated token may be used (within its lifetime); 0 means unlimited.
	TokenNumUses pulumi.IntPtrInput
	// If set, indicates that the
	// token generated using this role should never expire. The token should be renewed within the
	// duration specified by this value. At each renewal, the token's TTL will be set to the
	// value of this field. Specified in seconds.
	TokenPeriod pulumi.IntPtrInput
	// List of policies to encode onto generated tokens. Depending
	// on the auth method, this list may be supplemented by user/group/other values.
	TokenPolicies pulumi.StringArrayInput
	// The incremental lifetime for generated tokens in number of seconds.
	// Its current value will be referenced at renewal time.
	TokenTtl pulumi.IntPtrInput
	// The type of token that should be generated. Can be `service`,
	// `batch`, or `default` to use the mount's tuned default (which unless changed will be
	// `service` tokens). For token store roles, there are two additional possibilities:
	// `default-service` and `default-batch` which specify the type to return unless the client
	// requests a different type at generation time.
	TokenType pulumi.StringPtrInput
}

func (AuthBackendRoleState) ElementType() reflect.Type {
	return reflect.TypeOf((*authBackendRoleState)(nil)).Elem()
}

type authBackendRoleArgs struct {
	// Unique name of the auth backend to configure.
	Backend *string `pulumi:"backend"`
	// If set, defines a constraint on the groups
	// that can perform the login operation that they should be using the group
	// ID specified by this field.
	BoundGroupIds []string `pulumi:"boundGroupIds"`
	// If set, defines a constraint on the virtual machines
	// that can perform the login operation that the location in their identity
	// document must match the one specified by this field.
	BoundLocations []string `pulumi:"boundLocations"`
	// If set, defines a constraint on the virtual
	// machines that can perform the login operation that they be associated with
	// the resource group that matches the value specified by this field.
	BoundResourceGroups []string `pulumi:"boundResourceGroups"`
	// If set, defines a constraint on the virtual
	// machines that can perform the login operation that they must match the scale set
	// specified by this field.
	BoundScaleSets []string `pulumi:"boundScaleSets"`
	// If set, defines a constraint on the
	// service principals that can perform the login operation that they should be possess
	// the ids specified by this field.
	BoundServicePrincipalIds []string `pulumi:"boundServicePrincipalIds"`
	// If set, defines a constraint on the subscriptions
	// that can perform the login operation to ones which  matches the value specified by this
	// field.
	BoundSubscriptionIds []string `pulumi:"boundSubscriptionIds"`
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
	// *Available only for Vault Enterprise*.
	Namespace *string `pulumi:"namespace"`
	// The name of the role.
	Role string `pulumi:"role"`
	// List of CIDR blocks; if set, specifies blocks of IP
	// addresses which can authenticate successfully, and ties the resulting token to these blocks
	// as well.
	TokenBoundCidrs []string `pulumi:"tokenBoundCidrs"`
	// If set, will encode an
	// [explicit max TTL](https://www.vaultproject.io/docs/concepts/tokens.html#token-time-to-live-periodic-tokens-and-explicit-max-ttls)
	// onto the token in number of seconds. This is a hard cap even if `tokenTtl` and
	// `tokenMaxTtl` would otherwise allow a renewal.
	TokenExplicitMaxTtl *int `pulumi:"tokenExplicitMaxTtl"`
	// The maximum lifetime for generated tokens in number of seconds.
	// Its current value will be referenced at renewal time.
	TokenMaxTtl *int `pulumi:"tokenMaxTtl"`
	// If set, the default policy will not be set on
	// generated tokens; otherwise it will be added to the policies set in token_policies.
	TokenNoDefaultPolicy *bool `pulumi:"tokenNoDefaultPolicy"`
	// The [maximum number](https://www.vaultproject.io/api-docs/azure#token_num_uses)
	// of times a generated token may be used (within its lifetime); 0 means unlimited.
	TokenNumUses *int `pulumi:"tokenNumUses"`
	// If set, indicates that the
	// token generated using this role should never expire. The token should be renewed within the
	// duration specified by this value. At each renewal, the token's TTL will be set to the
	// value of this field. Specified in seconds.
	TokenPeriod *int `pulumi:"tokenPeriod"`
	// List of policies to encode onto generated tokens. Depending
	// on the auth method, this list may be supplemented by user/group/other values.
	TokenPolicies []string `pulumi:"tokenPolicies"`
	// The incremental lifetime for generated tokens in number of seconds.
	// Its current value will be referenced at renewal time.
	TokenTtl *int `pulumi:"tokenTtl"`
	// The type of token that should be generated. Can be `service`,
	// `batch`, or `default` to use the mount's tuned default (which unless changed will be
	// `service` tokens). For token store roles, there are two additional possibilities:
	// `default-service` and `default-batch` which specify the type to return unless the client
	// requests a different type at generation time.
	TokenType *string `pulumi:"tokenType"`
}

// The set of arguments for constructing a AuthBackendRole resource.
type AuthBackendRoleArgs struct {
	// Unique name of the auth backend to configure.
	Backend pulumi.StringPtrInput
	// If set, defines a constraint on the groups
	// that can perform the login operation that they should be using the group
	// ID specified by this field.
	BoundGroupIds pulumi.StringArrayInput
	// If set, defines a constraint on the virtual machines
	// that can perform the login operation that the location in their identity
	// document must match the one specified by this field.
	BoundLocations pulumi.StringArrayInput
	// If set, defines a constraint on the virtual
	// machines that can perform the login operation that they be associated with
	// the resource group that matches the value specified by this field.
	BoundResourceGroups pulumi.StringArrayInput
	// If set, defines a constraint on the virtual
	// machines that can perform the login operation that they must match the scale set
	// specified by this field.
	BoundScaleSets pulumi.StringArrayInput
	// If set, defines a constraint on the
	// service principals that can perform the login operation that they should be possess
	// the ids specified by this field.
	BoundServicePrincipalIds pulumi.StringArrayInput
	// If set, defines a constraint on the subscriptions
	// that can perform the login operation to ones which  matches the value specified by this
	// field.
	BoundSubscriptionIds pulumi.StringArrayInput
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
	// *Available only for Vault Enterprise*.
	Namespace pulumi.StringPtrInput
	// The name of the role.
	Role pulumi.StringInput
	// List of CIDR blocks; if set, specifies blocks of IP
	// addresses which can authenticate successfully, and ties the resulting token to these blocks
	// as well.
	TokenBoundCidrs pulumi.StringArrayInput
	// If set, will encode an
	// [explicit max TTL](https://www.vaultproject.io/docs/concepts/tokens.html#token-time-to-live-periodic-tokens-and-explicit-max-ttls)
	// onto the token in number of seconds. This is a hard cap even if `tokenTtl` and
	// `tokenMaxTtl` would otherwise allow a renewal.
	TokenExplicitMaxTtl pulumi.IntPtrInput
	// The maximum lifetime for generated tokens in number of seconds.
	// Its current value will be referenced at renewal time.
	TokenMaxTtl pulumi.IntPtrInput
	// If set, the default policy will not be set on
	// generated tokens; otherwise it will be added to the policies set in token_policies.
	TokenNoDefaultPolicy pulumi.BoolPtrInput
	// The [maximum number](https://www.vaultproject.io/api-docs/azure#token_num_uses)
	// of times a generated token may be used (within its lifetime); 0 means unlimited.
	TokenNumUses pulumi.IntPtrInput
	// If set, indicates that the
	// token generated using this role should never expire. The token should be renewed within the
	// duration specified by this value. At each renewal, the token's TTL will be set to the
	// value of this field. Specified in seconds.
	TokenPeriod pulumi.IntPtrInput
	// List of policies to encode onto generated tokens. Depending
	// on the auth method, this list may be supplemented by user/group/other values.
	TokenPolicies pulumi.StringArrayInput
	// The incremental lifetime for generated tokens in number of seconds.
	// Its current value will be referenced at renewal time.
	TokenTtl pulumi.IntPtrInput
	// The type of token that should be generated. Can be `service`,
	// `batch`, or `default` to use the mount's tuned default (which unless changed will be
	// `service` tokens). For token store roles, there are two additional possibilities:
	// `default-service` and `default-batch` which specify the type to return unless the client
	// requests a different type at generation time.
	TokenType pulumi.StringPtrInput
}

func (AuthBackendRoleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*authBackendRoleArgs)(nil)).Elem()
}

type AuthBackendRoleInput interface {
	pulumi.Input

	ToAuthBackendRoleOutput() AuthBackendRoleOutput
	ToAuthBackendRoleOutputWithContext(ctx context.Context) AuthBackendRoleOutput
}

func (*AuthBackendRole) ElementType() reflect.Type {
	return reflect.TypeOf((**AuthBackendRole)(nil)).Elem()
}

func (i *AuthBackendRole) ToAuthBackendRoleOutput() AuthBackendRoleOutput {
	return i.ToAuthBackendRoleOutputWithContext(context.Background())
}

func (i *AuthBackendRole) ToAuthBackendRoleOutputWithContext(ctx context.Context) AuthBackendRoleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AuthBackendRoleOutput)
}

// AuthBackendRoleArrayInput is an input type that accepts AuthBackendRoleArray and AuthBackendRoleArrayOutput values.
// You can construct a concrete instance of `AuthBackendRoleArrayInput` via:
//
//	AuthBackendRoleArray{ AuthBackendRoleArgs{...} }
type AuthBackendRoleArrayInput interface {
	pulumi.Input

	ToAuthBackendRoleArrayOutput() AuthBackendRoleArrayOutput
	ToAuthBackendRoleArrayOutputWithContext(context.Context) AuthBackendRoleArrayOutput
}

type AuthBackendRoleArray []AuthBackendRoleInput

func (AuthBackendRoleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AuthBackendRole)(nil)).Elem()
}

func (i AuthBackendRoleArray) ToAuthBackendRoleArrayOutput() AuthBackendRoleArrayOutput {
	return i.ToAuthBackendRoleArrayOutputWithContext(context.Background())
}

func (i AuthBackendRoleArray) ToAuthBackendRoleArrayOutputWithContext(ctx context.Context) AuthBackendRoleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AuthBackendRoleArrayOutput)
}

// AuthBackendRoleMapInput is an input type that accepts AuthBackendRoleMap and AuthBackendRoleMapOutput values.
// You can construct a concrete instance of `AuthBackendRoleMapInput` via:
//
//	AuthBackendRoleMap{ "key": AuthBackendRoleArgs{...} }
type AuthBackendRoleMapInput interface {
	pulumi.Input

	ToAuthBackendRoleMapOutput() AuthBackendRoleMapOutput
	ToAuthBackendRoleMapOutputWithContext(context.Context) AuthBackendRoleMapOutput
}

type AuthBackendRoleMap map[string]AuthBackendRoleInput

func (AuthBackendRoleMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AuthBackendRole)(nil)).Elem()
}

func (i AuthBackendRoleMap) ToAuthBackendRoleMapOutput() AuthBackendRoleMapOutput {
	return i.ToAuthBackendRoleMapOutputWithContext(context.Background())
}

func (i AuthBackendRoleMap) ToAuthBackendRoleMapOutputWithContext(ctx context.Context) AuthBackendRoleMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AuthBackendRoleMapOutput)
}

type AuthBackendRoleOutput struct{ *pulumi.OutputState }

func (AuthBackendRoleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AuthBackendRole)(nil)).Elem()
}

func (o AuthBackendRoleOutput) ToAuthBackendRoleOutput() AuthBackendRoleOutput {
	return o
}

func (o AuthBackendRoleOutput) ToAuthBackendRoleOutputWithContext(ctx context.Context) AuthBackendRoleOutput {
	return o
}

// Unique name of the auth backend to configure.
func (o AuthBackendRoleOutput) Backend() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AuthBackendRole) pulumi.StringPtrOutput { return v.Backend }).(pulumi.StringPtrOutput)
}

// If set, defines a constraint on the groups
// that can perform the login operation that they should be using the group
// ID specified by this field.
func (o AuthBackendRoleOutput) BoundGroupIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *AuthBackendRole) pulumi.StringArrayOutput { return v.BoundGroupIds }).(pulumi.StringArrayOutput)
}

// If set, defines a constraint on the virtual machines
// that can perform the login operation that the location in their identity
// document must match the one specified by this field.
func (o AuthBackendRoleOutput) BoundLocations() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *AuthBackendRole) pulumi.StringArrayOutput { return v.BoundLocations }).(pulumi.StringArrayOutput)
}

// If set, defines a constraint on the virtual
// machines that can perform the login operation that they be associated with
// the resource group that matches the value specified by this field.
func (o AuthBackendRoleOutput) BoundResourceGroups() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *AuthBackendRole) pulumi.StringArrayOutput { return v.BoundResourceGroups }).(pulumi.StringArrayOutput)
}

// If set, defines a constraint on the virtual
// machines that can perform the login operation that they must match the scale set
// specified by this field.
func (o AuthBackendRoleOutput) BoundScaleSets() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *AuthBackendRole) pulumi.StringArrayOutput { return v.BoundScaleSets }).(pulumi.StringArrayOutput)
}

// If set, defines a constraint on the
// service principals that can perform the login operation that they should be possess
// the ids specified by this field.
func (o AuthBackendRoleOutput) BoundServicePrincipalIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *AuthBackendRole) pulumi.StringArrayOutput { return v.BoundServicePrincipalIds }).(pulumi.StringArrayOutput)
}

// If set, defines a constraint on the subscriptions
// that can perform the login operation to ones which  matches the value specified by this
// field.
func (o AuthBackendRoleOutput) BoundSubscriptionIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *AuthBackendRole) pulumi.StringArrayOutput { return v.BoundSubscriptionIds }).(pulumi.StringArrayOutput)
}

// The namespace to provision the resource in.
// The value should not contain leading or trailing forward slashes.
// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
// *Available only for Vault Enterprise*.
func (o AuthBackendRoleOutput) Namespace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AuthBackendRole) pulumi.StringPtrOutput { return v.Namespace }).(pulumi.StringPtrOutput)
}

// The name of the role.
func (o AuthBackendRoleOutput) Role() pulumi.StringOutput {
	return o.ApplyT(func(v *AuthBackendRole) pulumi.StringOutput { return v.Role }).(pulumi.StringOutput)
}

// List of CIDR blocks; if set, specifies blocks of IP
// addresses which can authenticate successfully, and ties the resulting token to these blocks
// as well.
func (o AuthBackendRoleOutput) TokenBoundCidrs() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *AuthBackendRole) pulumi.StringArrayOutput { return v.TokenBoundCidrs }).(pulumi.StringArrayOutput)
}

// If set, will encode an
// [explicit max TTL](https://www.vaultproject.io/docs/concepts/tokens.html#token-time-to-live-periodic-tokens-and-explicit-max-ttls)
// onto the token in number of seconds. This is a hard cap even if `tokenTtl` and
// `tokenMaxTtl` would otherwise allow a renewal.
func (o AuthBackendRoleOutput) TokenExplicitMaxTtl() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *AuthBackendRole) pulumi.IntPtrOutput { return v.TokenExplicitMaxTtl }).(pulumi.IntPtrOutput)
}

// The maximum lifetime for generated tokens in number of seconds.
// Its current value will be referenced at renewal time.
func (o AuthBackendRoleOutput) TokenMaxTtl() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *AuthBackendRole) pulumi.IntPtrOutput { return v.TokenMaxTtl }).(pulumi.IntPtrOutput)
}

// If set, the default policy will not be set on
// generated tokens; otherwise it will be added to the policies set in token_policies.
func (o AuthBackendRoleOutput) TokenNoDefaultPolicy() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AuthBackendRole) pulumi.BoolPtrOutput { return v.TokenNoDefaultPolicy }).(pulumi.BoolPtrOutput)
}

// The [maximum number](https://www.vaultproject.io/api-docs/azure#token_num_uses)
// of times a generated token may be used (within its lifetime); 0 means unlimited.
func (o AuthBackendRoleOutput) TokenNumUses() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *AuthBackendRole) pulumi.IntPtrOutput { return v.TokenNumUses }).(pulumi.IntPtrOutput)
}

// If set, indicates that the
// token generated using this role should never expire. The token should be renewed within the
// duration specified by this value. At each renewal, the token's TTL will be set to the
// value of this field. Specified in seconds.
func (o AuthBackendRoleOutput) TokenPeriod() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *AuthBackendRole) pulumi.IntPtrOutput { return v.TokenPeriod }).(pulumi.IntPtrOutput)
}

// List of policies to encode onto generated tokens. Depending
// on the auth method, this list may be supplemented by user/group/other values.
func (o AuthBackendRoleOutput) TokenPolicies() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *AuthBackendRole) pulumi.StringArrayOutput { return v.TokenPolicies }).(pulumi.StringArrayOutput)
}

// The incremental lifetime for generated tokens in number of seconds.
// Its current value will be referenced at renewal time.
func (o AuthBackendRoleOutput) TokenTtl() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *AuthBackendRole) pulumi.IntPtrOutput { return v.TokenTtl }).(pulumi.IntPtrOutput)
}

// The type of token that should be generated. Can be `service`,
// `batch`, or `default` to use the mount's tuned default (which unless changed will be
// `service` tokens). For token store roles, there are two additional possibilities:
// `default-service` and `default-batch` which specify the type to return unless the client
// requests a different type at generation time.
func (o AuthBackendRoleOutput) TokenType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AuthBackendRole) pulumi.StringPtrOutput { return v.TokenType }).(pulumi.StringPtrOutput)
}

type AuthBackendRoleArrayOutput struct{ *pulumi.OutputState }

func (AuthBackendRoleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AuthBackendRole)(nil)).Elem()
}

func (o AuthBackendRoleArrayOutput) ToAuthBackendRoleArrayOutput() AuthBackendRoleArrayOutput {
	return o
}

func (o AuthBackendRoleArrayOutput) ToAuthBackendRoleArrayOutputWithContext(ctx context.Context) AuthBackendRoleArrayOutput {
	return o
}

func (o AuthBackendRoleArrayOutput) Index(i pulumi.IntInput) AuthBackendRoleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *AuthBackendRole {
		return vs[0].([]*AuthBackendRole)[vs[1].(int)]
	}).(AuthBackendRoleOutput)
}

type AuthBackendRoleMapOutput struct{ *pulumi.OutputState }

func (AuthBackendRoleMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AuthBackendRole)(nil)).Elem()
}

func (o AuthBackendRoleMapOutput) ToAuthBackendRoleMapOutput() AuthBackendRoleMapOutput {
	return o
}

func (o AuthBackendRoleMapOutput) ToAuthBackendRoleMapOutputWithContext(ctx context.Context) AuthBackendRoleMapOutput {
	return o
}

func (o AuthBackendRoleMapOutput) MapIndex(k pulumi.StringInput) AuthBackendRoleOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *AuthBackendRole {
		return vs[0].(map[string]*AuthBackendRole)[vs[1].(string)]
	}).(AuthBackendRoleOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AuthBackendRoleInput)(nil)).Elem(), &AuthBackendRole{})
	pulumi.RegisterInputType(reflect.TypeOf((*AuthBackendRoleArrayInput)(nil)).Elem(), AuthBackendRoleArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AuthBackendRoleMapInput)(nil)).Elem(), AuthBackendRoleMap{})
	pulumi.RegisterOutputType(AuthBackendRoleOutput{})
	pulumi.RegisterOutputType(AuthBackendRoleArrayOutput{})
	pulumi.RegisterOutputType(AuthBackendRoleMapOutput{})
}
