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

// ## Example Usage
//
// <!--Start PulumiCodeChooser -->
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
//			key, err := identity.NewOidcKey(ctx, "key", &identity.OidcKeyArgs{
//				Algorithm: pulumi.String("RS256"),
//			})
//			if err != nil {
//				return err
//			}
//			roleOidcRole, err := identity.NewOidcRole(ctx, "roleOidcRole", &identity.OidcRoleArgs{
//				Key: key.Name,
//			})
//			if err != nil {
//				return err
//			}
//			_, err = identity.NewOidcKeyAllowedClientID(ctx, "roleOidcKeyAllowedClientID", &identity.OidcKeyAllowedClientIDArgs{
//				KeyName:         key.Name,
//				AllowedClientId: roleOidcRole.ClientId,
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
type OidcKeyAllowedClientID struct {
	pulumi.CustomResourceState

	// Client ID to allow usage with the OIDC named key
	AllowedClientId pulumi.StringOutput `pulumi:"allowedClientId"`
	// Name of the OIDC Key allow the Client ID.
	KeyName pulumi.StringOutput `pulumi:"keyName"`
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
	// *Available only for Vault Enterprise*.
	Namespace pulumi.StringPtrOutput `pulumi:"namespace"`
}

// NewOidcKeyAllowedClientID registers a new resource with the given unique name, arguments, and options.
func NewOidcKeyAllowedClientID(ctx *pulumi.Context,
	name string, args *OidcKeyAllowedClientIDArgs, opts ...pulumi.ResourceOption) (*OidcKeyAllowedClientID, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AllowedClientId == nil {
		return nil, errors.New("invalid value for required argument 'AllowedClientId'")
	}
	if args.KeyName == nil {
		return nil, errors.New("invalid value for required argument 'KeyName'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource OidcKeyAllowedClientID
	err := ctx.RegisterResource("vault:identity/oidcKeyAllowedClientID:OidcKeyAllowedClientID", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetOidcKeyAllowedClientID gets an existing OidcKeyAllowedClientID resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetOidcKeyAllowedClientID(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *OidcKeyAllowedClientIDState, opts ...pulumi.ResourceOption) (*OidcKeyAllowedClientID, error) {
	var resource OidcKeyAllowedClientID
	err := ctx.ReadResource("vault:identity/oidcKeyAllowedClientID:OidcKeyAllowedClientID", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering OidcKeyAllowedClientID resources.
type oidcKeyAllowedClientIDState struct {
	// Client ID to allow usage with the OIDC named key
	AllowedClientId *string `pulumi:"allowedClientId"`
	// Name of the OIDC Key allow the Client ID.
	KeyName *string `pulumi:"keyName"`
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
	// *Available only for Vault Enterprise*.
	Namespace *string `pulumi:"namespace"`
}

type OidcKeyAllowedClientIDState struct {
	// Client ID to allow usage with the OIDC named key
	AllowedClientId pulumi.StringPtrInput
	// Name of the OIDC Key allow the Client ID.
	KeyName pulumi.StringPtrInput
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
	// *Available only for Vault Enterprise*.
	Namespace pulumi.StringPtrInput
}

func (OidcKeyAllowedClientIDState) ElementType() reflect.Type {
	return reflect.TypeOf((*oidcKeyAllowedClientIDState)(nil)).Elem()
}

type oidcKeyAllowedClientIDArgs struct {
	// Client ID to allow usage with the OIDC named key
	AllowedClientId string `pulumi:"allowedClientId"`
	// Name of the OIDC Key allow the Client ID.
	KeyName string `pulumi:"keyName"`
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
	// *Available only for Vault Enterprise*.
	Namespace *string `pulumi:"namespace"`
}

// The set of arguments for constructing a OidcKeyAllowedClientID resource.
type OidcKeyAllowedClientIDArgs struct {
	// Client ID to allow usage with the OIDC named key
	AllowedClientId pulumi.StringInput
	// Name of the OIDC Key allow the Client ID.
	KeyName pulumi.StringInput
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
	// *Available only for Vault Enterprise*.
	Namespace pulumi.StringPtrInput
}

func (OidcKeyAllowedClientIDArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*oidcKeyAllowedClientIDArgs)(nil)).Elem()
}

type OidcKeyAllowedClientIDInput interface {
	pulumi.Input

	ToOidcKeyAllowedClientIDOutput() OidcKeyAllowedClientIDOutput
	ToOidcKeyAllowedClientIDOutputWithContext(ctx context.Context) OidcKeyAllowedClientIDOutput
}

func (*OidcKeyAllowedClientID) ElementType() reflect.Type {
	return reflect.TypeOf((**OidcKeyAllowedClientID)(nil)).Elem()
}

func (i *OidcKeyAllowedClientID) ToOidcKeyAllowedClientIDOutput() OidcKeyAllowedClientIDOutput {
	return i.ToOidcKeyAllowedClientIDOutputWithContext(context.Background())
}

func (i *OidcKeyAllowedClientID) ToOidcKeyAllowedClientIDOutputWithContext(ctx context.Context) OidcKeyAllowedClientIDOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OidcKeyAllowedClientIDOutput)
}

// OidcKeyAllowedClientIDArrayInput is an input type that accepts OidcKeyAllowedClientIDArray and OidcKeyAllowedClientIDArrayOutput values.
// You can construct a concrete instance of `OidcKeyAllowedClientIDArrayInput` via:
//
//	OidcKeyAllowedClientIDArray{ OidcKeyAllowedClientIDArgs{...} }
type OidcKeyAllowedClientIDArrayInput interface {
	pulumi.Input

	ToOidcKeyAllowedClientIDArrayOutput() OidcKeyAllowedClientIDArrayOutput
	ToOidcKeyAllowedClientIDArrayOutputWithContext(context.Context) OidcKeyAllowedClientIDArrayOutput
}

type OidcKeyAllowedClientIDArray []OidcKeyAllowedClientIDInput

func (OidcKeyAllowedClientIDArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*OidcKeyAllowedClientID)(nil)).Elem()
}

