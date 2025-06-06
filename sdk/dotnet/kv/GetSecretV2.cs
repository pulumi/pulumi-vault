// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vault.kv
{
    public static class GetSecretV2
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
        ///     var kvv2 = new Vault.Mount("kvv2", new()
        ///     {
        ///         Path = "kvv2",
        ///         Type = "kv",
        ///         Options = 
        ///         {
        ///             { "version", "2" },
        ///         },
        ///         Description = "KV Version 2 secret engine mount",
        ///     });
        /// 
        ///     var exampleSecretV2 = new Vault.Kv.SecretV2("example", new()
        ///     {
        ///         Mount = kvv2.Path,
        ///         Name = "secret",
        ///         DeleteAllVersions = true,
        ///         DataJson = JsonSerializer.Serialize(new Dictionary&lt;string, object?&gt;
        ///         {
        ///             ["zip"] = "zap",
        ///             ["foo"] = "bar",
        ///         }),
        ///     });
        /// 
        ///     var example = Vault.kv.GetSecretV2.Invoke(new()
        ///     {
        ///         Mount = kvv2.Path,
        ///         Name = exampleSecretV2.Name,
        ///     });
        /// 
        /// });
        /// ```
        /// 
        /// ## Required Vault Capabilities
        /// 
        /// Use of this resource requires the `read` capability on the given path.
        /// </summary>
        public static Task<GetSecretV2Result> InvokeAsync(GetSecretV2Args args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetSecretV2Result>("vault:kv/getSecretV2:getSecretV2", args ?? new GetSecretV2Args(), options.WithDefaults());

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
        ///     var kvv2 = new Vault.Mount("kvv2", new()
        ///     {
        ///         Path = "kvv2",
        ///         Type = "kv",
        ///         Options = 
        ///         {
        ///             { "version", "2" },
        ///         },
        ///         Description = "KV Version 2 secret engine mount",
        ///     });
        /// 
        ///     var exampleSecretV2 = new Vault.Kv.SecretV2("example", new()
        ///     {
        ///         Mount = kvv2.Path,
        ///         Name = "secret",
        ///         DeleteAllVersions = true,
        ///         DataJson = JsonSerializer.Serialize(new Dictionary&lt;string, object?&gt;
        ///         {
        ///             ["zip"] = "zap",
        ///             ["foo"] = "bar",
        ///         }),
        ///     });
        /// 
        ///     var example = Vault.kv.GetSecretV2.Invoke(new()
        ///     {
        ///         Mount = kvv2.Path,
        ///         Name = exampleSecretV2.Name,
        ///     });
        /// 
        /// });
        /// ```
        /// 
        /// ## Required Vault Capabilities
        /// 
        /// Use of this resource requires the `read` capability on the given path.
        /// </summary>
        public static Output<GetSecretV2Result> Invoke(GetSecretV2InvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetSecretV2Result>("vault:kv/getSecretV2:getSecretV2", args ?? new GetSecretV2InvokeArgs(), options.WithDefaults());

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
        ///     var kvv2 = new Vault.Mount("kvv2", new()
        ///     {
        ///         Path = "kvv2",
        ///         Type = "kv",
        ///         Options = 
        ///         {
        ///             { "version", "2" },
        ///         },
        ///         Description = "KV Version 2 secret engine mount",
        ///     });
        /// 
        ///     var exampleSecretV2 = new Vault.Kv.SecretV2("example", new()
        ///     {
        ///         Mount = kvv2.Path,
        ///         Name = "secret",
        ///         DeleteAllVersions = true,
        ///         DataJson = JsonSerializer.Serialize(new Dictionary&lt;string, object?&gt;
        ///         {
        ///             ["zip"] = "zap",
        ///             ["foo"] = "bar",
        ///         }),
        ///     });
        /// 
        ///     var example = Vault.kv.GetSecretV2.Invoke(new()
        ///     {
        ///         Mount = kvv2.Path,
        ///         Name = exampleSecretV2.Name,
        ///     });
        /// 
        /// });
        /// ```
        /// 
        /// ## Required Vault Capabilities
        /// 
        /// Use of this resource requires the `read` capability on the given path.
        /// </summary>
        public static Output<GetSecretV2Result> Invoke(GetSecretV2InvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetSecretV2Result>("vault:kv/getSecretV2:getSecretV2", args ?? new GetSecretV2InvokeArgs(), options.WithDefaults());
    }


    public sealed class GetSecretV2Args : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Path where KV-V2 engine is mounted.
        /// </summary>
        [Input("mount", required: true)]
        public string Mount { get; set; } = null!;

        /// <summary>
        /// Full name of the secret. For a nested secret
        /// the name is the nested path excluding the mount and data
        /// prefix. For example, for a secret at `kvv2/data/foo/bar/baz`
        /// the name is `foo/bar/baz`.
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

        /// <summary>
        /// Version of the secret to retrieve.
        /// </summary>
        [Input("version")]
        public int? Version { get; set; }

        public GetSecretV2Args()
        {
        }
        public static new GetSecretV2Args Empty => new GetSecretV2Args();
    }

    public sealed class GetSecretV2InvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Path where KV-V2 engine is mounted.
        /// </summary>
        [Input("mount", required: true)]
        public Input<string> Mount { get; set; } = null!;

        /// <summary>
        /// Full name of the secret. For a nested secret
        /// the name is the nested path excluding the mount and data
        /// prefix. For example, for a secret at `kvv2/data/foo/bar/baz`
        /// the name is `foo/bar/baz`.
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

        /// <summary>
        /// Version of the secret to retrieve.
        /// </summary>
        [Input("version")]
        public Input<int>? Version { get; set; }

        public GetSecretV2InvokeArgs()
        {
        }
        public static new GetSecretV2InvokeArgs Empty => new GetSecretV2InvokeArgs();
    }


    [OutputType]
    public sealed class GetSecretV2Result
    {
        /// <summary>
        /// Time at which secret was created.
        /// </summary>
        public readonly string CreatedTime;
        /// <summary>
        /// Custom metadata for the secret.
        /// </summary>
        public readonly ImmutableDictionary<string, string> CustomMetadata;
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
        /// Deletion time for the secret.
        /// </summary>
        public readonly string DeletionTime;
        /// <summary>
        /// Indicates whether the secret has been destroyed.
        /// </summary>
        public readonly bool Destroyed;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string Mount;
        public readonly string Name;
        public readonly string? Namespace;
        /// <summary>
        /// Full path where the KVV2 secret is written.
        /// </summary>
        public readonly string Path;
        /// <summary>
        /// Version of the secret.
        /// </summary>
        public readonly int? Version;

        [OutputConstructor]
        private GetSecretV2Result(
            string createdTime,

            ImmutableDictionary<string, string> customMetadata,

            ImmutableDictionary<string, string> data,

            string dataJson,

            string deletionTime,

            bool destroyed,

            string id,

            string mount,

            string name,

            string? @namespace,

            string path,

            int? version)
        {
            CreatedTime = createdTime;
            CustomMetadata = customMetadata;
            Data = data;
            DataJson = dataJson;
            DeletionTime = deletionTime;
            Destroyed = destroyed;
            Id = id;
            Mount = mount;
            Name = name;
            Namespace = @namespace;
            Path = path;
            Version = version;
        }
    }
}
