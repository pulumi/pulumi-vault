// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package gcp

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-vault/sdk/v6/go/vault/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a resource to configure the [GCP auth backend within Vault](https://www.vaultproject.io/docs/auth/gcp.html).
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"os"
//
//	"github.com/pulumi/pulumi-vault/sdk/v6/go/vault/gcp"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func readFileOrPanic(path string) pulumi.StringPtrInput {
//		data, err := os.ReadFile(path)
//		if err != nil {
//			panic(err.Error())
//		}
//		return pulumi.String(string(data))
//	}
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := gcp.NewAuthBackend(ctx, "gcp", &gcp.AuthBackendArgs{
//				Credentials: readFileOrPanic("vault-gcp-credentials.json"),
//				CustomEndpoint: &gcp.AuthBackendCustomEndpointArgs{
//					Api:     pulumi.String("www.googleapis.com"),
//					Iam:     pulumi.String("iam.googleapis.com"),
//					Crm:     pulumi.String("cloudresourcemanager.googleapis.com"),
//					Compute: pulumi.String("compute.googleapis.com"),
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
// GCP authentication backends can be imported using the backend name, e.g.
//
// ```sh
// $ pulumi import vault:gcp/authBackend:AuthBackend gcp gcp
// ```
type AuthBackend struct {
	pulumi.CustomResourceState

	// The mount accessor related to the auth mount. It is useful for integration with [Identity Secrets Engine](https://www.vaultproject.io/docs/secrets/identity/index.html).
	Accessor pulumi.StringOutput `pulumi:"accessor"`
	// The clients email associated with the credentials
	ClientEmail pulumi.StringOutput `pulumi:"clientEmail"`
	// The Client ID of the credentials
	ClientId pulumi.StringOutput `pulumi:"clientId"`
	// A JSON string containing the contents of a GCP credentials file. If this value is empty, Vault will try to use Application Default Credentials from the machine on which the Vault server is running.
	Credentials pulumi.StringPtrOutput `pulumi:"credentials"`
	// Specifies overrides to
	// [service endpoints](https://cloud.google.com/apis/design/glossary#api_service_endpoint)
	// used when making API requests. This allows specific requests made during authentication
	// to target alternative service endpoints for use in [Private Google Access](https://cloud.google.com/vpc/docs/configure-private-google-access)
	// environments. Requires Vault 1.11+.
	//
	// Overrides are set at the subdomain level using the following keys:
	CustomEndpoint AuthBackendCustomEndpointPtrOutput `pulumi:"customEndpoint"`
	// A description of the auth method.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// If set, opts out of mount migration on path updates.
	// See here for more info on [Mount Migration](https://www.vaultproject.io/docs/concepts/mount-migration)
	DisableRemount pulumi.BoolPtrOutput `pulumi:"disableRemount"`
	// Specifies if the auth method is local only.
	Local pulumi.BoolPtrOutput `pulumi:"local"`
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
	// *Available only for Vault Enterprise*.
	Namespace pulumi.StringPtrOutput `pulumi:"namespace"`
	// The path to mount the auth method — this defaults to 'gcp'.
	Path pulumi.StringPtrOutput `pulumi:"path"`
	// The ID of the private key from the credentials
	PrivateKeyId pulumi.StringOutput `pulumi:"privateKeyId"`
	// The GCP Project ID
	ProjectId pulumi.StringOutput `pulumi:"projectId"`
	// Extra configuration block. Structure is documented below.
	//
	// The `tune` block is used to tune the auth backend:
	Tune AuthBackendTuneOutput `pulumi:"tune"`
}

// NewAuthBackend registers a new resource with the given unique name, arguments, and options.
func NewAuthBackend(ctx *pulumi.Context,
	name string, args *AuthBackendArgs, opts ...pulumi.ResourceOption) (*AuthBackend, error) {
	if args == nil {
		args = &AuthBackendArgs{}
	}

	if args.Credentials != nil {
		args.Credentials = pulumi.ToSecret(args.Credentials).(pulumi.StringPtrInput)
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"credentials",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource AuthBackend
	err := ctx.RegisterResource("vault:gcp/authBackend:AuthBackend", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAuthBackend gets an existing AuthBackend resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAuthBackend(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AuthBackendState, opts ...pulumi.ResourceOption) (*AuthBackend, error) {
	var resource AuthBackend
	err := ctx.ReadResource("vault:gcp/authBackend:AuthBackend", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AuthBackend resources.
type authBackendState struct {
	// The mount accessor related to the auth mount. It is useful for integration with [Identity Secrets Engine](https://www.vaultproject.io/docs/secrets/identity/index.html).
	Accessor *string `pulumi:"accessor"`
	// The clients email associated with the credentials
	ClientEmail *string `pulumi:"clientEmail"`
	// The Client ID of the credentials
	ClientId *string `pulumi:"clientId"`
	// A JSON string containing the contents of a GCP credentials file. If this value is empty, Vault will try to use Application Default Credentials from the machine on which the Vault server is running.
	Credentials *string `pulumi:"credentials"`
	// Specifies overrides to
	// [service endpoints](https://cloud.google.com/apis/design/glossary#api_service_endpoint)
	// used when making API requests. This allows specific requests made during authentication
	// to target alternative service endpoints for use in [Private Google Access](https://cloud.google.com/vpc/docs/configure-private-google-access)
	// environments. Requires Vault 1.11+.
	//
	// Overrides are set at the subdomain level using the following keys:
	CustomEndpoint *AuthBackendCustomEndpoint `pulumi:"customEndpoint"`
	// A description of the auth method.
	Description *string `pulumi:"description"`
	// If set, opts out of mount migration on path updates.
	// See here for more info on [Mount Migration](https://www.vaultproject.io/docs/concepts/mount-migration)
	DisableRemount *bool `pulumi:"disableRemount"`
	// Specifies if the auth method is local only.
	Local *bool `pulumi:"local"`
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
	// *Available only for Vault Enterprise*.
	Namespace *string `pulumi:"namespace"`
	// The path to mount the auth method — this defaults to 'gcp'.
	Path *string `pulumi:"path"`
	// The ID of the private key from the credentials
	PrivateKeyId *string `pulumi:"privateKeyId"`
	// The GCP Project ID
	ProjectId *string `pulumi:"projectId"`
	// Extra configuration block. Structure is documented below.
	//
	// The `tune` block is used to tune the auth backend:
	Tune *AuthBackendTune `pulumi:"tune"`
}

type AuthBackendState struct {
	// The mount accessor related to the auth mount. It is useful for integration with [Identity Secrets Engine](https://www.vaultproject.io/docs/secrets/identity/index.html).
	Accessor pulumi.StringPtrInput
	// The clients email associated with the credentials
	ClientEmail pulumi.StringPtrInput
	// The Client ID of the credentials
	ClientId pulumi.StringPtrInput
	// A JSON string containing the contents of a GCP credentials file. If this value is empty, Vault will try to use Application Default Credentials from the machine on which the Vault server is running.
	Credentials pulumi.StringPtrInput
	// Specifies overrides to
	// [service endpoints](https://cloud.google.com/apis/design/glossary#api_service_endpoint)
	// used when making API requests. This allows specific requests made during authentication
	// to target alternative service endpoints for use in [Private Google Access](https://cloud.google.com/vpc/docs/configure-private-google-access)
	// environments. Requires Vault 1.11+.
	//
	// Overrides are set at the subdomain level using the following keys:
	CustomEndpoint AuthBackendCustomEndpointPtrInput
	// A description of the auth method.
	Description pulumi.StringPtrInput
	// If set, opts out of mount migration on path updates.
	// See here for more info on [Mount Migration](https://www.vaultproject.io/docs/concepts/mount-migration)
	DisableRemount pulumi.BoolPtrInput
	// Specifies if the auth method is local only.
	Local pulumi.BoolPtrInput
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
	// *Available only for Vault Enterprise*.
	Namespace pulumi.StringPtrInput
	// The path to mount the auth method — this defaults to 'gcp'.
	Path pulumi.StringPtrInput
	// The ID of the private key from the credentials
	PrivateKeyId pulumi.StringPtrInput
	// The GCP Project ID
	ProjectId pulumi.StringPtrInput
	// Extra configuration block. Structure is documented below.
	//
	// The `tune` block is used to tune the auth backend:
	Tune AuthBackendTunePtrInput
}

func (AuthBackendState) ElementType() reflect.Type {
	return reflect.TypeOf((*authBackendState)(nil)).Elem()
}

type authBackendArgs struct {
	// The clients email associated with the credentials
	ClientEmail *string `pulumi:"clientEmail"`
	// The Client ID of the credentials
	ClientId *string `pulumi:"clientId"`
	// A JSON string containing the contents of a GCP credentials file. If this value is empty, Vault will try to use Application Default Credentials from the machine on which the Vault server is running.
	Credentials *string `pulumi:"credentials"`
	// Specifies overrides to
	// [service endpoints](https://cloud.google.com/apis/design/glossary#api_service_endpoint)
	// used when making API requests. This allows specific requests made during authentication
	// to target alternative service endpoints for use in [Private Google Access](https://cloud.google.com/vpc/docs/configure-private-google-access)
	// environments. Requires Vault 1.11+.
	//
	// Overrides are set at the subdomain level using the following keys:
	CustomEndpoint *AuthBackendCustomEndpoint `pulumi:"customEndpoint"`
	// A description of the auth method.
	Description *string `pulumi:"description"`
	// If set, opts out of mount migration on path updates.
	// See here for more info on [Mount Migration](https://www.vaultproject.io/docs/concepts/mount-migration)
	DisableRemount *bool `pulumi:"disableRemount"`
	// Specifies if the auth method is local only.
	Local *bool `pulumi:"local"`
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
	// *Available only for Vault Enterprise*.
	Namespace *string `pulumi:"namespace"`
	// The path to mount the auth method — this defaults to 'gcp'.
	Path *string `pulumi:"path"`
	// The ID of the private key from the credentials
	PrivateKeyId *string `pulumi:"privateKeyId"`
	// The GCP Project ID
	ProjectId *string `pulumi:"projectId"`
	// Extra configuration block. Structure is documented below.
	//
	// The `tune` block is used to tune the auth backend:
	Tune *AuthBackendTune `pulumi:"tune"`
}

// The set of arguments for constructing a AuthBackend resource.
type AuthBackendArgs struct {
	// The clients email associated with the credentials
	ClientEmail pulumi.StringPtrInput
	// The Client ID of the credentials
	ClientId pulumi.StringPtrInput
	// A JSON string containing the contents of a GCP credentials file. If this value is empty, Vault will try to use Application Default Credentials from the machine on which the Vault server is running.
	Credentials pulumi.StringPtrInput
	// Specifies overrides to
	// [service endpoints](https://cloud.google.com/apis/design/glossary#api_service_endpoint)
	// used when making API requests. This allows specific requests made during authentication
	// to target alternative service endpoints for use in [Private Google Access](https://cloud.google.com/vpc/docs/configure-private-google-access)
	// environments. Requires Vault 1.11+.
	//
	// Overrides are set at the subdomain level using the following keys:
	CustomEndpoint AuthBackendCustomEndpointPtrInput
	// A description of the auth method.
	Description pulumi.StringPtrInput
	// If set, opts out of mount migration on path updates.
	// See here for more info on [Mount Migration](https://www.vaultproject.io/docs/concepts/mount-migration)
	DisableRemount pulumi.BoolPtrInput
	// Specifies if the auth method is local only.
	Local pulumi.BoolPtrInput
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
	// *Available only for Vault Enterprise*.
	Namespace pulumi.StringPtrInput
	// The path to mount the auth method — this defaults to 'gcp'.
	Path pulumi.StringPtrInput
	// The ID of the private key from the credentials
	PrivateKeyId pulumi.StringPtrInput
	// The GCP Project ID
	ProjectId pulumi.StringPtrInput
	// Extra configuration block. Structure is documented below.
	//
	// The `tune` block is used to tune the auth backend:
	Tune AuthBackendTunePtrInput
}

func (AuthBackendArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*authBackendArgs)(nil)).Elem()
}

type AuthBackendInput interface {
	pulumi.Input

	ToAuthBackendOutput() AuthBackendOutput
	ToAuthBackendOutputWithContext(ctx context.Context) AuthBackendOutput
}

func (*AuthBackend) ElementType() reflect.Type {
	return reflect.TypeOf((**AuthBackend)(nil)).Elem()
}

func (i *AuthBackend) ToAuthBackendOutput() AuthBackendOutput {
	return i.ToAuthBackendOutputWithContext(context.Background())
}

func (i *AuthBackend) ToAuthBackendOutputWithContext(ctx context.Context) AuthBackendOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AuthBackendOutput)
}

// AuthBackendArrayInput is an input type that accepts AuthBackendArray and AuthBackendArrayOutput values.
// You can construct a concrete instance of `AuthBackendArrayInput` via:
//
//	AuthBackendArray{ AuthBackendArgs{...} }
type AuthBackendArrayInput interface {
	pulumi.Input

	ToAuthBackendArrayOutput() AuthBackendArrayOutput
	ToAuthBackendArrayOutputWithContext(context.Context) AuthBackendArrayOutput
}

type AuthBackendArray []AuthBackendInput

func (AuthBackendArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AuthBackend)(nil)).Elem()
}

func (i AuthBackendArray) ToAuthBackendArrayOutput() AuthBackendArrayOutput {
	return i.ToAuthBackendArrayOutputWithContext(context.Background())
}

func (i AuthBackendArray) ToAuthBackendArrayOutputWithContext(ctx context.Context) AuthBackendArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AuthBackendArrayOutput)
}

