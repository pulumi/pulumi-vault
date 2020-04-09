// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package aws

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

type SecretBackendRole struct {
	pulumi.CustomResourceState

	// The path the AWS secret backend is mounted at,
	// with no leading or trailing `/`s.
	Backend pulumi.StringOutput `pulumi:"backend"`
	// Specifies the type of credential to be used when
	// retrieving credentials from the role. Must be one of `iamUser`, `assumedRole`, or
	// `federationToken`.
	CredentialType pulumi.StringOutput `pulumi:"credentialType"`
	// The default TTL in seconds for STS credentials.
	// When a TTL is not specified when STS credentials are requested,
	// and a default TTL is specified on the role,
	// then this default TTL will be used. Valid only when `credentialType` is one of
	// `assumedRole` or `federationToken`.
	DefaultStsTtl pulumi.IntOutput `pulumi:"defaultStsTtl"`
	// The max allowed TTL in seconds for STS credentials
	// (credentials TTL are capped to `maxStsTtl`). Valid only when `credentialType` is
	// one of `assumedRole` or `federationToken`.
	MaxStsTtl pulumi.IntOutput `pulumi:"maxStsTtl"`
	// The name to identify this role within the backend.
	// Must be unique within the backend.
	Name pulumi.StringOutput `pulumi:"name"`
	// Specifies a list of AWS managed policy ARNs. The
	// behavior depends on the credential type. With `iamUser`, the policies will be
	// attached to IAM users when they are requested. With `assumedRole` and
	// `federationToken`, the policy ARNs will act as a filter on what the credentials
	// can do, similar to `policyDocument`. When `credentialType` is `iamUser` or
	// `federationToken`, at least one of `policyDocument` or `policyArns` must
	// be specified.
	PolicyArns pulumi.StringArrayOutput `pulumi:"policyArns"`
	// The IAM policy document for the role. The
	// behavior depends on the credential type. With `iamUser`, the policy document
	// will be attached to the IAM user generated and augment the permissions the IAM
	// user has. With `assumedRole` and `federationToken`, the policy document will
	// act as a filter on what the credentials can do, similar to `policyArns`.
	PolicyDocument pulumi.StringPtrOutput `pulumi:"policyDocument"`
	// Specifies the ARNs of the AWS roles this Vault role
	// is allowed to assume. Required when `credentialType` is `assumedRole` and
	// prohibited otherwise.
	RoleArns pulumi.StringArrayOutput `pulumi:"roleArns"`
}

