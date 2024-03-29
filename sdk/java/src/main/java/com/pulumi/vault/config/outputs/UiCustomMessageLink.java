// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vault.config.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class UiCustomMessageLink {
    /**
     * @return The URL of the hyperlink
     * 
     */
    private String href;
    /**
     * @return The title of the hyperlink
     * 
     */
    private String title;

    private UiCustomMessageLink() {}
    /**
     * @return The URL of the hyperlink
     * 
     */
    public String href() {
        return this.href;
    }
    /**
     * @return The title of the hyperlink
     * 
     */
    public String title() {
        return this.title;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(UiCustomMessageLink defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String href;
        private String title;
        public Builder() {}
        public Builder(UiCustomMessageLink defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.href = defaults.href;
    	      this.title = defaults.title;
        }

        @CustomType.Setter
        public Builder href(String href) {
            if (href == null) {
              throw new MissingRequiredPropertyException("UiCustomMessageLink", "href");
            }
            this.href = href;
            return this;
        }
        @CustomType.Setter
        public Builder title(String title) {
            if (title == null) {
              throw new MissingRequiredPropertyException("UiCustomMessageLink", "title");
            }
            this.title = title;
            return this;
        }
        public UiCustomMessageLink build() {
            final var _resultValue = new UiCustomMessageLink();
            _resultValue.href = href;
            _resultValue.title = title;
            return _resultValue;
        }
    }
}
