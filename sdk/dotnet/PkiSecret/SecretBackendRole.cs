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
    /// Creates a role on an PKI Secret Backend for Vault.
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
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
    ///     var role = new Vault.PkiSecret.SecretBackendRole("role", new()
    ///     {
    ///         Backend = pki.Path,
    ///         Ttl = "3600",
    ///         AllowIpSans = true,
    ///         KeyType = "rsa",
    ///         KeyBits = 4096,
    ///         AllowedDomains = new[]
    ///         {
    ///             "example.com",
    ///             "my.domain",
    ///         },
    ///         AllowSubdomains = true,
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// PKI secret backend roles can be imported using the `path`, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import vault:pkiSecret/secretBackendRole:SecretBackendRole role pki/roles/my_role
    /// ```
    /// </summary>
    [VaultResourceType("vault:pkiSecret/secretBackendRole:SecretBackendRole")]
    public partial class SecretBackendRole : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Flag to allow any name
        /// </summary>
        [Output("allowAnyName")]
        public Output<bool?> AllowAnyName { get; private set; } = null!;

        /// <summary>
        /// Flag to allow certificates matching the actual domain
        /// </summary>
        [Output("allowBareDomains")]
        public Output<bool?> AllowBareDomains { get; private set; } = null!;

        /// <summary>
        /// Flag to allow names containing glob patterns.
        /// </summary>
        [Output("allowGlobDomains")]
        public Output<bool?> AllowGlobDomains { get; private set; } = null!;

        /// <summary>
        /// Flag to allow IP SANs
        /// </summary>
        [Output("allowIpSans")]
        public Output<bool?> AllowIpSans { get; private set; } = null!;

        /// <summary>
        /// Flag to allow certificates for localhost
        /// </summary>
        [Output("allowLocalhost")]
        public Output<bool?> AllowLocalhost { get; private set; } = null!;

        /// <summary>
        /// Flag to allow certificates matching subdomains
        /// </summary>
        [Output("allowSubdomains")]
        public Output<bool?> AllowSubdomains { get; private set; } = null!;

        /// <summary>
        /// Flag to allow wildcard certificates.
        /// </summary>
        [Output("allowWildcardCertificates")]
        public Output<bool?> AllowWildcardCertificates { get; private set; } = null!;

        /// <summary>
        /// List of allowed domains for certificates
        /// </summary>
        [Output("allowedDomains")]
        public Output<ImmutableArray<string>> AllowedDomains { get; private set; } = null!;

        /// <summary>
        /// Flag, if set, `allowed_domains` can be specified using identity template expressions such as `{{identity.entity.aliases.&lt;mount accessor&gt;.name}}`.
        /// </summary>
        [Output("allowedDomainsTemplate")]
        public Output<bool?> AllowedDomainsTemplate { get; private set; } = null!;

        /// <summary>
        /// Defines allowed custom SANs
        /// </summary>
        [Output("allowedOtherSans")]
        public Output<ImmutableArray<string>> AllowedOtherSans { get; private set; } = null!;

        /// <summary>
        /// An array of allowed serial numbers to put in Subject
        /// </summary>
        [Output("allowedSerialNumbers")]
        public Output<ImmutableArray<string>> AllowedSerialNumbers { get; private set; } = null!;

        /// <summary>
        /// Defines allowed URI SANs
        /// </summary>
        [Output("allowedUriSans")]
        public Output<ImmutableArray<string>> AllowedUriSans { get; private set; } = null!;

        /// <summary>
        /// Flag, if set, `allowed_uri_sans` can be specified using identity template expressions such as `{{identity.entity.aliases.&lt;mount accessor&gt;.name}}`.
        /// </summary>
        [Output("allowedUriSansTemplate")]
        public Output<bool> AllowedUriSansTemplate { get; private set; } = null!;

        /// <summary>
        /// Defines allowed User IDs
        /// </summary>
        [Output("allowedUserIds")]
        public Output<ImmutableArray<string>> AllowedUserIds { get; private set; } = null!;

        /// <summary>
        /// The path the PKI secret backend is mounted at, with no leading or trailing `/`s.
        /// </summary>
        [Output("backend")]
        public Output<string> Backend { get; private set; } = null!;

        /// <summary>
        /// Flag to mark basic constraints valid when issuing non-CA certificates
        /// </summary>
        [Output("basicConstraintsValidForNonCa")]
        public Output<bool?> BasicConstraintsValidForNonCa { get; private set; } = null!;

        /// <summary>
        /// Flag to specify certificates for client use
        /// </summary>
        [Output("clientFlag")]
        public Output<bool?> ClientFlag { get; private set; } = null!;

        /// <summary>
        /// Flag to specify certificates for code signing use
        /// </summary>
        [Output("codeSigningFlag")]
        public Output<bool?> CodeSigningFlag { get; private set; } = null!;

        /// <summary>
        /// The country of generated certificates
        /// </summary>
        [Output("countries")]
        public Output<ImmutableArray<string>> Countries { get; private set; } = null!;

        /// <summary>
        /// Flag to specify certificates for email protection use
        /// </summary>
        [Output("emailProtectionFlag")]
        public Output<bool?> EmailProtectionFlag { get; private set; } = null!;

        /// <summary>
        /// Flag to allow only valid host names
        /// </summary>
        [Output("enforceHostnames")]
        public Output<bool?> EnforceHostnames { get; private set; } = null!;

        /// <summary>
        /// Specify the allowed extended key usage constraint on issued certificates
        /// </summary>
        [Output("extKeyUsages")]
        public Output<ImmutableArray<string>> ExtKeyUsages { get; private set; } = null!;

        /// <summary>
        /// Flag to generate leases with certificates
        /// </summary>
        [Output("generateLease")]
        public Output<bool?> GenerateLease { get; private set; } = null!;

        /// <summary>
        /// Specifies the default issuer of this request. May
        /// be the value `default`, a name, or an issuer ID. Use ACLs to prevent access to
        /// the `/pki/issuer/:issuer_ref/{issue,sign}/:name` paths to prevent users
        /// overriding the role's `issuer_ref` value.
        /// </summary>
        [Output("issuerRef")]
        public Output<string> IssuerRef { get; private set; } = null!;

        /// <summary>
        /// The number of bits of generated keys
        /// </summary>
        [Output("keyBits")]
        public Output<int?> KeyBits { get; private set; } = null!;

        /// <summary>
        /// The generated key type, choices: `rsa`, `ec`, `ed25519`, `any`
        /// Defaults to `rsa`
        /// </summary>
        [Output("keyType")]
        public Output<string?> KeyType { get; private set; } = null!;

        /// <summary>
        /// Specify the allowed key usage constraint on issued certificates
        /// </summary>
        [Output("keyUsages")]
        public Output<ImmutableArray<string>> KeyUsages { get; private set; } = null!;

        /// <summary>
        /// The locality of generated certificates
        /// </summary>
        [Output("localities")]
        public Output<ImmutableArray<string>> Localities { get; private set; } = null!;

        /// <summary>
        /// The maximum lease TTL, in seconds, for the role.
        /// </summary>
        [Output("maxTtl")]
        public Output<string> MaxTtl { get; private set; } = null!;

        /// <summary>
        /// The name to identify this role within the backend. Must be unique within the backend.
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
        /// Flag to not store certificates in the storage backend
        /// </summary>
        [Output("noStore")]
        public Output<bool?> NoStore { get; private set; } = null!;

        /// <summary>
        /// Specifies the duration by which to backdate the NotBefore property.
        /// </summary>
        [Output("notBeforeDuration")]
        public Output<string> NotBeforeDuration { get; private set; } = null!;

        /// <summary>
        /// The organization unit of generated certificates
        /// </summary>
        [Output("organizationUnit")]
        public Output<ImmutableArray<string>> OrganizationUnit { get; private set; } = null!;

        /// <summary>
        /// The organization of generated certificates
        /// </summary>
        [Output("organizations")]
        public Output<ImmutableArray<string>> Organizations { get; private set; } = null!;

        /// <summary>
        /// Specify the list of allowed policies OIDs. Use with Vault 1.10 or before. For Vault 1.11+, use `policy_identifier` blocks instead
        /// </summary>
        [Output("policyIdentifiers")]
        public Output<ImmutableArray<string>> PolicyIdentifiers { get; private set; } = null!;

        /// <summary>
        /// The postal code of generated certificates
        /// </summary>
        [Output("postalCodes")]
        public Output<ImmutableArray<string>> PostalCodes { get; private set; } = null!;

        /// <summary>
        /// The province of generated certificates
        /// </summary>
        [Output("provinces")]
        public Output<ImmutableArray<string>> Provinces { get; private set; } = null!;

        /// <summary>
        /// Flag to force CN usage
        /// </summary>
        [Output("requireCn")]
        public Output<bool?> RequireCn { get; private set; } = null!;

        /// <summary>
        /// Flag to specify certificates for server use
        /// </summary>
        [Output("serverFlag")]
        public Output<bool?> ServerFlag { get; private set; } = null!;

        /// <summary>
        /// The street address of generated certificates
        /// </summary>
        [Output("streetAddresses")]
        public Output<ImmutableArray<string>> StreetAddresses { get; private set; } = null!;

        /// <summary>
        /// The TTL, in seconds, for any certificate issued against this role.
        /// </summary>
        [Output("ttl")]
        public Output<string> Ttl { get; private set; } = null!;

        /// <summary>
        /// Flag to use the CN in the CSR
        /// </summary>
        [Output("useCsrCommonName")]
        public Output<bool?> UseCsrCommonName { get; private set; } = null!;

        /// <summary>
        /// Flag to use the SANs in the CSR
        /// </summary>
        [Output("useCsrSans")]
        public Output<bool?> UseCsrSans { get; private set; } = null!;


        /// <summary>
        /// Create a SecretBackendRole resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public SecretBackendRole(string name, SecretBackendRoleArgs args, CustomResourceOptions? options = null)
            : base("vault:pkiSecret/secretBackendRole:SecretBackendRole", name, args ?? new SecretBackendRoleArgs(), MakeResourceOptions(options, ""))
        {
        }

        private SecretBackendRole(string name, Input<string> id, SecretBackendRoleState? state = null, CustomResourceOptions? options = null)
            : base("vault:pkiSecret/secretBackendRole:SecretBackendRole", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing SecretBackendRole resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static SecretBackendRole Get(string name, Input<string> id, SecretBackendRoleState? state = null, CustomResourceOptions? options = null)
        {
            return new SecretBackendRole(name, id, state, options);
        }
    }

    public sealed class SecretBackendRoleArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Flag to allow any name
        /// </summary>
        [Input("allowAnyName")]
        public Input<bool>? AllowAnyName { get; set; }

        /// <summary>
        /// Flag to allow certificates matching the actual domain
        /// </summary>
        [Input("allowBareDomains")]
        public Input<bool>? AllowBareDomains { get; set; }

        /// <summary>
        /// Flag to allow names containing glob patterns.
        /// </summary>
        [Input("allowGlobDomains")]
        public Input<bool>? AllowGlobDomains { get; set; }

        /// <summary>
        /// Flag to allow IP SANs
        /// </summary>
        [Input("allowIpSans")]
        public Input<bool>? AllowIpSans { get; set; }

        /// <summary>
        /// Flag to allow certificates for localhost
        /// </summary>
        [Input("allowLocalhost")]
        public Input<bool>? AllowLocalhost { get; set; }

        /// <summary>
        /// Flag to allow certificates matching subdomains
        /// </summary>
        [Input("allowSubdomains")]
        public Input<bool>? AllowSubdomains { get; set; }

        /// <summary>
        /// Flag to allow wildcard certificates.
        /// </summary>
        [Input("allowWildcardCertificates")]
        public Input<bool>? AllowWildcardCertificates { get; set; }

        [Input("allowedDomains")]
        private InputList<string>? _allowedDomains;

        /// <summary>
        /// List of allowed domains for certificates
        /// </summary>
        public InputList<string> AllowedDomains
        {
            get => _allowedDomains ?? (_allowedDomains = new InputList<string>());
            set => _allowedDomains = value;
        }

        /// <summary>
        /// Flag, if set, `allowed_domains` can be specified using identity template expressions such as `{{identity.entity.aliases.&lt;mount accessor&gt;.name}}`.
        /// </summary>
        [Input("allowedDomainsTemplate")]
        public Input<bool>? AllowedDomainsTemplate { get; set; }

        [Input("allowedOtherSans")]
        private InputList<string>? _allowedOtherSans;

        /// <summary>
        /// Defines allowed custom SANs
        /// </summary>
        public InputList<string> AllowedOtherSans
        {
            get => _allowedOtherSans ?? (_allowedOtherSans = new InputList<string>());
            set => _allowedOtherSans = value;
        }

        [Input("allowedSerialNumbers")]
        private InputList<string>? _allowedSerialNumbers;

        /// <summary>
        /// An array of allowed serial numbers to put in Subject
        /// </summary>
        public InputList<string> AllowedSerialNumbers
        {
            get => _allowedSerialNumbers ?? (_allowedSerialNumbers = new InputList<string>());
            set => _allowedSerialNumbers = value;
        }

        [Input("allowedUriSans")]
        private InputList<string>? _allowedUriSans;

        /// <summary>
        /// Defines allowed URI SANs
        /// </summary>
        public InputList<string> AllowedUriSans
        {
            get => _allowedUriSans ?? (_allowedUriSans = new InputList<string>());
            set => _allowedUriSans = value;
        }

        /// <summary>
        /// Flag, if set, `allowed_uri_sans` can be specified using identity template expressions such as `{{identity.entity.aliases.&lt;mount accessor&gt;.name}}`.
        /// </summary>
        [Input("allowedUriSansTemplate")]
        public Input<bool>? AllowedUriSansTemplate { get; set; }

        [Input("allowedUserIds")]
        private InputList<string>? _allowedUserIds;

        /// <summary>
        /// Defines allowed User IDs
        /// </summary>
        public InputList<string> AllowedUserIds
        {
            get => _allowedUserIds ?? (_allowedUserIds = new InputList<string>());
            set => _allowedUserIds = value;
        }

        /// <summary>
        /// The path the PKI secret backend is mounted at, with no leading or trailing `/`s.
        /// </summary>
        [Input("backend", required: true)]
        public Input<string> Backend { get; set; } = null!;

        /// <summary>
        /// Flag to mark basic constraints valid when issuing non-CA certificates
        /// </summary>
        [Input("basicConstraintsValidForNonCa")]
        public Input<bool>? BasicConstraintsValidForNonCa { get; set; }

        /// <summary>
        /// Flag to specify certificates for client use
        /// </summary>
        [Input("clientFlag")]
        public Input<bool>? ClientFlag { get; set; }

        /// <summary>
        /// Flag to specify certificates for code signing use
        /// </summary>
        [Input("codeSigningFlag")]
        public Input<bool>? CodeSigningFlag { get; set; }

        [Input("countries")]
        private InputList<string>? _countries;

        /// <summary>
        /// The country of generated certificates
        /// </summary>
        public InputList<string> Countries
        {
            get => _countries ?? (_countries = new InputList<string>());
            set => _countries = value;
        }

        /// <summary>
        /// Flag to specify certificates for email protection use
        /// </summary>
        [Input("emailProtectionFlag")]
        public Input<bool>? EmailProtectionFlag { get; set; }

        /// <summary>
        /// Flag to allow only valid host names
        /// </summary>
        [Input("enforceHostnames")]
        public Input<bool>? EnforceHostnames { get; set; }

        [Input("extKeyUsages")]
        private InputList<string>? _extKeyUsages;

        /// <summary>
        /// Specify the allowed extended key usage constraint on issued certificates
        /// </summary>
        public InputList<string> ExtKeyUsages
        {
            get => _extKeyUsages ?? (_extKeyUsages = new InputList<string>());
            set => _extKeyUsages = value;
        }

        /// <summary>
        /// Flag to generate leases with certificates
        /// </summary>
        [Input("generateLease")]
        public Input<bool>? GenerateLease { get; set; }

        /// <summary>
        /// Specifies the default issuer of this request. May
        /// be the value `default`, a name, or an issuer ID. Use ACLs to prevent access to
        /// the `/pki/issuer/:issuer_ref/{issue,sign}/:name` paths to prevent users
        /// overriding the role's `issuer_ref` value.
        /// </summary>
        [Input("issuerRef")]
        public Input<string>? IssuerRef { get; set; }

        /// <summary>
        /// The number of bits of generated keys
        /// </summary>
        [Input("keyBits")]
        public Input<int>? KeyBits { get; set; }

        /// <summary>
        /// The generated key type, choices: `rsa`, `ec`, `ed25519`, `any`
        /// Defaults to `rsa`
        /// </summary>
        [Input("keyType")]
        public Input<string>? KeyType { get; set; }

        [Input("keyUsages")]
        private InputList<string>? _keyUsages;

        /// <summary>
        /// Specify the allowed key usage constraint on issued certificates
        /// </summary>
        public InputList<string> KeyUsages
        {
            get => _keyUsages ?? (_keyUsages = new InputList<string>());
            set => _keyUsages = value;
        }

        [Input("localities")]
        private InputList<string>? _localities;

        /// <summary>
        /// The locality of generated certificates
        /// </summary>
        public InputList<string> Localities
        {
            get => _localities ?? (_localities = new InputList<string>());
            set => _localities = value;
        }

        /// <summary>
        /// The maximum lease TTL, in seconds, for the role.
        /// </summary>
        [Input("maxTtl")]
        public Input<string>? MaxTtl { get; set; }

        /// <summary>
        /// The name to identify this role within the backend. Must be unique within the backend.
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

        /// <summary>
        /// Flag to not store certificates in the storage backend
        /// </summary>
        [Input("noStore")]
        public Input<bool>? NoStore { get; set; }

        /// <summary>
        /// Specifies the duration by which to backdate the NotBefore property.
        /// </summary>
        [Input("notBeforeDuration")]
        public Input<string>? NotBeforeDuration { get; set; }

        [Input("organizationUnit")]
        private InputList<string>? _organizationUnit;

        /// <summary>
        /// The organization unit of generated certificates
        /// </summary>
        public InputList<string> OrganizationUnit
        {
            get => _organizationUnit ?? (_organizationUnit = new InputList<string>());
            set => _organizationUnit = value;
        }

        [Input("organizations")]
        private InputList<string>? _organizations;

        /// <summary>
        /// The organization of generated certificates
        /// </summary>
        public InputList<string> Organizations
        {
            get => _organizations ?? (_organizations = new InputList<string>());
            set => _organizations = value;
        }

        [Input("policyIdentifiers")]
        private InputList<string>? _policyIdentifiers;

        /// <summary>
        /// Specify the list of allowed policies OIDs. Use with Vault 1.10 or before. For Vault 1.11+, use `policy_identifier` blocks instead
        /// </summary>
        public InputList<string> PolicyIdentifiers
        {
            get => _policyIdentifiers ?? (_policyIdentifiers = new InputList<string>());
            set => _policyIdentifiers = value;
        }

        [Input("postalCodes")]
        private InputList<string>? _postalCodes;

        /// <summary>
        /// The postal code of generated certificates
        /// </summary>
        public InputList<string> PostalCodes
        {
            get => _postalCodes ?? (_postalCodes = new InputList<string>());
            set => _postalCodes = value;
        }

        [Input("provinces")]
        private InputList<string>? _provinces;

        /// <summary>
        /// The province of generated certificates
        /// </summary>
        public InputList<string> Provinces
        {
            get => _provinces ?? (_provinces = new InputList<string>());
            set => _provinces = value;
        }

        /// <summary>
        /// Flag to force CN usage
        /// </summary>
        [Input("requireCn")]
        public Input<bool>? RequireCn { get; set; }

        /// <summary>
        /// Flag to specify certificates for server use
        /// </summary>
        [Input("serverFlag")]
        public Input<bool>? ServerFlag { get; set; }

        [Input("streetAddresses")]
        private InputList<string>? _streetAddresses;

        /// <summary>
        /// The street address of generated certificates
        /// </summary>
        public InputList<string> StreetAddresses
        {
            get => _streetAddresses ?? (_streetAddresses = new InputList<string>());
            set => _streetAddresses = value;
        }

        /// <summary>
        /// The TTL, in seconds, for any certificate issued against this role.
        /// </summary>
        [Input("ttl")]
        public Input<string>? Ttl { get; set; }

        /// <summary>
        /// Flag to use the CN in the CSR
        /// </summary>
        [Input("useCsrCommonName")]
        public Input<bool>? UseCsrCommonName { get; set; }

        /// <summary>
        /// Flag to use the SANs in the CSR
        /// </summary>
        [Input("useCsrSans")]
        public Input<bool>? UseCsrSans { get; set; }

        public SecretBackendRoleArgs()
        {
        }
        public static new SecretBackendRoleArgs Empty => new SecretBackendRoleArgs();
    }

    public sealed class SecretBackendRoleState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Flag to allow any name
        /// </summary>
        [Input("allowAnyName")]
        public Input<bool>? AllowAnyName { get; set; }

        /// <summary>
        /// Flag to allow certificates matching the actual domain
        /// </summary>
        [Input("allowBareDomains")]
        public Input<bool>? AllowBareDomains { get; set; }

        /// <summary>
        /// Flag to allow names containing glob patterns.
        /// </summary>
        [Input("allowGlobDomains")]
        public Input<bool>? AllowGlobDomains { get; set; }

        /// <summary>
        /// Flag to allow IP SANs
        /// </summary>
        [Input("allowIpSans")]
        public Input<bool>? AllowIpSans { get; set; }

        /// <summary>
        /// Flag to allow certificates for localhost
        /// </summary>
        [Input("allowLocalhost")]
        public Input<bool>? AllowLocalhost { get; set; }

        /// <summary>
        /// Flag to allow certificates matching subdomains
        /// </summary>
        [Input("allowSubdomains")]
        public Input<bool>? AllowSubdomains { get; set; }

        /// <summary>
        /// Flag to allow wildcard certificates.
        /// </summary>
        [Input("allowWildcardCertificates")]
        public Input<bool>? AllowWildcardCertificates { get; set; }

        [Input("allowedDomains")]
        private InputList<string>? _allowedDomains;

        /// <summary>
        /// List of allowed domains for certificates
        /// </summary>
        public InputList<string> AllowedDomains
        {
            get => _allowedDomains ?? (_allowedDomains = new InputList<string>());
            set => _allowedDomains = value;
        }

        /// <summary>
        /// Flag, if set, `allowed_domains` can be specified using identity template expressions such as `{{identity.entity.aliases.&lt;mount accessor&gt;.name}}`.
        /// </summary>
        [Input("allowedDomainsTemplate")]
        public Input<bool>? AllowedDomainsTemplate { get; set; }

        [Input("allowedOtherSans")]
        private InputList<string>? _allowedOtherSans;

        /// <summary>
        /// Defines allowed custom SANs
        /// </summary>
        public InputList<string> AllowedOtherSans
        {
            get => _allowedOtherSans ?? (_allowedOtherSans = new InputList<string>());
            set => _allowedOtherSans = value;
        }

        [Input("allowedSerialNumbers")]
        private InputList<string>? _allowedSerialNumbers;

        /// <summary>
        /// An array of allowed serial numbers to put in Subject
        /// </summary>
        public InputList<string> AllowedSerialNumbers
        {
            get => _allowedSerialNumbers ?? (_allowedSerialNumbers = new InputList<string>());
            set => _allowedSerialNumbers = value;
        }

        [Input("allowedUriSans")]
        private InputList<string>? _allowedUriSans;

        /// <summary>
        /// Defines allowed URI SANs
        /// </summary>
        public InputList<string> AllowedUriSans
        {
            get => _allowedUriSans ?? (_allowedUriSans = new InputList<string>());
            set => _allowedUriSans = value;
        }

        /// <summary>
        /// Flag, if set, `allowed_uri_sans` can be specified using identity template expressions such as `{{identity.entity.aliases.&lt;mount accessor&gt;.name}}`.
        /// </summary>
        [Input("allowedUriSansTemplate")]
        public Input<bool>? AllowedUriSansTemplate { get; set; }

        [Input("allowedUserIds")]
        private InputList<string>? _allowedUserIds;

        /// <summary>
        /// Defines allowed User IDs
        /// </summary>
        public InputList<string> AllowedUserIds
        {
            get => _allowedUserIds ?? (_allowedUserIds = new InputList<string>());
            set => _allowedUserIds = value;
        }

        /// <summary>
        /// The path the PKI secret backend is mounted at, with no leading or trailing `/`s.
        /// </summary>
        [Input("backend")]
        public Input<string>? Backend { get; set; }

        /// <summary>
        /// Flag to mark basic constraints valid when issuing non-CA certificates
        /// </summary>
        [Input("basicConstraintsValidForNonCa")]
        public Input<bool>? BasicConstraintsValidForNonCa { get; set; }

        /// <summary>
        /// Flag to specify certificates for client use
        /// </summary>
        [Input("clientFlag")]
        public Input<bool>? ClientFlag { get; set; }

        /// <summary>
        /// Flag to specify certificates for code signing use
        /// </summary>
        [Input("codeSigningFlag")]
        public Input<bool>? CodeSigningFlag { get; set; }

        [Input("countries")]
        private InputList<string>? _countries;

        /// <summary>
        /// The country of generated certificates
        /// </summary>
        public InputList<string> Countries
        {
            get => _countries ?? (_countries = new InputList<string>());
            set => _countries = value;
        }

        /// <summary>
        /// Flag to specify certificates for email protection use
        /// </summary>
        [Input("emailProtectionFlag")]
        public Input<bool>? EmailProtectionFlag { get; set; }

        /// <summary>
        /// Flag to allow only valid host names
        /// </summary>
        [Input("enforceHostnames")]
        public Input<bool>? EnforceHostnames { get; set; }

        [Input("extKeyUsages")]
        private InputList<string>? _extKeyUsages;

        /// <summary>
        /// Specify the allowed extended key usage constraint on issued certificates
        /// </summary>
        public InputList<string> ExtKeyUsages
        {
            get => _extKeyUsages ?? (_extKeyUsages = new InputList<string>());
            set => _extKeyUsages = value;
        }

        /// <summary>
        /// Flag to generate leases with certificates
        /// </summary>
        [Input("generateLease")]
        public Input<bool>? GenerateLease { get; set; }

        /// <summary>
        /// Specifies the default issuer of this request. May
        /// be the value `default`, a name, or an issuer ID. Use ACLs to prevent access to
        /// the `/pki/issuer/:issuer_ref/{issue,sign}/:name` paths to prevent users
        /// overriding the role's `issuer_ref` value.
        /// </summary>
        [Input("issuerRef")]
        public Input<string>? IssuerRef { get; set; }

        /// <summary>
        /// The number of bits of generated keys
        /// </summary>
        [Input("keyBits")]
        public Input<int>? KeyBits { get; set; }

        /// <summary>
        /// The generated key type, choices: `rsa`, `ec`, `ed25519`, `any`
        /// Defaults to `rsa`
        /// </summary>
        [Input("keyType")]
        public Input<string>? KeyType { get; set; }

        [Input("keyUsages")]
        private InputList<string>? _keyUsages;

        /// <summary>
        /// Specify the allowed key usage constraint on issued certificates
        /// </summary>
        public InputList<string> KeyUsages
        {
            get => _keyUsages ?? (_keyUsages = new InputList<string>());
            set => _keyUsages = value;
        }

        [Input("localities")]
        private InputList<string>? _localities;

        /// <summary>
        /// The locality of generated certificates
        /// </summary>
        public InputList<string> Localities
        {
            get => _localities ?? (_localities = new InputList<string>());
            set => _localities = value;
        }

        /// <summary>
        /// The maximum lease TTL, in seconds, for the role.
        /// </summary>
        [Input("maxTtl")]
        public Input<string>? MaxTtl { get; set; }

        /// <summary>
        /// The name to identify this role within the backend. Must be unique within the backend.
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

        /// <summary>
        /// Flag to not store certificates in the storage backend
        /// </summary>
        [Input("noStore")]
        public Input<bool>? NoStore { get; set; }

        /// <summary>
        /// Specifies the duration by which to backdate the NotBefore property.
        /// </summary>
        [Input("notBeforeDuration")]
        public Input<string>? NotBeforeDuration { get; set; }

        [Input("organizationUnit")]
        private InputList<string>? _organizationUnit;

        /// <summary>
        /// The organization unit of generated certificates
        /// </summary>
        public InputList<string> OrganizationUnit
        {
            get => _organizationUnit ?? (_organizationUnit = new InputList<string>());
            set => _organizationUnit = value;
        }

        [Input("organizations")]
        private InputList<string>? _organizations;

        /// <summary>
        /// The organization of generated certificates
        /// </summary>
        public InputList<string> Organizations
        {
            get => _organizations ?? (_organizations = new InputList<string>());
            set => _organizations = value;
        }

        [Input("policyIdentifiers")]
        private InputList<string>? _policyIdentifiers;

        /// <summary>
        /// Specify the list of allowed policies OIDs. Use with Vault 1.10 or before. For Vault 1.11+, use `policy_identifier` blocks instead
        /// </summary>
        public InputList<string> PolicyIdentifiers
        {
            get => _policyIdentifiers ?? (_policyIdentifiers = new InputList<string>());
            set => _policyIdentifiers = value;
        }

        [Input("postalCodes")]
        private InputList<string>? _postalCodes;

        /// <summary>
        /// The postal code of generated certificates
        /// </summary>
        public InputList<string> PostalCodes
        {
            get => _postalCodes ?? (_postalCodes = new InputList<string>());
            set => _postalCodes = value;
        }

        [Input("provinces")]
        private InputList<string>? _provinces;

        /// <summary>
        /// The province of generated certificates
        /// </summary>
        public InputList<string> Provinces
        {
            get => _provinces ?? (_provinces = new InputList<string>());
            set => _provinces = value;
        }

        /// <summary>
        /// Flag to force CN usage
        /// </summary>
        [Input("requireCn")]
        public Input<bool>? RequireCn { get; set; }

        /// <summary>
        /// Flag to specify certificates for server use
        /// </summary>
        [Input("serverFlag")]
        public Input<bool>? ServerFlag { get; set; }

        [Input("streetAddresses")]
        private InputList<string>? _streetAddresses;

        /// <summary>
        /// The street address of generated certificates
        /// </summary>
        public InputList<string> StreetAddresses
        {
            get => _streetAddresses ?? (_streetAddresses = new InputList<string>());
            set => _streetAddresses = value;
        }

        /// <summary>
        /// The TTL, in seconds, for any certificate issued against this role.
        /// </summary>
        [Input("ttl")]
        public Input<string>? Ttl { get; set; }

        /// <summary>
        /// Flag to use the CN in the CSR
        /// </summary>
        [Input("useCsrCommonName")]
        public Input<bool>? UseCsrCommonName { get; set; }

        /// <summary>
        /// Flag to use the SANs in the CSR
        /// </summary>
        [Input("useCsrSans")]
        public Input<bool>? UseCsrSans { get; set; }

        public SecretBackendRoleState()
        {
        }
        public static new SecretBackendRoleState Empty => new SecretBackendRoleState();
    }
}
