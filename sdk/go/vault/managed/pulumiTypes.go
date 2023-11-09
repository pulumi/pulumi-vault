// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package managed

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-vault/sdk/v5/go/vault/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

var _ = internal.GetEnvOrDefault

type KeysAw struct {
	// The AWS access key to use.
	AccessKey string `pulumi:"accessKey"`
	// If no existing key can be found in
	// the referenced backend, instructs Vault to generate a key within the backend.
	AllowGenerateKey *bool `pulumi:"allowGenerateKey"`
	// Controls the ability for Vault to replace through
	// generation or importing a key into the configured backend even
	// if a key is present, if set to `false` those operations are forbidden
	// if a key exists.
	AllowReplaceKey *bool `pulumi:"allowReplaceKey"`
	// Controls the ability for Vault to import a key to the
	// configured backend, if `false`, those operations will be forbidden.
	AllowStoreKey *bool `pulumi:"allowStoreKey"`
	// If `true`, allows usage from any mount point within the
	// namespace.
	AnyMount *bool `pulumi:"anyMount"`
	// The curve to use for an ECDSA key. Used when `keyType`
	// is `ECDSA`. Required if `allowGenerateKey` is `true`.
	Curve *string `pulumi:"curve"`
	// Used to specify a custom AWS endpoint.
	Endpoint *string `pulumi:"endpoint"`
	// The size in bits for an RSA key.
	KeyBits string `pulumi:"keyBits"`
	// The type of key to use.
	KeyType string `pulumi:"keyType"`
	// An identifier for the key.
	KmsKey string `pulumi:"kmsKey"`
	// A unique lowercase name that serves as identifying the key.
	Name string `pulumi:"name"`
	// The AWS region where the keys are stored (or will be stored).
	Region *string `pulumi:"region"`
	// The AWS access key to use.
	SecretKey string  `pulumi:"secretKey"`
	Uuid      *string `pulumi:"uuid"`
}

// KeysAwInput is an input type that accepts KeysAwArgs and KeysAwOutput values.
// You can construct a concrete instance of `KeysAwInput` via:
//
//	KeysAwArgs{...}
type KeysAwInput interface {
	pulumi.Input

	ToKeysAwOutput() KeysAwOutput
	ToKeysAwOutputWithContext(context.Context) KeysAwOutput
}

type KeysAwArgs struct {
	// The AWS access key to use.
	AccessKey pulumi.StringInput `pulumi:"accessKey"`
	// If no existing key can be found in
	// the referenced backend, instructs Vault to generate a key within the backend.
	AllowGenerateKey pulumi.BoolPtrInput `pulumi:"allowGenerateKey"`
	// Controls the ability for Vault to replace through
	// generation or importing a key into the configured backend even
	// if a key is present, if set to `false` those operations are forbidden
	// if a key exists.
	AllowReplaceKey pulumi.BoolPtrInput `pulumi:"allowReplaceKey"`
	// Controls the ability for Vault to import a key to the
	// configured backend, if `false`, those operations will be forbidden.
	AllowStoreKey pulumi.BoolPtrInput `pulumi:"allowStoreKey"`
	// If `true`, allows usage from any mount point within the
	// namespace.
	AnyMount pulumi.BoolPtrInput `pulumi:"anyMount"`
	// The curve to use for an ECDSA key. Used when `keyType`
	// is `ECDSA`. Required if `allowGenerateKey` is `true`.
	Curve pulumi.StringPtrInput `pulumi:"curve"`
	// Used to specify a custom AWS endpoint.
	Endpoint pulumi.StringPtrInput `pulumi:"endpoint"`
	// The size in bits for an RSA key.
	KeyBits pulumi.StringInput `pulumi:"keyBits"`
	// The type of key to use.
	KeyType pulumi.StringInput `pulumi:"keyType"`
	// An identifier for the key.
	KmsKey pulumi.StringInput `pulumi:"kmsKey"`
	// A unique lowercase name that serves as identifying the key.
	Name pulumi.StringInput `pulumi:"name"`
	// The AWS region where the keys are stored (or will be stored).
	Region pulumi.StringPtrInput `pulumi:"region"`
	// The AWS access key to use.
	SecretKey pulumi.StringInput    `pulumi:"secretKey"`
	Uuid      pulumi.StringPtrInput `pulumi:"uuid"`
}

func (KeysAwArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*KeysAw)(nil)).Elem()
}

