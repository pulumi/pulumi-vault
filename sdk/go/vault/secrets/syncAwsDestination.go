// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package secrets

import (
	"context"
	"reflect"

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
//	"github.com/pulumi/pulumi-vault/sdk/v6/go/vault/secrets"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := secrets.NewSyncAwsDestination(ctx, "aws", &secrets.SyncAwsDestinationArgs{
//				Name:               pulumi.String("aws-dest"),
//				AccessKeyId:        pulumi.Any(accessKeyId),
//				SecretAccessKey:    pulumi.Any(secretAccessKey),
//				Region:             pulumi.String("us-east-1"),
//				RoleArn:            pulumi.String("role-arn"),
//				ExternalId:         pulumi.String("external-id"),
//				SecretNameTemplate: pulumi.String("vault_{{ .MountAccessor | lowercase }}_{{ .SecretPath | lowercase }}"),
//				CustomTags: pulumi.StringMap{
//					"foo": pulumi.String("bar"),
//				},
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
// AWS Secrets sync destinations can be imported using the `name`, e.g.
//
// ```sh
// $ pulumi import vault:secrets/syncAwsDestination:SyncAwsDestination aws aws-dest
// ```
type SyncAwsDestination struct {
	pulumi.CustomResourceState

	// Access key id to authenticate against the AWS secrets manager.
	// Can be omitted and directly provided to Vault using the `AWS_ACCESS_KEY_ID` environment
	// variable.
	AccessKeyId pulumi.StringPtrOutput `pulumi:"accessKeyId"`
	// Custom tags to set on the secret managed at the destination.
	CustomTags pulumi.StringMapOutput `pulumi:"customTags"`
	// Optional extra protection that must match the trust policy granting access to the
	// AWS IAM role ARN. We recommend using a different random UUID per destination. The value is generated by users.
	// The field is mutable with no special condition, but users must be careful that the new value fits with the trust
	// relationship condition they set on AWS otherwise sync operations will start to fail due to client-side access
	// denied errors. Ignored if the `roleArn` field is empty.
	ExternalId pulumi.StringPtrOutput `pulumi:"externalId"`
	// Determines what level of information is synced as a distinct resource
	// at the destination. Supports `secret-path` and `secret-key`.
	Granularity pulumi.StringPtrOutput `pulumi:"granularity"`
	// Unique name of the AWS destination.
	Name pulumi.StringOutput `pulumi:"name"`
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
	Namespace pulumi.StringPtrOutput `pulumi:"namespace"`
	// Region where to manage the secrets manager entries.
	// Can be omitted and directly provided to Vault using the `AWS_REGION` environment
	// variable.
	Region pulumi.StringPtrOutput `pulumi:"region"`
	// Specifies a role to assume when connecting to AWS. When assuming a role,
	// Vault uses temporary STS credentials to authenticate. An initial session with the proper trust relationship must
	// exist for Vault to be able to assume this role. The role can be in a different account.
	// The value is mutable as long as the new role targets the same AWS account ID. If not, the BE will return an error.
	// It is possible to provide both an access key pair and a role to assume.
	RoleArn pulumi.StringPtrOutput `pulumi:"roleArn"`
	// Secret access key to authenticate against the AWS secrets manager.
	// Can be omitted and directly provided to Vault using the `AWS_SECRET_ACCESS_KEY` environment
	// variable.
	SecretAccessKey pulumi.StringPtrOutput `pulumi:"secretAccessKey"`
	// Template describing how to generate external secret names.
	// Supports a subset of the Go Template syntax.
	SecretNameTemplate pulumi.StringOutput `pulumi:"secretNameTemplate"`
	// The type of the secrets destination (`aws-sm`).
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewSyncAwsDestination registers a new resource with the given unique name, arguments, and options.
func NewSyncAwsDestination(ctx *pulumi.Context,
	name string, args *SyncAwsDestinationArgs, opts ...pulumi.ResourceOption) (*SyncAwsDestination, error) {
	if args == nil {
		args = &SyncAwsDestinationArgs{}
	}

	if args.SecretAccessKey != nil {
		args.SecretAccessKey = pulumi.ToSecret(args.SecretAccessKey).(pulumi.StringPtrInput)
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"secretAccessKey",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource SyncAwsDestination
	err := ctx.RegisterResource("vault:secrets/syncAwsDestination:SyncAwsDestination", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSyncAwsDestination gets an existing SyncAwsDestination resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSyncAwsDestination(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SyncAwsDestinationState, opts ...pulumi.ResourceOption) (*SyncAwsDestination, error) {
	var resource SyncAwsDestination
	err := ctx.ReadResource("vault:secrets/syncAwsDestination:SyncAwsDestination", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SyncAwsDestination resources.
type syncAwsDestinationState struct {
	// Access key id to authenticate against the AWS secrets manager.
	// Can be omitted and directly provided to Vault using the `AWS_ACCESS_KEY_ID` environment
	// variable.
	AccessKeyId *string `pulumi:"accessKeyId"`
	// Custom tags to set on the secret managed at the destination.
	CustomTags map[string]string `pulumi:"customTags"`
	// Optional extra protection that must match the trust policy granting access to the
	// AWS IAM role ARN. We recommend using a different random UUID per destination. The value is generated by users.
	// The field is mutable with no special condition, but users must be careful that the new value fits with the trust
	// relationship condition they set on AWS otherwise sync operations will start to fail due to client-side access
	// denied errors. Ignored if the `roleArn` field is empty.
	ExternalId *string `pulumi:"externalId"`
	// Determines what level of information is synced as a distinct resource
	// at the destination. Supports `secret-path` and `secret-key`.
	Granularity *string `pulumi:"granularity"`
	// Unique name of the AWS destination.
	Name *string `pulumi:"name"`
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
	Namespace *string `pulumi:"namespace"`
	// Region where to manage the secrets manager entries.
	// Can be omitted and directly provided to Vault using the `AWS_REGION` environment
	// variable.
	Region *string `pulumi:"region"`
	// Specifies a role to assume when connecting to AWS. When assuming a role,
	// Vault uses temporary STS credentials to authenticate. An initial session with the proper trust relationship must
	// exist for Vault to be able to assume this role. The role can be in a different account.
	// The value is mutable as long as the new role targets the same AWS account ID. If not, the BE will return an error.
	// It is possible to provide both an access key pair and a role to assume.
	RoleArn *string `pulumi:"roleArn"`
	// Secret access key to authenticate against the AWS secrets manager.
	// Can be omitted and directly provided to Vault using the `AWS_SECRET_ACCESS_KEY` environment
	// variable.
	SecretAccessKey *string `pulumi:"secretAccessKey"`
	// Template describing how to generate external secret names.
	// Supports a subset of the Go Template syntax.
	SecretNameTemplate *string `pulumi:"secretNameTemplate"`
	// The type of the secrets destination (`aws-sm`).
	Type *string `pulumi:"type"`
}

type SyncAwsDestinationState struct {
	// Access key id to authenticate against the AWS secrets manager.
	// Can be omitted and directly provided to Vault using the `AWS_ACCESS_KEY_ID` environment
	// variable.
	AccessKeyId pulumi.StringPtrInput
	// Custom tags to set on the secret managed at the destination.
	CustomTags pulumi.StringMapInput
	// Optional extra protection that must match the trust policy granting access to the
	// AWS IAM role ARN. We recommend using a different random UUID per destination. The value is generated by users.
	// The field is mutable with no special condition, but users must be careful that the new value fits with the trust
	// relationship condition they set on AWS otherwise sync operations will start to fail due to client-side access
	// denied errors. Ignored if the `roleArn` field is empty.
	ExternalId pulumi.StringPtrInput
	// Determines what level of information is synced as a distinct resource
	// at the destination. Supports `secret-path` and `secret-key`.
	Granularity pulumi.StringPtrInput
	// Unique name of the AWS destination.
	Name pulumi.StringPtrInput
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
	Namespace pulumi.StringPtrInput
	// Region where to manage the secrets manager entries.
	// Can be omitted and directly provided to Vault using the `AWS_REGION` environment
	// variable.
	Region pulumi.StringPtrInput
	// Specifies a role to assume when connecting to AWS. When assuming a role,
	// Vault uses temporary STS credentials to authenticate. An initial session with the proper trust relationship must
	// exist for Vault to be able to assume this role. The role can be in a different account.
	// The value is mutable as long as the new role targets the same AWS account ID. If not, the BE will return an error.
	// It is possible to provide both an access key pair and a role to assume.
	RoleArn pulumi.StringPtrInput
	// Secret access key to authenticate against the AWS secrets manager.
	// Can be omitted and directly provided to Vault using the `AWS_SECRET_ACCESS_KEY` environment
	// variable.
	SecretAccessKey pulumi.StringPtrInput
	// Template describing how to generate external secret names.
	// Supports a subset of the Go Template syntax.
	SecretNameTemplate pulumi.StringPtrInput
	// The type of the secrets destination (`aws-sm`).
	Type pulumi.StringPtrInput
}

func (SyncAwsDestinationState) ElementType() reflect.Type {
	return reflect.TypeOf((*syncAwsDestinationState)(nil)).Elem()
}

type syncAwsDestinationArgs struct {
	// Access key id to authenticate against the AWS secrets manager.
	// Can be omitted and directly provided to Vault using the `AWS_ACCESS_KEY_ID` environment
	// variable.
	AccessKeyId *string `pulumi:"accessKeyId"`
	// Custom tags to set on the secret managed at the destination.
	CustomTags map[string]string `pulumi:"customTags"`
	// Optional extra protection that must match the trust policy granting access to the
	// AWS IAM role ARN. We recommend using a different random UUID per destination. The value is generated by users.
	// The field is mutable with no special condition, but users must be careful that the new value fits with the trust
	// relationship condition they set on AWS otherwise sync operations will start to fail due to client-side access
	// denied errors. Ignored if the `roleArn` field is empty.
	ExternalId *string `pulumi:"externalId"`
	// Determines what level of information is synced as a distinct resource
	// at the destination. Supports `secret-path` and `secret-key`.
	Granularity *string `pulumi:"granularity"`
	// Unique name of the AWS destination.
	Name *string `pulumi:"name"`
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
	Namespace *string `pulumi:"namespace"`
	// Region where to manage the secrets manager entries.
	// Can be omitted and directly provided to Vault using the `AWS_REGION` environment
	// variable.
	Region *string `pulumi:"region"`
	// Specifies a role to assume when connecting to AWS. When assuming a role,
	// Vault uses temporary STS credentials to authenticate. An initial session with the proper trust relationship must
	// exist for Vault to be able to assume this role. The role can be in a different account.
	// The value is mutable as long as the new role targets the same AWS account ID. If not, the BE will return an error.
	// It is possible to provide both an access key pair and a role to assume.
	RoleArn *string `pulumi:"roleArn"`
	// Secret access key to authenticate against the AWS secrets manager.
	// Can be omitted and directly provided to Vault using the `AWS_SECRET_ACCESS_KEY` environment
	// variable.
	SecretAccessKey *string `pulumi:"secretAccessKey"`
	// Template describing how to generate external secret names.
	// Supports a subset of the Go Template syntax.
	SecretNameTemplate *string `pulumi:"secretNameTemplate"`
}

// The set of arguments for constructing a SyncAwsDestination resource.
type SyncAwsDestinationArgs struct {
	// Access key id to authenticate against the AWS secrets manager.
	// Can be omitted and directly provided to Vault using the `AWS_ACCESS_KEY_ID` environment
	// variable.
	AccessKeyId pulumi.StringPtrInput
	// Custom tags to set on the secret managed at the destination.
	CustomTags pulumi.StringMapInput
	// Optional extra protection that must match the trust policy granting access to the
	// AWS IAM role ARN. We recommend using a different random UUID per destination. The value is generated by users.
	// The field is mutable with no special condition, but users must be careful that the new value fits with the trust
	// relationship condition they set on AWS otherwise sync operations will start to fail due to client-side access
	// denied errors. Ignored if the `roleArn` field is empty.
	ExternalId pulumi.StringPtrInput
	// Determines what level of information is synced as a distinct resource
	// at the destination. Supports `secret-path` and `secret-key`.
	Granularity pulumi.StringPtrInput
	// Unique name of the AWS destination.
	Name pulumi.StringPtrInput
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
	Namespace pulumi.StringPtrInput
	// Region where to manage the secrets manager entries.
	// Can be omitted and directly provided to Vault using the `AWS_REGION` environment
	// variable.
	Region pulumi.StringPtrInput
	// Specifies a role to assume when connecting to AWS. When assuming a role,
	// Vault uses temporary STS credentials to authenticate. An initial session with the proper trust relationship must
	// exist for Vault to be able to assume this role. The role can be in a different account.
	// The value is mutable as long as the new role targets the same AWS account ID. If not, the BE will return an error.
	// It is possible to provide both an access key pair and a role to assume.
	RoleArn pulumi.StringPtrInput
	// Secret access key to authenticate against the AWS secrets manager.
	// Can be omitted and directly provided to Vault using the `AWS_SECRET_ACCESS_KEY` environment
	// variable.
	SecretAccessKey pulumi.StringPtrInput
	// Template describing how to generate external secret names.
	// Supports a subset of the Go Template syntax.
	SecretNameTemplate pulumi.StringPtrInput
}

func (SyncAwsDestinationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*syncAwsDestinationArgs)(nil)).Elem()
}

type SyncAwsDestinationInput interface {
	pulumi.Input

	ToSyncAwsDestinationOutput() SyncAwsDestinationOutput
	ToSyncAwsDestinationOutputWithContext(ctx context.Context) SyncAwsDestinationOutput
}

func (*SyncAwsDestination) ElementType() reflect.Type {
	return reflect.TypeOf((**SyncAwsDestination)(nil)).Elem()
}

func (i *SyncAwsDestination) ToSyncAwsDestinationOutput() SyncAwsDestinationOutput {
	return i.ToSyncAwsDestinationOutputWithContext(context.Background())
}

func (i *SyncAwsDestination) ToSyncAwsDestinationOutputWithContext(ctx context.Context) SyncAwsDestinationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SyncAwsDestinationOutput)
}

// SyncAwsDestinationArrayInput is an input type that accepts SyncAwsDestinationArray and SyncAwsDestinationArrayOutput values.
// You can construct a concrete instance of `SyncAwsDestinationArrayInput` via:
//
//	SyncAwsDestinationArray{ SyncAwsDestinationArgs{...} }
type SyncAwsDestinationArrayInput interface {
	pulumi.Input

	ToSyncAwsDestinationArrayOutput() SyncAwsDestinationArrayOutput
	ToSyncAwsDestinationArrayOutputWithContext(context.Context) SyncAwsDestinationArrayOutput
}

type SyncAwsDestinationArray []SyncAwsDestinationInput

func (SyncAwsDestinationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SyncAwsDestination)(nil)).Elem()
}

func (i SyncAwsDestinationArray) ToSyncAwsDestinationArrayOutput() SyncAwsDestinationArrayOutput {
	return i.ToSyncAwsDestinationArrayOutputWithContext(context.Background())
}

func (i SyncAwsDestinationArray) ToSyncAwsDestinationArrayOutputWithContext(ctx context.Context) SyncAwsDestinationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SyncAwsDestinationArrayOutput)
}

