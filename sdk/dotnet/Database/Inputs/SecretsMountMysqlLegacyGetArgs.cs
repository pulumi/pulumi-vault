// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vault.Database.Inputs
{

    public sealed class SecretsMountMysqlLegacyGetArgs : global::Pulumi.ResourceArgs
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
        /// Specify alternative authorization type. (Only 'gcp_iam' is valid currently)
        /// </summary>
        [Input("authType")]
        public Input<string>? AuthType { get; set; }

        /// <summary>
        /// Connection string to use to connect to the database.
        /// </summary>
        [Input("connectionUrl")]
        public Input<string>? ConnectionUrl { get; set; }

        [Input("data")]
        private InputMap<string>? _data;

        /// <summary>
        /// A map of sensitive data to pass to the endpoint. Useful for templated connection strings.
        /// 
        /// Supported list of database secrets engines that can be configured:
        /// </summary>
        public InputMap<string> Data
        {
            get => _data ?? (_data = new InputMap<string>());
            set => _data = value;
        }

        /// <summary>
        /// Maximum number of seconds a connection may be reused.
        /// </summary>
        [Input("maxConnectionLifetime")]
        public Input<int>? MaxConnectionLifetime { get; set; }

        /// <summary>
        /// Maximum number of idle connections to the database.
        /// </summary>
        [Input("maxIdleConnections")]
        public Input<int>? MaxIdleConnections { get; set; }

        /// <summary>
        /// Maximum number of open connections to the database.
        /// </summary>
        [Input("maxOpenConnections")]
        public Input<int>? MaxOpenConnections { get; set; }

        /// <summary>
        /// Name of the database connection.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        [Input("password")]
        private Input<string>? _password;

        /// <summary>
        /// The root credential password used in the connection URL
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

        [Input("serviceAccountJson")]
        private Input<string>? _serviceAccountJson;

        /// <summary>
        /// A JSON encoded credential for use with IAM authorization
        /// </summary>
        public Input<string>? ServiceAccountJson
        {
            get => _serviceAccountJson;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _serviceAccountJson = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// x509 CA file for validating the certificate presented by the MySQL server. Must be PEM encoded.
        /// </summary>
        [Input("tlsCa")]
        public Input<string>? TlsCa { get; set; }

        [Input("tlsCertificateKey")]
        private Input<string>? _tlsCertificateKey;

        /// <summary>
        /// x509 certificate for connecting to the database. This must be a PEM encoded version of the private key and the certificate combined.
        /// </summary>
        public Input<string>? TlsCertificateKey
        {
            get => _tlsCertificateKey;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _tlsCertificateKey = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// The root credential username used in the connection URL
        /// </summary>
        [Input("username")]
        public Input<string>? Username { get; set; }

        /// <summary>
        /// Username generation template.
        /// </summary>
        [Input("usernameTemplate")]
        public Input<string>? UsernameTemplate { get; set; }

        /// <summary>
        /// Whether the connection should be verified on
        /// initial configuration or not.
        /// </summary>
        [Input("verifyConnection")]
        public Input<bool>? VerifyConnection { get; set; }

        public SecretsMountMysqlLegacyGetArgs()
        {
        }
        public static new SecretsMountMysqlLegacyGetArgs Empty => new SecretsMountMysqlLegacyGetArgs();
    }
}
