// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vault.database.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class SecretsMountMysqlAurora {
    /**
     * @return A list of roles that are allowed to use this
     * connection.
     * 
     */
    private @Nullable List<String> allowedRoles;
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
     * @return A map of sensitive data to pass to the endpoint. Useful for templated connection strings.
     * 
     */
    private @Nullable Map<String,String> data;
    /**
     * @return Cancels all upcoming rotations of the root credential until unset. Requires Vault Enterprise 1.19+.
     * 
     * Supported list of database secrets engines that can be configured:
     * 
     */
    private @Nullable Boolean disableAutomatedRotation;
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
     * @return Name of the database connection.
     * 
     */
    private String name;
    /**
     * @return The root credential password used in the connection URL
     * 
     */
    private @Nullable String password;
    /**
     * @return Specifies the name of the plugin to use.
     * 
     */
    private @Nullable String pluginName;
    /**
     * @return A list of database statements to be executed to rotate the root user&#39;s credentials.
     * 
     */
    private @Nullable List<String> rootRotationStatements;
    /**
     * @return The amount of time in seconds Vault should wait before rotating the root credential.
     * A zero value tells Vault not to rotate the root credential. The minimum rotation period is 10 seconds. Requires Vault Enterprise 1.19+.
     * 
     */
    private @Nullable Integer rotationPeriod;
    /**
     * @return The schedule, in [cron-style time format](https://en.wikipedia.org/wiki/Cron),
     * defining the schedule on which Vault should rotate the root token. Requires Vault Enterprise 1.19+.
     * 
     */
    private @Nullable String rotationSchedule;
    /**
     * @return The maximum amount of time in seconds allowed to complete
     * a rotation when a scheduled token rotation occurs. The default rotation window is
     * unbound and the minimum allowable window is `3600`. Requires Vault Enterprise 1.19+.
     * 
     */
    private @Nullable Integer rotationWindow;
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
    /**
     * @return Whether the connection should be verified on
     * initial configuration or not.
     * 
     */
    private @Nullable Boolean verifyConnection;

    private SecretsMountMysqlAurora() {}
    /**
     * @return A list of roles that are allowed to use this
     * connection.
     * 
     */
    public List<String> allowedRoles() {
        return this.allowedRoles == null ? List.of() : this.allowedRoles;
    }
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
     * @return A map of sensitive data to pass to the endpoint. Useful for templated connection strings.
     * 
     */
    public Map<String,String> data() {
        return this.data == null ? Map.of() : this.data;
    }
    /**
     * @return Cancels all upcoming rotations of the root credential until unset. Requires Vault Enterprise 1.19+.
     * 
     * Supported list of database secrets engines that can be configured:
     * 
     */
    public Optional<Boolean> disableAutomatedRotation() {
        return Optional.ofNullable(this.disableAutomatedRotation);
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
     * @return Name of the database connection.
     * 
     */
    public String name() {
        return this.name;
    }
    /**
     * @return The root credential password used in the connection URL
     * 
     */
    public Optional<String> password() {
        return Optional.ofNullable(this.password);
    }
    /**
     * @return Specifies the name of the plugin to use.
     * 
     */
    public Optional<String> pluginName() {
        return Optional.ofNullable(this.pluginName);
    }
    /**
     * @return A list of database statements to be executed to rotate the root user&#39;s credentials.
     * 
     */
    public List<String> rootRotationStatements() {
        return this.rootRotationStatements == null ? List.of() : this.rootRotationStatements;
    }
    /**
     * @return The amount of time in seconds Vault should wait before rotating the root credential.
     * A zero value tells Vault not to rotate the root credential. The minimum rotation period is 10 seconds. Requires Vault Enterprise 1.19+.
     * 
     */
    public Optional<Integer> rotationPeriod() {
        return Optional.ofNullable(this.rotationPeriod);
    }
    /**
     * @return The schedule, in [cron-style time format](https://en.wikipedia.org/wiki/Cron),
     * defining the schedule on which Vault should rotate the root token. Requires Vault Enterprise 1.19+.
     * 
     */
    public Optional<String> rotationSchedule() {
        return Optional.ofNullable(this.rotationSchedule);
    }
    /**
     * @return The maximum amount of time in seconds allowed to complete
     * a rotation when a scheduled token rotation occurs. The default rotation window is
     * unbound and the minimum allowable window is `3600`. Requires Vault Enterprise 1.19+.
     * 
     */
    public Optional<Integer> rotationWindow() {
        return Optional.ofNullable(this.rotationWindow);
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
    /**
     * @return Whether the connection should be verified on
     * initial configuration or not.
     * 
     */
    public Optional<Boolean> verifyConnection() {
        return Optional.ofNullable(this.verifyConnection);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(SecretsMountMysqlAurora defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable List<String> allowedRoles;
        private @Nullable String authType;
        private @Nullable String connectionUrl;
        private @Nullable Map<String,String> data;
        private @Nullable Boolean disableAutomatedRotation;
        private @Nullable Integer maxConnectionLifetime;
        private @Nullable Integer maxIdleConnections;
        private @Nullable Integer maxOpenConnections;
        private String name;
        private @Nullable String password;
        private @Nullable String pluginName;
        private @Nullable List<String> rootRotationStatements;
        private @Nullable Integer rotationPeriod;
        private @Nullable String rotationSchedule;
        private @Nullable Integer rotationWindow;
        private @Nullable String serviceAccountJson;
        private @Nullable String tlsCa;
        private @Nullable String tlsCertificateKey;
        private @Nullable String username;
        private @Nullable String usernameTemplate;
        private @Nullable Boolean verifyConnection;
        public Builder() {}
        public Builder(SecretsMountMysqlAurora defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.allowedRoles = defaults.allowedRoles;
    	      this.authType = defaults.authType;
    	      this.connectionUrl = defaults.connectionUrl;
    	      this.data = defaults.data;
    	      this.disableAutomatedRotation = defaults.disableAutomatedRotation;
    	      this.maxConnectionLifetime = defaults.maxConnectionLifetime;
    	      this.maxIdleConnections = defaults.maxIdleConnections;
    	      this.maxOpenConnections = defaults.maxOpenConnections;
    	      this.name = defaults.name;
    	      this.password = defaults.password;
    	      this.pluginName = defaults.pluginName;
    	      this.rootRotationStatements = defaults.rootRotationStatements;
    	      this.rotationPeriod = defaults.rotationPeriod;
    	      this.rotationSchedule = defaults.rotationSchedule;
    	      this.rotationWindow = defaults.rotationWindow;
    	      this.serviceAccountJson = defaults.serviceAccountJson;
    	      this.tlsCa = defaults.tlsCa;
    	      this.tlsCertificateKey = defaults.tlsCertificateKey;
    	      this.username = defaults.username;
    	      this.usernameTemplate = defaults.usernameTemplate;
    	      this.verifyConnection = defaults.verifyConnection;
        }

        @CustomType.Setter
        public Builder allowedRoles(@Nullable List<String> allowedRoles) {

            this.allowedRoles = allowedRoles;
            return this;
        }
        public Builder allowedRoles(String... allowedRoles) {
            return allowedRoles(List.of(allowedRoles));
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
        public Builder data(@Nullable Map<String,String> data) {

            this.data = data;
            return this;
        }
        @CustomType.Setter
        public Builder disableAutomatedRotation(@Nullable Boolean disableAutomatedRotation) {

            this.disableAutomatedRotation = disableAutomatedRotation;
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
        public Builder name(String name) {
            if (name == null) {
              throw new MissingRequiredPropertyException("SecretsMountMysqlAurora", "name");
            }
            this.name = name;
            return this;
        }
        @CustomType.Setter
        public Builder password(@Nullable String password) {

            this.password = password;
            return this;
        }
        @CustomType.Setter
        public Builder pluginName(@Nullable String pluginName) {

            this.pluginName = pluginName;
            return this;
        }
        @CustomType.Setter
        public Builder rootRotationStatements(@Nullable List<String> rootRotationStatements) {

            this.rootRotationStatements = rootRotationStatements;
            return this;
        }
        public Builder rootRotationStatements(String... rootRotationStatements) {
            return rootRotationStatements(List.of(rootRotationStatements));
        }
        @CustomType.Setter
        public Builder rotationPeriod(@Nullable Integer rotationPeriod) {

            this.rotationPeriod = rotationPeriod;
            return this;
        }
        @CustomType.Setter
        public Builder rotationSchedule(@Nullable String rotationSchedule) {

            this.rotationSchedule = rotationSchedule;
            return this;
        }
        @CustomType.Setter
        public Builder rotationWindow(@Nullable Integer rotationWindow) {

            this.rotationWindow = rotationWindow;
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
        @CustomType.Setter
        public Builder verifyConnection(@Nullable Boolean verifyConnection) {

            this.verifyConnection = verifyConnection;
            return this;
        }
        public SecretsMountMysqlAurora build() {
            final var _resultValue = new SecretsMountMysqlAurora();
            _resultValue.allowedRoles = allowedRoles;
            _resultValue.authType = authType;
            _resultValue.connectionUrl = connectionUrl;
            _resultValue.data = data;
            _resultValue.disableAutomatedRotation = disableAutomatedRotation;
            _resultValue.maxConnectionLifetime = maxConnectionLifetime;
            _resultValue.maxIdleConnections = maxIdleConnections;
            _resultValue.maxOpenConnections = maxOpenConnections;
            _resultValue.name = name;
            _resultValue.password = password;
            _resultValue.pluginName = pluginName;
            _resultValue.rootRotationStatements = rootRotationStatements;
            _resultValue.rotationPeriod = rotationPeriod;
            _resultValue.rotationSchedule = rotationSchedule;
            _resultValue.rotationWindow = rotationWindow;
            _resultValue.serviceAccountJson = serviceAccountJson;
            _resultValue.tlsCa = tlsCa;
            _resultValue.tlsCertificateKey = tlsCertificateKey;
            _resultValue.username = username;
            _resultValue.usernameTemplate = usernameTemplate;
            _resultValue.verifyConnection = verifyConnection;
            return _resultValue;
        }
    }
}
