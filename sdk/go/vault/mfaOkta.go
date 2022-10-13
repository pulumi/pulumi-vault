// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vault

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a resource to manage [Okta MFA](https://www.vaultproject.io/docs/enterprise/mfa/mfa-okta).
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
//	"github.com/pulumi/pulumi-vault/sdk/v5/go/vault"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			userpass, err := vault.NewAuthBackend(ctx, "userpass", &vault.AuthBackendArgs{
//				Type: pulumi.String("userpass"),
//				Path: pulumi.String("userpass"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = vault.NewMfaOkta(ctx, "myOkta", &vault.MfaOktaArgs{
//				MountAccessor:  userpass.Accessor,
//				UsernameFormat: pulumi.String("user@example.com"),
//				OrgName:        pulumi.String("hashicorp"),
//				ApiToken:       pulumi.String("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9"),
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
// Mounts can be imported using the `path`, e.g.
//
// ```sh
//
//	$ pulumi import vault:index/mfaOkta:MfaOkta my_okta my_okta
//
// ```
type MfaOkta struct {
	pulumi.CustomResourceState

	// `(string: <required>)` - Okta API key.
	ApiToken pulumi.StringOutput `pulumi:"apiToken"`
	// `(string)` - If set, will be used as the base domain for API requests. Examples are `okta.com`,
	// `oktapreview.com`, and `okta-emea.com`.
	BaseUrl pulumi.StringPtrOutput `pulumi:"baseUrl"`
	// `(string: <required>)` - The mount to tie this method to for use in automatic mappings.
	// The mapping will use the Name field of Aliases associated with this mount as the username in the mapping.
	MountAccessor pulumi.StringOutput `pulumi:"mountAccessor"`
	// `(string: <required>)` – Name of the MFA method.
	Name pulumi.StringOutput `pulumi:"name"`
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
	// *Available only for Vault Enterprise*.
	Namespace pulumi.StringPtrOutput `pulumi:"namespace"`
	// `(string: <required>)` - Name of the organization to be used in the Okta API.
	OrgName pulumi.StringOutput `pulumi:"orgName"`
	// `(string: <required>)` - If set to true, the username will only match the
	// primary email for the account.
	PrimaryEmail pulumi.BoolPtrOutput `pulumi:"primaryEmail"`
	// `(string)` - A format string for mapping Identity names to MFA method names.
	// Values to substitute should be placed in `{{}}`. For example, `"{{alias.name}}@example.com"`.
	// If blank, the Alias's Name field will be used as-is. Currently-supported mappings:
	// - alias.name: The name returned by the mount configured via the `mountAccessor` parameter
	// - entity.name: The name configured for the Entity
	// - alias.metadata.`<key>`: The value of the Alias's metadata parameter
	// - entity.metadata.`<key>`: The value of the Entity's metadata parameter
	UsernameFormat pulumi.StringPtrOutput `pulumi:"usernameFormat"`
}

// NewMfaOkta registers a new resource with the given unique name, arguments, and options.
func NewMfaOkta(ctx *pulumi.Context,
	name string, args *MfaOktaArgs, opts ...pulumi.ResourceOption) (*MfaOkta, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ApiToken == nil {
		return nil, errors.New("invalid value for required argument 'ApiToken'")
	}
	if args.MountAccessor == nil {
		return nil, errors.New("invalid value for required argument 'MountAccessor'")
	}
	if args.OrgName == nil {
		return nil, errors.New("invalid value for required argument 'OrgName'")
	}
	if args.ApiToken != nil {
		args.ApiToken = pulumi.ToSecret(args.ApiToken).(pulumi.StringOutput)
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"apiToken",
	})
	opts = append(opts, secrets)
	var resource MfaOkta
	err := ctx.RegisterResource("vault:index/mfaOkta:MfaOkta", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetMfaOkta gets an existing MfaOkta resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetMfaOkta(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MfaOktaState, opts ...pulumi.ResourceOption) (*MfaOkta, error) {
	var resource MfaOkta
	err := ctx.ReadResource("vault:index/mfaOkta:MfaOkta", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering MfaOkta resources.
type mfaOktaState struct {
	// `(string: <required>)` - Okta API key.
	ApiToken *string `pulumi:"apiToken"`
	// `(string)` - If set, will be used as the base domain for API requests. Examples are `okta.com`,
	// `oktapreview.com`, and `okta-emea.com`.
	BaseUrl *string `pulumi:"baseUrl"`
	// `(string: <required>)` - The mount to tie this method to for use in automatic mappings.
	// The mapping will use the Name field of Aliases associated with this mount as the username in the mapping.
	MountAccessor *string `pulumi:"mountAccessor"`
	// `(string: <required>)` – Name of the MFA method.
	Name *string `pulumi:"name"`
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
	// *Available only for Vault Enterprise*.
	Namespace *string `pulumi:"namespace"`
	// `(string: <required>)` - Name of the organization to be used in the Okta API.
	OrgName *string `pulumi:"orgName"`
	// `(string: <required>)` - If set to true, the username will only match the
	// primary email for the account.
	PrimaryEmail *bool `pulumi:"primaryEmail"`
	// `(string)` - A format string for mapping Identity names to MFA method names.
	// Values to substitute should be placed in `{{}}`. For example, `"{{alias.name}}@example.com"`.
	// If blank, the Alias's Name field will be used as-is. Currently-supported mappings:
	// - alias.name: The name returned by the mount configured via the `mountAccessor` parameter
	// - entity.name: The name configured for the Entity
	// - alias.metadata.`<key>`: The value of the Alias's metadata parameter
	// - entity.metadata.`<key>`: The value of the Entity's metadata parameter
	UsernameFormat *string `pulumi:"usernameFormat"`
}

type MfaOktaState struct {
	// `(string: <required>)` - Okta API key.
	ApiToken pulumi.StringPtrInput
	// `(string)` - If set, will be used as the base domain for API requests. Examples are `okta.com`,
	// `oktapreview.com`, and `okta-emea.com`.
	BaseUrl pulumi.StringPtrInput
	// `(string: <required>)` - The mount to tie this method to for use in automatic mappings.
	// The mapping will use the Name field of Aliases associated with this mount as the username in the mapping.
	MountAccessor pulumi.StringPtrInput
	// `(string: <required>)` – Name of the MFA method.
	Name pulumi.StringPtrInput
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
	// *Available only for Vault Enterprise*.
	Namespace pulumi.StringPtrInput
	// `(string: <required>)` - Name of the organization to be used in the Okta API.
	OrgName pulumi.StringPtrInput
	// `(string: <required>)` - If set to true, the username will only match the
	// primary email for the account.
	PrimaryEmail pulumi.BoolPtrInput
	// `(string)` - A format string for mapping Identity names to MFA method names.
	// Values to substitute should be placed in `{{}}`. For example, `"{{alias.name}}@example.com"`.
	// If blank, the Alias's Name field will be used as-is. Currently-supported mappings:
	// - alias.name: The name returned by the mount configured via the `mountAccessor` parameter
	// - entity.name: The name configured for the Entity
	// - alias.metadata.`<key>`: The value of the Alias's metadata parameter
	// - entity.metadata.`<key>`: The value of the Entity's metadata parameter
	UsernameFormat pulumi.StringPtrInput
}

func (MfaOktaState) ElementType() reflect.Type {
	return reflect.TypeOf((*mfaOktaState)(nil)).Elem()
}

type mfaOktaArgs struct {
	// `(string: <required>)` - Okta API key.
	ApiToken string `pulumi:"apiToken"`
	// `(string)` - If set, will be used as the base domain for API requests. Examples are `okta.com`,
	// `oktapreview.com`, and `okta-emea.com`.
	BaseUrl *string `pulumi:"baseUrl"`
	// `(string: <required>)` - The mount to tie this method to for use in automatic mappings.
	// The mapping will use the Name field of Aliases associated with this mount as the username in the mapping.
	MountAccessor string `pulumi:"mountAccessor"`
	// `(string: <required>)` – Name of the MFA method.
	Name *string `pulumi:"name"`
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
	// *Available only for Vault Enterprise*.
	Namespace *string `pulumi:"namespace"`
	// `(string: <required>)` - Name of the organization to be used in the Okta API.
	OrgName string `pulumi:"orgName"`
	// `(string: <required>)` - If set to true, the username will only match the
	// primary email for the account.
	PrimaryEmail *bool `pulumi:"primaryEmail"`
	// `(string)` - A format string for mapping Identity names to MFA method names.
	// Values to substitute should be placed in `{{}}`. For example, `"{{alias.name}}@example.com"`.
	// If blank, the Alias's Name field will be used as-is. Currently-supported mappings:
	// - alias.name: The name returned by the mount configured via the `mountAccessor` parameter
	// - entity.name: The name configured for the Entity
	// - alias.metadata.`<key>`: The value of the Alias's metadata parameter
	// - entity.metadata.`<key>`: The value of the Entity's metadata parameter
	UsernameFormat *string `pulumi:"usernameFormat"`
}

// The set of arguments for constructing a MfaOkta resource.
type MfaOktaArgs struct {
	// `(string: <required>)` - Okta API key.
	ApiToken pulumi.StringInput
	// `(string)` - If set, will be used as the base domain for API requests. Examples are `okta.com`,
	// `oktapreview.com`, and `okta-emea.com`.
	BaseUrl pulumi.StringPtrInput
	// `(string: <required>)` - The mount to tie this method to for use in automatic mappings.
	// The mapping will use the Name field of Aliases associated with this mount as the username in the mapping.
	MountAccessor pulumi.StringInput
	// `(string: <required>)` – Name of the MFA method.
	Name pulumi.StringPtrInput
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
	// *Available only for Vault Enterprise*.
	Namespace pulumi.StringPtrInput
	// `(string: <required>)` - Name of the organization to be used in the Okta API.
	OrgName pulumi.StringInput
	// `(string: <required>)` - If set to true, the username will only match the
	// primary email for the account.
	PrimaryEmail pulumi.BoolPtrInput
	// `(string)` - A format string for mapping Identity names to MFA method names.
	// Values to substitute should be placed in `{{}}`. For example, `"{{alias.name}}@example.com"`.
	// If blank, the Alias's Name field will be used as-is. Currently-supported mappings:
	// - alias.name: The name returned by the mount configured via the `mountAccessor` parameter
	// - entity.name: The name configured for the Entity
	// - alias.metadata.`<key>`: The value of the Alias's metadata parameter
	// - entity.metadata.`<key>`: The value of the Entity's metadata parameter
	UsernameFormat pulumi.StringPtrInput
}

func (MfaOktaArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*mfaOktaArgs)(nil)).Elem()
}

type MfaOktaInput interface {
	pulumi.Input

	ToMfaOktaOutput() MfaOktaOutput
	ToMfaOktaOutputWithContext(ctx context.Context) MfaOktaOutput
}

func (*MfaOkta) ElementType() reflect.Type {
	return reflect.TypeOf((**MfaOkta)(nil)).Elem()
}

func (i *MfaOkta) ToMfaOktaOutput() MfaOktaOutput {
	return i.ToMfaOktaOutputWithContext(context.Background())
}

func (i *MfaOkta) ToMfaOktaOutputWithContext(ctx context.Context) MfaOktaOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MfaOktaOutput)
}

