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

// This resource supports the `/transform/template/{name}` Vault endpoint.
//
// It creates or updates a template with the given name. If a template with the name does not exist,
// it will be created. If the template exists, it will be updated with the new attributes.
//
// > Requires _Vault Enterprise with the Advanced Data Protection Transform Module_.
// See [Transform Secrets Engine](https://www.vaultproject.io/docs/secrets/transform)
// for more information.
//
// ## Example Usage
//
// Please note that the `pattern` below holds a regex. The regex shown
// is identical to the one in our [Setup](https://www.vaultproject.io/docs/secrets/transform#setup)
// docs, `(\d{4})-(\d{4})-(\d{4})-(\d{4})`. However, due to HCL, the
// backslashes must be escaped to appear correctly in Vault. For further
// assistance escaping your own custom regex, see String Literals.
//
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
//			transform, err := vault.NewMount(ctx, "transform", &vault.MountArgs{
//				Path: pulumi.String("transform"),
//				Type: pulumi.String("transform"),
//			})
//			if err != nil {
//				return err
//			}
//			numerics, err := transform.NewAlphabet(ctx, "numerics", &transform.AlphabetArgs{
//				Path:     transform.Path,
//				Alphabet: pulumi.String("0123456789"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = transform.NewTemplate(ctx, "test", &transform.TemplateArgs{
//				Path:         numerics.Path,
//				Type:         pulumi.String("regex"),
//				Pattern:      pulumi.String("(\\d{4})[- ](\\d{4})[- ](\\d{4})[- ](\\d{4})"),
//				Alphabet:     pulumi.String("numerics"),
//				EncodeFormat: pulumi.String("$1-$2-$3-$4"),
//				DecodeFormats: pulumi.Map{
//					"last-four-digits": pulumi.Any("$4"),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
type Template struct {
	pulumi.CustomResourceState

	// The alphabet to use for this template. This is only used during FPE transformations.
	Alphabet pulumi.StringPtrOutput `pulumi:"alphabet"`
	// Optional mapping of name to regular expression template, used to customize
	// the decoded output. (requires Vault Enterprise 1.9+)
	DecodeFormats pulumi.MapOutput `pulumi:"decodeFormats"`
	// The regular expression template used to format encoded values.
	// (requires Vault Enterprise 1.9+)
	EncodeFormat pulumi.StringPtrOutput `pulumi:"encodeFormat"`
	// The name of the template.
	Name pulumi.StringOutput `pulumi:"name"`
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
	// *Available only for Vault Enterprise*.
	Namespace pulumi.StringPtrOutput `pulumi:"namespace"`
	// Path to where the back-end is mounted within Vault.
	Path pulumi.StringOutput `pulumi:"path"`
	// The pattern used for matching. Currently, only regular expression pattern is supported.
	Pattern pulumi.StringPtrOutput `pulumi:"pattern"`
	// The pattern type to use for match detection. Currently, only regex is supported.
	Type pulumi.StringPtrOutput `pulumi:"type"`
}

