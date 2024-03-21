// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package aws

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-vault/sdk/v6/go/vault/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## Import
//
// AWS secret backends can be imported using the `path`, e.g.
//
// ```sh
// $ pulumi import vault:aws/secretBackend:SecretBackend aws aws
// ```
type SecretBackend struct {
	pulumi.CustomResourceState

	// The AWS Access Key ID this backend should use to
	// issue new credentials. Vault uses the official AWS SDK to authenticate, and thus can also use standard AWS environment credentials, shared file credentials or IAM role/ECS task credentials.
	AccessKey pulumi.StringPtrOutput `pulumi:"accessKey"`
	// The default TTL for credentials
	// issued by this backend.
	DefaultLeaseTtlSeconds pulumi.IntOutput `pulumi:"defaultLeaseTtlSeconds"`
	// A human-friendly description for this backend.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// If set, opts out of mount migration on path updates.
	// See here for more info on [Mount Migration](https://www.vaultproject.io/docs/concepts/mount-migration)
	DisableRemount pulumi.BoolPtrOutput `pulumi:"disableRemount"`
	// Specifies a custom HTTP IAM endpoint to use.
	IamEndpoint pulumi.StringPtrOutput `pulumi:"iamEndpoint"`
	// The audience claim value. Requires Vault 1.16+.
	IdentityTokenAudience pulumi.StringPtrOutput `pulumi:"identityTokenAudience"`
	// The key to use for signing identity tokens. Requires Vault 1.16+.
	IdentityTokenKey pulumi.StringPtrOutput `pulumi:"identityTokenKey"`
	// The TTL of generated identity tokens in seconds. Requires Vault 1.16+.
	IdentityTokenTtl pulumi.IntOutput `pulumi:"identityTokenTtl"`
	// Specifies whether the secrets mount will be marked as local. Local mounts are not replicated to performance replicas.
	Local pulumi.BoolPtrOutput `pulumi:"local"`
	// The maximum TTL that can be requested
	// for credentials issued by this backend.
	MaxLeaseTtlSeconds pulumi.IntOutput `pulumi:"maxLeaseTtlSeconds"`
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
	// *Available only for Vault Enterprise*.
	Namespace pulumi.StringPtrOutput `pulumi:"namespace"`
	// The unique path this backend should be mounted at. Must
	// not begin or end with a `/`. Defaults to `aws`.
	Path pulumi.StringPtrOutput `pulumi:"path"`
	// The AWS region to make API calls against. Defaults to us-east-1.
	Region pulumi.StringOutput `pulumi:"region"`
	// Role ARN to assume for plugin identity token federation. Requires Vault 1.16+.
	RoleArn pulumi.StringPtrOutput `pulumi:"roleArn"`
	// The AWS Secret Access Key to use when generating new credentials.
	SecretKey pulumi.StringPtrOutput `pulumi:"secretKey"`
	// Specifies a custom HTTP STS endpoint to use.
	StsEndpoint pulumi.StringPtrOutput `pulumi:"stsEndpoint"`
	// Template describing how dynamic usernames are generated. The username template is used to generate both IAM usernames (capped at 64 characters) and STS usernames (capped at 32 characters). If no template is provided the field defaults to the template:
	UsernameTemplate pulumi.StringOutput `pulumi:"usernameTemplate"`
}

