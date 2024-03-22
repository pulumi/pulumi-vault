// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vault.ldap.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class SecretBackendDynamicRoleState extends com.pulumi.resources.ResourceArgs {

    public static final SecretBackendDynamicRoleState Empty = new SecretBackendDynamicRoleState();

    /**
     * A templatized LDIF string used to create a user
     * account. This may contain multiple LDIF entries. The `creation_ldif` can also
     * be used to add the user account to an existing group. All LDIF entries are
     * performed in order. If Vault encounters an error while executing the
     * `creation_ldif` it will stop at the first error and not execute any remaining
     * LDIF entries. If an error occurs and `rollback_ldif` is specified, the LDIF
     * entries in `rollback_ldif` will be executed. See `rollback_ldif` for more
     * details. This field may optionally be provided as a base64 encoded string.
     * 
     */
    @Import(name="creationLdif")
    private @Nullable Output<String> creationLdif;

    /**
     * @return A templatized LDIF string used to create a user
     * account. This may contain multiple LDIF entries. The `creation_ldif` can also
     * be used to add the user account to an existing group. All LDIF entries are
     * performed in order. If Vault encounters an error while executing the
     * `creation_ldif` it will stop at the first error and not execute any remaining
     * LDIF entries. If an error occurs and `rollback_ldif` is specified, the LDIF
     * entries in `rollback_ldif` will be executed. See `rollback_ldif` for more
     * details. This field may optionally be provided as a base64 encoded string.
     * 
     */
    public Optional<Output<String>> creationLdif() {
        return Optional.ofNullable(this.creationLdif);
    }

    /**
     * Specifies the TTL for the leases associated with this role.
     * 
     */
    @Import(name="defaultTtl")
    private @Nullable Output<Integer> defaultTtl;

    /**
     * @return Specifies the TTL for the leases associated with this role.
     * 
     */
    public Optional<Output<Integer>> defaultTtl() {
        return Optional.ofNullable(this.defaultTtl);
    }

    /**
     * A templatized LDIF string used to delete the
     * user account once its TTL has expired. This may contain multiple LDIF
     * entries. All LDIF entries are performed in order. If Vault encounters an
     * error while executing an entry in the `deletion_ldif` it will attempt to
     * continue executing any remaining entries. This field may optionally be
     * provided as a base64 encoded string.
     * 
     */
    @Import(name="deletionLdif")
    private @Nullable Output<String> deletionLdif;

    /**
     * @return A templatized LDIF string used to delete the
     * user account once its TTL has expired. This may contain multiple LDIF
     * entries. All LDIF entries are performed in order. If Vault encounters an
     * error while executing an entry in the `deletion_ldif` it will attempt to
     * continue executing any remaining entries. This field may optionally be
     * provided as a base64 encoded string.
     * 
     */
    public Optional<Output<String>> deletionLdif() {
        return Optional.ofNullable(this.deletionLdif);
    }

    /**
     * Specifies the maximum TTL for the leases associated with this role.
     * 
     */
    @Import(name="maxTtl")
    private @Nullable Output<Integer> maxTtl;

    /**
     * @return Specifies the maximum TTL for the leases associated with this role.
     * 
     */
    public Optional<Output<Integer>> maxTtl() {
        return Optional.ofNullable(this.maxTtl);
    }

    /**
     * The unique path this backend should be mounted at. Must
     * not begin or end with a `/`. Defaults to `ldap`.
     * 
     */
    @Import(name="mount")
    private @Nullable Output<String> mount;

    /**
     * @return The unique path this backend should be mounted at. Must
     * not begin or end with a `/`. Defaults to `ldap`.
     * 
     */
    public Optional<Output<String>> mount() {
        return Optional.ofNullable(this.mount);
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
     * Name of the role.
     * 
     */
    @Import(name="roleName")
    private @Nullable Output<String> roleName;

    /**
     * @return Name of the role.
     * 
     */
    public Optional<Output<String>> roleName() {
        return Optional.ofNullable(this.roleName);
    }

    /**
     * A templatized LDIF string used to attempt to
     * rollback any changes in the event that execution of the `creation_ldif` results
     * in an error. This may contain multiple LDIF entries. All LDIF entries are
     * performed in order. If Vault encounters an error while executing an entry in
     * the `rollback_ldif` it will attempt to continue executing any remaining
     * entries. This field may optionally be provided as a base64 encoded string.
     * 
     */
    @Import(name="rollbackLdif")
    private @Nullable Output<String> rollbackLdif;

    /**
     * @return A templatized LDIF string used to attempt to
     * rollback any changes in the event that execution of the `creation_ldif` results
     * in an error. This may contain multiple LDIF entries. All LDIF entries are
     * performed in order. If Vault encounters an error while executing an entry in
     * the `rollback_ldif` it will attempt to continue executing any remaining
     * entries. This field may optionally be provided as a base64 encoded string.
     * 
     */
    public Optional<Output<String>> rollbackLdif() {
        return Optional.ofNullable(this.rollbackLdif);
    }

    /**
     * A template used to generate a dynamic
     * username. This will be used to fill in the `.Username` field within the
     * `creation_ldif` string.
     * 
     */
    @Import(name="usernameTemplate")
    private @Nullable Output<String> usernameTemplate;

    /**
     * @return A template used to generate a dynamic
     * username. This will be used to fill in the `.Username` field within the
     * `creation_ldif` string.
     * 
     */
    public Optional<Output<String>> usernameTemplate() {
        return Optional.ofNullable(this.usernameTemplate);
    }

    private SecretBackendDynamicRoleState() {}

    private SecretBackendDynamicRoleState(SecretBackendDynamicRoleState $) {
        this.creationLdif = $.creationLdif;
        this.defaultTtl = $.defaultTtl;
        this.deletionLdif = $.deletionLdif;
        this.maxTtl = $.maxTtl;
        this.mount = $.mount;
        this.namespace = $.namespace;
        this.roleName = $.roleName;
        this.rollbackLdif = $.rollbackLdif;
        this.usernameTemplate = $.usernameTemplate;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(SecretBackendDynamicRoleState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private SecretBackendDynamicRoleState $;

        public Builder() {
            $ = new SecretBackendDynamicRoleState();
        }

        public Builder(SecretBackendDynamicRoleState defaults) {
            $ = new SecretBackendDynamicRoleState(Objects.requireNonNull(defaults));
        }

        /**
         * @param creationLdif A templatized LDIF string used to create a user
         * account. This may contain multiple LDIF entries. The `creation_ldif` can also
         * be used to add the user account to an existing group. All LDIF entries are
         * performed in order. If Vault encounters an error while executing the
         * `creation_ldif` it will stop at the first error and not execute any remaining
         * LDIF entries. If an error occurs and `rollback_ldif` is specified, the LDIF
         * entries in `rollback_ldif` will be executed. See `rollback_ldif` for more
         * details. This field may optionally be provided as a base64 encoded string.
         * 
         * @return builder
         * 
         */
        public Builder creationLdif(@Nullable Output<String> creationLdif) {
            $.creationLdif = creationLdif;
            return this;
        }

        /**
         * @param creationLdif A templatized LDIF string used to create a user
         * account. This may contain multiple LDIF entries. The `creation_ldif` can also
         * be used to add the user account to an existing group. All LDIF entries are
         * performed in order. If Vault encounters an error while executing the
         * `creation_ldif` it will stop at the first error and not execute any remaining
         * LDIF entries. If an error occurs and `rollback_ldif` is specified, the LDIF
         * entries in `rollback_ldif` will be executed. See `rollback_ldif` for more
         * details. This field may optionally be provided as a base64 encoded string.
         * 
         * @return builder
         * 
         */
        public Builder creationLdif(String creationLdif) {
            return creationLdif(Output.of(creationLdif));
        }

        /**
         * @param defaultTtl Specifies the TTL for the leases associated with this role.
         * 
         * @return builder
         * 
         */
        public Builder defaultTtl(@Nullable Output<Integer> defaultTtl) {
            $.defaultTtl = defaultTtl;
            return this;
        }

        /**
         * @param defaultTtl Specifies the TTL for the leases associated with this role.
         * 
         * @return builder
         * 
         */
        public Builder defaultTtl(Integer defaultTtl) {
            return defaultTtl(Output.of(defaultTtl));
        }

        /**
         * @param deletionLdif A templatized LDIF string used to delete the
         * user account once its TTL has expired. This may contain multiple LDIF
         * entries. All LDIF entries are performed in order. If Vault encounters an
         * error while executing an entry in the `deletion_ldif` it will attempt to
         * continue executing any remaining entries. This field may optionally be
         * provided as a base64 encoded string.
         * 
         * @return builder
         * 
         */
        public Builder deletionLdif(@Nullable Output<String> deletionLdif) {
            $.deletionLdif = deletionLdif;
            return this;
        }

        /**
         * @param deletionLdif A templatized LDIF string used to delete the
         * user account once its TTL has expired. This may contain multiple LDIF
         * entries. All LDIF entries are performed in order. If Vault encounters an
         * error while executing an entry in the `deletion_ldif` it will attempt to
         * continue executing any remaining entries. This field may optionally be
         * provided as a base64 encoded string.
         * 
         * @return builder
         * 
         */
        public Builder deletionLdif(String deletionLdif) {
            return deletionLdif(Output.of(deletionLdif));
        }

        /**
         * @param maxTtl Specifies the maximum TTL for the leases associated with this role.
         * 
         * @return builder
         * 
         */
        public Builder maxTtl(@Nullable Output<Integer> maxTtl) {
            $.maxTtl = maxTtl;
            return this;
        }

        /**
         * @param maxTtl Specifies the maximum TTL for the leases associated with this role.
         * 
         * @return builder
         * 
         */
        public Builder maxTtl(Integer maxTtl) {
            return maxTtl(Output.of(maxTtl));
        }

        /**
         * @param mount The unique path this backend should be mounted at. Must
         * not begin or end with a `/`. Defaults to `ldap`.
         * 
         * @return builder
         * 
         */
        public Builder mount(@Nullable Output<String> mount) {
            $.mount = mount;
            return this;
        }

        /**
         * @param mount The unique path this backend should be mounted at. Must
         * not begin or end with a `/`. Defaults to `ldap`.
         * 
         * @return builder
         * 
         */
        public Builder mount(String mount) {
            return mount(Output.of(mount));
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
         * @param roleName Name of the role.
         * 
         * @return builder
         * 
         */
        public Builder roleName(@Nullable Output<String> roleName) {
            $.roleName = roleName;
            return this;
        }

        /**
         * @param roleName Name of the role.
         * 
         * @return builder
         * 
         */
        public Builder roleName(String roleName) {
            return roleName(Output.of(roleName));
        }

        /**
         * @param rollbackLdif A templatized LDIF string used to attempt to
         * rollback any changes in the event that execution of the `creation_ldif` results
         * in an error. This may contain multiple LDIF entries. All LDIF entries are
         * performed in order. If Vault encounters an error while executing an entry in
         * the `rollback_ldif` it will attempt to continue executing any remaining
         * entries. This field may optionally be provided as a base64 encoded string.
         * 
         * @return builder
         * 
         */
        public Builder rollbackLdif(@Nullable Output<String> rollbackLdif) {
            $.rollbackLdif = rollbackLdif;
            return this;
        }

        /**
         * @param rollbackLdif A templatized LDIF string used to attempt to
         * rollback any changes in the event that execution of the `creation_ldif` results
         * in an error. This may contain multiple LDIF entries. All LDIF entries are
         * performed in order. If Vault encounters an error while executing an entry in
         * the `rollback_ldif` it will attempt to continue executing any remaining
         * entries. This field may optionally be provided as a base64 encoded string.
         * 
         * @return builder
         * 
         */
        public Builder rollbackLdif(String rollbackLdif) {
            return rollbackLdif(Output.of(rollbackLdif));
        }

        /**
         * @param usernameTemplate A template used to generate a dynamic
         * username. This will be used to fill in the `.Username` field within the
         * `creation_ldif` string.
         * 
         * @return builder
         * 
         */
        public Builder usernameTemplate(@Nullable Output<String> usernameTemplate) {
            $.usernameTemplate = usernameTemplate;
            return this;
        }

        /**
         * @param usernameTemplate A template used to generate a dynamic
         * username. This will be used to fill in the `.Username` field within the
         * `creation_ldif` string.
         * 
         * @return builder
         * 
         */
        public Builder usernameTemplate(String usernameTemplate) {
            return usernameTemplate(Output.of(usernameTemplate));
        }

        public SecretBackendDynamicRoleState build() {
            return $;
        }
    }

}
