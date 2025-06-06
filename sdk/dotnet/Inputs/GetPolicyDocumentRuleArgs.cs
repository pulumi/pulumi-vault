// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vault.Inputs
{

    public sealed class GetPolicyDocumentRuleInputArgs : global::Pulumi.ResourceArgs
    {
        [Input("allowedParameters")]
        private InputList<Inputs.GetPolicyDocumentRuleAllowedParameterInputArgs>? _allowedParameters;

        /// <summary>
        /// Whitelists a list of keys and values that are permitted on the given path. See Parameters below.
        /// </summary>
        public InputList<Inputs.GetPolicyDocumentRuleAllowedParameterInputArgs> AllowedParameters
        {
            get => _allowedParameters ?? (_allowedParameters = new InputList<Inputs.GetPolicyDocumentRuleAllowedParameterInputArgs>());
            set => _allowedParameters = value;
        }

        [Input("capabilities", required: true)]
        private InputList<string>? _capabilities;

        /// <summary>
        /// A list of capabilities that this rule apply to `path`. For example, ["read", "write"].
        /// </summary>
        public InputList<string> Capabilities
        {
            get => _capabilities ?? (_capabilities = new InputList<string>());
            set => _capabilities = value;
        }

        [Input("deniedParameters")]
        private InputList<Inputs.GetPolicyDocumentRuleDeniedParameterInputArgs>? _deniedParameters;

        /// <summary>
        /// Blacklists a list of parameter and values. Any values specified here take precedence over `allowed_parameter`. See Parameters below.
        /// </summary>
        public InputList<Inputs.GetPolicyDocumentRuleDeniedParameterInputArgs> DeniedParameters
        {
            get => _deniedParameters ?? (_deniedParameters = new InputList<Inputs.GetPolicyDocumentRuleDeniedParameterInputArgs>());
            set => _deniedParameters = value;
        }

        /// <summary>
        /// Description of the rule. Will be added as a comment to rendered rule.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The maximum allowed TTL that clients can specify for a wrapped response.
        /// </summary>
        [Input("maxWrappingTtl")]
        public Input<string>? MaxWrappingTtl { get; set; }

        /// <summary>
        /// The minimum allowed TTL that clients can specify for a wrapped response.
        /// </summary>
        [Input("minWrappingTtl")]
        public Input<string>? MinWrappingTtl { get; set; }

        /// <summary>
        /// A path in Vault that this rule applies to.
        /// </summary>
        [Input("path", required: true)]
        public Input<string> Path { get; set; } = null!;

        [Input("requiredParameters")]
        private InputList<string>? _requiredParameters;

        /// <summary>
        /// A list of parameters that must be specified.
        /// </summary>
        public InputList<string> RequiredParameters
        {
            get => _requiredParameters ?? (_requiredParameters = new InputList<string>());
            set => _requiredParameters = value;
        }

        [Input("subscribeEventTypes")]
        private InputList<string>? _subscribeEventTypes;

        /// <summary>
        /// A list of event types to subscribe to when using `subscribe` capability.
        /// </summary>
        public InputList<string> SubscribeEventTypes
        {
            get => _subscribeEventTypes ?? (_subscribeEventTypes = new InputList<string>());
            set => _subscribeEventTypes = value;
        }

        public GetPolicyDocumentRuleInputArgs()
        {
        }
        public static new GetPolicyDocumentRuleInputArgs Empty => new GetPolicyDocumentRuleInputArgs();
    }
}
