// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package transform

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-vault/sdk/v5/go/vault/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This resource supports the "/transform/alphabet/{name}" Vault endpoint.
//
// It queries an existing alphabet by the given name.
//
// ## Example Usage
//
// <!--Start PulumiCodeChooser -->
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-vault/sdk/v5/go/vault"
//	"github.com/pulumi/pulumi-vault/sdk/v5/go/vault/transform"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			mountTransform, err := vault.NewMount(ctx, "mountTransform", &vault.MountArgs{
//				Path: pulumi.String("transform"),
//				Type: pulumi.String("transform"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = transform.NewAlphabet(ctx, "test", &transform.AlphabetArgs{
//				Path:     mountTransform.Path,
//				Alphabet: pulumi.String("0123456789"),
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
type Alphabet struct {
	pulumi.CustomResourceState

	// A string of characters that contains the alphabet set.
	Alphabet pulumi.StringPtrOutput `pulumi:"alphabet"`
	// The name of the alphabet.
	Name pulumi.StringOutput `pulumi:"name"`
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
	// *Available only for Vault Enterprise*.
	Namespace pulumi.StringPtrOutput `pulumi:"namespace"`
	// Path to where the back-end is mounted within Vault.
	Path pulumi.StringOutput `pulumi:"path"`
}

// NewAlphabet registers a new resource with the given unique name, arguments, and options.
func NewAlphabet(ctx *pulumi.Context,
	name string, args *AlphabetArgs, opts ...pulumi.ResourceOption) (*Alphabet, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Path == nil {
		return nil, errors.New("invalid value for required argument 'Path'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Alphabet
	err := ctx.RegisterResource("vault:transform/alphabet:Alphabet", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAlphabet gets an existing Alphabet resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAlphabet(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AlphabetState, opts ...pulumi.ResourceOption) (*Alphabet, error) {
	var resource Alphabet
	err := ctx.ReadResource("vault:transform/alphabet:Alphabet", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Alphabet resources.
type alphabetState struct {
	// A string of characters that contains the alphabet set.
	Alphabet *string `pulumi:"alphabet"`
	// The name of the alphabet.
	Name *string `pulumi:"name"`
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
	// *Available only for Vault Enterprise*.
	Namespace *string `pulumi:"namespace"`
	// Path to where the back-end is mounted within Vault.
	Path *string `pulumi:"path"`
}

type AlphabetState struct {
	// A string of characters that contains the alphabet set.
	Alphabet pulumi.StringPtrInput
	// The name of the alphabet.
	Name pulumi.StringPtrInput
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
	// *Available only for Vault Enterprise*.
	Namespace pulumi.StringPtrInput
	// Path to where the back-end is mounted within Vault.
	Path pulumi.StringPtrInput
}

func (AlphabetState) ElementType() reflect.Type {
	return reflect.TypeOf((*alphabetState)(nil)).Elem()
}

type alphabetArgs struct {
	// A string of characters that contains the alphabet set.
	Alphabet *string `pulumi:"alphabet"`
	// The name of the alphabet.
	Name *string `pulumi:"name"`
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
	// *Available only for Vault Enterprise*.
	Namespace *string `pulumi:"namespace"`
	// Path to where the back-end is mounted within Vault.
	Path string `pulumi:"path"`
}

// The set of arguments for constructing a Alphabet resource.
type AlphabetArgs struct {
	// A string of characters that contains the alphabet set.
	Alphabet pulumi.StringPtrInput
	// The name of the alphabet.
	Name pulumi.StringPtrInput
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
	// *Available only for Vault Enterprise*.
	Namespace pulumi.StringPtrInput
	// Path to where the back-end is mounted within Vault.
	Path pulumi.StringInput
}

func (AlphabetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*alphabetArgs)(nil)).Elem()
}

type AlphabetInput interface {
	pulumi.Input

	ToAlphabetOutput() AlphabetOutput
	ToAlphabetOutputWithContext(ctx context.Context) AlphabetOutput
}

func (*Alphabet) ElementType() reflect.Type {
	return reflect.TypeOf((**Alphabet)(nil)).Elem()
}

func (i *Alphabet) ToAlphabetOutput() AlphabetOutput {
	return i.ToAlphabetOutputWithContext(context.Background())
}

func (i *Alphabet) ToAlphabetOutputWithContext(ctx context.Context) AlphabetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AlphabetOutput)
}

// AlphabetArrayInput is an input type that accepts AlphabetArray and AlphabetArrayOutput values.
// You can construct a concrete instance of `AlphabetArrayInput` via:
//
//	AlphabetArray{ AlphabetArgs{...} }
type AlphabetArrayInput interface {
	pulumi.Input

	ToAlphabetArrayOutput() AlphabetArrayOutput
	ToAlphabetArrayOutputWithContext(context.Context) AlphabetArrayOutput
}

type AlphabetArray []AlphabetInput

func (AlphabetArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Alphabet)(nil)).Elem()
}

