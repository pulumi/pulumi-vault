// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package secrets

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-vault/sdk/v6/go/vault/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## Example Usage
//
// <!--Start PulumiCodeChooser -->
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-vault/sdk/v6/go/vault/secrets"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := secrets.NewSyncAzureDestination(ctx, "az", &secrets.SyncAzureDestinationArgs{
//				KeyVaultUri:        pulumi.Any(_var.Key_vault_uri),
//				ClientId:           pulumi.Any(_var.Client_id),
//				ClientSecret:       pulumi.Any(_var.Client_secret),
//				TenantId:           pulumi.Any(_var.Tenant_id),
//				SecretNameTemplate: pulumi.String("vault_{{ .MountAccessor | lowercase }}_{{ .SecretPath | lowercase }}"),
//				CustomTags: pulumi.Map{
//					"foo": pulumi.Any("bar"),
//				},
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
// ## Import
//
// Azure Secrets sync destinations can be imported using the `name`, e.g.
//
// ```sh
// $ pulumi import vault:secrets/syncAzureDestination:SyncAzureDestination az az-dest
// ```
type SyncAzureDestination struct {
	pulumi.CustomResourceState

	// Client ID of an Azure app registration.
	// Can be omitted and directly provided to Vault using the `AZURE_CLIENT_ID` environment
	// variable.
	ClientId pulumi.StringPtrOutput `pulumi:"clientId"`
	// Client Secret of an Azure app registration.
	// Can be omitted and directly provided to Vault using the `AZURE_CLIENT_SECRET` environment
	// variable.
	ClientSecret pulumi.StringPtrOutput `pulumi:"clientSecret"`
	// Specifies a cloud for the client. The default is Azure Public Cloud.
	Cloud pulumi.StringPtrOutput `pulumi:"cloud"`
	// Custom tags to set on the secret managed at the destination.
	CustomTags pulumi.MapOutput `pulumi:"customTags"`
	// Determines what level of information is synced as a distinct resource
	// at the destination. Supports `secret-path` and `secret-key`.
	Granularity pulumi.StringPtrOutput `pulumi:"granularity"`
	// URI of an existing Azure Key Vault instance.
	// Can be omitted and directly provided to Vault using the `KEY_VAULT_URI` environment
	// variable.
	KeyVaultUri pulumi.StringPtrOutput `pulumi:"keyVaultUri"`
	// Unique name of the Azure destination.
	Name pulumi.StringOutput `pulumi:"name"`
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
	Namespace pulumi.StringPtrOutput `pulumi:"namespace"`
	// Template describing how to generate external secret names.
	// Supports a subset of the Go Template syntax.
	SecretNameTemplate pulumi.StringOutput `pulumi:"secretNameTemplate"`
	// ID of the target Azure tenant.
	// Can be omitted and directly provided to Vault using the `AZURE_TENANT_ID` environment
	// variable.
	TenantId pulumi.StringPtrOutput `pulumi:"tenantId"`
	// The type of the secrets destination (`azure-kv`).
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewSyncAzureDestination registers a new resource with the given unique name, arguments, and options.
func NewSyncAzureDestination(ctx *pulumi.Context,
	name string, args *SyncAzureDestinationArgs, opts ...pulumi.ResourceOption) (*SyncAzureDestination, error) {
	if args == nil {
		args = &SyncAzureDestinationArgs{}
	}

	if args.ClientSecret != nil {
		args.ClientSecret = pulumi.ToSecret(args.ClientSecret).(pulumi.StringPtrInput)
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"clientSecret",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource SyncAzureDestination
	err := ctx.RegisterResource("vault:secrets/syncAzureDestination:SyncAzureDestination", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSyncAzureDestination gets an existing SyncAzureDestination resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSyncAzureDestination(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SyncAzureDestinationState, opts ...pulumi.ResourceOption) (*SyncAzureDestination, error) {
	var resource SyncAzureDestination
	err := ctx.ReadResource("vault:secrets/syncAzureDestination:SyncAzureDestination", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SyncAzureDestination resources.
type syncAzureDestinationState struct {
	// Client ID of an Azure app registration.
	// Can be omitted and directly provided to Vault using the `AZURE_CLIENT_ID` environment
	// variable.
	ClientId *string `pulumi:"clientId"`
	// Client Secret of an Azure app registration.
	// Can be omitted and directly provided to Vault using the `AZURE_CLIENT_SECRET` environment
	// variable.
	ClientSecret *string `pulumi:"clientSecret"`
	// Specifies a cloud for the client. The default is Azure Public Cloud.
	Cloud *string `pulumi:"cloud"`
	// Custom tags to set on the secret managed at the destination.
	CustomTags map[string]interface{} `pulumi:"customTags"`
	// Determines what level of information is synced as a distinct resource
	// at the destination. Supports `secret-path` and `secret-key`.
	Granularity *string `pulumi:"granularity"`
	// URI of an existing Azure Key Vault instance.
	// Can be omitted and directly provided to Vault using the `KEY_VAULT_URI` environment
	// variable.
	KeyVaultUri *string `pulumi:"keyVaultUri"`
	// Unique name of the Azure destination.
	Name *string `pulumi:"name"`
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
	Namespace *string `pulumi:"namespace"`
	// Template describing how to generate external secret names.
	// Supports a subset of the Go Template syntax.
	SecretNameTemplate *string `pulumi:"secretNameTemplate"`
	// ID of the target Azure tenant.
	// Can be omitted and directly provided to Vault using the `AZURE_TENANT_ID` environment
	// variable.
	TenantId *string `pulumi:"tenantId"`
	// The type of the secrets destination (`azure-kv`).
	Type *string `pulumi:"type"`
}

type SyncAzureDestinationState struct {
	// Client ID of an Azure app registration.
	// Can be omitted and directly provided to Vault using the `AZURE_CLIENT_ID` environment
	// variable.
	ClientId pulumi.StringPtrInput
	// Client Secret of an Azure app registration.
	// Can be omitted and directly provided to Vault using the `AZURE_CLIENT_SECRET` environment
	// variable.
	ClientSecret pulumi.StringPtrInput
	// Specifies a cloud for the client. The default is Azure Public Cloud.
	Cloud pulumi.StringPtrInput
	// Custom tags to set on the secret managed at the destination.
	CustomTags pulumi.MapInput
	// Determines what level of information is synced as a distinct resource
	// at the destination. Supports `secret-path` and `secret-key`.
	Granularity pulumi.StringPtrInput
	// URI of an existing Azure Key Vault instance.
	// Can be omitted and directly provided to Vault using the `KEY_VAULT_URI` environment
	// variable.
	KeyVaultUri pulumi.StringPtrInput
	// Unique name of the Azure destination.
	Name pulumi.StringPtrInput
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
	Namespace pulumi.StringPtrInput
	// Template describing how to generate external secret names.
	// Supports a subset of the Go Template syntax.
	SecretNameTemplate pulumi.StringPtrInput
	// ID of the target Azure tenant.
	// Can be omitted and directly provided to Vault using the `AZURE_TENANT_ID` environment
	// variable.
	TenantId pulumi.StringPtrInput
	// The type of the secrets destination (`azure-kv`).
	Type pulumi.StringPtrInput
}

func (SyncAzureDestinationState) ElementType() reflect.Type {
	return reflect.TypeOf((*syncAzureDestinationState)(nil)).Elem()
}

type syncAzureDestinationArgs struct {
	// Client ID of an Azure app registration.
	// Can be omitted and directly provided to Vault using the `AZURE_CLIENT_ID` environment
	// variable.
	ClientId *string `pulumi:"clientId"`
	// Client Secret of an Azure app registration.
	// Can be omitted and directly provided to Vault using the `AZURE_CLIENT_SECRET` environment
	// variable.
	ClientSecret *string `pulumi:"clientSecret"`
	// Specifies a cloud for the client. The default is Azure Public Cloud.
	Cloud *string `pulumi:"cloud"`
	// Custom tags to set on the secret managed at the destination.
	CustomTags map[string]interface{} `pulumi:"customTags"`
	// Determines what level of information is synced as a distinct resource
	// at the destination. Supports `secret-path` and `secret-key`.
	Granularity *string `pulumi:"granularity"`
	// URI of an existing Azure Key Vault instance.
	// Can be omitted and directly provided to Vault using the `KEY_VAULT_URI` environment
	// variable.
	KeyVaultUri *string `pulumi:"keyVaultUri"`
	// Unique name of the Azure destination.
	Name *string `pulumi:"name"`
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
	Namespace *string `pulumi:"namespace"`
	// Template describing how to generate external secret names.
	// Supports a subset of the Go Template syntax.
	SecretNameTemplate *string `pulumi:"secretNameTemplate"`
	// ID of the target Azure tenant.
	// Can be omitted and directly provided to Vault using the `AZURE_TENANT_ID` environment
	// variable.
	TenantId *string `pulumi:"tenantId"`
}

// The set of arguments for constructing a SyncAzureDestination resource.
type SyncAzureDestinationArgs struct {
	// Client ID of an Azure app registration.
	// Can be omitted and directly provided to Vault using the `AZURE_CLIENT_ID` environment
	// variable.
	ClientId pulumi.StringPtrInput
	// Client Secret of an Azure app registration.
	// Can be omitted and directly provided to Vault using the `AZURE_CLIENT_SECRET` environment
	// variable.
	ClientSecret pulumi.StringPtrInput
	// Specifies a cloud for the client. The default is Azure Public Cloud.
	Cloud pulumi.StringPtrInput
	// Custom tags to set on the secret managed at the destination.
	CustomTags pulumi.MapInput
	// Determines what level of information is synced as a distinct resource
	// at the destination. Supports `secret-path` and `secret-key`.
	Granularity pulumi.StringPtrInput
	// URI of an existing Azure Key Vault instance.
	// Can be omitted and directly provided to Vault using the `KEY_VAULT_URI` environment
	// variable.
	KeyVaultUri pulumi.StringPtrInput
	// Unique name of the Azure destination.
	Name pulumi.StringPtrInput
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
	Namespace pulumi.StringPtrInput
	// Template describing how to generate external secret names.
	// Supports a subset of the Go Template syntax.
	SecretNameTemplate pulumi.StringPtrInput
	// ID of the target Azure tenant.
	// Can be omitted and directly provided to Vault using the `AZURE_TENANT_ID` environment
	// variable.
	TenantId pulumi.StringPtrInput
}

func (SyncAzureDestinationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*syncAzureDestinationArgs)(nil)).Elem()
}

type SyncAzureDestinationInput interface {
	pulumi.Input

	ToSyncAzureDestinationOutput() SyncAzureDestinationOutput
	ToSyncAzureDestinationOutputWithContext(ctx context.Context) SyncAzureDestinationOutput
}

func (*SyncAzureDestination) ElementType() reflect.Type {
	return reflect.TypeOf((**SyncAzureDestination)(nil)).Elem()
}

func (i *SyncAzureDestination) ToSyncAzureDestinationOutput() SyncAzureDestinationOutput {
	return i.ToSyncAzureDestinationOutputWithContext(context.Background())
}

func (i *SyncAzureDestination) ToSyncAzureDestinationOutputWithContext(ctx context.Context) SyncAzureDestinationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SyncAzureDestinationOutput)
}

// SyncAzureDestinationArrayInput is an input type that accepts SyncAzureDestinationArray and SyncAzureDestinationArrayOutput values.
// You can construct a concrete instance of `SyncAzureDestinationArrayInput` via:
//
//	SyncAzureDestinationArray{ SyncAzureDestinationArgs{...} }
type SyncAzureDestinationArrayInput interface {
	pulumi.Input

	ToSyncAzureDestinationArrayOutput() SyncAzureDestinationArrayOutput
	ToSyncAzureDestinationArrayOutputWithContext(context.Context) SyncAzureDestinationArrayOutput
}

type SyncAzureDestinationArray []SyncAzureDestinationInput

func (SyncAzureDestinationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SyncAzureDestination)(nil)).Elem()
}

