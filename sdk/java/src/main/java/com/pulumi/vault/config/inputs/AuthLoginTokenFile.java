// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vault.config.inputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class AuthLoginTokenFile {
    private String filename;
    private @Nullable String namespace;

    private AuthLoginTokenFile() {}
    public String filename() {
        return this.filename;
    }
    public Optional<String> namespace() {
        return Optional.ofNullable(this.namespace);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(AuthLoginTokenFile defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String filename;
        private @Nullable String namespace;
        public Builder() {}
        public Builder(AuthLoginTokenFile defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.filename = defaults.filename;
    	      this.namespace = defaults.namespace;
        }

        @CustomType.Setter
        public Builder filename(String filename) {
            this.filename = Objects.requireNonNull(filename);
            return this;
        }
        @CustomType.Setter
        public Builder namespace(@Nullable String namespace) {
            this.namespace = namespace;
            return this;
        }
        public AuthLoginTokenFile build() {
            final var o = new AuthLoginTokenFile();
            o.filename = filename;
            o.namespace = namespace;
            return o;
        }
    }
}
