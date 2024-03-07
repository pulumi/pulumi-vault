// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vault;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.vault.MfaDuoArgs;
import com.pulumi.vault.Utilities;
import com.pulumi.vault.inputs.MfaDuoState;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides a resource to manage [Duo MFA](https://www.vaultproject.io/docs/enterprise/mfa/mfa-duo.html).
 * 
 * **Note** this feature is available only with Vault Enterprise.
 * 
 * ## Example Usage
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.vault.AuthBackend;
 * import com.pulumi.vault.AuthBackendArgs;
 * import com.pulumi.vault.MfaDuo;
 * import com.pulumi.vault.MfaDuoArgs;
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
 *         var userpass = new AuthBackend(&#34;userpass&#34;, AuthBackendArgs.builder()        
 *             .type(&#34;userpass&#34;)
 *             .path(&#34;userpass&#34;)
 *             .build());
 * 
 *         var myDuo = new MfaDuo(&#34;myDuo&#34;, MfaDuoArgs.builder()        
 *             .mountAccessor(userpass.accessor())
 *             .secretKey(&#34;8C7THtrIigh2rPZQMbguugt8IUftWhMRCOBzbuyz&#34;)
 *             .integrationKey(&#34;BIACEUEAXI20BNWTEYXT&#34;)
 *             .apiHostname(&#34;api-2b5c39f5.duosecurity.com&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Import
 * 
 * Mounts can be imported using the `path`, e.g.
 * 
 * ```sh
 * $ pulumi import vault:index/mfaDuo:MfaDuo my_duo my_duo
 * ```
 * 
 */
@ResourceType(type="vault:index/mfaDuo:MfaDuo")
public class MfaDuo extends com.pulumi.resources.CustomResource {
    /**
     * `(string: &lt;required&gt;)` - API hostname for Duo.
     * 
     */
    @Export(name="apiHostname", refs={String.class}, tree="[0]")
    private Output<String> apiHostname;

    /**
     * @return `(string: &lt;required&gt;)` - API hostname for Duo.
     * 
     */
    public Output<String> apiHostname() {
        return this.apiHostname;
    }
    /**
     * `(string: &lt;required&gt;)` - Integration key for Duo.
     * 
     */
    @Export(name="integrationKey", refs={String.class}, tree="[0]")
    private Output<String> integrationKey;

    /**
     * @return `(string: &lt;required&gt;)` - Integration key for Duo.
     * 
     */
    public Output<String> integrationKey() {
        return this.integrationKey;
    }
    /**
     * `(string: &lt;required&gt;)` - The mount to tie this method to for use in automatic mappings. The mapping will use the Name field of Aliases associated with this mount as the username in the mapping.
     * 
     */
    @Export(name="mountAccessor", refs={String.class}, tree="[0]")
    private Output<String> mountAccessor;

    /**
     * @return `(string: &lt;required&gt;)` - The mount to tie this method to for use in automatic mappings. The mapping will use the Name field of Aliases associated with this mount as the username in the mapping.
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
     * The `namespace` is always relative to the provider&#39;s configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
     * *Available only for Vault Enterprise*.
     * 
     */
    @Export(name="namespace", refs={String.class}, tree="[0]")
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
     * `(string)` - Push information for Duo.
     * 
     */
    @Export(name="pushInfo", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> pushInfo;

    /**
     * @return `(string)` - Push information for Duo.
     * 
     */
    public Output<Optional<String>> pushInfo() {
        return Codegen.optional(this.pushInfo);
    }
    /**
     * `(string: &lt;required&gt;)` - Secret key for Duo.
     * 
     */
    @Export(name="secretKey", refs={String.class}, tree="[0]")
    private Output<String> secretKey;

    /**
     * @return `(string: &lt;required&gt;)` - Secret key for Duo.
     * 
     */
    public Output<String> secretKey() {
        return this.secretKey;
    }
    /**
     * `(string)` - A format string for mapping Identity names to MFA method names. Values to substitute should be placed in `{{}}`. For example, `&#34;{{alias.name}}@example.com&#34;`. If blank, the Alias&#39;s Name field will be used as-is. Currently-supported mappings:
     * - alias.name: The name returned by the mount configured via the `mount_accessor` parameter
     * - entity.name: The name configured for the Entity
     * - alias.metadata.`&lt;key&gt;`: The value of the Alias&#39;s metadata parameter
     * - entity.metadata.`&lt;key&gt;`: The value of the Entity&#39;s metadata parameter
     * 
     */
    @Export(name="usernameFormat", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> usernameFormat;

    /**
     * @return `(string)` - A format string for mapping Identity names to MFA method names. Values to substitute should be placed in `{{}}`. For example, `&#34;{{alias.name}}@example.com&#34;`. If blank, the Alias&#39;s Name field will be used as-is. Currently-supported mappings:
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
    public MfaDuo(String name) {
        this(name, MfaDuoArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public MfaDuo(String name, MfaDuoArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public MfaDuo(String name, MfaDuoArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("vault:index/mfaDuo:MfaDuo", name, args == null ? MfaDuoArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private MfaDuo(String name, Output<String> id, @Nullable MfaDuoState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("vault:index/mfaDuo:MfaDuo", name, state, makeResourceOptions(options, id));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .additionalSecretOutputs(List.of(
                "integrationKey",
                "secretKey"
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
    public static MfaDuo get(String name, Output<String> id, @Nullable MfaDuoState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new MfaDuo(name, id, state, options);
    }
}
