// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vault
{
    /// <summary>
    /// Autopilot enables automated workflows for managing Raft clusters. The
    /// current feature set includes 3 main features: Server Stabilization, Dead
    /// Server Cleanup and State API. **These three features are introduced in
    /// Vault 1.7.**
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using Pulumi;
    /// using Vault = Pulumi.Vault;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var autopilot = new Vault.RaftAutopilot("autopilot", new()
    ///     {
    ///         CleanupDeadServers = true,
    ///         DeadServerLastContactThreshold = "24h0m0s",
    ///         LastContactThreshold = "10s",
    ///         MaxTrailingLogs = 1000,
    ///         MinQuorum = 3,
    ///         ServerStabilizationTime = "10s",
    ///     });
    /// 
    /// });
    /// ```
    /// </summary>
    [VaultResourceType("vault:index/raftAutopilot:RaftAutopilot")]
    public partial class RaftAutopilot : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Specifies whether to remove dead server nodes
        /// periodically or when a new server joins. This requires that `min-quorum` is also set.
        /// </summary>
        [Output("cleanupDeadServers")]
        public Output<bool?> CleanupDeadServers { get; private set; } = null!;

        /// <summary>
        /// Limit the amount of time a 
        /// server can go without leader contact before being considered failed. This only takes
        /// effect when `cleanup_dead_servers` is set.
        /// </summary>
        [Output("deadServerLastContactThreshold")]
        public Output<string?> DeadServerLastContactThreshold { get; private set; } = null!;

        /// <summary>
        /// Disables automatically upgrading Vault using autopilot. (Enterprise-only)
        /// </summary>
        [Output("disableUpgradeMigration")]
        public Output<bool?> DisableUpgradeMigration { get; private set; } = null!;

        /// <summary>
        /// Limit the amount of time a server can go 
        /// without leader contact before being considered unhealthy.
        /// </summary>
        [Output("lastContactThreshold")]
        public Output<string?> LastContactThreshold { get; private set; } = null!;

        /// <summary>
        /// Maximum number of log entries in the Raft log 
        /// that a server can be behind its leader before being considered unhealthy.
        /// </summary>
        [Output("maxTrailingLogs")]
        public Output<int?> MaxTrailingLogs { get; private set; } = null!;

        /// <summary>
        /// Minimum number of servers allowed in a cluster before 
        /// autopilot can prune dead servers. This should at least be 3. Applicable only for
        /// voting nodes.
        /// </summary>
        [Output("minQuorum")]
        public Output<int?> MinQuorum { get; private set; } = null!;

        /// <summary>
        /// The namespace to provision the resource in.
        /// The value should not contain leading or trailing forward slashes.
        /// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
        /// *Available only for Vault Enterprise*.
        /// </summary>
        [Output("namespace")]
        public Output<string?> Namespace { get; private set; } = null!;

        /// <summary>
        /// Minimum amount of time a server must be 
        /// stable in the 'healthy' state before being added to the cluster.
        /// </summary>
        [Output("serverStabilizationTime")]
        public Output<string?> ServerStabilizationTime { get; private set; } = null!;


        /// <summary>
        /// Create a RaftAutopilot resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public RaftAutopilot(string name, RaftAutopilotArgs? args = null, CustomResourceOptions? options = null)
            : base("vault:index/raftAutopilot:RaftAutopilot", name, args ?? new RaftAutopilotArgs(), MakeResourceOptions(options, ""))
        {
        }

        private RaftAutopilot(string name, Input<string> id, RaftAutopilotState? state = null, CustomResourceOptions? options = null)
            : base("vault:index/raftAutopilot:RaftAutopilot", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing RaftAutopilot resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static RaftAutopilot Get(string name, Input<string> id, RaftAutopilotState? state = null, CustomResourceOptions? options = null)
        {
            return new RaftAutopilot(name, id, state, options);
        }
    }

    public sealed class RaftAutopilotArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies whether to remove dead server nodes
        /// periodically or when a new server joins. This requires that `min-quorum` is also set.
        /// </summary>
        [Input("cleanupDeadServers")]
        public Input<bool>? CleanupDeadServers { get; set; }

        /// <summary>
        /// Limit the amount of time a 
        /// server can go without leader contact before being considered failed. This only takes
        /// effect when `cleanup_dead_servers` is set.
        /// </summary>
        [Input("deadServerLastContactThreshold")]
        public Input<string>? DeadServerLastContactThreshold { get; set; }

        /// <summary>
        /// Disables automatically upgrading Vault using autopilot. (Enterprise-only)
        /// </summary>
        [Input("disableUpgradeMigration")]
        public Input<bool>? DisableUpgradeMigration { get; set; }

        /// <summary>
        /// Limit the amount of time a server can go 
        /// without leader contact before being considered unhealthy.
        /// </summary>
        [Input("lastContactThreshold")]
        public Input<string>? LastContactThreshold { get; set; }

        /// <summary>
        /// Maximum number of log entries in the Raft log 
        /// that a server can be behind its leader before being considered unhealthy.
        /// </summary>
        [Input("maxTrailingLogs")]
        public Input<int>? MaxTrailingLogs { get; set; }

        /// <summary>
        /// Minimum number of servers allowed in a cluster before 
        /// autopilot can prune dead servers. This should at least be 3. Applicable only for
        /// voting nodes.
        /// </summary>
        [Input("minQuorum")]
        public Input<int>? MinQuorum { get; set; }

        /// <summary>
        /// The namespace to provision the resource in.
        /// The value should not contain leading or trailing forward slashes.
        /// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
        /// *Available only for Vault Enterprise*.
        /// </summary>
        [Input("namespace")]
        public Input<string>? Namespace { get; set; }

        /// <summary>
        /// Minimum amount of time a server must be 
        /// stable in the 'healthy' state before being added to the cluster.
        /// </summary>
        [Input("serverStabilizationTime")]
        public Input<string>? ServerStabilizationTime { get; set; }

        public RaftAutopilotArgs()
        {
        }
        public static new RaftAutopilotArgs Empty => new RaftAutopilotArgs();
    }

    public sealed class RaftAutopilotState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies whether to remove dead server nodes
        /// periodically or when a new server joins. This requires that `min-quorum` is also set.
        /// </summary>
        [Input("cleanupDeadServers")]
        public Input<bool>? CleanupDeadServers { get; set; }

        /// <summary>
        /// Limit the amount of time a 
        /// server can go without leader contact before being considered failed. This only takes
        /// effect when `cleanup_dead_servers` is set.
        /// </summary>
        [Input("deadServerLastContactThreshold")]
        public Input<string>? DeadServerLastContactThreshold { get; set; }

        /// <summary>
        /// Disables automatically upgrading Vault using autopilot. (Enterprise-only)
        /// </summary>
        [Input("disableUpgradeMigration")]
        public Input<bool>? DisableUpgradeMigration { get; set; }

        /// <summary>
        /// Limit the amount of time a server can go 
        /// without leader contact before being considered unhealthy.
        /// </summary>
        [Input("lastContactThreshold")]
        public Input<string>? LastContactThreshold { get; set; }

        /// <summary>
        /// Maximum number of log entries in the Raft log 
        /// that a server can be behind its leader before being considered unhealthy.
        /// </summary>
        [Input("maxTrailingLogs")]
        public Input<int>? MaxTrailingLogs { get; set; }

        /// <summary>
        /// Minimum number of servers allowed in a cluster before 
        /// autopilot can prune dead servers. This should at least be 3. Applicable only for
        /// voting nodes.
        /// </summary>
        [Input("minQuorum")]
        public Input<int>? MinQuorum { get; set; }

        /// <summary>
        /// The namespace to provision the resource in.
        /// The value should not contain leading or trailing forward slashes.
        /// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
        /// *Available only for Vault Enterprise*.
        /// </summary>
        [Input("namespace")]
        public Input<string>? Namespace { get; set; }

        /// <summary>
        /// Minimum amount of time a server must be 
        /// stable in the 'healthy' state before being added to the cluster.
        /// </summary>
        [Input("serverStabilizationTime")]
        public Input<string>? ServerStabilizationTime { get; set; }

        public RaftAutopilotState()
        {
        }
        public static new RaftAutopilotState Empty => new RaftAutopilotState();
    }
}
