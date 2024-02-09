// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package mongodbatlas

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-vault/sdk/v5/go/vault/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-vault/sdk/v5/go/vault"
//	"github.com/pulumi/pulumi-vault/sdk/v5/go/vault/mongodbatlas"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			mongo, err := vault.NewMount(ctx, "mongo", &vault.MountArgs{
//				Path:        pulumi.String("%s"),
//				Type:        pulumi.String("mongodbatlas"),
//				Description: pulumi.String("MongoDB Atlas secret engine mount"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = mongodbatlas.NewSecretBackend(ctx, "config", &mongodbatlas.SecretBackendArgs{
//				Mount:      pulumi.String("vault_mount.mongo.path"),
//				PrivateKey: pulumi.String("privateKey"),
//				PublicKey:  pulumi.String("publicKey"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = mongodbatlas.NewSecretRole(ctx, "role", &mongodbatlas.SecretRoleArgs{
//				Mount:          mongo.Path,
//				OrganizationId: pulumi.String("7cf5a45a9ccf6400e60981b7"),
//				ProjectId:      pulumi.String("5cf5a45a9ccf6400e60981b6"),
//				Roles:          pulumi.StringArray("ORG_READ_ONLY"),
//				IpAddresses:    pulumi.StringArray("192.168.1.5, 192.168.1.6"),
//				CidrBlocks:     pulumi.StringArray("192.168.1.3/35"),
//				ProjectRoles:   pulumi.StringArray("GROUP_READ_ONLY"),
//				Ttl:            pulumi.String("60"),
//				MaxTtl:         pulumi.String("120"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// ## Import
//
// The MongoDB Atlas secret role can be imported using the full path to the role of the form: `<mount_path>/roles/<role_name>` e.g.
//
// ```sh
//
//	$ pulumi import vault:mongodbatlas/secretRole:SecretRole example mongodbatlas/roles/example-role
//
// ```
type SecretRole struct {
	pulumi.CustomResourceState

	// Whitelist entry in CIDR notation to be added for the API key.
	CidrBlocks pulumi.StringArrayOutput `pulumi:"cidrBlocks"`
	// IP address to be added to the whitelist for the API key.
	IpAddresses pulumi.StringArrayOutput `pulumi:"ipAddresses"`
	// The maximum allowed lifetime of credentials issued using this role.
	MaxTtl pulumi.StringPtrOutput `pulumi:"maxTtl"`
	// Path where the MongoDB Atlas Secrets Engine is mounted.
	Mount pulumi.StringOutput `pulumi:"mount"`
	// The name of the role.
	Name pulumi.StringOutput `pulumi:"name"`
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
	// *Available only for Vault Enterprise*.
	Namespace pulumi.StringPtrOutput `pulumi:"namespace"`
	// Unique identifier for the organization to which the target API Key belongs.
	// Required if `projectId` is not set.
	OrganizationId pulumi.StringPtrOutput `pulumi:"organizationId"`
	// Unique identifier for the project to which the target API Key belongs.
	// Required if `organizationId is` not set.
	ProjectId pulumi.StringPtrOutput `pulumi:"projectId"`
	// Roles assigned when an org API key is assigned to a project API key.
	ProjectRoles pulumi.StringArrayOutput `pulumi:"projectRoles"`
	// List of roles that the API Key needs to have.
	Roles pulumi.StringArrayOutput `pulumi:"roles"`
	// Duration in seconds after which the issued credential should expire.
	Ttl pulumi.StringPtrOutput `pulumi:"ttl"`
}

