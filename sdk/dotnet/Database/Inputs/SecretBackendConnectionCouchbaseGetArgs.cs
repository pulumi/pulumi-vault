// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vault.Database.Inputs
{

    public sealed class SecretBackendConnectionCouchbaseGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("base64Pem")]
        private Input<string>? _base64Pem;

        /// <summary>
        /// Required if `tls` is `true`. Specifies the certificate authority of the Couchbase server, as a PEM certificate that has been base64 encoded.
        /// </summary>
        public Input<string>? Base64Pem
        {
            get => _base64Pem;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _base64Pem = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// Required for Couchbase versions prior to 6.5.0. This is only used to verify vault's connection to the server.
        /// </summary>
        [Input("bucketName")]
        public Input<string>? BucketName { get; set; }

        [Input("hosts", required: true)]
        private InputList<string>? _hosts;

        /// <summary>
        /// A set of Couchbase URIs to connect to. Must use `couchbases://` scheme if `tls` is `true`.
        /// </summary>
        public InputList<string> Hosts
        {
            get => _hosts ?? (_hosts = new InputList<string>());
            set => _hosts = value;
        }

        /// <summary>
        /// Whether to skip verification of the server
        /// certificate when using TLS.
        /// </summary>
        [Input("insecureTls")]
        public Input<bool>? InsecureTls { get; set; }

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
        /// Whether to use TLS when connecting to Redis.
        /// </summary>
        [Input("tls")]
        public Input<bool>? Tls { get; set; }

        /// <summary>
        /// The root credential username used in the connection URL.
        /// </summary>
        [Input("username", required: true)]
        public Input<string> Username { get; set; } = null!;

        /// <summary>
        /// - [Template](https://www.vaultproject.io/docs/concepts/username-templating) describing how dynamic usernames are generated.
        /// </summary>
        [Input("usernameTemplate")]
        public Input<string>? UsernameTemplate { get; set; }

        public SecretBackendConnectionCouchbaseGetArgs()
        {
        }
        public static new SecretBackendConnectionCouchbaseGetArgs Empty => new SecretBackendConnectionCouchbaseGetArgs();
    }
}
