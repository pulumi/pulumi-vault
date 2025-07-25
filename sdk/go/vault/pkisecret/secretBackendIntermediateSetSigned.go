// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package pkisecret

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
//	"github.com/pulumi/pulumi-vault/sdk/v7/go/vault"
//	"github.com/pulumi/pulumi-vault/sdk/v7/go/vault/pkisecret"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			root, err := vault.NewMount(ctx, "root", &vault.MountArgs{
//				Path:                   pulumi.String("pki-root"),
//				Type:                   pulumi.String("pki"),
//				Description:            pulumi.String("root"),
//				DefaultLeaseTtlSeconds: pulumi.Int(8640000),
//				MaxLeaseTtlSeconds:     pulumi.Int(8640000),
//			})
//			if err != nil {
//				return err
//			}
//			intermediate, err := vault.NewMount(ctx, "intermediate", &vault.MountArgs{
//				Path:                   pulumi.String("pki-int"),
//				Type:                   root.Type,
//				Description:            pulumi.String("intermediate"),
//				DefaultLeaseTtlSeconds: pulumi.Int(86400),
//				MaxLeaseTtlSeconds:     pulumi.Int(86400),
//			})
//			if err != nil {
//				return err
//			}
//			example, err := pkisecret.NewSecretBackendRootCert(ctx, "example", &pkisecret.SecretBackendRootCertArgs{
//				Backend:           root.Path,
//				Type:              pulumi.String("internal"),
//				CommonName:        pulumi.String("RootOrg Root CA"),
//				Ttl:               pulumi.String("86400"),
//				Format:            pulumi.String("pem"),
//				PrivateKeyFormat:  pulumi.String("der"),
//				KeyType:           pulumi.String("rsa"),
//				KeyBits:           pulumi.Int(4096),
//				ExcludeCnFromSans: pulumi.Bool(true),
//				Ou:                pulumi.String("Organizational Unit"),
//				Organization:      pulumi.String("RootOrg"),
//				Country:           pulumi.String("US"),
//				Locality:          pulumi.String("San Francisco"),
//				Province:          pulumi.String("CA"),
//			})
//			if err != nil {
//				return err
//			}
//			exampleSecretBackendIntermediateCertRequest, err := pkisecret.NewSecretBackendIntermediateCertRequest(ctx, "example", &pkisecret.SecretBackendIntermediateCertRequestArgs{
//				Backend:    intermediate.Path,
//				Type:       example.Type,
//				CommonName: pulumi.String("SubOrg Intermediate CA"),
//			})
//			if err != nil {
//				return err
//			}
//			exampleSecretBackendRootSignIntermediate, err := pkisecret.NewSecretBackendRootSignIntermediate(ctx, "example", &pkisecret.SecretBackendRootSignIntermediateArgs{
//				Backend:           root.Path,
//				Csr:               exampleSecretBackendIntermediateCertRequest.Csr,
//				CommonName:        pulumi.String("SubOrg Intermediate CA"),
//				ExcludeCnFromSans: pulumi.Bool(true),
//				Ou:                pulumi.String("SubUnit"),
//				Organization:      pulumi.String("SubOrg"),
//				Country:           pulumi.String("US"),
//				Locality:          pulumi.String("San Francisco"),
//				Province:          pulumi.String("CA"),
//				Revoke:            pulumi.Bool(true),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = pkisecret.NewSecretBackendIntermediateSetSigned(ctx, "example", &pkisecret.SecretBackendIntermediateSetSignedArgs{
//				Backend:     intermediate.Path,
//				Certificate: exampleSecretBackendRootSignIntermediate.Certificate,
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
type SecretBackendIntermediateSetSigned struct {
	pulumi.CustomResourceState

	// The PKI secret backend the resource belongs to.
	Backend pulumi.StringOutput `pulumi:"backend"`
	// Specifies the PEM encoded certificate. May optionally append additional
	// CA certificates to populate the whole chain, which will then enable returning the full chain from
	// issue and sign operations.
	Certificate pulumi.StringOutput `pulumi:"certificate"`
	// The imported issuers indicating which issuers were created as part of
	// this request.
	ImportedIssuers pulumi.StringArrayOutput `pulumi:"importedIssuers"`
	// The imported keys indicating which keys were created as part of this request.
	ImportedKeys pulumi.StringArrayOutput `pulumi:"importedKeys"`
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
	// *Available only for Vault Enterprise*.
	Namespace pulumi.StringPtrOutput `pulumi:"namespace"`
}