// NewSecretBackendRole registers a new resource with the given unique name, arguments, and options.
func NewSecretBackendRole(ctx *pulumi.Context,
	name string, args *SecretBackendRoleArgs, opts ...pulumi.ResourceOption) (*SecretBackendRole, error) {
	if args == nil || args.Backend == nil {
		return nil, errors.New("missing required argument 'Backend'")
	}
	if args == nil || args.CredentialType == nil {
		return nil, errors.New("missing required argument 'CredentialType'")
	}
	if args == nil {
		args = &SecretBackendRoleArgs{}
	}
	var resource SecretBackendRole
	err := ctx.RegisterResource("vault:aws/secretBackendRole:SecretBackendRole", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSecretBackendRole gets an existing SecretBackendRole resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSecretBackendRole(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SecretBackendRoleState, opts ...pulumi.ResourceOption) (*SecretBackendRole, error) {
	var resource SecretBackendRole
	err := ctx.ReadResource("vault:aws/secretBackendRole:SecretBackendRole", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SecretBackendRole resources.
type secretBackendRoleState struct {
	// The path the AWS secret backend is mounted at,
	// with no leading or trailing `/`s.
	Backend *string `pulumi:"backend"`
	// Specifies the type of credential to be used when
	// retrieving credentials from the role. Must be one of `iamUser`, `assumedRole`, or
	// `federationToken`.
	CredentialType *string `pulumi:"credentialType"`
	// The default TTL in seconds for STS credentials.
	// When a TTL is not specified when STS credentials are requested,
	// and a default TTL is specified on the role,
	// then this default TTL will be used. Valid only when `credentialType` is one of
	// `assumedRole` or `federationToken`.
	DefaultStsTtl *int `pulumi:"defaultStsTtl"`
	// The max allowed TTL in seconds for STS credentials
	// (credentials TTL are capped to `maxStsTtl`). Valid only when `credentialType` is
	// one of `assumedRole` or `federationToken`.
	MaxStsTtl *int `pulumi:"maxStsTtl"`
	// The name to identify this role within the backend.
	// Must be unique within the backend.
	Name *string `pulumi:"name"`
	// Specifies a list of AWS managed policy ARNs. The
	// behavior depends on the credential type. With `iamUser`, the policies will be
	// attached to IAM users when they are requested. With `assumedRole` and
	// `federationToken`, the policy ARNs will act as a filter on what the credentials
	// can do, similar to `policyDocument`. When `credentialType` is `iamUser` or
	// `federationToken`, at least one of `policyDocument` or `policyArns` must
	// be specified.
	PolicyArns []string `pulumi:"policyArns"`
	// The IAM policy document for the role. The
	// behavior depends on the credential type. With `iamUser`, the policy document
	// will be attached to the IAM user generated and augment the permissions the IAM
	// user has. With `assumedRole` and `federationToken`, the policy document will
	// act as a filter on what the credentials can do, similar to `policyArns`.
	PolicyDocument *string `pulumi:"policyDocument"`
	// Specifies the ARNs of the AWS roles this Vault role
	// is allowed to assume. Required when `credentialType` is `assumedRole` and
	// prohibited otherwise.
	RoleArns []string `pulumi:"roleArns"`
}

type SecretBackendRoleState struct {
	// The path the AWS secret backend is mounted at,
	// with no leading or trailing `/`s.
	Backend pulumi.StringPtrInput
	// Specifies the type of credential to be used when
	// retrieving credentials from the role. Must be one of `iamUser`, `assumedRole`, or
	// `federationToken`.
	CredentialType pulumi.StringPtrInput
	// The default TTL in seconds for STS credentials.
	// When a TTL is not specified when STS credentials are requested,
	// and a default TTL is specified on the role,
	// then this default TTL will be used. Valid only when `credentialType` is one of
	// `assumedRole` or `federationToken`.
	DefaultStsTtl pulumi.IntPtrInput
	// The max allowed TTL in seconds for STS credentials
	// (credentials TTL are capped to `maxStsTtl`). Valid only when `credentialType` is
	// one of `assumedRole` or `federationToken`.
	MaxStsTtl pulumi.IntPtrInput
	// The name to identify this role within the backend.
	// Must be unique within the backend.
	Name pulumi.StringPtrInput
	// Specifies a list of AWS managed policy ARNs. The
	// behavior depends on the credential type. With `iamUser`, the policies will be
	// attached to IAM users when they are requested. With `assumedRole` and
	// `federationToken`, the policy ARNs will act as a filter on what the credentials
	// can do, similar to `policyDocument`. When `credentialType` is `iamUser` or
	// `federationToken`, at least one of `policyDocument` or `policyArns` must
	// be specified.
	PolicyArns pulumi.StringArrayInput
	// The IAM policy document for the role. The
	// behavior depends on the credential type. With `iamUser`, the policy document
	// will be attached to the IAM user generated and augment the permissions the IAM
	// user has. With `assumedRole` and `federationToken`, the policy document will
	// act as a filter on what the credentials can do, similar to `policyArns`.
	PolicyDocument pulumi.StringPtrInput
	// Specifies the ARNs of the AWS roles this Vault role
	// is allowed to assume. Required when `credentialType` is `assumedRole` and
	// prohibited otherwise.
	RoleArns pulumi.StringArrayInput
}

func (SecretBackendRoleState) ElementType() reflect.Type {
	return reflect.TypeOf((*secretBackendRoleState)(nil)).Elem()
}

type secretBackendRoleArgs struct {
	// The path the AWS secret backend is mounted at,
	// with no leading or trailing `/`s.
	Backend string `pulumi:"backend"`
	// Specifies the type of credential to be used when
	// retrieving credentials from the role. Must be one of `iamUser`, `assumedRole`, or
	// `federationToken`.
	CredentialType string `pulumi:"credentialType"`
	// The default TTL in seconds for STS credentials.
	// When a TTL is not specified when STS credentials are requested,
	// and a default TTL is specified on the role,
	// then this default TTL will be used. Valid only when `credentialType` is one of
	// `assumedRole` or `federationToken`.
	DefaultStsTtl *int `pulumi:"defaultStsTtl"`
	// The max allowed TTL in seconds for STS credentials
	// (credentials TTL are capped to `maxStsTtl`). Valid only when `credentialType` is
	// one of `assumedRole` or `federationToken`.
	MaxStsTtl *int `pulumi:"maxStsTtl"`
	// The name to identify this role within the backend.
	// Must be unique within the backend.
	Name *string `pulumi:"name"`
	// Specifies a list of AWS managed policy ARNs. The
	// behavior depends on the credential type. With `iamUser`, the policies will be
	// attached to IAM users when they are requested. With `assumedRole` and
	// `federationToken`, the policy ARNs will act as a filter on what the credentials
	// can do, similar to `policyDocument`. When `credentialType` is `iamUser` or
	// `federationToken`, at least one of `policyDocument` or `policyArns` must
	// be specified.
	PolicyArns []string `pulumi:"policyArns"`
	// The IAM policy document for the role. The
	// behavior depends on the credential type. With `iamUser`, the policy document
	// will be attached to the IAM user generated and augment the permissions the IAM
	// user has. With `assumedRole` and `federationToken`, the policy document will
	// act as a filter on what the credentials can do, similar to `policyArns`.
	PolicyDocument *string `pulumi:"policyDocument"`
	// Specifies the ARNs of the AWS roles this Vault role
	// is allowed to assume. Required when `credentialType` is `assumedRole` and
	// prohibited otherwise.
	RoleArns []string `pulumi:"roleArns"`
}

// The set of arguments for constructing a SecretBackendRole resource.
type SecretBackendRoleArgs struct {
	// The path the AWS secret backend is mounted at,
	// with no leading or trailing `/`s.
	Backend pulumi.StringInput
	// Specifies the type of credential to be used when
	// retrieving credentials from the role. Must be one of `iamUser`, `assumedRole`, or
	// `federationToken`.
	CredentialType pulumi.StringInput
	// The default TTL in seconds for STS credentials.
	// When a TTL is not specified when STS credentials are requested,
	// and a default TTL is specified on the role,
	// then this default TTL will be used. Valid only when `credentialType` is one of
	// `assumedRole` or `federationToken`.
	DefaultStsTtl pulumi.IntPtrInput
	// The max allowed TTL in seconds for STS credentials
	// (credentials TTL are capped to `maxStsTtl`). Valid only when `credentialType` is
	// one of `assumedRole` or `federationToken`.
	MaxStsTtl pulumi.IntPtrInput
	// The name to identify this role within the backend.
	// Must be unique within the backend.
	Name pulumi.StringPtrInput
	// Specifies a list of AWS managed policy ARNs. The
	// behavior depends on the credential type. With `iamUser`, the policies will be
	// attached to IAM users when they are requested. With `assumedRole` and
	// `federationToken`, the policy ARNs will act as a filter on what the credentials
	// can do, similar to `policyDocument`. When `credentialType` is `iamUser` or
	// `federationToken`, at least one of `policyDocument` or `policyArns` must
	// be specified.
	PolicyArns pulumi.StringArrayInput
	// The IAM policy document for the role. The
	// behavior depends on the credential type. With `iamUser`, the policy document
	// will be attached to the IAM user generated and augment the permissions the IAM
	// user has. With `assumedRole` and `federationToken`, the policy document will
	// act as a filter on what the credentials can do, similar to `policyArns`.
	PolicyDocument pulumi.StringPtrInput
	// Specifies the ARNs of the AWS roles this Vault role
	// is allowed to assume. Required when `credentialType` is `assumedRole` and
	// prohibited otherwise.
	RoleArns pulumi.StringArrayInput
}

func (SecretBackendRoleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*secretBackendRoleArgs)(nil)).Elem()
}
