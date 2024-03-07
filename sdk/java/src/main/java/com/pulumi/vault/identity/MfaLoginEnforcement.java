// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vault.identity;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.vault.Utilities;
import com.pulumi.vault.identity.MfaLoginEnforcementArgs;
import com.pulumi.vault.identity.inputs.MfaLoginEnforcementState;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Resource for configuring MFA login-enforcement
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
 * import com.pulumi.vault.identity.MfaDuo;
 * import com.pulumi.vault.identity.MfaDuoArgs;
 * import com.pulumi.vault.identity.MfaLoginEnforcement;
 * import com.pulumi.vault.identity.MfaLoginEnforcementArgs;
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
 *         var exampleMfaDuo = new MfaDuo(&#34;exampleMfaDuo&#34;, MfaDuoArgs.builder()        
 *             .secretKey(&#34;secret-key&#34;)
 *             .integrationKey(&#34;int-key&#34;)
 *             .apiHostname(&#34;foo.baz&#34;)
 *             .pushInfo(&#34;push-info&#34;)
 *             .build());
 * 
 *         var exampleMfaLoginEnforcement = new MfaLoginEnforcement(&#34;exampleMfaLoginEnforcement&#34;, MfaLoginEnforcementArgs.builder()        
 *             .mfaMethodIds(exampleMfaDuo.methodId())
 *             .build());
 * 
 *     }
 * }
 * ```
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Import
 * 
 * Resource can be imported using its `name` field, e.g.
 * 
 * ```sh
 * $ pulumi import vault:identity/mfaLoginEnforcement:MfaLoginEnforcement example default
 * ```
 * 
 */
@ResourceType(type="vault:identity/mfaLoginEnforcement:MfaLoginEnforcement")
public class MfaLoginEnforcement extends com.pulumi.resources.CustomResource {
    /**
     * Set of auth method accessor IDs.
     * 
     */
    @Export(name="authMethodAccessors", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> authMethodAccessors;

    /**
     * @return Set of auth method accessor IDs.
     * 
     */
    public Output<Optional<List<String>>> authMethodAccessors() {
        return Codegen.optional(this.authMethodAccessors);
    }
    /**
     * Set of auth method types.
     * 
     */
    @Export(name="authMethodTypes", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> authMethodTypes;

    /**
     * @return Set of auth method types.
     * 
     */
    public Output<Optional<List<String>>> authMethodTypes() {
        return Codegen.optional(this.authMethodTypes);
    }
    /**
     * Set of identity entity IDs.
     * 
     */
    @Export(name="identityEntityIds", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> identityEntityIds;

    /**
     * @return Set of identity entity IDs.
     * 
     */
    public Output<Optional<List<String>>> identityEntityIds() {
        return Codegen.optional(this.identityEntityIds);
    }
    /**
     * Set of identity group IDs.
     * 
     */
    @Export(name="identityGroupIds", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> identityGroupIds;

    /**
     * @return Set of identity group IDs.
     * 
     */
    public Output<Optional<List<String>>> identityGroupIds() {
        return Codegen.optional(this.identityGroupIds);
    }
    /**
     * Set of MFA method UUIDs.
     * 
     */
    @Export(name="mfaMethodIds", refs={List.class,String.class}, tree="[0,1]")
    private Output<List<String>> mfaMethodIds;

    /**
     * @return Set of MFA method UUIDs.
     * 
     */
    public Output<List<String>> mfaMethodIds() {
        return this.mfaMethodIds;
    }
    /**
     * Login enforcement name.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return Login enforcement name.
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
    public MfaLoginEnforcement(String name) {
        this(name, MfaLoginEnforcementArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public MfaLoginEnforcement(String name, MfaLoginEnforcementArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public MfaLoginEnforcement(String name, MfaLoginEnforcementArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("vault:identity/mfaLoginEnforcement:MfaLoginEnforcement", name, args == null ? MfaLoginEnforcementArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private MfaLoginEnforcement(String name, Output<String> id, @Nullable MfaLoginEnforcementState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("vault:identity/mfaLoginEnforcement:MfaLoginEnforcement", name, state, makeResourceOptions(options, id));
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
    public static MfaLoginEnforcement get(String name, Output<String> id, @Nullable MfaLoginEnforcementState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new MfaLoginEnforcement(name, id, state, options);
    }
}
