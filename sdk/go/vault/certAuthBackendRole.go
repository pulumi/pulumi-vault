// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vault

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-vault/sdk/v7/go/vault/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a resource to create a role in an [Cert auth backend within Vault](https://www.vaultproject.io/docs/auth/cert.html).
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-std/sdk/go/std"
//	"github.com/pulumi/pulumi-vault/sdk/v7/go/vault"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			cert, err := vault.NewAuthBackend(ctx, "cert", &vault.AuthBackendArgs{
//				Path: pulumi.String("cert"),
//				Type: pulumi.String("cert"),
//			})
//			if err != nil {
//				return err
//			}
//			invokeFile, err := std.File(ctx, &std.FileArgs{
//				Input: "/path/to/certs/ca-cert.pem",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = vault.NewCertAuthBackendRole(ctx, "cert", &vault.CertAuthBackendRoleArgs{
//				Name:        pulumi.String("foo"),
//				Certificate: pulumi.String(invokeFile.Result),
//				Backend:     cert.Path,
//				AllowedNames: pulumi.StringArray{
//					pulumi.String("foo.example.org"),
//					pulumi.String("baz.example.org"),
//				},
//				TokenTtl:    pulumi.Int(300),
//				TokenMaxTtl: pulumi.Int(600),
//				TokenPolicies: pulumi.StringArray{
//					pulumi.String("foo"),
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
type CertAuthBackendRole struct {
	pulumi.CustomResourceState

	// Allowed the common names for authenticated client certificates
	AllowedCommonNames pulumi.StringArrayOutput `pulumi:"allowedCommonNames"`
	// Allowed alternative dns names for authenticated client certificates
	AllowedDnsSans pulumi.StringArrayOutput `pulumi:"allowedDnsSans"`
	// Allowed emails for authenticated client certificates
	AllowedEmailSans pulumi.StringArrayOutput `pulumi:"allowedEmailSans"`
	// DEPRECATED: Please use the individual `allowed_X_sans` parameters instead. Allowed subject names for authenticated client certificates
	AllowedNames pulumi.StringArrayOutput `pulumi:"allowedNames"`
	// Allowed organization units for authenticated client certificates.
	AllowedOrganizationalUnits pulumi.StringArrayOutput `pulumi:"allowedOrganizationalUnits"`
	// Allowed URIs for authenticated client certificates
	AllowedUriSans pulumi.StringArrayOutput `pulumi:"allowedUriSans"`
	// Path to the mounted Cert auth backend
	Backend pulumi.StringPtrOutput `pulumi:"backend"`
	// CA certificate used to validate client certificates
	Certificate pulumi.StringOutput `pulumi:"certificate"`
	// The name to display on tokens issued under this role.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// Name of the role
	Name pulumi.StringOutput `pulumi:"name"`
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
	// *Available only for Vault Enterprise*.
	Namespace pulumi.StringPtrOutput `pulumi:"namespace"`
	// Any additional CA certificates
	// needed to verify OCSP responses. Provided as base64 encoded PEM data.
	// Requires Vault version 1.13+.
	OcspCaCertificates pulumi.StringPtrOutput `pulumi:"ocspCaCertificates"`
	// If enabled, validate certificates'
	// revocation status using OCSP. Requires Vault version 1.13+.
	OcspEnabled pulumi.BoolOutput `pulumi:"ocspEnabled"`
	// If true and an OCSP response cannot
	// be fetched or is of an unknown status, the login will proceed as if the
	// certificate has not been revoked.
	// Requires Vault version 1.13+.
	OcspFailOpen pulumi.BoolOutput `pulumi:"ocspFailOpen"`
	// If set to true, rather than
	// accepting the first successful OCSP response, query all servers and consider
	// the certificate valid only if all servers agree.
	// Requires Vault version 1.13+.
	OcspQueryAllServers pulumi.BoolOutput `pulumi:"ocspQueryAllServers"`
	// : A comma-separated list of OCSP
	// server addresses. If unset, the OCSP server is determined from the
	// AuthorityInformationAccess extension on the certificate being inspected.
	// Requires Vault version 1.13+.
	OcspServersOverrides pulumi.StringArrayOutput `pulumi:"ocspServersOverrides"`
	// TLS extensions required on
	// client certificates
	RequiredExtensions pulumi.StringArrayOutput `pulumi:"requiredExtensions"`
	// Specifies the blocks of IP addresses which are allowed to use the generated token
	TokenBoundCidrs pulumi.StringArrayOutput `pulumi:"tokenBoundCidrs"`
	// Generated Token's Explicit Maximum TTL in seconds
	TokenExplicitMaxTtl pulumi.IntPtrOutput `pulumi:"tokenExplicitMaxTtl"`
	// The maximum lifetime of the generated token
	TokenMaxTtl pulumi.IntPtrOutput `pulumi:"tokenMaxTtl"`
	// If true, the 'default' policy will not automatically be added to generated tokens
	TokenNoDefaultPolicy pulumi.BoolPtrOutput `pulumi:"tokenNoDefaultPolicy"`
	// The maximum number of times a token may be used, a value of zero means unlimited
	TokenNumUses pulumi.IntPtrOutput `pulumi:"tokenNumUses"`
	// Generated Token's Period
	TokenPeriod pulumi.IntPtrOutput `pulumi:"tokenPeriod"`
	// Generated Token's Policies
	TokenPolicies pulumi.StringArrayOutput `pulumi:"tokenPolicies"`
	// The initial ttl of the token to generate in seconds
	TokenTtl pulumi.IntPtrOutput `pulumi:"tokenTtl"`
	// The type of token to generate, service or batch
	TokenType pulumi.StringPtrOutput `pulumi:"tokenType"`
}

// NewCertAuthBackendRole registers a new resource with the given unique name, arguments, and options.
func NewCertAuthBackendRole(ctx *pulumi.Context,
	name string, args *CertAuthBackendRoleArgs, opts ...pulumi.ResourceOption) (*CertAuthBackendRole, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Certificate == nil {
		return nil, errors.New("invalid value for required argument 'Certificate'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource CertAuthBackendRole
	err := ctx.RegisterResource("vault:index/certAuthBackendRole:CertAuthBackendRole", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCertAuthBackendRole gets an existing CertAuthBackendRole resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCertAuthBackendRole(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CertAuthBackendRoleState, opts ...pulumi.ResourceOption) (*CertAuthBackendRole, error) {
	var resource CertAuthBackendRole
	err := ctx.ReadResource("vault:index/certAuthBackendRole:CertAuthBackendRole", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CertAuthBackendRole resources.
type certAuthBackendRoleState struct {
	// Allowed the common names for authenticated client certificates
	AllowedCommonNames []string `pulumi:"allowedCommonNames"`
	// Allowed alternative dns names for authenticated client certificates
	AllowedDnsSans []string `pulumi:"allowedDnsSans"`
	// Allowed emails for authenticated client certificates
	AllowedEmailSans []string `pulumi:"allowedEmailSans"`
	// DEPRECATED: Please use the individual `allowed_X_sans` parameters instead. Allowed subject names for authenticated client certificates
	AllowedNames []string `pulumi:"allowedNames"`
	// Allowed organization units for authenticated client certificates.
	AllowedOrganizationalUnits []string `pulumi:"allowedOrganizationalUnits"`
	// Allowed URIs for authenticated client certificates
	AllowedUriSans []string `pulumi:"allowedUriSans"`
	// Path to the mounted Cert auth backend
	Backend *string `pulumi:"backend"`
	// CA certificate used to validate client certificates
	Certificate *string `pulumi:"certificate"`
	// The name to display on tokens issued under this role.
	DisplayName *string `pulumi:"displayName"`
	// Name of the role
	Name *string `pulumi:"name"`
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
	// *Available only for Vault Enterprise*.
	Namespace *string `pulumi:"namespace"`
	// Any additional CA certificates
	// needed to verify OCSP responses. Provided as base64 encoded PEM data.
	// Requires Vault version 1.13+.
	OcspCaCertificates *string `pulumi:"ocspCaCertificates"`
	// If enabled, validate certificates'
	// revocation status using OCSP. Requires Vault version 1.13+.
	OcspEnabled *bool `pulumi:"ocspEnabled"`
	// If true and an OCSP response cannot
	// be fetched or is of an unknown status, the login will proceed as if the
	// certificate has not been revoked.
	// Requires Vault version 1.13+.
	OcspFailOpen *bool `pulumi:"ocspFailOpen"`
	// If set to true, rather than
	// accepting the first successful OCSP response, query all servers and consider
	// the certificate valid only if all servers agree.
	// Requires Vault version 1.13+.
	OcspQueryAllServers *bool `pulumi:"ocspQueryAllServers"`
	// : A comma-separated list of OCSP
	// server addresses. If unset, the OCSP server is determined from the
	// AuthorityInformationAccess extension on the certificate being inspected.
	// Requires Vault version 1.13+.
	OcspServersOverrides []string `pulumi:"ocspServersOverrides"`
	// TLS extensions required on
	// client certificates
	RequiredExtensions []string `pulumi:"requiredExtensions"`
	// Specifies the blocks of IP addresses which are allowed to use the generated token
	TokenBoundCidrs []string `pulumi:"tokenBoundCidrs"`
	// Generated Token's Explicit Maximum TTL in seconds
	TokenExplicitMaxTtl *int `pulumi:"tokenExplicitMaxTtl"`
	// The maximum lifetime of the generated token
	TokenMaxTtl *int `pulumi:"tokenMaxTtl"`
	// If true, the 'default' policy will not automatically be added to generated tokens
	TokenNoDefaultPolicy *bool `pulumi:"tokenNoDefaultPolicy"`
	// The maximum number of times a token may be used, a value of zero means unlimited
	TokenNumUses *int `pulumi:"tokenNumUses"`
	// Generated Token's Period
	TokenPeriod *int `pulumi:"tokenPeriod"`
	// Generated Token's Policies
	TokenPolicies []string `pulumi:"tokenPolicies"`
	// The initial ttl of the token to generate in seconds
	TokenTtl *int `pulumi:"tokenTtl"`
	// The type of token to generate, service or batch
	TokenType *string `pulumi:"tokenType"`
}

type CertAuthBackendRoleState struct {
	// Allowed the common names for authenticated client certificates
	AllowedCommonNames pulumi.StringArrayInput
	// Allowed alternative dns names for authenticated client certificates
	AllowedDnsSans pulumi.StringArrayInput
	// Allowed emails for authenticated client certificates
	AllowedEmailSans pulumi.StringArrayInput
	// DEPRECATED: Please use the individual `allowed_X_sans` parameters instead. Allowed subject names for authenticated client certificates
	AllowedNames pulumi.StringArrayInput
	// Allowed organization units for authenticated client certificates.
	AllowedOrganizationalUnits pulumi.StringArrayInput
	// Allowed URIs for authenticated client certificates
	AllowedUriSans pulumi.StringArrayInput
	// Path to the mounted Cert auth backend
	Backend pulumi.StringPtrInput
	// CA certificate used to validate client certificates
	Certificate pulumi.StringPtrInput
	// The name to display on tokens issued under this role.
	DisplayName pulumi.StringPtrInput
	// Name of the role
	Name pulumi.StringPtrInput
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
	// *Available only for Vault Enterprise*.
	Namespace pulumi.StringPtrInput
	// Any additional CA certificates
	// needed to verify OCSP responses. Provided as base64 encoded PEM data.
	// Requires Vault version 1.13+.
	OcspCaCertificates pulumi.StringPtrInput
	// If enabled, validate certificates'
	// revocation status using OCSP. Requires Vault version 1.13+.
	OcspEnabled pulumi.BoolPtrInput
	// If true and an OCSP response cannot
	// be fetched or is of an unknown status, the login will proceed as if the
	// certificate has not been revoked.
	// Requires Vault version 1.13+.
	OcspFailOpen pulumi.BoolPtrInput
	// If set to true, rather than
	// accepting the first successful OCSP response, query all servers and consider
	// the certificate valid only if all servers agree.
	// Requires Vault version 1.13+.
	OcspQueryAllServers pulumi.BoolPtrInput
	// : A comma-separated list of OCSP
	// server addresses. If unset, the OCSP server is determined from the
	// AuthorityInformationAccess extension on the certificate being inspected.
	// Requires Vault version 1.13+.
	OcspServersOverrides pulumi.StringArrayInput
	// TLS extensions required on
	// client certificates
	RequiredExtensions pulumi.StringArrayInput
	// Specifies the blocks of IP addresses which are allowed to use the generated token
	TokenBoundCidrs pulumi.StringArrayInput
	// Generated Token's Explicit Maximum TTL in seconds
	TokenExplicitMaxTtl pulumi.IntPtrInput
	// The maximum lifetime of the generated token
	TokenMaxTtl pulumi.IntPtrInput
	// If true, the 'default' policy will not automatically be added to generated tokens
	TokenNoDefaultPolicy pulumi.BoolPtrInput
	// The maximum number of times a token may be used, a value of zero means unlimited
	TokenNumUses pulumi.IntPtrInput
	// Generated Token's Period
	TokenPeriod pulumi.IntPtrInput
	// Generated Token's Policies
	TokenPolicies pulumi.StringArrayInput
	// The initial ttl of the token to generate in seconds
	TokenTtl pulumi.IntPtrInput
	// The type of token to generate, service or batch
	TokenType pulumi.StringPtrInput
}

func (CertAuthBackendRoleState) ElementType() reflect.Type {
	return reflect.TypeOf((*certAuthBackendRoleState)(nil)).Elem()
}

type certAuthBackendRoleArgs struct {
	// Allowed the common names for authenticated client certificates
	AllowedCommonNames []string `pulumi:"allowedCommonNames"`
	// Allowed alternative dns names for authenticated client certificates
	AllowedDnsSans []string `pulumi:"allowedDnsSans"`
	// Allowed emails for authenticated client certificates
	AllowedEmailSans []string `pulumi:"allowedEmailSans"`
	// DEPRECATED: Please use the individual `allowed_X_sans` parameters instead. Allowed subject names for authenticated client certificates
	AllowedNames []string `pulumi:"allowedNames"`
	// Allowed organization units for authenticated client certificates.
	AllowedOrganizationalUnits []string `pulumi:"allowedOrganizationalUnits"`
	// Allowed URIs for authenticated client certificates
	AllowedUriSans []string `pulumi:"allowedUriSans"`
	// Path to the mounted Cert auth backend
	Backend *string `pulumi:"backend"`
	// CA certificate used to validate client certificates
	Certificate string `pulumi:"certificate"`
	// The name to display on tokens issued under this role.
	DisplayName *string `pulumi:"displayName"`
	// Name of the role
	Name *string `pulumi:"name"`
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
	// *Available only for Vault Enterprise*.
	Namespace *string `pulumi:"namespace"`
	// Any additional CA certificates
	// needed to verify OCSP responses. Provided as base64 encoded PEM data.
	// Requires Vault version 1.13+.
	OcspCaCertificates *string `pulumi:"ocspCaCertificates"`
	// If enabled, validate certificates'
	// revocation status using OCSP. Requires Vault version 1.13+.
	OcspEnabled *bool `pulumi:"ocspEnabled"`
	// If true and an OCSP response cannot
	// be fetched or is of an unknown status, the login will proceed as if the
	// certificate has not been revoked.
	// Requires Vault version 1.13+.
	OcspFailOpen *bool `pulumi:"ocspFailOpen"`
	// If set to true, rather than
	// accepting the first successful OCSP response, query all servers and consider
	// the certificate valid only if all servers agree.
	// Requires Vault version 1.13+.
	OcspQueryAllServers *bool `pulumi:"ocspQueryAllServers"`
	// : A comma-separated list of OCSP
	// server addresses. If unset, the OCSP server is determined from the
	// AuthorityInformationAccess extension on the certificate being inspected.
	// Requires Vault version 1.13+.
	OcspServersOverrides []string `pulumi:"ocspServersOverrides"`
	// TLS extensions required on
	// client certificates
	RequiredExtensions []string `pulumi:"requiredExtensions"`
	// Specifies the blocks of IP addresses which are allowed to use the generated token
	TokenBoundCidrs []string `pulumi:"tokenBoundCidrs"`
	// Generated Token's Explicit Maximum TTL in seconds
	TokenExplicitMaxTtl *int `pulumi:"tokenExplicitMaxTtl"`
	// The maximum lifetime of the generated token
	TokenMaxTtl *int `pulumi:"tokenMaxTtl"`
	// If true, the 'default' policy will not automatically be added to generated tokens
	TokenNoDefaultPolicy *bool `pulumi:"tokenNoDefaultPolicy"`
	// The maximum number of times a token may be used, a value of zero means unlimited
	TokenNumUses *int `pulumi:"tokenNumUses"`
	// Generated Token's Period
	TokenPeriod *int `pulumi:"tokenPeriod"`
	// Generated Token's Policies
	TokenPolicies []string `pulumi:"tokenPolicies"`
	// The initial ttl of the token to generate in seconds
	TokenTtl *int `pulumi:"tokenTtl"`
	// The type of token to generate, service or batch
	TokenType *string `pulumi:"tokenType"`
}

// The set of arguments for constructing a CertAuthBackendRole resource.
type CertAuthBackendRoleArgs struct {
	// Allowed the common names for authenticated client certificates
	AllowedCommonNames pulumi.StringArrayInput
	// Allowed alternative dns names for authenticated client certificates
	AllowedDnsSans pulumi.StringArrayInput
	// Allowed emails for authenticated client certificates
	AllowedEmailSans pulumi.StringArrayInput
	// DEPRECATED: Please use the individual `allowed_X_sans` parameters instead. Allowed subject names for authenticated client certificates
	AllowedNames pulumi.StringArrayInput
	// Allowed organization units for authenticated client certificates.
	AllowedOrganizationalUnits pulumi.StringArrayInput
	// Allowed URIs for authenticated client certificates
	AllowedUriSans pulumi.StringArrayInput
	// Path to the mounted Cert auth backend
	Backend pulumi.StringPtrInput
	// CA certificate used to validate client certificates
	Certificate pulumi.StringInput
	// The name to display on tokens issued under this role.
	DisplayName pulumi.StringPtrInput
	// Name of the role
	Name pulumi.StringPtrInput
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
	// *Available only for Vault Enterprise*.
	Namespace pulumi.StringPtrInput
	// Any additional CA certificates
	// needed to verify OCSP responses. Provided as base64 encoded PEM data.
	// Requires Vault version 1.13+.
	OcspCaCertificates pulumi.StringPtrInput
	// If enabled, validate certificates'
	// revocation status using OCSP. Requires Vault version 1.13+.
	OcspEnabled pulumi.BoolPtrInput
	// If true and an OCSP response cannot
	// be fetched or is of an unknown status, the login will proceed as if the
	// certificate has not been revoked.
	// Requires Vault version 1.13+.
	OcspFailOpen pulumi.BoolPtrInput
	// If set to true, rather than
	// accepting the first successful OCSP response, query all servers and consider
	// the certificate valid only if all servers agree.
	// Requires Vault version 1.13+.
	OcspQueryAllServers pulumi.BoolPtrInput
	// : A comma-separated list of OCSP
	// server addresses. If unset, the OCSP server is determined from the
	// AuthorityInformationAccess extension on the certificate being inspected.
	// Requires Vault version 1.13+.
	OcspServersOverrides pulumi.StringArrayInput
	// TLS extensions required on
	// client certificates
	RequiredExtensions pulumi.StringArrayInput
	// Specifies the blocks of IP addresses which are allowed to use the generated token
	TokenBoundCidrs pulumi.StringArrayInput
	// Generated Token's Explicit Maximum TTL in seconds
	TokenExplicitMaxTtl pulumi.IntPtrInput
	// The maximum lifetime of the generated token
	TokenMaxTtl pulumi.IntPtrInput
	// If true, the 'default' policy will not automatically be added to generated tokens
	TokenNoDefaultPolicy pulumi.BoolPtrInput
	// The maximum number of times a token may be used, a value of zero means unlimited
	TokenNumUses pulumi.IntPtrInput
	// Generated Token's Period
	TokenPeriod pulumi.IntPtrInput
	// Generated Token's Policies
	TokenPolicies pulumi.StringArrayInput
	// The initial ttl of the token to generate in seconds
	TokenTtl pulumi.IntPtrInput
	// The type of token to generate, service or batch
	TokenType pulumi.StringPtrInput
}

func (CertAuthBackendRoleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*certAuthBackendRoleArgs)(nil)).Elem()
}

type CertAuthBackendRoleInput interface {
	pulumi.Input

	ToCertAuthBackendRoleOutput() CertAuthBackendRoleOutput
	ToCertAuthBackendRoleOutputWithContext(ctx context.Context) CertAuthBackendRoleOutput
}

func (*CertAuthBackendRole) ElementType() reflect.Type {
	return reflect.TypeOf((**CertAuthBackendRole)(nil)).Elem()
}

func (i *CertAuthBackendRole) ToCertAuthBackendRoleOutput() CertAuthBackendRoleOutput {
	return i.ToCertAuthBackendRoleOutputWithContext(context.Background())
}

func (i *CertAuthBackendRole) ToCertAuthBackendRoleOutputWithContext(ctx context.Context) CertAuthBackendRoleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CertAuthBackendRoleOutput)
}

// CertAuthBackendRoleArrayInput is an input type that accepts CertAuthBackendRoleArray and CertAuthBackendRoleArrayOutput values.
// You can construct a concrete instance of `CertAuthBackendRoleArrayInput` via:
//
//	CertAuthBackendRoleArray{ CertAuthBackendRoleArgs{...} }
type CertAuthBackendRoleArrayInput interface {
	pulumi.Input

	ToCertAuthBackendRoleArrayOutput() CertAuthBackendRoleArrayOutput
	ToCertAuthBackendRoleArrayOutputWithContext(context.Context) CertAuthBackendRoleArrayOutput
}

type CertAuthBackendRoleArray []CertAuthBackendRoleInput

func (CertAuthBackendRoleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CertAuthBackendRole)(nil)).Elem()
}

func (i CertAuthBackendRoleArray) ToCertAuthBackendRoleArrayOutput() CertAuthBackendRoleArrayOutput {
	return i.ToCertAuthBackendRoleArrayOutputWithContext(context.Background())
}

func (i CertAuthBackendRoleArray) ToCertAuthBackendRoleArrayOutputWithContext(ctx context.Context) CertAuthBackendRoleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CertAuthBackendRoleArrayOutput)
}