// MfaOktaArrayInput is an input type that accepts MfaOktaArray and MfaOktaArrayOutput values.
// You can construct a concrete instance of `MfaOktaArrayInput` via:
//
//	MfaOktaArray{ MfaOktaArgs{...} }
type MfaOktaArrayInput interface {
	pulumi.Input

	ToMfaOktaArrayOutput() MfaOktaArrayOutput
	ToMfaOktaArrayOutputWithContext(context.Context) MfaOktaArrayOutput
}

type MfaOktaArray []MfaOktaInput

func (MfaOktaArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*MfaOkta)(nil)).Elem()
}

func (i MfaOktaArray) ToMfaOktaArrayOutput() MfaOktaArrayOutput {
	return i.ToMfaOktaArrayOutputWithContext(context.Background())
}

func (i MfaOktaArray) ToMfaOktaArrayOutputWithContext(ctx context.Context) MfaOktaArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MfaOktaArrayOutput)
}

// MfaOktaMapInput is an input type that accepts MfaOktaMap and MfaOktaMapOutput values.
// You can construct a concrete instance of `MfaOktaMapInput` via:
//
//	MfaOktaMap{ "key": MfaOktaArgs{...} }
type MfaOktaMapInput interface {
	pulumi.Input

	ToMfaOktaMapOutput() MfaOktaMapOutput
	ToMfaOktaMapOutputWithContext(context.Context) MfaOktaMapOutput
}

