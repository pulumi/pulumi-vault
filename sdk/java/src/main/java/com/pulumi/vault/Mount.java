// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vault;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.vault.MountArgs;
import com.pulumi.vault.Utilities;
import com.pulumi.vault.inputs.MountState;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.Object;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * This resource enables a new secrets engine at the given path.
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
 * import com.pulumi.vault.Mount;
 * import com.pulumi.vault.MountArgs;
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
 *         var example = new Mount(&#34;example&#34;, MountArgs.builder()        
 *             .description(&#34;This is an example mount&#34;)
 *             .path(&#34;dummy&#34;)
 *             .type(&#34;generic&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.vault.Mount;
 * import com.pulumi.vault.MountArgs;
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
 *         var kvv2_example = new Mount(&#34;kvv2-example&#34;, MountArgs.builder()        
 *             .description(&#34;This is an example KV Version 2 secret engine mount&#34;)
 *             .options(Map.ofEntries(
 *                 Map.entry(&#34;type&#34;, &#34;kv-v2&#34;),
 *                 Map.entry(&#34;version&#34;, &#34;2&#34;)
 *             ))
 *             .path(&#34;version2-example&#34;)
 *             .type(&#34;kv-v2&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.vault.Mount;
 * import com.pulumi.vault.MountArgs;
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
 *         var transit_example = new Mount(&#34;transit-example&#34;, MountArgs.builder()        
 *             .description(&#34;This is an example transit secret engine mount&#34;)
 *             .options(Map.of(&#34;convergent_encryption&#34;, false))
 *             .path(&#34;transit-example&#34;)
 *             .type(&#34;transit&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.vault.Mount;
 * import com.pulumi.vault.MountArgs;
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
 *         var pki_example = new Mount(&#34;pki-example&#34;, MountArgs.builder()        
 *             .defaultLeaseTtlSeconds(3600)
 *             .description(&#34;This is an example PKI mount&#34;)
 *             .maxLeaseTtlSeconds(86400)
 *             .path(&#34;pki-example&#34;)
 *             .type(&#34;pki&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Import
 * 
 * Mounts can be imported using the `path`, e.g.
 * 
 * ```sh
 * $ pulumi import vault:index/mount:Mount example dummy
 * ```
 * 
 */
@ResourceType(type="vault:index/mount:Mount")
public class Mount extends com.pulumi.resources.CustomResource {
    /**
     * The accessor for this mount.
     * 
     */
    @Export(name="accessor", refs={String.class}, tree="[0]")
    private Output<String> accessor;

    /**
     * @return The accessor for this mount.
     * 
     */
    public Output<String> accessor() {
        return this.accessor;
    }
    /**
     * Set of managed key registry entry names that the mount in question is allowed to access
     * 
     */
    @Export(name="allowedManagedKeys", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> allowedManagedKeys;

    /**
     * @return Set of managed key registry entry names that the mount in question is allowed to access
     * 
     */
    public Output<Optional<List<String>>> allowedManagedKeys() {
        return Codegen.optional(this.allowedManagedKeys);
    }
    /**
     * Specifies the list of keys that will not be HMAC&#39;d by audit devices in the request data object.
     * 
     */
    @Export(name="auditNonHmacRequestKeys", refs={List.class,String.class}, tree="[0,1]")
    private Output<List<String>> auditNonHmacRequestKeys;

    /**
     * @return Specifies the list of keys that will not be HMAC&#39;d by audit devices in the request data object.
     * 
     */
    public Output<List<String>> auditNonHmacRequestKeys() {
        return this.auditNonHmacRequestKeys;
    }
    /**
     * Specifies the list of keys that will not be HMAC&#39;d by audit devices in the response data object.
     * 
     */
    @Export(name="auditNonHmacResponseKeys", refs={List.class,String.class}, tree="[0,1]")
    private Output<List<String>> auditNonHmacResponseKeys;

    /**
     * @return Specifies the list of keys that will not be HMAC&#39;d by audit devices in the response data object.
     * 
     */
    public Output<List<String>> auditNonHmacResponseKeys() {
        return this.auditNonHmacResponseKeys;
    }
    /**
     * Default lease duration for tokens and secrets in seconds
     * 
     */
    @Export(name="defaultLeaseTtlSeconds", refs={Integer.class}, tree="[0]")
    private Output<Integer> defaultLeaseTtlSeconds;

    /**
     * @return Default lease duration for tokens and secrets in seconds
     * 
     */
    public Output<Integer> defaultLeaseTtlSeconds() {
        return this.defaultLeaseTtlSeconds;
    }
    /**
     * Human-friendly description of the mount
     * 
     */
    @Export(name="description", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> description;

    /**
     * @return Human-friendly description of the mount
     * 
     */
    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    /**
     * Boolean flag that can be explicitly set to true to enable the secrets engine to access Vault&#39;s external entropy source
     * 
     */
    @Export(name="externalEntropyAccess", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> externalEntropyAccess;

    /**
     * @return Boolean flag that can be explicitly set to true to enable the secrets engine to access Vault&#39;s external entropy source
     * 
     */
    public Output<Optional<Boolean>> externalEntropyAccess() {
        return Codegen.optional(this.externalEntropyAccess);
    }
    /**
     * Boolean flag that can be explicitly set to true to enforce local mount in HA environment
     * 
     */
    @Export(name="local", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> local;

    /**
     * @return Boolean flag that can be explicitly set to true to enforce local mount in HA environment
     * 
     */
    public Output<Optional<Boolean>> local() {
        return Codegen.optional(this.local);
    }
    /**
     * Maximum possible lease duration for tokens and secrets in seconds
     * 
     */
    @Export(name="maxLeaseTtlSeconds", refs={Integer.class}, tree="[0]")
    private Output<Integer> maxLeaseTtlSeconds;

    /**
     * @return Maximum possible lease duration for tokens and secrets in seconds
     * 
     */
    public Output<Integer> maxLeaseTtlSeconds() {
        return this.maxLeaseTtlSeconds;
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
     * Specifies mount type specific options that are passed to the backend
     * 
     */
    @Export(name="options", refs={Map.class,String.class,Object.class}, tree="[0,1,2]")
    private Output</* @Nullable */ Map<String,Object>> options;

    /**
     * @return Specifies mount type specific options that are passed to the backend
     * 
     */
    public Output<Optional<Map<String,Object>>> options() {
        return Codegen.optional(this.options);
    }
    /**
     * Where the secret backend will be mounted
     * 
     */
    @Export(name="path", refs={String.class}, tree="[0]")
    private Output<String> path;

    /**
     * @return Where the secret backend will be mounted
     * 
     */
    public Output<String> path() {
        return this.path;
    }
    /**
     * Boolean flag that can be explicitly set to true to enable seal wrapping for the mount, causing values stored by the mount to be wrapped by the seal&#39;s encryption capability
     * 
     */
    @Export(name="sealWrap", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> sealWrap;

    /**
     * @return Boolean flag that can be explicitly set to true to enable seal wrapping for the mount, causing values stored by the mount to be wrapped by the seal&#39;s encryption capability
     * 
     */
    public Output<Boolean> sealWrap() {
        return this.sealWrap;
    }
    /**
     * Type of the backend, such as &#34;aws&#34;
     * 
     */
    @Export(name="type", refs={String.class}, tree="[0]")
    private Output<String> type;

    /**
     * @return Type of the backend, such as &#34;aws&#34;
     * 
     */
    public Output<String> type() {
        return this.type;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Mount(String name) {
        this(name, MountArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Mount(String name, MountArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Mount(String name, MountArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("vault:index/mount:Mount", name, args == null ? MountArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private Mount(String name, Output<String> id, @Nullable MountState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("vault:index/mount:Mount", name, state, makeResourceOptions(options, id));
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
    public static Mount get(String name, Output<String> id, @Nullable MountState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Mount(name, id, state, options);
    }
}
