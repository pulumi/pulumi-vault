// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package gcp

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-vault/sdk/v5/go/vault/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Creates a Impersonated Account in the [GCP Secrets Engine](https://www.vaultproject.io/docs/secrets/gcp/index.html) for Vault.
//
// Each [impersonated account](https://www.vaultproject.io/docs/secrets/gcp/index.html#impersonated-accounts) is tied to a separately managed
// Service Account.
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
//	"github.com/pulumi/pulumi-gcp/sdk/v5/go/gcp/serviceAccount"
//	"github.com/pulumi/pulumi-vault/sdk/v5/go/vault/gcp"
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
//			this, err := serviceAccount.NewAccount(ctx, "this", &serviceAccount.AccountArgs{
//				AccountId: pulumi.String("my-awesome-account"),
//			})
//			if err != nil {
//				return err
//			}
//			gcp, err := gcp.NewSecretBackend(ctx, "gcp", &gcp.SecretBackendArgs{
//				Path:        pulumi.String("gcp"),
//				Credentials: readFileOrPanic("credentials.json"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = gcp.NewSecretImpersonatedAccount(ctx, "impersonatedAccount", &gcp.SecretImpersonatedAccountArgs{
//				Backend:             gcp.Path,
//				ImpersonatedAccount: pulumi.String("this"),
//				ServiceAccountEmail: this.Email,
//				TokenScopes: pulumi.StringArray{
//					pulumi.String("https://www.googleapis.com/auth/cloud-platform"),
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
// A impersonated account can be imported using its Vault Path. For example, referencing the example above,
//
// ```sh
//
//	$ pulumi import vault:gcp/secretImpersonatedAccount:SecretImpersonatedAccount impersonated_account gcp/impersonated-account/project_viewer
//
// ```
type SecretImpersonatedAccount struct {
	pulumi.CustomResourceState

	// Path where the GCP Secrets Engine is mounted
	Backend pulumi.StringOutput `pulumi:"backend"`
	// Name of the Impersonated Account to create
	ImpersonatedAccount pulumi.StringOutput `pulumi:"impersonatedAccount"`
	// Target namespace. (requires Enterprise)
	Namespace pulumi.StringPtrOutput `pulumi:"namespace"`
	// Email of the GCP service account to impersonate.
	ServiceAccountEmail pulumi.StringOutput `pulumi:"serviceAccountEmail"`
	// Project the service account belongs to.
	ServiceAccountProject pulumi.StringOutput `pulumi:"serviceAccountProject"`
	// List of OAuth scopes to assign to access tokens generated under this impersonated account.
	TokenScopes pulumi.StringArrayOutput `pulumi:"tokenScopes"`
}

// NewSecretImpersonatedAccount registers a new resource with the given unique name, arguments, and options.
func NewSecretImpersonatedAccount(ctx *pulumi.Context,
	name string, args *SecretImpersonatedAccountArgs, opts ...pulumi.ResourceOption) (*SecretImpersonatedAccount, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Backend == nil {
		return nil, errors.New("invalid value for required argument 'Backend'")
	}
	if args.ImpersonatedAccount == nil {
		return nil, errors.New("invalid value for required argument 'ImpersonatedAccount'")
	}
	if args.ServiceAccountEmail == nil {
		return nil, errors.New("invalid value for required argument 'ServiceAccountEmail'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource SecretImpersonatedAccount
	err := ctx.RegisterResource("vault:gcp/secretImpersonatedAccount:SecretImpersonatedAccount", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSecretImpersonatedAccount gets an existing SecretImpersonatedAccount resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSecretImpersonatedAccount(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SecretImpersonatedAccountState, opts ...pulumi.ResourceOption) (*SecretImpersonatedAccount, error) {
	var resource SecretImpersonatedAccount
	err := ctx.ReadResource("vault:gcp/secretImpersonatedAccount:SecretImpersonatedAccount", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SecretImpersonatedAccount resources.
type secretImpersonatedAccountState struct {
	// Path where the GCP Secrets Engine is mounted
	Backend *string `pulumi:"backend"`
	// Name of the Impersonated Account to create
	ImpersonatedAccount *string `pulumi:"impersonatedAccount"`
	// Target namespace. (requires Enterprise)
	Namespace *string `pulumi:"namespace"`
	// Email of the GCP service account to impersonate.
	ServiceAccountEmail *string `pulumi:"serviceAccountEmail"`
	// Project the service account belongs to.
	ServiceAccountProject *string `pulumi:"serviceAccountProject"`
	// List of OAuth scopes to assign to access tokens generated under this impersonated account.
	TokenScopes []string `pulumi:"tokenScopes"`
}

type SecretImpersonatedAccountState struct {
	// Path where the GCP Secrets Engine is mounted
	Backend pulumi.StringPtrInput
	// Name of the Impersonated Account to create
	ImpersonatedAccount pulumi.StringPtrInput
	// Target namespace. (requires Enterprise)
	Namespace pulumi.StringPtrInput
	// Email of the GCP service account to impersonate.
	ServiceAccountEmail pulumi.StringPtrInput
	// Project the service account belongs to.
	ServiceAccountProject pulumi.StringPtrInput
	// List of OAuth scopes to assign to access tokens generated under this impersonated account.
	TokenScopes pulumi.StringArrayInput
}

func (SecretImpersonatedAccountState) ElementType() reflect.Type {
	return reflect.TypeOf((*secretImpersonatedAccountState)(nil)).Elem()
}

type secretImpersonatedAccountArgs struct {
	// Path where the GCP Secrets Engine is mounted
	Backend string `pulumi:"backend"`
	// Name of the Impersonated Account to create
	ImpersonatedAccount string `pulumi:"impersonatedAccount"`
	// Target namespace. (requires Enterprise)
	Namespace *string `pulumi:"namespace"`
	// Email of the GCP service account to impersonate.
	ServiceAccountEmail string `pulumi:"serviceAccountEmail"`
	// List of OAuth scopes to assign to access tokens generated under this impersonated account.
	TokenScopes []string `pulumi:"tokenScopes"`
}

// The set of arguments for constructing a SecretImpersonatedAccount resource.
type SecretImpersonatedAccountArgs struct {
	// Path where the GCP Secrets Engine is mounted
	Backend pulumi.StringInput
	// Name of the Impersonated Account to create
	ImpersonatedAccount pulumi.StringInput
	// Target namespace. (requires Enterprise)
	Namespace pulumi.StringPtrInput
	// Email of the GCP service account to impersonate.
	ServiceAccountEmail pulumi.StringInput
	// List of OAuth scopes to assign to access tokens generated under this impersonated account.
	TokenScopes pulumi.StringArrayInput
}

func (SecretImpersonatedAccountArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*secretImpersonatedAccountArgs)(nil)).Elem()
}

type SecretImpersonatedAccountInput interface {
	pulumi.Input

	ToSecretImpersonatedAccountOutput() SecretImpersonatedAccountOutput
	ToSecretImpersonatedAccountOutputWithContext(ctx context.Context) SecretImpersonatedAccountOutput
}

func (*SecretImpersonatedAccount) ElementType() reflect.Type {
	return reflect.TypeOf((**SecretImpersonatedAccount)(nil)).Elem()
}

func (i *SecretImpersonatedAccount) ToSecretImpersonatedAccountOutput() SecretImpersonatedAccountOutput {
	return i.ToSecretImpersonatedAccountOutputWithContext(context.Background())
}

func (i *SecretImpersonatedAccount) ToSecretImpersonatedAccountOutputWithContext(ctx context.Context) SecretImpersonatedAccountOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecretImpersonatedAccountOutput)
}

func (i *SecretImpersonatedAccount) ToOutput(ctx context.Context) pulumix.Output[*SecretImpersonatedAccount] {
	return pulumix.Output[*SecretImpersonatedAccount]{
		OutputState: i.ToSecretImpersonatedAccountOutputWithContext(ctx).OutputState,
	}
}

// SecretImpersonatedAccountArrayInput is an input type that accepts SecretImpersonatedAccountArray and SecretImpersonatedAccountArrayOutput values.
// You can construct a concrete instance of `SecretImpersonatedAccountArrayInput` via:
//
//	SecretImpersonatedAccountArray{ SecretImpersonatedAccountArgs{...} }
type SecretImpersonatedAccountArrayInput interface {
	pulumi.Input

	ToSecretImpersonatedAccountArrayOutput() SecretImpersonatedAccountArrayOutput
	ToSecretImpersonatedAccountArrayOutputWithContext(context.Context) SecretImpersonatedAccountArrayOutput
}

type SecretImpersonatedAccountArray []SecretImpersonatedAccountInput

func (SecretImpersonatedAccountArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SecretImpersonatedAccount)(nil)).Elem()
}

func (i SecretImpersonatedAccountArray) ToSecretImpersonatedAccountArrayOutput() SecretImpersonatedAccountArrayOutput {
	return i.ToSecretImpersonatedAccountArrayOutputWithContext(context.Background())
}

func (i SecretImpersonatedAccountArray) ToSecretImpersonatedAccountArrayOutputWithContext(ctx context.Context) SecretImpersonatedAccountArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecretImpersonatedAccountArrayOutput)
}

