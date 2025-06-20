// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package identity

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-vault/sdk/v7/go/vault/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

var _ = internal.GetEnvOrDefault

type GetEntityAliasType struct {
	// Canonical ID of the Alias
	CanonicalId string `pulumi:"canonicalId"`
	// Creation time of the Alias
	CreationTime string `pulumi:"creationTime"`
	// ID of the alias
	Id string `pulumi:"id"`
	// Last update time of the alias
	LastUpdateTime string `pulumi:"lastUpdateTime"`
	// List of canonical IDs merged with this alias
	MergedFromCanonicalIds []string `pulumi:"mergedFromCanonicalIds"`
	// Arbitrary metadata
	Metadata map[string]string `pulumi:"metadata"`
	// Authentication mount acccessor which this alias belongs to
	MountAccessor string `pulumi:"mountAccessor"`
	// Authentication mount path which this alias belongs to
	MountPath string `pulumi:"mountPath"`
	// Authentication mount type which this alias belongs to
	MountType string `pulumi:"mountType"`
	// Name of the alias
	Name string `pulumi:"name"`
}

// GetEntityAliasTypeInput is an input type that accepts GetEntityAliasTypeArgs and GetEntityAliasTypeOutput values.
// You can construct a concrete instance of `GetEntityAliasTypeInput` via:
//
//	GetEntityAliasTypeArgs{...}
type GetEntityAliasTypeInput interface {
	pulumi.Input

	ToGetEntityAliasTypeOutput() GetEntityAliasTypeOutput
	ToGetEntityAliasTypeOutputWithContext(context.Context) GetEntityAliasTypeOutput
}

type GetEntityAliasTypeArgs struct {
	// Canonical ID of the Alias
	CanonicalId pulumi.StringInput `pulumi:"canonicalId"`
	// Creation time of the Alias
	CreationTime pulumi.StringInput `pulumi:"creationTime"`
	// ID of the alias
	Id pulumi.StringInput `pulumi:"id"`
	// Last update time of the alias
	LastUpdateTime pulumi.StringInput `pulumi:"lastUpdateTime"`
	// List of canonical IDs merged with this alias
	MergedFromCanonicalIds pulumi.StringArrayInput `pulumi:"mergedFromCanonicalIds"`
	// Arbitrary metadata
	Metadata pulumi.StringMapInput `pulumi:"metadata"`
	// Authentication mount acccessor which this alias belongs to
	MountAccessor pulumi.StringInput `pulumi:"mountAccessor"`
	// Authentication mount path which this alias belongs to
	MountPath pulumi.StringInput `pulumi:"mountPath"`
	// Authentication mount type which this alias belongs to
	MountType pulumi.StringInput `pulumi:"mountType"`
	// Name of the alias
	Name pulumi.StringInput `pulumi:"name"`
}

func (GetEntityAliasTypeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetEntityAliasType)(nil)).Elem()
}

func (i GetEntityAliasTypeArgs) ToGetEntityAliasTypeOutput() GetEntityAliasTypeOutput {
	return i.ToGetEntityAliasTypeOutputWithContext(context.Background())
}

func (i GetEntityAliasTypeArgs) ToGetEntityAliasTypeOutputWithContext(ctx context.Context) GetEntityAliasTypeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetEntityAliasTypeOutput)
}

// GetEntityAliasTypeArrayInput is an input type that accepts GetEntityAliasTypeArray and GetEntityAliasTypeArrayOutput values.
// You can construct a concrete instance of `GetEntityAliasTypeArrayInput` via:
//
//	GetEntityAliasTypeArray{ GetEntityAliasTypeArgs{...} }
type GetEntityAliasTypeArrayInput interface {
	pulumi.Input

	ToGetEntityAliasTypeArrayOutput() GetEntityAliasTypeArrayOutput
	ToGetEntityAliasTypeArrayOutputWithContext(context.Context) GetEntityAliasTypeArrayOutput
}

type GetEntityAliasTypeArray []GetEntityAliasTypeInput

func (GetEntityAliasTypeArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetEntityAliasType)(nil)).Elem()
}

func (i GetEntityAliasTypeArray) ToGetEntityAliasTypeArrayOutput() GetEntityAliasTypeArrayOutput {
	return i.ToGetEntityAliasTypeArrayOutputWithContext(context.Background())
}

