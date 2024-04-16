// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vault.jwt;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.vault.Utilities;
import com.pulumi.vault.jwt.AuthBackendArgs;
import com.pulumi.vault.jwt.inputs.AuthBackendState;
import com.pulumi.vault.jwt.outputs.AuthBackendTune;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides a resource for managing an
 * [JWT auth backend within Vault](https://www.vaultproject.io/docs/auth/jwt.html).
 * 
 * ## Example Usage
 * 
 * Manage JWT auth backend:
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.vault.jwt.AuthBackend;
 * import com.pulumi.vault.jwt.AuthBackendArgs;
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
 *         var example = new AuthBackend(&#34;example&#34;, AuthBackendArgs.builder()        
 *             .description(&#34;Demonstration of the Terraform JWT auth backend&#34;)
 *             .path(&#34;jwt&#34;)
 *             .oidcDiscoveryUrl(&#34;https://myco.auth0.com/&#34;)
 *             .boundIssuer(&#34;https://myco.auth0.com/&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * Manage OIDC auth backend:
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.vault.jwt.AuthBackend;
 * import com.pulumi.vault.jwt.AuthBackendArgs;
 * import com.pulumi.vault.jwt.inputs.AuthBackendTuneArgs;
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
 *         var example = new AuthBackend(&#34;example&#34;, AuthBackendArgs.builder()        
 *             .description(&#34;Demonstration of the Terraform JWT auth backend&#34;)
 *             .path(&#34;oidc&#34;)
 *             .type(&#34;oidc&#34;)
 *             .oidcDiscoveryUrl(&#34;https://myco.auth0.com/&#34;)
 *             .oidcClientId(&#34;1234567890&#34;)
 *             .oidcClientSecret(&#34;secret123456&#34;)
 *             .boundIssuer(&#34;https://myco.auth0.com/&#34;)
 *             .tune(AuthBackendTuneArgs.builder()
 *                 .listingVisibility(&#34;unauth&#34;)
 *                 .build())
 *             .build());
 * 
 *     }
 * }
 * ```
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * Configuring the auth backend with a `provider_config:
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.vault.jwt.AuthBackend;
 * import com.pulumi.vault.jwt.AuthBackendArgs;
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
 *         var gsuite = new AuthBackend(&#34;gsuite&#34;, AuthBackendArgs.builder()        
 *             .description(&#34;OIDC backend&#34;)
 *             .oidcDiscoveryUrl(&#34;https://accounts.google.com&#34;)
 *             .path(&#34;oidc&#34;)
 *             .type(&#34;oidc&#34;)
 *             .providerConfig(Map.ofEntries(
 *                 Map.entry(&#34;provider&#34;, &#34;gsuite&#34;),
 *                 Map.entry(&#34;fetch_groups&#34;, true),
 *                 Map.entry(&#34;fetch_user_info&#34;, true),
 *                 Map.entry(&#34;groups_recurse_max_depth&#34;, 1)
 *             ))
 *             .build());
 * 
 *     }
 * }
 * ```
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Import
 * 
 * JWT auth backend can be imported using the `path`, e.g.
 * 
 * ```sh
 * $ pulumi import vault:jwt/authBackend:AuthBackend oidc oidc
 * ```
 * or
 * 
 * ```sh
 * $ pulumi import vault:jwt/authBackend:AuthBackend jwt jwt
 * ```
 * 
 */
@ResourceType(type="vault:jwt/authBackend:AuthBackend")
public class AuthBackend extends com.pulumi.resources.CustomResource {
    /**
     * The accessor for this auth method
     * 
     */
    @Export(name="accessor", refs={String.class}, tree="[0]")
    private Output<String> accessor;

    /**
     * @return The accessor for this auth method
     * 
     */
    public Output<String> accessor() {
        return this.accessor;
    }
    /**
     * The value against which to match the iss claim in a JWT
     * 
     */
    @Export(name="boundIssuer", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> boundIssuer;

    /**
     * @return The value against which to match the iss claim in a JWT
     * 
     */
    public Output<Optional<String>> boundIssuer() {
        return Codegen.optional(this.boundIssuer);
    }
    /**
     * The default role to use if none is provided during login
     * 
     */
    @Export(name="defaultRole", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> defaultRole;

    /**
     * @return The default role to use if none is provided during login
     * 
     */
    public Output<Optional<String>> defaultRole() {
        return Codegen.optional(this.defaultRole);
    }
    /**
     * The description of the auth backend
     * 
     */
    @Export(name="description", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> description;

    /**
     * @return The description of the auth backend
     * 
     */
    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    /**
     * If set, opts out of mount migration on path updates.
     * See here for more info on [Mount Migration](https://www.vaultproject.io/docs/concepts/mount-migration)
     * 
     */
    @Export(name="disableRemount", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> disableRemount;

    /**
     * @return If set, opts out of mount migration on path updates.
     * See here for more info on [Mount Migration](https://www.vaultproject.io/docs/concepts/mount-migration)
     * 
     */
    public Output<Optional<Boolean>> disableRemount() {
        return Codegen.optional(this.disableRemount);
    }
    /**
     * The CA certificate or chain of certificates, in PEM format, to use to validate connections to the JWKS URL. If not set, system certificates are used.
     * 
     */
    @Export(name="jwksCaPem", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> jwksCaPem;

    /**
     * @return The CA certificate or chain of certificates, in PEM format, to use to validate connections to the JWKS URL. If not set, system certificates are used.
     * 
     */
    public Output<Optional<String>> jwksCaPem() {
        return Codegen.optional(this.jwksCaPem);
    }
    /**
     * JWKS URL to use to authenticate signatures. Cannot be used with &#34;oidc_discovery_url&#34; or &#34;jwt_validation_pubkeys&#34;.
     * 
     */
    @Export(name="jwksUrl", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> jwksUrl;

    /**
     * @return JWKS URL to use to authenticate signatures. Cannot be used with &#34;oidc_discovery_url&#34; or &#34;jwt_validation_pubkeys&#34;.
     * 
     */
    public Output<Optional<String>> jwksUrl() {
        return Codegen.optional(this.jwksUrl);
    }
    /**
     * A list of supported signing algorithms. Vault 1.1.0 defaults to [RS256] but future or past versions of Vault may differ
     * 
     */
    @Export(name="jwtSupportedAlgs", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> jwtSupportedAlgs;

    /**
     * @return A list of supported signing algorithms. Vault 1.1.0 defaults to [RS256] but future or past versions of Vault may differ
     * 
     */
    public Output<Optional<List<String>>> jwtSupportedAlgs() {
        return Codegen.optional(this.jwtSupportedAlgs);
    }
    /**
     * A list of PEM-encoded public keys to use to authenticate signatures locally. Cannot be used in combination with `oidc_discovery_url`
     * 
     */
    @Export(name="jwtValidationPubkeys", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> jwtValidationPubkeys;

    /**
     * @return A list of PEM-encoded public keys to use to authenticate signatures locally. Cannot be used in combination with `oidc_discovery_url`
     * 
     */
    public Output<Optional<List<String>>> jwtValidationPubkeys() {
        return Codegen.optional(this.jwtValidationPubkeys);
    }
    /**
     * Specifies if the auth method is local only.
     * 
     */
    @Export(name="local", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> local;

    /**
     * @return Specifies if the auth method is local only.
     * 
     */
    public Output<Optional<Boolean>> local() {
        return Codegen.optional(this.local);
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
     * Pass namespace in the OIDC state parameter instead of as a separate query parameter. With this setting, the allowed redirect URL(s) in Vault and on the provider side should not contain a namespace query parameter. This means only one redirect URL entry needs to be maintained on the OIDC provider side for all vault namespaces that will be authenticating against it. Defaults to true for new configs
     * 
     * * tune - (Optional) Extra configuration block. Structure is documented below.
     * 
     * The `tune` block is used to tune the auth backend:
     * 
     */
    @Export(name="namespaceInState", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> namespaceInState;

    /**
     * @return Pass namespace in the OIDC state parameter instead of as a separate query parameter. With this setting, the allowed redirect URL(s) in Vault and on the provider side should not contain a namespace query parameter. This means only one redirect URL entry needs to be maintained on the OIDC provider side for all vault namespaces that will be authenticating against it. Defaults to true for new configs
     * 
     * * tune - (Optional) Extra configuration block. Structure is documented below.
     * 
     * The `tune` block is used to tune the auth backend:
     * 
     */
    public Output<Optional<Boolean>> namespaceInState() {
        return Codegen.optional(this.namespaceInState);
    }
    /**
     * Client ID used for OIDC backends
     * 
     */
    @Export(name="oidcClientId", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> oidcClientId;

    /**
     * @return Client ID used for OIDC backends
     * 
     */
    public Output<Optional<String>> oidcClientId() {
        return Codegen.optional(this.oidcClientId);
    }
    /**
     * Client Secret used for OIDC backends
     * 
     */
    @Export(name="oidcClientSecret", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> oidcClientSecret;

    /**
     * @return Client Secret used for OIDC backends
     * 
     */
    public Output<Optional<String>> oidcClientSecret() {
        return Codegen.optional(this.oidcClientSecret);
    }
    /**
     * The CA certificate or chain of certificates, in PEM format, to use to validate connections to the OIDC Discovery URL. If not set, system certificates are used
     * 
     */
    @Export(name="oidcDiscoveryCaPem", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> oidcDiscoveryCaPem;

    /**
     * @return The CA certificate or chain of certificates, in PEM format, to use to validate connections to the OIDC Discovery URL. If not set, system certificates are used
     * 
     */
    public Output<Optional<String>> oidcDiscoveryCaPem() {
        return Codegen.optional(this.oidcDiscoveryCaPem);
    }
    /**
     * The OIDC Discovery URL, without any .well-known component (base path). Cannot be used in combination with `jwt_validation_pubkeys`
     * 
     */
    @Export(name="oidcDiscoveryUrl", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> oidcDiscoveryUrl;

    /**
     * @return The OIDC Discovery URL, without any .well-known component (base path). Cannot be used in combination with `jwt_validation_pubkeys`
     * 
     */
    public Output<Optional<String>> oidcDiscoveryUrl() {
        return Codegen.optional(this.oidcDiscoveryUrl);
    }
    /**
     * The response mode to be used in the OAuth2 request. Allowed values are `query` and `form_post`. Defaults to `query`. If using Vault namespaces, and `oidc_response_mode` is `form_post`, then `namespace_in_state` should be set to `false`.
     * 
     */
    @Export(name="oidcResponseMode", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> oidcResponseMode;

    /**
     * @return The response mode to be used in the OAuth2 request. Allowed values are `query` and `form_post`. Defaults to `query`. If using Vault namespaces, and `oidc_response_mode` is `form_post`, then `namespace_in_state` should be set to `false`.
     * 
     */
    public Output<Optional<String>> oidcResponseMode() {
        return Codegen.optional(this.oidcResponseMode);
    }
    /**
     * List of response types to request. Allowed values are &#39;code&#39; and &#39;id_token&#39;. Defaults to `[&#34;code&#34;]`. Note: `id_token` may only be used if `oidc_response_mode` is set to `form_post`.
     * 
     */
    @Export(name="oidcResponseTypes", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> oidcResponseTypes;

    /**
     * @return List of response types to request. Allowed values are &#39;code&#39; and &#39;id_token&#39;. Defaults to `[&#34;code&#34;]`. Note: `id_token` may only be used if `oidc_response_mode` is set to `form_post`.
     * 
     */
    public Output<Optional<List<String>>> oidcResponseTypes() {
        return Codegen.optional(this.oidcResponseTypes);
    }
    /**
     * Path to mount the JWT/OIDC auth backend
     * 
     */
    @Export(name="path", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> path;

    /**
     * @return Path to mount the JWT/OIDC auth backend
     * 
     */
    public Output<Optional<String>> path() {
        return Codegen.optional(this.path);
    }
    /**
     * Provider specific handling configuration. All values may be strings, and the provider will convert to the appropriate type when configuring Vault.
     * 
     */
    @Export(name="providerConfig", refs={Map.class,String.class}, tree="[0,1,1]")
    private Output</* @Nullable */ Map<String,String>> providerConfig;

    /**
     * @return Provider specific handling configuration. All values may be strings, and the provider will convert to the appropriate type when configuring Vault.
     * 
     */
    public Output<Optional<Map<String,String>>> providerConfig() {
        return Codegen.optional(this.providerConfig);
    }
    @Export(name="tune", refs={AuthBackendTune.class}, tree="[0]")
    private Output<AuthBackendTune> tune;

    public Output<AuthBackendTune> tune() {
        return this.tune;
    }
    /**
     * Type of auth backend. Should be one of `jwt` or `oidc`. Default - `jwt`
     * 
     */
    @Export(name="type", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> type;

    /**
     * @return Type of auth backend. Should be one of `jwt` or `oidc`. Default - `jwt`
     * 
     */
    public Output<Optional<String>> type() {
        return Codegen.optional(this.type);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public AuthBackend(String name) {
        this(name, AuthBackendArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public AuthBackend(String name, @Nullable AuthBackendArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public AuthBackend(String name, @Nullable AuthBackendArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("vault:jwt/authBackend:AuthBackend", name, args == null ? AuthBackendArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private AuthBackend(String name, Output<String> id, @Nullable AuthBackendState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("vault:jwt/authBackend:AuthBackend", name, state, makeResourceOptions(options, id));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .additionalSecretOutputs(List.of(
                "oidcClientSecret"
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
    public static AuthBackend get(String name, Output<String> id, @Nullable AuthBackendState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new AuthBackend(name, id, state, options);
    }
}
