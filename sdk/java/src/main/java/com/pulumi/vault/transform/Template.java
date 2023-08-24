// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vault.transform;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.vault.Utilities;
import com.pulumi.vault.transform.TemplateArgs;
import com.pulumi.vault.transform.inputs.TemplateState;
import java.lang.Object;
import java.lang.String;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

@ResourceType(type="vault:transform/template:Template")
public class Template extends com.pulumi.resources.CustomResource {
    /**
     * The alphabet to use for this template. This is only used during FPE transformations.
     * 
     */
    @Export(name="alphabet", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> alphabet;

    /**
     * @return The alphabet to use for this template. This is only used during FPE transformations.
     * 
     */
    public Output<Optional<String>> alphabet() {
        return Codegen.optional(this.alphabet);
    }
    /**
     * - Optional mapping of name to regular expression template, used to customize
     *   the decoded output. (requires Vault Enterprise 1.9+)
     * 
     */
    @Export(name="decodeFormats", refs={Map.class,String.class,Object.class}, tree="[0,1,2]")
    private Output</* @Nullable */ Map<String,Object>> decodeFormats;

    /**
     * @return - Optional mapping of name to regular expression template, used to customize
     * the decoded output. (requires Vault Enterprise 1.9+)
     * 
     */
    public Output<Optional<Map<String,Object>>> decodeFormats() {
        return Codegen.optional(this.decodeFormats);
    }
    /**
     * - The regular expression template used to format encoded values.
     *   (requires Vault Enterprise 1.9+)
     * 
     */
    @Export(name="encodeFormat", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> encodeFormat;

    /**
     * @return - The regular expression template used to format encoded values.
     * (requires Vault Enterprise 1.9+)
     * 
     */
    public Output<Optional<String>> encodeFormat() {
        return Codegen.optional(this.encodeFormat);
    }
    /**
     * The name of the template.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return The name of the template.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * Path to where the back-end is mounted within Vault.
     * 
     */
    @Export(name="path", refs={String.class}, tree="[0]")
    private Output<String> path;

    /**
     * @return Path to where the back-end is mounted within Vault.
     * 
     */
    public Output<String> path() {
        return this.path;
    }
    /**
     * The pattern used for matching. Currently, only regular expression pattern is supported.
     * 
     */
    @Export(name="pattern", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> pattern;

    /**
     * @return The pattern used for matching. Currently, only regular expression pattern is supported.
     * 
     */
    public Output<Optional<String>> pattern() {
        return Codegen.optional(this.pattern);
    }
    /**
     * The pattern type to use for match detection. Currently, only regex is supported.
     * 
     */
    @Export(name="type", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> type;

    /**
     * @return The pattern type to use for match detection. Currently, only regex is supported.
     * 
     */
    public Output<Optional<String>> type() {
        return Codegen.optional(this.type);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Template(String name) {
        this(name, TemplateArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Template(String name, TemplateArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Template(String name, TemplateArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("vault:transform/template:Template", name, args == null ? TemplateArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private Template(String name, Output<String> id, @Nullable TemplateState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("vault:transform/template:Template", name, state, makeResourceOptions(options, id));
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
    public static Template get(String name, Output<String> id, @Nullable TemplateState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Template(name, id, state, options);
    }
}
