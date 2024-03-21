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

// Manages policies for an Identity Group for Vault. The [Identity secrets engine](https://www.vaultproject.io/docs/secrets/identity/index.html) is the identity management solution for Vault.
//
// ## Example Usage
//
// ### Exclusive Policies
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
//				Type:             pulumi.String("internal"),
//				ExternalPolicies: pulumi.Bool(true),
//				Metadata: pulumi.StringMap{
//					"version": pulumi.String("2"),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			_, err = identity.NewGroupPolicies(ctx, "policies", &identity.GroupPoliciesArgs{
//				Policies: pulumi.StringArray{
//					pulumi.String("default"),
//					pulumi.String("test"),
//				},
//				Exclusive: pulumi.Bool(true),
//				GroupId:   internal.ID(),
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
// ### Non-exclusive Policies
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
//				Type:             pulumi.String("internal"),
//				ExternalPolicies: pulumi.Bool(true),
//				Metadata: pulumi.StringMap{
//					"version": pulumi.String("2"),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			_, err = identity.NewGroupPolicies(ctx, "default", &identity.GroupPoliciesArgs{
//				Policies: pulumi.StringArray{
//					pulumi.String("default"),
//					pulumi.String("test"),
//				},
//				Exclusive: pulumi.Bool(false),
//				GroupId:   internal.ID(),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = identity.NewGroupPolicies(ctx, "others", &identity.GroupPoliciesArgs{
//				Policies: pulumi.StringArray{
//					pulumi.String("others"),
//				},
//				Exclusive: pulumi.Bool(false),
//				GroupId:   internal.ID(),
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
type GroupPolicies struct {
	pulumi.CustomResourceState

	// Defaults to `true`.
	//
	// If `true`, this resource will take exclusive control of the policies assigned to the group and will set it equal to what is specified in the resource.
	//
	// If set to `false`, this resource will simply ensure that the policies specified in the resource are present in the group. When destroying the resource, the resource will ensure that the policies specified in the resource are removed.
	Exclusive pulumi.BoolPtrOutput `pulumi:"exclusive"`
	// Group ID to assign policies to.
	GroupId pulumi.StringOutput `pulumi:"groupId"`
	// The name of the group that are assigned the policies.
	GroupName pulumi.StringOutput `pulumi:"groupName"`
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
	// *Available only for Vault Enterprise*.
	Namespace pulumi.StringPtrOutput `pulumi:"namespace"`
	// List of policies to assign to the group
	Policies pulumi.StringArrayOutput `pulumi:"policies"`
}

