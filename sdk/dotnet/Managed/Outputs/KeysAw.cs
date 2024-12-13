// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vault.Managed.Outputs
{

    [OutputType]
    public sealed class KeysAw
    {
        /// <summary>
        /// The AWS access key to use
        /// </summary>
        public readonly string AccessKey;
        /// <summary>
        /// If no existing key can be found in the referenced backend, instructs Vault to generate a key within the backend
        /// </summary>
        public readonly bool? AllowGenerateKey;
        /// <summary>
        /// Controls the ability for Vault to replace through generation or importing a key into the configured backend even if a key is present, if set to false those operations are forbidden if a key exists.
        /// </summary>
        public readonly bool? AllowReplaceKey;
        /// <summary>
        /// Controls the ability for Vault to import a key to the configured backend, if 'false', those operations will be forbidden
        /// </summary>
        public readonly bool? AllowStoreKey;
        /// <summary>
        /// Allow usage from any mount point within the namespace if 'true'
        /// </summary>
        public readonly bool? AnyMount;
        /// <summary>
        /// The curve to use for an ECDSA key. Used when key_type is 'ECDSA'. Required if 'allow_generate_key' is true
        /// </summary>
        public readonly string? Curve;
        /// <summary>
        /// Used to specify a custom AWS endpoint
        /// </summary>
        public readonly string? Endpoint;
        /// <summary>
        /// The size in bits for an RSA key. This field is required when 'key_type' is 'RSA'
        /// </summary>
        public readonly string KeyBits;
        /// <summary>
        /// The type of key to use
        /// </summary>
        public readonly string KeyType;
        /// <summary>
        /// An identifier for the key
        /// </summary>
        public readonly string KmsKey;
        /// <summary>
        /// A unique lowercase name that serves as identifying the key
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The AWS region where the keys are stored (or will be stored)
        /// </summary>
        public readonly string? Region;
        /// <summary>
        /// The AWS secret key to use
        /// </summary>
        public readonly string SecretKey;
        /// <summary>
        /// ID of the managed key read from Vault
        /// </summary>
        public readonly string? Uuid;

        [OutputConstructor]
        private KeysAw(
            string accessKey,

            bool? allowGenerateKey,

            bool? allowReplaceKey,

            bool? allowStoreKey,

            bool? anyMount,

            string? curve,

            string? endpoint,

            string keyBits,

            string keyType,

            string kmsKey,

            string name,

            string? region,

            string secretKey,

            string? uuid)
        {
            AccessKey = accessKey;
            AllowGenerateKey = allowGenerateKey;
            AllowReplaceKey = allowReplaceKey;
            AllowStoreKey = allowStoreKey;
            AnyMount = anyMount;
            Curve = curve;
            Endpoint = endpoint;
            KeyBits = keyBits;
            KeyType = keyType;
            KmsKey = kmsKey;
            Name = name;
            Region = region;
            SecretKey = secretKey;
            Uuid = uuid;
        }
    }
}