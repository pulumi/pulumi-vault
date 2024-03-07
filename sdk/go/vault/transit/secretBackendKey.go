// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package transit

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-vault/sdk/v5/go/vault/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates an Encryption Keyring on a Transit Secret Backend for Vault.
//
// ## Example Usage
//
// <!--Start PulumiCodeChooser -->
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-vault/sdk/v5/go/vault"
//	"github.com/pulumi/pulumi-vault/sdk/v5/go/vault/transit"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			transit, err := vault.NewMount(ctx, "transit", &vault.MountArgs{
//				Path:                   pulumi.String("transit"),
//				Type:                   pulumi.String("transit"),
//				Description:            pulumi.String("Example description"),
//				DefaultLeaseTtlSeconds: pulumi.Int(3600),
//				MaxLeaseTtlSeconds:     pulumi.Int(86400),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = transit.NewSecretBackendKey(ctx, "key", &transit.SecretBackendKeyArgs{
//				Backend: transit.Path,
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// <!--End PulumiCodeChooser -->
//
// ## Deprecations
//
// * `autoRotateInterval` - Replaced by `autoRotatePeriod`.
//
// ## Import
//
// Transit secret backend keys can be imported using the `path`, e.g.
//
// ```sh
// $ pulumi import vault:transit/secretBackendKey:SecretBackendKey key transit/keys/my_key
// ```
type SecretBackendKey struct {
	pulumi.CustomResourceState

	// Enables taking backup of entire keyring in the plaintext format. Once set, this cannot be disabled.
	// * Refer to Vault API documentation on key backups for more information: [Backup Key](https://www.vaultproject.io/api-docs/secret/transit#backup-key)
	AllowPlaintextBackup pulumi.BoolPtrOutput `pulumi:"allowPlaintextBackup"`
	// Amount of time the key should live before being automatically rotated. A value of 0 disables automatic rotation for the
	// key.
	//
	// Deprecated: Use auto_rotate_period instead
	AutoRotateInterval pulumi.IntOutput `pulumi:"autoRotateInterval"`
	// Amount of seconds the key should live before being automatically rotated.
	// A value of 0 disables automatic rotation for the key.
	AutoRotatePeriod pulumi.IntOutput `pulumi:"autoRotatePeriod"`
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
	// The key size in bytes for algorithms that allow variable key sizes. Currently only applicable to HMAC, where it must be between 32 and 512 bytes.
	KeySize pulumi.IntPtrOutput `pulumi:"keySize"`
	// List of key versions in the keyring. This attribute is zero-indexed and will contain a map of values depending on the `type` of the encryption key.
	// * for key types `aes128-gcm96`, `aes256-gcm96` and `chacha20-poly1305`, each key version will be a map of a single value `id` which is just a hash of the key's metadata.
	// * for key types `ed25519`, `ecdsa-p256`, `ecdsa-p384`, `ecdsa-p521`, `rsa-2048`, `rsa-3072` and `rsa-4096`, each key version will be a map of the following:
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
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
	// *Available only for Vault Enterprise*.
	Namespace pulumi.StringPtrOutput `pulumi:"namespace"`
	// Whether or not the key supports decryption, based on key type.
	SupportsDecryption pulumi.BoolOutput `pulumi:"supportsDecryption"`
	// Whether or not the key supports derivation, based on key type.
	SupportsDerivation pulumi.BoolOutput `pulumi:"supportsDerivation"`
	// Whether or not the key supports encryption, based on key type.
	SupportsEncryption pulumi.BoolOutput `pulumi:"supportsEncryption"`
	// Whether or not the key supports signing, based on key type.
	SupportsSigning pulumi.BoolOutput `pulumi:"supportsSigning"`
	// Specifies the type of key to create. The currently-supported types are: `aes128-gcm96`, `aes256-gcm96` (default), `chacha20-poly1305`, `ed25519`, `ecdsa-p256`, `ecdsa-p384`, `ecdsa-p521`, `hmac`, `rsa-2048`, `rsa-3072` and `rsa-4096`.
	// * Refer to the Vault documentation on transit key types for more information: [Key Types](https://www.vaultproject.io/docs/secrets/transit#key-types)
	Type pulumi.StringPtrOutput `pulumi:"type"`
}

