// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vault.terraformcloud;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class SecretCredsArgs extends com.pulumi.resources.ResourceArgs {

    public static final SecretCredsArgs Empty = new SecretCredsArgs();

    @Import(name="backend", required=true)
    private Output<String> backend;

    public Output<String> backend() {
        return this.backend;
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
     * Name of the role.
     * 
     */
    @Import(name="role", required=true)
    private Output<String> role;

    /**
     * @return Name of the role.
     * 
     */
    public Output<String> role() {
        return this.role;
    }

    private SecretCredsArgs() {}

    private SecretCredsArgs(SecretCredsArgs $) {
        this.backend = $.backend;
        this.namespace = $.namespace;
        this.role = $.role;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(SecretCredsArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private SecretCredsArgs $;

        public Builder() {
            $ = new SecretCredsArgs();
        }

        public Builder(SecretCredsArgs defaults) {
            $ = new SecretCredsArgs(Objects.requireNonNull(defaults));
        }

        public Builder backend(Output<String> backend) {
            $.backend = backend;
            return this;
        }

        public Builder backend(String backend) {
            return backend(Output.of(backend));
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
         * @param role Name of the role.
         * 
         * @return builder
         * 
         */
        public Builder role(Output<String> role) {
            $.role = role;
            return this;
        }

        /**
         * @param role Name of the role.
         * 
         * @return builder
         * 
         */
        public Builder role(String role) {
            return role(Output.of(role));
        }

        public SecretCredsArgs build() {
            if ($.backend == null) {
                throw new MissingRequiredPropertyException("SecretCredsArgs", "backend");
            }
            if ($.role == null) {
                throw new MissingRequiredPropertyException("SecretCredsArgs", "role");
            }
            return $;
        }
    }

}
