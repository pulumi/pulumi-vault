// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vault

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The provider type for the vault package. By default, resources use package-wide configuration
// settings, however an explicit `Provider` instance may be created and passed during resource
// construction to achieve fine-grained programmatic control over provider settings. See the
// [documentation](https://www.pulumi.com/docs/reference/programming-model/#providers) for more information.
type Provider struct {
	pulumi.ProviderResourceState

	// If true, adds the value of the `address` argument to the Terraform process environment.
	AddAddressToEnv pulumi.StringPtrOutput `pulumi:"addAddressToEnv"`
	// URL of the root of the target Vault server.
	Address pulumi.StringOutput `pulumi:"address"`
	// Path to directory containing CA certificate files to validate the server's certificate.
	CaCertDir pulumi.StringPtrOutput `pulumi:"caCertDir"`
	// Path to a CA certificate file to validate the server's certificate.
	CaCertFile pulumi.StringPtrOutput `pulumi:"caCertFile"`
	// The namespace to use. Available only for Vault Enterprise.
	Namespace pulumi.StringPtrOutput `pulumi:"namespace"`
	// Name to use as the SNI host when connecting via TLS.
	TlsServerName pulumi.StringPtrOutput `pulumi:"tlsServerName"`
	// Token to use to authenticate to Vault.
	Token pulumi.StringOutput `pulumi:"token"`
	// Token name to use for creating the Vault child token.
	TokenName pulumi.StringPtrOutput `pulumi:"tokenName"`
	// Override the Vault server version, which is normally determined dynamically from the target Vault server
	VaultVersionOverride pulumi.StringPtrOutput `pulumi:"vaultVersionOverride"`
}

