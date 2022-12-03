// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package identity

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
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
// 	"github.com/pulumi/pulumi-vault/sdk/v5/go/vault/identity"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
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
// 	"github.com/pulumi/pulumi-vault/sdk/v5/go/vault/identity"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
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
// ## Caveats
//
// It's important to note that Vault identity groups names are *case-insensitive*. For example the following resources would be equivalent.
// Applying this configuration would result in the provider failing to create one of the identity groups, since the resources share the same `name`.
//
// This sort of pattern should be avoided:
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-vault/sdk/v5/go/vault/identity"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := identity.NewGroup(ctx, "internalIdentity/groupGroup", &identity.GroupArgs{
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
// 		_, err = identity.NewGroup(ctx, "internalGroup", &identity.GroupArgs{
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
//
// ## Import
//
// Identity group can be imported using the `id`, e.g.
//
// ```sh
//  $ pulumi import vault:identity/group:Group test 'fcbf1efb-2b69-4209-bed8-811e3475dad3'
// ```
type Group struct {
	pulumi.CustomResourceState

	// `false` by default. If set to `true`, this resource will ignore any Entity IDs returned from Vault or specified in the resource. You can use `identity.GroupMemberEntityIds` to manage Entity IDs for this group in a decoupled manner.
	ExternalMemberEntityIds pulumi.BoolPtrOutput `pulumi:"externalMemberEntityIds"`
	// `false` by default. If set to `true`, this resource will ignore any policies returned from Vault or specified in the resource. You can use `identity.GroupPolicies` to manage policies for this group in a decoupled manner.
	ExternalPolicies pulumi.BoolPtrOutput `pulumi:"externalPolicies"`
	// A list of Entity IDs to be assigned as group members. Not allowed on `external` groups.
	MemberEntityIds pulumi.StringArrayOutput `pulumi:"memberEntityIds"`
	// A list of Group IDs to be assigned as group members. Not allowed on `external` groups.
	MemberGroupIds pulumi.StringArrayOutput `pulumi:"memberGroupIds"`
	// A Map of additional metadata to associate with the group.
	Metadata pulumi.StringMapOutput `pulumi:"metadata"`
	// Name of the identity group to create.
	Name pulumi.StringOutput `pulumi:"name"`
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
	// *Available only for Vault Enterprise*.
	Namespace pulumi.StringPtrOutput `pulumi:"namespace"`
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
	// `false` by default. If set to `true`, this resource will ignore any Entity IDs returned from Vault or specified in the resource. You can use `identity.GroupMemberEntityIds` to manage Entity IDs for this group in a decoupled manner.
	ExternalMemberEntityIds *bool `pulumi:"externalMemberEntityIds"`
	// `false` by default. If set to `true`, this resource will ignore any policies returned from Vault or specified in the resource. You can use `identity.GroupPolicies` to manage policies for this group in a decoupled manner.
	ExternalPolicies *bool `pulumi:"externalPolicies"`
	// A list of Entity IDs to be assigned as group members. Not allowed on `external` groups.
	MemberEntityIds []string `pulumi:"memberEntityIds"`
	// A list of Group IDs to be assigned as group members. Not allowed on `external` groups.
	MemberGroupIds []string `pulumi:"memberGroupIds"`
	// A Map of additional metadata to associate with the group.
	Metadata map[string]string `pulumi:"metadata"`
	// Name of the identity group to create.
	Name *string `pulumi:"name"`
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
	// *Available only for Vault Enterprise*.
	Namespace *string `pulumi:"namespace"`
	// A list of policies to apply to the group.
	Policies []string `pulumi:"policies"`
	// Type of the group, internal or external. Defaults to `internal`.
	Type *string `pulumi:"type"`
}

type GroupState struct {
	// `false` by default. If set to `true`, this resource will ignore any Entity IDs returned from Vault or specified in the resource. You can use `identity.GroupMemberEntityIds` to manage Entity IDs for this group in a decoupled manner.
	ExternalMemberEntityIds pulumi.BoolPtrInput
	// `false` by default. If set to `true`, this resource will ignore any policies returned from Vault or specified in the resource. You can use `identity.GroupPolicies` to manage policies for this group in a decoupled manner.
	ExternalPolicies pulumi.BoolPtrInput
	// A list of Entity IDs to be assigned as group members. Not allowed on `external` groups.
	MemberEntityIds pulumi.StringArrayInput
	// A list of Group IDs to be assigned as group members. Not allowed on `external` groups.
	MemberGroupIds pulumi.StringArrayInput
	// A Map of additional metadata to associate with the group.
	Metadata pulumi.StringMapInput
	// Name of the identity group to create.
	Name pulumi.StringPtrInput
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
	// *Available only for Vault Enterprise*.
	Namespace pulumi.StringPtrInput
	// A list of policies to apply to the group.
	Policies pulumi.StringArrayInput
	// Type of the group, internal or external. Defaults to `internal`.
	Type pulumi.StringPtrInput
}

