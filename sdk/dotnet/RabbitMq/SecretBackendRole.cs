// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vault.RabbitMq
{
    /// <summary>
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-vault/blob/master/website/docs/r/rabbitmq_secret_backend_role.html.markdown.
    /// </summary>
    public partial class SecretBackendRole : Pulumi.CustomResource
    {
        /// <summary>
        /// The path the RabbitMQ secret backend is mounted at,
        /// with no leading or trailing `/`s.
        /// </summary>
        [Output("backend")]
        public Output<string> Backend { get; private set; } = null!;

        /// <summary>
        /// The name to identify this role within the backend.
        /// Must be unique within the backend.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Specifies a comma-separated RabbitMQ management tags.
        /// </summary>
        [Output("tags")]
        public Output<string?> Tags { get; private set; } = null!;

        /// <summary>
        /// Specifies a map of virtual hosts to permissions.
        /// </summary>
        [Output("vhosts")]
        public Output<ImmutableArray<Outputs.SecretBackendRoleVhosts>> Vhosts { get; private set; } = null!;


        /// <summary>
        /// Create a SecretBackendRole resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public SecretBackendRole(string name, SecretBackendRoleArgs args, CustomResourceOptions? options = null)
            : base("vault:rabbitMq/secretBackendRole:SecretBackendRole", name, args, MakeResourceOptions(options, ""))
        {
        }

        private SecretBackendRole(string name, Input<string> id, SecretBackendRoleState? state = null, CustomResourceOptions? options = null)
            : base("vault:rabbitMq/secretBackendRole:SecretBackendRole", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing SecretBackendRole resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static SecretBackendRole Get(string name, Input<string> id, SecretBackendRoleState? state = null, CustomResourceOptions? options = null)
        {
            return new SecretBackendRole(name, id, state, options);
        }
    }

    public sealed class SecretBackendRoleArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The path the RabbitMQ secret backend is mounted at,
        /// with no leading or trailing `/`s.
        /// </summary>
        [Input("backend", required: true)]
        public Input<string> Backend { get; set; } = null!;

        /// <summary>
        /// The name to identify this role within the backend.
        /// Must be unique within the backend.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Specifies a comma-separated RabbitMQ management tags.
        /// </summary>
        [Input("tags")]
        public Input<string>? Tags { get; set; }

        [Input("vhosts")]
        private InputList<Inputs.SecretBackendRoleVhostsArgs>? _vhosts;

        /// <summary>
        /// Specifies a map of virtual hosts to permissions.
        /// </summary>
        public InputList<Inputs.SecretBackendRoleVhostsArgs> Vhosts
        {
            get => _vhosts ?? (_vhosts = new InputList<Inputs.SecretBackendRoleVhostsArgs>());
            set => _vhosts = value;
        }

        public SecretBackendRoleArgs()
        {
        }
    }

    public sealed class SecretBackendRoleState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The path the RabbitMQ secret backend is mounted at,
        /// with no leading or trailing `/`s.
        /// </summary>
        [Input("backend")]
        public Input<string>? Backend { get; set; }

        /// <summary>
        /// The name to identify this role within the backend.
        /// Must be unique within the backend.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Specifies a comma-separated RabbitMQ management tags.
        /// </summary>
        [Input("tags")]
        public Input<string>? Tags { get; set; }

        [Input("vhosts")]
        private InputList<Inputs.SecretBackendRoleVhostsGetArgs>? _vhosts;

        /// <summary>
        /// Specifies a map of virtual hosts to permissions.
        /// </summary>
        public InputList<Inputs.SecretBackendRoleVhostsGetArgs> Vhosts
        {
            get => _vhosts ?? (_vhosts = new InputList<Inputs.SecretBackendRoleVhostsGetArgs>());
            set => _vhosts = value;
        }

        public SecretBackendRoleState()
        {
        }
    }

    namespace Inputs
    {

    public sealed class SecretBackendRoleVhostsArgs : Pulumi.ResourceArgs
    {
        [Input("configure", required: true)]
        public Input<string> Configure { get; set; } = null!;

        [Input("host", required: true)]
        public Input<string> Host { get; set; } = null!;

        [Input("read", required: true)]
        public Input<string> Read { get; set; } = null!;

        [Input("write", required: true)]
        public Input<string> Write { get; set; } = null!;

        public SecretBackendRoleVhostsArgs()
        {
        }
    }

    public sealed class SecretBackendRoleVhostsGetArgs : Pulumi.ResourceArgs
    {
        [Input("configure", required: true)]
        public Input<string> Configure { get; set; } = null!;

        [Input("host", required: true)]
        public Input<string> Host { get; set; } = null!;

        [Input("read", required: true)]
        public Input<string> Read { get; set; } = null!;

        [Input("write", required: true)]
        public Input<string> Write { get; set; } = null!;

        public SecretBackendRoleVhostsGetArgs()
        {
        }
    }
    }

    namespace Outputs
    {

    [OutputType]
    public sealed class SecretBackendRoleVhosts
    {
        public readonly string Configure;
        public readonly string Host;
        public readonly string Read;
        public readonly string Write;

        [OutputConstructor]
        private SecretBackendRoleVhosts(
            string configure,
            string host,
            string read,
            string write)
        {
            Configure = configure;
            Host = host;
            Read = read;
            Write = write;
        }
    }
    }
}
