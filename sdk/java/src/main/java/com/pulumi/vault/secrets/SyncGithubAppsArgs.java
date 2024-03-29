// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vault.secrets;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class SyncGithubAppsArgs extends com.pulumi.resources.ResourceArgs {

    public static final SyncGithubAppsArgs Empty = new SyncGithubAppsArgs();

    /**
     * The GitHub application ID.
     * 
     */
    @Import(name="appId", required=true)
    private Output<Integer> appId;

    /**
     * @return The GitHub application ID.
     * 
     */
    public Output<Integer> appId() {
        return this.appId;
    }

    /**
     * The user-defined name of the GitHub App configuration.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return The user-defined name of the GitHub App configuration.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * The namespace to provision the resource in.
     * The value should not contain leading or trailing forward slashes.
     * The `namespace` is always relative to the provider&#39;s configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
     * 
     */
    @Import(name="namespace")
    private @Nullable Output<String> namespace;

    /**
     * @return The namespace to provision the resource in.
     * The value should not contain leading or trailing forward slashes.
     * The `namespace` is always relative to the provider&#39;s configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
     * 
     */
    public Optional<Output<String>> namespace() {
        return Optional.ofNullable(this.namespace);
    }

    /**
     * The content of a PEM formatted private key generated on GitHub for the app.
     * 
     */
    @Import(name="privateKey", required=true)
    private Output<String> privateKey;

    /**
     * @return The content of a PEM formatted private key generated on GitHub for the app.
     * 
     */
    public Output<String> privateKey() {
        return this.privateKey;
    }

    private SyncGithubAppsArgs() {}

    private SyncGithubAppsArgs(SyncGithubAppsArgs $) {
        this.appId = $.appId;
        this.name = $.name;
        this.namespace = $.namespace;
        this.privateKey = $.privateKey;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(SyncGithubAppsArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private SyncGithubAppsArgs $;

        public Builder() {
            $ = new SyncGithubAppsArgs();
        }

        public Builder(SyncGithubAppsArgs defaults) {
            $ = new SyncGithubAppsArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param appId The GitHub application ID.
         * 
         * @return builder
         * 
         */
        public Builder appId(Output<Integer> appId) {
            $.appId = appId;
            return this;
        }

        /**
         * @param appId The GitHub application ID.
         * 
         * @return builder
         * 
         */
        public Builder appId(Integer appId) {
            return appId(Output.of(appId));
        }

        /**
         * @param name The user-defined name of the GitHub App configuration.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name The user-defined name of the GitHub App configuration.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param namespace The namespace to provision the resource in.
         * The value should not contain leading or trailing forward slashes.
         * The `namespace` is always relative to the provider&#39;s configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
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
         * 
         * @return builder
         * 
         */
        public Builder namespace(String namespace) {
            return namespace(Output.of(namespace));
        }

        /**
         * @param privateKey The content of a PEM formatted private key generated on GitHub for the app.
         * 
         * @return builder
         * 
         */
        public Builder privateKey(Output<String> privateKey) {
            $.privateKey = privateKey;
            return this;
        }

        /**
         * @param privateKey The content of a PEM formatted private key generated on GitHub for the app.
         * 
         * @return builder
         * 
         */
        public Builder privateKey(String privateKey) {
            return privateKey(Output.of(privateKey));
        }

        public SyncGithubAppsArgs build() {
            if ($.appId == null) {
                throw new MissingRequiredPropertyException("SyncGithubAppsArgs", "appId");
            }
            if ($.privateKey == null) {
                throw new MissingRequiredPropertyException("SyncGithubAppsArgs", "privateKey");
            }
            return $;
        }
    }

}
