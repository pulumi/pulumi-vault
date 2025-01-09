// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package approle

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-vault/sdk/v6/go/vault/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Logs into Vault using the AppRole auth backend. See the [Vault
// documentation](https://www.vaultproject.io/docs/auth/approle) for more
// information.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-vault/sdk/v6/go/vault"
//	"github.com/pulumi/pulumi-vault/sdk/v6/go/vault/approle"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			approle, err := vault.NewAuthBackend(ctx, "approle", &vault.AuthBackendArgs{
//				Type: pulumi.String("approle"),
//			})
//			if err != nil {
//				return err
//			}
//			example, err := approle.NewAuthBackendRole(ctx, "example", &approle.AuthBackendRoleArgs{
//				Backend:  approle.Path,
//				RoleName: pulumi.String("test-role"),
//				TokenPolicies: pulumi.StringArray{
//					pulumi.String("default"),
//					pulumi.String("dev"),
//					pulumi.String("prod"),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			id, err := approle.NewAuthBackendRoleSecretId(ctx, "id", &approle.AuthBackendRoleSecretIdArgs{
//				Backend:  approle.Path,
//				RoleName: example.RoleName,
//			})
//			if err != nil {
//				return err
//			}
//			_, err = approle.NewAuthBackendLogin(ctx, "login", &approle.AuthBackendLoginArgs{
//				Backend:  approle.Path,
//				RoleId:   example.RoleId,
//				SecretId: id.SecretId,
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
type AuthBackendLogin struct {
	pulumi.CustomResourceState

	// The accessor for the token.
	Accessor pulumi.StringOutput `pulumi:"accessor"`
	// The unique path of the Vault backend to log in with.
	Backend pulumi.StringPtrOutput `pulumi:"backend"`
	// The Vault token created.
	ClientToken pulumi.StringOutput `pulumi:"clientToken"`
	// How long the token is valid for, in seconds.
	LeaseDuration pulumi.IntOutput `pulumi:"leaseDuration"`
	// The date and time the lease started, in RFC 3339 format.
	LeaseStarted pulumi.StringOutput `pulumi:"leaseStarted"`
	// The metadata associated with the token.
	Metadata pulumi.StringMapOutput `pulumi:"metadata"`
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
	// *Available only for Vault Enterprise*.
	Namespace pulumi.StringPtrOutput `pulumi:"namespace"`
	// A list of policies applied to the token.
	Policies pulumi.StringArrayOutput `pulumi:"policies"`
	// Whether the token is renewable or not.
	Renewable pulumi.BoolOutput `pulumi:"renewable"`
	// The ID of the role to log in with.
	RoleId pulumi.StringOutput `pulumi:"roleId"`
	// The secret ID of the role to log in with. Required
	// unless `bindSecretId` is set to false on the role.
	SecretId pulumi.StringPtrOutput `pulumi:"secretId"`
}

// NewAuthBackendLogin registers a new resource with the given unique name, arguments, and options.
func NewAuthBackendLogin(ctx *pulumi.Context,
	name string, args *AuthBackendLoginArgs, opts ...pulumi.ResourceOption) (*AuthBackendLogin, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.RoleId == nil {
		return nil, errors.New("invalid value for required argument 'RoleId'")
	}
	if args.SecretId != nil {
		args.SecretId = pulumi.ToSecret(args.SecretId).(pulumi.StringPtrInput)
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"clientToken",
		"secretId",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource AuthBackendLogin
	err := ctx.RegisterResource("vault:appRole/authBackendLogin:AuthBackendLogin", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAuthBackendLogin gets an existing AuthBackendLogin resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAuthBackendLogin(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AuthBackendLoginState, opts ...pulumi.ResourceOption) (*AuthBackendLogin, error) {
	var resource AuthBackendLogin
	err := ctx.ReadResource("vault:appRole/authBackendLogin:AuthBackendLogin", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AuthBackendLogin resources.
type authBackendLoginState struct {
	// The accessor for the token.
	Accessor *string `pulumi:"accessor"`
	// The unique path of the Vault backend to log in with.
	Backend *string `pulumi:"backend"`
	// The Vault token created.
	ClientToken *string `pulumi:"clientToken"`
	// How long the token is valid for, in seconds.
	LeaseDuration *int `pulumi:"leaseDuration"`
	// The date and time the lease started, in RFC 3339 format.
	LeaseStarted *string `pulumi:"leaseStarted"`
	// The metadata associated with the token.
	Metadata map[string]string `pulumi:"metadata"`
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
	// *Available only for Vault Enterprise*.
	Namespace *string `pulumi:"namespace"`
	// A list of policies applied to the token.
	Policies []string `pulumi:"policies"`
	// Whether the token is renewable or not.
	Renewable *bool `pulumi:"renewable"`
	// The ID of the role to log in with.
	RoleId *string `pulumi:"roleId"`
	// The secret ID of the role to log in with. Required
	// unless `bindSecretId` is set to false on the role.
	SecretId *string `pulumi:"secretId"`
}

type AuthBackendLoginState struct {
	// The accessor for the token.
	Accessor pulumi.StringPtrInput
	// The unique path of the Vault backend to log in with.
	Backend pulumi.StringPtrInput
	// The Vault token created.
	ClientToken pulumi.StringPtrInput
	// How long the token is valid for, in seconds.
	LeaseDuration pulumi.IntPtrInput
	// The date and time the lease started, in RFC 3339 format.
	LeaseStarted pulumi.StringPtrInput
	// The metadata associated with the token.
	Metadata pulumi.StringMapInput
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
	// *Available only for Vault Enterprise*.
	Namespace pulumi.StringPtrInput
	// A list of policies applied to the token.
	Policies pulumi.StringArrayInput
	// Whether the token is renewable or not.
	Renewable pulumi.BoolPtrInput
	// The ID of the role to log in with.
	RoleId pulumi.StringPtrInput
	// The secret ID of the role to log in with. Required
	// unless `bindSecretId` is set to false on the role.
	SecretId pulumi.StringPtrInput
}

func (AuthBackendLoginState) ElementType() reflect.Type {
	return reflect.TypeOf((*authBackendLoginState)(nil)).Elem()
}

type authBackendLoginArgs struct {
	// The unique path of the Vault backend to log in with.
	Backend *string `pulumi:"backend"`
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
	// *Available only for Vault Enterprise*.
	Namespace *string `pulumi:"namespace"`
	// The ID of the role to log in with.
	RoleId string `pulumi:"roleId"`
	// The secret ID of the role to log in with. Required
	// unless `bindSecretId` is set to false on the role.
	SecretId *string `pulumi:"secretId"`
}

// The set of arguments for constructing a AuthBackendLogin resource.
type AuthBackendLoginArgs struct {
	// The unique path of the Vault backend to log in with.
	Backend pulumi.StringPtrInput
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
	// *Available only for Vault Enterprise*.
	Namespace pulumi.StringPtrInput
	// The ID of the role to log in with.
	RoleId pulumi.StringInput
	// The secret ID of the role to log in with. Required
	// unless `bindSecretId` is set to false on the role.
	SecretId pulumi.StringPtrInput
}

func (AuthBackendLoginArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*authBackendLoginArgs)(nil)).Elem()
}

type AuthBackendLoginInput interface {
	pulumi.Input

	ToAuthBackendLoginOutput() AuthBackendLoginOutput
	ToAuthBackendLoginOutputWithContext(ctx context.Context) AuthBackendLoginOutput
}

func (*AuthBackendLogin) ElementType() reflect.Type {
	return reflect.TypeOf((**AuthBackendLogin)(nil)).Elem()
}

func (i *AuthBackendLogin) ToAuthBackendLoginOutput() AuthBackendLoginOutput {
	return i.ToAuthBackendLoginOutputWithContext(context.Background())
}

func (i *AuthBackendLogin) ToAuthBackendLoginOutputWithContext(ctx context.Context) AuthBackendLoginOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AuthBackendLoginOutput)
}

// AuthBackendLoginArrayInput is an input type that accepts AuthBackendLoginArray and AuthBackendLoginArrayOutput values.
// You can construct a concrete instance of `AuthBackendLoginArrayInput` via:
//
//	AuthBackendLoginArray{ AuthBackendLoginArgs{...} }
type AuthBackendLoginArrayInput interface {
	pulumi.Input

	ToAuthBackendLoginArrayOutput() AuthBackendLoginArrayOutput
	ToAuthBackendLoginArrayOutputWithContext(context.Context) AuthBackendLoginArrayOutput
}

type AuthBackendLoginArray []AuthBackendLoginInput

func (AuthBackendLoginArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AuthBackendLogin)(nil)).Elem()
}

