// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ldap

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-vault/sdk/v5/go/vault/ldap"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			config, err := ldap.NewSecretBackend(ctx, "config", &ldap.SecretBackendArgs{
//				Path:     pulumi.String("my-custom-ldap"),
//				Binddn:   pulumi.String("CN=Administrator,CN=Users,DC=corp,DC=example,DC=net"),
//				Bindpass: pulumi.String("SuperSecretPassw0rd"),
//				Url:      pulumi.String("ldaps://localhost"),
//				Userdn:   pulumi.String("CN=Users,DC=corp,DC=example,DC=net"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = ldap.NewSecretBackendDynamicRole(ctx, "role", &ldap.SecretBackendDynamicRoleArgs{
//				Mount:        config.Path,
//				RoleName:     pulumi.String("alice"),
//				CreationLdif: pulumi.String("dn: cn={{.Username}},ou=users,dc=learn,dc=example\nobjectClass: person\nobjectClass: top\ncn: learn\nsn: {{.Password | utf16le | base64}}\nmemberOf: cn=dev,ou=groups,dc=learn,dc=example\nuserPassword: {{.Password}}\n"),
//				DeletionLdif: pulumi.String("dn: cn={{.Username}},ou=users,dc=learn,dc=example\nchangetype: delete\n  rollback_ldif = <<EOT\ndn: cn={{.Username}},ou=users,dc=learn,dc=example\nchangetype: delete\n"),
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
// LDAP secret backend dynamic role can be imported using the full path to the role of the form`<mount_path>/dynamic-role/<role_name>` e.g.
//
// ```sh
//
//	$ pulumi import vault:ldap/secretBackendDynamicRole:SecretBackendDynamicRole role ldap/role/dynamic-role
//
// ```
type SecretBackendDynamicRole struct {
	pulumi.CustomResourceState

	// A templatized LDIF string used to create a user
	// account. This may contain multiple LDIF entries. The `creationLdif` can also
	// be used to add the user account to an existing group. All LDIF entries are
	// performed in order. If Vault encounters an error while executing the
	// `creationLdif` it will stop at the first error and not execute any remaining
	// LDIF entries. If an error occurs and `rollbackLdif` is specified, the LDIF
	// entries in `rollbackLdif` will be executed. See `rollbackLdif` for more
	// details. This field may optionally be provided as a base64 encoded string.
	CreationLdif pulumi.StringOutput `pulumi:"creationLdif"`
	// Specifies the TTL for the leases associated with this role.
	DefaultTtl pulumi.IntPtrOutput `pulumi:"defaultTtl"`
	// A templatized LDIF string used to delete the
	// user account once its TTL has expired. This may contain multiple LDIF
	// entries. All LDIF entries are performed in order. If Vault encounters an
	// error while executing an entry in the `deletionLdif` it will attempt to
	// continue executing any remaining entries. This field may optionally be
	// provided as a base64 encoded string.
	DeletionLdif pulumi.StringOutput `pulumi:"deletionLdif"`
	// Specifies the maximum TTL for the leases associated with this role.
	MaxTtl pulumi.IntPtrOutput `pulumi:"maxTtl"`
	// The unique path this backend should be mounted at. Must
	// not begin or end with a `/`. Defaults to `ldap`.
	Mount pulumi.StringPtrOutput `pulumi:"mount"`
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
	// *Available only for Vault Enterprise*.
	Namespace pulumi.StringPtrOutput `pulumi:"namespace"`
	// Name of the role.
	RoleName pulumi.StringOutput `pulumi:"roleName"`
	// A templatized LDIF string used to attempt to
	// rollback any changes in the event that execution of the `creationLdif` results
	// in an error. This may contain multiple LDIF entries. All LDIF entries are
	// performed in order. If Vault encounters an error while executing an entry in
	// the `rollbackLdif` it will attempt to continue executing any remaining
	// entries. This field may optionally be provided as a base64 encoded string.
	RollbackLdif pulumi.StringPtrOutput `pulumi:"rollbackLdif"`
	// A template used to generate a dynamic
	// username. This will be used to fill in the `.Username` field within the
	// `creationLdif` string.
	UsernameTemplate pulumi.StringPtrOutput `pulumi:"usernameTemplate"`
}

