// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package vault

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-vault/sdk/v4/go/vault"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := vault.LookupAuthBackend(ctx, &vault.LookupAuthBackendArgs{
// 			Path: "userpass",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func LookupAuthBackend(ctx *pulumi.Context, args *LookupAuthBackendArgs, opts ...pulumi.InvokeOption) (*LookupAuthBackendResult, error) {
	var rv LookupAuthBackendResult
	err := ctx.Invoke("vault:index/getAuthBackend:getAuthBackend", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getAuthBackend.
type LookupAuthBackendArgs struct {
	// The auth backend mount point.
	Path string `pulumi:"path"`
}

// A collection of values returned by getAuthBackend.
type LookupAuthBackendResult struct {
	// The accessor for this auth method
	Accessor string `pulumi:"accessor"`
	// The default lease duration in seconds.
	DefaultLeaseTtlSeconds int `pulumi:"defaultLeaseTtlSeconds"`
	// A description of the auth method.
	Description string `pulumi:"description"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Specifies whether to show this mount in the UI-specific listing endpoint.
	ListingVisibility string `pulumi:"listingVisibility"`
	// Specifies if the auth method is local only.
	Local bool `pulumi:"local"`
	// The maximum lease duration in seconds.
	MaxLeaseTtlSeconds int    `pulumi:"maxLeaseTtlSeconds"`
	Path               string `pulumi:"path"`
	// The name of the auth method type.
	Type string `pulumi:"type"`
}
