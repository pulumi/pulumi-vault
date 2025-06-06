// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vault.RabbitMQ.Outputs
{

    [OutputType]
    public sealed class SecretBackendRoleVhostTopic
    {
        /// <summary>
        /// The vhost to set permissions for.
        /// </summary>
        public readonly string Host;
        /// <summary>
        /// Specifies a map of virtual hosts to permissions.
        /// </summary>
        public readonly ImmutableArray<Outputs.SecretBackendRoleVhostTopicVhost> Vhosts;

        [OutputConstructor]
        private SecretBackendRoleVhostTopic(
            string host,

            ImmutableArray<Outputs.SecretBackendRoleVhostTopicVhost> vhosts)
        {
            Host = host;
            Vhosts = vhosts;
        }
    }
}