// NewSecretBackendDynamicRole registers a new resource with the given unique name, arguments, and options.
func NewSecretBackendDynamicRole(ctx *pulumi.Context,
	name string, args *SecretBackendDynamicRoleArgs, opts ...pulumi.ResourceOption) (*SecretBackendDynamicRole, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.CreationLdif == nil {
		return nil, errors.New("invalid value for required argument 'CreationLdif'")
	}
	if args.DeletionLdif == nil {
		return nil, errors.New("invalid value for required argument 'DeletionLdif'")
	}
	if args.RoleName == nil {
		return nil, errors.New("invalid value for required argument 'RoleName'")
	}
	var resource SecretBackendDynamicRole
	err := ctx.RegisterResource("vault:ldap/secretBackendDynamicRole:SecretBackendDynamicRole", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSecretBackendDynamicRole gets an existing SecretBackendDynamicRole resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSecretBackendDynamicRole(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SecretBackendDynamicRoleState, opts ...pulumi.ResourceOption) (*SecretBackendDynamicRole, error) {
	var resource SecretBackendDynamicRole
	err := ctx.ReadResource("vault:ldap/secretBackendDynamicRole:SecretBackendDynamicRole", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SecretBackendDynamicRole resources.
type secretBackendDynamicRoleState struct {
	// A templatized LDIF string used to create a user
	// account. This may contain multiple LDIF entries. The `creationLdif` can also
	// be used to add the user account to an existing group. All LDIF entries are
	// performed in order. If Vault encounters an error while executing the
	// `creationLdif` it will stop at the first error and not execute any remaining
	// LDIF entries. If an error occurs and `rollbackLdif` is specified, the LDIF
	// entries in `rollbackLdif` will be executed. See `rollbackLdif` for more
	// details. This field may optionally be provided as a base64 encoded string.
	CreationLdif *string `pulumi:"creationLdif"`
	// Specifies the TTL for the leases associated with this role.
	DefaultTtl *int `pulumi:"defaultTtl"`
	// A templatized LDIF string used to delete the
	// user account once its TTL has expired. This may contain multiple LDIF
	// entries. All LDIF entries are performed in order. If Vault encounters an
	// error while executing an entry in the `deletionLdif` it will attempt to
	// continue executing any remaining entries. This field may optionally be
	// provided as a base64 encoded string.
	DeletionLdif *string `pulumi:"deletionLdif"`
	// Specifies the maximum TTL for the leases associated with this role.
	MaxTtl *int `pulumi:"maxTtl"`
	// The unique path this backend should be mounted at. Must
	// not begin or end with a `/`. Defaults to `ldap`.
	Mount *string `pulumi:"mount"`
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
	// *Available only for Vault Enterprise*.
	Namespace *string `pulumi:"namespace"`
	// Name of the role.
	RoleName *string `pulumi:"roleName"`
	// A templatized LDIF string used to attempt to
	// rollback any changes in the event that execution of the `creationLdif` results
	// in an error. This may contain multiple LDIF entries. All LDIF entries are
	// performed in order. If Vault encounters an error while executing an entry in
	// the `rollbackLdif` it will attempt to continue executing any remaining
	// entries. This field may optionally be provided as a base64 encoded string.
	RollbackLdif *string `pulumi:"rollbackLdif"`
	// A template used to generate a dynamic
	// username. This will be used to fill in the `.Username` field within the
	// `creationLdif` string.
	UsernameTemplate *string `pulumi:"usernameTemplate"`
}

type SecretBackendDynamicRoleState struct {
	// A templatized LDIF string used to create a user
	// account. This may contain multiple LDIF entries. The `creationLdif` can also
	// be used to add the user account to an existing group. All LDIF entries are
	// performed in order. If Vault encounters an error while executing the
	// `creationLdif` it will stop at the first error and not execute any remaining
	// LDIF entries. If an error occurs and `rollbackLdif` is specified, the LDIF
	// entries in `rollbackLdif` will be executed. See `rollbackLdif` for more
	// details. This field may optionally be provided as a base64 encoded string.
	CreationLdif pulumi.StringPtrInput
	// Specifies the TTL for the leases associated with this role.
	DefaultTtl pulumi.IntPtrInput
	// A templatized LDIF string used to delete the
	// user account once its TTL has expired. This may contain multiple LDIF
	// entries. All LDIF entries are performed in order. If Vault encounters an
	// error while executing an entry in the `deletionLdif` it will attempt to
	// continue executing any remaining entries. This field may optionally be
	// provided as a base64 encoded string.
	DeletionLdif pulumi.StringPtrInput
	// Specifies the maximum TTL for the leases associated with this role.
	MaxTtl pulumi.IntPtrInput
	// The unique path this backend should be mounted at. Must
	// not begin or end with a `/`. Defaults to `ldap`.
	Mount pulumi.StringPtrInput
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
	// *Available only for Vault Enterprise*.
	Namespace pulumi.StringPtrInput
	// Name of the role.
	RoleName pulumi.StringPtrInput
	// A templatized LDIF string used to attempt to
	// rollback any changes in the event that execution of the `creationLdif` results
	// in an error. This may contain multiple LDIF entries. All LDIF entries are
	// performed in order. If Vault encounters an error while executing an entry in
	// the `rollbackLdif` it will attempt to continue executing any remaining
	// entries. This field may optionally be provided as a base64 encoded string.
	RollbackLdif pulumi.StringPtrInput
	// A template used to generate a dynamic
	// username. This will be used to fill in the `.Username` field within the
	// `creationLdif` string.
	UsernameTemplate pulumi.StringPtrInput
}

func (SecretBackendDynamicRoleState) ElementType() reflect.Type {
	return reflect.TypeOf((*secretBackendDynamicRoleState)(nil)).Elem()
}

type secretBackendDynamicRoleArgs struct {
	// A templatized LDIF string used to create a user
	// account. This may contain multiple LDIF entries. The `creationLdif` can also
	// be used to add the user account to an existing group. All LDIF entries are
	// performed in order. If Vault encounters an error while executing the
	// `creationLdif` it will stop at the first error and not execute any remaining
	// LDIF entries. If an error occurs and `rollbackLdif` is specified, the LDIF
	// entries in `rollbackLdif` will be executed. See `rollbackLdif` for more
	// details. This field may optionally be provided as a base64 encoded string.
	CreationLdif string `pulumi:"creationLdif"`
	// Specifies the TTL for the leases associated with this role.
	DefaultTtl *int `pulumi:"defaultTtl"`
	// A templatized LDIF string used to delete the
	// user account once its TTL has expired. This may contain multiple LDIF
	// entries. All LDIF entries are performed in order. If Vault encounters an
	// error while executing an entry in the `deletionLdif` it will attempt to
	// continue executing any remaining entries. This field may optionally be
	// provided as a base64 encoded string.
	DeletionLdif string `pulumi:"deletionLdif"`
	// Specifies the maximum TTL for the leases associated with this role.
	MaxTtl *int `pulumi:"maxTtl"`
	// The unique path this backend should be mounted at. Must
	// not begin or end with a `/`. Defaults to `ldap`.
	Mount *string `pulumi:"mount"`
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
	// *Available only for Vault Enterprise*.
	Namespace *string `pulumi:"namespace"`
	// Name of the role.
	RoleName string `pulumi:"roleName"`
	// A templatized LDIF string used to attempt to
	// rollback any changes in the event that execution of the `creationLdif` results
	// in an error. This may contain multiple LDIF entries. All LDIF entries are
	// performed in order. If Vault encounters an error while executing an entry in
	// the `rollbackLdif` it will attempt to continue executing any remaining
	// entries. This field may optionally be provided as a base64 encoded string.
	RollbackLdif *string `pulumi:"rollbackLdif"`
	// A template used to generate a dynamic
	// username. This will be used to fill in the `.Username` field within the
	// `creationLdif` string.
	UsernameTemplate *string `pulumi:"usernameTemplate"`
}

// The set of arguments for constructing a SecretBackendDynamicRole resource.
type SecretBackendDynamicRoleArgs struct {
	// A templatized LDIF string used to create a user
	// account. This may contain multiple LDIF entries. The `creationLdif` can also
	// be used to add the user account to an existing group. All LDIF entries are
	// performed in order. If Vault encounters an error while executing the
	// `creationLdif` it will stop at the first error and not execute any remaining
	// LDIF entries. If an error occurs and `rollbackLdif` is specified, the LDIF
	// entries in `rollbackLdif` will be executed. See `rollbackLdif` for more
	// details. This field may optionally be provided as a base64 encoded string.
	CreationLdif pulumi.StringInput
	// Specifies the TTL for the leases associated with this role.
	DefaultTtl pulumi.IntPtrInput
	// A templatized LDIF string used to delete the
	// user account once its TTL has expired. This may contain multiple LDIF
	// entries. All LDIF entries are performed in order. If Vault encounters an
	// error while executing an entry in the `deletionLdif` it will attempt to
	// continue executing any remaining entries. This field may optionally be
	// provided as a base64 encoded string.
	DeletionLdif pulumi.StringInput
	// Specifies the maximum TTL for the leases associated with this role.
	MaxTtl pulumi.IntPtrInput
	// The unique path this backend should be mounted at. Must
	// not begin or end with a `/`. Defaults to `ldap`.
	Mount pulumi.StringPtrInput
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
	// *Available only for Vault Enterprise*.
	Namespace pulumi.StringPtrInput
	// Name of the role.
	RoleName pulumi.StringInput
	// A templatized LDIF string used to attempt to
	// rollback any changes in the event that execution of the `creationLdif` results
	// in an error. This may contain multiple LDIF entries. All LDIF entries are
	// performed in order. If Vault encounters an error while executing an entry in
	// the `rollbackLdif` it will attempt to continue executing any remaining
	// entries. This field may optionally be provided as a base64 encoded string.
	RollbackLdif pulumi.StringPtrInput
	// A template used to generate a dynamic
	// username. This will be used to fill in the `.Username` field within the
	// `creationLdif` string.
	UsernameTemplate pulumi.StringPtrInput
}

func (SecretBackendDynamicRoleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*secretBackendDynamicRoleArgs)(nil)).Elem()
}

type SecretBackendDynamicRoleInput interface {
	pulumi.Input

	ToSecretBackendDynamicRoleOutput() SecretBackendDynamicRoleOutput
	ToSecretBackendDynamicRoleOutputWithContext(ctx context.Context) SecretBackendDynamicRoleOutput
}

func (*SecretBackendDynamicRole) ElementType() reflect.Type {
	return reflect.TypeOf((**SecretBackendDynamicRole)(nil)).Elem()
}

func (i *SecretBackendDynamicRole) ToSecretBackendDynamicRoleOutput() SecretBackendDynamicRoleOutput {
	return i.ToSecretBackendDynamicRoleOutputWithContext(context.Background())
}

func (i *SecretBackendDynamicRole) ToSecretBackendDynamicRoleOutputWithContext(ctx context.Context) SecretBackendDynamicRoleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecretBackendDynamicRoleOutput)
}