func (i KeysAwArgs) ToKeysAwOutput() KeysAwOutput {
	return i.ToKeysAwOutputWithContext(context.Background())
}

func (i KeysAwArgs) ToKeysAwOutputWithContext(ctx context.Context) KeysAwOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeysAwOutput)
}

func (i KeysAwArgs) ToOutput(ctx context.Context) pulumix.Output[KeysAw] {
	return pulumix.Output[KeysAw]{
		OutputState: i.ToKeysAwOutputWithContext(ctx).OutputState,
	}
}

// KeysAwArrayInput is an input type that accepts KeysAwArray and KeysAwArrayOutput values.
// You can construct a concrete instance of `KeysAwArrayInput` via:
//
//	KeysAwArray{ KeysAwArgs{...} }
type KeysAwArrayInput interface {
	pulumi.Input

	ToKeysAwArrayOutput() KeysAwArrayOutput
	ToKeysAwArrayOutputWithContext(context.Context) KeysAwArrayOutput
}

type KeysAwArray []KeysAwInput

func (KeysAwArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]KeysAw)(nil)).Elem()
}

func (i KeysAwArray) ToKeysAwArrayOutput() KeysAwArrayOutput {
	return i.ToKeysAwArrayOutputWithContext(context.Background())
}

func (i KeysAwArray) ToKeysAwArrayOutputWithContext(ctx context.Context) KeysAwArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeysAwArrayOutput)
}

func (i KeysAwArray) ToOutput(ctx context.Context) pulumix.Output[[]KeysAw] {
	return pulumix.Output[[]KeysAw]{
		OutputState: i.ToKeysAwArrayOutputWithContext(ctx).OutputState,
	}
}

type KeysAwOutput struct{ *pulumi.OutputState }

func (KeysAwOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*KeysAw)(nil)).Elem()
}

func (o KeysAwOutput) ToKeysAwOutput() KeysAwOutput {
	return o
}

func (o KeysAwOutput) ToKeysAwOutputWithContext(ctx context.Context) KeysAwOutput {
	return o
}

func (o KeysAwOutput) ToOutput(ctx context.Context) pulumix.Output[KeysAw] {
	return pulumix.Output[KeysAw]{
		OutputState: o.OutputState,
	}
}

// The AWS access key to use.
func (o KeysAwOutput) AccessKey() pulumi.StringOutput {
	return o.ApplyT(func(v KeysAw) string { return v.AccessKey }).(pulumi.StringOutput)
}

// If no existing key can be found in
// the referenced backend, instructs Vault to generate a key within the backend.
func (o KeysAwOutput) AllowGenerateKey() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v KeysAw) *bool { return v.AllowGenerateKey }).(pulumi.BoolPtrOutput)
}

// Controls the ability for Vault to replace through
// generation or importing a key into the configured backend even
// if a key is present, if set to `false` those operations are forbidden
// if a key exists.
func (o KeysAwOutput) AllowReplaceKey() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v KeysAw) *bool { return v.AllowReplaceKey }).(pulumi.BoolPtrOutput)
}

// Controls the ability for Vault to import a key to the
// configured backend, if `false`, those operations will be forbidden.
func (o KeysAwOutput) AllowStoreKey() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v KeysAw) *bool { return v.AllowStoreKey }).(pulumi.BoolPtrOutput)
}

// If `true`, allows usage from any mount point within the
// namespace.
func (o KeysAwOutput) AnyMount() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v KeysAw) *bool { return v.AnyMount }).(pulumi.BoolPtrOutput)
}

// The curve to use for an ECDSA key. Used when `keyType`
// is `ECDSA`. Required if `allowGenerateKey` is `true`.
func (o KeysAwOutput) Curve() pulumi.StringPtrOutput {
	return o.ApplyT(func(v KeysAw) *string { return v.Curve }).(pulumi.StringPtrOutput)
}

// Used to specify a custom AWS endpoint.
func (o KeysAwOutput) Endpoint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v KeysAw) *string { return v.Endpoint }).(pulumi.StringPtrOutput)
}

// The size in bits for an RSA key.
func (o KeysAwOutput) KeyBits() pulumi.StringOutput {
	return o.ApplyT(func(v KeysAw) string { return v.KeyBits }).(pulumi.StringOutput)
}

