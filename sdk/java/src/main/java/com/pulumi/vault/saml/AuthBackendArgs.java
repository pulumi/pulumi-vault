// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vault.saml;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class AuthBackendArgs extends com.pulumi.resources.ResourceArgs {

    public static final AuthBackendArgs Empty = new AuthBackendArgs();

    /**
     * The well-formatted URLs of your Assertion Consumer Service (ACS)
     * that should receive a response from the identity provider.
     * 
     */
    @Import(name="acsUrls", required=true)
    private Output<List<String>> acsUrls;

    /**
     * @return The well-formatted URLs of your Assertion Consumer Service (ACS)
     * that should receive a response from the identity provider.
     * 
     */
    public Output<List<String>> acsUrls() {
        return this.acsUrls;
    }

    /**
     * The role to use if no role is provided during login.
     * 
     */
    @Import(name="defaultRole")
    private @Nullable Output<String> defaultRole;

    /**
     * @return The role to use if no role is provided during login.
     * 
     */
    public Optional<Output<String>> defaultRole() {
        return Optional.ofNullable(this.defaultRole);
    }

    /**
     * If set to `true`, opts out of mount migration on path updates.
     * See here for more info on [Mount Migration](https://www.vaultproject.io/docs/concepts/mount-migration)
     * 
     */
    @Import(name="disableRemount")
    private @Nullable Output<Boolean> disableRemount;

    /**
     * @return If set to `true`, opts out of mount migration on path updates.
     * See here for more info on [Mount Migration](https://www.vaultproject.io/docs/concepts/mount-migration)
     * 
     */
    public Optional<Output<Boolean>> disableRemount() {
        return Optional.ofNullable(this.disableRemount);
    }

    /**
     * The entity ID of the SAML authentication service provider.
     * 
     */
    @Import(name="entityId", required=true)
    private Output<String> entityId;

    /**
     * @return The entity ID of the SAML authentication service provider.
     * 
     */
    public Output<String> entityId() {
        return this.entityId;
    }

    /**
     * The PEM encoded certificate of the identity provider. Mutually exclusive
     * with `idp_metadata_url`.
     * 
     */
    @Import(name="idpCert")
    private @Nullable Output<String> idpCert;

    /**
     * @return The PEM encoded certificate of the identity provider. Mutually exclusive
     * with `idp_metadata_url`.
     * 
     */
    public Optional<Output<String>> idpCert() {
        return Optional.ofNullable(this.idpCert);
    }

    /**
     * The entity ID of the identity provider. Mutually exclusive with
     * `idp_metadata_url`.
     * 
     */
    @Import(name="idpEntityId")
    private @Nullable Output<String> idpEntityId;

    /**
     * @return The entity ID of the identity provider. Mutually exclusive with
     * `idp_metadata_url`.
     * 
     */
    public Optional<Output<String>> idpEntityId() {
        return Optional.ofNullable(this.idpEntityId);
    }

    /**
     * The metadata URL of the identity provider.
     * 
     */
    @Import(name="idpMetadataUrl")
    private @Nullable Output<String> idpMetadataUrl;

    /**
     * @return The metadata URL of the identity provider.
     * 
     */
    public Optional<Output<String>> idpMetadataUrl() {
        return Optional.ofNullable(this.idpMetadataUrl);
    }

    /**
     * The SSO URL of the identity provider. Mutually exclusive with
     * `idp_metadata_url`.
     * 
     */
    @Import(name="idpSsoUrl")
    private @Nullable Output<String> idpSsoUrl;

    /**
     * @return The SSO URL of the identity provider. Mutually exclusive with
     * `idp_metadata_url`.
     * 
     */
    public Optional<Output<String>> idpSsoUrl() {
        return Optional.ofNullable(this.idpSsoUrl);
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
     * Path where the auth backend will be mounted. Defaults to `auth/saml`
     * if not specified.
     * 
     */
    @Import(name="path")
    private @Nullable Output<String> path;

    /**
     * @return Path where the auth backend will be mounted. Defaults to `auth/saml`
     * if not specified.
     * 
     */
    public Optional<Output<String>> path() {
        return Optional.ofNullable(this.path);
    }

    /**
     * If set to `true`, logs additional, potentially sensitive
     * information during the SAML exchange according to the current logging level. Not
     * recommended for production.
     * 
     */
    @Import(name="verboseLogging")
    private @Nullable Output<Boolean> verboseLogging;

    /**
     * @return If set to `true`, logs additional, potentially sensitive
     * information during the SAML exchange according to the current logging level. Not
     * recommended for production.
     * 
     */
    public Optional<Output<Boolean>> verboseLogging() {
        return Optional.ofNullable(this.verboseLogging);
    }

    private AuthBackendArgs() {}

    private AuthBackendArgs(AuthBackendArgs $) {
        this.acsUrls = $.acsUrls;
        this.defaultRole = $.defaultRole;
        this.disableRemount = $.disableRemount;
        this.entityId = $.entityId;
        this.idpCert = $.idpCert;
        this.idpEntityId = $.idpEntityId;
        this.idpMetadataUrl = $.idpMetadataUrl;
        this.idpSsoUrl = $.idpSsoUrl;
        this.namespace = $.namespace;
        this.path = $.path;
        this.verboseLogging = $.verboseLogging;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(AuthBackendArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private AuthBackendArgs $;

        public Builder() {
            $ = new AuthBackendArgs();
        }

        public Builder(AuthBackendArgs defaults) {
            $ = new AuthBackendArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param acsUrls The well-formatted URLs of your Assertion Consumer Service (ACS)
         * that should receive a response from the identity provider.
         * 
         * @return builder
         * 
         */
        public Builder acsUrls(Output<List<String>> acsUrls) {
            $.acsUrls = acsUrls;
            return this;
        }

        /**
         * @param acsUrls The well-formatted URLs of your Assertion Consumer Service (ACS)
         * that should receive a response from the identity provider.
         * 
         * @return builder
         * 
         */
        public Builder acsUrls(List<String> acsUrls) {
            return acsUrls(Output.of(acsUrls));
        }

        /**
         * @param acsUrls The well-formatted URLs of your Assertion Consumer Service (ACS)
         * that should receive a response from the identity provider.
         * 
         * @return builder
         * 
         */
        public Builder acsUrls(String... acsUrls) {
            return acsUrls(List.of(acsUrls));
        }

        /**
         * @param defaultRole The role to use if no role is provided during login.
         * 
         * @return builder
         * 
         */
        public Builder defaultRole(@Nullable Output<String> defaultRole) {
            $.defaultRole = defaultRole;
            return this;
        }

        /**
         * @param defaultRole The role to use if no role is provided during login.
         * 
         * @return builder
         * 
         */
        public Builder defaultRole(String defaultRole) {
            return defaultRole(Output.of(defaultRole));
        }

        /**
         * @param disableRemount If set to `true`, opts out of mount migration on path updates.
         * See here for more info on [Mount Migration](https://www.vaultproject.io/docs/concepts/mount-migration)
         * 
         * @return builder
         * 
         */
        public Builder disableRemount(@Nullable Output<Boolean> disableRemount) {
            $.disableRemount = disableRemount;
            return this;
        }

        /**
         * @param disableRemount If set to `true`, opts out of mount migration on path updates.
         * See here for more info on [Mount Migration](https://www.vaultproject.io/docs/concepts/mount-migration)
         * 
         * @return builder
         * 
         */
        public Builder disableRemount(Boolean disableRemount) {
            return disableRemount(Output.of(disableRemount));
        }

        /**
         * @param entityId The entity ID of the SAML authentication service provider.
         * 
         * @return builder
         * 
         */
        public Builder entityId(Output<String> entityId) {
            $.entityId = entityId;
            return this;
        }

        /**
         * @param entityId The entity ID of the SAML authentication service provider.
         * 
         * @return builder
         * 
         */
        public Builder entityId(String entityId) {
            return entityId(Output.of(entityId));
        }

        /**
         * @param idpCert The PEM encoded certificate of the identity provider. Mutually exclusive
         * with `idp_metadata_url`.
         * 
         * @return builder
         * 
         */
        public Builder idpCert(@Nullable Output<String> idpCert) {
            $.idpCert = idpCert;
            return this;
        }

        /**
         * @param idpCert The PEM encoded certificate of the identity provider. Mutually exclusive
         * with `idp_metadata_url`.
         * 
         * @return builder
         * 
         */
        public Builder idpCert(String idpCert) {
            return idpCert(Output.of(idpCert));
        }

        /**
         * @param idpEntityId The entity ID of the identity provider. Mutually exclusive with
         * `idp_metadata_url`.
         * 
         * @return builder
         * 
         */
        public Builder idpEntityId(@Nullable Output<String> idpEntityId) {
            $.idpEntityId = idpEntityId;
            return this;
        }

        /**
         * @param idpEntityId The entity ID of the identity provider. Mutually exclusive with
         * `idp_metadata_url`.
         * 
         * @return builder
         * 
         */
        public Builder idpEntityId(String idpEntityId) {
            return idpEntityId(Output.of(idpEntityId));
        }

        /**
         * @param idpMetadataUrl The metadata URL of the identity provider.
         * 
         * @return builder
         * 
         */
        public Builder idpMetadataUrl(@Nullable Output<String> idpMetadataUrl) {
            $.idpMetadataUrl = idpMetadataUrl;
            return this;
        }

        /**
         * @param idpMetadataUrl The metadata URL of the identity provider.
         * 
         * @return builder
         * 
         */
        public Builder idpMetadataUrl(String idpMetadataUrl) {
            return idpMetadataUrl(Output.of(idpMetadataUrl));
        }

        /**
         * @param idpSsoUrl The SSO URL of the identity provider. Mutually exclusive with
         * `idp_metadata_url`.
         * 
         * @return builder
         * 
         */
        public Builder idpSsoUrl(@Nullable Output<String> idpSsoUrl) {
            $.idpSsoUrl = idpSsoUrl;
            return this;
        }

        /**
         * @param idpSsoUrl The SSO URL of the identity provider. Mutually exclusive with
         * `idp_metadata_url`.
         * 
         * @return builder
         * 
         */
        public Builder idpSsoUrl(String idpSsoUrl) {
            return idpSsoUrl(Output.of(idpSsoUrl));
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
         * @param path Path where the auth backend will be mounted. Defaults to `auth/saml`
         * if not specified.
         * 
         * @return builder
         * 
         */
        public Builder path(@Nullable Output<String> path) {
            $.path = path;
            return this;
        }

        /**
         * @param path Path where the auth backend will be mounted. Defaults to `auth/saml`
         * if not specified.
         * 
         * @return builder
         * 
         */
        public Builder path(String path) {
            return path(Output.of(path));
        }

        /**
         * @param verboseLogging If set to `true`, logs additional, potentially sensitive
         * information during the SAML exchange according to the current logging level. Not
         * recommended for production.
         * 
         * @return builder
         * 
         */
        public Builder verboseLogging(@Nullable Output<Boolean> verboseLogging) {
            $.verboseLogging = verboseLogging;
            return this;
        }

        /**
         * @param verboseLogging If set to `true`, logs additional, potentially sensitive
         * information during the SAML exchange according to the current logging level. Not
         * recommended for production.
         * 
         * @return builder
         * 
         */
        public Builder verboseLogging(Boolean verboseLogging) {
            return verboseLogging(Output.of(verboseLogging));
        }

        public AuthBackendArgs build() {
            if ($.acsUrls == null) {
                throw new MissingRequiredPropertyException("AuthBackendArgs", "acsUrls");
            }
            if ($.entityId == null) {
                throw new MissingRequiredPropertyException("AuthBackendArgs", "entityId");
            }
            return $;
        }
    }

}
