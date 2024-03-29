// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vault.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;


public final class ProviderClientAuthArgs extends com.pulumi.resources.ResourceArgs {

    public static final ProviderClientAuthArgs Empty = new ProviderClientAuthArgs();

    /**
     * Path to a file containing the client certificate.
     * 
     */
    @Import(name="certFile", required=true)
    private Output<String> certFile;

    /**
     * @return Path to a file containing the client certificate.
     * 
     */
    public Output<String> certFile() {
        return this.certFile;
    }

    /**
     * Path to a file containing the private key that the certificate was issued for.
     * 
     */
    @Import(name="keyFile", required=true)
    private Output<String> keyFile;

    /**
     * @return Path to a file containing the private key that the certificate was issued for.
     * 
     */
    public Output<String> keyFile() {
        return this.keyFile;
    }

    private ProviderClientAuthArgs() {}

    private ProviderClientAuthArgs(ProviderClientAuthArgs $) {
        this.certFile = $.certFile;
        this.keyFile = $.keyFile;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ProviderClientAuthArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ProviderClientAuthArgs $;

        public Builder() {
            $ = new ProviderClientAuthArgs();
        }

        public Builder(ProviderClientAuthArgs defaults) {
            $ = new ProviderClientAuthArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param certFile Path to a file containing the client certificate.
         * 
         * @return builder
         * 
         */
        public Builder certFile(Output<String> certFile) {
            $.certFile = certFile;
            return this;
        }

        /**
         * @param certFile Path to a file containing the client certificate.
         * 
         * @return builder
         * 
         */
        public Builder certFile(String certFile) {
            return certFile(Output.of(certFile));
        }

        /**
         * @param keyFile Path to a file containing the private key that the certificate was issued for.
         * 
         * @return builder
         * 
         */
        public Builder keyFile(Output<String> keyFile) {
            $.keyFile = keyFile;
            return this;
        }

        /**
         * @param keyFile Path to a file containing the private key that the certificate was issued for.
         * 
         * @return builder
         * 
         */
        public Builder keyFile(String keyFile) {
            return keyFile(Output.of(keyFile));
        }

        public ProviderClientAuthArgs build() {
            if ($.certFile == null) {
                throw new MissingRequiredPropertyException("ProviderClientAuthArgs", "certFile");
            }
            if ($.keyFile == null) {
                throw new MissingRequiredPropertyException("ProviderClientAuthArgs", "keyFile");
            }
            return $;
        }
    }

}
