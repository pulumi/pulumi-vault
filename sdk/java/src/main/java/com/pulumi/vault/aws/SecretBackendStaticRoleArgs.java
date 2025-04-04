// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vault.aws;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class SecretBackendStaticRoleArgs extends com.pulumi.resources.ResourceArgs {

    public static final SecretBackendStaticRoleArgs Empty = new SecretBackendStaticRoleArgs();

    /**
     * Specifies the ARN of the role that Vault should assume.
     * When provided, Vault will use AWS STS to assume this role and generate temporary credentials.
     * If `assume_role_arn` is provided, `assume_role_session_name` must also be provided.
     * Requires Vault 1.19+. *Available only for Vault Enterprise*.
     * 
     */
    @Import(name="assumeRoleArn")
    private @Nullable Output<String> assumeRoleArn;

    /**
     * @return Specifies the ARN of the role that Vault should assume.
     * When provided, Vault will use AWS STS to assume this role and generate temporary credentials.
     * If `assume_role_arn` is provided, `assume_role_session_name` must also be provided.
     * Requires Vault 1.19+. *Available only for Vault Enterprise*.
     * 
     */
    public Optional<Output<String>> assumeRoleArn() {
        return Optional.ofNullable(this.assumeRoleArn);
    }

    /**
     * Specifies the session name to use when assuming the role.
     * If `assume_role_session_name` is provided, `assume_role_arn` must also be provided.
     * Requires Vault 1.19+. *Available only for Vault Enterprise*.
     * 
     */
    @Import(name="assumeRoleSessionName")
    private @Nullable Output<String> assumeRoleSessionName;

    /**
     * @return Specifies the session name to use when assuming the role.
     * If `assume_role_session_name` is provided, `assume_role_arn` must also be provided.
     * Requires Vault 1.19+. *Available only for Vault Enterprise*.
     * 
     */
    public Optional<Output<String>> assumeRoleSessionName() {
        return Optional.ofNullable(this.assumeRoleSessionName);
    }

    /**
     * The unique path this backend should be mounted at. Must
     * not begin or end with a `/`. Defaults to `aws`
     * 
     */
    @Import(name="backend")
    private @Nullable Output<String> backend;

    /**
     * @return The unique path this backend should be mounted at. Must
     * not begin or end with a `/`. Defaults to `aws`
     * 
     */
    public Optional<Output<String>> backend() {
        return Optional.ofNullable(this.backend);
    }

    /**
     * Specifies the external ID to use when assuming the role.
     * Requires Vault 1.19+. *Available only for Vault Enterprise*.
     * 
     */
    @Import(name="externalId")
    private @Nullable Output<String> externalId;

    /**
     * @return Specifies the external ID to use when assuming the role.
     * Requires Vault 1.19+. *Available only for Vault Enterprise*.
     * 
     */
    public Optional<Output<String>> externalId() {
        return Optional.ofNullable(this.externalId);
    }

    /**
     * The name to identify this role within the backend.
     * Must be unique within the backend.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return The name to identify this role within the backend.
     * Must be unique within the backend.
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
     * How often Vault should rotate the password of the user entry.
     * 
     */
    @Import(name="rotationPeriod", required=true)
    private Output<Integer> rotationPeriod;

    /**
     * @return How often Vault should rotate the password of the user entry.
     * 
     */
    public Output<Integer> rotationPeriod() {
        return this.rotationPeriod;
    }

    /**
     * The username of the existing AWS IAM to manage password rotation for.
     * 
     */
    @Import(name="username", required=true)
    private Output<String> username;

    /**
     * @return The username of the existing AWS IAM to manage password rotation for.
     * 
     */
    public Output<String> username() {
        return this.username;
    }

    private SecretBackendStaticRoleArgs() {}

    private SecretBackendStaticRoleArgs(SecretBackendStaticRoleArgs $) {
        this.assumeRoleArn = $.assumeRoleArn;
        this.assumeRoleSessionName = $.assumeRoleSessionName;
        this.backend = $.backend;
        this.externalId = $.externalId;
        this.name = $.name;
        this.namespace = $.namespace;
        this.rotationPeriod = $.rotationPeriod;
        this.username = $.username;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(SecretBackendStaticRoleArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private SecretBackendStaticRoleArgs $;

        public Builder() {
            $ = new SecretBackendStaticRoleArgs();
        }

        public Builder(SecretBackendStaticRoleArgs defaults) {
            $ = new SecretBackendStaticRoleArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param assumeRoleArn Specifies the ARN of the role that Vault should assume.
         * When provided, Vault will use AWS STS to assume this role and generate temporary credentials.
         * If `assume_role_arn` is provided, `assume_role_session_name` must also be provided.
         * Requires Vault 1.19+. *Available only for Vault Enterprise*.
         * 
         * @return builder
         * 
         */
        public Builder assumeRoleArn(@Nullable Output<String> assumeRoleArn) {
            $.assumeRoleArn = assumeRoleArn;
            return this;
        }

        /**
         * @param assumeRoleArn Specifies the ARN of the role that Vault should assume.
         * When provided, Vault will use AWS STS to assume this role and generate temporary credentials.
         * If `assume_role_arn` is provided, `assume_role_session_name` must also be provided.
         * Requires Vault 1.19+. *Available only for Vault Enterprise*.
         * 
         * @return builder
         * 
         */
        public Builder assumeRoleArn(String assumeRoleArn) {
            return assumeRoleArn(Output.of(assumeRoleArn));
        }

        /**
         * @param assumeRoleSessionName Specifies the session name to use when assuming the role.
         * If `assume_role_session_name` is provided, `assume_role_arn` must also be provided.
         * Requires Vault 1.19+. *Available only for Vault Enterprise*.
         * 
         * @return builder
         * 
         */
        public Builder assumeRoleSessionName(@Nullable Output<String> assumeRoleSessionName) {
            $.assumeRoleSessionName = assumeRoleSessionName;
            return this;
        }

        /**
         * @param assumeRoleSessionName Specifies the session name to use when assuming the role.
         * If `assume_role_session_name` is provided, `assume_role_arn` must also be provided.
         * Requires Vault 1.19+. *Available only for Vault Enterprise*.
         * 
         * @return builder
         * 
         */
        public Builder assumeRoleSessionName(String assumeRoleSessionName) {
            return assumeRoleSessionName(Output.of(assumeRoleSessionName));
        }

        /**
         * @param backend The unique path this backend should be mounted at. Must
         * not begin or end with a `/`. Defaults to `aws`
         * 
         * @return builder
         * 
         */
        public Builder backend(@Nullable Output<String> backend) {
            $.backend = backend;
            return this;
        }

        /**
         * @param backend The unique path this backend should be mounted at. Must
         * not begin or end with a `/`. Defaults to `aws`
         * 
         * @return builder
         * 
         */
        public Builder backend(String backend) {
            return backend(Output.of(backend));
        }

        /**
         * @param externalId Specifies the external ID to use when assuming the role.
         * Requires Vault 1.19+. *Available only for Vault Enterprise*.
         * 
         * @return builder
         * 
         */
        public Builder externalId(@Nullable Output<String> externalId) {
            $.externalId = externalId;
            return this;
        }

        /**
         * @param externalId Specifies the external ID to use when assuming the role.
         * Requires Vault 1.19+. *Available only for Vault Enterprise*.
         * 
         * @return builder
         * 
         */
        public Builder externalId(String externalId) {
            return externalId(Output.of(externalId));
        }

        /**
         * @param name The name to identify this role within the backend.
         * Must be unique within the backend.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name The name to identify this role within the backend.
         * Must be unique within the backend.
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
         * @param rotationPeriod How often Vault should rotate the password of the user entry.
         * 
         * @return builder
         * 
         */
        public Builder rotationPeriod(Output<Integer> rotationPeriod) {
            $.rotationPeriod = rotationPeriod;
            return this;
        }

        /**
         * @param rotationPeriod How often Vault should rotate the password of the user entry.
         * 
         * @return builder
         * 
         */
        public Builder rotationPeriod(Integer rotationPeriod) {
            return rotationPeriod(Output.of(rotationPeriod));
        }

        /**
         * @param username The username of the existing AWS IAM to manage password rotation for.
         * 
         * @return builder
         * 
         */
        public Builder username(Output<String> username) {
            $.username = username;
            return this;
        }

        /**
         * @param username The username of the existing AWS IAM to manage password rotation for.
         * 
         * @return builder
         * 
         */
        public Builder username(String username) {
            return username(Output.of(username));
        }

        public SecretBackendStaticRoleArgs build() {
            if ($.rotationPeriod == null) {
                throw new MissingRequiredPropertyException("SecretBackendStaticRoleArgs", "rotationPeriod");
            }
            if ($.username == null) {
                throw new MissingRequiredPropertyException("SecretBackendStaticRoleArgs", "username");
            }
            return $;
        }
    }

}
