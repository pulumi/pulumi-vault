// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vault.pkiSecret;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class SecretBackendRootCertArgs extends com.pulumi.resources.ResourceArgs {

    public static final SecretBackendRootCertArgs Empty = new SecretBackendRootCertArgs();

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
     * CN of intermediate to create
     * 
     */
    @Import(name="commonName", required=true)
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
    @Import(name="country")
    private @Nullable Output<String> country;

    /**
     * @return The country
     * 
     */
    public Optional<Output<String>> country() {
        return Optional.ofNullable(this.country);
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
     * Provides a name to the specified issuer. The name must be unique
     * across all issuers and not be the reserved value `default`
     * 
     */
    @Import(name="issuerName")
    private @Nullable Output<String> issuerName;

    /**
     * @return Provides a name to the specified issuer. The name must be unique
     * across all issuers and not be the reserved value `default`
     * 
     */
    public Optional<Output<String>> issuerName() {
        return Optional.ofNullable(this.issuerName);
    }

    /**
     * The number of bits to use
     * 
     */
    @Import(name="keyBits")
    private @Nullable Output<Integer> keyBits;

    /**
     * @return The number of bits to use
     * 
     */
    public Optional<Output<Integer>> keyBits() {
        return Optional.ofNullable(this.keyBits);
    }

    /**
     * When a new key is created with this request, optionally specifies
     * the name for this. The global ref `default` may not be used as a name.
     * 
     */
    @Import(name="keyName")
    private @Nullable Output<String> keyName;

    /**
     * @return When a new key is created with this request, optionally specifies
     * the name for this. The global ref `default` may not be used as a name.
     * 
     */
    public Optional<Output<String>> keyName() {
        return Optional.ofNullable(this.keyName);
    }

    /**
     * Specifies the key (either default, by name, or by identifier) to use
     * for generating this request. Only suitable for `type=existing` requests.
     * 
     */
    @Import(name="keyRef")
    private @Nullable Output<String> keyRef;

    /**
     * @return Specifies the key (either default, by name, or by identifier) to use
     * for generating this request. Only suitable for `type=existing` requests.
     * 
     */
    public Optional<Output<String>> keyRef() {
        return Optional.ofNullable(this.keyRef);
    }

    /**
     * The desired key type
     * 
     */
    @Import(name="keyType")
    private @Nullable Output<String> keyType;

    /**
     * @return The desired key type
     * 
     */
    public Optional<Output<String>> keyType() {
        return Optional.ofNullable(this.keyType);
    }

    /**
     * The locality
     * 
     */
    @Import(name="locality")
    private @Nullable Output<String> locality;

    /**
     * @return The locality
     * 
     */
    public Optional<Output<String>> locality() {
        return Optional.ofNullable(this.locality);
    }

    /**
     * The ID of the previously configured managed key. This field is
     * required if `type` is `kms` and it conflicts with `managed_key_name`
     * 
     */
    @Import(name="managedKeyId")
    private @Nullable Output<String> managedKeyId;

    /**
     * @return The ID of the previously configured managed key. This field is
     * required if `type` is `kms` and it conflicts with `managed_key_name`
     * 
     */
    public Optional<Output<String>> managedKeyId() {
        return Optional.ofNullable(this.managedKeyId);
    }

    /**
     * The name of the previously configured managed key. This field is
     * required if `type` is `kms`  and it conflicts with `managed_key_id`
     * 
     */
    @Import(name="managedKeyName")
    private @Nullable Output<String> managedKeyName;

    /**
     * @return The name of the previously configured managed key. This field is
     * required if `type` is `kms`  and it conflicts with `managed_key_id`
     * 
     */
    public Optional<Output<String>> managedKeyName() {
        return Optional.ofNullable(this.managedKeyName);
    }

    /**
     * The maximum path length to encode in the generated certificate
     * 
     */
    @Import(name="maxPathLength")
    private @Nullable Output<Integer> maxPathLength;

    /**
     * @return The maximum path length to encode in the generated certificate
     * 
     */
    public Optional<Output<Integer>> maxPathLength() {
        return Optional.ofNullable(this.maxPathLength);
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
     * The organization
     * 
     */
    @Import(name="organization")
    private @Nullable Output<String> organization;

    /**
     * @return The organization
     * 
     */
    public Optional<Output<String>> organization() {
        return Optional.ofNullable(this.organization);
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
     * The organization unit
     * 
     */
    @Import(name="ou")
    private @Nullable Output<String> ou;

    /**
     * @return The organization unit
     * 
     */
    public Optional<Output<String>> ou() {
        return Optional.ofNullable(this.ou);
    }

    /**
     * List of domains for which certificates are allowed to be issued
     * 
     */
    @Import(name="permittedDnsDomains")
    private @Nullable Output<List<String>> permittedDnsDomains;

    /**
     * @return List of domains for which certificates are allowed to be issued
     * 
     */
    public Optional<Output<List<String>>> permittedDnsDomains() {
        return Optional.ofNullable(this.permittedDnsDomains);
    }

    /**
     * The postal code
     * 
     */
    @Import(name="postalCode")
    private @Nullable Output<String> postalCode;

    /**
     * @return The postal code
     * 
     */
    public Optional<Output<String>> postalCode() {
        return Optional.ofNullable(this.postalCode);
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
     * The province
     * 
     */
    @Import(name="province")
    private @Nullable Output<String> province;

    /**
     * @return The province
     * 
     */
    public Optional<Output<String>> province() {
        return Optional.ofNullable(this.province);
    }

    /**
     * The street address
     * 
     */
    @Import(name="streetAddress")
    private @Nullable Output<String> streetAddress;

    /**
     * @return The street address
     * 
     */
    public Optional<Output<String>> streetAddress() {
        return Optional.ofNullable(this.streetAddress);
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
     * Type of intermediate to create. Must be either \&#34;exported\&#34;, \&#34;internal\&#34;
     * or \&#34;kms\&#34;
     * 
     */
    @Import(name="type", required=true)
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
    @Import(name="uriSans")
    private @Nullable Output<List<String>> uriSans;

    /**
     * @return List of alternative URIs
     * 
     */
    public Optional<Output<List<String>>> uriSans() {
        return Optional.ofNullable(this.uriSans);
    }

    private SecretBackendRootCertArgs() {}

    private SecretBackendRootCertArgs(SecretBackendRootCertArgs $) {
        this.altNames = $.altNames;
        this.backend = $.backend;
        this.commonName = $.commonName;
        this.country = $.country;
        this.excludeCnFromSans = $.excludeCnFromSans;
        this.format = $.format;
        this.ipSans = $.ipSans;
        this.issuerName = $.issuerName;
        this.keyBits = $.keyBits;
        this.keyName = $.keyName;
        this.keyRef = $.keyRef;
        this.keyType = $.keyType;
        this.locality = $.locality;
        this.managedKeyId = $.managedKeyId;
        this.managedKeyName = $.managedKeyName;
        this.maxPathLength = $.maxPathLength;
        this.namespace = $.namespace;
        this.organization = $.organization;
        this.otherSans = $.otherSans;
        this.ou = $.ou;
        this.permittedDnsDomains = $.permittedDnsDomains;
        this.postalCode = $.postalCode;
        this.privateKeyFormat = $.privateKeyFormat;
        this.province = $.province;
        this.streetAddress = $.streetAddress;
        this.ttl = $.ttl;
        this.type = $.type;
        this.uriSans = $.uriSans;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(SecretBackendRootCertArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private SecretBackendRootCertArgs $;

        public Builder() {
            $ = new SecretBackendRootCertArgs();
        }

        public Builder(SecretBackendRootCertArgs defaults) {
            $ = new SecretBackendRootCertArgs(Objects.requireNonNull(defaults));
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
         * @param commonName CN of intermediate to create
         * 
         * @return builder
         * 
         */
        public Builder commonName(Output<String> commonName) {
            $.commonName = commonName;
            return this;
        }

        /**
         * @param commonName CN of intermediate to create
         * 
         * @return builder
         * 
         */
        public Builder commonName(String commonName) {
            return commonName(Output.of(commonName));
        }

        /**
         * @param country The country
         * 
         * @return builder
         * 
         */
        public Builder country(@Nullable Output<String> country) {
            $.country = country;
            return this;
        }

        /**
         * @param country The country
         * 
         * @return builder
         * 
         */
        public Builder country(String country) {
            return country(Output.of(country));
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
         * @param issuerName Provides a name to the specified issuer. The name must be unique
         * across all issuers and not be the reserved value `default`
         * 
         * @return builder
         * 
         */
        public Builder issuerName(@Nullable Output<String> issuerName) {
            $.issuerName = issuerName;
            return this;
        }

        /**
         * @param issuerName Provides a name to the specified issuer. The name must be unique
         * across all issuers and not be the reserved value `default`
         * 
         * @return builder
         * 
         */
        public Builder issuerName(String issuerName) {
            return issuerName(Output.of(issuerName));
        }

        /**
         * @param keyBits The number of bits to use
         * 
         * @return builder
         * 
         */
        public Builder keyBits(@Nullable Output<Integer> keyBits) {
            $.keyBits = keyBits;
            return this;
        }

        /**
         * @param keyBits The number of bits to use
         * 
         * @return builder
         * 
         */
        public Builder keyBits(Integer keyBits) {
            return keyBits(Output.of(keyBits));
        }

        /**
         * @param keyName When a new key is created with this request, optionally specifies
         * the name for this. The global ref `default` may not be used as a name.
         * 
         * @return builder
         * 
         */
        public Builder keyName(@Nullable Output<String> keyName) {
            $.keyName = keyName;
            return this;
        }

        /**
         * @param keyName When a new key is created with this request, optionally specifies
         * the name for this. The global ref `default` may not be used as a name.
         * 
         * @return builder
         * 
         */
        public Builder keyName(String keyName) {
            return keyName(Output.of(keyName));
        }

        /**
         * @param keyRef Specifies the key (either default, by name, or by identifier) to use
         * for generating this request. Only suitable for `type=existing` requests.
         * 
         * @return builder
         * 
         */
        public Builder keyRef(@Nullable Output<String> keyRef) {
            $.keyRef = keyRef;
            return this;
        }

        /**
         * @param keyRef Specifies the key (either default, by name, or by identifier) to use
         * for generating this request. Only suitable for `type=existing` requests.
         * 
         * @return builder
         * 
         */
        public Builder keyRef(String keyRef) {
            return keyRef(Output.of(keyRef));
        }

        /**
         * @param keyType The desired key type
         * 
         * @return builder
         * 
         */
        public Builder keyType(@Nullable Output<String> keyType) {
            $.keyType = keyType;
            return this;
        }

        /**
         * @param keyType The desired key type
         * 
         * @return builder
         * 
         */
        public Builder keyType(String keyType) {
            return keyType(Output.of(keyType));
        }

        /**
         * @param locality The locality
         * 
         * @return builder
         * 
         */
        public Builder locality(@Nullable Output<String> locality) {
            $.locality = locality;
            return this;
        }

        /**
         * @param locality The locality
         * 
         * @return builder
         * 
         */
        public Builder locality(String locality) {
            return locality(Output.of(locality));
        }

        /**
         * @param managedKeyId The ID of the previously configured managed key. This field is
         * required if `type` is `kms` and it conflicts with `managed_key_name`
         * 
         * @return builder
         * 
         */
        public Builder managedKeyId(@Nullable Output<String> managedKeyId) {
            $.managedKeyId = managedKeyId;
            return this;
        }

        /**
         * @param managedKeyId The ID of the previously configured managed key. This field is
         * required if `type` is `kms` and it conflicts with `managed_key_name`
         * 
         * @return builder
         * 
         */
        public Builder managedKeyId(String managedKeyId) {
            return managedKeyId(Output.of(managedKeyId));
        }

        /**
         * @param managedKeyName The name of the previously configured managed key. This field is
         * required if `type` is `kms`  and it conflicts with `managed_key_id`
         * 
         * @return builder
         * 
         */
        public Builder managedKeyName(@Nullable Output<String> managedKeyName) {
            $.managedKeyName = managedKeyName;
            return this;
        }

        /**
         * @param managedKeyName The name of the previously configured managed key. This field is
         * required if `type` is `kms`  and it conflicts with `managed_key_id`
         * 
         * @return builder
         * 
         */
        public Builder managedKeyName(String managedKeyName) {
            return managedKeyName(Output.of(managedKeyName));
        }

        /**
         * @param maxPathLength The maximum path length to encode in the generated certificate
         * 
         * @return builder
         * 
         */
        public Builder maxPathLength(@Nullable Output<Integer> maxPathLength) {
            $.maxPathLength = maxPathLength;
            return this;
        }

        /**
         * @param maxPathLength The maximum path length to encode in the generated certificate
         * 
         * @return builder
         * 
         */
        public Builder maxPathLength(Integer maxPathLength) {
            return maxPathLength(Output.of(maxPathLength));
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
         * @param organization The organization
         * 
         * @return builder
         * 
         */
        public Builder organization(@Nullable Output<String> organization) {
            $.organization = organization;
            return this;
        }

        /**
         * @param organization The organization
         * 
         * @return builder
         * 
         */
        public Builder organization(String organization) {
            return organization(Output.of(organization));
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
         * @param ou The organization unit
         * 
         * @return builder
         * 
         */
        public Builder ou(@Nullable Output<String> ou) {
            $.ou = ou;
            return this;
        }

        /**
         * @param ou The organization unit
         * 
         * @return builder
         * 
         */
        public Builder ou(String ou) {
            return ou(Output.of(ou));
        }

        /**
         * @param permittedDnsDomains List of domains for which certificates are allowed to be issued
         * 
         * @return builder
         * 
         */
        public Builder permittedDnsDomains(@Nullable Output<List<String>> permittedDnsDomains) {
            $.permittedDnsDomains = permittedDnsDomains;
            return this;
        }

        /**
         * @param permittedDnsDomains List of domains for which certificates are allowed to be issued
         * 
         * @return builder
         * 
         */
        public Builder permittedDnsDomains(List<String> permittedDnsDomains) {
            return permittedDnsDomains(Output.of(permittedDnsDomains));
        }

        /**
         * @param permittedDnsDomains List of domains for which certificates are allowed to be issued
         * 
         * @return builder
         * 
         */
        public Builder permittedDnsDomains(String... permittedDnsDomains) {
            return permittedDnsDomains(List.of(permittedDnsDomains));
        }

        /**
         * @param postalCode The postal code
         * 
         * @return builder
         * 
         */
        public Builder postalCode(@Nullable Output<String> postalCode) {
            $.postalCode = postalCode;
            return this;
        }

        /**
         * @param postalCode The postal code
         * 
         * @return builder
         * 
         */
        public Builder postalCode(String postalCode) {
            return postalCode(Output.of(postalCode));
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
         * @param province The province
         * 
         * @return builder
         * 
         */
        public Builder province(@Nullable Output<String> province) {
            $.province = province;
            return this;
        }

        /**
         * @param province The province
         * 
         * @return builder
         * 
         */
        public Builder province(String province) {
            return province(Output.of(province));
        }

        /**
         * @param streetAddress The street address
         * 
         * @return builder
         * 
         */
        public Builder streetAddress(@Nullable Output<String> streetAddress) {
            $.streetAddress = streetAddress;
            return this;
        }

        /**
         * @param streetAddress The street address
         * 
         * @return builder
         * 
         */
        public Builder streetAddress(String streetAddress) {
            return streetAddress(Output.of(streetAddress));
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
         * @param type Type of intermediate to create. Must be either \&#34;exported\&#34;, \&#34;internal\&#34;
         * or \&#34;kms\&#34;
         * 
         * @return builder
         * 
         */
        public Builder type(Output<String> type) {
            $.type = type;
            return this;
        }

        /**
         * @param type Type of intermediate to create. Must be either \&#34;exported\&#34;, \&#34;internal\&#34;
         * or \&#34;kms\&#34;
         * 
         * @return builder
         * 
         */
        public Builder type(String type) {
            return type(Output.of(type));
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

        public SecretBackendRootCertArgs build() {
            if ($.backend == null) {
                throw new MissingRequiredPropertyException("SecretBackendRootCertArgs", "backend");
            }
            if ($.commonName == null) {
                throw new MissingRequiredPropertyException("SecretBackendRootCertArgs", "commonName");
            }
            if ($.type == null) {
                throw new MissingRequiredPropertyException("SecretBackendRootCertArgs", "type");
            }
            return $;
        }
    }

}
