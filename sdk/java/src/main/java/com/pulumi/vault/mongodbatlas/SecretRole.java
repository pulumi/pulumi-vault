// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vault.mongodbatlas;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.vault.Utilities;
import com.pulumi.vault.mongodbatlas.SecretRoleArgs;
import com.pulumi.vault.mongodbatlas.inputs.SecretRoleState;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
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
 * import com.pulumi.vault.Mount;
 * import com.pulumi.vault.MountArgs;
 * import com.pulumi.vault.mongodbatlas.SecretBackend;
 * import com.pulumi.vault.mongodbatlas.SecretBackendArgs;
 * import com.pulumi.vault.mongodbatlas.SecretRole;
 * import com.pulumi.vault.mongodbatlas.SecretRoleArgs;
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
 *         var mongo = new Mount("mongo", MountArgs.builder()
 *             .path("%s")
 *             .type("mongodbatlas")
 *             .description("MongoDB Atlas secret engine mount")
 *             .build());
 * 
 *         var config = new SecretBackend("config", SecretBackendArgs.builder()
 *             .mount(mongo.path())
 *             .privateKey("privateKey")
 *             .publicKey("publicKey")
 *             .build());
 * 
 *         var role = new SecretRole("role", SecretRoleArgs.builder()
 *             .mount(mongo.path())
 *             .name("tf-test-role")
 *             .organizationId("7cf5a45a9ccf6400e60981b7")
 *             .projectId("5cf5a45a9ccf6400e60981b6")
 *             .roles("ORG_READ_ONLY")
 *             .ipAddresses("192.168.1.5, 192.168.1.6")
 *             .cidrBlocks("192.168.1.3/35")
 *             .projectRoles("GROUP_READ_ONLY")
 *             .ttl("60")
 *             .maxTtl("120")
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
 * The MongoDB Atlas secret role can be imported using the full path to the role
 * of the form: `&lt;mount_path&gt;/roles/&lt;role_name&gt;` e.g.
 * 
 * ```sh
 * $ pulumi import vault:mongodbatlas/secretRole:SecretRole example mongodbatlas/roles/example-role
 * ```
 * 
 */
@ResourceType(type="vault:mongodbatlas/secretRole:SecretRole")
public class SecretRole extends com.pulumi.resources.CustomResource {
    /**
     * Whitelist entry in CIDR notation to be added for the API key.
     * 
     */
    @Export(name="cidrBlocks", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> cidrBlocks;

    /**
     * @return Whitelist entry in CIDR notation to be added for the API key.
     * 
     */
    public Output<Optional<List<String>>> cidrBlocks() {
        return Codegen.optional(this.cidrBlocks);
    }
    /**
     * IP address to be added to the whitelist for the API key.
     * 
     */
    @Export(name="ipAddresses", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> ipAddresses;

    /**
     * @return IP address to be added to the whitelist for the API key.
     * 
     */
    public Output<Optional<List<String>>> ipAddresses() {
        return Codegen.optional(this.ipAddresses);
    }
    /**
     * The maximum allowed lifetime of credentials issued using this role.
     * 
     */
    @Export(name="maxTtl", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> maxTtl;

    /**
     * @return The maximum allowed lifetime of credentials issued using this role.
     * 
     */
    public Output<Optional<String>> maxTtl() {
        return Codegen.optional(this.maxTtl);
    }
    /**
     * Path where the MongoDB Atlas Secrets Engine is mounted.
     * 
     */
    @Export(name="mount", refs={String.class}, tree="[0]")
    private Output<String> mount;

    /**
     * @return Path where the MongoDB Atlas Secrets Engine is mounted.
     * 
     */
    public Output<String> mount() {
        return this.mount;
    }
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
     * Unique identifier for the organization to which the target API Key belongs.
     * Required if `project_id` is not set.
     * 
     */
    @Export(name="organizationId", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> organizationId;

    /**
     * @return Unique identifier for the organization to which the target API Key belongs.
     * Required if `project_id` is not set.
     * 
     */
    public Output<Optional<String>> organizationId() {
        return Codegen.optional(this.organizationId);
    }
    /**
     * Unique identifier for the project to which the target API Key belongs.
     * Required if `organization_id` is not set.
     * 
     */
    @Export(name="projectId", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> projectId;

    /**
     * @return Unique identifier for the project to which the target API Key belongs.
     * Required if `organization_id` is not set.
     * 
     */
    public Output<Optional<String>> projectId() {
        return Codegen.optional(this.projectId);
    }
    /**
     * Roles assigned when an org API key is assigned to a project API key. Possible values are `GROUP_CLUSTER_MANAGER`, `GROUP_DATA_ACCESS_ADMIN`, `GROUP_DATA_ACCESS_READ_ONLY`, `GROUP_DATA_ACCESS_READ_WRITE`, `GROUP_OWNER` and `GROUP_READ_ONLY`.
     * 
     */
    @Export(name="projectRoles", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> projectRoles;

    /**
     * @return Roles assigned when an org API key is assigned to a project API key. Possible values are `GROUP_CLUSTER_MANAGER`, `GROUP_DATA_ACCESS_ADMIN`, `GROUP_DATA_ACCESS_READ_ONLY`, `GROUP_DATA_ACCESS_READ_WRITE`, `GROUP_OWNER` and `GROUP_READ_ONLY`.
     * 
     */
    public Output<Optional<List<String>>> projectRoles() {
        return Codegen.optional(this.projectRoles);
    }
    /**
     * List of roles that the API Key needs to have. Possible values are `ORG_OWNER`, `ORG_MEMBER`, `ORG_GROUP_CREATOR`, `ORG_BILLING_ADMIN` and `ORG_READ_ONLY`.
     * 
     */
    @Export(name="roles", refs={List.class,String.class}, tree="[0,1]")
    private Output<List<String>> roles;

    /**
     * @return List of roles that the API Key needs to have. Possible values are `ORG_OWNER`, `ORG_MEMBER`, `ORG_GROUP_CREATOR`, `ORG_BILLING_ADMIN` and `ORG_READ_ONLY`.
     * 
     */
    public Output<List<String>> roles() {
        return this.roles;
    }
    /**
     * Duration in seconds after which the issued credential should expire.
     * 
     */
    @Export(name="ttl", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> ttl;

    /**
     * @return Duration in seconds after which the issued credential should expire.
     * 
     */
    public Output<Optional<String>> ttl() {
        return Codegen.optional(this.ttl);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public SecretRole(String name) {
        this(name, SecretRoleArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public SecretRole(String name, SecretRoleArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public SecretRole(String name, SecretRoleArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("vault:mongodbatlas/secretRole:SecretRole", name, args == null ? SecretRoleArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private SecretRole(String name, Output<String> id, @Nullable SecretRoleState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("vault:mongodbatlas/secretRole:SecretRole", name, state, makeResourceOptions(options, id));
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
    public static SecretRole get(String name, Output<String> id, @Nullable SecretRoleState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new SecretRole(name, id, state, options);
    }
}
