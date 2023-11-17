// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package azure

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-vault/sdk/v5/go/vault/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## Example Usage
// ### *Vault-1.9 And Above*
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-vault/sdk/v5/go/vault/azure"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := azure.NewBackend(ctx, "azure", &azure.BackendArgs{
//				ClientId:             pulumi.String("11111111-2222-3333-4444-333333333333"),
//				ClientSecret:         pulumi.String("12345678901234567890"),
//				Environment:          pulumi.String("AzurePublicCloud"),
//				SubscriptionId:       pulumi.String("11111111-2222-3333-4444-111111111111"),
//				TenantId:             pulumi.String("11111111-2222-3333-4444-222222222222"),
//				UseMicrosoftGraphApi: pulumi.Bool(true),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// ### *Vault-1.8 And Below*
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-vault/sdk/v5/go/vault/azure"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := azure.NewBackend(ctx, "azure", &azure.BackendArgs{
//				ClientId:             pulumi.String("11111111-2222-3333-4444-333333333333"),
//				ClientSecret:         pulumi.String("12345678901234567890"),
//				Environment:          pulumi.String("AzurePublicCloud"),
//				SubscriptionId:       pulumi.String("11111111-2222-3333-4444-111111111111"),
//				TenantId:             pulumi.String("11111111-2222-3333-4444-222222222222"),
//				UseMicrosoftGraphApi: pulumi.Bool(false),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
type Backend struct {
	pulumi.CustomResourceState

	// (`string:""`) - The OAuth2 client id to connect to Azure.
	ClientId pulumi.StringPtrOutput `pulumi:"clientId"`
	// (`string:""`) - The OAuth2 client secret to connect to Azure.
	ClientSecret pulumi.StringPtrOutput `pulumi:"clientSecret"`
	// Human-friendly description of the mount for the backend.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// If set, opts out of mount migration on path updates.
	// See here for more info on [Mount Migration](https://www.vaultproject.io/docs/concepts/mount-migration)
	DisableRemount pulumi.BoolPtrOutput `pulumi:"disableRemount"`
	// (`string:""`) - The Azure environment.
	Environment pulumi.StringPtrOutput `pulumi:"environment"`
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
	// *Available only for Vault Enterprise*.
	Namespace pulumi.StringPtrOutput `pulumi:"namespace"`
	// (`string: <optional>`) - The unique path this backend should be mounted at. Defaults to `azure`.
	Path pulumi.StringPtrOutput `pulumi:"path"`
	// (`string: <required>`) - The subscription id for the Azure Active Directory.
	SubscriptionId pulumi.StringOutput `pulumi:"subscriptionId"`
	// (`string: <required>`) - The tenant id for the Azure Active Directory.
	TenantId pulumi.StringOutput `pulumi:"tenantId"`
	// (`bool: <optional>`) - Indicates whether the secrets engine should use
	// the Microsoft Graph API. This parameter has been deprecated and will be ignored in `vault-1.12+`.
	// For more information, please refer to the [Vault docs](https://developer.hashicorp.com/vault/api-docs/secret/azure#use_microsoft_graph_api)
	UseMicrosoftGraphApi pulumi.BoolOutput `pulumi:"useMicrosoftGraphApi"`
}

