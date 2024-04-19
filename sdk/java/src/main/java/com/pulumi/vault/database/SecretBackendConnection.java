// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vault.database;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.vault.Utilities;
import com.pulumi.vault.database.SecretBackendConnectionArgs;
import com.pulumi.vault.database.inputs.SecretBackendConnectionState;
import com.pulumi.vault.database.outputs.SecretBackendConnectionCassandra;
import com.pulumi.vault.database.outputs.SecretBackendConnectionCouchbase;
import com.pulumi.vault.database.outputs.SecretBackendConnectionElasticsearch;
import com.pulumi.vault.database.outputs.SecretBackendConnectionHana;
import com.pulumi.vault.database.outputs.SecretBackendConnectionInfluxdb;
import com.pulumi.vault.database.outputs.SecretBackendConnectionMongodb;
import com.pulumi.vault.database.outputs.SecretBackendConnectionMongodbatlas;
import com.pulumi.vault.database.outputs.SecretBackendConnectionMssql;
import com.pulumi.vault.database.outputs.SecretBackendConnectionMysql;
import com.pulumi.vault.database.outputs.SecretBackendConnectionMysqlAurora;
import com.pulumi.vault.database.outputs.SecretBackendConnectionMysqlLegacy;
import com.pulumi.vault.database.outputs.SecretBackendConnectionMysqlRds;
import com.pulumi.vault.database.outputs.SecretBackendConnectionOracle;
import com.pulumi.vault.database.outputs.SecretBackendConnectionPostgresql;
import com.pulumi.vault.database.outputs.SecretBackendConnectionRedis;
import com.pulumi.vault.database.outputs.SecretBackendConnectionRedisElasticache;
import com.pulumi.vault.database.outputs.SecretBackendConnectionRedshift;
import com.pulumi.vault.database.outputs.SecretBackendConnectionSnowflake;
import java.lang.Boolean;
import java.lang.Object;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * ## Example Usage
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.vault.Mount;
 * import com.pulumi.vault.MountArgs;
 * import com.pulumi.vault.database.SecretBackendConnection;
 * import com.pulumi.vault.database.SecretBackendConnectionArgs;
 * import com.pulumi.vault.database.inputs.SecretBackendConnectionPostgresqlArgs;
 * import java.util.List;
 * import java.util.ArrayList;
 * import java.util.Map;
 * import java.io.File;
 * import java.nio.file.Files;
 * import java.nio.file.Paths;
 * 
 * public class App {
 *     public static void main(String[] args) {
 *         Pulumi.run(App::stack);
 *     }
 * 
 *     public static void stack(Context ctx) {
 *         var db = new Mount(&#34;db&#34;, MountArgs.builder()        
 *             .path(&#34;postgres&#34;)
 *             .type(&#34;database&#34;)
 *             .build());
 * 
 *         var postgres = new SecretBackendConnection(&#34;postgres&#34;, SecretBackendConnectionArgs.builder()        
 *             .backend(db.path())
 *             .name(&#34;postgres&#34;)
 *             .allowedRoles(            
 *                 &#34;dev&#34;,
 *                 &#34;prod&#34;)
 *             .postgresql(SecretBackendConnectionPostgresqlArgs.builder()
 *                 .connectionUrl(&#34;postgres://username:password@host:port/database&#34;)
 *                 .build())
 *             .build());
 * 
 *     }
 * }
 * ```
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Import
 * 
 * Database secret backend connections can be imported using the `backend`, `/config/`, and the `name` e.g.
 * 
 * ```sh
 * $ pulumi import vault:database/secretBackendConnection:SecretBackendConnection example postgres/config/postgres
 * ```
 * 
 */
@ResourceType(type="vault:database/secretBackendConnection:SecretBackendConnection")
public class SecretBackendConnection extends com.pulumi.resources.CustomResource {
    /**
     * A list of roles that are allowed to use this
     * connection.
     * 
     */
    @Export(name="allowedRoles", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> allowedRoles;

    /**
     * @return A list of roles that are allowed to use this
     * connection.
     * 
     */
    public Output<Optional<List<String>>> allowedRoles() {
        return Codegen.optional(this.allowedRoles);
    }
    /**
     * The unique name of the Vault mount to configure.
     * 
     */
    @Export(name="backend", refs={String.class}, tree="[0]")
    private Output<String> backend;

    /**
     * @return The unique name of the Vault mount to configure.
     * 
     */
    public Output<String> backend() {
        return this.backend;
    }
    /**
     * A nested block containing configuration options for Cassandra connections.
     * 
     */
    @Export(name="cassandra", refs={SecretBackendConnectionCassandra.class}, tree="[0]")
    private Output</* @Nullable */ SecretBackendConnectionCassandra> cassandra;

    /**
     * @return A nested block containing configuration options for Cassandra connections.
     * 
     */
    public Output<Optional<SecretBackendConnectionCassandra>> cassandra() {
        return Codegen.optional(this.cassandra);
    }
    /**
     * A nested block containing configuration options for Couchbase connections.
     * 
     */
    @Export(name="couchbase", refs={SecretBackendConnectionCouchbase.class}, tree="[0]")
    private Output</* @Nullable */ SecretBackendConnectionCouchbase> couchbase;

    /**
     * @return A nested block containing configuration options for Couchbase connections.
     * 
     */
    public Output<Optional<SecretBackendConnectionCouchbase>> couchbase() {
        return Codegen.optional(this.couchbase);
    }
    /**
     * A map of sensitive data to pass to the endpoint. Useful for templated connection strings.
     * 
     */
    @Export(name="data", refs={Map.class,String.class,Object.class}, tree="[0,1,2]")
    private Output</* @Nullable */ Map<String,Object>> data;

    /**
     * @return A map of sensitive data to pass to the endpoint. Useful for templated connection strings.
     * 
     */
    public Output<Optional<Map<String,Object>>> data() {
        return Codegen.optional(this.data);
    }
    /**
     * A nested block containing configuration options for Elasticsearch connections.
     * 
     */
    @Export(name="elasticsearch", refs={SecretBackendConnectionElasticsearch.class}, tree="[0]")
    private Output</* @Nullable */ SecretBackendConnectionElasticsearch> elasticsearch;

    /**
     * @return A nested block containing configuration options for Elasticsearch connections.
     * 
     */
    public Output<Optional<SecretBackendConnectionElasticsearch>> elasticsearch() {
        return Codegen.optional(this.elasticsearch);
    }
    /**
     * A nested block containing configuration options for SAP HanaDB connections.
     * 
     */
    @Export(name="hana", refs={SecretBackendConnectionHana.class}, tree="[0]")
    private Output</* @Nullable */ SecretBackendConnectionHana> hana;

    /**
     * @return A nested block containing configuration options for SAP HanaDB connections.
     * 
     */
    public Output<Optional<SecretBackendConnectionHana>> hana() {
        return Codegen.optional(this.hana);
    }
    /**
     * A nested block containing configuration options for InfluxDB connections.
     * 
     */
    @Export(name="influxdb", refs={SecretBackendConnectionInfluxdb.class}, tree="[0]")
    private Output</* @Nullable */ SecretBackendConnectionInfluxdb> influxdb;

    /**
     * @return A nested block containing configuration options for InfluxDB connections.
     * 
     */
    public Output<Optional<SecretBackendConnectionInfluxdb>> influxdb() {
        return Codegen.optional(this.influxdb);
    }
    /**
     * A nested block containing configuration options for MongoDB connections.
     * 
     */
    @Export(name="mongodb", refs={SecretBackendConnectionMongodb.class}, tree="[0]")
    private Output</* @Nullable */ SecretBackendConnectionMongodb> mongodb;

    /**
     * @return A nested block containing configuration options for MongoDB connections.
     * 
     */
    public Output<Optional<SecretBackendConnectionMongodb>> mongodb() {
        return Codegen.optional(this.mongodb);
    }
    /**
     * A nested block containing configuration options for MongoDB Atlas connections.
     * 
     */
    @Export(name="mongodbatlas", refs={SecretBackendConnectionMongodbatlas.class}, tree="[0]")
    private Output</* @Nullable */ SecretBackendConnectionMongodbatlas> mongodbatlas;

    /**
     * @return A nested block containing configuration options for MongoDB Atlas connections.
     * 
     */
    public Output<Optional<SecretBackendConnectionMongodbatlas>> mongodbatlas() {
        return Codegen.optional(this.mongodbatlas);
    }
    /**
     * A nested block containing configuration options for MSSQL connections.
     * 
     */
    @Export(name="mssql", refs={SecretBackendConnectionMssql.class}, tree="[0]")
    private Output</* @Nullable */ SecretBackendConnectionMssql> mssql;

    /**
     * @return A nested block containing configuration options for MSSQL connections.
     * 
     */
    public Output<Optional<SecretBackendConnectionMssql>> mssql() {
        return Codegen.optional(this.mssql);
    }
    /**
     * A nested block containing configuration options for MySQL connections.
     * 
     */
    @Export(name="mysql", refs={SecretBackendConnectionMysql.class}, tree="[0]")
    private Output</* @Nullable */ SecretBackendConnectionMysql> mysql;

    /**
     * @return A nested block containing configuration options for MySQL connections.
     * 
     */
    public Output<Optional<SecretBackendConnectionMysql>> mysql() {
        return Codegen.optional(this.mysql);
    }
    /**
     * A nested block containing configuration options for Aurora MySQL connections.
     * 
     */
    @Export(name="mysqlAurora", refs={SecretBackendConnectionMysqlAurora.class}, tree="[0]")
    private Output</* @Nullable */ SecretBackendConnectionMysqlAurora> mysqlAurora;

    /**
     * @return A nested block containing configuration options for Aurora MySQL connections.
     * 
     */
    public Output<Optional<SecretBackendConnectionMysqlAurora>> mysqlAurora() {
        return Codegen.optional(this.mysqlAurora);
    }
    /**
     * A nested block containing configuration options for legacy MySQL connections.
     * 
     */
    @Export(name="mysqlLegacy", refs={SecretBackendConnectionMysqlLegacy.class}, tree="[0]")
    private Output</* @Nullable */ SecretBackendConnectionMysqlLegacy> mysqlLegacy;

    /**
     * @return A nested block containing configuration options for legacy MySQL connections.
     * 
     */
    public Output<Optional<SecretBackendConnectionMysqlLegacy>> mysqlLegacy() {
        return Codegen.optional(this.mysqlLegacy);
    }
    /**
     * A nested block containing configuration options for RDS MySQL connections.
     * 
     */
    @Export(name="mysqlRds", refs={SecretBackendConnectionMysqlRds.class}, tree="[0]")
    private Output</* @Nullable */ SecretBackendConnectionMysqlRds> mysqlRds;

    /**
     * @return A nested block containing configuration options for RDS MySQL connections.
     * 
     */
    public Output<Optional<SecretBackendConnectionMysqlRds>> mysqlRds() {
        return Codegen.optional(this.mysqlRds);
    }
    /**
     * A unique name to give the database connection.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return A unique name to give the database connection.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * The namespace to provision the resource in.
     * The value should not contain leading or trailing forward slashes.
     * The `namespace` is always relative to the provider&#39;s configured namespace.
     * *Available only for Vault Enterprise*.
     * 
     */
    @Export(name="namespace", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> namespace;

    /**
     * @return The namespace to provision the resource in.
     * The value should not contain leading or trailing forward slashes.
     * The `namespace` is always relative to the provider&#39;s configured namespace.
     * *Available only for Vault Enterprise*.
     * 
     */
    public Output<Optional<String>> namespace() {
        return Codegen.optional(this.namespace);
    }
    /**
     * A nested block containing configuration options for Oracle connections.
     * 
     */
    @Export(name="oracle", refs={SecretBackendConnectionOracle.class}, tree="[0]")
    private Output</* @Nullable */ SecretBackendConnectionOracle> oracle;

    /**
     * @return A nested block containing configuration options for Oracle connections.
     * 
     */
    public Output<Optional<SecretBackendConnectionOracle>> oracle() {
        return Codegen.optional(this.oracle);
    }
    /**
     * Specifies the name of the plugin to use.
     * 
     */
    @Export(name="pluginName", refs={String.class}, tree="[0]")
    private Output<String> pluginName;

    /**
     * @return Specifies the name of the plugin to use.
     * 
     */
    public Output<String> pluginName() {
        return this.pluginName;
    }
    /**
     * A nested block containing configuration options for PostgreSQL connections.
     * 
     */
    @Export(name="postgresql", refs={SecretBackendConnectionPostgresql.class}, tree="[0]")
    private Output</* @Nullable */ SecretBackendConnectionPostgresql> postgresql;

    /**
     * @return A nested block containing configuration options for PostgreSQL connections.
     * 
     */
    public Output<Optional<SecretBackendConnectionPostgresql>> postgresql() {
        return Codegen.optional(this.postgresql);
    }
    /**
     * A nested block containing configuration options for Redis connections.
     * 
     */
    @Export(name="redis", refs={SecretBackendConnectionRedis.class}, tree="[0]")
    private Output</* @Nullable */ SecretBackendConnectionRedis> redis;

    /**
     * @return A nested block containing configuration options for Redis connections.
     * 
     */
    public Output<Optional<SecretBackendConnectionRedis>> redis() {
        return Codegen.optional(this.redis);
    }
    /**
     * A nested block containing configuration options for Redis ElastiCache connections.
     * 
     * Exactly one of the nested blocks of configuration options must be supplied.
     * 
     */
    @Export(name="redisElasticache", refs={SecretBackendConnectionRedisElasticache.class}, tree="[0]")
    private Output</* @Nullable */ SecretBackendConnectionRedisElasticache> redisElasticache;

    /**
     * @return A nested block containing configuration options for Redis ElastiCache connections.
     * 
     * Exactly one of the nested blocks of configuration options must be supplied.
     * 
     */
    public Output<Optional<SecretBackendConnectionRedisElasticache>> redisElasticache() {
        return Codegen.optional(this.redisElasticache);
    }
    /**
     * Connection parameters for the redshift-database-plugin plugin.
     * 
     */
    @Export(name="redshift", refs={SecretBackendConnectionRedshift.class}, tree="[0]")
    private Output</* @Nullable */ SecretBackendConnectionRedshift> redshift;

    /**
     * @return Connection parameters for the redshift-database-plugin plugin.
     * 
     */
    public Output<Optional<SecretBackendConnectionRedshift>> redshift() {
        return Codegen.optional(this.redshift);
    }
    /**
     * A list of database statements to be executed to rotate the root user&#39;s credentials.
     * 
     */
    @Export(name="rootRotationStatements", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> rootRotationStatements;

    /**
     * @return A list of database statements to be executed to rotate the root user&#39;s credentials.
     * 
     */
    public Output<Optional<List<String>>> rootRotationStatements() {
        return Codegen.optional(this.rootRotationStatements);
    }
    /**
     * A nested block containing configuration options for Snowflake connections.
     * 
     */
    @Export(name="snowflake", refs={SecretBackendConnectionSnowflake.class}, tree="[0]")
    private Output</* @Nullable */ SecretBackendConnectionSnowflake> snowflake;

    /**
     * @return A nested block containing configuration options for Snowflake connections.
     * 
     */
    public Output<Optional<SecretBackendConnectionSnowflake>> snowflake() {
        return Codegen.optional(this.snowflake);
    }
    /**
     * Whether the connection should be verified on
     * initial configuration or not.
     * 
     */
    @Export(name="verifyConnection", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> verifyConnection;

    /**
     * @return Whether the connection should be verified on
     * initial configuration or not.
     * 
     */
    public Output<Optional<Boolean>> verifyConnection() {
        return Codegen.optional(this.verifyConnection);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public SecretBackendConnection(String name) {
        this(name, SecretBackendConnectionArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public SecretBackendConnection(String name, SecretBackendConnectionArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public SecretBackendConnection(String name, SecretBackendConnectionArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("vault:database/secretBackendConnection:SecretBackendConnection", name, args == null ? SecretBackendConnectionArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private SecretBackendConnection(String name, Output<String> id, @Nullable SecretBackendConnectionState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("vault:database/secretBackendConnection:SecretBackendConnection", name, state, makeResourceOptions(options, id));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
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
    public static SecretBackendConnection get(String name, Output<String> id, @Nullable SecretBackendConnectionState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new SecretBackendConnection(name, id, state, options);
    }
}
