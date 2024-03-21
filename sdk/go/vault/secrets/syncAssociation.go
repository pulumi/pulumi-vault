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
//	"encoding/json"
//
//	"github.com/pulumi/pulumi-vault/sdk/v6/go/vault"
//	"github.com/pulumi/pulumi-vault/sdk/v6/go/vault/kv"
//	"github.com/pulumi/pulumi-vault/sdk/v6/go/vault/secrets"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			kvv2, err := vault.NewMount(ctx, "kvv2", &vault.MountArgs{
//				Path: pulumi.String("kvv2"),
//				Type: pulumi.String("kv"),
//				Options: pulumi.Map{
//					"version": pulumi.Any("2"),
//				},
//				Description: pulumi.String("KV Version 2 secret engine mount"),
//			})
//			if err != nil {
//				return err
//			}
//			tmpJSON0, err := json.Marshal(map[string]interface{}{
//				"dev":  "B!gS3cr3t",
//				"prod": "S3cureP4$$",
//			})
//			if err != nil {
//				return err
//			}
//			json0 := string(tmpJSON0)
//			token, err := kv.NewSecretV2(ctx, "token", &kv.SecretV2Args{
//				Mount:    kvv2.Path,
//				DataJson: pulumi.String(json0),
//			})
//			if err != nil {
//				return err
//			}
//			gh, err := secrets.NewSyncGhDestination(ctx, "gh", &secrets.SyncGhDestinationArgs{
//				AccessToken:        pulumi.Any(_var.Access_token),
//				RepositoryOwner:    pulumi.Any(_var.Repo_owner),
//				RepositoryName:     pulumi.String("repo-name-example"),
//				SecretNameTemplate: pulumi.String("vault_{{ .MountAccessor | lowercase }}_{{ .SecretPath | lowercase }}"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = secrets.NewSyncAssociation(ctx, "ghToken", &secrets.SyncAssociationArgs{
//				Type:       gh.Type,
//				Mount:      kvv2.Path,
//				SecretName: token.Name,
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
type SyncAssociation struct {
	pulumi.CustomResourceState

	// Specifies the mount where the secret is located.
	Mount pulumi.StringOutput `pulumi:"mount"`
	// Specifies the name of the destination.
	Name pulumi.StringOutput `pulumi:"name"`
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
	Namespace pulumi.StringPtrOutput `pulumi:"namespace"`
	// Specifies the name of the secret to synchronize.
	SecretName pulumi.StringOutput `pulumi:"secretName"`
	// Specifies the status of the association (for eg. `SYNCED`).
	SyncStatus pulumi.StringOutput `pulumi:"syncStatus"`
	// Specifies the destination type.
	Type pulumi.StringOutput `pulumi:"type"`
	// Duration string specifying when the secret was last updated.
	UpdatedAt pulumi.StringOutput `pulumi:"updatedAt"`
}

