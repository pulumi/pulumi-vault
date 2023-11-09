// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vault.transform;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.vault.Utilities;
import com.pulumi.vault.transform.AlphabetArgs;
import com.pulumi.vault.transform.inputs.AlphabetState;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * This resource supports the &#34;/transform/alphabet/{name}&#34; Vault endpoint.
 * 
 * It queries an existing alphabet by the given name.
 * 
 * ## Example Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.vault.Mount;
 * import com.pulumi.vault.MountArgs;
 * import com.pulumi.vault.transform.Alphabet;
 * import com.pulumi.vault.transform.AlphabetArgs;
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
 *         var mountTransform = new Mount(&#34;mountTransform&#34;, MountArgs.builder()        
 *             .path(&#34;transform&#34;)
 *             .type(&#34;transform&#34;)
 *             .build());
 * 
 *         var test = new Alphabet(&#34;test&#34;, AlphabetArgs.builder()        
 *             .path(mountTransform.path())
 *             .alphabet(&#34;0123456789&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 */
@ResourceType(type="vault:transform/alphabet:Alphabet")
public class Alphabet extends com.pulumi.resources.CustomResource {
    /**
     * A string of characters that contains the alphabet set.
     * 
     */
    @Export(name="alphabet", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> alphabet;

    /**
     * @return A string of characters that contains the alphabet set.
     * 
     */
    public Output<Optional<String>> alphabet() {
        return Codegen.optional(this.alphabet);
    }
    /**
     * The name of the alphabet.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return The name of the alphabet.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * The namespace to provision the resource in.
     * The value should not contain leading or trailing forward slashes.
     * The `namespace` is always relative to the provider&#39;s configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
     * *Available only for Vault Enterprise*.
     * 
     */
    @Export(name="namespace", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> namespace;

    /**
     * @return The namespace to provision the resource in.
     * The value should not contain leading or trailing forward slashes.
     * The `namespace` is always relative to the provider&#39;s configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
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
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Alphabet(String name) {
        this(name, AlphabetArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Alphabet(String name, AlphabetArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Alphabet(String name, AlphabetArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("vault:transform/alphabet:Alphabet", name, args == null ? AlphabetArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private Alphabet(String name, Output<String> id, @Nullable AlphabetState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("vault:transform/alphabet:Alphabet", name, state, makeResourceOptions(options, id));
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
    public static Alphabet get(String name, Output<String> id, @Nullable AlphabetState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Alphabet(name, id, state, options);
    }
}
