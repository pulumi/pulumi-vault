// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vault;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Boolean;
import java.lang.String;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class AuditArgs extends com.pulumi.resources.ResourceArgs {

    public static final AuditArgs Empty = new AuditArgs();

    /**
     * Human-friendly description of the audit device.
     * 
     */
    @Import(name="description")
    private @Nullable Output<String> description;

    /**
     * @return Human-friendly description of the audit device.
     * 
     */
    public Optional<Output<String>> description() {
        return Optional.ofNullable(this.description);
    }

    /**
     * Specifies if the audit device is a local only. Local audit devices are not replicated nor (if a secondary) removed by replication.
     * 
     */
    @Import(name="local")
    private @Nullable Output<Boolean> local;

    /**
     * @return Specifies if the audit device is a local only. Local audit devices are not replicated nor (if a secondary) removed by replication.
     * 
     */
    public Optional<Output<Boolean>> local() {
        return Optional.ofNullable(this.local);
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
     * Configuration options to pass to the audit device itself.
     * 
     * For a reference of the device types and their options, consult the [Vault documentation.](https://www.vaultproject.io/docs/audit/index.html)
     * 
     */
    @Import(name="options", required=true)
    private Output<Map<String,String>> options;

    /**
     * @return Configuration options to pass to the audit device itself.
     * 
     * For a reference of the device types and their options, consult the [Vault documentation.](https://www.vaultproject.io/docs/audit/index.html)
     * 
     */
    public Output<Map<String,String>> options() {
        return this.options;
    }

    /**
     * The path to mount the audit device. This defaults to the type.
     * 
     */
    @Import(name="path")
    private @Nullable Output<String> path;

    /**
     * @return The path to mount the audit device. This defaults to the type.
     * 
     */
    public Optional<Output<String>> path() {
        return Optional.ofNullable(this.path);
    }

    /**
     * Type of the audit device, such as &#39;file&#39;.
     * 
     */
    @Import(name="type", required=true)
    private Output<String> type;

    /**
     * @return Type of the audit device, such as &#39;file&#39;.
     * 
     */
    public Output<String> type() {
        return this.type;
    }

    private AuditArgs() {}

    private AuditArgs(AuditArgs $) {
        this.description = $.description;
        this.local = $.local;
        this.namespace = $.namespace;
        this.options = $.options;
        this.path = $.path;
        this.type = $.type;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(AuditArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private AuditArgs $;

        public Builder() {
            $ = new AuditArgs();
        }

        public Builder(AuditArgs defaults) {
            $ = new AuditArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param description Human-friendly description of the audit device.
         * 
         * @return builder
         * 
         */
        public Builder description(@Nullable Output<String> description) {
            $.description = description;
            return this;
        }

        /**
         * @param description Human-friendly description of the audit device.
         * 
         * @return builder
         * 
         */
        public Builder description(String description) {
            return description(Output.of(description));
        }

        /**
         * @param local Specifies if the audit device is a local only. Local audit devices are not replicated nor (if a secondary) removed by replication.
         * 
         * @return builder
         * 
         */
        public Builder local(@Nullable Output<Boolean> local) {
            $.local = local;
            return this;
        }

        /**
         * @param local Specifies if the audit device is a local only. Local audit devices are not replicated nor (if a secondary) removed by replication.
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
         * @param options Configuration options to pass to the audit device itself.
         * 
         * For a reference of the device types and their options, consult the [Vault documentation.](https://www.vaultproject.io/docs/audit/index.html)
         * 
         * @return builder
         * 
         */
        public Builder options(Output<Map<String,String>> options) {
            $.options = options;
            return this;
        }

        /**
         * @param options Configuration options to pass to the audit device itself.
         * 
         * For a reference of the device types and their options, consult the [Vault documentation.](https://www.vaultproject.io/docs/audit/index.html)
         * 
         * @return builder
         * 
         */
        public Builder options(Map<String,String> options) {
            return options(Output.of(options));
        }

        /**
         * @param path The path to mount the audit device. This defaults to the type.
         * 
         * @return builder
         * 
         */
        public Builder path(@Nullable Output<String> path) {
            $.path = path;
            return this;
        }

        /**
         * @param path The path to mount the audit device. This defaults to the type.
         * 
         * @return builder
         * 
         */
        public Builder path(String path) {
            return path(Output.of(path));
        }

        /**
         * @param type Type of the audit device, such as &#39;file&#39;.
         * 
         * @return builder
         * 
         */
        public Builder type(Output<String> type) {
            $.type = type;
            return this;
        }

        /**
         * @param type Type of the audit device, such as &#39;file&#39;.
         * 
         * @return builder
         * 
         */
        public Builder type(String type) {
            return type(Output.of(type));
        }

        public AuditArgs build() {
            if ($.options == null) {
                throw new MissingRequiredPropertyException("AuditArgs", "options");
            }
            if ($.type == null) {
                throw new MissingRequiredPropertyException("AuditArgs", "type");
            }
            return $;
        }
    }

}
