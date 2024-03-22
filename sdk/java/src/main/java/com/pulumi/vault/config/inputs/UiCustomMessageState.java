// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vault.config.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.vault.config.inputs.UiCustomMessageLinkArgs;
import java.lang.Boolean;
import java.lang.Object;
import java.lang.String;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class UiCustomMessageState extends com.pulumi.resources.ResourceArgs {

    public static final UiCustomMessageState Empty = new UiCustomMessageState();

    /**
     * A flag indicating whether the custom message is displayed pre-login (false) or post-login (true)
     * 
     */
    @Import(name="authenticated")
    private @Nullable Output<Boolean> authenticated;

    /**
     * @return A flag indicating whether the custom message is displayed pre-login (false) or post-login (true)
     * 
     */
    public Optional<Output<Boolean>> authenticated() {
        return Optional.ofNullable(this.authenticated);
    }

    /**
     * The ending time of the active period of the custom message. Can be omitted for non-expiring message
     * 
     */
    @Import(name="endTime")
    private @Nullable Output<String> endTime;

    /**
     * @return The ending time of the active period of the custom message. Can be omitted for non-expiring message
     * 
     */
    public Optional<Output<String>> endTime() {
        return Optional.ofNullable(this.endTime);
    }

    /**
     * A block containing a hyperlink associated with the custom message
     * 
     */
    @Import(name="link")
    private @Nullable Output<UiCustomMessageLinkArgs> link;

    /**
     * @return A block containing a hyperlink associated with the custom message
     * 
     */
    public Optional<Output<UiCustomMessageLinkArgs>> link() {
        return Optional.ofNullable(this.link);
    }

    /**
     * The base64-encoded content of the custom message
     * 
     */
    @Import(name="messageBase64")
    private @Nullable Output<String> messageBase64;

    /**
     * @return The base64-encoded content of the custom message
     * 
     */
    public Optional<Output<String>> messageBase64() {
        return Optional.ofNullable(this.messageBase64);
    }

    /**
     * Target namespace. (requires Enterprise)
     * 
     */
    @Import(name="namespace")
    private @Nullable Output<String> namespace;

    /**
     * @return Target namespace. (requires Enterprise)
     * 
     */
    public Optional<Output<String>> namespace() {
        return Optional.ofNullable(this.namespace);
    }

    /**
     * A map containing additional options for the custom message
     * 
     */
    @Import(name="options")
    private @Nullable Output<Map<String,Object>> options;

    /**
     * @return A map containing additional options for the custom message
     * 
     */
    public Optional<Output<Map<String,Object>>> options() {
        return Optional.ofNullable(this.options);
    }

    /**
     * The starting time of the active period of the custom message
     * 
     */
    @Import(name="startTime")
    private @Nullable Output<String> startTime;

    /**
     * @return The starting time of the active period of the custom message
     * 
     */
    public Optional<Output<String>> startTime() {
        return Optional.ofNullable(this.startTime);
    }

    /**
     * The title of the custom message
     * 
     */
    @Import(name="title")
    private @Nullable Output<String> title;

    /**
     * @return The title of the custom message
     * 
     */
    public Optional<Output<String>> title() {
        return Optional.ofNullable(this.title);
    }

    /**
     * The display type of custom message. Allowed values are banner and modal
     * 
     */
    @Import(name="type")
    private @Nullable Output<String> type;

    /**
     * @return The display type of custom message. Allowed values are banner and modal
     * 
     */
    public Optional<Output<String>> type() {
        return Optional.ofNullable(this.type);
    }

    private UiCustomMessageState() {}

    private UiCustomMessageState(UiCustomMessageState $) {
        this.authenticated = $.authenticated;
        this.endTime = $.endTime;
        this.link = $.link;
        this.messageBase64 = $.messageBase64;
        this.namespace = $.namespace;
        this.options = $.options;
        this.startTime = $.startTime;
        this.title = $.title;
        this.type = $.type;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(UiCustomMessageState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private UiCustomMessageState $;

        public Builder() {
            $ = new UiCustomMessageState();
        }

        public Builder(UiCustomMessageState defaults) {
            $ = new UiCustomMessageState(Objects.requireNonNull(defaults));
        }

        /**
         * @param authenticated A flag indicating whether the custom message is displayed pre-login (false) or post-login (true)
         * 
         * @return builder
         * 
         */
        public Builder authenticated(@Nullable Output<Boolean> authenticated) {
            $.authenticated = authenticated;
            return this;
        }

        /**
         * @param authenticated A flag indicating whether the custom message is displayed pre-login (false) or post-login (true)
         * 
         * @return builder
         * 
         */
        public Builder authenticated(Boolean authenticated) {
            return authenticated(Output.of(authenticated));
        }

        /**
         * @param endTime The ending time of the active period of the custom message. Can be omitted for non-expiring message
         * 
         * @return builder
         * 
         */
        public Builder endTime(@Nullable Output<String> endTime) {
            $.endTime = endTime;
            return this;
        }

        /**
         * @param endTime The ending time of the active period of the custom message. Can be omitted for non-expiring message
         * 
         * @return builder
         * 
         */
        public Builder endTime(String endTime) {
            return endTime(Output.of(endTime));
        }

        /**
         * @param link A block containing a hyperlink associated with the custom message
         * 
         * @return builder
         * 
         */
        public Builder link(@Nullable Output<UiCustomMessageLinkArgs> link) {
            $.link = link;
            return this;
        }

        /**
         * @param link A block containing a hyperlink associated with the custom message
         * 
         * @return builder
         * 
         */
        public Builder link(UiCustomMessageLinkArgs link) {
            return link(Output.of(link));
        }

        /**
         * @param messageBase64 The base64-encoded content of the custom message
         * 
         * @return builder
         * 
         */
        public Builder messageBase64(@Nullable Output<String> messageBase64) {
            $.messageBase64 = messageBase64;
            return this;
        }

        /**
         * @param messageBase64 The base64-encoded content of the custom message
         * 
         * @return builder
         * 
         */
        public Builder messageBase64(String messageBase64) {
            return messageBase64(Output.of(messageBase64));
        }

        /**
         * @param namespace Target namespace. (requires Enterprise)
         * 
         * @return builder
         * 
         */
        public Builder namespace(@Nullable Output<String> namespace) {
            $.namespace = namespace;
            return this;
        }

        /**
         * @param namespace Target namespace. (requires Enterprise)
         * 
         * @return builder
         * 
         */
        public Builder namespace(String namespace) {
            return namespace(Output.of(namespace));
        }

        /**
         * @param options A map containing additional options for the custom message
         * 
         * @return builder
         * 
         */
        public Builder options(@Nullable Output<Map<String,Object>> options) {
            $.options = options;
            return this;
        }

        /**
         * @param options A map containing additional options for the custom message
         * 
         * @return builder
         * 
         */
        public Builder options(Map<String,Object> options) {
            return options(Output.of(options));
        }

        /**
         * @param startTime The starting time of the active period of the custom message
         * 
         * @return builder
         * 
         */
        public Builder startTime(@Nullable Output<String> startTime) {
            $.startTime = startTime;
            return this;
        }

        /**
         * @param startTime The starting time of the active period of the custom message
         * 
         * @return builder
         * 
         */
        public Builder startTime(String startTime) {
            return startTime(Output.of(startTime));
        }

        /**
         * @param title The title of the custom message
         * 
         * @return builder
         * 
         */
        public Builder title(@Nullable Output<String> title) {
            $.title = title;
            return this;
        }

        /**
         * @param title The title of the custom message
         * 
         * @return builder
         * 
         */
        public Builder title(String title) {
            return title(Output.of(title));
        }

        /**
         * @param type The display type of custom message. Allowed values are banner and modal
         * 
         * @return builder
         * 
         */
        public Builder type(@Nullable Output<String> type) {
            $.type = type;
            return this;
        }

        /**
         * @param type The display type of custom message. Allowed values are banner and modal
         * 
         * @return builder
         * 
         */
        public Builder type(String type) {
            return type(Output.of(type));
        }

        public UiCustomMessageState build() {
            return $;
        }
    }

}