// NewSecretRole registers a new resource with the given unique name, arguments, and options.
func NewSecretRole(ctx *pulumi.Context,
	name string, args *SecretRoleArgs, opts ...pulumi.ResourceOption) (*SecretRole, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Mount == nil {
		return nil, errors.New("invalid value for required argument 'Mount'")
	}
	if args.Roles == nil {
		return nil, errors.New("invalid value for required argument 'Roles'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource SecretRole
	err := ctx.RegisterResource("vault:mongodbatlas/secretRole:SecretRole", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSecretRole gets an existing SecretRole resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSecretRole(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SecretRoleState, opts ...pulumi.ResourceOption) (*SecretRole, error) {
	var resource SecretRole
	err := ctx.ReadResource("vault:mongodbatlas/secretRole:SecretRole", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SecretRole resources.
type secretRoleState struct {
	// Whitelist entry in CIDR notation to be added for the API key.
	CidrBlocks []string `pulumi:"cidrBlocks"`
	// IP address to be added to the whitelist for the API key.
	IpAddresses []string `pulumi:"ipAddresses"`
	// The maximum allowed lifetime of credentials issued using this role.
	MaxTtl *string `pulumi:"maxTtl"`
	// Path where the MongoDB Atlas Secrets Engine is mounted.
	Mount *string `pulumi:"mount"`
	// The name of the role.
	Name *string `pulumi:"name"`
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
	// *Available only for Vault Enterprise*.
	Namespace *string `pulumi:"namespace"`
	// Unique identifier for the organization to which the target API Key belongs.
	// Required if `projectId` is not set.
	OrganizationId *string `pulumi:"organizationId"`
	// Unique identifier for the project to which the target API Key belongs.
	// Required if `organizationId is` not set.
	ProjectId *string `pulumi:"projectId"`
	// Roles assigned when an org API key is assigned to a project API key.
	ProjectRoles []string `pulumi:"projectRoles"`
	// List of roles that the API Key needs to have.
	Roles []string `pulumi:"roles"`
	// Duration in seconds after which the issued credential should expire.
	Ttl *string `pulumi:"ttl"`
}

type SecretRoleState struct {
	// Whitelist entry in CIDR notation to be added for the API key.
	CidrBlocks pulumi.StringArrayInput
	// IP address to be added to the whitelist for the API key.
	IpAddresses pulumi.StringArrayInput
	// The maximum allowed lifetime of credentials issued using this role.
	MaxTtl pulumi.StringPtrInput
	// Path where the MongoDB Atlas Secrets Engine is mounted.
	Mount pulumi.StringPtrInput
	// The name of the role.
	Name pulumi.StringPtrInput
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
	// *Available only for Vault Enterprise*.
	Namespace pulumi.StringPtrInput
	// Unique identifier for the organization to which the target API Key belongs.
	// Required if `projectId` is not set.
	OrganizationId pulumi.StringPtrInput
	// Unique identifier for the project to which the target API Key belongs.
	// Required if `organizationId is` not set.
	ProjectId pulumi.StringPtrInput
	// Roles assigned when an org API key is assigned to a project API key.
	ProjectRoles pulumi.StringArrayInput
	// List of roles that the API Key needs to have.
	Roles pulumi.StringArrayInput
	// Duration in seconds after which the issued credential should expire.
	Ttl pulumi.StringPtrInput
}

func (SecretRoleState) ElementType() reflect.Type {
	return reflect.TypeOf((*secretRoleState)(nil)).Elem()
}

type secretRoleArgs struct {
	// Whitelist entry in CIDR notation to be added for the API key.
	CidrBlocks []string `pulumi:"cidrBlocks"`
	// IP address to be added to the whitelist for the API key.
	IpAddresses []string `pulumi:"ipAddresses"`
	// The maximum allowed lifetime of credentials issued using this role.
	MaxTtl *string `pulumi:"maxTtl"`
	// Path where the MongoDB Atlas Secrets Engine is mounted.
	Mount string `pulumi:"mount"`
	// The name of the role.
	Name *string `pulumi:"name"`
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
	// *Available only for Vault Enterprise*.
	Namespace *string `pulumi:"namespace"`
	// Unique identifier for the organization to which the target API Key belongs.
	// Required if `projectId` is not set.
	OrganizationId *string `pulumi:"organizationId"`
	// Unique identifier for the project to which the target API Key belongs.
	// Required if `organizationId is` not set.
	ProjectId *string `pulumi:"projectId"`
	// Roles assigned when an org API key is assigned to a project API key.
	ProjectRoles []string `pulumi:"projectRoles"`
	// List of roles that the API Key needs to have.
	Roles []string `pulumi:"roles"`
	// Duration in seconds after which the issued credential should expire.
	Ttl *string `pulumi:"ttl"`
}

// The set of arguments for constructing a SecretRole resource.
type SecretRoleArgs struct {
	// Whitelist entry in CIDR notation to be added for the API key.
	CidrBlocks pulumi.StringArrayInput
	// IP address to be added to the whitelist for the API key.
	IpAddresses pulumi.StringArrayInput
	// The maximum allowed lifetime of credentials issued using this role.
	MaxTtl pulumi.StringPtrInput
	// Path where the MongoDB Atlas Secrets Engine is mounted.
	Mount pulumi.StringInput
	// The name of the role.
	Name pulumi.StringPtrInput
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
	// *Available only for Vault Enterprise*.
	Namespace pulumi.StringPtrInput
	// Unique identifier for the organization to which the target API Key belongs.
	// Required if `projectId` is not set.
	OrganizationId pulumi.StringPtrInput
	// Unique identifier for the project to which the target API Key belongs.
	// Required if `organizationId is` not set.
	ProjectId pulumi.StringPtrInput
	// Roles assigned when an org API key is assigned to a project API key.
	ProjectRoles pulumi.StringArrayInput
	// List of roles that the API Key needs to have.
	Roles pulumi.StringArrayInput
	// Duration in seconds after which the issued credential should expire.
	Ttl pulumi.StringPtrInput
}

func (SecretRoleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*secretRoleArgs)(nil)).Elem()
}

type SecretRoleInput interface {
	pulumi.Input

	ToSecretRoleOutput() SecretRoleOutput
	ToSecretRoleOutputWithContext(ctx context.Context) SecretRoleOutput
}

func (*SecretRole) ElementType() reflect.Type {
	return reflect.TypeOf((**SecretRole)(nil)).Elem()
}

func (i *SecretRole) ToSecretRoleOutput() SecretRoleOutput {
	return i.ToSecretRoleOutputWithContext(context.Background())
}

func (i *SecretRole) ToSecretRoleOutputWithContext(ctx context.Context) SecretRoleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecretRoleOutput)
}

// SecretRoleArrayInput is an input type that accepts SecretRoleArray and SecretRoleArrayOutput values.
// You can construct a concrete instance of `SecretRoleArrayInput` via:
//
//	SecretRoleArray{ SecretRoleArgs{...} }
type SecretRoleArrayInput interface {
	pulumi.Input

	ToSecretRoleArrayOutput() SecretRoleArrayOutput
	ToSecretRoleArrayOutputWithContext(context.Context) SecretRoleArrayOutput
}

type SecretRoleArray []SecretRoleInput

func (SecretRoleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SecretRole)(nil)).Elem()
}

