// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vault.appRole;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.vault.Utilities;
import com.pulumi.vault.appRole.AuthBackendRoleArgs;
import com.pulumi.vault.appRole.inputs.AuthBackendRoleState;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Manages an AppRole auth backend role in a Vault server. See the [Vault
 * documentation](https://www.vaultproject.io/docs/auth/approle) for more
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
 * import com.pulumi.vault.AuthBackend;
 * import com.pulumi.vault.AuthBackendArgs;
 * import com.pulumi.vault.appRole.AuthBackendRole;
 * import com.pulumi.vault.appRole.AuthBackendRoleArgs;
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
 *         var approle = new AuthBackend("approle", AuthBackendArgs.builder()
 *             .type("approle")
 *             .build());
 * 
 *         var example = new AuthBackendRole("example", AuthBackendRoleArgs.builder()
 *             .backend(approle.path())
 *             .roleName("test-role")
 *             .tokenPolicies(            
 *                 "default",
 *                 "dev",
 *                 "prod")
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
 * AppRole authentication backend roles can be imported using the `path`, e.g.
 * 
 * ```sh
 * $ pulumi import vault:appRole/authBackendRole:AuthBackendRole example auth/approle/role/test-role
 * ```
 * 
 */
@ResourceType(type="vault:appRole/authBackendRole:AuthBackendRole")
public class AuthBackendRole extends com.pulumi.resources.CustomResource {
    /**
     * The unique name of the auth backend to configure.
     * Defaults to `approle`.
     * 
     */
    @Export(name="backend", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> backend;

    /**
     * @return The unique name of the auth backend to configure.
     * Defaults to `approle`.
     * 
     */
    public Output<Optional<String>> backend() {
        return Codegen.optional(this.backend);
    }
    /**
     * Whether or not to require `secret_id` to be
     * presented when logging in using this AppRole. Defaults to `true`.
     * 
     */
    @Export(name="bindSecretId", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> bindSecretId;

    /**
     * @return Whether or not to require `secret_id` to be
     * presented when logging in using this AppRole. Defaults to `true`.
     * 
     */
    public Output<Optional<Boolean>> bindSecretId() {
        return Codegen.optional(this.bindSecretId);
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
     * The RoleID of this role. If not specified, one will be
     * auto-generated.
     * 
     */
    @Export(name="roleId", refs={String.class}, tree="[0]")
    private Output<String> roleId;

    /**
     * @return The RoleID of this role. If not specified, one will be
     * auto-generated.
     * 
     */
    public Output<String> roleId() {
        return this.roleId;
    }
    /**
     * The name of the role.
     * 
     */
    @Export(name="roleName", refs={String.class}, tree="[0]")
    private Output<String> roleName;

    /**
     * @return The name of the role.
     * 
     */
    public Output<String> roleName() {
        return this.roleName;
    }
    /**
     * If set,
     * specifies blocks of IP addresses which can perform the login operation.
     * 
     */
    @Export(name="secretIdBoundCidrs", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> secretIdBoundCidrs;

    /**
     * @return If set,
     * specifies blocks of IP addresses which can perform the login operation.
     * 
     */
    public Output<Optional<List<String>>> secretIdBoundCidrs() {
        return Codegen.optional(this.secretIdBoundCidrs);
    }
    /**
     * The number of times any particular SecretID
     * can be used to fetch a token from this AppRole, after which the SecretID will
     * expire. A value of zero will allow unlimited uses.
     * 
     */
    @Export(name="secretIdNumUses", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> secretIdNumUses;

    /**
     * @return The number of times any particular SecretID
     * can be used to fetch a token from this AppRole, after which the SecretID will
     * expire. A value of zero will allow unlimited uses.
     * 
     */
    public Output<Optional<Integer>> secretIdNumUses() {
        return Codegen.optional(this.secretIdNumUses);
    }
    /**
     * The number of seconds after which any SecretID
     * expires.
     * 
     */
    @Export(name="secretIdTtl", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> secretIdTtl;

    /**
     * @return The number of seconds after which any SecretID
     * expires.
     * 
     */
    public Output<Optional<Integer>> secretIdTtl() {
        return Codegen.optional(this.secretIdTtl);
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
    public AuthBackendRole(java.lang.String name) {
        this(name, AuthBackendRoleArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public AuthBackendRole(java.lang.String name, AuthBackendRoleArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public AuthBackendRole(java.lang.String name, AuthBackendRoleArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("vault:appRole/authBackendRole:AuthBackendRole", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private AuthBackendRole(java.lang.String name, Output<java.lang.String> id, @Nullable AuthBackendRoleState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("vault:appRole/authBackendRole:AuthBackendRole", name, state, makeResourceOptions(options, id), false);
    }

    private static AuthBackendRoleArgs makeArgs(AuthBackendRoleArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? AuthBackendRoleArgs.Empty : args;
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
    public static AuthBackendRole get(java.lang.String name, Output<java.lang.String> id, @Nullable AuthBackendRoleState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new AuthBackendRole(name, id, state, options);
    }
}