// The type of key to use.
func (o KeysAwOutput) KeyType() pulumi.StringOutput {
	return o.ApplyT(func(v KeysAw) string { return v.KeyType }).(pulumi.StringOutput)
}

// An identifier for the key.
func (o KeysAwOutput) KmsKey() pulumi.StringOutput {
	return o.ApplyT(func(v KeysAw) string { return v.KmsKey }).(pulumi.StringOutput)
}

// A unique lowercase name that serves as identifying the key.
func (o KeysAwOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v KeysAw) string { return v.Name }).(pulumi.StringOutput)
}

// The AWS region where the keys are stored (or will be stored).
func (o KeysAwOutput) Region() pulumi.StringPtrOutput {
	return o.ApplyT(func(v KeysAw) *string { return v.Region }).(pulumi.StringPtrOutput)
}

// The AWS access key to use.
func (o KeysAwOutput) SecretKey() pulumi.StringOutput {
	return o.ApplyT(func(v KeysAw) string { return v.SecretKey }).(pulumi.StringOutput)
}

func (o KeysAwOutput) Uuid() pulumi.StringPtrOutput {
	return o.ApplyT(func(v KeysAw) *string { return v.Uuid }).(pulumi.StringPtrOutput)
}

type KeysAwArrayOutput struct{ *pulumi.OutputState }

func (KeysAwArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]KeysAw)(nil)).Elem()
}

func (o KeysAwArrayOutput) ToKeysAwArrayOutput() KeysAwArrayOutput {
	return o
}

func (o KeysAwArrayOutput) ToKeysAwArrayOutputWithContext(ctx context.Context) KeysAwArrayOutput {
	return o
}

func (o KeysAwArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]KeysAw] {
	return pulumix.Output[[]KeysAw]{
		OutputState: o.OutputState,
	}
}

func (o KeysAwArrayOutput) Index(i pulumi.IntInput) KeysAwOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) KeysAw {
		return vs[0].([]KeysAw)[vs[1].(int)]
	}).(KeysAwOutput)
}

type KeysAzure struct {
	// If no existing key can be found in
	// the referenced backend, instructs Vault to generate a key within the backend.
	AllowGenerateKey *bool `pulumi:"allowGenerateKey"`
	// Controls the ability for Vault to replace through
	// generation or importing a key into the configured backend even
	// if a key is present, if set to `false` those operations are forbidden
	// if a key exists.
	AllowReplaceKey *bool `pulumi:"allowReplaceKey"`
	// Controls the ability for Vault to import a key to the
	// configured backend, if `false`, those operations will be forbidden.
	AllowStoreKey *bool `pulumi:"allowStoreKey"`
	// If `true`, allows usage from any mount point within the
	// namespace.
	AnyMount *bool `pulumi:"anyMount"`
	// The client id for credentials to query the Azure APIs.
	ClientId string `pulumi:"clientId"`
	// The client secret for credentials to query the Azure APIs.
	ClientSecret string `pulumi:"clientSecret"`
	// The Azure Cloud environment API endpoints to use.
	Environment *string `pulumi:"environment"`
	// The size in bits for an RSA key.
	KeyBits *string `pulumi:"keyBits"`
	// The Key Vault key to use for encryption and decryption.
	KeyName string `pulumi:"keyName"`
	// The type of key to use.
	KeyType string `pulumi:"keyType"`
	// A unique lowercase name that serves as identifying the key.
	Name string `pulumi:"name"`
	// The Azure Key Vault resource's DNS Suffix to connect to.
	Resource *string `pulumi:"resource"`
	// The tenant id for the Azure Active Directory organization.
	TenantId string  `pulumi:"tenantId"`
	Uuid     *string `pulumi:"uuid"`
	// The Key Vault vault to use for encryption and decryption.
	VaultName string `pulumi:"vaultName"`
}

// KeysAzureInput is an input type that accepts KeysAzureArgs and KeysAzureOutput values.
// You can construct a concrete instance of `KeysAzureInput` via:
//
//	KeysAzureArgs{...}
type KeysAzureInput interface {
	pulumi.Input

	ToKeysAzureOutput() KeysAzureOutput
	ToKeysAzureOutputWithContext(context.Context) KeysAzureOutput
}

