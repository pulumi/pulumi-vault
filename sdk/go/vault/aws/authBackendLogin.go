// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package aws

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-vault/sdk/v5/go/vault/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Logs into a Vault server using an AWS auth backend. Login can be
// accomplished using a signed identity request from IAM or using ec2
// instance metadata. For more information, see the [Vault
// documentation](https://www.vaultproject.io/docs/auth/aws.html).
type AuthBackendLogin struct {
	pulumi.CustomResourceState

	// The token's accessor.
	Accessor pulumi.StringOutput `pulumi:"accessor"`
	// The authentication type used to generate this token.
	AuthType pulumi.StringOutput `pulumi:"authType"`
	// The unique name of the AWS auth backend. Defaults to
	// 'aws'.
	Backend pulumi.StringPtrOutput `pulumi:"backend"`
	// The token returned by Vault.
	ClientToken pulumi.StringOutput `pulumi:"clientToken"`
	// The HTTP method used in the signed IAM
	// request.
	IamHttpRequestMethod pulumi.StringPtrOutput `pulumi:"iamHttpRequestMethod"`
	// The base64-encoded body of the signed
	// request.
	IamRequestBody pulumi.StringPtrOutput `pulumi:"iamRequestBody"`
	// The base64-encoded, JSON serialized
	// representation of the GetCallerIdentity HTTP request headers.
	IamRequestHeaders pulumi.StringPtrOutput `pulumi:"iamRequestHeaders"`
	// The base64-encoded HTTP URL used in the signed
	// request.
	IamRequestUrl pulumi.StringPtrOutput `pulumi:"iamRequestUrl"`
	// The base64-encoded EC2 instance identity document to
	// authenticate with. Can be retrieved from the EC2 metadata server.
	Identity pulumi.StringPtrOutput `pulumi:"identity"`
	// The duration in seconds the token will be valid, relative
	// to the time in `leaseStartTime`.
	LeaseDuration pulumi.IntOutput `pulumi:"leaseDuration"`
	// Time at which the lease was read, using the clock of the system where Terraform was running
	LeaseStartTime pulumi.StringOutput `pulumi:"leaseStartTime"`
	// A map of information returned by the Vault server about the
	// authentication used to generate this token.
	Metadata pulumi.MapOutput `pulumi:"metadata"`
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
	// *Available only for Vault Enterprise*.
	Namespace pulumi.StringPtrOutput `pulumi:"namespace"`
	// The unique nonce to be used for login requests. Can be
	// set to a user-specified value, or will contain the server-generated value
	// once a token is issued. EC2 instances can only acquire a single token until
	// the whitelist is tidied again unless they keep track of this nonce.
	Nonce pulumi.StringOutput `pulumi:"nonce"`
	// The PKCS#7 signature of the identity document to
	// authenticate with, with all newline characters removed. Can be retrieved from
	// the EC2 metadata server.
	Pkcs7 pulumi.StringPtrOutput `pulumi:"pkcs7"`
	// The Vault policies assigned to this token.
	Policies pulumi.StringArrayOutput `pulumi:"policies"`
	// Set to true if the token can be extended through renewal.
	Renewable pulumi.BoolOutput `pulumi:"renewable"`
	// The name of the AWS auth backend role to create tokens
	// against.
	Role pulumi.StringOutput `pulumi:"role"`
	// The base64-encoded SHA256 RSA signature of the
	// instance identity document to authenticate with, with all newline characters
	// removed. Can be retrieved from the EC2 metadata server.
	Signature pulumi.StringPtrOutput `pulumi:"signature"`
}

