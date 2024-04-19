// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vault.identity;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.vault.Utilities;
import com.pulumi.vault.identity.GroupMemberEntityIdsArgs;
import com.pulumi.vault.identity.inputs.GroupMemberEntityIdsState;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Manages member entities for an Identity Group for Vault. The [Identity secrets engine](https://www.vaultproject.io/docs/secrets/identity/index.html) is the identity management solution for Vault.
 * 
 * ## Example Usage
 * 
 * ### Exclusive Member Entities
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
 * import com.pulumi.vault.identity.Entity;
 * import com.pulumi.vault.identity.EntityArgs;
 * import com.pulumi.vault.identity.GroupMemberEntityIds;
 * import com.pulumi.vault.identity.GroupMemberEntityIdsArgs;
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
 *         var internal = new Group(&#34;internal&#34;, GroupArgs.builder()        
 *             .name(&#34;internal&#34;)
 *             .type(&#34;internal&#34;)
 *             .externalMemberEntityIds(true)
 *             .metadata(Map.of(&#34;version&#34;, &#34;2&#34;))
 *             .build());
 * 
 *         var user = new Entity(&#34;user&#34;, EntityArgs.builder()        
 *             .name(&#34;user&#34;)
 *             .build());
 * 
 *         var members = new GroupMemberEntityIds(&#34;members&#34;, GroupMemberEntityIdsArgs.builder()        
 *             .exclusive(true)
 *             .memberEntityIds(user.id())
 *             .groupId(internal.id())
 *             .build());
 * 
 *     }
 * }
 * ```
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ### Non-exclusive Member Entities
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
 * import com.pulumi.vault.identity.Entity;
 * import com.pulumi.vault.identity.EntityArgs;
 * import com.pulumi.vault.identity.GroupMemberEntityIds;
 * import com.pulumi.vault.identity.GroupMemberEntityIdsArgs;
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
 *         var internal = new Group(&#34;internal&#34;, GroupArgs.builder()        
 *             .name(&#34;internal&#34;)
 *             .type(&#34;internal&#34;)
 *             .externalMemberEntityIds(true)
 *             .metadata(Map.of(&#34;version&#34;, &#34;2&#34;))
 *             .build());
 * 
 *         var testUser = new Entity(&#34;testUser&#34;, EntityArgs.builder()        
 *             .name(&#34;test&#34;)
 *             .build());
 * 
 *         var secondTestUser = new Entity(&#34;secondTestUser&#34;, EntityArgs.builder()        
 *             .name(&#34;second_test&#34;)
 *             .build());
 * 
 *         var devUser = new Entity(&#34;devUser&#34;, EntityArgs.builder()        
 *             .name(&#34;dev&#34;)
 *             .build());
 * 
 *         var test = new GroupMemberEntityIds(&#34;test&#34;, GroupMemberEntityIdsArgs.builder()        
 *             .memberEntityIds(            
 *                 testUser.id(),
 *                 secondTestUser.id())
 *             .exclusive(false)
 *             .groupId(internal.id())
 *             .build());
 * 
 *         var others = new GroupMemberEntityIds(&#34;others&#34;, GroupMemberEntityIdsArgs.builder()        
 *             .memberEntityIds(devUser.id())
 *             .exclusive(false)
 *             .groupId(internal.id())
 *             .build());
 * 
 *     }
 * }
 * ```
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 */
@ResourceType(type="vault:identity/groupMemberEntityIds:GroupMemberEntityIds")
public class GroupMemberEntityIds extends com.pulumi.resources.CustomResource {
    /**
     * Defaults to `true`.
     * 
     * If `true`, this resource will take exclusive control of the member entities that belong to the group and will set it equal to what is specified in the resource.
     * 
     * If set to `false`, this resource will simply ensure that the member entities specified in the resource are present in the group. When destroying the resource, the resource will ensure that the member entities specified in the resource are removed.
     * 
     */
    @Export(name="exclusive", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> exclusive;

    /**
     * @return Defaults to `true`.
     * 
     * If `true`, this resource will take exclusive control of the member entities that belong to the group and will set it equal to what is specified in the resource.
     * 
     * If set to `false`, this resource will simply ensure that the member entities specified in the resource are present in the group. When destroying the resource, the resource will ensure that the member entities specified in the resource are removed.
     * 
     */
    public Output<Optional<Boolean>> exclusive() {
        return Codegen.optional(this.exclusive);
    }
    /**
     * Group ID to assign member entities to.
     * 
     */
    @Export(name="groupId", refs={String.class}, tree="[0]")
    private Output<String> groupId;

    /**
     * @return Group ID to assign member entities to.
     * 
     */
    public Output<String> groupId() {
        return this.groupId;
    }
    /**
     * List of member entities that belong to the group
     * 
     */
    @Export(name="memberEntityIds", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> memberEntityIds;

    /**
     * @return List of member entities that belong to the group
     * 
     */
    public Output<Optional<List<String>>> memberEntityIds() {
        return Codegen.optional(this.memberEntityIds);
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
    public GroupMemberEntityIds(String name) {
        this(name, GroupMemberEntityIdsArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public GroupMemberEntityIds(String name, GroupMemberEntityIdsArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public GroupMemberEntityIds(String name, GroupMemberEntityIdsArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("vault:identity/groupMemberEntityIds:GroupMemberEntityIds", name, args == null ? GroupMemberEntityIdsArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private GroupMemberEntityIds(String name, Output<String> id, @Nullable GroupMemberEntityIdsState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("vault:identity/groupMemberEntityIds:GroupMemberEntityIds", name, state, makeResourceOptions(options, id));
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
    public static GroupMemberEntityIds get(String name, Output<String> id, @Nullable GroupMemberEntityIdsState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new GroupMemberEntityIds(name, id, state, options);
    }
}
