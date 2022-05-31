// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vault

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a resource to manage [TOTP MFA](https://www.vaultproject.io/docs/enterprise/mfa/mfa-totp).
//
// **Note** this feature is available only with Vault Enterprise.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-vault/sdk/v5/go/vault"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := vault.NewMfaTotp(ctx, "myTotp", &vault.MfaTotpArgs{
// 			Algorithm: pulumi.String("SHA256"),
// 			Digits:    pulumi.Int(8),
// 			Issuer:    pulumi.String("hashicorp"),
// 			KeySize:   pulumi.Int(20),
// 			Period:    pulumi.Int(60),
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
// Mounts can be imported using the `path`, e.g.
//
// ```sh
//  $ pulumi import vault:index/mfaTotp:MfaTotp my_totp my_totp
// ```
type MfaTotp struct {
	pulumi.CustomResourceState

	// `(string)` - Specifies the hashing algorithm used to generate the TOTP code.
	// Options include `SHA1`, `SHA256` and `SHA512`
	Algorithm pulumi.StringPtrOutput `pulumi:"algorithm"`
	// `(int)` - The number of digits in the generated TOTP token.
	// This value can either be 6 or 8.
	Digits pulumi.IntPtrOutput `pulumi:"digits"`
	// `(string: <required>)` - The name of the key's issuing organization.
	Issuer pulumi.StringOutput `pulumi:"issuer"`
	// `(int)` - Specifies the size in bytes of the generated key.
	KeySize pulumi.IntPtrOutput `pulumi:"keySize"`
	// `(string: <required>)` – Name of the MFA method.
	Name pulumi.StringOutput `pulumi:"name"`
	// `(int)` - The length of time used to generate a counter for the TOTP token calculation.
	Period pulumi.IntPtrOutput `pulumi:"period"`
	// `(int)` - The pixel size of the generated square QR code.
	QrSize pulumi.IntPtrOutput `pulumi:"qrSize"`
	// `(int)` - The number of delay periods that are allowed when validating a TOTP token.
	// This value can either be 0 or 1.
	Skew pulumi.IntPtrOutput `pulumi:"skew"`
}

// NewMfaTotp registers a new resource with the given unique name, arguments, and options.
func NewMfaTotp(ctx *pulumi.Context,
	name string, args *MfaTotpArgs, opts ...pulumi.ResourceOption) (*MfaTotp, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Issuer == nil {
		return nil, errors.New("invalid value for required argument 'Issuer'")
	}
	var resource MfaTotp
	err := ctx.RegisterResource("vault:index/mfaTotp:MfaTotp", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetMfaTotp gets an existing MfaTotp resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetMfaTotp(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MfaTotpState, opts ...pulumi.ResourceOption) (*MfaTotp, error) {
	var resource MfaTotp
	err := ctx.ReadResource("vault:index/mfaTotp:MfaTotp", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering MfaTotp resources.
type mfaTotpState struct {
	// `(string)` - Specifies the hashing algorithm used to generate the TOTP code.
	// Options include `SHA1`, `SHA256` and `SHA512`
	Algorithm *string `pulumi:"algorithm"`
	// `(int)` - The number of digits in the generated TOTP token.
	// This value can either be 6 or 8.
	Digits *int `pulumi:"digits"`
	// `(string: <required>)` - The name of the key's issuing organization.
	Issuer *string `pulumi:"issuer"`
	// `(int)` - Specifies the size in bytes of the generated key.
	KeySize *int `pulumi:"keySize"`
	// `(string: <required>)` – Name of the MFA method.
	Name *string `pulumi:"name"`
	// `(int)` - The length of time used to generate a counter for the TOTP token calculation.
	Period *int `pulumi:"period"`
	// `(int)` - The pixel size of the generated square QR code.
	QrSize *int `pulumi:"qrSize"`
	// `(int)` - The number of delay periods that are allowed when validating a TOTP token.
	// This value can either be 0 or 1.
	Skew *int `pulumi:"skew"`
}

type MfaTotpState struct {
	// `(string)` - Specifies the hashing algorithm used to generate the TOTP code.
	// Options include `SHA1`, `SHA256` and `SHA512`
	Algorithm pulumi.StringPtrInput
	// `(int)` - The number of digits in the generated TOTP token.
	// This value can either be 6 or 8.
	Digits pulumi.IntPtrInput
	// `(string: <required>)` - The name of the key's issuing organization.
	Issuer pulumi.StringPtrInput
	// `(int)` - Specifies the size in bytes of the generated key.
	KeySize pulumi.IntPtrInput
	// `(string: <required>)` – Name of the MFA method.
	Name pulumi.StringPtrInput
	// `(int)` - The length of time used to generate a counter for the TOTP token calculation.
	Period pulumi.IntPtrInput
	// `(int)` - The pixel size of the generated square QR code.
	QrSize pulumi.IntPtrInput
	// `(int)` - The number of delay periods that are allowed when validating a TOTP token.
	// This value can either be 0 or 1.
	Skew pulumi.IntPtrInput
}

func (MfaTotpState) ElementType() reflect.Type {
	return reflect.TypeOf((*mfaTotpState)(nil)).Elem()
}

type mfaTotpArgs struct {
	// `(string)` - Specifies the hashing algorithm used to generate the TOTP code.
	// Options include `SHA1`, `SHA256` and `SHA512`
	Algorithm *string `pulumi:"algorithm"`
	// `(int)` - The number of digits in the generated TOTP token.
	// This value can either be 6 or 8.
	Digits *int `pulumi:"digits"`
	// `(string: <required>)` - The name of the key's issuing organization.
	Issuer string `pulumi:"issuer"`
	// `(int)` - Specifies the size in bytes of the generated key.
	KeySize *int `pulumi:"keySize"`
	// `(string: <required>)` – Name of the MFA method.
	Name *string `pulumi:"name"`
	// `(int)` - The length of time used to generate a counter for the TOTP token calculation.
	Period *int `pulumi:"period"`
	// `(int)` - The pixel size of the generated square QR code.
	QrSize *int `pulumi:"qrSize"`
	// `(int)` - The number of delay periods that are allowed when validating a TOTP token.
	// This value can either be 0 or 1.
	Skew *int `pulumi:"skew"`
}

// The set of arguments for constructing a MfaTotp resource.
type MfaTotpArgs struct {
	// `(string)` - Specifies the hashing algorithm used to generate the TOTP code.
	// Options include `SHA1`, `SHA256` and `SHA512`
	Algorithm pulumi.StringPtrInput
	// `(int)` - The number of digits in the generated TOTP token.
	// This value can either be 6 or 8.
	Digits pulumi.IntPtrInput
	// `(string: <required>)` - The name of the key's issuing organization.
	Issuer pulumi.StringInput
	// `(int)` - Specifies the size in bytes of the generated key.
	KeySize pulumi.IntPtrInput
	// `(string: <required>)` – Name of the MFA method.
	Name pulumi.StringPtrInput
	// `(int)` - The length of time used to generate a counter for the TOTP token calculation.
	Period pulumi.IntPtrInput
	// `(int)` - The pixel size of the generated square QR code.
	QrSize pulumi.IntPtrInput
	// `(int)` - The number of delay periods that are allowed when validating a TOTP token.
	// This value can either be 0 or 1.
	Skew pulumi.IntPtrInput
}

func (MfaTotpArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*mfaTotpArgs)(nil)).Elem()
}

type MfaTotpInput interface {
	pulumi.Input

	ToMfaTotpOutput() MfaTotpOutput
	ToMfaTotpOutputWithContext(ctx context.Context) MfaTotpOutput
}

func (*MfaTotp) ElementType() reflect.Type {
	return reflect.TypeOf((**MfaTotp)(nil)).Elem()
}

func (i *MfaTotp) ToMfaTotpOutput() MfaTotpOutput {
	return i.ToMfaTotpOutputWithContext(context.Background())
}

func (i *MfaTotp) ToMfaTotpOutputWithContext(ctx context.Context) MfaTotpOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MfaTotpOutput)
}

// MfaTotpArrayInput is an input type that accepts MfaTotpArray and MfaTotpArrayOutput values.
// You can construct a concrete instance of `MfaTotpArrayInput` via:
//
//          MfaTotpArray{ MfaTotpArgs{...} }
type MfaTotpArrayInput interface {
	pulumi.Input

	ToMfaTotpArrayOutput() MfaTotpArrayOutput
	ToMfaTotpArrayOutputWithContext(context.Context) MfaTotpArrayOutput
}

type MfaTotpArray []MfaTotpInput

func (MfaTotpArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*MfaTotp)(nil)).Elem()
}

