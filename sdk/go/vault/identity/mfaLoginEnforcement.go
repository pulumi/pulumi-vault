// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package identity

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource for configuring MFA login-enforcement
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-vault/sdk/v5/go/vault/identity"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		exampleMfaDuo, err := identity.NewMfaDuo(ctx, "exampleMfaDuo", &identity.MfaDuoArgs{
// 			SecretKey:      pulumi.String("secret-key"),
// 			IntegrationKey: pulumi.String("int-key"),
// 			ApiHostname:    pulumi.String("foo.baz"),
// 			PushInfo:       pulumi.String("push-info"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = identity.NewMfaLoginEnforcement(ctx, "exampleMfaLoginEnforcement", &identity.MfaLoginEnforcementArgs{
// 			MfaMethodIds: pulumi.StringArray{
// 				exampleMfaDuo.MethodId,
// 			},
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
// Resource can be imported using its `name` field, e.g.
//
// ```sh
//  $ pulumi import vault:identity/mfaLoginEnforcement:MfaLoginEnforcement example default
// ```
type MfaLoginEnforcement struct {
	pulumi.CustomResourceState

	// Set of auth method accessor IDs.
	AuthMethodAccessors pulumi.StringArrayOutput `pulumi:"authMethodAccessors"`
	// Set of auth method types.
	AuthMethodTypes pulumi.StringArrayOutput `pulumi:"authMethodTypes"`
	// Set of identity entity IDs.
	IdentityEntityIds pulumi.StringArrayOutput `pulumi:"identityEntityIds"`
	// Set of identity group IDs.
	IdentityGroupIds pulumi.StringArrayOutput `pulumi:"identityGroupIds"`
	// Set of MFA method UUIDs.
	MfaMethodIds pulumi.StringArrayOutput `pulumi:"mfaMethodIds"`
	// Login enforcement name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Target namespace. (requires Enterprise)
	Namespace pulumi.StringPtrOutput `pulumi:"namespace"`
	// Method's namespace ID.
	NamespaceId pulumi.StringOutput `pulumi:"namespaceId"`
	// Method's namespace path.
	NamespacePath pulumi.StringOutput `pulumi:"namespacePath"`
	// Resource UUID.
	Uuid pulumi.StringOutput `pulumi:"uuid"`
}

// NewMfaLoginEnforcement registers a new resource with the given unique name, arguments, and options.
func NewMfaLoginEnforcement(ctx *pulumi.Context,
	name string, args *MfaLoginEnforcementArgs, opts ...pulumi.ResourceOption) (*MfaLoginEnforcement, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.MfaMethodIds == nil {
		return nil, errors.New("invalid value for required argument 'MfaMethodIds'")
	}
	var resource MfaLoginEnforcement
	err := ctx.RegisterResource("vault:identity/mfaLoginEnforcement:MfaLoginEnforcement", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetMfaLoginEnforcement gets an existing MfaLoginEnforcement resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetMfaLoginEnforcement(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MfaLoginEnforcementState, opts ...pulumi.ResourceOption) (*MfaLoginEnforcement, error) {
	var resource MfaLoginEnforcement
	err := ctx.ReadResource("vault:identity/mfaLoginEnforcement:MfaLoginEnforcement", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering MfaLoginEnforcement resources.
type mfaLoginEnforcementState struct {
	// Set of auth method accessor IDs.
	AuthMethodAccessors []string `pulumi:"authMethodAccessors"`
	// Set of auth method types.
	AuthMethodTypes []string `pulumi:"authMethodTypes"`
	// Set of identity entity IDs.
	IdentityEntityIds []string `pulumi:"identityEntityIds"`
	// Set of identity group IDs.
	IdentityGroupIds []string `pulumi:"identityGroupIds"`
	// Set of MFA method UUIDs.
	MfaMethodIds []string `pulumi:"mfaMethodIds"`
	// Login enforcement name.
	Name *string `pulumi:"name"`
	// Target namespace. (requires Enterprise)
	Namespace *string `pulumi:"namespace"`
	// Method's namespace ID.
	NamespaceId *string `pulumi:"namespaceId"`
	// Method's namespace path.
	NamespacePath *string `pulumi:"namespacePath"`
	// Resource UUID.
	Uuid *string `pulumi:"uuid"`
}

type MfaLoginEnforcementState struct {
	// Set of auth method accessor IDs.
	AuthMethodAccessors pulumi.StringArrayInput
	// Set of auth method types.
	AuthMethodTypes pulumi.StringArrayInput
	// Set of identity entity IDs.
	IdentityEntityIds pulumi.StringArrayInput
	// Set of identity group IDs.
	IdentityGroupIds pulumi.StringArrayInput
	// Set of MFA method UUIDs.
	MfaMethodIds pulumi.StringArrayInput
	// Login enforcement name.
	Name pulumi.StringPtrInput
	// Target namespace. (requires Enterprise)
	Namespace pulumi.StringPtrInput
	// Method's namespace ID.
	NamespaceId pulumi.StringPtrInput
	// Method's namespace path.
	NamespacePath pulumi.StringPtrInput
	// Resource UUID.
	Uuid pulumi.StringPtrInput
}

func (MfaLoginEnforcementState) ElementType() reflect.Type {
	return reflect.TypeOf((*mfaLoginEnforcementState)(nil)).Elem()
}

type mfaLoginEnforcementArgs struct {
	// Set of auth method accessor IDs.
	AuthMethodAccessors []string `pulumi:"authMethodAccessors"`
	// Set of auth method types.
	AuthMethodTypes []string `pulumi:"authMethodTypes"`
	// Set of identity entity IDs.
	IdentityEntityIds []string `pulumi:"identityEntityIds"`
	// Set of identity group IDs.
	IdentityGroupIds []string `pulumi:"identityGroupIds"`
	// Set of MFA method UUIDs.
	MfaMethodIds []string `pulumi:"mfaMethodIds"`
	// Login enforcement name.
	Name *string `pulumi:"name"`
	// Target namespace. (requires Enterprise)
	Namespace *string `pulumi:"namespace"`
}

// The set of arguments for constructing a MfaLoginEnforcement resource.
type MfaLoginEnforcementArgs struct {
	// Set of auth method accessor IDs.
	AuthMethodAccessors pulumi.StringArrayInput
	// Set of auth method types.
	AuthMethodTypes pulumi.StringArrayInput
	// Set of identity entity IDs.
	IdentityEntityIds pulumi.StringArrayInput
	// Set of identity group IDs.
	IdentityGroupIds pulumi.StringArrayInput
	// Set of MFA method UUIDs.
	MfaMethodIds pulumi.StringArrayInput
	// Login enforcement name.
	Name pulumi.StringPtrInput
	// Target namespace. (requires Enterprise)
	Namespace pulumi.StringPtrInput
}

func (MfaLoginEnforcementArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*mfaLoginEnforcementArgs)(nil)).Elem()
}

type MfaLoginEnforcementInput interface {
	pulumi.Input

	ToMfaLoginEnforcementOutput() MfaLoginEnforcementOutput
	ToMfaLoginEnforcementOutputWithContext(ctx context.Context) MfaLoginEnforcementOutput
}

func (*MfaLoginEnforcement) ElementType() reflect.Type {
	return reflect.TypeOf((**MfaLoginEnforcement)(nil)).Elem()
}

func (i *MfaLoginEnforcement) ToMfaLoginEnforcementOutput() MfaLoginEnforcementOutput {
	return i.ToMfaLoginEnforcementOutputWithContext(context.Background())
}

func (i *MfaLoginEnforcement) ToMfaLoginEnforcementOutputWithContext(ctx context.Context) MfaLoginEnforcementOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MfaLoginEnforcementOutput)
}

