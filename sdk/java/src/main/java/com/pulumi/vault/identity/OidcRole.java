// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vault.identity;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.vault.Utilities;
import com.pulumi.vault.identity.OidcRoleArgs;
import com.pulumi.vault.identity.inputs.OidcRoleState;
import java.lang.Integer;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * ## Example Usage
 * 
 * You need to create a role with a named key.
 * At creation time, the key can be created independently of the role. However, the key must
 * exist before the role can be used to issue tokens. You must also configure the key with the
 * role&#39;s Client ID to allow the role to use the key.
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.vault.identity.OidcRole;
 * import com.pulumi.vault.identity.OidcRoleArgs;
 * import com.pulumi.vault.identity.OidcKey;
 * import com.pulumi.vault.identity.OidcKeyArgs;
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
 *         final var key = config.get("key").orElse("key");
 *         var role = new OidcRole("role", OidcRoleArgs.builder()
 *             .name("role")
 *             .key(key)
 *             .build());
 * 
 *         var keyOidcKey = new OidcKey("keyOidcKey", OidcKeyArgs.builder()
 *             .name(key)
 *             .algorithm("RS256")
 *             .allowedClientIds(role.clientId())
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * If you want to create the key first before creating the role, you can use a separate
 * resource to configure the allowed Client ID on
 * the key.
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.vault.identity.OidcKey;
 * import com.pulumi.vault.identity.OidcKeyArgs;
 * import com.pulumi.vault.identity.OidcRole;
 * import com.pulumi.vault.identity.OidcRoleArgs;
 * import com.pulumi.vault.identity.OidcKeyAllowedClientID;
 * import com.pulumi.vault.identity.OidcKeyAllowedClientIDArgs;
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
 *         var key = new OidcKey("key", OidcKeyArgs.builder()
 *             .name("key")
 *             .algorithm("RS256")
 *             .build());
 * 
 *         var role = new OidcRole("role", OidcRoleArgs.builder()
 *             .name("role")
 *             .key(key.name())
 *             .build());
 * 
 *         var roleOidcKeyAllowedClientID = new OidcKeyAllowedClientID("roleOidcKeyAllowedClientID", OidcKeyAllowedClientIDArgs.builder()
 *             .keyName(key.name())
 *             .allowedClientId(role.clientId())
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
 * The key can be imported with the role name, for example:
 * 
 * ```sh
 * $ pulumi import vault:identity/oidcRole:OidcRole role role
 * ```
 * 
 */
@ResourceType(type="vault:identity/oidcRole:OidcRole")
public class OidcRole extends com.pulumi.resources.CustomResource {
    /**
     * The value that will be included in the `aud` field of all the OIDC identity
     * tokens issued by this role
     * 
     */
    @Export(name="clientId", refs={String.class}, tree="[0]")
    private Output<String> clientId;

    /**
     * @return The value that will be included in the `aud` field of all the OIDC identity
     * tokens issued by this role
     * 
     */
    public Output<String> clientId() {
        return this.clientId;
    }
    /**
     * A configured named key, the key must already exist
     * before tokens can be issued.
     * 
     */
    @Export(name="key", refs={String.class}, tree="[0]")
    private Output<String> key;

    /**
     * @return A configured named key, the key must already exist
     * before tokens can be issued.
     * 
     */
    public Output<String> key() {
        return this.key;
    }
    /**
     * Name of the OIDC Role to create.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return Name of the OIDC Role to create.
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
     * The template string to use for generating tokens. This may be in
     * string-ified JSON or base64 format. See the
     * [documentation](https://www.vaultproject.io/docs/secrets/identity/index.html#token-contents-and-templates)
     * for the template format.
     * 
     */
    @Export(name="template", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> template;

    /**
     * @return The template string to use for generating tokens. This may be in
     * string-ified JSON or base64 format. See the
     * [documentation](https://www.vaultproject.io/docs/secrets/identity/index.html#token-contents-and-templates)
     * for the template format.
     * 
     */
    public Output<Optional<String>> template() {
        return Codegen.optional(this.template);
    }
    /**
     * TTL of the tokens generated against the role in number of seconds.
     * 
     */
    @Export(name="ttl", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> ttl;

    /**
     * @return TTL of the tokens generated against the role in number of seconds.
     * 
     */
    public Output<Optional<Integer>> ttl() {
        return Codegen.optional(this.ttl);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public OidcRole(java.lang.String name) {
        this(name, OidcRoleArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public OidcRole(java.lang.String name, OidcRoleArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public OidcRole(java.lang.String name, OidcRoleArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("vault:identity/oidcRole:OidcRole", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private OidcRole(java.lang.String name, Output<java.lang.String> id, @Nullable OidcRoleState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("vault:identity/oidcRole:OidcRole", name, state, makeResourceOptions(options, id), false);
    }

    private static OidcRoleArgs makeArgs(OidcRoleArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? OidcRoleArgs.Empty : args;
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
    public static OidcRole get(java.lang.String name, Output<java.lang.String> id, @Nullable OidcRoleState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new OidcRole(name, id, state, options);
    }
}
