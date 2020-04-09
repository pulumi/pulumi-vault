// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package okta

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides a resource to create a group in an
// [Okta auth backend within Vault](https://www.vaultproject.io/docs/auth/okta.html).
//
//
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-vault/blob/master/website/docs/r/okta_auth_backend_group.html.md.
type AuthBackendGroup struct {
	pulumi.CustomResourceState

	// Name of the group within the Okta
	GroupName pulumi.StringOutput `pulumi:"groupName"`
	// The path where the Okta auth backend is mounted
	Path pulumi.StringOutput `pulumi:"path"`
	// Vault policies to associate with this group
	Policies pulumi.StringArrayOutput `pulumi:"policies"`
}

// NewAuthBackendGroup registers a new resource with the given unique name, arguments, and options.
func NewAuthBackendGroup(ctx *pulumi.Context,
	name string, args *AuthBackendGroupArgs, opts ...pulumi.ResourceOption) (*AuthBackendGroup, error) {
	if args == nil || args.GroupName == nil {
		return nil, errors.New("missing required argument 'GroupName'")
	}
	if args == nil || args.Path == nil {
		return nil, errors.New("missing required argument 'Path'")
	}
	if args == nil {
		args = &AuthBackendGroupArgs{}
	}
	var resource AuthBackendGroup
	err := ctx.RegisterResource("vault:okta/authBackendGroup:AuthBackendGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAuthBackendGroup gets an existing AuthBackendGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAuthBackendGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AuthBackendGroupState, opts ...pulumi.ResourceOption) (*AuthBackendGroup, error) {
	var resource AuthBackendGroup
	err := ctx.ReadResource("vault:okta/authBackendGroup:AuthBackendGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AuthBackendGroup resources.
type authBackendGroupState struct {
	// Name of the group within the Okta
	GroupName *string `pulumi:"groupName"`
	// The path where the Okta auth backend is mounted
	Path *string `pulumi:"path"`
	// Vault policies to associate with this group
	Policies []string `pulumi:"policies"`
}

type AuthBackendGroupState struct {
	// Name of the group within the Okta
	GroupName pulumi.StringPtrInput
	// The path where the Okta auth backend is mounted
	Path pulumi.StringPtrInput
	// Vault policies to associate with this group
	Policies pulumi.StringArrayInput
}

func (AuthBackendGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*authBackendGroupState)(nil)).Elem()
}

type authBackendGroupArgs struct {
	// Name of the group within the Okta
	GroupName string `pulumi:"groupName"`
	// The path where the Okta auth backend is mounted
	Path string `pulumi:"path"`
	// Vault policies to associate with this group
	Policies []string `pulumi:"policies"`
}

// The set of arguments for constructing a AuthBackendGroup resource.
type AuthBackendGroupArgs struct {
	// Name of the group within the Okta
	GroupName pulumi.StringInput
	// The path where the Okta auth backend is mounted
	Path pulumi.StringInput
	// Vault policies to associate with this group
	Policies pulumi.StringArrayInput
}

func (AuthBackendGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*authBackendGroupArgs)(nil)).Elem()
}
