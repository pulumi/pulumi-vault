// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vault.kmip;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.vault.Utilities;
import com.pulumi.vault.kmip.SecretBackendArgs;
import com.pulumi.vault.kmip.inputs.SecretBackendState;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Manages KMIP Secret backends in a Vault server. This feature requires
 * Vault Enterprise. See the [Vault documentation](https://www.vaultproject.io/docs/secrets/kmip)
 * for more information.
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
 * import com.pulumi.vault.kmip.SecretBackend;
 * import com.pulumi.vault.kmip.SecretBackendArgs;
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
 *         var default_ = new SecretBackend("default", SecretBackendArgs.builder()
 *             .path("kmip")
 *             .description("Vault KMIP backend")
 *             .listenAddrs(            
 *                 "127.0.0.1:5696",
 *                 "127.0.0.1:8080")
 *             .tlsCaKeyType("rsa")
 *             .tlsCaKeyBits(4096)
 *             .defaultTlsClientKeyType("rsa")
 *             .defaultTlsClientKeyBits(4096)
 *             .defaultTlsClientTtl(86400)
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
 * KMIP Secret backend can be imported using the `path`, e.g.
 * 
 * ```sh
 * $ pulumi import vault:kmip/secretBackend:SecretBackend default kmip
 * ```
 * 
 */
@ResourceType(type="vault:kmip/secretBackend:SecretBackend")
public class SecretBackend extends com.pulumi.resources.CustomResource {
    /**
     * Client certificate key bits, valid values depend on key type.
     * 
     */
    @Export(name="defaultTlsClientKeyBits", refs={Integer.class}, tree="[0]")
    private Output<Integer> defaultTlsClientKeyBits;

    /**
     * @return Client certificate key bits, valid values depend on key type.
     * 
     */
    public Output<Integer> defaultTlsClientKeyBits() {
        return this.defaultTlsClientKeyBits;
    }
    /**
     * Client certificate key type, `rsa` or `ec`.
     * 
     */
    @Export(name="defaultTlsClientKeyType", refs={String.class}, tree="[0]")
    private Output<String> defaultTlsClientKeyType;

    /**
     * @return Client certificate key type, `rsa` or `ec`.
     * 
     */
    public Output<String> defaultTlsClientKeyType() {
        return this.defaultTlsClientKeyType;
    }
    /**
     * Client certificate TTL in seconds
     * 
     */
    @Export(name="defaultTlsClientTtl", refs={Integer.class}, tree="[0]")
    private Output<Integer> defaultTlsClientTtl;

    /**
     * @return Client certificate TTL in seconds
     * 
     */
    public Output<Integer> defaultTlsClientTtl() {
        return this.defaultTlsClientTtl;
    }
    /**
     * A human-friendly description for this backend.
     * 
     */
    @Export(name="description", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> description;

    /**
     * @return A human-friendly description for this backend.
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
     * Addresses the KMIP server should listen on (`host:port`).
     * 
     */
    @Export(name="listenAddrs", refs={List.class,String.class}, tree="[0,1]")
    private Output<List<String>> listenAddrs;

    /**
     * @return Addresses the KMIP server should listen on (`host:port`).
     * 
     */
    public Output<List<String>> listenAddrs() {
        return this.listenAddrs;
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
     * The unique path this backend should be mounted at. Must
     * not begin or end with a `/`. Defaults to `kmip`.
     * 
     */
    @Export(name="path", refs={String.class}, tree="[0]")
    private Output<String> path;

    /**
     * @return The unique path this backend should be mounted at. Must
     * not begin or end with a `/`. Defaults to `kmip`.
     * 
     */
    public Output<String> path() {
        return this.path;
    }
    /**
     * Hostnames to include in the server&#39;s TLS certificate as SAN DNS names. The first will be used as the common name (CN).
     * 
     */
    @Export(name="serverHostnames", refs={List.class,String.class}, tree="[0,1]")
    private Output<List<String>> serverHostnames;

    /**
     * @return Hostnames to include in the server&#39;s TLS certificate as SAN DNS names. The first will be used as the common name (CN).
     * 
     */
    public Output<List<String>> serverHostnames() {
        return this.serverHostnames;
    }
    /**
     * IPs to include in the server&#39;s TLS certificate as SAN IP addresses.
     * 
     */
    @Export(name="serverIps", refs={List.class,String.class}, tree="[0,1]")
    private Output<List<String>> serverIps;

    /**
     * @return IPs to include in the server&#39;s TLS certificate as SAN IP addresses.
     * 
     */
    public Output<List<String>> serverIps() {
        return this.serverIps;
    }
    /**
     * CA key bits, valid values depend on key type.
     * 
     */
    @Export(name="tlsCaKeyBits", refs={Integer.class}, tree="[0]")
    private Output<Integer> tlsCaKeyBits;

    /**
     * @return CA key bits, valid values depend on key type.
     * 
     */
    public Output<Integer> tlsCaKeyBits() {
        return this.tlsCaKeyBits;
    }
    /**
     * CA key type, rsa or ec.
     * 
     */
    @Export(name="tlsCaKeyType", refs={String.class}, tree="[0]")
    private Output<String> tlsCaKeyType;

    /**
     * @return CA key type, rsa or ec.
     * 
     */
    public Output<String> tlsCaKeyType() {
        return this.tlsCaKeyType;
    }
    /**
     * Minimum TLS version to accept.
     * 
     */
    @Export(name="tlsMinVersion", refs={String.class}, tree="[0]")
    private Output<String> tlsMinVersion;

    /**
     * @return Minimum TLS version to accept.
     * 
     */
    public Output<String> tlsMinVersion() {
        return this.tlsMinVersion;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public SecretBackend(java.lang.String name) {
        this(name, SecretBackendArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public SecretBackend(java.lang.String name, SecretBackendArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public SecretBackend(java.lang.String name, SecretBackendArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("vault:kmip/secretBackend:SecretBackend", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private SecretBackend(java.lang.String name, Output<java.lang.String> id, @Nullable SecretBackendState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("vault:kmip/secretBackend:SecretBackend", name, state, makeResourceOptions(options, id), false);
    }

    private static SecretBackendArgs makeArgs(SecretBackendArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? SecretBackendArgs.Empty : args;
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
    public static SecretBackend get(java.lang.String name, Output<java.lang.String> id, @Nullable SecretBackendState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new SecretBackend(name, id, state, options);
    }
}
