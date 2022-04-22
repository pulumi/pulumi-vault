// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package generic

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"fmt"
//
// 	"github.com/pulumi/pulumi-vault/sdk/v5/go/vault"
// 	"github.com/pulumi/pulumi-vault/sdk/v5/go/vault/generic"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		userpass, err := vault.NewAuthBackend(ctx, "userpass", &vault.AuthBackendArgs{
// 			Type: pulumi.String("userpass"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		u1, err := generic.NewEndpoint(ctx, "u1", &generic.EndpointArgs{
// 			Path:               pulumi.String("auth/userpass/users/u1"),
// 			IgnoreAbsentFields: pulumi.Bool(true),
// 			DataJson:           pulumi.String(fmt.Sprintf("%v%v%v%v", "{\n", "  \"policies\": [\"p1\"],\n", "  \"password\": \"changeme\"\n", "}\n")),
// 		}, pulumi.DependsOn([]pulumi.Resource{
// 			userpass,
// 		}))
// 		if err != nil {
// 			return err
// 		}
// 		u1Token, err := generic.NewEndpoint(ctx, "u1Token", &generic.EndpointArgs{
// 			Path:          pulumi.String("auth/userpass/login/u1"),
// 			DisableRead:   pulumi.Bool(true),
// 			DisableDelete: pulumi.Bool(true),
// 			DataJson:      pulumi.String(fmt.Sprintf("%v%v%v", "{\n", "  \"password\": \"changeme\"\n", "}\n")),
// 		}, pulumi.DependsOn([]pulumi.Resource{
// 			u1,
// 		}))
// 		if err != nil {
// 			return err
// 		}
// 		u1Entity, err := generic.NewEndpoint(ctx, "u1Entity", &generic.EndpointArgs{
// 			DisableRead:        pulumi.Bool(true),
// 			DisableDelete:      pulumi.Bool(true),
// 			Path:               pulumi.String("identity/lookup/entity"),
// 			IgnoreAbsentFields: pulumi.Bool(true),
// 			WriteFields: pulumi.StringArray{
// 				pulumi.String("id"),
// 			},
// 			DataJson: pulumi.String(fmt.Sprintf("%v%v%v%v", "{\n", "  \"alias_name\": \"u1\",\n", "  \"alias_mount_accessor\": vault_auth_backend.userpass.accessor\n", "}\n")),
// 		}, pulumi.DependsOn([]pulumi.Resource{
// 			u1Token,
// 		}))
// 		if err != nil {
// 			return err
// 		}
// 		ctx.Export("u1Id", u1Entity.WriteData.ApplyT(func(writeData map[string]string) (string, error) {
// 			return writeData.Id, nil
// 		}).(pulumi.StringOutput))
// 		return nil
// 	})
// }
// ```
// ## Required Vault Capabilities
//
// Use of this resource requires the `create` or `update` capability
// (depending on whether the resource already exists) on the given path. If
// `disableDelete` is false, the `delete` capbility is also required. If
// `disableDelete` is false, the `read` capbility is required.
//
// ## Import
//
// Import is not supported for this resource.
type Endpoint struct {
	pulumi.CustomResourceState

	// String containing a JSON-encoded object that will be
	// written to the given path as the secret data.
	DataJson pulumi.StringOutput `pulumi:"dataJson"`
	// Don't attempt to delete the path from Vault if true
	DisableDelete pulumi.BoolPtrOutput `pulumi:"disableDelete"`
	// True/false. Set this to true if your vault
	// authentication is not able to read the data or if the endpoint does
	// not support the `GET` method. Setting this to `true` will break drift
	// detection. You should set this to `true` for endpoints that are
	// write-only. Defaults to false.
	DisableRead pulumi.BoolPtrOutput `pulumi:"disableRead"`
	// When reading, disregard fields not present in data_json
	IgnoreAbsentFields pulumi.BoolPtrOutput `pulumi:"ignoreAbsentFields"`
	// The full logical path at which to write the given
	// data. Consult each backend's documentation to see which endpoints
	// support the `PUT` methods and to determine whether they also support
	// `DELETE` and `GET`.
	Path pulumi.StringOutput `pulumi:"path"`
	// Map of strings returned by write operation
	WriteData pulumi.StringMapOutput `pulumi:"writeData"`
	// JSON data returned by write operation
	WriteDataJson pulumi.StringOutput `pulumi:"writeDataJson"`
	// Top-level fields returned by write to persist in state
	WriteFields pulumi.StringArrayOutput `pulumi:"writeFields"`
}