// NewSecretBackend registers a new resource with the given unique name, arguments, and options.
func NewSecretBackend(ctx *pulumi.Context,
	name string, args *SecretBackendArgs, opts ...pulumi.ResourceOption) (*SecretBackend, error) {
	if args == nil {
		args = &SecretBackendArgs{}
	}

	if args.AccessKey != nil {
		args.AccessKey = pulumi.ToSecret(args.AccessKey).(pulumi.StringPtrInput)
	}
	if args.SecretKey != nil {
		args.SecretKey = pulumi.ToSecret(args.SecretKey).(pulumi.StringPtrInput)
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"accessKey",
		"secretKey",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource SecretBackend
	err := ctx.RegisterResource("vault:aws/secretBackend:SecretBackend", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSecretBackend gets an existing SecretBackend resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSecretBackend(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SecretBackendState, opts ...pulumi.ResourceOption) (*SecretBackend, error) {
	var resource SecretBackend
	err := ctx.ReadResource("vault:aws/secretBackend:SecretBackend", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SecretBackend resources.
type secretBackendState struct {
	// The AWS Access Key ID this backend should use to
	// issue new credentials. Vault uses the official AWS SDK to authenticate, and thus can also use standard AWS environment credentials, shared file credentials or IAM role/ECS task credentials.
	AccessKey *string `pulumi:"accessKey"`
	// The default TTL for credentials
	// issued by this backend.
	DefaultLeaseTtlSeconds *int `pulumi:"defaultLeaseTtlSeconds"`
	// A human-friendly description for this backend.
	Description *string `pulumi:"description"`
	// If set, opts out of mount migration on path updates.
	// See here for more info on [Mount Migration](https://www.vaultproject.io/docs/concepts/mount-migration)
	DisableRemount *bool `pulumi:"disableRemount"`
	// Specifies a custom HTTP IAM endpoint to use.
	IamEndpoint *string `pulumi:"iamEndpoint"`
	// The audience claim value. Requires Vault 1.16+.
	IdentityTokenAudience *string `pulumi:"identityTokenAudience"`
	// The key to use for signing identity tokens. Requires Vault 1.16+.
	IdentityTokenKey *string `pulumi:"identityTokenKey"`
	// The TTL of generated identity tokens in seconds. Requires Vault 1.16+.
	IdentityTokenTtl *int `pulumi:"identityTokenTtl"`
	// Specifies whether the secrets mount will be marked as local. Local mounts are not replicated to performance replicas.
	Local *bool `pulumi:"local"`
	// The maximum TTL that can be requested
	// for credentials issued by this backend.
	MaxLeaseTtlSeconds *int `pulumi:"maxLeaseTtlSeconds"`
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
	// *Available only for Vault Enterprise*.
	Namespace *string `pulumi:"namespace"`
	// The unique path this backend should be mounted at. Must
	// not begin or end with a `/`. Defaults to `aws`.
	Path *string `pulumi:"path"`
	// The AWS region to make API calls against. Defaults to us-east-1.
	Region *string `pulumi:"region"`
	// Role ARN to assume for plugin identity token federation. Requires Vault 1.16+.
	RoleArn *string `pulumi:"roleArn"`
	// The AWS Secret Access Key to use when generating new credentials.
	SecretKey *string `pulumi:"secretKey"`
	// Specifies a custom HTTP STS endpoint to use.
	StsEndpoint *string `pulumi:"stsEndpoint"`
	// Template describing how dynamic usernames are generated. The username template is used to generate both IAM usernames (capped at 64 characters) and STS usernames (capped at 32 characters). If no template is provided the field defaults to the template:
	UsernameTemplate *string `pulumi:"usernameTemplate"`
}

type SecretBackendState struct {
	// The AWS Access Key ID this backend should use to
	// issue new credentials. Vault uses the official AWS SDK to authenticate, and thus can also use standard AWS environment credentials, shared file credentials or IAM role/ECS task credentials.
	AccessKey pulumi.StringPtrInput
	// The default TTL for credentials
	// issued by this backend.
	DefaultLeaseTtlSeconds pulumi.IntPtrInput
	// A human-friendly description for this backend.
	Description pulumi.StringPtrInput
	// If set, opts out of mount migration on path updates.
	// See here for more info on [Mount Migration](https://www.vaultproject.io/docs/concepts/mount-migration)
	DisableRemount pulumi.BoolPtrInput
	// Specifies a custom HTTP IAM endpoint to use.
	IamEndpoint pulumi.StringPtrInput
	// The audience claim value. Requires Vault 1.16+.
	IdentityTokenAudience pulumi.StringPtrInput
	// The key to use for signing identity tokens. Requires Vault 1.16+.
	IdentityTokenKey pulumi.StringPtrInput
	// The TTL of generated identity tokens in seconds. Requires Vault 1.16+.
	IdentityTokenTtl pulumi.IntPtrInput
	// Specifies whether the secrets mount will be marked as local. Local mounts are not replicated to performance replicas.
	Local pulumi.BoolPtrInput
	// The maximum TTL that can be requested
	// for credentials issued by this backend.
	MaxLeaseTtlSeconds pulumi.IntPtrInput
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
	// *Available only for Vault Enterprise*.
	Namespace pulumi.StringPtrInput
	// The unique path this backend should be mounted at. Must
	// not begin or end with a `/`. Defaults to `aws`.
	Path pulumi.StringPtrInput
	// The AWS region to make API calls against. Defaults to us-east-1.
	Region pulumi.StringPtrInput
	// Role ARN to assume for plugin identity token federation. Requires Vault 1.16+.
	RoleArn pulumi.StringPtrInput
	// The AWS Secret Access Key to use when generating new credentials.
	SecretKey pulumi.StringPtrInput
	// Specifies a custom HTTP STS endpoint to use.
	StsEndpoint pulumi.StringPtrInput
	// Template describing how dynamic usernames are generated. The username template is used to generate both IAM usernames (capped at 64 characters) and STS usernames (capped at 32 characters). If no template is provided the field defaults to the template:
	UsernameTemplate pulumi.StringPtrInput
}

func (SecretBackendState) ElementType() reflect.Type {
	return reflect.TypeOf((*secretBackendState)(nil)).Elem()
}

type secretBackendArgs struct {
	// The AWS Access Key ID this backend should use to
	// issue new credentials. Vault uses the official AWS SDK to authenticate, and thus can also use standard AWS environment credentials, shared file credentials or IAM role/ECS task credentials.
	AccessKey *string `pulumi:"accessKey"`
	// The default TTL for credentials
	// issued by this backend.
	DefaultLeaseTtlSeconds *int `pulumi:"defaultLeaseTtlSeconds"`
	// A human-friendly description for this backend.
	Description *string `pulumi:"description"`
	// If set, opts out of mount migration on path updates.
	// See here for more info on [Mount Migration](https://www.vaultproject.io/docs/concepts/mount-migration)
	DisableRemount *bool `pulumi:"disableRemount"`
	// Specifies a custom HTTP IAM endpoint to use.
	IamEndpoint *string `pulumi:"iamEndpoint"`
	// The audience claim value. Requires Vault 1.16+.
	IdentityTokenAudience *string `pulumi:"identityTokenAudience"`
	// The key to use for signing identity tokens. Requires Vault 1.16+.
	IdentityTokenKey *string `pulumi:"identityTokenKey"`
	// The TTL of generated identity tokens in seconds. Requires Vault 1.16+.
	IdentityTokenTtl *int `pulumi:"identityTokenTtl"`
	// Specifies whether the secrets mount will be marked as local. Local mounts are not replicated to performance replicas.
	Local *bool `pulumi:"local"`
	// The maximum TTL that can be requested
	// for credentials issued by this backend.
	MaxLeaseTtlSeconds *int `pulumi:"maxLeaseTtlSeconds"`
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
	// *Available only for Vault Enterprise*.
	Namespace *string `pulumi:"namespace"`
	// The unique path this backend should be mounted at. Must
	// not begin or end with a `/`. Defaults to `aws`.
	Path *string `pulumi:"path"`
	// The AWS region to make API calls against. Defaults to us-east-1.
	Region *string `pulumi:"region"`
	// Role ARN to assume for plugin identity token federation. Requires Vault 1.16+.
	RoleArn *string `pulumi:"roleArn"`
	// The AWS Secret Access Key to use when generating new credentials.
	SecretKey *string `pulumi:"secretKey"`
	// Specifies a custom HTTP STS endpoint to use.
	StsEndpoint *string `pulumi:"stsEndpoint"`
	// Template describing how dynamic usernames are generated. The username template is used to generate both IAM usernames (capped at 64 characters) and STS usernames (capped at 32 characters). If no template is provided the field defaults to the template:
	UsernameTemplate *string `pulumi:"usernameTemplate"`
}

// The set of arguments for constructing a SecretBackend resource.
type SecretBackendArgs struct {
	// The AWS Access Key ID this backend should use to
	// issue new credentials. Vault uses the official AWS SDK to authenticate, and thus can also use standard AWS environment credentials, shared file credentials or IAM role/ECS task credentials.
	AccessKey pulumi.StringPtrInput
	// The default TTL for credentials
	// issued by this backend.
	DefaultLeaseTtlSeconds pulumi.IntPtrInput
	// A human-friendly description for this backend.
	Description pulumi.StringPtrInput
	// If set, opts out of mount migration on path updates.
	// See here for more info on [Mount Migration](https://www.vaultproject.io/docs/concepts/mount-migration)
	DisableRemount pulumi.BoolPtrInput
	// Specifies a custom HTTP IAM endpoint to use.
	IamEndpoint pulumi.StringPtrInput
	// The audience claim value. Requires Vault 1.16+.
	IdentityTokenAudience pulumi.StringPtrInput
	// The key to use for signing identity tokens. Requires Vault 1.16+.
	IdentityTokenKey pulumi.StringPtrInput
	// The TTL of generated identity tokens in seconds. Requires Vault 1.16+.
	IdentityTokenTtl pulumi.IntPtrInput
	// Specifies whether the secrets mount will be marked as local. Local mounts are not replicated to performance replicas.
	Local pulumi.BoolPtrInput
	// The maximum TTL that can be requested
	// for credentials issued by this backend.
	MaxLeaseTtlSeconds pulumi.IntPtrInput
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
	// *Available only for Vault Enterprise*.
	Namespace pulumi.StringPtrInput
	// The unique path this backend should be mounted at. Must
	// not begin or end with a `/`. Defaults to `aws`.
	Path pulumi.StringPtrInput
	// The AWS region to make API calls against. Defaults to us-east-1.
	Region pulumi.StringPtrInput
	// Role ARN to assume for plugin identity token federation. Requires Vault 1.16+.
	RoleArn pulumi.StringPtrInput
	// The AWS Secret Access Key to use when generating new credentials.
	SecretKey pulumi.StringPtrInput
	// Specifies a custom HTTP STS endpoint to use.
	StsEndpoint pulumi.StringPtrInput
	// Template describing how dynamic usernames are generated. The username template is used to generate both IAM usernames (capped at 64 characters) and STS usernames (capped at 32 characters). If no template is provided the field defaults to the template:
	UsernameTemplate pulumi.StringPtrInput
}

func (SecretBackendArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*secretBackendArgs)(nil)).Elem()
}

type SecretBackendInput interface {
	pulumi.Input

	ToSecretBackendOutput() SecretBackendOutput
	ToSecretBackendOutputWithContext(ctx context.Context) SecretBackendOutput
}

func (*SecretBackend) ElementType() reflect.Type {
	return reflect.TypeOf((**SecretBackend)(nil)).Elem()
}

func (i *SecretBackend) ToSecretBackendOutput() SecretBackendOutput {
	return i.ToSecretBackendOutputWithContext(context.Background())
}

func (i *SecretBackend) ToSecretBackendOutputWithContext(ctx context.Context) SecretBackendOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecretBackendOutput)
}

// SecretBackendArrayInput is an input type that accepts SecretBackendArray and SecretBackendArrayOutput values.
// You can construct a concrete instance of `SecretBackendArrayInput` via:
//
//	SecretBackendArray{ SecretBackendArgs{...} }
type SecretBackendArrayInput interface {
	pulumi.Input

	ToSecretBackendArrayOutput() SecretBackendArrayOutput
	ToSecretBackendArrayOutputWithContext(context.Context) SecretBackendArrayOutput
}

type SecretBackendArray []SecretBackendInput

func (SecretBackendArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SecretBackend)(nil)).Elem()
}

