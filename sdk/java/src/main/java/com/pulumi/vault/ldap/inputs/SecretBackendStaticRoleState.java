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


public final class SecretBackendStaticRoleState extends com.pulumi.resources.ResourceArgs {

    public static final SecretBackendStaticRoleState Empty = new SecretBackendStaticRoleState();

    /**
     * Distinguished name (DN) of the existing LDAP entry to manage
     * password rotation for. If given, it will take precedence over `username` for the LDAP
     * search performed during password rotation. Cannot be modified after creation.
     * 
     */
    @Import(name="dn")
    private @Nullable Output<String> dn;

    /**
     * @return Distinguished name (DN) of the existing LDAP entry to manage
     * password rotation for. If given, it will take precedence over `username` for the LDAP
     * search performed during password rotation. Cannot be modified after creation.
     * 
     */
    public Optional<Output<String>> dn() {
        return Optional.ofNullable(this.dn);
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
     * How often Vault should rotate the password of the user entry.
     * 
     */
    @Import(name="rotationPeriod")
    private @Nullable Output<Integer> rotationPeriod;

    /**
     * @return How often Vault should rotate the password of the user entry.
     * 
     */
    public Optional<Output<Integer>> rotationPeriod() {
        return Optional.ofNullable(this.rotationPeriod);
    }

    /**
     * The username of the existing LDAP entry to manage password rotation for.
     * 
     */
    @Import(name="username")
    private @Nullable Output<String> username;

    /**
     * @return The username of the existing LDAP entry to manage password rotation for.
     * 
     */
    public Optional<Output<String>> username() {
        return Optional.ofNullable(this.username);
    }

    private SecretBackendStaticRoleState() {}

    private SecretBackendStaticRoleState(SecretBackendStaticRoleState $) {
        this.dn = $.dn;
        this.mount = $.mount;
        this.namespace = $.namespace;
        this.roleName = $.roleName;
        this.rotationPeriod = $.rotationPeriod;
        this.username = $.username;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(SecretBackendStaticRoleState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private SecretBackendStaticRoleState $;

        public Builder() {
            $ = new SecretBackendStaticRoleState();
        }

        public Builder(SecretBackendStaticRoleState defaults) {
            $ = new SecretBackendStaticRoleState(Objects.requireNonNull(defaults));
        }

        /**
         * @param dn Distinguished name (DN) of the existing LDAP entry to manage
         * password rotation for. If given, it will take precedence over `username` for the LDAP
         * search performed during password rotation. Cannot be modified after creation.
         * 
         * @return builder
         * 
         */
        public Builder dn(@Nullable Output<String> dn) {
            $.dn = dn;
            return this;
        }

        /**
         * @param dn Distinguished name (DN) of the existing LDAP entry to manage
         * password rotation for. If given, it will take precedence over `username` for the LDAP
         * search performed during password rotation. Cannot be modified after creation.
         * 
         * @return builder
         * 
         */
        public Builder dn(String dn) {
            return dn(Output.of(dn));
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
         * @param rotationPeriod How often Vault should rotate the password of the user entry.
         * 
         * @return builder
         * 
         */
        public Builder rotationPeriod(@Nullable Output<Integer> rotationPeriod) {
            $.rotationPeriod = rotationPeriod;
            return this;
        }

        /**
         * @param rotationPeriod How often Vault should rotate the password of the user entry.
         * 
         * @return builder
         * 
         */
        public Builder rotationPeriod(Integer rotationPeriod) {
            return rotationPeriod(Output.of(rotationPeriod));
        }

        /**
         * @param username The username of the existing LDAP entry to manage password rotation for.
         * 
         * @return builder
         * 
         */
        public Builder username(@Nullable Output<String> username) {
            $.username = username;
            return this;
        }

        /**
         * @param username The username of the existing LDAP entry to manage password rotation for.
         * 
         * @return builder
         * 
         */
        public Builder username(String username) {
            return username(Output.of(username));
        }

        public SecretBackendStaticRoleState build() {
            return $;
        }
    }

}
