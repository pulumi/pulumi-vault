// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vault;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Boolean;
import java.lang.Double;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class QuotaRateLimitArgs extends com.pulumi.resources.ResourceArgs {

    public static final QuotaRateLimitArgs Empty = new QuotaRateLimitArgs();

    /**
     * If set, when a client reaches a rate limit threshold, the client will
     * be prohibited from any further requests until after the &#39;block_interval&#39; in seconds has elapsed.
     * 
     */
    @Import(name="blockInterval")
    private @Nullable Output<Integer> blockInterval;

    /**
     * @return If set, when a client reaches a rate limit threshold, the client will
     * be prohibited from any further requests until after the &#39;block_interval&#39; in seconds has elapsed.
     * 
     */
    public Optional<Output<Integer>> blockInterval() {
        return Optional.ofNullable(this.blockInterval);
    }

    /**
     * If set to `true` on a quota where path is set to a namespace, the same quota will be cumulatively applied to all child namespace. The inheritable parameter cannot be set to `true` if the path does not specify a namespace. Only the quotas associated with the root namespace are inheritable by default. Requires Vault 1.15+.
     * 
     */
    @Import(name="inheritable")
    private @Nullable Output<Boolean> inheritable;

    /**
     * @return If set to `true` on a quota where path is set to a namespace, the same quota will be cumulatively applied to all child namespace. The inheritable parameter cannot be set to `true` if the path does not specify a namespace. Only the quotas associated with the root namespace are inheritable by default. Requires Vault 1.15+.
     * 
     */
    public Optional<Output<Boolean>> inheritable() {
        return Optional.ofNullable(this.inheritable);
    }

    /**
     * The duration in seconds to enforce rate limiting for.
     * 
     */
    @Import(name="interval")
    private @Nullable Output<Integer> interval;

    /**
     * @return The duration in seconds to enforce rate limiting for.
     * 
     */
    public Optional<Output<Integer>> interval() {
        return Optional.ofNullable(this.interval);
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
     * The maximum number of requests at any given second to be allowed by the quota
     * rule. The `rate` must be positive.
     * 
     */
    @Import(name="rate", required=true)
    private Output<Double> rate;

    /**
     * @return The maximum number of requests at any given second to be allowed by the quota
     * rule. The `rate` must be positive.
     * 
     */
    public Output<Double> rate() {
        return this.rate;
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

    private QuotaRateLimitArgs() {}

    private QuotaRateLimitArgs(QuotaRateLimitArgs $) {
        this.blockInterval = $.blockInterval;
        this.inheritable = $.inheritable;
        this.interval = $.interval;
        this.name = $.name;
        this.namespace = $.namespace;
        this.path = $.path;
        this.rate = $.rate;
        this.role = $.role;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(QuotaRateLimitArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private QuotaRateLimitArgs $;

        public Builder() {
            $ = new QuotaRateLimitArgs();
        }

        public Builder(QuotaRateLimitArgs defaults) {
            $ = new QuotaRateLimitArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param blockInterval If set, when a client reaches a rate limit threshold, the client will
         * be prohibited from any further requests until after the &#39;block_interval&#39; in seconds has elapsed.
         * 
         * @return builder
         * 
         */
        public Builder blockInterval(@Nullable Output<Integer> blockInterval) {
            $.blockInterval = blockInterval;
            return this;
        }

        /**
         * @param blockInterval If set, when a client reaches a rate limit threshold, the client will
         * be prohibited from any further requests until after the &#39;block_interval&#39; in seconds has elapsed.
         * 
         * @return builder
         * 
         */
        public Builder blockInterval(Integer blockInterval) {
            return blockInterval(Output.of(blockInterval));
        }

        /**
         * @param inheritable If set to `true` on a quota where path is set to a namespace, the same quota will be cumulatively applied to all child namespace. The inheritable parameter cannot be set to `true` if the path does not specify a namespace. Only the quotas associated with the root namespace are inheritable by default. Requires Vault 1.15+.
         * 
         * @return builder
         * 
         */
        public Builder inheritable(@Nullable Output<Boolean> inheritable) {
            $.inheritable = inheritable;
            return this;
        }

        /**
         * @param inheritable If set to `true` on a quota where path is set to a namespace, the same quota will be cumulatively applied to all child namespace. The inheritable parameter cannot be set to `true` if the path does not specify a namespace. Only the quotas associated with the root namespace are inheritable by default. Requires Vault 1.15+.
         * 
         * @return builder
         * 
         */
        public Builder inheritable(Boolean inheritable) {
            return inheritable(Output.of(inheritable));
        }

        /**
         * @param interval The duration in seconds to enforce rate limiting for.
         * 
         * @return builder
         * 
         */
        public Builder interval(@Nullable Output<Integer> interval) {
            $.interval = interval;
            return this;
        }

        /**
         * @param interval The duration in seconds to enforce rate limiting for.
         * 
         * @return builder
         * 
         */
        public Builder interval(Integer interval) {
            return interval(Output.of(interval));
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
         * @param rate The maximum number of requests at any given second to be allowed by the quota
         * rule. The `rate` must be positive.
         * 
         * @return builder
         * 
         */
        public Builder rate(Output<Double> rate) {
            $.rate = rate;
            return this;
        }

        /**
         * @param rate The maximum number of requests at any given second to be allowed by the quota
         * rule. The `rate` must be positive.
         * 
         * @return builder
         * 
         */
        public Builder rate(Double rate) {
            return rate(Output.of(rate));
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

        public QuotaRateLimitArgs build() {
            if ($.rate == null) {
                throw new MissingRequiredPropertyException("QuotaRateLimitArgs", "rate");
            }
            return $;
        }
    }

}
