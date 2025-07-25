// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vault;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.vault.ScepAuthBackendRoleArgs;
import com.pulumi.vault.Utilities;
import com.pulumi.vault.inputs.ScepAuthBackendRoleState;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
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
 * import com.pulumi.vault.AuthBackend;
 * import com.pulumi.vault.AuthBackendArgs;
 * import com.pulumi.vault.ScepAuthBackendRole;
 * import com.pulumi.vault.ScepAuthBackendRoleArgs;
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
 *         var scep = new AuthBackend("scep", AuthBackendArgs.builder()
 *             .path("scep")
 *             .type("scep")
 *             .build());
 * 
 *         var scepScepAuthBackendRole = new ScepAuthBackendRole("scepScepAuthBackendRole", ScepAuthBackendRoleArgs.builder()
 *             .backend(scep.path())
 *             .name("scep_challenge")
 *             .authType("static-challenge")
 *             .challenge("well known secret")
 *             .tokenType("batch")
 *             .tokenTtl(300)
 *             .tokenMaxTtl(600)
 *             .tokenPolicies("scep-clients")
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 */
@ResourceType(type="vault:index/scepAuthBackendRole:ScepAuthBackendRole")
public class ScepAuthBackendRole extends com.pulumi.resources.CustomResource {
    /**
     * The authentication type to use. This can be either &#34;static-challenge&#34; or &#34;intune&#34;.
     * 
     */
    @Export(name="authType", refs={String.class}, tree="[0]")
    private Output<String> authType;

    /**
     * @return The authentication type to use. This can be either &#34;static-challenge&#34; or &#34;intune&#34;.
     * 
     */
    public Output<String> authType() {
        return this.authType;
    }
    /**
     * Path to the mounted SCEP auth backend.
     * 
     */
    @Export(name="backend", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> backend;

    /**
     * @return Path to the mounted SCEP auth backend.
     * 
     */
    public Output<Optional<String>> backend() {
        return Codegen.optional(this.backend);
    }
    /**
     * The static challenge to use if auth_type is &#34;static-challenge&#34;, not used for other auth types.
     * 
     */
    @Export(name="challenge", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> challenge;

    /**
     * @return The static challenge to use if auth_type is &#34;static-challenge&#34;, not used for other auth types.
     * 
     */
    public Output<Optional<String>> challenge() {
        return Codegen.optional(this.challenge);
    }
    @Export(name="displayName", refs={String.class}, tree="[0]")
    private Output<String> displayName;

    public Output<String> displayName() {
        return this.displayName;
    }
    /**
     * Name of the role.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return Name of the role.
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
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public ScepAuthBackendRole(java.lang.String name) {
        this(name, ScepAuthBackendRoleArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public ScepAuthBackendRole(java.lang.String name, ScepAuthBackendRoleArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public ScepAuthBackendRole(java.lang.String name, ScepAuthBackendRoleArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("vault:index/scepAuthBackendRole:ScepAuthBackendRole", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private ScepAuthBackendRole(java.lang.String name, Output<java.lang.String> id, @Nullable ScepAuthBackendRoleState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("vault:index/scepAuthBackendRole:ScepAuthBackendRole", name, state, makeResourceOptions(options, id), false);
    }

    private static ScepAuthBackendRoleArgs makeArgs(ScepAuthBackendRoleArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? ScepAuthBackendRoleArgs.Empty : args;
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
    public static ScepAuthBackendRole get(java.lang.String name, Output<java.lang.String> id, @Nullable ScepAuthBackendRoleState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new ScepAuthBackendRole(name, id, state, options);
    }
}
