// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vault.Database.Inputs
{

    public sealed class SecretBackendConnectionInfluxdbGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The number of seconds to use as a connection timeout.
        /// </summary>
        [Input("connectTimeout")]
        public Input<int>? ConnectTimeout { get; set; }

        /// <summary>
        /// Influxdb host to connect to.
        /// </summary>
        [Input("host", required: true)]
        public Input<string> Host { get; set; } = null!;

        /// <summary>
        /// Whether to skip verification of the server certificate when using TLS.
        /// </summary>
        [Input("insecureTls")]
        public Input<bool>? InsecureTls { get; set; }

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
        /// The transport port to use to connect to Influxdb.
        /// </summary>
        [Input("port")]
        public Input<int>? Port { get; set; }

        /// <summary>
        /// Whether to use TLS when connecting to Influxdb.
        /// </summary>
        [Input("tls")]
        public Input<bool>? Tls { get; set; }

        /// <summary>
        /// Specifies the username to use for superuser access.
        /// </summary>
        [Input("username", required: true)]
        public Input<string> Username { get; set; } = null!;

        /// <summary>
        /// Template describing how dynamic usernames are generated.
        /// </summary>
        [Input("usernameTemplate")]
        public Input<string>? UsernameTemplate { get; set; }

        public SecretBackendConnectionInfluxdbGetArgs()
        {
        }
        public static new SecretBackendConnectionInfluxdbGetArgs Empty => new SecretBackendConnectionInfluxdbGetArgs();
    }
}
