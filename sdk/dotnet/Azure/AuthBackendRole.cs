// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vault.Azure
{
    /// <summary>
    /// Manages an Azure auth backend role in a Vault server. Roles constrain the
    /// instances or principals that can perform the login operation against the
    /// backend. See the [Vault
    /// documentation](https://www.vaultproject.io/docs/auth/azure.html) for more
    /// information.
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Vault = Pulumi.Vault;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var azure = new Vault.AuthBackend("azure", new Vault.AuthBackendArgs
    ///         {
    ///             Type = "azure",
    ///         });
    ///         var example = new Vault.Azure.AuthBackendRole("example", new Vault.Azure.AuthBackendRoleArgs
    ///         {
    ///             Backend = azure.Path,
    ///             BoundResourceGroups = 
    ///             {
    ///                 "123456789012",
    ///             },
    ///             BoundSubscriptionIds = 
    ///             {
    ///                 "11111111-2222-3333-4444-555555555555",
    ///             },
    ///             MaxTtl = 120,
    ///             Policies = 
    ///             {
    ///                 "default",
    ///                 "dev",
    ///                 "prod",
    ///             },
    ///             Role = "test-role",
    ///             Ttl = 60,
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// Azure auth backend roles can be imported using `auth/`, the `backend` path, `/role/`, and the `role` name e.g.
    /// 
    /// ```sh
    ///  $ pulumi import vault:azure/authBackendRole:AuthBackendRole example auth/azure/role/test-role
    /// ```
    /// </summary>
    public partial class AuthBackendRole : Pulumi.CustomResource
    {
        /// <summary>
        /// Unique name of the auth backend to configure.
        /// </summary>
        [Output("backend")]
        public Output<string?> Backend { get; private set; } = null!;

        /// <summary>
        /// If set, defines a constraint on the groups
        /// that can perform the login operation that they should be using the group
        /// ID specified by this field.
        /// </summary>
        [Output("boundGroupIds")]
        public Output<ImmutableArray<string>> BoundGroupIds { get; private set; } = null!;

        /// <summary>
        /// If set, defines a constraint on the virtual machines
        /// that can perform the login operation that the location in their identity
        /// document must match the one specified by this field.
        /// </summary>
        [Output("boundLocations")]
        public Output<ImmutableArray<string>> BoundLocations { get; private set; } = null!;

        /// <summary>
        /// If set, defines a constraint on the virtual
        /// machiness that can perform the login operation that they be associated with
        /// the resource group that matches the value specified by this field.
        /// </summary>
        [Output("boundResourceGroups")]
        public Output<ImmutableArray<string>> BoundResourceGroups { get; private set; } = null!;

        /// <summary>
        /// If set, defines a constraint on the virtual
        /// machines that can perform the login operation that they must match the scale set
        /// specified by this field.
        /// </summary>
        [Output("boundScaleSets")]
        public Output<ImmutableArray<string>> BoundScaleSets { get; private set; } = null!;

        /// <summary>
        /// If set, defines a constraint on the
        /// service principals that can perform the login operation that they should be posess
        /// the ids specified by this field.
        /// </summary>
        [Output("boundServicePrincipalIds")]
        public Output<ImmutableArray<string>> BoundServicePrincipalIds { get; private set; } = null!;

        /// <summary>
        /// If set, defines a constraint on the subscriptions
        /// that can perform the login operation to ones which  matches the value specified by this
        /// field.
        /// </summary>
        [Output("boundSubscriptionIds")]
        public Output<ImmutableArray<string>> BoundSubscriptionIds { get; private set; } = null!;

        /// <summary>
        /// The maximum allowed lifetime of tokens issued using
        /// this role, provided as a number of seconds.
        /// </summary>
        [Output("maxTtl")]
        public Output<int?> MaxTtl { get; private set; } = null!;

        /// <summary>
        /// If set, indicates that the token generated using this
        /// role should never expire. The token should be renewed within the duration
        /// specified by this value. At each renewal, the token's TTL will be set to the
        /// value of this field. The maximum allowed lifetime of token issued using this
        /// role. Specified as a number of seconds.
        /// </summary>
        [Output("period")]
        public Output<int?> Period { get; private set; } = null!;

        /// <summary>
        /// An array of strings specifying the policies to be set
        /// on tokens issued using this role.
        /// </summary>
        [Output("policies")]
        public Output<ImmutableArray<string>> Policies { get; private set; } = null!;

        /// <summary>
        /// The name of the role.
        /// </summary>
        [Output("role")]
        public Output<string> Role { get; private set; } = null!;

        /// <summary>
        /// Specifies the blocks of IP addresses which are allowed to use the generated token
        /// </summary>
        [Output("tokenBoundCidrs")]
        public Output<ImmutableArray<string>> TokenBoundCidrs { get; private set; } = null!;

        /// <summary>
        /// Generated Token's Explicit Maximum TTL in seconds
        /// </summary>
        [Output("tokenExplicitMaxTtl")]
        public Output<int?> TokenExplicitMaxTtl { get; private set; } = null!;

        /// <summary>
        /// The maximum lifetime of the generated token
        /// </summary>
        [Output("tokenMaxTtl")]
        public Output<int?> TokenMaxTtl { get; private set; } = null!;

        /// <summary>
        /// If true, the 'default' policy will not automatically be added to generated tokens
        /// </summary>
        [Output("tokenNoDefaultPolicy")]
        public Output<bool?> TokenNoDefaultPolicy { get; private set; } = null!;

        /// <summary>
        /// The maximum number of times a token may be used, a value of zero means unlimited
        /// </summary>
        [Output("tokenNumUses")]
        public Output<int?> TokenNumUses { get; private set; } = null!;

        /// <summary>
        /// Generated Token's Period
        /// </summary>
        [Output("tokenPeriod")]
        public Output<int?> TokenPeriod { get; private set; } = null!;

        /// <summary>
        /// Generated Token's Policies
        /// </summary>
        [Output("tokenPolicies")]
        public Output<ImmutableArray<string>> TokenPolicies { get; private set; } = null!;

        /// <summary>
        /// The initial ttl of the token to generate in seconds
        /// </summary>
        [Output("tokenTtl")]
        public Output<int?> TokenTtl { get; private set; } = null!;

        /// <summary>
        /// The type of token to generate, service or batch
        /// </summary>
        [Output("tokenType")]
        public Output<string?> TokenType { get; private set; } = null!;

        /// <summary>
        /// The TTL period of tokens issued using this role, provided
        /// as a number of seconds.
        /// </summary>
        [Output("ttl")]
        public Output<int?> Ttl { get; private set; } = null!;


        /// <summary>
        /// Create a AuthBackendRole resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AuthBackendRole(string name, AuthBackendRoleArgs args, CustomResourceOptions? options = null)
            : base("vault:azure/authBackendRole:AuthBackendRole", name, args ?? new AuthBackendRoleArgs(), MakeResourceOptions(options, ""))
        {
        }

        private AuthBackendRole(string name, Input<string> id, AuthBackendRoleState? state = null, CustomResourceOptions? options = null)
            : base("vault:azure/authBackendRole:AuthBackendRole", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing AuthBackendRole resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static AuthBackendRole Get(string name, Input<string> id, AuthBackendRoleState? state = null, CustomResourceOptions? options = null)
        {
            return new AuthBackendRole(name, id, state, options);
        }
    }

    public sealed class AuthBackendRoleArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Unique name of the auth backend to configure.
        /// </summary>
        [Input("backend")]
        public Input<string>? Backend { get; set; }

        [Input("boundGroupIds")]
        private InputList<string>? _boundGroupIds;

        /// <summary>
        /// If set, defines a constraint on the groups
        /// that can perform the login operation that they should be using the group
        /// ID specified by this field.
        /// </summary>
        public InputList<string> BoundGroupIds
        {
            get => _boundGroupIds ?? (_boundGroupIds = new InputList<string>());
            set => _boundGroupIds = value;
        }

        [Input("boundLocations")]
        private InputList<string>? _boundLocations;

        /// <summary>
        /// If set, defines a constraint on the virtual machines
        /// that can perform the login operation that the location in their identity
        /// document must match the one specified by this field.
        /// </summary>
        public InputList<string> BoundLocations
        {
            get => _boundLocations ?? (_boundLocations = new InputList<string>());
            set => _boundLocations = value;
        }

        [Input("boundResourceGroups")]
        private InputList<string>? _boundResourceGroups;

        /// <summary>
        /// If set, defines a constraint on the virtual
        /// machiness that can perform the login operation that they be associated with
        /// the resource group that matches the value specified by this field.
        /// </summary>
        public InputList<string> BoundResourceGroups
        {
            get => _boundResourceGroups ?? (_boundResourceGroups = new InputList<string>());
            set => _boundResourceGroups = value;
        }

        [Input("boundScaleSets")]
        private InputList<string>? _boundScaleSets;

        /// <summary>
        /// If set, defines a constraint on the virtual
        /// machines that can perform the login operation that they must match the scale set
        /// specified by this field.
        /// </summary>
        public InputList<string> BoundScaleSets
        {
            get => _boundScaleSets ?? (_boundScaleSets = new InputList<string>());
            set => _boundScaleSets = value;
        }

        [Input("boundServicePrincipalIds")]
        private InputList<string>? _boundServicePrincipalIds;

        /// <summary>
        /// If set, defines a constraint on the
        /// service principals that can perform the login operation that they should be posess
        /// the ids specified by this field.
        /// </summary>
        public InputList<string> BoundServicePrincipalIds
        {
            get => _boundServicePrincipalIds ?? (_boundServicePrincipalIds = new InputList<string>());
            set => _boundServicePrincipalIds = value;
        }

        [Input("boundSubscriptionIds")]
        private InputList<string>? _boundSubscriptionIds;

        /// <summary>
        /// If set, defines a constraint on the subscriptions
        /// that can perform the login operation to ones which  matches the value specified by this
        /// field.
        /// </summary>
        public InputList<string> BoundSubscriptionIds
        {
            get => _boundSubscriptionIds ?? (_boundSubscriptionIds = new InputList<string>());
            set => _boundSubscriptionIds = value;
        }

        /// <summary>
        /// The maximum allowed lifetime of tokens issued using
        /// this role, provided as a number of seconds.
        /// </summary>
        [Input("maxTtl")]
        public Input<int>? MaxTtl { get; set; }

        /// <summary>
        /// If set, indicates that the token generated using this
        /// role should never expire. The token should be renewed within the duration
        /// specified by this value. At each renewal, the token's TTL will be set to the
        /// value of this field. The maximum allowed lifetime of token issued using this
        /// role. Specified as a number of seconds.
        /// </summary>
        [Input("period")]
        public Input<int>? Period { get; set; }

        [Input("policies")]
        private InputList<string>? _policies;

        /// <summary>
        /// An array of strings specifying the policies to be set
        /// on tokens issued using this role.
        /// </summary>
        [Obsolete(@"use `token_policies` instead if you are running Vault >= 1.2")]
        public InputList<string> Policies
        {
            get => _policies ?? (_policies = new InputList<string>());
            set => _policies = value;
        }

        /// <summary>
        /// The name of the role.
        /// </summary>
        [Input("role", required: true)]
        public Input<string> Role { get; set; } = null!;

        [Input("tokenBoundCidrs")]
        private InputList<string>? _tokenBoundCidrs;

        /// <summary>
        /// Specifies the blocks of IP addresses which are allowed to use the generated token
        /// </summary>
        public InputList<string> TokenBoundCidrs
        {
            get => _tokenBoundCidrs ?? (_tokenBoundCidrs = new InputList<string>());
            set => _tokenBoundCidrs = value;
        }

        /// <summary>
        /// Generated Token's Explicit Maximum TTL in seconds
        /// </summary>
        [Input("tokenExplicitMaxTtl")]
        public Input<int>? TokenExplicitMaxTtl { get; set; }

        /// <summary>
        /// The maximum lifetime of the generated token
        /// </summary>
        [Input("tokenMaxTtl")]
        public Input<int>? TokenMaxTtl { get; set; }

        /// <summary>
        /// If true, the 'default' policy will not automatically be added to generated tokens
        /// </summary>
        [Input("tokenNoDefaultPolicy")]
        public Input<bool>? TokenNoDefaultPolicy { get; set; }

        /// <summary>
        /// The maximum number of times a token may be used, a value of zero means unlimited
        /// </summary>
        [Input("tokenNumUses")]
        public Input<int>? TokenNumUses { get; set; }

        /// <summary>
        /// Generated Token's Period
        /// </summary>
        [Input("tokenPeriod")]
        public Input<int>? TokenPeriod { get; set; }

        [Input("tokenPolicies")]
        private InputList<string>? _tokenPolicies;

        /// <summary>
        /// Generated Token's Policies
        /// </summary>
        public InputList<string> TokenPolicies
        {
            get => _tokenPolicies ?? (_tokenPolicies = new InputList<string>());
            set => _tokenPolicies = value;
        }

        /// <summary>
        /// The initial ttl of the token to generate in seconds
        /// </summary>
        [Input("tokenTtl")]
        public Input<int>? TokenTtl { get; set; }

        /// <summary>
        /// The type of token to generate, service or batch
        /// </summary>
        [Input("tokenType")]
        public Input<string>? TokenType { get; set; }

        /// <summary>
        /// The TTL period of tokens issued using this role, provided
        /// as a number of seconds.
        /// </summary>
        [Input("ttl")]
        public Input<int>? Ttl { get; set; }

        public AuthBackendRoleArgs()
        {
        }
    }

    public sealed class AuthBackendRoleState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Unique name of the auth backend to configure.
        /// </summary>
        [Input("backend")]
        public Input<string>? Backend { get; set; }

        [Input("boundGroupIds")]
        private InputList<string>? _boundGroupIds;

        /// <summary>
        /// If set, defines a constraint on the groups
        /// that can perform the login operation that they should be using the group
        /// ID specified by this field.
        /// </summary>
        public InputList<string> BoundGroupIds
        {
            get => _boundGroupIds ?? (_boundGroupIds = new InputList<string>());
            set => _boundGroupIds = value;
        }

        [Input("boundLocations")]
        private InputList<string>? _boundLocations;

        /// <summary>
        /// If set, defines a constraint on the virtual machines
        /// that can perform the login operation that the location in their identity
        /// document must match the one specified by this field.
        /// </summary>
        public InputList<string> BoundLocations
        {
            get => _boundLocations ?? (_boundLocations = new InputList<string>());
            set => _boundLocations = value;
        }

        [Input("boundResourceGroups")]
        private InputList<string>? _boundResourceGroups;

        /// <summary>
        /// If set, defines a constraint on the virtual
        /// machiness that can perform the login operation that they be associated with
        /// the resource group that matches the value specified by this field.
        /// </summary>
        public InputList<string> BoundResourceGroups
        {
            get => _boundResourceGroups ?? (_boundResourceGroups = new InputList<string>());
            set => _boundResourceGroups = value;
        }

        [Input("boundScaleSets")]
        private InputList<string>? _boundScaleSets;

        /// <summary>
        /// If set, defines a constraint on the virtual
        /// machines that can perform the login operation that they must match the scale set
        /// specified by this field.
        /// </summary>
        public InputList<string> BoundScaleSets
        {
            get => _boundScaleSets ?? (_boundScaleSets = new InputList<string>());
            set => _boundScaleSets = value;
        }

        [Input("boundServicePrincipalIds")]
        private InputList<string>? _boundServicePrincipalIds;

        /// <summary>
        /// If set, defines a constraint on the
        /// service principals that can perform the login operation that they should be posess
        /// the ids specified by this field.
        /// </summary>
        public InputList<string> BoundServicePrincipalIds
        {
            get => _boundServicePrincipalIds ?? (_boundServicePrincipalIds = new InputList<string>());
            set => _boundServicePrincipalIds = value;
        }

        [Input("boundSubscriptionIds")]
        private InputList<string>? _boundSubscriptionIds;

        /// <summary>
        /// If set, defines a constraint on the subscriptions
        /// that can perform the login operation to ones which  matches the value specified by this
        /// field.
        /// </summary>
        public InputList<string> BoundSubscriptionIds
        {
            get => _boundSubscriptionIds ?? (_boundSubscriptionIds = new InputList<string>());
            set => _boundSubscriptionIds = value;
        }

        /// <summary>
        /// The maximum allowed lifetime of tokens issued using
        /// this role, provided as a number of seconds.
        /// </summary>
        [Input("maxTtl")]
        public Input<int>? MaxTtl { get; set; }

        /// <summary>
        /// If set, indicates that the token generated using this
        /// role should never expire. The token should be renewed within the duration
        /// specified by this value. At each renewal, the token's TTL will be set to the
        /// value of this field. The maximum allowed lifetime of token issued using this
        /// role. Specified as a number of seconds.
        /// </summary>
        [Input("period")]
        public Input<int>? Period { get; set; }

        [Input("policies")]
        private InputList<string>? _policies;

        /// <summary>
        /// An array of strings specifying the policies to be set
        /// on tokens issued using this role.
        /// </summary>
        [Obsolete(@"use `token_policies` instead if you are running Vault >= 1.2")]
        public InputList<string> Policies
        {
            get => _policies ?? (_policies = new InputList<string>());
            set => _policies = value;
        }

        /// <summary>
        /// The name of the role.
        /// </summary>
        [Input("role")]
        public Input<string>? Role { get; set; }

        [Input("tokenBoundCidrs")]
        private InputList<string>? _tokenBoundCidrs;

        /// <summary>
        /// Specifies the blocks of IP addresses which are allowed to use the generated token
        /// </summary>
        public InputList<string> TokenBoundCidrs
        {
            get => _tokenBoundCidrs ?? (_tokenBoundCidrs = new InputList<string>());
            set => _tokenBoundCidrs = value;
        }

        /// <summary>
        /// Generated Token's Explicit Maximum TTL in seconds
        /// </summary>
        [Input("tokenExplicitMaxTtl")]
        public Input<int>? TokenExplicitMaxTtl { get; set; }

        /// <summary>
        /// The maximum lifetime of the generated token
        /// </summary>
        [Input("tokenMaxTtl")]
        public Input<int>? TokenMaxTtl { get; set; }

        /// <summary>
        /// If true, the 'default' policy will not automatically be added to generated tokens
        /// </summary>
        [Input("tokenNoDefaultPolicy")]
        public Input<bool>? TokenNoDefaultPolicy { get; set; }

        /// <summary>
        /// The maximum number of times a token may be used, a value of zero means unlimited
        /// </summary>
        [Input("tokenNumUses")]
        public Input<int>? TokenNumUses { get; set; }

        /// <summary>
        /// Generated Token's Period
        /// </summary>
        [Input("tokenPeriod")]
        public Input<int>? TokenPeriod { get; set; }

        [Input("tokenPolicies")]
        private InputList<string>? _tokenPolicies;

        /// <summary>
        /// Generated Token's Policies
        /// </summary>
        public InputList<string> TokenPolicies
        {
            get => _tokenPolicies ?? (_tokenPolicies = new InputList<string>());
            set => _tokenPolicies = value;
        }

        /// <summary>
        /// The initial ttl of the token to generate in seconds
        /// </summary>
        [Input("tokenTtl")]
        public Input<int>? TokenTtl { get; set; }

        /// <summary>
        /// The type of token to generate, service or batch
        /// </summary>
        [Input("tokenType")]
        public Input<string>? TokenType { get; set; }

        /// <summary>
        /// The TTL period of tokens issued using this role, provided
        /// as a number of seconds.
        /// </summary>
        [Input("ttl")]
        public Input<int>? Ttl { get; set; }

        public AuthBackendRoleState()
        {
        }
    }
}
