// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vault.Gcp.Inputs
{

    public sealed class SecretStaticAccountBindingGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Resource or resource path for which IAM policy information will be bound. The resource path may be specified in a few different [formats](https://www.vaultproject.io/docs/secrets/gcp/index.html#bindings).
        /// </summary>
        [Input("resource", required: true)]
        public Input<string> Resource { get; set; } = null!;

        [Input("roles", required: true)]
        private InputList<string>? _roles;

        /// <summary>
        /// List of [GCP IAM roles](https://cloud.google.com/iam/docs/understanding-roles) for the resource.
        /// </summary>
        public InputList<string> Roles
        {
            get => _roles ?? (_roles = new InputList<string>());
            set => _roles = value;
        }

        public SecretStaticAccountBindingGetArgs()
        {
        }
        public static new SecretStaticAccountBindingGetArgs Empty => new SecretStaticAccountBindingGetArgs();
    }
}
