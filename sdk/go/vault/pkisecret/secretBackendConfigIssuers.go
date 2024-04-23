// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package pkisecret

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-vault/sdk/v6/go/vault/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-vault/sdk/v6/go/vault"
//	"github.com/pulumi/pulumi-vault/sdk/v6/go/vault/pkiSecret"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			pki, err := vault.NewMount(ctx, "pki", &vault.MountArgs{
//				Path:                   pulumi.String("pki"),
//				Type:                   pulumi.String("pki"),
//				DefaultLeaseTtlSeconds: pulumi.Int(3600),
//				MaxLeaseTtlSeconds:     pulumi.Int(86400),
//			})
//			if err != nil {
//				return err
//			}
//			root, err := pkiSecret.NewSecretBackendRootCert(ctx, "root", &pkiSecret.SecretBackendRootCertArgs{
//				Backend:    pki.Path,
//				Type:       pulumi.String("internal"),
//				CommonName: pulumi.String("test"),
//				Ttl:        pulumi.String("86400"),
//			})
//			if err != nil {
//				return err
//			}
//			example, err := pkiSecret.NewSecretBackendIssuer(ctx, "example", &pkiSecret.SecretBackendIssuerArgs{
//				Backend:    root.Backend,
//				IssuerRef:  root.IssuerId,
//				IssuerName: pulumi.String("example-issuer"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = pkiSecret.NewSecretBackendConfigIssuers(ctx, "config", &pkiSecret.SecretBackendConfigIssuersArgs{
//				Backend:                    pki.Path,
//				Default:                    example.IssuerId,
//				DefaultFollowsLatestIssuer: pulumi.Bool(true),
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
// PKI secret backend config issuers can be imported using the path, e.g.
//
// ```sh
// $ pulumi import vault:pkiSecret/secretBackendConfigIssuers:SecretBackendConfigIssuers config pki/config/issuers
// ```
type SecretBackendConfigIssuers struct {
	pulumi.CustomResourceState

	// The path the PKI secret backend is mounted at, with no
	// leading or trailing `/`s.
	Backend pulumi.StringOutput `pulumi:"backend"`
	// Specifies the default issuer by ID.
	Default pulumi.StringPtrOutput `pulumi:"default"`
	// Specifies whether a root creation
	// or an issuer import operation updates the default issuer to the newly added issuer.
	DefaultFollowsLatestIssuer pulumi.BoolOutput `pulumi:"defaultFollowsLatestIssuer"`
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
	// *Available only for Vault Enterprise*.
	Namespace pulumi.StringPtrOutput `pulumi:"namespace"`
}