// NewGroupPolicies registers a new resource with the given unique name, arguments, and options.
func NewGroupPolicies(ctx *pulumi.Context,
	name string, args *GroupPoliciesArgs, opts ...pulumi.ResourceOption) (*GroupPolicies, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.GroupId == nil {
		return nil, errors.New("invalid value for required argument 'GroupId'")
	}
	if args.Policies == nil {
		return nil, errors.New("invalid value for required argument 'Policies'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource GroupPolicies
	err := ctx.RegisterResource("vault:identity/groupPolicies:GroupPolicies", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetGroupPolicies gets an existing GroupPolicies resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGroupPolicies(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GroupPoliciesState, opts ...pulumi.ResourceOption) (*GroupPolicies, error) {
	var resource GroupPolicies
	err := ctx.ReadResource("vault:identity/groupPolicies:GroupPolicies", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering GroupPolicies resources.
type groupPoliciesState struct {
	// Defaults to `true`.
	//
	// If `true`, this resource will take exclusive control of the policies assigned to the group and will set it equal to what is specified in the resource.
	//
	// If set to `false`, this resource will simply ensure that the policies specified in the resource are present in the group. When destroying the resource, the resource will ensure that the policies specified in the resource are removed.
	Exclusive *bool `pulumi:"exclusive"`
	// Group ID to assign policies to.
	GroupId *string `pulumi:"groupId"`
	// The name of the group that are assigned the policies.
	GroupName *string `pulumi:"groupName"`
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
	// *Available only for Vault Enterprise*.
	Namespace *string `pulumi:"namespace"`
	// List of policies to assign to the group
	Policies []string `pulumi:"policies"`
}

type GroupPoliciesState struct {
	// Defaults to `true`.
	//
	// If `true`, this resource will take exclusive control of the policies assigned to the group and will set it equal to what is specified in the resource.
	//
	// If set to `false`, this resource will simply ensure that the policies specified in the resource are present in the group. When destroying the resource, the resource will ensure that the policies specified in the resource are removed.
	Exclusive pulumi.BoolPtrInput
	// Group ID to assign policies to.
	GroupId pulumi.StringPtrInput
	// The name of the group that are assigned the policies.
	GroupName pulumi.StringPtrInput
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
	// *Available only for Vault Enterprise*.
	Namespace pulumi.StringPtrInput
	// List of policies to assign to the group
	Policies pulumi.StringArrayInput
}

func (GroupPoliciesState) ElementType() reflect.Type {
	return reflect.TypeOf((*groupPoliciesState)(nil)).Elem()
}

type groupPoliciesArgs struct {
	// Defaults to `true`.
	//
	// If `true`, this resource will take exclusive control of the policies assigned to the group and will set it equal to what is specified in the resource.
	//
	// If set to `false`, this resource will simply ensure that the policies specified in the resource are present in the group. When destroying the resource, the resource will ensure that the policies specified in the resource are removed.
	Exclusive *bool `pulumi:"exclusive"`
	// Group ID to assign policies to.
	GroupId string `pulumi:"groupId"`
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
	// *Available only for Vault Enterprise*.
	Namespace *string `pulumi:"namespace"`
	// List of policies to assign to the group
	Policies []string `pulumi:"policies"`
}

// The set of arguments for constructing a GroupPolicies resource.
type GroupPoliciesArgs struct {
	// Defaults to `true`.
	//
	// If `true`, this resource will take exclusive control of the policies assigned to the group and will set it equal to what is specified in the resource.
	//
	// If set to `false`, this resource will simply ensure that the policies specified in the resource are present in the group. When destroying the resource, the resource will ensure that the policies specified in the resource are removed.
	Exclusive pulumi.BoolPtrInput
	// Group ID to assign policies to.
	GroupId pulumi.StringInput
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
	// *Available only for Vault Enterprise*.
	Namespace pulumi.StringPtrInput
	// List of policies to assign to the group
	Policies pulumi.StringArrayInput
}

func (GroupPoliciesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*groupPoliciesArgs)(nil)).Elem()
}

type GroupPoliciesInput interface {
	pulumi.Input

	ToGroupPoliciesOutput() GroupPoliciesOutput
	ToGroupPoliciesOutputWithContext(ctx context.Context) GroupPoliciesOutput
}

func (*GroupPolicies) ElementType() reflect.Type {
	return reflect.TypeOf((**GroupPolicies)(nil)).Elem()
}

func (i *GroupPolicies) ToGroupPoliciesOutput() GroupPoliciesOutput {
	return i.ToGroupPoliciesOutputWithContext(context.Background())
}

func (i *GroupPolicies) ToGroupPoliciesOutputWithContext(ctx context.Context) GroupPoliciesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GroupPoliciesOutput)
}

// GroupPoliciesArrayInput is an input type that accepts GroupPoliciesArray and GroupPoliciesArrayOutput values.
// You can construct a concrete instance of `GroupPoliciesArrayInput` via:
//
//	GroupPoliciesArray{ GroupPoliciesArgs{...} }
type GroupPoliciesArrayInput interface {
	pulumi.Input

	ToGroupPoliciesArrayOutput() GroupPoliciesArrayOutput
	ToGroupPoliciesArrayOutputWithContext(context.Context) GroupPoliciesArrayOutput
}

type GroupPoliciesArray []GroupPoliciesInput

func (GroupPoliciesArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*GroupPolicies)(nil)).Elem()
}

func (i GroupPoliciesArray) ToGroupPoliciesArrayOutput() GroupPoliciesArrayOutput {
	return i.ToGroupPoliciesArrayOutputWithContext(context.Background())
}

func (i GroupPoliciesArray) ToGroupPoliciesArrayOutputWithContext(ctx context.Context) GroupPoliciesArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GroupPoliciesArrayOutput)
}

// GroupPoliciesMapInput is an input type that accepts GroupPoliciesMap and GroupPoliciesMapOutput values.
// You can construct a concrete instance of `GroupPoliciesMapInput` via:
//
//	GroupPoliciesMap{ "key": GroupPoliciesArgs{...} }
type GroupPoliciesMapInput interface {
	pulumi.Input

	ToGroupPoliciesMapOutput() GroupPoliciesMapOutput
	ToGroupPoliciesMapOutputWithContext(context.Context) GroupPoliciesMapOutput
}

