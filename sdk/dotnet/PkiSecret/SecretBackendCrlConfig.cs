// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vault.PkiSecret
{
    /// <summary>
    /// Allows setting the duration for which the generated CRL should be marked valid. If the CRL is disabled, it will return a signed but zero-length CRL for any request. If enabled, it will re-build the CRL.
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Vault = Pulumi.Vault;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var pki = new Vault.Mount("pki", new Vault.MountArgs
    ///         {
    ///             DefaultLeaseTtlSeconds = 3600,
    ///             MaxLeaseTtlSeconds = 86400,
    ///             Path = "%s",
    ///             Type = "pki",
    ///         });
    ///         var crlConfig = new Vault.PkiSecret.SecretBackendCrlConfig("crlConfig", new Vault.PkiSecret.SecretBackendCrlConfigArgs
    ///         {
    ///             Backend = pki.Path,
    ///             Disable = false,
    ///             Expiry = "72h",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// </summary>
    public partial class SecretBackendCrlConfig : Pulumi.CustomResource
    {
        /// <summary>
        /// The path the PKI secret backend is mounted at, with no leading or trailing `/`s.
        /// </summary>
        [Output("backend")]
        public Output<string> Backend { get; private set; } = null!;

        /// <summary>
        /// Disables or enables CRL building.
        /// </summary>
        [Output("disable")]
        public Output<bool?> Disable { get; private set; } = null!;

        /// <summary>
        /// Specifies the time until expiration.
        /// </summary>
        [Output("expiry")]
        public Output<string?> Expiry { get; private set; } = null!;


        /// <summary>
        /// Create a SecretBackendCrlConfig resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public SecretBackendCrlConfig(string name, SecretBackendCrlConfigArgs args, CustomResourceOptions? options = null)
            : base("vault:pkisecret/secretBackendCrlConfig:SecretBackendCrlConfig", name, args ?? new SecretBackendCrlConfigArgs(), MakeResourceOptions(options, ""))
        {
        }

        private SecretBackendCrlConfig(string name, Input<string> id, SecretBackendCrlConfigState? state = null, CustomResourceOptions? options = null)
            : base("vault:pkisecret/secretBackendCrlConfig:SecretBackendCrlConfig", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing SecretBackendCrlConfig resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static SecretBackendCrlConfig Get(string name, Input<string> id, SecretBackendCrlConfigState? state = null, CustomResourceOptions? options = null)
        {
            return new SecretBackendCrlConfig(name, id, state, options);
        }
    }

    public sealed class SecretBackendCrlConfigArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The path the PKI secret backend is mounted at, with no leading or trailing `/`s.
        /// </summary>
        [Input("backend", required: true)]
        public Input<string> Backend { get; set; } = null!;

        /// <summary>
        /// Disables or enables CRL building.
        /// </summary>
        [Input("disable")]
        public Input<bool>? Disable { get; set; }

        /// <summary>
        /// Specifies the time until expiration.
        /// </summary>
        [Input("expiry")]
        public Input<string>? Expiry { get; set; }

        public SecretBackendCrlConfigArgs()
        {
        }
    }

    public sealed class SecretBackendCrlConfigState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The path the PKI secret backend is mounted at, with no leading or trailing `/`s.
        /// </summary>
        [Input("backend")]
        public Input<string>? Backend { get; set; }

        /// <summary>
        /// Disables or enables CRL building.
        /// </summary>
        [Input("disable")]
        public Input<bool>? Disable { get; set; }

        /// <summary>
        /// Specifies the time until expiration.
        /// </summary>
        [Input("expiry")]
        public Input<string>? Expiry { get; set; }

        public SecretBackendCrlConfigState()
        {
        }
    }
}
