// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vault.database.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class SecretBackendConnectionInfluxdbArgs extends com.pulumi.resources.ResourceArgs {

    public static final SecretBackendConnectionInfluxdbArgs Empty = new SecretBackendConnectionInfluxdbArgs();

    /**
     * The number of seconds to use as a connection
     * timeout.
     * 
     */
    @Import(name="connectTimeout")
    private @Nullable Output<Integer> connectTimeout;

    /**
     * @return The number of seconds to use as a connection
     * timeout.
     * 
     */
    public Optional<Output<Integer>> connectTimeout() {
        return Optional.ofNullable(this.connectTimeout);
    }

    /**
     * The host to connect to.
     * 
     */
    @Import(name="host", required=true)
    private Output<String> host;

    /**
     * @return The host to connect to.
     * 
     */
    public Output<String> host() {
        return this.host;
    }

    /**
     * Whether to skip verification of the server
     * certificate when using TLS.
     * 
     */
    @Import(name="insecureTls")
    private @Nullable Output<Boolean> insecureTls;

    /**
     * @return Whether to skip verification of the server
     * certificate when using TLS.
     * 
     */
    public Optional<Output<Boolean>> insecureTls() {
        return Optional.ofNullable(this.insecureTls);
    }

    /**
     * The password to authenticate with.
     * 
     */
    @Import(name="password", required=true)
    private Output<String> password;

    /**
     * @return The password to authenticate with.
     * 
     */
    public Output<String> password() {
        return this.password;
    }

    /**
     * Concatenated PEM blocks configuring the certificate
     * chain.
     * 
     */
    @Import(name="pemBundle")
    private @Nullable Output<String> pemBundle;

    /**
     * @return Concatenated PEM blocks configuring the certificate
     * chain.
     * 
     */
    public Optional<Output<String>> pemBundle() {
        return Optional.ofNullable(this.pemBundle);
    }

    /**
     * A JSON structure configuring the certificate chain.
     * 
     */
    @Import(name="pemJson")
    private @Nullable Output<String> pemJson;

    /**
     * @return A JSON structure configuring the certificate chain.
     * 
     */
    public Optional<Output<String>> pemJson() {
        return Optional.ofNullable(this.pemJson);
    }

    /**
     * The default port to connect to if no port is specified as
     * part of the host.
     * 
     */
    @Import(name="port")
    private @Nullable Output<Integer> port;

    /**
     * @return The default port to connect to if no port is specified as
     * part of the host.
     * 
     */
    public Optional<Output<Integer>> port() {
        return Optional.ofNullable(this.port);
    }

    /**
     * Whether to use TLS when connecting to Cassandra.
     * 
     */
    @Import(name="tls")
    private @Nullable Output<Boolean> tls;

    /**
     * @return Whether to use TLS when connecting to Cassandra.
     * 
     */
    public Optional<Output<Boolean>> tls() {
        return Optional.ofNullable(this.tls);
    }

    /**
     * The username to authenticate with.
     * 
     */
    @Import(name="username", required=true)
    private Output<String> username;

    /**
     * @return The username to authenticate with.
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

    private SecretBackendConnectionInfluxdbArgs() {}

    private SecretBackendConnectionInfluxdbArgs(SecretBackendConnectionInfluxdbArgs $) {
        this.connectTimeout = $.connectTimeout;
        this.host = $.host;
        this.insecureTls = $.insecureTls;
        this.password = $.password;
        this.pemBundle = $.pemBundle;
        this.pemJson = $.pemJson;
        this.port = $.port;
        this.tls = $.tls;
        this.username = $.username;
        this.usernameTemplate = $.usernameTemplate;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(SecretBackendConnectionInfluxdbArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private SecretBackendConnectionInfluxdbArgs $;

        public Builder() {
            $ = new SecretBackendConnectionInfluxdbArgs();
        }

        public Builder(SecretBackendConnectionInfluxdbArgs defaults) {
            $ = new SecretBackendConnectionInfluxdbArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param connectTimeout The number of seconds to use as a connection
         * timeout.
         * 
         * @return builder
         * 
         */
        public Builder connectTimeout(@Nullable Output<Integer> connectTimeout) {
            $.connectTimeout = connectTimeout;
            return this;
        }

        /**
         * @param connectTimeout The number of seconds to use as a connection
         * timeout.
         * 
         * @return builder
         * 
         */
        public Builder connectTimeout(Integer connectTimeout) {
            return connectTimeout(Output.of(connectTimeout));
        }

        /**
         * @param host The host to connect to.
         * 
         * @return builder
         * 
         */
        public Builder host(Output<String> host) {
            $.host = host;
            return this;
        }

        /**
         * @param host The host to connect to.
         * 
         * @return builder
         * 
         */
        public Builder host(String host) {
            return host(Output.of(host));
        }

        /**
         * @param insecureTls Whether to skip verification of the server
         * certificate when using TLS.
         * 
         * @return builder
         * 
         */
        public Builder insecureTls(@Nullable Output<Boolean> insecureTls) {
            $.insecureTls = insecureTls;
            return this;
        }

        /**
         * @param insecureTls Whether to skip verification of the server
         * certificate when using TLS.
         * 
         * @return builder
         * 
         */
        public Builder insecureTls(Boolean insecureTls) {
            return insecureTls(Output.of(insecureTls));
        }

        /**
         * @param password The password to authenticate with.
         * 
         * @return builder
         * 
         */
        public Builder password(Output<String> password) {
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
         * @param pemBundle Concatenated PEM blocks configuring the certificate
         * chain.
         * 
         * @return builder
         * 
         */
        public Builder pemBundle(@Nullable Output<String> pemBundle) {
            $.pemBundle = pemBundle;
            return this;
        }

        /**
         * @param pemBundle Concatenated PEM blocks configuring the certificate
         * chain.
         * 
         * @return builder
         * 
         */
        public Builder pemBundle(String pemBundle) {
            return pemBundle(Output.of(pemBundle));
        }

        /**
         * @param pemJson A JSON structure configuring the certificate chain.
         * 
         * @return builder
         * 
         */
        public Builder pemJson(@Nullable Output<String> pemJson) {
            $.pemJson = pemJson;
            return this;
        }

        /**
         * @param pemJson A JSON structure configuring the certificate chain.
         * 
         * @return builder
         * 
         */
        public Builder pemJson(String pemJson) {
            return pemJson(Output.of(pemJson));
        }

        /**
         * @param port The default port to connect to if no port is specified as
         * part of the host.
         * 
         * @return builder
         * 
         */
        public Builder port(@Nullable Output<Integer> port) {
            $.port = port;
            return this;
        }

        /**
         * @param port The default port to connect to if no port is specified as
         * part of the host.
         * 
         * @return builder
         * 
         */
        public Builder port(Integer port) {
            return port(Output.of(port));
        }

        /**
         * @param tls Whether to use TLS when connecting to Cassandra.
         * 
         * @return builder
         * 
         */
        public Builder tls(@Nullable Output<Boolean> tls) {
            $.tls = tls;
            return this;
        }

        /**
         * @param tls Whether to use TLS when connecting to Cassandra.
         * 
         * @return builder
         * 
         */
        public Builder tls(Boolean tls) {
            return tls(Output.of(tls));
        }

        /**
         * @param username The username to authenticate with.
         * 
         * @return builder
         * 
         */
        public Builder username(Output<String> username) {
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

        public SecretBackendConnectionInfluxdbArgs build() {
            if ($.host == null) {
                throw new MissingRequiredPropertyException("SecretBackendConnectionInfluxdbArgs", "host");
            }
            if ($.password == null) {
                throw new MissingRequiredPropertyException("SecretBackendConnectionInfluxdbArgs", "password");
            }
            if ($.username == null) {
                throw new MissingRequiredPropertyException("SecretBackendConnectionInfluxdbArgs", "username");
            }
            return $;
        }
    }

}
