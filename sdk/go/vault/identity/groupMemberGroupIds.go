// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package identity

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-vault/sdk/v6/go/vault/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages member groups for an Identity Group for Vault. The
// [Identity secrets engine](https://www.vaultproject.io/docs/secrets/identity/index.html)
// is the identity management solution for Vault.
//
// ## Example Usage
//
// ### Exclusive Member Groups
//
// <!--Start PulumiCodeChooser -->
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-vault/sdk/v6/go/vault/identity"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			internal, err := identity.NewGroup(ctx, "internal", &identity.GroupArgs{
//				Type:                   pulumi.String("internal"),
//				ExternalMemberGroupIds: pulumi.Bool(true),
//				Metadata: pulumi.StringMap{
//					"version": pulumi.String("2"),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			users, err := identity.NewGroup(ctx, "users", &identity.GroupArgs{
//				Metadata: pulumi.StringMap{
//					"version": pulumi.String("2"),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			_, err = identity.NewGroupMemberGroupIds(ctx, "members", &identity.GroupMemberGroupIdsArgs{
//				Exclusive: pulumi.Bool(true),
//				MemberGroupIds: pulumi.StringArray{
//					users.ID(),
//				},
//				GroupId: internal.ID(),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// <!--End PulumiCodeChooser -->
//
// ### Non-Exclusive Member Groups
//
// <!--Start PulumiCodeChooser -->
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-vault/sdk/v6/go/vault/identity"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			internal, err := identity.NewGroup(ctx, "internal", &identity.GroupArgs{
//				Type:                   pulumi.String("internal"),
//				ExternalMemberGroupIds: pulumi.Bool(true),
//				Metadata: pulumi.StringMap{
//					"version": pulumi.String("2"),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			users, err := identity.NewGroup(ctx, "users", &identity.GroupArgs{
//				Metadata: pulumi.StringMap{
//					"version": pulumi.String("2"),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			_, err = identity.NewGroupMemberGroupIds(ctx, "members", &identity.GroupMemberGroupIdsArgs{
//				Exclusive: pulumi.Bool(false),
//				MemberGroupIds: pulumi.StringArray{
//					users.ID(),
//				},
//				GroupId: internal.ID(),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// <!--End PulumiCodeChooser -->
type GroupMemberGroupIds struct {
	pulumi.CustomResourceState

	// Defaults to `true`.
	//
	// If `true`, this resource will take exclusive control of the member groups that belong to the group and will set
	// it equal to what is specified in the resource.
	//
	// If set to `false`, this resource will simply ensure that the member groups specified in the resource are present
	// in the group. When destroying the resource, the resource will ensure that the member groups specified in the resource
	// are removed.
	Exclusive pulumi.BoolPtrOutput `pulumi:"exclusive"`
	// Group ID to assign member entities to.
	GroupId pulumi.StringOutput `pulumi:"groupId"`
	// List of member groups that belong to the group
	MemberGroupIds pulumi.StringArrayOutput `pulumi:"memberGroupIds"`
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
	// *Available only for Vault Enterprise*.
	Namespace pulumi.StringPtrOutput `pulumi:"namespace"`
}

