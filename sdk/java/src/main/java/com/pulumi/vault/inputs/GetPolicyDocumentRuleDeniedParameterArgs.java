// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vault.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.List;
import java.util.Objects;


public final class GetPolicyDocumentRuleDeniedParameterArgs extends com.pulumi.resources.ResourceArgs {

    public static final GetPolicyDocumentRuleDeniedParameterArgs Empty = new GetPolicyDocumentRuleDeniedParameterArgs();

    @Import(name="key", required=true)
    private Output<String> key;

    public Output<String> key() {
        return this.key;
    }

    @Import(name="values", required=true)
    private Output<List<String>> values;

    public Output<List<String>> values() {
        return this.values;
    }

    private GetPolicyDocumentRuleDeniedParameterArgs() {}

    private GetPolicyDocumentRuleDeniedParameterArgs(GetPolicyDocumentRuleDeniedParameterArgs $) {
        this.key = $.key;
        this.values = $.values;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetPolicyDocumentRuleDeniedParameterArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetPolicyDocumentRuleDeniedParameterArgs $;

        public Builder() {
            $ = new GetPolicyDocumentRuleDeniedParameterArgs();
        }

        public Builder(GetPolicyDocumentRuleDeniedParameterArgs defaults) {
            $ = new GetPolicyDocumentRuleDeniedParameterArgs(Objects.requireNonNull(defaults));
        }

        public Builder key(Output<String> key) {
            $.key = key;
            return this;
        }

        public Builder key(String key) {
            return key(Output.of(key));
        }

        public Builder values(Output<List<String>> values) {
            $.values = values;
            return this;
        }

        public Builder values(List<String> values) {
            return values(Output.of(values));
        }

        public Builder values(String... values) {
            return values(List.of(values));
        }

        public GetPolicyDocumentRuleDeniedParameterArgs build() {
            if ($.key == null) {
                throw new MissingRequiredPropertyException("GetPolicyDocumentRuleDeniedParameterArgs", "key");
            }
            if ($.values == null) {
                throw new MissingRequiredPropertyException("GetPolicyDocumentRuleDeniedParameterArgs", "values");
            }
            return $;
        }
    }

}
