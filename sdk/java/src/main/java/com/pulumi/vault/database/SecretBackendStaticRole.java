// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vault.database;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.vault.Utilities;
import com.pulumi.vault.database.SecretBackendStaticRoleArgs;
import com.pulumi.vault.database.inputs.SecretBackendStaticRoleState;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Creates a Database Secret Backend static role in Vault. Database secret backend
 * static roles can be used to manage 1-to-1 mapping of a Vault Role to a user in a
 * database for the database.
 * 
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
 * import com.pulumi.vault.database.SecretBackendStaticRole;
 * import com.pulumi.vault.database.SecretBackendStaticRoleArgs;
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
 *         var db = new Mount("db", MountArgs.builder()
 *             .path("postgres")
 *             .type("database")
 *             .build());
 * 
 *         var postgres = new SecretBackendConnection("postgres", SecretBackendConnectionArgs.builder()
 *             .backend(db.path())
 *             .name("postgres")
 *             .allowedRoles("*")
 *             .postgresql(SecretBackendConnectionPostgresqlArgs.builder()
 *                 .connectionUrl("postgres://username:password}{@literal @}{@code host:port/database")
 *                 .build())
 *             .build());
 * 
 *         // configure a static role with period-based rotations
 *         var periodRole = new SecretBackendStaticRole("periodRole", SecretBackendStaticRoleArgs.builder()
 *             .backend(db.path())
 *             .name("my-period-role")
 *             .dbName(postgres.name())
 *             .username("example")
 *             .rotationPeriod(3600)
 *             .rotationStatements("ALTER USER \"}{{{@code name}}}{@code \" WITH PASSWORD '}{{{@code password}}}{@code ';")
 *             .build());
 * 
 *         // configure a static role with schedule-based rotations
 *         var scheduleRole = new SecretBackendStaticRole("scheduleRole", SecretBackendStaticRoleArgs.builder()
 *             .backend(db.path())
 *             .name("my-schedule-role")
 *             .dbName(postgres.name())
 *             .username("example")
 *             .rotationSchedule("0 0 * * SAT")
 *             .rotationWindow(172800)
 *             .rotationStatements("ALTER USER \"}{{{@code name}}}{@code \" WITH PASSWORD '}{{{@code password}}}{@code ';")
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
 * Database secret backend static roles can be imported using the `backend`, `/static-roles/`, and the `name` e.g.
 * 
 * ```sh
 * $ pulumi import vault:database/secretBackendStaticRole:SecretBackendStaticRole example postgres/static-roles/my-role
 * ```
 * 
 */
@ResourceType(type="vault:database/secretBackendStaticRole:SecretBackendStaticRole")
public class SecretBackendStaticRole extends com.pulumi.resources.CustomResource {
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
    @Export(name="credentialConfig", refs={Map.class,String.class}, tree="[0,1,1]")
    private Output</* @Nullable */ Map<String,String>> credentialConfig;

    public Output<Optional<Map<String,String>>> credentialConfig() {
        return Codegen.optional(this.credentialConfig);
    }
    /**
     * The credential type for the user, can be one of &#34;password&#34;, &#34;rsa_private_key&#34; or &#34;client_certificate&#34;.The configuration
     * can be done in `credential_config`.
     * 
     */
    @Export(name="credentialType", refs={String.class}, tree="[0]")
    private Output<String> credentialType;

    /**
     * @return The credential type for the user, can be one of &#34;password&#34;, &#34;rsa_private_key&#34; or &#34;client_certificate&#34;.The configuration
     * can be done in `credential_config`.
     * 
     */
    public Output<String> credentialType() {
        return this.credentialType;
    }
    /**
     * The unique name of the database connection to use for the static role.
     * 
     */
    @Export(name="dbName", refs={String.class}, tree="[0]")
    private Output<String> dbName;

    /**
     * @return The unique name of the database connection to use for the static role.
     * 
     */
    public Output<String> dbName() {
        return this.dbName;
    }
    /**
     * A unique name to give the static role.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return A unique name to give the static role.
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
     * The amount of time Vault should wait before rotating the password, in seconds.
     * Mutually exclusive with `rotation_schedule`.
     * 
     */
    @Export(name="rotationPeriod", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> rotationPeriod;

    /**
     * @return The amount of time Vault should wait before rotating the password, in seconds.
     * Mutually exclusive with `rotation_schedule`.
     * 
     */
    public Output<Optional<Integer>> rotationPeriod() {
        return Codegen.optional(this.rotationPeriod);
    }
    /**
     * A cron-style string that will define the schedule on which rotations should occur.
     * Mutually exclusive with `rotation_period`.
     * 
     * **Warning**: The `rotation_period` and `rotation_schedule` fields are
     * mutually exclusive. One of them must be set but not both.
     * 
     */
    @Export(name="rotationSchedule", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> rotationSchedule;

    /**
     * @return A cron-style string that will define the schedule on which rotations should occur.
     * Mutually exclusive with `rotation_period`.
     * 
     * **Warning**: The `rotation_period` and `rotation_schedule` fields are
     * mutually exclusive. One of them must be set but not both.
     * 
     */
    public Output<Optional<String>> rotationSchedule() {
        return Codegen.optional(this.rotationSchedule);
    }
    /**
     * Database statements to execute to rotate the password for the configured database user.
     * 
     */
    @Export(name="rotationStatements", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> rotationStatements;

    /**
     * @return Database statements to execute to rotate the password for the configured database user.
     * 
     */
    public Output<Optional<List<String>>> rotationStatements() {
        return Codegen.optional(this.rotationStatements);
    }
    /**
     * The amount of time, in seconds, in which rotations are allowed to occur starting
     * from a given `rotation_schedule`.
     * 
     */
    @Export(name="rotationWindow", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> rotationWindow;

    /**
     * @return The amount of time, in seconds, in which rotations are allowed to occur starting
     * from a given `rotation_schedule`.
     * 
     */
    public Output<Optional<Integer>> rotationWindow() {
        return Codegen.optional(this.rotationWindow);
    }
    /**
     * The password corresponding to the username in the database.
     * Required when using the Rootless Password Rotation workflow for static roles. Only enabled for
     * select DB engines (Postgres). Requires Vault 1.18+ Enterprise.
     * 
     */
    @Export(name="selfManagedPassword", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> selfManagedPassword;

    /**
     * @return The password corresponding to the username in the database.
     * Required when using the Rootless Password Rotation workflow for static roles. Only enabled for
     * select DB engines (Postgres). Requires Vault 1.18+ Enterprise.
     * 
     */
    public Output<Optional<String>> selfManagedPassword() {
        return Codegen.optional(this.selfManagedPassword);
    }
    /**
     * If set to true, Vault will skip the
     * initial secret rotation on import. Requires Vault 1.18+ Enterprise.
     * 
     */
    @Export(name="skipImportRotation", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> skipImportRotation;

    /**
     * @return If set to true, Vault will skip the
     * initial secret rotation on import. Requires Vault 1.18+ Enterprise.
     * 
     */
    public Output<Optional<Boolean>> skipImportRotation() {
        return Codegen.optional(this.skipImportRotation);
    }
    /**
     * The database username that this static role corresponds to.
     * 
     */
    @Export(name="username", refs={String.class}, tree="[0]")
    private Output<String> username;

    /**
     * @return The database username that this static role corresponds to.
     * 
     */
    public Output<String> username() {
        return this.username;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public SecretBackendStaticRole(java.lang.String name) {
        this(name, SecretBackendStaticRoleArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public SecretBackendStaticRole(java.lang.String name, SecretBackendStaticRoleArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public SecretBackendStaticRole(java.lang.String name, SecretBackendStaticRoleArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("vault:database/secretBackendStaticRole:SecretBackendStaticRole", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private SecretBackendStaticRole(java.lang.String name, Output<java.lang.String> id, @Nullable SecretBackendStaticRoleState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("vault:database/secretBackendStaticRole:SecretBackendStaticRole", name, state, makeResourceOptions(options, id), false);
    }

    private static SecretBackendStaticRoleArgs makeArgs(SecretBackendStaticRoleArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? SecretBackendStaticRoleArgs.Empty : args;
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<java.lang.String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .additionalSecretOutputs(List.of(
                "selfManagedPassword"
            ))
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
    public static SecretBackendStaticRole get(java.lang.String name, Output<java.lang.String> id, @Nullable SecretBackendStaticRoleState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new SecretBackendStaticRole(name, id, state, options);
    }
}
