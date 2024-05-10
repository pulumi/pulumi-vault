// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vault;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.vault.MfaTotpArgs;
import com.pulumi.vault.Utilities;
import com.pulumi.vault.inputs.MfaTotpState;
import java.lang.Integer;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides a resource to manage [TOTP MFA](https://www.vaultproject.io/docs/enterprise/mfa/mfa-totp).
 * 
 * **Note** this feature is available only with Vault Enterprise.
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
 * import com.pulumi.vault.MfaTotp;
 * import com.pulumi.vault.MfaTotpArgs;
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
 *         var myTotp = new MfaTotp("myTotp", MfaTotpArgs.builder()        
 *             .name("my_totp")
 *             .issuer("hashicorp")
 *             .period(60)
 *             .algorithm("SHA256")
 *             .digits(8)
 *             .keySize(20)
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
 * Mounts can be imported using the `path`, e.g.
 * 
 * ```sh
 * $ pulumi import vault:index/mfaTotp:MfaTotp my_totp my_totp
 * ```
 * 
 */
@ResourceType(type="vault:index/mfaTotp:MfaTotp")
public class MfaTotp extends com.pulumi.resources.CustomResource {
    /**
     * `(string)` - Specifies the hashing algorithm used to generate the TOTP code.
     * Options include `SHA1`, `SHA256` and `SHA512`
     * 
     */
    @Export(name="algorithm", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> algorithm;

    /**
     * @return `(string)` - Specifies the hashing algorithm used to generate the TOTP code.
     * Options include `SHA1`, `SHA256` and `SHA512`
     * 
     */
    public Output<Optional<String>> algorithm() {
        return Codegen.optional(this.algorithm);
    }
    /**
     * `(int)` - The number of digits in the generated TOTP token.
     * This value can either be 6 or 8.
     * 
     */
    @Export(name="digits", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> digits;

    /**
     * @return `(int)` - The number of digits in the generated TOTP token.
     * This value can either be 6 or 8.
     * 
     */
    public Output<Optional<Integer>> digits() {
        return Codegen.optional(this.digits);
    }
    /**
     * `(string: &lt;required&gt;)` - The name of the key&#39;s issuing organization.
     * 
     */
    @Export(name="issuer", refs={String.class}, tree="[0]")
    private Output<String> issuer;

    /**
     * @return `(string: &lt;required&gt;)` - The name of the key&#39;s issuing organization.
     * 
     */
    public Output<String> issuer() {
        return this.issuer;
    }
    /**
     * `(int)` - Specifies the size in bytes of the generated key.
     * 
     */
    @Export(name="keySize", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> keySize;

    /**
     * @return `(int)` - Specifies the size in bytes of the generated key.
     * 
     */
    public Output<Optional<Integer>> keySize() {
        return Codegen.optional(this.keySize);
    }
    /**
     * `(string: &lt;required&gt;)` – Name of the MFA method.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return `(string: &lt;required&gt;)` – Name of the MFA method.
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
     * `(int)` - The length of time used to generate a counter for the TOTP token calculation.
     * 
     */
    @Export(name="period", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> period;

    /**
     * @return `(int)` - The length of time used to generate a counter for the TOTP token calculation.
     * 
     */
    public Output<Optional<Integer>> period() {
        return Codegen.optional(this.period);
    }
    /**
     * `(int)` - The pixel size of the generated square QR code.
     * 
     */
    @Export(name="qrSize", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> qrSize;

    /**
     * @return `(int)` - The pixel size of the generated square QR code.
     * 
     */
    public Output<Optional<Integer>> qrSize() {
        return Codegen.optional(this.qrSize);
    }
    /**
     * `(int)` - The number of delay periods that are allowed when validating a TOTP token.
     * This value can either be 0 or 1.
     * 
     */
    @Export(name="skew", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> skew;

    /**
     * @return `(int)` - The number of delay periods that are allowed when validating a TOTP token.
     * This value can either be 0 or 1.
     * 
     */
    public Output<Optional<Integer>> skew() {
        return Codegen.optional(this.skew);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public MfaTotp(String name) {
        this(name, MfaTotpArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public MfaTotp(String name, MfaTotpArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public MfaTotp(String name, MfaTotpArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("vault:index/mfaTotp:MfaTotp", name, args == null ? MfaTotpArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private MfaTotp(String name, Output<String> id, @Nullable MfaTotpState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("vault:index/mfaTotp:MfaTotp", name, state, makeResourceOptions(options, id));
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
    public static MfaTotp get(String name, Output<String> id, @Nullable MfaTotpState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new MfaTotp(name, id, state, options);
    }
}
