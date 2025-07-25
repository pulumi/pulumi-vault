// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package kmip

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-vault/sdk/v7/go/vault/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages KMIP Secret roles in a Vault server. This feature requires
// Vault Enterprise. See the [Vault documentation](https://www.vaultproject.io/docs/secrets/kmip)
// for more information.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-vault/sdk/v7/go/vault/kmip"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_default, err := kmip.NewSecretBackend(ctx, "default", &kmip.SecretBackendArgs{
//				Path:        pulumi.String("kmip"),
//				Description: pulumi.String("Vault KMIP backend"),
//			})
//			if err != nil {
//				return err
//			}
//			dev, err := kmip.NewSecretScope(ctx, "dev", &kmip.SecretScopeArgs{
//				Path:  _default.Path,
//				Scope: pulumi.String("dev"),
//				Force: pulumi.Bool(true),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = kmip.NewSecretRole(ctx, "admin", &kmip.SecretRoleArgs{
//				Path:                   dev.Path,
//				Scope:                  dev.Scope,
//				Role:                   pulumi.String("admin"),
//				TlsClientKeyType:       pulumi.String("ec"),
//				TlsClientKeyBits:       pulumi.Int(256),
//				OperationActivate:      pulumi.Bool(true),
//				OperationGet:           pulumi.Bool(true),
//				OperationGetAttributes: pulumi.Bool(true),
//				OperationCreate:        pulumi.Bool(true),
//				OperationDestroy:       pulumi.Bool(true),
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
// KMIP Secret role can be imported using the `path`, e.g.
//
// ```sh
// $ pulumi import vault:kmip/secretRole:SecretRole admin kmip
// ```
type SecretRole struct {
	pulumi.CustomResourceState

	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
	// *Available only for Vault Enterprise*.
	Namespace pulumi.StringPtrOutput `pulumi:"namespace"`
	// Grant permission to use the KMIP Activate operation.
	OperationActivate pulumi.BoolOutput `pulumi:"operationActivate"`
	// Grant permission to use the KMIP Add Attribute operation.
	OperationAddAttribute pulumi.BoolOutput `pulumi:"operationAddAttribute"`
	// Grant all permissions to this role. May not be specified with any other `operation_*` params.
	OperationAll pulumi.BoolOutput `pulumi:"operationAll"`
	// Grant permission to use the KMIP Create operation.
	OperationCreate pulumi.BoolOutput `pulumi:"operationCreate"`
	// Grant permission to use the KMIP Destroy operation.
	OperationDestroy pulumi.BoolOutput `pulumi:"operationDestroy"`
	// Grant permission to use the KMIP Discover Version operation.
	OperationDiscoverVersions pulumi.BoolOutput `pulumi:"operationDiscoverVersions"`
	// Grant permission to use the KMIP Get operation.
	OperationGet pulumi.BoolOutput `pulumi:"operationGet"`
	// Grant permission to use the KMIP Get Atrribute List operation.
	OperationGetAttributeList pulumi.BoolOutput `pulumi:"operationGetAttributeList"`
	// Grant permission to use the KMIP Get Atrributes operation.
	OperationGetAttributes pulumi.BoolOutput `pulumi:"operationGetAttributes"`
	// Grant permission to use the KMIP Get Locate operation.
	OperationLocate pulumi.BoolOutput `pulumi:"operationLocate"`
	// Remove all permissions from this role. May not be specified with any other `operation_*` params.
	OperationNone pulumi.BoolOutput `pulumi:"operationNone"`
	// Grant permission to use the KMIP Register operation.
	OperationRegister pulumi.BoolOutput `pulumi:"operationRegister"`
	// Grant permission to use the KMIP Rekey operation.
	OperationRekey pulumi.BoolOutput `pulumi:"operationRekey"`
	// Grant permission to use the KMIP Revoke operation.
	OperationRevoke pulumi.BoolOutput `pulumi:"operationRevoke"`
	// The unique path this backend should be mounted at. Must
	// not begin or end with a `/`. Defaults to `kmip`.
	Path pulumi.StringOutput `pulumi:"path"`
	// Name of the role.
	Role pulumi.StringOutput `pulumi:"role"`
	// Name of the scope.
	Scope pulumi.StringOutput `pulumi:"scope"`
	// Client certificate key bits, valid values depend on key type.
	TlsClientKeyBits pulumi.IntPtrOutput `pulumi:"tlsClientKeyBits"`
	// Client certificate key type, `rsa` or `ec`.
	TlsClientKeyType pulumi.StringPtrOutput `pulumi:"tlsClientKeyType"`
	// Client certificate TTL in seconds.
	TlsClientTtl pulumi.IntPtrOutput `pulumi:"tlsClientTtl"`
}

