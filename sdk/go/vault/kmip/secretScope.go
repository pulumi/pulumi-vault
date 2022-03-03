// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package kmip

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages KMIP Secret Scopes in a Vault server. This feature requires
// Vault Enterprise. See the [Vault documentation](https://www.vaultproject.io/docs/secrets/kmip)
// for more information.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-vault/sdk/v5/go/vault/kmip"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := kmip.NewSecretBackend(ctx, "default", &kmip.SecretBackendArgs{
// 			Path:        pulumi.String("kmip"),
// 			Description: pulumi.String("Vault KMIP backend"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = kmip.NewSecretScope(ctx, "dev", &kmip.SecretScopeArgs{
// 			Path:  _default.Path,
// 			Scope: pulumi.String("dev"),
// 			Force: pulumi.Bool(true),
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
// KMIP Secret scope can be imported using the `path`, e.g.
//
// ```sh
//  $ pulumi import vault:kmip/secretScope:SecretScope dev kmip
// ```
type SecretScope struct {
	pulumi.CustomResourceState

	// Boolean field to force deletion even if there are managed objects in the scope.
	Force pulumi.BoolPtrOutput `pulumi:"force"`
	// The unique path this backend should be mounted at. Must
	// not begin or end with a `/`. Defaults to `kmip`.
	Path pulumi.StringOutput `pulumi:"path"`
	// Name of the scope.
	Scope pulumi.StringOutput `pulumi:"scope"`
}

// NewSecretScope registers a new resource with the given unique name, arguments, and options.
func NewSecretScope(ctx *pulumi.Context,
	name string, args *SecretScopeArgs, opts ...pulumi.ResourceOption) (*SecretScope, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Path == nil {
		return nil, errors.New("invalid value for required argument 'Path'")
	}
	if args.Scope == nil {
		return nil, errors.New("invalid value for required argument 'Scope'")
	}
	var resource SecretScope
	err := ctx.RegisterResource("vault:kmip/secretScope:SecretScope", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSecretScope gets an existing SecretScope resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSecretScope(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SecretScopeState, opts ...pulumi.ResourceOption) (*SecretScope, error) {
	var resource SecretScope
	err := ctx.ReadResource("vault:kmip/secretScope:SecretScope", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SecretScope resources.
type secretScopeState struct {
	// Boolean field to force deletion even if there are managed objects in the scope.
	Force *bool `pulumi:"force"`
	// The unique path this backend should be mounted at. Must
	// not begin or end with a `/`. Defaults to `kmip`.
	Path *string `pulumi:"path"`
	// Name of the scope.
	Scope *string `pulumi:"scope"`
}

type SecretScopeState struct {
	// Boolean field to force deletion even if there are managed objects in the scope.
	Force pulumi.BoolPtrInput
	// The unique path this backend should be mounted at. Must
	// not begin or end with a `/`. Defaults to `kmip`.
	Path pulumi.StringPtrInput
	// Name of the scope.
	Scope pulumi.StringPtrInput
}

func (SecretScopeState) ElementType() reflect.Type {
	return reflect.TypeOf((*secretScopeState)(nil)).Elem()
}

type secretScopeArgs struct {
	// Boolean field to force deletion even if there are managed objects in the scope.
	Force *bool `pulumi:"force"`
	// The unique path this backend should be mounted at. Must
	// not begin or end with a `/`. Defaults to `kmip`.
	Path string `pulumi:"path"`
	// Name of the scope.
	Scope string `pulumi:"scope"`
}

// The set of arguments for constructing a SecretScope resource.
type SecretScopeArgs struct {
	// Boolean field to force deletion even if there are managed objects in the scope.
	Force pulumi.BoolPtrInput
	// The unique path this backend should be mounted at. Must
	// not begin or end with a `/`. Defaults to `kmip`.
	Path pulumi.StringInput
	// Name of the scope.
	Scope pulumi.StringInput
}

func (SecretScopeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*secretScopeArgs)(nil)).Elem()
}

type SecretScopeInput interface {
	pulumi.Input

	ToSecretScopeOutput() SecretScopeOutput
	ToSecretScopeOutputWithContext(ctx context.Context) SecretScopeOutput
}

func (*SecretScope) ElementType() reflect.Type {
	return reflect.TypeOf((**SecretScope)(nil)).Elem()
}

func (i *SecretScope) ToSecretScopeOutput() SecretScopeOutput {
	return i.ToSecretScopeOutputWithContext(context.Background())
}

func (i *SecretScope) ToSecretScopeOutputWithContext(ctx context.Context) SecretScopeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecretScopeOutput)
}

