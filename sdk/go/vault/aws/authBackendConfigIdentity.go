// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package aws

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-vault/sdk/v6/go/vault/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages an AWS auth backend identity configuration in a Vault server. This configuration defines how Vault interacts
// with the identity store. See the [Vault documentation](https://www.vaultproject.io/docs/auth/aws.html) for more
// information.
//
// ## Example Usage
//
// <!--Start PulumiCodeChooser -->
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-vault/sdk/v6/go/vault"
//	"github.com/pulumi/pulumi-vault/sdk/v6/go/vault/aws"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			aws, err := vault.NewAuthBackend(ctx, "aws", &vault.AuthBackendArgs{
//				Type: pulumi.String("aws"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = aws.NewAuthBackendConfigIdentity(ctx, "example", &aws.AuthBackendConfigIdentityArgs{
//				Backend:  aws.Path,
//				IamAlias: pulumi.String("full_arn"),
//				IamMetadatas: pulumi.StringArray{
//					pulumi.String("canonical_arn"),
//					pulumi.String("account_id"),
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
// <!--End PulumiCodeChooser -->
//
// ## Import
//
// AWS auth backend identity config can be imported using `auth/`, the `backend` path, and `/config/identity` e.g.
//
// ```sh
// $ pulumi import vault:aws/authBackendConfigIdentity:AuthBackendConfigIdentity example auth/aws/config/identity
// ```
type AuthBackendConfigIdentity struct {
	pulumi.CustomResourceState

	// Unique name of the auth backend to configure.
	Backend pulumi.StringPtrOutput `pulumi:"backend"`
	// How to generate the identity alias when using the ec2 auth method. Valid choices are
	// `roleId`, `instanceId`, and `imageId`. Defaults to `roleId`
	Ec2Alias pulumi.StringPtrOutput `pulumi:"ec2Alias"`
	// The metadata to include on the token returned by the `login` endpoint. This metadata will be
	// added to both audit logs, and on the `ec2Alias`
	Ec2Metadatas pulumi.StringArrayOutput `pulumi:"ec2Metadatas"`
	// How to generate the identity alias when using the iam auth method. Valid choices are
	// `roleId`, `uniqueId`, and `fullArn`. Defaults to `roleId`
	IamAlias pulumi.StringPtrOutput `pulumi:"iamAlias"`
	// The metadata to include on the token returned by the `login` endpoint. This metadata will be
	// added to both audit logs, and on the `iamAlias`
	IamMetadatas pulumi.StringArrayOutput `pulumi:"iamMetadatas"`
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
	// *Available only for Vault Enterprise*.
	Namespace pulumi.StringPtrOutput `pulumi:"namespace"`
}

