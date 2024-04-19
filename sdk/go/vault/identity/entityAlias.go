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

// ## Example Usage
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
//			_, err := identity.NewEntityAlias(ctx, "test", &identity.EntityAliasArgs{
//				Name:          pulumi.String("user_1"),
//				MountAccessor: pulumi.String("token_1f2bd5"),
//				CanonicalId:   pulumi.String("49877D63-07AD-4B85-BDA8-B61626C477E8"),
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
// ## Import
//
// Identity entity alias can be imported using the `id`, e.g.
//
// ```sh
// $ pulumi import vault:identity/entityAlias:EntityAlias test "3856fb4d-3c91-dcaf-2401-68f446796bfb"
// ```
type EntityAlias struct {
	pulumi.CustomResourceState

	// Entity ID to which this alias belongs to.
	CanonicalId pulumi.StringOutput `pulumi:"canonicalId"`
	// Custom metadata to be associated with this alias.
	CustomMetadata pulumi.StringMapOutput `pulumi:"customMetadata"`
	// Accessor of the mount to which the alias should belong to.
	MountAccessor pulumi.StringOutput `pulumi:"mountAccessor"`
	// Name of the alias. Name should be the identifier of the client in the authentication source. For example, if the alias belongs to userpass backend, the name should be a valid username within userpass backend. If alias belongs to GitHub, it should be the GitHub username.
	Name pulumi.StringOutput `pulumi:"name"`
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
	// *Available only for Vault Enterprise*.
	Namespace pulumi.StringPtrOutput `pulumi:"namespace"`
}

// NewEntityAlias registers a new resource with the given unique name, arguments, and options.
func NewEntityAlias(ctx *pulumi.Context,
	name string, args *EntityAliasArgs, opts ...pulumi.ResourceOption) (*EntityAlias, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.CanonicalId == nil {
		return nil, errors.New("invalid value for required argument 'CanonicalId'")
	}
	if args.MountAccessor == nil {
		return nil, errors.New("invalid value for required argument 'MountAccessor'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource EntityAlias
	err := ctx.RegisterResource("vault:identity/entityAlias:EntityAlias", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetEntityAlias gets an existing EntityAlias resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEntityAlias(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EntityAliasState, opts ...pulumi.ResourceOption) (*EntityAlias, error) {
	var resource EntityAlias
	err := ctx.ReadResource("vault:identity/entityAlias:EntityAlias", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering EntityAlias resources.
type entityAliasState struct {
	// Entity ID to which this alias belongs to.
	CanonicalId *string `pulumi:"canonicalId"`
	// Custom metadata to be associated with this alias.
	CustomMetadata map[string]string `pulumi:"customMetadata"`
	// Accessor of the mount to which the alias should belong to.
	MountAccessor *string `pulumi:"mountAccessor"`
	// Name of the alias. Name should be the identifier of the client in the authentication source. For example, if the alias belongs to userpass backend, the name should be a valid username within userpass backend. If alias belongs to GitHub, it should be the GitHub username.
	Name *string `pulumi:"name"`
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
	// *Available only for Vault Enterprise*.
	Namespace *string `pulumi:"namespace"`
}

type EntityAliasState struct {
	// Entity ID to which this alias belongs to.
	CanonicalId pulumi.StringPtrInput
	// Custom metadata to be associated with this alias.
	CustomMetadata pulumi.StringMapInput
	// Accessor of the mount to which the alias should belong to.
	MountAccessor pulumi.StringPtrInput
	// Name of the alias. Name should be the identifier of the client in the authentication source. For example, if the alias belongs to userpass backend, the name should be a valid username within userpass backend. If alias belongs to GitHub, it should be the GitHub username.
	Name pulumi.StringPtrInput
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
	// *Available only for Vault Enterprise*.
	Namespace pulumi.StringPtrInput
}

func (EntityAliasState) ElementType() reflect.Type {
	return reflect.TypeOf((*entityAliasState)(nil)).Elem()
}

type entityAliasArgs struct {
	// Entity ID to which this alias belongs to.
	CanonicalId string `pulumi:"canonicalId"`
	// Custom metadata to be associated with this alias.
	CustomMetadata map[string]string `pulumi:"customMetadata"`
	// Accessor of the mount to which the alias should belong to.
	MountAccessor string `pulumi:"mountAccessor"`
	// Name of the alias. Name should be the identifier of the client in the authentication source. For example, if the alias belongs to userpass backend, the name should be a valid username within userpass backend. If alias belongs to GitHub, it should be the GitHub username.
	Name *string `pulumi:"name"`
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
	// *Available only for Vault Enterprise*.
	Namespace *string `pulumi:"namespace"`
}

// The set of arguments for constructing a EntityAlias resource.
type EntityAliasArgs struct {
	// Entity ID to which this alias belongs to.
	CanonicalId pulumi.StringInput
	// Custom metadata to be associated with this alias.
	CustomMetadata pulumi.StringMapInput
	// Accessor of the mount to which the alias should belong to.
	MountAccessor pulumi.StringInput
	// Name of the alias. Name should be the identifier of the client in the authentication source. For example, if the alias belongs to userpass backend, the name should be a valid username within userpass backend. If alias belongs to GitHub, it should be the GitHub username.
	Name pulumi.StringPtrInput
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
	// *Available only for Vault Enterprise*.
	Namespace pulumi.StringPtrInput
}

func (EntityAliasArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*entityAliasArgs)(nil)).Elem()
}

type EntityAliasInput interface {
	pulumi.Input

	ToEntityAliasOutput() EntityAliasOutput
	ToEntityAliasOutputWithContext(ctx context.Context) EntityAliasOutput
}

func (*EntityAlias) ElementType() reflect.Type {
	return reflect.TypeOf((**EntityAlias)(nil)).Elem()
}

func (i *EntityAlias) ToEntityAliasOutput() EntityAliasOutput {
	return i.ToEntityAliasOutputWithContext(context.Background())
}

func (i *EntityAlias) ToEntityAliasOutputWithContext(ctx context.Context) EntityAliasOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EntityAliasOutput)
}

// EntityAliasArrayInput is an input type that accepts EntityAliasArray and EntityAliasArrayOutput values.
// You can construct a concrete instance of `EntityAliasArrayInput` via:
//
//	EntityAliasArray{ EntityAliasArgs{...} }
type EntityAliasArrayInput interface {
	pulumi.Input

	ToEntityAliasArrayOutput() EntityAliasArrayOutput
	ToEntityAliasArrayOutputWithContext(context.Context) EntityAliasArrayOutput
}

type EntityAliasArray []EntityAliasInput

func (EntityAliasArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*EntityAlias)(nil)).Elem()
}

