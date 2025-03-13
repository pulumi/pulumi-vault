// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package aws

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
//	"github.com/pulumi/pulumi-vault/sdk/v6/go/vault/aws"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			aws, err := aws.NewSecretBackend(ctx, "aws", &aws.SecretBackendArgs{
//				Path:        pulumi.String("my-aws"),
//				Description: pulumi.String("Obtain AWS credentials."),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = aws.NewSecretBackendStaticRole(ctx, "role", &aws.SecretBackendStaticRoleArgs{
//				Backend:        aws.Path,
//				Name:           pulumi.String("test"),
//				Username:       pulumi.String("my-test-user"),
//				RotationPeriod: pulumi.Int(3600),
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
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-vault/sdk/v6/go/vault/aws"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			aws, err := aws.NewSecretBackend(ctx, "aws", &aws.SecretBackendArgs{
//				Path:        pulumi.String("my-aws"),
//				Description: pulumi.String("Obtain AWS credentials."),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = aws.NewSecretBackendStaticRole(ctx, "assume-role", &aws.SecretBackendStaticRoleArgs{
//				Backend:               aws.Path,
//				Name:                  pulumi.String("assume-role-test"),
//				Username:              pulumi.String("my-assume-role-user"),
//				AssumeRoleArn:         pulumi.String("arn:aws:iam::123456789012:role/assume-role"),
//				AssumeRoleSessionName: pulumi.String("assume-role-session"),
//				ExternalId:            pulumi.String("test-id"),
//				RotationPeriod:        pulumi.Int(3600),
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
// AWS secret backend static role can be imported using the full path to the role
// of the form: `<mount_path>/static-roles/<role_name>` e.g.
//
// ```sh
// $ pulumi import vault:aws/secretBackendStaticRole:SecretBackendStaticRole role aws/static-roles/example-role
// ```
type SecretBackendStaticRole struct {
	pulumi.CustomResourceState

	// Specifies the ARN of the role that Vault should assume.
	// When provided, Vault will use AWS STS to assume this role and generate temporary credentials.
	// If `assumeRoleArn` is provided, `assumeRoleSessionName` must also be provided.
	// Requires Vault 1.19+. *Available only for Vault Enterprise*.
	AssumeRoleArn pulumi.StringPtrOutput `pulumi:"assumeRoleArn"`
	// Specifies the session name to use when assuming the role.
	// If `assumeRoleSessionName` is provided, `assumeRoleArn` must also be provided.
	// Requires Vault 1.19+. *Available only for Vault Enterprise*.
	AssumeRoleSessionName pulumi.StringPtrOutput `pulumi:"assumeRoleSessionName"`
	// The unique path this backend should be mounted at. Must
	// not begin or end with a `/`. Defaults to `aws`
	Backend pulumi.StringPtrOutput `pulumi:"backend"`
	// Specifies the external ID to use when assuming the role.
	// Requires Vault 1.19+. *Available only for Vault Enterprise*.
	ExternalId pulumi.StringPtrOutput `pulumi:"externalId"`
	// The name to identify this role within the backend.
	// Must be unique within the backend.
	Name pulumi.StringOutput `pulumi:"name"`
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
	// *Available only for Vault Enterprise*.
	Namespace pulumi.StringPtrOutput `pulumi:"namespace"`
	// How often Vault should rotate the password of the user entry.
	RotationPeriod pulumi.IntOutput `pulumi:"rotationPeriod"`
	// The username of the existing AWS IAM to manage password rotation for.
	Username pulumi.StringOutput `pulumi:"username"`
}

