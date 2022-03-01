// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package identity

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## Example Usage
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
// 		_, err := identity.NewEntity(ctx, "test", &identity.EntityArgs{
// 			Metadata: pulumi.StringMap{
// 				"foo": pulumi.String("bar"),
// 			},
// 			Policies: pulumi.StringArray{
// 				pulumi.String("test"),
// 			},
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
// Identity entity can be imported using the `id`, e.g.
//
// ```sh
//  $ pulumi import vault:identity/entity:Entity test "ae6f8ued-0f1a-9f6b-2915-1a2be20dc053"
// ```
type Entity struct {
	pulumi.CustomResourceState

	// True/false Is this entity currently disabled. Defaults to `false`
	Disabled pulumi.BoolPtrOutput `pulumi:"disabled"`
	// `false` by default. If set to `true`, this resource will ignore any policies return from Vault or specified in the resource. You can use `identity.EntityPolicies` to manage policies for this entity in a decoupled manner.
	ExternalPolicies pulumi.BoolPtrOutput `pulumi:"externalPolicies"`
	// A Map of additional metadata to associate with the user.
	Metadata pulumi.StringMapOutput `pulumi:"metadata"`
	// Name of the identity entity to create.
	Name pulumi.StringOutput `pulumi:"name"`
	// A list of policies to apply to the entity.
	Policies pulumi.StringArrayOutput `pulumi:"policies"`
}

// NewEntity registers a new resource with the given unique name, arguments, and options.
func NewEntity(ctx *pulumi.Context,
	name string, args *EntityArgs, opts ...pulumi.ResourceOption) (*Entity, error) {
	if args == nil {
		args = &EntityArgs{}
	}

	var resource Entity
	err := ctx.RegisterResource("vault:identity/entity:Entity", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetEntity gets an existing Entity resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEntity(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EntityState, opts ...pulumi.ResourceOption) (*Entity, error) {
	var resource Entity
	err := ctx.ReadResource("vault:identity/entity:Entity", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Entity resources.
type entityState struct {
	// True/false Is this entity currently disabled. Defaults to `false`
	Disabled *bool `pulumi:"disabled"`
	// `false` by default. If set to `true`, this resource will ignore any policies return from Vault or specified in the resource. You can use `identity.EntityPolicies` to manage policies for this entity in a decoupled manner.
	ExternalPolicies *bool `pulumi:"externalPolicies"`
	// A Map of additional metadata to associate with the user.
	Metadata map[string]string `pulumi:"metadata"`
	// Name of the identity entity to create.
	Name *string `pulumi:"name"`
	// A list of policies to apply to the entity.
	Policies []string `pulumi:"policies"`
}

type EntityState struct {
	// True/false Is this entity currently disabled. Defaults to `false`
	Disabled pulumi.BoolPtrInput
	// `false` by default. If set to `true`, this resource will ignore any policies return from Vault or specified in the resource. You can use `identity.EntityPolicies` to manage policies for this entity in a decoupled manner.
	ExternalPolicies pulumi.BoolPtrInput
	// A Map of additional metadata to associate with the user.
	Metadata pulumi.StringMapInput
	// Name of the identity entity to create.
	Name pulumi.StringPtrInput
	// A list of policies to apply to the entity.
	Policies pulumi.StringArrayInput
}

func (EntityState) ElementType() reflect.Type {
	return reflect.TypeOf((*entityState)(nil)).Elem()
}

type entityArgs struct {
	// True/false Is this entity currently disabled. Defaults to `false`
	Disabled *bool `pulumi:"disabled"`
	// `false` by default. If set to `true`, this resource will ignore any policies return from Vault or specified in the resource. You can use `identity.EntityPolicies` to manage policies for this entity in a decoupled manner.
	ExternalPolicies *bool `pulumi:"externalPolicies"`
	// A Map of additional metadata to associate with the user.
	Metadata map[string]string `pulumi:"metadata"`
	// Name of the identity entity to create.
	Name *string `pulumi:"name"`
	// A list of policies to apply to the entity.
	Policies []string `pulumi:"policies"`
}

// The set of arguments for constructing a Entity resource.
type EntityArgs struct {
	// True/false Is this entity currently disabled. Defaults to `false`
	Disabled pulumi.BoolPtrInput
	// `false` by default. If set to `true`, this resource will ignore any policies return from Vault or specified in the resource. You can use `identity.EntityPolicies` to manage policies for this entity in a decoupled manner.
	ExternalPolicies pulumi.BoolPtrInput
	// A Map of additional metadata to associate with the user.
	Metadata pulumi.StringMapInput
	// Name of the identity entity to create.
	Name pulumi.StringPtrInput
	// A list of policies to apply to the entity.
	Policies pulumi.StringArrayInput
}

func (EntityArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*entityArgs)(nil)).Elem()
}

type EntityInput interface {
	pulumi.Input

	ToEntityOutput() EntityOutput
	ToEntityOutputWithContext(ctx context.Context) EntityOutput
}

func (*Entity) ElementType() reflect.Type {
	return reflect.TypeOf((**Entity)(nil)).Elem()
}

func (i *Entity) ToEntityOutput() EntityOutput {
	return i.ToEntityOutputWithContext(context.Background())
}

func (i *Entity) ToEntityOutputWithContext(ctx context.Context) EntityOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EntityOutput)
}

