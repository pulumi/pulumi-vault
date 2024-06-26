// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vault.inputs;

import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetNamespacePlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetNamespacePlainArgs Empty = new GetNamespacePlainArgs();

    /**
     * The namespace to provision the resource in.
     * The value should not contain leading or trailing forward slashes.
     * The `namespace` is always relative to the provider&#39;s configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
     * 
     */
    @Import(name="namespace")
    private @Nullable String namespace;

    /**
     * @return The namespace to provision the resource in.
     * The value should not contain leading or trailing forward slashes.
     * The `namespace` is always relative to the provider&#39;s configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
     * 
     */
    public Optional<String> namespace() {
        return Optional.ofNullable(this.namespace);
    }

    /**
     * The path of the namespace. Must not have a trailing `/`.
     * If not specified or empty, path attributes are set for the current namespace
     * based on the `namespace` arguments of the provider and this data source.
     * Other path related attributes will be empty in this case.
     * 
     */
    @Import(name="path")
    private @Nullable String path;

    /**
     * @return The path of the namespace. Must not have a trailing `/`.
     * If not specified or empty, path attributes are set for the current namespace
     * based on the `namespace` arguments of the provider and this data source.
     * Other path related attributes will be empty in this case.
     * 
     */
    public Optional<String> path() {
        return Optional.ofNullable(this.path);
    }

    private GetNamespacePlainArgs() {}

    private GetNamespacePlainArgs(GetNamespacePlainArgs $) {
        this.namespace = $.namespace;
        this.path = $.path;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetNamespacePlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetNamespacePlainArgs $;

        public Builder() {
            $ = new GetNamespacePlainArgs();
        }

        public Builder(GetNamespacePlainArgs defaults) {
            $ = new GetNamespacePlainArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param namespace The namespace to provision the resource in.
         * The value should not contain leading or trailing forward slashes.
         * The `namespace` is always relative to the provider&#39;s configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
         * 
         * @return builder
         * 
         */
        public Builder namespace(@Nullable String namespace) {
            $.namespace = namespace;
            return this;
        }

        /**
         * @param path The path of the namespace. Must not have a trailing `/`.
         * If not specified or empty, path attributes are set for the current namespace
         * based on the `namespace` arguments of the provider and this data source.
         * Other path related attributes will be empty in this case.
         * 
         * @return builder
         * 
         */
        public Builder path(@Nullable String path) {
            $.path = path;
            return this;
        }

        public GetNamespacePlainArgs build() {
            return $;
        }
    }

}