type KeysAzureArgs struct {
	// If no existing key can be found in
	// the referenced backend, instructs Vault to generate a key within the backend.
	AllowGenerateKey pulumi.BoolPtrInput `pulumi:"allowGenerateKey"`
	// Controls the ability for Vault to replace through
	// generation or importing a key into the configured backend even
	// if a key is present, if set to `false` those operations are forbidden
	// if a key exists.
	AllowReplaceKey pulumi.BoolPtrInput `pulumi:"allowReplaceKey"`
	// Controls the ability for Vault to import a key to the
	// configured backend, if `false`, those operations will be forbidden.
	AllowStoreKey pulumi.BoolPtrInput `pulumi:"allowStoreKey"`
	// If `true`, allows usage from any mount point within the
	// namespace.
	AnyMount pulumi.BoolPtrInput `pulumi:"anyMount"`
	// The client id for credentials to query the Azure APIs.
	ClientId pulumi.StringInput `pulumi:"clientId"`
	// The client secret for credentials to query the Azure APIs.
	ClientSecret pulumi.StringInput `pulumi:"clientSecret"`
	// The Azure Cloud environment API endpoints to use.
	Environment pulumi.StringPtrInput `pulumi:"environment"`
	// The size in bits for an RSA key.
	KeyBits pulumi.StringPtrInput `pulumi:"keyBits"`
	// The Key Vault key to use for encryption and decryption.
	KeyName pulumi.StringInput `pulumi:"keyName"`
	// The type of key to use.
	KeyType pulumi.StringInput `pulumi:"keyType"`
	// A unique lowercase name that serves as identifying the key.
	Name pulumi.StringInput `pulumi:"name"`
	// The Azure Key Vault resource's DNS Suffix to connect to.
	Resource pulumi.StringPtrInput `pulumi:"resource"`
	// The tenant id for the Azure Active Directory organization.
	TenantId pulumi.StringInput    `pulumi:"tenantId"`
	Uuid     pulumi.StringPtrInput `pulumi:"uuid"`
	// The Key Vault vault to use for encryption and decryption.
	VaultName pulumi.StringInput `pulumi:"vaultName"`
}

func (KeysAzureArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*KeysAzure)(nil)).Elem()
}

func (i KeysAzureArgs) ToKeysAzureOutput() KeysAzureOutput {
	return i.ToKeysAzureOutputWithContext(context.Background())
}

func (i KeysAzureArgs) ToKeysAzureOutputWithContext(ctx context.Context) KeysAzureOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeysAzureOutput)
}

func (i KeysAzureArgs) ToOutput(ctx context.Context) pulumix.Output[KeysAzure] {
	return pulumix.Output[KeysAzure]{
		OutputState: i.ToKeysAzureOutputWithContext(ctx).OutputState,
	}
}

// KeysAzureArrayInput is an input type that accepts KeysAzureArray and KeysAzureArrayOutput values.
// You can construct a concrete instance of `KeysAzureArrayInput` via:
//
//	KeysAzureArray{ KeysAzureArgs{...} }
type KeysAzureArrayInput interface {
	pulumi.Input

	ToKeysAzureArrayOutput() KeysAzureArrayOutput
	ToKeysAzureArrayOutputWithContext(context.Context) KeysAzureArrayOutput
}

type KeysAzureArray []KeysAzureInput

func (KeysAzureArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]KeysAzure)(nil)).Elem()
}

func (i KeysAzureArray) ToKeysAzureArrayOutput() KeysAzureArrayOutput {
	return i.ToKeysAzureArrayOutputWithContext(context.Background())
}

func (i KeysAzureArray) ToKeysAzureArrayOutputWithContext(ctx context.Context) KeysAzureArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeysAzureArrayOutput)
}

func (i KeysAzureArray) ToOutput(ctx context.Context) pulumix.Output[[]KeysAzure] {
	return pulumix.Output[[]KeysAzure]{
		OutputState: i.ToKeysAzureArrayOutputWithContext(ctx).OutputState,
	}
}

type KeysAzureOutput struct{ *pulumi.OutputState }

func (KeysAzureOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*KeysAzure)(nil)).Elem()
}

func (o KeysAzureOutput) ToKeysAzureOutput() KeysAzureOutput {
	return o
}

func (o KeysAzureOutput) ToKeysAzureOutputWithContext(ctx context.Context) KeysAzureOutput {
	return o
}

func (o KeysAzureOutput) ToOutput(ctx context.Context) pulumix.Output[KeysAzure] {
	return pulumix.Output[KeysAzure]{
		OutputState: o.OutputState,
	}
}

