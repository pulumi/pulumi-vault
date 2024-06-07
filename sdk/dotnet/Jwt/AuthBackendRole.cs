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
    /// Role for JWT backend:
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Vault = Pulumi.Vault;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var jwt = new Vault.Jwt.AuthBackend("jwt", new()
    ///     {
    ///         Path = "jwt",
    ///     });
    /// 
    ///     var example = new Vault.Jwt.AuthBackendRole("example", new()
    ///     {
    ///         Backend = jwt.Path,
    ///         RoleName = "test-role",
    ///         TokenPolicies = new[]
    ///         {
    ///             "default",
    ///             "dev",
    ///             "prod",
    ///         },
    ///         BoundAudiences = new[]
    ///         {
    ///             "https://myco.test",
    ///         },
    ///         BoundClaims = 
    ///         {
    ///             { "color", "red,green,blue" },
    ///         },
    ///         UserClaim = "https://vault/user",
    ///         RoleType = "jwt",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// Role for OIDC backend:
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Vault = Pulumi.Vault;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var oidc = new Vault.Jwt.AuthBackend("oidc", new()
    ///     {
    ///         Path = "oidc",
    ///         DefaultRole = "test-role",
    ///     });
    /// 
    ///     var example = new Vault.Jwt.AuthBackendRole("example", new()
    ///     {
    ///         Backend = oidc.Path,
    ///         RoleName = "test-role",
    ///         TokenPolicies = new[]
    ///         {
    ///             "default",
    ///             "dev",
    ///             "prod",
    ///         },
    ///         UserClaim = "https://vault/user",
    ///         RoleType = "oidc",
    ///         AllowedRedirectUris = new[]
    ///         {
    ///             "http://localhost:8200/ui/vault/auth/oidc/oidc/callback",
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// JWT authentication backend roles can be imported using the `path`, e.g.
    /// 
    /// ```sh
    /// $ pulumi import vault:jwt/authBackendRole:AuthBackendRole example auth/jwt/role/test-role
    /// ```
    /// </summary>
    [VaultResourceType("vault:jwt/authBackendRole:AuthBackendRole")]
    public partial class AuthBackendRole : global::Pulumi.CustomResource
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
        /// (For "jwt" roles, at least one of `bound_audiences`, `bound_subject`, `bound_claims`
        /// or `token_bound_cidrs` is required. Optional for "oidc" roles.) List of `aud` claims to match against.
        /// Any match is sufficient.
        /// </summary>
        [Output("boundAudiences")]
        public Output<ImmutableArray<string>> BoundAudiences { get; private set; } = null!;

        /// <summary>
        /// If set, a map of claims to values to match against.
        /// A claim's value must be a string, which may contain one value or multiple
        /// comma-separated values, e.g. `"red"` or `"red,green,blue"`.
        /// </summary>
        [Output("boundClaims")]
        public Output<ImmutableDictionary<string, object>?> BoundClaims { get; private set; } = null!;

        /// <summary>
        /// How to interpret values in the claims/values
        /// map (`bound_claims`): can be either `string` (exact match) or `glob` (wildcard
        /// match). Requires Vault 1.4.0 or above.
        /// </summary>
        [Output("boundClaimsType")]
        public Output<string> BoundClaimsType { get; private set; } = null!;

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
        /// Disable bound claim value parsing. Useful when values contain commas.
        /// </summary>
        [Output("disableBoundClaimsParsing")]
        public Output<bool?> DisableBoundClaimsParsing { get; private set; } = null!;

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
        /// Specifies the allowable elapsed time in seconds since the last time 
        /// the user was actively authenticated with the OIDC provider.
        /// </summary>
        [Output("maxAge")]
        public Output<int?> MaxAge { get; private set; } = null!;

        /// <summary>
        /// The namespace to provision the resource in.
        /// The value should not contain leading or trailing forward slashes.
        /// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
        /// *Available only for Vault Enterprise*.
        /// </summary>
        [Output("namespace")]
        public Output<string?> Namespace { get; private set; } = null!;

        /// <summary>
        /// The amount of leeway to add to not before (`nbf`) claims to account for
        /// clock skew, in seconds. Defaults to `60` seconds if set to `0` and can be disabled if set to `-1`.
        /// Only applicable with "jwt" roles.
        /// </summary>
        [Output("notBeforeLeeway")]
        public Output<int?> NotBeforeLeeway { get; private set; } = null!;

        /// <summary>
        /// If set, a list of OIDC scopes to be used with an OIDC role.
        /// The standard scope "openid" is automatically included and need not be specified.
        /// </summary>
        [Output("oidcScopes")]
        public Output<ImmutableArray<string>> OidcScopes { get; private set; } = null!;

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
        /// Specifies the blocks of IP addresses which are allowed to use the generated token
        /// </summary>
        [Output("tokenBoundCidrs")]
        public Output<ImmutableArray<string>> TokenBoundCidrs { get; private set; } = null!;

        /// <summary>
        /// Generated Token's Explicit Maximum TTL in seconds
        /// </summary>
        [Output("tokenExplicitMaxTtl")]
        public Output<int?> TokenExplicitMaxTtl { get; private set; } = null!;

        /// <summary>
        /// The maximum lifetime of the generated token
        /// </summary>
        [Output("tokenMaxTtl")]
        public Output<int?> TokenMaxTtl { get; private set; } = null!;

        /// <summary>
        /// If true, the 'default' policy will not automatically be added to generated tokens
        /// </summary>
        [Output("tokenNoDefaultPolicy")]
        public Output<bool?> TokenNoDefaultPolicy { get; private set; } = null!;

        /// <summary>
        /// The maximum number of times a token may be used, a value of zero means unlimited
        /// </summary>
        [Output("tokenNumUses")]
        public Output<int?> TokenNumUses { get; private set; } = null!;

        /// <summary>
        /// Generated Token's Period
        /// </summary>
        [Output("tokenPeriod")]
        public Output<int?> TokenPeriod { get; private set; } = null!;

        /// <summary>
        /// Generated Token's Policies
        /// </summary>
        [Output("tokenPolicies")]
        public Output<ImmutableArray<string>> TokenPolicies { get; private set; } = null!;

        /// <summary>
        /// The initial ttl of the token to generate in seconds
        /// </summary>
        [Output("tokenTtl")]
        public Output<int?> TokenTtl { get; private set; } = null!;

        /// <summary>
        /// The type of token to generate, service or batch
        /// </summary>
        [Output("tokenType")]
        public Output<string?> TokenType { get; private set; } = null!;

        /// <summary>
        /// The claim to use to uniquely identify
        /// the user; this will be used as the name for the Identity entity alias created
        /// due to a successful login.
        /// </summary>
        [Output("userClaim")]
        public Output<string> UserClaim { get; private set; } = null!;

        /// <summary>
        /// Specifies if the `user_claim` value uses
        /// [JSON pointer](https://www.vaultproject.io/docs/auth/jwt#claim-specifications-and-json-pointer)
        /// syntax for referencing claims. By default, the `user_claim` value will not use JSON pointer.
        /// Requires Vault 1.11+.
        /// </summary>
        [Output("userClaimJsonPointer")]
        public Output<bool?> UserClaimJsonPointer { get; private set; } = null!;

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

    public sealed class AuthBackendRoleArgs : global::Pulumi.ResourceArgs
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
        /// (For "jwt" roles, at least one of `bound_audiences`, `bound_subject`, `bound_claims`
        /// or `token_bound_cidrs` is required. Optional for "oidc" roles.) List of `aud` claims to match against.
        /// Any match is sufficient.
        /// </summary>
        public InputList<string> BoundAudiences
        {
            get => _boundAudiences ?? (_boundAudiences = new InputList<string>());
            set => _boundAudiences = value;
        }

        [Input("boundClaims")]
        private InputMap<object>? _boundClaims;

        /// <summary>
        /// If set, a map of claims to values to match against.
        /// A claim's value must be a string, which may contain one value or multiple
        /// comma-separated values, e.g. `"red"` or `"red,green,blue"`.
        /// </summary>
        public InputMap<object> BoundClaims
        {
            get => _boundClaims ?? (_boundClaims = new InputMap<object>());
            set => _boundClaims = value;
        }

        /// <summary>
        /// How to interpret values in the claims/values
        /// map (`bound_claims`): can be either `string` (exact match) or `glob` (wildcard
        /// match). Requires Vault 1.4.0 or above.
        /// </summary>
        [Input("boundClaimsType")]
        public Input<string>? BoundClaimsType { get; set; }

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
        /// Disable bound claim value parsing. Useful when values contain commas.
        /// </summary>
        [Input("disableBoundClaimsParsing")]
        public Input<bool>? DisableBoundClaimsParsing { get; set; }

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
        /// Specifies the allowable elapsed time in seconds since the last time 
        /// the user was actively authenticated with the OIDC provider.
        /// </summary>
        [Input("maxAge")]
        public Input<int>? MaxAge { get; set; }

        /// <summary>
        /// The namespace to provision the resource in.
        /// The value should not contain leading or trailing forward slashes.
        /// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
        /// *Available only for Vault Enterprise*.
        /// </summary>
        [Input("namespace")]
        public Input<string>? Namespace { get; set; }

        /// <summary>
        /// The amount of leeway to add to not before (`nbf`) claims to account for
        /// clock skew, in seconds. Defaults to `60` seconds if set to `0` and can be disabled if set to `-1`.
        /// Only applicable with "jwt" roles.
        /// </summary>
        [Input("notBeforeLeeway")]
        public Input<int>? NotBeforeLeeway { get; set; }

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
        /// Specifies the blocks of IP addresses which are allowed to use the generated token
        /// </summary>
        public InputList<string> TokenBoundCidrs
        {
            get => _tokenBoundCidrs ?? (_tokenBoundCidrs = new InputList<string>());
            set => _tokenBoundCidrs = value;
        }

        /// <summary>
        /// Generated Token's Explicit Maximum TTL in seconds
        /// </summary>
        [Input("tokenExplicitMaxTtl")]
        public Input<int>? TokenExplicitMaxTtl { get; set; }

        /// <summary>
        /// The maximum lifetime of the generated token
        /// </summary>
        [Input("tokenMaxTtl")]
        public Input<int>? TokenMaxTtl { get; set; }

        /// <summary>
        /// If true, the 'default' policy will not automatically be added to generated tokens
        /// </summary>
        [Input("tokenNoDefaultPolicy")]
        public Input<bool>? TokenNoDefaultPolicy { get; set; }

        /// <summary>
        /// The maximum number of times a token may be used, a value of zero means unlimited
        /// </summary>
        [Input("tokenNumUses")]
        public Input<int>? TokenNumUses { get; set; }

        /// <summary>
        /// Generated Token's Period
        /// </summary>
        [Input("tokenPeriod")]
        public Input<int>? TokenPeriod { get; set; }

        [Input("tokenPolicies")]
        private InputList<string>? _tokenPolicies;

        /// <summary>
        /// Generated Token's Policies
        /// </summary>
        public InputList<string> TokenPolicies
        {
            get => _tokenPolicies ?? (_tokenPolicies = new InputList<string>());
            set => _tokenPolicies = value;
        }

        /// <summary>
        /// The initial ttl of the token to generate in seconds
        /// </summary>
        [Input("tokenTtl")]
        public Input<int>? TokenTtl { get; set; }

        /// <summary>
        /// The type of token to generate, service or batch
        /// </summary>
        [Input("tokenType")]
        public Input<string>? TokenType { get; set; }

        /// <summary>
        /// The claim to use to uniquely identify
        /// the user; this will be used as the name for the Identity entity alias created
        /// due to a successful login.
        /// </summary>
        [Input("userClaim", required: true)]
        public Input<string> UserClaim { get; set; } = null!;

        /// <summary>
        /// Specifies if the `user_claim` value uses
        /// [JSON pointer](https://www.vaultproject.io/docs/auth/jwt#claim-specifications-and-json-pointer)
        /// syntax for referencing claims. By default, the `user_claim` value will not use JSON pointer.
        /// Requires Vault 1.11+.
        /// </summary>
        [Input("userClaimJsonPointer")]
        public Input<bool>? UserClaimJsonPointer { get; set; }

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
        public static new AuthBackendRoleArgs Empty => new AuthBackendRoleArgs();
    }

    public sealed class AuthBackendRoleState : global::Pulumi.ResourceArgs
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
        /// (For "jwt" roles, at least one of `bound_audiences`, `bound_subject`, `bound_claims`
        /// or `token_bound_cidrs` is required. Optional for "oidc" roles.) List of `aud` claims to match against.
        /// Any match is sufficient.
        /// </summary>
        public InputList<string> BoundAudiences
        {
            get => _boundAudiences ?? (_boundAudiences = new InputList<string>());
            set => _boundAudiences = value;
        }

        [Input("boundClaims")]
        private InputMap<object>? _boundClaims;

        /// <summary>
        /// If set, a map of claims to values to match against.
        /// A claim's value must be a string, which may contain one value or multiple
        /// comma-separated values, e.g. `"red"` or `"red,green,blue"`.
        /// </summary>
        public InputMap<object> BoundClaims
        {
            get => _boundClaims ?? (_boundClaims = new InputMap<object>());
            set => _boundClaims = value;
        }

        /// <summary>
        /// How to interpret values in the claims/values
        /// map (`bound_claims`): can be either `string` (exact match) or `glob` (wildcard
        /// match). Requires Vault 1.4.0 or above.
        /// </summary>
        [Input("boundClaimsType")]
        public Input<string>? BoundClaimsType { get; set; }

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
        /// Disable bound claim value parsing. Useful when values contain commas.
        /// </summary>
        [Input("disableBoundClaimsParsing")]
        public Input<bool>? DisableBoundClaimsParsing { get; set; }

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
        /// Specifies the allowable elapsed time in seconds since the last time 
        /// the user was actively authenticated with the OIDC provider.
        /// </summary>
        [Input("maxAge")]
        public Input<int>? MaxAge { get; set; }

        /// <summary>
        /// The namespace to provision the resource in.
        /// The value should not contain leading or trailing forward slashes.
        /// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
        /// *Available only for Vault Enterprise*.
        /// </summary>
        [Input("namespace")]
        public Input<string>? Namespace { get; set; }

        /// <summary>
        /// The amount of leeway to add to not before (`nbf`) claims to account for
        /// clock skew, in seconds. Defaults to `60` seconds if set to `0` and can be disabled if set to `-1`.
        /// Only applicable with "jwt" roles.
        /// </summary>
        [Input("notBeforeLeeway")]
        public Input<int>? NotBeforeLeeway { get; set; }

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
        /// Specifies the blocks of IP addresses which are allowed to use the generated token
        /// </summary>
        public InputList<string> TokenBoundCidrs
        {
            get => _tokenBoundCidrs ?? (_tokenBoundCidrs = new InputList<string>());
            set => _tokenBoundCidrs = value;
        }

        /// <summary>
        /// Generated Token's Explicit Maximum TTL in seconds
        /// </summary>
        [Input("tokenExplicitMaxTtl")]
        public Input<int>? TokenExplicitMaxTtl { get; set; }

        /// <summary>
        /// The maximum lifetime of the generated token
        /// </summary>
        [Input("tokenMaxTtl")]
        public Input<int>? TokenMaxTtl { get; set; }

        /// <summary>
        /// If true, the 'default' policy will not automatically be added to generated tokens
        /// </summary>
        [Input("tokenNoDefaultPolicy")]
        public Input<bool>? TokenNoDefaultPolicy { get; set; }

        /// <summary>
        /// The maximum number of times a token may be used, a value of zero means unlimited
        /// </summary>
        [Input("tokenNumUses")]
        public Input<int>? TokenNumUses { get; set; }

        /// <summary>
        /// Generated Token's Period
        /// </summary>
        [Input("tokenPeriod")]
        public Input<int>? TokenPeriod { get; set; }

        [Input("tokenPolicies")]
        private InputList<string>? _tokenPolicies;

        /// <summary>
        /// Generated Token's Policies
        /// </summary>
        public InputList<string> TokenPolicies
        {
            get => _tokenPolicies ?? (_tokenPolicies = new InputList<string>());
            set => _tokenPolicies = value;
        }

        /// <summary>
        /// The initial ttl of the token to generate in seconds
        /// </summary>
        [Input("tokenTtl")]
        public Input<int>? TokenTtl { get; set; }

        /// <summary>
        /// The type of token to generate, service or batch
        /// </summary>
        [Input("tokenType")]
        public Input<string>? TokenType { get; set; }

        /// <summary>
        /// The claim to use to uniquely identify
        /// the user; this will be used as the name for the Identity entity alias created
        /// due to a successful login.
        /// </summary>
        [Input("userClaim")]
        public Input<string>? UserClaim { get; set; }

        /// <summary>
        /// Specifies if the `user_claim` value uses
        /// [JSON pointer](https://www.vaultproject.io/docs/auth/jwt#claim-specifications-and-json-pointer)
        /// syntax for referencing claims. By default, the `user_claim` value will not use JSON pointer.
        /// Requires Vault 1.11+.
        /// </summary>
        [Input("userClaimJsonPointer")]
        public Input<bool>? UserClaimJsonPointer { get; set; }

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
        public static new AuthBackendRoleState Empty => new AuthBackendRoleState();
    }
}
