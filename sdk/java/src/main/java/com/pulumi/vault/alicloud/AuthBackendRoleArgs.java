// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vault.alicloud;

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
     * The role&#39;s arn.
     * 
     */
    @Import(name="arn", required=true)
    private Output<String> arn;

    /**
     * @return The role&#39;s arn.
     * 
     */
    public Output<String> arn() {
        return this.arn;
    }

    /**
     * Path to the mounted AliCloud auth backend.
     * Defaults to `alicloud`
     * 
     * For more details on the usage of each argument consult the [Vault AliCloud API documentation](https://www.vaultproject.io/api-docs/auth/alicloud).
     * 
     */
    @Import(name="backend")
    private @Nullable Output<String> backend;

    /**
     * @return Path to the mounted AliCloud auth backend.
     * Defaults to `alicloud`
     * 
     * For more details on the usage of each argument consult the [Vault AliCloud API documentation](https://www.vaultproject.io/api-docs/auth/alicloud).
     * 
     */
    public Optional<Output<String>> backend() {
        return Optional.ofNullable(this.backend);
    }

    /**
     * The namespace to provision the resource in.
     * The value should not contain leading or trailing forward slashes.
     * The `namespace` is always relative to the provider&#39;s configured namespace.
     * *Available only for Vault Enterprise*.
     * 
     */
    @Import(name="namespace")
    private @Nullable Output<String> namespace;

    /**
     * @return The namespace to provision the resource in.
     * The value should not contain leading or trailing forward slashes.
     * The `namespace` is always relative to the provider&#39;s configured namespace.
     * *Available only for Vault Enterprise*.
     * 
     */
    public Optional<Output<String>> namespace() {
        return Optional.ofNullable(this.namespace);
    }

    /**
     * Name of the role. Must correspond with the name of
     * the role reflected in the arn.
     * 
     */
    @Import(name="role", required=true)
    private Output<String> role;

    /**
     * @return Name of the role. Must correspond with the name of
     * the role reflected in the arn.
     * 
     */
    public Output<String> role() {
        return this.role;
    }

    /**
     * List of CIDR blocks; if set, specifies blocks of IP
     * addresses which can authenticate successfully, and ties the resulting token to these blocks
     * as well.
     * 
     */
    @Import(name="tokenBoundCidrs")
    private @Nullable Output<List<String>> tokenBoundCidrs;

    /**
     * @return List of CIDR blocks; if set, specifies blocks of IP
     * addresses which can authenticate successfully, and ties the resulting token to these blocks
     * as well.
     * 
     */
    public Optional<Output<List<String>>> tokenBoundCidrs() {
        return Optional.ofNullable(this.tokenBoundCidrs);
    }

    /**
     * If set, will encode an
     * [explicit max TTL](https://www.vaultproject.io/docs/concepts/tokens.html#token-time-to-live-periodic-tokens-and-explicit-max-ttls)
     * onto the token in number of seconds. This is a hard cap even if `token_ttl` and
     * `token_max_ttl` would otherwise allow a renewal.
     * 
     */
    @Import(name="tokenExplicitMaxTtl")
    private @Nullable Output<Integer> tokenExplicitMaxTtl;

    /**
     * @return If set, will encode an
     * [explicit max TTL](https://www.vaultproject.io/docs/concepts/tokens.html#token-time-to-live-periodic-tokens-and-explicit-max-ttls)
     * onto the token in number of seconds. This is a hard cap even if `token_ttl` and
     * `token_max_ttl` would otherwise allow a renewal.
     * 
     */
    public Optional<Output<Integer>> tokenExplicitMaxTtl() {
        return Optional.ofNullable(this.tokenExplicitMaxTtl);
    }

    /**
     * The maximum lifetime for generated tokens in number of seconds.
     * Its current value will be referenced at renewal time.
     * 
     */
    @Import(name="tokenMaxTtl")
    private @Nullable Output<Integer> tokenMaxTtl;

    /**
     * @return The maximum lifetime for generated tokens in number of seconds.
     * Its current value will be referenced at renewal time.
     * 
     */
    public Optional<Output<Integer>> tokenMaxTtl() {
        return Optional.ofNullable(this.tokenMaxTtl);
    }

    /**
     * If set, the default policy will not be set on
     * generated tokens; otherwise it will be added to the policies set in token_policies.
     * 
     */
    @Import(name="tokenNoDefaultPolicy")
    private @Nullable Output<Boolean> tokenNoDefaultPolicy;

    /**
     * @return If set, the default policy will not be set on
     * generated tokens; otherwise it will be added to the policies set in token_policies.
     * 
     */
    public Optional<Output<Boolean>> tokenNoDefaultPolicy() {
        return Optional.ofNullable(this.tokenNoDefaultPolicy);
    }

    /**
     * The [maximum number](https://www.vaultproject.io/api-docs/auth/alicloud#token_num_uses)
     * of times a generated token may be used (within its lifetime); 0 means unlimited.
     * 
     */
    @Import(name="tokenNumUses")
    private @Nullable Output<Integer> tokenNumUses;

    /**
     * @return The [maximum number](https://www.vaultproject.io/api-docs/auth/alicloud#token_num_uses)
     * of times a generated token may be used (within its lifetime); 0 means unlimited.
     * 
     */
    public Optional<Output<Integer>> tokenNumUses() {
        return Optional.ofNullable(this.tokenNumUses);
    }

    /**
     * If set, indicates that the
     * token generated using this role should never expire. The token should be renewed within the
     * duration specified by this value. At each renewal, the token&#39;s TTL will be set to the
     * value of this field. Specified in seconds.
     * 
     */
    @Import(name="tokenPeriod")
    private @Nullable Output<Integer> tokenPeriod;

    /**
     * @return If set, indicates that the
     * token generated using this role should never expire. The token should be renewed within the
     * duration specified by this value. At each renewal, the token&#39;s TTL will be set to the
     * value of this field. Specified in seconds.
     * 
     */
    public Optional<Output<Integer>> tokenPeriod() {
        return Optional.ofNullable(this.tokenPeriod);
    }

    /**
     * List of policies to encode onto generated tokens. Depending
     * on the auth method, this list may be supplemented by user/group/other values.
     * 
     */
    @Import(name="tokenPolicies")
    private @Nullable Output<List<String>> tokenPolicies;

    /**
     * @return List of policies to encode onto generated tokens. Depending
     * on the auth method, this list may be supplemented by user/group/other values.
     * 
     */
    public Optional<Output<List<String>>> tokenPolicies() {
        return Optional.ofNullable(this.tokenPolicies);
    }

    /**
     * The incremental lifetime for generated tokens in number of seconds.
     * Its current value will be referenced at renewal time.
     * 
     */
    @Import(name="tokenTtl")
    private @Nullable Output<Integer> tokenTtl;

    /**
     * @return The incremental lifetime for generated tokens in number of seconds.
     * Its current value will be referenced at renewal time.
     * 
     */
    public Optional<Output<Integer>> tokenTtl() {
        return Optional.ofNullable(this.tokenTtl);
    }

    /**
     * The type of token that should be generated. Can be `service`,
     * `batch`, or `default` to use the mount&#39;s tuned default (which unless changed will be
     * `service` tokens). For token store roles, there are two additional possibilities:
     * `default-service` and `default-batch` which specify the type to return unless the client
     * requests a different type at generation time.
     * 
     */
    @Import(name="tokenType")
    private @Nullable Output<String> tokenType;

    /**
     * @return The type of token that should be generated. Can be `service`,
     * `batch`, or `default` to use the mount&#39;s tuned default (which unless changed will be
     * `service` tokens). For token store roles, there are two additional possibilities:
     * `default-service` and `default-batch` which specify the type to return unless the client
     * requests a different type at generation time.
     * 
     */
    public Optional<Output<String>> tokenType() {
        return Optional.ofNullable(this.tokenType);
    }

    private AuthBackendRoleArgs() {}

    private AuthBackendRoleArgs(AuthBackendRoleArgs $) {
        this.arn = $.arn;
        this.backend = $.backend;
        this.namespace = $.namespace;
        this.role = $.role;
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
         * @param arn The role&#39;s arn.
         * 
         * @return builder
         * 
         */
        public Builder arn(Output<String> arn) {
            $.arn = arn;
            return this;
        }

        /**
         * @param arn The role&#39;s arn.
         * 
         * @return builder
         * 
         */
        public Builder arn(String arn) {
            return arn(Output.of(arn));
        }

        /**
         * @param backend Path to the mounted AliCloud auth backend.
         * Defaults to `alicloud`
         * 
         * For more details on the usage of each argument consult the [Vault AliCloud API documentation](https://www.vaultproject.io/api-docs/auth/alicloud).
         * 
         * @return builder
         * 
         */
        public Builder backend(@Nullable Output<String> backend) {
            $.backend = backend;
            return this;
        }

        /**
         * @param backend Path to the mounted AliCloud auth backend.
         * Defaults to `alicloud`
         * 
         * For more details on the usage of each argument consult the [Vault AliCloud API documentation](https://www.vaultproject.io/api-docs/auth/alicloud).
         * 
         * @return builder
         * 
         */
        public Builder backend(String backend) {
            return backend(Output.of(backend));
        }

        /**
         * @param namespace The namespace to provision the resource in.
         * The value should not contain leading or trailing forward slashes.
         * The `namespace` is always relative to the provider&#39;s configured namespace.
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
         * The `namespace` is always relative to the provider&#39;s configured namespace.
         * *Available only for Vault Enterprise*.
         * 
         * @return builder
         * 
         */
        public Builder namespace(String namespace) {
            return namespace(Output.of(namespace));
        }

        /**
         * @param role Name of the role. Must correspond with the name of
         * the role reflected in the arn.
         * 
         * @return builder
         * 
         */
        public Builder role(Output<String> role) {
            $.role = role;
            return this;
        }

        /**
         * @param role Name of the role. Must correspond with the name of
         * the role reflected in the arn.
         * 
         * @return builder
         * 
         */
        public Builder role(String role) {
            return role(Output.of(role));
        }

        /**
         * @param tokenBoundCidrs List of CIDR blocks; if set, specifies blocks of IP
         * addresses which can authenticate successfully, and ties the resulting token to these blocks
         * as well.
         * 
         * @return builder
         * 
         */
        public Builder tokenBoundCidrs(@Nullable Output<List<String>> tokenBoundCidrs) {
            $.tokenBoundCidrs = tokenBoundCidrs;
            return this;
        }

        /**
         * @param tokenBoundCidrs List of CIDR blocks; if set, specifies blocks of IP
         * addresses which can authenticate successfully, and ties the resulting token to these blocks
         * as well.
         * 
         * @return builder
         * 
         */
        public Builder tokenBoundCidrs(List<String> tokenBoundCidrs) {
            return tokenBoundCidrs(Output.of(tokenBoundCidrs));
        }

        /**
         * @param tokenBoundCidrs List of CIDR blocks; if set, specifies blocks of IP
         * addresses which can authenticate successfully, and ties the resulting token to these blocks
         * as well.
         * 
         * @return builder
         * 
         */
        public Builder tokenBoundCidrs(String... tokenBoundCidrs) {
            return tokenBoundCidrs(List.of(tokenBoundCidrs));
        }

        /**
         * @param tokenExplicitMaxTtl If set, will encode an
         * [explicit max TTL](https://www.vaultproject.io/docs/concepts/tokens.html#token-time-to-live-periodic-tokens-and-explicit-max-ttls)
         * onto the token in number of seconds. This is a hard cap even if `token_ttl` and
         * `token_max_ttl` would otherwise allow a renewal.
         * 
         * @return builder
         * 
         */
        public Builder tokenExplicitMaxTtl(@Nullable Output<Integer> tokenExplicitMaxTtl) {
            $.tokenExplicitMaxTtl = tokenExplicitMaxTtl;
            return this;
        }

        /**
         * @param tokenExplicitMaxTtl If set, will encode an
         * [explicit max TTL](https://www.vaultproject.io/docs/concepts/tokens.html#token-time-to-live-periodic-tokens-and-explicit-max-ttls)
         * onto the token in number of seconds. This is a hard cap even if `token_ttl` and
         * `token_max_ttl` would otherwise allow a renewal.
         * 
         * @return builder
         * 
         */
        public Builder tokenExplicitMaxTtl(Integer tokenExplicitMaxTtl) {
            return tokenExplicitMaxTtl(Output.of(tokenExplicitMaxTtl));
        }

        /**
         * @param tokenMaxTtl The maximum lifetime for generated tokens in number of seconds.
         * Its current value will be referenced at renewal time.
         * 
         * @return builder
         * 
         */
        public Builder tokenMaxTtl(@Nullable Output<Integer> tokenMaxTtl) {
            $.tokenMaxTtl = tokenMaxTtl;
            return this;
        }

        /**
         * @param tokenMaxTtl The maximum lifetime for generated tokens in number of seconds.
         * Its current value will be referenced at renewal time.
         * 
         * @return builder
         * 
         */
        public Builder tokenMaxTtl(Integer tokenMaxTtl) {
            return tokenMaxTtl(Output.of(tokenMaxTtl));
        }

        /**
         * @param tokenNoDefaultPolicy If set, the default policy will not be set on
         * generated tokens; otherwise it will be added to the policies set in token_policies.
         * 
         * @return builder
         * 
         */
        public Builder tokenNoDefaultPolicy(@Nullable Output<Boolean> tokenNoDefaultPolicy) {
            $.tokenNoDefaultPolicy = tokenNoDefaultPolicy;
            return this;
        }

        /**
         * @param tokenNoDefaultPolicy If set, the default policy will not be set on
         * generated tokens; otherwise it will be added to the policies set in token_policies.
         * 
         * @return builder
         * 
         */
        public Builder tokenNoDefaultPolicy(Boolean tokenNoDefaultPolicy) {
            return tokenNoDefaultPolicy(Output.of(tokenNoDefaultPolicy));
        }

        /**
         * @param tokenNumUses The [maximum number](https://www.vaultproject.io/api-docs/auth/alicloud#token_num_uses)
         * of times a generated token may be used (within its lifetime); 0 means unlimited.
         * 
         * @return builder
         * 
         */
        public Builder tokenNumUses(@Nullable Output<Integer> tokenNumUses) {
            $.tokenNumUses = tokenNumUses;
            return this;
        }

        /**
         * @param tokenNumUses The [maximum number](https://www.vaultproject.io/api-docs/auth/alicloud#token_num_uses)
         * of times a generated token may be used (within its lifetime); 0 means unlimited.
         * 
         * @return builder
         * 
         */
        public Builder tokenNumUses(Integer tokenNumUses) {
            return tokenNumUses(Output.of(tokenNumUses));
        }

        /**
         * @param tokenPeriod If set, indicates that the
         * token generated using this role should never expire. The token should be renewed within the
         * duration specified by this value. At each renewal, the token&#39;s TTL will be set to the
         * value of this field. Specified in seconds.
         * 
         * @return builder
         * 
         */
        public Builder tokenPeriod(@Nullable Output<Integer> tokenPeriod) {
            $.tokenPeriod = tokenPeriod;
            return this;
        }

        /**
         * @param tokenPeriod If set, indicates that the
         * token generated using this role should never expire. The token should be renewed within the
         * duration specified by this value. At each renewal, the token&#39;s TTL will be set to the
         * value of this field. Specified in seconds.
         * 
         * @return builder
         * 
         */
        public Builder tokenPeriod(Integer tokenPeriod) {
            return tokenPeriod(Output.of(tokenPeriod));
        }

        /**
         * @param tokenPolicies List of policies to encode onto generated tokens. Depending
         * on the auth method, this list may be supplemented by user/group/other values.
         * 
         * @return builder
         * 
         */
        public Builder tokenPolicies(@Nullable Output<List<String>> tokenPolicies) {
            $.tokenPolicies = tokenPolicies;
            return this;
        }

        /**
         * @param tokenPolicies List of policies to encode onto generated tokens. Depending
         * on the auth method, this list may be supplemented by user/group/other values.
         * 
         * @return builder
         * 
         */
        public Builder tokenPolicies(List<String> tokenPolicies) {
            return tokenPolicies(Output.of(tokenPolicies));
        }

        /**
         * @param tokenPolicies List of policies to encode onto generated tokens. Depending
         * on the auth method, this list may be supplemented by user/group/other values.
         * 
         * @return builder
         * 
         */
        public Builder tokenPolicies(String... tokenPolicies) {
            return tokenPolicies(List.of(tokenPolicies));
        }

        /**
         * @param tokenTtl The incremental lifetime for generated tokens in number of seconds.
         * Its current value will be referenced at renewal time.
         * 
         * @return builder
         * 
         */
        public Builder tokenTtl(@Nullable Output<Integer> tokenTtl) {
            $.tokenTtl = tokenTtl;
            return this;
        }

        /**
         * @param tokenTtl The incremental lifetime for generated tokens in number of seconds.
         * Its current value will be referenced at renewal time.
         * 
         * @return builder
         * 
         */
        public Builder tokenTtl(Integer tokenTtl) {
            return tokenTtl(Output.of(tokenTtl));
        }

        /**
         * @param tokenType The type of token that should be generated. Can be `service`,
         * `batch`, or `default` to use the mount&#39;s tuned default (which unless changed will be
         * `service` tokens). For token store roles, there are two additional possibilities:
         * `default-service` and `default-batch` which specify the type to return unless the client
         * requests a different type at generation time.
         * 
         * @return builder
         * 
         */
        public Builder tokenType(@Nullable Output<String> tokenType) {
            $.tokenType = tokenType;
            return this;
        }

        /**
         * @param tokenType The type of token that should be generated. Can be `service`,
         * `batch`, or `default` to use the mount&#39;s tuned default (which unless changed will be
         * `service` tokens). For token store roles, there are two additional possibilities:
         * `default-service` and `default-batch` which specify the type to return unless the client
         * requests a different type at generation time.
         * 
         * @return builder
         * 
         */
        public Builder tokenType(String tokenType) {
            return tokenType(Output.of(tokenType));
        }

        public AuthBackendRoleArgs build() {
            if ($.arn == null) {
                throw new MissingRequiredPropertyException("AuthBackendRoleArgs", "arn");
            }
            if ($.role == null) {
                throw new MissingRequiredPropertyException("AuthBackendRoleArgs", "role");
            }
            return $;
        }
    }

}
