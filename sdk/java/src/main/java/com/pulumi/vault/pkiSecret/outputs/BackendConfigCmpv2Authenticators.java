// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vault.pkiSecret.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Map;
import java.util.Objects;
import javax.annotation.Nullable;

@CustomType
public final class BackendConfigCmpv2Authenticators {
    /**
     * @return &#34;The accessor (required) and cert_role (optional) properties for cert auth backends&#34;.
     * 
     */
    private @Nullable Map<String,String> cert;

    private BackendConfigCmpv2Authenticators() {}
    /**
     * @return &#34;The accessor (required) and cert_role (optional) properties for cert auth backends&#34;.
     * 
     */
    public Map<String,String> cert() {
        return this.cert == null ? Map.of() : this.cert;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(BackendConfigCmpv2Authenticators defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable Map<String,String> cert;
        public Builder() {}
        public Builder(BackendConfigCmpv2Authenticators defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.cert = defaults.cert;
        }

        @CustomType.Setter
        public Builder cert(@Nullable Map<String,String> cert) {

            this.cert = cert;
            return this;
        }
        public BackendConfigCmpv2Authenticators build() {
            final var _resultValue = new BackendConfigCmpv2Authenticators();
            _resultValue.cert = cert;
            return _resultValue;
        }
    }
}
