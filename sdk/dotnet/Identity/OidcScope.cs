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
    /// Manages OIDC Scopes in a Vault server. See the [Vault documentation](https://www.vaultproject.io/api-docs/secret/identity/oidc-provider#create-or-update-a-scope)
    /// for more information.
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
    ///     var groups = new Vault.Identity.OidcScope("groups", new()
    ///     {
    ///         Name = "groups",
    ///         Template = "{\"groups\":{{identity.entity.groups.names}}}",
    ///         Description = "Vault OIDC Groups Scope",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// OIDC Scopes can be imported using the `name`, e.g.
    /// 
    /// ```sh
    /// $ pulumi import vault:identity/oidcScope:OidcScope groups groups
    /// ```
    /// </summary>
    [VaultResourceType("vault:identity/oidcScope:OidcScope")]
    public partial class OidcScope : global::Pulumi.CustomResource
    {
        /// <summary>
        /// A description of the scope.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The name of the scope. The `openid` scope name is reserved.
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
        /// The template string for the scope. This may be provided as escaped JSON or base64 encoded JSON.
        /// </summary>
        [Output("template")]
        public Output<string?> Template { get; private set; } = null!;


        /// <summary>
        /// Create a OidcScope resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public OidcScope(string name, OidcScopeArgs? args = null, CustomResourceOptions? options = null)
            : base("vault:identity/oidcScope:OidcScope", name, args ?? new OidcScopeArgs(), MakeResourceOptions(options, ""))
        {
        }

        private OidcScope(string name, Input<string> id, OidcScopeState? state = null, CustomResourceOptions? options = null)
            : base("vault:identity/oidcScope:OidcScope", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing OidcScope resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static OidcScope Get(string name, Input<string> id, OidcScopeState? state = null, CustomResourceOptions? options = null)
        {
            return new OidcScope(name, id, state, options);
        }
    }

    public sealed class OidcScopeArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// A description of the scope.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The name of the scope. The `openid` scope name is reserved.
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
        /// The template string for the scope. This may be provided as escaped JSON or base64 encoded JSON.
        /// </summary>
        [Input("template")]
        public Input<string>? Template { get; set; }

        public OidcScopeArgs()
        {
        }
        public static new OidcScopeArgs Empty => new OidcScopeArgs();
    }

    public sealed class OidcScopeState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// A description of the scope.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The name of the scope. The `openid` scope name is reserved.
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
        /// The template string for the scope. This may be provided as escaped JSON or base64 encoded JSON.
        /// </summary>
        [Input("template")]
        public Input<string>? Template { get; set; }

        public OidcScopeState()
        {
        }
        public static new OidcScopeState Empty => new OidcScopeState();
    }
}
