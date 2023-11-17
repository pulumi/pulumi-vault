// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package identity

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-vault/sdk/v5/go/vault/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource for configuring the pingid MFA method.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-vault/sdk/v5/go/vault/identity"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := identity.NewMfaPingid(ctx, "example", &identity.MfaPingidArgs{
//				SettingsFileBase64: pulumi.String("CnVzZV9iYXNlNjR[...]HBtCg=="),
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
// Resource can be imported using its `uuid` field, e.g.
//
// ```sh
//
//	$ pulumi import vault:identity/mfaPingid:MfaPingid example 0d89c36a-4ff5-4d70-8749-bb6a5598aeec
//
// ```
type MfaPingid struct {
	pulumi.CustomResourceState

	// The admin URL, derived from "settingsFileBase64"
	AdminUrl pulumi.StringOutput `pulumi:"adminUrl"`
	// A unique identifier of the organization, derived from "settingsFileBase64"
	AuthenticatorUrl pulumi.StringOutput `pulumi:"authenticatorUrl"`
	// The IDP URL, derived from "settingsFileBase64"
	IdpUrl pulumi.StringOutput `pulumi:"idpUrl"`
	// Method ID.
	MethodId pulumi.StringOutput `pulumi:"methodId"`
	// Mount accessor.
	MountAccessor pulumi.StringOutput `pulumi:"mountAccessor"`
	// Method name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Target namespace. (requires Enterprise)
	Namespace pulumi.StringPtrOutput `pulumi:"namespace"`
	// Method's namespace ID.
	NamespaceId pulumi.StringOutput `pulumi:"namespaceId"`
	// Method's namespace path.
	NamespacePath pulumi.StringOutput `pulumi:"namespacePath"`
	// The name of the PingID client organization, derived from "settingsFileBase64"
	OrgAlias pulumi.StringOutput `pulumi:"orgAlias"`
	// A base64-encoded third-party settings contents as retrieved from PingID's configuration page.
	SettingsFileBase64 pulumi.StringOutput `pulumi:"settingsFileBase64"`
	// MFA type.
	Type pulumi.StringOutput `pulumi:"type"`
	// Use signature value, derived from "settingsFileBase64"
	UseSignature pulumi.BoolOutput `pulumi:"useSignature"`
	// A template string for mapping Identity names to MFA methods.
	UsernameFormat pulumi.StringPtrOutput `pulumi:"usernameFormat"`
	// Resource UUID.
	Uuid pulumi.StringOutput `pulumi:"uuid"`
}

