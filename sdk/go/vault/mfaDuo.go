// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vault

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-vault/sdk/v7/go/vault/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a resource to manage [Duo MFA](https://www.vaultproject.io/docs/enterprise/mfa/mfa-duo.html).
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
//	"github.com/pulumi/pulumi-vault/sdk/v7/go/vault"
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
//			_, err = vault.NewMfaDuo(ctx, "my_duo", &vault.MfaDuoArgs{
//				Name:           pulumi.String("my_duo"),
//				MountAccessor:  userpass.Accessor,
//				SecretKey:      pulumi.String("8C7THtrIigh2rPZQMbguugt8IUftWhMRCOBzbuyz"),
//				IntegrationKey: pulumi.String("BIACEUEAXI20BNWTEYXT"),
//				ApiHostname:    pulumi.String("api-2b5c39f5.duosecurity.com"),
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
// $ pulumi import vault:index/mfaDuo:MfaDuo my_duo my_duo
// ```
type MfaDuo struct {
	pulumi.CustomResourceState

	// `(string: <required>)` - API hostname for Duo.
	ApiHostname pulumi.StringOutput `pulumi:"apiHostname"`
	// `(string: <required>)` - Integration key for Duo.
	IntegrationKey pulumi.StringOutput `pulumi:"integrationKey"`
	// `(string: <required>)` - The mount to tie this method to for use in automatic mappings. The mapping will use the Name field of Aliases associated with this mount as the username in the mapping.
	MountAccessor pulumi.StringOutput `pulumi:"mountAccessor"`
	// `(string: <required>)` – Name of the MFA method.
	Name pulumi.StringOutput `pulumi:"name"`
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
	// *Available only for Vault Enterprise*.
	Namespace pulumi.StringPtrOutput `pulumi:"namespace"`
	// `(string)` - Push information for Duo.
	PushInfo pulumi.StringPtrOutput `pulumi:"pushInfo"`
	// `(string: <required>)` - Secret key for Duo.
	SecretKey pulumi.StringOutput `pulumi:"secretKey"`
	// `(string)` - A format string for mapping Identity names to MFA method names. Values to substitute should be placed in `{{}}`. For example, `"{{alias.name}}@example.com"`. If blank, the Alias's Name field will be used as-is. Currently-supported mappings:
	// - alias.name: The name returned by the mount configured via the `mountAccessor` parameter
	// - entity.name: The name configured for the Entity
	// - alias.metadata.`<key>`: The value of the Alias's metadata parameter
	// - entity.metadata.`<key>`: The value of the Entity's metadata parameter
	UsernameFormat pulumi.StringPtrOutput `pulumi:"usernameFormat"`
}

// NewMfaDuo registers a new resource with the given unique name, arguments, and options.
func NewMfaDuo(ctx *pulumi.Context,
	name string, args *MfaDuoArgs, opts ...pulumi.ResourceOption) (*MfaDuo, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ApiHostname == nil {
		return nil, errors.New("invalid value for required argument 'ApiHostname'")
	}
	if args.IntegrationKey == nil {
		return nil, errors.New("invalid value for required argument 'IntegrationKey'")
	}
	if args.MountAccessor == nil {
		return nil, errors.New("invalid value for required argument 'MountAccessor'")
	}
	if args.SecretKey == nil {
		return nil, errors.New("invalid value for required argument 'SecretKey'")
	}
	if args.IntegrationKey != nil {
		args.IntegrationKey = pulumi.ToSecret(args.IntegrationKey).(pulumi.StringInput)
	}
	if args.SecretKey != nil {
		args.SecretKey = pulumi.ToSecret(args.SecretKey).(pulumi.StringInput)
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"integrationKey",
		"secretKey",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource MfaDuo
	err := ctx.RegisterResource("vault:index/mfaDuo:MfaDuo", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetMfaDuo gets an existing MfaDuo resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetMfaDuo(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MfaDuoState, opts ...pulumi.ResourceOption) (*MfaDuo, error) {
	var resource MfaDuo
	err := ctx.ReadResource("vault:index/mfaDuo:MfaDuo", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering MfaDuo resources.
type mfaDuoState struct {
	// `(string: <required>)` - API hostname for Duo.
	ApiHostname *string `pulumi:"apiHostname"`
	// `(string: <required>)` - Integration key for Duo.
	IntegrationKey *string `pulumi:"integrationKey"`
	// `(string: <required>)` - The mount to tie this method to for use in automatic mappings. The mapping will use the Name field of Aliases associated with this mount as the username in the mapping.
	MountAccessor *string `pulumi:"mountAccessor"`
	// `(string: <required>)` – Name of the MFA method.
	Name *string `pulumi:"name"`
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
	// *Available only for Vault Enterprise*.
	Namespace *string `pulumi:"namespace"`
	// `(string)` - Push information for Duo.
	PushInfo *string `pulumi:"pushInfo"`
	// `(string: <required>)` - Secret key for Duo.
	SecretKey *string `pulumi:"secretKey"`
	// `(string)` - A format string for mapping Identity names to MFA method names. Values to substitute should be placed in `{{}}`. For example, `"{{alias.name}}@example.com"`. If blank, the Alias's Name field will be used as-is. Currently-supported mappings:
	// - alias.name: The name returned by the mount configured via the `mountAccessor` parameter
	// - entity.name: The name configured for the Entity
	// - alias.metadata.`<key>`: The value of the Alias's metadata parameter
	// - entity.metadata.`<key>`: The value of the Entity's metadata parameter
	UsernameFormat *string `pulumi:"usernameFormat"`
}

type MfaDuoState struct {
	// `(string: <required>)` - API hostname for Duo.
	ApiHostname pulumi.StringPtrInput
	// `(string: <required>)` - Integration key for Duo.
	IntegrationKey pulumi.StringPtrInput
	// `(string: <required>)` - The mount to tie this method to for use in automatic mappings. The mapping will use the Name field of Aliases associated with this mount as the username in the mapping.
	MountAccessor pulumi.StringPtrInput
	// `(string: <required>)` – Name of the MFA method.
	Name pulumi.StringPtrInput
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
	// *Available only for Vault Enterprise*.
	Namespace pulumi.StringPtrInput
	// `(string)` - Push information for Duo.
	PushInfo pulumi.StringPtrInput
	// `(string: <required>)` - Secret key for Duo.
	SecretKey pulumi.StringPtrInput
	// `(string)` - A format string for mapping Identity names to MFA method names. Values to substitute should be placed in `{{}}`. For example, `"{{alias.name}}@example.com"`. If blank, the Alias's Name field will be used as-is. Currently-supported mappings:
	// - alias.name: The name returned by the mount configured via the `mountAccessor` parameter
	// - entity.name: The name configured for the Entity
	// - alias.metadata.`<key>`: The value of the Alias's metadata parameter
	// - entity.metadata.`<key>`: The value of the Entity's metadata parameter
	UsernameFormat pulumi.StringPtrInput
}

func (MfaDuoState) ElementType() reflect.Type {
	return reflect.TypeOf((*mfaDuoState)(nil)).Elem()
}

type mfaDuoArgs struct {
	// `(string: <required>)` - API hostname for Duo.
	ApiHostname string `pulumi:"apiHostname"`
	// `(string: <required>)` - Integration key for Duo.
	IntegrationKey string `pulumi:"integrationKey"`
	// `(string: <required>)` - The mount to tie this method to for use in automatic mappings. The mapping will use the Name field of Aliases associated with this mount as the username in the mapping.
	MountAccessor string `pulumi:"mountAccessor"`
	// `(string: <required>)` – Name of the MFA method.
	Name *string `pulumi:"name"`
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
	// *Available only for Vault Enterprise*.
	Namespace *string `pulumi:"namespace"`
	// `(string)` - Push information for Duo.
	PushInfo *string `pulumi:"pushInfo"`
	// `(string: <required>)` - Secret key for Duo.
	SecretKey string `pulumi:"secretKey"`
	// `(string)` - A format string for mapping Identity names to MFA method names. Values to substitute should be placed in `{{}}`. For example, `"{{alias.name}}@example.com"`. If blank, the Alias's Name field will be used as-is. Currently-supported mappings:
	// - alias.name: The name returned by the mount configured via the `mountAccessor` parameter
	// - entity.name: The name configured for the Entity
	// - alias.metadata.`<key>`: The value of the Alias's metadata parameter
	// - entity.metadata.`<key>`: The value of the Entity's metadata parameter
	UsernameFormat *string `pulumi:"usernameFormat"`
}

// The set of arguments for constructing a MfaDuo resource.
type MfaDuoArgs struct {
	// `(string: <required>)` - API hostname for Duo.
	ApiHostname pulumi.StringInput
	// `(string: <required>)` - Integration key for Duo.
	IntegrationKey pulumi.StringInput
	// `(string: <required>)` - The mount to tie this method to for use in automatic mappings. The mapping will use the Name field of Aliases associated with this mount as the username in the mapping.
	MountAccessor pulumi.StringInput
	// `(string: <required>)` – Name of the MFA method.
	Name pulumi.StringPtrInput
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
	// *Available only for Vault Enterprise*.
	Namespace pulumi.StringPtrInput
	// `(string)` - Push information for Duo.
	PushInfo pulumi.StringPtrInput
	// `(string: <required>)` - Secret key for Duo.
	SecretKey pulumi.StringInput
	// `(string)` - A format string for mapping Identity names to MFA method names. Values to substitute should be placed in `{{}}`. For example, `"{{alias.name}}@example.com"`. If blank, the Alias's Name field will be used as-is. Currently-supported mappings:
	// - alias.name: The name returned by the mount configured via the `mountAccessor` parameter
	// - entity.name: The name configured for the Entity
	// - alias.metadata.`<key>`: The value of the Alias's metadata parameter
	// - entity.metadata.`<key>`: The value of the Entity's metadata parameter
	UsernameFormat pulumi.StringPtrInput
}

func (MfaDuoArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*mfaDuoArgs)(nil)).Elem()
}

type MfaDuoInput interface {
	pulumi.Input

	ToMfaDuoOutput() MfaDuoOutput
	ToMfaDuoOutputWithContext(ctx context.Context) MfaDuoOutput
}

func (*MfaDuo) ElementType() reflect.Type {
	return reflect.TypeOf((**MfaDuo)(nil)).Elem()
}

func (i *MfaDuo) ToMfaDuoOutput() MfaDuoOutput {
	return i.ToMfaDuoOutputWithContext(context.Background())
}

func (i *MfaDuo) ToMfaDuoOutputWithContext(ctx context.Context) MfaDuoOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MfaDuoOutput)
}

// MfaDuoArrayInput is an input type that accepts MfaDuoArray and MfaDuoArrayOutput values.
// You can construct a concrete instance of `MfaDuoArrayInput` via:
//
//	MfaDuoArray{ MfaDuoArgs{...} }
type MfaDuoArrayInput interface {
	pulumi.Input

	ToMfaDuoArrayOutput() MfaDuoArrayOutput
	ToMfaDuoArrayOutputWithContext(context.Context) MfaDuoArrayOutput
}

type MfaDuoArray []MfaDuoInput

func (MfaDuoArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*MfaDuo)(nil)).Elem()
}