// SecretBackendDynamicRoleArrayInput is an input type that accepts SecretBackendDynamicRoleArray and SecretBackendDynamicRoleArrayOutput values.
// You can construct a concrete instance of `SecretBackendDynamicRoleArrayInput` via:
//
//	SecretBackendDynamicRoleArray{ SecretBackendDynamicRoleArgs{...} }
type SecretBackendDynamicRoleArrayInput interface {
	pulumi.Input

	ToSecretBackendDynamicRoleArrayOutput() SecretBackendDynamicRoleArrayOutput
	ToSecretBackendDynamicRoleArrayOutputWithContext(context.Context) SecretBackendDynamicRoleArrayOutput
}

type SecretBackendDynamicRoleArray []SecretBackendDynamicRoleInput

func (SecretBackendDynamicRoleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SecretBackendDynamicRole)(nil)).Elem()
}

func (i SecretBackendDynamicRoleArray) ToSecretBackendDynamicRoleArrayOutput() SecretBackendDynamicRoleArrayOutput {
	return i.ToSecretBackendDynamicRoleArrayOutputWithContext(context.Background())
}

func (i SecretBackendDynamicRoleArray) ToSecretBackendDynamicRoleArrayOutputWithContext(ctx context.Context) SecretBackendDynamicRoleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecretBackendDynamicRoleArrayOutput)
}

