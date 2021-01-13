// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package vault

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides a resource to manage Role Governing Policy (RGP) via [Sentinel](https://www.vaultproject.io/docs/enterprise/sentinel/index.html).
//
// **Note** this feature is available only with Vault Enterprise.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"fmt"
//
// 	"github.com/pulumi/pulumi-vault/sdk/v3/go/vault"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := vault.NewRgpPolicy(ctx, "allow_all", &vault.RgpPolicyArgs{
// 			EnforcementLevel: pulumi.String("soft-mandatory"),
// 			Policy:           pulumi.String(fmt.Sprintf("%v%v%v%v", "main = rule {\n", "  true\n", "}\n", "\n")),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type RgpPolicy struct {
	pulumi.CustomResourceState

	// Enforcement level of Sentinel policy. Can be either `advisory` or `soft-mandatory` or `hard-mandatory`
	EnforcementLevel pulumi.StringOutput `pulumi:"enforcementLevel"`
	// The name of the policy
	Name pulumi.StringOutput `pulumi:"name"`
	// String containing a Sentinel policy
	Policy pulumi.StringOutput `pulumi:"policy"`
}

// NewRgpPolicy registers a new resource with the given unique name, arguments, and options.
func NewRgpPolicy(ctx *pulumi.Context,
	name string, args *RgpPolicyArgs, opts ...pulumi.ResourceOption) (*RgpPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.EnforcementLevel == nil {
		return nil, errors.New("invalid value for required argument 'EnforcementLevel'")
	}
	if args.Policy == nil {
		return nil, errors.New("invalid value for required argument 'Policy'")
	}
	var resource RgpPolicy
	err := ctx.RegisterResource("vault:index/rgpPolicy:RgpPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRgpPolicy gets an existing RgpPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRgpPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RgpPolicyState, opts ...pulumi.ResourceOption) (*RgpPolicy, error) {
	var resource RgpPolicy
	err := ctx.ReadResource("vault:index/rgpPolicy:RgpPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RgpPolicy resources.
type rgpPolicyState struct {
	// Enforcement level of Sentinel policy. Can be either `advisory` or `soft-mandatory` or `hard-mandatory`
	EnforcementLevel *string `pulumi:"enforcementLevel"`
	// The name of the policy
	Name *string `pulumi:"name"`
	// String containing a Sentinel policy
	Policy *string `pulumi:"policy"`
}

type RgpPolicyState struct {
	// Enforcement level of Sentinel policy. Can be either `advisory` or `soft-mandatory` or `hard-mandatory`
	EnforcementLevel pulumi.StringPtrInput
	// The name of the policy
	Name pulumi.StringPtrInput
	// String containing a Sentinel policy
	Policy pulumi.StringPtrInput
}

func (RgpPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*rgpPolicyState)(nil)).Elem()
}

type rgpPolicyArgs struct {
	// Enforcement level of Sentinel policy. Can be either `advisory` or `soft-mandatory` or `hard-mandatory`
	EnforcementLevel string `pulumi:"enforcementLevel"`
	// The name of the policy
	Name *string `pulumi:"name"`
	// String containing a Sentinel policy
	Policy string `pulumi:"policy"`
}

// The set of arguments for constructing a RgpPolicy resource.
type RgpPolicyArgs struct {
	// Enforcement level of Sentinel policy. Can be either `advisory` or `soft-mandatory` or `hard-mandatory`
	EnforcementLevel pulumi.StringInput
	// The name of the policy
	Name pulumi.StringPtrInput
	// String containing a Sentinel policy
	Policy pulumi.StringInput
}

func (RgpPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*rgpPolicyArgs)(nil)).Elem()
}

type RgpPolicyInput interface {
	pulumi.Input

	ToRgpPolicyOutput() RgpPolicyOutput
	ToRgpPolicyOutputWithContext(ctx context.Context) RgpPolicyOutput
}

func (RgpPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((*RgpPolicy)(nil)).Elem()
}

func (i RgpPolicy) ToRgpPolicyOutput() RgpPolicyOutput {
	return i.ToRgpPolicyOutputWithContext(context.Background())
}

func (i RgpPolicy) ToRgpPolicyOutputWithContext(ctx context.Context) RgpPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RgpPolicyOutput)
}

type RgpPolicyOutput struct {
	*pulumi.OutputState
}

func (RgpPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RgpPolicyOutput)(nil)).Elem()
}

func (o RgpPolicyOutput) ToRgpPolicyOutput() RgpPolicyOutput {
	return o
}

func (o RgpPolicyOutput) ToRgpPolicyOutputWithContext(ctx context.Context) RgpPolicyOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(RgpPolicyOutput{})
}
