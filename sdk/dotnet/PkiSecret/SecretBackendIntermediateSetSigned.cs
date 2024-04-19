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
    ///         Description = "root",
    ///         DefaultLeaseTtlSeconds = 8640000,
    ///         MaxLeaseTtlSeconds = 8640000,
    ///     });
    /// 
    ///     var intermediate = new Vault.Mount("intermediate", new()
    ///     {
    ///         Path = "pki-int",
    ///         Type = root.Type,
    ///         Description = "intermediate",
    ///         DefaultLeaseTtlSeconds = 86400,
    ///         MaxLeaseTtlSeconds = 86400,
    ///     });
    /// 
    ///     var exampleSecretBackendRootCert = new Vault.PkiSecret.SecretBackendRootCert("exampleSecretBackendRootCert", new()
    ///     {
    ///         Backend = root.Path,
    ///         Type = "internal",
    ///         CommonName = "RootOrg Root CA",
    ///         Ttl = "86400",
    ///         Format = "pem",
    ///         PrivateKeyFormat = "der",
    ///         KeyType = "rsa",
    ///         KeyBits = 4096,
    ///         ExcludeCnFromSans = true,
    ///         Ou = "Organizational Unit",
    ///         Organization = "RootOrg",
    ///         Country = "US",
    ///         Locality = "San Francisco",
    ///         Province = "CA",
    ///     });
    /// 
    ///     var exampleSecretBackendIntermediateCertRequest = new Vault.PkiSecret.SecretBackendIntermediateCertRequest("exampleSecretBackendIntermediateCertRequest", new()
    ///     {
    ///         Backend = intermediate.Path,
    ///         Type = exampleSecretBackendRootCert.Type,
    ///         CommonName = "SubOrg Intermediate CA",
    ///     });
    /// 
    ///     var exampleSecretBackendRootSignIntermediate = new Vault.PkiSecret.SecretBackendRootSignIntermediate("exampleSecretBackendRootSignIntermediate", new()
    ///     {
    ///         Backend = root.Path,
    ///         Csr = exampleSecretBackendIntermediateCertRequest.Csr,
    ///         CommonName = "SubOrg Intermediate CA",
    ///         ExcludeCnFromSans = true,
    ///         Ou = "SubUnit",
    ///         Organization = "SubOrg",
    ///         Country = "US",
    ///         Locality = "San Francisco",
    ///         Province = "CA",
    ///         Revoke = true,
    ///     });
    /// 
    ///     var exampleSecretBackendIntermediateSetSigned = new Vault.PkiSecret.SecretBackendIntermediateSetSigned("exampleSecretBackendIntermediateSetSigned", new()
    ///     {
    ///         Backend = intermediate.Path,
    ///         Certificate = exampleSecretBackendRootSignIntermediate.Certificate,
    ///     });
    /// 
    /// });
    /// ```
    /// </summary>
    [VaultResourceType("vault:pkiSecret/secretBackendIntermediateSetSigned:SecretBackendIntermediateSetSigned")]
    public partial class SecretBackendIntermediateSetSigned : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The PKI secret backend the resource belongs to.
        /// </summary>
        [Output("backend")]
        public Output<string> Backend { get; private set; } = null!;

        /// <summary>
        /// Specifies the PEM encoded certificate. May optionally append additional
        /// CA certificates to populate the whole chain, which will then enable returning the full chain from
        /// issue and sign operations.
        /// </summary>
        [Output("certificate")]
        public Output<string> Certificate { get; private set; } = null!;

        /// <summary>
        /// The imported issuers indicating which issuers were created as part of
        /// this request.
        /// </summary>
        [Output("importedIssuers")]
        public Output<ImmutableArray<string>> ImportedIssuers { get; private set; } = null!;

        /// <summary>
        /// The imported keys indicating which keys were created as part of this request.
        /// </summary>
        [Output("importedKeys")]
        public Output<ImmutableArray<string>> ImportedKeys { get; private set; } = null!;

        /// <summary>
        /// The namespace to provision the resource in.
        /// The value should not contain leading or trailing forward slashes.
        /// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
        /// *Available only for Vault Enterprise*.
        /// </summary>
        [Output("namespace")]
        public Output<string?> Namespace { get; private set; } = null!;


        /// <summary>
        /// Create a SecretBackendIntermediateSetSigned resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public SecretBackendIntermediateSetSigned(string name, SecretBackendIntermediateSetSignedArgs args, CustomResourceOptions? options = null)
            : base("vault:pkiSecret/secretBackendIntermediateSetSigned:SecretBackendIntermediateSetSigned", name, args ?? new SecretBackendIntermediateSetSignedArgs(), MakeResourceOptions(options, ""))
        {
        }

        private SecretBackendIntermediateSetSigned(string name, Input<string> id, SecretBackendIntermediateSetSignedState? state = null, CustomResourceOptions? options = null)
            : base("vault:pkiSecret/secretBackendIntermediateSetSigned:SecretBackendIntermediateSetSigned", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing SecretBackendIntermediateSetSigned resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static SecretBackendIntermediateSetSigned Get(string name, Input<string> id, SecretBackendIntermediateSetSignedState? state = null, CustomResourceOptions? options = null)
        {
            return new SecretBackendIntermediateSetSigned(name, id, state, options);
        }
    }

    public sealed class SecretBackendIntermediateSetSignedArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The PKI secret backend the resource belongs to.
        /// </summary>
        [Input("backend", required: true)]
        public Input<string> Backend { get; set; } = null!;

        /// <summary>
        /// Specifies the PEM encoded certificate. May optionally append additional
        /// CA certificates to populate the whole chain, which will then enable returning the full chain from
        /// issue and sign operations.
        /// </summary>
        [Input("certificate", required: true)]
        public Input<string> Certificate { get; set; } = null!;

        /// <summary>
        /// The namespace to provision the resource in.
        /// The value should not contain leading or trailing forward slashes.
        /// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
        /// *Available only for Vault Enterprise*.
        /// </summary>
        [Input("namespace")]
        public Input<string>? Namespace { get; set; }

        public SecretBackendIntermediateSetSignedArgs()
        {
        }
        public static new SecretBackendIntermediateSetSignedArgs Empty => new SecretBackendIntermediateSetSignedArgs();
    }

    public sealed class SecretBackendIntermediateSetSignedState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The PKI secret backend the resource belongs to.
        /// </summary>
        [Input("backend")]
        public Input<string>? Backend { get; set; }

        /// <summary>
        /// Specifies the PEM encoded certificate. May optionally append additional
        /// CA certificates to populate the whole chain, which will then enable returning the full chain from
        /// issue and sign operations.
        /// </summary>
        [Input("certificate")]
        public Input<string>? Certificate { get; set; }

        [Input("importedIssuers")]
        private InputList<string>? _importedIssuers;

        /// <summary>
        /// The imported issuers indicating which issuers were created as part of
        /// this request.
        /// </summary>
        public InputList<string> ImportedIssuers
        {
            get => _importedIssuers ?? (_importedIssuers = new InputList<string>());
            set => _importedIssuers = value;
        }

        [Input("importedKeys")]
        private InputList<string>? _importedKeys;

        /// <summary>
        /// The imported keys indicating which keys were created as part of this request.
        /// </summary>
        public InputList<string> ImportedKeys
        {
            get => _importedKeys ?? (_importedKeys = new InputList<string>());
            set => _importedKeys = value;
        }

        /// <summary>
        /// The namespace to provision the resource in.
        /// The value should not contain leading or trailing forward slashes.
        /// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
        /// *Available only for Vault Enterprise*.
        /// </summary>
        [Input("namespace")]
        public Input<string>? Namespace { get; set; }

        public SecretBackendIntermediateSetSignedState()
        {
        }
        public static new SecretBackendIntermediateSetSignedState Empty => new SecretBackendIntermediateSetSignedState();
    }
}
