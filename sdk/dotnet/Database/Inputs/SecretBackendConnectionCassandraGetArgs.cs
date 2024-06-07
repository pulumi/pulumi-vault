// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vault.Database.Inputs
{

    public sealed class SecretBackendConnectionCassandraGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The number of seconds to use as a connection timeout.
        /// </summary>
        [Input("connectTimeout")]
        public Input<int>? ConnectTimeout { get; set; }

        [Input("hosts")]
        private InputList<string>? _hosts;

        /// <summary>
        /// Cassandra hosts to connect to.
        /// </summary>
        public InputList<string> Hosts
        {
            get => _hosts ?? (_hosts = new InputList<string>());
            set => _hosts = value;
        }

        /// <summary>
        /// Whether to skip verification of the server certificate when using TLS.
        /// </summary>
        [Input("insecureTls")]
        public Input<bool>? InsecureTls { get; set; }

        [Input("password")]
        private Input<string>? _password;

        /// <summary>
        /// The password to use when authenticating with Cassandra.
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

        [Input("pemBundle")]
        private Input<string>? _pemBundle;

        /// <summary>
        /// Concatenated PEM blocks containing a certificate and private key; a certificate, private key, and issuing CA certificate; or just a CA certificate.
        /// </summary>
        public Input<string>? PemBundle
        {
            get => _pemBundle;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _pemBundle = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        [Input("pemJson")]
        private Input<string>? _pemJson;

        /// <summary>
        /// Specifies JSON containing a certificate and private key; a certificate, private key, and issuing CA certificate; or just a CA certificate.
        /// </summary>
        public Input<string>? PemJson
        {
            get => _pemJson;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _pemJson = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// The transport port to use to connect to Cassandra.
        /// </summary>
        [Input("port")]
        public Input<int>? Port { get; set; }

        /// <summary>
        /// The CQL protocol version to use.
        /// </summary>
        [Input("protocolVersion")]
        public Input<int>? ProtocolVersion { get; set; }

        /// <summary>
        /// Whether to use TLS when connecting to Cassandra.
        /// </summary>
        [Input("tls")]
        public Input<bool>? Tls { get; set; }

        /// <summary>
        /// The username to use when authenticating with Cassandra.
        /// </summary>
        [Input("username")]
        public Input<string>? Username { get; set; }

        public SecretBackendConnectionCassandraGetArgs()
        {
        }
        public static new SecretBackendConnectionCassandraGetArgs Empty => new SecretBackendConnectionCassandraGetArgs();
    }
}
