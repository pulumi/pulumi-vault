// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vault.database.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class SecretBackendConnectionCouchbase {
    /**
     * @return Required if `tls` is `true`. Specifies the certificate authority of the Couchbase server, as a PEM certificate that has been base64 encoded.
     * 
     */
    private @Nullable String base64Pem;
    /**
     * @return Required for Couchbase versions prior to 6.5.0. This is only used to verify vault&#39;s connection to the server.
     * 
     */
    private @Nullable String bucketName;
    /**
     * @return The hosts to connect to.
     * 
     */
    private List<String> hosts;
    /**
     * @return Whether to skip verification of the server
     * certificate when using TLS.
     * 
     */
    private @Nullable Boolean insecureTls;
    /**
     * @return The password to authenticate with.
     * 
     */
    private String password;
    /**
     * @return Whether to use TLS when connecting to Cassandra.
     * 
     */
    private @Nullable Boolean tls;
    /**
     * @return The username to authenticate with.
     * 
     */
    private String username;
    /**
     * @return Template describing how dynamic usernames are generated.
     * 
     */
    private @Nullable String usernameTemplate;

    private SecretBackendConnectionCouchbase() {}
    /**
     * @return Required if `tls` is `true`. Specifies the certificate authority of the Couchbase server, as a PEM certificate that has been base64 encoded.
     * 
     */
    public Optional<String> base64Pem() {
        return Optional.ofNullable(this.base64Pem);
    }
    /**
     * @return Required for Couchbase versions prior to 6.5.0. This is only used to verify vault&#39;s connection to the server.
     * 
     */
    public Optional<String> bucketName() {
        return Optional.ofNullable(this.bucketName);
    }
    /**
     * @return The hosts to connect to.
     * 
     */
    public List<String> hosts() {
        return this.hosts;
    }
    /**
     * @return Whether to skip verification of the server
     * certificate when using TLS.
     * 
     */
    public Optional<Boolean> insecureTls() {
        return Optional.ofNullable(this.insecureTls);
    }
    /**
     * @return The password to authenticate with.
     * 
     */
    public String password() {
        return this.password;
    }
    /**
     * @return Whether to use TLS when connecting to Cassandra.
     * 
     */
    public Optional<Boolean> tls() {
        return Optional.ofNullable(this.tls);
    }
    /**
     * @return The username to authenticate with.
     * 
     */
    public String username() {
        return this.username;
    }
    /**
     * @return Template describing how dynamic usernames are generated.
     * 
     */
    public Optional<String> usernameTemplate() {
        return Optional.ofNullable(this.usernameTemplate);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(SecretBackendConnectionCouchbase defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable String base64Pem;
        private @Nullable String bucketName;
        private List<String> hosts;
        private @Nullable Boolean insecureTls;
        private String password;
        private @Nullable Boolean tls;
        private String username;
        private @Nullable String usernameTemplate;
        public Builder() {}
        public Builder(SecretBackendConnectionCouchbase defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.base64Pem = defaults.base64Pem;
    	      this.bucketName = defaults.bucketName;
    	      this.hosts = defaults.hosts;
    	      this.insecureTls = defaults.insecureTls;
    	      this.password = defaults.password;
    	      this.tls = defaults.tls;
    	      this.username = defaults.username;
    	      this.usernameTemplate = defaults.usernameTemplate;
        }

        @CustomType.Setter
        public Builder base64Pem(@Nullable String base64Pem) {
            this.base64Pem = base64Pem;
            return this;
        }
        @CustomType.Setter
        public Builder bucketName(@Nullable String bucketName) {
            this.bucketName = bucketName;
            return this;
        }
        @CustomType.Setter
        public Builder hosts(List<String> hosts) {
            this.hosts = Objects.requireNonNull(hosts);
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
        public Builder password(String password) {
            this.password = Objects.requireNonNull(password);
            return this;
        }
        @CustomType.Setter
        public Builder tls(@Nullable Boolean tls) {
            this.tls = tls;
            return this;
        }
        @CustomType.Setter
        public Builder username(String username) {
            this.username = Objects.requireNonNull(username);
            return this;
        }
        @CustomType.Setter
        public Builder usernameTemplate(@Nullable String usernameTemplate) {
            this.usernameTemplate = usernameTemplate;
            return this;
        }
        public SecretBackendConnectionCouchbase build() {
            final var o = new SecretBackendConnectionCouchbase();
            o.base64Pem = base64Pem;
            o.bucketName = bucketName;
            o.hosts = hosts;
            o.insecureTls = insecureTls;
            o.password = password;
            o.tls = tls;
            o.username = username;
            o.usernameTemplate = usernameTemplate;
            return o;
        }
    }
}
