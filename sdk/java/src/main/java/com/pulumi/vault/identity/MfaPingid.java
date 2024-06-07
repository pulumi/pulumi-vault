// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vault.identity;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.vault.Utilities;
import com.pulumi.vault.identity.MfaPingidArgs;
import com.pulumi.vault.identity.inputs.MfaPingidState;
import java.lang.Boolean;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Resource for configuring the pingid MFA method.
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
 * import com.pulumi.vault.identity.MfaPingid;
 * import com.pulumi.vault.identity.MfaPingidArgs;
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
 *         var example = new MfaPingid("example", MfaPingidArgs.builder()
 *             .settingsFileBase64("CnVzZV9iYXNlNjR[...]HBtCg==")
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
 * Resource can be imported using its `uuid` field, e.g.
 * 
 * ```sh
 * $ pulumi import vault:identity/mfaPingid:MfaPingid example 0d89c36a-4ff5-4d70-8749-bb6a5598aeec
 * ```
 * 
 */
@ResourceType(type="vault:identity/mfaPingid:MfaPingid")
public class MfaPingid extends com.pulumi.resources.CustomResource {
    /**
     * The admin URL, derived from &#34;settings_file_base64&#34;
     * 
     */
    @Export(name="adminUrl", refs={String.class}, tree="[0]")
    private Output<String> adminUrl;

    /**
     * @return The admin URL, derived from &#34;settings_file_base64&#34;
     * 
     */
    public Output<String> adminUrl() {
        return this.adminUrl;
    }
    /**
     * A unique identifier of the organization, derived from &#34;settings_file_base64&#34;
     * 
     */
    @Export(name="authenticatorUrl", refs={String.class}, tree="[0]")
    private Output<String> authenticatorUrl;

    /**
     * @return A unique identifier of the organization, derived from &#34;settings_file_base64&#34;
     * 
     */
    public Output<String> authenticatorUrl() {
        return this.authenticatorUrl;
    }
    /**
     * The IDP URL, derived from &#34;settings_file_base64&#34;
     * 
     */
    @Export(name="idpUrl", refs={String.class}, tree="[0]")
    private Output<String> idpUrl;

    /**
     * @return The IDP URL, derived from &#34;settings_file_base64&#34;
     * 
     */
    public Output<String> idpUrl() {
        return this.idpUrl;
    }
    /**
     * Method ID.
     * 
     */
    @Export(name="methodId", refs={String.class}, tree="[0]")
    private Output<String> methodId;

    /**
     * @return Method ID.
     * 
     */
    public Output<String> methodId() {
        return this.methodId;
    }
    /**
     * Mount accessor.
     * 
     */
    @Export(name="mountAccessor", refs={String.class}, tree="[0]")
    private Output<String> mountAccessor;

    /**
     * @return Mount accessor.
     * 
     */
    public Output<String> mountAccessor() {
        return this.mountAccessor;
    }
    /**
     * Method name.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return Method name.
     * 
     */
    public Output<String> name() {
        return this.name;
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
     * Method&#39;s namespace ID.
     * 
     */
    @Export(name="namespaceId", refs={String.class}, tree="[0]")
    private Output<String> namespaceId;

    /**
     * @return Method&#39;s namespace ID.
     * 
     */
    public Output<String> namespaceId() {
        return this.namespaceId;
    }
    /**
     * Method&#39;s namespace path.
     * 
     */
    @Export(name="namespacePath", refs={String.class}, tree="[0]")
    private Output<String> namespacePath;

    /**
     * @return Method&#39;s namespace path.
     * 
     */
    public Output<String> namespacePath() {
        return this.namespacePath;
    }
    /**
     * The name of the PingID client organization, derived from &#34;settings_file_base64&#34;
     * 
     */
    @Export(name="orgAlias", refs={String.class}, tree="[0]")
    private Output<String> orgAlias;

    /**
     * @return The name of the PingID client organization, derived from &#34;settings_file_base64&#34;
     * 
     */
    public Output<String> orgAlias() {
        return this.orgAlias;
    }
    /**
     * A base64-encoded third-party settings contents as retrieved from PingID&#39;s configuration page.
     * 
     */
    @Export(name="settingsFileBase64", refs={String.class}, tree="[0]")
    private Output<String> settingsFileBase64;

    /**
     * @return A base64-encoded third-party settings contents as retrieved from PingID&#39;s configuration page.
     * 
     */
    public Output<String> settingsFileBase64() {
        return this.settingsFileBase64;
    }
    /**
     * MFA type.
     * 
     */
    @Export(name="type", refs={String.class}, tree="[0]")
    private Output<String> type;

    /**
     * @return MFA type.
     * 
     */
    public Output<String> type() {
        return this.type;
    }
    /**
     * Use signature value, derived from &#34;settings_file_base64&#34;
     * 
     */
    @Export(name="useSignature", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> useSignature;

    /**
     * @return Use signature value, derived from &#34;settings_file_base64&#34;
     * 
     */
    public Output<Boolean> useSignature() {
        return this.useSignature;
    }
    /**
     * A template string for mapping Identity names to MFA methods.
     * 
     */
    @Export(name="usernameFormat", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> usernameFormat;

    /**
     * @return A template string for mapping Identity names to MFA methods.
     * 
     */
    public Output<Optional<String>> usernameFormat() {
        return Codegen.optional(this.usernameFormat);
    }
    /**
     * Resource UUID.
     * 
     */
    @Export(name="uuid", refs={String.class}, tree="[0]")
    private Output<String> uuid;

    /**
     * @return Resource UUID.
     * 
     */
    public Output<String> uuid() {
        return this.uuid;
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
        super("vault:identity/mfaPingid:MfaPingid", name, args == null ? MfaPingidArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private MfaPingid(String name, Output<String> id, @Nullable MfaPingidState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("vault:identity/mfaPingid:MfaPingid", name, state, makeResourceOptions(options, id));
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
