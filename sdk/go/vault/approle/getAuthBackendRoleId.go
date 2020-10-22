// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package approle

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Reads the Role ID of an AppRole from a Vault server.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-vault/sdk/v2/go/vault/appRole"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		opt0 := "my-approle-backend"
// 		role, err := appRole.GetAuthBackendRoleId(ctx, &appRole.GetAuthBackendRoleIdArgs{
// 			Backend:  &opt0,
// 			RoleName: "my-role",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		ctx.Export("role-id", role.RoleId)
// 		return nil
// 	})
// }
// ```
func GetAuthBackendRoleId(ctx *pulumi.Context, args *GetAuthBackendRoleIdArgs, opts ...pulumi.InvokeOption) (*GetAuthBackendRoleIdResult, error) {
	var rv GetAuthBackendRoleIdResult
	err := ctx.Invoke("vault:appRole/getAuthBackendRoleId:getAuthBackendRoleId", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getAuthBackendRoleId.
type GetAuthBackendRoleIdArgs struct {
	// The unique name for the AppRole backend the role to
	// retrieve a RoleID for resides in. Defaults to "approle".
	Backend *string `pulumi:"backend"`
	// The name of the role to retrieve the Role ID for.
	RoleName string `pulumi:"roleName"`
}

// A collection of values returned by getAuthBackendRoleId.
type GetAuthBackendRoleIdResult struct {
	Backend *string `pulumi:"backend"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The RoleID of the role.
	RoleId   string `pulumi:"roleId"`
	RoleName string `pulumi:"roleName"`
}