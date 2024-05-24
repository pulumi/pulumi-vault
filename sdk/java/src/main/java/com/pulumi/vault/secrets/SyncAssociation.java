// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vault.secrets;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.vault.Utilities;
import com.pulumi.vault.secrets.SyncAssociationArgs;
import com.pulumi.vault.secrets.inputs.SyncAssociationState;
import com.pulumi.vault.secrets.outputs.SyncAssociationMetadata;
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
 * import com.pulumi.vault.Mount;
 * import com.pulumi.vault.MountArgs;
 * import com.pulumi.vault.kv.SecretV2;
 * import com.pulumi.vault.kv.SecretV2Args;
 * import com.pulumi.vault.secrets.SyncGhDestination;
 * import com.pulumi.vault.secrets.SyncGhDestinationArgs;
 * import com.pulumi.vault.secrets.SyncAssociation;
 * import com.pulumi.vault.secrets.SyncAssociationArgs;
 * import static com.pulumi.codegen.internal.Serialization.*;
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
 *         var kvv2 = new Mount("kvv2", MountArgs.builder()
 *             .path("kvv2")
 *             .type("kv")
 *             .options(Map.of("version", "2"))
 *             .description("KV Version 2 secret engine mount")
 *             .build());
 * 
 *         var token = new SecretV2("token", SecretV2Args.builder()
 *             .mount(kvv2.path())
 *             .name("token")
 *             .dataJson(serializeJson(
 *                 jsonObject(
 *                     jsonProperty("dev", "B!gS3cr3t"),
 *                     jsonProperty("prod", "S3cureP4$$")
 *                 )))
 *             .build());
 * 
 *         var gh = new SyncGhDestination("gh", SyncGhDestinationArgs.builder()
 *             .name("gh-dest")
 *             .accessToken(accessToken)
 *             .repositoryOwner(repoOwner)
 *             .repositoryName("repo-name-example")
 *             .secretNameTemplate("vault_{{ .MountAccessor | lowercase }}_{{ .SecretPath | lowercase }}")
 *             .build());
 * 
 *         var ghToken = new SyncAssociation("ghToken", SyncAssociationArgs.builder()
 *             .name(gh.name())
 *             .type(gh.type())
 *             .mount(kvv2.path())
 *             .secretName(token.name())
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 */
@ResourceType(type="vault:secrets/syncAssociation:SyncAssociation")
public class SyncAssociation extends com.pulumi.resources.CustomResource {
    /**
     * Metadata for each subkey of the associated secret.
     * 
     */
    @Export(name="metadatas", refs={List.class,SyncAssociationMetadata.class}, tree="[0,1]")
    private Output<List<SyncAssociationMetadata>> metadatas;

    /**
     * @return Metadata for each subkey of the associated secret.
     * 
     */
    public Output<List<SyncAssociationMetadata>> metadatas() {
        return this.metadatas;
    }
    /**
     * Specifies the mount where the secret is located.
     * 
     */
    @Export(name="mount", refs={String.class}, tree="[0]")
    private Output<String> mount;

    /**
     * @return Specifies the mount where the secret is located.
     * 
     */
    public Output<String> mount() {
        return this.mount;
    }
    /**
     * Specifies the name of the destination.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return Specifies the name of the destination.
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
     * Specifies the name of the secret to synchronize.
     * 
     */
    @Export(name="secretName", refs={String.class}, tree="[0]")
    private Output<String> secretName;

    /**
     * @return Specifies the name of the secret to synchronize.
     * 
     */
    public Output<String> secretName() {
        return this.secretName;
    }
    /**
     * Specifies the destination type.
     * 
     */
    @Export(name="type", refs={String.class}, tree="[0]")
    private Output<String> type;

    /**
     * @return Specifies the destination type.
     * 
     */
    public Output<String> type() {
        return this.type;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public SyncAssociation(String name) {
        this(name, SyncAssociationArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public SyncAssociation(String name, SyncAssociationArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public SyncAssociation(String name, SyncAssociationArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("vault:secrets/syncAssociation:SyncAssociation", name, args == null ? SyncAssociationArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private SyncAssociation(String name, Output<String> id, @Nullable SyncAssociationState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("vault:secrets/syncAssociation:SyncAssociation", name, state, makeResourceOptions(options, id));
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
    public static SyncAssociation get(String name, Output<String> id, @Nullable SyncAssociationState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new SyncAssociation(name, id, state, options);
    }
}
