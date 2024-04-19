// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package database

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-vault/sdk/v6/go/vault/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## Example Usage
//
// <!--Start PulumiCodeChooser -->
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-vault/sdk/v6/go/vault"
//	"github.com/pulumi/pulumi-vault/sdk/v6/go/vault/database"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			db, err := vault.NewMount(ctx, "db", &vault.MountArgs{
//				Path: pulumi.String("postgres"),
//				Type: pulumi.String("database"),
//			})
//			if err != nil {
//				return err
//			}
//			postgres, err := database.NewSecretBackendConnection(ctx, "postgres", &database.SecretBackendConnectionArgs{
//				Backend: db.Path,
//				Name:    pulumi.String("postgres"),
//				AllowedRoles: pulumi.StringArray{
//					pulumi.String("dev"),
//					pulumi.String("prod"),
//				},
//				Postgresql: &database.SecretBackendConnectionPostgresqlArgs{
//					ConnectionUrl: pulumi.String("postgres://username:password@host:port/database"),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			_, err = database.NewSecretBackendRole(ctx, "role", &database.SecretBackendRoleArgs{
//				Backend: db.Path,
//				Name:    pulumi.String("dev"),
//				DbName:  postgres.Name,
//				CreationStatements: pulumi.StringArray{
//					pulumi.String("CREATE ROLE \"{{name}}\" WITH LOGIN PASSWORD '{{password}}' VALID UNTIL '{{expiration}}';"),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// <!--End PulumiCodeChooser -->
//
// ## Import
//
// Database secret backend roles can be imported using the `backend`, `/roles/`, and the `name` e.g.
//
// ```sh
// $ pulumi import vault:database/secretBackendRole:SecretBackendRole example postgres/roles/my-role
// ```
type SecretBackendRole struct {
	pulumi.CustomResourceState

	// The unique name of the Vault mount to configure.
	Backend pulumi.StringOutput `pulumi:"backend"`
	// The database statements to execute when
	// creating a user.
	CreationStatements pulumi.StringArrayOutput `pulumi:"creationStatements"`
	// Specifies the configuration
	// for the given `credentialType`.
	//
	// The following options are available for each `credentialType` value:
	CredentialConfig pulumi.MapOutput `pulumi:"credentialConfig"`
	// Specifies the type of credential that
	// will be generated for the role. Options include: `password`, `rsaPrivateKey`, `clientCertificate`.
	// See the plugin's API page for credential types supported by individual databases.
	CredentialType pulumi.StringOutput `pulumi:"credentialType"`
	// The unique name of the database connection to use for
	// the role.
	DbName pulumi.StringOutput `pulumi:"dbName"`
	// The default number of seconds for leases for this
	// role.
	DefaultTtl pulumi.IntPtrOutput `pulumi:"defaultTtl"`
	// The maximum number of seconds for leases for this
	// role.
	MaxTtl pulumi.IntPtrOutput `pulumi:"maxTtl"`
	// A unique name to give the role.
	Name pulumi.StringOutput `pulumi:"name"`
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured namespace.
	// *Available only for Vault Enterprise*.
	Namespace pulumi.StringPtrOutput `pulumi:"namespace"`
	// The database statements to execute when
	// renewing a user.
	RenewStatements pulumi.StringArrayOutput `pulumi:"renewStatements"`
	// The database statements to execute when
	// revoking a user.
	RevocationStatements pulumi.StringArrayOutput `pulumi:"revocationStatements"`
	// The database statements to execute when
	// rolling back creation due to an error.
	RollbackStatements pulumi.StringArrayOutput `pulumi:"rollbackStatements"`
}