// NewSecretBackendStaticRole registers a new resource with the given unique name, arguments, and options.
func NewSecretBackendStaticRole(ctx *pulumi.Context,
	name string, args *SecretBackendStaticRoleArgs, opts ...pulumi.ResourceOption) (*SecretBackendStaticRole, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.RotationPeriod == nil {
		return nil, errors.New("invalid value for required argument 'RotationPeriod'")
	}
	if args.Username == nil {
		return nil, errors.New("invalid value for required argument 'Username'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource SecretBackendStaticRole
	err := ctx.RegisterResource("vault:aws/secretBackendStaticRole:SecretBackendStaticRole", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSecretBackendStaticRole gets an existing SecretBackendStaticRole resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSecretBackendStaticRole(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SecretBackendStaticRoleState, opts ...pulumi.ResourceOption) (*SecretBackendStaticRole, error) {
	var resource SecretBackendStaticRole
	err := ctx.ReadResource("vault:aws/secretBackendStaticRole:SecretBackendStaticRole", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SecretBackendStaticRole resources.
type secretBackendStaticRoleState struct {
	// Specifies the ARN of the role that Vault should assume.
	// When provided, Vault will use AWS STS to assume this role and generate temporary credentials.
	// If `assumeRoleArn` is provided, `assumeRoleSessionName` must also be provided.
	// Requires Vault 1.19+. *Available only for Vault Enterprise*.
	AssumeRoleArn *string `pulumi:"assumeRoleArn"`
	// Specifies the session name to use when assuming the role.
	// If `assumeRoleSessionName` is provided, `assumeRoleArn` must also be provided.
	// Requires Vault 1.19+. *Available only for Vault Enterprise*.
	AssumeRoleSessionName *string `pulumi:"assumeRoleSessionName"`
	// The unique path this backend should be mounted at. Must
	// not begin or end with a `/`. Defaults to `aws`
	Backend *string `pulumi:"backend"`
	// Specifies the external ID to use when assuming the role.
	// Requires Vault 1.19+. *Available only for Vault Enterprise*.
	ExternalId *string `pulumi:"externalId"`
	// The name to identify this role within the backend.
	// Must be unique within the backend.
	Name *string `pulumi:"name"`
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
	// *Available only for Vault Enterprise*.
	Namespace *string `pulumi:"namespace"`
	// How often Vault should rotate the password of the user entry.
	RotationPeriod *int `pulumi:"rotationPeriod"`
	// The username of the existing AWS IAM to manage password rotation for.
	Username *string `pulumi:"username"`
}

type SecretBackendStaticRoleState struct {
	// Specifies the ARN of the role that Vault should assume.
	// When provided, Vault will use AWS STS to assume this role and generate temporary credentials.
	// If `assumeRoleArn` is provided, `assumeRoleSessionName` must also be provided.
	// Requires Vault 1.19+. *Available only for Vault Enterprise*.
	AssumeRoleArn pulumi.StringPtrInput
	// Specifies the session name to use when assuming the role.
	// If `assumeRoleSessionName` is provided, `assumeRoleArn` must also be provided.
	// Requires Vault 1.19+. *Available only for Vault Enterprise*.
	AssumeRoleSessionName pulumi.StringPtrInput
	// The unique path this backend should be mounted at. Must
	// not begin or end with a `/`. Defaults to `aws`
	Backend pulumi.StringPtrInput
	// Specifies the external ID to use when assuming the role.
	// Requires Vault 1.19+. *Available only for Vault Enterprise*.
	ExternalId pulumi.StringPtrInput
	// The name to identify this role within the backend.
	// Must be unique within the backend.
	Name pulumi.StringPtrInput
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
	// *Available only for Vault Enterprise*.
	Namespace pulumi.StringPtrInput
	// How often Vault should rotate the password of the user entry.
	RotationPeriod pulumi.IntPtrInput
	// The username of the existing AWS IAM to manage password rotation for.
	Username pulumi.StringPtrInput
}

func (SecretBackendStaticRoleState) ElementType() reflect.Type {
	return reflect.TypeOf((*secretBackendStaticRoleState)(nil)).Elem()
}

type secretBackendStaticRoleArgs struct {
	// Specifies the ARN of the role that Vault should assume.
	// When provided, Vault will use AWS STS to assume this role and generate temporary credentials.
	// If `assumeRoleArn` is provided, `assumeRoleSessionName` must also be provided.
	// Requires Vault 1.19+. *Available only for Vault Enterprise*.
	AssumeRoleArn *string `pulumi:"assumeRoleArn"`
	// Specifies the session name to use when assuming the role.
	// If `assumeRoleSessionName` is provided, `assumeRoleArn` must also be provided.
	// Requires Vault 1.19+. *Available only for Vault Enterprise*.
	AssumeRoleSessionName *string `pulumi:"assumeRoleSessionName"`
	// The unique path this backend should be mounted at. Must
	// not begin or end with a `/`. Defaults to `aws`
	Backend *string `pulumi:"backend"`
	// Specifies the external ID to use when assuming the role.
	// Requires Vault 1.19+. *Available only for Vault Enterprise*.
	ExternalId *string `pulumi:"externalId"`
	// The name to identify this role within the backend.
	// Must be unique within the backend.
	Name *string `pulumi:"name"`
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
	// *Available only for Vault Enterprise*.
	Namespace *string `pulumi:"namespace"`
	// How often Vault should rotate the password of the user entry.
	RotationPeriod int `pulumi:"rotationPeriod"`
	// The username of the existing AWS IAM to manage password rotation for.
	Username string `pulumi:"username"`
}

// The set of arguments for constructing a SecretBackendStaticRole resource.
type SecretBackendStaticRoleArgs struct {
	// Specifies the ARN of the role that Vault should assume.
	// When provided, Vault will use AWS STS to assume this role and generate temporary credentials.
	// If `assumeRoleArn` is provided, `assumeRoleSessionName` must also be provided.
	// Requires Vault 1.19+. *Available only for Vault Enterprise*.
	AssumeRoleArn pulumi.StringPtrInput
	// Specifies the session name to use when assuming the role.
	// If `assumeRoleSessionName` is provided, `assumeRoleArn` must also be provided.
	// Requires Vault 1.19+. *Available only for Vault Enterprise*.
	AssumeRoleSessionName pulumi.StringPtrInput
	// The unique path this backend should be mounted at. Must
	// not begin or end with a `/`. Defaults to `aws`
	Backend pulumi.StringPtrInput
	// Specifies the external ID to use when assuming the role.
	// Requires Vault 1.19+. *Available only for Vault Enterprise*.
	ExternalId pulumi.StringPtrInput
	// The name to identify this role within the backend.
	// Must be unique within the backend.
	Name pulumi.StringPtrInput
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
	// *Available only for Vault Enterprise*.
	Namespace pulumi.StringPtrInput
	// How often Vault should rotate the password of the user entry.
	RotationPeriod pulumi.IntInput
	// The username of the existing AWS IAM to manage password rotation for.
	Username pulumi.StringInput
}

func (SecretBackendStaticRoleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*secretBackendStaticRoleArgs)(nil)).Elem()
}

type SecretBackendStaticRoleInput interface {
	pulumi.Input

	ToSecretBackendStaticRoleOutput() SecretBackendStaticRoleOutput
	ToSecretBackendStaticRoleOutputWithContext(ctx context.Context) SecretBackendStaticRoleOutput
}

func (*SecretBackendStaticRole) ElementType() reflect.Type {
	return reflect.TypeOf((**SecretBackendStaticRole)(nil)).Elem()
}

func (i *SecretBackendStaticRole) ToSecretBackendStaticRoleOutput() SecretBackendStaticRoleOutput {
	return i.ToSecretBackendStaticRoleOutputWithContext(context.Background())
}

func (i *SecretBackendStaticRole) ToSecretBackendStaticRoleOutputWithContext(ctx context.Context) SecretBackendStaticRoleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecretBackendStaticRoleOutput)
}

// SecretBackendStaticRoleArrayInput is an input type that accepts SecretBackendStaticRoleArray and SecretBackendStaticRoleArrayOutput values.
// You can construct a concrete instance of `SecretBackendStaticRoleArrayInput` via:
//
//	SecretBackendStaticRoleArray{ SecretBackendStaticRoleArgs{...} }
type SecretBackendStaticRoleArrayInput interface {
	pulumi.Input

	ToSecretBackendStaticRoleArrayOutput() SecretBackendStaticRoleArrayOutput
	ToSecretBackendStaticRoleArrayOutputWithContext(context.Context) SecretBackendStaticRoleArrayOutput
}

type SecretBackendStaticRoleArray []SecretBackendStaticRoleInput

func (SecretBackendStaticRoleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SecretBackendStaticRole)(nil)).Elem()
}

