// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vault.pkiSecret.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class SecretBackendKeyState extends com.pulumi.resources.ResourceArgs {

    public static final SecretBackendKeyState Empty = new SecretBackendKeyState();

    /**
     * The path the PKI secret backend is mounted at, with no leading or trailing `/`s.
     * 
     */
    @Import(name="backend")
    private @Nullable Output<String> backend;

    /**
     * @return The path the PKI secret backend is mounted at, with no leading or trailing `/`s.
     * 
     */
    public Optional<Output<String>> backend() {
        return Optional.ofNullable(this.backend);
    }

    /**
     * Specifies the number of bits to use for the generated keys.
     * Allowed values are 0 (universal default); with `key_type=rsa`, allowed values are:
     * 2048 (default), 3072, or 4096; with `key_type=ec`, allowed values are: 224, 256 (default),
     * 384, or 521; ignored with `key_type=ed25519`.
     * 
     */
    @Import(name="keyBits")
    private @Nullable Output<Integer> keyBits;

    /**
     * @return Specifies the number of bits to use for the generated keys.
     * Allowed values are 0 (universal default); with `key_type=rsa`, allowed values are:
     * 2048 (default), 3072, or 4096; with `key_type=ec`, allowed values are: 224, 256 (default),
     * 384, or 521; ignored with `key_type=ed25519`.
     * 
     */
    public Optional<Output<Integer>> keyBits() {
        return Optional.ofNullable(this.keyBits);
    }

    /**
     * ID of the generated key.
     * 
     */
    @Import(name="keyId")
    private @Nullable Output<String> keyId;

    /**
     * @return ID of the generated key.
     * 
     */
    public Optional<Output<String>> keyId() {
        return Optional.ofNullable(this.keyId);
    }

    /**
     * When a new key is created with this request, optionally specifies the name for this.
     * The global ref `default` may not be used as a name.
     * 
     */
    @Import(name="keyName")
    private @Nullable Output<String> keyName;

    /**
     * @return When a new key is created with this request, optionally specifies the name for this.
     * The global ref `default` may not be used as a name.
     * 
     */
    public Optional<Output<String>> keyName() {
        return Optional.ofNullable(this.keyName);
    }

    /**
     * Specifies the desired key type; must be `rsa`, `ed25519` or `ec`.
     * 
     */
    @Import(name="keyType")
    private @Nullable Output<String> keyType;

    /**
     * @return Specifies the desired key type; must be `rsa`, `ed25519` or `ec`.
     * 
     */
    public Optional<Output<String>> keyType() {
        return Optional.ofNullable(this.keyType);
    }

    /**
     * The managed key&#39;s UUID.
     * 
     */
    @Import(name="managedKeyId")
    private @Nullable Output<String> managedKeyId;

    /**
     * @return The managed key&#39;s UUID.
     * 
     */
    public Optional<Output<String>> managedKeyId() {
        return Optional.ofNullable(this.managedKeyId);
    }

    /**
     * The managed key&#39;s configured name.
     * 
     */
    @Import(name="managedKeyName")
    private @Nullable Output<String> managedKeyName;

    /**
     * @return The managed key&#39;s configured name.
     * 
     */
    public Optional<Output<String>> managedKeyName() {
        return Optional.ofNullable(this.managedKeyName);
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
     * Specifies the type of the key to create. Can be `exported`,`internal` or `kms`.
     * 
     */
    @Import(name="type")
    private @Nullable Output<String> type;

    /**
     * @return Specifies the type of the key to create. Can be `exported`,`internal` or `kms`.
     * 
     */
    public Optional<Output<String>> type() {
        return Optional.ofNullable(this.type);
    }

    private SecretBackendKeyState() {}

    private SecretBackendKeyState(SecretBackendKeyState $) {
        this.backend = $.backend;
        this.keyBits = $.keyBits;
        this.keyId = $.keyId;
        this.keyName = $.keyName;
        this.keyType = $.keyType;
        this.managedKeyId = $.managedKeyId;
        this.managedKeyName = $.managedKeyName;
        this.namespace = $.namespace;
        this.type = $.type;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(SecretBackendKeyState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private SecretBackendKeyState $;

        public Builder() {
            $ = new SecretBackendKeyState();
        }

        public Builder(SecretBackendKeyState defaults) {
            $ = new SecretBackendKeyState(Objects.requireNonNull(defaults));
        }

        /**
         * @param backend The path the PKI secret backend is mounted at, with no leading or trailing `/`s.
         * 
         * @return builder
         * 
         */
        public Builder backend(@Nullable Output<String> backend) {
            $.backend = backend;
            return this;
        }

        /**
         * @param backend The path the PKI secret backend is mounted at, with no leading or trailing `/`s.
         * 
         * @return builder
         * 
         */
        public Builder backend(String backend) {
            return backend(Output.of(backend));
        }

        /**
         * @param keyBits Specifies the number of bits to use for the generated keys.
         * Allowed values are 0 (universal default); with `key_type=rsa`, allowed values are:
         * 2048 (default), 3072, or 4096; with `key_type=ec`, allowed values are: 224, 256 (default),
         * 384, or 521; ignored with `key_type=ed25519`.
         * 
         * @return builder
         * 
         */
        public Builder keyBits(@Nullable Output<Integer> keyBits) {
            $.keyBits = keyBits;
            return this;
        }

        /**
         * @param keyBits Specifies the number of bits to use for the generated keys.
         * Allowed values are 0 (universal default); with `key_type=rsa`, allowed values are:
         * 2048 (default), 3072, or 4096; with `key_type=ec`, allowed values are: 224, 256 (default),
         * 384, or 521; ignored with `key_type=ed25519`.
         * 
         * @return builder
         * 
         */
        public Builder keyBits(Integer keyBits) {
            return keyBits(Output.of(keyBits));
        }

        /**
         * @param keyId ID of the generated key.
         * 
         * @return builder
         * 
         */
        public Builder keyId(@Nullable Output<String> keyId) {
            $.keyId = keyId;
            return this;
        }

        /**
         * @param keyId ID of the generated key.
         * 
         * @return builder
         * 
         */
        public Builder keyId(String keyId) {
            return keyId(Output.of(keyId));
        }

        /**
         * @param keyName When a new key is created with this request, optionally specifies the name for this.
         * The global ref `default` may not be used as a name.
         * 
         * @return builder
         * 
         */
        public Builder keyName(@Nullable Output<String> keyName) {
            $.keyName = keyName;
            return this;
        }

        /**
         * @param keyName When a new key is created with this request, optionally specifies the name for this.
         * The global ref `default` may not be used as a name.
         * 
         * @return builder
         * 
         */
        public Builder keyName(String keyName) {
            return keyName(Output.of(keyName));
        }

        /**
         * @param keyType Specifies the desired key type; must be `rsa`, `ed25519` or `ec`.
         * 
         * @return builder
         * 
         */
        public Builder keyType(@Nullable Output<String> keyType) {
            $.keyType = keyType;
            return this;
        }

        /**
         * @param keyType Specifies the desired key type; must be `rsa`, `ed25519` or `ec`.
         * 
         * @return builder
         * 
         */
        public Builder keyType(String keyType) {
            return keyType(Output.of(keyType));
        }

        /**
         * @param managedKeyId The managed key&#39;s UUID.
         * 
         * @return builder
         * 
         */
        public Builder managedKeyId(@Nullable Output<String> managedKeyId) {
            $.managedKeyId = managedKeyId;
            return this;
        }

        /**
         * @param managedKeyId The managed key&#39;s UUID.
         * 
         * @return builder
         * 
         */
        public Builder managedKeyId(String managedKeyId) {
            return managedKeyId(Output.of(managedKeyId));
        }

        /**
         * @param managedKeyName The managed key&#39;s configured name.
         * 
         * @return builder
         * 
         */
        public Builder managedKeyName(@Nullable Output<String> managedKeyName) {
            $.managedKeyName = managedKeyName;
            return this;
        }

        /**
         * @param managedKeyName The managed key&#39;s configured name.
         * 
         * @return builder
         * 
         */
        public Builder managedKeyName(String managedKeyName) {
            return managedKeyName(Output.of(managedKeyName));
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
         * @param type Specifies the type of the key to create. Can be `exported`,`internal` or `kms`.
         * 
         * @return builder
         * 
         */
        public Builder type(@Nullable Output<String> type) {
            $.type = type;
            return this;
        }

        /**
         * @param type Specifies the type of the key to create. Can be `exported`,`internal` or `kms`.
         * 
         * @return builder
         * 
         */
        public Builder type(String type) {
            return type(Output.of(type));
        }

        public SecretBackendKeyState build() {
            return $;
        }
    }

}