// AuthBackendMapInput is an input type that accepts AuthBackendMap and AuthBackendMapOutput values.
// You can construct a concrete instance of `AuthBackendMapInput` via:
//
//	AuthBackendMap{ "key": AuthBackendArgs{...} }
type AuthBackendMapInput interface {
	pulumi.Input

	ToAuthBackendMapOutput() AuthBackendMapOutput
	ToAuthBackendMapOutputWithContext(context.Context) AuthBackendMapOutput
}

type AuthBackendMap map[string]AuthBackendInput

func (AuthBackendMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AuthBackend)(nil)).Elem()
}

func (i AuthBackendMap) ToAuthBackendMapOutput() AuthBackendMapOutput {
	return i.ToAuthBackendMapOutputWithContext(context.Background())
}

func (i AuthBackendMap) ToAuthBackendMapOutputWithContext(ctx context.Context) AuthBackendMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AuthBackendMapOutput)
}

type AuthBackendOutput struct{ *pulumi.OutputState }

func (AuthBackendOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AuthBackend)(nil)).Elem()
}

func (o AuthBackendOutput) ToAuthBackendOutput() AuthBackendOutput {
	return o
}

func (o AuthBackendOutput) ToAuthBackendOutputWithContext(ctx context.Context) AuthBackendOutput {
	return o
}

