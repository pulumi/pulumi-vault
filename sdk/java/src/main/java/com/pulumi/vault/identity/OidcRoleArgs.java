// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vault.identity;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class OidcRoleArgs extends com.pulumi.resources.ResourceArgs {

    public static final OidcRoleArgs Empty = new OidcRoleArgs();

    /**
     * The value that will be included in the `aud` field of all the OIDC identity
     * tokens issued by this role
     * 
     */
    @Import(name="clientId")
    private @Nullable Output<String> clientId;

    /**
     * @return The value that will be included in the `aud` field of all the OIDC identity
     * tokens issued by this role
     * 
     */
    public Optional<Output<String>> clientId() {
        return Optional.ofNullable(this.clientId);
    }

    /**
     * A configured named key, the key must already exist
     * before tokens can be issued.
     * 
     */
    @Import(name="key", required=true)
    private Output<String> key;

    /**
     * @return A configured named key, the key must already exist
     * before tokens can be issued.
     * 
     */
    public Output<String> key() {
        return this.key;
    }

    /**
     * Name of the OIDC Role to create.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return Name of the OIDC Role to create.
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
     * The template string to use for generating tokens. This may be in
     * string-ified JSON or base64 format. See the
     * [documentation](https://www.vaultproject.io/docs/secrets/identity/index.html#token-contents-and-templates)
     * for the template format.
     * 
     */
    @Import(name="template")
    private @Nullable Output<String> template;

    /**
     * @return The template string to use for generating tokens. This may be in
     * string-ified JSON or base64 format. See the
     * [documentation](https://www.vaultproject.io/docs/secrets/identity/index.html#token-contents-and-templates)
     * for the template format.
     * 
     */
    public Optional<Output<String>> template() {
        return Optional.ofNullable(this.template);
    }

    /**
     * TTL of the tokens generated against the role in number of seconds.
     * 
     */
    @Import(name="ttl")
    private @Nullable Output<Integer> ttl;

    /**
     * @return TTL of the tokens generated against the role in number of seconds.
     * 
     */
    public Optional<Output<Integer>> ttl() {
        return Optional.ofNullable(this.ttl);
    }

    private OidcRoleArgs() {}

    private OidcRoleArgs(OidcRoleArgs $) {
        this.clientId = $.clientId;
        this.key = $.key;
        this.name = $.name;
        this.namespace = $.namespace;
        this.template = $.template;
        this.ttl = $.ttl;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(OidcRoleArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private OidcRoleArgs $;

        public Builder() {
            $ = new OidcRoleArgs();
        }

        public Builder(OidcRoleArgs defaults) {
            $ = new OidcRoleArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param clientId The value that will be included in the `aud` field of all the OIDC identity
         * tokens issued by this role
         * 
         * @return builder
         * 
         */
        public Builder clientId(@Nullable Output<String> clientId) {
            $.clientId = clientId;
            return this;
        }

        /**
         * @param clientId The value that will be included in the `aud` field of all the OIDC identity
         * tokens issued by this role
         * 
         * @return builder
         * 
         */
        public Builder clientId(String clientId) {
            return clientId(Output.of(clientId));
        }

        /**
         * @param key A configured named key, the key must already exist
         * before tokens can be issued.
         * 
         * @return builder
         * 
         */
        public Builder key(Output<String> key) {
            $.key = key;
            return this;
        }

        /**
         * @param key A configured named key, the key must already exist
         * before tokens can be issued.
         * 
         * @return builder
         * 
         */
        public Builder key(String key) {
            return key(Output.of(key));
        }

        /**
         * @param name Name of the OIDC Role to create.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name Name of the OIDC Role to create.
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
         * @param template The template string to use for generating tokens. This may be in
         * string-ified JSON or base64 format. See the
         * [documentation](https://www.vaultproject.io/docs/secrets/identity/index.html#token-contents-and-templates)
         * for the template format.
         * 
         * @return builder
         * 
         */
        public Builder template(@Nullable Output<String> template) {
            $.template = template;
            return this;
        }

        /**
         * @param template The template string to use for generating tokens. This may be in
         * string-ified JSON or base64 format. See the
         * [documentation](https://www.vaultproject.io/docs/secrets/identity/index.html#token-contents-and-templates)
         * for the template format.
         * 
         * @return builder
         * 
         */
        public Builder template(String template) {
            return template(Output.of(template));
        }

        /**
         * @param ttl TTL of the tokens generated against the role in number of seconds.
         * 
         * @return builder
         * 
         */
        public Builder ttl(@Nullable Output<Integer> ttl) {
            $.ttl = ttl;
            return this;
        }

        /**
         * @param ttl TTL of the tokens generated against the role in number of seconds.
         * 
         * @return builder
         * 
         */
        public Builder ttl(Integer ttl) {
            return ttl(Output.of(ttl));
        }

        public OidcRoleArgs build() {
            if ($.key == null) {
                throw new MissingRequiredPropertyException("OidcRoleArgs", "key");
            }
            return $;
        }
    }

}