// NewBackend registers a new resource with the given unique name, arguments, and options.
func NewBackend(ctx *pulumi.Context,
	name string, args *BackendArgs, opts ...pulumi.ResourceOption) (*Backend, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.SubscriptionId == nil {
		return nil, errors.New("invalid value for required argument 'SubscriptionId'")
	}
	if args.TenantId == nil {
		return nil, errors.New("invalid value for required argument 'TenantId'")
	}
	if args.ClientId != nil {
		args.ClientId = pulumi.ToSecret(args.ClientId).(pulumi.StringPtrInput)
	}
	if args.ClientSecret != nil {
		args.ClientSecret = pulumi.ToSecret(args.ClientSecret).(pulumi.StringPtrInput)
	}
	if args.SubscriptionId != nil {
		args.SubscriptionId = pulumi.ToSecret(args.SubscriptionId).(pulumi.StringInput)
	}
	if args.TenantId != nil {
		args.TenantId = pulumi.ToSecret(args.TenantId).(pulumi.StringInput)
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"clientId",
		"clientSecret",
		"subscriptionId",
		"tenantId",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Backend
	err := ctx.RegisterResource("vault:azure/backend:Backend", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBackend gets an existing Backend resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBackend(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BackendState, opts ...pulumi.ResourceOption) (*Backend, error) {
	var resource Backend
	err := ctx.ReadResource("vault:azure/backend:Backend", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Backend resources.
type backendState struct {
	// (`string:""`) - The OAuth2 client id to connect to Azure.
	ClientId *string `pulumi:"clientId"`
	// (`string:""`) - The OAuth2 client secret to connect to Azure.
	ClientSecret *string `pulumi:"clientSecret"`
	// Human-friendly description of the mount for the backend.
	Description *string `pulumi:"description"`
	// If set, opts out of mount migration on path updates.
	// See here for more info on [Mount Migration](https://www.vaultproject.io/docs/concepts/mount-migration)
	DisableRemount *bool `pulumi:"disableRemount"`
	// (`string:""`) - The Azure environment.
	Environment *string `pulumi:"environment"`
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
	// *Available only for Vault Enterprise*.
	Namespace *string `pulumi:"namespace"`
	// (`string: <optional>`) - The unique path this backend should be mounted at. Defaults to `azure`.
	Path *string `pulumi:"path"`
	// (`string: <required>`) - The subscription id for the Azure Active Directory.
	SubscriptionId *string `pulumi:"subscriptionId"`
	// (`string: <required>`) - The tenant id for the Azure Active Directory.
	TenantId *string `pulumi:"tenantId"`
	// (`bool: <optional>`) - Indicates whether the secrets engine should use
	// the Microsoft Graph API. This parameter has been deprecated and will be ignored in `vault-1.12+`.
	// For more information, please refer to the [Vault docs](https://developer.hashicorp.com/vault/api-docs/secret/azure#use_microsoft_graph_api)
	UseMicrosoftGraphApi *bool `pulumi:"useMicrosoftGraphApi"`
}

type BackendState struct {
	// (`string:""`) - The OAuth2 client id to connect to Azure.
	ClientId pulumi.StringPtrInput
	// (`string:""`) - The OAuth2 client secret to connect to Azure.
	ClientSecret pulumi.StringPtrInput
	// Human-friendly description of the mount for the backend.
	Description pulumi.StringPtrInput
	// If set, opts out of mount migration on path updates.
	// See here for more info on [Mount Migration](https://www.vaultproject.io/docs/concepts/mount-migration)
	DisableRemount pulumi.BoolPtrInput
	// (`string:""`) - The Azure environment.
	Environment pulumi.StringPtrInput
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
	// *Available only for Vault Enterprise*.
	Namespace pulumi.StringPtrInput
	// (`string: <optional>`) - The unique path this backend should be mounted at. Defaults to `azure`.
	Path pulumi.StringPtrInput
	// (`string: <required>`) - The subscription id for the Azure Active Directory.
	SubscriptionId pulumi.StringPtrInput
	// (`string: <required>`) - The tenant id for the Azure Active Directory.
	TenantId pulumi.StringPtrInput
	// (`bool: <optional>`) - Indicates whether the secrets engine should use
	// the Microsoft Graph API. This parameter has been deprecated and will be ignored in `vault-1.12+`.
	// For more information, please refer to the [Vault docs](https://developer.hashicorp.com/vault/api-docs/secret/azure#use_microsoft_graph_api)
	UseMicrosoftGraphApi pulumi.BoolPtrInput
}

func (BackendState) ElementType() reflect.Type {
	return reflect.TypeOf((*backendState)(nil)).Elem()
}

type backendArgs struct {
	// (`string:""`) - The OAuth2 client id to connect to Azure.
	ClientId *string `pulumi:"clientId"`
	// (`string:""`) - The OAuth2 client secret to connect to Azure.
	ClientSecret *string `pulumi:"clientSecret"`
	// Human-friendly description of the mount for the backend.
	Description *string `pulumi:"description"`
	// If set, opts out of mount migration on path updates.
	// See here for more info on [Mount Migration](https://www.vaultproject.io/docs/concepts/mount-migration)
	DisableRemount *bool `pulumi:"disableRemount"`
	// (`string:""`) - The Azure environment.
	Environment *string `pulumi:"environment"`
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
	// *Available only for Vault Enterprise*.
	Namespace *string `pulumi:"namespace"`
	// (`string: <optional>`) - The unique path this backend should be mounted at. Defaults to `azure`.
	Path *string `pulumi:"path"`
	// (`string: <required>`) - The subscription id for the Azure Active Directory.
	SubscriptionId string `pulumi:"subscriptionId"`
	// (`string: <required>`) - The tenant id for the Azure Active Directory.
	TenantId string `pulumi:"tenantId"`
	// (`bool: <optional>`) - Indicates whether the secrets engine should use
	// the Microsoft Graph API. This parameter has been deprecated and will be ignored in `vault-1.12+`.
	// For more information, please refer to the [Vault docs](https://developer.hashicorp.com/vault/api-docs/secret/azure#use_microsoft_graph_api)
	UseMicrosoftGraphApi *bool `pulumi:"useMicrosoftGraphApi"`
}

// The set of arguments for constructing a Backend resource.
type BackendArgs struct {
	// (`string:""`) - The OAuth2 client id to connect to Azure.
	ClientId pulumi.StringPtrInput
	// (`string:""`) - The OAuth2 client secret to connect to Azure.
	ClientSecret pulumi.StringPtrInput
	// Human-friendly description of the mount for the backend.
	Description pulumi.StringPtrInput
	// If set, opts out of mount migration on path updates.
	// See here for more info on [Mount Migration](https://www.vaultproject.io/docs/concepts/mount-migration)
	DisableRemount pulumi.BoolPtrInput
	// (`string:""`) - The Azure environment.
	Environment pulumi.StringPtrInput
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
	// *Available only for Vault Enterprise*.
	Namespace pulumi.StringPtrInput
	// (`string: <optional>`) - The unique path this backend should be mounted at. Defaults to `azure`.
	Path pulumi.StringPtrInput
	// (`string: <required>`) - The subscription id for the Azure Active Directory.
	SubscriptionId pulumi.StringInput
	// (`string: <required>`) - The tenant id for the Azure Active Directory.
	TenantId pulumi.StringInput
	// (`bool: <optional>`) - Indicates whether the secrets engine should use
	// the Microsoft Graph API. This parameter has been deprecated and will be ignored in `vault-1.12+`.
	// For more information, please refer to the [Vault docs](https://developer.hashicorp.com/vault/api-docs/secret/azure#use_microsoft_graph_api)
	UseMicrosoftGraphApi pulumi.BoolPtrInput
}

func (BackendArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*backendArgs)(nil)).Elem()
}

type BackendInput interface {
	pulumi.Input

	ToBackendOutput() BackendOutput
	ToBackendOutputWithContext(ctx context.Context) BackendOutput
}

func (*Backend) ElementType() reflect.Type {
	return reflect.TypeOf((**Backend)(nil)).Elem()
}

func (i *Backend) ToBackendOutput() BackendOutput {
	return i.ToBackendOutputWithContext(context.Background())
}

func (i *Backend) ToBackendOutputWithContext(ctx context.Context) BackendOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BackendOutput)
}

// BackendArrayInput is an input type that accepts BackendArray and BackendArrayOutput values.
// You can construct a concrete instance of `BackendArrayInput` via:
//
//	BackendArray{ BackendArgs{...} }
type BackendArrayInput interface {
	pulumi.Input

	ToBackendArrayOutput() BackendArrayOutput
	ToBackendArrayOutputWithContext(context.Context) BackendArrayOutput
}

type BackendArray []BackendInput

func (BackendArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Backend)(nil)).Elem()
}

