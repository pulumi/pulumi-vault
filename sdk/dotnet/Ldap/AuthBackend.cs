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
    /// Provides a resource for managing an [LDAP auth backend within Vault](https://www.vaultproject.io/docs/auth/ldap.html).
    /// 
    /// ## Example Usage
    /// 
    /// &lt;!--Start PulumiCodeChooser --&gt;
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Vault = Pulumi.Vault;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var ldap = new Vault.Ldap.AuthBackend("ldap", new()
    ///     {
    ///         Discoverdn = false,
    ///         Groupdn = "OU=Groups,DC=example,DC=org",
    ///         Groupfilter = "(&amp;(objectClass=group)(member:1.2.840.113556.1.4.1941:={{.UserDN}}))",
    ///         Path = "ldap",
    ///         Upndomain = "EXAMPLE.ORG",
    ///         Url = "ldaps://dc-01.example.org",
    ///         Userattr = "sAMAccountName",
    ///         Userdn = "OU=Users,OU=Accounts,DC=example,DC=org",
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ## Import
    /// 
    /// LDAP authentication backends can be imported using the `path`, e.g.
    /// 
    /// ```sh
    /// $ pulumi import vault:ldap/authBackend:AuthBackend ldap ldap
    /// ```
    /// </summary>
    [VaultResourceType("vault:ldap/authBackend:AuthBackend")]
    public partial class AuthBackend : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The accessor for this auth mount.
        /// </summary>
        [Output("accessor")]
        public Output<string> Accessor { get; private set; } = null!;

        /// <summary>
        /// DN of object to bind when performing user search
        /// </summary>
        [Output("binddn")]
        public Output<string> Binddn { get; private set; } = null!;

        /// <summary>
        /// Password to use with `binddn` when performing user search
        /// </summary>
        [Output("bindpass")]
        public Output<string> Bindpass { get; private set; } = null!;

        /// <summary>
        /// Control case senstivity of objects fetched from LDAP, this is used for object matching in vault
        /// </summary>
        [Output("caseSensitiveNames")]
        public Output<bool> CaseSensitiveNames { get; private set; } = null!;

        /// <summary>
        /// Trusted CA to validate TLS certificate
        /// </summary>
        [Output("certificate")]
        public Output<string> Certificate { get; private set; } = null!;

        [Output("clientTlsCert")]
        public Output<string> ClientTlsCert { get; private set; } = null!;

        [Output("clientTlsKey")]
        public Output<string> ClientTlsKey { get; private set; } = null!;

        /// <summary>
        /// Prevents users from bypassing authentication when providing an empty password.
        /// </summary>
        [Output("denyNullBind")]
        public Output<bool> DenyNullBind { get; private set; } = null!;

        /// <summary>
        /// Description for the LDAP auth backend mount
        /// </summary>
        [Output("description")]
        public Output<string> Description { get; private set; } = null!;

        /// <summary>
        /// If set, opts out of mount migration on path updates.
        /// See here for more info on [Mount Migration](https://www.vaultproject.io/docs/concepts/mount-migration)
        /// </summary>
        [Output("disableRemount")]
        public Output<bool?> DisableRemount { get; private set; } = null!;

        /// <summary>
        /// Use anonymous bind to discover the bind DN of a user.
        /// </summary>
        [Output("discoverdn")]
        public Output<bool> Discoverdn { get; private set; } = null!;

        /// <summary>
        /// LDAP attribute to follow on objects returned by groupfilter
        /// </summary>
        [Output("groupattr")]
        public Output<string> Groupattr { get; private set; } = null!;

        /// <summary>
        /// Base DN under which to perform group search
        /// </summary>
        [Output("groupdn")]
        public Output<string> Groupdn { get; private set; } = null!;

        /// <summary>
        /// Go template used to construct group membership query
        /// </summary>
        [Output("groupfilter")]
        public Output<string> Groupfilter { get; private set; } = null!;

        /// <summary>
        /// Control whether or TLS certificates must be validated
        /// </summary>
        [Output("insecureTls")]
        public Output<bool> InsecureTls { get; private set; } = null!;

        /// <summary>
        /// Specifies if the auth method is local only.
        /// </summary>
        [Output("local")]
        public Output<bool?> Local { get; private set; } = null!;

        /// <summary>
        /// Sets the max page size for LDAP lookups, by default it's set to -1.
        /// *Available only for Vault 1.11.11+, 1.12.7+, and 1.13.3+*.
        /// </summary>
        [Output("maxPageSize")]
        public Output<int?> MaxPageSize { get; private set; } = null!;

        /// <summary>
        /// The namespace to provision the resource in.
        /// The value should not contain leading or trailing forward slashes.
        /// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
        /// *Available only for Vault Enterprise*.
        /// </summary>
        [Output("namespace")]
        public Output<string?> Namespace { get; private set; } = null!;

        /// <summary>
        /// Path to mount the LDAP auth backend under
        /// </summary>
        [Output("path")]
        public Output<string?> Path { get; private set; } = null!;

        /// <summary>
        /// Control use of TLS when conecting to LDAP
        /// </summary>
        [Output("starttls")]
        public Output<bool> Starttls { get; private set; } = null!;

        /// <summary>
        /// Maximum acceptable version of TLS
        /// </summary>
        [Output("tlsMaxVersion")]
        public Output<string> TlsMaxVersion { get; private set; } = null!;

        /// <summary>
        /// Minimum acceptable version of TLS
        /// </summary>
        [Output("tlsMinVersion")]
        public Output<string> TlsMinVersion { get; private set; } = null!;

        /// <summary>
        /// List of CIDR blocks; if set, specifies blocks of IP
        /// addresses which can authenticate successfully, and ties the resulting token to these blocks
        /// as well.
        /// </summary>
        [Output("tokenBoundCidrs")]
        public Output<ImmutableArray<string>> TokenBoundCidrs { get; private set; } = null!;

        /// <summary>
        /// If set, will encode an
        /// [explicit max TTL](https://www.vaultproject.io/docs/concepts/tokens.html#token-time-to-live-periodic-tokens-and-explicit-max-ttls)
        /// onto the token in number of seconds. This is a hard cap even if `token_ttl` and
        /// `token_max_ttl` would otherwise allow a renewal.
        /// </summary>
        [Output("tokenExplicitMaxTtl")]
        public Output<int?> TokenExplicitMaxTtl { get; private set; } = null!;

        /// <summary>
        /// The maximum lifetime for generated tokens in number of seconds.
        /// Its current value will be referenced at renewal time.
        /// </summary>
        [Output("tokenMaxTtl")]
        public Output<int?> TokenMaxTtl { get; private set; } = null!;

        /// <summary>
        /// If set, the default policy will not be set on
        /// generated tokens; otherwise it will be added to the policies set in token_policies.
        /// </summary>
        [Output("tokenNoDefaultPolicy")]
        public Output<bool?> TokenNoDefaultPolicy { get; private set; } = null!;

        /// <summary>
        /// The [maximum number](https://www.vaultproject.io/api-docs/ldap#token_num_uses)
        /// of times a generated token may be used (within its lifetime); 0 means unlimited.
        /// </summary>
        [Output("tokenNumUses")]
        public Output<int?> TokenNumUses { get; private set; } = null!;

        /// <summary>
        /// If set, indicates that the
        /// token generated using this role should never expire. The token should be renewed within the
        /// duration specified by this value. At each renewal, the token's TTL will be set to the
        /// value of this field. Specified in seconds.
        /// </summary>
        [Output("tokenPeriod")]
        public Output<int?> TokenPeriod { get; private set; } = null!;

        /// <summary>
        /// List of policies to encode onto generated tokens. Depending
        /// on the auth method, this list may be supplemented by user/group/other values.
        /// </summary>
        [Output("tokenPolicies")]
        public Output<ImmutableArray<string>> TokenPolicies { get; private set; } = null!;

        /// <summary>
        /// The incremental lifetime for generated tokens in number of seconds.
        /// Its current value will be referenced at renewal time.
        /// </summary>
        [Output("tokenTtl")]
        public Output<int?> TokenTtl { get; private set; } = null!;

        /// <summary>
        /// The type of token to generate, service or batch
        /// </summary>
        [Output("tokenType")]
        public Output<string?> TokenType { get; private set; } = null!;

        /// <summary>
        /// The `userPrincipalDomain` used to construct the UPN string for the authenticating user.
        /// </summary>
        [Output("upndomain")]
        public Output<string> Upndomain { get; private set; } = null!;

        /// <summary>
        /// The URL of the LDAP server
        /// </summary>
        [Output("url")]
        public Output<string> Url { get; private set; } = null!;

        /// <summary>
        /// Use the Active Directory tokenGroups constructed attribute of the user to find the group memberships
        /// </summary>
        [Output("useTokenGroups")]
        public Output<bool> UseTokenGroups { get; private set; } = null!;

        /// <summary>
        /// Attribute on user object matching username passed in
        /// </summary>
        [Output("userattr")]
        public Output<string> Userattr { get; private set; } = null!;

        /// <summary>
        /// Base DN under which to perform user search
        /// </summary>
        [Output("userdn")]
        public Output<string> Userdn { get; private set; } = null!;

        /// <summary>
        /// LDAP user search filter
        /// </summary>
        [Output("userfilter")]
        public Output<string> Userfilter { get; private set; } = null!;

        /// <summary>
        /// Force the auth method to use the username passed by the user as the alias name.
        /// </summary>
        [Output("usernameAsAlias")]
        public Output<bool> UsernameAsAlias { get; private set; } = null!;


        /// <summary>
        /// Create a AuthBackend resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AuthBackend(string name, AuthBackendArgs args, CustomResourceOptions? options = null)
            : base("vault:ldap/authBackend:AuthBackend", name, args ?? new AuthBackendArgs(), MakeResourceOptions(options, ""))
        {
        }

        private AuthBackend(string name, Input<string> id, AuthBackendState? state = null, CustomResourceOptions? options = null)
            : base("vault:ldap/authBackend:AuthBackend", name, state, MakeResourceOptions(options, id))
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
                    "clientTlsKey",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing AuthBackend resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static AuthBackend Get(string name, Input<string> id, AuthBackendState? state = null, CustomResourceOptions? options = null)
        {
            return new AuthBackend(name, id, state, options);
        }
    }

    public sealed class AuthBackendArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// DN of object to bind when performing user search
        /// </summary>
        [Input("binddn")]
        public Input<string>? Binddn { get; set; }

        [Input("bindpass")]
        private Input<string>? _bindpass;

        /// <summary>
        /// Password to use with `binddn` when performing user search
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
        /// Control case senstivity of objects fetched from LDAP, this is used for object matching in vault
        /// </summary>
        [Input("caseSensitiveNames")]
        public Input<bool>? CaseSensitiveNames { get; set; }

        /// <summary>
        /// Trusted CA to validate TLS certificate
        /// </summary>
        [Input("certificate")]
        public Input<string>? Certificate { get; set; }

        [Input("clientTlsCert")]
        public Input<string>? ClientTlsCert { get; set; }

        [Input("clientTlsKey")]
        private Input<string>? _clientTlsKey;
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
        /// Prevents users from bypassing authentication when providing an empty password.
        /// </summary>
        [Input("denyNullBind")]
        public Input<bool>? DenyNullBind { get; set; }

        /// <summary>
        /// Description for the LDAP auth backend mount
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
        /// Use anonymous bind to discover the bind DN of a user.
        /// </summary>
        [Input("discoverdn")]
        public Input<bool>? Discoverdn { get; set; }

        /// <summary>
        /// LDAP attribute to follow on objects returned by groupfilter
        /// </summary>
        [Input("groupattr")]
        public Input<string>? Groupattr { get; set; }

        /// <summary>
        /// Base DN under which to perform group search
        /// </summary>
        [Input("groupdn")]
        public Input<string>? Groupdn { get; set; }

        /// <summary>
        /// Go template used to construct group membership query
        /// </summary>
        [Input("groupfilter")]
        public Input<string>? Groupfilter { get; set; }

        /// <summary>
        /// Control whether or TLS certificates must be validated
        /// </summary>
        [Input("insecureTls")]
        public Input<bool>? InsecureTls { get; set; }

        /// <summary>
        /// Specifies if the auth method is local only.
        /// </summary>
        [Input("local")]
        public Input<bool>? Local { get; set; }

        /// <summary>
        /// Sets the max page size for LDAP lookups, by default it's set to -1.
        /// *Available only for Vault 1.11.11+, 1.12.7+, and 1.13.3+*.
        /// </summary>
        [Input("maxPageSize")]
        public Input<int>? MaxPageSize { get; set; }

        /// <summary>
        /// The namespace to provision the resource in.
        /// The value should not contain leading or trailing forward slashes.
        /// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
        /// *Available only for Vault Enterprise*.
        /// </summary>
        [Input("namespace")]
        public Input<string>? Namespace { get; set; }

        /// <summary>
        /// Path to mount the LDAP auth backend under
        /// </summary>
        [Input("path")]
        public Input<string>? Path { get; set; }

        /// <summary>
        /// Control use of TLS when conecting to LDAP
        /// </summary>
        [Input("starttls")]
        public Input<bool>? Starttls { get; set; }

        /// <summary>
        /// Maximum acceptable version of TLS
        /// </summary>
        [Input("tlsMaxVersion")]
        public Input<string>? TlsMaxVersion { get; set; }

        /// <summary>
        /// Minimum acceptable version of TLS
        /// </summary>
        [Input("tlsMinVersion")]
        public Input<string>? TlsMinVersion { get; set; }

        [Input("tokenBoundCidrs")]
        private InputList<string>? _tokenBoundCidrs;

        /// <summary>
        /// List of CIDR blocks; if set, specifies blocks of IP
        /// addresses which can authenticate successfully, and ties the resulting token to these blocks
        /// as well.
        /// </summary>
        public InputList<string> TokenBoundCidrs
        {
            get => _tokenBoundCidrs ?? (_tokenBoundCidrs = new InputList<string>());
            set => _tokenBoundCidrs = value;
        }

        /// <summary>
        /// If set, will encode an
        /// [explicit max TTL](https://www.vaultproject.io/docs/concepts/tokens.html#token-time-to-live-periodic-tokens-and-explicit-max-ttls)
        /// onto the token in number of seconds. This is a hard cap even if `token_ttl` and
        /// `token_max_ttl` would otherwise allow a renewal.
        /// </summary>
        [Input("tokenExplicitMaxTtl")]
        public Input<int>? TokenExplicitMaxTtl { get; set; }

        /// <summary>
        /// The maximum lifetime for generated tokens in number of seconds.
        /// Its current value will be referenced at renewal time.
        /// </summary>
        [Input("tokenMaxTtl")]
        public Input<int>? TokenMaxTtl { get; set; }

        /// <summary>
        /// If set, the default policy will not be set on
        /// generated tokens; otherwise it will be added to the policies set in token_policies.
        /// </summary>
        [Input("tokenNoDefaultPolicy")]
        public Input<bool>? TokenNoDefaultPolicy { get; set; }

        /// <summary>
        /// The [maximum number](https://www.vaultproject.io/api-docs/ldap#token_num_uses)
        /// of times a generated token may be used (within its lifetime); 0 means unlimited.
        /// </summary>
        [Input("tokenNumUses")]
        public Input<int>? TokenNumUses { get; set; }

        /// <summary>
        /// If set, indicates that the
        /// token generated using this role should never expire. The token should be renewed within the
        /// duration specified by this value. At each renewal, the token's TTL will be set to the
        /// value of this field. Specified in seconds.
        /// </summary>
        [Input("tokenPeriod")]
        public Input<int>? TokenPeriod { get; set; }

        [Input("tokenPolicies")]
        private InputList<string>? _tokenPolicies;

        /// <summary>
        /// List of policies to encode onto generated tokens. Depending
        /// on the auth method, this list may be supplemented by user/group/other values.
        /// </summary>
        public InputList<string> TokenPolicies
        {
            get => _tokenPolicies ?? (_tokenPolicies = new InputList<string>());
            set => _tokenPolicies = value;
        }

        /// <summary>
        /// The incremental lifetime for generated tokens in number of seconds.
        /// Its current value will be referenced at renewal time.
        /// </summary>
        [Input("tokenTtl")]
        public Input<int>? TokenTtl { get; set; }

        /// <summary>
        /// The type of token to generate, service or batch
        /// </summary>
        [Input("tokenType")]
        public Input<string>? TokenType { get; set; }

        /// <summary>
        /// The `userPrincipalDomain` used to construct the UPN string for the authenticating user.
        /// </summary>
        [Input("upndomain")]
        public Input<string>? Upndomain { get; set; }

        /// <summary>
        /// The URL of the LDAP server
        /// </summary>
        [Input("url", required: true)]
        public Input<string> Url { get; set; } = null!;

        /// <summary>
        /// Use the Active Directory tokenGroups constructed attribute of the user to find the group memberships
        /// </summary>
        [Input("useTokenGroups")]
        public Input<bool>? UseTokenGroups { get; set; }

        /// <summary>
        /// Attribute on user object matching username passed in
        /// </summary>
        [Input("userattr")]
        public Input<string>? Userattr { get; set; }

        /// <summary>
        /// Base DN under which to perform user search
        /// </summary>
        [Input("userdn")]
        public Input<string>? Userdn { get; set; }

        /// <summary>
        /// LDAP user search filter
        /// </summary>
        [Input("userfilter")]
        public Input<string>? Userfilter { get; set; }

        /// <summary>
        /// Force the auth method to use the username passed by the user as the alias name.
        /// </summary>
        [Input("usernameAsAlias")]
        public Input<bool>? UsernameAsAlias { get; set; }

        public AuthBackendArgs()
        {
        }
        public static new AuthBackendArgs Empty => new AuthBackendArgs();
    }

    public sealed class AuthBackendState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The accessor for this auth mount.
        /// </summary>
        [Input("accessor")]
        public Input<string>? Accessor { get; set; }

        /// <summary>
        /// DN of object to bind when performing user search
        /// </summary>
        [Input("binddn")]
        public Input<string>? Binddn { get; set; }

        [Input("bindpass")]
        private Input<string>? _bindpass;

        /// <summary>
        /// Password to use with `binddn` when performing user search
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
        /// Control case senstivity of objects fetched from LDAP, this is used for object matching in vault
        /// </summary>
        [Input("caseSensitiveNames")]
        public Input<bool>? CaseSensitiveNames { get; set; }

        /// <summary>
        /// Trusted CA to validate TLS certificate
        /// </summary>
        [Input("certificate")]
        public Input<string>? Certificate { get; set; }

        [Input("clientTlsCert")]
        public Input<string>? ClientTlsCert { get; set; }

        [Input("clientTlsKey")]
        private Input<string>? _clientTlsKey;
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
        /// Prevents users from bypassing authentication when providing an empty password.
        /// </summary>
        [Input("denyNullBind")]
        public Input<bool>? DenyNullBind { get; set; }

        /// <summary>
        /// Description for the LDAP auth backend mount
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
        /// Use anonymous bind to discover the bind DN of a user.
        /// </summary>
        [Input("discoverdn")]
        public Input<bool>? Discoverdn { get; set; }

        /// <summary>
        /// LDAP attribute to follow on objects returned by groupfilter
        /// </summary>
        [Input("groupattr")]
        public Input<string>? Groupattr { get; set; }

        /// <summary>
        /// Base DN under which to perform group search
        /// </summary>
        [Input("groupdn")]
        public Input<string>? Groupdn { get; set; }

        /// <summary>
        /// Go template used to construct group membership query
        /// </summary>
        [Input("groupfilter")]
        public Input<string>? Groupfilter { get; set; }

        /// <summary>
        /// Control whether or TLS certificates must be validated
        /// </summary>
        [Input("insecureTls")]
        public Input<bool>? InsecureTls { get; set; }

        /// <summary>
        /// Specifies if the auth method is local only.
        /// </summary>
        [Input("local")]
        public Input<bool>? Local { get; set; }

        /// <summary>
        /// Sets the max page size for LDAP lookups, by default it's set to -1.
        /// *Available only for Vault 1.11.11+, 1.12.7+, and 1.13.3+*.
        /// </summary>
        [Input("maxPageSize")]
        public Input<int>? MaxPageSize { get; set; }

        /// <summary>
        /// The namespace to provision the resource in.
        /// The value should not contain leading or trailing forward slashes.
        /// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
        /// *Available only for Vault Enterprise*.
        /// </summary>
        [Input("namespace")]
        public Input<string>? Namespace { get; set; }

        /// <summary>
        /// Path to mount the LDAP auth backend under
        /// </summary>
        [Input("path")]
        public Input<string>? Path { get; set; }

        /// <summary>
        /// Control use of TLS when conecting to LDAP
        /// </summary>
        [Input("starttls")]
        public Input<bool>? Starttls { get; set; }

        /// <summary>
        /// Maximum acceptable version of TLS
        /// </summary>
        [Input("tlsMaxVersion")]
        public Input<string>? TlsMaxVersion { get; set; }

        /// <summary>
        /// Minimum acceptable version of TLS
        /// </summary>
        [Input("tlsMinVersion")]
        public Input<string>? TlsMinVersion { get; set; }

        [Input("tokenBoundCidrs")]
        private InputList<string>? _tokenBoundCidrs;

        /// <summary>
        /// List of CIDR blocks; if set, specifies blocks of IP
        /// addresses which can authenticate successfully, and ties the resulting token to these blocks
        /// as well.
        /// </summary>
        public InputList<string> TokenBoundCidrs
        {
            get => _tokenBoundCidrs ?? (_tokenBoundCidrs = new InputList<string>());
            set => _tokenBoundCidrs = value;
        }

        /// <summary>
        /// If set, will encode an
        /// [explicit max TTL](https://www.vaultproject.io/docs/concepts/tokens.html#token-time-to-live-periodic-tokens-and-explicit-max-ttls)
        /// onto the token in number of seconds. This is a hard cap even if `token_ttl` and
        /// `token_max_ttl` would otherwise allow a renewal.
        /// </summary>
        [Input("tokenExplicitMaxTtl")]
        public Input<int>? TokenExplicitMaxTtl { get; set; }

        /// <summary>
        /// The maximum lifetime for generated tokens in number of seconds.
        /// Its current value will be referenced at renewal time.
        /// </summary>
        [Input("tokenMaxTtl")]
        public Input<int>? TokenMaxTtl { get; set; }

        /// <summary>
        /// If set, the default policy will not be set on
        /// generated tokens; otherwise it will be added to the policies set in token_policies.
        /// </summary>
        [Input("tokenNoDefaultPolicy")]
        public Input<bool>? TokenNoDefaultPolicy { get; set; }

        /// <summary>
        /// The [maximum number](https://www.vaultproject.io/api-docs/ldap#token_num_uses)
        /// of times a generated token may be used (within its lifetime); 0 means unlimited.
        /// </summary>
        [Input("tokenNumUses")]
        public Input<int>? TokenNumUses { get; set; }

        /// <summary>
        /// If set, indicates that the
        /// token generated using this role should never expire. The token should be renewed within the
        /// duration specified by this value. At each renewal, the token's TTL will be set to the
        /// value of this field. Specified in seconds.
        /// </summary>
        [Input("tokenPeriod")]
        public Input<int>? TokenPeriod { get; set; }

        [Input("tokenPolicies")]
        private InputList<string>? _tokenPolicies;

        /// <summary>
        /// List of policies to encode onto generated tokens. Depending
        /// on the auth method, this list may be supplemented by user/group/other values.
        /// </summary>
        public InputList<string> TokenPolicies
        {
            get => _tokenPolicies ?? (_tokenPolicies = new InputList<string>());
            set => _tokenPolicies = value;
        }

        /// <summary>
        /// The incremental lifetime for generated tokens in number of seconds.
        /// Its current value will be referenced at renewal time.
        /// </summary>
        [Input("tokenTtl")]
        public Input<int>? TokenTtl { get; set; }

        /// <summary>
        /// The type of token to generate, service or batch
        /// </summary>
        [Input("tokenType")]
        public Input<string>? TokenType { get; set; }

        /// <summary>
        /// The `userPrincipalDomain` used to construct the UPN string for the authenticating user.
        /// </summary>
        [Input("upndomain")]
        public Input<string>? Upndomain { get; set; }

        /// <summary>
        /// The URL of the LDAP server
        /// </summary>
        [Input("url")]
        public Input<string>? Url { get; set; }

        /// <summary>
        /// Use the Active Directory tokenGroups constructed attribute of the user to find the group memberships
        /// </summary>
        [Input("useTokenGroups")]
        public Input<bool>? UseTokenGroups { get; set; }

        /// <summary>
        /// Attribute on user object matching username passed in
        /// </summary>
        [Input("userattr")]
        public Input<string>? Userattr { get; set; }

        /// <summary>
        /// Base DN under which to perform user search
        /// </summary>
        [Input("userdn")]
        public Input<string>? Userdn { get; set; }

        /// <summary>
        /// LDAP user search filter
        /// </summary>
        [Input("userfilter")]
        public Input<string>? Userfilter { get; set; }

        /// <summary>
        /// Force the auth method to use the username passed by the user as the alias name.
        /// </summary>
        [Input("usernameAsAlias")]
        public Input<bool>? UsernameAsAlias { get; set; }

        public AuthBackendState()
        {
        }
        public static new AuthBackendState Empty => new AuthBackendState();
    }
}
