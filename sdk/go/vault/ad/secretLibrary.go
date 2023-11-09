// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ad

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-vault/sdk/v5/go/vault/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-vault/sdk/v5/go/vault/ad"
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
//			_, err = ad.NewSecretLibrary(ctx, "qa", &ad.SecretLibraryArgs{
//				Backend: config.Backend,
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
// AD secret backend libraries can be imported using the `path`, e.g.
//
// ```sh
//
//	$ pulumi import vault:ad/secretLibrary:SecretLibrary role ad/library/bob
//
// ```
type SecretLibrary struct {
	pulumi.CustomResourceState

	// The path the AD secret backend is mounted at,
	// with no leading or trailing `/`s.
	Backend pulumi.StringOutput `pulumi:"backend"`
	// Disable enforcing that service accounts must be checked in by the entity or client token that checked them out.
	DisableCheckInEnforcement pulumi.BoolPtrOutput `pulumi:"disableCheckInEnforcement"`
	// The maximum password time-to-live in seconds. Defaults to the configuration
	// maxTtl if not provided.
	MaxTtl pulumi.IntOutput `pulumi:"maxTtl"`
	// The name to identify this set of service accounts.
	// Must be unique within the backend.
	Name pulumi.StringOutput `pulumi:"name"`
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
	// *Available only for Vault Enterprise*.
	Namespace pulumi.StringPtrOutput `pulumi:"namespace"`
	// Specifies the slice of service accounts mapped to this set.
	ServiceAccountNames pulumi.StringArrayOutput `pulumi:"serviceAccountNames"`
	// The password time-to-live in seconds. Defaults to the configuration
	// ttl if not provided.
	Ttl pulumi.IntOutput `pulumi:"ttl"`
}

