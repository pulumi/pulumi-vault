// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vault.Inputs
{

    public sealed class ProviderAuthLoginGcpArgs : global::Pulumi.ResourceArgs
    {
        [Input("credentials")]
        public Input<string>? Credentials { get; set; }

        [Input("jwt")]
        public Input<string>? Jwt { get; set; }

        [Input("mount")]
        public Input<string>? Mount { get; set; }

        [Input("namespace")]
        public Input<string>? Namespace { get; set; }

        [Input("role", required: true)]
        public Input<string> Role { get; set; } = null!;

        [Input("serviceAccount")]
        public Input<string>? ServiceAccount { get; set; }

        public ProviderAuthLoginGcpArgs()
        {
        }
        public static new ProviderAuthLoginGcpArgs Empty => new ProviderAuthLoginGcpArgs();
    }
}