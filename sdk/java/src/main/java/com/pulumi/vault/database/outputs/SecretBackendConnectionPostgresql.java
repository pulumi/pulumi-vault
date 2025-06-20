// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vault.database.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class SecretBackendConnectionPostgresql {
    /**
     * @return Specify alternative authorization type. (Only &#39;gcp_iam&#39; is valid currently)
     * 
     */
    private @Nullable String authType;
    /**
     * @return Connection string to use to connect to the database.
     * 
     */
    private @Nullable String connectionUrl;
    /**
     * @return Disable special character escaping in username and password
     * 
     */
    private @Nullable Boolean disableEscaping;
    /**
     * @return Maximum number of seconds a connection may be reused.
     * 
     */
    private @Nullable Integer maxConnectionLifetime;
    /**
     * @return Maximum number of idle connections to the database.
     * 
     */
    private @Nullable Integer maxIdleConnections;
    /**
     * @return Maximum number of open connections to the database.
     * 
     */
    private @Nullable Integer maxOpenConnections;
    /**
     * @return The root credential password used in the connection URL
     * 
     */
    private @Nullable String password;
    /**
     * @return When set to `scram-sha-256`, passwords will be hashed by Vault before being sent to PostgreSQL.
     * 
     */
    private @Nullable String passwordAuthentication;
    /**
     * @return Version counter for root credential password write-only field
     * 
     */
    private @Nullable Integer passwordWoVersion;
    /**
     * @return The secret key used for the x509 client certificate. Must be PEM encoded.
     * 
     */
    private @Nullable String privateKey;
    /**
     * @return If set, allows onboarding static roles with a rootless connection configuration.
     * 
     */
    private @Nullable Boolean selfManaged;
    /**
     * @return A JSON encoded credential for use with IAM authorization
     * 
     */
    private @Nullable String serviceAccountJson;
    /**
     * @return The x509 CA file for validating the certificate presented by the PostgreSQL server. Must be PEM encoded.
     * 
     */
    private @Nullable String tlsCa;
    /**
     * @return The x509 client certificate for connecting to the database. Must be PEM encoded.
     * 
     */
    private @Nullable String tlsCertificate;
    /**
     * @return The root credential username used in the connection URL
     * 
     */
    private @Nullable String username;
    /**
     * @return Username generation template.
     * 
     */
    private @Nullable String usernameTemplate;

    private SecretBackendConnectionPostgresql() {}
    /**
     * @return Specify alternative authorization type. (Only &#39;gcp_iam&#39; is valid currently)
     * 
     */
    public Optional<String> authType() {
        return Optional.ofNullable(this.authType);
    }
    /**
     * @return Connection string to use to connect to the database.
     * 
     */
    public Optional<String> connectionUrl() {
        return Optional.ofNullable(this.connectionUrl);
    }
    /**
     * @return Disable special character escaping in username and password
     * 
     */
    public Optional<Boolean> disableEscaping() {
        return Optional.ofNullable(this.disableEscaping);
    }
    /**
     * @return Maximum number of seconds a connection may be reused.
     * 
     */
    public Optional<Integer> maxConnectionLifetime() {
        return Optional.ofNullable(this.maxConnectionLifetime);
    }
    /**
     * @return Maximum number of idle connections to the database.
     * 
     */
    public Optional<Integer> maxIdleConnections() {
        return Optional.ofNullable(this.maxIdleConnections);
    }
    /**
     * @return Maximum number of open connections to the database.
     * 
     */
    public Optional<Integer> maxOpenConnections() {
        return Optional.ofNullable(this.maxOpenConnections);
    }
    /**
     * @return The root credential password used in the connection URL
     * 
     */
    public Optional<String> password() {
        return Optional.ofNullable(this.password);
    }
    /**
     * @return When set to `scram-sha-256`, passwords will be hashed by Vault before being sent to PostgreSQL.
     * 
     */
    public Optional<String> passwordAuthentication() {
        return Optional.ofNullable(this.passwordAuthentication);
    }
    /**
     * @return Version counter for root credential password write-only field
     * 
     */
    public Optional<Integer> passwordWoVersion() {
        return Optional.ofNullable(this.passwordWoVersion);
    }
    /**
     * @return The secret key used for the x509 client certificate. Must be PEM encoded.
     * 
     */
    public Optional<String> privateKey() {
        return Optional.ofNullable(this.privateKey);
    }
    /**
     * @return If set, allows onboarding static roles with a rootless connection configuration.
     * 
     */
    public Optional<Boolean> selfManaged() {
        return Optional.ofNullable(this.selfManaged);
    }
    /**
     * @return A JSON encoded credential for use with IAM authorization
     * 
     */
    public Optional<String> serviceAccountJson() {
        return Optional.ofNullable(this.serviceAccountJson);
    }
    /**
     * @return The x509 CA file for validating the certificate presented by the PostgreSQL server. Must be PEM encoded.
     * 
     */
    public Optional<String> tlsCa() {
        return Optional.ofNullable(this.tlsCa);
    }
    /**
     * @return The x509 client certificate for connecting to the database. Must be PEM encoded.
     * 
     */
    public Optional<String> tlsCertificate() {
        return Optional.ofNullable(this.tlsCertificate);
    }
    /**
     * @return The root credential username used in the connection URL
     * 
     */
    public Optional<String> username() {
        return Optional.ofNullable(this.username);
    }
    /**
     * @return Username generation template.
     * 
     */
    public Optional<String> usernameTemplate() {
        return Optional.ofNullable(this.usernameTemplate);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(SecretBackendConnectionPostgresql defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable String authType;
        private @Nullable String connectionUrl;
        private @Nullable Boolean disableEscaping;
        private @Nullable Integer maxConnectionLifetime;
        private @Nullable Integer maxIdleConnections;
        private @Nullable Integer maxOpenConnections;
        private @Nullable String password;
        private @Nullable String passwordAuthentication;
        private @Nullable Integer passwordWoVersion;
        private @Nullable String privateKey;
        private @Nullable Boolean selfManaged;
        private @Nullable String serviceAccountJson;
        private @Nullable String tlsCa;
        private @Nullable String tlsCertificate;
        private @Nullable String username;
        private @Nullable String usernameTemplate;
        public Builder() {}
        public Builder(SecretBackendConnectionPostgresql defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.authType = defaults.authType;
    	      this.connectionUrl = defaults.connectionUrl;
    	      this.disableEscaping = defaults.disableEscaping;
    	      this.maxConnectionLifetime = defaults.maxConnectionLifetime;
    	      this.maxIdleConnections = defaults.maxIdleConnections;
    	      this.maxOpenConnections = defaults.maxOpenConnections;
    	      this.password = defaults.password;
    	      this.passwordAuthentication = defaults.passwordAuthentication;
    	      this.passwordWoVersion = defaults.passwordWoVersion;
    	      this.privateKey = defaults.privateKey;
    	      this.selfManaged = defaults.selfManaged;
    	      this.serviceAccountJson = defaults.serviceAccountJson;
    	      this.tlsCa = defaults.tlsCa;
    	      this.tlsCertificate = defaults.tlsCertificate;
    	      this.username = defaults.username;
    	      this.usernameTemplate = defaults.usernameTemplate;
        }

        @CustomType.Setter
        public Builder authType(@Nullable String authType) {

            this.authType = authType;
            return this;
        }
        @CustomType.Setter
        public Builder connectionUrl(@Nullable String connectionUrl) {

            this.connectionUrl = connectionUrl;
            return this;
        }
        @CustomType.Setter
        public Builder disableEscaping(@Nullable Boolean disableEscaping) {

            this.disableEscaping = disableEscaping;
            return this;
        }
        @CustomType.Setter
        public Builder maxConnectionLifetime(@Nullable Integer maxConnectionLifetime) {

            this.maxConnectionLifetime = maxConnectionLifetime;
            return this;
        }
        @CustomType.Setter
        public Builder maxIdleConnections(@Nullable Integer maxIdleConnections) {

            this.maxIdleConnections = maxIdleConnections;
            return this;
        }
        @CustomType.Setter
        public Builder maxOpenConnections(@Nullable Integer maxOpenConnections) {

            this.maxOpenConnections = maxOpenConnections;
            return this;
        }
        @CustomType.Setter
        public Builder password(@Nullable String password) {

            this.password = password;
            return this;
        }
        @CustomType.Setter
        public Builder passwordAuthentication(@Nullable String passwordAuthentication) {

            this.passwordAuthentication = passwordAuthentication;
            return this;
        }
        @CustomType.Setter
        public Builder passwordWoVersion(@Nullable Integer passwordWoVersion) {

            this.passwordWoVersion = passwordWoVersion;
            return this;
        }
        @CustomType.Setter
        public Builder privateKey(@Nullable String privateKey) {

            this.privateKey = privateKey;
            return this;
        }
        @CustomType.Setter
        public Builder selfManaged(@Nullable Boolean selfManaged) {

            this.selfManaged = selfManaged;
            return this;
        }
        @CustomType.Setter
        public Builder serviceAccountJson(@Nullable String serviceAccountJson) {

            this.serviceAccountJson = serviceAccountJson;
            return this;
        }
        @CustomType.Setter
        public Builder tlsCa(@Nullable String tlsCa) {

            this.tlsCa = tlsCa;
            return this;
        }
        @CustomType.Setter
        public Builder tlsCertificate(@Nullable String tlsCertificate) {

            this.tlsCertificate = tlsCertificate;
            return this;
        }
        @CustomType.Setter
        public Builder username(@Nullable String username) {

            this.username = username;
            return this;
        }
        @CustomType.Setter
        public Builder usernameTemplate(@Nullable String usernameTemplate) {

            this.usernameTemplate = usernameTemplate;
            return this;
        }
        public SecretBackendConnectionPostgresql build() {
            final var _resultValue = new SecretBackendConnectionPostgresql();
            _resultValue.authType = authType;
            _resultValue.connectionUrl = connectionUrl;
            _resultValue.disableEscaping = disableEscaping;
            _resultValue.maxConnectionLifetime = maxConnectionLifetime;
            _resultValue.maxIdleConnections = maxIdleConnections;
            _resultValue.maxOpenConnections = maxOpenConnections;
            _resultValue.password = password;
            _resultValue.passwordAuthentication = passwordAuthentication;
            _resultValue.passwordWoVersion = passwordWoVersion;
            _resultValue.privateKey = privateKey;
            _resultValue.selfManaged = selfManaged;
            _resultValue.serviceAccountJson = serviceAccountJson;
            _resultValue.tlsCa = tlsCa;
            _resultValue.tlsCertificate = tlsCertificate;
            _resultValue.username = username;
            _resultValue.usernameTemplate = usernameTemplate;
            return _resultValue;
        }
    }
}
