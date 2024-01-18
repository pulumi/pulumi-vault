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


public final class SecretsMountMysqlRdArgs extends com.pulumi.resources.ResourceArgs {

    public static final SecretsMountMysqlRdArgs Empty = new SecretsMountMysqlRdArgs();

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

    @Import(name="authType")
    private @Nullable Output<String> authType;

    public Optional<Output<String>> authType() {
        return Optional.ofNullable(this.authType);
    }

    /**
     * Specifies the Redshift DSN.
     * See [Vault docs](https://www.vaultproject.io/api-docs/secret/databases/redshift#sample-payload)
     * 
     */
    @Import(name="connectionUrl")
    private @Nullable Output<String> connectionUrl;

    /**
     * @return Specifies the Redshift DSN.
     * See [Vault docs](https://www.vaultproject.io/api-docs/secret/databases/redshift#sample-payload)
     * 
     */
    public Optional<Output<String>> connectionUrl() {
        return Optional.ofNullable(this.connectionUrl);
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
     * The maximum amount of time a connection may be reused.
     * 
     */
    @Import(name="maxConnectionLifetime")
    private @Nullable Output<Integer> maxConnectionLifetime;

    /**
     * @return The maximum amount of time a connection may be reused.
     * 
     */
    public Optional<Output<Integer>> maxConnectionLifetime() {
        return Optional.ofNullable(this.maxConnectionLifetime);
    }

    /**
     * The maximum number of idle connections to
     * the database.
     * 
     */
    @Import(name="maxIdleConnections")
    private @Nullable Output<Integer> maxIdleConnections;

    /**
     * @return The maximum number of idle connections to
     * the database.
     * 
     */
    public Optional<Output<Integer>> maxIdleConnections() {
        return Optional.ofNullable(this.maxIdleConnections);
    }

    /**
     * The maximum number of open connections to
     * the database.
     * 
     */
    @Import(name="maxOpenConnections")
    private @Nullable Output<Integer> maxOpenConnections;

    /**
     * @return The maximum number of open connections to
     * the database.
     * 
     */
    public Optional<Output<Integer>> maxOpenConnections() {
        return Optional.ofNullable(this.maxOpenConnections);
    }

    @Import(name="name", required=true)
    private Output<String> name;

    public Output<String> name() {
        return this.name;
    }

    /**
     * The root credential password used in the connection URL.
     * 
     */
    @Import(name="password")
    private @Nullable Output<String> password;

    /**
     * @return The root credential password used in the connection URL.
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

    @Import(name="serviceAccountJson")
    private @Nullable Output<String> serviceAccountJson;

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
     * The root credential username used in the connection URL.
     * 
     */
    @Import(name="username")
    private @Nullable Output<String> username;

    /**
     * @return The root credential username used in the connection URL.
     * 
     */
    public Optional<Output<String>> username() {
        return Optional.ofNullable(this.username);
    }

    /**
     * [Template](https://www.vaultproject.io/docs/concepts/username-templating) describing how dynamic usernames are generated.
     * 
     */
    @Import(name="usernameTemplate")
    private @Nullable Output<String> usernameTemplate;

    /**
     * @return [Template](https://www.vaultproject.io/docs/concepts/username-templating) describing how dynamic usernames are generated.
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

    private SecretsMountMysqlRdArgs() {}

    private SecretsMountMysqlRdArgs(SecretsMountMysqlRdArgs $) {
        this.allowedRoles = $.allowedRoles;
        this.authType = $.authType;
        this.connectionUrl = $.connectionUrl;
        this.data = $.data;
        this.maxConnectionLifetime = $.maxConnectionLifetime;
        this.maxIdleConnections = $.maxIdleConnections;
        this.maxOpenConnections = $.maxOpenConnections;
        this.name = $.name;
        this.password = $.password;
        this.pluginName = $.pluginName;
        this.rootRotationStatements = $.rootRotationStatements;
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
    public static Builder builder(SecretsMountMysqlRdArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private SecretsMountMysqlRdArgs $;

        public Builder() {
            $ = new SecretsMountMysqlRdArgs();
        }

        public Builder(SecretsMountMysqlRdArgs defaults) {
            $ = new SecretsMountMysqlRdArgs(Objects.requireNonNull(defaults));
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

        public Builder authType(@Nullable Output<String> authType) {
            $.authType = authType;
            return this;
        }

        public Builder authType(String authType) {
            return authType(Output.of(authType));
        }

        /**
         * @param connectionUrl Specifies the Redshift DSN.
         * See [Vault docs](https://www.vaultproject.io/api-docs/secret/databases/redshift#sample-payload)
         * 
         * @return builder
         * 
         */
        public Builder connectionUrl(@Nullable Output<String> connectionUrl) {
            $.connectionUrl = connectionUrl;
            return this;
        }

        /**
         * @param connectionUrl Specifies the Redshift DSN.
         * See [Vault docs](https://www.vaultproject.io/api-docs/secret/databases/redshift#sample-payload)
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
         * @param maxConnectionLifetime The maximum amount of time a connection may be reused.
         * 
         * @return builder
         * 
         */
        public Builder maxConnectionLifetime(@Nullable Output<Integer> maxConnectionLifetime) {
            $.maxConnectionLifetime = maxConnectionLifetime;
            return this;
        }

        /**
         * @param maxConnectionLifetime The maximum amount of time a connection may be reused.
         * 
         * @return builder
         * 
         */
        public Builder maxConnectionLifetime(Integer maxConnectionLifetime) {
            return maxConnectionLifetime(Output.of(maxConnectionLifetime));
        }

        /**
         * @param maxIdleConnections The maximum number of idle connections to
         * the database.
         * 
         * @return builder
         * 
         */
        public Builder maxIdleConnections(@Nullable Output<Integer> maxIdleConnections) {
            $.maxIdleConnections = maxIdleConnections;
            return this;
        }

        /**
         * @param maxIdleConnections The maximum number of idle connections to
         * the database.
         * 
         * @return builder
         * 
         */
        public Builder maxIdleConnections(Integer maxIdleConnections) {
            return maxIdleConnections(Output.of(maxIdleConnections));
        }

        /**
         * @param maxOpenConnections The maximum number of open connections to
         * the database.
         * 
         * @return builder
         * 
         */
        public Builder maxOpenConnections(@Nullable Output<Integer> maxOpenConnections) {
            $.maxOpenConnections = maxOpenConnections;
            return this;
        }

        /**
         * @param maxOpenConnections The maximum number of open connections to
         * the database.
         * 
         * @return builder
         * 
         */
        public Builder maxOpenConnections(Integer maxOpenConnections) {
            return maxOpenConnections(Output.of(maxOpenConnections));
        }

        public Builder name(Output<String> name) {
            $.name = name;
            return this;
        }

        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param password The root credential password used in the connection URL.
         * 
         * @return builder
         * 
         */
        public Builder password(@Nullable Output<String> password) {
            $.password = password;
            return this;
        }

        /**
         * @param password The root credential password used in the connection URL.
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

        public Builder serviceAccountJson(@Nullable Output<String> serviceAccountJson) {
            $.serviceAccountJson = serviceAccountJson;
            return this;
        }

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
         * @param username The root credential username used in the connection URL.
         * 
         * @return builder
         * 
         */
        public Builder username(@Nullable Output<String> username) {
            $.username = username;
            return this;
        }

        /**
         * @param username The root credential username used in the connection URL.
         * 
         * @return builder
         * 
         */
        public Builder username(String username) {
            return username(Output.of(username));
        }

        /**
         * @param usernameTemplate [Template](https://www.vaultproject.io/docs/concepts/username-templating) describing how dynamic usernames are generated.
         * 
         * @return builder
         * 
         */
        public Builder usernameTemplate(@Nullable Output<String> usernameTemplate) {
            $.usernameTemplate = usernameTemplate;
            return this;
        }

        /**
         * @param usernameTemplate [Template](https://www.vaultproject.io/docs/concepts/username-templating) describing how dynamic usernames are generated.
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

        public SecretsMountMysqlRdArgs build() {
            if ($.name == null) {
                throw new MissingRequiredPropertyException("SecretsMountMysqlRdArgs", "name");
            }
            return $;
        }
    }

}
