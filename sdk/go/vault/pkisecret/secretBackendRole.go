// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package pkisecret

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a role on an PKI Secret Backend for Vault.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-vault/sdk/v5/go/vault"
// 	"github.com/pulumi/pulumi-vault/sdk/v5/go/vault/pkiSecret"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		pki, err := vault.NewMount(ctx, "pki", &vault.MountArgs{
// 			Path:                   pulumi.String("pki"),
// 			Type:                   pulumi.String("pki"),
// 			DefaultLeaseTtlSeconds: pulumi.Int(3600),
// 			MaxLeaseTtlSeconds:     pulumi.Int(86400),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = pkiSecret.NewSecretBackendRole(ctx, "role", &pkiSecret.SecretBackendRoleArgs{
// 			Backend:     pki.Path,
// 			Ttl:         pulumi.String("3600"),
// 			AllowIpSans: pulumi.Bool(true),
// 			KeyType:     pulumi.String("rsa"),
// 			KeyBits:     pulumi.Int(4096),
// 			AllowedDomains: pulumi.StringArray{
// 				pulumi.String("example.com"),
// 				pulumi.String("my.domain"),
// 			},
// 			AllowSubdomains: pulumi.Bool(true),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## Import
//
// PKI secret backend roles can be imported using the `path`, e.g.
//
// ```sh
//  $ pulumi import vault:pkiSecret/secretBackendRole:SecretBackendRole role pki/roles/my_role
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
	// Flag, if set, `allowedDomains` can be specified using identity template expressions such as `{{identity.entity.aliases.<mount accessor>.name}}`.
	AllowedDomainsTemplate pulumi.BoolPtrOutput `pulumi:"allowedDomainsTemplate"`
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
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Backend == nil {
		return nil, errors.New("invalid value for required argument 'Backend'")
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
	// Flag, if set, `allowedDomains` can be specified using identity template expressions such as `{{identity.entity.aliases.<mount accessor>.name}}`.
	AllowedDomainsTemplate *bool `pulumi:"allowedDomainsTemplate"`
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
	// Flag, if set, `allowedDomains` can be specified using identity template expressions such as `{{identity.entity.aliases.<mount accessor>.name}}`.
	AllowedDomainsTemplate pulumi.BoolPtrInput
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
	// Flag, if set, `allowedDomains` can be specified using identity template expressions such as `{{identity.entity.aliases.<mount accessor>.name}}`.
	AllowedDomainsTemplate *bool `pulumi:"allowedDomainsTemplate"`
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
	// Flag, if set, `allowedDomains` can be specified using identity template expressions such as `{{identity.entity.aliases.<mount accessor>.name}}`.
	AllowedDomainsTemplate pulumi.BoolPtrInput
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

type SecretBackendRoleInput interface {
	pulumi.Input

	ToSecretBackendRoleOutput() SecretBackendRoleOutput
	ToSecretBackendRoleOutputWithContext(ctx context.Context) SecretBackendRoleOutput
}

func (*SecretBackendRole) ElementType() reflect.Type {
	return reflect.TypeOf((**SecretBackendRole)(nil)).Elem()
}

func (i *SecretBackendRole) ToSecretBackendRoleOutput() SecretBackendRoleOutput {
	return i.ToSecretBackendRoleOutputWithContext(context.Background())
}

func (i *SecretBackendRole) ToSecretBackendRoleOutputWithContext(ctx context.Context) SecretBackendRoleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecretBackendRoleOutput)
}

// SecretBackendRoleArrayInput is an input type that accepts SecretBackendRoleArray and SecretBackendRoleArrayOutput values.
// You can construct a concrete instance of `SecretBackendRoleArrayInput` via:
//
//          SecretBackendRoleArray{ SecretBackendRoleArgs{...} }
type SecretBackendRoleArrayInput interface {
	pulumi.Input

	ToSecretBackendRoleArrayOutput() SecretBackendRoleArrayOutput
	ToSecretBackendRoleArrayOutputWithContext(context.Context) SecretBackendRoleArrayOutput
}

type SecretBackendRoleArray []SecretBackendRoleInput

