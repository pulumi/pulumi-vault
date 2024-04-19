// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vault.Transform
{
    /// <summary>
    /// This resource supports the "/transform/role/{name}" Vault endpoint.
    /// 
    /// It creates or updates the role with the given name. If a role with the name does not exist, it will be created.
    /// If the role exists, it will be updated with the new attributes.
    /// 
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
    ///     var mountTransform = new Vault.Mount("mountTransform", new()
    ///     {
    ///         Path = "transform",
    ///         Type = "transform",
    ///     });
    /// 
    ///     var test = new Vault.Transform.Role("test", new()
    ///     {
    ///         Path = mountTransform.Path,
    ///         Transformations = new[]
    ///         {
    ///             "ccn-fpe",
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// </summary>
    [VaultResourceType("vault:transform/role:Role")]
    public partial class Role : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The name of the role.
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
        /// Path to where the back-end is mounted within Vault.
        /// </summary>
        [Output("path")]
        public Output<string> Path { get; private set; } = null!;

        /// <summary>
        /// A comma separated string or slice of transformations to use.
        /// </summary>
        [Output("transformations")]
        public Output<ImmutableArray<string>> Transformations { get; private set; } = null!;


        /// <summary>
        /// Create a Role resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Role(string name, RoleArgs args, CustomResourceOptions? options = null)
            : base("vault:transform/role:Role", name, args ?? new RoleArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Role(string name, Input<string> id, RoleState? state = null, CustomResourceOptions? options = null)
            : base("vault:transform/role:Role", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Role resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Role Get(string name, Input<string> id, RoleState? state = null, CustomResourceOptions? options = null)
        {
            return new Role(name, id, state, options);
        }
    }

    public sealed class RoleArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the role.
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
        /// Path to where the back-end is mounted within Vault.
        /// </summary>
        [Input("path", required: true)]
        public Input<string> Path { get; set; } = null!;

        [Input("transformations")]
        private InputList<string>? _transformations;

        /// <summary>
        /// A comma separated string or slice of transformations to use.
        /// </summary>
        public InputList<string> Transformations
        {
            get => _transformations ?? (_transformations = new InputList<string>());
            set => _transformations = value;
        }

        public RoleArgs()
        {
        }
        public static new RoleArgs Empty => new RoleArgs();
    }

    public sealed class RoleState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the role.
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
        /// Path to where the back-end is mounted within Vault.
        /// </summary>
        [Input("path")]
        public Input<string>? Path { get; set; }

        [Input("transformations")]
        private InputList<string>? _transformations;

        /// <summary>
        /// A comma separated string or slice of transformations to use.
        /// </summary>
        public InputList<string> Transformations
        {
            get => _transformations ?? (_transformations = new InputList<string>());
            set => _transformations = value;
        }

        public RoleState()
        {
        }
        public static new RoleState Empty => new RoleState();
    }
}