// SyncAwsDestinationMapInput is an input type that accepts SyncAwsDestinationMap and SyncAwsDestinationMapOutput values.
// You can construct a concrete instance of `SyncAwsDestinationMapInput` via:
//
//	SyncAwsDestinationMap{ "key": SyncAwsDestinationArgs{...} }
type SyncAwsDestinationMapInput interface {
	pulumi.Input

	ToSyncAwsDestinationMapOutput() SyncAwsDestinationMapOutput
	ToSyncAwsDestinationMapOutputWithContext(context.Context) SyncAwsDestinationMapOutput
}

type SyncAwsDestinationMap map[string]SyncAwsDestinationInput

func (SyncAwsDestinationMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SyncAwsDestination)(nil)).Elem()
}

func (i SyncAwsDestinationMap) ToSyncAwsDestinationMapOutput() SyncAwsDestinationMapOutput {
	return i.ToSyncAwsDestinationMapOutputWithContext(context.Background())
}

func (i SyncAwsDestinationMap) ToSyncAwsDestinationMapOutputWithContext(ctx context.Context) SyncAwsDestinationMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SyncAwsDestinationMapOutput)
}

type SyncAwsDestinationOutput struct{ *pulumi.OutputState }

func (SyncAwsDestinationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SyncAwsDestination)(nil)).Elem()
}

