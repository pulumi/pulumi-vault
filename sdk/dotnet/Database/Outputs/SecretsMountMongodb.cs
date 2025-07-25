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
    public sealed class SecretsMountMongodb
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
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Data;
        /// <summary>
        /// Cancels all upcoming rotations of the root credential until unset. Requires Vault Enterprise 1.19+.
        /// 
        /// Supported list of database secrets engines that can be configured:
        /// </summary>
        public readonly bool? DisableAutomatedRotation;
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
        /// Version counter for root credential password write-only field
        /// </summary>
        public readonly int? PasswordWoVersion;
        /// <summary>
        /// Specifies the name of the plugin to use.
        /// </summary>
        public readonly string? PluginName;
        /// <summary>
        /// A list of database statements to be executed to rotate the root user's credentials.
        /// </summary>
        public readonly ImmutableArray<string> RootRotationStatements;
        /// <summary>
        /// The amount of time in seconds Vault should wait before rotating the root credential.
        /// A zero value tells Vault not to rotate the root credential. The minimum rotation period is 10 seconds. Requires Vault Enterprise 1.19+.
        /// </summary>
        public readonly int? RotationPeriod;
        /// <summary>
        /// The schedule, in [cron-style time format](https://en.wikipedia.org/wiki/Cron),
        /// defining the schedule on which Vault should rotate the root token. Requires Vault Enterprise 1.19+.
        /// </summary>
        public readonly string? RotationSchedule;
        /// <summary>
        /// The maximum amount of time in seconds allowed to complete
        /// a rotation when a scheduled token rotation occurs. The default rotation window is
        /// unbound and the minimum allowable window is `3600`. Requires Vault Enterprise 1.19+.
        /// </summary>
        public readonly int? RotationWindow;
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
        private SecretsMountMongodb(
            ImmutableArray<string> allowedRoles,

            string? connectionUrl,

            ImmutableDictionary<string, string>? data,

            bool? disableAutomatedRotation,

            int? maxConnectionLifetime,

            int? maxIdleConnections,

            int? maxOpenConnections,

            string name,

            string? password,

            int? passwordWoVersion,

            string? pluginName,

            ImmutableArray<string> rootRotationStatements,

            int? rotationPeriod,

            string? rotationSchedule,

            int? rotationWindow,

            string? username,

            string? usernameTemplate,

            bool? verifyConnection)
        {
            AllowedRoles = allowedRoles;
            ConnectionUrl = connectionUrl;
            Data = data;
            DisableAutomatedRotation = disableAutomatedRotation;
            MaxConnectionLifetime = maxConnectionLifetime;
            MaxIdleConnections = maxIdleConnections;
            MaxOpenConnections = maxOpenConnections;
            Name = name;
            Password = password;
            PasswordWoVersion = passwordWoVersion;
            PluginName = pluginName;
            RootRotationStatements = rootRotationStatements;
            RotationPeriod = rotationPeriod;
            RotationSchedule = rotationSchedule;
            RotationWindow = rotationWindow;
            Username = username;
            UsernameTemplate = usernameTemplate;
            VerifyConnection = verifyConnection;
        }
    }
}
