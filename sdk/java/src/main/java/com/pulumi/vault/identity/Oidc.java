// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vault.identity;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.vault.Utilities;
import com.pulumi.vault.identity.OidcArgs;
import com.pulumi.vault.identity.inputs.OidcState;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Configure the [Identity Tokens Backend](https://www.vaultproject.io/docs/secrets/identity/index.html#identity-tokens).
 * 
 * The Identity secrets engine is the identity management solution for Vault. It internally maintains
 * the clients who are recognized by Vault.
 * 
 * &gt; **NOTE:** Each Vault server may only have one Identity Tokens Backend configuration. Multiple configurations of the resource against the same Vault server will cause a perpetual difference.
 * 
 * ## Example Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.vault.identity.Oidc;
 * import com.pulumi.vault.identity.OidcArgs;
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
 *         var server = new Oidc(&#34;server&#34;, OidcArgs.builder()        
 *             .issuer(&#34;https://www.acme.com&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 */
@ResourceType(type="vault:identity/oidc:Oidc")
public class Oidc extends com.pulumi.resources.CustomResource {
    /**
     * Issuer URL to be used in the iss claim of the token. If not set, Vault&#39;s
     * `api_addr` will be used. The issuer is a case sensitive URL using the https scheme that contains
     * scheme, host, and optionally, port number and path components, but no query or fragment
     * components.
     * 
     */
    @Export(name="issuer", refs={String.class}, tree="[0]")
    private Output<String> issuer;

    /**
     * @return Issuer URL to be used in the iss claim of the token. If not set, Vault&#39;s
     * `api_addr` will be used. The issuer is a case sensitive URL using the https scheme that contains
     * scheme, host, and optionally, port number and path components, but no query or fragment
     * components.
     * 
     */
    public Output<String> issuer() {
        return this.issuer;
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
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Oidc(String name) {
        this(name, OidcArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Oidc(String name, @Nullable OidcArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Oidc(String name, @Nullable OidcArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("vault:identity/oidc:Oidc", name, args == null ? OidcArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private Oidc(String name, Output<String> id, @Nullable OidcState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("vault:identity/oidc:Oidc", name, state, makeResourceOptions(options, id));
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
    public static Oidc get(String name, Output<String> id, @Nullable OidcState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Oidc(name, id, state, options);
    }
}
