// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package identity

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages member entities for an Identity Group for Vault. The [Identity secrets engine](https://www.vaultproject.io/docs/secrets/identity/index.html) is the identity management solution for Vault.
//
// ## Example Usage
// ### Exclusive Member Entities
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
// 		internal, err := identity.NewGroup(ctx, "internal", &identity.GroupArgs{
// 			Type:                    pulumi.String("internal"),
// 			ExternalMemberEntityIds: pulumi.Bool(true),
// 			Metadata: pulumi.StringMap{
// 				"version": pulumi.String("2"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		user, err := identity.NewEntity(ctx, "user", nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = identity.NewGroupMemberEntityIds(ctx, "members", &identity.GroupMemberEntityIdsArgs{
// 			Exclusive: pulumi.Bool(true),
// 			MemberEntityIds: pulumi.StringArray{
// 				user.ID(),
// 			},
// 			GroupId: internal.ID(),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### Non-exclusive Member Entities
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
// 		internal, err := identity.NewGroup(ctx, "internal", &identity.GroupArgs{
// 			Type:                    pulumi.String("internal"),
// 			ExternalMemberEntityIds: pulumi.Bool(true),
// 			Metadata: pulumi.StringMap{
// 				"version": pulumi.String("2"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		testUser, err := identity.NewEntity(ctx, "testUser", nil)
// 		if err != nil {
// 			return err
// 		}
// 		secondTestUser, err := identity.NewEntity(ctx, "secondTestUser", nil)
// 		if err != nil {
// 			return err
// 		}
// 		devUser, err := identity.NewEntity(ctx, "devUser", nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = identity.NewGroupMemberEntityIds(ctx, "test", &identity.GroupMemberEntityIdsArgs{
// 			MemberEntityIds: pulumi.StringArray{
// 				testUser.ID(),
// 				secondTestUser.ID(),
// 			},
// 			Exclusive: pulumi.Bool(false),
// 			GroupId:   internal.ID(),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = identity.NewGroupMemberEntityIds(ctx, "others", &identity.GroupMemberEntityIdsArgs{
// 			MemberEntityIds: pulumi.StringArray{
// 				devUser.ID(),
// 			},
// 			Exclusive: pulumi.Bool(false),
// 			GroupId:   internal.ID(),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type GroupMemberEntityIds struct {
	pulumi.CustomResourceState

	// Defaults to `true`.
	Exclusive pulumi.BoolPtrOutput `pulumi:"exclusive"`
	// Group ID to assign member entities to.
	GroupId pulumi.StringOutput `pulumi:"groupId"`
	// The name of the group that are assigned the member entities.
	GroupName pulumi.StringOutput `pulumi:"groupName"`
	// List of member entities that belong to the group
	MemberEntityIds pulumi.StringArrayOutput `pulumi:"memberEntityIds"`
}

