// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package identity

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-vault/sdk/v6/go/vault/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages OIDC Assignments in a Vault server. See the [Vault documentation](https://www.vaultproject.io/api-docs/secret/identity/oidc-provider#create-or-update-an-assignment)
// for more information.
//
// ## Example Usage
//
// <!--Start PulumiCodeChooser -->
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
//			internal, err := identity.NewGroup(ctx, "internal", &identity.GroupArgs{
//				Type: pulumi.String("internal"),
//				Policies: pulumi.StringArray{
//					pulumi.String("dev"),
//					pulumi.String("test"),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			test, err := identity.NewEntity(ctx, "test", &identity.EntityArgs{
//				Policies: pulumi.StringArray{
//					pulumi.String("test"),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			_, err = identity.NewOidcAssignment(ctx, "default", &identity.OidcAssignmentArgs{
//				EntityIds: pulumi.StringArray{
//					test.ID(),
//				},
//				GroupIds: pulumi.StringArray{
//					internal.ID(),
//				},
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
// OIDC Assignments can be imported using the `name`, e.g.
//
// ```sh
// $ pulumi import vault:identity/oidcAssignment:OidcAssignment default assignment
// ```
type OidcAssignment struct {
	pulumi.CustomResourceState

	// A set of Vault entity IDs.
	EntityIds pulumi.StringArrayOutput `pulumi:"entityIds"`
	// A set of Vault group IDs.
	GroupIds pulumi.StringArrayOutput `pulumi:"groupIds"`
	// The name of the assignment.
	Name pulumi.StringOutput `pulumi:"name"`
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
	// *Available only for Vault Enterprise*.
	Namespace pulumi.StringPtrOutput `pulumi:"namespace"`
}

// NewOidcAssignment registers a new resource with the given unique name, arguments, and options.
func NewOidcAssignment(ctx *pulumi.Context,
	name string, args *OidcAssignmentArgs, opts ...pulumi.ResourceOption) (*OidcAssignment, error) {
	if args == nil {
		args = &OidcAssignmentArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource OidcAssignment
	err := ctx.RegisterResource("vault:identity/oidcAssignment:OidcAssignment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetOidcAssignment gets an existing OidcAssignment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetOidcAssignment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *OidcAssignmentState, opts ...pulumi.ResourceOption) (*OidcAssignment, error) {
	var resource OidcAssignment
	err := ctx.ReadResource("vault:identity/oidcAssignment:OidcAssignment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering OidcAssignment resources.
type oidcAssignmentState struct {
	// A set of Vault entity IDs.
	EntityIds []string `pulumi:"entityIds"`
	// A set of Vault group IDs.
	GroupIds []string `pulumi:"groupIds"`
	// The name of the assignment.
	Name *string `pulumi:"name"`
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
	// *Available only for Vault Enterprise*.
	Namespace *string `pulumi:"namespace"`
}

type OidcAssignmentState struct {
	// A set of Vault entity IDs.
	EntityIds pulumi.StringArrayInput
	// A set of Vault group IDs.
	GroupIds pulumi.StringArrayInput
	// The name of the assignment.
	Name pulumi.StringPtrInput
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
	// *Available only for Vault Enterprise*.
	Namespace pulumi.StringPtrInput
}

func (OidcAssignmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*oidcAssignmentState)(nil)).Elem()
}

type oidcAssignmentArgs struct {
	// A set of Vault entity IDs.
	EntityIds []string `pulumi:"entityIds"`
	// A set of Vault group IDs.
	GroupIds []string `pulumi:"groupIds"`
	// The name of the assignment.
	Name *string `pulumi:"name"`
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
	// *Available only for Vault Enterprise*.
	Namespace *string `pulumi:"namespace"`
}

// The set of arguments for constructing a OidcAssignment resource.
type OidcAssignmentArgs struct {
	// A set of Vault entity IDs.
	EntityIds pulumi.StringArrayInput
	// A set of Vault group IDs.
	GroupIds pulumi.StringArrayInput
	// The name of the assignment.
	Name pulumi.StringPtrInput
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
	// *Available only for Vault Enterprise*.
	Namespace pulumi.StringPtrInput
}

func (OidcAssignmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*oidcAssignmentArgs)(nil)).Elem()
}

type OidcAssignmentInput interface {
	pulumi.Input

	ToOidcAssignmentOutput() OidcAssignmentOutput
	ToOidcAssignmentOutputWithContext(ctx context.Context) OidcAssignmentOutput
}

func (*OidcAssignment) ElementType() reflect.Type {
	return reflect.TypeOf((**OidcAssignment)(nil)).Elem()
}

func (i *OidcAssignment) ToOidcAssignmentOutput() OidcAssignmentOutput {
	return i.ToOidcAssignmentOutputWithContext(context.Background())
}

func (i *OidcAssignment) ToOidcAssignmentOutputWithContext(ctx context.Context) OidcAssignmentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OidcAssignmentOutput)
}

