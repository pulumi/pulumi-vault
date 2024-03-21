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
// <!--Start PulumiCodeChooser -->
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-vault/sdk/v6/go/vault/secrets"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := secrets.NewSyncVercelDestination(ctx, "vercel", &secrets.SyncVercelDestinationArgs{
//				AccessToken: pulumi.Any(_var.Access_token),
//				ProjectId:   pulumi.Any(_var.Project_id),
//				DeploymentEnvironments: pulumi.StringArray{
//					pulumi.String("development"),
//					pulumi.String("preview"),
//					pulumi.String("production"),
//				},
//				SecretNameTemplate: pulumi.String("vault_{{ .MountAccessor | lowercase }}_{{ .SecretPath | lowercase }}"),
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
//
// ## Import
//
// GitHub Secrets sync destinations can be imported using the `name`, e.g.
//
// ```sh
// $ pulumi import vault:secrets/syncVercelDestination:SyncVercelDestination vercel vercel-dest
// ```
type SyncVercelDestination struct {
	pulumi.CustomResourceState

	// Vercel API access token with the permissions to manage environment
	// variables.
	AccessToken pulumi.StringOutput `pulumi:"accessToken"`
	// Deployment environments where the environment variables
	// are available. Accepts `development`, `preview` and `production`.
	DeploymentEnvironments pulumi.StringArrayOutput `pulumi:"deploymentEnvironments"`
	// Unique name of the GitHub destination.
	Name pulumi.StringOutput `pulumi:"name"`
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
	Namespace pulumi.StringPtrOutput `pulumi:"namespace"`
	// Project ID where to manage environment variables.
	ProjectId pulumi.StringOutput `pulumi:"projectId"`
	// Template describing how to generate external secret names.
	// Supports a subset of the Go Template syntax.
	SecretNameTemplate pulumi.StringOutput `pulumi:"secretNameTemplate"`
	// Team ID where to manage environment variables.
	TeamId pulumi.StringPtrOutput `pulumi:"teamId"`
	// The type of the secrets destination (`vercel-project`).
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewSyncVercelDestination registers a new resource with the given unique name, arguments, and options.
func NewSyncVercelDestination(ctx *pulumi.Context,
	name string, args *SyncVercelDestinationArgs, opts ...pulumi.ResourceOption) (*SyncVercelDestination, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AccessToken == nil {
		return nil, errors.New("invalid value for required argument 'AccessToken'")
	}
	if args.DeploymentEnvironments == nil {
		return nil, errors.New("invalid value for required argument 'DeploymentEnvironments'")
	}
	if args.ProjectId == nil {
		return nil, errors.New("invalid value for required argument 'ProjectId'")
	}
	if args.AccessToken != nil {
		args.AccessToken = pulumi.ToSecret(args.AccessToken).(pulumi.StringInput)
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"accessToken",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource SyncVercelDestination
	err := ctx.RegisterResource("vault:secrets/syncVercelDestination:SyncVercelDestination", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSyncVercelDestination gets an existing SyncVercelDestination resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSyncVercelDestination(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SyncVercelDestinationState, opts ...pulumi.ResourceOption) (*SyncVercelDestination, error) {
	var resource SyncVercelDestination
	err := ctx.ReadResource("vault:secrets/syncVercelDestination:SyncVercelDestination", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SyncVercelDestination resources.
type syncVercelDestinationState struct {
	// Vercel API access token with the permissions to manage environment
	// variables.
	AccessToken *string `pulumi:"accessToken"`
	// Deployment environments where the environment variables
	// are available. Accepts `development`, `preview` and `production`.
	DeploymentEnvironments []string `pulumi:"deploymentEnvironments"`
	// Unique name of the GitHub destination.
	Name *string `pulumi:"name"`
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
	Namespace *string `pulumi:"namespace"`
	// Project ID where to manage environment variables.
	ProjectId *string `pulumi:"projectId"`
	// Template describing how to generate external secret names.
	// Supports a subset of the Go Template syntax.
	SecretNameTemplate *string `pulumi:"secretNameTemplate"`
	// Team ID where to manage environment variables.
	TeamId *string `pulumi:"teamId"`
	// The type of the secrets destination (`vercel-project`).
	Type *string `pulumi:"type"`
}

type SyncVercelDestinationState struct {
	// Vercel API access token with the permissions to manage environment
	// variables.
	AccessToken pulumi.StringPtrInput
	// Deployment environments where the environment variables
	// are available. Accepts `development`, `preview` and `production`.
	DeploymentEnvironments pulumi.StringArrayInput
	// Unique name of the GitHub destination.
	Name pulumi.StringPtrInput
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
	Namespace pulumi.StringPtrInput
	// Project ID where to manage environment variables.
	ProjectId pulumi.StringPtrInput
	// Template describing how to generate external secret names.
	// Supports a subset of the Go Template syntax.
	SecretNameTemplate pulumi.StringPtrInput
	// Team ID where to manage environment variables.
	TeamId pulumi.StringPtrInput
	// The type of the secrets destination (`vercel-project`).
	Type pulumi.StringPtrInput
}

func (SyncVercelDestinationState) ElementType() reflect.Type {
	return reflect.TypeOf((*syncVercelDestinationState)(nil)).Elem()
}

type syncVercelDestinationArgs struct {
	// Vercel API access token with the permissions to manage environment
	// variables.
	AccessToken string `pulumi:"accessToken"`
	// Deployment environments where the environment variables
	// are available. Accepts `development`, `preview` and `production`.
	DeploymentEnvironments []string `pulumi:"deploymentEnvironments"`
	// Unique name of the GitHub destination.
	Name *string `pulumi:"name"`
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
	Namespace *string `pulumi:"namespace"`
	// Project ID where to manage environment variables.
	ProjectId string `pulumi:"projectId"`
	// Template describing how to generate external secret names.
	// Supports a subset of the Go Template syntax.
	SecretNameTemplate *string `pulumi:"secretNameTemplate"`
	// Team ID where to manage environment variables.
	TeamId *string `pulumi:"teamId"`
}

// The set of arguments for constructing a SyncVercelDestination resource.
type SyncVercelDestinationArgs struct {
	// Vercel API access token with the permissions to manage environment
	// variables.
	AccessToken pulumi.StringInput
	// Deployment environments where the environment variables
	// are available. Accepts `development`, `preview` and `production`.
	DeploymentEnvironments pulumi.StringArrayInput
	// Unique name of the GitHub destination.
	Name pulumi.StringPtrInput
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
	Namespace pulumi.StringPtrInput
	// Project ID where to manage environment variables.
	ProjectId pulumi.StringInput
	// Template describing how to generate external secret names.
	// Supports a subset of the Go Template syntax.
	SecretNameTemplate pulumi.StringPtrInput
	// Team ID where to manage environment variables.
	TeamId pulumi.StringPtrInput
}

func (SyncVercelDestinationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*syncVercelDestinationArgs)(nil)).Elem()
}

type SyncVercelDestinationInput interface {
	pulumi.Input

	ToSyncVercelDestinationOutput() SyncVercelDestinationOutput
	ToSyncVercelDestinationOutputWithContext(ctx context.Context) SyncVercelDestinationOutput
}

func (*SyncVercelDestination) ElementType() reflect.Type {
	return reflect.TypeOf((**SyncVercelDestination)(nil)).Elem()
}

func (i *SyncVercelDestination) ToSyncVercelDestinationOutput() SyncVercelDestinationOutput {
	return i.ToSyncVercelDestinationOutputWithContext(context.Background())
}

func (i *SyncVercelDestination) ToSyncVercelDestinationOutputWithContext(ctx context.Context) SyncVercelDestinationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SyncVercelDestinationOutput)
}

// SyncVercelDestinationArrayInput is an input type that accepts SyncVercelDestinationArray and SyncVercelDestinationArrayOutput values.
// You can construct a concrete instance of `SyncVercelDestinationArrayInput` via:
//
//	SyncVercelDestinationArray{ SyncVercelDestinationArgs{...} }
type SyncVercelDestinationArrayInput interface {
	pulumi.Input

	ToSyncVercelDestinationArrayOutput() SyncVercelDestinationArrayOutput
	ToSyncVercelDestinationArrayOutputWithContext(context.Context) SyncVercelDestinationArrayOutput
}

type SyncVercelDestinationArray []SyncVercelDestinationInput

func (SyncVercelDestinationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SyncVercelDestination)(nil)).Elem()
}