func (o SyncAwsDestinationOutput) ToSyncAwsDestinationOutput() SyncAwsDestinationOutput {
	return o
}

func (o SyncAwsDestinationOutput) ToSyncAwsDestinationOutputWithContext(ctx context.Context) SyncAwsDestinationOutput {
	return o
}

// Access key id to authenticate against the AWS secrets manager.
// Can be omitted and directly provided to Vault using the `AWS_ACCESS_KEY_ID` environment
// variable.
func (o SyncAwsDestinationOutput) AccessKeyId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SyncAwsDestination) pulumi.StringPtrOutput { return v.AccessKeyId }).(pulumi.StringPtrOutput)
}

// Custom tags to set on the secret managed at the destination.
func (o SyncAwsDestinationOutput) CustomTags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *SyncAwsDestination) pulumi.StringMapOutput { return v.CustomTags }).(pulumi.StringMapOutput)
}

// Optional extra protection that must match the trust policy granting access to the
// AWS IAM role ARN. We recommend using a different random UUID per destination. The value is generated by users.
// The field is mutable with no special condition, but users must be careful that the new value fits with the trust
// relationship condition they set on AWS otherwise sync operations will start to fail due to client-side access
// denied errors. Ignored if the `roleArn` field is empty.
func (o SyncAwsDestinationOutput) ExternalId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SyncAwsDestination) pulumi.StringPtrOutput { return v.ExternalId }).(pulumi.StringPtrOutput)
}

