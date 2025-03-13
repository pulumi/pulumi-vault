// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vault.Database.Inputs
{

    public sealed class SecretsMountElasticsearchGetArgs : global::Pulumi.ResourceArgs
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
        /// The path to a PEM-encoded CA cert file to use to verify the Elasticsearch server's identity
        /// </summary>
        [Input("caCert")]
        public Input<string>? CaCert { get; set; }

        /// <summary>
        /// The path to a directory of PEM-encoded CA cert files to use to verify the Elasticsearch server's identity
        /// </summary>
        [Input("caPath")]
        public Input<string>? CaPath { get; set; }

        /// <summary>
        /// The path to the certificate for the Elasticsearch client to present for communication
        /// </summary>
        [Input("clientCert")]
        public Input<string>? ClientCert { get; set; }

        /// <summary>
        /// The path to the key for the Elasticsearch client to use for communication
        /// </summary>
        [Input("clientKey")]
        public Input<string>? ClientKey { get; set; }

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
        /// Whether to disable certificate verification
        /// </summary>
        [Input("insecure")]
        public Input<bool>? Insecure { get; set; }

        /// <summary>
        /// Name of the database connection.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        [Input("password", required: true)]
        private Input<string>? _password;

        /// <summary>
        /// The password to be used in the connection URL
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
        /// This, if set, is used to set the SNI host when connecting via TLS
        /// </summary>
        [Input("tlsServerName")]
        public Input<string>? TlsServerName { get; set; }

        /// <summary>
        /// The URL for Elasticsearch's API
        /// </summary>
        [Input("url", required: true)]
        public Input<string> Url { get; set; } = null!;

        /// <summary>
        /// The username to be used in the connection URL
        /// </summary>
        [Input("username", required: true)]
        public Input<string> Username { get; set; } = null!;

        /// <summary>
        /// Template describing how dynamic usernames are generated.
        /// </summary>
        [Input("usernameTemplate")]
        public Input<string>? UsernameTemplate { get; set; }

        /// <summary>
        /// Whether the connection should be verified on
        /// initial configuration or not.
        /// </summary>
        [Input("verifyConnection")]
        public Input<bool>? VerifyConnection { get; set; }

        public SecretsMountElasticsearchGetArgs()
        {
        }
        public static new SecretsMountElasticsearchGetArgs Empty => new SecretsMountElasticsearchGetArgs();
    }
}
