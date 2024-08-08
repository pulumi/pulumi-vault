// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vault.pkiSecret;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.vault.Utilities;
import com.pulumi.vault.pkiSecret.SecretBackendCrlConfigArgs;
import com.pulumi.vault.pkiSecret.inputs.SecretBackendCrlConfigState;
import java.lang.Boolean;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Allows setting the duration for which the generated CRL should be marked valid. If the CRL is disabled, it will return a signed but zero-length CRL for any request. If enabled, it will re-build the CRL.
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
 * import com.pulumi.vault.pkiSecret.SecretBackendCrlConfig;
 * import com.pulumi.vault.pkiSecret.SecretBackendCrlConfigArgs;
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
 *         var pki = new Mount("pki", MountArgs.builder()
 *             .path("%s")
 *             .type("pki")
 *             .defaultLeaseTtlSeconds(3600)
 *             .maxLeaseTtlSeconds(86400)
 *             .build());
 * 
 *         var crlConfig = new SecretBackendCrlConfig("crlConfig", SecretBackendCrlConfigArgs.builder()
 *             .backend(pki.path())
 *             .expiry("72h")
 *             .disable(false)
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 */
@ResourceType(type="vault:pkiSecret/secretBackendCrlConfig:SecretBackendCrlConfig")
public class SecretBackendCrlConfig extends com.pulumi.resources.CustomResource {
    /**
     * Enables periodic rebuilding of the CRL upon expiry. **Vault 1.12+**
     * 
     */
    @Export(name="autoRebuild", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> autoRebuild;

    /**
     * @return Enables periodic rebuilding of the CRL upon expiry. **Vault 1.12+**
     * 
     */
    public Output<Optional<Boolean>> autoRebuild() {
        return Codegen.optional(this.autoRebuild);
    }
    /**
     * Grace period before CRL expiry to attempt rebuild of CRL. **Vault 1.12+**
     * 
     */
    @Export(name="autoRebuildGracePeriod", refs={String.class}, tree="[0]")
    private Output<String> autoRebuildGracePeriod;

    /**
     * @return Grace period before CRL expiry to attempt rebuild of CRL. **Vault 1.12+**
     * 
     */
    public Output<String> autoRebuildGracePeriod() {
        return this.autoRebuildGracePeriod;
    }
    /**
     * The path the PKI secret backend is mounted at, with no leading or trailing `/`s.
     * 
     */
    @Export(name="backend", refs={String.class}, tree="[0]")
    private Output<String> backend;

    /**
     * @return The path the PKI secret backend is mounted at, with no leading or trailing `/`s.
     * 
     */
    public Output<String> backend() {
        return this.backend;
    }
    /**
     * Enable cross-cluster revocation request queues. **Vault 1.13+**
     * 
     */
    @Export(name="crossClusterRevocation", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> crossClusterRevocation;

    /**
     * @return Enable cross-cluster revocation request queues. **Vault 1.13+**
     * 
     */
    public Output<Boolean> crossClusterRevocation() {
        return this.crossClusterRevocation;
    }
    /**
     * Interval to check for new revocations on, to regenerate the delta CRL.
     * 
     */
    @Export(name="deltaRebuildInterval", refs={String.class}, tree="[0]")
    private Output<String> deltaRebuildInterval;

    /**
     * @return Interval to check for new revocations on, to regenerate the delta CRL.
     * 
     */
    public Output<String> deltaRebuildInterval() {
        return this.deltaRebuildInterval;
    }
    /**
     * Disables or enables CRL building.
     * 
     */
    @Export(name="disable", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> disable;

    /**
     * @return Disables or enables CRL building.
     * 
     */
    public Output<Optional<Boolean>> disable() {
        return Codegen.optional(this.disable);
    }
    /**
     * Enables building of delta CRLs with up-to-date revocation information,
     * augmenting the last complete CRL.  **Vault 1.12+**
     * 
     */
    @Export(name="enableDelta", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> enableDelta;

    /**
     * @return Enables building of delta CRLs with up-to-date revocation information,
     * augmenting the last complete CRL.  **Vault 1.12+**
     * 
     */
    public Output<Optional<Boolean>> enableDelta() {
        return Codegen.optional(this.enableDelta);
    }
    /**
     * Specifies the time until expiration.
     * 
     */
    @Export(name="expiry", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> expiry;

    /**
     * @return Specifies the time until expiration.
     * 
     */
    public Output<Optional<String>> expiry() {
        return Codegen.optional(this.expiry);
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
     * Disables the OCSP responder in Vault. **Vault 1.12+**
     * 
     */
    @Export(name="ocspDisable", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> ocspDisable;

    /**
     * @return Disables the OCSP responder in Vault. **Vault 1.12+**
     * 
     */
    public Output<Optional<Boolean>> ocspDisable() {
        return Codegen.optional(this.ocspDisable);
    }
    /**
     * The amount of time an OCSP response can be cached for, useful for OCSP stapling
     * refresh durations. **Vault 1.12+**
     * 
     */
    @Export(name="ocspExpiry", refs={String.class}, tree="[0]")
    private Output<String> ocspExpiry;

    /**
     * @return The amount of time an OCSP response can be cached for, useful for OCSP stapling
     * refresh durations. **Vault 1.12+**
     * 
     */
    public Output<String> ocspExpiry() {
        return this.ocspExpiry;
    }
    /**
     * Enables unified CRL and OCSP building. **Vault 1.13+**
     * 
     */
    @Export(name="unifiedCrl", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> unifiedCrl;

    /**
     * @return Enables unified CRL and OCSP building. **Vault 1.13+**
     * 
     */
    public Output<Boolean> unifiedCrl() {
        return this.unifiedCrl;
    }
    /**
     * Enables serving the unified CRL and OCSP on the existing, previously
     * cluster-local paths. **Vault 1.13+**
     * 
     */
    @Export(name="unifiedCrlOnExistingPaths", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> unifiedCrlOnExistingPaths;

    /**
     * @return Enables serving the unified CRL and OCSP on the existing, previously
     * cluster-local paths. **Vault 1.13+**
     * 
     */
    public Output<Boolean> unifiedCrlOnExistingPaths() {
        return this.unifiedCrlOnExistingPaths;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public SecretBackendCrlConfig(java.lang.String name) {
        this(name, SecretBackendCrlConfigArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public SecretBackendCrlConfig(java.lang.String name, SecretBackendCrlConfigArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public SecretBackendCrlConfig(java.lang.String name, SecretBackendCrlConfigArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("vault:pkiSecret/secretBackendCrlConfig:SecretBackendCrlConfig", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private SecretBackendCrlConfig(java.lang.String name, Output<java.lang.String> id, @Nullable SecretBackendCrlConfigState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("vault:pkiSecret/secretBackendCrlConfig:SecretBackendCrlConfig", name, state, makeResourceOptions(options, id), false);
    }

    private static SecretBackendCrlConfigArgs makeArgs(SecretBackendCrlConfigArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? SecretBackendCrlConfigArgs.Empty : args;
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
    public static SecretBackendCrlConfig get(java.lang.String name, Output<java.lang.String> id, @Nullable SecretBackendCrlConfigState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new SecretBackendCrlConfig(name, id, state, options);
    }
}
