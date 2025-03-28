// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vault.ssh;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.vault.Utilities;
import com.pulumi.vault.ssh.SecretBackendRoleArgs;
import com.pulumi.vault.ssh.inputs.SecretBackendRoleState;
import com.pulumi.vault.ssh.outputs.SecretBackendRoleAllowedUserKeyConfig;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides a resource to manage roles in an SSH secret backend
 * [SSH secret backend within Vault](https://www.vaultproject.io/docs/secrets/ssh/index.html).
 * 
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
 * import com.pulumi.vault.ssh.SecretBackendRole;
 * import com.pulumi.vault.ssh.SecretBackendRoleArgs;
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
 *         var example = new Mount("example", MountArgs.builder()
 *             .type("ssh")
 *             .build());
 * 
 *         var foo = new SecretBackendRole("foo", SecretBackendRoleArgs.builder()
 *             .name("my-role")
 *             .backend(example.path())
 *             .keyType("ca")
 *             .allowUserCertificates(true)
 *             .build());
 * 
 *         var bar = new SecretBackendRole("bar", SecretBackendRoleArgs.builder()
 *             .name("otp-role")
 *             .backend(example.path())
 *             .keyType("otp")
 *             .defaultUser("default")
 *             .allowedUsers("default,baz")
 *             .cidrList("0.0.0.0/0")
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
 * SSH secret backend roles can be imported using the `path`, e.g.
 * 
 * ```sh
 * $ pulumi import vault:ssh/secretBackendRole:SecretBackendRole foo ssh/roles/my-role
 * ```
 * 
 */
@ResourceType(type="vault:ssh/secretBackendRole:SecretBackendRole")
public class SecretBackendRole extends com.pulumi.resources.CustomResource {
    /**
     * When supplied, this value specifies a signing algorithm for the key. Possible values: ssh-rsa, rsa-sha2-256, rsa-sha2-512.
     * 
     */
    @Export(name="algorithmSigner", refs={String.class}, tree="[0]")
    private Output<String> algorithmSigner;

    /**
     * @return When supplied, this value specifies a signing algorithm for the key. Possible values: ssh-rsa, rsa-sha2-256, rsa-sha2-512.
     * 
     */
    public Output<String> algorithmSigner() {
        return this.algorithmSigner;
    }
    /**
     * Specifies if host certificates that are requested are allowed to use the base domains listed in `allowed_domains`.
     * 
     */
    @Export(name="allowBareDomains", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> allowBareDomains;

    /**
     * @return Specifies if host certificates that are requested are allowed to use the base domains listed in `allowed_domains`.
     * 
     */
    public Output<Optional<Boolean>> allowBareDomains() {
        return Codegen.optional(this.allowBareDomains);
    }
    /**
     * Allow signing certificates with no
     * valid principals (e.g. any valid principal). For backwards compatibility
     * only. The default of false is highly recommended.
     * 
     */
    @Export(name="allowEmptyPrincipals", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> allowEmptyPrincipals;

    /**
     * @return Allow signing certificates with no
     * valid principals (e.g. any valid principal). For backwards compatibility
     * only. The default of false is highly recommended.
     * 
     */
    public Output<Optional<Boolean>> allowEmptyPrincipals() {
        return Codegen.optional(this.allowEmptyPrincipals);
    }
    /**
     * Specifies if certificates are allowed to be signed for use as a &#39;host&#39;.
     * 
     */
    @Export(name="allowHostCertificates", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> allowHostCertificates;

    /**
     * @return Specifies if certificates are allowed to be signed for use as a &#39;host&#39;.
     * 
     */
    public Output<Optional<Boolean>> allowHostCertificates() {
        return Codegen.optional(this.allowHostCertificates);
    }
    /**
     * Specifies if host certificates that are requested are allowed to be subdomains of those listed in `allowed_domains`.
     * 
     */
    @Export(name="allowSubdomains", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> allowSubdomains;

    /**
     * @return Specifies if host certificates that are requested are allowed to be subdomains of those listed in `allowed_domains`.
     * 
     */
    public Output<Optional<Boolean>> allowSubdomains() {
        return Codegen.optional(this.allowSubdomains);
    }
    /**
     * Specifies if certificates are allowed to be signed for use as a &#39;user&#39;.
     * 
     */
    @Export(name="allowUserCertificates", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> allowUserCertificates;

    /**
     * @return Specifies if certificates are allowed to be signed for use as a &#39;user&#39;.
     * 
     */
    public Output<Optional<Boolean>> allowUserCertificates() {
        return Codegen.optional(this.allowUserCertificates);
    }
    /**
     * Specifies if users can override the key ID for a signed certificate with the `key_id` field.
     * 
     */
    @Export(name="allowUserKeyIds", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> allowUserKeyIds;

    /**
     * @return Specifies if users can override the key ID for a signed certificate with the `key_id` field.
     * 
     */
    public Output<Optional<Boolean>> allowUserKeyIds() {
        return Codegen.optional(this.allowUserKeyIds);
    }
    /**
     * Specifies a comma-separated list of critical options that certificates can have when signed.
     * 
     */
    @Export(name="allowedCriticalOptions", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> allowedCriticalOptions;

    /**
     * @return Specifies a comma-separated list of critical options that certificates can have when signed.
     * 
     */
    public Output<Optional<String>> allowedCriticalOptions() {
        return Codegen.optional(this.allowedCriticalOptions);
    }
    /**
     * The list of domains for which a client can request a host certificate.
     * 
     */
    @Export(name="allowedDomains", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> allowedDomains;

    /**
     * @return The list of domains for which a client can request a host certificate.
     * 
     */
    public Output<Optional<String>> allowedDomains() {
        return Codegen.optional(this.allowedDomains);
    }
    /**
     * Specifies if `allowed_domains` can be declared using
     * identity template policies. Non-templated domains are also permitted.
     * 
     */
    @Export(name="allowedDomainsTemplate", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> allowedDomainsTemplate;

    /**
     * @return Specifies if `allowed_domains` can be declared using
     * identity template policies. Non-templated domains are also permitted.
     * 
     */
    public Output<Boolean> allowedDomainsTemplate() {
        return this.allowedDomainsTemplate;
    }
    /**
     * Specifies a comma-separated list of extensions that certificates can have when signed.
     * 
     */
    @Export(name="allowedExtensions", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> allowedExtensions;

    /**
     * @return Specifies a comma-separated list of extensions that certificates can have when signed.
     * 
     */
    public Output<Optional<String>> allowedExtensions() {
        return Codegen.optional(this.allowedExtensions);
    }
    /**
     * Set of configuration blocks to define allowed\
     * user key configuration, like key type and their lengths. Can be specified multiple times.
     * *See Configuration-Options for more info*
     * 
     */
    @Export(name="allowedUserKeyConfigs", refs={List.class,SecretBackendRoleAllowedUserKeyConfig.class}, tree="[0,1]")
    private Output</* @Nullable */ List<SecretBackendRoleAllowedUserKeyConfig>> allowedUserKeyConfigs;

    /**
     * @return Set of configuration blocks to define allowed\
     * user key configuration, like key type and their lengths. Can be specified multiple times.
     * *See Configuration-Options for more info*
     * 
     */
    public Output<Optional<List<SecretBackendRoleAllowedUserKeyConfig>>> allowedUserKeyConfigs() {
        return Codegen.optional(this.allowedUserKeyConfigs);
    }
    /**
     * Specifies a comma-separated list of usernames that are to be allowed, only if certain usernames are to be allowed.
     * 
     */
    @Export(name="allowedUsers", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> allowedUsers;

    /**
     * @return Specifies a comma-separated list of usernames that are to be allowed, only if certain usernames are to be allowed.
     * 
     */
    public Output<Optional<String>> allowedUsers() {
        return Codegen.optional(this.allowedUsers);
    }
    /**
     * Specifies if `allowed_users` can be declared using identity template policies. Non-templated users are also permitted.
     * 
     */
    @Export(name="allowedUsersTemplate", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> allowedUsersTemplate;

    /**
     * @return Specifies if `allowed_users` can be declared using identity template policies. Non-templated users are also permitted.
     * 
     */
    public Output<Optional<Boolean>> allowedUsersTemplate() {
        return Codegen.optional(this.allowedUsersTemplate);
    }
    /**
     * The path where the SSH secret backend is mounted.
     * 
     */
    @Export(name="backend", refs={String.class}, tree="[0]")
    private Output<String> backend;

    /**
     * @return The path where the SSH secret backend is mounted.
     * 
     */
    public Output<String> backend() {
        return this.backend;
    }
    /**
     * The comma-separated string of CIDR blocks for which this role is applicable.
     * 
     */
    @Export(name="cidrList", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> cidrList;

    /**
     * @return The comma-separated string of CIDR blocks for which this role is applicable.
     * 
     */
    public Output<Optional<String>> cidrList() {
        return Codegen.optional(this.cidrList);
    }
    /**
     * Specifies a map of critical options that certificates have when signed.
     * 
     */
    @Export(name="defaultCriticalOptions", refs={Map.class,String.class}, tree="[0,1,1]")
    private Output</* @Nullable */ Map<String,String>> defaultCriticalOptions;

    /**
     * @return Specifies a map of critical options that certificates have when signed.
     * 
     */
    public Output<Optional<Map<String,String>>> defaultCriticalOptions() {
        return Codegen.optional(this.defaultCriticalOptions);
    }
    /**
     * Specifies a map of extensions that certificates have when signed.
     * 
     */
    @Export(name="defaultExtensions", refs={Map.class,String.class}, tree="[0,1,1]")
    private Output</* @Nullable */ Map<String,String>> defaultExtensions;

    /**
     * @return Specifies a map of extensions that certificates have when signed.
     * 
     */
    public Output<Optional<Map<String,String>>> defaultExtensions() {
        return Codegen.optional(this.defaultExtensions);
    }
    /**
     * Specifies the default username for which a credential will be generated.
     * 
     */
    @Export(name="defaultUser", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> defaultUser;

    /**
     * @return Specifies the default username for which a credential will be generated.
     * 
     */
    public Output<Optional<String>> defaultUser() {
        return Codegen.optional(this.defaultUser);
    }
    /**
     * If set, `default_users` can be specified using identity template values. A non-templated user is also permitted.
     * 
     */
    @Export(name="defaultUserTemplate", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> defaultUserTemplate;

    /**
     * @return If set, `default_users` can be specified using identity template values. A non-templated user is also permitted.
     * 
     */
    public Output<Optional<Boolean>> defaultUserTemplate() {
        return Codegen.optional(this.defaultUserTemplate);
    }
    /**
     * Specifies a custom format for the key id of a signed certificate.
     * 
     */
    @Export(name="keyIdFormat", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> keyIdFormat;

    /**
     * @return Specifies a custom format for the key id of a signed certificate.
     * 
     */
    public Output<Optional<String>> keyIdFormat() {
        return Codegen.optional(this.keyIdFormat);
    }
    /**
     * Specifies the type of credentials generated by this role. This can be either `otp`, `dynamic` or `ca`.
     * 
     */
    @Export(name="keyType", refs={String.class}, tree="[0]")
    private Output<String> keyType;

    /**
     * @return Specifies the type of credentials generated by this role. This can be either `otp`, `dynamic` or `ca`.
     * 
     */
    public Output<String> keyType() {
        return this.keyType;
    }
    /**
     * Specifies the maximum Time To Live value.
     * 
     */
    @Export(name="maxTtl", refs={String.class}, tree="[0]")
    private Output<String> maxTtl;

    /**
     * @return Specifies the maximum Time To Live value.
     * 
     */
    public Output<String> maxTtl() {
        return this.maxTtl;
    }
    /**
     * Specifies the name of the role to create.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return Specifies the name of the role to create.
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
     * Specifies the duration by which to backdate the ValidAfter property. Uses duration format strings.
     * 
     */
    @Export(name="notBeforeDuration", refs={String.class}, tree="[0]")
    private Output<String> notBeforeDuration;

    /**
     * @return Specifies the duration by which to backdate the ValidAfter property. Uses duration format strings.
     * 
     */
    public Output<String> notBeforeDuration() {
        return this.notBeforeDuration;
    }
    /**
     * Specifies the Time To Live value.
     * 
     */
    @Export(name="ttl", refs={String.class}, tree="[0]")
    private Output<String> ttl;

    /**
     * @return Specifies the Time To Live value.
     * 
     */
    public Output<String> ttl() {
        return this.ttl;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public SecretBackendRole(java.lang.String name) {
        this(name, SecretBackendRoleArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public SecretBackendRole(java.lang.String name, SecretBackendRoleArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public SecretBackendRole(java.lang.String name, SecretBackendRoleArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("vault:ssh/secretBackendRole:SecretBackendRole", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private SecretBackendRole(java.lang.String name, Output<java.lang.String> id, @Nullable SecretBackendRoleState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("vault:ssh/secretBackendRole:SecretBackendRole", name, state, makeResourceOptions(options, id), false);
    }

    private static SecretBackendRoleArgs makeArgs(SecretBackendRoleArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? SecretBackendRoleArgs.Empty : args;
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<java.lang.String> id) {
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
    public static SecretBackendRole get(java.lang.String name, Output<java.lang.String> id, @Nullable SecretBackendRoleState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new SecretBackendRole(name, id, state, options);
    }
}