// NewAuthBackendLogin registers a new resource with the given unique name, arguments, and options.
func NewAuthBackendLogin(ctx *pulumi.Context,
	name string, args *AuthBackendLoginArgs, opts ...pulumi.ResourceOption) (*AuthBackendLogin, error) {
	if args == nil {
		args = &AuthBackendLoginArgs{}
	}

	secrets := pulumi.AdditionalSecretOutputs([]string{
		"clientToken",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource AuthBackendLogin
	err := ctx.RegisterResource("vault:aws/authBackendLogin:AuthBackendLogin", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAuthBackendLogin gets an existing AuthBackendLogin resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAuthBackendLogin(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AuthBackendLoginState, opts ...pulumi.ResourceOption) (*AuthBackendLogin, error) {
	var resource AuthBackendLogin
	err := ctx.ReadResource("vault:aws/authBackendLogin:AuthBackendLogin", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AuthBackendLogin resources.
type authBackendLoginState struct {
	// The token's accessor.
	Accessor *string `pulumi:"accessor"`
	// The authentication type used to generate this token.
	AuthType *string `pulumi:"authType"`
	// The unique name of the AWS auth backend. Defaults to
	// 'aws'.
	Backend *string `pulumi:"backend"`
	// The token returned by Vault.
	ClientToken *string `pulumi:"clientToken"`
	// The HTTP method used in the signed IAM
	// request.
	IamHttpRequestMethod *string `pulumi:"iamHttpRequestMethod"`
	// The base64-encoded body of the signed
	// request.
	IamRequestBody *string `pulumi:"iamRequestBody"`
	// The base64-encoded, JSON serialized
	// representation of the GetCallerIdentity HTTP request headers.
	IamRequestHeaders *string `pulumi:"iamRequestHeaders"`
	// The base64-encoded HTTP URL used in the signed
	// request.
	IamRequestUrl *string `pulumi:"iamRequestUrl"`
	// The base64-encoded EC2 instance identity document to
	// authenticate with. Can be retrieved from the EC2 metadata server.
	Identity *string `pulumi:"identity"`
	// The duration in seconds the token will be valid, relative
	// to the time in `leaseStartTime`.
	LeaseDuration *int `pulumi:"leaseDuration"`
	// Time at which the lease was read, using the clock of the system where Terraform was running
	LeaseStartTime *string `pulumi:"leaseStartTime"`
	// A map of information returned by the Vault server about the
	// authentication used to generate this token.
	Metadata map[string]interface{} `pulumi:"metadata"`
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
	// *Available only for Vault Enterprise*.
	Namespace *string `pulumi:"namespace"`
	// The unique nonce to be used for login requests. Can be
	// set to a user-specified value, or will contain the server-generated value
	// once a token is issued. EC2 instances can only acquire a single token until
	// the whitelist is tidied again unless they keep track of this nonce.
	Nonce *string `pulumi:"nonce"`
	// The PKCS#7 signature of the identity document to
	// authenticate with, with all newline characters removed. Can be retrieved from
	// the EC2 metadata server.
	Pkcs7 *string `pulumi:"pkcs7"`
	// The Vault policies assigned to this token.
	Policies []string `pulumi:"policies"`
	// Set to true if the token can be extended through renewal.
	Renewable *bool `pulumi:"renewable"`
	// The name of the AWS auth backend role to create tokens
	// against.
	Role *string `pulumi:"role"`
	// The base64-encoded SHA256 RSA signature of the
	// instance identity document to authenticate with, with all newline characters
	// removed. Can be retrieved from the EC2 metadata server.
	Signature *string `pulumi:"signature"`
}

type AuthBackendLoginState struct {
	// The token's accessor.
	Accessor pulumi.StringPtrInput
	// The authentication type used to generate this token.
	AuthType pulumi.StringPtrInput
	// The unique name of the AWS auth backend. Defaults to
	// 'aws'.
	Backend pulumi.StringPtrInput
	// The token returned by Vault.
	ClientToken pulumi.StringPtrInput
	// The HTTP method used in the signed IAM
	// request.
	IamHttpRequestMethod pulumi.StringPtrInput
	// The base64-encoded body of the signed
	// request.
	IamRequestBody pulumi.StringPtrInput
	// The base64-encoded, JSON serialized
	// representation of the GetCallerIdentity HTTP request headers.
	IamRequestHeaders pulumi.StringPtrInput
	// The base64-encoded HTTP URL used in the signed
	// request.
	IamRequestUrl pulumi.StringPtrInput
	// The base64-encoded EC2 instance identity document to
	// authenticate with. Can be retrieved from the EC2 metadata server.
	Identity pulumi.StringPtrInput
	// The duration in seconds the token will be valid, relative
	// to the time in `leaseStartTime`.
	LeaseDuration pulumi.IntPtrInput
	// Time at which the lease was read, using the clock of the system where Terraform was running
	LeaseStartTime pulumi.StringPtrInput
	// A map of information returned by the Vault server about the
	// authentication used to generate this token.
	Metadata pulumi.MapInput
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
	// *Available only for Vault Enterprise*.
	Namespace pulumi.StringPtrInput
	// The unique nonce to be used for login requests. Can be
	// set to a user-specified value, or will contain the server-generated value
	// once a token is issued. EC2 instances can only acquire a single token until
	// the whitelist is tidied again unless they keep track of this nonce.
	Nonce pulumi.StringPtrInput
	// The PKCS#7 signature of the identity document to
	// authenticate with, with all newline characters removed. Can be retrieved from
	// the EC2 metadata server.
	Pkcs7 pulumi.StringPtrInput
	// The Vault policies assigned to this token.
	Policies pulumi.StringArrayInput
	// Set to true if the token can be extended through renewal.
	Renewable pulumi.BoolPtrInput
	// The name of the AWS auth backend role to create tokens
	// against.
	Role pulumi.StringPtrInput
	// The base64-encoded SHA256 RSA signature of the
	// instance identity document to authenticate with, with all newline characters
	// removed. Can be retrieved from the EC2 metadata server.
	Signature pulumi.StringPtrInput
}

func (AuthBackendLoginState) ElementType() reflect.Type {
	return reflect.TypeOf((*authBackendLoginState)(nil)).Elem()
}

type authBackendLoginArgs struct {
	// The unique name of the AWS auth backend. Defaults to
	// 'aws'.
	Backend *string `pulumi:"backend"`
	// The HTTP method used in the signed IAM
	// request.
	IamHttpRequestMethod *string `pulumi:"iamHttpRequestMethod"`
	// The base64-encoded body of the signed
	// request.
	IamRequestBody *string `pulumi:"iamRequestBody"`
	// The base64-encoded, JSON serialized
	// representation of the GetCallerIdentity HTTP request headers.
	IamRequestHeaders *string `pulumi:"iamRequestHeaders"`
	// The base64-encoded HTTP URL used in the signed
	// request.
	IamRequestUrl *string `pulumi:"iamRequestUrl"`
	// The base64-encoded EC2 instance identity document to
	// authenticate with. Can be retrieved from the EC2 metadata server.
	Identity *string `pulumi:"identity"`
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
	// *Available only for Vault Enterprise*.
	Namespace *string `pulumi:"namespace"`
	// The unique nonce to be used for login requests. Can be
	// set to a user-specified value, or will contain the server-generated value
	// once a token is issued. EC2 instances can only acquire a single token until
	// the whitelist is tidied again unless they keep track of this nonce.
	Nonce *string `pulumi:"nonce"`
	// The PKCS#7 signature of the identity document to
	// authenticate with, with all newline characters removed. Can be retrieved from
	// the EC2 metadata server.
	Pkcs7 *string `pulumi:"pkcs7"`
	// The name of the AWS auth backend role to create tokens
	// against.
	Role *string `pulumi:"role"`
	// The base64-encoded SHA256 RSA signature of the
	// instance identity document to authenticate with, with all newline characters
	// removed. Can be retrieved from the EC2 metadata server.
	Signature *string `pulumi:"signature"`
}

// The set of arguments for constructing a AuthBackendLogin resource.
type AuthBackendLoginArgs struct {
	// The unique name of the AWS auth backend. Defaults to
	// 'aws'.
	Backend pulumi.StringPtrInput
	// The HTTP method used in the signed IAM
	// request.
	IamHttpRequestMethod pulumi.StringPtrInput
	// The base64-encoded body of the signed
	// request.
	IamRequestBody pulumi.StringPtrInput
	// The base64-encoded, JSON serialized
	// representation of the GetCallerIdentity HTTP request headers.
	IamRequestHeaders pulumi.StringPtrInput
	// The base64-encoded HTTP URL used in the signed
	// request.
	IamRequestUrl pulumi.StringPtrInput
	// The base64-encoded EC2 instance identity document to
	// authenticate with. Can be retrieved from the EC2 metadata server.
	Identity pulumi.StringPtrInput
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
	// *Available only for Vault Enterprise*.
	Namespace pulumi.StringPtrInput
	// The unique nonce to be used for login requests. Can be
	// set to a user-specified value, or will contain the server-generated value
	// once a token is issued. EC2 instances can only acquire a single token until
	// the whitelist is tidied again unless they keep track of this nonce.
	Nonce pulumi.StringPtrInput
	// The PKCS#7 signature of the identity document to
	// authenticate with, with all newline characters removed. Can be retrieved from
	// the EC2 metadata server.
	Pkcs7 pulumi.StringPtrInput
	// The name of the AWS auth backend role to create tokens
	// against.
	Role pulumi.StringPtrInput
	// The base64-encoded SHA256 RSA signature of the
	// instance identity document to authenticate with, with all newline characters
	// removed. Can be retrieved from the EC2 metadata server.
	Signature pulumi.StringPtrInput
}

func (AuthBackendLoginArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*authBackendLoginArgs)(nil)).Elem()
}

type AuthBackendLoginInput interface {
	pulumi.Input

	ToAuthBackendLoginOutput() AuthBackendLoginOutput
	ToAuthBackendLoginOutputWithContext(ctx context.Context) AuthBackendLoginOutput
}

func (*AuthBackendLogin) ElementType() reflect.Type {
	return reflect.TypeOf((**AuthBackendLogin)(nil)).Elem()
}

func (i *AuthBackendLogin) ToAuthBackendLoginOutput() AuthBackendLoginOutput {
	return i.ToAuthBackendLoginOutputWithContext(context.Background())
}

func (i *AuthBackendLogin) ToAuthBackendLoginOutputWithContext(ctx context.Context) AuthBackendLoginOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AuthBackendLoginOutput)
}

// AuthBackendLoginArrayInput is an input type that accepts AuthBackendLoginArray and AuthBackendLoginArrayOutput values.
// You can construct a concrete instance of `AuthBackendLoginArrayInput` via:
//
//	AuthBackendLoginArray{ AuthBackendLoginArgs{...} }
type AuthBackendLoginArrayInput interface {
	pulumi.Input

	ToAuthBackendLoginArrayOutput() AuthBackendLoginArrayOutput
	ToAuthBackendLoginArrayOutputWithContext(context.Context) AuthBackendLoginArrayOutput
}

type AuthBackendLoginArray []AuthBackendLoginInput

func (AuthBackendLoginArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AuthBackendLogin)(nil)).Elem()
}

