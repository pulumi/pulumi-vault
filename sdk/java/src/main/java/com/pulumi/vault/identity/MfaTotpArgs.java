// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vault.identity;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class MfaTotpArgs extends com.pulumi.resources.ResourceArgs {

    public static final MfaTotpArgs Empty = new MfaTotpArgs();

    /**
     * Specifies the hashing algorithm used to generate the TOTP code. Options include SHA1, SHA256, SHA512.
     * 
     */
    @Import(name="algorithm")
    private @Nullable Output<String> algorithm;

    /**
     * @return Specifies the hashing algorithm used to generate the TOTP code. Options include SHA1, SHA256, SHA512.
     * 
     */
    public Optional<Output<String>> algorithm() {
        return Optional.ofNullable(this.algorithm);
    }

    /**
     * The number of digits in the generated TOTP token. This value can either be 6 or 8
     * 
     */
    @Import(name="digits")
    private @Nullable Output<Integer> digits;

    /**
     * @return The number of digits in the generated TOTP token. This value can either be 6 or 8
     * 
     */
    public Optional<Output<Integer>> digits() {
        return Optional.ofNullable(this.digits);
    }

    /**
     * The name of the key&#39;s issuing organization.
     * 
     */
    @Import(name="issuer", required=true)
    private Output<String> issuer;

    /**
     * @return The name of the key&#39;s issuing organization.
     * 
     */
    public Output<String> issuer() {
        return this.issuer;
    }

    /**
     * Specifies the size in bytes of the generated key.
     * 
     */
    @Import(name="keySize")
    private @Nullable Output<Integer> keySize;

    /**
     * @return Specifies the size in bytes of the generated key.
     * 
     */
    public Optional<Output<Integer>> keySize() {
        return Optional.ofNullable(this.keySize);
    }

    /**
     * The maximum number of consecutive failed validation attempts allowed.
     * 
     */
    @Import(name="maxValidationAttempts")
    private @Nullable Output<Integer> maxValidationAttempts;

    /**
     * @return The maximum number of consecutive failed validation attempts allowed.
     * 
     */
    public Optional<Output<Integer>> maxValidationAttempts() {
        return Optional.ofNullable(this.maxValidationAttempts);
    }

    /**
     * Target namespace. (requires Enterprise)
     * 
     */
    @Import(name="namespace")
    private @Nullable Output<String> namespace;

    /**
     * @return Target namespace. (requires Enterprise)
     * 
     */
    public Optional<Output<String>> namespace() {
        return Optional.ofNullable(this.namespace);
    }

    /**
     * The length of time in seconds used to generate a counter for the TOTP token calculation.
     * 
     */
    @Import(name="period")
    private @Nullable Output<Integer> period;

    /**
     * @return The length of time in seconds used to generate a counter for the TOTP token calculation.
     * 
     */
    public Optional<Output<Integer>> period() {
        return Optional.ofNullable(this.period);
    }

    /**
     * The pixel size of the generated square QR code.
     * 
     */
    @Import(name="qrSize")
    private @Nullable Output<Integer> qrSize;

    /**
     * @return The pixel size of the generated square QR code.
     * 
     */
    public Optional<Output<Integer>> qrSize() {
        return Optional.ofNullable(this.qrSize);
    }

    /**
     * The number of delay periods that are allowed when validating a TOTP token. This value can either be 0 or 1.
     * 
     */
    @Import(name="skew")
    private @Nullable Output<Integer> skew;

    /**
     * @return The number of delay periods that are allowed when validating a TOTP token. This value can either be 0 or 1.
     * 
     */
    public Optional<Output<Integer>> skew() {
        return Optional.ofNullable(this.skew);
    }

    private MfaTotpArgs() {}

    private MfaTotpArgs(MfaTotpArgs $) {
        this.algorithm = $.algorithm;
        this.digits = $.digits;
        this.issuer = $.issuer;
        this.keySize = $.keySize;
        this.maxValidationAttempts = $.maxValidationAttempts;
        this.namespace = $.namespace;
        this.period = $.period;
        this.qrSize = $.qrSize;
        this.skew = $.skew;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(MfaTotpArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private MfaTotpArgs $;

        public Builder() {
            $ = new MfaTotpArgs();
        }

        public Builder(MfaTotpArgs defaults) {
            $ = new MfaTotpArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param algorithm Specifies the hashing algorithm used to generate the TOTP code. Options include SHA1, SHA256, SHA512.
         * 
         * @return builder
         * 
         */
        public Builder algorithm(@Nullable Output<String> algorithm) {
            $.algorithm = algorithm;
            return this;
        }

        /**
         * @param algorithm Specifies the hashing algorithm used to generate the TOTP code. Options include SHA1, SHA256, SHA512.
         * 
         * @return builder
         * 
         */
        public Builder algorithm(String algorithm) {
            return algorithm(Output.of(algorithm));
        }

        /**
         * @param digits The number of digits in the generated TOTP token. This value can either be 6 or 8
         * 
         * @return builder
         * 
         */
        public Builder digits(@Nullable Output<Integer> digits) {
            $.digits = digits;
            return this;
        }

        /**
         * @param digits The number of digits in the generated TOTP token. This value can either be 6 or 8
         * 
         * @return builder
         * 
         */
        public Builder digits(Integer digits) {
            return digits(Output.of(digits));
        }

        /**
         * @param issuer The name of the key&#39;s issuing organization.
         * 
         * @return builder
         * 
         */
        public Builder issuer(Output<String> issuer) {
            $.issuer = issuer;
            return this;
        }

        /**
         * @param issuer The name of the key&#39;s issuing organization.
         * 
         * @return builder
         * 
         */
        public Builder issuer(String issuer) {
            return issuer(Output.of(issuer));
        }

        /**
         * @param keySize Specifies the size in bytes of the generated key.
         * 
         * @return builder
         * 
         */
        public Builder keySize(@Nullable Output<Integer> keySize) {
            $.keySize = keySize;
            return this;
        }

        /**
         * @param keySize Specifies the size in bytes of the generated key.
         * 
         * @return builder
         * 
         */
        public Builder keySize(Integer keySize) {
            return keySize(Output.of(keySize));
        }

        /**
         * @param maxValidationAttempts The maximum number of consecutive failed validation attempts allowed.
         * 
         * @return builder
         * 
         */
        public Builder maxValidationAttempts(@Nullable Output<Integer> maxValidationAttempts) {
            $.maxValidationAttempts = maxValidationAttempts;
            return this;
        }

        /**
         * @param maxValidationAttempts The maximum number of consecutive failed validation attempts allowed.
         * 
         * @return builder
         * 
         */
        public Builder maxValidationAttempts(Integer maxValidationAttempts) {
            return maxValidationAttempts(Output.of(maxValidationAttempts));
        }

        /**
         * @param namespace Target namespace. (requires Enterprise)
         * 
         * @return builder
         * 
         */
        public Builder namespace(@Nullable Output<String> namespace) {
            $.namespace = namespace;
            return this;
        }

        /**
         * @param namespace Target namespace. (requires Enterprise)
         * 
         * @return builder
         * 
         */
        public Builder namespace(String namespace) {
            return namespace(Output.of(namespace));
        }

        /**
         * @param period The length of time in seconds used to generate a counter for the TOTP token calculation.
         * 
         * @return builder
         * 
         */
        public Builder period(@Nullable Output<Integer> period) {
            $.period = period;
            return this;
        }

        /**
         * @param period The length of time in seconds used to generate a counter for the TOTP token calculation.
         * 
         * @return builder
         * 
         */
        public Builder period(Integer period) {
            return period(Output.of(period));
        }

        /**
         * @param qrSize The pixel size of the generated square QR code.
         * 
         * @return builder
         * 
         */
        public Builder qrSize(@Nullable Output<Integer> qrSize) {
            $.qrSize = qrSize;
            return this;
        }

        /**
         * @param qrSize The pixel size of the generated square QR code.
         * 
         * @return builder
         * 
         */
        public Builder qrSize(Integer qrSize) {
            return qrSize(Output.of(qrSize));
        }

        /**
         * @param skew The number of delay periods that are allowed when validating a TOTP token. This value can either be 0 or 1.
         * 
         * @return builder
         * 
         */
        public Builder skew(@Nullable Output<Integer> skew) {
            $.skew = skew;
            return this;
        }

        /**
         * @param skew The number of delay periods that are allowed when validating a TOTP token. This value can either be 0 or 1.
         * 
         * @return builder
         * 
         */
        public Builder skew(Integer skew) {
            return skew(Output.of(skew));
        }

        public MfaTotpArgs build() {
            $.issuer = Objects.requireNonNull($.issuer, "expected parameter 'issuer' to be non-null");
            return $;
        }
    }

}