func (i AuthBackendLoginArray) ToAuthBackendLoginArrayOutput() AuthBackendLoginArrayOutput {
	return i.ToAuthBackendLoginArrayOutputWithContext(context.Background())
}

func (i AuthBackendLoginArray) ToAuthBackendLoginArrayOutputWithContext(ctx context.Context) AuthBackendLoginArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AuthBackendLoginArrayOutput)
}

// AuthBackendLoginMapInput is an input type that accepts AuthBackendLoginMap and AuthBackendLoginMapOutput values.
// You can construct a concrete instance of `AuthBackendLoginMapInput` via:
//
//	AuthBackendLoginMap{ "key": AuthBackendLoginArgs{...} }
type AuthBackendLoginMapInput interface {
	pulumi.Input

	ToAuthBackendLoginMapOutput() AuthBackendLoginMapOutput
	ToAuthBackendLoginMapOutputWithContext(context.Context) AuthBackendLoginMapOutput
}

type AuthBackendLoginMap map[string]AuthBackendLoginInput

func (AuthBackendLoginMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AuthBackendLogin)(nil)).Elem()
}

func (i AuthBackendLoginMap) ToAuthBackendLoginMapOutput() AuthBackendLoginMapOutput {
	return i.ToAuthBackendLoginMapOutputWithContext(context.Background())
}

