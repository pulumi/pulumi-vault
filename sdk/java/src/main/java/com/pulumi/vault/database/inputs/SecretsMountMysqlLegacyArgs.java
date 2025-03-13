// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vault.database.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class SecretsMountMysqlLegacyArgs extends com.pulumi.resources.ResourceArgs {

    public static final SecretsMountMysqlLegacyArgs Empty = new SecretsMountMysqlLegacyArgs();

    /**
     * A list of roles that are allowed to use this
     * connection.
     * 
     */
    @Import(name="allowedRoles")
    private @Nullable Output<List<String>> allowedRoles;

    /**
     * @return A list of roles that are allowed to use this
     * connection.
     * 
     */
    public Optional<Output<List<String>>> allowedRoles() {
        return Optional.ofNullable(this.allowedRoles);
    }

    /**
     * Specify alternative authorization type. (Only &#39;gcp_iam&#39; is valid currently)
     * 
     */
    @Import(name="authType")
    private @Nullable Output<String> authType;

    /**
     * @return Specify alternative authorization type. (Only &#39;gcp_iam&#39; is valid currently)
     * 
     */
    public Optional<Output<String>> authType() {
        return Optional.ofNullable(this.authType);
    }

    /**
     * Connection string to use to connect to the database.
     * 
     */
    @Import(name="connectionUrl")
    private @Nullable Output<String> connectionUrl;

    /**
     * @return Connection string to use to connect to the database.
     * 
     */
    public Optional<Output<String>> connectionUrl() {
        return Optional.ofNullable(this.connectionUrl);
    }

    /**
     * A map of sensitive data to pass to the endpoint. Useful for templated connection strings.
     * 
     */
    @Import(name="data")
    private @Nullable Output<Map<String,String>> data;

    /**
     * @return A map of sensitive data to pass to the endpoint. Useful for templated connection strings.
     * 
     */
    public Optional<Output<Map<String,String>>> data() {
        return Optional.ofNullable(this.data);
    }

    /**
     * Cancels all upcoming rotations of the root credential until unset. Requires Vault Enterprise 1.19+.
     * 
     * Supported list of database secrets engines that can be configured:
     * 
     */
    @Import(name="disableAutomatedRotation")
    private @Nullable Output<Boolean> disableAutomatedRotation;

    /**
     * @return Cancels all upcoming rotations of the root credential until unset. Requires Vault Enterprise 1.19+.
     * 
     * Supported list of database secrets engines that can be configured:
     * 
     */
    public Optional<Output<Boolean>> disableAutomatedRotation() {
        return Optional.ofNullable(this.disableAutomatedRotation);
    }

    /**
     * Maximum number of seconds a connection may be reused.
     * 
     */
    @Import(name="maxConnectionLifetime")
    private @Nullable Output<Integer> maxConnectionLifetime;

    /**
     * @return Maximum number of seconds a connection may be reused.
     * 
     */
    public Optional<Output<Integer>> maxConnectionLifetime() {
        return Optional.ofNullable(this.maxConnectionLifetime);
    }

    /**
     * Maximum number of idle connections to the database.
     * 
     */
    @Import(name="maxIdleConnections")
    private @Nullable Output<Integer> maxIdleConnections;

    /**
     * @return Maximum number of idle connections to the database.
     * 
     */
    public Optional<Output<Integer>> maxIdleConnections() {
        return Optional.ofNullable(this.maxIdleConnections);
    }

    /**
     * Maximum number of open connections to the database.
     * 
     */
    @Import(name="maxOpenConnections")
    private @Nullable Output<Integer> maxOpenConnections;

    /**
     * @return Maximum number of open connections to the database.
     * 
     */
    public Optional<Output<Integer>> maxOpenConnections() {
        return Optional.ofNullable(this.maxOpenConnections);
    }

    /**
     * Name of the database connection.
     * 
     */
    @Import(name="name", required=true)
    private Output<String> name;

    /**
     * @return Name of the database connection.
     * 
     */
    public Output<String> name() {
        return this.name;
    }

    /**
     * The root credential password used in the connection URL
     * 
     */
    @Import(name="password")
    private @Nullable Output<String> password;

    /**
     * @return The root credential password used in the connection URL
     * 
     */
    public Optional<Output<String>> password() {
        return Optional.ofNullable(this.password);
    }

    /**
     * Specifies the name of the plugin to use.
     * 
     */
    @Import(name="pluginName")
    private @Nullable Output<String> pluginName;

    /**
     * @return Specifies the name of the plugin to use.
     * 
     */
    public Optional<Output<String>> pluginName() {
        return Optional.ofNullable(this.pluginName);
    }

    /**
     * A list of database statements to be executed to rotate the root user&#39;s credentials.
     * 
     */
    @Import(name="rootRotationStatements")
    private @Nullable Output<List<String>> rootRotationStatements;

    /**
     * @return A list of database statements to be executed to rotate the root user&#39;s credentials.
     * 
     */
    public Optional<Output<List<String>>> rootRotationStatements() {
        return Optional.ofNullable(this.rootRotationStatements);
    }

    /**
     * The amount of time in seconds Vault should wait before rotating the root credential.
     * A zero value tells Vault not to rotate the root credential. The minimum rotation period is 10 seconds. Requires Vault Enterprise 1.19+.
     * 
     */
    @Import(name="rotationPeriod")
    private @Nullable Output<Integer> rotationPeriod;

    /**
     * @return The amount of time in seconds Vault should wait before rotating the root credential.
     * A zero value tells Vault not to rotate the root credential. The minimum rotation period is 10 seconds. Requires Vault Enterprise 1.19+.
     * 
     */
    public Optional<Output<Integer>> rotationPeriod() {
        return Optional.ofNullable(this.rotationPeriod);
    }

    /**
     * The schedule, in [cron-style time format](https://en.wikipedia.org/wiki/Cron),
     * defining the schedule on which Vault should rotate the root token. Requires Vault Enterprise 1.19+.
     * 
     */
    @Import(name="rotationSchedule")
    private @Nullable Output<String> rotationSchedule;

    /**
     * @return The schedule, in [cron-style time format](https://en.wikipedia.org/wiki/Cron),
     * defining the schedule on which Vault should rotate the root token. Requires Vault Enterprise 1.19+.
     * 
     */
    public Optional<Output<String>> rotationSchedule() {
        return Optional.ofNullable(this.rotationSchedule);
    }

    /**
     * The maximum amount of time in seconds allowed to complete
     * a rotation when a scheduled token rotation occurs. The default rotation window is
     * unbound and the minimum allowable window is `3600`. Requires Vault Enterprise 1.19+.
     * 
     */
    @Import(name="rotationWindow")
    private @Nullable Output<Integer> rotationWindow;

    /**
     * @return The maximum amount of time in seconds allowed to complete
     * a rotation when a scheduled token rotation occurs. The default rotation window is
     * unbound and the minimum allowable window is `3600`. Requires Vault Enterprise 1.19+.
     * 
     */
    public Optional<Output<Integer>> rotationWindow() {
        return Optional.ofNullable(this.rotationWindow);
    }

    /**
     * A JSON encoded credential for use with IAM authorization
     * 
     */
    @Import(name="serviceAccountJson")
    private @Nullable Output<String> serviceAccountJson;

    /**
     * @return A JSON encoded credential for use with IAM authorization
     * 
     */
    public Optional<Output<String>> serviceAccountJson() {
        return Optional.ofNullable(this.serviceAccountJson);
    }

    /**
     * x509 CA file for validating the certificate presented by the MySQL server. Must be PEM encoded.
     * 
     */
    @Import(name="tlsCa")
    private @Nullable Output<String> tlsCa;

    /**
     * @return x509 CA file for validating the certificate presented by the MySQL server. Must be PEM encoded.
     * 
     */
    public Optional<Output<String>> tlsCa() {
        return Optional.ofNullable(this.tlsCa);
    }

    /**
     * x509 certificate for connecting to the database. This must be a PEM encoded version of the private key and the certificate combined.
     * 
     */
    @Import(name="tlsCertificateKey")
    private @Nullable Output<String> tlsCertificateKey;

    /**
     * @return x509 certificate for connecting to the database. This must be a PEM encoded version of the private key and the certificate combined.
     * 
     */
    public Optional<Output<String>> tlsCertificateKey() {
        return Optional.ofNullable(this.tlsCertificateKey);
    }

    /**
     * The root credential username used in the connection URL
     * 
     */
    @Import(name="username")
    private @Nullable Output<String> username;

    /**
     * @return The root credential username used in the connection URL
     * 
     */
    public Optional<Output<String>> username() {
        return Optional.ofNullable(this.username);
    }

    /**
     * Username generation template.
     * 
     */
    @Import(name="usernameTemplate")
    private @Nullable Output<String> usernameTemplate;

    /**
     * @return Username generation template.
     * 
     */
    public Optional<Output<String>> usernameTemplate() {
        return Optional.ofNullable(this.usernameTemplate);
    }

    /**
     * Whether the connection should be verified on
     * initial configuration or not.
     * 
     */
    @Import(name="verifyConnection")
    private @Nullable Output<Boolean> verifyConnection;

    /**
     * @return Whether the connection should be verified on
     * initial configuration or not.
     * 
     */
    public Optional<Output<Boolean>> verifyConnection() {
        return Optional.ofNullable(this.verifyConnection);
    }

    private SecretsMountMysqlLegacyArgs() {}

    private SecretsMountMysqlLegacyArgs(SecretsMountMysqlLegacyArgs $) {
        this.allowedRoles = $.allowedRoles;
        this.authType = $.authType;
        this.connectionUrl = $.connectionUrl;
        this.data = $.data;
        this.disableAutomatedRotation = $.disableAutomatedRotation;
        this.maxConnectionLifetime = $.maxConnectionLifetime;
        this.maxIdleConnections = $.maxIdleConnections;
        this.maxOpenConnections = $.maxOpenConnections;
        this.name = $.name;
        this.password = $.password;
        this.pluginName = $.pluginName;
        this.rootRotationStatements = $.rootRotationStatements;
        this.rotationPeriod = $.rotationPeriod;
        this.rotationSchedule = $.rotationSchedule;
        this.rotationWindow = $.rotationWindow;
        this.serviceAccountJson = $.serviceAccountJson;
        this.tlsCa = $.tlsCa;
        this.tlsCertificateKey = $.tlsCertificateKey;
        this.username = $.username;
        this.usernameTemplate = $.usernameTemplate;
        this.verifyConnection = $.verifyConnection;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(SecretsMountMysqlLegacyArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private SecretsMountMysqlLegacyArgs $;

        public Builder() {
            $ = new SecretsMountMysqlLegacyArgs();
        }

        public Builder(SecretsMountMysqlLegacyArgs defaults) {
            $ = new SecretsMountMysqlLegacyArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param allowedRoles A list of roles that are allowed to use this
         * connection.
         * 
         * @return builder
         * 
         */
        public Builder allowedRoles(@Nullable Output<List<String>> allowedRoles) {
            $.allowedRoles = allowedRoles;
            return this;
        }

        /**
         * @param allowedRoles A list of roles that are allowed to use this
         * connection.
         * 
         * @return builder
         * 
         */
        public Builder allowedRoles(List<String> allowedRoles) {
            return allowedRoles(Output.of(allowedRoles));
        }

        /**
         * @param allowedRoles A list of roles that are allowed to use this
         * connection.
         * 
         * @return builder
         * 
         */
        public Builder allowedRoles(String... allowedRoles) {
            return allowedRoles(List.of(allowedRoles));
        }

        /**
         * @param authType Specify alternative authorization type. (Only &#39;gcp_iam&#39; is valid currently)
         * 
         * @return builder
         * 
         */
        public Builder authType(@Nullable Output<String> authType) {
            $.authType = authType;
            return this;
        }

        /**
         * @param authType Specify alternative authorization type. (Only &#39;gcp_iam&#39; is valid currently)
         * 
         * @return builder
         * 
         */
        public Builder authType(String authType) {
            return authType(Output.of(authType));
        }

        /**
         * @param connectionUrl Connection string to use to connect to the database.
         * 
         * @return builder
         * 
         */
        public Builder connectionUrl(@Nullable Output<String> connectionUrl) {
            $.connectionUrl = connectionUrl;
            return this;
        }

        /**
         * @param connectionUrl Connection string to use to connect to the database.
         * 
         * @return builder
         * 
         */
        public Builder connectionUrl(String connectionUrl) {
            return connectionUrl(Output.of(connectionUrl));
        }

        /**
         * @param data A map of sensitive data to pass to the endpoint. Useful for templated connection strings.
         * 
         * @return builder
         * 
         */
        public Builder data(@Nullable Output<Map<String,String>> data) {
            $.data = data;
            return this;
        }

        /**
         * @param data A map of sensitive data to pass to the endpoint. Useful for templated connection strings.
         * 
         * @return builder
         * 
         */
        public Builder data(Map<String,String> data) {
            return data(Output.of(data));
        }

        /**
         * @param disableAutomatedRotation Cancels all upcoming rotations of the root credential until unset. Requires Vault Enterprise 1.19+.
         * 
         * Supported list of database secrets engines that can be configured:
         * 
         * @return builder
         * 
         */
        public Builder disableAutomatedRotation(@Nullable Output<Boolean> disableAutomatedRotation) {
            $.disableAutomatedRotation = disableAutomatedRotation;
            return this;
        }

        /**
         * @param disableAutomatedRotation Cancels all upcoming rotations of the root credential until unset. Requires Vault Enterprise 1.19+.
         * 
         * Supported list of database secrets engines that can be configured:
         * 
         * @return builder
         * 
         */
        public Builder disableAutomatedRotation(Boolean disableAutomatedRotation) {
            return disableAutomatedRotation(Output.of(disableAutomatedRotation));
        }

        /**
         * @param maxConnectionLifetime Maximum number of seconds a connection may be reused.
         * 
         * @return builder
         * 
         */
        public Builder maxConnectionLifetime(@Nullable Output<Integer> maxConnectionLifetime) {
            $.maxConnectionLifetime = maxConnectionLifetime;
            return this;
        }

        /**
         * @param maxConnectionLifetime Maximum number of seconds a connection may be reused.
         * 
         * @return builder
         * 
         */
        public Builder maxConnectionLifetime(Integer maxConnectionLifetime) {
            return maxConnectionLifetime(Output.of(maxConnectionLifetime));
        }

        /**
         * @param maxIdleConnections Maximum number of idle connections to the database.
         * 
         * @return builder
         * 
         */
        public Builder maxIdleConnections(@Nullable Output<Integer> maxIdleConnections) {
            $.maxIdleConnections = maxIdleConnections;
            return this;
        }

        /**
         * @param maxIdleConnections Maximum number of idle connections to the database.
         * 
         * @return builder
         * 
         */
        public Builder maxIdleConnections(Integer maxIdleConnections) {
            return maxIdleConnections(Output.of(maxIdleConnections));
        }

        /**
         * @param maxOpenConnections Maximum number of open connections to the database.
         * 
         * @return builder
         * 
         */
        public Builder maxOpenConnections(@Nullable Output<Integer> maxOpenConnections) {
            $.maxOpenConnections = maxOpenConnections;
            return this;
        }

        /**
         * @param maxOpenConnections Maximum number of open connections to the database.
         * 
         * @return builder
         * 
         */
        public Builder maxOpenConnections(Integer maxOpenConnections) {
            return maxOpenConnections(Output.of(maxOpenConnections));
        }

        /**
         * @param name Name of the database connection.
         * 
         * @return builder
         * 
         */
        public Builder name(Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name Name of the database connection.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param password The root credential password used in the connection URL
         * 
         * @return builder
         * 
         */
        public Builder password(@Nullable Output<String> password) {
            $.password = password;
            return this;
        }

        /**
         * @param password The root credential password used in the connection URL
         * 
         * @return builder
         * 
         */
        public Builder password(String password) {
            return password(Output.of(password));
        }

        /**
         * @param pluginName Specifies the name of the plugin to use.
         * 
         * @return builder
         * 
         */
        public Builder pluginName(@Nullable Output<String> pluginName) {
            $.pluginName = pluginName;
            return this;
        }

        /**
         * @param pluginName Specifies the name of the plugin to use.
         * 
         * @return builder
         * 
         */
        public Builder pluginName(String pluginName) {
            return pluginName(Output.of(pluginName));
        }

        /**
         * @param rootRotationStatements A list of database statements to be executed to rotate the root user&#39;s credentials.
         * 
         * @return builder
         * 
         */
        public Builder rootRotationStatements(@Nullable Output<List<String>> rootRotationStatements) {
            $.rootRotationStatements = rootRotationStatements;
            return this;
        }

        /**
         * @param rootRotationStatements A list of database statements to be executed to rotate the root user&#39;s credentials.
         * 
         * @return builder
         * 
         */
        public Builder rootRotationStatements(List<String> rootRotationStatements) {
            return rootRotationStatements(Output.of(rootRotationStatements));
        }

        /**
         * @param rootRotationStatements A list of database statements to be executed to rotate the root user&#39;s credentials.
         * 
         * @return builder
         * 
         */
        public Builder rootRotationStatements(String... rootRotationStatements) {
            return rootRotationStatements(List.of(rootRotationStatements));
        }

        /**
         * @param rotationPeriod The amount of time in seconds Vault should wait before rotating the root credential.
         * A zero value tells Vault not to rotate the root credential. The minimum rotation period is 10 seconds. Requires Vault Enterprise 1.19+.
         * 
         * @return builder
         * 
         */
        public Builder rotationPeriod(@Nullable Output<Integer> rotationPeriod) {
            $.rotationPeriod = rotationPeriod;
            return this;
        }

        /**
         * @param rotationPeriod The amount of time in seconds Vault should wait before rotating the root credential.
         * A zero value tells Vault not to rotate the root credential. The minimum rotation period is 10 seconds. Requires Vault Enterprise 1.19+.
         * 
         * @return builder
         * 
         */
        public Builder rotationPeriod(Integer rotationPeriod) {
            return rotationPeriod(Output.of(rotationPeriod));
        }

        /**
         * @param rotationSchedule The schedule, in [cron-style time format](https://en.wikipedia.org/wiki/Cron),
         * defining the schedule on which Vault should rotate the root token. Requires Vault Enterprise 1.19+.
         * 
         * @return builder
         * 
         */
        public Builder rotationSchedule(@Nullable Output<String> rotationSchedule) {
            $.rotationSchedule = rotationSchedule;
            return this;
        }

        /**
         * @param rotationSchedule The schedule, in [cron-style time format](https://en.wikipedia.org/wiki/Cron),
         * defining the schedule on which Vault should rotate the root token. Requires Vault Enterprise 1.19+.
         * 
         * @return builder
         * 
         */
        public Builder rotationSchedule(String rotationSchedule) {
            return rotationSchedule(Output.of(rotationSchedule));
        }

        /**
         * @param rotationWindow The maximum amount of time in seconds allowed to complete
         * a rotation when a scheduled token rotation occurs. The default rotation window is
         * unbound and the minimum allowable window is `3600`. Requires Vault Enterprise 1.19+.
         * 
         * @return builder
         * 
         */
        public Builder rotationWindow(@Nullable Output<Integer> rotationWindow) {
            $.rotationWindow = rotationWindow;
            return this;
        }

        /**
         * @param rotationWindow The maximum amount of time in seconds allowed to complete
         * a rotation when a scheduled token rotation occurs. The default rotation window is
         * unbound and the minimum allowable window is `3600`. Requires Vault Enterprise 1.19+.
         * 
         * @return builder
         * 
         */
        public Builder rotationWindow(Integer rotationWindow) {
            return rotationWindow(Output.of(rotationWindow));
        }

        /**
         * @param serviceAccountJson A JSON encoded credential for use with IAM authorization
         * 
         * @return builder
         * 
         */
        public Builder serviceAccountJson(@Nullable Output<String> serviceAccountJson) {
            $.serviceAccountJson = serviceAccountJson;
            return this;
        }

        /**
         * @param serviceAccountJson A JSON encoded credential for use with IAM authorization
         * 
         * @return builder
         * 
         */
        public Builder serviceAccountJson(String serviceAccountJson) {
            return serviceAccountJson(Output.of(serviceAccountJson));
        }

        /**
         * @param tlsCa x509 CA file for validating the certificate presented by the MySQL server. Must be PEM encoded.
         * 
         * @return builder
         * 
         */
        public Builder tlsCa(@Nullable Output<String> tlsCa) {
            $.tlsCa = tlsCa;
            return this;
        }

        /**
         * @param tlsCa x509 CA file for validating the certificate presented by the MySQL server. Must be PEM encoded.
         * 
         * @return builder
         * 
         */
        public Builder tlsCa(String tlsCa) {
            return tlsCa(Output.of(tlsCa));
        }

        /**
         * @param tlsCertificateKey x509 certificate for connecting to the database. This must be a PEM encoded version of the private key and the certificate combined.
         * 
         * @return builder
         * 
         */
        public Builder tlsCertificateKey(@Nullable Output<String> tlsCertificateKey) {
            $.tlsCertificateKey = tlsCertificateKey;
            return this;
        }

        /**
         * @param tlsCertificateKey x509 certificate for connecting to the database. This must be a PEM encoded version of the private key and the certificate combined.
         * 
         * @return builder
         * 
         */
        public Builder tlsCertificateKey(String tlsCertificateKey) {
            return tlsCertificateKey(Output.of(tlsCertificateKey));
        }

        /**
         * @param username The root credential username used in the connection URL
         * 
         * @return builder
         * 
         */
        public Builder username(@Nullable Output<String> username) {
            $.username = username;
            return this;
        }

        /**
         * @param username The root credential username used in the connection URL
         * 
         * @return builder
         * 
         */
        public Builder username(String username) {
            return username(Output.of(username));
        }

        /**
         * @param usernameTemplate Username generation template.
         * 
         * @return builder
         * 
         */
        public Builder usernameTemplate(@Nullable Output<String> usernameTemplate) {
            $.usernameTemplate = usernameTemplate;
            return this;
        }

        /**
         * @param usernameTemplate Username generation template.
         * 
         * @return builder
         * 
         */
        public Builder usernameTemplate(String usernameTemplate) {
            return usernameTemplate(Output.of(usernameTemplate));
        }

        /**
         * @param verifyConnection Whether the connection should be verified on
         * initial configuration or not.
         * 
         * @return builder
         * 
         */
        public Builder verifyConnection(@Nullable Output<Boolean> verifyConnection) {
            $.verifyConnection = verifyConnection;
            return this;
        }

        /**
         * @param verifyConnection Whether the connection should be verified on
         * initial configuration or not.
         * 
         * @return builder
         * 
         */
        public Builder verifyConnection(Boolean verifyConnection) {
            return verifyConnection(Output.of(verifyConnection));
        }

        public SecretsMountMysqlLegacyArgs build() {
            if ($.name == null) {
                throw new MissingRequiredPropertyException("SecretsMountMysqlLegacyArgs", "name");
            }
            return $;
        }
    }

}
