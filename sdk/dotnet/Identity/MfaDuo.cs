// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vault.Identity
{
    /// <summary>
    /// Resource for configuring the duo MFA method.
    /// 
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
    ///     var example = new Vault.Identity.MfaDuo("example", new()
    ///     {
    ///         ApiHostname = "api-xxxxxxxx.duosecurity.com",
    ///         SecretKey = "secret-key",
    ///         IntegrationKey = "secret-int-key",
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ## Import
    /// 
    /// Resource can be imported using its `uuid` field, e.g.
    /// 
    /// ```sh
    /// $ pulumi import vault:identity/mfaDuo:MfaDuo example 0d89c36a-4ff5-4d70-8749-bb6a5598aeec
    /// ```
    /// </summary>
    [VaultResourceType("vault:identity/mfaDuo:MfaDuo")]
    public partial class MfaDuo : global::Pulumi.CustomResource
    {
        /// <summary>
        /// API hostname for Duo
        /// </summary>
        [Output("apiHostname")]
        public Output<string> ApiHostname { get; private set; } = null!;

        /// <summary>
        /// Integration key for Duo
        /// </summary>
        [Output("integrationKey")]
        public Output<string> IntegrationKey { get; private set; } = null!;

        /// <summary>
        /// Method ID.
        /// </summary>
        [Output("methodId")]
        public Output<string> MethodId { get; private set; } = null!;

        /// <summary>
        /// Mount accessor.
        /// </summary>
        [Output("mountAccessor")]
        public Output<string> MountAccessor { get; private set; } = null!;

        /// <summary>
        /// Method name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Target namespace. (requires Enterprise)
        /// </summary>
        [Output("namespace")]
        public Output<string?> Namespace { get; private set; } = null!;

        /// <summary>
        /// Method's namespace ID.
        /// </summary>
        [Output("namespaceId")]
        public Output<string> NamespaceId { get; private set; } = null!;

        /// <summary>
        /// Method's namespace path.
        /// </summary>
        [Output("namespacePath")]
        public Output<string> NamespacePath { get; private set; } = null!;

        /// <summary>
        /// Push information for Duo.
        /// </summary>
        [Output("pushInfo")]
        public Output<string?> PushInfo { get; private set; } = null!;

        /// <summary>
        /// Secret key for Duo
        /// </summary>
        [Output("secretKey")]
        public Output<string> SecretKey { get; private set; } = null!;

        /// <summary>
        /// MFA type.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;

        /// <summary>
        /// Require passcode upon MFA validation.
        /// </summary>
        [Output("usePasscode")]
        public Output<bool?> UsePasscode { get; private set; } = null!;

        /// <summary>
        /// A template string for mapping Identity names to MFA methods.
        /// </summary>
        [Output("usernameFormat")]
        public Output<string?> UsernameFormat { get; private set; } = null!;

        /// <summary>
        /// Resource UUID.
        /// </summary>
        [Output("uuid")]
        public Output<string> Uuid { get; private set; } = null!;


        /// <summary>
        /// Create a MfaDuo resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public MfaDuo(string name, MfaDuoArgs args, CustomResourceOptions? options = null)
            : base("vault:identity/mfaDuo:MfaDuo", name, args ?? new MfaDuoArgs(), MakeResourceOptions(options, ""))
        {
        }

        private MfaDuo(string name, Input<string> id, MfaDuoState? state = null, CustomResourceOptions? options = null)
            : base("vault:identity/mfaDuo:MfaDuo", name, state, MakeResourceOptions(options, id))
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
        /// API hostname for Duo
        /// </summary>
        [Input("apiHostname", required: true)]
        public Input<string> ApiHostname { get; set; } = null!;

        [Input("integrationKey", required: true)]
        private Input<string>? _integrationKey;

        /// <summary>
        /// Integration key for Duo
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
        /// Target namespace. (requires Enterprise)
        /// </summary>
        [Input("namespace")]
        public Input<string>? Namespace { get; set; }

        /// <summary>
        /// Push information for Duo.
        /// </summary>
        [Input("pushInfo")]
        public Input<string>? PushInfo { get; set; }

        [Input("secretKey", required: true)]
        private Input<string>? _secretKey;

        /// <summary>
        /// Secret key for Duo
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
        /// Require passcode upon MFA validation.
        /// </summary>
        [Input("usePasscode")]
        public Input<bool>? UsePasscode { get; set; }

        /// <summary>
        /// A template string for mapping Identity names to MFA methods.
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
        /// API hostname for Duo
        /// </summary>
        [Input("apiHostname")]
        public Input<string>? ApiHostname { get; set; }

        [Input("integrationKey")]
        private Input<string>? _integrationKey;

        /// <summary>
        /// Integration key for Duo
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
        /// Method ID.
        /// </summary>
        [Input("methodId")]
        public Input<string>? MethodId { get; set; }

        /// <summary>
        /// Mount accessor.
        /// </summary>
        [Input("mountAccessor")]
        public Input<string>? MountAccessor { get; set; }

        /// <summary>
        /// Method name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Target namespace. (requires Enterprise)
        /// </summary>
        [Input("namespace")]
        public Input<string>? Namespace { get; set; }

        /// <summary>
        /// Method's namespace ID.
        /// </summary>
        [Input("namespaceId")]
        public Input<string>? NamespaceId { get; set; }

        /// <summary>
        /// Method's namespace path.
        /// </summary>
        [Input("namespacePath")]
        public Input<string>? NamespacePath { get; set; }

        /// <summary>
        /// Push information for Duo.
        /// </summary>
        [Input("pushInfo")]
        public Input<string>? PushInfo { get; set; }

        [Input("secretKey")]
        private Input<string>? _secretKey;

        /// <summary>
        /// Secret key for Duo
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
        /// MFA type.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        /// <summary>
        /// Require passcode upon MFA validation.
        /// </summary>
        [Input("usePasscode")]
        public Input<bool>? UsePasscode { get; set; }

        /// <summary>
        /// A template string for mapping Identity names to MFA methods.
        /// </summary>
        [Input("usernameFormat")]
        public Input<string>? UsernameFormat { get; set; }

        /// <summary>
        /// Resource UUID.
        /// </summary>
        [Input("uuid")]
        public Input<string>? Uuid { get; set; }

        public MfaDuoState()
        {
        }
        public static new MfaDuoState Empty => new MfaDuoState();
    }
}
