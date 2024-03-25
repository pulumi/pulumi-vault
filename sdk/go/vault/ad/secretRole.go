// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ad

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
//	"github.com/pulumi/pulumi-vault/sdk/v6/go/vault/ad"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			config, err := ad.NewSecretBackend(ctx, "config", &ad.SecretBackendArgs{
//				Backend:     pulumi.String("ad"),
//				Binddn:      pulumi.String("CN=Administrator,CN=Users,DC=corp,DC=example,DC=net"),
//				Bindpass:    pulumi.String("SuperSecretPassw0rd"),
//				Url:         pulumi.String("ldaps://ad"),
//				InsecureTls: pulumi.Bool(true),
//				Userdn:      pulumi.String("CN=Users,DC=corp,DC=example,DC=net"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = ad.NewSecretRole(ctx, "role", &ad.SecretRoleArgs{
//				Backend:            config.Backend,
//				Role:               pulumi.String("bob"),
//				ServiceAccountName: pulumi.String("Bob"),
//				Ttl:                pulumi.Int(60),
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
// AD secret backend roles can be imported using the `path`, e.g.
//
// ```sh
// $ pulumi import vault:ad/secretRole:SecretRole role ad/roles/bob
// ```
type SecretRole struct {
	pulumi.CustomResourceState

	// The path the AD secret backend is mounted at,
	// with no leading or trailing `/`s.
	Backend pulumi.StringOutput `pulumi:"backend"`
	// Timestamp of the last password rotation by Vault.
	LastVaultRotation pulumi.StringOutput `pulumi:"lastVaultRotation"`
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
	// *Available only for Vault Enterprise*.
	Namespace pulumi.StringPtrOutput `pulumi:"namespace"`
	// Timestamp of the last password set by Vault.
	PasswordLastSet pulumi.StringOutput `pulumi:"passwordLastSet"`
	// The name to identify this role within the backend.
	// Must be unique within the backend.
	Role pulumi.StringOutput `pulumi:"role"`
	// Specifies the name of the Active Directory service
	// account mapped to this role.
	ServiceAccountName pulumi.StringOutput `pulumi:"serviceAccountName"`
	// The password time-to-live in seconds. Defaults to the configuration
	// ttl if not provided.
	Ttl pulumi.IntPtrOutput `pulumi:"ttl"`
}