// NewMfaPingid registers a new resource with the given unique name, arguments, and options.
func NewMfaPingid(ctx *pulumi.Context,
	name string, args *MfaPingidArgs, opts ...pulumi.ResourceOption) (*MfaPingid, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.SettingsFileBase64 == nil {
		return nil, errors.New("invalid value for required argument 'SettingsFileBase64'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource MfaPingid
	err := ctx.RegisterResource("vault:identity/mfaPingid:MfaPingid", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetMfaPingid gets an existing MfaPingid resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetMfaPingid(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MfaPingidState, opts ...pulumi.ResourceOption) (*MfaPingid, error) {
	var resource MfaPingid
	err := ctx.ReadResource("vault:identity/mfaPingid:MfaPingid", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering MfaPingid resources.
type mfaPingidState struct {
	// The admin URL, derived from "settingsFileBase64"
	AdminUrl *string `pulumi:"adminUrl"`
	// A unique identifier of the organization, derived from "settingsFileBase64"
	AuthenticatorUrl *string `pulumi:"authenticatorUrl"`
	// The IDP URL, derived from "settingsFileBase64"
	IdpUrl *string `pulumi:"idpUrl"`
	// Method ID.
	MethodId *string `pulumi:"methodId"`
	// Mount accessor.
	MountAccessor *string `pulumi:"mountAccessor"`
	// Method name.
	Name *string `pulumi:"name"`
	// Target namespace. (requires Enterprise)
	Namespace *string `pulumi:"namespace"`
	// Method's namespace ID.
	NamespaceId *string `pulumi:"namespaceId"`
	// Method's namespace path.
	NamespacePath *string `pulumi:"namespacePath"`
	// The name of the PingID client organization, derived from "settingsFileBase64"
	OrgAlias *string `pulumi:"orgAlias"`
	// A base64-encoded third-party settings contents as retrieved from PingID's configuration page.
	SettingsFileBase64 *string `pulumi:"settingsFileBase64"`
	// MFA type.
	Type *string `pulumi:"type"`
	// Use signature value, derived from "settingsFileBase64"
	UseSignature *bool `pulumi:"useSignature"`
	// A template string for mapping Identity names to MFA methods.
	UsernameFormat *string `pulumi:"usernameFormat"`
	// Resource UUID.
	Uuid *string `pulumi:"uuid"`
}

type MfaPingidState struct {
	// The admin URL, derived from "settingsFileBase64"
	AdminUrl pulumi.StringPtrInput
	// A unique identifier of the organization, derived from "settingsFileBase64"
	AuthenticatorUrl pulumi.StringPtrInput
	// The IDP URL, derived from "settingsFileBase64"
	IdpUrl pulumi.StringPtrInput
	// Method ID.
	MethodId pulumi.StringPtrInput
	// Mount accessor.
	MountAccessor pulumi.StringPtrInput
	// Method name.
	Name pulumi.StringPtrInput
	// Target namespace. (requires Enterprise)
	Namespace pulumi.StringPtrInput
	// Method's namespace ID.
	NamespaceId pulumi.StringPtrInput
	// Method's namespace path.
	NamespacePath pulumi.StringPtrInput
	// The name of the PingID client organization, derived from "settingsFileBase64"
	OrgAlias pulumi.StringPtrInput
	// A base64-encoded third-party settings contents as retrieved from PingID's configuration page.
	SettingsFileBase64 pulumi.StringPtrInput
	// MFA type.
	Type pulumi.StringPtrInput
	// Use signature value, derived from "settingsFileBase64"
	UseSignature pulumi.BoolPtrInput
	// A template string for mapping Identity names to MFA methods.
	UsernameFormat pulumi.StringPtrInput
	// Resource UUID.
	Uuid pulumi.StringPtrInput
}

func (MfaPingidState) ElementType() reflect.Type {
	return reflect.TypeOf((*mfaPingidState)(nil)).Elem()
}

type mfaPingidArgs struct {
	// Target namespace. (requires Enterprise)
	Namespace *string `pulumi:"namespace"`
	// A base64-encoded third-party settings contents as retrieved from PingID's configuration page.
	SettingsFileBase64 string `pulumi:"settingsFileBase64"`
	// A template string for mapping Identity names to MFA methods.
	UsernameFormat *string `pulumi:"usernameFormat"`
}

// The set of arguments for constructing a MfaPingid resource.
type MfaPingidArgs struct {
	// Target namespace. (requires Enterprise)
	Namespace pulumi.StringPtrInput
	// A base64-encoded third-party settings contents as retrieved from PingID's configuration page.
	SettingsFileBase64 pulumi.StringInput
	// A template string for mapping Identity names to MFA methods.
	UsernameFormat pulumi.StringPtrInput
}

func (MfaPingidArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*mfaPingidArgs)(nil)).Elem()
}

type MfaPingidInput interface {
	pulumi.Input

	ToMfaPingidOutput() MfaPingidOutput
	ToMfaPingidOutputWithContext(ctx context.Context) MfaPingidOutput
}

func (*MfaPingid) ElementType() reflect.Type {
	return reflect.TypeOf((**MfaPingid)(nil)).Elem()
}

func (i *MfaPingid) ToMfaPingidOutput() MfaPingidOutput {
	return i.ToMfaPingidOutputWithContext(context.Background())
}

func (i *MfaPingid) ToMfaPingidOutputWithContext(ctx context.Context) MfaPingidOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MfaPingidOutput)
}

// MfaPingidArrayInput is an input type that accepts MfaPingidArray and MfaPingidArrayOutput values.
// You can construct a concrete instance of `MfaPingidArrayInput` via:
//
//	MfaPingidArray{ MfaPingidArgs{...} }
type MfaPingidArrayInput interface {
	pulumi.Input

	ToMfaPingidArrayOutput() MfaPingidArrayOutput
	ToMfaPingidArrayOutputWithContext(context.Context) MfaPingidArrayOutput
}

type MfaPingidArray []MfaPingidInput

func (MfaPingidArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*MfaPingid)(nil)).Elem()
}

