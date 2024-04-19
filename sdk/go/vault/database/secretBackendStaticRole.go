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

// Creates a Database Secret Backend static role in Vault. Database secret backend
// static roles can be used to manage 1-to-1 mapping of a Vault Role to a user in a
// database for the database.
//
// ## Example Usage
//
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
//				AllowedRoles: pulumi.StringArray{
//					pulumi.String("*"),
//				},
//				Postgresql: &database.SecretBackendConnectionPostgresqlArgs{
//					ConnectionUrl: pulumi.String("postgres://username:password@host:port/database"),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			// configure a static role with period-based rotations
//			_, err = database.NewSecretBackendStaticRole(ctx, "periodRole", &database.SecretBackendStaticRoleArgs{
//				Backend:        db.Path,
//				DbName:         postgres.Name,
//				Username:       pulumi.String("example"),
//				RotationPeriod: pulumi.Int(3600),
//				RotationStatements: pulumi.StringArray{
//					pulumi.String("ALTER USER \"{{name}}\" WITH PASSWORD '{{password}}';"),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			// configure a static role with schedule-based rotations
//			_, err = database.NewSecretBackendStaticRole(ctx, "scheduleRole", &database.SecretBackendStaticRoleArgs{
//				Backend:          db.Path,
//				DbName:           postgres.Name,
//				Username:         pulumi.String("example"),
//				RotationSchedule: pulumi.String("0 0 * * SAT"),
//				RotationWindow:   pulumi.Int(172800),
//				RotationStatements: pulumi.StringArray{
//					pulumi.String("ALTER USER \"{{name}}\" WITH PASSWORD '{{password}}';"),
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
//
// ## Import
//
// Database secret backend static roles can be imported using the `backend`, `/static-roles/`, and the `name` e.g.
//
// ```sh
// $ pulumi import vault:database/secretBackendStaticRole:SecretBackendStaticRole example postgres/static-roles/my-role
// ```
type SecretBackendStaticRole struct {
	pulumi.CustomResourceState

	// The unique name of the Vault mount to configure.
	Backend pulumi.StringOutput `pulumi:"backend"`
	// The unique name of the database connection to use for the static role.
	DbName pulumi.StringOutput `pulumi:"dbName"`
	// A unique name to give the static role.
	Name pulumi.StringOutput `pulumi:"name"`
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured namespace.
	// *Available only for Vault Enterprise*.
	Namespace pulumi.StringPtrOutput `pulumi:"namespace"`
	// The amount of time Vault should wait before rotating the password, in seconds.
	// Mutually exclusive with `rotationSchedule`.
	RotationPeriod pulumi.IntPtrOutput `pulumi:"rotationPeriod"`
	// A cron-style string that will define the schedule on which rotations should occur.
	// Mutually exclusive with `rotationPeriod`.
	//
	// **Warning**: The `rotationPeriod` and `rotationSchedule` fields are
	// mutually exclusive. One of them must be set but not both.
	RotationSchedule pulumi.StringPtrOutput `pulumi:"rotationSchedule"`
	// Database statements to execute to rotate the password for the configured database user.
	RotationStatements pulumi.StringArrayOutput `pulumi:"rotationStatements"`
	// The amount of time, in seconds, in which rotations are allowed to occur starting
	// from a given `rotationSchedule`.
	RotationWindow pulumi.IntPtrOutput `pulumi:"rotationWindow"`
	// The database username that this static role corresponds to.
	Username pulumi.StringOutput `pulumi:"username"`
}

