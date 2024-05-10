// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vault.aws;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.vault.Utilities;
import com.pulumi.vault.aws.SecretBackendStaticRoleArgs;
import com.pulumi.vault.aws.inputs.SecretBackendStaticRoleState;
import java.lang.Integer;
import java.lang.String;
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
 * import com.pulumi.vault.aws.SecretBackend;
 * import com.pulumi.vault.aws.SecretBackendArgs;
 * import com.pulumi.vault.aws.SecretBackendStaticRole;
 * import com.pulumi.vault.aws.SecretBackendStaticRoleArgs;
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
 *         var aws = new SecretBackend("aws", SecretBackendArgs.builder()        
 *             .path("my-aws")
 *             .description("Obtain AWS credentials.")
 *             .build());
 * 
 *         var role = new SecretBackendStaticRole("role", SecretBackendStaticRoleArgs.builder()        
 *             .backend(aws.path())
 *             .name("test")
 *             .username("my-test-user")
 *             .rotationPeriod("3600")
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
 * AWS secret backend static role can be imported using the full path to the role
 * of the form: `&lt;mount_path&gt;/static-roles/&lt;role_name&gt;` e.g.
 * 
 * ```sh
 * $ pulumi import vault:aws/secretBackendStaticRole:SecretBackendStaticRole role aws/static-roles/example-role
 * ```
 * 
 */
@ResourceType(type="vault:aws/secretBackendStaticRole:SecretBackendStaticRole")
public class SecretBackendStaticRole extends com.pulumi.resources.CustomResource {
    /**
     * The unique path this backend should be mounted at. Must
     * not begin or end with a `/`. Defaults to `aws`
     * 
     */
    @Export(name="backend", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> backend;

    /**
     * @return The unique path this backend should be mounted at. Must
     * not begin or end with a `/`. Defaults to `aws`
     * 
     */
    public Output<Optional<String>> backend() {
        return Codegen.optional(this.backend);
    }
    /**
     * The name to identify this role within the backend.
     * Must be unique within the backend.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return The name to identify this role within the backend.
     * Must be unique within the backend.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * The namespace to provision the resource in.
     * The value should not contain leading or trailing forward slashes.
     * The `namespace` is always relative to the provider&#39;s configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
     * *Available only for Vault Enterprise*.
     * 
     */
    @Export(name="namespace", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> namespace;

    /**
     * @return The namespace to provision the resource in.
     * The value should not contain leading or trailing forward slashes.
     * The `namespace` is always relative to the provider&#39;s configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
     * *Available only for Vault Enterprise*.
     * 
     */
    public Output<Optional<String>> namespace() {
        return Codegen.optional(this.namespace);
    }
    /**
     * How often Vault should rotate the password of the user entry.
     * 
     */
    @Export(name="rotationPeriod", refs={Integer.class}, tree="[0]")
    private Output<Integer> rotationPeriod;

    /**
     * @return How often Vault should rotate the password of the user entry.
     * 
     */
    public Output<Integer> rotationPeriod() {
        return this.rotationPeriod;
    }
    /**
     * The username of the existing AWS IAM to manage password rotation for.
     * 
     */
    @Export(name="username", refs={String.class}, tree="[0]")
    private Output<String> username;

    /**
     * @return The username of the existing AWS IAM to manage password rotation for.
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
        super("vault:aws/secretBackendStaticRole:SecretBackendStaticRole", name, args == null ? SecretBackendStaticRoleArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private SecretBackendStaticRole(String name, Output<String> id, @Nullable SecretBackendStaticRoleState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("vault:aws/secretBackendStaticRole:SecretBackendStaticRole", name, state, makeResourceOptions(options, id));
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
