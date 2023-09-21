// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vault.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class QuotaLeaseCountState extends com.pulumi.resources.ResourceArgs {

    public static final QuotaLeaseCountState Empty = new QuotaLeaseCountState();

    /**
     * The maximum number of leases to be allowed by the quota
     * rule. The `max_leases` must be positive.
     * 
     */
    @Import(name="maxLeases")
    private @Nullable Output<Integer> maxLeases;

    /**
     * @return The maximum number of leases to be allowed by the quota
     * rule. The `max_leases` must be positive.
     * 
     */
    public Optional<Output<Integer>> maxLeases() {
        return Optional.ofNullable(this.maxLeases);
    }

    /**
     * Name of the rate limit quota
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return Name of the rate limit quota
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
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
     * Path of the mount or namespace to apply the quota. A blank path configures a
     * global rate limit quota. For example `namespace1/` adds a quota to a full namespace,
     * `namespace1/auth/userpass` adds a `quota` to `userpass` in `namespace1`.
     * Updating this field on an existing quota can have &#34;moving&#34; effects. For example, updating
     * `auth/userpass` to `namespace1/auth/userpass` moves this quota from being a global mount quota to
     * a namespace specific mount quota. **Note, namespaces are supported in Enterprise only.**
     * 
     */
    @Import(name="path")
    private @Nullable Output<String> path;

    /**
     * @return Path of the mount or namespace to apply the quota. A blank path configures a
     * global rate limit quota. For example `namespace1/` adds a quota to a full namespace,
     * `namespace1/auth/userpass` adds a `quota` to `userpass` in `namespace1`.
     * Updating this field on an existing quota can have &#34;moving&#34; effects. For example, updating
     * `auth/userpass` to `namespace1/auth/userpass` moves this quota from being a global mount quota to
     * a namespace specific mount quota. **Note, namespaces are supported in Enterprise only.**
     * 
     */
    public Optional<Output<String>> path() {
        return Optional.ofNullable(this.path);
    }

    /**
     * If set on a quota where `path` is set to an auth mount with a concept of roles (such as /auth/approle/), this will make the quota restrict login requests to that mount that are made with the specified role.
     * 
     */
    @Import(name="role")
    private @Nullable Output<String> role;

    /**
     * @return If set on a quota where `path` is set to an auth mount with a concept of roles (such as /auth/approle/), this will make the quota restrict login requests to that mount that are made with the specified role.
     * 
     */
    public Optional<Output<String>> role() {
        return Optional.ofNullable(this.role);
    }

    private QuotaLeaseCountState() {}

    private QuotaLeaseCountState(QuotaLeaseCountState $) {
        this.maxLeases = $.maxLeases;
        this.name = $.name;
        this.namespace = $.namespace;
        this.path = $.path;
        this.role = $.role;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(QuotaLeaseCountState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private QuotaLeaseCountState $;

        public Builder() {
            $ = new QuotaLeaseCountState();
        }

        public Builder(QuotaLeaseCountState defaults) {
            $ = new QuotaLeaseCountState(Objects.requireNonNull(defaults));
        }

        /**
         * @param maxLeases The maximum number of leases to be allowed by the quota
         * rule. The `max_leases` must be positive.
         * 
         * @return builder
         * 
         */
        public Builder maxLeases(@Nullable Output<Integer> maxLeases) {
            $.maxLeases = maxLeases;
            return this;
        }

        /**
         * @param maxLeases The maximum number of leases to be allowed by the quota
         * rule. The `max_leases` must be positive.
         * 
         * @return builder
         * 
         */
        public Builder maxLeases(Integer maxLeases) {
            return maxLeases(Output.of(maxLeases));
        }

        /**
         * @param name Name of the rate limit quota
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name Name of the rate limit quota
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
         * @param path Path of the mount or namespace to apply the quota. A blank path configures a
         * global rate limit quota. For example `namespace1/` adds a quota to a full namespace,
         * `namespace1/auth/userpass` adds a `quota` to `userpass` in `namespace1`.
         * Updating this field on an existing quota can have &#34;moving&#34; effects. For example, updating
         * `auth/userpass` to `namespace1/auth/userpass` moves this quota from being a global mount quota to
         * a namespace specific mount quota. **Note, namespaces are supported in Enterprise only.**
         * 
         * @return builder
         * 
         */
        public Builder path(@Nullable Output<String> path) {
            $.path = path;
            return this;
        }

        /**
         * @param path Path of the mount or namespace to apply the quota. A blank path configures a
         * global rate limit quota. For example `namespace1/` adds a quota to a full namespace,
         * `namespace1/auth/userpass` adds a `quota` to `userpass` in `namespace1`.
         * Updating this field on an existing quota can have &#34;moving&#34; effects. For example, updating
         * `auth/userpass` to `namespace1/auth/userpass` moves this quota from being a global mount quota to
         * a namespace specific mount quota. **Note, namespaces are supported in Enterprise only.**
         * 
         * @return builder
         * 
         */
        public Builder path(String path) {
            return path(Output.of(path));
        }

        /**
         * @param role If set on a quota where `path` is set to an auth mount with a concept of roles (such as /auth/approle/), this will make the quota restrict login requests to that mount that are made with the specified role.
         * 
         * @return builder
         * 
         */
        public Builder role(@Nullable Output<String> role) {
            $.role = role;
            return this;
        }

        /**
         * @param role If set on a quota where `path` is set to an auth mount with a concept of roles (such as /auth/approle/), this will make the quota restrict login requests to that mount that are made with the specified role.
         * 
         * @return builder
         * 
         */
        public Builder role(String role) {
            return role(Output.of(role));
        }

        public QuotaLeaseCountState build() {
            return $;
        }
    }

}
