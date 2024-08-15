// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vault

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-vault/sdk/v6/go/vault/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## Import
//
// # Namespaces can be imported using its `name` as accessor id
//
// ```sh
// $ pulumi import vault:index/namespace:Namespace example <name>
// ```
//
// If the declared resource is imported and intends to support namespaces using a provider alias, then the name is relative to the namespace path.
//
// hcl
//
// provider "vault" {
//
// # Configuration options
//
//	namespace = "example"
//
//	alias     = "example"
//
// }
//
// resource "vault_namespace" "example2" {
//
//	provider = vault.example
//
//	path     = "example2"
//
// }
//
// ```sh
// $ pulumi import vault:index/namespace:Namespace example2 example2
// ```
//
// $ terraform state show vault_namespace.example2
//
// vault_namespace.example2:
//
// resource "vault_namespace" "example2" {
//
//	id           = "example/example2/"
//
//	namespace_id = <known after import>
//
//	path         = "example2"
//
//	path_fq      = "example2"
//
// }
type Namespace struct {
	pulumi.CustomResourceState

	// Custom metadata describing this namespace. Value type
	// is `map[string]string`. Requires Vault version 1.12+.
	CustomMetadata pulumi.StringMapOutput `pulumi:"customMetadata"`
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
	// *Available only for Vault Enterprise*.
	Namespace pulumi.StringPtrOutput `pulumi:"namespace"`
	// Vault server's internal ID of the namespace.
	NamespaceId pulumi.StringOutput `pulumi:"namespaceId"`
	// The path of the namespace. Must not have a trailing `/`.
	Path pulumi.StringOutput `pulumi:"path"`
	// The fully qualified path to the namespace. Useful when provisioning resources in a child `namespace`.
	// The path is relative to the provider's `namespace` argument.
	PathFq pulumi.StringOutput `pulumi:"pathFq"`
}

// NewNamespace registers a new resource with the given unique name, arguments, and options.
func NewNamespace(ctx *pulumi.Context,
	name string, args *NamespaceArgs, opts ...pulumi.ResourceOption) (*Namespace, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Path == nil {
		return nil, errors.New("invalid value for required argument 'Path'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Namespace
	err := ctx.RegisterResource("vault:index/namespace:Namespace", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNamespace gets an existing Namespace resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNamespace(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NamespaceState, opts ...pulumi.ResourceOption) (*Namespace, error) {
	var resource Namespace
	err := ctx.ReadResource("vault:index/namespace:Namespace", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Namespace resources.
type namespaceState struct {
	// Custom metadata describing this namespace. Value type
	// is `map[string]string`. Requires Vault version 1.12+.
	CustomMetadata map[string]string `pulumi:"customMetadata"`
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
	// *Available only for Vault Enterprise*.
	Namespace *string `pulumi:"namespace"`
	// Vault server's internal ID of the namespace.
	NamespaceId *string `pulumi:"namespaceId"`
	// The path of the namespace. Must not have a trailing `/`.
	Path *string `pulumi:"path"`
	// The fully qualified path to the namespace. Useful when provisioning resources in a child `namespace`.
	// The path is relative to the provider's `namespace` argument.
	PathFq *string `pulumi:"pathFq"`
}

type NamespaceState struct {
	// Custom metadata describing this namespace. Value type
	// is `map[string]string`. Requires Vault version 1.12+.
	CustomMetadata pulumi.StringMapInput
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
	// *Available only for Vault Enterprise*.
	Namespace pulumi.StringPtrInput
	// Vault server's internal ID of the namespace.
	NamespaceId pulumi.StringPtrInput
	// The path of the namespace. Must not have a trailing `/`.
	Path pulumi.StringPtrInput
	// The fully qualified path to the namespace. Useful when provisioning resources in a child `namespace`.
	// The path is relative to the provider's `namespace` argument.
	PathFq pulumi.StringPtrInput
}

func (NamespaceState) ElementType() reflect.Type {
	return reflect.TypeOf((*namespaceState)(nil)).Elem()
}

type namespaceArgs struct {
	// Custom metadata describing this namespace. Value type
	// is `map[string]string`. Requires Vault version 1.12+.
	CustomMetadata map[string]string `pulumi:"customMetadata"`
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
	// *Available only for Vault Enterprise*.
	Namespace *string `pulumi:"namespace"`
	// The path of the namespace. Must not have a trailing `/`.
	Path string `pulumi:"path"`
	// The fully qualified path to the namespace. Useful when provisioning resources in a child `namespace`.
	// The path is relative to the provider's `namespace` argument.
	PathFq *string `pulumi:"pathFq"`
}

// The set of arguments for constructing a Namespace resource.
type NamespaceArgs struct {
	// Custom metadata describing this namespace. Value type
	// is `map[string]string`. Requires Vault version 1.12+.
	CustomMetadata pulumi.StringMapInput
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
	// *Available only for Vault Enterprise*.
	Namespace pulumi.StringPtrInput
	// The path of the namespace. Must not have a trailing `/`.
	Path pulumi.StringInput
	// The fully qualified path to the namespace. Useful when provisioning resources in a child `namespace`.
	// The path is relative to the provider's `namespace` argument.
	PathFq pulumi.StringPtrInput
}

func (NamespaceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*namespaceArgs)(nil)).Elem()
}

type NamespaceInput interface {
	pulumi.Input

	ToNamespaceOutput() NamespaceOutput
	ToNamespaceOutputWithContext(ctx context.Context) NamespaceOutput
}

func (*Namespace) ElementType() reflect.Type {
	return reflect.TypeOf((**Namespace)(nil)).Elem()
}

func (i *Namespace) ToNamespaceOutput() NamespaceOutput {
	return i.ToNamespaceOutputWithContext(context.Background())
}

func (i *Namespace) ToNamespaceOutputWithContext(ctx context.Context) NamespaceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NamespaceOutput)
}

// NamespaceArrayInput is an input type that accepts NamespaceArray and NamespaceArrayOutput values.
// You can construct a concrete instance of `NamespaceArrayInput` via:
//
//	NamespaceArray{ NamespaceArgs{...} }
type NamespaceArrayInput interface {
	pulumi.Input

	ToNamespaceArrayOutput() NamespaceArrayOutput
	ToNamespaceArrayOutputWithContext(context.Context) NamespaceArrayOutput
}

type NamespaceArray []NamespaceInput

func (NamespaceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Namespace)(nil)).Elem()
}

