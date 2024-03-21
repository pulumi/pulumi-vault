// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ldap

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-vault/sdk/v6/go/vault/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## Example Usage
//
// <!--Start PulumiCodeChooser -->
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-vault/sdk/v6/go/vault/ldap"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := ldap.NewSecretBackend(ctx, "config", &ldap.SecretBackendArgs{
//				Binddn:      pulumi.String("CN=Administrator,CN=Users,DC=corp,DC=example,DC=net"),
//				Bindpass:    pulumi.String("SuperSecretPassw0rd"),
//				InsecureTls: pulumi.Bool(true),
//				Path:        pulumi.String("my-custom-ldap"),
//				Url:         pulumi.String("ldaps://localhost"),
//				Userdn:      pulumi.String("CN=Users,DC=corp,DC=example,DC=net"),
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
// LDAP secret backend can be imported using the `${mount}/config`, e.g.
//
// ```sh
// $ pulumi import vault:ldap/secretBackend:SecretBackend config ldap/config
// ```
type SecretBackend struct {
	pulumi.CustomResourceState

	// Accessor of the mount
	Accessor pulumi.StringOutput `pulumi:"accessor"`
	// List of managed key registry entry names that the mount in question is allowed to access
	AllowedManagedKeys pulumi.StringArrayOutput `pulumi:"allowedManagedKeys"`
	// Specifies the list of keys that will not be HMAC'd by audit devices in the request data object.
	AuditNonHmacRequestKeys pulumi.StringArrayOutput `pulumi:"auditNonHmacRequestKeys"`
	// Specifies the list of keys that will not be HMAC'd by audit devices in the response data object.
	AuditNonHmacResponseKeys pulumi.StringArrayOutput `pulumi:"auditNonHmacResponseKeys"`
	// Distinguished name of object to bind when performing user and group search.
	Binddn pulumi.StringOutput `pulumi:"binddn"`
	// Password to use along with binddn when performing user search.
	Bindpass pulumi.StringOutput `pulumi:"bindpass"`
	// CA certificate to use when verifying LDAP server certificate, must be
	// x509 PEM encoded.
	Certificate pulumi.StringPtrOutput `pulumi:"certificate"`
	// Client certificate to provide to the LDAP server, must be x509 PEM encoded.
	ClientTlsCert pulumi.StringPtrOutput `pulumi:"clientTlsCert"`
	// Client certificate key to provide to the LDAP server, must be x509 PEM encoded.
	ClientTlsKey pulumi.StringPtrOutput `pulumi:"clientTlsKey"`
	// Timeout, in seconds, when attempting to connect to the LDAP server before trying
	// the next URL in the configuration.
	ConnectionTimeout pulumi.IntPtrOutput `pulumi:"connectionTimeout"`
	// Default lease duration for secrets in seconds.
	DefaultLeaseTtlSeconds pulumi.IntOutput `pulumi:"defaultLeaseTtlSeconds"`
	// Human-friendly description of the mount for the Active Directory backend.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// If set, opts out of mount migration on path updates.
	DisableRemount pulumi.BoolPtrOutput `pulumi:"disableRemount"`
	// Enable the secrets engine to access Vault's external entropy source
	ExternalEntropyAccess pulumi.BoolPtrOutput `pulumi:"externalEntropyAccess"`
	// Skip LDAP server SSL Certificate verification. This is not recommended for production.
	// Defaults to `false`.
	InsecureTls pulumi.BoolPtrOutput `pulumi:"insecureTls"`
	// **Deprecated** use `passwordPolicy`. The desired length of passwords that Vault generates.
	// *Mutually exclusive with `passwordPolicy` on vault-1.11+*
	//
	// Deprecated: Length is deprecated and passwordPolicy should be used with Vault >= 1.5.
	Length pulumi.IntOutput `pulumi:"length"`
	// Mark the secrets engine as local-only. Local engines are not replicated or removed by
	// replication.Tolerance duration to use when checking the last rotation time.
	Local pulumi.BoolPtrOutput `pulumi:"local"`
	// Maximum possible lease duration for secrets in seconds.
	MaxLeaseTtlSeconds pulumi.IntOutput `pulumi:"maxLeaseTtlSeconds"`
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
	// *Available only for Vault Enterprise*.
	Namespace pulumi.StringPtrOutput `pulumi:"namespace"`
	// Specifies mount type specific options that are passed to the backend
	Options pulumi.MapOutput `pulumi:"options"`
	// Name of the password policy to use to generate passwords.
	PasswordPolicy pulumi.StringPtrOutput `pulumi:"passwordPolicy"`
	// The unique path this backend should be mounted at. Must
	// not begin or end with a `/`. Defaults to `ldap`.
	Path pulumi.StringPtrOutput `pulumi:"path"`
	// Timeout, in seconds, for the connection when making requests against the server
	// before returning back an error.
	RequestTimeout pulumi.IntOutput `pulumi:"requestTimeout"`
	// The LDAP schema to use when storing entry passwords. Valid schemas include `openldap`, `ad`, and `racf`. Default is `openldap`.
	Schema pulumi.StringOutput `pulumi:"schema"`
	// Enable seal wrapping for the mount, causing values stored by the mount to be wrapped by the seal's encryption capability
	SealWrap pulumi.BoolOutput `pulumi:"sealWrap"`
	// Issue a StartTLS command after establishing unencrypted connection.
	Starttls pulumi.BoolOutput `pulumi:"starttls"`
	// Enables userPrincipalDomain login with [username]@UPNDomain.
	Upndomain pulumi.StringOutput `pulumi:"upndomain"`
	// LDAP URL to connect to. Multiple URLs can be specified by concatenating
	// them with commas; they will be tried in-order. Defaults to `ldap://127.0.0.1`.
	Url pulumi.StringOutput `pulumi:"url"`
	// Attribute used when searching users. Defaults to `cn`.
	Userattr pulumi.StringOutput `pulumi:"userattr"`
	// LDAP domain to use for users (eg: ou=People,dc=example,dc=org)`.
	Userdn pulumi.StringPtrOutput `pulumi:"userdn"`
}

