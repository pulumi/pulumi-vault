// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vault
{
    public partial class Audit : Pulumi.CustomResource
    {
        /// <summary>
        /// Human-friendly description of the audit device.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// Configuration options to pass to the audit device itself.
        /// </summary>
        [Output("options")]
        public Output<ImmutableDictionary<string, string>> Options { get; private set; } = null!;

        /// <summary>
        /// The path to mount the audit device. This defaults to the type.
        /// </summary>
        [Output("path")]
        public Output<string> Path { get; private set; } = null!;

        /// <summary>
        /// Type of the audit device, such as 'file'.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a Audit resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Audit(string name, AuditArgs args, CustomResourceOptions? options = null)
            : base("vault:index/audit:Audit", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private Audit(string name, Input<string> id, AuditState? state = null, CustomResourceOptions? options = null)
            : base("vault:index/audit:Audit", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Audit resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Audit Get(string name, Input<string> id, AuditState? state = null, CustomResourceOptions? options = null)
        {
            return new Audit(name, id, state, options);
        }
    }

    public sealed class AuditArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Human-friendly description of the audit device.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("options", required: true)]
        private InputMap<string>? _options;

        /// <summary>
        /// Configuration options to pass to the audit device itself.
        /// </summary>
        public InputMap<string> Options
        {
            get => _options ?? (_options = new InputMap<string>());
            set => _options = value;
        }

        /// <summary>
        /// The path to mount the audit device. This defaults to the type.
        /// </summary>
        [Input("path")]
        public Input<string>? Path { get; set; }

        /// <summary>
        /// Type of the audit device, such as 'file'.
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        public AuditArgs()
        {
        }
    }

    public sealed class AuditState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Human-friendly description of the audit device.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("options")]
        private InputMap<string>? _options;

        /// <summary>
        /// Configuration options to pass to the audit device itself.
        /// </summary>
        public InputMap<string> Options
        {
            get => _options ?? (_options = new InputMap<string>());
            set => _options = value;
        }

        /// <summary>
        /// The path to mount the audit device. This defaults to the type.
        /// </summary>
        [Input("path")]
        public Input<string>? Path { get; set; }

        /// <summary>
        /// Type of the audit device, such as 'file'.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        public AuditState()
        {
        }
    }
}