func (i SecretBackendArray) ToSecretBackendArrayOutput() SecretBackendArrayOutput {
	return i.ToSecretBackendArrayOutputWithContext(context.Background())
}

func (i SecretBackendArray) ToSecretBackendArrayOutputWithContext(ctx context.Context) SecretBackendArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecretBackendArrayOutput)
}

// SecretBackendMapInput is an input type that accepts SecretBackendMap and SecretBackendMapOutput values.
// You can construct a concrete instance of `SecretBackendMapInput` via:
//
//	SecretBackendMap{ "key": SecretBackendArgs{...} }
type SecretBackendMapInput interface {
	pulumi.Input

	ToSecretBackendMapOutput() SecretBackendMapOutput
	ToSecretBackendMapOutputWithContext(context.Context) SecretBackendMapOutput
}

type SecretBackendMap map[string]SecretBackendInput

func (SecretBackendMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SecretBackend)(nil)).Elem()
}

func (i SecretBackendMap) ToSecretBackendMapOutput() SecretBackendMapOutput {
	return i.ToSecretBackendMapOutputWithContext(context.Background())
}

func (i SecretBackendMap) ToSecretBackendMapOutputWithContext(ctx context.Context) SecretBackendMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecretBackendMapOutput)
}

type SecretBackendOutput struct{ *pulumi.OutputState }

