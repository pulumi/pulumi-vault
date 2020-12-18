// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package vault

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

type MfaDuo struct {
	pulumi.CustomResourceState

	// API hostname for Duo.
	ApiHostname pulumi.StringOutput `pulumi:"apiHostname"`
	// Integration key for Duo.
	IntegrationKey pulumi.StringOutput `pulumi:"integrationKey"`
	// The mount to tie this method to for use in automatic mappings. The mapping will use the Name field of Aliases associated
	// with this mount as the username in the mapping.
	MountAccessor pulumi.StringOutput `pulumi:"mountAccessor"`
	// Name of the MFA method.
	Name pulumi.StringOutput `pulumi:"name"`
	// Push information for Duo.
	PushInfo pulumi.StringPtrOutput `pulumi:"pushInfo"`
	// Secret key for Duo.
	SecretKey pulumi.StringOutput `pulumi:"secretKey"`
	// A format string for mapping Identity names to MFA method names. Values to substitute should be placed in `{{}}`.
	UsernameFormat pulumi.StringPtrOutput `pulumi:"usernameFormat"`
}

// NewMfaDuo registers a new resource with the given unique name, arguments, and options.
func NewMfaDuo(ctx *pulumi.Context,
	name string, args *MfaDuoArgs, opts ...pulumi.ResourceOption) (*MfaDuo, error) {
	if args == nil || args.ApiHostname == nil {
		return nil, errors.New("missing required argument 'ApiHostname'")
	}
	if args == nil || args.IntegrationKey == nil {
		return nil, errors.New("missing required argument 'IntegrationKey'")
	}
	if args == nil || args.MountAccessor == nil {
		return nil, errors.New("missing required argument 'MountAccessor'")
	}
	if args == nil || args.SecretKey == nil {
		return nil, errors.New("missing required argument 'SecretKey'")
	}
	if args == nil {
		args = &MfaDuoArgs{}
	}
	var resource MfaDuo
	err := ctx.RegisterResource("vault:index/mfaDuo:MfaDuo", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetMfaDuo gets an existing MfaDuo resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetMfaDuo(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MfaDuoState, opts ...pulumi.ResourceOption) (*MfaDuo, error) {
	var resource MfaDuo
	err := ctx.ReadResource("vault:index/mfaDuo:MfaDuo", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering MfaDuo resources.
type mfaDuoState struct {
	// API hostname for Duo.
	ApiHostname *string `pulumi:"apiHostname"`
	// Integration key for Duo.
	IntegrationKey *string `pulumi:"integrationKey"`
	// The mount to tie this method to for use in automatic mappings. The mapping will use the Name field of Aliases associated
	// with this mount as the username in the mapping.
	MountAccessor *string `pulumi:"mountAccessor"`
	// Name of the MFA method.
	Name *string `pulumi:"name"`
	// Push information for Duo.
	PushInfo *string `pulumi:"pushInfo"`
	// Secret key for Duo.
	SecretKey *string `pulumi:"secretKey"`
	// A format string for mapping Identity names to MFA method names. Values to substitute should be placed in `{{}}`.
	UsernameFormat *string `pulumi:"usernameFormat"`
}

type MfaDuoState struct {
	// API hostname for Duo.
	ApiHostname pulumi.StringPtrInput
	// Integration key for Duo.
	IntegrationKey pulumi.StringPtrInput
	// The mount to tie this method to for use in automatic mappings. The mapping will use the Name field of Aliases associated
	// with this mount as the username in the mapping.
	MountAccessor pulumi.StringPtrInput
	// Name of the MFA method.
	Name pulumi.StringPtrInput
	// Push information for Duo.
	PushInfo pulumi.StringPtrInput
	// Secret key for Duo.
	SecretKey pulumi.StringPtrInput
	// A format string for mapping Identity names to MFA method names. Values to substitute should be placed in `{{}}`.
	UsernameFormat pulumi.StringPtrInput
}

func (MfaDuoState) ElementType() reflect.Type {
	return reflect.TypeOf((*mfaDuoState)(nil)).Elem()
}

type mfaDuoArgs struct {
	// API hostname for Duo.
	ApiHostname string `pulumi:"apiHostname"`
	// Integration key for Duo.
	IntegrationKey string `pulumi:"integrationKey"`
	// The mount to tie this method to for use in automatic mappings. The mapping will use the Name field of Aliases associated
	// with this mount as the username in the mapping.
	MountAccessor string `pulumi:"mountAccessor"`
	// Name of the MFA method.
	Name *string `pulumi:"name"`
	// Push information for Duo.
	PushInfo *string `pulumi:"pushInfo"`
	// Secret key for Duo.
	SecretKey string `pulumi:"secretKey"`
	// A format string for mapping Identity names to MFA method names. Values to substitute should be placed in `{{}}`.
	UsernameFormat *string `pulumi:"usernameFormat"`
}

// The set of arguments for constructing a MfaDuo resource.
type MfaDuoArgs struct {
	// API hostname for Duo.
	ApiHostname pulumi.StringInput
	// Integration key for Duo.
	IntegrationKey pulumi.StringInput
	// The mount to tie this method to for use in automatic mappings. The mapping will use the Name field of Aliases associated
	// with this mount as the username in the mapping.
	MountAccessor pulumi.StringInput
	// Name of the MFA method.
	Name pulumi.StringPtrInput
	// Push information for Duo.
	PushInfo pulumi.StringPtrInput
	// Secret key for Duo.
	SecretKey pulumi.StringInput
	// A format string for mapping Identity names to MFA method names. Values to substitute should be placed in `{{}}`.
	UsernameFormat pulumi.StringPtrInput
}

func (MfaDuoArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*mfaDuoArgs)(nil)).Elem()
}

type MfaDuoInput interface {
	pulumi.Input

	ToMfaDuoOutput() MfaDuoOutput
	ToMfaDuoOutputWithContext(ctx context.Context) MfaDuoOutput
}

func (MfaDuo) ElementType() reflect.Type {
	return reflect.TypeOf((*MfaDuo)(nil)).Elem()
}

func (i MfaDuo) ToMfaDuoOutput() MfaDuoOutput {
	return i.ToMfaDuoOutputWithContext(context.Background())
}

func (i MfaDuo) ToMfaDuoOutputWithContext(ctx context.Context) MfaDuoOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MfaDuoOutput)
}

type MfaDuoOutput struct {
	*pulumi.OutputState
}

func (MfaDuoOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MfaDuoOutput)(nil)).Elem()
}

func (o MfaDuoOutput) ToMfaDuoOutput() MfaDuoOutput {
	return o
}

func (o MfaDuoOutput) ToMfaDuoOutputWithContext(ctx context.Context) MfaDuoOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(MfaDuoOutput{})
}
