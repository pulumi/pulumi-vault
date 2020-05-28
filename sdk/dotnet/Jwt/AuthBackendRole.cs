// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vault.Jwt
{
    /// <summary>
    /// Manages an JWT/OIDC auth backend role in a Vault server. See the [Vault
    /// documentation](https://www.vaultproject.io/docs/auth/jwt.html) for more
    /// information.
    /// 
    /// ## Example Usage
    /// 
    /// 
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Vault = Pulumi.Vault;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var jwt = new Vault.Jwt.AuthBackend("jwt", new Vault.Jwt.AuthBackendArgs
    ///         {
    ///             Path = "jwt",
    ///         });
    ///         var example = new Vault.Jwt.AuthBackendRole("example", new Vault.Jwt.AuthBackendRoleArgs
    ///         {
    ///             Backend = jwt.Path,
    ///             RoleName = "test-role",
    ///             TokenPolicies = 
    ///             {
    ///                 "default",
    ///                 "dev",
    ///                 "prod",
    ///             },
    ///             BoundAudiences = 
    ///             {
    ///                 "https://myco.test",
    ///             },
    ///             UserClaim = "https://vault/user",
    ///             RoleType = "jwt",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// </summary>
    public partial class AuthBackendRole : Pulumi.CustomResource
    {
        /// <summary>
        /// The list of allowed values for redirect_uri during OIDC logins.
        /// Required for OIDC roles
        /// </summary>
        [Output("allowedRedirectUris")]
        public Output<ImmutableArray<string>> AllowedRedirectUris { get; private set; } = null!;

        /// <summary>
        /// The unique name of the auth backend to configure.
        /// Defaults to `jwt`.
        /// </summary>
        [Output("backend")]
        public Output<string?> Backend { get; private set; } = null!;

        /// <summary>
        /// (Required for roles of type `jwt`, optional for roles of
        /// type `oidc`) List of `aud` claims to match against. Any match is sufficient.
        /// </summary>
        [Output("boundAudiences")]
        public Output<ImmutableArray<string>> BoundAudiences { get; private set; } = null!;

        /// <summary>
        /// If set, a list of
        /// CIDRs valid as the source address for login requests. This value is also encoded into any resulting token.
        /// </summary>
        [Output("boundCidrs")]
        public Output<ImmutableArray<string>> BoundCidrs { get; private set; } = null!;

        /// <summary>
        /// If set, a map of claims/values to match against.
        /// The expected value may be a single string or a list of strings.
        /// </summary>
        [Output("boundClaims")]
        public Output<ImmutableDictionary<string, object>?> BoundClaims { get; private set; } = null!;

        /// <summary>
        /// If set, requires that the `sub` claim matches
        /// this value.
        /// </summary>
        [Output("boundSubject")]
        public Output<string?> BoundSubject { get; private set; } = null!;

        /// <summary>
        /// If set, a map of claims (keys) to be copied
        /// to specified metadata fields (values).
        /// </summary>
        [Output("claimMappings")]
        public Output<ImmutableDictionary<string, object>?> ClaimMappings { get; private set; } = null!;

        /// <summary>
        /// The amount of leeway to add to all claims to account for clock skew, in
        /// seconds. Defaults to `60` seconds if set to `0` and can be disabled if set to `-1`.
        /// Only applicable with "jwt" roles.
        /// </summary>
        [Output("clockSkewLeeway")]
        public Output<int?> ClockSkewLeeway { get; private set; } = null!;

        /// <summary>
        /// The amount of leeway to add to expiration (`exp`) claims to account for
        /// clock skew, in seconds. Defaults to `60` seconds if set to `0` and can be disabled if set to `-1`.
        /// Only applicable with "jwt" roles.
        /// </summary>
        [Output("expirationLeeway")]
        public Output<int?> ExpirationLeeway { get; private set; } = null!;

        /// <summary>
        /// The claim to use to uniquely identify
        /// the set of groups to which the user belongs; this will be used as the names
        /// for the Identity group aliases created due to a successful login. The claim
        /// value must be a list of strings.
        /// </summary>
        [Output("groupsClaim")]
        public Output<string?> GroupsClaim { get; private set; } = null!;

        /// <summary>
        /// (Optional; Deprecated. This field has been
        /// removed since Vault 1.1. If the groups claim is not at the top level, it can
        /// now be specified as a [JSONPointer](https://tools.ietf.org/html/rfc6901).)
        /// A pattern of delimiters
        /// used to allow the groups_claim to live outside of the top-level JWT structure.
        /// For instance, a groups_claim of meta/user.name/groups with this field
        /// set to // will expect nested structures named meta, user.name, and groups.
        /// If this field was set to /./ the groups information would expect to be
        /// via nested structures of meta, user, name, and groups.
        /// </summary>
        [Output("groupsClaimDelimiterPattern")]
        public Output<string?> GroupsClaimDelimiterPattern { get; private set; } = null!;

        /// <summary>
        /// The maximum allowed lifetime of tokens
        /// issued using this role, provided as a number of seconds.
        /// </summary>
        [Output("maxTtl")]
        public Output<int?> MaxTtl { get; private set; } = null!;

        /// <summary>
        /// The amount of leeway to add to not before (`nbf`) claims to account for
        /// clock skew, in seconds. Defaults to `60` seconds if set to `0` and can be disabled if set to `-1`.
        /// Only applicable with "jwt" roles.
        /// </summary>
        [Output("notBeforeLeeway")]
        public Output<int?> NotBeforeLeeway { get; private set; } = null!;

        /// <summary>
        /// If set, puts a use-count
        /// limitation on the issued token.
        /// </summary>
        [Output("numUses")]
        public Output<int?> NumUses { get; private set; } = null!;

        /// <summary>
        /// If set, a list of OIDC scopes to be used with an OIDC role.
        /// The standard scope "openid" is automatically included and need not be specified.
        /// </summary>
        [Output("oidcScopes")]
        public Output<ImmutableArray<string>> OidcScopes { get; private set; } = null!;

        /// <summary>
        /// If set, indicates that the
        /// token generated using this role should never expire. The token should be renewed within the
        /// duration specified by this value. At each renewal, the token's TTL will be set to the
        /// value of this field. Specified in seconds.
        /// </summary>
        [Output("period")]
        public Output<int?> Period { get; private set; } = null!;

        /// <summary>
        /// An array of strings
        /// specifying the policies to be set on tokens issued using this role.
        /// </summary>
        [Output("policies")]
        public Output<ImmutableArray<string>> Policies { get; private set; } = null!;

        /// <summary>
        /// The name of the role.
        /// </summary>
        [Output("roleName")]
        public Output<string> RoleName { get; private set; } = null!;

        /// <summary>
        /// Type of role, either "oidc" (default) or "jwt".
        /// </summary>
        [Output("roleType")]
        public Output<string> RoleType { get; private set; } = null!;

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
        /// The
        /// [period](https://www.vaultproject.io/docs/concepts/tokens.html#token-time-to-live-periodic-tokens-and-explicit-max-ttls),
        /// if any, in number of seconds to set on the token.
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
        /// The type of token that should be generated. Can be `service`,
        /// `batch`, or `default` to use the mount's tuned default (which unless changed will be
        /// `service` tokens). For token store roles, there are two additional possibilities:
        /// `default-service` and `default-batch` which specify the type to return unless the client
        /// requests a different type at generation time.
        /// </summary>
        [Output("tokenType")]
        public Output<string?> TokenType { get; private set; } = null!;

        /// <summary>
        /// The TTL period of tokens issued
        /// using this role, provided as a number of seconds.
        /// </summary>
        [Output("ttl")]
        public Output<int?> Ttl { get; private set; } = null!;

        /// <summary>
        /// The claim to use to uniquely identify
        /// the user; this will be used as the name for the Identity entity alias created
        /// due to a successful login.
        /// </summary>
        [Output("userClaim")]
        public Output<string> UserClaim { get; private set; } = null!;

        /// <summary>
        /// Log received OIDC tokens and claims when debug-level
        /// logging is active. Not recommended in production since sensitive information may be present
        /// in OIDC responses.
        /// </summary>
        [Output("verboseOidcLogging")]
        public Output<bool?> VerboseOidcLogging { get; private set; } = null!;


        /// <summary>
        /// Create a AuthBackendRole resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AuthBackendRole(string name, AuthBackendRoleArgs args, CustomResourceOptions? options = null)
            : base("vault:jwt/authBackendRole:AuthBackendRole", name, args ?? new AuthBackendRoleArgs(), MakeResourceOptions(options, ""))
        {
        }

        private AuthBackendRole(string name, Input<string> id, AuthBackendRoleState? state = null, CustomResourceOptions? options = null)
            : base("vault:jwt/authBackendRole:AuthBackendRole", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing AuthBackendRole resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static AuthBackendRole Get(string name, Input<string> id, AuthBackendRoleState? state = null, CustomResourceOptions? options = null)
        {
            return new AuthBackendRole(name, id, state, options);
        }
    }

    public sealed class AuthBackendRoleArgs : Pulumi.ResourceArgs
    {
        [Input("allowedRedirectUris")]
        private InputList<string>? _allowedRedirectUris;

        /// <summary>
        /// The list of allowed values for redirect_uri during OIDC logins.
        /// Required for OIDC roles
        /// </summary>
        public InputList<string> AllowedRedirectUris
        {
            get => _allowedRedirectUris ?? (_allowedRedirectUris = new InputList<string>());
            set => _allowedRedirectUris = value;
        }

        /// <summary>
        /// The unique name of the auth backend to configure.
        /// Defaults to `jwt`.
        /// </summary>
        [Input("backend")]
        public Input<string>? Backend { get; set; }

        [Input("boundAudiences")]
        private InputList<string>? _boundAudiences;

        /// <summary>
        /// (Required for roles of type `jwt`, optional for roles of
        /// type `oidc`) List of `aud` claims to match against. Any match is sufficient.
        /// </summary>
        public InputList<string> BoundAudiences
        {
            get => _boundAudiences ?? (_boundAudiences = new InputList<string>());
            set => _boundAudiences = value;
        }

        [Input("boundCidrs")]
        private InputList<string>? _boundCidrs;

        /// <summary>
        /// If set, a list of
        /// CIDRs valid as the source address for login requests. This value is also encoded into any resulting token.
        /// </summary>
        [Obsolete(@"use `token_bound_cidrs` instead if you are running Vault >= 1.2")]
        public InputList<string> BoundCidrs
        {
            get => _boundCidrs ?? (_boundCidrs = new InputList<string>());
            set => _boundCidrs = value;
        }

        [Input("boundClaims")]
        private InputMap<object>? _boundClaims;

        /// <summary>
        /// If set, a map of claims/values to match against.
        /// The expected value may be a single string or a list of strings.
        /// </summary>
        public InputMap<object> BoundClaims
        {
            get => _boundClaims ?? (_boundClaims = new InputMap<object>());
            set => _boundClaims = value;
        }

        /// <summary>
        /// If set, requires that the `sub` claim matches
        /// this value.
        /// </summary>
        [Input("boundSubject")]
        public Input<string>? BoundSubject { get; set; }

        [Input("claimMappings")]
        private InputMap<object>? _claimMappings;

        /// <summary>
        /// If set, a map of claims (keys) to be copied
        /// to specified metadata fields (values).
        /// </summary>
        public InputMap<object> ClaimMappings
        {
            get => _claimMappings ?? (_claimMappings = new InputMap<object>());
            set => _claimMappings = value;
        }

        /// <summary>
        /// The amount of leeway to add to all claims to account for clock skew, in
        /// seconds. Defaults to `60` seconds if set to `0` and can be disabled if set to `-1`.
        /// Only applicable with "jwt" roles.
        /// </summary>
        [Input("clockSkewLeeway")]
        public Input<int>? ClockSkewLeeway { get; set; }

        /// <summary>
        /// The amount of leeway to add to expiration (`exp`) claims to account for
        /// clock skew, in seconds. Defaults to `60` seconds if set to `0` and can be disabled if set to `-1`.
        /// Only applicable with "jwt" roles.
        /// </summary>
        [Input("expirationLeeway")]
        public Input<int>? ExpirationLeeway { get; set; }

        /// <summary>
        /// The claim to use to uniquely identify
        /// the set of groups to which the user belongs; this will be used as the names
        /// for the Identity group aliases created due to a successful login. The claim
        /// value must be a list of strings.
        /// </summary>
        [Input("groupsClaim")]
        public Input<string>? GroupsClaim { get; set; }

        /// <summary>
        /// (Optional; Deprecated. This field has been
        /// removed since Vault 1.1. If the groups claim is not at the top level, it can
        /// now be specified as a [JSONPointer](https://tools.ietf.org/html/rfc6901).)
        /// A pattern of delimiters
        /// used to allow the groups_claim to live outside of the top-level JWT structure.
        /// For instance, a groups_claim of meta/user.name/groups with this field
        /// set to // will expect nested structures named meta, user.name, and groups.
        /// If this field was set to /./ the groups information would expect to be
        /// via nested structures of meta, user, name, and groups.
        /// </summary>
        [Input("groupsClaimDelimiterPattern")]
        public Input<string>? GroupsClaimDelimiterPattern { get; set; }

        /// <summary>
        /// The maximum allowed lifetime of tokens
        /// issued using this role, provided as a number of seconds.
        /// </summary>
        [Input("maxTtl")]
        public Input<int>? MaxTtl { get; set; }

        /// <summary>
        /// The amount of leeway to add to not before (`nbf`) claims to account for
        /// clock skew, in seconds. Defaults to `60` seconds if set to `0` and can be disabled if set to `-1`.
        /// Only applicable with "jwt" roles.
        /// </summary>
        [Input("notBeforeLeeway")]
        public Input<int>? NotBeforeLeeway { get; set; }

        /// <summary>
        /// If set, puts a use-count
        /// limitation on the issued token.
        /// </summary>
        [Input("numUses")]
        public Input<int>? NumUses { get; set; }

        [Input("oidcScopes")]
        private InputList<string>? _oidcScopes;

        /// <summary>
        /// If set, a list of OIDC scopes to be used with an OIDC role.
        /// The standard scope "openid" is automatically included and need not be specified.
        /// </summary>
        public InputList<string> OidcScopes
        {
            get => _oidcScopes ?? (_oidcScopes = new InputList<string>());
            set => _oidcScopes = value;
        }

        /// <summary>
        /// If set, indicates that the
        /// token generated using this role should never expire. The token should be renewed within the
        /// duration specified by this value. At each renewal, the token's TTL will be set to the
        /// value of this field. Specified in seconds.
        /// </summary>
        [Input("period")]
        public Input<int>? Period { get; set; }

        [Input("policies")]
        private InputList<string>? _policies;

        /// <summary>
        /// An array of strings
        /// specifying the policies to be set on tokens issued using this role.
        /// </summary>
        [Obsolete(@"use `token_policies` instead if you are running Vault >= 1.2")]
        public InputList<string> Policies
        {
            get => _policies ?? (_policies = new InputList<string>());
            set => _policies = value;
        }

        /// <summary>
        /// The name of the role.
        /// </summary>
        [Input("roleName", required: true)]
        public Input<string> RoleName { get; set; } = null!;

        /// <summary>
        /// Type of role, either "oidc" (default) or "jwt".
        /// </summary>
        [Input("roleType")]
        public Input<string>? RoleType { get; set; }

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
        /// The
        /// [period](https://www.vaultproject.io/docs/concepts/tokens.html#token-time-to-live-periodic-tokens-and-explicit-max-ttls),
        /// if any, in number of seconds to set on the token.
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
        /// The type of token that should be generated. Can be `service`,
        /// `batch`, or `default` to use the mount's tuned default (which unless changed will be
        /// `service` tokens). For token store roles, there are two additional possibilities:
        /// `default-service` and `default-batch` which specify the type to return unless the client
        /// requests a different type at generation time.
        /// </summary>
        [Input("tokenType")]
        public Input<string>? TokenType { get; set; }

        /// <summary>
        /// The TTL period of tokens issued
        /// using this role, provided as a number of seconds.
        /// </summary>
        [Input("ttl")]
        public Input<int>? Ttl { get; set; }

        /// <summary>
        /// The claim to use to uniquely identify
        /// the user; this will be used as the name for the Identity entity alias created
        /// due to a successful login.
        /// </summary>
        [Input("userClaim", required: true)]
        public Input<string> UserClaim { get; set; } = null!;

        /// <summary>
        /// Log received OIDC tokens and claims when debug-level
        /// logging is active. Not recommended in production since sensitive information may be present
        /// in OIDC responses.
        /// </summary>
        [Input("verboseOidcLogging")]
        public Input<bool>? VerboseOidcLogging { get; set; }

        public AuthBackendRoleArgs()
        {
        }
    }

    public sealed class AuthBackendRoleState : Pulumi.ResourceArgs
    {
        [Input("allowedRedirectUris")]
        private InputList<string>? _allowedRedirectUris;

        /// <summary>
        /// The list of allowed values for redirect_uri during OIDC logins.
        /// Required for OIDC roles
        /// </summary>
        public InputList<string> AllowedRedirectUris
        {
            get => _allowedRedirectUris ?? (_allowedRedirectUris = new InputList<string>());
            set => _allowedRedirectUris = value;
        }

        /// <summary>
        /// The unique name of the auth backend to configure.
        /// Defaults to `jwt`.
        /// </summary>
        [Input("backend")]
        public Input<string>? Backend { get; set; }

        [Input("boundAudiences")]
        private InputList<string>? _boundAudiences;

        /// <summary>
        /// (Required for roles of type `jwt`, optional for roles of
        /// type `oidc`) List of `aud` claims to match against. Any match is sufficient.
        /// </summary>
        public InputList<string> BoundAudiences
        {
            get => _boundAudiences ?? (_boundAudiences = new InputList<string>());
            set => _boundAudiences = value;
        }

        [Input("boundCidrs")]
        private InputList<string>? _boundCidrs;

        /// <summary>
        /// If set, a list of
        /// CIDRs valid as the source address for login requests. This value is also encoded into any resulting token.
        /// </summary>
        [Obsolete(@"use `token_bound_cidrs` instead if you are running Vault >= 1.2")]
        public InputList<string> BoundCidrs
        {
            get => _boundCidrs ?? (_boundCidrs = new InputList<string>());
            set => _boundCidrs = value;
        }

        [Input("boundClaims")]
        private InputMap<object>? _boundClaims;

        /// <summary>
        /// If set, a map of claims/values to match against.
        /// The expected value may be a single string or a list of strings.
        /// </summary>
        public InputMap<object> BoundClaims
        {
            get => _boundClaims ?? (_boundClaims = new InputMap<object>());
            set => _boundClaims = value;
        }

        /// <summary>
        /// If set, requires that the `sub` claim matches
        /// this value.
        /// </summary>
        [Input("boundSubject")]
        public Input<string>? BoundSubject { get; set; }

        [Input("claimMappings")]
        private InputMap<object>? _claimMappings;

        /// <summary>
        /// If set, a map of claims (keys) to be copied
        /// to specified metadata fields (values).
        /// </summary>
        public InputMap<object> ClaimMappings
        {
            get => _claimMappings ?? (_claimMappings = new InputMap<object>());
            set => _claimMappings = value;
        }

        /// <summary>
        /// The amount of leeway to add to all claims to account for clock skew, in
        /// seconds. Defaults to `60` seconds if set to `0` and can be disabled if set to `-1`.
        /// Only applicable with "jwt" roles.
        /// </summary>
        [Input("clockSkewLeeway")]
        public Input<int>? ClockSkewLeeway { get; set; }

        /// <summary>
        /// The amount of leeway to add to expiration (`exp`) claims to account for
        /// clock skew, in seconds. Defaults to `60` seconds if set to `0` and can be disabled if set to `-1`.
        /// Only applicable with "jwt" roles.
        /// </summary>
        [Input("expirationLeeway")]
        public Input<int>? ExpirationLeeway { get; set; }

        /// <summary>
        /// The claim to use to uniquely identify
        /// the set of groups to which the user belongs; this will be used as the names
        /// for the Identity group aliases created due to a successful login. The claim
        /// value must be a list of strings.
        /// </summary>
        [Input("groupsClaim")]
        public Input<string>? GroupsClaim { get; set; }

        /// <summary>
        /// (Optional; Deprecated. This field has been
        /// removed since Vault 1.1. If the groups claim is not at the top level, it can
        /// now be specified as a [JSONPointer](https://tools.ietf.org/html/rfc6901).)
        /// A pattern of delimiters
        /// used to allow the groups_claim to live outside of the top-level JWT structure.
        /// For instance, a groups_claim of meta/user.name/groups with this field
        /// set to // will expect nested structures named meta, user.name, and groups.
        /// If this field was set to /./ the groups information would expect to be
        /// via nested structures of meta, user, name, and groups.
        /// </summary>
        [Input("groupsClaimDelimiterPattern")]
        public Input<string>? GroupsClaimDelimiterPattern { get; set; }

        /// <summary>
        /// The maximum allowed lifetime of tokens
        /// issued using this role, provided as a number of seconds.
        /// </summary>
        [Input("maxTtl")]
        public Input<int>? MaxTtl { get; set; }

        /// <summary>
        /// The amount of leeway to add to not before (`nbf`) claims to account for
        /// clock skew, in seconds. Defaults to `60` seconds if set to `0` and can be disabled if set to `-1`.
        /// Only applicable with "jwt" roles.
        /// </summary>
        [Input("notBeforeLeeway")]
        public Input<int>? NotBeforeLeeway { get; set; }

        /// <summary>
        /// If set, puts a use-count
        /// limitation on the issued token.
        /// </summary>
        [Input("numUses")]
        public Input<int>? NumUses { get; set; }

        [Input("oidcScopes")]
        private InputList<string>? _oidcScopes;

        /// <summary>
        /// If set, a list of OIDC scopes to be used with an OIDC role.
        /// The standard scope "openid" is automatically included and need not be specified.
        /// </summary>
        public InputList<string> OidcScopes
        {
            get => _oidcScopes ?? (_oidcScopes = new InputList<string>());
            set => _oidcScopes = value;
        }

        /// <summary>
        /// If set, indicates that the
        /// token generated using this role should never expire. The token should be renewed within the
        /// duration specified by this value. At each renewal, the token's TTL will be set to the
        /// value of this field. Specified in seconds.
        /// </summary>
        [Input("period")]
        public Input<int>? Period { get; set; }

        [Input("policies")]
        private InputList<string>? _policies;

        /// <summary>
        /// An array of strings
        /// specifying the policies to be set on tokens issued using this role.
        /// </summary>
        [Obsolete(@"use `token_policies` instead if you are running Vault >= 1.2")]
        public InputList<string> Policies
        {
            get => _policies ?? (_policies = new InputList<string>());
            set => _policies = value;
        }

        /// <summary>
        /// The name of the role.
        /// </summary>
        [Input("roleName")]
        public Input<string>? RoleName { get; set; }

        /// <summary>
        /// Type of role, either "oidc" (default) or "jwt".
        /// </summary>
        [Input("roleType")]
        public Input<string>? RoleType { get; set; }

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
        /// The
        /// [period](https://www.vaultproject.io/docs/concepts/tokens.html#token-time-to-live-periodic-tokens-and-explicit-max-ttls),
        /// if any, in number of seconds to set on the token.
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
        /// The type of token that should be generated. Can be `service`,
        /// `batch`, or `default` to use the mount's tuned default (which unless changed will be
        /// `service` tokens). For token store roles, there are two additional possibilities:
        /// `default-service` and `default-batch` which specify the type to return unless the client
        /// requests a different type at generation time.
        /// </summary>
        [Input("tokenType")]
        public Input<string>? TokenType { get; set; }

        /// <summary>
        /// The TTL period of tokens issued
        /// using this role, provided as a number of seconds.
        /// </summary>
        [Input("ttl")]
        public Input<int>? Ttl { get; set; }

        /// <summary>
        /// The claim to use to uniquely identify
        /// the user; this will be used as the name for the Identity entity alias created
        /// due to a successful login.
        /// </summary>
        [Input("userClaim")]
        public Input<string>? UserClaim { get; set; }

        /// <summary>
        /// Log received OIDC tokens and claims when debug-level
        /// logging is active. Not recommended in production since sensitive information may be present
        /// in OIDC responses.
        /// </summary>
        [Input("verboseOidcLogging")]
        public Input<bool>? VerboseOidcLogging { get; set; }

        public AuthBackendRoleState()
        {
        }
    }
}
