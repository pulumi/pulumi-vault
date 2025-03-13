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
    ///     var test = new Vault.PkiSecret.SecretBackendRootCert("test", new()
    ///     {
    ///         Backend = pki.Path,
    ///         Type = "internal",
    ///         CommonName = "Root CA",
    ///         Ttl = "315360000",
    ///         Format = "pem",
    ///         PrivateKeyFormat = "der",
    ///         KeyType = "rsa",
    ///         KeyBits = 4096,
    ///         ExcludeCnFromSans = true,
    ///         Ou = "My OU",
    ///         Organization = "My organization",
    ///     }, new CustomResourceOptions
    ///     {
    ///         DependsOn =
    ///         {
    ///             pki,
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// </summary>
    [VaultResourceType("vault:pkiSecret/secretBackendRootCert:SecretBackendRootCert")]
    public partial class SecretBackendRootCert : global::Pulumi.CustomResource
    {
        /// <summary>
        /// List of alternative names
        /// </summary>
        [Output("altNames")]
        public Output<ImmutableArray<string>> AltNames { get; private set; } = null!;

        /// <summary>
        /// The PKI secret backend the resource belongs to.
        /// </summary>
        [Output("backend")]
        public Output<string> Backend { get; private set; } = null!;

        /// <summary>
        /// The certificate.
        /// </summary>
        [Output("certificate")]
        public Output<string> Certificate { get; private set; } = null!;

        /// <summary>
        /// CN of intermediate to create
        /// </summary>
        [Output("commonName")]
        public Output<string> CommonName { get; private set; } = null!;

        /// <summary>
        /// The country
        /// </summary>
        [Output("country")]
        public Output<string?> Country { get; private set; } = null!;

        /// <summary>
        /// Flag to exclude CN from SANs
        /// </summary>
        [Output("excludeCnFromSans")]
        public Output<bool?> ExcludeCnFromSans { get; private set; } = null!;

        /// <summary>
        /// List of domains for which certificates are not allowed to be issued. Requires Vault version 1.19+.
        /// </summary>
        [Output("excludedDnsDomains")]
        public Output<ImmutableArray<string>> ExcludedDnsDomains { get; private set; } = null!;

        /// <summary>
        /// List of email addresses for which certificates are not allowed to be issued. Requires Vault version 1.19+.
        /// </summary>
        [Output("excludedEmailAddresses")]
        public Output<ImmutableArray<string>> ExcludedEmailAddresses { get; private set; } = null!;

        /// <summary>
        /// List of IP ranges for which certificates are not allowed to be issued. Requires Vault version 1.19+.
        /// </summary>
        [Output("excludedIpRanges")]
        public Output<ImmutableArray<string>> ExcludedIpRanges { get; private set; } = null!;

        /// <summary>
        /// List of URI domains for which certificates are not allowed to be issued. Requires Vault version 1.19+.
        /// </summary>
        [Output("excludedUriDomains")]
        public Output<ImmutableArray<string>> ExcludedUriDomains { get; private set; } = null!;

        /// <summary>
        /// The format of data
        /// </summary>
        [Output("format")]
        public Output<string?> Format { get; private set; } = null!;

        /// <summary>
        /// List of alternative IPs
        /// </summary>
        [Output("ipSans")]
        public Output<ImmutableArray<string>> IpSans { get; private set; } = null!;

        /// <summary>
        /// The ID of the generated issuer.
        /// </summary>
        [Output("issuerId")]
        public Output<string> IssuerId { get; private set; } = null!;

        /// <summary>
        /// Provides a name to the specified issuer. The name must be unique
        /// across all issuers and not be the reserved value `default`
        /// </summary>
        [Output("issuerName")]
        public Output<string> IssuerName { get; private set; } = null!;

        /// <summary>
        /// The issuing CA certificate.
        /// </summary>
        [Output("issuingCa")]
        public Output<string> IssuingCa { get; private set; } = null!;

        /// <summary>
        /// The number of bits to use
        /// </summary>
        [Output("keyBits")]
        public Output<int?> KeyBits { get; private set; } = null!;

        /// <summary>
        /// The ID of the generated key.
        /// </summary>
        [Output("keyId")]
        public Output<string> KeyId { get; private set; } = null!;

        /// <summary>
        /// When a new key is created with this request, optionally specifies
        /// the name for this. The global ref `default` may not be used as a name.
        /// </summary>
        [Output("keyName")]
        public Output<string> KeyName { get; private set; } = null!;

        /// <summary>
        /// Specifies the key (either default, by name, or by identifier) to use
        /// for generating this request. Only suitable for `type=existing` requests.
        /// </summary>
        [Output("keyRef")]
        public Output<string> KeyRef { get; private set; } = null!;

        /// <summary>
        /// The desired key type
        /// </summary>
        [Output("keyType")]
        public Output<string?> KeyType { get; private set; } = null!;

        /// <summary>
        /// The locality
        /// </summary>
        [Output("locality")]
        public Output<string?> Locality { get; private set; } = null!;

        /// <summary>
        /// The ID of the previously configured managed key. This field is
        /// required if `type` is `kms` and it conflicts with `managed_key_name`
        /// </summary>
        [Output("managedKeyId")]
        public Output<string> ManagedKeyId { get; private set; } = null!;

        /// <summary>
        /// The name of the previously configured managed key. This field is
        /// required if `type` is `kms`  and it conflicts with `managed_key_id`
        /// </summary>
        [Output("managedKeyName")]
        public Output<string> ManagedKeyName { get; private set; } = null!;

        /// <summary>
        /// The maximum path length to encode in the generated certificate
        /// </summary>
        [Output("maxPathLength")]
        public Output<int?> MaxPathLength { get; private set; } = null!;

        /// <summary>
        /// The namespace to provision the resource in.
        /// The value should not contain leading or trailing forward slashes.
        /// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
        /// *Available only for Vault Enterprise*.
        /// </summary>
        [Output("namespace")]
        public Output<string?> Namespace { get; private set; } = null!;

        /// <summary>
        /// Set the Not After field of the certificate with specified date value. The value format should be given in UTC format YYYY-MM-ddTHH:MM:SSZ. Supports the Y10K end date for IEEE 802.1AR-2018 standard devices, 9999-12-31T23:59:59Z.
        /// </summary>
        [Output("notAfter")]
        public Output<string?> NotAfter { get; private set; } = null!;

        /// <summary>
        /// The organization
        /// </summary>
        [Output("organization")]
        public Output<string?> Organization { get; private set; } = null!;

        /// <summary>
        /// List of other SANs
        /// </summary>
        [Output("otherSans")]
        public Output<ImmutableArray<string>> OtherSans { get; private set; } = null!;

        /// <summary>
        /// The organization unit
        /// </summary>
        [Output("ou")]
        public Output<string?> Ou { get; private set; } = null!;

        /// <summary>
        /// List of domains for which certificates are allowed to be issued
        /// </summary>
        [Output("permittedDnsDomains")]
        public Output<ImmutableArray<string>> PermittedDnsDomains { get; private set; } = null!;

        /// <summary>
        /// List of email addresses for which certificates are allowed to be issued. Requires Vault version 1.19+.
        /// </summary>
        [Output("permittedEmailAddresses")]
        public Output<ImmutableArray<string>> PermittedEmailAddresses { get; private set; } = null!;

        /// <summary>
        /// List of IP ranges for which certificates are allowed to be issued. Requires Vault version 1.19+.
        /// </summary>
        [Output("permittedIpRanges")]
        public Output<ImmutableArray<string>> PermittedIpRanges { get; private set; } = null!;

        /// <summary>
        /// List of URI domains for which certificates are allowed to be issued. Requires Vault version 1.19+.
        /// </summary>
        [Output("permittedUriDomains")]
        public Output<ImmutableArray<string>> PermittedUriDomains { get; private set; } = null!;

        /// <summary>
        /// The postal code
        /// </summary>
        [Output("postalCode")]
        public Output<string?> PostalCode { get; private set; } = null!;

        /// <summary>
        /// The private key format
        /// </summary>
        [Output("privateKeyFormat")]
        public Output<string?> PrivateKeyFormat { get; private set; } = null!;

        /// <summary>
        /// The province
        /// </summary>
        [Output("province")]
        public Output<string?> Province { get; private set; } = null!;

        /// <summary>
        /// The certificate's serial number, hex formatted.
        /// </summary>
        [Output("serialNumber")]
        public Output<string> SerialNumber { get; private set; } = null!;

        /// <summary>
        /// The number of bits to use in the signature algorithm
        /// </summary>
        [Output("signatureBits")]
        public Output<int> SignatureBits { get; private set; } = null!;

        /// <summary>
        /// The street address
        /// </summary>
        [Output("streetAddress")]
        public Output<string?> StreetAddress { get; private set; } = null!;

        /// <summary>
        /// Time to live
        /// </summary>
        [Output("ttl")]
        public Output<string?> Ttl { get; private set; } = null!;

        /// <summary>
        /// Type of intermediate to create. Must be either \"exported\", \"internal\"
        /// or \"kms\"
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;

        /// <summary>
        /// List of alternative URIs
        /// </summary>
        [Output("uriSans")]
        public Output<ImmutableArray<string>> UriSans { get; private set; } = null!;


        /// <summary>
        /// Create a SecretBackendRootCert resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public SecretBackendRootCert(string name, SecretBackendRootCertArgs args, CustomResourceOptions? options = null)
            : base("vault:pkiSecret/secretBackendRootCert:SecretBackendRootCert", name, args ?? new SecretBackendRootCertArgs(), MakeResourceOptions(options, ""))
        {
        }

        private SecretBackendRootCert(string name, Input<string> id, SecretBackendRootCertState? state = null, CustomResourceOptions? options = null)
            : base("vault:pkiSecret/secretBackendRootCert:SecretBackendRootCert", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing SecretBackendRootCert resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static SecretBackendRootCert Get(string name, Input<string> id, SecretBackendRootCertState? state = null, CustomResourceOptions? options = null)
        {
            return new SecretBackendRootCert(name, id, state, options);
        }
    }

    public sealed class SecretBackendRootCertArgs : global::Pulumi.ResourceArgs
    {
        [Input("altNames")]
        private InputList<string>? _altNames;

        /// <summary>
        /// List of alternative names
        /// </summary>
        public InputList<string> AltNames
        {
            get => _altNames ?? (_altNames = new InputList<string>());
            set => _altNames = value;
        }

        /// <summary>
        /// The PKI secret backend the resource belongs to.
        /// </summary>
        [Input("backend", required: true)]
        public Input<string> Backend { get; set; } = null!;

        /// <summary>
        /// CN of intermediate to create
        /// </summary>
        [Input("commonName", required: true)]
        public Input<string> CommonName { get; set; } = null!;

        /// <summary>
        /// The country
        /// </summary>
        [Input("country")]
        public Input<string>? Country { get; set; }

        /// <summary>
        /// Flag to exclude CN from SANs
        /// </summary>
        [Input("excludeCnFromSans")]
        public Input<bool>? ExcludeCnFromSans { get; set; }

        [Input("excludedDnsDomains")]
        private InputList<string>? _excludedDnsDomains;

        /// <summary>
        /// List of domains for which certificates are not allowed to be issued. Requires Vault version 1.19+.
        /// </summary>
        public InputList<string> ExcludedDnsDomains
        {
            get => _excludedDnsDomains ?? (_excludedDnsDomains = new InputList<string>());
            set => _excludedDnsDomains = value;
        }

        [Input("excludedEmailAddresses")]
        private InputList<string>? _excludedEmailAddresses;

        /// <summary>
        /// List of email addresses for which certificates are not allowed to be issued. Requires Vault version 1.19+.
        /// </summary>
        public InputList<string> ExcludedEmailAddresses
        {
            get => _excludedEmailAddresses ?? (_excludedEmailAddresses = new InputList<string>());
            set => _excludedEmailAddresses = value;
        }

        [Input("excludedIpRanges")]
        private InputList<string>? _excludedIpRanges;

        /// <summary>
        /// List of IP ranges for which certificates are not allowed to be issued. Requires Vault version 1.19+.
        /// </summary>
        public InputList<string> ExcludedIpRanges
        {
            get => _excludedIpRanges ?? (_excludedIpRanges = new InputList<string>());
            set => _excludedIpRanges = value;
        }

        [Input("excludedUriDomains")]
        private InputList<string>? _excludedUriDomains;

        /// <summary>
        /// List of URI domains for which certificates are not allowed to be issued. Requires Vault version 1.19+.
        /// </summary>
        public InputList<string> ExcludedUriDomains
        {
            get => _excludedUriDomains ?? (_excludedUriDomains = new InputList<string>());
            set => _excludedUriDomains = value;
        }

        /// <summary>
        /// The format of data
        /// </summary>
        [Input("format")]
        public Input<string>? Format { get; set; }

        [Input("ipSans")]
        private InputList<string>? _ipSans;

        /// <summary>
        /// List of alternative IPs
        /// </summary>
        public InputList<string> IpSans
        {
            get => _ipSans ?? (_ipSans = new InputList<string>());
            set => _ipSans = value;
        }

        /// <summary>
        /// Provides a name to the specified issuer. The name must be unique
        /// across all issuers and not be the reserved value `default`
        /// </summary>
        [Input("issuerName")]
        public Input<string>? IssuerName { get; set; }

        /// <summary>
        /// The number of bits to use
        /// </summary>
        [Input("keyBits")]
        public Input<int>? KeyBits { get; set; }

        /// <summary>
        /// When a new key is created with this request, optionally specifies
        /// the name for this. The global ref `default` may not be used as a name.
        /// </summary>
        [Input("keyName")]
        public Input<string>? KeyName { get; set; }

        /// <summary>
        /// Specifies the key (either default, by name, or by identifier) to use
        /// for generating this request. Only suitable for `type=existing` requests.
        /// </summary>
        [Input("keyRef")]
        public Input<string>? KeyRef { get; set; }

        /// <summary>
        /// The desired key type
        /// </summary>
        [Input("keyType")]
        public Input<string>? KeyType { get; set; }

        /// <summary>
        /// The locality
        /// </summary>
        [Input("locality")]
        public Input<string>? Locality { get; set; }

        /// <summary>
        /// The ID of the previously configured managed key. This field is
        /// required if `type` is `kms` and it conflicts with `managed_key_name`
        /// </summary>
        [Input("managedKeyId")]
        public Input<string>? ManagedKeyId { get; set; }

        /// <summary>
        /// The name of the previously configured managed key. This field is
        /// required if `type` is `kms`  and it conflicts with `managed_key_id`
        /// </summary>
        [Input("managedKeyName")]
        public Input<string>? ManagedKeyName { get; set; }

        /// <summary>
        /// The maximum path length to encode in the generated certificate
        /// </summary>
        [Input("maxPathLength")]
        public Input<int>? MaxPathLength { get; set; }

        /// <summary>
        /// The namespace to provision the resource in.
        /// The value should not contain leading or trailing forward slashes.
        /// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
        /// *Available only for Vault Enterprise*.
        /// </summary>
        [Input("namespace")]
        public Input<string>? Namespace { get; set; }

        /// <summary>
        /// Set the Not After field of the certificate with specified date value. The value format should be given in UTC format YYYY-MM-ddTHH:MM:SSZ. Supports the Y10K end date for IEEE 802.1AR-2018 standard devices, 9999-12-31T23:59:59Z.
        /// </summary>
        [Input("notAfter")]
        public Input<string>? NotAfter { get; set; }

        /// <summary>
        /// The organization
        /// </summary>
        [Input("organization")]
        public Input<string>? Organization { get; set; }

        [Input("otherSans")]
        private InputList<string>? _otherSans;

        /// <summary>
        /// List of other SANs
        /// </summary>
        public InputList<string> OtherSans
        {
            get => _otherSans ?? (_otherSans = new InputList<string>());
            set => _otherSans = value;
        }

        /// <summary>
        /// The organization unit
        /// </summary>
        [Input("ou")]
        public Input<string>? Ou { get; set; }

        [Input("permittedDnsDomains")]
        private InputList<string>? _permittedDnsDomains;

        /// <summary>
        /// List of domains for which certificates are allowed to be issued
        /// </summary>
        public InputList<string> PermittedDnsDomains
        {
            get => _permittedDnsDomains ?? (_permittedDnsDomains = new InputList<string>());
            set => _permittedDnsDomains = value;
        }

        [Input("permittedEmailAddresses")]
        private InputList<string>? _permittedEmailAddresses;

        /// <summary>
        /// List of email addresses for which certificates are allowed to be issued. Requires Vault version 1.19+.
        /// </summary>
        public InputList<string> PermittedEmailAddresses
        {
            get => _permittedEmailAddresses ?? (_permittedEmailAddresses = new InputList<string>());
            set => _permittedEmailAddresses = value;
        }

        [Input("permittedIpRanges")]
        private InputList<string>? _permittedIpRanges;

        /// <summary>
        /// List of IP ranges for which certificates are allowed to be issued. Requires Vault version 1.19+.
        /// </summary>
        public InputList<string> PermittedIpRanges
        {
            get => _permittedIpRanges ?? (_permittedIpRanges = new InputList<string>());
            set => _permittedIpRanges = value;
        }

        [Input("permittedUriDomains")]
        private InputList<string>? _permittedUriDomains;

        /// <summary>
        /// List of URI domains for which certificates are allowed to be issued. Requires Vault version 1.19+.
        /// </summary>
        public InputList<string> PermittedUriDomains
        {
            get => _permittedUriDomains ?? (_permittedUriDomains = new InputList<string>());
            set => _permittedUriDomains = value;
        }

        /// <summary>
        /// The postal code
        /// </summary>
        [Input("postalCode")]
        public Input<string>? PostalCode { get; set; }

        /// <summary>
        /// The private key format
        /// </summary>
        [Input("privateKeyFormat")]
        public Input<string>? PrivateKeyFormat { get; set; }

        /// <summary>
        /// The province
        /// </summary>
        [Input("province")]
        public Input<string>? Province { get; set; }

        /// <summary>
        /// The number of bits to use in the signature algorithm
        /// </summary>
        [Input("signatureBits")]
        public Input<int>? SignatureBits { get; set; }

        /// <summary>
        /// The street address
        /// </summary>
        [Input("streetAddress")]
        public Input<string>? StreetAddress { get; set; }

        /// <summary>
        /// Time to live
        /// </summary>
        [Input("ttl")]
        public Input<string>? Ttl { get; set; }

        /// <summary>
        /// Type of intermediate to create. Must be either \"exported\", \"internal\"
        /// or \"kms\"
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        [Input("uriSans")]
        private InputList<string>? _uriSans;

        /// <summary>
        /// List of alternative URIs
        /// </summary>
        public InputList<string> UriSans
        {
            get => _uriSans ?? (_uriSans = new InputList<string>());
            set => _uriSans = value;
        }

        public SecretBackendRootCertArgs()
        {
        }
        public static new SecretBackendRootCertArgs Empty => new SecretBackendRootCertArgs();
    }

    public sealed class SecretBackendRootCertState : global::Pulumi.ResourceArgs
    {
        [Input("altNames")]
        private InputList<string>? _altNames;

        /// <summary>
        /// List of alternative names
        /// </summary>
        public InputList<string> AltNames
        {
            get => _altNames ?? (_altNames = new InputList<string>());
            set => _altNames = value;
        }

        /// <summary>
        /// The PKI secret backend the resource belongs to.
        /// </summary>
        [Input("backend")]
        public Input<string>? Backend { get; set; }

        /// <summary>
        /// The certificate.
        /// </summary>
        [Input("certificate")]
        public Input<string>? Certificate { get; set; }

        /// <summary>
        /// CN of intermediate to create
        /// </summary>
        [Input("commonName")]
        public Input<string>? CommonName { get; set; }

        /// <summary>
        /// The country
        /// </summary>
        [Input("country")]
        public Input<string>? Country { get; set; }

        /// <summary>
        /// Flag to exclude CN from SANs
        /// </summary>
        [Input("excludeCnFromSans")]
        public Input<bool>? ExcludeCnFromSans { get; set; }

        [Input("excludedDnsDomains")]
        private InputList<string>? _excludedDnsDomains;

        /// <summary>
        /// List of domains for which certificates are not allowed to be issued. Requires Vault version 1.19+.
        /// </summary>
        public InputList<string> ExcludedDnsDomains
        {
            get => _excludedDnsDomains ?? (_excludedDnsDomains = new InputList<string>());
            set => _excludedDnsDomains = value;
        }

        [Input("excludedEmailAddresses")]
        private InputList<string>? _excludedEmailAddresses;

        /// <summary>
        /// List of email addresses for which certificates are not allowed to be issued. Requires Vault version 1.19+.
        /// </summary>
        public InputList<string> ExcludedEmailAddresses
        {
            get => _excludedEmailAddresses ?? (_excludedEmailAddresses = new InputList<string>());
            set => _excludedEmailAddresses = value;
        }

        [Input("excludedIpRanges")]
        private InputList<string>? _excludedIpRanges;

        /// <summary>
        /// List of IP ranges for which certificates are not allowed to be issued. Requires Vault version 1.19+.
        /// </summary>
        public InputList<string> ExcludedIpRanges
        {
            get => _excludedIpRanges ?? (_excludedIpRanges = new InputList<string>());
            set => _excludedIpRanges = value;
        }

        [Input("excludedUriDomains")]
        private InputList<string>? _excludedUriDomains;

        /// <summary>
        /// List of URI domains for which certificates are not allowed to be issued. Requires Vault version 1.19+.
        /// </summary>
        public InputList<string> ExcludedUriDomains
        {
            get => _excludedUriDomains ?? (_excludedUriDomains = new InputList<string>());
            set => _excludedUriDomains = value;
        }

        /// <summary>
        /// The format of data
        /// </summary>
        [Input("format")]
        public Input<string>? Format { get; set; }

        [Input("ipSans")]
        private InputList<string>? _ipSans;

        /// <summary>
        /// List of alternative IPs
        /// </summary>
        public InputList<string> IpSans
        {
            get => _ipSans ?? (_ipSans = new InputList<string>());
            set => _ipSans = value;
        }

        /// <summary>
        /// The ID of the generated issuer.
        /// </summary>
        [Input("issuerId")]
        public Input<string>? IssuerId { get; set; }

        /// <summary>
        /// Provides a name to the specified issuer. The name must be unique
        /// across all issuers and not be the reserved value `default`
        /// </summary>
        [Input("issuerName")]
        public Input<string>? IssuerName { get; set; }

        /// <summary>
        /// The issuing CA certificate.
        /// </summary>
        [Input("issuingCa")]
        public Input<string>? IssuingCa { get; set; }

        /// <summary>
        /// The number of bits to use
        /// </summary>
        [Input("keyBits")]
        public Input<int>? KeyBits { get; set; }

        /// <summary>
        /// The ID of the generated key.
        /// </summary>
        [Input("keyId")]
        public Input<string>? KeyId { get; set; }

        /// <summary>
        /// When a new key is created with this request, optionally specifies
        /// the name for this. The global ref `default` may not be used as a name.
        /// </summary>
        [Input("keyName")]
        public Input<string>? KeyName { get; set; }

        /// <summary>
        /// Specifies the key (either default, by name, or by identifier) to use
        /// for generating this request. Only suitable for `type=existing` requests.
        /// </summary>
        [Input("keyRef")]
        public Input<string>? KeyRef { get; set; }

        /// <summary>
        /// The desired key type
        /// </summary>
        [Input("keyType")]
        public Input<string>? KeyType { get; set; }

        /// <summary>
        /// The locality
        /// </summary>
        [Input("locality")]
        public Input<string>? Locality { get; set; }

        /// <summary>
        /// The ID of the previously configured managed key. This field is
        /// required if `type` is `kms` and it conflicts with `managed_key_name`
        /// </summary>
        [Input("managedKeyId")]
        public Input<string>? ManagedKeyId { get; set; }

        /// <summary>
        /// The name of the previously configured managed key. This field is
        /// required if `type` is `kms`  and it conflicts with `managed_key_id`
        /// </summary>
        [Input("managedKeyName")]
        public Input<string>? ManagedKeyName { get; set; }

        /// <summary>
        /// The maximum path length to encode in the generated certificate
        /// </summary>
        [Input("maxPathLength")]
        public Input<int>? MaxPathLength { get; set; }

        /// <summary>
        /// The namespace to provision the resource in.
        /// The value should not contain leading or trailing forward slashes.
        /// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
        /// *Available only for Vault Enterprise*.
        /// </summary>
        [Input("namespace")]
        public Input<string>? Namespace { get; set; }

        /// <summary>
        /// Set the Not After field of the certificate with specified date value. The value format should be given in UTC format YYYY-MM-ddTHH:MM:SSZ. Supports the Y10K end date for IEEE 802.1AR-2018 standard devices, 9999-12-31T23:59:59Z.
        /// </summary>
        [Input("notAfter")]
        public Input<string>? NotAfter { get; set; }

        /// <summary>
        /// The organization
        /// </summary>
        [Input("organization")]
        public Input<string>? Organization { get; set; }

        [Input("otherSans")]
        private InputList<string>? _otherSans;

        /// <summary>
        /// List of other SANs
        /// </summary>
        public InputList<string> OtherSans
        {
            get => _otherSans ?? (_otherSans = new InputList<string>());
            set => _otherSans = value;
        }

        /// <summary>
        /// The organization unit
        /// </summary>
        [Input("ou")]
        public Input<string>? Ou { get; set; }

        [Input("permittedDnsDomains")]
        private InputList<string>? _permittedDnsDomains;

        /// <summary>
        /// List of domains for which certificates are allowed to be issued
        /// </summary>
        public InputList<string> PermittedDnsDomains
        {
            get => _permittedDnsDomains ?? (_permittedDnsDomains = new InputList<string>());
            set => _permittedDnsDomains = value;
        }

        [Input("permittedEmailAddresses")]
        private InputList<string>? _permittedEmailAddresses;

        /// <summary>
        /// List of email addresses for which certificates are allowed to be issued. Requires Vault version 1.19+.
        /// </summary>
        public InputList<string> PermittedEmailAddresses
        {
            get => _permittedEmailAddresses ?? (_permittedEmailAddresses = new InputList<string>());
            set => _permittedEmailAddresses = value;
        }

        [Input("permittedIpRanges")]
        private InputList<string>? _permittedIpRanges;

        /// <summary>
        /// List of IP ranges for which certificates are allowed to be issued. Requires Vault version 1.19+.
        /// </summary>
        public InputList<string> PermittedIpRanges
        {
            get => _permittedIpRanges ?? (_permittedIpRanges = new InputList<string>());
            set => _permittedIpRanges = value;
        }

        [Input("permittedUriDomains")]
        private InputList<string>? _permittedUriDomains;

        /// <summary>
        /// List of URI domains for which certificates are allowed to be issued. Requires Vault version 1.19+.
        /// </summary>
        public InputList<string> PermittedUriDomains
        {
            get => _permittedUriDomains ?? (_permittedUriDomains = new InputList<string>());
            set => _permittedUriDomains = value;
        }

        /// <summary>
        /// The postal code
        /// </summary>
        [Input("postalCode")]
        public Input<string>? PostalCode { get; set; }

        /// <summary>
        /// The private key format
        /// </summary>
        [Input("privateKeyFormat")]
        public Input<string>? PrivateKeyFormat { get; set; }

        /// <summary>
        /// The province
        /// </summary>
        [Input("province")]
        public Input<string>? Province { get; set; }

        /// <summary>
        /// The certificate's serial number, hex formatted.
        /// </summary>
        [Input("serialNumber")]
        public Input<string>? SerialNumber { get; set; }

        /// <summary>
        /// The number of bits to use in the signature algorithm
        /// </summary>
        [Input("signatureBits")]
        public Input<int>? SignatureBits { get; set; }

        /// <summary>
        /// The street address
        /// </summary>
        [Input("streetAddress")]
        public Input<string>? StreetAddress { get; set; }

        /// <summary>
        /// Time to live
        /// </summary>
        [Input("ttl")]
        public Input<string>? Ttl { get; set; }

        /// <summary>
        /// Type of intermediate to create. Must be either \"exported\", \"internal\"
        /// or \"kms\"
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        [Input("uriSans")]
        private InputList<string>? _uriSans;

        /// <summary>
        /// List of alternative URIs
        /// </summary>
        public InputList<string> UriSans
        {
            get => _uriSans ?? (_uriSans = new InputList<string>());
            set => _uriSans = value;
        }

        public SecretBackendRootCertState()
        {
        }
        public static new SecretBackendRootCertState Empty => new SecretBackendRootCertState();
    }
}
