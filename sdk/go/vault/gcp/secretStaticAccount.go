// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package gcp

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-vault/sdk/v7/go/vault/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a Static Account in the [GCP Secrets Engine](https://www.vaultproject.io/docs/secrets/gcp/index.html) for Vault.
//
// Each [static account](https://www.vaultproject.io/docs/secrets/gcp/index.html#static-accounts) is tied to a separately managed
// Service Account, and can have one or more [bindings](https://www.vaultproject.io/docs/secrets/gcp/index.html#bindings) associated with it.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"fmt"
//
//	"github.com/pulumi/pulumi-google/sdk/go/google"
//	"github.com/pulumi/pulumi-std/sdk/go/std"
//	"github.com/pulumi/pulumi-vault/sdk/v7/go/vault/gcp"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			this, err := google.NewServiceAccount(ctx, "this", &google.ServiceAccountArgs{
//				AccountId: "my-awesome-account",
//			})
//			if err != nil {
//				return err
//			}
//			invokeFile, err := std.File(ctx, &std.FileArgs{
//				Input: "credentials.json",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			gcp, err := gcp.NewSecretBackend(ctx, "gcp", &gcp.SecretBackendArgs{
//				Path:        pulumi.String("gcp"),
//				Credentials: pulumi.String(invokeFile.Result),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = gcp.NewSecretStaticAccount(ctx, "static_account", &gcp.SecretStaticAccountArgs{
//				Backend:       gcp.Path,
//				StaticAccount: pulumi.String("project_viewer"),
//				SecretType:    pulumi.String("access_token"),
//				TokenScopes: pulumi.StringArray{
//					pulumi.String("https://www.googleapis.com/auth/cloud-platform"),
//				},
//				ServiceAccountEmail: this.Email,
//				Bindings: gcp.SecretStaticAccountBindingArray{
//					&gcp.SecretStaticAccountBindingArgs{
//						Resource: pulumi.Sprintf("//cloudresourcemanager.googleapis.com/projects/%v", this.Project),
//						Roles: pulumi.StringArray{
//							pulumi.String("roles/viewer"),
//						},
//					},
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
//
// ## Import
//
// A static account can be imported using its Vault Path. For example, referencing the example above,
//
// ```sh
// $ pulumi import vault:gcp/secretStaticAccount:SecretStaticAccount static_account gcp/static-account/project_viewer
// ```
type SecretStaticAccount struct {
	pulumi.CustomResourceState

	// Path where the GCP Secrets Engine is mounted
	Backend pulumi.StringOutput `pulumi:"backend"`
	// Bindings to create for this static account. This can be specified multiple times for multiple bindings. Structure is documented below.
	Bindings SecretStaticAccountBindingArrayOutput `pulumi:"bindings"`
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
	// *Available only for Vault Enterprise*.
	Namespace pulumi.StringPtrOutput `pulumi:"namespace"`
	// Type of secret generated for this static account. Accepted values: `accessToken`, `serviceAccountKey`. Defaults to `accessToken`.
	SecretType pulumi.StringOutput `pulumi:"secretType"`
	// Email of the GCP service account to manage.
	ServiceAccountEmail pulumi.StringOutput `pulumi:"serviceAccountEmail"`
	// Project the service account belongs to.
	ServiceAccountProject pulumi.StringOutput `pulumi:"serviceAccountProject"`
	// Name of the Static Account to create
	StaticAccount pulumi.StringOutput `pulumi:"staticAccount"`
	// List of OAuth scopes to assign to `accessToken` secrets generated under this static account (`accessToken` static accounts only).
	TokenScopes pulumi.StringArrayOutput `pulumi:"tokenScopes"`
}

// NewSecretStaticAccount registers a new resource with the given unique name, arguments, and options.
func NewSecretStaticAccount(ctx *pulumi.Context,
	name string, args *SecretStaticAccountArgs, opts ...pulumi.ResourceOption) (*SecretStaticAccount, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Backend == nil {
		return nil, errors.New("invalid value for required argument 'Backend'")
	}
	if args.ServiceAccountEmail == nil {
		return nil, errors.New("invalid value for required argument 'ServiceAccountEmail'")
	}
	if args.StaticAccount == nil {
		return nil, errors.New("invalid value for required argument 'StaticAccount'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource SecretStaticAccount
	err := ctx.RegisterResource("vault:gcp/secretStaticAccount:SecretStaticAccount", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSecretStaticAccount gets an existing SecretStaticAccount resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSecretStaticAccount(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SecretStaticAccountState, opts ...pulumi.ResourceOption) (*SecretStaticAccount, error) {
	var resource SecretStaticAccount
	err := ctx.ReadResource("vault:gcp/secretStaticAccount:SecretStaticAccount", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SecretStaticAccount resources.
type secretStaticAccountState struct {
	// Path where the GCP Secrets Engine is mounted
	Backend *string `pulumi:"backend"`
	// Bindings to create for this static account. This can be specified multiple times for multiple bindings. Structure is documented below.
	Bindings []SecretStaticAccountBinding `pulumi:"bindings"`
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
	// *Available only for Vault Enterprise*.
	Namespace *string `pulumi:"namespace"`
	// Type of secret generated for this static account. Accepted values: `accessToken`, `serviceAccountKey`. Defaults to `accessToken`.
	SecretType *string `pulumi:"secretType"`
	// Email of the GCP service account to manage.
	ServiceAccountEmail *string `pulumi:"serviceAccountEmail"`
	// Project the service account belongs to.
	ServiceAccountProject *string `pulumi:"serviceAccountProject"`
	// Name of the Static Account to create
	StaticAccount *string `pulumi:"staticAccount"`
	// List of OAuth scopes to assign to `accessToken` secrets generated under this static account (`accessToken` static accounts only).
	TokenScopes []string `pulumi:"tokenScopes"`
}

type SecretStaticAccountState struct {
	// Path where the GCP Secrets Engine is mounted
	Backend pulumi.StringPtrInput
	// Bindings to create for this static account. This can be specified multiple times for multiple bindings. Structure is documented below.
	Bindings SecretStaticAccountBindingArrayInput
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
	// *Available only for Vault Enterprise*.
	Namespace pulumi.StringPtrInput
	// Type of secret generated for this static account. Accepted values: `accessToken`, `serviceAccountKey`. Defaults to `accessToken`.
	SecretType pulumi.StringPtrInput
	// Email of the GCP service account to manage.
	ServiceAccountEmail pulumi.StringPtrInput
	// Project the service account belongs to.
	ServiceAccountProject pulumi.StringPtrInput
	// Name of the Static Account to create
	StaticAccount pulumi.StringPtrInput
	// List of OAuth scopes to assign to `accessToken` secrets generated under this static account (`accessToken` static accounts only).
	TokenScopes pulumi.StringArrayInput
}

func (SecretStaticAccountState) ElementType() reflect.Type {
	return reflect.TypeOf((*secretStaticAccountState)(nil)).Elem()
}

type secretStaticAccountArgs struct {
	// Path where the GCP Secrets Engine is mounted
	Backend string `pulumi:"backend"`
	// Bindings to create for this static account. This can be specified multiple times for multiple bindings. Structure is documented below.
	Bindings []SecretStaticAccountBinding `pulumi:"bindings"`
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
	// *Available only for Vault Enterprise*.
	Namespace *string `pulumi:"namespace"`
	// Type of secret generated for this static account. Accepted values: `accessToken`, `serviceAccountKey`. Defaults to `accessToken`.
	SecretType *string `pulumi:"secretType"`
	// Email of the GCP service account to manage.
	ServiceAccountEmail string `pulumi:"serviceAccountEmail"`
	// Name of the Static Account to create
	StaticAccount string `pulumi:"staticAccount"`
	// List of OAuth scopes to assign to `accessToken` secrets generated under this static account (`accessToken` static accounts only).
	TokenScopes []string `pulumi:"tokenScopes"`
}

// The set of arguments for constructing a SecretStaticAccount resource.
type SecretStaticAccountArgs struct {
	// Path where the GCP Secrets Engine is mounted
	Backend pulumi.StringInput
	// Bindings to create for this static account. This can be specified multiple times for multiple bindings. Structure is documented below.
	Bindings SecretStaticAccountBindingArrayInput
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
	// *Available only for Vault Enterprise*.
	Namespace pulumi.StringPtrInput
	// Type of secret generated for this static account. Accepted values: `accessToken`, `serviceAccountKey`. Defaults to `accessToken`.
	SecretType pulumi.StringPtrInput
	// Email of the GCP service account to manage.
	ServiceAccountEmail pulumi.StringInput
	// Name of the Static Account to create
	StaticAccount pulumi.StringInput
	// List of OAuth scopes to assign to `accessToken` secrets generated under this static account (`accessToken` static accounts only).
	TokenScopes pulumi.StringArrayInput
}

func (SecretStaticAccountArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*secretStaticAccountArgs)(nil)).Elem()
}

type SecretStaticAccountInput interface {
	pulumi.Input

	ToSecretStaticAccountOutput() SecretStaticAccountOutput
	ToSecretStaticAccountOutputWithContext(ctx context.Context) SecretStaticAccountOutput
}

func (*SecretStaticAccount) ElementType() reflect.Type {
	return reflect.TypeOf((**SecretStaticAccount)(nil)).Elem()
}

func (i *SecretStaticAccount) ToSecretStaticAccountOutput() SecretStaticAccountOutput {
	return i.ToSecretStaticAccountOutputWithContext(context.Background())
}

func (i *SecretStaticAccount) ToSecretStaticAccountOutputWithContext(ctx context.Context) SecretStaticAccountOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecretStaticAccountOutput)
}

// SecretStaticAccountArrayInput is an input type that accepts SecretStaticAccountArray and SecretStaticAccountArrayOutput values.
// You can construct a concrete instance of `SecretStaticAccountArrayInput` via:
//
//	SecretStaticAccountArray{ SecretStaticAccountArgs{...} }
type SecretStaticAccountArrayInput interface {
	pulumi.Input

	ToSecretStaticAccountArrayOutput() SecretStaticAccountArrayOutput
	ToSecretStaticAccountArrayOutputWithContext(context.Context) SecretStaticAccountArrayOutput
}

type SecretStaticAccountArray []SecretStaticAccountInput

func (SecretStaticAccountArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SecretStaticAccount)(nil)).Elem()
}

func (i SecretStaticAccountArray) ToSecretStaticAccountArrayOutput() SecretStaticAccountArrayOutput {
	return i.ToSecretStaticAccountArrayOutputWithContext(context.Background())
}

func (i SecretStaticAccountArray) ToSecretStaticAccountArrayOutputWithContext(ctx context.Context) SecretStaticAccountArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecretStaticAccountArrayOutput)
}

