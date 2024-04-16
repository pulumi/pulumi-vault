// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vault;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.vault.QuotaLeaseCountArgs;
import com.pulumi.vault.Utilities;
import com.pulumi.vault.inputs.QuotaLeaseCountState;
import java.lang.Integer;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Manage lease count quotas which enforce the number of leases that can be created.
 * A lease count quota can be created at the root level or defined on a namespace or mount by
 * specifying a path when creating the quota.
 * 
 * See [Vault&#39;s Documentation](https://www.vaultproject.io/docs/enterprise/lease-count-quotas) for more
 * information.
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
 * import com.pulumi.vault.QuotaLeaseCount;
 * import com.pulumi.vault.QuotaLeaseCountArgs;
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
 *         var global = new QuotaLeaseCount(&#34;global&#34;, QuotaLeaseCountArgs.builder()        
 *             .name(&#34;global&#34;)
 *             .path(&#34;&#34;)
 *             .maxLeases(100)
 *             .build());
 * 
 *     }
 * }
 * ```
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Import
 * 
 * Lease count quotas can be imported using their names
 * 
 * ```sh
 * $ pulumi import vault:index/quotaLeaseCount:QuotaLeaseCount global global
 * ```
 * 
 */
@ResourceType(type="vault:index/quotaLeaseCount:QuotaLeaseCount")
public class QuotaLeaseCount extends com.pulumi.resources.CustomResource {
    /**
     * The maximum number of leases to be allowed by the quota
     * rule. The `max_leases` must be positive.
     * 
     */
    @Export(name="maxLeases", refs={Integer.class}, tree="[0]")
    private Output<Integer> maxLeases;

    /**
     * @return The maximum number of leases to be allowed by the quota
     * rule. The `max_leases` must be positive.
     * 
     */
    public Output<Integer> maxLeases() {
        return this.maxLeases;
    }
    /**
     * Name of the rate limit quota
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return Name of the rate limit quota
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * The namespace to provision the resource in.
     * The value should not contain leading or trailing forward slashes.
     * The `namespace` is always relative to the provider&#39;s configured namespace.
     * *Available only for Vault Enterprise*.
     * 
     */
    @Export(name="namespace", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> namespace;

    /**
     * @return The namespace to provision the resource in.
     * The value should not contain leading or trailing forward slashes.
     * The `namespace` is always relative to the provider&#39;s configured namespace.
     * *Available only for Vault Enterprise*.
     * 
     */
    public Output<Optional<String>> namespace() {
        return Codegen.optional(this.namespace);
    }
    /**
     * Path of the mount or namespace to apply the quota. A blank path configures a
     * global rate limit quota. For example `namespace1/` adds a quota to a full namespace,
     * `namespace1/auth/userpass` adds a `quota` to `userpass` in `namespace1`.
     * Updating this field on an existing quota can have &#34;moving&#34; effects. For example, updating
     * `auth/userpass` to `namespace1/auth/userpass` moves this quota from being a global mount quota to
     * a namespace specific mount quota. **Note, namespaces are supported in Enterprise only.**
     * 
     */
    @Export(name="path", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> path;

    /**
     * @return Path of the mount or namespace to apply the quota. A blank path configures a
     * global rate limit quota. For example `namespace1/` adds a quota to a full namespace,
     * `namespace1/auth/userpass` adds a `quota` to `userpass` in `namespace1`.
     * Updating this field on an existing quota can have &#34;moving&#34; effects. For example, updating
     * `auth/userpass` to `namespace1/auth/userpass` moves this quota from being a global mount quota to
     * a namespace specific mount quota. **Note, namespaces are supported in Enterprise only.**
     * 
     */
    public Output<Optional<String>> path() {
        return Codegen.optional(this.path);
    }
    /**
     * If set on a quota where `path` is set to an auth mount with a concept of roles (such as /auth/approle/), this will make the quota restrict login requests to that mount that are made with the specified role.
     * 
     */
    @Export(name="role", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> role;

    /**
     * @return If set on a quota where `path` is set to an auth mount with a concept of roles (such as /auth/approle/), this will make the quota restrict login requests to that mount that are made with the specified role.
     * 
     */
    public Output<Optional<String>> role() {
        return Codegen.optional(this.role);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public QuotaLeaseCount(String name) {
        this(name, QuotaLeaseCountArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public QuotaLeaseCount(String name, QuotaLeaseCountArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public QuotaLeaseCount(String name, QuotaLeaseCountArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("vault:index/quotaLeaseCount:QuotaLeaseCount", name, args == null ? QuotaLeaseCountArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private QuotaLeaseCount(String name, Output<String> id, @Nullable QuotaLeaseCountState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("vault:index/quotaLeaseCount:QuotaLeaseCount", name, state, makeResourceOptions(options, id));
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
    public static QuotaLeaseCount get(String name, Output<String> id, @Nullable QuotaLeaseCountState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new QuotaLeaseCount(name, id, state, options);
    }
}
