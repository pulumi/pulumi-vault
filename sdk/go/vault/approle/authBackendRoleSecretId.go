// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package approle

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-vault/sdk/v6/go/vault/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages an AppRole auth backend SecretID in a Vault server. See the [Vault
// documentation](https://www.vaultproject.io/docs/auth/approle) for more
// information.
//
// ## Example Usage
//
// <!--Start PulumiCodeChooser -->
// ```go
// package main
//
// import (
//
//	"encoding/json"
//
//	"github.com/pulumi/pulumi-vault/sdk/v6/go/vault"
//	"github.com/pulumi/pulumi-vault/sdk/v6/go/vault/appRole"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			approle, err := vault.NewAuthBackend(ctx, "approle", &vault.AuthBackendArgs{
//				Type: pulumi.String("approle"),
//			})
//			if err != nil {
//				return err
//			}
//			example, err := appRole.NewAuthBackendRole(ctx, "example", &appRole.AuthBackendRoleArgs{
//				Backend:  approle.Path,
//				RoleName: pulumi.String("test-role"),
//				TokenPolicies: pulumi.StringArray{
//					pulumi.String("default"),
//					pulumi.String("dev"),
//					pulumi.String("prod"),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			tmpJSON0, err := json.Marshal(map[string]interface{}{
//				"hello": "world",
//			})
//			if err != nil {
//				return err
//			}
//			json0 := string(tmpJSON0)
//			_, err = appRole.NewAuthBackendRoleSecretId(ctx, "id", &appRole.AuthBackendRoleSecretIdArgs{
//				Backend:  approle.Path,
//				RoleName: example.RoleName,
//				Metadata: pulumi.String(json0),
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
type AuthBackendRoleSecretId struct {
	pulumi.CustomResourceState

	// The unique ID for this SecretID that can be safely logged.
	Accessor pulumi.StringOutput `pulumi:"accessor"`
	// Unique name of the auth backend to configure.
	Backend pulumi.StringPtrOutput `pulumi:"backend"`
	// If set, specifies blocks of IP addresses which can
	// perform the login operation using this SecretID.
	CidrLists pulumi.StringArrayOutput `pulumi:"cidrLists"`
	// A JSON-encoded string containing metadata in
	// key-value pairs to be set on tokens issued with this SecretID.
	Metadata pulumi.StringPtrOutput `pulumi:"metadata"`
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
	// *Available only for Vault Enterprise*.
	Namespace pulumi.StringPtrOutput `pulumi:"namespace"`
	// The name of the role to create the SecretID for.
	RoleName pulumi.StringOutput `pulumi:"roleName"`
	// The SecretID to be created. If set, uses "Push"
	// mode.  Defaults to Vault auto-generating SecretIDs.
	SecretId pulumi.StringOutput `pulumi:"secretId"`
	// Set to `true` to use the wrapped secret-id accessor as the resource ID.
	// If `false` (default value), a fresh secret ID will be regenerated whenever the wrapping token is expired or
	// invalidated through unwrapping.
	WithWrappedAccessor pulumi.BoolPtrOutput `pulumi:"withWrappedAccessor"`
	// The unique ID for the response-wrapped SecretID that can
	// be safely logged.
	WrappingAccessor pulumi.StringOutput `pulumi:"wrappingAccessor"`
	// The token used to retrieve a response-wrapped SecretID.
	WrappingToken pulumi.StringOutput `pulumi:"wrappingToken"`
	// If set, the SecretID response will be
	// [response-wrapped](https://www.vaultproject.io/docs/concepts/response-wrapping)
	// and available for the duration specified. Only a single unwrapping of the
	// token is allowed.
	WrappingTtl pulumi.StringPtrOutput `pulumi:"wrappingTtl"`
}

// NewAuthBackendRoleSecretId registers a new resource with the given unique name, arguments, and options.
func NewAuthBackendRoleSecretId(ctx *pulumi.Context,
	name string, args *AuthBackendRoleSecretIdArgs, opts ...pulumi.ResourceOption) (*AuthBackendRoleSecretId, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.RoleName == nil {
		return nil, errors.New("invalid value for required argument 'RoleName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("vault:appRole/authBackendRoleSecretID:AuthBackendRoleSecretID"),
		},
	})
	opts = append(opts, aliases)
	if args.SecretId != nil {
		args.SecretId = pulumi.ToSecret(args.SecretId).(pulumi.StringPtrInput)
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"secretId",
		"wrappingToken",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource AuthBackendRoleSecretId
	err := ctx.RegisterResource("vault:appRole/authBackendRoleSecretId:AuthBackendRoleSecretId", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAuthBackendRoleSecretId gets an existing AuthBackendRoleSecretId resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAuthBackendRoleSecretId(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AuthBackendRoleSecretIdState, opts ...pulumi.ResourceOption) (*AuthBackendRoleSecretId, error) {
	var resource AuthBackendRoleSecretId
	err := ctx.ReadResource("vault:appRole/authBackendRoleSecretId:AuthBackendRoleSecretId", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AuthBackendRoleSecretId resources.
type authBackendRoleSecretIdState struct {
	// The unique ID for this SecretID that can be safely logged.
	Accessor *string `pulumi:"accessor"`
	// Unique name of the auth backend to configure.
	Backend *string `pulumi:"backend"`
	// If set, specifies blocks of IP addresses which can
	// perform the login operation using this SecretID.
	CidrLists []string `pulumi:"cidrLists"`
	// A JSON-encoded string containing metadata in
	// key-value pairs to be set on tokens issued with this SecretID.
	Metadata *string `pulumi:"metadata"`
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
	// *Available only for Vault Enterprise*.
	Namespace *string `pulumi:"namespace"`
	// The name of the role to create the SecretID for.
	RoleName *string `pulumi:"roleName"`
	// The SecretID to be created. If set, uses "Push"
	// mode.  Defaults to Vault auto-generating SecretIDs.
	SecretId *string `pulumi:"secretId"`
	// Set to `true` to use the wrapped secret-id accessor as the resource ID.
	// If `false` (default value), a fresh secret ID will be regenerated whenever the wrapping token is expired or
	// invalidated through unwrapping.
	WithWrappedAccessor *bool `pulumi:"withWrappedAccessor"`
	// The unique ID for the response-wrapped SecretID that can
	// be safely logged.
	WrappingAccessor *string `pulumi:"wrappingAccessor"`
	// The token used to retrieve a response-wrapped SecretID.
	WrappingToken *string `pulumi:"wrappingToken"`
	// If set, the SecretID response will be
	// [response-wrapped](https://www.vaultproject.io/docs/concepts/response-wrapping)
	// and available for the duration specified. Only a single unwrapping of the
	// token is allowed.
	WrappingTtl *string `pulumi:"wrappingTtl"`
}

type AuthBackendRoleSecretIdState struct {
	// The unique ID for this SecretID that can be safely logged.
	Accessor pulumi.StringPtrInput
	// Unique name of the auth backend to configure.
	Backend pulumi.StringPtrInput
	// If set, specifies blocks of IP addresses which can
	// perform the login operation using this SecretID.
	CidrLists pulumi.StringArrayInput
	// A JSON-encoded string containing metadata in
	// key-value pairs to be set on tokens issued with this SecretID.
	Metadata pulumi.StringPtrInput
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
	// *Available only for Vault Enterprise*.
	Namespace pulumi.StringPtrInput
	// The name of the role to create the SecretID for.
	RoleName pulumi.StringPtrInput
	// The SecretID to be created. If set, uses "Push"
	// mode.  Defaults to Vault auto-generating SecretIDs.
	SecretId pulumi.StringPtrInput
	// Set to `true` to use the wrapped secret-id accessor as the resource ID.
	// If `false` (default value), a fresh secret ID will be regenerated whenever the wrapping token is expired or
	// invalidated through unwrapping.
	WithWrappedAccessor pulumi.BoolPtrInput
	// The unique ID for the response-wrapped SecretID that can
	// be safely logged.
	WrappingAccessor pulumi.StringPtrInput
	// The token used to retrieve a response-wrapped SecretID.
	WrappingToken pulumi.StringPtrInput
	// If set, the SecretID response will be
	// [response-wrapped](https://www.vaultproject.io/docs/concepts/response-wrapping)
	// and available for the duration specified. Only a single unwrapping of the
	// token is allowed.
	WrappingTtl pulumi.StringPtrInput
}

func (AuthBackendRoleSecretIdState) ElementType() reflect.Type {
	return reflect.TypeOf((*authBackendRoleSecretIdState)(nil)).Elem()
}

type authBackendRoleSecretIdArgs struct {
	// Unique name of the auth backend to configure.
	Backend *string `pulumi:"backend"`
	// If set, specifies blocks of IP addresses which can
	// perform the login operation using this SecretID.
	CidrLists []string `pulumi:"cidrLists"`
	// A JSON-encoded string containing metadata in
	// key-value pairs to be set on tokens issued with this SecretID.
	Metadata *string `pulumi:"metadata"`
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
	// *Available only for Vault Enterprise*.
	Namespace *string `pulumi:"namespace"`
	// The name of the role to create the SecretID for.
	RoleName string `pulumi:"roleName"`
	// The SecretID to be created. If set, uses "Push"
	// mode.  Defaults to Vault auto-generating SecretIDs.
	SecretId *string `pulumi:"secretId"`
	// Set to `true` to use the wrapped secret-id accessor as the resource ID.
	// If `false` (default value), a fresh secret ID will be regenerated whenever the wrapping token is expired or
	// invalidated through unwrapping.
	WithWrappedAccessor *bool `pulumi:"withWrappedAccessor"`
	// If set, the SecretID response will be
	// [response-wrapped](https://www.vaultproject.io/docs/concepts/response-wrapping)
	// and available for the duration specified. Only a single unwrapping of the
	// token is allowed.
	WrappingTtl *string `pulumi:"wrappingTtl"`
}

// The set of arguments for constructing a AuthBackendRoleSecretId resource.
type AuthBackendRoleSecretIdArgs struct {
	// Unique name of the auth backend to configure.
	Backend pulumi.StringPtrInput
	// If set, specifies blocks of IP addresses which can
	// perform the login operation using this SecretID.
	CidrLists pulumi.StringArrayInput
	// A JSON-encoded string containing metadata in
	// key-value pairs to be set on tokens issued with this SecretID.
	Metadata pulumi.StringPtrInput
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
	// *Available only for Vault Enterprise*.
	Namespace pulumi.StringPtrInput
	// The name of the role to create the SecretID for.
	RoleName pulumi.StringInput
	// The SecretID to be created. If set, uses "Push"
	// mode.  Defaults to Vault auto-generating SecretIDs.
	SecretId pulumi.StringPtrInput
	// Set to `true` to use the wrapped secret-id accessor as the resource ID.
	// If `false` (default value), a fresh secret ID will be regenerated whenever the wrapping token is expired or
	// invalidated through unwrapping.
	WithWrappedAccessor pulumi.BoolPtrInput
	// If set, the SecretID response will be
	// [response-wrapped](https://www.vaultproject.io/docs/concepts/response-wrapping)
	// and available for the duration specified. Only a single unwrapping of the
	// token is allowed.
	WrappingTtl pulumi.StringPtrInput
}

func (AuthBackendRoleSecretIdArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*authBackendRoleSecretIdArgs)(nil)).Elem()
}

type AuthBackendRoleSecretIdInput interface {
	pulumi.Input

	ToAuthBackendRoleSecretIdOutput() AuthBackendRoleSecretIdOutput
	ToAuthBackendRoleSecretIdOutputWithContext(ctx context.Context) AuthBackendRoleSecretIdOutput
}

func (*AuthBackendRoleSecretId) ElementType() reflect.Type {
	return reflect.TypeOf((**AuthBackendRoleSecretId)(nil)).Elem()
}

func (i *AuthBackendRoleSecretId) ToAuthBackendRoleSecretIdOutput() AuthBackendRoleSecretIdOutput {
	return i.ToAuthBackendRoleSecretIdOutputWithContext(context.Background())
}

func (i *AuthBackendRoleSecretId) ToAuthBackendRoleSecretIdOutputWithContext(ctx context.Context) AuthBackendRoleSecretIdOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AuthBackendRoleSecretIdOutput)
}

// AuthBackendRoleSecretIdArrayInput is an input type that accepts AuthBackendRoleSecretIdArray and AuthBackendRoleSecretIdArrayOutput values.
// You can construct a concrete instance of `AuthBackendRoleSecretIdArrayInput` via:
//
//	AuthBackendRoleSecretIdArray{ AuthBackendRoleSecretIdArgs{...} }
type AuthBackendRoleSecretIdArrayInput interface {
	pulumi.Input

	ToAuthBackendRoleSecretIdArrayOutput() AuthBackendRoleSecretIdArrayOutput
	ToAuthBackendRoleSecretIdArrayOutputWithContext(context.Context) AuthBackendRoleSecretIdArrayOutput
}

type AuthBackendRoleSecretIdArray []AuthBackendRoleSecretIdInput

func (AuthBackendRoleSecretIdArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AuthBackendRoleSecretId)(nil)).Elem()
}

