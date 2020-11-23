// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package identity

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// ## Import
//
// The key can be imported with the key name, for example
//
// ```sh
//  $ pulumi import vault:identity/oidcKey:OidcKey key key
// ```
type OidcKey struct {
	pulumi.CustomResourceState

	// Signing algorithm to use. Signing algorithm to use.
	// Allowed values are: RS256 (default), RS384, RS512, ES256, ES384, ES512, EdDSA.
	Algorithm pulumi.StringPtrOutput `pulumi:"algorithm"`
	// Array of role client ids allowed to use this key for signing. If empty, no roles are allowed. If "*", all roles are
	// allowed.
	AllowedClientIds pulumi.StringArrayOutput `pulumi:"allowedClientIds"`
	// Name of the OIDC Key to create.
	Name pulumi.StringOutput `pulumi:"name"`
	// How often to generate a new signing key in number of seconds
	RotationPeriod pulumi.IntPtrOutput `pulumi:"rotationPeriod"`
	// "Controls how long the public portion of a signing key will be
	// available for verification after being rotated in seconds.
	VerificationTtl pulumi.IntPtrOutput `pulumi:"verificationTtl"`
}

// NewOidcKey registers a new resource with the given unique name, arguments, and options.
func NewOidcKey(ctx *pulumi.Context,
	name string, args *OidcKeyArgs, opts ...pulumi.ResourceOption) (*OidcKey, error) {
	if args == nil {
		args = &OidcKeyArgs{}
	}
	var resource OidcKey
	err := ctx.RegisterResource("vault:identity/oidcKey:OidcKey", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetOidcKey gets an existing OidcKey resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetOidcKey(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *OidcKeyState, opts ...pulumi.ResourceOption) (*OidcKey, error) {
	var resource OidcKey
	err := ctx.ReadResource("vault:identity/oidcKey:OidcKey", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering OidcKey resources.
type oidcKeyState struct {
	// Signing algorithm to use. Signing algorithm to use.
	// Allowed values are: RS256 (default), RS384, RS512, ES256, ES384, ES512, EdDSA.
	Algorithm *string `pulumi:"algorithm"`
	// Array of role client ids allowed to use this key for signing. If empty, no roles are allowed. If "*", all roles are
	// allowed.
	AllowedClientIds []string `pulumi:"allowedClientIds"`
	// Name of the OIDC Key to create.
	Name *string `pulumi:"name"`
	// How often to generate a new signing key in number of seconds
	RotationPeriod *int `pulumi:"rotationPeriod"`
	// "Controls how long the public portion of a signing key will be
	// available for verification after being rotated in seconds.
	VerificationTtl *int `pulumi:"verificationTtl"`
}

type OidcKeyState struct {
	// Signing algorithm to use. Signing algorithm to use.
	// Allowed values are: RS256 (default), RS384, RS512, ES256, ES384, ES512, EdDSA.
	Algorithm pulumi.StringPtrInput
	// Array of role client ids allowed to use this key for signing. If empty, no roles are allowed. If "*", all roles are
	// allowed.
	AllowedClientIds pulumi.StringArrayInput
	// Name of the OIDC Key to create.
	Name pulumi.StringPtrInput
	// How often to generate a new signing key in number of seconds
	RotationPeriod pulumi.IntPtrInput
	// "Controls how long the public portion of a signing key will be
	// available for verification after being rotated in seconds.
	VerificationTtl pulumi.IntPtrInput
}

func (OidcKeyState) ElementType() reflect.Type {
	return reflect.TypeOf((*oidcKeyState)(nil)).Elem()
}

type oidcKeyArgs struct {
	// Signing algorithm to use. Signing algorithm to use.
	// Allowed values are: RS256 (default), RS384, RS512, ES256, ES384, ES512, EdDSA.
	Algorithm *string `pulumi:"algorithm"`
	// Array of role client ids allowed to use this key for signing. If empty, no roles are allowed. If "*", all roles are
	// allowed.
	AllowedClientIds []string `pulumi:"allowedClientIds"`
	// Name of the OIDC Key to create.
	Name *string `pulumi:"name"`
	// How often to generate a new signing key in number of seconds
	RotationPeriod *int `pulumi:"rotationPeriod"`
	// "Controls how long the public portion of a signing key will be
	// available for verification after being rotated in seconds.
	VerificationTtl *int `pulumi:"verificationTtl"`
}

// The set of arguments for constructing a OidcKey resource.
type OidcKeyArgs struct {
	// Signing algorithm to use. Signing algorithm to use.
	// Allowed values are: RS256 (default), RS384, RS512, ES256, ES384, ES512, EdDSA.
	Algorithm pulumi.StringPtrInput
	// Array of role client ids allowed to use this key for signing. If empty, no roles are allowed. If "*", all roles are
	// allowed.
	AllowedClientIds pulumi.StringArrayInput
	// Name of the OIDC Key to create.
	Name pulumi.StringPtrInput
	// How often to generate a new signing key in number of seconds
	RotationPeriod pulumi.IntPtrInput
	// "Controls how long the public portion of a signing key will be
	// available for verification after being rotated in seconds.
	VerificationTtl pulumi.IntPtrInput
}

func (OidcKeyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*oidcKeyArgs)(nil)).Elem()
}

type OidcKeyInput interface {
	pulumi.Input

	ToOidcKeyOutput() OidcKeyOutput
	ToOidcKeyOutputWithContext(ctx context.Context) OidcKeyOutput
}

func (OidcKey) ElementType() reflect.Type {
	return reflect.TypeOf((*OidcKey)(nil)).Elem()
}

func (i OidcKey) ToOidcKeyOutput() OidcKeyOutput {
	return i.ToOidcKeyOutputWithContext(context.Background())
}

func (i OidcKey) ToOidcKeyOutputWithContext(ctx context.Context) OidcKeyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OidcKeyOutput)
}

type OidcKeyOutput struct {
	*pulumi.OutputState
}

func (OidcKeyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*OidcKeyOutput)(nil)).Elem()
}

func (o OidcKeyOutput) ToOidcKeyOutput() OidcKeyOutput {
	return o
}

func (o OidcKeyOutput) ToOidcKeyOutputWithContext(ctx context.Context) OidcKeyOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(OidcKeyOutput{})
}
