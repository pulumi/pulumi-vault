// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vault;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.vault.RgpPolicyArgs;
import com.pulumi.vault.Utilities;
import com.pulumi.vault.inputs.RgpPolicyState;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides a resource to manage Role Governing Policy (RGP) via [Sentinel](https://www.vaultproject.io/docs/enterprise/sentinel/index.html).
 * 
 * **Note** this feature is available only with Vault Enterprise.
 * 
 * ## Example Usage
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.vault.RgpPolicy;
 * import com.pulumi.vault.RgpPolicyArgs;
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
 *         var allow_all = new RgpPolicy(&#34;allow-all&#34;, RgpPolicyArgs.builder()        
 *             .enforcementLevel(&#34;soft-mandatory&#34;)
 *             .policy(&#34;&#34;&#34;
 * main = rule {
 *   true
 * }
 * 
 *             &#34;&#34;&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 */
@ResourceType(type="vault:index/rgpPolicy:RgpPolicy")
public class RgpPolicy extends com.pulumi.resources.CustomResource {
    /**
     * Enforcement level of Sentinel policy. Can be either `advisory` or `soft-mandatory` or `hard-mandatory`
     * 
     */
    @Export(name="enforcementLevel", refs={String.class}, tree="[0]")
    private Output<String> enforcementLevel;

    /**
     * @return Enforcement level of Sentinel policy. Can be either `advisory` or `soft-mandatory` or `hard-mandatory`
     * 
     */
    public Output<String> enforcementLevel() {
        return this.enforcementLevel;
    }
    /**
     * The name of the policy
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return The name of the policy
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
     * String containing a Sentinel policy
     * 
     */
    @Export(name="policy", refs={String.class}, tree="[0]")
    private Output<String> policy;

    /**
     * @return String containing a Sentinel policy
     * 
     */
    public Output<String> policy() {
        return this.policy;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public RgpPolicy(String name) {
        this(name, RgpPolicyArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public RgpPolicy(String name, RgpPolicyArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public RgpPolicy(String name, RgpPolicyArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("vault:index/rgpPolicy:RgpPolicy", name, args == null ? RgpPolicyArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private RgpPolicy(String name, Output<String> id, @Nullable RgpPolicyState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("vault:index/rgpPolicy:RgpPolicy", name, state, makeResourceOptions(options, id));
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
    public static RgpPolicy get(String name, Output<String> id, @Nullable RgpPolicyState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new RgpPolicy(name, id, state, options);
    }
}