// If no existing key can be found in
// the referenced backend, instructs Vault to generate a key within the backend.
func (o KeysAzureOutput) AllowGenerateKey() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v KeysAzure) *bool { return v.AllowGenerateKey }).(pulumi.BoolPtrOutput)
}

// Controls the ability for Vault to replace through
// generation or importing a key into the configured backend even
// if a key is present, if set to `false` those operations are forbidden
// if a key exists.
func (o KeysAzureOutput) AllowReplaceKey() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v KeysAzure) *bool { return v.AllowReplaceKey }).(pulumi.BoolPtrOutput)
}

// Controls the ability for Vault to import a key to the
// configured backend, if `false`, those operations will be forbidden.
func (o KeysAzureOutput) AllowStoreKey() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v KeysAzure) *bool { return v.AllowStoreKey }).(pulumi.BoolPtrOutput)
}

// If `true`, allows usage from any mount point within the
// namespace.
func (o KeysAzureOutput) AnyMount() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v KeysAzure) *bool { return v.AnyMount }).(pulumi.BoolPtrOutput)
}

// The client id for credentials to query the Azure APIs.
func (o KeysAzureOutput) ClientId() pulumi.StringOutput {
	return o.ApplyT(func(v KeysAzure) string { return v.ClientId }).(pulumi.StringOutput)
}

// The client secret for credentials to query the Azure APIs.
func (o KeysAzureOutput) ClientSecret() pulumi.StringOutput {
	return o.ApplyT(func(v KeysAzure) string { return v.ClientSecret }).(pulumi.StringOutput)
}

// The Azure Cloud environment API endpoints to use.
func (o KeysAzureOutput) Environment() pulumi.StringPtrOutput {
	return o.ApplyT(func(v KeysAzure) *string { return v.Environment }).(pulumi.StringPtrOutput)
}

// The size in bits for an RSA key.
func (o KeysAzureOutput) KeyBits() pulumi.StringPtrOutput {
	return o.ApplyT(func(v KeysAzure) *string { return v.KeyBits }).(pulumi.StringPtrOutput)
}

// The Key Vault key to use for encryption and decryption.
func (o KeysAzureOutput) KeyName() pulumi.StringOutput {
	return o.ApplyT(func(v KeysAzure) string { return v.KeyName }).(pulumi.StringOutput)
}

// The type of key to use.
func (o KeysAzureOutput) KeyType() pulumi.StringOutput {
	return o.ApplyT(func(v KeysAzure) string { return v.KeyType }).(pulumi.StringOutput)
}

// A unique lowercase name that serves as identifying the key.
func (o KeysAzureOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v KeysAzure) string { return v.Name }).(pulumi.StringOutput)
}

// The Azure Key Vault resource's DNS Suffix to connect to.
func (o KeysAzureOutput) Resource() pulumi.StringPtrOutput {
	return o.ApplyT(func(v KeysAzure) *string { return v.Resource }).(pulumi.StringPtrOutput)
}

// The tenant id for the Azure Active Directory organization.
func (o KeysAzureOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v KeysAzure) string { return v.TenantId }).(pulumi.StringOutput)
}

func (o KeysAzureOutput) Uuid() pulumi.StringPtrOutput {
	return o.ApplyT(func(v KeysAzure) *string { return v.Uuid }).(pulumi.StringPtrOutput)
}

// The Key Vault vault to use for encryption and decryption.
func (o KeysAzureOutput) VaultName() pulumi.StringOutput {
	return o.ApplyT(func(v KeysAzure) string { return v.VaultName }).(pulumi.StringOutput)
}

type KeysAzureArrayOutput struct{ *pulumi.OutputState }

func (KeysAzureArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]KeysAzure)(nil)).Elem()
}

func (o KeysAzureArrayOutput) ToKeysAzureArrayOutput() KeysAzureArrayOutput {
	return o
}

func (o KeysAzureArrayOutput) ToKeysAzureArrayOutputWithContext(ctx context.Context) KeysAzureArrayOutput {
	return o
}

func (o KeysAzureArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]KeysAzure] {
	return pulumix.Output[[]KeysAzure]{
		OutputState: o.OutputState,
	}
}

func (o KeysAzureArrayOutput) Index(i pulumi.IntInput) KeysAzureOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) KeysAzure {
		return vs[0].([]KeysAzure)[vs[1].(int)]
	}).(KeysAzureOutput)
}