func (i AuthBackendRoleSecretIdArray) ToAuthBackendRoleSecretIdArrayOutput() AuthBackendRoleSecretIdArrayOutput {
	return i.ToAuthBackendRoleSecretIdArrayOutputWithContext(context.Background())
}

func (i AuthBackendRoleSecretIdArray) ToAuthBackendRoleSecretIdArrayOutputWithContext(ctx context.Context) AuthBackendRoleSecretIdArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AuthBackendRoleSecretIdArrayOutput)
}

// AuthBackendRoleSecretIdMapInput is an input type that accepts AuthBackendRoleSecretIdMap and AuthBackendRoleSecretIdMapOutput values.
// You can construct a concrete instance of `AuthBackendRoleSecretIdMapInput` via:
//
//	AuthBackendRoleSecretIdMap{ "key": AuthBackendRoleSecretIdArgs{...} }
type AuthBackendRoleSecretIdMapInput interface {
	pulumi.Input

	ToAuthBackendRoleSecretIdMapOutput() AuthBackendRoleSecretIdMapOutput
	ToAuthBackendRoleSecretIdMapOutputWithContext(context.Context) AuthBackendRoleSecretIdMapOutput
}

type AuthBackendRoleSecretIdMap map[string]AuthBackendRoleSecretIdInput

func (AuthBackendRoleSecretIdMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AuthBackendRoleSecretId)(nil)).Elem()
}

func (i AuthBackendRoleSecretIdMap) ToAuthBackendRoleSecretIdMapOutput() AuthBackendRoleSecretIdMapOutput {
	return i.ToAuthBackendRoleSecretIdMapOutputWithContext(context.Background())
}

func (i AuthBackendRoleSecretIdMap) ToAuthBackendRoleSecretIdMapOutputWithContext(ctx context.Context) AuthBackendRoleSecretIdMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AuthBackendRoleSecretIdMapOutput)
}

type AuthBackendRoleSecretIdOutput struct{ *pulumi.OutputState }

func (AuthBackendRoleSecretIdOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AuthBackendRoleSecretId)(nil)).Elem()
}

func (o AuthBackendRoleSecretIdOutput) ToAuthBackendRoleSecretIdOutput() AuthBackendRoleSecretIdOutput {
	return o
}

func (o AuthBackendRoleSecretIdOutput) ToAuthBackendRoleSecretIdOutputWithContext(ctx context.Context) AuthBackendRoleSecretIdOutput {
	return o
}

// The unique ID for this SecretID that can be safely logged.
func (o AuthBackendRoleSecretIdOutput) Accessor() pulumi.StringOutput {
	return o.ApplyT(func(v *AuthBackendRoleSecretId) pulumi.StringOutput { return v.Accessor }).(pulumi.StringOutput)
}

