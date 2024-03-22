// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vault.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class NomadSecretRoleState extends com.pulumi.resources.ResourceArgs {

    public static final NomadSecretRoleState Empty = new NomadSecretRoleState();

    /**
     * The unique path this backend should be mounted at.
     * 
     */
    @Import(name="backend")
    private @Nullable Output<String> backend;

    /**
     * @return The unique path this backend should be mounted at.
     * 
     */
    public Optional<Output<String>> backend() {
        return Optional.ofNullable(this.backend);
    }

    /**
     * Specifies if the generated token should be global. Defaults to
     * false.
     * 
     */
    @Import(name="global")
    private @Nullable Output<Boolean> global;

    /**
     * @return Specifies if the generated token should be global. Defaults to
     * false.
     * 
     */
    public Optional<Output<Boolean>> global() {
        return Optional.ofNullable(this.global);
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
     * List of policies attached to the generated token. This setting is only used
     * when `type` is &#39;client&#39;.
     * 
     */
    @Import(name="policies")
    private @Nullable Output<List<String>> policies;

    /**
     * @return List of policies attached to the generated token. This setting is only used
     * when `type` is &#39;client&#39;.
     * 
     */
    public Optional<Output<List<String>>> policies() {
        return Optional.ofNullable(this.policies);
    }

    /**
     * The name to identify this role within the backend.
     * Must be unique within the backend.
     * 
     */
    @Import(name="role")
    private @Nullable Output<String> role;

    /**
     * @return The name to identify this role within the backend.
     * Must be unique within the backend.
     * 
     */
    public Optional<Output<String>> role() {
        return Optional.ofNullable(this.role);
    }

    /**
     * Specifies the type of token to create when using this role. Valid
     * settings are &#39;client&#39; and &#39;management&#39;. Defaults to &#39;client&#39;.
     * 
     */
    @Import(name="type")
    private @Nullable Output<String> type;

    /**
     * @return Specifies the type of token to create when using this role. Valid
     * settings are &#39;client&#39; and &#39;management&#39;. Defaults to &#39;client&#39;.
     * 
     */
    public Optional<Output<String>> type() {
        return Optional.ofNullable(this.type);
    }

    private NomadSecretRoleState() {}

    private NomadSecretRoleState(NomadSecretRoleState $) {
        this.backend = $.backend;
        this.global = $.global;
        this.namespace = $.namespace;
        this.policies = $.policies;
        this.role = $.role;
        this.type = $.type;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(NomadSecretRoleState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private NomadSecretRoleState $;

        public Builder() {
            $ = new NomadSecretRoleState();
        }

        public Builder(NomadSecretRoleState defaults) {
            $ = new NomadSecretRoleState(Objects.requireNonNull(defaults));
        }

        /**
         * @param backend The unique path this backend should be mounted at.
         * 
         * @return builder
         * 
         */
        public Builder backend(@Nullable Output<String> backend) {
            $.backend = backend;
            return this;
        }

        /**
         * @param backend The unique path this backend should be mounted at.
         * 
         * @return builder
         * 
         */
        public Builder backend(String backend) {
            return backend(Output.of(backend));
        }

        /**
         * @param global Specifies if the generated token should be global. Defaults to
         * false.
         * 
         * @return builder
         * 
         */
        public Builder global(@Nullable Output<Boolean> global) {
            $.global = global;
            return this;
        }

        /**
         * @param global Specifies if the generated token should be global. Defaults to
         * false.
         * 
         * @return builder
         * 
         */
        public Builder global(Boolean global) {
            return global(Output.of(global));
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
         * @param policies List of policies attached to the generated token. This setting is only used
         * when `type` is &#39;client&#39;.
         * 
         * @return builder
         * 
         */
        public Builder policies(@Nullable Output<List<String>> policies) {
            $.policies = policies;
            return this;
        }

        /**
         * @param policies List of policies attached to the generated token. This setting is only used
         * when `type` is &#39;client&#39;.
         * 
         * @return builder
         * 
         */
        public Builder policies(List<String> policies) {
            return policies(Output.of(policies));
        }

        /**
         * @param policies List of policies attached to the generated token. This setting is only used
         * when `type` is &#39;client&#39;.
         * 
         * @return builder
         * 
         */
        public Builder policies(String... policies) {
            return policies(List.of(policies));
        }

        /**
         * @param role The name to identify this role within the backend.
         * Must be unique within the backend.
         * 
         * @return builder
         * 
         */
        public Builder role(@Nullable Output<String> role) {
            $.role = role;
            return this;
        }

        /**
         * @param role The name to identify this role within the backend.
         * Must be unique within the backend.
         * 
         * @return builder
         * 
         */
        public Builder role(String role) {
            return role(Output.of(role));
        }

        /**
         * @param type Specifies the type of token to create when using this role. Valid
         * settings are &#39;client&#39; and &#39;management&#39;. Defaults to &#39;client&#39;.
         * 
         * @return builder
         * 
         */
        public Builder type(@Nullable Output<String> type) {
            $.type = type;
            return this;
        }

        /**
         * @param type Specifies the type of token to create when using this role. Valid
         * settings are &#39;client&#39; and &#39;management&#39;. Defaults to &#39;client&#39;.
         * 
         * @return builder
         * 
         */
        public Builder type(String type) {
            return type(Output.of(type));
        }

        public NomadSecretRoleState build() {
            return $;
        }
    }

}