// SecretStaticAccountMapInput is an input type that accepts SecretStaticAccountMap and SecretStaticAccountMapOutput values.
// You can construct a concrete instance of `SecretStaticAccountMapInput` via:
//
//	SecretStaticAccountMap{ "key": SecretStaticAccountArgs{...} }
type SecretStaticAccountMapInput interface {
	pulumi.Input

	ToSecretStaticAccountMapOutput() SecretStaticAccountMapOutput
	ToSecretStaticAccountMapOutputWithContext(context.Context) SecretStaticAccountMapOutput
}

type SecretStaticAccountMap map[string]SecretStaticAccountInput

func (SecretStaticAccountMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SecretStaticAccount)(nil)).Elem()
}

func (i SecretStaticAccountMap) ToSecretStaticAccountMapOutput() SecretStaticAccountMapOutput {
	return i.ToSecretStaticAccountMapOutputWithContext(context.Background())
}

func (i SecretStaticAccountMap) ToSecretStaticAccountMapOutputWithContext(ctx context.Context) SecretStaticAccountMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecretStaticAccountMapOutput)
}

type SecretStaticAccountOutput struct{ *pulumi.OutputState }

func (SecretStaticAccountOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SecretStaticAccount)(nil)).Elem()
}

func (o SecretStaticAccountOutput) ToSecretStaticAccountOutput() SecretStaticAccountOutput {
	return o
}

