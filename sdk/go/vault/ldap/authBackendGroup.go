// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ldap

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a resource to create a group in an [LDAP auth backend within Vault](https://www.vaultproject.io/docs/auth/ldap.html).
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-vault/sdk/v4/go/vault/ldap"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		ldap, err := ldap.NewAuthBackend(ctx, "ldap", &ldap.AuthBackendArgs{
// 			Path:        pulumi.String("ldap"),
// 			Url:         pulumi.String("ldaps://dc-01.example.org"),
// 			Userdn:      pulumi.String("OU=Users,OU=Accounts,DC=example,DC=org"),
// 			Userattr:    pulumi.String("sAMAccountName"),
// 			Upndomain:   pulumi.String("EXAMPLE.ORG"),
// 			Discoverdn:  pulumi.Bool(false),
// 			Groupdn:     pulumi.String("OU=Groups,DC=example,DC=org"),
// 			Groupfilter: pulumi.String("(&(objectClass=group)(member:1.2.840.113556.1.4.1941:={{.UserDN}}))"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = ldap.NewAuthBackendGroup(ctx, "group", &ldap.AuthBackendGroupArgs{
// 			Groupname: pulumi.String("dba"),
// 			Policies: pulumi.StringArray{
// 				pulumi.String("dba"),
// 			},
// 			Backend: ldap.Path,
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
// LDAP authentication backend groups can be imported using the `path`, e.g.
//
// ```sh
//  $ pulumi import vault:ldap/authBackendGroup:AuthBackendGroup foo auth/ldap/groups/foo
// ```
type AuthBackendGroup struct {
	pulumi.CustomResourceState

	// Path to the authentication backend
	Backend pulumi.StringPtrOutput `pulumi:"backend"`
	// The LDAP groupname
	Groupname pulumi.StringOutput `pulumi:"groupname"`
	// Policies which should be granted to members of the group
	Policies pulumi.StringArrayOutput `pulumi:"policies"`
}

// NewAuthBackendGroup registers a new resource with the given unique name, arguments, and options.
func NewAuthBackendGroup(ctx *pulumi.Context,
	name string, args *AuthBackendGroupArgs, opts ...pulumi.ResourceOption) (*AuthBackendGroup, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Groupname == nil {
		return nil, errors.New("invalid value for required argument 'Groupname'")
	}
	var resource AuthBackendGroup
	err := ctx.RegisterResource("vault:ldap/authBackendGroup:AuthBackendGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAuthBackendGroup gets an existing AuthBackendGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAuthBackendGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AuthBackendGroupState, opts ...pulumi.ResourceOption) (*AuthBackendGroup, error) {
	var resource AuthBackendGroup
	err := ctx.ReadResource("vault:ldap/authBackendGroup:AuthBackendGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AuthBackendGroup resources.
type authBackendGroupState struct {
	// Path to the authentication backend
	Backend *string `pulumi:"backend"`
	// The LDAP groupname
	Groupname *string `pulumi:"groupname"`
	// Policies which should be granted to members of the group
	Policies []string `pulumi:"policies"`
}

type AuthBackendGroupState struct {
	// Path to the authentication backend
	Backend pulumi.StringPtrInput
	// The LDAP groupname
	Groupname pulumi.StringPtrInput
	// Policies which should be granted to members of the group
	Policies pulumi.StringArrayInput
}

func (AuthBackendGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*authBackendGroupState)(nil)).Elem()
}

type authBackendGroupArgs struct {
	// Path to the authentication backend
	Backend *string `pulumi:"backend"`
	// The LDAP groupname
	Groupname string `pulumi:"groupname"`
	// Policies which should be granted to members of the group
	Policies []string `pulumi:"policies"`
}

// The set of arguments for constructing a AuthBackendGroup resource.
type AuthBackendGroupArgs struct {
	// Path to the authentication backend
	Backend pulumi.StringPtrInput
	// The LDAP groupname
	Groupname pulumi.StringInput
	// Policies which should be granted to members of the group
	Policies pulumi.StringArrayInput
}

func (AuthBackendGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*authBackendGroupArgs)(nil)).Elem()
}

type AuthBackendGroupInput interface {
	pulumi.Input

	ToAuthBackendGroupOutput() AuthBackendGroupOutput
	ToAuthBackendGroupOutputWithContext(ctx context.Context) AuthBackendGroupOutput
}

func (*AuthBackendGroup) ElementType() reflect.Type {
	return reflect.TypeOf((*AuthBackendGroup)(nil))
}

func (i *AuthBackendGroup) ToAuthBackendGroupOutput() AuthBackendGroupOutput {
	return i.ToAuthBackendGroupOutputWithContext(context.Background())
}

func (i *AuthBackendGroup) ToAuthBackendGroupOutputWithContext(ctx context.Context) AuthBackendGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AuthBackendGroupOutput)
}

func (i *AuthBackendGroup) ToAuthBackendGroupPtrOutput() AuthBackendGroupPtrOutput {
	return i.ToAuthBackendGroupPtrOutputWithContext(context.Background())
}

func (i *AuthBackendGroup) ToAuthBackendGroupPtrOutputWithContext(ctx context.Context) AuthBackendGroupPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AuthBackendGroupPtrOutput)
}

type AuthBackendGroupPtrInput interface {
	pulumi.Input

	ToAuthBackendGroupPtrOutput() AuthBackendGroupPtrOutput
	ToAuthBackendGroupPtrOutputWithContext(ctx context.Context) AuthBackendGroupPtrOutput
}

type authBackendGroupPtrType AuthBackendGroupArgs

func (*authBackendGroupPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AuthBackendGroup)(nil))
}

func (i *authBackendGroupPtrType) ToAuthBackendGroupPtrOutput() AuthBackendGroupPtrOutput {
	return i.ToAuthBackendGroupPtrOutputWithContext(context.Background())
}

