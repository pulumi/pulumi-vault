// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vault.database.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class SecretBackendConnectionCassandraArgs extends com.pulumi.resources.ResourceArgs {

    public static final SecretBackendConnectionCassandraArgs Empty = new SecretBackendConnectionCassandraArgs();

    /**
     * The number of seconds to use as a connection timeout.
     * 
     */
    @Import(name="connectTimeout")
    private @Nullable Output<Integer> connectTimeout;

    /**
     * @return The number of seconds to use as a connection timeout.
     * 
     */
    public Optional<Output<Integer>> connectTimeout() {
        return Optional.ofNullable(this.connectTimeout);
    }

    /**
     * Cassandra hosts to connect to.
     * 
     */
    @Import(name="hosts")
    private @Nullable Output<List<String>> hosts;

    /**
     * @return Cassandra hosts to connect to.
     * 
     */
    public Optional<Output<List<String>>> hosts() {
        return Optional.ofNullable(this.hosts);
    }

    /**
     * Whether to skip verification of the server certificate when using TLS.
     * 
     */
    @Import(name="insecureTls")
    private @Nullable Output<Boolean> insecureTls;

    /**
     * @return Whether to skip verification of the server certificate when using TLS.
     * 
     */
    public Optional<Output<Boolean>> insecureTls() {
        return Optional.ofNullable(this.insecureTls);
    }

    /**
     * The password to use when authenticating with Cassandra.
     * 
     */
    @Import(name="password")
    private @Nullable Output<String> password;

    /**
     * @return The password to use when authenticating with Cassandra.
     * 
     */
    public Optional<Output<String>> password() {
        return Optional.ofNullable(this.password);
    }

    /**
     * Concatenated PEM blocks containing a certificate and private key; a certificate, private key, and issuing CA certificate; or just a CA certificate.
     * 
     */
    @Import(name="pemBundle")
    private @Nullable Output<String> pemBundle;

    /**
     * @return Concatenated PEM blocks containing a certificate and private key; a certificate, private key, and issuing CA certificate; or just a CA certificate.
     * 
     */
    public Optional<Output<String>> pemBundle() {
        return Optional.ofNullable(this.pemBundle);
    }

    /**
     * Specifies JSON containing a certificate and private key; a certificate, private key, and issuing CA certificate; or just a CA certificate.
     * 
     */
    @Import(name="pemJson")
    private @Nullable Output<String> pemJson;

    /**
     * @return Specifies JSON containing a certificate and private key; a certificate, private key, and issuing CA certificate; or just a CA certificate.
     * 
     */
    public Optional<Output<String>> pemJson() {
        return Optional.ofNullable(this.pemJson);
    }

    /**
     * The transport port to use to connect to Cassandra.
     * 
     */
    @Import(name="port")
    private @Nullable Output<Integer> port;

    /**
     * @return The transport port to use to connect to Cassandra.
     * 
     */
    public Optional<Output<Integer>> port() {
        return Optional.ofNullable(this.port);
    }

    /**
     * The CQL protocol version to use.
     * 
     */
    @Import(name="protocolVersion")
    private @Nullable Output<Integer> protocolVersion;

    /**
     * @return The CQL protocol version to use.
     * 
     */
    public Optional<Output<Integer>> protocolVersion() {
        return Optional.ofNullable(this.protocolVersion);
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
     * The username to use when authenticating with Cassandra.
     * 
     */
    @Import(name="username")
    private @Nullable Output<String> username;

    /**
     * @return The username to use when authenticating with Cassandra.
     * 
     */
    public Optional<Output<String>> username() {
        return Optional.ofNullable(this.username);
    }

    private SecretBackendConnectionCassandraArgs() {}

    private SecretBackendConnectionCassandraArgs(SecretBackendConnectionCassandraArgs $) {
        this.connectTimeout = $.connectTimeout;
        this.hosts = $.hosts;
        this.insecureTls = $.insecureTls;
        this.password = $.password;
        this.pemBundle = $.pemBundle;
        this.pemJson = $.pemJson;
        this.port = $.port;
        this.protocolVersion = $.protocolVersion;
        this.tls = $.tls;
        this.username = $.username;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(SecretBackendConnectionCassandraArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private SecretBackendConnectionCassandraArgs $;

        public Builder() {
            $ = new SecretBackendConnectionCassandraArgs();
        }

        public Builder(SecretBackendConnectionCassandraArgs defaults) {
            $ = new SecretBackendConnectionCassandraArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param connectTimeout The number of seconds to use as a connection timeout.
         * 
         * @return builder
         * 
         */
        public Builder connectTimeout(@Nullable Output<Integer> connectTimeout) {
            $.connectTimeout = connectTimeout;
            return this;
        }

        /**
         * @param connectTimeout The number of seconds to use as a connection timeout.
         * 
         * @return builder
         * 
         */
        public Builder connectTimeout(Integer connectTimeout) {
            return connectTimeout(Output.of(connectTimeout));
        }

        /**
         * @param hosts Cassandra hosts to connect to.
         * 
         * @return builder
         * 
         */
        public Builder hosts(@Nullable Output<List<String>> hosts) {
            $.hosts = hosts;
            return this;
        }

        /**
         * @param hosts Cassandra hosts to connect to.
         * 
         * @return builder
         * 
         */
        public Builder hosts(List<String> hosts) {
            return hosts(Output.of(hosts));
        }

        /**
         * @param hosts Cassandra hosts to connect to.
         * 
         * @return builder
         * 
         */
        public Builder hosts(String... hosts) {
            return hosts(List.of(hosts));
        }

        /**
         * @param insecureTls Whether to skip verification of the server certificate when using TLS.
         * 
         * @return builder
         * 
         */
        public Builder insecureTls(@Nullable Output<Boolean> insecureTls) {
            $.insecureTls = insecureTls;
            return this;
        }

        /**
         * @param insecureTls Whether to skip verification of the server certificate when using TLS.
         * 
         * @return builder
         * 
         */
        public Builder insecureTls(Boolean insecureTls) {
            return insecureTls(Output.of(insecureTls));
        }

        /**
         * @param password The password to use when authenticating with Cassandra.
         * 
         * @return builder
         * 
         */
        public Builder password(@Nullable Output<String> password) {
            $.password = password;
            return this;
        }

        /**
         * @param password The password to use when authenticating with Cassandra.
         * 
         * @return builder
         * 
         */
        public Builder password(String password) {
            return password(Output.of(password));
        }

        /**
         * @param pemBundle Concatenated PEM blocks containing a certificate and private key; a certificate, private key, and issuing CA certificate; or just a CA certificate.
         * 
         * @return builder
         * 
         */
        public Builder pemBundle(@Nullable Output<String> pemBundle) {
            $.pemBundle = pemBundle;
            return this;
        }

        /**
         * @param pemBundle Concatenated PEM blocks containing a certificate and private key; a certificate, private key, and issuing CA certificate; or just a CA certificate.
         * 
         * @return builder
         * 
         */
        public Builder pemBundle(String pemBundle) {
            return pemBundle(Output.of(pemBundle));
        }

        /**
         * @param pemJson Specifies JSON containing a certificate and private key; a certificate, private key, and issuing CA certificate; or just a CA certificate.
         * 
         * @return builder
         * 
         */
        public Builder pemJson(@Nullable Output<String> pemJson) {
            $.pemJson = pemJson;
            return this;
        }

        /**
         * @param pemJson Specifies JSON containing a certificate and private key; a certificate, private key, and issuing CA certificate; or just a CA certificate.
         * 
         * @return builder
         * 
         */
        public Builder pemJson(String pemJson) {
            return pemJson(Output.of(pemJson));
        }

        /**
         * @param port The transport port to use to connect to Cassandra.
         * 
         * @return builder
         * 
         */
        public Builder port(@Nullable Output<Integer> port) {
            $.port = port;
            return this;
        }

        /**
         * @param port The transport port to use to connect to Cassandra.
         * 
         * @return builder
         * 
         */
        public Builder port(Integer port) {
            return port(Output.of(port));
        }

        /**
         * @param protocolVersion The CQL protocol version to use.
         * 
         * @return builder
         * 
         */
        public Builder protocolVersion(@Nullable Output<Integer> protocolVersion) {
            $.protocolVersion = protocolVersion;
            return this;
        }

        /**
         * @param protocolVersion The CQL protocol version to use.
         * 
         * @return builder
         * 
         */
        public Builder protocolVersion(Integer protocolVersion) {
            return protocolVersion(Output.of(protocolVersion));
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
         * @param username The username to use when authenticating with Cassandra.
         * 
         * @return builder
         * 
         */
        public Builder username(@Nullable Output<String> username) {
            $.username = username;
            return this;
        }

        /**
         * @param username The username to use when authenticating with Cassandra.
         * 
         * @return builder
         * 
         */
        public Builder username(String username) {
            return username(Output.of(username));
        }

        public SecretBackendConnectionCassandraArgs build() {
            return $;
        }
    }

}
