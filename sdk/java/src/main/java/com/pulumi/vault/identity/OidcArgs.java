// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vault.identity;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class OidcArgs extends com.pulumi.resources.ResourceArgs {

    public static final OidcArgs Empty = new OidcArgs();

    /**
     * Issuer URL to be used in the iss claim of the token. If not set, Vault&#39;s
     * `api_addr` will be used. The issuer is a case sensitive URL using the https scheme that contains
     * scheme, host, and optionally, port number and path components, but no query or fragment
     * components.
     * 
     */
    @Import(name="issuer")
    private @Nullable Output<String> issuer;

    /**
     * @return Issuer URL to be used in the iss claim of the token. If not set, Vault&#39;s
     * `api_addr` will be used. The issuer is a case sensitive URL using the https scheme that contains
     * scheme, host, and optionally, port number and path components, but no query or fragment
     * components.
     * 
     */
    public Optional<Output<String>> issuer() {
        return Optional.ofNullable(this.issuer);
    }

    /**
     * The namespace to provision the resource in.
     * The value should not contain leading or trailing forward slashes.
     * The `namespace` is always relative to the provider&#39;s configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
     * *Available only for Vault Enterprise*.
     * 
     */
    @Import(name="namespace")
    private @Nullable Output<String> namespace;

    /**
     * @return The namespace to provision the resource in.
     * The value should not contain leading or trailing forward slashes.
     * The `namespace` is always relative to the provider&#39;s configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
     * *Available only for Vault Enterprise*.
     * 
     */
    public Optional<Output<String>> namespace() {
        return Optional.ofNullable(this.namespace);
    }

    private OidcArgs() {}

    private OidcArgs(OidcArgs $) {
        this.issuer = $.issuer;
        this.namespace = $.namespace;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(OidcArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private OidcArgs $;

        public Builder() {
            $ = new OidcArgs();
        }

        public Builder(OidcArgs defaults) {
            $ = new OidcArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param issuer Issuer URL to be used in the iss claim of the token. If not set, Vault&#39;s
         * `api_addr` will be used. The issuer is a case sensitive URL using the https scheme that contains
         * scheme, host, and optionally, port number and path components, but no query or fragment
         * components.
         * 
         * @return builder
         * 
         */
        public Builder issuer(@Nullable Output<String> issuer) {
            $.issuer = issuer;
            return this;
        }

        /**
         * @param issuer Issuer URL to be used in the iss claim of the token. If not set, Vault&#39;s
         * `api_addr` will be used. The issuer is a case sensitive URL using the https scheme that contains
         * scheme, host, and optionally, port number and path components, but no query or fragment
         * components.
         * 
         * @return builder
         * 
         */
        public Builder issuer(String issuer) {
            return issuer(Output.of(issuer));
        }

        /**
         * @param namespace The namespace to provision the resource in.
         * The value should not contain leading or trailing forward slashes.
         * The `namespace` is always relative to the provider&#39;s configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
         * *Available only for Vault Enterprise*.
         * 
         * @return builder
         * 
         */
        public Builder namespace(@Nullable Output<String> namespace) {
            $.namespace = namespace;
            return this;
        }

        /**
         * @param namespace The namespace to provision the resource in.
         * The value should not contain leading or trailing forward slashes.
         * The `namespace` is always relative to the provider&#39;s configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
         * *Available only for Vault Enterprise*.
         * 
         * @return builder
         * 
         */
        public Builder namespace(String namespace) {
            return namespace(Output.of(namespace));
        }

        public OidcArgs build() {
            return $;
        }
    }

}