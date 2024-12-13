// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vault.appRole;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class AuthBackendRoleArgs extends com.pulumi.resources.ResourceArgs {

    public static final AuthBackendRoleArgs Empty = new AuthBackendRoleArgs();

    /**
     * The unique name of the auth backend to configure.
     * Defaults to `approle`.
     * 
     */
    @Import(name="backend")
    private @Nullable Output<String> backend;

    /**
     * @return The unique name of the auth backend to configure.
     * Defaults to `approle`.
     * 
     */
    public Optional<Output<String>> backend() {
        return Optional.ofNullable(this.backend);
    }

    /**
     * Whether or not to require `secret_id` to be
     * presented when logging in using this AppRole. Defaults to `true`.
     * 
     */
    @Import(name="bindSecretId")
    private @Nullable Output<Boolean> bindSecretId;

    /**
     * @return Whether or not to require `secret_id` to be
     * presented when logging in using this AppRole. Defaults to `true`.
     * 
     */
    public Optional<Output<Boolean>> bindSecretId() {
        return Optional.ofNullable(this.bindSecretId);
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
     * The RoleID of this role. If not specified, one will be
     * auto-generated.
     * 
     */
    @Import(name="roleId")
    private @Nullable Output<String> roleId;

    /**
     * @return The RoleID of this role. If not specified, one will be
     * auto-generated.
     * 
     */
    public Optional<Output<String>> roleId() {
        return Optional.ofNullable(this.roleId);
    }

    /**
     * The name of the role.
     * 
     */
    @Import(name="roleName", required=true)
    private Output<String> roleName;

    /**
     * @return The name of the role.
     * 
     */
    public Output<String> roleName() {
        return this.roleName;
    }

    /**
     * If set,
     * specifies blocks of IP addresses which can perform the login operation.
     * 
     */
    @Import(name="secretIdBoundCidrs")
    private @Nullable Output<List<String>> secretIdBoundCidrs;

    /**
     * @return If set,
     * specifies blocks of IP addresses which can perform the login operation.
     * 
     */
    public Optional<Output<List<String>>> secretIdBoundCidrs() {
        return Optional.ofNullable(this.secretIdBoundCidrs);
    }

    /**
     * The number of times any particular SecretID
     * can be used to fetch a token from this AppRole, after which the SecretID will
     * expire. A value of zero will allow unlimited uses.
     * 
     */
    @Import(name="secretIdNumUses")
    private @Nullable Output<Integer> secretIdNumUses;

    /**
     * @return The number of times any particular SecretID
     * can be used to fetch a token from this AppRole, after which the SecretID will
     * expire. A value of zero will allow unlimited uses.
     * 
     */
    public Optional<Output<Integer>> secretIdNumUses() {
        return Optional.ofNullable(this.secretIdNumUses);
    }

    /**
     * The number of seconds after which any SecretID
     * expires.
     * 
     */
    @Import(name="secretIdTtl")
    private @Nullable Output<Integer> secretIdTtl;

    /**
     * @return The number of seconds after which any SecretID
     * expires.
     * 
     */
    public Optional<Output<Integer>> secretIdTtl() {
        return Optional.ofNullable(this.secretIdTtl);
    }

    /**
     * Specifies the blocks of IP addresses which are allowed to use the generated token
     * 
     */
    @Import(name="tokenBoundCidrs")
    private @Nullable Output<List<String>> tokenBoundCidrs;

    /**
     * @return Specifies the blocks of IP addresses which are allowed to use the generated token
     * 
     */
    public Optional<Output<List<String>>> tokenBoundCidrs() {
        return Optional.ofNullable(this.tokenBoundCidrs);
    }

    /**
     * Generated Token&#39;s Explicit Maximum TTL in seconds
     * 
     */
    @Import(name="tokenExplicitMaxTtl")
    private @Nullable Output<Integer> tokenExplicitMaxTtl;

    /**
     * @return Generated Token&#39;s Explicit Maximum TTL in seconds
     * 
     */
    public Optional<Output<Integer>> tokenExplicitMaxTtl() {
        return Optional.ofNullable(this.tokenExplicitMaxTtl);
    }

    /**
     * The maximum lifetime of the generated token
     * 
     */
    @Import(name="tokenMaxTtl")
    private @Nullable Output<Integer> tokenMaxTtl;

    /**
     * @return The maximum lifetime of the generated token
     * 
     */
    public Optional<Output<Integer>> tokenMaxTtl() {
        return Optional.ofNullable(this.tokenMaxTtl);
    }

    /**
     * If true, the &#39;default&#39; policy will not automatically be added to generated tokens
     * 
     */
    @Import(name="tokenNoDefaultPolicy")
    private @Nullable Output<Boolean> tokenNoDefaultPolicy;

    /**
     * @return If true, the &#39;default&#39; policy will not automatically be added to generated tokens
     * 
     */
    public Optional<Output<Boolean>> tokenNoDefaultPolicy() {
        return Optional.ofNullable(this.tokenNoDefaultPolicy);
    }

    /**
     * The maximum number of times a token may be used, a value of zero means unlimited
     * 
     */
    @Import(name="tokenNumUses")
    private @Nullable Output<Integer> tokenNumUses;

    /**
     * @return The maximum number of times a token may be used, a value of zero means unlimited
     * 
     */
    public Optional<Output<Integer>> tokenNumUses() {
        return Optional.ofNullable(this.tokenNumUses);
    }

    /**
     * Generated Token&#39;s Period
     * 
     */
    @Import(name="tokenPeriod")
    private @Nullable Output<Integer> tokenPeriod;

    /**
     * @return Generated Token&#39;s Period
     * 
     */
    public Optional<Output<Integer>> tokenPeriod() {
        return Optional.ofNullable(this.tokenPeriod);
    }

    /**
     * Generated Token&#39;s Policies
     * 
     */
    @Import(name="tokenPolicies")
    private @Nullable Output<List<String>> tokenPolicies;

    /**
     * @return Generated Token&#39;s Policies
     * 
     */
    public Optional<Output<List<String>>> tokenPolicies() {
        return Optional.ofNullable(this.tokenPolicies);
    }

    /**
     * The initial ttl of the token to generate in seconds
     * 
     */
    @Import(name="tokenTtl")
    private @Nullable Output<Integer> tokenTtl;

    /**
     * @return The initial ttl of the token to generate in seconds
     * 
     */
    public Optional<Output<Integer>> tokenTtl() {
        return Optional.ofNullable(this.tokenTtl);
    }

    /**
     * The type of token to generate, service or batch
     * 
     */
    @Import(name="tokenType")
    private @Nullable Output<String> tokenType;

    /**
     * @return The type of token to generate, service or batch
     * 
     */
    public Optional<Output<String>> tokenType() {
        return Optional.ofNullable(this.tokenType);
    }

    private AuthBackendRoleArgs() {}

    private AuthBackendRoleArgs(AuthBackendRoleArgs $) {
        this.backend = $.backend;
        this.bindSecretId = $.bindSecretId;
        this.namespace = $.namespace;
        this.roleId = $.roleId;
        this.roleName = $.roleName;
        this.secretIdBoundCidrs = $.secretIdBoundCidrs;
        this.secretIdNumUses = $.secretIdNumUses;
        this.secretIdTtl = $.secretIdTtl;
        this.tokenBoundCidrs = $.tokenBoundCidrs;
        this.tokenExplicitMaxTtl = $.tokenExplicitMaxTtl;
        this.tokenMaxTtl = $.tokenMaxTtl;
        this.tokenNoDefaultPolicy = $.tokenNoDefaultPolicy;
        this.tokenNumUses = $.tokenNumUses;
        this.tokenPeriod = $.tokenPeriod;
        this.tokenPolicies = $.tokenPolicies;
        this.tokenTtl = $.tokenTtl;
        this.tokenType = $.tokenType;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(AuthBackendRoleArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private AuthBackendRoleArgs $;

        public Builder() {
            $ = new AuthBackendRoleArgs();
        }

        public Builder(AuthBackendRoleArgs defaults) {
            $ = new AuthBackendRoleArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param backend The unique name of the auth backend to configure.
         * Defaults to `approle`.
         * 
         * @return builder
         * 
         */
        public Builder backend(@Nullable Output<String> backend) {
            $.backend = backend;
            return this;
        }

        /**
         * @param backend The unique name of the auth backend to configure.
         * Defaults to `approle`.
         * 
         * @return builder
         * 
         */
        public Builder backend(String backend) {
            return backend(Output.of(backend));
        }

        /**
         * @param bindSecretId Whether or not to require `secret_id` to be
         * presented when logging in using this AppRole. Defaults to `true`.
         * 
         * @return builder
         * 
         */
        public Builder bindSecretId(@Nullable Output<Boolean> bindSecretId) {
            $.bindSecretId = bindSecretId;
            return this;
        }

        /**
         * @param bindSecretId Whether or not to require `secret_id` to be
         * presented when logging in using this AppRole. Defaults to `true`.
         * 
         * @return builder
         * 
         */
        public Builder bindSecretId(Boolean bindSecretId) {
            return bindSecretId(Output.of(bindSecretId));
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
         * @param roleId The RoleID of this role. If not specified, one will be
         * auto-generated.
         * 
         * @return builder
         * 
         */
        public Builder roleId(@Nullable Output<String> roleId) {
            $.roleId = roleId;
            return this;
        }

        /**
         * @param roleId The RoleID of this role. If not specified, one will be
         * auto-generated.
         * 
         * @return builder
         * 
         */
        public Builder roleId(String roleId) {
            return roleId(Output.of(roleId));
        }

        /**
         * @param roleName The name of the role.
         * 
         * @return builder
         * 
         */
        public Builder roleName(Output<String> roleName) {
            $.roleName = roleName;
            return this;
        }

        /**
         * @param roleName The name of the role.
         * 
         * @return builder
         * 
         */
        public Builder roleName(String roleName) {
            return roleName(Output.of(roleName));
        }

        /**
         * @param secretIdBoundCidrs If set,
         * specifies blocks of IP addresses which can perform the login operation.
         * 
         * @return builder
         * 
         */
        public Builder secretIdBoundCidrs(@Nullable Output<List<String>> secretIdBoundCidrs) {
            $.secretIdBoundCidrs = secretIdBoundCidrs;
            return this;
        }

        /**
         * @param secretIdBoundCidrs If set,
         * specifies blocks of IP addresses which can perform the login operation.
         * 
         * @return builder
         * 
         */
        public Builder secretIdBoundCidrs(List<String> secretIdBoundCidrs) {
            return secretIdBoundCidrs(Output.of(secretIdBoundCidrs));
        }

        /**
         * @param secretIdBoundCidrs If set,
         * specifies blocks of IP addresses which can perform the login operation.
         * 
         * @return builder
         * 
         */
        public Builder secretIdBoundCidrs(String... secretIdBoundCidrs) {
            return secretIdBoundCidrs(List.of(secretIdBoundCidrs));
        }

        /**
         * @param secretIdNumUses The number of times any particular SecretID
         * can be used to fetch a token from this AppRole, after which the SecretID will
         * expire. A value of zero will allow unlimited uses.
         * 
         * @return builder
         * 
         */
        public Builder secretIdNumUses(@Nullable Output<Integer> secretIdNumUses) {
            $.secretIdNumUses = secretIdNumUses;
            return this;
        }

        /**
         * @param secretIdNumUses The number of times any particular SecretID
         * can be used to fetch a token from this AppRole, after which the SecretID will
         * expire. A value of zero will allow unlimited uses.
         * 
         * @return builder
         * 
         */
        public Builder secretIdNumUses(Integer secretIdNumUses) {
            return secretIdNumUses(Output.of(secretIdNumUses));
        }

        /**
         * @param secretIdTtl The number of seconds after which any SecretID
         * expires.
         * 
         * @return builder
         * 
         */
        public Builder secretIdTtl(@Nullable Output<Integer> secretIdTtl) {
            $.secretIdTtl = secretIdTtl;
            return this;
        }

        /**
         * @param secretIdTtl The number of seconds after which any SecretID
         * expires.
         * 
         * @return builder
         * 
         */
        public Builder secretIdTtl(Integer secretIdTtl) {
            return secretIdTtl(Output.of(secretIdTtl));
        }

        /**
         * @param tokenBoundCidrs Specifies the blocks of IP addresses which are allowed to use the generated token
         * 
         * @return builder
         * 
         */
        public Builder tokenBoundCidrs(@Nullable Output<List<String>> tokenBoundCidrs) {
            $.tokenBoundCidrs = tokenBoundCidrs;
            return this;
        }

        /**
         * @param tokenBoundCidrs Specifies the blocks of IP addresses which are allowed to use the generated token
         * 
         * @return builder
         * 
         */
        public Builder tokenBoundCidrs(List<String> tokenBoundCidrs) {
            return tokenBoundCidrs(Output.of(tokenBoundCidrs));
        }

        /**
         * @param tokenBoundCidrs Specifies the blocks of IP addresses which are allowed to use the generated token
         * 
         * @return builder
         * 
         */
        public Builder tokenBoundCidrs(String... tokenBoundCidrs) {
            return tokenBoundCidrs(List.of(tokenBoundCidrs));
        }

        /**
         * @param tokenExplicitMaxTtl Generated Token&#39;s Explicit Maximum TTL in seconds
         * 
         * @return builder
         * 
         */
        public Builder tokenExplicitMaxTtl(@Nullable Output<Integer> tokenExplicitMaxTtl) {
            $.tokenExplicitMaxTtl = tokenExplicitMaxTtl;
            return this;
        }

        /**
         * @param tokenExplicitMaxTtl Generated Token&#39;s Explicit Maximum TTL in seconds
         * 
         * @return builder
         * 
         */
        public Builder tokenExplicitMaxTtl(Integer tokenExplicitMaxTtl) {
            return tokenExplicitMaxTtl(Output.of(tokenExplicitMaxTtl));
        }

        /**
         * @param tokenMaxTtl The maximum lifetime of the generated token
         * 
         * @return builder
         * 
         */
        public Builder tokenMaxTtl(@Nullable Output<Integer> tokenMaxTtl) {
            $.tokenMaxTtl = tokenMaxTtl;
            return this;
        }

        /**
         * @param tokenMaxTtl The maximum lifetime of the generated token
         * 
         * @return builder
         * 
         */
        public Builder tokenMaxTtl(Integer tokenMaxTtl) {
            return tokenMaxTtl(Output.of(tokenMaxTtl));
        }

        /**
         * @param tokenNoDefaultPolicy If true, the &#39;default&#39; policy will not automatically be added to generated tokens
         * 
         * @return builder
         * 
         */
        public Builder tokenNoDefaultPolicy(@Nullable Output<Boolean> tokenNoDefaultPolicy) {
            $.tokenNoDefaultPolicy = tokenNoDefaultPolicy;
            return this;
        }

        /**
         * @param tokenNoDefaultPolicy If true, the &#39;default&#39; policy will not automatically be added to generated tokens
         * 
         * @return builder
         * 
         */
        public Builder tokenNoDefaultPolicy(Boolean tokenNoDefaultPolicy) {
            return tokenNoDefaultPolicy(Output.of(tokenNoDefaultPolicy));
        }

        /**
         * @param tokenNumUses The maximum number of times a token may be used, a value of zero means unlimited
         * 
         * @return builder
         * 
         */
        public Builder tokenNumUses(@Nullable Output<Integer> tokenNumUses) {
            $.tokenNumUses = tokenNumUses;
            return this;
        }

        /**
         * @param tokenNumUses The maximum number of times a token may be used, a value of zero means unlimited
         * 
         * @return builder
         * 
         */
        public Builder tokenNumUses(Integer tokenNumUses) {
            return tokenNumUses(Output.of(tokenNumUses));
        }

        /**
         * @param tokenPeriod Generated Token&#39;s Period
         * 
         * @return builder
         * 
         */
        public Builder tokenPeriod(@Nullable Output<Integer> tokenPeriod) {
            $.tokenPeriod = tokenPeriod;
            return this;
        }

        /**
         * @param tokenPeriod Generated Token&#39;s Period
         * 
         * @return builder
         * 
         */
        public Builder tokenPeriod(Integer tokenPeriod) {
            return tokenPeriod(Output.of(tokenPeriod));
        }

        /**
         * @param tokenPolicies Generated Token&#39;s Policies
         * 
         * @return builder
         * 
         */
        public Builder tokenPolicies(@Nullable Output<List<String>> tokenPolicies) {
            $.tokenPolicies = tokenPolicies;
            return this;
        }

        /**
         * @param tokenPolicies Generated Token&#39;s Policies
         * 
         * @return builder
         * 
         */
        public Builder tokenPolicies(List<String> tokenPolicies) {
            return tokenPolicies(Output.of(tokenPolicies));
        }

        /**
         * @param tokenPolicies Generated Token&#39;s Policies
         * 
         * @return builder
         * 
         */
        public Builder tokenPolicies(String... tokenPolicies) {
            return tokenPolicies(List.of(tokenPolicies));
        }

        /**
         * @param tokenTtl The initial ttl of the token to generate in seconds
         * 
         * @return builder
         * 
         */
        public Builder tokenTtl(@Nullable Output<Integer> tokenTtl) {
            $.tokenTtl = tokenTtl;
            return this;
        }

        /**
         * @param tokenTtl The initial ttl of the token to generate in seconds
         * 
         * @return builder
         * 
         */
        public Builder tokenTtl(Integer tokenTtl) {
            return tokenTtl(Output.of(tokenTtl));
        }

        /**
         * @param tokenType The type of token to generate, service or batch
         * 
         * @return builder
         * 
         */
        public Builder tokenType(@Nullable Output<String> tokenType) {
            $.tokenType = tokenType;
            return this;
        }

        /**
         * @param tokenType The type of token to generate, service or batch
         * 
         * @return builder
         * 
         */
        public Builder tokenType(String tokenType) {
            return tokenType(Output.of(tokenType));
        }

        public AuthBackendRoleArgs build() {
            if ($.roleName == null) {
                throw new MissingRequiredPropertyException("AuthBackendRoleArgs", "roleName");
            }
            return $;
        }
    }

}