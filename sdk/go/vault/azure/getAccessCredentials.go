// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package azure

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetAccessCredentials(ctx *pulumi.Context, args *GetAccessCredentialsArgs, opts ...pulumi.InvokeOption) (*GetAccessCredentialsResult, error) {
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
	// If 'validate_creds' is true,
	// the number of seconds after which to give up validating credentials. Defaults
	// to 1,200 (20 minutes).
	MaxCredValidationSeconds *int `pulumi:"maxCredValidationSeconds"`
	// If 'validate_creds' is true,
	// the number of seconds to wait between each test of generated credentials.
	// Defaults to 7.
	NumSecondsBetweenTests *int `pulumi:"numSecondsBetweenTests"`
	// If 'validate_creds' is true,
	// the number of sequential successes required to validate generated
	// credentials. Defaults to 8.
	NumSequentialSuccesses *int `pulumi:"numSequentialSuccesses"`
	// The name of the Azure secret backend role to read
	// credentials from, with no leading or trailing `/`s.
	Role string `pulumi:"role"`
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
	ClientSecret string `pulumi:"clientSecret"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The duration of the secret lease, in seconds relative
	// to the time the data was requested. Once this time has passed any plan
	// generated with this data may fail to apply.
	LeaseDuration int `pulumi:"leaseDuration"`
	// The lease identifier assigned by Vault.
	LeaseId                  string `pulumi:"leaseId"`
	LeaseRenewable           bool   `pulumi:"leaseRenewable"`
	LeaseStartTime           string `pulumi:"leaseStartTime"`
	MaxCredValidationSeconds *int   `pulumi:"maxCredValidationSeconds"`
	NumSecondsBetweenTests   *int   `pulumi:"numSecondsBetweenTests"`
	NumSequentialSuccesses   *int   `pulumi:"numSequentialSuccesses"`
	Role                     string `pulumi:"role"`
	ValidateCreds            *bool  `pulumi:"validateCreds"`
}
