// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vault.transit.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetDecryptArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetDecryptArgs Empty = new GetDecryptArgs();

    @Import(name="backend", required=true)
    private Output<String> backend;

    public Output<String> backend() {
        return this.backend;
    }

    @Import(name="ciphertext", required=true)
    private Output<String> ciphertext;

    public Output<String> ciphertext() {
        return this.ciphertext;
    }

    @Import(name="context")
    private @Nullable Output<String> context;

    public Optional<Output<String>> context() {
        return Optional.ofNullable(this.context);
    }

    @Import(name="key", required=true)
    private Output<String> key;

    public Output<String> key() {
        return this.key;
    }

    @Import(name="namespace")
    private @Nullable Output<String> namespace;

    public Optional<Output<String>> namespace() {
        return Optional.ofNullable(this.namespace);
    }

    private GetDecryptArgs() {}

    private GetDecryptArgs(GetDecryptArgs $) {
        this.backend = $.backend;
        this.ciphertext = $.ciphertext;
        this.context = $.context;
        this.key = $.key;
        this.namespace = $.namespace;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetDecryptArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetDecryptArgs $;

        public Builder() {
            $ = new GetDecryptArgs();
        }

        public Builder(GetDecryptArgs defaults) {
            $ = new GetDecryptArgs(Objects.requireNonNull(defaults));
        }

        public Builder backend(Output<String> backend) {
            $.backend = backend;
            return this;
        }

        public Builder backend(String backend) {
            return backend(Output.of(backend));
        }

        public Builder ciphertext(Output<String> ciphertext) {
            $.ciphertext = ciphertext;
            return this;
        }

        public Builder ciphertext(String ciphertext) {
            return ciphertext(Output.of(ciphertext));
        }

        public Builder context(@Nullable Output<String> context) {
            $.context = context;
            return this;
        }

        public Builder context(String context) {
            return context(Output.of(context));
        }

        public Builder key(Output<String> key) {
            $.key = key;
            return this;
        }

        public Builder key(String key) {
            return key(Output.of(key));
        }

        public Builder namespace(@Nullable Output<String> namespace) {
            $.namespace = namespace;
            return this;
        }

        public Builder namespace(String namespace) {
            return namespace(Output.of(namespace));
        }

        public GetDecryptArgs build() {
            if ($.backend == null) {
                throw new MissingRequiredPropertyException("GetDecryptArgs", "backend");
            }
            if ($.ciphertext == null) {
                throw new MissingRequiredPropertyException("GetDecryptArgs", "ciphertext");
            }
            if ($.key == null) {
                throw new MissingRequiredPropertyException("GetDecryptArgs", "key");
            }
            return $;
        }
    }

}
