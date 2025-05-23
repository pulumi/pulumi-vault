// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vault.ad;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.vault.Utilities;
import com.pulumi.vault.ad.SecretRoleArgs;
import com.pulumi.vault.ad.inputs.SecretRoleState;
import java.lang.Integer;
import java.lang.String;
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
 * import com.pulumi.vault.ad.SecretBackend;
 * import com.pulumi.vault.ad.SecretBackendArgs;
 * import com.pulumi.vault.ad.SecretRole;
 * import com.pulumi.vault.ad.SecretRoleArgs;
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
 *         var config = new SecretBackend("config", SecretBackendArgs.builder()
 *             .backend("ad")
 *             .binddn("CN=Administrator,CN=Users,DC=corp,DC=example,DC=net")
 *             .bindpass("SuperSecretPassw0rd")
 *             .url("ldaps://ad")
 *             .insecureTls(true)
 *             .userdn("CN=Users,DC=corp,DC=example,DC=net")
 *             .build());
 * 
 *         var role = new SecretRole("role", SecretRoleArgs.builder()
 *             .backend(config.backend())
 *             .role("bob")
 *             .serviceAccountName("Bob")
 *             .ttl(60)
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
 * AD secret backend roles can be imported using the `path`, e.g.
 * 
 * ```sh
 * $ pulumi import vault:ad/secretRole:SecretRole role ad/roles/bob
 * ```
 * 
 */
@ResourceType(type="vault:ad/secretRole:SecretRole")
public class SecretRole extends com.pulumi.resources.CustomResource {
    /**
     * The path the AD secret backend is mounted at,
     * with no leading or trailing `/`s.
     * 
     */
    @Export(name="backend", refs={String.class}, tree="[0]")
    private Output<String> backend;

    /**
     * @return The path the AD secret backend is mounted at,
     * with no leading or trailing `/`s.
     * 
     */
    public Output<String> backend() {
        return this.backend;
    }
    /**
     * Timestamp of the last password rotation by Vault.
     * 
     */
    @Export(name="lastVaultRotation", refs={String.class}, tree="[0]")
    private Output<String> lastVaultRotation;

    /**
     * @return Timestamp of the last password rotation by Vault.
     * 
     */
    public Output<String> lastVaultRotation() {
        return this.lastVaultRotation;
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
     * Timestamp of the last password set by Vault.
     * 
     */
    @Export(name="passwordLastSet", refs={String.class}, tree="[0]")
    private Output<String> passwordLastSet;

    /**
     * @return Timestamp of the last password set by Vault.
     * 
     */
    public Output<String> passwordLastSet() {
        return this.passwordLastSet;
    }
    /**
     * The name to identify this role within the backend.
     * Must be unique within the backend.
     * 
     */
    @Export(name="role", refs={String.class}, tree="[0]")
    private Output<String> role;

    /**
     * @return The name to identify this role within the backend.
     * Must be unique within the backend.
     * 
     */
    public Output<String> role() {
        return this.role;
    }
    /**
     * Specifies the name of the Active Directory service
     * account mapped to this role.
     * 
     */
    @Export(name="serviceAccountName", refs={String.class}, tree="[0]")
    private Output<String> serviceAccountName;

    /**
     * @return Specifies the name of the Active Directory service
     * account mapped to this role.
     * 
     */
    public Output<String> serviceAccountName() {
        return this.serviceAccountName;
    }
    /**
     * The password time-to-live in seconds. Defaults to the configuration
     * ttl if not provided.
     * 
     */
    @Export(name="ttl", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> ttl;

    /**
     * @return The password time-to-live in seconds. Defaults to the configuration
     * ttl if not provided.
     * 
     */
    public Output<Optional<Integer>> ttl() {
        return Codegen.optional(this.ttl);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public SecretRole(java.lang.String name) {
        this(name, SecretRoleArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public SecretRole(java.lang.String name, SecretRoleArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public SecretRole(java.lang.String name, SecretRoleArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("vault:ad/secretRole:SecretRole", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private SecretRole(java.lang.String name, Output<java.lang.String> id, @Nullable SecretRoleState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("vault:ad/secretRole:SecretRole", name, state, makeResourceOptions(options, id), false);
    }

    private static SecretRoleArgs makeArgs(SecretRoleArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? SecretRoleArgs.Empty : args;
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
    public static SecretRole get(java.lang.String name, Output<java.lang.String> id, @Nullable SecretRoleState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new SecretRole(name, id, state, options);
    }
}
