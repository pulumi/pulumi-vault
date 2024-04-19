// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package secrets

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
//	"os"
//
//	"github.com/pulumi/pulumi-vault/sdk/v6/go/vault/secrets"
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
//			_, err := secrets.NewSyncGithubApps(ctx, "github-apps", &secrets.SyncGithubAppsArgs{
//				AppId:      pulumi.Any(_var.App_id),
//				PrivateKey: readFileOrPanic(_var.Privatekey_file),
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
// GitHub Apps Secrets sync configuration endpoint can be imported using the `name`, e.g.
//
// ```sh
// $ pulumi import vault:secrets/syncGithubApps:SyncGithubApps gh github-apps
// ```
type SyncGithubApps struct {
	pulumi.CustomResourceState

	// The GitHub application ID.
	AppId pulumi.IntOutput `pulumi:"appId"`
	// A fingerprint of a private key.
	Fingerprint pulumi.StringOutput `pulumi:"fingerprint"`
	// The user-defined name of the GitHub App configuration.
	Name pulumi.StringOutput `pulumi:"name"`
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
	Namespace pulumi.StringPtrOutput `pulumi:"namespace"`
	// The content of a PEM formatted private key generated on GitHub for the app.
	PrivateKey pulumi.StringOutput `pulumi:"privateKey"`
}

// NewSyncGithubApps registers a new resource with the given unique name, arguments, and options.
func NewSyncGithubApps(ctx *pulumi.Context,
	name string, args *SyncGithubAppsArgs, opts ...pulumi.ResourceOption) (*SyncGithubApps, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AppId == nil {
		return nil, errors.New("invalid value for required argument 'AppId'")
	}
	if args.PrivateKey == nil {
		return nil, errors.New("invalid value for required argument 'PrivateKey'")
	}
	if args.PrivateKey != nil {
		args.PrivateKey = pulumi.ToSecret(args.PrivateKey).(pulumi.StringInput)
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"privateKey",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource SyncGithubApps
	err := ctx.RegisterResource("vault:secrets/syncGithubApps:SyncGithubApps", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSyncGithubApps gets an existing SyncGithubApps resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSyncGithubApps(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SyncGithubAppsState, opts ...pulumi.ResourceOption) (*SyncGithubApps, error) {
	var resource SyncGithubApps
	err := ctx.ReadResource("vault:secrets/syncGithubApps:SyncGithubApps", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SyncGithubApps resources.
type syncGithubAppsState struct {
	// The GitHub application ID.
	AppId *int `pulumi:"appId"`
	// A fingerprint of a private key.
	Fingerprint *string `pulumi:"fingerprint"`
	// The user-defined name of the GitHub App configuration.
	Name *string `pulumi:"name"`
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
	Namespace *string `pulumi:"namespace"`
	// The content of a PEM formatted private key generated on GitHub for the app.
	PrivateKey *string `pulumi:"privateKey"`
}

type SyncGithubAppsState struct {
	// The GitHub application ID.
	AppId pulumi.IntPtrInput
	// A fingerprint of a private key.
	Fingerprint pulumi.StringPtrInput
	// The user-defined name of the GitHub App configuration.
	Name pulumi.StringPtrInput
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
	Namespace pulumi.StringPtrInput
	// The content of a PEM formatted private key generated on GitHub for the app.
	PrivateKey pulumi.StringPtrInput
}

func (SyncGithubAppsState) ElementType() reflect.Type {
	return reflect.TypeOf((*syncGithubAppsState)(nil)).Elem()
}

type syncGithubAppsArgs struct {
	// The GitHub application ID.
	AppId int `pulumi:"appId"`
	// The user-defined name of the GitHub App configuration.
	Name *string `pulumi:"name"`
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
	Namespace *string `pulumi:"namespace"`
	// The content of a PEM formatted private key generated on GitHub for the app.
	PrivateKey string `pulumi:"privateKey"`
}

// The set of arguments for constructing a SyncGithubApps resource.
type SyncGithubAppsArgs struct {
	// The GitHub application ID.
	AppId pulumi.IntInput
	// The user-defined name of the GitHub App configuration.
	Name pulumi.StringPtrInput
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
	Namespace pulumi.StringPtrInput
	// The content of a PEM formatted private key generated on GitHub for the app.
	PrivateKey pulumi.StringInput
}

func (SyncGithubAppsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*syncGithubAppsArgs)(nil)).Elem()
}

type SyncGithubAppsInput interface {
	pulumi.Input

	ToSyncGithubAppsOutput() SyncGithubAppsOutput
	ToSyncGithubAppsOutputWithContext(ctx context.Context) SyncGithubAppsOutput
}

func (*SyncGithubApps) ElementType() reflect.Type {
	return reflect.TypeOf((**SyncGithubApps)(nil)).Elem()
}

func (i *SyncGithubApps) ToSyncGithubAppsOutput() SyncGithubAppsOutput {
	return i.ToSyncGithubAppsOutputWithContext(context.Background())
}

func (i *SyncGithubApps) ToSyncGithubAppsOutputWithContext(ctx context.Context) SyncGithubAppsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SyncGithubAppsOutput)
}

// SyncGithubAppsArrayInput is an input type that accepts SyncGithubAppsArray and SyncGithubAppsArrayOutput values.
// You can construct a concrete instance of `SyncGithubAppsArrayInput` via:
//
//	SyncGithubAppsArray{ SyncGithubAppsArgs{...} }
type SyncGithubAppsArrayInput interface {
	pulumi.Input

	ToSyncGithubAppsArrayOutput() SyncGithubAppsArrayOutput
	ToSyncGithubAppsArrayOutputWithContext(context.Context) SyncGithubAppsArrayOutput
}

type SyncGithubAppsArray []SyncGithubAppsInput

func (SyncGithubAppsArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SyncGithubApps)(nil)).Elem()
}

