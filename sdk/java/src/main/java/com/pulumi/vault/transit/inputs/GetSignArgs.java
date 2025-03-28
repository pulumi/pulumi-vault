// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vault.transit.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetSignArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetSignArgs Empty = new GetSignArgs();

    @Import(name="batchInputs")
    private @Nullable Output<List<Map<String,String>>> batchInputs;

    public Optional<Output<List<Map<String,String>>>> batchInputs() {
        return Optional.ofNullable(this.batchInputs);
    }

    /**
     * The results returned from Vault if using `batch_input`
     * 
     */
    @Import(name="batchResults")
    private @Nullable Output<List<Map<String,String>>> batchResults;

    /**
     * @return The results returned from Vault if using `batch_input`
     * 
     */
    public Optional<Output<List<Map<String,String>>>> batchResults() {
        return Optional.ofNullable(this.batchResults);
    }

    @Import(name="context")
    private @Nullable Output<String> context;

    public Optional<Output<String>> context() {
        return Optional.ofNullable(this.context);
    }

    @Import(name="hashAlgorithm")
    private @Nullable Output<String> hashAlgorithm;

    public Optional<Output<String>> hashAlgorithm() {
        return Optional.ofNullable(this.hashAlgorithm);
    }

    @Import(name="input")
    private @Nullable Output<String> input;

    public Optional<Output<String>> input() {
        return Optional.ofNullable(this.input);
    }

    @Import(name="keyVersion")
    private @Nullable Output<Integer> keyVersion;

    public Optional<Output<Integer>> keyVersion() {
        return Optional.ofNullable(this.keyVersion);
    }

    @Import(name="marshalingAlgorithm")
    private @Nullable Output<String> marshalingAlgorithm;

    public Optional<Output<String>> marshalingAlgorithm() {
        return Optional.ofNullable(this.marshalingAlgorithm);
    }

    @Import(name="name", required=true)
    private Output<String> name;

    public Output<String> name() {
        return this.name;
    }

    @Import(name="namespace")
    private @Nullable Output<String> namespace;

    public Optional<Output<String>> namespace() {
        return Optional.ofNullable(this.namespace);
    }

    @Import(name="path", required=true)
    private Output<String> path;

    public Output<String> path() {
        return this.path;
    }

    @Import(name="prehashed")
    private @Nullable Output<Boolean> prehashed;

    public Optional<Output<Boolean>> prehashed() {
        return Optional.ofNullable(this.prehashed);
    }

    @Import(name="reference")
    private @Nullable Output<String> reference;

    public Optional<Output<String>> reference() {
        return Optional.ofNullable(this.reference);
    }

    @Import(name="saltLength")
    private @Nullable Output<String> saltLength;

    public Optional<Output<String>> saltLength() {
        return Optional.ofNullable(this.saltLength);
    }

    /**
     * The signature returned from Vault if using `input`
     * 
     */
    @Import(name="signature")
    private @Nullable Output<String> signature;

    /**
     * @return The signature returned from Vault if using `input`
     * 
     */
    public Optional<Output<String>> signature() {
        return Optional.ofNullable(this.signature);
    }

    @Import(name="signatureAlgorithm")
    private @Nullable Output<String> signatureAlgorithm;

    public Optional<Output<String>> signatureAlgorithm() {
        return Optional.ofNullable(this.signatureAlgorithm);
    }

    @Import(name="signatureContext")
    private @Nullable Output<String> signatureContext;

    public Optional<Output<String>> signatureContext() {
        return Optional.ofNullable(this.signatureContext);
    }

    private GetSignArgs() {}

    private GetSignArgs(GetSignArgs $) {
        this.batchInputs = $.batchInputs;
        this.batchResults = $.batchResults;
        this.context = $.context;
        this.hashAlgorithm = $.hashAlgorithm;
        this.input = $.input;
        this.keyVersion = $.keyVersion;
        this.marshalingAlgorithm = $.marshalingAlgorithm;
        this.name = $.name;
        this.namespace = $.namespace;
        this.path = $.path;
        this.prehashed = $.prehashed;
        this.reference = $.reference;
        this.saltLength = $.saltLength;
        this.signature = $.signature;
        this.signatureAlgorithm = $.signatureAlgorithm;
        this.signatureContext = $.signatureContext;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetSignArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetSignArgs $;

        public Builder() {
            $ = new GetSignArgs();
        }

        public Builder(GetSignArgs defaults) {
            $ = new GetSignArgs(Objects.requireNonNull(defaults));
        }

        public Builder batchInputs(@Nullable Output<List<Map<String,String>>> batchInputs) {
            $.batchInputs = batchInputs;
            return this;
        }

        public Builder batchInputs(List<Map<String,String>> batchInputs) {
            return batchInputs(Output.of(batchInputs));
        }

        public Builder batchInputs(Map<String,String>... batchInputs) {
            return batchInputs(List.of(batchInputs));
        }

        /**
         * @param batchResults The results returned from Vault if using `batch_input`
         * 
         * @return builder
         * 
         */
        public Builder batchResults(@Nullable Output<List<Map<String,String>>> batchResults) {
            $.batchResults = batchResults;
            return this;
        }

        /**
         * @param batchResults The results returned from Vault if using `batch_input`
         * 
         * @return builder
         * 
         */
        public Builder batchResults(List<Map<String,String>> batchResults) {
            return batchResults(Output.of(batchResults));
        }

        /**
         * @param batchResults The results returned from Vault if using `batch_input`
         * 
         * @return builder
         * 
         */
        public Builder batchResults(Map<String,String>... batchResults) {
            return batchResults(List.of(batchResults));
        }

        public Builder context(@Nullable Output<String> context) {
            $.context = context;
            return this;
        }

        public Builder context(String context) {
            return context(Output.of(context));
        }

        public Builder hashAlgorithm(@Nullable Output<String> hashAlgorithm) {
            $.hashAlgorithm = hashAlgorithm;
            return this;
        }

        public Builder hashAlgorithm(String hashAlgorithm) {
            return hashAlgorithm(Output.of(hashAlgorithm));
        }

        public Builder input(@Nullable Output<String> input) {
            $.input = input;
            return this;
        }

        public Builder input(String input) {
            return input(Output.of(input));
        }

        public Builder keyVersion(@Nullable Output<Integer> keyVersion) {
            $.keyVersion = keyVersion;
            return this;
        }

        public Builder keyVersion(Integer keyVersion) {
            return keyVersion(Output.of(keyVersion));
        }

        public Builder marshalingAlgorithm(@Nullable Output<String> marshalingAlgorithm) {
            $.marshalingAlgorithm = marshalingAlgorithm;
            return this;
        }

        public Builder marshalingAlgorithm(String marshalingAlgorithm) {
            return marshalingAlgorithm(Output.of(marshalingAlgorithm));
        }

        public Builder name(Output<String> name) {
            $.name = name;
            return this;
        }

        public Builder name(String name) {
            return name(Output.of(name));
        }

        public Builder namespace(@Nullable Output<String> namespace) {
            $.namespace = namespace;
            return this;
        }

        public Builder namespace(String namespace) {
            return namespace(Output.of(namespace));
        }

        public Builder path(Output<String> path) {
            $.path = path;
            return this;
        }

        public Builder path(String path) {
            return path(Output.of(path));
        }

        public Builder prehashed(@Nullable Output<Boolean> prehashed) {
            $.prehashed = prehashed;
            return this;
        }

        public Builder prehashed(Boolean prehashed) {
            return prehashed(Output.of(prehashed));
        }

        public Builder reference(@Nullable Output<String> reference) {
            $.reference = reference;
            return this;
        }

        public Builder reference(String reference) {
            return reference(Output.of(reference));
        }

        public Builder saltLength(@Nullable Output<String> saltLength) {
            $.saltLength = saltLength;
            return this;
        }

        public Builder saltLength(String saltLength) {
            return saltLength(Output.of(saltLength));
        }

        /**
         * @param signature The signature returned from Vault if using `input`
         * 
         * @return builder
         * 
         */
        public Builder signature(@Nullable Output<String> signature) {
            $.signature = signature;
            return this;
        }

        /**
         * @param signature The signature returned from Vault if using `input`
         * 
         * @return builder
         * 
         */
        public Builder signature(String signature) {
            return signature(Output.of(signature));
        }

        public Builder signatureAlgorithm(@Nullable Output<String> signatureAlgorithm) {
            $.signatureAlgorithm = signatureAlgorithm;
            return this;
        }

        public Builder signatureAlgorithm(String signatureAlgorithm) {
            return signatureAlgorithm(Output.of(signatureAlgorithm));
        }

        public Builder signatureContext(@Nullable Output<String> signatureContext) {
            $.signatureContext = signatureContext;
            return this;
        }

        public Builder signatureContext(String signatureContext) {
            return signatureContext(Output.of(signatureContext));
        }

        public GetSignArgs build() {
            if ($.name == null) {
                throw new MissingRequiredPropertyException("GetSignArgs", "name");
            }
            if ($.path == null) {
                throw new MissingRequiredPropertyException("GetSignArgs", "path");
            }
            return $;
        }
    }

}
