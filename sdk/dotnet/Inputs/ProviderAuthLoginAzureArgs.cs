// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vault.Inputs
{

    public sealed class ProviderAuthLoginAzureArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The identity's client ID.
        /// </summary>
        [Input("clientId")]
        public Input<string>? ClientId { get; set; }

        /// <summary>
        /// A signed JSON Web Token. If not specified on will be created automatically
        /// </summary>
        [Input("jwt")]
        public Input<string>? Jwt { get; set; }

        /// <summary>
        /// The path where the authentication engine is mounted.
        /// </summary>
        [Input("mount")]
        public Input<string>? Mount { get; set; }

        /// <summary>
        /// The authentication engine's namespace. Conflicts with use_root_namespace
        /// </summary>
        [Input("namespace")]
        public Input<string>? Namespace { get; set; }

        /// <summary>
        /// The resource group for the machine that generated the MSI token. This information can be obtained through instance metadata.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// Name of the login role.
        /// </summary>
        [Input("role", required: true)]
        public Input<string> Role { get; set; } = null!;

        /// <summary>
        /// The scopes to include in the token request.
        /// </summary>
        [Input("scope")]
        public Input<string>? Scope { get; set; }

        /// <summary>
        /// The subscription ID for the machine that generated the MSI token. This information can be obtained through instance metadata.
        /// </summary>
        [Input("subscriptionId", required: true)]
        public Input<string> SubscriptionId { get; set; } = null!;

        /// <summary>
        /// Provides the tenant ID to use in a multi-tenant authentication scenario.
        /// </summary>
        [Input("tenantId")]
        public Input<string>? TenantId { get; set; }

        /// <summary>
        /// Authenticate to the root Vault namespace. Conflicts with namespace
        /// </summary>
        [Input("useRootNamespace")]
        public Input<bool>? UseRootNamespace { get; set; }

        /// <summary>
        /// The virtual machine name for the machine that generated the MSI token. This information can be obtained through instance metadata.
        /// </summary>
        [Input("vmName")]
        public Input<string>? VmName { get; set; }

        /// <summary>
        /// The virtual machine scale set name for the machine that generated the MSI token. This information can be obtained through instance metadata.
        /// </summary>
        [Input("vmssName")]
        public Input<string>? VmssName { get; set; }

        public ProviderAuthLoginAzureArgs()
        {
        }
        public static new ProviderAuthLoginAzureArgs Empty => new ProviderAuthLoginAzureArgs();
    }
}
