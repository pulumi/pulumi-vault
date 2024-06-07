// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vault.Inputs
{

    public sealed class GetPolicyDocumentRuleAllowedParameterArgs : global::Pulumi.InvokeArgs
    {
        [Input("key", required: true)]
        public string Key { get; set; } = null!;

        [Input("values", required: true)]
        private List<string>? _values;
        public List<string> Values
        {
            get => _values ?? (_values = new List<string>());
            set => _values = value;
        }

        public GetPolicyDocumentRuleAllowedParameterArgs()
        {
        }
        public static new GetPolicyDocumentRuleAllowedParameterArgs Empty => new GetPolicyDocumentRuleAllowedParameterArgs();
    }
}
