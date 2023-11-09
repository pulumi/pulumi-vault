// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vault.Database.Inputs
{

    public sealed class SecretsMountRediArgs : global::Pulumi.ResourceArgs
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
        /// The path to a PEM-encoded CA cert file to use to verify the Elasticsearch server's identity.
        /// </summary>
        [Input("caCert")]
        public Input<string>? CaCert { get; set; }

        [Input("data")]
        private InputMap<object>? _data;

        /// <summary>
        /// A map of sensitive data to pass to the endpoint. Useful for templated connection strings.
        /// 
        /// Supported list of database secrets engines that can be configured:
        /// </summary>
        public InputMap<object> Data
        {
            get => _data ?? (_data = new InputMap<object>());
            set => _data = value;
        }

        /// <summary>
        /// The host to connect to.
        /// </summary>
        [Input("host", required: true)]
        public Input<string> Host { get; set; } = null!;

        /// <summary>
        /// Whether to skip verification of the server
        /// certificate when using TLS.
        /// </summary>
        [Input("insecureTls")]
        public Input<bool>? InsecureTls { get; set; }

        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        [Input("password", required: true)]
        private Input<string>? _password;

        /// <summary>
        /// The root credential password used in the connection URL.
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
        /// The default port to connect to if no port is specified as
        /// part of the host.
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
        /// Whether to use TLS when connecting to Cassandra.
        /// </summary>
        [Input("tls")]
        public Input<bool>? Tls { get; set; }

        /// <summary>
        /// The root credential username used in the connection URL.
        /// </summary>
        [Input("username", required: true)]
        public Input<string> Username { get; set; } = null!;

        /// <summary>
        /// Whether the connection should be verified on
        /// initial configuration or not.
        /// </summary>
        [Input("verifyConnection")]
        public Input<bool>? VerifyConnection { get; set; }

        public SecretsMountRediArgs()
        {
        }
        public static new SecretsMountRediArgs Empty => new SecretsMountRediArgs();
    }
}
