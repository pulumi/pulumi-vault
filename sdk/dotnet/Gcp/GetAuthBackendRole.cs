// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi.Utilities;

namespace Pulumi.Vault.Gcp
{
    public static class GetAuthBackendRole
    {
        /// <summary>
        /// Reads a GCP auth role from a Vault server.
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
        ///         var role = Output.Create(Vault.Gcp.GetAuthBackendRole.InvokeAsync(new Vault.Gcp.GetAuthBackendRoleArgs
        ///         {
        ///             Backend = "my-gcp-backend",
        ///             RoleName = "my-role",
        ///         }));
        ///         this.Role_id = role.Apply(role =&gt; role.RoleId);
        ///     }
        /// 
        ///     [Output("role-id")]
        ///     public Output&lt;string&gt; Role_id { get; set; }
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetAuthBackendRoleResult> InvokeAsync(GetAuthBackendRoleArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetAuthBackendRoleResult>("vault:gcp/getAuthBackendRole:getAuthBackendRole", args ?? new GetAuthBackendRoleArgs(), options.WithVersion());

        /// <summary>
        /// Reads a GCP auth role from a Vault server.
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
        ///         var role = Output.Create(Vault.Gcp.GetAuthBackendRole.InvokeAsync(new Vault.Gcp.GetAuthBackendRoleArgs
        ///         {
        ///             Backend = "my-gcp-backend",
        ///             RoleName = "my-role",
        ///         }));
        ///         this.Role_id = role.Apply(role =&gt; role.RoleId);
        ///     }
        /// 
        ///     [Output("role-id")]
        ///     public Output&lt;string&gt; Role_id { get; set; }
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetAuthBackendRoleResult> Invoke(GetAuthBackendRoleInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetAuthBackendRoleResult>("vault:gcp/getAuthBackendRole:getAuthBackendRole", args ?? new GetAuthBackendRoleInvokeArgs(), options.WithVersion());
    }


    public sealed class GetAuthBackendRoleArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The unique name for the GCP backend from which to fetch the role. Defaults to "gcp".
        /// </summary>
        [Input("backend")]
        public string? Backend { get; set; }

        /// <summary>
        /// The name of the role to retrieve the Role ID for.
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

        public GetAuthBackendRoleArgs()
        {
        }
    }

    public sealed class GetAuthBackendRoleInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The unique name for the GCP backend from which to fetch the role. Defaults to "gcp".
        /// </summary>
        [Input("backend")]
        public Input<string>? Backend { get; set; }

        /// <summary>
        /// The name of the role to retrieve the Role ID for.
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
        /// (Optional) If set, indicates that the
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

        public GetAuthBackendRoleInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetAuthBackendRoleResult
    {
        public readonly string? Backend;
        /// <summary>
        /// GCP regions bound to the role. Returned when `type` is `gce`.
        /// </summary>
        public readonly ImmutableArray<string> BoundInstanceGroups;
        /// <summary>
        /// GCP labels bound to the role. Returned when `type` is `gce`.
        /// </summary>
        public readonly ImmutableArray<string> BoundLabels;
        /// <summary>
        /// GCP projects bound to the role.
        /// </summary>
        public readonly ImmutableArray<string> BoundProjects;
        /// <summary>
        /// GCP regions bound to the role. Returned when `type` is `gce`.
        /// </summary>
        public readonly ImmutableArray<string> BoundRegions;
        /// <summary>
        /// GCP service accounts bound to the role. Returned when `type` is `iam`.
        /// </summary>
        public readonly ImmutableArray<string> BoundServiceAccounts;
        /// <summary>
        /// GCP zones bound to the role. Returned when `type` is `gce`.
        /// </summary>
        public readonly ImmutableArray<string> BoundZones;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The RoleID of the GCP role.
        /// </summary>
        public readonly string RoleId;
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
        /// <summary>
        /// Type of GCP role. Expected values are `iam` or `gce`.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetAuthBackendRoleResult(
            string? backend,

            ImmutableArray<string> boundInstanceGroups,

            ImmutableArray<string> boundLabels,

            ImmutableArray<string> boundProjects,

            ImmutableArray<string> boundRegions,

            ImmutableArray<string> boundServiceAccounts,

            ImmutableArray<string> boundZones,

            string id,

            string roleId,

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

            string type)
        {
            Backend = backend;
            BoundInstanceGroups = boundInstanceGroups;
            BoundLabels = boundLabels;
            BoundProjects = boundProjects;
            BoundRegions = boundRegions;
            BoundServiceAccounts = boundServiceAccounts;
            BoundZones = boundZones;
            Id = id;
            RoleId = roleId;
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
            Type = type;
        }
    }
}