// CertAuthBackendRoleMapInput is an input type that accepts CertAuthBackendRoleMap and CertAuthBackendRoleMapOutput values.
// You can construct a concrete instance of `CertAuthBackendRoleMapInput` via:
//
//	CertAuthBackendRoleMap{ "key": CertAuthBackendRoleArgs{...} }
type CertAuthBackendRoleMapInput interface {
	pulumi.Input

	ToCertAuthBackendRoleMapOutput() CertAuthBackendRoleMapOutput
	ToCertAuthBackendRoleMapOutputWithContext(context.Context) CertAuthBackendRoleMapOutput
}

type CertAuthBackendRoleMap map[string]CertAuthBackendRoleInput

func (CertAuthBackendRoleMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CertAuthBackendRole)(nil)).Elem()
}

func (i CertAuthBackendRoleMap) ToCertAuthBackendRoleMapOutput() CertAuthBackendRoleMapOutput {
	return i.ToCertAuthBackendRoleMapOutputWithContext(context.Background())
}

func (i CertAuthBackendRoleMap) ToCertAuthBackendRoleMapOutputWithContext(ctx context.Context) CertAuthBackendRoleMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CertAuthBackendRoleMapOutput)
}

type CertAuthBackendRoleOutput struct{ *pulumi.OutputState }

