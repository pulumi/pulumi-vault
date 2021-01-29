// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package vault

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-vault/sdk/v3/go/vault"
// 	"github.com/pulumi/pulumi-vault/sdk/v3/go/vault/"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := vault.NewAuthBackend(ctx, "example", &vault.AuthBackendArgs{
// 			Tune: &vault.AuthBackendTuneArgs{
// 				ListingVisibility: pulumi.String("unauth"),
// 				MaxLeaseTtl:       pulumi.String("90000s"),
// 			},
// 			Type: pulumi.String("github"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## Import
//
// Auth methods can be imported using the `path`, e.g.
//
// ```sh
//  $ pulumi import vault:index/authBackend:AuthBackend example github
// ```
type AuthBackend struct {
	pulumi.CustomResourceState

	// The accessor for this auth method
	Accessor pulumi.StringOutput `pulumi:"accessor"`
	// (Optional; Deprecated, use `tune.default_lease_ttl` if you are using Vault provider version >= 1.8) The default lease duration in seconds.
	//
	// Deprecated: Use the tune configuration block to avoid forcing creation of new resource on an update
	DefaultLeaseTtlSeconds pulumi.IntOutput `pulumi:"defaultLeaseTtlSeconds"`
	// A description of the auth method
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Specifies whether to show this mount in
	// the UI-specific listing endpoint. Valid values are "unauth" or "hidden".
	//
	// Deprecated: Use the tune configuration block to avoid forcing creation of new resource on an update
	ListingVisibility pulumi.StringOutput `pulumi:"listingVisibility"`
	// Specifies if the auth method is local only.
	Local pulumi.BoolPtrOutput `pulumi:"local"`
	// (Optional; Deprecated, use `tune.max_lease_ttl` if you are using Vault provider version >= 1.8) The maximum lease duration in seconds.
	//
	// Deprecated: Use the tune configuration block to avoid forcing creation of new resource on an update
	MaxLeaseTtlSeconds pulumi.IntOutput `pulumi:"maxLeaseTtlSeconds"`
	// The path to mount the auth method — this defaults to the name of the type
	Path pulumi.StringOutput `pulumi:"path"`
	// Extra configuration block. Structure is documented below.
	Tune AuthBackendTuneOutput `pulumi:"tune"`
	// The name of the auth method type
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewAuthBackend registers a new resource with the given unique name, arguments, and options.
func NewAuthBackend(ctx *pulumi.Context,
	name string, args *AuthBackendArgs, opts ...pulumi.ResourceOption) (*AuthBackend, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Type == nil {
		return nil, errors.New("invalid value for required argument 'Type'")
	}
	var resource AuthBackend
	err := ctx.RegisterResource("vault:index/authBackend:AuthBackend", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAuthBackend gets an existing AuthBackend resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAuthBackend(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AuthBackendState, opts ...pulumi.ResourceOption) (*AuthBackend, error) {
	var resource AuthBackend
	err := ctx.ReadResource("vault:index/authBackend:AuthBackend", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AuthBackend resources.
type authBackendState struct {
	// The accessor for this auth method
	Accessor *string `pulumi:"accessor"`
	// (Optional; Deprecated, use `tune.default_lease_ttl` if you are using Vault provider version >= 1.8) The default lease duration in seconds.
	//
	// Deprecated: Use the tune configuration block to avoid forcing creation of new resource on an update
	DefaultLeaseTtlSeconds *int `pulumi:"defaultLeaseTtlSeconds"`
	// A description of the auth method
	Description *string `pulumi:"description"`
	// Specifies whether to show this mount in
	// the UI-specific listing endpoint. Valid values are "unauth" or "hidden".
	//
	// Deprecated: Use the tune configuration block to avoid forcing creation of new resource on an update
	ListingVisibility *string `pulumi:"listingVisibility"`
	// Specifies if the auth method is local only.
	Local *bool `pulumi:"local"`
	// (Optional; Deprecated, use `tune.max_lease_ttl` if you are using Vault provider version >= 1.8) The maximum lease duration in seconds.
	//
	// Deprecated: Use the tune configuration block to avoid forcing creation of new resource on an update
	MaxLeaseTtlSeconds *int `pulumi:"maxLeaseTtlSeconds"`
	// The path to mount the auth method — this defaults to the name of the type
	Path *string `pulumi:"path"`
	// Extra configuration block. Structure is documented below.
	Tune *AuthBackendTune `pulumi:"tune"`
	// The name of the auth method type
	Type *string `pulumi:"type"`
}

type AuthBackendState struct {
	// The accessor for this auth method
	Accessor pulumi.StringPtrInput
	// (Optional; Deprecated, use `tune.default_lease_ttl` if you are using Vault provider version >= 1.8) The default lease duration in seconds.
	//
	// Deprecated: Use the tune configuration block to avoid forcing creation of new resource on an update
	DefaultLeaseTtlSeconds pulumi.IntPtrInput
	// A description of the auth method
	Description pulumi.StringPtrInput
	// Specifies whether to show this mount in
	// the UI-specific listing endpoint. Valid values are "unauth" or "hidden".
	//
	// Deprecated: Use the tune configuration block to avoid forcing creation of new resource on an update
	ListingVisibility pulumi.StringPtrInput
	// Specifies if the auth method is local only.
	Local pulumi.BoolPtrInput
	// (Optional; Deprecated, use `tune.max_lease_ttl` if you are using Vault provider version >= 1.8) The maximum lease duration in seconds.
	//
	// Deprecated: Use the tune configuration block to avoid forcing creation of new resource on an update
	MaxLeaseTtlSeconds pulumi.IntPtrInput
	// The path to mount the auth method — this defaults to the name of the type
	Path pulumi.StringPtrInput
	// Extra configuration block. Structure is documented below.
	Tune AuthBackendTunePtrInput
	// The name of the auth method type
	Type pulumi.StringPtrInput
}

func (AuthBackendState) ElementType() reflect.Type {
	return reflect.TypeOf((*authBackendState)(nil)).Elem()
}

type authBackendArgs struct {
	// (Optional; Deprecated, use `tune.default_lease_ttl` if you are using Vault provider version >= 1.8) The default lease duration in seconds.
	//
	// Deprecated: Use the tune configuration block to avoid forcing creation of new resource on an update
	DefaultLeaseTtlSeconds *int `pulumi:"defaultLeaseTtlSeconds"`
	// A description of the auth method
	Description *string `pulumi:"description"`
	// Specifies whether to show this mount in
	// the UI-specific listing endpoint. Valid values are "unauth" or "hidden".
	//
	// Deprecated: Use the tune configuration block to avoid forcing creation of new resource on an update
	ListingVisibility *string `pulumi:"listingVisibility"`
	// Specifies if the auth method is local only.
	Local *bool `pulumi:"local"`
	// (Optional; Deprecated, use `tune.max_lease_ttl` if you are using Vault provider version >= 1.8) The maximum lease duration in seconds.
	//
	// Deprecated: Use the tune configuration block to avoid forcing creation of new resource on an update
	MaxLeaseTtlSeconds *int `pulumi:"maxLeaseTtlSeconds"`
	// The path to mount the auth method — this defaults to the name of the type
	Path *string `pulumi:"path"`
	// Extra configuration block. Structure is documented below.
	Tune *AuthBackendTune `pulumi:"tune"`
	// The name of the auth method type
	Type string `pulumi:"type"`
}

// The set of arguments for constructing a AuthBackend resource.
type AuthBackendArgs struct {
	// (Optional; Deprecated, use `tune.default_lease_ttl` if you are using Vault provider version >= 1.8) The default lease duration in seconds.
	//
	// Deprecated: Use the tune configuration block to avoid forcing creation of new resource on an update
	DefaultLeaseTtlSeconds pulumi.IntPtrInput
	// A description of the auth method
	Description pulumi.StringPtrInput
	// Specifies whether to show this mount in
	// the UI-specific listing endpoint. Valid values are "unauth" or "hidden".
	//
	// Deprecated: Use the tune configuration block to avoid forcing creation of new resource on an update
	ListingVisibility pulumi.StringPtrInput
	// Specifies if the auth method is local only.
	Local pulumi.BoolPtrInput
	// (Optional; Deprecated, use `tune.max_lease_ttl` if you are using Vault provider version >= 1.8) The maximum lease duration in seconds.
	//
	// Deprecated: Use the tune configuration block to avoid forcing creation of new resource on an update
	MaxLeaseTtlSeconds pulumi.IntPtrInput
	// The path to mount the auth method — this defaults to the name of the type
	Path pulumi.StringPtrInput
	// Extra configuration block. Structure is documented below.
	Tune AuthBackendTunePtrInput
	// The name of the auth method type
	Type pulumi.StringInput
}

func (AuthBackendArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*authBackendArgs)(nil)).Elem()
}

type AuthBackendInput interface {
	pulumi.Input

	ToAuthBackendOutput() AuthBackendOutput
	ToAuthBackendOutputWithContext(ctx context.Context) AuthBackendOutput
}

func (*AuthBackend) ElementType() reflect.Type {
	return reflect.TypeOf((*AuthBackend)(nil))
}

func (i *AuthBackend) ToAuthBackendOutput() AuthBackendOutput {
	return i.ToAuthBackendOutputWithContext(context.Background())
}

func (i *AuthBackend) ToAuthBackendOutputWithContext(ctx context.Context) AuthBackendOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AuthBackendOutput)
}

type AuthBackendOutput struct {
	*pulumi.OutputState
}

func (AuthBackendOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AuthBackend)(nil))
}

func (o AuthBackendOutput) ToAuthBackendOutput() AuthBackendOutput {
	return o
}

func (o AuthBackendOutput) ToAuthBackendOutputWithContext(ctx context.Context) AuthBackendOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(AuthBackendOutput{})
}