func (i BackendArray) ToBackendArrayOutput() BackendArrayOutput {
	return i.ToBackendArrayOutputWithContext(context.Background())
}

func (i BackendArray) ToBackendArrayOutputWithContext(ctx context.Context) BackendArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BackendArrayOutput)
}

// BackendMapInput is an input type that accepts BackendMap and BackendMapOutput values.
// You can construct a concrete instance of `BackendMapInput` via:
//
//	BackendMap{ "key": BackendArgs{...} }
type BackendMapInput interface {
	pulumi.Input

	ToBackendMapOutput() BackendMapOutput
	ToBackendMapOutputWithContext(context.Context) BackendMapOutput
}

type BackendMap map[string]BackendInput

func (BackendMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Backend)(nil)).Elem()
}

func (i BackendMap) ToBackendMapOutput() BackendMapOutput {
	return i.ToBackendMapOutputWithContext(context.Background())
}

func (i BackendMap) ToBackendMapOutputWithContext(ctx context.Context) BackendMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BackendMapOutput)
}

type BackendOutput struct{ *pulumi.OutputState }

func (BackendOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Backend)(nil)).Elem()
}

func (o BackendOutput) ToBackendOutput() BackendOutput {
	return o
}

func (o BackendOutput) ToBackendOutputWithContext(ctx context.Context) BackendOutput {
	return o
}

