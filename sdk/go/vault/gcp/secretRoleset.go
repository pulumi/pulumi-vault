// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package gcp

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a Roleset in the [GCP Secrets Engine](https://www.vaultproject.io/docs/secrets/gcp/index.html) for Vault.
//
// Each Roleset is [tied](https://www.vaultproject.io/docs/secrets/gcp/index.html#service-accounts-are-tied-to-rolesets) to a Service Account, and can have one or more [bindings](https://www.vaultproject.io/docs/secrets/gcp/index.html#roleset-bindings) associated with it.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"fmt"
// 	"io/ioutil"
//
// 	"github.com/pulumi/pulumi-vault/sdk/v5/go/vault/gcp"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func readFileOrPanic(path string) pulumi.StringPtrInput {
// 	data, err := ioutil.ReadFile(path)
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	return pulumi.String(string(data))
// }
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		project := "my-awesome-project"
// 		gcp, err := gcp.NewSecretBackend(ctx, "gcp", &gcp.SecretBackendArgs{
// 			Path:        pulumi.String("gcp"),
// 			Credentials: readFileOrPanic("credentials.json"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = gcp.NewSecretRoleset(ctx, "roleset", &gcp.SecretRolesetArgs{
// 			Backend:    gcp.Path,
// 			Roleset:    pulumi.String("project_viewer"),
// 			SecretType: pulumi.String("access_token"),
// 			Project:    pulumi.String(project),
// 			TokenScopes: pulumi.StringArray{
// 				pulumi.String("https://www.googleapis.com/auth/cloud-platform"),
// 			},
// 			Bindings: gcp.SecretRolesetBindingArray{
// 				&gcp.SecretRolesetBindingArgs{
// 					Resource: pulumi.String(fmt.Sprintf("%v%v", "//cloudresourcemanager.googleapis.com/projects/", project)),
// 					Roles: pulumi.StringArray{
// 						pulumi.String("roles/viewer"),
// 					},
// 				},
// 			},
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
// A roleset can be imported using its Vault Path. For example, referencing the example above,
//
// ```sh
//  $ pulumi import vault:gcp/secretRoleset:SecretRoleset roleset gcp/roleset/project_viewer
// ```
type SecretRoleset struct {
	pulumi.CustomResourceState

	// Path where the GCP Secrets Engine is mounted
	Backend pulumi.StringOutput `pulumi:"backend"`
	// Bindings to create for this roleset. This can be specified multiple times for multiple bindings. Structure is documented below.
	Bindings SecretRolesetBindingArrayOutput `pulumi:"bindings"`
	// Name of the GCP project that this roleset's service account will belong to.
	Project pulumi.StringOutput `pulumi:"project"`
	// Name of the Roleset to create
	Roleset pulumi.StringOutput `pulumi:"roleset"`
	// Type of secret generated for this role set. Accepted values: `accessToken`, `serviceAccountKey`. Defaults to `accessToken`.
	SecretType pulumi.StringOutput `pulumi:"secretType"`
	// Email of the service account created by Vault for this Roleset
	ServiceAccountEmail pulumi.StringOutput `pulumi:"serviceAccountEmail"`
	// List of OAuth scopes to assign to `accessToken` secrets generated under this role set (`accessToken` role sets only).
	TokenScopes pulumi.StringArrayOutput `pulumi:"tokenScopes"`
}

// NewSecretRoleset registers a new resource with the given unique name, arguments, and options.
func NewSecretRoleset(ctx *pulumi.Context,
	name string, args *SecretRolesetArgs, opts ...pulumi.ResourceOption) (*SecretRoleset, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Backend == nil {
		return nil, errors.New("invalid value for required argument 'Backend'")
	}
	if args.Bindings == nil {
		return nil, errors.New("invalid value for required argument 'Bindings'")
	}
	if args.Project == nil {
		return nil, errors.New("invalid value for required argument 'Project'")
	}
	if args.Roleset == nil {
		return nil, errors.New("invalid value for required argument 'Roleset'")
	}
	var resource SecretRoleset
	err := ctx.RegisterResource("vault:gcp/secretRoleset:SecretRoleset", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSecretRoleset gets an existing SecretRoleset resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSecretRoleset(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SecretRolesetState, opts ...pulumi.ResourceOption) (*SecretRoleset, error) {
	var resource SecretRoleset
	err := ctx.ReadResource("vault:gcp/secretRoleset:SecretRoleset", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SecretRoleset resources.
type secretRolesetState struct {
	// Path where the GCP Secrets Engine is mounted
	Backend *string `pulumi:"backend"`
	// Bindings to create for this roleset. This can be specified multiple times for multiple bindings. Structure is documented below.
	Bindings []SecretRolesetBinding `pulumi:"bindings"`
	// Name of the GCP project that this roleset's service account will belong to.
	Project *string `pulumi:"project"`
	// Name of the Roleset to create
	Roleset *string `pulumi:"roleset"`
	// Type of secret generated for this role set. Accepted values: `accessToken`, `serviceAccountKey`. Defaults to `accessToken`.
	SecretType *string `pulumi:"secretType"`
	// Email of the service account created by Vault for this Roleset
	ServiceAccountEmail *string `pulumi:"serviceAccountEmail"`
	// List of OAuth scopes to assign to `accessToken` secrets generated under this role set (`accessToken` role sets only).
	TokenScopes []string `pulumi:"tokenScopes"`
}

type SecretRolesetState struct {
	// Path where the GCP Secrets Engine is mounted
	Backend pulumi.StringPtrInput
	// Bindings to create for this roleset. This can be specified multiple times for multiple bindings. Structure is documented below.
	Bindings SecretRolesetBindingArrayInput
	// Name of the GCP project that this roleset's service account will belong to.
	Project pulumi.StringPtrInput
	// Name of the Roleset to create
	Roleset pulumi.StringPtrInput
	// Type of secret generated for this role set. Accepted values: `accessToken`, `serviceAccountKey`. Defaults to `accessToken`.
	SecretType pulumi.StringPtrInput
	// Email of the service account created by Vault for this Roleset
	ServiceAccountEmail pulumi.StringPtrInput
	// List of OAuth scopes to assign to `accessToken` secrets generated under this role set (`accessToken` role sets only).
	TokenScopes pulumi.StringArrayInput
}

func (SecretRolesetState) ElementType() reflect.Type {
	return reflect.TypeOf((*secretRolesetState)(nil)).Elem()
}

type secretRolesetArgs struct {
	// Path where the GCP Secrets Engine is mounted
	Backend string `pulumi:"backend"`
	// Bindings to create for this roleset. This can be specified multiple times for multiple bindings. Structure is documented below.
	Bindings []SecretRolesetBinding `pulumi:"bindings"`
	// Name of the GCP project that this roleset's service account will belong to.
	Project string `pulumi:"project"`
	// Name of the Roleset to create
	Roleset string `pulumi:"roleset"`
	// Type of secret generated for this role set. Accepted values: `accessToken`, `serviceAccountKey`. Defaults to `accessToken`.
	SecretType *string `pulumi:"secretType"`
	// List of OAuth scopes to assign to `accessToken` secrets generated under this role set (`accessToken` role sets only).
	TokenScopes []string `pulumi:"tokenScopes"`
}

// The set of arguments for constructing a SecretRoleset resource.
type SecretRolesetArgs struct {
	// Path where the GCP Secrets Engine is mounted
	Backend pulumi.StringInput
	// Bindings to create for this roleset. This can be specified multiple times for multiple bindings. Structure is documented below.
	Bindings SecretRolesetBindingArrayInput
	// Name of the GCP project that this roleset's service account will belong to.
	Project pulumi.StringInput
	// Name of the Roleset to create
	Roleset pulumi.StringInput
	// Type of secret generated for this role set. Accepted values: `accessToken`, `serviceAccountKey`. Defaults to `accessToken`.
	SecretType pulumi.StringPtrInput
	// List of OAuth scopes to assign to `accessToken` secrets generated under this role set (`accessToken` role sets only).
	TokenScopes pulumi.StringArrayInput
}

func (SecretRolesetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*secretRolesetArgs)(nil)).Elem()
}

type SecretRolesetInput interface {
	pulumi.Input

	ToSecretRolesetOutput() SecretRolesetOutput
	ToSecretRolesetOutputWithContext(ctx context.Context) SecretRolesetOutput
}

func (*SecretRoleset) ElementType() reflect.Type {
	return reflect.TypeOf((**SecretRoleset)(nil)).Elem()
}

func (i *SecretRoleset) ToSecretRolesetOutput() SecretRolesetOutput {
	return i.ToSecretRolesetOutputWithContext(context.Background())
}

func (i *SecretRoleset) ToSecretRolesetOutputWithContext(ctx context.Context) SecretRolesetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecretRolesetOutput)
}

// SecretRolesetArrayInput is an input type that accepts SecretRolesetArray and SecretRolesetArrayOutput values.
// You can construct a concrete instance of `SecretRolesetArrayInput` via:
//
//          SecretRolesetArray{ SecretRolesetArgs{...} }
type SecretRolesetArrayInput interface {
	pulumi.Input

	ToSecretRolesetArrayOutput() SecretRolesetArrayOutput
	ToSecretRolesetArrayOutputWithContext(context.Context) SecretRolesetArrayOutput
}

type SecretRolesetArray []SecretRolesetInput

func (SecretRolesetArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SecretRoleset)(nil)).Elem()
}

