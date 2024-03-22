// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vault.pkiSecret;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class SecretBackendIssuerArgs extends com.pulumi.resources.ResourceArgs {

    public static final SecretBackendIssuerArgs Empty = new SecretBackendIssuerArgs();

    /**
     * The path the PKI secret backend is mounted at, with no
     * leading or trailing `/`s.
     * 
     */
    @Import(name="backend", required=true)
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
     * Specifies the URL values for the CRL
     * Distribution Points field.
     * 
     */
    @Import(name="crlDistributionPoints")
    private @Nullable Output<List<String>> crlDistributionPoints;

    /**
     * @return Specifies the URL values for the CRL
     * Distribution Points field.
     * 
     */
    public Optional<Output<List<String>>> crlDistributionPoints() {
        return Optional.ofNullable(this.crlDistributionPoints);
    }

    /**
     * Specifies that the AIA URL values should
     * be templated.
     * 
     */
    @Import(name="enableAiaUrlTemplating")
    private @Nullable Output<Boolean> enableAiaUrlTemplating;

    /**
     * @return Specifies that the AIA URL values should
     * be templated.
     * 
     */
    public Optional<Output<Boolean>> enableAiaUrlTemplating() {
        return Optional.ofNullable(this.enableAiaUrlTemplating);
    }

    /**
     * Name of the issuer.
     * 
     */
    @Import(name="issuerName")
    private @Nullable Output<String> issuerName;

    /**
     * @return Name of the issuer.
     * 
     */
    public Optional<Output<String>> issuerName() {
        return Optional.ofNullable(this.issuerName);
    }

    /**
     * Reference to an existing issuer.
     * 
     */
    @Import(name="issuerRef", required=true)
    private Output<String> issuerRef;

    /**
     * @return Reference to an existing issuer.
     * 
     */
    public Output<String> issuerRef() {
        return this.issuerRef;
    }

    /**
     * Specifies the URL values for the Issuing
     * Certificate field.
     * 
     */
    @Import(name="issuingCertificates")
    private @Nullable Output<List<String>> issuingCertificates;

    /**
     * @return Specifies the URL values for the Issuing
     * Certificate field.
     * 
     */
    public Optional<Output<List<String>>> issuingCertificates() {
        return Optional.ofNullable(this.issuingCertificates);
    }

    /**
     * Behavior of a leaf&#39;s NotAfter field during
     * issuance.
     * 
     */
    @Import(name="leafNotAfterBehavior")
    private @Nullable Output<String> leafNotAfterBehavior;

    /**
     * @return Behavior of a leaf&#39;s NotAfter field during
     * issuance.
     * 
     */
    public Optional<Output<String>> leafNotAfterBehavior() {
        return Optional.ofNullable(this.leafNotAfterBehavior);
    }

    /**
     * Chain of issuer references to build this issuer&#39;s
     * computed CAChain field from, when non-empty.
     * 
     */
    @Import(name="manualChains")
    private @Nullable Output<List<String>> manualChains;

    /**
     * @return Chain of issuer references to build this issuer&#39;s
     * computed CAChain field from, when non-empty.
     * 
     */
    public Optional<Output<List<String>>> manualChains() {
        return Optional.ofNullable(this.manualChains);
    }

    /**
     * The namespace to provision the resource in.
     * The value should not contain leading or trailing forward slashes.
     * The `namespace` is always relative to the provider&#39;s configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
     * *Available only for Vault Enterprise*.
     * 
     */
    @Import(name="namespace")
    private @Nullable Output<String> namespace;

    /**
     * @return The namespace to provision the resource in.
     * The value should not contain leading or trailing forward slashes.
     * The `namespace` is always relative to the provider&#39;s configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
     * *Available only for Vault Enterprise*.
     * 
     */
    public Optional<Output<String>> namespace() {
        return Optional.ofNullable(this.namespace);
    }

    /**
     * Specifies the URL values for the OCSP Servers field.
     * 
     */
    @Import(name="ocspServers")
    private @Nullable Output<List<String>> ocspServers;

    /**
     * @return Specifies the URL values for the OCSP Servers field.
     * 
     */
    public Optional<Output<List<String>>> ocspServers() {
        return Optional.ofNullable(this.ocspServers);
    }

    /**
     * Which signature algorithm to use
     * when building CRLs.
     * 
     */
    @Import(name="revocationSignatureAlgorithm")
    private @Nullable Output<String> revocationSignatureAlgorithm;

    /**
     * @return Which signature algorithm to use
     * when building CRLs.
     * 
     */
    public Optional<Output<String>> revocationSignatureAlgorithm() {
        return Optional.ofNullable(this.revocationSignatureAlgorithm);
    }

    /**
     * Allowed usages for this issuer.
     * 
     */
    @Import(name="usage")
    private @Nullable Output<String> usage;

    /**
     * @return Allowed usages for this issuer.
     * 
     */
    public Optional<Output<String>> usage() {
        return Optional.ofNullable(this.usage);
    }

    private SecretBackendIssuerArgs() {}

    private SecretBackendIssuerArgs(SecretBackendIssuerArgs $) {
        this.backend = $.backend;
        this.crlDistributionPoints = $.crlDistributionPoints;
        this.enableAiaUrlTemplating = $.enableAiaUrlTemplating;
        this.issuerName = $.issuerName;
        this.issuerRef = $.issuerRef;
        this.issuingCertificates = $.issuingCertificates;
        this.leafNotAfterBehavior = $.leafNotAfterBehavior;
        this.manualChains = $.manualChains;
        this.namespace = $.namespace;
        this.ocspServers = $.ocspServers;
        this.revocationSignatureAlgorithm = $.revocationSignatureAlgorithm;
        this.usage = $.usage;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(SecretBackendIssuerArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private SecretBackendIssuerArgs $;

        public Builder() {
            $ = new SecretBackendIssuerArgs();
        }

        public Builder(SecretBackendIssuerArgs defaults) {
            $ = new SecretBackendIssuerArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param backend The path the PKI secret backend is mounted at, with no
         * leading or trailing `/`s.
         * 
         * @return builder
         * 
         */
        public Builder backend(Output<String> backend) {
            $.backend = backend;
            return this;
        }

        /**
         * @param backend The path the PKI secret backend is mounted at, with no
         * leading or trailing `/`s.
         * 
         * @return builder
         * 
         */
        public Builder backend(String backend) {
            return backend(Output.of(backend));
        }

        /**
         * @param crlDistributionPoints Specifies the URL values for the CRL
         * Distribution Points field.
         * 
         * @return builder
         * 
         */
        public Builder crlDistributionPoints(@Nullable Output<List<String>> crlDistributionPoints) {
            $.crlDistributionPoints = crlDistributionPoints;
            return this;
        }

        /**
         * @param crlDistributionPoints Specifies the URL values for the CRL
         * Distribution Points field.
         * 
         * @return builder
         * 
         */
        public Builder crlDistributionPoints(List<String> crlDistributionPoints) {
            return crlDistributionPoints(Output.of(crlDistributionPoints));
        }

        /**
         * @param crlDistributionPoints Specifies the URL values for the CRL
         * Distribution Points field.
         * 
         * @return builder
         * 
         */
        public Builder crlDistributionPoints(String... crlDistributionPoints) {
            return crlDistributionPoints(List.of(crlDistributionPoints));
        }

        /**
         * @param enableAiaUrlTemplating Specifies that the AIA URL values should
         * be templated.
         * 
         * @return builder
         * 
         */
        public Builder enableAiaUrlTemplating(@Nullable Output<Boolean> enableAiaUrlTemplating) {
            $.enableAiaUrlTemplating = enableAiaUrlTemplating;
            return this;
        }

        /**
         * @param enableAiaUrlTemplating Specifies that the AIA URL values should
         * be templated.
         * 
         * @return builder
         * 
         */
        public Builder enableAiaUrlTemplating(Boolean enableAiaUrlTemplating) {
            return enableAiaUrlTemplating(Output.of(enableAiaUrlTemplating));
        }

        /**
         * @param issuerName Name of the issuer.
         * 
         * @return builder
         * 
         */
        public Builder issuerName(@Nullable Output<String> issuerName) {
            $.issuerName = issuerName;
            return this;
        }

        /**
         * @param issuerName Name of the issuer.
         * 
         * @return builder
         * 
         */
        public Builder issuerName(String issuerName) {
            return issuerName(Output.of(issuerName));
        }

        /**
         * @param issuerRef Reference to an existing issuer.
         * 
         * @return builder
         * 
         */
        public Builder issuerRef(Output<String> issuerRef) {
            $.issuerRef = issuerRef;
            return this;
        }

        /**
         * @param issuerRef Reference to an existing issuer.
         * 
         * @return builder
         * 
         */
        public Builder issuerRef(String issuerRef) {
            return issuerRef(Output.of(issuerRef));
        }

        /**
         * @param issuingCertificates Specifies the URL values for the Issuing
         * Certificate field.
         * 
         * @return builder
         * 
         */
        public Builder issuingCertificates(@Nullable Output<List<String>> issuingCertificates) {
            $.issuingCertificates = issuingCertificates;
            return this;
        }

        /**
         * @param issuingCertificates Specifies the URL values for the Issuing
         * Certificate field.
         * 
         * @return builder
         * 
         */
        public Builder issuingCertificates(List<String> issuingCertificates) {
            return issuingCertificates(Output.of(issuingCertificates));
        }

        /**
         * @param issuingCertificates Specifies the URL values for the Issuing
         * Certificate field.
         * 
         * @return builder
         * 
         */
        public Builder issuingCertificates(String... issuingCertificates) {
            return issuingCertificates(List.of(issuingCertificates));
        }

        /**
         * @param leafNotAfterBehavior Behavior of a leaf&#39;s NotAfter field during
         * issuance.
         * 
         * @return builder
         * 
         */
        public Builder leafNotAfterBehavior(@Nullable Output<String> leafNotAfterBehavior) {
            $.leafNotAfterBehavior = leafNotAfterBehavior;
            return this;
        }

        /**
         * @param leafNotAfterBehavior Behavior of a leaf&#39;s NotAfter field during
         * issuance.
         * 
         * @return builder
         * 
         */
        public Builder leafNotAfterBehavior(String leafNotAfterBehavior) {
            return leafNotAfterBehavior(Output.of(leafNotAfterBehavior));
        }

        /**
         * @param manualChains Chain of issuer references to build this issuer&#39;s
         * computed CAChain field from, when non-empty.
         * 
         * @return builder
         * 
         */
        public Builder manualChains(@Nullable Output<List<String>> manualChains) {
            $.manualChains = manualChains;
            return this;
        }

        /**
         * @param manualChains Chain of issuer references to build this issuer&#39;s
         * computed CAChain field from, when non-empty.
         * 
         * @return builder
         * 
         */
        public Builder manualChains(List<String> manualChains) {
            return manualChains(Output.of(manualChains));
        }

        /**
         * @param manualChains Chain of issuer references to build this issuer&#39;s
         * computed CAChain field from, when non-empty.
         * 
         * @return builder
         * 
         */
        public Builder manualChains(String... manualChains) {
            return manualChains(List.of(manualChains));
        }

        /**
         * @param namespace The namespace to provision the resource in.
         * The value should not contain leading or trailing forward slashes.
         * The `namespace` is always relative to the provider&#39;s configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
         * *Available only for Vault Enterprise*.
         * 
         * @return builder
         * 
         */
        public Builder namespace(@Nullable Output<String> namespace) {
            $.namespace = namespace;
            return this;
        }

        /**
         * @param namespace The namespace to provision the resource in.
         * The value should not contain leading or trailing forward slashes.
         * The `namespace` is always relative to the provider&#39;s configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
         * *Available only for Vault Enterprise*.
         * 
         * @return builder
         * 
         */
        public Builder namespace(String namespace) {
            return namespace(Output.of(namespace));
        }

        /**
         * @param ocspServers Specifies the URL values for the OCSP Servers field.
         * 
         * @return builder
         * 
         */
        public Builder ocspServers(@Nullable Output<List<String>> ocspServers) {
            $.ocspServers = ocspServers;
            return this;
        }

        /**
         * @param ocspServers Specifies the URL values for the OCSP Servers field.
         * 
         * @return builder
         * 
         */
        public Builder ocspServers(List<String> ocspServers) {
            return ocspServers(Output.of(ocspServers));
        }

        /**
         * @param ocspServers Specifies the URL values for the OCSP Servers field.
         * 
         * @return builder
         * 
         */
        public Builder ocspServers(String... ocspServers) {
            return ocspServers(List.of(ocspServers));
        }

        /**
         * @param revocationSignatureAlgorithm Which signature algorithm to use
         * when building CRLs.
         * 
         * @return builder
         * 
         */
        public Builder revocationSignatureAlgorithm(@Nullable Output<String> revocationSignatureAlgorithm) {
            $.revocationSignatureAlgorithm = revocationSignatureAlgorithm;
            return this;
        }

        /**
         * @param revocationSignatureAlgorithm Which signature algorithm to use
         * when building CRLs.
         * 
         * @return builder
         * 
         */
        public Builder revocationSignatureAlgorithm(String revocationSignatureAlgorithm) {
            return revocationSignatureAlgorithm(Output.of(revocationSignatureAlgorithm));
        }

        /**
         * @param usage Allowed usages for this issuer.
         * 
         * @return builder
         * 
         */
        public Builder usage(@Nullable Output<String> usage) {
            $.usage = usage;
            return this;
        }

        /**
         * @param usage Allowed usages for this issuer.
         * 
         * @return builder
         * 
         */
        public Builder usage(String usage) {
            return usage(Output.of(usage));
        }

        public SecretBackendIssuerArgs build() {
            if ($.backend == null) {
                throw new MissingRequiredPropertyException("SecretBackendIssuerArgs", "backend");
            }
            if ($.issuerRef == null) {
                throw new MissingRequiredPropertyException("SecretBackendIssuerArgs", "issuerRef");
            }
            return $;
        }
    }

}
