// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vault.transform;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class AlphabetArgs extends com.pulumi.resources.ResourceArgs {

    public static final AlphabetArgs Empty = new AlphabetArgs();

    /**
     * A string of characters that contains the alphabet set.
     * 
     */
    @Import(name="alphabet")
    private @Nullable Output<String> alphabet;

    /**
     * @return A string of characters that contains the alphabet set.
     * 
     */
    public Optional<Output<String>> alphabet() {
        return Optional.ofNullable(this.alphabet);
    }

    /**
     * The name of the alphabet.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return The name of the alphabet.
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
     * Path to where the back-end is mounted within Vault.
     * 
     */
    @Import(name="path", required=true)
    private Output<String> path;

    /**
     * @return Path to where the back-end is mounted within Vault.
     * 
     */
    public Output<String> path() {
        return this.path;
    }

    private AlphabetArgs() {}

    private AlphabetArgs(AlphabetArgs $) {
        this.alphabet = $.alphabet;
        this.name = $.name;
        this.namespace = $.namespace;
        this.path = $.path;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(AlphabetArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private AlphabetArgs $;

        public Builder() {
            $ = new AlphabetArgs();
        }

        public Builder(AlphabetArgs defaults) {
            $ = new AlphabetArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param alphabet A string of characters that contains the alphabet set.
         * 
         * @return builder
         * 
         */
        public Builder alphabet(@Nullable Output<String> alphabet) {
            $.alphabet = alphabet;
            return this;
        }

        /**
         * @param alphabet A string of characters that contains the alphabet set.
         * 
         * @return builder
         * 
         */
        public Builder alphabet(String alphabet) {
            return alphabet(Output.of(alphabet));
        }

        /**
         * @param name The name of the alphabet.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name The name of the alphabet.
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
         * @param path Path to where the back-end is mounted within Vault.
         * 
         * @return builder
         * 
         */
        public Builder path(Output<String> path) {
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

        public AlphabetArgs build() {
            if ($.path == null) {
                throw new MissingRequiredPropertyException("AlphabetArgs", "path");
            }
            return $;
        }
    }

}