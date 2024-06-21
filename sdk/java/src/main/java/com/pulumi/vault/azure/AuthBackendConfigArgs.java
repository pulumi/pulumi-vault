// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vault.azure;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class AuthBackendConfigArgs extends com.pulumi.resources.ResourceArgs {

    public static final AuthBackendConfigArgs Empty = new AuthBackendConfigArgs();

    /**
     * The path the Azure auth backend being configured was
     * mounted at.  Defaults to `azure`.
     * 
     */
    @Import(name="backend")
    private @Nullable Output<String> backend;

    /**
     * @return The path the Azure auth backend being configured was
     * mounted at.  Defaults to `azure`.
     * 
     */
    public Optional<Output<String>> backend() {
        return Optional.ofNullable(this.backend);
    }

    /**
     * The client id for credentials to query the Azure APIs.
     * Currently read permissions to query compute resources are required.
     * 
     */
    @Import(name="clientId")
    private @Nullable Output<String> clientId;

    /**
     * @return The client id for credentials to query the Azure APIs.
     * Currently read permissions to query compute resources are required.
     * 
     */
    public Optional<Output<String>> clientId() {
        return Optional.ofNullable(this.clientId);
    }

    /**
     * The client secret for credentials to query the
     * Azure APIs.
     * 
     */
    @Import(name="clientSecret")
    private @Nullable Output<String> clientSecret;

    /**
     * @return The client secret for credentials to query the
     * Azure APIs.
     * 
     */
    public Optional<Output<String>> clientSecret() {
        return Optional.ofNullable(this.clientSecret);
    }

    /**
     * The Azure cloud environment. Valid values:
     * AzurePublicCloud, AzureUSGovernmentCloud, AzureChinaCloud,
     * AzureGermanCloud.  Defaults to `AzurePublicCloud`.
     * 
     */
    @Import(name="environment")
    private @Nullable Output<String> environment;

    /**
     * @return The Azure cloud environment. Valid values:
     * AzurePublicCloud, AzureUSGovernmentCloud, AzureChinaCloud,
     * AzureGermanCloud.  Defaults to `AzurePublicCloud`.
     * 
     */
    public Optional<Output<String>> environment() {
        return Optional.ofNullable(this.environment);
    }

    /**
     * The audience claim value for plugin identity tokens. Requires Vault 1.17+.
     * *Available only for Vault Enterprise*
     * 
     */
    @Import(name="identityTokenAudience")
    private @Nullable Output<String> identityTokenAudience;

    /**
     * @return The audience claim value for plugin identity tokens. Requires Vault 1.17+.
     * *Available only for Vault Enterprise*
     * 
     */
    public Optional<Output<String>> identityTokenAudience() {
        return Optional.ofNullable(this.identityTokenAudience);
    }

    /**
     * The TTL of generated identity tokens in seconds.
     * 
     */
    @Import(name="identityTokenTtl")
    private @Nullable Output<Integer> identityTokenTtl;

    /**
     * @return The TTL of generated identity tokens in seconds.
     * 
     */
    public Optional<Output<Integer>> identityTokenTtl() {
        return Optional.ofNullable(this.identityTokenTtl);
    }

    /**
     * The namespace to provision the resource in.
     * The value should not contain leading or trailing forward slashes.
     * The `namespace` is always relative to the provider&#39;s configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
     * *Available only for Vault Enterprise*.
     * 
     */
    @Import(name="namespace")
    private @Nullable Output<String> namespace;

    /**
     * @return The namespace to provision the resource in.
     * The value should not contain leading or trailing forward slashes.
     * The `namespace` is always relative to the provider&#39;s configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
     * *Available only for Vault Enterprise*.
     * 
     */
    public Optional<Output<String>> namespace() {
        return Optional.ofNullable(this.namespace);
    }

    /**
     * The configured URL for the application registered in
     * Azure Active Directory.
     * 
     */
    @Import(name="resource", required=true)
    private Output<String> resource;

    /**
     * @return The configured URL for the application registered in
     * Azure Active Directory.
     * 
     */
    public Output<String> resource() {
        return this.resource;
    }

    /**
     * The tenant id for the Azure Active Directory
     * organization.
     * 
     */
    @Import(name="tenantId", required=true)
    private Output<String> tenantId;

    /**
     * @return The tenant id for the Azure Active Directory
     * organization.
     * 
     */
    public Output<String> tenantId() {
        return this.tenantId;
    }

    private AuthBackendConfigArgs() {}

    private AuthBackendConfigArgs(AuthBackendConfigArgs $) {
        this.backend = $.backend;
        this.clientId = $.clientId;
        this.clientSecret = $.clientSecret;
        this.environment = $.environment;
        this.identityTokenAudience = $.identityTokenAudience;
        this.identityTokenTtl = $.identityTokenTtl;
        this.namespace = $.namespace;
        this.resource = $.resource;
        this.tenantId = $.tenantId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(AuthBackendConfigArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private AuthBackendConfigArgs $;

        public Builder() {
            $ = new AuthBackendConfigArgs();
        }

        public Builder(AuthBackendConfigArgs defaults) {
            $ = new AuthBackendConfigArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param backend The path the Azure auth backend being configured was
         * mounted at.  Defaults to `azure`.
         * 
         * @return builder
         * 
         */
        public Builder backend(@Nullable Output<String> backend) {
            $.backend = backend;
            return this;
        }

        /**
         * @param backend The path the Azure auth backend being configured was
         * mounted at.  Defaults to `azure`.
         * 
         * @return builder
         * 
         */
        public Builder backend(String backend) {
            return backend(Output.of(backend));
        }

        /**
         * @param clientId The client id for credentials to query the Azure APIs.
         * Currently read permissions to query compute resources are required.
         * 
         * @return builder
         * 
         */
        public Builder clientId(@Nullable Output<String> clientId) {
            $.clientId = clientId;
            return this;
        }

        /**
         * @param clientId The client id for credentials to query the Azure APIs.
         * Currently read permissions to query compute resources are required.
         * 
         * @return builder
         * 
         */
        public Builder clientId(String clientId) {
            return clientId(Output.of(clientId));
        }

        /**
         * @param clientSecret The client secret for credentials to query the
         * Azure APIs.
         * 
         * @return builder
         * 
         */
        public Builder clientSecret(@Nullable Output<String> clientSecret) {
            $.clientSecret = clientSecret;
            return this;
        }

        /**
         * @param clientSecret The client secret for credentials to query the
         * Azure APIs.
         * 
         * @return builder
         * 
         */
        public Builder clientSecret(String clientSecret) {
            return clientSecret(Output.of(clientSecret));
        }

        /**
         * @param environment The Azure cloud environment. Valid values:
         * AzurePublicCloud, AzureUSGovernmentCloud, AzureChinaCloud,
         * AzureGermanCloud.  Defaults to `AzurePublicCloud`.
         * 
         * @return builder
         * 
         */
        public Builder environment(@Nullable Output<String> environment) {
            $.environment = environment;
            return this;
        }

        /**
         * @param environment The Azure cloud environment. Valid values:
         * AzurePublicCloud, AzureUSGovernmentCloud, AzureChinaCloud,
         * AzureGermanCloud.  Defaults to `AzurePublicCloud`.
         * 
         * @return builder
         * 
         */
        public Builder environment(String environment) {
            return environment(Output.of(environment));
        }

        /**
         * @param identityTokenAudience The audience claim value for plugin identity tokens. Requires Vault 1.17+.
         * *Available only for Vault Enterprise*
         * 
         * @return builder
         * 
         */
        public Builder identityTokenAudience(@Nullable Output<String> identityTokenAudience) {
            $.identityTokenAudience = identityTokenAudience;
            return this;
        }

        /**
         * @param identityTokenAudience The audience claim value for plugin identity tokens. Requires Vault 1.17+.
         * *Available only for Vault Enterprise*
         * 
         * @return builder
         * 
         */
        public Builder identityTokenAudience(String identityTokenAudience) {
            return identityTokenAudience(Output.of(identityTokenAudience));
        }

        /**
         * @param identityTokenTtl The TTL of generated identity tokens in seconds.
         * 
         * @return builder
         * 
         */
        public Builder identityTokenTtl(@Nullable Output<Integer> identityTokenTtl) {
            $.identityTokenTtl = identityTokenTtl;
            return this;
        }

        /**
         * @param identityTokenTtl The TTL of generated identity tokens in seconds.
         * 
         * @return builder
         * 
         */
        public Builder identityTokenTtl(Integer identityTokenTtl) {
            return identityTokenTtl(Output.of(identityTokenTtl));
        }

        /**
         * @param namespace The namespace to provision the resource in.
         * The value should not contain leading or trailing forward slashes.
         * The `namespace` is always relative to the provider&#39;s configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
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
         * The `namespace` is always relative to the provider&#39;s configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
         * *Available only for Vault Enterprise*.
         * 
         * @return builder
         * 
         */
        public Builder namespace(String namespace) {
            return namespace(Output.of(namespace));
        }

        /**
         * @param resource The configured URL for the application registered in
         * Azure Active Directory.
         * 
         * @return builder
         * 
         */
        public Builder resource(Output<String> resource) {
            $.resource = resource;
            return this;
        }

        /**
         * @param resource The configured URL for the application registered in
         * Azure Active Directory.
         * 
         * @return builder
         * 
         */
        public Builder resource(String resource) {
            return resource(Output.of(resource));
        }

        /**
         * @param tenantId The tenant id for the Azure Active Directory
         * organization.
         * 
         * @return builder
         * 
         */
        public Builder tenantId(Output<String> tenantId) {
            $.tenantId = tenantId;
            return this;
        }

        /**
         * @param tenantId The tenant id for the Azure Active Directory
         * organization.
         * 
         * @return builder
         * 
         */
        public Builder tenantId(String tenantId) {
            return tenantId(Output.of(tenantId));
        }

        public AuthBackendConfigArgs build() {
            if ($.resource == null) {
                throw new MissingRequiredPropertyException("AuthBackendConfigArgs", "resource");
            }
            if ($.tenantId == null) {
                throw new MissingRequiredPropertyException("AuthBackendConfigArgs", "tenantId");
            }
            return $;
        }
    }

}
