// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ldap

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-vault/sdk/v6/go/vault/internal"
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
//	"github.com/pulumi/pulumi-vault/sdk/v6/go/vault/ldap"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			config, err := ldap.NewSecretBackend(ctx, "config", &ldap.SecretBackendArgs{
//				Path:        pulumi.String("my-custom-ldap"),
//				Binddn:      pulumi.String("CN=Administrator,CN=Users,DC=corp,DC=example,DC=net"),
//				Bindpass:    pulumi.String("SuperSecretPassw0rd"),
//				Url:         pulumi.String("ldaps://localhost"),
//				InsecureTls: pulumi.Bool(true),
//				Userdn:      pulumi.String("CN=Users,DC=corp,DC=example,DC=net"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = ldap.NewSecretBackendStaticRole(ctx, "role", &ldap.SecretBackendStaticRoleArgs{
//				Mount:          config.Path,
//				Username:       pulumi.String("alice"),
//				Dn:             pulumi.String("cn=alice,ou=Users,DC=corp,DC=example,DC=net"),
//				RoleName:       pulumi.String("alice"),
//				RotationPeriod: pulumi.Int(60),
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
// LDAP secret backend static role can be imported using the full path to the role
// of the form: `<mount_path>/static-role/<role_name>` e.g.
//
// ```sh
// $ pulumi import vault:ldap/secretBackendStaticRole:SecretBackendStaticRole role ldap/static-role/example-role
// ```
type SecretBackendStaticRole struct {
	pulumi.CustomResourceState

	// Distinguished name (DN) of the existing LDAP entry to manage
	// password rotation for. If given, it will take precedence over `username` for the LDAP
	// search performed during password rotation. Cannot be modified after creation.
	Dn pulumi.StringPtrOutput `pulumi:"dn"`
	// The unique path this backend should be mounted at. Must
	// not begin or end with a `/`. Defaults to `ldap`.
	Mount pulumi.StringPtrOutput `pulumi:"mount"`
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
	// *Available only for Vault Enterprise*.
	Namespace pulumi.StringPtrOutput `pulumi:"namespace"`
	// Name of the role.
	RoleName pulumi.StringOutput `pulumi:"roleName"`
	// How often Vault should rotate the password of the user entry.
	RotationPeriod pulumi.IntOutput `pulumi:"rotationPeriod"`
	// Causes vault to skip the initial secret rotation on import. Not applicable to updates.
	// Requires Vault 1.16 or above.
	SkipImportRotation pulumi.BoolPtrOutput `pulumi:"skipImportRotation"`
	// The username of the existing LDAP entry to manage password rotation for.
	Username pulumi.StringOutput `pulumi:"username"`
}

