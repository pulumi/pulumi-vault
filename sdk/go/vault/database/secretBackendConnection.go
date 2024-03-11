// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package database

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-vault/sdk/v5/go/vault/internal"
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
//	"github.com/pulumi/pulumi-vault/sdk/v5/go/vault"
//	"github.com/pulumi/pulumi-vault/sdk/v5/go/vault/database"
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
//			_, err = database.NewSecretBackendConnection(ctx, "postgres", &database.SecretBackendConnectionArgs{
//				Backend: db.Path,
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
//			return nil
//		})
//	}
//
// ```
// <!--End PulumiCodeChooser -->
//
// ## Import
//
// Database secret backend connections can be imported using the `backend`, `/config/`, and the `name` e.g.
//
// ```sh
// $ pulumi import vault:database/secretBackendConnection:SecretBackendConnection example postgres/config/postgres
// ```
type SecretBackendConnection struct {
	pulumi.CustomResourceState

	// A list of roles that are allowed to use this
	// connection.
	AllowedRoles pulumi.StringArrayOutput `pulumi:"allowedRoles"`
	// The unique name of the Vault mount to configure.
	Backend pulumi.StringOutput `pulumi:"backend"`
	// A nested block containing configuration options for Cassandra connections.
	Cassandra SecretBackendConnectionCassandraPtrOutput `pulumi:"cassandra"`
	// A nested block containing configuration options for Couchbase connections.
	Couchbase SecretBackendConnectionCouchbasePtrOutput `pulumi:"couchbase"`
	// A map of sensitive data to pass to the endpoint. Useful for templated connection strings.
	Data pulumi.MapOutput `pulumi:"data"`
	// A nested block containing configuration options for Elasticsearch connections.
	Elasticsearch SecretBackendConnectionElasticsearchPtrOutput `pulumi:"elasticsearch"`
	// A nested block containing configuration options for SAP HanaDB connections.
	Hana SecretBackendConnectionHanaPtrOutput `pulumi:"hana"`
	// A nested block containing configuration options for InfluxDB connections.
	Influxdb SecretBackendConnectionInfluxdbPtrOutput `pulumi:"influxdb"`
	// A nested block containing configuration options for MongoDB connections.
	Mongodb SecretBackendConnectionMongodbPtrOutput `pulumi:"mongodb"`
	// A nested block containing configuration options for MongoDB Atlas connections.
	Mongodbatlas SecretBackendConnectionMongodbatlasPtrOutput `pulumi:"mongodbatlas"`
	// A nested block containing configuration options for MSSQL connections.
	Mssql SecretBackendConnectionMssqlPtrOutput `pulumi:"mssql"`
	// A nested block containing configuration options for MySQL connections.
	Mysql SecretBackendConnectionMysqlPtrOutput `pulumi:"mysql"`
	// A nested block containing configuration options for Aurora MySQL connections.
	MysqlAurora SecretBackendConnectionMysqlAuroraPtrOutput `pulumi:"mysqlAurora"`
	// A nested block containing configuration options for legacy MySQL connections.
	MysqlLegacy SecretBackendConnectionMysqlLegacyPtrOutput `pulumi:"mysqlLegacy"`
	// A nested block containing configuration options for RDS MySQL connections.
	MysqlRds SecretBackendConnectionMysqlRdsPtrOutput `pulumi:"mysqlRds"`
	// A unique name to give the database connection.
	Name pulumi.StringOutput `pulumi:"name"`
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured namespace.
	// *Available only for Vault Enterprise*.
	Namespace pulumi.StringPtrOutput `pulumi:"namespace"`
	// A nested block containing configuration options for Oracle connections.
	Oracle SecretBackendConnectionOraclePtrOutput `pulumi:"oracle"`
	// Specifies the name of the plugin to use.
	PluginName pulumi.StringOutput `pulumi:"pluginName"`
	// A nested block containing configuration options for PostgreSQL connections.
	Postgresql SecretBackendConnectionPostgresqlPtrOutput `pulumi:"postgresql"`
	// A nested block containing configuration options for Redis connections.
	Redis SecretBackendConnectionRedisPtrOutput `pulumi:"redis"`
	// A nested block containing configuration options for Redis ElastiCache connections.
	//
	// Exactly one of the nested blocks of configuration options must be supplied.
	RedisElasticache SecretBackendConnectionRedisElasticachePtrOutput `pulumi:"redisElasticache"`
	// Connection parameters for the redshift-database-plugin plugin.
	Redshift SecretBackendConnectionRedshiftPtrOutput `pulumi:"redshift"`
	// A list of database statements to be executed to rotate the root user's credentials.
	RootRotationStatements pulumi.StringArrayOutput `pulumi:"rootRotationStatements"`
	// A nested block containing configuration options for Snowflake connections.
	Snowflake SecretBackendConnectionSnowflakePtrOutput `pulumi:"snowflake"`
	// Whether the connection should be verified on
	// initial configuration or not.
	VerifyConnection pulumi.BoolPtrOutput `pulumi:"verifyConnection"`
}