// The mount accessor related to the auth mount. It is useful for integration with [Identity Secrets Engine](https://www.vaultproject.io/docs/secrets/identity/index.html).
func (o AuthBackendOutput) Accessor() pulumi.StringOutput {
	return o.ApplyT(func(v *AuthBackend) pulumi.StringOutput { return v.Accessor }).(pulumi.StringOutput)
}

// The clients email associated with the credentials
func (o AuthBackendOutput) ClientEmail() pulumi.StringOutput {
	return o.ApplyT(func(v *AuthBackend) pulumi.StringOutput { return v.ClientEmail }).(pulumi.StringOutput)
}

// The Client ID of the credentials
func (o AuthBackendOutput) ClientId() pulumi.StringOutput {
	return o.ApplyT(func(v *AuthBackend) pulumi.StringOutput { return v.ClientId }).(pulumi.StringOutput)
}

// A JSON string containing the contents of a GCP credentials file. If this value is empty, Vault will try to use Application Default Credentials from the machine on which the Vault server is running.
func (o AuthBackendOutput) Credentials() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AuthBackend) pulumi.StringPtrOutput { return v.Credentials }).(pulumi.StringPtrOutput)
}

// Specifies overrides to
// [service endpoints](https://cloud.google.com/apis/design/glossary#api_service_endpoint)
// used when making API requests. This allows specific requests made during authentication
// to target alternative service endpoints for use in [Private Google Access](https://cloud.google.com/vpc/docs/configure-private-google-access)
// environments. Requires Vault 1.11+.
//
// Overrides are set at the subdomain level using the following keys:
func (o AuthBackendOutput) CustomEndpoint() AuthBackendCustomEndpointPtrOutput {
	return o.ApplyT(func(v *AuthBackend) AuthBackendCustomEndpointPtrOutput { return v.CustomEndpoint }).(AuthBackendCustomEndpointPtrOutput)
}