func (CertAuthBackendRoleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CertAuthBackendRole)(nil)).Elem()
}

func (o CertAuthBackendRoleOutput) ToCertAuthBackendRoleOutput() CertAuthBackendRoleOutput {
	return o
}

func (o CertAuthBackendRoleOutput) ToCertAuthBackendRoleOutputWithContext(ctx context.Context) CertAuthBackendRoleOutput {
	return o
}

// Allowed the common names for authenticated client certificates
func (o CertAuthBackendRoleOutput) AllowedCommonNames() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *CertAuthBackendRole) pulumi.StringArrayOutput { return v.AllowedCommonNames }).(pulumi.StringArrayOutput)
}

// Allowed alternative dns names for authenticated client certificates
func (o CertAuthBackendRoleOutput) AllowedDnsSans() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *CertAuthBackendRole) pulumi.StringArrayOutput { return v.AllowedDnsSans }).(pulumi.StringArrayOutput)
}

// Allowed emails for authenticated client certificates
func (o CertAuthBackendRoleOutput) AllowedEmailSans() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *CertAuthBackendRole) pulumi.StringArrayOutput { return v.AllowedEmailSans }).(pulumi.StringArrayOutput)
}

// DEPRECATED: Please use the individual `allowed_X_sans` parameters instead. Allowed subject names for authenticated client certificates
func (o CertAuthBackendRoleOutput) AllowedNames() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *CertAuthBackendRole) pulumi.StringArrayOutput { return v.AllowedNames }).(pulumi.StringArrayOutput)
}