// NewSecretBackendKey registers a new resource with the given unique name, arguments, and options.
func NewSecretBackendKey(ctx *pulumi.Context,
	name string, args *SecretBackendKeyArgs, opts ...pulumi.ResourceOption) (*SecretBackendKey, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Backend == nil {
		return nil, errors.New("invalid value for required argument 'Backend'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
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
	// * Refer to Vault API documentation on key backups for more information: [Backup Key](https://www.vaultproject.io/api-docs/secret/transit#backup-key)
	AllowPlaintextBackup *bool `pulumi:"allowPlaintextBackup"`
	// Amount of time the key should live before being automatically rotated. A value of 0 disables automatic rotation for the
	// key.
	//
	// Deprecated: Use auto_rotate_period instead
	AutoRotateInterval *int `pulumi:"autoRotateInterval"`
	// Amount of seconds the key should live before being automatically rotated.
	// A value of 0 disables automatic rotation for the key.
	AutoRotatePeriod *int `pulumi:"autoRotatePeriod"`
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
	// The key size in bytes for algorithms that allow variable key sizes. Currently only applicable to HMAC, where it must be between 32 and 512 bytes.
	KeySize *int `pulumi:"keySize"`
	// List of key versions in the keyring. This attribute is zero-indexed and will contain a map of values depending on the `type` of the encryption key.
	// * for key types `aes128-gcm96`, `aes256-gcm96` and `chacha20-poly1305`, each key version will be a map of a single value `id` which is just a hash of the key's metadata.
	// * for key types `ed25519`, `ecdsa-p256`, `ecdsa-p384`, `ecdsa-p521`, `rsa-2048`, `rsa-3072` and `rsa-4096`, each key version will be a map of the following:
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
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
	// *Available only for Vault Enterprise*.
	Namespace *string `pulumi:"namespace"`
	// Whether or not the key supports decryption, based on key type.
	SupportsDecryption *bool `pulumi:"supportsDecryption"`
	// Whether or not the key supports derivation, based on key type.
	SupportsDerivation *bool `pulumi:"supportsDerivation"`
	// Whether or not the key supports encryption, based on key type.
	SupportsEncryption *bool `pulumi:"supportsEncryption"`
	// Whether or not the key supports signing, based on key type.
	SupportsSigning *bool `pulumi:"supportsSigning"`
	// Specifies the type of key to create. The currently-supported types are: `aes128-gcm96`, `aes256-gcm96` (default), `chacha20-poly1305`, `ed25519`, `ecdsa-p256`, `ecdsa-p384`, `ecdsa-p521`, `hmac`, `rsa-2048`, `rsa-3072` and `rsa-4096`.
	// * Refer to the Vault documentation on transit key types for more information: [Key Types](https://www.vaultproject.io/docs/secrets/transit#key-types)
	Type *string `pulumi:"type"`
}

type SecretBackendKeyState struct {
	// Enables taking backup of entire keyring in the plaintext format. Once set, this cannot be disabled.
	// * Refer to Vault API documentation on key backups for more information: [Backup Key](https://www.vaultproject.io/api-docs/secret/transit#backup-key)
	AllowPlaintextBackup pulumi.BoolPtrInput
	// Amount of time the key should live before being automatically rotated. A value of 0 disables automatic rotation for the
	// key.
	//
	// Deprecated: Use auto_rotate_period instead
	AutoRotateInterval pulumi.IntPtrInput
	// Amount of seconds the key should live before being automatically rotated.
	// A value of 0 disables automatic rotation for the key.
	AutoRotatePeriod pulumi.IntPtrInput
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
	// The key size in bytes for algorithms that allow variable key sizes. Currently only applicable to HMAC, where it must be between 32 and 512 bytes.
	KeySize pulumi.IntPtrInput
	// List of key versions in the keyring. This attribute is zero-indexed and will contain a map of values depending on the `type` of the encryption key.
	// * for key types `aes128-gcm96`, `aes256-gcm96` and `chacha20-poly1305`, each key version will be a map of a single value `id` which is just a hash of the key's metadata.
	// * for key types `ed25519`, `ecdsa-p256`, `ecdsa-p384`, `ecdsa-p521`, `rsa-2048`, `rsa-3072` and `rsa-4096`, each key version will be a map of the following:
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
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
	// *Available only for Vault Enterprise*.
	Namespace pulumi.StringPtrInput
	// Whether or not the key supports decryption, based on key type.
	SupportsDecryption pulumi.BoolPtrInput
	// Whether or not the key supports derivation, based on key type.
	SupportsDerivation pulumi.BoolPtrInput
	// Whether or not the key supports encryption, based on key type.
	SupportsEncryption pulumi.BoolPtrInput
	// Whether or not the key supports signing, based on key type.
	SupportsSigning pulumi.BoolPtrInput
	// Specifies the type of key to create. The currently-supported types are: `aes128-gcm96`, `aes256-gcm96` (default), `chacha20-poly1305`, `ed25519`, `ecdsa-p256`, `ecdsa-p384`, `ecdsa-p521`, `hmac`, `rsa-2048`, `rsa-3072` and `rsa-4096`.
	// * Refer to the Vault documentation on transit key types for more information: [Key Types](https://www.vaultproject.io/docs/secrets/transit#key-types)
	Type pulumi.StringPtrInput
}

func (SecretBackendKeyState) ElementType() reflect.Type {
	return reflect.TypeOf((*secretBackendKeyState)(nil)).Elem()
}

type secretBackendKeyArgs struct {
	// Enables taking backup of entire keyring in the plaintext format. Once set, this cannot be disabled.
	// * Refer to Vault API documentation on key backups for more information: [Backup Key](https://www.vaultproject.io/api-docs/secret/transit#backup-key)
	AllowPlaintextBackup *bool `pulumi:"allowPlaintextBackup"`
	// Amount of time the key should live before being automatically rotated. A value of 0 disables automatic rotation for the
	// key.
	//
	// Deprecated: Use auto_rotate_period instead
	AutoRotateInterval *int `pulumi:"autoRotateInterval"`
	// Amount of seconds the key should live before being automatically rotated.
	// A value of 0 disables automatic rotation for the key.
	AutoRotatePeriod *int `pulumi:"autoRotatePeriod"`
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
	// The key size in bytes for algorithms that allow variable key sizes. Currently only applicable to HMAC, where it must be between 32 and 512 bytes.
	KeySize *int `pulumi:"keySize"`
	// Minimum key version to use for decryption.
	MinDecryptionVersion *int `pulumi:"minDecryptionVersion"`
	// Minimum key version to use for encryption
	MinEncryptionVersion *int `pulumi:"minEncryptionVersion"`
	// The name to identify this key within the backend. Must be unique within the backend.
	Name *string `pulumi:"name"`
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
	// *Available only for Vault Enterprise*.
	Namespace *string `pulumi:"namespace"`
	// Specifies the type of key to create. The currently-supported types are: `aes128-gcm96`, `aes256-gcm96` (default), `chacha20-poly1305`, `ed25519`, `ecdsa-p256`, `ecdsa-p384`, `ecdsa-p521`, `hmac`, `rsa-2048`, `rsa-3072` and `rsa-4096`.
	// * Refer to the Vault documentation on transit key types for more information: [Key Types](https://www.vaultproject.io/docs/secrets/transit#key-types)
	Type *string `pulumi:"type"`
}

// The set of arguments for constructing a SecretBackendKey resource.
type SecretBackendKeyArgs struct {
	// Enables taking backup of entire keyring in the plaintext format. Once set, this cannot be disabled.
	// * Refer to Vault API documentation on key backups for more information: [Backup Key](https://www.vaultproject.io/api-docs/secret/transit#backup-key)
	AllowPlaintextBackup pulumi.BoolPtrInput
	// Amount of time the key should live before being automatically rotated. A value of 0 disables automatic rotation for the
	// key.
	//
	// Deprecated: Use auto_rotate_period instead
	AutoRotateInterval pulumi.IntPtrInput
	// Amount of seconds the key should live before being automatically rotated.
	// A value of 0 disables automatic rotation for the key.
	AutoRotatePeriod pulumi.IntPtrInput
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
	// The key size in bytes for algorithms that allow variable key sizes. Currently only applicable to HMAC, where it must be between 32 and 512 bytes.
	KeySize pulumi.IntPtrInput
	// Minimum key version to use for decryption.
	MinDecryptionVersion pulumi.IntPtrInput
	// Minimum key version to use for encryption
	MinEncryptionVersion pulumi.IntPtrInput
	// The name to identify this key within the backend. Must be unique within the backend.
	Name pulumi.StringPtrInput
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
	// *Available only for Vault Enterprise*.
	Namespace pulumi.StringPtrInput
	// Specifies the type of key to create. The currently-supported types are: `aes128-gcm96`, `aes256-gcm96` (default), `chacha20-poly1305`, `ed25519`, `ecdsa-p256`, `ecdsa-p384`, `ecdsa-p521`, `hmac`, `rsa-2048`, `rsa-3072` and `rsa-4096`.
	// * Refer to the Vault documentation on transit key types for more information: [Key Types](https://www.vaultproject.io/docs/secrets/transit#key-types)
	Type pulumi.StringPtrInput
}

func (SecretBackendKeyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*secretBackendKeyArgs)(nil)).Elem()
}

