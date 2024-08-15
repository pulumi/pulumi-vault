// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vault

import (
	"context"
	"reflect"

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
//	"github.com/pulumi/pulumi-vault/sdk/v6/go/vault"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			main, err := vault.GetRaftAutopilotState(ctx, nil, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("failure-tolerance", main.FailureTolerance)
//			return nil
//		})
//	}
//
// ```
func GetRaftAutopilotState(ctx *pulumi.Context, args *GetRaftAutopilotStateArgs, opts ...pulumi.InvokeOption) (*GetRaftAutopilotStateResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetRaftAutopilotStateResult
	err := ctx.Invoke("vault:index/getRaftAutopilotState:getRaftAutopilotState", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getRaftAutopilotState.
type GetRaftAutopilotStateArgs struct {
	// The namespace of the target resource.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
	// *Available only for Vault Enterprise*.
	Namespace *string `pulumi:"namespace"`
}

// A collection of values returned by getRaftAutopilotState.
type GetRaftAutopilotStateResult struct {
	// How many nodes could fail before the cluster becomes unhealthy.
	FailureTolerance int `pulumi:"failureTolerance"`
	// Cluster health status.
	Healthy bool `pulumi:"healthy"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The current leader of Vault.
	Leader    string  `pulumi:"leader"`
	Namespace *string `pulumi:"namespace"`
	// The cluster-level optimistic failure tolerance.
	OptimisticFailureTolerance int `pulumi:"optimisticFailureTolerance"`
	// Additional output related to redundancy zones stored as a serialized map of strings.
	RedundancyZones map[string]string `pulumi:"redundancyZones"`
	// Additional output related to redundancy zones.
	RedundancyZonesJson string `pulumi:"redundancyZonesJson"`
	// Additionaly output related to servers in the cluster stored as a serialized map of strings.
	Servers map[string]string `pulumi:"servers"`
	// Additionaly output related to servers in the cluster.
	ServersJson string `pulumi:"serversJson"`
	// Additional output related to upgrade information stored as a serialized map of strings.
	UpgradeInfo map[string]string `pulumi:"upgradeInfo"`
	// Additional output related to upgrade information.
	UpgradeInfoJson string `pulumi:"upgradeInfoJson"`
	// The voters in the Vault cluster.
	Voters []string `pulumi:"voters"`
}

func GetRaftAutopilotStateOutput(ctx *pulumi.Context, args GetRaftAutopilotStateOutputArgs, opts ...pulumi.InvokeOption) GetRaftAutopilotStateResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetRaftAutopilotStateResult, error) {
			args := v.(GetRaftAutopilotStateArgs)
			r, err := GetRaftAutopilotState(ctx, &args, opts...)
			var s GetRaftAutopilotStateResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetRaftAutopilotStateResultOutput)
}

// A collection of arguments for invoking getRaftAutopilotState.
type GetRaftAutopilotStateOutputArgs struct {
	// The namespace of the target resource.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
	// *Available only for Vault Enterprise*.
	Namespace pulumi.StringPtrInput `pulumi:"namespace"`
}

func (GetRaftAutopilotStateOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetRaftAutopilotStateArgs)(nil)).Elem()
}

// A collection of values returned by getRaftAutopilotState.
type GetRaftAutopilotStateResultOutput struct{ *pulumi.OutputState }

func (GetRaftAutopilotStateResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetRaftAutopilotStateResult)(nil)).Elem()
}

func (o GetRaftAutopilotStateResultOutput) ToGetRaftAutopilotStateResultOutput() GetRaftAutopilotStateResultOutput {
	return o
}

func (o GetRaftAutopilotStateResultOutput) ToGetRaftAutopilotStateResultOutputWithContext(ctx context.Context) GetRaftAutopilotStateResultOutput {
	return o
}

// How many nodes could fail before the cluster becomes unhealthy.
func (o GetRaftAutopilotStateResultOutput) FailureTolerance() pulumi.IntOutput {
	return o.ApplyT(func(v GetRaftAutopilotStateResult) int { return v.FailureTolerance }).(pulumi.IntOutput)
}

// Cluster health status.
func (o GetRaftAutopilotStateResultOutput) Healthy() pulumi.BoolOutput {
	return o.ApplyT(func(v GetRaftAutopilotStateResult) bool { return v.Healthy }).(pulumi.BoolOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetRaftAutopilotStateResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetRaftAutopilotStateResult) string { return v.Id }).(pulumi.StringOutput)
}

// The current leader of Vault.
func (o GetRaftAutopilotStateResultOutput) Leader() pulumi.StringOutput {
	return o.ApplyT(func(v GetRaftAutopilotStateResult) string { return v.Leader }).(pulumi.StringOutput)
}

func (o GetRaftAutopilotStateResultOutput) Namespace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetRaftAutopilotStateResult) *string { return v.Namespace }).(pulumi.StringPtrOutput)
}

// The cluster-level optimistic failure tolerance.
func (o GetRaftAutopilotStateResultOutput) OptimisticFailureTolerance() pulumi.IntOutput {
	return o.ApplyT(func(v GetRaftAutopilotStateResult) int { return v.OptimisticFailureTolerance }).(pulumi.IntOutput)
}

// Additional output related to redundancy zones stored as a serialized map of strings.
func (o GetRaftAutopilotStateResultOutput) RedundancyZones() pulumi.StringMapOutput {
	return o.ApplyT(func(v GetRaftAutopilotStateResult) map[string]string { return v.RedundancyZones }).(pulumi.StringMapOutput)
}

// Additional output related to redundancy zones.
func (o GetRaftAutopilotStateResultOutput) RedundancyZonesJson() pulumi.StringOutput {
	return o.ApplyT(func(v GetRaftAutopilotStateResult) string { return v.RedundancyZonesJson }).(pulumi.StringOutput)
}

// Additionaly output related to servers in the cluster stored as a serialized map of strings.
func (o GetRaftAutopilotStateResultOutput) Servers() pulumi.StringMapOutput {
	return o.ApplyT(func(v GetRaftAutopilotStateResult) map[string]string { return v.Servers }).(pulumi.StringMapOutput)
}

// Additionaly output related to servers in the cluster.
func (o GetRaftAutopilotStateResultOutput) ServersJson() pulumi.StringOutput {
	return o.ApplyT(func(v GetRaftAutopilotStateResult) string { return v.ServersJson }).(pulumi.StringOutput)
}

// Additional output related to upgrade information stored as a serialized map of strings.
func (o GetRaftAutopilotStateResultOutput) UpgradeInfo() pulumi.StringMapOutput {
	return o.ApplyT(func(v GetRaftAutopilotStateResult) map[string]string { return v.UpgradeInfo }).(pulumi.StringMapOutput)
}

// Additional output related to upgrade information.
func (o GetRaftAutopilotStateResultOutput) UpgradeInfoJson() pulumi.StringOutput {
	return o.ApplyT(func(v GetRaftAutopilotStateResult) string { return v.UpgradeInfoJson }).(pulumi.StringOutput)
}

// The voters in the Vault cluster.
func (o GetRaftAutopilotStateResultOutput) Voters() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetRaftAutopilotStateResult) []string { return v.Voters }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetRaftAutopilotStateResultOutput{})
}
