// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vault.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class PluginState extends com.pulumi.resources.ResourceArgs {

    public static final PluginState Empty = new PluginState();

    /**
     * List of additional args to pass to the plugin.
     * 
     */
    @Import(name="args")
    private @Nullable Output<List<String>> args;

    /**
     * @return List of additional args to pass to the plugin.
     * 
     */
    public Optional<Output<List<String>>> args() {
        return Optional.ofNullable(this.args);
    }

    /**
     * Command to execute the plugin, relative to the server&#39;s configured `plugin_directory`.
     * 
     */
    @Import(name="command")
    private @Nullable Output<String> command;

    /**
     * @return Command to execute the plugin, relative to the server&#39;s configured `plugin_directory`.
     * 
     */
    public Optional<Output<String>> command() {
        return Optional.ofNullable(this.command);
    }

    /**
     * List of additional environment variables to run the plugin with in KEY=VALUE form.
     * 
     */
    @Import(name="envs")
    private @Nullable Output<List<String>> envs;

    /**
     * @return List of additional environment variables to run the plugin with in KEY=VALUE form.
     * 
     */
    public Optional<Output<List<String>>> envs() {
        return Optional.ofNullable(this.envs);
    }

    /**
     * Name of the plugin.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return Name of the plugin.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * Specifies OCI image to run. If specified, setting
     * `command`, `args`, and `env` will update the container&#39;s entrypoint, args, and
     * environment variables (append-only) respectively.
     * 
     */
    @Import(name="ociImage")
    private @Nullable Output<String> ociImage;

    /**
     * @return Specifies OCI image to run. If specified, setting
     * `command`, `args`, and `env` will update the container&#39;s entrypoint, args, and
     * environment variables (append-only) respectively.
     * 
     */
    public Optional<Output<String>> ociImage() {
        return Optional.ofNullable(this.ociImage);
    }

    /**
     * Vault plugin runtime to use if `oci_image` is specified.
     * 
     */
    @Import(name="runtime")
    private @Nullable Output<String> runtime;

    /**
     * @return Vault plugin runtime to use if `oci_image` is specified.
     * 
     */
    public Optional<Output<String>> runtime() {
        return Optional.ofNullable(this.runtime);
    }

    /**
     * SHA256 sum of the plugin binary.
     * 
     */
    @Import(name="sha256")
    private @Nullable Output<String> sha256;

    /**
     * @return SHA256 sum of the plugin binary.
     * 
     */
    public Optional<Output<String>> sha256() {
        return Optional.ofNullable(this.sha256);
    }

    /**
     * Type of plugin; one of &#34;auth&#34;, &#34;secret&#34;, or &#34;database&#34;.
     * 
     */
    @Import(name="type")
    private @Nullable Output<String> type;

    /**
     * @return Type of plugin; one of &#34;auth&#34;, &#34;secret&#34;, or &#34;database&#34;.
     * 
     */
    public Optional<Output<String>> type() {
        return Optional.ofNullable(this.type);
    }

    /**
     * Semantic version of the plugin.
     * 
     */
    @Import(name="version")
    private @Nullable Output<String> version;

    /**
     * @return Semantic version of the plugin.
     * 
     */
    public Optional<Output<String>> version() {
        return Optional.ofNullable(this.version);
    }

    private PluginState() {}

    private PluginState(PluginState $) {
        this.args = $.args;
        this.command = $.command;
        this.envs = $.envs;
        this.name = $.name;
        this.ociImage = $.ociImage;
        this.runtime = $.runtime;
        this.sha256 = $.sha256;
        this.type = $.type;
        this.version = $.version;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(PluginState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private PluginState $;

        public Builder() {
            $ = new PluginState();
        }

        public Builder(PluginState defaults) {
            $ = new PluginState(Objects.requireNonNull(defaults));
        }

        /**
         * @param args List of additional args to pass to the plugin.
         * 
         * @return builder
         * 
         */
        public Builder args(@Nullable Output<List<String>> args) {
            $.args = args;
            return this;
        }

        /**
         * @param args List of additional args to pass to the plugin.
         * 
         * @return builder
         * 
         */
        public Builder args(List<String> args) {
            return args(Output.of(args));
        }

        /**
         * @param args List of additional args to pass to the plugin.
         * 
         * @return builder
         * 
         */
        public Builder args(String... args) {
            return args(List.of(args));
        }

        /**
         * @param command Command to execute the plugin, relative to the server&#39;s configured `plugin_directory`.
         * 
         * @return builder
         * 
         */
        public Builder command(@Nullable Output<String> command) {
            $.command = command;
            return this;
        }

        /**
         * @param command Command to execute the plugin, relative to the server&#39;s configured `plugin_directory`.
         * 
         * @return builder
         * 
         */
        public Builder command(String command) {
            return command(Output.of(command));
        }

        /**
         * @param envs List of additional environment variables to run the plugin with in KEY=VALUE form.
         * 
         * @return builder
         * 
         */
        public Builder envs(@Nullable Output<List<String>> envs) {
            $.envs = envs;
            return this;
        }

        /**
         * @param envs List of additional environment variables to run the plugin with in KEY=VALUE form.
         * 
         * @return builder
         * 
         */
        public Builder envs(List<String> envs) {
            return envs(Output.of(envs));
        }

        /**
         * @param envs List of additional environment variables to run the plugin with in KEY=VALUE form.
         * 
         * @return builder
         * 
         */
        public Builder envs(String... envs) {
            return envs(List.of(envs));
        }

        /**
         * @param name Name of the plugin.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name Name of the plugin.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param ociImage Specifies OCI image to run. If specified, setting
         * `command`, `args`, and `env` will update the container&#39;s entrypoint, args, and
         * environment variables (append-only) respectively.
         * 
         * @return builder
         * 
         */
        public Builder ociImage(@Nullable Output<String> ociImage) {
            $.ociImage = ociImage;
            return this;
        }

        /**
         * @param ociImage Specifies OCI image to run. If specified, setting
         * `command`, `args`, and `env` will update the container&#39;s entrypoint, args, and
         * environment variables (append-only) respectively.
         * 
         * @return builder
         * 
         */
        public Builder ociImage(String ociImage) {
            return ociImage(Output.of(ociImage));
        }

        /**
         * @param runtime Vault plugin runtime to use if `oci_image` is specified.
         * 
         * @return builder
         * 
         */
        public Builder runtime(@Nullable Output<String> runtime) {
            $.runtime = runtime;
            return this;
        }

        /**
         * @param runtime Vault plugin runtime to use if `oci_image` is specified.
         * 
         * @return builder
         * 
         */
        public Builder runtime(String runtime) {
            return runtime(Output.of(runtime));
        }

        /**
         * @param sha256 SHA256 sum of the plugin binary.
         * 
         * @return builder
         * 
         */
        public Builder sha256(@Nullable Output<String> sha256) {
            $.sha256 = sha256;
            return this;
        }

        /**
         * @param sha256 SHA256 sum of the plugin binary.
         * 
         * @return builder
         * 
         */
        public Builder sha256(String sha256) {
            return sha256(Output.of(sha256));
        }

        /**
         * @param type Type of plugin; one of &#34;auth&#34;, &#34;secret&#34;, or &#34;database&#34;.
         * 
         * @return builder
         * 
         */
        public Builder type(@Nullable Output<String> type) {
            $.type = type;
            return this;
        }

        /**
         * @param type Type of plugin; one of &#34;auth&#34;, &#34;secret&#34;, or &#34;database&#34;.
         * 
         * @return builder
         * 
         */
        public Builder type(String type) {
            return type(Output.of(type));
        }

        /**
         * @param version Semantic version of the plugin.
         * 
         * @return builder
         * 
         */
        public Builder version(@Nullable Output<String> version) {
            $.version = version;
            return this;
        }

        /**
         * @param version Semantic version of the plugin.
         * 
         * @return builder
         * 
         */
        public Builder version(String version) {
            return version(Output.of(version));
        }

        public PluginState build() {
            return $;
        }
    }

}