// NewSecretBackend registers a new resource with the given unique name, arguments, and options.
func NewSecretBackend(ctx *pulumi.Context,
	name string, args *SecretBackendArgs, opts ...pulumi.ResourceOption) (*SecretBackend, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Binddn == nil {
		return nil, errors.New("invalid value for required argument 'Binddn'")
	}
	if args.Bindpass == nil {
		return nil, errors.New("invalid value for required argument 'Bindpass'")
	}
	if args.Bindpass != nil {
		args.Bindpass = pulumi.ToSecret(args.Bindpass).(pulumi.StringInput)
	}
	if args.ClientTlsCert != nil {
		args.ClientTlsCert = pulumi.ToSecret(args.ClientTlsCert).(pulumi.StringPtrInput)
	}
	if args.ClientTlsKey != nil {
		args.ClientTlsKey = pulumi.ToSecret(args.ClientTlsKey).(pulumi.StringPtrInput)
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"bindpass",
		"clientTlsCert",
		"clientTlsKey",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource SecretBackend
	err := ctx.RegisterResource("vault:ldap/secretBackend:SecretBackend", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSecretBackend gets an existing SecretBackend resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSecretBackend(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SecretBackendState, opts ...pulumi.ResourceOption) (*SecretBackend, error) {
	var resource SecretBackend
	err := ctx.ReadResource("vault:ldap/secretBackend:SecretBackend", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SecretBackend resources.
type secretBackendState struct {
	// Accessor of the mount
	Accessor *string `pulumi:"accessor"`
	// List of managed key registry entry names that the mount in question is allowed to access
	AllowedManagedKeys []string `pulumi:"allowedManagedKeys"`
	// Specifies the list of keys that will not be HMAC'd by audit devices in the request data object.
	AuditNonHmacRequestKeys []string `pulumi:"auditNonHmacRequestKeys"`
	// Specifies the list of keys that will not be HMAC'd by audit devices in the response data object.
	AuditNonHmacResponseKeys []string `pulumi:"auditNonHmacResponseKeys"`
	// Distinguished name of object to bind when performing user and group search.
	Binddn *string `pulumi:"binddn"`
	// Password to use along with binddn when performing user search.
	Bindpass *string `pulumi:"bindpass"`
	// CA certificate to use when verifying LDAP server certificate, must be
	// x509 PEM encoded.
	Certificate *string `pulumi:"certificate"`
	// Client certificate to provide to the LDAP server, must be x509 PEM encoded.
	ClientTlsCert *string `pulumi:"clientTlsCert"`
	// Client certificate key to provide to the LDAP server, must be x509 PEM encoded.
	ClientTlsKey *string `pulumi:"clientTlsKey"`
	// Timeout, in seconds, when attempting to connect to the LDAP server before trying
	// the next URL in the configuration.
	ConnectionTimeout *int `pulumi:"connectionTimeout"`
	// Default lease duration for secrets in seconds.
	DefaultLeaseTtlSeconds *int `pulumi:"defaultLeaseTtlSeconds"`
	// Human-friendly description of the mount for the Active Directory backend.
	Description *string `pulumi:"description"`
	// If set, opts out of mount migration on path updates.
	DisableRemount *bool `pulumi:"disableRemount"`
	// Enable the secrets engine to access Vault's external entropy source
	ExternalEntropyAccess *bool `pulumi:"externalEntropyAccess"`
	// Skip LDAP server SSL Certificate verification. This is not recommended for production.
	// Defaults to `false`.
	InsecureTls *bool `pulumi:"insecureTls"`
	// **Deprecated** use `passwordPolicy`. The desired length of passwords that Vault generates.
	// *Mutually exclusive with `passwordPolicy` on vault-1.11+*
	//
	// Deprecated: Length is deprecated and passwordPolicy should be used with Vault >= 1.5.
	Length *int `pulumi:"length"`
	// Mark the secrets engine as local-only. Local engines are not replicated or removed by
	// replication.Tolerance duration to use when checking the last rotation time.
	Local *bool `pulumi:"local"`
	// Maximum possible lease duration for secrets in seconds.
	MaxLeaseTtlSeconds *int `pulumi:"maxLeaseTtlSeconds"`
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
	// *Available only for Vault Enterprise*.
	Namespace *string `pulumi:"namespace"`
	// Specifies mount type specific options that are passed to the backend
	Options map[string]interface{} `pulumi:"options"`
	// Name of the password policy to use to generate passwords.
	PasswordPolicy *string `pulumi:"passwordPolicy"`
	// The unique path this backend should be mounted at. Must
	// not begin or end with a `/`. Defaults to `ldap`.
	Path *string `pulumi:"path"`
	// Timeout, in seconds, for the connection when making requests against the server
	// before returning back an error.
	RequestTimeout *int `pulumi:"requestTimeout"`
	// The LDAP schema to use when storing entry passwords. Valid schemas include `openldap`, `ad`, and `racf`. Default is `openldap`.
	Schema *string `pulumi:"schema"`
	// Enable seal wrapping for the mount, causing values stored by the mount to be wrapped by the seal's encryption capability
	SealWrap *bool `pulumi:"sealWrap"`
	// Issue a StartTLS command after establishing unencrypted connection.
	Starttls *bool `pulumi:"starttls"`
	// Enables userPrincipalDomain login with [username]@UPNDomain.
	Upndomain *string `pulumi:"upndomain"`
	// LDAP URL to connect to. Multiple URLs can be specified by concatenating
	// them with commas; they will be tried in-order. Defaults to `ldap://127.0.0.1`.
	Url *string `pulumi:"url"`
	// Attribute used when searching users. Defaults to `cn`.
	Userattr *string `pulumi:"userattr"`
	// LDAP domain to use for users (eg: ou=People,dc=example,dc=org)`.
	Userdn *string `pulumi:"userdn"`
}

type SecretBackendState struct {
	// Accessor of the mount
	Accessor pulumi.StringPtrInput
	// List of managed key registry entry names that the mount in question is allowed to access
	AllowedManagedKeys pulumi.StringArrayInput
	// Specifies the list of keys that will not be HMAC'd by audit devices in the request data object.
	AuditNonHmacRequestKeys pulumi.StringArrayInput
	// Specifies the list of keys that will not be HMAC'd by audit devices in the response data object.
	AuditNonHmacResponseKeys pulumi.StringArrayInput
	// Distinguished name of object to bind when performing user and group search.
	Binddn pulumi.StringPtrInput
	// Password to use along with binddn when performing user search.
	Bindpass pulumi.StringPtrInput
	// CA certificate to use when verifying LDAP server certificate, must be
	// x509 PEM encoded.
	Certificate pulumi.StringPtrInput
	// Client certificate to provide to the LDAP server, must be x509 PEM encoded.
	ClientTlsCert pulumi.StringPtrInput
	// Client certificate key to provide to the LDAP server, must be x509 PEM encoded.
	ClientTlsKey pulumi.StringPtrInput
	// Timeout, in seconds, when attempting to connect to the LDAP server before trying
	// the next URL in the configuration.
	ConnectionTimeout pulumi.IntPtrInput
	// Default lease duration for secrets in seconds.
	DefaultLeaseTtlSeconds pulumi.IntPtrInput
	// Human-friendly description of the mount for the Active Directory backend.
	Description pulumi.StringPtrInput
	// If set, opts out of mount migration on path updates.
	DisableRemount pulumi.BoolPtrInput
	// Enable the secrets engine to access Vault's external entropy source
	ExternalEntropyAccess pulumi.BoolPtrInput
	// Skip LDAP server SSL Certificate verification. This is not recommended for production.
	// Defaults to `false`.
	InsecureTls pulumi.BoolPtrInput
	// **Deprecated** use `passwordPolicy`. The desired length of passwords that Vault generates.
	// *Mutually exclusive with `passwordPolicy` on vault-1.11+*
	//
	// Deprecated: Length is deprecated and passwordPolicy should be used with Vault >= 1.5.
	Length pulumi.IntPtrInput
	// Mark the secrets engine as local-only. Local engines are not replicated or removed by
	// replication.Tolerance duration to use when checking the last rotation time.
	Local pulumi.BoolPtrInput
	// Maximum possible lease duration for secrets in seconds.
	MaxLeaseTtlSeconds pulumi.IntPtrInput
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
	// *Available only for Vault Enterprise*.
	Namespace pulumi.StringPtrInput
	// Specifies mount type specific options that are passed to the backend
	Options pulumi.MapInput
	// Name of the password policy to use to generate passwords.
	PasswordPolicy pulumi.StringPtrInput
	// The unique path this backend should be mounted at. Must
	// not begin or end with a `/`. Defaults to `ldap`.
	Path pulumi.StringPtrInput
	// Timeout, in seconds, for the connection when making requests against the server
	// before returning back an error.
	RequestTimeout pulumi.IntPtrInput
	// The LDAP schema to use when storing entry passwords. Valid schemas include `openldap`, `ad`, and `racf`. Default is `openldap`.
	Schema pulumi.StringPtrInput
	// Enable seal wrapping for the mount, causing values stored by the mount to be wrapped by the seal's encryption capability
	SealWrap pulumi.BoolPtrInput
	// Issue a StartTLS command after establishing unencrypted connection.
	Starttls pulumi.BoolPtrInput
	// Enables userPrincipalDomain login with [username]@UPNDomain.
	Upndomain pulumi.StringPtrInput
	// LDAP URL to connect to. Multiple URLs can be specified by concatenating
	// them with commas; they will be tried in-order. Defaults to `ldap://127.0.0.1`.
	Url pulumi.StringPtrInput
	// Attribute used when searching users. Defaults to `cn`.
	Userattr pulumi.StringPtrInput
	// LDAP domain to use for users (eg: ou=People,dc=example,dc=org)`.
	Userdn pulumi.StringPtrInput
}

func (SecretBackendState) ElementType() reflect.Type {
	return reflect.TypeOf((*secretBackendState)(nil)).Elem()
}

type secretBackendArgs struct {
	// List of managed key registry entry names that the mount in question is allowed to access
	AllowedManagedKeys []string `pulumi:"allowedManagedKeys"`
	// Specifies the list of keys that will not be HMAC'd by audit devices in the request data object.
	AuditNonHmacRequestKeys []string `pulumi:"auditNonHmacRequestKeys"`
	// Specifies the list of keys that will not be HMAC'd by audit devices in the response data object.
	AuditNonHmacResponseKeys []string `pulumi:"auditNonHmacResponseKeys"`
	// Distinguished name of object to bind when performing user and group search.
	Binddn string `pulumi:"binddn"`
	// Password to use along with binddn when performing user search.
	Bindpass string `pulumi:"bindpass"`
	// CA certificate to use when verifying LDAP server certificate, must be
	// x509 PEM encoded.
	Certificate *string `pulumi:"certificate"`
	// Client certificate to provide to the LDAP server, must be x509 PEM encoded.
	ClientTlsCert *string `pulumi:"clientTlsCert"`
	// Client certificate key to provide to the LDAP server, must be x509 PEM encoded.
	ClientTlsKey *string `pulumi:"clientTlsKey"`
	// Timeout, in seconds, when attempting to connect to the LDAP server before trying
	// the next URL in the configuration.
	ConnectionTimeout *int `pulumi:"connectionTimeout"`
	// Default lease duration for secrets in seconds.
	DefaultLeaseTtlSeconds *int `pulumi:"defaultLeaseTtlSeconds"`
	// Human-friendly description of the mount for the Active Directory backend.
	Description *string `pulumi:"description"`
	// If set, opts out of mount migration on path updates.
	DisableRemount *bool `pulumi:"disableRemount"`
	// Enable the secrets engine to access Vault's external entropy source
	ExternalEntropyAccess *bool `pulumi:"externalEntropyAccess"`
	// Skip LDAP server SSL Certificate verification. This is not recommended for production.
	// Defaults to `false`.
	InsecureTls *bool `pulumi:"insecureTls"`
	// **Deprecated** use `passwordPolicy`. The desired length of passwords that Vault generates.
	// *Mutually exclusive with `passwordPolicy` on vault-1.11+*
	//
	// Deprecated: Length is deprecated and passwordPolicy should be used with Vault >= 1.5.
	Length *int `pulumi:"length"`
	// Mark the secrets engine as local-only. Local engines are not replicated or removed by
	// replication.Tolerance duration to use when checking the last rotation time.
	Local *bool `pulumi:"local"`
	// Maximum possible lease duration for secrets in seconds.
	MaxLeaseTtlSeconds *int `pulumi:"maxLeaseTtlSeconds"`
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
	// *Available only for Vault Enterprise*.
	Namespace *string `pulumi:"namespace"`
	// Specifies mount type specific options that are passed to the backend
	Options map[string]interface{} `pulumi:"options"`
	// Name of the password policy to use to generate passwords.
	PasswordPolicy *string `pulumi:"passwordPolicy"`
	// The unique path this backend should be mounted at. Must
	// not begin or end with a `/`. Defaults to `ldap`.
	Path *string `pulumi:"path"`
	// Timeout, in seconds, for the connection when making requests against the server
	// before returning back an error.
	RequestTimeout *int `pulumi:"requestTimeout"`
	// The LDAP schema to use when storing entry passwords. Valid schemas include `openldap`, `ad`, and `racf`. Default is `openldap`.
	Schema *string `pulumi:"schema"`
	// Enable seal wrapping for the mount, causing values stored by the mount to be wrapped by the seal's encryption capability
	SealWrap *bool `pulumi:"sealWrap"`
	// Issue a StartTLS command after establishing unencrypted connection.
	Starttls *bool `pulumi:"starttls"`
	// Enables userPrincipalDomain login with [username]@UPNDomain.
	Upndomain *string `pulumi:"upndomain"`
	// LDAP URL to connect to. Multiple URLs can be specified by concatenating
	// them with commas; they will be tried in-order. Defaults to `ldap://127.0.0.1`.
	Url *string `pulumi:"url"`
	// Attribute used when searching users. Defaults to `cn`.
	Userattr *string `pulumi:"userattr"`
	// LDAP domain to use for users (eg: ou=People,dc=example,dc=org)`.
	Userdn *string `pulumi:"userdn"`
}

// The set of arguments for constructing a SecretBackend resource.
type SecretBackendArgs struct {
	// List of managed key registry entry names that the mount in question is allowed to access
	AllowedManagedKeys pulumi.StringArrayInput
	// Specifies the list of keys that will not be HMAC'd by audit devices in the request data object.
	AuditNonHmacRequestKeys pulumi.StringArrayInput
	// Specifies the list of keys that will not be HMAC'd by audit devices in the response data object.
	AuditNonHmacResponseKeys pulumi.StringArrayInput
	// Distinguished name of object to bind when performing user and group search.
	Binddn pulumi.StringInput
	// Password to use along with binddn when performing user search.
	Bindpass pulumi.StringInput
	// CA certificate to use when verifying LDAP server certificate, must be
	// x509 PEM encoded.
	Certificate pulumi.StringPtrInput
	// Client certificate to provide to the LDAP server, must be x509 PEM encoded.
	ClientTlsCert pulumi.StringPtrInput
	// Client certificate key to provide to the LDAP server, must be x509 PEM encoded.
	ClientTlsKey pulumi.StringPtrInput
	// Timeout, in seconds, when attempting to connect to the LDAP server before trying
	// the next URL in the configuration.
	ConnectionTimeout pulumi.IntPtrInput
	// Default lease duration for secrets in seconds.
	DefaultLeaseTtlSeconds pulumi.IntPtrInput
	// Human-friendly description of the mount for the Active Directory backend.
	Description pulumi.StringPtrInput
	// If set, opts out of mount migration on path updates.
	DisableRemount pulumi.BoolPtrInput
	// Enable the secrets engine to access Vault's external entropy source
	ExternalEntropyAccess pulumi.BoolPtrInput
	// Skip LDAP server SSL Certificate verification. This is not recommended for production.
	// Defaults to `false`.
	InsecureTls pulumi.BoolPtrInput
	// **Deprecated** use `passwordPolicy`. The desired length of passwords that Vault generates.
	// *Mutually exclusive with `passwordPolicy` on vault-1.11+*
	//
	// Deprecated: Length is deprecated and passwordPolicy should be used with Vault >= 1.5.
	Length pulumi.IntPtrInput
	// Mark the secrets engine as local-only. Local engines are not replicated or removed by
	// replication.Tolerance duration to use when checking the last rotation time.
	Local pulumi.BoolPtrInput
	// Maximum possible lease duration for secrets in seconds.
	MaxLeaseTtlSeconds pulumi.IntPtrInput
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
	// *Available only for Vault Enterprise*.
	Namespace pulumi.StringPtrInput
	// Specifies mount type specific options that are passed to the backend
	Options pulumi.MapInput
	// Name of the password policy to use to generate passwords.
	PasswordPolicy pulumi.StringPtrInput
	// The unique path this backend should be mounted at. Must
	// not begin or end with a `/`. Defaults to `ldap`.
	Path pulumi.StringPtrInput
	// Timeout, in seconds, for the connection when making requests against the server
	// before returning back an error.
	RequestTimeout pulumi.IntPtrInput
	// The LDAP schema to use when storing entry passwords. Valid schemas include `openldap`, `ad`, and `racf`. Default is `openldap`.
	Schema pulumi.StringPtrInput
	// Enable seal wrapping for the mount, causing values stored by the mount to be wrapped by the seal's encryption capability
	SealWrap pulumi.BoolPtrInput
	// Issue a StartTLS command after establishing unencrypted connection.
	Starttls pulumi.BoolPtrInput
	// Enables userPrincipalDomain login with [username]@UPNDomain.
	Upndomain pulumi.StringPtrInput
	// LDAP URL to connect to. Multiple URLs can be specified by concatenating
	// them with commas; they will be tried in-order. Defaults to `ldap://127.0.0.1`.
	Url pulumi.StringPtrInput
	// Attribute used when searching users. Defaults to `cn`.
	Userattr pulumi.StringPtrInput
	// LDAP domain to use for users (eg: ou=People,dc=example,dc=org)`.
	Userdn pulumi.StringPtrInput
}

func (SecretBackendArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*secretBackendArgs)(nil)).Elem()
}

type SecretBackendInput interface {
	pulumi.Input

	ToSecretBackendOutput() SecretBackendOutput
	ToSecretBackendOutputWithContext(ctx context.Context) SecretBackendOutput
}

func (*SecretBackend) ElementType() reflect.Type {
	return reflect.TypeOf((**SecretBackend)(nil)).Elem()
}

func (i *SecretBackend) ToSecretBackendOutput() SecretBackendOutput {
	return i.ToSecretBackendOutputWithContext(context.Background())
}

func (i *SecretBackend) ToSecretBackendOutputWithContext(ctx context.Context) SecretBackendOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecretBackendOutput)
}

// SecretBackendArrayInput is an input type that accepts SecretBackendArray and SecretBackendArrayOutput values.
// You can construct a concrete instance of `SecretBackendArrayInput` via:
//
//	SecretBackendArray{ SecretBackendArgs{...} }
type SecretBackendArrayInput interface {
	pulumi.Input

	ToSecretBackendArrayOutput() SecretBackendArrayOutput
	ToSecretBackendArrayOutputWithContext(context.Context) SecretBackendArrayOutput
}

type SecretBackendArray []SecretBackendInput

func (SecretBackendArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SecretBackend)(nil)).Elem()
}