// NewSecretBackendStaticRole registers a new resource with the given unique name, arguments, and options.
func NewSecretBackendStaticRole(ctx *pulumi.Context,
	name string, args *SecretBackendStaticRoleArgs, opts ...pulumi.ResourceOption) (*SecretBackendStaticRole, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Backend == nil {
		return nil, errors.New("invalid value for required argument 'Backend'")
	}
	if args.DbName == nil {
		return nil, errors.New("invalid value for required argument 'DbName'")
	}
	if args.Username == nil {
		return nil, errors.New("invalid value for required argument 'Username'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource SecretBackendStaticRole
	err := ctx.RegisterResource("vault:database/secretBackendStaticRole:SecretBackendStaticRole", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSecretBackendStaticRole gets an existing SecretBackendStaticRole resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSecretBackendStaticRole(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SecretBackendStaticRoleState, opts ...pulumi.ResourceOption) (*SecretBackendStaticRole, error) {
	var resource SecretBackendStaticRole
	err := ctx.ReadResource("vault:database/secretBackendStaticRole:SecretBackendStaticRole", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SecretBackendStaticRole resources.
type secretBackendStaticRoleState struct {
	// The unique name of the Vault mount to configure.
	Backend *string `pulumi:"backend"`
	// The unique name of the database connection to use for the static role.
	DbName *string `pulumi:"dbName"`
	// A unique name to give the static role.
	Name *string `pulumi:"name"`
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured namespace.
	// *Available only for Vault Enterprise*.
	Namespace *string `pulumi:"namespace"`
	// The amount of time Vault should wait before rotating the password, in seconds.
	// Mutually exclusive with `rotationSchedule`.
	RotationPeriod *int `pulumi:"rotationPeriod"`
	// A cron-style string that will define the schedule on which rotations should occur.
	// Mutually exclusive with `rotationPeriod`.
	//
	// **Warning**: The `rotationPeriod` and `rotationSchedule` fields are
	// mutually exclusive. One of them must be set but not both.
	RotationSchedule *string `pulumi:"rotationSchedule"`
	// Database statements to execute to rotate the password for the configured database user.
	RotationStatements []string `pulumi:"rotationStatements"`
	// The amount of time, in seconds, in which rotations are allowed to occur starting
	// from a given `rotationSchedule`.
	RotationWindow *int `pulumi:"rotationWindow"`
	// The database username that this static role corresponds to.
	Username *string `pulumi:"username"`
}

type SecretBackendStaticRoleState struct {
	// The unique name of the Vault mount to configure.
	Backend pulumi.StringPtrInput
	// The unique name of the database connection to use for the static role.
	DbName pulumi.StringPtrInput
	// A unique name to give the static role.
	Name pulumi.StringPtrInput
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured namespace.
	// *Available only for Vault Enterprise*.
	Namespace pulumi.StringPtrInput
	// The amount of time Vault should wait before rotating the password, in seconds.
	// Mutually exclusive with `rotationSchedule`.
	RotationPeriod pulumi.IntPtrInput
	// A cron-style string that will define the schedule on which rotations should occur.
	// Mutually exclusive with `rotationPeriod`.
	//
	// **Warning**: The `rotationPeriod` and `rotationSchedule` fields are
	// mutually exclusive. One of them must be set but not both.
	RotationSchedule pulumi.StringPtrInput
	// Database statements to execute to rotate the password for the configured database user.
	RotationStatements pulumi.StringArrayInput
	// The amount of time, in seconds, in which rotations are allowed to occur starting
	// from a given `rotationSchedule`.
	RotationWindow pulumi.IntPtrInput
	// The database username that this static role corresponds to.
	Username pulumi.StringPtrInput
}

func (SecretBackendStaticRoleState) ElementType() reflect.Type {
	return reflect.TypeOf((*secretBackendStaticRoleState)(nil)).Elem()
}

type secretBackendStaticRoleArgs struct {
	// The unique name of the Vault mount to configure.
	Backend string `pulumi:"backend"`
	// The unique name of the database connection to use for the static role.
	DbName string `pulumi:"dbName"`
	// A unique name to give the static role.
	Name *string `pulumi:"name"`
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured namespace.
	// *Available only for Vault Enterprise*.
	Namespace *string `pulumi:"namespace"`
	// The amount of time Vault should wait before rotating the password, in seconds.
	// Mutually exclusive with `rotationSchedule`.
	RotationPeriod *int `pulumi:"rotationPeriod"`
	// A cron-style string that will define the schedule on which rotations should occur.
	// Mutually exclusive with `rotationPeriod`.
	//
	// **Warning**: The `rotationPeriod` and `rotationSchedule` fields are
	// mutually exclusive. One of them must be set but not both.
	RotationSchedule *string `pulumi:"rotationSchedule"`
	// Database statements to execute to rotate the password for the configured database user.
	RotationStatements []string `pulumi:"rotationStatements"`
	// The amount of time, in seconds, in which rotations are allowed to occur starting
	// from a given `rotationSchedule`.
	RotationWindow *int `pulumi:"rotationWindow"`
	// The database username that this static role corresponds to.
	Username string `pulumi:"username"`
}

// The set of arguments for constructing a SecretBackendStaticRole resource.
type SecretBackendStaticRoleArgs struct {
	// The unique name of the Vault mount to configure.
	Backend pulumi.StringInput
	// The unique name of the database connection to use for the static role.
	DbName pulumi.StringInput
	// A unique name to give the static role.
	Name pulumi.StringPtrInput
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured namespace.
	// *Available only for Vault Enterprise*.
	Namespace pulumi.StringPtrInput
	// The amount of time Vault should wait before rotating the password, in seconds.
	// Mutually exclusive with `rotationSchedule`.
	RotationPeriod pulumi.IntPtrInput
	// A cron-style string that will define the schedule on which rotations should occur.
	// Mutually exclusive with `rotationPeriod`.
	//
	// **Warning**: The `rotationPeriod` and `rotationSchedule` fields are
	// mutually exclusive. One of them must be set but not both.
	RotationSchedule pulumi.StringPtrInput
	// Database statements to execute to rotate the password for the configured database user.
	RotationStatements pulumi.StringArrayInput
	// The amount of time, in seconds, in which rotations are allowed to occur starting
	// from a given `rotationSchedule`.
	RotationWindow pulumi.IntPtrInput
	// The database username that this static role corresponds to.
	Username pulumi.StringInput
}

func (SecretBackendStaticRoleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*secretBackendStaticRoleArgs)(nil)).Elem()
}

type SecretBackendStaticRoleInput interface {
	pulumi.Input

	ToSecretBackendStaticRoleOutput() SecretBackendStaticRoleOutput
	ToSecretBackendStaticRoleOutputWithContext(ctx context.Context) SecretBackendStaticRoleOutput
}

func (*SecretBackendStaticRole) ElementType() reflect.Type {
	return reflect.TypeOf((**SecretBackendStaticRole)(nil)).Elem()
}

func (i *SecretBackendStaticRole) ToSecretBackendStaticRoleOutput() SecretBackendStaticRoleOutput {
	return i.ToSecretBackendStaticRoleOutputWithContext(context.Background())
}

func (i *SecretBackendStaticRole) ToSecretBackendStaticRoleOutputWithContext(ctx context.Context) SecretBackendStaticRoleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecretBackendStaticRoleOutput)
}

// SecretBackendStaticRoleArrayInput is an input type that accepts SecretBackendStaticRoleArray and SecretBackendStaticRoleArrayOutput values.
// You can construct a concrete instance of `SecretBackendStaticRoleArrayInput` via:
//
//	SecretBackendStaticRoleArray{ SecretBackendStaticRoleArgs{...} }
type SecretBackendStaticRoleArrayInput interface {
	pulumi.Input

	ToSecretBackendStaticRoleArrayOutput() SecretBackendStaticRoleArrayOutput
	ToSecretBackendStaticRoleArrayOutputWithContext(context.Context) SecretBackendStaticRoleArrayOutput
}

type SecretBackendStaticRoleArray []SecretBackendStaticRoleInput

func (SecretBackendStaticRoleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SecretBackendStaticRole)(nil)).Elem()
}

