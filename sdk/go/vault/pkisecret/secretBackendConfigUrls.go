// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package pkisecret

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Allows setting the issuing certificate endpoints, CRL distribution points, and OCSP server endpoints that will be encoded into issued certificates.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-vault/sdk/v5/go/vault"
// 	"github.com/pulumi/pulumi-vault/sdk/v5/go/vault/pkiSecret"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		root, err := vault.NewMount(ctx, "root", &vault.MountArgs{
// 			Path:                   pulumi.String("pki-root"),
// 			Type:                   pulumi.String("pki"),
// 			Description:            pulumi.String("root PKI"),
// 			DefaultLeaseTtlSeconds: pulumi.Int(8640000),
// 			MaxLeaseTtlSeconds:     pulumi.Int(8640000),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = pkiSecret.NewSecretBackendConfigUrls(ctx, "example", &pkiSecret.SecretBackendConfigUrlsArgs{
// 			Backend: root.Path,
// 			IssuingCertificates: pulumi.StringArray{
// 				pulumi.String("http://127.0.0.1:8200/v1/pki/ca"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## Import
//
// The PKI config URLs can be imported using the resource's `id`.
//
// In the case of the example above the `id` would be `pki-root/config/urls`,
//
// where the `pki-root` component is the resource's `backend`, e.g.
//
// ```sh
//  $ pulumi import vault:pkiSecret/secretBackendConfigUrls:SecretBackendConfigUrls example pki-root/config/urls
// ```
type SecretBackendConfigUrls struct {
	pulumi.CustomResourceState

	// The path the PKI secret backend is mounted at, with no leading or trailing `/`s.
	Backend pulumi.StringOutput `pulumi:"backend"`
	// Specifies the URL values for the CRL Distribution Points field.
	CrlDistributionPoints pulumi.StringArrayOutput `pulumi:"crlDistributionPoints"`
	// Specifies the URL values for the Issuing Certificate field.
	IssuingCertificates pulumi.StringArrayOutput `pulumi:"issuingCertificates"`
	// Specifies the URL values for the OCSP Servers field.
	OcspServers pulumi.StringArrayOutput `pulumi:"ocspServers"`
}

