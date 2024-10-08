// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vault.pkiSecret;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.vault.Utilities;
import com.pulumi.vault.pkiSecret.BackendConfigClusterArgs;
import com.pulumi.vault.pkiSecret.inputs.BackendConfigClusterState;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Allows setting the cluster-local&#39;s API mount path and AIA distribution point on a particular performance replication cluster.
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
 * import com.pulumi.vault.pkiSecret.BackendConfigCluster;
 * import com.pulumi.vault.pkiSecret.BackendConfigClusterArgs;
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
 *         var root = new Mount("root", MountArgs.builder()
 *             .path("pki-root")
 *             .type("pki")
 *             .description("root PKI")
 *             .defaultLeaseTtlSeconds(8640000)
 *             .maxLeaseTtlSeconds(8640000)
 *             .build());
 * 
 *         var example = new BackendConfigCluster("example", BackendConfigClusterArgs.builder()
 *             .backend(root.path())
 *             .path("http://127.0.0.1:8200/v1/pki-root")
 *             .aiaPath("http://127.0.0.1:8200/v1/pki-root")
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
 * The PKI config cluster can be imported using the resource&#39;s `id`.
 * In the case of the example above the `id` would be `pki-root/config/cluster`,
 * where the `pki-root` component is the resource&#39;s `backend`, e.g.
 * 
 * ```sh
 * $ pulumi import vault:pkiSecret/backendConfigCluster:BackendConfigCluster example pki-root/config/cluster
 * ```
 * 
 */
@ResourceType(type="vault:pkiSecret/backendConfigCluster:BackendConfigCluster")
public class BackendConfigCluster extends com.pulumi.resources.CustomResource {
    /**
     * Specifies the path to this performance replication cluster&#39;s AIA distribution point.
     * 
     */
    @Export(name="aiaPath", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> aiaPath;

    /**
     * @return Specifies the path to this performance replication cluster&#39;s AIA distribution point.
     * 
     */
    public Output<Optional<String>> aiaPath() {
        return Codegen.optional(this.aiaPath);
    }
    /**
     * The path the PKI secret backend is mounted at, with no leading or trailing `/`s.
     * 
     */
    @Export(name="backend", refs={String.class}, tree="[0]")
    private Output<String> backend;

    /**
     * @return The path the PKI secret backend is mounted at, with no leading or trailing `/`s.
     * 
     */
    public Output<String> backend() {
        return this.backend;
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
     * Specifies the path to this performance replication cluster&#39;s API mount path.
     * 
     */
    @Export(name="path", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> path;

    /**
     * @return Specifies the path to this performance replication cluster&#39;s API mount path.
     * 
     */
    public Output<Optional<String>> path() {
        return Codegen.optional(this.path);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public BackendConfigCluster(java.lang.String name) {
        this(name, BackendConfigClusterArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public BackendConfigCluster(java.lang.String name, BackendConfigClusterArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public BackendConfigCluster(java.lang.String name, BackendConfigClusterArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("vault:pkiSecret/backendConfigCluster:BackendConfigCluster", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private BackendConfigCluster(java.lang.String name, Output<java.lang.String> id, @Nullable BackendConfigClusterState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("vault:pkiSecret/backendConfigCluster:BackendConfigCluster", name, state, makeResourceOptions(options, id), false);
    }

    private static BackendConfigClusterArgs makeArgs(BackendConfigClusterArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? BackendConfigClusterArgs.Empty : args;
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
    public static BackendConfigCluster get(java.lang.String name, Output<java.lang.String> id, @Nullable BackendConfigClusterState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new BackendConfigCluster(name, id, state, options);
    }
}