func (i GetEntityAliasTypeArray) ToGetEntityAliasTypeArrayOutputWithContext(ctx context.Context) GetEntityAliasTypeArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetEntityAliasTypeArrayOutput)
}

type GetEntityAliasTypeOutput struct{ *pulumi.OutputState }

func (GetEntityAliasTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetEntityAliasType)(nil)).Elem()
}

func (o GetEntityAliasTypeOutput) ToGetEntityAliasTypeOutput() GetEntityAliasTypeOutput {
	return o
}

func (o GetEntityAliasTypeOutput) ToGetEntityAliasTypeOutputWithContext(ctx context.Context) GetEntityAliasTypeOutput {
	return o
}

// Canonical ID of the Alias
func (o GetEntityAliasTypeOutput) CanonicalId() pulumi.StringOutput {
	return o.ApplyT(func(v GetEntityAliasType) string { return v.CanonicalId }).(pulumi.StringOutput)
}

// Creation time of the Alias
func (o GetEntityAliasTypeOutput) CreationTime() pulumi.StringOutput {
	return o.ApplyT(func(v GetEntityAliasType) string { return v.CreationTime }).(pulumi.StringOutput)
}

// ID of the alias
func (o GetEntityAliasTypeOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetEntityAliasType) string { return v.Id }).(pulumi.StringOutput)
}

// Last update time of the alias
func (o GetEntityAliasTypeOutput) LastUpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v GetEntityAliasType) string { return v.LastUpdateTime }).(pulumi.StringOutput)
}

// List of canonical IDs merged with this alias
func (o GetEntityAliasTypeOutput) MergedFromCanonicalIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetEntityAliasType) []string { return v.MergedFromCanonicalIds }).(pulumi.StringArrayOutput)
}

// Arbitrary metadata
func (o GetEntityAliasTypeOutput) Metadata() pulumi.StringMapOutput {
	return o.ApplyT(func(v GetEntityAliasType) map[string]string { return v.Metadata }).(pulumi.StringMapOutput)
}

// Authentication mount acccessor which this alias belongs to
func (o GetEntityAliasTypeOutput) MountAccessor() pulumi.StringOutput {
	return o.ApplyT(func(v GetEntityAliasType) string { return v.MountAccessor }).(pulumi.StringOutput)
}

// Authentication mount path which this alias belongs to
func (o GetEntityAliasTypeOutput) MountPath() pulumi.StringOutput {
	return o.ApplyT(func(v GetEntityAliasType) string { return v.MountPath }).(pulumi.StringOutput)
}

// Authentication mount type which this alias belongs to
func (o GetEntityAliasTypeOutput) MountType() pulumi.StringOutput {
	return o.ApplyT(func(v GetEntityAliasType) string { return v.MountType }).(pulumi.StringOutput)
}

// Name of the alias
func (o GetEntityAliasTypeOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GetEntityAliasType) string { return v.Name }).(pulumi.StringOutput)
}

type GetEntityAliasTypeArrayOutput struct{ *pulumi.OutputState }

func (GetEntityAliasTypeArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetEntityAliasType)(nil)).Elem()
}

func (o GetEntityAliasTypeArrayOutput) ToGetEntityAliasTypeArrayOutput() GetEntityAliasTypeArrayOutput {
	return o
}

func (o GetEntityAliasTypeArrayOutput) ToGetEntityAliasTypeArrayOutputWithContext(ctx context.Context) GetEntityAliasTypeArrayOutput {
	return o
}

func (o GetEntityAliasTypeArrayOutput) Index(i pulumi.IntInput) GetEntityAliasTypeOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) GetEntityAliasType {
		return vs[0].([]GetEntityAliasType)[vs[1].(int)]
	}).(GetEntityAliasTypeOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*GetEntityAliasTypeInput)(nil)).Elem(), GetEntityAliasTypeArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*GetEntityAliasTypeArrayInput)(nil)).Elem(), GetEntityAliasTypeArray{})
	pulumi.RegisterOutputType(GetEntityAliasTypeOutput{})
	pulumi.RegisterOutputType(GetEntityAliasTypeArrayOutput{})
}
