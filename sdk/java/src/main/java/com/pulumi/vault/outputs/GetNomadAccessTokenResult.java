// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vault.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class GetNomadAccessTokenResult {
    /**
     * @return The public identifier for a specific token. It can be used
     * to look up information about a token or to revoke a token.
     * 
     */
    private String accessorId;
    private String backend;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    private @Nullable String namespace;
    private String role;
    /**
     * @return The token to be used when making requests to Nomad and should be kept private.
     * 
     */
    private String secretId;

    private GetNomadAccessTokenResult() {}
    /**
     * @return The public identifier for a specific token. It can be used
     * to look up information about a token or to revoke a token.
     * 
     */
    public String accessorId() {
        return this.accessorId;
    }
    public String backend() {
        return this.backend;
    }
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    public String id() {
        return this.id;
    }
    public Optional<String> namespace() {
        return Optional.ofNullable(this.namespace);
    }
    public String role() {
        return this.role;
    }
    /**
     * @return The token to be used when making requests to Nomad and should be kept private.
     * 
     */
    public String secretId() {
        return this.secretId;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetNomadAccessTokenResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String accessorId;
        private String backend;
        private String id;
        private @Nullable String namespace;
        private String role;
        private String secretId;
        public Builder() {}
        public Builder(GetNomadAccessTokenResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.accessorId = defaults.accessorId;
    	      this.backend = defaults.backend;
    	      this.id = defaults.id;
    	      this.namespace = defaults.namespace;
    	      this.role = defaults.role;
    	      this.secretId = defaults.secretId;
        }

        @CustomType.Setter
        public Builder accessorId(String accessorId) {
            this.accessorId = Objects.requireNonNull(accessorId);
            return this;
        }
        @CustomType.Setter
        public Builder backend(String backend) {
            this.backend = Objects.requireNonNull(backend);
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        @CustomType.Setter
        public Builder namespace(@Nullable String namespace) {
            this.namespace = namespace;
            return this;
        }
        @CustomType.Setter
        public Builder role(String role) {
            this.role = Objects.requireNonNull(role);
            return this;
        }
        @CustomType.Setter
        public Builder secretId(String secretId) {
            this.secretId = Objects.requireNonNull(secretId);
            return this;
        }
        public GetNomadAccessTokenResult build() {
            final var o = new GetNomadAccessTokenResult();
            o.accessorId = accessorId;
            o.backend = backend;
            o.id = id;
            o.namespace = namespace;
            o.role = role;
            o.secretId = secretId;
            return o;
        }
    }
}