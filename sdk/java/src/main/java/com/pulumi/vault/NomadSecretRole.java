// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vault;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.vault.NomadSecretRoleArgs;
import com.pulumi.vault.Utilities;
import com.pulumi.vault.inputs.NomadSecretRoleState;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
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
 * import com.pulumi.vault.NomadSecretBackend;
 * import com.pulumi.vault.NomadSecretBackendArgs;
 * import com.pulumi.vault.NomadSecretRole;
 * import com.pulumi.vault.NomadSecretRoleArgs;
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
 *         var config = new NomadSecretBackend("config", NomadSecretBackendArgs.builder()
 *             .backend("nomad")
 *             .description("test description")
 *             .defaultLeaseTtlSeconds(3600)
 *             .maxLeaseTtlSeconds(7200)
 *             .address("https://127.0.0.1:4646")
 *             .token("ae20ceaa-...")
 *             .build());
 * 
 *         var test = new NomadSecretRole("test", NomadSecretRoleArgs.builder()
 *             .backend(config.backend())
 *             .role("test")
 *             .type("client")
 *             .policies("readonly")
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
 * Nomad secret role can be imported using the `backend`, e.g.
 * 
 * ```sh
 * $ pulumi import vault:index/nomadSecretRole:NomadSecretRole bob nomad/role/bob
 * ```
 * 
 */
@ResourceType(type="vault:index/nomadSecretRole:NomadSecretRole")
public class NomadSecretRole extends com.pulumi.resources.CustomResource {
    /**
     * The unique path this backend should be mounted at.
     * 
     */
    @Export(name="backend", refs={String.class}, tree="[0]")
    private Output<String> backend;

    /**
     * @return The unique path this backend should be mounted at.
     * 
     */
    public Output<String> backend() {
        return this.backend;
    }
    /**
     * Specifies if the generated token should be global. Defaults to
     * false.
     * 
     */
    @Export(name="global", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> global;

    /**
     * @return Specifies if the generated token should be global. Defaults to
     * false.
     * 
     */
    public Output<Boolean> global() {
        return this.global;
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
     * List of policies attached to the generated token. This setting is only used
     * when `type` is &#39;client&#39;.
     * 
     */
    @Export(name="policies", refs={List.class,String.class}, tree="[0,1]")
    private Output<List<String>> policies;

    /**
     * @return List of policies attached to the generated token. This setting is only used
     * when `type` is &#39;client&#39;.
     * 
     */
    public Output<List<String>> policies() {
        return this.policies;
    }
    /**
     * The name to identify this role within the backend.
     * Must be unique within the backend.
     * 
     */
    @Export(name="role", refs={String.class}, tree="[0]")
    private Output<String> role;

    /**
     * @return The name to identify this role within the backend.
     * Must be unique within the backend.
     * 
     */
    public Output<String> role() {
        return this.role;
    }
    /**
     * Specifies the type of token to create when using this role. Valid
     * settings are &#39;client&#39; and &#39;management&#39;. Defaults to &#39;client&#39;.
     * 
     */
    @Export(name="type", refs={String.class}, tree="[0]")
    private Output<String> type;

    /**
     * @return Specifies the type of token to create when using this role. Valid
     * settings are &#39;client&#39; and &#39;management&#39;. Defaults to &#39;client&#39;.
     * 
     */
    public Output<String> type() {
        return this.type;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public NomadSecretRole(java.lang.String name) {
        this(name, NomadSecretRoleArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public NomadSecretRole(java.lang.String name, NomadSecretRoleArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public NomadSecretRole(java.lang.String name, NomadSecretRoleArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("vault:index/nomadSecretRole:NomadSecretRole", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private NomadSecretRole(java.lang.String name, Output<java.lang.String> id, @Nullable NomadSecretRoleState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("vault:index/nomadSecretRole:NomadSecretRole", name, state, makeResourceOptions(options, id), false);
    }

    private static NomadSecretRoleArgs makeArgs(NomadSecretRoleArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? NomadSecretRoleArgs.Empty : args;
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
    public static NomadSecretRole get(java.lang.String name, Output<java.lang.String> id, @Nullable NomadSecretRoleState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new NomadSecretRole(name, id, state, options);
    }
}