// SecretBackendDynamicRoleMapInput is an input type that accepts SecretBackendDynamicRoleMap and SecretBackendDynamicRoleMapOutput values.
// You can construct a concrete instance of `SecretBackendDynamicRoleMapInput` via:
//
//	SecretBackendDynamicRoleMap{ "key": SecretBackendDynamicRoleArgs{...} }
type SecretBackendDynamicRoleMapInput interface {
	pulumi.Input

	ToSecretBackendDynamicRoleMapOutput() SecretBackendDynamicRoleMapOutput
	ToSecretBackendDynamicRoleMapOutputWithContext(context.Context) SecretBackendDynamicRoleMapOutput
}

type SecretBackendDynamicRoleMap map[string]SecretBackendDynamicRoleInput

func (SecretBackendDynamicRoleMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SecretBackendDynamicRole)(nil)).Elem()
}

func (i SecretBackendDynamicRoleMap) ToSecretBackendDynamicRoleMapOutput() SecretBackendDynamicRoleMapOutput {
	return i.ToSecretBackendDynamicRoleMapOutputWithContext(context.Background())
}

func (i SecretBackendDynamicRoleMap) ToSecretBackendDynamicRoleMapOutputWithContext(ctx context.Context) SecretBackendDynamicRoleMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecretBackendDynamicRoleMapOutput)
}

