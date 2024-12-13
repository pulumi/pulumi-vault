// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vault.identity;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GroupMemberEntityIdsArgs extends com.pulumi.resources.ResourceArgs {

    public static final GroupMemberEntityIdsArgs Empty = new GroupMemberEntityIdsArgs();

    /**
     * Defaults to `true`.
     * 
     * If `true`, this resource will take exclusive control of the member entities that belong to the group and will set it equal to what is specified in the resource.
     * 
     * If set to `false`, this resource will simply ensure that the member entities specified in the resource are present in the group. When destroying the resource, the resource will ensure that the member entities specified in the resource are removed.
     * 
     */
    @Import(name="exclusive")
    private @Nullable Output<Boolean> exclusive;

    /**
     * @return Defaults to `true`.
     * 
     * If `true`, this resource will take exclusive control of the member entities that belong to the group and will set it equal to what is specified in the resource.
     * 
     * If set to `false`, this resource will simply ensure that the member entities specified in the resource are present in the group. When destroying the resource, the resource will ensure that the member entities specified in the resource are removed.
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
     * List of member entities that belong to the group
     * 
     */
    @Import(name="memberEntityIds")
    private @Nullable Output<List<String>> memberEntityIds;

    /**
     * @return List of member entities that belong to the group
     * 
     */
    public Optional<Output<List<String>>> memberEntityIds() {
        return Optional.ofNullable(this.memberEntityIds);
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

    private GroupMemberEntityIdsArgs() {}

    private GroupMemberEntityIdsArgs(GroupMemberEntityIdsArgs $) {
        this.exclusive = $.exclusive;
        this.groupId = $.groupId;
        this.memberEntityIds = $.memberEntityIds;
        this.namespace = $.namespace;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GroupMemberEntityIdsArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GroupMemberEntityIdsArgs $;

        public Builder() {
            $ = new GroupMemberEntityIdsArgs();
        }

        public Builder(GroupMemberEntityIdsArgs defaults) {
            $ = new GroupMemberEntityIdsArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param exclusive Defaults to `true`.
         * 
         * If `true`, this resource will take exclusive control of the member entities that belong to the group and will set it equal to what is specified in the resource.
         * 
         * If set to `false`, this resource will simply ensure that the member entities specified in the resource are present in the group. When destroying the resource, the resource will ensure that the member entities specified in the resource are removed.
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
         * If `true`, this resource will take exclusive control of the member entities that belong to the group and will set it equal to what is specified in the resource.
         * 
         * If set to `false`, this resource will simply ensure that the member entities specified in the resource are present in the group. When destroying the resource, the resource will ensure that the member entities specified in the resource are removed.
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
         * @param memberEntityIds List of member entities that belong to the group
         * 
         * @return builder
         * 
         */
        public Builder memberEntityIds(@Nullable Output<List<String>> memberEntityIds) {
            $.memberEntityIds = memberEntityIds;
            return this;
        }

        /**
         * @param memberEntityIds List of member entities that belong to the group
         * 
         * @return builder
         * 
         */
        public Builder memberEntityIds(List<String> memberEntityIds) {
            return memberEntityIds(Output.of(memberEntityIds));
        }

        /**
         * @param memberEntityIds List of member entities that belong to the group
         * 
         * @return builder
         * 
         */
        public Builder memberEntityIds(String... memberEntityIds) {
            return memberEntityIds(List.of(memberEntityIds));
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

        public GroupMemberEntityIdsArgs build() {
            if ($.groupId == null) {
                throw new MissingRequiredPropertyException("GroupMemberEntityIdsArgs", "groupId");
            }
            return $;
        }
    }

}