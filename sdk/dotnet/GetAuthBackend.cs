// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vault
{
    public static class GetAuthBackend
    {
        /// <summary>
        /// {{% examples %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetAuthBackendResult> InvokeAsync(GetAuthBackendArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetAuthBackendResult>("vault:index/getAuthBackend:getAuthBackend", args ?? new GetAuthBackendArgs(), options.WithVersion());
    }


    public sealed class GetAuthBackendArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The auth backend mount point.
        /// </summary>
        [Input("path", required: true)]
        public string Path { get; set; } = null!;

        public GetAuthBackendArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetAuthBackendResult
    {
        /// <summary>
        /// The accessor for this auth method
        /// </summary>
        public readonly string Accessor;
        /// <summary>
        /// The default lease duration in seconds.
        /// </summary>
        public readonly int DefaultLeaseTtlSeconds;
        /// <summary>
        /// A description of the auth method.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// id is the provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Speficies whether to show this mount in the UI-specific listing endpoint.
        /// </summary>
        public readonly string ListingVisibility;
        /// <summary>
        /// Specifies if the auth method is local only.
        /// </summary>
        public readonly bool Local;
        /// <summary>
        /// The maximum lease duration in seconds.
        /// </summary>
        public readonly int MaxLeaseTtlSeconds;
        public readonly string Path;
        /// <summary>
        /// The name of the auth method type.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetAuthBackendResult(
            string accessor,

            int defaultLeaseTtlSeconds,

            string description,

            string id,

            string listingVisibility,

            bool local,

            int maxLeaseTtlSeconds,

            string path,

            string type)
        {
            Accessor = accessor;
            DefaultLeaseTtlSeconds = defaultLeaseTtlSeconds;
            Description = description;
            Id = id;
            ListingVisibility = listingVisibility;
            Local = local;
            MaxLeaseTtlSeconds = maxLeaseTtlSeconds;
            Path = path;
            Type = type;
        }
    }
}
