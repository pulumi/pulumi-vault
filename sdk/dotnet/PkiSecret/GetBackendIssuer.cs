// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vault.PkiSecret
{
    public static class GetBackendIssuer
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
        ///     var pki = new Vault.Mount("pki", new()
        ///     {
        ///         Path = "pki",
        ///         Type = "pki",
        ///         Description = "PKI secret engine mount",
        ///     });
        /// 
        ///     var root = new Vault.PkiSecret.SecretBackendRootCert("root", new()
        ///     {
        ///         Backend = pki.Path,
        ///         Type = "internal",
        ///         CommonName = "example",
        ///         Ttl = "86400",
        ///         IssuerName = "example",
        ///     });
        /// 
        ///     var example = Vault.PkiSecret.GetBackendIssuer.Invoke(new()
        ///     {
        ///         Backend = root.Path,
        ///         IssuerRef = root.IssuerId,
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Task<GetBackendIssuerResult> InvokeAsync(GetBackendIssuerArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetBackendIssuerResult>("vault:pkiSecret/getBackendIssuer:getBackendIssuer", args ?? new GetBackendIssuerArgs(), options.WithDefaults());

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
        ///     var pki = new Vault.Mount("pki", new()
        ///     {
        ///         Path = "pki",
        ///         Type = "pki",
        ///         Description = "PKI secret engine mount",
        ///     });
        /// 
        ///     var root = new Vault.PkiSecret.SecretBackendRootCert("root", new()
        ///     {
        ///         Backend = pki.Path,
        ///         Type = "internal",
        ///         CommonName = "example",
        ///         Ttl = "86400",
        ///         IssuerName = "example",
        ///     });
        /// 
        ///     var example = Vault.PkiSecret.GetBackendIssuer.Invoke(new()
        ///     {
        ///         Backend = root.Path,
        ///         IssuerRef = root.IssuerId,
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetBackendIssuerResult> Invoke(GetBackendIssuerInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetBackendIssuerResult>("vault:pkiSecret/getBackendIssuer:getBackendIssuer", args ?? new GetBackendIssuerInvokeArgs(), options.WithDefaults());

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
        ///     var pki = new Vault.Mount("pki", new()
        ///     {
        ///         Path = "pki",
        ///         Type = "pki",
        ///         Description = "PKI secret engine mount",
        ///     });
        /// 
        ///     var root = new Vault.PkiSecret.SecretBackendRootCert("root", new()
        ///     {
        ///         Backend = pki.Path,
        ///         Type = "internal",
        ///         CommonName = "example",
        ///         Ttl = "86400",
        ///         IssuerName = "example",
        ///     });
        /// 
        ///     var example = Vault.PkiSecret.GetBackendIssuer.Invoke(new()
        ///     {
        ///         Backend = root.Path,
        ///         IssuerRef = root.IssuerId,
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetBackendIssuerResult> Invoke(GetBackendIssuerInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetBackendIssuerResult>("vault:pkiSecret/getBackendIssuer:getBackendIssuer", args ?? new GetBackendIssuerInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetBackendIssuerArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The path to the PKI secret backend to
        /// read the issuer from, with no leading or trailing `/`s.
        /// </summary>
        [Input("backend", required: true)]
        public string Backend { get; set; } = null!;

        /// <summary>
        /// This determines whether this
        /// issuer is able to issue certificates where the chain of trust (including the
        /// issued certificate) contain critical extensions not processed by Vault.
        /// </summary>
        [Input("disableCriticalExtensionChecks")]
        public bool? DisableCriticalExtensionChecks { get; set; }

        /// <summary>
        /// This determines whether this issuer is able
        /// to issue certificates where the chain of trust (including the final issued
        /// certificate) contains a link in which the subject of the issuing certificate
        /// does not match the named issuer of the certificate it signed.
        /// </summary>
        [Input("disableNameChecks")]
        public bool? DisableNameChecks { get; set; }

        /// <summary>
        /// This determines whether this
        /// issuer is able to issue certificates where the chain of trust (including the
        /// final issued certificate) violates the name constraints critical extension of
        /// one of the issuer certificates in the chain.
        /// </summary>
        [Input("disableNameConstraintChecks")]
        public bool? DisableNameConstraintChecks { get; set; }

        /// <summary>
        /// This determines whether this issuer
        /// is able to issue certificates where the chain of trust (including the final
        /// issued certificate) is longer than allowed by a certificate authority in that
        /// chain.
        /// </summary>
        [Input("disablePathLengthChecks")]
        public bool? DisablePathLengthChecks { get; set; }

        /// <summary>
        /// Reference to an existing issuer.
        /// </summary>
        [Input("issuerRef", required: true)]
        public string IssuerRef { get; set; } = null!;

        /// <summary>
        /// The namespace of the target resource.
        /// The value should not contain leading or trailing forward slashes.
        /// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
        /// *Available only for Vault Enterprise*.
        /// </summary>
        [Input("namespace")]
        public string? Namespace { get; set; }

        public GetBackendIssuerArgs()
        {
        }
        public static new GetBackendIssuerArgs Empty => new GetBackendIssuerArgs();
    }

    public sealed class GetBackendIssuerInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The path to the PKI secret backend to
        /// read the issuer from, with no leading or trailing `/`s.
        /// </summary>
        [Input("backend", required: true)]
        public Input<string> Backend { get; set; } = null!;

        /// <summary>
        /// This determines whether this
        /// issuer is able to issue certificates where the chain of trust (including the
        /// issued certificate) contain critical extensions not processed by Vault.
        /// </summary>
        [Input("disableCriticalExtensionChecks")]
        public Input<bool>? DisableCriticalExtensionChecks { get; set; }

        /// <summary>
        /// This determines whether this issuer is able
        /// to issue certificates where the chain of trust (including the final issued
        /// certificate) contains a link in which the subject of the issuing certificate
        /// does not match the named issuer of the certificate it signed.
        /// </summary>
        [Input("disableNameChecks")]
        public Input<bool>? DisableNameChecks { get; set; }

        /// <summary>
        /// This determines whether this
        /// issuer is able to issue certificates where the chain of trust (including the
        /// final issued certificate) violates the name constraints critical extension of
        /// one of the issuer certificates in the chain.
        /// </summary>
        [Input("disableNameConstraintChecks")]
        public Input<bool>? DisableNameConstraintChecks { get; set; }

        /// <summary>
        /// This determines whether this issuer
        /// is able to issue certificates where the chain of trust (including the final
        /// issued certificate) is longer than allowed by a certificate authority in that
        /// chain.
        /// </summary>
        [Input("disablePathLengthChecks")]
        public Input<bool>? DisablePathLengthChecks { get; set; }

        /// <summary>
        /// Reference to an existing issuer.
        /// </summary>
        [Input("issuerRef", required: true)]
        public Input<string> IssuerRef { get; set; } = null!;

        /// <summary>
        /// The namespace of the target resource.
        /// The value should not contain leading or trailing forward slashes.
        /// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
        /// *Available only for Vault Enterprise*.
        /// </summary>
        [Input("namespace")]
        public Input<string>? Namespace { get; set; }

        public GetBackendIssuerInvokeArgs()
        {
        }
        public static new GetBackendIssuerInvokeArgs Empty => new GetBackendIssuerInvokeArgs();
    }


    [OutputType]
    public sealed class GetBackendIssuerResult
    {
        public readonly string Backend;
        /// <summary>
        /// The CA chain as a list of format specific certificates.
        /// </summary>
        public readonly ImmutableArray<string> CaChains;
        /// <summary>
        /// Certificate associated with this issuer.
        /// </summary>
        public readonly string Certificate;
        /// <summary>
        /// This determines whether this
        /// issuer is able to issue certificates where the chain of trust (including the
        /// issued certificate) contain critical extensions not processed by Vault.
        /// </summary>
        public readonly bool? DisableCriticalExtensionChecks;
        /// <summary>
        /// This determines whether this issuer is able
        /// to issue certificates where the chain of trust (including the final issued
        /// certificate) contains a link in which the subject of the issuing certificate
        /// does not match the named issuer of the certificate it signed.
        /// </summary>
        public readonly bool? DisableNameChecks;
        /// <summary>
        /// This determines whether this
        /// issuer is able to issue certificates where the chain of trust (including the
        /// final issued certificate) violates the name constraints critical extension of
        /// one of the issuer certificates in the chain.
        /// </summary>
        public readonly bool? DisableNameConstraintChecks;
        /// <summary>
        /// This determines whether this issuer
        /// is able to issue certificates where the chain of trust (including the final
        /// issued certificate) is longer than allowed by a certificate authority in that
        /// chain.
        /// </summary>
        public readonly bool? DisablePathLengthChecks;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// ID of the issuer.
        /// </summary>
        public readonly string IssuerId;
        /// <summary>
        /// Name of the issuer.
        /// </summary>
        public readonly string IssuerName;
        public readonly string IssuerRef;
        /// <summary>
        /// ID of the key used by the issuer.
        /// </summary>
        public readonly string KeyId;
        /// <summary>
        /// Behavior of a leaf's NotAfter field during issuance.
        /// </summary>
        public readonly string LeafNotAfterBehavior;
        /// <summary>
        /// Chain of issuer references to build this issuer's computed 
        /// CAChain field from, when non-empty.
        /// </summary>
        public readonly ImmutableArray<string> ManualChains;
        public readonly string? Namespace;
        /// <summary>
        /// Allowed usages for this issuer.
        /// </summary>
        public readonly string Usage;

        [OutputConstructor]
        private GetBackendIssuerResult(
            string backend,

            ImmutableArray<string> caChains,

            string certificate,

            bool? disableCriticalExtensionChecks,

            bool? disableNameChecks,

            bool? disableNameConstraintChecks,

            bool? disablePathLengthChecks,

            string id,

            string issuerId,

            string issuerName,

            string issuerRef,

            string keyId,

            string leafNotAfterBehavior,

            ImmutableArray<string> manualChains,

            string? @namespace,

            string usage)
        {
            Backend = backend;
            CaChains = caChains;
            Certificate = certificate;
            DisableCriticalExtensionChecks = disableCriticalExtensionChecks;
            DisableNameChecks = disableNameChecks;
            DisableNameConstraintChecks = disableNameConstraintChecks;
            DisablePathLengthChecks = disablePathLengthChecks;
            Id = id;
            IssuerId = issuerId;
            IssuerName = issuerName;
            IssuerRef = issuerRef;
            KeyId = keyId;
            LeafNotAfterBehavior = leafNotAfterBehavior;
            ManualChains = manualChains;
            Namespace = @namespace;
            Usage = usage;
        }
    }
}
