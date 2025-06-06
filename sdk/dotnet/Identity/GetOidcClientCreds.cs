// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vault.Identity
{
    public static class GetOidcClientCreds
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
        ///     var app = new Vault.Identity.OidcClient("app", new()
        ///     {
        ///         Name = "application",
        ///         RedirectUris = new[]
        ///         {
        ///             "http://127.0.0.1:9200/v1/auth-methods/oidc:authenticate:callback",
        ///             "http://127.0.0.1:8251/callback",
        ///             "http://127.0.0.1:8080/callback",
        ///         },
        ///         IdTokenTtl = 2400,
        ///         AccessTokenTtl = 7200,
        ///     });
        /// 
        ///     var creds = Vault.Identity.GetOidcClientCreds.Invoke(new()
        ///     {
        ///         Name = app.Name,
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Task<GetOidcClientCredsResult> InvokeAsync(GetOidcClientCredsArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetOidcClientCredsResult>("vault:identity/getOidcClientCreds:getOidcClientCreds", args ?? new GetOidcClientCredsArgs(), options.WithDefaults());

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
        ///     var app = new Vault.Identity.OidcClient("app", new()
        ///     {
        ///         Name = "application",
        ///         RedirectUris = new[]
        ///         {
        ///             "http://127.0.0.1:9200/v1/auth-methods/oidc:authenticate:callback",
        ///             "http://127.0.0.1:8251/callback",
        ///             "http://127.0.0.1:8080/callback",
        ///         },
        ///         IdTokenTtl = 2400,
        ///         AccessTokenTtl = 7200,
        ///     });
        /// 
        ///     var creds = Vault.Identity.GetOidcClientCreds.Invoke(new()
        ///     {
        ///         Name = app.Name,
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetOidcClientCredsResult> Invoke(GetOidcClientCredsInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetOidcClientCredsResult>("vault:identity/getOidcClientCreds:getOidcClientCreds", args ?? new GetOidcClientCredsInvokeArgs(), options.WithDefaults());

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
        ///     var app = new Vault.Identity.OidcClient("app", new()
        ///     {
        ///         Name = "application",
        ///         RedirectUris = new[]
        ///         {
        ///             "http://127.0.0.1:9200/v1/auth-methods/oidc:authenticate:callback",
        ///             "http://127.0.0.1:8251/callback",
        ///             "http://127.0.0.1:8080/callback",
        ///         },
        ///         IdTokenTtl = 2400,
        ///         AccessTokenTtl = 7200,
        ///     });
        /// 
        ///     var creds = Vault.Identity.GetOidcClientCreds.Invoke(new()
        ///     {
        ///         Name = app.Name,
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetOidcClientCredsResult> Invoke(GetOidcClientCredsInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetOidcClientCredsResult>("vault:identity/getOidcClientCreds:getOidcClientCreds", args ?? new GetOidcClientCredsInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetOidcClientCredsArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the OIDC Client in Vault.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// The namespace of the target resource.
        /// The value should not contain leading or trailing forward slashes.
        /// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
        /// *Available only for Vault Enterprise*.
        /// </summary>
        [Input("namespace")]
        public string? Namespace { get; set; }

        public GetOidcClientCredsArgs()
        {
        }
        public static new GetOidcClientCredsArgs Empty => new GetOidcClientCredsArgs();
    }

    public sealed class GetOidcClientCredsInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the OIDC Client in Vault.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// The namespace of the target resource.
        /// The value should not contain leading or trailing forward slashes.
        /// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
        /// *Available only for Vault Enterprise*.
        /// </summary>
        [Input("namespace")]
        public Input<string>? Namespace { get; set; }

        public GetOidcClientCredsInvokeArgs()
        {
        }
        public static new GetOidcClientCredsInvokeArgs Empty => new GetOidcClientCredsInvokeArgs();
    }


    [OutputType]
    public sealed class GetOidcClientCredsResult
    {
        /// <summary>
        /// The Client ID returned by Vault.
        /// </summary>
        public readonly string ClientId;
        /// <summary>
        /// The Client Secret Key returned by Vault.
        /// For public OpenID Clients `client_secret` is set to an empty string `""`
        /// </summary>
        public readonly string ClientSecret;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string Name;
        public readonly string? Namespace;

        [OutputConstructor]
        private GetOidcClientCredsResult(
            string clientId,

            string clientSecret,

            string id,

            string name,

            string? @namespace)
        {
            ClientId = clientId;
            ClientSecret = clientSecret;
            Id = id;
            Name = name;
            Namespace = @namespace;
        }
    }
}
