// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vault

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a resource to manage [PingID MFA](https://www.vaultproject.io/docs/enterprise/mfa/mfa-pingid).
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
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		cfg := config.New(ctx, "")
// 		settingsFile := cfg.RequireObject("settingsFile")
// 		userpass, err := vault.NewAuthBackend(ctx, "userpass", &vault.AuthBackendArgs{
// 			Type: pulumi.String("userpass"),
// 			Path: pulumi.String("userpass"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = vault.NewMfaPingid(ctx, "myPingid", &vault.MfaPingidArgs{
// 			MountAccessor:      userpass.Accessor,
// 			UsernameFormat:     pulumi.String("user@example.com"),
// 			SettingsFileBase64: pulumi.Any(settingsFile),
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
//  $ pulumi import vault:index/mfaPingid:MfaPingid my_pingid my_pingid
// ```
type MfaPingid struct {
	pulumi.CustomResourceState

	// Admin URL computed by Vault.
	AdminUrl pulumi.StringOutput `pulumi:"adminUrl"`
	// Authenticator URL computed by Vault.
	AuthenticatorUrl pulumi.StringOutput `pulumi:"authenticatorUrl"`
	// IDP URL computed by Vault.
	IdpUrl pulumi.StringOutput `pulumi:"idpUrl"`
	// `(string: <required>)` - The mount to tie this method to for use in automatic mappings.
	// The mapping will use the Name field of Aliases associated with this mount as the username in the mapping.
	MountAccessor pulumi.StringOutput `pulumi:"mountAccessor"`
	// `(string: <required>)` – Name of the MFA method.
	Name pulumi.StringOutput `pulumi:"name"`
	// Namespace ID computed by Vault.
	NamespaceId pulumi.StringOutput `pulumi:"namespaceId"`
	// Org Alias computed by Vault.
	OrgAlias pulumi.StringOutput `pulumi:"orgAlias"`
	// `(string: <required>)` - A base64-encoded third-party settings file retrieved
	// from PingID's configuration page.
	SettingsFileBase64 pulumi.StringOutput `pulumi:"settingsFileBase64"`
	// Type of configuration computed by Vault.
	Type pulumi.StringOutput `pulumi:"type"`
	// If set, enables use of PingID signature. Computed by Vault
	UseSignature pulumi.BoolOutput `pulumi:"useSignature"`
	// `(string)` - A format string for mapping Identity names to MFA method names.
	// Values to substitute should be placed in `{{}}`. For example, `"{{alias.name}}@example.com"`.
	// If blank, the Alias's Name field will be used as-is. Currently-supported mappings:
	// - alias.name: The name returned by the mount configured via the `mountAccessor` parameter
	// - entity.name: The name configured for the Entity
	// - alias.metadata.`<key>`: The value of the Alias's metadata parameter
	// - entity.metadata.`<key>`: The value of the Entity's metadata parameter
	UsernameFormat pulumi.StringPtrOutput `pulumi:"usernameFormat"`
}

// NewMfaPingid registers a new resource with the given unique name, arguments, and options.
func NewMfaPingid(ctx *pulumi.Context,
	name string, args *MfaPingidArgs, opts ...pulumi.ResourceOption) (*MfaPingid, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.MountAccessor == nil {
		return nil, errors.New("invalid value for required argument 'MountAccessor'")
	}
	if args.SettingsFileBase64 == nil {
		return nil, errors.New("invalid value for required argument 'SettingsFileBase64'")
	}
	var resource MfaPingid
	err := ctx.RegisterResource("vault:index/mfaPingid:MfaPingid", name, args, &resource, opts...)
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
	err := ctx.ReadResource("vault:index/mfaPingid:MfaPingid", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering MfaPingid resources.
type mfaPingidState struct {
	// Admin URL computed by Vault.
	AdminUrl *string `pulumi:"adminUrl"`
	// Authenticator URL computed by Vault.
	AuthenticatorUrl *string `pulumi:"authenticatorUrl"`
	// IDP URL computed by Vault.
	IdpUrl *string `pulumi:"idpUrl"`
	// `(string: <required>)` - The mount to tie this method to for use in automatic mappings.
	// The mapping will use the Name field of Aliases associated with this mount as the username in the mapping.
	MountAccessor *string `pulumi:"mountAccessor"`
	// `(string: <required>)` – Name of the MFA method.
	Name *string `pulumi:"name"`
	// Namespace ID computed by Vault.
	NamespaceId *string `pulumi:"namespaceId"`
	// Org Alias computed by Vault.
	OrgAlias *string `pulumi:"orgAlias"`
	// `(string: <required>)` - A base64-encoded third-party settings file retrieved
	// from PingID's configuration page.
	SettingsFileBase64 *string `pulumi:"settingsFileBase64"`
	// Type of configuration computed by Vault.
	Type *string `pulumi:"type"`
	// If set, enables use of PingID signature. Computed by Vault
	UseSignature *bool `pulumi:"useSignature"`
	// `(string)` - A format string for mapping Identity names to MFA method names.
	// Values to substitute should be placed in `{{}}`. For example, `"{{alias.name}}@example.com"`.
	// If blank, the Alias's Name field will be used as-is. Currently-supported mappings:
	// - alias.name: The name returned by the mount configured via the `mountAccessor` parameter
	// - entity.name: The name configured for the Entity
	// - alias.metadata.`<key>`: The value of the Alias's metadata parameter
	// - entity.metadata.`<key>`: The value of the Entity's metadata parameter
	UsernameFormat *string `pulumi:"usernameFormat"`
}

type MfaPingidState struct {
	// Admin URL computed by Vault.
	AdminUrl pulumi.StringPtrInput
	// Authenticator URL computed by Vault.
	AuthenticatorUrl pulumi.StringPtrInput
	// IDP URL computed by Vault.
	IdpUrl pulumi.StringPtrInput
	// `(string: <required>)` - The mount to tie this method to for use in automatic mappings.
	// The mapping will use the Name field of Aliases associated with this mount as the username in the mapping.
	MountAccessor pulumi.StringPtrInput
	// `(string: <required>)` – Name of the MFA method.
	Name pulumi.StringPtrInput
	// Namespace ID computed by Vault.
	NamespaceId pulumi.StringPtrInput
	// Org Alias computed by Vault.
	OrgAlias pulumi.StringPtrInput
	// `(string: <required>)` - A base64-encoded third-party settings file retrieved
	// from PingID's configuration page.
	SettingsFileBase64 pulumi.StringPtrInput
	// Type of configuration computed by Vault.
	Type pulumi.StringPtrInput
	// If set, enables use of PingID signature. Computed by Vault
	UseSignature pulumi.BoolPtrInput
	// `(string)` - A format string for mapping Identity names to MFA method names.
	// Values to substitute should be placed in `{{}}`. For example, `"{{alias.name}}@example.com"`.
	// If blank, the Alias's Name field will be used as-is. Currently-supported mappings:
	// - alias.name: The name returned by the mount configured via the `mountAccessor` parameter
	// - entity.name: The name configured for the Entity
	// - alias.metadata.`<key>`: The value of the Alias's metadata parameter
	// - entity.metadata.`<key>`: The value of the Entity's metadata parameter
	UsernameFormat pulumi.StringPtrInput
}

func (MfaPingidState) ElementType() reflect.Type {
	return reflect.TypeOf((*mfaPingidState)(nil)).Elem()
}

type mfaPingidArgs struct {
	// `(string: <required>)` - The mount to tie this method to for use in automatic mappings.
	// The mapping will use the Name field of Aliases associated with this mount as the username in the mapping.
	MountAccessor string `pulumi:"mountAccessor"`
	// `(string: <required>)` – Name of the MFA method.
	Name *string `pulumi:"name"`
	// `(string: <required>)` - A base64-encoded third-party settings file retrieved
	// from PingID's configuration page.
	SettingsFileBase64 string `pulumi:"settingsFileBase64"`
	// `(string)` - A format string for mapping Identity names to MFA method names.
	// Values to substitute should be placed in `{{}}`. For example, `"{{alias.name}}@example.com"`.
	// If blank, the Alias's Name field will be used as-is. Currently-supported mappings:
	// - alias.name: The name returned by the mount configured via the `mountAccessor` parameter
	// - entity.name: The name configured for the Entity
	// - alias.metadata.`<key>`: The value of the Alias's metadata parameter
	// - entity.metadata.`<key>`: The value of the Entity's metadata parameter
	UsernameFormat *string `pulumi:"usernameFormat"`
}

// The set of arguments for constructing a MfaPingid resource.
type MfaPingidArgs struct {
	// `(string: <required>)` - The mount to tie this method to for use in automatic mappings.
	// The mapping will use the Name field of Aliases associated with this mount as the username in the mapping.
	MountAccessor pulumi.StringInput
	// `(string: <required>)` – Name of the MFA method.
	Name pulumi.StringPtrInput
	// `(string: <required>)` - A base64-encoded third-party settings file retrieved
	// from PingID's configuration page.
	SettingsFileBase64 pulumi.StringInput
	// `(string)` - A format string for mapping Identity names to MFA method names.
	// Values to substitute should be placed in `{{}}`. For example, `"{{alias.name}}@example.com"`.
	// If blank, the Alias's Name field will be used as-is. Currently-supported mappings:
	// - alias.name: The name returned by the mount configured via the `mountAccessor` parameter
	// - entity.name: The name configured for the Entity
	// - alias.metadata.`<key>`: The value of the Alias's metadata parameter
	// - entity.metadata.`<key>`: The value of the Entity's metadata parameter
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
//          MfaPingidArray{ MfaPingidArgs{...} }
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
//          MfaPingidMap{ "key": MfaPingidArgs{...} }
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
