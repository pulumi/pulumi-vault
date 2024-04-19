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
    /// ### *Vault-1.9 And Above*
    /// 
    /// &lt;!--Start PulumiCodeChooser --&gt;
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Vault = Pulumi.Vault;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var azure = new Vault.Azure.Backend("azure", new()
    ///     {
    ///         UseMicrosoftGraphApi = true,
    ///         SubscriptionId = "11111111-2222-3333-4444-111111111111",
    ///         TenantId = "11111111-2222-3333-4444-222222222222",
    ///         ClientId = "11111111-2222-3333-4444-333333333333",
    ///         ClientSecret = "12345678901234567890",
    ///         Environment = "AzurePublicCloud",
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ### *Vault-1.8 And Below*
    /// 
    /// &lt;!--Start PulumiCodeChooser --&gt;
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Vault = Pulumi.Vault;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var azure = new Vault.Azure.Backend("azure", new()
    ///     {
    ///         UseMicrosoftGraphApi = false,
    ///         SubscriptionId = "11111111-2222-3333-4444-111111111111",
    ///         TenantId = "11111111-2222-3333-4444-222222222222",
    ///         ClientId = "11111111-2222-3333-4444-333333333333",
    ///         ClientSecret = "12345678901234567890",
    ///         Environment = "AzurePublicCloud",
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// </summary>
    [VaultResourceType("vault:azure/backend:Backend")]
    public partial class Backend : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The OAuth2 client id to connect to Azure.
        /// </summary>
        [Output("clientId")]
        public Output<string?> ClientId { get; private set; } = null!;

        /// <summary>
        /// The OAuth2 client secret to connect to Azure.
        /// </summary>
        [Output("clientSecret")]
        public Output<string?> ClientSecret { get; private set; } = null!;

        /// <summary>
        /// Human-friendly description of the mount for the backend.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// If set, opts out of mount migration on path updates.
        /// See here for more info on [Mount Migration](https://www.vaultproject.io/docs/concepts/mount-migration)
        /// </summary>
        [Output("disableRemount")]
        public Output<bool?> DisableRemount { get; private set; } = null!;

        /// <summary>
        /// The Azure environment.
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
        /// The unique path this backend should be mounted at. Defaults to `azure`.
        /// </summary>
        [Output("path")]
        public Output<string?> Path { get; private set; } = null!;

        /// <summary>
        /// The subscription id for the Azure Active Directory.
        /// </summary>
        [Output("subscriptionId")]
        public Output<string> SubscriptionId { get; private set; } = null!;

        /// <summary>
        /// The tenant id for the Azure Active Directory.
        /// </summary>
        [Output("tenantId")]
        public Output<string> TenantId { get; private set; } = null!;

        /// <summary>
        /// Use the Microsoft Graph API. Should be set to true on vault-1.10+
        /// </summary>
        [Output("useMicrosoftGraphApi")]
        public Output<bool> UseMicrosoftGraphApi { get; private set; } = null!;


        /// <summary>
        /// Create a Backend resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Backend(string name, BackendArgs args, CustomResourceOptions? options = null)
            : base("vault:azure/backend:Backend", name, args ?? new BackendArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Backend(string name, Input<string> id, BackendState? state = null, CustomResourceOptions? options = null)
            : base("vault:azure/backend:Backend", name, state, MakeResourceOptions(options, id))
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
                    "subscriptionId",
                    "tenantId",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Backend resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Backend Get(string name, Input<string> id, BackendState? state = null, CustomResourceOptions? options = null)
        {
            return new Backend(name, id, state, options);
        }
    }

    public sealed class BackendArgs : global::Pulumi.ResourceArgs
    {
        [Input("clientId")]
        private Input<string>? _clientId;

        /// <summary>
        /// The OAuth2 client id to connect to Azure.
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
        /// The OAuth2 client secret to connect to Azure.
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
        /// Human-friendly description of the mount for the backend.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// If set, opts out of mount migration on path updates.
        /// See here for more info on [Mount Migration](https://www.vaultproject.io/docs/concepts/mount-migration)
        /// </summary>
        [Input("disableRemount")]
        public Input<bool>? DisableRemount { get; set; }

        /// <summary>
        /// The Azure environment.
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
        /// The unique path this backend should be mounted at. Defaults to `azure`.
        /// </summary>
        [Input("path")]
        public Input<string>? Path { get; set; }

        [Input("subscriptionId", required: true)]
        private Input<string>? _subscriptionId;

        /// <summary>
        /// The subscription id for the Azure Active Directory.
        /// </summary>
        public Input<string>? SubscriptionId
        {
            get => _subscriptionId;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _subscriptionId = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        [Input("tenantId", required: true)]
        private Input<string>? _tenantId;

        /// <summary>
        /// The tenant id for the Azure Active Directory.
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

        /// <summary>
        /// Use the Microsoft Graph API. Should be set to true on vault-1.10+
        /// </summary>
        [Input("useMicrosoftGraphApi")]
        public Input<bool>? UseMicrosoftGraphApi { get; set; }

        public BackendArgs()
        {
        }
        public static new BackendArgs Empty => new BackendArgs();
    }

    public sealed class BackendState : global::Pulumi.ResourceArgs
    {
        [Input("clientId")]
        private Input<string>? _clientId;

        /// <summary>
        /// The OAuth2 client id to connect to Azure.
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
        /// The OAuth2 client secret to connect to Azure.
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
        /// Human-friendly description of the mount for the backend.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// If set, opts out of mount migration on path updates.
        /// See here for more info on [Mount Migration](https://www.vaultproject.io/docs/concepts/mount-migration)
        /// </summary>
        [Input("disableRemount")]
        public Input<bool>? DisableRemount { get; set; }

        /// <summary>
        /// The Azure environment.
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
        /// The unique path this backend should be mounted at. Defaults to `azure`.
        /// </summary>
        [Input("path")]
        public Input<string>? Path { get; set; }

        [Input("subscriptionId")]
        private Input<string>? _subscriptionId;

        /// <summary>
        /// The subscription id for the Azure Active Directory.
        /// </summary>
        public Input<string>? SubscriptionId
        {
            get => _subscriptionId;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _subscriptionId = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        [Input("tenantId")]
        private Input<string>? _tenantId;

        /// <summary>
        /// The tenant id for the Azure Active Directory.
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

        /// <summary>
        /// Use the Microsoft Graph API. Should be set to true on vault-1.10+
        /// </summary>
        [Input("useMicrosoftGraphApi")]
        public Input<bool>? UseMicrosoftGraphApi { get; set; }

        public BackendState()
        {
        }
        public static new BackendState Empty => new BackendState();
    }
}
