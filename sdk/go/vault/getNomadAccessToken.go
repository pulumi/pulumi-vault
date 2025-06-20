// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vault

import (
	"context"
	"reflect"

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
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			config, err := vault.NewNomadSecretBackend(ctx, "config", &vault.NomadSecretBackendArgs{
//				Backend:                pulumi.String("nomad"),
//				Description:            pulumi.String("test description"),
//				DefaultLeaseTtlSeconds: pulumi.Int(3600),
//				MaxLeaseTtlSeconds:     pulumi.Int(7200),
//				Address:                pulumi.String("https://127.0.0.1:4646"),
//				Token:                  pulumi.String("ae20ceaa-..."),
//			})
//			if err != nil {
//				return err
//			}
//			test, err := vault.NewNomadSecretRole(ctx, "test", &vault.NomadSecretRoleArgs{
//				Backend: config.Backend,
//				Role:    pulumi.String("test"),
//				Type:    pulumi.String("client"),
//				Policies: pulumi.StringArray{
//					pulumi.String("readonly"),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			_ = pulumi.All(config.Backend, test.Role).ApplyT(func(_args []interface{}) (vault.GetNomadAccessTokenResult, error) {
//				backend := _args[0].(*string)
//				role := _args[1].(string)
//				return vault.GetNomadAccessTokenResult(interface{}(vault.GetNomadAccessTokenOutput(ctx, vault.GetNomadAccessTokenOutputArgs{
//					Backend: backend,
//					Role:    role,
//				}, nil))), nil
//			}).(vault.GetNomadAccessTokenResultOutput)
//			return nil
//		})
//	}
//
// ```
func GetNomadAccessToken(ctx *pulumi.Context, args *GetNomadAccessTokenArgs, opts ...pulumi.InvokeOption) (*GetNomadAccessTokenResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetNomadAccessTokenResult
	err := ctx.Invoke("vault:index/getNomadAccessToken:getNomadAccessToken", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getNomadAccessToken.
type GetNomadAccessTokenArgs struct {
	// The path to the Nomad secret backend to
	// read credentials from, with no leading or trailing `/`s.
	Backend string `pulumi:"backend"`
	// The namespace of the target resource.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
	// *Available only for Vault Enterprise*.
	Namespace *string `pulumi:"namespace"`
	// The name of the Nomad secret backend role to generate
	// a token for, with no leading or trailing `/`s.
	Role string `pulumi:"role"`
}

// A collection of values returned by getNomadAccessToken.
type GetNomadAccessTokenResult struct {
	// The public identifier for a specific token. It can be used
	// to look up information about a token or to revoke a token.
	AccessorId string `pulumi:"accessorId"`
	Backend    string `pulumi:"backend"`
	// The provider-assigned unique ID for this managed resource.
	Id        string  `pulumi:"id"`
	Namespace *string `pulumi:"namespace"`
	Role      string  `pulumi:"role"`
	// The token to be used when making requests to Nomad and should be kept private.
	SecretId string `pulumi:"secretId"`
}

func GetNomadAccessTokenOutput(ctx *pulumi.Context, args GetNomadAccessTokenOutputArgs, opts ...pulumi.InvokeOption) GetNomadAccessTokenResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (GetNomadAccessTokenResultOutput, error) {
			args := v.(GetNomadAccessTokenArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("vault:index/getNomadAccessToken:getNomadAccessToken", args, GetNomadAccessTokenResultOutput{}, options).(GetNomadAccessTokenResultOutput), nil
		}).(GetNomadAccessTokenResultOutput)
}

// A collection of arguments for invoking getNomadAccessToken.
type GetNomadAccessTokenOutputArgs struct {
	// The path to the Nomad secret backend to
	// read credentials from, with no leading or trailing `/`s.
	Backend pulumi.StringInput `pulumi:"backend"`
	// The namespace of the target resource.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
	// *Available only for Vault Enterprise*.
	Namespace pulumi.StringPtrInput `pulumi:"namespace"`
	// The name of the Nomad secret backend role to generate
	// a token for, with no leading or trailing `/`s.
	Role pulumi.StringInput `pulumi:"role"`
}

func (GetNomadAccessTokenOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetNomadAccessTokenArgs)(nil)).Elem()
}

// A collection of values returned by getNomadAccessToken.
type GetNomadAccessTokenResultOutput struct{ *pulumi.OutputState }

func (GetNomadAccessTokenResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetNomadAccessTokenResult)(nil)).Elem()
}

func (o GetNomadAccessTokenResultOutput) ToGetNomadAccessTokenResultOutput() GetNomadAccessTokenResultOutput {
	return o
}

func (o GetNomadAccessTokenResultOutput) ToGetNomadAccessTokenResultOutputWithContext(ctx context.Context) GetNomadAccessTokenResultOutput {
	return o
}

// The public identifier for a specific token. It can be used
// to look up information about a token or to revoke a token.
func (o GetNomadAccessTokenResultOutput) AccessorId() pulumi.StringOutput {
	return o.ApplyT(func(v GetNomadAccessTokenResult) string { return v.AccessorId }).(pulumi.StringOutput)
}

func (o GetNomadAccessTokenResultOutput) Backend() pulumi.StringOutput {
	return o.ApplyT(func(v GetNomadAccessTokenResult) string { return v.Backend }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetNomadAccessTokenResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetNomadAccessTokenResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetNomadAccessTokenResultOutput) Namespace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetNomadAccessTokenResult) *string { return v.Namespace }).(pulumi.StringPtrOutput)
}

func (o GetNomadAccessTokenResultOutput) Role() pulumi.StringOutput {
	return o.ApplyT(func(v GetNomadAccessTokenResult) string { return v.Role }).(pulumi.StringOutput)
}

// The token to be used when making requests to Nomad and should be kept private.
func (o GetNomadAccessTokenResultOutput) SecretId() pulumi.StringOutput {
	return o.ApplyT(func(v GetNomadAccessTokenResult) string { return v.SecretId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetNomadAccessTokenResultOutput{})
}
