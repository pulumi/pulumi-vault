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
    ///     var key = new Vault.Identity.OidcKey("key", new()
    ///     {
    ///         Algorithm = "RS256",
    ///     });
    /// 
    ///     var roleOidcRole = new Vault.Identity.OidcRole("roleOidcRole", new()
    ///     {
    ///         Key = key.Name,
    ///     });
    /// 
    ///     var roleOidcKeyAllowedClientID = new Vault.Identity.OidcKeyAllowedClientID("roleOidcKeyAllowedClientID", new()
    ///     {
    ///         KeyName = key.Name,
    ///         AllowedClientId = roleOidcRole.ClientId,
    ///     });
    /// 
    /// });
    /// ```
    /// </summary>
    [VaultResourceType("vault:identity/oidcKeyAllowedClientID:OidcKeyAllowedClientID")]
    public partial class OidcKeyAllowedClientID : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Client ID to allow usage with the OIDC named key
        /// </summary>
        [Output("allowedClientId")]
        public Output<string> AllowedClientId { get; private set; } = null!;

        /// <summary>
        /// Name of the OIDC Key allow the Client ID.
        /// </summary>
        [Output("keyName")]
        public Output<string> KeyName { get; private set; } = null!;

        /// <summary>
        /// The namespace to provision the resource in.
        /// The value should not contain leading or trailing forward slashes.
        /// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
        /// *Available only for Vault Enterprise*.
        /// </summary>
        [Output("namespace")]
        public Output<string?> Namespace { get; private set; } = null!;


        /// <summary>
        /// Create a OidcKeyAllowedClientID resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public OidcKeyAllowedClientID(string name, OidcKeyAllowedClientIDArgs args, CustomResourceOptions? options = null)
            : base("vault:identity/oidcKeyAllowedClientID:OidcKeyAllowedClientID", name, args ?? new OidcKeyAllowedClientIDArgs(), MakeResourceOptions(options, ""))
        {
        }

        private OidcKeyAllowedClientID(string name, Input<string> id, OidcKeyAllowedClientIDState? state = null, CustomResourceOptions? options = null)
            : base("vault:identity/oidcKeyAllowedClientID:OidcKeyAllowedClientID", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing OidcKeyAllowedClientID resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static OidcKeyAllowedClientID Get(string name, Input<string> id, OidcKeyAllowedClientIDState? state = null, CustomResourceOptions? options = null)
        {
            return new OidcKeyAllowedClientID(name, id, state, options);
        }
    }

    public sealed class OidcKeyAllowedClientIDArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Client ID to allow usage with the OIDC named key
        /// </summary>
        [Input("allowedClientId", required: true)]
        public Input<string> AllowedClientId { get; set; } = null!;

        /// <summary>
        /// Name of the OIDC Key allow the Client ID.
        /// </summary>
        [Input("keyName", required: true)]
        public Input<string> KeyName { get; set; } = null!;

        /// <summary>
        /// The namespace to provision the resource in.
        /// The value should not contain leading or trailing forward slashes.
        /// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
        /// *Available only for Vault Enterprise*.
        /// </summary>
        [Input("namespace")]
        public Input<string>? Namespace { get; set; }

        public OidcKeyAllowedClientIDArgs()
        {
        }
        public static new OidcKeyAllowedClientIDArgs Empty => new OidcKeyAllowedClientIDArgs();
    }

    public sealed class OidcKeyAllowedClientIDState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Client ID to allow usage with the OIDC named key
        /// </summary>
        [Input("allowedClientId")]
        public Input<string>? AllowedClientId { get; set; }

        /// <summary>
        /// Name of the OIDC Key allow the Client ID.
        /// </summary>
        [Input("keyName")]
        public Input<string>? KeyName { get; set; }

        /// <summary>
        /// The namespace to provision the resource in.
        /// The value should not contain leading or trailing forward slashes.
        /// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
        /// *Available only for Vault Enterprise*.
        /// </summary>
        [Input("namespace")]
        public Input<string>? Namespace { get; set; }

        public OidcKeyAllowedClientIDState()
        {
        }
        public static new OidcKeyAllowedClientIDState Empty => new OidcKeyAllowedClientIDState();
    }
}