// NewSecretBackendIntermediateSetSigned registers a new resource with the given unique name, arguments, and options.
func NewSecretBackendIntermediateSetSigned(ctx *pulumi.Context,
	name string, args *SecretBackendIntermediateSetSignedArgs, opts ...pulumi.ResourceOption) (*SecretBackendIntermediateSetSigned, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Backend == nil {
		return nil, errors.New("invalid value for required argument 'Backend'")
	}
	if args.Certificate == nil {
		return nil, errors.New("invalid value for required argument 'Certificate'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource SecretBackendIntermediateSetSigned
	err := ctx.RegisterResource("vault:pkiSecret/secretBackendIntermediateSetSigned:SecretBackendIntermediateSetSigned", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSecretBackendIntermediateSetSigned gets an existing SecretBackendIntermediateSetSigned resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSecretBackendIntermediateSetSigned(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SecretBackendIntermediateSetSignedState, opts ...pulumi.ResourceOption) (*SecretBackendIntermediateSetSigned, error) {
	var resource SecretBackendIntermediateSetSigned
	err := ctx.ReadResource("vault:pkiSecret/secretBackendIntermediateSetSigned:SecretBackendIntermediateSetSigned", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SecretBackendIntermediateSetSigned resources.
type secretBackendIntermediateSetSignedState struct {
	// The PKI secret backend the resource belongs to.
	Backend *string `pulumi:"backend"`
	// Specifies the PEM encoded certificate. May optionally append additional
	// CA certificates to populate the whole chain, which will then enable returning the full chain from
	// issue and sign operations.
	Certificate *string `pulumi:"certificate"`
	// The imported issuers indicating which issuers were created as part of
	// this request.
	ImportedIssuers []string `pulumi:"importedIssuers"`
	// The imported keys indicating which keys were created as part of this request.
	ImportedKeys []string `pulumi:"importedKeys"`
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
	// *Available only for Vault Enterprise*.
	Namespace *string `pulumi:"namespace"`
}

type SecretBackendIntermediateSetSignedState struct {
	// The PKI secret backend the resource belongs to.
	Backend pulumi.StringPtrInput
	// Specifies the PEM encoded certificate. May optionally append additional
	// CA certificates to populate the whole chain, which will then enable returning the full chain from
	// issue and sign operations.
	Certificate pulumi.StringPtrInput
	// The imported issuers indicating which issuers were created as part of
	// this request.
	ImportedIssuers pulumi.StringArrayInput
	// The imported keys indicating which keys were created as part of this request.
	ImportedKeys pulumi.StringArrayInput
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
	// *Available only for Vault Enterprise*.
	Namespace pulumi.StringPtrInput
}

func (SecretBackendIntermediateSetSignedState) ElementType() reflect.Type {
	return reflect.TypeOf((*secretBackendIntermediateSetSignedState)(nil)).Elem()
}

type secretBackendIntermediateSetSignedArgs struct {
	// The PKI secret backend the resource belongs to.
	Backend string `pulumi:"backend"`
	// Specifies the PEM encoded certificate. May optionally append additional
	// CA certificates to populate the whole chain, which will then enable returning the full chain from
	// issue and sign operations.
	Certificate string `pulumi:"certificate"`
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
	// *Available only for Vault Enterprise*.
	Namespace *string `pulumi:"namespace"`
}

// The set of arguments for constructing a SecretBackendIntermediateSetSigned resource.
type SecretBackendIntermediateSetSignedArgs struct {
	// The PKI secret backend the resource belongs to.
	Backend pulumi.StringInput
	// Specifies the PEM encoded certificate. May optionally append additional
	// CA certificates to populate the whole chain, which will then enable returning the full chain from
	// issue and sign operations.
	Certificate pulumi.StringInput
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
	// *Available only for Vault Enterprise*.
	Namespace pulumi.StringPtrInput
}

func (SecretBackendIntermediateSetSignedArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*secretBackendIntermediateSetSignedArgs)(nil)).Elem()
}

type SecretBackendIntermediateSetSignedInput interface {
	pulumi.Input

	ToSecretBackendIntermediateSetSignedOutput() SecretBackendIntermediateSetSignedOutput
	ToSecretBackendIntermediateSetSignedOutputWithContext(ctx context.Context) SecretBackendIntermediateSetSignedOutput
}

func (*SecretBackendIntermediateSetSigned) ElementType() reflect.Type {
	return reflect.TypeOf((**SecretBackendIntermediateSetSigned)(nil)).Elem()
}

func (i *SecretBackendIntermediateSetSigned) ToSecretBackendIntermediateSetSignedOutput() SecretBackendIntermediateSetSignedOutput {
	return i.ToSecretBackendIntermediateSetSignedOutputWithContext(context.Background())
}

func (i *SecretBackendIntermediateSetSigned) ToSecretBackendIntermediateSetSignedOutputWithContext(ctx context.Context) SecretBackendIntermediateSetSignedOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecretBackendIntermediateSetSignedOutput)
}

// SecretBackendIntermediateSetSignedArrayInput is an input type that accepts SecretBackendIntermediateSetSignedArray and SecretBackendIntermediateSetSignedArrayOutput values.
// You can construct a concrete instance of `SecretBackendIntermediateSetSignedArrayInput` via:
//
//	SecretBackendIntermediateSetSignedArray{ SecretBackendIntermediateSetSignedArgs{...} }
type SecretBackendIntermediateSetSignedArrayInput interface {
	pulumi.Input

	ToSecretBackendIntermediateSetSignedArrayOutput() SecretBackendIntermediateSetSignedArrayOutput
	ToSecretBackendIntermediateSetSignedArrayOutputWithContext(context.Context) SecretBackendIntermediateSetSignedArrayOutput
}

type SecretBackendIntermediateSetSignedArray []SecretBackendIntermediateSetSignedInput

func (SecretBackendIntermediateSetSignedArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SecretBackendIntermediateSetSigned)(nil)).Elem()
}

