// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vault.Generic
{
    /// <summary>
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-vault/blob/master/website/docs/r/generic_secret.html.markdown.
    /// </summary>
    public partial class Secret : Pulumi.CustomResource
    {
        /// <summary>
        /// A mapping whose keys are the top-level data keys returned from
        /// Vault and whose values are the corresponding values. This map can only
        /// represent string data, so any non-string values returned from Vault are
        /// serialized as JSON.
        /// </summary>
        [Output("data")]
        public Output<ImmutableDictionary<string, object>> Data { get; private set; } = null!;

        /// <summary>
        /// String containing a JSON-encoded object that will be
        /// written as the secret data at the given path.
        /// </summary>
        [Output("dataJson")]
        public Output<string> DataJson { get; private set; } = null!;

        /// <summary>
        /// True/false. Set this to true if your vault
        /// authentication is not able to read the data. Setting this to `true` will
        /// break drift detection. Defaults to false.
        /// </summary>
        [Output("disableRead")]
        public Output<bool?> DisableRead { get; private set; } = null!;

        /// <summary>
        /// The full logical path at which to write the given data.
        /// To write data into the "generic" secret backend mounted in Vault by default,
        /// this should be prefixed with `secret/`. Writing to other backends with this
        /// resource is possible; consult each backend's documentation to see which
        /// endpoints support the `PUT` and `DELETE` methods.
        /// </summary>
        [Output("path")]
        public Output<string> Path { get; private set; } = null!;


        /// <summary>
        /// Create a Secret resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Secret(string name, SecretArgs args, CustomResourceOptions? options = null)
            : base("vault:generic/secret:Secret", name, args, MakeResourceOptions(options, ""))
        {
        }

        private Secret(string name, Input<string> id, SecretState? state = null, CustomResourceOptions? options = null)
            : base("vault:generic/secret:Secret", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Secret resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Secret Get(string name, Input<string> id, SecretState? state = null, CustomResourceOptions? options = null)
        {
            return new Secret(name, id, state, options);
        }
    }

    public sealed class SecretArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// String containing a JSON-encoded object that will be
        /// written as the secret data at the given path.
        /// </summary>
        [Input("dataJson", required: true)]
        public Input<string> DataJson { get; set; } = null!;

        /// <summary>
        /// True/false. Set this to true if your vault
        /// authentication is not able to read the data. Setting this to `true` will
        /// break drift detection. Defaults to false.
        /// </summary>
        [Input("disableRead")]
        public Input<bool>? DisableRead { get; set; }

        /// <summary>
        /// The full logical path at which to write the given data.
        /// To write data into the "generic" secret backend mounted in Vault by default,
        /// this should be prefixed with `secret/`. Writing to other backends with this
        /// resource is possible; consult each backend's documentation to see which
        /// endpoints support the `PUT` and `DELETE` methods.
        /// </summary>
        [Input("path", required: true)]
        public Input<string> Path { get; set; } = null!;

        public SecretArgs()
        {
        }
    }

    public sealed class SecretState : Pulumi.ResourceArgs
    {
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
        /// String containing a JSON-encoded object that will be
        /// written as the secret data at the given path.
        /// </summary>
        [Input("dataJson")]
        public Input<string>? DataJson { get; set; }

        /// <summary>
        /// True/false. Set this to true if your vault
        /// authentication is not able to read the data. Setting this to `true` will
        /// break drift detection. Defaults to false.
        /// </summary>
        [Input("disableRead")]
        public Input<bool>? DisableRead { get; set; }

        /// <summary>
        /// The full logical path at which to write the given data.
        /// To write data into the "generic" secret backend mounted in Vault by default,
        /// this should be prefixed with `secret/`. Writing to other backends with this
        /// resource is possible; consult each backend's documentation to see which
        /// endpoints support the `PUT` and `DELETE` methods.
        /// </summary>
        [Input("path")]
        public Input<string>? Path { get; set; }

        public SecretState()
        {
        }
    }
}
