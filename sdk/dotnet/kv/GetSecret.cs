// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vault.kv
{
    public static class GetSecret
    {
        /// <summary>
        /// ## Example Usage
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using System.Text.Json;
        /// using Pulumi;
        /// using Vault = Pulumi.Vault;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var kvv1 = new Vault.Mount("kvv1", new()
        ///     {
        ///         Path = "kvv1",
        ///         Type = "kv",
        ///         Options = 
        ///         {
        ///             { "version", "1" },
        ///         },
        ///         Description = "KV Version 1 secret engine mount",
        ///     });
        /// 
        ///     var secret = new Vault.Kv.Secret("secret", new()
        ///     {
        ///         Path = kvv1.Path.Apply(path =&gt; $"{path}/secret"),
        ///         DataJson = JsonSerializer.Serialize(new Dictionary&lt;string, object?&gt;
        ///         {
        ///             ["zip"] = "zap",
        ///             ["foo"] = "bar",
        ///         }),
        ///     });
        /// 
        ///     var secretData = Vault.kv.GetSecret.Invoke(new()
        ///     {
        ///         Path = secret.Path,
        ///     });
        /// 
        /// });
        /// ```
        /// 
        /// ## Required Vault Capabilities
        /// 
        /// Use of this resource requires the `read` capability on the given path.
        /// </summary>
        public static Task<GetSecretResult> InvokeAsync(GetSecretArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetSecretResult>("vault:kv/getSecret:getSecret", args ?? new GetSecretArgs(), options.WithDefaults());

        /// <summary>
        /// ## Example Usage
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using System.Text.Json;
        /// using Pulumi;
        /// using Vault = Pulumi.Vault;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var kvv1 = new Vault.Mount("kvv1", new()
        ///     {
        ///         Path = "kvv1",
        ///         Type = "kv",
        ///         Options = 
        ///         {
        ///             { "version", "1" },
        ///         },
        ///         Description = "KV Version 1 secret engine mount",
        ///     });
        /// 
        ///     var secret = new Vault.Kv.Secret("secret", new()
        ///     {
        ///         Path = kvv1.Path.Apply(path =&gt; $"{path}/secret"),
        ///         DataJson = JsonSerializer.Serialize(new Dictionary&lt;string, object?&gt;
        ///         {
        ///             ["zip"] = "zap",
        ///             ["foo"] = "bar",
        ///         }),
        ///     });
        /// 
        ///     var secretData = Vault.kv.GetSecret.Invoke(new()
        ///     {
        ///         Path = secret.Path,
        ///     });
        /// 
        /// });
        /// ```
        /// 
        /// ## Required Vault Capabilities
        /// 
        /// Use of this resource requires the `read` capability on the given path.
        /// </summary>
        public static Output<GetSecretResult> Invoke(GetSecretInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetSecretResult>("vault:kv/getSecret:getSecret", args ?? new GetSecretInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// ## Example Usage
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using System.Text.Json;
        /// using Pulumi;
        /// using Vault = Pulumi.Vault;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var kvv1 = new Vault.Mount("kvv1", new()
        ///     {
        ///         Path = "kvv1",
        ///         Type = "kv",
        ///         Options = 
        ///         {
        ///             { "version", "1" },
        ///         },
        ///         Description = "KV Version 1 secret engine mount",
        ///     });
        /// 
        ///     var secret = new Vault.Kv.Secret("secret", new()
        ///     {
        ///         Path = kvv1.Path.Apply(path =&gt; $"{path}/secret"),
        ///         DataJson = JsonSerializer.Serialize(new Dictionary&lt;string, object?&gt;
        ///         {
        ///             ["zip"] = "zap",
        ///             ["foo"] = "bar",
        ///         }),
        ///     });
        /// 
        ///     var secretData = Vault.kv.GetSecret.Invoke(new()
        ///     {
        ///         Path = secret.Path,
        ///     });
        /// 
        /// });
        /// ```
        /// 
        /// ## Required Vault Capabilities
        /// 
        /// Use of this resource requires the `read` capability on the given path.
        /// </summary>
        public static Output<GetSecretResult> Invoke(GetSecretInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetSecretResult>("vault:kv/getSecret:getSecret", args ?? new GetSecretInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetSecretArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The namespace of the target resource.
        /// The value should not contain leading or trailing forward slashes.
        /// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
        /// *Available only for Vault Enterprise*.
        /// </summary>
        [Input("namespace")]
        public string? Namespace { get; set; }

        /// <summary>
        /// Full path of the KV-V1 secret.
        /// </summary>
        [Input("path", required: true)]
        public string Path { get; set; } = null!;

        public GetSecretArgs()
        {
        }
        public static new GetSecretArgs Empty => new GetSecretArgs();
    }

    public sealed class GetSecretInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The namespace of the target resource.
        /// The value should not contain leading or trailing forward slashes.
        /// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
        /// *Available only for Vault Enterprise*.
        /// </summary>
        [Input("namespace")]
        public Input<string>? Namespace { get; set; }

        /// <summary>
        /// Full path of the KV-V1 secret.
        /// </summary>
        [Input("path", required: true)]
        public Input<string> Path { get; set; } = null!;

        public GetSecretInvokeArgs()
        {
        }
        public static new GetSecretInvokeArgs Empty => new GetSecretInvokeArgs();
    }


    [OutputType]
    public sealed class GetSecretResult
    {
        /// <summary>
        /// A mapping whose keys are the top-level data keys returned from
        /// Vault and whose values are the corresponding values. This map can only
        /// represent string data, so any non-string values returned from Vault are
        /// serialized as JSON.
        /// </summary>
        public readonly ImmutableDictionary<string, string> Data;
        /// <summary>
        /// JSON-encoded string that that is
        /// read as the secret data at the given path.
        /// </summary>
        public readonly string DataJson;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The duration of the secret lease, in seconds. Once 
        /// this time has passed any plan generated with this data may fail to apply.
        /// </summary>
        public readonly int LeaseDuration;
        /// <summary>
        /// The lease identifier assigned by Vault, if any.
        /// </summary>
        public readonly string LeaseId;
        /// <summary>
        /// True if the duration of this lease can be extended 
        /// through renewal.
        /// </summary>
        public readonly bool LeaseRenewable;
        public readonly string? Namespace;
        public readonly string Path;

        [OutputConstructor]
        private GetSecretResult(
            ImmutableDictionary<string, string> data,

            string dataJson,

            string id,

            int leaseDuration,

            string leaseId,

            bool leaseRenewable,

            string? @namespace,

            string path)
        {
            Data = data;
            DataJson = dataJson;
            Id = id;
            LeaseDuration = leaseDuration;
            LeaseId = leaseId;
            LeaseRenewable = leaseRenewable;
            Namespace = @namespace;
            Path = path;
        }
    }
}
