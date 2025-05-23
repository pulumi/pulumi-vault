// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vault.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.List;
import java.util.Objects;


public final class GetPolicyDocumentRuleAllowedParameterArgs extends com.pulumi.resources.ResourceArgs {

    public static final GetPolicyDocumentRuleAllowedParameterArgs Empty = new GetPolicyDocumentRuleAllowedParameterArgs();

    /**
     * Name of permitted key.
     * 
     */
    @Import(name="key", required=true)
    private Output<String> key;

    /**
     * @return Name of permitted key.
     * 
     */
    public Output<String> key() {
        return this.key;
    }

    /**
     * A list of values what are permitted by policy rule.
     * 
     */
    @Import(name="values", required=true)
    private Output<List<String>> values;

    /**
     * @return A list of values what are permitted by policy rule.
     * 
     */
    public Output<List<String>> values() {
        return this.values;
    }

    private GetPolicyDocumentRuleAllowedParameterArgs() {}

    private GetPolicyDocumentRuleAllowedParameterArgs(GetPolicyDocumentRuleAllowedParameterArgs $) {
        this.key = $.key;
        this.values = $.values;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetPolicyDocumentRuleAllowedParameterArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetPolicyDocumentRuleAllowedParameterArgs $;

        public Builder() {
            $ = new GetPolicyDocumentRuleAllowedParameterArgs();
        }

        public Builder(GetPolicyDocumentRuleAllowedParameterArgs defaults) {
            $ = new GetPolicyDocumentRuleAllowedParameterArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param key Name of permitted key.
         * 
         * @return builder
         * 
         */
        public Builder key(Output<String> key) {
            $.key = key;
            return this;
        }

        /**
         * @param key Name of permitted key.
         * 
         * @return builder
         * 
         */
        public Builder key(String key) {
            return key(Output.of(key));
        }

        /**
         * @param values A list of values what are permitted by policy rule.
         * 
         * @return builder
         * 
         */
        public Builder values(Output<List<String>> values) {
            $.values = values;
            return this;
        }

        /**
         * @param values A list of values what are permitted by policy rule.
         * 
         * @return builder
         * 
         */
        public Builder values(List<String> values) {
            return values(Output.of(values));
        }

        /**
         * @param values A list of values what are permitted by policy rule.
         * 
         * @return builder
         * 
         */
        public Builder values(String... values) {
            return values(List.of(values));
        }

        public GetPolicyDocumentRuleAllowedParameterArgs build() {
            if ($.key == null) {
                throw new MissingRequiredPropertyException("GetPolicyDocumentRuleAllowedParameterArgs", "key");
            }
            if ($.values == null) {
                throw new MissingRequiredPropertyException("GetPolicyDocumentRuleAllowedParameterArgs", "values");
            }
            return $;
        }
    }

}
