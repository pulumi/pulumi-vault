// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vault.consul;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.vault.Utilities;
import com.pulumi.vault.consul.SecretBackendArgs;
import com.pulumi.vault.consul.inputs.SecretBackendState;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * ## Example Usage
 * 
 * ### Creating a standard backend resource:
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.vault.consul.SecretBackend;
 * import com.pulumi.vault.consul.SecretBackendArgs;
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
 *         var test = new SecretBackend("test", SecretBackendArgs.builder()
 *             .path("consul")
 *             .description("Manages the Consul backend")
 *             .address("127.0.0.1:8500")
 *             .token("4240861b-ce3d-8530-115a-521ff070dd29")
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ### Creating a backend resource to bootstrap a new Consul instance:
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.vault.consul.SecretBackend;
 * import com.pulumi.vault.consul.SecretBackendArgs;
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
 *         var test = new SecretBackend("test", SecretBackendArgs.builder()
 *             .path("consul")
 *             .description("Bootstrap the Consul backend")
 *             .address("127.0.0.1:8500")
 *             .bootstrap(true)
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
 * Consul secret backends can be imported using the `path`, e.g.
 * 
 * ```sh
 * $ pulumi import vault:consul/secretBackend:SecretBackend example consul
 * ```
 * 
 */
@ResourceType(type="vault:consul/secretBackend:SecretBackend")
public class SecretBackend extends com.pulumi.resources.CustomResource {
    /**
     * Specifies the address of the Consul instance, provided as &#34;host:port&#34; like &#34;127.0.0.1:8500&#34;.
     * 
     */
    @Export(name="address", refs={String.class}, tree="[0]")
    private Output<String> address;

    /**
     * @return Specifies the address of the Consul instance, provided as &#34;host:port&#34; like &#34;127.0.0.1:8500&#34;.
     * 
     */
    public Output<String> address() {
        return this.address;
    }
    /**
     * Denotes a backend resource that is used to bootstrap the Consul ACL system. Only one resource may be used to bootstrap.
     * 
     */
    @Export(name="bootstrap", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> bootstrap;

    /**
     * @return Denotes a backend resource that is used to bootstrap the Consul ACL system. Only one resource may be used to bootstrap.
     * 
     */
    public Output<Optional<Boolean>> bootstrap() {
        return Codegen.optional(this.bootstrap);
    }
    /**
     * CA certificate to use when verifying Consul server certificate, must be x509 PEM encoded.
     * 
     */
    @Export(name="caCert", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> caCert;

    /**
     * @return CA certificate to use when verifying Consul server certificate, must be x509 PEM encoded.
     * 
     */
    public Output<Optional<String>> caCert() {
        return Codegen.optional(this.caCert);
    }
    /**
     * Client certificate used for Consul&#39;s TLS communication, must be x509 PEM encoded and if
     * this is set you need to also set client_key.
     * 
     */
    @Export(name="clientCert", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> clientCert;

    /**
     * @return Client certificate used for Consul&#39;s TLS communication, must be x509 PEM encoded and if
     * this is set you need to also set client_key.
     * 
     */
    public Output<Optional<String>> clientCert() {
        return Codegen.optional(this.clientCert);
    }
    /**
     * Client key used for Consul&#39;s TLS communication, must be x509 PEM encoded and if this is set
     * you need to also set client_cert.
     * 
     */
    @Export(name="clientKey", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> clientKey;

    /**
     * @return Client key used for Consul&#39;s TLS communication, must be x509 PEM encoded and if this is set
     * you need to also set client_cert.
     * 
     */
    public Output<Optional<String>> clientKey() {
        return Codegen.optional(this.clientKey);
    }
    /**
     * The default TTL for credentials issued by this backend.
     * 
     */
    @Export(name="defaultLeaseTtlSeconds", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> defaultLeaseTtlSeconds;

    /**
     * @return The default TTL for credentials issued by this backend.
     * 
     */
    public Output<Optional<Integer>> defaultLeaseTtlSeconds() {
        return Codegen.optional(this.defaultLeaseTtlSeconds);
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
     * Specifies if the secret backend is local only.
     * 
     */
    @Export(name="local", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> local;

    /**
     * @return Specifies if the secret backend is local only.
     * 
     */
    public Output<Optional<Boolean>> local() {
        return Codegen.optional(this.local);
    }
    /**
     * The maximum TTL that can be requested
     * for credentials issued by this backend.
     * 
     */
    @Export(name="maxLeaseTtlSeconds", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> maxLeaseTtlSeconds;

    /**
     * @return The maximum TTL that can be requested
     * for credentials issued by this backend.
     * 
     */
    public Output<Optional<Integer>> maxLeaseTtlSeconds() {
        return Codegen.optional(this.maxLeaseTtlSeconds);
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
     * The unique location this backend should be mounted at. Must not begin or end with a `/`. Defaults
     * to `consul`.
     * 
     */
    @Export(name="path", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> path;

    /**
     * @return The unique location this backend should be mounted at. Must not begin or end with a `/`. Defaults
     * to `consul`.
     * 
     */
    public Output<Optional<String>> path() {
        return Codegen.optional(this.path);
    }
    /**
     * Specifies the URL scheme to use. Defaults to `http`.
     * 
     */
    @Export(name="scheme", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> scheme;

    /**
     * @return Specifies the URL scheme to use. Defaults to `http`.
     * 
     */
    public Output<Optional<String>> scheme() {
        return Codegen.optional(this.scheme);
    }
    /**
     * Specifies the Consul token to use when managing or issuing new tokens.
     * 
     */
    @Export(name="token", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> token;

    /**
     * @return Specifies the Consul token to use when managing or issuing new tokens.
     * 
     */
    public Output<Optional<String>> token() {
        return Codegen.optional(this.token);
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
        super("vault:consul/secretBackend:SecretBackend", name, args == null ? SecretBackendArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private SecretBackend(String name, Output<String> id, @Nullable SecretBackendState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("vault:consul/secretBackend:SecretBackend", name, state, makeResourceOptions(options, id));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .additionalSecretOutputs(List.of(
                "clientCert",
                "clientKey",
                "token"
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