func (i SecretBackendArray) ToSecretBackendArrayOutput() SecretBackendArrayOutput {
	return i.ToSecretBackendArrayOutputWithContext(context.Background())
}

func (i SecretBackendArray) ToSecretBackendArrayOutputWithContext(ctx context.Context) SecretBackendArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecretBackendArrayOutput)
}

// SecretBackendMapInput is an input type that accepts SecretBackendMap and SecretBackendMapOutput values.
// You can construct a concrete instance of `SecretBackendMapInput` via:
//
//	SecretBackendMap{ "key": SecretBackendArgs{...} }
type SecretBackendMapInput interface {
	pulumi.Input

	ToSecretBackendMapOutput() SecretBackendMapOutput
	ToSecretBackendMapOutputWithContext(context.Context) SecretBackendMapOutput
}

type SecretBackendMap map[string]SecretBackendInput

func (SecretBackendMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SecretBackend)(nil)).Elem()
}

func (i SecretBackendMap) ToSecretBackendMapOutput() SecretBackendMapOutput {
	return i.ToSecretBackendMapOutputWithContext(context.Background())
}

func (i SecretBackendMap) ToSecretBackendMapOutputWithContext(ctx context.Context) SecretBackendMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecretBackendMapOutput)
}

type SecretBackendOutput struct{ *pulumi.OutputState }

