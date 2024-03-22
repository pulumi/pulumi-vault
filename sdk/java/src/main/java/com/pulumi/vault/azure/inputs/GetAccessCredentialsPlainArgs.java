// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vault.azure.inputs;

import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetAccessCredentialsPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetAccessCredentialsPlainArgs Empty = new GetAccessCredentialsPlainArgs();

    /**
     * The path to the Azure secret backend to
     * read credentials from, with no leading or trailing `/`s.
     * 
     */
    @Import(name="backend", required=true)
    private String backend;

    /**
     * @return The path to the Azure secret backend to
     * read credentials from, with no leading or trailing `/`s.
     * 
     */
    public String backend() {
        return this.backend;
    }

    /**
     * The Azure environment to use during credential validation.
     * Defaults to the environment configured in the Vault backend.
     * Some possible values: `AzurePublicCloud`, `AzureGovernmentCloud`
     * *See the caveats section for more information on this field.*
     * 
     */
    @Import(name="environment")
    private @Nullable String environment;

    /**
     * @return The Azure environment to use during credential validation.
     * Defaults to the environment configured in the Vault backend.
     * Some possible values: `AzurePublicCloud`, `AzureGovernmentCloud`
     * *See the caveats section for more information on this field.*
     * 
     */
    public Optional<String> environment() {
        return Optional.ofNullable(this.environment);
    }

    /**
     * If &#39;validate_creds&#39; is true,
     * the number of seconds after which to give up validating credentials. Defaults
     * to 300.
     * 
     */
    @Import(name="maxCredValidationSeconds")
    private @Nullable Integer maxCredValidationSeconds;

    /**
     * @return If &#39;validate_creds&#39; is true,
     * the number of seconds after which to give up validating credentials. Defaults
     * to 300.
     * 
     */
    public Optional<Integer> maxCredValidationSeconds() {
        return Optional.ofNullable(this.maxCredValidationSeconds);
    }

    /**
     * The namespace of the target resource.
     * The value should not contain leading or trailing forward slashes.
     * The `namespace` is always relative to the provider&#39;s configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
     * *Available only for Vault Enterprise*.
     * 
     */
    @Import(name="namespace")
    private @Nullable String namespace;

    /**
     * @return The namespace of the target resource.
     * The value should not contain leading or trailing forward slashes.
     * The `namespace` is always relative to the provider&#39;s configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
     * *Available only for Vault Enterprise*.
     * 
     */
    public Optional<String> namespace() {
        return Optional.ofNullable(this.namespace);
    }

    /**
     * If &#39;validate_creds&#39; is true,
     * the number of seconds to wait between each test of generated credentials.
     * Defaults to 1.
     * 
     */
    @Import(name="numSecondsBetweenTests")
    private @Nullable Integer numSecondsBetweenTests;

    /**
     * @return If &#39;validate_creds&#39; is true,
     * the number of seconds to wait between each test of generated credentials.
     * Defaults to 1.
     * 
     */
    public Optional<Integer> numSecondsBetweenTests() {
        return Optional.ofNullable(this.numSecondsBetweenTests);
    }

    /**
     * If &#39;validate_creds&#39; is true,
     * the number of sequential successes required to validate generated
     * credentials. Defaults to 8.
     * 
     */
    @Import(name="numSequentialSuccesses")
    private @Nullable Integer numSequentialSuccesses;

    /**
     * @return If &#39;validate_creds&#39; is true,
     * the number of sequential successes required to validate generated
     * credentials. Defaults to 8.
     * 
     */
    public Optional<Integer> numSequentialSuccesses() {
        return Optional.ofNullable(this.numSequentialSuccesses);
    }

    /**
     * The name of the Azure secret backend role to read
     * credentials from, with no leading or trailing `/`s.
     * 
     */
    @Import(name="role", required=true)
    private String role;

    /**
     * @return The name of the Azure secret backend role to read
     * credentials from, with no leading or trailing `/`s.
     * 
     */
    public String role() {
        return this.role;
    }

    /**
     * The subscription ID to use during credential
     * validation. Defaults to the subscription ID configured in the Vault `backend`.
     * *See the caveats section for more information on this field.*
     * 
     */
    @Import(name="subscriptionId")
    private @Nullable String subscriptionId;

    /**
     * @return The subscription ID to use during credential
     * validation. Defaults to the subscription ID configured in the Vault `backend`.
     * *See the caveats section for more information on this field.*
     * 
     */
    public Optional<String> subscriptionId() {
        return Optional.ofNullable(this.subscriptionId);
    }

    /**
     * The tenant ID to use during credential validation.
     * Defaults to the tenant ID configured in the Vault `backend`.
     * *See the caveats section for more information on this field.*
     * 
     */
    @Import(name="tenantId")
    private @Nullable String tenantId;

    /**
     * @return The tenant ID to use during credential validation.
     * Defaults to the tenant ID configured in the Vault `backend`.
     * *See the caveats section for more information on this field.*
     * 
     */
    public Optional<String> tenantId() {
        return Optional.ofNullable(this.tenantId);
    }

    /**
     * Whether generated credentials should be
     * validated before being returned. Defaults to `false`, which returns
     * credentials without checking whether they have fully propagated throughout
     * Azure Active Directory. Designating `true` activates testing.
     * 
     */
    @Import(name="validateCreds")
    private @Nullable Boolean validateCreds;

    /**
     * @return Whether generated credentials should be
     * validated before being returned. Defaults to `false`, which returns
     * credentials without checking whether they have fully propagated throughout
     * Azure Active Directory. Designating `true` activates testing.
     * 
     */
    public Optional<Boolean> validateCreds() {
        return Optional.ofNullable(this.validateCreds);
    }

    private GetAccessCredentialsPlainArgs() {}

    private GetAccessCredentialsPlainArgs(GetAccessCredentialsPlainArgs $) {
        this.backend = $.backend;
        this.environment = $.environment;
        this.maxCredValidationSeconds = $.maxCredValidationSeconds;
        this.namespace = $.namespace;
        this.numSecondsBetweenTests = $.numSecondsBetweenTests;
        this.numSequentialSuccesses = $.numSequentialSuccesses;
        this.role = $.role;
        this.subscriptionId = $.subscriptionId;
        this.tenantId = $.tenantId;
        this.validateCreds = $.validateCreds;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetAccessCredentialsPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetAccessCredentialsPlainArgs $;

        public Builder() {
            $ = new GetAccessCredentialsPlainArgs();
        }

        public Builder(GetAccessCredentialsPlainArgs defaults) {
            $ = new GetAccessCredentialsPlainArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param backend The path to the Azure secret backend to
         * read credentials from, with no leading or trailing `/`s.
         * 
         * @return builder
         * 
         */
        public Builder backend(String backend) {
            $.backend = backend;
            return this;
        }

        /**
         * @param environment The Azure environment to use during credential validation.
         * Defaults to the environment configured in the Vault backend.
         * Some possible values: `AzurePublicCloud`, `AzureGovernmentCloud`
         * *See the caveats section for more information on this field.*
         * 
         * @return builder
         * 
         */
        public Builder environment(@Nullable String environment) {
            $.environment = environment;
            return this;
        }

        /**
         * @param maxCredValidationSeconds If &#39;validate_creds&#39; is true,
         * the number of seconds after which to give up validating credentials. Defaults
         * to 300.
         * 
         * @return builder
         * 
         */
        public Builder maxCredValidationSeconds(@Nullable Integer maxCredValidationSeconds) {
            $.maxCredValidationSeconds = maxCredValidationSeconds;
            return this;
        }

        /**
         * @param namespace The namespace of the target resource.
         * The value should not contain leading or trailing forward slashes.
         * The `namespace` is always relative to the provider&#39;s configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
         * *Available only for Vault Enterprise*.
         * 
         * @return builder
         * 
         */
        public Builder namespace(@Nullable String namespace) {
            $.namespace = namespace;
            return this;
        }

        /**
         * @param numSecondsBetweenTests If &#39;validate_creds&#39; is true,
         * the number of seconds to wait between each test of generated credentials.
         * Defaults to 1.
         * 
         * @return builder
         * 
         */
        public Builder numSecondsBetweenTests(@Nullable Integer numSecondsBetweenTests) {
            $.numSecondsBetweenTests = numSecondsBetweenTests;
            return this;
        }

        /**
         * @param numSequentialSuccesses If &#39;validate_creds&#39; is true,
         * the number of sequential successes required to validate generated
         * credentials. Defaults to 8.
         * 
         * @return builder
         * 
         */
        public Builder numSequentialSuccesses(@Nullable Integer numSequentialSuccesses) {
            $.numSequentialSuccesses = numSequentialSuccesses;
            return this;
        }

        /**
         * @param role The name of the Azure secret backend role to read
         * credentials from, with no leading or trailing `/`s.
         * 
         * @return builder
         * 
         */
        public Builder role(String role) {
            $.role = role;
            return this;
        }

        /**
         * @param subscriptionId The subscription ID to use during credential
         * validation. Defaults to the subscription ID configured in the Vault `backend`.
         * *See the caveats section for more information on this field.*
         * 
         * @return builder
         * 
         */
        public Builder subscriptionId(@Nullable String subscriptionId) {
            $.subscriptionId = subscriptionId;
            return this;
        }

        /**
         * @param tenantId The tenant ID to use during credential validation.
         * Defaults to the tenant ID configured in the Vault `backend`.
         * *See the caveats section for more information on this field.*
         * 
         * @return builder
         * 
         */
        public Builder tenantId(@Nullable String tenantId) {
            $.tenantId = tenantId;
            return this;
        }

        /**
         * @param validateCreds Whether generated credentials should be
         * validated before being returned. Defaults to `false`, which returns
         * credentials without checking whether they have fully propagated throughout
         * Azure Active Directory. Designating `true` activates testing.
         * 
         * @return builder
         * 
         */
        public Builder validateCreds(@Nullable Boolean validateCreds) {
            $.validateCreds = validateCreds;
            return this;
        }

        public GetAccessCredentialsPlainArgs build() {
            if ($.backend == null) {
                throw new MissingRequiredPropertyException("GetAccessCredentialsPlainArgs", "backend");
            }
            if ($.role == null) {
                throw new MissingRequiredPropertyException("GetAccessCredentialsPlainArgs", "role");
            }
            return $;
        }
    }

}
