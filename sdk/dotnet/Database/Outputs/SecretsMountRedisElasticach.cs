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
    public sealed class SecretsMountRedisElasticach
    {
        /// <summary>
        /// A list of roles that are allowed to use this
        /// connection.
        /// </summary>
        public readonly ImmutableArray<string> AllowedRoles;
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
        /// Name of the database connection.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The AWS secret key id to use to talk to ElastiCache. If omitted the credentials chain provider is used instead.
        /// </summary>
        public readonly string? Password;
        /// <summary>
        /// Specifies the name of the plugin to use.
        /// </summary>
        public readonly string? PluginName;
        /// <summary>
        /// The AWS region where the ElastiCache cluster is hosted. If omitted the plugin tries to infer the region from the environment.
        /// </summary>
        public readonly string? Region;
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
        /// The configuration endpoint for the ElastiCache cluster to connect to.
        /// </summary>
        public readonly string Url;
        /// <summary>
        /// The AWS access key id to use to talk to ElastiCache. If omitted the credentials chain provider is used instead.
        /// </summary>
        public readonly string? Username;
        /// <summary>
        /// Whether the connection should be verified on
        /// initial configuration or not.
        /// </summary>
        public readonly bool? VerifyConnection;

        [OutputConstructor]
        private SecretsMountRedisElasticach(
            ImmutableArray<string> allowedRoles,

            ImmutableDictionary<string, string>? data,

            bool? disableAutomatedRotation,

            string name,

            string? password,

            string? pluginName,

            string? region,

            ImmutableArray<string> rootRotationStatements,

            int? rotationPeriod,

            string? rotationSchedule,

            int? rotationWindow,

            string url,

            string? username,

            bool? verifyConnection)
        {
            AllowedRoles = allowedRoles;
            Data = data;
            DisableAutomatedRotation = disableAutomatedRotation;
            Name = name;
            Password = password;
            PluginName = pluginName;
            Region = region;
            RootRotationStatements = rootRotationStatements;
            RotationPeriod = rotationPeriod;
            RotationSchedule = rotationSchedule;
            RotationWindow = rotationWindow;
            Url = url;
            Username = username;
            VerifyConnection = verifyConnection;
        }
    }
}
