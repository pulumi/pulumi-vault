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
    public sealed class SecretBackendConnectionMysql
    {
        /// <summary>
        /// Specify alternative authorization type. (Only 'gcp_iam' is valid currently)
        /// </summary>
        public readonly string? AuthType;
        /// <summary>
        /// Connection string to use to connect to the database.
        /// </summary>
        public readonly string? ConnectionUrl;
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
        /// The root credential password used in the connection URL
        /// </summary>
        public readonly string? Password;
        /// <summary>
        /// A JSON encoded credential for use with IAM authorization
        /// </summary>
        public readonly string? ServiceAccountJson;
        /// <summary>
        /// x509 CA file for validating the certificate presented by the MySQL server. Must be PEM encoded.
        /// </summary>
        public readonly string? TlsCa;
        /// <summary>
        /// x509 certificate for connecting to the database. This must be a PEM encoded version of the private key and the certificate combined.
        /// </summary>
        public readonly string? TlsCertificateKey;
        /// <summary>
        /// The root credential username used in the connection URL
        /// </summary>
        public readonly string? Username;
        /// <summary>
        /// Username generation template.
        /// </summary>
        public readonly string? UsernameTemplate;

        [OutputConstructor]
        private SecretBackendConnectionMysql(
            string? authType,

            string? connectionUrl,

            int? maxConnectionLifetime,

            int? maxIdleConnections,

            int? maxOpenConnections,

            string? password,

            string? serviceAccountJson,

            string? tlsCa,

            string? tlsCertificateKey,

            string? username,

            string? usernameTemplate)
        {
            AuthType = authType;
            ConnectionUrl = connectionUrl;
            MaxConnectionLifetime = maxConnectionLifetime;
            MaxIdleConnections = maxIdleConnections;
            MaxOpenConnections = maxOpenConnections;
            Password = password;
            ServiceAccountJson = serviceAccountJson;
            TlsCa = tlsCa;
            TlsCertificateKey = tlsCertificateKey;
            Username = username;
            UsernameTemplate = usernameTemplate;
        }
    }
}
