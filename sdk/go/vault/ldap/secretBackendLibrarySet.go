// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ldap

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-vault/sdk/v7/go/vault/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-vault/sdk/v7/go/vault/ldap"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			config, err := ldap.NewSecretBackend(ctx, "config", &ldap.SecretBackendArgs{
//				Path:        pulumi.String("ldap"),
//				Binddn:      pulumi.String("CN=Administrator,CN=Users,DC=corp,DC=example,DC=net"),
//				Bindpass:    pulumi.String("SuperSecretPassw0rd"),
//				Url:         pulumi.String("ldaps://localhost"),
//				InsecureTls: pulumi.Bool(true),
//				Userdn:      pulumi.String("CN=Users,DC=corp,DC=example,DC=net"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = ldap.NewSecretBackendLibrarySet(ctx, "qa", &ldap.SecretBackendLibrarySetArgs{
//				Mount: config.Path,
//				Name:  pulumi.String("qa"),
//				ServiceAccountNames: pulumi.StringArray{
//					pulumi.String("Bob"),
//					pulumi.String("Mary"),
//				},
//				Ttl:                       pulumi.Int(60),
//				DisableCheckInEnforcement: pulumi.Bool(true),
//				MaxTtl:                    pulumi.Int(120),
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
// LDAP secret backend libraries can be imported using the `path`, e.g.
//
// ```sh
// $ pulumi import vault:ldap/secretBackendLibrarySet:SecretBackendLibrarySet qa ldap/library/bob
// ```
type SecretBackendLibrarySet struct {
	pulumi.CustomResourceState

	// Disable enforcing that service
	// accounts must be checked in by the entity or client token that checked them
	// out. Defaults to false.
	DisableCheckInEnforcement pulumi.BoolPtrOutput `pulumi:"disableCheckInEnforcement"`
	// The maximum password time-to-live in seconds. Defaults
	// to the configuration maxTtl if not provided.
	MaxTtl pulumi.IntOutput `pulumi:"maxTtl"`
	// The path where the LDAP secrets backend is mounted.
	Mount pulumi.StringPtrOutput `pulumi:"mount"`
	// The name to identify this set of service accounts.
	// Must be unique within the backend.
	Name pulumi.StringOutput `pulumi:"name"`
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
	// *Available only for Vault Enterprise*.
	Namespace pulumi.StringPtrOutput `pulumi:"namespace"`
	// Specifies the slice of service accounts mapped to this set.
	ServiceAccountNames pulumi.StringArrayOutput `pulumi:"serviceAccountNames"`
	// The password time-to-live in seconds. Defaults to the configuration
	// ttl if not provided.
	Ttl pulumi.IntOutput `pulumi:"ttl"`
}