// Allowed organization units for authenticated client certificates.
func (o CertAuthBackendRoleOutput) AllowedOrganizationalUnits() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *CertAuthBackendRole) pulumi.StringArrayOutput { return v.AllowedOrganizationalUnits }).(pulumi.StringArrayOutput)
}

// Allowed URIs for authenticated client certificates
func (o CertAuthBackendRoleOutput) AllowedUriSans() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *CertAuthBackendRole) pulumi.StringArrayOutput { return v.AllowedUriSans }).(pulumi.StringArrayOutput)
}

// Path to the mounted Cert auth backend
func (o CertAuthBackendRoleOutput) Backend() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CertAuthBackendRole) pulumi.StringPtrOutput { return v.Backend }).(pulumi.StringPtrOutput)
}

// CA certificate used to validate client certificates
func (o CertAuthBackendRoleOutput) Certificate() pulumi.StringOutput {
	return o.ApplyT(func(v *CertAuthBackendRole) pulumi.StringOutput { return v.Certificate }).(pulumi.StringOutput)
}

// The name to display on tokens issued under this role.
func (o CertAuthBackendRoleOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v *CertAuthBackendRole) pulumi.StringOutput { return v.DisplayName }).(pulumi.StringOutput)
}

// Name of the role
func (o CertAuthBackendRoleOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *CertAuthBackendRole) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The namespace to provision the resource in.
// The value should not contain leading or trailing forward slashes.
// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
// *Available only for Vault Enterprise*.
func (o CertAuthBackendRoleOutput) Namespace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CertAuthBackendRole) pulumi.StringPtrOutput { return v.Namespace }).(pulumi.StringPtrOutput)
}

