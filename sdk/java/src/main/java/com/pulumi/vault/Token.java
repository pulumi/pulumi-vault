// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vault;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.vault.TokenArgs;
import com.pulumi.vault.Utilities;
import com.pulumi.vault.inputs.TokenState;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
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
 * import com.pulumi.vault.Token;
 * import com.pulumi.vault.TokenArgs;
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
 *         var example = new Token("example", TokenArgs.builder()
 *             .roleName("app")
 *             .policies(            
 *                 "policy1",
 *                 "policy2")
 *             .renewable(true)
 *             .ttl("24h")
 *             .renewMinLease(43200)
 *             .renewIncrement(86400)
 *             .metadata(Map.of("purpose", "service-account"))
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
 * Tokens can be imported using its `id` as accessor id, e.g.
 * 
 * ```sh
 * $ pulumi import vault:index/token:Token example &lt;accessor_id&gt;
 * ```
 * 
 */
@ResourceType(type="vault:index/token:Token")
public class Token extends com.pulumi.resources.CustomResource {
    /**
     * String containing the client token if stored in present file
     * 
     */
    @Export(name="clientToken", refs={String.class}, tree="[0]")
    private Output<String> clientToken;

    /**
     * @return String containing the client token if stored in present file
     * 
     */
    public Output<String> clientToken() {
        return this.clientToken;
    }
    /**
     * String containing the token display name
     * 
     */
    @Export(name="displayName", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> displayName;

    /**
     * @return String containing the token display name
     * 
     */
    public Output<Optional<String>> displayName() {
        return Codegen.optional(this.displayName);
    }
    /**
     * The explicit max TTL of this token. This is specified as a numeric string with suffix like &#34;30s&#34; ro &#34;5m&#34;
     * 
     */
    @Export(name="explicitMaxTtl", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> explicitMaxTtl;

    /**
     * @return The explicit max TTL of this token. This is specified as a numeric string with suffix like &#34;30s&#34; ro &#34;5m&#34;
     * 
     */
    public Output<Optional<String>> explicitMaxTtl() {
        return Codegen.optional(this.explicitMaxTtl);
    }
    /**
     * String containing the token lease duration if present in state file
     * 
     */
    @Export(name="leaseDuration", refs={Integer.class}, tree="[0]")
    private Output<Integer> leaseDuration;

    /**
     * @return String containing the token lease duration if present in state file
     * 
     */
    public Output<Integer> leaseDuration() {
        return this.leaseDuration;
    }
    /**
     * String containing the token lease started time if present in state file
     * 
     */
    @Export(name="leaseStarted", refs={String.class}, tree="[0]")
    private Output<String> leaseStarted;

    /**
     * @return String containing the token lease started time if present in state file
     * 
     */
    public Output<String> leaseStarted() {
        return this.leaseStarted;
    }
    /**
     * Metadata to be set on this token
     * 
     */
    @Export(name="metadata", refs={Map.class,String.class}, tree="[0,1,1]")
    private Output</* @Nullable */ Map<String,String>> metadata;

    /**
     * @return Metadata to be set on this token
     * 
     */
    public Output<Optional<Map<String,String>>> metadata() {
        return Codegen.optional(this.metadata);
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
     * Flag to not attach the default policy to this token
     * 
     */
    @Export(name="noDefaultPolicy", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> noDefaultPolicy;

    /**
     * @return Flag to not attach the default policy to this token
     * 
     */
    public Output<Optional<Boolean>> noDefaultPolicy() {
        return Codegen.optional(this.noDefaultPolicy);
    }
    /**
     * Flag to create a token without parent
     * 
     */
    @Export(name="noParent", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> noParent;

    /**
     * @return Flag to create a token without parent
     * 
     */
    public Output<Boolean> noParent() {
        return this.noParent;
    }
    /**
     * The number of allowed uses of this token
     * 
     */
    @Export(name="numUses", refs={Integer.class}, tree="[0]")
    private Output<Integer> numUses;

    /**
     * @return The number of allowed uses of this token
     * 
     */
    public Output<Integer> numUses() {
        return this.numUses;
    }
    /**
     * The period of this token. This is specified as a numeric string with suffix like &#34;30s&#34; ro &#34;5m&#34;
     * 
     */
    @Export(name="period", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> period;

    /**
     * @return The period of this token. This is specified as a numeric string with suffix like &#34;30s&#34; ro &#34;5m&#34;
     * 
     */
    public Output<Optional<String>> period() {
        return Codegen.optional(this.period);
    }
    /**
     * List of policies to attach to this token
     * 
     */
    @Export(name="policies", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> policies;

    /**
     * @return List of policies to attach to this token
     * 
     */
    public Output<Optional<List<String>>> policies() {
        return Codegen.optional(this.policies);
    }
    /**
     * The renew increment. This is specified in seconds
     * 
     */
    @Export(name="renewIncrement", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> renewIncrement;

    /**
     * @return The renew increment. This is specified in seconds
     * 
     */
    public Output<Optional<Integer>> renewIncrement() {
        return Codegen.optional(this.renewIncrement);
    }
    /**
     * The minimal lease to renew this token
     * 
     */
    @Export(name="renewMinLease", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> renewMinLease;

    /**
     * @return The minimal lease to renew this token
     * 
     */
    public Output<Optional<Integer>> renewMinLease() {
        return Codegen.optional(this.renewMinLease);
    }
    /**
     * Flag to allow to renew this token
     * 
     */
    @Export(name="renewable", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> renewable;

    /**
     * @return Flag to allow to renew this token
     * 
     */
    public Output<Boolean> renewable() {
        return this.renewable;
    }
    /**
     * The token role name
     * 
     */
    @Export(name="roleName", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> roleName;

    /**
     * @return The token role name
     * 
     */
    public Output<Optional<String>> roleName() {
        return Codegen.optional(this.roleName);
    }
    /**
     * The TTL period of this token. This is specified as a numeric string with suffix like &#34;30s&#34; ro &#34;5m&#34;
     * 
     */
    @Export(name="ttl", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> ttl;

    /**
     * @return The TTL period of this token. This is specified as a numeric string with suffix like &#34;30s&#34; ro &#34;5m&#34;
     * 
     */
    public Output<Optional<String>> ttl() {
        return Codegen.optional(this.ttl);
    }
    /**
     * The client wrapped token.
     * 
     */
    @Export(name="wrappedToken", refs={String.class}, tree="[0]")
    private Output<String> wrappedToken;

    /**
     * @return The client wrapped token.
     * 
     */
    public Output<String> wrappedToken() {
        return this.wrappedToken;
    }
    /**
     * The client wrapping accessor.
     * 
     */
    @Export(name="wrappingAccessor", refs={String.class}, tree="[0]")
    private Output<String> wrappingAccessor;

    /**
     * @return The client wrapping accessor.
     * 
     */
    public Output<String> wrappingAccessor() {
        return this.wrappingAccessor;
    }
    /**
     * The TTL period of the wrapped token.
     * 
     */
    @Export(name="wrappingTtl", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> wrappingTtl;

    /**
     * @return The TTL period of the wrapped token.
     * 
     */
    public Output<Optional<String>> wrappingTtl() {
        return Codegen.optional(this.wrappingTtl);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Token(java.lang.String name) {
        this(name, TokenArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Token(java.lang.String name, @Nullable TokenArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Token(java.lang.String name, @Nullable TokenArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("vault:index/token:Token", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private Token(java.lang.String name, Output<java.lang.String> id, @Nullable TokenState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("vault:index/token:Token", name, state, makeResourceOptions(options, id), false);
    }

    private static TokenArgs makeArgs(@Nullable TokenArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? TokenArgs.Empty : args;
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<java.lang.String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .additionalSecretOutputs(List.of(
                "clientToken",
                "wrappedToken",
                "wrappingAccessor"
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
    public static Token get(java.lang.String name, Output<java.lang.String> id, @Nullable TokenState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Token(name, id, state, options);
    }
}
