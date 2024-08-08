// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vault.azure;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.vault.Utilities;
import com.pulumi.vault.azure.BackendArgs;
import com.pulumi.vault.azure.inputs.BackendState;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * ## Example Usage
 * 
 * ### *Vault-1.9 And Above*
 * 
 * You can setup the Azure secrets engine with Workload Identity Federation (WIF) for a secret-less configuration:
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.vault.azure.Backend;
 * import com.pulumi.vault.azure.BackendArgs;
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
 *         var azure = new Backend("azure", BackendArgs.builder()
 *             .subscriptionId("11111111-2222-3333-4444-111111111111")
 *             .tenantId("11111111-2222-3333-4444-222222222222")
 *             .clientId("11111111-2222-3333-4444-333333333333")
 *             .identityTokenAudience("<TOKEN_AUDIENCE>")
 *             .identityTokenTtl("<TOKEN_TTL>")
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
 * import com.pulumi.vault.azure.Backend;
 * import com.pulumi.vault.azure.BackendArgs;
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
 *         var azure = new Backend("azure", BackendArgs.builder()
 *             .useMicrosoftGraphApi(true)
 *             .subscriptionId("11111111-2222-3333-4444-111111111111")
 *             .tenantId("11111111-2222-3333-4444-222222222222")
 *             .clientId("11111111-2222-3333-4444-333333333333")
 *             .clientSecret("12345678901234567890")
 *             .environment("AzurePublicCloud")
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ### *Vault-1.8 And Below*
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.vault.azure.Backend;
 * import com.pulumi.vault.azure.BackendArgs;
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
 *         var azure = new Backend("azure", BackendArgs.builder()
 *             .useMicrosoftGraphApi(false)
 *             .subscriptionId("11111111-2222-3333-4444-111111111111")
 *             .tenantId("11111111-2222-3333-4444-222222222222")
 *             .clientId("11111111-2222-3333-4444-333333333333")
 *             .clientSecret("12345678901234567890")
 *             .environment("AzurePublicCloud")
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 */
@ResourceType(type="vault:azure/backend:Backend")
public class Backend extends com.pulumi.resources.CustomResource {
    /**
     * The OAuth2 client id to connect to Azure.
     * 
     */
    @Export(name="clientId", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> clientId;

    /**
     * @return The OAuth2 client id to connect to Azure.
     * 
     */
    public Output<Optional<String>> clientId() {
        return Codegen.optional(this.clientId);
    }
    /**
     * The OAuth2 client secret to connect to Azure.
     * 
     */
    @Export(name="clientSecret", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> clientSecret;

    /**
     * @return The OAuth2 client secret to connect to Azure.
     * 
     */
    public Output<Optional<String>> clientSecret() {
        return Codegen.optional(this.clientSecret);
    }
    /**
     * Human-friendly description of the mount for the backend.
     * 
     */
    @Export(name="description", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> description;

    /**
     * @return Human-friendly description of the mount for the backend.
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
     * The Azure environment.
     * 
     */
    @Export(name="environment", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> environment;

    /**
     * @return The Azure environment.
     * 
     */
    public Output<Optional<String>> environment() {
        return Codegen.optional(this.environment);
    }
    /**
     * The audience claim value. Requires Vault 1.17+.
     * *Available only for Vault Enterprise*
     * 
     */
    @Export(name="identityTokenAudience", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> identityTokenAudience;

    /**
     * @return The audience claim value. Requires Vault 1.17+.
     * *Available only for Vault Enterprise*
     * 
     */
    public Output<Optional<String>> identityTokenAudience() {
        return Codegen.optional(this.identityTokenAudience);
    }
    /**
     * The key to use for signing identity tokens. Requires Vault 1.17+.
     * *Available only for Vault Enterprise*
     * 
     */
    @Export(name="identityTokenKey", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> identityTokenKey;

    /**
     * @return The key to use for signing identity tokens. Requires Vault 1.17+.
     * *Available only for Vault Enterprise*
     * 
     */
    public Output<Optional<String>> identityTokenKey() {
        return Codegen.optional(this.identityTokenKey);
    }
    /**
     * The TTL of generated identity tokens in seconds. Requires Vault 1.17+.
     * *Available only for Vault Enterprise*
     * 
     */
    @Export(name="identityTokenTtl", refs={Integer.class}, tree="[0]")
    private Output<Integer> identityTokenTtl;

    /**
     * @return The TTL of generated identity tokens in seconds. Requires Vault 1.17+.
     * *Available only for Vault Enterprise*
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
     * The unique path this backend should be mounted at. Defaults to `azure`.
     * 
     */
    @Export(name="path", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> path;

    /**
     * @return The unique path this backend should be mounted at. Defaults to `azure`.
     * 
     */
    public Output<Optional<String>> path() {
        return Codegen.optional(this.path);
    }
    /**
     * The subscription id for the Azure Active Directory.
     * 
     */
    @Export(name="subscriptionId", refs={String.class}, tree="[0]")
    private Output<String> subscriptionId;

    /**
     * @return The subscription id for the Azure Active Directory.
     * 
     */
    public Output<String> subscriptionId() {
        return this.subscriptionId;
    }
    /**
     * The tenant id for the Azure Active Directory.
     * 
     */
    @Export(name="tenantId", refs={String.class}, tree="[0]")
    private Output<String> tenantId;

    /**
     * @return The tenant id for the Azure Active Directory.
     * 
     */
    public Output<String> tenantId() {
        return this.tenantId;
    }
    /**
     * Use the Microsoft Graph API. Should be set to true on vault-1.10+
     * 
     * @deprecated
     * This field is not supported in Vault-1.12+ and is the default behavior. This field will be removed in future version of the provider.
     * 
     */
    @Deprecated /* This field is not supported in Vault-1.12+ and is the default behavior. This field will be removed in future version of the provider. */
    @Export(name="useMicrosoftGraphApi", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> useMicrosoftGraphApi;

    /**
     * @return Use the Microsoft Graph API. Should be set to true on vault-1.10+
     * 
     */
    public Output<Boolean> useMicrosoftGraphApi() {
        return this.useMicrosoftGraphApi;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Backend(String name) {
        this(name, BackendArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Backend(String name, BackendArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Backend(String name, BackendArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("vault:azure/backend:Backend", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()));
    }

    private Backend(String name, Output<String> id, @Nullable BackendState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("vault:azure/backend:Backend", name, state, makeResourceOptions(options, id));
    }

    private static BackendArgs makeArgs(BackendArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? BackendArgs.Empty : args;
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .additionalSecretOutputs(List.of(
                "clientId",
                "clientSecret",
                "subscriptionId",
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
    public static Backend get(String name, Output<String> id, @Nullable BackendState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Backend(name, id, state, options);
    }
}
