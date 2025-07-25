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


public final class BackendConfigScepAuthenticatorsArgs extends com.pulumi.resources.ResourceArgs {

    public static final BackendConfigScepAuthenticatorsArgs Empty = new BackendConfigScepAuthenticatorsArgs();

    /**
     * The accessor and cert_role properties for cert auth backends
     * 
     */
    @Import(name="cert")
    private @Nullable Output<Map<String,String>> cert;

    /**
     * @return The accessor and cert_role properties for cert auth backends
     * 
     */
    public Optional<Output<Map<String,String>>> cert() {
        return Optional.ofNullable(this.cert);
    }

    /**
     * The accessor property for SCEP auth backends
     * 
     */
    @Import(name="scep")
    private @Nullable Output<Map<String,String>> scep;

    /**
     * @return The accessor property for SCEP auth backends
     * 
     */
    public Optional<Output<Map<String,String>>> scep() {
        return Optional.ofNullable(this.scep);
    }

    private BackendConfigScepAuthenticatorsArgs() {}

    private BackendConfigScepAuthenticatorsArgs(BackendConfigScepAuthenticatorsArgs $) {
        this.cert = $.cert;
        this.scep = $.scep;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(BackendConfigScepAuthenticatorsArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private BackendConfigScepAuthenticatorsArgs $;

        public Builder() {
            $ = new BackendConfigScepAuthenticatorsArgs();
        }

        public Builder(BackendConfigScepAuthenticatorsArgs defaults) {
            $ = new BackendConfigScepAuthenticatorsArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param cert The accessor and cert_role properties for cert auth backends
         * 
         * @return builder
         * 
         */
        public Builder cert(@Nullable Output<Map<String,String>> cert) {
            $.cert = cert;
            return this;
        }

        /**
         * @param cert The accessor and cert_role properties for cert auth backends
         * 
         * @return builder
         * 
         */
        public Builder cert(Map<String,String> cert) {
            return cert(Output.of(cert));
        }

        /**
         * @param scep The accessor property for SCEP auth backends
         * 
         * @return builder
         * 
         */
        public Builder scep(@Nullable Output<Map<String,String>> scep) {
            $.scep = scep;
            return this;
        }

        /**
         * @param scep The accessor property for SCEP auth backends
         * 
         * @return builder
         * 
         */
        public Builder scep(Map<String,String> scep) {
            return scep(Output.of(scep));
        }

        public BackendConfigScepAuthenticatorsArgs build() {
            return $;
        }
    }

}
