// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vault.identity.inputs;

import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetGroupPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetGroupPlainArgs Empty = new GetGroupPlainArgs();

    /**
     * ID of the alias.
     * 
     */
    @Import(name="aliasId")
    private @Nullable String aliasId;

    /**
     * @return ID of the alias.
     * 
     */
    public Optional<String> aliasId() {
        return Optional.ofNullable(this.aliasId);
    }

    /**
     * Accessor of the mount to which the alias belongs to.
     * This should be supplied in conjunction with `alias_name`.
     * 
     * The lookup criteria can be `group_name`, `group_id`, `alias_id`, or a combination of
     * `alias_name` and `alias_mount_accessor`.
     * 
     */
    @Import(name="aliasMountAccessor")
    private @Nullable String aliasMountAccessor;

    /**
     * @return Accessor of the mount to which the alias belongs to.
     * This should be supplied in conjunction with `alias_name`.
     * 
     * The lookup criteria can be `group_name`, `group_id`, `alias_id`, or a combination of
     * `alias_name` and `alias_mount_accessor`.
     * 
     */
    public Optional<String> aliasMountAccessor() {
        return Optional.ofNullable(this.aliasMountAccessor);
    }

    /**
     * Name of the alias. This should be supplied in conjunction with
     * `alias_mount_accessor`.
     * 
     */
    @Import(name="aliasName")
    private @Nullable String aliasName;

    /**
     * @return Name of the alias. This should be supplied in conjunction with
     * `alias_mount_accessor`.
     * 
     */
    public Optional<String> aliasName() {
        return Optional.ofNullable(this.aliasName);
    }

    /**
     * ID of the group.
     * 
     */
    @Import(name="groupId")
    private @Nullable String groupId;

    /**
     * @return ID of the group.
     * 
     */
    public Optional<String> groupId() {
        return Optional.ofNullable(this.groupId);
    }

    /**
     * Name of the group.
     * 
     */
    @Import(name="groupName")
    private @Nullable String groupName;

    /**
     * @return Name of the group.
     * 
     */
    public Optional<String> groupName() {
        return Optional.ofNullable(this.groupName);
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

    private GetGroupPlainArgs() {}

    private GetGroupPlainArgs(GetGroupPlainArgs $) {
        this.aliasId = $.aliasId;
        this.aliasMountAccessor = $.aliasMountAccessor;
        this.aliasName = $.aliasName;
        this.groupId = $.groupId;
        this.groupName = $.groupName;
        this.namespace = $.namespace;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetGroupPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetGroupPlainArgs $;

        public Builder() {
            $ = new GetGroupPlainArgs();
        }

        public Builder(GetGroupPlainArgs defaults) {
            $ = new GetGroupPlainArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param aliasId ID of the alias.
         * 
         * @return builder
         * 
         */
        public Builder aliasId(@Nullable String aliasId) {
            $.aliasId = aliasId;
            return this;
        }

        /**
         * @param aliasMountAccessor Accessor of the mount to which the alias belongs to.
         * This should be supplied in conjunction with `alias_name`.
         * 
         * The lookup criteria can be `group_name`, `group_id`, `alias_id`, or a combination of
         * `alias_name` and `alias_mount_accessor`.
         * 
         * @return builder
         * 
         */
        public Builder aliasMountAccessor(@Nullable String aliasMountAccessor) {
            $.aliasMountAccessor = aliasMountAccessor;
            return this;
        }

        /**
         * @param aliasName Name of the alias. This should be supplied in conjunction with
         * `alias_mount_accessor`.
         * 
         * @return builder
         * 
         */
        public Builder aliasName(@Nullable String aliasName) {
            $.aliasName = aliasName;
            return this;
        }

        /**
         * @param groupId ID of the group.
         * 
         * @return builder
         * 
         */
        public Builder groupId(@Nullable String groupId) {
            $.groupId = groupId;
            return this;
        }

        /**
         * @param groupName Name of the group.
         * 
         * @return builder
         * 
         */
        public Builder groupName(@Nullable String groupName) {
            $.groupName = groupName;
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

        public GetGroupPlainArgs build() {
            return $;
        }
    }

}
