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


public final class SecretsMountElasticsearchArgs extends com.pulumi.resources.ResourceArgs {

    public static final SecretsMountElasticsearchArgs Empty = new SecretsMountElasticsearchArgs();

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
     * The path to a PEM-encoded CA cert file to use to verify the Elasticsearch server&#39;s identity
     * 
     */
    @Import(name="caCert")
    private @Nullable Output<String> caCert;

    /**
     * @return The path to a PEM-encoded CA cert file to use to verify the Elasticsearch server&#39;s identity
     * 
     */
    public Optional<Output<String>> caCert() {
        return Optional.ofNullable(this.caCert);
    }

    /**
     * The path to a directory of PEM-encoded CA cert files to use to verify the Elasticsearch server&#39;s identity
     * 
     */
    @Import(name="caPath")
    private @Nullable Output<String> caPath;

    /**
     * @return The path to a directory of PEM-encoded CA cert files to use to verify the Elasticsearch server&#39;s identity
     * 
     */
    public Optional<Output<String>> caPath() {
        return Optional.ofNullable(this.caPath);
    }

    /**
     * The path to the certificate for the Elasticsearch client to present for communication
     * 
     */
    @Import(name="clientCert")
    private @Nullable Output<String> clientCert;

    /**
     * @return The path to the certificate for the Elasticsearch client to present for communication
     * 
     */
    public Optional<Output<String>> clientCert() {
        return Optional.ofNullable(this.clientCert);
    }

    /**
     * The path to the key for the Elasticsearch client to use for communication
     * 
     */
    @Import(name="clientKey")
    private @Nullable Output<String> clientKey;

