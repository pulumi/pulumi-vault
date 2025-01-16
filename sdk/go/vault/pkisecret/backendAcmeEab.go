// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package pkisecret

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-vault/sdk/v6/go/vault/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Allows creating ACME EAB (External Account Binding) tokens and deleting unused ones.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-vault/sdk/v6/go/vault"
//	"github.com/pulumi/pulumi-vault/sdk/v6/go/vault/pkisecret"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			test, err := vault.NewMount(ctx, "test", &vault.MountArgs{
//				Path:        pulumi.String("pki"),
//				Type:        pulumi.String("pki"),
//				Description: pulumi.String("PKI secret engine mount"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = pkisecret.NewBackendAcmeEab(ctx, "test", &pkisecret.BackendAcmeEabArgs{
//				Backend: test.Path,
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// ## Import
//
// # As EAB tokens are only available on initial creation there is no possibility to
//
// import or update this resource.
type BackendAcmeEab struct {
	pulumi.CustomResourceState

	// The ACME directory to which the key belongs
	AcmeDirectory pulumi.StringOutput `pulumi:"acmeDirectory"`
	// The path to the PKI secret backend to
	// create the EAB token within, with no leading or trailing `/`s.
	Backend pulumi.StringOutput `pulumi:"backend"`
	// An RFC3339 formatted date time when the EAB token was created
	CreatedOn pulumi.StringOutput `pulumi:"createdOn"`
	// The identifier of a specific ACME EAB token
	EabId pulumi.StringOutput `pulumi:"eabId"`
	// Create an EAB token that is specific to an issuer's ACME directory.
	Issuer pulumi.StringPtrOutput `pulumi:"issuer"`
	// The EAB token
	Key pulumi.StringOutput `pulumi:"key"`
	// The key type of the EAB key
	KeyType pulumi.StringOutput `pulumi:"keyType"`
	// The namespace of the target resource.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
	// *Available only for Vault Enterprise*.
	Namespace pulumi.StringPtrOutput `pulumi:"namespace"`
	// Create an EAB token that is specific to a role's ACME directory.
	//
	// **NOTE**: Within Vault ACME there are different ACME directories which an EAB token is associated with;
	//
	// 1. Default directory (`pki/acme/`) - Do not specify a value for issuer nor role parameters.
	// 2. Issuer specific (`pki/issuer/:issuer_ref/acme/`) - Specify a value for the issuer parameter
	// 3. Role specific (`pki/roles/:role/acme/`) - Specify a value for the role parameter
	// 4. Issuer and Role specific (`pki/issuer/:issuer_ref/roles/:role/acme/`) - Specify a value for both the issuer and role parameters
	Role pulumi.StringPtrOutput `pulumi:"role"`
}

