// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ad

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## Import
//
// AD secret backend can be imported using the `backend`, e.g.
//
// ```sh
//  $ pulumi import vault:ad/secretBackend:SecretBackend ad ad
// ```
type SecretBackend struct {
	pulumi.CustomResourceState

	// Use anonymous binds when performing LDAP group searches
	// (if true the initial credentials will still be used for the initial connection test).
	AnonymousGroupSearch pulumi.BoolPtrOutput `pulumi:"anonymousGroupSearch"`
	// The unique path this backend should be mounted at. Must
	// not begin or end with a `/`. Defaults to `ad`.
	Backend pulumi.StringPtrOutput `pulumi:"backend"`
	// Distinguished name of object to bind when performing user and group search.
	Binddn pulumi.StringOutput `pulumi:"binddn"`
	// Password to use along with binddn when performing user search.
	Bindpass pulumi.StringOutput `pulumi:"bindpass"`
	// If set, user and group names assigned to policies within the
	// backend will be case sensitive. Otherwise, names will be normalized to lower case.
	CaseSensitiveNames pulumi.BoolPtrOutput `pulumi:"caseSensitiveNames"`
	// CA certificate to use when verifying LDAP server certificate, must be
	// x509 PEM encoded.
	Certificate pulumi.StringPtrOutput `pulumi:"certificate"`
	// Client certificate to provide to the LDAP server, must be x509 PEM encoded.
	ClientTlsCert pulumi.StringPtrOutput `pulumi:"clientTlsCert"`
	// Client certificate key to provide to the LDAP server, must be x509 PEM encoded.
	ClientTlsKey pulumi.StringPtrOutput `pulumi:"clientTlsKey"`
	// Default lease duration for secrets in seconds.
	DefaultLeaseTtlSeconds pulumi.IntOutput `pulumi:"defaultLeaseTtlSeconds"`
	// Denies an unauthenticated LDAP bind request if the user's password is empty;
	// defaults to true.
	DenyNullBind pulumi.BoolPtrOutput `pulumi:"denyNullBind"`
	// Human-friendly description of the mount for the Active Directory backend.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Use anonymous bind to discover the bind Distinguished Name of a user.
	Discoverdn pulumi.BoolPtrOutput `pulumi:"discoverdn"`
	// Text to insert the password into, ex. "customPrefix{{PASSWORD}}customSuffix". This
	// setting is deprecated and should instead use `passwordPolicy`.
	//
	// Deprecated: Formatter is deprecated and password_policy should be used with Vault >= 1.5.
	Formatter pulumi.StringOutput `pulumi:"formatter"`
	// LDAP attribute to follow on objects returned by <groupfilter> in order to enumerate
	// user group membership. Examples: `cn` or `memberOf`, etc. Defaults to `cn`.
	Groupattr pulumi.StringPtrOutput `pulumi:"groupattr"`
	// LDAP search base to use for group membership search (eg: ou=Groups,dc=example,dc=org).
	Groupdn pulumi.StringPtrOutput `pulumi:"groupdn"`
	// Go template for querying group membership of user (optional) The template can access
	// the following context variables: UserDN, Username. Defaults to `(|(memberUid={{.Username}})(member={{.UserDN}})(uniqueMember={{.UserDN}}))`
	Groupfilter pulumi.StringPtrOutput `pulumi:"groupfilter"`
	// Skip LDAP server SSL Certificate verification. This is not recommended for production.
	// Defaults to `false`.
	InsecureTls pulumi.BoolPtrOutput `pulumi:"insecureTls"`
	// The number of seconds after a Vault rotation where, if Active Directory
	// shows a later rotation, it should be considered out-of-band
	LastRotationTolerance pulumi.IntOutput `pulumi:"lastRotationTolerance"`
	// The desired length of passwords that Vault generates. This
	// setting is deprecated and should instead use `passwordPolicy`.
	//
	// Deprecated: Length is deprecated and password_policy should be used with Vault >= 1.5.
	Length pulumi.IntOutput `pulumi:"length"`
	// Mark the secrets engine as local-only. Local engines are not replicated or removed by
	// replication.Tolerance duration to use when checking the last rotation time.
	Local pulumi.BoolPtrOutput `pulumi:"local"`
	// Maximum possible lease duration for secrets in seconds.
	MaxLeaseTtlSeconds pulumi.IntOutput `pulumi:"maxLeaseTtlSeconds"`
	// In seconds, the maximum password time-to-live.
	MaxTtl pulumi.IntOutput `pulumi:"maxTtl"`
	// Name of the password policy to use to generate passwords.
	PasswordPolicy pulumi.StringPtrOutput `pulumi:"passwordPolicy"`
	// Timeout, in seconds, for the connection when making requests against the server
	// before returning back an error.
	RequestTimeout pulumi.IntPtrOutput `pulumi:"requestTimeout"`
	// Issue a StartTLS command after establishing unencrypted connection.
	Starttls pulumi.BoolOutput `pulumi:"starttls"`
	// Maximum TLS version to use. Accepted values are `tls10`, `tls11`,
	// `tls12` or `tls13`. Defaults to `tls12`.
	TlsMaxVersion pulumi.StringOutput `pulumi:"tlsMaxVersion"`
	// Minimum TLS version to use. Accepted values are `tls10`, `tls11`,
	// `tls12` or `tls13`. Defaults to `tls12`.
	TlsMinVersion pulumi.StringOutput `pulumi:"tlsMinVersion"`
	// In seconds, the default password time-to-live.
	Ttl pulumi.IntOutput `pulumi:"ttl"`
	// Enables userPrincipalDomain login with [username]@UPNDomain.
	Upndomain pulumi.StringOutput `pulumi:"upndomain"`
	// LDAP URL to connect to. Multiple URLs can be specified by concatenating
	// them with commas; they will be tried in-order. Defaults to `ldap://127.0.0.1`.
	Url pulumi.StringPtrOutput `pulumi:"url"`
	// In Vault 1.1.1 a fix for handling group CN values of
	// different cases unfortunately introduced a regression that could cause previously defined groups
	// to not be found due to a change in the resulting name. If set true, the pre-1.1.1 behavior for
	// matching group CNs will be used. This is only needed in some upgrade scenarios for backwards
	// compatibility. It is enabled by default if the config is upgraded but disabled by default on
	// new configurations.
	UsePre111GroupCnBehavior pulumi.BoolOutput `pulumi:"usePre111GroupCnBehavior"`
	// If true, use the Active Directory tokenGroups constructed attribute of the
	// user to find the group memberships. This will find all security groups including nested ones.
	UseTokenGroups pulumi.BoolPtrOutput `pulumi:"useTokenGroups"`
	// Attribute used when searching users. Defaults to `cn`.
	Userattr pulumi.StringPtrOutput `pulumi:"userattr"`
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
	var resource SecretBackend
	err := ctx.RegisterResource("vault:ad/secretBackend:SecretBackend", name, args, &resource, opts...)
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
	err := ctx.ReadResource("vault:ad/secretBackend:SecretBackend", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SecretBackend resources.
type secretBackendState struct {
	// Use anonymous binds when performing LDAP group searches
	// (if true the initial credentials will still be used for the initial connection test).
	AnonymousGroupSearch *bool `pulumi:"anonymousGroupSearch"`
	// The unique path this backend should be mounted at. Must
	// not begin or end with a `/`. Defaults to `ad`.
	Backend *string `pulumi:"backend"`
	// Distinguished name of object to bind when performing user and group search.
	Binddn *string `pulumi:"binddn"`
	// Password to use along with binddn when performing user search.
	Bindpass *string `pulumi:"bindpass"`
	// If set, user and group names assigned to policies within the
	// backend will be case sensitive. Otherwise, names will be normalized to lower case.
	CaseSensitiveNames *bool `pulumi:"caseSensitiveNames"`
	// CA certificate to use when verifying LDAP server certificate, must be
	// x509 PEM encoded.
	Certificate *string `pulumi:"certificate"`
	// Client certificate to provide to the LDAP server, must be x509 PEM encoded.
	ClientTlsCert *string `pulumi:"clientTlsCert"`
	// Client certificate key to provide to the LDAP server, must be x509 PEM encoded.
	ClientTlsKey *string `pulumi:"clientTlsKey"`
	// Default lease duration for secrets in seconds.
	DefaultLeaseTtlSeconds *int `pulumi:"defaultLeaseTtlSeconds"`
	// Denies an unauthenticated LDAP bind request if the user's password is empty;
	// defaults to true.
	DenyNullBind *bool `pulumi:"denyNullBind"`
	// Human-friendly description of the mount for the Active Directory backend.
	Description *string `pulumi:"description"`
	// Use anonymous bind to discover the bind Distinguished Name of a user.
	Discoverdn *bool `pulumi:"discoverdn"`
	// Text to insert the password into, ex. "customPrefix{{PASSWORD}}customSuffix". This
	// setting is deprecated and should instead use `passwordPolicy`.
	//
	// Deprecated: Formatter is deprecated and password_policy should be used with Vault >= 1.5.
	Formatter *string `pulumi:"formatter"`
	// LDAP attribute to follow on objects returned by <groupfilter> in order to enumerate
	// user group membership. Examples: `cn` or `memberOf`, etc. Defaults to `cn`.
	Groupattr *string `pulumi:"groupattr"`
	// LDAP search base to use for group membership search (eg: ou=Groups,dc=example,dc=org).
	Groupdn *string `pulumi:"groupdn"`
	// Go template for querying group membership of user (optional) The template can access
	// the following context variables: UserDN, Username. Defaults to `(|(memberUid={{.Username}})(member={{.UserDN}})(uniqueMember={{.UserDN}}))`
	Groupfilter *string `pulumi:"groupfilter"`
	// Skip LDAP server SSL Certificate verification. This is not recommended for production.
	// Defaults to `false`.
	InsecureTls *bool `pulumi:"insecureTls"`
	// The number of seconds after a Vault rotation where, if Active Directory
	// shows a later rotation, it should be considered out-of-band
	LastRotationTolerance *int `pulumi:"lastRotationTolerance"`
	// The desired length of passwords that Vault generates. This
	// setting is deprecated and should instead use `passwordPolicy`.
	//
	// Deprecated: Length is deprecated and password_policy should be used with Vault >= 1.5.
	Length *int `pulumi:"length"`
	// Mark the secrets engine as local-only. Local engines are not replicated or removed by
	// replication.Tolerance duration to use when checking the last rotation time.
	Local *bool `pulumi:"local"`
	// Maximum possible lease duration for secrets in seconds.
	MaxLeaseTtlSeconds *int `pulumi:"maxLeaseTtlSeconds"`
	// In seconds, the maximum password time-to-live.
	MaxTtl *int `pulumi:"maxTtl"`
	// Name of the password policy to use to generate passwords.
	PasswordPolicy *string `pulumi:"passwordPolicy"`
	// Timeout, in seconds, for the connection when making requests against the server
	// before returning back an error.
	RequestTimeout *int `pulumi:"requestTimeout"`
	// Issue a StartTLS command after establishing unencrypted connection.
	Starttls *bool `pulumi:"starttls"`
	// Maximum TLS version to use. Accepted values are `tls10`, `tls11`,
	// `tls12` or `tls13`. Defaults to `tls12`.
	TlsMaxVersion *string `pulumi:"tlsMaxVersion"`
	// Minimum TLS version to use. Accepted values are `tls10`, `tls11`,
	// `tls12` or `tls13`. Defaults to `tls12`.
	TlsMinVersion *string `pulumi:"tlsMinVersion"`
	// In seconds, the default password time-to-live.
	Ttl *int `pulumi:"ttl"`
	// Enables userPrincipalDomain login with [username]@UPNDomain.
	Upndomain *string `pulumi:"upndomain"`
	// LDAP URL to connect to. Multiple URLs can be specified by concatenating
	// them with commas; they will be tried in-order. Defaults to `ldap://127.0.0.1`.
	Url *string `pulumi:"url"`
	// In Vault 1.1.1 a fix for handling group CN values of
	// different cases unfortunately introduced a regression that could cause previously defined groups
	// to not be found due to a change in the resulting name. If set true, the pre-1.1.1 behavior for
	// matching group CNs will be used. This is only needed in some upgrade scenarios for backwards
	// compatibility. It is enabled by default if the config is upgraded but disabled by default on
	// new configurations.
	UsePre111GroupCnBehavior *bool `pulumi:"usePre111GroupCnBehavior"`
	// If true, use the Active Directory tokenGroups constructed attribute of the
	// user to find the group memberships. This will find all security groups including nested ones.
	UseTokenGroups *bool `pulumi:"useTokenGroups"`
	// Attribute used when searching users. Defaults to `cn`.
	Userattr *string `pulumi:"userattr"`
	// LDAP domain to use for users (eg: ou=People,dc=example,dc=org)`.
	Userdn *string `pulumi:"userdn"`
}

type SecretBackendState struct {
	// Use anonymous binds when performing LDAP group searches
	// (if true the initial credentials will still be used for the initial connection test).
	AnonymousGroupSearch pulumi.BoolPtrInput
	// The unique path this backend should be mounted at. Must
	// not begin or end with a `/`. Defaults to `ad`.
	Backend pulumi.StringPtrInput
	// Distinguished name of object to bind when performing user and group search.
	Binddn pulumi.StringPtrInput
	// Password to use along with binddn when performing user search.
	Bindpass pulumi.StringPtrInput
	// If set, user and group names assigned to policies within the
	// backend will be case sensitive. Otherwise, names will be normalized to lower case.
	CaseSensitiveNames pulumi.BoolPtrInput
	// CA certificate to use when verifying LDAP server certificate, must be
	// x509 PEM encoded.
	Certificate pulumi.StringPtrInput
	// Client certificate to provide to the LDAP server, must be x509 PEM encoded.
	ClientTlsCert pulumi.StringPtrInput
	// Client certificate key to provide to the LDAP server, must be x509 PEM encoded.
	ClientTlsKey pulumi.StringPtrInput
	// Default lease duration for secrets in seconds.
	DefaultLeaseTtlSeconds pulumi.IntPtrInput
	// Denies an unauthenticated LDAP bind request if the user's password is empty;
	// defaults to true.
	DenyNullBind pulumi.BoolPtrInput
	// Human-friendly description of the mount for the Active Directory backend.
	Description pulumi.StringPtrInput
	// Use anonymous bind to discover the bind Distinguished Name of a user.
	Discoverdn pulumi.BoolPtrInput
	// Text to insert the password into, ex. "customPrefix{{PASSWORD}}customSuffix". This
	// setting is deprecated and should instead use `passwordPolicy`.
	//
	// Deprecated: Formatter is deprecated and password_policy should be used with Vault >= 1.5.
	Formatter pulumi.StringPtrInput
	// LDAP attribute to follow on objects returned by <groupfilter> in order to enumerate
	// user group membership. Examples: `cn` or `memberOf`, etc. Defaults to `cn`.
	Groupattr pulumi.StringPtrInput
	// LDAP search base to use for group membership search (eg: ou=Groups,dc=example,dc=org).
	Groupdn pulumi.StringPtrInput
	// Go template for querying group membership of user (optional) The template can access
	// the following context variables: UserDN, Username. Defaults to `(|(memberUid={{.Username}})(member={{.UserDN}})(uniqueMember={{.UserDN}}))`
	Groupfilter pulumi.StringPtrInput
	// Skip LDAP server SSL Certificate verification. This is not recommended for production.
	// Defaults to `false`.
	InsecureTls pulumi.BoolPtrInput
	// The number of seconds after a Vault rotation where, if Active Directory
	// shows a later rotation, it should be considered out-of-band
	LastRotationTolerance pulumi.IntPtrInput
	// The desired length of passwords that Vault generates. This
	// setting is deprecated and should instead use `passwordPolicy`.
	//
	// Deprecated: Length is deprecated and password_policy should be used with Vault >= 1.5.
	Length pulumi.IntPtrInput
	// Mark the secrets engine as local-only. Local engines are not replicated or removed by
	// replication.Tolerance duration to use when checking the last rotation time.
	Local pulumi.BoolPtrInput
	// Maximum possible lease duration for secrets in seconds.
	MaxLeaseTtlSeconds pulumi.IntPtrInput
	// In seconds, the maximum password time-to-live.
	MaxTtl pulumi.IntPtrInput
	// Name of the password policy to use to generate passwords.
	PasswordPolicy pulumi.StringPtrInput
	// Timeout, in seconds, for the connection when making requests against the server
	// before returning back an error.
	RequestTimeout pulumi.IntPtrInput
	// Issue a StartTLS command after establishing unencrypted connection.
	Starttls pulumi.BoolPtrInput
	// Maximum TLS version to use. Accepted values are `tls10`, `tls11`,
	// `tls12` or `tls13`. Defaults to `tls12`.
	TlsMaxVersion pulumi.StringPtrInput
	// Minimum TLS version to use. Accepted values are `tls10`, `tls11`,
	// `tls12` or `tls13`. Defaults to `tls12`.
	TlsMinVersion pulumi.StringPtrInput
	// In seconds, the default password time-to-live.
	Ttl pulumi.IntPtrInput
	// Enables userPrincipalDomain login with [username]@UPNDomain.
	Upndomain pulumi.StringPtrInput
	// LDAP URL to connect to. Multiple URLs can be specified by concatenating
	// them with commas; they will be tried in-order. Defaults to `ldap://127.0.0.1`.
	Url pulumi.StringPtrInput
	// In Vault 1.1.1 a fix for handling group CN values of
	// different cases unfortunately introduced a regression that could cause previously defined groups
	// to not be found due to a change in the resulting name. If set true, the pre-1.1.1 behavior for
	// matching group CNs will be used. This is only needed in some upgrade scenarios for backwards
	// compatibility. It is enabled by default if the config is upgraded but disabled by default on
	// new configurations.
	UsePre111GroupCnBehavior pulumi.BoolPtrInput
	// If true, use the Active Directory tokenGroups constructed attribute of the
	// user to find the group memberships. This will find all security groups including nested ones.
	UseTokenGroups pulumi.BoolPtrInput
	// Attribute used when searching users. Defaults to `cn`.
	Userattr pulumi.StringPtrInput
	// LDAP domain to use for users (eg: ou=People,dc=example,dc=org)`.
	Userdn pulumi.StringPtrInput
}

func (SecretBackendState) ElementType() reflect.Type {
	return reflect.TypeOf((*secretBackendState)(nil)).Elem()
}

type secretBackendArgs struct {
	// Use anonymous binds when performing LDAP group searches
	// (if true the initial credentials will still be used for the initial connection test).
	AnonymousGroupSearch *bool `pulumi:"anonymousGroupSearch"`
	// The unique path this backend should be mounted at. Must
	// not begin or end with a `/`. Defaults to `ad`.
	Backend *string `pulumi:"backend"`
	// Distinguished name of object to bind when performing user and group search.
	Binddn string `pulumi:"binddn"`
	// Password to use along with binddn when performing user search.
	Bindpass string `pulumi:"bindpass"`
	// If set, user and group names assigned to policies within the
	// backend will be case sensitive. Otherwise, names will be normalized to lower case.
	CaseSensitiveNames *bool `pulumi:"caseSensitiveNames"`
	// CA certificate to use when verifying LDAP server certificate, must be
	// x509 PEM encoded.
	Certificate *string `pulumi:"certificate"`
	// Client certificate to provide to the LDAP server, must be x509 PEM encoded.
	ClientTlsCert *string `pulumi:"clientTlsCert"`
	// Client certificate key to provide to the LDAP server, must be x509 PEM encoded.
	ClientTlsKey *string `pulumi:"clientTlsKey"`
	// Default lease duration for secrets in seconds.
	DefaultLeaseTtlSeconds *int `pulumi:"defaultLeaseTtlSeconds"`
	// Denies an unauthenticated LDAP bind request if the user's password is empty;
	// defaults to true.
	DenyNullBind *bool `pulumi:"denyNullBind"`
	// Human-friendly description of the mount for the Active Directory backend.
	Description *string `pulumi:"description"`
	// Use anonymous bind to discover the bind Distinguished Name of a user.
	Discoverdn *bool `pulumi:"discoverdn"`
	// Text to insert the password into, ex. "customPrefix{{PASSWORD}}customSuffix". This
	// setting is deprecated and should instead use `passwordPolicy`.
	//
	// Deprecated: Formatter is deprecated and password_policy should be used with Vault >= 1.5.
	Formatter *string `pulumi:"formatter"`
	// LDAP attribute to follow on objects returned by <groupfilter> in order to enumerate
	// user group membership. Examples: `cn` or `memberOf`, etc. Defaults to `cn`.
	Groupattr *string `pulumi:"groupattr"`
	// LDAP search base to use for group membership search (eg: ou=Groups,dc=example,dc=org).
	Groupdn *string `pulumi:"groupdn"`
	// Go template for querying group membership of user (optional) The template can access
	// the following context variables: UserDN, Username. Defaults to `(|(memberUid={{.Username}})(member={{.UserDN}})(uniqueMember={{.UserDN}}))`
	Groupfilter *string `pulumi:"groupfilter"`
	// Skip LDAP server SSL Certificate verification. This is not recommended for production.
	// Defaults to `false`.
	InsecureTls *bool `pulumi:"insecureTls"`
	// The number of seconds after a Vault rotation where, if Active Directory
	// shows a later rotation, it should be considered out-of-band
	LastRotationTolerance *int `pulumi:"lastRotationTolerance"`
	// The desired length of passwords that Vault generates. This
	// setting is deprecated and should instead use `passwordPolicy`.
	//
	// Deprecated: Length is deprecated and password_policy should be used with Vault >= 1.5.
	Length *int `pulumi:"length"`
	// Mark the secrets engine as local-only. Local engines are not replicated or removed by
	// replication.Tolerance duration to use when checking the last rotation time.
	Local *bool `pulumi:"local"`
	// Maximum possible lease duration for secrets in seconds.
	MaxLeaseTtlSeconds *int `pulumi:"maxLeaseTtlSeconds"`
	// In seconds, the maximum password time-to-live.
	MaxTtl *int `pulumi:"maxTtl"`
	// Name of the password policy to use to generate passwords.
	PasswordPolicy *string `pulumi:"passwordPolicy"`
	// Timeout, in seconds, for the connection when making requests against the server
	// before returning back an error.
	RequestTimeout *int `pulumi:"requestTimeout"`
	// Issue a StartTLS command after establishing unencrypted connection.
	Starttls *bool `pulumi:"starttls"`
	// Maximum TLS version to use. Accepted values are `tls10`, `tls11`,
	// `tls12` or `tls13`. Defaults to `tls12`.
	TlsMaxVersion *string `pulumi:"tlsMaxVersion"`
	// Minimum TLS version to use. Accepted values are `tls10`, `tls11`,
	// `tls12` or `tls13`. Defaults to `tls12`.
	TlsMinVersion *string `pulumi:"tlsMinVersion"`
	// In seconds, the default password time-to-live.
	Ttl *int `pulumi:"ttl"`
	// Enables userPrincipalDomain login with [username]@UPNDomain.
	Upndomain *string `pulumi:"upndomain"`
	// LDAP URL to connect to. Multiple URLs can be specified by concatenating
	// them with commas; they will be tried in-order. Defaults to `ldap://127.0.0.1`.
	Url *string `pulumi:"url"`
	// In Vault 1.1.1 a fix for handling group CN values of
	// different cases unfortunately introduced a regression that could cause previously defined groups
	// to not be found due to a change in the resulting name. If set true, the pre-1.1.1 behavior for
	// matching group CNs will be used. This is only needed in some upgrade scenarios for backwards
	// compatibility. It is enabled by default if the config is upgraded but disabled by default on
	// new configurations.
	UsePre111GroupCnBehavior *bool `pulumi:"usePre111GroupCnBehavior"`
	// If true, use the Active Directory tokenGroups constructed attribute of the
	// user to find the group memberships. This will find all security groups including nested ones.
	UseTokenGroups *bool `pulumi:"useTokenGroups"`
	// Attribute used when searching users. Defaults to `cn`.
	Userattr *string `pulumi:"userattr"`
	// LDAP domain to use for users (eg: ou=People,dc=example,dc=org)`.
	Userdn *string `pulumi:"userdn"`
}

// The set of arguments for constructing a SecretBackend resource.
type SecretBackendArgs struct {
	// Use anonymous binds when performing LDAP group searches
	// (if true the initial credentials will still be used for the initial connection test).
	AnonymousGroupSearch pulumi.BoolPtrInput
	// The unique path this backend should be mounted at. Must
	// not begin or end with a `/`. Defaults to `ad`.
	Backend pulumi.StringPtrInput
	// Distinguished name of object to bind when performing user and group search.
	Binddn pulumi.StringInput
	// Password to use along with binddn when performing user search.
	Bindpass pulumi.StringInput
	// If set, user and group names assigned to policies within the
	// backend will be case sensitive. Otherwise, names will be normalized to lower case.
	CaseSensitiveNames pulumi.BoolPtrInput
	// CA certificate to use when verifying LDAP server certificate, must be
	// x509 PEM encoded.
	Certificate pulumi.StringPtrInput
	// Client certificate to provide to the LDAP server, must be x509 PEM encoded.
	ClientTlsCert pulumi.StringPtrInput
	// Client certificate key to provide to the LDAP server, must be x509 PEM encoded.
	ClientTlsKey pulumi.StringPtrInput
	// Default lease duration for secrets in seconds.
	DefaultLeaseTtlSeconds pulumi.IntPtrInput
	// Denies an unauthenticated LDAP bind request if the user's password is empty;
	// defaults to true.
	DenyNullBind pulumi.BoolPtrInput
	// Human-friendly description of the mount for the Active Directory backend.
	Description pulumi.StringPtrInput
	// Use anonymous bind to discover the bind Distinguished Name of a user.
	Discoverdn pulumi.BoolPtrInput
	// Text to insert the password into, ex. "customPrefix{{PASSWORD}}customSuffix". This
	// setting is deprecated and should instead use `passwordPolicy`.
	//
	// Deprecated: Formatter is deprecated and password_policy should be used with Vault >= 1.5.
	Formatter pulumi.StringPtrInput
	// LDAP attribute to follow on objects returned by <groupfilter> in order to enumerate
	// user group membership. Examples: `cn` or `memberOf`, etc. Defaults to `cn`.
	Groupattr pulumi.StringPtrInput
	// LDAP search base to use for group membership search (eg: ou=Groups,dc=example,dc=org).
	Groupdn pulumi.StringPtrInput
	// Go template for querying group membership of user (optional) The template can access
	// the following context variables: UserDN, Username. Defaults to `(|(memberUid={{.Username}})(member={{.UserDN}})(uniqueMember={{.UserDN}}))`
	Groupfilter pulumi.StringPtrInput
	// Skip LDAP server SSL Certificate verification. This is not recommended for production.
	// Defaults to `false`.
	InsecureTls pulumi.BoolPtrInput
	// The number of seconds after a Vault rotation where, if Active Directory
	// shows a later rotation, it should be considered out-of-band
	LastRotationTolerance pulumi.IntPtrInput
	// The desired length of passwords that Vault generates. This
	// setting is deprecated and should instead use `passwordPolicy`.
	//
	// Deprecated: Length is deprecated and password_policy should be used with Vault >= 1.5.
	Length pulumi.IntPtrInput
	// Mark the secrets engine as local-only. Local engines are not replicated or removed by
	// replication.Tolerance duration to use when checking the last rotation time.
	Local pulumi.BoolPtrInput
	// Maximum possible lease duration for secrets in seconds.
	MaxLeaseTtlSeconds pulumi.IntPtrInput
	// In seconds, the maximum password time-to-live.
	MaxTtl pulumi.IntPtrInput
	// Name of the password policy to use to generate passwords.
	PasswordPolicy pulumi.StringPtrInput
	// Timeout, in seconds, for the connection when making requests against the server
	// before returning back an error.
	RequestTimeout pulumi.IntPtrInput
	// Issue a StartTLS command after establishing unencrypted connection.
	Starttls pulumi.BoolPtrInput
	// Maximum TLS version to use. Accepted values are `tls10`, `tls11`,
	// `tls12` or `tls13`. Defaults to `tls12`.
	TlsMaxVersion pulumi.StringPtrInput
	// Minimum TLS version to use. Accepted values are `tls10`, `tls11`,
	// `tls12` or `tls13`. Defaults to `tls12`.
	TlsMinVersion pulumi.StringPtrInput
	// In seconds, the default password time-to-live.
	Ttl pulumi.IntPtrInput
	// Enables userPrincipalDomain login with [username]@UPNDomain.
	Upndomain pulumi.StringPtrInput
	// LDAP URL to connect to. Multiple URLs can be specified by concatenating
	// them with commas; they will be tried in-order. Defaults to `ldap://127.0.0.1`.
	Url pulumi.StringPtrInput
	// In Vault 1.1.1 a fix for handling group CN values of
	// different cases unfortunately introduced a regression that could cause previously defined groups
	// to not be found due to a change in the resulting name. If set true, the pre-1.1.1 behavior for
	// matching group CNs will be used. This is only needed in some upgrade scenarios for backwards
	// compatibility. It is enabled by default if the config is upgraded but disabled by default on
	// new configurations.
	UsePre111GroupCnBehavior pulumi.BoolPtrInput
	// If true, use the Active Directory tokenGroups constructed attribute of the
	// user to find the group memberships. This will find all security groups including nested ones.
	UseTokenGroups pulumi.BoolPtrInput
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
	return reflect.TypeOf((*SecretBackend)(nil))
}

func (i *SecretBackend) ToSecretBackendOutput() SecretBackendOutput {
	return i.ToSecretBackendOutputWithContext(context.Background())
}

func (i *SecretBackend) ToSecretBackendOutputWithContext(ctx context.Context) SecretBackendOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecretBackendOutput)
}