func (i SyncGithubAppsArray) ToSyncGithubAppsArrayOutput() SyncGithubAppsArrayOutput {
	return i.ToSyncGithubAppsArrayOutputWithContext(context.Background())
}

func (i SyncGithubAppsArray) ToSyncGithubAppsArrayOutputWithContext(ctx context.Context) SyncGithubAppsArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SyncGithubAppsArrayOutput)
}

// SyncGithubAppsMapInput is an input type that accepts SyncGithubAppsMap and SyncGithubAppsMapOutput values.
// You can construct a concrete instance of `SyncGithubAppsMapInput` via:
//
//	SyncGithubAppsMap{ "key": SyncGithubAppsArgs{...} }
type SyncGithubAppsMapInput interface {
	pulumi.Input

	ToSyncGithubAppsMapOutput() SyncGithubAppsMapOutput
	ToSyncGithubAppsMapOutputWithContext(context.Context) SyncGithubAppsMapOutput
}

type SyncGithubAppsMap map[string]SyncGithubAppsInput

func (SyncGithubAppsMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SyncGithubApps)(nil)).Elem()
}

func (i SyncGithubAppsMap) ToSyncGithubAppsMapOutput() SyncGithubAppsMapOutput {
	return i.ToSyncGithubAppsMapOutputWithContext(context.Background())
}

func (i SyncGithubAppsMap) ToSyncGithubAppsMapOutputWithContext(ctx context.Context) SyncGithubAppsMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SyncGithubAppsMapOutput)
}

type SyncGithubAppsOutput struct{ *pulumi.OutputState }

func (SyncGithubAppsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SyncGithubApps)(nil)).Elem()
}

func (o SyncGithubAppsOutput) ToSyncGithubAppsOutput() SyncGithubAppsOutput {
	return o
}

func (o SyncGithubAppsOutput) ToSyncGithubAppsOutputWithContext(ctx context.Context) SyncGithubAppsOutput {
	return o
}

// The GitHub application ID.
func (o SyncGithubAppsOutput) AppId() pulumi.IntOutput {
	return o.ApplyT(func(v *SyncGithubApps) pulumi.IntOutput { return v.AppId }).(pulumi.IntOutput)
}

// A fingerprint of a private key.
func (o SyncGithubAppsOutput) Fingerprint() pulumi.StringOutput {
	return o.ApplyT(func(v *SyncGithubApps) pulumi.StringOutput { return v.Fingerprint }).(pulumi.StringOutput)
}

// The user-defined name of the GitHub App configuration.
func (o SyncGithubAppsOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SyncGithubApps) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The namespace to provision the resource in.
// The value should not contain leading or trailing forward slashes.
// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
func (o SyncGithubAppsOutput) Namespace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SyncGithubApps) pulumi.StringPtrOutput { return v.Namespace }).(pulumi.StringPtrOutput)
}

// The content of a PEM formatted private key generated on GitHub for the app.
func (o SyncGithubAppsOutput) PrivateKey() pulumi.StringOutput {
	return o.ApplyT(func(v *SyncGithubApps) pulumi.StringOutput { return v.PrivateKey }).(pulumi.StringOutput)
}

type SyncGithubAppsArrayOutput struct{ *pulumi.OutputState }

func (SyncGithubAppsArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SyncGithubApps)(nil)).Elem()
}

func (o SyncGithubAppsArrayOutput) ToSyncGithubAppsArrayOutput() SyncGithubAppsArrayOutput {
	return o
}

func (o SyncGithubAppsArrayOutput) ToSyncGithubAppsArrayOutputWithContext(ctx context.Context) SyncGithubAppsArrayOutput {
	return o
}

func (o SyncGithubAppsArrayOutput) Index(i pulumi.IntInput) SyncGithubAppsOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SyncGithubApps {
		return vs[0].([]*SyncGithubApps)[vs[1].(int)]
	}).(SyncGithubAppsOutput)
}

type SyncGithubAppsMapOutput struct{ *pulumi.OutputState }

func (SyncGithubAppsMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SyncGithubApps)(nil)).Elem()
}

func (o SyncGithubAppsMapOutput) ToSyncGithubAppsMapOutput() SyncGithubAppsMapOutput {
	return o
}

func (o SyncGithubAppsMapOutput) ToSyncGithubAppsMapOutputWithContext(ctx context.Context) SyncGithubAppsMapOutput {
	return o
}

func (o SyncGithubAppsMapOutput) MapIndex(k pulumi.StringInput) SyncGithubAppsOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SyncGithubApps {
		return vs[0].(map[string]*SyncGithubApps)[vs[1].(string)]
	}).(SyncGithubAppsOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SyncGithubAppsInput)(nil)).Elem(), &SyncGithubApps{})
	pulumi.RegisterInputType(reflect.TypeOf((*SyncGithubAppsArrayInput)(nil)).Elem(), SyncGithubAppsArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SyncGithubAppsMapInput)(nil)).Elem(), SyncGithubAppsMap{})
	pulumi.RegisterOutputType(SyncGithubAppsOutput{})
	pulumi.RegisterOutputType(SyncGithubAppsArrayOutput{})
	pulumi.RegisterOutputType(SyncGithubAppsMapOutput{})
}
