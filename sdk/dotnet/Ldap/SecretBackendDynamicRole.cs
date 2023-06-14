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
    ///         Userdn = "CN=Users,DC=corp,DC=example,DC=net",
    ///     });
    /// 
    ///     var role = new Vault.Ldap.SecretBackendDynamicRole("role", new()
    ///     {
    ///         Mount = config.Path,
    ///         RoleName = "alice",
    ///         CreationLdif = @"dn: cn={{.Username}},ou=users,dc=learn,dc=example
    /// objectClass: person
    /// objectClass: top
    /// cn: learn
    /// sn: {{.Password | utf16le | base64}}
    /// memberOf: cn=dev,ou=groups,dc=learn,dc=example
    /// userPassword: {{.Password}}
    /// ",
    ///         DeletionLdif = @"dn: cn={{.Username}},ou=users,dc=learn,dc=example
    /// changetype: delete
    ///   rollback_ldif = &lt;&lt;EOT
    /// dn: cn={{.Username}},ou=users,dc=learn,dc=example
    /// changetype: delete
    /// ",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// LDAP secret backend dynamic role can be imported using the full path to the role of the form`&lt;mount_path&gt;/dynamic-role/&lt;role_name&gt;` e.g.
    /// 
    /// ```sh
    ///  $ pulumi import vault:ldap/secretBackendDynamicRole:SecretBackendDynamicRole role ldap/role/dynamic-role
    /// ```
    /// </summary>
    [VaultResourceType("vault:ldap/secretBackendDynamicRole:SecretBackendDynamicRole")]
    public partial class SecretBackendDynamicRole : global::Pulumi.CustomResource
    {
        /// <summary>
        /// A templatized LDIF string used to create a user
        /// account. This may contain multiple LDIF entries. The `creation_ldif` can also
        /// be used to add the user account to an existing group. All LDIF entries are
        /// performed in order. If Vault encounters an error while executing the
        /// `creation_ldif` it will stop at the first error and not execute any remaining
        /// LDIF entries. If an error occurs and `rollback_ldif` is specified, the LDIF
        /// entries in `rollback_ldif` will be executed. See `rollback_ldif` for more
        /// details. This field may optionally be provided as a base64 encoded string.
        /// </summary>
        [Output("creationLdif")]
        public Output<string> CreationLdif { get; private set; } = null!;

        /// <summary>
        /// Specifies the TTL for the leases associated with this role.
        /// </summary>
        [Output("defaultTtl")]
        public Output<int?> DefaultTtl { get; private set; } = null!;

        /// <summary>
        /// A templatized LDIF string used to delete the
        /// user account once its TTL has expired. This may contain multiple LDIF
        /// entries. All LDIF entries are performed in order. If Vault encounters an
        /// error while executing an entry in the `deletion_ldif` it will attempt to
        /// continue executing any remaining entries. This field may optionally be
        /// provided as a base64 encoded string.
        /// </summary>
        [Output("deletionLdif")]
        public Output<string> DeletionLdif { get; private set; } = null!;

        /// <summary>
        /// Specifies the maximum TTL for the leases associated with this role.
        /// </summary>
        [Output("maxTtl")]
        public Output<int?> MaxTtl { get; private set; } = null!;

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
        /// A templatized LDIF string used to attempt to
        /// rollback any changes in the event that execution of the `creation_ldif` results
        /// in an error. This may contain multiple LDIF entries. All LDIF entries are
        /// performed in order. If Vault encounters an error while executing an entry in
        /// the `rollback_ldif` it will attempt to continue executing any remaining
        /// entries. This field may optionally be provided as a base64 encoded string.
        /// </summary>
        [Output("rollbackLdif")]
        public Output<string?> RollbackLdif { get; private set; } = null!;

        /// <summary>
        /// A template used to generate a dynamic
        /// username. This will be used to fill in the `.Username` field within the
        /// `creation_ldif` string.
        /// </summary>
        [Output("usernameTemplate")]
        public Output<string?> UsernameTemplate { get; private set; } = null!;


        /// <summary>
        /// Create a SecretBackendDynamicRole resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public SecretBackendDynamicRole(string name, SecretBackendDynamicRoleArgs args, CustomResourceOptions? options = null)
            : base("vault:ldap/secretBackendDynamicRole:SecretBackendDynamicRole", name, args ?? new SecretBackendDynamicRoleArgs(), MakeResourceOptions(options, ""))
        {
        }

        private SecretBackendDynamicRole(string name, Input<string> id, SecretBackendDynamicRoleState? state = null, CustomResourceOptions? options = null)
            : base("vault:ldap/secretBackendDynamicRole:SecretBackendDynamicRole", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing SecretBackendDynamicRole resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static SecretBackendDynamicRole Get(string name, Input<string> id, SecretBackendDynamicRoleState? state = null, CustomResourceOptions? options = null)
        {
            return new SecretBackendDynamicRole(name, id, state, options);
        }
    }

    public sealed class SecretBackendDynamicRoleArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// A templatized LDIF string used to create a user
        /// account. This may contain multiple LDIF entries. The `creation_ldif` can also
        /// be used to add the user account to an existing group. All LDIF entries are
        /// performed in order. If Vault encounters an error while executing the
        /// `creation_ldif` it will stop at the first error and not execute any remaining
        /// LDIF entries. If an error occurs and `rollback_ldif` is specified, the LDIF
        /// entries in `rollback_ldif` will be executed. See `rollback_ldif` for more
        /// details. This field may optionally be provided as a base64 encoded string.
        /// </summary>
        [Input("creationLdif", required: true)]
        public Input<string> CreationLdif { get; set; } = null!;

        /// <summary>
        /// Specifies the TTL for the leases associated with this role.
        /// </summary>
        [Input("defaultTtl")]
        public Input<int>? DefaultTtl { get; set; }

        /// <summary>
        /// A templatized LDIF string used to delete the
        /// user account once its TTL has expired. This may contain multiple LDIF
        /// entries. All LDIF entries are performed in order. If Vault encounters an
        /// error while executing an entry in the `deletion_ldif` it will attempt to
        /// continue executing any remaining entries. This field may optionally be
        /// provided as a base64 encoded string.
        /// </summary>
        [Input("deletionLdif", required: true)]
        public Input<string> DeletionLdif { get; set; } = null!;

        /// <summary>
        /// Specifies the maximum TTL for the leases associated with this role.
        /// </summary>
        [Input("maxTtl")]
        public Input<int>? MaxTtl { get; set; }

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
        /// A templatized LDIF string used to attempt to
        /// rollback any changes in the event that execution of the `creation_ldif` results
        /// in an error. This may contain multiple LDIF entries. All LDIF entries are
        /// performed in order. If Vault encounters an error while executing an entry in
        /// the `rollback_ldif` it will attempt to continue executing any remaining
        /// entries. This field may optionally be provided as a base64 encoded string.
        /// </summary>
        [Input("rollbackLdif")]
        public Input<string>? RollbackLdif { get; set; }

        /// <summary>
        /// A template used to generate a dynamic
        /// username. This will be used to fill in the `.Username` field within the
        /// `creation_ldif` string.
        /// </summary>
        [Input("usernameTemplate")]
        public Input<string>? UsernameTemplate { get; set; }

        public SecretBackendDynamicRoleArgs()
        {
        }
        public static new SecretBackendDynamicRoleArgs Empty => new SecretBackendDynamicRoleArgs();
    }

    public sealed class SecretBackendDynamicRoleState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// A templatized LDIF string used to create a user
        /// account. This may contain multiple LDIF entries. The `creation_ldif` can also
        /// be used to add the user account to an existing group. All LDIF entries are
        /// performed in order. If Vault encounters an error while executing the
        /// `creation_ldif` it will stop at the first error and not execute any remaining
        /// LDIF entries. If an error occurs and `rollback_ldif` is specified, the LDIF
        /// entries in `rollback_ldif` will be executed. See `rollback_ldif` for more
        /// details. This field may optionally be provided as a base64 encoded string.
        /// </summary>
        [Input("creationLdif")]
        public Input<string>? CreationLdif { get; set; }

        /// <summary>
        /// Specifies the TTL for the leases associated with this role.
        /// </summary>
        [Input("defaultTtl")]
        public Input<int>? DefaultTtl { get; set; }

        /// <summary>
        /// A templatized LDIF string used to delete the
        /// user account once its TTL has expired. This may contain multiple LDIF
        /// entries. All LDIF entries are performed in order. If Vault encounters an
        /// error while executing an entry in the `deletion_ldif` it will attempt to
        /// continue executing any remaining entries. This field may optionally be
        /// provided as a base64 encoded string.
        /// </summary>
        [Input("deletionLdif")]
        public Input<string>? DeletionLdif { get; set; }

        /// <summary>
        /// Specifies the maximum TTL for the leases associated with this role.
        /// </summary>
        [Input("maxTtl")]
        public Input<int>? MaxTtl { get; set; }

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
        /// A templatized LDIF string used to attempt to
        /// rollback any changes in the event that execution of the `creation_ldif` results
        /// in an error. This may contain multiple LDIF entries. All LDIF entries are
        /// performed in order. If Vault encounters an error while executing an entry in
        /// the `rollback_ldif` it will attempt to continue executing any remaining
        /// entries. This field may optionally be provided as a base64 encoded string.
        /// </summary>
        [Input("rollbackLdif")]
        public Input<string>? RollbackLdif { get; set; }

        /// <summary>
        /// A template used to generate a dynamic
        /// username. This will be used to fill in the `.Username` field within the
        /// `creation_ldif` string.
        /// </summary>
        [Input("usernameTemplate")]
        public Input<string>? UsernameTemplate { get; set; }

        public SecretBackendDynamicRoleState()
        {
        }
        public static new SecretBackendDynamicRoleState Empty => new SecretBackendDynamicRoleState();
    }
}
