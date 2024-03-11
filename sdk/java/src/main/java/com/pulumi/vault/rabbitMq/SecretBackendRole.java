// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vault.rabbitMq;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.vault.Utilities;
import com.pulumi.vault.rabbitMq.SecretBackendRoleArgs;
import com.pulumi.vault.rabbitMq.inputs.SecretBackendRoleState;
import com.pulumi.vault.rabbitMq.outputs.SecretBackendRoleVhost;
import com.pulumi.vault.rabbitMq.outputs.SecretBackendRoleVhostTopic;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * ## Example Usage
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.vault.rabbitMq.SecretBackend;
 * import com.pulumi.vault.rabbitMq.SecretBackendArgs;
 * import com.pulumi.vault.rabbitMq.SecretBackendRole;
 * import com.pulumi.vault.rabbitMq.SecretBackendRoleArgs;
 * import com.pulumi.vault.rabbitMq.inputs.SecretBackendRoleVhostArgs;
 * import com.pulumi.vault.rabbitMq.inputs.SecretBackendRoleVhostTopicArgs;
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
 *         var rabbitmq = new SecretBackend(&#34;rabbitmq&#34;, SecretBackendArgs.builder()        
 *             .connectionUri(&#34;https://.....&#34;)
 *             .username(&#34;user&#34;)
 *             .password(&#34;password&#34;)
 *             .build());
 * 
 *         var role = new SecretBackendRole(&#34;role&#34;, SecretBackendRoleArgs.builder()        
 *             .backend(rabbitmq.path())
 *             .tags(&#34;tag1,tag2&#34;)
 *             .vhosts(SecretBackendRoleVhostArgs.builder()
 *                 .host(&#34;/&#34;)
 *                 .configure(&#34;&#34;)
 *                 .read(&#34;.*&#34;)
 *                 .write(&#34;&#34;)
 *                 .build())
 *             .vhostTopics(SecretBackendRoleVhostTopicArgs.builder()
 *                 .vhosts(SecretBackendRoleVhostTopicVhostArgs.builder()
 *                     .topic(&#34;amq.topic&#34;)
 *                     .read(&#34;.*&#34;)
 *                     .write(&#34;&#34;)
 *                     .build())
 *                 .host(&#34;/&#34;)
 *                 .build())
 *             .build());
 * 
 *     }
 * }
 * ```
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Import
 * 
 * RabbitMQ secret backend roles can be imported using the `path`, e.g.
 * 
 * ```sh
 * $ pulumi import vault:rabbitMq/secretBackendRole:SecretBackendRole role rabbitmq/roles/deploy
 * ```
 * 
 */
@ResourceType(type="vault:rabbitMq/secretBackendRole:SecretBackendRole")
public class SecretBackendRole extends com.pulumi.resources.CustomResource {
    /**
     * The path the RabbitMQ secret backend is mounted at,
     * with no leading or trailing `/`s.
     * 
     */
    @Export(name="backend", refs={String.class}, tree="[0]")
    private Output<String> backend;

    /**
     * @return The path the RabbitMQ secret backend is mounted at,
     * with no leading or trailing `/`s.
     * 
     */
    public Output<String> backend() {
        return this.backend;
    }
    /**
     * The name to identify this role within the backend.
     * Must be unique within the backend.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return The name to identify this role within the backend.
     * Must be unique within the backend.
     * 
     */
    public Output<String> name() {
        return this.name;
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
     * Specifies a comma-separated RabbitMQ management tags.
     * 
     */
    @Export(name="tags", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> tags;

    /**
     * @return Specifies a comma-separated RabbitMQ management tags.
     * 
     */
    public Output<Optional<String>> tags() {
        return Codegen.optional(this.tags);
    }
    /**
     * Specifies a map of virtual hosts and exchanges to topic permissions. This option requires RabbitMQ 3.7.0 or later.
     * 
     */
    @Export(name="vhostTopics", refs={List.class,SecretBackendRoleVhostTopic.class}, tree="[0,1]")
    private Output</* @Nullable */ List<SecretBackendRoleVhostTopic>> vhostTopics;

    /**
     * @return Specifies a map of virtual hosts and exchanges to topic permissions. This option requires RabbitMQ 3.7.0 or later.
     * 
     */
    public Output<Optional<List<SecretBackendRoleVhostTopic>>> vhostTopics() {
        return Codegen.optional(this.vhostTopics);
    }
    /**
     * Specifies a map of virtual hosts to permissions.
     * 
     */
    @Export(name="vhosts", refs={List.class,SecretBackendRoleVhost.class}, tree="[0,1]")
    private Output</* @Nullable */ List<SecretBackendRoleVhost>> vhosts;

    /**
     * @return Specifies a map of virtual hosts to permissions.
     * 
     */
    public Output<Optional<List<SecretBackendRoleVhost>>> vhosts() {
        return Codegen.optional(this.vhosts);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public SecretBackendRole(String name) {
        this(name, SecretBackendRoleArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public SecretBackendRole(String name, SecretBackendRoleArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public SecretBackendRole(String name, SecretBackendRoleArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("vault:rabbitMq/secretBackendRole:SecretBackendRole", name, args == null ? SecretBackendRoleArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private SecretBackendRole(String name, Output<String> id, @Nullable SecretBackendRoleState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("vault:rabbitMq/secretBackendRole:SecretBackendRole", name, state, makeResourceOptions(options, id));
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
    public static SecretBackendRole get(String name, Output<String> id, @Nullable SecretBackendRoleState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new SecretBackendRole(name, id, state, options);
    }
}