func (i NamespaceArray) ToNamespaceArrayOutput() NamespaceArrayOutput {
	return i.ToNamespaceArrayOutputWithContext(context.Background())
}

func (i NamespaceArray) ToNamespaceArrayOutputWithContext(ctx context.Context) NamespaceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NamespaceArrayOutput)
}

// NamespaceMapInput is an input type that accepts NamespaceMap and NamespaceMapOutput values.
// You can construct a concrete instance of `NamespaceMapInput` via:
//
//	NamespaceMap{ "key": NamespaceArgs{...} }
type NamespaceMapInput interface {
	pulumi.Input

	ToNamespaceMapOutput() NamespaceMapOutput
	ToNamespaceMapOutputWithContext(context.Context) NamespaceMapOutput
}

type NamespaceMap map[string]NamespaceInput

func (NamespaceMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Namespace)(nil)).Elem()
}

func (i NamespaceMap) ToNamespaceMapOutput() NamespaceMapOutput {
	return i.ToNamespaceMapOutputWithContext(context.Background())
}

func (i NamespaceMap) ToNamespaceMapOutputWithContext(ctx context.Context) NamespaceMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NamespaceMapOutput)
}

type NamespaceOutput struct{ *pulumi.OutputState }

func (NamespaceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Namespace)(nil)).Elem()
}

func (o NamespaceOutput) ToNamespaceOutput() NamespaceOutput {
	return o
}

func (o NamespaceOutput) ToNamespaceOutputWithContext(ctx context.Context) NamespaceOutput {
	return o
}

// Custom metadata describing this namespace. Value type
// is `map[string]string`. Requires Vault version 1.12+.
func (o NamespaceOutput) CustomMetadata() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Namespace) pulumi.StringMapOutput { return v.CustomMetadata }).(pulumi.StringMapOutput)
}

// The namespace to provision the resource in.
// The value should not contain leading or trailing forward slashes.
// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
// *Available only for Vault Enterprise*.
func (o NamespaceOutput) Namespace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Namespace) pulumi.StringPtrOutput { return v.Namespace }).(pulumi.StringPtrOutput)
}

// Vault server's internal ID of the namespace.
func (o NamespaceOutput) NamespaceId() pulumi.StringOutput {
	return o.ApplyT(func(v *Namespace) pulumi.StringOutput { return v.NamespaceId }).(pulumi.StringOutput)
}

// The path of the namespace. Must not have a trailing `/`.
func (o NamespaceOutput) Path() pulumi.StringOutput {
	return o.ApplyT(func(v *Namespace) pulumi.StringOutput { return v.Path }).(pulumi.StringOutput)
}

// The fully qualified path to the namespace. Useful when provisioning resources in a child `namespace`.
// The path is relative to the provider's `namespace` argument.
func (o NamespaceOutput) PathFq() pulumi.StringOutput {
	return o.ApplyT(func(v *Namespace) pulumi.StringOutput { return v.PathFq }).(pulumi.StringOutput)
}

type NamespaceArrayOutput struct{ *pulumi.OutputState }

func (NamespaceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Namespace)(nil)).Elem()
}

func (o NamespaceArrayOutput) ToNamespaceArrayOutput() NamespaceArrayOutput {
	return o
}

func (o NamespaceArrayOutput) ToNamespaceArrayOutputWithContext(ctx context.Context) NamespaceArrayOutput {
	return o
}

func (o NamespaceArrayOutput) Index(i pulumi.IntInput) NamespaceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Namespace {
		return vs[0].([]*Namespace)[vs[1].(int)]
	}).(NamespaceOutput)
}

type NamespaceMapOutput struct{ *pulumi.OutputState }

func (NamespaceMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Namespace)(nil)).Elem()
}

func (o NamespaceMapOutput) ToNamespaceMapOutput() NamespaceMapOutput {
	return o
}

func (o NamespaceMapOutput) ToNamespaceMapOutputWithContext(ctx context.Context) NamespaceMapOutput {
	return o
}

func (o NamespaceMapOutput) MapIndex(k pulumi.StringInput) NamespaceOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Namespace {
		return vs[0].(map[string]*Namespace)[vs[1].(string)]
	}).(NamespaceOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*NamespaceInput)(nil)).Elem(), &Namespace{})
	pulumi.RegisterInputType(reflect.TypeOf((*NamespaceArrayInput)(nil)).Elem(), NamespaceArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*NamespaceMapInput)(nil)).Elem(), NamespaceMap{})
	pulumi.RegisterOutputType(NamespaceOutput{})
	pulumi.RegisterOutputType(NamespaceArrayOutput{})
	pulumi.RegisterOutputType(NamespaceMapOutput{})
}
