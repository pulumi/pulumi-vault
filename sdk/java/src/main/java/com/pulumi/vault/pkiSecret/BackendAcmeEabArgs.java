// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vault.pkiSecret;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class BackendAcmeEabArgs extends com.pulumi.resources.ResourceArgs {

    public static final BackendAcmeEabArgs Empty = new BackendAcmeEabArgs();

    /**
     * The path to the PKI secret backend to
     * create the EAB token within, with no leading or trailing `/`s.
     * 
     */
    @Import(name="backend", required=true)
    private Output<String> backend;

    /**
     * @return The path to the PKI secret backend to
     * create the EAB token within, with no leading or trailing `/`s.
     * 
     */
    public Output<String> backend() {
        return this.backend;
    }

    /**
     * Create an EAB token that is specific to an issuer&#39;s ACME directory.
     * 
     */
    @Import(name="issuer")
    private @Nullable Output<String> issuer;

    /**
     * @return Create an EAB token that is specific to an issuer&#39;s ACME directory.
     * 
     */
    public Optional<Output<String>> issuer() {
        return Optional.ofNullable(this.issuer);
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

    /**
     * Create an EAB token that is specific to a role&#39;s ACME directory.
     * 
     * **NOTE**: Within Vault ACME there are different ACME directories which an EAB token is associated with;
     * 
     * 1. Default directory (`pki/acme/`) - Do not specify a value for issuer nor role parameters.
     * 2. Issuer specific (`pki/issuer/:issuer_ref/acme/`) - Specify a value for the issuer parameter
     * 3. Role specific (`pki/roles/:role/acme/`) - Specify a value for the role parameter
     * 4. Issuer and Role specific (`pki/issuer/:issuer_ref/roles/:role/acme/`) - Specify a value for both the issuer and role parameters
     * 
     */
    @Import(name="role")
    private @Nullable Output<String> role;

    /**
     * @return Create an EAB token that is specific to a role&#39;s ACME directory.
     * 
     * **NOTE**: Within Vault ACME there are different ACME directories which an EAB token is associated with;
     * 
     * 1. Default directory (`pki/acme/`) - Do not specify a value for issuer nor role parameters.
     * 2. Issuer specific (`pki/issuer/:issuer_ref/acme/`) - Specify a value for the issuer parameter
     * 3. Role specific (`pki/roles/:role/acme/`) - Specify a value for the role parameter
     * 4. Issuer and Role specific (`pki/issuer/:issuer_ref/roles/:role/acme/`) - Specify a value for both the issuer and role parameters
     * 
     */
    public Optional<Output<String>> role() {
        return Optional.ofNullable(this.role);
    }

    private BackendAcmeEabArgs() {}

    private BackendAcmeEabArgs(BackendAcmeEabArgs $) {
        this.backend = $.backend;
        this.issuer = $.issuer;
        this.namespace = $.namespace;
        this.role = $.role;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(BackendAcmeEabArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private BackendAcmeEabArgs $;

        public Builder() {
            $ = new BackendAcmeEabArgs();
        }

        public Builder(BackendAcmeEabArgs defaults) {
            $ = new BackendAcmeEabArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param backend The path to the PKI secret backend to
         * create the EAB token within, with no leading or trailing `/`s.
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
         * create the EAB token within, with no leading or trailing `/`s.
         * 
         * @return builder
         * 
         */
        public Builder backend(String backend) {
            return backend(Output.of(backend));
        }

        /**
         * @param issuer Create an EAB token that is specific to an issuer&#39;s ACME directory.
         * 
         * @return builder
         * 
         */
        public Builder issuer(@Nullable Output<String> issuer) {
            $.issuer = issuer;
            return this;
        }

        /**
         * @param issuer Create an EAB token that is specific to an issuer&#39;s ACME directory.
         * 
         * @return builder
         * 
         */
        public Builder issuer(String issuer) {
            return issuer(Output.of(issuer));
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

        /**
         * @param role Create an EAB token that is specific to a role&#39;s ACME directory.
         * 
         * **NOTE**: Within Vault ACME there are different ACME directories which an EAB token is associated with;
         * 
         * 1. Default directory (`pki/acme/`) - Do not specify a value for issuer nor role parameters.
         * 2. Issuer specific (`pki/issuer/:issuer_ref/acme/`) - Specify a value for the issuer parameter
         * 3. Role specific (`pki/roles/:role/acme/`) - Specify a value for the role parameter
         * 4. Issuer and Role specific (`pki/issuer/:issuer_ref/roles/:role/acme/`) - Specify a value for both the issuer and role parameters
         * 
         * @return builder
         * 
         */
        public Builder role(@Nullable Output<String> role) {
            $.role = role;
            return this;
        }

        /**
         * @param role Create an EAB token that is specific to a role&#39;s ACME directory.
         * 
         * **NOTE**: Within Vault ACME there are different ACME directories which an EAB token is associated with;
         * 
         * 1. Default directory (`pki/acme/`) - Do not specify a value for issuer nor role parameters.
         * 2. Issuer specific (`pki/issuer/:issuer_ref/acme/`) - Specify a value for the issuer parameter
         * 3. Role specific (`pki/roles/:role/acme/`) - Specify a value for the role parameter
         * 4. Issuer and Role specific (`pki/issuer/:issuer_ref/roles/:role/acme/`) - Specify a value for both the issuer and role parameters
         * 
         * @return builder
         * 
         */
        public Builder role(String role) {
            return role(Output.of(role));
        }

        public BackendAcmeEabArgs build() {
            if ($.backend == null) {
                throw new MissingRequiredPropertyException("BackendAcmeEabArgs", "backend");
            }
            return $;
        }
    }

}
