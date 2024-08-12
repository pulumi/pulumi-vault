// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vault.ldap;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.vault.Utilities;
import com.pulumi.vault.ldap.SecretBackendLibrarySetArgs;
import com.pulumi.vault.ldap.inputs.SecretBackendLibrarySetState;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
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
 * import com.pulumi.vault.ldap.SecretBackend;
 * import com.pulumi.vault.ldap.SecretBackendArgs;
 * import com.pulumi.vault.ldap.SecretBackendLibrarySet;
 * import com.pulumi.vault.ldap.SecretBackendLibrarySetArgs;
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
 *         var config = new SecretBackend("config", SecretBackendArgs.builder()
 *             .path("ldap")
 *             .binddn("CN=Administrator,CN=Users,DC=corp,DC=example,DC=net")
 *             .bindpass("SuperSecretPassw0rd")
 *             .url("ldaps://localhost")
 *             .insecureTls("true")
 *             .userdn("CN=Users,DC=corp,DC=example,DC=net")
 *             .build());
 * 
 *         var qa = new SecretBackendLibrarySet("qa", SecretBackendLibrarySetArgs.builder()
 *             .mount(config.path())
 *             .name("qa")
 *             .serviceAccountNames(            
 *                 "Bob",
 *                 "Mary")
 *             .ttl(60)
 *             .disableCheckInEnforcement(true)
 *             .maxTtl(120)
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
 * LDAP secret backend libraries can be imported using the `path`, e.g.
 * 
 * ```sh
 * $ pulumi import vault:ldap/secretBackendLibrarySet:SecretBackendLibrarySet qa ldap/library/bob
 * ```
 * 
 */
@ResourceType(type="vault:ldap/secretBackendLibrarySet:SecretBackendLibrarySet")
public class SecretBackendLibrarySet extends com.pulumi.resources.CustomResource {
    /**
     * Disable enforcing that service
     * accounts must be checked in by the entity or client token that checked them
     * out. Defaults to false.
     * 
     */
    @Export(name="disableCheckInEnforcement", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> disableCheckInEnforcement;

    /**
     * @return Disable enforcing that service
     * accounts must be checked in by the entity or client token that checked them
     * out. Defaults to false.
     * 
     */
    public Output<Optional<Boolean>> disableCheckInEnforcement() {
        return Codegen.optional(this.disableCheckInEnforcement);
    }
    /**
     * The maximum password time-to-live in seconds. Defaults
     * to the configuration max_ttl if not provided.
     * 
     */
    @Export(name="maxTtl", refs={Integer.class}, tree="[0]")
    private Output<Integer> maxTtl;

    /**
     * @return The maximum password time-to-live in seconds. Defaults
     * to the configuration max_ttl if not provided.
     * 
     */
    public Output<Integer> maxTtl() {
        return this.maxTtl;
    }
    /**
     * The path where the LDAP secrets backend is mounted.
     * 
     */
    @Export(name="mount", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> mount;

    /**
     * @return The path where the LDAP secrets backend is mounted.
     * 
     */
    public Output<Optional<String>> mount() {
        return Codegen.optional(this.mount);
    }
    /**
     * The name to identify this set of service accounts.
     * Must be unique within the backend.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return The name to identify this set of service accounts.
     * Must be unique within the backend.
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
     * Specifies the slice of service accounts mapped to this set.
     * 
     */
    @Export(name="serviceAccountNames", refs={List.class,String.class}, tree="[0,1]")
    private Output<List<String>> serviceAccountNames;

    /**
     * @return Specifies the slice of service accounts mapped to this set.
     * 
     */
    public Output<List<String>> serviceAccountNames() {
        return this.serviceAccountNames;
    }
    /**
     * The password time-to-live in seconds. Defaults to the configuration
     * ttl if not provided.
     * 
     */
    @Export(name="ttl", refs={Integer.class}, tree="[0]")
    private Output<Integer> ttl;

    /**
     * @return The password time-to-live in seconds. Defaults to the configuration
     * ttl if not provided.
     * 
     */
    public Output<Integer> ttl() {
        return this.ttl;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public SecretBackendLibrarySet(java.lang.String name) {
        this(name, SecretBackendLibrarySetArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public SecretBackendLibrarySet(java.lang.String name, SecretBackendLibrarySetArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public SecretBackendLibrarySet(java.lang.String name, SecretBackendLibrarySetArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("vault:ldap/secretBackendLibrarySet:SecretBackendLibrarySet", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private SecretBackendLibrarySet(java.lang.String name, Output<java.lang.String> id, @Nullable SecretBackendLibrarySetState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("vault:ldap/secretBackendLibrarySet:SecretBackendLibrarySet", name, state, makeResourceOptions(options, id), false);
    }

    private static SecretBackendLibrarySetArgs makeArgs(SecretBackendLibrarySetArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? SecretBackendLibrarySetArgs.Empty : args;
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
    public static SecretBackendLibrarySet get(java.lang.String name, Output<java.lang.String> id, @Nullable SecretBackendLibrarySetState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new SecretBackendLibrarySet(name, id, state, options);
    }
}
