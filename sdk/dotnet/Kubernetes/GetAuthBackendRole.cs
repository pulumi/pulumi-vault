// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vault.Kubernetes
{
    public static class GetAuthBackendRole
    {
        /// <summary>
        /// Reads the Role of an Kubernetes from a Vault server. See the [Vault
        /// documentation](https://www.vaultproject.io/api-docs/auth/kubernetes#read-role) for more
        /// information.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using Vault = Pulumi.Vault;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var role = Output.Create(Vault.Kubernetes.GetAuthBackendRole.InvokeAsync(new Vault.Kubernetes.GetAuthBackendRoleArgs
        ///         {
        ///             Backend = "my-kubernetes-backend",
        ///             RoleName = "my-role",
        ///         }));
        ///         this.Policies = role.Apply(role =&gt; role.Policies);
        ///     }
        /// 
        ///     [Output("policies")]
        ///     public Output&lt;string&gt; Policies { get; set; }
        /// }
        /// ```
        /// 
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetAuthBackendRoleResult> InvokeAsync(GetAuthBackendRoleArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetAuthBackendRoleResult>("vault:kubernetes/getAuthBackendRole:getAuthBackendRole", args ?? new GetAuthBackendRoleArgs(), options.WithVersion());
    }


    public sealed class GetAuthBackendRoleArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// (Optional) Audience claim to verify in the JWT.
        /// </summary>
        [Input("audience")]
        public string? Audience { get; set; }

        /// <summary>
        /// The unique name for the Kubernetes backend the role to
        /// retrieve Role attributes for resides in. Defaults to "kubernetes".
        /// </summary>
        [Input("backend")]
        public string? Backend { get; set; }

        [Input("boundCidrs")]
        private List<string>? _boundCidrs;
        [Obsolete(@"use `token_bound_cidrs` instead if you are running Vault >= 1.2")]
        public List<string> BoundCidrs
        {
            get => _boundCidrs ?? (_boundCidrs = new List<string>());
            set => _boundCidrs = value;
        }

        [Input("maxTtl")]
        public int? MaxTtl { get; set; }

        [Input("numUses")]
        public int? NumUses { get; set; }

        [Input("period")]
        public int? Period { get; set; }

        [Input("policies")]
        private List<string>? _policies;
        [Obsolete(@"use `token_policies` instead if you are running Vault >= 1.2")]
        public List<string> Policies
        {
            get => _policies ?? (_policies = new List<string>());
            set => _policies = value;
        }

        /// <summary>
        /// The name of the role to retrieve the Role attributes for.
        /// </summary>
        [Input("roleName", required: true)]
        public string RoleName { get; set; } = null!;

        [Input("tokenBoundCidrs")]
        private List<string>? _tokenBoundCidrs;

        /// <summary>
        /// List of CIDR blocks; if set, specifies blocks of IP
        /// addresses which can authenticate successfully, and ties the resulting token to these blocks
        /// as well.
        /// </summary>
        public List<string> TokenBoundCidrs
        {
            get => _tokenBoundCidrs ?? (_tokenBoundCidrs = new List<string>());
            set => _tokenBoundCidrs = value;
        }

        /// <summary>
        /// If set, will encode an
        /// [explicit max TTL](https://www.vaultproject.io/docs/concepts/tokens.html#token-time-to-live-periodic-tokens-and-explicit-max-ttls)
        /// onto the token in number of seconds. This is a hard cap even if `token_ttl` and
        /// `token_max_ttl` would otherwise allow a renewal.
        /// </summary>
        [Input("tokenExplicitMaxTtl")]
        public int? TokenExplicitMaxTtl { get; set; }

        /// <summary>
        /// The maximum lifetime for generated tokens in number of seconds.
        /// Its current value will be referenced at renewal time.
        /// </summary>
        [Input("tokenMaxTtl")]
        public int? TokenMaxTtl { get; set; }

        /// <summary>
        /// If set, the default policy will not be set on
        /// generated tokens; otherwise it will be added to the policies set in token_policies.
        /// </summary>
        [Input("tokenNoDefaultPolicy")]
        public bool? TokenNoDefaultPolicy { get; set; }

        /// <summary>
        /// The
        /// [period](https://www.vaultproject.io/docs/concepts/tokens.html#token-time-to-live-periodic-tokens-and-explicit-max-ttls),
        /// if any, in number of seconds to set on the token.
        /// </summary>
        [Input("tokenNumUses")]
        public int? TokenNumUses { get; set; }

        /// <summary>
        /// (Optional) If set, indicates that the
        /// token generated using this role should never expire. The token should be renewed within the
        /// duration specified by this value. At each renewal, the token's TTL will be set to the
        /// value of this field. Specified in seconds.
        /// </summary>
        [Input("tokenPeriod")]
        public int? TokenPeriod { get; set; }

        [Input("tokenPolicies")]
        private List<string>? _tokenPolicies;

        /// <summary>
        /// List of policies to encode onto generated tokens. Depending
        /// on the auth method, this list may be supplemented by user/group/other values.
        /// </summary>
        public List<string> TokenPolicies
        {
            get => _tokenPolicies ?? (_tokenPolicies = new List<string>());
            set => _tokenPolicies = value;
        }

        /// <summary>
        /// The incremental lifetime for generated tokens in number of seconds.
        /// Its current value will be referenced at renewal time.
        /// </summary>
        [Input("tokenTtl")]
        public int? TokenTtl { get; set; }

        /// <summary>
        /// The type of token that should be generated. Can be `service`,
        /// `batch`, or `default` to use the mount's tuned default (which unless changed will be
        /// `service` tokens). For token store roles, there are two additional possibilities:
        /// `default-service` and `default-batch` which specify the type to return unless the client
        /// requests a different type at generation time.
        /// </summary>
        [Input("tokenType")]
        public string? TokenType { get; set; }

        [Input("ttl")]
        public int? Ttl { get; set; }

        public GetAuthBackendRoleArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetAuthBackendRoleResult
    {
        /// <summary>
        /// (Optional) Audience claim to verify in the JWT.
        /// </summary>
        public readonly string? Audience;
        public readonly string? Backend;
        public readonly ImmutableArray<string> BoundCidrs;
        /// <summary>
        /// List of service account names able to access this role. If set to "*" all names are allowed, both this and bound_service_account_namespaces can not be "*".
        /// </summary>
        public readonly ImmutableArray<string> BoundServiceAccountNames;
        /// <summary>
        /// List of namespaces allowed to access this role. If set to "*" all namespaces are allowed, both this and bound_service_account_names can not be set to "*".
        /// </summary>
        public readonly ImmutableArray<string> BoundServiceAccountNamespaces;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly int? MaxTtl;
        public readonly int? NumUses;
        public readonly int? Period;
        public readonly ImmutableArray<string> Policies;
        public readonly string RoleName;
        /// <summary>
        /// List of CIDR blocks; if set, specifies blocks of IP
        /// addresses which can authenticate successfully, and ties the resulting token to these blocks
        /// as well.
        /// </summary>
        public readonly ImmutableArray<string> TokenBoundCidrs;
        /// <summary>
        /// If set, will encode an
        /// [explicit max TTL](https://www.vaultproject.io/docs/concepts/tokens.html#token-time-to-live-periodic-tokens-and-explicit-max-ttls)
        /// onto the token in number of seconds. This is a hard cap even if `token_ttl` and
        /// `token_max_ttl` would otherwise allow a renewal.
        /// </summary>
        public readonly int? TokenExplicitMaxTtl;
        /// <summary>
        /// The maximum lifetime for generated tokens in number of seconds.
        /// Its current value will be referenced at renewal time.
        /// </summary>
        public readonly int? TokenMaxTtl;
        /// <summary>
        /// If set, the default policy will not be set on
        /// generated tokens; otherwise it will be added to the policies set in token_policies.
        /// </summary>
        public readonly bool? TokenNoDefaultPolicy;
        /// <summary>
        /// The
        /// [period](https://www.vaultproject.io/docs/concepts/tokens.html#token-time-to-live-periodic-tokens-and-explicit-max-ttls),
        /// if any, in number of seconds to set on the token.
        /// </summary>
        public readonly int? TokenNumUses;
        /// <summary>
        /// (Optional) If set, indicates that the
        /// token generated using this role should never expire. The token should be renewed within the
        /// duration specified by this value. At each renewal, the token's TTL will be set to the
        /// value of this field. Specified in seconds.
        /// </summary>
        public readonly int? TokenPeriod;
        /// <summary>
        /// List of policies to encode onto generated tokens. Depending
        /// on the auth method, this list may be supplemented by user/group/other values.
        /// </summary>
        public readonly ImmutableArray<string> TokenPolicies;
        /// <summary>
        /// The incremental lifetime for generated tokens in number of seconds.
        /// Its current value will be referenced at renewal time.
        /// </summary>
        public readonly int? TokenTtl;
        /// <summary>
        /// The type of token that should be generated. Can be `service`,
        /// `batch`, or `default` to use the mount's tuned default (which unless changed will be
        /// `service` tokens). For token store roles, there are two additional possibilities:
        /// `default-service` and `default-batch` which specify the type to return unless the client
        /// requests a different type at generation time.
        /// </summary>
        public readonly string? TokenType;
        public readonly int? Ttl;

        [OutputConstructor]
        private GetAuthBackendRoleResult(
            string? audience,

            string? backend,

            ImmutableArray<string> boundCidrs,

            ImmutableArray<string> boundServiceAccountNames,

            ImmutableArray<string> boundServiceAccountNamespaces,

            string id,

            int? maxTtl,

            int? numUses,

            int? period,

            ImmutableArray<string> policies,

            string roleName,

            ImmutableArray<string> tokenBoundCidrs,

            int? tokenExplicitMaxTtl,

            int? tokenMaxTtl,

            bool? tokenNoDefaultPolicy,

            int? tokenNumUses,

            int? tokenPeriod,

            ImmutableArray<string> tokenPolicies,

            int? tokenTtl,

            string? tokenType,

            int? ttl)
        {
            Audience = audience;
            Backend = backend;
            BoundCidrs = boundCidrs;
            BoundServiceAccountNames = boundServiceAccountNames;
            BoundServiceAccountNamespaces = boundServiceAccountNamespaces;
            Id = id;
            MaxTtl = maxTtl;
            NumUses = numUses;
            Period = period;
            Policies = policies;
            RoleName = roleName;
            TokenBoundCidrs = tokenBoundCidrs;
            TokenExplicitMaxTtl = tokenExplicitMaxTtl;
            TokenMaxTtl = tokenMaxTtl;
            TokenNoDefaultPolicy = tokenNoDefaultPolicy;
            TokenNumUses = tokenNumUses;
            TokenPeriod = tokenPeriod;
            TokenPolicies = tokenPolicies;
            TokenTtl = tokenTtl;
            TokenType = tokenType;
            Ttl = ttl;
        }
    }
}
