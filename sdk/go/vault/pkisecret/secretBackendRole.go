// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package pkisecret

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Creates a role on an PKI Secret Backend for Vault.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"fmt"
//
// 	"github.com/pulumi/pulumi-vault/sdk/v2/go/vault/pkiSecret"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		pki, err := pkiSecret.NewSecretBackend(ctx, "pki", &pkiSecret.SecretBackendArgs{
// 			DefaultLeaseTtlSeconds: pulumi.Int(3600),
// 			MaxLeaseTtlSeconds:     pulumi.Int(86400),
// 			Path:                   pulumi.String(fmt.Sprintf("%v%v", "%", "s")),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = pkiSecret.NewSecretBackendRole(ctx, "role", &pkiSecret.SecretBackendRoleArgs{
// 			Backend: pki.Path,
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type SecretBackendRole struct {
	pulumi.CustomResourceState

	// Flag to allow any name
	AllowAnyName pulumi.BoolPtrOutput `pulumi:"allowAnyName"`
	// Flag to allow certificates matching the actual domain
	AllowBareDomains pulumi.BoolPtrOutput `pulumi:"allowBareDomains"`
	// Flag to allow names containing glob patterns.
	AllowGlobDomains pulumi.BoolPtrOutput `pulumi:"allowGlobDomains"`
	// Flag to allow IP SANs
	AllowIpSans pulumi.BoolPtrOutput `pulumi:"allowIpSans"`
	// Flag to allow certificates for localhost
	AllowLocalhost pulumi.BoolPtrOutput `pulumi:"allowLocalhost"`
	// Flag to allow certificates matching subdomains
	AllowSubdomains pulumi.BoolPtrOutput `pulumi:"allowSubdomains"`
	// List of allowed domains for certificates
	AllowedDomains pulumi.StringArrayOutput `pulumi:"allowedDomains"`
	// Defines allowed custom SANs
	AllowedOtherSans pulumi.StringArrayOutput `pulumi:"allowedOtherSans"`
	// Defines allowed URI SANs
	AllowedUriSans pulumi.StringArrayOutput `pulumi:"allowedUriSans"`
	// The path the PKI secret backend is mounted at, with no leading or trailing `/`s.
	Backend pulumi.StringOutput `pulumi:"backend"`
	// Flag to mark basic constraints valid when issuing non-CA certificates
	BasicConstraintsValidForNonCa pulumi.BoolPtrOutput `pulumi:"basicConstraintsValidForNonCa"`
	// Flag to specify certificates for client use
	ClientFlag pulumi.BoolPtrOutput `pulumi:"clientFlag"`
	// Flag to specify certificates for code signing use
	CodeSigningFlag pulumi.BoolPtrOutput `pulumi:"codeSigningFlag"`
	// The country of generated certificates
	Countries pulumi.StringArrayOutput `pulumi:"countries"`
	// Flag to specify certificates for email protection use
	EmailProtectionFlag pulumi.BoolPtrOutput `pulumi:"emailProtectionFlag"`
	// Flag to allow only valid host names
	EnforceHostnames pulumi.BoolPtrOutput `pulumi:"enforceHostnames"`
	// Specify the allowed extended key usage constraint on issued certificates
	ExtKeyUsages pulumi.StringArrayOutput `pulumi:"extKeyUsages"`
	// Flag to generate leases with certificates
	GenerateLease pulumi.BoolPtrOutput `pulumi:"generateLease"`
	// The number of bits of generated keys
	KeyBits pulumi.IntPtrOutput `pulumi:"keyBits"`
	// The type of generated keys
	KeyType pulumi.StringPtrOutput `pulumi:"keyType"`
	// Specify the allowed key usage constraint on issued certificates
	KeyUsages pulumi.StringArrayOutput `pulumi:"keyUsages"`
	// The locality of generated certificates
	Localities pulumi.StringArrayOutput `pulumi:"localities"`
	// The maximum TTL
	MaxTtl pulumi.StringPtrOutput `pulumi:"maxTtl"`
	// The name to identify this role within the backend. Must be unique within the backend.
	Name pulumi.StringOutput `pulumi:"name"`
	// Flag to not store certificates in the storage backend
	NoStore pulumi.BoolPtrOutput `pulumi:"noStore"`
	// Specifies the duration by which to backdate the NotBefore property.
	NotBeforeDuration pulumi.StringOutput `pulumi:"notBeforeDuration"`
	// The organization unit of generated certificates
	OrganizationUnit pulumi.StringArrayOutput `pulumi:"organizationUnit"`
	// The organization of generated certificates
	Organizations pulumi.StringArrayOutput `pulumi:"organizations"`
	// Specify the list of allowed policies IODs
	PolicyIdentifiers pulumi.StringArrayOutput `pulumi:"policyIdentifiers"`
	// The postal code of generated certificates
	PostalCodes pulumi.StringArrayOutput `pulumi:"postalCodes"`
	// The province of generated certificates
	Provinces pulumi.StringArrayOutput `pulumi:"provinces"`
	// Flag to force CN usage
	RequireCn pulumi.BoolPtrOutput `pulumi:"requireCn"`
	// Flag to specify certificates for server use
	ServerFlag pulumi.BoolPtrOutput `pulumi:"serverFlag"`
	// The street address of generated certificates
	StreetAddresses pulumi.StringArrayOutput `pulumi:"streetAddresses"`
	// The TTL
	Ttl pulumi.StringPtrOutput `pulumi:"ttl"`
	// Flag to use the CN in the CSR
	UseCsrCommonName pulumi.BoolPtrOutput `pulumi:"useCsrCommonName"`
	// Flag to use the SANs in the CSR
	UseCsrSans pulumi.BoolPtrOutput `pulumi:"useCsrSans"`
}