// NewSecretBackendConfigUrls registers a new resource with the given unique name, arguments, and options.
func NewSecretBackendConfigUrls(ctx *pulumi.Context,
	name string, args *SecretBackendConfigUrlsArgs, opts ...pulumi.ResourceOption) (*SecretBackendConfigUrls, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Backend == nil {
		return nil, errors.New("invalid value for required argument 'Backend'")
	}
	var resource SecretBackendConfigUrls
	err := ctx.RegisterResource("vault:pkiSecret/secretBackendConfigUrls:SecretBackendConfigUrls", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSecretBackendConfigUrls gets an existing SecretBackendConfigUrls resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSecretBackendConfigUrls(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SecretBackendConfigUrlsState, opts ...pulumi.ResourceOption) (*SecretBackendConfigUrls, error) {
	var resource SecretBackendConfigUrls
	err := ctx.ReadResource("vault:pkiSecret/secretBackendConfigUrls:SecretBackendConfigUrls", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SecretBackendConfigUrls resources.
type secretBackendConfigUrlsState struct {
	// The path the PKI secret backend is mounted at, with no leading or trailing `/`s.
	Backend *string `pulumi:"backend"`
	// Specifies the URL values for the CRL Distribution Points field.
	CrlDistributionPoints []string `pulumi:"crlDistributionPoints"`
	// Specifies the URL values for the Issuing Certificate field.
	IssuingCertificates []string `pulumi:"issuingCertificates"`
	// Specifies the URL values for the OCSP Servers field.
	OcspServers []string `pulumi:"ocspServers"`
}

type SecretBackendConfigUrlsState struct {
	// The path the PKI secret backend is mounted at, with no leading or trailing `/`s.
	Backend pulumi.StringPtrInput
	// Specifies the URL values for the CRL Distribution Points field.
	CrlDistributionPoints pulumi.StringArrayInput
	// Specifies the URL values for the Issuing Certificate field.
	IssuingCertificates pulumi.StringArrayInput
	// Specifies the URL values for the OCSP Servers field.
	OcspServers pulumi.StringArrayInput
}

func (SecretBackendConfigUrlsState) ElementType() reflect.Type {
	return reflect.TypeOf((*secretBackendConfigUrlsState)(nil)).Elem()
}

type secretBackendConfigUrlsArgs struct {
	// The path the PKI secret backend is mounted at, with no leading or trailing `/`s.
	Backend string `pulumi:"backend"`
	// Specifies the URL values for the CRL Distribution Points field.
	CrlDistributionPoints []string `pulumi:"crlDistributionPoints"`
	// Specifies the URL values for the Issuing Certificate field.
	IssuingCertificates []string `pulumi:"issuingCertificates"`
	// Specifies the URL values for the OCSP Servers field.
	OcspServers []string `pulumi:"ocspServers"`
}

// The set of arguments for constructing a SecretBackendConfigUrls resource.
type SecretBackendConfigUrlsArgs struct {
	// The path the PKI secret backend is mounted at, with no leading or trailing `/`s.
	Backend pulumi.StringInput
	// Specifies the URL values for the CRL Distribution Points field.
	CrlDistributionPoints pulumi.StringArrayInput
	// Specifies the URL values for the Issuing Certificate field.
	IssuingCertificates pulumi.StringArrayInput
	// Specifies the URL values for the OCSP Servers field.
	OcspServers pulumi.StringArrayInput
}

func (SecretBackendConfigUrlsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*secretBackendConfigUrlsArgs)(nil)).Elem()
}

type SecretBackendConfigUrlsInput interface {
	pulumi.Input

	ToSecretBackendConfigUrlsOutput() SecretBackendConfigUrlsOutput
	ToSecretBackendConfigUrlsOutputWithContext(ctx context.Context) SecretBackendConfigUrlsOutput
}

func (*SecretBackendConfigUrls) ElementType() reflect.Type {
	return reflect.TypeOf((**SecretBackendConfigUrls)(nil)).Elem()
}

func (i *SecretBackendConfigUrls) ToSecretBackendConfigUrlsOutput() SecretBackendConfigUrlsOutput {
	return i.ToSecretBackendConfigUrlsOutputWithContext(context.Background())
}

func (i *SecretBackendConfigUrls) ToSecretBackendConfigUrlsOutputWithContext(ctx context.Context) SecretBackendConfigUrlsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecretBackendConfigUrlsOutput)
}

// SecretBackendConfigUrlsArrayInput is an input type that accepts SecretBackendConfigUrlsArray and SecretBackendConfigUrlsArrayOutput values.
// You can construct a concrete instance of `SecretBackendConfigUrlsArrayInput` via:
//
//          SecretBackendConfigUrlsArray{ SecretBackendConfigUrlsArgs{...} }
type SecretBackendConfigUrlsArrayInput interface {
	pulumi.Input

	ToSecretBackendConfigUrlsArrayOutput() SecretBackendConfigUrlsArrayOutput
	ToSecretBackendConfigUrlsArrayOutputWithContext(context.Context) SecretBackendConfigUrlsArrayOutput
}

type SecretBackendConfigUrlsArray []SecretBackendConfigUrlsInput

func (SecretBackendConfigUrlsArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SecretBackendConfigUrls)(nil)).Elem()
}

func (i SecretBackendConfigUrlsArray) ToSecretBackendConfigUrlsArrayOutput() SecretBackendConfigUrlsArrayOutput {
	return i.ToSecretBackendConfigUrlsArrayOutputWithContext(context.Background())
}

func (i SecretBackendConfigUrlsArray) ToSecretBackendConfigUrlsArrayOutputWithContext(ctx context.Context) SecretBackendConfigUrlsArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecretBackendConfigUrlsArrayOutput)
}

