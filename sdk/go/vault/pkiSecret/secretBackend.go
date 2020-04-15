// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package pkiSecret

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Creates an PKI Secret Backend for Vault. PKI secret backends can then issue certificates, once a role has been added to
// the backend.
type SecretBackend struct {
	pulumi.CustomResourceState

	// The default TTL for credentials issued by this backend.
	DefaultLeaseTtlSeconds pulumi.IntOutput `pulumi:"defaultLeaseTtlSeconds"`
	// A human-friendly description for this backend.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The maximum TTL that can be requested for credentials issued by this backend.
	MaxLeaseTtlSeconds pulumi.IntOutput `pulumi:"maxLeaseTtlSeconds"`
	// The unique path this backend should be mounted at. Must not begin or end with a `/`.
	Path pulumi.StringOutput `pulumi:"path"`
}

// NewSecretBackend registers a new resource with the given unique name, arguments, and options.
func NewSecretBackend(ctx *pulumi.Context,
	name string, args *SecretBackendArgs, opts ...pulumi.ResourceOption) (*SecretBackend, error) {
	if args == nil || args.Path == nil {
		return nil, errors.New("missing required argument 'Path'")
	}
	if args == nil {
		args = &SecretBackendArgs{}
	}
	var resource SecretBackend
	err := ctx.RegisterResource("vault:pkiSecret/secretBackend:SecretBackend", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSecretBackend gets an existing SecretBackend resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSecretBackend(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SecretBackendState, opts ...pulumi.ResourceOption) (*SecretBackend, error) {
	var resource SecretBackend
	err := ctx.ReadResource("vault:pkiSecret/secretBackend:SecretBackend", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SecretBackend resources.
type secretBackendState struct {
	// The default TTL for credentials issued by this backend.
	DefaultLeaseTtlSeconds *int `pulumi:"defaultLeaseTtlSeconds"`
	// A human-friendly description for this backend.
	Description *string `pulumi:"description"`
	// The maximum TTL that can be requested for credentials issued by this backend.
	MaxLeaseTtlSeconds *int `pulumi:"maxLeaseTtlSeconds"`
	// The unique path this backend should be mounted at. Must not begin or end with a `/`.
	Path *string `pulumi:"path"`
}

type SecretBackendState struct {
	// The default TTL for credentials issued by this backend.
	DefaultLeaseTtlSeconds pulumi.IntPtrInput
	// A human-friendly description for this backend.
	Description pulumi.StringPtrInput
	// The maximum TTL that can be requested for credentials issued by this backend.
	MaxLeaseTtlSeconds pulumi.IntPtrInput
	// The unique path this backend should be mounted at. Must not begin or end with a `/`.
	Path pulumi.StringPtrInput
}

func (SecretBackendState) ElementType() reflect.Type {
	return reflect.TypeOf((*secretBackendState)(nil)).Elem()
}

type secretBackendArgs struct {
	// The default TTL for credentials issued by this backend.
	DefaultLeaseTtlSeconds *int `pulumi:"defaultLeaseTtlSeconds"`
	// A human-friendly description for this backend.
	Description *string `pulumi:"description"`
	// The maximum TTL that can be requested for credentials issued by this backend.
	MaxLeaseTtlSeconds *int `pulumi:"maxLeaseTtlSeconds"`
	// The unique path this backend should be mounted at. Must not begin or end with a `/`.
	Path string `pulumi:"path"`
}

// The set of arguments for constructing a SecretBackend resource.
type SecretBackendArgs struct {
	// The default TTL for credentials issued by this backend.
	DefaultLeaseTtlSeconds pulumi.IntPtrInput
	// A human-friendly description for this backend.
	Description pulumi.StringPtrInput
	// The maximum TTL that can be requested for credentials issued by this backend.
	MaxLeaseTtlSeconds pulumi.IntPtrInput
	// The unique path this backend should be mounted at. Must not begin or end with a `/`.
	Path pulumi.StringInput
}

func (SecretBackendArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*secretBackendArgs)(nil)).Elem()
}
