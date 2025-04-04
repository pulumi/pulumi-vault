// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vault.ldap;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.vault.Utilities;
import com.pulumi.vault.ldap.AuthBackendArgs;
import com.pulumi.vault.ldap.inputs.AuthBackendState;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides a resource for managing an [LDAP auth backend within Vault](https://www.vaultproject.io/docs/auth/ldap.html).
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
 * import com.pulumi.vault.ldap.AuthBackend;
 * import com.pulumi.vault.ldap.AuthBackendArgs;
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
 *         var ldap = new AuthBackend("ldap", AuthBackendArgs.builder()
 *             .path("ldap")
 *             .url("ldaps://dc-01.example.org")
 *             .userdn("OU=Users,OU=Accounts,DC=example,DC=org")
 *             .userattr("sAMAccountName")
 *             .upndomain("EXAMPLE.ORG")
 *             .discoverdn(false)
 *             .groupdn("OU=Groups,DC=example,DC=org")
 *             .groupfilter("(&(objectClass=group)(member:1.2.840.113556.1.4.1941:={{.UserDN}}))")
 *             .rotationSchedule("0 * * * SAT")
 *             .rotationWindow(3600)
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
 * LDAP authentication backends can be imported using the `path`, e.g.
 * 
 * ```sh
 * $ pulumi import vault:ldap/authBackend:AuthBackend ldap ldap
 * ```
 * 
 */
@ResourceType(type="vault:ldap/authBackend:AuthBackend")
public class AuthBackend extends com.pulumi.resources.CustomResource {
    /**
     * The accessor for this auth mount.
     * 
     */
    @Export(name="accessor", refs={String.class}, tree="[0]")
    private Output<String> accessor;

    /**
     * @return The accessor for this auth mount.
     * 
     */
    public Output<String> accessor() {
        return this.accessor;
    }
    /**
     * DN of object to bind when performing user search
     * 
     */
    @Export(name="binddn", refs={String.class}, tree="[0]")
    private Output<String> binddn;

    /**
     * @return DN of object to bind when performing user search
     * 
     */
    public Output<String> binddn() {
        return this.binddn;
    }
    /**
     * Password to use with `binddn` when performing user search
     * 
     */
    @Export(name="bindpass", refs={String.class}, tree="[0]")
    private Output<String> bindpass;

    /**
     * @return Password to use with `binddn` when performing user search
     * 
     */
    public Output<String> bindpass() {
        return this.bindpass;
    }
    /**
     * Control case senstivity of objects fetched from LDAP, this is used for object matching in vault
     * 
     */
    @Export(name="caseSensitiveNames", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> caseSensitiveNames;

    /**
     * @return Control case senstivity of objects fetched from LDAP, this is used for object matching in vault
     * 
     */
    public Output<Boolean> caseSensitiveNames() {
        return this.caseSensitiveNames;
    }
    /**
     * Trusted CA to validate TLS certificate
     * 
     */
    @Export(name="certificate", refs={String.class}, tree="[0]")
    private Output<String> certificate;

    /**
     * @return Trusted CA to validate TLS certificate
     * 
     */
    public Output<String> certificate() {
        return this.certificate;
    }
    @Export(name="clientTlsCert", refs={String.class}, tree="[0]")
    private Output<String> clientTlsCert;

    public Output<String> clientTlsCert() {
        return this.clientTlsCert;
    }
    @Export(name="clientTlsKey", refs={String.class}, tree="[0]")
    private Output<String> clientTlsKey;

    public Output<String> clientTlsKey() {
        return this.clientTlsKey;
    }
    /**
     * Timeout in seconds when connecting to LDAP before attempting to connect to the next server in the URL provided in `url` (integer: 30)
     * 
     */
    @Export(name="connectionTimeout", refs={Integer.class}, tree="[0]")
    private Output<Integer> connectionTimeout;

    /**
     * @return Timeout in seconds when connecting to LDAP before attempting to connect to the next server in the URL provided in `url` (integer: 30)
     * 
     */
    public Output<Integer> connectionTimeout() {
        return this.connectionTimeout;
    }
    /**
     * Prevents users from bypassing authentication when providing an empty password.
     * 
     */
    @Export(name="denyNullBind", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> denyNullBind;

    /**
     * @return Prevents users from bypassing authentication when providing an empty password.
     * 
     */
    public Output<Boolean> denyNullBind() {
        return this.denyNullBind;
    }
    /**
     * Description for the LDAP auth backend mount
     * 
     */
    @Export(name="description", refs={String.class}, tree="[0]")
    private Output<String> description;

    /**
     * @return Description for the LDAP auth backend mount
     * 
     */
    public Output<String> description() {
        return this.description;
    }
    /**
     * Cancels all upcoming rotations of the root credential until unset. Requires Vault Enterprise 1.19+.
     * 
     */
    @Export(name="disableAutomatedRotation", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> disableAutomatedRotation;

    /**
     * @return Cancels all upcoming rotations of the root credential until unset. Requires Vault Enterprise 1.19+.
     * 
     */
    public Output<Optional<Boolean>> disableAutomatedRotation() {
        return Codegen.optional(this.disableAutomatedRotation);
    }
    /**
     * If set, opts out of mount migration on path updates.
     * See here for more info on [Mount Migration](https://www.vaultproject.io/docs/concepts/mount-migration)
     * 
     */
    @Export(name="disableRemount", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> disableRemount;

    /**
     * @return If set, opts out of mount migration on path updates.
     * See here for more info on [Mount Migration](https://www.vaultproject.io/docs/concepts/mount-migration)
     * 
     */
    public Output<Optional<Boolean>> disableRemount() {
        return Codegen.optional(this.disableRemount);
    }
    /**
     * Use anonymous bind to discover the bind DN of a user.
     * 
     */
    @Export(name="discoverdn", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> discoverdn;

    /**
     * @return Use anonymous bind to discover the bind DN of a user.
     * 
     */
    public Output<Boolean> discoverdn() {
        return this.discoverdn;
    }
    /**
     * LDAP attribute to follow on objects returned by groupfilter
     * 
     */
    @Export(name="groupattr", refs={String.class}, tree="[0]")
    private Output<String> groupattr;

    /**
     * @return LDAP attribute to follow on objects returned by groupfilter
     * 
     */
    public Output<String> groupattr() {
        return this.groupattr;
    }
    /**
     * Base DN under which to perform group search
     * 
     */
    @Export(name="groupdn", refs={String.class}, tree="[0]")
    private Output<String> groupdn;

    /**
     * @return Base DN under which to perform group search
     * 
     */
    public Output<String> groupdn() {
        return this.groupdn;
    }
    /**
     * Go template used to construct group membership query
     * 
     */
    @Export(name="groupfilter", refs={String.class}, tree="[0]")
    private Output<String> groupfilter;

    /**
     * @return Go template used to construct group membership query
     * 
     */
    public Output<String> groupfilter() {
        return this.groupfilter;
    }
    /**
     * Control whether or TLS certificates must be validated
     * 
     */
    @Export(name="insecureTls", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> insecureTls;

    /**
     * @return Control whether or TLS certificates must be validated
     * 
     */
    public Output<Boolean> insecureTls() {
        return this.insecureTls;
    }
    /**
     * Specifies if the auth method is local only.
     * 
     */
    @Export(name="local", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> local;

    /**
     * @return Specifies if the auth method is local only.
     * 
     */
    public Output<Optional<Boolean>> local() {
        return Codegen.optional(this.local);
    }
    /**
     * Sets the max page size for LDAP lookups, by default it&#39;s set to -1.
     * *Available only for Vault 1.11.11+, 1.12.7+, and 1.13.3+*.
     * 
     */
    @Export(name="maxPageSize", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> maxPageSize;

    /**
     * @return Sets the max page size for LDAP lookups, by default it&#39;s set to -1.
     * *Available only for Vault 1.11.11+, 1.12.7+, and 1.13.3+*.
     * 
     */
    public Output<Optional<Integer>> maxPageSize() {
        return Codegen.optional(this.maxPageSize);
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
     * Path to mount the LDAP auth backend under
     * 
     */
    @Export(name="path", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> path;

    /**
     * @return Path to mount the LDAP auth backend under
     * 
     */
    public Output<Optional<String>> path() {
        return Codegen.optional(this.path);
    }
    /**
     * The amount of time in seconds Vault should wait before rotating the root credential.
     * A zero value tells Vault not to rotate the root credential. The minimum rotation period is 10 seconds. Requires Vault Enterprise 1.19+.
     * 
     */
    @Export(name="rotationPeriod", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> rotationPeriod;

    /**
     * @return The amount of time in seconds Vault should wait before rotating the root credential.
     * A zero value tells Vault not to rotate the root credential. The minimum rotation period is 10 seconds. Requires Vault Enterprise 1.19+.
     * 
     */
    public Output<Optional<Integer>> rotationPeriod() {
        return Codegen.optional(this.rotationPeriod);
    }
    /**
     * The schedule, in [cron-style time format](https://en.wikipedia.org/wiki/Cron),
     * defining the schedule on which Vault should rotate the root token. Requires Vault Enterprise 1.19+.
     * 
     */
    @Export(name="rotationSchedule", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> rotationSchedule;

    /**
     * @return The schedule, in [cron-style time format](https://en.wikipedia.org/wiki/Cron),
     * defining the schedule on which Vault should rotate the root token. Requires Vault Enterprise 1.19+.
     * 
     */
    public Output<Optional<String>> rotationSchedule() {
        return Codegen.optional(this.rotationSchedule);
    }
    /**
     * The maximum amount of time in seconds allowed to complete
     * a rotation when a scheduled token rotation occurs. The default rotation window is
     * unbound and the minimum allowable window is `3600`. Requires Vault Enterprise 1.19+.
     * 
     */
    @Export(name="rotationWindow", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> rotationWindow;

    /**
     * @return The maximum amount of time in seconds allowed to complete
     * a rotation when a scheduled token rotation occurs. The default rotation window is
     * unbound and the minimum allowable window is `3600`. Requires Vault Enterprise 1.19+.
     * 
     */
    public Output<Optional<Integer>> rotationWindow() {
        return Codegen.optional(this.rotationWindow);
    }
    /**
     * Control use of TLS when conecting to LDAP
     * 
     */
    @Export(name="starttls", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> starttls;

    /**
     * @return Control use of TLS when conecting to LDAP
     * 
     */
    public Output<Boolean> starttls() {
        return this.starttls;
    }
    /**
     * Maximum acceptable version of TLS
     * 
     */
    @Export(name="tlsMaxVersion", refs={String.class}, tree="[0]")
    private Output<String> tlsMaxVersion;

    /**
     * @return Maximum acceptable version of TLS
     * 
     */
    public Output<String> tlsMaxVersion() {
        return this.tlsMaxVersion;
    }
    /**
     * Minimum acceptable version of TLS
     * 
     */
    @Export(name="tlsMinVersion", refs={String.class}, tree="[0]")
    private Output<String> tlsMinVersion;

    /**
     * @return Minimum acceptable version of TLS
     * 
     */
    public Output<String> tlsMinVersion() {
        return this.tlsMinVersion;
    }
    /**
     * Specifies the blocks of IP addresses which are allowed to use the generated token
     * 
     */
    @Export(name="tokenBoundCidrs", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> tokenBoundCidrs;

    /**
     * @return Specifies the blocks of IP addresses which are allowed to use the generated token
     * 
     */
    public Output<Optional<List<String>>> tokenBoundCidrs() {
        return Codegen.optional(this.tokenBoundCidrs);
    }
    /**
     * Generated Token&#39;s Explicit Maximum TTL in seconds
     * 
     */
    @Export(name="tokenExplicitMaxTtl", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> tokenExplicitMaxTtl;

    /**
     * @return Generated Token&#39;s Explicit Maximum TTL in seconds
     * 
     */
    public Output<Optional<Integer>> tokenExplicitMaxTtl() {
        return Codegen.optional(this.tokenExplicitMaxTtl);
    }
    /**
     * The maximum lifetime of the generated token
     * 
     */
    @Export(name="tokenMaxTtl", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> tokenMaxTtl;

    /**
     * @return The maximum lifetime of the generated token
     * 
     */
    public Output<Optional<Integer>> tokenMaxTtl() {
        return Codegen.optional(this.tokenMaxTtl);
    }
    /**
     * If true, the &#39;default&#39; policy will not automatically be added to generated tokens
     * 
     */
    @Export(name="tokenNoDefaultPolicy", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> tokenNoDefaultPolicy;

    /**
     * @return If true, the &#39;default&#39; policy will not automatically be added to generated tokens
     * 
     */
    public Output<Optional<Boolean>> tokenNoDefaultPolicy() {
        return Codegen.optional(this.tokenNoDefaultPolicy);
    }
    /**
     * The maximum number of times a token may be used, a value of zero means unlimited
     * 
     */
    @Export(name="tokenNumUses", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> tokenNumUses;

    /**
     * @return The maximum number of times a token may be used, a value of zero means unlimited
     * 
     */
    public Output<Optional<Integer>> tokenNumUses() {
        return Codegen.optional(this.tokenNumUses);
    }
    /**
     * Generated Token&#39;s Period
     * 
     */
    @Export(name="tokenPeriod", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> tokenPeriod;

    /**
     * @return Generated Token&#39;s Period
     * 
     */
    public Output<Optional<Integer>> tokenPeriod() {
        return Codegen.optional(this.tokenPeriod);
    }
    /**
     * Generated Token&#39;s Policies
     * 
     */
    @Export(name="tokenPolicies", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> tokenPolicies;

    /**
     * @return Generated Token&#39;s Policies
     * 
     */
    public Output<Optional<List<String>>> tokenPolicies() {
        return Codegen.optional(this.tokenPolicies);
    }
    /**
     * The initial ttl of the token to generate in seconds
     * 
     */
    @Export(name="tokenTtl", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> tokenTtl;

    /**
     * @return The initial ttl of the token to generate in seconds
     * 
     */
    public Output<Optional<Integer>> tokenTtl() {
        return Codegen.optional(this.tokenTtl);
    }
    /**
     * The type of token to generate, service or batch
     * 
     */
    @Export(name="tokenType", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> tokenType;

    /**
     * @return The type of token to generate, service or batch
     * 
     */
    public Output<Optional<String>> tokenType() {
        return Codegen.optional(this.tokenType);
    }
    /**
     * The `userPrincipalDomain` used to construct the UPN string for the authenticating user.
     * 
     */
    @Export(name="upndomain", refs={String.class}, tree="[0]")
    private Output<String> upndomain;

    /**
     * @return The `userPrincipalDomain` used to construct the UPN string for the authenticating user.
     * 
     */
    public Output<String> upndomain() {
        return this.upndomain;
    }
    /**
     * The URL of the LDAP server
     * 
     */
    @Export(name="url", refs={String.class}, tree="[0]")
    private Output<String> url;

    /**
     * @return The URL of the LDAP server
     * 
     */
    public Output<String> url() {
        return this.url;
    }
    /**
     * Use the Active Directory tokenGroups constructed attribute of the user to find the group memberships
     * 
     */
    @Export(name="useTokenGroups", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> useTokenGroups;

    /**
     * @return Use the Active Directory tokenGroups constructed attribute of the user to find the group memberships
     * 
     */
    public Output<Boolean> useTokenGroups() {
        return this.useTokenGroups;
    }
    /**
     * Attribute on user object matching username passed in
     * 
     */
    @Export(name="userattr", refs={String.class}, tree="[0]")
    private Output<String> userattr;

    /**
     * @return Attribute on user object matching username passed in
     * 
     */
    public Output<String> userattr() {
        return this.userattr;
    }
    /**
     * Base DN under which to perform user search
     * 
     */
    @Export(name="userdn", refs={String.class}, tree="[0]")
    private Output<String> userdn;

    /**
     * @return Base DN under which to perform user search
     * 
     */
    public Output<String> userdn() {
        return this.userdn;
    }
    /**
     * LDAP user search filter
     * 
     */
    @Export(name="userfilter", refs={String.class}, tree="[0]")
    private Output<String> userfilter;

    /**
     * @return LDAP user search filter
     * 
     */
    public Output<String> userfilter() {
        return this.userfilter;
    }
    /**
     * Force the auth method to use the username passed by the user as the alias name.
     * 
     */
    @Export(name="usernameAsAlias", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> usernameAsAlias;

    /**
     * @return Force the auth method to use the username passed by the user as the alias name.
     * 
     */
    public Output<Boolean> usernameAsAlias() {
        return this.usernameAsAlias;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public AuthBackend(java.lang.String name) {
        this(name, AuthBackendArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public AuthBackend(java.lang.String name, AuthBackendArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public AuthBackend(java.lang.String name, AuthBackendArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("vault:ldap/authBackend:AuthBackend", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private AuthBackend(java.lang.String name, Output<java.lang.String> id, @Nullable AuthBackendState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("vault:ldap/authBackend:AuthBackend", name, state, makeResourceOptions(options, id), false);
    }

    private static AuthBackendArgs makeArgs(AuthBackendArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? AuthBackendArgs.Empty : args;
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<java.lang.String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .additionalSecretOutputs(List.of(
                "bindpass",
                "clientTlsKey"
            ))
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
    public static AuthBackend get(java.lang.String name, Output<java.lang.String> id, @Nullable AuthBackendState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new AuthBackend(name, id, state, options);
    }
}
