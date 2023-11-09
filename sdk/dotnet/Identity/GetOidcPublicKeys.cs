// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vault.Identity
{
    public static class GetOidcPublicKeys
    {
        /// <summary>
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
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
        ///         AllowedClientIds = new[]
        ///         {
        ///             vault_identity_oidc_client.Test.Client_id,
        ///         },
        ///     });
        /// 
        ///     var publicKeys = Vault.Identity.GetOidcPublicKeys.Invoke(new()
        ///     {
        ///         Name = provider.Name,
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetOidcPublicKeysResult> InvokeAsync(GetOidcPublicKeysArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetOidcPublicKeysResult>("vault:identity/getOidcPublicKeys:getOidcPublicKeys", args ?? new GetOidcPublicKeysArgs(), options.WithDefaults());

        /// <summary>
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
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
        ///         AllowedClientIds = new[]
        ///         {
        ///             vault_identity_oidc_client.Test.Client_id,
        ///         },
        ///     });
        /// 
        ///     var publicKeys = Vault.Identity.GetOidcPublicKeys.Invoke(new()
        ///     {
        ///         Name = provider.Name,
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetOidcPublicKeysResult> Invoke(GetOidcPublicKeysInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetOidcPublicKeysResult>("vault:identity/getOidcPublicKeys:getOidcPublicKeys", args ?? new GetOidcPublicKeysInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetOidcPublicKeysArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the OIDC Provider in Vault.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// The namespace of the target resource.
        /// The value should not contain leading or trailing forward slashes.
        /// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
        /// *Available only for Vault Enterprise*.
        /// </summary>
        [Input("namespace")]
        public string? Namespace { get; set; }

        public GetOidcPublicKeysArgs()
        {
        }
        public static new GetOidcPublicKeysArgs Empty => new GetOidcPublicKeysArgs();
    }

    public sealed class GetOidcPublicKeysInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the OIDC Provider in Vault.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// The namespace of the target resource.
        /// The value should not contain leading or trailing forward slashes.
        /// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
        /// *Available only for Vault Enterprise*.
        /// </summary>
        [Input("namespace")]
        public Input<string>? Namespace { get; set; }

        public GetOidcPublicKeysInvokeArgs()
        {
        }
        public static new GetOidcPublicKeysInvokeArgs Empty => new GetOidcPublicKeysInvokeArgs();
    }


    [OutputType]
    public sealed class GetOidcPublicKeysResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The public portion of keys for an OIDC provider. 
        /// Clients can use them to validate the authenticity of an identity token.
        /// </summary>
        public readonly ImmutableArray<ImmutableDictionary<string, object>> Keys;
        public readonly string Name;
        public readonly string? Namespace;

        [OutputConstructor]
        private GetOidcPublicKeysResult(
            string id,

            ImmutableArray<ImmutableDictionary<string, object>> keys,

            string name,

            string? @namespace)
        {
            Id = id;
            Keys = keys;
            Name = name;
            Namespace = @namespace;
        }
    }
}