// A description of the auth method.
func (o AuthBackendOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AuthBackend) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// If set, opts out of mount migration on path updates.
// See here for more info on [Mount Migration](https://www.vaultproject.io/docs/concepts/mount-migration)
func (o AuthBackendOutput) DisableRemount() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AuthBackend) pulumi.BoolPtrOutput { return v.DisableRemount }).(pulumi.BoolPtrOutput)
}

// Specifies if the auth method is local only.
func (o AuthBackendOutput) Local() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AuthBackend) pulumi.BoolPtrOutput { return v.Local }).(pulumi.BoolPtrOutput)
}

// The namespace to provision the resource in.
// The value should not contain leading or trailing forward slashes.
// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
// *Available only for Vault Enterprise*.
func (o AuthBackendOutput) Namespace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AuthBackend) pulumi.StringPtrOutput { return v.Namespace }).(pulumi.StringPtrOutput)
}

// The path to mount the auth method — this defaults to 'gcp'.
func (o AuthBackendOutput) Path() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AuthBackend) pulumi.StringPtrOutput { return v.Path }).(pulumi.StringPtrOutput)
}

// The ID of the private key from the credentials
func (o AuthBackendOutput) PrivateKeyId() pulumi.StringOutput {
	return o.ApplyT(func(v *AuthBackend) pulumi.StringOutput { return v.PrivateKeyId }).(pulumi.StringOutput)
}