func (i SecretRoleArray) ToSecretRoleArrayOutput() SecretRoleArrayOutput {
	return i.ToSecretRoleArrayOutputWithContext(context.Background())
}

func (i SecretRoleArray) ToSecretRoleArrayOutputWithContext(ctx context.Context) SecretRoleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecretRoleArrayOutput)
}

// SecretRoleMapInput is an input type that accepts SecretRoleMap and SecretRoleMapOutput values.
// You can construct a concrete instance of `SecretRoleMapInput` via:
//
//	SecretRoleMap{ "key": SecretRoleArgs{...} }
type SecretRoleMapInput interface {
	pulumi.Input

	ToSecretRoleMapOutput() SecretRoleMapOutput
	ToSecretRoleMapOutputWithContext(context.Context) SecretRoleMapOutput
}

type SecretRoleMap map[string]SecretRoleInput

func (SecretRoleMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SecretRole)(nil)).Elem()
}

func (i SecretRoleMap) ToSecretRoleMapOutput() SecretRoleMapOutput {
	return i.ToSecretRoleMapOutputWithContext(context.Background())
}

func (i SecretRoleMap) ToSecretRoleMapOutputWithContext(ctx context.Context) SecretRoleMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecretRoleMapOutput)
}

type SecretRoleOutput struct{ *pulumi.OutputState }

func (SecretRoleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SecretRole)(nil)).Elem()
}

func (o SecretRoleOutput) ToSecretRoleOutput() SecretRoleOutput {
	return o
}

func (o SecretRoleOutput) ToSecretRoleOutputWithContext(ctx context.Context) SecretRoleOutput {
	return o
}

