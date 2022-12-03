// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ssh

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a resource to manage CA information in an SSH secret backend
// [SSH secret backend within Vault](https://www.vaultproject.io/docs/secrets/ssh/index.html).
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-vault/sdk/v5/go/vault"
// 	"github.com/pulumi/pulumi-vault/sdk/v5/go/vault/ssh"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		example, err := vault.NewMount(ctx, "example", &vault.MountArgs{
// 			Type: pulumi.String("ssh"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = ssh.NewSecretBackendCa(ctx, "foo", &ssh.SecretBackendCaArgs{
// 			Backend: example.Path,
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
// SSH secret backend CAs can be imported using the `path`, e.g.
//
// ```sh
//  $ pulumi import vault:ssh/secretBackendCa:SecretBackendCa foo ssh
// ```
type SecretBackendCa struct {
	pulumi.CustomResourceState

	// The path where the SSH secret backend is mounted. Defaults to 'ssh'
	Backend pulumi.StringPtrOutput `pulumi:"backend"`
	// Whether Vault should generate the signing key pair internally. Defaults to true
	GenerateSigningKey pulumi.BoolPtrOutput `pulumi:"generateSigningKey"`
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
	// *Available only for Vault Enterprise*.
	Namespace pulumi.StringPtrOutput `pulumi:"namespace"`
	// The private key part the SSH CA key pair; required if generateSigningKey is false.
	PrivateKey pulumi.StringOutput `pulumi:"privateKey"`
	// The public key part the SSH CA key pair; required if generateSigningKey is false.
	PublicKey pulumi.StringOutput `pulumi:"publicKey"`
}

// NewSecretBackendCa registers a new resource with the given unique name, arguments, and options.
func NewSecretBackendCa(ctx *pulumi.Context,
	name string, args *SecretBackendCaArgs, opts ...pulumi.ResourceOption) (*SecretBackendCa, error) {
	if args == nil {
		args = &SecretBackendCaArgs{}
	}

	if args.PrivateKey != nil {
		args.PrivateKey = pulumi.ToSecret(args.PrivateKey).(pulumi.StringPtrOutput)
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"privateKey",
	})
	opts = append(opts, secrets)
	var resource SecretBackendCa
	err := ctx.RegisterResource("vault:ssh/secretBackendCa:SecretBackendCa", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSecretBackendCa gets an existing SecretBackendCa resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSecretBackendCa(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SecretBackendCaState, opts ...pulumi.ResourceOption) (*SecretBackendCa, error) {
	var resource SecretBackendCa
	err := ctx.ReadResource("vault:ssh/secretBackendCa:SecretBackendCa", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SecretBackendCa resources.
type secretBackendCaState struct {
	// The path where the SSH secret backend is mounted. Defaults to 'ssh'
	Backend *string `pulumi:"backend"`
	// Whether Vault should generate the signing key pair internally. Defaults to true
	GenerateSigningKey *bool `pulumi:"generateSigningKey"`
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
	// *Available only for Vault Enterprise*.
	Namespace *string `pulumi:"namespace"`
	// The private key part the SSH CA key pair; required if generateSigningKey is false.
	PrivateKey *string `pulumi:"privateKey"`
	// The public key part the SSH CA key pair; required if generateSigningKey is false.
	PublicKey *string `pulumi:"publicKey"`
}

type SecretBackendCaState struct {
	// The path where the SSH secret backend is mounted. Defaults to 'ssh'
	Backend pulumi.StringPtrInput
	// Whether Vault should generate the signing key pair internally. Defaults to true
	GenerateSigningKey pulumi.BoolPtrInput
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
	// *Available only for Vault Enterprise*.
	Namespace pulumi.StringPtrInput
	// The private key part the SSH CA key pair; required if generateSigningKey is false.
	PrivateKey pulumi.StringPtrInput
	// The public key part the SSH CA key pair; required if generateSigningKey is false.
	PublicKey pulumi.StringPtrInput
}

func (SecretBackendCaState) ElementType() reflect.Type {
	return reflect.TypeOf((*secretBackendCaState)(nil)).Elem()
}

type secretBackendCaArgs struct {
	// The path where the SSH secret backend is mounted. Defaults to 'ssh'
	Backend *string `pulumi:"backend"`
	// Whether Vault should generate the signing key pair internally. Defaults to true
	GenerateSigningKey *bool `pulumi:"generateSigningKey"`
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
	// *Available only for Vault Enterprise*.
	Namespace *string `pulumi:"namespace"`
	// The private key part the SSH CA key pair; required if generateSigningKey is false.
	PrivateKey *string `pulumi:"privateKey"`
	// The public key part the SSH CA key pair; required if generateSigningKey is false.
	PublicKey *string `pulumi:"publicKey"`
}

// The set of arguments for constructing a SecretBackendCa resource.
type SecretBackendCaArgs struct {
	// The path where the SSH secret backend is mounted. Defaults to 'ssh'
	Backend pulumi.StringPtrInput
	// Whether Vault should generate the signing key pair internally. Defaults to true
	GenerateSigningKey pulumi.BoolPtrInput
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
	// *Available only for Vault Enterprise*.
	Namespace pulumi.StringPtrInput
	// The private key part the SSH CA key pair; required if generateSigningKey is false.
	PrivateKey pulumi.StringPtrInput
	// The public key part the SSH CA key pair; required if generateSigningKey is false.
	PublicKey pulumi.StringPtrInput
}

func (SecretBackendCaArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*secretBackendCaArgs)(nil)).Elem()
}

type SecretBackendCaInput interface {
	pulumi.Input

	ToSecretBackendCaOutput() SecretBackendCaOutput
	ToSecretBackendCaOutputWithContext(ctx context.Context) SecretBackendCaOutput
}

func (*SecretBackendCa) ElementType() reflect.Type {
	return reflect.TypeOf((**SecretBackendCa)(nil)).Elem()
}

func (i *SecretBackendCa) ToSecretBackendCaOutput() SecretBackendCaOutput {
	return i.ToSecretBackendCaOutputWithContext(context.Background())
}

func (i *SecretBackendCa) ToSecretBackendCaOutputWithContext(ctx context.Context) SecretBackendCaOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecretBackendCaOutput)
}

