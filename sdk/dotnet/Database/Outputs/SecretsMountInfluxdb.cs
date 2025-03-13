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
    public sealed class SecretsMountInfluxdb
    {
        /// <summary>
        /// A list of roles that are allowed to use this
        /// connection.
        /// </summary>
        public readonly ImmutableArray<string> AllowedRoles;
        /// <summary>
        /// The number of seconds to use as a connection timeout.
        /// </summary>
        public readonly int? ConnectTimeout;
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
        /// Influxdb host to connect to.
        /// </summary>
        public readonly string Host;
        /// <summary>
        /// Whether to skip verification of the server certificate when using TLS.
        /// </summary>
        public readonly bool? InsecureTls;
        /// <summary>
        /// Name of the database connection.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Specifies the password corresponding to the given username.
        /// </summary>
        public readonly string Password;
        /// <summary>
        /// Concatenated PEM blocks containing a certificate and private key; a certificate, private key, and issuing CA certificate; or just a CA certificate.
        /// </summary>
        public readonly string? PemBundle;
        /// <summary>
        /// Specifies JSON containing a certificate and private key; a certificate, private key, and issuing CA certificate; or just a CA certificate.
        /// </summary>
        public readonly string? PemJson;
        /// <summary>
        /// Specifies the name of the plugin to use.
        /// </summary>
        public readonly string? PluginName;
        /// <summary>
        /// The transport port to use to connect to Influxdb.
        /// </summary>
        public readonly int? Port;
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
        /// Whether to use TLS when connecting to Influxdb.
        /// </summary>
        public readonly bool? Tls;
        /// <summary>
        /// Specifies the username to use for superuser access.
        /// </summary>
        public readonly string Username;
        /// <summary>
        /// Template describing how dynamic usernames are generated.
        /// </summary>
        public readonly string? UsernameTemplate;
        /// <summary>
        /// Whether the connection should be verified on
        /// initial configuration or not.
        /// </summary>
        public readonly bool? VerifyConnection;

        [OutputConstructor]
        private SecretsMountInfluxdb(
            ImmutableArray<string> allowedRoles,

            int? connectTimeout,

            ImmutableDictionary<string, string>? data,

            bool? disableAutomatedRotation,

            string host,

            bool? insecureTls,

            string name,

            string password,

            string? pemBundle,

            string? pemJson,

            string? pluginName,

            int? port,

            ImmutableArray<string> rootRotationStatements,

            int? rotationPeriod,

            string? rotationSchedule,

            int? rotationWindow,

            bool? tls,

            string username,

            string? usernameTemplate,

            bool? verifyConnection)
        {
            AllowedRoles = allowedRoles;
            ConnectTimeout = connectTimeout;
            Data = data;
            DisableAutomatedRotation = disableAutomatedRotation;
            Host = host;
            InsecureTls = insecureTls;
            Name = name;
            Password = password;
            PemBundle = pemBundle;
            PemJson = pemJson;
            PluginName = pluginName;
            Port = port;
            RootRotationStatements = rootRotationStatements;
            RotationPeriod = rotationPeriod;
            RotationSchedule = rotationSchedule;
            RotationWindow = rotationWindow;
            Tls = tls;
            Username = username;
            UsernameTemplate = usernameTemplate;
            VerifyConnection = verifyConnection;
        }
    }
}
