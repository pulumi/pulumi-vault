// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vault.azure;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.vault.Utilities;
import com.pulumi.vault.azure.AuthBackendConfigArgs;
import com.pulumi.vault.azure.inputs.AuthBackendConfigState;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * ## Example Usage
 * 
 * You can setup the Azure auth engine with Workload Identity Federation (WIF) for a secret-less configuration:
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
 * import com.pulumi.vault.azure.AuthBackendConfig;
 * import com.pulumi.vault.azure.AuthBackendConfigArgs;
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
 *         var example = new AuthBackend("example", AuthBackendArgs.builder()
 *             .type("azure")
 *             .identityTokenKey("example-key")
 *             .build());
 * 
 *         var exampleAuthBackendConfig = new AuthBackendConfig("exampleAuthBackendConfig", AuthBackendConfigArgs.builder()
 *             .backend(example.path())
 *             .tenantId("11111111-2222-3333-4444-555555555555")
 *             .clientId("11111111-2222-3333-4444-555555555555")
 *             .identityTokenAudience("<TOKEN_AUDIENCE>")
 *             .identityTokenTtl("<TOKEN_TTL>")
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
 * import com.pulumi.vault.azure.AuthBackendConfig;
 * import com.pulumi.vault.azure.AuthBackendConfigArgs;
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
 *         var example = new AuthBackend("example", AuthBackendArgs.builder()
 *             .type("azure")
 *             .build());
 * 
 *         var exampleAuthBackendConfig = new AuthBackendConfig("exampleAuthBackendConfig", AuthBackendConfigArgs.builder()
 *             .backend(example.path())
 *             .tenantId("11111111-2222-3333-4444-555555555555")
 *             .clientId("11111111-2222-3333-4444-555555555555")
 *             .clientSecret("01234567890123456789")
 *             .resource("https://vault.hashicorp.com")
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
 * Azure auth backends can be imported using `auth/`, the `backend` path, and `/config` e.g.
 * 
 * ```sh
 * $ pulumi import vault:azure/authBackendConfig:AuthBackendConfig example auth/azure/config
 * ```
 * 
 */
@ResourceType(type="vault:azure/authBackendConfig:AuthBackendConfig")
public class AuthBackendConfig extends com.pulumi.resources.CustomResource {
    /**
     * The path the Azure auth backend being configured was
     * mounted at.  Defaults to `azure`.
     * 
     */
    @Export(name="backend", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> backend;

    /**
     * @return The path the Azure auth backend being configured was
     * mounted at.  Defaults to `azure`.
     * 
     */
    public Output<Optional<String>> backend() {
        return Codegen.optional(this.backend);
    }
    /**
     * The client id for credentials to query the Azure APIs.
     * Currently read permissions to query compute resources are required.
     * 
     */
    @Export(name="clientId", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> clientId;

    /**
     * @return The client id for credentials to query the Azure APIs.
     * Currently read permissions to query compute resources are required.
     * 
     */
    public Output<Optional<String>> clientId() {
        return Codegen.optional(this.clientId);
    }
    /**
     * The client secret for credentials to query the
     * Azure APIs.
     * 
     */
    @Export(name="clientSecret", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> clientSecret;

    /**
     * @return The client secret for credentials to query the
     * Azure APIs.
     * 
     */
    public Output<Optional<String>> clientSecret() {
        return Codegen.optional(this.clientSecret);
    }
    /**
     * Cancels all upcoming rotations of the root credential until unset. Requires Vault Enterprise 1.19+.
     * *Available only for Vault Enterprise*
     * 
     */
    @Export(name="disableAutomatedRotation", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> disableAutomatedRotation;

    /**
     * @return Cancels all upcoming rotations of the root credential until unset. Requires Vault Enterprise 1.19+.
     * *Available only for Vault Enterprise*
     * 
     */
    public Output<Optional<Boolean>> disableAutomatedRotation() {
        return Codegen.optional(this.disableAutomatedRotation);
    }
    /**
     * The Azure cloud environment. Valid values:
     * AzurePublicCloud, AzureUSGovernmentCloud, AzureChinaCloud,
     * AzureGermanCloud.  Defaults to `AzurePublicCloud`.
     * 
     */
    @Export(name="environment", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> environment;

    /**
     * @return The Azure cloud environment. Valid values:
     * AzurePublicCloud, AzureUSGovernmentCloud, AzureChinaCloud,
     * AzureGermanCloud.  Defaults to `AzurePublicCloud`.
     * 
     */
    public Output<Optional<String>> environment() {
        return Codegen.optional(this.environment);
    }
    /**
     * The audience claim value for plugin identity tokens. Requires Vault 1.17+.
     * *Available only for Vault Enterprise*
     * 
     */
    @Export(name="identityTokenAudience", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> identityTokenAudience;

    /**
     * @return The audience claim value for plugin identity tokens. Requires Vault 1.17+.
     * *Available only for Vault Enterprise*
     * 
     */
    public Output<Optional<String>> identityTokenAudience() {
        return Codegen.optional(this.identityTokenAudience);
    }
    /**
     * The TTL of generated identity tokens in seconds.
     * 
     */
    @Export(name="identityTokenTtl", refs={Integer.class}, tree="[0]")
    private Output<Integer> identityTokenTtl;

    /**
     * @return The TTL of generated identity tokens in seconds.
     * 
     */
    public Output<Integer> identityTokenTtl() {
        return this.identityTokenTtl;
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
     * The configured URL for the application registered in
     * Azure Active Directory.
     * 
     */
    @Export(name="resource", refs={String.class}, tree="[0]")
    private Output<String> resource;

    /**
     * @return The configured URL for the application registered in
     * Azure Active Directory.
     * 
     */
    public Output<String> resource() {
        return this.resource;
    }
    /**
     * The amount of time in seconds Vault should wait before rotating the root credential.
     * A zero value tells Vault not to rotate the root credential. The minimum rotation period is 10 seconds. Requires Vault Enterprise 1.19+.
     * *Available only for Vault Enterprise*
     * 
     */
    @Export(name="rotationPeriod", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> rotationPeriod;

    /**
     * @return The amount of time in seconds Vault should wait before rotating the root credential.
     * A zero value tells Vault not to rotate the root credential. The minimum rotation period is 10 seconds. Requires Vault Enterprise 1.19+.
     * *Available only for Vault Enterprise*
     * 
     */
    public Output<Optional<Integer>> rotationPeriod() {
        return Codegen.optional(this.rotationPeriod);
    }
    /**
     * The schedule, in [cron-style time format](https://en.wikipedia.org/wiki/Cron),
     * defining the schedule on which Vault should rotate the root token. Requires Vault Enterprise 1.19+.
     * *Available only for Vault Enterprise*
     * 
     */
    @Export(name="rotationSchedule", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> rotationSchedule;

    /**
     * @return The schedule, in [cron-style time format](https://en.wikipedia.org/wiki/Cron),
     * defining the schedule on which Vault should rotate the root token. Requires Vault Enterprise 1.19+.
     * *Available only for Vault Enterprise*
     * 
     */
    public Output<Optional<String>> rotationSchedule() {
        return Codegen.optional(this.rotationSchedule);
    }
    /**
     * The maximum amount of time in seconds allowed to complete
     * a rotation when a scheduled token rotation occurs. The default rotation window is
     * unbound and the minimum allowable window is `3600`. Requires Vault Enterprise 1.19+.
     * *Available only for Vault Enterprise*
     * 
     */
    @Export(name="rotationWindow", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> rotationWindow;

    /**
     * @return The maximum amount of time in seconds allowed to complete
     * a rotation when a scheduled token rotation occurs. The default rotation window is
     * unbound and the minimum allowable window is `3600`. Requires Vault Enterprise 1.19+.
     * *Available only for Vault Enterprise*
     * 
     */
    public Output<Optional<Integer>> rotationWindow() {
        return Codegen.optional(this.rotationWindow);
    }
    /**
     * The tenant id for the Azure Active Directory
     * organization.
     * 
     */
    @Export(name="tenantId", refs={String.class}, tree="[0]")
    private Output<String> tenantId;

    /**
     * @return The tenant id for the Azure Active Directory
     * organization.
     * 
     */
    public Output<String> tenantId() {
        return this.tenantId;
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
        super("vault:azure/authBackendConfig:AuthBackendConfig", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private AuthBackendConfig(java.lang.String name, Output<java.lang.String> id, @Nullable AuthBackendConfigState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("vault:azure/authBackendConfig:AuthBackendConfig", name, state, makeResourceOptions(options, id), false);
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
                "clientId",
                "clientSecret",
                "tenantId"
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
