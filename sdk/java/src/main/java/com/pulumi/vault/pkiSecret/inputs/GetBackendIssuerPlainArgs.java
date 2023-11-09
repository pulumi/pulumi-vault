// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vault.pkiSecret.inputs;

import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetBackendIssuerPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetBackendIssuerPlainArgs Empty = new GetBackendIssuerPlainArgs();

    /**
     * The path to the PKI secret backend to
     * read the issuer from, with no leading or trailing `/`s.
     * 
     */
    @Import(name="backend", required=true)
    private String backend;

    /**
     * @return The path to the PKI secret backend to
     * read the issuer from, with no leading or trailing `/`s.
     * 
     */
    public String backend() {
        return this.backend;
    }

    /**
     * Reference to an existing issuer.
     * 
     */
    @Import(name="issuerRef", required=true)
    private String issuerRef;

    /**
     * @return Reference to an existing issuer.
     * 
     */
    public String issuerRef() {
        return this.issuerRef;
    }

    /**
     * The namespace of the target resource.
     * The value should not contain leading or trailing forward slashes.
     * The `namespace` is always relative to the provider&#39;s configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
     * *Available only for Vault Enterprise*.
     * 
     */
    @Import(name="namespace")
    private @Nullable String namespace;

    /**
     * @return The namespace of the target resource.
     * The value should not contain leading or trailing forward slashes.
     * The `namespace` is always relative to the provider&#39;s configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
     * *Available only for Vault Enterprise*.
     * 
     */
    public Optional<String> namespace() {
        return Optional.ofNullable(this.namespace);
    }

    private GetBackendIssuerPlainArgs() {}

    private GetBackendIssuerPlainArgs(GetBackendIssuerPlainArgs $) {
        this.backend = $.backend;
        this.issuerRef = $.issuerRef;
        this.namespace = $.namespace;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetBackendIssuerPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetBackendIssuerPlainArgs $;

        public Builder() {
            $ = new GetBackendIssuerPlainArgs();
        }

        public Builder(GetBackendIssuerPlainArgs defaults) {
            $ = new GetBackendIssuerPlainArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param backend The path to the PKI secret backend to
         * read the issuer from, with no leading or trailing `/`s.
         * 
         * @return builder
         * 
         */
        public Builder backend(String backend) {
            $.backend = backend;
            return this;
        }

        /**
         * @param issuerRef Reference to an existing issuer.
         * 
         * @return builder
         * 
         */
        public Builder issuerRef(String issuerRef) {
            $.issuerRef = issuerRef;
            return this;
        }

        /**
         * @param namespace The namespace of the target resource.
         * The value should not contain leading or trailing forward slashes.
         * The `namespace` is always relative to the provider&#39;s configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
         * *Available only for Vault Enterprise*.
         * 
         * @return builder
         * 
         */
        public Builder namespace(@Nullable String namespace) {
            $.namespace = namespace;
            return this;
        }

        public GetBackendIssuerPlainArgs build() {
            $.backend = Objects.requireNonNull($.backend, "expected parameter 'backend' to be non-null");
            $.issuerRef = Objects.requireNonNull($.issuerRef, "expected parameter 'issuerRef' to be non-null");
            return $;
        }
    }

}
