// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vault;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import com.pulumi.vault.inputs.AuthBackendTuneArgs;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class AuthBackendArgs extends com.pulumi.resources.ResourceArgs {

    public static final AuthBackendArgs Empty = new AuthBackendArgs();

    /**
     * A description of the auth method.
     * 
     */
    @Import(name="description")
    private @Nullable Output<String> description;

    /**
     * @return A description of the auth method.
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
     * Specifies if the auth method is local only.
     * 
     */
    @Import(name="local")
    private @Nullable Output<Boolean> local;

    /**
     * @return Specifies if the auth method is local only.
     * 
     */
    public Optional<Output<Boolean>> local() {
        return Optional.ofNullable(this.local);
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
     * The path to mount the auth method — this defaults to the name of the type.
     * 
     */
    @Import(name="path")
    private @Nullable Output<String> path;

    /**
     * @return The path to mount the auth method — this defaults to the name of the type.
     * 
     */
    public Optional<Output<String>> path() {
        return Optional.ofNullable(this.path);
    }

    /**
     * Extra configuration block. Structure is documented below.
     * 
     * The `tune` block is used to tune the auth backend:
     * 
     */
    @Import(name="tune")
    private @Nullable Output<AuthBackendTuneArgs> tune;

    /**
     * @return Extra configuration block. Structure is documented below.
     * 
     * The `tune` block is used to tune the auth backend:
     * 
     */
    public Optional<Output<AuthBackendTuneArgs>> tune() {
        return Optional.ofNullable(this.tune);
    }

    /**
     * The name of the auth method type.
     * 
     */
    @Import(name="type", required=true)
    private Output<String> type;

    /**
     * @return The name of the auth method type.
     * 
     */
    public Output<String> type() {
        return this.type;
    }

    private AuthBackendArgs() {}

    private AuthBackendArgs(AuthBackendArgs $) {
        this.description = $.description;
        this.disableRemount = $.disableRemount;
        this.local = $.local;
        this.namespace = $.namespace;
        this.path = $.path;
        this.tune = $.tune;
        this.type = $.type;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(AuthBackendArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private AuthBackendArgs $;

        public Builder() {
            $ = new AuthBackendArgs();
        }

        public Builder(AuthBackendArgs defaults) {
            $ = new AuthBackendArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param description A description of the auth method.
         * 
         * @return builder
         * 
         */
        public Builder description(@Nullable Output<String> description) {
            $.description = description;
            return this;
        }

        /**
         * @param description A description of the auth method.
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
         * @param local Specifies if the auth method is local only.
         * 
         * @return builder
         * 
         */
        public Builder local(@Nullable Output<Boolean> local) {
            $.local = local;
            return this;
        }

        /**
         * @param local Specifies if the auth method is local only.
         * 
         * @return builder
         * 
         */
        public Builder local(Boolean local) {
            return local(Output.of(local));
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
         * @param path The path to mount the auth method — this defaults to the name of the type.
         * 
         * @return builder
         * 
         */
        public Builder path(@Nullable Output<String> path) {
            $.path = path;
            return this;
        }

        /**
         * @param path The path to mount the auth method — this defaults to the name of the type.
         * 
         * @return builder
         * 
         */
        public Builder path(String path) {
            return path(Output.of(path));
        }

        /**
         * @param tune Extra configuration block. Structure is documented below.
         * 
         * The `tune` block is used to tune the auth backend:
         * 
         * @return builder
         * 
         */
        public Builder tune(@Nullable Output<AuthBackendTuneArgs> tune) {
            $.tune = tune;
            return this;
        }

        /**
         * @param tune Extra configuration block. Structure is documented below.
         * 
         * The `tune` block is used to tune the auth backend:
         * 
         * @return builder
         * 
         */
        public Builder tune(AuthBackendTuneArgs tune) {
            return tune(Output.of(tune));
        }

        /**
         * @param type The name of the auth method type.
         * 
         * @return builder
         * 
         */
        public Builder type(Output<String> type) {
            $.type = type;
            return this;
        }

        /**
         * @param type The name of the auth method type.
         * 
         * @return builder
         * 
         */
        public Builder type(String type) {
            return type(Output.of(type));
        }

        public AuthBackendArgs build() {
            if ($.type == null) {
                throw new MissingRequiredPropertyException("AuthBackendArgs", "type");
            }
            return $;
        }
    }

}
