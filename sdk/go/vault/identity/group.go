// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package identity

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Creates an Identity Group for Vault. The [Identity secrets engine](https://www.vaultproject.io/docs/secrets/identity/index.html) is the identity management solution for Vault.
//
// A group can contain multiple entities as its members. A group can also have subgroups. Policies set on the group is granted to all members of the group. During request time, when the token's entity ID is being evaluated for the policies that it has access to; along with the policies on the entity itself, policies that are inherited due to group memberships are also granted.
//
// ## Example Usage
// ### Internal Group
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-vault/sdk/v3/go/vault/identity"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := identity.NewGroup(ctx, "internal", &identity.GroupArgs{
// 			Metadata: pulumi.StringMap{
// 				"version": pulumi.String("2"),
// 			},
// 			Policies: pulumi.StringArray{
// 				pulumi.String("dev"),
// 				pulumi.String("test"),
// 			},
// 			Type: pulumi.String("internal"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### External Group
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-vault/sdk/v3/go/vault/identity"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := identity.NewGroup(ctx, "group", &identity.GroupArgs{
// 			Metadata: pulumi.StringMap{
// 				"version": pulumi.String("1"),
// 			},
// 			Policies: pulumi.StringArray{
// 				pulumi.String("test"),
// 			},
// 			Type: pulumi.String("external"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type Group struct {
	pulumi.CustomResourceState

	// Manage member entities externally through `vault_identity_group_policies_member_entity_ids`
	ExternalMemberEntityIds pulumi.BoolPtrOutput `pulumi:"externalMemberEntityIds"`
	// Manage policies externally through `vault_identity_group_policies`, allows using group ID in assigned policies.
	ExternalPolicies pulumi.BoolPtrOutput `pulumi:"externalPolicies"`
	// A list of Entity IDs to be assigned as group members. Not allowed on `external` groups.
	MemberEntityIds pulumi.StringArrayOutput `pulumi:"memberEntityIds"`
	// A list of Group IDs to be assigned as group members.
	MemberGroupIds pulumi.StringArrayOutput `pulumi:"memberGroupIds"`
	// A Map of additional metadata to associate with the group.
	Metadata pulumi.StringMapOutput `pulumi:"metadata"`
	// Name of the identity group to create.
	Name pulumi.StringOutput `pulumi:"name"`
	// A list of policies to apply to the group.
	Policies pulumi.StringArrayOutput `pulumi:"policies"`
	// Type of the group, internal or external. Defaults to `internal`.
	Type pulumi.StringPtrOutput `pulumi:"type"`
}