func (i *SecretBackend) ToSecretBackendPtrOutput() SecretBackendPtrOutput {
	return i.ToSecretBackendPtrOutputWithContext(context.Background())
}

func (i *SecretBackend) ToSecretBackendPtrOutputWithContext(ctx context.Context) SecretBackendPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecretBackendPtrOutput)
}

type SecretBackendPtrInput interface {
	pulumi.Input

	ToSecretBackendPtrOutput() SecretBackendPtrOutput
	ToSecretBackendPtrOutputWithContext(ctx context.Context) SecretBackendPtrOutput
}

type secretBackendPtrType SecretBackendArgs

func (*secretBackendPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SecretBackend)(nil))
}

func (i *secretBackendPtrType) ToSecretBackendPtrOutput() SecretBackendPtrOutput {
	return i.ToSecretBackendPtrOutputWithContext(context.Background())
}

func (i *secretBackendPtrType) ToSecretBackendPtrOutputWithContext(ctx context.Context) SecretBackendPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecretBackendPtrOutput)
}

// SecretBackendArrayInput is an input type that accepts SecretBackendArray and SecretBackendArrayOutput values.
// You can construct a concrete instance of `SecretBackendArrayInput` via:
//
//          SecretBackendArray{ SecretBackendArgs{...} }
type SecretBackendArrayInput interface {
	pulumi.Input

	ToSecretBackendArrayOutput() SecretBackendArrayOutput
	ToSecretBackendArrayOutputWithContext(context.Context) SecretBackendArrayOutput
}

