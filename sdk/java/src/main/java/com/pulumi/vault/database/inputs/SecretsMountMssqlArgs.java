// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vault.database.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
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


public final class SecretsMountMssqlArgs extends com.pulumi.resources.ResourceArgs {

    public static final SecretsMountMssqlArgs Empty = new SecretsMountMssqlArgs();

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
     * Set to true when the target is a Contained Database, e.g. AzureSQL.
     * 
     */
    @Import(name="containedDb")
    private @Nullable Output<Boolean> containedDb;

    /**
     * @return Set to true when the target is a Contained Database, e.g. AzureSQL.
     * 
     */
    public Optional<Output<Boolean>> containedDb() {
        return Optional.ofNullable(this.containedDb);
    }

    /**
     * A map of sensitive data to pass to the endpoint. Useful for templated connection strings.
     * 
     * Supported list of database secrets engines that can be configured:
     * 
     */
    @Import(name="data")
    private @Nullable Output<Map<String,Object>> data;

    /**
     * @return A map of sensitive data to pass to the endpoint. Useful for templated connection strings.
     * 
     * Supported list of database secrets engines that can be configured:
     * 
     */
    public Optional<Output<Map<String,Object>>> data() {
        return Optional.ofNullable(this.data);
    }

    /**
     * Disable special character escaping in username and password
     * 
     */
    @Import(name="disableEscaping")
    private @Nullable Output<Boolean> disableEscaping;

    /**
     * @return Disable special character escaping in username and password
     * 
     */
    public Optional<Output<Boolean>> disableEscaping() {
        return Optional.ofNullable(this.disableEscaping);
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

    private SecretsMountMssqlArgs() {}

    private SecretsMountMssqlArgs(SecretsMountMssqlArgs $) {
        this.allowedRoles = $.allowedRoles;
        this.connectionUrl = $.connectionUrl;
        this.containedDb = $.containedDb;
        this.data = $.data;
        this.disableEscaping = $.disableEscaping;
        this.maxConnectionLifetime = $.maxConnectionLifetime;
        this.maxIdleConnections = $.maxIdleConnections;
        this.maxOpenConnections = $.maxOpenConnections;
        this.name = $.name;
        this.password = $.password;
        this.pluginName = $.pluginName;
        this.rootRotationStatements = $.rootRotationStatements;
        this.username = $.username;
        this.usernameTemplate = $.usernameTemplate;
        this.verifyConnection = $.verifyConnection;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(SecretsMountMssqlArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private SecretsMountMssqlArgs $;

        public Builder() {
            $ = new SecretsMountMssqlArgs();
        }

        public Builder(SecretsMountMssqlArgs defaults) {
            $ = new SecretsMountMssqlArgs(Objects.requireNonNull(defaults));
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
         * @param containedDb Set to true when the target is a Contained Database, e.g. AzureSQL.
         * 
         * @return builder
         * 
         */
        public Builder containedDb(@Nullable Output<Boolean> containedDb) {
            $.containedDb = containedDb;
            return this;
        }

        /**
         * @param containedDb Set to true when the target is a Contained Database, e.g. AzureSQL.
         * 
         * @return builder
         * 
         */
        public Builder containedDb(Boolean containedDb) {
            return containedDb(Output.of(containedDb));
        }

        /**
         * @param data A map of sensitive data to pass to the endpoint. Useful for templated connection strings.
         * 
         * Supported list of database secrets engines that can be configured:
         * 
         * @return builder
         * 
         */
        public Builder data(@Nullable Output<Map<String,Object>> data) {
            $.data = data;
            return this;
        }

        /**
         * @param data A map of sensitive data to pass to the endpoint. Useful for templated connection strings.
         * 
         * Supported list of database secrets engines that can be configured:
         * 
         * @return builder
         * 
         */
        public Builder data(Map<String,Object> data) {
            return data(Output.of(data));
        }

        /**
         * @param disableEscaping Disable special character escaping in username and password
         * 
         * @return builder
         * 
         */
        public Builder disableEscaping(@Nullable Output<Boolean> disableEscaping) {
            $.disableEscaping = disableEscaping;
            return this;
        }

        /**
         * @param disableEscaping Disable special character escaping in username and password
         * 
         * @return builder
         * 
         */
        public Builder disableEscaping(Boolean disableEscaping) {
            return disableEscaping(Output.of(disableEscaping));
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

        public SecretsMountMssqlArgs build() {
            if ($.name == null) {
                throw new MissingRequiredPropertyException("SecretsMountMssqlArgs", "name");
            }
            return $;
        }
    }

}