// NewTemplate registers a new resource with the given unique name, arguments, and options.
func NewTemplate(ctx *pulumi.Context,
	name string, args *TemplateArgs, opts ...pulumi.ResourceOption) (*Template, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Path == nil {
		return nil, errors.New("invalid value for required argument 'Path'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Template
	err := ctx.RegisterResource("vault:transform/template:Template", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTemplate gets an existing Template resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTemplate(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TemplateState, opts ...pulumi.ResourceOption) (*Template, error) {
	var resource Template
	err := ctx.ReadResource("vault:transform/template:Template", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Template resources.
type templateState struct {
	// The alphabet to use for this template. This is only used during FPE transformations.
	Alphabet *string `pulumi:"alphabet"`
	// Optional mapping of name to regular expression template, used to customize
	// the decoded output. (requires Vault Enterprise 1.9+)
	DecodeFormats map[string]interface{} `pulumi:"decodeFormats"`
	// The regular expression template used to format encoded values.
	// (requires Vault Enterprise 1.9+)
	EncodeFormat *string `pulumi:"encodeFormat"`
	// The name of the template.
	Name *string `pulumi:"name"`
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
	// *Available only for Vault Enterprise*.
	Namespace *string `pulumi:"namespace"`
	// Path to where the back-end is mounted within Vault.
	Path *string `pulumi:"path"`
	// The pattern used for matching. Currently, only regular expression pattern is supported.
	Pattern *string `pulumi:"pattern"`
	// The pattern type to use for match detection. Currently, only regex is supported.
	Type *string `pulumi:"type"`
}

type TemplateState struct {
	// The alphabet to use for this template. This is only used during FPE transformations.
	Alphabet pulumi.StringPtrInput
	// Optional mapping of name to regular expression template, used to customize
	// the decoded output. (requires Vault Enterprise 1.9+)
	DecodeFormats pulumi.MapInput
	// The regular expression template used to format encoded values.
	// (requires Vault Enterprise 1.9+)
	EncodeFormat pulumi.StringPtrInput
	// The name of the template.
	Name pulumi.StringPtrInput
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
	// *Available only for Vault Enterprise*.
	Namespace pulumi.StringPtrInput
	// Path to where the back-end is mounted within Vault.
	Path pulumi.StringPtrInput
	// The pattern used for matching. Currently, only regular expression pattern is supported.
	Pattern pulumi.StringPtrInput
	// The pattern type to use for match detection. Currently, only regex is supported.
	Type pulumi.StringPtrInput
}

func (TemplateState) ElementType() reflect.Type {
	return reflect.TypeOf((*templateState)(nil)).Elem()
}

type templateArgs struct {
	// The alphabet to use for this template. This is only used during FPE transformations.
	Alphabet *string `pulumi:"alphabet"`
	// Optional mapping of name to regular expression template, used to customize
	// the decoded output. (requires Vault Enterprise 1.9+)
	DecodeFormats map[string]interface{} `pulumi:"decodeFormats"`
	// The regular expression template used to format encoded values.
	// (requires Vault Enterprise 1.9+)
	EncodeFormat *string `pulumi:"encodeFormat"`
	// The name of the template.
	Name *string `pulumi:"name"`
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
	// *Available only for Vault Enterprise*.
	Namespace *string `pulumi:"namespace"`
	// Path to where the back-end is mounted within Vault.
	Path string `pulumi:"path"`
	// The pattern used for matching. Currently, only regular expression pattern is supported.
	Pattern *string `pulumi:"pattern"`
	// The pattern type to use for match detection. Currently, only regex is supported.
	Type *string `pulumi:"type"`
}

// The set of arguments for constructing a Template resource.
type TemplateArgs struct {
	// The alphabet to use for this template. This is only used during FPE transformations.
	Alphabet pulumi.StringPtrInput
	// Optional mapping of name to regular expression template, used to customize
	// the decoded output. (requires Vault Enterprise 1.9+)
	DecodeFormats pulumi.MapInput
	// The regular expression template used to format encoded values.
	// (requires Vault Enterprise 1.9+)
	EncodeFormat pulumi.StringPtrInput
	// The name of the template.
	Name pulumi.StringPtrInput
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
	// *Available only for Vault Enterprise*.
	Namespace pulumi.StringPtrInput
	// Path to where the back-end is mounted within Vault.
	Path pulumi.StringInput
	// The pattern used for matching. Currently, only regular expression pattern is supported.
	Pattern pulumi.StringPtrInput
	// The pattern type to use for match detection. Currently, only regex is supported.
	Type pulumi.StringPtrInput
}

func (TemplateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*templateArgs)(nil)).Elem()
}

type TemplateInput interface {
	pulumi.Input

	ToTemplateOutput() TemplateOutput
	ToTemplateOutputWithContext(ctx context.Context) TemplateOutput
}

func (*Template) ElementType() reflect.Type {
	return reflect.TypeOf((**Template)(nil)).Elem()
}

func (i *Template) ToTemplateOutput() TemplateOutput {
	return i.ToTemplateOutputWithContext(context.Background())
}

func (i *Template) ToTemplateOutputWithContext(ctx context.Context) TemplateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TemplateOutput)
}

// TemplateArrayInput is an input type that accepts TemplateArray and TemplateArrayOutput values.
// You can construct a concrete instance of `TemplateArrayInput` via:
//
//	TemplateArray{ TemplateArgs{...} }
type TemplateArrayInput interface {
	pulumi.Input

	ToTemplateArrayOutput() TemplateArrayOutput
	ToTemplateArrayOutputWithContext(context.Context) TemplateArrayOutput
}

type TemplateArray []TemplateInput

func (TemplateArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Template)(nil)).Elem()
}

func (i TemplateArray) ToTemplateArrayOutput() TemplateArrayOutput {
	return i.ToTemplateArrayOutputWithContext(context.Background())
}

func (i TemplateArray) ToTemplateArrayOutputWithContext(ctx context.Context) TemplateArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TemplateArrayOutput)
}