func (i SecretBackendStaticRoleArray) ToSecretBackendStaticRoleArrayOutput() SecretBackendStaticRoleArrayOutput {
	return i.ToSecretBackendStaticRoleArrayOutputWithContext(context.Background())
}

func (i SecretBackendStaticRoleArray) ToSecretBackendStaticRoleArrayOutputWithContext(ctx context.Context) SecretBackendStaticRoleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecretBackendStaticRoleArrayOutput)
}

// SecretBackendStaticRoleMapInput is an input type that accepts SecretBackendStaticRoleMap and SecretBackendStaticRoleMapOutput values.
// You can construct a concrete instance of `SecretBackendStaticRoleMapInput` via:
//
//	SecretBackendStaticRoleMap{ "key": SecretBackendStaticRoleArgs{...} }
type SecretBackendStaticRoleMapInput interface {
	pulumi.Input

	ToSecretBackendStaticRoleMapOutput() SecretBackendStaticRoleMapOutput
	ToSecretBackendStaticRoleMapOutputWithContext(context.Context) SecretBackendStaticRoleMapOutput
}

type SecretBackendStaticRoleMap map[string]SecretBackendStaticRoleInput

func (SecretBackendStaticRoleMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SecretBackendStaticRole)(nil)).Elem()
}

func (i SecretBackendStaticRoleMap) ToSecretBackendStaticRoleMapOutput() SecretBackendStaticRoleMapOutput {
	return i.ToSecretBackendStaticRoleMapOutputWithContext(context.Background())
}

