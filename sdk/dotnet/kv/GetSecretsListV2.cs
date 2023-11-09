// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vault.kv
{
    public static class GetSecretsListV2
    {
        /// <summary>
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
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
        ///     var awsSecret = new Vault.Kv.SecretV2("awsSecret", new()
        ///     {
        ///         Mount = kvv2.Path,
        ///         DataJson = JsonSerializer.Serialize(new Dictionary&lt;string, object?&gt;
        ///         {
        ///             ["zip"] = "zap",
        ///         }),
        ///     });
        /// 
        ///     var azureSecret = new Vault.Kv.SecretV2("azureSecret", new()
        ///     {
        ///         Mount = kvv2.Path,
        ///         DataJson = JsonSerializer.Serialize(new Dictionary&lt;string, object?&gt;
        ///         {
        ///             ["foo"] = "bar",
        ///         }),
        ///     });
        /// 
        ///     var nestedSecret = new Vault.Kv.SecretV2("nestedSecret", new()
        ///     {
        ///         Mount = kvv2.Path,
        ///         DataJson = JsonSerializer.Serialize(new Dictionary&lt;string, object?&gt;
        ///         {
        ///             ["password"] = "test",
        ///         }),
        ///     });
        /// 
        ///     var secrets = Vault.kv.GetSecretsListV2.Invoke(new()
        ///     {
        ///         Mount = kvv2.Path,
        ///     });
        /// 
        ///     var nestedSecrets = Vault.kv.GetSecretsListV2.Invoke(new()
        ///     {
        ///         Mount = kvv2.Path,
        ///         Name = vault_kv_secret_v2.Test_2.Name,
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// ## Required Vault Capabilities
        /// 
        /// Use of this resource requires the `read` capability on the given path.
        /// </summary>
        public static Task<GetSecretsListV2Result> InvokeAsync(GetSecretsListV2Args args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetSecretsListV2Result>("vault:kv/getSecretsListV2:getSecretsListV2", args ?? new GetSecretsListV2Args(), options.WithDefaults());

        /// <summary>
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
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
        ///     var awsSecret = new Vault.Kv.SecretV2("awsSecret", new()
        ///     {
        ///         Mount = kvv2.Path,
        ///         DataJson = JsonSerializer.Serialize(new Dictionary&lt;string, object?&gt;
        ///         {
        ///             ["zip"] = "zap",
        ///         }),
        ///     });
        /// 
        ///     var azureSecret = new Vault.Kv.SecretV2("azureSecret", new()
        ///     {
        ///         Mount = kvv2.Path,
        ///         DataJson = JsonSerializer.Serialize(new Dictionary&lt;string, object?&gt;
        ///         {
        ///             ["foo"] = "bar",
        ///         }),
        ///     });
        /// 
        ///     var nestedSecret = new Vault.Kv.SecretV2("nestedSecret", new()
        ///     {
        ///         Mount = kvv2.Path,
        ///         DataJson = JsonSerializer.Serialize(new Dictionary&lt;string, object?&gt;
        ///         {
        ///             ["password"] = "test",
        ///         }),
        ///     });
        /// 
        ///     var secrets = Vault.kv.GetSecretsListV2.Invoke(new()
        ///     {
        ///         Mount = kvv2.Path,
        ///     });
        /// 
        ///     var nestedSecrets = Vault.kv.GetSecretsListV2.Invoke(new()
        ///     {
        ///         Mount = kvv2.Path,
        ///         Name = vault_kv_secret_v2.Test_2.Name,
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// ## Required Vault Capabilities
        /// 
        /// Use of this resource requires the `read` capability on the given path.
        /// </summary>
        public static Output<GetSecretsListV2Result> Invoke(GetSecretsListV2InvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetSecretsListV2Result>("vault:kv/getSecretsListV2:getSecretsListV2", args ?? new GetSecretsListV2InvokeArgs(), options.WithDefaults());
    }


    public sealed class GetSecretsListV2Args : global::Pulumi.InvokeArgs
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
        [Input("name")]
        public string? Name { get; set; }

        /// <summary>
        /// The namespace of the target resource.
        /// The value should not contain leading or trailing forward slashes.
        /// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
        /// *Available only for Vault Enterprise*.
        /// </summary>
        [Input("namespace")]
        public string? Namespace { get; set; }

        public GetSecretsListV2Args()
        {
        }
        public static new GetSecretsListV2Args Empty => new GetSecretsListV2Args();
    }

    public sealed class GetSecretsListV2InvokeArgs : global::Pulumi.InvokeArgs
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
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The namespace of the target resource.
        /// The value should not contain leading or trailing forward slashes.
        /// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
        /// *Available only for Vault Enterprise*.
        /// </summary>
        [Input("namespace")]
        public Input<string>? Namespace { get; set; }

        public GetSecretsListV2InvokeArgs()
        {
        }
        public static new GetSecretsListV2InvokeArgs Empty => new GetSecretsListV2InvokeArgs();
    }


    [OutputType]
    public sealed class GetSecretsListV2Result
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string Mount;
        public readonly string? Name;
        /// <summary>
        /// List of all secret names listed under the given path.
        /// </summary>
        public readonly ImmutableArray<string> Names;
        public readonly string? Namespace;
        /// <summary>
        /// Full path where the KV-V2 secrets are listed.
        /// </summary>
        public readonly string Path;

        [OutputConstructor]
        private GetSecretsListV2Result(
            string id,

            string mount,

            string? name,

            ImmutableArray<string> names,

            string? @namespace,

            string path)
        {
            Id = id;
            Mount = mount;
            Name = name;
            Names = names;
            Namespace = @namespace;
            Path = path;
        }
    }
}