type SecretBackendKeyInput interface {
	pulumi.Input

	ToSecretBackendKeyOutput() SecretBackendKeyOutput
	ToSecretBackendKeyOutputWithContext(ctx context.Context) SecretBackendKeyOutput
}

func (*SecretBackendKey) ElementType() reflect.Type {
	return reflect.TypeOf((**SecretBackendKey)(nil)).Elem()
}

func (i *SecretBackendKey) ToSecretBackendKeyOutput() SecretBackendKeyOutput {
	return i.ToSecretBackendKeyOutputWithContext(context.Background())
}

func (i *SecretBackendKey) ToSecretBackendKeyOutputWithContext(ctx context.Context) SecretBackendKeyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecretBackendKeyOutput)
}

// SecretBackendKeyArrayInput is an input type that accepts SecretBackendKeyArray and SecretBackendKeyArrayOutput values.
// You can construct a concrete instance of `SecretBackendKeyArrayInput` via:
//
//	SecretBackendKeyArray{ SecretBackendKeyArgs{...} }
type SecretBackendKeyArrayInput interface {
	pulumi.Input

	ToSecretBackendKeyArrayOutput() SecretBackendKeyArrayOutput
	ToSecretBackendKeyArrayOutputWithContext(context.Context) SecretBackendKeyArrayOutput
}