// NewGroupMemberEntityIds registers a new resource with the given unique name, arguments, and options.
func NewGroupMemberEntityIds(ctx *pulumi.Context,
	name string, args *GroupMemberEntityIdsArgs, opts ...pulumi.ResourceOption) (*GroupMemberEntityIds, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.GroupId == nil {
		return nil, errors.New("invalid value for required argument 'GroupId'")
	}
	var resource GroupMemberEntityIds
	err := ctx.RegisterResource("vault:identity/groupMemberEntityIds:GroupMemberEntityIds", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetGroupMemberEntityIds gets an existing GroupMemberEntityIds resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGroupMemberEntityIds(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GroupMemberEntityIdsState, opts ...pulumi.ResourceOption) (*GroupMemberEntityIds, error) {
	var resource GroupMemberEntityIds
	err := ctx.ReadResource("vault:identity/groupMemberEntityIds:GroupMemberEntityIds", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering GroupMemberEntityIds resources.
type groupMemberEntityIdsState struct {
	// Defaults to `true`.
	Exclusive *bool `pulumi:"exclusive"`
	// Group ID to assign member entities to.
	GroupId *string `pulumi:"groupId"`
	// The name of the group that are assigned the member entities.
	GroupName *string `pulumi:"groupName"`
	// List of member entities that belong to the group
	MemberEntityIds []string `pulumi:"memberEntityIds"`
}

type GroupMemberEntityIdsState struct {
	// Defaults to `true`.
	Exclusive pulumi.BoolPtrInput
	// Group ID to assign member entities to.
	GroupId pulumi.StringPtrInput
	// The name of the group that are assigned the member entities.
	GroupName pulumi.StringPtrInput
	// List of member entities that belong to the group
	MemberEntityIds pulumi.StringArrayInput
}

func (GroupMemberEntityIdsState) ElementType() reflect.Type {
	return reflect.TypeOf((*groupMemberEntityIdsState)(nil)).Elem()
}

type groupMemberEntityIdsArgs struct {
	// Defaults to `true`.
	Exclusive *bool `pulumi:"exclusive"`
	// Group ID to assign member entities to.
	GroupId string `pulumi:"groupId"`
	// List of member entities that belong to the group
	MemberEntityIds []string `pulumi:"memberEntityIds"`
}

// The set of arguments for constructing a GroupMemberEntityIds resource.
type GroupMemberEntityIdsArgs struct {
	// Defaults to `true`.
	Exclusive pulumi.BoolPtrInput
	// Group ID to assign member entities to.
	GroupId pulumi.StringInput
	// List of member entities that belong to the group
	MemberEntityIds pulumi.StringArrayInput
}

func (GroupMemberEntityIdsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*groupMemberEntityIdsArgs)(nil)).Elem()
}

type GroupMemberEntityIdsInput interface {
	pulumi.Input

	ToGroupMemberEntityIdsOutput() GroupMemberEntityIdsOutput
	ToGroupMemberEntityIdsOutputWithContext(ctx context.Context) GroupMemberEntityIdsOutput
}

func (*GroupMemberEntityIds) ElementType() reflect.Type {
	return reflect.TypeOf((**GroupMemberEntityIds)(nil)).Elem()
}

func (i *GroupMemberEntityIds) ToGroupMemberEntityIdsOutput() GroupMemberEntityIdsOutput {
	return i.ToGroupMemberEntityIdsOutputWithContext(context.Background())
}

func (i *GroupMemberEntityIds) ToGroupMemberEntityIdsOutputWithContext(ctx context.Context) GroupMemberEntityIdsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GroupMemberEntityIdsOutput)
}

// GroupMemberEntityIdsArrayInput is an input type that accepts GroupMemberEntityIdsArray and GroupMemberEntityIdsArrayOutput values.
// You can construct a concrete instance of `GroupMemberEntityIdsArrayInput` via:
//
//          GroupMemberEntityIdsArray{ GroupMemberEntityIdsArgs{...} }
type GroupMemberEntityIdsArrayInput interface {
	pulumi.Input

	ToGroupMemberEntityIdsArrayOutput() GroupMemberEntityIdsArrayOutput
	ToGroupMemberEntityIdsArrayOutputWithContext(context.Context) GroupMemberEntityIdsArrayOutput
}

type GroupMemberEntityIdsArray []GroupMemberEntityIdsInput

func (GroupMemberEntityIdsArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*GroupMemberEntityIds)(nil)).Elem()
}

func (i GroupMemberEntityIdsArray) ToGroupMemberEntityIdsArrayOutput() GroupMemberEntityIdsArrayOutput {
	return i.ToGroupMemberEntityIdsArrayOutputWithContext(context.Background())
}

func (i GroupMemberEntityIdsArray) ToGroupMemberEntityIdsArrayOutputWithContext(ctx context.Context) GroupMemberEntityIdsArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GroupMemberEntityIdsArrayOutput)
}