func (o SecretStaticAccountOutput) ToSecretStaticAccountOutputWithContext(ctx context.Context) SecretStaticAccountOutput {
	return o
}

// Path where the GCP Secrets Engine is mounted
func (o SecretStaticAccountOutput) Backend() pulumi.StringOutput {
	return o.ApplyT(func(v *SecretStaticAccount) pulumi.StringOutput { return v.Backend }).(pulumi.StringOutput)
}

// Bindings to create for this static account. This can be specified multiple times for multiple bindings. Structure is documented below.
func (o SecretStaticAccountOutput) Bindings() SecretStaticAccountBindingArrayOutput {
	return o.ApplyT(func(v *SecretStaticAccount) SecretStaticAccountBindingArrayOutput { return v.Bindings }).(SecretStaticAccountBindingArrayOutput)
}

// The namespace to provision the resource in.
// The value should not contain leading or trailing forward slashes.
// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
// *Available only for Vault Enterprise*.
func (o SecretStaticAccountOutput) Namespace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecretStaticAccount) pulumi.StringPtrOutput { return v.Namespace }).(pulumi.StringPtrOutput)
}

// Type of secret generated for this static account. Accepted values: `accessToken`, `serviceAccountKey`. Defaults to `accessToken`.
func (o SecretStaticAccountOutput) SecretType() pulumi.StringOutput {
	return o.ApplyT(func(v *SecretStaticAccount) pulumi.StringOutput { return v.SecretType }).(pulumi.StringOutput)
}

