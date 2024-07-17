// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vault.Inputs
{

    public sealed class GetPolicyDocumentRuleDeniedParameterArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// name of permitted or denied parameter.
        /// </summary>
        [Input("key", required: true)]
        public string Key { get; set; } = null!;

        [Input("values", required: true)]
        private List<string>? _values;

        /// <summary>
        /// list of values what are permitted or denied by policy rule.
        /// </summary>
        public List<string> Values
        {
            get => _values ?? (_values = new List<string>());
            set => _values = value;
        }

        public GetPolicyDocumentRuleDeniedParameterArgs()
        {
        }
        public static new GetPolicyDocumentRuleDeniedParameterArgs Empty => new GetPolicyDocumentRuleDeniedParameterArgs();
    }
}