// MfaLoginEnforcementArrayInput is an input type that accepts MfaLoginEnforcementArray and MfaLoginEnforcementArrayOutput values.
// You can construct a concrete instance of `MfaLoginEnforcementArrayInput` via:
//
//          MfaLoginEnforcementArray{ MfaLoginEnforcementArgs{...} }
type MfaLoginEnforcementArrayInput interface {
	pulumi.Input

	ToMfaLoginEnforcementArrayOutput() MfaLoginEnforcementArrayOutput
	ToMfaLoginEnforcementArrayOutputWithContext(context.Context) MfaLoginEnforcementArrayOutput
}

type MfaLoginEnforcementArray []MfaLoginEnforcementInput

func (MfaLoginEnforcementArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*MfaLoginEnforcement)(nil)).Elem()
}

func (i MfaLoginEnforcementArray) ToMfaLoginEnforcementArrayOutput() MfaLoginEnforcementArrayOutput {
	return i.ToMfaLoginEnforcementArrayOutputWithContext(context.Background())
}

func (i MfaLoginEnforcementArray) ToMfaLoginEnforcementArrayOutputWithContext(ctx context.Context) MfaLoginEnforcementArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MfaLoginEnforcementArrayOutput)
}

// MfaLoginEnforcementMapInput is an input type that accepts MfaLoginEnforcementMap and MfaLoginEnforcementMapOutput values.
// You can construct a concrete instance of `MfaLoginEnforcementMapInput` via:
//
//          MfaLoginEnforcementMap{ "key": MfaLoginEnforcementArgs{...} }
type MfaLoginEnforcementMapInput interface {
	pulumi.Input

	ToMfaLoginEnforcementMapOutput() MfaLoginEnforcementMapOutput
	ToMfaLoginEnforcementMapOutputWithContext(context.Context) MfaLoginEnforcementMapOutput
}

type MfaLoginEnforcementMap map[string]MfaLoginEnforcementInput

func (MfaLoginEnforcementMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*MfaLoginEnforcement)(nil)).Elem()
}

func (i MfaLoginEnforcementMap) ToMfaLoginEnforcementMapOutput() MfaLoginEnforcementMapOutput {
	return i.ToMfaLoginEnforcementMapOutputWithContext(context.Background())
}

func (i MfaLoginEnforcementMap) ToMfaLoginEnforcementMapOutputWithContext(ctx context.Context) MfaLoginEnforcementMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MfaLoginEnforcementMapOutput)
}

type MfaLoginEnforcementOutput struct{ *pulumi.OutputState }

