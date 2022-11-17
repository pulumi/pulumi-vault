// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vault.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetNomadAccessTokenArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetNomadAccessTokenArgs Empty = new GetNomadAccessTokenArgs();

    /**
     * The path to the Nomad secret backend to
     * read credentials from, with no leading or trailing `/`s.
     * 
     */
    @Import(name="backend", required=true)
    private Output<String> backend;

    /**
     * @return The path to the Nomad secret backend to
     * read credentials from, with no leading or trailing `/`s.
     * 
     */
    public Output<String> backend() {
        return this.backend;
    }

    /**
     * The namespace of the target resource.
     * The value should not contain leading or trailing forward slashes.
     * The `namespace` is always relative to the provider&#39;s configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
     * *Available only for Vault Enterprise*.
     * 
     */
    @Import(name="namespace")
    private @Nullable Output<String> namespace;

    /**
     * @return The namespace of the target resource.
     * The value should not contain leading or trailing forward slashes.
     * The `namespace` is always relative to the provider&#39;s configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
     * *Available only for Vault Enterprise*.
     * 
     */
    public Optional<Output<String>> namespace() {
        return Optional.ofNullable(this.namespace);
    }

    /**
     * The name of the Nomad secret backend role to generate
     * a token for, with no leading or trailing `/`s.
     * 
     */
    @Import(name="role", required=true)
    private Output<String> role;

    /**
     * @return The name of the Nomad secret backend role to generate
     * a token for, with no leading or trailing `/`s.
     * 
     */
    public Output<String> role() {
        return this.role;
    }

    private GetNomadAccessTokenArgs() {}

    private GetNomadAccessTokenArgs(GetNomadAccessTokenArgs $) {
        this.backend = $.backend;
        this.namespace = $.namespace;
        this.role = $.role;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetNomadAccessTokenArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetNomadAccessTokenArgs $;

        public Builder() {
            $ = new GetNomadAccessTokenArgs();
        }

        public Builder(GetNomadAccessTokenArgs defaults) {
            $ = new GetNomadAccessTokenArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param backend The path to the Nomad secret backend to
         * read credentials from, with no leading or trailing `/`s.
         * 
         * @return builder
         * 
         */
        public Builder backend(Output<String> backend) {
            $.backend = backend;
            return this;
        }

        /**
         * @param backend The path to the Nomad secret backend to
         * read credentials from, with no leading or trailing `/`s.
         * 
         * @return builder
         * 
         */
        public Builder backend(String backend) {
            return backend(Output.of(backend));
        }

        /**
         * @param namespace The namespace of the target resource.
         * The value should not contain leading or trailing forward slashes.
         * The `namespace` is always relative to the provider&#39;s configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
         * *Available only for Vault Enterprise*.
         * 
         * @return builder
         * 
         */
        public Builder namespace(@Nullable Output<String> namespace) {
            $.namespace = namespace;
            return this;
        }

        /**
         * @param namespace The namespace of the target resource.
         * The value should not contain leading or trailing forward slashes.
         * The `namespace` is always relative to the provider&#39;s configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
         * *Available only for Vault Enterprise*.
         * 
         * @return builder
         * 
         */
        public Builder namespace(String namespace) {
            return namespace(Output.of(namespace));
        }

        /**
         * @param role The name of the Nomad secret backend role to generate
         * a token for, with no leading or trailing `/`s.
         * 
         * @return builder
         * 
         */
        public Builder role(Output<String> role) {
            $.role = role;
            return this;
        }

        /**
         * @param role The name of the Nomad secret backend role to generate
         * a token for, with no leading or trailing `/`s.
         * 
         * @return builder
         * 
         */
        public Builder role(String role) {
            return role(Output.of(role));
        }

        public GetNomadAccessTokenArgs build() {
            $.backend = Objects.requireNonNull($.backend, "expected parameter 'backend' to be non-null");
            $.role = Objects.requireNonNull($.role, "expected parameter 'role' to be non-null");
            return $;
        }
    }

}