// Any additional CA certificates
// needed to verify OCSP responses. Provided as base64 encoded PEM data.
// Requires Vault version 1.13+.
func (o CertAuthBackendRoleOutput) OcspCaCertificates() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CertAuthBackendRole) pulumi.StringPtrOutput { return v.OcspCaCertificates }).(pulumi.StringPtrOutput)
}

// If enabled, validate certificates'
// revocation status using OCSP. Requires Vault version 1.13+.
func (o CertAuthBackendRoleOutput) OcspEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v *CertAuthBackendRole) pulumi.BoolOutput { return v.OcspEnabled }).(pulumi.BoolOutput)
}

// If true and an OCSP response cannot
// be fetched or is of an unknown status, the login will proceed as if the
// certificate has not been revoked.
// Requires Vault version 1.13+.
func (o CertAuthBackendRoleOutput) OcspFailOpen() pulumi.BoolOutput {
	return o.ApplyT(func(v *CertAuthBackendRole) pulumi.BoolOutput { return v.OcspFailOpen }).(pulumi.BoolOutput)
}

// If set to true, rather than
// accepting the first successful OCSP response, query all servers and consider
// the certificate valid only if all servers agree.
// Requires Vault version 1.13+.
func (o CertAuthBackendRoleOutput) OcspQueryAllServers() pulumi.BoolOutput {
	return o.ApplyT(func(v *CertAuthBackendRole) pulumi.BoolOutput { return v.OcspQueryAllServers }).(pulumi.BoolOutput)
}