func (SecretBackendOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SecretBackend)(nil)).Elem()
}

func (o SecretBackendOutput) ToSecretBackendOutput() SecretBackendOutput {
	return o
}

func (o SecretBackendOutput) ToSecretBackendOutputWithContext(ctx context.Context) SecretBackendOutput {
	return o
}

// The AWS Access Key ID this backend should use to
// issue new credentials. Vault uses the official AWS SDK to authenticate, and thus can also use standard AWS environment credentials, shared file credentials or IAM role/ECS task credentials.
func (o SecretBackendOutput) AccessKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecretBackend) pulumi.StringPtrOutput { return v.AccessKey }).(pulumi.StringPtrOutput)
}

// The default TTL for credentials
// issued by this backend.
func (o SecretBackendOutput) DefaultLeaseTtlSeconds() pulumi.IntOutput {
	return o.ApplyT(func(v *SecretBackend) pulumi.IntOutput { return v.DefaultLeaseTtlSeconds }).(pulumi.IntOutput)
}

// A human-friendly description for this backend.
func (o SecretBackendOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecretBackend) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// If set, opts out of mount migration on path updates.
// See here for more info on [Mount Migration](https://www.vaultproject.io/docs/concepts/mount-migration)
func (o SecretBackendOutput) DisableRemount() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SecretBackend) pulumi.BoolPtrOutput { return v.DisableRemount }).(pulumi.BoolPtrOutput)
}

