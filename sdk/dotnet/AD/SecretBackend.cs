// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vault.AD
{
    /// <summary>
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Vault = Pulumi.Vault;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var config = new Vault.AD.SecretBackend("config", new()
    ///     {
    ///         Backend = "ad",
    ///         Binddn = "CN=Administrator,CN=Users,DC=corp,DC=example,DC=net",
    ///         Bindpass = "SuperSecretPassw0rd",
    ///         Url = "ldaps://ad",
    ///         InsecureTls = true,
    ///         Userdn = "CN=Users,DC=corp,DC=example,DC=net",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// AD secret backend can be imported using the `backend`, e.g.
    /// 
    /// ```sh
    /// $ pulumi import vault:ad/secretBackend:SecretBackend ad ad
    /// ```
    /// </summary>
    [VaultResourceType("vault:ad/secretBackend:SecretBackend")]
    public partial class SecretBackend : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Use anonymous binds when performing LDAP group searches
        /// (if true the initial credentials will still be used for the initial connection test).
        /// </summary>
        [Output("anonymousGroupSearch")]
        public Output<bool?> AnonymousGroupSearch { get; private set; } = null!;

        /// <summary>
        /// The unique path this backend should be mounted at. Must
        /// not begin or end with a `/`. Defaults to `ad`.
        /// </summary>
        [Output("backend")]
        public Output<string?> Backend { get; private set; } = null!;

        /// <summary>
        /// Distinguished name of object to bind when performing user and group search.
        /// </summary>
        [Output("binddn")]
        public Output<string> Binddn { get; private set; } = null!;

        /// <summary>
        /// Password to use along with binddn when performing user search.
        /// </summary>
        [Output("bindpass")]
        public Output<string> Bindpass { get; private set; } = null!;

        /// <summary>
        /// If set, user and group names assigned to policies within the
        /// backend will be case sensitive. Otherwise, names will be normalized to lower case.
        /// </summary>
        [Output("caseSensitiveNames")]
        public Output<bool?> CaseSensitiveNames { get; private set; } = null!;

        /// <summary>
        /// CA certificate to use when verifying LDAP server certificate, must be
        /// x509 PEM encoded.
        /// </summary>
        [Output("certificate")]
        public Output<string?> Certificate { get; private set; } = null!;

        /// <summary>
        /// Client certificate to provide to the LDAP server, must be x509 PEM encoded.
        /// </summary>
        [Output("clientTlsCert")]
        public Output<string?> ClientTlsCert { get; private set; } = null!;

        /// <summary>
        /// Client certificate key to provide to the LDAP server, must be x509 PEM encoded.
        /// </summary>
        [Output("clientTlsKey")]
        public Output<string?> ClientTlsKey { get; private set; } = null!;

        /// <summary>
        /// Default lease duration for secrets in seconds.
        /// </summary>
        [Output("defaultLeaseTtlSeconds")]
        public Output<int> DefaultLeaseTtlSeconds { get; private set; } = null!;

        /// <summary>
        /// Denies an unauthenticated LDAP bind request if the user's password is empty;
        /// defaults to true.
        /// </summary>
        [Output("denyNullBind")]
        public Output<bool?> DenyNullBind { get; private set; } = null!;

        /// <summary>
        /// Human-friendly description of the mount for the Active Directory backend.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// If set, opts out of mount migration on path updates.
        /// See here for more info on [Mount Migration](https://www.vaultproject.io/docs/concepts/mount-migration)
        /// </summary>
        [Output("disableRemount")]
        public Output<bool?> DisableRemount { get; private set; } = null!;

        /// <summary>
        /// Use anonymous bind to discover the bind Distinguished Name of a user.
        /// </summary>
        [Output("discoverdn")]
        public Output<bool?> Discoverdn { get; private set; } = null!;

        /// <summary>
        /// LDAP attribute to follow on objects returned by &lt;groupfilter&gt; in order to enumerate
        /// user group membership. Examples: `cn` or `memberOf`, etc. Defaults to `cn`.
        /// </summary>
        [Output("groupattr")]
        public Output<string?> Groupattr { get; private set; } = null!;

        /// <summary>
        /// LDAP search base to use for group membership search (eg: ou=Groups,dc=example,dc=org).
        /// </summary>
        [Output("groupdn")]
        public Output<string?> Groupdn { get; private set; } = null!;

        /// <summary>
        /// Go template for querying group membership of user (optional) The template can access
        /// the following context variables: UserDN, Username. Defaults to `(|(memberUid={{.Username}})(member={{.UserDN}})(uniqueMember={{.UserDN}}))`
        /// </summary>
        [Output("groupfilter")]
        public Output<string?> Groupfilter { get; private set; } = null!;

        /// <summary>
        /// Skip LDAP server SSL Certificate verification. This is not recommended for production.
        /// Defaults to `false`.
        /// </summary>
        [Output("insecureTls")]
        public Output<bool?> InsecureTls { get; private set; } = null!;

        /// <summary>
        /// The number of seconds after a Vault rotation where, if Active Directory
        /// shows a later rotation, it should be considered out-of-band
        /// </summary>
        [Output("lastRotationTolerance")]
        public Output<int> LastRotationTolerance { get; private set; } = null!;

        /// <summary>
        /// Mark the secrets engine as local-only. Local engines are not replicated or removed by
        /// replication.Tolerance duration to use when checking the last rotation time.
        /// </summary>
        [Output("local")]
        public Output<bool?> Local { get; private set; } = null!;

        /// <summary>
        /// Maximum possible lease duration for secrets in seconds.
        /// </summary>
        [Output("maxLeaseTtlSeconds")]
        public Output<int> MaxLeaseTtlSeconds { get; private set; } = null!;

        /// <summary>
        /// In seconds, the maximum password time-to-live.
        /// </summary>
        [Output("maxTtl")]
        public Output<int> MaxTtl { get; private set; } = null!;

        /// <summary>
        /// The namespace to provision the resource in.
        /// The value should not contain leading or trailing forward slashes.
        /// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
        /// *Available only for Vault Enterprise*.
        /// </summary>
        [Output("namespace")]
        public Output<string?> Namespace { get; private set; } = null!;

        /// <summary>
        /// Name of the password policy to use to generate passwords.
        /// </summary>
        [Output("passwordPolicy")]
        public Output<string?> PasswordPolicy { get; private set; } = null!;

        /// <summary>
        /// Timeout, in seconds, for the connection when making requests against the server
        /// before returning back an error.
        /// </summary>
        [Output("requestTimeout")]
        public Output<int?> RequestTimeout { get; private set; } = null!;

        /// <summary>
        /// Issue a StartTLS command after establishing unencrypted connection.
        /// </summary>
        [Output("starttls")]
        public Output<bool> Starttls { get; private set; } = null!;

        /// <summary>
        /// Maximum TLS version to use. Accepted values are `tls10`, `tls11`,
        /// `tls12` or `tls13`. Defaults to `tls12`.
        /// </summary>
        [Output("tlsMaxVersion")]
        public Output<string> TlsMaxVersion { get; private set; } = null!;

        /// <summary>
        /// Minimum TLS version to use. Accepted values are `tls10`, `tls11`,
        /// `tls12` or `tls13`. Defaults to `tls12`.
        /// </summary>
        [Output("tlsMinVersion")]
        public Output<string> TlsMinVersion { get; private set; } = null!;

        /// <summary>
        /// In seconds, the default password time-to-live.
        /// </summary>
        [Output("ttl")]
        public Output<int> Ttl { get; private set; } = null!;

        /// <summary>
        /// Enables userPrincipalDomain login with [username]@UPNDomain.
        /// </summary>
        [Output("upndomain")]
        public Output<string> Upndomain { get; private set; } = null!;

        /// <summary>
        /// LDAP URL to connect to. Multiple URLs can be specified by concatenating
        /// them with commas; they will be tried in-order. Defaults to `ldap://127.0.0.1`.
        /// </summary>
        [Output("url")]
        public Output<string?> Url { get; private set; } = null!;

        /// <summary>
        /// In Vault 1.1.1 a fix for handling group CN values of
        /// different cases unfortunately introduced a regression that could cause previously defined groups
        /// to not be found due to a change in the resulting name. If set true, the pre-1.1.1 behavior for
        /// matching group CNs will be used. This is only needed in some upgrade scenarios for backwards
        /// compatibility. It is enabled by default if the config is upgraded but disabled by default on
        /// new configurations.
        /// </summary>
        [Output("usePre111GroupCnBehavior")]
        public Output<bool> UsePre111GroupCnBehavior { get; private set; } = null!;

        /// <summary>
        /// If true, use the Active Directory tokenGroups constructed attribute of the
        /// user to find the group memberships. This will find all security groups including nested ones.
        /// </summary>
        [Output("useTokenGroups")]
        public Output<bool?> UseTokenGroups { get; private set; } = null!;

        /// <summary>
        /// Attribute used when searching users. Defaults to `cn`.
        /// </summary>
        [Output("userattr")]
        public Output<string?> Userattr { get; private set; } = null!;

        /// <summary>
        /// LDAP domain to use for users (eg: ou=People,dc=example,dc=org)`.
        /// </summary>
        [Output("userdn")]
        public Output<string?> Userdn { get; private set; } = null!;


        /// <summary>
        /// Create a SecretBackend resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public SecretBackend(string name, SecretBackendArgs args, CustomResourceOptions? options = null)
            : base("vault:ad/secretBackend:SecretBackend", name, args ?? new SecretBackendArgs(), MakeResourceOptions(options, ""))
        {
        }

        private SecretBackend(string name, Input<string> id, SecretBackendState? state = null, CustomResourceOptions? options = null)
            : base("vault:ad/secretBackend:SecretBackend", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                AdditionalSecretOutputs =
                {
                    "bindpass",
                    "clientTlsCert",
                    "clientTlsKey",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing SecretBackend resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static SecretBackend Get(string name, Input<string> id, SecretBackendState? state = null, CustomResourceOptions? options = null)
        {
            return new SecretBackend(name, id, state, options);
        }
    }

    public sealed class SecretBackendArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Use anonymous binds when performing LDAP group searches
        /// (if true the initial credentials will still be used for the initial connection test).
        /// </summary>
        [Input("anonymousGroupSearch")]
        public Input<bool>? AnonymousGroupSearch { get; set; }

        /// <summary>
        /// The unique path this backend should be mounted at. Must
        /// not begin or end with a `/`. Defaults to `ad`.
        /// </summary>
        [Input("backend")]
        public Input<string>? Backend { get; set; }

        /// <summary>
        /// Distinguished name of object to bind when performing user and group search.
        /// </summary>
        [Input("binddn", required: true)]
        public Input<string> Binddn { get; set; } = null!;

        [Input("bindpass", required: true)]
        private Input<string>? _bindpass;

        /// <summary>
        /// Password to use along with binddn when performing user search.
        /// </summary>
        public Input<string>? Bindpass
        {
            get => _bindpass;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _bindpass = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// If set, user and group names assigned to policies within the
        /// backend will be case sensitive. Otherwise, names will be normalized to lower case.
        /// </summary>
        [Input("caseSensitiveNames")]
        public Input<bool>? CaseSensitiveNames { get; set; }

        /// <summary>
        /// CA certificate to use when verifying LDAP server certificate, must be
        /// x509 PEM encoded.
        /// </summary>
        [Input("certificate")]
        public Input<string>? Certificate { get; set; }

        [Input("clientTlsCert")]
        private Input<string>? _clientTlsCert;

        /// <summary>
        /// Client certificate to provide to the LDAP server, must be x509 PEM encoded.
        /// </summary>
        public Input<string>? ClientTlsCert
        {
            get => _clientTlsCert;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _clientTlsCert = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        [Input("clientTlsKey")]
        private Input<string>? _clientTlsKey;

        /// <summary>
        /// Client certificate key to provide to the LDAP server, must be x509 PEM encoded.
        /// </summary>
        public Input<string>? ClientTlsKey
        {
            get => _clientTlsKey;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _clientTlsKey = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// Default lease duration for secrets in seconds.
        /// </summary>
        [Input("defaultLeaseTtlSeconds")]
        public Input<int>? DefaultLeaseTtlSeconds { get; set; }

        /// <summary>
        /// Denies an unauthenticated LDAP bind request if the user's password is empty;
        /// defaults to true.
        /// </summary>
        [Input("denyNullBind")]
        public Input<bool>? DenyNullBind { get; set; }

        /// <summary>
        /// Human-friendly description of the mount for the Active Directory backend.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// If set, opts out of mount migration on path updates.
        /// See here for more info on [Mount Migration](https://www.vaultproject.io/docs/concepts/mount-migration)
        /// </summary>
        [Input("disableRemount")]
        public Input<bool>? DisableRemount { get; set; }

        /// <summary>
        /// Use anonymous bind to discover the bind Distinguished Name of a user.
        /// </summary>
        [Input("discoverdn")]
        public Input<bool>? Discoverdn { get; set; }

        /// <summary>
        /// LDAP attribute to follow on objects returned by &lt;groupfilter&gt; in order to enumerate
        /// user group membership. Examples: `cn` or `memberOf`, etc. Defaults to `cn`.
        /// </summary>
        [Input("groupattr")]
        public Input<string>? Groupattr { get; set; }

        /// <summary>
        /// LDAP search base to use for group membership search (eg: ou=Groups,dc=example,dc=org).
        /// </summary>
        [Input("groupdn")]
        public Input<string>? Groupdn { get; set; }

        /// <summary>
        /// Go template for querying group membership of user (optional) The template can access
        /// the following context variables: UserDN, Username. Defaults to `(|(memberUid={{.Username}})(member={{.UserDN}})(uniqueMember={{.UserDN}}))`
        /// </summary>
        [Input("groupfilter")]
        public Input<string>? Groupfilter { get; set; }

        /// <summary>
        /// Skip LDAP server SSL Certificate verification. This is not recommended for production.
        /// Defaults to `false`.
        /// </summary>
        [Input("insecureTls")]
        public Input<bool>? InsecureTls { get; set; }

        /// <summary>
        /// The number of seconds after a Vault rotation where, if Active Directory
        /// shows a later rotation, it should be considered out-of-band
        /// </summary>
        [Input("lastRotationTolerance")]
        public Input<int>? LastRotationTolerance { get; set; }

        /// <summary>
        /// Mark the secrets engine as local-only. Local engines are not replicated or removed by
        /// replication.Tolerance duration to use when checking the last rotation time.
        /// </summary>
        [Input("local")]
        public Input<bool>? Local { get; set; }

        /// <summary>
        /// Maximum possible lease duration for secrets in seconds.
        /// </summary>
        [Input("maxLeaseTtlSeconds")]
        public Input<int>? MaxLeaseTtlSeconds { get; set; }

        /// <summary>
        /// In seconds, the maximum password time-to-live.
        /// </summary>
        [Input("maxTtl")]
        public Input<int>? MaxTtl { get; set; }

        /// <summary>
        /// The namespace to provision the resource in.
        /// The value should not contain leading or trailing forward slashes.
        /// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
        /// *Available only for Vault Enterprise*.
        /// </summary>
        [Input("namespace")]
        public Input<string>? Namespace { get; set; }

        /// <summary>
        /// Name of the password policy to use to generate passwords.
        /// </summary>
        [Input("passwordPolicy")]
        public Input<string>? PasswordPolicy { get; set; }

        /// <summary>
        /// Timeout, in seconds, for the connection when making requests against the server
        /// before returning back an error.
        /// </summary>
        [Input("requestTimeout")]
        public Input<int>? RequestTimeout { get; set; }

        /// <summary>
        /// Issue a StartTLS command after establishing unencrypted connection.
        /// </summary>
        [Input("starttls")]
        public Input<bool>? Starttls { get; set; }

        /// <summary>
        /// Maximum TLS version to use. Accepted values are `tls10`, `tls11`,
        /// `tls12` or `tls13`. Defaults to `tls12`.
        /// </summary>
        [Input("tlsMaxVersion")]
        public Input<string>? TlsMaxVersion { get; set; }

        /// <summary>
        /// Minimum TLS version to use. Accepted values are `tls10`, `tls11`,
        /// `tls12` or `tls13`. Defaults to `tls12`.
        /// </summary>
        [Input("tlsMinVersion")]
        public Input<string>? TlsMinVersion { get; set; }

        /// <summary>
        /// In seconds, the default password time-to-live.
        /// </summary>
        [Input("ttl")]
        public Input<int>? Ttl { get; set; }

        /// <summary>
        /// Enables userPrincipalDomain login with [username]@UPNDomain.
        /// </summary>
        [Input("upndomain")]
        public Input<string>? Upndomain { get; set; }

        /// <summary>
        /// LDAP URL to connect to. Multiple URLs can be specified by concatenating
        /// them with commas; they will be tried in-order. Defaults to `ldap://127.0.0.1`.
        /// </summary>
        [Input("url")]
        public Input<string>? Url { get; set; }

        /// <summary>
        /// In Vault 1.1.1 a fix for handling group CN values of
        /// different cases unfortunately introduced a regression that could cause previously defined groups
        /// to not be found due to a change in the resulting name. If set true, the pre-1.1.1 behavior for
        /// matching group CNs will be used. This is only needed in some upgrade scenarios for backwards
        /// compatibility. It is enabled by default if the config is upgraded but disabled by default on
        /// new configurations.
        /// </summary>
        [Input("usePre111GroupCnBehavior")]
        public Input<bool>? UsePre111GroupCnBehavior { get; set; }

        /// <summary>
        /// If true, use the Active Directory tokenGroups constructed attribute of the
        /// user to find the group memberships. This will find all security groups including nested ones.
        /// </summary>
        [Input("useTokenGroups")]
        public Input<bool>? UseTokenGroups { get; set; }

        /// <summary>
        /// Attribute used when searching users. Defaults to `cn`.
        /// </summary>
        [Input("userattr")]
        public Input<string>? Userattr { get; set; }

        /// <summary>
        /// LDAP domain to use for users (eg: ou=People,dc=example,dc=org)`.
        /// </summary>
        [Input("userdn")]
        public Input<string>? Userdn { get; set; }

        public SecretBackendArgs()
        {
        }
        public static new SecretBackendArgs Empty => new SecretBackendArgs();
    }

    public sealed class SecretBackendState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Use anonymous binds when performing LDAP group searches
        /// (if true the initial credentials will still be used for the initial connection test).
        /// </summary>
        [Input("anonymousGroupSearch")]
        public Input<bool>? AnonymousGroupSearch { get; set; }

        /// <summary>
        /// The unique path this backend should be mounted at. Must
        /// not begin or end with a `/`. Defaults to `ad`.
        /// </summary>
        [Input("backend")]
        public Input<string>? Backend { get; set; }

        /// <summary>
        /// Distinguished name of object to bind when performing user and group search.
        /// </summary>
        [Input("binddn")]
        public Input<string>? Binddn { get; set; }

        [Input("bindpass")]
        private Input<string>? _bindpass;

        /// <summary>
        /// Password to use along with binddn when performing user search.
        /// </summary>
        public Input<string>? Bindpass
        {
            get => _bindpass;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _bindpass = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// If set, user and group names assigned to policies within the
        /// backend will be case sensitive. Otherwise, names will be normalized to lower case.
        /// </summary>
        [Input("caseSensitiveNames")]
        public Input<bool>? CaseSensitiveNames { get; set; }

        /// <summary>
        /// CA certificate to use when verifying LDAP server certificate, must be
        /// x509 PEM encoded.
        /// </summary>
        [Input("certificate")]
        public Input<string>? Certificate { get; set; }

        [Input("clientTlsCert")]
        private Input<string>? _clientTlsCert;

        /// <summary>
        /// Client certificate to provide to the LDAP server, must be x509 PEM encoded.
        /// </summary>
        public Input<string>? ClientTlsCert
        {
            get => _clientTlsCert;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _clientTlsCert = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        [Input("clientTlsKey")]
        private Input<string>? _clientTlsKey;

        /// <summary>
        /// Client certificate key to provide to the LDAP server, must be x509 PEM encoded.
        /// </summary>
        public Input<string>? ClientTlsKey
        {
            get => _clientTlsKey;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _clientTlsKey = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// Default lease duration for secrets in seconds.
        /// </summary>
        [Input("defaultLeaseTtlSeconds")]
        public Input<int>? DefaultLeaseTtlSeconds { get; set; }

        /// <summary>
        /// Denies an unauthenticated LDAP bind request if the user's password is empty;
        /// defaults to true.
        /// </summary>
        [Input("denyNullBind")]
        public Input<bool>? DenyNullBind { get; set; }

        /// <summary>
        /// Human-friendly description of the mount for the Active Directory backend.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// If set, opts out of mount migration on path updates.
        /// See here for more info on [Mount Migration](https://www.vaultproject.io/docs/concepts/mount-migration)
        /// </summary>
        [Input("disableRemount")]
        public Input<bool>? DisableRemount { get; set; }

        /// <summary>
        /// Use anonymous bind to discover the bind Distinguished Name of a user.
        /// </summary>
        [Input("discoverdn")]
        public Input<bool>? Discoverdn { get; set; }

        /// <summary>
        /// LDAP attribute to follow on objects returned by &lt;groupfilter&gt; in order to enumerate
        /// user group membership. Examples: `cn` or `memberOf`, etc. Defaults to `cn`.
        /// </summary>
        [Input("groupattr")]
        public Input<string>? Groupattr { get; set; }

        /// <summary>
        /// LDAP search base to use for group membership search (eg: ou=Groups,dc=example,dc=org).
        /// </summary>
        [Input("groupdn")]
        public Input<string>? Groupdn { get; set; }

        /// <summary>
        /// Go template for querying group membership of user (optional) The template can access
        /// the following context variables: UserDN, Username. Defaults to `(|(memberUid={{.Username}})(member={{.UserDN}})(uniqueMember={{.UserDN}}))`
        /// </summary>
        [Input("groupfilter")]
        public Input<string>? Groupfilter { get; set; }

        /// <summary>
        /// Skip LDAP server SSL Certificate verification. This is not recommended for production.
        /// Defaults to `false`.
        /// </summary>
        [Input("insecureTls")]
        public Input<bool>? InsecureTls { get; set; }

        /// <summary>
        /// The number of seconds after a Vault rotation where, if Active Directory
        /// shows a later rotation, it should be considered out-of-band
        /// </summary>
        [Input("lastRotationTolerance")]
        public Input<int>? LastRotationTolerance { get; set; }

        /// <summary>
        /// Mark the secrets engine as local-only. Local engines are not replicated or removed by
        /// replication.Tolerance duration to use when checking the last rotation time.
        /// </summary>
        [Input("local")]
        public Input<bool>? Local { get; set; }

        /// <summary>
        /// Maximum possible lease duration for secrets in seconds.
        /// </summary>
        [Input("maxLeaseTtlSeconds")]
        public Input<int>? MaxLeaseTtlSeconds { get; set; }

        /// <summary>
        /// In seconds, the maximum password time-to-live.
        /// </summary>
        [Input("maxTtl")]
        public Input<int>? MaxTtl { get; set; }

        /// <summary>
        /// The namespace to provision the resource in.
        /// The value should not contain leading or trailing forward slashes.
        /// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
        /// *Available only for Vault Enterprise*.
        /// </summary>
        [Input("namespace")]
        public Input<string>? Namespace { get; set; }

        /// <summary>
        /// Name of the password policy to use to generate passwords.
        /// </summary>
        [Input("passwordPolicy")]
        public Input<string>? PasswordPolicy { get; set; }

        /// <summary>
        /// Timeout, in seconds, for the connection when making requests against the server
        /// before returning back an error.
        /// </summary>
        [Input("requestTimeout")]
        public Input<int>? RequestTimeout { get; set; }

        /// <summary>
        /// Issue a StartTLS command after establishing unencrypted connection.
        /// </summary>
        [Input("starttls")]
        public Input<bool>? Starttls { get; set; }

        /// <summary>
        /// Maximum TLS version to use. Accepted values are `tls10`, `tls11`,
        /// `tls12` or `tls13`. Defaults to `tls12`.
        /// </summary>
        [Input("tlsMaxVersion")]
        public Input<string>? TlsMaxVersion { get; set; }

        /// <summary>
        /// Minimum TLS version to use. Accepted values are `tls10`, `tls11`,
        /// `tls12` or `tls13`. Defaults to `tls12`.
        /// </summary>
        [Input("tlsMinVersion")]
        public Input<string>? TlsMinVersion { get; set; }

        /// <summary>
        /// In seconds, the default password time-to-live.
        /// </summary>
        [Input("ttl")]
        public Input<int>? Ttl { get; set; }

        /// <summary>
        /// Enables userPrincipalDomain login with [username]@UPNDomain.
        /// </summary>
        [Input("upndomain")]
        public Input<string>? Upndomain { get; set; }

        /// <summary>
        /// LDAP URL to connect to. Multiple URLs can be specified by concatenating
        /// them with commas; they will be tried in-order. Defaults to `ldap://127.0.0.1`.
        /// </summary>
        [Input("url")]
        public Input<string>? Url { get; set; }

        /// <summary>
        /// In Vault 1.1.1 a fix for handling group CN values of
        /// different cases unfortunately introduced a regression that could cause previously defined groups
        /// to not be found due to a change in the resulting name. If set true, the pre-1.1.1 behavior for
        /// matching group CNs will be used. This is only needed in some upgrade scenarios for backwards
        /// compatibility. It is enabled by default if the config is upgraded but disabled by default on
        /// new configurations.
        /// </summary>
        [Input("usePre111GroupCnBehavior")]
        public Input<bool>? UsePre111GroupCnBehavior { get; set; }

        /// <summary>
        /// If true, use the Active Directory tokenGroups constructed attribute of the
        /// user to find the group memberships. This will find all security groups including nested ones.
        /// </summary>
        [Input("useTokenGroups")]
        public Input<bool>? UseTokenGroups { get; set; }

        /// <summary>
        /// Attribute used when searching users. Defaults to `cn`.
        /// </summary>
        [Input("userattr")]
        public Input<string>? Userattr { get; set; }

        /// <summary>
        /// LDAP domain to use for users (eg: ou=People,dc=example,dc=org)`.
        /// </summary>
        [Input("userdn")]
        public Input<string>? Userdn { get; set; }

        public SecretBackendState()
        {
        }
        public static new SecretBackendState Empty => new SecretBackendState();
    }
}
