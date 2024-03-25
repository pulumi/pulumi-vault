// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vault.Identity
{
    /// <summary>
    /// Manages OIDC Assignments in a Vault server. See the [Vault documentation](https://www.vaultproject.io/api-docs/secret/identity/oidc-provider#create-or-update-an-assignment)
    /// for more information.
    /// 
    /// ## Example Usage
    /// 
    /// &lt;!--Start PulumiCodeChooser --&gt;
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Vault = Pulumi.Vault;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var @internal = new Vault.Identity.Group("internal", new()
    ///     {
    ///         Type = "internal",
    ///         Policies = new[]
    ///         {
    ///             "dev",
    ///             "test",
    ///         },
    ///     });
    /// 
    ///     var test = new Vault.Identity.Entity("test", new()
    ///     {
    ///         Policies = new[]
    ///         {
    ///             "test",
    ///         },
    ///     });
    /// 
    ///     var @default = new Vault.Identity.OidcAssignment("default", new()
    ///     {
    ///         EntityIds = new[]
    ///         {
    ///             test.Id,
    ///         },
    ///         GroupIds = new[]
    ///         {
    ///             @internal.Id,
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ## Import
    /// 
    /// OIDC Assignments can be imported using the `name`, e.g.
    /// 
    /// ```sh
    /// $ pulumi import vault:identity/oidcAssignment:OidcAssignment default assignment
    /// ```
    /// </summary>
    [VaultResourceType("vault:identity/oidcAssignment:OidcAssignment")]
    public partial class OidcAssignment : global::Pulumi.CustomResource
    {
        /// <summary>
        /// A set of Vault entity IDs.
        /// </summary>
        [Output("entityIds")]
        public Output<ImmutableArray<string>> EntityIds { get; private set; } = null!;

        /// <summary>
        /// A set of Vault group IDs.
        /// </summary>
        [Output("groupIds")]
        public Output<ImmutableArray<string>> GroupIds { get; private set; } = null!;

        /// <summary>
        /// The name of the assignment.
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
        /// Create a OidcAssignment resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public OidcAssignment(string name, OidcAssignmentArgs? args = null, CustomResourceOptions? options = null)
            : base("vault:identity/oidcAssignment:OidcAssignment", name, args ?? new OidcAssignmentArgs(), MakeResourceOptions(options, ""))
        {
        }

        private OidcAssignment(string name, Input<string> id, OidcAssignmentState? state = null, CustomResourceOptions? options = null)
            : base("vault:identity/oidcAssignment:OidcAssignment", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing OidcAssignment resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static OidcAssignment Get(string name, Input<string> id, OidcAssignmentState? state = null, CustomResourceOptions? options = null)
        {
            return new OidcAssignment(name, id, state, options);
        }
    }

    public sealed class OidcAssignmentArgs : global::Pulumi.ResourceArgs
    {
        [Input("entityIds")]
        private InputList<string>? _entityIds;

        /// <summary>
        /// A set of Vault entity IDs.
        /// </summary>
        public InputList<string> EntityIds
        {
            get => _entityIds ?? (_entityIds = new InputList<string>());
            set => _entityIds = value;
        }

        [Input("groupIds")]
        private InputList<string>? _groupIds;

        /// <summary>
        /// A set of Vault group IDs.
        /// </summary>
        public InputList<string> GroupIds
        {
            get => _groupIds ?? (_groupIds = new InputList<string>());
            set => _groupIds = value;
        }

        /// <summary>
        /// The name of the assignment.
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

        public OidcAssignmentArgs()
        {
        }
        public static new OidcAssignmentArgs Empty => new OidcAssignmentArgs();
    }

    public sealed class OidcAssignmentState : global::Pulumi.ResourceArgs
    {
        [Input("entityIds")]
        private InputList<string>? _entityIds;

        /// <summary>
        /// A set of Vault entity IDs.
        /// </summary>
        public InputList<string> EntityIds
        {
            get => _entityIds ?? (_entityIds = new InputList<string>());
            set => _entityIds = value;
        }

        [Input("groupIds")]
        private InputList<string>? _groupIds;

        /// <summary>
        /// A set of Vault group IDs.
        /// </summary>
        public InputList<string> GroupIds
        {
            get => _groupIds ?? (_groupIds = new InputList<string>());
            set => _groupIds = value;
        }

        /// <summary>
        /// The name of the assignment.
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

        public OidcAssignmentState()
        {
        }
        public static new OidcAssignmentState Empty => new OidcAssignmentState();
    }
}
