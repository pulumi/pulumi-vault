// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package vault

import (
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

type Token struct {
	pulumi.CustomResourceState

	// String containing the client token if stored in present file
	ClientToken pulumi.StringOutput `pulumi:"clientToken"`
	// String containing the token display name
	DisplayName pulumi.StringPtrOutput `pulumi:"displayName"`
	// String containing the client token encrypted with the given `pgpKey` if stored in present file
	EncryptedClientToken pulumi.StringOutput `pulumi:"encryptedClientToken"`
	// The explicit max TTL of this token
	ExplicitMaxTtl pulumi.StringPtrOutput `pulumi:"explicitMaxTtl"`
	// String containing the token lease duration if present in state file
	LeaseDuration pulumi.IntOutput `pulumi:"leaseDuration"`
	// String containing the token lease started time if present in state file
	LeaseStarted pulumi.StringOutput `pulumi:"leaseStarted"`
	// Flag to not attach the default policy to this token
	NoDefaultPolicy pulumi.BoolPtrOutput `pulumi:"noDefaultPolicy"`
	// Flag to create a token without parent
	NoParent pulumi.BoolOutput `pulumi:"noParent"`
	// The number of allowed uses of this token
	NumUses pulumi.IntOutput `pulumi:"numUses"`
	// The period of this token
	Period pulumi.StringPtrOutput `pulumi:"period"`
	// The PGP key (base64 encoded) to encrypt the token.
	PgpKey pulumi.StringPtrOutput `pulumi:"pgpKey"`
	// List of policies to attach to this token
	Policies pulumi.StringArrayOutput `pulumi:"policies"`
	// The renew increment
	RenewIncrement pulumi.IntPtrOutput `pulumi:"renewIncrement"`
	// The minimal lease to renew this token
	RenewMinLease pulumi.IntPtrOutput `pulumi:"renewMinLease"`
	// Flag to allow to renew this token
	Renewable pulumi.BoolOutput `pulumi:"renewable"`
	// The token role name
	RoleName pulumi.StringPtrOutput `pulumi:"roleName"`
	// The TTL period of this token
	Ttl pulumi.StringPtrOutput `pulumi:"ttl"`
	// The client wrapped token.
	WrappedToken pulumi.StringOutput `pulumi:"wrappedToken"`
	// The client wrapping accessor.
	WrappingAccessor pulumi.StringOutput `pulumi:"wrappingAccessor"`
	// The TTL period of the wrapped token.
	WrappingTtl pulumi.StringPtrOutput `pulumi:"wrappingTtl"`
}

// NewToken registers a new resource with the given unique name, arguments, and options.
func NewToken(ctx *pulumi.Context,
	name string, args *TokenArgs, opts ...pulumi.ResourceOption) (*Token, error) {
	if args == nil {
		args = &TokenArgs{}
	}
	var resource Token
	err := ctx.RegisterResource("vault:index/token:Token", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetToken gets an existing Token resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetToken(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TokenState, opts ...pulumi.ResourceOption) (*Token, error) {
	var resource Token
	err := ctx.ReadResource("vault:index/token:Token", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Token resources.
type tokenState struct {
	// String containing the client token if stored in present file
	ClientToken *string `pulumi:"clientToken"`
	// String containing the token display name
	DisplayName *string `pulumi:"displayName"`
	// String containing the client token encrypted with the given `pgpKey` if stored in present file
	EncryptedClientToken *string `pulumi:"encryptedClientToken"`
	// The explicit max TTL of this token
	ExplicitMaxTtl *string `pulumi:"explicitMaxTtl"`
	// String containing the token lease duration if present in state file
	LeaseDuration *int `pulumi:"leaseDuration"`
	// String containing the token lease started time if present in state file
	LeaseStarted *string `pulumi:"leaseStarted"`
	// Flag to not attach the default policy to this token
	NoDefaultPolicy *bool `pulumi:"noDefaultPolicy"`
	// Flag to create a token without parent
	NoParent *bool `pulumi:"noParent"`
	// The number of allowed uses of this token
	NumUses *int `pulumi:"numUses"`
	// The period of this token
	Period *string `pulumi:"period"`
	// The PGP key (base64 encoded) to encrypt the token.
	PgpKey *string `pulumi:"pgpKey"`
	// List of policies to attach to this token
	Policies []string `pulumi:"policies"`
	// The renew increment
	RenewIncrement *int `pulumi:"renewIncrement"`
	// The minimal lease to renew this token
	RenewMinLease *int `pulumi:"renewMinLease"`
	// Flag to allow to renew this token
	Renewable *bool `pulumi:"renewable"`
	// The token role name
	RoleName *string `pulumi:"roleName"`
	// The TTL period of this token
	Ttl *string `pulumi:"ttl"`
	// The client wrapped token.
	WrappedToken *string `pulumi:"wrappedToken"`
	// The client wrapping accessor.
	WrappingAccessor *string `pulumi:"wrappingAccessor"`
	// The TTL period of the wrapped token.
	WrappingTtl *string `pulumi:"wrappingTtl"`
}

type TokenState struct {
	// String containing the client token if stored in present file
	ClientToken pulumi.StringPtrInput
	// String containing the token display name
	DisplayName pulumi.StringPtrInput
	// String containing the client token encrypted with the given `pgpKey` if stored in present file
	EncryptedClientToken pulumi.StringPtrInput
	// The explicit max TTL of this token
	ExplicitMaxTtl pulumi.StringPtrInput
	// String containing the token lease duration if present in state file
	LeaseDuration pulumi.IntPtrInput
	// String containing the token lease started time if present in state file
	LeaseStarted pulumi.StringPtrInput
	// Flag to not attach the default policy to this token
	NoDefaultPolicy pulumi.BoolPtrInput
	// Flag to create a token without parent
	NoParent pulumi.BoolPtrInput
	// The number of allowed uses of this token
	NumUses pulumi.IntPtrInput
	// The period of this token
	Period pulumi.StringPtrInput
	// The PGP key (base64 encoded) to encrypt the token.
	PgpKey pulumi.StringPtrInput
	// List of policies to attach to this token
	Policies pulumi.StringArrayInput
	// The renew increment
	RenewIncrement pulumi.IntPtrInput
	// The minimal lease to renew this token
	RenewMinLease pulumi.IntPtrInput
	// Flag to allow to renew this token
	Renewable pulumi.BoolPtrInput
	// The token role name
	RoleName pulumi.StringPtrInput
	// The TTL period of this token
	Ttl pulumi.StringPtrInput
	// The client wrapped token.
	WrappedToken pulumi.StringPtrInput
	// The client wrapping accessor.
	WrappingAccessor pulumi.StringPtrInput
	// The TTL period of the wrapped token.
	WrappingTtl pulumi.StringPtrInput
}

func (TokenState) ElementType() reflect.Type {
	return reflect.TypeOf((*tokenState)(nil)).Elem()
}

type tokenArgs struct {
	// String containing the token display name
	DisplayName *string `pulumi:"displayName"`
	// The explicit max TTL of this token
	ExplicitMaxTtl *string `pulumi:"explicitMaxTtl"`
	// Flag to not attach the default policy to this token
	NoDefaultPolicy *bool `pulumi:"noDefaultPolicy"`
	// Flag to create a token without parent
	NoParent *bool `pulumi:"noParent"`
	// The number of allowed uses of this token
	NumUses *int `pulumi:"numUses"`
	// The period of this token
	Period *string `pulumi:"period"`
	// The PGP key (base64 encoded) to encrypt the token.
	PgpKey *string `pulumi:"pgpKey"`
	// List of policies to attach to this token
	Policies []string `pulumi:"policies"`
	// The renew increment
	RenewIncrement *int `pulumi:"renewIncrement"`
	// The minimal lease to renew this token
	RenewMinLease *int `pulumi:"renewMinLease"`
	// Flag to allow to renew this token
	Renewable *bool `pulumi:"renewable"`
	// The token role name
	RoleName *string `pulumi:"roleName"`
	// The TTL period of this token
	Ttl *string `pulumi:"ttl"`
	// The TTL period of the wrapped token.
	WrappingTtl *string `pulumi:"wrappingTtl"`
}

// The set of arguments for constructing a Token resource.
type TokenArgs struct {
	// String containing the token display name
	DisplayName pulumi.StringPtrInput
	// The explicit max TTL of this token
	ExplicitMaxTtl pulumi.StringPtrInput
	// Flag to not attach the default policy to this token
	NoDefaultPolicy pulumi.BoolPtrInput
	// Flag to create a token without parent
	NoParent pulumi.BoolPtrInput
	// The number of allowed uses of this token
	NumUses pulumi.IntPtrInput
	// The period of this token
	Period pulumi.StringPtrInput
	// The PGP key (base64 encoded) to encrypt the token.
	PgpKey pulumi.StringPtrInput
	// List of policies to attach to this token
	Policies pulumi.StringArrayInput
	// The renew increment
	RenewIncrement pulumi.IntPtrInput
	// The minimal lease to renew this token
	RenewMinLease pulumi.IntPtrInput
	// Flag to allow to renew this token
	Renewable pulumi.BoolPtrInput
	// The token role name
	RoleName pulumi.StringPtrInput
	// The TTL period of this token
	Ttl pulumi.StringPtrInput
	// The TTL period of the wrapped token.
	WrappingTtl pulumi.StringPtrInput
}

func (TokenArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*tokenArgs)(nil)).Elem()
}