func (i SyncAzureDestinationArray) ToSyncAzureDestinationArrayOutput() SyncAzureDestinationArrayOutput {
	return i.ToSyncAzureDestinationArrayOutputWithContext(context.Background())
}

func (i SyncAzureDestinationArray) ToSyncAzureDestinationArrayOutputWithContext(ctx context.Context) SyncAzureDestinationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SyncAzureDestinationArrayOutput)
}

// SyncAzureDestinationMapInput is an input type that accepts SyncAzureDestinationMap and SyncAzureDestinationMapOutput values.
// You can construct a concrete instance of `SyncAzureDestinationMapInput` via:
//
//	SyncAzureDestinationMap{ "key": SyncAzureDestinationArgs{...} }
type SyncAzureDestinationMapInput interface {
	pulumi.Input

	ToSyncAzureDestinationMapOutput() SyncAzureDestinationMapOutput
	ToSyncAzureDestinationMapOutputWithContext(context.Context) SyncAzureDestinationMapOutput
}

type SyncAzureDestinationMap map[string]SyncAzureDestinationInput

func (SyncAzureDestinationMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SyncAzureDestination)(nil)).Elem()
}

func (i SyncAzureDestinationMap) ToSyncAzureDestinationMapOutput() SyncAzureDestinationMapOutput {
	return i.ToSyncAzureDestinationMapOutputWithContext(context.Background())
}

