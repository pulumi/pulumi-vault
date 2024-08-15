// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vault.saml;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class AuthBackendRoleArgs extends com.pulumi.resources.ResourceArgs {

    public static final AuthBackendRoleArgs Empty = new AuthBackendRoleArgs();

    /**
     * Mapping of attribute names to values that are expected to
     * exist in the SAML assertion.
     * 
     */
    @Import(name="boundAttributes")
    private @Nullable Output<Map<String,String>> boundAttributes;

    /**
     * @return Mapping of attribute names to values that are expected to
     * exist in the SAML assertion.
     * 
     */
    public Optional<Output<Map<String,String>>> boundAttributes() {
        return Optional.ofNullable(this.boundAttributes);
    }

    /**
     * The type of matching assertion to perform on
     * `bound_attributes_type`.
     * 
     */
    @Import(name="boundAttributesType")
    private @Nullable Output<String> boundAttributesType;

    /**
     * @return The type of matching assertion to perform on
     * `bound_attributes_type`.
     * 
     */
    public Optional<Output<String>> boundAttributesType() {
        return Optional.ofNullable(this.boundAttributesType);
    }

    /**
     * List of subjects being asserted for SAML authentication.
     * 
     */
    @Import(name="boundSubjects")
    private @Nullable Output<List<String>> boundSubjects;

    /**
     * @return List of subjects being asserted for SAML authentication.
     * 
     */
    public Optional<Output<List<String>>> boundSubjects() {
        return Optional.ofNullable(this.boundSubjects);
    }

    /**
     * The type of matching assertion to perform on `bound_subjects`.
     * 
     */
    @Import(name="boundSubjectsType")
    private @Nullable Output<String> boundSubjectsType;

    /**
     * @return The type of matching assertion to perform on `bound_subjects`.
     * 
     */
    public Optional<Output<String>> boundSubjectsType() {
        return Optional.ofNullable(this.boundSubjectsType);
    }

    /**
     * The attribute to use to identify the set of groups to which the
     * user belongs.
     * 
     */
    @Import(name="groupsAttribute")
    private @Nullable Output<String> groupsAttribute;

    /**
     * @return The attribute to use to identify the set of groups to which the
     * user belongs.
     * 
     */
    public Optional<Output<String>> groupsAttribute() {
        return Optional.ofNullable(this.groupsAttribute);
    }

    /**
     * Unique name of the role.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return Unique name of the role.
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
     * Path where the auth backend is mounted.
     * 
     */
    @Import(name="path", required=true)
    private Output<String> path;

    /**
     * @return Path where the auth backend is mounted.
     * 
     */
    public Output<String> path() {
        return this.path;
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
        this.boundAttributes = $.boundAttributes;
        this.boundAttributesType = $.boundAttributesType;
        this.boundSubjects = $.boundSubjects;
        this.boundSubjectsType = $.boundSubjectsType;
        this.groupsAttribute = $.groupsAttribute;
        this.name = $.name;
        this.namespace = $.namespace;
        this.path = $.path;
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
         * @param boundAttributes Mapping of attribute names to values that are expected to
         * exist in the SAML assertion.
         * 
         * @return builder
         * 
         */
        public Builder boundAttributes(@Nullable Output<Map<String,String>> boundAttributes) {
            $.boundAttributes = boundAttributes;
            return this;
        }

        /**
         * @param boundAttributes Mapping of attribute names to values that are expected to
         * exist in the SAML assertion.
         * 
         * @return builder
         * 
         */
        public Builder boundAttributes(Map<String,String> boundAttributes) {
            return boundAttributes(Output.of(boundAttributes));
        }

        /**
         * @param boundAttributesType The type of matching assertion to perform on
         * `bound_attributes_type`.
         * 
         * @return builder
         * 
         */
        public Builder boundAttributesType(@Nullable Output<String> boundAttributesType) {
            $.boundAttributesType = boundAttributesType;
            return this;
        }

        /**
         * @param boundAttributesType The type of matching assertion to perform on
         * `bound_attributes_type`.
         * 
         * @return builder
         * 
         */
        public Builder boundAttributesType(String boundAttributesType) {
            return boundAttributesType(Output.of(boundAttributesType));
        }

        /**
         * @param boundSubjects List of subjects being asserted for SAML authentication.
         * 
         * @return builder
         * 
         */
        public Builder boundSubjects(@Nullable Output<List<String>> boundSubjects) {
            $.boundSubjects = boundSubjects;
            return this;
        }

        /**
         * @param boundSubjects List of subjects being asserted for SAML authentication.
         * 
         * @return builder
         * 
         */
        public Builder boundSubjects(List<String> boundSubjects) {
            return boundSubjects(Output.of(boundSubjects));
        }

        /**
         * @param boundSubjects List of subjects being asserted for SAML authentication.
         * 
         * @return builder
         * 
         */
        public Builder boundSubjects(String... boundSubjects) {
            return boundSubjects(List.of(boundSubjects));
        }

        /**
         * @param boundSubjectsType The type of matching assertion to perform on `bound_subjects`.
         * 
         * @return builder
         * 
         */
        public Builder boundSubjectsType(@Nullable Output<String> boundSubjectsType) {
            $.boundSubjectsType = boundSubjectsType;
            return this;
        }

        /**
         * @param boundSubjectsType The type of matching assertion to perform on `bound_subjects`.
         * 
         * @return builder
         * 
         */
        public Builder boundSubjectsType(String boundSubjectsType) {
            return boundSubjectsType(Output.of(boundSubjectsType));
        }

        /**
         * @param groupsAttribute The attribute to use to identify the set of groups to which the
         * user belongs.
         * 
         * @return builder
         * 
         */
        public Builder groupsAttribute(@Nullable Output<String> groupsAttribute) {
            $.groupsAttribute = groupsAttribute;
            return this;
        }

        /**
         * @param groupsAttribute The attribute to use to identify the set of groups to which the
         * user belongs.
         * 
         * @return builder
         * 
         */
        public Builder groupsAttribute(String groupsAttribute) {
            return groupsAttribute(Output.of(groupsAttribute));
        }

        /**
         * @param name Unique name of the role.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name Unique name of the role.
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
         * @param path Path where the auth backend is mounted.
         * 
         * @return builder
         * 
         */
        public Builder path(Output<String> path) {
            $.path = path;
            return this;
        }

        /**
         * @param path Path where the auth backend is mounted.
         * 
         * @return builder
         * 
         */
        public Builder path(String path) {
            return path(Output.of(path));
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
            if ($.path == null) {
                throw new MissingRequiredPropertyException("AuthBackendRoleArgs", "path");
            }
            return $;
        }
    }

}
