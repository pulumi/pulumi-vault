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
    public sealed class SecretBackendConnectionOracle
    {
        /// <summary>
        /// A URL containing connection information. See
        /// the [Vault
        /// docs](https://www.vaultproject.io/api-docs/secret/databases/mongodb.html#sample-payload)
        /// for an example.
        /// </summary>
        public readonly string? ConnectionUrl;
        /// <summary>
        /// Enable the built-in session disconnect mechanism.
        /// </summary>
        public readonly bool? DisconnectSessions;
        /// <summary>
        /// The maximum number of seconds to keep
        /// a connection alive for.
        /// </summary>
        public readonly int? MaxConnectionLifetime;
        /// <summary>
        /// The maximum number of idle connections to
        /// maintain.
        /// </summary>
        public readonly int? MaxIdleConnections;
        /// <summary>
        /// The maximum number of open connections to
        /// use.
        /// </summary>
        public readonly int? MaxOpenConnections;
        /// <summary>
        /// The password to authenticate with.
        /// </summary>
        public readonly string? Password;
        /// <summary>
        /// Enable spliting statements after semi-colons.
        /// </summary>
        public readonly bool? SplitStatements;
        /// <summary>
        /// The username to authenticate with.
        /// </summary>
        public readonly string? Username;
        /// <summary>
        /// Template describing how dynamic usernames are generated.
        /// </summary>
        public readonly string? UsernameTemplate;

        [OutputConstructor]
        private SecretBackendConnectionOracle(
            string? connectionUrl,

            bool? disconnectSessions,

            int? maxConnectionLifetime,

            int? maxIdleConnections,

            int? maxOpenConnections,

            string? password,

            bool? splitStatements,

            string? username,

            string? usernameTemplate)
        {
            ConnectionUrl = connectionUrl;
            DisconnectSessions = disconnectSessions;
            MaxConnectionLifetime = maxConnectionLifetime;
            MaxIdleConnections = maxIdleConnections;
            MaxOpenConnections = maxOpenConnections;
            Password = password;
            SplitStatements = splitStatements;
            Username = username;
            UsernameTemplate = usernameTemplate;
        }
    }
}
