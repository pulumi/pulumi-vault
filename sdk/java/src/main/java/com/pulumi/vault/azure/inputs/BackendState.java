// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vault.azure.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class BackendState extends com.pulumi.resources.ResourceArgs {

    public static final BackendState Empty = new BackendState();

    /**
     * The OAuth2 client id to connect to Azure.
     * 
     */
    @Import(name="clientId")
    private @Nullable Output<String> clientId;

    /**
     * @return The OAuth2 client id to connect to Azure.
     * 
     */
    public Optional<Output<String>> clientId() {
        return Optional.ofNullable(this.clientId);
    }

    /**
     * The OAuth2 client secret to connect to Azure.
     * 
     */
    @Import(name="clientSecret")
    private @Nullable Output<String> clientSecret;

    /**
     * @return The OAuth2 client secret to connect to Azure.
     * 
     */
    public Optional<Output<String>> clientSecret() {
        return Optional.ofNullable(this.clientSecret);
    }

    /**
     * Human-friendly description of the mount for the backend.
     * 
     */
    @Import(name="description")
    private @Nullable Output<String> description;

    /**
     * @return Human-friendly description of the mount for the backend.
     * 
     */
    public Optional<Output<String>> description() {
        return Optional.ofNullable(this.description);
    }

    /**
     * If set, opts out of mount migration on path updates.
     * See here for more info on [Mount Migration](https://www.vaultproject.io/docs/concepts/mount-migration)
     * 
     */
    @Import(name="disableRemount")
    private @Nullable Output<Boolean> disableRemount;

    /**
     * @return If set, opts out of mount migration on path updates.
     * See here for more info on [Mount Migration](https://www.vaultproject.io/docs/concepts/mount-migration)
     * 
     */
    public Optional<Output<Boolean>> disableRemount() {
        return Optional.ofNullable(this.disableRemount);
    }

    /**
     * The Azure environment.
     * 
     */
    @Import(name="environment")
    private @Nullable Output<String> environment;

    /**
     * @return The Azure environment.
     * 
     */
    public Optional<Output<String>> environment() {
        return Optional.ofNullable(this.environment);
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
     * The unique path this backend should be mounted at. Defaults to `azure`.
     * 
     */
    @Import(name="path")
    private @Nullable Output<String> path;

    /**
     * @return The unique path this backend should be mounted at. Defaults to `azure`.
     * 
     */
    public Optional<Output<String>> path() {
        return Optional.ofNullable(this.path);
    }

    /**
     * The subscription id for the Azure Active Directory.
     * 
     */
    @Import(name="subscriptionId")
    private @Nullable Output<String> subscriptionId;

    /**
     * @return The subscription id for the Azure Active Directory.
     * 
     */
    public Optional<Output<String>> subscriptionId() {
        return Optional.ofNullable(this.subscriptionId);
    }

    /**
     * The tenant id for the Azure Active Directory.
     * 
     */
    @Import(name="tenantId")
    private @Nullable Output<String> tenantId;

    /**
     * @return The tenant id for the Azure Active Directory.
     * 
     */
    public Optional<Output<String>> tenantId() {
        return Optional.ofNullable(this.tenantId);
    }

    /**
     * Use the Microsoft Graph API. Should be set to true on vault-1.10+
     * 
     */
    @Import(name="useMicrosoftGraphApi")
    private @Nullable Output<Boolean> useMicrosoftGraphApi;

    /**
     * @return Use the Microsoft Graph API. Should be set to true on vault-1.10+
     * 
     */
    public Optional<Output<Boolean>> useMicrosoftGraphApi() {
        return Optional.ofNullable(this.useMicrosoftGraphApi);
    }

    private BackendState() {}

    private BackendState(BackendState $) {
        this.clientId = $.clientId;
        this.clientSecret = $.clientSecret;
        this.description = $.description;
        this.disableRemount = $.disableRemount;
        this.environment = $.environment;
        this.namespace = $.namespace;
        this.path = $.path;
        this.subscriptionId = $.subscriptionId;
        this.tenantId = $.tenantId;
        this.useMicrosoftGraphApi = $.useMicrosoftGraphApi;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(BackendState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private BackendState $;

        public Builder() {
            $ = new BackendState();
        }

        public Builder(BackendState defaults) {
            $ = new BackendState(Objects.requireNonNull(defaults));
        }

        /**
         * @param clientId The OAuth2 client id to connect to Azure.
         * 
         * @return builder
         * 
         */
        public Builder clientId(@Nullable Output<String> clientId) {
            $.clientId = clientId;
            return this;
        }

        /**
         * @param clientId The OAuth2 client id to connect to Azure.
         * 
         * @return builder
         * 
         */
        public Builder clientId(String clientId) {
            return clientId(Output.of(clientId));
        }

        /**
         * @param clientSecret The OAuth2 client secret to connect to Azure.
         * 
         * @return builder
         * 
         */
        public Builder clientSecret(@Nullable Output<String> clientSecret) {
            $.clientSecret = clientSecret;
            return this;
        }

        /**
         * @param clientSecret The OAuth2 client secret to connect to Azure.
         * 
         * @return builder
         * 
         */
        public Builder clientSecret(String clientSecret) {
            return clientSecret(Output.of(clientSecret));
        }

        /**
         * @param description Human-friendly description of the mount for the backend.
         * 
         * @return builder
         * 
         */
        public Builder description(@Nullable Output<String> description) {
            $.description = description;
            return this;
        }

        /**
         * @param description Human-friendly description of the mount for the backend.
         * 
         * @return builder
         * 
         */
        public Builder description(String description) {
            return description(Output.of(description));
        }

        /**
         * @param disableRemount If set, opts out of mount migration on path updates.
         * See here for more info on [Mount Migration](https://www.vaultproject.io/docs/concepts/mount-migration)
         * 
         * @return builder
         * 
         */
        public Builder disableRemount(@Nullable Output<Boolean> disableRemount) {
            $.disableRemount = disableRemount;
            return this;
        }

        /**
         * @param disableRemount If set, opts out of mount migration on path updates.
         * See here for more info on [Mount Migration](https://www.vaultproject.io/docs/concepts/mount-migration)
         * 
         * @return builder
         * 
         */
        public Builder disableRemount(Boolean disableRemount) {
            return disableRemount(Output.of(disableRemount));
        }

        /**
         * @param environment The Azure environment.
         * 
         * @return builder
         * 
         */
        public Builder environment(@Nullable Output<String> environment) {
            $.environment = environment;
            return this;
        }

        /**
         * @param environment The Azure environment.
         * 
         * @return builder
         * 
         */
        public Builder environment(String environment) {
            return environment(Output.of(environment));
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
         * @param path The unique path this backend should be mounted at. Defaults to `azure`.
         * 
         * @return builder
         * 
         */
        public Builder path(@Nullable Output<String> path) {
            $.path = path;
            return this;
        }

        /**
         * @param path The unique path this backend should be mounted at. Defaults to `azure`.
         * 
         * @return builder
         * 
         */
        public Builder path(String path) {
            return path(Output.of(path));
        }

        /**
         * @param subscriptionId The subscription id for the Azure Active Directory.
         * 
         * @return builder
         * 
         */
        public Builder subscriptionId(@Nullable Output<String> subscriptionId) {
            $.subscriptionId = subscriptionId;
            return this;
        }

        /**
         * @param subscriptionId The subscription id for the Azure Active Directory.
         * 
         * @return builder
         * 
         */
        public Builder subscriptionId(String subscriptionId) {
            return subscriptionId(Output.of(subscriptionId));
        }

        /**
         * @param tenantId The tenant id for the Azure Active Directory.
         * 
         * @return builder
         * 
         */
        public Builder tenantId(@Nullable Output<String> tenantId) {
            $.tenantId = tenantId;
            return this;
        }

        /**
         * @param tenantId The tenant id for the Azure Active Directory.
         * 
         * @return builder
         * 
         */
        public Builder tenantId(String tenantId) {
            return tenantId(Output.of(tenantId));
        }

        /**
         * @param useMicrosoftGraphApi Use the Microsoft Graph API. Should be set to true on vault-1.10+
         * 
         * @return builder
         * 
         */
        public Builder useMicrosoftGraphApi(@Nullable Output<Boolean> useMicrosoftGraphApi) {
            $.useMicrosoftGraphApi = useMicrosoftGraphApi;
            return this;
        }

        /**
         * @param useMicrosoftGraphApi Use the Microsoft Graph API. Should be set to true on vault-1.10+
         * 
         * @return builder
         * 
         */
        public Builder useMicrosoftGraphApi(Boolean useMicrosoftGraphApi) {
            return useMicrosoftGraphApi(Output.of(useMicrosoftGraphApi));
        }

        public BackendState build() {
            return $;
        }
    }

}
