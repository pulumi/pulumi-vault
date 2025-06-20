// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package terraformcloud

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
//	"github.com/pulumi/pulumi-vault/sdk/v7/go/vault/terraformcloud"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			test, err := terraformcloud.NewSecretBackend(ctx, "test", &terraformcloud.SecretBackendArgs{
//				Backend:     pulumi.String("terraform"),
//				Description: pulumi.String("Manages the Terraform Cloud backend"),
//				Token:       pulumi.String("V0idfhi2iksSDU234ucdbi2nidsi..."),
//			})
//			if err != nil {
//				return err
//			}
//			example, err := terraformcloud.NewSecretRole(ctx, "example", &terraformcloud.SecretRoleArgs{
//				Backend:      test.Backend,
//				Name:         pulumi.String("test-role"),
//				Organization: pulumi.String("example-organization-name"),
//				TeamId:       pulumi.String("team-ieF4isC..."),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = terraformcloud.NewSecretCreds(ctx, "token", &terraformcloud.SecretCredsArgs{
//				Backend: test.Backend,
//				Role:    example.Name,
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
type SecretCreds struct {
	pulumi.CustomResourceState

	Backend pulumi.StringOutput `pulumi:"backend"`
	// The lease associated with the token. Only user tokens will have a
	// Vault lease associated with them.
	LeaseId pulumi.StringOutput `pulumi:"leaseId"`
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
	// *Available only for Vault Enterprise*.
	Namespace pulumi.StringPtrOutput `pulumi:"namespace"`
	// The organization associated with the token provided.
	Organization pulumi.StringOutput `pulumi:"organization"`
	// Name of the role.
	Role pulumi.StringOutput `pulumi:"role"`
	// The team id associated with the token provided.
	TeamId pulumi.StringOutput `pulumi:"teamId"`
	// The actual token that was generated and can be used with API calls
	// to identify the user of the call.
	Token pulumi.StringOutput `pulumi:"token"`
	// The public identifier for a specific token. It can be used
	// to look up information about a token or to revoke a token.
	TokenId pulumi.StringOutput `pulumi:"tokenId"`
}

