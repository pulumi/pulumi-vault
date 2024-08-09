// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vault.gcp;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.vault.Utilities;
import com.pulumi.vault.gcp.SecretImpersonatedAccountArgs;
import com.pulumi.vault.gcp.inputs.SecretImpersonatedAccountState;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Creates a Impersonated Account in the [GCP Secrets Engine](https://www.vaultproject.io/docs/secrets/gcp/index.html) for Vault.
 * 
 * Each [impersonated account](https://www.vaultproject.io/docs/secrets/gcp/index.html#impersonated-accounts) is tied to a separately managed
 * Service Account.
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
 * import com.pulumi.google.serviceAccount;
 * import com.pulumi.google.ServiceAccountArgs;
 * import com.pulumi.vault.gcp.SecretBackend;
 * import com.pulumi.vault.gcp.SecretBackendArgs;
 * import com.pulumi.vault.gcp.SecretImpersonatedAccount;
 * import com.pulumi.vault.gcp.SecretImpersonatedAccountArgs;
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
 *         var this_ = new ServiceAccount("this", ServiceAccountArgs.builder()
 *             .accountId("my-awesome-account")
 *             .build());
 * 
 *         var gcp = new SecretBackend("gcp", SecretBackendArgs.builder()
 *             .path("gcp")
 *             .credentials(StdFunctions.file(FileArgs.builder()
 *                 .input("credentials.json")
 *                 .build()).result())
 *             .build());
 * 
 *         var impersonatedAccount = new SecretImpersonatedAccount("impersonatedAccount", SecretImpersonatedAccountArgs.builder()
 *             .backend(gcp.path())
 *             .impersonatedAccount("this")
 *             .serviceAccountEmail(this_.email())
 *             .tokenScopes("https://www.googleapis.com/auth/cloud-platform")
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
 * A impersonated account can be imported using its Vault Path. For example, referencing the example above,
 * 
 * ```sh
 * $ pulumi import vault:gcp/secretImpersonatedAccount:SecretImpersonatedAccount impersonated_account gcp/impersonated-account/project_viewer
 * ```
 * 
 */
@ResourceType(type="vault:gcp/secretImpersonatedAccount:SecretImpersonatedAccount")
public class SecretImpersonatedAccount extends com.pulumi.resources.CustomResource {
    /**
     * Path where the GCP Secrets Engine is mounted
     * 
     */
    @Export(name="backend", refs={String.class}, tree="[0]")
    private Output<String> backend;

    /**
     * @return Path where the GCP Secrets Engine is mounted
     * 
     */
    public Output<String> backend() {
        return this.backend;
    }
    /**
     * Name of the Impersonated Account to create
     * 
     */
    @Export(name="impersonatedAccount", refs={String.class}, tree="[0]")
    private Output<String> impersonatedAccount;

    /**
     * @return Name of the Impersonated Account to create
     * 
     */
    public Output<String> impersonatedAccount() {
        return this.impersonatedAccount;
    }
    /**
     * Target namespace. (requires Enterprise)
     * 
     */
    @Export(name="namespace", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> namespace;

    /**
     * @return Target namespace. (requires Enterprise)
     * 
     */
    public Output<Optional<String>> namespace() {
        return Codegen.optional(this.namespace);
    }
    /**
     * Email of the GCP service account to impersonate.
     * 
     */
    @Export(name="serviceAccountEmail", refs={String.class}, tree="[0]")
    private Output<String> serviceAccountEmail;

    /**
     * @return Email of the GCP service account to impersonate.
     * 
     */
    public Output<String> serviceAccountEmail() {
        return this.serviceAccountEmail;
    }
    /**
     * Project the service account belongs to.
     * 
     */
    @Export(name="serviceAccountProject", refs={String.class}, tree="[0]")
    private Output<String> serviceAccountProject;

    /**
     * @return Project the service account belongs to.
     * 
     */
    public Output<String> serviceAccountProject() {
        return this.serviceAccountProject;
    }
    /**
     * List of OAuth scopes to assign to access tokens generated under this impersonated account.
     * 
     */
    @Export(name="tokenScopes", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> tokenScopes;

    /**
     * @return List of OAuth scopes to assign to access tokens generated under this impersonated account.
     * 
     */
    public Output<Optional<List<String>>> tokenScopes() {
        return Codegen.optional(this.tokenScopes);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public SecretImpersonatedAccount(java.lang.String name) {
        this(name, SecretImpersonatedAccountArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public SecretImpersonatedAccount(java.lang.String name, SecretImpersonatedAccountArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public SecretImpersonatedAccount(java.lang.String name, SecretImpersonatedAccountArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("vault:gcp/secretImpersonatedAccount:SecretImpersonatedAccount", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private SecretImpersonatedAccount(java.lang.String name, Output<java.lang.String> id, @Nullable SecretImpersonatedAccountState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("vault:gcp/secretImpersonatedAccount:SecretImpersonatedAccount", name, state, makeResourceOptions(options, id), false);
    }

    private static SecretImpersonatedAccountArgs makeArgs(SecretImpersonatedAccountArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? SecretImpersonatedAccountArgs.Empty : args;
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
    public static SecretImpersonatedAccount get(java.lang.String name, Output<java.lang.String> id, @Nullable SecretImpersonatedAccountState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new SecretImpersonatedAccount(name, id, state, options);
    }
}
