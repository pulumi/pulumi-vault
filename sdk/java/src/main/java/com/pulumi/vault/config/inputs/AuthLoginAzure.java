// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vault.config.inputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class AuthLoginAzure {
    private @Nullable String clientId;
    private @Nullable String jwt;
    private @Nullable String mount;
    private @Nullable String namespace;
    private String resourceGroupName;
    private String role;
    private @Nullable String scope;
    private String subscriptionId;
    private @Nullable String tenantId;
    private @Nullable String vmName;
    private @Nullable String vmssName;

    private AuthLoginAzure() {}
    public Optional<String> clientId() {
        return Optional.ofNullable(this.clientId);
    }
    public Optional<String> jwt() {
        return Optional.ofNullable(this.jwt);
    }
    public Optional<String> mount() {
        return Optional.ofNullable(this.mount);
    }
    public Optional<String> namespace() {
        return Optional.ofNullable(this.namespace);
    }
    public String resourceGroupName() {
        return this.resourceGroupName;
    }
    public String role() {
        return this.role;
    }
    public Optional<String> scope() {
        return Optional.ofNullable(this.scope);
    }
    public String subscriptionId() {
        return this.subscriptionId;
    }
    public Optional<String> tenantId() {
        return Optional.ofNullable(this.tenantId);
    }
    public Optional<String> vmName() {
        return Optional.ofNullable(this.vmName);
    }
    public Optional<String> vmssName() {
        return Optional.ofNullable(this.vmssName);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(AuthLoginAzure defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable String clientId;
        private @Nullable String jwt;
        private @Nullable String mount;
        private @Nullable String namespace;
        private String resourceGroupName;
        private String role;
        private @Nullable String scope;
        private String subscriptionId;
        private @Nullable String tenantId;
        private @Nullable String vmName;
        private @Nullable String vmssName;
        public Builder() {}
        public Builder(AuthLoginAzure defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.clientId = defaults.clientId;
    	      this.jwt = defaults.jwt;
    	      this.mount = defaults.mount;
    	      this.namespace = defaults.namespace;
    	      this.resourceGroupName = defaults.resourceGroupName;
    	      this.role = defaults.role;
    	      this.scope = defaults.scope;
    	      this.subscriptionId = defaults.subscriptionId;
    	      this.tenantId = defaults.tenantId;
    	      this.vmName = defaults.vmName;
    	      this.vmssName = defaults.vmssName;
        }

        @CustomType.Setter
        public Builder clientId(@Nullable String clientId) {
            this.clientId = clientId;
            return this;
        }
        @CustomType.Setter
        public Builder jwt(@Nullable String jwt) {
            this.jwt = jwt;
            return this;
        }
        @CustomType.Setter
        public Builder mount(@Nullable String mount) {
            this.mount = mount;
            return this;
        }
        @CustomType.Setter
        public Builder namespace(@Nullable String namespace) {
            this.namespace = namespace;
            return this;
        }
        @CustomType.Setter
        public Builder resourceGroupName(String resourceGroupName) {
            this.resourceGroupName = Objects.requireNonNull(resourceGroupName);
            return this;
        }
        @CustomType.Setter
        public Builder role(String role) {
            this.role = Objects.requireNonNull(role);
            return this;
        }
        @CustomType.Setter
        public Builder scope(@Nullable String scope) {
            this.scope = scope;
            return this;
        }
        @CustomType.Setter
        public Builder subscriptionId(String subscriptionId) {
            this.subscriptionId = Objects.requireNonNull(subscriptionId);
            return this;
        }
        @CustomType.Setter
        public Builder tenantId(@Nullable String tenantId) {
            this.tenantId = tenantId;
            return this;
        }
        @CustomType.Setter
        public Builder vmName(@Nullable String vmName) {
            this.vmName = vmName;
            return this;
        }
        @CustomType.Setter
        public Builder vmssName(@Nullable String vmssName) {
            this.vmssName = vmssName;
            return this;
        }
        public AuthLoginAzure build() {
            final var o = new AuthLoginAzure();
            o.clientId = clientId;
            o.jwt = jwt;
            o.mount = mount;
            o.namespace = namespace;
            o.resourceGroupName = resourceGroupName;
            o.role = role;
            o.scope = scope;
            o.subscriptionId = subscriptionId;
            o.tenantId = tenantId;
            o.vmName = vmName;
            o.vmssName = vmssName;
            return o;
        }
    }
}