func (MfaLoginEnforcementOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MfaLoginEnforcement)(nil)).Elem()
}

func (o MfaLoginEnforcementOutput) ToMfaLoginEnforcementOutput() MfaLoginEnforcementOutput {
	return o
}

func (o MfaLoginEnforcementOutput) ToMfaLoginEnforcementOutputWithContext(ctx context.Context) MfaLoginEnforcementOutput {
	return o
}

// Set of auth method accessor IDs.
func (o MfaLoginEnforcementOutput) AuthMethodAccessors() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *MfaLoginEnforcement) pulumi.StringArrayOutput { return v.AuthMethodAccessors }).(pulumi.StringArrayOutput)
}

// Set of auth method types.
func (o MfaLoginEnforcementOutput) AuthMethodTypes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *MfaLoginEnforcement) pulumi.StringArrayOutput { return v.AuthMethodTypes }).(pulumi.StringArrayOutput)
}

// Set of identity entity IDs.
func (o MfaLoginEnforcementOutput) IdentityEntityIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *MfaLoginEnforcement) pulumi.StringArrayOutput { return v.IdentityEntityIds }).(pulumi.StringArrayOutput)
}

// Set of identity group IDs.
func (o MfaLoginEnforcementOutput) IdentityGroupIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *MfaLoginEnforcement) pulumi.StringArrayOutput { return v.IdentityGroupIds }).(pulumi.StringArrayOutput)
}

// Set of MFA method UUIDs.
func (o MfaLoginEnforcementOutput) MfaMethodIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *MfaLoginEnforcement) pulumi.StringArrayOutput { return v.MfaMethodIds }).(pulumi.StringArrayOutput)
}

// Login enforcement name.
func (o MfaLoginEnforcementOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *MfaLoginEnforcement) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Target namespace. (requires Enterprise)
func (o MfaLoginEnforcementOutput) Namespace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MfaLoginEnforcement) pulumi.StringPtrOutput { return v.Namespace }).(pulumi.StringPtrOutput)
}

// Method's namespace ID.
func (o MfaLoginEnforcementOutput) NamespaceId() pulumi.StringOutput {
	return o.ApplyT(func(v *MfaLoginEnforcement) pulumi.StringOutput { return v.NamespaceId }).(pulumi.StringOutput)
}

// Method's namespace path.
func (o MfaLoginEnforcementOutput) NamespacePath() pulumi.StringOutput {
	return o.ApplyT(func(v *MfaLoginEnforcement) pulumi.StringOutput { return v.NamespacePath }).(pulumi.StringOutput)
}

// Resource UUID.
func (o MfaLoginEnforcementOutput) Uuid() pulumi.StringOutput {
	return o.ApplyT(func(v *MfaLoginEnforcement) pulumi.StringOutput { return v.Uuid }).(pulumi.StringOutput)
}

type MfaLoginEnforcementArrayOutput struct{ *pulumi.OutputState }

func (MfaLoginEnforcementArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*MfaLoginEnforcement)(nil)).Elem()
}

func (o MfaLoginEnforcementArrayOutput) ToMfaLoginEnforcementArrayOutput() MfaLoginEnforcementArrayOutput {
	return o
}

func (o MfaLoginEnforcementArrayOutput) ToMfaLoginEnforcementArrayOutputWithContext(ctx context.Context) MfaLoginEnforcementArrayOutput {
	return o
}

func (o MfaLoginEnforcementArrayOutput) Index(i pulumi.IntInput) MfaLoginEnforcementOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *MfaLoginEnforcement {
		return vs[0].([]*MfaLoginEnforcement)[vs[1].(int)]
	}).(MfaLoginEnforcementOutput)
}

type MfaLoginEnforcementMapOutput struct{ *pulumi.OutputState }

func (MfaLoginEnforcementMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*MfaLoginEnforcement)(nil)).Elem()
}

func (o MfaLoginEnforcementMapOutput) ToMfaLoginEnforcementMapOutput() MfaLoginEnforcementMapOutput {
	return o
}

func (o MfaLoginEnforcementMapOutput) ToMfaLoginEnforcementMapOutputWithContext(ctx context.Context) MfaLoginEnforcementMapOutput {
	return o
}

func (o MfaLoginEnforcementMapOutput) MapIndex(k pulumi.StringInput) MfaLoginEnforcementOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *MfaLoginEnforcement {
		return vs[0].(map[string]*MfaLoginEnforcement)[vs[1].(string)]
	}).(MfaLoginEnforcementOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*MfaLoginEnforcementInput)(nil)).Elem(), &MfaLoginEnforcement{})
	pulumi.RegisterInputType(reflect.TypeOf((*MfaLoginEnforcementArrayInput)(nil)).Elem(), MfaLoginEnforcementArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*MfaLoginEnforcementMapInput)(nil)).Elem(), MfaLoginEnforcementMap{})
	pulumi.RegisterOutputType(MfaLoginEnforcementOutput{})
	pulumi.RegisterOutputType(MfaLoginEnforcementArrayOutput{})
	pulumi.RegisterOutputType(MfaLoginEnforcementMapOutput{})
}
