// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vault.database.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.Object;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class SecretsMountRedi {
    /**
     * @return A list of roles that are allowed to use this
     * connection.
     * 
     */
    private @Nullable List<String> allowedRoles;
    /**
     * @return The path to a PEM-encoded CA cert file to use to verify the Elasticsearch server&#39;s identity.
     * 
     */
    private @Nullable String caCert;
    /**
     * @return A map of sensitive data to pass to the endpoint. Useful for templated connection strings.
     * 
     * Supported list of database secrets engines that can be configured:
     * 
     */
    private @Nullable Map<String,Object> data;
    /**
     * @return The host to connect to.
     * 
     */
    private String host;
    /**
     * @return Whether to skip verification of the server
     * certificate when using TLS.
     * 
     */
    private @Nullable Boolean insecureTls;
    private String name;
    /**
     * @return The root credential password used in the connection URL.
     * 
     */
    private String password;
    /**
     * @return Specifies the name of the plugin to use.
     * 
     */
    private @Nullable String pluginName;
    /**
     * @return The default port to connect to if no port is specified as
     * part of the host.
     * 
     */
    private @Nullable Integer port;
    /**
     * @return A list of database statements to be executed to rotate the root user&#39;s credentials.
     * 
     */
    private @Nullable List<String> rootRotationStatements;
    /**
     * @return Whether to use TLS when connecting to Cassandra.
     * 
     */
    private @Nullable Boolean tls;
    /**
     * @return The root credential username used in the connection URL.
     * 
     */
    private String username;
    /**
     * @return Whether the connection should be verified on
     * initial configuration or not.
     * 
     */
    private @Nullable Boolean verifyConnection;

    private SecretsMountRedi() {}
    /**
     * @return A list of roles that are allowed to use this
     * connection.
     * 
     */
    public List<String> allowedRoles() {
        return this.allowedRoles == null ? List.of() : this.allowedRoles;
    }
    /**
     * @return The path to a PEM-encoded CA cert file to use to verify the Elasticsearch server&#39;s identity.
     * 
     */
    public Optional<String> caCert() {
        return Optional.ofNullable(this.caCert);
    }
    /**
     * @return A map of sensitive data to pass to the endpoint. Useful for templated connection strings.
     * 
     * Supported list of database secrets engines that can be configured:
     * 
     */
    public Map<String,Object> data() {
        return this.data == null ? Map.of() : this.data;
    }
    /**
     * @return The host to connect to.
     * 
     */
    public String host() {
        return this.host;
    }
    /**
     * @return Whether to skip verification of the server
     * certificate when using TLS.
     * 
     */
    public Optional<Boolean> insecureTls() {
        return Optional.ofNullable(this.insecureTls);
    }
    public String name() {
        return this.name;
    }
    /**
     * @return The root credential password used in the connection URL.
     * 
     */
    public String password() {
        return this.password;
    }
    /**
     * @return Specifies the name of the plugin to use.
     * 
     */
    public Optional<String> pluginName() {
        return Optional.ofNullable(this.pluginName);
    }
    /**
     * @return The default port to connect to if no port is specified as
     * part of the host.
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
     * @return Whether to use TLS when connecting to Cassandra.
     * 
     */
    public Optional<Boolean> tls() {
        return Optional.ofNullable(this.tls);
    }
    /**
     * @return The root credential username used in the connection URL.
     * 
     */
    public String username() {
        return this.username;
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

    public static Builder builder(SecretsMountRedi defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable List<String> allowedRoles;
        private @Nullable String caCert;
        private @Nullable Map<String,Object> data;
        private String host;
        private @Nullable Boolean insecureTls;
        private String name;
        private String password;
        private @Nullable String pluginName;
        private @Nullable Integer port;
        private @Nullable List<String> rootRotationStatements;
        private @Nullable Boolean tls;
        private String username;
        private @Nullable Boolean verifyConnection;
        public Builder() {}
        public Builder(SecretsMountRedi defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.allowedRoles = defaults.allowedRoles;
    	      this.caCert = defaults.caCert;
    	      this.data = defaults.data;
    	      this.host = defaults.host;
    	      this.insecureTls = defaults.insecureTls;
    	      this.name = defaults.name;
    	      this.password = defaults.password;
    	      this.pluginName = defaults.pluginName;
    	      this.port = defaults.port;
    	      this.rootRotationStatements = defaults.rootRotationStatements;
    	      this.tls = defaults.tls;
    	      this.username = defaults.username;
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
        public Builder caCert(@Nullable String caCert) {

            this.caCert = caCert;
            return this;
        }
        @CustomType.Setter
        public Builder data(@Nullable Map<String,Object> data) {

            this.data = data;
            return this;
        }
        @CustomType.Setter
        public Builder host(String host) {
            if (host == null) {
              throw new MissingRequiredPropertyException("SecretsMountRedi", "host");
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
              throw new MissingRequiredPropertyException("SecretsMountRedi", "name");
            }
            this.name = name;
            return this;
        }
        @CustomType.Setter
        public Builder password(String password) {
            if (password == null) {
              throw new MissingRequiredPropertyException("SecretsMountRedi", "password");
            }
            this.password = password;
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
              throw new MissingRequiredPropertyException("SecretsMountRedi", "username");
            }
            this.username = username;
            return this;
        }
        @CustomType.Setter
        public Builder verifyConnection(@Nullable Boolean verifyConnection) {

            this.verifyConnection = verifyConnection;
            return this;
        }
        public SecretsMountRedi build() {
            final var _resultValue = new SecretsMountRedi();
            _resultValue.allowedRoles = allowedRoles;
            _resultValue.caCert = caCert;
            _resultValue.data = data;
            _resultValue.host = host;
            _resultValue.insecureTls = insecureTls;
            _resultValue.name = name;
            _resultValue.password = password;
            _resultValue.pluginName = pluginName;
            _resultValue.port = port;
            _resultValue.rootRotationStatements = rootRotationStatements;
            _resultValue.tls = tls;
            _resultValue.username = username;
            _resultValue.verifyConnection = verifyConnection;
            return _resultValue;
        }
    }
}
