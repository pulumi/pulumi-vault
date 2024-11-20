// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vault.kubernetes;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.vault.Utilities;
import com.pulumi.vault.kubernetes.AuthBackendConfigArgs;
import com.pulumi.vault.kubernetes.inputs.AuthBackendConfigState;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Manages an Kubernetes auth backend config in a Vault server. See the [Vault
 * documentation](https://www.vaultproject.io/docs/auth/kubernetes.html) for more
 * information.
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
 * import com.pulumi.vault.AuthBackend;
 * import com.pulumi.vault.AuthBackendArgs;
 * import com.pulumi.vault.kubernetes.AuthBackendConfig;
 * import com.pulumi.vault.kubernetes.AuthBackendConfigArgs;
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
 *         var kubernetes = new AuthBackend("kubernetes", AuthBackendArgs.builder()
 *             .type("kubernetes")
 *             .build());
 * 
 *         var example = new AuthBackendConfig("example", AuthBackendConfigArgs.builder()
 *             .backend(kubernetes.path())
 *             .kubernetesHost("http://example.com:443")
 *             .kubernetesCaCert("""
 * -----BEGIN CERTIFICATE-----
 * example
 * -----END CERTIFICATE-----            """)
 *             .tokenReviewerJwt("ZXhhbXBsZQo=")
 *             .issuer("api")
 *             .disableIssValidation("true")
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
 * Kubernetes authentication backend can be imported using the `path`, e.g.
 * 
 * ```sh
 * $ pulumi import vault:kubernetes/authBackendConfig:AuthBackendConfig config auth/kubernetes/config
 * ```
 * 
 */
@ResourceType(type="vault:kubernetes/authBackendConfig:AuthBackendConfig")
public class AuthBackendConfig extends com.pulumi.resources.CustomResource {
    /**
     * Unique name of the kubernetes backend to configure.
     * 
     */
    @Export(name="backend", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> backend;

    /**
     * @return Unique name of the kubernetes backend to configure.
     * 
     */
    public Output<Optional<String>> backend() {
        return Codegen.optional(this.backend);
    }
    /**
     * Disable JWT issuer validation. Allows to skip ISS validation. Requires Vault `v1.5.4+` or Vault auth kubernetes plugin `v0.7.1+`
     * 
     */
    @Export(name="disableIssValidation", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> disableIssValidation;

    /**
     * @return Disable JWT issuer validation. Allows to skip ISS validation. Requires Vault `v1.5.4+` or Vault auth kubernetes plugin `v0.7.1+`
     * 
     */
    public Output<Boolean> disableIssValidation() {
        return this.disableIssValidation;
    }
    /**
     * Disable defaulting to the local CA cert and service account JWT when running in a Kubernetes pod. Requires Vault `v1.5.4+` or Vault auth kubernetes plugin `v0.7.1+`
     * 
     */
    @Export(name="disableLocalCaJwt", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> disableLocalCaJwt;

    /**
     * @return Disable defaulting to the local CA cert and service account JWT when running in a Kubernetes pod. Requires Vault `v1.5.4+` or Vault auth kubernetes plugin `v0.7.1+`
     * 
     */
    public Output<Boolean> disableLocalCaJwt() {
        return this.disableLocalCaJwt;
    }
    /**
     * JWT issuer. If no issuer is specified, `kubernetes.io/serviceaccount` will be used as the default issuer.
     * 
     */
    @Export(name="issuer", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> issuer;

    /**
     * @return JWT issuer. If no issuer is specified, `kubernetes.io/serviceaccount` will be used as the default issuer.
     * 
     */
    public Output<Optional<String>> issuer() {
        return Codegen.optional(this.issuer);
    }
    /**
     * PEM encoded CA cert for use by the TLS client used to talk with the Kubernetes API.
     * 
     */
    @Export(name="kubernetesCaCert", refs={String.class}, tree="[0]")
    private Output<String> kubernetesCaCert;

    /**
     * @return PEM encoded CA cert for use by the TLS client used to talk with the Kubernetes API.
     * 
     */
    public Output<String> kubernetesCaCert() {
        return this.kubernetesCaCert;
    }
    /**
     * Host must be a host string, a host:port pair, or a URL to the base of the Kubernetes API server.
     * 
     */
    @Export(name="kubernetesHost", refs={String.class}, tree="[0]")
    private Output<String> kubernetesHost;

    /**
     * @return Host must be a host string, a host:port pair, or a URL to the base of the Kubernetes API server.
     * 
     */
    public Output<String> kubernetesHost() {
        return this.kubernetesHost;
    }
    /**
     * The namespace to provision the resource in.
     * The value should not contain leading or trailing forward slashes.
     * The `namespace` is always relative to the provider&#39;s configured namespace.
     * *Available only for Vault Enterprise*.
     * 
     */
    @Export(name="namespace", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> namespace;

    /**
     * @return The namespace to provision the resource in.
     * The value should not contain leading or trailing forward slashes.
     * The `namespace` is always relative to the provider&#39;s configured namespace.
     * *Available only for Vault Enterprise*.
     * 
     */
    public Output<Optional<String>> namespace() {
        return Codegen.optional(this.namespace);
    }
    /**
     * List of PEM-formatted public keys or certificates used to verify the signatures of Kubernetes service account JWTs. If a certificate is given, its public key will be extracted. Not every installation of Kubernetes exposes these keys.
     * 
     */
    @Export(name="pemKeys", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> pemKeys;

    /**
     * @return List of PEM-formatted public keys or certificates used to verify the signatures of Kubernetes service account JWTs. If a certificate is given, its public key will be extracted. Not every installation of Kubernetes exposes these keys.
     * 
     */
    public Output<Optional<List<String>>> pemKeys() {
        return Codegen.optional(this.pemKeys);
    }
    /**
     * A service account JWT (or other token) used as a bearer token to access the TokenReview API to validate other JWTs during login. If not set the JWT used for login will be used to access the API.
     * 
     */
    @Export(name="tokenReviewerJwt", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> tokenReviewerJwt;

    /**
     * @return A service account JWT (or other token) used as a bearer token to access the TokenReview API to validate other JWTs during login. If not set the JWT used for login will be used to access the API.
     * 
     */
    public Output<Optional<String>> tokenReviewerJwt() {
        return Codegen.optional(this.tokenReviewerJwt);
    }
    /**
     * Use annotations from the client token&#39;s associated service account as alias metadata for the Vault entity. Requires Vault `v1.16+` or Vault auth kubernetes plugin `v0.18.0+`
     * 
     */
    @Export(name="useAnnotationsAsAliasMetadata", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> useAnnotationsAsAliasMetadata;

    /**
     * @return Use annotations from the client token&#39;s associated service account as alias metadata for the Vault entity. Requires Vault `v1.16+` or Vault auth kubernetes plugin `v0.18.0+`
     * 
     */
    public Output<Boolean> useAnnotationsAsAliasMetadata() {
        return this.useAnnotationsAsAliasMetadata;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public AuthBackendConfig(java.lang.String name) {
        this(name, AuthBackendConfigArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public AuthBackendConfig(java.lang.String name, AuthBackendConfigArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public AuthBackendConfig(java.lang.String name, AuthBackendConfigArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("vault:kubernetes/authBackendConfig:AuthBackendConfig", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private AuthBackendConfig(java.lang.String name, Output<java.lang.String> id, @Nullable AuthBackendConfigState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("vault:kubernetes/authBackendConfig:AuthBackendConfig", name, state, makeResourceOptions(options, id), false);
    }

    private static AuthBackendConfigArgs makeArgs(AuthBackendConfigArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? AuthBackendConfigArgs.Empty : args;
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<java.lang.String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .additionalSecretOutputs(List.of(
                "tokenReviewerJwt"
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
    public static AuthBackendConfig get(java.lang.String name, Output<java.lang.String> id, @Nullable AuthBackendConfigState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new AuthBackendConfig(name, id, state, options);
    }
}
