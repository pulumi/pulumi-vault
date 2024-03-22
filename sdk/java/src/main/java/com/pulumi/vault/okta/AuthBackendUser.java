// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vault.okta;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.vault.Utilities;
import com.pulumi.vault.okta.AuthBackendUserArgs;
import com.pulumi.vault.okta.inputs.AuthBackendUserState;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides a resource to create a user in an
 * [Okta auth backend within Vault](https://www.vaultproject.io/docs/auth/okta.html).
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
 * import com.pulumi.vault.okta.AuthBackend;
 * import com.pulumi.vault.okta.AuthBackendArgs;
 * import com.pulumi.vault.okta.AuthBackendUser;
 * import com.pulumi.vault.okta.AuthBackendUserArgs;
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
 *         var example = new AuthBackend(&#34;example&#34;, AuthBackendArgs.builder()        
 *             .path(&#34;user_okta&#34;)
 *             .organization(&#34;dummy&#34;)
 *             .build());
 * 
 *         var foo = new AuthBackendUser(&#34;foo&#34;, AuthBackendUserArgs.builder()        
 *             .path(example.path())
 *             .username(&#34;foo&#34;)
 *             .groups(            
 *                 &#34;one&#34;,
 *                 &#34;two&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Import
 * 
 * Okta authentication backend users can be imported using its `path/user` ID format, e.g.
 * 
 * ```sh
 * $ pulumi import vault:okta/authBackendUser:AuthBackendUser example okta/foo
 * ```
 * 
 */
@ResourceType(type="vault:okta/authBackendUser:AuthBackendUser")
public class AuthBackendUser extends com.pulumi.resources.CustomResource {
    /**
     * List of Okta groups to associate with this user
     * 
     */
    @Export(name="groups", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> groups;

    /**
     * @return List of Okta groups to associate with this user
     * 
     */
    public Output<Optional<List<String>>> groups() {
        return Codegen.optional(this.groups);
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
     * The path where the Okta auth backend is mounted
     * 
     */
    @Export(name="path", refs={String.class}, tree="[0]")
    private Output<String> path;

    /**
     * @return The path where the Okta auth backend is mounted
     * 
     */
    public Output<String> path() {
        return this.path;
    }
    /**
     * List of Vault policies to associate with this user
     * 
     */
    @Export(name="policies", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> policies;

    /**
     * @return List of Vault policies to associate with this user
     * 
     */
    public Output<Optional<List<String>>> policies() {
        return Codegen.optional(this.policies);
    }
    /**
     * Name of the user within Okta
     * 
     */
    @Export(name="username", refs={String.class}, tree="[0]")
    private Output<String> username;

    /**
     * @return Name of the user within Okta
     * 
     */
    public Output<String> username() {
        return this.username;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public AuthBackendUser(String name) {
        this(name, AuthBackendUserArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public AuthBackendUser(String name, AuthBackendUserArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public AuthBackendUser(String name, AuthBackendUserArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("vault:okta/authBackendUser:AuthBackendUser", name, args == null ? AuthBackendUserArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private AuthBackendUser(String name, Output<String> id, @Nullable AuthBackendUserState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("vault:okta/authBackendUser:AuthBackendUser", name, state, makeResourceOptions(options, id));
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
    public static AuthBackendUser get(String name, Output<String> id, @Nullable AuthBackendUserState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new AuthBackendUser(name, id, state, options);
    }
}