func (i AuthBackendLoginMap) ToAuthBackendLoginMapOutputWithContext(ctx context.Context) AuthBackendLoginMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AuthBackendLoginMapOutput)
}

type AuthBackendLoginOutput struct{ *pulumi.OutputState }

func (AuthBackendLoginOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AuthBackendLogin)(nil)).Elem()
}

func (o AuthBackendLoginOutput) ToAuthBackendLoginOutput() AuthBackendLoginOutput {
	return o
}

func (o AuthBackendLoginOutput) ToAuthBackendLoginOutputWithContext(ctx context.Context) AuthBackendLoginOutput {
	return o
}

// The accessor for the token.
func (o AuthBackendLoginOutput) Accessor() pulumi.StringOutput {
	return o.ApplyT(func(v *AuthBackendLogin) pulumi.StringOutput { return v.Accessor }).(pulumi.StringOutput)
}

// The unique path of the Vault backend to log in with.
func (o AuthBackendLoginOutput) Backend() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AuthBackendLogin) pulumi.StringPtrOutput { return v.Backend }).(pulumi.StringPtrOutput)
}

// The Vault token created.
func (o AuthBackendLoginOutput) ClientToken() pulumi.StringOutput {
	return o.ApplyT(func(v *AuthBackendLogin) pulumi.StringOutput { return v.ClientToken }).(pulumi.StringOutput)
}

// How long the token is valid for, in seconds.
func (o AuthBackendLoginOutput) LeaseDuration() pulumi.IntOutput {
	return o.ApplyT(func(v *AuthBackendLogin) pulumi.IntOutput { return v.LeaseDuration }).(pulumi.IntOutput)
}

