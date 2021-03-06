// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vault.Database.Inputs
{

    public sealed class SecretBackendConnectionMongodbGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// A URL containing connection information. See
        /// the [Vault
        /// docs](https://www.vaultproject.io/api-docs/secret/databases/oracle.html#sample-payload)
        /// for an example.
        /// </summary>
        [Input("connectionUrl")]
        public Input<string>? ConnectionUrl { get; set; }

        /// <summary>
        /// The maximum number of seconds to keep
        /// a connection alive for.
        /// </summary>
        [Input("maxConnectionLifetime")]
        public Input<int>? MaxConnectionLifetime { get; set; }

        /// <summary>
        /// The maximum number of idle connections to
        /// maintain.
        /// </summary>
        [Input("maxIdleConnections")]
        public Input<int>? MaxIdleConnections { get; set; }

        /// <summary>
        /// The maximum number of open connections to
        /// use.
        /// </summary>
        [Input("maxOpenConnections")]
        public Input<int>? MaxOpenConnections { get; set; }

        public SecretBackendConnectionMongodbGetArgs()
        {
        }
    }
}
