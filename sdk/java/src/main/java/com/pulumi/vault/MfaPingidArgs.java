// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vault;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class MfaPingidArgs extends com.pulumi.resources.ResourceArgs {

    public static final MfaPingidArgs Empty = new MfaPingidArgs();

    /**
     * `(string: &lt;required&gt;)` - The mount to tie this method to for use in automatic mappings.
     * The mapping will use the Name field of Aliases associated with this mount as the username in the mapping.
     * 
     */
    @Import(name="mountAccessor", required=true)
    private Output<String> mountAccessor;

    /**
     * @return `(string: &lt;required&gt;)` - The mount to tie this method to for use in automatic mappings.
     * The mapping will use the Name field of Aliases associated with this mount as the username in the mapping.
     * 
     */
    public Output<String> mountAccessor() {
        return this.mountAccessor;
    }

    /**
     * `(string: &lt;required&gt;)` – Name of the MFA method.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return `(string: &lt;required&gt;)` – Name of the MFA method.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
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
     * `(string: &lt;required&gt;)` - A base64-encoded third-party settings file retrieved
     * from PingID&#39;s configuration page.
     * 
     */
    @Import(name="settingsFileBase64", required=true)
    private Output<String> settingsFileBase64;

    /**
     * @return `(string: &lt;required&gt;)` - A base64-encoded third-party settings file retrieved
     * from PingID&#39;s configuration page.
     * 
     */
    public Output<String> settingsFileBase64() {
        return this.settingsFileBase64;
    }

    /**
     * `(string)` - A format string for mapping Identity names to MFA method names.
     * Values to substitute should be placed in `{{}}`. For example, `&#34;{{alias.name}}@example.com&#34;`.
     * If blank, the Alias&#39;s Name field will be used as-is. Currently-supported mappings:
     * - alias.name: The name returned by the mount configured via the `mount_accessor` parameter
     * - entity.name: The name configured for the Entity
     * - alias.metadata.`&lt;key&gt;`: The value of the Alias&#39;s metadata parameter
     * - entity.metadata.`&lt;key&gt;`: The value of the Entity&#39;s metadata parameter
     * 
     */
    @Import(name="usernameFormat")
    private @Nullable Output<String> usernameFormat;

    /**
     * @return `(string)` - A format string for mapping Identity names to MFA method names.
     * Values to substitute should be placed in `{{}}`. For example, `&#34;{{alias.name}}@example.com&#34;`.
     * If blank, the Alias&#39;s Name field will be used as-is. Currently-supported mappings:
     * - alias.name: The name returned by the mount configured via the `mount_accessor` parameter
     * - entity.name: The name configured for the Entity
     * - alias.metadata.`&lt;key&gt;`: The value of the Alias&#39;s metadata parameter
     * - entity.metadata.`&lt;key&gt;`: The value of the Entity&#39;s metadata parameter
     * 
     */
    public Optional<Output<String>> usernameFormat() {
        return Optional.ofNullable(this.usernameFormat);
    }

    private MfaPingidArgs() {}

    private MfaPingidArgs(MfaPingidArgs $) {
        this.mountAccessor = $.mountAccessor;
        this.name = $.name;
        this.namespace = $.namespace;
        this.settingsFileBase64 = $.settingsFileBase64;
        this.usernameFormat = $.usernameFormat;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(MfaPingidArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private MfaPingidArgs $;

        public Builder() {
            $ = new MfaPingidArgs();
        }

        public Builder(MfaPingidArgs defaults) {
            $ = new MfaPingidArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param mountAccessor `(string: &lt;required&gt;)` - The mount to tie this method to for use in automatic mappings.
         * The mapping will use the Name field of Aliases associated with this mount as the username in the mapping.
         * 
         * @return builder
         * 
         */
        public Builder mountAccessor(Output<String> mountAccessor) {
            $.mountAccessor = mountAccessor;
            return this;
        }

        /**
         * @param mountAccessor `(string: &lt;required&gt;)` - The mount to tie this method to for use in automatic mappings.
         * The mapping will use the Name field of Aliases associated with this mount as the username in the mapping.
         * 
         * @return builder
         * 
         */
        public Builder mountAccessor(String mountAccessor) {
            return mountAccessor(Output.of(mountAccessor));
        }

        /**
         * @param name `(string: &lt;required&gt;)` – Name of the MFA method.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name `(string: &lt;required&gt;)` – Name of the MFA method.
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
         * @param settingsFileBase64 `(string: &lt;required&gt;)` - A base64-encoded third-party settings file retrieved
         * from PingID&#39;s configuration page.
         * 
         * @return builder
         * 
         */
        public Builder settingsFileBase64(Output<String> settingsFileBase64) {
            $.settingsFileBase64 = settingsFileBase64;
            return this;
        }

        /**
         * @param settingsFileBase64 `(string: &lt;required&gt;)` - A base64-encoded third-party settings file retrieved
         * from PingID&#39;s configuration page.
         * 
         * @return builder
         * 
         */
        public Builder settingsFileBase64(String settingsFileBase64) {
            return settingsFileBase64(Output.of(settingsFileBase64));
        }

        /**
         * @param usernameFormat `(string)` - A format string for mapping Identity names to MFA method names.
         * Values to substitute should be placed in `{{}}`. For example, `&#34;{{alias.name}}@example.com&#34;`.
         * If blank, the Alias&#39;s Name field will be used as-is. Currently-supported mappings:
         * - alias.name: The name returned by the mount configured via the `mount_accessor` parameter
         * - entity.name: The name configured for the Entity
         * - alias.metadata.`&lt;key&gt;`: The value of the Alias&#39;s metadata parameter
         * - entity.metadata.`&lt;key&gt;`: The value of the Entity&#39;s metadata parameter
         * 
         * @return builder
         * 
         */
        public Builder usernameFormat(@Nullable Output<String> usernameFormat) {
            $.usernameFormat = usernameFormat;
            return this;
        }

        /**
         * @param usernameFormat `(string)` - A format string for mapping Identity names to MFA method names.
         * Values to substitute should be placed in `{{}}`. For example, `&#34;{{alias.name}}@example.com&#34;`.
         * If blank, the Alias&#39;s Name field will be used as-is. Currently-supported mappings:
         * - alias.name: The name returned by the mount configured via the `mount_accessor` parameter
         * - entity.name: The name configured for the Entity
         * - alias.metadata.`&lt;key&gt;`: The value of the Alias&#39;s metadata parameter
         * - entity.metadata.`&lt;key&gt;`: The value of the Entity&#39;s metadata parameter
         * 
         * @return builder
         * 
         */
        public Builder usernameFormat(String usernameFormat) {
            return usernameFormat(Output.of(usernameFormat));
        }

        public MfaPingidArgs build() {
            if ($.mountAccessor == null) {
                throw new MissingRequiredPropertyException("MfaPingidArgs", "mountAccessor");
            }
            if ($.settingsFileBase64 == null) {
                throw new MissingRequiredPropertyException("MfaPingidArgs", "settingsFileBase64");
            }
            return $;
        }
    }

}