func (i *authBackendGroupPtrType) ToAuthBackendGroupPtrOutputWithContext(ctx context.Context) AuthBackendGroupPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AuthBackendGroupPtrOutput)
}

// AuthBackendGroupArrayInput is an input type that accepts AuthBackendGroupArray and AuthBackendGroupArrayOutput values.
// You can construct a concrete instance of `AuthBackendGroupArrayInput` via:
//
//          AuthBackendGroupArray{ AuthBackendGroupArgs{...} }
type AuthBackendGroupArrayInput interface {
	pulumi.Input

	ToAuthBackendGroupArrayOutput() AuthBackendGroupArrayOutput
	ToAuthBackendGroupArrayOutputWithContext(context.Context) AuthBackendGroupArrayOutput
}

type AuthBackendGroupArray []AuthBackendGroupInput

func (AuthBackendGroupArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*AuthBackendGroup)(nil))
}

func (i AuthBackendGroupArray) ToAuthBackendGroupArrayOutput() AuthBackendGroupArrayOutput {
	return i.ToAuthBackendGroupArrayOutputWithContext(context.Background())
}

func (i AuthBackendGroupArray) ToAuthBackendGroupArrayOutputWithContext(ctx context.Context) AuthBackendGroupArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AuthBackendGroupArrayOutput)
}

// AuthBackendGroupMapInput is an input type that accepts AuthBackendGroupMap and AuthBackendGroupMapOutput values.
// You can construct a concrete instance of `AuthBackendGroupMapInput` via:
//
//          AuthBackendGroupMap{ "key": AuthBackendGroupArgs{...} }
type AuthBackendGroupMapInput interface {
	pulumi.Input

	ToAuthBackendGroupMapOutput() AuthBackendGroupMapOutput
	ToAuthBackendGroupMapOutputWithContext(context.Context) AuthBackendGroupMapOutput
}

type AuthBackendGroupMap map[string]AuthBackendGroupInput

func (AuthBackendGroupMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*AuthBackendGroup)(nil))
}

func (i AuthBackendGroupMap) ToAuthBackendGroupMapOutput() AuthBackendGroupMapOutput {
	return i.ToAuthBackendGroupMapOutputWithContext(context.Background())
}

func (i AuthBackendGroupMap) ToAuthBackendGroupMapOutputWithContext(ctx context.Context) AuthBackendGroupMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AuthBackendGroupMapOutput)
}

type AuthBackendGroupOutput struct {
	*pulumi.OutputState
}

func (AuthBackendGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AuthBackendGroup)(nil))
}

func (o AuthBackendGroupOutput) ToAuthBackendGroupOutput() AuthBackendGroupOutput {
	return o
}

func (o AuthBackendGroupOutput) ToAuthBackendGroupOutputWithContext(ctx context.Context) AuthBackendGroupOutput {
	return o
}

func (o AuthBackendGroupOutput) ToAuthBackendGroupPtrOutput() AuthBackendGroupPtrOutput {
	return o.ToAuthBackendGroupPtrOutputWithContext(context.Background())
}

func (o AuthBackendGroupOutput) ToAuthBackendGroupPtrOutputWithContext(ctx context.Context) AuthBackendGroupPtrOutput {
	return o.ApplyT(func(v AuthBackendGroup) *AuthBackendGroup {
		return &v
	}).(AuthBackendGroupPtrOutput)
}

type AuthBackendGroupPtrOutput struct {
	*pulumi.OutputState
}

func (AuthBackendGroupPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AuthBackendGroup)(nil))
}

func (o AuthBackendGroupPtrOutput) ToAuthBackendGroupPtrOutput() AuthBackendGroupPtrOutput {
	return o
}

func (o AuthBackendGroupPtrOutput) ToAuthBackendGroupPtrOutputWithContext(ctx context.Context) AuthBackendGroupPtrOutput {
	return o
}

type AuthBackendGroupArrayOutput struct{ *pulumi.OutputState }

func (AuthBackendGroupArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AuthBackendGroup)(nil))
}

func (o AuthBackendGroupArrayOutput) ToAuthBackendGroupArrayOutput() AuthBackendGroupArrayOutput {
	return o
}

func (o AuthBackendGroupArrayOutput) ToAuthBackendGroupArrayOutputWithContext(ctx context.Context) AuthBackendGroupArrayOutput {
	return o
}

func (o AuthBackendGroupArrayOutput) Index(i pulumi.IntInput) AuthBackendGroupOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) AuthBackendGroup {
		return vs[0].([]AuthBackendGroup)[vs[1].(int)]
	}).(AuthBackendGroupOutput)
}

type AuthBackendGroupMapOutput struct{ *pulumi.OutputState }

func (AuthBackendGroupMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]AuthBackendGroup)(nil))
}

func (o AuthBackendGroupMapOutput) ToAuthBackendGroupMapOutput() AuthBackendGroupMapOutput {
	return o
}

func (o AuthBackendGroupMapOutput) ToAuthBackendGroupMapOutputWithContext(ctx context.Context) AuthBackendGroupMapOutput {
	return o
}

func (o AuthBackendGroupMapOutput) MapIndex(k pulumi.StringInput) AuthBackendGroupOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) AuthBackendGroup {
		return vs[0].(map[string]AuthBackendGroup)[vs[1].(string)]
	}).(AuthBackendGroupOutput)
}

func init() {
	pulumi.RegisterOutputType(AuthBackendGroupOutput{})
	pulumi.RegisterOutputType(AuthBackendGroupPtrOutput{})
	pulumi.RegisterOutputType(AuthBackendGroupArrayOutput{})
	pulumi.RegisterOutputType(AuthBackendGroupMapOutput{})
}
