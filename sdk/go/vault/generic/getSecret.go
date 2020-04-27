// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package generic

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func LookupSecret(ctx *pulumi.Context, args *LookupSecretArgs, opts ...pulumi.InvokeOption) (*LookupSecretResult, error) {
	var rv LookupSecretResult
	err := ctx.Invoke("vault:generic/getSecret:getSecret", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSecret.
type LookupSecretArgs struct {
	// The full logical path from which to request data.
	// To read data from the "generic" secret backend mounted in Vault by
	// default, this should be prefixed with `secret/`. Reading from other backends
	// with this data source is possible; consult each backend's documentation
	// to see which endpoints support the `GET` method.
	Path    string `pulumi:"path"`
	Version *int   `pulumi:"version"`
}

// A collection of values returned by getSecret.
type LookupSecretResult struct {
	// A mapping whose keys are the top-level data keys returned from
	// Vault and whose values are the corresponding values. This map can only
	// represent string data, so any non-string values returned from Vault are
	// serialized as JSON.
	Data map[string]interface{} `pulumi:"data"`
	// A string containing the full data payload retrieved from
	// Vault, serialized in JSON format.
	DataJson string `pulumi:"dataJson"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The duration of the secret lease, in seconds relative
	// to the time the data was requested. Once this time has passed any plan
	// generated with this data may fail to apply.
	LeaseDuration int `pulumi:"leaseDuration"`
	// The lease identifier assigned by Vault, if any.
	LeaseId        string `pulumi:"leaseId"`
	LeaseRenewable bool   `pulumi:"leaseRenewable"`
	LeaseStartTime string `pulumi:"leaseStartTime"`
	Path           string `pulumi:"path"`
	Version        *int   `pulumi:"version"`
}
