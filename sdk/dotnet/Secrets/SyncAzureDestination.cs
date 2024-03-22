// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vault.Secrets
{
    /// <summary>
    /// ## Example Usage
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
    ///     var az = new Vault.Secrets.SyncAzureDestination("az", new()
    ///     {
    ///         KeyVaultUri = @var.Key_vault_uri,
    ///         ClientId = @var.Client_id,
    ///         ClientSecret = @var.Client_secret,
    ///         TenantId = @var.Tenant_id,
    ///         SecretNameTemplate = "vault_{{ .MountAccessor | lowercase }}_{{ .SecretPath | lowercase }}",
    ///         CustomTags = 
    ///         {
    ///             { "foo", "bar" },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ## Import
    /// 
    /// Azure Secrets sync destinations can be imported using the `name`, e.g.
    /// 
    /// ```sh
    /// $ pulumi import vault:secrets/syncAzureDestination:SyncAzureDestination az az-dest
    /// ```
    /// </summary>
    [VaultResourceType("vault:secrets/syncAzureDestination:SyncAzureDestination")]
    public partial class SyncAzureDestination : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Client ID of an Azure app registration.
        /// Can be omitted and directly provided to Vault using the `AZURE_CLIENT_ID` environment
        /// variable.
        /// </summary>
        [Output("clientId")]
        public Output<string?> ClientId { get; private set; } = null!;

        /// <summary>
        /// Client Secret of an Azure app registration.
        /// Can be omitted and directly provided to Vault using the `AZURE_CLIENT_SECRET` environment
        /// variable.
        /// </summary>
        [Output("clientSecret")]
        public Output<string?> ClientSecret { get; private set; } = null!;

        /// <summary>
        /// Specifies a cloud for the client. The default is Azure Public Cloud.
        /// </summary>
        [Output("cloud")]
        public Output<string?> Cloud { get; private set; } = null!;

        /// <summary>
        /// Custom tags to set on the secret managed at the destination.
        /// </summary>
        [Output("customTags")]
        public Output<ImmutableDictionary<string, object>?> CustomTags { get; private set; } = null!;

        /// <summary>
        /// URI of an existing Azure Key Vault instance.
        /// Can be omitted and directly provided to Vault using the `KEY_VAULT_URI` environment
        /// variable.
        /// </summary>
        [Output("keyVaultUri")]
        public Output<string?> KeyVaultUri { get; private set; } = null!;

        /// <summary>
        /// Unique name of the Azure destination.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The namespace to provision the resource in.
        /// The value should not contain leading or trailing forward slashes.
        /// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
        /// </summary>
        [Output("namespace")]
        public Output<string?> Namespace { get; private set; } = null!;

        /// <summary>
        /// Template describing how to generate external secret names.
        /// Supports a subset of the Go Template syntax.
        /// </summary>
        [Output("secretNameTemplate")]
        public Output<string> SecretNameTemplate { get; private set; } = null!;

        /// <summary>
        /// ID of the target Azure tenant.
        /// Can be omitted and directly provided to Vault using the `AZURE_TENANT_ID` environment
        /// variable.
        /// </summary>
        [Output("tenantId")]
        public Output<string?> TenantId { get; private set; } = null!;

        /// <summary>
        /// The type of the secrets destination (`azure-kv`).
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a SyncAzureDestination resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public SyncAzureDestination(string name, SyncAzureDestinationArgs? args = null, CustomResourceOptions? options = null)
            : base("vault:secrets/syncAzureDestination:SyncAzureDestination", name, args ?? new SyncAzureDestinationArgs(), MakeResourceOptions(options, ""))
        {
        }

        private SyncAzureDestination(string name, Input<string> id, SyncAzureDestinationState? state = null, CustomResourceOptions? options = null)
            : base("vault:secrets/syncAzureDestination:SyncAzureDestination", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                AdditionalSecretOutputs =
                {
                    "clientSecret",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing SyncAzureDestination resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static SyncAzureDestination Get(string name, Input<string> id, SyncAzureDestinationState? state = null, CustomResourceOptions? options = null)
        {
            return new SyncAzureDestination(name, id, state, options);
        }
    }

    public sealed class SyncAzureDestinationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Client ID of an Azure app registration.
        /// Can be omitted and directly provided to Vault using the `AZURE_CLIENT_ID` environment
        /// variable.
        /// </summary>
        [Input("clientId")]
        public Input<string>? ClientId { get; set; }

        [Input("clientSecret")]
        private Input<string>? _clientSecret;

        /// <summary>
        /// Client Secret of an Azure app registration.
        /// Can be omitted and directly provided to Vault using the `AZURE_CLIENT_SECRET` environment
        /// variable.
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
        /// Specifies a cloud for the client. The default is Azure Public Cloud.
        /// </summary>
        [Input("cloud")]
        public Input<string>? Cloud { get; set; }

        [Input("customTags")]
        private InputMap<object>? _customTags;

        /// <summary>
        /// Custom tags to set on the secret managed at the destination.
        /// </summary>
        public InputMap<object> CustomTags
        {
            get => _customTags ?? (_customTags = new InputMap<object>());
            set => _customTags = value;
        }

        /// <summary>
        /// URI of an existing Azure Key Vault instance.
        /// Can be omitted and directly provided to Vault using the `KEY_VAULT_URI` environment
        /// variable.
        /// </summary>
        [Input("keyVaultUri")]
        public Input<string>? KeyVaultUri { get; set; }

        /// <summary>
        /// Unique name of the Azure destination.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The namespace to provision the resource in.
        /// The value should not contain leading or trailing forward slashes.
        /// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
        /// </summary>
        [Input("namespace")]
        public Input<string>? Namespace { get; set; }

        /// <summary>
        /// Template describing how to generate external secret names.
        /// Supports a subset of the Go Template syntax.
        /// </summary>
        [Input("secretNameTemplate")]
        public Input<string>? SecretNameTemplate { get; set; }

        /// <summary>
        /// ID of the target Azure tenant.
        /// Can be omitted and directly provided to Vault using the `AZURE_TENANT_ID` environment
        /// variable.
        /// </summary>
        [Input("tenantId")]
        public Input<string>? TenantId { get; set; }

        public SyncAzureDestinationArgs()
        {
        }
        public static new SyncAzureDestinationArgs Empty => new SyncAzureDestinationArgs();
    }

    public sealed class SyncAzureDestinationState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Client ID of an Azure app registration.
        /// Can be omitted and directly provided to Vault using the `AZURE_CLIENT_ID` environment
        /// variable.
        /// </summary>
        [Input("clientId")]
        public Input<string>? ClientId { get; set; }

        [Input("clientSecret")]
        private Input<string>? _clientSecret;

        /// <summary>
        /// Client Secret of an Azure app registration.
        /// Can be omitted and directly provided to Vault using the `AZURE_CLIENT_SECRET` environment
        /// variable.
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
        /// Specifies a cloud for the client. The default is Azure Public Cloud.
        /// </summary>
        [Input("cloud")]
        public Input<string>? Cloud { get; set; }

        [Input("customTags")]
        private InputMap<object>? _customTags;

        /// <summary>
        /// Custom tags to set on the secret managed at the destination.
        /// </summary>
        public InputMap<object> CustomTags
        {
            get => _customTags ?? (_customTags = new InputMap<object>());
            set => _customTags = value;
        }

        /// <summary>
        /// URI of an existing Azure Key Vault instance.
        /// Can be omitted and directly provided to Vault using the `KEY_VAULT_URI` environment
        /// variable.
        /// </summary>
        [Input("keyVaultUri")]
        public Input<string>? KeyVaultUri { get; set; }

        /// <summary>
        /// Unique name of the Azure destination.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The namespace to provision the resource in.
        /// The value should not contain leading or trailing forward slashes.
        /// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
        /// </summary>
        [Input("namespace")]
        public Input<string>? Namespace { get; set; }

        /// <summary>
        /// Template describing how to generate external secret names.
        /// Supports a subset of the Go Template syntax.
        /// </summary>
        [Input("secretNameTemplate")]
        public Input<string>? SecretNameTemplate { get; set; }

        /// <summary>
        /// ID of the target Azure tenant.
        /// Can be omitted and directly provided to Vault using the `AZURE_TENANT_ID` environment
        /// variable.
        /// </summary>
        [Input("tenantId")]
        public Input<string>? TenantId { get; set; }

        /// <summary>
        /// The type of the secrets destination (`azure-kv`).
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        public SyncAzureDestinationState()
        {
        }
        public static new SyncAzureDestinationState Empty => new SyncAzureDestinationState();
    }
}
