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
    public sealed class SecretsMountCouchbase
    {
        /// <summary>
        /// A list of roles that are allowed to use this
        /// connection.
        /// </summary>
        public readonly ImmutableArray<string> AllowedRoles;
        /// <summary>
        /// Required if `tls` is `true`. Specifies the certificate authority of the Couchbase server, as a PEM certificate that has been base64 encoded.
        /// </summary>
        public readonly string? Base64Pem;
        /// <summary>
        /// Required for Couchbase versions prior to 6.5.0. This is only used to verify vault's connection to the server.
        /// </summary>
        public readonly string? BucketName;
        /// <summary>
        /// A map of sensitive data to pass to the endpoint. Useful for templated connection strings.
        /// 
        /// Supported list of database secrets engines that can be configured:
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Data;
        /// <summary>
        /// A set of Couchbase URIs to connect to. Must use `couchbases://` scheme if `tls` is `true`.
        /// </summary>
        public readonly ImmutableArray<string> Hosts;
        /// <summary>
        /// Specifies whether to skip verification of the server certificate when using TLS.
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
        /// Specifies the name of the plugin to use.
        /// </summary>
        public readonly string? PluginName;
        /// <summary>
        /// A list of database statements to be executed to rotate the root user's credentials.
        /// </summary>
        public readonly ImmutableArray<string> RootRotationStatements;
        /// <summary>
        /// Specifies whether to use TLS when connecting to Couchbase.
        /// </summary>
        public readonly bool? Tls;
        /// <summary>
        /// Specifies the username for Vault to use.
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
        private SecretsMountCouchbase(
            ImmutableArray<string> allowedRoles,

            string? base64Pem,

            string? bucketName,

            ImmutableDictionary<string, string>? data,

            ImmutableArray<string> hosts,

            bool? insecureTls,

            string name,

            string password,

            string? pluginName,

            ImmutableArray<string> rootRotationStatements,

            bool? tls,

            string username,

            string? usernameTemplate,

            bool? verifyConnection)
        {
            AllowedRoles = allowedRoles;
            Base64Pem = base64Pem;
            BucketName = bucketName;
            Data = data;
            Hosts = hosts;
            InsecureTls = insecureTls;
            Name = name;
            Password = password;
            PluginName = pluginName;
            RootRotationStatements = rootRotationStatements;
            Tls = tls;
            Username = username;
            UsernameTemplate = usernameTemplate;
            VerifyConnection = verifyConnection;
        }
    }
}