// NewSecretBackendLibrarySet registers a new resource with the given unique name, arguments, and options.
func NewSecretBackendLibrarySet(ctx *pulumi.Context,
	name string, args *SecretBackendLibrarySetArgs, opts ...pulumi.ResourceOption) (*SecretBackendLibrarySet, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ServiceAccountNames == nil {
		return nil, errors.New("invalid value for required argument 'ServiceAccountNames'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource SecretBackendLibrarySet
	err := ctx.RegisterResource("vault:ldap/secretBackendLibrarySet:SecretBackendLibrarySet", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSecretBackendLibrarySet gets an existing SecretBackendLibrarySet resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSecretBackendLibrarySet(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SecretBackendLibrarySetState, opts ...pulumi.ResourceOption) (*SecretBackendLibrarySet, error) {
	var resource SecretBackendLibrarySet
	err := ctx.ReadResource("vault:ldap/secretBackendLibrarySet:SecretBackendLibrarySet", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SecretBackendLibrarySet resources.
type secretBackendLibrarySetState struct {
	// Disable enforcing that service
	// accounts must be checked in by the entity or client token that checked them
	// out. Defaults to false.
	DisableCheckInEnforcement *bool `pulumi:"disableCheckInEnforcement"`
	// The maximum password time-to-live in seconds. Defaults
	// to the configuration maxTtl if not provided.
	MaxTtl *int `pulumi:"maxTtl"`
	// The path where the LDAP secrets backend is mounted.
	Mount *string `pulumi:"mount"`
	// The name to identify this set of service accounts.
	// Must be unique within the backend.
	Name *string `pulumi:"name"`
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
	// *Available only for Vault Enterprise*.
	Namespace *string `pulumi:"namespace"`
	// Specifies the slice of service accounts mapped to this set.
	ServiceAccountNames []string `pulumi:"serviceAccountNames"`
	// The password time-to-live in seconds. Defaults to the configuration
	// ttl if not provided.
	Ttl *int `pulumi:"ttl"`
}

type SecretBackendLibrarySetState struct {
	// Disable enforcing that service
	// accounts must be checked in by the entity or client token that checked them
	// out. Defaults to false.
	DisableCheckInEnforcement pulumi.BoolPtrInput
	// The maximum password time-to-live in seconds. Defaults
	// to the configuration maxTtl if not provided.
	MaxTtl pulumi.IntPtrInput
	// The path where the LDAP secrets backend is mounted.
	Mount pulumi.StringPtrInput
	// The name to identify this set of service accounts.
	// Must be unique within the backend.
	Name pulumi.StringPtrInput
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
	// *Available only for Vault Enterprise*.
	Namespace pulumi.StringPtrInput
	// Specifies the slice of service accounts mapped to this set.
	ServiceAccountNames pulumi.StringArrayInput
	// The password time-to-live in seconds. Defaults to the configuration
	// ttl if not provided.
	Ttl pulumi.IntPtrInput
}

func (SecretBackendLibrarySetState) ElementType() reflect.Type {
	return reflect.TypeOf((*secretBackendLibrarySetState)(nil)).Elem()
}

type secretBackendLibrarySetArgs struct {
	// Disable enforcing that service
	// accounts must be checked in by the entity or client token that checked them
	// out. Defaults to false.
	DisableCheckInEnforcement *bool `pulumi:"disableCheckInEnforcement"`
	// The maximum password time-to-live in seconds. Defaults
	// to the configuration maxTtl if not provided.
	MaxTtl *int `pulumi:"maxTtl"`
	// The path where the LDAP secrets backend is mounted.
	Mount *string `pulumi:"mount"`
	// The name to identify this set of service accounts.
	// Must be unique within the backend.
	Name *string `pulumi:"name"`
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
	// *Available only for Vault Enterprise*.
	Namespace *string `pulumi:"namespace"`
	// Specifies the slice of service accounts mapped to this set.
	ServiceAccountNames []string `pulumi:"serviceAccountNames"`
	// The password time-to-live in seconds. Defaults to the configuration
	// ttl if not provided.
	Ttl *int `pulumi:"ttl"`
}

// The set of arguments for constructing a SecretBackendLibrarySet resource.
type SecretBackendLibrarySetArgs struct {
	// Disable enforcing that service
	// accounts must be checked in by the entity or client token that checked them
	// out. Defaults to false.
	DisableCheckInEnforcement pulumi.BoolPtrInput
	// The maximum password time-to-live in seconds. Defaults
	// to the configuration maxTtl if not provided.
	MaxTtl pulumi.IntPtrInput
	// The path where the LDAP secrets backend is mounted.
	Mount pulumi.StringPtrInput
	// The name to identify this set of service accounts.
	// Must be unique within the backend.
	Name pulumi.StringPtrInput
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
	// *Available only for Vault Enterprise*.
	Namespace pulumi.StringPtrInput
	// Specifies the slice of service accounts mapped to this set.
	ServiceAccountNames pulumi.StringArrayInput
	// The password time-to-live in seconds. Defaults to the configuration
	// ttl if not provided.
	Ttl pulumi.IntPtrInput
}

func (SecretBackendLibrarySetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*secretBackendLibrarySetArgs)(nil)).Elem()
}

type SecretBackendLibrarySetInput interface {
	pulumi.Input

	ToSecretBackendLibrarySetOutput() SecretBackendLibrarySetOutput
	ToSecretBackendLibrarySetOutputWithContext(ctx context.Context) SecretBackendLibrarySetOutput
}

func (*SecretBackendLibrarySet) ElementType() reflect.Type {
	return reflect.TypeOf((**SecretBackendLibrarySet)(nil)).Elem()
}

func (i *SecretBackendLibrarySet) ToSecretBackendLibrarySetOutput() SecretBackendLibrarySetOutput {
	return i.ToSecretBackendLibrarySetOutputWithContext(context.Background())
}

func (i *SecretBackendLibrarySet) ToSecretBackendLibrarySetOutputWithContext(ctx context.Context) SecretBackendLibrarySetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecretBackendLibrarySetOutput)
}

// SecretBackendLibrarySetArrayInput is an input type that accepts SecretBackendLibrarySetArray and SecretBackendLibrarySetArrayOutput values.
// You can construct a concrete instance of `SecretBackendLibrarySetArrayInput` via:
//
//	SecretBackendLibrarySetArray{ SecretBackendLibrarySetArgs{...} }
type SecretBackendLibrarySetArrayInput interface {
	pulumi.Input

	ToSecretBackendLibrarySetArrayOutput() SecretBackendLibrarySetArrayOutput
	ToSecretBackendLibrarySetArrayOutputWithContext(context.Context) SecretBackendLibrarySetArrayOutput
}

type SecretBackendLibrarySetArray []SecretBackendLibrarySetInput

func (SecretBackendLibrarySetArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SecretBackendLibrarySet)(nil)).Elem()
}

