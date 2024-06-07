// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package azure

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-vault/sdk/v6/go/vault/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"fmt"
//
//	"github.com/pulumi/pulumi-vault/sdk/v6/go/vault/azure"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			azure, err := azure.NewBackend(ctx, "azure", &azure.BackendArgs{
//				SubscriptionId: pulumi.Any(subscriptionId),
//				TenantId:       pulumi.Any(tenantId),
//				ClientSecret:   pulumi.Any(clientSecret),
//				ClientId:       pulumi.Any(clientId),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = azure.NewBackendRole(ctx, "generated_role", &azure.BackendRoleArgs{
//				Backend:        azure.Path,
//				Role:           pulumi.String("generated_role"),
//				SignInAudience: pulumi.String("AzureADMyOrg"),
//				Tags: pulumi.StringArray{
//					pulumi.String("team:engineering"),
//					pulumi.String("environment:development"),
//				},
//				Ttl:    pulumi.String("300"),
//				MaxTtl: pulumi.String("600"),
//				AzureRoles: azure.BackendRoleAzureRoleArray{
//					&azure.BackendRoleAzureRoleArgs{
//						RoleName: pulumi.String("Reader"),
//						Scope:    pulumi.String(fmt.Sprintf("/subscriptions/%v/resourceGroups/azure-vault-group", subscriptionId)),
//					},
//				},
//			})
//			if err != nil {
//				return err
//			}
//			_, err = azure.NewBackendRole(ctx, "existing_object_id", &azure.BackendRoleArgs{
//				Backend:             azure.Path,
//				Role:                pulumi.String("existing_object_id"),
//				ApplicationObjectId: pulumi.String("11111111-2222-3333-4444-44444444444"),
//				Ttl:                 pulumi.String("300"),
//				MaxTtl:              pulumi.String("600"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
type BackendRole struct {
	pulumi.CustomResourceState

	// Application Object ID for an existing service principal that will
	// be used instead of creating dynamic service principals. If present, `azureRoles` and `permanentlyDelete` will be ignored.
	ApplicationObjectId pulumi.StringPtrOutput `pulumi:"applicationObjectId"`
	// List of Azure groups to be assigned to the generated service principal.
	AzureGroups BackendRoleAzureGroupArrayOutput `pulumi:"azureGroups"`
	// List of Azure roles to be assigned to the generated service principal.
	AzureRoles BackendRoleAzureRoleArrayOutput `pulumi:"azureRoles"`
	// Path to the mounted Azure auth backend
	Backend pulumi.StringPtrOutput `pulumi:"backend"`
	// Human-friendly description of the mount for the backend.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Specifies the maximum TTL for service principals generated using this role. Accepts time
	// suffixed strings ("1h") or an integer number of seconds. Defaults to the system/engine max TTL time.
	MaxTtl pulumi.StringPtrOutput `pulumi:"maxTtl"`
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
	// *Available only for Vault Enterprise*.
	Namespace pulumi.StringPtrOutput `pulumi:"namespace"`
	// Indicates whether the applications and service principals created by Vault will be permanently
	// deleted when the corresponding leases expire. Defaults to `false`. For Vault v1.12+.
	PermanentlyDelete pulumi.BoolOutput `pulumi:"permanentlyDelete"`
	// Name of the Azure role
	Role pulumi.StringOutput `pulumi:"role"`
	// Specifies the security principal types that are allowed to sign in to the application.
	// Valid values are: AzureADMyOrg, AzureADMultipleOrgs, AzureADandPersonalMicrosoftAccount, PersonalMicrosoftAccount. Requires Vault 1.16+.
	SignInAudience pulumi.StringPtrOutput `pulumi:"signInAudience"`
	// A list of Azure tags to attach to an application. Requires Vault 1.16+.
	Tags pulumi.StringArrayOutput `pulumi:"tags"`
	// Specifies the default TTL for service principals generated using this role.
	// Accepts time suffixed strings ("1h") or an integer number of seconds. Defaults to the system/engine default TTL time.
	Ttl pulumi.StringPtrOutput `pulumi:"ttl"`
}