// Email of the GCP service account to manage.
func (o SecretStaticAccountOutput) ServiceAccountEmail() pulumi.StringOutput {
	return o.ApplyT(func(v *SecretStaticAccount) pulumi.StringOutput { return v.ServiceAccountEmail }).(pulumi.StringOutput)
}

// Project the service account belongs to.
func (o SecretStaticAccountOutput) ServiceAccountProject() pulumi.StringOutput {
	return o.ApplyT(func(v *SecretStaticAccount) pulumi.StringOutput { return v.ServiceAccountProject }).(pulumi.StringOutput)
}

// Name of the Static Account to create
func (o SecretStaticAccountOutput) StaticAccount() pulumi.StringOutput {
	return o.ApplyT(func(v *SecretStaticAccount) pulumi.StringOutput { return v.StaticAccount }).(pulumi.StringOutput)
}

// List of OAuth scopes to assign to `accessToken` secrets generated under this static account (`accessToken` static accounts only).
func (o SecretStaticAccountOutput) TokenScopes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *SecretStaticAccount) pulumi.StringArrayOutput { return v.TokenScopes }).(pulumi.StringArrayOutput)
}

type SecretStaticAccountArrayOutput struct{ *pulumi.OutputState }

func (SecretStaticAccountArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SecretStaticAccount)(nil)).Elem()
}

func (o SecretStaticAccountArrayOutput) ToSecretStaticAccountArrayOutput() SecretStaticAccountArrayOutput {
	return o
}

func (o SecretStaticAccountArrayOutput) ToSecretStaticAccountArrayOutputWithContext(ctx context.Context) SecretStaticAccountArrayOutput {
	return o
}

func (o SecretStaticAccountArrayOutput) Index(i pulumi.IntInput) SecretStaticAccountOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SecretStaticAccount {
		return vs[0].([]*SecretStaticAccount)[vs[1].(int)]
	}).(SecretStaticAccountOutput)
}

type SecretStaticAccountMapOutput struct{ *pulumi.OutputState }

func (SecretStaticAccountMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SecretStaticAccount)(nil)).Elem()
}

func (o SecretStaticAccountMapOutput) ToSecretStaticAccountMapOutput() SecretStaticAccountMapOutput {
	return o
}

func (o SecretStaticAccountMapOutput) ToSecretStaticAccountMapOutputWithContext(ctx context.Context) SecretStaticAccountMapOutput {
	return o
}

func (o SecretStaticAccountMapOutput) MapIndex(k pulumi.StringInput) SecretStaticAccountOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SecretStaticAccount {
		return vs[0].(map[string]*SecretStaticAccount)[vs[1].(string)]
	}).(SecretStaticAccountOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SecretStaticAccountInput)(nil)).Elem(), &SecretStaticAccount{})
	pulumi.RegisterInputType(reflect.TypeOf((*SecretStaticAccountArrayInput)(nil)).Elem(), SecretStaticAccountArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SecretStaticAccountMapInput)(nil)).Elem(), SecretStaticAccountMap{})
	pulumi.RegisterOutputType(SecretStaticAccountOutput{})
	pulumi.RegisterOutputType(SecretStaticAccountArrayOutput{})
	pulumi.RegisterOutputType(SecretStaticAccountMapOutput{})
}