// NewSecretBackendRole registers a new resource with the given unique name, arguments, and options.
func NewSecretBackendRole(ctx *pulumi.Context,
	name string, args *SecretBackendRoleArgs, opts ...pulumi.ResourceOption) (*SecretBackendRole, error) {
	if args == nil || args.Backend == nil {
		return nil, errors.New("missing required argument 'Backend'")
	}
	if args == nil {
		args = &SecretBackendRoleArgs{}
	}
	var resource SecretBackendRole
	err := ctx.RegisterResource("vault:pkiSecret/secretBackendRole:SecretBackendRole", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSecretBackendRole gets an existing SecretBackendRole resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSecretBackendRole(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SecretBackendRoleState, opts ...pulumi.ResourceOption) (*SecretBackendRole, error) {
	var resource SecretBackendRole
	err := ctx.ReadResource("vault:pkiSecret/secretBackendRole:SecretBackendRole", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SecretBackendRole resources.
type secretBackendRoleState struct {
	// Flag to allow any name
	AllowAnyName *bool `pulumi:"allowAnyName"`
	// Flag to allow certificates matching the actual domain
	AllowBareDomains *bool `pulumi:"allowBareDomains"`
	// Flag to allow names containing glob patterns.
	AllowGlobDomains *bool `pulumi:"allowGlobDomains"`
	// Flag to allow IP SANs
	AllowIpSans *bool `pulumi:"allowIpSans"`
	// Flag to allow certificates for localhost
	AllowLocalhost *bool `pulumi:"allowLocalhost"`
	// Flag to allow certificates matching subdomains
	AllowSubdomains *bool `pulumi:"allowSubdomains"`
	// List of allowed domains for certificates
	AllowedDomains []string `pulumi:"allowedDomains"`
	// Defines allowed custom SANs
	AllowedOtherSans []string `pulumi:"allowedOtherSans"`
	// Defines allowed URI SANs
	AllowedUriSans []string `pulumi:"allowedUriSans"`
	// The path the PKI secret backend is mounted at, with no leading or trailing `/`s.
	Backend *string `pulumi:"backend"`
	// Flag to mark basic constraints valid when issuing non-CA certificates
	BasicConstraintsValidForNonCa *bool `pulumi:"basicConstraintsValidForNonCa"`
	// Flag to specify certificates for client use
	ClientFlag *bool `pulumi:"clientFlag"`
	// Flag to specify certificates for code signing use
	CodeSigningFlag *bool `pulumi:"codeSigningFlag"`
	// The country of generated certificates
	Countries []string `pulumi:"countries"`
	// Flag to specify certificates for email protection use
	EmailProtectionFlag *bool `pulumi:"emailProtectionFlag"`
	// Flag to allow only valid host names
	EnforceHostnames *bool `pulumi:"enforceHostnames"`
	// Specify the allowed extended key usage constraint on issued certificates
	ExtKeyUsages []string `pulumi:"extKeyUsages"`
	// Flag to generate leases with certificates
	GenerateLease *bool `pulumi:"generateLease"`
	// The number of bits of generated keys
	KeyBits *int `pulumi:"keyBits"`
	// The type of generated keys
	KeyType *string `pulumi:"keyType"`
	// Specify the allowed key usage constraint on issued certificates
	KeyUsages []string `pulumi:"keyUsages"`
	// The locality of generated certificates
	Localities []string `pulumi:"localities"`
	// The maximum TTL
	MaxTtl *string `pulumi:"maxTtl"`
	// The name to identify this role within the backend. Must be unique within the backend.
	Name *string `pulumi:"name"`
	// Flag to not store certificates in the storage backend
	NoStore *bool `pulumi:"noStore"`
	// Specifies the duration by which to backdate the NotBefore property.
	NotBeforeDuration *string `pulumi:"notBeforeDuration"`
	// The organization unit of generated certificates
	OrganizationUnit []string `pulumi:"organizationUnit"`
	// The organization of generated certificates
	Organizations []string `pulumi:"organizations"`
	// Specify the list of allowed policies IODs
	PolicyIdentifiers []string `pulumi:"policyIdentifiers"`
	// The postal code of generated certificates
	PostalCodes []string `pulumi:"postalCodes"`
	// The province of generated certificates
	Provinces []string `pulumi:"provinces"`
	// Flag to force CN usage
	RequireCn *bool `pulumi:"requireCn"`
	// Flag to specify certificates for server use
	ServerFlag *bool `pulumi:"serverFlag"`
	// The street address of generated certificates
	StreetAddresses []string `pulumi:"streetAddresses"`
	// The TTL
	Ttl *string `pulumi:"ttl"`
	// Flag to use the CN in the CSR
	UseCsrCommonName *bool `pulumi:"useCsrCommonName"`
	// Flag to use the SANs in the CSR
	UseCsrSans *bool `pulumi:"useCsrSans"`
}

type SecretBackendRoleState struct {
	// Flag to allow any name
	AllowAnyName pulumi.BoolPtrInput
	// Flag to allow certificates matching the actual domain
	AllowBareDomains pulumi.BoolPtrInput
	// Flag to allow names containing glob patterns.
	AllowGlobDomains pulumi.BoolPtrInput
	// Flag to allow IP SANs
	AllowIpSans pulumi.BoolPtrInput
	// Flag to allow certificates for localhost
	AllowLocalhost pulumi.BoolPtrInput
	// Flag to allow certificates matching subdomains
	AllowSubdomains pulumi.BoolPtrInput
	// List of allowed domains for certificates
	AllowedDomains pulumi.StringArrayInput
	// Defines allowed custom SANs
	AllowedOtherSans pulumi.StringArrayInput
	// Defines allowed URI SANs
	AllowedUriSans pulumi.StringArrayInput
	// The path the PKI secret backend is mounted at, with no leading or trailing `/`s.
	Backend pulumi.StringPtrInput
	// Flag to mark basic constraints valid when issuing non-CA certificates
	BasicConstraintsValidForNonCa pulumi.BoolPtrInput
	// Flag to specify certificates for client use
	ClientFlag pulumi.BoolPtrInput
	// Flag to specify certificates for code signing use
	CodeSigningFlag pulumi.BoolPtrInput
	// The country of generated certificates
	Countries pulumi.StringArrayInput
	// Flag to specify certificates for email protection use
	EmailProtectionFlag pulumi.BoolPtrInput
	// Flag to allow only valid host names
	EnforceHostnames pulumi.BoolPtrInput
	// Specify the allowed extended key usage constraint on issued certificates
	ExtKeyUsages pulumi.StringArrayInput
	// Flag to generate leases with certificates
	GenerateLease pulumi.BoolPtrInput
	// The number of bits of generated keys
	KeyBits pulumi.IntPtrInput
	// The type of generated keys
	KeyType pulumi.StringPtrInput
	// Specify the allowed key usage constraint on issued certificates
	KeyUsages pulumi.StringArrayInput
	// The locality of generated certificates
	Localities pulumi.StringArrayInput
	// The maximum TTL
	MaxTtl pulumi.StringPtrInput
	// The name to identify this role within the backend. Must be unique within the backend.
	Name pulumi.StringPtrInput
	// Flag to not store certificates in the storage backend
	NoStore pulumi.BoolPtrInput
	// Specifies the duration by which to backdate the NotBefore property.
	NotBeforeDuration pulumi.StringPtrInput
	// The organization unit of generated certificates
	OrganizationUnit pulumi.StringArrayInput
	// The organization of generated certificates
	Organizations pulumi.StringArrayInput
	// Specify the list of allowed policies IODs
	PolicyIdentifiers pulumi.StringArrayInput
	// The postal code of generated certificates
	PostalCodes pulumi.StringArrayInput
	// The province of generated certificates
	Provinces pulumi.StringArrayInput
	// Flag to force CN usage
	RequireCn pulumi.BoolPtrInput
	// Flag to specify certificates for server use
	ServerFlag pulumi.BoolPtrInput
	// The street address of generated certificates
	StreetAddresses pulumi.StringArrayInput
	// The TTL
	Ttl pulumi.StringPtrInput
	// Flag to use the CN in the CSR
	UseCsrCommonName pulumi.BoolPtrInput
	// Flag to use the SANs in the CSR
	UseCsrSans pulumi.BoolPtrInput
}

func (SecretBackendRoleState) ElementType() reflect.Type {
	return reflect.TypeOf((*secretBackendRoleState)(nil)).Elem()
}

type secretBackendRoleArgs struct {
	// Flag to allow any name
	AllowAnyName *bool `pulumi:"allowAnyName"`
	// Flag to allow certificates matching the actual domain
	AllowBareDomains *bool `pulumi:"allowBareDomains"`
	// Flag to allow names containing glob patterns.
	AllowGlobDomains *bool `pulumi:"allowGlobDomains"`
	// Flag to allow IP SANs
	AllowIpSans *bool `pulumi:"allowIpSans"`
	// Flag to allow certificates for localhost
	AllowLocalhost *bool `pulumi:"allowLocalhost"`
	// Flag to allow certificates matching subdomains
	AllowSubdomains *bool `pulumi:"allowSubdomains"`
	// List of allowed domains for certificates
	AllowedDomains []string `pulumi:"allowedDomains"`
	// Defines allowed custom SANs
	AllowedOtherSans []string `pulumi:"allowedOtherSans"`
	// Defines allowed URI SANs
	AllowedUriSans []string `pulumi:"allowedUriSans"`
	// The path the PKI secret backend is mounted at, with no leading or trailing `/`s.
	Backend string `pulumi:"backend"`
	// Flag to mark basic constraints valid when issuing non-CA certificates
	BasicConstraintsValidForNonCa *bool `pulumi:"basicConstraintsValidForNonCa"`
	// Flag to specify certificates for client use
	ClientFlag *bool `pulumi:"clientFlag"`
	// Flag to specify certificates for code signing use
	CodeSigningFlag *bool `pulumi:"codeSigningFlag"`
	// The country of generated certificates
	Countries []string `pulumi:"countries"`
	// Flag to specify certificates for email protection use
	EmailProtectionFlag *bool `pulumi:"emailProtectionFlag"`
	// Flag to allow only valid host names
	EnforceHostnames *bool `pulumi:"enforceHostnames"`
	// Specify the allowed extended key usage constraint on issued certificates
	ExtKeyUsages []string `pulumi:"extKeyUsages"`
	// Flag to generate leases with certificates
	GenerateLease *bool `pulumi:"generateLease"`
	// The number of bits of generated keys
	KeyBits *int `pulumi:"keyBits"`
	// The type of generated keys
	KeyType *string `pulumi:"keyType"`
	// Specify the allowed key usage constraint on issued certificates
	KeyUsages []string `pulumi:"keyUsages"`
	// The locality of generated certificates
	Localities []string `pulumi:"localities"`
	// The maximum TTL
	MaxTtl *string `pulumi:"maxTtl"`
	// The name to identify this role within the backend. Must be unique within the backend.
	Name *string `pulumi:"name"`
	// Flag to not store certificates in the storage backend
	NoStore *bool `pulumi:"noStore"`
	// Specifies the duration by which to backdate the NotBefore property.
	NotBeforeDuration *string `pulumi:"notBeforeDuration"`
	// The organization unit of generated certificates
	OrganizationUnit []string `pulumi:"organizationUnit"`
	// The organization of generated certificates
	Organizations []string `pulumi:"organizations"`
	// Specify the list of allowed policies IODs
	PolicyIdentifiers []string `pulumi:"policyIdentifiers"`
	// The postal code of generated certificates
	PostalCodes []string `pulumi:"postalCodes"`
	// The province of generated certificates
	Provinces []string `pulumi:"provinces"`
	// Flag to force CN usage
	RequireCn *bool `pulumi:"requireCn"`
	// Flag to specify certificates for server use
	ServerFlag *bool `pulumi:"serverFlag"`
	// The street address of generated certificates
	StreetAddresses []string `pulumi:"streetAddresses"`
	// The TTL
	Ttl *string `pulumi:"ttl"`
	// Flag to use the CN in the CSR
	UseCsrCommonName *bool `pulumi:"useCsrCommonName"`
	// Flag to use the SANs in the CSR
	UseCsrSans *bool `pulumi:"useCsrSans"`
}

// The set of arguments for constructing a SecretBackendRole resource.
type SecretBackendRoleArgs struct {
	// Flag to allow any name
	AllowAnyName pulumi.BoolPtrInput
	// Flag to allow certificates matching the actual domain
	AllowBareDomains pulumi.BoolPtrInput
	// Flag to allow names containing glob patterns.
	AllowGlobDomains pulumi.BoolPtrInput
	// Flag to allow IP SANs
	AllowIpSans pulumi.BoolPtrInput
	// Flag to allow certificates for localhost
	AllowLocalhost pulumi.BoolPtrInput
	// Flag to allow certificates matching subdomains
	AllowSubdomains pulumi.BoolPtrInput
	// List of allowed domains for certificates
	AllowedDomains pulumi.StringArrayInput
	// Defines allowed custom SANs
	AllowedOtherSans pulumi.StringArrayInput
	// Defines allowed URI SANs
	AllowedUriSans pulumi.StringArrayInput
	// The path the PKI secret backend is mounted at, with no leading or trailing `/`s.
	Backend pulumi.StringInput
	// Flag to mark basic constraints valid when issuing non-CA certificates
	BasicConstraintsValidForNonCa pulumi.BoolPtrInput
	// Flag to specify certificates for client use
	ClientFlag pulumi.BoolPtrInput
	// Flag to specify certificates for code signing use
	CodeSigningFlag pulumi.BoolPtrInput
	// The country of generated certificates
	Countries pulumi.StringArrayInput
	// Flag to specify certificates for email protection use
	EmailProtectionFlag pulumi.BoolPtrInput
	// Flag to allow only valid host names
	EnforceHostnames pulumi.BoolPtrInput
	// Specify the allowed extended key usage constraint on issued certificates
	ExtKeyUsages pulumi.StringArrayInput
	// Flag to generate leases with certificates
	GenerateLease pulumi.BoolPtrInput
	// The number of bits of generated keys
	KeyBits pulumi.IntPtrInput
	// The type of generated keys
	KeyType pulumi.StringPtrInput
	// Specify the allowed key usage constraint on issued certificates
	KeyUsages pulumi.StringArrayInput
	// The locality of generated certificates
	Localities pulumi.StringArrayInput
	// The maximum TTL
	MaxTtl pulumi.StringPtrInput
	// The name to identify this role within the backend. Must be unique within the backend.
	Name pulumi.StringPtrInput
	// Flag to not store certificates in the storage backend
	NoStore pulumi.BoolPtrInput
	// Specifies the duration by which to backdate the NotBefore property.
	NotBeforeDuration pulumi.StringPtrInput
	// The organization unit of generated certificates
	OrganizationUnit pulumi.StringArrayInput
	// The organization of generated certificates
	Organizations pulumi.StringArrayInput
	// Specify the list of allowed policies IODs
	PolicyIdentifiers pulumi.StringArrayInput
	// The postal code of generated certificates
	PostalCodes pulumi.StringArrayInput
	// The province of generated certificates
	Provinces pulumi.StringArrayInput
	// Flag to force CN usage
	RequireCn pulumi.BoolPtrInput
	// Flag to specify certificates for server use
	ServerFlag pulumi.BoolPtrInput
	// The street address of generated certificates
	StreetAddresses pulumi.StringArrayInput
	// The TTL
	Ttl pulumi.StringPtrInput
	// Flag to use the CN in the CSR
	UseCsrCommonName pulumi.BoolPtrInput
	// Flag to use the SANs in the CSR
	UseCsrSans pulumi.BoolPtrInput
}

func (SecretBackendRoleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*secretBackendRoleArgs)(nil)).Elem()
}