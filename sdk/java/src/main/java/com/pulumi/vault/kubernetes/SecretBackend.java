// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vault.kubernetes;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.vault.Utilities;
import com.pulumi.vault.kubernetes.SecretBackendArgs;
import com.pulumi.vault.kubernetes.inputs.SecretBackendState;
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
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.vault.kubernetes.SecretBackend;
 * import com.pulumi.vault.kubernetes.SecretBackendArgs;
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
 *         var config = new SecretBackend("config", SecretBackendArgs.builder()
 *             .path("kubernetes")
 *             .description("kubernetes secrets engine description")
 *             .defaultLeaseTtlSeconds(43200)
 *             .maxLeaseTtlSeconds(86400)
 *             .kubernetesHost("https://127.0.0.1:61233")
 *             .kubernetesCaCert(StdFunctions.file(FileArgs.builder()
 *                 .input("/path/to/cert")
 *                 .build()).result())
 *             .serviceAccountJwt(StdFunctions.file(FileArgs.builder()
 *                 .input("/path/to/token")
 *                 .build()).result())
 *             .disableLocalCaJwt(false)
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
 * The Kubernetes secret backend can be imported using its `path` e.g.
 * 
 * ```sh
 * $ pulumi import vault:kubernetes/secretBackend:SecretBackend config kubernetes
 * ```
 * 
 */
@ResourceType(type="vault:kubernetes/secretBackend:SecretBackend")
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
     * List of headers to allow and pass from the request to the plugin
     * 
     */
    @Export(name="allowedResponseHeaders", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> allowedResponseHeaders;

    /**
     * @return List of headers to allow and pass from the request to the plugin
     * 
     */
    public Output<Optional<List<String>>> allowedResponseHeaders() {
        return Codegen.optional(this.allowedResponseHeaders);
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
     * Default lease duration for tokens and secrets in seconds
     * 
     */
    @Export(name="defaultLeaseTtlSeconds", refs={Integer.class}, tree="[0]")
    private Output<Integer> defaultLeaseTtlSeconds;

    /**
     * @return Default lease duration for tokens and secrets in seconds
     * 
     */
    public Output<Integer> defaultLeaseTtlSeconds() {
        return this.defaultLeaseTtlSeconds;
    }
    /**
     * List of headers to allow and pass from the request to the plugin
     * 
     */
    @Export(name="delegatedAuthAccessors", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> delegatedAuthAccessors;

    /**
     * @return List of headers to allow and pass from the request to the plugin
     * 
     */
    public Output<Optional<List<String>>> delegatedAuthAccessors() {
        return Codegen.optional(this.delegatedAuthAccessors);
    }
    /**
     * Human-friendly description of the mount
     * 
     */
    @Export(name="description", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> description;

    /**
     * @return Human-friendly description of the mount
     * 
     */
    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    /**
     * Disable defaulting to the local CA certificate and
     * service account JWT when Vault is running in a Kubernetes pod.
     * 
     */
    @Export(name="disableLocalCaJwt", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> disableLocalCaJwt;

    /**
     * @return Disable defaulting to the local CA certificate and
     * service account JWT when Vault is running in a Kubernetes pod.
     * 
     */
    public Output<Optional<Boolean>> disableLocalCaJwt() {
        return Codegen.optional(this.disableLocalCaJwt);
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
     * The key to use for signing plugin workload identity tokens
     * 
     */
    @Export(name="identityTokenKey", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> identityTokenKey;

    /**
     * @return The key to use for signing plugin workload identity tokens
     * 
     */
    public Output<Optional<String>> identityTokenKey() {
        return Codegen.optional(this.identityTokenKey);
    }
    /**
     * A PEM-encoded CA certificate used by the
     * secrets engine to verify the Kubernetes API server certificate. Defaults to the local
     * pod’s CA if Vault is running in Kubernetes. Otherwise, defaults to the root CA set where
     * Vault is running.
     * 
     */
    @Export(name="kubernetesCaCert", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> kubernetesCaCert;

    /**
     * @return A PEM-encoded CA certificate used by the
     * secrets engine to verify the Kubernetes API server certificate. Defaults to the local
     * pod’s CA if Vault is running in Kubernetes. Otherwise, defaults to the root CA set where
     * Vault is running.
     * 
     */
    public Output<Optional<String>> kubernetesCaCert() {
        return Codegen.optional(this.kubernetesCaCert);
    }
    /**
     * The Kubernetes API URL to connect to. Required if the
     * standard pod environment variables `KUBERNETES_SERVICE_HOST` or `KUBERNETES_SERVICE_PORT`
     * are not set on the host that Vault is running on.
     * 
     */
    @Export(name="kubernetesHost", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> kubernetesHost;

    /**
     * @return The Kubernetes API URL to connect to. Required if the
     * standard pod environment variables `KUBERNETES_SERVICE_HOST` or `KUBERNETES_SERVICE_PORT`
     * are not set on the host that Vault is running on.
     * 
     */
    public Output<Optional<String>> kubernetesHost() {
        return Codegen.optional(this.kubernetesHost);
    }
    /**
     * Specifies whether to show this mount in the UI-specific listing endpoint
     * 
     */
    @Export(name="listingVisibility", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> listingVisibility;

    /**
     * @return Specifies whether to show this mount in the UI-specific listing endpoint
     * 
     */
    public Output<Optional<String>> listingVisibility() {
        return Codegen.optional(this.listingVisibility);
    }
    /**
     * Local mount flag that can be explicitly set to true to enforce local mount in HA environment
     * 
     */
    @Export(name="local", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> local;

    /**
     * @return Local mount flag that can be explicitly set to true to enforce local mount in HA environment
     * 
     */
    public Output<Optional<Boolean>> local() {
        return Codegen.optional(this.local);
    }
    /**
     * Maximum possible lease duration for tokens and secrets in seconds
     * 
     */
    @Export(name="maxLeaseTtlSeconds", refs={Integer.class}, tree="[0]")
    private Output<Integer> maxLeaseTtlSeconds;

    /**
     * @return Maximum possible lease duration for tokens and secrets in seconds
     * 
     */
    public Output<Integer> maxLeaseTtlSeconds() {
        return this.maxLeaseTtlSeconds;
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
     * List of headers to allow and pass from the request to the plugin
     * 
     */
    @Export(name="passthroughRequestHeaders", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> passthroughRequestHeaders;

    /**
     * @return List of headers to allow and pass from the request to the plugin
     * 
     */
    public Output<Optional<List<String>>> passthroughRequestHeaders() {
        return Codegen.optional(this.passthroughRequestHeaders);
    }
    /**
     * Where the secret backend will be mounted
     * 
     */
    @Export(name="path", refs={String.class}, tree="[0]")
    private Output<String> path;

    /**
     * @return Where the secret backend will be mounted
     * 
     */
    public Output<String> path() {
        return this.path;
    }
    /**
     * Specifies the semantic version of the plugin to use, e.g. &#39;v1.0.0&#39;
     * 
     */
    @Export(name="pluginVersion", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> pluginVersion;

    /**
     * @return Specifies the semantic version of the plugin to use, e.g. &#39;v1.0.0&#39;
     * 
     */
    public Output<Optional<String>> pluginVersion() {
        return Codegen.optional(this.pluginVersion);
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
     * The JSON web token of the service account used by the
     * secrets engine to manage Kubernetes credentials. Defaults to the local pod’s JWT if Vault
     * is running in Kubernetes.
     * 
     */
    @Export(name="serviceAccountJwt", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> serviceAccountJwt;

    /**
     * @return The JSON web token of the service account used by the
     * secrets engine to manage Kubernetes credentials. Defaults to the local pod’s JWT if Vault
     * is running in Kubernetes.
     * 
     */
    public Output<Optional<String>> serviceAccountJwt() {
        return Codegen.optional(this.serviceAccountJwt);
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
        super("vault:kubernetes/secretBackend:SecretBackend", name, args == null ? SecretBackendArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private SecretBackend(String name, Output<String> id, @Nullable SecretBackendState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("vault:kubernetes/secretBackend:SecretBackend", name, state, makeResourceOptions(options, id));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .additionalSecretOutputs(List.of(
                "serviceAccountJwt"
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