// NewEndpoint registers a new resource with the given unique name, arguments, and options.
func NewEndpoint(ctx *pulumi.Context,
	name string, args *EndpointArgs, opts ...pulumi.ResourceOption) (*Endpoint, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DataJson == nil {
		return nil, errors.New("invalid value for required argument 'DataJson'")
	}
	if args.Path == nil {
		return nil, errors.New("invalid value for required argument 'Path'")
	}
	var resource Endpoint
	err := ctx.RegisterResource("vault:generic/endpoint:Endpoint", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetEndpoint gets an existing Endpoint resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEndpoint(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EndpointState, opts ...pulumi.ResourceOption) (*Endpoint, error) {
	var resource Endpoint
	err := ctx.ReadResource("vault:generic/endpoint:Endpoint", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Endpoint resources.
type endpointState struct {
	// String containing a JSON-encoded object that will be
	// written to the given path as the secret data.
	DataJson *string `pulumi:"dataJson"`
	// Don't attempt to delete the path from Vault if true
	DisableDelete *bool `pulumi:"disableDelete"`
	// True/false. Set this to true if your vault
	// authentication is not able to read the data or if the endpoint does
	// not support the `GET` method. Setting this to `true` will break drift
	// detection. You should set this to `true` for endpoints that are
	// write-only. Defaults to false.
	DisableRead *bool `pulumi:"disableRead"`
	// When reading, disregard fields not present in data_json
	IgnoreAbsentFields *bool `pulumi:"ignoreAbsentFields"`
	// The full logical path at which to write the given
	// data. Consult each backend's documentation to see which endpoints
	// support the `PUT` methods and to determine whether they also support
	// `DELETE` and `GET`.
	Path *string `pulumi:"path"`
	// Map of strings returned by write operation
	WriteData map[string]string `pulumi:"writeData"`
	// JSON data returned by write operation
	WriteDataJson *string `pulumi:"writeDataJson"`
	// Top-level fields returned by write to persist in state
	WriteFields []string `pulumi:"writeFields"`
}

type EndpointState struct {
	// String containing a JSON-encoded object that will be
	// written to the given path as the secret data.
	DataJson pulumi.StringPtrInput
	// Don't attempt to delete the path from Vault if true
	DisableDelete pulumi.BoolPtrInput
	// True/false. Set this to true if your vault
	// authentication is not able to read the data or if the endpoint does
	// not support the `GET` method. Setting this to `true` will break drift
	// detection. You should set this to `true` for endpoints that are
	// write-only. Defaults to false.
	DisableRead pulumi.BoolPtrInput
	// When reading, disregard fields not present in data_json
	IgnoreAbsentFields pulumi.BoolPtrInput
	// The full logical path at which to write the given
	// data. Consult each backend's documentation to see which endpoints
	// support the `PUT` methods and to determine whether they also support
	// `DELETE` and `GET`.
	Path pulumi.StringPtrInput
	// Map of strings returned by write operation
	WriteData pulumi.StringMapInput
	// JSON data returned by write operation
	WriteDataJson pulumi.StringPtrInput
	// Top-level fields returned by write to persist in state
	WriteFields pulumi.StringArrayInput
}

func (EndpointState) ElementType() reflect.Type {
	return reflect.TypeOf((*endpointState)(nil)).Elem()
}

type endpointArgs struct {
	// String containing a JSON-encoded object that will be
	// written to the given path as the secret data.
	DataJson string `pulumi:"dataJson"`
	// Don't attempt to delete the path from Vault if true
	DisableDelete *bool `pulumi:"disableDelete"`
	// True/false. Set this to true if your vault
	// authentication is not able to read the data or if the endpoint does
	// not support the `GET` method. Setting this to `true` will break drift
	// detection. You should set this to `true` for endpoints that are
	// write-only. Defaults to false.
	DisableRead *bool `pulumi:"disableRead"`
	// When reading, disregard fields not present in data_json
	IgnoreAbsentFields *bool `pulumi:"ignoreAbsentFields"`
	// The full logical path at which to write the given
	// data. Consult each backend's documentation to see which endpoints
	// support the `PUT` methods and to determine whether they also support
	// `DELETE` and `GET`.
	Path string `pulumi:"path"`
	// Top-level fields returned by write to persist in state
	WriteFields []string `pulumi:"writeFields"`
}

// The set of arguments for constructing a Endpoint resource.
type EndpointArgs struct {
	// String containing a JSON-encoded object that will be
	// written to the given path as the secret data.
	DataJson pulumi.StringInput
	// Don't attempt to delete the path from Vault if true
	DisableDelete pulumi.BoolPtrInput
	// True/false. Set this to true if your vault
	// authentication is not able to read the data or if the endpoint does
	// not support the `GET` method. Setting this to `true` will break drift
	// detection. You should set this to `true` for endpoints that are
	// write-only. Defaults to false.
	DisableRead pulumi.BoolPtrInput
	// When reading, disregard fields not present in data_json
	IgnoreAbsentFields pulumi.BoolPtrInput
	// The full logical path at which to write the given
	// data. Consult each backend's documentation to see which endpoints
	// support the `PUT` methods and to determine whether they also support
	// `DELETE` and `GET`.
	Path pulumi.StringInput
	// Top-level fields returned by write to persist in state
	WriteFields pulumi.StringArrayInput
}

func (EndpointArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*endpointArgs)(nil)).Elem()
}

type EndpointInput interface {
	pulumi.Input

	ToEndpointOutput() EndpointOutput
	ToEndpointOutputWithContext(ctx context.Context) EndpointOutput
}

func (*Endpoint) ElementType() reflect.Type {
	return reflect.TypeOf((**Endpoint)(nil)).Elem()
}

func (i *Endpoint) ToEndpointOutput() EndpointOutput {
	return i.ToEndpointOutputWithContext(context.Background())
}

func (i *Endpoint) ToEndpointOutputWithContext(ctx context.Context) EndpointOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EndpointOutput)
}