func (i MfaDuoArray) ToMfaDuoArrayOutput() MfaDuoArrayOutput {
	return i.ToMfaDuoArrayOutputWithContext(context.Background())
}

func (i MfaDuoArray) ToMfaDuoArrayOutputWithContext(ctx context.Context) MfaDuoArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MfaDuoArrayOutput)
}

// MfaDuoMapInput is an input type that accepts MfaDuoMap and MfaDuoMapOutput values.
// You can construct a concrete instance of `MfaDuoMapInput` via:
//
//	MfaDuoMap{ "key": MfaDuoArgs{...} }
type MfaDuoMapInput interface {
	pulumi.Input

	ToMfaDuoMapOutput() MfaDuoMapOutput
	ToMfaDuoMapOutputWithContext(context.Context) MfaDuoMapOutput
}

type MfaDuoMap map[string]MfaDuoInput

func (MfaDuoMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*MfaDuo)(nil)).Elem()
}

func (i MfaDuoMap) ToMfaDuoMapOutput() MfaDuoMapOutput {
	return i.ToMfaDuoMapOutputWithContext(context.Background())
}

func (i MfaDuoMap) ToMfaDuoMapOutputWithContext(ctx context.Context) MfaDuoMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MfaDuoMapOutput)
}

type MfaDuoOutput struct{ *pulumi.OutputState }