// NewSecretBackendConnection registers a new resource with the given unique name, arguments, and options.
func NewSecretBackendConnection(ctx *pulumi.Context,
	name string, args *SecretBackendConnectionArgs, opts ...pulumi.ResourceOption) (*SecretBackendConnection, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Backend == nil {
		return nil, errors.New("invalid value for required argument 'Backend'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource SecretBackendConnection
	err := ctx.RegisterResource("vault:database/secretBackendConnection:SecretBackendConnection", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSecretBackendConnection gets an existing SecretBackendConnection resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSecretBackendConnection(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SecretBackendConnectionState, opts ...pulumi.ResourceOption) (*SecretBackendConnection, error) {
	var resource SecretBackendConnection
	err := ctx.ReadResource("vault:database/secretBackendConnection:SecretBackendConnection", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SecretBackendConnection resources.
type secretBackendConnectionState struct {
	// A list of roles that are allowed to use this
	// connection.
	AllowedRoles []string `pulumi:"allowedRoles"`
	// The unique name of the Vault mount to configure.
	Backend *string `pulumi:"backend"`
	// A nested block containing configuration options for Cassandra connections.
	Cassandra *SecretBackendConnectionCassandra `pulumi:"cassandra"`
	// A nested block containing configuration options for Couchbase connections.
	Couchbase *SecretBackendConnectionCouchbase `pulumi:"couchbase"`
	// A map of sensitive data to pass to the endpoint. Useful for templated connection strings.
	Data map[string]interface{} `pulumi:"data"`
	// A nested block containing configuration options for Elasticsearch connections.
	Elasticsearch *SecretBackendConnectionElasticsearch `pulumi:"elasticsearch"`
	// A nested block containing configuration options for SAP HanaDB connections.
	Hana *SecretBackendConnectionHana `pulumi:"hana"`
	// A nested block containing configuration options for InfluxDB connections.
	Influxdb *SecretBackendConnectionInfluxdb `pulumi:"influxdb"`
	// A nested block containing configuration options for MongoDB connections.
	Mongodb *SecretBackendConnectionMongodb `pulumi:"mongodb"`
	// A nested block containing configuration options for MongoDB Atlas connections.
	Mongodbatlas *SecretBackendConnectionMongodbatlas `pulumi:"mongodbatlas"`
	// A nested block containing configuration options for MSSQL connections.
	Mssql *SecretBackendConnectionMssql `pulumi:"mssql"`
	// A nested block containing configuration options for MySQL connections.
	Mysql *SecretBackendConnectionMysql `pulumi:"mysql"`
	// A nested block containing configuration options for Aurora MySQL connections.
	MysqlAurora *SecretBackendConnectionMysqlAurora `pulumi:"mysqlAurora"`
	// A nested block containing configuration options for legacy MySQL connections.
	MysqlLegacy *SecretBackendConnectionMysqlLegacy `pulumi:"mysqlLegacy"`
	// A nested block containing configuration options for RDS MySQL connections.
	MysqlRds *SecretBackendConnectionMysqlRds `pulumi:"mysqlRds"`
	// A unique name to give the database connection.
	Name *string `pulumi:"name"`
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured namespace.
	// *Available only for Vault Enterprise*.
	Namespace *string `pulumi:"namespace"`
	// A nested block containing configuration options for Oracle connections.
	Oracle *SecretBackendConnectionOracle `pulumi:"oracle"`
	// Specifies the name of the plugin to use.
	PluginName *string `pulumi:"pluginName"`
	// A nested block containing configuration options for PostgreSQL connections.
	Postgresql *SecretBackendConnectionPostgresql `pulumi:"postgresql"`
	// A nested block containing configuration options for Redis connections.
	Redis *SecretBackendConnectionRedis `pulumi:"redis"`
	// A nested block containing configuration options for Redis ElastiCache connections.
	//
	// Exactly one of the nested blocks of configuration options must be supplied.
	RedisElasticache *SecretBackendConnectionRedisElasticache `pulumi:"redisElasticache"`
	// Connection parameters for the redshift-database-plugin plugin.
	Redshift *SecretBackendConnectionRedshift `pulumi:"redshift"`
	// A list of database statements to be executed to rotate the root user's credentials.
	RootRotationStatements []string `pulumi:"rootRotationStatements"`
	// A nested block containing configuration options for Snowflake connections.
	Snowflake *SecretBackendConnectionSnowflake `pulumi:"snowflake"`
	// Whether the connection should be verified on
	// initial configuration or not.
	VerifyConnection *bool `pulumi:"verifyConnection"`
}

type SecretBackendConnectionState struct {
	// A list of roles that are allowed to use this
	// connection.
	AllowedRoles pulumi.StringArrayInput
	// The unique name of the Vault mount to configure.
	Backend pulumi.StringPtrInput
	// A nested block containing configuration options for Cassandra connections.
	Cassandra SecretBackendConnectionCassandraPtrInput
	// A nested block containing configuration options for Couchbase connections.
	Couchbase SecretBackendConnectionCouchbasePtrInput
	// A map of sensitive data to pass to the endpoint. Useful for templated connection strings.
	Data pulumi.MapInput
	// A nested block containing configuration options for Elasticsearch connections.
	Elasticsearch SecretBackendConnectionElasticsearchPtrInput
	// A nested block containing configuration options for SAP HanaDB connections.
	Hana SecretBackendConnectionHanaPtrInput
	// A nested block containing configuration options for InfluxDB connections.
	Influxdb SecretBackendConnectionInfluxdbPtrInput
	// A nested block containing configuration options for MongoDB connections.
	Mongodb SecretBackendConnectionMongodbPtrInput
	// A nested block containing configuration options for MongoDB Atlas connections.
	Mongodbatlas SecretBackendConnectionMongodbatlasPtrInput
	// A nested block containing configuration options for MSSQL connections.
	Mssql SecretBackendConnectionMssqlPtrInput
	// A nested block containing configuration options for MySQL connections.
	Mysql SecretBackendConnectionMysqlPtrInput
	// A nested block containing configuration options for Aurora MySQL connections.
	MysqlAurora SecretBackendConnectionMysqlAuroraPtrInput
	// A nested block containing configuration options for legacy MySQL connections.
	MysqlLegacy SecretBackendConnectionMysqlLegacyPtrInput
	// A nested block containing configuration options for RDS MySQL connections.
	MysqlRds SecretBackendConnectionMysqlRdsPtrInput
	// A unique name to give the database connection.
	Name pulumi.StringPtrInput
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured namespace.
	// *Available only for Vault Enterprise*.
	Namespace pulumi.StringPtrInput
	// A nested block containing configuration options for Oracle connections.
	Oracle SecretBackendConnectionOraclePtrInput
	// Specifies the name of the plugin to use.
	PluginName pulumi.StringPtrInput
	// A nested block containing configuration options for PostgreSQL connections.
	Postgresql SecretBackendConnectionPostgresqlPtrInput
	// A nested block containing configuration options for Redis connections.
	Redis SecretBackendConnectionRedisPtrInput
	// A nested block containing configuration options for Redis ElastiCache connections.
	//
	// Exactly one of the nested blocks of configuration options must be supplied.
	RedisElasticache SecretBackendConnectionRedisElasticachePtrInput
	// Connection parameters for the redshift-database-plugin plugin.
	Redshift SecretBackendConnectionRedshiftPtrInput
	// A list of database statements to be executed to rotate the root user's credentials.
	RootRotationStatements pulumi.StringArrayInput
	// A nested block containing configuration options for Snowflake connections.
	Snowflake SecretBackendConnectionSnowflakePtrInput
	// Whether the connection should be verified on
	// initial configuration or not.
	VerifyConnection pulumi.BoolPtrInput
}

func (SecretBackendConnectionState) ElementType() reflect.Type {
	return reflect.TypeOf((*secretBackendConnectionState)(nil)).Elem()
}

type secretBackendConnectionArgs struct {
	// A list of roles that are allowed to use this
	// connection.
	AllowedRoles []string `pulumi:"allowedRoles"`
	// The unique name of the Vault mount to configure.
	Backend string `pulumi:"backend"`
	// A nested block containing configuration options for Cassandra connections.
	Cassandra *SecretBackendConnectionCassandra `pulumi:"cassandra"`
	// A nested block containing configuration options for Couchbase connections.
	Couchbase *SecretBackendConnectionCouchbase `pulumi:"couchbase"`
	// A map of sensitive data to pass to the endpoint. Useful for templated connection strings.
	Data map[string]interface{} `pulumi:"data"`
	// A nested block containing configuration options for Elasticsearch connections.
	Elasticsearch *SecretBackendConnectionElasticsearch `pulumi:"elasticsearch"`
	// A nested block containing configuration options for SAP HanaDB connections.
	Hana *SecretBackendConnectionHana `pulumi:"hana"`
	// A nested block containing configuration options for InfluxDB connections.
	Influxdb *SecretBackendConnectionInfluxdb `pulumi:"influxdb"`
	// A nested block containing configuration options for MongoDB connections.
	Mongodb *SecretBackendConnectionMongodb `pulumi:"mongodb"`
	// A nested block containing configuration options for MongoDB Atlas connections.
	Mongodbatlas *SecretBackendConnectionMongodbatlas `pulumi:"mongodbatlas"`
	// A nested block containing configuration options for MSSQL connections.
	Mssql *SecretBackendConnectionMssql `pulumi:"mssql"`
	// A nested block containing configuration options for MySQL connections.
	Mysql *SecretBackendConnectionMysql `pulumi:"mysql"`
	// A nested block containing configuration options for Aurora MySQL connections.
	MysqlAurora *SecretBackendConnectionMysqlAurora `pulumi:"mysqlAurora"`
	// A nested block containing configuration options for legacy MySQL connections.
	MysqlLegacy *SecretBackendConnectionMysqlLegacy `pulumi:"mysqlLegacy"`
	// A nested block containing configuration options for RDS MySQL connections.
	MysqlRds *SecretBackendConnectionMysqlRds `pulumi:"mysqlRds"`
	// A unique name to give the database connection.
	Name *string `pulumi:"name"`
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured namespace.
	// *Available only for Vault Enterprise*.
	Namespace *string `pulumi:"namespace"`
	// A nested block containing configuration options for Oracle connections.
	Oracle *SecretBackendConnectionOracle `pulumi:"oracle"`
	// Specifies the name of the plugin to use.
	PluginName *string `pulumi:"pluginName"`
	// A nested block containing configuration options for PostgreSQL connections.
	Postgresql *SecretBackendConnectionPostgresql `pulumi:"postgresql"`
	// A nested block containing configuration options for Redis connections.
	Redis *SecretBackendConnectionRedis `pulumi:"redis"`
	// A nested block containing configuration options for Redis ElastiCache connections.
	//
	// Exactly one of the nested blocks of configuration options must be supplied.
	RedisElasticache *SecretBackendConnectionRedisElasticache `pulumi:"redisElasticache"`
	// Connection parameters for the redshift-database-plugin plugin.
	Redshift *SecretBackendConnectionRedshift `pulumi:"redshift"`
	// A list of database statements to be executed to rotate the root user's credentials.
	RootRotationStatements []string `pulumi:"rootRotationStatements"`
	// A nested block containing configuration options for Snowflake connections.
	Snowflake *SecretBackendConnectionSnowflake `pulumi:"snowflake"`
	// Whether the connection should be verified on
	// initial configuration or not.
	VerifyConnection *bool `pulumi:"verifyConnection"`
}

// The set of arguments for constructing a SecretBackendConnection resource.
type SecretBackendConnectionArgs struct {
	// A list of roles that are allowed to use this
	// connection.
	AllowedRoles pulumi.StringArrayInput
	// The unique name of the Vault mount to configure.
	Backend pulumi.StringInput
	// A nested block containing configuration options for Cassandra connections.
	Cassandra SecretBackendConnectionCassandraPtrInput
	// A nested block containing configuration options for Couchbase connections.
	Couchbase SecretBackendConnectionCouchbasePtrInput
	// A map of sensitive data to pass to the endpoint. Useful for templated connection strings.
	Data pulumi.MapInput
	// A nested block containing configuration options for Elasticsearch connections.
	Elasticsearch SecretBackendConnectionElasticsearchPtrInput
	// A nested block containing configuration options for SAP HanaDB connections.
	Hana SecretBackendConnectionHanaPtrInput
	// A nested block containing configuration options for InfluxDB connections.
	Influxdb SecretBackendConnectionInfluxdbPtrInput
	// A nested block containing configuration options for MongoDB connections.
	Mongodb SecretBackendConnectionMongodbPtrInput
	// A nested block containing configuration options for MongoDB Atlas connections.
	Mongodbatlas SecretBackendConnectionMongodbatlasPtrInput
	// A nested block containing configuration options for MSSQL connections.
	Mssql SecretBackendConnectionMssqlPtrInput
	// A nested block containing configuration options for MySQL connections.
	Mysql SecretBackendConnectionMysqlPtrInput
	// A nested block containing configuration options for Aurora MySQL connections.
	MysqlAurora SecretBackendConnectionMysqlAuroraPtrInput
	// A nested block containing configuration options for legacy MySQL connections.
	MysqlLegacy SecretBackendConnectionMysqlLegacyPtrInput
	// A nested block containing configuration options for RDS MySQL connections.
	MysqlRds SecretBackendConnectionMysqlRdsPtrInput
	// A unique name to give the database connection.
	Name pulumi.StringPtrInput
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured namespace.
	// *Available only for Vault Enterprise*.
	Namespace pulumi.StringPtrInput
	// A nested block containing configuration options for Oracle connections.
	Oracle SecretBackendConnectionOraclePtrInput
	// Specifies the name of the plugin to use.
	PluginName pulumi.StringPtrInput
	// A nested block containing configuration options for PostgreSQL connections.
	Postgresql SecretBackendConnectionPostgresqlPtrInput
	// A nested block containing configuration options for Redis connections.
	Redis SecretBackendConnectionRedisPtrInput
	// A nested block containing configuration options for Redis ElastiCache connections.
	//
	// Exactly one of the nested blocks of configuration options must be supplied.
	RedisElasticache SecretBackendConnectionRedisElasticachePtrInput
	// Connection parameters for the redshift-database-plugin plugin.
	Redshift SecretBackendConnectionRedshiftPtrInput
	// A list of database statements to be executed to rotate the root user's credentials.
	RootRotationStatements pulumi.StringArrayInput
	// A nested block containing configuration options for Snowflake connections.
	Snowflake SecretBackendConnectionSnowflakePtrInput
	// Whether the connection should be verified on
	// initial configuration or not.
	VerifyConnection pulumi.BoolPtrInput
}

func (SecretBackendConnectionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*secretBackendConnectionArgs)(nil)).Elem()
}

type SecretBackendConnectionInput interface {
	pulumi.Input

	ToSecretBackendConnectionOutput() SecretBackendConnectionOutput
	ToSecretBackendConnectionOutputWithContext(ctx context.Context) SecretBackendConnectionOutput
}

func (*SecretBackendConnection) ElementType() reflect.Type {
	return reflect.TypeOf((**SecretBackendConnection)(nil)).Elem()
}

func (i *SecretBackendConnection) ToSecretBackendConnectionOutput() SecretBackendConnectionOutput {
	return i.ToSecretBackendConnectionOutputWithContext(context.Background())
}

func (i *SecretBackendConnection) ToSecretBackendConnectionOutputWithContext(ctx context.Context) SecretBackendConnectionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecretBackendConnectionOutput)
}

// SecretBackendConnectionArrayInput is an input type that accepts SecretBackendConnectionArray and SecretBackendConnectionArrayOutput values.
// You can construct a concrete instance of `SecretBackendConnectionArrayInput` via:
//
//	SecretBackendConnectionArray{ SecretBackendConnectionArgs{...} }
type SecretBackendConnectionArrayInput interface {
	pulumi.Input

	ToSecretBackendConnectionArrayOutput() SecretBackendConnectionArrayOutput
	ToSecretBackendConnectionArrayOutputWithContext(context.Context) SecretBackendConnectionArrayOutput
}

type SecretBackendConnectionArray []SecretBackendConnectionInput

func (SecretBackendConnectionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SecretBackendConnection)(nil)).Elem()
}

