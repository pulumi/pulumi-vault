// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vault
{
    /// <summary>
    /// Provides a resource to manage [Duo MFA](https://www.vaultproject.io/docs/enterprise/mfa/mfa-duo.html).
    /// 
    /// **Note** this feature is available only with Vault Enterprise.
    /// 
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
    ///     var userpass = new Vault.AuthBackend("userpass", new()
    ///     {
    ///         Type = "userpass",
    ///         Path = "userpass",
    ///     });
    /// 
    ///     var myDuo = new Vault.MfaDuo("my_duo", new()
    ///     {
    ///         Name = "my_duo",
    ///         MountAccessor = userpass.Accessor,
    ///         SecretKey = "8C7THtrIigh2rPZQMbguugt8IUftWhMRCOBzbuyz",
    ///         IntegrationKey = "BIACEUEAXI20BNWTEYXT",
    ///         ApiHostname = "api-2b5c39f5.duosecurity.com",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Mounts can be imported using the `path`, e.g.
    /// 
    /// ```sh
    /// $ pulumi import vault:index/mfaDuo:MfaDuo my_duo my_duo
    /// ```
    /// </summary>
    [VaultResourceType("vault:index/mfaDuo:MfaDuo")]
    public partial class MfaDuo : global::Pulumi.CustomResource
    {
        /// <summary>
        /// `(string: &lt;required&gt;)` - API hostname for Duo.
        /// </summary>
        [Output("apiHostname")]
        public Output<string> ApiHostname { get; private set; } = null!;

        /// <summary>
        /// `(string: &lt;required&gt;)` - Integration key for Duo.
        /// </summary>
        [Output("integrationKey")]
        public Output<string> IntegrationKey { get; private set; } = null!;

        /// <summary>
        /// `(string: &lt;required&gt;)` - The mount to tie this method to for use in automatic mappings. The mapping will use the Name field of Aliases associated with this mount as the username in the mapping.
        /// </summary>
        [Output("mountAccessor")]
        public Output<string> MountAccessor { get; private set; } = null!;

        /// <summary>
        /// `(string: &lt;required&gt;)` – Name of the MFA method.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The namespace to provision the resource in.
        /// The value should not contain leading or trailing forward slashes.
        /// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
        /// *Available only for Vault Enterprise*.
        /// </summary>
        [Output("namespace")]
        public Output<string?> Namespace { get; private set; } = null!;

        /// <summary>
        /// `(string)` - Push information for Duo.
        /// </summary>
        [Output("pushInfo")]
        public Output<string?> PushInfo { get; private set; } = null!;

        /// <summary>
        /// `(string: &lt;required&gt;)` - Secret key for Duo.
        /// </summary>
        [Output("secretKey")]
        public Output<string> SecretKey { get; private set; } = null!;

        /// <summary>
        /// `(string)` - A format string for mapping Identity names to MFA method names. Values to substitute should be placed in `{{}}`. For example, `"{{alias.name}}@example.com"`. If blank, the Alias's Name field will be used as-is. Currently-supported mappings:
        /// - alias.name: The name returned by the mount configured via the `mount_accessor` parameter
        /// - entity.name: The name configured for the Entity
        /// - alias.metadata.`&lt;key&gt;`: The value of the Alias's metadata parameter
        /// - entity.metadata.`&lt;key&gt;`: The value of the Entity's metadata parameter
        /// </summary>
        [Output("usernameFormat")]
        public Output<string?> UsernameFormat { get; private set; } = null!;


        /// <summary>
        /// Create a MfaDuo resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public MfaDuo(string name, MfaDuoArgs args, CustomResourceOptions? options = null)
            : base("vault:index/mfaDuo:MfaDuo", name, args ?? new MfaDuoArgs(), MakeResourceOptions(options, ""))
        {
        }

        private MfaDuo(string name, Input<string> id, MfaDuoState? state = null, CustomResourceOptions? options = null)
            : base("vault:index/mfaDuo:MfaDuo", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                AdditionalSecretOutputs =
                {
                    "integrationKey",
                    "secretKey",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing MfaDuo resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static MfaDuo Get(string name, Input<string> id, MfaDuoState? state = null, CustomResourceOptions? options = null)
        {
            return new MfaDuo(name, id, state, options);
        }
    }

    public sealed class MfaDuoArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// `(string: &lt;required&gt;)` - API hostname for Duo.
        /// </summary>
        [Input("apiHostname", required: true)]
        public Input<string> ApiHostname { get; set; } = null!;

        [Input("integrationKey", required: true)]
        private Input<string>? _integrationKey;

        /// <summary>
        /// `(string: &lt;required&gt;)` - Integration key for Duo.
        /// </summary>
        public Input<string>? IntegrationKey
        {
            get => _integrationKey;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _integrationKey = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// `(string: &lt;required&gt;)` - The mount to tie this method to for use in automatic mappings. The mapping will use the Name field of Aliases associated with this mount as the username in the mapping.
        /// </summary>
        [Input("mountAccessor", required: true)]
        public Input<string> MountAccessor { get; set; } = null!;

        /// <summary>
        /// `(string: &lt;required&gt;)` – Name of the MFA method.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The namespace to provision the resource in.
        /// The value should not contain leading or trailing forward slashes.
        /// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
        /// *Available only for Vault Enterprise*.
        /// </summary>
        [Input("namespace")]
        public Input<string>? Namespace { get; set; }

        /// <summary>
        /// `(string)` - Push information for Duo.
        /// </summary>
        [Input("pushInfo")]
        public Input<string>? PushInfo { get; set; }

        [Input("secretKey", required: true)]
        private Input<string>? _secretKey;

        /// <summary>
        /// `(string: &lt;required&gt;)` - Secret key for Duo.
        /// </summary>
        public Input<string>? SecretKey
        {
            get => _secretKey;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _secretKey = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// `(string)` - A format string for mapping Identity names to MFA method names. Values to substitute should be placed in `{{}}`. For example, `"{{alias.name}}@example.com"`. If blank, the Alias's Name field will be used as-is. Currently-supported mappings:
        /// - alias.name: The name returned by the mount configured via the `mount_accessor` parameter
        /// - entity.name: The name configured for the Entity
        /// - alias.metadata.`&lt;key&gt;`: The value of the Alias's metadata parameter
        /// - entity.metadata.`&lt;key&gt;`: The value of the Entity's metadata parameter
        /// </summary>
        [Input("usernameFormat")]
        public Input<string>? UsernameFormat { get; set; }

        public MfaDuoArgs()
        {
        }
        public static new MfaDuoArgs Empty => new MfaDuoArgs();
    }

    public sealed class MfaDuoState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// `(string: &lt;required&gt;)` - API hostname for Duo.
        /// </summary>
        [Input("apiHostname")]
        public Input<string>? ApiHostname { get; set; }

        [Input("integrationKey")]
        private Input<string>? _integrationKey;

        /// <summary>
        /// `(string: &lt;required&gt;)` - Integration key for Duo.
        /// </summary>
        public Input<string>? IntegrationKey
        {
            get => _integrationKey;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _integrationKey = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// `(string: &lt;required&gt;)` - The mount to tie this method to for use in automatic mappings. The mapping will use the Name field of Aliases associated with this mount as the username in the mapping.
        /// </summary>
        [Input("mountAccessor")]
        public Input<string>? MountAccessor { get; set; }

        /// <summary>
        /// `(string: &lt;required&gt;)` – Name of the MFA method.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The namespace to provision the resource in.
        /// The value should not contain leading or trailing forward slashes.
        /// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
        /// *Available only for Vault Enterprise*.
        /// </summary>
        [Input("namespace")]
        public Input<string>? Namespace { get; set; }

        /// <summary>
        /// `(string)` - Push information for Duo.
        /// </summary>
        [Input("pushInfo")]
        public Input<string>? PushInfo { get; set; }

        [Input("secretKey")]
        private Input<string>? _secretKey;

        /// <summary>
        /// `(string: &lt;required&gt;)` - Secret key for Duo.
        /// </summary>
        public Input<string>? SecretKey
        {
            get => _secretKey;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _secretKey = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// `(string)` - A format string for mapping Identity names to MFA method names. Values to substitute should be placed in `{{}}`. For example, `"{{alias.name}}@example.com"`. If blank, the Alias's Name field will be used as-is. Currently-supported mappings:
        /// - alias.name: The name returned by the mount configured via the `mount_accessor` parameter
        /// - entity.name: The name configured for the Entity
        /// - alias.metadata.`&lt;key&gt;`: The value of the Alias's metadata parameter
        /// - entity.metadata.`&lt;key&gt;`: The value of the Entity's metadata parameter
        /// </summary>
        [Input("usernameFormat")]
        public Input<string>? UsernameFormat { get; set; }

        public MfaDuoState()
        {
        }
        public static new MfaDuoState Empty => new MfaDuoState();
    }
}
