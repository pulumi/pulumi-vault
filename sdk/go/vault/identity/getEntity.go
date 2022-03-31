// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package identity

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-vault/sdk/v5/go/vault/identity"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := identity.LookupEntity(ctx, &identity.LookupEntityArgs{
// 			EntityName: pulumi.StringRef("entity_12345"),
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ## Required Vault Capabilities
//
// Use of this resource requires the `create` capability on `/identity/lookup/entity`.
func LookupEntity(ctx *pulumi.Context, args *LookupEntityArgs, opts ...pulumi.InvokeOption) (*LookupEntityResult, error) {
	var rv LookupEntityResult
	err := ctx.Invoke("vault:identity/getEntity:getEntity", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getEntity.
type LookupEntityArgs struct {
	// ID of the alias.
	AliasId *string `pulumi:"aliasId"`
	// Accessor of the mount to which the alias belongs to.
	// This should be supplied in conjunction with `aliasName`.
	AliasMountAccessor *string `pulumi:"aliasMountAccessor"`
	// Name of the alias. This should be supplied in conjunction with
	// `aliasMountAccessor`.
	AliasName *string `pulumi:"aliasName"`
	// ID of the entity.
	EntityId *string `pulumi:"entityId"`
	// Name of the entity.
	EntityName *string `pulumi:"entityName"`
}

// A collection of values returned by getEntity.
type LookupEntityResult struct {
	AliasId            string `pulumi:"aliasId"`
	AliasMountAccessor string `pulumi:"aliasMountAccessor"`
	AliasName          string `pulumi:"aliasName"`
	// A list of entity alias. Structure is documented below.
	Aliases []GetEntityAliasType `pulumi:"aliases"`
	// Creation time of the Alias
	CreationTime string `pulumi:"creationTime"`
	// A string containing the full data payload retrieved from
	// Vault, serialized in JSON format.
	DataJson string `pulumi:"dataJson"`
	// List of Group IDs of which the entity is directly a member of
	DirectGroupIds []string `pulumi:"directGroupIds"`
	// Whether the entity is disabled
	Disabled   bool   `pulumi:"disabled"`
	EntityId   string `pulumi:"entityId"`
	EntityName string `pulumi:"entityName"`
	// List of all Group IDs of which the entity is a member of
	GroupIds []string `pulumi:"groupIds"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// List of all Group IDs of which the entity is a member of transitively
	InheritedGroupIds []string `pulumi:"inheritedGroupIds"`
	// Last update time of the alias
	LastUpdateTime string `pulumi:"lastUpdateTime"`
	// Other entity IDs which is merged with this entity
	MergedEntityIds []string `pulumi:"mergedEntityIds"`
	// Arbitrary metadata
	Metadata map[string]interface{} `pulumi:"metadata"`
	// Namespace of which the entity is part of
	NamespaceId string `pulumi:"namespaceId"`
	// List of policies attached to the entity
	Policies []string `pulumi:"policies"`
}

func LookupEntityOutput(ctx *pulumi.Context, args LookupEntityOutputArgs, opts ...pulumi.InvokeOption) LookupEntityResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupEntityResult, error) {
			args := v.(LookupEntityArgs)
			r, err := LookupEntity(ctx, &args, opts...)
			return *r, err
		}).(LookupEntityResultOutput)
}

// A collection of arguments for invoking getEntity.
type LookupEntityOutputArgs struct {
	// ID of the alias.
	AliasId pulumi.StringPtrInput `pulumi:"aliasId"`
	// Accessor of the mount to which the alias belongs to.
	// This should be supplied in conjunction with `aliasName`.
	AliasMountAccessor pulumi.StringPtrInput `pulumi:"aliasMountAccessor"`
	// Name of the alias. This should be supplied in conjunction with
	// `aliasMountAccessor`.
	AliasName pulumi.StringPtrInput `pulumi:"aliasName"`
	// ID of the entity.
	EntityId pulumi.StringPtrInput `pulumi:"entityId"`
	// Name of the entity.
	EntityName pulumi.StringPtrInput `pulumi:"entityName"`
}

func (LookupEntityOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupEntityArgs)(nil)).Elem()
}

// A collection of values returned by getEntity.
type LookupEntityResultOutput struct{ *pulumi.OutputState }

func (LookupEntityResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupEntityResult)(nil)).Elem()
}

func (o LookupEntityResultOutput) ToLookupEntityResultOutput() LookupEntityResultOutput {
	return o
}

func (o LookupEntityResultOutput) ToLookupEntityResultOutputWithContext(ctx context.Context) LookupEntityResultOutput {
	return o
}

func (o LookupEntityResultOutput) AliasId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEntityResult) string { return v.AliasId }).(pulumi.StringOutput)
}

func (o LookupEntityResultOutput) AliasMountAccessor() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEntityResult) string { return v.AliasMountAccessor }).(pulumi.StringOutput)
}

func (o LookupEntityResultOutput) AliasName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEntityResult) string { return v.AliasName }).(pulumi.StringOutput)
}

// A list of entity alias. Structure is documented below.
func (o LookupEntityResultOutput) Aliases() GetEntityAliasTypeArrayOutput {
	return o.ApplyT(func(v LookupEntityResult) []GetEntityAliasType { return v.Aliases }).(GetEntityAliasTypeArrayOutput)
}

// Creation time of the Alias
func (o LookupEntityResultOutput) CreationTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEntityResult) string { return v.CreationTime }).(pulumi.StringOutput)
}

// A string containing the full data payload retrieved from
// Vault, serialized in JSON format.
func (o LookupEntityResultOutput) DataJson() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEntityResult) string { return v.DataJson }).(pulumi.StringOutput)
}

// List of Group IDs of which the entity is directly a member of
func (o LookupEntityResultOutput) DirectGroupIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupEntityResult) []string { return v.DirectGroupIds }).(pulumi.StringArrayOutput)
}

// Whether the entity is disabled
func (o LookupEntityResultOutput) Disabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupEntityResult) bool { return v.Disabled }).(pulumi.BoolOutput)
}

func (o LookupEntityResultOutput) EntityId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEntityResult) string { return v.EntityId }).(pulumi.StringOutput)
}

func (o LookupEntityResultOutput) EntityName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEntityResult) string { return v.EntityName }).(pulumi.StringOutput)
}

// List of all Group IDs of which the entity is a member of
func (o LookupEntityResultOutput) GroupIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupEntityResult) []string { return v.GroupIds }).(pulumi.StringArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupEntityResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEntityResult) string { return v.Id }).(pulumi.StringOutput)
}

// List of all Group IDs of which the entity is a member of transitively
func (o LookupEntityResultOutput) InheritedGroupIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupEntityResult) []string { return v.InheritedGroupIds }).(pulumi.StringArrayOutput)
}

// Last update time of the alias
func (o LookupEntityResultOutput) LastUpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEntityResult) string { return v.LastUpdateTime }).(pulumi.StringOutput)
}

// Other entity IDs which is merged with this entity
func (o LookupEntityResultOutput) MergedEntityIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupEntityResult) []string { return v.MergedEntityIds }).(pulumi.StringArrayOutput)
}

// Arbitrary metadata
func (o LookupEntityResultOutput) Metadata() pulumi.MapOutput {
	return o.ApplyT(func(v LookupEntityResult) map[string]interface{} { return v.Metadata }).(pulumi.MapOutput)
}

// Namespace of which the entity is part of
func (o LookupEntityResultOutput) NamespaceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEntityResult) string { return v.NamespaceId }).(pulumi.StringOutput)
}

// List of policies attached to the entity
func (o LookupEntityResultOutput) Policies() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupEntityResult) []string { return v.Policies }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupEntityResultOutput{})
}
