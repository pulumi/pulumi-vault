// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vault.Identity
{
    /// <summary>
    /// Resource for configuring the okta MFA method.
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
    ///     var example = new Vault.Identity.MfaOkta("example", new()
    ///     {
    ///         OrgName = "org1",
    ///         ApiToken = "token1",
    ///         BaseUrl = "qux.baz.com",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Resource can be imported using its `uuid` field, e.g.
    /// 
    /// ```sh
    /// $ pulumi import vault:identity/mfaOkta:MfaOkta example 0d89c36a-4ff5-4d70-8749-bb6a5598aeec
    /// ```
    /// </summary>
    [VaultResourceType("vault:identity/mfaOkta:MfaOkta")]
    public partial class MfaOkta : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Okta API token.
        /// </summary>
        [Output("apiToken")]
        public Output<string> ApiToken { get; private set; } = null!;

        /// <summary>
        /// The base domain to use for API requests.
        /// </summary>
        [Output("baseUrl")]
        public Output<string?> BaseUrl { get; private set; } = null!;

        /// <summary>
        /// Method ID.
        /// </summary>
        [Output("methodId")]
        public Output<string> MethodId { get; private set; } = null!;

        /// <summary>
        /// Mount accessor.
        /// </summary>
        [Output("mountAccessor")]
        public Output<string> MountAccessor { get; private set; } = null!;

        /// <summary>
        /// Method name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Target namespace. (requires Enterprise)
        /// </summary>
        [Output("namespace")]
        public Output<string?> Namespace { get; private set; } = null!;

        /// <summary>
        /// Method's namespace ID.
        /// </summary>
        [Output("namespaceId")]
        public Output<string> NamespaceId { get; private set; } = null!;

        /// <summary>
        /// Method's namespace path.
        /// </summary>
        [Output("namespacePath")]
        public Output<string> NamespacePath { get; private set; } = null!;

        /// <summary>
        /// Name of the organization to be used in the Okta API.
        /// </summary>
        [Output("orgName")]
        public Output<string> OrgName { get; private set; } = null!;

        /// <summary>
        /// Only match the primary email for the account.
        /// </summary>
        [Output("primaryEmail")]
        public Output<bool?> PrimaryEmail { get; private set; } = null!;

        /// <summary>
        /// MFA type.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;

        /// <summary>
        /// A template string for mapping Identity names to MFA methods.
        /// </summary>
        [Output("usernameFormat")]
        public Output<string?> UsernameFormat { get; private set; } = null!;

        /// <summary>
        /// Resource UUID.
        /// </summary>
        [Output("uuid")]
        public Output<string> Uuid { get; private set; } = null!;


        /// <summary>
        /// Create a MfaOkta resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public MfaOkta(string name, MfaOktaArgs args, CustomResourceOptions? options = null)
            : base("vault:identity/mfaOkta:MfaOkta", name, args ?? new MfaOktaArgs(), MakeResourceOptions(options, ""))
        {
        }

        private MfaOkta(string name, Input<string> id, MfaOktaState? state = null, CustomResourceOptions? options = null)
            : base("vault:identity/mfaOkta:MfaOkta", name, state, MakeResourceOptions(options, id))
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
        /// Okta API token.
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
        /// The base domain to use for API requests.
        /// </summary>
        [Input("baseUrl")]
        public Input<string>? BaseUrl { get; set; }

        /// <summary>
        /// Target namespace. (requires Enterprise)
        /// </summary>
        [Input("namespace")]
        public Input<string>? Namespace { get; set; }

        /// <summary>
        /// Name of the organization to be used in the Okta API.
        /// </summary>
        [Input("orgName", required: true)]
        public Input<string> OrgName { get; set; } = null!;

        /// <summary>
        /// Only match the primary email for the account.
        /// </summary>
        [Input("primaryEmail")]
        public Input<bool>? PrimaryEmail { get; set; }

        /// <summary>
        /// A template string for mapping Identity names to MFA methods.
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
        /// Okta API token.
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
        /// The base domain to use for API requests.
        /// </summary>
        [Input("baseUrl")]
        public Input<string>? BaseUrl { get; set; }

        /// <summary>
        /// Method ID.
        /// </summary>
        [Input("methodId")]
        public Input<string>? MethodId { get; set; }

        /// <summary>
        /// Mount accessor.
        /// </summary>
        [Input("mountAccessor")]
        public Input<string>? MountAccessor { get; set; }

        /// <summary>
        /// Method name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Target namespace. (requires Enterprise)
        /// </summary>
        [Input("namespace")]
        public Input<string>? Namespace { get; set; }

        /// <summary>
        /// Method's namespace ID.
        /// </summary>
        [Input("namespaceId")]
        public Input<string>? NamespaceId { get; set; }

        /// <summary>
        /// Method's namespace path.
        /// </summary>
        [Input("namespacePath")]
        public Input<string>? NamespacePath { get; set; }

        /// <summary>
        /// Name of the organization to be used in the Okta API.
        /// </summary>
        [Input("orgName")]
        public Input<string>? OrgName { get; set; }

        /// <summary>
        /// Only match the primary email for the account.
        /// </summary>
        [Input("primaryEmail")]
        public Input<bool>? PrimaryEmail { get; set; }

        /// <summary>
        /// MFA type.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        /// <summary>
        /// A template string for mapping Identity names to MFA methods.
        /// </summary>
        [Input("usernameFormat")]
        public Input<string>? UsernameFormat { get; set; }

        /// <summary>
        /// Resource UUID.
        /// </summary>
        [Input("uuid")]
        public Input<string>? Uuid { get; set; }

        public MfaOktaState()
        {
        }
        public static new MfaOktaState Empty => new MfaOktaState();
    }
}