// SecretBackendConfigUrlsMapInput is an input type that accepts SecretBackendConfigUrlsMap and SecretBackendConfigUrlsMapOutput values.
// You can construct a concrete instance of `SecretBackendConfigUrlsMapInput` via:
//
//          SecretBackendConfigUrlsMap{ "key": SecretBackendConfigUrlsArgs{...} }
type SecretBackendConfigUrlsMapInput interface {
	pulumi.Input

	ToSecretBackendConfigUrlsMapOutput() SecretBackendConfigUrlsMapOutput
	ToSecretBackendConfigUrlsMapOutputWithContext(context.Context) SecretBackendConfigUrlsMapOutput
}

type SecretBackendConfigUrlsMap map[string]SecretBackendConfigUrlsInput

func (SecretBackendConfigUrlsMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SecretBackendConfigUrls)(nil)).Elem()
}

func (i SecretBackendConfigUrlsMap) ToSecretBackendConfigUrlsMapOutput() SecretBackendConfigUrlsMapOutput {
	return i.ToSecretBackendConfigUrlsMapOutputWithContext(context.Background())
}

func (i SecretBackendConfigUrlsMap) ToSecretBackendConfigUrlsMapOutputWithContext(ctx context.Context) SecretBackendConfigUrlsMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecretBackendConfigUrlsMapOutput)
}

type SecretBackendConfigUrlsOutput struct{ *pulumi.OutputState }

func (SecretBackendConfigUrlsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SecretBackendConfigUrls)(nil)).Elem()
}

func (o SecretBackendConfigUrlsOutput) ToSecretBackendConfigUrlsOutput() SecretBackendConfigUrlsOutput {
	return o
}

func (o SecretBackendConfigUrlsOutput) ToSecretBackendConfigUrlsOutputWithContext(ctx context.Context) SecretBackendConfigUrlsOutput {
	return o
}

type SecretBackendConfigUrlsArrayOutput struct{ *pulumi.OutputState }

func (SecretBackendConfigUrlsArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SecretBackendConfigUrls)(nil)).Elem()
}

func (o SecretBackendConfigUrlsArrayOutput) ToSecretBackendConfigUrlsArrayOutput() SecretBackendConfigUrlsArrayOutput {
	return o
}

func (o SecretBackendConfigUrlsArrayOutput) ToSecretBackendConfigUrlsArrayOutputWithContext(ctx context.Context) SecretBackendConfigUrlsArrayOutput {
	return o
}

func (o SecretBackendConfigUrlsArrayOutput) Index(i pulumi.IntInput) SecretBackendConfigUrlsOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SecretBackendConfigUrls {
		return vs[0].([]*SecretBackendConfigUrls)[vs[1].(int)]
	}).(SecretBackendConfigUrlsOutput)
}

type SecretBackendConfigUrlsMapOutput struct{ *pulumi.OutputState }

func (SecretBackendConfigUrlsMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SecretBackendConfigUrls)(nil)).Elem()
}

func (o SecretBackendConfigUrlsMapOutput) ToSecretBackendConfigUrlsMapOutput() SecretBackendConfigUrlsMapOutput {
	return o
}

func (o SecretBackendConfigUrlsMapOutput) ToSecretBackendConfigUrlsMapOutputWithContext(ctx context.Context) SecretBackendConfigUrlsMapOutput {
	return o
}

func (o SecretBackendConfigUrlsMapOutput) MapIndex(k pulumi.StringInput) SecretBackendConfigUrlsOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SecretBackendConfigUrls {
		return vs[0].(map[string]*SecretBackendConfigUrls)[vs[1].(string)]
	}).(SecretBackendConfigUrlsOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SecretBackendConfigUrlsInput)(nil)).Elem(), &SecretBackendConfigUrls{})
	pulumi.RegisterInputType(reflect.TypeOf((*SecretBackendConfigUrlsArrayInput)(nil)).Elem(), SecretBackendConfigUrlsArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SecretBackendConfigUrlsMapInput)(nil)).Elem(), SecretBackendConfigUrlsMap{})
	pulumi.RegisterOutputType(SecretBackendConfigUrlsOutput{})
	pulumi.RegisterOutputType(SecretBackendConfigUrlsArrayOutput{})
	pulumi.RegisterOutputType(SecretBackendConfigUrlsMapOutput{})
}