type SecretBackendArray []SecretBackendInput

func (SecretBackendArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*SecretBackend)(nil))
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
//          SecretBackendMap{ "key": SecretBackendArgs{...} }
type SecretBackendMapInput interface {
	pulumi.Input

	ToSecretBackendMapOutput() SecretBackendMapOutput
	ToSecretBackendMapOutputWithContext(context.Context) SecretBackendMapOutput
}

type SecretBackendMap map[string]SecretBackendInput

func (SecretBackendMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*SecretBackend)(nil))
}

func (i SecretBackendMap) ToSecretBackendMapOutput() SecretBackendMapOutput {
	return i.ToSecretBackendMapOutputWithContext(context.Background())
}

func (i SecretBackendMap) ToSecretBackendMapOutputWithContext(ctx context.Context) SecretBackendMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecretBackendMapOutput)
}

type SecretBackendOutput struct {
	*pulumi.OutputState
}

func (SecretBackendOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SecretBackend)(nil))
}

func (o SecretBackendOutput) ToSecretBackendOutput() SecretBackendOutput {
	return o
}

func (o SecretBackendOutput) ToSecretBackendOutputWithContext(ctx context.Context) SecretBackendOutput {
	return o
}

func (o SecretBackendOutput) ToSecretBackendPtrOutput() SecretBackendPtrOutput {
	return o.ToSecretBackendPtrOutputWithContext(context.Background())
}