func (i SecretImpersonatedAccountArray) ToOutput(ctx context.Context) pulumix.Output[[]*SecretImpersonatedAccount] {
	return pulumix.Output[[]*SecretImpersonatedAccount]{
		OutputState: i.ToSecretImpersonatedAccountArrayOutputWithContext(ctx).OutputState,
	}
}

// SecretImpersonatedAccountMapInput is an input type that accepts SecretImpersonatedAccountMap and SecretImpersonatedAccountMapOutput values.
// You can construct a concrete instance of `SecretImpersonatedAccountMapInput` via:
//
//	SecretImpersonatedAccountMap{ "key": SecretImpersonatedAccountArgs{...} }
type SecretImpersonatedAccountMapInput interface {
	pulumi.Input

	ToSecretImpersonatedAccountMapOutput() SecretImpersonatedAccountMapOutput
	ToSecretImpersonatedAccountMapOutputWithContext(context.Context) SecretImpersonatedAccountMapOutput
}

type SecretImpersonatedAccountMap map[string]SecretImpersonatedAccountInput

func (SecretImpersonatedAccountMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SecretImpersonatedAccount)(nil)).Elem()
}

func (i SecretImpersonatedAccountMap) ToSecretImpersonatedAccountMapOutput() SecretImpersonatedAccountMapOutput {
	return i.ToSecretImpersonatedAccountMapOutputWithContext(context.Background())
}