func (i SecretRolesetArray) ToSecretRolesetArrayOutput() SecretRolesetArrayOutput {
	return i.ToSecretRolesetArrayOutputWithContext(context.Background())
}

func (i SecretRolesetArray) ToSecretRolesetArrayOutputWithContext(ctx context.Context) SecretRolesetArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecretRolesetArrayOutput)
}

// SecretRolesetMapInput is an input type that accepts SecretRolesetMap and SecretRolesetMapOutput values.
// You can construct a concrete instance of `SecretRolesetMapInput` via:
//
//          SecretRolesetMap{ "key": SecretRolesetArgs{...} }
type SecretRolesetMapInput interface {
	pulumi.Input

	ToSecretRolesetMapOutput() SecretRolesetMapOutput
	ToSecretRolesetMapOutputWithContext(context.Context) SecretRolesetMapOutput
}

type SecretRolesetMap map[string]SecretRolesetInput

func (SecretRolesetMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SecretRoleset)(nil)).Elem()
}

func (i SecretRolesetMap) ToSecretRolesetMapOutput() SecretRolesetMapOutput {
	return i.ToSecretRolesetMapOutputWithContext(context.Background())
}

func (i SecretRolesetMap) ToSecretRolesetMapOutputWithContext(ctx context.Context) SecretRolesetMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecretRolesetMapOutput)
}