// NewProvider registers a new resource with the given unique name, arguments, and options.
func NewProvider(ctx *pulumi.Context,
	name string, args *ProviderArgs, opts ...pulumi.ResourceOption) (*Provider, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Address == nil {
		return nil, errors.New("invalid value for required argument 'Address'")
	}
	if args.Token == nil {
		return nil, errors.New("invalid value for required argument 'Token'")
	}
	if isZero(args.MaxLeaseTtlSeconds) {
		args.MaxLeaseTtlSeconds = pulumi.IntPtr(getEnvOrDefault(1200, parseEnvInt, "TERRAFORM_VAULT_MAX_TTL").(int))
	}
	if isZero(args.MaxRetries) {
		args.MaxRetries = pulumi.IntPtr(getEnvOrDefault(2, parseEnvInt, "VAULT_MAX_RETRIES").(int))
	}
	if isZero(args.SkipTlsVerify) {
		args.SkipTlsVerify = pulumi.BoolPtr(getEnvOrDefault(false, parseEnvBool, "VAULT_SKIP_VERIFY").(bool))
	}
	var resource Provider
	err := ctx.RegisterResource("pulumi:providers:vault", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

type providerArgs struct {
	// If true, adds the value of the `address` argument to the Terraform process environment.
	AddAddressToEnv *string `pulumi:"addAddressToEnv"`
	// URL of the root of the target Vault server.
	Address string `pulumi:"address"`
	// Login to vault with an existing auth method using auth/<mount>/login
	AuthLogin *ProviderAuthLogin `pulumi:"authLogin"`
	// Login to vault using the AWS method
	AuthLoginAws *ProviderAuthLoginAws `pulumi:"authLoginAws"`
	// Login to vault using the azure method
	AuthLoginAzure *ProviderAuthLoginAzure `pulumi:"authLoginAzure"`
	// Login to vault using the cert method
	AuthLoginCert *ProviderAuthLoginCert `pulumi:"authLoginCert"`
	// Login to vault using the gcp method
	AuthLoginGcp *ProviderAuthLoginGcp `pulumi:"authLoginGcp"`
	// Login to vault using the jwt method
	AuthLoginJwt *ProviderAuthLoginJwt `pulumi:"authLoginJwt"`
	// Login to vault using the kerberos method
	AuthLoginKerberos *ProviderAuthLoginKerberos `pulumi:"authLoginKerberos"`
	// Login to vault using the OCI method
	AuthLoginOci *ProviderAuthLoginOci `pulumi:"authLoginOci"`
	// Login to vault using the oidc method
	AuthLoginOidc *ProviderAuthLoginOidc `pulumi:"authLoginOidc"`
	// Login to vault using the radius method
	AuthLoginRadius *ProviderAuthLoginRadius `pulumi:"authLoginRadius"`
	// Login to vault using the userpass method
	AuthLoginUserpass *ProviderAuthLoginUserpass `pulumi:"authLoginUserpass"`
	// Path to directory containing CA certificate files to validate the server's certificate.
	CaCertDir *string `pulumi:"caCertDir"`
	// Path to a CA certificate file to validate the server's certificate.
	CaCertFile *string `pulumi:"caCertFile"`
	// Client authentication credentials.
	ClientAuth *ProviderClientAuth `pulumi:"clientAuth"`
	// The headers to send with each Vault request.
	Headers []ProviderHeader `pulumi:"headers"`
	// Maximum TTL for secret leases requested by this provider.
	MaxLeaseTtlSeconds *int `pulumi:"maxLeaseTtlSeconds"`
	// Maximum number of retries when a 5xx error code is encountered.
	MaxRetries *int `pulumi:"maxRetries"`
	// Maximum number of retries for Client Controlled Consistency related operations
	MaxRetriesCcc *int `pulumi:"maxRetriesCcc"`
	// The namespace to use. Available only for Vault Enterprise.
	Namespace *string `pulumi:"namespace"`
	// Set this to true to prevent the creation of ephemeral child token used by this provider.
	SkipChildToken *bool `pulumi:"skipChildToken"`
	// Skip the dynamic fetching of the Vault server version.
	SkipGetVaultVersion *bool `pulumi:"skipGetVaultVersion"`
	// Set this to true only if the target Vault server is an insecure development instance.
	SkipTlsVerify *bool `pulumi:"skipTlsVerify"`
	// Name to use as the SNI host when connecting via TLS.
	TlsServerName *string `pulumi:"tlsServerName"`
	// Token to use to authenticate to Vault.
	Token string `pulumi:"token"`
	// Token name to use for creating the Vault child token.
	TokenName *string `pulumi:"tokenName"`
	// Override the Vault server version, which is normally determined dynamically from the target Vault server
	VaultVersionOverride *string `pulumi:"vaultVersionOverride"`
}

// The set of arguments for constructing a Provider resource.
type ProviderArgs struct {
	// If true, adds the value of the `address` argument to the Terraform process environment.
	AddAddressToEnv pulumi.StringPtrInput
	// URL of the root of the target Vault server.
	Address pulumi.StringInput
	// Login to vault with an existing auth method using auth/<mount>/login
	AuthLogin ProviderAuthLoginPtrInput
	// Login to vault using the AWS method
	AuthLoginAws ProviderAuthLoginAwsPtrInput
	// Login to vault using the azure method
	AuthLoginAzure ProviderAuthLoginAzurePtrInput
	// Login to vault using the cert method
	AuthLoginCert ProviderAuthLoginCertPtrInput
	// Login to vault using the gcp method
	AuthLoginGcp ProviderAuthLoginGcpPtrInput
	// Login to vault using the jwt method
	AuthLoginJwt ProviderAuthLoginJwtPtrInput
	// Login to vault using the kerberos method
	AuthLoginKerberos ProviderAuthLoginKerberosPtrInput
	// Login to vault using the OCI method
	AuthLoginOci ProviderAuthLoginOciPtrInput
	// Login to vault using the oidc method
	AuthLoginOidc ProviderAuthLoginOidcPtrInput
	// Login to vault using the radius method
	AuthLoginRadius ProviderAuthLoginRadiusPtrInput
	// Login to vault using the userpass method
	AuthLoginUserpass ProviderAuthLoginUserpassPtrInput
	// Path to directory containing CA certificate files to validate the server's certificate.
	CaCertDir pulumi.StringPtrInput
	// Path to a CA certificate file to validate the server's certificate.
	CaCertFile pulumi.StringPtrInput
	// Client authentication credentials.
	ClientAuth ProviderClientAuthPtrInput
	// The headers to send with each Vault request.
	Headers ProviderHeaderArrayInput
	// Maximum TTL for secret leases requested by this provider.
	MaxLeaseTtlSeconds pulumi.IntPtrInput
	// Maximum number of retries when a 5xx error code is encountered.
	MaxRetries pulumi.IntPtrInput
	// Maximum number of retries for Client Controlled Consistency related operations
	MaxRetriesCcc pulumi.IntPtrInput
	// The namespace to use. Available only for Vault Enterprise.
	Namespace pulumi.StringPtrInput
	// Set this to true to prevent the creation of ephemeral child token used by this provider.
	SkipChildToken pulumi.BoolPtrInput
	// Skip the dynamic fetching of the Vault server version.
	SkipGetVaultVersion pulumi.BoolPtrInput
	// Set this to true only if the target Vault server is an insecure development instance.
	SkipTlsVerify pulumi.BoolPtrInput
	// Name to use as the SNI host when connecting via TLS.
	TlsServerName pulumi.StringPtrInput
	// Token to use to authenticate to Vault.
	Token pulumi.StringInput
	// Token name to use for creating the Vault child token.
	TokenName pulumi.StringPtrInput
	// Override the Vault server version, which is normally determined dynamically from the target Vault server
	VaultVersionOverride pulumi.StringPtrInput
}

func (ProviderArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*providerArgs)(nil)).Elem()
}

