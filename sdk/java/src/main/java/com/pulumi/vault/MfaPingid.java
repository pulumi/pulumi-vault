// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vault;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.vault.MfaPingidArgs;
import com.pulumi.vault.Utilities;
import com.pulumi.vault.inputs.MfaPingidState;
import java.lang.Boolean;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides a resource to manage [PingID MFA](https://www.vaultproject.io/docs/enterprise/mfa/mfa-pingid).
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
 * import com.pulumi.vault.MfaPingid;
 * import com.pulumi.vault.MfaPingidArgs;
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
 *         final var config = ctx.config();
 *         final var settingsFile = config.get("settingsFile");
 *         var userpass = new AuthBackend("userpass", AuthBackendArgs.builder()
 *             .type("userpass")
 *             .path("userpass")
 *             .build());
 * 
 *         var myPingid = new MfaPingid("myPingid", MfaPingidArgs.builder()
 *             .name("my_pingid")
 *             .mountAccessor(userpass.accessor())
 *             .usernameFormat("user}{@literal @}{@code example.com")
 *             .settingsFileBase64(settingsFile)
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
 * $ pulumi import vault:index/mfaPingid:MfaPingid my_pingid my_pingid
 * ```
 * 
 */
@ResourceType(type="vault:index/mfaPingid:MfaPingid")
public class MfaPingid extends com.pulumi.resources.CustomResource {
    /**
     * `(string)` – Admin URL computed by Vault
     * 
     */
    @Export(name="adminUrl", refs={String.class}, tree="[0]")
    private Output<String> adminUrl;

    /**
     * @return `(string)` – Admin URL computed by Vault
     * 
     */
    public Output<String> adminUrl() {
        return this.adminUrl;
    }
    /**
     * `(string)` – Authenticator URL computed by Vault
     * 
     */
    @Export(name="authenticatorUrl", refs={String.class}, tree="[0]")
    private Output<String> authenticatorUrl;

    /**
     * @return `(string)` – Authenticator URL computed by Vault
     * 
     */
    public Output<String> authenticatorUrl() {
        return this.authenticatorUrl;
    }
    /**
     * `(string)` – IDP URL computed by Vault
     * 
     */
    @Export(name="idpUrl", refs={String.class}, tree="[0]")
    private Output<String> idpUrl;

    /**
     * @return `(string)` – IDP URL computed by Vault
     * 
     */
    public Output<String> idpUrl() {
        return this.idpUrl;
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
     * `(string)` – Namespace ID computed by Vault
     * 
     */
    @Export(name="namespaceId", refs={String.class}, tree="[0]")
    private Output<String> namespaceId;

    /**
     * @return `(string)` – Namespace ID computed by Vault
     * 
     */
    public Output<String> namespaceId() {
        return this.namespaceId;
    }
    /**
     * `(string)` – Org Alias computed by Vault
     * 
     */
    @Export(name="orgAlias", refs={String.class}, tree="[0]")
    private Output<String> orgAlias;

    /**
     * @return `(string)` – Org Alias computed by Vault
     * 
     */
    public Output<String> orgAlias() {
        return this.orgAlias;
    }
    /**
     * `(string: &lt;required&gt;)` - A base64-encoded third-party settings file retrieved
     * from PingID&#39;s configuration page.
     * 
     */
    @Export(name="settingsFileBase64", refs={String.class}, tree="[0]")
    private Output<String> settingsFileBase64;

    /**
     * @return `(string: &lt;required&gt;)` - A base64-encoded third-party settings file retrieved
     * from PingID&#39;s configuration page.
     * 
     */
    public Output<String> settingsFileBase64() {
        return this.settingsFileBase64;
    }
    /**
     * `(string)` – Type of configuration computed by Vault
     * 
     */
    @Export(name="type", refs={String.class}, tree="[0]")
    private Output<String> type;

    /**
     * @return `(string)` – Type of configuration computed by Vault
     * 
     */
    public Output<String> type() {
        return this.type;
    }
    /**
     * `(string)` – If set to true, enables use of PingID signature. Computed by Vault
     * 
     */
    @Export(name="useSignature", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> useSignature;

    /**
     * @return `(string)` – If set to true, enables use of PingID signature. Computed by Vault
     * 
     */
    public Output<Boolean> useSignature() {
        return this.useSignature;
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
    public MfaPingid(java.lang.String name) {
        this(name, MfaPingidArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public MfaPingid(java.lang.String name, MfaPingidArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public MfaPingid(java.lang.String name, MfaPingidArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("vault:index/mfaPingid:MfaPingid", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private MfaPingid(java.lang.String name, Output<java.lang.String> id, @Nullable MfaPingidState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("vault:index/mfaPingid:MfaPingid", name, state, makeResourceOptions(options, id), false);
    }

    private static MfaPingidArgs makeArgs(MfaPingidArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? MfaPingidArgs.Empty : args;
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
    public static MfaPingid get(java.lang.String name, Output<java.lang.String> id, @Nullable MfaPingidState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new MfaPingid(name, id, state, options);
    }
}
