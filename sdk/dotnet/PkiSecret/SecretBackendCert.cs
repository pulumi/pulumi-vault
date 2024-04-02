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
    ///     var app = new Vault.PkiSecret.SecretBackendCert("app", new()
    ///     {
    ///         Backend = vault_mount.Intermediate.Path,
    ///         CommonName = "app.my.domain",
    ///     }, new CustomResourceOptions
    ///     {
    ///         DependsOn =
    ///         {
    ///             vault_pki_secret_backend_role.Admin, 
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// </summary>
    [VaultResourceType("vault:pkiSecret/secretBackendCert:SecretBackendCert")]
    public partial class SecretBackendCert : global::Pulumi.CustomResource
    {
        /// <summary>
        /// List of alternative names
        /// </summary>
        [Output("altNames")]
        public Output<ImmutableArray<string>> AltNames { get; private set; } = null!;

        /// <summary>
        /// If set to `true`, certs will be renewed if the expiration is within `min_seconds_remaining`. Default `false`
        /// </summary>
        [Output("autoRenew")]
        public Output<bool?> AutoRenew { get; private set; } = null!;

        /// <summary>
        /// The PKI secret backend the resource belongs to.
        /// </summary>
        [Output("backend")]
        public Output<string> Backend { get; private set; } = null!;

        /// <summary>
        /// The CA chain
        /// </summary>
        [Output("caChain")]
        public Output<string> CaChain { get; private set; } = null!;

        /// <summary>
        /// The certificate
        /// </summary>
        [Output("certificate")]
        public Output<string> Certificate { get; private set; } = null!;

        /// <summary>
        /// CN of certificate to create
        /// </summary>
        [Output("commonName")]
        public Output<string> CommonName { get; private set; } = null!;

        /// <summary>
        /// Flag to exclude CN from SANs
        /// </summary>
        [Output("excludeCnFromSans")]
        public Output<bool?> ExcludeCnFromSans { get; private set; } = null!;

        /// <summary>
        /// The expiration date of the certificate in unix epoch format
        /// </summary>
        [Output("expiration")]
        public Output<int> Expiration { get; private set; } = null!;

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
        /// Specifies the default issuer of this request.
        /// </summary>
        [Output("issuerRef")]
        public Output<string?> IssuerRef { get; private set; } = null!;

        /// <summary>
        /// The issuing CA
        /// </summary>
        [Output("issuingCa")]
        public Output<string> IssuingCa { get; private set; } = null!;

        /// <summary>
        /// Generate a new certificate when the expiration is within this number of seconds, default is 604800 (7 days)
        /// </summary>
        [Output("minSecondsRemaining")]
        public Output<int?> MinSecondsRemaining { get; private set; } = null!;

        /// <summary>
        /// Name of the role to create the certificate against
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
        /// List of other SANs
        /// </summary>
        [Output("otherSans")]
        public Output<ImmutableArray<string>> OtherSans { get; private set; } = null!;

        /// <summary>
        /// The private key
        /// </summary>
        [Output("privateKey")]
        public Output<string> PrivateKey { get; private set; } = null!;

        /// <summary>
        /// The private key format
        /// </summary>
        [Output("privateKeyFormat")]
        public Output<string?> PrivateKeyFormat { get; private set; } = null!;

        /// <summary>
        /// The private key type
        /// </summary>
        [Output("privateKeyType")]
        public Output<string> PrivateKeyType { get; private set; } = null!;

        /// <summary>
        /// `true` if the current time (during refresh) is after the start of the early renewal window declared by `min_seconds_remaining`, and `false` otherwise; if `auto_renew` is set to `true` then the provider will plan to replace the certificate once renewal is pending.
        /// </summary>
        [Output("renewPending")]
        public Output<bool> RenewPending { get; private set; } = null!;

        /// <summary>
        /// If set to `true`, the certificate will be revoked on resource destruction.
        /// </summary>
        [Output("revoke")]
        public Output<bool?> Revoke { get; private set; } = null!;

        /// <summary>
        /// The serial number
        /// </summary>
        [Output("serialNumber")]
        public Output<string> SerialNumber { get; private set; } = null!;

        /// <summary>
        /// Time to live
        /// </summary>
        [Output("ttl")]
        public Output<string?> Ttl { get; private set; } = null!;

        /// <summary>
        /// List of alternative URIs
        /// </summary>
        [Output("uriSans")]
        public Output<ImmutableArray<string>> UriSans { get; private set; } = null!;

        /// <summary>
        /// List of Subject User IDs
        /// </summary>
        [Output("userIds")]
        public Output<ImmutableArray<string>> UserIds { get; private set; } = null!;


        /// <summary>
        /// Create a SecretBackendCert resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public SecretBackendCert(string name, SecretBackendCertArgs args, CustomResourceOptions? options = null)
            : base("vault:pkiSecret/secretBackendCert:SecretBackendCert", name, args ?? new SecretBackendCertArgs(), MakeResourceOptions(options, ""))
        {
        }

        private SecretBackendCert(string name, Input<string> id, SecretBackendCertState? state = null, CustomResourceOptions? options = null)
            : base("vault:pkiSecret/secretBackendCert:SecretBackendCert", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                AdditionalSecretOutputs =
                {
                    "privateKey",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing SecretBackendCert resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static SecretBackendCert Get(string name, Input<string> id, SecretBackendCertState? state = null, CustomResourceOptions? options = null)
        {
            return new SecretBackendCert(name, id, state, options);
        }
    }

    public sealed class SecretBackendCertArgs : global::Pulumi.ResourceArgs
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
        /// If set to `true`, certs will be renewed if the expiration is within `min_seconds_remaining`. Default `false`
        /// </summary>
        [Input("autoRenew")]
        public Input<bool>? AutoRenew { get; set; }

        /// <summary>
        /// The PKI secret backend the resource belongs to.
        /// </summary>
        [Input("backend", required: true)]
        public Input<string> Backend { get; set; } = null!;

        /// <summary>
        /// CN of certificate to create
        /// </summary>
        [Input("commonName", required: true)]
        public Input<string> CommonName { get; set; } = null!;

        /// <summary>
        /// Flag to exclude CN from SANs
        /// </summary>
        [Input("excludeCnFromSans")]
        public Input<bool>? ExcludeCnFromSans { get; set; }

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
        /// Specifies the default issuer of this request.
        /// </summary>
        [Input("issuerRef")]
        public Input<string>? IssuerRef { get; set; }

        /// <summary>
        /// Generate a new certificate when the expiration is within this number of seconds, default is 604800 (7 days)
        /// </summary>
        [Input("minSecondsRemaining")]
        public Input<int>? MinSecondsRemaining { get; set; }

        /// <summary>
        /// Name of the role to create the certificate against
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
        /// The private key format
        /// </summary>
        [Input("privateKeyFormat")]
        public Input<string>? PrivateKeyFormat { get; set; }

        /// <summary>
        /// If set to `true`, the certificate will be revoked on resource destruction.
        /// </summary>
        [Input("revoke")]
        public Input<bool>? Revoke { get; set; }

        /// <summary>
        /// Time to live
        /// </summary>
        [Input("ttl")]
        public Input<string>? Ttl { get; set; }

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

        [Input("userIds")]
        private InputList<string>? _userIds;

        /// <summary>
        /// List of Subject User IDs
        /// </summary>
        public InputList<string> UserIds
        {
            get => _userIds ?? (_userIds = new InputList<string>());
            set => _userIds = value;
        }

        public SecretBackendCertArgs()
        {
        }
        public static new SecretBackendCertArgs Empty => new SecretBackendCertArgs();
    }

    public sealed class SecretBackendCertState : global::Pulumi.ResourceArgs
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
        /// If set to `true`, certs will be renewed if the expiration is within `min_seconds_remaining`. Default `false`
        /// </summary>
        [Input("autoRenew")]
        public Input<bool>? AutoRenew { get; set; }

        /// <summary>
        /// The PKI secret backend the resource belongs to.
        /// </summary>
        [Input("backend")]
        public Input<string>? Backend { get; set; }

        /// <summary>
        /// The CA chain
        /// </summary>
        [Input("caChain")]
        public Input<string>? CaChain { get; set; }

        /// <summary>
        /// The certificate
        /// </summary>
        [Input("certificate")]
        public Input<string>? Certificate { get; set; }

        /// <summary>
        /// CN of certificate to create
        /// </summary>
        [Input("commonName")]
        public Input<string>? CommonName { get; set; }

        /// <summary>
        /// Flag to exclude CN from SANs
        /// </summary>
        [Input("excludeCnFromSans")]
        public Input<bool>? ExcludeCnFromSans { get; set; }

        /// <summary>
        /// The expiration date of the certificate in unix epoch format
        /// </summary>
        [Input("expiration")]
        public Input<int>? Expiration { get; set; }

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
        /// Specifies the default issuer of this request.
        /// </summary>
        [Input("issuerRef")]
        public Input<string>? IssuerRef { get; set; }

        /// <summary>
        /// The issuing CA
        /// </summary>
        [Input("issuingCa")]
        public Input<string>? IssuingCa { get; set; }

        /// <summary>
        /// Generate a new certificate when the expiration is within this number of seconds, default is 604800 (7 days)
        /// </summary>
        [Input("minSecondsRemaining")]
        public Input<int>? MinSecondsRemaining { get; set; }

        /// <summary>
        /// Name of the role to create the certificate against
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

        [Input("privateKey")]
        private Input<string>? _privateKey;

        /// <summary>
        /// The private key
        /// </summary>
        public Input<string>? PrivateKey
        {
            get => _privateKey;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _privateKey = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// The private key format
        /// </summary>
        [Input("privateKeyFormat")]
        public Input<string>? PrivateKeyFormat { get; set; }

        /// <summary>
        /// The private key type
        /// </summary>
        [Input("privateKeyType")]
        public Input<string>? PrivateKeyType { get; set; }

        /// <summary>
        /// `true` if the current time (during refresh) is after the start of the early renewal window declared by `min_seconds_remaining`, and `false` otherwise; if `auto_renew` is set to `true` then the provider will plan to replace the certificate once renewal is pending.
        /// </summary>
        [Input("renewPending")]
        public Input<bool>? RenewPending { get; set; }

        /// <summary>
        /// If set to `true`, the certificate will be revoked on resource destruction.
        /// </summary>
        [Input("revoke")]
        public Input<bool>? Revoke { get; set; }

        /// <summary>
        /// The serial number
        /// </summary>
        [Input("serialNumber")]
        public Input<string>? SerialNumber { get; set; }

        /// <summary>
        /// Time to live
        /// </summary>
        [Input("ttl")]
        public Input<string>? Ttl { get; set; }

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

        [Input("userIds")]
        private InputList<string>? _userIds;

        /// <summary>
        /// List of Subject User IDs
        /// </summary>
        public InputList<string> UserIds
        {
            get => _userIds ?? (_userIds = new InputList<string>());
            set => _userIds = value;
        }

        public SecretBackendCertState()
        {
        }
        public static new SecretBackendCertState Empty => new SecretBackendCertState();
    }
}