func (i SecretBackendLibrarySetArray) ToSecretBackendLibrarySetArrayOutput() SecretBackendLibrarySetArrayOutput {
	return i.ToSecretBackendLibrarySetArrayOutputWithContext(context.Background())
}

func (i SecretBackendLibrarySetArray) ToSecretBackendLibrarySetArrayOutputWithContext(ctx context.Context) SecretBackendLibrarySetArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecretBackendLibrarySetArrayOutput)
}

// SecretBackendLibrarySetMapInput is an input type that accepts SecretBackendLibrarySetMap and SecretBackendLibrarySetMapOutput values.
// You can construct a concrete instance of `SecretBackendLibrarySetMapInput` via:
//
//	SecretBackendLibrarySetMap{ "key": SecretBackendLibrarySetArgs{...} }
type SecretBackendLibrarySetMapInput interface {
	pulumi.Input

	ToSecretBackendLibrarySetMapOutput() SecretBackendLibrarySetMapOutput
	ToSecretBackendLibrarySetMapOutputWithContext(context.Context) SecretBackendLibrarySetMapOutput
}

type SecretBackendLibrarySetMap map[string]SecretBackendLibrarySetInput

func (SecretBackendLibrarySetMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SecretBackendLibrarySet)(nil)).Elem()
}

func (i SecretBackendLibrarySetMap) ToSecretBackendLibrarySetMapOutput() SecretBackendLibrarySetMapOutput {
	return i.ToSecretBackendLibrarySetMapOutputWithContext(context.Background())
}

func (i SecretBackendLibrarySetMap) ToSecretBackendLibrarySetMapOutputWithContext(ctx context.Context) SecretBackendLibrarySetMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecretBackendLibrarySetMapOutput)
}

type SecretBackendLibrarySetOutput struct{ *pulumi.OutputState }

func (SecretBackendLibrarySetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SecretBackendLibrarySet)(nil)).Elem()
}

func (o SecretBackendLibrarySetOutput) ToSecretBackendLibrarySetOutput() SecretBackendLibrarySetOutput {
	return o
}

func (o SecretBackendLibrarySetOutput) ToSecretBackendLibrarySetOutputWithContext(ctx context.Context) SecretBackendLibrarySetOutput {
	return o
}

// Disable enforcing that service
// accounts must be checked in by the entity or client token that checked them
// out. Defaults to false.
func (o SecretBackendLibrarySetOutput) DisableCheckInEnforcement() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SecretBackendLibrarySet) pulumi.BoolPtrOutput { return v.DisableCheckInEnforcement }).(pulumi.BoolPtrOutput)
}

