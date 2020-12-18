// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package approle

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Logs into Vault using the AppRole auth backend. See the [Vault
// documentation](https://www.vaultproject.io/docs/auth/approle.html) for more
// information.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-vault/sdk/v3/go/vault"
// 	"github.com/pulumi/pulumi-vault/sdk/v3/go/vault/appRole"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		approle, err := vault.NewAuthBackend(ctx, "approle", &vault.AuthBackendArgs{
// 			Type: pulumi.String("approle"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		example, err := appRole.NewAuthBackendRole(ctx, "example", &appRole.AuthBackendRoleArgs{
// 			Backend: approle.Path,
// 			Policies: pulumi.StringArray{
// 				pulumi.String("default"),
// 				pulumi.String("dev"),
// 				pulumi.String("prod"),
// 			},
// 			RoleName: pulumi.String("test-role"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		id, err := appRole.NewAuthBackendRoleSecretID(ctx, "id", &appRole.AuthBackendRoleSecretIDArgs{
// 			Backend:  approle.Path,
// 			RoleName: example.RoleName,
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = appRole.NewAuthBackendLogin(ctx, "login", &appRole.AuthBackendLoginArgs{
// 			Backend:  approle.Path,
// 			RoleId:   example.RoleId,
// 			SecretId: id.SecretId,
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
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
	if args == nil || args.RoleId == nil {
		return nil, errors.New("missing required argument 'RoleId'")
	}
	if args == nil {
		args = &AuthBackendLoginArgs{}
	}
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

func (AuthBackendLogin) ElementType() reflect.Type {
	return reflect.TypeOf((*AuthBackendLogin)(nil)).Elem()
}

func (i AuthBackendLogin) ToAuthBackendLoginOutput() AuthBackendLoginOutput {
	return i.ToAuthBackendLoginOutputWithContext(context.Background())
}

func (i AuthBackendLogin) ToAuthBackendLoginOutputWithContext(ctx context.Context) AuthBackendLoginOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AuthBackendLoginOutput)
}

type AuthBackendLoginOutput struct {
	*pulumi.OutputState
}

func (AuthBackendLoginOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AuthBackendLoginOutput)(nil)).Elem()
}

func (o AuthBackendLoginOutput) ToAuthBackendLoginOutput() AuthBackendLoginOutput {
	return o
}

func (o AuthBackendLoginOutput) ToAuthBackendLoginOutputWithContext(ctx context.Context) AuthBackendLoginOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(AuthBackendLoginOutput{})
}