// The GCP Project ID
func (o AuthBackendOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v *AuthBackend) pulumi.StringOutput { return v.ProjectId }).(pulumi.StringOutput)
}

// Extra configuration block. Structure is documented below.
//
// The `tune` block is used to tune the auth backend:
func (o AuthBackendOutput) Tune() AuthBackendTuneOutput {
	return o.ApplyT(func(v *AuthBackend) AuthBackendTuneOutput { return v.Tune }).(AuthBackendTuneOutput)
}

type AuthBackendArrayOutput struct{ *pulumi.OutputState }

func (AuthBackendArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AuthBackend)(nil)).Elem()
}

func (o AuthBackendArrayOutput) ToAuthBackendArrayOutput() AuthBackendArrayOutput {
	return o
}

func (o AuthBackendArrayOutput) ToAuthBackendArrayOutputWithContext(ctx context.Context) AuthBackendArrayOutput {
	return o
}

func (o AuthBackendArrayOutput) Index(i pulumi.IntInput) AuthBackendOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *AuthBackend {
		return vs[0].([]*AuthBackend)[vs[1].(int)]
	}).(AuthBackendOutput)
}

type AuthBackendMapOutput struct{ *pulumi.OutputState }

func (AuthBackendMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AuthBackend)(nil)).Elem()
}

func (o AuthBackendMapOutput) ToAuthBackendMapOutput() AuthBackendMapOutput {
	return o
}

func (o AuthBackendMapOutput) ToAuthBackendMapOutputWithContext(ctx context.Context) AuthBackendMapOutput {
	return o
}

func (o AuthBackendMapOutput) MapIndex(k pulumi.StringInput) AuthBackendOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *AuthBackend {
		return vs[0].(map[string]*AuthBackend)[vs[1].(string)]
	}).(AuthBackendOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AuthBackendInput)(nil)).Elem(), &AuthBackend{})
	pulumi.RegisterInputType(reflect.TypeOf((*AuthBackendArrayInput)(nil)).Elem(), AuthBackendArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AuthBackendMapInput)(nil)).Elem(), AuthBackendMap{})
	pulumi.RegisterOutputType(AuthBackendOutput{})
	pulumi.RegisterOutputType(AuthBackendArrayOutput{})
	pulumi.RegisterOutputType(AuthBackendMapOutput{})
}
