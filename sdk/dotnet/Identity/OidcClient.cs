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
    /// Manages OIDC Clients in a Vault server. See the [Vault documentation](https://www.vaultproject.io/api-docs/secret/identity/oidc-provider#create-or-update-an-assignment)
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
    ///     var test = new Vault.Identity.OidcAssignment("test", new()
    ///     {
    ///         Name = "my-assignment",
    ///         EntityIds = new[]
    ///         {
    ///             "ascbascas-2231a-sdfaa",
    ///         },
    ///         GroupIds = new[]
    ///         {
    ///             "sajkdsad-32414-sfsada",
    ///         },
    ///     });
    /// 
    ///     var testOidcClient = new Vault.Identity.OidcClient("test", new()
    ///     {
    ///         Name = "my-app",
    ///         RedirectUris = new[]
    ///         {
    ///             "http://127.0.0.1:9200/v1/auth-methods/oidc:authenticate:callback",
    ///             "http://127.0.0.1:8251/callback",
    ///             "http://127.0.0.1:8080/callback",
    ///         },
    ///         Assignments = new[]
    ///         {
    ///             test.Name,
    ///         },
    ///         IdTokenTtl = 2400,
    ///         AccessTokenTtl = 7200,
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// OIDC Clients can be imported using the `name`, e.g.
    /// 
    /// ```sh
    /// $ pulumi import vault:identity/oidcClient:OidcClient test my-app
    /// ```
    /// </summary>
    [VaultResourceType("vault:identity/oidcClient:OidcClient")]
    public partial class OidcClient : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The time-to-live for access tokens obtained by the client.
        /// </summary>
        [Output("accessTokenTtl")]
        public Output<int> AccessTokenTtl { get; private set; } = null!;

        /// <summary>
        /// A list of assignment resources associated with the client.
        /// </summary>
        [Output("assignments")]
        public Output<ImmutableArray<string>> Assignments { get; private set; } = null!;

        /// <summary>
        /// The Client ID returned by Vault.
        /// </summary>
        [Output("clientId")]
        public Output<string> ClientId { get; private set; } = null!;

        /// <summary>
        /// The Client Secret Key returned by Vault.
        /// For public OpenID Clients `client_secret` is set to an empty string `""`
        /// </summary>
        [Output("clientSecret")]
        public Output<string> ClientSecret { get; private set; } = null!;

        /// <summary>
        /// The client type based on its ability to maintain confidentiality of credentials.
        /// The following client types are supported: `confidential`, `public`. Defaults to `confidential`.
        /// </summary>
        [Output("clientType")]
        public Output<string> ClientType { get; private set; } = null!;

        /// <summary>
        /// The time-to-live for ID tokens obtained by the client. 
        /// The value should be less than the `verification_ttl` on the key.
        /// </summary>
        [Output("idTokenTtl")]
        public Output<int> IdTokenTtl { get; private set; } = null!;

        /// <summary>
        /// A reference to a named key resource in Vault.
        /// This cannot be modified after creation. If not provided, the `default`
        /// key is used.
        /// </summary>
        [Output("key")]
        public Output<string> Key { get; private set; } = null!;

        /// <summary>
        /// The name of the client.
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
        /// Redirection URI values used by the client. 
        /// One of these values must exactly match the `redirect_uri` parameter value
        /// used in each authentication request.
        /// </summary>
        [Output("redirectUris")]
        public Output<ImmutableArray<string>> RedirectUris { get; private set; } = null!;


        /// <summary>
        /// Create a OidcClient resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public OidcClient(string name, OidcClientArgs? args = null, CustomResourceOptions? options = null)
            : base("vault:identity/oidcClient:OidcClient", name, args ?? new OidcClientArgs(), MakeResourceOptions(options, ""))
        {
        }

        private OidcClient(string name, Input<string> id, OidcClientState? state = null, CustomResourceOptions? options = null)
            : base("vault:identity/oidcClient:OidcClient", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                AdditionalSecretOutputs =
                {
                    "clientSecret",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing OidcClient resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static OidcClient Get(string name, Input<string> id, OidcClientState? state = null, CustomResourceOptions? options = null)
        {
            return new OidcClient(name, id, state, options);
        }
    }

    public sealed class OidcClientArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The time-to-live for access tokens obtained by the client.
        /// </summary>
        [Input("accessTokenTtl")]
        public Input<int>? AccessTokenTtl { get; set; }

        [Input("assignments")]
        private InputList<string>? _assignments;

        /// <summary>
        /// A list of assignment resources associated with the client.
        /// </summary>
        public InputList<string> Assignments
        {
            get => _assignments ?? (_assignments = new InputList<string>());
            set => _assignments = value;
        }

        /// <summary>
        /// The client type based on its ability to maintain confidentiality of credentials.
        /// The following client types are supported: `confidential`, `public`. Defaults to `confidential`.
        /// </summary>
        [Input("clientType")]
        public Input<string>? ClientType { get; set; }

        /// <summary>
        /// The time-to-live for ID tokens obtained by the client. 
        /// The value should be less than the `verification_ttl` on the key.
        /// </summary>
        [Input("idTokenTtl")]
        public Input<int>? IdTokenTtl { get; set; }

        /// <summary>
        /// A reference to a named key resource in Vault.
        /// This cannot be modified after creation. If not provided, the `default`
        /// key is used.
        /// </summary>
        [Input("key")]
        public Input<string>? Key { get; set; }

        /// <summary>
        /// The name of the client.
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

        [Input("redirectUris")]
        private InputList<string>? _redirectUris;

        /// <summary>
        /// Redirection URI values used by the client. 
        /// One of these values must exactly match the `redirect_uri` parameter value
        /// used in each authentication request.
        /// </summary>
        public InputList<string> RedirectUris
        {
            get => _redirectUris ?? (_redirectUris = new InputList<string>());
            set => _redirectUris = value;
        }

        public OidcClientArgs()
        {
        }
        public static new OidcClientArgs Empty => new OidcClientArgs();
    }

    public sealed class OidcClientState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The time-to-live for access tokens obtained by the client.
        /// </summary>
        [Input("accessTokenTtl")]
        public Input<int>? AccessTokenTtl { get; set; }

        [Input("assignments")]
        private InputList<string>? _assignments;

        /// <summary>
        /// A list of assignment resources associated with the client.
        /// </summary>
        public InputList<string> Assignments
        {
            get => _assignments ?? (_assignments = new InputList<string>());
            set => _assignments = value;
        }

        /// <summary>
        /// The Client ID returned by Vault.
        /// </summary>
        [Input("clientId")]
        public Input<string>? ClientId { get; set; }

        [Input("clientSecret")]
        private Input<string>? _clientSecret;

        /// <summary>
        /// The Client Secret Key returned by Vault.
        /// For public OpenID Clients `client_secret` is set to an empty string `""`
        /// </summary>
        public Input<string>? ClientSecret
        {
            get => _clientSecret;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _clientSecret = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// The client type based on its ability to maintain confidentiality of credentials.
        /// The following client types are supported: `confidential`, `public`. Defaults to `confidential`.
        /// </summary>
        [Input("clientType")]
        public Input<string>? ClientType { get; set; }

        /// <summary>
        /// The time-to-live for ID tokens obtained by the client. 
        /// The value should be less than the `verification_ttl` on the key.
        /// </summary>
        [Input("idTokenTtl")]
        public Input<int>? IdTokenTtl { get; set; }

        /// <summary>
        /// A reference to a named key resource in Vault.
        /// This cannot be modified after creation. If not provided, the `default`
        /// key is used.
        /// </summary>
        [Input("key")]
        public Input<string>? Key { get; set; }

        /// <summary>
        /// The name of the client.
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

        [Input("redirectUris")]
        private InputList<string>? _redirectUris;

        /// <summary>
        /// Redirection URI values used by the client. 
        /// One of these values must exactly match the `redirect_uri` parameter value
        /// used in each authentication request.
        /// </summary>
        public InputList<string> RedirectUris
        {
            get => _redirectUris ?? (_redirectUris = new InputList<string>());
            set => _redirectUris = value;
        }

        public OidcClientState()
        {
        }
        public static new OidcClientState Empty => new OidcClientState();
    }
}
