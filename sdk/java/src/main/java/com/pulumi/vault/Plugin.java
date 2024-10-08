// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vault;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.vault.PluginArgs;
import com.pulumi.vault.Utilities;
import com.pulumi.vault.inputs.PluginState;
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
 * import com.pulumi.vault.Plugin;
 * import com.pulumi.vault.PluginArgs;
 * import com.pulumi.vault.AuthBackend;
 * import com.pulumi.vault.AuthBackendArgs;
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
 *         var jwt = new Plugin("jwt", PluginArgs.builder()
 *             .type("auth")
 *             .name("jwt")
 *             .command("vault-plugin-auth-jwt")
 *             .version("v0.17.0")
 *             .sha256("6bd0a803ed742aa3ce35e4fa23d2c8d550e6c1567bf63410cec489c28b68b0fc")
 *             .envs("HTTP_PROXY=http://proxy.example.com:8080")
 *             .build());
 * 
 *         var jwtAuth = new AuthBackend("jwtAuth", AuthBackendArgs.builder()
 *             .type(jwt.name())
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
 * Plugins can be imported using `:type/name/:name` or `:type/version/:version/name/:name` as the ID if the version is non-empty, e.g.
 * 
 * ```sh
 * $ pulumi import vault:index/plugin:Plugin jwt auth/name/jwt
 * ```
 * ```sh
 * $ pulumi import vault:index/plugin:Plugin jwt auth/version/v0.17.0/name/jwt
 * ```
 * 
 */
@ResourceType(type="vault:index/plugin:Plugin")
public class Plugin extends com.pulumi.resources.CustomResource {
    /**
     * List of additional args to pass to the plugin.
     * 
     */
    @Export(name="args", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> args;

    /**
     * @return List of additional args to pass to the plugin.
     * 
     */
    public Output<Optional<List<String>>> args() {
        return Codegen.optional(this.args);
    }
    /**
     * Command to execute the plugin, relative to the server&#39;s configured `plugin_directory`.
     * 
     */
    @Export(name="command", refs={String.class}, tree="[0]")
    private Output<String> command;

    /**
     * @return Command to execute the plugin, relative to the server&#39;s configured `plugin_directory`.
     * 
     */
    public Output<String> command() {
        return this.command;
    }
    /**
     * List of additional environment variables to run the plugin with in KEY=VALUE form.
     * 
     */
    @Export(name="envs", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> envs;

    /**
     * @return List of additional environment variables to run the plugin with in KEY=VALUE form.
     * 
     */
    public Output<Optional<List<String>>> envs() {
        return Codegen.optional(this.envs);
    }
    /**
     * Name of the plugin.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return Name of the plugin.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * Specifies OCI image to run. If specified, setting
     * `command`, `args`, and `env` will update the container&#39;s entrypoint, args, and
     * environment variables (append-only) respectively.
     * 
     */
    @Export(name="ociImage", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> ociImage;

    /**
     * @return Specifies OCI image to run. If specified, setting
     * `command`, `args`, and `env` will update the container&#39;s entrypoint, args, and
     * environment variables (append-only) respectively.
     * 
     */
    public Output<Optional<String>> ociImage() {
        return Codegen.optional(this.ociImage);
    }
    /**
     * Vault plugin runtime to use if `oci_image` is specified.
     * 
     */
    @Export(name="runtime", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> runtime;

    /**
     * @return Vault plugin runtime to use if `oci_image` is specified.
     * 
     */
    public Output<Optional<String>> runtime() {
        return Codegen.optional(this.runtime);
    }
    /**
     * SHA256 sum of the plugin binary.
     * 
     */
    @Export(name="sha256", refs={String.class}, tree="[0]")
    private Output<String> sha256;

    /**
     * @return SHA256 sum of the plugin binary.
     * 
     */
    public Output<String> sha256() {
        return this.sha256;
    }
    /**
     * Type of plugin; one of &#34;auth&#34;, &#34;secret&#34;, or &#34;database&#34;.
     * 
     */
    @Export(name="type", refs={String.class}, tree="[0]")
    private Output<String> type;

    /**
     * @return Type of plugin; one of &#34;auth&#34;, &#34;secret&#34;, or &#34;database&#34;.
     * 
     */
    public Output<String> type() {
        return this.type;
    }
    /**
     * Semantic version of the plugin.
     * 
     */
    @Export(name="version", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> version;

    /**
     * @return Semantic version of the plugin.
     * 
     */
    public Output<Optional<String>> version() {
        return Codegen.optional(this.version);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Plugin(java.lang.String name) {
        this(name, PluginArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Plugin(java.lang.String name, PluginArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Plugin(java.lang.String name, PluginArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("vault:index/plugin:Plugin", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private Plugin(java.lang.String name, Output<java.lang.String> id, @Nullable PluginState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("vault:index/plugin:Plugin", name, state, makeResourceOptions(options, id), false);
    }

    private static PluginArgs makeArgs(PluginArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? PluginArgs.Empty : args;
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<java.lang.String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .additionalSecretOutputs(List.of(
                "envs"
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
    public static Plugin get(java.lang.String name, Output<java.lang.String> id, @Nullable PluginState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Plugin(name, id, state, options);
    }
}