func (i SecretBackendConnectionArray) ToSecretBackendConnectionArrayOutput() SecretBackendConnectionArrayOutput {
	return i.ToSecretBackendConnectionArrayOutputWithContext(context.Background())
}

func (i SecretBackendConnectionArray) ToSecretBackendConnectionArrayOutputWithContext(ctx context.Context) SecretBackendConnectionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecretBackendConnectionArrayOutput)
}

// SecretBackendConnectionMapInput is an input type that accepts SecretBackendConnectionMap and SecretBackendConnectionMapOutput values.
// You can construct a concrete instance of `SecretBackendConnectionMapInput` via:
//
//	SecretBackendConnectionMap{ "key": SecretBackendConnectionArgs{...} }
type SecretBackendConnectionMapInput interface {
	pulumi.Input

	ToSecretBackendConnectionMapOutput() SecretBackendConnectionMapOutput
	ToSecretBackendConnectionMapOutputWithContext(context.Context) SecretBackendConnectionMapOutput
}

type SecretBackendConnectionMap map[string]SecretBackendConnectionInput

func (SecretBackendConnectionMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SecretBackendConnection)(nil)).Elem()
}

func (i SecretBackendConnectionMap) ToSecretBackendConnectionMapOutput() SecretBackendConnectionMapOutput {
	return i.ToSecretBackendConnectionMapOutputWithContext(context.Background())
}

