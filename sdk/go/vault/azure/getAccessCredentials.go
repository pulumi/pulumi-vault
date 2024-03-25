// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package azure

import (
	"context"
	"reflect"

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
//	"github.com/pulumi/pulumi-vault/sdk/v6/go/vault/azure"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := azure.GetAccessCredentials(ctx, &azure.GetAccessCredentialsArgs{
//				Role:                     "my-role",
//				ValidateCreds:            pulumi.BoolRef(true),
//				NumSequentialSuccesses:   pulumi.IntRef(8),
//				NumSecondsBetweenTests:   pulumi.IntRef(1),
//				MaxCredValidationSeconds: pulumi.IntRef(300),
//			}, nil)
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
// ## Caveats
//
// The `validateCreds` option requires read-access to the `backend` config endpoint.
// If the effective Vault role does not have the required permissions then valid values
// are required to be set for: `subscriptionId`, `tenantId`, `environment`.
func GetAccessCredentials(ctx *pulumi.Context, args *GetAccessCredentialsArgs, opts ...pulumi.InvokeOption) (*GetAccessCredentialsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetAccessCredentialsResult
	err := ctx.Invoke("vault:azure/getAccessCredentials:getAccessCredentials", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getAccessCredentials.
type GetAccessCredentialsArgs struct {
	// The path to the Azure secret backend to
	// read credentials from, with no leading or trailing `/`s.
	Backend string `pulumi:"backend"`
	// The Azure environment to use during credential validation.
	// Defaults to the environment configured in the Vault backend.
	// Some possible values: `AzurePublicCloud`, `AzureGovernmentCloud`
	// *See the caveats section for more information on this field.*
	Environment *string `pulumi:"environment"`
	// If 'validate_creds' is true,
	// the number of seconds after which to give up validating credentials. Defaults
	// to 300.
	MaxCredValidationSeconds *int `pulumi:"maxCredValidationSeconds"`
	// The namespace of the target resource.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
	// *Available only for Vault Enterprise*.
	Namespace *string `pulumi:"namespace"`
	// If 'validate_creds' is true,
	// the number of seconds to wait between each test of generated credentials.
	// Defaults to 1.
	NumSecondsBetweenTests *int `pulumi:"numSecondsBetweenTests"`
	// If 'validate_creds' is true,
	// the number of sequential successes required to validate generated
	// credentials. Defaults to 8.
	NumSequentialSuccesses *int `pulumi:"numSequentialSuccesses"`
	// The name of the Azure secret backend role to read
	// credentials from, with no leading or trailing `/`s.
	Role string `pulumi:"role"`
	// The subscription ID to use during credential
	// validation. Defaults to the subscription ID configured in the Vault `backend`.
	// *See the caveats section for more information on this field.*
	SubscriptionId *string `pulumi:"subscriptionId"`
	// The tenant ID to use during credential validation.
	// Defaults to the tenant ID configured in the Vault `backend`.
	// *See the caveats section for more information on this field.*
	TenantId *string `pulumi:"tenantId"`
	// Whether generated credentials should be
	// validated before being returned. Defaults to `false`, which returns
	// credentials without checking whether they have fully propagated throughout
	// Azure Active Directory. Designating `true` activates testing.
	ValidateCreds *bool `pulumi:"validateCreds"`
}

// A collection of values returned by getAccessCredentials.
type GetAccessCredentialsResult struct {
	Backend string `pulumi:"backend"`
	// The client id for credentials to query the Azure APIs.
	ClientId string `pulumi:"clientId"`
	// The client secret for credentials to query the Azure APIs.
	ClientSecret string  `pulumi:"clientSecret"`
	Environment  *string `pulumi:"environment"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The duration of the secret lease, in seconds relative
	// to the time the data was requested. Once this time has passed any plan
	// generated with this data may fail to apply.
	LeaseDuration int `pulumi:"leaseDuration"`
	// The lease identifier assigned by Vault.
	LeaseId                  string  `pulumi:"leaseId"`
	LeaseRenewable           bool    `pulumi:"leaseRenewable"`
	LeaseStartTime           string  `pulumi:"leaseStartTime"`
	MaxCredValidationSeconds *int    `pulumi:"maxCredValidationSeconds"`
	Namespace                *string `pulumi:"namespace"`
	NumSecondsBetweenTests   *int    `pulumi:"numSecondsBetweenTests"`
	NumSequentialSuccesses   *int    `pulumi:"numSequentialSuccesses"`
	Role                     string  `pulumi:"role"`
	SubscriptionId           *string `pulumi:"subscriptionId"`
	TenantId                 *string `pulumi:"tenantId"`
	ValidateCreds            *bool   `pulumi:"validateCreds"`
}

func GetAccessCredentialsOutput(ctx *pulumi.Context, args GetAccessCredentialsOutputArgs, opts ...pulumi.InvokeOption) GetAccessCredentialsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetAccessCredentialsResult, error) {
			args := v.(GetAccessCredentialsArgs)
			r, err := GetAccessCredentials(ctx, &args, opts...)
			var s GetAccessCredentialsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetAccessCredentialsResultOutput)
}

// A collection of arguments for invoking getAccessCredentials.
type GetAccessCredentialsOutputArgs struct {
	// The path to the Azure secret backend to
	// read credentials from, with no leading or trailing `/`s.
	Backend pulumi.StringInput `pulumi:"backend"`
	// The Azure environment to use during credential validation.
	// Defaults to the environment configured in the Vault backend.
	// Some possible values: `AzurePublicCloud`, `AzureGovernmentCloud`
	// *See the caveats section for more information on this field.*
	Environment pulumi.StringPtrInput `pulumi:"environment"`
	// If 'validate_creds' is true,
	// the number of seconds after which to give up validating credentials. Defaults
	// to 300.
	MaxCredValidationSeconds pulumi.IntPtrInput `pulumi:"maxCredValidationSeconds"`
	// The namespace of the target resource.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
	// *Available only for Vault Enterprise*.
	Namespace pulumi.StringPtrInput `pulumi:"namespace"`
	// If 'validate_creds' is true,
	// the number of seconds to wait between each test of generated credentials.
	// Defaults to 1.
	NumSecondsBetweenTests pulumi.IntPtrInput `pulumi:"numSecondsBetweenTests"`
	// If 'validate_creds' is true,
	// the number of sequential successes required to validate generated
	// credentials. Defaults to 8.
	NumSequentialSuccesses pulumi.IntPtrInput `pulumi:"numSequentialSuccesses"`
	// The name of the Azure secret backend role to read
	// credentials from, with no leading or trailing `/`s.
	Role pulumi.StringInput `pulumi:"role"`
	// The subscription ID to use during credential
	// validation. Defaults to the subscription ID configured in the Vault `backend`.
	// *See the caveats section for more information on this field.*
	SubscriptionId pulumi.StringPtrInput `pulumi:"subscriptionId"`
	// The tenant ID to use during credential validation.
	// Defaults to the tenant ID configured in the Vault `backend`.
	// *See the caveats section for more information on this field.*
	TenantId pulumi.StringPtrInput `pulumi:"tenantId"`
	// Whether generated credentials should be
	// validated before being returned. Defaults to `false`, which returns
	// credentials without checking whether they have fully propagated throughout
	// Azure Active Directory. Designating `true` activates testing.
	ValidateCreds pulumi.BoolPtrInput `pulumi:"validateCreds"`
}

func (GetAccessCredentialsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetAccessCredentialsArgs)(nil)).Elem()
}

// A collection of values returned by getAccessCredentials.
type GetAccessCredentialsResultOutput struct{ *pulumi.OutputState }

func (GetAccessCredentialsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetAccessCredentialsResult)(nil)).Elem()
}

func (o GetAccessCredentialsResultOutput) ToGetAccessCredentialsResultOutput() GetAccessCredentialsResultOutput {
	return o
}

func (o GetAccessCredentialsResultOutput) ToGetAccessCredentialsResultOutputWithContext(ctx context.Context) GetAccessCredentialsResultOutput {
	return o
}

func (o GetAccessCredentialsResultOutput) Backend() pulumi.StringOutput {
	return o.ApplyT(func(v GetAccessCredentialsResult) string { return v.Backend }).(pulumi.StringOutput)
}

// The client id for credentials to query the Azure APIs.
func (o GetAccessCredentialsResultOutput) ClientId() pulumi.StringOutput {
	return o.ApplyT(func(v GetAccessCredentialsResult) string { return v.ClientId }).(pulumi.StringOutput)
}

// The client secret for credentials to query the Azure APIs.
func (o GetAccessCredentialsResultOutput) ClientSecret() pulumi.StringOutput {
	return o.ApplyT(func(v GetAccessCredentialsResult) string { return v.ClientSecret }).(pulumi.StringOutput)
}

func (o GetAccessCredentialsResultOutput) Environment() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetAccessCredentialsResult) *string { return v.Environment }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetAccessCredentialsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetAccessCredentialsResult) string { return v.Id }).(pulumi.StringOutput)
}

// The duration of the secret lease, in seconds relative
// to the time the data was requested. Once this time has passed any plan
// generated with this data may fail to apply.
func (o GetAccessCredentialsResultOutput) LeaseDuration() pulumi.IntOutput {
	return o.ApplyT(func(v GetAccessCredentialsResult) int { return v.LeaseDuration }).(pulumi.IntOutput)
}

// The lease identifier assigned by Vault.
func (o GetAccessCredentialsResultOutput) LeaseId() pulumi.StringOutput {
	return o.ApplyT(func(v GetAccessCredentialsResult) string { return v.LeaseId }).(pulumi.StringOutput)
}

func (o GetAccessCredentialsResultOutput) LeaseRenewable() pulumi.BoolOutput {
	return o.ApplyT(func(v GetAccessCredentialsResult) bool { return v.LeaseRenewable }).(pulumi.BoolOutput)
}

func (o GetAccessCredentialsResultOutput) LeaseStartTime() pulumi.StringOutput {
	return o.ApplyT(func(v GetAccessCredentialsResult) string { return v.LeaseStartTime }).(pulumi.StringOutput)
}

func (o GetAccessCredentialsResultOutput) MaxCredValidationSeconds() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GetAccessCredentialsResult) *int { return v.MaxCredValidationSeconds }).(pulumi.IntPtrOutput)
}

func (o GetAccessCredentialsResultOutput) Namespace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetAccessCredentialsResult) *string { return v.Namespace }).(pulumi.StringPtrOutput)
}

func (o GetAccessCredentialsResultOutput) NumSecondsBetweenTests() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GetAccessCredentialsResult) *int { return v.NumSecondsBetweenTests }).(pulumi.IntPtrOutput)
}

func (o GetAccessCredentialsResultOutput) NumSequentialSuccesses() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GetAccessCredentialsResult) *int { return v.NumSequentialSuccesses }).(pulumi.IntPtrOutput)
}

func (o GetAccessCredentialsResultOutput) Role() pulumi.StringOutput {
	return o.ApplyT(func(v GetAccessCredentialsResult) string { return v.Role }).(pulumi.StringOutput)
}

func (o GetAccessCredentialsResultOutput) SubscriptionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetAccessCredentialsResult) *string { return v.SubscriptionId }).(pulumi.StringPtrOutput)
}

func (o GetAccessCredentialsResultOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetAccessCredentialsResult) *string { return v.TenantId }).(pulumi.StringPtrOutput)
}

func (o GetAccessCredentialsResultOutput) ValidateCreds() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetAccessCredentialsResult) *bool { return v.ValidateCreds }).(pulumi.BoolPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetAccessCredentialsResultOutput{})
}