// NewGroupMemberGroupIds registers a new resource with the given unique name, arguments, and options.
func NewGroupMemberGroupIds(ctx *pulumi.Context,
	name string, args *GroupMemberGroupIdsArgs, opts ...pulumi.ResourceOption) (*GroupMemberGroupIds, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.GroupId == nil {
		return nil, errors.New("invalid value for required argument 'GroupId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource GroupMemberGroupIds
	err := ctx.RegisterResource("vault:identity/groupMemberGroupIds:GroupMemberGroupIds", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetGroupMemberGroupIds gets an existing GroupMemberGroupIds resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGroupMemberGroupIds(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GroupMemberGroupIdsState, opts ...pulumi.ResourceOption) (*GroupMemberGroupIds, error) {
	var resource GroupMemberGroupIds
	err := ctx.ReadResource("vault:identity/groupMemberGroupIds:GroupMemberGroupIds", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering GroupMemberGroupIds resources.
type groupMemberGroupIdsState struct {
	// Defaults to `true`.
	//
	// If `true`, this resource will take exclusive control of the member groups that belong to the group and will set
	// it equal to what is specified in the resource.
	//
	// If set to `false`, this resource will simply ensure that the member groups specified in the resource are present
	// in the group. When destroying the resource, the resource will ensure that the member groups specified in the resource
	// are removed.
	Exclusive *bool `pulumi:"exclusive"`
	// Group ID to assign member entities to.
	GroupId *string `pulumi:"groupId"`
	// List of member groups that belong to the group
	MemberGroupIds []string `pulumi:"memberGroupIds"`
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
	// *Available only for Vault Enterprise*.
	Namespace *string `pulumi:"namespace"`
}

type GroupMemberGroupIdsState struct {
	// Defaults to `true`.
	//
	// If `true`, this resource will take exclusive control of the member groups that belong to the group and will set
	// it equal to what is specified in the resource.
	//
	// If set to `false`, this resource will simply ensure that the member groups specified in the resource are present
	// in the group. When destroying the resource, the resource will ensure that the member groups specified in the resource
	// are removed.
	Exclusive pulumi.BoolPtrInput
	// Group ID to assign member entities to.
	GroupId pulumi.StringPtrInput
	// List of member groups that belong to the group
	MemberGroupIds pulumi.StringArrayInput
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
	// *Available only for Vault Enterprise*.
	Namespace pulumi.StringPtrInput
}

func (GroupMemberGroupIdsState) ElementType() reflect.Type {
	return reflect.TypeOf((*groupMemberGroupIdsState)(nil)).Elem()
}

type groupMemberGroupIdsArgs struct {
	// Defaults to `true`.
	//
	// If `true`, this resource will take exclusive control of the member groups that belong to the group and will set
	// it equal to what is specified in the resource.
	//
	// If set to `false`, this resource will simply ensure that the member groups specified in the resource are present
	// in the group. When destroying the resource, the resource will ensure that the member groups specified in the resource
	// are removed.
	Exclusive *bool `pulumi:"exclusive"`
	// Group ID to assign member entities to.
	GroupId string `pulumi:"groupId"`
	// List of member groups that belong to the group
	MemberGroupIds []string `pulumi:"memberGroupIds"`
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
	// *Available only for Vault Enterprise*.
	Namespace *string `pulumi:"namespace"`
}

// The set of arguments for constructing a GroupMemberGroupIds resource.
type GroupMemberGroupIdsArgs struct {
	// Defaults to `true`.
	//
	// If `true`, this resource will take exclusive control of the member groups that belong to the group and will set
	// it equal to what is specified in the resource.
	//
	// If set to `false`, this resource will simply ensure that the member groups specified in the resource are present
	// in the group. When destroying the resource, the resource will ensure that the member groups specified in the resource
	// are removed.
	Exclusive pulumi.BoolPtrInput
	// Group ID to assign member entities to.
	GroupId pulumi.StringInput
	// List of member groups that belong to the group
	MemberGroupIds pulumi.StringArrayInput
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
	// *Available only for Vault Enterprise*.
	Namespace pulumi.StringPtrInput
}

func (GroupMemberGroupIdsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*groupMemberGroupIdsArgs)(nil)).Elem()
}

type GroupMemberGroupIdsInput interface {
	pulumi.Input

	ToGroupMemberGroupIdsOutput() GroupMemberGroupIdsOutput
	ToGroupMemberGroupIdsOutputWithContext(ctx context.Context) GroupMemberGroupIdsOutput
}

func (*GroupMemberGroupIds) ElementType() reflect.Type {
	return reflect.TypeOf((**GroupMemberGroupIds)(nil)).Elem()
}

func (i *GroupMemberGroupIds) ToGroupMemberGroupIdsOutput() GroupMemberGroupIdsOutput {
	return i.ToGroupMemberGroupIdsOutputWithContext(context.Background())
}

func (i *GroupMemberGroupIds) ToGroupMemberGroupIdsOutputWithContext(ctx context.Context) GroupMemberGroupIdsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GroupMemberGroupIdsOutput)
}

// GroupMemberGroupIdsArrayInput is an input type that accepts GroupMemberGroupIdsArray and GroupMemberGroupIdsArrayOutput values.
// You can construct a concrete instance of `GroupMemberGroupIdsArrayInput` via:
//
//	GroupMemberGroupIdsArray{ GroupMemberGroupIdsArgs{...} }
type GroupMemberGroupIdsArrayInput interface {
	pulumi.Input

	ToGroupMemberGroupIdsArrayOutput() GroupMemberGroupIdsArrayOutput
	ToGroupMemberGroupIdsArrayOutputWithContext(context.Context) GroupMemberGroupIdsArrayOutput
}

type GroupMemberGroupIdsArray []GroupMemberGroupIdsInput

func (GroupMemberGroupIdsArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*GroupMemberGroupIds)(nil)).Elem()
}

func (i GroupMemberGroupIdsArray) ToGroupMemberGroupIdsArrayOutput() GroupMemberGroupIdsArrayOutput {
	return i.ToGroupMemberGroupIdsArrayOutputWithContext(context.Background())
}

func (i GroupMemberGroupIdsArray) ToGroupMemberGroupIdsArrayOutputWithContext(ctx context.Context) GroupMemberGroupIdsArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GroupMemberGroupIdsArrayOutput)
}

// GroupMemberGroupIdsMapInput is an input type that accepts GroupMemberGroupIdsMap and GroupMemberGroupIdsMapOutput values.
// You can construct a concrete instance of `GroupMemberGroupIdsMapInput` via:
//
//	GroupMemberGroupIdsMap{ "key": GroupMemberGroupIdsArgs{...} }
type GroupMemberGroupIdsMapInput interface {
	pulumi.Input

	ToGroupMemberGroupIdsMapOutput() GroupMemberGroupIdsMapOutput
	ToGroupMemberGroupIdsMapOutputWithContext(context.Context) GroupMemberGroupIdsMapOutput
}