// NewBackendRole registers a new resource with the given unique name, arguments, and options.
func NewBackendRole(ctx *pulumi.Context,
	name string, args *BackendRoleArgs, opts ...pulumi.ResourceOption) (*BackendRole, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Role == nil {
		return nil, errors.New("invalid value for required argument 'Role'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource BackendRole
	err := ctx.RegisterResource("vault:azure/backendRole:BackendRole", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBackendRole gets an existing BackendRole resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBackendRole(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BackendRoleState, opts ...pulumi.ResourceOption) (*BackendRole, error) {
	var resource BackendRole
	err := ctx.ReadResource("vault:azure/backendRole:BackendRole", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering BackendRole resources.
type backendRoleState struct {
	// Application Object ID for an existing service principal that will
	// be used instead of creating dynamic service principals. If present, `azureRoles` and `permanentlyDelete` will be ignored.
	ApplicationObjectId *string `pulumi:"applicationObjectId"`
	// List of Azure groups to be assigned to the generated service principal.
	AzureGroups []BackendRoleAzureGroup `pulumi:"azureGroups"`
	// List of Azure roles to be assigned to the generated service principal.
	AzureRoles []BackendRoleAzureRole `pulumi:"azureRoles"`
	// Path to the mounted Azure auth backend
	Backend *string `pulumi:"backend"`
	// Human-friendly description of the mount for the backend.
	Description *string `pulumi:"description"`
	// Specifies the maximum TTL for service principals generated using this role. Accepts time
	// suffixed strings ("1h") or an integer number of seconds. Defaults to the system/engine max TTL time.
	MaxTtl *string `pulumi:"maxTtl"`
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
	// *Available only for Vault Enterprise*.
	Namespace *string `pulumi:"namespace"`
	// Indicates whether the applications and service principals created by Vault will be permanently
	// deleted when the corresponding leases expire. Defaults to `false`. For Vault v1.12+.
	PermanentlyDelete *bool `pulumi:"permanentlyDelete"`
	// Name of the Azure role
	Role *string `pulumi:"role"`
	// Specifies the security principal types that are allowed to sign in to the application.
	// Valid values are: AzureADMyOrg, AzureADMultipleOrgs, AzureADandPersonalMicrosoftAccount, PersonalMicrosoftAccount. Requires Vault 1.16+.
	SignInAudience *string `pulumi:"signInAudience"`
	// A list of Azure tags to attach to an application. Requires Vault 1.16+.
	Tags []string `pulumi:"tags"`
	// Specifies the default TTL for service principals generated using this role.
	// Accepts time suffixed strings ("1h") or an integer number of seconds. Defaults to the system/engine default TTL time.
	Ttl *string `pulumi:"ttl"`
}

type BackendRoleState struct {
	// Application Object ID for an existing service principal that will
	// be used instead of creating dynamic service principals. If present, `azureRoles` and `permanentlyDelete` will be ignored.
	ApplicationObjectId pulumi.StringPtrInput
	// List of Azure groups to be assigned to the generated service principal.
	AzureGroups BackendRoleAzureGroupArrayInput
	// List of Azure roles to be assigned to the generated service principal.
	AzureRoles BackendRoleAzureRoleArrayInput
	// Path to the mounted Azure auth backend
	Backend pulumi.StringPtrInput
	// Human-friendly description of the mount for the backend.
	Description pulumi.StringPtrInput
	// Specifies the maximum TTL for service principals generated using this role. Accepts time
	// suffixed strings ("1h") or an integer number of seconds. Defaults to the system/engine max TTL time.
	MaxTtl pulumi.StringPtrInput
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
	// *Available only for Vault Enterprise*.
	Namespace pulumi.StringPtrInput
	// Indicates whether the applications and service principals created by Vault will be permanently
	// deleted when the corresponding leases expire. Defaults to `false`. For Vault v1.12+.
	PermanentlyDelete pulumi.BoolPtrInput
	// Name of the Azure role
	Role pulumi.StringPtrInput
	// Specifies the security principal types that are allowed to sign in to the application.
	// Valid values are: AzureADMyOrg, AzureADMultipleOrgs, AzureADandPersonalMicrosoftAccount, PersonalMicrosoftAccount. Requires Vault 1.16+.
	SignInAudience pulumi.StringPtrInput
	// A list of Azure tags to attach to an application. Requires Vault 1.16+.
	Tags pulumi.StringArrayInput
	// Specifies the default TTL for service principals generated using this role.
	// Accepts time suffixed strings ("1h") or an integer number of seconds. Defaults to the system/engine default TTL time.
	Ttl pulumi.StringPtrInput
}

func (BackendRoleState) ElementType() reflect.Type {
	return reflect.TypeOf((*backendRoleState)(nil)).Elem()
}

type backendRoleArgs struct {
	// Application Object ID for an existing service principal that will
	// be used instead of creating dynamic service principals. If present, `azureRoles` and `permanentlyDelete` will be ignored.
	ApplicationObjectId *string `pulumi:"applicationObjectId"`
	// List of Azure groups to be assigned to the generated service principal.
	AzureGroups []BackendRoleAzureGroup `pulumi:"azureGroups"`
	// List of Azure roles to be assigned to the generated service principal.
	AzureRoles []BackendRoleAzureRole `pulumi:"azureRoles"`
	// Path to the mounted Azure auth backend
	Backend *string `pulumi:"backend"`
	// Human-friendly description of the mount for the backend.
	Description *string `pulumi:"description"`
	// Specifies the maximum TTL for service principals generated using this role. Accepts time
	// suffixed strings ("1h") or an integer number of seconds. Defaults to the system/engine max TTL time.
	MaxTtl *string `pulumi:"maxTtl"`
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
	// *Available only for Vault Enterprise*.
	Namespace *string `pulumi:"namespace"`
	// Indicates whether the applications and service principals created by Vault will be permanently
	// deleted when the corresponding leases expire. Defaults to `false`. For Vault v1.12+.
	PermanentlyDelete *bool `pulumi:"permanentlyDelete"`
	// Name of the Azure role
	Role string `pulumi:"role"`
	// Specifies the security principal types that are allowed to sign in to the application.
	// Valid values are: AzureADMyOrg, AzureADMultipleOrgs, AzureADandPersonalMicrosoftAccount, PersonalMicrosoftAccount. Requires Vault 1.16+.
	SignInAudience *string `pulumi:"signInAudience"`
	// A list of Azure tags to attach to an application. Requires Vault 1.16+.
	Tags []string `pulumi:"tags"`
	// Specifies the default TTL for service principals generated using this role.
	// Accepts time suffixed strings ("1h") or an integer number of seconds. Defaults to the system/engine default TTL time.
	Ttl *string `pulumi:"ttl"`
}

// The set of arguments for constructing a BackendRole resource.
type BackendRoleArgs struct {
	// Application Object ID for an existing service principal that will
	// be used instead of creating dynamic service principals. If present, `azureRoles` and `permanentlyDelete` will be ignored.
	ApplicationObjectId pulumi.StringPtrInput
	// List of Azure groups to be assigned to the generated service principal.
	AzureGroups BackendRoleAzureGroupArrayInput
	// List of Azure roles to be assigned to the generated service principal.
	AzureRoles BackendRoleAzureRoleArrayInput
	// Path to the mounted Azure auth backend
	Backend pulumi.StringPtrInput
	// Human-friendly description of the mount for the backend.
	Description pulumi.StringPtrInput
	// Specifies the maximum TTL for service principals generated using this role. Accepts time
	// suffixed strings ("1h") or an integer number of seconds. Defaults to the system/engine max TTL time.
	MaxTtl pulumi.StringPtrInput
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
	// *Available only for Vault Enterprise*.
	Namespace pulumi.StringPtrInput
	// Indicates whether the applications and service principals created by Vault will be permanently
	// deleted when the corresponding leases expire. Defaults to `false`. For Vault v1.12+.
	PermanentlyDelete pulumi.BoolPtrInput
	// Name of the Azure role
	Role pulumi.StringInput
	// Specifies the security principal types that are allowed to sign in to the application.
	// Valid values are: AzureADMyOrg, AzureADMultipleOrgs, AzureADandPersonalMicrosoftAccount, PersonalMicrosoftAccount. Requires Vault 1.16+.
	SignInAudience pulumi.StringPtrInput
	// A list of Azure tags to attach to an application. Requires Vault 1.16+.
	Tags pulumi.StringArrayInput
	// Specifies the default TTL for service principals generated using this role.
	// Accepts time suffixed strings ("1h") or an integer number of seconds. Defaults to the system/engine default TTL time.
	Ttl pulumi.StringPtrInput
}

func (BackendRoleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*backendRoleArgs)(nil)).Elem()
}

type BackendRoleInput interface {
	pulumi.Input

	ToBackendRoleOutput() BackendRoleOutput
	ToBackendRoleOutputWithContext(ctx context.Context) BackendRoleOutput
}

func (*BackendRole) ElementType() reflect.Type {
	return reflect.TypeOf((**BackendRole)(nil)).Elem()
}

func (i *BackendRole) ToBackendRoleOutput() BackendRoleOutput {
	return i.ToBackendRoleOutputWithContext(context.Background())
}

func (i *BackendRole) ToBackendRoleOutputWithContext(ctx context.Context) BackendRoleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BackendRoleOutput)
}

// BackendRoleArrayInput is an input type that accepts BackendRoleArray and BackendRoleArrayOutput values.
// You can construct a concrete instance of `BackendRoleArrayInput` via:
//
//	BackendRoleArray{ BackendRoleArgs{...} }
type BackendRoleArrayInput interface {
	pulumi.Input

	ToBackendRoleArrayOutput() BackendRoleArrayOutput
	ToBackendRoleArrayOutputWithContext(context.Context) BackendRoleArrayOutput
}

type BackendRoleArray []BackendRoleInput

func (BackendRoleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*BackendRole)(nil)).Elem()
}

