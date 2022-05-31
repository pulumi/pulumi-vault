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
    /// Manages member entities for an Identity Group for Vault. The [Identity secrets engine](https://www.vaultproject.io/docs/secrets/identity/index.html) is the identity management solution for Vault.
    /// 
    /// ## Example Usage
    /// ### Exclusive Member Entities
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Vault = Pulumi.Vault;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var @internal = new Vault.Identity.Group("internal", new Vault.Identity.GroupArgs
    ///         {
    ///             Type = "internal",
    ///             ExternalMemberEntityIds = true,
    ///             Metadata = 
    ///             {
    ///                 { "version", "2" },
    ///             },
    ///         });
    ///         var user = new Vault.Identity.Entity("user", new Vault.Identity.EntityArgs
    ///         {
    ///         });
    ///         var members = new Vault.Identity.GroupMemberEntityIds("members", new Vault.Identity.GroupMemberEntityIdsArgs
    ///         {
    ///             Exclusive = true,
    ///             MemberEntityIds = 
    ///             {
    ///                 user.Id,
    ///             },
    ///             GroupId = @internal.Id,
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// ### Non-exclusive Member Entities
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Vault = Pulumi.Vault;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var @internal = new Vault.Identity.Group("internal", new Vault.Identity.GroupArgs
    ///         {
    ///             Type = "internal",
    ///             ExternalMemberEntityIds = true,
    ///             Metadata = 
    ///             {
    ///                 { "version", "2" },
    ///             },
    ///         });
    ///         var testUser = new Vault.Identity.Entity("testUser", new Vault.Identity.EntityArgs
    ///         {
    ///         });
    ///         var secondTestUser = new Vault.Identity.Entity("secondTestUser", new Vault.Identity.EntityArgs
    ///         {
    ///         });
    ///         var devUser = new Vault.Identity.Entity("devUser", new Vault.Identity.EntityArgs
    ///         {
    ///         });
    ///         var test = new Vault.Identity.GroupMemberEntityIds("test", new Vault.Identity.GroupMemberEntityIdsArgs
    ///         {
    ///             MemberEntityIds = 
    ///             {
    ///                 testUser.Id,
    ///                 secondTestUser.Id,
    ///             },
    ///             Exclusive = false,
    ///             GroupId = @internal.Id,
    ///         });
    ///         var others = new Vault.Identity.GroupMemberEntityIds("others", new Vault.Identity.GroupMemberEntityIdsArgs
    ///         {
    ///             MemberEntityIds = 
    ///             {
    ///                 devUser.Id,
    ///             },
    ///             Exclusive = false,
    ///             GroupId = @internal.Id,
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// </summary>
    [VaultResourceType("vault:identity/groupMemberEntityIds:GroupMemberEntityIds")]
    public partial class GroupMemberEntityIds : Pulumi.CustomResource
    {
        /// <summary>
        /// Defaults to `true`.
        /// </summary>
        [Output("exclusive")]
        public Output<bool?> Exclusive { get; private set; } = null!;

        /// <summary>
        /// Group ID to assign member entities to.
        /// </summary>
        [Output("groupId")]
        public Output<string> GroupId { get; private set; } = null!;

        /// <summary>
        /// The name of the group that are assigned the member entities.  
        /// *Deprecated: The value for group_name may not always be accurate*
        /// *use* `data.vault_identity_group.*.group_name`, *or* `vault_identity_group.*.group_name` *instead.*
        /// </summary>
        [Output("groupName")]
        public Output<string> GroupName { get; private set; } = null!;

        /// <summary>
        /// List of member entities that belong to the group
        /// </summary>
        [Output("memberEntityIds")]
        public Output<ImmutableArray<string>> MemberEntityIds { get; private set; } = null!;


        /// <summary>
        /// Create a GroupMemberEntityIds resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public GroupMemberEntityIds(string name, GroupMemberEntityIdsArgs args, CustomResourceOptions? options = null)
            : base("vault:identity/groupMemberEntityIds:GroupMemberEntityIds", name, args ?? new GroupMemberEntityIdsArgs(), MakeResourceOptions(options, ""))
        {
        }

        private GroupMemberEntityIds(string name, Input<string> id, GroupMemberEntityIdsState? state = null, CustomResourceOptions? options = null)
            : base("vault:identity/groupMemberEntityIds:GroupMemberEntityIds", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing GroupMemberEntityIds resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static GroupMemberEntityIds Get(string name, Input<string> id, GroupMemberEntityIdsState? state = null, CustomResourceOptions? options = null)
        {
            return new GroupMemberEntityIds(name, id, state, options);
        }
    }

    public sealed class GroupMemberEntityIdsArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Defaults to `true`.
        /// </summary>
        [Input("exclusive")]
        public Input<bool>? Exclusive { get; set; }

        /// <summary>
        /// Group ID to assign member entities to.
        /// </summary>
        [Input("groupId", required: true)]
        public Input<string> GroupId { get; set; } = null!;

        [Input("memberEntityIds")]
        private InputList<string>? _memberEntityIds;

        /// <summary>
        /// List of member entities that belong to the group
        /// </summary>
        public InputList<string> MemberEntityIds
        {
            get => _memberEntityIds ?? (_memberEntityIds = new InputList<string>());
            set => _memberEntityIds = value;
        }

        public GroupMemberEntityIdsArgs()
        {
        }
    }

    public sealed class GroupMemberEntityIdsState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Defaults to `true`.
        /// </summary>
        [Input("exclusive")]
        public Input<bool>? Exclusive { get; set; }

        /// <summary>
        /// Group ID to assign member entities to.
        /// </summary>
        [Input("groupId")]
        public Input<string>? GroupId { get; set; }

        /// <summary>
        /// The name of the group that are assigned the member entities.  
        /// *Deprecated: The value for group_name may not always be accurate*
        /// *use* `data.vault_identity_group.*.group_name`, *or* `vault_identity_group.*.group_name` *instead.*
        /// </summary>
        [Input("groupName")]
        public Input<string>? GroupName { get; set; }

        [Input("memberEntityIds")]
        private InputList<string>? _memberEntityIds;

        /// <summary>
        /// List of member entities that belong to the group
        /// </summary>
        public InputList<string> MemberEntityIds
        {
            get => _memberEntityIds ?? (_memberEntityIds = new InputList<string>());
            set => _memberEntityIds = value;
        }

        public GroupMemberEntityIdsState()
        {
        }
    }
}
