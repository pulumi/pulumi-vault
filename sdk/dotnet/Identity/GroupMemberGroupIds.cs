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
    /// Manages member groups for an Identity Group for Vault. The
    /// [Identity secrets engine](https://www.vaultproject.io/docs/secrets/identity/index.html)
    /// is the identity management solution for Vault.
    /// 
    /// ## Example Usage
    /// 
    /// ### Exclusive Member Groups
    /// 
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
    ///         ExternalMemberGroupIds = true,
    ///         Metadata = 
    ///         {
    ///             { "version", "2" },
    ///         },
    ///     });
    /// 
    ///     var users = new Vault.Identity.Group("users", new()
    ///     {
    ///         Metadata = 
    ///         {
    ///             { "version", "2" },
    ///         },
    ///     });
    /// 
    ///     var members = new Vault.Identity.GroupMemberGroupIds("members", new()
    ///     {
    ///         Exclusive = true,
    ///         MemberGroupIds = new[]
    ///         {
    ///             users.Id,
    ///         },
    ///         GroupId = @internal.Id,
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ### Non-Exclusive Member Groups
    /// 
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
    ///         ExternalMemberGroupIds = true,
    ///         Metadata = 
    ///         {
    ///             { "version", "2" },
    ///         },
    ///     });
    /// 
    ///     var users = new Vault.Identity.Group("users", new()
    ///     {
    ///         Metadata = 
    ///         {
    ///             { "version", "2" },
    ///         },
    ///     });
    /// 
    ///     var members = new Vault.Identity.GroupMemberGroupIds("members", new()
    ///     {
    ///         Exclusive = false,
    ///         MemberGroupIds = new[]
    ///         {
    ///             users.Id,
    ///         },
    ///         GroupId = @internal.Id,
    ///     });
    /// 
    /// });
    /// ```
    /// </summary>
    [VaultResourceType("vault:identity/groupMemberGroupIds:GroupMemberGroupIds")]
    public partial class GroupMemberGroupIds : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Defaults to `true`.
        /// 
        /// If `true`, this resource will take exclusive control of the member groups that belong to the group and will set
        /// it equal to what is specified in the resource.
        /// 
        /// If set to `false`, this resource will simply ensure that the member groups specified in the resource are present
        /// in the group. When destroying the resource, the resource will ensure that the member groups specified in the resource
        /// are removed.
        /// </summary>
        [Output("exclusive")]
        public Output<bool?> Exclusive { get; private set; } = null!;

        /// <summary>
        /// Group ID to assign member entities to.
        /// </summary>
        [Output("groupId")]
        public Output<string> GroupId { get; private set; } = null!;

        /// <summary>
        /// List of member groups that belong to the group
        /// </summary>
        [Output("memberGroupIds")]
        public Output<ImmutableArray<string>> MemberGroupIds { get; private set; } = null!;

        /// <summary>
        /// The namespace to provision the resource in.
        /// The value should not contain leading or trailing forward slashes.
        /// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
        /// *Available only for Vault Enterprise*.
        /// </summary>
        [Output("namespace")]
        public Output<string?> Namespace { get; private set; } = null!;


        /// <summary>
        /// Create a GroupMemberGroupIds resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public GroupMemberGroupIds(string name, GroupMemberGroupIdsArgs args, CustomResourceOptions? options = null)
            : base("vault:identity/groupMemberGroupIds:GroupMemberGroupIds", name, args ?? new GroupMemberGroupIdsArgs(), MakeResourceOptions(options, ""))
        {
        }

        private GroupMemberGroupIds(string name, Input<string> id, GroupMemberGroupIdsState? state = null, CustomResourceOptions? options = null)
            : base("vault:identity/groupMemberGroupIds:GroupMemberGroupIds", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing GroupMemberGroupIds resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static GroupMemberGroupIds Get(string name, Input<string> id, GroupMemberGroupIdsState? state = null, CustomResourceOptions? options = null)
        {
            return new GroupMemberGroupIds(name, id, state, options);
        }
    }

    public sealed class GroupMemberGroupIdsArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Defaults to `true`.
        /// 
        /// If `true`, this resource will take exclusive control of the member groups that belong to the group and will set
        /// it equal to what is specified in the resource.
        /// 
        /// If set to `false`, this resource will simply ensure that the member groups specified in the resource are present
        /// in the group. When destroying the resource, the resource will ensure that the member groups specified in the resource
        /// are removed.
        /// </summary>
        [Input("exclusive")]
        public Input<bool>? Exclusive { get; set; }

        /// <summary>
        /// Group ID to assign member entities to.
        /// </summary>
        [Input("groupId", required: true)]
        public Input<string> GroupId { get; set; } = null!;

        [Input("memberGroupIds")]
        private InputList<string>? _memberGroupIds;

        /// <summary>
        /// List of member groups that belong to the group
        /// </summary>
        public InputList<string> MemberGroupIds
        {
            get => _memberGroupIds ?? (_memberGroupIds = new InputList<string>());
            set => _memberGroupIds = value;
        }

        /// <summary>
        /// The namespace to provision the resource in.
        /// The value should not contain leading or trailing forward slashes.
        /// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
        /// *Available only for Vault Enterprise*.
        /// </summary>
        [Input("namespace")]
        public Input<string>? Namespace { get; set; }

        public GroupMemberGroupIdsArgs()
        {
        }
        public static new GroupMemberGroupIdsArgs Empty => new GroupMemberGroupIdsArgs();
    }

    public sealed class GroupMemberGroupIdsState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Defaults to `true`.
        /// 
        /// If `true`, this resource will take exclusive control of the member groups that belong to the group and will set
        /// it equal to what is specified in the resource.
        /// 
        /// If set to `false`, this resource will simply ensure that the member groups specified in the resource are present
        /// in the group. When destroying the resource, the resource will ensure that the member groups specified in the resource
        /// are removed.
        /// </summary>
        [Input("exclusive")]
        public Input<bool>? Exclusive { get; set; }

        /// <summary>
        /// Group ID to assign member entities to.
        /// </summary>
        [Input("groupId")]
        public Input<string>? GroupId { get; set; }

        [Input("memberGroupIds")]
        private InputList<string>? _memberGroupIds;

        /// <summary>
        /// List of member groups that belong to the group
        /// </summary>
        public InputList<string> MemberGroupIds
        {
            get => _memberGroupIds ?? (_memberGroupIds = new InputList<string>());
            set => _memberGroupIds = value;
        }

        /// <summary>
        /// The namespace to provision the resource in.
        /// The value should not contain leading or trailing forward slashes.
        /// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
        /// *Available only for Vault Enterprise*.
        /// </summary>
        [Input("namespace")]
        public Input<string>? Namespace { get; set; }

        public GroupMemberGroupIdsState()
        {
        }
        public static new GroupMemberGroupIdsState Empty => new GroupMemberGroupIdsState();
    }
}