func (i SecretBackendConnectionMap) ToSecretBackendConnectionMapOutputWithContext(ctx context.Context) SecretBackendConnectionMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecretBackendConnectionMapOutput)
}

type SecretBackendConnectionOutput struct{ *pulumi.OutputState }

func (SecretBackendConnectionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SecretBackendConnection)(nil)).Elem()
}

func (o SecretBackendConnectionOutput) ToSecretBackendConnectionOutput() SecretBackendConnectionOutput {
	return o
}

func (o SecretBackendConnectionOutput) ToSecretBackendConnectionOutputWithContext(ctx context.Context) SecretBackendConnectionOutput {
	return o
}

// A list of roles that are allowed to use this
// connection.
func (o SecretBackendConnectionOutput) AllowedRoles() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *SecretBackendConnection) pulumi.StringArrayOutput { return v.AllowedRoles }).(pulumi.StringArrayOutput)
}

// The unique name of the Vault mount to configure.
func (o SecretBackendConnectionOutput) Backend() pulumi.StringOutput {
	return o.ApplyT(func(v *SecretBackendConnection) pulumi.StringOutput { return v.Backend }).(pulumi.StringOutput)
}

// A nested block containing configuration options for Cassandra connections.
func (o SecretBackendConnectionOutput) Cassandra() SecretBackendConnectionCassandraPtrOutput {
	return o.ApplyT(func(v *SecretBackendConnection) SecretBackendConnectionCassandraPtrOutput { return v.Cassandra }).(SecretBackendConnectionCassandraPtrOutput)
}

