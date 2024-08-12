// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vault;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.vault.AuditRequestHeaderArgs;
import com.pulumi.vault.Utilities;
import com.pulumi.vault.inputs.AuditRequestHeaderState;
import java.lang.Boolean;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Manages additional request headers that appear in audited requests.
 * 
 * &gt; **Note**
 * Because of the way the [sys/config/auditing/request-headers API](https://www.vaultproject.io/api-docs/system/config-auditing)
 * is implemented in Vault, this resource will manage existing audited headers with
 * matching names without requiring import.
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
 * import com.pulumi.vault.AuditRequestHeader;
 * import com.pulumi.vault.AuditRequestHeaderArgs;
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
 *         var xForwardedFor = new AuditRequestHeader("xForwardedFor", AuditRequestHeaderArgs.builder()
 *             .name("X-Forwarded-For")
 *             .hmac(false)
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 */
@ResourceType(type="vault:index/auditRequestHeader:AuditRequestHeader")
public class AuditRequestHeader extends com.pulumi.resources.CustomResource {
    /**
     * Whether this header&#39;s value should be HMAC&#39;d in the audit logs.
     * 
     */
    @Export(name="hmac", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> hmac;

    /**
     * @return Whether this header&#39;s value should be HMAC&#39;d in the audit logs.
     * 
     */
    public Output<Optional<Boolean>> hmac() {
        return Codegen.optional(this.hmac);
    }
    /**
     * The name of the request header to audit.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return The name of the request header to audit.
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
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public AuditRequestHeader(java.lang.String name) {
        this(name, AuditRequestHeaderArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public AuditRequestHeader(java.lang.String name, @Nullable AuditRequestHeaderArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public AuditRequestHeader(java.lang.String name, @Nullable AuditRequestHeaderArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("vault:index/auditRequestHeader:AuditRequestHeader", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private AuditRequestHeader(java.lang.String name, Output<java.lang.String> id, @Nullable AuditRequestHeaderState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("vault:index/auditRequestHeader:AuditRequestHeader", name, state, makeResourceOptions(options, id), false);
    }

    private static AuditRequestHeaderArgs makeArgs(@Nullable AuditRequestHeaderArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? AuditRequestHeaderArgs.Empty : args;
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
    public static AuditRequestHeader get(java.lang.String name, Output<java.lang.String> id, @Nullable AuditRequestHeaderState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new AuditRequestHeader(name, id, state, options);
    }
}
