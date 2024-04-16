// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vault;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.vault.AuditArgs;
import com.pulumi.vault.Utilities;
import com.pulumi.vault.inputs.AuditState;
import java.lang.Boolean;
import java.lang.String;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * ## Example Usage
 * 
 * ### File Audit Device)
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.vault.Audit;
 * import com.pulumi.vault.AuditArgs;
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
 *         var test = new Audit(&#34;test&#34;, AuditArgs.builder()        
 *             .type(&#34;file&#34;)
 *             .options(Map.of(&#34;file_path&#34;, &#34;C:/temp/audit.txt&#34;))
 *             .build());
 * 
 *     }
 * }
 * ```
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ### Socket Audit Device)
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.vault.Audit;
 * import com.pulumi.vault.AuditArgs;
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
 *         var test = new Audit(&#34;test&#34;, AuditArgs.builder()        
 *             .type(&#34;socket&#34;)
 *             .path(&#34;app_socket&#34;)
 *             .local(false)
 *             .options(Map.ofEntries(
 *                 Map.entry(&#34;address&#34;, &#34;127.0.0.1:8000&#34;),
 *                 Map.entry(&#34;socket_type&#34;, &#34;tcp&#34;),
 *                 Map.entry(&#34;description&#34;, &#34;application x socket&#34;)
 *             ))
 *             .build());
 * 
 *     }
 * }
 * ```
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Import
 * 
 * Audit devices can be imported using the `path`, e.g.
 * 
 * ```sh
 * $ pulumi import vault:index/audit:Audit test syslog
 * ```
 * 
 */
@ResourceType(type="vault:index/audit:Audit")
public class Audit extends com.pulumi.resources.CustomResource {
    /**
     * Human-friendly description of the audit device.
     * 
     */
    @Export(name="description", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> description;

    /**
     * @return Human-friendly description of the audit device.
     * 
     */
    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    /**
     * Specifies if the audit device is a local only. Local audit devices are not replicated nor (if a secondary) removed by replication.
     * 
     */
    @Export(name="local", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> local;

    /**
     * @return Specifies if the audit device is a local only. Local audit devices are not replicated nor (if a secondary) removed by replication.
     * 
     */
    public Output<Optional<Boolean>> local() {
        return Codegen.optional(this.local);
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
     * Configuration options to pass to the audit device itself.
     * 
     * For a reference of the device types and their options, consult the [Vault documentation.](https://www.vaultproject.io/docs/audit/index.html)
     * 
     */
    @Export(name="options", refs={Map.class,String.class}, tree="[0,1,1]")
    private Output<Map<String,String>> options;

    /**
     * @return Configuration options to pass to the audit device itself.
     * 
     * For a reference of the device types and their options, consult the [Vault documentation.](https://www.vaultproject.io/docs/audit/index.html)
     * 
     */
    public Output<Map<String,String>> options() {
        return this.options;
    }
    /**
     * The path to mount the audit device. This defaults to the type.
     * 
     */
    @Export(name="path", refs={String.class}, tree="[0]")
    private Output<String> path;

    /**
     * @return The path to mount the audit device. This defaults to the type.
     * 
     */
    public Output<String> path() {
        return this.path;
    }
    /**
     * Type of the audit device, such as &#39;file&#39;.
     * 
     */
    @Export(name="type", refs={String.class}, tree="[0]")
    private Output<String> type;

    /**
     * @return Type of the audit device, such as &#39;file&#39;.
     * 
     */
    public Output<String> type() {
        return this.type;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Audit(String name) {
        this(name, AuditArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Audit(String name, AuditArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Audit(String name, AuditArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("vault:index/audit:Audit", name, args == null ? AuditArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private Audit(String name, Output<String> id, @Nullable AuditState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("vault:index/audit:Audit", name, state, makeResourceOptions(options, id));
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
    public static Audit get(String name, Output<String> id, @Nullable AuditState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Audit(name, id, state, options);
    }
}
