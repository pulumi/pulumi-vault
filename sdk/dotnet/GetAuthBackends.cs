// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vault
{
    public static class GetAuthBackends
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
        ///     var example = Vault.GetAuthBackends.Invoke();
        /// 
        /// });
        /// ```
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Vault = Pulumi.Vault;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var example_filter = Vault.GetAuthBackends.Invoke(new()
        ///     {
        ///         Type = "kubernetes",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Task<GetAuthBackendsResult> InvokeAsync(GetAuthBackendsArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetAuthBackendsResult>("vault:index/getAuthBackends:getAuthBackends", args ?? new GetAuthBackendsArgs(), options.WithDefaults());

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
        ///     var example = Vault.GetAuthBackends.Invoke();
        /// 
        /// });
        /// ```
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Vault = Pulumi.Vault;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var example_filter = Vault.GetAuthBackends.Invoke(new()
        ///     {
        ///         Type = "kubernetes",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetAuthBackendsResult> Invoke(GetAuthBackendsInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetAuthBackendsResult>("vault:index/getAuthBackends:getAuthBackends", args ?? new GetAuthBackendsInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetAuthBackendsArgs : global::Pulumi.InvokeArgs
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
        /// The name of the auth method type. Allows filtering of backends returned by type.
        /// </summary>
        [Input("type")]
        public string? Type { get; set; }

        public GetAuthBackendsArgs()
        {
        }
        public static new GetAuthBackendsArgs Empty => new GetAuthBackendsArgs();
    }

    public sealed class GetAuthBackendsInvokeArgs : global::Pulumi.InvokeArgs
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
        /// The name of the auth method type. Allows filtering of backends returned by type.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        public GetAuthBackendsInvokeArgs()
        {
        }
        public static new GetAuthBackendsInvokeArgs Empty => new GetAuthBackendsInvokeArgs();
    }


    [OutputType]
    public sealed class GetAuthBackendsResult
    {
        /// <summary>
        /// The accessor IDs for the auth methods.
        /// </summary>
        public readonly ImmutableArray<string> Accessors;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string? Namespace;
        /// <summary>
        /// List of auth backend mount points.
        /// </summary>
        public readonly ImmutableArray<string> Paths;
        public readonly string? Type;

        [OutputConstructor]
        private GetAuthBackendsResult(
            ImmutableArray<string> accessors,

            string id,

            string? @namespace,

            ImmutableArray<string> paths,

            string? type)
        {
            Accessors = accessors;
            Id = id;
            Namespace = @namespace;
            Paths = paths;
            Type = type;
        }
    }
}