type MfaOktaMap map[string]MfaOktaInput

func (MfaOktaMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*MfaOkta)(nil)).Elem()
}

func (i MfaOktaMap) ToMfaOktaMapOutput() MfaOktaMapOutput {
	return i.ToMfaOktaMapOutputWithContext(context.Background())
}

func (i MfaOktaMap) ToMfaOktaMapOutputWithContext(ctx context.Context) MfaOktaMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MfaOktaMapOutput)
}

type MfaOktaOutput struct{ *pulumi.OutputState }

func (MfaOktaOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MfaOkta)(nil)).Elem()
}

func (o MfaOktaOutput) ToMfaOktaOutput() MfaOktaOutput {
	return o
}

func (o MfaOktaOutput) ToMfaOktaOutputWithContext(ctx context.Context) MfaOktaOutput {
	return o
}

// `(string: <required>)` - Okta API key.
func (o MfaOktaOutput) ApiToken() pulumi.StringOutput {
	return o.ApplyT(func(v *MfaOkta) pulumi.StringOutput { return v.ApiToken }).(pulumi.StringOutput)
}

// `(string)` - If set, will be used as the base domain for API requests. Examples are `okta.com`,
// `oktapreview.com`, and `okta-emea.com`.
func (o MfaOktaOutput) BaseUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MfaOkta) pulumi.StringPtrOutput { return v.BaseUrl }).(pulumi.StringPtrOutput)
}

