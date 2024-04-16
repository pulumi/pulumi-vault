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
    /// Provides a resource to manage Role Governing Policy (RGP) via [Sentinel](https://www.vaultproject.io/docs/enterprise/sentinel/index.html).
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
    ///     var allow_all = new Vault.RgpPolicy("allow-all", new()
    ///     {
    ///         Name = "allow-all",
    ///         EnforcementLevel = "soft-mandatory",
    ///         Policy = @"main = rule {
    ///   true
    /// }
    /// ",
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// </summary>
    [VaultResourceType("vault:index/rgpPolicy:RgpPolicy")]
    public partial class RgpPolicy : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Enforcement level of Sentinel policy. Can be either `advisory` or `soft-mandatory` or `hard-mandatory`
        /// </summary>
        [Output("enforcementLevel")]
        public Output<string> EnforcementLevel { get; private set; } = null!;

        /// <summary>
        /// The name of the policy
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The namespace to provision the resource in.
        /// The value should not contain leading or trailing forward slashes.
        /// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
        /// *Available only for Vault Enterprise*.
        /// </summary>
        [Output("namespace")]
        public Output<string?> Namespace { get; private set; } = null!;

        /// <summary>
        /// String containing a Sentinel policy
        /// </summary>
        [Output("policy")]
        public Output<string> Policy { get; private set; } = null!;


        /// <summary>
        /// Create a RgpPolicy resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public RgpPolicy(string name, RgpPolicyArgs args, CustomResourceOptions? options = null)
            : base("vault:index/rgpPolicy:RgpPolicy", name, args ?? new RgpPolicyArgs(), MakeResourceOptions(options, ""))
        {
        }

        private RgpPolicy(string name, Input<string> id, RgpPolicyState? state = null, CustomResourceOptions? options = null)
            : base("vault:index/rgpPolicy:RgpPolicy", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing RgpPolicy resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static RgpPolicy Get(string name, Input<string> id, RgpPolicyState? state = null, CustomResourceOptions? options = null)
        {
            return new RgpPolicy(name, id, state, options);
        }
    }

    public sealed class RgpPolicyArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Enforcement level of Sentinel policy. Can be either `advisory` or `soft-mandatory` or `hard-mandatory`
        /// </summary>
        [Input("enforcementLevel", required: true)]
        public Input<string> EnforcementLevel { get; set; } = null!;

        /// <summary>
        /// The name of the policy
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The namespace to provision the resource in.
        /// The value should not contain leading or trailing forward slashes.
        /// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
        /// *Available only for Vault Enterprise*.
        /// </summary>
        [Input("namespace")]
        public Input<string>? Namespace { get; set; }

        /// <summary>
        /// String containing a Sentinel policy
        /// </summary>
        [Input("policy", required: true)]
        public Input<string> Policy { get; set; } = null!;

        public RgpPolicyArgs()
        {
        }
        public static new RgpPolicyArgs Empty => new RgpPolicyArgs();
    }

    public sealed class RgpPolicyState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Enforcement level of Sentinel policy. Can be either `advisory` or `soft-mandatory` or `hard-mandatory`
        /// </summary>
        [Input("enforcementLevel")]
        public Input<string>? EnforcementLevel { get; set; }

        /// <summary>
        /// The name of the policy
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The namespace to provision the resource in.
        /// The value should not contain leading or trailing forward slashes.
        /// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
        /// *Available only for Vault Enterprise*.
        /// </summary>
        [Input("namespace")]
        public Input<string>? Namespace { get; set; }

        /// <summary>
        /// String containing a Sentinel policy
        /// </summary>
        [Input("policy")]
        public Input<string>? Policy { get; set; }

        public RgpPolicyState()
        {
        }
        public static new RgpPolicyState Empty => new RgpPolicyState();
    }
}
