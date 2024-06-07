// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vault.kmip;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.vault.Utilities;
import com.pulumi.vault.kmip.SecretScopeArgs;
import com.pulumi.vault.kmip.inputs.SecretScopeState;
import java.lang.Boolean;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Manages KMIP Secret Scopes in a Vault server. This feature requires
 * Vault Enterprise. See the [Vault documentation](https://www.vaultproject.io/docs/secrets/kmip)
 * for more information.
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
 * import com.pulumi.vault.kmip.SecretBackend;
 * import com.pulumi.vault.kmip.SecretBackendArgs;
 * import com.pulumi.vault.kmip.SecretScope;
 * import com.pulumi.vault.kmip.SecretScopeArgs;
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
 *         var default_ = new SecretBackend("default", SecretBackendArgs.builder()
 *             .path("kmip")
 *             .description("Vault KMIP backend")
 *             .build());
 * 
 *         var dev = new SecretScope("dev", SecretScopeArgs.builder()
 *             .path(default_.path())
 *             .scope("dev")
 *             .force(true)
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
 * KMIP Secret scope can be imported using the `path`, e.g.
 * 
 * ```sh
 * $ pulumi import vault:kmip/secretScope:SecretScope dev kmip
 * ```
 * 
 */
@ResourceType(type="vault:kmip/secretScope:SecretScope")
public class SecretScope extends com.pulumi.resources.CustomResource {
    /**
     * Boolean field to force deletion even if there are managed objects in the scope.
     * 
     */
    @Export(name="force", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> force;

    /**
     * @return Boolean field to force deletion even if there are managed objects in the scope.
     * 
     */
    public Output<Optional<Boolean>> force() {
        return Codegen.optional(this.force);
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
     * The unique path this backend should be mounted at. Must
     * not begin or end with a `/`. Defaults to `kmip`.
     * 
     */
    @Export(name="path", refs={String.class}, tree="[0]")
    private Output<String> path;

    /**
     * @return The unique path this backend should be mounted at. Must
     * not begin or end with a `/`. Defaults to `kmip`.
     * 
     */
    public Output<String> path() {
        return this.path;
    }
    /**
     * Name of the scope.
     * 
     */
    @Export(name="scope", refs={String.class}, tree="[0]")
    private Output<String> scope;

    /**
     * @return Name of the scope.
     * 
     */
    public Output<String> scope() {
        return this.scope;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public SecretScope(String name) {
        this(name, SecretScopeArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public SecretScope(String name, SecretScopeArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public SecretScope(String name, SecretScopeArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("vault:kmip/secretScope:SecretScope", name, args == null ? SecretScopeArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private SecretScope(String name, Output<String> id, @Nullable SecretScopeState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("vault:kmip/secretScope:SecretScope", name, state, makeResourceOptions(options, id));
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
    public static SecretScope get(String name, Output<String> id, @Nullable SecretScopeState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new SecretScope(name, id, state, options);
    }
}