type SecretBackendKeyArray []SecretBackendKeyInput

func (SecretBackendKeyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SecretBackendKey)(nil)).Elem()
}

func (i SecretBackendKeyArray) ToSecretBackendKeyArrayOutput() SecretBackendKeyArrayOutput {
	return i.ToSecretBackendKeyArrayOutputWithContext(context.Background())
}

func (i SecretBackendKeyArray) ToSecretBackendKeyArrayOutputWithContext(ctx context.Context) SecretBackendKeyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecretBackendKeyArrayOutput)
}

// SecretBackendKeyMapInput is an input type that accepts SecretBackendKeyMap and SecretBackendKeyMapOutput values.
// You can construct a concrete instance of `SecretBackendKeyMapInput` via:
//
//	SecretBackendKeyMap{ "key": SecretBackendKeyArgs{...} }
type SecretBackendKeyMapInput interface {
	pulumi.Input

	ToSecretBackendKeyMapOutput() SecretBackendKeyMapOutput
	ToSecretBackendKeyMapOutputWithContext(context.Context) SecretBackendKeyMapOutput
}

type SecretBackendKeyMap map[string]SecretBackendKeyInput

func (SecretBackendKeyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SecretBackendKey)(nil)).Elem()
}

func (i SecretBackendKeyMap) ToSecretBackendKeyMapOutput() SecretBackendKeyMapOutput {
	return i.ToSecretBackendKeyMapOutputWithContext(context.Background())
}

func (i SecretBackendKeyMap) ToSecretBackendKeyMapOutputWithContext(ctx context.Context) SecretBackendKeyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecretBackendKeyMapOutput)
}

type SecretBackendKeyOutput struct{ *pulumi.OutputState }

func (SecretBackendKeyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SecretBackendKey)(nil)).Elem()
}

func (o SecretBackendKeyOutput) ToSecretBackendKeyOutput() SecretBackendKeyOutput {
	return o
}

func (o SecretBackendKeyOutput) ToSecretBackendKeyOutputWithContext(ctx context.Context) SecretBackendKeyOutput {
	return o
}

// Enables taking backup of entire keyring in the plaintext format. Once set, this cannot be disabled.
// * Refer to Vault API documentation on key backups for more information: [Backup Key](https://www.vaultproject.io/api-docs/secret/transit#backup-key)
func (o SecretBackendKeyOutput) AllowPlaintextBackup() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SecretBackendKey) pulumi.BoolPtrOutput { return v.AllowPlaintextBackup }).(pulumi.BoolPtrOutput)
}

