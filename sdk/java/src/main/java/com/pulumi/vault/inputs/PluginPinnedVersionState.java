// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vault.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class PluginPinnedVersionState extends com.pulumi.resources.ResourceArgs {

    public static final PluginPinnedVersionState Empty = new PluginPinnedVersionState();

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
     * Semantic version of the plugin to pin.
     * 
     */
    @Import(name="version")
    private @Nullable Output<String> version;

    /**
     * @return Semantic version of the plugin to pin.
     * 
     */
    public Optional<Output<String>> version() {
        return Optional.ofNullable(this.version);
    }

    private PluginPinnedVersionState() {}

    private PluginPinnedVersionState(PluginPinnedVersionState $) {
        this.name = $.name;
        this.type = $.type;
        this.version = $.version;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(PluginPinnedVersionState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private PluginPinnedVersionState $;

        public Builder() {
            $ = new PluginPinnedVersionState();
        }

        public Builder(PluginPinnedVersionState defaults) {
            $ = new PluginPinnedVersionState(Objects.requireNonNull(defaults));
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
         * @param version Semantic version of the plugin to pin.
         * 
         * @return builder
         * 
         */
        public Builder version(@Nullable Output<String> version) {
            $.version = version;
            return this;
        }

        /**
         * @param version Semantic version of the plugin to pin.
         * 
         * @return builder
         * 
         */
        public Builder version(String version) {
            return version(Output.of(version));
        }

        public PluginPinnedVersionState build() {
            return $;
        }
    }

}
