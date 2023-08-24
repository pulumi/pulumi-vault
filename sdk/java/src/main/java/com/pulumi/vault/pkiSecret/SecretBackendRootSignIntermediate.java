// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vault.pkiSecret;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.vault.Utilities;
import com.pulumi.vault.pkiSecret.SecretBackendRootSignIntermediateArgs;
import com.pulumi.vault.pkiSecret.inputs.SecretBackendRootSignIntermediateState;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Creates PKI certificate.
 * 
 * ## Example Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.vault.pkiSecret.SecretBackendRootSignIntermediate;
 * import com.pulumi.vault.pkiSecret.SecretBackendRootSignIntermediateArgs;
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
 *         var root = new SecretBackendRootSignIntermediate(&#34;root&#34;, SecretBackendRootSignIntermediateArgs.builder()        
 *             .backend(vault_mount.root().path())
 *             .csr(vault_pki_secret_backend_intermediate_cert_request.intermediate().csr())
 *             .commonName(&#34;Intermediate CA&#34;)
 *             .excludeCnFromSans(true)
 *             .ou(&#34;My OU&#34;)
 *             .organization(&#34;My organization&#34;)
 *             .build(), CustomResourceOptions.builder()
 *                 .dependsOn(vault_pki_secret_backend_intermediate_cert_request.intermediate())
 *                 .build());
 * 
 *     }
 * }
 * ```
 * ## Deprecations
 * 
 * * `serial` - Use `serial_number` instead.
 * 
 */
@ResourceType(type="vault:pkiSecret/secretBackendRootSignIntermediate:SecretBackendRootSignIntermediate")
public class SecretBackendRootSignIntermediate extends com.pulumi.resources.CustomResource {
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
     * A list of the issuing and intermediate CA certificates in the `format` specified.
     * 
     */
    @Export(name="caChains", refs={List.class,String.class}, tree="[0,1]")
    private Output<List<String>> caChains;

    /**
     * @return A list of the issuing and intermediate CA certificates in the `format` specified.
     * 
     */
    public Output<List<String>> caChains() {
        return this.caChains;
    }
    /**
     * The intermediate CA certificate in the `format` specified.
     * 
     */
    @Export(name="certificate", refs={String.class}, tree="[0]")
    private Output<String> certificate;

    /**
     * @return The intermediate CA certificate in the `format` specified.
     * 
     */
    public Output<String> certificate() {
        return this.certificate;
    }
    /**
     * The concatenation of the intermediate CA and the issuing CA certificates (PEM encoded).
     * Requires the `format` to be set to any of: pem, pem_bundle. The value will be empty for all other formats.
     * 
     */
    @Export(name="certificateBundle", refs={String.class}, tree="[0]")
    private Output<String> certificateBundle;

    /**
     * @return The concatenation of the intermediate CA and the issuing CA certificates (PEM encoded).
     * Requires the `format` to be set to any of: pem, pem_bundle. The value will be empty for all other formats.
     * 
     */
    public Output<String> certificateBundle() {
        return this.certificateBundle;
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
     * The CSR
     * 
     */
    @Export(name="csr", refs={String.class}, tree="[0]")
    private Output<String> csr;

    /**
     * @return The CSR
     * 
     */
    public Output<String> csr() {
        return this.csr;
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
     * Specifies the default issuer of this request. May
     * be the value `default`, a name, or an issuer ID. Use ACLs to prevent access to
     * the `/pki/issuer/:issuer_ref/{issue,sign}/:name` paths to prevent users
     * overriding the role&#39;s `issuer_ref` value.
     * 
     */
    @Export(name="issuerRef", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> issuerRef;

    /**
     * @return Specifies the default issuer of this request. May
     * be the value `default`, a name, or an issuer ID. Use ACLs to prevent access to
     * the `/pki/issuer/:issuer_ref/{issue,sign}/:name` paths to prevent users
     * overriding the role&#39;s `issuer_ref` value.
     * 
     */
    public Output<Optional<String>> issuerRef() {
        return Codegen.optional(this.issuerRef);
    }
    /**
     * The issuing CA certificate in the `format` specified.
     * 
     */
    @Export(name="issuingCa", refs={String.class}, tree="[0]")
    private Output<String> issuingCa;

    /**
     * @return The issuing CA certificate in the `format` specified.
     * 
     */
    public Output<String> issuingCa() {
        return this.issuingCa;
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
     * If set to `true`, the certificate will be revoked on resource destruction.
     * 
     */
    @Export(name="revoke", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> revoke;

    /**
     * @return If set to `true`, the certificate will be revoked on resource destruction.
     * 
     */
    public Output<Optional<Boolean>> revoke() {
        return Codegen.optional(this.revoke);
    }
    /**
     * The serial number.
     * 
     * @deprecated
     * Use serial_number instead
     * 
     */
    @Deprecated /* Use serial_number instead */
    @Export(name="serial", refs={String.class}, tree="[0]")
    private Output<String> serial;

    /**
     * @return The serial number.
     * 
     */
    public Output<String> serial() {
        return this.serial;
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
     * Preserve CSR values
     * 
     */
    @Export(name="useCsrValues", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> useCsrValues;

    /**
     * @return Preserve CSR values
     * 
     */
    public Output<Optional<Boolean>> useCsrValues() {
        return Codegen.optional(this.useCsrValues);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public SecretBackendRootSignIntermediate(String name) {
        this(name, SecretBackendRootSignIntermediateArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public SecretBackendRootSignIntermediate(String name, SecretBackendRootSignIntermediateArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public SecretBackendRootSignIntermediate(String name, SecretBackendRootSignIntermediateArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("vault:pkiSecret/secretBackendRootSignIntermediate:SecretBackendRootSignIntermediate", name, args == null ? SecretBackendRootSignIntermediateArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private SecretBackendRootSignIntermediate(String name, Output<String> id, @Nullable SecretBackendRootSignIntermediateState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("vault:pkiSecret/secretBackendRootSignIntermediate:SecretBackendRootSignIntermediate", name, state, makeResourceOptions(options, id));
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
    public static SecretBackendRootSignIntermediate get(String name, Output<String> id, @Nullable SecretBackendRootSignIntermediateState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new SecretBackendRootSignIntermediate(name, id, state, options);
    }
}
