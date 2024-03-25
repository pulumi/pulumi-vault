// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vault.ldap;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class AuthBackendGroupArgs extends com.pulumi.resources.ResourceArgs {

    public static final AuthBackendGroupArgs Empty = new AuthBackendGroupArgs();

    /**
     * Path to the authentication backend
     * 
     * For more details on the usage of each argument consult the [Vault LDAP API documentation](https://www.vaultproject.io/api-docs/auth/ldap).
     * 
     */
    @Import(name="backend")
    private @Nullable Output<String> backend;

    /**
     * @return Path to the authentication backend
     * 
     * For more details on the usage of each argument consult the [Vault LDAP API documentation](https://www.vaultproject.io/api-docs/auth/ldap).
     * 
     */
    public Optional<Output<String>> backend() {
        return Optional.ofNullable(this.backend);
    }

    /**
     * The LDAP groupname
     * 
     */
    @Import(name="groupname", required=true)
    private Output<String> groupname;

    /**
     * @return The LDAP groupname
     * 
     */
    public Output<String> groupname() {
        return this.groupname;
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
     * Policies which should be granted to members of the group
     * 
     */
    @Import(name="policies")
    private @Nullable Output<List<String>> policies;

    /**
     * @return Policies which should be granted to members of the group
     * 
     */
    public Optional<Output<List<String>>> policies() {
        return Optional.ofNullable(this.policies);
    }

    private AuthBackendGroupArgs() {}

    private AuthBackendGroupArgs(AuthBackendGroupArgs $) {
        this.backend = $.backend;
        this.groupname = $.groupname;
        this.namespace = $.namespace;
        this.policies = $.policies;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(AuthBackendGroupArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private AuthBackendGroupArgs $;

        public Builder() {
            $ = new AuthBackendGroupArgs();
        }

        public Builder(AuthBackendGroupArgs defaults) {
            $ = new AuthBackendGroupArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param backend Path to the authentication backend
         * 
         * For more details on the usage of each argument consult the [Vault LDAP API documentation](https://www.vaultproject.io/api-docs/auth/ldap).
         * 
         * @return builder
         * 
         */
        public Builder backend(@Nullable Output<String> backend) {
            $.backend = backend;
            return this;
        }

        /**
         * @param backend Path to the authentication backend
         * 
         * For more details on the usage of each argument consult the [Vault LDAP API documentation](https://www.vaultproject.io/api-docs/auth/ldap).
         * 
         * @return builder
         * 
         */
        public Builder backend(String backend) {
            return backend(Output.of(backend));
        }

        /**
         * @param groupname The LDAP groupname
         * 
         * @return builder
         * 
         */
        public Builder groupname(Output<String> groupname) {
            $.groupname = groupname;
            return this;
        }

        /**
         * @param groupname The LDAP groupname
         * 
         * @return builder
         * 
         */
        public Builder groupname(String groupname) {
            return groupname(Output.of(groupname));
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
         * @param policies Policies which should be granted to members of the group
         * 
         * @return builder
         * 
         */
        public Builder policies(@Nullable Output<List<String>> policies) {
            $.policies = policies;
            return this;
        }

        /**
         * @param policies Policies which should be granted to members of the group
         * 
         * @return builder
         * 
         */
        public Builder policies(List<String> policies) {
            return policies(Output.of(policies));
        }

        /**
         * @param policies Policies which should be granted to members of the group
         * 
         * @return builder
         * 
         */
        public Builder policies(String... policies) {
            return policies(List.of(policies));
        }

        public AuthBackendGroupArgs build() {
            if ($.groupname == null) {
                throw new MissingRequiredPropertyException("AuthBackendGroupArgs", "groupname");
            }
            return $;
        }
    }

}
