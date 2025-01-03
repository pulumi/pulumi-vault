// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vault
{
    public static class GetNamespaces
    {
        /// <summary>
        /// ## Example Usage
        /// 
        /// ### Child namespaces
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Vault = Pulumi.Vault;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var children = Vault.GetNamespaces.Invoke();
        /// 
        /// });
        /// ```
        /// 
        /// ### Nested namespace
        /// 
        /// To fetch the details of nested namespaces:
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Vault = Pulumi.Vault;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var children = Vault.GetNamespaces.Invoke(new()
        ///     {
        ///         Namespace = "parent",
        ///     });
        /// 
        ///     var child = .ToDictionary(item =&gt; {
        ///         var __key = item.Key;
        ///         return __key;
        ///     }, item =&gt; {
        ///         var __key = item.Key;
        ///         return Vault.GetNamespace.Invoke(new()
        ///         {
        ///             Namespace = _arg0_.Namespace,
        ///             Path = __key,
        ///         });
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Task<GetNamespacesResult> InvokeAsync(GetNamespacesArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetNamespacesResult>("vault:index/getNamespaces:getNamespaces", args ?? new GetNamespacesArgs(), options.WithDefaults());

        /// <summary>
        /// ## Example Usage
        /// 
        /// ### Child namespaces
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Vault = Pulumi.Vault;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var children = Vault.GetNamespaces.Invoke();
        /// 
        /// });
        /// ```
        /// 
        /// ### Nested namespace
        /// 
        /// To fetch the details of nested namespaces:
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Vault = Pulumi.Vault;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var children = Vault.GetNamespaces.Invoke(new()
        ///     {
        ///         Namespace = "parent",
        ///     });
        /// 
        ///     var child = .ToDictionary(item =&gt; {
        ///         var __key = item.Key;
        ///         return __key;
        ///     }, item =&gt; {
        ///         var __key = item.Key;
        ///         return Vault.GetNamespace.Invoke(new()
        ///         {
        ///             Namespace = _arg0_.Namespace,
        ///             Path = __key,
        ///         });
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetNamespacesResult> Invoke(GetNamespacesInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetNamespacesResult>("vault:index/getNamespaces:getNamespaces", args ?? new GetNamespacesInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// ## Example Usage
        /// 
        /// ### Child namespaces
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Vault = Pulumi.Vault;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var children = Vault.GetNamespaces.Invoke();
        /// 
        /// });
        /// ```
        /// 
        /// ### Nested namespace
        /// 
        /// To fetch the details of nested namespaces:
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Vault = Pulumi.Vault;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var children = Vault.GetNamespaces.Invoke(new()
        ///     {
        ///         Namespace = "parent",
        ///     });
        /// 
        ///     var child = .ToDictionary(item =&gt; {
        ///         var __key = item.Key;
        ///         return __key;
        ///     }, item =&gt; {
        ///         var __key = item.Key;
        ///         return Vault.GetNamespace.Invoke(new()
        ///         {
        ///             Namespace = _arg0_.Namespace,
        ///             Path = __key,
        ///         });
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetNamespacesResult> Invoke(GetNamespacesInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetNamespacesResult>("vault:index/getNamespaces:getNamespaces", args ?? new GetNamespacesInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetNamespacesArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The namespace to provision the resource in.
        /// The value should not contain leading or trailing forward slashes.
        /// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
        /// </summary>
        [Input("namespace")]
        public string? Namespace { get; set; }

        public GetNamespacesArgs()
        {
        }
        public static new GetNamespacesArgs Empty => new GetNamespacesArgs();
    }

    public sealed class GetNamespacesInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The namespace to provision the resource in.
        /// The value should not contain leading or trailing forward slashes.
        /// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
        /// </summary>
        [Input("namespace")]
        public Input<string>? Namespace { get; set; }

        public GetNamespacesInvokeArgs()
        {
        }
        public static new GetNamespacesInvokeArgs Empty => new GetNamespacesInvokeArgs();
    }


    [OutputType]
    public sealed class GetNamespacesResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string? Namespace;
        /// <summary>
        /// Set of the paths of direct child namespaces.
        /// </summary>
        public readonly ImmutableArray<string> Paths;

        [OutputConstructor]
        private GetNamespacesResult(
            string id,

            string? @namespace,

            ImmutableArray<string> paths)
        {
            Id = id;
            Namespace = @namespace;
            Paths = paths;
        }
    }
}
