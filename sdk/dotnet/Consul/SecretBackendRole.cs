// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vault.Consul
{
    /// <summary>
    /// Manages a Consul secrets role for a Consul secrets engine in Vault. Consul secret backends can then issue Consul tokens.
    /// 
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-vault/blob/master/website/docs/r/consul_secret_backend_role.html.markdown.
    /// </summary>
    public partial class SecretBackendRole : Pulumi.CustomResource
    {
        /// <summary>
        /// The name of the Consul secrets engine role to create.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The unique name of an existing Consul secrets backend mount. Must not begin or end with a `/`.
        /// </summary>
        [Output("path")]
        public Output<string> Path { get; private set; } = null!;

        /// <summary>
        /// The list of Consul ACL policies to associate with these roles.
        /// </summary>
        [Output("policies")]
        public Output<ImmutableArray<string>> Policies { get; private set; } = null!;


        /// <summary>
        /// Create a SecretBackendRole resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public SecretBackendRole(string name, SecretBackendRoleArgs args, CustomResourceOptions? options = null)
            : base("vault:consul/secretBackendRole:SecretBackendRole", name, args, MakeResourceOptions(options, ""))
        {
        }

        private SecretBackendRole(string name, Input<string> id, SecretBackendRoleState? state = null, CustomResourceOptions? options = null)
            : base("vault:consul/secretBackendRole:SecretBackendRole", name, state, MakeResourceOptions(options, id))
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
        /// The name of the Consul secrets engine role to create.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The unique name of an existing Consul secrets backend mount. Must not begin or end with a `/`.
        /// </summary>
        [Input("path", required: true)]
        public Input<string> Path { get; set; } = null!;

        [Input("policies", required: true)]
        private InputList<string>? _policies;

        /// <summary>
        /// The list of Consul ACL policies to associate with these roles.
        /// </summary>
        public InputList<string> Policies
        {
            get => _policies ?? (_policies = new InputList<string>());
            set => _policies = value;
        }

        public SecretBackendRoleArgs()
        {
        }
    }

    public sealed class SecretBackendRoleState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the Consul secrets engine role to create.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The unique name of an existing Consul secrets backend mount. Must not begin or end with a `/`.
        /// </summary>
        [Input("path")]
        public Input<string>? Path { get; set; }

        [Input("policies")]
        private InputList<string>? _policies;

        /// <summary>
        /// The list of Consul ACL policies to associate with these roles.
        /// </summary>
        public InputList<string> Policies
        {
            get => _policies ?? (_policies = new InputList<string>());
            set => _policies = value;
        }

        public SecretBackendRoleState()
        {
        }
    }
}