func (SecretBackendOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SecretBackend)(nil)).Elem()
}

func (o SecretBackendOutput) ToSecretBackendOutput() SecretBackendOutput {
	return o
}

func (o SecretBackendOutput) ToSecretBackendOutputWithContext(ctx context.Context) SecretBackendOutput {
	return o
}

// Accessor of the mount
func (o SecretBackendOutput) Accessor() pulumi.StringOutput {
	return o.ApplyT(func(v *SecretBackend) pulumi.StringOutput { return v.Accessor }).(pulumi.StringOutput)
}

// List of managed key registry entry names that the mount in question is allowed to access
func (o SecretBackendOutput) AllowedManagedKeys() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *SecretBackend) pulumi.StringArrayOutput { return v.AllowedManagedKeys }).(pulumi.StringArrayOutput)
}

// Specifies the list of keys that will not be HMAC'd by audit devices in the request data object.
func (o SecretBackendOutput) AuditNonHmacRequestKeys() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *SecretBackend) pulumi.StringArrayOutput { return v.AuditNonHmacRequestKeys }).(pulumi.StringArrayOutput)
}

// Specifies the list of keys that will not be HMAC'd by audit devices in the response data object.
func (o SecretBackendOutput) AuditNonHmacResponseKeys() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *SecretBackend) pulumi.StringArrayOutput { return v.AuditNonHmacResponseKeys }).(pulumi.StringArrayOutput)
}

