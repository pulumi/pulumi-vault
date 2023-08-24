// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vault.azure;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.vault.Utilities;
import com.pulumi.vault.azure.BackendRoleArgs;
import com.pulumi.vault.azure.inputs.BackendRoleState;
import com.pulumi.vault.azure.outputs.BackendRoleAzureGroup;
import com.pulumi.vault.azure.outputs.BackendRoleAzureRole;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * ## Example Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.vault.azure.Backend;
 * import com.pulumi.vault.azure.BackendArgs;
 * import com.pulumi.vault.azure.BackendRole;
 * import com.pulumi.vault.azure.BackendRoleArgs;
 * import com.pulumi.vault.azure.inputs.BackendRoleAzureRoleArgs;
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
 *         var azure = new Backend(&#34;azure&#34;, BackendArgs.builder()        
 *             .subscriptionId(var_.subscription_id())
 *             .tenantId(var_.tenant_id())
 *             .clientSecret(var_.client_secret())
 *             .clientId(var_.client_id())
 *             .build());
 * 
 *         var generatedRole = new BackendRole(&#34;generatedRole&#34;, BackendRoleArgs.builder()        
 *             .backend(azure.path())
 *             .role(&#34;generated_role&#34;)
 *             .ttl(300)
 *             .maxTtl(600)
 *             .azureRoles(BackendRoleAzureRoleArgs.builder()
 *                 .roleName(&#34;Reader&#34;)
 *                 .scope(String.format(&#34;/subscriptions/%s/resourceGroups/azure-vault-group&#34;, var_.subscription_id()))
 *                 .build())
 *             .build());
 * 
 *         var existingObjectId = new BackendRole(&#34;existingObjectId&#34;, BackendRoleArgs.builder()        
 *             .backend(azure.path())
 *             .role(&#34;existing_object_id&#34;)
 *             .applicationObjectId(&#34;11111111-2222-3333-4444-44444444444&#34;)
 *             .ttl(300)
 *             .maxTtl(600)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 */
@ResourceType(type="vault:azure/backendRole:BackendRole")
public class BackendRole extends com.pulumi.resources.CustomResource {
    /**
     * Application Object ID for an existing service principal that will
     * be used instead of creating dynamic service principals. If present, `azure_roles` will be ignored.
     * 
     */
    @Export(name="applicationObjectId", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> applicationObjectId;

    /**
     * @return Application Object ID for an existing service principal that will
     * be used instead of creating dynamic service principals. If present, `azure_roles` will be ignored.
     * 
     */
    public Output<Optional<String>> applicationObjectId() {
        return Codegen.optional(this.applicationObjectId);
    }
    /**
     * List of Azure groups to be assigned to the generated service principal.
     * 
     */
    @Export(name="azureGroups", refs={List.class,BackendRoleAzureGroup.class}, tree="[0,1]")
    private Output</* @Nullable */ List<BackendRoleAzureGroup>> azureGroups;

    /**
     * @return List of Azure groups to be assigned to the generated service principal.
     * 
     */
    public Output<Optional<List<BackendRoleAzureGroup>>> azureGroups() {
        return Codegen.optional(this.azureGroups);
    }
    /**
     * List of Azure roles to be assigned to the generated service principal.
     * 
     */
    @Export(name="azureRoles", refs={List.class,BackendRoleAzureRole.class}, tree="[0,1]")
    private Output</* @Nullable */ List<BackendRoleAzureRole>> azureRoles;

    /**
     * @return List of Azure roles to be assigned to the generated service principal.
     * 
     */
    public Output<Optional<List<BackendRoleAzureRole>>> azureRoles() {
        return Codegen.optional(this.azureRoles);
    }
    /**
     * Path to the mounted Azure auth backend
     * 
     */
    @Export(name="backend", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> backend;

    /**
     * @return Path to the mounted Azure auth backend
     * 
     */
    public Output<Optional<String>> backend() {
        return Codegen.optional(this.backend);
    }
    /**
     * Human-friendly description of the mount for the backend.
     * 
     */
    @Export(name="description", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> description;

    /**
     * @return Human-friendly description of the mount for the backend.
     * 
     */
    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    /**
     * Specifies the maximum TTL for service principals generated using this role. Accepts time
     * suffixed strings (&#34;1h&#34;) or an integer number of seconds. Defaults to the system/engine max TTL time.
     * 
     */
    @Export(name="maxTtl", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> maxTtl;

    /**
     * @return Specifies the maximum TTL for service principals generated using this role. Accepts time
     * suffixed strings (&#34;1h&#34;) or an integer number of seconds. Defaults to the system/engine max TTL time.
     * 
     */
    public Output<Optional<String>> maxTtl() {
        return Codegen.optional(this.maxTtl);
    }
    /**
     * The namespace to provision the resource in.
     * The value should not contain leading or trailing forward slashes.
     * The `namespace` is always relative to the provider&#39;s configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
     * *Available only for Vault Enterprise*.
     * 
     */
    @Export(name="namespace", refs={String.class}, tree="[0]")
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
     * Name of the Azure role
     * 
     */
    @Export(name="role", refs={String.class}, tree="[0]")
    private Output<String> role;

    /**
     * @return Name of the Azure role
     * 
     */
    public Output<String> role() {
        return this.role;
    }
    /**
     * Specifies the default TTL for service principals generated using this role.
     * Accepts time suffixed strings (&#34;1h&#34;) or an integer number of seconds. Defaults to the system/engine default TTL time.
     * 
     */
    @Export(name="ttl", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> ttl;

    /**
     * @return Specifies the default TTL for service principals generated using this role.
     * Accepts time suffixed strings (&#34;1h&#34;) or an integer number of seconds. Defaults to the system/engine default TTL time.
     * 
     */
    public Output<Optional<String>> ttl() {
        return Codegen.optional(this.ttl);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public BackendRole(String name) {
        this(name, BackendRoleArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public BackendRole(String name, BackendRoleArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public BackendRole(String name, BackendRoleArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("vault:azure/backendRole:BackendRole", name, args == null ? BackendRoleArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private BackendRole(String name, Output<String> id, @Nullable BackendRoleState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("vault:azure/backendRole:BackendRole", name, state, makeResourceOptions(options, id));
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
    public static BackendRole get(String name, Output<String> id, @Nullable BackendRoleState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new BackendRole(name, id, state, options);
    }
}