func (i OidcKeyAllowedClientIDArray) ToOidcKeyAllowedClientIDArrayOutput() OidcKeyAllowedClientIDArrayOutput {
	return i.ToOidcKeyAllowedClientIDArrayOutputWithContext(context.Background())
}

func (i OidcKeyAllowedClientIDArray) ToOidcKeyAllowedClientIDArrayOutputWithContext(ctx context.Context) OidcKeyAllowedClientIDArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OidcKeyAllowedClientIDArrayOutput)
}

// OidcKeyAllowedClientIDMapInput is an input type that accepts OidcKeyAllowedClientIDMap and OidcKeyAllowedClientIDMapOutput values.
// You can construct a concrete instance of `OidcKeyAllowedClientIDMapInput` via:
//
//	OidcKeyAllowedClientIDMap{ "key": OidcKeyAllowedClientIDArgs{...} }
type OidcKeyAllowedClientIDMapInput interface {
	pulumi.Input

	ToOidcKeyAllowedClientIDMapOutput() OidcKeyAllowedClientIDMapOutput
	ToOidcKeyAllowedClientIDMapOutputWithContext(context.Context) OidcKeyAllowedClientIDMapOutput
}

type OidcKeyAllowedClientIDMap map[string]OidcKeyAllowedClientIDInput

func (OidcKeyAllowedClientIDMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*OidcKeyAllowedClientID)(nil)).Elem()
}

func (i OidcKeyAllowedClientIDMap) ToOidcKeyAllowedClientIDMapOutput() OidcKeyAllowedClientIDMapOutput {
	return i.ToOidcKeyAllowedClientIDMapOutputWithContext(context.Background())
}

func (i OidcKeyAllowedClientIDMap) ToOidcKeyAllowedClientIDMapOutputWithContext(ctx context.Context) OidcKeyAllowedClientIDMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OidcKeyAllowedClientIDMapOutput)
}

type OidcKeyAllowedClientIDOutput struct{ *pulumi.OutputState }

func (OidcKeyAllowedClientIDOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**OidcKeyAllowedClientID)(nil)).Elem()
}

func (o OidcKeyAllowedClientIDOutput) ToOidcKeyAllowedClientIDOutput() OidcKeyAllowedClientIDOutput {
	return o
}

func (o OidcKeyAllowedClientIDOutput) ToOidcKeyAllowedClientIDOutputWithContext(ctx context.Context) OidcKeyAllowedClientIDOutput {
	return o
}

// Client ID to allow usage with the OIDC named key
func (o OidcKeyAllowedClientIDOutput) AllowedClientId() pulumi.StringOutput {
	return o.ApplyT(func(v *OidcKeyAllowedClientID) pulumi.StringOutput { return v.AllowedClientId }).(pulumi.StringOutput)
}

// Name of the OIDC Key allow the Client ID.
func (o OidcKeyAllowedClientIDOutput) KeyName() pulumi.StringOutput {
	return o.ApplyT(func(v *OidcKeyAllowedClientID) pulumi.StringOutput { return v.KeyName }).(pulumi.StringOutput)
}

// The namespace to provision the resource in.
// The value should not contain leading or trailing forward slashes.
// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
// *Available only for Vault Enterprise*.
func (o OidcKeyAllowedClientIDOutput) Namespace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OidcKeyAllowedClientID) pulumi.StringPtrOutput { return v.Namespace }).(pulumi.StringPtrOutput)
}

type OidcKeyAllowedClientIDArrayOutput struct{ *pulumi.OutputState }

func (OidcKeyAllowedClientIDArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*OidcKeyAllowedClientID)(nil)).Elem()
}

func (o OidcKeyAllowedClientIDArrayOutput) ToOidcKeyAllowedClientIDArrayOutput() OidcKeyAllowedClientIDArrayOutput {
	return o
}

func (o OidcKeyAllowedClientIDArrayOutput) ToOidcKeyAllowedClientIDArrayOutputWithContext(ctx context.Context) OidcKeyAllowedClientIDArrayOutput {
	return o
}

func (o OidcKeyAllowedClientIDArrayOutput) Index(i pulumi.IntInput) OidcKeyAllowedClientIDOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *OidcKeyAllowedClientID {
		return vs[0].([]*OidcKeyAllowedClientID)[vs[1].(int)]
	}).(OidcKeyAllowedClientIDOutput)
}

type OidcKeyAllowedClientIDMapOutput struct{ *pulumi.OutputState }

func (OidcKeyAllowedClientIDMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*OidcKeyAllowedClientID)(nil)).Elem()
}

func (o OidcKeyAllowedClientIDMapOutput) ToOidcKeyAllowedClientIDMapOutput() OidcKeyAllowedClientIDMapOutput {
	return o
}

func (o OidcKeyAllowedClientIDMapOutput) ToOidcKeyAllowedClientIDMapOutputWithContext(ctx context.Context) OidcKeyAllowedClientIDMapOutput {
	return o
}

func (o OidcKeyAllowedClientIDMapOutput) MapIndex(k pulumi.StringInput) OidcKeyAllowedClientIDOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *OidcKeyAllowedClientID {
		return vs[0].(map[string]*OidcKeyAllowedClientID)[vs[1].(string)]
	}).(OidcKeyAllowedClientIDOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*OidcKeyAllowedClientIDInput)(nil)).Elem(), &OidcKeyAllowedClientID{})
	pulumi.RegisterInputType(reflect.TypeOf((*OidcKeyAllowedClientIDArrayInput)(nil)).Elem(), OidcKeyAllowedClientIDArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*OidcKeyAllowedClientIDMapInput)(nil)).Elem(), OidcKeyAllowedClientIDMap{})
	pulumi.RegisterOutputType(OidcKeyAllowedClientIDOutput{})
	pulumi.RegisterOutputType(OidcKeyAllowedClientIDArrayOutput{})
	pulumi.RegisterOutputType(OidcKeyAllowedClientIDMapOutput{})
}