// NewAuthBackendConfigIdentity registers a new resource with the given unique name, arguments, and options.
func NewAuthBackendConfigIdentity(ctx *pulumi.Context,
	name string, args *AuthBackendConfigIdentityArgs, opts ...pulumi.ResourceOption) (*AuthBackendConfigIdentity, error) {
	if args == nil {
		args = &AuthBackendConfigIdentityArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource AuthBackendConfigIdentity
	err := ctx.RegisterResource("vault:aws/authBackendConfigIdentity:AuthBackendConfigIdentity", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAuthBackendConfigIdentity gets an existing AuthBackendConfigIdentity resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAuthBackendConfigIdentity(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AuthBackendConfigIdentityState, opts ...pulumi.ResourceOption) (*AuthBackendConfigIdentity, error) {
	var resource AuthBackendConfigIdentity
	err := ctx.ReadResource("vault:aws/authBackendConfigIdentity:AuthBackendConfigIdentity", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AuthBackendConfigIdentity resources.
type authBackendConfigIdentityState struct {
	// Unique name of the auth backend to configure.
	Backend *string `pulumi:"backend"`
	// How to generate the identity alias when using the ec2 auth method. Valid choices are
	// `roleId`, `instanceId`, and `imageId`. Defaults to `roleId`
	Ec2Alias *string `pulumi:"ec2Alias"`
	// The metadata to include on the token returned by the `login` endpoint. This metadata will be
	// added to both audit logs, and on the `ec2Alias`
	Ec2Metadatas []string `pulumi:"ec2Metadatas"`
	// How to generate the identity alias when using the iam auth method. Valid choices are
	// `roleId`, `uniqueId`, and `fullArn`. Defaults to `roleId`
	IamAlias *string `pulumi:"iamAlias"`
	// The metadata to include on the token returned by the `login` endpoint. This metadata will be
	// added to both audit logs, and on the `iamAlias`
	IamMetadatas []string `pulumi:"iamMetadatas"`
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
	// *Available only for Vault Enterprise*.
	Namespace *string `pulumi:"namespace"`
}

type AuthBackendConfigIdentityState struct {
	// Unique name of the auth backend to configure.
	Backend pulumi.StringPtrInput
	// How to generate the identity alias when using the ec2 auth method. Valid choices are
	// `roleId`, `instanceId`, and `imageId`. Defaults to `roleId`
	Ec2Alias pulumi.StringPtrInput
	// The metadata to include on the token returned by the `login` endpoint. This metadata will be
	// added to both audit logs, and on the `ec2Alias`
	Ec2Metadatas pulumi.StringArrayInput
	// How to generate the identity alias when using the iam auth method. Valid choices are
	// `roleId`, `uniqueId`, and `fullArn`. Defaults to `roleId`
	IamAlias pulumi.StringPtrInput
	// The metadata to include on the token returned by the `login` endpoint. This metadata will be
	// added to both audit logs, and on the `iamAlias`
	IamMetadatas pulumi.StringArrayInput
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
	// *Available only for Vault Enterprise*.
	Namespace pulumi.StringPtrInput
}

func (AuthBackendConfigIdentityState) ElementType() reflect.Type {
	return reflect.TypeOf((*authBackendConfigIdentityState)(nil)).Elem()
}

type authBackendConfigIdentityArgs struct {
	// Unique name of the auth backend to configure.
	Backend *string `pulumi:"backend"`
	// How to generate the identity alias when using the ec2 auth method. Valid choices are
	// `roleId`, `instanceId`, and `imageId`. Defaults to `roleId`
	Ec2Alias *string `pulumi:"ec2Alias"`
	// The metadata to include on the token returned by the `login` endpoint. This metadata will be
	// added to both audit logs, and on the `ec2Alias`
	Ec2Metadatas []string `pulumi:"ec2Metadatas"`
	// How to generate the identity alias when using the iam auth method. Valid choices are
	// `roleId`, `uniqueId`, and `fullArn`. Defaults to `roleId`
	IamAlias *string `pulumi:"iamAlias"`
	// The metadata to include on the token returned by the `login` endpoint. This metadata will be
	// added to both audit logs, and on the `iamAlias`
	IamMetadatas []string `pulumi:"iamMetadatas"`
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
	// *Available only for Vault Enterprise*.
	Namespace *string `pulumi:"namespace"`
}

// The set of arguments for constructing a AuthBackendConfigIdentity resource.
type AuthBackendConfigIdentityArgs struct {
	// Unique name of the auth backend to configure.
	Backend pulumi.StringPtrInput
	// How to generate the identity alias when using the ec2 auth method. Valid choices are
	// `roleId`, `instanceId`, and `imageId`. Defaults to `roleId`
	Ec2Alias pulumi.StringPtrInput
	// The metadata to include on the token returned by the `login` endpoint. This metadata will be
	// added to both audit logs, and on the `ec2Alias`
	Ec2Metadatas pulumi.StringArrayInput
	// How to generate the identity alias when using the iam auth method. Valid choices are
	// `roleId`, `uniqueId`, and `fullArn`. Defaults to `roleId`
	IamAlias pulumi.StringPtrInput
	// The metadata to include on the token returned by the `login` endpoint. This metadata will be
	// added to both audit logs, and on the `iamAlias`
	IamMetadatas pulumi.StringArrayInput
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
	// *Available only for Vault Enterprise*.
	Namespace pulumi.StringPtrInput
}

func (AuthBackendConfigIdentityArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*authBackendConfigIdentityArgs)(nil)).Elem()
}

type AuthBackendConfigIdentityInput interface {
	pulumi.Input

	ToAuthBackendConfigIdentityOutput() AuthBackendConfigIdentityOutput
	ToAuthBackendConfigIdentityOutputWithContext(ctx context.Context) AuthBackendConfigIdentityOutput
}

func (*AuthBackendConfigIdentity) ElementType() reflect.Type {
	return reflect.TypeOf((**AuthBackendConfigIdentity)(nil)).Elem()
}

func (i *AuthBackendConfigIdentity) ToAuthBackendConfigIdentityOutput() AuthBackendConfigIdentityOutput {
	return i.ToAuthBackendConfigIdentityOutputWithContext(context.Background())
}

func (i *AuthBackendConfigIdentity) ToAuthBackendConfigIdentityOutputWithContext(ctx context.Context) AuthBackendConfigIdentityOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AuthBackendConfigIdentityOutput)
}

// AuthBackendConfigIdentityArrayInput is an input type that accepts AuthBackendConfigIdentityArray and AuthBackendConfigIdentityArrayOutput values.
// You can construct a concrete instance of `AuthBackendConfigIdentityArrayInput` via:
//
//	AuthBackendConfigIdentityArray{ AuthBackendConfigIdentityArgs{...} }
type AuthBackendConfigIdentityArrayInput interface {
	pulumi.Input

	ToAuthBackendConfigIdentityArrayOutput() AuthBackendConfigIdentityArrayOutput
	ToAuthBackendConfigIdentityArrayOutputWithContext(context.Context) AuthBackendConfigIdentityArrayOutput
}

type AuthBackendConfigIdentityArray []AuthBackendConfigIdentityInput

func (AuthBackendConfigIdentityArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AuthBackendConfigIdentity)(nil)).Elem()
}

