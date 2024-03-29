// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vault.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import com.pulumi.vault.outputs.GetPolicyDocumentRule;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class GetPolicyDocumentResult {
    /**
     * @return The above arguments serialized as a standard Vault HCL policy document.
     * 
     */
    private String hcl;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    private @Nullable String namespace;
    private List<GetPolicyDocumentRule> rules;

    private GetPolicyDocumentResult() {}
    /**
     * @return The above arguments serialized as a standard Vault HCL policy document.
     * 
     */
    public String hcl() {
        return this.hcl;
    }
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    public String id() {
        return this.id;
    }
    public Optional<String> namespace() {
        return Optional.ofNullable(this.namespace);
    }
    public List<GetPolicyDocumentRule> rules() {
        return this.rules;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetPolicyDocumentResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String hcl;
        private String id;
        private @Nullable String namespace;
        private List<GetPolicyDocumentRule> rules;
        public Builder() {}
        public Builder(GetPolicyDocumentResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.hcl = defaults.hcl;
    	      this.id = defaults.id;
    	      this.namespace = defaults.namespace;
    	      this.rules = defaults.rules;
        }

        @CustomType.Setter
        public Builder hcl(String hcl) {
            if (hcl == null) {
              throw new MissingRequiredPropertyException("GetPolicyDocumentResult", "hcl");
            }
            this.hcl = hcl;
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            if (id == null) {
              throw new MissingRequiredPropertyException("GetPolicyDocumentResult", "id");
            }
            this.id = id;
            return this;
        }
        @CustomType.Setter
        public Builder namespace(@Nullable String namespace) {

            this.namespace = namespace;
            return this;
        }
        @CustomType.Setter
        public Builder rules(List<GetPolicyDocumentRule> rules) {
            if (rules == null) {
              throw new MissingRequiredPropertyException("GetPolicyDocumentResult", "rules");
            }
            this.rules = rules;
            return this;
        }
        public Builder rules(GetPolicyDocumentRule... rules) {
            return rules(List.of(rules));
        }
        public GetPolicyDocumentResult build() {
            final var _resultValue = new GetPolicyDocumentResult();
            _resultValue.hcl = hcl;
            _resultValue.id = id;
            _resultValue.namespace = namespace;
            _resultValue.rules = rules;
            return _resultValue;
        }
    }
}
