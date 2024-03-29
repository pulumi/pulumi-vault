// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vault.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Boolean;
import java.lang.String;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ProviderAuthLoginArgs extends com.pulumi.resources.ResourceArgs {

    public static final ProviderAuthLoginArgs Empty = new ProviderAuthLoginArgs();

    @Import(name="method")
    private @Nullable Output<String> method;

    public Optional<Output<String>> method() {
        return Optional.ofNullable(this.method);
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

    @Import(name="parameters")
    private @Nullable Output<Map<String,String>> parameters;

    public Optional<Output<Map<String,String>>> parameters() {
        return Optional.ofNullable(this.parameters);
    }

    @Import(name="path", required=true)
    private Output<String> path;

    public Output<String> path() {
        return this.path;
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

    private ProviderAuthLoginArgs() {}

    private ProviderAuthLoginArgs(ProviderAuthLoginArgs $) {
        this.method = $.method;
        this.namespace = $.namespace;
        this.parameters = $.parameters;
        this.path = $.path;
        this.useRootNamespace = $.useRootNamespace;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ProviderAuthLoginArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ProviderAuthLoginArgs $;

        public Builder() {
            $ = new ProviderAuthLoginArgs();
        }

        public Builder(ProviderAuthLoginArgs defaults) {
            $ = new ProviderAuthLoginArgs(Objects.requireNonNull(defaults));
        }

        public Builder method(@Nullable Output<String> method) {
            $.method = method;
            return this;
        }

        public Builder method(String method) {
            return method(Output.of(method));
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

        public Builder parameters(@Nullable Output<Map<String,String>> parameters) {
            $.parameters = parameters;
            return this;
        }

        public Builder parameters(Map<String,String> parameters) {
            return parameters(Output.of(parameters));
        }

        public Builder path(Output<String> path) {
            $.path = path;
            return this;
        }

        public Builder path(String path) {
            return path(Output.of(path));
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

        public ProviderAuthLoginArgs build() {
            if ($.path == null) {
                throw new MissingRequiredPropertyException("ProviderAuthLoginArgs", "path");
            }
            return $;
        }
    }

}
