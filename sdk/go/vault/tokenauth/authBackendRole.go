// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package tokenauth

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-vault/sdk/v6/go/vault/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages Token auth backend role in a Vault server. See the [Vault
// documentation](https://www.vaultproject.io/docs/auth/token.html) for more
// information.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-vault/sdk/v6/go/vault/tokenauth"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := tokenauth.NewAuthBackendRole(ctx, "example", &tokenauth.AuthBackendRoleArgs{
//				RoleName: pulumi.String("my-role"),
//				AllowedPolicies: pulumi.StringArray{
//					pulumi.String("dev"),
//					pulumi.String("test"),
//				},
//				DisallowedPolicies: pulumi.StringArray{
//					pulumi.String("default"),
//				},
//				AllowedEntityAliases: pulumi.StringArray{
//					pulumi.String("test_entity"),
//				},
//				Orphan:              pulumi.Bool(true),
//				TokenPeriod:         pulumi.Int(86400),
//				Renewable:           pulumi.Bool(true),
//				TokenExplicitMaxTtl: pulumi.Int(115200),
//				PathSuffix:          pulumi.String("path-suffix"),
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
// Token auth backend roles can be imported with `auth/token/roles/` followed by the `role_name`, e.g.
//
// ```sh
// $ pulumi import vault:tokenauth/authBackendRole:AuthBackendRole example auth/token/roles/my-role
// ```
type AuthBackendRole struct {
	pulumi.CustomResourceState

	// List of allowed entity aliases.
	AllowedEntityAliases pulumi.StringArrayOutput `pulumi:"allowedEntityAliases"`
	// List of allowed policies for given role.
	AllowedPolicies pulumi.StringArrayOutput `pulumi:"allowedPolicies"`
	// Set of allowed policies with glob match for given role.
	AllowedPoliciesGlobs pulumi.StringArrayOutput `pulumi:"allowedPoliciesGlobs"`
	// List of disallowed policies for given role.
	DisallowedPolicies pulumi.StringArrayOutput `pulumi:"disallowedPolicies"`
	// Set of disallowed policies with glob match for given role.
	DisallowedPoliciesGlobs pulumi.StringArrayOutput `pulumi:"disallowedPoliciesGlobs"`
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
	// *Available only for Vault Enterprise*.
	Namespace pulumi.StringPtrOutput `pulumi:"namespace"`
	// If true, tokens created against this policy will be orphan tokens.
	Orphan pulumi.BoolPtrOutput `pulumi:"orphan"`
	// Tokens created against this role will have the given suffix as part of their path in addition to the role name.
	//
	// > Due to a bug the resource. This *will* cause all existing tokens issued by this role to be revoked.
	PathSuffix pulumi.StringPtrOutput `pulumi:"pathSuffix"`
	// Whether to disable the ability of the token to be renewed past its initial TTL.
	Renewable pulumi.BoolPtrOutput `pulumi:"renewable"`
	// The name of the role.
	RoleName pulumi.StringOutput `pulumi:"roleName"`
	// Specifies the blocks of IP addresses which are allowed to use the generated token
	TokenBoundCidrs pulumi.StringArrayOutput `pulumi:"tokenBoundCidrs"`
	// Generated Token's Explicit Maximum TTL in seconds
	TokenExplicitMaxTtl pulumi.IntPtrOutput `pulumi:"tokenExplicitMaxTtl"`
	// The maximum lifetime of the generated token
	TokenMaxTtl pulumi.IntPtrOutput `pulumi:"tokenMaxTtl"`
	// If true, the 'default' policy will not automatically be added to generated tokens
	TokenNoDefaultPolicy pulumi.BoolPtrOutput `pulumi:"tokenNoDefaultPolicy"`
	// The maximum number of times a token may be used, a value of zero means unlimited
	TokenNumUses pulumi.IntPtrOutput `pulumi:"tokenNumUses"`
	// Generated Token's Period
	TokenPeriod pulumi.IntPtrOutput `pulumi:"tokenPeriod"`
	// Generated Token's Policies
	TokenPolicies pulumi.StringArrayOutput `pulumi:"tokenPolicies"`
	// The initial ttl of the token to generate in seconds
	TokenTtl pulumi.IntPtrOutput `pulumi:"tokenTtl"`
	// The type of token to generate, service or batch
	TokenType pulumi.StringPtrOutput `pulumi:"tokenType"`
}

// NewAuthBackendRole registers a new resource with the given unique name, arguments, and options.
func NewAuthBackendRole(ctx *pulumi.Context,
	name string, args *AuthBackendRoleArgs, opts ...pulumi.ResourceOption) (*AuthBackendRole, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.RoleName == nil {
		return nil, errors.New("invalid value for required argument 'RoleName'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource AuthBackendRole
	err := ctx.RegisterResource("vault:tokenauth/authBackendRole:AuthBackendRole", name, args, &resource, opts...)
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
	err := ctx.ReadResource("vault:tokenauth/authBackendRole:AuthBackendRole", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AuthBackendRole resources.
type authBackendRoleState struct {
	// List of allowed entity aliases.
	AllowedEntityAliases []string `pulumi:"allowedEntityAliases"`
	// List of allowed policies for given role.
	AllowedPolicies []string `pulumi:"allowedPolicies"`
	// Set of allowed policies with glob match for given role.
	AllowedPoliciesGlobs []string `pulumi:"allowedPoliciesGlobs"`
	// List of disallowed policies for given role.
	DisallowedPolicies []string `pulumi:"disallowedPolicies"`
	// Set of disallowed policies with glob match for given role.
	DisallowedPoliciesGlobs []string `pulumi:"disallowedPoliciesGlobs"`
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
	// *Available only for Vault Enterprise*.
	Namespace *string `pulumi:"namespace"`
	// If true, tokens created against this policy will be orphan tokens.
	Orphan *bool `pulumi:"orphan"`
	// Tokens created against this role will have the given suffix as part of their path in addition to the role name.
	//
	// > Due to a bug the resource. This *will* cause all existing tokens issued by this role to be revoked.
	PathSuffix *string `pulumi:"pathSuffix"`
	// Whether to disable the ability of the token to be renewed past its initial TTL.
	Renewable *bool `pulumi:"renewable"`
	// The name of the role.
	RoleName *string `pulumi:"roleName"`
	// Specifies the blocks of IP addresses which are allowed to use the generated token
	TokenBoundCidrs []string `pulumi:"tokenBoundCidrs"`
	// Generated Token's Explicit Maximum TTL in seconds
	TokenExplicitMaxTtl *int `pulumi:"tokenExplicitMaxTtl"`
	// The maximum lifetime of the generated token
	TokenMaxTtl *int `pulumi:"tokenMaxTtl"`
	// If true, the 'default' policy will not automatically be added to generated tokens
	TokenNoDefaultPolicy *bool `pulumi:"tokenNoDefaultPolicy"`
	// The maximum number of times a token may be used, a value of zero means unlimited
	TokenNumUses *int `pulumi:"tokenNumUses"`
	// Generated Token's Period
	TokenPeriod *int `pulumi:"tokenPeriod"`
	// Generated Token's Policies
	TokenPolicies []string `pulumi:"tokenPolicies"`
	// The initial ttl of the token to generate in seconds
	TokenTtl *int `pulumi:"tokenTtl"`
	// The type of token to generate, service or batch
	TokenType *string `pulumi:"tokenType"`
}

type AuthBackendRoleState struct {
	// List of allowed entity aliases.
	AllowedEntityAliases pulumi.StringArrayInput
	// List of allowed policies for given role.
	AllowedPolicies pulumi.StringArrayInput
	// Set of allowed policies with glob match for given role.
	AllowedPoliciesGlobs pulumi.StringArrayInput
	// List of disallowed policies for given role.
	DisallowedPolicies pulumi.StringArrayInput
	// Set of disallowed policies with glob match for given role.
	DisallowedPoliciesGlobs pulumi.StringArrayInput
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
	// *Available only for Vault Enterprise*.
	Namespace pulumi.StringPtrInput
	// If true, tokens created against this policy will be orphan tokens.
	Orphan pulumi.BoolPtrInput
	// Tokens created against this role will have the given suffix as part of their path in addition to the role name.
	//
	// > Due to a bug the resource. This *will* cause all existing tokens issued by this role to be revoked.
	PathSuffix pulumi.StringPtrInput
	// Whether to disable the ability of the token to be renewed past its initial TTL.
	Renewable pulumi.BoolPtrInput
	// The name of the role.
	RoleName pulumi.StringPtrInput
	// Specifies the blocks of IP addresses which are allowed to use the generated token
	TokenBoundCidrs pulumi.StringArrayInput
	// Generated Token's Explicit Maximum TTL in seconds
	TokenExplicitMaxTtl pulumi.IntPtrInput
	// The maximum lifetime of the generated token
	TokenMaxTtl pulumi.IntPtrInput
	// If true, the 'default' policy will not automatically be added to generated tokens
	TokenNoDefaultPolicy pulumi.BoolPtrInput
	// The maximum number of times a token may be used, a value of zero means unlimited
	TokenNumUses pulumi.IntPtrInput
	// Generated Token's Period
	TokenPeriod pulumi.IntPtrInput
	// Generated Token's Policies
	TokenPolicies pulumi.StringArrayInput
	// The initial ttl of the token to generate in seconds
	TokenTtl pulumi.IntPtrInput
	// The type of token to generate, service or batch
	TokenType pulumi.StringPtrInput
}

func (AuthBackendRoleState) ElementType() reflect.Type {
	return reflect.TypeOf((*authBackendRoleState)(nil)).Elem()
}

type authBackendRoleArgs struct {
	// List of allowed entity aliases.
	AllowedEntityAliases []string `pulumi:"allowedEntityAliases"`
	// List of allowed policies for given role.
	AllowedPolicies []string `pulumi:"allowedPolicies"`
	// Set of allowed policies with glob match for given role.
	AllowedPoliciesGlobs []string `pulumi:"allowedPoliciesGlobs"`
	// List of disallowed policies for given role.
	DisallowedPolicies []string `pulumi:"disallowedPolicies"`
	// Set of disallowed policies with glob match for given role.
	DisallowedPoliciesGlobs []string `pulumi:"disallowedPoliciesGlobs"`
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
	// *Available only for Vault Enterprise*.
	Namespace *string `pulumi:"namespace"`
	// If true, tokens created against this policy will be orphan tokens.
	Orphan *bool `pulumi:"orphan"`
	// Tokens created against this role will have the given suffix as part of their path in addition to the role name.
	//
	// > Due to a bug the resource. This *will* cause all existing tokens issued by this role to be revoked.
	PathSuffix *string `pulumi:"pathSuffix"`
	// Whether to disable the ability of the token to be renewed past its initial TTL.
	Renewable *bool `pulumi:"renewable"`
	// The name of the role.
	RoleName string `pulumi:"roleName"`
	// Specifies the blocks of IP addresses which are allowed to use the generated token
	TokenBoundCidrs []string `pulumi:"tokenBoundCidrs"`
	// Generated Token's Explicit Maximum TTL in seconds
	TokenExplicitMaxTtl *int `pulumi:"tokenExplicitMaxTtl"`
	// The maximum lifetime of the generated token
	TokenMaxTtl *int `pulumi:"tokenMaxTtl"`
	// If true, the 'default' policy will not automatically be added to generated tokens
	TokenNoDefaultPolicy *bool `pulumi:"tokenNoDefaultPolicy"`
	// The maximum number of times a token may be used, a value of zero means unlimited
	TokenNumUses *int `pulumi:"tokenNumUses"`
	// Generated Token's Period
	TokenPeriod *int `pulumi:"tokenPeriod"`
	// Generated Token's Policies
	TokenPolicies []string `pulumi:"tokenPolicies"`
	// The initial ttl of the token to generate in seconds
	TokenTtl *int `pulumi:"tokenTtl"`
	// The type of token to generate, service or batch
	TokenType *string `pulumi:"tokenType"`
}

// The set of arguments for constructing a AuthBackendRole resource.
type AuthBackendRoleArgs struct {
	// List of allowed entity aliases.
	AllowedEntityAliases pulumi.StringArrayInput
	// List of allowed policies for given role.
	AllowedPolicies pulumi.StringArrayInput
	// Set of allowed policies with glob match for given role.
	AllowedPoliciesGlobs pulumi.StringArrayInput
	// List of disallowed policies for given role.
	DisallowedPolicies pulumi.StringArrayInput
	// Set of disallowed policies with glob match for given role.
	DisallowedPoliciesGlobs pulumi.StringArrayInput
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
	// *Available only for Vault Enterprise*.
	Namespace pulumi.StringPtrInput
	// If true, tokens created against this policy will be orphan tokens.
	Orphan pulumi.BoolPtrInput
	// Tokens created against this role will have the given suffix as part of their path in addition to the role name.
	//
	// > Due to a bug the resource. This *will* cause all existing tokens issued by this role to be revoked.
	PathSuffix pulumi.StringPtrInput
	// Whether to disable the ability of the token to be renewed past its initial TTL.
	Renewable pulumi.BoolPtrInput
	// The name of the role.
	RoleName pulumi.StringInput
	// Specifies the blocks of IP addresses which are allowed to use the generated token
	TokenBoundCidrs pulumi.StringArrayInput
	// Generated Token's Explicit Maximum TTL in seconds
	TokenExplicitMaxTtl pulumi.IntPtrInput
	// The maximum lifetime of the generated token
	TokenMaxTtl pulumi.IntPtrInput
	// If true, the 'default' policy will not automatically be added to generated tokens
	TokenNoDefaultPolicy pulumi.BoolPtrInput
	// The maximum number of times a token may be used, a value of zero means unlimited
	TokenNumUses pulumi.IntPtrInput
	// Generated Token's Period
	TokenPeriod pulumi.IntPtrInput
	// Generated Token's Policies
	TokenPolicies pulumi.StringArrayInput
	// The initial ttl of the token to generate in seconds
	TokenTtl pulumi.IntPtrInput
	// The type of token to generate, service or batch
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

// List of allowed entity aliases.
func (o AuthBackendRoleOutput) AllowedEntityAliases() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *AuthBackendRole) pulumi.StringArrayOutput { return v.AllowedEntityAliases }).(pulumi.StringArrayOutput)
}

// List of allowed policies for given role.
func (o AuthBackendRoleOutput) AllowedPolicies() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *AuthBackendRole) pulumi.StringArrayOutput { return v.AllowedPolicies }).(pulumi.StringArrayOutput)
}

