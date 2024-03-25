// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package transform

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-vault/sdk/v6/go/vault/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Transformation struct {
	pulumi.CustomResourceState

	// The set of roles allowed to perform this transformation.
	AllowedRoles pulumi.StringArrayOutput `pulumi:"allowedRoles"`
	// If true, this transform can be deleted.
	// Otherwise, deletion is blocked while this value remains false. Default: `false`
	// *Only supported on vault-1.12+*
	DeletionAllowed pulumi.BoolPtrOutput `pulumi:"deletionAllowed"`
	// The character used to replace data when in masking mode
	MaskingCharacter pulumi.StringPtrOutput `pulumi:"maskingCharacter"`
	// The name of the transformation.
	Name pulumi.StringOutput `pulumi:"name"`
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
	// *Available only for Vault Enterprise*.
	Namespace pulumi.StringPtrOutput `pulumi:"namespace"`
	// Path to where the back-end is mounted within Vault.
	Path pulumi.StringOutput `pulumi:"path"`
	// The name of the template to use.
	Template pulumi.StringPtrOutput `pulumi:"template"`
	// Templates configured for transformation.
	Templates pulumi.StringArrayOutput `pulumi:"templates"`
	// The source of where the tweak value comes from. Only valid when in FPE mode.
	TweakSource pulumi.StringPtrOutput `pulumi:"tweakSource"`
	// The type of transformation to perform.
	Type pulumi.StringPtrOutput `pulumi:"type"`
}