func (i EntityAliasArray) ToEntityAliasArrayOutput() EntityAliasArrayOutput {
	return i.ToEntityAliasArrayOutputWithContext(context.Background())
}

func (i EntityAliasArray) ToEntityAliasArrayOutputWithContext(ctx context.Context) EntityAliasArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EntityAliasArrayOutput)
}

// EntityAliasMapInput is an input type that accepts EntityAliasMap and EntityAliasMapOutput values.
// You can construct a concrete instance of `EntityAliasMapInput` via:
//
//	EntityAliasMap{ "key": EntityAliasArgs{...} }
type EntityAliasMapInput interface {
	pulumi.Input

	ToEntityAliasMapOutput() EntityAliasMapOutput
	ToEntityAliasMapOutputWithContext(context.Context) EntityAliasMapOutput
}

type EntityAliasMap map[string]EntityAliasInput

func (EntityAliasMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*EntityAlias)(nil)).Elem()
}

func (i EntityAliasMap) ToEntityAliasMapOutput() EntityAliasMapOutput {
	return i.ToEntityAliasMapOutputWithContext(context.Background())
}

func (i EntityAliasMap) ToEntityAliasMapOutputWithContext(ctx context.Context) EntityAliasMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EntityAliasMapOutput)
}

type EntityAliasOutput struct{ *pulumi.OutputState }

func (EntityAliasOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EntityAlias)(nil)).Elem()
}

func (o EntityAliasOutput) ToEntityAliasOutput() EntityAliasOutput {
	return o
}

func (o EntityAliasOutput) ToEntityAliasOutputWithContext(ctx context.Context) EntityAliasOutput {
	return o
}

// Entity ID to which this alias belongs to.
func (o EntityAliasOutput) CanonicalId() pulumi.StringOutput {
	return o.ApplyT(func(v *EntityAlias) pulumi.StringOutput { return v.CanonicalId }).(pulumi.StringOutput)
}

// Custom metadata to be associated with this alias.
func (o EntityAliasOutput) CustomMetadata() pulumi.StringMapOutput {
	return o.ApplyT(func(v *EntityAlias) pulumi.StringMapOutput { return v.CustomMetadata }).(pulumi.StringMapOutput)
}

// Accessor of the mount to which the alias should belong to.
func (o EntityAliasOutput) MountAccessor() pulumi.StringOutput {
	return o.ApplyT(func(v *EntityAlias) pulumi.StringOutput { return v.MountAccessor }).(pulumi.StringOutput)
}

// Name of the alias. Name should be the identifier of the client in the authentication source. For example, if the alias belongs to userpass backend, the name should be a valid username within userpass backend. If alias belongs to GitHub, it should be the GitHub username.
func (o EntityAliasOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *EntityAlias) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The namespace to provision the resource in.
// The value should not contain leading or trailing forward slashes.
// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
// *Available only for Vault Enterprise*.
func (o EntityAliasOutput) Namespace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EntityAlias) pulumi.StringPtrOutput { return v.Namespace }).(pulumi.StringPtrOutput)
}

type EntityAliasArrayOutput struct{ *pulumi.OutputState }

func (EntityAliasArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*EntityAlias)(nil)).Elem()
}

func (o EntityAliasArrayOutput) ToEntityAliasArrayOutput() EntityAliasArrayOutput {
	return o
}

func (o EntityAliasArrayOutput) ToEntityAliasArrayOutputWithContext(ctx context.Context) EntityAliasArrayOutput {
	return o
}

func (o EntityAliasArrayOutput) Index(i pulumi.IntInput) EntityAliasOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *EntityAlias {
		return vs[0].([]*EntityAlias)[vs[1].(int)]
	}).(EntityAliasOutput)
}

type EntityAliasMapOutput struct{ *pulumi.OutputState }

func (EntityAliasMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*EntityAlias)(nil)).Elem()
}

func (o EntityAliasMapOutput) ToEntityAliasMapOutput() EntityAliasMapOutput {
	return o
}

func (o EntityAliasMapOutput) ToEntityAliasMapOutputWithContext(ctx context.Context) EntityAliasMapOutput {
	return o
}

func (o EntityAliasMapOutput) MapIndex(k pulumi.StringInput) EntityAliasOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *EntityAlias {
		return vs[0].(map[string]*EntityAlias)[vs[1].(string)]
	}).(EntityAliasOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*EntityAliasInput)(nil)).Elem(), &EntityAlias{})
	pulumi.RegisterInputType(reflect.TypeOf((*EntityAliasArrayInput)(nil)).Elem(), EntityAliasArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*EntityAliasMapInput)(nil)).Elem(), EntityAliasMap{})
	pulumi.RegisterOutputType(EntityAliasOutput{})
	pulumi.RegisterOutputType(EntityAliasArrayOutput{})
	pulumi.RegisterOutputType(EntityAliasMapOutput{})
}
