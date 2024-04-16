// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vault.transform;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.vault.Utilities;
import com.pulumi.vault.transform.RoleArgs;
import com.pulumi.vault.transform.inputs.RoleState;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * This resource supports the &#34;/transform/role/{name}&#34; Vault endpoint.
 * 
 * It creates or updates the role with the given name. If a role with the name does not exist, it will be created.
 * If the role exists, it will be updated with the new attributes.
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
 * import com.pulumi.vault.Mount;
 * import com.pulumi.vault.MountArgs;
 * import com.pulumi.vault.transform.Role;
 * import com.pulumi.vault.transform.RoleArgs;
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
 *         var mountTransform = new Mount(&#34;mountTransform&#34;, MountArgs.builder()        
 *             .path(&#34;transform&#34;)
 *             .type(&#34;transform&#34;)
 *             .build());
 * 
 *         var test = new Role(&#34;test&#34;, RoleArgs.builder()        
 *             .path(mountTransform.path())
 *             .name(&#34;payments&#34;)
 *             .transformations(&#34;ccn-fpe&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 */
@ResourceType(type="vault:transform/role:Role")
public class Role extends com.pulumi.resources.CustomResource {
    /**
     * The name of the role.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return The name of the role.
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
     * Path to where the back-end is mounted within Vault.
     * 
     */
    @Export(name="path", refs={String.class}, tree="[0]")
    private Output<String> path;

    /**
     * @return Path to where the back-end is mounted within Vault.
     * 
     */
    public Output<String> path() {
        return this.path;
    }
    /**
     * A comma separated string or slice of transformations to use.
     * 
     */
    @Export(name="transformations", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> transformations;

    /**
     * @return A comma separated string or slice of transformations to use.
     * 
     */
    public Output<Optional<List<String>>> transformations() {
        return Codegen.optional(this.transformations);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Role(String name) {
        this(name, RoleArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Role(String name, RoleArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Role(String name, RoleArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("vault:transform/role:Role", name, args == null ? RoleArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private Role(String name, Output<String> id, @Nullable RoleState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("vault:transform/role:Role", name, state, makeResourceOptions(options, id));
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
    public static Role get(String name, Output<String> id, @Nullable RoleState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Role(name, id, state, options);
    }
}