// NewSecretRole registers a new resource with the given unique name, arguments, and options.
func NewSecretRole(ctx *pulumi.Context,
	name string, args *SecretRoleArgs, opts ...pulumi.ResourceOption) (*SecretRole, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Backend == nil {
		return nil, errors.New("invalid value for required argument 'Backend'")
	}
	if args.Role == nil {
		return nil, errors.New("invalid value for required argument 'Role'")
	}
	if args.ServiceAccountName == nil {
		return nil, errors.New("invalid value for required argument 'ServiceAccountName'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource SecretRole
	err := ctx.RegisterResource("vault:ad/secretRole:SecretRole", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSecretRole gets an existing SecretRole resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSecretRole(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SecretRoleState, opts ...pulumi.ResourceOption) (*SecretRole, error) {
	var resource SecretRole
	err := ctx.ReadResource("vault:ad/secretRole:SecretRole", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SecretRole resources.
type secretRoleState struct {
	// The path the AD secret backend is mounted at,
	// with no leading or trailing `/`s.
	Backend *string `pulumi:"backend"`
	// Timestamp of the last password rotation by Vault.
	LastVaultRotation *string `pulumi:"lastVaultRotation"`
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
	// *Available only for Vault Enterprise*.
	Namespace *string `pulumi:"namespace"`
	// Timestamp of the last password set by Vault.
	PasswordLastSet *string `pulumi:"passwordLastSet"`
	// The name to identify this role within the backend.
	// Must be unique within the backend.
	Role *string `pulumi:"role"`
	// Specifies the name of the Active Directory service
	// account mapped to this role.
	ServiceAccountName *string `pulumi:"serviceAccountName"`
	// The password time-to-live in seconds. Defaults to the configuration
	// ttl if not provided.
	Ttl *int `pulumi:"ttl"`
}

type SecretRoleState struct {
	// The path the AD secret backend is mounted at,
	// with no leading or trailing `/`s.
	Backend pulumi.StringPtrInput
	// Timestamp of the last password rotation by Vault.
	LastVaultRotation pulumi.StringPtrInput
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
	// *Available only for Vault Enterprise*.
	Namespace pulumi.StringPtrInput
	// Timestamp of the last password set by Vault.
	PasswordLastSet pulumi.StringPtrInput
	// The name to identify this role within the backend.
	// Must be unique within the backend.
	Role pulumi.StringPtrInput
	// Specifies the name of the Active Directory service
	// account mapped to this role.
	ServiceAccountName pulumi.StringPtrInput
	// The password time-to-live in seconds. Defaults to the configuration
	// ttl if not provided.
	Ttl pulumi.IntPtrInput
}

func (SecretRoleState) ElementType() reflect.Type {
	return reflect.TypeOf((*secretRoleState)(nil)).Elem()
}

type secretRoleArgs struct {
	// The path the AD secret backend is mounted at,
	// with no leading or trailing `/`s.
	Backend string `pulumi:"backend"`
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
	// *Available only for Vault Enterprise*.
	Namespace *string `pulumi:"namespace"`
	// The name to identify this role within the backend.
	// Must be unique within the backend.
	Role string `pulumi:"role"`
	// Specifies the name of the Active Directory service
	// account mapped to this role.
	ServiceAccountName string `pulumi:"serviceAccountName"`
	// The password time-to-live in seconds. Defaults to the configuration
	// ttl if not provided.
	Ttl *int `pulumi:"ttl"`
}

// The set of arguments for constructing a SecretRole resource.
type SecretRoleArgs struct {
	// The path the AD secret backend is mounted at,
	// with no leading or trailing `/`s.
	Backend pulumi.StringInput
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
	// *Available only for Vault Enterprise*.
	Namespace pulumi.StringPtrInput
	// The name to identify this role within the backend.
	// Must be unique within the backend.
	Role pulumi.StringInput
	// Specifies the name of the Active Directory service
	// account mapped to this role.
	ServiceAccountName pulumi.StringInput
	// The password time-to-live in seconds. Defaults to the configuration
	// ttl if not provided.
	Ttl pulumi.IntPtrInput
}

func (SecretRoleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*secretRoleArgs)(nil)).Elem()
}

type SecretRoleInput interface {
	pulumi.Input

	ToSecretRoleOutput() SecretRoleOutput
	ToSecretRoleOutputWithContext(ctx context.Context) SecretRoleOutput
}

func (*SecretRole) ElementType() reflect.Type {
	return reflect.TypeOf((**SecretRole)(nil)).Elem()
}

func (i *SecretRole) ToSecretRoleOutput() SecretRoleOutput {
	return i.ToSecretRoleOutputWithContext(context.Background())
}

func (i *SecretRole) ToSecretRoleOutputWithContext(ctx context.Context) SecretRoleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecretRoleOutput)
}

// SecretRoleArrayInput is an input type that accepts SecretRoleArray and SecretRoleArrayOutput values.
// You can construct a concrete instance of `SecretRoleArrayInput` via:
//
//	SecretRoleArray{ SecretRoleArgs{...} }
type SecretRoleArrayInput interface {
	pulumi.Input

	ToSecretRoleArrayOutput() SecretRoleArrayOutput
	ToSecretRoleArrayOutputWithContext(context.Context) SecretRoleArrayOutput
}

type SecretRoleArray []SecretRoleInput

func (SecretRoleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SecretRole)(nil)).Elem()
}

func (i SecretRoleArray) ToSecretRoleArrayOutput() SecretRoleArrayOutput {
	return i.ToSecretRoleArrayOutputWithContext(context.Background())
}

func (i SecretRoleArray) ToSecretRoleArrayOutputWithContext(ctx context.Context) SecretRoleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecretRoleArrayOutput)
}

// SecretRoleMapInput is an input type that accepts SecretRoleMap and SecretRoleMapOutput values.
// You can construct a concrete instance of `SecretRoleMapInput` via:
//
//	SecretRoleMap{ "key": SecretRoleArgs{...} }
type SecretRoleMapInput interface {
	pulumi.Input

	ToSecretRoleMapOutput() SecretRoleMapOutput
	ToSecretRoleMapOutputWithContext(context.Context) SecretRoleMapOutput
}

type SecretRoleMap map[string]SecretRoleInput

func (SecretRoleMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SecretRole)(nil)).Elem()
}

func (i SecretRoleMap) ToSecretRoleMapOutput() SecretRoleMapOutput {
	return i.ToSecretRoleMapOutputWithContext(context.Background())
}

