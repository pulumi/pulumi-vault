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
import java.lang.String;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * This resource supports the `/transform/template/{name}` Vault endpoint.
 * 
 * It creates or updates a template with the given name. If a template with the name does not exist,
 * it will be created. If the template exists, it will be updated with the new attributes.
 * 
 * &gt; Requires *Vault Enterprise with the Advanced Data Protection Transform Module*.
 * See [Transform Secrets Engine](https://www.vaultproject.io/docs/secrets/transform)
 * for more information.
 * 
 * ## Example Usage
 * 
 * Please note that the `pattern` below holds a regex. The regex shown
 * is identical to the one in our [Setup](https://www.vaultproject.io/docs/secrets/transform#setup)
 * docs, `(\d{4})-(\d{4})-(\d{4})-(\d{4})`. However, due to HCL, the
 * backslashes must be escaped to appear correctly in Vault. For further
 * assistance escaping your own custom regex, see String Literals.
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
 * import com.pulumi.vault.transform.Alphabet;
 * import com.pulumi.vault.transform.AlphabetArgs;
 * import com.pulumi.vault.transform.Template;
 * import com.pulumi.vault.transform.TemplateArgs;
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
 *         var transform = new Mount("transform", MountArgs.builder()
 *             .path("transform")
 *             .type("transform")
 *             .build());
 * 
 *         var numerics = new Alphabet("numerics", AlphabetArgs.builder()
 *             .path(transform.path())
 *             .name("numerics")
 *             .alphabet("0123456789")
 *             .build());
 * 
 *         var test = new Template("test", TemplateArgs.builder()
 *             .path(numerics.path())
 *             .name("ccn")
 *             .type("regex")
 *             .pattern("(\\d{4})[- ](\\d{4})[- ](\\d{4})[- ](\\d{4})")
 *             .alphabet("numerics")
 *             .encodeFormat("$1-$2-$3-$4")
 *             .decodeFormats(Map.of("last-four-digits", "$4"))
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 */
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
     * Optional mapping of name to regular expression template, used to customize
     * the decoded output. (requires Vault Enterprise 1.9+)
     * 
     */
    @Export(name="decodeFormats", refs={Map.class,String.class}, tree="[0,1,1]")
    private Output</* @Nullable */ Map<String,String>> decodeFormats;

    /**
     * @return Optional mapping of name to regular expression template, used to customize
     * the decoded output. (requires Vault Enterprise 1.9+)
     * 
     */
    public Output<Optional<Map<String,String>>> decodeFormats() {
        return Codegen.optional(this.decodeFormats);
    }
    /**
     * The regular expression template used to format encoded values.
     * (requires Vault Enterprise 1.9+)
     * 
     */
    @Export(name="encodeFormat", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> encodeFormat;

    /**
     * @return The regular expression template used to format encoded values.
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
    public Template(java.lang.String name) {
        this(name, TemplateArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Template(java.lang.String name, TemplateArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Template(java.lang.String name, TemplateArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("vault:transform/template:Template", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private Template(java.lang.String name, Output<java.lang.String> id, @Nullable TemplateState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("vault:transform/template:Template", name, state, makeResourceOptions(options, id), false);
    }

    private static TemplateArgs makeArgs(TemplateArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? TemplateArgs.Empty : args;
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
    public static Template get(java.lang.String name, Output<java.lang.String> id, @Nullable TemplateState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Template(name, id, state, options);
    }
}