// OidcAssignmentArrayInput is an input type that accepts OidcAssignmentArray and OidcAssignmentArrayOutput values.
// You can construct a concrete instance of `OidcAssignmentArrayInput` via:
//
//	OidcAssignmentArray{ OidcAssignmentArgs{...} }
type OidcAssignmentArrayInput interface {
	pulumi.Input

	ToOidcAssignmentArrayOutput() OidcAssignmentArrayOutput
	ToOidcAssignmentArrayOutputWithContext(context.Context) OidcAssignmentArrayOutput
}

type OidcAssignmentArray []OidcAssignmentInput

func (OidcAssignmentArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*OidcAssignment)(nil)).Elem()
}

func (i OidcAssignmentArray) ToOidcAssignmentArrayOutput() OidcAssignmentArrayOutput {
	return i.ToOidcAssignmentArrayOutputWithContext(context.Background())
}

func (i OidcAssignmentArray) ToOidcAssignmentArrayOutputWithContext(ctx context.Context) OidcAssignmentArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OidcAssignmentArrayOutput)
}

// OidcAssignmentMapInput is an input type that accepts OidcAssignmentMap and OidcAssignmentMapOutput values.
// You can construct a concrete instance of `OidcAssignmentMapInput` via:
//
//	OidcAssignmentMap{ "key": OidcAssignmentArgs{...} }
type OidcAssignmentMapInput interface {
	pulumi.Input

	ToOidcAssignmentMapOutput() OidcAssignmentMapOutput
	ToOidcAssignmentMapOutputWithContext(context.Context) OidcAssignmentMapOutput
}

type OidcAssignmentMap map[string]OidcAssignmentInput

func (OidcAssignmentMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*OidcAssignment)(nil)).Elem()
}

func (i OidcAssignmentMap) ToOidcAssignmentMapOutput() OidcAssignmentMapOutput {
	return i.ToOidcAssignmentMapOutputWithContext(context.Background())
}

func (i OidcAssignmentMap) ToOidcAssignmentMapOutputWithContext(ctx context.Context) OidcAssignmentMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OidcAssignmentMapOutput)
}

type OidcAssignmentOutput struct{ *pulumi.OutputState }

func (OidcAssignmentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**OidcAssignment)(nil)).Elem()
}

func (o OidcAssignmentOutput) ToOidcAssignmentOutput() OidcAssignmentOutput {
	return o
}

func (o OidcAssignmentOutput) ToOidcAssignmentOutputWithContext(ctx context.Context) OidcAssignmentOutput {
	return o
}

// A set of Vault entity IDs.
func (o OidcAssignmentOutput) EntityIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *OidcAssignment) pulumi.StringArrayOutput { return v.EntityIds }).(pulumi.StringArrayOutput)
}

// A set of Vault group IDs.
func (o OidcAssignmentOutput) GroupIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *OidcAssignment) pulumi.StringArrayOutput { return v.GroupIds }).(pulumi.StringArrayOutput)
}

// The name of the assignment.
func (o OidcAssignmentOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *OidcAssignment) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The namespace to provision the resource in.
// The value should not contain leading or trailing forward slashes.
// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
// *Available only for Vault Enterprise*.
func (o OidcAssignmentOutput) Namespace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OidcAssignment) pulumi.StringPtrOutput { return v.Namespace }).(pulumi.StringPtrOutput)
}

type OidcAssignmentArrayOutput struct{ *pulumi.OutputState }

func (OidcAssignmentArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*OidcAssignment)(nil)).Elem()
}

func (o OidcAssignmentArrayOutput) ToOidcAssignmentArrayOutput() OidcAssignmentArrayOutput {
	return o
}

func (o OidcAssignmentArrayOutput) ToOidcAssignmentArrayOutputWithContext(ctx context.Context) OidcAssignmentArrayOutput {
	return o
}

func (o OidcAssignmentArrayOutput) Index(i pulumi.IntInput) OidcAssignmentOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *OidcAssignment {
		return vs[0].([]*OidcAssignment)[vs[1].(int)]
	}).(OidcAssignmentOutput)
}

type OidcAssignmentMapOutput struct{ *pulumi.OutputState }

func (OidcAssignmentMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*OidcAssignment)(nil)).Elem()
}

func (o OidcAssignmentMapOutput) ToOidcAssignmentMapOutput() OidcAssignmentMapOutput {
	return o
}

func (o OidcAssignmentMapOutput) ToOidcAssignmentMapOutputWithContext(ctx context.Context) OidcAssignmentMapOutput {
	return o
}

func (o OidcAssignmentMapOutput) MapIndex(k pulumi.StringInput) OidcAssignmentOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *OidcAssignment {
		return vs[0].(map[string]*OidcAssignment)[vs[1].(string)]
	}).(OidcAssignmentOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*OidcAssignmentInput)(nil)).Elem(), &OidcAssignment{})
	pulumi.RegisterInputType(reflect.TypeOf((*OidcAssignmentArrayInput)(nil)).Elem(), OidcAssignmentArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*OidcAssignmentMapInput)(nil)).Elem(), OidcAssignmentMap{})
	pulumi.RegisterOutputType(OidcAssignmentOutput{})
	pulumi.RegisterOutputType(OidcAssignmentArrayOutput{})
	pulumi.RegisterOutputType(OidcAssignmentMapOutput{})
}