func (i MfaTotpArray) ToMfaTotpArrayOutput() MfaTotpArrayOutput {
	return i.ToMfaTotpArrayOutputWithContext(context.Background())
}

func (i MfaTotpArray) ToMfaTotpArrayOutputWithContext(ctx context.Context) MfaTotpArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MfaTotpArrayOutput)
}

// MfaTotpMapInput is an input type that accepts MfaTotpMap and MfaTotpMapOutput values.
// You can construct a concrete instance of `MfaTotpMapInput` via:
//
//          MfaTotpMap{ "key": MfaTotpArgs{...} }
type MfaTotpMapInput interface {
	pulumi.Input

	ToMfaTotpMapOutput() MfaTotpMapOutput
	ToMfaTotpMapOutputWithContext(context.Context) MfaTotpMapOutput
}

type MfaTotpMap map[string]MfaTotpInput

func (MfaTotpMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*MfaTotp)(nil)).Elem()
}

func (i MfaTotpMap) ToMfaTotpMapOutput() MfaTotpMapOutput {
	return i.ToMfaTotpMapOutputWithContext(context.Background())
}

func (i MfaTotpMap) ToMfaTotpMapOutputWithContext(ctx context.Context) MfaTotpMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MfaTotpMapOutput)
}

type MfaTotpOutput struct{ *pulumi.OutputState }