func (i MfaPingidArray) ToMfaPingidArrayOutput() MfaPingidArrayOutput {
	return i.ToMfaPingidArrayOutputWithContext(context.Background())
}

func (i MfaPingidArray) ToMfaPingidArrayOutputWithContext(ctx context.Context) MfaPingidArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MfaPingidArrayOutput)
}

// MfaPingidMapInput is an input type that accepts MfaPingidMap and MfaPingidMapOutput values.
// You can construct a concrete instance of `MfaPingidMapInput` via:
//
//	MfaPingidMap{ "key": MfaPingidArgs{...} }
type MfaPingidMapInput interface {
	pulumi.Input

	ToMfaPingidMapOutput() MfaPingidMapOutput
	ToMfaPingidMapOutputWithContext(context.Context) MfaPingidMapOutput
}

type MfaPingidMap map[string]MfaPingidInput

func (MfaPingidMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*MfaPingid)(nil)).Elem()
}

func (i MfaPingidMap) ToMfaPingidMapOutput() MfaPingidMapOutput {
	return i.ToMfaPingidMapOutputWithContext(context.Background())
}

func (i MfaPingidMap) ToMfaPingidMapOutputWithContext(ctx context.Context) MfaPingidMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MfaPingidMapOutput)
}

type MfaPingidOutput struct{ *pulumi.OutputState }

func (MfaPingidOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MfaPingid)(nil)).Elem()
}

func (o MfaPingidOutput) ToMfaPingidOutput() MfaPingidOutput {
	return o
}

func (o MfaPingidOutput) ToMfaPingidOutputWithContext(ctx context.Context) MfaPingidOutput {
	return o
}

// The admin URL, derived from "settingsFileBase64"
func (o MfaPingidOutput) AdminUrl() pulumi.StringOutput {
	return o.ApplyT(func(v *MfaPingid) pulumi.StringOutput { return v.AdminUrl }).(pulumi.StringOutput)
}

// A unique identifier of the organization, derived from "settingsFileBase64"
func (o MfaPingidOutput) AuthenticatorUrl() pulumi.StringOutput {
	return o.ApplyT(func(v *MfaPingid) pulumi.StringOutput { return v.AuthenticatorUrl }).(pulumi.StringOutput)
}

// The IDP URL, derived from "settingsFileBase64"
func (o MfaPingidOutput) IdpUrl() pulumi.StringOutput {
	return o.ApplyT(func(v *MfaPingid) pulumi.StringOutput { return v.IdpUrl }).(pulumi.StringOutput)
}

// Method ID.
func (o MfaPingidOutput) MethodId() pulumi.StringOutput {
	return o.ApplyT(func(v *MfaPingid) pulumi.StringOutput { return v.MethodId }).(pulumi.StringOutput)
}

// Mount accessor.
func (o MfaPingidOutput) MountAccessor() pulumi.StringOutput {
	return o.ApplyT(func(v *MfaPingid) pulumi.StringOutput { return v.MountAccessor }).(pulumi.StringOutput)
}

// Method name.
func (o MfaPingidOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *MfaPingid) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Target namespace. (requires Enterprise)
func (o MfaPingidOutput) Namespace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MfaPingid) pulumi.StringPtrOutput { return v.Namespace }).(pulumi.StringPtrOutput)
}