func (i SyncAzureDestinationMap) ToSyncAzureDestinationMapOutputWithContext(ctx context.Context) SyncAzureDestinationMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SyncAzureDestinationMapOutput)
}

type SyncAzureDestinationOutput struct{ *pulumi.OutputState }

func (SyncAzureDestinationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SyncAzureDestination)(nil)).Elem()
}

func (o SyncAzureDestinationOutput) ToSyncAzureDestinationOutput() SyncAzureDestinationOutput {
	return o
}

func (o SyncAzureDestinationOutput) ToSyncAzureDestinationOutputWithContext(ctx context.Context) SyncAzureDestinationOutput {
	return o
}

// Client ID of an Azure app registration.
// Can be omitted and directly provided to Vault using the `AZURE_CLIENT_ID` environment
// variable.
func (o SyncAzureDestinationOutput) ClientId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SyncAzureDestination) pulumi.StringPtrOutput { return v.ClientId }).(pulumi.StringPtrOutput)
}

// Client Secret of an Azure app registration.
// Can be omitted and directly provided to Vault using the `AZURE_CLIENT_SECRET` environment
// variable.
func (o SyncAzureDestinationOutput) ClientSecret() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SyncAzureDestination) pulumi.StringPtrOutput { return v.ClientSecret }).(pulumi.StringPtrOutput)
}

