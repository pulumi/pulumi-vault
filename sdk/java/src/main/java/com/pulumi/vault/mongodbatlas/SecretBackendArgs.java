// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vault.mongodbatlas;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class SecretBackendArgs extends com.pulumi.resources.ResourceArgs {

    public static final SecretBackendArgs Empty = new SecretBackendArgs();

    /**
     * Path where the MongoDB Atlas Secrets Engine is mounted.
     * 
     */
    @Import(name="mount", required=true)
    private Output<String> mount;

    /**
     * @return Path where the MongoDB Atlas Secrets Engine is mounted.
     * 
     */
    public Output<String> mount() {
        return this.mount;
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
     * Specifies the Private API Key used to authenticate with the MongoDB Atlas API.
     * 
     */
    @Import(name="privateKey", required=true)
    private Output<String> privateKey;

    /**
     * @return Specifies the Private API Key used to authenticate with the MongoDB Atlas API.
     * 
     */
    public Output<String> privateKey() {
        return this.privateKey;
    }

    /**
     * Specifies the Public API Key used to authenticate with the MongoDB Atlas API.
     * 
     */
    @Import(name="publicKey", required=true)
    private Output<String> publicKey;

    /**
     * @return Specifies the Public API Key used to authenticate with the MongoDB Atlas API.
     * 
     */
    public Output<String> publicKey() {
        return this.publicKey;
    }

    private SecretBackendArgs() {}

    private SecretBackendArgs(SecretBackendArgs $) {
        this.mount = $.mount;
        this.namespace = $.namespace;
        this.privateKey = $.privateKey;
        this.publicKey = $.publicKey;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(SecretBackendArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private SecretBackendArgs $;

        public Builder() {
            $ = new SecretBackendArgs();
        }

        public Builder(SecretBackendArgs defaults) {
            $ = new SecretBackendArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param mount Path where the MongoDB Atlas Secrets Engine is mounted.
         * 
         * @return builder
         * 
         */
        public Builder mount(Output<String> mount) {
            $.mount = mount;
            return this;
        }

        /**
         * @param mount Path where the MongoDB Atlas Secrets Engine is mounted.
         * 
         * @return builder
         * 
         */
        public Builder mount(String mount) {
            return mount(Output.of(mount));
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
         * @param privateKey Specifies the Private API Key used to authenticate with the MongoDB Atlas API.
         * 
         * @return builder
         * 
         */
        public Builder privateKey(Output<String> privateKey) {
            $.privateKey = privateKey;
            return this;
        }

        /**
         * @param privateKey Specifies the Private API Key used to authenticate with the MongoDB Atlas API.
         * 
         * @return builder
         * 
         */
        public Builder privateKey(String privateKey) {
            return privateKey(Output.of(privateKey));
        }

        /**
         * @param publicKey Specifies the Public API Key used to authenticate with the MongoDB Atlas API.
         * 
         * @return builder
         * 
         */
        public Builder publicKey(Output<String> publicKey) {
            $.publicKey = publicKey;
            return this;
        }

        /**
         * @param publicKey Specifies the Public API Key used to authenticate with the MongoDB Atlas API.
         * 
         * @return builder
         * 
         */
        public Builder publicKey(String publicKey) {
            return publicKey(Output.of(publicKey));
        }

        public SecretBackendArgs build() {
            if ($.mount == null) {
                throw new MissingRequiredPropertyException("SecretBackendArgs", "mount");
            }
            if ($.privateKey == null) {
                throw new MissingRequiredPropertyException("SecretBackendArgs", "privateKey");
            }
            if ($.publicKey == null) {
                throw new MissingRequiredPropertyException("SecretBackendArgs", "publicKey");
            }
            return $;
        }
    }

}
