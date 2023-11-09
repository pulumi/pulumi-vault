// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vault.database.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class SecretBackendConnectionMysqlLegacyArgs extends com.pulumi.resources.ResourceArgs {

    public static final SecretBackendConnectionMysqlLegacyArgs Empty = new SecretBackendConnectionMysqlLegacyArgs();

    /**
     * A URL containing connection information. See
     * the [Vault
     * docs](https://www.vaultproject.io/api-docs/secret/databases/mongodb.html#sample-payload)
     * for an example.
     * 
     */
    @Import(name="connectionUrl")
    private @Nullable Output<String> connectionUrl;

    /**
     * @return A URL containing connection information. See
     * the [Vault
     * docs](https://www.vaultproject.io/api-docs/secret/databases/mongodb.html#sample-payload)
     * for an example.
     * 
     */
    public Optional<Output<String>> connectionUrl() {
        return Optional.ofNullable(this.connectionUrl);
    }

    /**
     * The maximum number of seconds to keep
     * a connection alive for.
     * 
     */
    @Import(name="maxConnectionLifetime")
    private @Nullable Output<Integer> maxConnectionLifetime;

    /**
     * @return The maximum number of seconds to keep
     * a connection alive for.
     * 
     */
    public Optional<Output<Integer>> maxConnectionLifetime() {
        return Optional.ofNullable(this.maxConnectionLifetime);
    }

    /**
     * The maximum number of idle connections to
     * maintain.
     * 
     */
    @Import(name="maxIdleConnections")
    private @Nullable Output<Integer> maxIdleConnections;

    /**
     * @return The maximum number of idle connections to
     * maintain.
     * 
     */
    public Optional<Output<Integer>> maxIdleConnections() {
        return Optional.ofNullable(this.maxIdleConnections);
    }

    /**
     * The maximum number of open connections to
     * use.
     * 
     */
    @Import(name="maxOpenConnections")
    private @Nullable Output<Integer> maxOpenConnections;

    /**
     * @return The maximum number of open connections to
     * use.
     * 
     */
    public Optional<Output<Integer>> maxOpenConnections() {
        return Optional.ofNullable(this.maxOpenConnections);
    }

    /**
     * The password to authenticate with.
     * 
     */
    @Import(name="password")
    private @Nullable Output<String> password;

    /**
     * @return The password to authenticate with.
     * 
     */
    public Optional<Output<String>> password() {
        return Optional.ofNullable(this.password);
    }

    /**
     * The username to authenticate with.
     * 
     */
    @Import(name="username")
    private @Nullable Output<String> username;

    /**
     * @return The username to authenticate with.
     * 
     */
    public Optional<Output<String>> username() {
        return Optional.ofNullable(this.username);
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

    private SecretBackendConnectionMysqlLegacyArgs() {}

    private SecretBackendConnectionMysqlLegacyArgs(SecretBackendConnectionMysqlLegacyArgs $) {
        this.connectionUrl = $.connectionUrl;
        this.maxConnectionLifetime = $.maxConnectionLifetime;
        this.maxIdleConnections = $.maxIdleConnections;
        this.maxOpenConnections = $.maxOpenConnections;
        this.password = $.password;
        this.username = $.username;
        this.usernameTemplate = $.usernameTemplate;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(SecretBackendConnectionMysqlLegacyArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private SecretBackendConnectionMysqlLegacyArgs $;

        public Builder() {
            $ = new SecretBackendConnectionMysqlLegacyArgs();
        }

        public Builder(SecretBackendConnectionMysqlLegacyArgs defaults) {
            $ = new SecretBackendConnectionMysqlLegacyArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param connectionUrl A URL containing connection information. See
         * the [Vault
         * docs](https://www.vaultproject.io/api-docs/secret/databases/mongodb.html#sample-payload)
         * for an example.
         * 
         * @return builder
         * 
         */
        public Builder connectionUrl(@Nullable Output<String> connectionUrl) {
            $.connectionUrl = connectionUrl;
            return this;
        }

        /**
         * @param connectionUrl A URL containing connection information. See
         * the [Vault
         * docs](https://www.vaultproject.io/api-docs/secret/databases/mongodb.html#sample-payload)
         * for an example.
         * 
         * @return builder
         * 
         */
        public Builder connectionUrl(String connectionUrl) {
            return connectionUrl(Output.of(connectionUrl));
        }

        /**
         * @param maxConnectionLifetime The maximum number of seconds to keep
         * a connection alive for.
         * 
         * @return builder
         * 
         */
        public Builder maxConnectionLifetime(@Nullable Output<Integer> maxConnectionLifetime) {
            $.maxConnectionLifetime = maxConnectionLifetime;
            return this;
        }

        /**
         * @param maxConnectionLifetime The maximum number of seconds to keep
         * a connection alive for.
         * 
         * @return builder
         * 
         */
        public Builder maxConnectionLifetime(Integer maxConnectionLifetime) {
            return maxConnectionLifetime(Output.of(maxConnectionLifetime));
        }

        /**
         * @param maxIdleConnections The maximum number of idle connections to
         * maintain.
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
         * maintain.
         * 
         * @return builder
         * 
         */
        public Builder maxIdleConnections(Integer maxIdleConnections) {
            return maxIdleConnections(Output.of(maxIdleConnections));
        }

        /**
         * @param maxOpenConnections The maximum number of open connections to
         * use.
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
         * use.
         * 
         * @return builder
         * 
         */
        public Builder maxOpenConnections(Integer maxOpenConnections) {
            return maxOpenConnections(Output.of(maxOpenConnections));
        }

        /**
         * @param password The password to authenticate with.
         * 
         * @return builder
         * 
         */
        public Builder password(@Nullable Output<String> password) {
            $.password = password;
            return this;
        }

        /**
         * @param password The password to authenticate with.
         * 
         * @return builder
         * 
         */
        public Builder password(String password) {
            return password(Output.of(password));
        }

        /**
         * @param username The username to authenticate with.
         * 
         * @return builder
         * 
         */
        public Builder username(@Nullable Output<String> username) {
            $.username = username;
            return this;
        }

        /**
         * @param username The username to authenticate with.
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

        public SecretBackendConnectionMysqlLegacyArgs build() {
            return $;
        }
    }

}
