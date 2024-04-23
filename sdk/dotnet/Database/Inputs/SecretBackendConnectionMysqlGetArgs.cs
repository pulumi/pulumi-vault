// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vault.Database.Inputs
{

    public sealed class SecretBackendConnectionMysqlGetArgs : global::Pulumi.ResourceArgs
    {
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

        public SecretBackendConnectionMysqlGetArgs()
        {
        }
        public static new SecretBackendConnectionMysqlGetArgs Empty => new SecretBackendConnectionMysqlGetArgs();
    }
}