type SecretBackendDynamicRoleOutput struct{ *pulumi.OutputState }

func (SecretBackendDynamicRoleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SecretBackendDynamicRole)(nil)).Elem()
}

func (o SecretBackendDynamicRoleOutput) ToSecretBackendDynamicRoleOutput() SecretBackendDynamicRoleOutput {
	return o
}

func (o SecretBackendDynamicRoleOutput) ToSecretBackendDynamicRoleOutputWithContext(ctx context.Context) SecretBackendDynamicRoleOutput {
	return o
}

// A templatized LDIF string used to create a user
// account. This may contain multiple LDIF entries. The `creationLdif` can also
// be used to add the user account to an existing group. All LDIF entries are
// performed in order. If Vault encounters an error while executing the
// `creationLdif` it will stop at the first error and not execute any remaining
// LDIF entries. If an error occurs and `rollbackLdif` is specified, the LDIF
// entries in `rollbackLdif` will be executed. See `rollbackLdif` for more
// details. This field may optionally be provided as a base64 encoded string.
func (o SecretBackendDynamicRoleOutput) CreationLdif() pulumi.StringOutput {
	return o.ApplyT(func(v *SecretBackendDynamicRole) pulumi.StringOutput { return v.CreationLdif }).(pulumi.StringOutput)
}

// Specifies the TTL for the leases associated with this role.
func (o SecretBackendDynamicRoleOutput) DefaultTtl() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SecretBackendDynamicRole) pulumi.IntPtrOutput { return v.DefaultTtl }).(pulumi.IntPtrOutput)
}

// A templatized LDIF string used to delete the
// user account once its TTL has expired. This may contain multiple LDIF
// entries. All LDIF entries are performed in order. If Vault encounters an
// error while executing an entry in the `deletionLdif` it will attempt to
// continue executing any remaining entries. This field may optionally be
// provided as a base64 encoded string.
func (o SecretBackendDynamicRoleOutput) DeletionLdif() pulumi.StringOutput {
	return o.ApplyT(func(v *SecretBackendDynamicRole) pulumi.StringOutput { return v.DeletionLdif }).(pulumi.StringOutput)
}

// Specifies the maximum TTL for the leases associated with this role.
func (o SecretBackendDynamicRoleOutput) MaxTtl() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SecretBackendDynamicRole) pulumi.IntPtrOutput { return v.MaxTtl }).(pulumi.IntPtrOutput)
}

