// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vault.identity;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.vault.Utilities;
import com.pulumi.vault.identity.EntityPoliciesArgs;
import com.pulumi.vault.identity.inputs.EntityPoliciesState;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Manages policies for an Identity Entity for Vault. The [Identity secrets engine](https://www.vaultproject.io/docs/secrets/identity/index.html) is the identity management solution for Vault.
 * 
 * ## Example Usage
 * 
 * ### Exclusive Policies
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.vault.identity.Entity;
 * import com.pulumi.vault.identity.EntityArgs;
 * import com.pulumi.vault.identity.EntityPolicies;
 * import com.pulumi.vault.identity.EntityPoliciesArgs;
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
 *         var entity = new Entity("entity", EntityArgs.builder()        
 *             .name("entity")
 *             .externalPolicies(true)
 *             .build());
 * 
 *         var policies = new EntityPolicies("policies", EntityPoliciesArgs.builder()        
 *             .policies(            
 *                 "default",
 *                 "test")
 *             .exclusive(true)
 *             .entityId(entity.id())
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ### Non-exclusive Policies
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.vault.identity.Entity;
 * import com.pulumi.vault.identity.EntityArgs;
 * import com.pulumi.vault.identity.EntityPolicies;
 * import com.pulumi.vault.identity.EntityPoliciesArgs;
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
 *         var entity = new Entity("entity", EntityArgs.builder()        
 *             .name("entity")
 *             .externalPolicies(true)
 *             .build());
 * 
 *         var default_ = new EntityPolicies("default", EntityPoliciesArgs.builder()        
 *             .policies(            
 *                 "default",
 *                 "test")
 *             .exclusive(false)
 *             .entityId(entity.id())
 *             .build());
 * 
 *         var others = new EntityPolicies("others", EntityPoliciesArgs.builder()        
 *             .policies("others")
 *             .exclusive(false)
 *             .entityId(entity.id())
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 */
@ResourceType(type="vault:identity/entityPolicies:EntityPolicies")
public class EntityPolicies extends com.pulumi.resources.CustomResource {
    /**
     * Entity ID to assign policies to.
     * 
     */
    @Export(name="entityId", refs={String.class}, tree="[0]")
    private Output<String> entityId;

    /**
     * @return Entity ID to assign policies to.
     * 
     */
    public Output<String> entityId() {
        return this.entityId;
    }
    /**
     * The name of the entity that are assigned the policies.
     * 
     */
    @Export(name="entityName", refs={String.class}, tree="[0]")
    private Output<String> entityName;

    /**
     * @return The name of the entity that are assigned the policies.
     * 
     */
    public Output<String> entityName() {
        return this.entityName;
    }
    /**
     * Defaults to `true`.
     * 
     * If `true`, this resource will take exclusive control of the policies assigned to the entity and will set it equal to what is specified in the resource.
     * 
     * If set to `false`, this resource will simply ensure that the policies specified in the resource are present in the entity. When destroying the resource, the resource will ensure that the policies specified in the resource are removed.
     * 
     */
    @Export(name="exclusive", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> exclusive;

    /**
     * @return Defaults to `true`.
     * 
     * If `true`, this resource will take exclusive control of the policies assigned to the entity and will set it equal to what is specified in the resource.
     * 
     * If set to `false`, this resource will simply ensure that the policies specified in the resource are present in the entity. When destroying the resource, the resource will ensure that the policies specified in the resource are removed.
     * 
     */
    public Output<Optional<Boolean>> exclusive() {
        return Codegen.optional(this.exclusive);
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
     * List of policies to assign to the entity
     * 
     */
    @Export(name="policies", refs={List.class,String.class}, tree="[0,1]")
    private Output<List<String>> policies;

    /**
     * @return List of policies to assign to the entity
     * 
     */
    public Output<List<String>> policies() {
        return this.policies;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public EntityPolicies(String name) {
        this(name, EntityPoliciesArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public EntityPolicies(String name, EntityPoliciesArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public EntityPolicies(String name, EntityPoliciesArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("vault:identity/entityPolicies:EntityPolicies", name, args == null ? EntityPoliciesArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private EntityPolicies(String name, Output<String> id, @Nullable EntityPoliciesState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("vault:identity/entityPolicies:EntityPolicies", name, state, makeResourceOptions(options, id));
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
    public static EntityPolicies get(String name, Output<String> id, @Nullable EntityPoliciesState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new EntityPolicies(name, id, state, options);
    }
}
