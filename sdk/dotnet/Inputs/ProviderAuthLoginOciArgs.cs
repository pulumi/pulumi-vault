// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vault.Inputs
{

    public sealed class ProviderAuthLoginOciArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Authentication type to use when getting OCI credentials.
        /// </summary>
        [Input("authType", required: true)]
        public Input<string> AuthType { get; set; } = null!;

        /// <summary>
        /// The path where the authentication engine is mounted.
        /// </summary>
        [Input("mount")]
        public Input<string>? Mount { get; set; }

        /// <summary>
        /// The authentication engine's namespace. Conflicts with use_root_namespace
        /// </summary>
        [Input("namespace")]
        public Input<string>? Namespace { get; set; }

        /// <summary>
        /// Name of the login role.
        /// </summary>
        [Input("role", required: true)]
        public Input<string> Role { get; set; } = null!;

        /// <summary>
        /// Authenticate to the root Vault namespace. Conflicts with namespace
        /// </summary>
        [Input("useRootNamespace")]
        public Input<bool>? UseRootNamespace { get; set; }

        public ProviderAuthLoginOciArgs()
        {
        }
        public static new ProviderAuthLoginOciArgs Empty => new ProviderAuthLoginOciArgs();
    }
}
