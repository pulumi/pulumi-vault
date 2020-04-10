// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package github

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages policy mappings for Github Users authenticated via Github. See the [Vault
// documentation](https://www.vaultproject.io/docs/auth/github.html) for more
// information.
type User struct {
	pulumi.CustomResourceState

	// Path where the github auth backend is mounted. Defaults to `github`
	// if not specified.
	Backend pulumi.StringPtrOutput `pulumi:"backend"`
	// An array of strings specifying the policies to be set on tokens issued
	// using this role.
	Policies pulumi.StringArrayOutput `pulumi:"policies"`
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
	// GitHub user name.
	User pulumi.StringOutput `pulumi:"user"`
}

// NewUser registers a new resource with the given unique name, arguments, and options.
func NewUser(ctx *pulumi.Context,
	name string, args *UserArgs, opts ...pulumi.ResourceOption) (*User, error) {
	if args == nil || args.User == nil {
		return nil, errors.New("missing required argument 'User'")
	}
	if args == nil {
		args = &UserArgs{}
	}
	var resource User
	err := ctx.RegisterResource("vault:github/user:User", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetUser gets an existing User resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetUser(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *UserState, opts ...pulumi.ResourceOption) (*User, error) {
	var resource User
	err := ctx.ReadResource("vault:github/user:User", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering User resources.
type userState struct {
	// Path where the github auth backend is mounted. Defaults to `github`
	// if not specified.
	Backend *string `pulumi:"backend"`
	// An array of strings specifying the policies to be set on tokens issued
	// using this role.
	Policies []string `pulumi:"policies"`
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
	// GitHub user name.
	User *string `pulumi:"user"`
}

type UserState struct {
	// Path where the github auth backend is mounted. Defaults to `github`
	// if not specified.
	Backend pulumi.StringPtrInput
	// An array of strings specifying the policies to be set on tokens issued
	// using this role.
	Policies pulumi.StringArrayInput
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
	// GitHub user name.
	User pulumi.StringPtrInput
}

func (UserState) ElementType() reflect.Type {
	return reflect.TypeOf((*userState)(nil)).Elem()
}

type userArgs struct {
	// Path where the github auth backend is mounted. Defaults to `github`
	// if not specified.
	Backend *string `pulumi:"backend"`
	// An array of strings specifying the policies to be set on tokens issued
	// using this role.
	Policies []string `pulumi:"policies"`
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
	// GitHub user name.
	User string `pulumi:"user"`
}

// The set of arguments for constructing a User resource.
type UserArgs struct {
	// Path where the github auth backend is mounted. Defaults to `github`
	// if not specified.
	Backend pulumi.StringPtrInput
	// An array of strings specifying the policies to be set on tokens issued
	// using this role.
	Policies pulumi.StringArrayInput
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
	// GitHub user name.
	User pulumi.StringInput
}

func (UserArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*userArgs)(nil)).Elem()
}
