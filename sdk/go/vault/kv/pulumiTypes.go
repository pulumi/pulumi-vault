// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package kv

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-vault/sdk/v6/go/vault/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

var _ = internal.GetEnvOrDefault

type SecretV2CustomMetadata struct {
	// If true, all keys will require the cas parameter to be set on all write requests.
	CasRequired *bool `pulumi:"casRequired"`
	// A mapping whose keys are the top-level data keys returned from
	// Vault and whose values are the corresponding values. This map can only
	// represent string data, so any non-string values returned from Vault are
	// serialized as JSON.
	Data map[string]string `pulumi:"data"`
	// If set, specifies the length of time before a version is deleted.
	DeleteVersionAfter *int `pulumi:"deleteVersionAfter"`
	// The number of versions to keep per key.
	MaxVersions *int `pulumi:"maxVersions"`
}

// SecretV2CustomMetadataInput is an input type that accepts SecretV2CustomMetadataArgs and SecretV2CustomMetadataOutput values.
// You can construct a concrete instance of `SecretV2CustomMetadataInput` via:
//
//	SecretV2CustomMetadataArgs{...}
type SecretV2CustomMetadataInput interface {
	pulumi.Input

	ToSecretV2CustomMetadataOutput() SecretV2CustomMetadataOutput
	ToSecretV2CustomMetadataOutputWithContext(context.Context) SecretV2CustomMetadataOutput
}

type SecretV2CustomMetadataArgs struct {
	// If true, all keys will require the cas parameter to be set on all write requests.
	CasRequired pulumi.BoolPtrInput `pulumi:"casRequired"`
	// A mapping whose keys are the top-level data keys returned from
	// Vault and whose values are the corresponding values. This map can only
	// represent string data, so any non-string values returned from Vault are
	// serialized as JSON.
	Data pulumi.StringMapInput `pulumi:"data"`
	// If set, specifies the length of time before a version is deleted.
	DeleteVersionAfter pulumi.IntPtrInput `pulumi:"deleteVersionAfter"`
	// The number of versions to keep per key.
	MaxVersions pulumi.IntPtrInput `pulumi:"maxVersions"`
}

func (SecretV2CustomMetadataArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SecretV2CustomMetadata)(nil)).Elem()
}

func (i SecretV2CustomMetadataArgs) ToSecretV2CustomMetadataOutput() SecretV2CustomMetadataOutput {
	return i.ToSecretV2CustomMetadataOutputWithContext(context.Background())
}

func (i SecretV2CustomMetadataArgs) ToSecretV2CustomMetadataOutputWithContext(ctx context.Context) SecretV2CustomMetadataOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecretV2CustomMetadataOutput)
}

func (i SecretV2CustomMetadataArgs) ToSecretV2CustomMetadataPtrOutput() SecretV2CustomMetadataPtrOutput {
	return i.ToSecretV2CustomMetadataPtrOutputWithContext(context.Background())
}

func (i SecretV2CustomMetadataArgs) ToSecretV2CustomMetadataPtrOutputWithContext(ctx context.Context) SecretV2CustomMetadataPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecretV2CustomMetadataOutput).ToSecretV2CustomMetadataPtrOutputWithContext(ctx)
}

// SecretV2CustomMetadataPtrInput is an input type that accepts SecretV2CustomMetadataArgs, SecretV2CustomMetadataPtr and SecretV2CustomMetadataPtrOutput values.
// You can construct a concrete instance of `SecretV2CustomMetadataPtrInput` via:
//
//	        SecretV2CustomMetadataArgs{...}
//
//	or:
//
//	        nil
type SecretV2CustomMetadataPtrInput interface {
	pulumi.Input

	ToSecretV2CustomMetadataPtrOutput() SecretV2CustomMetadataPtrOutput
	ToSecretV2CustomMetadataPtrOutputWithContext(context.Context) SecretV2CustomMetadataPtrOutput
}

type secretV2CustomMetadataPtrType SecretV2CustomMetadataArgs

func SecretV2CustomMetadataPtr(v *SecretV2CustomMetadataArgs) SecretV2CustomMetadataPtrInput {
	return (*secretV2CustomMetadataPtrType)(v)
}

func (*secretV2CustomMetadataPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SecretV2CustomMetadata)(nil)).Elem()
}

func (i *secretV2CustomMetadataPtrType) ToSecretV2CustomMetadataPtrOutput() SecretV2CustomMetadataPtrOutput {
	return i.ToSecretV2CustomMetadataPtrOutputWithContext(context.Background())
}

func (i *secretV2CustomMetadataPtrType) ToSecretV2CustomMetadataPtrOutputWithContext(ctx context.Context) SecretV2CustomMetadataPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecretV2CustomMetadataPtrOutput)
}

type SecretV2CustomMetadataOutput struct{ *pulumi.OutputState }