// The date and time the lease started, in RFC 3339 format.
func (o AuthBackendLoginOutput) LeaseStarted() pulumi.StringOutput {
	return o.ApplyT(func(v *AuthBackendLogin) pulumi.StringOutput { return v.LeaseStarted }).(pulumi.StringOutput)
}

// The metadata associated with the token.
func (o AuthBackendLoginOutput) Metadata() pulumi.StringMapOutput {
	return o.ApplyT(func(v *AuthBackendLogin) pulumi.StringMapOutput { return v.Metadata }).(pulumi.StringMapOutput)
}

// The namespace to provision the resource in.
// The value should not contain leading or trailing forward slashes.
// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
// *Available only for Vault Enterprise*.
func (o AuthBackendLoginOutput) Namespace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AuthBackendLogin) pulumi.StringPtrOutput { return v.Namespace }).(pulumi.StringPtrOutput)
}

// A list of policies applied to the token.
func (o AuthBackendLoginOutput) Policies() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *AuthBackendLogin) pulumi.StringArrayOutput { return v.Policies }).(pulumi.StringArrayOutput)
}

// Whether the token is renewable or not.
func (o AuthBackendLoginOutput) Renewable() pulumi.BoolOutput {
	return o.ApplyT(func(v *AuthBackendLogin) pulumi.BoolOutput { return v.Renewable }).(pulumi.BoolOutput)
}

// The ID of the role to log in with.
func (o AuthBackendLoginOutput) RoleId() pulumi.StringOutput {
	return o.ApplyT(func(v *AuthBackendLogin) pulumi.StringOutput { return v.RoleId }).(pulumi.StringOutput)
}

// The secret ID of the role to log in with. Required
// unless `bindSecretId` is set to false on the role.
func (o AuthBackendLoginOutput) SecretId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AuthBackendLogin) pulumi.StringPtrOutput { return v.SecretId }).(pulumi.StringPtrOutput)
}

type AuthBackendLoginArrayOutput struct{ *pulumi.OutputState }

func (AuthBackendLoginArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AuthBackendLogin)(nil)).Elem()
}

func (o AuthBackendLoginArrayOutput) ToAuthBackendLoginArrayOutput() AuthBackendLoginArrayOutput {
	return o
}

func (o AuthBackendLoginArrayOutput) ToAuthBackendLoginArrayOutputWithContext(ctx context.Context) AuthBackendLoginArrayOutput {
	return o
}

func (o AuthBackendLoginArrayOutput) Index(i pulumi.IntInput) AuthBackendLoginOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *AuthBackendLogin {
		return vs[0].([]*AuthBackendLogin)[vs[1].(int)]
	}).(AuthBackendLoginOutput)
}

type AuthBackendLoginMapOutput struct{ *pulumi.OutputState }

func (AuthBackendLoginMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AuthBackendLogin)(nil)).Elem()
}

func (o AuthBackendLoginMapOutput) ToAuthBackendLoginMapOutput() AuthBackendLoginMapOutput {
	return o
}

func (o AuthBackendLoginMapOutput) ToAuthBackendLoginMapOutputWithContext(ctx context.Context) AuthBackendLoginMapOutput {
	return o
}

func (o AuthBackendLoginMapOutput) MapIndex(k pulumi.StringInput) AuthBackendLoginOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *AuthBackendLogin {
		return vs[0].(map[string]*AuthBackendLogin)[vs[1].(string)]
	}).(AuthBackendLoginOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AuthBackendLoginInput)(nil)).Elem(), &AuthBackendLogin{})
	pulumi.RegisterInputType(reflect.TypeOf((*AuthBackendLoginArrayInput)(nil)).Elem(), AuthBackendLoginArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AuthBackendLoginMapInput)(nil)).Elem(), AuthBackendLoginMap{})
	pulumi.RegisterOutputType(AuthBackendLoginOutput{})
	pulumi.RegisterOutputType(AuthBackendLoginArrayOutput{})
	pulumi.RegisterOutputType(AuthBackendLoginMapOutput{})
}
