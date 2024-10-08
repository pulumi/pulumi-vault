// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vault.kv.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Integer;
import java.lang.String;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class GetSecretSubkeysV2Result {
    /**
     * @return Subkeys for the KV-V2 secret stored as a serialized map of strings.
     * 
     */
    private Map<String,String> data;
    /**
     * @return Subkeys for the KV-V2 secret read from Vault.
     * 
     */
    private String dataJson;
    private @Nullable Integer depth;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    private String mount;
    private String name;
    private @Nullable String namespace;
    /**
     * @return Full path where the KV-V2 secrets are listed.
     * 
     */
    private String path;
    private @Nullable Integer version;

    private GetSecretSubkeysV2Result() {}
    /**
     * @return Subkeys for the KV-V2 secret stored as a serialized map of strings.
     * 
     */
    public Map<String,String> data() {
        return this.data;
    }
    /**
     * @return Subkeys for the KV-V2 secret read from Vault.
     * 
     */
    public String dataJson() {
        return this.dataJson;
    }
    public Optional<Integer> depth() {
        return Optional.ofNullable(this.depth);
    }
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    public String id() {
        return this.id;
    }
    public String mount() {
        return this.mount;
    }
    public String name() {
        return this.name;
    }
    public Optional<String> namespace() {
        return Optional.ofNullable(this.namespace);
    }
    /**
     * @return Full path where the KV-V2 secrets are listed.
     * 
     */
    public String path() {
        return this.path;
    }
    public Optional<Integer> version() {
        return Optional.ofNullable(this.version);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetSecretSubkeysV2Result defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private Map<String,String> data;
        private String dataJson;
        private @Nullable Integer depth;
        private String id;
        private String mount;
        private String name;
        private @Nullable String namespace;
        private String path;
        private @Nullable Integer version;
        public Builder() {}
        public Builder(GetSecretSubkeysV2Result defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.data = defaults.data;
    	      this.dataJson = defaults.dataJson;
    	      this.depth = defaults.depth;
    	      this.id = defaults.id;
    	      this.mount = defaults.mount;
    	      this.name = defaults.name;
    	      this.namespace = defaults.namespace;
    	      this.path = defaults.path;
    	      this.version = defaults.version;
        }

        @CustomType.Setter
        public Builder data(Map<String,String> data) {
            if (data == null) {
              throw new MissingRequiredPropertyException("GetSecretSubkeysV2Result", "data");
            }
            this.data = data;
            return this;
        }
        @CustomType.Setter
        public Builder dataJson(String dataJson) {
            if (dataJson == null) {
              throw new MissingRequiredPropertyException("GetSecretSubkeysV2Result", "dataJson");
            }
            this.dataJson = dataJson;
            return this;
        }
        @CustomType.Setter
        public Builder depth(@Nullable Integer depth) {

            this.depth = depth;
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            if (id == null) {
              throw new MissingRequiredPropertyException("GetSecretSubkeysV2Result", "id");
            }
            this.id = id;
            return this;
        }
        @CustomType.Setter
        public Builder mount(String mount) {
            if (mount == null) {
              throw new MissingRequiredPropertyException("GetSecretSubkeysV2Result", "mount");
            }
            this.mount = mount;
            return this;
        }
        @CustomType.Setter
        public Builder name(String name) {
            if (name == null) {
              throw new MissingRequiredPropertyException("GetSecretSubkeysV2Result", "name");
            }
            this.name = name;
            return this;
        }
        @CustomType.Setter
        public Builder namespace(@Nullable String namespace) {

            this.namespace = namespace;
            return this;
        }
        @CustomType.Setter
        public Builder path(String path) {
            if (path == null) {
              throw new MissingRequiredPropertyException("GetSecretSubkeysV2Result", "path");
            }
            this.path = path;
            return this;
        }
        @CustomType.Setter
        public Builder version(@Nullable Integer version) {

            this.version = version;
            return this;
        }
        public GetSecretSubkeysV2Result build() {
            final var _resultValue = new GetSecretSubkeysV2Result();
            _resultValue.data = data;
            _resultValue.dataJson = dataJson;
            _resultValue.depth = depth;
            _resultValue.id = id;
            _resultValue.mount = mount;
            _resultValue.name = name;
            _resultValue.namespace = namespace;
            _resultValue.path = path;
            _resultValue.version = version;
            return _resultValue;
        }
    }
}