type GroupMemberGroupIdsMap map[string]GroupMemberGroupIdsInput

func (GroupMemberGroupIdsMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*GroupMemberGroupIds)(nil)).Elem()
}

func (i GroupMemberGroupIdsMap) ToGroupMemberGroupIdsMapOutput() GroupMemberGroupIdsMapOutput {
	return i.ToGroupMemberGroupIdsMapOutputWithContext(context.Background())
}

func (i GroupMemberGroupIdsMap) ToGroupMemberGroupIdsMapOutputWithContext(ctx context.Context) GroupMemberGroupIdsMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GroupMemberGroupIdsMapOutput)
}

type GroupMemberGroupIdsOutput struct{ *pulumi.OutputState }

func (GroupMemberGroupIdsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GroupMemberGroupIds)(nil)).Elem()
}

func (o GroupMemberGroupIdsOutput) ToGroupMemberGroupIdsOutput() GroupMemberGroupIdsOutput {
	return o
}

func (o GroupMemberGroupIdsOutput) ToGroupMemberGroupIdsOutputWithContext(ctx context.Context) GroupMemberGroupIdsOutput {
	return o
}

// Defaults to `true`.
//
// If `true`, this resource will take exclusive control of the member groups that belong to the group and will set
// it equal to what is specified in the resource.
//
// If set to `false`, this resource will simply ensure that the member groups specified in the resource are present
// in the group. When destroying the resource, the resource will ensure that the member groups specified in the resource
// are removed.
func (o GroupMemberGroupIdsOutput) Exclusive() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *GroupMemberGroupIds) pulumi.BoolPtrOutput { return v.Exclusive }).(pulumi.BoolPtrOutput)
}

// Group ID to assign member entities to.
func (o GroupMemberGroupIdsOutput) GroupId() pulumi.StringOutput {
	return o.ApplyT(func(v *GroupMemberGroupIds) pulumi.StringOutput { return v.GroupId }).(pulumi.StringOutput)
}

// List of member groups that belong to the group
func (o GroupMemberGroupIdsOutput) MemberGroupIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *GroupMemberGroupIds) pulumi.StringArrayOutput { return v.MemberGroupIds }).(pulumi.StringArrayOutput)
}

// The namespace to provision the resource in.
// The value should not contain leading or trailing forward slashes.
// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
// *Available only for Vault Enterprise*.
func (o GroupMemberGroupIdsOutput) Namespace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GroupMemberGroupIds) pulumi.StringPtrOutput { return v.Namespace }).(pulumi.StringPtrOutput)
}

type GroupMemberGroupIdsArrayOutput struct{ *pulumi.OutputState }

func (GroupMemberGroupIdsArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*GroupMemberGroupIds)(nil)).Elem()
}

func (o GroupMemberGroupIdsArrayOutput) ToGroupMemberGroupIdsArrayOutput() GroupMemberGroupIdsArrayOutput {
	return o
}

func (o GroupMemberGroupIdsArrayOutput) ToGroupMemberGroupIdsArrayOutputWithContext(ctx context.Context) GroupMemberGroupIdsArrayOutput {
	return o
}

func (o GroupMemberGroupIdsArrayOutput) Index(i pulumi.IntInput) GroupMemberGroupIdsOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *GroupMemberGroupIds {
		return vs[0].([]*GroupMemberGroupIds)[vs[1].(int)]
	}).(GroupMemberGroupIdsOutput)
}

type GroupMemberGroupIdsMapOutput struct{ *pulumi.OutputState }

func (GroupMemberGroupIdsMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*GroupMemberGroupIds)(nil)).Elem()
}

func (o GroupMemberGroupIdsMapOutput) ToGroupMemberGroupIdsMapOutput() GroupMemberGroupIdsMapOutput {
	return o
}

func (o GroupMemberGroupIdsMapOutput) ToGroupMemberGroupIdsMapOutputWithContext(ctx context.Context) GroupMemberGroupIdsMapOutput {
	return o
}

func (o GroupMemberGroupIdsMapOutput) MapIndex(k pulumi.StringInput) GroupMemberGroupIdsOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *GroupMemberGroupIds {
		return vs[0].(map[string]*GroupMemberGroupIds)[vs[1].(string)]
	}).(GroupMemberGroupIdsOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*GroupMemberGroupIdsInput)(nil)).Elem(), &GroupMemberGroupIds{})
	pulumi.RegisterInputType(reflect.TypeOf((*GroupMemberGroupIdsArrayInput)(nil)).Elem(), GroupMemberGroupIdsArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*GroupMemberGroupIdsMapInput)(nil)).Elem(), GroupMemberGroupIdsMap{})
	pulumi.RegisterOutputType(GroupMemberGroupIdsOutput{})
	pulumi.RegisterOutputType(GroupMemberGroupIdsArrayOutput{})
	pulumi.RegisterOutputType(GroupMemberGroupIdsMapOutput{})
}
