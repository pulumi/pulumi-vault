// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vault.terraformcloud;

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
     * Specifies the address of the Terraform Cloud instance, provided as &#34;host:port&#34; like &#34;127.0.0.1:8500&#34;.
     * 
     */
    @Import(name="address")
    private @Nullable Output<String> address;

    /**
     * @return Specifies the address of the Terraform Cloud instance, provided as &#34;host:port&#34; like &#34;127.0.0.1:8500&#34;.
     * 
     */
    public Optional<Output<String>> address() {
        return Optional.ofNullable(this.address);
    }

    /**
     * Unique name of the Vault Terraform Cloud mount to configure
     * 
     */
    @Import(name="backend")
    private @Nullable Output<String> backend;

    /**
     * @return Unique name of the Vault Terraform Cloud mount to configure
     * 
     */
    public Optional<Output<String>> backend() {
        return Optional.ofNullable(this.backend);
    }

    /**
     * Specifies the base path for the Terraform Cloud or Enterprise API.
     * 
     */
    @Import(name="basePath")
    private @Nullable Output<String> basePath;

    /**
     * @return Specifies the base path for the Terraform Cloud or Enterprise API.
     * 
     */
    public Optional<Output<String>> basePath() {
        return Optional.ofNullable(this.basePath);
    }

    /**
     * The default TTL for credentials issued by this backend.
     * 
     */
    @Import(name="defaultLeaseTtlSeconds")
    private @Nullable Output<Integer> defaultLeaseTtlSeconds;

    /**
     * @return The default TTL for credentials issued by this backend.
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
     * The maximum TTL that can be requested
     * for credentials issued by this backend.
     * 
     */
    @Import(name="maxLeaseTtlSeconds")
    private @Nullable Output<Integer> maxLeaseTtlSeconds;

    /**
     * @return The maximum TTL that can be requested
     * for credentials issued by this backend.
     * 
     */
    public Optional<Output<Integer>> maxLeaseTtlSeconds() {
        return Optional.ofNullable(this.maxLeaseTtlSeconds);
    }

    /**
     * The namespace to provision the resource in.
     * The value should not contain leading or trailing forward slashes.
     * The `namespace` is always relative to the provider&#39;s configured namespace.
     * *Available only for Vault Enterprise*.
     * 
     */
    @Import(name="namespace")
    private @Nullable Output<String> namespace;

    /**
     * @return The namespace to provision the resource in.
     * The value should not contain leading or trailing forward slashes.
     * The `namespace` is always relative to the provider&#39;s configured namespace.
     * *Available only for Vault Enterprise*.
     * 
     */
    public Optional<Output<String>> namespace() {
        return Optional.ofNullable(this.namespace);
    }

    /**
     * Specifies the Terraform Cloud access token to use.
     * 
     */
    @Import(name="token")
    private @Nullable Output<String> token;

    /**
     * @return Specifies the Terraform Cloud access token to use.
     * 
     */
    public Optional<Output<String>> token() {
        return Optional.ofNullable(this.token);
    }

    private SecretBackendArgs() {}

    private SecretBackendArgs(SecretBackendArgs $) {
        this.address = $.address;
        this.backend = $.backend;
        this.basePath = $.basePath;
        this.defaultLeaseTtlSeconds = $.defaultLeaseTtlSeconds;
        this.description = $.description;
        this.disableRemount = $.disableRemount;
        this.maxLeaseTtlSeconds = $.maxLeaseTtlSeconds;
        this.namespace = $.namespace;
        this.token = $.token;
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
         * @param address Specifies the address of the Terraform Cloud instance, provided as &#34;host:port&#34; like &#34;127.0.0.1:8500&#34;.
         * 
         * @return builder
         * 
         */
        public Builder address(@Nullable Output<String> address) {
            $.address = address;
            return this;
        }

        /**
         * @param address Specifies the address of the Terraform Cloud instance, provided as &#34;host:port&#34; like &#34;127.0.0.1:8500&#34;.
         * 
         * @return builder
         * 
         */
        public Builder address(String address) {
            return address(Output.of(address));
        }

        /**
         * @param backend Unique name of the Vault Terraform Cloud mount to configure
         * 
         * @return builder
         * 
         */
        public Builder backend(@Nullable Output<String> backend) {
            $.backend = backend;
            return this;
        }

        /**
         * @param backend Unique name of the Vault Terraform Cloud mount to configure
         * 
         * @return builder
         * 
         */
        public Builder backend(String backend) {
            return backend(Output.of(backend));
        }

        /**
         * @param basePath Specifies the base path for the Terraform Cloud or Enterprise API.
         * 
         * @return builder
         * 
         */
        public Builder basePath(@Nullable Output<String> basePath) {
            $.basePath = basePath;
            return this;
        }

        /**
         * @param basePath Specifies the base path for the Terraform Cloud or Enterprise API.
         * 
         * @return builder
         * 
         */
        public Builder basePath(String basePath) {
            return basePath(Output.of(basePath));
        }

        /**
         * @param defaultLeaseTtlSeconds The default TTL for credentials issued by this backend.
         * 
         * @return builder
         * 
         */
        public Builder defaultLeaseTtlSeconds(@Nullable Output<Integer> defaultLeaseTtlSeconds) {
            $.defaultLeaseTtlSeconds = defaultLeaseTtlSeconds;
            return this;
        }

        /**
         * @param defaultLeaseTtlSeconds The default TTL for credentials issued by this backend.
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
         * @param maxLeaseTtlSeconds The maximum TTL that can be requested
         * for credentials issued by this backend.
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
         * for credentials issued by this backend.
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
         * The `namespace` is always relative to the provider&#39;s configured namespace.
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
         * The `namespace` is always relative to the provider&#39;s configured namespace.
         * *Available only for Vault Enterprise*.
         * 
         * @return builder
         * 
         */
        public Builder namespace(String namespace) {
            return namespace(Output.of(namespace));
        }

        /**
         * @param token Specifies the Terraform Cloud access token to use.
         * 
         * @return builder
         * 
         */
        public Builder token(@Nullable Output<String> token) {
            $.token = token;
            return this;
        }

        /**
         * @param token Specifies the Terraform Cloud access token to use.
         * 
         * @return builder
         * 
         */
        public Builder token(String token) {
            return token(Output.of(token));
        }

        public SecretBackendArgs build() {
            return $;
        }
    }

}