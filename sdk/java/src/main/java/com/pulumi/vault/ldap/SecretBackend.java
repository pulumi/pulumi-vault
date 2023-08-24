// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vault.ldap;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.vault.Utilities;
import com.pulumi.vault.ldap.SecretBackendArgs;
import com.pulumi.vault.ldap.inputs.SecretBackendState;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.Object;
import java.lang.String;
import java.util.List;
import java.util.Map;
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
 * import com.pulumi.vault.ldap.SecretBackend;
 * import com.pulumi.vault.ldap.SecretBackendArgs;
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
 *             .binddn(&#34;CN=Administrator,CN=Users,DC=corp,DC=example,DC=net&#34;)
 *             .bindpass(&#34;SuperSecretPassw0rd&#34;)
 *             .insecureTls(&#34;true&#34;)
 *             .path(&#34;my-custom-ldap&#34;)
 *             .url(&#34;ldaps://localhost&#34;)
 *             .userdn(&#34;CN=Users,DC=corp,DC=example,DC=net&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * LDAP secret backend can be imported using the `${mount}/config`, e.g.
 * 
 * ```sh
 *  $ pulumi import vault:ldap/secretBackend:SecretBackend config ldap/config
 * ```
 * 
 */
@ResourceType(type="vault:ldap/secretBackend:SecretBackend")
public class SecretBackend extends com.pulumi.resources.CustomResource {
    /**
     * Accessor of the mount
     * 
     */
    @Export(name="accessor", refs={String.class}, tree="[0]")
    private Output<String> accessor;

    /**
     * @return Accessor of the mount
     * 
     */
    public Output<String> accessor() {
        return this.accessor;
    }
    /**
     * List of managed key registry entry names that the mount in question is allowed to access
     * 
     */
    @Export(name="allowedManagedKeys", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> allowedManagedKeys;

    /**
     * @return List of managed key registry entry names that the mount in question is allowed to access
     * 
     */
    public Output<Optional<List<String>>> allowedManagedKeys() {
        return Codegen.optional(this.allowedManagedKeys);
    }
    /**
     * Specifies the list of keys that will not be HMAC&#39;d by audit devices in the request data object.
     * 
     */
    @Export(name="auditNonHmacRequestKeys", refs={List.class,String.class}, tree="[0,1]")
    private Output<List<String>> auditNonHmacRequestKeys;

    /**
     * @return Specifies the list of keys that will not be HMAC&#39;d by audit devices in the request data object.
     * 
     */
    public Output<List<String>> auditNonHmacRequestKeys() {
        return this.auditNonHmacRequestKeys;
    }
    /**
     * Specifies the list of keys that will not be HMAC&#39;d by audit devices in the response data object.
     * 
     */
    @Export(name="auditNonHmacResponseKeys", refs={List.class,String.class}, tree="[0,1]")
    private Output<List<String>> auditNonHmacResponseKeys;

    /**
     * @return Specifies the list of keys that will not be HMAC&#39;d by audit devices in the response data object.
     * 
     */
    public Output<List<String>> auditNonHmacResponseKeys() {
        return this.auditNonHmacResponseKeys;
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
     * Timeout, in seconds, when attempting to connect to the LDAP server before trying
     * the next URL in the configuration.
     * 
     */
    @Export(name="connectionTimeout", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> connectionTimeout;

    /**
     * @return Timeout, in seconds, when attempting to connect to the LDAP server before trying
     * the next URL in the configuration.
     * 
     */
    public Output<Optional<Integer>> connectionTimeout() {
        return Codegen.optional(this.connectionTimeout);
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
     * 
     */
    @Export(name="disableRemount", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> disableRemount;

    /**
     * @return If set, opts out of mount migration on path updates.
     * 
     */
    public Output<Optional<Boolean>> disableRemount() {
        return Codegen.optional(this.disableRemount);
    }
    /**
     * Enable the secrets engine to access Vault&#39;s external entropy source
     * 
     */
    @Export(name="externalEntropyAccess", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> externalEntropyAccess;

    /**
     * @return Enable the secrets engine to access Vault&#39;s external entropy source
     * 
     */
    public Output<Optional<Boolean>> externalEntropyAccess() {
        return Codegen.optional(this.externalEntropyAccess);
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
     * Specifies mount type specific options that are passed to the backend
     * 
     */
    @Export(name="options", refs={Map.class,String.class,Object.class}, tree="[0,1,2]")
    private Output</* @Nullable */ Map<String,Object>> options;

    /**
     * @return Specifies mount type specific options that are passed to the backend
     * 
     */
    public Output<Optional<Map<String,Object>>> options() {
        return Codegen.optional(this.options);
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
     * The unique path this backend should be mounted at. Must
     * not begin or end with a `/`. Defaults to `ldap`.
     * 
     */
    @Export(name="path", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> path;

    /**
     * @return The unique path this backend should be mounted at. Must
     * not begin or end with a `/`. Defaults to `ldap`.
     * 
     */
    public Output<Optional<String>> path() {
        return Codegen.optional(this.path);
    }
    /**
     * Timeout, in seconds, for the connection when making requests against the server
     * before returning back an error.
     * 
     */
    @Export(name="requestTimeout", refs={Integer.class}, tree="[0]")
    private Output<Integer> requestTimeout;

    /**
     * @return Timeout, in seconds, for the connection when making requests against the server
     * before returning back an error.
     * 
     */
    public Output<Integer> requestTimeout() {
        return this.requestTimeout;
    }
    /**
     * The LDAP schema to use when storing entry passwords. Valid schemas include openldap, ad, and racf.
     * 
     */
    @Export(name="schema", refs={String.class}, tree="[0]")
    private Output<String> schema;

    /**
     * @return The LDAP schema to use when storing entry passwords. Valid schemas include openldap, ad, and racf.
     * 
     */
    public Output<String> schema() {
        return this.schema;
    }
    /**
     * Enable seal wrapping for the mount, causing values stored by the mount to be wrapped by the seal&#39;s encryption capability
     * 
     */
    @Export(name="sealWrap", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> sealWrap;

    /**
     * @return Enable seal wrapping for the mount, causing values stored by the mount to be wrapped by the seal&#39;s encryption capability
     * 
     */
    public Output<Boolean> sealWrap() {
        return this.sealWrap;
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
    private Output<String> url;

    /**
     * @return LDAP URL to connect to. Multiple URLs can be specified by concatenating
     * them with commas; they will be tried in-order. Defaults to `ldap://127.0.0.1`.
     * 
     */
    public Output<String> url() {
        return this.url;
    }
    /**
     * Attribute used when searching users. Defaults to `cn`.
     * 
     */
    @Export(name="userattr", refs={String.class}, tree="[0]")
    private Output<String> userattr;

    /**
     * @return Attribute used when searching users. Defaults to `cn`.
     * 
     */
    public Output<String> userattr() {
        return this.userattr;
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
        super("vault:ldap/secretBackend:SecretBackend", name, args == null ? SecretBackendArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private SecretBackend(String name, Output<String> id, @Nullable SecretBackendState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("vault:ldap/secretBackend:SecretBackend", name, state, makeResourceOptions(options, id));
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