// The maximum password time-to-live in seconds. Defaults
// to the configuration maxTtl if not provided.
func (o SecretBackendLibrarySetOutput) MaxTtl() pulumi.IntOutput {
	return o.ApplyT(func(v *SecretBackendLibrarySet) pulumi.IntOutput { return v.MaxTtl }).(pulumi.IntOutput)
}

// The path where the LDAP secrets backend is mounted.
func (o SecretBackendLibrarySetOutput) Mount() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecretBackendLibrarySet) pulumi.StringPtrOutput { return v.Mount }).(pulumi.StringPtrOutput)
}

// The name to identify this set of service accounts.
// Must be unique within the backend.
func (o SecretBackendLibrarySetOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SecretBackendLibrarySet) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The namespace to provision the resource in.
// The value should not contain leading or trailing forward slashes.
// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
// *Available only for Vault Enterprise*.
func (o SecretBackendLibrarySetOutput) Namespace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecretBackendLibrarySet) pulumi.StringPtrOutput { return v.Namespace }).(pulumi.StringPtrOutput)
}

// Specifies the slice of service accounts mapped to this set.
func (o SecretBackendLibrarySetOutput) ServiceAccountNames() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *SecretBackendLibrarySet) pulumi.StringArrayOutput { return v.ServiceAccountNames }).(pulumi.StringArrayOutput)
}

// The password time-to-live in seconds. Defaults to the configuration
// ttl if not provided.
func (o SecretBackendLibrarySetOutput) Ttl() pulumi.IntOutput {
	return o.ApplyT(func(v *SecretBackendLibrarySet) pulumi.IntOutput { return v.Ttl }).(pulumi.IntOutput)
}

type SecretBackendLibrarySetArrayOutput struct{ *pulumi.OutputState }

func (SecretBackendLibrarySetArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SecretBackendLibrarySet)(nil)).Elem()
}

func (o SecretBackendLibrarySetArrayOutput) ToSecretBackendLibrarySetArrayOutput() SecretBackendLibrarySetArrayOutput {
	return o
}

func (o SecretBackendLibrarySetArrayOutput) ToSecretBackendLibrarySetArrayOutputWithContext(ctx context.Context) SecretBackendLibrarySetArrayOutput {
	return o
}

func (o SecretBackendLibrarySetArrayOutput) Index(i pulumi.IntInput) SecretBackendLibrarySetOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SecretBackendLibrarySet {
		return vs[0].([]*SecretBackendLibrarySet)[vs[1].(int)]
	}).(SecretBackendLibrarySetOutput)
}

type SecretBackendLibrarySetMapOutput struct{ *pulumi.OutputState }

func (SecretBackendLibrarySetMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SecretBackendLibrarySet)(nil)).Elem()
}

func (o SecretBackendLibrarySetMapOutput) ToSecretBackendLibrarySetMapOutput() SecretBackendLibrarySetMapOutput {
	return o
}

func (o SecretBackendLibrarySetMapOutput) ToSecretBackendLibrarySetMapOutputWithContext(ctx context.Context) SecretBackendLibrarySetMapOutput {
	return o
}

func (o SecretBackendLibrarySetMapOutput) MapIndex(k pulumi.StringInput) SecretBackendLibrarySetOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SecretBackendLibrarySet {
		return vs[0].(map[string]*SecretBackendLibrarySet)[vs[1].(string)]
	}).(SecretBackendLibrarySetOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SecretBackendLibrarySetInput)(nil)).Elem(), &SecretBackendLibrarySet{})
	pulumi.RegisterInputType(reflect.TypeOf((*SecretBackendLibrarySetArrayInput)(nil)).Elem(), SecretBackendLibrarySetArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SecretBackendLibrarySetMapInput)(nil)).Elem(), SecretBackendLibrarySetMap{})
	pulumi.RegisterOutputType(SecretBackendLibrarySetOutput{})
	pulumi.RegisterOutputType(SecretBackendLibrarySetArrayOutput{})
	pulumi.RegisterOutputType(SecretBackendLibrarySetMapOutput{})
}
