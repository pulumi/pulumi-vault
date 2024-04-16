// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vault.identity;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.vault.Utilities;
import com.pulumi.vault.identity.GroupPoliciesArgs;
import com.pulumi.vault.identity.inputs.GroupPoliciesState;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Manages policies for an Identity Group for Vault. The [Identity secrets engine](https://www.vaultproject.io/docs/secrets/identity/index.html) is the identity management solution for Vault.
 * 
 * ## Example Usage
 * 
 * ### Exclusive Policies
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
 * import com.pulumi.vault.identity.GroupPolicies;
 * import com.pulumi.vault.identity.GroupPoliciesArgs;
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
 *             .externalPolicies(true)
 *             .metadata(Map.of(&#34;version&#34;, &#34;2&#34;))
 *             .build());
 * 
 *         var policies = new GroupPolicies(&#34;policies&#34;, GroupPoliciesArgs.builder()        
 *             .policies(            
 *                 &#34;default&#34;,
 *                 &#34;test&#34;)
 *             .exclusive(true)
 *             .groupId(internal.id())
 *             .build());
 * 
 *     }
 * }
 * ```
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ### Non-exclusive Policies
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
 * import com.pulumi.vault.identity.GroupPolicies;
 * import com.pulumi.vault.identity.GroupPoliciesArgs;
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
 *             .externalPolicies(true)
 *             .metadata(Map.of(&#34;version&#34;, &#34;2&#34;))
 *             .build());
 * 
 *         var default_ = new GroupPolicies(&#34;default&#34;, GroupPoliciesArgs.builder()        
 *             .policies(            
 *                 &#34;default&#34;,
 *                 &#34;test&#34;)
 *             .exclusive(false)
 *             .groupId(internal.id())
 *             .build());
 * 
 *         var others = new GroupPolicies(&#34;others&#34;, GroupPoliciesArgs.builder()        
 *             .policies(&#34;others&#34;)
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
@ResourceType(type="vault:identity/groupPolicies:GroupPolicies")
public class GroupPolicies extends com.pulumi.resources.CustomResource {
    /**
     * Defaults to `true`.
     * 
     * If `true`, this resource will take exclusive control of the policies assigned to the group and will set it equal to what is specified in the resource.
     * 
     * If set to `false`, this resource will simply ensure that the policies specified in the resource are present in the group. When destroying the resource, the resource will ensure that the policies specified in the resource are removed.
     * 
     */
    @Export(name="exclusive", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> exclusive;

    /**
     * @return Defaults to `true`.
     * 
     * If `true`, this resource will take exclusive control of the policies assigned to the group and will set it equal to what is specified in the resource.
     * 
     * If set to `false`, this resource will simply ensure that the policies specified in the resource are present in the group. When destroying the resource, the resource will ensure that the policies specified in the resource are removed.
     * 
     */
    public Output<Optional<Boolean>> exclusive() {
        return Codegen.optional(this.exclusive);
    }
    /**
     * Group ID to assign policies to.
     * 
     */
    @Export(name="groupId", refs={String.class}, tree="[0]")
    private Output<String> groupId;

    /**
     * @return Group ID to assign policies to.
     * 
     */
    public Output<String> groupId() {
        return this.groupId;
    }
    /**
     * The name of the group that are assigned the policies.
     * 
     */
    @Export(name="groupName", refs={String.class}, tree="[0]")
    private Output<String> groupName;

    /**
     * @return The name of the group that are assigned the policies.
     * 
     */
    public Output<String> groupName() {
        return this.groupName;
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
     * List of policies to assign to the group
     * 
     */
    @Export(name="policies", refs={List.class,String.class}, tree="[0,1]")
    private Output<List<String>> policies;

    /**
     * @return List of policies to assign to the group
     * 
     */
    public Output<List<String>> policies() {
        return this.policies;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public GroupPolicies(String name) {
        this(name, GroupPoliciesArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public GroupPolicies(String name, GroupPoliciesArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public GroupPolicies(String name, GroupPoliciesArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("vault:identity/groupPolicies:GroupPolicies", name, args == null ? GroupPoliciesArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private GroupPolicies(String name, Output<String> id, @Nullable GroupPoliciesState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("vault:identity/groupPolicies:GroupPolicies", name, state, makeResourceOptions(options, id));
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
    public static GroupPolicies get(String name, Output<String> id, @Nullable GroupPoliciesState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new GroupPolicies(name, id, state, options);
    }
}