// NewBackendAcmeEab registers a new resource with the given unique name, arguments, and options.
func NewBackendAcmeEab(ctx *pulumi.Context,
	name string, args *BackendAcmeEabArgs, opts ...pulumi.ResourceOption) (*BackendAcmeEab, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Backend == nil {
		return nil, errors.New("invalid value for required argument 'Backend'")
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"key",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource BackendAcmeEab
	err := ctx.RegisterResource("vault:pkiSecret/backendAcmeEab:BackendAcmeEab", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBackendAcmeEab gets an existing BackendAcmeEab resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBackendAcmeEab(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BackendAcmeEabState, opts ...pulumi.ResourceOption) (*BackendAcmeEab, error) {
	var resource BackendAcmeEab
	err := ctx.ReadResource("vault:pkiSecret/backendAcmeEab:BackendAcmeEab", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering BackendAcmeEab resources.
type backendAcmeEabState struct {
	// The ACME directory to which the key belongs
	AcmeDirectory *string `pulumi:"acmeDirectory"`
	// The path to the PKI secret backend to
	// create the EAB token within, with no leading or trailing `/`s.
	Backend *string `pulumi:"backend"`
	// An RFC3339 formatted date time when the EAB token was created
	CreatedOn *string `pulumi:"createdOn"`
	// The identifier of a specific ACME EAB token
	EabId *string `pulumi:"eabId"`
	// Create an EAB token that is specific to an issuer's ACME directory.
	Issuer *string `pulumi:"issuer"`
	// The EAB token
	Key *string `pulumi:"key"`
	// The key type of the EAB key
	KeyType *string `pulumi:"keyType"`
	// The namespace of the target resource.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
	// *Available only for Vault Enterprise*.
	Namespace *string `pulumi:"namespace"`
	// Create an EAB token that is specific to a role's ACME directory.
	//
	// **NOTE**: Within Vault ACME there are different ACME directories which an EAB token is associated with;
	//
	// 1. Default directory (`pki/acme/`) - Do not specify a value for issuer nor role parameters.
	// 2. Issuer specific (`pki/issuer/:issuer_ref/acme/`) - Specify a value for the issuer parameter
	// 3. Role specific (`pki/roles/:role/acme/`) - Specify a value for the role parameter
	// 4. Issuer and Role specific (`pki/issuer/:issuer_ref/roles/:role/acme/`) - Specify a value for both the issuer and role parameters
	Role *string `pulumi:"role"`
}

type BackendAcmeEabState struct {
	// The ACME directory to which the key belongs
	AcmeDirectory pulumi.StringPtrInput
	// The path to the PKI secret backend to
	// create the EAB token within, with no leading or trailing `/`s.
	Backend pulumi.StringPtrInput
	// An RFC3339 formatted date time when the EAB token was created
	CreatedOn pulumi.StringPtrInput
	// The identifier of a specific ACME EAB token
	EabId pulumi.StringPtrInput
	// Create an EAB token that is specific to an issuer's ACME directory.
	Issuer pulumi.StringPtrInput
	// The EAB token
	Key pulumi.StringPtrInput
	// The key type of the EAB key
	KeyType pulumi.StringPtrInput
	// The namespace of the target resource.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
	// *Available only for Vault Enterprise*.
	Namespace pulumi.StringPtrInput
	// Create an EAB token that is specific to a role's ACME directory.
	//
	// **NOTE**: Within Vault ACME there are different ACME directories which an EAB token is associated with;
	//
	// 1. Default directory (`pki/acme/`) - Do not specify a value for issuer nor role parameters.
	// 2. Issuer specific (`pki/issuer/:issuer_ref/acme/`) - Specify a value for the issuer parameter
	// 3. Role specific (`pki/roles/:role/acme/`) - Specify a value for the role parameter
	// 4. Issuer and Role specific (`pki/issuer/:issuer_ref/roles/:role/acme/`) - Specify a value for both the issuer and role parameters
	Role pulumi.StringPtrInput
}

func (BackendAcmeEabState) ElementType() reflect.Type {
	return reflect.TypeOf((*backendAcmeEabState)(nil)).Elem()
}

type backendAcmeEabArgs struct {
	// The path to the PKI secret backend to
	// create the EAB token within, with no leading or trailing `/`s.
	Backend string `pulumi:"backend"`
	// Create an EAB token that is specific to an issuer's ACME directory.
	Issuer *string `pulumi:"issuer"`
	// The namespace of the target resource.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
	// *Available only for Vault Enterprise*.
	Namespace *string `pulumi:"namespace"`
	// Create an EAB token that is specific to a role's ACME directory.
	//
	// **NOTE**: Within Vault ACME there are different ACME directories which an EAB token is associated with;
	//
	// 1. Default directory (`pki/acme/`) - Do not specify a value for issuer nor role parameters.
	// 2. Issuer specific (`pki/issuer/:issuer_ref/acme/`) - Specify a value for the issuer parameter
	// 3. Role specific (`pki/roles/:role/acme/`) - Specify a value for the role parameter
	// 4. Issuer and Role specific (`pki/issuer/:issuer_ref/roles/:role/acme/`) - Specify a value for both the issuer and role parameters
	Role *string `pulumi:"role"`
}

// The set of arguments for constructing a BackendAcmeEab resource.
type BackendAcmeEabArgs struct {
	// The path to the PKI secret backend to
	// create the EAB token within, with no leading or trailing `/`s.
	Backend pulumi.StringInput
	// Create an EAB token that is specific to an issuer's ACME directory.
	Issuer pulumi.StringPtrInput
	// The namespace of the target resource.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
	// *Available only for Vault Enterprise*.
	Namespace pulumi.StringPtrInput
	// Create an EAB token that is specific to a role's ACME directory.
	//
	// **NOTE**: Within Vault ACME there are different ACME directories which an EAB token is associated with;
	//
	// 1. Default directory (`pki/acme/`) - Do not specify a value for issuer nor role parameters.
	// 2. Issuer specific (`pki/issuer/:issuer_ref/acme/`) - Specify a value for the issuer parameter
	// 3. Role specific (`pki/roles/:role/acme/`) - Specify a value for the role parameter
	// 4. Issuer and Role specific (`pki/issuer/:issuer_ref/roles/:role/acme/`) - Specify a value for both the issuer and role parameters
	Role pulumi.StringPtrInput
}

func (BackendAcmeEabArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*backendAcmeEabArgs)(nil)).Elem()
}

type BackendAcmeEabInput interface {
	pulumi.Input

	ToBackendAcmeEabOutput() BackendAcmeEabOutput
	ToBackendAcmeEabOutputWithContext(ctx context.Context) BackendAcmeEabOutput
}

func (*BackendAcmeEab) ElementType() reflect.Type {
	return reflect.TypeOf((**BackendAcmeEab)(nil)).Elem()
}

func (i *BackendAcmeEab) ToBackendAcmeEabOutput() BackendAcmeEabOutput {
	return i.ToBackendAcmeEabOutputWithContext(context.Background())
}

func (i *BackendAcmeEab) ToBackendAcmeEabOutputWithContext(ctx context.Context) BackendAcmeEabOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BackendAcmeEabOutput)
}

// BackendAcmeEabArrayInput is an input type that accepts BackendAcmeEabArray and BackendAcmeEabArrayOutput values.
// You can construct a concrete instance of `BackendAcmeEabArrayInput` via:
//
//	BackendAcmeEabArray{ BackendAcmeEabArgs{...} }
type BackendAcmeEabArrayInput interface {
	pulumi.Input

	ToBackendAcmeEabArrayOutput() BackendAcmeEabArrayOutput
	ToBackendAcmeEabArrayOutputWithContext(context.Context) BackendAcmeEabArrayOutput
}

type BackendAcmeEabArray []BackendAcmeEabInput

func (BackendAcmeEabArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*BackendAcmeEab)(nil)).Elem()
}