// Determines what level of information is synced as a distinct resource
// at the destination. Supports `secret-path` and `secret-key`.
func (o SyncAwsDestinationOutput) Granularity() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SyncAwsDestination) pulumi.StringPtrOutput { return v.Granularity }).(pulumi.StringPtrOutput)
}

// Unique name of the AWS destination.
func (o SyncAwsDestinationOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SyncAwsDestination) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The namespace to provision the resource in.
// The value should not contain leading or trailing forward slashes.
// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
func (o SyncAwsDestinationOutput) Namespace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SyncAwsDestination) pulumi.StringPtrOutput { return v.Namespace }).(pulumi.StringPtrOutput)
}

// Region where to manage the secrets manager entries.
// Can be omitted and directly provided to Vault using the `AWS_REGION` environment
// variable.
func (o SyncAwsDestinationOutput) Region() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SyncAwsDestination) pulumi.StringPtrOutput { return v.Region }).(pulumi.StringPtrOutput)
}

// Specifies a role to assume when connecting to AWS. When assuming a role,
// Vault uses temporary STS credentials to authenticate. An initial session with the proper trust relationship must
// exist for Vault to be able to assume this role. The role can be in a different account.
// The value is mutable as long as the new role targets the same AWS account ID. If not, the BE will return an error.
// It is possible to provide both an access key pair and a role to assume.
func (o SyncAwsDestinationOutput) RoleArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SyncAwsDestination) pulumi.StringPtrOutput { return v.RoleArn }).(pulumi.StringPtrOutput)
}

