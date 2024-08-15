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
public final class SecretsMountInfluxdb {
    /**
     * @return A list of roles that are allowed to use this
     * connection.
     * 
     */
    private @Nullable List<String> allowedRoles;
    /**
     * @return The number of seconds to use as a connection timeout.
     * 
     */
    private @Nullable Integer connectTimeout;
    /**
     * @return A map of sensitive data to pass to the endpoint. Useful for templated connection strings.
     * 
     * Supported list of database secrets engines that can be configured:
     * 
     */
    private @Nullable Map<String,String> data;
    /**
     * @return Influxdb host to connect to.
     * 
     */
    private String host;
    /**
     * @return Whether to skip verification of the server certificate when using TLS.
     * 
     */
    private @Nullable Boolean insecureTls;
    /**
     * @return Name of the database connection.
     * 
     */
    private String name;
    /**
     * @return Specifies the password corresponding to the given username.
     * 
     */
    private String password;
    /**
     * @return Concatenated PEM blocks containing a certificate and private key; a certificate, private key, and issuing CA certificate; or just a CA certificate.
     * 
     */
    private @Nullable String pemBundle;
    /**
     * @return Specifies JSON containing a certificate and private key; a certificate, private key, and issuing CA certificate; or just a CA certificate.
     * 
     */
    private @Nullable String pemJson;
    /**
     * @return Specifies the name of the plugin to use.
     * 
     */
    private @Nullable String pluginName;
    /**
     * @return The transport port to use to connect to Influxdb.
     * 
     */
    private @Nullable Integer port;
    /**
     * @return A list of database statements to be executed to rotate the root user&#39;s credentials.
     * 
     */
    private @Nullable List<String> rootRotationStatements;
    /**
     * @return Whether to use TLS when connecting to Influxdb.
     * 
     */
    private @Nullable Boolean tls;
    /**
     * @return Specifies the username to use for superuser access.
     * 
     */
    private String username;
    /**
     * @return Template describing how dynamic usernames are generated.
     * 
     */
    private @Nullable String usernameTemplate;
    /**
     * @return Whether the connection should be verified on
     * initial configuration or not.
     * 
     */
    private @Nullable Boolean verifyConnection;