// Distinguished name of object to bind when performing user and group search.
func (o SecretBackendOutput) Binddn() pulumi.StringOutput {
	return o.ApplyT(func(v *SecretBackend) pulumi.StringOutput { return v.Binddn }).(pulumi.StringOutput)
}

// Password to use along with binddn when performing user search.
func (o SecretBackendOutput) Bindpass() pulumi.StringOutput {
	return o.ApplyT(func(v *SecretBackend) pulumi.StringOutput { return v.Bindpass }).(pulumi.StringOutput)
}

// CA certificate to use when verifying LDAP server certificate, must be
// x509 PEM encoded.
func (o SecretBackendOutput) Certificate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecretBackend) pulumi.StringPtrOutput { return v.Certificate }).(pulumi.StringPtrOutput)
}

// Client certificate to provide to the LDAP server, must be x509 PEM encoded.
func (o SecretBackendOutput) ClientTlsCert() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecretBackend) pulumi.StringPtrOutput { return v.ClientTlsCert }).(pulumi.StringPtrOutput)
}

// Client certificate key to provide to the LDAP server, must be x509 PEM encoded.
func (o SecretBackendOutput) ClientTlsKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecretBackend) pulumi.StringPtrOutput { return v.ClientTlsKey }).(pulumi.StringPtrOutput)
}

