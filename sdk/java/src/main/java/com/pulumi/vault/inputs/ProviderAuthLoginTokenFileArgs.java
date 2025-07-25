// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vault.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ProviderAuthLoginTokenFileArgs extends com.pulumi.resources.ResourceArgs {

    public static final ProviderAuthLoginTokenFileArgs Empty = new ProviderAuthLoginTokenFileArgs();

    /**
     * The name of a file containing a single line that is a valid Vault token
     * 
     */
    @Import(name="filename")
    private @Nullable Output<String> filename;

    /**
     * @return The name of a file containing a single line that is a valid Vault token
     * 
     */
    public Optional<Output<String>> filename() {
        return Optional.ofNullable(this.filename);
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

    private ProviderAuthLoginTokenFileArgs() {}

    private ProviderAuthLoginTokenFileArgs(ProviderAuthLoginTokenFileArgs $) {
        this.filename = $.filename;
        this.namespace = $.namespace;
        this.useRootNamespace = $.useRootNamespace;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ProviderAuthLoginTokenFileArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ProviderAuthLoginTokenFileArgs $;

        public Builder() {
            $ = new ProviderAuthLoginTokenFileArgs();
        }

        public Builder(ProviderAuthLoginTokenFileArgs defaults) {
            $ = new ProviderAuthLoginTokenFileArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param filename The name of a file containing a single line that is a valid Vault token
         * 
         * @return builder
         * 
         */
        public Builder filename(@Nullable Output<String> filename) {
            $.filename = filename;
            return this;
        }

        /**
         * @param filename The name of a file containing a single line that is a valid Vault token
         * 
         * @return builder
         * 
         */
        public Builder filename(String filename) {
            return filename(Output.of(filename));
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

        public ProviderAuthLoginTokenFileArgs build() {
            return $;
        }
    }

}
