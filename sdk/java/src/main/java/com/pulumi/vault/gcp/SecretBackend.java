// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vault.gcp;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.vault.Utilities;
import com.pulumi.vault.gcp.SecretBackendArgs;
import com.pulumi.vault.gcp.inputs.SecretBackendState;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * ## Example Usage
 * 
 * You can setup the GCP secret backend with Workload Identity Federation (WIF) for a secret-less configuration:
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.vault.gcp.SecretBackend;
 * import com.pulumi.vault.gcp.SecretBackendArgs;
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
 *         var gcp = new SecretBackend("gcp", SecretBackendArgs.builder()
 *             .identityTokenKey("example-key")
 *             .identityTokenTtl(1800)
 *             .identityTokenAudience("<TOKEN_AUDIENCE>")
 *             .serviceAccountEmail("<SERVICE_ACCOUNT_EMAIL>")
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.vault.gcp.SecretBackend;
 * import com.pulumi.vault.gcp.SecretBackendArgs;
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
 *         var gcp = new SecretBackend("gcp", SecretBackendArgs.builder()
 *             .credentials(StdFunctions.file(FileArgs.builder()
 *                 .input("credentials.json")
 *                 .build()).result())
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 */
@ResourceType(type="vault:gcp/secretBackend:SecretBackend")
public class SecretBackend extends com.pulumi.resources.CustomResource {
    /**
     * The accessor of the created GCP mount.
     * 
     */
    @Export(name="accessor", refs={String.class}, tree="[0]")
    private Output<String> accessor;

    /**
     * @return The accessor of the created GCP mount.
     * 
     */
    public Output<String> accessor() {
        return this.accessor;
    }
    /**
     * JSON-encoded credentials to use to connect to GCP
     * 
     */
    @Export(name="credentials", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> credentials;

    /**
     * @return JSON-encoded credentials to use to connect to GCP
     * 
     */
    public Output<Optional<String>> credentials() {
        return Codegen.optional(this.credentials);
    }
    /**
     * The default TTL for credentials
     * issued by this backend. Defaults to &#39;0&#39;.
     * 
     */
    @Export(name="defaultLeaseTtlSeconds", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> defaultLeaseTtlSeconds;

    /**
     * @return The default TTL for credentials
     * issued by this backend. Defaults to &#39;0&#39;.
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
     * The audience claim value for plugin identity
     * tokens. Must match an allowed audience configured for the target [Workload Identity Pool](https://cloud.google.com/iam/docs/workload-identity-federation-with-other-providers#prepare).
     * Mutually exclusive with `credentials`.  Requires Vault 1.17+. *Available only for Vault Enterprise*.
     * 
     */
    @Export(name="identityTokenAudience", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> identityTokenAudience;

    /**
     * @return The audience claim value for plugin identity
     * tokens. Must match an allowed audience configured for the target [Workload Identity Pool](https://cloud.google.com/iam/docs/workload-identity-federation-with-other-providers#prepare).
     * Mutually exclusive with `credentials`.  Requires Vault 1.17+. *Available only for Vault Enterprise*.
     * 
     */
    public Output<Optional<String>> identityTokenAudience() {
        return Codegen.optional(this.identityTokenAudience);
    }
    /**
     * The key to use for signing plugin identity
     * tokens. Requires Vault 1.17+. *Available only for Vault Enterprise*.
     * 
     */
    @Export(name="identityTokenKey", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> identityTokenKey;

    /**
     * @return The key to use for signing plugin identity
     * tokens. Requires Vault 1.17+. *Available only for Vault Enterprise*.
     * 
     */
    public Output<Optional<String>> identityTokenKey() {
        return Codegen.optional(this.identityTokenKey);
    }
    /**
     * The TTL of generated tokens.
     * 
     */
    @Export(name="identityTokenTtl", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> identityTokenTtl;

    /**
     * @return The TTL of generated tokens.
     * 
     */
    public Output<Optional<Integer>> identityTokenTtl() {
        return Codegen.optional(this.identityTokenTtl);
    }
    /**
     * Boolean flag that can be explicitly set to true to enforce local mount in HA environment
     * 
     */
    @Export(name="local", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> local;

    /**
     * @return Boolean flag that can be explicitly set to true to enforce local mount in HA environment
     * 
     */
    public Output<Optional<Boolean>> local() {
        return Codegen.optional(this.local);
    }
    /**
     * The maximum TTL that can be requested
     * for credentials issued by this backend. Defaults to &#39;0&#39;.
     * 
     */
    @Export(name="maxLeaseTtlSeconds", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> maxLeaseTtlSeconds;

    /**
     * @return The maximum TTL that can be requested
     * for credentials issued by this backend. Defaults to &#39;0&#39;.
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
     * The unique path this backend should be mounted at. Must
     * not begin or end with a `/`. Defaults to `gcp`.
     * 
     */
    @Export(name="path", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> path;

    /**
     * @return The unique path this backend should be mounted at. Must
     * not begin or end with a `/`. Defaults to `gcp`.
     * 
     */
    public Output<Optional<String>> path() {
        return Codegen.optional(this.path);
    }
    /**
     * Service Account to impersonate for plugin workload identity federation.
     * Required with `identity_token_audience`. Requires Vault 1.17+. *Available only for Vault Enterprise*.
     * 
     */
    @Export(name="serviceAccountEmail", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> serviceAccountEmail;

    /**
     * @return Service Account to impersonate for plugin workload identity federation.
     * Required with `identity_token_audience`. Requires Vault 1.17+. *Available only for Vault Enterprise*.
     * 
     */
    public Output<Optional<String>> serviceAccountEmail() {
        return Codegen.optional(this.serviceAccountEmail);
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
    public SecretBackend(String name, @Nullable SecretBackendArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public SecretBackend(String name, @Nullable SecretBackendArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("vault:gcp/secretBackend:SecretBackend", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()));
    }

    private SecretBackend(String name, Output<String> id, @Nullable SecretBackendState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("vault:gcp/secretBackend:SecretBackend", name, state, makeResourceOptions(options, id));
    }

    private static SecretBackendArgs makeArgs(@Nullable SecretBackendArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? SecretBackendArgs.Empty : args;
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .additionalSecretOutputs(List.of(
                "credentials"
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
