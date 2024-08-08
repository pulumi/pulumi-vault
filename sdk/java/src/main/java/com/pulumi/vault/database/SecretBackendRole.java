// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vault.database;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.vault.Utilities;
import com.pulumi.vault.database.SecretBackendRoleArgs;
import com.pulumi.vault.database.inputs.SecretBackendRoleState;
import java.lang.Integer;
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
 * <pre>
 * {@code
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
 * import com.pulumi.vault.database.SecretBackendRole;
 * import com.pulumi.vault.database.SecretBackendRoleArgs;
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
 *         var db = new Mount("db", MountArgs.builder()
 *             .path("postgres")
 *             .type("database")
 *             .build());
 * 
 *         var postgres = new SecretBackendConnection("postgres", SecretBackendConnectionArgs.builder()
 *             .backend(db.path())
 *             .name("postgres")
 *             .allowedRoles(            
 *                 "dev",
 *                 "prod")
 *             .postgresql(SecretBackendConnectionPostgresqlArgs.builder()
 *                 .connectionUrl("postgres://username:password{@literal @}host:port/database")
 *                 .build())
 *             .build());
 * 
 *         var role = new SecretBackendRole("role", SecretBackendRoleArgs.builder()
 *             .backend(db.path())
 *             .name("dev")
 *             .dbName(postgres.name())
 *             .creationStatements("CREATE ROLE \"{{name}}\" WITH LOGIN PASSWORD '{{password}}' VALID UNTIL '{{expiration}}';")
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Import
 * 
 * Database secret backend roles can be imported using the `backend`, `/roles/`, and the `name` e.g.
 * 
 * ```sh
 * $ pulumi import vault:database/secretBackendRole:SecretBackendRole example postgres/roles/my-role
 * ```
 * 
 */
@ResourceType(type="vault:database/secretBackendRole:SecretBackendRole")
public class SecretBackendRole extends com.pulumi.resources.CustomResource {
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
     * The database statements to execute when
     * creating a user.
     * 
     */
    @Export(name="creationStatements", refs={List.class,String.class}, tree="[0,1]")
    private Output<List<String>> creationStatements;

    /**
     * @return The database statements to execute when
     * creating a user.
     * 
     */
    public Output<List<String>> creationStatements() {
        return this.creationStatements;
    }
    /**
     * Specifies the configuration
     * for the given `credential_type`.
     * 
     * The following options are available for each `credential_type` value:
     * 
     */
    @Export(name="credentialConfig", refs={Map.class,String.class,Object.class}, tree="[0,1,2]")
    private Output</* @Nullable */ Map<String,Object>> credentialConfig;

    /**
     * @return Specifies the configuration
     * for the given `credential_type`.
     * 
     * The following options are available for each `credential_type` value:
     * 
     */
    public Output<Optional<Map<String,Object>>> credentialConfig() {
        return Codegen.optional(this.credentialConfig);
    }
    /**
     * Specifies the type of credential that
     * will be generated for the role. Options include: `password`, `rsa_private_key`, `client_certificate`.
     * See the plugin&#39;s API page for credential types supported by individual databases.
     * 
     */
    @Export(name="credentialType", refs={String.class}, tree="[0]")
    private Output<String> credentialType;

    /**
     * @return Specifies the type of credential that
     * will be generated for the role. Options include: `password`, `rsa_private_key`, `client_certificate`.
     * See the plugin&#39;s API page for credential types supported by individual databases.
     * 
     */
    public Output<String> credentialType() {
        return this.credentialType;
    }
    /**
     * The unique name of the database connection to use for
     * the role.
     * 
     */
    @Export(name="dbName", refs={String.class}, tree="[0]")
    private Output<String> dbName;

    /**
     * @return The unique name of the database connection to use for
     * the role.
     * 
     */
    public Output<String> dbName() {
        return this.dbName;
    }
    /**
     * The default number of seconds for leases for this
     * role.
     * 
     */
    @Export(name="defaultTtl", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> defaultTtl;

    /**
     * @return The default number of seconds for leases for this
     * role.
     * 
     */
    public Output<Optional<Integer>> defaultTtl() {
        return Codegen.optional(this.defaultTtl);
    }
    /**
     * The maximum number of seconds for leases for this
     * role.
     * 
     */
    @Export(name="maxTtl", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> maxTtl;

    /**
     * @return The maximum number of seconds for leases for this
     * role.
     * 
     */
    public Output<Optional<Integer>> maxTtl() {
        return Codegen.optional(this.maxTtl);
    }
    /**
     * A unique name to give the role.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return A unique name to give the role.
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
     * The database statements to execute when
     * renewing a user.
     * 
     */
    @Export(name="renewStatements", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> renewStatements;

    /**
     * @return The database statements to execute when
     * renewing a user.
     * 
     */
    public Output<Optional<List<String>>> renewStatements() {
        return Codegen.optional(this.renewStatements);
    }
    /**
     * The database statements to execute when
     * revoking a user.
     * 
     */
    @Export(name="revocationStatements", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> revocationStatements;

    /**
     * @return The database statements to execute when
     * revoking a user.
     * 
     */
    public Output<Optional<List<String>>> revocationStatements() {
        return Codegen.optional(this.revocationStatements);
    }
    /**
     * The database statements to execute when
     * rolling back creation due to an error.
     * 
     */
    @Export(name="rollbackStatements", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> rollbackStatements;

    /**
     * @return The database statements to execute when
     * rolling back creation due to an error.
     * 
     */
    public Output<Optional<List<String>>> rollbackStatements() {
        return Codegen.optional(this.rollbackStatements);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public SecretBackendRole(String name) {
        this(name, SecretBackendRoleArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public SecretBackendRole(String name, SecretBackendRoleArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public SecretBackendRole(String name, SecretBackendRoleArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("vault:database/secretBackendRole:SecretBackendRole", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()));
    }

    private SecretBackendRole(String name, Output<String> id, @Nullable SecretBackendRoleState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("vault:database/secretBackendRole:SecretBackendRole", name, state, makeResourceOptions(options, id));
    }

    private static SecretBackendRoleArgs makeArgs(SecretBackendRoleArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? SecretBackendRoleArgs.Empty : args;
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
    public static SecretBackendRole get(String name, Output<String> id, @Nullable SecretBackendRoleState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new SecretBackendRole(name, id, state, options);
    }
}