func (GroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*groupState)(nil)).Elem()
}

type groupArgs struct {
	// `false` by default. If set to `true`, this resource will ignore any Entity IDs returned from Vault or specified in the resource. You can use `identity.GroupMemberEntityIds` to manage Entity IDs for this group in a decoupled manner.
	ExternalMemberEntityIds *bool `pulumi:"externalMemberEntityIds"`
	// `false` by default. If set to `true`, this resource will ignore any policies returned from Vault or specified in the resource. You can use `identity.GroupPolicies` to manage policies for this group in a decoupled manner.
	ExternalPolicies *bool `pulumi:"externalPolicies"`
	// A list of Entity IDs to be assigned as group members. Not allowed on `external` groups.
	MemberEntityIds []string `pulumi:"memberEntityIds"`
	// A list of Group IDs to be assigned as group members. Not allowed on `external` groups.
	MemberGroupIds []string `pulumi:"memberGroupIds"`
	// A Map of additional metadata to associate with the group.
	Metadata map[string]string `pulumi:"metadata"`
	// Name of the identity group to create.
	Name *string `pulumi:"name"`
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
	// *Available only for Vault Enterprise*.
	Namespace *string `pulumi:"namespace"`
	// A list of policies to apply to the group.
	Policies []string `pulumi:"policies"`
	// Type of the group, internal or external. Defaults to `internal`.
	Type *string `pulumi:"type"`
}

// The set of arguments for constructing a Group resource.
type GroupArgs struct {
	// `false` by default. If set to `true`, this resource will ignore any Entity IDs returned from Vault or specified in the resource. You can use `identity.GroupMemberEntityIds` to manage Entity IDs for this group in a decoupled manner.
	ExternalMemberEntityIds pulumi.BoolPtrInput
	// `false` by default. If set to `true`, this resource will ignore any policies returned from Vault or specified in the resource. You can use `identity.GroupPolicies` to manage policies for this group in a decoupled manner.
	ExternalPolicies pulumi.BoolPtrInput
	// A list of Entity IDs to be assigned as group members. Not allowed on `external` groups.
	MemberEntityIds pulumi.StringArrayInput
	// A list of Group IDs to be assigned as group members. Not allowed on `external` groups.
	MemberGroupIds pulumi.StringArrayInput
	// A Map of additional metadata to associate with the group.
	Metadata pulumi.StringMapInput
	// Name of the identity group to create.
	Name pulumi.StringPtrInput
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
	// *Available only for Vault Enterprise*.
	Namespace pulumi.StringPtrInput
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

func (*Group) ElementType() reflect.Type {
	return reflect.TypeOf((**Group)(nil)).Elem()
}

func (i *Group) ToGroupOutput() GroupOutput {
	return i.ToGroupOutputWithContext(context.Background())
}

func (i *Group) ToGroupOutputWithContext(ctx context.Context) GroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GroupOutput)
}

// GroupArrayInput is an input type that accepts GroupArray and GroupArrayOutput values.
// You can construct a concrete instance of `GroupArrayInput` via:
//
//          GroupArray{ GroupArgs{...} }
type GroupArrayInput interface {
	pulumi.Input

	ToGroupArrayOutput() GroupArrayOutput
	ToGroupArrayOutputWithContext(context.Context) GroupArrayOutput
}

type GroupArray []GroupInput

func (GroupArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Group)(nil)).Elem()
}

func (i GroupArray) ToGroupArrayOutput() GroupArrayOutput {
	return i.ToGroupArrayOutputWithContext(context.Background())
}

func (i GroupArray) ToGroupArrayOutputWithContext(ctx context.Context) GroupArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GroupArrayOutput)
}

// GroupMapInput is an input type that accepts GroupMap and GroupMapOutput values.
// You can construct a concrete instance of `GroupMapInput` via:
//
//          GroupMap{ "key": GroupArgs{...} }
type GroupMapInput interface {
	pulumi.Input

	ToGroupMapOutput() GroupMapOutput
	ToGroupMapOutputWithContext(context.Context) GroupMapOutput
}