// Method's namespace ID.
func (o MfaPingidOutput) NamespaceId() pulumi.StringOutput {
	return o.ApplyT(func(v *MfaPingid) pulumi.StringOutput { return v.NamespaceId }).(pulumi.StringOutput)
}

// Method's namespace path.
func (o MfaPingidOutput) NamespacePath() pulumi.StringOutput {
	return o.ApplyT(func(v *MfaPingid) pulumi.StringOutput { return v.NamespacePath }).(pulumi.StringOutput)
}

// The name of the PingID client organization, derived from "settingsFileBase64"
func (o MfaPingidOutput) OrgAlias() pulumi.StringOutput {
	return o.ApplyT(func(v *MfaPingid) pulumi.StringOutput { return v.OrgAlias }).(pulumi.StringOutput)
}

// A base64-encoded third-party settings contents as retrieved from PingID's configuration page.
func (o MfaPingidOutput) SettingsFileBase64() pulumi.StringOutput {
	return o.ApplyT(func(v *MfaPingid) pulumi.StringOutput { return v.SettingsFileBase64 }).(pulumi.StringOutput)
}

// MFA type.
func (o MfaPingidOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *MfaPingid) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// Use signature value, derived from "settingsFileBase64"
func (o MfaPingidOutput) UseSignature() pulumi.BoolOutput {
	return o.ApplyT(func(v *MfaPingid) pulumi.BoolOutput { return v.UseSignature }).(pulumi.BoolOutput)
}

// A template string for mapping Identity names to MFA methods.
func (o MfaPingidOutput) UsernameFormat() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MfaPingid) pulumi.StringPtrOutput { return v.UsernameFormat }).(pulumi.StringPtrOutput)
}

// Resource UUID.
func (o MfaPingidOutput) Uuid() pulumi.StringOutput {
	return o.ApplyT(func(v *MfaPingid) pulumi.StringOutput { return v.Uuid }).(pulumi.StringOutput)
}

type MfaPingidArrayOutput struct{ *pulumi.OutputState }

func (MfaPingidArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*MfaPingid)(nil)).Elem()
}

func (o MfaPingidArrayOutput) ToMfaPingidArrayOutput() MfaPingidArrayOutput {
	return o
}

func (o MfaPingidArrayOutput) ToMfaPingidArrayOutputWithContext(ctx context.Context) MfaPingidArrayOutput {
	return o
}

func (o MfaPingidArrayOutput) Index(i pulumi.IntInput) MfaPingidOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *MfaPingid {
		return vs[0].([]*MfaPingid)[vs[1].(int)]
	}).(MfaPingidOutput)
}

type MfaPingidMapOutput struct{ *pulumi.OutputState }

func (MfaPingidMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*MfaPingid)(nil)).Elem()
}

func (o MfaPingidMapOutput) ToMfaPingidMapOutput() MfaPingidMapOutput {
	return o
}

func (o MfaPingidMapOutput) ToMfaPingidMapOutputWithContext(ctx context.Context) MfaPingidMapOutput {
	return o
}

func (o MfaPingidMapOutput) MapIndex(k pulumi.StringInput) MfaPingidOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *MfaPingid {
		return vs[0].(map[string]*MfaPingid)[vs[1].(string)]
	}).(MfaPingidOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*MfaPingidInput)(nil)).Elem(), &MfaPingid{})
	pulumi.RegisterInputType(reflect.TypeOf((*MfaPingidArrayInput)(nil)).Elem(), MfaPingidArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*MfaPingidMapInput)(nil)).Elem(), MfaPingidMap{})
	pulumi.RegisterOutputType(MfaPingidOutput{})
	pulumi.RegisterOutputType(MfaPingidArrayOutput{})
	pulumi.RegisterOutputType(MfaPingidMapOutput{})
}