func (i SecretImpersonatedAccountMap) ToSecretImpersonatedAccountMapOutputWithContext(ctx context.Context) SecretImpersonatedAccountMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecretImpersonatedAccountMapOutput)
}

func (i SecretImpersonatedAccountMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*SecretImpersonatedAccount] {
	return pulumix.Output[map[string]*SecretImpersonatedAccount]{
		OutputState: i.ToSecretImpersonatedAccountMapOutputWithContext(ctx).OutputState,
	}
}

type SecretImpersonatedAccountOutput struct{ *pulumi.OutputState }

func (SecretImpersonatedAccountOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SecretImpersonatedAccount)(nil)).Elem()
}

func (o SecretImpersonatedAccountOutput) ToSecretImpersonatedAccountOutput() SecretImpersonatedAccountOutput {
	return o
}

func (o SecretImpersonatedAccountOutput) ToSecretImpersonatedAccountOutputWithContext(ctx context.Context) SecretImpersonatedAccountOutput {
	return o
}

func (o SecretImpersonatedAccountOutput) ToOutput(ctx context.Context) pulumix.Output[*SecretImpersonatedAccount] {
	return pulumix.Output[*SecretImpersonatedAccount]{
		OutputState: o.OutputState,
	}
}

// Path where the GCP Secrets Engine is mounted
func (o SecretImpersonatedAccountOutput) Backend() pulumi.StringOutput {
	return o.ApplyT(func(v *SecretImpersonatedAccount) pulumi.StringOutput { return v.Backend }).(pulumi.StringOutput)
}