type KeysPkc struct {
	// If no existing key can be found in
	// the referenced backend, instructs Vault to generate a key within the backend.
	AllowGenerateKey *bool `pulumi:"allowGenerateKey"`
	// Controls the ability for Vault to replace through
	// generation or importing a key into the configured backend even
	// if a key is present, if set to `false` those operations are forbidden
	// if a key exists.
	AllowReplaceKey *bool `pulumi:"allowReplaceKey"`
	// Controls the ability for Vault to import a key to the
	// configured backend, if `false`, those operations will be forbidden.
	AllowStoreKey *bool `pulumi:"allowStoreKey"`
	// If `true`, allows usage from any mount point within the
	// namespace.
	AnyMount *bool `pulumi:"anyMount"`
	// The curve to use for an ECDSA key. Used when `keyType`
	// is `ECDSA`. Required if `allowGenerateKey` is `true`.
	Curve *string `pulumi:"curve"`
	// Force all operations to open up a read-write session to
	// the HSM.
	ForceRwSession *string `pulumi:"forceRwSession"`
	// The size in bits for an RSA key.
	KeyBits *string `pulumi:"keyBits"`
	// The id of a PKCS#11 key to use.
	KeyId string `pulumi:"keyId"`
	// The label of the key to use.
	KeyLabel string `pulumi:"keyLabel"`
	// The name of the kmsLibrary stanza to use from Vault's config
	// to lookup the local library path.
	Library string `pulumi:"library"`
	// The encryption/decryption mechanism to use, specified as a
	// hexadecimal (prefixed by 0x) string.
	Mechanism string `pulumi:"mechanism"`
	// A unique lowercase name that serves as identifying the key.
	Name string `pulumi:"name"`
	// The PIN for login.
	Pin string `pulumi:"pin"`
	// The slot number to use, specified as a string in a decimal format
	// (e.g. `2305843009213693953`).
	Slot *string `pulumi:"slot"`
	// The slot token label to use.
	TokenLabel *string `pulumi:"tokenLabel"`
	Uuid       *string `pulumi:"uuid"`
}

// KeysPkcInput is an input type that accepts KeysPkcArgs and KeysPkcOutput values.
// You can construct a concrete instance of `KeysPkcInput` via:
//
//	KeysPkcArgs{...}
type KeysPkcInput interface {
	pulumi.Input

	ToKeysPkcOutput() KeysPkcOutput
	ToKeysPkcOutputWithContext(context.Context) KeysPkcOutput
}

type KeysPkcArgs struct {
	// If no existing key can be found in
	// the referenced backend, instructs Vault to generate a key within the backend.
	AllowGenerateKey pulumi.BoolPtrInput `pulumi:"allowGenerateKey"`
	// Controls the ability for Vault to replace through
	// generation or importing a key into the configured backend even
	// if a key is present, if set to `false` those operations are forbidden
	// if a key exists.
	AllowReplaceKey pulumi.BoolPtrInput `pulumi:"allowReplaceKey"`
	// Controls the ability for Vault to import a key to the
	// configured backend, if `false`, those operations will be forbidden.
	AllowStoreKey pulumi.BoolPtrInput `pulumi:"allowStoreKey"`
	// If `true`, allows usage from any mount point within the
	// namespace.
	AnyMount pulumi.BoolPtrInput `pulumi:"anyMount"`
	// The curve to use for an ECDSA key. Used when `keyType`
	// is `ECDSA`. Required if `allowGenerateKey` is `true`.
	Curve pulumi.StringPtrInput `pulumi:"curve"`
	// Force all operations to open up a read-write session to
	// the HSM.
	ForceRwSession pulumi.StringPtrInput `pulumi:"forceRwSession"`
	// The size in bits for an RSA key.
	KeyBits pulumi.StringPtrInput `pulumi:"keyBits"`
	// The id of a PKCS#11 key to use.
	KeyId pulumi.StringInput `pulumi:"keyId"`
	// The label of the key to use.
	KeyLabel pulumi.StringInput `pulumi:"keyLabel"`
	// The name of the kmsLibrary stanza to use from Vault's config
	// to lookup the local library path.
	Library pulumi.StringInput `pulumi:"library"`
	// The encryption/decryption mechanism to use, specified as a
	// hexadecimal (prefixed by 0x) string.
	Mechanism pulumi.StringInput `pulumi:"mechanism"`
	// A unique lowercase name that serves as identifying the key.
	Name pulumi.StringInput `pulumi:"name"`
	// The PIN for login.
	Pin pulumi.StringInput `pulumi:"pin"`
	// The slot number to use, specified as a string in a decimal format
	// (e.g. `2305843009213693953`).
	Slot pulumi.StringPtrInput `pulumi:"slot"`
	// The slot token label to use.
	TokenLabel pulumi.StringPtrInput `pulumi:"tokenLabel"`
	Uuid       pulumi.StringPtrInput `pulumi:"uuid"`
}

