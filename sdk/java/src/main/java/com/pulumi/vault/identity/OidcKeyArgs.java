// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vault.identity;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class OidcKeyArgs extends com.pulumi.resources.ResourceArgs {

    public static final OidcKeyArgs Empty = new OidcKeyArgs();

    /**
     * Signing algorithm to use. Signing algorithm to use.
     * Allowed values are: RS256 (default), RS384, RS512, ES256, ES384, ES512, EdDSA.
     * 
     */
    @Import(name="algorithm")
    private @Nullable Output<String> algorithm;

    /**
     * @return Signing algorithm to use. Signing algorithm to use.
     * Allowed values are: RS256 (default), RS384, RS512, ES256, ES384, ES512, EdDSA.
     * 
     */
    public Optional<Output<String>> algorithm() {
        return Optional.ofNullable(this.algorithm);
    }

    /**
     * Array of role client ids allowed to use this key for signing. If empty, no roles are allowed. If &#34;*&#34;, all roles are
     * allowed.
     * 
     */
    @Import(name="allowedClientIds")
    private @Nullable Output<List<String>> allowedClientIds;

    /**
     * @return Array of role client ids allowed to use this key for signing. If empty, no roles are allowed. If &#34;*&#34;, all roles are
     * allowed.
     * 
     */
    public Optional<Output<List<String>>> allowedClientIds() {
        return Optional.ofNullable(this.allowedClientIds);
    }

    /**
     * Name of the OIDC Key to create.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return Name of the OIDC Key to create.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * The namespace to provision the resource in.
     * The value should not contain leading or trailing forward slashes.
     * The `namespace` is always relative to the provider&#39;s configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
     * *Available only for Vault Enterprise*.
     * 
     */
    @Import(name="namespace")
    private @Nullable Output<String> namespace;

    /**
     * @return The namespace to provision the resource in.
     * The value should not contain leading or trailing forward slashes.
     * The `namespace` is always relative to the provider&#39;s configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
     * *Available only for Vault Enterprise*.
     * 
     */
    public Optional<Output<String>> namespace() {
        return Optional.ofNullable(this.namespace);
    }

    /**
     * How often to generate a new signing key in number of seconds
     * 
     */
    @Import(name="rotationPeriod")
    private @Nullable Output<Integer> rotationPeriod;

    /**
     * @return How often to generate a new signing key in number of seconds
     * 
     */
    public Optional<Output<Integer>> rotationPeriod() {
        return Optional.ofNullable(this.rotationPeriod);
    }

    /**
     * &#34;Controls how long the public portion of a signing key will be
     * available for verification after being rotated in seconds.
     * 
     */
    @Import(name="verificationTtl")
    private @Nullable Output<Integer> verificationTtl;

    /**
     * @return &#34;Controls how long the public portion of a signing key will be
     * available for verification after being rotated in seconds.
     * 
     */
    public Optional<Output<Integer>> verificationTtl() {
        return Optional.ofNullable(this.verificationTtl);
    }

    private OidcKeyArgs() {}

    private OidcKeyArgs(OidcKeyArgs $) {
        this.algorithm = $.algorithm;
        this.allowedClientIds = $.allowedClientIds;
        this.name = $.name;
        this.namespace = $.namespace;
        this.rotationPeriod = $.rotationPeriod;
        this.verificationTtl = $.verificationTtl;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(OidcKeyArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private OidcKeyArgs $;

        public Builder() {
            $ = new OidcKeyArgs();
        }

        public Builder(OidcKeyArgs defaults) {
            $ = new OidcKeyArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param algorithm Signing algorithm to use. Signing algorithm to use.
         * Allowed values are: RS256 (default), RS384, RS512, ES256, ES384, ES512, EdDSA.
         * 
         * @return builder
         * 
         */
        public Builder algorithm(@Nullable Output<String> algorithm) {
            $.algorithm = algorithm;
            return this;
        }

        /**
         * @param algorithm Signing algorithm to use. Signing algorithm to use.
         * Allowed values are: RS256 (default), RS384, RS512, ES256, ES384, ES512, EdDSA.
         * 
         * @return builder
         * 
         */
        public Builder algorithm(String algorithm) {
            return algorithm(Output.of(algorithm));
        }

        /**
         * @param allowedClientIds Array of role client ids allowed to use this key for signing. If empty, no roles are allowed. If &#34;*&#34;, all roles are
         * allowed.
         * 
         * @return builder
         * 
         */
        public Builder allowedClientIds(@Nullable Output<List<String>> allowedClientIds) {
            $.allowedClientIds = allowedClientIds;
            return this;
        }

        /**
         * @param allowedClientIds Array of role client ids allowed to use this key for signing. If empty, no roles are allowed. If &#34;*&#34;, all roles are
         * allowed.
         * 
         * @return builder
         * 
         */
        public Builder allowedClientIds(List<String> allowedClientIds) {
            return allowedClientIds(Output.of(allowedClientIds));
        }

        /**
         * @param allowedClientIds Array of role client ids allowed to use this key for signing. If empty, no roles are allowed. If &#34;*&#34;, all roles are
         * allowed.
         * 
         * @return builder
         * 
         */
        public Builder allowedClientIds(String... allowedClientIds) {
            return allowedClientIds(List.of(allowedClientIds));
        }

        /**
         * @param name Name of the OIDC Key to create.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name Name of the OIDC Key to create.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param namespace The namespace to provision the resource in.
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
         * @param namespace The namespace to provision the resource in.
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
         * @param rotationPeriod How often to generate a new signing key in number of seconds
         * 
         * @return builder
         * 
         */
        public Builder rotationPeriod(@Nullable Output<Integer> rotationPeriod) {
            $.rotationPeriod = rotationPeriod;
            return this;
        }

        /**
         * @param rotationPeriod How often to generate a new signing key in number of seconds
         * 
         * @return builder
         * 
         */
        public Builder rotationPeriod(Integer rotationPeriod) {
            return rotationPeriod(Output.of(rotationPeriod));
        }

        /**
         * @param verificationTtl &#34;Controls how long the public portion of a signing key will be
         * available for verification after being rotated in seconds.
         * 
         * @return builder
         * 
         */
        public Builder verificationTtl(@Nullable Output<Integer> verificationTtl) {
            $.verificationTtl = verificationTtl;
            return this;
        }

        /**
         * @param verificationTtl &#34;Controls how long the public portion of a signing key will be
         * available for verification after being rotated in seconds.
         * 
         * @return builder
         * 
         */
        public Builder verificationTtl(Integer verificationTtl) {
            return verificationTtl(Output.of(verificationTtl));
        }

        public OidcKeyArgs build() {
            return $;
        }
    }

}