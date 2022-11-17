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
 * ```java
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
 * public class App {
 *     public static void main(String[] args) {
 *         Pulumi.run(App::stack);
 *     }
 * 
 *     public static void stack(Context ctx) {
 *         final var config = ctx.config();
 *         final var settingsFile = config.get(&#34;settingsFile&#34;);
 *         var userpass = new AuthBackend(&#34;userpass&#34;, AuthBackendArgs.builder()        
 *             .type(&#34;userpass&#34;)
 *             .path(&#34;userpass&#34;)
 *             .build());
 * 
 *         var myPingid = new MfaPingid(&#34;myPingid&#34;, MfaPingidArgs.builder()        
 *             .mountAccessor(userpass.accessor())
 *             .usernameFormat(&#34;user@example.com&#34;)
 *             .settingsFileBase64(settingsFile)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * Mounts can be imported using the `path`, e.g.
 * 
 * ```sh
 *  $ pulumi import vault:index/mfaPingid:MfaPingid my_pingid my_pingid
 * ```
 * 
 */
@ResourceType(type="vault:index/mfaPingid:MfaPingid")
public class MfaPingid extends com.pulumi.resources.CustomResource {
    /**
     * Admin URL computed by Vault.
     * 
     */
    @Export(name="adminUrl", type=String.class, parameters={})
    private Output<String> adminUrl;

    /**
     * @return Admin URL computed by Vault.
     * 
     */
    public Output<String> adminUrl() {
        return this.adminUrl;
    }
    /**
     * Authenticator URL computed by Vault.
     * 
     */
    @Export(name="authenticatorUrl", type=String.class, parameters={})
    private Output<String> authenticatorUrl;

    /**
     * @return Authenticator URL computed by Vault.
     * 
     */
    public Output<String> authenticatorUrl() {
        return this.authenticatorUrl;
    }
    /**
     * IDP URL computed by Vault.
     * 
     */
    @Export(name="idpUrl", type=String.class, parameters={})
    private Output<String> idpUrl;

    /**
     * @return IDP URL computed by Vault.
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
    @Export(name="mountAccessor", type=String.class, parameters={})
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
    @Export(name="name", type=String.class, parameters={})
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
     * The `namespace` is always relative to the provider&#39;s configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
     * *Available only for Vault Enterprise*.
     * 
     */
    @Export(name="namespace", type=String.class, parameters={})
    private Output</* @Nullable */ String> namespace;

    /**
     * @return The namespace to provision the resource in.
     * The value should not contain leading or trailing forward slashes.
     * The `namespace` is always relative to the provider&#39;s configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
     * *Available only for Vault Enterprise*.
     * 
     */
    public Output<Optional<String>> namespace() {
        return Codegen.optional(this.namespace);
    }
    /**
     * Namespace ID computed by Vault.
     * 
     */
    @Export(name="namespaceId", type=String.class, parameters={})
    private Output<String> namespaceId;

    /**
     * @return Namespace ID computed by Vault.
     * 
     */
    public Output<String> namespaceId() {
        return this.namespaceId;
    }
    /**
     * Org Alias computed by Vault.
     * 
     */
    @Export(name="orgAlias", type=String.class, parameters={})
    private Output<String> orgAlias;

    /**
     * @return Org Alias computed by Vault.
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
    @Export(name="settingsFileBase64", type=String.class, parameters={})
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
     * Type of configuration computed by Vault.
     * 
     */
    @Export(name="type", type=String.class, parameters={})
    private Output<String> type;

    /**
     * @return Type of configuration computed by Vault.
     * 
     */
    public Output<String> type() {
        return this.type;
    }
    /**
     * If set, enables use of PingID signature. Computed by Vault
     * 
     */
    @Export(name="useSignature", type=Boolean.class, parameters={})
    private Output<Boolean> useSignature;

    /**
     * @return If set, enables use of PingID signature. Computed by Vault
     * 
     */
    public Output<Boolean> useSignature() {
        return this.useSignature;
    }
    /**
     * `(string)` - A format string for mapping Identity names to MFA method names.
     * Values to substitute should be placed in `{{}}`. For example, `&#34;{{alias.name}}@example.com&#34;`.
     * If blank, the Alias&#39;s Name field will be used as-is. Currently-supported mappings:
     * - alias.name: The name returned by the mount configured via the `mount_accessor` parameter
     * - entity.name: The name configured for the Entity
     * - alias.metadata.`&lt;key&gt;`: The value of the Alias&#39;s metadata parameter
     * - entity.metadata.`&lt;key&gt;`: The value of the Entity&#39;s metadata parameter
     * 
     */
    @Export(name="usernameFormat", type=String.class, parameters={})
    private Output</* @Nullable */ String> usernameFormat;

    /**
     * @return `(string)` - A format string for mapping Identity names to MFA method names.
     * Values to substitute should be placed in `{{}}`. For example, `&#34;{{alias.name}}@example.com&#34;`.
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
    public MfaPingid(String name) {
        this(name, MfaPingidArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public MfaPingid(String name, MfaPingidArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public MfaPingid(String name, MfaPingidArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("vault:index/mfaPingid:MfaPingid", name, args == null ? MfaPingidArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private MfaPingid(String name, Output<String> id, @Nullable MfaPingidState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("vault:index/mfaPingid:MfaPingid", name, state, makeResourceOptions(options, id));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
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
    public static MfaPingid get(String name, Output<String> id, @Nullable MfaPingidState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new MfaPingid(name, id, state, options);
    }
}