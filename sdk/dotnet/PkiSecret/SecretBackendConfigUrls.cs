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
    /// Allows setting the issuing certificate endpoints, CRL distribution points, and OCSP server endpoints that will be encoded into issued certificates.
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
    ///     var root = new Vault.Mount("root", new()
    ///     {
    ///         Path = "pki-root",
    ///         Type = "pki",
    ///         Description = "root PKI",
    ///         DefaultLeaseTtlSeconds = 8640000,
    ///         MaxLeaseTtlSeconds = 8640000,
    ///     });
    /// 
    ///     var example = new Vault.PkiSecret.SecretBackendConfigUrls("example", new()
    ///     {
    ///         Backend = root.Path,
    ///         IssuingCertificates = new[]
    ///         {
    ///             "http://127.0.0.1:8200/v1/pki/ca",
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ## Import
    /// 
    /// The PKI config URLs can be imported using the resource's `id`.
    /// In the case of the example above the `id` would be `pki-root/config/urls`,
    /// where the `pki-root` component is the resource's `backend`, e.g.
    /// 
    /// ```sh
    /// $ pulumi import vault:pkiSecret/secretBackendConfigUrls:SecretBackendConfigUrls example pki-root/config/urls
    /// ```
    /// </summary>
    [VaultResourceType("vault:pkiSecret/secretBackendConfigUrls:SecretBackendConfigUrls")]
    public partial class SecretBackendConfigUrls : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The path the PKI secret backend is mounted at, with no leading or trailing `/`s.
        /// </summary>
        [Output("backend")]
        public Output<string> Backend { get; private set; } = null!;

        /// <summary>
        /// Specifies the URL values for the CRL Distribution Points field.
        /// </summary>
        [Output("crlDistributionPoints")]
        public Output<ImmutableArray<string>> CrlDistributionPoints { get; private set; } = null!;

        /// <summary>
        /// Specifies that templating of AIA fields is allowed.
        /// </summary>
        [Output("enableTemplating")]
        public Output<bool?> EnableTemplating { get; private set; } = null!;

        /// <summary>
        /// Specifies the URL values for the Issuing Certificate field.
        /// </summary>
        [Output("issuingCertificates")]
        public Output<ImmutableArray<string>> IssuingCertificates { get; private set; } = null!;

        /// <summary>
        /// The namespace to provision the resource in.
        /// The value should not contain leading or trailing forward slashes.
        /// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
        /// *Available only for Vault Enterprise*.
        /// </summary>
        [Output("namespace")]
        public Output<string?> Namespace { get; private set; } = null!;

        /// <summary>
        /// Specifies the URL values for the OCSP Servers field.
        /// </summary>
        [Output("ocspServers")]
        public Output<ImmutableArray<string>> OcspServers { get; private set; } = null!;


        /// <summary>
        /// Create a SecretBackendConfigUrls resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public SecretBackendConfigUrls(string name, SecretBackendConfigUrlsArgs args, CustomResourceOptions? options = null)
            : base("vault:pkiSecret/secretBackendConfigUrls:SecretBackendConfigUrls", name, args ?? new SecretBackendConfigUrlsArgs(), MakeResourceOptions(options, ""))
        {
        }

        private SecretBackendConfigUrls(string name, Input<string> id, SecretBackendConfigUrlsState? state = null, CustomResourceOptions? options = null)
            : base("vault:pkiSecret/secretBackendConfigUrls:SecretBackendConfigUrls", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing SecretBackendConfigUrls resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static SecretBackendConfigUrls Get(string name, Input<string> id, SecretBackendConfigUrlsState? state = null, CustomResourceOptions? options = null)
        {
            return new SecretBackendConfigUrls(name, id, state, options);
        }
    }

    public sealed class SecretBackendConfigUrlsArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The path the PKI secret backend is mounted at, with no leading or trailing `/`s.
        /// </summary>
        [Input("backend", required: true)]
        public Input<string> Backend { get; set; } = null!;

        [Input("crlDistributionPoints")]
        private InputList<string>? _crlDistributionPoints;

        /// <summary>
        /// Specifies the URL values for the CRL Distribution Points field.
        /// </summary>
        public InputList<string> CrlDistributionPoints
        {
            get => _crlDistributionPoints ?? (_crlDistributionPoints = new InputList<string>());
            set => _crlDistributionPoints = value;
        }

        /// <summary>
        /// Specifies that templating of AIA fields is allowed.
        /// </summary>
        [Input("enableTemplating")]
        public Input<bool>? EnableTemplating { get; set; }

        [Input("issuingCertificates")]
        private InputList<string>? _issuingCertificates;

        /// <summary>
        /// Specifies the URL values for the Issuing Certificate field.
        /// </summary>
        public InputList<string> IssuingCertificates
        {
            get => _issuingCertificates ?? (_issuingCertificates = new InputList<string>());
            set => _issuingCertificates = value;
        }

        /// <summary>
        /// The namespace to provision the resource in.
        /// The value should not contain leading or trailing forward slashes.
        /// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
        /// *Available only for Vault Enterprise*.
        /// </summary>
        [Input("namespace")]
        public Input<string>? Namespace { get; set; }

        [Input("ocspServers")]
        private InputList<string>? _ocspServers;

        /// <summary>
        /// Specifies the URL values for the OCSP Servers field.
        /// </summary>
        public InputList<string> OcspServers
        {
            get => _ocspServers ?? (_ocspServers = new InputList<string>());
            set => _ocspServers = value;
        }

        public SecretBackendConfigUrlsArgs()
        {
        }
        public static new SecretBackendConfigUrlsArgs Empty => new SecretBackendConfigUrlsArgs();
    }

    public sealed class SecretBackendConfigUrlsState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The path the PKI secret backend is mounted at, with no leading or trailing `/`s.
        /// </summary>
        [Input("backend")]
        public Input<string>? Backend { get; set; }

        [Input("crlDistributionPoints")]
        private InputList<string>? _crlDistributionPoints;

        /// <summary>
        /// Specifies the URL values for the CRL Distribution Points field.
        /// </summary>
        public InputList<string> CrlDistributionPoints
        {
            get => _crlDistributionPoints ?? (_crlDistributionPoints = new InputList<string>());
            set => _crlDistributionPoints = value;
        }

        /// <summary>
        /// Specifies that templating of AIA fields is allowed.
        /// </summary>
        [Input("enableTemplating")]
        public Input<bool>? EnableTemplating { get; set; }

        [Input("issuingCertificates")]
        private InputList<string>? _issuingCertificates;

        /// <summary>
        /// Specifies the URL values for the Issuing Certificate field.
        /// </summary>
        public InputList<string> IssuingCertificates
        {
            get => _issuingCertificates ?? (_issuingCertificates = new InputList<string>());
            set => _issuingCertificates = value;
        }

        /// <summary>
        /// The namespace to provision the resource in.
        /// The value should not contain leading or trailing forward slashes.
        /// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
        /// *Available only for Vault Enterprise*.
        /// </summary>
        [Input("namespace")]
        public Input<string>? Namespace { get; set; }

        [Input("ocspServers")]
        private InputList<string>? _ocspServers;

        /// <summary>
        /// Specifies the URL values for the OCSP Servers field.
        /// </summary>
        public InputList<string> OcspServers
        {
            get => _ocspServers ?? (_ocspServers = new InputList<string>());
            set => _ocspServers = value;
        }

        public SecretBackendConfigUrlsState()
        {
        }
        public static new SecretBackendConfigUrlsState Empty => new SecretBackendConfigUrlsState();
    }
}