// Timeout, in seconds, when attempting to connect to the LDAP server before trying
// the next URL in the configuration.
func (o SecretBackendOutput) ConnectionTimeout() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SecretBackend) pulumi.IntPtrOutput { return v.ConnectionTimeout }).(pulumi.IntPtrOutput)
}

// Default lease duration for secrets in seconds.
func (o SecretBackendOutput) DefaultLeaseTtlSeconds() pulumi.IntOutput {
	return o.ApplyT(func(v *SecretBackend) pulumi.IntOutput { return v.DefaultLeaseTtlSeconds }).(pulumi.IntOutput)
}

// Human-friendly description of the mount for the Active Directory backend.
func (o SecretBackendOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecretBackend) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// If set, opts out of mount migration on path updates.
func (o SecretBackendOutput) DisableRemount() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SecretBackend) pulumi.BoolPtrOutput { return v.DisableRemount }).(pulumi.BoolPtrOutput)
}

// Enable the secrets engine to access Vault's external entropy source
func (o SecretBackendOutput) ExternalEntropyAccess() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SecretBackend) pulumi.BoolPtrOutput { return v.ExternalEntropyAccess }).(pulumi.BoolPtrOutput)
}

// Skip LDAP server SSL Certificate verification. This is not recommended for production.
// Defaults to `false`.
func (o SecretBackendOutput) InsecureTls() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SecretBackend) pulumi.BoolPtrOutput { return v.InsecureTls }).(pulumi.BoolPtrOutput)
}

