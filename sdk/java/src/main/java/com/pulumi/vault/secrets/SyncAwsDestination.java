// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vault.secrets;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.vault.Utilities;
import com.pulumi.vault.secrets.SyncAwsDestinationArgs;
import com.pulumi.vault.secrets.inputs.SyncAwsDestinationState;
import java.lang.Object;
import java.lang.String;
import java.util.List;
import java.util.Map;
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
 * import com.pulumi.vault.secrets.SyncAwsDestination;
 * import com.pulumi.vault.secrets.SyncAwsDestinationArgs;
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
 *         var aws = new SyncAwsDestination(&#34;aws&#34;, SyncAwsDestinationArgs.builder()        
 *             .accessKeyId(var_.access_key_id())
 *             .secretAccessKey(var_.secret_access_key())
 *             .region(&#34;us-east-1&#34;)
 *             .secretNameTemplate(&#34;vault_{{ .MountAccessor | lowercase }}_{{ .SecretPath | lowercase }}&#34;)
 *             .customTags(Map.of(&#34;foo&#34;, &#34;bar&#34;))
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * AWS Secrets sync destinations can be imported using the `name`, e.g.
 * 
 * ```sh
 *  $ pulumi import vault:secrets/syncAwsDestination:SyncAwsDestination aws aws-dest
 * ```
 * 
 */
@ResourceType(type="vault:secrets/syncAwsDestination:SyncAwsDestination")
public class SyncAwsDestination extends com.pulumi.resources.CustomResource {
    /**
     * Access key id to authenticate against the AWS secrets manager.
     * Can be omitted and directly provided to Vault using the `AWS_ACCESS_KEY_ID` environment
     * variable.
     * 
     */
    @Export(name="accessKeyId", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> accessKeyId;

    /**
     * @return Access key id to authenticate against the AWS secrets manager.
     * Can be omitted and directly provided to Vault using the `AWS_ACCESS_KEY_ID` environment
     * variable.
     * 
     */
    public Output<Optional<String>> accessKeyId() {
        return Codegen.optional(this.accessKeyId);
    }
    /**
     * Custom tags to set on the secret managed at the destination.
     * 
     */
    @Export(name="customTags", refs={Map.class,String.class,Object.class}, tree="[0,1,2]")
    private Output</* @Nullable */ Map<String,Object>> customTags;

    /**
     * @return Custom tags to set on the secret managed at the destination.
     * 
     */
    public Output<Optional<Map<String,Object>>> customTags() {
        return Codegen.optional(this.customTags);
    }
    /**
     * Unique name of the AWS destination.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return Unique name of the AWS destination.
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
     * Region where to manage the secrets manager entries.
     * Can be omitted and directly provided to Vault using the `AWS_REGION` environment
     * variable.
     * 
     */
    @Export(name="region", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> region;

    /**
     * @return Region where to manage the secrets manager entries.
     * Can be omitted and directly provided to Vault using the `AWS_REGION` environment
     * variable.
     * 
     */
    public Output<Optional<String>> region() {
        return Codegen.optional(this.region);
    }
    /**
     * Secret access key to authenticate against the AWS secrets manager.
     * Can be omitted and directly provided to Vault using the `AWS_SECRET_ACCESS_KEY` environment
     * variable.
     * 
     */
    @Export(name="secretAccessKey", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> secretAccessKey;

    /**
     * @return Secret access key to authenticate against the AWS secrets manager.
     * Can be omitted and directly provided to Vault using the `AWS_SECRET_ACCESS_KEY` environment
     * variable.
     * 
     */
    public Output<Optional<String>> secretAccessKey() {
        return Codegen.optional(this.secretAccessKey);
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
     * The type of the secrets destination (`aws-sm`).
     * 
     */
    @Export(name="type", refs={String.class}, tree="[0]")
    private Output<String> type;

    /**
     * @return The type of the secrets destination (`aws-sm`).
     * 
     */
    public Output<String> type() {
        return this.type;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public SyncAwsDestination(String name) {
        this(name, SyncAwsDestinationArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public SyncAwsDestination(String name, @Nullable SyncAwsDestinationArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public SyncAwsDestination(String name, @Nullable SyncAwsDestinationArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("vault:secrets/syncAwsDestination:SyncAwsDestination", name, args == null ? SyncAwsDestinationArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private SyncAwsDestination(String name, Output<String> id, @Nullable SyncAwsDestinationState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("vault:secrets/syncAwsDestination:SyncAwsDestination", name, state, makeResourceOptions(options, id));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .additionalSecretOutputs(List.of(
                "secretAccessKey"
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
    public static SyncAwsDestination get(String name, Output<String> id, @Nullable SyncAwsDestinationState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new SyncAwsDestination(name, id, state, options);
    }
}