type GroupPoliciesMap map[string]GroupPoliciesInput

func (GroupPoliciesMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*GroupPolicies)(nil)).Elem()
}

func (i GroupPoliciesMap) ToGroupPoliciesMapOutput() GroupPoliciesMapOutput {
	return i.ToGroupPoliciesMapOutputWithContext(context.Background())
}

func (i GroupPoliciesMap) ToGroupPoliciesMapOutputWithContext(ctx context.Context) GroupPoliciesMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GroupPoliciesMapOutput)
}

type GroupPoliciesOutput struct{ *pulumi.OutputState }

func (GroupPoliciesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GroupPolicies)(nil)).Elem()
}

func (o GroupPoliciesOutput) ToGroupPoliciesOutput() GroupPoliciesOutput {
	return o
}

func (o GroupPoliciesOutput) ToGroupPoliciesOutputWithContext(ctx context.Context) GroupPoliciesOutput {
	return o
}

// Defaults to `true`.
//
// If `true`, this resource will take exclusive control of the policies assigned to the group and will set it equal to what is specified in the resource.
//
// If set to `false`, this resource will simply ensure that the policies specified in the resource are present in the group. When destroying the resource, the resource will ensure that the policies specified in the resource are removed.
func (o GroupPoliciesOutput) Exclusive() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *GroupPolicies) pulumi.BoolPtrOutput { return v.Exclusive }).(pulumi.BoolPtrOutput)
}

// Group ID to assign policies to.
func (o GroupPoliciesOutput) GroupId() pulumi.StringOutput {
	return o.ApplyT(func(v *GroupPolicies) pulumi.StringOutput { return v.GroupId }).(pulumi.StringOutput)
}

// The name of the group that are assigned the policies.
func (o GroupPoliciesOutput) GroupName() pulumi.StringOutput {
	return o.ApplyT(func(v *GroupPolicies) pulumi.StringOutput { return v.GroupName }).(pulumi.StringOutput)
}

// The namespace to provision the resource in.
// The value should not contain leading or trailing forward slashes.
// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
// *Available only for Vault Enterprise*.
func (o GroupPoliciesOutput) Namespace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GroupPolicies) pulumi.StringPtrOutput { return v.Namespace }).(pulumi.StringPtrOutput)
}

// List of policies to assign to the group
func (o GroupPoliciesOutput) Policies() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *GroupPolicies) pulumi.StringArrayOutput { return v.Policies }).(pulumi.StringArrayOutput)
}

type GroupPoliciesArrayOutput struct{ *pulumi.OutputState }

func (GroupPoliciesArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*GroupPolicies)(nil)).Elem()
}

func (o GroupPoliciesArrayOutput) ToGroupPoliciesArrayOutput() GroupPoliciesArrayOutput {
	return o
}

func (o GroupPoliciesArrayOutput) ToGroupPoliciesArrayOutputWithContext(ctx context.Context) GroupPoliciesArrayOutput {
	return o
}

func (o GroupPoliciesArrayOutput) Index(i pulumi.IntInput) GroupPoliciesOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *GroupPolicies {
		return vs[0].([]*GroupPolicies)[vs[1].(int)]
	}).(GroupPoliciesOutput)
}

type GroupPoliciesMapOutput struct{ *pulumi.OutputState }

func (GroupPoliciesMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*GroupPolicies)(nil)).Elem()
}

func (o GroupPoliciesMapOutput) ToGroupPoliciesMapOutput() GroupPoliciesMapOutput {
	return o
}

func (o GroupPoliciesMapOutput) ToGroupPoliciesMapOutputWithContext(ctx context.Context) GroupPoliciesMapOutput {
	return o
}

func (o GroupPoliciesMapOutput) MapIndex(k pulumi.StringInput) GroupPoliciesOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *GroupPolicies {
		return vs[0].(map[string]*GroupPolicies)[vs[1].(string)]
	}).(GroupPoliciesOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*GroupPoliciesInput)(nil)).Elem(), &GroupPolicies{})
	pulumi.RegisterInputType(reflect.TypeOf((*GroupPoliciesArrayInput)(nil)).Elem(), GroupPoliciesArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*GroupPoliciesMapInput)(nil)).Elem(), GroupPoliciesMap{})
	pulumi.RegisterOutputType(GroupPoliciesOutput{})
	pulumi.RegisterOutputType(GroupPoliciesArrayOutput{})
	pulumi.RegisterOutputType(GroupPoliciesMapOutput{})
}
