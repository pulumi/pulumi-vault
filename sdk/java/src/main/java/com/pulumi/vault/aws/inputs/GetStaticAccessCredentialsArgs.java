// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vault.aws.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetStaticAccessCredentialsArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetStaticAccessCredentialsArgs Empty = new GetStaticAccessCredentialsArgs();

    @Import(name="backend", required=true)
    private Output<String> backend;

    public Output<String> backend() {
        return this.backend;
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

    private GetStaticAccessCredentialsArgs() {}

    private GetStaticAccessCredentialsArgs(GetStaticAccessCredentialsArgs $) {
        this.backend = $.backend;
        this.name = $.name;
        this.namespace = $.namespace;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetStaticAccessCredentialsArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetStaticAccessCredentialsArgs $;

        public Builder() {
            $ = new GetStaticAccessCredentialsArgs();
        }

        public Builder(GetStaticAccessCredentialsArgs defaults) {
            $ = new GetStaticAccessCredentialsArgs(Objects.requireNonNull(defaults));
        }

        public Builder backend(Output<String> backend) {
            $.backend = backend;
            return this;
        }

        public Builder backend(String backend) {
            return backend(Output.of(backend));
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

        public GetStaticAccessCredentialsArgs build() {
            if ($.backend == null) {
                throw new MissingRequiredPropertyException("GetStaticAccessCredentialsArgs", "backend");
            }
            if ($.name == null) {
                throw new MissingRequiredPropertyException("GetStaticAccessCredentialsArgs", "name");
            }
            return $;
        }
    }

}