// TemplateMapInput is an input type that accepts TemplateMap and TemplateMapOutput values.
// You can construct a concrete instance of `TemplateMapInput` via:
//
//	TemplateMap{ "key": TemplateArgs{...} }
type TemplateMapInput interface {
	pulumi.Input

	ToTemplateMapOutput() TemplateMapOutput
	ToTemplateMapOutputWithContext(context.Context) TemplateMapOutput
}

type TemplateMap map[string]TemplateInput

func (TemplateMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Template)(nil)).Elem()
}

func (i TemplateMap) ToTemplateMapOutput() TemplateMapOutput {
	return i.ToTemplateMapOutputWithContext(context.Background())
}

func (i TemplateMap) ToTemplateMapOutputWithContext(ctx context.Context) TemplateMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TemplateMapOutput)
}

type TemplateOutput struct{ *pulumi.OutputState }

func (TemplateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Template)(nil)).Elem()
}

func (o TemplateOutput) ToTemplateOutput() TemplateOutput {
	return o
}

func (o TemplateOutput) ToTemplateOutputWithContext(ctx context.Context) TemplateOutput {
	return o
}

// The alphabet to use for this template. This is only used during FPE transformations.
func (o TemplateOutput) Alphabet() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Template) pulumi.StringPtrOutput { return v.Alphabet }).(pulumi.StringPtrOutput)
}

// Optional mapping of name to regular expression template, used to customize
// the decoded output. (requires Vault Enterprise 1.9+)
func (o TemplateOutput) DecodeFormats() pulumi.MapOutput {
	return o.ApplyT(func(v *Template) pulumi.MapOutput { return v.DecodeFormats }).(pulumi.MapOutput)
}

// The regular expression template used to format encoded values.
// (requires Vault Enterprise 1.9+)
func (o TemplateOutput) EncodeFormat() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Template) pulumi.StringPtrOutput { return v.EncodeFormat }).(pulumi.StringPtrOutput)
}

// The name of the template.
func (o TemplateOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Template) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The namespace to provision the resource in.
// The value should not contain leading or trailing forward slashes.
// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
// *Available only for Vault Enterprise*.
func (o TemplateOutput) Namespace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Template) pulumi.StringPtrOutput { return v.Namespace }).(pulumi.StringPtrOutput)
}

// Path to where the back-end is mounted within Vault.
func (o TemplateOutput) Path() pulumi.StringOutput {
	return o.ApplyT(func(v *Template) pulumi.StringOutput { return v.Path }).(pulumi.StringOutput)
}

// The pattern used for matching. Currently, only regular expression pattern is supported.
func (o TemplateOutput) Pattern() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Template) pulumi.StringPtrOutput { return v.Pattern }).(pulumi.StringPtrOutput)
}

// The pattern type to use for match detection. Currently, only regex is supported.
func (o TemplateOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Template) pulumi.StringPtrOutput { return v.Type }).(pulumi.StringPtrOutput)
}

type TemplateArrayOutput struct{ *pulumi.OutputState }

func (TemplateArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Template)(nil)).Elem()
}

func (o TemplateArrayOutput) ToTemplateArrayOutput() TemplateArrayOutput {
	return o
}

func (o TemplateArrayOutput) ToTemplateArrayOutputWithContext(ctx context.Context) TemplateArrayOutput {
	return o
}

func (o TemplateArrayOutput) Index(i pulumi.IntInput) TemplateOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Template {
		return vs[0].([]*Template)[vs[1].(int)]
	}).(TemplateOutput)
}

type TemplateMapOutput struct{ *pulumi.OutputState }

func (TemplateMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Template)(nil)).Elem()
}

func (o TemplateMapOutput) ToTemplateMapOutput() TemplateMapOutput {
	return o
}

func (o TemplateMapOutput) ToTemplateMapOutputWithContext(ctx context.Context) TemplateMapOutput {
	return o
}

func (o TemplateMapOutput) MapIndex(k pulumi.StringInput) TemplateOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Template {
		return vs[0].(map[string]*Template)[vs[1].(string)]
	}).(TemplateOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*TemplateInput)(nil)).Elem(), &Template{})
	pulumi.RegisterInputType(reflect.TypeOf((*TemplateArrayInput)(nil)).Elem(), TemplateArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*TemplateMapInput)(nil)).Elem(), TemplateMap{})
	pulumi.RegisterOutputType(TemplateOutput{})
	pulumi.RegisterOutputType(TemplateArrayOutput{})
	pulumi.RegisterOutputType(TemplateMapOutput{})
}
