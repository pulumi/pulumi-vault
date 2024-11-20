// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vault.Database.Outputs
{

    [OutputType]
    public sealed class SecretBackendConnectionCassandra
    {
        /// <summary>
        /// The number of seconds to use as a connection timeout.
        /// </summary>
        public readonly int? ConnectTimeout;
        /// <summary>
        /// Cassandra hosts to connect to.
        /// </summary>
        public readonly ImmutableArray<string> Hosts;
        /// <summary>
        /// Whether to skip verification of the server certificate when using TLS.
        /// </summary>
        public readonly bool? InsecureTls;
        /// <summary>
        /// The password to use when authenticating with Cassandra.
        /// </summary>
        public readonly string? Password;
        /// <summary>
        /// Concatenated PEM blocks containing a certificate and private key; a certificate, private key, and issuing CA certificate; or just a CA certificate.
        /// </summary>
        public readonly string? PemBundle;
        /// <summary>
        /// Specifies JSON containing a certificate and private key; a certificate, private key, and issuing CA certificate; or just a CA certificate.
        /// </summary>
        public readonly string? PemJson;
        /// <summary>
        /// The transport port to use to connect to Cassandra.
        /// </summary>
        public readonly int? Port;
        /// <summary>
        /// The CQL protocol version to use.
        /// </summary>
        public readonly int? ProtocolVersion;
        /// <summary>
        /// Skip permissions checks when a connection to Cassandra is first created. These checks ensure that Vault is able to create roles, but can be resource intensive in clusters with many roles.
        /// </summary>
        public readonly bool? SkipVerification;
        /// <summary>
        /// Whether to use TLS when connecting to Cassandra.
        /// </summary>
        public readonly bool? Tls;
        /// <summary>
        /// The username to use when authenticating with Cassandra.
        /// </summary>
        public readonly string? Username;

        [OutputConstructor]
        private SecretBackendConnectionCassandra(
            int? connectTimeout,

            ImmutableArray<string> hosts,

            bool? insecureTls,

            string? password,

            string? pemBundle,

            string? pemJson,

            int? port,

            int? protocolVersion,

            bool? skipVerification,

            bool? tls,

            string? username)
        {
            ConnectTimeout = connectTimeout;
            Hosts = hosts;
            InsecureTls = insecureTls;
            Password = password;
            PemBundle = pemBundle;
            PemJson = pemJson;
            Port = port;
            ProtocolVersion = protocolVersion;
            SkipVerification = skipVerification;
            Tls = tls;
            Username = username;
        }
    }
}
