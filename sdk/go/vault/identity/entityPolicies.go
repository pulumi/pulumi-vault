// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package identity

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-vault/sdk/v5/go/vault/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages policies for an Identity Entity for Vault. The [Identity secrets engine](https://www.vaultproject.io/docs/secrets/identity/index.html) is the identity management solution for Vault.
//
// ## Example Usage
// ### Exclusive Policies
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-vault/sdk/v5/go/vault/identity"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			entity, err := identity.NewEntity(ctx, "entity", &identity.EntityArgs{
//				ExternalPolicies: pulumi.Bool(true),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = identity.NewEntityPolicies(ctx, "policies", &identity.EntityPoliciesArgs{
//				Policies: pulumi.StringArray{
//					pulumi.String("default"),
//					pulumi.String("test"),
//				},
//				Exclusive: pulumi.Bool(true),
//				EntityId:  entity.ID(),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// ### Non-exclusive Policies
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-vault/sdk/v5/go/vault/identity"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			entity, err := identity.NewEntity(ctx, "entity", &identity.EntityArgs{
//				ExternalPolicies: pulumi.Bool(true),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = identity.NewEntityPolicies(ctx, "default", &identity.EntityPoliciesArgs{
//				Policies: pulumi.StringArray{
//					pulumi.String("default"),
//					pulumi.String("test"),
//				},
//				Exclusive: pulumi.Bool(false),
//				EntityId:  entity.ID(),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = identity.NewEntityPolicies(ctx, "others", &identity.EntityPoliciesArgs{
//				Policies: pulumi.StringArray{
//					pulumi.String("others"),
//				},
//				Exclusive: pulumi.Bool(false),
//				EntityId:  entity.ID(),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
type EntityPolicies struct {
	pulumi.CustomResourceState

	// Entity ID to assign policies to.
	EntityId pulumi.StringOutput `pulumi:"entityId"`
	// The name of the entity that are assigned the policies.
	EntityName pulumi.StringOutput `pulumi:"entityName"`
	// Defaults to `true`.
	//
	// If `true`, this resource will take exclusive control of the policies assigned to the entity and will set it equal to what is specified in the resource.
	//
	// If set to `false`, this resource will simply ensure that the policies specified in the resource are present in the entity. When destroying the resource, the resource will ensure that the policies specified in the resource are removed.
	Exclusive pulumi.BoolPtrOutput `pulumi:"exclusive"`
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
	// *Available only for Vault Enterprise*.
	Namespace pulumi.StringPtrOutput `pulumi:"namespace"`
	// List of policies to assign to the entity
	Policies pulumi.StringArrayOutput `pulumi:"policies"`
}

// NewEntityPolicies registers a new resource with the given unique name, arguments, and options.
func NewEntityPolicies(ctx *pulumi.Context,
	name string, args *EntityPoliciesArgs, opts ...pulumi.ResourceOption) (*EntityPolicies, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.EntityId == nil {
		return nil, errors.New("invalid value for required argument 'EntityId'")
	}
	if args.Policies == nil {
		return nil, errors.New("invalid value for required argument 'Policies'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource EntityPolicies
	err := ctx.RegisterResource("vault:identity/entityPolicies:EntityPolicies", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetEntityPolicies gets an existing EntityPolicies resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEntityPolicies(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EntityPoliciesState, opts ...pulumi.ResourceOption) (*EntityPolicies, error) {
	var resource EntityPolicies
	err := ctx.ReadResource("vault:identity/entityPolicies:EntityPolicies", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering EntityPolicies resources.
type entityPoliciesState struct {
	// Entity ID to assign policies to.
	EntityId *string `pulumi:"entityId"`
	// The name of the entity that are assigned the policies.
	EntityName *string `pulumi:"entityName"`
	// Defaults to `true`.
	//
	// If `true`, this resource will take exclusive control of the policies assigned to the entity and will set it equal to what is specified in the resource.
	//
	// If set to `false`, this resource will simply ensure that the policies specified in the resource are present in the entity. When destroying the resource, the resource will ensure that the policies specified in the resource are removed.
	Exclusive *bool `pulumi:"exclusive"`
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
	// *Available only for Vault Enterprise*.
	Namespace *string `pulumi:"namespace"`
	// List of policies to assign to the entity
	Policies []string `pulumi:"policies"`
}

type EntityPoliciesState struct {
	// Entity ID to assign policies to.
	EntityId pulumi.StringPtrInput
	// The name of the entity that are assigned the policies.
	EntityName pulumi.StringPtrInput
	// Defaults to `true`.
	//
	// If `true`, this resource will take exclusive control of the policies assigned to the entity and will set it equal to what is specified in the resource.
	//
	// If set to `false`, this resource will simply ensure that the policies specified in the resource are present in the entity. When destroying the resource, the resource will ensure that the policies specified in the resource are removed.
	Exclusive pulumi.BoolPtrInput
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
	// *Available only for Vault Enterprise*.
	Namespace pulumi.StringPtrInput
	// List of policies to assign to the entity
	Policies pulumi.StringArrayInput
}

func (EntityPoliciesState) ElementType() reflect.Type {
	return reflect.TypeOf((*entityPoliciesState)(nil)).Elem()
}

type entityPoliciesArgs struct {
	// Entity ID to assign policies to.
	EntityId string `pulumi:"entityId"`
	// Defaults to `true`.
	//
	// If `true`, this resource will take exclusive control of the policies assigned to the entity and will set it equal to what is specified in the resource.
	//
	// If set to `false`, this resource will simply ensure that the policies specified in the resource are present in the entity. When destroying the resource, the resource will ensure that the policies specified in the resource are removed.
	Exclusive *bool `pulumi:"exclusive"`
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
	// *Available only for Vault Enterprise*.
	Namespace *string `pulumi:"namespace"`
	// List of policies to assign to the entity
	Policies []string `pulumi:"policies"`
}

// The set of arguments for constructing a EntityPolicies resource.
type EntityPoliciesArgs struct {
	// Entity ID to assign policies to.
	EntityId pulumi.StringInput
	// Defaults to `true`.
	//
	// If `true`, this resource will take exclusive control of the policies assigned to the entity and will set it equal to what is specified in the resource.
	//
	// If set to `false`, this resource will simply ensure that the policies specified in the resource are present in the entity. When destroying the resource, the resource will ensure that the policies specified in the resource are removed.
	Exclusive pulumi.BoolPtrInput
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
	// *Available only for Vault Enterprise*.
	Namespace pulumi.StringPtrInput
	// List of policies to assign to the entity
	Policies pulumi.StringArrayInput
}

func (EntityPoliciesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*entityPoliciesArgs)(nil)).Elem()
}

type EntityPoliciesInput interface {
	pulumi.Input

	ToEntityPoliciesOutput() EntityPoliciesOutput
	ToEntityPoliciesOutputWithContext(ctx context.Context) EntityPoliciesOutput
}

func (*EntityPolicies) ElementType() reflect.Type {
	return reflect.TypeOf((**EntityPolicies)(nil)).Elem()
}

func (i *EntityPolicies) ToEntityPoliciesOutput() EntityPoliciesOutput {
	return i.ToEntityPoliciesOutputWithContext(context.Background())
}

func (i *EntityPolicies) ToEntityPoliciesOutputWithContext(ctx context.Context) EntityPoliciesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EntityPoliciesOutput)
}

// EntityPoliciesArrayInput is an input type that accepts EntityPoliciesArray and EntityPoliciesArrayOutput values.
// You can construct a concrete instance of `EntityPoliciesArrayInput` via:
//
//	EntityPoliciesArray{ EntityPoliciesArgs{...} }
type EntityPoliciesArrayInput interface {
	pulumi.Input

	ToEntityPoliciesArrayOutput() EntityPoliciesArrayOutput
	ToEntityPoliciesArrayOutputWithContext(context.Context) EntityPoliciesArrayOutput
}

type EntityPoliciesArray []EntityPoliciesInput

func (EntityPoliciesArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*EntityPolicies)(nil)).Elem()
}

func (i EntityPoliciesArray) ToEntityPoliciesArrayOutput() EntityPoliciesArrayOutput {
	return i.ToEntityPoliciesArrayOutputWithContext(context.Background())
}

func (i EntityPoliciesArray) ToEntityPoliciesArrayOutputWithContext(ctx context.Context) EntityPoliciesArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EntityPoliciesArrayOutput)
}

