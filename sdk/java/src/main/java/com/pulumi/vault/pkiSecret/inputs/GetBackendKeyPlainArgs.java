// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vault.pkiSecret.inputs;

import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetBackendKeyPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetBackendKeyPlainArgs Empty = new GetBackendKeyPlainArgs();

    /**
     * The path to the PKI secret backend to
     * read the key from, with no leading or trailing `/`s.
     * 
     */
    @Import(name="backend", required=true)
    private String backend;

    /**
     * @return The path to the PKI secret backend to
     * read the key from, with no leading or trailing `/`s.
     * 
     */
    public String backend() {
        return this.backend;
    }

    /**
     * Reference to an existing key.
     * 
     */
    @Import(name="keyRef", required=true)
    private String keyRef;

    /**
     * @return Reference to an existing key.
     * 
     */
    public String keyRef() {
        return this.keyRef;
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

    private GetBackendKeyPlainArgs() {}

    private GetBackendKeyPlainArgs(GetBackendKeyPlainArgs $) {
        this.backend = $.backend;
        this.keyRef = $.keyRef;
        this.namespace = $.namespace;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetBackendKeyPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetBackendKeyPlainArgs $;

        public Builder() {
            $ = new GetBackendKeyPlainArgs();
        }

        public Builder(GetBackendKeyPlainArgs defaults) {
            $ = new GetBackendKeyPlainArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param backend The path to the PKI secret backend to
         * read the key from, with no leading or trailing `/`s.
         * 
         * @return builder
         * 
         */
        public Builder backend(String backend) {
            $.backend = backend;
            return this;
        }

        /**
         * @param keyRef Reference to an existing key.
         * 
         * @return builder
         * 
         */
        public Builder keyRef(String keyRef) {
            $.keyRef = keyRef;
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

        public GetBackendKeyPlainArgs build() {
            if ($.backend == null) {
                throw new MissingRequiredPropertyException("GetBackendKeyPlainArgs", "backend");
            }
            if ($.keyRef == null) {
                throw new MissingRequiredPropertyException("GetBackendKeyPlainArgs", "keyRef");
            }
            return $;
        }
    }

}
