// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vault.Managed.Inputs
{

    public sealed class KeysAwArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The AWS access key to use.
        /// </summary>
        [Input("accessKey", required: true)]
        public Input<string> AccessKey { get; set; } = null!;

        /// <summary>
        /// If no existing key can be found in 
        /// the referenced backend, instructs Vault to generate a key within the backend.
        /// </summary>
        [Input("allowGenerateKey")]
        public Input<bool>? AllowGenerateKey { get; set; }

        /// <summary>
        /// Controls the ability for Vault to replace through
        /// generation or importing a key into the configured backend even
        /// if a key is present, if set to `false` those operations are forbidden
        /// if a key exists.
        /// </summary>
        [Input("allowReplaceKey")]
        public Input<bool>? AllowReplaceKey { get; set; }

        /// <summary>
        /// Controls the ability for Vault to import a key to the
        /// configured backend, if `false`, those operations will be forbidden.
        /// </summary>
        [Input("allowStoreKey")]
        public Input<bool>? AllowStoreKey { get; set; }

        /// <summary>
        /// If `true`, allows usage from any mount point within the
        /// namespace.
        /// </summary>
        [Input("anyMount")]
        public Input<bool>? AnyMount { get; set; }

        /// <summary>
        /// The curve to use for an ECDSA key. Used when `key_type` 
        /// is `ECDSA`. Required if `allow_generate_key` is `true`.
        /// </summary>
        [Input("curve")]
        public Input<string>? Curve { get; set; }

        /// <summary>
        /// Used to specify a custom AWS endpoint.
        /// </summary>
        [Input("endpoint")]
        public Input<string>? Endpoint { get; set; }

        /// <summary>
        /// The size in bits for an RSA key.
        /// </summary>
        [Input("keyBits", required: true)]
        public Input<string> KeyBits { get; set; } = null!;

        /// <summary>
        /// The type of key to use.
        /// </summary>
        [Input("keyType", required: true)]
        public Input<string> KeyType { get; set; } = null!;

        /// <summary>
        /// An identifier for the key.
        /// </summary>
        [Input("kmsKey", required: true)]
        public Input<string> KmsKey { get; set; } = null!;

        /// <summary>
        /// A unique lowercase name that serves as identifying the key.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// The AWS region where the keys are stored (or will be stored).
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// The AWS access key to use.
        /// </summary>
        [Input("secretKey", required: true)]
        public Input<string> SecretKey { get; set; } = null!;

        [Input("uuid")]
        public Input<string>? Uuid { get; set; }

        public KeysAwArgs()
        {
        }
        public static new KeysAwArgs Empty => new KeysAwArgs();
    }
}
