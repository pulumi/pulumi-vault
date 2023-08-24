// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vault.ad;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.vault.Utilities;
import com.pulumi.vault.ad.SecretBackendArgs;
import com.pulumi.vault.ad.inputs.SecretBackendState;
import java.lang.Boolean;
import java.lang.Integer;
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
 * import com.pulumi.vault.ad.SecretBackend;
 * import com.pulumi.vault.ad.SecretBackendArgs;
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
 *         var config = new SecretBackend(&#34;config&#34;, SecretBackendArgs.builder()        
 *             .backend(&#34;ad&#34;)
 *             .binddn(&#34;CN=Administrator,CN=Users,DC=corp,DC=example,DC=net&#34;)
 *             .bindpass(&#34;SuperSecretPassw0rd&#34;)
 *             .insecureTls(&#34;true&#34;)
 *             .url(&#34;ldaps://ad&#34;)
 *             .userdn(&#34;CN=Users,DC=corp,DC=example,DC=net&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * AD secret backend can be imported using the `backend`, e.g.
 * 
 * ```sh
 *  $ pulumi import vault:ad/secretBackend:SecretBackend ad ad
 * ```
 * 
 */
@ResourceType(type="vault:ad/secretBackend:SecretBackend")
public class SecretBackend extends com.pulumi.resources.CustomResource {
    /**
     * Use anonymous binds when performing LDAP group searches
     * (if true the initial credentials will still be used for the initial connection test).
     * 
     */
    @Export(name="anonymousGroupSearch", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> anonymousGroupSearch;

    /**
     * @return Use anonymous binds when performing LDAP group searches
     * (if true the initial credentials will still be used for the initial connection test).
     * 
     */
    public Output<Optional<Boolean>> anonymousGroupSearch() {
        return Codegen.optional(this.anonymousGroupSearch);
    }
    /**
     * The unique path this backend should be mounted at. Must
     * not begin or end with a `/`. Defaults to `ad`.
     * 
     */
    @Export(name="backend", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> backend;

    /**
     * @return The unique path this backend should be mounted at. Must
     * not begin or end with a `/`. Defaults to `ad`.
     * 
     */
    public Output<Optional<String>> backend() {
        return Codegen.optional(this.backend);
    }
    /**
     * Distinguished name of object to bind when performing user and group search.
     * 
     */
    @Export(name="binddn", refs={String.class}, tree="[0]")
    private Output<String> binddn;

    /**
     * @return Distinguished name of object to bind when performing user and group search.
     * 
     */
    public Output<String> binddn() {
        return this.binddn;
    }
    /**
     * Password to use along with binddn when performing user search.
     * 
     */
    @Export(name="bindpass", refs={String.class}, tree="[0]")
    private Output<String> bindpass;

    /**
     * @return Password to use along with binddn when performing user search.
     * 
     */
    public Output<String> bindpass() {
        return this.bindpass;
    }
    /**
     * If set, user and group names assigned to policies within the
     * backend will be case sensitive. Otherwise, names will be normalized to lower case.
     * 
     */
    @Export(name="caseSensitiveNames", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> caseSensitiveNames;

    /**
     * @return If set, user and group names assigned to policies within the
     * backend will be case sensitive. Otherwise, names will be normalized to lower case.
     * 
     */
    public Output<Optional<Boolean>> caseSensitiveNames() {
        return Codegen.optional(this.caseSensitiveNames);
    }
    /**
     * CA certificate to use when verifying LDAP server certificate, must be
     * x509 PEM encoded.
     * 
     */
    @Export(name="certificate", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> certificate;

    /**
     * @return CA certificate to use when verifying LDAP server certificate, must be
     * x509 PEM encoded.
     * 
     */
    public Output<Optional<String>> certificate() {
        return Codegen.optional(this.certificate);
    }
    /**
     * Client certificate to provide to the LDAP server, must be x509 PEM encoded.
     * 
     */
    @Export(name="clientTlsCert", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> clientTlsCert;

    /**
     * @return Client certificate to provide to the LDAP server, must be x509 PEM encoded.
     * 
     */
    public Output<Optional<String>> clientTlsCert() {
        return Codegen.optional(this.clientTlsCert);
    }
    /**
     * Client certificate key to provide to the LDAP server, must be x509 PEM encoded.
     * 
     */
    @Export(name="clientTlsKey", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> clientTlsKey;

    /**
     * @return Client certificate key to provide to the LDAP server, must be x509 PEM encoded.
     * 
     */
    public Output<Optional<String>> clientTlsKey() {
        return Codegen.optional(this.clientTlsKey);
    }
    /**
     * Default lease duration for secrets in seconds.
     * 
     */
    @Export(name="defaultLeaseTtlSeconds", refs={Integer.class}, tree="[0]")
    private Output<Integer> defaultLeaseTtlSeconds;

    /**
     * @return Default lease duration for secrets in seconds.
     * 
     */
    public Output<Integer> defaultLeaseTtlSeconds() {
        return this.defaultLeaseTtlSeconds;
    }
    /**
     * Denies an unauthenticated LDAP bind request if the user&#39;s password is empty;
     * defaults to true.
     * 
     */
    @Export(name="denyNullBind", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> denyNullBind;

    /**
     * @return Denies an unauthenticated LDAP bind request if the user&#39;s password is empty;
     * defaults to true.
     * 
     */
    public Output<Optional<Boolean>> denyNullBind() {
        return Codegen.optional(this.denyNullBind);
    }
    /**
     * Human-friendly description of the mount for the Active Directory backend.
     * 
     */
    @Export(name="description", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> description;

    /**
     * @return Human-friendly description of the mount for the Active Directory backend.
     * 
     */
    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
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
     * Use anonymous bind to discover the bind Distinguished Name of a user.
     * 
     */
    @Export(name="discoverdn", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> discoverdn;

    /**
     * @return Use anonymous bind to discover the bind Distinguished Name of a user.
     * 
     */
    public Output<Optional<Boolean>> discoverdn() {
        return Codegen.optional(this.discoverdn);
    }
    /**
     * **Deprecated** use `password_policy`. Text to insert the password into, ex. &#34;customPrefix{{PASSWORD}}customSuffix&#34;.
     * 
     * @deprecated
     * Formatter is deprecated and password_policy should be used with Vault &gt;= 1.5.
     * 
     */
    @Deprecated /* Formatter is deprecated and password_policy should be used with Vault >= 1.5. */
    @Export(name="formatter", refs={String.class}, tree="[0]")
    private Output<String> formatter;

    /**
     * @return **Deprecated** use `password_policy`. Text to insert the password into, ex. &#34;customPrefix{{PASSWORD}}customSuffix&#34;.
     * 
     */
    public Output<String> formatter() {
        return this.formatter;
    }
    /**
     * LDAP attribute to follow on objects returned by &lt;groupfilter&gt; in order to enumerate
     * user group membership. Examples: `cn` or `memberOf`, etc. Defaults to `cn`.
     * 
     */
    @Export(name="groupattr", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> groupattr;

    /**
     * @return LDAP attribute to follow on objects returned by &lt;groupfilter&gt; in order to enumerate
     * user group membership. Examples: `cn` or `memberOf`, etc. Defaults to `cn`.
     * 
     */
    public Output<Optional<String>> groupattr() {
        return Codegen.optional(this.groupattr);
    }
    /**
     * LDAP search base to use for group membership search (eg: ou=Groups,dc=example,dc=org).
     * 
     */
    @Export(name="groupdn", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> groupdn;

    /**
     * @return LDAP search base to use for group membership search (eg: ou=Groups,dc=example,dc=org).
     * 
     */
    public Output<Optional<String>> groupdn() {
        return Codegen.optional(this.groupdn);
    }
    /**
     * Go template for querying group membership of user (optional) The template can access
     * the following context variables: UserDN, Username. Defaults to `(|(memberUid={{.Username}})(member={{.UserDN}})(uniqueMember={{.UserDN}}))`
     * 
     */
    @Export(name="groupfilter", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> groupfilter;

    /**
     * @return Go template for querying group membership of user (optional) The template can access
     * the following context variables: UserDN, Username. Defaults to `(|(memberUid={{.Username}})(member={{.UserDN}})(uniqueMember={{.UserDN}}))`
     * 
     */
    public Output<Optional<String>> groupfilter() {
        return Codegen.optional(this.groupfilter);
    }
    /**
     * Skip LDAP server SSL Certificate verification. This is not recommended for production.
     * Defaults to `false`.
     * 
     */
    @Export(name="insecureTls", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> insecureTls;

    /**
     * @return Skip LDAP server SSL Certificate verification. This is not recommended for production.
     * Defaults to `false`.
     * 
     */
    public Output<Optional<Boolean>> insecureTls() {
        return Codegen.optional(this.insecureTls);
    }
    /**
     * The number of seconds after a Vault rotation where, if Active Directory
     * shows a later rotation, it should be considered out-of-band
     * 
     */
    @Export(name="lastRotationTolerance", refs={Integer.class}, tree="[0]")
    private Output<Integer> lastRotationTolerance;

    /**
     * @return The number of seconds after a Vault rotation where, if Active Directory
     * shows a later rotation, it should be considered out-of-band
     * 
     */
    public Output<Integer> lastRotationTolerance() {
        return this.lastRotationTolerance;
    }
    /**
     * **Deprecated** use `password_policy`. The desired length of passwords that Vault generates.
     * *Mutually exclusive with `password_policy` on vault-1.11+*
     * 
     * @deprecated
     * Length is deprecated and password_policy should be used with Vault &gt;= 1.5.
     * 
     */
    @Deprecated /* Length is deprecated and password_policy should be used with Vault >= 1.5. */
    @Export(name="length", refs={Integer.class}, tree="[0]")
    private Output<Integer> length;

    /**
     * @return **Deprecated** use `password_policy`. The desired length of passwords that Vault generates.
     * *Mutually exclusive with `password_policy` on vault-1.11+*
     * 
     */
    public Output<Integer> length() {
        return this.length;
    }
    /**
     * Mark the secrets engine as local-only. Local engines are not replicated or removed by
     * replication.Tolerance duration to use when checking the last rotation time.
     * 
     */
    @Export(name="local", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> local;

    /**
     * @return Mark the secrets engine as local-only. Local engines are not replicated or removed by
     * replication.Tolerance duration to use when checking the last rotation time.
     * 
     */
    public Output<Optional<Boolean>> local() {
        return Codegen.optional(this.local);
    }
    /**
     * Maximum possible lease duration for secrets in seconds.
     * 
     */
    @Export(name="maxLeaseTtlSeconds", refs={Integer.class}, tree="[0]")
    private Output<Integer> maxLeaseTtlSeconds;

    /**
     * @return Maximum possible lease duration for secrets in seconds.
     * 
     */
    public Output<Integer> maxLeaseTtlSeconds() {
        return this.maxLeaseTtlSeconds;
    }
    /**
     * In seconds, the maximum password time-to-live.
     * 
     */
    @Export(name="maxTtl", refs={Integer.class}, tree="[0]")
    private Output<Integer> maxTtl;

    /**
     * @return In seconds, the maximum password time-to-live.
     * 
     */
    public Output<Integer> maxTtl() {
        return this.maxTtl;
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
     * Name of the password policy to use to generate passwords.
     * 
     */
    @Export(name="passwordPolicy", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> passwordPolicy;

    /**
     * @return Name of the password policy to use to generate passwords.
     * 
     */
    public Output<Optional<String>> passwordPolicy() {
        return Codegen.optional(this.passwordPolicy);
    }
    /**
     * Timeout, in seconds, for the connection when making requests against the server
     * before returning back an error.
     * 
     */
    @Export(name="requestTimeout", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> requestTimeout;

    /**
     * @return Timeout, in seconds, for the connection when making requests against the server
     * before returning back an error.
     * 
     */
    public Output<Optional<Integer>> requestTimeout() {
        return Codegen.optional(this.requestTimeout);
    }
    /**
     * Issue a StartTLS command after establishing unencrypted connection.
     * 
     */
    @Export(name="starttls", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> starttls;

    /**
     * @return Issue a StartTLS command after establishing unencrypted connection.
     * 
     */
    public Output<Boolean> starttls() {
        return this.starttls;
    }
    /**
     * Maximum TLS version to use. Accepted values are `tls10`, `tls11`,
     * `tls12` or `tls13`. Defaults to `tls12`.
     * 
     */
    @Export(name="tlsMaxVersion", refs={String.class}, tree="[0]")
    private Output<String> tlsMaxVersion;

    /**
     * @return Maximum TLS version to use. Accepted values are `tls10`, `tls11`,
     * `tls12` or `tls13`. Defaults to `tls12`.
     * 
     */
    public Output<String> tlsMaxVersion() {
        return this.tlsMaxVersion;
    }
    /**
     * Minimum TLS version to use. Accepted values are `tls10`, `tls11`,
     * `tls12` or `tls13`. Defaults to `tls12`.
     * 
     */
    @Export(name="tlsMinVersion", refs={String.class}, tree="[0]")
    private Output<String> tlsMinVersion;

    /**
     * @return Minimum TLS version to use. Accepted values are `tls10`, `tls11`,
     * `tls12` or `tls13`. Defaults to `tls12`.
     * 
     */
    public Output<String> tlsMinVersion() {
        return this.tlsMinVersion;
    }
    /**
     * In seconds, the default password time-to-live.
     * 
     */
    @Export(name="ttl", refs={Integer.class}, tree="[0]")
    private Output<Integer> ttl;

    /**
     * @return In seconds, the default password time-to-live.
     * 
     */
    public Output<Integer> ttl() {
        return this.ttl;
    }
    /**
     * Enables userPrincipalDomain login with [username]@UPNDomain.
     * 
     */
    @Export(name="upndomain", refs={String.class}, tree="[0]")
    private Output<String> upndomain;

    /**
     * @return Enables userPrincipalDomain login with [username]@UPNDomain.
     * 
     */
    public Output<String> upndomain() {
        return this.upndomain;
    }
    /**
     * LDAP URL to connect to. Multiple URLs can be specified by concatenating
     * them with commas; they will be tried in-order. Defaults to `ldap://127.0.0.1`.
     * 
     */
    @Export(name="url", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> url;

    /**
     * @return LDAP URL to connect to. Multiple URLs can be specified by concatenating
     * them with commas; they will be tried in-order. Defaults to `ldap://127.0.0.1`.
     * 
     */
    public Output<Optional<String>> url() {
        return Codegen.optional(this.url);
    }
    /**
     * In Vault 1.1.1 a fix for handling group CN values of
     * different cases unfortunately introduced a regression that could cause previously defined groups
     * to not be found due to a change in the resulting name. If set true, the pre-1.1.1 behavior for
     * matching group CNs will be used. This is only needed in some upgrade scenarios for backwards
     * compatibility. It is enabled by default if the config is upgraded but disabled by default on
     * new configurations.
     * 
     */
    @Export(name="usePre111GroupCnBehavior", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> usePre111GroupCnBehavior;

    /**
     * @return In Vault 1.1.1 a fix for handling group CN values of
     * different cases unfortunately introduced a regression that could cause previously defined groups
     * to not be found due to a change in the resulting name. If set true, the pre-1.1.1 behavior for
     * matching group CNs will be used. This is only needed in some upgrade scenarios for backwards
     * compatibility. It is enabled by default if the config is upgraded but disabled by default on
     * new configurations.
     * 
     */
    public Output<Boolean> usePre111GroupCnBehavior() {
        return this.usePre111GroupCnBehavior;
    }
    /**
     * If true, use the Active Directory tokenGroups constructed attribute of the
     * user to find the group memberships. This will find all security groups including nested ones.
     * 
     */
    @Export(name="useTokenGroups", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> useTokenGroups;

    /**
     * @return If true, use the Active Directory tokenGroups constructed attribute of the
     * user to find the group memberships. This will find all security groups including nested ones.
     * 
     */
    public Output<Optional<Boolean>> useTokenGroups() {
        return Codegen.optional(this.useTokenGroups);
    }
    /**
     * Attribute used when searching users. Defaults to `cn`.
     * 
     */
    @Export(name="userattr", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> userattr;

    /**
     * @return Attribute used when searching users. Defaults to `cn`.
     * 
     */
    public Output<Optional<String>> userattr() {
        return Codegen.optional(this.userattr);
    }
    /**
     * LDAP domain to use for users (eg: ou=People,dc=example,dc=org)`.
     * 
     */
    @Export(name="userdn", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> userdn;

    /**
     * @return LDAP domain to use for users (eg: ou=People,dc=example,dc=org)`.
     * 
     */
    public Output<Optional<String>> userdn() {
        return Codegen.optional(this.userdn);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public SecretBackend(String name) {
        this(name, SecretBackendArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public SecretBackend(String name, SecretBackendArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public SecretBackend(String name, SecretBackendArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("vault:ad/secretBackend:SecretBackend", name, args == null ? SecretBackendArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private SecretBackend(String name, Output<String> id, @Nullable SecretBackendState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("vault:ad/secretBackend:SecretBackend", name, state, makeResourceOptions(options, id));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .additionalSecretOutputs(List.of(
                "bindpass",
                "clientTlsCert",
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
    public static SecretBackend get(String name, Output<String> id, @Nullable SecretBackendState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new SecretBackend(name, id, state, options);
    }
}