// A nested block containing configuration options for Couchbase connections.
func (o SecretBackendConnectionOutput) Couchbase() SecretBackendConnectionCouchbasePtrOutput {
	return o.ApplyT(func(v *SecretBackendConnection) SecretBackendConnectionCouchbasePtrOutput { return v.Couchbase }).(SecretBackendConnectionCouchbasePtrOutput)
}

// A map of sensitive data to pass to the endpoint. Useful for templated connection strings.
func (o SecretBackendConnectionOutput) Data() pulumi.MapOutput {
	return o.ApplyT(func(v *SecretBackendConnection) pulumi.MapOutput { return v.Data }).(pulumi.MapOutput)
}

// A nested block containing configuration options for Elasticsearch connections.
func (o SecretBackendConnectionOutput) Elasticsearch() SecretBackendConnectionElasticsearchPtrOutput {
	return o.ApplyT(func(v *SecretBackendConnection) SecretBackendConnectionElasticsearchPtrOutput { return v.Elasticsearch }).(SecretBackendConnectionElasticsearchPtrOutput)
}

// A nested block containing configuration options for SAP HanaDB connections.
func (o SecretBackendConnectionOutput) Hana() SecretBackendConnectionHanaPtrOutput {
	return o.ApplyT(func(v *SecretBackendConnection) SecretBackendConnectionHanaPtrOutput { return v.Hana }).(SecretBackendConnectionHanaPtrOutput)
}

