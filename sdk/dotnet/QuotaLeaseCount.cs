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
    /// Manage lease count quotas which enforce the number of leases that can be created.
    /// A lease count quota can be created at the root level or defined on a namespace or mount by
    /// specifying a path when creating the quota.
    /// 
    /// See [Vault's Documentation](https://www.vaultproject.io/docs/enterprise/lease-count-quotas) for more
    /// information.
    /// 
    /// **Note** this feature is available only with Vault Enterprise.
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
    ///     var @global = new Vault.QuotaLeaseCount("global", new()
    ///     {
    ///         MaxLeases = 100,
    ///         Path = "",
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ## Import
    /// 
    /// Lease count quotas can be imported using their names
    /// 
    /// ```sh
    /// $ pulumi import vault:index/quotaLeaseCount:QuotaLeaseCount global global
    /// ```
    /// </summary>
    [VaultResourceType("vault:index/quotaLeaseCount:QuotaLeaseCount")]
    public partial class QuotaLeaseCount : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The maximum number of leases to be allowed by the quota
        /// rule. The `max_leases` must be positive.
        /// </summary>
        [Output("maxLeases")]
        public Output<int> MaxLeases { get; private set; } = null!;

        /// <summary>
        /// Name of the rate limit quota
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The namespace to provision the resource in.
        /// The value should not contain leading or trailing forward slashes.
        /// The `namespace` is always relative to the provider's configured namespace.
        /// *Available only for Vault Enterprise*.
        /// </summary>
        [Output("namespace")]
        public Output<string?> Namespace { get; private set; } = null!;

        /// <summary>
        /// Path of the mount or namespace to apply the quota. A blank path configures a
        /// global rate limit quota. For example `namespace1/` adds a quota to a full namespace,
        /// `namespace1/auth/userpass` adds a `quota` to `userpass` in `namespace1`.
        /// Updating this field on an existing quota can have "moving" effects. For example, updating
        /// `auth/userpass` to `namespace1/auth/userpass` moves this quota from being a global mount quota to
        /// a namespace specific mount quota. **Note, namespaces are supported in Enterprise only.**
        /// </summary>
        [Output("path")]
        public Output<string?> Path { get; private set; } = null!;

        /// <summary>
        /// If set on a quota where `path` is set to an auth mount with a concept of roles (such as /auth/approle/), this will make the quota restrict login requests to that mount that are made with the specified role.
        /// </summary>
        [Output("role")]
        public Output<string?> Role { get; private set; } = null!;


        /// <summary>
        /// Create a QuotaLeaseCount resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public QuotaLeaseCount(string name, QuotaLeaseCountArgs args, CustomResourceOptions? options = null)
            : base("vault:index/quotaLeaseCount:QuotaLeaseCount", name, args ?? new QuotaLeaseCountArgs(), MakeResourceOptions(options, ""))
        {
        }

        private QuotaLeaseCount(string name, Input<string> id, QuotaLeaseCountState? state = null, CustomResourceOptions? options = null)
            : base("vault:index/quotaLeaseCount:QuotaLeaseCount", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing QuotaLeaseCount resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static QuotaLeaseCount Get(string name, Input<string> id, QuotaLeaseCountState? state = null, CustomResourceOptions? options = null)
        {
            return new QuotaLeaseCount(name, id, state, options);
        }
    }

    public sealed class QuotaLeaseCountArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The maximum number of leases to be allowed by the quota
        /// rule. The `max_leases` must be positive.
        /// </summary>
        [Input("maxLeases", required: true)]
        public Input<int> MaxLeases { get; set; } = null!;

        /// <summary>
        /// Name of the rate limit quota
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The namespace to provision the resource in.
        /// The value should not contain leading or trailing forward slashes.
        /// The `namespace` is always relative to the provider's configured namespace.
        /// *Available only for Vault Enterprise*.
        /// </summary>
        [Input("namespace")]
        public Input<string>? Namespace { get; set; }

        /// <summary>
        /// Path of the mount or namespace to apply the quota. A blank path configures a
        /// global rate limit quota. For example `namespace1/` adds a quota to a full namespace,
        /// `namespace1/auth/userpass` adds a `quota` to `userpass` in `namespace1`.
        /// Updating this field on an existing quota can have "moving" effects. For example, updating
        /// `auth/userpass` to `namespace1/auth/userpass` moves this quota from being a global mount quota to
        /// a namespace specific mount quota. **Note, namespaces are supported in Enterprise only.**
        /// </summary>
        [Input("path")]
        public Input<string>? Path { get; set; }

        /// <summary>
        /// If set on a quota where `path` is set to an auth mount with a concept of roles (such as /auth/approle/), this will make the quota restrict login requests to that mount that are made with the specified role.
        /// </summary>
        [Input("role")]
        public Input<string>? Role { get; set; }

        public QuotaLeaseCountArgs()
        {
        }
        public static new QuotaLeaseCountArgs Empty => new QuotaLeaseCountArgs();
    }

    public sealed class QuotaLeaseCountState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The maximum number of leases to be allowed by the quota
        /// rule. The `max_leases` must be positive.
        /// </summary>
        [Input("maxLeases")]
        public Input<int>? MaxLeases { get; set; }

        /// <summary>
        /// Name of the rate limit quota
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The namespace to provision the resource in.
        /// The value should not contain leading or trailing forward slashes.
        /// The `namespace` is always relative to the provider's configured namespace.
        /// *Available only for Vault Enterprise*.
        /// </summary>
        [Input("namespace")]
        public Input<string>? Namespace { get; set; }

        /// <summary>
        /// Path of the mount or namespace to apply the quota. A blank path configures a
        /// global rate limit quota. For example `namespace1/` adds a quota to a full namespace,
        /// `namespace1/auth/userpass` adds a `quota` to `userpass` in `namespace1`.
        /// Updating this field on an existing quota can have "moving" effects. For example, updating
        /// `auth/userpass` to `namespace1/auth/userpass` moves this quota from being a global mount quota to
        /// a namespace specific mount quota. **Note, namespaces are supported in Enterprise only.**
        /// </summary>
        [Input("path")]
        public Input<string>? Path { get; set; }

        /// <summary>
        /// If set on a quota where `path` is set to an auth mount with a concept of roles (such as /auth/approle/), this will make the quota restrict login requests to that mount that are made with the specified role.
        /// </summary>
        [Input("role")]
        public Input<string>? Role { get; set; }

        public QuotaLeaseCountState()
        {
        }
        public static new QuotaLeaseCountState Empty => new QuotaLeaseCountState();
    }
}