// : A comma-separated list of OCSP
// server addresses. If unset, the OCSP server is determined from the
// AuthorityInformationAccess extension on the certificate being inspected.
// Requires Vault version 1.13+.
func (o CertAuthBackendRoleOutput) OcspServersOverrides() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *CertAuthBackendRole) pulumi.StringArrayOutput { return v.OcspServersOverrides }).(pulumi.StringArrayOutput)
}

// TLS extensions required on
// client certificates
func (o CertAuthBackendRoleOutput) RequiredExtensions() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *CertAuthBackendRole) pulumi.StringArrayOutput { return v.RequiredExtensions }).(pulumi.StringArrayOutput)
}

// Specifies the blocks of IP addresses which are allowed to use the generated token
func (o CertAuthBackendRoleOutput) TokenBoundCidrs() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *CertAuthBackendRole) pulumi.StringArrayOutput { return v.TokenBoundCidrs }).(pulumi.StringArrayOutput)
}

// Generated Token's Explicit Maximum TTL in seconds
func (o CertAuthBackendRoleOutput) TokenExplicitMaxTtl() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *CertAuthBackendRole) pulumi.IntPtrOutput { return v.TokenExplicitMaxTtl }).(pulumi.IntPtrOutput)
}

// The maximum lifetime of the generated token
func (o CertAuthBackendRoleOutput) TokenMaxTtl() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *CertAuthBackendRole) pulumi.IntPtrOutput { return v.TokenMaxTtl }).(pulumi.IntPtrOutput)
}