func (i AlphabetArray) ToAlphabetArrayOutput() AlphabetArrayOutput {
	return i.ToAlphabetArrayOutputWithContext(context.Background())
}

func (i AlphabetArray) ToAlphabetArrayOutputWithContext(ctx context.Context) AlphabetArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AlphabetArrayOutput)
}

// AlphabetMapInput is an input type that accepts AlphabetMap and AlphabetMapOutput values.
// You can construct a concrete instance of `AlphabetMapInput` via:
//
//	AlphabetMap{ "key": AlphabetArgs{...} }
type AlphabetMapInput interface {
	pulumi.Input

	ToAlphabetMapOutput() AlphabetMapOutput
	ToAlphabetMapOutputWithContext(context.Context) AlphabetMapOutput
}

type AlphabetMap map[string]AlphabetInput

func (AlphabetMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Alphabet)(nil)).Elem()
}

func (i AlphabetMap) ToAlphabetMapOutput() AlphabetMapOutput {
	return i.ToAlphabetMapOutputWithContext(context.Background())
}

func (i AlphabetMap) ToAlphabetMapOutputWithContext(ctx context.Context) AlphabetMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AlphabetMapOutput)
}

type AlphabetOutput struct{ *pulumi.OutputState }

func (AlphabetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Alphabet)(nil)).Elem()
}

func (o AlphabetOutput) ToAlphabetOutput() AlphabetOutput {
	return o
}

func (o AlphabetOutput) ToAlphabetOutputWithContext(ctx context.Context) AlphabetOutput {
	return o
}

// A string of characters that contains the alphabet set.
func (o AlphabetOutput) Alphabet() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Alphabet) pulumi.StringPtrOutput { return v.Alphabet }).(pulumi.StringPtrOutput)
}

// The name of the alphabet.
func (o AlphabetOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Alphabet) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The namespace to provision the resource in.
// The value should not contain leading or trailing forward slashes.
// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
// *Available only for Vault Enterprise*.
func (o AlphabetOutput) Namespace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Alphabet) pulumi.StringPtrOutput { return v.Namespace }).(pulumi.StringPtrOutput)
}

// Path to where the back-end is mounted within Vault.
func (o AlphabetOutput) Path() pulumi.StringOutput {
	return o.ApplyT(func(v *Alphabet) pulumi.StringOutput { return v.Path }).(pulumi.StringOutput)
}

type AlphabetArrayOutput struct{ *pulumi.OutputState }

func (AlphabetArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Alphabet)(nil)).Elem()
}

func (o AlphabetArrayOutput) ToAlphabetArrayOutput() AlphabetArrayOutput {
	return o
}

func (o AlphabetArrayOutput) ToAlphabetArrayOutputWithContext(ctx context.Context) AlphabetArrayOutput {
	return o
}

func (o AlphabetArrayOutput) Index(i pulumi.IntInput) AlphabetOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Alphabet {
		return vs[0].([]*Alphabet)[vs[1].(int)]
	}).(AlphabetOutput)
}

type AlphabetMapOutput struct{ *pulumi.OutputState }

func (AlphabetMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Alphabet)(nil)).Elem()
}

func (o AlphabetMapOutput) ToAlphabetMapOutput() AlphabetMapOutput {
	return o
}

func (o AlphabetMapOutput) ToAlphabetMapOutputWithContext(ctx context.Context) AlphabetMapOutput {
	return o
}

func (o AlphabetMapOutput) MapIndex(k pulumi.StringInput) AlphabetOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Alphabet {
		return vs[0].(map[string]*Alphabet)[vs[1].(string)]
	}).(AlphabetOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AlphabetInput)(nil)).Elem(), &Alphabet{})
	pulumi.RegisterInputType(reflect.TypeOf((*AlphabetArrayInput)(nil)).Elem(), AlphabetArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AlphabetMapInput)(nil)).Elem(), AlphabetMap{})
	pulumi.RegisterOutputType(AlphabetOutput{})
	pulumi.RegisterOutputType(AlphabetArrayOutput{})
	pulumi.RegisterOutputType(AlphabetMapOutput{})
}
