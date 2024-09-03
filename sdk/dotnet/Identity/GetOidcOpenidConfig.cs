// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vault.Identity
{
    public static class GetOidcOpenidConfig
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
        ///         Name = "key",
        ///         AllowedClientIds = new[]
        ///         {
        ///             "*",
        ///         },
        ///         RotationPeriod = 3600,
        ///         VerificationTtl = 3600,
        ///     });
        /// 
        ///     var app = new Vault.Identity.OidcClient("app", new()
        ///     {
        ///         Name = "application",
        ///         Key = key.Name,
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
        ///     var provider = new Vault.Identity.OidcProvider("provider", new()
        ///     {
        ///         Name = "provider",
        ///         AllowedClientIds = new[]
        ///         {
        ///             test.ClientId,
        ///         },
        ///     });
        /// 
        ///     var config = Vault.Identity.GetOidcOpenidConfig.Invoke(new()
        ///     {
        ///         Name = provider.Name,
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Task<GetOidcOpenidConfigResult> InvokeAsync(GetOidcOpenidConfigArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetOidcOpenidConfigResult>("vault:identity/getOidcOpenidConfig:getOidcOpenidConfig", args ?? new GetOidcOpenidConfigArgs(), options.WithDefaults());

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
        ///         Name = "key",
        ///         AllowedClientIds = new[]
        ///         {
        ///             "*",
        ///         },
        ///         RotationPeriod = 3600,
        ///         VerificationTtl = 3600,
        ///     });
        /// 
        ///     var app = new Vault.Identity.OidcClient("app", new()
        ///     {
        ///         Name = "application",
        ///         Key = key.Name,
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
        ///     var provider = new Vault.Identity.OidcProvider("provider", new()
        ///     {
        ///         Name = "provider",
        ///         AllowedClientIds = new[]
        ///         {
        ///             test.ClientId,
        ///         },
        ///     });
        /// 
        ///     var config = Vault.Identity.GetOidcOpenidConfig.Invoke(new()
        ///     {
        ///         Name = provider.Name,
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetOidcOpenidConfigResult> Invoke(GetOidcOpenidConfigInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetOidcOpenidConfigResult>("vault:identity/getOidcOpenidConfig:getOidcOpenidConfig", args ?? new GetOidcOpenidConfigInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetOidcOpenidConfigArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the OIDC Provider in Vault.
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

        public GetOidcOpenidConfigArgs()
        {
        }
        public static new GetOidcOpenidConfigArgs Empty => new GetOidcOpenidConfigArgs();
    }

    public sealed class GetOidcOpenidConfigInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the OIDC Provider in Vault.
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

        public GetOidcOpenidConfigInvokeArgs()
        {
        }
        public static new GetOidcOpenidConfigInvokeArgs Empty => new GetOidcOpenidConfigInvokeArgs();
    }


    [OutputType]
    public sealed class GetOidcOpenidConfigResult
    {
        /// <summary>
        /// The Authorization Endpoint for the provider.
        /// </summary>
        public readonly string AuthorizationEndpoint;
        /// <summary>
        /// The grant types supported by the provider.
        /// </summary>
        public readonly ImmutableArray<string> GrantTypesSupporteds;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The signing algorithms supported by
        /// the provider.
        /// </summary>
        public readonly ImmutableArray<string> IdTokenSigningAlgValuesSupporteds;
        /// <summary>
        /// The URL of the issuer for the provider.
        /// </summary>
        public readonly string Issuer;
        /// <summary>
        /// The well known keys URI for the provider.
        /// </summary>
        public readonly string JwksUri;
        public readonly string Name;
        public readonly string? Namespace;
        /// <summary>
        /// Specifies whether Request URI Parameter is
        /// supported by the provider.
        /// </summary>
        public readonly bool RequestUriParameterSupported;
        /// <summary>
        /// The response types supported by the provider.
        /// </summary>
        public readonly ImmutableArray<string> ResponseTypesSupporteds;
        /// <summary>
        /// The scopes supported by the provider.
        /// </summary>
        public readonly ImmutableArray<string> ScopesSupporteds;
        /// <summary>
        /// The subject types supported by the provider.
        /// </summary>
        public readonly ImmutableArray<string> SubjectTypesSupporteds;
        /// <summary>
        /// The Token Endpoint for the provider.
        /// </summary>
        public readonly string TokenEndpoint;
        /// <summary>
        /// The token endpoint auth methods supported by the provider.
        /// </summary>
        public readonly ImmutableArray<string> TokenEndpointAuthMethodsSupporteds;
        /// <summary>
        /// The User Info Endpoint for the provider
        /// </summary>
        public readonly string UserinfoEndpoint;

        [OutputConstructor]
        private GetOidcOpenidConfigResult(
            string authorizationEndpoint,

            ImmutableArray<string> grantTypesSupporteds,

            string id,

            ImmutableArray<string> idTokenSigningAlgValuesSupporteds,

            string issuer,

            string jwksUri,

            string name,

            string? @namespace,

            bool requestUriParameterSupported,

            ImmutableArray<string> responseTypesSupporteds,

            ImmutableArray<string> scopesSupporteds,

            ImmutableArray<string> subjectTypesSupporteds,

            string tokenEndpoint,

            ImmutableArray<string> tokenEndpointAuthMethodsSupporteds,

            string userinfoEndpoint)
        {
            AuthorizationEndpoint = authorizationEndpoint;
            GrantTypesSupporteds = grantTypesSupporteds;
            Id = id;
            IdTokenSigningAlgValuesSupporteds = idTokenSigningAlgValuesSupporteds;
            Issuer = issuer;
            JwksUri = jwksUri;
            Name = name;
            Namespace = @namespace;
            RequestUriParameterSupported = requestUriParameterSupported;
            ResponseTypesSupporteds = responseTypesSupporteds;
            ScopesSupporteds = scopesSupporteds;
            SubjectTypesSupporteds = subjectTypesSupporteds;
            TokenEndpoint = tokenEndpoint;
            TokenEndpointAuthMethodsSupporteds = tokenEndpointAuthMethodsSupporteds;
            UserinfoEndpoint = userinfoEndpoint;
        }
    }
}
