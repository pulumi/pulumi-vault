// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vault.transform.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class RoleState extends com.pulumi.resources.ResourceArgs {

    public static final RoleState Empty = new RoleState();

    /**
     * The name of the role.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return The name of the role.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * Path to where the back-end is mounted within Vault.
     * 
     */
    @Import(name="path")
    private @Nullable Output<String> path;

    /**
     * @return Path to where the back-end is mounted within Vault.
     * 
     */
    public Optional<Output<String>> path() {
        return Optional.ofNullable(this.path);
    }

    /**
     * A comma separated string or slice of transformations to use.
     * 
     */
    @Import(name="transformations")
    private @Nullable Output<List<String>> transformations;

    /**
     * @return A comma separated string or slice of transformations to use.
     * 
     */
    public Optional<Output<List<String>>> transformations() {
        return Optional.ofNullable(this.transformations);
    }

    private RoleState() {}

    private RoleState(RoleState $) {
        this.name = $.name;
        this.path = $.path;
        this.transformations = $.transformations;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(RoleState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private RoleState $;

        public Builder() {
            $ = new RoleState();
        }

        public Builder(RoleState defaults) {
            $ = new RoleState(Objects.requireNonNull(defaults));
        }

        /**
         * @param name The name of the role.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name The name of the role.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param path Path to where the back-end is mounted within Vault.
         * 
         * @return builder
         * 
         */
        public Builder path(@Nullable Output<String> path) {
            $.path = path;
            return this;
        }

        /**
         * @param path Path to where the back-end is mounted within Vault.
         * 
         * @return builder
         * 
         */
        public Builder path(String path) {
            return path(Output.of(path));
        }

        /**
         * @param transformations A comma separated string or slice of transformations to use.
         * 
         * @return builder
         * 
         */
        public Builder transformations(@Nullable Output<List<String>> transformations) {
            $.transformations = transformations;
            return this;
        }

        /**
         * @param transformations A comma separated string or slice of transformations to use.
         * 
         * @return builder
         * 
         */
        public Builder transformations(List<String> transformations) {
            return transformations(Output.of(transformations));
        }

        /**
         * @param transformations A comma separated string or slice of transformations to use.
         * 
         * @return builder
         * 
         */
        public Builder transformations(String... transformations) {
            return transformations(List.of(transformations));
        }

        public RoleState build() {
            return $;
        }
    }

}