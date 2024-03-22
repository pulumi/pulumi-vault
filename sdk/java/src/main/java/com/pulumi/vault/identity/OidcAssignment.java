// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vault.identity;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.vault.Utilities;
import com.pulumi.vault.identity.OidcAssignmentArgs;
import com.pulumi.vault.identity.inputs.OidcAssignmentState;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Manages OIDC Assignments in a Vault server. See the [Vault documentation](https://www.vaultproject.io/api-docs/secret/identity/oidc-provider#create-or-update-an-assignment)
 * for more information.
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
 * import com.pulumi.vault.identity.Entity;
 * import com.pulumi.vault.identity.EntityArgs;
 * import com.pulumi.vault.identity.OidcAssignment;
 * import com.pulumi.vault.identity.OidcAssignmentArgs;
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
 *             .type(&#34;internal&#34;)
 *             .policies(            
 *                 &#34;dev&#34;,
 *                 &#34;test&#34;)
 *             .build());
 * 
 *         var test = new Entity(&#34;test&#34;, EntityArgs.builder()        
 *             .policies(&#34;test&#34;)
 *             .build());
 * 
 *         var default_ = new OidcAssignment(&#34;default&#34;, OidcAssignmentArgs.builder()        
 *             .entityIds(test.id())
 *             .groupIds(internal.id())
 *             .build());
 * 
 *     }
 * }
 * ```
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Import
 * 
 * OIDC Assignments can be imported using the `name`, e.g.
 * 
 * ```sh
 * $ pulumi import vault:identity/oidcAssignment:OidcAssignment default assignment
 * ```
 * 
 */
@ResourceType(type="vault:identity/oidcAssignment:OidcAssignment")
public class OidcAssignment extends com.pulumi.resources.CustomResource {
    /**
     * A set of Vault entity IDs.
     * 
     */
    @Export(name="entityIds", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> entityIds;

    /**
     * @return A set of Vault entity IDs.
     * 
     */
    public Output<Optional<List<String>>> entityIds() {
        return Codegen.optional(this.entityIds);
    }
    /**
     * A set of Vault group IDs.
     * 
     */
    @Export(name="groupIds", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> groupIds;

    /**
     * @return A set of Vault group IDs.
     * 
     */
    public Output<Optional<List<String>>> groupIds() {
        return Codegen.optional(this.groupIds);
    }
    /**
     * The name of the assignment.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return The name of the assignment.
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
    public OidcAssignment(String name) {
        this(name, OidcAssignmentArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public OidcAssignment(String name, @Nullable OidcAssignmentArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public OidcAssignment(String name, @Nullable OidcAssignmentArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("vault:identity/oidcAssignment:OidcAssignment", name, args == null ? OidcAssignmentArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private OidcAssignment(String name, Output<String> id, @Nullable OidcAssignmentState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("vault:identity/oidcAssignment:OidcAssignment", name, state, makeResourceOptions(options, id));
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
    public static OidcAssignment get(String name, Output<String> id, @Nullable OidcAssignmentState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new OidcAssignment(name, id, state, options);
    }
}
