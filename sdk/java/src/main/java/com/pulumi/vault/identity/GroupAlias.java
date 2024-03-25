// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vault.identity;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.vault.Utilities;
import com.pulumi.vault.identity.GroupAliasArgs;
import com.pulumi.vault.identity.inputs.GroupAliasState;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Creates an Identity Group Alias for Vault. The [Identity secrets engine](https://www.vaultproject.io/docs/secrets/identity/index.html) is the identity management solution for Vault.
 * 
 * Group aliases allows entity membership in external groups to be managed semi-automatically. External group serves as a mapping to a group that is outside of the identity store. External groups can have one (and only one) alias. This alias should map to a notion of group that is outside of the identity store. For example, groups in LDAP, and teams in GitHub. A username in LDAP, belonging to a group in LDAP, can get its entity ID added as a member of a group in Vault automatically during logins and token renewals. This works only if the group in Vault is an external group and has an alias that maps to the group in LDAP. If the user is removed from the group in LDAP, that change gets reflected in Vault only upon the subsequent login or renewal operation.
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
 * import com.pulumi.vault.identity.Group;
 * import com.pulumi.vault.identity.GroupArgs;
 * import com.pulumi.vault.AuthBackend;
 * import com.pulumi.vault.AuthBackendArgs;
 * import com.pulumi.vault.identity.GroupAlias;
 * import com.pulumi.vault.identity.GroupAliasArgs;
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
 *         var group = new Group(&#34;group&#34;, GroupArgs.builder()        
 *             .type(&#34;external&#34;)
 *             .policies(&#34;test&#34;)
 *             .build());
 * 
 *         var github = new AuthBackend(&#34;github&#34;, AuthBackendArgs.builder()        
 *             .type(&#34;github&#34;)
 *             .path(&#34;github&#34;)
 *             .build());
 * 
 *         var group_alias = new GroupAlias(&#34;group-alias&#34;, GroupAliasArgs.builder()        
 *             .name(&#34;Github_Team_Slug&#34;)
 *             .mountAccessor(github.accessor())
 *             .canonicalId(group.id())
 *             .build());
 * 
 *     }
 * }
 * ```
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Import
 * 
 * The group alias can be imported with the group alias `id`, for example:
 * 
 * ```sh
 * $ pulumi import vault:identity/groupAlias:GroupAlias group-alias id
 * ```
 * 
 * Group aliases can also be imported using the UUID of the alias record, e.g.
 * 
 * ```sh
 * $ pulumi import vault:identity/groupAlias:GroupAlias alias_name 63104e20-88e4-11eb-8d04-cf7ac9d60157
 * ```
 * 
 */
@ResourceType(type="vault:identity/groupAlias:GroupAlias")
public class GroupAlias extends com.pulumi.resources.CustomResource {
    /**
     * ID of the group to which this is an alias.
     * 
     */
    @Export(name="canonicalId", refs={String.class}, tree="[0]")
    private Output<String> canonicalId;

    /**
     * @return ID of the group to which this is an alias.
     * 
     */
    public Output<String> canonicalId() {
        return this.canonicalId;
    }
    /**
     * Mount accessor of the authentication backend to which this alias belongs to.
     * 
     */
    @Export(name="mountAccessor", refs={String.class}, tree="[0]")
    private Output<String> mountAccessor;

    /**
     * @return Mount accessor of the authentication backend to which this alias belongs to.
     * 
     */
    public Output<String> mountAccessor() {
        return this.mountAccessor;
    }
    /**
     * Name of the group alias to create.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return Name of the group alias to create.
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
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public GroupAlias(String name) {
        this(name, GroupAliasArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public GroupAlias(String name, GroupAliasArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public GroupAlias(String name, GroupAliasArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("vault:identity/groupAlias:GroupAlias", name, args == null ? GroupAliasArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private GroupAlias(String name, Output<String> id, @Nullable GroupAliasState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("vault:identity/groupAlias:GroupAlias", name, state, makeResourceOptions(options, id));
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
    public static GroupAlias get(String name, Output<String> id, @Nullable GroupAliasState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new GroupAlias(name, id, state, options);
    }
}
