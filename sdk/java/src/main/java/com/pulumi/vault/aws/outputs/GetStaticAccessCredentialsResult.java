// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vault.aws.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class GetStaticAccessCredentialsResult {
    private String accessKey;
    private String backend;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    private String name;
    private @Nullable String namespace;
    private String secretKey;

    private GetStaticAccessCredentialsResult() {}
    public String accessKey() {
        return this.accessKey;
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
    public String name() {
        return this.name;
    }
    public Optional<String> namespace() {
        return Optional.ofNullable(this.namespace);
    }
    public String secretKey() {
        return this.secretKey;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetStaticAccessCredentialsResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String accessKey;
        private String backend;
        private String id;
        private String name;
        private @Nullable String namespace;
        private String secretKey;
        public Builder() {}
        public Builder(GetStaticAccessCredentialsResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.accessKey = defaults.accessKey;
    	      this.backend = defaults.backend;
    	      this.id = defaults.id;
    	      this.name = defaults.name;
    	      this.namespace = defaults.namespace;
    	      this.secretKey = defaults.secretKey;
        }

        @CustomType.Setter
        public Builder accessKey(String accessKey) {
            this.accessKey = Objects.requireNonNull(accessKey);
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
        public Builder name(String name) {
            this.name = Objects.requireNonNull(name);
            return this;
        }
        @CustomType.Setter
        public Builder namespace(@Nullable String namespace) {
            this.namespace = namespace;
            return this;
        }
        @CustomType.Setter
        public Builder secretKey(String secretKey) {
            this.secretKey = Objects.requireNonNull(secretKey);
            return this;
        }
        public GetStaticAccessCredentialsResult build() {
            final var o = new GetStaticAccessCredentialsResult();
            o.accessKey = accessKey;
            o.backend = backend;
            o.id = id;
            o.name = name;
            o.namespace = namespace;
            o.secretKey = secretKey;
            return o;
        }
    }
}