// NewSecretLibrary registers a new resource with the given unique name, arguments, and options.
func NewSecretLibrary(ctx *pulumi.Context,
	name string, args *SecretLibraryArgs, opts ...pulumi.ResourceOption) (*SecretLibrary, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Backend == nil {
		return nil, errors.New("invalid value for required argument 'Backend'")
	}
	if args.ServiceAccountNames == nil {
		return nil, errors.New("invalid value for required argument 'ServiceAccountNames'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource SecretLibrary
	err := ctx.RegisterResource("vault:ad/secretLibrary:SecretLibrary", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSecretLibrary gets an existing SecretLibrary resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSecretLibrary(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SecretLibraryState, opts ...pulumi.ResourceOption) (*SecretLibrary, error) {
	var resource SecretLibrary
	err := ctx.ReadResource("vault:ad/secretLibrary:SecretLibrary", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SecretLibrary resources.
type secretLibraryState struct {
	// The path the AD secret backend is mounted at,
	// with no leading or trailing `/`s.
	Backend *string `pulumi:"backend"`
	// Disable enforcing that service accounts must be checked in by the entity or client token that checked them out.
	DisableCheckInEnforcement *bool `pulumi:"disableCheckInEnforcement"`
	// The maximum password time-to-live in seconds. Defaults to the configuration
	// maxTtl if not provided.
	MaxTtl *int `pulumi:"maxTtl"`
	// The name to identify this set of service accounts.
	// Must be unique within the backend.
	Name *string `pulumi:"name"`
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
	// *Available only for Vault Enterprise*.
	Namespace *string `pulumi:"namespace"`
	// Specifies the slice of service accounts mapped to this set.
	ServiceAccountNames []string `pulumi:"serviceAccountNames"`
	// The password time-to-live in seconds. Defaults to the configuration
	// ttl if not provided.
	Ttl *int `pulumi:"ttl"`
}

type SecretLibraryState struct {
	// The path the AD secret backend is mounted at,
	// with no leading or trailing `/`s.
	Backend pulumi.StringPtrInput
	// Disable enforcing that service accounts must be checked in by the entity or client token that checked them out.
	DisableCheckInEnforcement pulumi.BoolPtrInput
	// The maximum password time-to-live in seconds. Defaults to the configuration
	// maxTtl if not provided.
	MaxTtl pulumi.IntPtrInput
	// The name to identify this set of service accounts.
	// Must be unique within the backend.
	Name pulumi.StringPtrInput
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
	// *Available only for Vault Enterprise*.
	Namespace pulumi.StringPtrInput
	// Specifies the slice of service accounts mapped to this set.
	ServiceAccountNames pulumi.StringArrayInput
	// The password time-to-live in seconds. Defaults to the configuration
	// ttl if not provided.
	Ttl pulumi.IntPtrInput
}

func (SecretLibraryState) ElementType() reflect.Type {
	return reflect.TypeOf((*secretLibraryState)(nil)).Elem()
}

type secretLibraryArgs struct {
	// The path the AD secret backend is mounted at,
	// with no leading or trailing `/`s.
	Backend string `pulumi:"backend"`
	// Disable enforcing that service accounts must be checked in by the entity or client token that checked them out.
	DisableCheckInEnforcement *bool `pulumi:"disableCheckInEnforcement"`
	// The maximum password time-to-live in seconds. Defaults to the configuration
	// maxTtl if not provided.
	MaxTtl *int `pulumi:"maxTtl"`
	// The name to identify this set of service accounts.
	// Must be unique within the backend.
	Name *string `pulumi:"name"`
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
	// *Available only for Vault Enterprise*.
	Namespace *string `pulumi:"namespace"`
	// Specifies the slice of service accounts mapped to this set.
	ServiceAccountNames []string `pulumi:"serviceAccountNames"`
	// The password time-to-live in seconds. Defaults to the configuration
	// ttl if not provided.
	Ttl *int `pulumi:"ttl"`
}

// The set of arguments for constructing a SecretLibrary resource.
type SecretLibraryArgs struct {
	// The path the AD secret backend is mounted at,
	// with no leading or trailing `/`s.
	Backend pulumi.StringInput
	// Disable enforcing that service accounts must be checked in by the entity or client token that checked them out.
	DisableCheckInEnforcement pulumi.BoolPtrInput
	// The maximum password time-to-live in seconds. Defaults to the configuration
	// maxTtl if not provided.
	MaxTtl pulumi.IntPtrInput
	// The name to identify this set of service accounts.
	// Must be unique within the backend.
	Name pulumi.StringPtrInput
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
	// *Available only for Vault Enterprise*.
	Namespace pulumi.StringPtrInput
	// Specifies the slice of service accounts mapped to this set.
	ServiceAccountNames pulumi.StringArrayInput
	// The password time-to-live in seconds. Defaults to the configuration
	// ttl if not provided.
	Ttl pulumi.IntPtrInput
}

func (SecretLibraryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*secretLibraryArgs)(nil)).Elem()
}

type SecretLibraryInput interface {
	pulumi.Input

	ToSecretLibraryOutput() SecretLibraryOutput
	ToSecretLibraryOutputWithContext(ctx context.Context) SecretLibraryOutput
}

func (*SecretLibrary) ElementType() reflect.Type {
	return reflect.TypeOf((**SecretLibrary)(nil)).Elem()
}

func (i *SecretLibrary) ToSecretLibraryOutput() SecretLibraryOutput {
	return i.ToSecretLibraryOutputWithContext(context.Background())
}

func (i *SecretLibrary) ToSecretLibraryOutputWithContext(ctx context.Context) SecretLibraryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecretLibraryOutput)
}

func (i *SecretLibrary) ToOutput(ctx context.Context) pulumix.Output[*SecretLibrary] {
	return pulumix.Output[*SecretLibrary]{
		OutputState: i.ToSecretLibraryOutputWithContext(ctx).OutputState,
	}
}

// SecretLibraryArrayInput is an input type that accepts SecretLibraryArray and SecretLibraryArrayOutput values.
// You can construct a concrete instance of `SecretLibraryArrayInput` via:
//
//	SecretLibraryArray{ SecretLibraryArgs{...} }
type SecretLibraryArrayInput interface {
	pulumi.Input

	ToSecretLibraryArrayOutput() SecretLibraryArrayOutput
	ToSecretLibraryArrayOutputWithContext(context.Context) SecretLibraryArrayOutput
}

type SecretLibraryArray []SecretLibraryInput

func (SecretLibraryArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SecretLibrary)(nil)).Elem()
}

func (i SecretLibraryArray) ToSecretLibraryArrayOutput() SecretLibraryArrayOutput {
	return i.ToSecretLibraryArrayOutputWithContext(context.Background())
}

func (i SecretLibraryArray) ToSecretLibraryArrayOutputWithContext(ctx context.Context) SecretLibraryArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecretLibraryArrayOutput)
}

func (i SecretLibraryArray) ToOutput(ctx context.Context) pulumix.Output[[]*SecretLibrary] {
	return pulumix.Output[[]*SecretLibrary]{
		OutputState: i.ToSecretLibraryArrayOutputWithContext(ctx).OutputState,
	}
}

// SecretLibraryMapInput is an input type that accepts SecretLibraryMap and SecretLibraryMapOutput values.
// You can construct a concrete instance of `SecretLibraryMapInput` via:
//
//	SecretLibraryMap{ "key": SecretLibraryArgs{...} }
type SecretLibraryMapInput interface {
	pulumi.Input

	ToSecretLibraryMapOutput() SecretLibraryMapOutput
	ToSecretLibraryMapOutputWithContext(context.Context) SecretLibraryMapOutput
}

type SecretLibraryMap map[string]SecretLibraryInput

func (SecretLibraryMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SecretLibrary)(nil)).Elem()
}

func (i SecretLibraryMap) ToSecretLibraryMapOutput() SecretLibraryMapOutput {
	return i.ToSecretLibraryMapOutputWithContext(context.Background())
}

func (i SecretLibraryMap) ToSecretLibraryMapOutputWithContext(ctx context.Context) SecretLibraryMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecretLibraryMapOutput)
}