// Specifies a custom HTTP IAM endpoint to use.
func (o SecretBackendOutput) IamEndpoint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecretBackend) pulumi.StringPtrOutput { return v.IamEndpoint }).(pulumi.StringPtrOutput)
}

// The audience claim value. Requires Vault 1.16+.
func (o SecretBackendOutput) IdentityTokenAudience() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecretBackend) pulumi.StringPtrOutput { return v.IdentityTokenAudience }).(pulumi.StringPtrOutput)
}

// The key to use for signing identity tokens. Requires Vault 1.16+.
func (o SecretBackendOutput) IdentityTokenKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecretBackend) pulumi.StringPtrOutput { return v.IdentityTokenKey }).(pulumi.StringPtrOutput)
}

// The TTL of generated identity tokens in seconds. Requires Vault 1.16+.
func (o SecretBackendOutput) IdentityTokenTtl() pulumi.IntOutput {
	return o.ApplyT(func(v *SecretBackend) pulumi.IntOutput { return v.IdentityTokenTtl }).(pulumi.IntOutput)
}

// Specifies whether the secrets mount will be marked as local. Local mounts are not replicated to performance replicas.
func (o SecretBackendOutput) Local() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SecretBackend) pulumi.BoolPtrOutput { return v.Local }).(pulumi.BoolPtrOutput)
}

