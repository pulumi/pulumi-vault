// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vault.ssh.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class SecretBackendCaState extends com.pulumi.resources.ResourceArgs {

    public static final SecretBackendCaState Empty = new SecretBackendCaState();

    /**
     * The path where the SSH secret backend is mounted. Defaults to &#39;ssh&#39;
     * 
     */
    @Import(name="backend")
    private @Nullable Output<String> backend;

    /**
     * @return The path where the SSH secret backend is mounted. Defaults to &#39;ssh&#39;
     * 
     */
    public Optional<Output<String>> backend() {
        return Optional.ofNullable(this.backend);
    }

    /**
     * Whether Vault should generate the signing key pair internally. Defaults to true
     * 
     */
    @Import(name="generateSigningKey")
    private @Nullable Output<Boolean> generateSigningKey;

    /**
     * @return Whether Vault should generate the signing key pair internally. Defaults to true
     * 
     */
    public Optional<Output<Boolean>> generateSigningKey() {
        return Optional.ofNullable(this.generateSigningKey);
    }

    /**
     * Specifies the desired key bits for the generated SSH CA key when `generate_signing_key` is set to `true`.
     * 
     */
    @Import(name="keyBits")
    private @Nullable Output<Integer> keyBits;

    /**
     * @return Specifies the desired key bits for the generated SSH CA key when `generate_signing_key` is set to `true`.
     * 
     */
    public Optional<Output<Integer>> keyBits() {
        return Optional.ofNullable(this.keyBits);
    }

    /**
     * Specifies the desired key type for the generated SSH CA key when `generate_signing_key` is set to `true`.
     * 
     */
    @Import(name="keyType")
    private @Nullable Output<String> keyType;

    /**
     * @return Specifies the desired key type for the generated SSH CA key when `generate_signing_key` is set to `true`.
     * 
     */
    public Optional<Output<String>> keyType() {
        return Optional.ofNullable(this.keyType);
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
     * Private key part the SSH CA key pair; required if generate_signing_key is false.
     * 
     */
    @Import(name="privateKey")
    private @Nullable Output<String> privateKey;

    /**
     * @return Private key part the SSH CA key pair; required if generate_signing_key is false.
     * 
     */
    public Optional<Output<String>> privateKey() {
        return Optional.ofNullable(this.privateKey);
    }

    /**
     * The public key part the SSH CA key pair; required if generate_signing_key is false.
     * 
     */
    @Import(name="publicKey")
    private @Nullable Output<String> publicKey;

    /**
     * @return The public key part the SSH CA key pair; required if generate_signing_key is false.
     * 
     */
    public Optional<Output<String>> publicKey() {
        return Optional.ofNullable(this.publicKey);
    }

    private SecretBackendCaState() {}

    private SecretBackendCaState(SecretBackendCaState $) {
        this.backend = $.backend;
        this.generateSigningKey = $.generateSigningKey;
        this.keyBits = $.keyBits;
        this.keyType = $.keyType;
        this.namespace = $.namespace;
        this.privateKey = $.privateKey;
        this.publicKey = $.publicKey;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(SecretBackendCaState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private SecretBackendCaState $;

        public Builder() {
            $ = new SecretBackendCaState();
        }

        public Builder(SecretBackendCaState defaults) {
            $ = new SecretBackendCaState(Objects.requireNonNull(defaults));
        }

        /**
         * @param backend The path where the SSH secret backend is mounted. Defaults to &#39;ssh&#39;
         * 
         * @return builder
         * 
         */
        public Builder backend(@Nullable Output<String> backend) {
            $.backend = backend;
            return this;
        }

        /**
         * @param backend The path where the SSH secret backend is mounted. Defaults to &#39;ssh&#39;
         * 
         * @return builder
         * 
         */
        public Builder backend(String backend) {
            return backend(Output.of(backend));
        }

        /**
         * @param generateSigningKey Whether Vault should generate the signing key pair internally. Defaults to true
         * 
         * @return builder
         * 
         */
        public Builder generateSigningKey(@Nullable Output<Boolean> generateSigningKey) {
            $.generateSigningKey = generateSigningKey;
            return this;
        }

        /**
         * @param generateSigningKey Whether Vault should generate the signing key pair internally. Defaults to true
         * 
         * @return builder
         * 
         */
        public Builder generateSigningKey(Boolean generateSigningKey) {
            return generateSigningKey(Output.of(generateSigningKey));
        }

        /**
         * @param keyBits Specifies the desired key bits for the generated SSH CA key when `generate_signing_key` is set to `true`.
         * 
         * @return builder
         * 
         */
        public Builder keyBits(@Nullable Output<Integer> keyBits) {
            $.keyBits = keyBits;
            return this;
        }

        /**
         * @param keyBits Specifies the desired key bits for the generated SSH CA key when `generate_signing_key` is set to `true`.
         * 
         * @return builder
         * 
         */
        public Builder keyBits(Integer keyBits) {
            return keyBits(Output.of(keyBits));
        }

        /**
         * @param keyType Specifies the desired key type for the generated SSH CA key when `generate_signing_key` is set to `true`.
         * 
         * @return builder
         * 
         */
        public Builder keyType(@Nullable Output<String> keyType) {
            $.keyType = keyType;
            return this;
        }

        /**
         * @param keyType Specifies the desired key type for the generated SSH CA key when `generate_signing_key` is set to `true`.
         * 
         * @return builder
         * 
         */
        public Builder keyType(String keyType) {
            return keyType(Output.of(keyType));
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
         * @param privateKey Private key part the SSH CA key pair; required if generate_signing_key is false.
         * 
         * @return builder
         * 
         */
        public Builder privateKey(@Nullable Output<String> privateKey) {
            $.privateKey = privateKey;
            return this;
        }

        /**
         * @param privateKey Private key part the SSH CA key pair; required if generate_signing_key is false.
         * 
         * @return builder
         * 
         */
        public Builder privateKey(String privateKey) {
            return privateKey(Output.of(privateKey));
        }

        /**
         * @param publicKey The public key part the SSH CA key pair; required if generate_signing_key is false.
         * 
         * @return builder
         * 
         */
        public Builder publicKey(@Nullable Output<String> publicKey) {
            $.publicKey = publicKey;
            return this;
        }

        /**
         * @param publicKey The public key part the SSH CA key pair; required if generate_signing_key is false.
         * 
         * @return builder
         * 
         */
        public Builder publicKey(String publicKey) {
            return publicKey(Output.of(publicKey));
        }

        public SecretBackendCaState build() {
            return $;
        }
    }

}