func (i SecretBackendIntermediateSetSignedArray) ToSecretBackendIntermediateSetSignedArrayOutput() SecretBackendIntermediateSetSignedArrayOutput {
	return i.ToSecretBackendIntermediateSetSignedArrayOutputWithContext(context.Background())
}

func (i SecretBackendIntermediateSetSignedArray) ToSecretBackendIntermediateSetSignedArrayOutputWithContext(ctx context.Context) SecretBackendIntermediateSetSignedArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecretBackendIntermediateSetSignedArrayOutput)
}

// SecretBackendIntermediateSetSignedMapInput is an input type that accepts SecretBackendIntermediateSetSignedMap and SecretBackendIntermediateSetSignedMapOutput values.
// You can construct a concrete instance of `SecretBackendIntermediateSetSignedMapInput` via:
//
//	SecretBackendIntermediateSetSignedMap{ "key": SecretBackendIntermediateSetSignedArgs{...} }
type SecretBackendIntermediateSetSignedMapInput interface {
	pulumi.Input

	ToSecretBackendIntermediateSetSignedMapOutput() SecretBackendIntermediateSetSignedMapOutput
	ToSecretBackendIntermediateSetSignedMapOutputWithContext(context.Context) SecretBackendIntermediateSetSignedMapOutput
}

type SecretBackendIntermediateSetSignedMap map[string]SecretBackendIntermediateSetSignedInput

func (SecretBackendIntermediateSetSignedMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SecretBackendIntermediateSetSigned)(nil)).Elem()
}

func (i SecretBackendIntermediateSetSignedMap) ToSecretBackendIntermediateSetSignedMapOutput() SecretBackendIntermediateSetSignedMapOutput {
	return i.ToSecretBackendIntermediateSetSignedMapOutputWithContext(context.Background())
}

func (i SecretBackendIntermediateSetSignedMap) ToSecretBackendIntermediateSetSignedMapOutputWithContext(ctx context.Context) SecretBackendIntermediateSetSignedMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecretBackendIntermediateSetSignedMapOutput)
}

type SecretBackendIntermediateSetSignedOutput struct{ *pulumi.OutputState }

func (SecretBackendIntermediateSetSignedOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SecretBackendIntermediateSetSigned)(nil)).Elem()
}

func (o SecretBackendIntermediateSetSignedOutput) ToSecretBackendIntermediateSetSignedOutput() SecretBackendIntermediateSetSignedOutput {
	return o
}

func (o SecretBackendIntermediateSetSignedOutput) ToSecretBackendIntermediateSetSignedOutputWithContext(ctx context.Context) SecretBackendIntermediateSetSignedOutput {
	return o
}

