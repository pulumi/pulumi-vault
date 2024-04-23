// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vault from "@pulumi/vault";
 *
 * const db = new vault.database.SecretsMount("db", {
 *     path: "db",
 *     mssqls: [{
 *         name: "db1",
 *         username: "sa",
 *         password: "super_secret_1",
 *         connectionUrl: "sqlserver://{{username}}:{{password}}@127.0.0.1:1433",
 *         allowedRoles: ["dev1"],
 *     }],
 *     postgresqls: [{
 *         name: "db2",
 *         username: "postgres",
 *         password: "super_secret_2",
 *         connectionUrl: "postgresql://{{username}}:{{password}}@127.0.0.1:5432/postgres",
 *         verifyConnection: true,
 *         allowedRoles: ["dev2"],
 *     }],
 * });
 * const dev1 = new vault.database.SecretBackendRole("dev1", {
 *     name: "dev1",
 *     backend: db.path,
 *     dbName: db.mssqls.apply(mssqls => mssqls?.[0]?.name),
 *     creationStatements: [
 *         "CREATE LOGIN [{{name}}] WITH PASSWORD = '{{password}}';",
 *         "CREATE USER [{{name}}] FOR LOGIN [{{name}}];",
 *         "GRANT SELECT ON SCHEMA::dbo TO [{{name}}];",
 *     ],
 * });
 * const dev2 = new vault.database.SecretBackendRole("dev2", {
 *     name: "dev2",
 *     backend: db.path,
 *     dbName: db.postgresqls.apply(postgresqls => postgresqls?.[0]?.name),
 *     creationStatements: [
 *         "CREATE ROLE \"{{name}}\" WITH LOGIN PASSWORD '{{password}}' VALID UNTIL '{{expiration}}';",
 *         "GRANT SELECT ON ALL TABLES IN SCHEMA public TO \"{{name}}\";",
 *     ],
 * });
 * ```
 *
 * ## Import
 *
 * Database secret backend connections can be imported using the `path` e.g.
 *
 * ```sh
 * $ pulumi import vault:database/secretsMount:SecretsMount db db
 * ```
 */
export class SecretsMount extends pulumi.CustomResource {
    /**
     * Get an existing SecretsMount resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SecretsMountState, opts?: pulumi.CustomResourceOptions): SecretsMount {
        return new SecretsMount(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'vault:database/secretsMount:SecretsMount';

    /**
     * Returns true if the given object is an instance of SecretsMount.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SecretsMount {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SecretsMount.__pulumiType;
    }

    /**
     * Accessor of the mount
     */
    public /*out*/ readonly accessor!: pulumi.Output<string>;
    /**
     * Set of managed key registry entry names that the mount in question is allowed to access
     *
     * The following arguments are common to all database engines:
     */
    public readonly allowedManagedKeys!: pulumi.Output<string[] | undefined>;
    /**
     * Specifies the list of keys that will not be HMAC'd by audit devices in the request data object.
     */
    public readonly auditNonHmacRequestKeys!: pulumi.Output<string[]>;
    /**
     * Specifies the list of keys that will not be HMAC'd by audit devices in the response data object.
     */
    public readonly auditNonHmacResponseKeys!: pulumi.Output<string[]>;
    /**
     * A nested block containing configuration options for Cassandra connections.  
     * *See Configuration Options for more info*
     */
    public readonly cassandras!: pulumi.Output<outputs.database.SecretsMountCassandra[] | undefined>;
    /**
     * A nested block containing configuration options for Couchbase connections.  
     * *See Configuration Options for more info*
     */
    public readonly couchbases!: pulumi.Output<outputs.database.SecretsMountCouchbase[] | undefined>;
    /**
     * Default lease duration for tokens and secrets in seconds
     */
    public readonly defaultLeaseTtlSeconds!: pulumi.Output<number>;
    /**
     * Human-friendly description of the mount
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * A nested block containing configuration options for Elasticsearch connections.  
     * *See Configuration Options for more info*
     */
    public readonly elasticsearches!: pulumi.Output<outputs.database.SecretsMountElasticsearch[] | undefined>;
    /**
     * The total number of database secrets engines configured.
     */
    public /*out*/ readonly engineCount!: pulumi.Output<number>;
    /**
     * Boolean flag that can be explicitly set to true to enable the secrets engine to access Vault's external entropy source
     */
    public readonly externalEntropyAccess!: pulumi.Output<boolean | undefined>;
    /**
     * A nested block containing configuration options for SAP HanaDB connections.  
     * *See Configuration Options for more info*
     */
    public readonly hanas!: pulumi.Output<outputs.database.SecretsMountHana[] | undefined>;
    /**
     * A nested block containing configuration options for InfluxDB connections.  
     * *See Configuration Options for more info*
     */
    public readonly influxdbs!: pulumi.Output<outputs.database.SecretsMountInfluxdb[] | undefined>;
    /**
     * Boolean flag that can be explicitly set to true to enforce local mount in HA environment
     */
    public readonly local!: pulumi.Output<boolean | undefined>;
    /**
     * Maximum possible lease duration for tokens and secrets in seconds
     */
    public readonly maxLeaseTtlSeconds!: pulumi.Output<number>;
    /**
     * A nested block containing configuration options for MongoDB Atlas connections.  
     * *See Configuration Options for more info*
     */
    public readonly mongodbatlas!: pulumi.Output<outputs.database.SecretsMountMongodbatla[] | undefined>;
    /**
     * A nested block containing configuration options for MongoDB connections.  
     * *See Configuration Options for more info*
     */
    public readonly mongodbs!: pulumi.Output<outputs.database.SecretsMountMongodb[] | undefined>;
    /**
     * A nested block containing configuration options for MSSQL connections.  
     * *See Configuration Options for more info*
     */
    public readonly mssqls!: pulumi.Output<outputs.database.SecretsMountMssql[] | undefined>;
    /**
     * A nested block containing configuration options for Aurora MySQL connections.  
     * *See Configuration Options for more info*
     */
    public readonly mysqlAuroras!: pulumi.Output<outputs.database.SecretsMountMysqlAurora[] | undefined>;
    /**
     * A nested block containing configuration options for legacy MySQL connections.  
     * *See Configuration Options for more info*
     */
    public readonly mysqlLegacies!: pulumi.Output<outputs.database.SecretsMountMysqlLegacy[] | undefined>;
    /**
     * A nested block containing configuration options for RDS MySQL connections.  
     * *See Configuration Options for more info*
     */
    public readonly mysqlRds!: pulumi.Output<outputs.database.SecretsMountMysqlRd[] | undefined>;
    /**
     * A nested block containing configuration options for MySQL connections.  
     * *See Configuration Options for more info*
     */
    public readonly mysqls!: pulumi.Output<outputs.database.SecretsMountMysql[] | undefined>;
    /**
     * Target namespace. (requires Enterprise)
     */
    public readonly namespace!: pulumi.Output<string | undefined>;
    /**
     * Specifies mount type specific options that are passed to the backend
     */
    public readonly options!: pulumi.Output<{[key: string]: any} | undefined>;
    /**
     * A nested block containing configuration options for Oracle connections.  
     * *See Configuration Options for more info*
     */
    public readonly oracles!: pulumi.Output<outputs.database.SecretsMountOracle[] | undefined>;
    /**
     * Where the secret backend will be mounted
     */
    public readonly path!: pulumi.Output<string>;
    /**
     * A nested block containing configuration options for PostgreSQL connections.  
     * *See Configuration Options for more info*
     */
    public readonly postgresqls!: pulumi.Output<outputs.database.SecretsMountPostgresql[] | undefined>;
    /**
     * A nested block containing configuration options for Redis connections.  
     * *See Configuration Options for more info*
     */
    public readonly redis!: pulumi.Output<outputs.database.SecretsMountRedi[] | undefined>;
    /**
     * A nested block containing configuration options for Redis ElastiCache connections.  
     * *See Configuration Options for more info*
     */
    public readonly redisElasticaches!: pulumi.Output<outputs.database.SecretsMountRedisElasticach[] | undefined>;
    /**
     * A nested block containing configuration options for AWS Redshift connections.  
     * *See Configuration Options for more info*
     */
    public readonly redshifts!: pulumi.Output<outputs.database.SecretsMountRedshift[] | undefined>;
    /**
     * Boolean flag that can be explicitly set to true to enable seal wrapping for the mount, causing values stored by the mount to be wrapped by the seal's encryption capability
     */
    public readonly sealWrap!: pulumi.Output<boolean>;
    /**
     * A nested block containing configuration options for Snowflake connections.  
     * *See Configuration Options for more info*
     */
    public readonly snowflakes!: pulumi.Output<outputs.database.SecretsMountSnowflake[] | undefined>;

    /**
     * Create a SecretsMount resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SecretsMountArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SecretsMountArgs | SecretsMountState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SecretsMountState | undefined;
            resourceInputs["accessor"] = state ? state.accessor : undefined;
            resourceInputs["allowedManagedKeys"] = state ? state.allowedManagedKeys : undefined;
            resourceInputs["auditNonHmacRequestKeys"] = state ? state.auditNonHmacRequestKeys : undefined;
            resourceInputs["auditNonHmacResponseKeys"] = state ? state.auditNonHmacResponseKeys : undefined;
            resourceInputs["cassandras"] = state ? state.cassandras : undefined;
            resourceInputs["couchbases"] = state ? state.couchbases : undefined;
            resourceInputs["defaultLeaseTtlSeconds"] = state ? state.defaultLeaseTtlSeconds : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["elasticsearches"] = state ? state.elasticsearches : undefined;
            resourceInputs["engineCount"] = state ? state.engineCount : undefined;
            resourceInputs["externalEntropyAccess"] = state ? state.externalEntropyAccess : undefined;
            resourceInputs["hanas"] = state ? state.hanas : undefined;
            resourceInputs["influxdbs"] = state ? state.influxdbs : undefined;
            resourceInputs["local"] = state ? state.local : undefined;
            resourceInputs["maxLeaseTtlSeconds"] = state ? state.maxLeaseTtlSeconds : undefined;
            resourceInputs["mongodbatlas"] = state ? state.mongodbatlas : undefined;
            resourceInputs["mongodbs"] = state ? state.mongodbs : undefined;
            resourceInputs["mssqls"] = state ? state.mssqls : undefined;
            resourceInputs["mysqlAuroras"] = state ? state.mysqlAuroras : undefined;
            resourceInputs["mysqlLegacies"] = state ? state.mysqlLegacies : undefined;
            resourceInputs["mysqlRds"] = state ? state.mysqlRds : undefined;
            resourceInputs["mysqls"] = state ? state.mysqls : undefined;
            resourceInputs["namespace"] = state ? state.namespace : undefined;
            resourceInputs["options"] = state ? state.options : undefined;
            resourceInputs["oracles"] = state ? state.oracles : undefined;
            resourceInputs["path"] = state ? state.path : undefined;
            resourceInputs["postgresqls"] = state ? state.postgresqls : undefined;
            resourceInputs["redis"] = state ? state.redis : undefined;
            resourceInputs["redisElasticaches"] = state ? state.redisElasticaches : undefined;
            resourceInputs["redshifts"] = state ? state.redshifts : undefined;
            resourceInputs["sealWrap"] = state ? state.sealWrap : undefined;
            resourceInputs["snowflakes"] = state ? state.snowflakes : undefined;
        } else {
            const args = argsOrState as SecretsMountArgs | undefined;
            if ((!args || args.path === undefined) && !opts.urn) {
                throw new Error("Missing required property 'path'");
            }
            resourceInputs["allowedManagedKeys"] = args ? args.allowedManagedKeys : undefined;
            resourceInputs["auditNonHmacRequestKeys"] = args ? args.auditNonHmacRequestKeys : undefined;
            resourceInputs["auditNonHmacResponseKeys"] = args ? args.auditNonHmacResponseKeys : undefined;
            resourceInputs["cassandras"] = args ? args.cassandras : undefined;
            resourceInputs["couchbases"] = args ? args.couchbases : undefined;
            resourceInputs["defaultLeaseTtlSeconds"] = args ? args.defaultLeaseTtlSeconds : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["elasticsearches"] = args ? args.elasticsearches : undefined;
            resourceInputs["externalEntropyAccess"] = args ? args.externalEntropyAccess : undefined;
            resourceInputs["hanas"] = args ? args.hanas : undefined;
            resourceInputs["influxdbs"] = args ? args.influxdbs : undefined;
            resourceInputs["local"] = args ? args.local : undefined;
            resourceInputs["maxLeaseTtlSeconds"] = args ? args.maxLeaseTtlSeconds : undefined;
            resourceInputs["mongodbatlas"] = args ? args.mongodbatlas : undefined;
            resourceInputs["mongodbs"] = args ? args.mongodbs : undefined;
            resourceInputs["mssqls"] = args ? args.mssqls : undefined;
            resourceInputs["mysqlAuroras"] = args ? args.mysqlAuroras : undefined;
            resourceInputs["mysqlLegacies"] = args ? args.mysqlLegacies : undefined;
            resourceInputs["mysqlRds"] = args ? args.mysqlRds : undefined;
            resourceInputs["mysqls"] = args ? args.mysqls : undefined;
            resourceInputs["namespace"] = args ? args.namespace : undefined;
            resourceInputs["options"] = args ? args.options : undefined;
            resourceInputs["oracles"] = args ? args.oracles : undefined;
            resourceInputs["path"] = args ? args.path : undefined;
            resourceInputs["postgresqls"] = args ? args.postgresqls : undefined;
            resourceInputs["redis"] = args ? args.redis : undefined;
            resourceInputs["redisElasticaches"] = args ? args.redisElasticaches : undefined;
            resourceInputs["redshifts"] = args ? args.redshifts : undefined;
            resourceInputs["sealWrap"] = args ? args.sealWrap : undefined;
            resourceInputs["snowflakes"] = args ? args.snowflakes : undefined;
            resourceInputs["accessor"] = undefined /*out*/;
            resourceInputs["engineCount"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(SecretsMount.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SecretsMount resources.
 */
export interface SecretsMountState {
    /**
     * Accessor of the mount
     */
    accessor?: pulumi.Input<string>;
    /**
     * Set of managed key registry entry names that the mount in question is allowed to access
     *
     * The following arguments are common to all database engines:
     */
    allowedManagedKeys?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Specifies the list of keys that will not be HMAC'd by audit devices in the request data object.
     */
    auditNonHmacRequestKeys?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Specifies the list of keys that will not be HMAC'd by audit devices in the response data object.
     */
    auditNonHmacResponseKeys?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A nested block containing configuration options for Cassandra connections.  
     * *See Configuration Options for more info*
     */
    cassandras?: pulumi.Input<pulumi.Input<inputs.database.SecretsMountCassandra>[]>;
    /**
     * A nested block containing configuration options for Couchbase connections.  
     * *See Configuration Options for more info*
     */
    couchbases?: pulumi.Input<pulumi.Input<inputs.database.SecretsMountCouchbase>[]>;
    /**
     * Default lease duration for tokens and secrets in seconds
     */
    defaultLeaseTtlSeconds?: pulumi.Input<number>;
    /**
     * Human-friendly description of the mount
     */
    description?: pulumi.Input<string>;
    /**
     * A nested block containing configuration options for Elasticsearch connections.  
     * *See Configuration Options for more info*
     */
    elasticsearches?: pulumi.Input<pulumi.Input<inputs.database.SecretsMountElasticsearch>[]>;
    /**
     * The total number of database secrets engines configured.
     */
    engineCount?: pulumi.Input<number>;
    /**
     * Boolean flag that can be explicitly set to true to enable the secrets engine to access Vault's external entropy source
     */
    externalEntropyAccess?: pulumi.Input<boolean>;
    /**
     * A nested block containing configuration options for SAP HanaDB connections.  
     * *See Configuration Options for more info*
     */
    hanas?: pulumi.Input<pulumi.Input<inputs.database.SecretsMountHana>[]>;
    /**
     * A nested block containing configuration options for InfluxDB connections.  
     * *See Configuration Options for more info*
     */
    influxdbs?: pulumi.Input<pulumi.Input<inputs.database.SecretsMountInfluxdb>[]>;
    /**
     * Boolean flag that can be explicitly set to true to enforce local mount in HA environment
     */
    local?: pulumi.Input<boolean>;
    /**
     * Maximum possible lease duration for tokens and secrets in seconds
     */
    maxLeaseTtlSeconds?: pulumi.Input<number>;
    /**
     * A nested block containing configuration options for MongoDB Atlas connections.  
     * *See Configuration Options for more info*
     */
    mongodbatlas?: pulumi.Input<pulumi.Input<inputs.database.SecretsMountMongodbatla>[]>;
    /**
     * A nested block containing configuration options for MongoDB connections.  
     * *See Configuration Options for more info*
     */
    mongodbs?: pulumi.Input<pulumi.Input<inputs.database.SecretsMountMongodb>[]>;
    /**
     * A nested block containing configuration options for MSSQL connections.  
     * *See Configuration Options for more info*
     */
    mssqls?: pulumi.Input<pulumi.Input<inputs.database.SecretsMountMssql>[]>;
    /**
     * A nested block containing configuration options for Aurora MySQL connections.  
     * *See Configuration Options for more info*
     */
    mysqlAuroras?: pulumi.Input<pulumi.Input<inputs.database.SecretsMountMysqlAurora>[]>;
    /**
     * A nested block containing configuration options for legacy MySQL connections.  
     * *See Configuration Options for more info*
     */
    mysqlLegacies?: pulumi.Input<pulumi.Input<inputs.database.SecretsMountMysqlLegacy>[]>;
    /**
     * A nested block containing configuration options for RDS MySQL connections.  
     * *See Configuration Options for more info*
     */
    mysqlRds?: pulumi.Input<pulumi.Input<inputs.database.SecretsMountMysqlRd>[]>;
    /**
     * A nested block containing configuration options for MySQL connections.  
     * *See Configuration Options for more info*
     */
    mysqls?: pulumi.Input<pulumi.Input<inputs.database.SecretsMountMysql>[]>;
    /**
     * Target namespace. (requires Enterprise)
     */
    namespace?: pulumi.Input<string>;
    /**
     * Specifies mount type specific options that are passed to the backend
     */
    options?: pulumi.Input<{[key: string]: any}>;
    /**
     * A nested block containing configuration options for Oracle connections.  
     * *See Configuration Options for more info*
     */
    oracles?: pulumi.Input<pulumi.Input<inputs.database.SecretsMountOracle>[]>;
    /**
     * Where the secret backend will be mounted
     */
    path?: pulumi.Input<string>;
    /**
     * A nested block containing configuration options for PostgreSQL connections.  
     * *See Configuration Options for more info*
     */
    postgresqls?: pulumi.Input<pulumi.Input<inputs.database.SecretsMountPostgresql>[]>;
    /**
     * A nested block containing configuration options for Redis connections.  
     * *See Configuration Options for more info*
     */
    redis?: pulumi.Input<pulumi.Input<inputs.database.SecretsMountRedi>[]>;
    /**
     * A nested block containing configuration options for Redis ElastiCache connections.  
     * *See Configuration Options for more info*
     */
    redisElasticaches?: pulumi.Input<pulumi.Input<inputs.database.SecretsMountRedisElasticach>[]>;
    /**
     * A nested block containing configuration options for AWS Redshift connections.  
     * *See Configuration Options for more info*
     */
    redshifts?: pulumi.Input<pulumi.Input<inputs.database.SecretsMountRedshift>[]>;
    /**
     * Boolean flag that can be explicitly set to true to enable seal wrapping for the mount, causing values stored by the mount to be wrapped by the seal's encryption capability
     */
    sealWrap?: pulumi.Input<boolean>;
    /**
     * A nested block containing configuration options for Snowflake connections.  
     * *See Configuration Options for more info*
     */
    snowflakes?: pulumi.Input<pulumi.Input<inputs.database.SecretsMountSnowflake>[]>;
}

/**
 * The set of arguments for constructing a SecretsMount resource.
 */
export interface SecretsMountArgs {
    /**
     * Set of managed key registry entry names that the mount in question is allowed to access
     *
     * The following arguments are common to all database engines:
     */
    allowedManagedKeys?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Specifies the list of keys that will not be HMAC'd by audit devices in the request data object.
     */
    auditNonHmacRequestKeys?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Specifies the list of keys that will not be HMAC'd by audit devices in the response data object.
     */
    auditNonHmacResponseKeys?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A nested block containing configuration options for Cassandra connections.  
     * *See Configuration Options for more info*
     */
    cassandras?: pulumi.Input<pulumi.Input<inputs.database.SecretsMountCassandra>[]>;
    /**
     * A nested block containing configuration options for Couchbase connections.  
     * *See Configuration Options for more info*
     */
    couchbases?: pulumi.Input<pulumi.Input<inputs.database.SecretsMountCouchbase>[]>;
    /**
     * Default lease duration for tokens and secrets in seconds
     */
    defaultLeaseTtlSeconds?: pulumi.Input<number>;
    /**
     * Human-friendly description of the mount
     */
    description?: pulumi.Input<string>;
    /**
     * A nested block containing configuration options for Elasticsearch connections.  
     * *See Configuration Options for more info*
     */
    elasticsearches?: pulumi.Input<pulumi.Input<inputs.database.SecretsMountElasticsearch>[]>;
    /**
     * Boolean flag that can be explicitly set to true to enable the secrets engine to access Vault's external entropy source
     */
    externalEntropyAccess?: pulumi.Input<boolean>;
    /**
     * A nested block containing configuration options for SAP HanaDB connections.  
     * *See Configuration Options for more info*
     */
    hanas?: pulumi.Input<pulumi.Input<inputs.database.SecretsMountHana>[]>;
    /**
     * A nested block containing configuration options for InfluxDB connections.  
     * *See Configuration Options for more info*
     */
    influxdbs?: pulumi.Input<pulumi.Input<inputs.database.SecretsMountInfluxdb>[]>;
    /**
     * Boolean flag that can be explicitly set to true to enforce local mount in HA environment
     */
    local?: pulumi.Input<boolean>;
    /**
     * Maximum possible lease duration for tokens and secrets in seconds
     */
    maxLeaseTtlSeconds?: pulumi.Input<number>;
    /**
     * A nested block containing configuration options for MongoDB Atlas connections.  
     * *See Configuration Options for more info*
     */
    mongodbatlas?: pulumi.Input<pulumi.Input<inputs.database.SecretsMountMongodbatla>[]>;
    /**
     * A nested block containing configuration options for MongoDB connections.  
     * *See Configuration Options for more info*
     */
    mongodbs?: pulumi.Input<pulumi.Input<inputs.database.SecretsMountMongodb>[]>;
    /**
     * A nested block containing configuration options for MSSQL connections.  
     * *See Configuration Options for more info*
     */
    mssqls?: pulumi.Input<pulumi.Input<inputs.database.SecretsMountMssql>[]>;
    /**
     * A nested block containing configuration options for Aurora MySQL connections.  
     * *See Configuration Options for more info*
     */
    mysqlAuroras?: pulumi.Input<pulumi.Input<inputs.database.SecretsMountMysqlAurora>[]>;
    /**
     * A nested block containing configuration options for legacy MySQL connections.  
     * *See Configuration Options for more info*
     */
    mysqlLegacies?: pulumi.Input<pulumi.Input<inputs.database.SecretsMountMysqlLegacy>[]>;
    /**
     * A nested block containing configuration options for RDS MySQL connections.  
     * *See Configuration Options for more info*
     */
    mysqlRds?: pulumi.Input<pulumi.Input<inputs.database.SecretsMountMysqlRd>[]>;
    /**
     * A nested block containing configuration options for MySQL connections.  
     * *See Configuration Options for more info*
     */
    mysqls?: pulumi.Input<pulumi.Input<inputs.database.SecretsMountMysql>[]>;
    /**
     * Target namespace. (requires Enterprise)
     */
    namespace?: pulumi.Input<string>;
    /**
     * Specifies mount type specific options that are passed to the backend
     */
    options?: pulumi.Input<{[key: string]: any}>;
    /**
     * A nested block containing configuration options for Oracle connections.  
     * *See Configuration Options for more info*
     */
    oracles?: pulumi.Input<pulumi.Input<inputs.database.SecretsMountOracle>[]>;
    /**
     * Where the secret backend will be mounted
     */
    path: pulumi.Input<string>;
    /**
     * A nested block containing configuration options for PostgreSQL connections.  
     * *See Configuration Options for more info*
     */
    postgresqls?: pulumi.Input<pulumi.Input<inputs.database.SecretsMountPostgresql>[]>;
    /**
     * A nested block containing configuration options for Redis connections.  
     * *See Configuration Options for more info*
     */
    redis?: pulumi.Input<pulumi.Input<inputs.database.SecretsMountRedi>[]>;
    /**
     * A nested block containing configuration options for Redis ElastiCache connections.  
     * *See Configuration Options for more info*
     */
    redisElasticaches?: pulumi.Input<pulumi.Input<inputs.database.SecretsMountRedisElasticach>[]>;
    /**
     * A nested block containing configuration options for AWS Redshift connections.  
     * *See Configuration Options for more info*
     */
    redshifts?: pulumi.Input<pulumi.Input<inputs.database.SecretsMountRedshift>[]>;
    /**
     * Boolean flag that can be explicitly set to true to enable seal wrapping for the mount, causing values stored by the mount to be wrapped by the seal's encryption capability
     */
    sealWrap?: pulumi.Input<boolean>;
    /**
     * A nested block containing configuration options for Snowflake connections.  
     * *See Configuration Options for more info*
     */
    snowflakes?: pulumi.Input<pulumi.Input<inputs.database.SecretsMountSnowflake>[]>;
}
