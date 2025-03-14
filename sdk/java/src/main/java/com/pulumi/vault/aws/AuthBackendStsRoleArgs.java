// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vault.aws;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class AuthBackendStsRoleArgs extends com.pulumi.resources.ResourceArgs {

    public static final AuthBackendStsRoleArgs Empty = new AuthBackendStsRoleArgs();

    /**
     * The AWS account ID to configure the STS role for.
     * 
     */
    @Import(name="accountId", required=true)
    private Output<String> accountId;

    /**
     * @return The AWS account ID to configure the STS role for.
     * 
     */
    public Output<String> accountId() {
        return this.accountId;
    }

    /**
     * The path the AWS auth backend being configured was
     * mounted at.  Defaults to `aws`.
     * 
     */
    @Import(name="backend")
    private @Nullable Output<String> backend;

    /**
     * @return The path the AWS auth backend being configured was
     * mounted at.  Defaults to `aws`.
     * 
     */
    public Optional<Output<String>> backend() {
        return Optional.ofNullable(this.backend);
    }

    /**
     * External ID expected by the STS role. The associated STS role must be configured to require the external ID. Requires Vault 1.17+.
     * 
     */
    @Import(name="externalId")
    private @Nullable Output<String> externalId;

    /**
     * @return External ID expected by the STS role. The associated STS role must be configured to require the external ID. Requires Vault 1.17+.
     * 
     */
    public Optional<Output<String>> externalId() {
        return Optional.ofNullable(this.externalId);
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
     * The STS role to assume when verifying requests made
     * by EC2 instances in the account specified by `account_id`.
     * 
     */
    @Import(name="stsRole", required=true)
    private Output<String> stsRole;

    /**
     * @return The STS role to assume when verifying requests made
     * by EC2 instances in the account specified by `account_id`.
     * 
     */
    public Output<String> stsRole() {
        return this.stsRole;
    }

    private AuthBackendStsRoleArgs() {}

    private AuthBackendStsRoleArgs(AuthBackendStsRoleArgs $) {
        this.accountId = $.accountId;
        this.backend = $.backend;
        this.externalId = $.externalId;
        this.namespace = $.namespace;
        this.stsRole = $.stsRole;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(AuthBackendStsRoleArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private AuthBackendStsRoleArgs $;

        public Builder() {
            $ = new AuthBackendStsRoleArgs();
        }

        public Builder(AuthBackendStsRoleArgs defaults) {
            $ = new AuthBackendStsRoleArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param accountId The AWS account ID to configure the STS role for.
         * 
         * @return builder
         * 
         */
        public Builder accountId(Output<String> accountId) {
            $.accountId = accountId;
            return this;
        }

        /**
         * @param accountId The AWS account ID to configure the STS role for.
         * 
         * @return builder
         * 
         */
        public Builder accountId(String accountId) {
            return accountId(Output.of(accountId));
        }

        /**
         * @param backend The path the AWS auth backend being configured was
         * mounted at.  Defaults to `aws`.
         * 
         * @return builder
         * 
         */
        public Builder backend(@Nullable Output<String> backend) {
            $.backend = backend;
            return this;
        }

        /**
         * @param backend The path the AWS auth backend being configured was
         * mounted at.  Defaults to `aws`.
         * 
         * @return builder
         * 
         */
        public Builder backend(String backend) {
            return backend(Output.of(backend));
        }

        /**
         * @param externalId External ID expected by the STS role. The associated STS role must be configured to require the external ID. Requires Vault 1.17+.
         * 
         * @return builder
         * 
         */
        public Builder externalId(@Nullable Output<String> externalId) {
            $.externalId = externalId;
            return this;
        }

        /**
         * @param externalId External ID expected by the STS role. The associated STS role must be configured to require the external ID. Requires Vault 1.17+.
         * 
         * @return builder
         * 
         */
        public Builder externalId(String externalId) {
            return externalId(Output.of(externalId));
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
         * @param stsRole The STS role to assume when verifying requests made
         * by EC2 instances in the account specified by `account_id`.
         * 
         * @return builder
         * 
         */
        public Builder stsRole(Output<String> stsRole) {
            $.stsRole = stsRole;
            return this;
        }

        /**
         * @param stsRole The STS role to assume when verifying requests made
         * by EC2 instances in the account specified by `account_id`.
         * 
         * @return builder
         * 
         */
        public Builder stsRole(String stsRole) {
            return stsRole(Output.of(stsRole));
        }

        public AuthBackendStsRoleArgs build() {
            if ($.accountId == null) {
                throw new MissingRequiredPropertyException("AuthBackendStsRoleArgs", "accountId");
            }
            if ($.stsRole == null) {
                throw new MissingRequiredPropertyException("AuthBackendStsRoleArgs", "stsRole");
            }
            return $;
        }
    }

}
