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
    public sealed class SecretsMountRedi
    {
        /// <summary>
        /// A list of roles that are allowed to use this
        /// connection.
        /// </summary>
        public readonly ImmutableArray<string> AllowedRoles;
        /// <summary>
        /// The path to a PEM-encoded CA cert file to use to verify the Elasticsearch server's identity.
        /// </summary>
        public readonly string? CaCert;
        /// <summary>
        /// A map of sensitive data to pass to the endpoint. Useful for templated connection strings.
        /// 
        /// Supported list of database secrets engines that can be configured:
        /// </summary>
        public readonly ImmutableDictionary<string, object>? Data;
        /// <summary>
        /// The host to connect to.
        /// </summary>
        public readonly string Host;
        /// <summary>
        /// Whether to skip verification of the server
        /// certificate when using TLS.
        /// </summary>
        public readonly bool? InsecureTls;
        /// <summary>
        /// Name of the database connection.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The root credential password used in the connection URL.
        /// </summary>
        public readonly string Password;
        /// <summary>
        /// Specifies the name of the plugin to use.
        /// </summary>
        public readonly string? PluginName;
        /// <summary>
        /// The default port to connect to if no port is specified as
        /// part of the host.
        /// </summary>
        public readonly int? Port;
        /// <summary>
        /// A list of database statements to be executed to rotate the root user's credentials.
        /// </summary>
        public readonly ImmutableArray<string> RootRotationStatements;
        /// <summary>
        /// Whether to use TLS when connecting to Cassandra.
        /// </summary>
        public readonly bool? Tls;
        /// <summary>
        /// The root credential username used in the connection URL.
        /// </summary>
        public readonly string Username;
        /// <summary>
        /// Whether the connection should be verified on
        /// initial configuration or not.
        /// </summary>
        public readonly bool? VerifyConnection;

        [OutputConstructor]
        private SecretsMountRedi(
            ImmutableArray<string> allowedRoles,

            string? caCert,

            ImmutableDictionary<string, object>? data,

            string host,

            bool? insecureTls,

            string name,

            string password,

            string? pluginName,

            int? port,

            ImmutableArray<string> rootRotationStatements,

            bool? tls,

            string username,

            bool? verifyConnection)
        {
            AllowedRoles = allowedRoles;
            CaCert = caCert;
            Data = data;
            Host = host;
            InsecureTls = insecureTls;
            Name = name;
            Password = password;
            PluginName = pluginName;
            Port = port;
            RootRotationStatements = rootRotationStatements;
            Tls = tls;
            Username = username;
            VerifyConnection = verifyConnection;
        }
    }
}
