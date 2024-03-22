// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vault.TerraformCloud
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
    ///     var test = new Vault.TerraformCloud.SecretBackend("test", new()
    ///     {
    ///         Backend = "terraform",
    ///         Description = "Manages the Terraform Cloud backend",
    ///         Token = "V0idfhi2iksSDU234ucdbi2nidsi...",
    ///     });
    /// 
    ///     var example = new Vault.TerraformCloud.SecretRole("example", new()
    ///     {
    ///         Backend = test.Backend,
    ///         Organization = "example-organization-name",
    ///         TeamId = "team-ieF4isC...",
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ## Import
    /// 
    /// Terraform Cloud secret backend roles can be imported using the `backend`, `/roles/`, and the `name` e.g.
    /// 
    /// ```sh
    /// $ pulumi import vault:terraformcloud/secretRole:SecretRole example terraform/roles/my-role
    /// ```
    /// </summary>
    [VaultResourceType("vault:terraformcloud/secretRole:SecretRole")]
    public partial class SecretRole : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The path of the Terraform Cloud Secret Backend the role belongs to.
        /// </summary>
        [Output("backend")]
        public Output<string?> Backend { get; private set; } = null!;

        /// <summary>
        /// Maximum TTL for leases associated with this role, in seconds.
        /// </summary>
        [Output("maxTtl")]
        public Output<int?> MaxTtl { get; private set; } = null!;

        /// <summary>
        /// The name of an existing role against which to create this Terraform Cloud credential
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
        /// Name of the Terraform Cloud or Enterprise organization
        /// </summary>
        [Output("organization")]
        public Output<string?> Organization { get; private set; } = null!;

        /// <summary>
        /// ID of the Terraform Cloud or Enterprise team under organization (e.g., settings/teams/team-xxxxxxxxxxxxx)
        /// </summary>
        [Output("teamId")]
        public Output<string?> TeamId { get; private set; } = null!;

        /// <summary>
        /// Specifies the TTL for this role.
        /// </summary>
        [Output("ttl")]
        public Output<int?> Ttl { get; private set; } = null!;

        /// <summary>
        /// ID of the Terraform Cloud or Enterprise user (e.g., user-xxxxxxxxxxxxxxxx)
        /// </summary>
        [Output("userId")]
        public Output<string?> UserId { get; private set; } = null!;


        /// <summary>
        /// Create a SecretRole resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public SecretRole(string name, SecretRoleArgs? args = null, CustomResourceOptions? options = null)
            : base("vault:terraformcloud/secretRole:SecretRole", name, args ?? new SecretRoleArgs(), MakeResourceOptions(options, ""))
        {
        }

        private SecretRole(string name, Input<string> id, SecretRoleState? state = null, CustomResourceOptions? options = null)
            : base("vault:terraformcloud/secretRole:SecretRole", name, state, MakeResourceOptions(options, id))
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
        /// <summary>
        /// The path of the Terraform Cloud Secret Backend the role belongs to.
        /// </summary>
        [Input("backend")]
        public Input<string>? Backend { get; set; }

        /// <summary>
        /// Maximum TTL for leases associated with this role, in seconds.
        /// </summary>
        [Input("maxTtl")]
        public Input<int>? MaxTtl { get; set; }

        /// <summary>
        /// The name of an existing role against which to create this Terraform Cloud credential
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
        /// Name of the Terraform Cloud or Enterprise organization
        /// </summary>
        [Input("organization")]
        public Input<string>? Organization { get; set; }

        /// <summary>
        /// ID of the Terraform Cloud or Enterprise team under organization (e.g., settings/teams/team-xxxxxxxxxxxxx)
        /// </summary>
        [Input("teamId")]
        public Input<string>? TeamId { get; set; }

        /// <summary>
        /// Specifies the TTL for this role.
        /// </summary>
        [Input("ttl")]
        public Input<int>? Ttl { get; set; }

        /// <summary>
        /// ID of the Terraform Cloud or Enterprise user (e.g., user-xxxxxxxxxxxxxxxx)
        /// </summary>
        [Input("userId")]
        public Input<string>? UserId { get; set; }

        public SecretRoleArgs()
        {
        }
        public static new SecretRoleArgs Empty => new SecretRoleArgs();
    }

    public sealed class SecretRoleState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The path of the Terraform Cloud Secret Backend the role belongs to.
        /// </summary>
        [Input("backend")]
        public Input<string>? Backend { get; set; }

        /// <summary>
        /// Maximum TTL for leases associated with this role, in seconds.
        /// </summary>
        [Input("maxTtl")]
        public Input<int>? MaxTtl { get; set; }

        /// <summary>
        /// The name of an existing role against which to create this Terraform Cloud credential
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
        /// Name of the Terraform Cloud or Enterprise organization
        /// </summary>
        [Input("organization")]
        public Input<string>? Organization { get; set; }

        /// <summary>
        /// ID of the Terraform Cloud or Enterprise team under organization (e.g., settings/teams/team-xxxxxxxxxxxxx)
        /// </summary>
        [Input("teamId")]
        public Input<string>? TeamId { get; set; }

        /// <summary>
        /// Specifies the TTL for this role.
        /// </summary>
        [Input("ttl")]
        public Input<int>? Ttl { get; set; }

        /// <summary>
        /// ID of the Terraform Cloud or Enterprise user (e.g., user-xxxxxxxxxxxxxxxx)
        /// </summary>
        [Input("userId")]
        public Input<string>? UserId { get; set; }

        public SecretRoleState()
        {
        }
        public static new SecretRoleState Empty => new SecretRoleState();
    }
}