func (KeysPkcArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*KeysPkc)(nil)).Elem()
}

func (i KeysPkcArgs) ToKeysPkcOutput() KeysPkcOutput {
	return i.ToKeysPkcOutputWithContext(context.Background())
}

func (i KeysPkcArgs) ToKeysPkcOutputWithContext(ctx context.Context) KeysPkcOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeysPkcOutput)
}

func (i KeysPkcArgs) ToOutput(ctx context.Context) pulumix.Output[KeysPkc] {
	return pulumix.Output[KeysPkc]{
		OutputState: i.ToKeysPkcOutputWithContext(ctx).OutputState,
	}
}

// KeysPkcArrayInput is an input type that accepts KeysPkcArray and KeysPkcArrayOutput values.
// You can construct a concrete instance of `KeysPkcArrayInput` via:
//
//	KeysPkcArray{ KeysPkcArgs{...} }
type KeysPkcArrayInput interface {
	pulumi.Input

	ToKeysPkcArrayOutput() KeysPkcArrayOutput
	ToKeysPkcArrayOutputWithContext(context.Context) KeysPkcArrayOutput
}

type KeysPkcArray []KeysPkcInput

func (KeysPkcArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]KeysPkc)(nil)).Elem()
}

func (i KeysPkcArray) ToKeysPkcArrayOutput() KeysPkcArrayOutput {
	return i.ToKeysPkcArrayOutputWithContext(context.Background())
}

func (i KeysPkcArray) ToKeysPkcArrayOutputWithContext(ctx context.Context) KeysPkcArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeysPkcArrayOutput)
}

func (i KeysPkcArray) ToOutput(ctx context.Context) pulumix.Output[[]KeysPkc] {
	return pulumix.Output[[]KeysPkc]{
		OutputState: i.ToKeysPkcArrayOutputWithContext(ctx).OutputState,
	}
}

type KeysPkcOutput struct{ *pulumi.OutputState }

func (KeysPkcOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*KeysPkc)(nil)).Elem()
}

func (o KeysPkcOutput) ToKeysPkcOutput() KeysPkcOutput {
	return o
}

func (o KeysPkcOutput) ToKeysPkcOutputWithContext(ctx context.Context) KeysPkcOutput {
	return o
}

func (o KeysPkcOutput) ToOutput(ctx context.Context) pulumix.Output[KeysPkc] {
	return pulumix.Output[KeysPkc]{
		OutputState: o.OutputState,
	}
}

// If no existing key can be found in
// the referenced backend, instructs Vault to generate a key within the backend.
func (o KeysPkcOutput) AllowGenerateKey() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v KeysPkc) *bool { return v.AllowGenerateKey }).(pulumi.BoolPtrOutput)
}

// Controls the ability for Vault to replace through
// generation or importing a key into the configured backend even
// if a key is present, if set to `false` those operations are forbidden
// if a key exists.
func (o KeysPkcOutput) AllowReplaceKey() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v KeysPkc) *bool { return v.AllowReplaceKey }).(pulumi.BoolPtrOutput)
}

// Controls the ability for Vault to import a key to the
// configured backend, if `false`, those operations will be forbidden.
func (o KeysPkcOutput) AllowStoreKey() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v KeysPkc) *bool { return v.AllowStoreKey }).(pulumi.BoolPtrOutput)
}

// If `true`, allows usage from any mount point within the
// namespace.
func (o KeysPkcOutput) AnyMount() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v KeysPkc) *bool { return v.AnyMount }).(pulumi.BoolPtrOutput)
}

// The curve to use for an ECDSA key. Used when `keyType`
// is `ECDSA`. Required if `allowGenerateKey` is `true`.
func (o KeysPkcOutput) Curve() pulumi.StringPtrOutput {
	return o.ApplyT(func(v KeysPkc) *string { return v.Curve }).(pulumi.StringPtrOutput)
}

