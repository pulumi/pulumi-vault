// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vault.Transit
{
    public static class GetDecrypt
    {
        /// <summary>
        /// This is a data source which can be used to decrypt ciphertext using a Vault Transit key.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Vault = Pulumi.Vault;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var test = Vault.Transit.GetDecrypt.Invoke(new()
        ///     {
        ///         Backend = "transit",
        ///         Ciphertext = "vault:v1:S3GtnJ5GUNCWV+/pdL9+g1Feu/nzAv+RlmTmE91Tu0rBkeIU8MEb2nSspC/1IQ==",
        ///         Key = "test",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetDecryptResult> InvokeAsync(GetDecryptArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetDecryptResult>("vault:transit/getDecrypt:getDecrypt", args ?? new GetDecryptArgs(), options.WithDefaults());

        /// <summary>
        /// This is a data source which can be used to decrypt ciphertext using a Vault Transit key.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Vault = Pulumi.Vault;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var test = Vault.Transit.GetDecrypt.Invoke(new()
        ///     {
        ///         Backend = "transit",
        ///         Ciphertext = "vault:v1:S3GtnJ5GUNCWV+/pdL9+g1Feu/nzAv+RlmTmE91Tu0rBkeIU8MEb2nSspC/1IQ==",
        ///         Key = "test",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetDecryptResult> Invoke(GetDecryptInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetDecryptResult>("vault:transit/getDecrypt:getDecrypt", args ?? new GetDecryptInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetDecryptArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The path the transit secret backend is mounted at, with no leading or trailing `/`.
        /// </summary>
        [Input("backend", required: true)]
        public string Backend { get; set; } = null!;

        /// <summary>
        /// Ciphertext to be decoded.
        /// </summary>
        [Input("ciphertext", required: true)]
        public string Ciphertext { get; set; } = null!;

        /// <summary>
        /// Context for key derivation. This is required if key derivation is enabled for this key.
        /// </summary>
        [Input("context")]
        public string? Context { get; set; }

        /// <summary>
        /// Specifies the name of the transit key to decrypt against.
        /// </summary>
        [Input("key", required: true)]
        public string Key { get; set; } = null!;

        [Input("namespace")]
        public string? Namespace { get; set; }

        public GetDecryptArgs()
        {
        }
        public static new GetDecryptArgs Empty => new GetDecryptArgs();
    }

    public sealed class GetDecryptInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The path the transit secret backend is mounted at, with no leading or trailing `/`.
        /// </summary>
        [Input("backend", required: true)]
        public Input<string> Backend { get; set; } = null!;

        /// <summary>
        /// Ciphertext to be decoded.
        /// </summary>
        [Input("ciphertext", required: true)]
        public Input<string> Ciphertext { get; set; } = null!;

        /// <summary>
        /// Context for key derivation. This is required if key derivation is enabled for this key.
        /// </summary>
        [Input("context")]
        public Input<string>? Context { get; set; }

        /// <summary>
        /// Specifies the name of the transit key to decrypt against.
        /// </summary>
        [Input("key", required: true)]
        public Input<string> Key { get; set; } = null!;

        [Input("namespace")]
        public Input<string>? Namespace { get; set; }

        public GetDecryptInvokeArgs()
        {
        }
        public static new GetDecryptInvokeArgs Empty => new GetDecryptInvokeArgs();
    }


    [OutputType]
    public sealed class GetDecryptResult
    {
        public readonly string Backend;
        public readonly string Ciphertext;
        public readonly string? Context;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string Key;
        public readonly string? Namespace;
        /// <summary>
        /// Decrypted plaintext returned from Vault
        /// </summary>
        public readonly string Plaintext;

        [OutputConstructor]
        private GetDecryptResult(
            string backend,

            string ciphertext,

            string? context,

            string id,

            string key,

            string? @namespace,

            string plaintext)
        {
            Backend = backend;
            Ciphertext = ciphertext;
            Context = context;
            Id = id;
            Key = key;
            Namespace = @namespace;
            Plaintext = plaintext;
        }
    }
}
