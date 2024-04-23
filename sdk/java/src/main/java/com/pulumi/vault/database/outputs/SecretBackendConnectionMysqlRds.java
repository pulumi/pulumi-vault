// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vault.database.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class SecretBackendConnectionMysqlRds {
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
     * @return A JSON encoded credential for use with IAM authorization
     * 
     */
    private @Nullable String serviceAccountJson;
    /**
     * @return x509 CA file for validating the certificate presented by the MySQL server. Must be PEM encoded.
     * 
     */
    private @Nullable String tlsCa;
    /**
     * @return x509 certificate for connecting to the database. This must be a PEM encoded version of the private key and the certificate combined.
     * 
     */
    private @Nullable String tlsCertificateKey;
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

    private SecretBackendConnectionMysqlRds() {}
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
     * @return A JSON encoded credential for use with IAM authorization
     * 
     */
    public Optional<String> serviceAccountJson() {
        return Optional.ofNullable(this.serviceAccountJson);
    }
    /**
     * @return x509 CA file for validating the certificate presented by the MySQL server. Must be PEM encoded.
     * 
     */
    public Optional<String> tlsCa() {
        return Optional.ofNullable(this.tlsCa);
    }
    /**
     * @return x509 certificate for connecting to the database. This must be a PEM encoded version of the private key and the certificate combined.
     * 
     */
    public Optional<String> tlsCertificateKey() {
        return Optional.ofNullable(this.tlsCertificateKey);
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

    public static Builder builder(SecretBackendConnectionMysqlRds defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable String authType;
        private @Nullable String connectionUrl;
        private @Nullable Integer maxConnectionLifetime;
        private @Nullable Integer maxIdleConnections;
        private @Nullable Integer maxOpenConnections;
        private @Nullable String password;
        private @Nullable String serviceAccountJson;
        private @Nullable String tlsCa;
        private @Nullable String tlsCertificateKey;
        private @Nullable String username;
        private @Nullable String usernameTemplate;
        public Builder() {}
        public Builder(SecretBackendConnectionMysqlRds defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.authType = defaults.authType;
    	      this.connectionUrl = defaults.connectionUrl;
    	      this.maxConnectionLifetime = defaults.maxConnectionLifetime;
    	      this.maxIdleConnections = defaults.maxIdleConnections;
    	      this.maxOpenConnections = defaults.maxOpenConnections;
    	      this.password = defaults.password;
    	      this.serviceAccountJson = defaults.serviceAccountJson;
    	      this.tlsCa = defaults.tlsCa;
    	      this.tlsCertificateKey = defaults.tlsCertificateKey;
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
        public Builder tlsCertificateKey(@Nullable String tlsCertificateKey) {

            this.tlsCertificateKey = tlsCertificateKey;
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
        public SecretBackendConnectionMysqlRds build() {
            final var _resultValue = new SecretBackendConnectionMysqlRds();
            _resultValue.authType = authType;
            _resultValue.connectionUrl = connectionUrl;
            _resultValue.maxConnectionLifetime = maxConnectionLifetime;
            _resultValue.maxIdleConnections = maxIdleConnections;
            _resultValue.maxOpenConnections = maxOpenConnections;
            _resultValue.password = password;
            _resultValue.serviceAccountJson = serviceAccountJson;
            _resultValue.tlsCa = tlsCa;
            _resultValue.tlsCertificateKey = tlsCertificateKey;
            _resultValue.username = username;
            _resultValue.usernameTemplate = usernameTemplate;
            return _resultValue;
        }
    }
}
