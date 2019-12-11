// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vault.Identity
{
    /// <summary>
    /// Manages policies for an Identity Entity for Vault. The [Identity secrets engine](https://www.vaultproject.io/docs/secrets/identity/index.html) is the identity management solution for Vault.
    /// 
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-vault/blob/master/website/docs/r/identity_entity_policies.html.markdown.
    /// </summary>
    public partial class EntityPolicies : Pulumi.CustomResource
    {
        /// <summary>
        /// Entity ID to assign policies to.
        /// </summary>
        [Output("entityId")]
        public Output<string> EntityId { get; private set; } = null!;

        /// <summary>
        /// The name of the entity that are assigned the policies.
        /// </summary>
        [Output("entityName")]
        public Output<string> EntityName { get; private set; } = null!;

        /// <summary>
        /// Defaults to `true`.
        /// </summary>
        [Output("exclusive")]
        public Output<bool?> Exclusive { get; private set; } = null!;

        /// <summary>
        /// List of policies to assign to the entity
        /// </summary>
        [Output("policies")]
        public Output<ImmutableArray<string>> Policies { get; private set; } = null!;


        /// <summary>
        /// Create a EntityPolicies resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public EntityPolicies(string name, EntityPoliciesArgs args, CustomResourceOptions? options = null)
            : base("vault:identity/entityPolicies:EntityPolicies", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private EntityPolicies(string name, Input<string> id, EntityPoliciesState? state = null, CustomResourceOptions? options = null)
            : base("vault:identity/entityPolicies:EntityPolicies", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing EntityPolicies resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static EntityPolicies Get(string name, Input<string> id, EntityPoliciesState? state = null, CustomResourceOptions? options = null)
        {
            return new EntityPolicies(name, id, state, options);
        }
    }

    public sealed class EntityPoliciesArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Entity ID to assign policies to.
        /// </summary>
        [Input("entityId", required: true)]
        public Input<string> EntityId { get; set; } = null!;

        /// <summary>
        /// Defaults to `true`.
        /// </summary>
        [Input("exclusive")]
        public Input<bool>? Exclusive { get; set; }

        [Input("policies", required: true)]
        private InputList<string>? _policies;

        /// <summary>
        /// List of policies to assign to the entity
        /// </summary>
        public InputList<string> Policies
        {
            get => _policies ?? (_policies = new InputList<string>());
            set => _policies = value;
        }

        public EntityPoliciesArgs()
        {
        }
    }

    public sealed class EntityPoliciesState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Entity ID to assign policies to.
        /// </summary>
        [Input("entityId")]
        public Input<string>? EntityId { get; set; }

        /// <summary>
        /// The name of the entity that are assigned the policies.
        /// </summary>
        [Input("entityName")]
        public Input<string>? EntityName { get; set; }

        /// <summary>
        /// Defaults to `true`.
        /// </summary>
        [Input("exclusive")]
        public Input<bool>? Exclusive { get; set; }

        [Input("policies")]
        private InputList<string>? _policies;

        /// <summary>
        /// List of policies to assign to the entity
        /// </summary>
        public InputList<string> Policies
        {
            get => _policies ?? (_policies = new InputList<string>());
            set => _policies = value;
        }

        public EntityPoliciesState()
        {
        }
    }
}