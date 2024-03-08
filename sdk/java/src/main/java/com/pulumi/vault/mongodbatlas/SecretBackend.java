// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vault.mongodbatlas;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.vault.Utilities;
import com.pulumi.vault.mongodbatlas.SecretBackendArgs;
import com.pulumi.vault.mongodbatlas.inputs.SecretBackendState;
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
 * import com.pulumi.vault.mongodbatlas.SecretBackend;
 * import com.pulumi.vault.mongodbatlas.SecretBackendArgs;
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
 *         var mongo = new Mount(&#34;mongo&#34;, MountArgs.builder()        
 *             .description(&#34;MongoDB Atlas secret engine mount&#34;)
 *             .path(&#34;mongodbatlas&#34;)
 *             .type(&#34;mongodbatlas&#34;)
 *             .build());
 * 
 *         var config = new SecretBackend(&#34;config&#34;, SecretBackendArgs.builder()        
 *             .mount(&#34;vault_mount.mongo.path&#34;)
 *             .privateKey(&#34;privateKey&#34;)
 *             .publicKey(&#34;publicKey&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Import
 * 
 * MongoDB Atlas secret backends can be imported using the `${mount}/config`, e.g.
 * 
 * ```sh
 * $ pulumi import vault:mongodbatlas/secretBackend:SecretBackend config mongodbatlas/config
 * ```
 * 
 */
@ResourceType(type="vault:mongodbatlas/secretBackend:SecretBackend")
public class SecretBackend extends com.pulumi.resources.CustomResource {
    /**
     * Path where the MongoDB Atlas Secrets Engine is mounted.
     * 
     */
    @Export(name="mount", refs={String.class}, tree="[0]")
    private Output<String> mount;

    /**
     * @return Path where the MongoDB Atlas Secrets Engine is mounted.
     * 
     */
    public Output<String> mount() {
        return this.mount;
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
     * Path where MongoDB Atlas configuration is located
     * 
     */
    @Export(name="path", refs={String.class}, tree="[0]")
    private Output<String> path;

    /**
     * @return Path where MongoDB Atlas configuration is located
     * 
     */
    public Output<String> path() {
        return this.path;
    }
    /**
     * Specifies the Private API Key used to authenticate with the MongoDB Atlas API.
     * 
     */
    @Export(name="privateKey", refs={String.class}, tree="[0]")
    private Output<String> privateKey;

    /**
     * @return Specifies the Private API Key used to authenticate with the MongoDB Atlas API.
     * 
     */
    public Output<String> privateKey() {
        return this.privateKey;
    }
    /**
     * Specifies the Public API Key used to authenticate with the MongoDB Atlas API.
     * 
     */
    @Export(name="publicKey", refs={String.class}, tree="[0]")
    private Output<String> publicKey;

    /**
     * @return Specifies the Public API Key used to authenticate with the MongoDB Atlas API.
     * 
     */
    public Output<String> publicKey() {
        return this.publicKey;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public SecretBackend(String name) {
        this(name, SecretBackendArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public SecretBackend(String name, SecretBackendArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public SecretBackend(String name, SecretBackendArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("vault:mongodbatlas/secretBackend:SecretBackend", name, args == null ? SecretBackendArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private SecretBackend(String name, Output<String> id, @Nullable SecretBackendState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("vault:mongodbatlas/secretBackend:SecretBackend", name, state, makeResourceOptions(options, id));
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
    public static SecretBackend get(String name, Output<String> id, @Nullable SecretBackendState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new SecretBackend(name, id, state, options);
    }
}