// GroupMemberEntityIdsMapInput is an input type that accepts GroupMemberEntityIdsMap and GroupMemberEntityIdsMapOutput values.
// You can construct a concrete instance of `GroupMemberEntityIdsMapInput` via:
//
//          GroupMemberEntityIdsMap{ "key": GroupMemberEntityIdsArgs{...} }
type GroupMemberEntityIdsMapInput interface {
	pulumi.Input

	ToGroupMemberEntityIdsMapOutput() GroupMemberEntityIdsMapOutput
	ToGroupMemberEntityIdsMapOutputWithContext(context.Context) GroupMemberEntityIdsMapOutput
}

type GroupMemberEntityIdsMap map[string]GroupMemberEntityIdsInput

func (GroupMemberEntityIdsMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*GroupMemberEntityIds)(nil)).Elem()
}

func (i GroupMemberEntityIdsMap) ToGroupMemberEntityIdsMapOutput() GroupMemberEntityIdsMapOutput {
	return i.ToGroupMemberEntityIdsMapOutputWithContext(context.Background())
}

func (i GroupMemberEntityIdsMap) ToGroupMemberEntityIdsMapOutputWithContext(ctx context.Context) GroupMemberEntityIdsMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GroupMemberEntityIdsMapOutput)
}

type GroupMemberEntityIdsOutput struct{ *pulumi.OutputState }

func (GroupMemberEntityIdsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GroupMemberEntityIds)(nil)).Elem()
}

func (o GroupMemberEntityIdsOutput) ToGroupMemberEntityIdsOutput() GroupMemberEntityIdsOutput {
	return o
}

func (o GroupMemberEntityIdsOutput) ToGroupMemberEntityIdsOutputWithContext(ctx context.Context) GroupMemberEntityIdsOutput {
	return o
}

type GroupMemberEntityIdsArrayOutput struct{ *pulumi.OutputState }

func (GroupMemberEntityIdsArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*GroupMemberEntityIds)(nil)).Elem()
}

func (o GroupMemberEntityIdsArrayOutput) ToGroupMemberEntityIdsArrayOutput() GroupMemberEntityIdsArrayOutput {
	return o
}

func (o GroupMemberEntityIdsArrayOutput) ToGroupMemberEntityIdsArrayOutputWithContext(ctx context.Context) GroupMemberEntityIdsArrayOutput {
	return o
}

func (o GroupMemberEntityIdsArrayOutput) Index(i pulumi.IntInput) GroupMemberEntityIdsOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *GroupMemberEntityIds {
		return vs[0].([]*GroupMemberEntityIds)[vs[1].(int)]
	}).(GroupMemberEntityIdsOutput)
}

type GroupMemberEntityIdsMapOutput struct{ *pulumi.OutputState }

func (GroupMemberEntityIdsMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*GroupMemberEntityIds)(nil)).Elem()
}

func (o GroupMemberEntityIdsMapOutput) ToGroupMemberEntityIdsMapOutput() GroupMemberEntityIdsMapOutput {
	return o
}

func (o GroupMemberEntityIdsMapOutput) ToGroupMemberEntityIdsMapOutputWithContext(ctx context.Context) GroupMemberEntityIdsMapOutput {
	return o
}

func (o GroupMemberEntityIdsMapOutput) MapIndex(k pulumi.StringInput) GroupMemberEntityIdsOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *GroupMemberEntityIds {
		return vs[0].(map[string]*GroupMemberEntityIds)[vs[1].(string)]
	}).(GroupMemberEntityIdsOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*GroupMemberEntityIdsInput)(nil)).Elem(), &GroupMemberEntityIds{})
	pulumi.RegisterInputType(reflect.TypeOf((*GroupMemberEntityIdsArrayInput)(nil)).Elem(), GroupMemberEntityIdsArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*GroupMemberEntityIdsMapInput)(nil)).Elem(), GroupMemberEntityIdsMap{})
	pulumi.RegisterOutputType(GroupMemberEntityIdsOutput{})
	pulumi.RegisterOutputType(GroupMemberEntityIdsArrayOutput{})
	pulumi.RegisterOutputType(GroupMemberEntityIdsMapOutput{})
}