// NewSecretBackendStaticRole registers a new resource with the given unique name, arguments, and options.
func NewSecretBackendStaticRole(ctx *pulumi.Context,
	name string, args *SecretBackendStaticRoleArgs, opts ...pulumi.ResourceOption) (*SecretBackendStaticRole, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.RoleName == nil {
		return nil, errors.New("invalid value for required argument 'RoleName'")
	}
	if args.RotationPeriod == nil {
		return nil, errors.New("invalid value for required argument 'RotationPeriod'")
	}
	if args.Username == nil {
		return nil, errors.New("invalid value for required argument 'Username'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource SecretBackendStaticRole
	err := ctx.RegisterResource("vault:ldap/secretBackendStaticRole:SecretBackendStaticRole", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSecretBackendStaticRole gets an existing SecretBackendStaticRole resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSecretBackendStaticRole(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SecretBackendStaticRoleState, opts ...pulumi.ResourceOption) (*SecretBackendStaticRole, error) {
	var resource SecretBackendStaticRole
	err := ctx.ReadResource("vault:ldap/secretBackendStaticRole:SecretBackendStaticRole", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SecretBackendStaticRole resources.
type secretBackendStaticRoleState struct {
	// Distinguished name (DN) of the existing LDAP entry to manage
	// password rotation for. If given, it will take precedence over `username` for the LDAP
	// search performed during password rotation. Cannot be modified after creation.
	Dn *string `pulumi:"dn"`
	// The unique path this backend should be mounted at. Must
	// not begin or end with a `/`. Defaults to `ldap`.
	Mount *string `pulumi:"mount"`
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
	// *Available only for Vault Enterprise*.
	Namespace *string `pulumi:"namespace"`
	// Name of the role.
	RoleName *string `pulumi:"roleName"`
	// How often Vault should rotate the password of the user entry.
	RotationPeriod *int `pulumi:"rotationPeriod"`
	// Causes vault to skip the initial secret rotation on import. Not applicable to updates.
	// Requires Vault 1.16 or above.
	SkipImportRotation *bool `pulumi:"skipImportRotation"`
	// The username of the existing LDAP entry to manage password rotation for.
	Username *string `pulumi:"username"`
}

type SecretBackendStaticRoleState struct {
	// Distinguished name (DN) of the existing LDAP entry to manage
	// password rotation for. If given, it will take precedence over `username` for the LDAP
	// search performed during password rotation. Cannot be modified after creation.
	Dn pulumi.StringPtrInput
	// The unique path this backend should be mounted at. Must
	// not begin or end with a `/`. Defaults to `ldap`.
	Mount pulumi.StringPtrInput
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
	// *Available only for Vault Enterprise*.
	Namespace pulumi.StringPtrInput
	// Name of the role.
	RoleName pulumi.StringPtrInput
	// How often Vault should rotate the password of the user entry.
	RotationPeriod pulumi.IntPtrInput
	// Causes vault to skip the initial secret rotation on import. Not applicable to updates.
	// Requires Vault 1.16 or above.
	SkipImportRotation pulumi.BoolPtrInput
	// The username of the existing LDAP entry to manage password rotation for.
	Username pulumi.StringPtrInput
}

func (SecretBackendStaticRoleState) ElementType() reflect.Type {
	return reflect.TypeOf((*secretBackendStaticRoleState)(nil)).Elem()
}

type secretBackendStaticRoleArgs struct {
	// Distinguished name (DN) of the existing LDAP entry to manage
	// password rotation for. If given, it will take precedence over `username` for the LDAP
	// search performed during password rotation. Cannot be modified after creation.
	Dn *string `pulumi:"dn"`
	// The unique path this backend should be mounted at. Must
	// not begin or end with a `/`. Defaults to `ldap`.
	Mount *string `pulumi:"mount"`
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
	// *Available only for Vault Enterprise*.
	Namespace *string `pulumi:"namespace"`
	// Name of the role.
	RoleName string `pulumi:"roleName"`
	// How often Vault should rotate the password of the user entry.
	RotationPeriod int `pulumi:"rotationPeriod"`
	// Causes vault to skip the initial secret rotation on import. Not applicable to updates.
	// Requires Vault 1.16 or above.
	SkipImportRotation *bool `pulumi:"skipImportRotation"`
	// The username of the existing LDAP entry to manage password rotation for.
	Username string `pulumi:"username"`
}

// The set of arguments for constructing a SecretBackendStaticRole resource.
type SecretBackendStaticRoleArgs struct {
	// Distinguished name (DN) of the existing LDAP entry to manage
	// password rotation for. If given, it will take precedence over `username` for the LDAP
	// search performed during password rotation. Cannot be modified after creation.
	Dn pulumi.StringPtrInput
	// The unique path this backend should be mounted at. Must
	// not begin or end with a `/`. Defaults to `ldap`.
	Mount pulumi.StringPtrInput
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
	// *Available only for Vault Enterprise*.
	Namespace pulumi.StringPtrInput
	// Name of the role.
	RoleName pulumi.StringInput
	// How often Vault should rotate the password of the user entry.
	RotationPeriod pulumi.IntInput
	// Causes vault to skip the initial secret rotation on import. Not applicable to updates.
	// Requires Vault 1.16 or above.
	SkipImportRotation pulumi.BoolPtrInput
	// The username of the existing LDAP entry to manage password rotation for.
	Username pulumi.StringInput
}

func (SecretBackendStaticRoleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*secretBackendStaticRoleArgs)(nil)).Elem()
}

type SecretBackendStaticRoleInput interface {
	pulumi.Input

	ToSecretBackendStaticRoleOutput() SecretBackendStaticRoleOutput
	ToSecretBackendStaticRoleOutputWithContext(ctx context.Context) SecretBackendStaticRoleOutput
}

func (*SecretBackendStaticRole) ElementType() reflect.Type {
	return reflect.TypeOf((**SecretBackendStaticRole)(nil)).Elem()
}

func (i *SecretBackendStaticRole) ToSecretBackendStaticRoleOutput() SecretBackendStaticRoleOutput {
	return i.ToSecretBackendStaticRoleOutputWithContext(context.Background())
}

func (i *SecretBackendStaticRole) ToSecretBackendStaticRoleOutputWithContext(ctx context.Context) SecretBackendStaticRoleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecretBackendStaticRoleOutput)
}

// SecretBackendStaticRoleArrayInput is an input type that accepts SecretBackendStaticRoleArray and SecretBackendStaticRoleArrayOutput values.
// You can construct a concrete instance of `SecretBackendStaticRoleArrayInput` via:
//
//	SecretBackendStaticRoleArray{ SecretBackendStaticRoleArgs{...} }
type SecretBackendStaticRoleArrayInput interface {
	pulumi.Input

	ToSecretBackendStaticRoleArrayOutput() SecretBackendStaticRoleArrayOutput
	ToSecretBackendStaticRoleArrayOutputWithContext(context.Context) SecretBackendStaticRoleArrayOutput
}

type SecretBackendStaticRoleArray []SecretBackendStaticRoleInput

func (SecretBackendStaticRoleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SecretBackendStaticRole)(nil)).Elem()
}

