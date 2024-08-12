// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vault.terraformcloud;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.vault.Utilities;
import com.pulumi.vault.terraformcloud.SecretRoleArgs;
import com.pulumi.vault.terraformcloud.inputs.SecretRoleState;
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
 * import com.pulumi.vault.terraformcloud.SecretBackend;
 * import com.pulumi.vault.terraformcloud.SecretBackendArgs;
 * import com.pulumi.vault.terraformcloud.SecretRole;
 * import com.pulumi.vault.terraformcloud.SecretRoleArgs;
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
 *         var test = new SecretBackend("test", SecretBackendArgs.builder()
 *             .backend("terraform")
 *             .description("Manages the Terraform Cloud backend")
 *             .token("V0idfhi2iksSDU234ucdbi2nidsi...")
 *             .build());
 * 
 *         var example = new SecretRole("example", SecretRoleArgs.builder()
 *             .backend(test.backend())
 *             .name("test-role")
 *             .organization("example-organization-name")
 *             .teamId("team-ieF4isC...")
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
 * Terraform Cloud secret backend roles can be imported using the `backend`, `/roles/`, and the `name` e.g.
 * 
 * ```sh
 * $ pulumi import vault:terraformcloud/secretRole:SecretRole example terraform/roles/my-role
 * ```
 * 
 */
@ResourceType(type="vault:terraformcloud/secretRole:SecretRole")
public class SecretRole extends com.pulumi.resources.CustomResource {
    @Export(name="backend", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> backend;

    public Output<Optional<String>> backend() {
        return Codegen.optional(this.backend);
    }
    /**
     * Maximum TTL for leases associated with this role, in seconds.
     * 
     */
    @Export(name="maxTtl", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> maxTtl;

    /**
     * @return Maximum TTL for leases associated with this role, in seconds.
     * 
     */
    public Output<Optional<Integer>> maxTtl() {
        return Codegen.optional(this.maxTtl);
    }
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

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
    @Export(name="organization", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> organization;

    public Output<Optional<String>> organization() {
        return Codegen.optional(this.organization);
    }
    @Export(name="teamId", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> teamId;

    public Output<Optional<String>> teamId() {
        return Codegen.optional(this.teamId);
    }
    /**
     * Specifies the TTL for this role.
     * 
     */
    @Export(name="ttl", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> ttl;

    /**
     * @return Specifies the TTL for this role.
     * 
     */
    public Output<Optional<Integer>> ttl() {
        return Codegen.optional(this.ttl);
    }
    @Export(name="userId", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> userId;

    public Output<Optional<String>> userId() {
        return Codegen.optional(this.userId);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public SecretRole(java.lang.String name) {
        this(name, SecretRoleArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public SecretRole(java.lang.String name, @Nullable SecretRoleArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public SecretRole(java.lang.String name, @Nullable SecretRoleArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("vault:terraformcloud/secretRole:SecretRole", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private SecretRole(java.lang.String name, Output<java.lang.String> id, @Nullable SecretRoleState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("vault:terraformcloud/secretRole:SecretRole", name, state, makeResourceOptions(options, id), false);
    }

    private static SecretRoleArgs makeArgs(@Nullable SecretRoleArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? SecretRoleArgs.Empty : args;
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
    public static SecretRole get(java.lang.String name, Output<java.lang.String> id, @Nullable SecretRoleState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new SecretRole(name, id, state, options);
    }
}
