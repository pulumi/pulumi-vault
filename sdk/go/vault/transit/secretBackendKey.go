// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package transit

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Creates an Encryption Keyring on a Transit Secret Backend for Vault.
//
//
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-vault/blob/master/website/docs/r/transit_secret_backend_key.html.md.
type SecretBackendKey struct {
	pulumi.CustomResourceState

	// Enables taking backup of entire keyring in the plaintext format. Once set, this cannot be disabled.
	// * Refer to Vault API documentation on key backups for more information: [Backup Key](https://www.vaultproject.io/api/secret/transit/index.html#backup-key)
	AllowPlaintextBackup pulumi.BoolPtrOutput `pulumi:"allowPlaintextBackup"`
	// The path the transit secret backend is mounted at, with no leading or trailing `/`s.
	Backend pulumi.StringOutput `pulumi:"backend"`
	// Whether or not to support convergent encryption, where the same plaintext creates the same ciphertext. This requires `derived` to be set to `true`.
	ConvergentEncryption pulumi.BoolPtrOutput `pulumi:"convergentEncryption"`
	// Specifies if the key is allowed to be deleted.
	DeletionAllowed pulumi.BoolPtrOutput `pulumi:"deletionAllowed"`
	// Specifies if key derivation is to be used. If enabled, all encrypt/decrypt requests to this key must provide a context which is used for key derivation.
	Derived pulumi.BoolPtrOutput `pulumi:"derived"`
	// Enables keys to be exportable. This allows for all valid private keys in the keyring to be exported. Once set, this cannot be disabled.
	Exportable pulumi.BoolPtrOutput `pulumi:"exportable"`
	// List of key versions in the keyring. This attribute is zero-indexed and will contain a map of values depending on the `type` of the encryption key.
	// * for key types `aes256-gcm96` and `chacha20-poly1305`, each key version will be a map of a single value `id` which is just a hash of the key's metadata.
	// * for key types `ed25519`, `ecdsa-p256`, `rsa-2048` and `rsa-4096`, each key version will be a map of the following:
	Keys pulumi.MapArrayOutput `pulumi:"keys"`
	// Latest key version available. This value is 1-indexed, so if `latestVersion` is `1`, then the key's information can be referenced from `keys` by selecting element `0`
	LatestVersion pulumi.IntOutput `pulumi:"latestVersion"`
	// Minimum key version available for use. If keys have been archived by increasing `minDecryptionVersion`, this attribute will reflect that change.
	MinAvailableVersion pulumi.IntOutput `pulumi:"minAvailableVersion"`
	// Minimum key version to use for decryption.
	MinDecryptionVersion pulumi.IntPtrOutput `pulumi:"minDecryptionVersion"`
	// Minimum key version to use for encryption
	MinEncryptionVersion pulumi.IntPtrOutput `pulumi:"minEncryptionVersion"`
	// The name to identify this key within the backend. Must be unique within the backend.
	Name pulumi.StringOutput `pulumi:"name"`
	// Whether or not the key supports decryption, based on key type.
	SupportsDecryption pulumi.BoolOutput `pulumi:"supportsDecryption"`
	// Whether or not the key supports derivation, based on key type.
	SupportsDerivation pulumi.BoolOutput `pulumi:"supportsDerivation"`
	// Whether or not the key supports encryption, based on key type.
	SupportsEncryption pulumi.BoolOutput `pulumi:"supportsEncryption"`
	// Whether or not the key supports signing, based on key type.
	SupportsSigning pulumi.BoolOutput `pulumi:"supportsSigning"`
	// Specifies the type of key to create. The currently-supported types are: `aes256-gcm96` (default), `chacha20-poly1305`, `ed25519`, `ecdsa-p256`, `rsa-2048` and `rsa-4096`.
	// * Refer to the Vault documentation on transit key types for more information: [Key Types](https://www.vaultproject.io/docs/secrets/transit/index.html#key-types)
	Type pulumi.StringPtrOutput `pulumi:"type"`
}