// Specifies a cloud for the client. The default is Azure Public Cloud.
func (o SyncAzureDestinationOutput) Cloud() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SyncAzureDestination) pulumi.StringPtrOutput { return v.Cloud }).(pulumi.StringPtrOutput)
}

// Custom tags to set on the secret managed at the destination.
func (o SyncAzureDestinationOutput) CustomTags() pulumi.MapOutput {
	return o.ApplyT(func(v *SyncAzureDestination) pulumi.MapOutput { return v.CustomTags }).(pulumi.MapOutput)
}

// Determines what level of information is synced as a distinct resource
// at the destination. Supports `secret-path` and `secret-key`.
func (o SyncAzureDestinationOutput) Granularity() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SyncAzureDestination) pulumi.StringPtrOutput { return v.Granularity }).(pulumi.StringPtrOutput)
}

// URI of an existing Azure Key Vault instance.
// Can be omitted and directly provided to Vault using the `KEY_VAULT_URI` environment
// variable.
func (o SyncAzureDestinationOutput) KeyVaultUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SyncAzureDestination) pulumi.StringPtrOutput { return v.KeyVaultUri }).(pulumi.StringPtrOutput)
}

// Unique name of the Azure destination.
func (o SyncAzureDestinationOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SyncAzureDestination) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The namespace to provision the resource in.
// The value should not contain leading or trailing forward slashes.
// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
func (o SyncAzureDestinationOutput) Namespace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SyncAzureDestination) pulumi.StringPtrOutput { return v.Namespace }).(pulumi.StringPtrOutput)
}

