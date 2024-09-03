// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vault.database;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.vault.Utilities;
import com.pulumi.vault.database.SecretsMountArgs;
import com.pulumi.vault.database.inputs.SecretsMountState;
import com.pulumi.vault.database.outputs.SecretsMountCassandra;
import com.pulumi.vault.database.outputs.SecretsMountCouchbase;
import com.pulumi.vault.database.outputs.SecretsMountElasticsearch;
import com.pulumi.vault.database.outputs.SecretsMountHana;
import com.pulumi.vault.database.outputs.SecretsMountInfluxdb;
import com.pulumi.vault.database.outputs.SecretsMountMongodb;
import com.pulumi.vault.database.outputs.SecretsMountMongodbatla;
import com.pulumi.vault.database.outputs.SecretsMountMssql;
import com.pulumi.vault.database.outputs.SecretsMountMysql;
import com.pulumi.vault.database.outputs.SecretsMountMysqlAurora;
import com.pulumi.vault.database.outputs.SecretsMountMysqlLegacy;
import com.pulumi.vault.database.outputs.SecretsMountMysqlRd;
import com.pulumi.vault.database.outputs.SecretsMountOracle;
import com.pulumi.vault.database.outputs.SecretsMountPostgresql;
import com.pulumi.vault.database.outputs.SecretsMountRedi;
import com.pulumi.vault.database.outputs.SecretsMountRedisElasticach;
import com.pulumi.vault.database.outputs.SecretsMountRedshift;
import com.pulumi.vault.database.outputs.SecretsMountSnowflake;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * ## Example Usage
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.vault.database.SecretsMount;
 * import com.pulumi.vault.database.SecretsMountArgs;
 * import com.pulumi.vault.database.inputs.SecretsMountMssqlArgs;
 * import com.pulumi.vault.database.inputs.SecretsMountPostgresqlArgs;
 * import com.pulumi.vault.database.SecretBackendRole;
 * import com.pulumi.vault.database.SecretBackendRoleArgs;
 * import java.util.List;
 * import java.util.ArrayList;
 * import java.util.Map;
 * import java.io.File;
 * import java.nio.file.Files;
 * import java.nio.file.Paths;
 * 
 * public class App }{{@code
 *     public static void main(String[] args) }{{@code
 *         Pulumi.run(App::stack);
 *     }}{@code
 * 
 *     public static void stack(Context ctx) }{{@code
 *         var db = new SecretsMount("db", SecretsMountArgs.builder()
 *             .path("db")
 *             .mssqls(SecretsMountMssqlArgs.builder()
 *                 .name("db1")
 *                 .username("sa")
 *                 .password("super_secret_1")
 *                 .connectionUrl("sqlserver://}{{{@code username}}}{@code :}{{{@code password}}}{@literal @}{@code 127.0.0.1:1433")
 *                 .allowedRoles("dev1")
 *                 .build())
 *             .postgresqls(SecretsMountPostgresqlArgs.builder()
 *                 .name("db2")
 *                 .username("postgres")
 *                 .password("super_secret_2")
 *                 .connectionUrl("postgresql://}{{{@code username}}}{@code :}{{{@code password}}}{@literal @}{@code 127.0.0.1:5432/postgres")
 *                 .verifyConnection(true)
 *                 .allowedRoles("dev2")
 *                 .build())
 *             .build());
 * 
 *         var dev1 = new SecretBackendRole("dev1", SecretBackendRoleArgs.builder()
 *             .name("dev1")
 *             .backend(db.path())
 *             .dbName(db.mssqls().applyValue(mssqls -> mssqls[0].name()))
 *             .creationStatements(            
 *                 "CREATE LOGIN [}{{{@code name}}}{@code ] WITH PASSWORD = '}{{{@code password}}}{@code ';",
 *                 "CREATE USER [}{{{@code name}}}{@code ] FOR LOGIN [}{{{@code name}}}{@code ];",
 *                 "GRANT SELECT ON SCHEMA::dbo TO [}{{{@code name}}}{@code ];")
 *             .build());
 * 
 *         var dev2 = new SecretBackendRole("dev2", SecretBackendRoleArgs.builder()
 *             .name("dev2")
 *             .backend(db.path())
 *             .dbName(db.postgresqls().applyValue(postgresqls -> postgresqls[0].name()))
 *             .creationStatements(            
 *                 "CREATE ROLE \"}{{{@code name}}}{@code \" WITH LOGIN PASSWORD '}{{{@code password}}}{@code ' VALID UNTIL '}{{{@code expiration}}}{@code ';",
 *                 "GRANT SELECT ON ALL TABLES IN SCHEMA public TO \"}{{{@code name}}}{@code \";")
 *             .build());
 * 
 *     }}{@code
 * }}{@code
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Import
 * 
 * Database secret backend connections can be imported using the `path` e.g.
 * 
 * ```sh
 * $ pulumi import vault:database/secretsMount:SecretsMount db db
 * ```
 * 
 */
@ResourceType(type="vault:database/secretsMount:SecretsMount")
public class SecretsMount extends com.pulumi.resources.CustomResource {
    /**
     * Accessor of the mount
     * 
     */
    @Export(name="accessor", refs={String.class}, tree="[0]")
    private Output<String> accessor;

    /**
     * @return Accessor of the mount
     * 
     */
    public Output<String> accessor() {
        return this.accessor;
    }
    /**
     * Set of managed key registry entry names that the mount in question is allowed to access
     * 
     * The following arguments are common to all database engines:
     * 
     */
    @Export(name="allowedManagedKeys", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> allowedManagedKeys;

    /**
     * @return Set of managed key registry entry names that the mount in question is allowed to access
     * 
     * The following arguments are common to all database engines:
     * 
     */
    public Output<Optional<List<String>>> allowedManagedKeys() {
        return Codegen.optional(this.allowedManagedKeys);
    }
    /**
     * List of headers to allow and pass from the request to the plugin
     * 
     */
    @Export(name="allowedResponseHeaders", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> allowedResponseHeaders;

    /**
     * @return List of headers to allow and pass from the request to the plugin
     * 
     */
    public Output<Optional<List<String>>> allowedResponseHeaders() {
        return Codegen.optional(this.allowedResponseHeaders);
    }
    /**
     * Specifies the list of keys that will not be HMAC&#39;d by audit devices in the request data object.
     * 
     */
    @Export(name="auditNonHmacRequestKeys", refs={List.class,String.class}, tree="[0,1]")
    private Output<List<String>> auditNonHmacRequestKeys;

    /**
     * @return Specifies the list of keys that will not be HMAC&#39;d by audit devices in the request data object.
     * 
     */
    public Output<List<String>> auditNonHmacRequestKeys() {
        return this.auditNonHmacRequestKeys;
    }
    /**
     * Specifies the list of keys that will not be HMAC&#39;d by audit devices in the response data object.
     * 
     */
    @Export(name="auditNonHmacResponseKeys", refs={List.class,String.class}, tree="[0,1]")
    private Output<List<String>> auditNonHmacResponseKeys;

    /**
     * @return Specifies the list of keys that will not be HMAC&#39;d by audit devices in the response data object.
     * 
     */
    public Output<List<String>> auditNonHmacResponseKeys() {
        return this.auditNonHmacResponseKeys;
    }
    /**
     * A nested block containing configuration options for Cassandra connections.*See Configuration Options for more info*
     * 
     */
    @Export(name="cassandras", refs={List.class,SecretsMountCassandra.class}, tree="[0,1]")
    private Output</* @Nullable */ List<SecretsMountCassandra>> cassandras;

    /**
     * @return A nested block containing configuration options for Cassandra connections.*See Configuration Options for more info*
     * 
     */
    public Output<Optional<List<SecretsMountCassandra>>> cassandras() {
        return Codegen.optional(this.cassandras);
    }
    /**
     * A nested block containing configuration options for Couchbase connections.*See Configuration Options for more info*
     * 
     */
    @Export(name="couchbases", refs={List.class,SecretsMountCouchbase.class}, tree="[0,1]")
    private Output</* @Nullable */ List<SecretsMountCouchbase>> couchbases;

    /**
     * @return A nested block containing configuration options for Couchbase connections.*See Configuration Options for more info*
     * 
     */
    public Output<Optional<List<SecretsMountCouchbase>>> couchbases() {
        return Codegen.optional(this.couchbases);
    }
    /**
     * Default lease duration for tokens and secrets in seconds
     * 
     */
    @Export(name="defaultLeaseTtlSeconds", refs={Integer.class}, tree="[0]")
    private Output<Integer> defaultLeaseTtlSeconds;

    /**
     * @return Default lease duration for tokens and secrets in seconds
     * 
     */
    public Output<Integer> defaultLeaseTtlSeconds() {
        return this.defaultLeaseTtlSeconds;
    }
    /**
     * List of headers to allow and pass from the request to the plugin
     * 
     */
    @Export(name="delegatedAuthAccessors", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> delegatedAuthAccessors;

    /**
     * @return List of headers to allow and pass from the request to the plugin
     * 
     */
    public Output<Optional<List<String>>> delegatedAuthAccessors() {
        return Codegen.optional(this.delegatedAuthAccessors);
    }
    /**
     * Human-friendly description of the mount
     * 
     */
    @Export(name="description", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> description;

    /**
     * @return Human-friendly description of the mount
     * 
     */
    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    /**
     * A nested block containing configuration options for Elasticsearch connections.*See Configuration Options for more info*
     * 
     */
    @Export(name="elasticsearches", refs={List.class,SecretsMountElasticsearch.class}, tree="[0,1]")
    private Output</* @Nullable */ List<SecretsMountElasticsearch>> elasticsearches;

    /**
     * @return A nested block containing configuration options for Elasticsearch connections.*See Configuration Options for more info*
     * 
     */
    public Output<Optional<List<SecretsMountElasticsearch>>> elasticsearches() {
        return Codegen.optional(this.elasticsearches);
    }
    /**
     * The total number of database secrets engines configured.
     * 
     */
    @Export(name="engineCount", refs={Integer.class}, tree="[0]")
    private Output<Integer> engineCount;

    /**
     * @return The total number of database secrets engines configured.
     * 
     */
    public Output<Integer> engineCount() {
        return this.engineCount;
    }
    /**
     * Boolean flag that can be explicitly set to true to enable the secrets engine to access Vault&#39;s external entropy source
     * 
     */
    @Export(name="externalEntropyAccess", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> externalEntropyAccess;

    /**
     * @return Boolean flag that can be explicitly set to true to enable the secrets engine to access Vault&#39;s external entropy source
     * 
     */
    public Output<Optional<Boolean>> externalEntropyAccess() {
        return Codegen.optional(this.externalEntropyAccess);
    }
    /**
     * A nested block containing configuration options for SAP HanaDB connections.*See Configuration Options for more info*
     * 
     */
    @Export(name="hanas", refs={List.class,SecretsMountHana.class}, tree="[0,1]")
    private Output</* @Nullable */ List<SecretsMountHana>> hanas;

    /**
     * @return A nested block containing configuration options for SAP HanaDB connections.*See Configuration Options for more info*
     * 
     */
    public Output<Optional<List<SecretsMountHana>>> hanas() {
        return Codegen.optional(this.hanas);
    }
    /**
     * The key to use for signing plugin workload identity tokens
     * 
     */
    @Export(name="identityTokenKey", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> identityTokenKey;

    /**
     * @return The key to use for signing plugin workload identity tokens
     * 
     */
    public Output<Optional<String>> identityTokenKey() {
        return Codegen.optional(this.identityTokenKey);
    }
    /**
     * A nested block containing configuration options for InfluxDB connections.*See Configuration Options for more info*
     * 
     */
    @Export(name="influxdbs", refs={List.class,SecretsMountInfluxdb.class}, tree="[0,1]")
    private Output</* @Nullable */ List<SecretsMountInfluxdb>> influxdbs;

    /**
     * @return A nested block containing configuration options for InfluxDB connections.*See Configuration Options for more info*
     * 
     */
    public Output<Optional<List<SecretsMountInfluxdb>>> influxdbs() {
        return Codegen.optional(this.influxdbs);
    }
    /**
     * Specifies whether to show this mount in the UI-specific listing endpoint
     * 
     */
    @Export(name="listingVisibility", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> listingVisibility;

    /**
     * @return Specifies whether to show this mount in the UI-specific listing endpoint
     * 
     */
    public Output<Optional<String>> listingVisibility() {
        return Codegen.optional(this.listingVisibility);
    }
    /**
     * Boolean flag that can be explicitly set to true to enforce local mount in HA environment
     * 
     */
    @Export(name="local", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> local;

    /**
     * @return Boolean flag that can be explicitly set to true to enforce local mount in HA environment
     * 
     */
    public Output<Optional<Boolean>> local() {
        return Codegen.optional(this.local);
    }
    /**
     * Maximum possible lease duration for tokens and secrets in seconds
     * 
     */
    @Export(name="maxLeaseTtlSeconds", refs={Integer.class}, tree="[0]")
    private Output<Integer> maxLeaseTtlSeconds;

    /**
     * @return Maximum possible lease duration for tokens and secrets in seconds
     * 
     */
    public Output<Integer> maxLeaseTtlSeconds() {
        return this.maxLeaseTtlSeconds;
    }
    /**
     * A nested block containing configuration options for MongoDB Atlas connections.*See Configuration Options for more info*
     * 
     */
    @Export(name="mongodbatlas", refs={List.class,SecretsMountMongodbatla.class}, tree="[0,1]")
    private Output</* @Nullable */ List<SecretsMountMongodbatla>> mongodbatlas;

    /**
     * @return A nested block containing configuration options for MongoDB Atlas connections.*See Configuration Options for more info*
     * 
     */
    public Output<Optional<List<SecretsMountMongodbatla>>> mongodbatlas() {
        return Codegen.optional(this.mongodbatlas);
    }
    /**
     * A nested block containing configuration options for MongoDB connections.*See Configuration Options for more info*
     * 
     */
    @Export(name="mongodbs", refs={List.class,SecretsMountMongodb.class}, tree="[0,1]")
    private Output</* @Nullable */ List<SecretsMountMongodb>> mongodbs;

    /**
     * @return A nested block containing configuration options for MongoDB connections.*See Configuration Options for more info*
     * 
     */
    public Output<Optional<List<SecretsMountMongodb>>> mongodbs() {
        return Codegen.optional(this.mongodbs);
    }
    /**
     * A nested block containing configuration options for MSSQL connections.*See Configuration Options for more info*
     * 
     */
    @Export(name="mssqls", refs={List.class,SecretsMountMssql.class}, tree="[0,1]")
    private Output</* @Nullable */ List<SecretsMountMssql>> mssqls;

    /**
     * @return A nested block containing configuration options for MSSQL connections.*See Configuration Options for more info*
     * 
     */
    public Output<Optional<List<SecretsMountMssql>>> mssqls() {
        return Codegen.optional(this.mssqls);
    }
    /**
     * A nested block containing configuration options for Aurora MySQL connections.*See Configuration Options for more info*
     * 
     */
    @Export(name="mysqlAuroras", refs={List.class,SecretsMountMysqlAurora.class}, tree="[0,1]")
    private Output</* @Nullable */ List<SecretsMountMysqlAurora>> mysqlAuroras;

    /**
     * @return A nested block containing configuration options for Aurora MySQL connections.*See Configuration Options for more info*
     * 
     */
    public Output<Optional<List<SecretsMountMysqlAurora>>> mysqlAuroras() {
        return Codegen.optional(this.mysqlAuroras);
    }
    /**
     * A nested block containing configuration options for legacy MySQL connections.*See Configuration Options for more info*
     * 
     */
    @Export(name="mysqlLegacies", refs={List.class,SecretsMountMysqlLegacy.class}, tree="[0,1]")
    private Output</* @Nullable */ List<SecretsMountMysqlLegacy>> mysqlLegacies;

    /**
     * @return A nested block containing configuration options for legacy MySQL connections.*See Configuration Options for more info*
     * 
     */
    public Output<Optional<List<SecretsMountMysqlLegacy>>> mysqlLegacies() {
        return Codegen.optional(this.mysqlLegacies);
    }
    /**
     * A nested block containing configuration options for RDS MySQL connections.*See Configuration Options for more info*
     * 
     */
    @Export(name="mysqlRds", refs={List.class,SecretsMountMysqlRd.class}, tree="[0,1]")
    private Output</* @Nullable */ List<SecretsMountMysqlRd>> mysqlRds;

    /**
     * @return A nested block containing configuration options for RDS MySQL connections.*See Configuration Options for more info*
     * 
     */
    public Output<Optional<List<SecretsMountMysqlRd>>> mysqlRds() {
        return Codegen.optional(this.mysqlRds);
    }
    /**
     * A nested block containing configuration options for MySQL connections.*See Configuration Options for more info*
     * 
     */
    @Export(name="mysqls", refs={List.class,SecretsMountMysql.class}, tree="[0,1]")
    private Output</* @Nullable */ List<SecretsMountMysql>> mysqls;

    /**
     * @return A nested block containing configuration options for MySQL connections.*See Configuration Options for more info*
     * 
     */
    public Output<Optional<List<SecretsMountMysql>>> mysqls() {
        return Codegen.optional(this.mysqls);
    }
    /**
     * Target namespace. (requires Enterprise)
     * 
     */
    @Export(name="namespace", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> namespace;

    /**
     * @return Target namespace. (requires Enterprise)
     * 
     */
    public Output<Optional<String>> namespace() {
        return Codegen.optional(this.namespace);
    }
    /**
     * Specifies mount type specific options that are passed to the backend
     * 
     */
    @Export(name="options", refs={Map.class,String.class}, tree="[0,1,1]")
    private Output</* @Nullable */ Map<String,String>> options;

    /**
     * @return Specifies mount type specific options that are passed to the backend
     * 
     */
    public Output<Optional<Map<String,String>>> options() {
        return Codegen.optional(this.options);
    }
    /**
     * A nested block containing configuration options for Oracle connections.*See Configuration Options for more info*
     * 
     */
    @Export(name="oracles", refs={List.class,SecretsMountOracle.class}, tree="[0,1]")
    private Output</* @Nullable */ List<SecretsMountOracle>> oracles;

    /**
     * @return A nested block containing configuration options for Oracle connections.*See Configuration Options for more info*
     * 
     */
    public Output<Optional<List<SecretsMountOracle>>> oracles() {
        return Codegen.optional(this.oracles);
    }
    /**
     * List of headers to allow and pass from the request to the plugin
     * 
     */
    @Export(name="passthroughRequestHeaders", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> passthroughRequestHeaders;

    /**
     * @return List of headers to allow and pass from the request to the plugin
     * 
     */
    public Output<Optional<List<String>>> passthroughRequestHeaders() {
        return Codegen.optional(this.passthroughRequestHeaders);
    }
    /**
     * Where the secret backend will be mounted
     * 
     */
    @Export(name="path", refs={String.class}, tree="[0]")
    private Output<String> path;

    /**
     * @return Where the secret backend will be mounted
     * 
     */
    public Output<String> path() {
        return this.path;
    }
    /**
     * Specifies the semantic version of the plugin to use, e.g. &#39;v1.0.0&#39;
     * 
     */
    @Export(name="pluginVersion", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> pluginVersion;

    /**
     * @return Specifies the semantic version of the plugin to use, e.g. &#39;v1.0.0&#39;
     * 
     */
    public Output<Optional<String>> pluginVersion() {
        return Codegen.optional(this.pluginVersion);
    }
    /**
     * A nested block containing configuration options for PostgreSQL connections.*See Configuration Options for more info*
     * 
     */
    @Export(name="postgresqls", refs={List.class,SecretsMountPostgresql.class}, tree="[0,1]")
    private Output</* @Nullable */ List<SecretsMountPostgresql>> postgresqls;

    /**
     * @return A nested block containing configuration options for PostgreSQL connections.*See Configuration Options for more info*
     * 
     */
    public Output<Optional<List<SecretsMountPostgresql>>> postgresqls() {
        return Codegen.optional(this.postgresqls);
    }
    /**
     * A nested block containing configuration options for Redis connections.*See Configuration Options for more info*
     * 
     */
    @Export(name="redis", refs={List.class,SecretsMountRedi.class}, tree="[0,1]")
    private Output</* @Nullable */ List<SecretsMountRedi>> redis;

    /**
     * @return A nested block containing configuration options for Redis connections.*See Configuration Options for more info*
     * 
     */
    public Output<Optional<List<SecretsMountRedi>>> redis() {
        return Codegen.optional(this.redis);
    }
    /**
     * A nested block containing configuration options for Redis ElastiCache connections.*See Configuration Options for more info*
     * 
     */
    @Export(name="redisElasticaches", refs={List.class,SecretsMountRedisElasticach.class}, tree="[0,1]")
    private Output</* @Nullable */ List<SecretsMountRedisElasticach>> redisElasticaches;

    /**
     * @return A nested block containing configuration options for Redis ElastiCache connections.*See Configuration Options for more info*
     * 
     */
    public Output<Optional<List<SecretsMountRedisElasticach>>> redisElasticaches() {
        return Codegen.optional(this.redisElasticaches);
    }
    /**
     * A nested block containing configuration options for AWS Redshift connections.*See Configuration Options for more info*
     * 
     */
    @Export(name="redshifts", refs={List.class,SecretsMountRedshift.class}, tree="[0,1]")
    private Output</* @Nullable */ List<SecretsMountRedshift>> redshifts;

    /**
     * @return A nested block containing configuration options for AWS Redshift connections.*See Configuration Options for more info*
     * 
     */
    public Output<Optional<List<SecretsMountRedshift>>> redshifts() {
        return Codegen.optional(this.redshifts);
    }
    /**
     * Boolean flag that can be explicitly set to true to enable seal wrapping for the mount, causing values stored by the mount to be wrapped by the seal&#39;s encryption capability
     * 
     */
    @Export(name="sealWrap", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> sealWrap;

    /**
     * @return Boolean flag that can be explicitly set to true to enable seal wrapping for the mount, causing values stored by the mount to be wrapped by the seal&#39;s encryption capability
     * 
     */
    public Output<Boolean> sealWrap() {
        return this.sealWrap;
    }
    /**
     * A nested block containing configuration options for Snowflake connections.*See Configuration Options for more info*
     * 
     */
    @Export(name="snowflakes", refs={List.class,SecretsMountSnowflake.class}, tree="[0,1]")
    private Output</* @Nullable */ List<SecretsMountSnowflake>> snowflakes;

    /**
     * @return A nested block containing configuration options for Snowflake connections.*See Configuration Options for more info*
     * 
     */
    public Output<Optional<List<SecretsMountSnowflake>>> snowflakes() {
        return Codegen.optional(this.snowflakes);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public SecretsMount(java.lang.String name) {
        this(name, SecretsMountArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public SecretsMount(java.lang.String name, SecretsMountArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public SecretsMount(java.lang.String name, SecretsMountArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("vault:database/secretsMount:SecretsMount", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private SecretsMount(java.lang.String name, Output<java.lang.String> id, @Nullable SecretsMountState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("vault:database/secretsMount:SecretsMount", name, state, makeResourceOptions(options, id), false);
    }

    private static SecretsMountArgs makeArgs(SecretsMountArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? SecretsMountArgs.Empty : args;
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<java.lang.String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .build();
        return com.pulumi.resources.CustomResourceOptions.merge(defaultOptions, options, id);
    }

    /**
     * Get an existing Host resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state
     * @param options Optional settings to control the behavior of the CustomResource.
     */
    public static SecretsMount get(java.lang.String name, Output<java.lang.String> id, @Nullable SecretsMountState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new SecretsMount(name, id, state, options);
    }
}
