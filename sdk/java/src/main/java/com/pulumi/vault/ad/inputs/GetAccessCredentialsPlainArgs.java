// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vault.ad.inputs;

import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetAccessCredentialsPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetAccessCredentialsPlainArgs Empty = new GetAccessCredentialsPlainArgs();

    /**
     * The path to the AD secret backend to
     * read credentials from, with no leading or trailing `/`s.
     * 
     */
    @Import(name="backend", required=true)
    private String backend;

    /**
     * @return The path to the AD secret backend to
     * read credentials from, with no leading or trailing `/`s.
     * 
     */
    public String backend() {
        return this.backend;
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
     * The name of the AD secret backend role to read
     * credentials from, with no leading or trailing `/`s.
     * 
     */
    @Import(name="role", required=true)
    private String role;

    /**
     * @return The name of the AD secret backend role to read
     * credentials from, with no leading or trailing `/`s.
     * 
     */
    public String role() {
        return this.role;
    }

    private GetAccessCredentialsPlainArgs() {}

    private GetAccessCredentialsPlainArgs(GetAccessCredentialsPlainArgs $) {
        this.backend = $.backend;
        this.namespace = $.namespace;
        this.role = $.role;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetAccessCredentialsPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetAccessCredentialsPlainArgs $;

        public Builder() {
            $ = new GetAccessCredentialsPlainArgs();
        }

        public Builder(GetAccessCredentialsPlainArgs defaults) {
            $ = new GetAccessCredentialsPlainArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param backend The path to the AD secret backend to
         * read credentials from, with no leading or trailing `/`s.
         * 
         * @return builder
         * 
         */
        public Builder backend(String backend) {
            $.backend = backend;
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
         * @param role The name of the AD secret backend role to read
         * credentials from, with no leading or trailing `/`s.
         * 
         * @return builder
         * 
         */
        public Builder role(String role) {
            $.role = role;
            return this;
        }

        public GetAccessCredentialsPlainArgs build() {
            if ($.backend == null) {
                throw new MissingRequiredPropertyException("GetAccessCredentialsPlainArgs", "backend");
            }
            if ($.role == null) {
                throw new MissingRequiredPropertyException("GetAccessCredentialsPlainArgs", "role");
            }
            return $;
        }
    }

}
