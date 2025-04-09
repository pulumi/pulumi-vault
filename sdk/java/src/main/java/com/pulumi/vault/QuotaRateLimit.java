// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vault;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.vault.QuotaRateLimitArgs;
import com.pulumi.vault.Utilities;
import com.pulumi.vault.inputs.QuotaRateLimitState;
import java.lang.Boolean;
import java.lang.Double;
import java.lang.Integer;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Manage rate limit quotas which enforce API rate limiting using a token bucket algorithm.
 * A rate limit quota can be created at the root level or defined on a namespace or mount by
 * specifying a path when creating the quota.
 * 
 * See [Vault&#39;s Documentation](https://www.vaultproject.io/docs/concepts/resource-quotas) for more
 * information.
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
 * import com.pulumi.vault.QuotaRateLimit;
 * import com.pulumi.vault.QuotaRateLimitArgs;
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
 *         var global = new QuotaRateLimit("global", QuotaRateLimitArgs.builder()
 *             .name("global")
 *             .path("")
 *             .rate(100.0)
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Import
 * 
 * Rate limit quotas can be imported using their names
 * 
 * ```sh
 * $ pulumi import vault:index/quotaRateLimit:QuotaRateLimit global global
 * ```
 * 
 */
@ResourceType(type="vault:index/quotaRateLimit:QuotaRateLimit")
public class QuotaRateLimit extends com.pulumi.resources.CustomResource {
    /**
     * If set, when a client reaches a rate limit threshold, the client will
     * be prohibited from any further requests until after the &#39;block_interval&#39; in seconds has elapsed.
     * 
     */
    @Export(name="blockInterval", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> blockInterval;

    /**
     * @return If set, when a client reaches a rate limit threshold, the client will
     * be prohibited from any further requests until after the &#39;block_interval&#39; in seconds has elapsed.
     * 
     */
    public Output<Optional<Integer>> blockInterval() {
        return Codegen.optional(this.blockInterval);
    }
    /**
     * If set to `true` on a quota where path is set to a namespace, the same quota will be cumulatively applied to all child namespace. The inheritable parameter cannot be set to `true` if the path does not specify a namespace. Only the quotas associated with the root namespace are inheritable by default. Requires Vault 1.15+.
     * 
     */
    @Export(name="inheritable", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> inheritable;

    /**
     * @return If set to `true` on a quota where path is set to a namespace, the same quota will be cumulatively applied to all child namespace. The inheritable parameter cannot be set to `true` if the path does not specify a namespace. Only the quotas associated with the root namespace are inheritable by default. Requires Vault 1.15+.
     * 
     */
    public Output<Optional<Boolean>> inheritable() {
        return Codegen.optional(this.inheritable);
    }
    /**
     * The duration in seconds to enforce rate limiting for.
     * 
     */
    @Export(name="interval", refs={Integer.class}, tree="[0]")
    private Output<Integer> interval;

    /**
     * @return The duration in seconds to enforce rate limiting for.
     * 
     */
    public Output<Integer> interval() {
        return this.interval;
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
     * The maximum number of requests at any given second to be allowed by the quota
     * rule. The `rate` must be positive.
     * 
     */
    @Export(name="rate", refs={Double.class}, tree="[0]")
    private Output<Double> rate;

    /**
     * @return The maximum number of requests at any given second to be allowed by the quota
     * rule. The `rate` must be positive.
     * 
     */
    public Output<Double> rate() {
        return this.rate;
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
    public QuotaRateLimit(java.lang.String name) {
        this(name, QuotaRateLimitArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public QuotaRateLimit(java.lang.String name, QuotaRateLimitArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public QuotaRateLimit(java.lang.String name, QuotaRateLimitArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("vault:index/quotaRateLimit:QuotaRateLimit", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private QuotaRateLimit(java.lang.String name, Output<java.lang.String> id, @Nullable QuotaRateLimitState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("vault:index/quotaRateLimit:QuotaRateLimit", name, state, makeResourceOptions(options, id), false);
    }

    private static QuotaRateLimitArgs makeArgs(QuotaRateLimitArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? QuotaRateLimitArgs.Empty : args;
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
    public static QuotaRateLimit get(java.lang.String name, Output<java.lang.String> id, @Nullable QuotaRateLimitState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new QuotaRateLimit(name, id, state, options);
    }
}
