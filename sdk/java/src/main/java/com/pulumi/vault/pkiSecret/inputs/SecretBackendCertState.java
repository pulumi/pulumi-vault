// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vault.pkiSecret.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class SecretBackendCertState extends com.pulumi.resources.ResourceArgs {

    public static final SecretBackendCertState Empty = new SecretBackendCertState();

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
    @Import(name="backend")
    private @Nullable Output<String> backend;

    /**
     * @return The PKI secret backend the resource belongs to.
     * 
     */
    public Optional<Output<String>> backend() {
        return Optional.ofNullable(this.backend);
    }

    /**
     * The CA chain
     * 
     */
    @Import(name="caChain")
    private @Nullable Output<String> caChain;

    /**
     * @return The CA chain
     * 
     */
    public Optional<Output<String>> caChain() {
        return Optional.ofNullable(this.caChain);
    }

    /**
     * A base 64 encoded value or an empty string to associate with the certificate&#39;s serial number. The role&#39;s no_store_metadata must be set to false, otherwise an error is returned when specified.
     * 
     */
    @Import(name="certMetadata")
    private @Nullable Output<String> certMetadata;

    /**
     * @return A base 64 encoded value or an empty string to associate with the certificate&#39;s serial number. The role&#39;s no_store_metadata must be set to false, otherwise an error is returned when specified.
     * 
     */
    public Optional<Output<String>> certMetadata() {
        return Optional.ofNullable(this.certMetadata);
    }

    /**
     * The certificate
     * 
     */
    @Import(name="certificate")
    private @Nullable Output<String> certificate;

    /**
     * @return The certificate
     * 
     */
    public Optional<Output<String>> certificate() {
        return Optional.ofNullable(this.certificate);
    }

    /**
     * CN of certificate to create
     * 
     */
    @Import(name="commonName")
    private @Nullable Output<String> commonName;

    /**
     * @return CN of certificate to create
     * 
     */
    public Optional<Output<String>> commonName() {
        return Optional.ofNullable(this.commonName);
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
     * The expiration date of the certificate in unix epoch format
     * 
     */
    @Import(name="expiration")
    private @Nullable Output<Integer> expiration;

    /**
     * @return The expiration date of the certificate in unix epoch format
     * 
     */
    public Optional<Output<Integer>> expiration() {
        return Optional.ofNullable(this.expiration);
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
     * Specifies the default issuer of this request.
     * 
     */
    @Import(name="issuerRef")
    private @Nullable Output<String> issuerRef;

    /**
     * @return Specifies the default issuer of this request.
     * 
     */
    public Optional<Output<String>> issuerRef() {
        return Optional.ofNullable(this.issuerRef);
    }

    /**
     * The issuing CA
     * 
     */
    @Import(name="issuingCa")
    private @Nullable Output<String> issuingCa;

    /**
     * @return The issuing CA
     * 
     */
    public Optional<Output<String>> issuingCa() {
        return Optional.ofNullable(this.issuingCa);
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
     * Set the Not After field of the certificate with specified date value. The value format should be given in UTC format YYYY-MM-ddTHH:MM:SSZ. Supports the Y10K end date for IEEE 802.1AR-2018 standard devices, 9999-12-31T23:59:59Z.
     * 
     */
    @Import(name="notAfter")
    private @Nullable Output<String> notAfter;

    /**
     * @return Set the Not After field of the certificate with specified date value. The value format should be given in UTC format YYYY-MM-ddTHH:MM:SSZ. Supports the Y10K end date for IEEE 802.1AR-2018 standard devices, 9999-12-31T23:59:59Z.
     * 
     */
    public Optional<Output<String>> notAfter() {
        return Optional.ofNullable(this.notAfter);
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
     * The private key
     * 
     */
    @Import(name="privateKey")
    private @Nullable Output<String> privateKey;

    /**
     * @return The private key
     * 
     */
    public Optional<Output<String>> privateKey() {
        return Optional.ofNullable(this.privateKey);
    }

    /**
     * The private key format
     * 
     */
    @Import(name="privateKeyFormat")
    private @Nullable Output<String> privateKeyFormat;

    /**
     * @return The private key format
     * 
     */
    public Optional<Output<String>> privateKeyFormat() {
        return Optional.ofNullable(this.privateKeyFormat);
    }

    /**
     * The private key type
     * 
     */
    @Import(name="privateKeyType")
    private @Nullable Output<String> privateKeyType;

    /**
     * @return The private key type
     * 
     */
    public Optional<Output<String>> privateKeyType() {
        return Optional.ofNullable(this.privateKeyType);
    }

    /**
     * `true` if the current time (during refresh) is after the start of the early renewal window declared by `min_seconds_remaining`, and `false` otherwise; if `auto_renew` is set to `true` then the provider will plan to replace the certificate once renewal is pending.
     * 
     */
    @Import(name="renewPending")
    private @Nullable Output<Boolean> renewPending;

    /**
     * @return `true` if the current time (during refresh) is after the start of the early renewal window declared by `min_seconds_remaining`, and `false` otherwise; if `auto_renew` is set to `true` then the provider will plan to replace the certificate once renewal is pending.
     * 
     */
    public Optional<Output<Boolean>> renewPending() {
        return Optional.ofNullable(this.renewPending);
    }

    /**
     * If set to `true`, the certificate will be revoked on resource destruction using the `revoke` PKI API. Conflicts with `revoke_with_key`. Default `false`.
     * 
     */
    @Import(name="revoke")
    private @Nullable Output<Boolean> revoke;

    /**
     * @return If set to `true`, the certificate will be revoked on resource destruction using the `revoke` PKI API. Conflicts with `revoke_with_key`. Default `false`.
     * 
     */
    public Optional<Output<Boolean>> revoke() {
        return Optional.ofNullable(this.revoke);
    }

    /**
     * If set to `true`, the certificate will be revoked on resource destruction using the `revoke-with-key` PKI API. Conflicts with `revoke`. Default `false`
     * 
     */
    @Import(name="revokeWithKey")
    private @Nullable Output<Boolean> revokeWithKey;

    /**
     * @return If set to `true`, the certificate will be revoked on resource destruction using the `revoke-with-key` PKI API. Conflicts with `revoke`. Default `false`
     * 
     */
    public Optional<Output<Boolean>> revokeWithKey() {
        return Optional.ofNullable(this.revokeWithKey);
    }

    /**
     * The serial number
     * 
     */
    @Import(name="serialNumber")
    private @Nullable Output<String> serialNumber;

    /**
     * @return The serial number
     * 
     */
    public Optional<Output<String>> serialNumber() {
        return Optional.ofNullable(this.serialNumber);
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

    /**
     * List of Subject User IDs
     * 
     */
    @Import(name="userIds")
    private @Nullable Output<List<String>> userIds;

    /**
     * @return List of Subject User IDs
     * 
     */
    public Optional<Output<List<String>>> userIds() {
        return Optional.ofNullable(this.userIds);
    }

    private SecretBackendCertState() {}

    private SecretBackendCertState(SecretBackendCertState $) {
        this.altNames = $.altNames;
        this.autoRenew = $.autoRenew;
        this.backend = $.backend;
        this.caChain = $.caChain;
        this.certMetadata = $.certMetadata;
        this.certificate = $.certificate;
        this.commonName = $.commonName;
        this.excludeCnFromSans = $.excludeCnFromSans;
        this.expiration = $.expiration;
        this.format = $.format;
        this.ipSans = $.ipSans;
        this.issuerRef = $.issuerRef;
        this.issuingCa = $.issuingCa;
        this.minSecondsRemaining = $.minSecondsRemaining;
        this.name = $.name;
        this.namespace = $.namespace;
        this.notAfter = $.notAfter;
        this.otherSans = $.otherSans;
        this.privateKey = $.privateKey;
        this.privateKeyFormat = $.privateKeyFormat;
        this.privateKeyType = $.privateKeyType;
        this.renewPending = $.renewPending;
        this.revoke = $.revoke;
        this.revokeWithKey = $.revokeWithKey;
        this.serialNumber = $.serialNumber;
        this.ttl = $.ttl;
        this.uriSans = $.uriSans;
        this.userIds = $.userIds;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(SecretBackendCertState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private SecretBackendCertState $;

        public Builder() {
            $ = new SecretBackendCertState();
        }

        public Builder(SecretBackendCertState defaults) {
            $ = new SecretBackendCertState(Objects.requireNonNull(defaults));
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
        public Builder backend(@Nullable Output<String> backend) {
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
         * @param caChain The CA chain
         * 
         * @return builder
         * 
         */
        public Builder caChain(@Nullable Output<String> caChain) {
            $.caChain = caChain;
            return this;
        }

        /**
         * @param caChain The CA chain
         * 
         * @return builder
         * 
         */
        public Builder caChain(String caChain) {
            return caChain(Output.of(caChain));
        }

        /**
         * @param certMetadata A base 64 encoded value or an empty string to associate with the certificate&#39;s serial number. The role&#39;s no_store_metadata must be set to false, otherwise an error is returned when specified.
         * 
         * @return builder
         * 
         */
        public Builder certMetadata(@Nullable Output<String> certMetadata) {
            $.certMetadata = certMetadata;
            return this;
        }

        /**
         * @param certMetadata A base 64 encoded value or an empty string to associate with the certificate&#39;s serial number. The role&#39;s no_store_metadata must be set to false, otherwise an error is returned when specified.
         * 
         * @return builder
         * 
         */
        public Builder certMetadata(String certMetadata) {
            return certMetadata(Output.of(certMetadata));
        }

        /**
         * @param certificate The certificate
         * 
         * @return builder
         * 
         */
        public Builder certificate(@Nullable Output<String> certificate) {
            $.certificate = certificate;
            return this;
        }

        /**
         * @param certificate The certificate
         * 
         * @return builder
         * 
         */
        public Builder certificate(String certificate) {
            return certificate(Output.of(certificate));
        }

        /**
         * @param commonName CN of certificate to create
         * 
         * @return builder
         * 
         */
        public Builder commonName(@Nullable Output<String> commonName) {
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
         * @param expiration The expiration date of the certificate in unix epoch format
         * 
         * @return builder
         * 
         */
        public Builder expiration(@Nullable Output<Integer> expiration) {
            $.expiration = expiration;
            return this;
        }

        /**
         * @param expiration The expiration date of the certificate in unix epoch format
         * 
         * @return builder
         * 
         */
        public Builder expiration(Integer expiration) {
            return expiration(Output.of(expiration));
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
         * @param issuerRef Specifies the default issuer of this request.
         * 
         * @return builder
         * 
         */
        public Builder issuerRef(@Nullable Output<String> issuerRef) {
            $.issuerRef = issuerRef;
            return this;
        }

        /**
         * @param issuerRef Specifies the default issuer of this request.
         * 
         * @return builder
         * 
         */
        public Builder issuerRef(String issuerRef) {
            return issuerRef(Output.of(issuerRef));
        }

        /**
         * @param issuingCa The issuing CA
         * 
         * @return builder
         * 
         */
        public Builder issuingCa(@Nullable Output<String> issuingCa) {
            $.issuingCa = issuingCa;
            return this;
        }

        /**
         * @param issuingCa The issuing CA
         * 
         * @return builder
         * 
         */
        public Builder issuingCa(String issuingCa) {
            return issuingCa(Output.of(issuingCa));
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
         * @param notAfter Set the Not After field of the certificate with specified date value. The value format should be given in UTC format YYYY-MM-ddTHH:MM:SSZ. Supports the Y10K end date for IEEE 802.1AR-2018 standard devices, 9999-12-31T23:59:59Z.
         * 
         * @return builder
         * 
         */
        public Builder notAfter(@Nullable Output<String> notAfter) {
            $.notAfter = notAfter;
            return this;
        }

        /**
         * @param notAfter Set the Not After field of the certificate with specified date value. The value format should be given in UTC format YYYY-MM-ddTHH:MM:SSZ. Supports the Y10K end date for IEEE 802.1AR-2018 standard devices, 9999-12-31T23:59:59Z.
         * 
         * @return builder
         * 
         */
        public Builder notAfter(String notAfter) {
            return notAfter(Output.of(notAfter));
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
         * @param privateKey The private key
         * 
         * @return builder
         * 
         */
        public Builder privateKey(@Nullable Output<String> privateKey) {
            $.privateKey = privateKey;
            return this;
        }

        /**
         * @param privateKey The private key
         * 
         * @return builder
         * 
         */
        public Builder privateKey(String privateKey) {
            return privateKey(Output.of(privateKey));
        }

        /**
         * @param privateKeyFormat The private key format
         * 
         * @return builder
         * 
         */
        public Builder privateKeyFormat(@Nullable Output<String> privateKeyFormat) {
            $.privateKeyFormat = privateKeyFormat;
            return this;
        }

        /**
         * @param privateKeyFormat The private key format
         * 
         * @return builder
         * 
         */
        public Builder privateKeyFormat(String privateKeyFormat) {
            return privateKeyFormat(Output.of(privateKeyFormat));
        }

        /**
         * @param privateKeyType The private key type
         * 
         * @return builder
         * 
         */
        public Builder privateKeyType(@Nullable Output<String> privateKeyType) {
            $.privateKeyType = privateKeyType;
            return this;
        }

        /**
         * @param privateKeyType The private key type
         * 
         * @return builder
         * 
         */
        public Builder privateKeyType(String privateKeyType) {
            return privateKeyType(Output.of(privateKeyType));
        }

        /**
         * @param renewPending `true` if the current time (during refresh) is after the start of the early renewal window declared by `min_seconds_remaining`, and `false` otherwise; if `auto_renew` is set to `true` then the provider will plan to replace the certificate once renewal is pending.
         * 
         * @return builder
         * 
         */
        public Builder renewPending(@Nullable Output<Boolean> renewPending) {
            $.renewPending = renewPending;
            return this;
        }

        /**
         * @param renewPending `true` if the current time (during refresh) is after the start of the early renewal window declared by `min_seconds_remaining`, and `false` otherwise; if `auto_renew` is set to `true` then the provider will plan to replace the certificate once renewal is pending.
         * 
         * @return builder
         * 
         */
        public Builder renewPending(Boolean renewPending) {
            return renewPending(Output.of(renewPending));
        }

        /**
         * @param revoke If set to `true`, the certificate will be revoked on resource destruction using the `revoke` PKI API. Conflicts with `revoke_with_key`. Default `false`.
         * 
         * @return builder
         * 
         */
        public Builder revoke(@Nullable Output<Boolean> revoke) {
            $.revoke = revoke;
            return this;
        }

        /**
         * @param revoke If set to `true`, the certificate will be revoked on resource destruction using the `revoke` PKI API. Conflicts with `revoke_with_key`. Default `false`.
         * 
         * @return builder
         * 
         */
        public Builder revoke(Boolean revoke) {
            return revoke(Output.of(revoke));
        }

        /**
         * @param revokeWithKey If set to `true`, the certificate will be revoked on resource destruction using the `revoke-with-key` PKI API. Conflicts with `revoke`. Default `false`
         * 
         * @return builder
         * 
         */
        public Builder revokeWithKey(@Nullable Output<Boolean> revokeWithKey) {
            $.revokeWithKey = revokeWithKey;
            return this;
        }

        /**
         * @param revokeWithKey If set to `true`, the certificate will be revoked on resource destruction using the `revoke-with-key` PKI API. Conflicts with `revoke`. Default `false`
         * 
         * @return builder
         * 
         */
        public Builder revokeWithKey(Boolean revokeWithKey) {
            return revokeWithKey(Output.of(revokeWithKey));
        }

        /**
         * @param serialNumber The serial number
         * 
         * @return builder
         * 
         */
        public Builder serialNumber(@Nullable Output<String> serialNumber) {
            $.serialNumber = serialNumber;
            return this;
        }

        /**
         * @param serialNumber The serial number
         * 
         * @return builder
         * 
         */
        public Builder serialNumber(String serialNumber) {
            return serialNumber(Output.of(serialNumber));
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

        /**
         * @param userIds List of Subject User IDs
         * 
         * @return builder
         * 
         */
        public Builder userIds(@Nullable Output<List<String>> userIds) {
            $.userIds = userIds;
            return this;
        }

        /**
         * @param userIds List of Subject User IDs
         * 
         * @return builder
         * 
         */
        public Builder userIds(List<String> userIds) {
            return userIds(Output.of(userIds));
        }

        /**
         * @param userIds List of Subject User IDs
         * 
         * @return builder
         * 
         */
        public Builder userIds(String... userIds) {
            return userIds(List.of(userIds));
        }

        public SecretBackendCertState build() {
            return $;
        }
    }

}
