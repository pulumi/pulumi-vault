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
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Vault = Pulumi.Vault;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var rabbitmq = new Vault.RabbitMQ.SecretBackend("rabbitmq", new()
    ///     {
    ///         ConnectionUri = "https://.....",
    ///         Username = "user",
    ///         Password = "password",
    ///     });
    /// 
    ///     var role = new Vault.RabbitMQ.SecretBackendRole("role", new()
    ///     {
    ///         Backend = rabbitmq.Path,
    ///         Name = "deploy",
    ///         Tags = "tag1,tag2",
    ///         Vhosts = new[]
    ///         {
    ///             new Vault.RabbitMQ.Inputs.SecretBackendRoleVhostArgs
    ///             {
    ///                 Host = "/",
    ///                 Configure = "",
    ///                 Read = ".*",
    ///                 Write = "",
    ///             },
    ///         },
    ///         VhostTopics = new[]
    ///         {
    ///             new Vault.RabbitMQ.Inputs.SecretBackendRoleVhostTopicArgs
    ///             {
    ///                 Vhosts = new[]
    ///                 {
    ///                     new Vault.RabbitMQ.Inputs.SecretBackendRoleVhostTopicVhostArgs
    ///                     {
    ///                         Topic = "amq.topic",
    ///                         Read = ".*",
    ///                         Write = "",
    ///                     },
    ///                 },
    ///                 Host = "/",
    ///             },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// RabbitMQ secret backend roles can be imported using the `path`, e.g.
    /// 
    /// ```sh
    /// $ pulumi import vault:rabbitMq/secretBackendRole:SecretBackendRole role rabbitmq/roles/deploy
    /// ```
    /// </summary>
    [VaultResourceType("vault:rabbitMq/secretBackendRole:SecretBackendRole")]
    public partial class SecretBackendRole : global::Pulumi.CustomResource
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
        /// The namespace to provision the resource in.
        /// The value should not contain leading or trailing forward slashes.
        /// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
        /// *Available only for Vault Enterprise*.
        /// </summary>
        [Output("namespace")]
        public Output<string?> Namespace { get; private set; } = null!;

        /// <summary>
        /// Specifies a comma-separated RabbitMQ management tags.
        /// </summary>
        [Output("tags")]
        public Output<string?> Tags { get; private set; } = null!;

        /// <summary>
        /// Specifies a map of virtual hosts and exchanges to topic permissions. This option requires RabbitMQ 3.7.0 or later.
        /// </summary>
        [Output("vhostTopics")]
        public Output<ImmutableArray<Outputs.SecretBackendRoleVhostTopic>> VhostTopics { get; private set; } = null!;

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

    public sealed class SecretBackendRoleArgs : global::Pulumi.ResourceArgs
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
        /// The namespace to provision the resource in.
        /// The value should not contain leading or trailing forward slashes.
        /// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
        /// *Available only for Vault Enterprise*.
        /// </summary>
        [Input("namespace")]
        public Input<string>? Namespace { get; set; }

        /// <summary>
        /// Specifies a comma-separated RabbitMQ management tags.
        /// </summary>
        [Input("tags")]
        public Input<string>? Tags { get; set; }

        [Input("vhostTopics")]
        private InputList<Inputs.SecretBackendRoleVhostTopicArgs>? _vhostTopics;

        /// <summary>
        /// Specifies a map of virtual hosts and exchanges to topic permissions. This option requires RabbitMQ 3.7.0 or later.
        /// </summary>
        public InputList<Inputs.SecretBackendRoleVhostTopicArgs> VhostTopics
        {
            get => _vhostTopics ?? (_vhostTopics = new InputList<Inputs.SecretBackendRoleVhostTopicArgs>());
            set => _vhostTopics = value;
        }

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
        public static new SecretBackendRoleArgs Empty => new SecretBackendRoleArgs();
    }

    public sealed class SecretBackendRoleState : global::Pulumi.ResourceArgs
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
        /// The namespace to provision the resource in.
        /// The value should not contain leading or trailing forward slashes.
        /// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
        /// *Available only for Vault Enterprise*.
        /// </summary>
        [Input("namespace")]
        public Input<string>? Namespace { get; set; }

        /// <summary>
        /// Specifies a comma-separated RabbitMQ management tags.
        /// </summary>
        [Input("tags")]
        public Input<string>? Tags { get; set; }

        [Input("vhostTopics")]
        private InputList<Inputs.SecretBackendRoleVhostTopicGetArgs>? _vhostTopics;

        /// <summary>
        /// Specifies a map of virtual hosts and exchanges to topic permissions. This option requires RabbitMQ 3.7.0 or later.
        /// </summary>
        public InputList<Inputs.SecretBackendRoleVhostTopicGetArgs> VhostTopics
        {
            get => _vhostTopics ?? (_vhostTopics = new InputList<Inputs.SecretBackendRoleVhostTopicGetArgs>());
            set => _vhostTopics = value;
        }

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
        public static new SecretBackendRoleState Empty => new SecretBackendRoleState();
    }
}