// EntityPoliciesMapInput is an input type that accepts EntityPoliciesMap and EntityPoliciesMapOutput values.
// You can construct a concrete instance of `EntityPoliciesMapInput` via:
//
//	EntityPoliciesMap{ "key": EntityPoliciesArgs{...} }
type EntityPoliciesMapInput interface {
	pulumi.Input

	ToEntityPoliciesMapOutput() EntityPoliciesMapOutput
	ToEntityPoliciesMapOutputWithContext(context.Context) EntityPoliciesMapOutput
}

type EntityPoliciesMap map[string]EntityPoliciesInput

func (EntityPoliciesMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*EntityPolicies)(nil)).Elem()
}

func (i EntityPoliciesMap) ToEntityPoliciesMapOutput() EntityPoliciesMapOutput {
	return i.ToEntityPoliciesMapOutputWithContext(context.Background())
}

func (i EntityPoliciesMap) ToEntityPoliciesMapOutputWithContext(ctx context.Context) EntityPoliciesMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EntityPoliciesMapOutput)
}

type EntityPoliciesOutput struct{ *pulumi.OutputState }

func (EntityPoliciesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EntityPolicies)(nil)).Elem()
}

func (o EntityPoliciesOutput) ToEntityPoliciesOutput() EntityPoliciesOutput {
	return o
}