func (i SyncVercelDestinationArray) ToSyncVercelDestinationArrayOutput() SyncVercelDestinationArrayOutput {
	return i.ToSyncVercelDestinationArrayOutputWithContext(context.Background())
}

func (i SyncVercelDestinationArray) ToSyncVercelDestinationArrayOutputWithContext(ctx context.Context) SyncVercelDestinationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SyncVercelDestinationArrayOutput)
}

// SyncVercelDestinationMapInput is an input type that accepts SyncVercelDestinationMap and SyncVercelDestinationMapOutput values.
// You can construct a concrete instance of `SyncVercelDestinationMapInput` via:
//
//	SyncVercelDestinationMap{ "key": SyncVercelDestinationArgs{...} }
type SyncVercelDestinationMapInput interface {
	pulumi.Input

	ToSyncVercelDestinationMapOutput() SyncVercelDestinationMapOutput
	ToSyncVercelDestinationMapOutputWithContext(context.Context) SyncVercelDestinationMapOutput
}

type SyncVercelDestinationMap map[string]SyncVercelDestinationInput

func (SyncVercelDestinationMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SyncVercelDestination)(nil)).Elem()
}

func (i SyncVercelDestinationMap) ToSyncVercelDestinationMapOutput() SyncVercelDestinationMapOutput {
	return i.ToSyncVercelDestinationMapOutputWithContext(context.Background())
}

