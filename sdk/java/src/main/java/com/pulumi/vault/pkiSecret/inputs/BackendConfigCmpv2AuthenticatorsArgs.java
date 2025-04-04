// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vault.pkiSecret.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class BackendConfigCmpv2AuthenticatorsArgs extends com.pulumi.resources.ResourceArgs {

    public static final BackendConfigCmpv2AuthenticatorsArgs Empty = new BackendConfigCmpv2AuthenticatorsArgs();

    /**
     * &#34;The accessor (required) and cert_role (optional) properties for cert auth backends&#34;.
     * 
     */
    @Import(name="cert")
    private @Nullable Output<Map<String,String>> cert;

    /**
     * @return &#34;The accessor (required) and cert_role (optional) properties for cert auth backends&#34;.
     * 
     */
    public Optional<Output<Map<String,String>>> cert() {
        return Optional.ofNullable(this.cert);
    }

    private BackendConfigCmpv2AuthenticatorsArgs() {}

    private BackendConfigCmpv2AuthenticatorsArgs(BackendConfigCmpv2AuthenticatorsArgs $) {
        this.cert = $.cert;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(BackendConfigCmpv2AuthenticatorsArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private BackendConfigCmpv2AuthenticatorsArgs $;

        public Builder() {
            $ = new BackendConfigCmpv2AuthenticatorsArgs();
        }

        public Builder(BackendConfigCmpv2AuthenticatorsArgs defaults) {
            $ = new BackendConfigCmpv2AuthenticatorsArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param cert &#34;The accessor (required) and cert_role (optional) properties for cert auth backends&#34;.
         * 
         * @return builder
         * 
         */
        public Builder cert(@Nullable Output<Map<String,String>> cert) {
            $.cert = cert;
            return this;
        }

        /**
         * @param cert &#34;The accessor (required) and cert_role (optional) properties for cert auth backends&#34;.
         * 
         * @return builder
         * 
         */
        public Builder cert(Map<String,String> cert) {
            return cert(Output.of(cert));
        }

        public BackendConfigCmpv2AuthenticatorsArgs build() {
            return $;
        }
    }

}
