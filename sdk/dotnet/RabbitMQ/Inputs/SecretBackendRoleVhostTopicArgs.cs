// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vault.RabbitMQ.Inputs
{

    public sealed class SecretBackendRoleVhostTopicArgs : Pulumi.ResourceArgs
    {
        [Input("host", required: true)]
        public Input<string> Host { get; set; } = null!;

        [Input("vhosts")]
        private InputList<Inputs.SecretBackendRoleVhostTopicVhostArgs>? _vhosts;

        /// <summary>
        /// Specifies a map of virtual hosts to permissions.
        /// </summary>
        public InputList<Inputs.SecretBackendRoleVhostTopicVhostArgs> Vhosts
        {
            get => _vhosts ?? (_vhosts = new InputList<Inputs.SecretBackendRoleVhostTopicVhostArgs>());
            set => _vhosts = value;
        }

        public SecretBackendRoleVhostTopicArgs()
        {
        }
    }
}