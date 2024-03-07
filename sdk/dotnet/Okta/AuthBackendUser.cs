// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vault.Okta
{
    /// <summary>
    /// Provides a resource to create a user in an
    /// [Okta auth backend within Vault](https://www.vaultproject.io/docs/auth/okta.html).
    /// 
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
    ///     var example = new Vault.Okta.AuthBackend("example", new()
    ///     {
    ///         Path = "user_okta",
    ///         Organization = "dummy",
    ///     });
    /// 
    ///     var foo = new Vault.Okta.AuthBackendUser("foo", new()
    ///     {
    ///         Path = example.Path,
    ///         Username = "foo",
    ///         Groups = new[]
    ///         {
    ///             "one",
    ///             "two",
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ## Import
    /// 
    /// Okta authentication backend users can be imported using its `path/user` ID format, e.g.
    /// 
    /// ```sh
    /// $ pulumi import vault:okta/authBackendUser:AuthBackendUser example okta/foo
    /// ```
    /// </summary>
    [VaultResourceType("vault:okta/authBackendUser:AuthBackendUser")]
    public partial class AuthBackendUser : global::Pulumi.CustomResource
    {
        /// <summary>
        /// List of Okta groups to associate with this user
        /// </summary>
        [Output("groups")]
        public Output<ImmutableArray<string>> Groups { get; private set; } = null!;

        /// <summary>
        /// The namespace to provision the resource in.
        /// The value should not contain leading or trailing forward slashes.
        /// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
        /// *Available only for Vault Enterprise*.
        /// </summary>
        [Output("namespace")]
        public Output<string?> Namespace { get; private set; } = null!;

        /// <summary>
        /// The path where the Okta auth backend is mounted
        /// </summary>
        [Output("path")]
        public Output<string> Path { get; private set; } = null!;

        /// <summary>
        /// List of Vault policies to associate with this user
        /// </summary>
        [Output("policies")]
        public Output<ImmutableArray<string>> Policies { get; private set; } = null!;

        /// <summary>
        /// Name of the user within Okta
        /// </summary>
        [Output("username")]
        public Output<string> Username { get; private set; } = null!;


        /// <summary>
        /// Create a AuthBackendUser resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AuthBackendUser(string name, AuthBackendUserArgs args, CustomResourceOptions? options = null)
            : base("vault:okta/authBackendUser:AuthBackendUser", name, args ?? new AuthBackendUserArgs(), MakeResourceOptions(options, ""))
        {
        }

        private AuthBackendUser(string name, Input<string> id, AuthBackendUserState? state = null, CustomResourceOptions? options = null)
            : base("vault:okta/authBackendUser:AuthBackendUser", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing AuthBackendUser resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static AuthBackendUser Get(string name, Input<string> id, AuthBackendUserState? state = null, CustomResourceOptions? options = null)
        {
            return new AuthBackendUser(name, id, state, options);
        }
    }

    public sealed class AuthBackendUserArgs : global::Pulumi.ResourceArgs
    {
        [Input("groups")]
        private InputList<string>? _groups;

        /// <summary>
        /// List of Okta groups to associate with this user
        /// </summary>
        public InputList<string> Groups
        {
            get => _groups ?? (_groups = new InputList<string>());
            set => _groups = value;
        }

        /// <summary>
        /// The namespace to provision the resource in.
        /// The value should not contain leading or trailing forward slashes.
        /// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
        /// *Available only for Vault Enterprise*.
        /// </summary>
        [Input("namespace")]
        public Input<string>? Namespace { get; set; }

        /// <summary>
        /// The path where the Okta auth backend is mounted
        /// </summary>
        [Input("path", required: true)]
        public Input<string> Path { get; set; } = null!;

        [Input("policies")]
        private InputList<string>? _policies;

        /// <summary>
        /// List of Vault policies to associate with this user
        /// </summary>
        public InputList<string> Policies
        {
            get => _policies ?? (_policies = new InputList<string>());
            set => _policies = value;
        }

        /// <summary>
        /// Name of the user within Okta
        /// </summary>
        [Input("username", required: true)]
        public Input<string> Username { get; set; } = null!;

        public AuthBackendUserArgs()
        {
        }
        public static new AuthBackendUserArgs Empty => new AuthBackendUserArgs();
    }

    public sealed class AuthBackendUserState : global::Pulumi.ResourceArgs
    {
        [Input("groups")]
        private InputList<string>? _groups;

        /// <summary>
        /// List of Okta groups to associate with this user
        /// </summary>
        public InputList<string> Groups
        {
            get => _groups ?? (_groups = new InputList<string>());
            set => _groups = value;
        }

        /// <summary>
        /// The namespace to provision the resource in.
        /// The value should not contain leading or trailing forward slashes.
        /// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
        /// *Available only for Vault Enterprise*.
        /// </summary>
        [Input("namespace")]
        public Input<string>? Namespace { get; set; }

        /// <summary>
        /// The path where the Okta auth backend is mounted
        /// </summary>
        [Input("path")]
        public Input<string>? Path { get; set; }

        [Input("policies")]
        private InputList<string>? _policies;

        /// <summary>
        /// List of Vault policies to associate with this user
        /// </summary>
        public InputList<string> Policies
        {
            get => _policies ?? (_policies = new InputList<string>());
            set => _policies = value;
        }

        /// <summary>
        /// Name of the user within Okta
        /// </summary>
        [Input("username")]
        public Input<string>? Username { get; set; }

        public AuthBackendUserState()
        {
        }
        public static new AuthBackendUserState Empty => new AuthBackendUserState();
    }
}