func (i BackendAcmeEabArray) ToBackendAcmeEabArrayOutput() BackendAcmeEabArrayOutput {
	return i.ToBackendAcmeEabArrayOutputWithContext(context.Background())
}

func (i BackendAcmeEabArray) ToBackendAcmeEabArrayOutputWithContext(ctx context.Context) BackendAcmeEabArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BackendAcmeEabArrayOutput)
}

// BackendAcmeEabMapInput is an input type that accepts BackendAcmeEabMap and BackendAcmeEabMapOutput values.
// You can construct a concrete instance of `BackendAcmeEabMapInput` via:
//
//	BackendAcmeEabMap{ "key": BackendAcmeEabArgs{...} }
type BackendAcmeEabMapInput interface {
	pulumi.Input

	ToBackendAcmeEabMapOutput() BackendAcmeEabMapOutput
	ToBackendAcmeEabMapOutputWithContext(context.Context) BackendAcmeEabMapOutput
}

type BackendAcmeEabMap map[string]BackendAcmeEabInput

func (BackendAcmeEabMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*BackendAcmeEab)(nil)).Elem()
}

func (i BackendAcmeEabMap) ToBackendAcmeEabMapOutput() BackendAcmeEabMapOutput {
	return i.ToBackendAcmeEabMapOutputWithContext(context.Background())
}

func (i BackendAcmeEabMap) ToBackendAcmeEabMapOutputWithContext(ctx context.Context) BackendAcmeEabMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BackendAcmeEabMapOutput)
}

type BackendAcmeEabOutput struct{ *pulumi.OutputState }

func (BackendAcmeEabOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BackendAcmeEab)(nil)).Elem()
}

func (o BackendAcmeEabOutput) ToBackendAcmeEabOutput() BackendAcmeEabOutput {
	return o
}

func (o BackendAcmeEabOutput) ToBackendAcmeEabOutputWithContext(ctx context.Context) BackendAcmeEabOutput {
	return o
}

// The ACME directory to which the key belongs
func (o BackendAcmeEabOutput) AcmeDirectory() pulumi.StringOutput {
	return o.ApplyT(func(v *BackendAcmeEab) pulumi.StringOutput { return v.AcmeDirectory }).(pulumi.StringOutput)
}

// The path to the PKI secret backend to
// create the EAB token within, with no leading or trailing `/`s.
func (o BackendAcmeEabOutput) Backend() pulumi.StringOutput {
	return o.ApplyT(func(v *BackendAcmeEab) pulumi.StringOutput { return v.Backend }).(pulumi.StringOutput)
}

