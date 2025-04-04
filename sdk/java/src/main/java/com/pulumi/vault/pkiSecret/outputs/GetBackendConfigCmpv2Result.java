// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vault.pkiSecret.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import com.pulumi.vault.pkiSecret.outputs.GetBackendConfigCmpv2Authenticator;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class GetBackendConfigCmpv2Result {
    private List<String> auditFields;
    private List<GetBackendConfigCmpv2Authenticator> authenticators;
    private String backend;
    private String defaultPathPolicy;
    private @Nullable List<String> disabledValidations;
    private Boolean enableSentinelParsing;
    private Boolean enabled;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    private String lastUpdated;
    private @Nullable String namespace;

    private GetBackendConfigCmpv2Result() {}
    public List<String> auditFields() {
        return this.auditFields;
    }
    public List<GetBackendConfigCmpv2Authenticator> authenticators() {
        return this.authenticators;
    }
    public String backend() {
        return this.backend;
    }
    public String defaultPathPolicy() {
        return this.defaultPathPolicy;
    }
    public List<String> disabledValidations() {
        return this.disabledValidations == null ? List.of() : this.disabledValidations;
    }
    public Boolean enableSentinelParsing() {
        return this.enableSentinelParsing;
    }
    public Boolean enabled() {
        return this.enabled;
    }
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    public String id() {
        return this.id;
    }
    public String lastUpdated() {
        return this.lastUpdated;
    }
    public Optional<String> namespace() {
        return Optional.ofNullable(this.namespace);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetBackendConfigCmpv2Result defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private List<String> auditFields;
        private List<GetBackendConfigCmpv2Authenticator> authenticators;
        private String backend;
        private String defaultPathPolicy;
        private @Nullable List<String> disabledValidations;
        private Boolean enableSentinelParsing;
        private Boolean enabled;
        private String id;
        private String lastUpdated;
        private @Nullable String namespace;
        public Builder() {}
        public Builder(GetBackendConfigCmpv2Result defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.auditFields = defaults.auditFields;
    	      this.authenticators = defaults.authenticators;
    	      this.backend = defaults.backend;
    	      this.defaultPathPolicy = defaults.defaultPathPolicy;
    	      this.disabledValidations = defaults.disabledValidations;
    	      this.enableSentinelParsing = defaults.enableSentinelParsing;
    	      this.enabled = defaults.enabled;
    	      this.id = defaults.id;
    	      this.lastUpdated = defaults.lastUpdated;
    	      this.namespace = defaults.namespace;
        }

        @CustomType.Setter
        public Builder auditFields(List<String> auditFields) {
            if (auditFields == null) {
              throw new MissingRequiredPropertyException("GetBackendConfigCmpv2Result", "auditFields");
            }
            this.auditFields = auditFields;
            return this;
        }
        public Builder auditFields(String... auditFields) {
            return auditFields(List.of(auditFields));
        }
        @CustomType.Setter
        public Builder authenticators(List<GetBackendConfigCmpv2Authenticator> authenticators) {
            if (authenticators == null) {
              throw new MissingRequiredPropertyException("GetBackendConfigCmpv2Result", "authenticators");
            }
            this.authenticators = authenticators;
            return this;
        }
        public Builder authenticators(GetBackendConfigCmpv2Authenticator... authenticators) {
            return authenticators(List.of(authenticators));
        }
        @CustomType.Setter
        public Builder backend(String backend) {
            if (backend == null) {
              throw new MissingRequiredPropertyException("GetBackendConfigCmpv2Result", "backend");
            }
            this.backend = backend;
            return this;
        }
        @CustomType.Setter
        public Builder defaultPathPolicy(String defaultPathPolicy) {
            if (defaultPathPolicy == null) {
              throw new MissingRequiredPropertyException("GetBackendConfigCmpv2Result", "defaultPathPolicy");
            }
            this.defaultPathPolicy = defaultPathPolicy;
            return this;
        }
        @CustomType.Setter
        public Builder disabledValidations(@Nullable List<String> disabledValidations) {

            this.disabledValidations = disabledValidations;
            return this;
        }
        public Builder disabledValidations(String... disabledValidations) {
            return disabledValidations(List.of(disabledValidations));
        }
        @CustomType.Setter
        public Builder enableSentinelParsing(Boolean enableSentinelParsing) {
            if (enableSentinelParsing == null) {
              throw new MissingRequiredPropertyException("GetBackendConfigCmpv2Result", "enableSentinelParsing");
            }
            this.enableSentinelParsing = enableSentinelParsing;
            return this;
        }
        @CustomType.Setter
        public Builder enabled(Boolean enabled) {
            if (enabled == null) {
              throw new MissingRequiredPropertyException("GetBackendConfigCmpv2Result", "enabled");
            }
            this.enabled = enabled;
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            if (id == null) {
              throw new MissingRequiredPropertyException("GetBackendConfigCmpv2Result", "id");
            }
            this.id = id;
            return this;
        }
        @CustomType.Setter
        public Builder lastUpdated(String lastUpdated) {
            if (lastUpdated == null) {
              throw new MissingRequiredPropertyException("GetBackendConfigCmpv2Result", "lastUpdated");
            }
            this.lastUpdated = lastUpdated;
            return this;
        }
        @CustomType.Setter
        public Builder namespace(@Nullable String namespace) {

            this.namespace = namespace;
            return this;
        }
        public GetBackendConfigCmpv2Result build() {
            final var _resultValue = new GetBackendConfigCmpv2Result();
            _resultValue.auditFields = auditFields;
            _resultValue.authenticators = authenticators;
            _resultValue.backend = backend;
            _resultValue.defaultPathPolicy = defaultPathPolicy;
            _resultValue.disabledValidations = disabledValidations;
            _resultValue.enableSentinelParsing = enableSentinelParsing;
            _resultValue.enabled = enabled;
            _resultValue.id = id;
            _resultValue.lastUpdated = lastUpdated;
            _resultValue.namespace = namespace;
            return _resultValue;
        }
    }
}
