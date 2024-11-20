// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vault.appRole.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class AuthBackendRoleSecretIdState extends com.pulumi.resources.ResourceArgs {

    public static final AuthBackendRoleSecretIdState Empty = new AuthBackendRoleSecretIdState();

    /**
     * The unique ID for this SecretID that can be safely logged.
     * 
     */
    @Import(name="accessor")
    private @Nullable Output<String> accessor;

    /**
     * @return The unique ID for this SecretID that can be safely logged.
     * 
     */
    public Optional<Output<String>> accessor() {
        return Optional.ofNullable(this.accessor);
    }

    /**
     * Unique name of the auth backend to configure.
     * 
     */
    @Import(name="backend")
    private @Nullable Output<String> backend;

    /**
     * @return Unique name of the auth backend to configure.
     * 
     */
    public Optional<Output<String>> backend() {
        return Optional.ofNullable(this.backend);
    }

    /**
     * If set, specifies blocks of IP addresses which can
     * perform the login operation using this SecretID.
     * 
     */
    @Import(name="cidrLists")
    private @Nullable Output<List<String>> cidrLists;

    /**
     * @return If set, specifies blocks of IP addresses which can
     * perform the login operation using this SecretID.
     * 
     */
    public Optional<Output<List<String>>> cidrLists() {
        return Optional.ofNullable(this.cidrLists);
    }

    /**
     * A JSON-encoded string containing metadata in
     * key-value pairs to be set on tokens issued with this SecretID.
     * 
     */
    @Import(name="metadata")
    private @Nullable Output<String> metadata;

    /**
     * @return A JSON-encoded string containing metadata in
     * key-value pairs to be set on tokens issued with this SecretID.
     * 
     */
    public Optional<Output<String>> metadata() {
        return Optional.ofNullable(this.metadata);
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
     * The number of uses for the secret-id.
     * 
     */
    @Import(name="numUses")
    private @Nullable Output<Integer> numUses;

    /**
     * @return The number of uses for the secret-id.
     * 
     */
    public Optional<Output<Integer>> numUses() {
        return Optional.ofNullable(this.numUses);
    }

    /**
     * The name of the role to create the SecretID for.
     * 
     */
    @Import(name="roleName")
    private @Nullable Output<String> roleName;

    /**
     * @return The name of the role to create the SecretID for.
     * 
     */
    public Optional<Output<String>> roleName() {
        return Optional.ofNullable(this.roleName);
    }

    /**
     * The SecretID to be created. If set, uses &#34;Push&#34;
     * mode.  Defaults to Vault auto-generating SecretIDs.
     * 
     */
    @Import(name="secretId")
    private @Nullable Output<String> secretId;

    /**
     * @return The SecretID to be created. If set, uses &#34;Push&#34;
     * mode.  Defaults to Vault auto-generating SecretIDs.
     * 
     */
    public Optional<Output<String>> secretId() {
        return Optional.ofNullable(this.secretId);
    }

    /**
     * The TTL duration of the SecretID.
     * 
     */
    @Import(name="ttl")
    private @Nullable Output<Integer> ttl;

    /**
     * @return The TTL duration of the SecretID.
     * 
     */
    public Optional<Output<Integer>> ttl() {
        return Optional.ofNullable(this.ttl);
    }

    /**
     * Set to `true` to use the wrapped secret-id accessor as the resource ID.
     * If `false` (default value), a fresh secret ID will be regenerated whenever the wrapping token is expired or
     * invalidated through unwrapping.
     * 
     */
    @Import(name="withWrappedAccessor")
    private @Nullable Output<Boolean> withWrappedAccessor;

    /**
     * @return Set to `true` to use the wrapped secret-id accessor as the resource ID.
     * If `false` (default value), a fresh secret ID will be regenerated whenever the wrapping token is expired or
     * invalidated through unwrapping.
     * 
     */
    public Optional<Output<Boolean>> withWrappedAccessor() {
        return Optional.ofNullable(this.withWrappedAccessor);
    }

    /**
     * The unique ID for the response-wrapped SecretID that can
     * be safely logged.
     * 
     */
    @Import(name="wrappingAccessor")
    private @Nullable Output<String> wrappingAccessor;

    /**
     * @return The unique ID for the response-wrapped SecretID that can
     * be safely logged.
     * 
     */
    public Optional<Output<String>> wrappingAccessor() {
        return Optional.ofNullable(this.wrappingAccessor);
    }

    /**
     * The token used to retrieve a response-wrapped SecretID.
     * 
     */
    @Import(name="wrappingToken")
    private @Nullable Output<String> wrappingToken;

    /**
     * @return The token used to retrieve a response-wrapped SecretID.
     * 
     */
    public Optional<Output<String>> wrappingToken() {
        return Optional.ofNullable(this.wrappingToken);
    }

    /**
     * If set, the SecretID response will be
     * [response-wrapped](https://www.vaultproject.io/docs/concepts/response-wrapping)
     * and available for the duration specified. Only a single unwrapping of the
     * token is allowed.
     * 
     */
    @Import(name="wrappingTtl")
    private @Nullable Output<String> wrappingTtl;

    /**
     * @return If set, the SecretID response will be
     * [response-wrapped](https://www.vaultproject.io/docs/concepts/response-wrapping)
     * and available for the duration specified. Only a single unwrapping of the
     * token is allowed.
     * 
     */
    public Optional<Output<String>> wrappingTtl() {
        return Optional.ofNullable(this.wrappingTtl);
    }

    private AuthBackendRoleSecretIdState() {}

    private AuthBackendRoleSecretIdState(AuthBackendRoleSecretIdState $) {
        this.accessor = $.accessor;
        this.backend = $.backend;
        this.cidrLists = $.cidrLists;
        this.metadata = $.metadata;
        this.namespace = $.namespace;
        this.numUses = $.numUses;
        this.roleName = $.roleName;
        this.secretId = $.secretId;
        this.ttl = $.ttl;
        this.withWrappedAccessor = $.withWrappedAccessor;
        this.wrappingAccessor = $.wrappingAccessor;
        this.wrappingToken = $.wrappingToken;
        this.wrappingTtl = $.wrappingTtl;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(AuthBackendRoleSecretIdState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private AuthBackendRoleSecretIdState $;

        public Builder() {
            $ = new AuthBackendRoleSecretIdState();
        }

        public Builder(AuthBackendRoleSecretIdState defaults) {
            $ = new AuthBackendRoleSecretIdState(Objects.requireNonNull(defaults));
        }

        /**
         * @param accessor The unique ID for this SecretID that can be safely logged.
         * 
         * @return builder
         * 
         */
        public Builder accessor(@Nullable Output<String> accessor) {
            $.accessor = accessor;
            return this;
        }

        /**
         * @param accessor The unique ID for this SecretID that can be safely logged.
         * 
         * @return builder
         * 
         */
        public Builder accessor(String accessor) {
            return accessor(Output.of(accessor));
        }

        /**
         * @param backend Unique name of the auth backend to configure.
         * 
         * @return builder
         * 
         */
        public Builder backend(@Nullable Output<String> backend) {
            $.backend = backend;
            return this;
        }

        /**
         * @param backend Unique name of the auth backend to configure.
         * 
         * @return builder
         * 
         */
        public Builder backend(String backend) {
            return backend(Output.of(backend));
        }

        /**
         * @param cidrLists If set, specifies blocks of IP addresses which can
         * perform the login operation using this SecretID.
         * 
         * @return builder
         * 
         */
        public Builder cidrLists(@Nullable Output<List<String>> cidrLists) {
            $.cidrLists = cidrLists;
            return this;
        }

        /**
         * @param cidrLists If set, specifies blocks of IP addresses which can
         * perform the login operation using this SecretID.
         * 
         * @return builder
         * 
         */
        public Builder cidrLists(List<String> cidrLists) {
            return cidrLists(Output.of(cidrLists));
        }

        /**
         * @param cidrLists If set, specifies blocks of IP addresses which can
         * perform the login operation using this SecretID.
         * 
         * @return builder
         * 
         */
        public Builder cidrLists(String... cidrLists) {
            return cidrLists(List.of(cidrLists));
        }

        /**
         * @param metadata A JSON-encoded string containing metadata in
         * key-value pairs to be set on tokens issued with this SecretID.
         * 
         * @return builder
         * 
         */
        public Builder metadata(@Nullable Output<String> metadata) {
            $.metadata = metadata;
            return this;
        }

        /**
         * @param metadata A JSON-encoded string containing metadata in
         * key-value pairs to be set on tokens issued with this SecretID.
         * 
         * @return builder
         * 
         */
        public Builder metadata(String metadata) {
            return metadata(Output.of(metadata));
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
         * @param numUses The number of uses for the secret-id.
         * 
         * @return builder
         * 
         */
        public Builder numUses(@Nullable Output<Integer> numUses) {
            $.numUses = numUses;
            return this;
        }

        /**
         * @param numUses The number of uses for the secret-id.
         * 
         * @return builder
         * 
         */
        public Builder numUses(Integer numUses) {
            return numUses(Output.of(numUses));
        }

        /**
         * @param roleName The name of the role to create the SecretID for.
         * 
         * @return builder
         * 
         */
        public Builder roleName(@Nullable Output<String> roleName) {
            $.roleName = roleName;
            return this;
        }

        /**
         * @param roleName The name of the role to create the SecretID for.
         * 
         * @return builder
         * 
         */
        public Builder roleName(String roleName) {
            return roleName(Output.of(roleName));
        }

        /**
         * @param secretId The SecretID to be created. If set, uses &#34;Push&#34;
         * mode.  Defaults to Vault auto-generating SecretIDs.
         * 
         * @return builder
         * 
         */
        public Builder secretId(@Nullable Output<String> secretId) {
            $.secretId = secretId;
            return this;
        }

        /**
         * @param secretId The SecretID to be created. If set, uses &#34;Push&#34;
         * mode.  Defaults to Vault auto-generating SecretIDs.
         * 
         * @return builder
         * 
         */
        public Builder secretId(String secretId) {
            return secretId(Output.of(secretId));
        }

        /**
         * @param ttl The TTL duration of the SecretID.
         * 
         * @return builder
         * 
         */
        public Builder ttl(@Nullable Output<Integer> ttl) {
            $.ttl = ttl;
            return this;
        }

        /**
         * @param ttl The TTL duration of the SecretID.
         * 
         * @return builder
         * 
         */
        public Builder ttl(Integer ttl) {
            return ttl(Output.of(ttl));
        }

        /**
         * @param withWrappedAccessor Set to `true` to use the wrapped secret-id accessor as the resource ID.
         * If `false` (default value), a fresh secret ID will be regenerated whenever the wrapping token is expired or
         * invalidated through unwrapping.
         * 
         * @return builder
         * 
         */
        public Builder withWrappedAccessor(@Nullable Output<Boolean> withWrappedAccessor) {
            $.withWrappedAccessor = withWrappedAccessor;
            return this;
        }

        /**
         * @param withWrappedAccessor Set to `true` to use the wrapped secret-id accessor as the resource ID.
         * If `false` (default value), a fresh secret ID will be regenerated whenever the wrapping token is expired or
         * invalidated through unwrapping.
         * 
         * @return builder
         * 
         */
        public Builder withWrappedAccessor(Boolean withWrappedAccessor) {
            return withWrappedAccessor(Output.of(withWrappedAccessor));
        }

        /**
         * @param wrappingAccessor The unique ID for the response-wrapped SecretID that can
         * be safely logged.
         * 
         * @return builder
         * 
         */
        public Builder wrappingAccessor(@Nullable Output<String> wrappingAccessor) {
            $.wrappingAccessor = wrappingAccessor;
            return this;
        }

        /**
         * @param wrappingAccessor The unique ID for the response-wrapped SecretID that can
         * be safely logged.
         * 
         * @return builder
         * 
         */
        public Builder wrappingAccessor(String wrappingAccessor) {
            return wrappingAccessor(Output.of(wrappingAccessor));
        }

        /**
         * @param wrappingToken The token used to retrieve a response-wrapped SecretID.
         * 
         * @return builder
         * 
         */
        public Builder wrappingToken(@Nullable Output<String> wrappingToken) {
            $.wrappingToken = wrappingToken;
            return this;
        }

        /**
         * @param wrappingToken The token used to retrieve a response-wrapped SecretID.
         * 
         * @return builder
         * 
         */
        public Builder wrappingToken(String wrappingToken) {
            return wrappingToken(Output.of(wrappingToken));
        }

        /**
         * @param wrappingTtl If set, the SecretID response will be
         * [response-wrapped](https://www.vaultproject.io/docs/concepts/response-wrapping)
         * and available for the duration specified. Only a single unwrapping of the
         * token is allowed.
         * 
         * @return builder
         * 
         */
        public Builder wrappingTtl(@Nullable Output<String> wrappingTtl) {
            $.wrappingTtl = wrappingTtl;
            return this;
        }

        /**
         * @param wrappingTtl If set, the SecretID response will be
         * [response-wrapped](https://www.vaultproject.io/docs/concepts/response-wrapping)
         * and available for the duration specified. Only a single unwrapping of the
         * token is allowed.
         * 
         * @return builder
         * 
         */
        public Builder wrappingTtl(String wrappingTtl) {
            return wrappingTtl(Output.of(wrappingTtl));
        }

        public AuthBackendRoleSecretIdState build() {
            return $;
        }
    }

}
