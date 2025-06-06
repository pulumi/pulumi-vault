// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vault.PkiSecret.Inputs
{

    public sealed class BackendConfigCmpv2AuthenticatorsGetArgs : global::Pulumi.ResourceArgs
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

        public BackendConfigCmpv2AuthenticatorsGetArgs()
        {
        }
        public static new BackendConfigCmpv2AuthenticatorsGetArgs Empty => new BackendConfigCmpv2AuthenticatorsGetArgs();
    }
}