// NewSecretBackendConfigIssuers registers a new resource with the given unique name, arguments, and options.
func NewSecretBackendConfigIssuers(ctx *pulumi.Context,
	name string, args *SecretBackendConfigIssuersArgs, opts ...pulumi.ResourceOption) (*SecretBackendConfigIssuers, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Backend == nil {
		return nil, errors.New("invalid value for required argument 'Backend'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource SecretBackendConfigIssuers
	err := ctx.RegisterResource("vault:pkiSecret/secretBackendConfigIssuers:SecretBackendConfigIssuers", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSecretBackendConfigIssuers gets an existing SecretBackendConfigIssuers resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSecretBackendConfigIssuers(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SecretBackendConfigIssuersState, opts ...pulumi.ResourceOption) (*SecretBackendConfigIssuers, error) {
	var resource SecretBackendConfigIssuers
	err := ctx.ReadResource("vault:pkiSecret/secretBackendConfigIssuers:SecretBackendConfigIssuers", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SecretBackendConfigIssuers resources.
type secretBackendConfigIssuersState struct {
	// The path the PKI secret backend is mounted at, with no
	// leading or trailing `/`s.
	Backend *string `pulumi:"backend"`
	// Specifies the default issuer by ID.
	Default *string `pulumi:"default"`
	// Specifies whether a root creation
	// or an issuer import operation updates the default issuer to the newly added issuer.
	DefaultFollowsLatestIssuer *bool `pulumi:"defaultFollowsLatestIssuer"`
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
	// *Available only for Vault Enterprise*.
	Namespace *string `pulumi:"namespace"`
}

type SecretBackendConfigIssuersState struct {
	// The path the PKI secret backend is mounted at, with no
	// leading or trailing `/`s.
	Backend pulumi.StringPtrInput
	// Specifies the default issuer by ID.
	Default pulumi.StringPtrInput
	// Specifies whether a root creation
	// or an issuer import operation updates the default issuer to the newly added issuer.
	DefaultFollowsLatestIssuer pulumi.BoolPtrInput
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
	// *Available only for Vault Enterprise*.
	Namespace pulumi.StringPtrInput
}

func (SecretBackendConfigIssuersState) ElementType() reflect.Type {
	return reflect.TypeOf((*secretBackendConfigIssuersState)(nil)).Elem()
}

type secretBackendConfigIssuersArgs struct {
	// The path the PKI secret backend is mounted at, with no
	// leading or trailing `/`s.
	Backend string `pulumi:"backend"`
	// Specifies the default issuer by ID.
	Default *string `pulumi:"default"`
	// Specifies whether a root creation
	// or an issuer import operation updates the default issuer to the newly added issuer.
	DefaultFollowsLatestIssuer *bool `pulumi:"defaultFollowsLatestIssuer"`
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
	// *Available only for Vault Enterprise*.
	Namespace *string `pulumi:"namespace"`
}

// The set of arguments for constructing a SecretBackendConfigIssuers resource.
type SecretBackendConfigIssuersArgs struct {
	// The path the PKI secret backend is mounted at, with no
	// leading or trailing `/`s.
	Backend pulumi.StringInput
	// Specifies the default issuer by ID.
	Default pulumi.StringPtrInput
	// Specifies whether a root creation
	// or an issuer import operation updates the default issuer to the newly added issuer.
	DefaultFollowsLatestIssuer pulumi.BoolPtrInput
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
	// *Available only for Vault Enterprise*.
	Namespace pulumi.StringPtrInput
}

func (SecretBackendConfigIssuersArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*secretBackendConfigIssuersArgs)(nil)).Elem()
}

type SecretBackendConfigIssuersInput interface {
	pulumi.Input

	ToSecretBackendConfigIssuersOutput() SecretBackendConfigIssuersOutput
	ToSecretBackendConfigIssuersOutputWithContext(ctx context.Context) SecretBackendConfigIssuersOutput
}

func (*SecretBackendConfigIssuers) ElementType() reflect.Type {
	return reflect.TypeOf((**SecretBackendConfigIssuers)(nil)).Elem()
}

func (i *SecretBackendConfigIssuers) ToSecretBackendConfigIssuersOutput() SecretBackendConfigIssuersOutput {
	return i.ToSecretBackendConfigIssuersOutputWithContext(context.Background())
}

func (i *SecretBackendConfigIssuers) ToSecretBackendConfigIssuersOutputWithContext(ctx context.Context) SecretBackendConfigIssuersOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecretBackendConfigIssuersOutput)
}

// SecretBackendConfigIssuersArrayInput is an input type that accepts SecretBackendConfigIssuersArray and SecretBackendConfigIssuersArrayOutput values.
// You can construct a concrete instance of `SecretBackendConfigIssuersArrayInput` via:
//
//	SecretBackendConfigIssuersArray{ SecretBackendConfigIssuersArgs{...} }
type SecretBackendConfigIssuersArrayInput interface {
	pulumi.Input

	ToSecretBackendConfigIssuersArrayOutput() SecretBackendConfigIssuersArrayOutput
	ToSecretBackendConfigIssuersArrayOutputWithContext(context.Context) SecretBackendConfigIssuersArrayOutput
}

type SecretBackendConfigIssuersArray []SecretBackendConfigIssuersInput

func (SecretBackendConfigIssuersArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SecretBackendConfigIssuers)(nil)).Elem()
}

func (i SecretBackendConfigIssuersArray) ToSecretBackendConfigIssuersArrayOutput() SecretBackendConfigIssuersArrayOutput {
	return i.ToSecretBackendConfigIssuersArrayOutputWithContext(context.Background())
}

func (i SecretBackendConfigIssuersArray) ToSecretBackendConfigIssuersArrayOutputWithContext(ctx context.Context) SecretBackendConfigIssuersArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecretBackendConfigIssuersArrayOutput)
}

// SecretBackendConfigIssuersMapInput is an input type that accepts SecretBackendConfigIssuersMap and SecretBackendConfigIssuersMapOutput values.
// You can construct a concrete instance of `SecretBackendConfigIssuersMapInput` via:
//
//	SecretBackendConfigIssuersMap{ "key": SecretBackendConfigIssuersArgs{...} }
type SecretBackendConfigIssuersMapInput interface {
	pulumi.Input

	ToSecretBackendConfigIssuersMapOutput() SecretBackendConfigIssuersMapOutput
	ToSecretBackendConfigIssuersMapOutputWithContext(context.Context) SecretBackendConfigIssuersMapOutput
}

type SecretBackendConfigIssuersMap map[string]SecretBackendConfigIssuersInput

func (SecretBackendConfigIssuersMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SecretBackendConfigIssuers)(nil)).Elem()
}