type GroupMap map[string]GroupInput

func (GroupMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Group)(nil)).Elem()
}

func (i GroupMap) ToGroupMapOutput() GroupMapOutput {
	return i.ToGroupMapOutputWithContext(context.Background())
}

func (i GroupMap) ToGroupMapOutputWithContext(ctx context.Context) GroupMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GroupMapOutput)
}

type GroupOutput struct{ *pulumi.OutputState }

func (GroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Group)(nil)).Elem()
}

func (o GroupOutput) ToGroupOutput() GroupOutput {
	return o
}

func (o GroupOutput) ToGroupOutputWithContext(ctx context.Context) GroupOutput {
	return o
}

// `false` by default. If set to `true`, this resource will ignore any Entity IDs returned from Vault or specified in the resource. You can use `identity.GroupMemberEntityIds` to manage Entity IDs for this group in a decoupled manner.
func (o GroupOutput) ExternalMemberEntityIds() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Group) pulumi.BoolPtrOutput { return v.ExternalMemberEntityIds }).(pulumi.BoolPtrOutput)
}

// `false` by default. If set to `true`, this resource will ignore any policies returned from Vault or specified in the resource. You can use `identity.GroupPolicies` to manage policies for this group in a decoupled manner.
func (o GroupOutput) ExternalPolicies() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Group) pulumi.BoolPtrOutput { return v.ExternalPolicies }).(pulumi.BoolPtrOutput)
}

// A list of Entity IDs to be assigned as group members. Not allowed on `external` groups.
func (o GroupOutput) MemberEntityIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Group) pulumi.StringArrayOutput { return v.MemberEntityIds }).(pulumi.StringArrayOutput)
}

// A list of Group IDs to be assigned as group members. Not allowed on `external` groups.
func (o GroupOutput) MemberGroupIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Group) pulumi.StringArrayOutput { return v.MemberGroupIds }).(pulumi.StringArrayOutput)
}

// A Map of additional metadata to associate with the group.
func (o GroupOutput) Metadata() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Group) pulumi.StringMapOutput { return v.Metadata }).(pulumi.StringMapOutput)
}

// Name of the identity group to create.
func (o GroupOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Group) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The namespace to provision the resource in.
// The value should not contain leading or trailing forward slashes.
// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
// *Available only for Vault Enterprise*.
func (o GroupOutput) Namespace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Group) pulumi.StringPtrOutput { return v.Namespace }).(pulumi.StringPtrOutput)
}

// A list of policies to apply to the group.
func (o GroupOutput) Policies() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Group) pulumi.StringArrayOutput { return v.Policies }).(pulumi.StringArrayOutput)
}

// Type of the group, internal or external. Defaults to `internal`.
func (o GroupOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Group) pulumi.StringPtrOutput { return v.Type }).(pulumi.StringPtrOutput)
}

type GroupArrayOutput struct{ *pulumi.OutputState }

func (GroupArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Group)(nil)).Elem()
}

func (o GroupArrayOutput) ToGroupArrayOutput() GroupArrayOutput {
	return o
}

func (o GroupArrayOutput) ToGroupArrayOutputWithContext(ctx context.Context) GroupArrayOutput {
	return o
}

func (o GroupArrayOutput) Index(i pulumi.IntInput) GroupOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Group {
		return vs[0].([]*Group)[vs[1].(int)]
	}).(GroupOutput)
}

type GroupMapOutput struct{ *pulumi.OutputState }

func (GroupMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Group)(nil)).Elem()
}

func (o GroupMapOutput) ToGroupMapOutput() GroupMapOutput {
	return o
}

func (o GroupMapOutput) ToGroupMapOutputWithContext(ctx context.Context) GroupMapOutput {
	return o
}

func (o GroupMapOutput) MapIndex(k pulumi.StringInput) GroupOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Group {
		return vs[0].(map[string]*Group)[vs[1].(string)]
	}).(GroupOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*GroupInput)(nil)).Elem(), &Group{})
	pulumi.RegisterInputType(reflect.TypeOf((*GroupArrayInput)(nil)).Elem(), GroupArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*GroupMapInput)(nil)).Elem(), GroupMap{})
	pulumi.RegisterOutputType(GroupOutput{})
	pulumi.RegisterOutputType(GroupArrayOutput{})
	pulumi.RegisterOutputType(GroupMapOutput{})
}
