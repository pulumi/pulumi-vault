// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vault.pkiSecret.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Map;
import java.util.Objects;
import javax.annotation.Nullable;

@CustomType
public final class GetBackendConfigEstAuthenticator {
    /**
     * @return &#34;The accessor and cert_role properties for cert auth backends&#34;.
     * 
     */
    private @Nullable Map<String,String> cert;
    /**
     * @return &#34;The accessor property for user pass auth backends&#34;.
     * 
     */
    private @Nullable Map<String,String> userpass;

    private GetBackendConfigEstAuthenticator() {}
    /**
     * @return &#34;The accessor and cert_role properties for cert auth backends&#34;.
     * 
     */
    public Map<String,String> cert() {
        return this.cert == null ? Map.of() : this.cert;
    }
    /**
     * @return &#34;The accessor property for user pass auth backends&#34;.
     * 
     */
    public Map<String,String> userpass() {
        return this.userpass == null ? Map.of() : this.userpass;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetBackendConfigEstAuthenticator defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable Map<String,String> cert;
        private @Nullable Map<String,String> userpass;
        public Builder() {}
        public Builder(GetBackendConfigEstAuthenticator defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.cert = defaults.cert;
    	      this.userpass = defaults.userpass;
        }

        @CustomType.Setter
        public Builder cert(@Nullable Map<String,String> cert) {

            this.cert = cert;
            return this;
        }
        @CustomType.Setter
        public Builder userpass(@Nullable Map<String,String> userpass) {

            this.userpass = userpass;
            return this;
        }
        public GetBackendConfigEstAuthenticator build() {
            final var _resultValue = new GetBackendConfigEstAuthenticator();
            _resultValue.cert = cert;
            _resultValue.userpass = userpass;
            return _resultValue;
        }
    }
}
