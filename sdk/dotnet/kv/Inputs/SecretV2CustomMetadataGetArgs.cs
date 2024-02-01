// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vault.kv.Inputs
{

    public sealed class SecretV2CustomMetadataGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// If true, all keys will require the cas parameter to be set on all write requests.
        /// </summary>
        [Input("casRequired")]
        public Input<bool>? CasRequired { get; set; }

        [Input("data")]
        private InputMap<object>? _data;

        /// <summary>
        /// A mapping whose keys are the top-level data keys returned from
        /// Vault and whose values are the corresponding values. This map can only
        /// represent string data, so any non-string values returned from Vault are
        /// serialized as JSON.
        /// </summary>
        public InputMap<object> Data
        {
            get => _data ?? (_data = new InputMap<object>());
            set => _data = value;
        }

        /// <summary>
        /// If set, specifies the length of time before a version is deleted.
        /// </summary>
        [Input("deleteVersionAfter")]
        public Input<int>? DeleteVersionAfter { get; set; }

        /// <summary>
        /// The number of versions to keep per key.
        /// </summary>
        [Input("maxVersions")]
        public Input<int>? MaxVersions { get; set; }

        public SecretV2CustomMetadataGetArgs()
        {
        }
        public static new SecretV2CustomMetadataGetArgs Empty => new SecretV2CustomMetadataGetArgs();
    }
}
