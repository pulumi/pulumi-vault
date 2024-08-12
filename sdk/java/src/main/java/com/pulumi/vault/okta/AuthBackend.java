// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vault.okta;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.vault.Utilities;
import com.pulumi.vault.okta.AuthBackendArgs;
import com.pulumi.vault.okta.inputs.AuthBackendState;
import com.pulumi.vault.okta.outputs.AuthBackendGroup;
import com.pulumi.vault.okta.outputs.AuthBackendUser;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides a resource for managing an
 * [Okta auth backend within Vault](https://www.vaultproject.io/docs/auth/okta.html).
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
 * import com.pulumi.vault.okta.AuthBackend;
 * import com.pulumi.vault.okta.AuthBackendArgs;
 * import com.pulumi.vault.okta.inputs.AuthBackendGroupArgs;
 * import com.pulumi.vault.okta.inputs.AuthBackendUserArgs;
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
 *         var example = new AuthBackend("example", AuthBackendArgs.builder()
 *             .description("Demonstration of the Terraform Okta auth backend")
 *             .organization("example")
 *             .token("something that should be kept secret")
 *             .groups(AuthBackendGroupArgs.builder()
 *                 .groupName("foo")
 *                 .policies(                
 *                     "one",
 *                     "two")
 *                 .build())
 *             .users(AuthBackendUserArgs.builder()
 *                 .username("bar")
 *                 .groups("foo")
 *                 .build())
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
 * Okta authentication backends can be imported using its `path`, e.g.
 * 
 * ```sh
 * $ pulumi import vault:okta/authBackend:AuthBackend example okta
 * ```
 * 
 */
@ResourceType(type="vault:okta/authBackend:AuthBackend")
public class AuthBackend extends com.pulumi.resources.CustomResource {
    /**
     * The mount accessor related to the auth mount. It is useful for integration with [Identity Secrets Engine](https://www.vaultproject.io/docs/secrets/identity/index.html).
     * 
     */
    @Export(name="accessor", refs={String.class}, tree="[0]")
    private Output<String> accessor;

    /**
     * @return The mount accessor related to the auth mount. It is useful for integration with [Identity Secrets Engine](https://www.vaultproject.io/docs/secrets/identity/index.html).
     * 
     */
    public Output<String> accessor() {
        return this.accessor;
    }
    /**
     * The Okta url. Examples: oktapreview.com, okta.com
     * 
     */
    @Export(name="baseUrl", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> baseUrl;

    /**
     * @return The Okta url. Examples: oktapreview.com, okta.com
     * 
     */
    public Output<Optional<String>> baseUrl() {
        return Codegen.optional(this.baseUrl);
    }
    /**
     * When true, requests by Okta for a MFA check will be bypassed. This also disallows certain status checks on the account, such as whether the password is expired.
     * 
     */
    @Export(name="bypassOktaMfa", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> bypassOktaMfa;

    /**
     * @return When true, requests by Okta for a MFA check will be bypassed. This also disallows certain status checks on the account, such as whether the password is expired.
     * 
     */
    public Output<Optional<Boolean>> bypassOktaMfa() {
        return Codegen.optional(this.bypassOktaMfa);
    }
    /**
     * The description of the auth backend
     * 
     */
    @Export(name="description", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> description;

    /**
     * @return The description of the auth backend
     * 
     */
    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    /**
     * If set, opts out of mount migration on path updates.
     * See here for more info on [Mount Migration](https://www.vaultproject.io/docs/concepts/mount-migration)
     * 
     */
    @Export(name="disableRemount", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> disableRemount;

    /**
     * @return If set, opts out of mount migration on path updates.
     * See here for more info on [Mount Migration](https://www.vaultproject.io/docs/concepts/mount-migration)
     * 
     */
    public Output<Optional<Boolean>> disableRemount() {
        return Codegen.optional(this.disableRemount);
    }
    /**
     * Associate Okta groups with policies within Vault.
     * See below for more details.
     * 
     */
    @Export(name="groups", refs={List.class,AuthBackendGroup.class}, tree="[0,1]")
    private Output<List<AuthBackendGroup>> groups;

    /**
     * @return Associate Okta groups with policies within Vault.
     * See below for more details.
     * 
     */
    public Output<List<AuthBackendGroup>> groups() {
        return this.groups;
    }
    /**
     * Maximum duration after which authentication will be expired
     * [See the documentation for info on valid duration formats](https://golang.org/pkg/time/#ParseDuration).
     * 
     * @deprecated
     * Deprecated. Please use `token_max_ttl` instead.
     * 
     */
    @Deprecated /* Deprecated. Please use `token_max_ttl` instead. */
    @Export(name="maxTtl", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> maxTtl;

    /**
     * @return Maximum duration after which authentication will be expired
     * [See the documentation for info on valid duration formats](https://golang.org/pkg/time/#ParseDuration).
     * 
     */
    public Output<Optional<String>> maxTtl() {
        return Codegen.optional(this.maxTtl);
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
     * The Okta organization. This will be the first part of the url `https://XXX.okta.com`
     * 
     */
    @Export(name="organization", refs={String.class}, tree="[0]")
    private Output<String> organization;

    /**
     * @return The Okta organization. This will be the first part of the url `https://XXX.okta.com`
     * 
     */
    public Output<String> organization() {
        return this.organization;
    }
    /**
     * Path to mount the Okta auth backend. Default to path `okta`.
     * 
     */
    @Export(name="path", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> path;

    /**
     * @return Path to mount the Okta auth backend. Default to path `okta`.
     * 
     */
    public Output<Optional<String>> path() {
        return Codegen.optional(this.path);
    }
    /**
     * The Okta API token. This is required to query Okta for user group membership.
     * If this is not supplied only locally configured groups will be enabled.
     * 
     */
    @Export(name="token", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> token;

    /**
     * @return The Okta API token. This is required to query Okta for user group membership.
     * If this is not supplied only locally configured groups will be enabled.
     * 
     */
    public Output<Optional<String>> token() {
        return Codegen.optional(this.token);
    }
    /**
     * Specifies the blocks of IP addresses which are allowed to use the generated token
     * 
     */
    @Export(name="tokenBoundCidrs", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> tokenBoundCidrs;

    /**
     * @return Specifies the blocks of IP addresses which are allowed to use the generated token
     * 
     */
    public Output<Optional<List<String>>> tokenBoundCidrs() {
        return Codegen.optional(this.tokenBoundCidrs);
    }
    /**
     * Generated Token&#39;s Explicit Maximum TTL in seconds
     * 
     */
    @Export(name="tokenExplicitMaxTtl", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> tokenExplicitMaxTtl;

    /**
     * @return Generated Token&#39;s Explicit Maximum TTL in seconds
     * 
     */
    public Output<Optional<Integer>> tokenExplicitMaxTtl() {
        return Codegen.optional(this.tokenExplicitMaxTtl);
    }
    /**
     * The maximum lifetime of the generated token
     * 
     */
    @Export(name="tokenMaxTtl", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> tokenMaxTtl;

    /**
     * @return The maximum lifetime of the generated token
     * 
     */
    public Output<Optional<Integer>> tokenMaxTtl() {
        return Codegen.optional(this.tokenMaxTtl);
    }
    /**
     * If true, the &#39;default&#39; policy will not automatically be added to generated tokens
     * 
     */
    @Export(name="tokenNoDefaultPolicy", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> tokenNoDefaultPolicy;

    /**
     * @return If true, the &#39;default&#39; policy will not automatically be added to generated tokens
     * 
     */
    public Output<Optional<Boolean>> tokenNoDefaultPolicy() {
        return Codegen.optional(this.tokenNoDefaultPolicy);
    }
    /**
     * The maximum number of times a token may be used, a value of zero means unlimited
     * 
     */
    @Export(name="tokenNumUses", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> tokenNumUses;

    /**
     * @return The maximum number of times a token may be used, a value of zero means unlimited
     * 
     */
    public Output<Optional<Integer>> tokenNumUses() {
        return Codegen.optional(this.tokenNumUses);
    }
    /**
     * Generated Token&#39;s Period
     * 
     */
    @Export(name="tokenPeriod", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> tokenPeriod;

    /**
     * @return Generated Token&#39;s Period
     * 
     */
    public Output<Optional<Integer>> tokenPeriod() {
        return Codegen.optional(this.tokenPeriod);
    }
    /**
     * Generated Token&#39;s Policies
     * 
     */
    @Export(name="tokenPolicies", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> tokenPolicies;

    /**
     * @return Generated Token&#39;s Policies
     * 
     */
    public Output<Optional<List<String>>> tokenPolicies() {
        return Codegen.optional(this.tokenPolicies);
    }
    /**
     * The initial ttl of the token to generate in seconds
     * 
     */
    @Export(name="tokenTtl", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> tokenTtl;

    /**
     * @return The initial ttl of the token to generate in seconds
     * 
     */
    public Output<Optional<Integer>> tokenTtl() {
        return Codegen.optional(this.tokenTtl);
    }
    /**
     * The type of token to generate, service or batch
     * 
     */
    @Export(name="tokenType", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> tokenType;

    /**
     * @return The type of token to generate, service or batch
     * 
     */
    public Output<Optional<String>> tokenType() {
        return Codegen.optional(this.tokenType);
    }
    /**
     * Duration after which authentication will be expired.
     * [See the documentation for info on valid duration formats](https://golang.org/pkg/time/#ParseDuration).
     * 
     * @deprecated
     * Deprecated. Please use `token_ttl` instead.
     * 
     */
    @Deprecated /* Deprecated. Please use `token_ttl` instead. */
    @Export(name="ttl", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> ttl;

    /**
     * @return Duration after which authentication will be expired.
     * [See the documentation for info on valid duration formats](https://golang.org/pkg/time/#ParseDuration).
     * 
     */
    public Output<Optional<String>> ttl() {
        return Codegen.optional(this.ttl);
    }
    /**
     * Associate Okta users with groups or policies within Vault.
     * See below for more details.
     * 
     */
    @Export(name="users", refs={List.class,AuthBackendUser.class}, tree="[0,1]")
    private Output<List<AuthBackendUser>> users;

    /**
     * @return Associate Okta users with groups or policies within Vault.
     * See below for more details.
     * 
     */
    public Output<List<AuthBackendUser>> users() {
        return this.users;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public AuthBackend(java.lang.String name) {
        this(name, AuthBackendArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public AuthBackend(java.lang.String name, AuthBackendArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public AuthBackend(java.lang.String name, AuthBackendArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("vault:okta/authBackend:AuthBackend", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private AuthBackend(java.lang.String name, Output<java.lang.String> id, @Nullable AuthBackendState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("vault:okta/authBackend:AuthBackend", name, state, makeResourceOptions(options, id), false);
    }

    private static AuthBackendArgs makeArgs(AuthBackendArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? AuthBackendArgs.Empty : args;
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<java.lang.String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .additionalSecretOutputs(List.of(
                "token"
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
    public static AuthBackend get(java.lang.String name, Output<java.lang.String> id, @Nullable AuthBackendState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new AuthBackend(name, id, state, options);
    }
}