// NewSecretRole registers a new resource with the given unique name, arguments, and options.
func NewSecretRole(ctx *pulumi.Context,
	name string, args *SecretRoleArgs, opts ...pulumi.ResourceOption) (*SecretRole, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Path == nil {
		return nil, errors.New("invalid value for required argument 'Path'")
	}
	if args.Role == nil {
		return nil, errors.New("invalid value for required argument 'Role'")
	}
	if args.Scope == nil {
		return nil, errors.New("invalid value for required argument 'Scope'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource SecretRole
	err := ctx.RegisterResource("vault:kmip/secretRole:SecretRole", name, args, &resource, opts...)
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
	err := ctx.ReadResource("vault:kmip/secretRole:SecretRole", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SecretRole resources.
type secretRoleState struct {
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
	// *Available only for Vault Enterprise*.
	Namespace *string `pulumi:"namespace"`
	// Grant permission to use the KMIP Activate operation.
	OperationActivate *bool `pulumi:"operationActivate"`
	// Grant permission to use the KMIP Add Attribute operation.
	OperationAddAttribute *bool `pulumi:"operationAddAttribute"`
	// Grant all permissions to this role. May not be specified with any other `operation_*` params.
	OperationAll *bool `pulumi:"operationAll"`
	// Grant permission to use the KMIP Create operation.
	OperationCreate *bool `pulumi:"operationCreate"`
	// Grant permission to use the KMIP Destroy operation.
	OperationDestroy *bool `pulumi:"operationDestroy"`
	// Grant permission to use the KMIP Discover Version operation.
	OperationDiscoverVersions *bool `pulumi:"operationDiscoverVersions"`
	// Grant permission to use the KMIP Get operation.
	OperationGet *bool `pulumi:"operationGet"`
	// Grant permission to use the KMIP Get Atrribute List operation.
	OperationGetAttributeList *bool `pulumi:"operationGetAttributeList"`
	// Grant permission to use the KMIP Get Atrributes operation.
	OperationGetAttributes *bool `pulumi:"operationGetAttributes"`
	// Grant permission to use the KMIP Get Locate operation.
	OperationLocate *bool `pulumi:"operationLocate"`
	// Remove all permissions from this role. May not be specified with any other `operation_*` params.
	OperationNone *bool `pulumi:"operationNone"`
	// Grant permission to use the KMIP Register operation.
	OperationRegister *bool `pulumi:"operationRegister"`
	// Grant permission to use the KMIP Rekey operation.
	OperationRekey *bool `pulumi:"operationRekey"`
	// Grant permission to use the KMIP Revoke operation.
	OperationRevoke *bool `pulumi:"operationRevoke"`
	// The unique path this backend should be mounted at. Must
	// not begin or end with a `/`. Defaults to `kmip`.
	Path *string `pulumi:"path"`
	// Name of the role.
	Role *string `pulumi:"role"`
	// Name of the scope.
	Scope *string `pulumi:"scope"`
	// Client certificate key bits, valid values depend on key type.
	TlsClientKeyBits *int `pulumi:"tlsClientKeyBits"`
	// Client certificate key type, `rsa` or `ec`.
	TlsClientKeyType *string `pulumi:"tlsClientKeyType"`
	// Client certificate TTL in seconds.
	TlsClientTtl *int `pulumi:"tlsClientTtl"`
}

type SecretRoleState struct {
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
	// *Available only for Vault Enterprise*.
	Namespace pulumi.StringPtrInput
	// Grant permission to use the KMIP Activate operation.
	OperationActivate pulumi.BoolPtrInput
	// Grant permission to use the KMIP Add Attribute operation.
	OperationAddAttribute pulumi.BoolPtrInput
	// Grant all permissions to this role. May not be specified with any other `operation_*` params.
	OperationAll pulumi.BoolPtrInput
	// Grant permission to use the KMIP Create operation.
	OperationCreate pulumi.BoolPtrInput
	// Grant permission to use the KMIP Destroy operation.
	OperationDestroy pulumi.BoolPtrInput
	// Grant permission to use the KMIP Discover Version operation.
	OperationDiscoverVersions pulumi.BoolPtrInput
	// Grant permission to use the KMIP Get operation.
	OperationGet pulumi.BoolPtrInput
	// Grant permission to use the KMIP Get Atrribute List operation.
	OperationGetAttributeList pulumi.BoolPtrInput
	// Grant permission to use the KMIP Get Atrributes operation.
	OperationGetAttributes pulumi.BoolPtrInput
	// Grant permission to use the KMIP Get Locate operation.
	OperationLocate pulumi.BoolPtrInput
	// Remove all permissions from this role. May not be specified with any other `operation_*` params.
	OperationNone pulumi.BoolPtrInput
	// Grant permission to use the KMIP Register operation.
	OperationRegister pulumi.BoolPtrInput
	// Grant permission to use the KMIP Rekey operation.
	OperationRekey pulumi.BoolPtrInput
	// Grant permission to use the KMIP Revoke operation.
	OperationRevoke pulumi.BoolPtrInput
	// The unique path this backend should be mounted at. Must
	// not begin or end with a `/`. Defaults to `kmip`.
	Path pulumi.StringPtrInput
	// Name of the role.
	Role pulumi.StringPtrInput
	// Name of the scope.
	Scope pulumi.StringPtrInput
	// Client certificate key bits, valid values depend on key type.
	TlsClientKeyBits pulumi.IntPtrInput
	// Client certificate key type, `rsa` or `ec`.
	TlsClientKeyType pulumi.StringPtrInput
	// Client certificate TTL in seconds.
	TlsClientTtl pulumi.IntPtrInput
}

func (SecretRoleState) ElementType() reflect.Type {
	return reflect.TypeOf((*secretRoleState)(nil)).Elem()
}

type secretRoleArgs struct {
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
	// *Available only for Vault Enterprise*.
	Namespace *string `pulumi:"namespace"`
	// Grant permission to use the KMIP Activate operation.
	OperationActivate *bool `pulumi:"operationActivate"`
	// Grant permission to use the KMIP Add Attribute operation.
	OperationAddAttribute *bool `pulumi:"operationAddAttribute"`
	// Grant all permissions to this role. May not be specified with any other `operation_*` params.
	OperationAll *bool `pulumi:"operationAll"`
	// Grant permission to use the KMIP Create operation.
	OperationCreate *bool `pulumi:"operationCreate"`
	// Grant permission to use the KMIP Destroy operation.
	OperationDestroy *bool `pulumi:"operationDestroy"`
	// Grant permission to use the KMIP Discover Version operation.
	OperationDiscoverVersions *bool `pulumi:"operationDiscoverVersions"`
	// Grant permission to use the KMIP Get operation.
	OperationGet *bool `pulumi:"operationGet"`
	// Grant permission to use the KMIP Get Atrribute List operation.
	OperationGetAttributeList *bool `pulumi:"operationGetAttributeList"`
	// Grant permission to use the KMIP Get Atrributes operation.
	OperationGetAttributes *bool `pulumi:"operationGetAttributes"`
	// Grant permission to use the KMIP Get Locate operation.
	OperationLocate *bool `pulumi:"operationLocate"`
	// Remove all permissions from this role. May not be specified with any other `operation_*` params.
	OperationNone *bool `pulumi:"operationNone"`
	// Grant permission to use the KMIP Register operation.
	OperationRegister *bool `pulumi:"operationRegister"`
	// Grant permission to use the KMIP Rekey operation.
	OperationRekey *bool `pulumi:"operationRekey"`
	// Grant permission to use the KMIP Revoke operation.
	OperationRevoke *bool `pulumi:"operationRevoke"`
	// The unique path this backend should be mounted at. Must
	// not begin or end with a `/`. Defaults to `kmip`.
	Path string `pulumi:"path"`
	// Name of the role.
	Role string `pulumi:"role"`
	// Name of the scope.
	Scope string `pulumi:"scope"`
	// Client certificate key bits, valid values depend on key type.
	TlsClientKeyBits *int `pulumi:"tlsClientKeyBits"`
	// Client certificate key type, `rsa` or `ec`.
	TlsClientKeyType *string `pulumi:"tlsClientKeyType"`
	// Client certificate TTL in seconds.
	TlsClientTtl *int `pulumi:"tlsClientTtl"`
}

// The set of arguments for constructing a SecretRole resource.
type SecretRoleArgs struct {
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
	// *Available only for Vault Enterprise*.
	Namespace pulumi.StringPtrInput
	// Grant permission to use the KMIP Activate operation.
	OperationActivate pulumi.BoolPtrInput
	// Grant permission to use the KMIP Add Attribute operation.
	OperationAddAttribute pulumi.BoolPtrInput
	// Grant all permissions to this role. May not be specified with any other `operation_*` params.
	OperationAll pulumi.BoolPtrInput
	// Grant permission to use the KMIP Create operation.
	OperationCreate pulumi.BoolPtrInput
	// Grant permission to use the KMIP Destroy operation.
	OperationDestroy pulumi.BoolPtrInput
	// Grant permission to use the KMIP Discover Version operation.
	OperationDiscoverVersions pulumi.BoolPtrInput
	// Grant permission to use the KMIP Get operation.
	OperationGet pulumi.BoolPtrInput
	// Grant permission to use the KMIP Get Atrribute List operation.
	OperationGetAttributeList pulumi.BoolPtrInput
	// Grant permission to use the KMIP Get Atrributes operation.
	OperationGetAttributes pulumi.BoolPtrInput
	// Grant permission to use the KMIP Get Locate operation.
	OperationLocate pulumi.BoolPtrInput
	// Remove all permissions from this role. May not be specified with any other `operation_*` params.
	OperationNone pulumi.BoolPtrInput
	// Grant permission to use the KMIP Register operation.
	OperationRegister pulumi.BoolPtrInput
	// Grant permission to use the KMIP Rekey operation.
	OperationRekey pulumi.BoolPtrInput
	// Grant permission to use the KMIP Revoke operation.
	OperationRevoke pulumi.BoolPtrInput
	// The unique path this backend should be mounted at. Must
	// not begin or end with a `/`. Defaults to `kmip`.
	Path pulumi.StringInput
	// Name of the role.
	Role pulumi.StringInput
	// Name of the scope.
	Scope pulumi.StringInput
	// Client certificate key bits, valid values depend on key type.
	TlsClientKeyBits pulumi.IntPtrInput
	// Client certificate key type, `rsa` or `ec`.
	TlsClientKeyType pulumi.StringPtrInput
	// Client certificate TTL in seconds.
	TlsClientTtl pulumi.IntPtrInput
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

// The namespace to provision the resource in.
// The value should not contain leading or trailing forward slashes.
// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
// *Available only for Vault Enterprise*.
func (o SecretRoleOutput) Namespace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecretRole) pulumi.StringPtrOutput { return v.Namespace }).(pulumi.StringPtrOutput)
}

// Grant permission to use the KMIP Activate operation.
func (o SecretRoleOutput) OperationActivate() pulumi.BoolOutput {
	return o.ApplyT(func(v *SecretRole) pulumi.BoolOutput { return v.OperationActivate }).(pulumi.BoolOutput)
}

// Grant permission to use the KMIP Add Attribute operation.
func (o SecretRoleOutput) OperationAddAttribute() pulumi.BoolOutput {
	return o.ApplyT(func(v *SecretRole) pulumi.BoolOutput { return v.OperationAddAttribute }).(pulumi.BoolOutput)
}

// Grant all permissions to this role. May not be specified with any other `operation_*` params.
func (o SecretRoleOutput) OperationAll() pulumi.BoolOutput {
	return o.ApplyT(func(v *SecretRole) pulumi.BoolOutput { return v.OperationAll }).(pulumi.BoolOutput)
}

// Grant permission to use the KMIP Create operation.
func (o SecretRoleOutput) OperationCreate() pulumi.BoolOutput {
	return o.ApplyT(func(v *SecretRole) pulumi.BoolOutput { return v.OperationCreate }).(pulumi.BoolOutput)
}

// Grant permission to use the KMIP Destroy operation.
func (o SecretRoleOutput) OperationDestroy() pulumi.BoolOutput {
	return o.ApplyT(func(v *SecretRole) pulumi.BoolOutput { return v.OperationDestroy }).(pulumi.BoolOutput)
}

// Grant permission to use the KMIP Discover Version operation.
func (o SecretRoleOutput) OperationDiscoverVersions() pulumi.BoolOutput {
	return o.ApplyT(func(v *SecretRole) pulumi.BoolOutput { return v.OperationDiscoverVersions }).(pulumi.BoolOutput)
}

// Grant permission to use the KMIP Get operation.
func (o SecretRoleOutput) OperationGet() pulumi.BoolOutput {
	return o.ApplyT(func(v *SecretRole) pulumi.BoolOutput { return v.OperationGet }).(pulumi.BoolOutput)
}

// Grant permission to use the KMIP Get Atrribute List operation.
func (o SecretRoleOutput) OperationGetAttributeList() pulumi.BoolOutput {
	return o.ApplyT(func(v *SecretRole) pulumi.BoolOutput { return v.OperationGetAttributeList }).(pulumi.BoolOutput)
}

// Grant permission to use the KMIP Get Atrributes operation.
func (o SecretRoleOutput) OperationGetAttributes() pulumi.BoolOutput {
	return o.ApplyT(func(v *SecretRole) pulumi.BoolOutput { return v.OperationGetAttributes }).(pulumi.BoolOutput)
}

// Grant permission to use the KMIP Get Locate operation.
func (o SecretRoleOutput) OperationLocate() pulumi.BoolOutput {
	return o.ApplyT(func(v *SecretRole) pulumi.BoolOutput { return v.OperationLocate }).(pulumi.BoolOutput)
}

// Remove all permissions from this role. May not be specified with any other `operation_*` params.
func (o SecretRoleOutput) OperationNone() pulumi.BoolOutput {
	return o.ApplyT(func(v *SecretRole) pulumi.BoolOutput { return v.OperationNone }).(pulumi.BoolOutput)
}

// Grant permission to use the KMIP Register operation.
func (o SecretRoleOutput) OperationRegister() pulumi.BoolOutput {
	return o.ApplyT(func(v *SecretRole) pulumi.BoolOutput { return v.OperationRegister }).(pulumi.BoolOutput)
}

// Grant permission to use the KMIP Rekey operation.
func (o SecretRoleOutput) OperationRekey() pulumi.BoolOutput {
	return o.ApplyT(func(v *SecretRole) pulumi.BoolOutput { return v.OperationRekey }).(pulumi.BoolOutput)
}

// Grant permission to use the KMIP Revoke operation.
func (o SecretRoleOutput) OperationRevoke() pulumi.BoolOutput {
	return o.ApplyT(func(v *SecretRole) pulumi.BoolOutput { return v.OperationRevoke }).(pulumi.BoolOutput)
}

// The unique path this backend should be mounted at. Must
// not begin or end with a `/`. Defaults to `kmip`.
func (o SecretRoleOutput) Path() pulumi.StringOutput {
	return o.ApplyT(func(v *SecretRole) pulumi.StringOutput { return v.Path }).(pulumi.StringOutput)
}

// Name of the role.
func (o SecretRoleOutput) Role() pulumi.StringOutput {
	return o.ApplyT(func(v *SecretRole) pulumi.StringOutput { return v.Role }).(pulumi.StringOutput)
}

// Name of the scope.
func (o SecretRoleOutput) Scope() pulumi.StringOutput {
	return o.ApplyT(func(v *SecretRole) pulumi.StringOutput { return v.Scope }).(pulumi.StringOutput)
}

// Client certificate key bits, valid values depend on key type.
func (o SecretRoleOutput) TlsClientKeyBits() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SecretRole) pulumi.IntPtrOutput { return v.TlsClientKeyBits }).(pulumi.IntPtrOutput)
}

// Client certificate key type, `rsa` or `ec`.
func (o SecretRoleOutput) TlsClientKeyType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecretRole) pulumi.StringPtrOutput { return v.TlsClientKeyType }).(pulumi.StringPtrOutput)
}

// Client certificate TTL in seconds.
func (o SecretRoleOutput) TlsClientTtl() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SecretRole) pulumi.IntPtrOutput { return v.TlsClientTtl }).(pulumi.IntPtrOutput)
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