    /**
     * @return The path to the key for the Elasticsearch client to use for communication
     * 
     */
    public Optional<Output<String>> clientKey() {
        return Optional.ofNullable(this.clientKey);
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
     * Whether to disable certificate verification
     * 
     */
    @Import(name="insecure")
    private @Nullable Output<Boolean> insecure;

    /**
     * @return Whether to disable certificate verification
     * 
     */
    public Optional<Output<Boolean>> insecure() {
        return Optional.ofNullable(this.insecure);
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
     * The password to be used in the connection URL
     * 
     */
    @Import(name="password", required=true)
    private Output<String> password;

    /**
     * @return The password to be used in the connection URL
     * 
     */
    public Output<String> password() {
        return this.password;
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
     * This, if set, is used to set the SNI host when connecting via TLS
     * 
     */
    @Import(name="tlsServerName")
    private @Nullable Output<String> tlsServerName;

    /**
     * @return This, if set, is used to set the SNI host when connecting via TLS
     * 
     */
    public Optional<Output<String>> tlsServerName() {
        return Optional.ofNullable(this.tlsServerName);
    }

    /**
     * The URL for Elasticsearch&#39;s API
     * 
     */
    @Import(name="url", required=true)
    private Output<String> url;

    /**
     * @return The URL for Elasticsearch&#39;s API
     * 
     */
    public Output<String> url() {
        return this.url;
    }

    /**
     * The username to be used in the connection URL
     * 
     */
    @Import(name="username", required=true)
    private Output<String> username;

    /**
     * @return The username to be used in the connection URL
     * 
     */
    public Output<String> username() {
        return this.username;
    }

    /**
     * Template describing how dynamic usernames are generated.
     * 
     */
    @Import(name="usernameTemplate")
    private @Nullable Output<String> usernameTemplate;

    /**
     * @return Template describing how dynamic usernames are generated.
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

    private SecretsMountElasticsearchArgs() {}

    private SecretsMountElasticsearchArgs(SecretsMountElasticsearchArgs $) {
        this.allowedRoles = $.allowedRoles;
        this.caCert = $.caCert;
        this.caPath = $.caPath;
        this.clientCert = $.clientCert;
        this.clientKey = $.clientKey;
        this.data = $.data;
        this.disableAutomatedRotation = $.disableAutomatedRotation;
        this.insecure = $.insecure;
        this.name = $.name;
        this.password = $.password;
        this.pluginName = $.pluginName;
        this.rootRotationStatements = $.rootRotationStatements;
        this.rotationPeriod = $.rotationPeriod;
        this.rotationSchedule = $.rotationSchedule;
        this.rotationWindow = $.rotationWindow;
        this.tlsServerName = $.tlsServerName;
        this.url = $.url;
        this.username = $.username;
        this.usernameTemplate = $.usernameTemplate;
        this.verifyConnection = $.verifyConnection;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(SecretsMountElasticsearchArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private SecretsMountElasticsearchArgs $;

        public Builder() {
            $ = new SecretsMountElasticsearchArgs();
        }

        public Builder(SecretsMountElasticsearchArgs defaults) {
            $ = new SecretsMountElasticsearchArgs(Objects.requireNonNull(defaults));
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
         * @param caCert The path to a PEM-encoded CA cert file to use to verify the Elasticsearch server&#39;s identity
         * 
         * @return builder
         * 
         */
        public Builder caCert(@Nullable Output<String> caCert) {
            $.caCert = caCert;
            return this;
        }

        /**
         * @param caCert The path to a PEM-encoded CA cert file to use to verify the Elasticsearch server&#39;s identity
         * 
         * @return builder
         * 
         */
        public Builder caCert(String caCert) {
            return caCert(Output.of(caCert));
        }

        /**
         * @param caPath The path to a directory of PEM-encoded CA cert files to use to verify the Elasticsearch server&#39;s identity
         * 
         * @return builder
         * 
         */
        public Builder caPath(@Nullable Output<String> caPath) {
            $.caPath = caPath;
            return this;
        }

        /**
         * @param caPath The path to a directory of PEM-encoded CA cert files to use to verify the Elasticsearch server&#39;s identity
         * 
         * @return builder
         * 
         */
        public Builder caPath(String caPath) {
            return caPath(Output.of(caPath));
        }

        /**
         * @param clientCert The path to the certificate for the Elasticsearch client to present for communication
         * 
         * @return builder
         * 
         */
        public Builder clientCert(@Nullable Output<String> clientCert) {
            $.clientCert = clientCert;
            return this;
        }

        /**
         * @param clientCert The path to the certificate for the Elasticsearch client to present for communication
         * 
         * @return builder
         * 
         */
        public Builder clientCert(String clientCert) {
            return clientCert(Output.of(clientCert));
        }

        /**
         * @param clientKey The path to the key for the Elasticsearch client to use for communication
         * 
         * @return builder
         * 
         */
        public Builder clientKey(@Nullable Output<String> clientKey) {
            $.clientKey = clientKey;
            return this;
        }

        /**
         * @param clientKey The path to the key for the Elasticsearch client to use for communication
         * 
         * @return builder
         * 
         */
        public Builder clientKey(String clientKey) {
            return clientKey(Output.of(clientKey));
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
         * @param insecure Whether to disable certificate verification
         * 
         * @return builder
         * 
         */
        public Builder insecure(@Nullable Output<Boolean> insecure) {
            $.insecure = insecure;
            return this;
        }

        /**
         * @param insecure Whether to disable certificate verification
         * 
         * @return builder
         * 
         */
        public Builder insecure(Boolean insecure) {
            return insecure(Output.of(insecure));
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
         * @param password The password to be used in the connection URL
         * 
         * @return builder
         * 
         */
        public Builder password(Output<String> password) {
            $.password = password;
            return this;
        }

        /**
         * @param password The password to be used in the connection URL
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
         * @param tlsServerName This, if set, is used to set the SNI host when connecting via TLS
         * 
         * @return builder
         * 
         */
        public Builder tlsServerName(@Nullable Output<String> tlsServerName) {
            $.tlsServerName = tlsServerName;
            return this;
        }

        /**
         * @param tlsServerName This, if set, is used to set the SNI host when connecting via TLS
         * 
         * @return builder
         * 
         */
        public Builder tlsServerName(String tlsServerName) {
            return tlsServerName(Output.of(tlsServerName));
        }

        /**
         * @param url The URL for Elasticsearch&#39;s API
         * 
         * @return builder
         * 
         */
        public Builder url(Output<String> url) {
            $.url = url;
            return this;
        }

        /**
         * @param url The URL for Elasticsearch&#39;s API
         * 
         * @return builder
         * 
         */
        public Builder url(String url) {
            return url(Output.of(url));
        }

        /**
         * @param username The username to be used in the connection URL
         * 
         * @return builder
         * 
         */
        public Builder username(Output<String> username) {
            $.username = username;
            return this;
        }

        /**
         * @param username The username to be used in the connection URL
         * 
         * @return builder
         * 
         */
        public Builder username(String username) {
            return username(Output.of(username));
        }

        /**
         * @param usernameTemplate Template describing how dynamic usernames are generated.
         * 
         * @return builder
         * 
         */
        public Builder usernameTemplate(@Nullable Output<String> usernameTemplate) {
            $.usernameTemplate = usernameTemplate;
            return this;
        }

        /**
         * @param usernameTemplate Template describing how dynamic usernames are generated.
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

        public SecretsMountElasticsearchArgs build() {
            if ($.name == null) {
                throw new MissingRequiredPropertyException("SecretsMountElasticsearchArgs", "name");
            }
            if ($.password == null) {
                throw new MissingRequiredPropertyException("SecretsMountElasticsearchArgs", "password");
            }
            if ($.url == null) {
                throw new MissingRequiredPropertyException("SecretsMountElasticsearchArgs", "url");
            }
            if ($.username == null) {
                throw new MissingRequiredPropertyException("SecretsMountElasticsearchArgs", "username");
            }
            return $;
        }
    }

}
