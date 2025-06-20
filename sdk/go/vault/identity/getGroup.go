// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package identity

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-vault/sdk/v7/go/vault/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-vault/sdk/v7/go/vault/identity"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := identity.LookupGroup(ctx, &identity.LookupGroupArgs{
//				GroupName: pulumi.StringRef("user"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// ## Required Vault Capabilities
//
// Use of this resource requires the `create` capability on `/identity/lookup/group`.
func LookupGroup(ctx *pulumi.Context, args *LookupGroupArgs, opts ...pulumi.InvokeOption) (*LookupGroupResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupGroupResult
	err := ctx.Invoke("vault:identity/getGroup:getGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getGroup.
type LookupGroupArgs struct {
	// ID of the alias.
	AliasId *string `pulumi:"aliasId"`
	// Accessor of the mount to which the alias belongs to.
	// This should be supplied in conjunction with `aliasName`.
	//
	// The lookup criteria can be `groupName`, `groupId`, `aliasId`, or a combination of
	// `aliasName` and `aliasMountAccessor`.
	AliasMountAccessor *string `pulumi:"aliasMountAccessor"`
	// Name of the alias. This should be supplied in conjunction with
	// `aliasMountAccessor`.
	AliasName *string `pulumi:"aliasName"`
	// ID of the group.
	GroupId *string `pulumi:"groupId"`
	// Name of the group.
	GroupName *string `pulumi:"groupName"`
	// The namespace of the target resource.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
	// *Available only for Vault Enterprise*.
	Namespace *string `pulumi:"namespace"`
}

// A collection of values returned by getGroup.
type LookupGroupResult struct {
	// Canonical ID of the Alias
	AliasCanonicalId string `pulumi:"aliasCanonicalId"`
	// Creation time of the Alias
	AliasCreationTime string `pulumi:"aliasCreationTime"`
	AliasId           string `pulumi:"aliasId"`
	// Last update time of the alias
	AliasLastUpdateTime string `pulumi:"aliasLastUpdateTime"`
	// List of canonical IDs merged with this alias
	AliasMergedFromCanonicalIds []string `pulumi:"aliasMergedFromCanonicalIds"`
	// Arbitrary metadata
	AliasMetadata      map[string]string `pulumi:"aliasMetadata"`
	AliasMountAccessor string            `pulumi:"aliasMountAccessor"`
	// Authentication mount path which this alias belongs to
	AliasMountPath string `pulumi:"aliasMountPath"`
	// Authentication mount type which this alias belongs to
	AliasMountType string `pulumi:"aliasMountType"`
	AliasName      string `pulumi:"aliasName"`
	// Creation timestamp of the group
	CreationTime string `pulumi:"creationTime"`
	// A string containing the full data payload retrieved from
	// Vault, serialized in JSON format.
	DataJson  string `pulumi:"dataJson"`
	GroupId   string `pulumi:"groupId"`
	GroupName string `pulumi:"groupName"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Last updated time of the group
	LastUpdateTime string `pulumi:"lastUpdateTime"`
	// List of Entity IDs which are members of this group
	MemberEntityIds []string `pulumi:"memberEntityIds"`
	// List of Group IDs which are members of this group
	MemberGroupIds []string `pulumi:"memberGroupIds"`
	// Arbitrary metadata
	Metadata map[string]string `pulumi:"metadata"`
	// Modify index of the group
	ModifyIndex int     `pulumi:"modifyIndex"`
	Namespace   *string `pulumi:"namespace"`
	// Namespace of which the group is part of
	NamespaceId string `pulumi:"namespaceId"`
	// List of Group IDs which are parents of this group.
	ParentGroupIds []string `pulumi:"parentGroupIds"`
	// List of policies attached to the group
	Policies []string `pulumi:"policies"`
	// Type of group
	Type string `pulumi:"type"`
}

func LookupGroupOutput(ctx *pulumi.Context, args LookupGroupOutputArgs, opts ...pulumi.InvokeOption) LookupGroupResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupGroupResultOutput, error) {
			args := v.(LookupGroupArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("vault:identity/getGroup:getGroup", args, LookupGroupResultOutput{}, options).(LookupGroupResultOutput), nil
		}).(LookupGroupResultOutput)
}

// A collection of arguments for invoking getGroup.
type LookupGroupOutputArgs struct {
	// ID of the alias.
	AliasId pulumi.StringPtrInput `pulumi:"aliasId"`
	// Accessor of the mount to which the alias belongs to.
	// This should be supplied in conjunction with `aliasName`.
	//
	// The lookup criteria can be `groupName`, `groupId`, `aliasId`, or a combination of
	// `aliasName` and `aliasMountAccessor`.
	AliasMountAccessor pulumi.StringPtrInput `pulumi:"aliasMountAccessor"`
	// Name of the alias. This should be supplied in conjunction with
	// `aliasMountAccessor`.
	AliasName pulumi.StringPtrInput `pulumi:"aliasName"`
	// ID of the group.
	GroupId pulumi.StringPtrInput `pulumi:"groupId"`
	// Name of the group.
	GroupName pulumi.StringPtrInput `pulumi:"groupName"`
	// The namespace of the target resource.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
	// *Available only for Vault Enterprise*.
	Namespace pulumi.StringPtrInput `pulumi:"namespace"`
}

func (LookupGroupOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupGroupArgs)(nil)).Elem()
}

// A collection of values returned by getGroup.
type LookupGroupResultOutput struct{ *pulumi.OutputState }

func (LookupGroupResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupGroupResult)(nil)).Elem()
}

func (o LookupGroupResultOutput) ToLookupGroupResultOutput() LookupGroupResultOutput {
	return o
}

func (o LookupGroupResultOutput) ToLookupGroupResultOutputWithContext(ctx context.Context) LookupGroupResultOutput {
	return o
}

// Canonical ID of the Alias
func (o LookupGroupResultOutput) AliasCanonicalId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGroupResult) string { return v.AliasCanonicalId }).(pulumi.StringOutput)
}

// Creation time of the Alias
func (o LookupGroupResultOutput) AliasCreationTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGroupResult) string { return v.AliasCreationTime }).(pulumi.StringOutput)
}

func (o LookupGroupResultOutput) AliasId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGroupResult) string { return v.AliasId }).(pulumi.StringOutput)
}

// Last update time of the alias
func (o LookupGroupResultOutput) AliasLastUpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGroupResult) string { return v.AliasLastUpdateTime }).(pulumi.StringOutput)
}

// List of canonical IDs merged with this alias
func (o LookupGroupResultOutput) AliasMergedFromCanonicalIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupGroupResult) []string { return v.AliasMergedFromCanonicalIds }).(pulumi.StringArrayOutput)
}

// Arbitrary metadata
func (o LookupGroupResultOutput) AliasMetadata() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupGroupResult) map[string]string { return v.AliasMetadata }).(pulumi.StringMapOutput)
}

func (o LookupGroupResultOutput) AliasMountAccessor() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGroupResult) string { return v.AliasMountAccessor }).(pulumi.StringOutput)
}

// Authentication mount path which this alias belongs to
func (o LookupGroupResultOutput) AliasMountPath() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGroupResult) string { return v.AliasMountPath }).(pulumi.StringOutput)
}

// Authentication mount type which this alias belongs to
func (o LookupGroupResultOutput) AliasMountType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGroupResult) string { return v.AliasMountType }).(pulumi.StringOutput)
}

func (o LookupGroupResultOutput) AliasName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGroupResult) string { return v.AliasName }).(pulumi.StringOutput)
}

// Creation timestamp of the group
func (o LookupGroupResultOutput) CreationTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGroupResult) string { return v.CreationTime }).(pulumi.StringOutput)
}

// A string containing the full data payload retrieved from
// Vault, serialized in JSON format.
func (o LookupGroupResultOutput) DataJson() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGroupResult) string { return v.DataJson }).(pulumi.StringOutput)
}

func (o LookupGroupResultOutput) GroupId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGroupResult) string { return v.GroupId }).(pulumi.StringOutput)
}

func (o LookupGroupResultOutput) GroupName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGroupResult) string { return v.GroupName }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupGroupResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGroupResult) string { return v.Id }).(pulumi.StringOutput)
}

// Last updated time of the group
func (o LookupGroupResultOutput) LastUpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGroupResult) string { return v.LastUpdateTime }).(pulumi.StringOutput)
}

// List of Entity IDs which are members of this group
func (o LookupGroupResultOutput) MemberEntityIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupGroupResult) []string { return v.MemberEntityIds }).(pulumi.StringArrayOutput)
}

// List of Group IDs which are members of this group
func (o LookupGroupResultOutput) MemberGroupIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupGroupResult) []string { return v.MemberGroupIds }).(pulumi.StringArrayOutput)
}

// Arbitrary metadata
func (o LookupGroupResultOutput) Metadata() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupGroupResult) map[string]string { return v.Metadata }).(pulumi.StringMapOutput)
}

// Modify index of the group
func (o LookupGroupResultOutput) ModifyIndex() pulumi.IntOutput {
	return o.ApplyT(func(v LookupGroupResult) int { return v.ModifyIndex }).(pulumi.IntOutput)
}

func (o LookupGroupResultOutput) Namespace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupGroupResult) *string { return v.Namespace }).(pulumi.StringPtrOutput)
}

// Namespace of which the group is part of
func (o LookupGroupResultOutput) NamespaceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGroupResult) string { return v.NamespaceId }).(pulumi.StringOutput)
}

// List of Group IDs which are parents of this group.
func (o LookupGroupResultOutput) ParentGroupIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupGroupResult) []string { return v.ParentGroupIds }).(pulumi.StringArrayOutput)
}

// List of policies attached to the group
func (o LookupGroupResultOutput) Policies() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupGroupResult) []string { return v.Policies }).(pulumi.StringArrayOutput)
}

// Type of group
func (o LookupGroupResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGroupResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupGroupResultOutput{})
}
