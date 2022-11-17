// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vault.consul;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.vault.Utilities;
import com.pulumi.vault.consul.SecretBackendRoleArgs;
import com.pulumi.vault.consul.inputs.SecretBackendRoleState;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Manages a Consul secrets role for a Consul secrets engine in Vault. Consul secret backends can then issue Consul tokens.
 * 
 * ## Example Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.vault.consul.SecretBackend;
 * import com.pulumi.vault.consul.SecretBackendArgs;
 * import com.pulumi.vault.consul.SecretBackendRole;
 * import com.pulumi.vault.consul.SecretBackendRoleArgs;
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
 *         var test = new SecretBackend(&#34;test&#34;, SecretBackendArgs.builder()        
 *             .path(&#34;consul&#34;)
 *             .description(&#34;Manages the Consul backend&#34;)
 *             .address(&#34;127.0.0.1:8500&#34;)
 *             .token(&#34;4240861b-ce3d-8530-115a-521ff070dd29&#34;)
 *             .build());
 * 
 *         var example = new SecretBackendRole(&#34;example&#34;, SecretBackendRoleArgs.builder()        
 *             .backend(test.path())
 *             .consulPolicies(&#34;example-policy&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * ## Note About Required Arguments
 * 
 * *At least one* of the four arguments `consul_policies`, `consul_roles`, `service_identities`, or
 * `node_identities` is required for a token. If desired, any combination of the four arguments up-to and
 * including all four, is valid.
 * 
 * ## Import
 * 
 * Consul secret backend roles can be imported using the `backend`, `/roles/`, and the `name` e.g.
 * 
 * ```sh
 *  $ pulumi import vault:consul/secretBackendRole:SecretBackendRole example consul/roles/my-role
 * ```
 * 
 */
@ResourceType(type="vault:consul/secretBackendRole:SecretBackendRole")
public class SecretBackendRole extends com.pulumi.resources.CustomResource {
    /**
     * The unique name of an existing Consul secrets backend mount. Must not begin or end with a `/`. One of `path` or `backend` is required.
     * 
     */
    @Export(name="backend", type=String.class, parameters={})
    private Output</* @Nullable */ String> backend;

    /**
     * @return The unique name of an existing Consul secrets backend mount. Must not begin or end with a `/`. One of `path` or `backend` is required.
     * 
     */
    public Output<Optional<String>> backend() {
        return Codegen.optional(this.backend);
    }
    /**
     * The Consul namespace that the token will be created in.
     * Applicable for Vault 1.10+ and Consul 1.7+&#34;.
     * 
     */
    @Export(name="consulNamespace", type=String.class, parameters={})
    private Output<String> consulNamespace;

    /**
     * @return The Consul namespace that the token will be created in.
     * Applicable for Vault 1.10+ and Consul 1.7+&#34;.
     * 
     */
    public Output<String> consulNamespace() {
        return this.consulNamespace;
    }
    /**
     * &lt;sup&gt;&lt;a href=&#34;#note-about-required-arguments&#34;&gt;SEE NOTE&lt;/a&gt;&lt;/sup&gt; The list of Consul ACL policies to associate with these roles.
     * 
     */
    @Export(name="consulPolicies", type=List.class, parameters={String.class})
    private Output</* @Nullable */ List<String>> consulPolicies;

    /**
     * @return &lt;sup&gt;&lt;a href=&#34;#note-about-required-arguments&#34;&gt;SEE NOTE&lt;/a&gt;&lt;/sup&gt; The list of Consul ACL policies to associate with these roles.
     * 
     */
    public Output<Optional<List<String>>> consulPolicies() {
        return Codegen.optional(this.consulPolicies);
    }
    /**
     * &lt;sup&gt;&lt;a href=&#34;#note-about-required-arguments&#34;&gt;SEE NOTE&lt;/a&gt;&lt;/sup&gt; Set of Consul roles to attach to the token.
     * Applicable for Vault 1.10+ with Consul 1.5+.
     * 
     */
    @Export(name="consulRoles", type=List.class, parameters={String.class})
    private Output</* @Nullable */ List<String>> consulRoles;

    /**
     * @return &lt;sup&gt;&lt;a href=&#34;#note-about-required-arguments&#34;&gt;SEE NOTE&lt;/a&gt;&lt;/sup&gt; Set of Consul roles to attach to the token.
     * Applicable for Vault 1.10+ with Consul 1.5+.
     * 
     */
    public Output<Optional<List<String>>> consulRoles() {
        return Codegen.optional(this.consulRoles);
    }
    /**
     * Indicates that the token should not be replicated globally and instead be local to the current datacenter.
     * 
     */
    @Export(name="local", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> local;

    /**
     * @return Indicates that the token should not be replicated globally and instead be local to the current datacenter.
     * 
     */
    public Output<Optional<Boolean>> local() {
        return Codegen.optional(this.local);
    }
    /**
     * Maximum TTL for leases associated with this role, in seconds.
     * 
     */
    @Export(name="maxTtl", type=Integer.class, parameters={})
    private Output</* @Nullable */ Integer> maxTtl;

    /**
     * @return Maximum TTL for leases associated with this role, in seconds.
     * 
     */
    public Output<Optional<Integer>> maxTtl() {
        return Codegen.optional(this.maxTtl);
    }
    /**
     * The name of the Consul secrets engine role to create.
     * 
     */
    @Export(name="name", type=String.class, parameters={})
    private Output<String> name;

    /**
     * @return The name of the Consul secrets engine role to create.
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
    @Export(name="namespace", type=String.class, parameters={})
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
     * &lt;sup&gt;&lt;a href=&#34;#note-about-required-arguments&#34;&gt;SEE NOTE&lt;/a&gt;&lt;/sup&gt; Set of Consul node
     * identities to attach to the token. Applicable for Vault 1.11+ with Consul 1.8+.
     * 
     */
    @Export(name="nodeIdentities", type=List.class, parameters={String.class})
    private Output</* @Nullable */ List<String>> nodeIdentities;

    /**
     * @return &lt;sup&gt;&lt;a href=&#34;#note-about-required-arguments&#34;&gt;SEE NOTE&lt;/a&gt;&lt;/sup&gt; Set of Consul node
     * identities to attach to the token. Applicable for Vault 1.11+ with Consul 1.8+.
     * 
     */
    public Output<Optional<List<String>>> nodeIdentities() {
        return Codegen.optional(this.nodeIdentities);
    }
    /**
     * The admin partition that the token will be created in.
     * Applicable for Vault 1.10+ and Consul 1.11+&#34;.
     * 
     */
    @Export(name="partition", type=String.class, parameters={})
    private Output<String> partition;

    /**
     * @return The admin partition that the token will be created in.
     * Applicable for Vault 1.10+ and Consul 1.11+&#34;.
     * 
     */
    public Output<String> partition() {
        return this.partition;
    }
    /**
     * The list of Consul ACL policies to associate with these roles.
     * **NOTE:** The new parameter `consul_policies` should be used in favor of this. This parameter,
     * `policies`, remains supported for legacy users, but Vault has deprecated this field.
     * 
     */
    @Export(name="policies", type=List.class, parameters={String.class})
    private Output</* @Nullable */ List<String>> policies;

    /**
     * @return The list of Consul ACL policies to associate with these roles.
     * **NOTE:** The new parameter `consul_policies` should be used in favor of this. This parameter,
     * `policies`, remains supported for legacy users, but Vault has deprecated this field.
     * 
     */
    public Output<Optional<List<String>>> policies() {
        return Codegen.optional(this.policies);
    }
    /**
     * &lt;sup&gt;&lt;a href=&#34;#note-about-required-arguments&#34;&gt;SEE NOTE&lt;/a&gt;&lt;/sup&gt; Set of Consul
     * service identities to attach to the token. Applicable for Vault 1.11+ with Consul 1.5+.
     * 
     */
    @Export(name="serviceIdentities", type=List.class, parameters={String.class})
    private Output</* @Nullable */ List<String>> serviceIdentities;

    /**
     * @return &lt;sup&gt;&lt;a href=&#34;#note-about-required-arguments&#34;&gt;SEE NOTE&lt;/a&gt;&lt;/sup&gt; Set of Consul
     * service identities to attach to the token. Applicable for Vault 1.11+ with Consul 1.5+.
     * 
     */
    public Output<Optional<List<String>>> serviceIdentities() {
        return Codegen.optional(this.serviceIdentities);
    }
    /**
     * Specifies the type of token to create when using this role. Valid values are &#34;client&#34; or &#34;management&#34;.
     * *Deprecated: Consul 1.11 and later removed the legacy ACL system which supported this field.*
     * 
     * @deprecated
     * Consul 1.11 and later removed the legacy ACL system which supported this field.
     * 
     */
    @Deprecated /* Consul 1.11 and later removed the legacy ACL system which supported this field. */
    @Export(name="tokenType", type=String.class, parameters={})
    private Output</* @Nullable */ String> tokenType;

    /**
     * @return Specifies the type of token to create when using this role. Valid values are &#34;client&#34; or &#34;management&#34;.
     * *Deprecated: Consul 1.11 and later removed the legacy ACL system which supported this field.*
     * 
     */
    public Output<Optional<String>> tokenType() {
        return Codegen.optional(this.tokenType);
    }
    /**
     * Specifies the TTL for this role.
     * 
     */
    @Export(name="ttl", type=Integer.class, parameters={})
    private Output</* @Nullable */ Integer> ttl;

    /**
     * @return Specifies the TTL for this role.
     * 
     */
    public Output<Optional<Integer>> ttl() {
        return Codegen.optional(this.ttl);
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
    public SecretBackendRole(String name, @Nullable SecretBackendRoleArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public SecretBackendRole(String name, @Nullable SecretBackendRoleArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("vault:consul/secretBackendRole:SecretBackendRole", name, args == null ? SecretBackendRoleArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private SecretBackendRole(String name, Output<String> id, @Nullable SecretBackendRoleState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("vault:consul/secretBackendRole:SecretBackendRole", name, state, makeResourceOptions(options, id));
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