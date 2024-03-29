// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vault.gcp.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class AuthBackendCustomEndpoint {
    /**
     * @return Replaces the service endpoint used in API requests to `https://www.googleapis.com`.
     * 
     */
    private @Nullable String api;
    /**
     * @return Replaces the service endpoint used in API requests to `https://compute.googleapis.com`.
     * 
     * The endpoint value provided for a given key has the form of `scheme://host:port`.
     * The `scheme://` and `:port` portions of the endpoint value are optional.
     * 
     */
    private @Nullable String compute;
    /**
     * @return Replaces the service endpoint used in API requests to `https://cloudresourcemanager.googleapis.com`.
     * 
     */
    private @Nullable String crm;
    /**
     * @return Replaces the service endpoint used in API requests to `https://iam.googleapis.com`.
     * 
     */
    private @Nullable String iam;

    private AuthBackendCustomEndpoint() {}
    /**
     * @return Replaces the service endpoint used in API requests to `https://www.googleapis.com`.
     * 
     */
    public Optional<String> api() {
        return Optional.ofNullable(this.api);
    }
    /**
     * @return Replaces the service endpoint used in API requests to `https://compute.googleapis.com`.
     * 
     * The endpoint value provided for a given key has the form of `scheme://host:port`.
     * The `scheme://` and `:port` portions of the endpoint value are optional.
     * 
     */
    public Optional<String> compute() {
        return Optional.ofNullable(this.compute);
    }
    /**
     * @return Replaces the service endpoint used in API requests to `https://cloudresourcemanager.googleapis.com`.
     * 
     */
    public Optional<String> crm() {
        return Optional.ofNullable(this.crm);
    }
    /**
     * @return Replaces the service endpoint used in API requests to `https://iam.googleapis.com`.
     * 
     */
    public Optional<String> iam() {
        return Optional.ofNullable(this.iam);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(AuthBackendCustomEndpoint defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable String api;
        private @Nullable String compute;
        private @Nullable String crm;
        private @Nullable String iam;
        public Builder() {}
        public Builder(AuthBackendCustomEndpoint defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.api = defaults.api;
    	      this.compute = defaults.compute;
    	      this.crm = defaults.crm;
    	      this.iam = defaults.iam;
        }

        @CustomType.Setter
        public Builder api(@Nullable String api) {

            this.api = api;
            return this;
        }
        @CustomType.Setter
        public Builder compute(@Nullable String compute) {

            this.compute = compute;
            return this;
        }
        @CustomType.Setter
        public Builder crm(@Nullable String crm) {

            this.crm = crm;
            return this;
        }
        @CustomType.Setter
        public Builder iam(@Nullable String iam) {

            this.iam = iam;
            return this;
        }
        public AuthBackendCustomEndpoint build() {
            final var _resultValue = new AuthBackendCustomEndpoint();
            _resultValue.api = api;
            _resultValue.compute = compute;
            _resultValue.crm = crm;
            _resultValue.iam = iam;
            return _resultValue;
        }
    }
}