// NewSecretBackendRole registers a new resource with the given unique name, arguments, and options.
func NewSecretBackendRole(ctx *pulumi.Context,
	name string, args *SecretBackendRoleArgs, opts ...pulumi.ResourceOption) (*SecretBackendRole, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Backend == nil {
		return nil, errors.New("invalid value for required argument 'Backend'")
	}
	if args.CreationStatements == nil {
		return nil, errors.New("invalid value for required argument 'CreationStatements'")
	}
	if args.DbName == nil {
		return nil, errors.New("invalid value for required argument 'DbName'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource SecretBackendRole
	err := ctx.RegisterResource("vault:database/secretBackendRole:SecretBackendRole", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSecretBackendRole gets an existing SecretBackendRole resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSecretBackendRole(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SecretBackendRoleState, opts ...pulumi.ResourceOption) (*SecretBackendRole, error) {
	var resource SecretBackendRole
	err := ctx.ReadResource("vault:database/secretBackendRole:SecretBackendRole", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SecretBackendRole resources.
type secretBackendRoleState struct {
	// The unique name of the Vault mount to configure.
	Backend *string `pulumi:"backend"`
	// The database statements to execute when
	// creating a user.
	CreationStatements []string `pulumi:"creationStatements"`
	// Specifies the configuration
	// for the given `credentialType`.
	//
	// The following options are available for each `credentialType` value:
	CredentialConfig map[string]interface{} `pulumi:"credentialConfig"`
	// Specifies the type of credential that
	// will be generated for the role. Options include: `password`, `rsaPrivateKey`, `clientCertificate`.
	// See the plugin's API page for credential types supported by individual databases.
	CredentialType *string `pulumi:"credentialType"`
	// The unique name of the database connection to use for
	// the role.
	DbName *string `pulumi:"dbName"`
	// The default number of seconds for leases for this
	// role.
	DefaultTtl *int `pulumi:"defaultTtl"`
	// The maximum number of seconds for leases for this
	// role.
	MaxTtl *int `pulumi:"maxTtl"`
	// A unique name to give the role.
	Name *string `pulumi:"name"`
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured namespace.
	// *Available only for Vault Enterprise*.
	Namespace *string `pulumi:"namespace"`
	// The database statements to execute when
	// renewing a user.
	RenewStatements []string `pulumi:"renewStatements"`
	// The database statements to execute when
	// revoking a user.
	RevocationStatements []string `pulumi:"revocationStatements"`
	// The database statements to execute when
	// rolling back creation due to an error.
	RollbackStatements []string `pulumi:"rollbackStatements"`
}

type SecretBackendRoleState struct {
	// The unique name of the Vault mount to configure.
	Backend pulumi.StringPtrInput
	// The database statements to execute when
	// creating a user.
	CreationStatements pulumi.StringArrayInput
	// Specifies the configuration
	// for the given `credentialType`.
	//
	// The following options are available for each `credentialType` value:
	CredentialConfig pulumi.MapInput
	// Specifies the type of credential that
	// will be generated for the role. Options include: `password`, `rsaPrivateKey`, `clientCertificate`.
	// See the plugin's API page for credential types supported by individual databases.
	CredentialType pulumi.StringPtrInput
	// The unique name of the database connection to use for
	// the role.
	DbName pulumi.StringPtrInput
	// The default number of seconds for leases for this
	// role.
	DefaultTtl pulumi.IntPtrInput
	// The maximum number of seconds for leases for this
	// role.
	MaxTtl pulumi.IntPtrInput
	// A unique name to give the role.
	Name pulumi.StringPtrInput
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured namespace.
	// *Available only for Vault Enterprise*.
	Namespace pulumi.StringPtrInput
	// The database statements to execute when
	// renewing a user.
	RenewStatements pulumi.StringArrayInput
	// The database statements to execute when
	// revoking a user.
	RevocationStatements pulumi.StringArrayInput
	// The database statements to execute when
	// rolling back creation due to an error.
	RollbackStatements pulumi.StringArrayInput
}

func (SecretBackendRoleState) ElementType() reflect.Type {
	return reflect.TypeOf((*secretBackendRoleState)(nil)).Elem()
}

type secretBackendRoleArgs struct {
	// The unique name of the Vault mount to configure.
	Backend string `pulumi:"backend"`
	// The database statements to execute when
	// creating a user.
	CreationStatements []string `pulumi:"creationStatements"`
	// Specifies the configuration
	// for the given `credentialType`.
	//
	// The following options are available for each `credentialType` value:
	CredentialConfig map[string]interface{} `pulumi:"credentialConfig"`
	// Specifies the type of credential that
	// will be generated for the role. Options include: `password`, `rsaPrivateKey`, `clientCertificate`.
	// See the plugin's API page for credential types supported by individual databases.
	CredentialType *string `pulumi:"credentialType"`
	// The unique name of the database connection to use for
	// the role.
	DbName string `pulumi:"dbName"`
	// The default number of seconds for leases for this
	// role.
	DefaultTtl *int `pulumi:"defaultTtl"`
	// The maximum number of seconds for leases for this
	// role.
	MaxTtl *int `pulumi:"maxTtl"`
	// A unique name to give the role.
	Name *string `pulumi:"name"`
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured namespace.
	// *Available only for Vault Enterprise*.
	Namespace *string `pulumi:"namespace"`
	// The database statements to execute when
	// renewing a user.
	RenewStatements []string `pulumi:"renewStatements"`
	// The database statements to execute when
	// revoking a user.
	RevocationStatements []string `pulumi:"revocationStatements"`
	// The database statements to execute when
	// rolling back creation due to an error.
	RollbackStatements []string `pulumi:"rollbackStatements"`
}

// The set of arguments for constructing a SecretBackendRole resource.
type SecretBackendRoleArgs struct {
	// The unique name of the Vault mount to configure.
	Backend pulumi.StringInput
	// The database statements to execute when
	// creating a user.
	CreationStatements pulumi.StringArrayInput
	// Specifies the configuration
	// for the given `credentialType`.
	//
	// The following options are available for each `credentialType` value:
	CredentialConfig pulumi.MapInput
	// Specifies the type of credential that
	// will be generated for the role. Options include: `password`, `rsaPrivateKey`, `clientCertificate`.
	// See the plugin's API page for credential types supported by individual databases.
	CredentialType pulumi.StringPtrInput
	// The unique name of the database connection to use for
	// the role.
	DbName pulumi.StringInput
	// The default number of seconds for leases for this
	// role.
	DefaultTtl pulumi.IntPtrInput
	// The maximum number of seconds for leases for this
	// role.
	MaxTtl pulumi.IntPtrInput
	// A unique name to give the role.
	Name pulumi.StringPtrInput
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured namespace.
	// *Available only for Vault Enterprise*.
	Namespace pulumi.StringPtrInput
	// The database statements to execute when
	// renewing a user.
	RenewStatements pulumi.StringArrayInput
	// The database statements to execute when
	// revoking a user.
	RevocationStatements pulumi.StringArrayInput
	// The database statements to execute when
	// rolling back creation due to an error.
	RollbackStatements pulumi.StringArrayInput
}

func (SecretBackendRoleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*secretBackendRoleArgs)(nil)).Elem()
}

type SecretBackendRoleInput interface {
	pulumi.Input

	ToSecretBackendRoleOutput() SecretBackendRoleOutput
	ToSecretBackendRoleOutputWithContext(ctx context.Context) SecretBackendRoleOutput
}

func (*SecretBackendRole) ElementType() reflect.Type {
	return reflect.TypeOf((**SecretBackendRole)(nil)).Elem()
}

func (i *SecretBackendRole) ToSecretBackendRoleOutput() SecretBackendRoleOutput {
	return i.ToSecretBackendRoleOutputWithContext(context.Background())
}

func (i *SecretBackendRole) ToSecretBackendRoleOutputWithContext(ctx context.Context) SecretBackendRoleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecretBackendRoleOutput)
}

// SecretBackendRoleArrayInput is an input type that accepts SecretBackendRoleArray and SecretBackendRoleArrayOutput values.
// You can construct a concrete instance of `SecretBackendRoleArrayInput` via:
//
//	SecretBackendRoleArray{ SecretBackendRoleArgs{...} }
type SecretBackendRoleArrayInput interface {
	pulumi.Input

	ToSecretBackendRoleArrayOutput() SecretBackendRoleArrayOutput
	ToSecretBackendRoleArrayOutputWithContext(context.Context) SecretBackendRoleArrayOutput
}

type SecretBackendRoleArray []SecretBackendRoleInput

func (SecretBackendRoleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SecretBackendRole)(nil)).Elem()
}

