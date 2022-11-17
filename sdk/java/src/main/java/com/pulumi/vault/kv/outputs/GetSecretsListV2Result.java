// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vault.kv.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class GetSecretsListV2Result {
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    private String mount;
    private @Nullable String name;
    /**
     * @return List of all secret names listed under the given path.
     * 
     */
    private List<String> names;
    private @Nullable String namespace;
    /**
     * @return Full path where the KV-V2 secrets are listed.
     * 
     */
    private String path;

    private GetSecretsListV2Result() {}
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
    public Optional<String> name() {
        return Optional.ofNullable(this.name);
    }
    /**
     * @return List of all secret names listed under the given path.
     * 
     */
    public List<String> names() {
        return this.names;
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

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetSecretsListV2Result defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String id;
        private String mount;
        private @Nullable String name;
        private List<String> names;
        private @Nullable String namespace;
        private String path;
        public Builder() {}
        public Builder(GetSecretsListV2Result defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.id = defaults.id;
    	      this.mount = defaults.mount;
    	      this.name = defaults.name;
    	      this.names = defaults.names;
    	      this.namespace = defaults.namespace;
    	      this.path = defaults.path;
        }

        @CustomType.Setter
        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        @CustomType.Setter
        public Builder mount(String mount) {
            this.mount = Objects.requireNonNull(mount);
            return this;
        }
        @CustomType.Setter
        public Builder name(@Nullable String name) {
            this.name = name;
            return this;
        }
        @CustomType.Setter
        public Builder names(List<String> names) {
            this.names = Objects.requireNonNull(names);
            return this;
        }
        public Builder names(String... names) {
            return names(List.of(names));
        }
        @CustomType.Setter
        public Builder namespace(@Nullable String namespace) {
            this.namespace = namespace;
            return this;
        }
        @CustomType.Setter
        public Builder path(String path) {
            this.path = Objects.requireNonNull(path);
            return this;
        }
        public GetSecretsListV2Result build() {
            final var o = new GetSecretsListV2Result();
            o.id = id;
            o.mount = mount;
            o.name = name;
            o.names = names;
            o.namespace = namespace;
            o.path = path;
            return o;
        }
    }
}