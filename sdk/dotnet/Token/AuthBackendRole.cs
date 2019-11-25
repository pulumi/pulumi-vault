// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

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
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-vault/blob/master/website/docs/r/token_auth_backend_role.html.markdown.
    /// </summary>
    public partial class AuthBackendRole : Pulumi.CustomResource
    {
        /// <summary>
        /// List of allowed policies for given role.
        /// </summary>
        [Output("allowedPolicies")]
        public Output<ImmutableArray<string>> AllowedPolicies { get; private set; } = null!;

        /// <summary>
        /// If set, a list of
        /// CIDRs valid as the source address for login requests. This value is also encoded into any resulting token.
        /// </summary>
        [Output("boundCidrs")]
        public Output<ImmutableArray<string>> BoundCidrs { get; private set; } = null!;

        /// <summary>
        /// List of disallowed policies for given role.
        /// </summary>
        [Output("disallowedPolicies")]
        public Output<ImmutableArray<string>> DisallowedPolicies { get; private set; } = null!;

        /// <summary>
        /// Number of seconds after which issued tokens can no longer be renewed.
        /// </summary>
        [Output("explicitMaxTtl")]
        public Output<string?> ExplicitMaxTtl { get; private set; } = null!;

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
        /// If set, indicates that the
        /// token generated using this role should never expire. The token should be renewed within the
        /// duration specified by this value. At each renewal, the token's TTL will be set to the
        /// value of this field. Specified in seconds.
        /// </summary>
        [Output("period")]
        public Output<string?> Period { get; private set; } = null!;

        /// <summary>
        /// Wether to disable the ability of the token to be renewed past its initial TTL.
        /// </summary>
        [Output("renewable")]
        public Output<bool?> Renewable { get; private set; } = null!;

        /// <summary>
        /// The name of the role.
        /// </summary>
        [Output("roleName")]
        public Output<string> RoleName { get; private set; } = null!;

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
        /// Generated Token's Policies
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
        /// Create a AuthBackendRole resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AuthBackendRole(string name, AuthBackendRoleArgs args, CustomResourceOptions? options = null)
            : base("vault:token/authBackendRole:AuthBackendRole", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private AuthBackendRole(string name, Input<string> id, AuthBackendRoleState? state = null, CustomResourceOptions? options = null)
            : base("vault:token/authBackendRole:AuthBackendRole", name, state, MakeResourceOptions(options, id))
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

        [Input("boundCidrs")]
        private InputList<string>? _boundCidrs;

        /// <summary>
        /// If set, a list of
        /// CIDRs valid as the source address for login requests. This value is also encoded into any resulting token.
        /// </summary>
        public InputList<string> BoundCidrs
        {
            get => _boundCidrs ?? (_boundCidrs = new InputList<string>());
            set => _boundCidrs = value;
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

        /// <summary>
        /// Number of seconds after which issued tokens can no longer be renewed.
        /// </summary>
        [Input("explicitMaxTtl")]
        public Input<string>? ExplicitMaxTtl { get; set; }

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
        /// If set, indicates that the
        /// token generated using this role should never expire. The token should be renewed within the
        /// duration specified by this value. At each renewal, the token's TTL will be set to the
        /// value of this field. Specified in seconds.
        /// </summary>
        [Input("period")]
        public Input<string>? Period { get; set; }

        /// <summary>
        /// Wether to disable the ability of the token to be renewed past its initial TTL.
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
        /// Generated Token's Policies
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

        public AuthBackendRoleArgs()
        {
        }
    }

    public sealed class AuthBackendRoleState : Pulumi.ResourceArgs
    {
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

        [Input("boundCidrs")]
        private InputList<string>? _boundCidrs;

        /// <summary>
        /// If set, a list of
        /// CIDRs valid as the source address for login requests. This value is also encoded into any resulting token.
        /// </summary>
        public InputList<string> BoundCidrs
        {
            get => _boundCidrs ?? (_boundCidrs = new InputList<string>());
            set => _boundCidrs = value;
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

        /// <summary>
        /// Number of seconds after which issued tokens can no longer be renewed.
        /// </summary>
        [Input("explicitMaxTtl")]
        public Input<string>? ExplicitMaxTtl { get; set; }

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
        /// If set, indicates that the
        /// token generated using this role should never expire. The token should be renewed within the
        /// duration specified by this value. At each renewal, the token's TTL will be set to the
        /// value of this field. Specified in seconds.
        /// </summary>
        [Input("period")]
        public Input<string>? Period { get; set; }

        /// <summary>
        /// Wether to disable the ability of the token to be renewed past its initial TTL.
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
        /// Generated Token's Policies
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

        public AuthBackendRoleState()
        {
        }
    }
}
