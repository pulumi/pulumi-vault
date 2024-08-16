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
    public sealed class SecretsMountOracle
    {
        /// <summary>
        /// A list of roles that are allowed to use this
        /// connection.
        /// </summary>
        public readonly ImmutableArray<string> AllowedRoles;
        /// <summary>
        /// Connection string to use to connect to the database.
        /// </summary>
        public readonly string? ConnectionUrl;
        /// <summary>
        /// A map of sensitive data to pass to the endpoint. Useful for templated connection strings.
        /// 
        /// Supported list of database secrets engines that can be configured:
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Data;
        /// <summary>
        /// Set to true to disconnect any open sessions prior to running the revocation statements.
        /// </summary>
        public readonly bool? DisconnectSessions;
        /// <summary>
        /// Maximum number of seconds a connection may be reused.
        /// </summary>
        public readonly int? MaxConnectionLifetime;
        /// <summary>
        /// Maximum number of idle connections to the database.
        /// </summary>
        public readonly int? MaxIdleConnections;
        /// <summary>
        /// Maximum number of open connections to the database.
        /// </summary>
        public readonly int? MaxOpenConnections;
        /// <summary>
        /// Name of the database connection.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The root credential password used in the connection URL
        /// </summary>
        public readonly string? Password;
        /// <summary>
        /// Specifies the name of the plugin to use.
        /// </summary>
        public readonly string? PluginName;
        /// <summary>
        /// A list of database statements to be executed to rotate the root user's credentials.
        /// </summary>
        public readonly ImmutableArray<string> RootRotationStatements;
        /// <summary>
        /// Set to true in order to split statements after semi-colons.
        /// </summary>
        public readonly bool? SplitStatements;
        /// <summary>
        /// The root credential username used in the connection URL
        /// </summary>
        public readonly string? Username;
        /// <summary>
        /// Username generation template.
        /// </summary>
        public readonly string? UsernameTemplate;
        /// <summary>
        /// Whether the connection should be verified on
        /// initial configuration or not.
        /// </summary>
        public readonly bool? VerifyConnection;

        [OutputConstructor]
        private SecretsMountOracle(
            ImmutableArray<string> allowedRoles,

            string? connectionUrl,

            ImmutableDictionary<string, string>? data,

            bool? disconnectSessions,

            int? maxConnectionLifetime,

            int? maxIdleConnections,

            int? maxOpenConnections,

            string name,

            string? password,

            string? pluginName,

            ImmutableArray<string> rootRotationStatements,

            bool? splitStatements,

            string? username,

            string? usernameTemplate,

            bool? verifyConnection)
        {
            AllowedRoles = allowedRoles;
            ConnectionUrl = connectionUrl;
            Data = data;
            DisconnectSessions = disconnectSessions;
            MaxConnectionLifetime = maxConnectionLifetime;
            MaxIdleConnections = maxIdleConnections;
            MaxOpenConnections = maxOpenConnections;
            Name = name;
            Password = password;
            PluginName = pluginName;
            RootRotationStatements = rootRotationStatements;
            SplitStatements = splitStatements;
            Username = username;
            UsernameTemplate = usernameTemplate;
            VerifyConnection = verifyConnection;
        }
    }
}
