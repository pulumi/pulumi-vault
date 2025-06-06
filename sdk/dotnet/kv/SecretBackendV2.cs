// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vault.kv
{
    /// <summary>
    /// Configures KV-V2 backend level settings that are applied to
    /// every key in the key-value store.
    /// 
    /// For more information on Vault's KV-V2 secret backend
    /// [see here](https://www.vaultproject.io/docs/secrets/kv/kv-v2).
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
    ///     var kvv2 = new Vault.Mount("kvv2", new()
    ///     {
    ///         Path = "kvv2",
    ///         Type = "kv",
    ///         Options = 
    ///         {
    ///             { "version", "2" },
    ///         },
    ///         Description = "KV Version 2 secret engine mount",
    ///     });
    /// 
    ///     var example = new Vault.Kv.SecretBackendV2("example", new()
    ///     {
    ///         Mount = kvv2.Path,
    ///         MaxVersions = 5,
    ///         DeleteVersionAfter = 12600,
    ///         CasRequired = true,
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Required Vault Capabilities
    /// 
    /// Use of this resource requires the `create` or `update` capability
    /// (depending on whether the resource already exists) on the given path,
    /// the `delete` capability if the resource is removed from configuration,
    /// and the `read` capability for drift detection (by default).
    /// 
    /// ## Import
    /// 
    /// The KV-V2 secret backend can be imported using its unique ID,
    /// the `${mount}/config`, e.g.
    /// 
    /// ```sh
    /// $ pulumi import vault:kv/secretBackendV2:SecretBackendV2 example kvv2/config
    /// ```
    /// </summary>
    [VaultResourceType("vault:kv/secretBackendV2:SecretBackendV2")]
    public partial class SecretBackendV2 : global::Pulumi.CustomResource
    {
        /// <summary>
        /// If true, all keys will require the cas
        /// parameter to be set on all write requests.
        /// </summary>
        [Output("casRequired")]
        public Output<bool> CasRequired { get; private set; } = null!;

        /// <summary>
        /// If set, specifies the length of time before
        /// a version is deleted. Accepts duration in integer seconds.
        /// </summary>
        [Output("deleteVersionAfter")]
        public Output<int?> DeleteVersionAfter { get; private set; } = null!;

        /// <summary>
        /// The number of versions to keep per key.
        /// </summary>
        [Output("maxVersions")]
        public Output<int> MaxVersions { get; private set; } = null!;

        /// <summary>
        /// Path where KV-V2 engine is mounted.
        /// </summary>
        [Output("mount")]
        public Output<string> Mount { get; private set; } = null!;

        /// <summary>
        /// The namespace to provision the resource in.
        /// The value should not contain leading or trailing forward slashes.
        /// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
        /// *Available only for Vault Enterprise*.
        /// </summary>
        [Output("namespace")]
        public Output<string?> Namespace { get; private set; } = null!;


        /// <summary>
        /// Create a SecretBackendV2 resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public SecretBackendV2(string name, SecretBackendV2Args args, CustomResourceOptions? options = null)
            : base("vault:kv/secretBackendV2:SecretBackendV2", name, args ?? new SecretBackendV2Args(), MakeResourceOptions(options, ""))
        {
        }

        private SecretBackendV2(string name, Input<string> id, SecretBackendV2State? state = null, CustomResourceOptions? options = null)
            : base("vault:kv/secretBackendV2:SecretBackendV2", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing SecretBackendV2 resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static SecretBackendV2 Get(string name, Input<string> id, SecretBackendV2State? state = null, CustomResourceOptions? options = null)
        {
            return new SecretBackendV2(name, id, state, options);
        }
    }

    public sealed class SecretBackendV2Args : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// If true, all keys will require the cas
        /// parameter to be set on all write requests.
        /// </summary>
        [Input("casRequired")]
        public Input<bool>? CasRequired { get; set; }

        /// <summary>
        /// If set, specifies the length of time before
        /// a version is deleted. Accepts duration in integer seconds.
        /// </summary>
        [Input("deleteVersionAfter")]
        public Input<int>? DeleteVersionAfter { get; set; }

        /// <summary>
        /// The number of versions to keep per key.
        /// </summary>
        [Input("maxVersions")]
        public Input<int>? MaxVersions { get; set; }

        /// <summary>
        /// Path where KV-V2 engine is mounted.
        /// </summary>
        [Input("mount", required: true)]
        public Input<string> Mount { get; set; } = null!;

        /// <summary>
        /// The namespace to provision the resource in.
        /// The value should not contain leading or trailing forward slashes.
        /// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
        /// *Available only for Vault Enterprise*.
        /// </summary>
        [Input("namespace")]
        public Input<string>? Namespace { get; set; }

        public SecretBackendV2Args()
        {
        }
        public static new SecretBackendV2Args Empty => new SecretBackendV2Args();
    }

    public sealed class SecretBackendV2State : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// If true, all keys will require the cas
        /// parameter to be set on all write requests.
        /// </summary>
        [Input("casRequired")]
        public Input<bool>? CasRequired { get; set; }

        /// <summary>
        /// If set, specifies the length of time before
        /// a version is deleted. Accepts duration in integer seconds.
        /// </summary>
        [Input("deleteVersionAfter")]
        public Input<int>? DeleteVersionAfter { get; set; }

        /// <summary>
        /// The number of versions to keep per key.
        /// </summary>
        [Input("maxVersions")]
        public Input<int>? MaxVersions { get; set; }

        /// <summary>
        /// Path where KV-V2 engine is mounted.
        /// </summary>
        [Input("mount")]
        public Input<string>? Mount { get; set; }

        /// <summary>
        /// The namespace to provision the resource in.
        /// The value should not contain leading or trailing forward slashes.
        /// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
        /// *Available only for Vault Enterprise*.
        /// </summary>
        [Input("namespace")]
        public Input<string>? Namespace { get; set; }

        public SecretBackendV2State()
        {
        }
        public static new SecretBackendV2State Empty => new SecretBackendV2State();
    }
}
