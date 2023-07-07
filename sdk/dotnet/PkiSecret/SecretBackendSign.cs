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
    /// using Pulumi;
    /// using Vault = Pulumi.Vault;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var test = new Vault.PkiSecret.SecretBackendSign("test", new()
    ///     {
    ///         Backend = vault_mount.Pki.Path,
    ///         Csr = @"-----BEGIN CERTIFICATE REQUEST-----
    /// MIIEqDCCApACAQAwYzELMAkGA1UEBhMCQVUxEzARBgNVBAgMClNvbWUtU3RhdGUx
    /// ITAfBgNVBAoMGEludGVybmV0IFdpZGdpdHMgUHR5IEx0ZDEcMBoGA1UEAwwTY2Vy
    /// dC50ZXN0Lm15LmRvbWFpbjCCAiIwDQYJKoZIhvcNAQEBBQADggIPADCCAgoCggIB
    /// AJupYCQ8UVCWII1Zof1c6YcSSaM9hEaDU78cfKP5RoSeH10BvrWRfT+mzCONVpNP
    /// CW9Iabtvk6hm0ot6ilnndEyVJbc0g7hdDLBX5BM25D+DGZGJRKUz1V+uBrWmXtIt
    /// Vonj7JTDTe7ViH0GDsB7CvqXFGXO2a2cDYBchLkL6vQiFPshxvUsLtwxuy/qdYgy
    /// X6ya+AUoZcoQGy1XxNjfH6cPtWSWQGEp1oPR6vL9hU3laTZb3C+VV4jZem+he8/0
    /// V+qV6fLG92WTXm2hmf8nrtUqqJ+C7mW/RJod+TviviBadIX0OHXW7k5HVsZood01
    /// te8vMRUNJNiZfa9EMIK5oncbQn0LcM3Wo9VrjpL7jREb/4HCS2gswYGv7hzk9cCS
    /// kVY4rDucchKbApuI3kfzmO7GFOF5eiSkYZpY/czNn7VVM3WCu6dpOX4+3rhgrZQw
    /// kY14L930DaLVRUgve/zKVP2D2GHdEOs+MbV7s96UgigT9pXly/yHPj+1sSYqmnaD
    /// 5b7jSeJusmzO/nrwXVGLsnezR87VzHl9Ux9g5s6zh+R+PrZuVxYsLvoUpaasH47O
    /// gIcBzSb/6pSGZKAUizmYsHsR1k88dAvsQ+FsUDaNokdi9VndEB4QPmiFmjyLV+0I
    /// 1TFoXop4sW11NPz1YCq+IxnYrEaIN3PyhY0GvBJDFY1/AgMBAAGgADANBgkqhkiG
    /// 9w0BAQsFAAOCAgEActuqnqS8Y9UF7e08w7tR3FPzGecWreuvxILrlFEZJxiLPFqL
    /// It7uJvtypCVQvz6UQzKdBYO7tMpRaWViB8DrWzXNZjLMrg+QHcpveg8C0Ett4scG
    /// fnvLk6fTDFYrnGvwHTqiHos5i0y3bFLyS1BGwSpdLAykGtvC+VM8mRyw/Y7CPcKN
    /// 77kebY/9xduW1g2uxWLr0x90RuQDv9psPojT+59tRLGSp5Kt0IeD3QtnAZEFE4aN
    /// vt+Pd69eg3BgZ8ZeDgoqAw3yppvOkpAFiE5pw2qPZaM4SRphl4d2Lek2zNIMyZqv
    /// do5zh356HOgXtDaSg0POnRGrN/Ua+LMCRTg6GEPUnx9uQb/zt8Zu0hIexDGyykp1
    /// OGqtWlv/Nc8UYuS38v0BeB6bMPeoqQUjkqs8nHlAEFn0KlgYdtDC+7SdQx6wS4te
    /// dBKRNDfC4lS3jYJgs55jHqonZgkpSi3bamlxpfpW0ukGBcmq91wRe4bOw/4uD/vf
    /// UwqMWOdCYcU3mdYNjTWy22ORW3SGFQxMBwpUEURCSoeqWr6aJeQ7KAYkx1PrB5T8
    /// OTEc13lWf+B0PU9UJuGTsmpIuImPDVd0EVDayr3mT5dDbqTVDbe8ppf2IswABmf0
    /// o3DybUeUmknYjl109rdSf+76nuREICHatxXgN3xCMFuBaN4WLO+ksd6Y1Ys=
    /// -----END CERTIFICATE REQUEST-----
    /// ",
    ///         CommonName = "test.my.domain",
    ///     }, new CustomResourceOptions
    ///     {
    ///         DependsOn = new[]
    ///         {
    ///             vault_pki_secret_backend_role.Admin,
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// ## Deprecations
    /// 
    /// * `serial` - Use `serial_number` instead.
    /// </summary>
    [VaultResourceType("vault:pkiSecret/secretBackendSign:SecretBackendSign")]
    public partial class SecretBackendSign : global::Pulumi.CustomResource
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
        [Output("caChains")]
        public Output<ImmutableArray<string>> CaChains { get; private set; } = null!;

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
        /// The CSR
        /// </summary>
        [Output("csr")]
        public Output<string> Csr { get; private set; } = null!;

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
        /// Specifies the default issuer of this request. Can
        /// be the value `default`, a name, or an issuer ID. Use ACLs to prevent access to
        /// the `/pki/issuer/:issuer_ref/{issue,sign}/:name` paths to prevent users
        /// overriding the role's `issuer_ref` value.
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
        /// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
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
        /// `true` if the current time (during refresh) is after the start of the early renewal window declared by `min_seconds_remaining`, and `false` otherwise; if `auto_renew` is set to `true` then the provider will plan to replace the certificate once renewal is pending.
        /// </summary>
        [Output("renewPending")]
        public Output<bool> RenewPending { get; private set; } = null!;

        /// <summary>
        /// The serial number.
        /// </summary>
        [Output("serial")]
        public Output<string> Serial { get; private set; } = null!;

        /// <summary>
        /// The certificate's serial number, hex formatted.
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
        /// Create a SecretBackendSign resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public SecretBackendSign(string name, SecretBackendSignArgs args, CustomResourceOptions? options = null)
            : base("vault:pkiSecret/secretBackendSign:SecretBackendSign", name, args ?? new SecretBackendSignArgs(), MakeResourceOptions(options, ""))
        {
        }

        private SecretBackendSign(string name, Input<string> id, SecretBackendSignState? state = null, CustomResourceOptions? options = null)
            : base("vault:pkiSecret/secretBackendSign:SecretBackendSign", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing SecretBackendSign resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static SecretBackendSign Get(string name, Input<string> id, SecretBackendSignState? state = null, CustomResourceOptions? options = null)
        {
            return new SecretBackendSign(name, id, state, options);
        }
    }

    public sealed class SecretBackendSignArgs : global::Pulumi.ResourceArgs
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
        /// The CSR
        /// </summary>
        [Input("csr", required: true)]
        public Input<string> Csr { get; set; } = null!;

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
        /// Specifies the default issuer of this request. Can
        /// be the value `default`, a name, or an issuer ID. Use ACLs to prevent access to
        /// the `/pki/issuer/:issuer_ref/{issue,sign}/:name` paths to prevent users
        /// overriding the role's `issuer_ref` value.
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
        /// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
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

        public SecretBackendSignArgs()
        {
        }
        public static new SecretBackendSignArgs Empty => new SecretBackendSignArgs();
    }

    public sealed class SecretBackendSignState : global::Pulumi.ResourceArgs
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

        [Input("caChains")]
        private InputList<string>? _caChains;

        /// <summary>
        /// The CA chain
        /// </summary>
        public InputList<string> CaChains
        {
            get => _caChains ?? (_caChains = new InputList<string>());
            set => _caChains = value;
        }

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
        /// The CSR
        /// </summary>
        [Input("csr")]
        public Input<string>? Csr { get; set; }

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
        /// Specifies the default issuer of this request. Can
        /// be the value `default`, a name, or an issuer ID. Use ACLs to prevent access to
        /// the `/pki/issuer/:issuer_ref/{issue,sign}/:name` paths to prevent users
        /// overriding the role's `issuer_ref` value.
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
        /// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
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
        /// `true` if the current time (during refresh) is after the start of the early renewal window declared by `min_seconds_remaining`, and `false` otherwise; if `auto_renew` is set to `true` then the provider will plan to replace the certificate once renewal is pending.
        /// </summary>
        [Input("renewPending")]
        public Input<bool>? RenewPending { get; set; }

        /// <summary>
        /// The serial number.
        /// </summary>
        [Input("serial")]
        public Input<string>? Serial { get; set; }

        /// <summary>
        /// The certificate's serial number, hex formatted.
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

        public SecretBackendSignState()
        {
        }
        public static new SecretBackendSignState Empty => new SecretBackendSignState();
    }
}