func (MfaDuoOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MfaDuo)(nil)).Elem()
}

func (o MfaDuoOutput) ToMfaDuoOutput() MfaDuoOutput {
	return o
}

func (o MfaDuoOutput) ToMfaDuoOutputWithContext(ctx context.Context) MfaDuoOutput {
	return o
}

// `(string: <required>)` - API hostname for Duo.
func (o MfaDuoOutput) ApiHostname() pulumi.StringOutput {
	return o.ApplyT(func(v *MfaDuo) pulumi.StringOutput { return v.ApiHostname }).(pulumi.StringOutput)
}

// `(string: <required>)` - Integration key for Duo.
func (o MfaDuoOutput) IntegrationKey() pulumi.StringOutput {
	return o.ApplyT(func(v *MfaDuo) pulumi.StringOutput { return v.IntegrationKey }).(pulumi.StringOutput)
}

// `(string: <required>)` - The mount to tie this method to for use in automatic mappings. The mapping will use the Name field of Aliases associated with this mount as the username in the mapping.
func (o MfaDuoOutput) MountAccessor() pulumi.StringOutput {
	return o.ApplyT(func(v *MfaDuo) pulumi.StringOutput { return v.MountAccessor }).(pulumi.StringOutput)
}

// `(string: <required>)` – Name of the MFA method.
func (o MfaDuoOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *MfaDuo) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The namespace to provision the resource in.
// The value should not contain leading or trailing forward slashes.
// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
// *Available only for Vault Enterprise*.
func (o MfaDuoOutput) Namespace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MfaDuo) pulumi.StringPtrOutput { return v.Namespace }).(pulumi.StringPtrOutput)
}