// SecretScopeArrayInput is an input type that accepts SecretScopeArray and SecretScopeArrayOutput values.
// You can construct a concrete instance of `SecretScopeArrayInput` via:
//
//          SecretScopeArray{ SecretScopeArgs{...} }
type SecretScopeArrayInput interface {
	pulumi.Input

	ToSecretScopeArrayOutput() SecretScopeArrayOutput
	ToSecretScopeArrayOutputWithContext(context.Context) SecretScopeArrayOutput
}

type SecretScopeArray []SecretScopeInput

func (SecretScopeArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SecretScope)(nil)).Elem()
}

func (i SecretScopeArray) ToSecretScopeArrayOutput() SecretScopeArrayOutput {
	return i.ToSecretScopeArrayOutputWithContext(context.Background())
}

func (i SecretScopeArray) ToSecretScopeArrayOutputWithContext(ctx context.Context) SecretScopeArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecretScopeArrayOutput)
}

// SecretScopeMapInput is an input type that accepts SecretScopeMap and SecretScopeMapOutput values.
// You can construct a concrete instance of `SecretScopeMapInput` via:
//
//          SecretScopeMap{ "key": SecretScopeArgs{...} }
type SecretScopeMapInput interface {
	pulumi.Input

	ToSecretScopeMapOutput() SecretScopeMapOutput
	ToSecretScopeMapOutputWithContext(context.Context) SecretScopeMapOutput
}

type SecretScopeMap map[string]SecretScopeInput

func (SecretScopeMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SecretScope)(nil)).Elem()
}

func (i SecretScopeMap) ToSecretScopeMapOutput() SecretScopeMapOutput {
	return i.ToSecretScopeMapOutputWithContext(context.Background())
}

func (i SecretScopeMap) ToSecretScopeMapOutputWithContext(ctx context.Context) SecretScopeMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecretScopeMapOutput)
}

type SecretScopeOutput struct{ *pulumi.OutputState }

func (SecretScopeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SecretScope)(nil)).Elem()
}

func (o SecretScopeOutput) ToSecretScopeOutput() SecretScopeOutput {
	return o
}

func (o SecretScopeOutput) ToSecretScopeOutputWithContext(ctx context.Context) SecretScopeOutput {
	return o
}

type SecretScopeArrayOutput struct{ *pulumi.OutputState }

func (SecretScopeArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SecretScope)(nil)).Elem()
}

func (o SecretScopeArrayOutput) ToSecretScopeArrayOutput() SecretScopeArrayOutput {
	return o
}

func (o SecretScopeArrayOutput) ToSecretScopeArrayOutputWithContext(ctx context.Context) SecretScopeArrayOutput {
	return o
}

func (o SecretScopeArrayOutput) Index(i pulumi.IntInput) SecretScopeOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SecretScope {
		return vs[0].([]*SecretScope)[vs[1].(int)]
	}).(SecretScopeOutput)
}

type SecretScopeMapOutput struct{ *pulumi.OutputState }

func (SecretScopeMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SecretScope)(nil)).Elem()
}

func (o SecretScopeMapOutput) ToSecretScopeMapOutput() SecretScopeMapOutput {
	return o
}

func (o SecretScopeMapOutput) ToSecretScopeMapOutputWithContext(ctx context.Context) SecretScopeMapOutput {
	return o
}

func (o SecretScopeMapOutput) MapIndex(k pulumi.StringInput) SecretScopeOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SecretScope {
		return vs[0].(map[string]*SecretScope)[vs[1].(string)]
	}).(SecretScopeOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SecretScopeInput)(nil)).Elem(), &SecretScope{})
	pulumi.RegisterInputType(reflect.TypeOf((*SecretScopeArrayInput)(nil)).Elem(), SecretScopeArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SecretScopeMapInput)(nil)).Elem(), SecretScopeMap{})
	pulumi.RegisterOutputType(SecretScopeOutput{})
	pulumi.RegisterOutputType(SecretScopeArrayOutput{})
	pulumi.RegisterOutputType(SecretScopeMapOutput{})
}