// NewSecretCreds registers a new resource with the given unique name, arguments, and options.
func NewSecretCreds(ctx *pulumi.Context,
	name string, args *SecretCredsArgs, opts ...pulumi.ResourceOption) (*SecretCreds, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Backend == nil {
		return nil, errors.New("invalid value for required argument 'Backend'")
	}
	if args.Role == nil {
		return nil, errors.New("invalid value for required argument 'Role'")
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"leaseId",
		"token",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource SecretCreds
	err := ctx.RegisterResource("vault:terraformcloud/secretCreds:SecretCreds", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSecretCreds gets an existing SecretCreds resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSecretCreds(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SecretCredsState, opts ...pulumi.ResourceOption) (*SecretCreds, error) {
	var resource SecretCreds
	err := ctx.ReadResource("vault:terraformcloud/secretCreds:SecretCreds", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SecretCreds resources.
type secretCredsState struct {
	Backend *string `pulumi:"backend"`
	// The lease associated with the token. Only user tokens will have a
	// Vault lease associated with them.
	LeaseId *string `pulumi:"leaseId"`
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
	// *Available only for Vault Enterprise*.
	Namespace *string `pulumi:"namespace"`
	// The organization associated with the token provided.
	Organization *string `pulumi:"organization"`
	// Name of the role.
	Role *string `pulumi:"role"`
	// The team id associated with the token provided.
	TeamId *string `pulumi:"teamId"`
	// The actual token that was generated and can be used with API calls
	// to identify the user of the call.
	Token *string `pulumi:"token"`
	// The public identifier for a specific token. It can be used
	// to look up information about a token or to revoke a token.
	TokenId *string `pulumi:"tokenId"`
}

type SecretCredsState struct {
	Backend pulumi.StringPtrInput
	// The lease associated with the token. Only user tokens will have a
	// Vault lease associated with them.
	LeaseId pulumi.StringPtrInput
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
	// *Available only for Vault Enterprise*.
	Namespace pulumi.StringPtrInput
	// The organization associated with the token provided.
	Organization pulumi.StringPtrInput
	// Name of the role.
	Role pulumi.StringPtrInput
	// The team id associated with the token provided.
	TeamId pulumi.StringPtrInput
	// The actual token that was generated and can be used with API calls
	// to identify the user of the call.
	Token pulumi.StringPtrInput
	// The public identifier for a specific token. It can be used
	// to look up information about a token or to revoke a token.
	TokenId pulumi.StringPtrInput
}

func (SecretCredsState) ElementType() reflect.Type {
	return reflect.TypeOf((*secretCredsState)(nil)).Elem()
}

type secretCredsArgs struct {
	Backend string `pulumi:"backend"`
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
	// *Available only for Vault Enterprise*.
	Namespace *string `pulumi:"namespace"`
	// Name of the role.
	Role string `pulumi:"role"`
}

// The set of arguments for constructing a SecretCreds resource.
type SecretCredsArgs struct {
	Backend pulumi.StringInput
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
	// *Available only for Vault Enterprise*.
	Namespace pulumi.StringPtrInput
	// Name of the role.
	Role pulumi.StringInput
}

func (SecretCredsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*secretCredsArgs)(nil)).Elem()
}

type SecretCredsInput interface {
	pulumi.Input

	ToSecretCredsOutput() SecretCredsOutput
	ToSecretCredsOutputWithContext(ctx context.Context) SecretCredsOutput
}

func (*SecretCreds) ElementType() reflect.Type {
	return reflect.TypeOf((**SecretCreds)(nil)).Elem()
}

func (i *SecretCreds) ToSecretCredsOutput() SecretCredsOutput {
	return i.ToSecretCredsOutputWithContext(context.Background())
}

func (i *SecretCreds) ToSecretCredsOutputWithContext(ctx context.Context) SecretCredsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecretCredsOutput)
}

// SecretCredsArrayInput is an input type that accepts SecretCredsArray and SecretCredsArrayOutput values.
// You can construct a concrete instance of `SecretCredsArrayInput` via:
//
//	SecretCredsArray{ SecretCredsArgs{...} }
type SecretCredsArrayInput interface {
	pulumi.Input

	ToSecretCredsArrayOutput() SecretCredsArrayOutput
	ToSecretCredsArrayOutputWithContext(context.Context) SecretCredsArrayOutput
}

type SecretCredsArray []SecretCredsInput

func (SecretCredsArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SecretCreds)(nil)).Elem()
}

func (i SecretCredsArray) ToSecretCredsArrayOutput() SecretCredsArrayOutput {
	return i.ToSecretCredsArrayOutputWithContext(context.Background())
}

func (i SecretCredsArray) ToSecretCredsArrayOutputWithContext(ctx context.Context) SecretCredsArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecretCredsArrayOutput)
}

// SecretCredsMapInput is an input type that accepts SecretCredsMap and SecretCredsMapOutput values.
// You can construct a concrete instance of `SecretCredsMapInput` via:
//
//	SecretCredsMap{ "key": SecretCredsArgs{...} }
type SecretCredsMapInput interface {
	pulumi.Input

	ToSecretCredsMapOutput() SecretCredsMapOutput
	ToSecretCredsMapOutputWithContext(context.Context) SecretCredsMapOutput
}

type SecretCredsMap map[string]SecretCredsInput

func (SecretCredsMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SecretCreds)(nil)).Elem()
}

func (i SecretCredsMap) ToSecretCredsMapOutput() SecretCredsMapOutput {
	return i.ToSecretCredsMapOutputWithContext(context.Background())
}

func (i SecretCredsMap) ToSecretCredsMapOutputWithContext(ctx context.Context) SecretCredsMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecretCredsMapOutput)
}

type SecretCredsOutput struct{ *pulumi.OutputState }

func (SecretCredsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SecretCreds)(nil)).Elem()
}

func (o SecretCredsOutput) ToSecretCredsOutput() SecretCredsOutput {
	return o
}