// An RFC3339 formatted date time when the EAB token was created
func (o BackendAcmeEabOutput) CreatedOn() pulumi.StringOutput {
	return o.ApplyT(func(v *BackendAcmeEab) pulumi.StringOutput { return v.CreatedOn }).(pulumi.StringOutput)
}

// The identifier of a specific ACME EAB token
func (o BackendAcmeEabOutput) EabId() pulumi.StringOutput {
	return o.ApplyT(func(v *BackendAcmeEab) pulumi.StringOutput { return v.EabId }).(pulumi.StringOutput)
}

// Create an EAB token that is specific to an issuer's ACME directory.
func (o BackendAcmeEabOutput) Issuer() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BackendAcmeEab) pulumi.StringPtrOutput { return v.Issuer }).(pulumi.StringPtrOutput)
}

// The EAB token
func (o BackendAcmeEabOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v *BackendAcmeEab) pulumi.StringOutput { return v.Key }).(pulumi.StringOutput)
}

// The key type of the EAB key
func (o BackendAcmeEabOutput) KeyType() pulumi.StringOutput {
	return o.ApplyT(func(v *BackendAcmeEab) pulumi.StringOutput { return v.KeyType }).(pulumi.StringOutput)
}

// The namespace of the target resource.
// The value should not contain leading or trailing forward slashes.
// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
// *Available only for Vault Enterprise*.
func (o BackendAcmeEabOutput) Namespace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BackendAcmeEab) pulumi.StringPtrOutput { return v.Namespace }).(pulumi.StringPtrOutput)
}

// Create an EAB token that is specific to a role's ACME directory.
//
// **NOTE**: Within Vault ACME there are different ACME directories which an EAB token is associated with;
//
// 1. Default directory (`pki/acme/`) - Do not specify a value for issuer nor role parameters.
// 2. Issuer specific (`pki/issuer/:issuer_ref/acme/`) - Specify a value for the issuer parameter
// 3. Role specific (`pki/roles/:role/acme/`) - Specify a value for the role parameter
// 4. Issuer and Role specific (`pki/issuer/:issuer_ref/roles/:role/acme/`) - Specify a value for both the issuer and role parameters
func (o BackendAcmeEabOutput) Role() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BackendAcmeEab) pulumi.StringPtrOutput { return v.Role }).(pulumi.StringPtrOutput)
}

type BackendAcmeEabArrayOutput struct{ *pulumi.OutputState }

func (BackendAcmeEabArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*BackendAcmeEab)(nil)).Elem()
}

func (o BackendAcmeEabArrayOutput) ToBackendAcmeEabArrayOutput() BackendAcmeEabArrayOutput {
	return o
}

func (o BackendAcmeEabArrayOutput) ToBackendAcmeEabArrayOutputWithContext(ctx context.Context) BackendAcmeEabArrayOutput {
	return o
}

func (o BackendAcmeEabArrayOutput) Index(i pulumi.IntInput) BackendAcmeEabOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *BackendAcmeEab {
		return vs[0].([]*BackendAcmeEab)[vs[1].(int)]
	}).(BackendAcmeEabOutput)
}

type BackendAcmeEabMapOutput struct{ *pulumi.OutputState }

func (BackendAcmeEabMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*BackendAcmeEab)(nil)).Elem()
}

func (o BackendAcmeEabMapOutput) ToBackendAcmeEabMapOutput() BackendAcmeEabMapOutput {
	return o
}

func (o BackendAcmeEabMapOutput) ToBackendAcmeEabMapOutputWithContext(ctx context.Context) BackendAcmeEabMapOutput {
	return o
}

func (o BackendAcmeEabMapOutput) MapIndex(k pulumi.StringInput) BackendAcmeEabOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *BackendAcmeEab {
		return vs[0].(map[string]*BackendAcmeEab)[vs[1].(string)]
	}).(BackendAcmeEabOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*BackendAcmeEabInput)(nil)).Elem(), &BackendAcmeEab{})
	pulumi.RegisterInputType(reflect.TypeOf((*BackendAcmeEabArrayInput)(nil)).Elem(), BackendAcmeEabArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*BackendAcmeEabMapInput)(nil)).Elem(), BackendAcmeEabMap{})
	pulumi.RegisterOutputType(BackendAcmeEabOutput{})
	pulumi.RegisterOutputType(BackendAcmeEabArrayOutput{})
	pulumi.RegisterOutputType(BackendAcmeEabMapOutput{})
}