// NewSecretBackendKey registers a new resource with the given unique name, arguments, and options.
func NewSecretBackendKey(ctx *pulumi.Context,
	name string, args *SecretBackendKeyArgs, opts ...pulumi.ResourceOption) (*SecretBackendKey, error) {
	if args == nil || args.Backend == nil {
		return nil, errors.New("missing required argument 'Backend'")
	}
	if args == nil {
		args = &SecretBackendKeyArgs{}
	}
	var resource SecretBackendKey
	err := ctx.RegisterResource("vault:transit/secretBackendKey:SecretBackendKey", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSecretBackendKey gets an existing SecretBackendKey resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSecretBackendKey(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SecretBackendKeyState, opts ...pulumi.ResourceOption) (*SecretBackendKey, error) {
	var resource SecretBackendKey
	err := ctx.ReadResource("vault:transit/secretBackendKey:SecretBackendKey", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SecretBackendKey resources.
type secretBackendKeyState struct {
	// Enables taking backup of entire keyring in the plaintext format. Once set, this cannot be disabled.
	// * Refer to Vault API documentation on key backups for more information: [Backup Key](https://www.vaultproject.io/api/secret/transit/index.html#backup-key)
	AllowPlaintextBackup *bool `pulumi:"allowPlaintextBackup"`
	// The path the transit secret backend is mounted at, with no leading or trailing `/`s.
	Backend *string `pulumi:"backend"`
	// Whether or not to support convergent encryption, where the same plaintext creates the same ciphertext. This requires `derived` to be set to `true`.
	ConvergentEncryption *bool `pulumi:"convergentEncryption"`
	// Specifies if the key is allowed to be deleted.
	DeletionAllowed *bool `pulumi:"deletionAllowed"`
	// Specifies if key derivation is to be used. If enabled, all encrypt/decrypt requests to this key must provide a context which is used for key derivation.
	Derived *bool `pulumi:"derived"`
	// Enables keys to be exportable. This allows for all valid private keys in the keyring to be exported. Once set, this cannot be disabled.
	Exportable *bool `pulumi:"exportable"`
	// List of key versions in the keyring. This attribute is zero-indexed and will contain a map of values depending on the `type` of the encryption key.
	// * for key types `aes256-gcm96` and `chacha20-poly1305`, each key version will be a map of a single value `id` which is just a hash of the key's metadata.
	// * for key types `ed25519`, `ecdsa-p256`, `rsa-2048` and `rsa-4096`, each key version will be a map of the following:
	Keys []map[string]interface{} `pulumi:"keys"`
	// Latest key version available. This value is 1-indexed, so if `latestVersion` is `1`, then the key's information can be referenced from `keys` by selecting element `0`
	LatestVersion *int `pulumi:"latestVersion"`
	// Minimum key version available for use. If keys have been archived by increasing `minDecryptionVersion`, this attribute will reflect that change.
	MinAvailableVersion *int `pulumi:"minAvailableVersion"`
	// Minimum key version to use for decryption.
	MinDecryptionVersion *int `pulumi:"minDecryptionVersion"`
	// Minimum key version to use for encryption
	MinEncryptionVersion *int `pulumi:"minEncryptionVersion"`
	// The name to identify this key within the backend. Must be unique within the backend.
	Name *string `pulumi:"name"`
	// Whether or not the key supports decryption, based on key type.
	SupportsDecryption *bool `pulumi:"supportsDecryption"`
	// Whether or not the key supports derivation, based on key type.
	SupportsDerivation *bool `pulumi:"supportsDerivation"`
	// Whether or not the key supports encryption, based on key type.
	SupportsEncryption *bool `pulumi:"supportsEncryption"`
	// Whether or not the key supports signing, based on key type.
	SupportsSigning *bool `pulumi:"supportsSigning"`
	// Specifies the type of key to create. The currently-supported types are: `aes256-gcm96` (default), `chacha20-poly1305`, `ed25519`, `ecdsa-p256`, `rsa-2048` and `rsa-4096`.
	// * Refer to the Vault documentation on transit key types for more information: [Key Types](https://www.vaultproject.io/docs/secrets/transit/index.html#key-types)
	Type *string `pulumi:"type"`
}

type SecretBackendKeyState struct {
	// Enables taking backup of entire keyring in the plaintext format. Once set, this cannot be disabled.
	// * Refer to Vault API documentation on key backups for more information: [Backup Key](https://www.vaultproject.io/api/secret/transit/index.html#backup-key)
	AllowPlaintextBackup pulumi.BoolPtrInput
	// The path the transit secret backend is mounted at, with no leading or trailing `/`s.
	Backend pulumi.StringPtrInput
	// Whether or not to support convergent encryption, where the same plaintext creates the same ciphertext. This requires `derived` to be set to `true`.
	ConvergentEncryption pulumi.BoolPtrInput
	// Specifies if the key is allowed to be deleted.
	DeletionAllowed pulumi.BoolPtrInput
	// Specifies if key derivation is to be used. If enabled, all encrypt/decrypt requests to this key must provide a context which is used for key derivation.
	Derived pulumi.BoolPtrInput
	// Enables keys to be exportable. This allows for all valid private keys in the keyring to be exported. Once set, this cannot be disabled.
	Exportable pulumi.BoolPtrInput
	// List of key versions in the keyring. This attribute is zero-indexed and will contain a map of values depending on the `type` of the encryption key.
	// * for key types `aes256-gcm96` and `chacha20-poly1305`, each key version will be a map of a single value `id` which is just a hash of the key's metadata.
	// * for key types `ed25519`, `ecdsa-p256`, `rsa-2048` and `rsa-4096`, each key version will be a map of the following:
	Keys pulumi.MapArrayInput
	// Latest key version available. This value is 1-indexed, so if `latestVersion` is `1`, then the key's information can be referenced from `keys` by selecting element `0`
	LatestVersion pulumi.IntPtrInput
	// Minimum key version available for use. If keys have been archived by increasing `minDecryptionVersion`, this attribute will reflect that change.
	MinAvailableVersion pulumi.IntPtrInput
	// Minimum key version to use for decryption.
	MinDecryptionVersion pulumi.IntPtrInput
	// Minimum key version to use for encryption
	MinEncryptionVersion pulumi.IntPtrInput
	// The name to identify this key within the backend. Must be unique within the backend.
	Name pulumi.StringPtrInput
	// Whether or not the key supports decryption, based on key type.
	SupportsDecryption pulumi.BoolPtrInput
	// Whether or not the key supports derivation, based on key type.
	SupportsDerivation pulumi.BoolPtrInput
	// Whether or not the key supports encryption, based on key type.
	SupportsEncryption pulumi.BoolPtrInput
	// Whether or not the key supports signing, based on key type.
	SupportsSigning pulumi.BoolPtrInput
	// Specifies the type of key to create. The currently-supported types are: `aes256-gcm96` (default), `chacha20-poly1305`, `ed25519`, `ecdsa-p256`, `rsa-2048` and `rsa-4096`.
	// * Refer to the Vault documentation on transit key types for more information: [Key Types](https://www.vaultproject.io/docs/secrets/transit/index.html#key-types)
	Type pulumi.StringPtrInput
}

func (SecretBackendKeyState) ElementType() reflect.Type {
	return reflect.TypeOf((*secretBackendKeyState)(nil)).Elem()
}

type secretBackendKeyArgs struct {
	// Enables taking backup of entire keyring in the plaintext format. Once set, this cannot be disabled.
	// * Refer to Vault API documentation on key backups for more information: [Backup Key](https://www.vaultproject.io/api/secret/transit/index.html#backup-key)
	AllowPlaintextBackup *bool `pulumi:"allowPlaintextBackup"`
	// The path the transit secret backend is mounted at, with no leading or trailing `/`s.
	Backend string `pulumi:"backend"`
	// Whether or not to support convergent encryption, where the same plaintext creates the same ciphertext. This requires `derived` to be set to `true`.
	ConvergentEncryption *bool `pulumi:"convergentEncryption"`
	// Specifies if the key is allowed to be deleted.
	DeletionAllowed *bool `pulumi:"deletionAllowed"`
	// Specifies if key derivation is to be used. If enabled, all encrypt/decrypt requests to this key must provide a context which is used for key derivation.
	Derived *bool `pulumi:"derived"`
	// Enables keys to be exportable. This allows for all valid private keys in the keyring to be exported. Once set, this cannot be disabled.
	Exportable *bool `pulumi:"exportable"`
	// Minimum key version to use for decryption.
	MinDecryptionVersion *int `pulumi:"minDecryptionVersion"`
	// Minimum key version to use for encryption
	MinEncryptionVersion *int `pulumi:"minEncryptionVersion"`
	// The name to identify this key within the backend. Must be unique within the backend.
	Name *string `pulumi:"name"`
	// Specifies the type of key to create. The currently-supported types are: `aes256-gcm96` (default), `chacha20-poly1305`, `ed25519`, `ecdsa-p256`, `rsa-2048` and `rsa-4096`.
	// * Refer to the Vault documentation on transit key types for more information: [Key Types](https://www.vaultproject.io/docs/secrets/transit/index.html#key-types)
	Type *string `pulumi:"type"`
}

// The set of arguments for constructing a SecretBackendKey resource.
type SecretBackendKeyArgs struct {
	// Enables taking backup of entire keyring in the plaintext format. Once set, this cannot be disabled.
	// * Refer to Vault API documentation on key backups for more information: [Backup Key](https://www.vaultproject.io/api/secret/transit/index.html#backup-key)
	AllowPlaintextBackup pulumi.BoolPtrInput
	// The path the transit secret backend is mounted at, with no leading or trailing `/`s.
	Backend pulumi.StringInput
	// Whether or not to support convergent encryption, where the same plaintext creates the same ciphertext. This requires `derived` to be set to `true`.
	ConvergentEncryption pulumi.BoolPtrInput
	// Specifies if the key is allowed to be deleted.
	DeletionAllowed pulumi.BoolPtrInput
	// Specifies if key derivation is to be used. If enabled, all encrypt/decrypt requests to this key must provide a context which is used for key derivation.
	Derived pulumi.BoolPtrInput
	// Enables keys to be exportable. This allows for all valid private keys in the keyring to be exported. Once set, this cannot be disabled.
	Exportable pulumi.BoolPtrInput
	// Minimum key version to use for decryption.
	MinDecryptionVersion pulumi.IntPtrInput
	// Minimum key version to use for encryption
	MinEncryptionVersion pulumi.IntPtrInput
	// The name to identify this key within the backend. Must be unique within the backend.
	Name pulumi.StringPtrInput
	// Specifies the type of key to create. The currently-supported types are: `aes256-gcm96` (default), `chacha20-poly1305`, `ed25519`, `ecdsa-p256`, `rsa-2048` and `rsa-4096`.
	// * Refer to the Vault documentation on transit key types for more information: [Key Types](https://www.vaultproject.io/docs/secrets/transit/index.html#key-types)
	Type pulumi.StringPtrInput
}

func (SecretBackendKeyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*secretBackendKeyArgs)(nil)).Elem()
}
