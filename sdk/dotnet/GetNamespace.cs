// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vault
{
    public static class GetNamespace
    {
        /// <summary>
        /// ## Example Usage
        /// 
        /// ### Current namespace
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Vault = Pulumi.Vault;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var current = Vault.GetNamespace.Invoke();
        /// 
        /// });
        /// ```
        /// 
        /// ### Single namespace
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Vault = Pulumi.Vault;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var ns1 = Vault.GetNamespace.Invoke(new()
        ///     {
        ///         Path = "ns1",
        ///     });
        /// 
        /// });
        /// ```
        /// 
        /// ### Nested namespace
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Vault = Pulumi.Vault;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var child = Vault.GetNamespace.Invoke(new()
        ///     {
        ///         Namespace = "parent",
        ///         Path = "child",
        ///     });
        /// 
        ///     var fullPath = child.Apply(getNamespaceResult =&gt; getNamespaceResult.Id);
        /// 
        ///     // -&gt; foo/parent/child/
        ///     var pathFq = child.Apply(getNamespaceResult =&gt; getNamespaceResult.PathFq);
        /// 
        /// });
        /// ```
        /// </summary>
        public static Task<GetNamespaceResult> InvokeAsync(GetNamespaceArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetNamespaceResult>("vault:index/getNamespace:getNamespace", args ?? new GetNamespaceArgs(), options.WithDefaults());

        /// <summary>
        /// ## Example Usage
        /// 
        /// ### Current namespace
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Vault = Pulumi.Vault;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var current = Vault.GetNamespace.Invoke();
        /// 
        /// });
        /// ```
        /// 
        /// ### Single namespace
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Vault = Pulumi.Vault;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var ns1 = Vault.GetNamespace.Invoke(new()
        ///     {
        ///         Path = "ns1",
        ///     });
        /// 
        /// });
        /// ```
        /// 
        /// ### Nested namespace
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Vault = Pulumi.Vault;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var child = Vault.GetNamespace.Invoke(new()
        ///     {
        ///         Namespace = "parent",
        ///         Path = "child",
        ///     });
        /// 
        ///     var fullPath = child.Apply(getNamespaceResult =&gt; getNamespaceResult.Id);
        /// 
        ///     // -&gt; foo/parent/child/
        ///     var pathFq = child.Apply(getNamespaceResult =&gt; getNamespaceResult.PathFq);
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetNamespaceResult> Invoke(GetNamespaceInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetNamespaceResult>("vault:index/getNamespace:getNamespace", args ?? new GetNamespaceInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetNamespaceArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The namespace to provision the resource in.
        /// The value should not contain leading or trailing forward slashes.
        /// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
        /// </summary>
        [Input("namespace")]
        public string? Namespace { get; set; }

        /// <summary>
        /// The path of the namespace. Must not have a trailing `/`.
        /// If not specified or empty, path attributes are set for the current namespace
        /// based on the `namespace` arguments of the provider and this data source.
        /// Other path related attributes will be empty in this case.
        /// </summary>
        [Input("path")]
        public string? Path { get; set; }

        public GetNamespaceArgs()
        {
        }
        public static new GetNamespaceArgs Empty => new GetNamespaceArgs();
    }

    public sealed class GetNamespaceInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The namespace to provision the resource in.
        /// The value should not contain leading or trailing forward slashes.
        /// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
        /// </summary>
        [Input("namespace")]
        public Input<string>? Namespace { get; set; }

        /// <summary>
        /// The path of the namespace. Must not have a trailing `/`.
        /// If not specified or empty, path attributes are set for the current namespace
        /// based on the `namespace` arguments of the provider and this data source.
        /// Other path related attributes will be empty in this case.
        /// </summary>
        [Input("path")]
        public Input<string>? Path { get; set; }

        public GetNamespaceInvokeArgs()
        {
        }
        public static new GetNamespaceInvokeArgs Empty => new GetNamespaceInvokeArgs();
    }


    [OutputType]
    public sealed class GetNamespaceResult
    {
        /// <summary>
        /// (Optional) A map of strings containing arbitrary metadata for the namespace.
        /// Only fetched if `path` is specified.
        /// *Requires Vault 1.12+.*
        /// </summary>
        public readonly ImmutableDictionary<string, string> CustomMetadata;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string? Namespace;
        /// <summary>
        /// Vault server's internal ID of the namespace.
        /// Only fetched if `path` is specified.
        /// </summary>
        public readonly string NamespaceId;
        public readonly string? Path;
        /// <summary>
        /// The fully qualified path to the namespace. Useful when provisioning resources in a child `namespace`.
        /// The path is relative to the provider's `namespace` argument.
        /// </summary>
        public readonly string PathFq;

        [OutputConstructor]
        private GetNamespaceResult(
            ImmutableDictionary<string, string> customMetadata,

            string id,

            string? @namespace,

            string namespaceId,

            string? path,

            string pathFq)
        {
            CustomMetadata = customMetadata;
            Id = id;
            Namespace = @namespace;
            NamespaceId = namespaceId;
            Path = path;
            PathFq = pathFq;
        }
    }
}