func (i SecretBackendStaticRoleMap) ToSecretBackendStaticRoleMapOutputWithContext(ctx context.Context) SecretBackendStaticRoleMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecretBackendStaticRoleMapOutput)
}

type SecretBackendStaticRoleOutput struct{ *pulumi.OutputState }

func (SecretBackendStaticRoleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SecretBackendStaticRole)(nil)).Elem()
}

func (o SecretBackendStaticRoleOutput) ToSecretBackendStaticRoleOutput() SecretBackendStaticRoleOutput {
	return o
}

func (o SecretBackendStaticRoleOutput) ToSecretBackendStaticRoleOutputWithContext(ctx context.Context) SecretBackendStaticRoleOutput {
	return o
}

// Specifies the ARN of the role that Vault should assume.
// When provided, Vault will use AWS STS to assume this role and generate temporary credentials.
// If `assumeRoleArn` is provided, `assumeRoleSessionName` must also be provided.
// Requires Vault 1.19+. *Available only for Vault Enterprise*.
func (o SecretBackendStaticRoleOutput) AssumeRoleArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecretBackendStaticRole) pulumi.StringPtrOutput { return v.AssumeRoleArn }).(pulumi.StringPtrOutput)
}

// Specifies the session name to use when assuming the role.
// If `assumeRoleSessionName` is provided, `assumeRoleArn` must also be provided.
// Requires Vault 1.19+. *Available only for Vault Enterprise*.
func (o SecretBackendStaticRoleOutput) AssumeRoleSessionName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecretBackendStaticRole) pulumi.StringPtrOutput { return v.AssumeRoleSessionName }).(pulumi.StringPtrOutput)
}

