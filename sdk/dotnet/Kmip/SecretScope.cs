// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vault.Kmip
{
    /// <summary>
    /// Manages KMIP Secret Scopes in a Vault server. This feature requires
    /// Vault Enterprise. See the [Vault documentation](https://www.vaultproject.io/docs/secrets/kmip)
    /// for more information.
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
    ///     var @default = new Vault.Kmip.SecretBackend("default", new()
    ///     {
    ///         Path = "kmip",
    ///         Description = "Vault KMIP backend",
    ///     });
    /// 
    ///     var dev = new Vault.Kmip.SecretScope("dev", new()
    ///     {
    ///         Path = @default.Path,
    ///         Scope = "dev",
    ///         Force = true,
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ## Import
    /// 
    /// KMIP Secret scope can be imported using the `path`, e.g.
    /// 
    /// ```sh
    /// $ pulumi import vault:kmip/secretScope:SecretScope dev kmip
    /// ```
    /// </summary>
    [VaultResourceType("vault:kmip/secretScope:SecretScope")]
    public partial class SecretScope : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Boolean field to force deletion even if there are managed objects in the scope.
        /// </summary>
        [Output("force")]
        public Output<bool?> Force { get; private set; } = null!;

        /// <summary>
        /// The namespace to provision the resource in.
        /// The value should not contain leading or trailing forward slashes.
        /// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
        /// *Available only for Vault Enterprise*.
        /// </summary>
        [Output("namespace")]
        public Output<string?> Namespace { get; private set; } = null!;

        /// <summary>
        /// The unique path this backend should be mounted at. Must
        /// not begin or end with a `/`. Defaults to `kmip`.
        /// </summary>
        [Output("path")]
        public Output<string> Path { get; private set; } = null!;

        /// <summary>
        /// Name of the scope.
        /// </summary>
        [Output("scope")]
        public Output<string> Scope { get; private set; } = null!;


        /// <summary>
        /// Create a SecretScope resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public SecretScope(string name, SecretScopeArgs args, CustomResourceOptions? options = null)
            : base("vault:kmip/secretScope:SecretScope", name, args ?? new SecretScopeArgs(), MakeResourceOptions(options, ""))
        {
        }

        private SecretScope(string name, Input<string> id, SecretScopeState? state = null, CustomResourceOptions? options = null)
            : base("vault:kmip/secretScope:SecretScope", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing SecretScope resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static SecretScope Get(string name, Input<string> id, SecretScopeState? state = null, CustomResourceOptions? options = null)
        {
            return new SecretScope(name, id, state, options);
        }
    }

    public sealed class SecretScopeArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Boolean field to force deletion even if there are managed objects in the scope.
        /// </summary>
        [Input("force")]
        public Input<bool>? Force { get; set; }

        /// <summary>
        /// The namespace to provision the resource in.
        /// The value should not contain leading or trailing forward slashes.
        /// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
        /// *Available only for Vault Enterprise*.
        /// </summary>
        [Input("namespace")]
        public Input<string>? Namespace { get; set; }

        /// <summary>
        /// The unique path this backend should be mounted at. Must
        /// not begin or end with a `/`. Defaults to `kmip`.
        /// </summary>
        [Input("path", required: true)]
        public Input<string> Path { get; set; } = null!;

        /// <summary>
        /// Name of the scope.
        /// </summary>
        [Input("scope", required: true)]
        public Input<string> Scope { get; set; } = null!;

        public SecretScopeArgs()
        {
        }
        public static new SecretScopeArgs Empty => new SecretScopeArgs();
    }

    public sealed class SecretScopeState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Boolean field to force deletion even if there are managed objects in the scope.
        /// </summary>
        [Input("force")]
        public Input<bool>? Force { get; set; }

        /// <summary>
        /// The namespace to provision the resource in.
        /// The value should not contain leading or trailing forward slashes.
        /// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
        /// *Available only for Vault Enterprise*.
        /// </summary>
        [Input("namespace")]
        public Input<string>? Namespace { get; set; }

        /// <summary>
        /// The unique path this backend should be mounted at. Must
        /// not begin or end with a `/`. Defaults to `kmip`.
        /// </summary>
        [Input("path")]
        public Input<string>? Path { get; set; }

        /// <summary>
        /// Name of the scope.
        /// </summary>
        [Input("scope")]
        public Input<string>? Scope { get; set; }

        public SecretScopeState()
        {
        }
        public static new SecretScopeState Empty => new SecretScopeState();
    }
}
