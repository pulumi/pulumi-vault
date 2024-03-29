// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vault.database.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class SecretBackendConnectionCouchbaseArgs extends com.pulumi.resources.ResourceArgs {

    public static final SecretBackendConnectionCouchbaseArgs Empty = new SecretBackendConnectionCouchbaseArgs();

    /**
     * Required if `tls` is `true`. Specifies the certificate authority of the Couchbase server, as a PEM certificate that has been base64 encoded.
     * 
     */
    @Import(name="base64Pem")
    private @Nullable Output<String> base64Pem;

    /**
     * @return Required if `tls` is `true`. Specifies the certificate authority of the Couchbase server, as a PEM certificate that has been base64 encoded.
     * 
     */
    public Optional<Output<String>> base64Pem() {
        return Optional.ofNullable(this.base64Pem);
    }

    /**
     * Required for Couchbase versions prior to 6.5.0. This is only used to verify vault&#39;s connection to the server.
     * 
     */
    @Import(name="bucketName")
    private @Nullable Output<String> bucketName;

    /**
     * @return Required for Couchbase versions prior to 6.5.0. This is only used to verify vault&#39;s connection to the server.
     * 
     */
    public Optional<Output<String>> bucketName() {
        return Optional.ofNullable(this.bucketName);
    }

    /**
     * The hosts to connect to.
     * 
     */
    @Import(name="hosts", required=true)
    private Output<List<String>> hosts;

    /**
     * @return The hosts to connect to.
     * 
     */
    public Output<List<String>> hosts() {
        return this.hosts;
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

    private SecretBackendConnectionCouchbaseArgs() {}

    private SecretBackendConnectionCouchbaseArgs(SecretBackendConnectionCouchbaseArgs $) {
        this.base64Pem = $.base64Pem;
        this.bucketName = $.bucketName;
        this.hosts = $.hosts;
        this.insecureTls = $.insecureTls;
        this.password = $.password;
        this.tls = $.tls;
        this.username = $.username;
        this.usernameTemplate = $.usernameTemplate;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(SecretBackendConnectionCouchbaseArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private SecretBackendConnectionCouchbaseArgs $;

        public Builder() {
            $ = new SecretBackendConnectionCouchbaseArgs();
        }

        public Builder(SecretBackendConnectionCouchbaseArgs defaults) {
            $ = new SecretBackendConnectionCouchbaseArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param base64Pem Required if `tls` is `true`. Specifies the certificate authority of the Couchbase server, as a PEM certificate that has been base64 encoded.
         * 
         * @return builder
         * 
         */
        public Builder base64Pem(@Nullable Output<String> base64Pem) {
            $.base64Pem = base64Pem;
            return this;
        }

        /**
         * @param base64Pem Required if `tls` is `true`. Specifies the certificate authority of the Couchbase server, as a PEM certificate that has been base64 encoded.
         * 
         * @return builder
         * 
         */
        public Builder base64Pem(String base64Pem) {
            return base64Pem(Output.of(base64Pem));
        }

        /**
         * @param bucketName Required for Couchbase versions prior to 6.5.0. This is only used to verify vault&#39;s connection to the server.
         * 
         * @return builder
         * 
         */
        public Builder bucketName(@Nullable Output<String> bucketName) {
            $.bucketName = bucketName;
            return this;
        }

        /**
         * @param bucketName Required for Couchbase versions prior to 6.5.0. This is only used to verify vault&#39;s connection to the server.
         * 
         * @return builder
         * 
         */
        public Builder bucketName(String bucketName) {
            return bucketName(Output.of(bucketName));
        }

        /**
         * @param hosts The hosts to connect to.
         * 
         * @return builder
         * 
         */
        public Builder hosts(Output<List<String>> hosts) {
            $.hosts = hosts;
            return this;
        }

        /**
         * @param hosts The hosts to connect to.
         * 
         * @return builder
         * 
         */
        public Builder hosts(List<String> hosts) {
            return hosts(Output.of(hosts));
        }

        /**
         * @param hosts The hosts to connect to.
         * 
         * @return builder
         * 
         */
        public Builder hosts(String... hosts) {
            return hosts(List.of(hosts));
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

        public SecretBackendConnectionCouchbaseArgs build() {
            if ($.hosts == null) {
                throw new MissingRequiredPropertyException("SecretBackendConnectionCouchbaseArgs", "hosts");
            }
            if ($.password == null) {
                throw new MissingRequiredPropertyException("SecretBackendConnectionCouchbaseArgs", "password");
            }
            if ($.username == null) {
                throw new MissingRequiredPropertyException("SecretBackendConnectionCouchbaseArgs", "username");
            }
            return $;
        }
    }

}