// (`string:""`) - The OAuth2 client id to connect to Azure.
func (o BackendOutput) ClientId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Backend) pulumi.StringPtrOutput { return v.ClientId }).(pulumi.StringPtrOutput)
}

// (`string:""`) - The OAuth2 client secret to connect to Azure.
func (o BackendOutput) ClientSecret() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Backend) pulumi.StringPtrOutput { return v.ClientSecret }).(pulumi.StringPtrOutput)
}

// Human-friendly description of the mount for the backend.
func (o BackendOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Backend) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// If set, opts out of mount migration on path updates.
// See here for more info on [Mount Migration](https://www.vaultproject.io/docs/concepts/mount-migration)
func (o BackendOutput) DisableRemount() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Backend) pulumi.BoolPtrOutput { return v.DisableRemount }).(pulumi.BoolPtrOutput)
}

// (`string:""`) - The Azure environment.
func (o BackendOutput) Environment() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Backend) pulumi.StringPtrOutput { return v.Environment }).(pulumi.StringPtrOutput)
}

// The namespace to provision the resource in.
// The value should not contain leading or trailing forward slashes.
// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
// *Available only for Vault Enterprise*.
func (o BackendOutput) Namespace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Backend) pulumi.StringPtrOutput { return v.Namespace }).(pulumi.StringPtrOutput)
}

// (`string: <optional>`) - The unique path this backend should be mounted at. Defaults to `azure`.
func (o BackendOutput) Path() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Backend) pulumi.StringPtrOutput { return v.Path }).(pulumi.StringPtrOutput)
}

// (`string: <required>`) - The subscription id for the Azure Active Directory.
func (o BackendOutput) SubscriptionId() pulumi.StringOutput {
	return o.ApplyT(func(v *Backend) pulumi.StringOutput { return v.SubscriptionId }).(pulumi.StringOutput)
}

// (`string: <required>`) - The tenant id for the Azure Active Directory.
func (o BackendOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v *Backend) pulumi.StringOutput { return v.TenantId }).(pulumi.StringOutput)
}

// (`bool: <optional>`) - Indicates whether the secrets engine should use
// the Microsoft Graph API. This parameter has been deprecated and will be ignored in `vault-1.12+`.
// For more information, please refer to the [Vault docs](https://developer.hashicorp.com/vault/api-docs/secret/azure#use_microsoft_graph_api)
func (o BackendOutput) UseMicrosoftGraphApi() pulumi.BoolOutput {
	return o.ApplyT(func(v *Backend) pulumi.BoolOutput { return v.UseMicrosoftGraphApi }).(pulumi.BoolOutput)
}

type BackendArrayOutput struct{ *pulumi.OutputState }

func (BackendArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Backend)(nil)).Elem()
}

func (o BackendArrayOutput) ToBackendArrayOutput() BackendArrayOutput {
	return o
}

func (o BackendArrayOutput) ToBackendArrayOutputWithContext(ctx context.Context) BackendArrayOutput {
	return o
}

func (o BackendArrayOutput) Index(i pulumi.IntInput) BackendOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Backend {
		return vs[0].([]*Backend)[vs[1].(int)]
	}).(BackendOutput)
}

type BackendMapOutput struct{ *pulumi.OutputState }

func (BackendMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Backend)(nil)).Elem()
}

func (o BackendMapOutput) ToBackendMapOutput() BackendMapOutput {
	return o
}

func (o BackendMapOutput) ToBackendMapOutputWithContext(ctx context.Context) BackendMapOutput {
	return o
}

func (o BackendMapOutput) MapIndex(k pulumi.StringInput) BackendOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Backend {
		return vs[0].(map[string]*Backend)[vs[1].(string)]
	}).(BackendOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*BackendInput)(nil)).Elem(), &Backend{})
	pulumi.RegisterInputType(reflect.TypeOf((*BackendArrayInput)(nil)).Elem(), BackendArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*BackendMapInput)(nil)).Elem(), BackendMap{})
	pulumi.RegisterOutputType(BackendOutput{})
	pulumi.RegisterOutputType(BackendArrayOutput{})
	pulumi.RegisterOutputType(BackendMapOutput{})
}
