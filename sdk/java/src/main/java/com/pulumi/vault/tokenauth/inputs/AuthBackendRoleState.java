// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vault.tokenauth.inputs;

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
     * List of allowed entity aliases.
     * 
     */
    @Import(name="allowedEntityAliases")
    private @Nullable Output<List<String>> allowedEntityAliases;

    /**
     * @return List of allowed entity aliases.
     * 
     */
    public Optional<Output<List<String>>> allowedEntityAliases() {
        return Optional.ofNullable(this.allowedEntityAliases);
    }

    /**
     * List of allowed policies for given role.
     * 
     */
    @Import(name="allowedPolicies")
    private @Nullable Output<List<String>> allowedPolicies;

    /**
     * @return List of allowed policies for given role.
     * 
     */
    public Optional<Output<List<String>>> allowedPolicies() {
        return Optional.ofNullable(this.allowedPolicies);
    }

    /**
     * Set of allowed policies with glob match for given role.
     * 
     */
    @Import(name="allowedPoliciesGlobs")
    private @Nullable Output<List<String>> allowedPoliciesGlobs;

    /**
     * @return Set of allowed policies with glob match for given role.
     * 
     */
    public Optional<Output<List<String>>> allowedPoliciesGlobs() {
        return Optional.ofNullable(this.allowedPoliciesGlobs);
    }

    /**
     * List of disallowed policies for given role.
     * 
     */
    @Import(name="disallowedPolicies")
    private @Nullable Output<List<String>> disallowedPolicies;

    /**
     * @return List of disallowed policies for given role.
     * 
     */
    public Optional<Output<List<String>>> disallowedPolicies() {
        return Optional.ofNullable(this.disallowedPolicies);
    }

    /**
     * Set of disallowed policies with glob match for given role.
     * 
     */
    @Import(name="disallowedPoliciesGlobs")
    private @Nullable Output<List<String>> disallowedPoliciesGlobs;

    /**
     * @return Set of disallowed policies with glob match for given role.
     * 
     */
    public Optional<Output<List<String>>> disallowedPoliciesGlobs() {
        return Optional.ofNullable(this.disallowedPoliciesGlobs);
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
     * If true, tokens created against this policy will be orphan tokens.
     * 
     */
    @Import(name="orphan")
    private @Nullable Output<Boolean> orphan;

    /**
     * @return If true, tokens created against this policy will be orphan tokens.
     * 
     */
    public Optional<Output<Boolean>> orphan() {
        return Optional.ofNullable(this.orphan);
    }

    /**
     * Tokens created against this role will have the given suffix as part of their path in addition to the role name.
     * 
     * &gt; Due to a bug the resource. This *will* cause all existing tokens issued by this role to be revoked.
     * 
     */
    @Import(name="pathSuffix")
    private @Nullable Output<String> pathSuffix;

    /**
     * @return Tokens created against this role will have the given suffix as part of their path in addition to the role name.
     * 
     * &gt; Due to a bug the resource. This *will* cause all existing tokens issued by this role to be revoked.
     * 
     */
    public Optional<Output<String>> pathSuffix() {
        return Optional.ofNullable(this.pathSuffix);
    }

    /**
     * Whether to disable the ability of the token to be renewed past its initial TTL.
     * 
     */
    @Import(name="renewable")
    private @Nullable Output<Boolean> renewable;

    /**
     * @return Whether to disable the ability of the token to be renewed past its initial TTL.
     * 
     */
    public Optional<Output<Boolean>> renewable() {
        return Optional.ofNullable(this.renewable);
    }

    /**
     * The name of the role.
     * 
     */
    @Import(name="roleName")
    private @Nullable Output<String> roleName;

    /**
     * @return The name of the role.
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
        this.allowedEntityAliases = $.allowedEntityAliases;
        this.allowedPolicies = $.allowedPolicies;
        this.allowedPoliciesGlobs = $.allowedPoliciesGlobs;
        this.disallowedPolicies = $.disallowedPolicies;
        this.disallowedPoliciesGlobs = $.disallowedPoliciesGlobs;
        this.namespace = $.namespace;
        this.orphan = $.orphan;
        this.pathSuffix = $.pathSuffix;
        this.renewable = $.renewable;
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
         * @param allowedEntityAliases List of allowed entity aliases.
         * 
         * @return builder
         * 
         */
        public Builder allowedEntityAliases(@Nullable Output<List<String>> allowedEntityAliases) {
            $.allowedEntityAliases = allowedEntityAliases;
            return this;
        }

        /**
         * @param allowedEntityAliases List of allowed entity aliases.
         * 
         * @return builder
         * 
         */
        public Builder allowedEntityAliases(List<String> allowedEntityAliases) {
            return allowedEntityAliases(Output.of(allowedEntityAliases));
        }

        /**
         * @param allowedEntityAliases List of allowed entity aliases.
         * 
         * @return builder
         * 
         */
        public Builder allowedEntityAliases(String... allowedEntityAliases) {
            return allowedEntityAliases(List.of(allowedEntityAliases));
        }

        /**
         * @param allowedPolicies List of allowed policies for given role.
         * 
         * @return builder
         * 
         */
        public Builder allowedPolicies(@Nullable Output<List<String>> allowedPolicies) {
            $.allowedPolicies = allowedPolicies;
            return this;
        }

        /**
         * @param allowedPolicies List of allowed policies for given role.
         * 
         * @return builder
         * 
         */
        public Builder allowedPolicies(List<String> allowedPolicies) {
            return allowedPolicies(Output.of(allowedPolicies));
        }

        /**
         * @param allowedPolicies List of allowed policies for given role.
         * 
         * @return builder
         * 
         */
        public Builder allowedPolicies(String... allowedPolicies) {
            return allowedPolicies(List.of(allowedPolicies));
        }

        /**
         * @param allowedPoliciesGlobs Set of allowed policies with glob match for given role.
         * 
         * @return builder
         * 
         */
        public Builder allowedPoliciesGlobs(@Nullable Output<List<String>> allowedPoliciesGlobs) {
            $.allowedPoliciesGlobs = allowedPoliciesGlobs;
            return this;
        }

        /**
         * @param allowedPoliciesGlobs Set of allowed policies with glob match for given role.
         * 
         * @return builder
         * 
         */
        public Builder allowedPoliciesGlobs(List<String> allowedPoliciesGlobs) {
            return allowedPoliciesGlobs(Output.of(allowedPoliciesGlobs));
        }

        /**
         * @param allowedPoliciesGlobs Set of allowed policies with glob match for given role.
         * 
         * @return builder
         * 
         */
        public Builder allowedPoliciesGlobs(String... allowedPoliciesGlobs) {
            return allowedPoliciesGlobs(List.of(allowedPoliciesGlobs));
        }

        /**
         * @param disallowedPolicies List of disallowed policies for given role.
         * 
         * @return builder
         * 
         */
        public Builder disallowedPolicies(@Nullable Output<List<String>> disallowedPolicies) {
            $.disallowedPolicies = disallowedPolicies;
            return this;
        }

        /**
         * @param disallowedPolicies List of disallowed policies for given role.
         * 
         * @return builder
         * 
         */
        public Builder disallowedPolicies(List<String> disallowedPolicies) {
            return disallowedPolicies(Output.of(disallowedPolicies));
        }

        /**
         * @param disallowedPolicies List of disallowed policies for given role.
         * 
         * @return builder
         * 
         */
        public Builder disallowedPolicies(String... disallowedPolicies) {
            return disallowedPolicies(List.of(disallowedPolicies));
        }

        /**
         * @param disallowedPoliciesGlobs Set of disallowed policies with glob match for given role.
         * 
         * @return builder
         * 
         */
        public Builder disallowedPoliciesGlobs(@Nullable Output<List<String>> disallowedPoliciesGlobs) {
            $.disallowedPoliciesGlobs = disallowedPoliciesGlobs;
            return this;
        }

        /**
         * @param disallowedPoliciesGlobs Set of disallowed policies with glob match for given role.
         * 
         * @return builder
         * 
         */
        public Builder disallowedPoliciesGlobs(List<String> disallowedPoliciesGlobs) {
            return disallowedPoliciesGlobs(Output.of(disallowedPoliciesGlobs));
        }

        /**
         * @param disallowedPoliciesGlobs Set of disallowed policies with glob match for given role.
         * 
         * @return builder
         * 
         */
        public Builder disallowedPoliciesGlobs(String... disallowedPoliciesGlobs) {
            return disallowedPoliciesGlobs(List.of(disallowedPoliciesGlobs));
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
         * @param orphan If true, tokens created against this policy will be orphan tokens.
         * 
         * @return builder
         * 
         */
        public Builder orphan(@Nullable Output<Boolean> orphan) {
            $.orphan = orphan;
            return this;
        }

        /**
         * @param orphan If true, tokens created against this policy will be orphan tokens.
         * 
         * @return builder
         * 
         */
        public Builder orphan(Boolean orphan) {
            return orphan(Output.of(orphan));
        }

        /**
         * @param pathSuffix Tokens created against this role will have the given suffix as part of their path in addition to the role name.
         * 
         * &gt; Due to a bug the resource. This *will* cause all existing tokens issued by this role to be revoked.
         * 
         * @return builder
         * 
         */
        public Builder pathSuffix(@Nullable Output<String> pathSuffix) {
            $.pathSuffix = pathSuffix;
            return this;
        }

        /**
         * @param pathSuffix Tokens created against this role will have the given suffix as part of their path in addition to the role name.
         * 
         * &gt; Due to a bug the resource. This *will* cause all existing tokens issued by this role to be revoked.
         * 
         * @return builder
         * 
         */
        public Builder pathSuffix(String pathSuffix) {
            return pathSuffix(Output.of(pathSuffix));
        }

        /**
         * @param renewable Whether to disable the ability of the token to be renewed past its initial TTL.
         * 
         * @return builder
         * 
         */
        public Builder renewable(@Nullable Output<Boolean> renewable) {
            $.renewable = renewable;
            return this;
        }

        /**
         * @param renewable Whether to disable the ability of the token to be renewed past its initial TTL.
         * 
         * @return builder
         * 
         */
        public Builder renewable(Boolean renewable) {
            return renewable(Output.of(renewable));
        }

        /**
         * @param roleName The name of the role.
         * 
         * @return builder
         * 
         */
        public Builder roleName(@Nullable Output<String> roleName) {
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
