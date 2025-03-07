// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vault.pkiSecret.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class BackendAcmeEabState extends com.pulumi.resources.ResourceArgs {

    public static final BackendAcmeEabState Empty = new BackendAcmeEabState();

    /**
     * The ACME directory to which the key belongs
     * 
     */
    @Import(name="acmeDirectory")
    private @Nullable Output<String> acmeDirectory;

    /**
     * @return The ACME directory to which the key belongs
     * 
     */
    public Optional<Output<String>> acmeDirectory() {
        return Optional.ofNullable(this.acmeDirectory);
    }

    /**
     * The path to the PKI secret backend to
     * create the EAB token within, with no leading or trailing `/`s.
     * 
     */
    @Import(name="backend")
    private @Nullable Output<String> backend;

    /**
     * @return The path to the PKI secret backend to
     * create the EAB token within, with no leading or trailing `/`s.
     * 
     */
    public Optional<Output<String>> backend() {
        return Optional.ofNullable(this.backend);
    }

    /**
     * An RFC3339 formatted date time when the EAB token was created
     * 
     */
    @Import(name="createdOn")
    private @Nullable Output<String> createdOn;

    /**
     * @return An RFC3339 formatted date time when the EAB token was created
     * 
     */
    public Optional<Output<String>> createdOn() {
        return Optional.ofNullable(this.createdOn);
    }

    /**
     * The identifier of a specific ACME EAB token
     * 
     */
    @Import(name="eabId")
    private @Nullable Output<String> eabId;

    /**
     * @return The identifier of a specific ACME EAB token
     * 
     */
    public Optional<Output<String>> eabId() {
        return Optional.ofNullable(this.eabId);
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
     * The EAB token
     * 
     */
    @Import(name="key")
    private @Nullable Output<String> key;

    /**
     * @return The EAB token
     * 
     */
    public Optional<Output<String>> key() {
        return Optional.ofNullable(this.key);
    }

    /**
     * The key type of the EAB key
     * 
     */
    @Import(name="keyType")
    private @Nullable Output<String> keyType;

    /**
     * @return The key type of the EAB key
     * 
     */
    public Optional<Output<String>> keyType() {
        return Optional.ofNullable(this.keyType);
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

    private BackendAcmeEabState() {}

    private BackendAcmeEabState(BackendAcmeEabState $) {
        this.acmeDirectory = $.acmeDirectory;
        this.backend = $.backend;
        this.createdOn = $.createdOn;
        this.eabId = $.eabId;
        this.issuer = $.issuer;
        this.key = $.key;
        this.keyType = $.keyType;
        this.namespace = $.namespace;
        this.role = $.role;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(BackendAcmeEabState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private BackendAcmeEabState $;

        public Builder() {
            $ = new BackendAcmeEabState();
        }

        public Builder(BackendAcmeEabState defaults) {
            $ = new BackendAcmeEabState(Objects.requireNonNull(defaults));
        }

        /**
         * @param acmeDirectory The ACME directory to which the key belongs
         * 
         * @return builder
         * 
         */
        public Builder acmeDirectory(@Nullable Output<String> acmeDirectory) {
            $.acmeDirectory = acmeDirectory;
            return this;
        }

        /**
         * @param acmeDirectory The ACME directory to which the key belongs
         * 
         * @return builder
         * 
         */
        public Builder acmeDirectory(String acmeDirectory) {
            return acmeDirectory(Output.of(acmeDirectory));
        }

        /**
         * @param backend The path to the PKI secret backend to
         * create the EAB token within, with no leading or trailing `/`s.
         * 
         * @return builder
         * 
         */
        public Builder backend(@Nullable Output<String> backend) {
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
         * @param createdOn An RFC3339 formatted date time when the EAB token was created
         * 
         * @return builder
         * 
         */
        public Builder createdOn(@Nullable Output<String> createdOn) {
            $.createdOn = createdOn;
            return this;
        }

        /**
         * @param createdOn An RFC3339 formatted date time when the EAB token was created
         * 
         * @return builder
         * 
         */
        public Builder createdOn(String createdOn) {
            return createdOn(Output.of(createdOn));
        }

        /**
         * @param eabId The identifier of a specific ACME EAB token
         * 
         * @return builder
         * 
         */
        public Builder eabId(@Nullable Output<String> eabId) {
            $.eabId = eabId;
            return this;
        }

        /**
         * @param eabId The identifier of a specific ACME EAB token
         * 
         * @return builder
         * 
         */
        public Builder eabId(String eabId) {
            return eabId(Output.of(eabId));
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
         * @param key The EAB token
         * 
         * @return builder
         * 
         */
        public Builder key(@Nullable Output<String> key) {
            $.key = key;
            return this;
        }

        /**
         * @param key The EAB token
         * 
         * @return builder
         * 
         */
        public Builder key(String key) {
            return key(Output.of(key));
        }

        /**
         * @param keyType The key type of the EAB key
         * 
         * @return builder
         * 
         */
        public Builder keyType(@Nullable Output<String> keyType) {
            $.keyType = keyType;
            return this;
        }

        /**
         * @param keyType The key type of the EAB key
         * 
         * @return builder
         * 
         */
        public Builder keyType(String keyType) {
            return keyType(Output.of(keyType));
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

        public BackendAcmeEabState build() {
            return $;
        }
    }

}
