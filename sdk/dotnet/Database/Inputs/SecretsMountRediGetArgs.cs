// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vault.Database.Inputs
{

    public sealed class SecretsMountRediGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("allowedRoles")]
        private InputList<string>? _allowedRoles;

        /// <summary>
        /// A list of roles that are allowed to use this
        /// connection.
        /// </summary>
        public InputList<string> AllowedRoles
        {
            get => _allowedRoles ?? (_allowedRoles = new InputList<string>());
            set => _allowedRoles = value;
        }

        /// <summary>
        /// The contents of a PEM-encoded CA cert file to use to verify the Redis server's identity.
        /// </summary>
        [Input("caCert")]
        public Input<string>? CaCert { get; set; }

        [Input("data")]
        private InputMap<string>? _data;

        /// <summary>
        /// A map of sensitive data to pass to the endpoint. Useful for templated connection strings.
        /// </summary>
        public InputMap<string> Data
        {
            get => _data ?? (_data = new InputMap<string>());
            set => _data = value;
        }

        /// <summary>
        /// Cancels all upcoming rotations of the root credential until unset. Requires Vault Enterprise 1.19+.
        /// 
        /// Supported list of database secrets engines that can be configured:
        /// </summary>
        [Input("disableAutomatedRotation")]
        public Input<bool>? DisableAutomatedRotation { get; set; }

        /// <summary>
        /// Specifies the host to connect to
        /// </summary>
        [Input("host", required: true)]
        public Input<string> Host { get; set; } = null!;

        /// <summary>
        /// Specifies whether to skip verification of the server certificate when using TLS.
        /// </summary>
        [Input("insecureTls")]
        public Input<bool>? InsecureTls { get; set; }

        /// <summary>
        /// Name of the database connection.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        [Input("password", required: true)]
        private Input<string>? _password;

        /// <summary>
        /// Specifies the password corresponding to the given username.
        /// </summary>
        public Input<string>? Password
        {
            get => _password;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _password = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// Specifies the name of the plugin to use.
        /// </summary>
        [Input("pluginName")]
        public Input<string>? PluginName { get; set; }

        /// <summary>
        /// The transport port to use to connect to Redis.
        /// </summary>
        [Input("port")]
        public Input<int>? Port { get; set; }

        [Input("rootRotationStatements")]
        private InputList<string>? _rootRotationStatements;

        /// <summary>
        /// A list of database statements to be executed to rotate the root user's credentials.
        /// </summary>
        public InputList<string> RootRotationStatements
        {
            get => _rootRotationStatements ?? (_rootRotationStatements = new InputList<string>());
            set => _rootRotationStatements = value;
        }

        /// <summary>
        /// The amount of time in seconds Vault should wait before rotating the root credential.
        /// A zero value tells Vault not to rotate the root credential. The minimum rotation period is 10 seconds. Requires Vault Enterprise 1.19+.
        /// </summary>
        [Input("rotationPeriod")]
        public Input<int>? RotationPeriod { get; set; }

        /// <summary>
        /// The schedule, in [cron-style time format](https://en.wikipedia.org/wiki/Cron),
        /// defining the schedule on which Vault should rotate the root token. Requires Vault Enterprise 1.19+.
        /// </summary>
        [Input("rotationSchedule")]
        public Input<string>? RotationSchedule { get; set; }

        /// <summary>
        /// The maximum amount of time in seconds allowed to complete
        /// a rotation when a scheduled token rotation occurs. The default rotation window is
        /// unbound and the minimum allowable window is `3600`. Requires Vault Enterprise 1.19+.
        /// </summary>
        [Input("rotationWindow")]
        public Input<int>? RotationWindow { get; set; }

        /// <summary>
        /// Specifies whether to use TLS when connecting to Redis.
        /// </summary>
        [Input("tls")]
        public Input<bool>? Tls { get; set; }

        /// <summary>
        /// Specifies the username for Vault to use.
        /// </summary>
        [Input("username", required: true)]
        public Input<string> Username { get; set; } = null!;

        /// <summary>
        /// Whether the connection should be verified on
        /// initial configuration or not.
        /// </summary>
        [Input("verifyConnection")]
        public Input<bool>? VerifyConnection { get; set; }

        public SecretsMountRediGetArgs()
        {
        }
        public static new SecretsMountRediGetArgs Empty => new SecretsMountRediGetArgs();
    }
}