// Set of allowed policies with glob match for given role.
func (o AuthBackendRoleOutput) AllowedPoliciesGlobs() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *AuthBackendRole) pulumi.StringArrayOutput { return v.AllowedPoliciesGlobs }).(pulumi.StringArrayOutput)
}

// List of disallowed policies for given role.
func (o AuthBackendRoleOutput) DisallowedPolicies() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *AuthBackendRole) pulumi.StringArrayOutput { return v.DisallowedPolicies }).(pulumi.StringArrayOutput)
}

// Set of disallowed policies with glob match for given role.
func (o AuthBackendRoleOutput) DisallowedPoliciesGlobs() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *AuthBackendRole) pulumi.StringArrayOutput { return v.DisallowedPoliciesGlobs }).(pulumi.StringArrayOutput)
}

// The namespace to provision the resource in.
// The value should not contain leading or trailing forward slashes.
// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
// *Available only for Vault Enterprise*.
func (o AuthBackendRoleOutput) Namespace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AuthBackendRole) pulumi.StringPtrOutput { return v.Namespace }).(pulumi.StringPtrOutput)
}

// If true, tokens created against this policy will be orphan tokens.
func (o AuthBackendRoleOutput) Orphan() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AuthBackendRole) pulumi.BoolPtrOutput { return v.Orphan }).(pulumi.BoolPtrOutput)
}