func (i BackendRoleArray) ToBackendRoleArrayOutput() BackendRoleArrayOutput {
	return i.ToBackendRoleArrayOutputWithContext(context.Background())
}

func (i BackendRoleArray) ToBackendRoleArrayOutputWithContext(ctx context.Context) BackendRoleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BackendRoleArrayOutput)
}

// BackendRoleMapInput is an input type that accepts BackendRoleMap and BackendRoleMapOutput values.
// You can construct a concrete instance of `BackendRoleMapInput` via:
//
//	BackendRoleMap{ "key": BackendRoleArgs{...} }
type BackendRoleMapInput interface {
	pulumi.Input

	ToBackendRoleMapOutput() BackendRoleMapOutput
	ToBackendRoleMapOutputWithContext(context.Context) BackendRoleMapOutput
}

type BackendRoleMap map[string]BackendRoleInput

func (BackendRoleMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*BackendRole)(nil)).Elem()
}

func (i BackendRoleMap) ToBackendRoleMapOutput() BackendRoleMapOutput {
	return i.ToBackendRoleMapOutputWithContext(context.Background())
}

func (i BackendRoleMap) ToBackendRoleMapOutputWithContext(ctx context.Context) BackendRoleMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BackendRoleMapOutput)
}

type BackendRoleOutput struct{ *pulumi.OutputState }