// **Deprecated** use `passwordPolicy`. The desired length of passwords that Vault generates.
// *Mutually exclusive with `passwordPolicy` on vault-1.11+*
//
// Deprecated: Length is deprecated and passwordPolicy should be used with Vault >= 1.5.
func (o SecretBackendOutput) Length() pulumi.IntOutput {
	return o.ApplyT(func(v *SecretBackend) pulumi.IntOutput { return v.Length }).(pulumi.IntOutput)
}

// Mark the secrets engine as local-only. Local engines are not replicated or removed by
// replication.Tolerance duration to use when checking the last rotation time.
func (o SecretBackendOutput) Local() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SecretBackend) pulumi.BoolPtrOutput { return v.Local }).(pulumi.BoolPtrOutput)
}

// Maximum possible lease duration for secrets in seconds.
func (o SecretBackendOutput) MaxLeaseTtlSeconds() pulumi.IntOutput {
	return o.ApplyT(func(v *SecretBackend) pulumi.IntOutput { return v.MaxLeaseTtlSeconds }).(pulumi.IntOutput)
}

// The namespace to provision the resource in.
// The value should not contain leading or trailing forward slashes.
// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
// *Available only for Vault Enterprise*.
func (o SecretBackendOutput) Namespace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecretBackend) pulumi.StringPtrOutput { return v.Namespace }).(pulumi.StringPtrOutput)
}

// Specifies mount type specific options that are passed to the backend
func (o SecretBackendOutput) Options() pulumi.MapOutput {
	return o.ApplyT(func(v *SecretBackend) pulumi.MapOutput { return v.Options }).(pulumi.MapOutput)
}