func (i AuthBackendLoginArray) ToAuthBackendLoginArrayOutput() AuthBackendLoginArrayOutput {
	return i.ToAuthBackendLoginArrayOutputWithContext(context.Background())
}

func (i AuthBackendLoginArray) ToAuthBackendLoginArrayOutputWithContext(ctx context.Context) AuthBackendLoginArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AuthBackendLoginArrayOutput)
}

// AuthBackendLoginMapInput is an input type that accepts AuthBackendLoginMap and AuthBackendLoginMapOutput values.
// You can construct a concrete instance of `AuthBackendLoginMapInput` via:
//
//	AuthBackendLoginMap{ "key": AuthBackendLoginArgs{...} }
type AuthBackendLoginMapInput interface {
	pulumi.Input

	ToAuthBackendLoginMapOutput() AuthBackendLoginMapOutput
	ToAuthBackendLoginMapOutputWithContext(context.Context) AuthBackendLoginMapOutput
}

type AuthBackendLoginMap map[string]AuthBackendLoginInput

func (AuthBackendLoginMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AuthBackendLogin)(nil)).Elem()
}

func (i AuthBackendLoginMap) ToAuthBackendLoginMapOutput() AuthBackendLoginMapOutput {
	return i.ToAuthBackendLoginMapOutputWithContext(context.Background())
}