func (i AuthBackendConfigIdentityArray) ToAuthBackendConfigIdentityArrayOutput() AuthBackendConfigIdentityArrayOutput {
	return i.ToAuthBackendConfigIdentityArrayOutputWithContext(context.Background())
}

func (i AuthBackendConfigIdentityArray) ToAuthBackendConfigIdentityArrayOutputWithContext(ctx context.Context) AuthBackendConfigIdentityArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AuthBackendConfigIdentityArrayOutput)
}

// AuthBackendConfigIdentityMapInput is an input type that accepts AuthBackendConfigIdentityMap and AuthBackendConfigIdentityMapOutput values.
// You can construct a concrete instance of `AuthBackendConfigIdentityMapInput` via:
//
//	AuthBackendConfigIdentityMap{ "key": AuthBackendConfigIdentityArgs{...} }
type AuthBackendConfigIdentityMapInput interface {
	pulumi.Input

	ToAuthBackendConfigIdentityMapOutput() AuthBackendConfigIdentityMapOutput
	ToAuthBackendConfigIdentityMapOutputWithContext(context.Context) AuthBackendConfigIdentityMapOutput
}

type AuthBackendConfigIdentityMap map[string]AuthBackendConfigIdentityInput

func (AuthBackendConfigIdentityMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AuthBackendConfigIdentity)(nil)).Elem()
}

func (i AuthBackendConfigIdentityMap) ToAuthBackendConfigIdentityMapOutput() AuthBackendConfigIdentityMapOutput {
	return i.ToAuthBackendConfigIdentityMapOutputWithContext(context.Background())
}

func (i AuthBackendConfigIdentityMap) ToAuthBackendConfigIdentityMapOutputWithContext(ctx context.Context) AuthBackendConfigIdentityMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AuthBackendConfigIdentityMapOutput)
}

type AuthBackendConfigIdentityOutput struct{ *pulumi.OutputState }

func (AuthBackendConfigIdentityOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AuthBackendConfigIdentity)(nil)).Elem()
}

func (o AuthBackendConfigIdentityOutput) ToAuthBackendConfigIdentityOutput() AuthBackendConfigIdentityOutput {
	return o
}

func (o AuthBackendConfigIdentityOutput) ToAuthBackendConfigIdentityOutputWithContext(ctx context.Context) AuthBackendConfigIdentityOutput {
	return o
}

// Unique name of the auth backend to configure.
func (o AuthBackendConfigIdentityOutput) Backend() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AuthBackendConfigIdentity) pulumi.StringPtrOutput { return v.Backend }).(pulumi.StringPtrOutput)
}