// `(string: <required>)` - The mount to tie this method to for use in automatic mappings.
// The mapping will use the Name field of Aliases associated with this mount as the username in the mapping.
func (o MfaOktaOutput) MountAccessor() pulumi.StringOutput {
	return o.ApplyT(func(v *MfaOkta) pulumi.StringOutput { return v.MountAccessor }).(pulumi.StringOutput)
}

// `(string: <required>)` – Name of the MFA method.
func (o MfaOktaOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *MfaOkta) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The namespace to provision the resource in.
// The value should not contain leading or trailing forward slashes.
// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
// *Available only for Vault Enterprise*.
func (o MfaOktaOutput) Namespace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MfaOkta) pulumi.StringPtrOutput { return v.Namespace }).(pulumi.StringPtrOutput)
}

// `(string: <required>)` - Name of the organization to be used in the Okta API.
func (o MfaOktaOutput) OrgName() pulumi.StringOutput {
	return o.ApplyT(func(v *MfaOkta) pulumi.StringOutput { return v.OrgName }).(pulumi.StringOutput)
}

// `(string: <required>)` - If set to true, the username will only match the
// primary email for the account.
func (o MfaOktaOutput) PrimaryEmail() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *MfaOkta) pulumi.BoolPtrOutput { return v.PrimaryEmail }).(pulumi.BoolPtrOutput)
}

// `(string)` - A format string for mapping Identity names to MFA method names.
// Values to substitute should be placed in `{{}}`. For example, `"{{alias.name}}@example.com"`.
// If blank, the Alias's Name field will be used as-is. Currently-supported mappings:
// - alias.name: The name returned by the mount configured via the `mountAccessor` parameter
// - entity.name: The name configured for the Entity
// - alias.metadata.`<key>`: The value of the Alias's metadata parameter
// - entity.metadata.`<key>`: The value of the Entity's metadata parameter
func (o MfaOktaOutput) UsernameFormat() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MfaOkta) pulumi.StringPtrOutput { return v.UsernameFormat }).(pulumi.StringPtrOutput)
}

type MfaOktaArrayOutput struct{ *pulumi.OutputState }

func (MfaOktaArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*MfaOkta)(nil)).Elem()
}

func (o MfaOktaArrayOutput) ToMfaOktaArrayOutput() MfaOktaArrayOutput {
	return o
}

func (o MfaOktaArrayOutput) ToMfaOktaArrayOutputWithContext(ctx context.Context) MfaOktaArrayOutput {
	return o
}

func (o MfaOktaArrayOutput) Index(i pulumi.IntInput) MfaOktaOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *MfaOkta {
		return vs[0].([]*MfaOkta)[vs[1].(int)]
	}).(MfaOktaOutput)
}

type MfaOktaMapOutput struct{ *pulumi.OutputState }

func (MfaOktaMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*MfaOkta)(nil)).Elem()
}

func (o MfaOktaMapOutput) ToMfaOktaMapOutput() MfaOktaMapOutput {
	return o
}

func (o MfaOktaMapOutput) ToMfaOktaMapOutputWithContext(ctx context.Context) MfaOktaMapOutput {
	return o
}

func (o MfaOktaMapOutput) MapIndex(k pulumi.StringInput) MfaOktaOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *MfaOkta {
		return vs[0].(map[string]*MfaOkta)[vs[1].(string)]
	}).(MfaOktaOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*MfaOktaInput)(nil)).Elem(), &MfaOkta{})
	pulumi.RegisterInputType(reflect.TypeOf((*MfaOktaArrayInput)(nil)).Elem(), MfaOktaArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*MfaOktaMapInput)(nil)).Elem(), MfaOktaMap{})
	pulumi.RegisterOutputType(MfaOktaOutput{})
	pulumi.RegisterOutputType(MfaOktaArrayOutput{})
	pulumi.RegisterOutputType(MfaOktaMapOutput{})
}
