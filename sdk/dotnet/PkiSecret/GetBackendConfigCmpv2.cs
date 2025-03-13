// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vault.PkiSecret
{
    public static class GetBackendConfigCmpv2
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
        ///     var pki = new Vault.Mount("pki", new()
        ///     {
        ///         Path = "pki",
        ///         Type = "pki",
        ///         Description = "PKI secret engine mount",
        ///     });
        /// 
        ///     var cmpv2Config = Vault.PkiSecret.GetBackendConfigCmpv2.Invoke(new()
        ///     {
        ///         Backend = pki.Path,
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Task<GetBackendConfigCmpv2Result> InvokeAsync(GetBackendConfigCmpv2Args args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetBackendConfigCmpv2Result>("vault:pkiSecret/getBackendConfigCmpv2:getBackendConfigCmpv2", args ?? new GetBackendConfigCmpv2Args(), options.WithDefaults());

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
        ///     var pki = new Vault.Mount("pki", new()
        ///     {
        ///         Path = "pki",
        ///         Type = "pki",
        ///         Description = "PKI secret engine mount",
        ///     });
        /// 
        ///     var cmpv2Config = Vault.PkiSecret.GetBackendConfigCmpv2.Invoke(new()
        ///     {
        ///         Backend = pki.Path,
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetBackendConfigCmpv2Result> Invoke(GetBackendConfigCmpv2InvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetBackendConfigCmpv2Result>("vault:pkiSecret/getBackendConfigCmpv2:getBackendConfigCmpv2", args ?? new GetBackendConfigCmpv2InvokeArgs(), options.WithDefaults());

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
        ///     var pki = new Vault.Mount("pki", new()
        ///     {
        ///         Path = "pki",
        ///         Type = "pki",
        ///         Description = "PKI secret engine mount",
        ///     });
        /// 
        ///     var cmpv2Config = Vault.PkiSecret.GetBackendConfigCmpv2.Invoke(new()
        ///     {
        ///         Backend = pki.Path,
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetBackendConfigCmpv2Result> Invoke(GetBackendConfigCmpv2InvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetBackendConfigCmpv2Result>("vault:pkiSecret/getBackendConfigCmpv2:getBackendConfigCmpv2", args ?? new GetBackendConfigCmpv2InvokeArgs(), options.WithDefaults());
    }


    public sealed class GetBackendConfigCmpv2Args : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The path to the PKI secret backend to
        /// read the CMPv2 configuration from, with no leading or trailing `/`s.
        /// 
        /// # Attributes Reference
        /// </summary>
        [Input("backend", required: true)]
        public string Backend { get; set; } = null!;

        [Input("disabledValidations")]
        private List<string>? _disabledValidations;

        /// <summary>
        /// A comma-separated list of validations not to perform on CMPv2 messages.
        /// </summary>
        public List<string> DisabledValidations
        {
            get => _disabledValidations ?? (_disabledValidations = new List<string>());
            set => _disabledValidations = value;
        }

        /// <summary>
        /// The namespace of the target resource.
        /// The value should not contain leading or trailing forward slashes.
        /// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
        /// *Available only for Vault Enterprise*.
        /// </summary>
        [Input("namespace")]
        public string? Namespace { get; set; }

        public GetBackendConfigCmpv2Args()
        {
        }
        public static new GetBackendConfigCmpv2Args Empty => new GetBackendConfigCmpv2Args();
    }

    public sealed class GetBackendConfigCmpv2InvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The path to the PKI secret backend to
        /// read the CMPv2 configuration from, with no leading or trailing `/`s.
        /// 
        /// # Attributes Reference
        /// </summary>
        [Input("backend", required: true)]
        public Input<string> Backend { get; set; } = null!;

        [Input("disabledValidations")]
        private InputList<string>? _disabledValidations;

        /// <summary>
        /// A comma-separated list of validations not to perform on CMPv2 messages.
        /// </summary>
        public InputList<string> DisabledValidations
        {
            get => _disabledValidations ?? (_disabledValidations = new InputList<string>());
            set => _disabledValidations = value;
        }

        /// <summary>
        /// The namespace of the target resource.
        /// The value should not contain leading or trailing forward slashes.
        /// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
        /// *Available only for Vault Enterprise*.
        /// </summary>
        [Input("namespace")]
        public Input<string>? Namespace { get; set; }

        public GetBackendConfigCmpv2InvokeArgs()
        {
        }
        public static new GetBackendConfigCmpv2InvokeArgs Empty => new GetBackendConfigCmpv2InvokeArgs();
    }


    [OutputType]
    public sealed class GetBackendConfigCmpv2Result
    {
        public readonly ImmutableArray<string> AuditFields;
        public readonly ImmutableArray<Outputs.GetBackendConfigCmpv2AuthenticatorResult> Authenticators;
        public readonly string Backend;
        public readonly string DefaultPathPolicy;
        public readonly ImmutableArray<string> DisabledValidations;
        public readonly bool EnableSentinelParsing;
        public readonly bool Enabled;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string LastUpdated;
        public readonly string? Namespace;

        [OutputConstructor]
        private GetBackendConfigCmpv2Result(
            ImmutableArray<string> auditFields,

            ImmutableArray<Outputs.GetBackendConfigCmpv2AuthenticatorResult> authenticators,

            string backend,

            string defaultPathPolicy,

            ImmutableArray<string> disabledValidations,

            bool enableSentinelParsing,

            bool enabled,

            string id,

            string lastUpdated,

            string? @namespace)
        {
            AuditFields = auditFields;
            Authenticators = authenticators;
            Backend = backend;
            DefaultPathPolicy = defaultPathPolicy;
            DisabledValidations = disabledValidations;
            EnableSentinelParsing = enableSentinelParsing;
            Enabled = enabled;
            Id = id;
            LastUpdated = lastUpdated;
            Namespace = @namespace;
        }
    }
}
