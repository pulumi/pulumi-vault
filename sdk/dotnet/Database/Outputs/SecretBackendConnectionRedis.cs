// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vault.Database.Outputs
{

    [OutputType]
    public sealed class SecretBackendConnectionRedis
    {
        /// <summary>
        /// The contents of a PEM-encoded CA cert file to use to verify the Redis server's identity.
        /// </summary>
        public readonly string? CaCert;
        /// <summary>
        /// Specifies the host to connect to
        /// </summary>
        public readonly string Host;
        /// <summary>
        /// Specifies whether to skip verification of the server certificate when using TLS.
        /// </summary>
        public readonly bool? InsecureTls;
        /// <summary>
        /// Specifies the password corresponding to the given username.
        /// </summary>
        public readonly string Password;
        /// <summary>
        /// The transport port to use to connect to Redis.
        /// </summary>
        public readonly int? Port;
        /// <summary>
        /// Specifies whether to use TLS when connecting to Redis.
        /// </summary>
        public readonly bool? Tls;
        /// <summary>
        /// Specifies the username for Vault to use.
        /// </summary>
        public readonly string Username;

        [OutputConstructor]
        private SecretBackendConnectionRedis(
            string? caCert,

            string host,

            bool? insecureTls,

            string password,

            int? port,

            bool? tls,

            string username)
        {
            CaCert = caCert;
            Host = host;
            InsecureTls = insecureTls;
            Password = password;
            Port = port;
            Tls = tls;
            Username = username;
        }
    }
}
