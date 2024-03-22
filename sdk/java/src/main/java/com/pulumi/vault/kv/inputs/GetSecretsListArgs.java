// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vault.kv.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetSecretsListArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetSecretsListArgs Empty = new GetSecretsListArgs();

    /**
     * The namespace of the target resource.
     * The value should not contain leading or trailing forward slashes.
     * The `namespace` is always relative to the provider&#39;s configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
     * *Available only for Vault Enterprise*.
     * 
     */
    @Import(name="namespace")
    private @Nullable Output<String> namespace;

    /**
     * @return The namespace of the target resource.
     * The value should not contain leading or trailing forward slashes.
     * The `namespace` is always relative to the provider&#39;s configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
     * *Available only for Vault Enterprise*.
     * 
     */
    public Optional<Output<String>> namespace() {
        return Optional.ofNullable(this.namespace);
    }

    /**
     * Full KV-V1 path where secrets will be listed.
     * 
     */
    @Import(name="path", required=true)
    private Output<String> path;

    /**
     * @return Full KV-V1 path where secrets will be listed.
     * 
     */
    public Output<String> path() {
        return this.path;
    }

    private GetSecretsListArgs() {}

    private GetSecretsListArgs(GetSecretsListArgs $) {
        this.namespace = $.namespace;
        this.path = $.path;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetSecretsListArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetSecretsListArgs $;

        public Builder() {
            $ = new GetSecretsListArgs();
        }

        public Builder(GetSecretsListArgs defaults) {
            $ = new GetSecretsListArgs(Objects.requireNonNull(defaults));
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
        public Builder namespace(@Nullable Output<String> namespace) {
            $.namespace = namespace;
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
        public Builder namespace(String namespace) {
            return namespace(Output.of(namespace));
        }

        /**
         * @param path Full KV-V1 path where secrets will be listed.
         * 
         * @return builder
         * 
         */
        public Builder path(Output<String> path) {
            $.path = path;
            return this;
        }

        /**
         * @param path Full KV-V1 path where secrets will be listed.
         * 
         * @return builder
         * 
         */
        public Builder path(String path) {
            return path(Output.of(path));
        }

        public GetSecretsListArgs build() {
            if ($.path == null) {
                throw new MissingRequiredPropertyException("GetSecretsListArgs", "path");
            }
            return $;
        }
    }

}
