// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vault.config.inputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class AuthLoginRadius {
    /**
     * @return The path where the authentication engine is mounted.
     * 
     */
    private @Nullable String mount;
    /**
     * @return The authentication engine&#39;s namespace. Conflicts with use_root_namespace
     * 
     */
    private @Nullable String namespace;
    /**
     * @return The Radius password for username.
     * 
     */
    private @Nullable String password;
    /**
     * @return Authenticate to the root Vault namespace. Conflicts with namespace
     * 
     */
    private @Nullable Boolean useRootNamespace;
    /**
     * @return The Radius username.
     * 
     */
    private @Nullable String username;

    private AuthLoginRadius() {}
    /**
     * @return The path where the authentication engine is mounted.
     * 
     */
    public Optional<String> mount() {
        return Optional.ofNullable(this.mount);
    }
    /**
     * @return The authentication engine&#39;s namespace. Conflicts with use_root_namespace
     * 
     */
    public Optional<String> namespace() {
        return Optional.ofNullable(this.namespace);
    }
    /**
     * @return The Radius password for username.
     * 
     */
    public Optional<String> password() {
        return Optional.ofNullable(this.password);
    }
    /**
     * @return Authenticate to the root Vault namespace. Conflicts with namespace
     * 
     */
    public Optional<Boolean> useRootNamespace() {
        return Optional.ofNullable(this.useRootNamespace);
    }
    /**
     * @return The Radius username.
     * 
     */
    public Optional<String> username() {
        return Optional.ofNullable(this.username);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(AuthLoginRadius defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable String mount;
        private @Nullable String namespace;
        private @Nullable String password;
        private @Nullable Boolean useRootNamespace;
        private @Nullable String username;
        public Builder() {}
        public Builder(AuthLoginRadius defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.mount = defaults.mount;
    	      this.namespace = defaults.namespace;
    	      this.password = defaults.password;
    	      this.useRootNamespace = defaults.useRootNamespace;
    	      this.username = defaults.username;
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
        public Builder password(@Nullable String password) {

            this.password = password;
            return this;
        }
        @CustomType.Setter
        public Builder useRootNamespace(@Nullable Boolean useRootNamespace) {

            this.useRootNamespace = useRootNamespace;
            return this;
        }
        @CustomType.Setter
        public Builder username(@Nullable String username) {

            this.username = username;
            return this;
        }
        public AuthLoginRadius build() {
            final var _resultValue = new AuthLoginRadius();
            _resultValue.mount = mount;
            _resultValue.namespace = namespace;
            _resultValue.password = password;
            _resultValue.useRootNamespace = useRootNamespace;
            _resultValue.username = username;
            return _resultValue;
        }
    }
}
