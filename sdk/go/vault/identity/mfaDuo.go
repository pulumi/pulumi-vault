// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package identity

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-vault/sdk/v6/go/vault/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource for configuring the duo MFA method.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-vault/sdk/v6/go/vault/identity"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := identity.NewMfaDuo(ctx, "example", &identity.MfaDuoArgs{
//				ApiHostname:    pulumi.String("api-xxxxxxxx.duosecurity.com"),
//				SecretKey:      pulumi.String("secret-key"),
//				IntegrationKey: pulumi.String("secret-int-key"),
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
// $ pulumi import vault:identity/mfaDuo:MfaDuo example 0d89c36a-4ff5-4d70-8749-bb6a5598aeec
// ```
type MfaDuo struct {
	pulumi.CustomResourceState

	// API hostname for Duo
	ApiHostname pulumi.StringOutput `pulumi:"apiHostname"`
	// Integration key for Duo
	IntegrationKey pulumi.StringOutput `pulumi:"integrationKey"`
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
	// Push information for Duo.
	PushInfo pulumi.StringPtrOutput `pulumi:"pushInfo"`
	// Secret key for Duo
	SecretKey pulumi.StringOutput `pulumi:"secretKey"`
	// MFA type.
	Type pulumi.StringOutput `pulumi:"type"`
	// Require passcode upon MFA validation.
	UsePasscode pulumi.BoolPtrOutput `pulumi:"usePasscode"`
	// A template string for mapping Identity names to MFA methods.
	UsernameFormat pulumi.StringPtrOutput `pulumi:"usernameFormat"`
	// Resource UUID.
	Uuid pulumi.StringOutput `pulumi:"uuid"`
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
	err := ctx.RegisterResource("vault:identity/mfaDuo:MfaDuo", name, args, &resource, opts...)
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
	err := ctx.ReadResource("vault:identity/mfaDuo:MfaDuo", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering MfaDuo resources.
type mfaDuoState struct {
	// API hostname for Duo
	ApiHostname *string `pulumi:"apiHostname"`
	// Integration key for Duo
	IntegrationKey *string `pulumi:"integrationKey"`
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
	// Push information for Duo.
	PushInfo *string `pulumi:"pushInfo"`
	// Secret key for Duo
	SecretKey *string `pulumi:"secretKey"`
	// MFA type.
	Type *string `pulumi:"type"`
	// Require passcode upon MFA validation.
	UsePasscode *bool `pulumi:"usePasscode"`
	// A template string for mapping Identity names to MFA methods.
	UsernameFormat *string `pulumi:"usernameFormat"`
	// Resource UUID.
	Uuid *string `pulumi:"uuid"`
}

type MfaDuoState struct {
	// API hostname for Duo
	ApiHostname pulumi.StringPtrInput
	// Integration key for Duo
	IntegrationKey pulumi.StringPtrInput
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
	// Push information for Duo.
	PushInfo pulumi.StringPtrInput
	// Secret key for Duo
	SecretKey pulumi.StringPtrInput
	// MFA type.
	Type pulumi.StringPtrInput
	// Require passcode upon MFA validation.
	UsePasscode pulumi.BoolPtrInput
	// A template string for mapping Identity names to MFA methods.
	UsernameFormat pulumi.StringPtrInput
	// Resource UUID.
	Uuid pulumi.StringPtrInput
}

func (MfaDuoState) ElementType() reflect.Type {
	return reflect.TypeOf((*mfaDuoState)(nil)).Elem()
}

type mfaDuoArgs struct {
	// API hostname for Duo
	ApiHostname string `pulumi:"apiHostname"`
	// Integration key for Duo
	IntegrationKey string `pulumi:"integrationKey"`
	// Target namespace. (requires Enterprise)
	Namespace *string `pulumi:"namespace"`
	// Push information for Duo.
	PushInfo *string `pulumi:"pushInfo"`
	// Secret key for Duo
	SecretKey string `pulumi:"secretKey"`
	// Require passcode upon MFA validation.
	UsePasscode *bool `pulumi:"usePasscode"`
	// A template string for mapping Identity names to MFA methods.
	UsernameFormat *string `pulumi:"usernameFormat"`
}

// The set of arguments for constructing a MfaDuo resource.
type MfaDuoArgs struct {
	// API hostname for Duo
	ApiHostname pulumi.StringInput
	// Integration key for Duo
	IntegrationKey pulumi.StringInput
	// Target namespace. (requires Enterprise)
	Namespace pulumi.StringPtrInput
	// Push information for Duo.
	PushInfo pulumi.StringPtrInput
	// Secret key for Duo
	SecretKey pulumi.StringInput
	// Require passcode upon MFA validation.
	UsePasscode pulumi.BoolPtrInput
	// A template string for mapping Identity names to MFA methods.
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

// API hostname for Duo
func (o MfaDuoOutput) ApiHostname() pulumi.StringOutput {
	return o.ApplyT(func(v *MfaDuo) pulumi.StringOutput { return v.ApiHostname }).(pulumi.StringOutput)
}

// Integration key for Duo
func (o MfaDuoOutput) IntegrationKey() pulumi.StringOutput {
	return o.ApplyT(func(v *MfaDuo) pulumi.StringOutput { return v.IntegrationKey }).(pulumi.StringOutput)
}

// Method ID.
func (o MfaDuoOutput) MethodId() pulumi.StringOutput {
	return o.ApplyT(func(v *MfaDuo) pulumi.StringOutput { return v.MethodId }).(pulumi.StringOutput)
}

// Mount accessor.
func (o MfaDuoOutput) MountAccessor() pulumi.StringOutput {
	return o.ApplyT(func(v *MfaDuo) pulumi.StringOutput { return v.MountAccessor }).(pulumi.StringOutput)
}

// Method name.
func (o MfaDuoOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *MfaDuo) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Target namespace. (requires Enterprise)
func (o MfaDuoOutput) Namespace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MfaDuo) pulumi.StringPtrOutput { return v.Namespace }).(pulumi.StringPtrOutput)
}

