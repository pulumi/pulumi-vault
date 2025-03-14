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
public final class SecretsMountCassandra {
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
     * @return Cassandra hosts to connect to.
     * 
     */
    private @Nullable List<String> hosts;
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
     * @return The password to use when authenticating with Cassandra.
     * 
     */
    private @Nullable String password;
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
     * @return The transport port to use to connect to Cassandra.
     * 
     */
    private @Nullable Integer port;
    /**
     * @return The CQL protocol version to use.
     * 
     */
    private @Nullable Integer protocolVersion;
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
     * @return Skip permissions checks when a connection to Cassandra is first created. These checks ensure that Vault is able to create roles, but can be resource intensive in clusters with many roles.
     * 
     */
    private @Nullable Boolean skipVerification;
    /**
     * @return Whether to use TLS when connecting to Cassandra.
     * 
     */
    private @Nullable Boolean tls;
    /**
     * @return The username to use when authenticating with Cassandra.
     * 
     */
    private @Nullable String username;
    /**
     * @return Whether the connection should be verified on
     * initial configuration or not.
     * 
     */
    private @Nullable Boolean verifyConnection;

    private SecretsMountCassandra() {}
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
     * @return Cassandra hosts to connect to.
     * 
     */
    public List<String> hosts() {
        return this.hosts == null ? List.of() : this.hosts;
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
     * @return The password to use when authenticating with Cassandra.
     * 
     */
    public Optional<String> password() {
        return Optional.ofNullable(this.password);
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
     * @return The transport port to use to connect to Cassandra.
     * 
     */
    public Optional<Integer> port() {
        return Optional.ofNullable(this.port);
    }
    /**
     * @return The CQL protocol version to use.
     * 
     */
    public Optional<Integer> protocolVersion() {
        return Optional.ofNullable(this.protocolVersion);
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
     * @return Skip permissions checks when a connection to Cassandra is first created. These checks ensure that Vault is able to create roles, but can be resource intensive in clusters with many roles.
     * 
     */
    public Optional<Boolean> skipVerification() {
        return Optional.ofNullable(this.skipVerification);
    }
    /**
     * @return Whether to use TLS when connecting to Cassandra.
     * 
     */
    public Optional<Boolean> tls() {
        return Optional.ofNullable(this.tls);
    }
    /**
     * @return The username to use when authenticating with Cassandra.
     * 
     */
    public Optional<String> username() {
        return Optional.ofNullable(this.username);
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

    public static Builder builder(SecretsMountCassandra defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable List<String> allowedRoles;
        private @Nullable Integer connectTimeout;
        private @Nullable Map<String,String> data;
        private @Nullable Boolean disableAutomatedRotation;
        private @Nullable List<String> hosts;
        private @Nullable Boolean insecureTls;
        private String name;
        private @Nullable String password;
        private @Nullable String pemBundle;
        private @Nullable String pemJson;
        private @Nullable String pluginName;
        private @Nullable Integer port;
        private @Nullable Integer protocolVersion;
        private @Nullable List<String> rootRotationStatements;
        private @Nullable Integer rotationPeriod;
        private @Nullable String rotationSchedule;
        private @Nullable Integer rotationWindow;
        private @Nullable Boolean skipVerification;
        private @Nullable Boolean tls;
        private @Nullable String username;
        private @Nullable Boolean verifyConnection;
        public Builder() {}
        public Builder(SecretsMountCassandra defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.allowedRoles = defaults.allowedRoles;
    	      this.connectTimeout = defaults.connectTimeout;
    	      this.data = defaults.data;
    	      this.disableAutomatedRotation = defaults.disableAutomatedRotation;
    	      this.hosts = defaults.hosts;
    	      this.insecureTls = defaults.insecureTls;
    	      this.name = defaults.name;
    	      this.password = defaults.password;
    	      this.pemBundle = defaults.pemBundle;
    	      this.pemJson = defaults.pemJson;
    	      this.pluginName = defaults.pluginName;
    	      this.port = defaults.port;
    	      this.protocolVersion = defaults.protocolVersion;
    	      this.rootRotationStatements = defaults.rootRotationStatements;
    	      this.rotationPeriod = defaults.rotationPeriod;
    	      this.rotationSchedule = defaults.rotationSchedule;
    	      this.rotationWindow = defaults.rotationWindow;
    	      this.skipVerification = defaults.skipVerification;
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
        public Builder disableAutomatedRotation(@Nullable Boolean disableAutomatedRotation) {

            this.disableAutomatedRotation = disableAutomatedRotation;
            return this;
        }
        @CustomType.Setter
        public Builder hosts(@Nullable List<String> hosts) {

            this.hosts = hosts;
            return this;
        }
        public Builder hosts(String... hosts) {
            return hosts(List.of(hosts));
        }
        @CustomType.Setter
        public Builder insecureTls(@Nullable Boolean insecureTls) {

            this.insecureTls = insecureTls;
            return this;
        }
        @CustomType.Setter
        public Builder name(String name) {
            if (name == null) {
              throw new MissingRequiredPropertyException("SecretsMountCassandra", "name");
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
        public Builder protocolVersion(@Nullable Integer protocolVersion) {

            this.protocolVersion = protocolVersion;
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
        public Builder skipVerification(@Nullable Boolean skipVerification) {

            this.skipVerification = skipVerification;
            return this;
        }
        @CustomType.Setter
        public Builder tls(@Nullable Boolean tls) {

            this.tls = tls;
            return this;
        }
        @CustomType.Setter
        public Builder username(@Nullable String username) {

            this.username = username;
            return this;
        }
        @CustomType.Setter
        public Builder verifyConnection(@Nullable Boolean verifyConnection) {

            this.verifyConnection = verifyConnection;
            return this;
        }
        public SecretsMountCassandra build() {
            final var _resultValue = new SecretsMountCassandra();
            _resultValue.allowedRoles = allowedRoles;
            _resultValue.connectTimeout = connectTimeout;
            _resultValue.data = data;
            _resultValue.disableAutomatedRotation = disableAutomatedRotation;
            _resultValue.hosts = hosts;
            _resultValue.insecureTls = insecureTls;
            _resultValue.name = name;
            _resultValue.password = password;
            _resultValue.pemBundle = pemBundle;
            _resultValue.pemJson = pemJson;
            _resultValue.pluginName = pluginName;
            _resultValue.port = port;
            _resultValue.protocolVersion = protocolVersion;
            _resultValue.rootRotationStatements = rootRotationStatements;
            _resultValue.rotationPeriod = rotationPeriod;
            _resultValue.rotationSchedule = rotationSchedule;
            _resultValue.rotationWindow = rotationWindow;
            _resultValue.skipVerification = skipVerification;
            _resultValue.tls = tls;
            _resultValue.username = username;
            _resultValue.verifyConnection = verifyConnection;
            return _resultValue;
        }
    }
}
