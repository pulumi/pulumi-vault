// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ssh

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-vault/sdk/v6/go/vault/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

var _ = internal.GetEnvOrDefault

type SecretBackendRoleAllowedUserKeyConfig struct {
	// List of allowed key lengths, vault-1.10 and above
	Lengths []int `pulumi:"lengths"`
	// Key type, choices:
	// rsa, ecdsa, ec, dsa, ed25519, ssh-rsa, ssh-dss, ssh-ed25519, ecdsa-sha2-nistp256, ecdsa-sha2-nistp384, ecdsa-sha2-nistp521
	Type string `pulumi:"type"`
}

// SecretBackendRoleAllowedUserKeyConfigInput is an input type that accepts SecretBackendRoleAllowedUserKeyConfigArgs and SecretBackendRoleAllowedUserKeyConfigOutput values.
// You can construct a concrete instance of `SecretBackendRoleAllowedUserKeyConfigInput` via:
//
//	SecretBackendRoleAllowedUserKeyConfigArgs{...}
type SecretBackendRoleAllowedUserKeyConfigInput interface {
	pulumi.Input

	ToSecretBackendRoleAllowedUserKeyConfigOutput() SecretBackendRoleAllowedUserKeyConfigOutput
	ToSecretBackendRoleAllowedUserKeyConfigOutputWithContext(context.Context) SecretBackendRoleAllowedUserKeyConfigOutput
}

type SecretBackendRoleAllowedUserKeyConfigArgs struct {
	// List of allowed key lengths, vault-1.10 and above
	Lengths pulumi.IntArrayInput `pulumi:"lengths"`
	// Key type, choices:
	// rsa, ecdsa, ec, dsa, ed25519, ssh-rsa, ssh-dss, ssh-ed25519, ecdsa-sha2-nistp256, ecdsa-sha2-nistp384, ecdsa-sha2-nistp521
	Type pulumi.StringInput `pulumi:"type"`
}

func (SecretBackendRoleAllowedUserKeyConfigArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SecretBackendRoleAllowedUserKeyConfig)(nil)).Elem()
}

func (i SecretBackendRoleAllowedUserKeyConfigArgs) ToSecretBackendRoleAllowedUserKeyConfigOutput() SecretBackendRoleAllowedUserKeyConfigOutput {
	return i.ToSecretBackendRoleAllowedUserKeyConfigOutputWithContext(context.Background())
}

func (i SecretBackendRoleAllowedUserKeyConfigArgs) ToSecretBackendRoleAllowedUserKeyConfigOutputWithContext(ctx context.Context) SecretBackendRoleAllowedUserKeyConfigOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecretBackendRoleAllowedUserKeyConfigOutput)
}

// SecretBackendRoleAllowedUserKeyConfigArrayInput is an input type that accepts SecretBackendRoleAllowedUserKeyConfigArray and SecretBackendRoleAllowedUserKeyConfigArrayOutput values.
// You can construct a concrete instance of `SecretBackendRoleAllowedUserKeyConfigArrayInput` via:
//
//	SecretBackendRoleAllowedUserKeyConfigArray{ SecretBackendRoleAllowedUserKeyConfigArgs{...} }
type SecretBackendRoleAllowedUserKeyConfigArrayInput interface {
	pulumi.Input

	ToSecretBackendRoleAllowedUserKeyConfigArrayOutput() SecretBackendRoleAllowedUserKeyConfigArrayOutput
	ToSecretBackendRoleAllowedUserKeyConfigArrayOutputWithContext(context.Context) SecretBackendRoleAllowedUserKeyConfigArrayOutput
}

type SecretBackendRoleAllowedUserKeyConfigArray []SecretBackendRoleAllowedUserKeyConfigInput

func (SecretBackendRoleAllowedUserKeyConfigArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SecretBackendRoleAllowedUserKeyConfig)(nil)).Elem()
}

func (i SecretBackendRoleAllowedUserKeyConfigArray) ToSecretBackendRoleAllowedUserKeyConfigArrayOutput() SecretBackendRoleAllowedUserKeyConfigArrayOutput {
	return i.ToSecretBackendRoleAllowedUserKeyConfigArrayOutputWithContext(context.Background())
}

func (i SecretBackendRoleAllowedUserKeyConfigArray) ToSecretBackendRoleAllowedUserKeyConfigArrayOutputWithContext(ctx context.Context) SecretBackendRoleAllowedUserKeyConfigArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecretBackendRoleAllowedUserKeyConfigArrayOutput)
}

type SecretBackendRoleAllowedUserKeyConfigOutput struct{ *pulumi.OutputState }

func (SecretBackendRoleAllowedUserKeyConfigOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SecretBackendRoleAllowedUserKeyConfig)(nil)).Elem()
}

func (o SecretBackendRoleAllowedUserKeyConfigOutput) ToSecretBackendRoleAllowedUserKeyConfigOutput() SecretBackendRoleAllowedUserKeyConfigOutput {
	return o
}

func (o SecretBackendRoleAllowedUserKeyConfigOutput) ToSecretBackendRoleAllowedUserKeyConfigOutputWithContext(ctx context.Context) SecretBackendRoleAllowedUserKeyConfigOutput {
	return o
}

// List of allowed key lengths, vault-1.10 and above
func (o SecretBackendRoleAllowedUserKeyConfigOutput) Lengths() pulumi.IntArrayOutput {
	return o.ApplyT(func(v SecretBackendRoleAllowedUserKeyConfig) []int { return v.Lengths }).(pulumi.IntArrayOutput)
}

// Key type, choices:
// rsa, ecdsa, ec, dsa, ed25519, ssh-rsa, ssh-dss, ssh-ed25519, ecdsa-sha2-nistp256, ecdsa-sha2-nistp384, ecdsa-sha2-nistp521
func (o SecretBackendRoleAllowedUserKeyConfigOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v SecretBackendRoleAllowedUserKeyConfig) string { return v.Type }).(pulumi.StringOutput)
}

type SecretBackendRoleAllowedUserKeyConfigArrayOutput struct{ *pulumi.OutputState }

func (SecretBackendRoleAllowedUserKeyConfigArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SecretBackendRoleAllowedUserKeyConfig)(nil)).Elem()
}

func (o SecretBackendRoleAllowedUserKeyConfigArrayOutput) ToSecretBackendRoleAllowedUserKeyConfigArrayOutput() SecretBackendRoleAllowedUserKeyConfigArrayOutput {
	return o
}

func (o SecretBackendRoleAllowedUserKeyConfigArrayOutput) ToSecretBackendRoleAllowedUserKeyConfigArrayOutputWithContext(ctx context.Context) SecretBackendRoleAllowedUserKeyConfigArrayOutput {
	return o
}

func (o SecretBackendRoleAllowedUserKeyConfigArrayOutput) Index(i pulumi.IntInput) SecretBackendRoleAllowedUserKeyConfigOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SecretBackendRoleAllowedUserKeyConfig {
		return vs[0].([]SecretBackendRoleAllowedUserKeyConfig)[vs[1].(int)]
	}).(SecretBackendRoleAllowedUserKeyConfigOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SecretBackendRoleAllowedUserKeyConfigInput)(nil)).Elem(), SecretBackendRoleAllowedUserKeyConfigArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SecretBackendRoleAllowedUserKeyConfigArrayInput)(nil)).Elem(), SecretBackendRoleAllowedUserKeyConfigArray{})
	pulumi.RegisterOutputType(SecretBackendRoleAllowedUserKeyConfigOutput{})
	pulumi.RegisterOutputType(SecretBackendRoleAllowedUserKeyConfigArrayOutput{})
}
