// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package transform

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-vault/sdk/v6/go/vault/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source supports the "/transform/decode/{role_name}" Vault endpoint.
//
// It decodes the provided value using a named role.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-vault/sdk/v6/go/vault"
//	"github.com/pulumi/pulumi-vault/sdk/v6/go/vault/transform"
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
//			_, err = transform.NewTransformation(ctx, "ccn-fpe", &transform.TransformationArgs{
//				Path:        transform.Path,
//				Name:        pulumi.String("ccn-fpe"),
//				Type:        pulumi.String("fpe"),
//				Template:    pulumi.String("builtin/creditcardnumber"),
//				TweakSource: pulumi.String("internal"),
//				AllowedRoles: pulumi.StringArray{
//					pulumi.String("payments"),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			payments, err := transform.NewRole(ctx, "payments", &transform.RoleArgs{
//				Path: ccn_fpe.Path,
//				Name: pulumi.String("payments"),
//				Transformations: pulumi.StringArray{
//					pulumi.String("ccn-fpe"),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			_ = transform.GetDecodeOutput(ctx, transform.GetDecodeOutputArgs{
//				Path:     payments.Path,
//				RoleName: pulumi.String("payments"),
//				Value:    pulumi.String("9300-3376-4943-8903"),
//			}, nil)
//			return nil
//		})
//	}
//
// ```
func GetDecode(ctx *pulumi.Context, args *GetDecodeArgs, opts ...pulumi.InvokeOption) (*GetDecodeResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetDecodeResult
	err := ctx.Invoke("vault:transform/getDecode:getDecode", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getDecode.
type GetDecodeArgs struct {
	// Specifies a list of items to be decoded in a single batch. If this parameter is set, the top-level parameters 'value', 'transformation' and 'tweak' will be ignored. Each batch item within the list can specify these parameters instead.
	BatchInputs []map[string]string `pulumi:"batchInputs"`
	// The result of decoding a batch.
	BatchResults []map[string]string `pulumi:"batchResults"`
	// The result of decoding a value.
	DecodedValue *string `pulumi:"decodedValue"`
	// The namespace of the target resource.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
	// *Available only for Vault Enterprise*.
	Namespace *string `pulumi:"namespace"`
	// Path to where the back-end is mounted within Vault.
	Path string `pulumi:"path"`
	// The name of the role.
	RoleName string `pulumi:"roleName"`
	// The transformation to perform. If no value is provided and the role contains a single transformation, this value will be inferred from the role.
	Transformation *string `pulumi:"transformation"`
	// The tweak value to use. Only applicable for FPE transformations
	Tweak *string `pulumi:"tweak"`
	// The value in which to decode.
	Value *string `pulumi:"value"`
}

// A collection of values returned by getDecode.
type GetDecodeResult struct {
	BatchInputs  []map[string]string `pulumi:"batchInputs"`
	BatchResults []map[string]string `pulumi:"batchResults"`
	DecodedValue string              `pulumi:"decodedValue"`
	// The provider-assigned unique ID for this managed resource.
	Id             string  `pulumi:"id"`
	Namespace      *string `pulumi:"namespace"`
	Path           string  `pulumi:"path"`
	RoleName       string  `pulumi:"roleName"`
	Transformation *string `pulumi:"transformation"`
	Tweak          *string `pulumi:"tweak"`
	Value          *string `pulumi:"value"`
}

func GetDecodeOutput(ctx *pulumi.Context, args GetDecodeOutputArgs, opts ...pulumi.InvokeOption) GetDecodeResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (GetDecodeResultOutput, error) {
			args := v.(GetDecodeArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("vault:transform/getDecode:getDecode", args, GetDecodeResultOutput{}, options).(GetDecodeResultOutput), nil
		}).(GetDecodeResultOutput)
}

// A collection of arguments for invoking getDecode.
type GetDecodeOutputArgs struct {
	// Specifies a list of items to be decoded in a single batch. If this parameter is set, the top-level parameters 'value', 'transformation' and 'tweak' will be ignored. Each batch item within the list can specify these parameters instead.
	BatchInputs pulumi.StringMapArrayInput `pulumi:"batchInputs"`
	// The result of decoding a batch.
	BatchResults pulumi.StringMapArrayInput `pulumi:"batchResults"`
	// The result of decoding a value.
	DecodedValue pulumi.StringPtrInput `pulumi:"decodedValue"`
	// The namespace of the target resource.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
	// *Available only for Vault Enterprise*.
	Namespace pulumi.StringPtrInput `pulumi:"namespace"`
	// Path to where the back-end is mounted within Vault.
	Path pulumi.StringInput `pulumi:"path"`
	// The name of the role.
	RoleName pulumi.StringInput `pulumi:"roleName"`
	// The transformation to perform. If no value is provided and the role contains a single transformation, this value will be inferred from the role.
	Transformation pulumi.StringPtrInput `pulumi:"transformation"`
	// The tweak value to use. Only applicable for FPE transformations
	Tweak pulumi.StringPtrInput `pulumi:"tweak"`
	// The value in which to decode.
	Value pulumi.StringPtrInput `pulumi:"value"`
}

func (GetDecodeOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetDecodeArgs)(nil)).Elem()
}

// A collection of values returned by getDecode.
type GetDecodeResultOutput struct{ *pulumi.OutputState }

func (GetDecodeResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetDecodeResult)(nil)).Elem()
}

func (o GetDecodeResultOutput) ToGetDecodeResultOutput() GetDecodeResultOutput {
	return o
}

func (o GetDecodeResultOutput) ToGetDecodeResultOutputWithContext(ctx context.Context) GetDecodeResultOutput {
	return o
}

func (o GetDecodeResultOutput) BatchInputs() pulumi.StringMapArrayOutput {
	return o.ApplyT(func(v GetDecodeResult) []map[string]string { return v.BatchInputs }).(pulumi.StringMapArrayOutput)
}

func (o GetDecodeResultOutput) BatchResults() pulumi.StringMapArrayOutput {
	return o.ApplyT(func(v GetDecodeResult) []map[string]string { return v.BatchResults }).(pulumi.StringMapArrayOutput)
}

func (o GetDecodeResultOutput) DecodedValue() pulumi.StringOutput {
	return o.ApplyT(func(v GetDecodeResult) string { return v.DecodedValue }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetDecodeResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetDecodeResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetDecodeResultOutput) Namespace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetDecodeResult) *string { return v.Namespace }).(pulumi.StringPtrOutput)
}

func (o GetDecodeResultOutput) Path() pulumi.StringOutput {
	return o.ApplyT(func(v GetDecodeResult) string { return v.Path }).(pulumi.StringOutput)
}

func (o GetDecodeResultOutput) RoleName() pulumi.StringOutput {
	return o.ApplyT(func(v GetDecodeResult) string { return v.RoleName }).(pulumi.StringOutput)
}

func (o GetDecodeResultOutput) Transformation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetDecodeResult) *string { return v.Transformation }).(pulumi.StringPtrOutput)
}

func (o GetDecodeResultOutput) Tweak() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetDecodeResult) *string { return v.Tweak }).(pulumi.StringPtrOutput)
}

func (o GetDecodeResultOutput) Value() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetDecodeResult) *string { return v.Value }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetDecodeResultOutput{})
}