// Tokens created against this role will have the given suffix as part of their path in addition to the role name.
//
// > Due to a bug the resource. This *will* cause all existing tokens issued by this role to be revoked.
func (o AuthBackendRoleOutput) PathSuffix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AuthBackendRole) pulumi.StringPtrOutput { return v.PathSuffix }).(pulumi.StringPtrOutput)
}

// Whether to disable the ability of the token to be renewed past its initial TTL.
func (o AuthBackendRoleOutput) Renewable() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AuthBackendRole) pulumi.BoolPtrOutput { return v.Renewable }).(pulumi.BoolPtrOutput)
}

// The name of the role.
func (o AuthBackendRoleOutput) RoleName() pulumi.StringOutput {
	return o.ApplyT(func(v *AuthBackendRole) pulumi.StringOutput { return v.RoleName }).(pulumi.StringOutput)
}

// Specifies the blocks of IP addresses which are allowed to use the generated token
func (o AuthBackendRoleOutput) TokenBoundCidrs() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *AuthBackendRole) pulumi.StringArrayOutput { return v.TokenBoundCidrs }).(pulumi.StringArrayOutput)
}

// Generated Token's Explicit Maximum TTL in seconds
func (o AuthBackendRoleOutput) TokenExplicitMaxTtl() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *AuthBackendRole) pulumi.IntPtrOutput { return v.TokenExplicitMaxTtl }).(pulumi.IntPtrOutput)
}

