// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vault.transform.inputs;

import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetDecodePlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetDecodePlainArgs Empty = new GetDecodePlainArgs();

    /**
     * Specifies a list of items to be decoded in a single batch. If this parameter is set, the top-level parameters &#39;value&#39;, &#39;transformation&#39; and &#39;tweak&#39; will be ignored. Each batch item within the list can specify these parameters instead.
     * 
     */
    @Import(name="batchInputs")
    private @Nullable List<Map<String,String>> batchInputs;

    /**
     * @return Specifies a list of items to be decoded in a single batch. If this parameter is set, the top-level parameters &#39;value&#39;, &#39;transformation&#39; and &#39;tweak&#39; will be ignored. Each batch item within the list can specify these parameters instead.
     * 
     */
    public Optional<List<Map<String,String>>> batchInputs() {
        return Optional.ofNullable(this.batchInputs);
    }

    /**
     * The result of decoding a batch.
     * 
     */
    @Import(name="batchResults")
    private @Nullable List<Map<String,String>> batchResults;

    /**
     * @return The result of decoding a batch.
     * 
     */
    public Optional<List<Map<String,String>>> batchResults() {
        return Optional.ofNullable(this.batchResults);
    }

    /**
     * The result of decoding a value.
     * 
     */
    @Import(name="decodedValue")
    private @Nullable String decodedValue;

    /**
     * @return The result of decoding a value.
     * 
     */
    public Optional<String> decodedValue() {
        return Optional.ofNullable(this.decodedValue);
    }

    /**
     * The namespace of the target resource.
     * The value should not contain leading or trailing forward slashes.
     * The `namespace` is always relative to the provider&#39;s configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
     * *Available only for Vault Enterprise*.
     * 
     */
    @Import(name="namespace")
    private @Nullable String namespace;

    /**
     * @return The namespace of the target resource.
     * The value should not contain leading or trailing forward slashes.
     * The `namespace` is always relative to the provider&#39;s configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
     * *Available only for Vault Enterprise*.
     * 
     */
    public Optional<String> namespace() {
        return Optional.ofNullable(this.namespace);
    }

    /**
     * Path to where the back-end is mounted within Vault.
     * 
     */
    @Import(name="path", required=true)
    private String path;

    /**
     * @return Path to where the back-end is mounted within Vault.
     * 
     */
    public String path() {
        return this.path;
    }

    /**
     * The name of the role.
     * 
     */
    @Import(name="roleName", required=true)
    private String roleName;

    /**
     * @return The name of the role.
     * 
     */
    public String roleName() {
        return this.roleName;
    }

    /**
     * The transformation to perform. If no value is provided and the role contains a single transformation, this value will be inferred from the role.
     * 
     */
    @Import(name="transformation")
    private @Nullable String transformation;

    /**
     * @return The transformation to perform. If no value is provided and the role contains a single transformation, this value will be inferred from the role.
     * 
     */
    public Optional<String> transformation() {
        return Optional.ofNullable(this.transformation);
    }

    /**
     * The tweak value to use. Only applicable for FPE transformations
     * 
     */
    @Import(name="tweak")
    private @Nullable String tweak;

    /**
     * @return The tweak value to use. Only applicable for FPE transformations
     * 
     */
    public Optional<String> tweak() {
        return Optional.ofNullable(this.tweak);
    }

    /**
     * The value in which to decode.
     * 
     */
    @Import(name="value")
    private @Nullable String value;

    /**
     * @return The value in which to decode.
     * 
     */
    public Optional<String> value() {
        return Optional.ofNullable(this.value);
    }

    private GetDecodePlainArgs() {}

    private GetDecodePlainArgs(GetDecodePlainArgs $) {
        this.batchInputs = $.batchInputs;
        this.batchResults = $.batchResults;
        this.decodedValue = $.decodedValue;
        this.namespace = $.namespace;
        this.path = $.path;
        this.roleName = $.roleName;
        this.transformation = $.transformation;
        this.tweak = $.tweak;
        this.value = $.value;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetDecodePlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetDecodePlainArgs $;

        public Builder() {
            $ = new GetDecodePlainArgs();
        }

        public Builder(GetDecodePlainArgs defaults) {
            $ = new GetDecodePlainArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param batchInputs Specifies a list of items to be decoded in a single batch. If this parameter is set, the top-level parameters &#39;value&#39;, &#39;transformation&#39; and &#39;tweak&#39; will be ignored. Each batch item within the list can specify these parameters instead.
         * 
         * @return builder
         * 
         */
        public Builder batchInputs(@Nullable List<Map<String,String>> batchInputs) {
            $.batchInputs = batchInputs;
            return this;
        }

        /**
         * @param batchInputs Specifies a list of items to be decoded in a single batch. If this parameter is set, the top-level parameters &#39;value&#39;, &#39;transformation&#39; and &#39;tweak&#39; will be ignored. Each batch item within the list can specify these parameters instead.
         * 
         * @return builder
         * 
         */
        public Builder batchInputs(Map<String,String>... batchInputs) {
            return batchInputs(List.of(batchInputs));
        }

        /**
         * @param batchResults The result of decoding a batch.
         * 
         * @return builder
         * 
         */
        public Builder batchResults(@Nullable List<Map<String,String>> batchResults) {
            $.batchResults = batchResults;
            return this;
        }

        /**
         * @param batchResults The result of decoding a batch.
         * 
         * @return builder
         * 
         */
        public Builder batchResults(Map<String,String>... batchResults) {
            return batchResults(List.of(batchResults));
        }

        /**
         * @param decodedValue The result of decoding a value.
         * 
         * @return builder
         * 
         */
        public Builder decodedValue(@Nullable String decodedValue) {
            $.decodedValue = decodedValue;
            return this;
        }

        /**
         * @param namespace The namespace of the target resource.
         * The value should not contain leading or trailing forward slashes.
         * The `namespace` is always relative to the provider&#39;s configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
         * *Available only for Vault Enterprise*.
         * 
         * @return builder
         * 
         */
        public Builder namespace(@Nullable String namespace) {
            $.namespace = namespace;
            return this;
        }

        /**
         * @param path Path to where the back-end is mounted within Vault.
         * 
         * @return builder
         * 
         */
        public Builder path(String path) {
            $.path = path;
            return this;
        }

        /**
         * @param roleName The name of the role.
         * 
         * @return builder
         * 
         */
        public Builder roleName(String roleName) {
            $.roleName = roleName;
            return this;
        }

        /**
         * @param transformation The transformation to perform. If no value is provided and the role contains a single transformation, this value will be inferred from the role.
         * 
         * @return builder
         * 
         */
        public Builder transformation(@Nullable String transformation) {
            $.transformation = transformation;
            return this;
        }

        /**
         * @param tweak The tweak value to use. Only applicable for FPE transformations
         * 
         * @return builder
         * 
         */
        public Builder tweak(@Nullable String tweak) {
            $.tweak = tweak;
            return this;
        }

        /**
         * @param value The value in which to decode.
         * 
         * @return builder
         * 
         */
        public Builder value(@Nullable String value) {
            $.value = value;
            return this;
        }

        public GetDecodePlainArgs build() {
            if ($.path == null) {
                throw new MissingRequiredPropertyException("GetDecodePlainArgs", "path");
            }
            if ($.roleName == null) {
                throw new MissingRequiredPropertyException("GetDecodePlainArgs", "roleName");
            }
            return $;
        }
    }

}