func (i SecretBackendStaticRoleArray) ToSecretBackendStaticRoleArrayOutput() SecretBackendStaticRoleArrayOutput {
	return i.ToSecretBackendStaticRoleArrayOutputWithContext(context.Background())
}

func (i SecretBackendStaticRoleArray) ToSecretBackendStaticRoleArrayOutputWithContext(ctx context.Context) SecretBackendStaticRoleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecretBackendStaticRoleArrayOutput)
}

// SecretBackendStaticRoleMapInput is an input type that accepts SecretBackendStaticRoleMap and SecretBackendStaticRoleMapOutput values.
// You can construct a concrete instance of `SecretBackendStaticRoleMapInput` via:
//
//	SecretBackendStaticRoleMap{ "key": SecretBackendStaticRoleArgs{...} }
type SecretBackendStaticRoleMapInput interface {
	pulumi.Input

	ToSecretBackendStaticRoleMapOutput() SecretBackendStaticRoleMapOutput
	ToSecretBackendStaticRoleMapOutputWithContext(context.Context) SecretBackendStaticRoleMapOutput
}

type SecretBackendStaticRoleMap map[string]SecretBackendStaticRoleInput

func (SecretBackendStaticRoleMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SecretBackendStaticRole)(nil)).Elem()
}

func (i SecretBackendStaticRoleMap) ToSecretBackendStaticRoleMapOutput() SecretBackendStaticRoleMapOutput {
	return i.ToSecretBackendStaticRoleMapOutputWithContext(context.Background())
}

func (i SecretBackendStaticRoleMap) ToSecretBackendStaticRoleMapOutputWithContext(ctx context.Context) SecretBackendStaticRoleMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecretBackendStaticRoleMapOutput)
}

type SecretBackendStaticRoleOutput struct{ *pulumi.OutputState }

func (SecretBackendStaticRoleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SecretBackendStaticRole)(nil)).Elem()
}

func (o SecretBackendStaticRoleOutput) ToSecretBackendStaticRoleOutput() SecretBackendStaticRoleOutput {
	return o
}

func (o SecretBackendStaticRoleOutput) ToSecretBackendStaticRoleOutputWithContext(ctx context.Context) SecretBackendStaticRoleOutput {
	return o
}

// Distinguished name (DN) of the existing LDAP entry to manage
// password rotation for. If given, it will take precedence over `username` for the LDAP
// search performed during password rotation. Cannot be modified after creation.
func (o SecretBackendStaticRoleOutput) Dn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecretBackendStaticRole) pulumi.StringPtrOutput { return v.Dn }).(pulumi.StringPtrOutput)
}