// If true, the 'default' policy will not automatically be added to generated tokens
func (o CertAuthBackendRoleOutput) TokenNoDefaultPolicy() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *CertAuthBackendRole) pulumi.BoolPtrOutput { return v.TokenNoDefaultPolicy }).(pulumi.BoolPtrOutput)
}

// The maximum number of times a token may be used, a value of zero means unlimited
func (o CertAuthBackendRoleOutput) TokenNumUses() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *CertAuthBackendRole) pulumi.IntPtrOutput { return v.TokenNumUses }).(pulumi.IntPtrOutput)
}

// Generated Token's Period
func (o CertAuthBackendRoleOutput) TokenPeriod() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *CertAuthBackendRole) pulumi.IntPtrOutput { return v.TokenPeriod }).(pulumi.IntPtrOutput)
}

// Generated Token's Policies
func (o CertAuthBackendRoleOutput) TokenPolicies() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *CertAuthBackendRole) pulumi.StringArrayOutput { return v.TokenPolicies }).(pulumi.StringArrayOutput)
}

// The initial ttl of the token to generate in seconds
func (o CertAuthBackendRoleOutput) TokenTtl() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *CertAuthBackendRole) pulumi.IntPtrOutput { return v.TokenTtl }).(pulumi.IntPtrOutput)
}