// Unique name of the auth backend to configure.
func (o AuthBackendRoleSecretIdOutput) Backend() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AuthBackendRoleSecretId) pulumi.StringPtrOutput { return v.Backend }).(pulumi.StringPtrOutput)
}

// If set, specifies blocks of IP addresses which can
// perform the login operation using this SecretID.
func (o AuthBackendRoleSecretIdOutput) CidrLists() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *AuthBackendRoleSecretId) pulumi.StringArrayOutput { return v.CidrLists }).(pulumi.StringArrayOutput)
}

// A JSON-encoded string containing metadata in
// key-value pairs to be set on tokens issued with this SecretID.
func (o AuthBackendRoleSecretIdOutput) Metadata() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AuthBackendRoleSecretId) pulumi.StringPtrOutput { return v.Metadata }).(pulumi.StringPtrOutput)
}

// The namespace to provision the resource in.
// The value should not contain leading or trailing forward slashes.
// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
// *Available only for Vault Enterprise*.
func (o AuthBackendRoleSecretIdOutput) Namespace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AuthBackendRoleSecretId) pulumi.StringPtrOutput { return v.Namespace }).(pulumi.StringPtrOutput)
}

// The name of the role to create the SecretID for.
func (o AuthBackendRoleSecretIdOutput) RoleName() pulumi.StringOutput {
	return o.ApplyT(func(v *AuthBackendRoleSecretId) pulumi.StringOutput { return v.RoleName }).(pulumi.StringOutput)
}