func (i SecretLibraryMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*SecretLibrary] {
	return pulumix.Output[map[string]*SecretLibrary]{
		OutputState: i.ToSecretLibraryMapOutputWithContext(ctx).OutputState,
	}
}

type SecretLibraryOutput struct{ *pulumi.OutputState }

func (SecretLibraryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SecretLibrary)(nil)).Elem()
}

func (o SecretLibraryOutput) ToSecretLibraryOutput() SecretLibraryOutput {
	return o
}

func (o SecretLibraryOutput) ToSecretLibraryOutputWithContext(ctx context.Context) SecretLibraryOutput {
	return o
}

func (o SecretLibraryOutput) ToOutput(ctx context.Context) pulumix.Output[*SecretLibrary] {
	return pulumix.Output[*SecretLibrary]{
		OutputState: o.OutputState,
	}
}

// The path the AD secret backend is mounted at,
// with no leading or trailing `/`s.
func (o SecretLibraryOutput) Backend() pulumi.StringOutput {
	return o.ApplyT(func(v *SecretLibrary) pulumi.StringOutput { return v.Backend }).(pulumi.StringOutput)
}

// Disable enforcing that service accounts must be checked in by the entity or client token that checked them out.
func (o SecretLibraryOutput) DisableCheckInEnforcement() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SecretLibrary) pulumi.BoolPtrOutput { return v.DisableCheckInEnforcement }).(pulumi.BoolPtrOutput)
}

// The maximum password time-to-live in seconds. Defaults to the configuration
// maxTtl if not provided.
func (o SecretLibraryOutput) MaxTtl() pulumi.IntOutput {
	return o.ApplyT(func(v *SecretLibrary) pulumi.IntOutput { return v.MaxTtl }).(pulumi.IntOutput)
}

// The name to identify this set of service accounts.
// Must be unique within the backend.
func (o SecretLibraryOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SecretLibrary) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The namespace to provision the resource in.
// The value should not contain leading or trailing forward slashes.
// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
// *Available only for Vault Enterprise*.
func (o SecretLibraryOutput) Namespace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecretLibrary) pulumi.StringPtrOutput { return v.Namespace }).(pulumi.StringPtrOutput)
}

// Specifies the slice of service accounts mapped to this set.
func (o SecretLibraryOutput) ServiceAccountNames() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *SecretLibrary) pulumi.StringArrayOutput { return v.ServiceAccountNames }).(pulumi.StringArrayOutput)
}

// The password time-to-live in seconds. Defaults to the configuration
// ttl if not provided.
func (o SecretLibraryOutput) Ttl() pulumi.IntOutput {
	return o.ApplyT(func(v *SecretLibrary) pulumi.IntOutput { return v.Ttl }).(pulumi.IntOutput)
}

type SecretLibraryArrayOutput struct{ *pulumi.OutputState }

func (SecretLibraryArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SecretLibrary)(nil)).Elem()
}

func (o SecretLibraryArrayOutput) ToSecretLibraryArrayOutput() SecretLibraryArrayOutput {
	return o
}

func (o SecretLibraryArrayOutput) ToSecretLibraryArrayOutputWithContext(ctx context.Context) SecretLibraryArrayOutput {
	return o
}

func (o SecretLibraryArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*SecretLibrary] {
	return pulumix.Output[[]*SecretLibrary]{
		OutputState: o.OutputState,
	}
}

func (o SecretLibraryArrayOutput) Index(i pulumi.IntInput) SecretLibraryOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SecretLibrary {
		return vs[0].([]*SecretLibrary)[vs[1].(int)]
	}).(SecretLibraryOutput)
}

type SecretLibraryMapOutput struct{ *pulumi.OutputState }

func (SecretLibraryMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SecretLibrary)(nil)).Elem()
}

func (o SecretLibraryMapOutput) ToSecretLibraryMapOutput() SecretLibraryMapOutput {
	return o
}

func (o SecretLibraryMapOutput) ToSecretLibraryMapOutputWithContext(ctx context.Context) SecretLibraryMapOutput {
	return o
}

func (o SecretLibraryMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*SecretLibrary] {
	return pulumix.Output[map[string]*SecretLibrary]{
		OutputState: o.OutputState,
	}
}

func (o SecretLibraryMapOutput) MapIndex(k pulumi.StringInput) SecretLibraryOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SecretLibrary {
		return vs[0].(map[string]*SecretLibrary)[vs[1].(string)]
	}).(SecretLibraryOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SecretLibraryInput)(nil)).Elem(), &SecretLibrary{})
	pulumi.RegisterInputType(reflect.TypeOf((*SecretLibraryArrayInput)(nil)).Elem(), SecretLibraryArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SecretLibraryMapInput)(nil)).Elem(), SecretLibraryMap{})
	pulumi.RegisterOutputType(SecretLibraryOutput{})
	pulumi.RegisterOutputType(SecretLibraryArrayOutput{})
	pulumi.RegisterOutputType(SecretLibraryMapOutput{})
}
