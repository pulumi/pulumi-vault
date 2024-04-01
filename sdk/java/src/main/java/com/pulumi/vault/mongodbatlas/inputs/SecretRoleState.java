// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vault.mongodbatlas.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class SecretRoleState extends com.pulumi.resources.ResourceArgs {

    public static final SecretRoleState Empty = new SecretRoleState();

    /**
     * Whitelist entry in CIDR notation to be added for the API key.
     * 
     */
    @Import(name="cidrBlocks")
    private @Nullable Output<List<String>> cidrBlocks;

    /**
     * @return Whitelist entry in CIDR notation to be added for the API key.
     * 
     */
    public Optional<Output<List<String>>> cidrBlocks() {
        return Optional.ofNullable(this.cidrBlocks);
    }

    /**
     * IP address to be added to the whitelist for the API key.
     * 
     */
    @Import(name="ipAddresses")
    private @Nullable Output<List<String>> ipAddresses;

    /**
     * @return IP address to be added to the whitelist for the API key.
     * 
     */
    public Optional<Output<List<String>>> ipAddresses() {
        return Optional.ofNullable(this.ipAddresses);
    }

    /**
     * The maximum allowed lifetime of credentials issued using this role.
     * 
     */
    @Import(name="maxTtl")
    private @Nullable Output<String> maxTtl;

    /**
     * @return The maximum allowed lifetime of credentials issued using this role.
     * 
     */
    public Optional<Output<String>> maxTtl() {
        return Optional.ofNullable(this.maxTtl);
    }

    /**
     * Path where the MongoDB Atlas Secrets Engine is mounted.
     * 
     */
    @Import(name="mount")
    private @Nullable Output<String> mount;

    /**
     * @return Path where the MongoDB Atlas Secrets Engine is mounted.
     * 
     */
    public Optional<Output<String>> mount() {
        return Optional.ofNullable(this.mount);
    }

    /**
     * The name of the role.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return The name of the role.
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
     * Unique identifier for the organization to which the target API Key belongs.
     * Required if `project_id` is not set.
     * 
     */
    @Import(name="organizationId")
    private @Nullable Output<String> organizationId;

    /**
     * @return Unique identifier for the organization to which the target API Key belongs.
     * Required if `project_id` is not set.
     * 
     */
    public Optional<Output<String>> organizationId() {
        return Optional.ofNullable(this.organizationId);
    }

    /**
     * Unique identifier for the project to which the target API Key belongs.
     * Required if `organization_id` is not set.
     * 
     */
    @Import(name="projectId")
    private @Nullable Output<String> projectId;

    /**
     * @return Unique identifier for the project to which the target API Key belongs.
     * Required if `organization_id` is not set.
     * 
     */
    public Optional<Output<String>> projectId() {
        return Optional.ofNullable(this.projectId);
    }

    /**
     * Roles assigned when an org API key is assigned to a project API key. Possible values are `GROUP_CLUSTER_MANAGER`, `GROUP_DATA_ACCESS_ADMIN`, `GROUP_DATA_ACCESS_READ_ONLY`, `GROUP_DATA_ACCESS_READ_WRITE`, `GROUP_OWNER` and `GROUP_READ_ONLY`.
     * 
     */
    @Import(name="projectRoles")
    private @Nullable Output<List<String>> projectRoles;

    /**
     * @return Roles assigned when an org API key is assigned to a project API key. Possible values are `GROUP_CLUSTER_MANAGER`, `GROUP_DATA_ACCESS_ADMIN`, `GROUP_DATA_ACCESS_READ_ONLY`, `GROUP_DATA_ACCESS_READ_WRITE`, `GROUP_OWNER` and `GROUP_READ_ONLY`.
     * 
     */
    public Optional<Output<List<String>>> projectRoles() {
        return Optional.ofNullable(this.projectRoles);
    }

    /**
     * List of roles that the API Key needs to have. Possible values are `ORG_OWNER`, `ORG_MEMBER`, `ORG_GROUP_CREATOR`, `ORG_BILLING_ADMIN` and `ORG_READ_ONLY`.
     * 
     */
    @Import(name="roles")
    private @Nullable Output<List<String>> roles;

    /**
     * @return List of roles that the API Key needs to have. Possible values are `ORG_OWNER`, `ORG_MEMBER`, `ORG_GROUP_CREATOR`, `ORG_BILLING_ADMIN` and `ORG_READ_ONLY`.
     * 
     */
    public Optional<Output<List<String>>> roles() {
        return Optional.ofNullable(this.roles);
    }

    /**
     * Duration in seconds after which the issued credential should expire.
     * 
     */
    @Import(name="ttl")
    private @Nullable Output<String> ttl;

    /**
     * @return Duration in seconds after which the issued credential should expire.
     * 
     */
    public Optional<Output<String>> ttl() {
        return Optional.ofNullable(this.ttl);
    }

    private SecretRoleState() {}

    private SecretRoleState(SecretRoleState $) {
        this.cidrBlocks = $.cidrBlocks;
        this.ipAddresses = $.ipAddresses;
        this.maxTtl = $.maxTtl;
        this.mount = $.mount;
        this.name = $.name;
        this.namespace = $.namespace;
        this.organizationId = $.organizationId;
        this.projectId = $.projectId;
        this.projectRoles = $.projectRoles;
        this.roles = $.roles;
        this.ttl = $.ttl;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(SecretRoleState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private SecretRoleState $;

        public Builder() {
            $ = new SecretRoleState();
        }

        public Builder(SecretRoleState defaults) {
            $ = new SecretRoleState(Objects.requireNonNull(defaults));
        }

        /**
         * @param cidrBlocks Whitelist entry in CIDR notation to be added for the API key.
         * 
         * @return builder
         * 
         */
        public Builder cidrBlocks(@Nullable Output<List<String>> cidrBlocks) {
            $.cidrBlocks = cidrBlocks;
            return this;
        }

        /**
         * @param cidrBlocks Whitelist entry in CIDR notation to be added for the API key.
         * 
         * @return builder
         * 
         */
        public Builder cidrBlocks(List<String> cidrBlocks) {
            return cidrBlocks(Output.of(cidrBlocks));
        }

        /**
         * @param cidrBlocks Whitelist entry in CIDR notation to be added for the API key.
         * 
         * @return builder
         * 
         */
        public Builder cidrBlocks(String... cidrBlocks) {
            return cidrBlocks(List.of(cidrBlocks));
        }

        /**
         * @param ipAddresses IP address to be added to the whitelist for the API key.
         * 
         * @return builder
         * 
         */
        public Builder ipAddresses(@Nullable Output<List<String>> ipAddresses) {
            $.ipAddresses = ipAddresses;
            return this;
        }

        /**
         * @param ipAddresses IP address to be added to the whitelist for the API key.
         * 
         * @return builder
         * 
         */
        public Builder ipAddresses(List<String> ipAddresses) {
            return ipAddresses(Output.of(ipAddresses));
        }

        /**
         * @param ipAddresses IP address to be added to the whitelist for the API key.
         * 
         * @return builder
         * 
         */
        public Builder ipAddresses(String... ipAddresses) {
            return ipAddresses(List.of(ipAddresses));
        }

        /**
         * @param maxTtl The maximum allowed lifetime of credentials issued using this role.
         * 
         * @return builder
         * 
         */
        public Builder maxTtl(@Nullable Output<String> maxTtl) {
            $.maxTtl = maxTtl;
            return this;
        }

        /**
         * @param maxTtl The maximum allowed lifetime of credentials issued using this role.
         * 
         * @return builder
         * 
         */
        public Builder maxTtl(String maxTtl) {
            return maxTtl(Output.of(maxTtl));
        }

        /**
         * @param mount Path where the MongoDB Atlas Secrets Engine is mounted.
         * 
         * @return builder
         * 
         */
        public Builder mount(@Nullable Output<String> mount) {
            $.mount = mount;
            return this;
        }

        /**
         * @param mount Path where the MongoDB Atlas Secrets Engine is mounted.
         * 
         * @return builder
         * 
         */
        public Builder mount(String mount) {
            return mount(Output.of(mount));
        }

        /**
         * @param name The name of the role.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name The name of the role.
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
         * @param organizationId Unique identifier for the organization to which the target API Key belongs.
         * Required if `project_id` is not set.
         * 
         * @return builder
         * 
         */
        public Builder organizationId(@Nullable Output<String> organizationId) {
            $.organizationId = organizationId;
            return this;
        }

        /**
         * @param organizationId Unique identifier for the organization to which the target API Key belongs.
         * Required if `project_id` is not set.
         * 
         * @return builder
         * 
         */
        public Builder organizationId(String organizationId) {
            return organizationId(Output.of(organizationId));
        }

        /**
         * @param projectId Unique identifier for the project to which the target API Key belongs.
         * Required if `organization_id` is not set.
         * 
         * @return builder
         * 
         */
        public Builder projectId(@Nullable Output<String> projectId) {
            $.projectId = projectId;
            return this;
        }

        /**
         * @param projectId Unique identifier for the project to which the target API Key belongs.
         * Required if `organization_id` is not set.
         * 
         * @return builder
         * 
         */
        public Builder projectId(String projectId) {
            return projectId(Output.of(projectId));
        }

        /**
         * @param projectRoles Roles assigned when an org API key is assigned to a project API key. Possible values are `GROUP_CLUSTER_MANAGER`, `GROUP_DATA_ACCESS_ADMIN`, `GROUP_DATA_ACCESS_READ_ONLY`, `GROUP_DATA_ACCESS_READ_WRITE`, `GROUP_OWNER` and `GROUP_READ_ONLY`.
         * 
         * @return builder
         * 
         */
        public Builder projectRoles(@Nullable Output<List<String>> projectRoles) {
            $.projectRoles = projectRoles;
            return this;
        }

        /**
         * @param projectRoles Roles assigned when an org API key is assigned to a project API key. Possible values are `GROUP_CLUSTER_MANAGER`, `GROUP_DATA_ACCESS_ADMIN`, `GROUP_DATA_ACCESS_READ_ONLY`, `GROUP_DATA_ACCESS_READ_WRITE`, `GROUP_OWNER` and `GROUP_READ_ONLY`.
         * 
         * @return builder
         * 
         */
        public Builder projectRoles(List<String> projectRoles) {
            return projectRoles(Output.of(projectRoles));
        }

        /**
         * @param projectRoles Roles assigned when an org API key is assigned to a project API key. Possible values are `GROUP_CLUSTER_MANAGER`, `GROUP_DATA_ACCESS_ADMIN`, `GROUP_DATA_ACCESS_READ_ONLY`, `GROUP_DATA_ACCESS_READ_WRITE`, `GROUP_OWNER` and `GROUP_READ_ONLY`.
         * 
         * @return builder
         * 
         */
        public Builder projectRoles(String... projectRoles) {
            return projectRoles(List.of(projectRoles));
        }

        /**
         * @param roles List of roles that the API Key needs to have. Possible values are `ORG_OWNER`, `ORG_MEMBER`, `ORG_GROUP_CREATOR`, `ORG_BILLING_ADMIN` and `ORG_READ_ONLY`.
         * 
         * @return builder
         * 
         */
        public Builder roles(@Nullable Output<List<String>> roles) {
            $.roles = roles;
            return this;
        }

        /**
         * @param roles List of roles that the API Key needs to have. Possible values are `ORG_OWNER`, `ORG_MEMBER`, `ORG_GROUP_CREATOR`, `ORG_BILLING_ADMIN` and `ORG_READ_ONLY`.
         * 
         * @return builder
         * 
         */
        public Builder roles(List<String> roles) {
            return roles(Output.of(roles));
        }

        /**
         * @param roles List of roles that the API Key needs to have. Possible values are `ORG_OWNER`, `ORG_MEMBER`, `ORG_GROUP_CREATOR`, `ORG_BILLING_ADMIN` and `ORG_READ_ONLY`.
         * 
         * @return builder
         * 
         */
        public Builder roles(String... roles) {
            return roles(List.of(roles));
        }

        /**
         * @param ttl Duration in seconds after which the issued credential should expire.
         * 
         * @return builder
         * 
         */
        public Builder ttl(@Nullable Output<String> ttl) {
            $.ttl = ttl;
            return this;
        }

        /**
         * @param ttl Duration in seconds after which the issued credential should expire.
         * 
         * @return builder
         * 
         */
        public Builder ttl(String ttl) {
            return ttl(Output.of(ttl));
        }

        public SecretRoleState build() {
            return $;
        }
    }

}