// The type of token to generate, service or batch
func (o CertAuthBackendRoleOutput) TokenType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CertAuthBackendRole) pulumi.StringPtrOutput { return v.TokenType }).(pulumi.StringPtrOutput)
}

type CertAuthBackendRoleArrayOutput struct{ *pulumi.OutputState }

func (CertAuthBackendRoleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CertAuthBackendRole)(nil)).Elem()
}

func (o CertAuthBackendRoleArrayOutput) ToCertAuthBackendRoleArrayOutput() CertAuthBackendRoleArrayOutput {
	return o
}

func (o CertAuthBackendRoleArrayOutput) ToCertAuthBackendRoleArrayOutputWithContext(ctx context.Context) CertAuthBackendRoleArrayOutput {
	return o
}

func (o CertAuthBackendRoleArrayOutput) Index(i pulumi.IntInput) CertAuthBackendRoleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *CertAuthBackendRole {
		return vs[0].([]*CertAuthBackendRole)[vs[1].(int)]
	}).(CertAuthBackendRoleOutput)
}

type CertAuthBackendRoleMapOutput struct{ *pulumi.OutputState }

func (CertAuthBackendRoleMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CertAuthBackendRole)(nil)).Elem()
}

func (o CertAuthBackendRoleMapOutput) ToCertAuthBackendRoleMapOutput() CertAuthBackendRoleMapOutput {
	return o
}

func (o CertAuthBackendRoleMapOutput) ToCertAuthBackendRoleMapOutputWithContext(ctx context.Context) CertAuthBackendRoleMapOutput {
	return o
}

func (o CertAuthBackendRoleMapOutput) MapIndex(k pulumi.StringInput) CertAuthBackendRoleOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *CertAuthBackendRole {
		return vs[0].(map[string]*CertAuthBackendRole)[vs[1].(string)]
	}).(CertAuthBackendRoleOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*CertAuthBackendRoleInput)(nil)).Elem(), &CertAuthBackendRole{})
	pulumi.RegisterInputType(reflect.TypeOf((*CertAuthBackendRoleArrayInput)(nil)).Elem(), CertAuthBackendRoleArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*CertAuthBackendRoleMapInput)(nil)).Elem(), CertAuthBackendRoleMap{})
	pulumi.RegisterOutputType(CertAuthBackendRoleOutput{})
	pulumi.RegisterOutputType(CertAuthBackendRoleArrayOutput{})
	pulumi.RegisterOutputType(CertAuthBackendRoleMapOutput{})
}
