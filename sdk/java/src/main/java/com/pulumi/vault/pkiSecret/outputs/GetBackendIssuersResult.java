// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vault.pkiSecret.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Object;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class GetBackendIssuersResult {
    private String backend;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    /**
     * @return Map of issuer strings read from Vault.
     * 
     */
    private Map<String,Object> keyInfo;
    /**
     * @return JSON-encoded issuer data read from Vault.
     * 
     */
    private String keyInfoJson;
    /**
     * @return Keys used by issuers under the backend path.
     * 
     */
    private List<String> keys;
    private @Nullable String namespace;

    private GetBackendIssuersResult() {}
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
    /**
     * @return Map of issuer strings read from Vault.
     * 
     */
    public Map<String,Object> keyInfo() {
        return this.keyInfo;
    }
    /**
     * @return JSON-encoded issuer data read from Vault.
     * 
     */
    public String keyInfoJson() {
        return this.keyInfoJson;
    }
    /**
     * @return Keys used by issuers under the backend path.
     * 
     */
    public List<String> keys() {
        return this.keys;
    }
    public Optional<String> namespace() {
        return Optional.ofNullable(this.namespace);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetBackendIssuersResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String backend;
        private String id;
        private Map<String,Object> keyInfo;
        private String keyInfoJson;
        private List<String> keys;
        private @Nullable String namespace;
        public Builder() {}
        public Builder(GetBackendIssuersResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.backend = defaults.backend;
    	      this.id = defaults.id;
    	      this.keyInfo = defaults.keyInfo;
    	      this.keyInfoJson = defaults.keyInfoJson;
    	      this.keys = defaults.keys;
    	      this.namespace = defaults.namespace;
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
        public Builder keyInfo(Map<String,Object> keyInfo) {
            this.keyInfo = Objects.requireNonNull(keyInfo);
            return this;
        }
        @CustomType.Setter
        public Builder keyInfoJson(String keyInfoJson) {
            this.keyInfoJson = Objects.requireNonNull(keyInfoJson);
            return this;
        }
        @CustomType.Setter
        public Builder keys(List<String> keys) {
            this.keys = Objects.requireNonNull(keys);
            return this;
        }
        public Builder keys(String... keys) {
            return keys(List.of(keys));
        }
        @CustomType.Setter
        public Builder namespace(@Nullable String namespace) {
            this.namespace = namespace;
            return this;
        }
        public GetBackendIssuersResult build() {
            final var _resultValue = new GetBackendIssuersResult();
            _resultValue.backend = backend;
            _resultValue.id = id;
            _resultValue.keyInfo = keyInfo;
            _resultValue.keyInfoJson = keyInfoJson;
            _resultValue.keys = keys;
            _resultValue.namespace = namespace;
            return _resultValue;
        }
    }
}
