// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vault.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ProviderAuthLoginRadiusArgs extends com.pulumi.resources.ResourceArgs {

    public static final ProviderAuthLoginRadiusArgs Empty = new ProviderAuthLoginRadiusArgs();

    /**
     * The path where the authentication engine is mounted.
     * 
     */
    @Import(name="mount")
    private @Nullable Output<String> mount;

    /**
     * @return The path where the authentication engine is mounted.
     * 
     */
    public Optional<Output<String>> mount() {
        return Optional.ofNullable(this.mount);
    }

    /**
     * The authentication engine&#39;s namespace. Conflicts with use_root_namespace
     * 
     */
    @Import(name="namespace")
    private @Nullable Output<String> namespace;

    /**
     * @return The authentication engine&#39;s namespace. Conflicts with use_root_namespace
     * 
     */
    public Optional<Output<String>> namespace() {
        return Optional.ofNullable(this.namespace);
    }

    /**
     * The Radius password for username.
     * 
     */
    @Import(name="password", required=true)
    private Output<String> password;

    /**
     * @return The Radius password for username.
     * 
     */
    public Output<String> password() {
        return this.password;
    }

    /**
     * Authenticate to the root Vault namespace. Conflicts with namespace
     * 
     */
    @Import(name="useRootNamespace")
    private @Nullable Output<Boolean> useRootNamespace;

    /**
     * @return Authenticate to the root Vault namespace. Conflicts with namespace
     * 
     */
    public Optional<Output<Boolean>> useRootNamespace() {
        return Optional.ofNullable(this.useRootNamespace);
    }

    /**
     * The Radius username.
     * 
     */
    @Import(name="username", required=true)
    private Output<String> username;

    /**
     * @return The Radius username.
     * 
     */
    public Output<String> username() {
        return this.username;
    }

    private ProviderAuthLoginRadiusArgs() {}

    private ProviderAuthLoginRadiusArgs(ProviderAuthLoginRadiusArgs $) {
        this.mount = $.mount;
        this.namespace = $.namespace;
        this.password = $.password;
        this.useRootNamespace = $.useRootNamespace;
        this.username = $.username;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ProviderAuthLoginRadiusArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ProviderAuthLoginRadiusArgs $;

        public Builder() {
            $ = new ProviderAuthLoginRadiusArgs();
        }

        public Builder(ProviderAuthLoginRadiusArgs defaults) {
            $ = new ProviderAuthLoginRadiusArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param mount The path where the authentication engine is mounted.
         * 
         * @return builder
         * 
         */
        public Builder mount(@Nullable Output<String> mount) {
            $.mount = mount;
            return this;
        }

        /**
         * @param mount The path where the authentication engine is mounted.
         * 
         * @return builder
         * 
         */
        public Builder mount(String mount) {
            return mount(Output.of(mount));
        }

        /**
         * @param namespace The authentication engine&#39;s namespace. Conflicts with use_root_namespace
         * 
         * @return builder
         * 
         */
        public Builder namespace(@Nullable Output<String> namespace) {
            $.namespace = namespace;
            return this;
        }

        /**
         * @param namespace The authentication engine&#39;s namespace. Conflicts with use_root_namespace
         * 
         * @return builder
         * 
         */
        public Builder namespace(String namespace) {
            return namespace(Output.of(namespace));
        }

        /**
         * @param password The Radius password for username.
         * 
         * @return builder
         * 
         */
        public Builder password(Output<String> password) {
            $.password = password;
            return this;
        }

        /**
         * @param password The Radius password for username.
         * 
         * @return builder
         * 
         */
        public Builder password(String password) {
            return password(Output.of(password));
        }

        /**
         * @param useRootNamespace Authenticate to the root Vault namespace. Conflicts with namespace
         * 
         * @return builder
         * 
         */
        public Builder useRootNamespace(@Nullable Output<Boolean> useRootNamespace) {
            $.useRootNamespace = useRootNamespace;
            return this;
        }

        /**
         * @param useRootNamespace Authenticate to the root Vault namespace. Conflicts with namespace
         * 
         * @return builder
         * 
         */
        public Builder useRootNamespace(Boolean useRootNamespace) {
            return useRootNamespace(Output.of(useRootNamespace));
        }

        /**
         * @param username The Radius username.
         * 
         * @return builder
         * 
         */
        public Builder username(Output<String> username) {
            $.username = username;
            return this;
        }

        /**
         * @param username The Radius username.
         * 
         * @return builder
         * 
         */
        public Builder username(String username) {
            return username(Output.of(username));
        }

        public ProviderAuthLoginRadiusArgs build() {
            if ($.password == null) {
                throw new MissingRequiredPropertyException("ProviderAuthLoginRadiusArgs", "password");
            }
            if ($.username == null) {
                throw new MissingRequiredPropertyException("ProviderAuthLoginRadiusArgs", "username");
            }
            return $;
        }
    }

}
