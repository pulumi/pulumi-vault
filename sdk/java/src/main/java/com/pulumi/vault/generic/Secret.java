// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vault.generic;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.vault.Utilities;
import com.pulumi.vault.generic.SecretArgs;
import com.pulumi.vault.generic.inputs.SecretState;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * ## Import
 * 
 * Generic secrets can be imported using the `path`, e.g.
 * 
 * ```sh
 * $ pulumi import vault:generic/secret:Secret example secret/foo
 * ```
 * 
 */
@ResourceType(type="vault:generic/secret:Secret")
public class Secret extends com.pulumi.resources.CustomResource {
    /**
     * A mapping whose keys are the top-level data keys returned from
     * Vault and whose values are the corresponding values. This map can only
     * represent string data, so any non-string values returned from Vault are
     * serialized as JSON.
     * 
     */
    @Export(name="data", refs={Map.class,String.class}, tree="[0,1,1]")
    private Output<Map<String,String>> data;

    /**
     * @return A mapping whose keys are the top-level data keys returned from
     * Vault and whose values are the corresponding values. This map can only
     * represent string data, so any non-string values returned from Vault are
     * serialized as JSON.
     * 
     */
    public Output<Map<String,String>> data() {
        return this.data;
    }
    /**
     * String containing a JSON-encoded object that will be
     * written as the secret data at the given path.
     * 
     */
    @Export(name="dataJson", refs={String.class}, tree="[0]")
    private Output<String> dataJson;

    /**
     * @return String containing a JSON-encoded object that will be
     * written as the secret data at the given path.
     * 
     */
    public Output<String> dataJson() {
        return this.dataJson;
    }
    /**
     * true/false.  Only applicable for kv-v2 stores.
     * If set to `true`, permanently deletes all versions for
     * the specified key. The default behavior is to only delete the latest version of the
     * secret.
     * 
     */
    @Export(name="deleteAllVersions", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> deleteAllVersions;

    /**
     * @return true/false.  Only applicable for kv-v2 stores.
     * If set to `true`, permanently deletes all versions for
     * the specified key. The default behavior is to only delete the latest version of the
     * secret.
     * 
     */
    public Output<Optional<Boolean>> deleteAllVersions() {
        return Codegen.optional(this.deleteAllVersions);
    }
    /**
     * true/false. Set this to true if your vault
     * authentication is not able to read the data. Setting this to `true` will
     * break drift detection. Defaults to false.
     * 
     */
    @Export(name="disableRead", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> disableRead;

    /**
     * @return true/false. Set this to true if your vault
     * authentication is not able to read the data. Setting this to `true` will
     * break drift detection. Defaults to false.
     * 
     */
    public Output<Optional<Boolean>> disableRead() {
        return Codegen.optional(this.disableRead);
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
     * The full logical path at which to write the given data.
     * To write data into the &#34;generic&#34; secret backend mounted in Vault by default,
     * this should be prefixed with `secret/`. Writing to other backends with this
     * resource is possible; consult each backend&#39;s documentation to see which
     * endpoints support the `PUT` and `DELETE` methods.
     * 
     */
    @Export(name="path", refs={String.class}, tree="[0]")
    private Output<String> path;

    /**
     * @return The full logical path at which to write the given data.
     * To write data into the &#34;generic&#34; secret backend mounted in Vault by default,
     * this should be prefixed with `secret/`. Writing to other backends with this
     * resource is possible; consult each backend&#39;s documentation to see which
     * endpoints support the `PUT` and `DELETE` methods.
     * 
     */
    public Output<String> path() {
        return this.path;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Secret(java.lang.String name) {
        this(name, SecretArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Secret(java.lang.String name, SecretArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Secret(java.lang.String name, SecretArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("vault:generic/secret:Secret", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private Secret(java.lang.String name, Output<java.lang.String> id, @Nullable SecretState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("vault:generic/secret:Secret", name, state, makeResourceOptions(options, id), false);
    }

    private static SecretArgs makeArgs(SecretArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? SecretArgs.Empty : args;
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<java.lang.String> id) {
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
    public static Secret get(java.lang.String name, Output<java.lang.String> id, @Nullable SecretState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Secret(name, id, state, options);
    }
}