// EndpointArrayInput is an input type that accepts EndpointArray and EndpointArrayOutput values.
// You can construct a concrete instance of `EndpointArrayInput` via:
//
//          EndpointArray{ EndpointArgs{...} }
type EndpointArrayInput interface {
	pulumi.Input

	ToEndpointArrayOutput() EndpointArrayOutput
	ToEndpointArrayOutputWithContext(context.Context) EndpointArrayOutput
}

type EndpointArray []EndpointInput

func (EndpointArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Endpoint)(nil)).Elem()
}

func (i EndpointArray) ToEndpointArrayOutput() EndpointArrayOutput {
	return i.ToEndpointArrayOutputWithContext(context.Background())
}

func (i EndpointArray) ToEndpointArrayOutputWithContext(ctx context.Context) EndpointArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EndpointArrayOutput)
}

// EndpointMapInput is an input type that accepts EndpointMap and EndpointMapOutput values.
// You can construct a concrete instance of `EndpointMapInput` via:
//
//          EndpointMap{ "key": EndpointArgs{...} }
type EndpointMapInput interface {
	pulumi.Input

	ToEndpointMapOutput() EndpointMapOutput
	ToEndpointMapOutputWithContext(context.Context) EndpointMapOutput
}

type EndpointMap map[string]EndpointInput

func (EndpointMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Endpoint)(nil)).Elem()
}

func (i EndpointMap) ToEndpointMapOutput() EndpointMapOutput {
	return i.ToEndpointMapOutputWithContext(context.Background())
}

func (i EndpointMap) ToEndpointMapOutputWithContext(ctx context.Context) EndpointMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EndpointMapOutput)
}

type EndpointOutput struct{ *pulumi.OutputState }

func (EndpointOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Endpoint)(nil)).Elem()
}

func (o EndpointOutput) ToEndpointOutput() EndpointOutput {
	return o
}

func (o EndpointOutput) ToEndpointOutputWithContext(ctx context.Context) EndpointOutput {
	return o
}

type EndpointArrayOutput struct{ *pulumi.OutputState }

func (EndpointArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Endpoint)(nil)).Elem()
}

func (o EndpointArrayOutput) ToEndpointArrayOutput() EndpointArrayOutput {
	return o
}

func (o EndpointArrayOutput) ToEndpointArrayOutputWithContext(ctx context.Context) EndpointArrayOutput {
	return o
}

func (o EndpointArrayOutput) Index(i pulumi.IntInput) EndpointOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Endpoint {
		return vs[0].([]*Endpoint)[vs[1].(int)]
	}).(EndpointOutput)
}

type EndpointMapOutput struct{ *pulumi.OutputState }

func (EndpointMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Endpoint)(nil)).Elem()
}

func (o EndpointMapOutput) ToEndpointMapOutput() EndpointMapOutput {
	return o
}

func (o EndpointMapOutput) ToEndpointMapOutputWithContext(ctx context.Context) EndpointMapOutput {
	return o
}

func (o EndpointMapOutput) MapIndex(k pulumi.StringInput) EndpointOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Endpoint {
		return vs[0].(map[string]*Endpoint)[vs[1].(string)]
	}).(EndpointOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*EndpointInput)(nil)).Elem(), &Endpoint{})
	pulumi.RegisterInputType(reflect.TypeOf((*EndpointArrayInput)(nil)).Elem(), EndpointArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*EndpointMapInput)(nil)).Elem(), EndpointMap{})
	pulumi.RegisterOutputType(EndpointOutput{})
	pulumi.RegisterOutputType(EndpointArrayOutput{})
	pulumi.RegisterOutputType(EndpointMapOutput{})
}
