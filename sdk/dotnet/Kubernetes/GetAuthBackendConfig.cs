// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vault.Kubernetes
{
    public static partial class Invokes
    {
        /// <summary>
        /// Reads the Role of an Kubernetes from a Vault server. See the [Vault
        /// documentation](https://www.vaultproject.io/api/auth/kubernetes/index.html#read-config) for more
        /// information.
        /// 
        /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-vault/blob/master/website/docs/d/kubernetes_auth_backend_config.html.markdown.
        /// </summary>
        public static Task<GetAuthBackendConfigResult> GetAuthBackendConfig(GetAuthBackendConfigArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetAuthBackendConfigResult>("vault:kubernetes/getAuthBackendConfig:getAuthBackendConfig", args ?? ResourceArgs.Empty, options.WithVersion());
    }

    public sealed class GetAuthBackendConfigArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The unique name for the Kubernetes backend the config to
        /// retrieve Role attributes for resides in. Defaults to "kubernetes".
        /// </summary>
        [Input("backend")]
        public Input<string>? Backend { get; set; }

        [Input("kubernetesCaCert")]
        public Input<string>? KubernetesCaCert { get; set; }

        [Input("kubernetesHost")]
        public Input<string>? KubernetesHost { get; set; }

        [Input("pemKeys")]
        private InputList<string>? _pemKeys;
        public InputList<string> PemKeys
        {
            get => _pemKeys ?? (_pemKeys = new InputList<string>());
            set => _pemKeys = value;
        }

        public GetAuthBackendConfigArgs()
        {
        }
    }

    [OutputType]
    public sealed class GetAuthBackendConfigResult
    {
        public readonly string? Backend;
        /// <summary>
        /// PEM encoded CA cert for use by the TLS client used to talk with the Kubernetes API.
        /// </summary>
        public readonly string KubernetesCaCert;
        /// <summary>
        /// Host must be a host string, a host:port pair, or a URL to the base of the Kubernetes API server.
        /// </summary>
        public readonly string KubernetesHost;
        /// <summary>
        /// Optional list of PEM-formatted public keys or certificates used to verify the signatures of Kubernetes service account JWTs. If a certificate is given, its public key will be extracted. Not every installation of Kubernetes exposes these keys.
        /// </summary>
        public readonly ImmutableArray<string> PemKeys;
        /// <summary>
        /// id is the provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;

        [OutputConstructor]
        private GetAuthBackendConfigResult(
            string? backend,
            string kubernetesCaCert,
            string kubernetesHost,
            ImmutableArray<string> pemKeys,
            string id)
        {
            Backend = backend;
            KubernetesCaCert = kubernetesCaCert;
            KubernetesHost = kubernetesHost;
            PemKeys = pemKeys;
            Id = id;
        }
    }
}