// NewSyncAssociation registers a new resource with the given unique name, arguments, and options.
func NewSyncAssociation(ctx *pulumi.Context,
	name string, args *SyncAssociationArgs, opts ...pulumi.ResourceOption) (*SyncAssociation, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Mount == nil {
		return nil, errors.New("invalid value for required argument 'Mount'")
	}
	if args.SecretName == nil {
		return nil, errors.New("invalid value for required argument 'SecretName'")
	}
	if args.Type == nil {
		return nil, errors.New("invalid value for required argument 'Type'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource SyncAssociation
	err := ctx.RegisterResource("vault:secrets/syncAssociation:SyncAssociation", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSyncAssociation gets an existing SyncAssociation resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSyncAssociation(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SyncAssociationState, opts ...pulumi.ResourceOption) (*SyncAssociation, error) {
	var resource SyncAssociation
	err := ctx.ReadResource("vault:secrets/syncAssociation:SyncAssociation", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SyncAssociation resources.
type syncAssociationState struct {
	// Specifies the mount where the secret is located.
	Mount *string `pulumi:"mount"`
	// Specifies the name of the destination.
	Name *string `pulumi:"name"`
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
	Namespace *string `pulumi:"namespace"`
	// Specifies the name of the secret to synchronize.
	SecretName *string `pulumi:"secretName"`
	// Specifies the status of the association (for eg. `SYNCED`).
	SyncStatus *string `pulumi:"syncStatus"`
	// Specifies the destination type.
	Type *string `pulumi:"type"`
	// Duration string specifying when the secret was last updated.
	UpdatedAt *string `pulumi:"updatedAt"`
}

type SyncAssociationState struct {
	// Specifies the mount where the secret is located.
	Mount pulumi.StringPtrInput
	// Specifies the name of the destination.
	Name pulumi.StringPtrInput
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
	Namespace pulumi.StringPtrInput
	// Specifies the name of the secret to synchronize.
	SecretName pulumi.StringPtrInput
	// Specifies the status of the association (for eg. `SYNCED`).
	SyncStatus pulumi.StringPtrInput
	// Specifies the destination type.
	Type pulumi.StringPtrInput
	// Duration string specifying when the secret was last updated.
	UpdatedAt pulumi.StringPtrInput
}

func (SyncAssociationState) ElementType() reflect.Type {
	return reflect.TypeOf((*syncAssociationState)(nil)).Elem()
}

type syncAssociationArgs struct {
	// Specifies the mount where the secret is located.
	Mount string `pulumi:"mount"`
	// Specifies the name of the destination.
	Name *string `pulumi:"name"`
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
	Namespace *string `pulumi:"namespace"`
	// Specifies the name of the secret to synchronize.
	SecretName string `pulumi:"secretName"`
	// Specifies the destination type.
	Type string `pulumi:"type"`
}

// The set of arguments for constructing a SyncAssociation resource.
type SyncAssociationArgs struct {
	// Specifies the mount where the secret is located.
	Mount pulumi.StringInput
	// Specifies the name of the destination.
	Name pulumi.StringPtrInput
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
	Namespace pulumi.StringPtrInput
	// Specifies the name of the secret to synchronize.
	SecretName pulumi.StringInput
	// Specifies the destination type.
	Type pulumi.StringInput
}

func (SyncAssociationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*syncAssociationArgs)(nil)).Elem()
}

type SyncAssociationInput interface {
	pulumi.Input

	ToSyncAssociationOutput() SyncAssociationOutput
	ToSyncAssociationOutputWithContext(ctx context.Context) SyncAssociationOutput
}

func (*SyncAssociation) ElementType() reflect.Type {
	return reflect.TypeOf((**SyncAssociation)(nil)).Elem()
}

func (i *SyncAssociation) ToSyncAssociationOutput() SyncAssociationOutput {
	return i.ToSyncAssociationOutputWithContext(context.Background())
}

func (i *SyncAssociation) ToSyncAssociationOutputWithContext(ctx context.Context) SyncAssociationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SyncAssociationOutput)
}

// SyncAssociationArrayInput is an input type that accepts SyncAssociationArray and SyncAssociationArrayOutput values.
// You can construct a concrete instance of `SyncAssociationArrayInput` via:
//
//	SyncAssociationArray{ SyncAssociationArgs{...} }
type SyncAssociationArrayInput interface {
	pulumi.Input

	ToSyncAssociationArrayOutput() SyncAssociationArrayOutput
	ToSyncAssociationArrayOutputWithContext(context.Context) SyncAssociationArrayOutput
}

type SyncAssociationArray []SyncAssociationInput

func (SyncAssociationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SyncAssociation)(nil)).Elem()
}

func (i SyncAssociationArray) ToSyncAssociationArrayOutput() SyncAssociationArrayOutput {
	return i.ToSyncAssociationArrayOutputWithContext(context.Background())
}

func (i SyncAssociationArray) ToSyncAssociationArrayOutputWithContext(ctx context.Context) SyncAssociationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SyncAssociationArrayOutput)
}

// SyncAssociationMapInput is an input type that accepts SyncAssociationMap and SyncAssociationMapOutput values.
// You can construct a concrete instance of `SyncAssociationMapInput` via:
//
//	SyncAssociationMap{ "key": SyncAssociationArgs{...} }
type SyncAssociationMapInput interface {
	pulumi.Input

	ToSyncAssociationMapOutput() SyncAssociationMapOutput
	ToSyncAssociationMapOutputWithContext(context.Context) SyncAssociationMapOutput
}

type SyncAssociationMap map[string]SyncAssociationInput

