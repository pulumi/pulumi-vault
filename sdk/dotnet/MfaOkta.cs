// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vault
{
    /// <summary>
    /// Provides a resource to manage [Okta MFA](https://www.vaultproject.io/docs/enterprise/mfa/mfa-okta).
    /// 
    /// **Note** this feature is available only with Vault Enterprise.
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
    ///     var userpass = new Vault.AuthBackend("userpass", new()
    ///     {
    ///         Type = "userpass",
    ///         Path = "userpass",
    ///     });
    /// 
    ///     var myOkta = new Vault.MfaOkta("myOkta", new()
    ///     {
    ///         MountAccessor = userpass.Accessor,
    ///         UsernameFormat = "user@example.com",
    ///         OrgName = "hashicorp",
    ///         ApiToken = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9",
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ## Import
    /// 
    /// Mounts can be imported using the `path`, e.g.
    /// 
    /// ```sh
    /// $ pulumi import vault:index/mfaOkta:MfaOkta my_okta my_okta
    /// ```
    /// </summary>
    [VaultResourceType("vault:index/mfaOkta:MfaOkta")]
    public partial class MfaOkta : global::Pulumi.CustomResource
    {
        /// <summary>
        /// `(string: &lt;required&gt;)` - Okta API key.
        /// </summary>
        [Output("apiToken")]
        public Output<string> ApiToken { get; private set; } = null!;

        /// <summary>
        /// `(string)` - If set, will be used as the base domain for API requests. Examples are `okta.com`, 
        /// `oktapreview.com`, and `okta-emea.com`.
        /// </summary>
        [Output("baseUrl")]
        public Output<string?> BaseUrl { get; private set; } = null!;

        /// <summary>
        /// `(string: &lt;required&gt;)` - The mount to tie this method to for use in automatic mappings. 
        /// The mapping will use the Name field of Aliases associated with this mount as the username in the mapping.
        /// </summary>
        [Output("mountAccessor")]
        public Output<string> MountAccessor { get; private set; } = null!;

        /// <summary>
        /// `(string: &lt;required&gt;)` – Name of the MFA method.
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
        /// `(string: &lt;required&gt;)` - Name of the organization to be used in the Okta API.
        /// </summary>
        [Output("orgName")]
        public Output<string> OrgName { get; private set; } = null!;

        /// <summary>
        /// `(string: &lt;required&gt;)` - If set to true, the username will only match the 
        /// primary email for the account.
        /// </summary>
        [Output("primaryEmail")]
        public Output<bool?> PrimaryEmail { get; private set; } = null!;

        /// <summary>
        /// `(string)` - A format string for mapping Identity names to MFA method names. 
        /// Values to substitute should be placed in `{{}}`. For example, `"{{alias.name}}@example.com"`.
        /// If blank, the Alias's Name field will be used as-is. Currently-supported mappings:
        /// - alias.name: The name returned by the mount configured via the `mount_accessor` parameter
        /// - entity.name: The name configured for the Entity
        /// - alias.metadata.`&lt;key&gt;`: The value of the Alias's metadata parameter
        /// - entity.metadata.`&lt;key&gt;`: The value of the Entity's metadata parameter
        /// </summary>
        [Output("usernameFormat")]
        public Output<string?> UsernameFormat { get; private set; } = null!;


        /// <summary>
        /// Create a MfaOkta resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public MfaOkta(string name, MfaOktaArgs args, CustomResourceOptions? options = null)
            : base("vault:index/mfaOkta:MfaOkta", name, args ?? new MfaOktaArgs(), MakeResourceOptions(options, ""))
        {
        }

        private MfaOkta(string name, Input<string> id, MfaOktaState? state = null, CustomResourceOptions? options = null)
            : base("vault:index/mfaOkta:MfaOkta", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                AdditionalSecretOutputs =
                {
                    "apiToken",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing MfaOkta resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static MfaOkta Get(string name, Input<string> id, MfaOktaState? state = null, CustomResourceOptions? options = null)
        {
            return new MfaOkta(name, id, state, options);
        }
    }

    public sealed class MfaOktaArgs : global::Pulumi.ResourceArgs
    {
        [Input("apiToken", required: true)]
        private Input<string>? _apiToken;

        /// <summary>
        /// `(string: &lt;required&gt;)` - Okta API key.
        /// </summary>
        public Input<string>? ApiToken
        {
            get => _apiToken;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _apiToken = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// `(string)` - If set, will be used as the base domain for API requests. Examples are `okta.com`, 
        /// `oktapreview.com`, and `okta-emea.com`.
        /// </summary>
        [Input("baseUrl")]
        public Input<string>? BaseUrl { get; set; }

        /// <summary>
        /// `(string: &lt;required&gt;)` - The mount to tie this method to for use in automatic mappings. 
        /// The mapping will use the Name field of Aliases associated with this mount as the username in the mapping.
        /// </summary>
        [Input("mountAccessor", required: true)]
        public Input<string> MountAccessor { get; set; } = null!;

        /// <summary>
        /// `(string: &lt;required&gt;)` – Name of the MFA method.
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
        /// `(string: &lt;required&gt;)` - Name of the organization to be used in the Okta API.
        /// </summary>
        [Input("orgName", required: true)]
        public Input<string> OrgName { get; set; } = null!;

        /// <summary>
        /// `(string: &lt;required&gt;)` - If set to true, the username will only match the 
        /// primary email for the account.
        /// </summary>
        [Input("primaryEmail")]
        public Input<bool>? PrimaryEmail { get; set; }

        /// <summary>
        /// `(string)` - A format string for mapping Identity names to MFA method names. 
        /// Values to substitute should be placed in `{{}}`. For example, `"{{alias.name}}@example.com"`.
        /// If blank, the Alias's Name field will be used as-is. Currently-supported mappings:
        /// - alias.name: The name returned by the mount configured via the `mount_accessor` parameter
        /// - entity.name: The name configured for the Entity
        /// - alias.metadata.`&lt;key&gt;`: The value of the Alias's metadata parameter
        /// - entity.metadata.`&lt;key&gt;`: The value of the Entity's metadata parameter
        /// </summary>
        [Input("usernameFormat")]
        public Input<string>? UsernameFormat { get; set; }

        public MfaOktaArgs()
        {
        }
        public static new MfaOktaArgs Empty => new MfaOktaArgs();
    }

    public sealed class MfaOktaState : global::Pulumi.ResourceArgs
    {
        [Input("apiToken")]
        private Input<string>? _apiToken;

        /// <summary>
        /// `(string: &lt;required&gt;)` - Okta API key.
        /// </summary>
        public Input<string>? ApiToken
        {
            get => _apiToken;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _apiToken = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// `(string)` - If set, will be used as the base domain for API requests. Examples are `okta.com`, 
        /// `oktapreview.com`, and `okta-emea.com`.
        /// </summary>
        [Input("baseUrl")]
        public Input<string>? BaseUrl { get; set; }

        /// <summary>
        /// `(string: &lt;required&gt;)` - The mount to tie this method to for use in automatic mappings. 
        /// The mapping will use the Name field of Aliases associated with this mount as the username in the mapping.
        /// </summary>
        [Input("mountAccessor")]
        public Input<string>? MountAccessor { get; set; }

        /// <summary>
        /// `(string: &lt;required&gt;)` – Name of the MFA method.
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
        /// `(string: &lt;required&gt;)` - Name of the organization to be used in the Okta API.
        /// </summary>
        [Input("orgName")]
        public Input<string>? OrgName { get; set; }

        /// <summary>
        /// `(string: &lt;required&gt;)` - If set to true, the username will only match the 
        /// primary email for the account.
        /// </summary>
        [Input("primaryEmail")]
        public Input<bool>? PrimaryEmail { get; set; }

        /// <summary>
        /// `(string)` - A format string for mapping Identity names to MFA method names. 
        /// Values to substitute should be placed in `{{}}`. For example, `"{{alias.name}}@example.com"`.
        /// If blank, the Alias's Name field will be used as-is. Currently-supported mappings:
        /// - alias.name: The name returned by the mount configured via the `mount_accessor` parameter
        /// - entity.name: The name configured for the Entity
        /// - alias.metadata.`&lt;key&gt;`: The value of the Alias's metadata parameter
        /// - entity.metadata.`&lt;key&gt;`: The value of the Entity's metadata parameter
        /// </summary>
        [Input("usernameFormat")]
        public Input<string>? UsernameFormat { get; set; }

        public MfaOktaState()
        {
        }
        public static new MfaOktaState Empty => new MfaOktaState();
    }
}