type SecretRolesetOutput struct{ *pulumi.OutputState }

func (SecretRolesetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SecretRoleset)(nil)).Elem()
}

func (o SecretRolesetOutput) ToSecretRolesetOutput() SecretRolesetOutput {
	return o
}

func (o SecretRolesetOutput) ToSecretRolesetOutputWithContext(ctx context.Context) SecretRolesetOutput {
	return o
}

type SecretRolesetArrayOutput struct{ *pulumi.OutputState }

func (SecretRolesetArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SecretRoleset)(nil)).Elem()
}

func (o SecretRolesetArrayOutput) ToSecretRolesetArrayOutput() SecretRolesetArrayOutput {
	return o
}

func (o SecretRolesetArrayOutput) ToSecretRolesetArrayOutputWithContext(ctx context.Context) SecretRolesetArrayOutput {
	return o
}

func (o SecretRolesetArrayOutput) Index(i pulumi.IntInput) SecretRolesetOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SecretRoleset {
		return vs[0].([]*SecretRoleset)[vs[1].(int)]
	}).(SecretRolesetOutput)
}

type SecretRolesetMapOutput struct{ *pulumi.OutputState }

func (SecretRolesetMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SecretRoleset)(nil)).Elem()
}

func (o SecretRolesetMapOutput) ToSecretRolesetMapOutput() SecretRolesetMapOutput {
	return o
}

func (o SecretRolesetMapOutput) ToSecretRolesetMapOutputWithContext(ctx context.Context) SecretRolesetMapOutput {
	return o
}

func (o SecretRolesetMapOutput) MapIndex(k pulumi.StringInput) SecretRolesetOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SecretRoleset {
		return vs[0].(map[string]*SecretRoleset)[vs[1].(string)]
	}).(SecretRolesetOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SecretRolesetInput)(nil)).Elem(), &SecretRoleset{})
	pulumi.RegisterInputType(reflect.TypeOf((*SecretRolesetArrayInput)(nil)).Elem(), SecretRolesetArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SecretRolesetMapInput)(nil)).Elem(), SecretRolesetMap{})
	pulumi.RegisterOutputType(SecretRolesetOutput{})
	pulumi.RegisterOutputType(SecretRolesetArrayOutput{})
	pulumi.RegisterOutputType(SecretRolesetMapOutput{})
}