// Secret access key to authenticate against the AWS secrets manager.
// Can be omitted and directly provided to Vault using the `AWS_SECRET_ACCESS_KEY` environment
// variable.
func (o SyncAwsDestinationOutput) SecretAccessKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SyncAwsDestination) pulumi.StringPtrOutput { return v.SecretAccessKey }).(pulumi.StringPtrOutput)
}

// Template describing how to generate external secret names.
// Supports a subset of the Go Template syntax.
func (o SyncAwsDestinationOutput) SecretNameTemplate() pulumi.StringOutput {
	return o.ApplyT(func(v *SyncAwsDestination) pulumi.StringOutput { return v.SecretNameTemplate }).(pulumi.StringOutput)
}

// The type of the secrets destination (`aws-sm`).
func (o SyncAwsDestinationOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *SyncAwsDestination) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

type SyncAwsDestinationArrayOutput struct{ *pulumi.OutputState }

func (SyncAwsDestinationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SyncAwsDestination)(nil)).Elem()
}

func (o SyncAwsDestinationArrayOutput) ToSyncAwsDestinationArrayOutput() SyncAwsDestinationArrayOutput {
	return o
}

func (o SyncAwsDestinationArrayOutput) ToSyncAwsDestinationArrayOutputWithContext(ctx context.Context) SyncAwsDestinationArrayOutput {
	return o
}

func (o SyncAwsDestinationArrayOutput) Index(i pulumi.IntInput) SyncAwsDestinationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SyncAwsDestination {
		return vs[0].([]*SyncAwsDestination)[vs[1].(int)]
	}).(SyncAwsDestinationOutput)
}

type SyncAwsDestinationMapOutput struct{ *pulumi.OutputState }

func (SyncAwsDestinationMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SyncAwsDestination)(nil)).Elem()
}

func (o SyncAwsDestinationMapOutput) ToSyncAwsDestinationMapOutput() SyncAwsDestinationMapOutput {
	return o
}

func (o SyncAwsDestinationMapOutput) ToSyncAwsDestinationMapOutputWithContext(ctx context.Context) SyncAwsDestinationMapOutput {
	return o
}

func (o SyncAwsDestinationMapOutput) MapIndex(k pulumi.StringInput) SyncAwsDestinationOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SyncAwsDestination {
		return vs[0].(map[string]*SyncAwsDestination)[vs[1].(string)]
	}).(SyncAwsDestinationOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SyncAwsDestinationInput)(nil)).Elem(), &SyncAwsDestination{})
	pulumi.RegisterInputType(reflect.TypeOf((*SyncAwsDestinationArrayInput)(nil)).Elem(), SyncAwsDestinationArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SyncAwsDestinationMapInput)(nil)).Elem(), SyncAwsDestinationMap{})
	pulumi.RegisterOutputType(SyncAwsDestinationOutput{})
	pulumi.RegisterOutputType(SyncAwsDestinationArrayOutput{})
	pulumi.RegisterOutputType(SyncAwsDestinationMapOutput{})
}
