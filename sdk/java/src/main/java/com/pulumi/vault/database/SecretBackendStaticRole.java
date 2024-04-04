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
import java.lang.Integer;
import java.lang.String;
import java.util.List;
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
 * import com.pulumi.vault.database.SecretBackendStaticRole;
 * import com.pulumi.vault.database.SecretBackendStaticRoleArgs;
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
 *             .allowedRoles(&#34;*&#34;)
 *             .postgresql(SecretBackendConnectionPostgresqlArgs.builder()
 *                 .connectionUrl(&#34;postgres://username:password@host:port/database&#34;)
 *                 .build())
 *             .build());
 * 
 *         // configure a static role with period-based rotations
 *         var periodRole = new SecretBackendStaticRole(&#34;periodRole&#34;, SecretBackendStaticRoleArgs.builder()        
 *             .backend(db.path())
 *             .dbName(postgres.name())
 *             .username(&#34;example&#34;)
 *             .rotationPeriod(&#34;3600&#34;)
 *             .rotationStatements(&#34;ALTER USER \&#34;{{name}}\&#34; WITH PASSWORD &#39;{{password}}&#39;;&#34;)
 *             .build());
 * 
 *         // configure a static role with schedule-based rotations
 *         var scheduleRole = new SecretBackendStaticRole(&#34;scheduleRole&#34;, SecretBackendStaticRoleArgs.builder()        
 *             .backend(db.path())
 *             .dbName(postgres.name())
 *             .username(&#34;example&#34;)
 *             .rotationSchedule(&#34;0 0 * * SAT&#34;)
 *             .rotationWindow(&#34;172800&#34;)
 *             .rotationStatements(&#34;ALTER USER \&#34;{{name}}\&#34; WITH PASSWORD &#39;{{password}}&#39;;&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
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
    public SecretBackendStaticRole(String name) {
        this(name, SecretBackendStaticRoleArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public SecretBackendStaticRole(String name, SecretBackendStaticRoleArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public SecretBackendStaticRole(String name, SecretBackendStaticRoleArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("vault:database/secretBackendStaticRole:SecretBackendStaticRole", name, args == null ? SecretBackendStaticRoleArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private SecretBackendStaticRole(String name, Output<String> id, @Nullable SecretBackendStaticRoleState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("vault:database/secretBackendStaticRole:SecretBackendStaticRole", name, state, makeResourceOptions(options, id));
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
    public static SecretBackendStaticRole get(String name, Output<String> id, @Nullable SecretBackendStaticRoleState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new SecretBackendStaticRole(name, id, state, options);
    }
}
