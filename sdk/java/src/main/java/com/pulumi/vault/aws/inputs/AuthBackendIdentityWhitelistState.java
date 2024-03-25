// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vault.aws.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class AuthBackendIdentityWhitelistState extends com.pulumi.resources.ResourceArgs {

    public static final AuthBackendIdentityWhitelistState Empty = new AuthBackendIdentityWhitelistState();

    /**
     * The path of the AWS backend being configured.
     * 
     */
    @Import(name="backend")
    private @Nullable Output<String> backend;

    /**
     * @return The path of the AWS backend being configured.
     * 
     */
    public Optional<Output<String>> backend() {
        return Optional.ofNullable(this.backend);
    }

    /**
     * If set to true, disables the periodic
     * tidying of the identity-whitelist entries.
     * 
     */
    @Import(name="disablePeriodicTidy")
    private @Nullable Output<Boolean> disablePeriodicTidy;

    /**
     * @return If set to true, disables the periodic
     * tidying of the identity-whitelist entries.
     * 
     */
    public Optional<Output<Boolean>> disablePeriodicTidy() {
        return Optional.ofNullable(this.disablePeriodicTidy);
    }

    /**
     * The namespace to provision the resource in.
     * The value should not contain leading or trailing forward slashes.
     * The `namespace` is always relative to the provider&#39;s configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
     * *Available only for Vault Enterprise*.
     * 
     */
    @Import(name="namespace")
    private @Nullable Output<String> namespace;

    /**
     * @return The namespace to provision the resource in.
     * The value should not contain leading or trailing forward slashes.
     * The `namespace` is always relative to the provider&#39;s configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
     * *Available only for Vault Enterprise*.
     * 
     */
    public Optional<Output<String>> namespace() {
        return Optional.ofNullable(this.namespace);
    }

    /**
     * The amount of extra time, in minutes, that must
     * have passed beyond the roletag expiration, before it is removed from the
     * backend storage.
     * 
     */
    @Import(name="safetyBuffer")
    private @Nullable Output<Integer> safetyBuffer;

    /**
     * @return The amount of extra time, in minutes, that must
     * have passed beyond the roletag expiration, before it is removed from the
     * backend storage.
     * 
     */
    public Optional<Output<Integer>> safetyBuffer() {
        return Optional.ofNullable(this.safetyBuffer);
    }

    private AuthBackendIdentityWhitelistState() {}

    private AuthBackendIdentityWhitelistState(AuthBackendIdentityWhitelistState $) {
        this.backend = $.backend;
        this.disablePeriodicTidy = $.disablePeriodicTidy;
        this.namespace = $.namespace;
        this.safetyBuffer = $.safetyBuffer;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(AuthBackendIdentityWhitelistState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private AuthBackendIdentityWhitelistState $;

        public Builder() {
            $ = new AuthBackendIdentityWhitelistState();
        }

        public Builder(AuthBackendIdentityWhitelistState defaults) {
            $ = new AuthBackendIdentityWhitelistState(Objects.requireNonNull(defaults));
        }

        /**
         * @param backend The path of the AWS backend being configured.
         * 
         * @return builder
         * 
         */
        public Builder backend(@Nullable Output<String> backend) {
            $.backend = backend;
            return this;
        }

        /**
         * @param backend The path of the AWS backend being configured.
         * 
         * @return builder
         * 
         */
        public Builder backend(String backend) {
            return backend(Output.of(backend));
        }

        /**
         * @param disablePeriodicTidy If set to true, disables the periodic
         * tidying of the identity-whitelist entries.
         * 
         * @return builder
         * 
         */
        public Builder disablePeriodicTidy(@Nullable Output<Boolean> disablePeriodicTidy) {
            $.disablePeriodicTidy = disablePeriodicTidy;
            return this;
        }

        /**
         * @param disablePeriodicTidy If set to true, disables the periodic
         * tidying of the identity-whitelist entries.
         * 
         * @return builder
         * 
         */
        public Builder disablePeriodicTidy(Boolean disablePeriodicTidy) {
            return disablePeriodicTidy(Output.of(disablePeriodicTidy));
        }

        /**
         * @param namespace The namespace to provision the resource in.
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
         * @param namespace The namespace to provision the resource in.
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
         * @param safetyBuffer The amount of extra time, in minutes, that must
         * have passed beyond the roletag expiration, before it is removed from the
         * backend storage.
         * 
         * @return builder
         * 
         */
        public Builder safetyBuffer(@Nullable Output<Integer> safetyBuffer) {
            $.safetyBuffer = safetyBuffer;
            return this;
        }

        /**
         * @param safetyBuffer The amount of extra time, in minutes, that must
         * have passed beyond the roletag expiration, before it is removed from the
         * backend storage.
         * 
         * @return builder
         * 
         */
        public Builder safetyBuffer(Integer safetyBuffer) {
            return safetyBuffer(Output.of(safetyBuffer));
        }

        public AuthBackendIdentityWhitelistState build() {
            return $;
        }
    }

}