func (i AuthBackendLoginMap) ToAuthBackendLoginMapOutputWithContext(ctx context.Context) AuthBackendLoginMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AuthBackendLoginMapOutput)
}

type AuthBackendLoginOutput struct{ *pulumi.OutputState }

func (AuthBackendLoginOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AuthBackendLogin)(nil)).Elem()
}

func (o AuthBackendLoginOutput) ToAuthBackendLoginOutput() AuthBackendLoginOutput {
	return o
}

func (o AuthBackendLoginOutput) ToAuthBackendLoginOutputWithContext(ctx context.Context) AuthBackendLoginOutput {
	return o
}

// The token's accessor.
func (o AuthBackendLoginOutput) Accessor() pulumi.StringOutput {
	return o.ApplyT(func(v *AuthBackendLogin) pulumi.StringOutput { return v.Accessor }).(pulumi.StringOutput)
}

// The authentication type used to generate this token.
func (o AuthBackendLoginOutput) AuthType() pulumi.StringOutput {
	return o.ApplyT(func(v *AuthBackendLogin) pulumi.StringOutput { return v.AuthType }).(pulumi.StringOutput)
}

// The unique name of the AWS auth backend. Defaults to
// 'aws'.
func (o AuthBackendLoginOutput) Backend() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AuthBackendLogin) pulumi.StringPtrOutput { return v.Backend }).(pulumi.StringPtrOutput)
}

