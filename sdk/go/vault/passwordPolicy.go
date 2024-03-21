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

// Provides a resource to manage Password Policies
//
// **Note** this feature is available only Vault 1.5+
//
// ## Example Usage
//
// <!--Start PulumiCodeChooser -->
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-vault/sdk/v6/go/vault"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := vault.NewPasswordPolicy(ctx, "alphanumeric", &vault.PasswordPolicyArgs{
//				Policy: pulumi.String(`    length = 20
//	    rule "charset" {
//	      charset = "abcdefghijklmnopqrstuvwxyz0123456789"
//	    }
//
// `),
//
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
// Password policies can be imported using the `name`, e.g.
//
// ```sh
// $ pulumi import vault:index/passwordPolicy:PasswordPolicy alphanumeric alphanumeric
// ```
type PasswordPolicy struct {
	pulumi.CustomResourceState

	// The name of the password policy.
	Name pulumi.StringOutput `pulumi:"name"`
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
	// *Available only for Vault Enterprise*.
	Namespace pulumi.StringPtrOutput `pulumi:"namespace"`
	// String containing a password policy.
	Policy pulumi.StringOutput `pulumi:"policy"`
}

// NewPasswordPolicy registers a new resource with the given unique name, arguments, and options.
func NewPasswordPolicy(ctx *pulumi.Context,
	name string, args *PasswordPolicyArgs, opts ...pulumi.ResourceOption) (*PasswordPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Policy == nil {
		return nil, errors.New("invalid value for required argument 'Policy'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource PasswordPolicy
	err := ctx.RegisterResource("vault:index/passwordPolicy:PasswordPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPasswordPolicy gets an existing PasswordPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPasswordPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PasswordPolicyState, opts ...pulumi.ResourceOption) (*PasswordPolicy, error) {
	var resource PasswordPolicy
	err := ctx.ReadResource("vault:index/passwordPolicy:PasswordPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PasswordPolicy resources.
type passwordPolicyState struct {
	// The name of the password policy.
	Name *string `pulumi:"name"`
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
	// *Available only for Vault Enterprise*.
	Namespace *string `pulumi:"namespace"`
	// String containing a password policy.
	Policy *string `pulumi:"policy"`
}

type PasswordPolicyState struct {
	// The name of the password policy.
	Name pulumi.StringPtrInput
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
	// *Available only for Vault Enterprise*.
	Namespace pulumi.StringPtrInput
	// String containing a password policy.
	Policy pulumi.StringPtrInput
}

func (PasswordPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*passwordPolicyState)(nil)).Elem()
}

type passwordPolicyArgs struct {
	// The name of the password policy.
	Name *string `pulumi:"name"`
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
	// *Available only for Vault Enterprise*.
	Namespace *string `pulumi:"namespace"`
	// String containing a password policy.
	Policy string `pulumi:"policy"`
}

// The set of arguments for constructing a PasswordPolicy resource.
type PasswordPolicyArgs struct {
	// The name of the password policy.
	Name pulumi.StringPtrInput
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
	// *Available only for Vault Enterprise*.
	Namespace pulumi.StringPtrInput
	// String containing a password policy.
	Policy pulumi.StringInput
}

func (PasswordPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*passwordPolicyArgs)(nil)).Elem()
}

type PasswordPolicyInput interface {
	pulumi.Input

	ToPasswordPolicyOutput() PasswordPolicyOutput
	ToPasswordPolicyOutputWithContext(ctx context.Context) PasswordPolicyOutput
}

func (*PasswordPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**PasswordPolicy)(nil)).Elem()
}

func (i *PasswordPolicy) ToPasswordPolicyOutput() PasswordPolicyOutput {
	return i.ToPasswordPolicyOutputWithContext(context.Background())
}

func (i *PasswordPolicy) ToPasswordPolicyOutputWithContext(ctx context.Context) PasswordPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PasswordPolicyOutput)
}

// PasswordPolicyArrayInput is an input type that accepts PasswordPolicyArray and PasswordPolicyArrayOutput values.
// You can construct a concrete instance of `PasswordPolicyArrayInput` via:
//
//	PasswordPolicyArray{ PasswordPolicyArgs{...} }
type PasswordPolicyArrayInput interface {
	pulumi.Input

	ToPasswordPolicyArrayOutput() PasswordPolicyArrayOutput
	ToPasswordPolicyArrayOutputWithContext(context.Context) PasswordPolicyArrayOutput
}

type PasswordPolicyArray []PasswordPolicyInput

