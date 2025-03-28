// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vault.pkiSecret.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetBackendConfigCmpv2Args extends com.pulumi.resources.InvokeArgs {

    public static final GetBackendConfigCmpv2Args Empty = new GetBackendConfigCmpv2Args();

    /**
     * The path to the PKI secret backend to
     * read the CMPv2 configuration from, with no leading or trailing `/`s.
     * 
     * # Attributes Reference
     * 
     */
    @Import(name="backend", required=true)
    private Output<String> backend;

    /**
     * @return The path to the PKI secret backend to
     * read the CMPv2 configuration from, with no leading or trailing `/`s.
     * 
     * # Attributes Reference
     * 
     */
    public Output<String> backend() {
        return this.backend;
    }

    /**
     * A comma-separated list of validations not to perform on CMPv2 messages.
     * 
     */
    @Import(name="disabledValidations")
    private @Nullable Output<List<String>> disabledValidations;

    /**
     * @return A comma-separated list of validations not to perform on CMPv2 messages.
     * 
     */
    public Optional<Output<List<String>>> disabledValidations() {
        return Optional.ofNullable(this.disabledValidations);
    }

    /**
     * The namespace of the target resource.
     * The value should not contain leading or trailing forward slashes.
     * The `namespace` is always relative to the provider&#39;s configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
     * *Available only for Vault Enterprise*.
     * 
     */
    @Import(name="namespace")
    private @Nullable Output<String> namespace;

    /**
     * @return The namespace of the target resource.
     * The value should not contain leading or trailing forward slashes.
     * The `namespace` is always relative to the provider&#39;s configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
     * *Available only for Vault Enterprise*.
     * 
     */
    public Optional<Output<String>> namespace() {
        return Optional.ofNullable(this.namespace);
    }

    private GetBackendConfigCmpv2Args() {}

    private GetBackendConfigCmpv2Args(GetBackendConfigCmpv2Args $) {
        this.backend = $.backend;
        this.disabledValidations = $.disabledValidations;
        this.namespace = $.namespace;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetBackendConfigCmpv2Args defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetBackendConfigCmpv2Args $;

        public Builder() {
            $ = new GetBackendConfigCmpv2Args();
        }

        public Builder(GetBackendConfigCmpv2Args defaults) {
            $ = new GetBackendConfigCmpv2Args(Objects.requireNonNull(defaults));
        }

        /**
         * @param backend The path to the PKI secret backend to
         * read the CMPv2 configuration from, with no leading or trailing `/`s.
         * 
         * # Attributes Reference
         * 
         * @return builder
         * 
         */
        public Builder backend(Output<String> backend) {
            $.backend = backend;
            return this;
        }

        /**
         * @param backend The path to the PKI secret backend to
         * read the CMPv2 configuration from, with no leading or trailing `/`s.
         * 
         * # Attributes Reference
         * 
         * @return builder
         * 
         */
        public Builder backend(String backend) {
            return backend(Output.of(backend));
        }

        /**
         * @param disabledValidations A comma-separated list of validations not to perform on CMPv2 messages.
         * 
         * @return builder
         * 
         */
        public Builder disabledValidations(@Nullable Output<List<String>> disabledValidations) {
            $.disabledValidations = disabledValidations;
            return this;
        }

        /**
         * @param disabledValidations A comma-separated list of validations not to perform on CMPv2 messages.
         * 
         * @return builder
         * 
         */
        public Builder disabledValidations(List<String> disabledValidations) {
            return disabledValidations(Output.of(disabledValidations));
        }

        /**
         * @param disabledValidations A comma-separated list of validations not to perform on CMPv2 messages.
         * 
         * @return builder
         * 
         */
        public Builder disabledValidations(String... disabledValidations) {
            return disabledValidations(List.of(disabledValidations));
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
        public Builder namespace(@Nullable Output<String> namespace) {
            $.namespace = namespace;
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
        public Builder namespace(String namespace) {
            return namespace(Output.of(namespace));
        }

        public GetBackendConfigCmpv2Args build() {
            if ($.backend == null) {
                throw new MissingRequiredPropertyException("GetBackendConfigCmpv2Args", "backend");
            }
            return $;
        }
    }

}
