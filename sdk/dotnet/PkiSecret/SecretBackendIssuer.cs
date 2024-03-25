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
    ///     var pki = new Vault.Mount("pki", new()
    ///     {
    ///         Path = "pki",
    ///         Type = "pki",
    ///         DefaultLeaseTtlSeconds = 3600,
    ///         MaxLeaseTtlSeconds = 86400,
    ///     });
    /// 
    ///     var root = new Vault.PkiSecret.SecretBackendRootCert("root", new()
    ///     {
    ///         Backend = pki.Path,
    ///         Type = "internal",
    ///         CommonName = "test",
    ///         Ttl = "86400",
    ///     });
    /// 
    ///     var example = new Vault.PkiSecret.SecretBackendIssuer("example", new()
    ///     {
    ///         Backend = root.Backend,
    ///         IssuerRef = root.IssuerId,
    ///         IssuerName = "example-issuer",
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ## Import
    /// 
    /// PKI secret backend issuer can be imported using the `id`, e.g.
    /// 
    /// ```sh
    /// $ pulumi import vault:pkiSecret/secretBackendIssuer:SecretBackendIssuer example pki/issuer/bf9b0d48-d0dd-652c-30be-77d04fc7e94d
    /// ```
    /// </summary>
    [VaultResourceType("vault:pkiSecret/secretBackendIssuer:SecretBackendIssuer")]
    public partial class SecretBackendIssuer : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The path the PKI secret backend is mounted at, with no
        /// leading or trailing `/`s.
        /// </summary>
        [Output("backend")]
        public Output<string> Backend { get; private set; } = null!;

        /// <summary>
        /// Specifies the URL values for the CRL
        /// Distribution Points field.
        /// </summary>
        [Output("crlDistributionPoints")]
        public Output<ImmutableArray<string>> CrlDistributionPoints { get; private set; } = null!;

        /// <summary>
        /// Specifies that the AIA URL values should
        /// be templated.
        /// </summary>
        [Output("enableAiaUrlTemplating")]
        public Output<bool?> EnableAiaUrlTemplating { get; private set; } = null!;

        /// <summary>
        /// ID of the issuer.
        /// </summary>
        [Output("issuerId")]
        public Output<string> IssuerId { get; private set; } = null!;

        /// <summary>
        /// Name of the issuer.
        /// </summary>
        [Output("issuerName")]
        public Output<string?> IssuerName { get; private set; } = null!;

        /// <summary>
        /// Reference to an existing issuer.
        /// </summary>
        [Output("issuerRef")]
        public Output<string> IssuerRef { get; private set; } = null!;

        /// <summary>
        /// Specifies the URL values for the Issuing
        /// Certificate field.
        /// </summary>
        [Output("issuingCertificates")]
        public Output<ImmutableArray<string>> IssuingCertificates { get; private set; } = null!;

        /// <summary>
        /// Behavior of a leaf's NotAfter field during
        /// issuance.
        /// </summary>
        [Output("leafNotAfterBehavior")]
        public Output<string> LeafNotAfterBehavior { get; private set; } = null!;

        /// <summary>
        /// Chain of issuer references to build this issuer's
        /// computed CAChain field from, when non-empty.
        /// </summary>
        [Output("manualChains")]
        public Output<ImmutableArray<string>> ManualChains { get; private set; } = null!;

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
        /// Which signature algorithm to use
        /// when building CRLs.
        /// </summary>
        [Output("revocationSignatureAlgorithm")]
        public Output<string> RevocationSignatureAlgorithm { get; private set; } = null!;

        /// <summary>
        /// Allowed usages for this issuer.
        /// </summary>
        [Output("usage")]
        public Output<string> Usage { get; private set; } = null!;


        /// <summary>
        /// Create a SecretBackendIssuer resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public SecretBackendIssuer(string name, SecretBackendIssuerArgs args, CustomResourceOptions? options = null)
            : base("vault:pkiSecret/secretBackendIssuer:SecretBackendIssuer", name, args ?? new SecretBackendIssuerArgs(), MakeResourceOptions(options, ""))
        {
        }

        private SecretBackendIssuer(string name, Input<string> id, SecretBackendIssuerState? state = null, CustomResourceOptions? options = null)
            : base("vault:pkiSecret/secretBackendIssuer:SecretBackendIssuer", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing SecretBackendIssuer resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static SecretBackendIssuer Get(string name, Input<string> id, SecretBackendIssuerState? state = null, CustomResourceOptions? options = null)
        {
            return new SecretBackendIssuer(name, id, state, options);
        }
    }

    public sealed class SecretBackendIssuerArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The path the PKI secret backend is mounted at, with no
        /// leading or trailing `/`s.
        /// </summary>
        [Input("backend", required: true)]
        public Input<string> Backend { get; set; } = null!;

        [Input("crlDistributionPoints")]
        private InputList<string>? _crlDistributionPoints;

        /// <summary>
        /// Specifies the URL values for the CRL
        /// Distribution Points field.
        /// </summary>
        public InputList<string> CrlDistributionPoints
        {
            get => _crlDistributionPoints ?? (_crlDistributionPoints = new InputList<string>());
            set => _crlDistributionPoints = value;
        }

        /// <summary>
        /// Specifies that the AIA URL values should
        /// be templated.
        /// </summary>
        [Input("enableAiaUrlTemplating")]
        public Input<bool>? EnableAiaUrlTemplating { get; set; }

        /// <summary>
        /// Name of the issuer.
        /// </summary>
        [Input("issuerName")]
        public Input<string>? IssuerName { get; set; }

        /// <summary>
        /// Reference to an existing issuer.
        /// </summary>
        [Input("issuerRef", required: true)]
        public Input<string> IssuerRef { get; set; } = null!;

        [Input("issuingCertificates")]
        private InputList<string>? _issuingCertificates;

        /// <summary>
        /// Specifies the URL values for the Issuing
        /// Certificate field.
        /// </summary>
        public InputList<string> IssuingCertificates
        {
            get => _issuingCertificates ?? (_issuingCertificates = new InputList<string>());
            set => _issuingCertificates = value;
        }

        /// <summary>
        /// Behavior of a leaf's NotAfter field during
        /// issuance.
        /// </summary>
        [Input("leafNotAfterBehavior")]
        public Input<string>? LeafNotAfterBehavior { get; set; }

        [Input("manualChains")]
        private InputList<string>? _manualChains;

        /// <summary>
        /// Chain of issuer references to build this issuer's
        /// computed CAChain field from, when non-empty.
        /// </summary>
        public InputList<string> ManualChains
        {
            get => _manualChains ?? (_manualChains = new InputList<string>());
            set => _manualChains = value;
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

        /// <summary>
        /// Which signature algorithm to use
        /// when building CRLs.
        /// </summary>
        [Input("revocationSignatureAlgorithm")]
        public Input<string>? RevocationSignatureAlgorithm { get; set; }

        /// <summary>
        /// Allowed usages for this issuer.
        /// </summary>
        [Input("usage")]
        public Input<string>? Usage { get; set; }

        public SecretBackendIssuerArgs()
        {
        }
        public static new SecretBackendIssuerArgs Empty => new SecretBackendIssuerArgs();
    }

    public sealed class SecretBackendIssuerState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The path the PKI secret backend is mounted at, with no
        /// leading or trailing `/`s.
        /// </summary>
        [Input("backend")]
        public Input<string>? Backend { get; set; }

        [Input("crlDistributionPoints")]
        private InputList<string>? _crlDistributionPoints;

        /// <summary>
        /// Specifies the URL values for the CRL
        /// Distribution Points field.
        /// </summary>
        public InputList<string> CrlDistributionPoints
        {
            get => _crlDistributionPoints ?? (_crlDistributionPoints = new InputList<string>());
            set => _crlDistributionPoints = value;
        }

        /// <summary>
        /// Specifies that the AIA URL values should
        /// be templated.
        /// </summary>
        [Input("enableAiaUrlTemplating")]
        public Input<bool>? EnableAiaUrlTemplating { get; set; }

        /// <summary>
        /// ID of the issuer.
        /// </summary>
        [Input("issuerId")]
        public Input<string>? IssuerId { get; set; }

        /// <summary>
        /// Name of the issuer.
        /// </summary>
        [Input("issuerName")]
        public Input<string>? IssuerName { get; set; }

        /// <summary>
        /// Reference to an existing issuer.
        /// </summary>
        [Input("issuerRef")]
        public Input<string>? IssuerRef { get; set; }

        [Input("issuingCertificates")]
        private InputList<string>? _issuingCertificates;

        /// <summary>
        /// Specifies the URL values for the Issuing
        /// Certificate field.
        /// </summary>
        public InputList<string> IssuingCertificates
        {
            get => _issuingCertificates ?? (_issuingCertificates = new InputList<string>());
            set => _issuingCertificates = value;
        }

        /// <summary>
        /// Behavior of a leaf's NotAfter field during
        /// issuance.
        /// </summary>
        [Input("leafNotAfterBehavior")]
        public Input<string>? LeafNotAfterBehavior { get; set; }

        [Input("manualChains")]
        private InputList<string>? _manualChains;

        /// <summary>
        /// Chain of issuer references to build this issuer's
        /// computed CAChain field from, when non-empty.
        /// </summary>
        public InputList<string> ManualChains
        {
            get => _manualChains ?? (_manualChains = new InputList<string>());
            set => _manualChains = value;
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

        /// <summary>
        /// Which signature algorithm to use
        /// when building CRLs.
        /// </summary>
        [Input("revocationSignatureAlgorithm")]
        public Input<string>? RevocationSignatureAlgorithm { get; set; }

        /// <summary>
        /// Allowed usages for this issuer.
        /// </summary>
        [Input("usage")]
        public Input<string>? Usage { get; set; }

        public SecretBackendIssuerState()
        {
        }
        public static new SecretBackendIssuerState Empty => new SecretBackendIssuerState();
    }
}
