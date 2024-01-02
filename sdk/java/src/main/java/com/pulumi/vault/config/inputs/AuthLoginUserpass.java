// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vault.config.inputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class AuthLoginUserpass {
    private @Nullable String mount;
    private @Nullable String namespace;
    private @Nullable String password;
    private @Nullable String passwordFile;
    private @Nullable Boolean useRootNamespace;
    private String username;

    private AuthLoginUserpass() {}
    public Optional<String> mount() {
        return Optional.ofNullable(this.mount);
    }
    public Optional<String> namespace() {
        return Optional.ofNullable(this.namespace);
    }
    public Optional<String> password() {
        return Optional.ofNullable(this.password);
    }
    public Optional<String> passwordFile() {
        return Optional.ofNullable(this.passwordFile);
    }
    public Optional<Boolean> useRootNamespace() {
        return Optional.ofNullable(this.useRootNamespace);
    }
    public String username() {
        return this.username;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(AuthLoginUserpass defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable String mount;
        private @Nullable String namespace;
        private @Nullable String password;
        private @Nullable String passwordFile;
        private @Nullable Boolean useRootNamespace;
        private String username;
        public Builder() {}
        public Builder(AuthLoginUserpass defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.mount = defaults.mount;
    	      this.namespace = defaults.namespace;
    	      this.password = defaults.password;
    	      this.passwordFile = defaults.passwordFile;
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
        public Builder passwordFile(@Nullable String passwordFile) {

            this.passwordFile = passwordFile;
            return this;
        }
        @CustomType.Setter
        public Builder useRootNamespace(@Nullable Boolean useRootNamespace) {

            this.useRootNamespace = useRootNamespace;
            return this;
        }
        @CustomType.Setter
        public Builder username(String username) {
            if (username == null) {
              throw new MissingRequiredPropertyException("AuthLoginUserpass", "username");
            }
            this.username = username;
            return this;
        }
        public AuthLoginUserpass build() {
            final var _resultValue = new AuthLoginUserpass();
            _resultValue.mount = mount;
            _resultValue.namespace = namespace;
            _resultValue.password = password;
            _resultValue.passwordFile = passwordFile;
            _resultValue.useRootNamespace = useRootNamespace;
            _resultValue.username = username;
            return _resultValue;
        }
    }
}
