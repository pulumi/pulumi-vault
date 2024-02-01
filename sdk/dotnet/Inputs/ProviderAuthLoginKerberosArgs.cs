// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vault.Inputs
{

    public sealed class ProviderAuthLoginKerberosArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Disable the Kerberos FAST negotiation.
        /// </summary>
        [Input("disableFastNegotiation")]
        public Input<bool>? DisableFastNegotiation { get; set; }

        /// <summary>
        /// The Kerberos keytab file containing the entry of the login entity.
        /// </summary>
        [Input("keytabPath")]
        public Input<string>? KeytabPath { get; set; }

        /// <summary>
        /// A valid Kerberos configuration file e.g. /etc/krb5.conf.
        /// </summary>
        [Input("krb5confPath")]
        public Input<string>? Krb5confPath { get; set; }

        /// <summary>
        /// The path where the authentication engine is mounted.
        /// </summary>
        [Input("mount")]
        public Input<string>? Mount { get; set; }

        /// <summary>
        /// The authentication engine's namespace. Conflicts with use_root_namespace
        /// </summary>
        [Input("namespace")]
        public Input<string>? Namespace { get; set; }

        /// <summary>
        /// The Kerberos server's authoritative authentication domain
        /// </summary>
        [Input("realm")]
        public Input<string>? Realm { get; set; }

        /// <summary>
        /// Strip the host from the username found in the keytab.
        /// </summary>
        [Input("removeInstanceName")]
        public Input<bool>? RemoveInstanceName { get; set; }

        /// <summary>
        /// The service principle name.
        /// </summary>
        [Input("service")]
        public Input<string>? Service { get; set; }

        /// <summary>
        /// Simple and Protected GSSAPI Negotiation Mechanism (SPNEGO) token
        /// </summary>
        [Input("token")]
        public Input<string>? Token { get; set; }

        /// <summary>
        /// Authenticate to the root Vault namespace. Conflicts with namespace
        /// </summary>
        [Input("useRootNamespace")]
        public Input<bool>? UseRootNamespace { get; set; }

        /// <summary>
        /// The username to login into Kerberos with.
        /// </summary>
        [Input("username")]
        public Input<string>? Username { get; set; }

        public ProviderAuthLoginKerberosArgs()
        {
        }
        public static new ProviderAuthLoginKerberosArgs Empty => new ProviderAuthLoginKerberosArgs();
    }
}
