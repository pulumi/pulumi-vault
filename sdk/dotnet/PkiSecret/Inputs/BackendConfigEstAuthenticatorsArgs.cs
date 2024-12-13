// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vault.PkiSecret.Inputs
{

    public sealed class BackendConfigEstAuthenticatorsArgs : global::Pulumi.ResourceArgs
    {
        [Input("cert")]
        private InputMap<string>? _cert;

        /// <summary>
        /// "The accessor (required) and cert_role (optional) properties for cert auth backends".
        /// </summary>
        public InputMap<string> Cert
        {
            get => _cert ?? (_cert = new InputMap<string>());
            set => _cert = value;
        }

        [Input("userpass")]
        private InputMap<string>? _userpass;

        /// <summary>
        /// "The accessor (required) property for user pass auth backends".
        /// </summary>
        public InputMap<string> Userpass
        {
            get => _userpass ?? (_userpass = new InputMap<string>());
            set => _userpass = value;
        }

        public BackendConfigEstAuthenticatorsArgs()
        {
        }
        public static new BackendConfigEstAuthenticatorsArgs Empty => new BackendConfigEstAuthenticatorsArgs();
    }
}