func (SyncAssociationMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SyncAssociation)(nil)).Elem()
}

func (i SyncAssociationMap) ToSyncAssociationMapOutput() SyncAssociationMapOutput {
	return i.ToSyncAssociationMapOutputWithContext(context.Background())
}

func (i SyncAssociationMap) ToSyncAssociationMapOutputWithContext(ctx context.Context) SyncAssociationMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SyncAssociationMapOutput)
}

type SyncAssociationOutput struct{ *pulumi.OutputState }

func (SyncAssociationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SyncAssociation)(nil)).Elem()
}

func (o SyncAssociationOutput) ToSyncAssociationOutput() SyncAssociationOutput {
	return o
}

func (o SyncAssociationOutput) ToSyncAssociationOutputWithContext(ctx context.Context) SyncAssociationOutput {
	return o
}

// Specifies the mount where the secret is located.
func (o SyncAssociationOutput) Mount() pulumi.StringOutput {
	return o.ApplyT(func(v *SyncAssociation) pulumi.StringOutput { return v.Mount }).(pulumi.StringOutput)
}

// Specifies the name of the destination.
func (o SyncAssociationOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SyncAssociation) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The namespace to provision the resource in.
// The value should not contain leading or trailing forward slashes.
// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
func (o SyncAssociationOutput) Namespace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SyncAssociation) pulumi.StringPtrOutput { return v.Namespace }).(pulumi.StringPtrOutput)
}

// Specifies the name of the secret to synchronize.
func (o SyncAssociationOutput) SecretName() pulumi.StringOutput {
	return o.ApplyT(func(v *SyncAssociation) pulumi.StringOutput { return v.SecretName }).(pulumi.StringOutput)
}

// Specifies the status of the association (for eg. `SYNCED`).
func (o SyncAssociationOutput) SyncStatus() pulumi.StringOutput {
	return o.ApplyT(func(v *SyncAssociation) pulumi.StringOutput { return v.SyncStatus }).(pulumi.StringOutput)
}

// Specifies the destination type.
func (o SyncAssociationOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *SyncAssociation) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// Duration string specifying when the secret was last updated.
func (o SyncAssociationOutput) UpdatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *SyncAssociation) pulumi.StringOutput { return v.UpdatedAt }).(pulumi.StringOutput)
}

type SyncAssociationArrayOutput struct{ *pulumi.OutputState }

func (SyncAssociationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SyncAssociation)(nil)).Elem()
}

func (o SyncAssociationArrayOutput) ToSyncAssociationArrayOutput() SyncAssociationArrayOutput {
	return o
}

func (o SyncAssociationArrayOutput) ToSyncAssociationArrayOutputWithContext(ctx context.Context) SyncAssociationArrayOutput {
	return o
}

func (o SyncAssociationArrayOutput) Index(i pulumi.IntInput) SyncAssociationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SyncAssociation {
		return vs[0].([]*SyncAssociation)[vs[1].(int)]
	}).(SyncAssociationOutput)
}

type SyncAssociationMapOutput struct{ *pulumi.OutputState }

func (SyncAssociationMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SyncAssociation)(nil)).Elem()
}

func (o SyncAssociationMapOutput) ToSyncAssociationMapOutput() SyncAssociationMapOutput {
	return o
}

func (o SyncAssociationMapOutput) ToSyncAssociationMapOutputWithContext(ctx context.Context) SyncAssociationMapOutput {
	return o
}

func (o SyncAssociationMapOutput) MapIndex(k pulumi.StringInput) SyncAssociationOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SyncAssociation {
		return vs[0].(map[string]*SyncAssociation)[vs[1].(string)]
	}).(SyncAssociationOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SyncAssociationInput)(nil)).Elem(), &SyncAssociation{})
	pulumi.RegisterInputType(reflect.TypeOf((*SyncAssociationArrayInput)(nil)).Elem(), SyncAssociationArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SyncAssociationMapInput)(nil)).Elem(), SyncAssociationMap{})
	pulumi.RegisterOutputType(SyncAssociationOutput{})
	pulumi.RegisterOutputType(SyncAssociationArrayOutput{})
	pulumi.RegisterOutputType(SyncAssociationMapOutput{})
}
