// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package azure

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

type Backend struct {
	pulumi.CustomResourceState

	// - The OAuth2 client id to connect to Azure.
	ClientId pulumi.StringPtrOutput `pulumi:"clientId"`
	// - The OAuth2 client secret to connect to Azure.
	ClientSecret pulumi.StringPtrOutput `pulumi:"clientSecret"`
	// Human-friendly description of the mount for the backend.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// - The Azure environment.
	Environment pulumi.StringPtrOutput `pulumi:"environment"`
	// - The unique path this backend should be mounted at. Defaults to `azure`.
	Path pulumi.StringPtrOutput `pulumi:"path"`
	// - The subscription id for the Azure Active Directory.
	SubscriptionId pulumi.StringOutput `pulumi:"subscriptionId"`
	// - The tenant id for the Azure Active Directory.
	TenantId pulumi.StringOutput `pulumi:"tenantId"`
}

// NewBackend registers a new resource with the given unique name, arguments, and options.
func NewBackend(ctx *pulumi.Context,
	name string, args *BackendArgs, opts ...pulumi.ResourceOption) (*Backend, error) {
	if args == nil || args.SubscriptionId == nil {
		return nil, errors.New("missing required argument 'SubscriptionId'")
	}
	if args == nil || args.TenantId == nil {
		return nil, errors.New("missing required argument 'TenantId'")
	}
	if args == nil {
		args = &BackendArgs{}
	}
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
	// - The OAuth2 client id to connect to Azure.
	ClientId *string `pulumi:"clientId"`
	// - The OAuth2 client secret to connect to Azure.
	ClientSecret *string `pulumi:"clientSecret"`
	// Human-friendly description of the mount for the backend.
	Description *string `pulumi:"description"`
	// - The Azure environment.
	Environment *string `pulumi:"environment"`
	// - The unique path this backend should be mounted at. Defaults to `azure`.
	Path *string `pulumi:"path"`
	// - The subscription id for the Azure Active Directory.
	SubscriptionId *string `pulumi:"subscriptionId"`
	// - The tenant id for the Azure Active Directory.
	TenantId *string `pulumi:"tenantId"`
}

type BackendState struct {
	// - The OAuth2 client id to connect to Azure.
	ClientId pulumi.StringPtrInput
	// - The OAuth2 client secret to connect to Azure.
	ClientSecret pulumi.StringPtrInput
	// Human-friendly description of the mount for the backend.
	Description pulumi.StringPtrInput
	// - The Azure environment.
	Environment pulumi.StringPtrInput
	// - The unique path this backend should be mounted at. Defaults to `azure`.
	Path pulumi.StringPtrInput
	// - The subscription id for the Azure Active Directory.
	SubscriptionId pulumi.StringPtrInput
	// - The tenant id for the Azure Active Directory.
	TenantId pulumi.StringPtrInput
}

func (BackendState) ElementType() reflect.Type {
	return reflect.TypeOf((*backendState)(nil)).Elem()
}

type backendArgs struct {
	// - The OAuth2 client id to connect to Azure.
	ClientId *string `pulumi:"clientId"`
	// - The OAuth2 client secret to connect to Azure.
	ClientSecret *string `pulumi:"clientSecret"`
	// Human-friendly description of the mount for the backend.
	Description *string `pulumi:"description"`
	// - The Azure environment.
	Environment *string `pulumi:"environment"`
	// - The unique path this backend should be mounted at. Defaults to `azure`.
	Path *string `pulumi:"path"`
	// - The subscription id for the Azure Active Directory.
	SubscriptionId string `pulumi:"subscriptionId"`
	// - The tenant id for the Azure Active Directory.
	TenantId string `pulumi:"tenantId"`
}

// The set of arguments for constructing a Backend resource.
type BackendArgs struct {
	// - The OAuth2 client id to connect to Azure.
	ClientId pulumi.StringPtrInput
	// - The OAuth2 client secret to connect to Azure.
	ClientSecret pulumi.StringPtrInput
	// Human-friendly description of the mount for the backend.
	Description pulumi.StringPtrInput
	// - The Azure environment.
	Environment pulumi.StringPtrInput
	// - The unique path this backend should be mounted at. Defaults to `azure`.
	Path pulumi.StringPtrInput
	// - The subscription id for the Azure Active Directory.
	SubscriptionId pulumi.StringInput
	// - The tenant id for the Azure Active Directory.
	TenantId pulumi.StringInput
}

func (BackendArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*backendArgs)(nil)).Elem()
}

type BackendInput interface {
	pulumi.Input

	ToBackendOutput() BackendOutput
	ToBackendOutputWithContext(ctx context.Context) BackendOutput
}

func (Backend) ElementType() reflect.Type {
	return reflect.TypeOf((*Backend)(nil)).Elem()
}

func (i Backend) ToBackendOutput() BackendOutput {
	return i.ToBackendOutputWithContext(context.Background())
}

func (i Backend) ToBackendOutputWithContext(ctx context.Context) BackendOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BackendOutput)
}

type BackendOutput struct {
	*pulumi.OutputState
}

func (BackendOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BackendOutput)(nil)).Elem()
}

func (o BackendOutput) ToBackendOutput() BackendOutput {
	return o
}

func (o BackendOutput) ToBackendOutputWithContext(ctx context.Context) BackendOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(BackendOutput{})
}