// Whitelist entry in CIDR notation to be added for the API key.
func (o SecretRoleOutput) CidrBlocks() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *SecretRole) pulumi.StringArrayOutput { return v.CidrBlocks }).(pulumi.StringArrayOutput)
}

// IP address to be added to the whitelist for the API key.
func (o SecretRoleOutput) IpAddresses() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *SecretRole) pulumi.StringArrayOutput { return v.IpAddresses }).(pulumi.StringArrayOutput)
}

// The maximum allowed lifetime of credentials issued using this role.
func (o SecretRoleOutput) MaxTtl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecretRole) pulumi.StringPtrOutput { return v.MaxTtl }).(pulumi.StringPtrOutput)
}

// Path where the MongoDB Atlas Secrets Engine is mounted.
func (o SecretRoleOutput) Mount() pulumi.StringOutput {
	return o.ApplyT(func(v *SecretRole) pulumi.StringOutput { return v.Mount }).(pulumi.StringOutput)
}

// The name of the role.
func (o SecretRoleOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SecretRole) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The namespace to provision the resource in.
// The value should not contain leading or trailing forward slashes.
// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
// *Available only for Vault Enterprise*.
func (o SecretRoleOutput) Namespace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecretRole) pulumi.StringPtrOutput { return v.Namespace }).(pulumi.StringPtrOutput)
}

// Unique identifier for the organization to which the target API Key belongs.
// Required if `projectId` is not set.
func (o SecretRoleOutput) OrganizationId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecretRole) pulumi.StringPtrOutput { return v.OrganizationId }).(pulumi.StringPtrOutput)
}

// Unique identifier for the project to which the target API Key belongs.
// Required if `organizationId is` not set.
func (o SecretRoleOutput) ProjectId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecretRole) pulumi.StringPtrOutput { return v.ProjectId }).(pulumi.StringPtrOutput)
}

// Roles assigned when an org API key is assigned to a project API key.
func (o SecretRoleOutput) ProjectRoles() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *SecretRole) pulumi.StringArrayOutput { return v.ProjectRoles }).(pulumi.StringArrayOutput)
}

// List of roles that the API Key needs to have.
func (o SecretRoleOutput) Roles() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *SecretRole) pulumi.StringArrayOutput { return v.Roles }).(pulumi.StringArrayOutput)
}

// Duration in seconds after which the issued credential should expire.
func (o SecretRoleOutput) Ttl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecretRole) pulumi.StringPtrOutput { return v.Ttl }).(pulumi.StringPtrOutput)
}

type SecretRoleArrayOutput struct{ *pulumi.OutputState }

func (SecretRoleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SecretRole)(nil)).Elem()
}

func (o SecretRoleArrayOutput) ToSecretRoleArrayOutput() SecretRoleArrayOutput {
	return o
}

func (o SecretRoleArrayOutput) ToSecretRoleArrayOutputWithContext(ctx context.Context) SecretRoleArrayOutput {
	return o
}

func (o SecretRoleArrayOutput) Index(i pulumi.IntInput) SecretRoleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SecretRole {
		return vs[0].([]*SecretRole)[vs[1].(int)]
	}).(SecretRoleOutput)
}

type SecretRoleMapOutput struct{ *pulumi.OutputState }

func (SecretRoleMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SecretRole)(nil)).Elem()
}

func (o SecretRoleMapOutput) ToSecretRoleMapOutput() SecretRoleMapOutput {
	return o
}

func (o SecretRoleMapOutput) ToSecretRoleMapOutputWithContext(ctx context.Context) SecretRoleMapOutput {
	return o
}

func (o SecretRoleMapOutput) MapIndex(k pulumi.StringInput) SecretRoleOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SecretRole {
		return vs[0].(map[string]*SecretRole)[vs[1].(string)]
	}).(SecretRoleOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SecretRoleInput)(nil)).Elem(), &SecretRole{})
	pulumi.RegisterInputType(reflect.TypeOf((*SecretRoleArrayInput)(nil)).Elem(), SecretRoleArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SecretRoleMapInput)(nil)).Elem(), SecretRoleMap{})
	pulumi.RegisterOutputType(SecretRoleOutput{})
	pulumi.RegisterOutputType(SecretRoleArrayOutput{})
	pulumi.RegisterOutputType(SecretRoleMapOutput{})
}
