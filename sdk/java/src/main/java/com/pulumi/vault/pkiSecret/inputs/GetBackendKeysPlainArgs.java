// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vault.pkiSecret.inputs;

import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetBackendKeysPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetBackendKeysPlainArgs Empty = new GetBackendKeysPlainArgs();

    /**
     * The path to the PKI secret backend to
     * read the keys from, with no leading or trailing `/`s.
     * 
     */
    @Import(name="backend", required=true)
    private String backend;

    /**
     * @return The path to the PKI secret backend to
     * read the keys from, with no leading or trailing `/`s.
     * 
     */
    public String backend() {
        return this.backend;
    }

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

    private GetBackendKeysPlainArgs() {}

    private GetBackendKeysPlainArgs(GetBackendKeysPlainArgs $) {
        this.backend = $.backend;
        this.namespace = $.namespace;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetBackendKeysPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetBackendKeysPlainArgs $;

        public Builder() {
            $ = new GetBackendKeysPlainArgs();
        }

        public Builder(GetBackendKeysPlainArgs defaults) {
            $ = new GetBackendKeysPlainArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param backend The path to the PKI secret backend to
         * read the keys from, with no leading or trailing `/`s.
         * 
         * @return builder
         * 
         */
        public Builder backend(String backend) {
            $.backend = backend;
            return this;
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

        public GetBackendKeysPlainArgs build() {
            if ($.backend == null) {
                throw new MissingRequiredPropertyException("GetBackendKeysPlainArgs", "backend");
            }
            return $;
        }
    }

}
