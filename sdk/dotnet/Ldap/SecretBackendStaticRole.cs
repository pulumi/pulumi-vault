// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vault.Ldap
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
    ///     var config = new Vault.Ldap.SecretBackend("config", new()
    ///     {
    ///         Path = "my-custom-ldap",
    ///         Binddn = "CN=Administrator,CN=Users,DC=corp,DC=example,DC=net",
    ///         Bindpass = "SuperSecretPassw0rd",
    ///         Url = "ldaps://localhost",
    ///         InsecureTls = true,
    ///         Userdn = "CN=Users,DC=corp,DC=example,DC=net",
    ///     });
    /// 
    ///     var role = new Vault.Ldap.SecretBackendStaticRole("role", new()
    ///     {
    ///         Mount = config.Path,
    ///         Username = "alice",
    ///         Dn = "cn=alice,ou=Users,DC=corp,DC=example,DC=net",
    ///         RoleName = "alice",
    ///         RotationPeriod = 60,
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// LDAP secret backend static role can be imported using the full path to the role of the form`&lt;mount_path&gt;/static-role/&lt;role_name&gt;` e.g.
    /// 
    /// ```sh
    ///  $ pulumi import vault:ldap/secretBackendStaticRole:SecretBackendStaticRole role ldap/static-role/example-role
    /// ```
    /// </summary>
    [VaultResourceType("vault:ldap/secretBackendStaticRole:SecretBackendStaticRole")]
    public partial class SecretBackendStaticRole : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Distinguished name (DN) of the existing LDAP entry to manage
        /// password rotation for. If given, it will take precedence over `username` for the LDAP
        /// search performed during password rotation. Cannot be modified after creation.
        /// </summary>
        [Output("dn")]
        public Output<string?> Dn { get; private set; } = null!;

        /// <summary>
        /// The unique path this backend should be mounted at. Must
        /// not begin or end with a `/`. Defaults to `ldap`.
        /// </summary>
        [Output("mount")]
        public Output<string?> Mount { get; private set; } = null!;

        /// <summary>
        /// The namespace to provision the resource in.
        /// The value should not contain leading or trailing forward slashes.
        /// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
        /// *Available only for Vault Enterprise*.
        /// </summary>
        [Output("namespace")]
        public Output<string?> Namespace { get; private set; } = null!;

        /// <summary>
        /// Name of the role.
        /// </summary>
        [Output("roleName")]
        public Output<string> RoleName { get; private set; } = null!;

        /// <summary>
        /// How often Vault should rotate the password of the user entry.
        /// </summary>
        [Output("rotationPeriod")]
        public Output<int> RotationPeriod { get; private set; } = null!;

        /// <summary>
        /// The username of the existing LDAP entry to manage password rotation for.
        /// </summary>
        [Output("username")]
        public Output<string> Username { get; private set; } = null!;


        /// <summary>
        /// Create a SecretBackendStaticRole resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public SecretBackendStaticRole(string name, SecretBackendStaticRoleArgs args, CustomResourceOptions? options = null)
            : base("vault:ldap/secretBackendStaticRole:SecretBackendStaticRole", name, args ?? new SecretBackendStaticRoleArgs(), MakeResourceOptions(options, ""))
        {
        }

        private SecretBackendStaticRole(string name, Input<string> id, SecretBackendStaticRoleState? state = null, CustomResourceOptions? options = null)
            : base("vault:ldap/secretBackendStaticRole:SecretBackendStaticRole", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing SecretBackendStaticRole resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static SecretBackendStaticRole Get(string name, Input<string> id, SecretBackendStaticRoleState? state = null, CustomResourceOptions? options = null)
        {
            return new SecretBackendStaticRole(name, id, state, options);
        }
    }

    public sealed class SecretBackendStaticRoleArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Distinguished name (DN) of the existing LDAP entry to manage
        /// password rotation for. If given, it will take precedence over `username` for the LDAP
        /// search performed during password rotation. Cannot be modified after creation.
        /// </summary>
        [Input("dn")]
        public Input<string>? Dn { get; set; }

        /// <summary>
        /// The unique path this backend should be mounted at. Must
        /// not begin or end with a `/`. Defaults to `ldap`.
        /// </summary>
        [Input("mount")]
        public Input<string>? Mount { get; set; }

        /// <summary>
        /// The namespace to provision the resource in.
        /// The value should not contain leading or trailing forward slashes.
        /// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
        /// *Available only for Vault Enterprise*.
        /// </summary>
        [Input("namespace")]
        public Input<string>? Namespace { get; set; }

        /// <summary>
        /// Name of the role.
        /// </summary>
        [Input("roleName", required: true)]
        public Input<string> RoleName { get; set; } = null!;

        /// <summary>
        /// How often Vault should rotate the password of the user entry.
        /// </summary>
        [Input("rotationPeriod", required: true)]
        public Input<int> RotationPeriod { get; set; } = null!;

        /// <summary>
        /// The username of the existing LDAP entry to manage password rotation for.
        /// </summary>
        [Input("username", required: true)]
        public Input<string> Username { get; set; } = null!;

        public SecretBackendStaticRoleArgs()
        {
        }
        public static new SecretBackendStaticRoleArgs Empty => new SecretBackendStaticRoleArgs();
    }

    public sealed class SecretBackendStaticRoleState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Distinguished name (DN) of the existing LDAP entry to manage
        /// password rotation for. If given, it will take precedence over `username` for the LDAP
        /// search performed during password rotation. Cannot be modified after creation.
        /// </summary>
        [Input("dn")]
        public Input<string>? Dn { get; set; }

        /// <summary>
        /// The unique path this backend should be mounted at. Must
        /// not begin or end with a `/`. Defaults to `ldap`.
        /// </summary>
        [Input("mount")]
        public Input<string>? Mount { get; set; }

        /// <summary>
        /// The namespace to provision the resource in.
        /// The value should not contain leading or trailing forward slashes.
        /// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
        /// *Available only for Vault Enterprise*.
        /// </summary>
        [Input("namespace")]
        public Input<string>? Namespace { get; set; }

        /// <summary>
        /// Name of the role.
        /// </summary>
        [Input("roleName")]
        public Input<string>? RoleName { get; set; }

        /// <summary>
        /// How often Vault should rotate the password of the user entry.
        /// </summary>
        [Input("rotationPeriod")]
        public Input<int>? RotationPeriod { get; set; }

        /// <summary>
        /// The username of the existing LDAP entry to manage password rotation for.
        /// </summary>
        [Input("username")]
        public Input<string>? Username { get; set; }

        public SecretBackendStaticRoleState()
        {
        }
        public static new SecretBackendStaticRoleState Empty => new SecretBackendStaticRoleState();
    }
}