// The token returned by Vault.
func (o AuthBackendLoginOutput) ClientToken() pulumi.StringOutput {
	return o.ApplyT(func(v *AuthBackendLogin) pulumi.StringOutput { return v.ClientToken }).(pulumi.StringOutput)
}

// The HTTP method used in the signed IAM
// request.
func (o AuthBackendLoginOutput) IamHttpRequestMethod() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AuthBackendLogin) pulumi.StringPtrOutput { return v.IamHttpRequestMethod }).(pulumi.StringPtrOutput)
}

// The base64-encoded body of the signed
// request.
func (o AuthBackendLoginOutput) IamRequestBody() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AuthBackendLogin) pulumi.StringPtrOutput { return v.IamRequestBody }).(pulumi.StringPtrOutput)
}

// The base64-encoded, JSON serialized
// representation of the GetCallerIdentity HTTP request headers.
func (o AuthBackendLoginOutput) IamRequestHeaders() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AuthBackendLogin) pulumi.StringPtrOutput { return v.IamRequestHeaders }).(pulumi.StringPtrOutput)
}

// The base64-encoded HTTP URL used in the signed
// request.
func (o AuthBackendLoginOutput) IamRequestUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AuthBackendLogin) pulumi.StringPtrOutput { return v.IamRequestUrl }).(pulumi.StringPtrOutput)
}

// The base64-encoded EC2 instance identity document to
// authenticate with. Can be retrieved from the EC2 metadata server.
func (o AuthBackendLoginOutput) Identity() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AuthBackendLogin) pulumi.StringPtrOutput { return v.Identity }).(pulumi.StringPtrOutput)
}

// The duration in seconds the token will be valid, relative
// to the time in `leaseStartTime`.
func (o AuthBackendLoginOutput) LeaseDuration() pulumi.IntOutput {
	return o.ApplyT(func(v *AuthBackendLogin) pulumi.IntOutput { return v.LeaseDuration }).(pulumi.IntOutput)
}

// Time at which the lease was read, using the clock of the system where Terraform was running
func (o AuthBackendLoginOutput) LeaseStartTime() pulumi.StringOutput {
	return o.ApplyT(func(v *AuthBackendLogin) pulumi.StringOutput { return v.LeaseStartTime }).(pulumi.StringOutput)
}

// A map of information returned by the Vault server about the
// authentication used to generate this token.
func (o AuthBackendLoginOutput) Metadata() pulumi.MapOutput {
	return o.ApplyT(func(v *AuthBackendLogin) pulumi.MapOutput { return v.Metadata }).(pulumi.MapOutput)
}

// The namespace to provision the resource in.
// The value should not contain leading or trailing forward slashes.
// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
// *Available only for Vault Enterprise*.
func (o AuthBackendLoginOutput) Namespace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AuthBackendLogin) pulumi.StringPtrOutput { return v.Namespace }).(pulumi.StringPtrOutput)
}