// The maximum TTL that can be requested
// for credentials issued by this backend.
func (o SecretBackendOutput) MaxLeaseTtlSeconds() pulumi.IntOutput {
	return o.ApplyT(func(v *SecretBackend) pulumi.IntOutput { return v.MaxLeaseTtlSeconds }).(pulumi.IntOutput)
}

// The namespace to provision the resource in.
// The value should not contain leading or trailing forward slashes.
// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
// *Available only for Vault Enterprise*.
func (o SecretBackendOutput) Namespace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecretBackend) pulumi.StringPtrOutput { return v.Namespace }).(pulumi.StringPtrOutput)
}

// The unique path this backend should be mounted at. Must
// not begin or end with a `/`. Defaults to `aws`.
func (o SecretBackendOutput) Path() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecretBackend) pulumi.StringPtrOutput { return v.Path }).(pulumi.StringPtrOutput)
}

// The AWS region to make API calls against. Defaults to us-east-1.
func (o SecretBackendOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *SecretBackend) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

// Role ARN to assume for plugin identity token federation. Requires Vault 1.16+.
func (o SecretBackendOutput) RoleArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecretBackend) pulumi.StringPtrOutput { return v.RoleArn }).(pulumi.StringPtrOutput)
}

// The AWS Secret Access Key to use when generating new credentials.
func (o SecretBackendOutput) SecretKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecretBackend) pulumi.StringPtrOutput { return v.SecretKey }).(pulumi.StringPtrOutput)
}

// Specifies a custom HTTP STS endpoint to use.
func (o SecretBackendOutput) StsEndpoint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecretBackend) pulumi.StringPtrOutput { return v.StsEndpoint }).(pulumi.StringPtrOutput)
}

// Template describing how dynamic usernames are generated. The username template is used to generate both IAM usernames (capped at 64 characters) and STS usernames (capped at 32 characters). If no template is provided the field defaults to the template:
func (o SecretBackendOutput) UsernameTemplate() pulumi.StringOutput {
	return o.ApplyT(func(v *SecretBackend) pulumi.StringOutput { return v.UsernameTemplate }).(pulumi.StringOutput)
}

type SecretBackendArrayOutput struct{ *pulumi.OutputState }

func (SecretBackendArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SecretBackend)(nil)).Elem()
}

func (o SecretBackendArrayOutput) ToSecretBackendArrayOutput() SecretBackendArrayOutput {
	return o
}

func (o SecretBackendArrayOutput) ToSecretBackendArrayOutputWithContext(ctx context.Context) SecretBackendArrayOutput {
	return o
}

func (o SecretBackendArrayOutput) Index(i pulumi.IntInput) SecretBackendOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SecretBackend {
		return vs[0].([]*SecretBackend)[vs[1].(int)]
	}).(SecretBackendOutput)
}

type SecretBackendMapOutput struct{ *pulumi.OutputState }

func (SecretBackendMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SecretBackend)(nil)).Elem()
}

func (o SecretBackendMapOutput) ToSecretBackendMapOutput() SecretBackendMapOutput {
	return o
}

func (o SecretBackendMapOutput) ToSecretBackendMapOutputWithContext(ctx context.Context) SecretBackendMapOutput {
	return o
}

func (o SecretBackendMapOutput) MapIndex(k pulumi.StringInput) SecretBackendOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SecretBackend {
		return vs[0].(map[string]*SecretBackend)[vs[1].(string)]
	}).(SecretBackendOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SecretBackendInput)(nil)).Elem(), &SecretBackend{})
	pulumi.RegisterInputType(reflect.TypeOf((*SecretBackendArrayInput)(nil)).Elem(), SecretBackendArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SecretBackendMapInput)(nil)).Elem(), SecretBackendMap{})
	pulumi.RegisterOutputType(SecretBackendOutput{})
	pulumi.RegisterOutputType(SecretBackendArrayOutput{})
	pulumi.RegisterOutputType(SecretBackendMapOutput{})
}