// The PKI secret backend the resource belongs to.
func (o SecretBackendIntermediateSetSignedOutput) Backend() pulumi.StringOutput {
	return o.ApplyT(func(v *SecretBackendIntermediateSetSigned) pulumi.StringOutput { return v.Backend }).(pulumi.StringOutput)
}

// Specifies the PEM encoded certificate. May optionally append additional
// CA certificates to populate the whole chain, which will then enable returning the full chain from
// issue and sign operations.
func (o SecretBackendIntermediateSetSignedOutput) Certificate() pulumi.StringOutput {
	return o.ApplyT(func(v *SecretBackendIntermediateSetSigned) pulumi.StringOutput { return v.Certificate }).(pulumi.StringOutput)
}

// The imported issuers indicating which issuers were created as part of
// this request.
func (o SecretBackendIntermediateSetSignedOutput) ImportedIssuers() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *SecretBackendIntermediateSetSigned) pulumi.StringArrayOutput { return v.ImportedIssuers }).(pulumi.StringArrayOutput)
}

// The imported keys indicating which keys were created as part of this request.
func (o SecretBackendIntermediateSetSignedOutput) ImportedKeys() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *SecretBackendIntermediateSetSigned) pulumi.StringArrayOutput { return v.ImportedKeys }).(pulumi.StringArrayOutput)
}

// The namespace to provision the resource in.
// The value should not contain leading or trailing forward slashes.
// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
// *Available only for Vault Enterprise*.
func (o SecretBackendIntermediateSetSignedOutput) Namespace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecretBackendIntermediateSetSigned) pulumi.StringPtrOutput { return v.Namespace }).(pulumi.StringPtrOutput)
}

type SecretBackendIntermediateSetSignedArrayOutput struct{ *pulumi.OutputState }

func (SecretBackendIntermediateSetSignedArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SecretBackendIntermediateSetSigned)(nil)).Elem()
}

func (o SecretBackendIntermediateSetSignedArrayOutput) ToSecretBackendIntermediateSetSignedArrayOutput() SecretBackendIntermediateSetSignedArrayOutput {
	return o
}

func (o SecretBackendIntermediateSetSignedArrayOutput) ToSecretBackendIntermediateSetSignedArrayOutputWithContext(ctx context.Context) SecretBackendIntermediateSetSignedArrayOutput {
	return o
}

func (o SecretBackendIntermediateSetSignedArrayOutput) Index(i pulumi.IntInput) SecretBackendIntermediateSetSignedOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SecretBackendIntermediateSetSigned {
		return vs[0].([]*SecretBackendIntermediateSetSigned)[vs[1].(int)]
	}).(SecretBackendIntermediateSetSignedOutput)
}

type SecretBackendIntermediateSetSignedMapOutput struct{ *pulumi.OutputState }

func (SecretBackendIntermediateSetSignedMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SecretBackendIntermediateSetSigned)(nil)).Elem()
}

func (o SecretBackendIntermediateSetSignedMapOutput) ToSecretBackendIntermediateSetSignedMapOutput() SecretBackendIntermediateSetSignedMapOutput {
	return o
}

func (o SecretBackendIntermediateSetSignedMapOutput) ToSecretBackendIntermediateSetSignedMapOutputWithContext(ctx context.Context) SecretBackendIntermediateSetSignedMapOutput {
	return o
}

func (o SecretBackendIntermediateSetSignedMapOutput) MapIndex(k pulumi.StringInput) SecretBackendIntermediateSetSignedOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SecretBackendIntermediateSetSigned {
		return vs[0].(map[string]*SecretBackendIntermediateSetSigned)[vs[1].(string)]
	}).(SecretBackendIntermediateSetSignedOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SecretBackendIntermediateSetSignedInput)(nil)).Elem(), &SecretBackendIntermediateSetSigned{})
	pulumi.RegisterInputType(reflect.TypeOf((*SecretBackendIntermediateSetSignedArrayInput)(nil)).Elem(), SecretBackendIntermediateSetSignedArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SecretBackendIntermediateSetSignedMapInput)(nil)).Elem(), SecretBackendIntermediateSetSignedMap{})
	pulumi.RegisterOutputType(SecretBackendIntermediateSetSignedOutput{})
	pulumi.RegisterOutputType(SecretBackendIntermediateSetSignedArrayOutput{})
	pulumi.RegisterOutputType(SecretBackendIntermediateSetSignedMapOutput{})
}
