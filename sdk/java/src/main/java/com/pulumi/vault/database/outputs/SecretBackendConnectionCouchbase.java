// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vault.database.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
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
     * @return A set of Couchbase URIs to connect to. Must use `couchbases://` scheme if `tls` is `true`.
     * 
     */
    private List<String> hosts;
    /**
     * @return Specifies whether to skip verification of the server certificate when using TLS.
     * 
     */
    private @Nullable Boolean insecureTls;
    /**
     * @return Specifies the password corresponding to the given username.
     * 
     */
    private String password;
    /**
     * @return Specifies whether to use TLS when connecting to Couchbase.
     * 
     */
    private @Nullable Boolean tls;
    /**
     * @return Specifies the username for Vault to use.
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
     * @return A set of Couchbase URIs to connect to. Must use `couchbases://` scheme if `tls` is `true`.
     * 
     */
    public List<String> hosts() {
        return this.hosts;
    }
    /**
     * @return Specifies whether to skip verification of the server certificate when using TLS.
     * 
     */
    public Optional<Boolean> insecureTls() {
        return Optional.ofNullable(this.insecureTls);
    }
    /**
     * @return Specifies the password corresponding to the given username.
     * 
     */
    public String password() {
        return this.password;
    }
    /**
     * @return Specifies whether to use TLS when connecting to Couchbase.
     * 
     */
    public Optional<Boolean> tls() {
        return Optional.ofNullable(this.tls);
    }
    /**
     * @return Specifies the username for Vault to use.
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
            if (hosts == null) {
              throw new MissingRequiredPropertyException("SecretBackendConnectionCouchbase", "hosts");
            }
            this.hosts = hosts;
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
            if (password == null) {
              throw new MissingRequiredPropertyException("SecretBackendConnectionCouchbase", "password");
            }
            this.password = password;
            return this;
        }
        @CustomType.Setter
        public Builder tls(@Nullable Boolean tls) {

            this.tls = tls;
            return this;
        }
        @CustomType.Setter
        public Builder username(String username) {
            if (username == null) {
              throw new MissingRequiredPropertyException("SecretBackendConnectionCouchbase", "username");
            }
            this.username = username;
            return this;
        }
        @CustomType.Setter
        public Builder usernameTemplate(@Nullable String usernameTemplate) {

            this.usernameTemplate = usernameTemplate;
            return this;
        }
        public SecretBackendConnectionCouchbase build() {
            final var _resultValue = new SecretBackendConnectionCouchbase();
            _resultValue.base64Pem = base64Pem;
            _resultValue.bucketName = bucketName;
            _resultValue.hosts = hosts;
            _resultValue.insecureTls = insecureTls;
            _resultValue.password = password;
            _resultValue.tls = tls;
            _resultValue.username = username;
            _resultValue.usernameTemplate = usernameTemplate;
            return _resultValue;
        }
    }
}