// SecretBackendCaArrayInput is an input type that accepts SecretBackendCaArray and SecretBackendCaArrayOutput values.
// You can construct a concrete instance of `SecretBackendCaArrayInput` via:
//
//          SecretBackendCaArray{ SecretBackendCaArgs{...} }
type SecretBackendCaArrayInput interface {
	pulumi.Input

	ToSecretBackendCaArrayOutput() SecretBackendCaArrayOutput
	ToSecretBackendCaArrayOutputWithContext(context.Context) SecretBackendCaArrayOutput
}

type SecretBackendCaArray []SecretBackendCaInput

func (SecretBackendCaArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SecretBackendCa)(nil)).Elem()
}

func (i SecretBackendCaArray) ToSecretBackendCaArrayOutput() SecretBackendCaArrayOutput {
	return i.ToSecretBackendCaArrayOutputWithContext(context.Background())
}

func (i SecretBackendCaArray) ToSecretBackendCaArrayOutputWithContext(ctx context.Context) SecretBackendCaArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecretBackendCaArrayOutput)
}

// SecretBackendCaMapInput is an input type that accepts SecretBackendCaMap and SecretBackendCaMapOutput values.
// You can construct a concrete instance of `SecretBackendCaMapInput` via:
//
//          SecretBackendCaMap{ "key": SecretBackendCaArgs{...} }
type SecretBackendCaMapInput interface {
	pulumi.Input

	ToSecretBackendCaMapOutput() SecretBackendCaMapOutput
	ToSecretBackendCaMapOutputWithContext(context.Context) SecretBackendCaMapOutput
}

type SecretBackendCaMap map[string]SecretBackendCaInput

func (SecretBackendCaMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SecretBackendCa)(nil)).Elem()
}

func (i SecretBackendCaMap) ToSecretBackendCaMapOutput() SecretBackendCaMapOutput {
	return i.ToSecretBackendCaMapOutputWithContext(context.Background())
}

