// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vault.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Double;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class QuotaRateLimitState extends com.pulumi.resources.ResourceArgs {

    public static final QuotaRateLimitState Empty = new QuotaRateLimitState();

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
    @Import(name="rate")
    private @Nullable Output<Double> rate;

    /**
     * @return The maximum number of requests at any given second to be allowed by the quota
     * rule. The `rate` must be positive.
     * 
     */
    public Optional<Output<Double>> rate() {
        return Optional.ofNullable(this.rate);
    }

    private QuotaRateLimitState() {}

    private QuotaRateLimitState(QuotaRateLimitState $) {
        this.blockInterval = $.blockInterval;
        this.interval = $.interval;
        this.name = $.name;
        this.namespace = $.namespace;
        this.path = $.path;
        this.rate = $.rate;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(QuotaRateLimitState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private QuotaRateLimitState $;

        public Builder() {
            $ = new QuotaRateLimitState();
        }

        public Builder(QuotaRateLimitState defaults) {
            $ = new QuotaRateLimitState(Objects.requireNonNull(defaults));
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
        public Builder rate(@Nullable Output<Double> rate) {
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

        public QuotaRateLimitState build() {
            return $;
        }
    }

}