// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vault.AD
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
    ///     var config = new Vault.AD.SecretBackend("config", new()
    ///     {
    ///         Backend = "ad",
    ///         Binddn = "CN=Administrator,CN=Users,DC=corp,DC=example,DC=net",
    ///         Bindpass = "SuperSecretPassw0rd",
    ///         Url = "ldaps://ad",
    ///         InsecureTls = true,
    ///         Userdn = "CN=Users,DC=corp,DC=example,DC=net",
    ///     });
    /// 
    ///     var role = new Vault.AD.SecretRole("role", new()
    ///     {
    ///         Backend = config.Backend,
    ///         Role = "bob",
    ///         ServiceAccountName = "Bob",
    ///         Ttl = 60,
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// AD secret backend roles can be imported using the `path`, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import vault:ad/secretRole:SecretRole role ad/roles/bob
    /// ```
    /// </summary>
    [VaultResourceType("vault:ad/secretRole:SecretRole")]
    public partial class SecretRole : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The path the AD secret backend is mounted at,
        /// with no leading or trailing `/`s.
        /// </summary>
        [Output("backend")]
        public Output<string> Backend { get; private set; } = null!;

        /// <summary>
        /// Timestamp of the last password rotation by Vault.
        /// </summary>
        [Output("lastVaultRotation")]
        public Output<string> LastVaultRotation { get; private set; } = null!;

        /// <summary>
        /// The namespace to provision the resource in.
        /// The value should not contain leading or trailing forward slashes.
        /// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
        /// *Available only for Vault Enterprise*.
        /// </summary>
        [Output("namespace")]
        public Output<string?> Namespace { get; private set; } = null!;

        /// <summary>
        /// Timestamp of the last password set by Vault.
        /// </summary>
        [Output("passwordLastSet")]
        public Output<string> PasswordLastSet { get; private set; } = null!;

        /// <summary>
        /// The name to identify this role within the backend.
        /// Must be unique within the backend.
        /// </summary>
        [Output("role")]
        public Output<string> Role { get; private set; } = null!;

        /// <summary>
        /// Specifies the name of the Active Directory service
        /// account mapped to this role.
        /// </summary>
        [Output("serviceAccountName")]
        public Output<string> ServiceAccountName { get; private set; } = null!;

        /// <summary>
        /// The password time-to-live in seconds. Defaults to the configuration
        /// ttl if not provided.
        /// </summary>
        [Output("ttl")]
        public Output<int?> Ttl { get; private set; } = null!;


        /// <summary>
        /// Create a SecretRole resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public SecretRole(string name, SecretRoleArgs args, CustomResourceOptions? options = null)
            : base("vault:ad/secretRole:SecretRole", name, args ?? new SecretRoleArgs(), MakeResourceOptions(options, ""))
        {
        }

        private SecretRole(string name, Input<string> id, SecretRoleState? state = null, CustomResourceOptions? options = null)
            : base("vault:ad/secretRole:SecretRole", name, state, MakeResourceOptions(options, id))
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
        /// The path the AD secret backend is mounted at,
        /// with no leading or trailing `/`s.
        /// </summary>
        [Input("backend", required: true)]
        public Input<string> Backend { get; set; } = null!;

        /// <summary>
        /// The namespace to provision the resource in.
        /// The value should not contain leading or trailing forward slashes.
        /// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
        /// *Available only for Vault Enterprise*.
        /// </summary>
        [Input("namespace")]
        public Input<string>? Namespace { get; set; }

        /// <summary>
        /// The name to identify this role within the backend.
        /// Must be unique within the backend.
        /// </summary>
        [Input("role", required: true)]
        public Input<string> Role { get; set; } = null!;

        /// <summary>
        /// Specifies the name of the Active Directory service
        /// account mapped to this role.
        /// </summary>
        [Input("serviceAccountName", required: true)]
        public Input<string> ServiceAccountName { get; set; } = null!;

        /// <summary>
        /// The password time-to-live in seconds. Defaults to the configuration
        /// ttl if not provided.
        /// </summary>
        [Input("ttl")]
        public Input<int>? Ttl { get; set; }

        public SecretRoleArgs()
        {
        }
        public static new SecretRoleArgs Empty => new SecretRoleArgs();
    }

    public sealed class SecretRoleState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The path the AD secret backend is mounted at,
        /// with no leading or trailing `/`s.
        /// </summary>
        [Input("backend")]
        public Input<string>? Backend { get; set; }

        /// <summary>
        /// Timestamp of the last password rotation by Vault.
        /// </summary>
        [Input("lastVaultRotation")]
        public Input<string>? LastVaultRotation { get; set; }

        /// <summary>
        /// The namespace to provision the resource in.
        /// The value should not contain leading or trailing forward slashes.
        /// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
        /// *Available only for Vault Enterprise*.
        /// </summary>
        [Input("namespace")]
        public Input<string>? Namespace { get; set; }

        /// <summary>
        /// Timestamp of the last password set by Vault.
        /// </summary>
        [Input("passwordLastSet")]
        public Input<string>? PasswordLastSet { get; set; }

        /// <summary>
        /// The name to identify this role within the backend.
        /// Must be unique within the backend.
        /// </summary>
        [Input("role")]
        public Input<string>? Role { get; set; }

        /// <summary>
        /// Specifies the name of the Active Directory service
        /// account mapped to this role.
        /// </summary>
        [Input("serviceAccountName")]
        public Input<string>? ServiceAccountName { get; set; }

        /// <summary>
        /// The password time-to-live in seconds. Defaults to the configuration
        /// ttl if not provided.
        /// </summary>
        [Input("ttl")]
        public Input<int>? Ttl { get; set; }

        public SecretRoleState()
        {
        }
        public static new SecretRoleState Empty => new SecretRoleState();
    }
}
