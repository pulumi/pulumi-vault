// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vault.pkiSecret;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.vault.Utilities;
import com.pulumi.vault.pkiSecret.SecretBackendSignArgs;
import com.pulumi.vault.pkiSecret.inputs.SecretBackendSignState;
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
 * import com.pulumi.vault.pkiSecret.SecretBackendSign;
 * import com.pulumi.vault.pkiSecret.SecretBackendSignArgs;
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
 *         var test = new SecretBackendSign("test", SecretBackendSignArgs.builder()
 *             .backend(pki.path())
 *             .name(admin.name())
 *             .csr("""
 * -----BEGIN CERTIFICATE REQUEST-----
 * MIIEqDCCApACAQAwYzELMAkGA1UEBhMCQVUxEzARBgNVBAgMClNvbWUtU3RhdGUx
 * ITAfBgNVBAoMGEludGVybmV0IFdpZGdpdHMgUHR5IEx0ZDEcMBoGA1UEAwwTY2Vy
 * dC50ZXN0Lm15LmRvbWFpbjCCAiIwDQYJKoZIhvcNAQEBBQADggIPADCCAgoCggIB
 * AJupYCQ8UVCWII1Zof1c6YcSSaM9hEaDU78cfKP5RoSeH10BvrWRfT+mzCONVpNP
 * CW9Iabtvk6hm0ot6ilnndEyVJbc0g7hdDLBX5BM25D+DGZGJRKUz1V+uBrWmXtIt
 * Vonj7JTDTe7ViH0GDsB7CvqXFGXO2a2cDYBchLkL6vQiFPshxvUsLtwxuy/qdYgy
 * X6ya+AUoZcoQGy1XxNjfH6cPtWSWQGEp1oPR6vL9hU3laTZb3C+VV4jZem+he8/0
 * V+qV6fLG92WTXm2hmf8nrtUqqJ+C7mW/RJod+TviviBadIX0OHXW7k5HVsZood01
 * te8vMRUNJNiZfa9EMIK5oncbQn0LcM3Wo9VrjpL7jREb/4HCS2gswYGv7hzk9cCS
 * kVY4rDucchKbApuI3kfzmO7GFOF5eiSkYZpY/czNn7VVM3WCu6dpOX4+3rhgrZQw
 * kY14L930DaLVRUgve/zKVP2D2GHdEOs+MbV7s96UgigT9pXly/yHPj+1sSYqmnaD
 * 5b7jSeJusmzO/nrwXVGLsnezR87VzHl9Ux9g5s6zh+R+PrZuVxYsLvoUpaasH47O
 * gIcBzSb/6pSGZKAUizmYsHsR1k88dAvsQ+FsUDaNokdi9VndEB4QPmiFmjyLV+0I
 * 1TFoXop4sW11NPz1YCq+IxnYrEaIN3PyhY0GvBJDFY1/AgMBAAGgADANBgkqhkiG
 * 9w0BAQsFAAOCAgEActuqnqS8Y9UF7e08w7tR3FPzGecWreuvxILrlFEZJxiLPFqL
 * It7uJvtypCVQvz6UQzKdBYO7tMpRaWViB8DrWzXNZjLMrg+QHcpveg8C0Ett4scG
 * fnvLk6fTDFYrnGvwHTqiHos5i0y3bFLyS1BGwSpdLAykGtvC+VM8mRyw/Y7CPcKN
 * 77kebY/9xduW1g2uxWLr0x90RuQDv9psPojT+59tRLGSp5Kt0IeD3QtnAZEFE4aN
 * vt+Pd69eg3BgZ8ZeDgoqAw3yppvOkpAFiE5pw2qPZaM4SRphl4d2Lek2zNIMyZqv
 * do5zh356HOgXtDaSg0POnRGrN/Ua+LMCRTg6GEPUnx9uQb/zt8Zu0hIexDGyykp1
 * OGqtWlv/Nc8UYuS38v0BeB6bMPeoqQUjkqs8nHlAEFn0KlgYdtDC+7SdQx6wS4te
 * dBKRNDfC4lS3jYJgs55jHqonZgkpSi3bamlxpfpW0ukGBcmq91wRe4bOw/4uD/vf
 * UwqMWOdCYcU3mdYNjTWy22ORW3SGFQxMBwpUEURCSoeqWr6aJeQ7KAYkx1PrB5T8
 * OTEc13lWf+B0PU9UJuGTsmpIuImPDVd0EVDayr3mT5dDbqTVDbe8ppf2IswABmf0
 * o3DybUeUmknYjl109rdSf+76nuREICHatxXgN3xCMFuBaN4WLO+ksd6Y1Ys=
 * -----END CERTIFICATE REQUEST-----
 *             """)
 *             .commonName("test.my.domain")
 *             .build(), CustomResourceOptions.builder()
 *                 .dependsOn(admin)
 *                 .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 */
@ResourceType(type="vault:pkiSecret/secretBackendSign:SecretBackendSign")
public class SecretBackendSign extends com.pulumi.resources.CustomResource {
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
     * If set to `true`, certs will be renewed if the expiration is within `min_seconds_remaining`. Default `false`
     * 
     */
    @Export(name="autoRenew", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> autoRenew;

    /**
     * @return If set to `true`, certs will be renewed if the expiration is within `min_seconds_remaining`. Default `false`
     * 
     */
    public Output<Optional<Boolean>> autoRenew() {
        return Codegen.optional(this.autoRenew);
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
     * The CA chain
     * 
     */
    @Export(name="caChains", refs={List.class,String.class}, tree="[0,1]")
    private Output<List<String>> caChains;

    /**
     * @return The CA chain
     * 
     */
    public Output<List<String>> caChains() {
        return this.caChains;
    }
    /**
     * A base 64 encoded value or an empty string to associate with the certificate&#39;s serial number. The role&#39;s no_store_metadata must be set to false, otherwise an error is returned when specified.
     * 
     */
    @Export(name="certMetadata", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> certMetadata;

    /**
     * @return A base 64 encoded value or an empty string to associate with the certificate&#39;s serial number. The role&#39;s no_store_metadata must be set to false, otherwise an error is returned when specified.
     * 
     */
    public Output<Optional<String>> certMetadata() {
        return Codegen.optional(this.certMetadata);
    }
    /**
     * The certificate
     * 
     */
    @Export(name="certificate", refs={String.class}, tree="[0]")
    private Output<String> certificate;

    /**
     * @return The certificate
     * 
     */
    public Output<String> certificate() {
        return this.certificate;
    }
    /**
     * CN of certificate to create
     * 
     */
    @Export(name="commonName", refs={String.class}, tree="[0]")
    private Output<String> commonName;

    /**
     * @return CN of certificate to create
     * 
     */
    public Output<String> commonName() {
        return this.commonName;
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
     * The expiration date of the certificate in unix epoch format
     * 
     */
    @Export(name="expiration", refs={Integer.class}, tree="[0]")
    private Output<Integer> expiration;

    /**
     * @return The expiration date of the certificate in unix epoch format
     * 
     */
    public Output<Integer> expiration() {
        return this.expiration;
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
     * Specifies the default issuer of this request. Can
     * be the value `default`, a name, or an issuer ID. Use ACLs to prevent access to
     * the `/pki/issuer/:issuer_ref/{issue,sign}/:name` paths to prevent users
     * overriding the role&#39;s `issuer_ref` value.
     * 
     */
    @Export(name="issuerRef", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> issuerRef;

    /**
     * @return Specifies the default issuer of this request. Can
     * be the value `default`, a name, or an issuer ID. Use ACLs to prevent access to
     * the `/pki/issuer/:issuer_ref/{issue,sign}/:name` paths to prevent users
     * overriding the role&#39;s `issuer_ref` value.
     * 
     */
    public Output<Optional<String>> issuerRef() {
        return Codegen.optional(this.issuerRef);
    }
    /**
     * The issuing CA
     * 
     */
    @Export(name="issuingCa", refs={String.class}, tree="[0]")
    private Output<String> issuingCa;

    /**
     * @return The issuing CA
     * 
     */
    public Output<String> issuingCa() {
        return this.issuingCa;
    }
    /**
     * Generate a new certificate when the expiration is within this number of seconds, default is 604800 (7 days)
     * 
     */
    @Export(name="minSecondsRemaining", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> minSecondsRemaining;

    /**
     * @return Generate a new certificate when the expiration is within this number of seconds, default is 604800 (7 days)
     * 
     */
    public Output<Optional<Integer>> minSecondsRemaining() {
        return Codegen.optional(this.minSecondsRemaining);
    }
    /**
     * Name of the role to create the certificate against
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return Name of the role to create the certificate against
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
     * Set the Not After field of the certificate with specified date value. The value format should be given in UTC format YYYY-MM-ddTHH:MM:SSZ. Supports the Y10K end date for IEEE 802.1AR-2018 standard devices, 9999-12-31T23:59:59Z.
     * 
     */
    @Export(name="notAfter", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> notAfter;

    /**
     * @return Set the Not After field of the certificate with specified date value. The value format should be given in UTC format YYYY-MM-ddTHH:MM:SSZ. Supports the Y10K end date for IEEE 802.1AR-2018 standard devices, 9999-12-31T23:59:59Z.
     * 
     */
    public Output<Optional<String>> notAfter() {
        return Codegen.optional(this.notAfter);
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
     * `true` if the current time (during refresh) is after the start of the early renewal window declared by `min_seconds_remaining`, and `false` otherwise; if `auto_renew` is set to `true` then the provider will plan to replace the certificate once renewal is pending.
     * 
     */
    @Export(name="renewPending", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> renewPending;

    /**
     * @return `true` if the current time (during refresh) is after the start of the early renewal window declared by `min_seconds_remaining`, and `false` otherwise; if `auto_renew` is set to `true` then the provider will plan to replace the certificate once renewal is pending.
     * 
     */
    public Output<Boolean> renewPending() {
        return this.renewPending;
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
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public SecretBackendSign(java.lang.String name) {
        this(name, SecretBackendSignArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public SecretBackendSign(java.lang.String name, SecretBackendSignArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public SecretBackendSign(java.lang.String name, SecretBackendSignArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("vault:pkiSecret/secretBackendSign:SecretBackendSign", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private SecretBackendSign(java.lang.String name, Output<java.lang.String> id, @Nullable SecretBackendSignState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("vault:pkiSecret/secretBackendSign:SecretBackendSign", name, state, makeResourceOptions(options, id), false);
    }

    private static SecretBackendSignArgs makeArgs(SecretBackendSignArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? SecretBackendSignArgs.Empty : args;
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
    public static SecretBackendSign get(java.lang.String name, Output<java.lang.String> id, @Nullable SecretBackendSignState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new SecretBackendSign(name, id, state, options);
    }
}