// A nested block containing configuration options for InfluxDB connections.
func (o SecretBackendConnectionOutput) Influxdb() SecretBackendConnectionInfluxdbPtrOutput {
	return o.ApplyT(func(v *SecretBackendConnection) SecretBackendConnectionInfluxdbPtrOutput { return v.Influxdb }).(SecretBackendConnectionInfluxdbPtrOutput)
}

// A nested block containing configuration options for MongoDB connections.
func (o SecretBackendConnectionOutput) Mongodb() SecretBackendConnectionMongodbPtrOutput {
	return o.ApplyT(func(v *SecretBackendConnection) SecretBackendConnectionMongodbPtrOutput { return v.Mongodb }).(SecretBackendConnectionMongodbPtrOutput)
}

// A nested block containing configuration options for MongoDB Atlas connections.
func (o SecretBackendConnectionOutput) Mongodbatlas() SecretBackendConnectionMongodbatlasPtrOutput {
	return o.ApplyT(func(v *SecretBackendConnection) SecretBackendConnectionMongodbatlasPtrOutput { return v.Mongodbatlas }).(SecretBackendConnectionMongodbatlasPtrOutput)
}

// A nested block containing configuration options for MSSQL connections.
func (o SecretBackendConnectionOutput) Mssql() SecretBackendConnectionMssqlPtrOutput {
	return o.ApplyT(func(v *SecretBackendConnection) SecretBackendConnectionMssqlPtrOutput { return v.Mssql }).(SecretBackendConnectionMssqlPtrOutput)
}