// EntityArrayInput is an input type that accepts EntityArray and EntityArrayOutput values.
// You can construct a concrete instance of `EntityArrayInput` via:
//
//          EntityArray{ EntityArgs{...} }
type EntityArrayInput interface {
	pulumi.Input

	ToEntityArrayOutput() EntityArrayOutput
	ToEntityArrayOutputWithContext(context.Context) EntityArrayOutput
}

type EntityArray []EntityInput

func (EntityArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Entity)(nil)).Elem()
}

func (i EntityArray) ToEntityArrayOutput() EntityArrayOutput {
	return i.ToEntityArrayOutputWithContext(context.Background())
}

func (i EntityArray) ToEntityArrayOutputWithContext(ctx context.Context) EntityArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EntityArrayOutput)
}

// EntityMapInput is an input type that accepts EntityMap and EntityMapOutput values.
// You can construct a concrete instance of `EntityMapInput` via:
//
//          EntityMap{ "key": EntityArgs{...} }
type EntityMapInput interface {
	pulumi.Input

	ToEntityMapOutput() EntityMapOutput
	ToEntityMapOutputWithContext(context.Context) EntityMapOutput
}

type EntityMap map[string]EntityInput

func (EntityMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Entity)(nil)).Elem()
}

func (i EntityMap) ToEntityMapOutput() EntityMapOutput {
	return i.ToEntityMapOutputWithContext(context.Background())
}

func (i EntityMap) ToEntityMapOutputWithContext(ctx context.Context) EntityMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EntityMapOutput)
}

type EntityOutput struct{ *pulumi.OutputState }

func (EntityOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Entity)(nil)).Elem()
}

func (o EntityOutput) ToEntityOutput() EntityOutput {
	return o
}

func (o EntityOutput) ToEntityOutputWithContext(ctx context.Context) EntityOutput {
	return o
}

type EntityArrayOutput struct{ *pulumi.OutputState }

func (EntityArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Entity)(nil)).Elem()
}

func (o EntityArrayOutput) ToEntityArrayOutput() EntityArrayOutput {
	return o
}

func (o EntityArrayOutput) ToEntityArrayOutputWithContext(ctx context.Context) EntityArrayOutput {
	return o
}

func (o EntityArrayOutput) Index(i pulumi.IntInput) EntityOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Entity {
		return vs[0].([]*Entity)[vs[1].(int)]
	}).(EntityOutput)
}

type EntityMapOutput struct{ *pulumi.OutputState }

func (EntityMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Entity)(nil)).Elem()
}

func (o EntityMapOutput) ToEntityMapOutput() EntityMapOutput {
	return o
}

func (o EntityMapOutput) ToEntityMapOutputWithContext(ctx context.Context) EntityMapOutput {
	return o
}

func (o EntityMapOutput) MapIndex(k pulumi.StringInput) EntityOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Entity {
		return vs[0].(map[string]*Entity)[vs[1].(string)]
	}).(EntityOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*EntityInput)(nil)).Elem(), &Entity{})
	pulumi.RegisterInputType(reflect.TypeOf((*EntityArrayInput)(nil)).Elem(), EntityArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*EntityMapInput)(nil)).Elem(), EntityMap{})
	pulumi.RegisterOutputType(EntityOutput{})
	pulumi.RegisterOutputType(EntityArrayOutput{})
	pulumi.RegisterOutputType(EntityMapOutput{})
}
