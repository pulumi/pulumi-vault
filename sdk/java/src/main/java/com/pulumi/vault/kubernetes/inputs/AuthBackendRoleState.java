// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vault.kubernetes.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class AuthBackendRoleState extends com.pulumi.resources.ResourceArgs {

    public static final AuthBackendRoleState Empty = new AuthBackendRoleState();

    /**
     * Configures how identity aliases are generated.
     * Valid choices are: `serviceaccount_uid`, `serviceaccount_name`. (vault-1.9+)
     * 
     */
    @Import(name="aliasNameSource")
    private @Nullable Output<String> aliasNameSource;

    /**
     * @return Configures how identity aliases are generated.
     * Valid choices are: `serviceaccount_uid`, `serviceaccount_name`. (vault-1.9+)
     * 
     */
    public Optional<Output<String>> aliasNameSource() {
        return Optional.ofNullable(this.aliasNameSource);
    }

    /**
     * Audience claim to verify in the JWT.
     * 
     * &gt; Please see [alias_name_source](https://www.vaultproject.io/api-docs/auth/kubernetes#alias_name_source)
     * before setting this to something other its default value. There are **important** security
     * implications to be aware of.
     * 
     */
    @Import(name="audience")
    private @Nullable Output<String> audience;

    /**
     * @return Audience claim to verify in the JWT.
     * 
     * &gt; Please see [alias_name_source](https://www.vaultproject.io/api-docs/auth/kubernetes#alias_name_source)
     * before setting this to something other its default value. There are **important** security
     * implications to be aware of.
     * 
     */
    public Optional<Output<String>> audience() {
        return Optional.ofNullable(this.audience);
    }

    /**
     * Unique name of the kubernetes backend to configure.
     * 
     */
    @Import(name="backend")
    private @Nullable Output<String> backend;

    /**
     * @return Unique name of the kubernetes backend to configure.
     * 
     */
    public Optional<Output<String>> backend() {
        return Optional.ofNullable(this.backend);
    }

    /**
     * List of service account names able to access this role. If set to `[&#34;*&#34;]` all names are allowed, both this and bound_service_account_namespaces can not be &#34;*&#34;.
     * 
     */
    @Import(name="boundServiceAccountNames")
    private @Nullable Output<List<String>> boundServiceAccountNames;

    /**
     * @return List of service account names able to access this role. If set to `[&#34;*&#34;]` all names are allowed, both this and bound_service_account_namespaces can not be &#34;*&#34;.
     * 
     */
    public Optional<Output<List<String>>> boundServiceAccountNames() {
        return Optional.ofNullable(this.boundServiceAccountNames);
    }

    /**
     * List of namespaces allowed to access this role. If set to `[&#34;*&#34;]` all namespaces are allowed, both this and bound_service_account_names can not be set to &#34;*&#34;.
     * 
     */
    @Import(name="boundServiceAccountNamespaces")
    private @Nullable Output<List<String>> boundServiceAccountNamespaces;

    /**
     * @return List of namespaces allowed to access this role. If set to `[&#34;*&#34;]` all namespaces are allowed, both this and bound_service_account_names can not be set to &#34;*&#34;.
     * 
     */
    public Optional<Output<List<String>>> boundServiceAccountNamespaces() {
        return Optional.ofNullable(this.boundServiceAccountNamespaces);
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

    private AuthBackendRoleState() {}

    private AuthBackendRoleState(AuthBackendRoleState $) {
        this.aliasNameSource = $.aliasNameSource;
        this.audience = $.audience;
        this.backend = $.backend;
        this.boundServiceAccountNames = $.boundServiceAccountNames;
        this.boundServiceAccountNamespaces = $.boundServiceAccountNamespaces;
        this.namespace = $.namespace;
        this.roleName = $.roleName;
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
    public static Builder builder(AuthBackendRoleState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private AuthBackendRoleState $;

        public Builder() {
            $ = new AuthBackendRoleState();
        }

        public Builder(AuthBackendRoleState defaults) {
            $ = new AuthBackendRoleState(Objects.requireNonNull(defaults));
        }

        /**
         * @param aliasNameSource Configures how identity aliases are generated.
         * Valid choices are: `serviceaccount_uid`, `serviceaccount_name`. (vault-1.9+)
         * 
         * @return builder
         * 
         */
        public Builder aliasNameSource(@Nullable Output<String> aliasNameSource) {
            $.aliasNameSource = aliasNameSource;
            return this;
        }

        /**
         * @param aliasNameSource Configures how identity aliases are generated.
         * Valid choices are: `serviceaccount_uid`, `serviceaccount_name`. (vault-1.9+)
         * 
         * @return builder
         * 
         */
        public Builder aliasNameSource(String aliasNameSource) {
            return aliasNameSource(Output.of(aliasNameSource));
        }

        /**
         * @param audience Audience claim to verify in the JWT.
         * 
         * &gt; Please see [alias_name_source](https://www.vaultproject.io/api-docs/auth/kubernetes#alias_name_source)
         * before setting this to something other its default value. There are **important** security
         * implications to be aware of.
         * 
         * @return builder
         * 
         */
        public Builder audience(@Nullable Output<String> audience) {
            $.audience = audience;
            return this;
        }

        /**
         * @param audience Audience claim to verify in the JWT.
         * 
         * &gt; Please see [alias_name_source](https://www.vaultproject.io/api-docs/auth/kubernetes#alias_name_source)
         * before setting this to something other its default value. There are **important** security
         * implications to be aware of.
         * 
         * @return builder
         * 
         */
        public Builder audience(String audience) {
            return audience(Output.of(audience));
        }

        /**
         * @param backend Unique name of the kubernetes backend to configure.
         * 
         * @return builder
         * 
         */
        public Builder backend(@Nullable Output<String> backend) {
            $.backend = backend;
            return this;
        }

        /**
         * @param backend Unique name of the kubernetes backend to configure.
         * 
         * @return builder
         * 
         */
        public Builder backend(String backend) {
            return backend(Output.of(backend));
        }

        /**
         * @param boundServiceAccountNames List of service account names able to access this role. If set to `[&#34;*&#34;]` all names are allowed, both this and bound_service_account_namespaces can not be &#34;*&#34;.
         * 
         * @return builder
         * 
         */
        public Builder boundServiceAccountNames(@Nullable Output<List<String>> boundServiceAccountNames) {
            $.boundServiceAccountNames = boundServiceAccountNames;
            return this;
        }

        /**
         * @param boundServiceAccountNames List of service account names able to access this role. If set to `[&#34;*&#34;]` all names are allowed, both this and bound_service_account_namespaces can not be &#34;*&#34;.
         * 
         * @return builder
         * 
         */
        public Builder boundServiceAccountNames(List<String> boundServiceAccountNames) {
            return boundServiceAccountNames(Output.of(boundServiceAccountNames));
        }

        /**
         * @param boundServiceAccountNames List of service account names able to access this role. If set to `[&#34;*&#34;]` all names are allowed, both this and bound_service_account_namespaces can not be &#34;*&#34;.
         * 
         * @return builder
         * 
         */
        public Builder boundServiceAccountNames(String... boundServiceAccountNames) {
            return boundServiceAccountNames(List.of(boundServiceAccountNames));
        }

        /**
         * @param boundServiceAccountNamespaces List of namespaces allowed to access this role. If set to `[&#34;*&#34;]` all namespaces are allowed, both this and bound_service_account_names can not be set to &#34;*&#34;.
         * 
         * @return builder
         * 
         */
        public Builder boundServiceAccountNamespaces(@Nullable Output<List<String>> boundServiceAccountNamespaces) {
            $.boundServiceAccountNamespaces = boundServiceAccountNamespaces;
            return this;
        }

        /**
         * @param boundServiceAccountNamespaces List of namespaces allowed to access this role. If set to `[&#34;*&#34;]` all namespaces are allowed, both this and bound_service_account_names can not be set to &#34;*&#34;.
         * 
         * @return builder
         * 
         */
        public Builder boundServiceAccountNamespaces(List<String> boundServiceAccountNamespaces) {
            return boundServiceAccountNamespaces(Output.of(boundServiceAccountNamespaces));
        }

        /**
         * @param boundServiceAccountNamespaces List of namespaces allowed to access this role. If set to `[&#34;*&#34;]` all namespaces are allowed, both this and bound_service_account_names can not be set to &#34;*&#34;.
         * 
         * @return builder
         * 
         */
        public Builder boundServiceAccountNamespaces(String... boundServiceAccountNamespaces) {
            return boundServiceAccountNamespaces(List.of(boundServiceAccountNamespaces));
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

        public AuthBackendRoleState build() {
            return $;
        }
    }

}