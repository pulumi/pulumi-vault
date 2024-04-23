// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vault.Consul
{
    /// <summary>
    /// Manages a Consul secrets role for a Consul secrets engine in Vault. Consul secret backends can then issue Consul tokens.
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
    ///     var test = new Vault.Consul.SecretBackend("test", new()
    ///     {
    ///         Path = "consul",
    ///         Description = "Manages the Consul backend",
    ///         Address = "127.0.0.1:8500",
    ///         Token = "4240861b-ce3d-8530-115a-521ff070dd29",
    ///     });
    /// 
    ///     var example = new Vault.Consul.SecretBackendRole("example", new()
    ///     {
    ///         Name = "test-role",
    ///         Backend = test.Path,
    ///         ConsulPolicies = new[]
    ///         {
    ///             "example-policy",
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Note About Required Arguments
    /// 
    /// *At least one* of the four arguments `consul_policies`, `consul_roles`, `service_identities`, or
    /// `node_identities` is required for a token. If desired, any combination of the four arguments up-to and
    /// including all four, is valid.
    /// 
    /// ## Import
    /// 
    /// Consul secret backend roles can be imported using the `backend`, `/roles/`, and the `name` e.g.
    /// 
    /// ```sh
    /// $ pulumi import vault:consul/secretBackendRole:SecretBackendRole example consul/roles/my-role
    /// ```
    /// </summary>
    [VaultResourceType("vault:consul/secretBackendRole:SecretBackendRole")]
    public partial class SecretBackendRole : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The unique name of an existing Consul secrets backend mount. Must not begin or end with a `/`. One of `path` or `backend` is required.
        /// </summary>
        [Output("backend")]
        public Output<string?> Backend { get; private set; } = null!;

        /// <summary>
        /// The Consul namespace that the token will be created in.
        /// Applicable for Vault 1.10+ and Consul 1.7+".
        /// </summary>
        [Output("consulNamespace")]
        public Output<string> ConsulNamespace { get; private set; } = null!;

        /// <summary>
        /// &lt;sup&gt;&lt;a href="#note-about-required-arguments"&gt;SEE NOTE&lt;/a&gt;&lt;/sup&gt; The list of Consul ACL policies to associate with these roles.
        /// </summary>
        [Output("consulPolicies")]
        public Output<ImmutableArray<string>> ConsulPolicies { get; private set; } = null!;

        /// <summary>
        /// &lt;sup&gt;&lt;a href="#note-about-required-arguments"&gt;SEE NOTE&lt;/a&gt;&lt;/sup&gt; Set of Consul roles to attach to the token.
        /// Applicable for Vault 1.10+ with Consul 1.5+.
        /// </summary>
        [Output("consulRoles")]
        public Output<ImmutableArray<string>> ConsulRoles { get; private set; } = null!;

        /// <summary>
        /// Indicates that the token should not be replicated globally and instead be local to the current datacenter.
        /// </summary>
        [Output("local")]
        public Output<bool?> Local { get; private set; } = null!;

        /// <summary>
        /// Maximum TTL for leases associated with this role, in seconds.
        /// </summary>
        [Output("maxTtl")]
        public Output<int?> MaxTtl { get; private set; } = null!;

        /// <summary>
        /// The name of the Consul secrets engine role to create.
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
        /// &lt;sup&gt;&lt;a href="#note-about-required-arguments"&gt;SEE NOTE&lt;/a&gt;&lt;/sup&gt; Set of Consul node
        /// identities to attach to the token. Applicable for Vault 1.11+ with Consul 1.8+.
        /// </summary>
        [Output("nodeIdentities")]
        public Output<ImmutableArray<string>> NodeIdentities { get; private set; } = null!;

        /// <summary>
        /// The admin partition that the token will be created in.
        /// Applicable for Vault 1.10+ and Consul 1.11+".
        /// </summary>
        [Output("partition")]
        public Output<string> Partition { get; private set; } = null!;

        /// <summary>
        /// The list of Consul ACL policies to associate with these roles.
        /// **NOTE:** The new parameter `consul_policies` should be used in favor of this. This parameter,
        /// `policies`, remains supported for legacy users, but Vault has deprecated this field.
        /// </summary>
        [Output("policies")]
        public Output<ImmutableArray<string>> Policies { get; private set; } = null!;

        /// <summary>
        /// &lt;sup&gt;&lt;a href="#note-about-required-arguments"&gt;SEE NOTE&lt;/a&gt;&lt;/sup&gt; Set of Consul
        /// service identities to attach to the token. Applicable for Vault 1.11+ with Consul 1.5+.
        /// </summary>
        [Output("serviceIdentities")]
        public Output<ImmutableArray<string>> ServiceIdentities { get; private set; } = null!;

        /// <summary>
        /// Specifies the TTL for this role.
        /// </summary>
        [Output("ttl")]
        public Output<int?> Ttl { get; private set; } = null!;


        /// <summary>
        /// Create a SecretBackendRole resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public SecretBackendRole(string name, SecretBackendRoleArgs? args = null, CustomResourceOptions? options = null)
            : base("vault:consul/secretBackendRole:SecretBackendRole", name, args ?? new SecretBackendRoleArgs(), MakeResourceOptions(options, ""))
        {
        }

        private SecretBackendRole(string name, Input<string> id, SecretBackendRoleState? state = null, CustomResourceOptions? options = null)
            : base("vault:consul/secretBackendRole:SecretBackendRole", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing SecretBackendRole resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static SecretBackendRole Get(string name, Input<string> id, SecretBackendRoleState? state = null, CustomResourceOptions? options = null)
        {
            return new SecretBackendRole(name, id, state, options);
        }
    }

    public sealed class SecretBackendRoleArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The unique name of an existing Consul secrets backend mount. Must not begin or end with a `/`. One of `path` or `backend` is required.
        /// </summary>
        [Input("backend")]
        public Input<string>? Backend { get; set; }

        /// <summary>
        /// The Consul namespace that the token will be created in.
        /// Applicable for Vault 1.10+ and Consul 1.7+".
        /// </summary>
        [Input("consulNamespace")]
        public Input<string>? ConsulNamespace { get; set; }

        [Input("consulPolicies")]
        private InputList<string>? _consulPolicies;

        /// <summary>
        /// &lt;sup&gt;&lt;a href="#note-about-required-arguments"&gt;SEE NOTE&lt;/a&gt;&lt;/sup&gt; The list of Consul ACL policies to associate with these roles.
        /// </summary>
        public InputList<string> ConsulPolicies
        {
            get => _consulPolicies ?? (_consulPolicies = new InputList<string>());
            set => _consulPolicies = value;
        }

        [Input("consulRoles")]
        private InputList<string>? _consulRoles;

        /// <summary>
        /// &lt;sup&gt;&lt;a href="#note-about-required-arguments"&gt;SEE NOTE&lt;/a&gt;&lt;/sup&gt; Set of Consul roles to attach to the token.
        /// Applicable for Vault 1.10+ with Consul 1.5+.
        /// </summary>
        public InputList<string> ConsulRoles
        {
            get => _consulRoles ?? (_consulRoles = new InputList<string>());
            set => _consulRoles = value;
        }

        /// <summary>
        /// Indicates that the token should not be replicated globally and instead be local to the current datacenter.
        /// </summary>
        [Input("local")]
        public Input<bool>? Local { get; set; }

        /// <summary>
        /// Maximum TTL for leases associated with this role, in seconds.
        /// </summary>
        [Input("maxTtl")]
        public Input<int>? MaxTtl { get; set; }

        /// <summary>
        /// The name of the Consul secrets engine role to create.
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

        [Input("nodeIdentities")]
        private InputList<string>? _nodeIdentities;

        /// <summary>
        /// &lt;sup&gt;&lt;a href="#note-about-required-arguments"&gt;SEE NOTE&lt;/a&gt;&lt;/sup&gt; Set of Consul node
        /// identities to attach to the token. Applicable for Vault 1.11+ with Consul 1.8+.
        /// </summary>
        public InputList<string> NodeIdentities
        {
            get => _nodeIdentities ?? (_nodeIdentities = new InputList<string>());
            set => _nodeIdentities = value;
        }

        /// <summary>
        /// The admin partition that the token will be created in.
        /// Applicable for Vault 1.10+ and Consul 1.11+".
        /// </summary>
        [Input("partition")]
        public Input<string>? Partition { get; set; }

        [Input("policies")]
        private InputList<string>? _policies;

        /// <summary>
        /// The list of Consul ACL policies to associate with these roles.
        /// **NOTE:** The new parameter `consul_policies` should be used in favor of this. This parameter,
        /// `policies`, remains supported for legacy users, but Vault has deprecated this field.
        /// </summary>
        public InputList<string> Policies
        {
            get => _policies ?? (_policies = new InputList<string>());
            set => _policies = value;
        }

        [Input("serviceIdentities")]
        private InputList<string>? _serviceIdentities;

        /// <summary>
        /// &lt;sup&gt;&lt;a href="#note-about-required-arguments"&gt;SEE NOTE&lt;/a&gt;&lt;/sup&gt; Set of Consul
        /// service identities to attach to the token. Applicable for Vault 1.11+ with Consul 1.5+.
        /// </summary>
        public InputList<string> ServiceIdentities
        {
            get => _serviceIdentities ?? (_serviceIdentities = new InputList<string>());
            set => _serviceIdentities = value;
        }

        /// <summary>
        /// Specifies the TTL for this role.
        /// </summary>
        [Input("ttl")]
        public Input<int>? Ttl { get; set; }

        public SecretBackendRoleArgs()
        {
        }
        public static new SecretBackendRoleArgs Empty => new SecretBackendRoleArgs();
    }

    public sealed class SecretBackendRoleState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The unique name of an existing Consul secrets backend mount. Must not begin or end with a `/`. One of `path` or `backend` is required.
        /// </summary>
        [Input("backend")]
        public Input<string>? Backend { get; set; }

        /// <summary>
        /// The Consul namespace that the token will be created in.
        /// Applicable for Vault 1.10+ and Consul 1.7+".
        /// </summary>
        [Input("consulNamespace")]
        public Input<string>? ConsulNamespace { get; set; }

        [Input("consulPolicies")]
        private InputList<string>? _consulPolicies;

        /// <summary>
        /// &lt;sup&gt;&lt;a href="#note-about-required-arguments"&gt;SEE NOTE&lt;/a&gt;&lt;/sup&gt; The list of Consul ACL policies to associate with these roles.
        /// </summary>
        public InputList<string> ConsulPolicies
        {
            get => _consulPolicies ?? (_consulPolicies = new InputList<string>());
            set => _consulPolicies = value;
        }

        [Input("consulRoles")]
        private InputList<string>? _consulRoles;

        /// <summary>
        /// &lt;sup&gt;&lt;a href="#note-about-required-arguments"&gt;SEE NOTE&lt;/a&gt;&lt;/sup&gt; Set of Consul roles to attach to the token.
        /// Applicable for Vault 1.10+ with Consul 1.5+.
        /// </summary>
        public InputList<string> ConsulRoles
        {
            get => _consulRoles ?? (_consulRoles = new InputList<string>());
            set => _consulRoles = value;
        }

        /// <summary>
        /// Indicates that the token should not be replicated globally and instead be local to the current datacenter.
        /// </summary>
        [Input("local")]
        public Input<bool>? Local { get; set; }

        /// <summary>
        /// Maximum TTL for leases associated with this role, in seconds.
        /// </summary>
        [Input("maxTtl")]
        public Input<int>? MaxTtl { get; set; }

        /// <summary>
        /// The name of the Consul secrets engine role to create.
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

        [Input("nodeIdentities")]
        private InputList<string>? _nodeIdentities;

        /// <summary>
        /// &lt;sup&gt;&lt;a href="#note-about-required-arguments"&gt;SEE NOTE&lt;/a&gt;&lt;/sup&gt; Set of Consul node
        /// identities to attach to the token. Applicable for Vault 1.11+ with Consul 1.8+.
        /// </summary>
        public InputList<string> NodeIdentities
        {
            get => _nodeIdentities ?? (_nodeIdentities = new InputList<string>());
            set => _nodeIdentities = value;
        }

        /// <summary>
        /// The admin partition that the token will be created in.
        /// Applicable for Vault 1.10+ and Consul 1.11+".
        /// </summary>
        [Input("partition")]
        public Input<string>? Partition { get; set; }

        [Input("policies")]
        private InputList<string>? _policies;

        /// <summary>
        /// The list of Consul ACL policies to associate with these roles.
        /// **NOTE:** The new parameter `consul_policies` should be used in favor of this. This parameter,
        /// `policies`, remains supported for legacy users, but Vault has deprecated this field.
        /// </summary>
        public InputList<string> Policies
        {
            get => _policies ?? (_policies = new InputList<string>());
            set => _policies = value;
        }

        [Input("serviceIdentities")]
        private InputList<string>? _serviceIdentities;

        /// <summary>
        /// &lt;sup&gt;&lt;a href="#note-about-required-arguments"&gt;SEE NOTE&lt;/a&gt;&lt;/sup&gt; Set of Consul
        /// service identities to attach to the token. Applicable for Vault 1.11+ with Consul 1.5+.
        /// </summary>
        public InputList<string> ServiceIdentities
        {
            get => _serviceIdentities ?? (_serviceIdentities = new InputList<string>());
            set => _serviceIdentities = value;
        }

        /// <summary>
        /// Specifies the TTL for this role.
        /// </summary>
        [Input("ttl")]
        public Input<int>? Ttl { get; set; }

        public SecretBackendRoleState()
        {
        }
        public static new SecretBackendRoleState Empty => new SecretBackendRoleState();
    }
}