// The unique path this backend should be mounted at. Must
// not begin or end with a `/`. Defaults to `ldap`.
func (o SecretBackendDynamicRoleOutput) Mount() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecretBackendDynamicRole) pulumi.StringPtrOutput { return v.Mount }).(pulumi.StringPtrOutput)
}

// The namespace to provision the resource in.
// The value should not contain leading or trailing forward slashes.
// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
// *Available only for Vault Enterprise*.
func (o SecretBackendDynamicRoleOutput) Namespace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecretBackendDynamicRole) pulumi.StringPtrOutput { return v.Namespace }).(pulumi.StringPtrOutput)
}

// Name of the role.
func (o SecretBackendDynamicRoleOutput) RoleName() pulumi.StringOutput {
	return o.ApplyT(func(v *SecretBackendDynamicRole) pulumi.StringOutput { return v.RoleName }).(pulumi.StringOutput)
}

// A templatized LDIF string used to attempt to
// rollback any changes in the event that execution of the `creationLdif` results
// in an error. This may contain multiple LDIF entries. All LDIF entries are
// performed in order. If Vault encounters an error while executing an entry in
// the `rollbackLdif` it will attempt to continue executing any remaining
// entries. This field may optionally be provided as a base64 encoded string.
func (o SecretBackendDynamicRoleOutput) RollbackLdif() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecretBackendDynamicRole) pulumi.StringPtrOutput { return v.RollbackLdif }).(pulumi.StringPtrOutput)
}

// A template used to generate a dynamic
// username. This will be used to fill in the `.Username` field within the
// `creationLdif` string.
func (o SecretBackendDynamicRoleOutput) UsernameTemplate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecretBackendDynamicRole) pulumi.StringPtrOutput { return v.UsernameTemplate }).(pulumi.StringPtrOutput)
}

type SecretBackendDynamicRoleArrayOutput struct{ *pulumi.OutputState }

func (SecretBackendDynamicRoleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SecretBackendDynamicRole)(nil)).Elem()
}

func (o SecretBackendDynamicRoleArrayOutput) ToSecretBackendDynamicRoleArrayOutput() SecretBackendDynamicRoleArrayOutput {
	return o
}

func (o SecretBackendDynamicRoleArrayOutput) ToSecretBackendDynamicRoleArrayOutputWithContext(ctx context.Context) SecretBackendDynamicRoleArrayOutput {
	return o
}

func (o SecretBackendDynamicRoleArrayOutput) Index(i pulumi.IntInput) SecretBackendDynamicRoleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SecretBackendDynamicRole {
		return vs[0].([]*SecretBackendDynamicRole)[vs[1].(int)]
	}).(SecretBackendDynamicRoleOutput)
}

type SecretBackendDynamicRoleMapOutput struct{ *pulumi.OutputState }

func (SecretBackendDynamicRoleMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SecretBackendDynamicRole)(nil)).Elem()
}

func (o SecretBackendDynamicRoleMapOutput) ToSecretBackendDynamicRoleMapOutput() SecretBackendDynamicRoleMapOutput {
	return o
}

func (o SecretBackendDynamicRoleMapOutput) ToSecretBackendDynamicRoleMapOutputWithContext(ctx context.Context) SecretBackendDynamicRoleMapOutput {
	return o
}

func (o SecretBackendDynamicRoleMapOutput) MapIndex(k pulumi.StringInput) SecretBackendDynamicRoleOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SecretBackendDynamicRole {
		return vs[0].(map[string]*SecretBackendDynamicRole)[vs[1].(string)]
	}).(SecretBackendDynamicRoleOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SecretBackendDynamicRoleInput)(nil)).Elem(), &SecretBackendDynamicRole{})
	pulumi.RegisterInputType(reflect.TypeOf((*SecretBackendDynamicRoleArrayInput)(nil)).Elem(), SecretBackendDynamicRoleArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SecretBackendDynamicRoleMapInput)(nil)).Elem(), SecretBackendDynamicRoleMap{})
	pulumi.RegisterOutputType(SecretBackendDynamicRoleOutput{})
	pulumi.RegisterOutputType(SecretBackendDynamicRoleArrayOutput{})
	pulumi.RegisterOutputType(SecretBackendDynamicRoleMapOutput{})
}
