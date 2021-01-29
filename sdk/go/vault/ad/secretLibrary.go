// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ad

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

type SecretLibrary struct {
	pulumi.CustomResourceState

	// The mount path for the AD backend.
	Backend pulumi.StringOutput `pulumi:"backend"`
	// Disable enforcing that service accounts must be checked in by the entity or client token that checked them out.
	DisableCheckInEnforcement pulumi.BoolPtrOutput `pulumi:"disableCheckInEnforcement"`
	// The maximum amount of time, in seconds, a check-out last with renewal before Vault automatically checks it back in.
	MaxTtl pulumi.IntOutput `pulumi:"maxTtl"`
	// The name of the set of service accounts.
	Name pulumi.StringOutput `pulumi:"name"`
	// The names of all the service accounts that can be checked out from this set. These service accounts must already exist
	// in Active Directory.
	ServiceAccountNames pulumi.StringArrayOutput `pulumi:"serviceAccountNames"`
	// The amount of time, in seconds, a single check-out lasts before Vault automatically checks it back in.
	Ttl pulumi.IntOutput `pulumi:"ttl"`
}

// NewSecretLibrary registers a new resource with the given unique name, arguments, and options.
func NewSecretLibrary(ctx *pulumi.Context,
	name string, args *SecretLibraryArgs, opts ...pulumi.ResourceOption) (*SecretLibrary, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Backend == nil {
		return nil, errors.New("invalid value for required argument 'Backend'")
	}
	if args.ServiceAccountNames == nil {
		return nil, errors.New("invalid value for required argument 'ServiceAccountNames'")
	}
	var resource SecretLibrary
	err := ctx.RegisterResource("vault:ad/secretLibrary:SecretLibrary", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSecretLibrary gets an existing SecretLibrary resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSecretLibrary(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SecretLibraryState, opts ...pulumi.ResourceOption) (*SecretLibrary, error) {
	var resource SecretLibrary
	err := ctx.ReadResource("vault:ad/secretLibrary:SecretLibrary", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SecretLibrary resources.
type secretLibraryState struct {
	// The mount path for the AD backend.
	Backend *string `pulumi:"backend"`
	// Disable enforcing that service accounts must be checked in by the entity or client token that checked them out.
	DisableCheckInEnforcement *bool `pulumi:"disableCheckInEnforcement"`
	// The maximum amount of time, in seconds, a check-out last with renewal before Vault automatically checks it back in.
	MaxTtl *int `pulumi:"maxTtl"`
	// The name of the set of service accounts.
	Name *string `pulumi:"name"`
	// The names of all the service accounts that can be checked out from this set. These service accounts must already exist
	// in Active Directory.
	ServiceAccountNames []string `pulumi:"serviceAccountNames"`
	// The amount of time, in seconds, a single check-out lasts before Vault automatically checks it back in.
	Ttl *int `pulumi:"ttl"`
}

type SecretLibraryState struct {
	// The mount path for the AD backend.
	Backend pulumi.StringPtrInput
	// Disable enforcing that service accounts must be checked in by the entity or client token that checked them out.
	DisableCheckInEnforcement pulumi.BoolPtrInput
	// The maximum amount of time, in seconds, a check-out last with renewal before Vault automatically checks it back in.
	MaxTtl pulumi.IntPtrInput
	// The name of the set of service accounts.
	Name pulumi.StringPtrInput
	// The names of all the service accounts that can be checked out from this set. These service accounts must already exist
	// in Active Directory.
	ServiceAccountNames pulumi.StringArrayInput
	// The amount of time, in seconds, a single check-out lasts before Vault automatically checks it back in.
	Ttl pulumi.IntPtrInput
}

func (SecretLibraryState) ElementType() reflect.Type {
	return reflect.TypeOf((*secretLibraryState)(nil)).Elem()
}

type secretLibraryArgs struct {
	// The mount path for the AD backend.
	Backend string `pulumi:"backend"`
	// Disable enforcing that service accounts must be checked in by the entity or client token that checked them out.
	DisableCheckInEnforcement *bool `pulumi:"disableCheckInEnforcement"`
	// The maximum amount of time, in seconds, a check-out last with renewal before Vault automatically checks it back in.
	MaxTtl *int `pulumi:"maxTtl"`
	// The name of the set of service accounts.
	Name *string `pulumi:"name"`
	// The names of all the service accounts that can be checked out from this set. These service accounts must already exist
	// in Active Directory.
	ServiceAccountNames []string `pulumi:"serviceAccountNames"`
	// The amount of time, in seconds, a single check-out lasts before Vault automatically checks it back in.
	Ttl *int `pulumi:"ttl"`
}

// The set of arguments for constructing a SecretLibrary resource.
type SecretLibraryArgs struct {
	// The mount path for the AD backend.
	Backend pulumi.StringInput
	// Disable enforcing that service accounts must be checked in by the entity or client token that checked them out.
	DisableCheckInEnforcement pulumi.BoolPtrInput
	// The maximum amount of time, in seconds, a check-out last with renewal before Vault automatically checks it back in.
	MaxTtl pulumi.IntPtrInput
	// The name of the set of service accounts.
	Name pulumi.StringPtrInput
	// The names of all the service accounts that can be checked out from this set. These service accounts must already exist
	// in Active Directory.
	ServiceAccountNames pulumi.StringArrayInput
	// The amount of time, in seconds, a single check-out lasts before Vault automatically checks it back in.
	Ttl pulumi.IntPtrInput
}

func (SecretLibraryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*secretLibraryArgs)(nil)).Elem()
}

type SecretLibraryInput interface {
	pulumi.Input

	ToSecretLibraryOutput() SecretLibraryOutput
	ToSecretLibraryOutputWithContext(ctx context.Context) SecretLibraryOutput
}

func (*SecretLibrary) ElementType() reflect.Type {
	return reflect.TypeOf((*SecretLibrary)(nil))
}

func (i *SecretLibrary) ToSecretLibraryOutput() SecretLibraryOutput {
	return i.ToSecretLibraryOutputWithContext(context.Background())
}

func (i *SecretLibrary) ToSecretLibraryOutputWithContext(ctx context.Context) SecretLibraryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecretLibraryOutput)
}

type SecretLibraryOutput struct {
	*pulumi.OutputState
}

func (SecretLibraryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SecretLibrary)(nil))
}

func (o SecretLibraryOutput) ToSecretLibraryOutput() SecretLibraryOutput {
	return o
}

func (o SecretLibraryOutput) ToSecretLibraryOutputWithContext(ctx context.Context) SecretLibraryOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(SecretLibraryOutput{})
}