// NewTransformation registers a new resource with the given unique name, arguments, and options.
func NewTransformation(ctx *pulumi.Context,
	name string, args *TransformationArgs, opts ...pulumi.ResourceOption) (*Transformation, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Path == nil {
		return nil, errors.New("invalid value for required argument 'Path'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Transformation
	err := ctx.RegisterResource("vault:transform/transformation:Transformation", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTransformation gets an existing Transformation resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTransformation(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TransformationState, opts ...pulumi.ResourceOption) (*Transformation, error) {
	var resource Transformation
	err := ctx.ReadResource("vault:transform/transformation:Transformation", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Transformation resources.
type transformationState struct {
	// The set of roles allowed to perform this transformation.
	AllowedRoles []string `pulumi:"allowedRoles"`
	// If true, this transform can be deleted.
	// Otherwise, deletion is blocked while this value remains false. Default: `false`
	// *Only supported on vault-1.12+*
	DeletionAllowed *bool `pulumi:"deletionAllowed"`
	// The character used to replace data when in masking mode
	MaskingCharacter *string `pulumi:"maskingCharacter"`
	// The name of the transformation.
	Name *string `pulumi:"name"`
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
	// *Available only for Vault Enterprise*.
	Namespace *string `pulumi:"namespace"`
	// Path to where the back-end is mounted within Vault.
	Path *string `pulumi:"path"`
	// The name of the template to use.
	Template *string `pulumi:"template"`
	// Templates configured for transformation.
	Templates []string `pulumi:"templates"`
	// The source of where the tweak value comes from. Only valid when in FPE mode.
	TweakSource *string `pulumi:"tweakSource"`
	// The type of transformation to perform.
	Type *string `pulumi:"type"`
}

type TransformationState struct {
	// The set of roles allowed to perform this transformation.
	AllowedRoles pulumi.StringArrayInput
	// If true, this transform can be deleted.
	// Otherwise, deletion is blocked while this value remains false. Default: `false`
	// *Only supported on vault-1.12+*
	DeletionAllowed pulumi.BoolPtrInput
	// The character used to replace data when in masking mode
	MaskingCharacter pulumi.StringPtrInput
	// The name of the transformation.
	Name pulumi.StringPtrInput
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
	// *Available only for Vault Enterprise*.
	Namespace pulumi.StringPtrInput
	// Path to where the back-end is mounted within Vault.
	Path pulumi.StringPtrInput
	// The name of the template to use.
	Template pulumi.StringPtrInput
	// Templates configured for transformation.
	Templates pulumi.StringArrayInput
	// The source of where the tweak value comes from. Only valid when in FPE mode.
	TweakSource pulumi.StringPtrInput
	// The type of transformation to perform.
	Type pulumi.StringPtrInput
}

func (TransformationState) ElementType() reflect.Type {
	return reflect.TypeOf((*transformationState)(nil)).Elem()
}

type transformationArgs struct {
	// The set of roles allowed to perform this transformation.
	AllowedRoles []string `pulumi:"allowedRoles"`
	// If true, this transform can be deleted.
	// Otherwise, deletion is blocked while this value remains false. Default: `false`
	// *Only supported on vault-1.12+*
	DeletionAllowed *bool `pulumi:"deletionAllowed"`
	// The character used to replace data when in masking mode
	MaskingCharacter *string `pulumi:"maskingCharacter"`
	// The name of the transformation.
	Name *string `pulumi:"name"`
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
	// *Available only for Vault Enterprise*.
	Namespace *string `pulumi:"namespace"`
	// Path to where the back-end is mounted within Vault.
	Path string `pulumi:"path"`
	// The name of the template to use.
	Template *string `pulumi:"template"`
	// Templates configured for transformation.
	Templates []string `pulumi:"templates"`
	// The source of where the tweak value comes from. Only valid when in FPE mode.
	TweakSource *string `pulumi:"tweakSource"`
	// The type of transformation to perform.
	Type *string `pulumi:"type"`
}

// The set of arguments for constructing a Transformation resource.
type TransformationArgs struct {
	// The set of roles allowed to perform this transformation.
	AllowedRoles pulumi.StringArrayInput
	// If true, this transform can be deleted.
	// Otherwise, deletion is blocked while this value remains false. Default: `false`
	// *Only supported on vault-1.12+*
	DeletionAllowed pulumi.BoolPtrInput
	// The character used to replace data when in masking mode
	MaskingCharacter pulumi.StringPtrInput
	// The name of the transformation.
	Name pulumi.StringPtrInput
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
	// *Available only for Vault Enterprise*.
	Namespace pulumi.StringPtrInput
	// Path to where the back-end is mounted within Vault.
	Path pulumi.StringInput
	// The name of the template to use.
	Template pulumi.StringPtrInput
	// Templates configured for transformation.
	Templates pulumi.StringArrayInput
	// The source of where the tweak value comes from. Only valid when in FPE mode.
	TweakSource pulumi.StringPtrInput
	// The type of transformation to perform.
	Type pulumi.StringPtrInput
}

func (TransformationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*transformationArgs)(nil)).Elem()
}

type TransformationInput interface {
	pulumi.Input

	ToTransformationOutput() TransformationOutput
	ToTransformationOutputWithContext(ctx context.Context) TransformationOutput
}

func (*Transformation) ElementType() reflect.Type {
	return reflect.TypeOf((**Transformation)(nil)).Elem()
}

func (i *Transformation) ToTransformationOutput() TransformationOutput {
	return i.ToTransformationOutputWithContext(context.Background())
}

func (i *Transformation) ToTransformationOutputWithContext(ctx context.Context) TransformationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TransformationOutput)
}

// TransformationArrayInput is an input type that accepts TransformationArray and TransformationArrayOutput values.
// You can construct a concrete instance of `TransformationArrayInput` via:
//
//	TransformationArray{ TransformationArgs{...} }
type TransformationArrayInput interface {
	pulumi.Input

	ToTransformationArrayOutput() TransformationArrayOutput
	ToTransformationArrayOutputWithContext(context.Context) TransformationArrayOutput
}

type TransformationArray []TransformationInput

func (TransformationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Transformation)(nil)).Elem()
}

func (i TransformationArray) ToTransformationArrayOutput() TransformationArrayOutput {
	return i.ToTransformationArrayOutputWithContext(context.Background())
}

func (i TransformationArray) ToTransformationArrayOutputWithContext(ctx context.Context) TransformationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TransformationArrayOutput)
}

// TransformationMapInput is an input type that accepts TransformationMap and TransformationMapOutput values.
// You can construct a concrete instance of `TransformationMapInput` via:
//
//	TransformationMap{ "key": TransformationArgs{...} }
type TransformationMapInput interface {
	pulumi.Input

	ToTransformationMapOutput() TransformationMapOutput
	ToTransformationMapOutputWithContext(context.Context) TransformationMapOutput
}

type TransformationMap map[string]TransformationInput

func (TransformationMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Transformation)(nil)).Elem()
}

func (i TransformationMap) ToTransformationMapOutput() TransformationMapOutput {
	return i.ToTransformationMapOutputWithContext(context.Background())
}

