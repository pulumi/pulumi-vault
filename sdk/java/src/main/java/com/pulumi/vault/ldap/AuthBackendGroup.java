// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vault.ldap;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.vault.Utilities;
import com.pulumi.vault.ldap.AuthBackendGroupArgs;
import com.pulumi.vault.ldap.inputs.AuthBackendGroupState;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides a resource to create a group in an [LDAP auth backend within Vault](https://www.vaultproject.io/docs/auth/ldap.html).
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
 * import com.pulumi.vault.ldap.AuthBackend;
 * import com.pulumi.vault.ldap.AuthBackendArgs;
 * import com.pulumi.vault.ldap.AuthBackendGroup;
 * import com.pulumi.vault.ldap.AuthBackendGroupArgs;
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
 *         var ldap = new AuthBackend("ldap", AuthBackendArgs.builder()
 *             .path("ldap")
 *             .url("ldaps://dc-01.example.org")
 *             .userdn("OU=Users,OU=Accounts,DC=example,DC=org")
 *             .userattr("sAMAccountName")
 *             .upndomain("EXAMPLE.ORG")
 *             .discoverdn(false)
 *             .groupdn("OU=Groups,DC=example,DC=org")
 *             .groupfilter("(&(objectClass=group)(member:1.2.840.113556.1.4.1941:={{.UserDN}}))")
 *             .build());
 * 
 *         var group = new AuthBackendGroup("group", AuthBackendGroupArgs.builder()
 *             .groupname("dba")
 *             .policies("dba")
 *             .backend(ldap.path())
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
 * LDAP authentication backend groups can be imported using the `path`, e.g.
 * 
 * ```sh
 * $ pulumi import vault:ldap/authBackendGroup:AuthBackendGroup foo auth/ldap/groups/foo
 * ```
 * 
 */
@ResourceType(type="vault:ldap/authBackendGroup:AuthBackendGroup")
public class AuthBackendGroup extends com.pulumi.resources.CustomResource {
    /**
     * Path to the authentication backend
     * 
     * For more details on the usage of each argument consult the [Vault LDAP API documentation](https://www.vaultproject.io/api-docs/auth/ldap).
     * 
     */
    @Export(name="backend", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> backend;

    /**
     * @return Path to the authentication backend
     * 
     * For more details on the usage of each argument consult the [Vault LDAP API documentation](https://www.vaultproject.io/api-docs/auth/ldap).
     * 
     */
    public Output<Optional<String>> backend() {
        return Codegen.optional(this.backend);
    }
    /**
     * The LDAP groupname
     * 
     */
    @Export(name="groupname", refs={String.class}, tree="[0]")
    private Output<String> groupname;

    /**
     * @return The LDAP groupname
     * 
     */
    public Output<String> groupname() {
        return this.groupname;
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
     * Policies which should be granted to members of the group
     * 
     */
    @Export(name="policies", refs={List.class,String.class}, tree="[0,1]")
    private Output<List<String>> policies;

    /**
     * @return Policies which should be granted to members of the group
     * 
     */
    public Output<List<String>> policies() {
        return this.policies;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public AuthBackendGroup(String name) {
        this(name, AuthBackendGroupArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public AuthBackendGroup(String name, AuthBackendGroupArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public AuthBackendGroup(String name, AuthBackendGroupArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("vault:ldap/authBackendGroup:AuthBackendGroup", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()));
    }

    private AuthBackendGroup(String name, Output<String> id, @Nullable AuthBackendGroupState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("vault:ldap/authBackendGroup:AuthBackendGroup", name, state, makeResourceOptions(options, id));
    }

    private static AuthBackendGroupArgs makeArgs(AuthBackendGroupArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? AuthBackendGroupArgs.Empty : args;
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
    public static AuthBackendGroup get(String name, Output<String> id, @Nullable AuthBackendGroupState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new AuthBackendGroup(name, id, state, options);
    }
}
