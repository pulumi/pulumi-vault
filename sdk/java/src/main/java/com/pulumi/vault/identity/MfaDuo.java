// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vault.identity;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.vault.Utilities;
import com.pulumi.vault.identity.MfaDuoArgs;
import com.pulumi.vault.identity.inputs.MfaDuoState;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Resource for configuring the duo MFA method.
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
 * import com.pulumi.vault.identity.MfaDuo;
 * import com.pulumi.vault.identity.MfaDuoArgs;
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
 *         var example = new MfaDuo("example", MfaDuoArgs.builder()
 *             .apiHostname("api-xxxxxxxx.duosecurity.com")
 *             .secretKey("secret-key")
 *             .integrationKey("secret-int-key")
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
 * $ pulumi import vault:identity/mfaDuo:MfaDuo example 0d89c36a-4ff5-4d70-8749-bb6a5598aeec
 * ```
 * 
 */
@ResourceType(type="vault:identity/mfaDuo:MfaDuo")
public class MfaDuo extends com.pulumi.resources.CustomResource {
    /**
     * API hostname for Duo
     * 
     */
    @Export(name="apiHostname", refs={String.class}, tree="[0]")
    private Output<String> apiHostname;

    /**
     * @return API hostname for Duo
     * 
     */
    public Output<String> apiHostname() {
        return this.apiHostname;
    }
    /**
     * Integration key for Duo
     * 
     */
    @Export(name="integrationKey", refs={String.class}, tree="[0]")
    private Output<String> integrationKey;

    /**
     * @return Integration key for Duo
     * 
     */
    public Output<String> integrationKey() {
        return this.integrationKey;
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
     * Push information for Duo.
     * 
     */
    @Export(name="pushInfo", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> pushInfo;

    /**
     * @return Push information for Duo.
     * 
     */
    public Output<Optional<String>> pushInfo() {
        return Codegen.optional(this.pushInfo);
    }
    /**
     * Secret key for Duo
     * 
     */
    @Export(name="secretKey", refs={String.class}, tree="[0]")
    private Output<String> secretKey;

    /**
     * @return Secret key for Duo
     * 
     */
    public Output<String> secretKey() {
        return this.secretKey;
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
     * Require passcode upon MFA validation.
     * 
     */
    @Export(name="usePasscode", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> usePasscode;

    /**
     * @return Require passcode upon MFA validation.
     * 
     */
    public Output<Optional<Boolean>> usePasscode() {
        return Codegen.optional(this.usePasscode);
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
        super("vault:identity/mfaDuo:MfaDuo", name, args == null ? MfaDuoArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private MfaDuo(String name, Output<String> id, @Nullable MfaDuoState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("vault:identity/mfaDuo:MfaDuo", name, state, makeResourceOptions(options, id));
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