func (i SyncVercelDestinationMap) ToSyncVercelDestinationMapOutputWithContext(ctx context.Context) SyncVercelDestinationMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SyncVercelDestinationMapOutput)
}

type SyncVercelDestinationOutput struct{ *pulumi.OutputState }

func (SyncVercelDestinationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SyncVercelDestination)(nil)).Elem()
}

func (o SyncVercelDestinationOutput) ToSyncVercelDestinationOutput() SyncVercelDestinationOutput {
	return o
}

func (o SyncVercelDestinationOutput) ToSyncVercelDestinationOutputWithContext(ctx context.Context) SyncVercelDestinationOutput {
	return o
}

// Vercel API access token with the permissions to manage environment
// variables.
func (o SyncVercelDestinationOutput) AccessToken() pulumi.StringOutput {
	return o.ApplyT(func(v *SyncVercelDestination) pulumi.StringOutput { return v.AccessToken }).(pulumi.StringOutput)
}

// Deployment environments where the environment variables
// are available. Accepts `development`, `preview` and `production`.
func (o SyncVercelDestinationOutput) DeploymentEnvironments() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *SyncVercelDestination) pulumi.StringArrayOutput { return v.DeploymentEnvironments }).(pulumi.StringArrayOutput)
}

// Unique name of the GitHub destination.
func (o SyncVercelDestinationOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SyncVercelDestination) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The namespace to provision the resource in.
// The value should not contain leading or trailing forward slashes.
// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
func (o SyncVercelDestinationOutput) Namespace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SyncVercelDestination) pulumi.StringPtrOutput { return v.Namespace }).(pulumi.StringPtrOutput)
}

// Project ID where to manage environment variables.
func (o SyncVercelDestinationOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v *SyncVercelDestination) pulumi.StringOutput { return v.ProjectId }).(pulumi.StringOutput)
}

// Template describing how to generate external secret names.
// Supports a subset of the Go Template syntax.
func (o SyncVercelDestinationOutput) SecretNameTemplate() pulumi.StringOutput {
	return o.ApplyT(func(v *SyncVercelDestination) pulumi.StringOutput { return v.SecretNameTemplate }).(pulumi.StringOutput)
}

// Team ID where to manage environment variables.
func (o SyncVercelDestinationOutput) TeamId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SyncVercelDestination) pulumi.StringPtrOutput { return v.TeamId }).(pulumi.StringPtrOutput)
}

// The type of the secrets destination (`vercel-project`).
func (o SyncVercelDestinationOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *SyncVercelDestination) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

type SyncVercelDestinationArrayOutput struct{ *pulumi.OutputState }

func (SyncVercelDestinationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SyncVercelDestination)(nil)).Elem()
}

func (o SyncVercelDestinationArrayOutput) ToSyncVercelDestinationArrayOutput() SyncVercelDestinationArrayOutput {
	return o
}

func (o SyncVercelDestinationArrayOutput) ToSyncVercelDestinationArrayOutputWithContext(ctx context.Context) SyncVercelDestinationArrayOutput {
	return o
}

func (o SyncVercelDestinationArrayOutput) Index(i pulumi.IntInput) SyncVercelDestinationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SyncVercelDestination {
		return vs[0].([]*SyncVercelDestination)[vs[1].(int)]
	}).(SyncVercelDestinationOutput)
}

type SyncVercelDestinationMapOutput struct{ *pulumi.OutputState }

func (SyncVercelDestinationMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SyncVercelDestination)(nil)).Elem()
}

func (o SyncVercelDestinationMapOutput) ToSyncVercelDestinationMapOutput() SyncVercelDestinationMapOutput {
	return o
}

func (o SyncVercelDestinationMapOutput) ToSyncVercelDestinationMapOutputWithContext(ctx context.Context) SyncVercelDestinationMapOutput {
	return o
}

func (o SyncVercelDestinationMapOutput) MapIndex(k pulumi.StringInput) SyncVercelDestinationOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SyncVercelDestination {
		return vs[0].(map[string]*SyncVercelDestination)[vs[1].(string)]
	}).(SyncVercelDestinationOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SyncVercelDestinationInput)(nil)).Elem(), &SyncVercelDestination{})
	pulumi.RegisterInputType(reflect.TypeOf((*SyncVercelDestinationArrayInput)(nil)).Elem(), SyncVercelDestinationArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SyncVercelDestinationMapInput)(nil)).Elem(), SyncVercelDestinationMap{})
	pulumi.RegisterOutputType(SyncVercelDestinationOutput{})
	pulumi.RegisterOutputType(SyncVercelDestinationArrayOutput{})
	pulumi.RegisterOutputType(SyncVercelDestinationMapOutput{})
}
