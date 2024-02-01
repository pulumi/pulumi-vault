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


public final class ProviderAuthLoginJwtArgs extends com.pulumi.resources.ResourceArgs {

    public static final ProviderAuthLoginJwtArgs Empty = new ProviderAuthLoginJwtArgs();

    /**
     * A signed JSON Web Token.
     * 
     */
    @Import(name="jwt", required=true)
    private Output<String> jwt;

    /**
     * @return A signed JSON Web Token.
     * 
     */
    public Output<String> jwt() {
        return this.jwt;
    }

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
     * Name of the login role.
     * 
     */
    @Import(name="role", required=true)
    private Output<String> role;

    /**
     * @return Name of the login role.
     * 
     */
    public Output<String> role() {
        return this.role;
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

    private ProviderAuthLoginJwtArgs() {}

    private ProviderAuthLoginJwtArgs(ProviderAuthLoginJwtArgs $) {
        this.jwt = $.jwt;
        this.mount = $.mount;
        this.namespace = $.namespace;
        this.role = $.role;
        this.useRootNamespace = $.useRootNamespace;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ProviderAuthLoginJwtArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ProviderAuthLoginJwtArgs $;

        public Builder() {
            $ = new ProviderAuthLoginJwtArgs();
        }

        public Builder(ProviderAuthLoginJwtArgs defaults) {
            $ = new ProviderAuthLoginJwtArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param jwt A signed JSON Web Token.
         * 
         * @return builder
         * 
         */
        public Builder jwt(Output<String> jwt) {
            $.jwt = jwt;
            return this;
        }

        /**
         * @param jwt A signed JSON Web Token.
         * 
         * @return builder
         * 
         */
        public Builder jwt(String jwt) {
            return jwt(Output.of(jwt));
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
         * @param role Name of the login role.
         * 
         * @return builder
         * 
         */
        public Builder role(Output<String> role) {
            $.role = role;
            return this;
        }

        /**
         * @param role Name of the login role.
         * 
         * @return builder
         * 
         */
        public Builder role(String role) {
            return role(Output.of(role));
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

        public ProviderAuthLoginJwtArgs build() {
            if ($.jwt == null) {
                throw new MissingRequiredPropertyException("ProviderAuthLoginJwtArgs", "jwt");
            }
            if ($.role == null) {
                throw new MissingRequiredPropertyException("ProviderAuthLoginJwtArgs", "role");
            }
            return $;
        }
    }

}
