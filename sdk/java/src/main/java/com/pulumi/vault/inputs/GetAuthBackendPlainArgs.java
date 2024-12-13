// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vault.inputs;

import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetAuthBackendPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetAuthBackendPlainArgs Empty = new GetAuthBackendPlainArgs();

    /**
     * The namespace of the target resource.
     * The value should not contain leading or trailing forward slashes.
     * The `namespace` is always relative to the provider&#39;s configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
     * *Available only for Vault Enterprise*.
     * 
     */
    @Import(name="namespace")
    private @Nullable String namespace;

    /**
     * @return The namespace of the target resource.
     * The value should not contain leading or trailing forward slashes.
     * The `namespace` is always relative to the provider&#39;s configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
     * *Available only for Vault Enterprise*.
     * 
     */
    public Optional<String> namespace() {
        return Optional.ofNullable(this.namespace);
    }

    /**
     * The auth backend mount point.
     * 
     */
    @Import(name="path", required=true)
    private String path;

    /**
     * @return The auth backend mount point.
     * 
     */
    public String path() {
        return this.path;
    }

    private GetAuthBackendPlainArgs() {}

    private GetAuthBackendPlainArgs(GetAuthBackendPlainArgs $) {
        this.namespace = $.namespace;
        this.path = $.path;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetAuthBackendPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetAuthBackendPlainArgs $;

        public Builder() {
            $ = new GetAuthBackendPlainArgs();
        }

        public Builder(GetAuthBackendPlainArgs defaults) {
            $ = new GetAuthBackendPlainArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param namespace The namespace of the target resource.
         * The value should not contain leading or trailing forward slashes.
         * The `namespace` is always relative to the provider&#39;s configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
         * *Available only for Vault Enterprise*.
         * 
         * @return builder
         * 
         */
        public Builder namespace(@Nullable String namespace) {
            $.namespace = namespace;
            return this;
        }

        /**
         * @param path The auth backend mount point.
         * 
         * @return builder
         * 
         */
        public Builder path(String path) {
            $.path = path;
            return this;
        }

        public GetAuthBackendPlainArgs build() {
            if ($.path == null) {
                throw new MissingRequiredPropertyException("GetAuthBackendPlainArgs", "path");
            }
            return $;
        }
    }

}