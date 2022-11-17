// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vault.aws;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.vault.Utilities;
import com.pulumi.vault.aws.AuthBackendRoletagBlacklistArgs;
import com.pulumi.vault.aws.inputs.AuthBackendRoletagBlacklistState;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Configures the periodic tidying operation of the blacklisted role tag entries.
 * 
 * ## Example Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.vault.AuthBackend;
 * import com.pulumi.vault.AuthBackendArgs;
 * import com.pulumi.vault.aws.AuthBackendRoletagBlacklist;
 * import com.pulumi.vault.aws.AuthBackendRoletagBlacklistArgs;
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
 *         var exampleAuthBackend = new AuthBackend(&#34;exampleAuthBackend&#34;, AuthBackendArgs.builder()        
 *             .type(&#34;aws&#34;)
 *             .build());
 * 
 *         var exampleAuthBackendRoletagBlacklist = new AuthBackendRoletagBlacklist(&#34;exampleAuthBackendRoletagBlacklist&#34;, AuthBackendRoletagBlacklistArgs.builder()        
 *             .backend(exampleAuthBackend.path())
 *             .safetyBuffer(360)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 */
@ResourceType(type="vault:aws/authBackendRoletagBlacklist:AuthBackendRoletagBlacklist")
public class AuthBackendRoletagBlacklist extends com.pulumi.resources.CustomResource {
    /**
     * The path the AWS auth backend being configured was
     * mounted at.
     * 
     */
    @Export(name="backend", type=String.class, parameters={})
    private Output<String> backend;

    /**
     * @return The path the AWS auth backend being configured was
     * mounted at.
     * 
     */
    public Output<String> backend() {
        return this.backend;
    }
    /**
     * If set to true, disables the periodic
     * tidying of the roletag blacklist entries. Defaults to false.
     * 
     */
    @Export(name="disablePeriodicTidy", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> disablePeriodicTidy;

    /**
     * @return If set to true, disables the periodic
     * tidying of the roletag blacklist entries. Defaults to false.
     * 
     */
    public Output<Optional<Boolean>> disablePeriodicTidy() {
        return Codegen.optional(this.disablePeriodicTidy);
    }
    /**
     * The namespace to provision the resource in.
     * The value should not contain leading or trailing forward slashes.
     * The `namespace` is always relative to the provider&#39;s configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
     * *Available only for Vault Enterprise*.
     * 
     */
    @Export(name="namespace", type=String.class, parameters={})
    private Output</* @Nullable */ String> namespace;

    /**
     * @return The namespace to provision the resource in.
     * The value should not contain leading or trailing forward slashes.
     * The `namespace` is always relative to the provider&#39;s configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
     * *Available only for Vault Enterprise*.
     * 
     */
    public Output<Optional<String>> namespace() {
        return Codegen.optional(this.namespace);
    }
    /**
     * The amount of extra time that must have passed
     * beyond the roletag expiration, before it is removed from the backend storage.
     * Defaults to 259,200 seconds, or 72 hours.
     * 
     */
    @Export(name="safetyBuffer", type=Integer.class, parameters={})
    private Output</* @Nullable */ Integer> safetyBuffer;

    /**
     * @return The amount of extra time that must have passed
     * beyond the roletag expiration, before it is removed from the backend storage.
     * Defaults to 259,200 seconds, or 72 hours.
     * 
     */
    public Output<Optional<Integer>> safetyBuffer() {
        return Codegen.optional(this.safetyBuffer);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public AuthBackendRoletagBlacklist(String name) {
        this(name, AuthBackendRoletagBlacklistArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public AuthBackendRoletagBlacklist(String name, AuthBackendRoletagBlacklistArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public AuthBackendRoletagBlacklist(String name, AuthBackendRoletagBlacklistArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("vault:aws/authBackendRoletagBlacklist:AuthBackendRoletagBlacklist", name, args == null ? AuthBackendRoletagBlacklistArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private AuthBackendRoletagBlacklist(String name, Output<String> id, @Nullable AuthBackendRoletagBlacklistState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("vault:aws/authBackendRoletagBlacklist:AuthBackendRoletagBlacklist", name, state, makeResourceOptions(options, id));
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
    public static AuthBackendRoletagBlacklist get(String name, Output<String> id, @Nullable AuthBackendRoletagBlacklistState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new AuthBackendRoletagBlacklist(name, id, state, options);
    }
}