type ProviderInput interface {
	pulumi.Input

	ToProviderOutput() ProviderOutput
	ToProviderOutputWithContext(ctx context.Context) ProviderOutput
}

func (*Provider) ElementType() reflect.Type {
	return reflect.TypeOf((**Provider)(nil)).Elem()
}

func (i *Provider) ToProviderOutput() ProviderOutput {
	return i.ToProviderOutputWithContext(context.Background())
}

func (i *Provider) ToProviderOutputWithContext(ctx context.Context) ProviderOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProviderOutput)
}

type ProviderOutput struct{ *pulumi.OutputState }

func (ProviderOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Provider)(nil)).Elem()
}

func (o ProviderOutput) ToProviderOutput() ProviderOutput {
	return o
}

func (o ProviderOutput) ToProviderOutputWithContext(ctx context.Context) ProviderOutput {
	return o
}

// If true, adds the value of the `address` argument to the Terraform process environment.
func (o ProviderOutput) AddAddressToEnv() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Provider) pulumi.StringPtrOutput { return v.AddAddressToEnv }).(pulumi.StringPtrOutput)
}

// URL of the root of the target Vault server.
func (o ProviderOutput) Address() pulumi.StringOutput {
	return o.ApplyT(func(v *Provider) pulumi.StringOutput { return v.Address }).(pulumi.StringOutput)
}

// Path to directory containing CA certificate files to validate the server's certificate.
func (o ProviderOutput) CaCertDir() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Provider) pulumi.StringPtrOutput { return v.CaCertDir }).(pulumi.StringPtrOutput)
}

// Path to a CA certificate file to validate the server's certificate.
func (o ProviderOutput) CaCertFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Provider) pulumi.StringPtrOutput { return v.CaCertFile }).(pulumi.StringPtrOutput)
}

// The namespace to use. Available only for Vault Enterprise.
func (o ProviderOutput) Namespace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Provider) pulumi.StringPtrOutput { return v.Namespace }).(pulumi.StringPtrOutput)
}

// Name to use as the SNI host when connecting via TLS.
func (o ProviderOutput) TlsServerName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Provider) pulumi.StringPtrOutput { return v.TlsServerName }).(pulumi.StringPtrOutput)
}

// Token to use to authenticate to Vault.
func (o ProviderOutput) Token() pulumi.StringOutput {
	return o.ApplyT(func(v *Provider) pulumi.StringOutput { return v.Token }).(pulumi.StringOutput)
}

// Token name to use for creating the Vault child token.
func (o ProviderOutput) TokenName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Provider) pulumi.StringPtrOutput { return v.TokenName }).(pulumi.StringPtrOutput)
}

// Override the Vault server version, which is normally determined dynamically from the target Vault server
func (o ProviderOutput) VaultVersionOverride() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Provider) pulumi.StringPtrOutput { return v.VaultVersionOverride }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ProviderInput)(nil)).Elem(), &Provider{})
	pulumi.RegisterOutputType(ProviderOutput{})
}