func (i SecretRoleMap) ToSecretRoleMapOutputWithContext(ctx context.Context) SecretRoleMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecretRoleMapOutput)
}

type SecretRoleOutput struct{ *pulumi.OutputState }

func (SecretRoleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SecretRole)(nil)).Elem()
}

func (o SecretRoleOutput) ToSecretRoleOutput() SecretRoleOutput {
	return o
}

func (o SecretRoleOutput) ToSecretRoleOutputWithContext(ctx context.Context) SecretRoleOutput {
	return o
}

// The path the AD secret backend is mounted at,
// with no leading or trailing `/`s.
func (o SecretRoleOutput) Backend() pulumi.StringOutput {
	return o.ApplyT(func(v *SecretRole) pulumi.StringOutput { return v.Backend }).(pulumi.StringOutput)
}

// Timestamp of the last password rotation by Vault.
func (o SecretRoleOutput) LastVaultRotation() pulumi.StringOutput {
	return o.ApplyT(func(v *SecretRole) pulumi.StringOutput { return v.LastVaultRotation }).(pulumi.StringOutput)
}

// The namespace to provision the resource in.
// The value should not contain leading or trailing forward slashes.
// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
// *Available only for Vault Enterprise*.
func (o SecretRoleOutput) Namespace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecretRole) pulumi.StringPtrOutput { return v.Namespace }).(pulumi.StringPtrOutput)
}

// Timestamp of the last password set by Vault.
func (o SecretRoleOutput) PasswordLastSet() pulumi.StringOutput {
	return o.ApplyT(func(v *SecretRole) pulumi.StringOutput { return v.PasswordLastSet }).(pulumi.StringOutput)
}

// The name to identify this role within the backend.
// Must be unique within the backend.
func (o SecretRoleOutput) Role() pulumi.StringOutput {
	return o.ApplyT(func(v *SecretRole) pulumi.StringOutput { return v.Role }).(pulumi.StringOutput)
}

// Specifies the name of the Active Directory service
// account mapped to this role.
func (o SecretRoleOutput) ServiceAccountName() pulumi.StringOutput {
	return o.ApplyT(func(v *SecretRole) pulumi.StringOutput { return v.ServiceAccountName }).(pulumi.StringOutput)
}

// The password time-to-live in seconds. Defaults to the configuration
// ttl if not provided.
func (o SecretRoleOutput) Ttl() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SecretRole) pulumi.IntPtrOutput { return v.Ttl }).(pulumi.IntPtrOutput)
}

type SecretRoleArrayOutput struct{ *pulumi.OutputState }

func (SecretRoleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SecretRole)(nil)).Elem()
}

func (o SecretRoleArrayOutput) ToSecretRoleArrayOutput() SecretRoleArrayOutput {
	return o
}

func (o SecretRoleArrayOutput) ToSecretRoleArrayOutputWithContext(ctx context.Context) SecretRoleArrayOutput {
	return o
}

func (o SecretRoleArrayOutput) Index(i pulumi.IntInput) SecretRoleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SecretRole {
		return vs[0].([]*SecretRole)[vs[1].(int)]
	}).(SecretRoleOutput)
}

type SecretRoleMapOutput struct{ *pulumi.OutputState }

func (SecretRoleMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SecretRole)(nil)).Elem()
}

func (o SecretRoleMapOutput) ToSecretRoleMapOutput() SecretRoleMapOutput {
	return o
}

func (o SecretRoleMapOutput) ToSecretRoleMapOutputWithContext(ctx context.Context) SecretRoleMapOutput {
	return o
}

func (o SecretRoleMapOutput) MapIndex(k pulumi.StringInput) SecretRoleOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SecretRole {
		return vs[0].(map[string]*SecretRole)[vs[1].(string)]
	}).(SecretRoleOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SecretRoleInput)(nil)).Elem(), &SecretRole{})
	pulumi.RegisterInputType(reflect.TypeOf((*SecretRoleArrayInput)(nil)).Elem(), SecretRoleArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SecretRoleMapInput)(nil)).Elem(), SecretRoleMap{})
	pulumi.RegisterOutputType(SecretRoleOutput{})
	pulumi.RegisterOutputType(SecretRoleArrayOutput{})
	pulumi.RegisterOutputType(SecretRoleMapOutput{})
}