// How to generate the identity alias when using the ec2 auth method. Valid choices are
// `roleId`, `instanceId`, and `imageId`. Defaults to `roleId`
func (o AuthBackendConfigIdentityOutput) Ec2Alias() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AuthBackendConfigIdentity) pulumi.StringPtrOutput { return v.Ec2Alias }).(pulumi.StringPtrOutput)
}

// The metadata to include on the token returned by the `login` endpoint. This metadata will be
// added to both audit logs, and on the `ec2Alias`
func (o AuthBackendConfigIdentityOutput) Ec2Metadatas() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *AuthBackendConfigIdentity) pulumi.StringArrayOutput { return v.Ec2Metadatas }).(pulumi.StringArrayOutput)
}

// How to generate the identity alias when using the iam auth method. Valid choices are
// `roleId`, `uniqueId`, and `fullArn`. Defaults to `roleId`
func (o AuthBackendConfigIdentityOutput) IamAlias() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AuthBackendConfigIdentity) pulumi.StringPtrOutput { return v.IamAlias }).(pulumi.StringPtrOutput)
}

// The metadata to include on the token returned by the `login` endpoint. This metadata will be
// added to both audit logs, and on the `iamAlias`
func (o AuthBackendConfigIdentityOutput) IamMetadatas() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *AuthBackendConfigIdentity) pulumi.StringArrayOutput { return v.IamMetadatas }).(pulumi.StringArrayOutput)
}

// The namespace to provision the resource in.
// The value should not contain leading or trailing forward slashes.
// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
// *Available only for Vault Enterprise*.
func (o AuthBackendConfigIdentityOutput) Namespace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AuthBackendConfigIdentity) pulumi.StringPtrOutput { return v.Namespace }).(pulumi.StringPtrOutput)
}

type AuthBackendConfigIdentityArrayOutput struct{ *pulumi.OutputState }

func (AuthBackendConfigIdentityArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AuthBackendConfigIdentity)(nil)).Elem()
}

func (o AuthBackendConfigIdentityArrayOutput) ToAuthBackendConfigIdentityArrayOutput() AuthBackendConfigIdentityArrayOutput {
	return o
}

func (o AuthBackendConfigIdentityArrayOutput) ToAuthBackendConfigIdentityArrayOutputWithContext(ctx context.Context) AuthBackendConfigIdentityArrayOutput {
	return o
}

func (o AuthBackendConfigIdentityArrayOutput) Index(i pulumi.IntInput) AuthBackendConfigIdentityOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *AuthBackendConfigIdentity {
		return vs[0].([]*AuthBackendConfigIdentity)[vs[1].(int)]
	}).(AuthBackendConfigIdentityOutput)
}

type AuthBackendConfigIdentityMapOutput struct{ *pulumi.OutputState }

func (AuthBackendConfigIdentityMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AuthBackendConfigIdentity)(nil)).Elem()
}

func (o AuthBackendConfigIdentityMapOutput) ToAuthBackendConfigIdentityMapOutput() AuthBackendConfigIdentityMapOutput {
	return o
}

func (o AuthBackendConfigIdentityMapOutput) ToAuthBackendConfigIdentityMapOutputWithContext(ctx context.Context) AuthBackendConfigIdentityMapOutput {
	return o
}

func (o AuthBackendConfigIdentityMapOutput) MapIndex(k pulumi.StringInput) AuthBackendConfigIdentityOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *AuthBackendConfigIdentity {
		return vs[0].(map[string]*AuthBackendConfigIdentity)[vs[1].(string)]
	}).(AuthBackendConfigIdentityOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AuthBackendConfigIdentityInput)(nil)).Elem(), &AuthBackendConfigIdentity{})
	pulumi.RegisterInputType(reflect.TypeOf((*AuthBackendConfigIdentityArrayInput)(nil)).Elem(), AuthBackendConfigIdentityArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AuthBackendConfigIdentityMapInput)(nil)).Elem(), AuthBackendConfigIdentityMap{})
	pulumi.RegisterOutputType(AuthBackendConfigIdentityOutput{})
	pulumi.RegisterOutputType(AuthBackendConfigIdentityArrayOutput{})
	pulumi.RegisterOutputType(AuthBackendConfigIdentityMapOutput{})
}
