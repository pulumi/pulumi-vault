// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vault.kv.inputs;

import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetSecretSubkeysV2PlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetSecretSubkeysV2PlainArgs Empty = new GetSecretSubkeysV2PlainArgs();

    /**
     * Specifies the deepest nesting level to provide in the output.
     * If non-zero, keys that reside at the specified depth value will be
     * artificially treated as leaves and will thus be `null` even if further
     * underlying sub-keys exist.
     * 
     */
    @Import(name="depth")
    private @Nullable Integer depth;

    /**
     * @return Specifies the deepest nesting level to provide in the output.
     * If non-zero, keys that reside at the specified depth value will be
     * artificially treated as leaves and will thus be `null` even if further
     * underlying sub-keys exist.
     * 
     */
    public Optional<Integer> depth() {
        return Optional.ofNullable(this.depth);
    }

    /**
     * Path where KV-V2 engine is mounted.
     * 
     */
    @Import(name="mount", required=true)
    private String mount;

    /**
     * @return Path where KV-V2 engine is mounted.
     * 
     */
    public String mount() {
        return this.mount;
    }

    /**
     * Full name of the secret. For a nested secret
     * the name is the nested path excluding the mount and data
     * prefix. For example, for a secret at `kvv2/data/foo/bar/baz`
     * the name is `foo/bar/baz`.
     * 
     */
    @Import(name="name", required=true)
    private String name;

    /**
     * @return Full name of the secret. For a nested secret
     * the name is the nested path excluding the mount and data
     * prefix. For example, for a secret at `kvv2/data/foo/bar/baz`
     * the name is `foo/bar/baz`.
     * 
     */
    public String name() {
        return this.name;
    }

    /**
     * The namespace of the target resource.
     * The value should not contain leading or trailing forward slashes.
     * The `namespace` is always relative to the provider&#39;s configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
     * *Available only for Vault Enterprise*.
     * 
     */
    @Import(name="namespace")
    private @Nullable String namespace;

    /**
     * @return The namespace of the target resource.
     * The value should not contain leading or trailing forward slashes.
     * The `namespace` is always relative to the provider&#39;s configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
     * *Available only for Vault Enterprise*.
     * 
     */
    public Optional<String> namespace() {
        return Optional.ofNullable(this.namespace);
    }

    /**
     * Specifies the version to return. If not
     * set the latest version is returned.
     * 
     */
    @Import(name="version")
    private @Nullable Integer version;

    /**
     * @return Specifies the version to return. If not
     * set the latest version is returned.
     * 
     */
    public Optional<Integer> version() {
        return Optional.ofNullable(this.version);
    }

    private GetSecretSubkeysV2PlainArgs() {}

    private GetSecretSubkeysV2PlainArgs(GetSecretSubkeysV2PlainArgs $) {
        this.depth = $.depth;
        this.mount = $.mount;
        this.name = $.name;
        this.namespace = $.namespace;
        this.version = $.version;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetSecretSubkeysV2PlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetSecretSubkeysV2PlainArgs $;

        public Builder() {
            $ = new GetSecretSubkeysV2PlainArgs();
        }

        public Builder(GetSecretSubkeysV2PlainArgs defaults) {
            $ = new GetSecretSubkeysV2PlainArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param depth Specifies the deepest nesting level to provide in the output.
         * If non-zero, keys that reside at the specified depth value will be
         * artificially treated as leaves and will thus be `null` even if further
         * underlying sub-keys exist.
         * 
         * @return builder
         * 
         */
        public Builder depth(@Nullable Integer depth) {
            $.depth = depth;
            return this;
        }

        /**
         * @param mount Path where KV-V2 engine is mounted.
         * 
         * @return builder
         * 
         */
        public Builder mount(String mount) {
            $.mount = mount;
            return this;
        }

        /**
         * @param name Full name of the secret. For a nested secret
         * the name is the nested path excluding the mount and data
         * prefix. For example, for a secret at `kvv2/data/foo/bar/baz`
         * the name is `foo/bar/baz`.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            $.name = name;
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
        public Builder namespace(@Nullable String namespace) {
            $.namespace = namespace;
            return this;
        }

        /**
         * @param version Specifies the version to return. If not
         * set the latest version is returned.
         * 
         * @return builder
         * 
         */
        public Builder version(@Nullable Integer version) {
            $.version = version;
            return this;
        }

        public GetSecretSubkeysV2PlainArgs build() {
            if ($.mount == null) {
                throw new MissingRequiredPropertyException("GetSecretSubkeysV2PlainArgs", "mount");
            }
            if ($.name == null) {
                throw new MissingRequiredPropertyException("GetSecretSubkeysV2PlainArgs", "name");
            }
            return $;
        }
    }

}