// Amount of time the key should live before being automatically rotated. A value of 0 disables automatic rotation for the
// key.
//
// Deprecated: Use auto_rotate_period instead
func (o SecretBackendKeyOutput) AutoRotateInterval() pulumi.IntOutput {
	return o.ApplyT(func(v *SecretBackendKey) pulumi.IntOutput { return v.AutoRotateInterval }).(pulumi.IntOutput)
}

// Amount of seconds the key should live before being automatically rotated.
// A value of 0 disables automatic rotation for the key.
func (o SecretBackendKeyOutput) AutoRotatePeriod() pulumi.IntOutput {
	return o.ApplyT(func(v *SecretBackendKey) pulumi.IntOutput { return v.AutoRotatePeriod }).(pulumi.IntOutput)
}

// The path the transit secret backend is mounted at, with no leading or trailing `/`s.
func (o SecretBackendKeyOutput) Backend() pulumi.StringOutput {
	return o.ApplyT(func(v *SecretBackendKey) pulumi.StringOutput { return v.Backend }).(pulumi.StringOutput)
}

// Whether or not to support convergent encryption, where the same plaintext creates the same ciphertext. This requires `derived` to be set to `true`.
func (o SecretBackendKeyOutput) ConvergentEncryption() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SecretBackendKey) pulumi.BoolPtrOutput { return v.ConvergentEncryption }).(pulumi.BoolPtrOutput)
}

// Specifies if the key is allowed to be deleted.
func (o SecretBackendKeyOutput) DeletionAllowed() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SecretBackendKey) pulumi.BoolPtrOutput { return v.DeletionAllowed }).(pulumi.BoolPtrOutput)
}

// Specifies if key derivation is to be used. If enabled, all encrypt/decrypt requests to this key must provide a context which is used for key derivation.
func (o SecretBackendKeyOutput) Derived() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SecretBackendKey) pulumi.BoolPtrOutput { return v.Derived }).(pulumi.BoolPtrOutput)
}

// Enables keys to be exportable. This allows for all valid private keys in the keyring to be exported. Once set, this cannot be disabled.
func (o SecretBackendKeyOutput) Exportable() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SecretBackendKey) pulumi.BoolPtrOutput { return v.Exportable }).(pulumi.BoolPtrOutput)
}

// The key size in bytes for algorithms that allow variable key sizes. Currently only applicable to HMAC, where it must be between 32 and 512 bytes.
func (o SecretBackendKeyOutput) KeySize() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SecretBackendKey) pulumi.IntPtrOutput { return v.KeySize }).(pulumi.IntPtrOutput)
}

// List of key versions in the keyring. This attribute is zero-indexed and will contain a map of values depending on the `type` of the encryption key.
// * for key types `aes128-gcm96`, `aes256-gcm96` and `chacha20-poly1305`, each key version will be a map of a single value `id` which is just a hash of the key's metadata.
// * for key types `ed25519`, `ecdsa-p256`, `ecdsa-p384`, `ecdsa-p521`, `rsa-2048`, `rsa-3072` and `rsa-4096`, each key version will be a map of the following:
func (o SecretBackendKeyOutput) Keys() pulumi.MapArrayOutput {
	return o.ApplyT(func(v *SecretBackendKey) pulumi.MapArrayOutput { return v.Keys }).(pulumi.MapArrayOutput)
}

// Latest key version available. This value is 1-indexed, so if `latestVersion` is `1`, then the key's information can be referenced from `keys` by selecting element `0`
func (o SecretBackendKeyOutput) LatestVersion() pulumi.IntOutput {
	return o.ApplyT(func(v *SecretBackendKey) pulumi.IntOutput { return v.LatestVersion }).(pulumi.IntOutput)
}

// Minimum key version available for use. If keys have been archived by increasing `minDecryptionVersion`, this attribute will reflect that change.
func (o SecretBackendKeyOutput) MinAvailableVersion() pulumi.IntOutput {
	return o.ApplyT(func(v *SecretBackendKey) pulumi.IntOutput { return v.MinAvailableVersion }).(pulumi.IntOutput)
}

// Minimum key version to use for decryption.
func (o SecretBackendKeyOutput) MinDecryptionVersion() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SecretBackendKey) pulumi.IntPtrOutput { return v.MinDecryptionVersion }).(pulumi.IntPtrOutput)
}

// Minimum key version to use for encryption
func (o SecretBackendKeyOutput) MinEncryptionVersion() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SecretBackendKey) pulumi.IntPtrOutput { return v.MinEncryptionVersion }).(pulumi.IntPtrOutput)
}

