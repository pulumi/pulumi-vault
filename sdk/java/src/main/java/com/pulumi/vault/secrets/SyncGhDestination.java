// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vault.secrets;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.vault.Utilities;
import com.pulumi.vault.secrets.SyncGhDestinationArgs;
import com.pulumi.vault.secrets.inputs.SyncGhDestinationState;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * ## Example Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.vault.secrets.SyncGhDestination;
 * import com.pulumi.vault.secrets.SyncGhDestinationArgs;
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
 *         var gh = new SyncGhDestination(&#34;gh&#34;, SyncGhDestinationArgs.builder()        
 *             .accessToken(var_.access_token())
 *             .repositoryOwner(var_.repo_owner())
 *             .repositoryName(&#34;repo-name-example&#34;)
 *             .secretNameTemplate(&#34;vault_{{ .MountAccessor | lowercase }}_{{ .SecretPath | lowercase }}&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * GitHub Secrets sync destinations can be imported using the `name`, e.g.
 * 
 * ```sh
 *  $ pulumi import vault:secrets/syncGhDestination:SyncGhDestination gh gh-dest
 * ```
 * 
 */
@ResourceType(type="vault:secrets/syncGhDestination:SyncGhDestination")
public class SyncGhDestination extends com.pulumi.resources.CustomResource {
    /**
     * Fine-grained or personal access token.
     * Can be omitted and directly provided to Vault using the `GITHUB_ACCESS_TOKEN` environment
     * variable.
     * 
     */
    @Export(name="accessToken", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> accessToken;

    /**
     * @return Fine-grained or personal access token.
     * Can be omitted and directly provided to Vault using the `GITHUB_ACCESS_TOKEN` environment
     * variable.
     * 
     */
    public Output<Optional<String>> accessToken() {
        return Codegen.optional(this.accessToken);
    }
    /**
     * Unique name of the GitHub destination.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return Unique name of the GitHub destination.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * The namespace to provision the resource in.
     * The value should not contain leading or trailing forward slashes.
     * The `namespace` is always relative to the provider&#39;s configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
     * 
     */
    @Export(name="namespace", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> namespace;

    /**
     * @return The namespace to provision the resource in.
     * The value should not contain leading or trailing forward slashes.
     * The `namespace` is always relative to the provider&#39;s configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
     * 
     */
    public Output<Optional<String>> namespace() {
        return Codegen.optional(this.namespace);
    }
    /**
     * Name of the repository.
     * Can be omitted and directly provided to Vault using the `GITHUB_REPOSITORY_NAME` environment
     * variable.
     * 
     */
    @Export(name="repositoryName", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> repositoryName;

    /**
     * @return Name of the repository.
     * Can be omitted and directly provided to Vault using the `GITHUB_REPOSITORY_NAME` environment
     * variable.
     * 
     */
    public Output<Optional<String>> repositoryName() {
        return Codegen.optional(this.repositoryName);
    }
    /**
     * GitHub organization or username that owns the repository.
     * Can be omitted and directly provided to Vault using the `GITHUB_REPOSITORY_OWNER` environment
     * variable.
     * 
     */
    @Export(name="repositoryOwner", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> repositoryOwner;

    /**
     * @return GitHub organization or username that owns the repository.
     * Can be omitted and directly provided to Vault using the `GITHUB_REPOSITORY_OWNER` environment
     * variable.
     * 
     */
    public Output<Optional<String>> repositoryOwner() {
        return Codegen.optional(this.repositoryOwner);
    }
    /**
     * Template describing how to generate external secret names.
     * Supports a subset of the Go Template syntax.
     * 
     */
    @Export(name="secretNameTemplate", refs={String.class}, tree="[0]")
    private Output<String> secretNameTemplate;

    /**
     * @return Template describing how to generate external secret names.
     * Supports a subset of the Go Template syntax.
     * 
     */
    public Output<String> secretNameTemplate() {
        return this.secretNameTemplate;
    }
    /**
     * The type of the secrets destination (`gh`).
     * 
     */
    @Export(name="type", refs={String.class}, tree="[0]")
    private Output<String> type;

    /**
     * @return The type of the secrets destination (`gh`).
     * 
     */
    public Output<String> type() {
        return this.type;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public SyncGhDestination(String name) {
        this(name, SyncGhDestinationArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public SyncGhDestination(String name, @Nullable SyncGhDestinationArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public SyncGhDestination(String name, @Nullable SyncGhDestinationArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("vault:secrets/syncGhDestination:SyncGhDestination", name, args == null ? SyncGhDestinationArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private SyncGhDestination(String name, Output<String> id, @Nullable SyncGhDestinationState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("vault:secrets/syncGhDestination:SyncGhDestination", name, state, makeResourceOptions(options, id));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .additionalSecretOutputs(List.of(
                "accessToken"
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
    public static SyncGhDestination get(String name, Output<String> id, @Nullable SyncGhDestinationState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new SyncGhDestination(name, id, state, options);
    }
}