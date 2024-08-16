// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vault.Ldap
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
    ///     var config = new Vault.Ldap.SecretBackend("config", new()
    ///     {
    ///         Path = "my-custom-ldap",
    ///         Binddn = "CN=Administrator,CN=Users,DC=corp,DC=example,DC=net",
    ///         Bindpass = "SuperSecretPassw0rd",
    ///         Url = "ldaps://localhost",
    ///         InsecureTls = true,
    ///         Userdn = "CN=Users,DC=corp,DC=example,DC=net",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// LDAP secret backend can be imported using the `${mount}/config`, e.g.
    /// 
    /// ```sh
    /// $ pulumi import vault:ldap/secretBackend:SecretBackend config ldap/config
    /// ```
    /// </summary>
    [VaultResourceType("vault:ldap/secretBackend:SecretBackend")]
    public partial class SecretBackend : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Accessor of the mount
        /// </summary>
        [Output("accessor")]
        public Output<string> Accessor { get; private set; } = null!;

        /// <summary>
        /// List of managed key registry entry names that the mount in question is allowed to access
        /// </summary>
        [Output("allowedManagedKeys")]
        public Output<ImmutableArray<string>> AllowedManagedKeys { get; private set; } = null!;

        /// <summary>
        /// List of headers to allow and pass from the request to the plugin
        /// </summary>
        [Output("allowedResponseHeaders")]
        public Output<ImmutableArray<string>> AllowedResponseHeaders { get; private set; } = null!;

        /// <summary>
        /// Specifies the list of keys that will not be HMAC'd by audit devices in the request data object.
        /// </summary>
        [Output("auditNonHmacRequestKeys")]
        public Output<ImmutableArray<string>> AuditNonHmacRequestKeys { get; private set; } = null!;

        /// <summary>
        /// Specifies the list of keys that will not be HMAC'd by audit devices in the response data object.
        /// </summary>
        [Output("auditNonHmacResponseKeys")]
        public Output<ImmutableArray<string>> AuditNonHmacResponseKeys { get; private set; } = null!;

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
        /// Timeout, in seconds, when attempting to connect to the LDAP server before trying
        /// the next URL in the configuration.
        /// </summary>
        [Output("connectionTimeout")]
        public Output<int?> ConnectionTimeout { get; private set; } = null!;

        /// <summary>
        /// Default lease duration for secrets in seconds.
        /// </summary>
        [Output("defaultLeaseTtlSeconds")]
        public Output<int> DefaultLeaseTtlSeconds { get; private set; } = null!;

        /// <summary>
        /// List of headers to allow and pass from the request to the plugin
        /// </summary>
        [Output("delegatedAuthAccessors")]
        public Output<ImmutableArray<string>> DelegatedAuthAccessors { get; private set; } = null!;

        /// <summary>
        /// Human-friendly description of the mount for the Active Directory backend.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// If set, opts out of mount migration on path updates.
        /// </summary>
        [Output("disableRemount")]
        public Output<bool?> DisableRemount { get; private set; } = null!;

        /// <summary>
        /// Enable the secrets engine to access Vault's external entropy source
        /// </summary>
        [Output("externalEntropyAccess")]
        public Output<bool?> ExternalEntropyAccess { get; private set; } = null!;

        /// <summary>
        /// The key to use for signing plugin workload identity tokens
        /// </summary>
        [Output("identityTokenKey")]
        public Output<string?> IdentityTokenKey { get; private set; } = null!;

        /// <summary>
        /// Skip LDAP server SSL Certificate verification. This is not recommended for production.
        /// Defaults to `false`.
        /// </summary>
        [Output("insecureTls")]
        public Output<bool?> InsecureTls { get; private set; } = null!;

        /// <summary>
        /// Specifies whether to show this mount in the UI-specific listing endpoint
        /// </summary>
        [Output("listingVisibility")]
        public Output<string?> ListingVisibility { get; private set; } = null!;

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
        /// The namespace to provision the resource in.
        /// The value should not contain leading or trailing forward slashes.
        /// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
        /// *Available only for Vault Enterprise*.
        /// </summary>
        [Output("namespace")]
        public Output<string?> Namespace { get; private set; } = null!;

        /// <summary>
        /// Specifies mount type specific options that are passed to the backend
        /// </summary>
        [Output("options")]
        public Output<ImmutableDictionary<string, string>?> Options { get; private set; } = null!;

        /// <summary>
        /// List of headers to allow and pass from the request to the plugin
        /// </summary>
        [Output("passthroughRequestHeaders")]
        public Output<ImmutableArray<string>> PassthroughRequestHeaders { get; private set; } = null!;

        /// <summary>
        /// Name of the password policy to use to generate passwords.
        /// </summary>
        [Output("passwordPolicy")]
        public Output<string?> PasswordPolicy { get; private set; } = null!;

        /// <summary>
        /// The unique path this backend should be mounted at. Must
        /// not begin or end with a `/`. Defaults to `ldap`.
        /// </summary>
        [Output("path")]
        public Output<string?> Path { get; private set; } = null!;

        /// <summary>
        /// Specifies the semantic version of the plugin to use, e.g. 'v1.0.0'
        /// </summary>
        [Output("pluginVersion")]
        public Output<string?> PluginVersion { get; private set; } = null!;

        /// <summary>
        /// Timeout, in seconds, for the connection when making requests against the server
        /// before returning back an error.
        /// </summary>
        [Output("requestTimeout")]
        public Output<int> RequestTimeout { get; private set; } = null!;

        /// <summary>
        /// The LDAP schema to use when storing entry passwords. Valid schemas include `openldap`, `ad`, and `racf`. Default is `openldap`.
        /// </summary>
        [Output("schema")]
        public Output<string> Schema { get; private set; } = null!;

        /// <summary>
        /// Enable seal wrapping for the mount, causing values stored by the mount to be wrapped by the seal's encryption capability
        /// </summary>
        [Output("sealWrap")]
        public Output<bool> SealWrap { get; private set; } = null!;

        /// <summary>
        /// If set to true, static roles will not be rotated during import.
        /// Defaults to false. Requires Vault 1.16 or above.
        /// </summary>
        [Output("skipStaticRoleImportRotation")]
        public Output<bool?> SkipStaticRoleImportRotation { get; private set; } = null!;

        /// <summary>
        /// Issue a StartTLS command after establishing unencrypted connection.
        /// </summary>
        [Output("starttls")]
        public Output<bool> Starttls { get; private set; } = null!;

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
        public Output<string> Url { get; private set; } = null!;

        /// <summary>
        /// Attribute used when searching users. Defaults to `cn`.
        /// </summary>
        [Output("userattr")]
        public Output<string> Userattr { get; private set; } = null!;

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
            : base("vault:ldap/secretBackend:SecretBackend", name, args ?? new SecretBackendArgs(), MakeResourceOptions(options, ""))
        {
        }

        private SecretBackend(string name, Input<string> id, SecretBackendState? state = null, CustomResourceOptions? options = null)
            : base("vault:ldap/secretBackend:SecretBackend", name, state, MakeResourceOptions(options, id))
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
        [Input("allowedManagedKeys")]
        private InputList<string>? _allowedManagedKeys;

        /// <summary>
        /// List of managed key registry entry names that the mount in question is allowed to access
        /// </summary>
        public InputList<string> AllowedManagedKeys
        {
            get => _allowedManagedKeys ?? (_allowedManagedKeys = new InputList<string>());
            set => _allowedManagedKeys = value;
        }

        [Input("allowedResponseHeaders")]
        private InputList<string>? _allowedResponseHeaders;

        /// <summary>
        /// List of headers to allow and pass from the request to the plugin
        /// </summary>
        public InputList<string> AllowedResponseHeaders
        {
            get => _allowedResponseHeaders ?? (_allowedResponseHeaders = new InputList<string>());
            set => _allowedResponseHeaders = value;
        }

        [Input("auditNonHmacRequestKeys")]
        private InputList<string>? _auditNonHmacRequestKeys;

        /// <summary>
        /// Specifies the list of keys that will not be HMAC'd by audit devices in the request data object.
        /// </summary>
        public InputList<string> AuditNonHmacRequestKeys
        {
            get => _auditNonHmacRequestKeys ?? (_auditNonHmacRequestKeys = new InputList<string>());
            set => _auditNonHmacRequestKeys = value;
        }

        [Input("auditNonHmacResponseKeys")]
        private InputList<string>? _auditNonHmacResponseKeys;

        /// <summary>
        /// Specifies the list of keys that will not be HMAC'd by audit devices in the response data object.
        /// </summary>
        public InputList<string> AuditNonHmacResponseKeys
        {
            get => _auditNonHmacResponseKeys ?? (_auditNonHmacResponseKeys = new InputList<string>());
            set => _auditNonHmacResponseKeys = value;
        }

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
        /// Timeout, in seconds, when attempting to connect to the LDAP server before trying
        /// the next URL in the configuration.
        /// </summary>
        [Input("connectionTimeout")]
        public Input<int>? ConnectionTimeout { get; set; }

        /// <summary>
        /// Default lease duration for secrets in seconds.
        /// </summary>
        [Input("defaultLeaseTtlSeconds")]
        public Input<int>? DefaultLeaseTtlSeconds { get; set; }

        [Input("delegatedAuthAccessors")]
        private InputList<string>? _delegatedAuthAccessors;

        /// <summary>
        /// List of headers to allow and pass from the request to the plugin
        /// </summary>
        public InputList<string> DelegatedAuthAccessors
        {
            get => _delegatedAuthAccessors ?? (_delegatedAuthAccessors = new InputList<string>());
            set => _delegatedAuthAccessors = value;
        }

        /// <summary>
        /// Human-friendly description of the mount for the Active Directory backend.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// If set, opts out of mount migration on path updates.
        /// </summary>
        [Input("disableRemount")]
        public Input<bool>? DisableRemount { get; set; }

        /// <summary>
        /// Enable the secrets engine to access Vault's external entropy source
        /// </summary>
        [Input("externalEntropyAccess")]
        public Input<bool>? ExternalEntropyAccess { get; set; }

        /// <summary>
        /// The key to use for signing plugin workload identity tokens
        /// </summary>
        [Input("identityTokenKey")]
        public Input<string>? IdentityTokenKey { get; set; }

        /// <summary>
        /// Skip LDAP server SSL Certificate verification. This is not recommended for production.
        /// Defaults to `false`.
        /// </summary>
        [Input("insecureTls")]
        public Input<bool>? InsecureTls { get; set; }

        /// <summary>
        /// Specifies whether to show this mount in the UI-specific listing endpoint
        /// </summary>
        [Input("listingVisibility")]
        public Input<string>? ListingVisibility { get; set; }

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
        /// The namespace to provision the resource in.
        /// The value should not contain leading or trailing forward slashes.
        /// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
        /// *Available only for Vault Enterprise*.
        /// </summary>
        [Input("namespace")]
        public Input<string>? Namespace { get; set; }

        [Input("options")]
        private InputMap<string>? _options;

        /// <summary>
        /// Specifies mount type specific options that are passed to the backend
        /// </summary>
        public InputMap<string> Options
        {
            get => _options ?? (_options = new InputMap<string>());
            set => _options = value;
        }

        [Input("passthroughRequestHeaders")]
        private InputList<string>? _passthroughRequestHeaders;

        /// <summary>
        /// List of headers to allow and pass from the request to the plugin
        /// </summary>
        public InputList<string> PassthroughRequestHeaders
        {
            get => _passthroughRequestHeaders ?? (_passthroughRequestHeaders = new InputList<string>());
            set => _passthroughRequestHeaders = value;
        }

        /// <summary>
        /// Name of the password policy to use to generate passwords.
        /// </summary>
        [Input("passwordPolicy")]
        public Input<string>? PasswordPolicy { get; set; }

        /// <summary>
        /// The unique path this backend should be mounted at. Must
        /// not begin or end with a `/`. Defaults to `ldap`.
        /// </summary>
        [Input("path")]
        public Input<string>? Path { get; set; }

        /// <summary>
        /// Specifies the semantic version of the plugin to use, e.g. 'v1.0.0'
        /// </summary>
        [Input("pluginVersion")]
        public Input<string>? PluginVersion { get; set; }

        /// <summary>
        /// Timeout, in seconds, for the connection when making requests against the server
        /// before returning back an error.
        /// </summary>
        [Input("requestTimeout")]
        public Input<int>? RequestTimeout { get; set; }

        /// <summary>
        /// The LDAP schema to use when storing entry passwords. Valid schemas include `openldap`, `ad`, and `racf`. Default is `openldap`.
        /// </summary>
        [Input("schema")]
        public Input<string>? Schema { get; set; }

        /// <summary>
        /// Enable seal wrapping for the mount, causing values stored by the mount to be wrapped by the seal's encryption capability
        /// </summary>
        [Input("sealWrap")]
        public Input<bool>? SealWrap { get; set; }

        /// <summary>
        /// If set to true, static roles will not be rotated during import.
        /// Defaults to false. Requires Vault 1.16 or above.
        /// </summary>
        [Input("skipStaticRoleImportRotation")]
        public Input<bool>? SkipStaticRoleImportRotation { get; set; }

        /// <summary>
        /// Issue a StartTLS command after establishing unencrypted connection.
        /// </summary>
        [Input("starttls")]
        public Input<bool>? Starttls { get; set; }

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
        /// Accessor of the mount
        /// </summary>
        [Input("accessor")]
        public Input<string>? Accessor { get; set; }

        [Input("allowedManagedKeys")]
        private InputList<string>? _allowedManagedKeys;

        /// <summary>
        /// List of managed key registry entry names that the mount in question is allowed to access
        /// </summary>
        public InputList<string> AllowedManagedKeys
        {
            get => _allowedManagedKeys ?? (_allowedManagedKeys = new InputList<string>());
            set => _allowedManagedKeys = value;
        }

        [Input("allowedResponseHeaders")]
        private InputList<string>? _allowedResponseHeaders;

        /// <summary>
        /// List of headers to allow and pass from the request to the plugin
        /// </summary>
        public InputList<string> AllowedResponseHeaders
        {
            get => _allowedResponseHeaders ?? (_allowedResponseHeaders = new InputList<string>());
            set => _allowedResponseHeaders = value;
        }

        [Input("auditNonHmacRequestKeys")]
        private InputList<string>? _auditNonHmacRequestKeys;

        /// <summary>
        /// Specifies the list of keys that will not be HMAC'd by audit devices in the request data object.
        /// </summary>
        public InputList<string> AuditNonHmacRequestKeys
        {
            get => _auditNonHmacRequestKeys ?? (_auditNonHmacRequestKeys = new InputList<string>());
            set => _auditNonHmacRequestKeys = value;
        }

        [Input("auditNonHmacResponseKeys")]
        private InputList<string>? _auditNonHmacResponseKeys;

        /// <summary>
        /// Specifies the list of keys that will not be HMAC'd by audit devices in the response data object.
        /// </summary>
        public InputList<string> AuditNonHmacResponseKeys
        {
            get => _auditNonHmacResponseKeys ?? (_auditNonHmacResponseKeys = new InputList<string>());
            set => _auditNonHmacResponseKeys = value;
        }

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
        /// Timeout, in seconds, when attempting to connect to the LDAP server before trying
        /// the next URL in the configuration.
        /// </summary>
        [Input("connectionTimeout")]
        public Input<int>? ConnectionTimeout { get; set; }

        /// <summary>
        /// Default lease duration for secrets in seconds.
        /// </summary>
        [Input("defaultLeaseTtlSeconds")]
        public Input<int>? DefaultLeaseTtlSeconds { get; set; }

        [Input("delegatedAuthAccessors")]
        private InputList<string>? _delegatedAuthAccessors;

        /// <summary>
        /// List of headers to allow and pass from the request to the plugin
        /// </summary>
        public InputList<string> DelegatedAuthAccessors
        {
            get => _delegatedAuthAccessors ?? (_delegatedAuthAccessors = new InputList<string>());
            set => _delegatedAuthAccessors = value;
        }

        /// <summary>
        /// Human-friendly description of the mount for the Active Directory backend.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// If set, opts out of mount migration on path updates.
        /// </summary>
        [Input("disableRemount")]
        public Input<bool>? DisableRemount { get; set; }

        /// <summary>
        /// Enable the secrets engine to access Vault's external entropy source
        /// </summary>
        [Input("externalEntropyAccess")]
        public Input<bool>? ExternalEntropyAccess { get; set; }

        /// <summary>
        /// The key to use for signing plugin workload identity tokens
        /// </summary>
        [Input("identityTokenKey")]
        public Input<string>? IdentityTokenKey { get; set; }

        /// <summary>
        /// Skip LDAP server SSL Certificate verification. This is not recommended for production.
        /// Defaults to `false`.
        /// </summary>
        [Input("insecureTls")]
        public Input<bool>? InsecureTls { get; set; }

        /// <summary>
        /// Specifies whether to show this mount in the UI-specific listing endpoint
        /// </summary>
        [Input("listingVisibility")]
        public Input<string>? ListingVisibility { get; set; }

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
        /// The namespace to provision the resource in.
        /// The value should not contain leading or trailing forward slashes.
        /// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
        /// *Available only for Vault Enterprise*.
        /// </summary>
        [Input("namespace")]
        public Input<string>? Namespace { get; set; }

        [Input("options")]
        private InputMap<string>? _options;

        /// <summary>
        /// Specifies mount type specific options that are passed to the backend
        /// </summary>
        public InputMap<string> Options
        {
            get => _options ?? (_options = new InputMap<string>());
            set => _options = value;
        }

        [Input("passthroughRequestHeaders")]
        private InputList<string>? _passthroughRequestHeaders;

        /// <summary>
        /// List of headers to allow and pass from the request to the plugin
        /// </summary>
        public InputList<string> PassthroughRequestHeaders
        {
            get => _passthroughRequestHeaders ?? (_passthroughRequestHeaders = new InputList<string>());
            set => _passthroughRequestHeaders = value;
        }

        /// <summary>
        /// Name of the password policy to use to generate passwords.
        /// </summary>
        [Input("passwordPolicy")]
        public Input<string>? PasswordPolicy { get; set; }

        /// <summary>
        /// The unique path this backend should be mounted at. Must
        /// not begin or end with a `/`. Defaults to `ldap`.
        /// </summary>
        [Input("path")]
        public Input<string>? Path { get; set; }

        /// <summary>
        /// Specifies the semantic version of the plugin to use, e.g. 'v1.0.0'
        /// </summary>
        [Input("pluginVersion")]
        public Input<string>? PluginVersion { get; set; }

        /// <summary>
        /// Timeout, in seconds, for the connection when making requests against the server
        /// before returning back an error.
        /// </summary>
        [Input("requestTimeout")]
        public Input<int>? RequestTimeout { get; set; }

        /// <summary>
        /// The LDAP schema to use when storing entry passwords. Valid schemas include `openldap`, `ad`, and `racf`. Default is `openldap`.
        /// </summary>
        [Input("schema")]
        public Input<string>? Schema { get; set; }

        /// <summary>
        /// Enable seal wrapping for the mount, causing values stored by the mount to be wrapped by the seal's encryption capability
        /// </summary>
        [Input("sealWrap")]
        public Input<bool>? SealWrap { get; set; }

        /// <summary>
        /// If set to true, static roles will not be rotated during import.
        /// Defaults to false. Requires Vault 1.16 or above.
        /// </summary>
        [Input("skipStaticRoleImportRotation")]
        public Input<bool>? SkipStaticRoleImportRotation { get; set; }

        /// <summary>
        /// Issue a StartTLS command after establishing unencrypted connection.
        /// </summary>
        [Input("starttls")]
        public Input<bool>? Starttls { get; set; }

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