func (i SecretBackendStaticRoleArray) ToSecretBackendStaticRoleArrayOutput() SecretBackendStaticRoleArrayOutput {
	return i.ToSecretBackendStaticRoleArrayOutputWithContext(context.Background())
}

func (i SecretBackendStaticRoleArray) ToSecretBackendStaticRoleArrayOutputWithContext(ctx context.Context) SecretBackendStaticRoleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecretBackendStaticRoleArrayOutput)
}

// SecretBackendStaticRoleMapInput is an input type that accepts SecretBackendStaticRoleMap and SecretBackendStaticRoleMapOutput values.
// You can construct a concrete instance of `SecretBackendStaticRoleMapInput` via:
//
//	SecretBackendStaticRoleMap{ "key": SecretBackendStaticRoleArgs{...} }
type SecretBackendStaticRoleMapInput interface {
	pulumi.Input

	ToSecretBackendStaticRoleMapOutput() SecretBackendStaticRoleMapOutput
	ToSecretBackendStaticRoleMapOutputWithContext(context.Context) SecretBackendStaticRoleMapOutput
}

type SecretBackendStaticRoleMap map[string]SecretBackendStaticRoleInput

func (SecretBackendStaticRoleMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SecretBackendStaticRole)(nil)).Elem()
}

func (i SecretBackendStaticRoleMap) ToSecretBackendStaticRoleMapOutput() SecretBackendStaticRoleMapOutput {
	return i.ToSecretBackendStaticRoleMapOutputWithContext(context.Background())
}

func (i SecretBackendStaticRoleMap) ToSecretBackendStaticRoleMapOutputWithContext(ctx context.Context) SecretBackendStaticRoleMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecretBackendStaticRoleMapOutput)
}

type SecretBackendStaticRoleOutput struct{ *pulumi.OutputState }

func (SecretBackendStaticRoleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SecretBackendStaticRole)(nil)).Elem()
}

func (o SecretBackendStaticRoleOutput) ToSecretBackendStaticRoleOutput() SecretBackendStaticRoleOutput {
	return o
}

func (o SecretBackendStaticRoleOutput) ToSecretBackendStaticRoleOutputWithContext(ctx context.Context) SecretBackendStaticRoleOutput {
	return o
}

// The unique name of the Vault mount to configure.
func (o SecretBackendStaticRoleOutput) Backend() pulumi.StringOutput {
	return o.ApplyT(func(v *SecretBackendStaticRole) pulumi.StringOutput { return v.Backend }).(pulumi.StringOutput)
}