func (o SecretBackendOutput) ToSecretBackendPtrOutputWithContext(ctx context.Context) SecretBackendPtrOutput {
	return o.ApplyT(func(v SecretBackend) *SecretBackend {
		return &v
	}).(SecretBackendPtrOutput)
}

type SecretBackendPtrOutput struct {
	*pulumi.OutputState
}

func (SecretBackendPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SecretBackend)(nil))
}

func (o SecretBackendPtrOutput) ToSecretBackendPtrOutput() SecretBackendPtrOutput {
	return o
}

func (o SecretBackendPtrOutput) ToSecretBackendPtrOutputWithContext(ctx context.Context) SecretBackendPtrOutput {
	return o
}

type SecretBackendArrayOutput struct{ *pulumi.OutputState }

func (SecretBackendArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SecretBackend)(nil))
}

func (o SecretBackendArrayOutput) ToSecretBackendArrayOutput() SecretBackendArrayOutput {
	return o
}

func (o SecretBackendArrayOutput) ToSecretBackendArrayOutputWithContext(ctx context.Context) SecretBackendArrayOutput {
	return o
}

func (o SecretBackendArrayOutput) Index(i pulumi.IntInput) SecretBackendOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SecretBackend {
		return vs[0].([]SecretBackend)[vs[1].(int)]
	}).(SecretBackendOutput)
}

type SecretBackendMapOutput struct{ *pulumi.OutputState }

func (SecretBackendMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]SecretBackend)(nil))
}

func (o SecretBackendMapOutput) ToSecretBackendMapOutput() SecretBackendMapOutput {
	return o
}

func (o SecretBackendMapOutput) ToSecretBackendMapOutputWithContext(ctx context.Context) SecretBackendMapOutput {
	return o
}

func (o SecretBackendMapOutput) MapIndex(k pulumi.StringInput) SecretBackendOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) SecretBackend {
		return vs[0].(map[string]SecretBackend)[vs[1].(string)]
	}).(SecretBackendOutput)
}

func init() {
	pulumi.RegisterOutputType(SecretBackendOutput{})
	pulumi.RegisterOutputType(SecretBackendPtrOutput{})
	pulumi.RegisterOutputType(SecretBackendArrayOutput{})
	pulumi.RegisterOutputType(SecretBackendMapOutput{})
}
