// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vault.rabbitMq.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class SecretBackendRoleVhost {
    private String configure;
    private String host;
    private String read;
    private String write;

    private SecretBackendRoleVhost() {}
    public String configure() {
        return this.configure;
    }
    public String host() {
        return this.host;
    }
    public String read() {
        return this.read;
    }
    public String write() {
        return this.write;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(SecretBackendRoleVhost defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String configure;
        private String host;
        private String read;
        private String write;
        public Builder() {}
        public Builder(SecretBackendRoleVhost defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.configure = defaults.configure;
    	      this.host = defaults.host;
    	      this.read = defaults.read;
    	      this.write = defaults.write;
        }

        @CustomType.Setter
        public Builder configure(String configure) {
            this.configure = Objects.requireNonNull(configure);
            return this;
        }
        @CustomType.Setter
        public Builder host(String host) {
            this.host = Objects.requireNonNull(host);
            return this;
        }
        @CustomType.Setter
        public Builder read(String read) {
            this.read = Objects.requireNonNull(read);
            return this;
        }
        @CustomType.Setter
        public Builder write(String write) {
            this.write = Objects.requireNonNull(write);
            return this;
        }
        public SecretBackendRoleVhost build() {
            final var o = new SecretBackendRoleVhost();
            o.configure = configure;
            o.host = host;
            o.read = read;
            o.write = write;
            return o;
        }
    }
}