// The unique name of the database connection to use for the static role.
func (o SecretBackendStaticRoleOutput) DbName() pulumi.StringOutput {
	return o.ApplyT(func(v *SecretBackendStaticRole) pulumi.StringOutput { return v.DbName }).(pulumi.StringOutput)
}

// A unique name to give the static role.
func (o SecretBackendStaticRoleOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SecretBackendStaticRole) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The namespace to provision the resource in.
// The value should not contain leading or trailing forward slashes.
// The `namespace` is always relative to the provider's configured namespace.
// *Available only for Vault Enterprise*.
func (o SecretBackendStaticRoleOutput) Namespace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecretBackendStaticRole) pulumi.StringPtrOutput { return v.Namespace }).(pulumi.StringPtrOutput)
}

// The amount of time Vault should wait before rotating the password, in seconds.
// Mutually exclusive with `rotationSchedule`.
func (o SecretBackendStaticRoleOutput) RotationPeriod() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SecretBackendStaticRole) pulumi.IntPtrOutput { return v.RotationPeriod }).(pulumi.IntPtrOutput)
}

// A cron-style string that will define the schedule on which rotations should occur.
// Mutually exclusive with `rotationPeriod`.
//
// **Warning**: The `rotationPeriod` and `rotationSchedule` fields are
// mutually exclusive. One of them must be set but not both.
func (o SecretBackendStaticRoleOutput) RotationSchedule() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecretBackendStaticRole) pulumi.StringPtrOutput { return v.RotationSchedule }).(pulumi.StringPtrOutput)
}

// Database statements to execute to rotate the password for the configured database user.
func (o SecretBackendStaticRoleOutput) RotationStatements() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *SecretBackendStaticRole) pulumi.StringArrayOutput { return v.RotationStatements }).(pulumi.StringArrayOutput)
}

// The amount of time, in seconds, in which rotations are allowed to occur starting
// from a given `rotationSchedule`.
func (o SecretBackendStaticRoleOutput) RotationWindow() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SecretBackendStaticRole) pulumi.IntPtrOutput { return v.RotationWindow }).(pulumi.IntPtrOutput)
}

// The database username that this static role corresponds to.
func (o SecretBackendStaticRoleOutput) Username() pulumi.StringOutput {
	return o.ApplyT(func(v *SecretBackendStaticRole) pulumi.StringOutput { return v.Username }).(pulumi.StringOutput)
}

type SecretBackendStaticRoleArrayOutput struct{ *pulumi.OutputState }

func (SecretBackendStaticRoleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SecretBackendStaticRole)(nil)).Elem()
}

func (o SecretBackendStaticRoleArrayOutput) ToSecretBackendStaticRoleArrayOutput() SecretBackendStaticRoleArrayOutput {
	return o
}

func (o SecretBackendStaticRoleArrayOutput) ToSecretBackendStaticRoleArrayOutputWithContext(ctx context.Context) SecretBackendStaticRoleArrayOutput {
	return o
}

func (o SecretBackendStaticRoleArrayOutput) Index(i pulumi.IntInput) SecretBackendStaticRoleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SecretBackendStaticRole {
		return vs[0].([]*SecretBackendStaticRole)[vs[1].(int)]
	}).(SecretBackendStaticRoleOutput)
}

type SecretBackendStaticRoleMapOutput struct{ *pulumi.OutputState }

func (SecretBackendStaticRoleMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SecretBackendStaticRole)(nil)).Elem()
}

func (o SecretBackendStaticRoleMapOutput) ToSecretBackendStaticRoleMapOutput() SecretBackendStaticRoleMapOutput {
	return o
}

func (o SecretBackendStaticRoleMapOutput) ToSecretBackendStaticRoleMapOutputWithContext(ctx context.Context) SecretBackendStaticRoleMapOutput {
	return o
}

func (o SecretBackendStaticRoleMapOutput) MapIndex(k pulumi.StringInput) SecretBackendStaticRoleOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SecretBackendStaticRole {
		return vs[0].(map[string]*SecretBackendStaticRole)[vs[1].(string)]
	}).(SecretBackendStaticRoleOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SecretBackendStaticRoleInput)(nil)).Elem(), &SecretBackendStaticRole{})
	pulumi.RegisterInputType(reflect.TypeOf((*SecretBackendStaticRoleArrayInput)(nil)).Elem(), SecretBackendStaticRoleArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SecretBackendStaticRoleMapInput)(nil)).Elem(), SecretBackendStaticRoleMap{})
	pulumi.RegisterOutputType(SecretBackendStaticRoleOutput{})
	pulumi.RegisterOutputType(SecretBackendStaticRoleArrayOutput{})
	pulumi.RegisterOutputType(SecretBackendStaticRoleMapOutput{})
}