    private SecretsMountInfluxdb() {}
    /**
     * @return A list of roles that are allowed to use this
     * connection.
     * 
     */
    public List<String> allowedRoles() {
        return this.allowedRoles == null ? List.of() : this.allowedRoles;
    }
    /**
     * @return The number of seconds to use as a connection timeout.
     * 
     */
    public Optional<Integer> connectTimeout() {
        return Optional.ofNullable(this.connectTimeout);
    }
    /**
     * @return A map of sensitive data to pass to the endpoint. Useful for templated connection strings.
     * 
     * Supported list of database secrets engines that can be configured:
     * 
     */
    public Map<String,String> data() {
        return this.data == null ? Map.of() : this.data;
    }
    /**
     * @return Influxdb host to connect to.
     * 
     */
    public String host() {
        return this.host;
    }
    /**
     * @return Whether to skip verification of the server certificate when using TLS.
     * 
     */
    public Optional<Boolean> insecureTls() {
        return Optional.ofNullable(this.insecureTls);
    }
    /**
     * @return Name of the database connection.
     * 
     */
    public String name() {
        return this.name;
    }
    /**
     * @return Specifies the password corresponding to the given username.
     * 
     */
    public String password() {
        return this.password;
    }
    /**
     * @return Concatenated PEM blocks containing a certificate and private key; a certificate, private key, and issuing CA certificate; or just a CA certificate.
     * 
     */
    public Optional<String> pemBundle() {
        return Optional.ofNullable(this.pemBundle);
    }
    /**
     * @return Specifies JSON containing a certificate and private key; a certificate, private key, and issuing CA certificate; or just a CA certificate.
     * 
     */
    public Optional<String> pemJson() {
        return Optional.ofNullable(this.pemJson);
    }
    /**
     * @return Specifies the name of the plugin to use.
     * 
     */
    public Optional<String> pluginName() {
        return Optional.ofNullable(this.pluginName);
    }
    /**
     * @return The transport port to use to connect to Influxdb.
     * 
     */
    public Optional<Integer> port() {
        return Optional.ofNullable(this.port);
    }
    /**
     * @return A list of database statements to be executed to rotate the root user&#39;s credentials.
     * 
     */
    public List<String> rootRotationStatements() {
        return this.rootRotationStatements == null ? List.of() : this.rootRotationStatements;
    }
    /**
     * @return Whether to use TLS when connecting to Influxdb.
     * 
     */
    public Optional<Boolean> tls() {
        return Optional.ofNullable(this.tls);
    }
    /**
     * @return Specifies the username to use for superuser access.
     * 
     */
    public String username() {
        return this.username;
    }
    /**
     * @return Template describing how dynamic usernames are generated.
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

    public static Builder builder(SecretsMountInfluxdb defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable List<String> allowedRoles;
        private @Nullable Integer connectTimeout;
        private @Nullable Map<String,String> data;
        private String host;
        private @Nullable Boolean insecureTls;
        private String name;
        private String password;
        private @Nullable String pemBundle;
        private @Nullable String pemJson;
        private @Nullable String pluginName;
        private @Nullable Integer port;
        private @Nullable List<String> rootRotationStatements;
        private @Nullable Boolean tls;
        private String username;
        private @Nullable String usernameTemplate;
        private @Nullable Boolean verifyConnection;
        public Builder() {}
        public Builder(SecretsMountInfluxdb defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.allowedRoles = defaults.allowedRoles;
    	      this.connectTimeout = defaults.connectTimeout;
    	      this.data = defaults.data;
    	      this.host = defaults.host;
    	      this.insecureTls = defaults.insecureTls;
    	      this.name = defaults.name;
    	      this.password = defaults.password;
    	      this.pemBundle = defaults.pemBundle;
    	      this.pemJson = defaults.pemJson;
    	      this.pluginName = defaults.pluginName;
    	      this.port = defaults.port;
    	      this.rootRotationStatements = defaults.rootRotationStatements;
    	      this.tls = defaults.tls;
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
        public Builder connectTimeout(@Nullable Integer connectTimeout) {

            this.connectTimeout = connectTimeout;
            return this;
        }
        @CustomType.Setter
        public Builder data(@Nullable Map<String,String> data) {

            this.data = data;
            return this;
        }
        @CustomType.Setter
        public Builder host(String host) {
            if (host == null) {
              throw new MissingRequiredPropertyException("SecretsMountInfluxdb", "host");
            }
            this.host = host;
            return this;
        }
        @CustomType.Setter
        public Builder insecureTls(@Nullable Boolean insecureTls) {

            this.insecureTls = insecureTls;
            return this;
        }
        @CustomType.Setter
        public Builder name(String name) {
            if (name == null) {
              throw new MissingRequiredPropertyException("SecretsMountInfluxdb", "name");
            }
            this.name = name;
            return this;
        }
        @CustomType.Setter
        public Builder password(String password) {
            if (password == null) {
              throw new MissingRequiredPropertyException("SecretsMountInfluxdb", "password");
            }
            this.password = password;
            return this;
        }
        @CustomType.Setter
        public Builder pemBundle(@Nullable String pemBundle) {

            this.pemBundle = pemBundle;
            return this;
        }
        @CustomType.Setter
        public Builder pemJson(@Nullable String pemJson) {

            this.pemJson = pemJson;
            return this;
        }
        @CustomType.Setter
        public Builder pluginName(@Nullable String pluginName) {

            this.pluginName = pluginName;
            return this;
        }
        @CustomType.Setter
        public Builder port(@Nullable Integer port) {

            this.port = port;
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
        public Builder tls(@Nullable Boolean tls) {

            this.tls = tls;
            return this;
        }
        @CustomType.Setter
        public Builder username(String username) {
            if (username == null) {
              throw new MissingRequiredPropertyException("SecretsMountInfluxdb", "username");
            }
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
        public SecretsMountInfluxdb build() {
            final var _resultValue = new SecretsMountInfluxdb();
            _resultValue.allowedRoles = allowedRoles;
            _resultValue.connectTimeout = connectTimeout;
            _resultValue.data = data;
            _resultValue.host = host;
            _resultValue.insecureTls = insecureTls;
            _resultValue.name = name;
            _resultValue.password = password;
            _resultValue.pemBundle = pemBundle;
            _resultValue.pemJson = pemJson;
            _resultValue.pluginName = pluginName;
            _resultValue.port = port;
            _resultValue.rootRotationStatements = rootRotationStatements;
            _resultValue.tls = tls;
            _resultValue.username = username;
            _resultValue.usernameTemplate = usernameTemplate;
            _resultValue.verifyConnection = verifyConnection;
            return _resultValue;
        }
    }
}
