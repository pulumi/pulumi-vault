// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vault.kv;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.vault.Utilities;
import com.pulumi.vault.kv.SecretV2Args;
import com.pulumi.vault.kv.inputs.SecretV2State;
import com.pulumi.vault.kv.outputs.SecretV2CustomMetadata;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.Object;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Writes a KV-V2 secret to a given path in Vault.
 * 
 * For more information on Vault&#39;s KV-V2 secret backend
 * [see here](https://www.vaultproject.io/docs/secrets/kv/kv-v2).
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
 * import com.pulumi.vault.kv.SecretV2;
 * import com.pulumi.vault.kv.SecretV2Args;
 * import com.pulumi.vault.kv.inputs.SecretV2CustomMetadataArgs;
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
 *         var example = new SecretV2("example", SecretV2Args.builder()
 *             .mount(kvv2.path())
 *             .name("secret")
 *             .cas(1)
 *             .deleteAllVersions(true)
 *             .dataJson(serializeJson(
 *                 jsonObject(
 *                     jsonProperty("zip", "zap"),
 *                     jsonProperty("foo", "bar")
 *                 )))
 *             .customMetadata(SecretV2CustomMetadataArgs.builder()
 *                 .maxVersions(5)
 *                 .data(Map.ofEntries(
 *                     Map.entry("foo", "vault{@literal @}example.com"),
 *                     Map.entry("bar", "12345")
 *                 ))
 *                 .build())
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Required Vault Capabilities
 * 
 * Use of this resource requires the `create` or `update` capability
 * (depending on whether the resource already exists) on the given path,
 * the `delete` capability if the resource is removed from configuration,
 * and the `read` capability for drift detection (by default).
 * 
 * ### Custom Metadata Configuration Options
 * 
 * * `max_versions` - (Optional) The number of versions to keep per key.
 * 
 * * `cas_required` - (Optional) If true, all keys will require the cas
 * parameter to be set on all write requests.
 * 
 * * `delete_version_after` - (Optional) If set, specifies the length of time before
 * a version is deleted. Accepts duration in integer seconds.
 * 
 * * `data` - (Optional) A string to string map describing the secret.
 * 
 * ## Import
 * 
 * KV-V2 secrets can be imported using the `path`, e.g.
 * 
 * ```sh
 * $ pulumi import vault:kv/secretV2:SecretV2 example kvv2/data/secret
 * ```
 * 
 */
@ResourceType(type="vault:kv/secretV2:SecretV2")
public class SecretV2 extends com.pulumi.resources.CustomResource {
    /**
     * This flag is required if `cas_required` is set to true
     * on either the secret or the engine&#39;s config. In order for a
     * write operation to be successful, cas must be set to the current version
     * of the secret.
     * 
     */
    @Export(name="cas", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> cas;

    /**
     * @return This flag is required if `cas_required` is set to true
     * on either the secret or the engine&#39;s config. In order for a
     * write operation to be successful, cas must be set to the current version
     * of the secret.
     * 
     */
    public Output<Optional<Integer>> cas() {
        return Codegen.optional(this.cas);
    }
    /**
     * A nested block that allows configuring metadata for the
     * KV secret. Refer to the
     * Configuration Options for more info.
     * 
     */
    @Export(name="customMetadata", refs={SecretV2CustomMetadata.class}, tree="[0]")
    private Output<SecretV2CustomMetadata> customMetadata;

    /**
     * @return A nested block that allows configuring metadata for the
     * KV secret. Refer to the
     * Configuration Options for more info.
     * 
     */
    public Output<SecretV2CustomMetadata> customMetadata() {
        return this.customMetadata;
    }
    /**
     * A mapping whose keys are the top-level data keys returned from
     * Vault and whose values are the corresponding values. This map can only
     * represent string data, so any non-string values returned from Vault are
     * serialized as JSON.
     * 
     */
    @Export(name="data", refs={Map.class,String.class,Object.class}, tree="[0,1,2]")
    private Output<Map<String,Object>> data;

    /**
     * @return A mapping whose keys are the top-level data keys returned from
     * Vault and whose values are the corresponding values. This map can only
     * represent string data, so any non-string values returned from Vault are
     * serialized as JSON.
     * 
     */
    public Output<Map<String,Object>> data() {
        return this.data;
    }
    /**
     * JSON-encoded string that will be
     * written as the secret data at the given path.
     * 
     */
    @Export(name="dataJson", refs={String.class}, tree="[0]")
    private Output<String> dataJson;

    /**
     * @return JSON-encoded string that will be
     * written as the secret data at the given path.
     * 
     */
    public Output<String> dataJson() {
        return this.dataJson;
    }
    /**
     * If set to true, permanently deletes all
     * versions for the specified key.
     * 
     */
    @Export(name="deleteAllVersions", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> deleteAllVersions;

    /**
     * @return If set to true, permanently deletes all
     * versions for the specified key.
     * 
     */
    public Output<Optional<Boolean>> deleteAllVersions() {
        return Codegen.optional(this.deleteAllVersions);
    }
    /**
     * If set to true, disables reading secret from Vault;
     * note: drift won&#39;t be detected.
     * 
     */
    @Export(name="disableRead", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> disableRead;

    /**
     * @return If set to true, disables reading secret from Vault;
     * note: drift won&#39;t be detected.
     * 
     */
    public Output<Optional<Boolean>> disableRead() {
        return Codegen.optional(this.disableRead);
    }
    /**
     * Metadata associated with this secret read from Vault.
     * 
     */
    @Export(name="metadata", refs={Map.class,String.class,Object.class}, tree="[0,1,2]")
    private Output<Map<String,Object>> metadata;

    /**
     * @return Metadata associated with this secret read from Vault.
     * 
     */
    public Output<Map<String,Object>> metadata() {
        return this.metadata;
    }
    /**
     * Path where KV-V2 engine is mounted.
     * 
     */
    @Export(name="mount", refs={String.class}, tree="[0]")
    private Output<String> mount;

    /**
     * @return Path where KV-V2 engine is mounted.
     * 
     */
    public Output<String> mount() {
        return this.mount;
    }
    /**
     * Full name of the secret. For a nested secret
     * the name is the nested path excluding the mount and data
     * prefix. For example, for a secret at `kvv2/data/foo/bar/baz`
     * the name is `foo/bar/baz`.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return Full name of the secret. For a nested secret
     * the name is the nested path excluding the mount and data
     * prefix. For example, for a secret at `kvv2/data/foo/bar/baz`
     * the name is `foo/bar/baz`.
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
     * An object that holds option settings.
     * 
     */
    @Export(name="options", refs={Map.class,String.class,Object.class}, tree="[0,1,2]")
    private Output</* @Nullable */ Map<String,Object>> options;

    /**
     * @return An object that holds option settings.
     * 
     */
    public Output<Optional<Map<String,Object>>> options() {
        return Codegen.optional(this.options);
    }
    /**
     * Full path where the KV-V2 secret will be written.
     * 
     */
    @Export(name="path", refs={String.class}, tree="[0]")
    private Output<String> path;

    /**
     * @return Full path where the KV-V2 secret will be written.
     * 
     */
    public Output<String> path() {
        return this.path;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public SecretV2(String name) {
        this(name, SecretV2Args.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public SecretV2(String name, SecretV2Args args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public SecretV2(String name, SecretV2Args args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("vault:kv/secretV2:SecretV2", name, args == null ? SecretV2Args.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private SecretV2(String name, Output<String> id, @Nullable SecretV2State state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("vault:kv/secretV2:SecretV2", name, state, makeResourceOptions(options, id));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .additionalSecretOutputs(List.of(
                "data",
                "dataJson"
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
    public static SecretV2 get(String name, Output<String> id, @Nullable SecretV2State state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new SecretV2(name, id, state, options);
    }
}