func (SecretBackendRoleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SecretBackendRole)(nil)).Elem()
}

func (i SecretBackendRoleArray) ToSecretBackendRoleArrayOutput() SecretBackendRoleArrayOutput {
	return i.ToSecretBackendRoleArrayOutputWithContext(context.Background())
}

func (i SecretBackendRoleArray) ToSecretBackendRoleArrayOutputWithContext(ctx context.Context) SecretBackendRoleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecretBackendRoleArrayOutput)
}

// SecretBackendRoleMapInput is an input type that accepts SecretBackendRoleMap and SecretBackendRoleMapOutput values.
// You can construct a concrete instance of `SecretBackendRoleMapInput` via:
//
//          SecretBackendRoleMap{ "key": SecretBackendRoleArgs{...} }
type SecretBackendRoleMapInput interface {
	pulumi.Input

	ToSecretBackendRoleMapOutput() SecretBackendRoleMapOutput
	ToSecretBackendRoleMapOutputWithContext(context.Context) SecretBackendRoleMapOutput
}

type SecretBackendRoleMap map[string]SecretBackendRoleInput

func (SecretBackendRoleMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SecretBackendRole)(nil)).Elem()
}

func (i SecretBackendRoleMap) ToSecretBackendRoleMapOutput() SecretBackendRoleMapOutput {
	return i.ToSecretBackendRoleMapOutputWithContext(context.Background())
}

func (i SecretBackendRoleMap) ToSecretBackendRoleMapOutputWithContext(ctx context.Context) SecretBackendRoleMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecretBackendRoleMapOutput)
}

type SecretBackendRoleOutput struct{ *pulumi.OutputState }

func (SecretBackendRoleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SecretBackendRole)(nil)).Elem()
}

func (o SecretBackendRoleOutput) ToSecretBackendRoleOutput() SecretBackendRoleOutput {
	return o
}

func (o SecretBackendRoleOutput) ToSecretBackendRoleOutputWithContext(ctx context.Context) SecretBackendRoleOutput {
	return o
}

type SecretBackendRoleArrayOutput struct{ *pulumi.OutputState }

func (SecretBackendRoleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SecretBackendRole)(nil)).Elem()
}

func (o SecretBackendRoleArrayOutput) ToSecretBackendRoleArrayOutput() SecretBackendRoleArrayOutput {
	return o
}

func (o SecretBackendRoleArrayOutput) ToSecretBackendRoleArrayOutputWithContext(ctx context.Context) SecretBackendRoleArrayOutput {
	return o
}

func (o SecretBackendRoleArrayOutput) Index(i pulumi.IntInput) SecretBackendRoleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SecretBackendRole {
		return vs[0].([]*SecretBackendRole)[vs[1].(int)]
	}).(SecretBackendRoleOutput)
}

type SecretBackendRoleMapOutput struct{ *pulumi.OutputState }

func (SecretBackendRoleMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SecretBackendRole)(nil)).Elem()
}

func (o SecretBackendRoleMapOutput) ToSecretBackendRoleMapOutput() SecretBackendRoleMapOutput {
	return o
}

func (o SecretBackendRoleMapOutput) ToSecretBackendRoleMapOutputWithContext(ctx context.Context) SecretBackendRoleMapOutput {
	return o
}

func (o SecretBackendRoleMapOutput) MapIndex(k pulumi.StringInput) SecretBackendRoleOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SecretBackendRole {
		return vs[0].(map[string]*SecretBackendRole)[vs[1].(string)]
	}).(SecretBackendRoleOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SecretBackendRoleInput)(nil)).Elem(), &SecretBackendRole{})
	pulumi.RegisterInputType(reflect.TypeOf((*SecretBackendRoleArrayInput)(nil)).Elem(), SecretBackendRoleArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SecretBackendRoleMapInput)(nil)).Elem(), SecretBackendRoleMap{})
	pulumi.RegisterOutputType(SecretBackendRoleOutput{})
	pulumi.RegisterOutputType(SecretBackendRoleArrayOutput{})
	pulumi.RegisterOutputType(SecretBackendRoleMapOutput{})
}
