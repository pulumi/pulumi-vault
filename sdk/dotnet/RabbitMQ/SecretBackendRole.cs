// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vault.RabbitMQ
{
    /// <summary>
    /// ## Import
    /// 
    /// RabbitMQ secret backend roles can be imported using the `path`, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import vault:rabbitMq/secretBackendRole:SecretBackendRole role rabbitmq/roles/deploy
    /// ```
    /// </summary>
    [VaultResourceType("vault:rabbitMq/secretBackendRole:SecretBackendRole")]
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
        public Output<ImmutableArray<Outputs.SecretBackendRoleVhost>> Vhosts { get; private set; } = null!;


        /// <summary>
        /// Create a SecretBackendRole resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public SecretBackendRole(string name, SecretBackendRoleArgs args, CustomResourceOptions? options = null)
            : base("vault:rabbitMq/secretBackendRole:SecretBackendRole", name, args ?? new SecretBackendRoleArgs(), MakeResourceOptions(options, ""))
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
        private InputList<Inputs.SecretBackendRoleVhostArgs>? _vhosts;

        /// <summary>
        /// Specifies a map of virtual hosts to permissions.
        /// </summary>
        public InputList<Inputs.SecretBackendRoleVhostArgs> Vhosts
        {
            get => _vhosts ?? (_vhosts = new InputList<Inputs.SecretBackendRoleVhostArgs>());
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
        private InputList<Inputs.SecretBackendRoleVhostGetArgs>? _vhosts;

        /// <summary>
        /// Specifies a map of virtual hosts to permissions.
        /// </summary>
        public InputList<Inputs.SecretBackendRoleVhostGetArgs> Vhosts
        {
            get => _vhosts ?? (_vhosts = new InputList<Inputs.SecretBackendRoleVhostGetArgs>());
            set => _vhosts = value;
        }

        public SecretBackendRoleState()
        {
        }
    }
}