func (BackendRoleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BackendRole)(nil)).Elem()
}

func (o BackendRoleOutput) ToBackendRoleOutput() BackendRoleOutput {
	return o
}

func (o BackendRoleOutput) ToBackendRoleOutputWithContext(ctx context.Context) BackendRoleOutput {
	return o
}

// Application Object ID for an existing service principal that will
// be used instead of creating dynamic service principals. If present, `azureRoles` and `permanentlyDelete` will be ignored.
func (o BackendRoleOutput) ApplicationObjectId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BackendRole) pulumi.StringPtrOutput { return v.ApplicationObjectId }).(pulumi.StringPtrOutput)
}

// List of Azure groups to be assigned to the generated service principal.
func (o BackendRoleOutput) AzureGroups() BackendRoleAzureGroupArrayOutput {
	return o.ApplyT(func(v *BackendRole) BackendRoleAzureGroupArrayOutput { return v.AzureGroups }).(BackendRoleAzureGroupArrayOutput)
}

// List of Azure roles to be assigned to the generated service principal.
func (o BackendRoleOutput) AzureRoles() BackendRoleAzureRoleArrayOutput {
	return o.ApplyT(func(v *BackendRole) BackendRoleAzureRoleArrayOutput { return v.AzureRoles }).(BackendRoleAzureRoleArrayOutput)
}

// Path to the mounted Azure auth backend
func (o BackendRoleOutput) Backend() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BackendRole) pulumi.StringPtrOutput { return v.Backend }).(pulumi.StringPtrOutput)
}