func (o EntityPoliciesOutput) ToEntityPoliciesOutputWithContext(ctx context.Context) EntityPoliciesOutput {
	return o
}

// Entity ID to assign policies to.
func (o EntityPoliciesOutput) EntityId() pulumi.StringOutput {
	return o.ApplyT(func(v *EntityPolicies) pulumi.StringOutput { return v.EntityId }).(pulumi.StringOutput)
}

// The name of the entity that are assigned the policies.
func (o EntityPoliciesOutput) EntityName() pulumi.StringOutput {
	return o.ApplyT(func(v *EntityPolicies) pulumi.StringOutput { return v.EntityName }).(pulumi.StringOutput)
}

// Defaults to `true`.
//
// If `true`, this resource will take exclusive control of the policies assigned to the entity and will set it equal to what is specified in the resource.
//
// If set to `false`, this resource will simply ensure that the policies specified in the resource are present in the entity. When destroying the resource, the resource will ensure that the policies specified in the resource are removed.
func (o EntityPoliciesOutput) Exclusive() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *EntityPolicies) pulumi.BoolPtrOutput { return v.Exclusive }).(pulumi.BoolPtrOutput)
}

// The namespace to provision the resource in.
// The value should not contain leading or trailing forward slashes.
// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
// *Available only for Vault Enterprise*.
func (o EntityPoliciesOutput) Namespace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EntityPolicies) pulumi.StringPtrOutput { return v.Namespace }).(pulumi.StringPtrOutput)
}

// List of policies to assign to the entity
func (o EntityPoliciesOutput) Policies() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *EntityPolicies) pulumi.StringArrayOutput { return v.Policies }).(pulumi.StringArrayOutput)
}

type EntityPoliciesArrayOutput struct{ *pulumi.OutputState }

func (EntityPoliciesArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*EntityPolicies)(nil)).Elem()
}

func (o EntityPoliciesArrayOutput) ToEntityPoliciesArrayOutput() EntityPoliciesArrayOutput {
	return o
}

func (o EntityPoliciesArrayOutput) ToEntityPoliciesArrayOutputWithContext(ctx context.Context) EntityPoliciesArrayOutput {
	return o
}

func (o EntityPoliciesArrayOutput) Index(i pulumi.IntInput) EntityPoliciesOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *EntityPolicies {
		return vs[0].([]*EntityPolicies)[vs[1].(int)]
	}).(EntityPoliciesOutput)
}

type EntityPoliciesMapOutput struct{ *pulumi.OutputState }

func (EntityPoliciesMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*EntityPolicies)(nil)).Elem()
}

func (o EntityPoliciesMapOutput) ToEntityPoliciesMapOutput() EntityPoliciesMapOutput {
	return o
}

func (o EntityPoliciesMapOutput) ToEntityPoliciesMapOutputWithContext(ctx context.Context) EntityPoliciesMapOutput {
	return o
}

func (o EntityPoliciesMapOutput) MapIndex(k pulumi.StringInput) EntityPoliciesOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *EntityPolicies {
		return vs[0].(map[string]*EntityPolicies)[vs[1].(string)]
	}).(EntityPoliciesOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*EntityPoliciesInput)(nil)).Elem(), &EntityPolicies{})
	pulumi.RegisterInputType(reflect.TypeOf((*EntityPoliciesArrayInput)(nil)).Elem(), EntityPoliciesArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*EntityPoliciesMapInput)(nil)).Elem(), EntityPoliciesMap{})
	pulumi.RegisterOutputType(EntityPoliciesOutput{})
	pulumi.RegisterOutputType(EntityPoliciesArrayOutput{})
	pulumi.RegisterOutputType(EntityPoliciesMapOutput{})
}
