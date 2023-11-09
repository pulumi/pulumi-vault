// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vault.gcp;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class SecretBackendArgs extends com.pulumi.resources.ResourceArgs {

    public static final SecretBackendArgs Empty = new SecretBackendArgs();

    /**
     * JSON-encoded credentials to use to connect to GCP
     * 
     */
    @Import(name="credentials")
    private @Nullable Output<String> credentials;

    /**
     * @return JSON-encoded credentials to use to connect to GCP
     * 
     */
    public Optional<Output<String>> credentials() {
        return Optional.ofNullable(this.credentials);
    }

    /**
     * The default TTL for credentials
     * issued by this backend. Defaults to &#39;0&#39;.
     * 
     */
    @Import(name="defaultLeaseTtlSeconds")
    private @Nullable Output<Integer> defaultLeaseTtlSeconds;

    /**
     * @return The default TTL for credentials
     * issued by this backend. Defaults to &#39;0&#39;.
     * 
     */
    public Optional<Output<Integer>> defaultLeaseTtlSeconds() {
        return Optional.ofNullable(this.defaultLeaseTtlSeconds);
    }

    /**
     * A human-friendly description for this backend.
     * 
     */
    @Import(name="description")
    private @Nullable Output<String> description;

    /**
     * @return A human-friendly description for this backend.
     * 
     */
    public Optional<Output<String>> description() {
        return Optional.ofNullable(this.description);
    }

    /**
     * If set, opts out of mount migration on path updates.
     * See here for more info on [Mount Migration](https://www.vaultproject.io/docs/concepts/mount-migration)
     * 
     */
    @Import(name="disableRemount")
    private @Nullable Output<Boolean> disableRemount;

    /**
     * @return If set, opts out of mount migration on path updates.
     * See here for more info on [Mount Migration](https://www.vaultproject.io/docs/concepts/mount-migration)
     * 
     */
    public Optional<Output<Boolean>> disableRemount() {
        return Optional.ofNullable(this.disableRemount);
    }

    /**
     * Boolean flag that can be explicitly set to true to enforce local mount in HA environment
     * 
     */
    @Import(name="local")
    private @Nullable Output<Boolean> local;

    /**
     * @return Boolean flag that can be explicitly set to true to enforce local mount in HA environment
     * 
     */
    public Optional<Output<Boolean>> local() {
        return Optional.ofNullable(this.local);
    }

    /**
     * The maximum TTL that can be requested
     * for credentials issued by this backend. Defaults to &#39;0&#39;.
     * 
     */
    @Import(name="maxLeaseTtlSeconds")
    private @Nullable Output<Integer> maxLeaseTtlSeconds;

    /**
     * @return The maximum TTL that can be requested
     * for credentials issued by this backend. Defaults to &#39;0&#39;.
     * 
     */
    public Optional<Output<Integer>> maxLeaseTtlSeconds() {
        return Optional.ofNullable(this.maxLeaseTtlSeconds);
    }

    /**
     * The namespace to provision the resource in.
     * The value should not contain leading or trailing forward slashes.
     * The `namespace` is always relative to the provider&#39;s configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
     * *Available only for Vault Enterprise*.
     * 
     */
    @Import(name="namespace")
    private @Nullable Output<String> namespace;

    /**
     * @return The namespace to provision the resource in.
     * The value should not contain leading or trailing forward slashes.
     * The `namespace` is always relative to the provider&#39;s configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
     * *Available only for Vault Enterprise*.
     * 
     */
    public Optional<Output<String>> namespace() {
        return Optional.ofNullable(this.namespace);
    }

    /**
     * The unique path this backend should be mounted at. Must
     * not begin or end with a `/`. Defaults to `gcp`.
     * 
     */
    @Import(name="path")
    private @Nullable Output<String> path;

    /**
     * @return The unique path this backend should be mounted at. Must
     * not begin or end with a `/`. Defaults to `gcp`.
     * 
     */
    public Optional<Output<String>> path() {
        return Optional.ofNullable(this.path);
    }

    private SecretBackendArgs() {}

    private SecretBackendArgs(SecretBackendArgs $) {
        this.credentials = $.credentials;
        this.defaultLeaseTtlSeconds = $.defaultLeaseTtlSeconds;
        this.description = $.description;
        this.disableRemount = $.disableRemount;
        this.local = $.local;
        this.maxLeaseTtlSeconds = $.maxLeaseTtlSeconds;
        this.namespace = $.namespace;
        this.path = $.path;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(SecretBackendArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private SecretBackendArgs $;

        public Builder() {
            $ = new SecretBackendArgs();
        }

        public Builder(SecretBackendArgs defaults) {
            $ = new SecretBackendArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param credentials JSON-encoded credentials to use to connect to GCP
         * 
         * @return builder
         * 
         */
        public Builder credentials(@Nullable Output<String> credentials) {
            $.credentials = credentials;
            return this;
        }

        /**
         * @param credentials JSON-encoded credentials to use to connect to GCP
         * 
         * @return builder
         * 
         */
        public Builder credentials(String credentials) {
            return credentials(Output.of(credentials));
        }

        /**
         * @param defaultLeaseTtlSeconds The default TTL for credentials
         * issued by this backend. Defaults to &#39;0&#39;.
         * 
         * @return builder
         * 
         */
        public Builder defaultLeaseTtlSeconds(@Nullable Output<Integer> defaultLeaseTtlSeconds) {
            $.defaultLeaseTtlSeconds = defaultLeaseTtlSeconds;
            return this;
        }

        /**
         * @param defaultLeaseTtlSeconds The default TTL for credentials
         * issued by this backend. Defaults to &#39;0&#39;.
         * 
         * @return builder
         * 
         */
        public Builder defaultLeaseTtlSeconds(Integer defaultLeaseTtlSeconds) {
            return defaultLeaseTtlSeconds(Output.of(defaultLeaseTtlSeconds));
        }

        /**
         * @param description A human-friendly description for this backend.
         * 
         * @return builder
         * 
         */
        public Builder description(@Nullable Output<String> description) {
            $.description = description;
            return this;
        }

        /**
         * @param description A human-friendly description for this backend.
         * 
         * @return builder
         * 
         */
        public Builder description(String description) {
            return description(Output.of(description));
        }

        /**
         * @param disableRemount If set, opts out of mount migration on path updates.
         * See here for more info on [Mount Migration](https://www.vaultproject.io/docs/concepts/mount-migration)
         * 
         * @return builder
         * 
         */
        public Builder disableRemount(@Nullable Output<Boolean> disableRemount) {
            $.disableRemount = disableRemount;
            return this;
        }

        /**
         * @param disableRemount If set, opts out of mount migration on path updates.
         * See here for more info on [Mount Migration](https://www.vaultproject.io/docs/concepts/mount-migration)
         * 
         * @return builder
         * 
         */
        public Builder disableRemount(Boolean disableRemount) {
            return disableRemount(Output.of(disableRemount));
        }

        /**
         * @param local Boolean flag that can be explicitly set to true to enforce local mount in HA environment
         * 
         * @return builder
         * 
         */
        public Builder local(@Nullable Output<Boolean> local) {
            $.local = local;
            return this;
        }

        /**
         * @param local Boolean flag that can be explicitly set to true to enforce local mount in HA environment
         * 
         * @return builder
         * 
         */
        public Builder local(Boolean local) {
            return local(Output.of(local));
        }

        /**
         * @param maxLeaseTtlSeconds The maximum TTL that can be requested
         * for credentials issued by this backend. Defaults to &#39;0&#39;.
         * 
         * @return builder
         * 
         */
        public Builder maxLeaseTtlSeconds(@Nullable Output<Integer> maxLeaseTtlSeconds) {
            $.maxLeaseTtlSeconds = maxLeaseTtlSeconds;
            return this;
        }

        /**
         * @param maxLeaseTtlSeconds The maximum TTL that can be requested
         * for credentials issued by this backend. Defaults to &#39;0&#39;.
         * 
         * @return builder
         * 
         */
        public Builder maxLeaseTtlSeconds(Integer maxLeaseTtlSeconds) {
            return maxLeaseTtlSeconds(Output.of(maxLeaseTtlSeconds));
        }

        /**
         * @param namespace The namespace to provision the resource in.
         * The value should not contain leading or trailing forward slashes.
         * The `namespace` is always relative to the provider&#39;s configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
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
         * The `namespace` is always relative to the provider&#39;s configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
         * *Available only for Vault Enterprise*.
         * 
         * @return builder
         * 
         */
        public Builder namespace(String namespace) {
            return namespace(Output.of(namespace));
        }

        /**
         * @param path The unique path this backend should be mounted at. Must
         * not begin or end with a `/`. Defaults to `gcp`.
         * 
         * @return builder
         * 
         */
        public Builder path(@Nullable Output<String> path) {
            $.path = path;
            return this;
        }

        /**
         * @param path The unique path this backend should be mounted at. Must
         * not begin or end with a `/`. Defaults to `gcp`.
         * 
         * @return builder
         * 
         */
        public Builder path(String path) {
            return path(Output.of(path));
        }

        public SecretBackendArgs build() {
            return $;
        }
    }

}