// The unique nonce to be used for login requests. Can be
// set to a user-specified value, or will contain the server-generated value
// once a token is issued. EC2 instances can only acquire a single token until
// the whitelist is tidied again unless they keep track of this nonce.
func (o AuthBackendLoginOutput) Nonce() pulumi.StringOutput {
	return o.ApplyT(func(v *AuthBackendLogin) pulumi.StringOutput { return v.Nonce }).(pulumi.StringOutput)
}

// The PKCS#7 signature of the identity document to
// authenticate with, with all newline characters removed. Can be retrieved from
// the EC2 metadata server.
func (o AuthBackendLoginOutput) Pkcs7() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AuthBackendLogin) pulumi.StringPtrOutput { return v.Pkcs7 }).(pulumi.StringPtrOutput)
}

// The Vault policies assigned to this token.
func (o AuthBackendLoginOutput) Policies() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *AuthBackendLogin) pulumi.StringArrayOutput { return v.Policies }).(pulumi.StringArrayOutput)
}

// Set to true if the token can be extended through renewal.
func (o AuthBackendLoginOutput) Renewable() pulumi.BoolOutput {
	return o.ApplyT(func(v *AuthBackendLogin) pulumi.BoolOutput { return v.Renewable }).(pulumi.BoolOutput)
}

// The name of the AWS auth backend role to create tokens
// against.
func (o AuthBackendLoginOutput) Role() pulumi.StringOutput {
	return o.ApplyT(func(v *AuthBackendLogin) pulumi.StringOutput { return v.Role }).(pulumi.StringOutput)
}

// The base64-encoded SHA256 RSA signature of the
// instance identity document to authenticate with, with all newline characters
// removed. Can be retrieved from the EC2 metadata server.
func (o AuthBackendLoginOutput) Signature() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AuthBackendLogin) pulumi.StringPtrOutput { return v.Signature }).(pulumi.StringPtrOutput)
}

type AuthBackendLoginArrayOutput struct{ *pulumi.OutputState }

func (AuthBackendLoginArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AuthBackendLogin)(nil)).Elem()
}

func (o AuthBackendLoginArrayOutput) ToAuthBackendLoginArrayOutput() AuthBackendLoginArrayOutput {
	return o
}

func (o AuthBackendLoginArrayOutput) ToAuthBackendLoginArrayOutputWithContext(ctx context.Context) AuthBackendLoginArrayOutput {
	return o
}

func (o AuthBackendLoginArrayOutput) Index(i pulumi.IntInput) AuthBackendLoginOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *AuthBackendLogin {
		return vs[0].([]*AuthBackendLogin)[vs[1].(int)]
	}).(AuthBackendLoginOutput)
}

type AuthBackendLoginMapOutput struct{ *pulumi.OutputState }

func (AuthBackendLoginMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AuthBackendLogin)(nil)).Elem()
}

func (o AuthBackendLoginMapOutput) ToAuthBackendLoginMapOutput() AuthBackendLoginMapOutput {
	return o
}

func (o AuthBackendLoginMapOutput) ToAuthBackendLoginMapOutputWithContext(ctx context.Context) AuthBackendLoginMapOutput {
	return o
}

func (o AuthBackendLoginMapOutput) MapIndex(k pulumi.StringInput) AuthBackendLoginOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *AuthBackendLogin {
		return vs[0].(map[string]*AuthBackendLogin)[vs[1].(string)]
	}).(AuthBackendLoginOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AuthBackendLoginInput)(nil)).Elem(), &AuthBackendLogin{})
	pulumi.RegisterInputType(reflect.TypeOf((*AuthBackendLoginArrayInput)(nil)).Elem(), AuthBackendLoginArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AuthBackendLoginMapInput)(nil)).Elem(), AuthBackendLoginMap{})
	pulumi.RegisterOutputType(AuthBackendLoginOutput{})
	pulumi.RegisterOutputType(AuthBackendLoginArrayOutput{})
	pulumi.RegisterOutputType(AuthBackendLoginMapOutput{})
}