// NewGroup registers a new resource with the given unique name, arguments, and options.
func NewGroup(ctx *pulumi.Context,
	name string, args *GroupArgs, opts ...pulumi.ResourceOption) (*Group, error) {
	if args == nil {
		args = &GroupArgs{}
	}
	var resource Group
	err := ctx.RegisterResource("vault:identity/group:Group", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetGroup gets an existing Group resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GroupState, opts ...pulumi.ResourceOption) (*Group, error) {
	var resource Group
	err := ctx.ReadResource("vault:identity/group:Group", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Group resources.
type groupState struct {
	// Manage member entities externally through `vault_identity_group_policies_member_entity_ids`
	ExternalMemberEntityIds *bool `pulumi:"externalMemberEntityIds"`
	// Manage policies externally through `vault_identity_group_policies`, allows using group ID in assigned policies.
	ExternalPolicies *bool `pulumi:"externalPolicies"`
	// A list of Entity IDs to be assigned as group members. Not allowed on `external` groups.
	MemberEntityIds []string `pulumi:"memberEntityIds"`
	// A list of Group IDs to be assigned as group members.
	MemberGroupIds []string `pulumi:"memberGroupIds"`
	// A Map of additional metadata to associate with the group.
	Metadata map[string]string `pulumi:"metadata"`
	// Name of the identity group to create.
	Name *string `pulumi:"name"`
	// A list of policies to apply to the group.
	Policies []string `pulumi:"policies"`
	// Type of the group, internal or external. Defaults to `internal`.
	Type *string `pulumi:"type"`
}

type GroupState struct {
	// Manage member entities externally through `vault_identity_group_policies_member_entity_ids`
	ExternalMemberEntityIds pulumi.BoolPtrInput
	// Manage policies externally through `vault_identity_group_policies`, allows using group ID in assigned policies.
	ExternalPolicies pulumi.BoolPtrInput
	// A list of Entity IDs to be assigned as group members. Not allowed on `external` groups.
	MemberEntityIds pulumi.StringArrayInput
	// A list of Group IDs to be assigned as group members.
	MemberGroupIds pulumi.StringArrayInput
	// A Map of additional metadata to associate with the group.
	Metadata pulumi.StringMapInput
	// Name of the identity group to create.
	Name pulumi.StringPtrInput
	// A list of policies to apply to the group.
	Policies pulumi.StringArrayInput
	// Type of the group, internal or external. Defaults to `internal`.
	Type pulumi.StringPtrInput
}

func (GroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*groupState)(nil)).Elem()
}

type groupArgs struct {
	// Manage member entities externally through `vault_identity_group_policies_member_entity_ids`
	ExternalMemberEntityIds *bool `pulumi:"externalMemberEntityIds"`
	// Manage policies externally through `vault_identity_group_policies`, allows using group ID in assigned policies.
	ExternalPolicies *bool `pulumi:"externalPolicies"`
	// A list of Entity IDs to be assigned as group members. Not allowed on `external` groups.
	MemberEntityIds []string `pulumi:"memberEntityIds"`
	// A list of Group IDs to be assigned as group members.
	MemberGroupIds []string `pulumi:"memberGroupIds"`
	// A Map of additional metadata to associate with the group.
	Metadata map[string]string `pulumi:"metadata"`
	// Name of the identity group to create.
	Name *string `pulumi:"name"`
	// A list of policies to apply to the group.
	Policies []string `pulumi:"policies"`
	// Type of the group, internal or external. Defaults to `internal`.
	Type *string `pulumi:"type"`
}

// The set of arguments for constructing a Group resource.
type GroupArgs struct {
	// Manage member entities externally through `vault_identity_group_policies_member_entity_ids`
	ExternalMemberEntityIds pulumi.BoolPtrInput
	// Manage policies externally through `vault_identity_group_policies`, allows using group ID in assigned policies.
	ExternalPolicies pulumi.BoolPtrInput
	// A list of Entity IDs to be assigned as group members. Not allowed on `external` groups.
	MemberEntityIds pulumi.StringArrayInput
	// A list of Group IDs to be assigned as group members.
	MemberGroupIds pulumi.StringArrayInput
	// A Map of additional metadata to associate with the group.
	Metadata pulumi.StringMapInput
	// Name of the identity group to create.
	Name pulumi.StringPtrInput
	// A list of policies to apply to the group.
	Policies pulumi.StringArrayInput
	// Type of the group, internal or external. Defaults to `internal`.
	Type pulumi.StringPtrInput
}

func (GroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*groupArgs)(nil)).Elem()
}

type GroupInput interface {
	pulumi.Input

	ToGroupOutput() GroupOutput
	ToGroupOutputWithContext(ctx context.Context) GroupOutput
}

func (Group) ElementType() reflect.Type {
	return reflect.TypeOf((*Group)(nil)).Elem()
}

func (i Group) ToGroupOutput() GroupOutput {
	return i.ToGroupOutputWithContext(context.Background())
}

func (i Group) ToGroupOutputWithContext(ctx context.Context) GroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GroupOutput)
}

type GroupOutput struct {
	*pulumi.OutputState
}

func (GroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GroupOutput)(nil)).Elem()
}

func (o GroupOutput) ToGroupOutput() GroupOutput {
	return o
}

func (o GroupOutput) ToGroupOutputWithContext(ctx context.Context) GroupOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(GroupOutput{})
}