// Template describing how to generate external secret names.
// Supports a subset of the Go Template syntax.
func (o SyncAzureDestinationOutput) SecretNameTemplate() pulumi.StringOutput {
	return o.ApplyT(func(v *SyncAzureDestination) pulumi.StringOutput { return v.SecretNameTemplate }).(pulumi.StringOutput)
}

// ID of the target Azure tenant.
// Can be omitted and directly provided to Vault using the `AZURE_TENANT_ID` environment
// variable.
func (o SyncAzureDestinationOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SyncAzureDestination) pulumi.StringPtrOutput { return v.TenantId }).(pulumi.StringPtrOutput)
}

// The type of the secrets destination (`azure-kv`).
func (o SyncAzureDestinationOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *SyncAzureDestination) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

type SyncAzureDestinationArrayOutput struct{ *pulumi.OutputState }

func (SyncAzureDestinationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SyncAzureDestination)(nil)).Elem()
}

func (o SyncAzureDestinationArrayOutput) ToSyncAzureDestinationArrayOutput() SyncAzureDestinationArrayOutput {
	return o
}

func (o SyncAzureDestinationArrayOutput) ToSyncAzureDestinationArrayOutputWithContext(ctx context.Context) SyncAzureDestinationArrayOutput {
	return o
}

func (o SyncAzureDestinationArrayOutput) Index(i pulumi.IntInput) SyncAzureDestinationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SyncAzureDestination {
		return vs[0].([]*SyncAzureDestination)[vs[1].(int)]
	}).(SyncAzureDestinationOutput)
}

type SyncAzureDestinationMapOutput struct{ *pulumi.OutputState }

func (SyncAzureDestinationMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SyncAzureDestination)(nil)).Elem()
}

func (o SyncAzureDestinationMapOutput) ToSyncAzureDestinationMapOutput() SyncAzureDestinationMapOutput {
	return o
}

func (o SyncAzureDestinationMapOutput) ToSyncAzureDestinationMapOutputWithContext(ctx context.Context) SyncAzureDestinationMapOutput {
	return o
}

func (o SyncAzureDestinationMapOutput) MapIndex(k pulumi.StringInput) SyncAzureDestinationOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SyncAzureDestination {
		return vs[0].(map[string]*SyncAzureDestination)[vs[1].(string)]
	}).(SyncAzureDestinationOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SyncAzureDestinationInput)(nil)).Elem(), &SyncAzureDestination{})
	pulumi.RegisterInputType(reflect.TypeOf((*SyncAzureDestinationArrayInput)(nil)).Elem(), SyncAzureDestinationArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SyncAzureDestinationMapInput)(nil)).Elem(), SyncAzureDestinationMap{})
	pulumi.RegisterOutputType(SyncAzureDestinationOutput{})
	pulumi.RegisterOutputType(SyncAzureDestinationArrayOutput{})
	pulumi.RegisterOutputType(SyncAzureDestinationMapOutput{})
}