func (o SecretCredsOutput) ToSecretCredsOutputWithContext(ctx context.Context) SecretCredsOutput {
	return o
}

func (o SecretCredsOutput) Backend() pulumi.StringOutput {
	return o.ApplyT(func(v *SecretCreds) pulumi.StringOutput { return v.Backend }).(pulumi.StringOutput)
}

// The lease associated with the token. Only user tokens will have a
// Vault lease associated with them.
func (o SecretCredsOutput) LeaseId() pulumi.StringOutput {
	return o.ApplyT(func(v *SecretCreds) pulumi.StringOutput { return v.LeaseId }).(pulumi.StringOutput)
}

// The namespace to provision the resource in.
// The value should not contain leading or trailing forward slashes.
// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
// *Available only for Vault Enterprise*.
func (o SecretCredsOutput) Namespace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecretCreds) pulumi.StringPtrOutput { return v.Namespace }).(pulumi.StringPtrOutput)
}

// The organization associated with the token provided.
func (o SecretCredsOutput) Organization() pulumi.StringOutput {
	return o.ApplyT(func(v *SecretCreds) pulumi.StringOutput { return v.Organization }).(pulumi.StringOutput)
}

// Name of the role.
func (o SecretCredsOutput) Role() pulumi.StringOutput {
	return o.ApplyT(func(v *SecretCreds) pulumi.StringOutput { return v.Role }).(pulumi.StringOutput)
}

// The team id associated with the token provided.
func (o SecretCredsOutput) TeamId() pulumi.StringOutput {
	return o.ApplyT(func(v *SecretCreds) pulumi.StringOutput { return v.TeamId }).(pulumi.StringOutput)
}

// The actual token that was generated and can be used with API calls
// to identify the user of the call.
func (o SecretCredsOutput) Token() pulumi.StringOutput {
	return o.ApplyT(func(v *SecretCreds) pulumi.StringOutput { return v.Token }).(pulumi.StringOutput)
}

// The public identifier for a specific token. It can be used
// to look up information about a token or to revoke a token.
func (o SecretCredsOutput) TokenId() pulumi.StringOutput {
	return o.ApplyT(func(v *SecretCreds) pulumi.StringOutput { return v.TokenId }).(pulumi.StringOutput)
}

type SecretCredsArrayOutput struct{ *pulumi.OutputState }

func (SecretCredsArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SecretCreds)(nil)).Elem()
}

func (o SecretCredsArrayOutput) ToSecretCredsArrayOutput() SecretCredsArrayOutput {
	return o
}

func (o SecretCredsArrayOutput) ToSecretCredsArrayOutputWithContext(ctx context.Context) SecretCredsArrayOutput {
	return o
}

func (o SecretCredsArrayOutput) Index(i pulumi.IntInput) SecretCredsOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SecretCreds {
		return vs[0].([]*SecretCreds)[vs[1].(int)]
	}).(SecretCredsOutput)
}

type SecretCredsMapOutput struct{ *pulumi.OutputState }

func (SecretCredsMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SecretCreds)(nil)).Elem()
}

func (o SecretCredsMapOutput) ToSecretCredsMapOutput() SecretCredsMapOutput {
	return o
}

func (o SecretCredsMapOutput) ToSecretCredsMapOutputWithContext(ctx context.Context) SecretCredsMapOutput {
	return o
}

func (o SecretCredsMapOutput) MapIndex(k pulumi.StringInput) SecretCredsOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SecretCreds {
		return vs[0].(map[string]*SecretCreds)[vs[1].(string)]
	}).(SecretCredsOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SecretCredsInput)(nil)).Elem(), &SecretCreds{})
	pulumi.RegisterInputType(reflect.TypeOf((*SecretCredsArrayInput)(nil)).Elem(), SecretCredsArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SecretCredsMapInput)(nil)).Elem(), SecretCredsMap{})
	pulumi.RegisterOutputType(SecretCredsOutput{})
	pulumi.RegisterOutputType(SecretCredsArrayOutput{})
	pulumi.RegisterOutputType(SecretCredsMapOutput{})
}
