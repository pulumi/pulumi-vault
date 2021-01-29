// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package vault

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// ## Import
//
// Nomad secret role can be imported using the `backend`, e.g.
//
// ```sh
//  $ pulumi import vault:index/nomadSecretRole:NomadSecretRole bob nomad/role/bob
// ```
type NomadSecretRole struct {
	pulumi.CustomResourceState

	// The unique path this backend should be mounted at. Must
	// not begin or end with a `/`. Defaults to `nomad`.
	Backend pulumi.StringOutput `pulumi:"backend"`
	// Specifies if the generated token should be global. Defaults to
	// false.
	Global pulumi.BoolOutput `pulumi:"global"`
	// List of policies attached to the generated token. This setting is only used
	// when `type` is 'client'.
	Policies pulumi.StringArrayOutput `pulumi:"policies"`
	// The name to identify this role within the backend.
	// Must be unique within the backend.
	Role pulumi.StringOutput `pulumi:"role"`
	// Specifies the type of token to create when using this role. Valid
	// settings are 'client' and 'management'. Defaults to 'client'.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewNomadSecretRole registers a new resource with the given unique name, arguments, and options.
func NewNomadSecretRole(ctx *pulumi.Context,
	name string, args *NomadSecretRoleArgs, opts ...pulumi.ResourceOption) (*NomadSecretRole, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Backend == nil {
		return nil, errors.New("invalid value for required argument 'Backend'")
	}
	if args.Role == nil {
		return nil, errors.New("invalid value for required argument 'Role'")
	}
	var resource NomadSecretRole
	err := ctx.RegisterResource("vault:index/nomadSecretRole:NomadSecretRole", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNomadSecretRole gets an existing NomadSecretRole resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNomadSecretRole(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NomadSecretRoleState, opts ...pulumi.ResourceOption) (*NomadSecretRole, error) {
	var resource NomadSecretRole
	err := ctx.ReadResource("vault:index/nomadSecretRole:NomadSecretRole", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering NomadSecretRole resources.
type nomadSecretRoleState struct {
	// The unique path this backend should be mounted at. Must
	// not begin or end with a `/`. Defaults to `nomad`.
	Backend *string `pulumi:"backend"`
	// Specifies if the generated token should be global. Defaults to
	// false.
	Global *bool `pulumi:"global"`
	// List of policies attached to the generated token. This setting is only used
	// when `type` is 'client'.
	Policies []string `pulumi:"policies"`
	// The name to identify this role within the backend.
	// Must be unique within the backend.
	Role *string `pulumi:"role"`
	// Specifies the type of token to create when using this role. Valid
	// settings are 'client' and 'management'. Defaults to 'client'.
	Type *string `pulumi:"type"`
}

type NomadSecretRoleState struct {
	// The unique path this backend should be mounted at. Must
	// not begin or end with a `/`. Defaults to `nomad`.
	Backend pulumi.StringPtrInput
	// Specifies if the generated token should be global. Defaults to
	// false.
	Global pulumi.BoolPtrInput
	// List of policies attached to the generated token. This setting is only used
	// when `type` is 'client'.
	Policies pulumi.StringArrayInput
	// The name to identify this role within the backend.
	// Must be unique within the backend.
	Role pulumi.StringPtrInput
	// Specifies the type of token to create when using this role. Valid
	// settings are 'client' and 'management'. Defaults to 'client'.
	Type pulumi.StringPtrInput
}

func (NomadSecretRoleState) ElementType() reflect.Type {
	return reflect.TypeOf((*nomadSecretRoleState)(nil)).Elem()
}

type nomadSecretRoleArgs struct {
	// The unique path this backend should be mounted at. Must
	// not begin or end with a `/`. Defaults to `nomad`.
	Backend string `pulumi:"backend"`
	// Specifies if the generated token should be global. Defaults to
	// false.
	Global *bool `pulumi:"global"`
	// List of policies attached to the generated token. This setting is only used
	// when `type` is 'client'.
	Policies []string `pulumi:"policies"`
	// The name to identify this role within the backend.
	// Must be unique within the backend.
	Role string `pulumi:"role"`
	// Specifies the type of token to create when using this role. Valid
	// settings are 'client' and 'management'. Defaults to 'client'.
	Type *string `pulumi:"type"`
}

// The set of arguments for constructing a NomadSecretRole resource.
type NomadSecretRoleArgs struct {
	// The unique path this backend should be mounted at. Must
	// not begin or end with a `/`. Defaults to `nomad`.
	Backend pulumi.StringInput
	// Specifies if the generated token should be global. Defaults to
	// false.
	Global pulumi.BoolPtrInput
	// List of policies attached to the generated token. This setting is only used
	// when `type` is 'client'.
	Policies pulumi.StringArrayInput
	// The name to identify this role within the backend.
	// Must be unique within the backend.
	Role pulumi.StringInput
	// Specifies the type of token to create when using this role. Valid
	// settings are 'client' and 'management'. Defaults to 'client'.
	Type pulumi.StringPtrInput
}

func (NomadSecretRoleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*nomadSecretRoleArgs)(nil)).Elem()
}

type NomadSecretRoleInput interface {
	pulumi.Input

	ToNomadSecretRoleOutput() NomadSecretRoleOutput
	ToNomadSecretRoleOutputWithContext(ctx context.Context) NomadSecretRoleOutput
}

func (*NomadSecretRole) ElementType() reflect.Type {
	return reflect.TypeOf((*NomadSecretRole)(nil))
}

func (i *NomadSecretRole) ToNomadSecretRoleOutput() NomadSecretRoleOutput {
	return i.ToNomadSecretRoleOutputWithContext(context.Background())
}

func (i *NomadSecretRole) ToNomadSecretRoleOutputWithContext(ctx context.Context) NomadSecretRoleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NomadSecretRoleOutput)
}

type NomadSecretRoleOutput struct {
	*pulumi.OutputState
}

func (NomadSecretRoleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NomadSecretRole)(nil))
}

func (o NomadSecretRoleOutput) ToNomadSecretRoleOutput() NomadSecretRoleOutput {
	return o
}

func (o NomadSecretRoleOutput) ToNomadSecretRoleOutputWithContext(ctx context.Context) NomadSecretRoleOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(NomadSecretRoleOutput{})
}
