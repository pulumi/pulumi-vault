// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vault.identity.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class GetOidcClientCredsResult {
    /**
     * @return The Client ID returned by Vault.
     * 
     */
    private String clientId;
    /**
     * @return The Client Secret Key returned by Vault.
     * For public OpenID Clients `client_secret` is set to an empty string `&#34;&#34;`
     * 
     */
    private String clientSecret;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    private String name;
    private @Nullable String namespace;

    private GetOidcClientCredsResult() {}
    /**
     * @return The Client ID returned by Vault.
     * 
     */
    public String clientId() {
        return this.clientId;
    }
    /**
     * @return The Client Secret Key returned by Vault.
     * For public OpenID Clients `client_secret` is set to an empty string `&#34;&#34;`
     * 
     */
    public String clientSecret() {
        return this.clientSecret;
    }
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    public String id() {
        return this.id;
    }
    public String name() {
        return this.name;
    }
    public Optional<String> namespace() {
        return Optional.ofNullable(this.namespace);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetOidcClientCredsResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String clientId;
        private String clientSecret;
        private String id;
        private String name;
        private @Nullable String namespace;
        public Builder() {}
        public Builder(GetOidcClientCredsResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.clientId = defaults.clientId;
    	      this.clientSecret = defaults.clientSecret;
    	      this.id = defaults.id;
    	      this.name = defaults.name;
    	      this.namespace = defaults.namespace;
        }

        @CustomType.Setter
        public Builder clientId(String clientId) {
            if (clientId == null) {
              throw new MissingRequiredPropertyException("GetOidcClientCredsResult", "clientId");
            }
            this.clientId = clientId;
            return this;
        }
        @CustomType.Setter
        public Builder clientSecret(String clientSecret) {
            if (clientSecret == null) {
              throw new MissingRequiredPropertyException("GetOidcClientCredsResult", "clientSecret");
            }
            this.clientSecret = clientSecret;
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            if (id == null) {
              throw new MissingRequiredPropertyException("GetOidcClientCredsResult", "id");
            }
            this.id = id;
            return this;
        }
        @CustomType.Setter
        public Builder name(String name) {
            if (name == null) {
              throw new MissingRequiredPropertyException("GetOidcClientCredsResult", "name");
            }
            this.name = name;
            return this;
        }
        @CustomType.Setter
        public Builder namespace(@Nullable String namespace) {

            this.namespace = namespace;
            return this;
        }
        public GetOidcClientCredsResult build() {
            final var _resultValue = new GetOidcClientCredsResult();
            _resultValue.clientId = clientId;
            _resultValue.clientSecret = clientSecret;
            _resultValue.id = id;
            _resultValue.name = name;
            _resultValue.namespace = namespace;
            return _resultValue;
        }
    }
}