// `(string)` - Push information for Duo.
func (o MfaDuoOutput) PushInfo() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MfaDuo) pulumi.StringPtrOutput { return v.PushInfo }).(pulumi.StringPtrOutput)
}

// `(string: <required>)` - Secret key for Duo.
func (o MfaDuoOutput) SecretKey() pulumi.StringOutput {
	return o.ApplyT(func(v *MfaDuo) pulumi.StringOutput { return v.SecretKey }).(pulumi.StringOutput)
}

// `(string)` - A format string for mapping Identity names to MFA method names. Values to substitute should be placed in `{{}}`. For example, `"{{alias.name}}@example.com"`. If blank, the Alias's Name field will be used as-is. Currently-supported mappings:
// - alias.name: The name returned by the mount configured via the `mountAccessor` parameter
// - entity.name: The name configured for the Entity
// - alias.metadata.`<key>`: The value of the Alias's metadata parameter
// - entity.metadata.`<key>`: The value of the Entity's metadata parameter
func (o MfaDuoOutput) UsernameFormat() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MfaDuo) pulumi.StringPtrOutput { return v.UsernameFormat }).(pulumi.StringPtrOutput)
}

type MfaDuoArrayOutput struct{ *pulumi.OutputState }

func (MfaDuoArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*MfaDuo)(nil)).Elem()
}

func (o MfaDuoArrayOutput) ToMfaDuoArrayOutput() MfaDuoArrayOutput {
	return o
}

func (o MfaDuoArrayOutput) ToMfaDuoArrayOutputWithContext(ctx context.Context) MfaDuoArrayOutput {
	return o
}

func (o MfaDuoArrayOutput) Index(i pulumi.IntInput) MfaDuoOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *MfaDuo {
		return vs[0].([]*MfaDuo)[vs[1].(int)]
	}).(MfaDuoOutput)
}

type MfaDuoMapOutput struct{ *pulumi.OutputState }

func (MfaDuoMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*MfaDuo)(nil)).Elem()
}

func (o MfaDuoMapOutput) ToMfaDuoMapOutput() MfaDuoMapOutput {
	return o
}

func (o MfaDuoMapOutput) ToMfaDuoMapOutputWithContext(ctx context.Context) MfaDuoMapOutput {
	return o
}

func (o MfaDuoMapOutput) MapIndex(k pulumi.StringInput) MfaDuoOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *MfaDuo {
		return vs[0].(map[string]*MfaDuo)[vs[1].(string)]
	}).(MfaDuoOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*MfaDuoInput)(nil)).Elem(), &MfaDuo{})
	pulumi.RegisterInputType(reflect.TypeOf((*MfaDuoArrayInput)(nil)).Elem(), MfaDuoArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*MfaDuoMapInput)(nil)).Elem(), MfaDuoMap{})
	pulumi.RegisterOutputType(MfaDuoOutput{})
	pulumi.RegisterOutputType(MfaDuoArrayOutput{})
	pulumi.RegisterOutputType(MfaDuoMapOutput{})
}