func (i SecretBackendRoleArray) ToSecretBackendRoleArrayOutput() SecretBackendRoleArrayOutput {
	return i.ToSecretBackendRoleArrayOutputWithContext(context.Background())
}

func (i SecretBackendRoleArray) ToSecretBackendRoleArrayOutputWithContext(ctx context.Context) SecretBackendRoleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecretBackendRoleArrayOutput)
}

// SecretBackendRoleMapInput is an input type that accepts SecretBackendRoleMap and SecretBackendRoleMapOutput values.
// You can construct a concrete instance of `SecretBackendRoleMapInput` via:
//
//	SecretBackendRoleMap{ "key": SecretBackendRoleArgs{...} }
type SecretBackendRoleMapInput interface {
	pulumi.Input

	ToSecretBackendRoleMapOutput() SecretBackendRoleMapOutput
	ToSecretBackendRoleMapOutputWithContext(context.Context) SecretBackendRoleMapOutput
}

type SecretBackendRoleMap map[string]SecretBackendRoleInput

func (SecretBackendRoleMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SecretBackendRole)(nil)).Elem()
}

func (i SecretBackendRoleMap) ToSecretBackendRoleMapOutput() SecretBackendRoleMapOutput {
	return i.ToSecretBackendRoleMapOutputWithContext(context.Background())
}

