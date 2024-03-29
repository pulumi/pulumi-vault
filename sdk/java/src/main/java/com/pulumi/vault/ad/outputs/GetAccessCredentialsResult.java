// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vault.ad.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class GetAccessCredentialsResult {
    private String backend;
    /**
     * @return The current set password on the Active Directory service account.
     * 
     */
    private String currentPassword;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    /**
     * @return The current set password on the Active Directory service account, provided because AD is eventually consistent.
     * 
     */
    private String lastPassword;
    private @Nullable String namespace;
    private String role;
    /**
     * @return The Active Directory service account username.
     * 
     */
    private String username;

    private GetAccessCredentialsResult() {}
    public String backend() {
        return this.backend;
    }
    /**
     * @return The current set password on the Active Directory service account.
     * 
     */
    public String currentPassword() {
        return this.currentPassword;
    }
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @return The current set password on the Active Directory service account, provided because AD is eventually consistent.
     * 
     */
    public String lastPassword() {
        return this.lastPassword;
    }
    public Optional<String> namespace() {
        return Optional.ofNullable(this.namespace);
    }
    public String role() {
        return this.role;
    }
    /**
     * @return The Active Directory service account username.
     * 
     */
    public String username() {
        return this.username;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetAccessCredentialsResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String backend;
        private String currentPassword;
        private String id;
        private String lastPassword;
        private @Nullable String namespace;
        private String role;
        private String username;
        public Builder() {}
        public Builder(GetAccessCredentialsResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.backend = defaults.backend;
    	      this.currentPassword = defaults.currentPassword;
    	      this.id = defaults.id;
    	      this.lastPassword = defaults.lastPassword;
    	      this.namespace = defaults.namespace;
    	      this.role = defaults.role;
    	      this.username = defaults.username;
        }

        @CustomType.Setter
        public Builder backend(String backend) {
            if (backend == null) {
              throw new MissingRequiredPropertyException("GetAccessCredentialsResult", "backend");
            }
            this.backend = backend;
            return this;
        }
        @CustomType.Setter
        public Builder currentPassword(String currentPassword) {
            if (currentPassword == null) {
              throw new MissingRequiredPropertyException("GetAccessCredentialsResult", "currentPassword");
            }
            this.currentPassword = currentPassword;
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            if (id == null) {
              throw new MissingRequiredPropertyException("GetAccessCredentialsResult", "id");
            }
            this.id = id;
            return this;
        }
        @CustomType.Setter
        public Builder lastPassword(String lastPassword) {
            if (lastPassword == null) {
              throw new MissingRequiredPropertyException("GetAccessCredentialsResult", "lastPassword");
            }
            this.lastPassword = lastPassword;
            return this;
        }
        @CustomType.Setter
        public Builder namespace(@Nullable String namespace) {

            this.namespace = namespace;
            return this;
        }
        @CustomType.Setter
        public Builder role(String role) {
            if (role == null) {
              throw new MissingRequiredPropertyException("GetAccessCredentialsResult", "role");
            }
            this.role = role;
            return this;
        }
        @CustomType.Setter
        public Builder username(String username) {
            if (username == null) {
              throw new MissingRequiredPropertyException("GetAccessCredentialsResult", "username");
            }
            this.username = username;
            return this;
        }
        public GetAccessCredentialsResult build() {
            final var _resultValue = new GetAccessCredentialsResult();
            _resultValue.backend = backend;
            _resultValue.currentPassword = currentPassword;
            _resultValue.id = id;
            _resultValue.lastPassword = lastPassword;
            _resultValue.namespace = namespace;
            _resultValue.role = role;
            _resultValue.username = username;
            return _resultValue;
        }
    }
}
