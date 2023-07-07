// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vault.pkiSecret;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class SecretBackendSignArgs extends com.pulumi.resources.ResourceArgs {

    public static final SecretBackendSignArgs Empty = new SecretBackendSignArgs();

    /**
     * List of alternative names
     * 
     */
    @Import(name="altNames")
    private @Nullable Output<List<String>> altNames;

    /**
     * @return List of alternative names
     * 
     */
    public Optional<Output<List<String>>> altNames() {
        return Optional.ofNullable(this.altNames);
    }

    /**
     * If set to `true`, certs will be renewed if the expiration is within `min_seconds_remaining`. Default `false`
     * 
     */
    @Import(name="autoRenew")
    private @Nullable Output<Boolean> autoRenew;

    /**
     * @return If set to `true`, certs will be renewed if the expiration is within `min_seconds_remaining`. Default `false`
     * 
     */
    public Optional<Output<Boolean>> autoRenew() {
        return Optional.ofNullable(this.autoRenew);
    }

    /**
     * The PKI secret backend the resource belongs to.
     * 
     */
    @Import(name="backend", required=true)
    private Output<String> backend;

    /**
     * @return The PKI secret backend the resource belongs to.
     * 
     */
    public Output<String> backend() {
        return this.backend;
    }

    /**
     * CN of certificate to create
     * 
     */
    @Import(name="commonName", required=true)
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
    @Import(name="csr", required=true)
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
    @Import(name="excludeCnFromSans")
    private @Nullable Output<Boolean> excludeCnFromSans;

    /**
     * @return Flag to exclude CN from SANs
     * 
     */
    public Optional<Output<Boolean>> excludeCnFromSans() {
        return Optional.ofNullable(this.excludeCnFromSans);
    }

    /**
     * The format of data
     * 
     */
    @Import(name="format")
    private @Nullable Output<String> format;

    /**
     * @return The format of data
     * 
     */
    public Optional<Output<String>> format() {
        return Optional.ofNullable(this.format);
    }

    /**
     * List of alternative IPs
     * 
     */
    @Import(name="ipSans")
    private @Nullable Output<List<String>> ipSans;

    /**
     * @return List of alternative IPs
     * 
     */
    public Optional<Output<List<String>>> ipSans() {
        return Optional.ofNullable(this.ipSans);
    }

    /**
     * Specifies the default issuer of this request. Can
     * be the value `default`, a name, or an issuer ID. Use ACLs to prevent access to
     * the `/pki/issuer/:issuer_ref/{issue,sign}/:name` paths to prevent users
     * overriding the role&#39;s `issuer_ref` value.
     * 
     */
    @Import(name="issuerRef")
    private @Nullable Output<String> issuerRef;

    /**
     * @return Specifies the default issuer of this request. Can
     * be the value `default`, a name, or an issuer ID. Use ACLs to prevent access to
     * the `/pki/issuer/:issuer_ref/{issue,sign}/:name` paths to prevent users
     * overriding the role&#39;s `issuer_ref` value.
     * 
     */
    public Optional<Output<String>> issuerRef() {
        return Optional.ofNullable(this.issuerRef);
    }

    /**
     * Generate a new certificate when the expiration is within this number of seconds, default is 604800 (7 days)
     * 
     */
    @Import(name="minSecondsRemaining")
    private @Nullable Output<Integer> minSecondsRemaining;

    /**
     * @return Generate a new certificate when the expiration is within this number of seconds, default is 604800 (7 days)
     * 
     */
    public Optional<Output<Integer>> minSecondsRemaining() {
        return Optional.ofNullable(this.minSecondsRemaining);
    }

    /**
     * Name of the role to create the certificate against
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return Name of the role to create the certificate against
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * The namespace to provision the resource in.
     * The value should not contain leading or trailing forward slashes.
     * The `namespace` is always relative to the provider&#39;s configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
     * *Available only for Vault Enterprise*.
     * 
     */
    @Import(name="namespace")
    private @Nullable Output<String> namespace;

    /**
     * @return The namespace to provision the resource in.
     * The value should not contain leading or trailing forward slashes.
     * The `namespace` is always relative to the provider&#39;s configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
     * *Available only for Vault Enterprise*.
     * 
     */
    public Optional<Output<String>> namespace() {
        return Optional.ofNullable(this.namespace);
    }

    /**
     * List of other SANs
     * 
     */
    @Import(name="otherSans")
    private @Nullable Output<List<String>> otherSans;

    /**
     * @return List of other SANs
     * 
     */
    public Optional<Output<List<String>>> otherSans() {
        return Optional.ofNullable(this.otherSans);
    }

    /**
     * Time to live
     * 
     */
    @Import(name="ttl")
    private @Nullable Output<String> ttl;

    /**
     * @return Time to live
     * 
     */
    public Optional<Output<String>> ttl() {
        return Optional.ofNullable(this.ttl);
    }

    /**
     * List of alternative URIs
     * 
     */
    @Import(name="uriSans")
    private @Nullable Output<List<String>> uriSans;

    /**
     * @return List of alternative URIs
     * 
     */
    public Optional<Output<List<String>>> uriSans() {
        return Optional.ofNullable(this.uriSans);
    }

    private SecretBackendSignArgs() {}

    private SecretBackendSignArgs(SecretBackendSignArgs $) {
        this.altNames = $.altNames;
        this.autoRenew = $.autoRenew;
        this.backend = $.backend;
        this.commonName = $.commonName;
        this.csr = $.csr;
        this.excludeCnFromSans = $.excludeCnFromSans;
        this.format = $.format;
        this.ipSans = $.ipSans;
        this.issuerRef = $.issuerRef;
        this.minSecondsRemaining = $.minSecondsRemaining;
        this.name = $.name;
        this.namespace = $.namespace;
        this.otherSans = $.otherSans;
        this.ttl = $.ttl;
        this.uriSans = $.uriSans;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(SecretBackendSignArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private SecretBackendSignArgs $;

        public Builder() {
            $ = new SecretBackendSignArgs();
        }

        public Builder(SecretBackendSignArgs defaults) {
            $ = new SecretBackendSignArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param altNames List of alternative names
         * 
         * @return builder
         * 
         */
        public Builder altNames(@Nullable Output<List<String>> altNames) {
            $.altNames = altNames;
            return this;
        }

        /**
         * @param altNames List of alternative names
         * 
         * @return builder
         * 
         */
        public Builder altNames(List<String> altNames) {
            return altNames(Output.of(altNames));
        }

        /**
         * @param altNames List of alternative names
         * 
         * @return builder
         * 
         */
        public Builder altNames(String... altNames) {
            return altNames(List.of(altNames));
        }

        /**
         * @param autoRenew If set to `true`, certs will be renewed if the expiration is within `min_seconds_remaining`. Default `false`
         * 
         * @return builder
         * 
         */
        public Builder autoRenew(@Nullable Output<Boolean> autoRenew) {
            $.autoRenew = autoRenew;
            return this;
        }

        /**
         * @param autoRenew If set to `true`, certs will be renewed if the expiration is within `min_seconds_remaining`. Default `false`
         * 
         * @return builder
         * 
         */
        public Builder autoRenew(Boolean autoRenew) {
            return autoRenew(Output.of(autoRenew));
        }

        /**
         * @param backend The PKI secret backend the resource belongs to.
         * 
         * @return builder
         * 
         */
        public Builder backend(Output<String> backend) {
            $.backend = backend;
            return this;
        }

        /**
         * @param backend The PKI secret backend the resource belongs to.
         * 
         * @return builder
         * 
         */
        public Builder backend(String backend) {
            return backend(Output.of(backend));
        }

        /**
         * @param commonName CN of certificate to create
         * 
         * @return builder
         * 
         */
        public Builder commonName(Output<String> commonName) {
            $.commonName = commonName;
            return this;
        }

        /**
         * @param commonName CN of certificate to create
         * 
         * @return builder
         * 
         */
        public Builder commonName(String commonName) {
            return commonName(Output.of(commonName));
        }

        /**
         * @param csr The CSR
         * 
         * @return builder
         * 
         */
        public Builder csr(Output<String> csr) {
            $.csr = csr;
            return this;
        }

        /**
         * @param csr The CSR
         * 
         * @return builder
         * 
         */
        public Builder csr(String csr) {
            return csr(Output.of(csr));
        }

        /**
         * @param excludeCnFromSans Flag to exclude CN from SANs
         * 
         * @return builder
         * 
         */
        public Builder excludeCnFromSans(@Nullable Output<Boolean> excludeCnFromSans) {
            $.excludeCnFromSans = excludeCnFromSans;
            return this;
        }

        /**
         * @param excludeCnFromSans Flag to exclude CN from SANs
         * 
         * @return builder
         * 
         */
        public Builder excludeCnFromSans(Boolean excludeCnFromSans) {
            return excludeCnFromSans(Output.of(excludeCnFromSans));
        }

        /**
         * @param format The format of data
         * 
         * @return builder
         * 
         */
        public Builder format(@Nullable Output<String> format) {
            $.format = format;
            return this;
        }

        /**
         * @param format The format of data
         * 
         * @return builder
         * 
         */
        public Builder format(String format) {
            return format(Output.of(format));
        }

        /**
         * @param ipSans List of alternative IPs
         * 
         * @return builder
         * 
         */
        public Builder ipSans(@Nullable Output<List<String>> ipSans) {
            $.ipSans = ipSans;
            return this;
        }

        /**
         * @param ipSans List of alternative IPs
         * 
         * @return builder
         * 
         */
        public Builder ipSans(List<String> ipSans) {
            return ipSans(Output.of(ipSans));
        }

        /**
         * @param ipSans List of alternative IPs
         * 
         * @return builder
         * 
         */
        public Builder ipSans(String... ipSans) {
            return ipSans(List.of(ipSans));
        }

        /**
         * @param issuerRef Specifies the default issuer of this request. Can
         * be the value `default`, a name, or an issuer ID. Use ACLs to prevent access to
         * the `/pki/issuer/:issuer_ref/{issue,sign}/:name` paths to prevent users
         * overriding the role&#39;s `issuer_ref` value.
         * 
         * @return builder
         * 
         */
        public Builder issuerRef(@Nullable Output<String> issuerRef) {
            $.issuerRef = issuerRef;
            return this;
        }

        /**
         * @param issuerRef Specifies the default issuer of this request. Can
         * be the value `default`, a name, or an issuer ID. Use ACLs to prevent access to
         * the `/pki/issuer/:issuer_ref/{issue,sign}/:name` paths to prevent users
         * overriding the role&#39;s `issuer_ref` value.
         * 
         * @return builder
         * 
         */
        public Builder issuerRef(String issuerRef) {
            return issuerRef(Output.of(issuerRef));
        }

        /**
         * @param minSecondsRemaining Generate a new certificate when the expiration is within this number of seconds, default is 604800 (7 days)
         * 
         * @return builder
         * 
         */
        public Builder minSecondsRemaining(@Nullable Output<Integer> minSecondsRemaining) {
            $.minSecondsRemaining = minSecondsRemaining;
            return this;
        }

        /**
         * @param minSecondsRemaining Generate a new certificate when the expiration is within this number of seconds, default is 604800 (7 days)
         * 
         * @return builder
         * 
         */
        public Builder minSecondsRemaining(Integer minSecondsRemaining) {
            return minSecondsRemaining(Output.of(minSecondsRemaining));
        }

        /**
         * @param name Name of the role to create the certificate against
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name Name of the role to create the certificate against
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param namespace The namespace to provision the resource in.
         * The value should not contain leading or trailing forward slashes.
         * The `namespace` is always relative to the provider&#39;s configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
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
         * The `namespace` is always relative to the provider&#39;s configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
         * *Available only for Vault Enterprise*.
         * 
         * @return builder
         * 
         */
        public Builder namespace(String namespace) {
            return namespace(Output.of(namespace));
        }

        /**
         * @param otherSans List of other SANs
         * 
         * @return builder
         * 
         */
        public Builder otherSans(@Nullable Output<List<String>> otherSans) {
            $.otherSans = otherSans;
            return this;
        }

        /**
         * @param otherSans List of other SANs
         * 
         * @return builder
         * 
         */
        public Builder otherSans(List<String> otherSans) {
            return otherSans(Output.of(otherSans));
        }

        /**
         * @param otherSans List of other SANs
         * 
         * @return builder
         * 
         */
        public Builder otherSans(String... otherSans) {
            return otherSans(List.of(otherSans));
        }

        /**
         * @param ttl Time to live
         * 
         * @return builder
         * 
         */
        public Builder ttl(@Nullable Output<String> ttl) {
            $.ttl = ttl;
            return this;
        }

        /**
         * @param ttl Time to live
         * 
         * @return builder
         * 
         */
        public Builder ttl(String ttl) {
            return ttl(Output.of(ttl));
        }

        /**
         * @param uriSans List of alternative URIs
         * 
         * @return builder
         * 
         */
        public Builder uriSans(@Nullable Output<List<String>> uriSans) {
            $.uriSans = uriSans;
            return this;
        }

        /**
         * @param uriSans List of alternative URIs
         * 
         * @return builder
         * 
         */
        public Builder uriSans(List<String> uriSans) {
            return uriSans(Output.of(uriSans));
        }

        /**
         * @param uriSans List of alternative URIs
         * 
         * @return builder
         * 
         */
        public Builder uriSans(String... uriSans) {
            return uriSans(List.of(uriSans));
        }

        public SecretBackendSignArgs build() {
            $.backend = Objects.requireNonNull($.backend, "expected parameter 'backend' to be non-null");
            $.commonName = Objects.requireNonNull($.commonName, "expected parameter 'commonName' to be non-null");
            $.csr = Objects.requireNonNull($.csr, "expected parameter 'csr' to be non-null");
            return $;
        }
    }

}
