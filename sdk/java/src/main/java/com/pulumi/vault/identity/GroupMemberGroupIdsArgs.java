// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vault.identity;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GroupMemberGroupIdsArgs extends com.pulumi.resources.ResourceArgs {

    public static final GroupMemberGroupIdsArgs Empty = new GroupMemberGroupIdsArgs();

    /**
     * Defaults to `true`.
     * 
     */
    @Import(name="exclusive")
    private @Nullable Output<Boolean> exclusive;

    /**
     * @return Defaults to `true`.
     * 
     */
    public Optional<Output<Boolean>> exclusive() {
        return Optional.ofNullable(this.exclusive);
    }

    /**
     * Group ID to assign member entities to.
     * 
     */
    @Import(name="groupId", required=true)
    private Output<String> groupId;

    /**
     * @return Group ID to assign member entities to.
     * 
     */
    public Output<String> groupId() {
        return this.groupId;
    }

    /**
     * List of member groups that belong to the group
     * 
     */
    @Import(name="memberGroupIds")
    private @Nullable Output<List<String>> memberGroupIds;

    /**
     * @return List of member groups that belong to the group
     * 
     */
    public Optional<Output<List<String>>> memberGroupIds() {
        return Optional.ofNullable(this.memberGroupIds);
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

    private GroupMemberGroupIdsArgs() {}

    private GroupMemberGroupIdsArgs(GroupMemberGroupIdsArgs $) {
        this.exclusive = $.exclusive;
        this.groupId = $.groupId;
        this.memberGroupIds = $.memberGroupIds;
        this.namespace = $.namespace;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GroupMemberGroupIdsArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GroupMemberGroupIdsArgs $;

        public Builder() {
            $ = new GroupMemberGroupIdsArgs();
        }

        public Builder(GroupMemberGroupIdsArgs defaults) {
            $ = new GroupMemberGroupIdsArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param exclusive Defaults to `true`.
         * 
         * @return builder
         * 
         */
        public Builder exclusive(@Nullable Output<Boolean> exclusive) {
            $.exclusive = exclusive;
            return this;
        }

        /**
         * @param exclusive Defaults to `true`.
         * 
         * @return builder
         * 
         */
        public Builder exclusive(Boolean exclusive) {
            return exclusive(Output.of(exclusive));
        }

        /**
         * @param groupId Group ID to assign member entities to.
         * 
         * @return builder
         * 
         */
        public Builder groupId(Output<String> groupId) {
            $.groupId = groupId;
            return this;
        }

        /**
         * @param groupId Group ID to assign member entities to.
         * 
         * @return builder
         * 
         */
        public Builder groupId(String groupId) {
            return groupId(Output.of(groupId));
        }

        /**
         * @param memberGroupIds List of member groups that belong to the group
         * 
         * @return builder
         * 
         */
        public Builder memberGroupIds(@Nullable Output<List<String>> memberGroupIds) {
            $.memberGroupIds = memberGroupIds;
            return this;
        }

        /**
         * @param memberGroupIds List of member groups that belong to the group
         * 
         * @return builder
         * 
         */
        public Builder memberGroupIds(List<String> memberGroupIds) {
            return memberGroupIds(Output.of(memberGroupIds));
        }

        /**
         * @param memberGroupIds List of member groups that belong to the group
         * 
         * @return builder
         * 
         */
        public Builder memberGroupIds(String... memberGroupIds) {
            return memberGroupIds(List.of(memberGroupIds));
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

        public GroupMemberGroupIdsArgs build() {
            $.groupId = Objects.requireNonNull($.groupId, "expected parameter 'groupId' to be non-null");
            return $;
        }
    }

}