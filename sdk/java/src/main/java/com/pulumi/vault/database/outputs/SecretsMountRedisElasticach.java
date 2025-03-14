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
public final class SecretsMountRedisElasticach {
    /**
     * @return A list of roles that are allowed to use this
     * connection.
     * 
     */
    private @Nullable List<String> allowedRoles;
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
     * @return Name of the database connection.
     * 
     */
    private String name;
    /**
     * @return The AWS secret key id to use to talk to ElastiCache. If omitted the credentials chain provider is used instead.
     * 
     */
    private @Nullable String password;
    /**
     * @return Specifies the name of the plugin to use.
     * 
     */
    private @Nullable String pluginName;
    /**
     * @return The AWS region where the ElastiCache cluster is hosted. If omitted the plugin tries to infer the region from the environment.
     * 
     */
    private @Nullable String region;
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
     * @return The configuration endpoint for the ElastiCache cluster to connect to.
     * 
     */
    private String url;
    /**
     * @return The AWS access key id to use to talk to ElastiCache. If omitted the credentials chain provider is used instead.
     * 
     */
    private @Nullable String username;
    /**
     * @return Whether the connection should be verified on
     * initial configuration or not.
     * 
     */
    private @Nullable Boolean verifyConnection;

    private SecretsMountRedisElasticach() {}
    /**
     * @return A list of roles that are allowed to use this
     * connection.
     * 
     */
    public List<String> allowedRoles() {
        return this.allowedRoles == null ? List.of() : this.allowedRoles;
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
     * @return Name of the database connection.
     * 
     */
    public String name() {
        return this.name;
    }
    /**
     * @return The AWS secret key id to use to talk to ElastiCache. If omitted the credentials chain provider is used instead.
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
     * @return The AWS region where the ElastiCache cluster is hosted. If omitted the plugin tries to infer the region from the environment.
     * 
     */
    public Optional<String> region() {
        return Optional.ofNullable(this.region);
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
     * @return The configuration endpoint for the ElastiCache cluster to connect to.
     * 
     */
    public String url() {
        return this.url;
    }
    /**
     * @return The AWS access key id to use to talk to ElastiCache. If omitted the credentials chain provider is used instead.
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

    public static Builder builder(SecretsMountRedisElasticach defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable List<String> allowedRoles;
        private @Nullable Map<String,String> data;
        private @Nullable Boolean disableAutomatedRotation;
        private String name;
        private @Nullable String password;
        private @Nullable String pluginName;
        private @Nullable String region;
        private @Nullable List<String> rootRotationStatements;
        private @Nullable Integer rotationPeriod;
        private @Nullable String rotationSchedule;
        private @Nullable Integer rotationWindow;
        private String url;
        private @Nullable String username;
        private @Nullable Boolean verifyConnection;
        public Builder() {}
        public Builder(SecretsMountRedisElasticach defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.allowedRoles = defaults.allowedRoles;
    	      this.data = defaults.data;
    	      this.disableAutomatedRotation = defaults.disableAutomatedRotation;
    	      this.name = defaults.name;
    	      this.password = defaults.password;
    	      this.pluginName = defaults.pluginName;
    	      this.region = defaults.region;
    	      this.rootRotationStatements = defaults.rootRotationStatements;
    	      this.rotationPeriod = defaults.rotationPeriod;
    	      this.rotationSchedule = defaults.rotationSchedule;
    	      this.rotationWindow = defaults.rotationWindow;
    	      this.url = defaults.url;
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
        public Builder name(String name) {
            if (name == null) {
              throw new MissingRequiredPropertyException("SecretsMountRedisElasticach", "name");
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
        public Builder region(@Nullable String region) {

            this.region = region;
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
        public Builder url(String url) {
            if (url == null) {
              throw new MissingRequiredPropertyException("SecretsMountRedisElasticach", "url");
            }
            this.url = url;
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
        public SecretsMountRedisElasticach build() {
            final var _resultValue = new SecretsMountRedisElasticach();
            _resultValue.allowedRoles = allowedRoles;
            _resultValue.data = data;
            _resultValue.disableAutomatedRotation = disableAutomatedRotation;
            _resultValue.name = name;
            _resultValue.password = password;
            _resultValue.pluginName = pluginName;
            _resultValue.region = region;
            _resultValue.rootRotationStatements = rootRotationStatements;
            _resultValue.rotationPeriod = rotationPeriod;
            _resultValue.rotationSchedule = rotationSchedule;
            _resultValue.rotationWindow = rotationWindow;
            _resultValue.url = url;
            _resultValue.username = username;
            _resultValue.verifyConnection = verifyConnection;
            return _resultValue;
        }
    }
}
