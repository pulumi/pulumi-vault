// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vault

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a resource to manage Endpoint Governing Policy (EGP) via [Sentinel](https://www.vaultproject.io/docs/enterprise/sentinel/index.html).
//
// **Note** this feature is available only with Vault Enterprise.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"fmt"
//
//	"github.com/pulumi/pulumi-vault/sdk/v5/go/vault"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := vault.NewEgpPolicy(ctx, "allow-all", &vault.EgpPolicyArgs{
//				EnforcementLevel: pulumi.String("soft-mandatory"),
//				Paths: pulumi.StringArray{
//					pulumi.String("*"),
//				},
//				Policy: pulumi.String(fmt.Sprintf("main = rule {\n  true\n}\n\n")),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
type EgpPolicy struct {
	pulumi.CustomResourceState

	// Enforcement level of Sentinel policy. Can be either `advisory` or `soft-mandatory` or `hard-mandatory`
	EnforcementLevel pulumi.StringOutput `pulumi:"enforcementLevel"`
	// The name of the policy
	Name pulumi.StringOutput `pulumi:"name"`
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
	// *Available only for Vault Enterprise*.
	Namespace pulumi.StringPtrOutput `pulumi:"namespace"`
	// List of paths to which the policy will be applied to
	Paths pulumi.StringArrayOutput `pulumi:"paths"`
	// String containing a Sentinel policy
	Policy pulumi.StringOutput `pulumi:"policy"`
}

// NewEgpPolicy registers a new resource with the given unique name, arguments, and options.
func NewEgpPolicy(ctx *pulumi.Context,
	name string, args *EgpPolicyArgs, opts ...pulumi.ResourceOption) (*EgpPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.EnforcementLevel == nil {
		return nil, errors.New("invalid value for required argument 'EnforcementLevel'")
	}
	if args.Paths == nil {
		return nil, errors.New("invalid value for required argument 'Paths'")
	}
	if args.Policy == nil {
		return nil, errors.New("invalid value for required argument 'Policy'")
	}
	var resource EgpPolicy
	err := ctx.RegisterResource("vault:index/egpPolicy:EgpPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetEgpPolicy gets an existing EgpPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEgpPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EgpPolicyState, opts ...pulumi.ResourceOption) (*EgpPolicy, error) {
	var resource EgpPolicy
	err := ctx.ReadResource("vault:index/egpPolicy:EgpPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering EgpPolicy resources.
type egpPolicyState struct {
	// Enforcement level of Sentinel policy. Can be either `advisory` or `soft-mandatory` or `hard-mandatory`
	EnforcementLevel *string `pulumi:"enforcementLevel"`
	// The name of the policy
	Name *string `pulumi:"name"`
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
	// *Available only for Vault Enterprise*.
	Namespace *string `pulumi:"namespace"`
	// List of paths to which the policy will be applied to
	Paths []string `pulumi:"paths"`
	// String containing a Sentinel policy
	Policy *string `pulumi:"policy"`
}

type EgpPolicyState struct {
	// Enforcement level of Sentinel policy. Can be either `advisory` or `soft-mandatory` or `hard-mandatory`
	EnforcementLevel pulumi.StringPtrInput
	// The name of the policy
	Name pulumi.StringPtrInput
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
	// *Available only for Vault Enterprise*.
	Namespace pulumi.StringPtrInput
	// List of paths to which the policy will be applied to
	Paths pulumi.StringArrayInput
	// String containing a Sentinel policy
	Policy pulumi.StringPtrInput
}

func (EgpPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*egpPolicyState)(nil)).Elem()
}

type egpPolicyArgs struct {
	// Enforcement level of Sentinel policy. Can be either `advisory` or `soft-mandatory` or `hard-mandatory`
	EnforcementLevel string `pulumi:"enforcementLevel"`
	// The name of the policy
	Name *string `pulumi:"name"`
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
	// *Available only for Vault Enterprise*.
	Namespace *string `pulumi:"namespace"`
	// List of paths to which the policy will be applied to
	Paths []string `pulumi:"paths"`
	// String containing a Sentinel policy
	Policy string `pulumi:"policy"`
}

// The set of arguments for constructing a EgpPolicy resource.
type EgpPolicyArgs struct {
	// Enforcement level of Sentinel policy. Can be either `advisory` or `soft-mandatory` or `hard-mandatory`
	EnforcementLevel pulumi.StringInput
	// The name of the policy
	Name pulumi.StringPtrInput
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
	// *Available only for Vault Enterprise*.
	Namespace pulumi.StringPtrInput
	// List of paths to which the policy will be applied to
	Paths pulumi.StringArrayInput
	// String containing a Sentinel policy
	Policy pulumi.StringInput
}

func (EgpPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*egpPolicyArgs)(nil)).Elem()
}

type EgpPolicyInput interface {
	pulumi.Input

	ToEgpPolicyOutput() EgpPolicyOutput
	ToEgpPolicyOutputWithContext(ctx context.Context) EgpPolicyOutput
}

func (*EgpPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**EgpPolicy)(nil)).Elem()
}

func (i *EgpPolicy) ToEgpPolicyOutput() EgpPolicyOutput {
	return i.ToEgpPolicyOutputWithContext(context.Background())
}

func (i *EgpPolicy) ToEgpPolicyOutputWithContext(ctx context.Context) EgpPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EgpPolicyOutput)
}