// Force all operations to open up a read-write session to
// the HSM.
func (o KeysPkcOutput) ForceRwSession() pulumi.StringPtrOutput {
	return o.ApplyT(func(v KeysPkc) *string { return v.ForceRwSession }).(pulumi.StringPtrOutput)
}

// The size in bits for an RSA key.
func (o KeysPkcOutput) KeyBits() pulumi.StringPtrOutput {
	return o.ApplyT(func(v KeysPkc) *string { return v.KeyBits }).(pulumi.StringPtrOutput)
}

// The id of a PKCS#11 key to use.
func (o KeysPkcOutput) KeyId() pulumi.StringOutput {
	return o.ApplyT(func(v KeysPkc) string { return v.KeyId }).(pulumi.StringOutput)
}

// The label of the key to use.
func (o KeysPkcOutput) KeyLabel() pulumi.StringOutput {
	return o.ApplyT(func(v KeysPkc) string { return v.KeyLabel }).(pulumi.StringOutput)
}

// The name of the kmsLibrary stanza to use from Vault's config
// to lookup the local library path.
func (o KeysPkcOutput) Library() pulumi.StringOutput {
	return o.ApplyT(func(v KeysPkc) string { return v.Library }).(pulumi.StringOutput)
}

// The encryption/decryption mechanism to use, specified as a
// hexadecimal (prefixed by 0x) string.
func (o KeysPkcOutput) Mechanism() pulumi.StringOutput {
	return o.ApplyT(func(v KeysPkc) string { return v.Mechanism }).(pulumi.StringOutput)
}

// A unique lowercase name that serves as identifying the key.
func (o KeysPkcOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v KeysPkc) string { return v.Name }).(pulumi.StringOutput)
}

// The PIN for login.
func (o KeysPkcOutput) Pin() pulumi.StringOutput {
	return o.ApplyT(func(v KeysPkc) string { return v.Pin }).(pulumi.StringOutput)
}

// The slot number to use, specified as a string in a decimal format
// (e.g. `2305843009213693953`).
func (o KeysPkcOutput) Slot() pulumi.StringPtrOutput {
	return o.ApplyT(func(v KeysPkc) *string { return v.Slot }).(pulumi.StringPtrOutput)
}

// The slot token label to use.
func (o KeysPkcOutput) TokenLabel() pulumi.StringPtrOutput {
	return o.ApplyT(func(v KeysPkc) *string { return v.TokenLabel }).(pulumi.StringPtrOutput)
}

func (o KeysPkcOutput) Uuid() pulumi.StringPtrOutput {
	return o.ApplyT(func(v KeysPkc) *string { return v.Uuid }).(pulumi.StringPtrOutput)
}

type KeysPkcArrayOutput struct{ *pulumi.OutputState }

func (KeysPkcArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]KeysPkc)(nil)).Elem()
}

func (o KeysPkcArrayOutput) ToKeysPkcArrayOutput() KeysPkcArrayOutput {
	return o
}

func (o KeysPkcArrayOutput) ToKeysPkcArrayOutputWithContext(ctx context.Context) KeysPkcArrayOutput {
	return o
}

func (o KeysPkcArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]KeysPkc] {
	return pulumix.Output[[]KeysPkc]{
		OutputState: o.OutputState,
	}
}

func (o KeysPkcArrayOutput) Index(i pulumi.IntInput) KeysPkcOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) KeysPkc {
		return vs[0].([]KeysPkc)[vs[1].(int)]
	}).(KeysPkcOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*KeysAwInput)(nil)).Elem(), KeysAwArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*KeysAwArrayInput)(nil)).Elem(), KeysAwArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*KeysAzureInput)(nil)).Elem(), KeysAzureArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*KeysAzureArrayInput)(nil)).Elem(), KeysAzureArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*KeysPkcInput)(nil)).Elem(), KeysPkcArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*KeysPkcArrayInput)(nil)).Elem(), KeysPkcArray{})
	pulumi.RegisterOutputType(KeysAwOutput{})
	pulumi.RegisterOutputType(KeysAwArrayOutput{})
	pulumi.RegisterOutputType(KeysAzureOutput{})
	pulumi.RegisterOutputType(KeysAzureArrayOutput{})
	pulumi.RegisterOutputType(KeysPkcOutput{})
	pulumi.RegisterOutputType(KeysPkcArrayOutput{})
}