// Method's namespace ID.
func (o MfaDuoOutput) NamespaceId() pulumi.StringOutput {
	return o.ApplyT(func(v *MfaDuo) pulumi.StringOutput { return v.NamespaceId }).(pulumi.StringOutput)
}

// Method's namespace path.
func (o MfaDuoOutput) NamespacePath() pulumi.StringOutput {
	return o.ApplyT(func(v *MfaDuo) pulumi.StringOutput { return v.NamespacePath }).(pulumi.StringOutput)
}

// Push information for Duo.
func (o MfaDuoOutput) PushInfo() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MfaDuo) pulumi.StringPtrOutput { return v.PushInfo }).(pulumi.StringPtrOutput)
}

// Secret key for Duo
func (o MfaDuoOutput) SecretKey() pulumi.StringOutput {
	return o.ApplyT(func(v *MfaDuo) pulumi.StringOutput { return v.SecretKey }).(pulumi.StringOutput)
}

// MFA type.
func (o MfaDuoOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *MfaDuo) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// Require passcode upon MFA validation.
func (o MfaDuoOutput) UsePasscode() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *MfaDuo) pulumi.BoolPtrOutput { return v.UsePasscode }).(pulumi.BoolPtrOutput)
}

// A template string for mapping Identity names to MFA methods.
func (o MfaDuoOutput) UsernameFormat() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MfaDuo) pulumi.StringPtrOutput { return v.UsernameFormat }).(pulumi.StringPtrOutput)
}

// Resource UUID.
func (o MfaDuoOutput) Uuid() pulumi.StringOutput {
	return o.ApplyT(func(v *MfaDuo) pulumi.StringOutput { return v.Uuid }).(pulumi.StringOutput)
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