// The unique path this backend should be mounted at. Must
// not begin or end with a `/`. Defaults to `aws`
func (o SecretBackendStaticRoleOutput) Backend() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecretBackendStaticRole) pulumi.StringPtrOutput { return v.Backend }).(pulumi.StringPtrOutput)
}

// Specifies the external ID to use when assuming the role.
// Requires Vault 1.19+. *Available only for Vault Enterprise*.
func (o SecretBackendStaticRoleOutput) ExternalId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecretBackendStaticRole) pulumi.StringPtrOutput { return v.ExternalId }).(pulumi.StringPtrOutput)
}

// The name to identify this role within the backend.
// Must be unique within the backend.
func (o SecretBackendStaticRoleOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SecretBackendStaticRole) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The namespace to provision the resource in.
// The value should not contain leading or trailing forward slashes.
// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
// *Available only for Vault Enterprise*.
func (o SecretBackendStaticRoleOutput) Namespace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecretBackendStaticRole) pulumi.StringPtrOutput { return v.Namespace }).(pulumi.StringPtrOutput)
}

// How often Vault should rotate the password of the user entry.
func (o SecretBackendStaticRoleOutput) RotationPeriod() pulumi.IntOutput {
	return o.ApplyT(func(v *SecretBackendStaticRole) pulumi.IntOutput { return v.RotationPeriod }).(pulumi.IntOutput)
}

// The username of the existing AWS IAM to manage password rotation for.
func (o SecretBackendStaticRoleOutput) Username() pulumi.StringOutput {
	return o.ApplyT(func(v *SecretBackendStaticRole) pulumi.StringOutput { return v.Username }).(pulumi.StringOutput)
}

type SecretBackendStaticRoleArrayOutput struct{ *pulumi.OutputState }

func (SecretBackendStaticRoleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SecretBackendStaticRole)(nil)).Elem()
}

func (o SecretBackendStaticRoleArrayOutput) ToSecretBackendStaticRoleArrayOutput() SecretBackendStaticRoleArrayOutput {
	return o
}

func (o SecretBackendStaticRoleArrayOutput) ToSecretBackendStaticRoleArrayOutputWithContext(ctx context.Context) SecretBackendStaticRoleArrayOutput {
	return o
}

func (o SecretBackendStaticRoleArrayOutput) Index(i pulumi.IntInput) SecretBackendStaticRoleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SecretBackendStaticRole {
		return vs[0].([]*SecretBackendStaticRole)[vs[1].(int)]
	}).(SecretBackendStaticRoleOutput)
}

type SecretBackendStaticRoleMapOutput struct{ *pulumi.OutputState }

func (SecretBackendStaticRoleMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SecretBackendStaticRole)(nil)).Elem()
}

func (o SecretBackendStaticRoleMapOutput) ToSecretBackendStaticRoleMapOutput() SecretBackendStaticRoleMapOutput {
	return o
}

func (o SecretBackendStaticRoleMapOutput) ToSecretBackendStaticRoleMapOutputWithContext(ctx context.Context) SecretBackendStaticRoleMapOutput {
	return o
}

func (o SecretBackendStaticRoleMapOutput) MapIndex(k pulumi.StringInput) SecretBackendStaticRoleOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SecretBackendStaticRole {
		return vs[0].(map[string]*SecretBackendStaticRole)[vs[1].(string)]
	}).(SecretBackendStaticRoleOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SecretBackendStaticRoleInput)(nil)).Elem(), &SecretBackendStaticRole{})
	pulumi.RegisterInputType(reflect.TypeOf((*SecretBackendStaticRoleArrayInput)(nil)).Elem(), SecretBackendStaticRoleArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SecretBackendStaticRoleMapInput)(nil)).Elem(), SecretBackendStaticRoleMap{})
	pulumi.RegisterOutputType(SecretBackendStaticRoleOutput{})
	pulumi.RegisterOutputType(SecretBackendStaticRoleArrayOutput{})
	pulumi.RegisterOutputType(SecretBackendStaticRoleMapOutput{})
}
