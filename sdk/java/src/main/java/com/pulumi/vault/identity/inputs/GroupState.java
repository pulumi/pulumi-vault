// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vault.identity.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GroupState extends com.pulumi.resources.ResourceArgs {

    public static final GroupState Empty = new GroupState();

    /**
     * `false` by default. If set to `true`, this resource will ignore any Entity IDs
     * returned from Vault or specified in the resource. You can use
     * `vault.identity.GroupMemberEntityIds` to manage Entity IDs for this group in a
     * decoupled manner.
     * 
     */
    @Import(name="externalMemberEntityIds")
    private @Nullable Output<Boolean> externalMemberEntityIds;

    /**
     * @return `false` by default. If set to `true`, this resource will ignore any Entity IDs
     * returned from Vault or specified in the resource. You can use
     * `vault.identity.GroupMemberEntityIds` to manage Entity IDs for this group in a
     * decoupled manner.
     * 
     */
    public Optional<Output<Boolean>> externalMemberEntityIds() {
        return Optional.ofNullable(this.externalMemberEntityIds);
    }

    /**
     * `false` by default. If set to `true`, this resource will ignore any Group IDs
     * returned from Vault or specified in the resource. You can use
     * `vault.identity.GroupMemberGroupIds` to manage Group IDs for this group in a
     * decoupled manner.
     * 
     */
    @Import(name="externalMemberGroupIds")
    private @Nullable Output<Boolean> externalMemberGroupIds;

    /**
     * @return `false` by default. If set to `true`, this resource will ignore any Group IDs
     * returned from Vault or specified in the resource. You can use
     * `vault.identity.GroupMemberGroupIds` to manage Group IDs for this group in a
     * decoupled manner.
     * 
     */
    public Optional<Output<Boolean>> externalMemberGroupIds() {
        return Optional.ofNullable(this.externalMemberGroupIds);
    }

    /**
     * `false` by default. If set to `true`, this resource will ignore any policies returned from
     * Vault or specified in the resource. You can use `vault.identity.GroupPolicies` to manage
     * policies for this group in a decoupled manner.
     * 
     */
    @Import(name="externalPolicies")
    private @Nullable Output<Boolean> externalPolicies;

    /**
     * @return `false` by default. If set to `true`, this resource will ignore any policies returned from
     * Vault or specified in the resource. You can use `vault.identity.GroupPolicies` to manage
     * policies for this group in a decoupled manner.
     * 
     */
    public Optional<Output<Boolean>> externalPolicies() {
        return Optional.ofNullable(this.externalPolicies);
    }

    /**
     * A list of Entity IDs to be assigned as group members. Not allowed on `external` groups.
     * 
     */
    @Import(name="memberEntityIds")
    private @Nullable Output<List<String>> memberEntityIds;

    /**
     * @return A list of Entity IDs to be assigned as group members. Not allowed on `external` groups.
     * 
     */
    public Optional<Output<List<String>>> memberEntityIds() {
        return Optional.ofNullable(this.memberEntityIds);
    }

    /**
     * A list of Group IDs to be assigned as group members. Not allowed on `external` groups.
     * 
     */
    @Import(name="memberGroupIds")
    private @Nullable Output<List<String>> memberGroupIds;

    /**
     * @return A list of Group IDs to be assigned as group members. Not allowed on `external` groups.
     * 
     */
    public Optional<Output<List<String>>> memberGroupIds() {
        return Optional.ofNullable(this.memberGroupIds);
    }

    /**
     * A Map of additional metadata to associate with the group.
     * 
     */
    @Import(name="metadata")
    private @Nullable Output<Map<String,String>> metadata;

    /**
     * @return A Map of additional metadata to associate with the group.
     * 
     */
    public Optional<Output<Map<String,String>>> metadata() {
        return Optional.ofNullable(this.metadata);
    }

    /**
     * Name of the identity group to create.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return Name of the identity group to create.
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
     * A list of policies to apply to the group.
     * 
     */
    @Import(name="policies")
    private @Nullable Output<List<String>> policies;

    /**
     * @return A list of policies to apply to the group.
     * 
     */
    public Optional<Output<List<String>>> policies() {
        return Optional.ofNullable(this.policies);
    }

    /**
     * Type of the group, internal or external. Defaults to `internal`.
     * 
     */
    @Import(name="type")
    private @Nullable Output<String> type;

    /**
     * @return Type of the group, internal or external. Defaults to `internal`.
     * 
     */
    public Optional<Output<String>> type() {
        return Optional.ofNullable(this.type);
    }

    private GroupState() {}

    private GroupState(GroupState $) {
        this.externalMemberEntityIds = $.externalMemberEntityIds;
        this.externalMemberGroupIds = $.externalMemberGroupIds;
        this.externalPolicies = $.externalPolicies;
        this.memberEntityIds = $.memberEntityIds;
        this.memberGroupIds = $.memberGroupIds;
        this.metadata = $.metadata;
        this.name = $.name;
        this.namespace = $.namespace;
        this.policies = $.policies;
        this.type = $.type;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GroupState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GroupState $;

        public Builder() {
            $ = new GroupState();
        }

        public Builder(GroupState defaults) {
            $ = new GroupState(Objects.requireNonNull(defaults));
        }

        /**
         * @param externalMemberEntityIds `false` by default. If set to `true`, this resource will ignore any Entity IDs
         * returned from Vault or specified in the resource. You can use
         * `vault.identity.GroupMemberEntityIds` to manage Entity IDs for this group in a
         * decoupled manner.
         * 
         * @return builder
         * 
         */
        public Builder externalMemberEntityIds(@Nullable Output<Boolean> externalMemberEntityIds) {
            $.externalMemberEntityIds = externalMemberEntityIds;
            return this;
        }

        /**
         * @param externalMemberEntityIds `false` by default. If set to `true`, this resource will ignore any Entity IDs
         * returned from Vault or specified in the resource. You can use
         * `vault.identity.GroupMemberEntityIds` to manage Entity IDs for this group in a
         * decoupled manner.
         * 
         * @return builder
         * 
         */
        public Builder externalMemberEntityIds(Boolean externalMemberEntityIds) {
            return externalMemberEntityIds(Output.of(externalMemberEntityIds));
        }

        /**
         * @param externalMemberGroupIds `false` by default. If set to `true`, this resource will ignore any Group IDs
         * returned from Vault or specified in the resource. You can use
         * `vault.identity.GroupMemberGroupIds` to manage Group IDs for this group in a
         * decoupled manner.
         * 
         * @return builder
         * 
         */
        public Builder externalMemberGroupIds(@Nullable Output<Boolean> externalMemberGroupIds) {
            $.externalMemberGroupIds = externalMemberGroupIds;
            return this;
        }

        /**
         * @param externalMemberGroupIds `false` by default. If set to `true`, this resource will ignore any Group IDs
         * returned from Vault or specified in the resource. You can use
         * `vault.identity.GroupMemberGroupIds` to manage Group IDs for this group in a
         * decoupled manner.
         * 
         * @return builder
         * 
         */
        public Builder externalMemberGroupIds(Boolean externalMemberGroupIds) {
            return externalMemberGroupIds(Output.of(externalMemberGroupIds));
        }

        /**
         * @param externalPolicies `false` by default. If set to `true`, this resource will ignore any policies returned from
         * Vault or specified in the resource. You can use `vault.identity.GroupPolicies` to manage
         * policies for this group in a decoupled manner.
         * 
         * @return builder
         * 
         */
        public Builder externalPolicies(@Nullable Output<Boolean> externalPolicies) {
            $.externalPolicies = externalPolicies;
            return this;
        }

        /**
         * @param externalPolicies `false` by default. If set to `true`, this resource will ignore any policies returned from
         * Vault or specified in the resource. You can use `vault.identity.GroupPolicies` to manage
         * policies for this group in a decoupled manner.
         * 
         * @return builder
         * 
         */
        public Builder externalPolicies(Boolean externalPolicies) {
            return externalPolicies(Output.of(externalPolicies));
        }

        /**
         * @param memberEntityIds A list of Entity IDs to be assigned as group members. Not allowed on `external` groups.
         * 
         * @return builder
         * 
         */
        public Builder memberEntityIds(@Nullable Output<List<String>> memberEntityIds) {
            $.memberEntityIds = memberEntityIds;
            return this;
        }

        /**
         * @param memberEntityIds A list of Entity IDs to be assigned as group members. Not allowed on `external` groups.
         * 
         * @return builder
         * 
         */
        public Builder memberEntityIds(List<String> memberEntityIds) {
            return memberEntityIds(Output.of(memberEntityIds));
        }

        /**
         * @param memberEntityIds A list of Entity IDs to be assigned as group members. Not allowed on `external` groups.
         * 
         * @return builder
         * 
         */
        public Builder memberEntityIds(String... memberEntityIds) {
            return memberEntityIds(List.of(memberEntityIds));
        }

        /**
         * @param memberGroupIds A list of Group IDs to be assigned as group members. Not allowed on `external` groups.
         * 
         * @return builder
         * 
         */
        public Builder memberGroupIds(@Nullable Output<List<String>> memberGroupIds) {
            $.memberGroupIds = memberGroupIds;
            return this;
        }

        /**
         * @param memberGroupIds A list of Group IDs to be assigned as group members. Not allowed on `external` groups.
         * 
         * @return builder
         * 
         */
        public Builder memberGroupIds(List<String> memberGroupIds) {
            return memberGroupIds(Output.of(memberGroupIds));
        }

        /**
         * @param memberGroupIds A list of Group IDs to be assigned as group members. Not allowed on `external` groups.
         * 
         * @return builder
         * 
         */
        public Builder memberGroupIds(String... memberGroupIds) {
            return memberGroupIds(List.of(memberGroupIds));
        }

        /**
         * @param metadata A Map of additional metadata to associate with the group.
         * 
         * @return builder
         * 
         */
        public Builder metadata(@Nullable Output<Map<String,String>> metadata) {
            $.metadata = metadata;
            return this;
        }

        /**
         * @param metadata A Map of additional metadata to associate with the group.
         * 
         * @return builder
         * 
         */
        public Builder metadata(Map<String,String> metadata) {
            return metadata(Output.of(metadata));
        }

        /**
         * @param name Name of the identity group to create.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name Name of the identity group to create.
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
         * @param policies A list of policies to apply to the group.
         * 
         * @return builder
         * 
         */
        public Builder policies(@Nullable Output<List<String>> policies) {
            $.policies = policies;
            return this;
        }

        /**
         * @param policies A list of policies to apply to the group.
         * 
         * @return builder
         * 
         */
        public Builder policies(List<String> policies) {
            return policies(Output.of(policies));
        }

        /**
         * @param policies A list of policies to apply to the group.
         * 
         * @return builder
         * 
         */
        public Builder policies(String... policies) {
            return policies(List.of(policies));
        }

        /**
         * @param type Type of the group, internal or external. Defaults to `internal`.
         * 
         * @return builder
         * 
         */
        public Builder type(@Nullable Output<String> type) {
            $.type = type;
            return this;
        }

        /**
         * @param type Type of the group, internal or external. Defaults to `internal`.
         * 
         * @return builder
         * 
         */
        public Builder type(String type) {
            return type(Output.of(type));
        }

        public GroupState build() {
            return $;
        }
    }

}