// A nested block containing configuration options for MySQL connections.
func (o SecretBackendConnectionOutput) Mysql() SecretBackendConnectionMysqlPtrOutput {
	return o.ApplyT(func(v *SecretBackendConnection) SecretBackendConnectionMysqlPtrOutput { return v.Mysql }).(SecretBackendConnectionMysqlPtrOutput)
}

// A nested block containing configuration options for Aurora MySQL connections.
func (o SecretBackendConnectionOutput) MysqlAurora() SecretBackendConnectionMysqlAuroraPtrOutput {
	return o.ApplyT(func(v *SecretBackendConnection) SecretBackendConnectionMysqlAuroraPtrOutput { return v.MysqlAurora }).(SecretBackendConnectionMysqlAuroraPtrOutput)
}

// A nested block containing configuration options for legacy MySQL connections.
func (o SecretBackendConnectionOutput) MysqlLegacy() SecretBackendConnectionMysqlLegacyPtrOutput {
	return o.ApplyT(func(v *SecretBackendConnection) SecretBackendConnectionMysqlLegacyPtrOutput { return v.MysqlLegacy }).(SecretBackendConnectionMysqlLegacyPtrOutput)
}

// A nested block containing configuration options for RDS MySQL connections.
func (o SecretBackendConnectionOutput) MysqlRds() SecretBackendConnectionMysqlRdsPtrOutput {
	return o.ApplyT(func(v *SecretBackendConnection) SecretBackendConnectionMysqlRdsPtrOutput { return v.MysqlRds }).(SecretBackendConnectionMysqlRdsPtrOutput)
}

// A unique name to give the database connection.
func (o SecretBackendConnectionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SecretBackendConnection) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The namespace to provision the resource in.
// The value should not contain leading or trailing forward slashes.
// The `namespace` is always relative to the provider's configured namespace.
// *Available only for Vault Enterprise*.
func (o SecretBackendConnectionOutput) Namespace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecretBackendConnection) pulumi.StringPtrOutput { return v.Namespace }).(pulumi.StringPtrOutput)
}

// A nested block containing configuration options for Oracle connections.
func (o SecretBackendConnectionOutput) Oracle() SecretBackendConnectionOraclePtrOutput {
	return o.ApplyT(func(v *SecretBackendConnection) SecretBackendConnectionOraclePtrOutput { return v.Oracle }).(SecretBackendConnectionOraclePtrOutput)
}

// Specifies the name of the plugin to use.
func (o SecretBackendConnectionOutput) PluginName() pulumi.StringOutput {
	return o.ApplyT(func(v *SecretBackendConnection) pulumi.StringOutput { return v.PluginName }).(pulumi.StringOutput)
}

// A nested block containing configuration options for PostgreSQL connections.
func (o SecretBackendConnectionOutput) Postgresql() SecretBackendConnectionPostgresqlPtrOutput {
	return o.ApplyT(func(v *SecretBackendConnection) SecretBackendConnectionPostgresqlPtrOutput { return v.Postgresql }).(SecretBackendConnectionPostgresqlPtrOutput)
}