func (i SecretBackendRoleMap) ToSecretBackendRoleMapOutputWithContext(ctx context.Context) SecretBackendRoleMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecretBackendRoleMapOutput)
}

type SecretBackendRoleOutput struct{ *pulumi.OutputState }

func (SecretBackendRoleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SecretBackendRole)(nil)).Elem()
}

func (o SecretBackendRoleOutput) ToSecretBackendRoleOutput() SecretBackendRoleOutput {
	return o
}

func (o SecretBackendRoleOutput) ToSecretBackendRoleOutputWithContext(ctx context.Context) SecretBackendRoleOutput {
	return o
}

// The unique name of the Vault mount to configure.
func (o SecretBackendRoleOutput) Backend() pulumi.StringOutput {
	return o.ApplyT(func(v *SecretBackendRole) pulumi.StringOutput { return v.Backend }).(pulumi.StringOutput)
}

// The database statements to execute when
// creating a user.
func (o SecretBackendRoleOutput) CreationStatements() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *SecretBackendRole) pulumi.StringArrayOutput { return v.CreationStatements }).(pulumi.StringArrayOutput)
}

// Specifies the configuration
// for the given `credentialType`.
//
// The following options are available for each `credentialType` value:
func (o SecretBackendRoleOutput) CredentialConfig() pulumi.MapOutput {
	return o.ApplyT(func(v *SecretBackendRole) pulumi.MapOutput { return v.CredentialConfig }).(pulumi.MapOutput)
}

// Specifies the type of credential that
// will be generated for the role. Options include: `password`, `rsaPrivateKey`, `clientCertificate`.
// See the plugin's API page for credential types supported by individual databases.
func (o SecretBackendRoleOutput) CredentialType() pulumi.StringOutput {
	return o.ApplyT(func(v *SecretBackendRole) pulumi.StringOutput { return v.CredentialType }).(pulumi.StringOutput)
}