func (i SecretBackendCaMap) ToSecretBackendCaMapOutputWithContext(ctx context.Context) SecretBackendCaMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecretBackendCaMapOutput)
}

type SecretBackendCaOutput struct{ *pulumi.OutputState }

func (SecretBackendCaOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SecretBackendCa)(nil)).Elem()
}

func (o SecretBackendCaOutput) ToSecretBackendCaOutput() SecretBackendCaOutput {
	return o
}

func (o SecretBackendCaOutput) ToSecretBackendCaOutputWithContext(ctx context.Context) SecretBackendCaOutput {
	return o
}

// The path where the SSH secret backend is mounted. Defaults to 'ssh'
func (o SecretBackendCaOutput) Backend() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecretBackendCa) pulumi.StringPtrOutput { return v.Backend }).(pulumi.StringPtrOutput)
}

// Whether Vault should generate the signing key pair internally. Defaults to true
func (o SecretBackendCaOutput) GenerateSigningKey() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SecretBackendCa) pulumi.BoolPtrOutput { return v.GenerateSigningKey }).(pulumi.BoolPtrOutput)
}

// The namespace to provision the resource in.
// The value should not contain leading or trailing forward slashes.
// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
// *Available only for Vault Enterprise*.
func (o SecretBackendCaOutput) Namespace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecretBackendCa) pulumi.StringPtrOutput { return v.Namespace }).(pulumi.StringPtrOutput)
}

// The private key part the SSH CA key pair; required if generateSigningKey is false.
func (o SecretBackendCaOutput) PrivateKey() pulumi.StringOutput {
	return o.ApplyT(func(v *SecretBackendCa) pulumi.StringOutput { return v.PrivateKey }).(pulumi.StringOutput)
}

// The public key part the SSH CA key pair; required if generateSigningKey is false.
func (o SecretBackendCaOutput) PublicKey() pulumi.StringOutput {
	return o.ApplyT(func(v *SecretBackendCa) pulumi.StringOutput { return v.PublicKey }).(pulumi.StringOutput)
}

type SecretBackendCaArrayOutput struct{ *pulumi.OutputState }

func (SecretBackendCaArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SecretBackendCa)(nil)).Elem()
}

func (o SecretBackendCaArrayOutput) ToSecretBackendCaArrayOutput() SecretBackendCaArrayOutput {
	return o
}

func (o SecretBackendCaArrayOutput) ToSecretBackendCaArrayOutputWithContext(ctx context.Context) SecretBackendCaArrayOutput {
	return o
}

func (o SecretBackendCaArrayOutput) Index(i pulumi.IntInput) SecretBackendCaOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SecretBackendCa {
		return vs[0].([]*SecretBackendCa)[vs[1].(int)]
	}).(SecretBackendCaOutput)
}

type SecretBackendCaMapOutput struct{ *pulumi.OutputState }

func (SecretBackendCaMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SecretBackendCa)(nil)).Elem()
}

func (o SecretBackendCaMapOutput) ToSecretBackendCaMapOutput() SecretBackendCaMapOutput {
	return o
}

func (o SecretBackendCaMapOutput) ToSecretBackendCaMapOutputWithContext(ctx context.Context) SecretBackendCaMapOutput {
	return o
}

func (o SecretBackendCaMapOutput) MapIndex(k pulumi.StringInput) SecretBackendCaOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SecretBackendCa {
		return vs[0].(map[string]*SecretBackendCa)[vs[1].(string)]
	}).(SecretBackendCaOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SecretBackendCaInput)(nil)).Elem(), &SecretBackendCa{})
	pulumi.RegisterInputType(reflect.TypeOf((*SecretBackendCaArrayInput)(nil)).Elem(), SecretBackendCaArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SecretBackendCaMapInput)(nil)).Elem(), SecretBackendCaMap{})
	pulumi.RegisterOutputType(SecretBackendCaOutput{})
	pulumi.RegisterOutputType(SecretBackendCaArrayOutput{})
	pulumi.RegisterOutputType(SecretBackendCaMapOutput{})
}
