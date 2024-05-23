// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vault.appRole;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.vault.Utilities;
import com.pulumi.vault.appRole.AuthBackendLoginArgs;
import com.pulumi.vault.appRole.inputs.AuthBackendLoginState;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Logs into Vault using the AppRole auth backend. See the [Vault
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
 * import com.pulumi.vault.appRole.AuthBackendRoleSecretId;
 * import com.pulumi.vault.appRole.AuthBackendRoleSecretIdArgs;
 * import com.pulumi.vault.appRole.AuthBackendLogin;
 * import com.pulumi.vault.appRole.AuthBackendLoginArgs;
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
 *         var id = new AuthBackendRoleSecretId("id", AuthBackendRoleSecretIdArgs.builder()
 *             .backend(approle.path())
 *             .roleName(example.roleName())
 *             .build());
 * 
 *         var login = new AuthBackendLogin("login", AuthBackendLoginArgs.builder()
 *             .backend(approle.path())
 *             .roleId(example.roleId())
 *             .secretId(id.secretId())
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 */
@ResourceType(type="vault:appRole/authBackendLogin:AuthBackendLogin")
public class AuthBackendLogin extends com.pulumi.resources.CustomResource {
    /**
     * The accessor for the token.
     * 
     */
    @Export(name="accessor", refs={String.class}, tree="[0]")
    private Output<String> accessor;

    /**
     * @return The accessor for the token.
     * 
     */
    public Output<String> accessor() {
        return this.accessor;
    }
    /**
     * The unique path of the Vault backend to log in with.
     * 
     */
    @Export(name="backend", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> backend;

    /**
     * @return The unique path of the Vault backend to log in with.
     * 
     */
    public Output<Optional<String>> backend() {
        return Codegen.optional(this.backend);
    }
    /**
     * The Vault token created.
     * 
     */
    @Export(name="clientToken", refs={String.class}, tree="[0]")
    private Output<String> clientToken;

    /**
     * @return The Vault token created.
     * 
     */
    public Output<String> clientToken() {
        return this.clientToken;
    }
    /**
     * How long the token is valid for, in seconds.
     * 
     */
    @Export(name="leaseDuration", refs={Integer.class}, tree="[0]")
    private Output<Integer> leaseDuration;

    /**
     * @return How long the token is valid for, in seconds.
     * 
     */
    public Output<Integer> leaseDuration() {
        return this.leaseDuration;
    }
    /**
     * The date and time the lease started, in RFC 3339 format.
     * 
     */
    @Export(name="leaseStarted", refs={String.class}, tree="[0]")
    private Output<String> leaseStarted;

    /**
     * @return The date and time the lease started, in RFC 3339 format.
     * 
     */
    public Output<String> leaseStarted() {
        return this.leaseStarted;
    }
    /**
     * The metadata associated with the token.
     * 
     */
    @Export(name="metadata", refs={Map.class,String.class}, tree="[0,1,1]")
    private Output<Map<String,String>> metadata;

    /**
     * @return The metadata associated with the token.
     * 
     */
    public Output<Map<String,String>> metadata() {
        return this.metadata;
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
     * A list of policies applied to the token.
     * 
     */
    @Export(name="policies", refs={List.class,String.class}, tree="[0,1]")
    private Output<List<String>> policies;

    /**
     * @return A list of policies applied to the token.
     * 
     */
    public Output<List<String>> policies() {
        return this.policies;
    }
    /**
     * Whether the token is renewable or not.
     * 
     */
    @Export(name="renewable", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> renewable;

    /**
     * @return Whether the token is renewable or not.
     * 
     */
    public Output<Boolean> renewable() {
        return this.renewable;
    }
    /**
     * The ID of the role to log in with.
     * 
     */
    @Export(name="roleId", refs={String.class}, tree="[0]")
    private Output<String> roleId;

    /**
     * @return The ID of the role to log in with.
     * 
     */
    public Output<String> roleId() {
        return this.roleId;
    }
    /**
     * The secret ID of the role to log in with. Required
     * unless `bind_secret_id` is set to false on the role.
     * 
     */
    @Export(name="secretId", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> secretId;

    /**
     * @return The secret ID of the role to log in with. Required
     * unless `bind_secret_id` is set to false on the role.
     * 
     */
    public Output<Optional<String>> secretId() {
        return Codegen.optional(this.secretId);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public AuthBackendLogin(String name) {
        this(name, AuthBackendLoginArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public AuthBackendLogin(String name, AuthBackendLoginArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public AuthBackendLogin(String name, AuthBackendLoginArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("vault:appRole/authBackendLogin:AuthBackendLogin", name, args == null ? AuthBackendLoginArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private AuthBackendLogin(String name, Output<String> id, @Nullable AuthBackendLoginState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("vault:appRole/authBackendLogin:AuthBackendLogin", name, state, makeResourceOptions(options, id));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .additionalSecretOutputs(List.of(
                "clientToken",
                "secretId"
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
    public static AuthBackendLogin get(String name, Output<String> id, @Nullable AuthBackendLoginState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new AuthBackendLogin(name, id, state, options);
    }
}
