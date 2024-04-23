// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vault.Azure
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
    ///     var example = new Vault.AuthBackend("example", new()
    ///     {
    ///         Type = "azure",
    ///     });
    /// 
    ///     var exampleAuthBackendConfig = new Vault.Azure.AuthBackendConfig("example", new()
    ///     {
    ///         Backend = example.Path,
    ///         TenantId = "11111111-2222-3333-4444-555555555555",
    ///         ClientId = "11111111-2222-3333-4444-555555555555",
    ///         ClientSecret = "01234567890123456789",
    ///         Resource = "https://vault.hashicorp.com",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Azure auth backends can be imported using `auth/`, the `backend` path, and `/config` e.g.
    /// 
    /// ```sh
    /// $ pulumi import vault:azure/authBackendConfig:AuthBackendConfig example auth/azure/config
    /// ```
    /// </summary>
    [VaultResourceType("vault:azure/authBackendConfig:AuthBackendConfig")]
    public partial class AuthBackendConfig : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The path the Azure auth backend being configured was
        /// mounted at.  Defaults to `azure`.
        /// </summary>
        [Output("backend")]
        public Output<string?> Backend { get; private set; } = null!;

        /// <summary>
        /// The client id for credentials to query the Azure APIs.
        /// Currently read permissions to query compute resources are required.
        /// </summary>
        [Output("clientId")]
        public Output<string?> ClientId { get; private set; } = null!;

        /// <summary>
        /// The client secret for credentials to query the
        /// Azure APIs.
        /// </summary>
        [Output("clientSecret")]
        public Output<string?> ClientSecret { get; private set; } = null!;

        /// <summary>
        /// The Azure cloud environment. Valid values:
        /// AzurePublicCloud, AzureUSGovernmentCloud, AzureChinaCloud,
        /// AzureGermanCloud.  Defaults to `AzurePublicCloud`.
        /// </summary>
        [Output("environment")]
        public Output<string?> Environment { get; private set; } = null!;

        /// <summary>
        /// The namespace to provision the resource in.
        /// The value should not contain leading or trailing forward slashes.
        /// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
        /// *Available only for Vault Enterprise*.
        /// </summary>
        [Output("namespace")]
        public Output<string?> Namespace { get; private set; } = null!;

        /// <summary>
        /// The configured URL for the application registered in
        /// Azure Active Directory.
        /// </summary>
        [Output("resource")]
        public Output<string> Resource { get; private set; } = null!;

        /// <summary>
        /// The tenant id for the Azure Active Directory
        /// organization.
        /// </summary>
        [Output("tenantId")]
        public Output<string> TenantId { get; private set; } = null!;


        /// <summary>
        /// Create a AuthBackendConfig resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AuthBackendConfig(string name, AuthBackendConfigArgs args, CustomResourceOptions? options = null)
            : base("vault:azure/authBackendConfig:AuthBackendConfig", name, args ?? new AuthBackendConfigArgs(), MakeResourceOptions(options, ""))
        {
        }

        private AuthBackendConfig(string name, Input<string> id, AuthBackendConfigState? state = null, CustomResourceOptions? options = null)
            : base("vault:azure/authBackendConfig:AuthBackendConfig", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                AdditionalSecretOutputs =
                {
                    "clientId",
                    "clientSecret",
                    "tenantId",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing AuthBackendConfig resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static AuthBackendConfig Get(string name, Input<string> id, AuthBackendConfigState? state = null, CustomResourceOptions? options = null)
        {
            return new AuthBackendConfig(name, id, state, options);
        }
    }

    public sealed class AuthBackendConfigArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The path the Azure auth backend being configured was
        /// mounted at.  Defaults to `azure`.
        /// </summary>
        [Input("backend")]
        public Input<string>? Backend { get; set; }

        [Input("clientId")]
        private Input<string>? _clientId;

        /// <summary>
        /// The client id for credentials to query the Azure APIs.
        /// Currently read permissions to query compute resources are required.
        /// </summary>
        public Input<string>? ClientId
        {
            get => _clientId;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _clientId = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        [Input("clientSecret")]
        private Input<string>? _clientSecret;

        /// <summary>
        /// The client secret for credentials to query the
        /// Azure APIs.
        /// </summary>
        public Input<string>? ClientSecret
        {
            get => _clientSecret;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _clientSecret = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// The Azure cloud environment. Valid values:
        /// AzurePublicCloud, AzureUSGovernmentCloud, AzureChinaCloud,
        /// AzureGermanCloud.  Defaults to `AzurePublicCloud`.
        /// </summary>
        [Input("environment")]
        public Input<string>? Environment { get; set; }

        /// <summary>
        /// The namespace to provision the resource in.
        /// The value should not contain leading or trailing forward slashes.
        /// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
        /// *Available only for Vault Enterprise*.
        /// </summary>
        [Input("namespace")]
        public Input<string>? Namespace { get; set; }

        /// <summary>
        /// The configured URL for the application registered in
        /// Azure Active Directory.
        /// </summary>
        [Input("resource", required: true)]
        public Input<string> Resource { get; set; } = null!;

        [Input("tenantId", required: true)]
        private Input<string>? _tenantId;

        /// <summary>
        /// The tenant id for the Azure Active Directory
        /// organization.
        /// </summary>
        public Input<string>? TenantId
        {
            get => _tenantId;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _tenantId = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        public AuthBackendConfigArgs()
        {
        }
        public static new AuthBackendConfigArgs Empty => new AuthBackendConfigArgs();
    }

    public sealed class AuthBackendConfigState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The path the Azure auth backend being configured was
        /// mounted at.  Defaults to `azure`.
        /// </summary>
        [Input("backend")]
        public Input<string>? Backend { get; set; }

        [Input("clientId")]
        private Input<string>? _clientId;

        /// <summary>
        /// The client id for credentials to query the Azure APIs.
        /// Currently read permissions to query compute resources are required.
        /// </summary>
        public Input<string>? ClientId
        {
            get => _clientId;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _clientId = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        [Input("clientSecret")]
        private Input<string>? _clientSecret;

        /// <summary>
        /// The client secret for credentials to query the
        /// Azure APIs.
        /// </summary>
        public Input<string>? ClientSecret
        {
            get => _clientSecret;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _clientSecret = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// The Azure cloud environment. Valid values:
        /// AzurePublicCloud, AzureUSGovernmentCloud, AzureChinaCloud,
        /// AzureGermanCloud.  Defaults to `AzurePublicCloud`.
        /// </summary>
        [Input("environment")]
        public Input<string>? Environment { get; set; }

        /// <summary>
        /// The namespace to provision the resource in.
        /// The value should not contain leading or trailing forward slashes.
        /// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
        /// *Available only for Vault Enterprise*.
        /// </summary>
        [Input("namespace")]
        public Input<string>? Namespace { get; set; }

        /// <summary>
        /// The configured URL for the application registered in
        /// Azure Active Directory.
        /// </summary>
        [Input("resource")]
        public Input<string>? Resource { get; set; }

        [Input("tenantId")]
        private Input<string>? _tenantId;

        /// <summary>
        /// The tenant id for the Azure Active Directory
        /// organization.
        /// </summary>
        public Input<string>? TenantId
        {
            get => _tenantId;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _tenantId = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        public AuthBackendConfigState()
        {
        }
        public static new AuthBackendConfigState Empty => new AuthBackendConfigState();
    }
}
