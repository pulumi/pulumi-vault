// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vault.TokenAuth
{
    /// <summary>
    /// Manages Token auth backend role in a Vault server. See the [Vault
    /// documentation](https://www.vaultproject.io/docs/auth/token.html) for more
    /// information.
    /// 
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
    ///     var example = new Vault.TokenAuth.AuthBackendRole("example", new()
    ///     {
    ///         RoleName = "my-role",
    ///         AllowedPolicies = new[]
    ///         {
    ///             "dev",
    ///             "test",
    ///         },
    ///         DisallowedPolicies = new[]
    ///         {
    ///             "default",
    ///         },
    ///         AllowedEntityAliases = new[]
    ///         {
    ///             "test_entity",
    ///         },
    ///         Orphan = true,
    ///         TokenPeriod = 86400,
    ///         Renewable = true,
    ///         TokenExplicitMaxTtl = 115200,
    ///         PathSuffix = "path-suffix",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Token auth backend roles can be imported with `auth/token/roles/` followed by the `role_name`, e.g.
    /// 
    /// ```sh
    /// $ pulumi import vault:tokenauth/authBackendRole:AuthBackendRole example auth/token/roles/my-role
    /// ```
    /// </summary>
    [VaultResourceType("vault:tokenauth/authBackendRole:AuthBackendRole")]
    public partial class AuthBackendRole : global::Pulumi.CustomResource
    {
        /// <summary>
        /// List of allowed entity aliases.
        /// </summary>
        [Output("allowedEntityAliases")]
        public Output<ImmutableArray<string>> AllowedEntityAliases { get; private set; } = null!;

        /// <summary>
        /// List of allowed policies for given role.
        /// </summary>
        [Output("allowedPolicies")]
        public Output<ImmutableArray<string>> AllowedPolicies { get; private set; } = null!;

        /// <summary>
        /// Set of allowed policies with glob match for given role.
        /// </summary>
        [Output("allowedPoliciesGlobs")]
        public Output<ImmutableArray<string>> AllowedPoliciesGlobs { get; private set; } = null!;

        /// <summary>
        /// List of disallowed policies for given role.
        /// </summary>
        [Output("disallowedPolicies")]
        public Output<ImmutableArray<string>> DisallowedPolicies { get; private set; } = null!;

        /// <summary>
        /// Set of disallowed policies with glob match for given role.
        /// </summary>
        [Output("disallowedPoliciesGlobs")]
        public Output<ImmutableArray<string>> DisallowedPoliciesGlobs { get; private set; } = null!;

        /// <summary>
        /// The namespace to provision the resource in.
        /// The value should not contain leading or trailing forward slashes.
        /// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
        /// *Available only for Vault Enterprise*.
        /// </summary>
        [Output("namespace")]
        public Output<string?> Namespace { get; private set; } = null!;

        /// <summary>
        /// If true, tokens created against this policy will be orphan tokens.
        /// </summary>
        [Output("orphan")]
        public Output<bool?> Orphan { get; private set; } = null!;

        /// <summary>
        /// Tokens created against this role will have the given suffix as part of their path in addition to the role name.
        /// </summary>
        [Output("pathSuffix")]
        public Output<string?> PathSuffix { get; private set; } = null!;

        /// <summary>
        /// Whether to disable the ability of the token to be renewed past its initial TTL.
        /// </summary>
        [Output("renewable")]
        public Output<bool?> Renewable { get; private set; } = null!;

        /// <summary>
        /// The name of the role.
        /// </summary>
        [Output("roleName")]
        public Output<string> RoleName { get; private set; } = null!;

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
        /// Create a AuthBackendRole resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AuthBackendRole(string name, AuthBackendRoleArgs args, CustomResourceOptions? options = null)
            : base("vault:tokenauth/authBackendRole:AuthBackendRole", name, args ?? new AuthBackendRoleArgs(), MakeResourceOptions(options, ""))
        {
        }

        private AuthBackendRole(string name, Input<string> id, AuthBackendRoleState? state = null, CustomResourceOptions? options = null)
            : base("vault:tokenauth/authBackendRole:AuthBackendRole", name, state, MakeResourceOptions(options, id))
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
        [Input("allowedEntityAliases")]
        private InputList<string>? _allowedEntityAliases;

        /// <summary>
        /// List of allowed entity aliases.
        /// </summary>
        public InputList<string> AllowedEntityAliases
        {
            get => _allowedEntityAliases ?? (_allowedEntityAliases = new InputList<string>());
            set => _allowedEntityAliases = value;
        }

        [Input("allowedPolicies")]
        private InputList<string>? _allowedPolicies;

        /// <summary>
        /// List of allowed policies for given role.
        /// </summary>
        public InputList<string> AllowedPolicies
        {
            get => _allowedPolicies ?? (_allowedPolicies = new InputList<string>());
            set => _allowedPolicies = value;
        }

        [Input("allowedPoliciesGlobs")]
        private InputList<string>? _allowedPoliciesGlobs;

        /// <summary>
        /// Set of allowed policies with glob match for given role.
        /// </summary>
        public InputList<string> AllowedPoliciesGlobs
        {
            get => _allowedPoliciesGlobs ?? (_allowedPoliciesGlobs = new InputList<string>());
            set => _allowedPoliciesGlobs = value;
        }

        [Input("disallowedPolicies")]
        private InputList<string>? _disallowedPolicies;

        /// <summary>
        /// List of disallowed policies for given role.
        /// </summary>
        public InputList<string> DisallowedPolicies
        {
            get => _disallowedPolicies ?? (_disallowedPolicies = new InputList<string>());
            set => _disallowedPolicies = value;
        }

        [Input("disallowedPoliciesGlobs")]
        private InputList<string>? _disallowedPoliciesGlobs;

        /// <summary>
        /// Set of disallowed policies with glob match for given role.
        /// </summary>
        public InputList<string> DisallowedPoliciesGlobs
        {
            get => _disallowedPoliciesGlobs ?? (_disallowedPoliciesGlobs = new InputList<string>());
            set => _disallowedPoliciesGlobs = value;
        }

        /// <summary>
        /// The namespace to provision the resource in.
        /// The value should not contain leading or trailing forward slashes.
        /// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
        /// *Available only for Vault Enterprise*.
        /// </summary>
        [Input("namespace")]
        public Input<string>? Namespace { get; set; }

        /// <summary>
        /// If true, tokens created against this policy will be orphan tokens.
        /// </summary>
        [Input("orphan")]
        public Input<bool>? Orphan { get; set; }

        /// <summary>
        /// Tokens created against this role will have the given suffix as part of their path in addition to the role name.
        /// </summary>
        [Input("pathSuffix")]
        public Input<string>? PathSuffix { get; set; }

        /// <summary>
        /// Whether to disable the ability of the token to be renewed past its initial TTL.
        /// </summary>
        [Input("renewable")]
        public Input<bool>? Renewable { get; set; }

        /// <summary>
        /// The name of the role.
        /// </summary>
        [Input("roleName", required: true)]
        public Input<string> RoleName { get; set; } = null!;

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

        public AuthBackendRoleArgs()
        {
        }
        public static new AuthBackendRoleArgs Empty => new AuthBackendRoleArgs();
    }

    public sealed class AuthBackendRoleState : global::Pulumi.ResourceArgs
    {
        [Input("allowedEntityAliases")]
        private InputList<string>? _allowedEntityAliases;

        /// <summary>
        /// List of allowed entity aliases.
        /// </summary>
        public InputList<string> AllowedEntityAliases
        {
            get => _allowedEntityAliases ?? (_allowedEntityAliases = new InputList<string>());
            set => _allowedEntityAliases = value;
        }

        [Input("allowedPolicies")]
        private InputList<string>? _allowedPolicies;

        /// <summary>
        /// List of allowed policies for given role.
        /// </summary>
        public InputList<string> AllowedPolicies
        {
            get => _allowedPolicies ?? (_allowedPolicies = new InputList<string>());
            set => _allowedPolicies = value;
        }

        [Input("allowedPoliciesGlobs")]
        private InputList<string>? _allowedPoliciesGlobs;

        /// <summary>
        /// Set of allowed policies with glob match for given role.
        /// </summary>
        public InputList<string> AllowedPoliciesGlobs
        {
            get => _allowedPoliciesGlobs ?? (_allowedPoliciesGlobs = new InputList<string>());
            set => _allowedPoliciesGlobs = value;
        }

        [Input("disallowedPolicies")]
        private InputList<string>? _disallowedPolicies;

        /// <summary>
        /// List of disallowed policies for given role.
        /// </summary>
        public InputList<string> DisallowedPolicies
        {
            get => _disallowedPolicies ?? (_disallowedPolicies = new InputList<string>());
            set => _disallowedPolicies = value;
        }

        [Input("disallowedPoliciesGlobs")]
        private InputList<string>? _disallowedPoliciesGlobs;

        /// <summary>
        /// Set of disallowed policies with glob match for given role.
        /// </summary>
        public InputList<string> DisallowedPoliciesGlobs
        {
            get => _disallowedPoliciesGlobs ?? (_disallowedPoliciesGlobs = new InputList<string>());
            set => _disallowedPoliciesGlobs = value;
        }

        /// <summary>
        /// The namespace to provision the resource in.
        /// The value should not contain leading or trailing forward slashes.
        /// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
        /// *Available only for Vault Enterprise*.
        /// </summary>
        [Input("namespace")]
        public Input<string>? Namespace { get; set; }

        /// <summary>
        /// If true, tokens created against this policy will be orphan tokens.
        /// </summary>
        [Input("orphan")]
        public Input<bool>? Orphan { get; set; }

        /// <summary>
        /// Tokens created against this role will have the given suffix as part of their path in addition to the role name.
        /// </summary>
        [Input("pathSuffix")]
        public Input<string>? PathSuffix { get; set; }

        /// <summary>
        /// Whether to disable the ability of the token to be renewed past its initial TTL.
        /// </summary>
        [Input("renewable")]
        public Input<bool>? Renewable { get; set; }

        /// <summary>
        /// The name of the role.
        /// </summary>
        [Input("roleName")]
        public Input<string>? RoleName { get; set; }

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

        public AuthBackendRoleState()
        {
        }
        public static new AuthBackendRoleState Empty => new AuthBackendRoleState();
    }
}
