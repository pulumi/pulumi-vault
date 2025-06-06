// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vault
{
    /// <summary>
    /// ## Import
    /// 
    /// Namespaces can be imported using its `name` as accessor id
    /// 
    /// ```sh
    /// $ pulumi import vault:index/namespace:Namespace example &lt;name&gt;
    /// ```
    /// 
    /// If the declared resource is imported and intends to support namespaces using a provider alias, then the name is relative to the namespace path.
    /// 
    /// hcl
    /// 
    /// provider "vault" {
    /// 
    /// # Configuration options
    /// 
    ///   namespace = "example"
    /// 
    ///   alias     = "example"
    /// 
    /// }
    /// 
    /// resource "vault_namespace" "example2" {
    /// 
    ///   provider = vault.example
    /// 
    ///   path     = "example2"
    /// 
    /// }
    /// 
    /// ```sh
    /// $ pulumi import vault:index/namespace:Namespace example2 example2
    /// ```
    /// 
    /// $ terraform state show vault_namespace.example2
    /// 
    /// vault_namespace.example2:
    /// 
    /// resource "vault_namespace" "example2" {
    /// 
    ///     id           = "example/example2/"
    ///     
    ///     namespace_id = &lt;known after import&gt;
    ///     
    ///     path         = "example2"
    ///     
    ///     path_fq      = "example2"
    /// 
    /// }
    /// </summary>
    [VaultResourceType("vault:index/namespace:Namespace")]
    public partial class Namespace : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Custom metadata describing this namespace. Value type
        /// is `map[string]string`. Requires Vault version 1.12+.
        /// </summary>
        [Output("customMetadata")]
        public Output<ImmutableDictionary<string, string>> CustomMetadata { get; private set; } = null!;

        /// <summary>
        /// The namespace to provision the resource in.
        /// The value should not contain leading or trailing forward slashes.
        /// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
        /// *Available only for Vault Enterprise*.
        /// </summary>
        [Output("namespace")]
        public Output<string?> TargetNamespace { get; private set; } = null!;

        /// <summary>
        /// Vault server's internal ID of the namespace.
        /// </summary>
        [Output("namespaceId")]
        public Output<string> NamespaceId { get; private set; } = null!;

        /// <summary>
        /// The path of the namespace. Must not have a trailing `/`.
        /// </summary>
        [Output("path")]
        public Output<string> Path { get; private set; } = null!;

        /// <summary>
        /// The fully qualified path to the namespace. Useful when provisioning resources in a child `namespace`.
        /// The path is relative to the provider's `namespace` argument.
        /// </summary>
        [Output("pathFq")]
        public Output<string> PathFq { get; private set; } = null!;


        /// <summary>
        /// Create a Namespace resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Namespace(string name, NamespaceArgs args, CustomResourceOptions? options = null)
            : base("vault:index/namespace:Namespace", name, args ?? new NamespaceArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Namespace(string name, Input<string> id, NamespaceState? state = null, CustomResourceOptions? options = null)
            : base("vault:index/namespace:Namespace", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Namespace resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Namespace Get(string name, Input<string> id, NamespaceState? state = null, CustomResourceOptions? options = null)
        {
            return new Namespace(name, id, state, options);
        }
    }

    public sealed class NamespaceArgs : global::Pulumi.ResourceArgs
    {
        [Input("customMetadata")]
        private InputMap<string>? _customMetadata;

        /// <summary>
        /// Custom metadata describing this namespace. Value type
        /// is `map[string]string`. Requires Vault version 1.12+.
        /// </summary>
        public InputMap<string> CustomMetadata
        {
            get => _customMetadata ?? (_customMetadata = new InputMap<string>());
            set => _customMetadata = value;
        }

        /// <summary>
        /// The namespace to provision the resource in.
        /// The value should not contain leading or trailing forward slashes.
        /// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
        /// *Available only for Vault Enterprise*.
        /// </summary>
        [Input("namespace")]
        public Input<string>? TargetNamespace { get; set; }

        /// <summary>
        /// The path of the namespace. Must not have a trailing `/`.
        /// </summary>
        [Input("path", required: true)]
        public Input<string> Path { get; set; } = null!;

        /// <summary>
        /// The fully qualified path to the namespace. Useful when provisioning resources in a child `namespace`.
        /// The path is relative to the provider's `namespace` argument.
        /// </summary>
        [Input("pathFq")]
        public Input<string>? PathFq { get; set; }

        public NamespaceArgs()
        {
        }
        public static new NamespaceArgs Empty => new NamespaceArgs();
    }

    public sealed class NamespaceState : global::Pulumi.ResourceArgs
    {
        [Input("customMetadata")]
        private InputMap<string>? _customMetadata;

        /// <summary>
        /// Custom metadata describing this namespace. Value type
        /// is `map[string]string`. Requires Vault version 1.12+.
        /// </summary>
        public InputMap<string> CustomMetadata
        {
            get => _customMetadata ?? (_customMetadata = new InputMap<string>());
            set => _customMetadata = value;
        }

        /// <summary>
        /// The namespace to provision the resource in.
        /// The value should not contain leading or trailing forward slashes.
        /// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
        /// *Available only for Vault Enterprise*.
        /// </summary>
        [Input("namespace")]
        public Input<string>? TargetNamespace { get; set; }

        /// <summary>
        /// Vault server's internal ID of the namespace.
        /// </summary>
        [Input("namespaceId")]
        public Input<string>? NamespaceId { get; set; }

        /// <summary>
        /// The path of the namespace. Must not have a trailing `/`.
        /// </summary>
        [Input("path")]
        public Input<string>? Path { get; set; }

        /// <summary>
        /// The fully qualified path to the namespace. Useful when provisioning resources in a child `namespace`.
        /// The path is relative to the provider's `namespace` argument.
        /// </summary>
        [Input("pathFq")]
        public Input<string>? PathFq { get; set; }

        public NamespaceState()
        {
        }
        public static new NamespaceState Empty => new NamespaceState();
    }
}