func (i TransformationMap) ToTransformationMapOutputWithContext(ctx context.Context) TransformationMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TransformationMapOutput)
}

type TransformationOutput struct{ *pulumi.OutputState }

func (TransformationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Transformation)(nil)).Elem()
}

func (o TransformationOutput) ToTransformationOutput() TransformationOutput {
	return o
}

func (o TransformationOutput) ToTransformationOutputWithContext(ctx context.Context) TransformationOutput {
	return o
}

// The set of roles allowed to perform this transformation.
func (o TransformationOutput) AllowedRoles() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Transformation) pulumi.StringArrayOutput { return v.AllowedRoles }).(pulumi.StringArrayOutput)
}

// If true, this transform can be deleted.
// Otherwise, deletion is blocked while this value remains false. Default: `false`
// *Only supported on vault-1.12+*
func (o TransformationOutput) DeletionAllowed() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Transformation) pulumi.BoolPtrOutput { return v.DeletionAllowed }).(pulumi.BoolPtrOutput)
}

// The character used to replace data when in masking mode
func (o TransformationOutput) MaskingCharacter() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Transformation) pulumi.StringPtrOutput { return v.MaskingCharacter }).(pulumi.StringPtrOutput)
}

// The name of the transformation.
func (o TransformationOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Transformation) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The namespace to provision the resource in.
// The value should not contain leading or trailing forward slashes.
// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
// *Available only for Vault Enterprise*.
func (o TransformationOutput) Namespace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Transformation) pulumi.StringPtrOutput { return v.Namespace }).(pulumi.StringPtrOutput)
}

// Path to where the back-end is mounted within Vault.
func (o TransformationOutput) Path() pulumi.StringOutput {
	return o.ApplyT(func(v *Transformation) pulumi.StringOutput { return v.Path }).(pulumi.StringOutput)
}

// The name of the template to use.
func (o TransformationOutput) Template() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Transformation) pulumi.StringPtrOutput { return v.Template }).(pulumi.StringPtrOutput)
}

// Templates configured for transformation.
func (o TransformationOutput) Templates() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Transformation) pulumi.StringArrayOutput { return v.Templates }).(pulumi.StringArrayOutput)
}

// The source of where the tweak value comes from. Only valid when in FPE mode.
func (o TransformationOutput) TweakSource() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Transformation) pulumi.StringPtrOutput { return v.TweakSource }).(pulumi.StringPtrOutput)
}

// The type of transformation to perform.
func (o TransformationOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Transformation) pulumi.StringPtrOutput { return v.Type }).(pulumi.StringPtrOutput)
}

type TransformationArrayOutput struct{ *pulumi.OutputState }

func (TransformationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Transformation)(nil)).Elem()
}

func (o TransformationArrayOutput) ToTransformationArrayOutput() TransformationArrayOutput {
	return o
}

func (o TransformationArrayOutput) ToTransformationArrayOutputWithContext(ctx context.Context) TransformationArrayOutput {
	return o
}

func (o TransformationArrayOutput) Index(i pulumi.IntInput) TransformationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Transformation {
		return vs[0].([]*Transformation)[vs[1].(int)]
	}).(TransformationOutput)
}

type TransformationMapOutput struct{ *pulumi.OutputState }

func (TransformationMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Transformation)(nil)).Elem()
}

func (o TransformationMapOutput) ToTransformationMapOutput() TransformationMapOutput {
	return o
}

func (o TransformationMapOutput) ToTransformationMapOutputWithContext(ctx context.Context) TransformationMapOutput {
	return o
}

func (o TransformationMapOutput) MapIndex(k pulumi.StringInput) TransformationOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Transformation {
		return vs[0].(map[string]*Transformation)[vs[1].(string)]
	}).(TransformationOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*TransformationInput)(nil)).Elem(), &Transformation{})
	pulumi.RegisterInputType(reflect.TypeOf((*TransformationArrayInput)(nil)).Elem(), TransformationArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*TransformationMapInput)(nil)).Elem(), TransformationMap{})
	pulumi.RegisterOutputType(TransformationOutput{})
	pulumi.RegisterOutputType(TransformationArrayOutput{})
	pulumi.RegisterOutputType(TransformationMapOutput{})
}