func (SecretV2CustomMetadataOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SecretV2CustomMetadata)(nil)).Elem()
}

func (o SecretV2CustomMetadataOutput) ToSecretV2CustomMetadataOutput() SecretV2CustomMetadataOutput {
	return o
}

func (o SecretV2CustomMetadataOutput) ToSecretV2CustomMetadataOutputWithContext(ctx context.Context) SecretV2CustomMetadataOutput {
	return o
}

func (o SecretV2CustomMetadataOutput) ToSecretV2CustomMetadataPtrOutput() SecretV2CustomMetadataPtrOutput {
	return o.ToSecretV2CustomMetadataPtrOutputWithContext(context.Background())
}

func (o SecretV2CustomMetadataOutput) ToSecretV2CustomMetadataPtrOutputWithContext(ctx context.Context) SecretV2CustomMetadataPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SecretV2CustomMetadata) *SecretV2CustomMetadata {
		return &v
	}).(SecretV2CustomMetadataPtrOutput)
}

// If true, all keys will require the cas parameter to be set on all write requests.
func (o SecretV2CustomMetadataOutput) CasRequired() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v SecretV2CustomMetadata) *bool { return v.CasRequired }).(pulumi.BoolPtrOutput)
}

// A mapping whose keys are the top-level data keys returned from
// Vault and whose values are the corresponding values. This map can only
// represent string data, so any non-string values returned from Vault are
// serialized as JSON.
func (o SecretV2CustomMetadataOutput) Data() pulumi.StringMapOutput {
	return o.ApplyT(func(v SecretV2CustomMetadata) map[string]string { return v.Data }).(pulumi.StringMapOutput)
}

// If set, specifies the length of time before a version is deleted.
func (o SecretV2CustomMetadataOutput) DeleteVersionAfter() pulumi.IntPtrOutput {
	return o.ApplyT(func(v SecretV2CustomMetadata) *int { return v.DeleteVersionAfter }).(pulumi.IntPtrOutput)
}

// The number of versions to keep per key.
func (o SecretV2CustomMetadataOutput) MaxVersions() pulumi.IntPtrOutput {
	return o.ApplyT(func(v SecretV2CustomMetadata) *int { return v.MaxVersions }).(pulumi.IntPtrOutput)
}

type SecretV2CustomMetadataPtrOutput struct{ *pulumi.OutputState }

func (SecretV2CustomMetadataPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SecretV2CustomMetadata)(nil)).Elem()
}

func (o SecretV2CustomMetadataPtrOutput) ToSecretV2CustomMetadataPtrOutput() SecretV2CustomMetadataPtrOutput {
	return o
}

func (o SecretV2CustomMetadataPtrOutput) ToSecretV2CustomMetadataPtrOutputWithContext(ctx context.Context) SecretV2CustomMetadataPtrOutput {
	return o
}

func (o SecretV2CustomMetadataPtrOutput) Elem() SecretV2CustomMetadataOutput {
	return o.ApplyT(func(v *SecretV2CustomMetadata) SecretV2CustomMetadata {
		if v != nil {
			return *v
		}
		var ret SecretV2CustomMetadata
		return ret
	}).(SecretV2CustomMetadataOutput)
}

// If true, all keys will require the cas parameter to be set on all write requests.
func (o SecretV2CustomMetadataPtrOutput) CasRequired() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SecretV2CustomMetadata) *bool {
		if v == nil {
			return nil
		}
		return v.CasRequired
	}).(pulumi.BoolPtrOutput)
}

// A mapping whose keys are the top-level data keys returned from
// Vault and whose values are the corresponding values. This map can only
// represent string data, so any non-string values returned from Vault are
// serialized as JSON.
func (o SecretV2CustomMetadataPtrOutput) Data() pulumi.StringMapOutput {
	return o.ApplyT(func(v *SecretV2CustomMetadata) map[string]string {
		if v == nil {
			return nil
		}
		return v.Data
	}).(pulumi.StringMapOutput)
}

// If set, specifies the length of time before a version is deleted.
func (o SecretV2CustomMetadataPtrOutput) DeleteVersionAfter() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SecretV2CustomMetadata) *int {
		if v == nil {
			return nil
		}
		return v.DeleteVersionAfter
	}).(pulumi.IntPtrOutput)
}

// The number of versions to keep per key.
func (o SecretV2CustomMetadataPtrOutput) MaxVersions() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SecretV2CustomMetadata) *int {
		if v == nil {
			return nil
		}
		return v.MaxVersions
	}).(pulumi.IntPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SecretV2CustomMetadataInput)(nil)).Elem(), SecretV2CustomMetadataArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SecretV2CustomMetadataPtrInput)(nil)).Elem(), SecretV2CustomMetadataArgs{})
	pulumi.RegisterOutputType(SecretV2CustomMetadataOutput{})
	pulumi.RegisterOutputType(SecretV2CustomMetadataPtrOutput{})
}