// Human-friendly description of the mount for the backend.
func (o BackendRoleOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BackendRole) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Specifies the maximum TTL for service principals generated using this role. Accepts time
// suffixed strings ("1h") or an integer number of seconds. Defaults to the system/engine max TTL time.
func (o BackendRoleOutput) MaxTtl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BackendRole) pulumi.StringPtrOutput { return v.MaxTtl }).(pulumi.StringPtrOutput)
}

// The namespace to provision the resource in.
// The value should not contain leading or trailing forward slashes.
// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
// *Available only for Vault Enterprise*.
func (o BackendRoleOutput) Namespace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BackendRole) pulumi.StringPtrOutput { return v.Namespace }).(pulumi.StringPtrOutput)
}

// Indicates whether the applications and service principals created by Vault will be permanently
// deleted when the corresponding leases expire. Defaults to `false`. For Vault v1.12+.
func (o BackendRoleOutput) PermanentlyDelete() pulumi.BoolOutput {
	return o.ApplyT(func(v *BackendRole) pulumi.BoolOutput { return v.PermanentlyDelete }).(pulumi.BoolOutput)
}

// Name of the Azure role
func (o BackendRoleOutput) Role() pulumi.StringOutput {
	return o.ApplyT(func(v *BackendRole) pulumi.StringOutput { return v.Role }).(pulumi.StringOutput)
}

// Specifies the security principal types that are allowed to sign in to the application.
// Valid values are: AzureADMyOrg, AzureADMultipleOrgs, AzureADandPersonalMicrosoftAccount, PersonalMicrosoftAccount. Requires Vault 1.16+.
func (o BackendRoleOutput) SignInAudience() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BackendRole) pulumi.StringPtrOutput { return v.SignInAudience }).(pulumi.StringPtrOutput)
}

// A list of Azure tags to attach to an application. Requires Vault 1.16+.
func (o BackendRoleOutput) Tags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *BackendRole) pulumi.StringArrayOutput { return v.Tags }).(pulumi.StringArrayOutput)
}

// Specifies the default TTL for service principals generated using this role.
// Accepts time suffixed strings ("1h") or an integer number of seconds. Defaults to the system/engine default TTL time.
func (o BackendRoleOutput) Ttl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BackendRole) pulumi.StringPtrOutput { return v.Ttl }).(pulumi.StringPtrOutput)
}

type BackendRoleArrayOutput struct{ *pulumi.OutputState }

func (BackendRoleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*BackendRole)(nil)).Elem()
}

func (o BackendRoleArrayOutput) ToBackendRoleArrayOutput() BackendRoleArrayOutput {
	return o
}

func (o BackendRoleArrayOutput) ToBackendRoleArrayOutputWithContext(ctx context.Context) BackendRoleArrayOutput {
	return o
}

func (o BackendRoleArrayOutput) Index(i pulumi.IntInput) BackendRoleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *BackendRole {
		return vs[0].([]*BackendRole)[vs[1].(int)]
	}).(BackendRoleOutput)
}

type BackendRoleMapOutput struct{ *pulumi.OutputState }

func (BackendRoleMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*BackendRole)(nil)).Elem()
}

func (o BackendRoleMapOutput) ToBackendRoleMapOutput() BackendRoleMapOutput {
	return o
}

func (o BackendRoleMapOutput) ToBackendRoleMapOutputWithContext(ctx context.Context) BackendRoleMapOutput {
	return o
}

func (o BackendRoleMapOutput) MapIndex(k pulumi.StringInput) BackendRoleOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *BackendRole {
		return vs[0].(map[string]*BackendRole)[vs[1].(string)]
	}).(BackendRoleOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*BackendRoleInput)(nil)).Elem(), &BackendRole{})
	pulumi.RegisterInputType(reflect.TypeOf((*BackendRoleArrayInput)(nil)).Elem(), BackendRoleArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*BackendRoleMapInput)(nil)).Elem(), BackendRoleMap{})
	pulumi.RegisterOutputType(BackendRoleOutput{})
	pulumi.RegisterOutputType(BackendRoleArrayOutput{})
	pulumi.RegisterOutputType(BackendRoleMapOutput{})
}
