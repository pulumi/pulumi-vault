// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vault.kv;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import com.pulumi.vault.kv.inputs.SecretV2CustomMetadataArgs;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.Object;
import java.lang.String;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class SecretV2Args extends com.pulumi.resources.ResourceArgs {

    public static final SecretV2Args Empty = new SecretV2Args();

    /**
     * This flag is required if `cas_required` is set to true
     * on either the secret or the engine&#39;s config. In order for a
     * write operation to be successful, cas must be set to the current version
     * of the secret.
     * 
     */
    @Import(name="cas")
    private @Nullable Output<Integer> cas;

    /**
     * @return This flag is required if `cas_required` is set to true
     * on either the secret or the engine&#39;s config. In order for a
     * write operation to be successful, cas must be set to the current version
     * of the secret.
     * 
     */
    public Optional<Output<Integer>> cas() {
        return Optional.ofNullable(this.cas);
    }

    /**
     * A nested block that allows configuring metadata for the
     * KV secret. Refer to the
     * Configuration Options for more info.
     * 
     */
    @Import(name="customMetadata")
    private @Nullable Output<SecretV2CustomMetadataArgs> customMetadata;

    /**
     * @return A nested block that allows configuring metadata for the
     * KV secret. Refer to the
     * Configuration Options for more info.
     * 
     */
    public Optional<Output<SecretV2CustomMetadataArgs>> customMetadata() {
        return Optional.ofNullable(this.customMetadata);
    }

    /**
     * JSON-encoded string that will be
     * written as the secret data at the given path.
     * 
     */
    @Import(name="dataJson", required=true)
    private Output<String> dataJson;

    /**
     * @return JSON-encoded string that will be
     * written as the secret data at the given path.
     * 
     */
    public Output<String> dataJson() {
        return this.dataJson;
    }

    /**
     * If set to true, permanently deletes all
     * versions for the specified key.
     * 
     */
    @Import(name="deleteAllVersions")
    private @Nullable Output<Boolean> deleteAllVersions;

    /**
     * @return If set to true, permanently deletes all
     * versions for the specified key.
     * 
     */
    public Optional<Output<Boolean>> deleteAllVersions() {
        return Optional.ofNullable(this.deleteAllVersions);
    }

    /**
     * If set to true, disables reading secret from Vault;
     * note: drift won&#39;t be detected.
     * 
     */
    @Import(name="disableRead")
    private @Nullable Output<Boolean> disableRead;

    /**
     * @return If set to true, disables reading secret from Vault;
     * note: drift won&#39;t be detected.
     * 
     */
    public Optional<Output<Boolean>> disableRead() {
        return Optional.ofNullable(this.disableRead);
    }

    /**
     * Path where KV-V2 engine is mounted.
     * 
     */
    @Import(name="mount", required=true)
    private Output<String> mount;

    /**
     * @return Path where KV-V2 engine is mounted.
     * 
     */
    public Output<String> mount() {
        return this.mount;
    }

    /**
     * Full name of the secret. For a nested secret
     * the name is the nested path excluding the mount and data
     * prefix. For example, for a secret at `kvv2/data/foo/bar/baz`
     * the name is `foo/bar/baz`.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return Full name of the secret. For a nested secret
     * the name is the nested path excluding the mount and data
     * prefix. For example, for a secret at `kvv2/data/foo/bar/baz`
     * the name is `foo/bar/baz`.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
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
     * An object that holds option settings.
     * 
     */
    @Import(name="options")
    private @Nullable Output<Map<String,Object>> options;

    /**
     * @return An object that holds option settings.
     * 
     */
    public Optional<Output<Map<String,Object>>> options() {
        return Optional.ofNullable(this.options);
    }

    private SecretV2Args() {}

    private SecretV2Args(SecretV2Args $) {
        this.cas = $.cas;
        this.customMetadata = $.customMetadata;
        this.dataJson = $.dataJson;
        this.deleteAllVersions = $.deleteAllVersions;
        this.disableRead = $.disableRead;
        this.mount = $.mount;
        this.name = $.name;
        this.namespace = $.namespace;
        this.options = $.options;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(SecretV2Args defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private SecretV2Args $;

        public Builder() {
            $ = new SecretV2Args();
        }

        public Builder(SecretV2Args defaults) {
            $ = new SecretV2Args(Objects.requireNonNull(defaults));
        }

        /**
         * @param cas This flag is required if `cas_required` is set to true
         * on either the secret or the engine&#39;s config. In order for a
         * write operation to be successful, cas must be set to the current version
         * of the secret.
         * 
         * @return builder
         * 
         */
        public Builder cas(@Nullable Output<Integer> cas) {
            $.cas = cas;
            return this;
        }

        /**
         * @param cas This flag is required if `cas_required` is set to true
         * on either the secret or the engine&#39;s config. In order for a
         * write operation to be successful, cas must be set to the current version
         * of the secret.
         * 
         * @return builder
         * 
         */
        public Builder cas(Integer cas) {
            return cas(Output.of(cas));
        }

        /**
         * @param customMetadata A nested block that allows configuring metadata for the
         * KV secret. Refer to the
         * Configuration Options for more info.
         * 
         * @return builder
         * 
         */
        public Builder customMetadata(@Nullable Output<SecretV2CustomMetadataArgs> customMetadata) {
            $.customMetadata = customMetadata;
            return this;
        }

        /**
         * @param customMetadata A nested block that allows configuring metadata for the
         * KV secret. Refer to the
         * Configuration Options for more info.
         * 
         * @return builder
         * 
         */
        public Builder customMetadata(SecretV2CustomMetadataArgs customMetadata) {
            return customMetadata(Output.of(customMetadata));
        }

        /**
         * @param dataJson JSON-encoded string that will be
         * written as the secret data at the given path.
         * 
         * @return builder
         * 
         */
        public Builder dataJson(Output<String> dataJson) {
            $.dataJson = dataJson;
            return this;
        }

        /**
         * @param dataJson JSON-encoded string that will be
         * written as the secret data at the given path.
         * 
         * @return builder
         * 
         */
        public Builder dataJson(String dataJson) {
            return dataJson(Output.of(dataJson));
        }

        /**
         * @param deleteAllVersions If set to true, permanently deletes all
         * versions for the specified key.
         * 
         * @return builder
         * 
         */
        public Builder deleteAllVersions(@Nullable Output<Boolean> deleteAllVersions) {
            $.deleteAllVersions = deleteAllVersions;
            return this;
        }

        /**
         * @param deleteAllVersions If set to true, permanently deletes all
         * versions for the specified key.
         * 
         * @return builder
         * 
         */
        public Builder deleteAllVersions(Boolean deleteAllVersions) {
            return deleteAllVersions(Output.of(deleteAllVersions));
        }

        /**
         * @param disableRead If set to true, disables reading secret from Vault;
         * note: drift won&#39;t be detected.
         * 
         * @return builder
         * 
         */
        public Builder disableRead(@Nullable Output<Boolean> disableRead) {
            $.disableRead = disableRead;
            return this;
        }

        /**
         * @param disableRead If set to true, disables reading secret from Vault;
         * note: drift won&#39;t be detected.
         * 
         * @return builder
         * 
         */
        public Builder disableRead(Boolean disableRead) {
            return disableRead(Output.of(disableRead));
        }

        /**
         * @param mount Path where KV-V2 engine is mounted.
         * 
         * @return builder
         * 
         */
        public Builder mount(Output<String> mount) {
            $.mount = mount;
            return this;
        }

        /**
         * @param mount Path where KV-V2 engine is mounted.
         * 
         * @return builder
         * 
         */
        public Builder mount(String mount) {
            return mount(Output.of(mount));
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
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
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
            return name(Output.of(name));
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
         * @param options An object that holds option settings.
         * 
         * @return builder
         * 
         */
        public Builder options(@Nullable Output<Map<String,Object>> options) {
            $.options = options;
            return this;
        }

        /**
         * @param options An object that holds option settings.
         * 
         * @return builder
         * 
         */
        public Builder options(Map<String,Object> options) {
            return options(Output.of(options));
        }

        public SecretV2Args build() {
            if ($.dataJson == null) {
                throw new MissingRequiredPropertyException("SecretV2Args", "dataJson");
            }
            if ($.mount == null) {
                throw new MissingRequiredPropertyException("SecretV2Args", "mount");
            }
            return $;
        }
    }

}
