// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vault.Kubernetes
{
    public static class GetAuthBackendConfig
    {
        /// <summary>
        /// Reads the Role of an Kubernetes from a Vault server. See the [Vault
        /// documentation](https://www.vaultproject.io/api-docs/auth/kubernetes#read-config) for more
        /// information.
        /// </summary>
        public static Task<GetAuthBackendConfigResult> InvokeAsync(GetAuthBackendConfigArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetAuthBackendConfigResult>("vault:kubernetes/getAuthBackendConfig:getAuthBackendConfig", args ?? new GetAuthBackendConfigArgs(), options.WithDefaults());

        /// <summary>
        /// Reads the Role of an Kubernetes from a Vault server. See the [Vault
        /// documentation](https://www.vaultproject.io/api-docs/auth/kubernetes#read-config) for more
        /// information.
        /// </summary>
        public static Output<GetAuthBackendConfigResult> Invoke(GetAuthBackendConfigInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetAuthBackendConfigResult>("vault:kubernetes/getAuthBackendConfig:getAuthBackendConfig", args ?? new GetAuthBackendConfigInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetAuthBackendConfigArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The unique name for the Kubernetes backend the config to
        /// retrieve Role attributes for resides in. Defaults to "kubernetes".
        /// </summary>
        [Input("backend")]
        public string? Backend { get; set; }

        /// <summary>
        /// (Optional) Disable JWT issuer validation. Allows to skip ISS validation. Requires Vault `v1.5.4+` or Vault auth kubernetes plugin `v0.7.1+`
        /// </summary>
        [Input("disableIssValidation")]
        public bool? DisableIssValidation { get; set; }

        /// <summary>
        /// (Optional) Disable defaulting to the local CA cert and service account JWT when running in a Kubernetes pod. Requires Vault `v1.5.4+` or Vault auth kubernetes plugin `v0.7.1+`
        /// </summary>
        [Input("disableLocalCaJwt")]
        public bool? DisableLocalCaJwt { get; set; }

        /// <summary>
        /// Optional JWT issuer. If no issuer is specified, `kubernetes.io/serviceaccount` will be used as the default issuer.
        /// </summary>
        [Input("issuer")]
        public string? Issuer { get; set; }

        /// <summary>
        /// PEM encoded CA cert for use by the TLS client used to talk with the Kubernetes API.
        /// </summary>
        [Input("kubernetesCaCert")]
        public string? KubernetesCaCert { get; set; }

        /// <summary>
        /// Host must be a host string, a host:port pair, or a URL to the base of the Kubernetes API server.
        /// </summary>
        [Input("kubernetesHost")]
        public string? KubernetesHost { get; set; }

        /// <summary>
        /// The namespace of the target resource.
        /// The value should not contain leading or trailing forward slashes.
        /// The `namespace` is always relative to the provider's configured namespace.
        /// *Available only for Vault Enterprise*.
        /// </summary>
        [Input("namespace")]
        public string? Namespace { get; set; }

        [Input("pemKeys")]
        private List<string>? _pemKeys;

        /// <summary>
        /// Optional list of PEM-formatted public keys or certificates used to verify the signatures of Kubernetes service account JWTs. If a certificate is given, its public key will be extracted. Not every installation of Kubernetes exposes these keys.
        /// </summary>
        public List<string> PemKeys
        {
            get => _pemKeys ?? (_pemKeys = new List<string>());
            set => _pemKeys = value;
        }

        /// <summary>
        /// (Optional) Use annotations from the client token's associated service account as alias metadata for the Vault entity. Requires Vault `v1.16+` or Vault auth kubernetes plugin `v0.18.0+`
        /// </summary>
        [Input("useAnnotationsAsAliasMetadata")]
        public bool? UseAnnotationsAsAliasMetadata { get; set; }

        public GetAuthBackendConfigArgs()
        {
        }
        public static new GetAuthBackendConfigArgs Empty => new GetAuthBackendConfigArgs();
    }

    public sealed class GetAuthBackendConfigInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The unique name for the Kubernetes backend the config to
        /// retrieve Role attributes for resides in. Defaults to "kubernetes".
        /// </summary>
        [Input("backend")]
        public Input<string>? Backend { get; set; }

        /// <summary>
        /// (Optional) Disable JWT issuer validation. Allows to skip ISS validation. Requires Vault `v1.5.4+` or Vault auth kubernetes plugin `v0.7.1+`
        /// </summary>
        [Input("disableIssValidation")]
        public Input<bool>? DisableIssValidation { get; set; }

        /// <summary>
        /// (Optional) Disable defaulting to the local CA cert and service account JWT when running in a Kubernetes pod. Requires Vault `v1.5.4+` or Vault auth kubernetes plugin `v0.7.1+`
        /// </summary>
        [Input("disableLocalCaJwt")]
        public Input<bool>? DisableLocalCaJwt { get; set; }

        /// <summary>
        /// Optional JWT issuer. If no issuer is specified, `kubernetes.io/serviceaccount` will be used as the default issuer.
        /// </summary>
        [Input("issuer")]
        public Input<string>? Issuer { get; set; }

        /// <summary>
        /// PEM encoded CA cert for use by the TLS client used to talk with the Kubernetes API.
        /// </summary>
        [Input("kubernetesCaCert")]
        public Input<string>? KubernetesCaCert { get; set; }

        /// <summary>
        /// Host must be a host string, a host:port pair, or a URL to the base of the Kubernetes API server.
        /// </summary>
        [Input("kubernetesHost")]
        public Input<string>? KubernetesHost { get; set; }

        /// <summary>
        /// The namespace of the target resource.
        /// The value should not contain leading or trailing forward slashes.
        /// The `namespace` is always relative to the provider's configured namespace.
        /// *Available only for Vault Enterprise*.
        /// </summary>
        [Input("namespace")]
        public Input<string>? Namespace { get; set; }

        [Input("pemKeys")]
        private InputList<string>? _pemKeys;

        /// <summary>
        /// Optional list of PEM-formatted public keys or certificates used to verify the signatures of Kubernetes service account JWTs. If a certificate is given, its public key will be extracted. Not every installation of Kubernetes exposes these keys.
        /// </summary>
        public InputList<string> PemKeys
        {
            get => _pemKeys ?? (_pemKeys = new InputList<string>());
            set => _pemKeys = value;
        }

        /// <summary>
        /// (Optional) Use annotations from the client token's associated service account as alias metadata for the Vault entity. Requires Vault `v1.16+` or Vault auth kubernetes plugin `v0.18.0+`
        /// </summary>
        [Input("useAnnotationsAsAliasMetadata")]
        public Input<bool>? UseAnnotationsAsAliasMetadata { get; set; }

        public GetAuthBackendConfigInvokeArgs()
        {
        }
        public static new GetAuthBackendConfigInvokeArgs Empty => new GetAuthBackendConfigInvokeArgs();
    }


    [OutputType]
    public sealed class GetAuthBackendConfigResult
    {
        public readonly string? Backend;
        /// <summary>
        /// (Optional) Disable JWT issuer validation. Allows to skip ISS validation. Requires Vault `v1.5.4+` or Vault auth kubernetes plugin `v0.7.1+`
        /// </summary>
        public readonly bool DisableIssValidation;
        /// <summary>
        /// (Optional) Disable defaulting to the local CA cert and service account JWT when running in a Kubernetes pod. Requires Vault `v1.5.4+` or Vault auth kubernetes plugin `v0.7.1+`
        /// </summary>
        public readonly bool DisableLocalCaJwt;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Optional JWT issuer. If no issuer is specified, `kubernetes.io/serviceaccount` will be used as the default issuer.
        /// </summary>
        public readonly string Issuer;
        /// <summary>
        /// PEM encoded CA cert for use by the TLS client used to talk with the Kubernetes API.
        /// </summary>
        public readonly string KubernetesCaCert;
        /// <summary>
        /// Host must be a host string, a host:port pair, or a URL to the base of the Kubernetes API server.
        /// </summary>
        public readonly string KubernetesHost;
        public readonly string? Namespace;
        /// <summary>
        /// Optional list of PEM-formatted public keys or certificates used to verify the signatures of Kubernetes service account JWTs. If a certificate is given, its public key will be extracted. Not every installation of Kubernetes exposes these keys.
        /// </summary>
        public readonly ImmutableArray<string> PemKeys;
        /// <summary>
        /// (Optional) Use annotations from the client token's associated service account as alias metadata for the Vault entity. Requires Vault `v1.16+` or Vault auth kubernetes plugin `v0.18.0+`
        /// </summary>
        public readonly bool UseAnnotationsAsAliasMetadata;

        [OutputConstructor]
        private GetAuthBackendConfigResult(
            string? backend,

            bool disableIssValidation,

            bool disableLocalCaJwt,

            string id,

            string issuer,

            string kubernetesCaCert,

            string kubernetesHost,

            string? @namespace,

            ImmutableArray<string> pemKeys,

            bool useAnnotationsAsAliasMetadata)
        {
            Backend = backend;
            DisableIssValidation = disableIssValidation;
            DisableLocalCaJwt = disableLocalCaJwt;
            Id = id;
            Issuer = issuer;
            KubernetesCaCert = kubernetesCaCert;
            KubernetesHost = kubernetesHost;
            Namespace = @namespace;
            PemKeys = pemKeys;
            UseAnnotationsAsAliasMetadata = useAnnotationsAsAliasMetadata;
        }
    }
}