func (PasswordPolicyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*PasswordPolicy)(nil)).Elem()
}

func (i PasswordPolicyArray) ToPasswordPolicyArrayOutput() PasswordPolicyArrayOutput {
	return i.ToPasswordPolicyArrayOutputWithContext(context.Background())
}

func (i PasswordPolicyArray) ToPasswordPolicyArrayOutputWithContext(ctx context.Context) PasswordPolicyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PasswordPolicyArrayOutput)
}

// PasswordPolicyMapInput is an input type that accepts PasswordPolicyMap and PasswordPolicyMapOutput values.
// You can construct a concrete instance of `PasswordPolicyMapInput` via:
//
//	PasswordPolicyMap{ "key": PasswordPolicyArgs{...} }
type PasswordPolicyMapInput interface {
	pulumi.Input

	ToPasswordPolicyMapOutput() PasswordPolicyMapOutput
	ToPasswordPolicyMapOutputWithContext(context.Context) PasswordPolicyMapOutput
}

type PasswordPolicyMap map[string]PasswordPolicyInput

func (PasswordPolicyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*PasswordPolicy)(nil)).Elem()
}

func (i PasswordPolicyMap) ToPasswordPolicyMapOutput() PasswordPolicyMapOutput {
	return i.ToPasswordPolicyMapOutputWithContext(context.Background())
}

func (i PasswordPolicyMap) ToPasswordPolicyMapOutputWithContext(ctx context.Context) PasswordPolicyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PasswordPolicyMapOutput)
}

type PasswordPolicyOutput struct{ *pulumi.OutputState }

func (PasswordPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PasswordPolicy)(nil)).Elem()
}

func (o PasswordPolicyOutput) ToPasswordPolicyOutput() PasswordPolicyOutput {
	return o
}

func (o PasswordPolicyOutput) ToPasswordPolicyOutputWithContext(ctx context.Context) PasswordPolicyOutput {
	return o
}

// The name of the password policy.
func (o PasswordPolicyOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *PasswordPolicy) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The namespace to provision the resource in.
// The value should not contain leading or trailing forward slashes.
// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
// *Available only for Vault Enterprise*.
func (o PasswordPolicyOutput) Namespace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PasswordPolicy) pulumi.StringPtrOutput { return v.Namespace }).(pulumi.StringPtrOutput)
}

// String containing a password policy.
func (o PasswordPolicyOutput) Policy() pulumi.StringOutput {
	return o.ApplyT(func(v *PasswordPolicy) pulumi.StringOutput { return v.Policy }).(pulumi.StringOutput)
}

type PasswordPolicyArrayOutput struct{ *pulumi.OutputState }

func (PasswordPolicyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*PasswordPolicy)(nil)).Elem()
}

func (o PasswordPolicyArrayOutput) ToPasswordPolicyArrayOutput() PasswordPolicyArrayOutput {
	return o
}

func (o PasswordPolicyArrayOutput) ToPasswordPolicyArrayOutputWithContext(ctx context.Context) PasswordPolicyArrayOutput {
	return o
}

func (o PasswordPolicyArrayOutput) Index(i pulumi.IntInput) PasswordPolicyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *PasswordPolicy {
		return vs[0].([]*PasswordPolicy)[vs[1].(int)]
	}).(PasswordPolicyOutput)
}

type PasswordPolicyMapOutput struct{ *pulumi.OutputState }

func (PasswordPolicyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*PasswordPolicy)(nil)).Elem()
}

func (o PasswordPolicyMapOutput) ToPasswordPolicyMapOutput() PasswordPolicyMapOutput {
	return o
}

func (o PasswordPolicyMapOutput) ToPasswordPolicyMapOutputWithContext(ctx context.Context) PasswordPolicyMapOutput {
	return o
}

func (o PasswordPolicyMapOutput) MapIndex(k pulumi.StringInput) PasswordPolicyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *PasswordPolicy {
		return vs[0].(map[string]*PasswordPolicy)[vs[1].(string)]
	}).(PasswordPolicyOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*PasswordPolicyInput)(nil)).Elem(), &PasswordPolicy{})
	pulumi.RegisterInputType(reflect.TypeOf((*PasswordPolicyArrayInput)(nil)).Elem(), PasswordPolicyArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*PasswordPolicyMapInput)(nil)).Elem(), PasswordPolicyMap{})
	pulumi.RegisterOutputType(PasswordPolicyOutput{})
	pulumi.RegisterOutputType(PasswordPolicyArrayOutput{})
	pulumi.RegisterOutputType(PasswordPolicyMapOutput{})
}