// The SecretID to be created. If set, uses "Push"
// mode.  Defaults to Vault auto-generating SecretIDs.
func (o AuthBackendRoleSecretIdOutput) SecretId() pulumi.StringOutput {
	return o.ApplyT(func(v *AuthBackendRoleSecretId) pulumi.StringOutput { return v.SecretId }).(pulumi.StringOutput)
}

// Set to `true` to use the wrapped secret-id accessor as the resource ID.
// If `false` (default value), a fresh secret ID will be regenerated whenever the wrapping token is expired or
// invalidated through unwrapping.
func (o AuthBackendRoleSecretIdOutput) WithWrappedAccessor() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AuthBackendRoleSecretId) pulumi.BoolPtrOutput { return v.WithWrappedAccessor }).(pulumi.BoolPtrOutput)
}

// The unique ID for the response-wrapped SecretID that can
// be safely logged.
func (o AuthBackendRoleSecretIdOutput) WrappingAccessor() pulumi.StringOutput {
	return o.ApplyT(func(v *AuthBackendRoleSecretId) pulumi.StringOutput { return v.WrappingAccessor }).(pulumi.StringOutput)
}

// The token used to retrieve a response-wrapped SecretID.
func (o AuthBackendRoleSecretIdOutput) WrappingToken() pulumi.StringOutput {
	return o.ApplyT(func(v *AuthBackendRoleSecretId) pulumi.StringOutput { return v.WrappingToken }).(pulumi.StringOutput)
}

// If set, the SecretID response will be
// [response-wrapped](https://www.vaultproject.io/docs/concepts/response-wrapping)
// and available for the duration specified. Only a single unwrapping of the
// token is allowed.
func (o AuthBackendRoleSecretIdOutput) WrappingTtl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AuthBackendRoleSecretId) pulumi.StringPtrOutput { return v.WrappingTtl }).(pulumi.StringPtrOutput)
}

type AuthBackendRoleSecretIdArrayOutput struct{ *pulumi.OutputState }

func (AuthBackendRoleSecretIdArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AuthBackendRoleSecretId)(nil)).Elem()
}

func (o AuthBackendRoleSecretIdArrayOutput) ToAuthBackendRoleSecretIdArrayOutput() AuthBackendRoleSecretIdArrayOutput {
	return o
}

func (o AuthBackendRoleSecretIdArrayOutput) ToAuthBackendRoleSecretIdArrayOutputWithContext(ctx context.Context) AuthBackendRoleSecretIdArrayOutput {
	return o
}

func (o AuthBackendRoleSecretIdArrayOutput) Index(i pulumi.IntInput) AuthBackendRoleSecretIdOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *AuthBackendRoleSecretId {
		return vs[0].([]*AuthBackendRoleSecretId)[vs[1].(int)]
	}).(AuthBackendRoleSecretIdOutput)
}

type AuthBackendRoleSecretIdMapOutput struct{ *pulumi.OutputState }

func (AuthBackendRoleSecretIdMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AuthBackendRoleSecretId)(nil)).Elem()
}

func (o AuthBackendRoleSecretIdMapOutput) ToAuthBackendRoleSecretIdMapOutput() AuthBackendRoleSecretIdMapOutput {
	return o
}

func (o AuthBackendRoleSecretIdMapOutput) ToAuthBackendRoleSecretIdMapOutputWithContext(ctx context.Context) AuthBackendRoleSecretIdMapOutput {
	return o
}

func (o AuthBackendRoleSecretIdMapOutput) MapIndex(k pulumi.StringInput) AuthBackendRoleSecretIdOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *AuthBackendRoleSecretId {
		return vs[0].(map[string]*AuthBackendRoleSecretId)[vs[1].(string)]
	}).(AuthBackendRoleSecretIdOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AuthBackendRoleSecretIdInput)(nil)).Elem(), &AuthBackendRoleSecretId{})
	pulumi.RegisterInputType(reflect.TypeOf((*AuthBackendRoleSecretIdArrayInput)(nil)).Elem(), AuthBackendRoleSecretIdArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AuthBackendRoleSecretIdMapInput)(nil)).Elem(), AuthBackendRoleSecretIdMap{})
	pulumi.RegisterOutputType(AuthBackendRoleSecretIdOutput{})
	pulumi.RegisterOutputType(AuthBackendRoleSecretIdArrayOutput{})
	pulumi.RegisterOutputType(AuthBackendRoleSecretIdMapOutput{})
}