// The unique name of the database connection to use for
// the role.
func (o SecretBackendRoleOutput) DbName() pulumi.StringOutput {
	return o.ApplyT(func(v *SecretBackendRole) pulumi.StringOutput { return v.DbName }).(pulumi.StringOutput)
}

// The default number of seconds for leases for this
// role.
func (o SecretBackendRoleOutput) DefaultTtl() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SecretBackendRole) pulumi.IntPtrOutput { return v.DefaultTtl }).(pulumi.IntPtrOutput)
}

// The maximum number of seconds for leases for this
// role.
func (o SecretBackendRoleOutput) MaxTtl() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SecretBackendRole) pulumi.IntPtrOutput { return v.MaxTtl }).(pulumi.IntPtrOutput)
}

// A unique name to give the role.
func (o SecretBackendRoleOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SecretBackendRole) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The namespace to provision the resource in.
// The value should not contain leading or trailing forward slashes.
// The `namespace` is always relative to the provider's configured namespace.
// *Available only for Vault Enterprise*.
func (o SecretBackendRoleOutput) Namespace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecretBackendRole) pulumi.StringPtrOutput { return v.Namespace }).(pulumi.StringPtrOutput)
}

// The database statements to execute when
// renewing a user.
func (o SecretBackendRoleOutput) RenewStatements() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *SecretBackendRole) pulumi.StringArrayOutput { return v.RenewStatements }).(pulumi.StringArrayOutput)
}

// The database statements to execute when
// revoking a user.
func (o SecretBackendRoleOutput) RevocationStatements() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *SecretBackendRole) pulumi.StringArrayOutput { return v.RevocationStatements }).(pulumi.StringArrayOutput)
}

// The database statements to execute when
// rolling back creation due to an error.
func (o SecretBackendRoleOutput) RollbackStatements() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *SecretBackendRole) pulumi.StringArrayOutput { return v.RollbackStatements }).(pulumi.StringArrayOutput)
}

type SecretBackendRoleArrayOutput struct{ *pulumi.OutputState }

func (SecretBackendRoleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SecretBackendRole)(nil)).Elem()
}

func (o SecretBackendRoleArrayOutput) ToSecretBackendRoleArrayOutput() SecretBackendRoleArrayOutput {
	return o
}

func (o SecretBackendRoleArrayOutput) ToSecretBackendRoleArrayOutputWithContext(ctx context.Context) SecretBackendRoleArrayOutput {
	return o
}

func (o SecretBackendRoleArrayOutput) Index(i pulumi.IntInput) SecretBackendRoleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SecretBackendRole {
		return vs[0].([]*SecretBackendRole)[vs[1].(int)]
	}).(SecretBackendRoleOutput)
}

type SecretBackendRoleMapOutput struct{ *pulumi.OutputState }

func (SecretBackendRoleMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SecretBackendRole)(nil)).Elem()
}

func (o SecretBackendRoleMapOutput) ToSecretBackendRoleMapOutput() SecretBackendRoleMapOutput {
	return o
}

func (o SecretBackendRoleMapOutput) ToSecretBackendRoleMapOutputWithContext(ctx context.Context) SecretBackendRoleMapOutput {
	return o
}

func (o SecretBackendRoleMapOutput) MapIndex(k pulumi.StringInput) SecretBackendRoleOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SecretBackendRole {
		return vs[0].(map[string]*SecretBackendRole)[vs[1].(string)]
	}).(SecretBackendRoleOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SecretBackendRoleInput)(nil)).Elem(), &SecretBackendRole{})
	pulumi.RegisterInputType(reflect.TypeOf((*SecretBackendRoleArrayInput)(nil)).Elem(), SecretBackendRoleArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SecretBackendRoleMapInput)(nil)).Elem(), SecretBackendRoleMap{})
	pulumi.RegisterOutputType(SecretBackendRoleOutput{})
	pulumi.RegisterOutputType(SecretBackendRoleArrayOutput{})
	pulumi.RegisterOutputType(SecretBackendRoleMapOutput{})
}