// The maximum lifetime of the generated token
func (o AuthBackendRoleOutput) TokenMaxTtl() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *AuthBackendRole) pulumi.IntPtrOutput { return v.TokenMaxTtl }).(pulumi.IntPtrOutput)
}

// If true, the 'default' policy will not automatically be added to generated tokens
func (o AuthBackendRoleOutput) TokenNoDefaultPolicy() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AuthBackendRole) pulumi.BoolPtrOutput { return v.TokenNoDefaultPolicy }).(pulumi.BoolPtrOutput)
}

// The maximum number of times a token may be used, a value of zero means unlimited
func (o AuthBackendRoleOutput) TokenNumUses() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *AuthBackendRole) pulumi.IntPtrOutput { return v.TokenNumUses }).(pulumi.IntPtrOutput)
}

// Generated Token's Period
func (o AuthBackendRoleOutput) TokenPeriod() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *AuthBackendRole) pulumi.IntPtrOutput { return v.TokenPeriod }).(pulumi.IntPtrOutput)
}

// Generated Token's Policies
func (o AuthBackendRoleOutput) TokenPolicies() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *AuthBackendRole) pulumi.StringArrayOutput { return v.TokenPolicies }).(pulumi.StringArrayOutput)
}

// The initial ttl of the token to generate in seconds
func (o AuthBackendRoleOutput) TokenTtl() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *AuthBackendRole) pulumi.IntPtrOutput { return v.TokenTtl }).(pulumi.IntPtrOutput)
}

// The type of token to generate, service or batch
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