// A nested block containing configuration options for Redis connections.
func (o SecretBackendConnectionOutput) Redis() SecretBackendConnectionRedisPtrOutput {
	return o.ApplyT(func(v *SecretBackendConnection) SecretBackendConnectionRedisPtrOutput { return v.Redis }).(SecretBackendConnectionRedisPtrOutput)
}

// A nested block containing configuration options for Redis ElastiCache connections.
//
// Exactly one of the nested blocks of configuration options must be supplied.
func (o SecretBackendConnectionOutput) RedisElasticache() SecretBackendConnectionRedisElasticachePtrOutput {
	return o.ApplyT(func(v *SecretBackendConnection) SecretBackendConnectionRedisElasticachePtrOutput {
		return v.RedisElasticache
	}).(SecretBackendConnectionRedisElasticachePtrOutput)
}

// Connection parameters for the redshift-database-plugin plugin.
func (o SecretBackendConnectionOutput) Redshift() SecretBackendConnectionRedshiftPtrOutput {
	return o.ApplyT(func(v *SecretBackendConnection) SecretBackendConnectionRedshiftPtrOutput { return v.Redshift }).(SecretBackendConnectionRedshiftPtrOutput)
}

// A list of database statements to be executed to rotate the root user's credentials.
func (o SecretBackendConnectionOutput) RootRotationStatements() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *SecretBackendConnection) pulumi.StringArrayOutput { return v.RootRotationStatements }).(pulumi.StringArrayOutput)
}

// A nested block containing configuration options for Snowflake connections.
func (o SecretBackendConnectionOutput) Snowflake() SecretBackendConnectionSnowflakePtrOutput {
	return o.ApplyT(func(v *SecretBackendConnection) SecretBackendConnectionSnowflakePtrOutput { return v.Snowflake }).(SecretBackendConnectionSnowflakePtrOutput)
}

// Whether the connection should be verified on
// initial configuration or not.
func (o SecretBackendConnectionOutput) VerifyConnection() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SecretBackendConnection) pulumi.BoolPtrOutput { return v.VerifyConnection }).(pulumi.BoolPtrOutput)
}

type SecretBackendConnectionArrayOutput struct{ *pulumi.OutputState }

func (SecretBackendConnectionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SecretBackendConnection)(nil)).Elem()
}

func (o SecretBackendConnectionArrayOutput) ToSecretBackendConnectionArrayOutput() SecretBackendConnectionArrayOutput {
	return o
}

func (o SecretBackendConnectionArrayOutput) ToSecretBackendConnectionArrayOutputWithContext(ctx context.Context) SecretBackendConnectionArrayOutput {
	return o
}

func (o SecretBackendConnectionArrayOutput) Index(i pulumi.IntInput) SecretBackendConnectionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SecretBackendConnection {
		return vs[0].([]*SecretBackendConnection)[vs[1].(int)]
	}).(SecretBackendConnectionOutput)
}

type SecretBackendConnectionMapOutput struct{ *pulumi.OutputState }

func (SecretBackendConnectionMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SecretBackendConnection)(nil)).Elem()
}

func (o SecretBackendConnectionMapOutput) ToSecretBackendConnectionMapOutput() SecretBackendConnectionMapOutput {
	return o
}

func (o SecretBackendConnectionMapOutput) ToSecretBackendConnectionMapOutputWithContext(ctx context.Context) SecretBackendConnectionMapOutput {
	return o
}

func (o SecretBackendConnectionMapOutput) MapIndex(k pulumi.StringInput) SecretBackendConnectionOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SecretBackendConnection {
		return vs[0].(map[string]*SecretBackendConnection)[vs[1].(string)]
	}).(SecretBackendConnectionOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SecretBackendConnectionInput)(nil)).Elem(), &SecretBackendConnection{})
	pulumi.RegisterInputType(reflect.TypeOf((*SecretBackendConnectionArrayInput)(nil)).Elem(), SecretBackendConnectionArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SecretBackendConnectionMapInput)(nil)).Elem(), SecretBackendConnectionMap{})
	pulumi.RegisterOutputType(SecretBackendConnectionOutput{})
	pulumi.RegisterOutputType(SecretBackendConnectionArrayOutput{})
	pulumi.RegisterOutputType(SecretBackendConnectionMapOutput{})
}