// Name of the password policy to use to generate passwords.
func (o SecretBackendOutput) PasswordPolicy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecretBackend) pulumi.StringPtrOutput { return v.PasswordPolicy }).(pulumi.StringPtrOutput)
}

// The unique path this backend should be mounted at. Must
// not begin or end with a `/`. Defaults to `ldap`.
func (o SecretBackendOutput) Path() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecretBackend) pulumi.StringPtrOutput { return v.Path }).(pulumi.StringPtrOutput)
}

// Timeout, in seconds, for the connection when making requests against the server
// before returning back an error.
func (o SecretBackendOutput) RequestTimeout() pulumi.IntOutput {
	return o.ApplyT(func(v *SecretBackend) pulumi.IntOutput { return v.RequestTimeout }).(pulumi.IntOutput)
}

// The LDAP schema to use when storing entry passwords. Valid schemas include `openldap`, `ad`, and `racf`. Default is `openldap`.
func (o SecretBackendOutput) Schema() pulumi.StringOutput {
	return o.ApplyT(func(v *SecretBackend) pulumi.StringOutput { return v.Schema }).(pulumi.StringOutput)
}

// Enable seal wrapping for the mount, causing values stored by the mount to be wrapped by the seal's encryption capability
func (o SecretBackendOutput) SealWrap() pulumi.BoolOutput {
	return o.ApplyT(func(v *SecretBackend) pulumi.BoolOutput { return v.SealWrap }).(pulumi.BoolOutput)
}

// Issue a StartTLS command after establishing unencrypted connection.
func (o SecretBackendOutput) Starttls() pulumi.BoolOutput {
	return o.ApplyT(func(v *SecretBackend) pulumi.BoolOutput { return v.Starttls }).(pulumi.BoolOutput)
}

// Enables userPrincipalDomain login with [username]@UPNDomain.
func (o SecretBackendOutput) Upndomain() pulumi.StringOutput {
	return o.ApplyT(func(v *SecretBackend) pulumi.StringOutput { return v.Upndomain }).(pulumi.StringOutput)
}

// LDAP URL to connect to. Multiple URLs can be specified by concatenating
// them with commas; they will be tried in-order. Defaults to `ldap://127.0.0.1`.
func (o SecretBackendOutput) Url() pulumi.StringOutput {
	return o.ApplyT(func(v *SecretBackend) pulumi.StringOutput { return v.Url }).(pulumi.StringOutput)
}

// Attribute used when searching users. Defaults to `cn`.
func (o SecretBackendOutput) Userattr() pulumi.StringOutput {
	return o.ApplyT(func(v *SecretBackend) pulumi.StringOutput { return v.Userattr }).(pulumi.StringOutput)
}

// LDAP domain to use for users (eg: ou=People,dc=example,dc=org)`.
func (o SecretBackendOutput) Userdn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecretBackend) pulumi.StringPtrOutput { return v.Userdn }).(pulumi.StringPtrOutput)
}

type SecretBackendArrayOutput struct{ *pulumi.OutputState }

func (SecretBackendArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SecretBackend)(nil)).Elem()
}

func (o SecretBackendArrayOutput) ToSecretBackendArrayOutput() SecretBackendArrayOutput {
	return o
}

func (o SecretBackendArrayOutput) ToSecretBackendArrayOutputWithContext(ctx context.Context) SecretBackendArrayOutput {
	return o
}

func (o SecretBackendArrayOutput) Index(i pulumi.IntInput) SecretBackendOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SecretBackend {
		return vs[0].([]*SecretBackend)[vs[1].(int)]
	}).(SecretBackendOutput)
}

type SecretBackendMapOutput struct{ *pulumi.OutputState }

func (SecretBackendMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SecretBackend)(nil)).Elem()
}

func (o SecretBackendMapOutput) ToSecretBackendMapOutput() SecretBackendMapOutput {
	return o
}

func (o SecretBackendMapOutput) ToSecretBackendMapOutputWithContext(ctx context.Context) SecretBackendMapOutput {
	return o
}

func (o SecretBackendMapOutput) MapIndex(k pulumi.StringInput) SecretBackendOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SecretBackend {
		return vs[0].(map[string]*SecretBackend)[vs[1].(string)]
	}).(SecretBackendOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SecretBackendInput)(nil)).Elem(), &SecretBackend{})
	pulumi.RegisterInputType(reflect.TypeOf((*SecretBackendArrayInput)(nil)).Elem(), SecretBackendArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SecretBackendMapInput)(nil)).Elem(), SecretBackendMap{})
	pulumi.RegisterOutputType(SecretBackendOutput{})
	pulumi.RegisterOutputType(SecretBackendArrayOutput{})
	pulumi.RegisterOutputType(SecretBackendMapOutput{})
}
