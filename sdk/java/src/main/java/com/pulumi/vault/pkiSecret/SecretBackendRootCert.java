// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vault.pkiSecret;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.vault.Utilities;
import com.pulumi.vault.pkiSecret.SecretBackendRootCertArgs;
import com.pulumi.vault.pkiSecret.inputs.SecretBackendRootCertState;
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
 * import com.pulumi.vault.pkiSecret.SecretBackendRootCert;
 * import com.pulumi.vault.pkiSecret.SecretBackendRootCertArgs;
 * import com.pulumi.resources.CustomResourceOptions;
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
 *         var test = new SecretBackendRootCert("test", SecretBackendRootCertArgs.builder()
 *             .backend(pki.path())
 *             .type("internal")
 *             .commonName("Root CA")
 *             .ttl("315360000")
 *             .format("pem")
 *             .privateKeyFormat("der")
 *             .keyType("rsa")
 *             .keyBits(4096)
 *             .excludeCnFromSans(true)
 *             .ou("My OU")
 *             .organization("My organization")
 *             .build(), CustomResourceOptions.builder()
 *                 .dependsOn(pki)
 *                 .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 */
@ResourceType(type="vault:pkiSecret/secretBackendRootCert:SecretBackendRootCert")
public class SecretBackendRootCert extends com.pulumi.resources.CustomResource {
    /**
     * List of alternative names
     * 
     */
    @Export(name="altNames", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> altNames;

    /**
     * @return List of alternative names
     * 
     */
    public Output<Optional<List<String>>> altNames() {
        return Codegen.optional(this.altNames);
    }
    /**
     * The PKI secret backend the resource belongs to.
     * 
     */
    @Export(name="backend", refs={String.class}, tree="[0]")
    private Output<String> backend;

    /**
     * @return The PKI secret backend the resource belongs to.
     * 
     */
    public Output<String> backend() {
        return this.backend;
    }
    /**
     * The certificate.
     * 
     */
    @Export(name="certificate", refs={String.class}, tree="[0]")
    private Output<String> certificate;

    /**
     * @return The certificate.
     * 
     */
    public Output<String> certificate() {
        return this.certificate;
    }
    /**
     * CN of intermediate to create
     * 
     */
    @Export(name="commonName", refs={String.class}, tree="[0]")
    private Output<String> commonName;

    /**
     * @return CN of intermediate to create
     * 
     */
    public Output<String> commonName() {
        return this.commonName;
    }
    /**
     * The country
     * 
     */
    @Export(name="country", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> country;

    /**
     * @return The country
     * 
     */
    public Output<Optional<String>> country() {
        return Codegen.optional(this.country);
    }
    /**
     * Flag to exclude CN from SANs
     * 
     */
    @Export(name="excludeCnFromSans", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> excludeCnFromSans;

    /**
     * @return Flag to exclude CN from SANs
     * 
     */
    public Output<Optional<Boolean>> excludeCnFromSans() {
        return Codegen.optional(this.excludeCnFromSans);
    }
    /**
     * The format of data
     * 
     */
    @Export(name="format", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> format;

    /**
     * @return The format of data
     * 
     */
    public Output<Optional<String>> format() {
        return Codegen.optional(this.format);
    }
    /**
     * List of alternative IPs
     * 
     */
    @Export(name="ipSans", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> ipSans;

    /**
     * @return List of alternative IPs
     * 
     */
    public Output<Optional<List<String>>> ipSans() {
        return Codegen.optional(this.ipSans);
    }
    /**
     * The ID of the generated issuer.
     * 
     */
    @Export(name="issuerId", refs={String.class}, tree="[0]")
    private Output<String> issuerId;

    /**
     * @return The ID of the generated issuer.
     * 
     */
    public Output<String> issuerId() {
        return this.issuerId;
    }
    /**
     * Provides a name to the specified issuer. The name must be unique
     * across all issuers and not be the reserved value `default`
     * 
     */
    @Export(name="issuerName", refs={String.class}, tree="[0]")
    private Output<String> issuerName;

    /**
     * @return Provides a name to the specified issuer. The name must be unique
     * across all issuers and not be the reserved value `default`
     * 
     */
    public Output<String> issuerName() {
        return this.issuerName;
    }
    /**
     * The issuing CA certificate.
     * 
     */
    @Export(name="issuingCa", refs={String.class}, tree="[0]")
    private Output<String> issuingCa;

    /**
     * @return The issuing CA certificate.
     * 
     */
    public Output<String> issuingCa() {
        return this.issuingCa;
    }
    /**
     * The number of bits to use
     * 
     */
    @Export(name="keyBits", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> keyBits;

    /**
     * @return The number of bits to use
     * 
     */
    public Output<Optional<Integer>> keyBits() {
        return Codegen.optional(this.keyBits);
    }
    /**
     * The ID of the generated key.
     * 
     */
    @Export(name="keyId", refs={String.class}, tree="[0]")
    private Output<String> keyId;

    /**
     * @return The ID of the generated key.
     * 
     */
    public Output<String> keyId() {
        return this.keyId;
    }
    /**
     * When a new key is created with this request, optionally specifies
     * the name for this. The global ref `default` may not be used as a name.
     * 
     */
    @Export(name="keyName", refs={String.class}, tree="[0]")
    private Output<String> keyName;

    /**
     * @return When a new key is created with this request, optionally specifies
     * the name for this. The global ref `default` may not be used as a name.
     * 
     */
    public Output<String> keyName() {
        return this.keyName;
    }
    /**
     * Specifies the key (either default, by name, or by identifier) to use
     * for generating this request. Only suitable for `type=existing` requests.
     * 
     */
    @Export(name="keyRef", refs={String.class}, tree="[0]")
    private Output<String> keyRef;

    /**
     * @return Specifies the key (either default, by name, or by identifier) to use
     * for generating this request. Only suitable for `type=existing` requests.
     * 
     */
    public Output<String> keyRef() {
        return this.keyRef;
    }
    /**
     * The desired key type
     * 
     */
    @Export(name="keyType", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> keyType;

    /**
     * @return The desired key type
     * 
     */
    public Output<Optional<String>> keyType() {
        return Codegen.optional(this.keyType);
    }
    /**
     * The locality
     * 
     */
    @Export(name="locality", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> locality;

    /**
     * @return The locality
     * 
     */
    public Output<Optional<String>> locality() {
        return Codegen.optional(this.locality);
    }
    /**
     * The ID of the previously configured managed key. This field is
     * required if `type` is `kms` and it conflicts with `managed_key_name`
     * 
     */
    @Export(name="managedKeyId", refs={String.class}, tree="[0]")
    private Output<String> managedKeyId;

    /**
     * @return The ID of the previously configured managed key. This field is
     * required if `type` is `kms` and it conflicts with `managed_key_name`
     * 
     */
    public Output<String> managedKeyId() {
        return this.managedKeyId;
    }
    /**
     * The name of the previously configured managed key. This field is
     * required if `type` is `kms`  and it conflicts with `managed_key_id`
     * 
     */
    @Export(name="managedKeyName", refs={String.class}, tree="[0]")
    private Output<String> managedKeyName;

    /**
     * @return The name of the previously configured managed key. This field is
     * required if `type` is `kms`  and it conflicts with `managed_key_id`
     * 
     */
    public Output<String> managedKeyName() {
        return this.managedKeyName;
    }
    /**
     * The maximum path length to encode in the generated certificate
     * 
     */
    @Export(name="maxPathLength", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> maxPathLength;

    /**
     * @return The maximum path length to encode in the generated certificate
     * 
     */
    public Output<Optional<Integer>> maxPathLength() {
        return Codegen.optional(this.maxPathLength);
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
     * The organization
     * 
     */
    @Export(name="organization", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> organization;

    /**
     * @return The organization
     * 
     */
    public Output<Optional<String>> organization() {
        return Codegen.optional(this.organization);
    }
    /**
     * List of other SANs
     * 
     */
    @Export(name="otherSans", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> otherSans;

    /**
     * @return List of other SANs
     * 
     */
    public Output<Optional<List<String>>> otherSans() {
        return Codegen.optional(this.otherSans);
    }
    /**
     * The organization unit
     * 
     */
    @Export(name="ou", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> ou;

    /**
     * @return The organization unit
     * 
     */
    public Output<Optional<String>> ou() {
        return Codegen.optional(this.ou);
    }
    /**
     * List of domains for which certificates are allowed to be issued
     * 
     */
    @Export(name="permittedDnsDomains", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> permittedDnsDomains;

    /**
     * @return List of domains for which certificates are allowed to be issued
     * 
     */
    public Output<Optional<List<String>>> permittedDnsDomains() {
        return Codegen.optional(this.permittedDnsDomains);
    }
    /**
     * The postal code
     * 
     */
    @Export(name="postalCode", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> postalCode;

    /**
     * @return The postal code
     * 
     */
    public Output<Optional<String>> postalCode() {
        return Codegen.optional(this.postalCode);
    }
    /**
     * The private key format
     * 
     */
    @Export(name="privateKeyFormat", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> privateKeyFormat;

    /**
     * @return The private key format
     * 
     */
    public Output<Optional<String>> privateKeyFormat() {
        return Codegen.optional(this.privateKeyFormat);
    }
    /**
     * The province
     * 
     */
    @Export(name="province", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> province;

    /**
     * @return The province
     * 
     */
    public Output<Optional<String>> province() {
        return Codegen.optional(this.province);
    }
    /**
     * The certificate&#39;s serial number, hex formatted.
     * 
     */
    @Export(name="serialNumber", refs={String.class}, tree="[0]")
    private Output<String> serialNumber;

    /**
     * @return The certificate&#39;s serial number, hex formatted.
     * 
     */
    public Output<String> serialNumber() {
        return this.serialNumber;
    }
    /**
     * The street address
     * 
     */
    @Export(name="streetAddress", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> streetAddress;

    /**
     * @return The street address
     * 
     */
    public Output<Optional<String>> streetAddress() {
        return Codegen.optional(this.streetAddress);
    }
    /**
     * Time to live
     * 
     */
    @Export(name="ttl", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> ttl;

    /**
     * @return Time to live
     * 
     */
    public Output<Optional<String>> ttl() {
        return Codegen.optional(this.ttl);
    }
    /**
     * Type of intermediate to create. Must be either \&#34;exported\&#34;, \&#34;internal\&#34;
     * or \&#34;kms\&#34;
     * 
     */
    @Export(name="type", refs={String.class}, tree="[0]")
    private Output<String> type;

    /**
     * @return Type of intermediate to create. Must be either \&#34;exported\&#34;, \&#34;internal\&#34;
     * or \&#34;kms\&#34;
     * 
     */
    public Output<String> type() {
        return this.type;
    }
    /**
     * List of alternative URIs
     * 
     */
    @Export(name="uriSans", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> uriSans;

    /**
     * @return List of alternative URIs
     * 
     */
    public Output<Optional<List<String>>> uriSans() {
        return Codegen.optional(this.uriSans);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public SecretBackendRootCert(java.lang.String name) {
        this(name, SecretBackendRootCertArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public SecretBackendRootCert(java.lang.String name, SecretBackendRootCertArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public SecretBackendRootCert(java.lang.String name, SecretBackendRootCertArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("vault:pkiSecret/secretBackendRootCert:SecretBackendRootCert", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private SecretBackendRootCert(java.lang.String name, Output<java.lang.String> id, @Nullable SecretBackendRootCertState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("vault:pkiSecret/secretBackendRootCert:SecretBackendRootCert", name, state, makeResourceOptions(options, id), false);
    }

    private static SecretBackendRootCertArgs makeArgs(SecretBackendRootCertArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? SecretBackendRootCertArgs.Empty : args;
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
    public static SecretBackendRootCert get(java.lang.String name, Output<java.lang.String> id, @Nullable SecretBackendRootCertState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new SecretBackendRootCert(name, id, state, options);
    }
}
