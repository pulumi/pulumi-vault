// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vault;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.vault.MfaOktaArgs;
import com.pulumi.vault.Utilities;
import com.pulumi.vault.inputs.MfaOktaState;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides a resource to manage [Okta MFA](https://www.vaultproject.io/docs/enterprise/mfa/mfa-okta).
 * 
 * **Note** this feature is available only with Vault Enterprise.
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
 * import com.pulumi.vault.MfaOkta;
 * import com.pulumi.vault.MfaOktaArgs;
 * import java.util.List;
 * import java.util.ArrayList;
 * import java.util.Map;
 * import java.io.File;
 * import java.nio.file.Files;
 * import java.nio.file.Paths;
 * 
 * public class App }{{@code
 *     public static void main(String[] args) }{{@code
 *         Pulumi.run(App::stack);
 *     }}{@code
 * 
 *     public static void stack(Context ctx) }{{@code
 *         var userpass = new AuthBackend("userpass", AuthBackendArgs.builder()
 *             .type("userpass")
 *             .path("userpass")
 *             .build());
 * 
 *         var myOkta = new MfaOkta("myOkta", MfaOktaArgs.builder()
 *             .name("my_okta")
 *             .mountAccessor(userpass.accessor())
 *             .usernameFormat("user}{@literal @}{@code example.com")
 *             .orgName("hashicorp")
 *             .apiToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9")
 *             .build());
 * 
 *     }}{@code
 * }}{@code
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Import
 * 
 * Mounts can be imported using the `path`, e.g.
 * 
 * ```sh
 * $ pulumi import vault:index/mfaOkta:MfaOkta my_okta my_okta
 * ```
 * 
 */
@ResourceType(type="vault:index/mfaOkta:MfaOkta")
public class MfaOkta extends com.pulumi.resources.CustomResource {
    /**
     * `(string: &lt;required&gt;)` - Okta API key.
     * 
     */
    @Export(name="apiToken", refs={String.class}, tree="[0]")
    private Output<String> apiToken;

    /**
     * @return `(string: &lt;required&gt;)` - Okta API key.
     * 
     */
    public Output<String> apiToken() {
        return this.apiToken;
    }
    /**
     * `(string)` - If set, will be used as the base domain for API requests. Examples are `okta.com`,
     * `oktapreview.com`, and `okta-emea.com`.
     * 
     */
    @Export(name="baseUrl", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> baseUrl;

    /**
     * @return `(string)` - If set, will be used as the base domain for API requests. Examples are `okta.com`,
     * `oktapreview.com`, and `okta-emea.com`.
     * 
     */
    public Output<Optional<String>> baseUrl() {
        return Codegen.optional(this.baseUrl);
    }
    /**
     * `(string: &lt;required&gt;)` - The mount to tie this method to for use in automatic mappings.
     * The mapping will use the Name field of Aliases associated with this mount as the username in the mapping.
     * 
     */
    @Export(name="mountAccessor", refs={String.class}, tree="[0]")
    private Output<String> mountAccessor;

    /**
     * @return `(string: &lt;required&gt;)` - The mount to tie this method to for use in automatic mappings.
     * The mapping will use the Name field of Aliases associated with this mount as the username in the mapping.
     * 
     */
    public Output<String> mountAccessor() {
        return this.mountAccessor;
    }
    /**
     * `(string: &lt;required&gt;)` – Name of the MFA method.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return `(string: &lt;required&gt;)` – Name of the MFA method.
     * 
     */
    public Output<String> name() {
        return this.name;
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
     * `(string: &lt;required&gt;)` - Name of the organization to be used in the Okta API.
     * 
     */
    @Export(name="orgName", refs={String.class}, tree="[0]")
    private Output<String> orgName;

    /**
     * @return `(string: &lt;required&gt;)` - Name of the organization to be used in the Okta API.
     * 
     */
    public Output<String> orgName() {
        return this.orgName;
    }
    /**
     * `(string: &lt;required&gt;)` - If set to true, the username will only match the
     * primary email for the account.
     * 
     */
    @Export(name="primaryEmail", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> primaryEmail;

    /**
     * @return `(string: &lt;required&gt;)` - If set to true, the username will only match the
     * primary email for the account.
     * 
     */
    public Output<Optional<Boolean>> primaryEmail() {
        return Codegen.optional(this.primaryEmail);
    }
    /**
     * `(string)` - A format string for mapping Identity names to MFA method names.
     * Values to substitute should be placed in `{{}}`. For example, `&#34;{{alias.name}}{@literal @}example.com&#34;`.
     * If blank, the Alias&#39;s Name field will be used as-is. Currently-supported mappings:
     * - alias.name: The name returned by the mount configured via the `mount_accessor` parameter
     * - entity.name: The name configured for the Entity
     * - alias.metadata.`&lt;key&gt;`: The value of the Alias&#39;s metadata parameter
     * - entity.metadata.`&lt;key&gt;`: The value of the Entity&#39;s metadata parameter
     * 
     */
    @Export(name="usernameFormat", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> usernameFormat;

    /**
     * @return `(string)` - A format string for mapping Identity names to MFA method names.
     * Values to substitute should be placed in `{{}}`. For example, `&#34;{{alias.name}}{@literal @}example.com&#34;`.
     * If blank, the Alias&#39;s Name field will be used as-is. Currently-supported mappings:
     * - alias.name: The name returned by the mount configured via the `mount_accessor` parameter
     * - entity.name: The name configured for the Entity
     * - alias.metadata.`&lt;key&gt;`: The value of the Alias&#39;s metadata parameter
     * - entity.metadata.`&lt;key&gt;`: The value of the Entity&#39;s metadata parameter
     * 
     */
    public Output<Optional<String>> usernameFormat() {
        return Codegen.optional(this.usernameFormat);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public MfaOkta(java.lang.String name) {
        this(name, MfaOktaArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public MfaOkta(java.lang.String name, MfaOktaArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public MfaOkta(java.lang.String name, MfaOktaArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("vault:index/mfaOkta:MfaOkta", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private MfaOkta(java.lang.String name, Output<java.lang.String> id, @Nullable MfaOktaState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("vault:index/mfaOkta:MfaOkta", name, state, makeResourceOptions(options, id), false);
    }

    private static MfaOktaArgs makeArgs(MfaOktaArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? MfaOktaArgs.Empty : args;
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<java.lang.String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .additionalSecretOutputs(List.of(
                "apiToken"
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
    public static MfaOkta get(java.lang.String name, Output<java.lang.String> id, @Nullable MfaOktaState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new MfaOkta(name, id, state, options);
    }
}
