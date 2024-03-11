// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vault.pkiSecret;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.vault.Utilities;
import com.pulumi.vault.pkiSecret.SecretBackendConfigIssuersArgs;
import com.pulumi.vault.pkiSecret.inputs.SecretBackendConfigIssuersState;
import java.lang.Boolean;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * ## Example Usage
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.vault.Mount;
 * import com.pulumi.vault.MountArgs;
 * import com.pulumi.vault.pkiSecret.SecretBackendRootCert;
 * import com.pulumi.vault.pkiSecret.SecretBackendRootCertArgs;
 * import com.pulumi.vault.pkiSecret.SecretBackendIssuer;
 * import com.pulumi.vault.pkiSecret.SecretBackendIssuerArgs;
 * import com.pulumi.vault.pkiSecret.SecretBackendConfigIssuers;
 * import com.pulumi.vault.pkiSecret.SecretBackendConfigIssuersArgs;
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
 *         var pki = new Mount(&#34;pki&#34;, MountArgs.builder()        
 *             .path(&#34;pki&#34;)
 *             .type(&#34;pki&#34;)
 *             .defaultLeaseTtlSeconds(3600)
 *             .maxLeaseTtlSeconds(86400)
 *             .build());
 * 
 *         var root = new SecretBackendRootCert(&#34;root&#34;, SecretBackendRootCertArgs.builder()        
 *             .backend(pki.path())
 *             .type(&#34;internal&#34;)
 *             .commonName(&#34;test&#34;)
 *             .ttl(&#34;86400&#34;)
 *             .build());
 * 
 *         var example = new SecretBackendIssuer(&#34;example&#34;, SecretBackendIssuerArgs.builder()        
 *             .backend(root.backend())
 *             .issuerRef(root.issuerId())
 *             .issuerName(&#34;example-issuer&#34;)
 *             .build());
 * 
 *         var config = new SecretBackendConfigIssuers(&#34;config&#34;, SecretBackendConfigIssuersArgs.builder()        
 *             .backend(pki.path())
 *             .default_(example.issuerId())
 *             .defaultFollowsLatestIssuer(true)
 *             .build());
 * 
 *     }
 * }
 * ```
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Import
 * 
 * PKI secret backend config issuers can be imported using the path, e.g.
 * 
 * ```sh
 * $ pulumi import vault:pkiSecret/secretBackendConfigIssuers:SecretBackendConfigIssuers config pki/config/issuers
 * ```
 * 
 */
@ResourceType(type="vault:pkiSecret/secretBackendConfigIssuers:SecretBackendConfigIssuers")
public class SecretBackendConfigIssuers extends com.pulumi.resources.CustomResource {
    /**
     * The path the PKI secret backend is mounted at, with no
     * leading or trailing `/`s.
     * 
     */
    @Export(name="backend", refs={String.class}, tree="[0]")
    private Output<String> backend;

    /**
     * @return The path the PKI secret backend is mounted at, with no
     * leading or trailing `/`s.
     * 
     */
    public Output<String> backend() {
        return this.backend;
    }
    /**
     * Specifies the default issuer by ID.
     * 
     */
    @Export(name="default", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> default_;

    /**
     * @return Specifies the default issuer by ID.
     * 
     */
    public Output<Optional<String>> default_() {
        return Codegen.optional(this.default_);
    }
    /**
     * Specifies whether a root creation
     * or an issuer import operation updates the default issuer to the newly added issuer.
     * 
     */
    @Export(name="defaultFollowsLatestIssuer", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> defaultFollowsLatestIssuer;

    /**
     * @return Specifies whether a root creation
     * or an issuer import operation updates the default issuer to the newly added issuer.
     * 
     */
    public Output<Boolean> defaultFollowsLatestIssuer() {
        return this.defaultFollowsLatestIssuer;
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
    public SecretBackendConfigIssuers(String name) {
        this(name, SecretBackendConfigIssuersArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public SecretBackendConfigIssuers(String name, SecretBackendConfigIssuersArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public SecretBackendConfigIssuers(String name, SecretBackendConfigIssuersArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("vault:pkiSecret/secretBackendConfigIssuers:SecretBackendConfigIssuers", name, args == null ? SecretBackendConfigIssuersArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private SecretBackendConfigIssuers(String name, Output<String> id, @Nullable SecretBackendConfigIssuersState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("vault:pkiSecret/secretBackendConfigIssuers:SecretBackendConfigIssuers", name, state, makeResourceOptions(options, id));
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
    public static SecretBackendConfigIssuers get(String name, Output<String> id, @Nullable SecretBackendConfigIssuersState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new SecretBackendConfigIssuers(name, id, state, options);
    }
}