// Name of the Impersonated Account to create
func (o SecretImpersonatedAccountOutput) ImpersonatedAccount() pulumi.StringOutput {
	return o.ApplyT(func(v *SecretImpersonatedAccount) pulumi.StringOutput { return v.ImpersonatedAccount }).(pulumi.StringOutput)
}

// Target namespace. (requires Enterprise)
func (o SecretImpersonatedAccountOutput) Namespace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecretImpersonatedAccount) pulumi.StringPtrOutput { return v.Namespace }).(pulumi.StringPtrOutput)
}

// Email of the GCP service account to impersonate.
func (o SecretImpersonatedAccountOutput) ServiceAccountEmail() pulumi.StringOutput {
	return o.ApplyT(func(v *SecretImpersonatedAccount) pulumi.StringOutput { return v.ServiceAccountEmail }).(pulumi.StringOutput)
}

// Project the service account belongs to.
func (o SecretImpersonatedAccountOutput) ServiceAccountProject() pulumi.StringOutput {
	return o.ApplyT(func(v *SecretImpersonatedAccount) pulumi.StringOutput { return v.ServiceAccountProject }).(pulumi.StringOutput)
}

// List of OAuth scopes to assign to access tokens generated under this impersonated account.
func (o SecretImpersonatedAccountOutput) TokenScopes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *SecretImpersonatedAccount) pulumi.StringArrayOutput { return v.TokenScopes }).(pulumi.StringArrayOutput)
}

type SecretImpersonatedAccountArrayOutput struct{ *pulumi.OutputState }

func (SecretImpersonatedAccountArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SecretImpersonatedAccount)(nil)).Elem()
}

func (o SecretImpersonatedAccountArrayOutput) ToSecretImpersonatedAccountArrayOutput() SecretImpersonatedAccountArrayOutput {
	return o
}

func (o SecretImpersonatedAccountArrayOutput) ToSecretImpersonatedAccountArrayOutputWithContext(ctx context.Context) SecretImpersonatedAccountArrayOutput {
	return o
}

func (o SecretImpersonatedAccountArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*SecretImpersonatedAccount] {
	return pulumix.Output[[]*SecretImpersonatedAccount]{
		OutputState: o.OutputState,
	}
}

func (o SecretImpersonatedAccountArrayOutput) Index(i pulumi.IntInput) SecretImpersonatedAccountOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SecretImpersonatedAccount {
		return vs[0].([]*SecretImpersonatedAccount)[vs[1].(int)]
	}).(SecretImpersonatedAccountOutput)
}

type SecretImpersonatedAccountMapOutput struct{ *pulumi.OutputState }

func (SecretImpersonatedAccountMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SecretImpersonatedAccount)(nil)).Elem()
}

func (o SecretImpersonatedAccountMapOutput) ToSecretImpersonatedAccountMapOutput() SecretImpersonatedAccountMapOutput {
	return o
}

func (o SecretImpersonatedAccountMapOutput) ToSecretImpersonatedAccountMapOutputWithContext(ctx context.Context) SecretImpersonatedAccountMapOutput {
	return o
}

func (o SecretImpersonatedAccountMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*SecretImpersonatedAccount] {
	return pulumix.Output[map[string]*SecretImpersonatedAccount]{
		OutputState: o.OutputState,
	}
}

func (o SecretImpersonatedAccountMapOutput) MapIndex(k pulumi.StringInput) SecretImpersonatedAccountOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SecretImpersonatedAccount {
		return vs[0].(map[string]*SecretImpersonatedAccount)[vs[1].(string)]
	}).(SecretImpersonatedAccountOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SecretImpersonatedAccountInput)(nil)).Elem(), &SecretImpersonatedAccount{})
	pulumi.RegisterInputType(reflect.TypeOf((*SecretImpersonatedAccountArrayInput)(nil)).Elem(), SecretImpersonatedAccountArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SecretImpersonatedAccountMapInput)(nil)).Elem(), SecretImpersonatedAccountMap{})
	pulumi.RegisterOutputType(SecretImpersonatedAccountOutput{})
	pulumi.RegisterOutputType(SecretImpersonatedAccountArrayOutput{})
	pulumi.RegisterOutputType(SecretImpersonatedAccountMapOutput{})
}