func (MfaTotpOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MfaTotp)(nil)).Elem()
}

func (o MfaTotpOutput) ToMfaTotpOutput() MfaTotpOutput {
	return o
}

func (o MfaTotpOutput) ToMfaTotpOutputWithContext(ctx context.Context) MfaTotpOutput {
	return o
}

type MfaTotpArrayOutput struct{ *pulumi.OutputState }

func (MfaTotpArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*MfaTotp)(nil)).Elem()
}

func (o MfaTotpArrayOutput) ToMfaTotpArrayOutput() MfaTotpArrayOutput {
	return o
}

func (o MfaTotpArrayOutput) ToMfaTotpArrayOutputWithContext(ctx context.Context) MfaTotpArrayOutput {
	return o
}

func (o MfaTotpArrayOutput) Index(i pulumi.IntInput) MfaTotpOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *MfaTotp {
		return vs[0].([]*MfaTotp)[vs[1].(int)]
	}).(MfaTotpOutput)
}

type MfaTotpMapOutput struct{ *pulumi.OutputState }

func (MfaTotpMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*MfaTotp)(nil)).Elem()
}

func (o MfaTotpMapOutput) ToMfaTotpMapOutput() MfaTotpMapOutput {
	return o
}

func (o MfaTotpMapOutput) ToMfaTotpMapOutputWithContext(ctx context.Context) MfaTotpMapOutput {
	return o
}

func (o MfaTotpMapOutput) MapIndex(k pulumi.StringInput) MfaTotpOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *MfaTotp {
		return vs[0].(map[string]*MfaTotp)[vs[1].(string)]
	}).(MfaTotpOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*MfaTotpInput)(nil)).Elem(), &MfaTotp{})
	pulumi.RegisterInputType(reflect.TypeOf((*MfaTotpArrayInput)(nil)).Elem(), MfaTotpArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*MfaTotpMapInput)(nil)).Elem(), MfaTotpMap{})
	pulumi.RegisterOutputType(MfaTotpOutput{})
	pulumi.RegisterOutputType(MfaTotpArrayOutput{})
	pulumi.RegisterOutputType(MfaTotpMapOutput{})
}