// The unique path this backend should be mounted at. Must
// not begin or end with a `/`. Defaults to `ldap`.
func (o SecretBackendStaticRoleOutput) Mount() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecretBackendStaticRole) pulumi.StringPtrOutput { return v.Mount }).(pulumi.StringPtrOutput)
}

// The namespace to provision the resource in.
// The value should not contain leading or trailing forward slashes.
// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
// *Available only for Vault Enterprise*.
func (o SecretBackendStaticRoleOutput) Namespace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecretBackendStaticRole) pulumi.StringPtrOutput { return v.Namespace }).(pulumi.StringPtrOutput)
}

// Name of the role.
func (o SecretBackendStaticRoleOutput) RoleName() pulumi.StringOutput {
	return o.ApplyT(func(v *SecretBackendStaticRole) pulumi.StringOutput { return v.RoleName }).(pulumi.StringOutput)
}

// How often Vault should rotate the password of the user entry.
func (o SecretBackendStaticRoleOutput) RotationPeriod() pulumi.IntOutput {
	return o.ApplyT(func(v *SecretBackendStaticRole) pulumi.IntOutput { return v.RotationPeriod }).(pulumi.IntOutput)
}

// Causes vault to skip the initial secret rotation on import. Not applicable to updates.
// Requires Vault 1.16 or above.
func (o SecretBackendStaticRoleOutput) SkipImportRotation() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SecretBackendStaticRole) pulumi.BoolPtrOutput { return v.SkipImportRotation }).(pulumi.BoolPtrOutput)
}

// The username of the existing LDAP entry to manage password rotation for.
func (o SecretBackendStaticRoleOutput) Username() pulumi.StringOutput {
	return o.ApplyT(func(v *SecretBackendStaticRole) pulumi.StringOutput { return v.Username }).(pulumi.StringOutput)
}

type SecretBackendStaticRoleArrayOutput struct{ *pulumi.OutputState }

func (SecretBackendStaticRoleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SecretBackendStaticRole)(nil)).Elem()
}

func (o SecretBackendStaticRoleArrayOutput) ToSecretBackendStaticRoleArrayOutput() SecretBackendStaticRoleArrayOutput {
	return o
}

func (o SecretBackendStaticRoleArrayOutput) ToSecretBackendStaticRoleArrayOutputWithContext(ctx context.Context) SecretBackendStaticRoleArrayOutput {
	return o
}

func (o SecretBackendStaticRoleArrayOutput) Index(i pulumi.IntInput) SecretBackendStaticRoleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SecretBackendStaticRole {
		return vs[0].([]*SecretBackendStaticRole)[vs[1].(int)]
	}).(SecretBackendStaticRoleOutput)
}

type SecretBackendStaticRoleMapOutput struct{ *pulumi.OutputState }

func (SecretBackendStaticRoleMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SecretBackendStaticRole)(nil)).Elem()
}

func (o SecretBackendStaticRoleMapOutput) ToSecretBackendStaticRoleMapOutput() SecretBackendStaticRoleMapOutput {
	return o
}

func (o SecretBackendStaticRoleMapOutput) ToSecretBackendStaticRoleMapOutputWithContext(ctx context.Context) SecretBackendStaticRoleMapOutput {
	return o
}

func (o SecretBackendStaticRoleMapOutput) MapIndex(k pulumi.StringInput) SecretBackendStaticRoleOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SecretBackendStaticRole {
		return vs[0].(map[string]*SecretBackendStaticRole)[vs[1].(string)]
	}).(SecretBackendStaticRoleOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SecretBackendStaticRoleInput)(nil)).Elem(), &SecretBackendStaticRole{})
	pulumi.RegisterInputType(reflect.TypeOf((*SecretBackendStaticRoleArrayInput)(nil)).Elem(), SecretBackendStaticRoleArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SecretBackendStaticRoleMapInput)(nil)).Elem(), SecretBackendStaticRoleMap{})
	pulumi.RegisterOutputType(SecretBackendStaticRoleOutput{})
	pulumi.RegisterOutputType(SecretBackendStaticRoleArrayOutput{})
	pulumi.RegisterOutputType(SecretBackendStaticRoleMapOutput{})
}
