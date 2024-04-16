// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vault.MongoDBAtlas
{
    /// <summary>
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
    ///     var mongo = new Vault.Mount("mongo", new()
    ///     {
    ///         Path = "%s",
    ///         Type = "mongodbatlas",
    ///         Description = "MongoDB Atlas secret engine mount",
    ///     });
    /// 
    ///     var config = new Vault.MongoDBAtlas.SecretBackend("config", new()
    ///     {
    ///         Mount = mongo.Path,
    ///         PrivateKey = "privateKey",
    ///         PublicKey = "publicKey",
    ///     });
    /// 
    ///     var role = new Vault.MongoDBAtlas.SecretRole("role", new()
    ///     {
    ///         Mount = mongo.Path,
    ///         Name = "tf-test-role",
    ///         OrganizationId = "7cf5a45a9ccf6400e60981b7",
    ///         ProjectId = "5cf5a45a9ccf6400e60981b6",
    ///         Roles = new[]
    ///         {
    ///             "ORG_READ_ONLY",
    ///         },
    ///         IpAddresses = "192.168.1.5, 192.168.1.6",
    ///         CidrBlocks = "192.168.1.3/35",
    ///         ProjectRoles = new[]
    ///         {
    ///             "GROUP_READ_ONLY",
    ///         },
    ///         Ttl = "60",
    ///         MaxTtl = "120",
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ## Import
    /// 
    /// The MongoDB Atlas secret role can be imported using the full path to the role
    /// of the form: `&lt;mount_path&gt;/roles/&lt;role_name&gt;` e.g.
    /// 
    /// ```sh
    /// $ pulumi import vault:mongodbatlas/secretRole:SecretRole example mongodbatlas/roles/example-role
    /// ```
    /// </summary>
    [VaultResourceType("vault:mongodbatlas/secretRole:SecretRole")]
    public partial class SecretRole : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Whitelist entry in CIDR notation to be added for the API key.
        /// </summary>
        [Output("cidrBlocks")]
        public Output<ImmutableArray<string>> CidrBlocks { get; private set; } = null!;

        /// <summary>
        /// IP address to be added to the whitelist for the API key.
        /// </summary>
        [Output("ipAddresses")]
        public Output<ImmutableArray<string>> IpAddresses { get; private set; } = null!;

        /// <summary>
        /// The maximum allowed lifetime of credentials issued using this role.
        /// </summary>
        [Output("maxTtl")]
        public Output<string?> MaxTtl { get; private set; } = null!;

        /// <summary>
        /// Path where the MongoDB Atlas Secrets Engine is mounted.
        /// </summary>
        [Output("mount")]
        public Output<string> Mount { get; private set; } = null!;

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
        /// Unique identifier for the organization to which the target API Key belongs. 
        /// Required if `project_id` is not set.
        /// </summary>
        [Output("organizationId")]
        public Output<string?> OrganizationId { get; private set; } = null!;

        /// <summary>
        /// Unique identifier for the project to which the target API Key belongs.
        /// Required if `organization_id` is not set.
        /// </summary>
        [Output("projectId")]
        public Output<string?> ProjectId { get; private set; } = null!;

        /// <summary>
        /// Roles assigned when an org API key is assigned to a project API key. Possible values are `GROUP_CLUSTER_MANAGER`, `GROUP_DATA_ACCESS_ADMIN`, `GROUP_DATA_ACCESS_READ_ONLY`, `GROUP_DATA_ACCESS_READ_WRITE`, `GROUP_OWNER` and `GROUP_READ_ONLY`.
        /// </summary>
        [Output("projectRoles")]
        public Output<ImmutableArray<string>> ProjectRoles { get; private set; } = null!;

        /// <summary>
        /// List of roles that the API Key needs to have. Possible values are `ORG_OWNER`, `ORG_MEMBER`, `ORG_GROUP_CREATOR`, `ORG_BILLING_ADMIN` and `ORG_READ_ONLY`.
        /// </summary>
        [Output("roles")]
        public Output<ImmutableArray<string>> Roles { get; private set; } = null!;

        /// <summary>
        /// Duration in seconds after which the issued credential should expire.
        /// </summary>
        [Output("ttl")]
        public Output<string?> Ttl { get; private set; } = null!;


        /// <summary>
        /// Create a SecretRole resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public SecretRole(string name, SecretRoleArgs args, CustomResourceOptions? options = null)
            : base("vault:mongodbatlas/secretRole:SecretRole", name, args ?? new SecretRoleArgs(), MakeResourceOptions(options, ""))
        {
        }

        private SecretRole(string name, Input<string> id, SecretRoleState? state = null, CustomResourceOptions? options = null)
            : base("vault:mongodbatlas/secretRole:SecretRole", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing SecretRole resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static SecretRole Get(string name, Input<string> id, SecretRoleState? state = null, CustomResourceOptions? options = null)
        {
            return new SecretRole(name, id, state, options);
        }
    }

    public sealed class SecretRoleArgs : global::Pulumi.ResourceArgs
    {
        [Input("cidrBlocks")]
        private InputList<string>? _cidrBlocks;

        /// <summary>
        /// Whitelist entry in CIDR notation to be added for the API key.
        /// </summary>
        public InputList<string> CidrBlocks
        {
            get => _cidrBlocks ?? (_cidrBlocks = new InputList<string>());
            set => _cidrBlocks = value;
        }

        [Input("ipAddresses")]
        private InputList<string>? _ipAddresses;

        /// <summary>
        /// IP address to be added to the whitelist for the API key.
        /// </summary>
        public InputList<string> IpAddresses
        {
            get => _ipAddresses ?? (_ipAddresses = new InputList<string>());
            set => _ipAddresses = value;
        }

        /// <summary>
        /// The maximum allowed lifetime of credentials issued using this role.
        /// </summary>
        [Input("maxTtl")]
        public Input<string>? MaxTtl { get; set; }

        /// <summary>
        /// Path where the MongoDB Atlas Secrets Engine is mounted.
        /// </summary>
        [Input("mount", required: true)]
        public Input<string> Mount { get; set; } = null!;

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
        /// Unique identifier for the organization to which the target API Key belongs. 
        /// Required if `project_id` is not set.
        /// </summary>
        [Input("organizationId")]
        public Input<string>? OrganizationId { get; set; }

        /// <summary>
        /// Unique identifier for the project to which the target API Key belongs.
        /// Required if `organization_id` is not set.
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        [Input("projectRoles")]
        private InputList<string>? _projectRoles;

        /// <summary>
        /// Roles assigned when an org API key is assigned to a project API key. Possible values are `GROUP_CLUSTER_MANAGER`, `GROUP_DATA_ACCESS_ADMIN`, `GROUP_DATA_ACCESS_READ_ONLY`, `GROUP_DATA_ACCESS_READ_WRITE`, `GROUP_OWNER` and `GROUP_READ_ONLY`.
        /// </summary>
        public InputList<string> ProjectRoles
        {
            get => _projectRoles ?? (_projectRoles = new InputList<string>());
            set => _projectRoles = value;
        }

        [Input("roles", required: true)]
        private InputList<string>? _roles;

        /// <summary>
        /// List of roles that the API Key needs to have. Possible values are `ORG_OWNER`, `ORG_MEMBER`, `ORG_GROUP_CREATOR`, `ORG_BILLING_ADMIN` and `ORG_READ_ONLY`.
        /// </summary>
        public InputList<string> Roles
        {
            get => _roles ?? (_roles = new InputList<string>());
            set => _roles = value;
        }

        /// <summary>
        /// Duration in seconds after which the issued credential should expire.
        /// </summary>
        [Input("ttl")]
        public Input<string>? Ttl { get; set; }

        public SecretRoleArgs()
        {
        }
        public static new SecretRoleArgs Empty => new SecretRoleArgs();
    }

    public sealed class SecretRoleState : global::Pulumi.ResourceArgs
    {
        [Input("cidrBlocks")]
        private InputList<string>? _cidrBlocks;

        /// <summary>
        /// Whitelist entry in CIDR notation to be added for the API key.
        /// </summary>
        public InputList<string> CidrBlocks
        {
            get => _cidrBlocks ?? (_cidrBlocks = new InputList<string>());
            set => _cidrBlocks = value;
        }

        [Input("ipAddresses")]
        private InputList<string>? _ipAddresses;

        /// <summary>
        /// IP address to be added to the whitelist for the API key.
        /// </summary>
        public InputList<string> IpAddresses
        {
            get => _ipAddresses ?? (_ipAddresses = new InputList<string>());
            set => _ipAddresses = value;
        }

        /// <summary>
        /// The maximum allowed lifetime of credentials issued using this role.
        /// </summary>
        [Input("maxTtl")]
        public Input<string>? MaxTtl { get; set; }

        /// <summary>
        /// Path where the MongoDB Atlas Secrets Engine is mounted.
        /// </summary>
        [Input("mount")]
        public Input<string>? Mount { get; set; }

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
        /// Unique identifier for the organization to which the target API Key belongs. 
        /// Required if `project_id` is not set.
        /// </summary>
        [Input("organizationId")]
        public Input<string>? OrganizationId { get; set; }

        /// <summary>
        /// Unique identifier for the project to which the target API Key belongs.
        /// Required if `organization_id` is not set.
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        [Input("projectRoles")]
        private InputList<string>? _projectRoles;

        /// <summary>
        /// Roles assigned when an org API key is assigned to a project API key. Possible values are `GROUP_CLUSTER_MANAGER`, `GROUP_DATA_ACCESS_ADMIN`, `GROUP_DATA_ACCESS_READ_ONLY`, `GROUP_DATA_ACCESS_READ_WRITE`, `GROUP_OWNER` and `GROUP_READ_ONLY`.
        /// </summary>
        public InputList<string> ProjectRoles
        {
            get => _projectRoles ?? (_projectRoles = new InputList<string>());
            set => _projectRoles = value;
        }

        [Input("roles")]
        private InputList<string>? _roles;

        /// <summary>
        /// List of roles that the API Key needs to have. Possible values are `ORG_OWNER`, `ORG_MEMBER`, `ORG_GROUP_CREATOR`, `ORG_BILLING_ADMIN` and `ORG_READ_ONLY`.
        /// </summary>
        public InputList<string> Roles
        {
            get => _roles ?? (_roles = new InputList<string>());
            set => _roles = value;
        }

        /// <summary>
        /// Duration in seconds after which the issued credential should expire.
        /// </summary>
        [Input("ttl")]
        public Input<string>? Ttl { get; set; }

        public SecretRoleState()
        {
        }
        public static new SecretRoleState Empty => new SecretRoleState();
    }
}