func (i SecretBackendConfigIssuersMap) ToSecretBackendConfigIssuersMapOutput() SecretBackendConfigIssuersMapOutput {
	return i.ToSecretBackendConfigIssuersMapOutputWithContext(context.Background())
}

func (i SecretBackendConfigIssuersMap) ToSecretBackendConfigIssuersMapOutputWithContext(ctx context.Context) SecretBackendConfigIssuersMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecretBackendConfigIssuersMapOutput)
}

type SecretBackendConfigIssuersOutput struct{ *pulumi.OutputState }

func (SecretBackendConfigIssuersOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SecretBackendConfigIssuers)(nil)).Elem()
}

func (o SecretBackendConfigIssuersOutput) ToSecretBackendConfigIssuersOutput() SecretBackendConfigIssuersOutput {
	return o
}

func (o SecretBackendConfigIssuersOutput) ToSecretBackendConfigIssuersOutputWithContext(ctx context.Context) SecretBackendConfigIssuersOutput {
	return o
}

// The path the PKI secret backend is mounted at, with no
// leading or trailing `/`s.
func (o SecretBackendConfigIssuersOutput) Backend() pulumi.StringOutput {
	return o.ApplyT(func(v *SecretBackendConfigIssuers) pulumi.StringOutput { return v.Backend }).(pulumi.StringOutput)
}

// Specifies the default issuer by ID.
func (o SecretBackendConfigIssuersOutput) Default() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecretBackendConfigIssuers) pulumi.StringPtrOutput { return v.Default }).(pulumi.StringPtrOutput)
}

// Specifies whether a root creation
// or an issuer import operation updates the default issuer to the newly added issuer.
func (o SecretBackendConfigIssuersOutput) DefaultFollowsLatestIssuer() pulumi.BoolOutput {
	return o.ApplyT(func(v *SecretBackendConfigIssuers) pulumi.BoolOutput { return v.DefaultFollowsLatestIssuer }).(pulumi.BoolOutput)
}

// The namespace to provision the resource in.
// The value should not contain leading or trailing forward slashes.
// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
// *Available only for Vault Enterprise*.
func (o SecretBackendConfigIssuersOutput) Namespace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecretBackendConfigIssuers) pulumi.StringPtrOutput { return v.Namespace }).(pulumi.StringPtrOutput)
}

type SecretBackendConfigIssuersArrayOutput struct{ *pulumi.OutputState }

func (SecretBackendConfigIssuersArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SecretBackendConfigIssuers)(nil)).Elem()
}

func (o SecretBackendConfigIssuersArrayOutput) ToSecretBackendConfigIssuersArrayOutput() SecretBackendConfigIssuersArrayOutput {
	return o
}

func (o SecretBackendConfigIssuersArrayOutput) ToSecretBackendConfigIssuersArrayOutputWithContext(ctx context.Context) SecretBackendConfigIssuersArrayOutput {
	return o
}

func (o SecretBackendConfigIssuersArrayOutput) Index(i pulumi.IntInput) SecretBackendConfigIssuersOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SecretBackendConfigIssuers {
		return vs[0].([]*SecretBackendConfigIssuers)[vs[1].(int)]
	}).(SecretBackendConfigIssuersOutput)
}

type SecretBackendConfigIssuersMapOutput struct{ *pulumi.OutputState }

func (SecretBackendConfigIssuersMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SecretBackendConfigIssuers)(nil)).Elem()
}

func (o SecretBackendConfigIssuersMapOutput) ToSecretBackendConfigIssuersMapOutput() SecretBackendConfigIssuersMapOutput {
	return o
}

func (o SecretBackendConfigIssuersMapOutput) ToSecretBackendConfigIssuersMapOutputWithContext(ctx context.Context) SecretBackendConfigIssuersMapOutput {
	return o
}

func (o SecretBackendConfigIssuersMapOutput) MapIndex(k pulumi.StringInput) SecretBackendConfigIssuersOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SecretBackendConfigIssuers {
		return vs[0].(map[string]*SecretBackendConfigIssuers)[vs[1].(string)]
	}).(SecretBackendConfigIssuersOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SecretBackendConfigIssuersInput)(nil)).Elem(), &SecretBackendConfigIssuers{})
	pulumi.RegisterInputType(reflect.TypeOf((*SecretBackendConfigIssuersArrayInput)(nil)).Elem(), SecretBackendConfigIssuersArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SecretBackendConfigIssuersMapInput)(nil)).Elem(), SecretBackendConfigIssuersMap{})
	pulumi.RegisterOutputType(SecretBackendConfigIssuersOutput{})
	pulumi.RegisterOutputType(SecretBackendConfigIssuersArrayOutput{})
	pulumi.RegisterOutputType(SecretBackendConfigIssuersMapOutput{})
}