// The name to identify this key within the backend. Must be unique within the backend.
func (o SecretBackendKeyOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SecretBackendKey) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The namespace to provision the resource in.
// The value should not contain leading or trailing forward slashes.
// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
// *Available only for Vault Enterprise*.
func (o SecretBackendKeyOutput) Namespace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecretBackendKey) pulumi.StringPtrOutput { return v.Namespace }).(pulumi.StringPtrOutput)
}

// Whether or not the key supports decryption, based on key type.
func (o SecretBackendKeyOutput) SupportsDecryption() pulumi.BoolOutput {
	return o.ApplyT(func(v *SecretBackendKey) pulumi.BoolOutput { return v.SupportsDecryption }).(pulumi.BoolOutput)
}

// Whether or not the key supports derivation, based on key type.
func (o SecretBackendKeyOutput) SupportsDerivation() pulumi.BoolOutput {
	return o.ApplyT(func(v *SecretBackendKey) pulumi.BoolOutput { return v.SupportsDerivation }).(pulumi.BoolOutput)
}

// Whether or not the key supports encryption, based on key type.
func (o SecretBackendKeyOutput) SupportsEncryption() pulumi.BoolOutput {
	return o.ApplyT(func(v *SecretBackendKey) pulumi.BoolOutput { return v.SupportsEncryption }).(pulumi.BoolOutput)
}

// Whether or not the key supports signing, based on key type.
func (o SecretBackendKeyOutput) SupportsSigning() pulumi.BoolOutput {
	return o.ApplyT(func(v *SecretBackendKey) pulumi.BoolOutput { return v.SupportsSigning }).(pulumi.BoolOutput)
}

// Specifies the type of key to create. The currently-supported types are: `aes128-gcm96`, `aes256-gcm96` (default), `chacha20-poly1305`, `ed25519`, `ecdsa-p256`, `ecdsa-p384`, `ecdsa-p521`, `hmac`, `rsa-2048`, `rsa-3072` and `rsa-4096`.
// * Refer to the Vault documentation on transit key types for more information: [Key Types](https://www.vaultproject.io/docs/secrets/transit#key-types)
func (o SecretBackendKeyOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecretBackendKey) pulumi.StringPtrOutput { return v.Type }).(pulumi.StringPtrOutput)
}

type SecretBackendKeyArrayOutput struct{ *pulumi.OutputState }

func (SecretBackendKeyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SecretBackendKey)(nil)).Elem()
}

func (o SecretBackendKeyArrayOutput) ToSecretBackendKeyArrayOutput() SecretBackendKeyArrayOutput {
	return o
}

func (o SecretBackendKeyArrayOutput) ToSecretBackendKeyArrayOutputWithContext(ctx context.Context) SecretBackendKeyArrayOutput {
	return o
}

func (o SecretBackendKeyArrayOutput) Index(i pulumi.IntInput) SecretBackendKeyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SecretBackendKey {
		return vs[0].([]*SecretBackendKey)[vs[1].(int)]
	}).(SecretBackendKeyOutput)
}

type SecretBackendKeyMapOutput struct{ *pulumi.OutputState }

func (SecretBackendKeyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SecretBackendKey)(nil)).Elem()
}

func (o SecretBackendKeyMapOutput) ToSecretBackendKeyMapOutput() SecretBackendKeyMapOutput {
	return o
}

func (o SecretBackendKeyMapOutput) ToSecretBackendKeyMapOutputWithContext(ctx context.Context) SecretBackendKeyMapOutput {
	return o
}

func (o SecretBackendKeyMapOutput) MapIndex(k pulumi.StringInput) SecretBackendKeyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SecretBackendKey {
		return vs[0].(map[string]*SecretBackendKey)[vs[1].(string)]
	}).(SecretBackendKeyOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SecretBackendKeyInput)(nil)).Elem(), &SecretBackendKey{})
	pulumi.RegisterInputType(reflect.TypeOf((*SecretBackendKeyArrayInput)(nil)).Elem(), SecretBackendKeyArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SecretBackendKeyMapInput)(nil)).Elem(), SecretBackendKeyMap{})
	pulumi.RegisterOutputType(SecretBackendKeyOutput{})
	pulumi.RegisterOutputType(SecretBackendKeyArrayOutput{})
	pulumi.RegisterOutputType(SecretBackendKeyMapOutput{})
}
