// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vault.Inputs
{

    public sealed class ProviderAuthLoginCertArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Path to a file containing the client certificate.
        /// </summary>
        [Input("certFile", required: true)]
        public Input<string> CertFile { get; set; } = null!;

        /// <summary>
        /// Path to a file containing the private key that the certificate was issued for.
        /// </summary>
        [Input("keyFile", required: true)]
        public Input<string> KeyFile { get; set; } = null!;

        /// <summary>
        /// The path where the authentication engine is mounted.
        /// </summary>
        [Input("mount")]
        public Input<string>? Mount { get; set; }

        /// <summary>
        /// Name of the certificate's role
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The authentication engine's namespace. Conflicts with use_root_namespace
        /// </summary>
        [Input("namespace")]
        public Input<string>? Namespace { get; set; }

        /// <summary>
        /// Authenticate to the root Vault namespace. Conflicts with namespace
        /// </summary>
        [Input("useRootNamespace")]
        public Input<bool>? UseRootNamespace { get; set; }

        public ProviderAuthLoginCertArgs()
        {
        }
        public static new ProviderAuthLoginCertArgs Empty => new ProviderAuthLoginCertArgs();
    }
}
