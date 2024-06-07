// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vault.secrets;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.vault.Utilities;
import com.pulumi.vault.secrets.SyncGithubAppsArgs;
import com.pulumi.vault.secrets.inputs.SyncGithubAppsState;
import java.lang.Integer;
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
 * import com.pulumi.vault.secrets.SyncGithubApps;
 * import com.pulumi.vault.secrets.SyncGithubAppsArgs;
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
 *         var github_apps = new SyncGithubApps("github-apps", SyncGithubAppsArgs.builder()
 *             .name("gh-apps")
 *             .appId(appId)
 *             .privateKey(StdFunctions.file(FileArgs.builder()
 *                 .input(privatekeyFile)
 *                 .build()).result())
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
 * GitHub Apps Secrets sync configuration endpoint can be imported using the `name`, e.g.
 * 
 * ```sh
 * $ pulumi import vault:secrets/syncGithubApps:SyncGithubApps gh github-apps
 * ```
 * 
 */
@ResourceType(type="vault:secrets/syncGithubApps:SyncGithubApps")
public class SyncGithubApps extends com.pulumi.resources.CustomResource {
    /**
     * The GitHub application ID.
     * 
     */
    @Export(name="appId", refs={Integer.class}, tree="[0]")
    private Output<Integer> appId;

    /**
     * @return The GitHub application ID.
     * 
     */
    public Output<Integer> appId() {
        return this.appId;
    }
    /**
     * A fingerprint of a private key.
     * 
     */
    @Export(name="fingerprint", refs={String.class}, tree="[0]")
    private Output<String> fingerprint;

    /**
     * @return A fingerprint of a private key.
     * 
     */
    public Output<String> fingerprint() {
        return this.fingerprint;
    }
    /**
     * The user-defined name of the GitHub App configuration.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return The user-defined name of the GitHub App configuration.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * The namespace to provision the resource in.
     * The value should not contain leading or trailing forward slashes.
     * The `namespace` is always relative to the provider&#39;s configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
     * 
     */
    @Export(name="namespace", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> namespace;

    /**
     * @return The namespace to provision the resource in.
     * The value should not contain leading or trailing forward slashes.
     * The `namespace` is always relative to the provider&#39;s configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
     * 
     */
    public Output<Optional<String>> namespace() {
        return Codegen.optional(this.namespace);
    }
    /**
     * The content of a PEM formatted private key generated on GitHub for the app.
     * 
     */
    @Export(name="privateKey", refs={String.class}, tree="[0]")
    private Output<String> privateKey;

    /**
     * @return The content of a PEM formatted private key generated on GitHub for the app.
     * 
     */
    public Output<String> privateKey() {
        return this.privateKey;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public SyncGithubApps(String name) {
        this(name, SyncGithubAppsArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public SyncGithubApps(String name, SyncGithubAppsArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public SyncGithubApps(String name, SyncGithubAppsArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("vault:secrets/syncGithubApps:SyncGithubApps", name, args == null ? SyncGithubAppsArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private SyncGithubApps(String name, Output<String> id, @Nullable SyncGithubAppsState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("vault:secrets/syncGithubApps:SyncGithubApps", name, state, makeResourceOptions(options, id));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .additionalSecretOutputs(List.of(
                "privateKey"
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
    public static SyncGithubApps get(String name, Output<String> id, @Nullable SyncGithubAppsState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new SyncGithubApps(name, id, state, options);
    }
}