// EgpPolicyArrayInput is an input type that accepts EgpPolicyArray and EgpPolicyArrayOutput values.
// You can construct a concrete instance of `EgpPolicyArrayInput` via:
//
//	EgpPolicyArray{ EgpPolicyArgs{...} }
type EgpPolicyArrayInput interface {
	pulumi.Input

	ToEgpPolicyArrayOutput() EgpPolicyArrayOutput
	ToEgpPolicyArrayOutputWithContext(context.Context) EgpPolicyArrayOutput
}

type EgpPolicyArray []EgpPolicyInput

func (EgpPolicyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*EgpPolicy)(nil)).Elem()
}

func (i EgpPolicyArray) ToEgpPolicyArrayOutput() EgpPolicyArrayOutput {
	return i.ToEgpPolicyArrayOutputWithContext(context.Background())
}

func (i EgpPolicyArray) ToEgpPolicyArrayOutputWithContext(ctx context.Context) EgpPolicyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EgpPolicyArrayOutput)
}

// EgpPolicyMapInput is an input type that accepts EgpPolicyMap and EgpPolicyMapOutput values.
// You can construct a concrete instance of `EgpPolicyMapInput` via:
//
//	EgpPolicyMap{ "key": EgpPolicyArgs{...} }
type EgpPolicyMapInput interface {
	pulumi.Input

	ToEgpPolicyMapOutput() EgpPolicyMapOutput
	ToEgpPolicyMapOutputWithContext(context.Context) EgpPolicyMapOutput
}

type EgpPolicyMap map[string]EgpPolicyInput

func (EgpPolicyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*EgpPolicy)(nil)).Elem()
}

func (i EgpPolicyMap) ToEgpPolicyMapOutput() EgpPolicyMapOutput {
	return i.ToEgpPolicyMapOutputWithContext(context.Background())
}

func (i EgpPolicyMap) ToEgpPolicyMapOutputWithContext(ctx context.Context) EgpPolicyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EgpPolicyMapOutput)
}

type EgpPolicyOutput struct{ *pulumi.OutputState }

func (EgpPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EgpPolicy)(nil)).Elem()
}

func (o EgpPolicyOutput) ToEgpPolicyOutput() EgpPolicyOutput {
	return o
}

func (o EgpPolicyOutput) ToEgpPolicyOutputWithContext(ctx context.Context) EgpPolicyOutput {
	return o
}

// Enforcement level of Sentinel policy. Can be either `advisory` or `soft-mandatory` or `hard-mandatory`
func (o EgpPolicyOutput) EnforcementLevel() pulumi.StringOutput {
	return o.ApplyT(func(v *EgpPolicy) pulumi.StringOutput { return v.EnforcementLevel }).(pulumi.StringOutput)
}

// The name of the policy
func (o EgpPolicyOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *EgpPolicy) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The namespace to provision the resource in.
// The value should not contain leading or trailing forward slashes.
// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
// *Available only for Vault Enterprise*.
func (o EgpPolicyOutput) Namespace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EgpPolicy) pulumi.StringPtrOutput { return v.Namespace }).(pulumi.StringPtrOutput)
}

// List of paths to which the policy will be applied to
func (o EgpPolicyOutput) Paths() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *EgpPolicy) pulumi.StringArrayOutput { return v.Paths }).(pulumi.StringArrayOutput)
}

// String containing a Sentinel policy
func (o EgpPolicyOutput) Policy() pulumi.StringOutput {
	return o.ApplyT(func(v *EgpPolicy) pulumi.StringOutput { return v.Policy }).(pulumi.StringOutput)
}

type EgpPolicyArrayOutput struct{ *pulumi.OutputState }

func (EgpPolicyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*EgpPolicy)(nil)).Elem()
}

func (o EgpPolicyArrayOutput) ToEgpPolicyArrayOutput() EgpPolicyArrayOutput {
	return o
}

func (o EgpPolicyArrayOutput) ToEgpPolicyArrayOutputWithContext(ctx context.Context) EgpPolicyArrayOutput {
	return o
}

func (o EgpPolicyArrayOutput) Index(i pulumi.IntInput) EgpPolicyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *EgpPolicy {
		return vs[0].([]*EgpPolicy)[vs[1].(int)]
	}).(EgpPolicyOutput)
}

type EgpPolicyMapOutput struct{ *pulumi.OutputState }

func (EgpPolicyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*EgpPolicy)(nil)).Elem()
}

func (o EgpPolicyMapOutput) ToEgpPolicyMapOutput() EgpPolicyMapOutput {
	return o
}

func (o EgpPolicyMapOutput) ToEgpPolicyMapOutputWithContext(ctx context.Context) EgpPolicyMapOutput {
	return o
}

func (o EgpPolicyMapOutput) MapIndex(k pulumi.StringInput) EgpPolicyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *EgpPolicy {
		return vs[0].(map[string]*EgpPolicy)[vs[1].(string)]
	}).(EgpPolicyOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*EgpPolicyInput)(nil)).Elem(), &EgpPolicy{})
	pulumi.RegisterInputType(reflect.TypeOf((*EgpPolicyArrayInput)(nil)).Elem(), EgpPolicyArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*EgpPolicyMapInput)(nil)).Elem(), EgpPolicyMap{})
	pulumi.RegisterOutputType(EgpPolicyOutput{})
	pulumi.RegisterOutputType(EgpPolicyArrayOutput{})
	pulumi.RegisterOutputType(EgpPolicyMapOutput{})
}
