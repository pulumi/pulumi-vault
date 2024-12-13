// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vault.kv
{
    public static class GetSecretsList
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
        ///     var awsSecret = new Vault.Kv.Secret("aws_secret", new()
        ///     {
        ///         Path = kvv1.Path.Apply(path =&gt; $"{path}/aws-secret"),
        ///         DataJson = JsonSerializer.Serialize(new Dictionary&lt;string, object?&gt;
        ///         {
        ///             ["zip"] = "zap",
        ///         }),
        ///     });
        /// 
        ///     var azureSecret = new Vault.Kv.Secret("azure_secret", new()
        ///     {
        ///         Path = kvv1.Path.Apply(path =&gt; $"{path}/azure-secret"),
        ///         DataJson = JsonSerializer.Serialize(new Dictionary&lt;string, object?&gt;
        ///         {
        ///             ["foo"] = "bar",
        ///         }),
        ///     });
        /// 
        ///     var secrets = Vault.kv.GetSecretsList.Invoke(new()
        ///     {
        ///         Path = kvv1.Path,
        ///     });
        /// 
        /// });
        /// ```
        /// 
        /// ## Required Vault Capabilities
        /// 
        /// Use of this resource requires the `read` capability on the given path.
        /// </summary>
        public static Task<GetSecretsListResult> InvokeAsync(GetSecretsListArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetSecretsListResult>("vault:kv/getSecretsList:getSecretsList", args ?? new GetSecretsListArgs(), options.WithDefaults());

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
        ///     var awsSecret = new Vault.Kv.Secret("aws_secret", new()
        ///     {
        ///         Path = kvv1.Path.Apply(path =&gt; $"{path}/aws-secret"),
        ///         DataJson = JsonSerializer.Serialize(new Dictionary&lt;string, object?&gt;
        ///         {
        ///             ["zip"] = "zap",
        ///         }),
        ///     });
        /// 
        ///     var azureSecret = new Vault.Kv.Secret("azure_secret", new()
        ///     {
        ///         Path = kvv1.Path.Apply(path =&gt; $"{path}/azure-secret"),
        ///         DataJson = JsonSerializer.Serialize(new Dictionary&lt;string, object?&gt;
        ///         {
        ///             ["foo"] = "bar",
        ///         }),
        ///     });
        /// 
        ///     var secrets = Vault.kv.GetSecretsList.Invoke(new()
        ///     {
        ///         Path = kvv1.Path,
        ///     });
        /// 
        /// });
        /// ```
        /// 
        /// ## Required Vault Capabilities
        /// 
        /// Use of this resource requires the `read` capability on the given path.
        /// </summary>
        public static Output<GetSecretsListResult> Invoke(GetSecretsListInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetSecretsListResult>("vault:kv/getSecretsList:getSecretsList", args ?? new GetSecretsListInvokeArgs(), options.WithDefaults());

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
        ///     var awsSecret = new Vault.Kv.Secret("aws_secret", new()
        ///     {
        ///         Path = kvv1.Path.Apply(path =&gt; $"{path}/aws-secret"),
        ///         DataJson = JsonSerializer.Serialize(new Dictionary&lt;string, object?&gt;
        ///         {
        ///             ["zip"] = "zap",
        ///         }),
        ///     });
        /// 
        ///     var azureSecret = new Vault.Kv.Secret("azure_secret", new()
        ///     {
        ///         Path = kvv1.Path.Apply(path =&gt; $"{path}/azure-secret"),
        ///         DataJson = JsonSerializer.Serialize(new Dictionary&lt;string, object?&gt;
        ///         {
        ///             ["foo"] = "bar",
        ///         }),
        ///     });
        /// 
        ///     var secrets = Vault.kv.GetSecretsList.Invoke(new()
        ///     {
        ///         Path = kvv1.Path,
        ///     });
        /// 
        /// });
        /// ```
        /// 
        /// ## Required Vault Capabilities
        /// 
        /// Use of this resource requires the `read` capability on the given path.
        /// </summary>
        public static Output<GetSecretsListResult> Invoke(GetSecretsListInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetSecretsListResult>("vault:kv/getSecretsList:getSecretsList", args ?? new GetSecretsListInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetSecretsListArgs : global::Pulumi.InvokeArgs
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
        /// Full KV-V1 path where secrets will be listed.
        /// </summary>
        [Input("path", required: true)]
        public string Path { get; set; } = null!;

        public GetSecretsListArgs()
        {
        }
        public static new GetSecretsListArgs Empty => new GetSecretsListArgs();
    }

    public sealed class GetSecretsListInvokeArgs : global::Pulumi.InvokeArgs
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
        /// Full KV-V1 path where secrets will be listed.
        /// </summary>
        [Input("path", required: true)]
        public Input<string> Path { get; set; } = null!;

        public GetSecretsListInvokeArgs()
        {
        }
        public static new GetSecretsListInvokeArgs Empty => new GetSecretsListInvokeArgs();
    }


    [OutputType]
    public sealed class GetSecretsListResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// List of all secret names listed under the given path.
        /// </summary>
        public readonly ImmutableArray<string> Names;
        public readonly string? Namespace;
        public readonly string Path;

        [OutputConstructor]
        private GetSecretsListResult(
            string id,

            ImmutableArray<string> names,

            string? @namespace,

            string path)
        {
            Id = id;
            Names = names;
            Namespace = @namespace;
            Path = path;
        }
    }
}
