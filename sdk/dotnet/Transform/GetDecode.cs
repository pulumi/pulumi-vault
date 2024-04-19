// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vault.Transform
{
    public static class GetDecode
    {
        /// <summary>
        /// This data source supports the "/transform/decode/{role_name}" Vault endpoint.
        /// 
        /// It decodes the provided value using a named role.
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
        ///     var transform = new Vault.Mount("transform", new()
        ///     {
        ///         Path = "transform",
        ///         Type = "transform",
        ///     });
        /// 
        ///     var ccn_fpe = new Vault.Transform.Transformation("ccn-fpe", new()
        ///     {
        ///         Path = transform.Path,
        ///         Type = "fpe",
        ///         Template = "builtin/creditcardnumber",
        ///         TweakSource = "internal",
        ///         AllowedRoles = new[]
        ///         {
        ///             "payments",
        ///         },
        ///     });
        /// 
        ///     var payments = new Vault.Transform.Role("payments", new()
        ///     {
        ///         Path = ccn_fpe.Path,
        ///         Transformations = new[]
        ///         {
        ///             "ccn-fpe",
        ///         },
        ///     });
        /// 
        ///     var test = Vault.Transform.GetDecode.Invoke(new()
        ///     {
        ///         Path = payments.Path,
        ///         RoleName = "payments",
        ///         Value = "9300-3376-4943-8903",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Task<GetDecodeResult> InvokeAsync(GetDecodeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetDecodeResult>("vault:transform/getDecode:getDecode", args ?? new GetDecodeArgs(), options.WithDefaults());

        /// <summary>
        /// This data source supports the "/transform/decode/{role_name}" Vault endpoint.
        /// 
        /// It decodes the provided value using a named role.
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
        ///     var transform = new Vault.Mount("transform", new()
        ///     {
        ///         Path = "transform",
        ///         Type = "transform",
        ///     });
        /// 
        ///     var ccn_fpe = new Vault.Transform.Transformation("ccn-fpe", new()
        ///     {
        ///         Path = transform.Path,
        ///         Type = "fpe",
        ///         Template = "builtin/creditcardnumber",
        ///         TweakSource = "internal",
        ///         AllowedRoles = new[]
        ///         {
        ///             "payments",
        ///         },
        ///     });
        /// 
        ///     var payments = new Vault.Transform.Role("payments", new()
        ///     {
        ///         Path = ccn_fpe.Path,
        ///         Transformations = new[]
        ///         {
        ///             "ccn-fpe",
        ///         },
        ///     });
        /// 
        ///     var test = Vault.Transform.GetDecode.Invoke(new()
        ///     {
        ///         Path = payments.Path,
        ///         RoleName = "payments",
        ///         Value = "9300-3376-4943-8903",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetDecodeResult> Invoke(GetDecodeInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetDecodeResult>("vault:transform/getDecode:getDecode", args ?? new GetDecodeInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetDecodeArgs : global::Pulumi.InvokeArgs
    {
        [Input("batchInputs")]
        private List<ImmutableDictionary<string, object>>? _batchInputs;

        /// <summary>
        /// Specifies a list of items to be decoded in a single batch. If this parameter is set, the top-level parameters 'value', 'transformation' and 'tweak' will be ignored. Each batch item within the list can specify these parameters instead.
        /// </summary>
        public List<ImmutableDictionary<string, object>> BatchInputs
        {
            get => _batchInputs ?? (_batchInputs = new List<ImmutableDictionary<string, object>>());
            set => _batchInputs = value;
        }

        [Input("batchResults")]
        private List<ImmutableDictionary<string, object>>? _batchResults;

        /// <summary>
        /// The result of decoding a batch.
        /// </summary>
        public List<ImmutableDictionary<string, object>> BatchResults
        {
            get => _batchResults ?? (_batchResults = new List<ImmutableDictionary<string, object>>());
            set => _batchResults = value;
        }

        /// <summary>
        /// The result of decoding a value.
        /// </summary>
        [Input("decodedValue")]
        public string? DecodedValue { get; set; }

        /// <summary>
        /// The namespace of the target resource.
        /// The value should not contain leading or trailing forward slashes.
        /// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
        /// *Available only for Vault Enterprise*.
        /// </summary>
        [Input("namespace")]
        public string? Namespace { get; set; }

        /// <summary>
        /// Path to where the back-end is mounted within Vault.
        /// </summary>
        [Input("path", required: true)]
        public string Path { get; set; } = null!;

        /// <summary>
        /// The name of the role.
        /// </summary>
        [Input("roleName", required: true)]
        public string RoleName { get; set; } = null!;

        /// <summary>
        /// The transformation to perform. If no value is provided and the role contains a single transformation, this value will be inferred from the role.
        /// </summary>
        [Input("transformation")]
        public string? Transformation { get; set; }

        /// <summary>
        /// The tweak value to use. Only applicable for FPE transformations
        /// </summary>
        [Input("tweak")]
        public string? Tweak { get; set; }

        /// <summary>
        /// The value in which to decode.
        /// </summary>
        [Input("value")]
        public string? Value { get; set; }

        public GetDecodeArgs()
        {
        }
        public static new GetDecodeArgs Empty => new GetDecodeArgs();
    }

    public sealed class GetDecodeInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("batchInputs")]
        private InputList<ImmutableDictionary<string, object>>? _batchInputs;

        /// <summary>
        /// Specifies a list of items to be decoded in a single batch. If this parameter is set, the top-level parameters 'value', 'transformation' and 'tweak' will be ignored. Each batch item within the list can specify these parameters instead.
        /// </summary>
        public InputList<ImmutableDictionary<string, object>> BatchInputs
        {
            get => _batchInputs ?? (_batchInputs = new InputList<ImmutableDictionary<string, object>>());
            set => _batchInputs = value;
        }

        [Input("batchResults")]
        private InputList<ImmutableDictionary<string, object>>? _batchResults;

        /// <summary>
        /// The result of decoding a batch.
        /// </summary>
        public InputList<ImmutableDictionary<string, object>> BatchResults
        {
            get => _batchResults ?? (_batchResults = new InputList<ImmutableDictionary<string, object>>());
            set => _batchResults = value;
        }

        /// <summary>
        /// The result of decoding a value.
        /// </summary>
        [Input("decodedValue")]
        public Input<string>? DecodedValue { get; set; }

        /// <summary>
        /// The namespace of the target resource.
        /// The value should not contain leading or trailing forward slashes.
        /// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
        /// *Available only for Vault Enterprise*.
        /// </summary>
        [Input("namespace")]
        public Input<string>? Namespace { get; set; }

        /// <summary>
        /// Path to where the back-end is mounted within Vault.
        /// </summary>
        [Input("path", required: true)]
        public Input<string> Path { get; set; } = null!;

        /// <summary>
        /// The name of the role.
        /// </summary>
        [Input("roleName", required: true)]
        public Input<string> RoleName { get; set; } = null!;

        /// <summary>
        /// The transformation to perform. If no value is provided and the role contains a single transformation, this value will be inferred from the role.
        /// </summary>
        [Input("transformation")]
        public Input<string>? Transformation { get; set; }

        /// <summary>
        /// The tweak value to use. Only applicable for FPE transformations
        /// </summary>
        [Input("tweak")]
        public Input<string>? Tweak { get; set; }

        /// <summary>
        /// The value in which to decode.
        /// </summary>
        [Input("value")]
        public Input<string>? Value { get; set; }

        public GetDecodeInvokeArgs()
        {
        }
        public static new GetDecodeInvokeArgs Empty => new GetDecodeInvokeArgs();
    }


    [OutputType]
    public sealed class GetDecodeResult
    {
        public readonly ImmutableArray<ImmutableDictionary<string, object>> BatchInputs;
        public readonly ImmutableArray<ImmutableDictionary<string, object>> BatchResults;
        public readonly string DecodedValue;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string? Namespace;
        public readonly string Path;
        public readonly string RoleName;
        public readonly string? Transformation;
        public readonly string? Tweak;
        public readonly string? Value;

        [OutputConstructor]
        private GetDecodeResult(
            ImmutableArray<ImmutableDictionary<string, object>> batchInputs,

            ImmutableArray<ImmutableDictionary<string, object>> batchResults,

            string decodedValue,

            string id,

            string? @namespace,

            string path,

            string roleName,

            string? transformation,

            string? tweak,

            string? value)
        {
            BatchInputs = batchInputs;
            BatchResults = batchResults;
            DecodedValue = decodedValue;
            Id = id;
            Namespace = @namespace;
            Path = path;
            RoleName = roleName;
            Transformation = transformation;
            Tweak = tweak;
            Value = value;
        }
    }
}
