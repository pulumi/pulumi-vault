// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vault.gcp.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class SecretImpersonatedAccountState extends com.pulumi.resources.ResourceArgs {

    public static final SecretImpersonatedAccountState Empty = new SecretImpersonatedAccountState();

    /**
     * Path where the GCP Secrets Engine is mounted
     * 
     */
    @Import(name="backend")
    private @Nullable Output<String> backend;

    /**
     * @return Path where the GCP Secrets Engine is mounted
     * 
     */
    public Optional<Output<String>> backend() {
        return Optional.ofNullable(this.backend);
    }

    /**
     * Name of the Impersonated Account to create
     * 
     */
    @Import(name="impersonatedAccount")
    private @Nullable Output<String> impersonatedAccount;

    /**
     * @return Name of the Impersonated Account to create
     * 
     */
    public Optional<Output<String>> impersonatedAccount() {
        return Optional.ofNullable(this.impersonatedAccount);
    }

    /**
     * Target namespace. (requires Enterprise)
     * 
     */
    @Import(name="namespace")
    private @Nullable Output<String> namespace;

    /**
     * @return Target namespace. (requires Enterprise)
     * 
     */
    public Optional<Output<String>> namespace() {
        return Optional.ofNullable(this.namespace);
    }

    /**
     * Email of the GCP service account to impersonate.
     * 
     */
    @Import(name="serviceAccountEmail")
    private @Nullable Output<String> serviceAccountEmail;

    /**
     * @return Email of the GCP service account to impersonate.
     * 
     */
    public Optional<Output<String>> serviceAccountEmail() {
        return Optional.ofNullable(this.serviceAccountEmail);
    }

    /**
     * Project the service account belongs to.
     * 
     */
    @Import(name="serviceAccountProject")
    private @Nullable Output<String> serviceAccountProject;

    /**
     * @return Project the service account belongs to.
     * 
     */
    public Optional<Output<String>> serviceAccountProject() {
        return Optional.ofNullable(this.serviceAccountProject);
    }

    /**
     * List of OAuth scopes to assign to access tokens generated under this impersonated account.
     * 
     */
    @Import(name="tokenScopes")
    private @Nullable Output<List<String>> tokenScopes;

    /**
     * @return List of OAuth scopes to assign to access tokens generated under this impersonated account.
     * 
     */
    public Optional<Output<List<String>>> tokenScopes() {
        return Optional.ofNullable(this.tokenScopes);
    }

    /**
     * Specifies the default TTL for service principals generated using this role.
     * Accepts time suffixed strings (&#34;1h&#34;) or an integer number of seconds. Defaults to the system/engine default TTL time.
     * 
     */
    @Import(name="ttl")
    private @Nullable Output<String> ttl;

    /**
     * @return Specifies the default TTL for service principals generated using this role.
     * Accepts time suffixed strings (&#34;1h&#34;) or an integer number of seconds. Defaults to the system/engine default TTL time.
     * 
     */
    public Optional<Output<String>> ttl() {
        return Optional.ofNullable(this.ttl);
    }

    private SecretImpersonatedAccountState() {}

    private SecretImpersonatedAccountState(SecretImpersonatedAccountState $) {
        this.backend = $.backend;
        this.impersonatedAccount = $.impersonatedAccount;
        this.namespace = $.namespace;
        this.serviceAccountEmail = $.serviceAccountEmail;
        this.serviceAccountProject = $.serviceAccountProject;
        this.tokenScopes = $.tokenScopes;
        this.ttl = $.ttl;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(SecretImpersonatedAccountState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private SecretImpersonatedAccountState $;

        public Builder() {
            $ = new SecretImpersonatedAccountState();
        }

        public Builder(SecretImpersonatedAccountState defaults) {
            $ = new SecretImpersonatedAccountState(Objects.requireNonNull(defaults));
        }

        /**
         * @param backend Path where the GCP Secrets Engine is mounted
         * 
         * @return builder
         * 
         */
        public Builder backend(@Nullable Output<String> backend) {
            $.backend = backend;
            return this;
        }

        /**
         * @param backend Path where the GCP Secrets Engine is mounted
         * 
         * @return builder
         * 
         */
        public Builder backend(String backend) {
            return backend(Output.of(backend));
        }

        /**
         * @param impersonatedAccount Name of the Impersonated Account to create
         * 
         * @return builder
         * 
         */
        public Builder impersonatedAccount(@Nullable Output<String> impersonatedAccount) {
            $.impersonatedAccount = impersonatedAccount;
            return this;
        }

        /**
         * @param impersonatedAccount Name of the Impersonated Account to create
         * 
         * @return builder
         * 
         */
        public Builder impersonatedAccount(String impersonatedAccount) {
            return impersonatedAccount(Output.of(impersonatedAccount));
        }

        /**
         * @param namespace Target namespace. (requires Enterprise)
         * 
         * @return builder
         * 
         */
        public Builder namespace(@Nullable Output<String> namespace) {
            $.namespace = namespace;
            return this;
        }

        /**
         * @param namespace Target namespace. (requires Enterprise)
         * 
         * @return builder
         * 
         */
        public Builder namespace(String namespace) {
            return namespace(Output.of(namespace));
        }

        /**
         * @param serviceAccountEmail Email of the GCP service account to impersonate.
         * 
         * @return builder
         * 
         */
        public Builder serviceAccountEmail(@Nullable Output<String> serviceAccountEmail) {
            $.serviceAccountEmail = serviceAccountEmail;
            return this;
        }

        /**
         * @param serviceAccountEmail Email of the GCP service account to impersonate.
         * 
         * @return builder
         * 
         */
        public Builder serviceAccountEmail(String serviceAccountEmail) {
            return serviceAccountEmail(Output.of(serviceAccountEmail));
        }

        /**
         * @param serviceAccountProject Project the service account belongs to.
         * 
         * @return builder
         * 
         */
        public Builder serviceAccountProject(@Nullable Output<String> serviceAccountProject) {
            $.serviceAccountProject = serviceAccountProject;
            return this;
        }

        /**
         * @param serviceAccountProject Project the service account belongs to.
         * 
         * @return builder
         * 
         */
        public Builder serviceAccountProject(String serviceAccountProject) {
            return serviceAccountProject(Output.of(serviceAccountProject));
        }

        /**
         * @param tokenScopes List of OAuth scopes to assign to access tokens generated under this impersonated account.
         * 
         * @return builder
         * 
         */
        public Builder tokenScopes(@Nullable Output<List<String>> tokenScopes) {
            $.tokenScopes = tokenScopes;
            return this;
        }

        /**
         * @param tokenScopes List of OAuth scopes to assign to access tokens generated under this impersonated account.
         * 
         * @return builder
         * 
         */
        public Builder tokenScopes(List<String> tokenScopes) {
            return tokenScopes(Output.of(tokenScopes));
        }

        /**
         * @param tokenScopes List of OAuth scopes to assign to access tokens generated under this impersonated account.
         * 
         * @return builder
         * 
         */
        public Builder tokenScopes(String... tokenScopes) {
            return tokenScopes(List.of(tokenScopes));
        }

        /**
         * @param ttl Specifies the default TTL for service principals generated using this role.
         * Accepts time suffixed strings (&#34;1h&#34;) or an integer number of seconds. Defaults to the system/engine default TTL time.
         * 
         * @return builder
         * 
         */
        public Builder ttl(@Nullable Output<String> ttl) {
            $.ttl = ttl;
            return this;
        }

        /**
         * @param ttl Specifies the default TTL for service principals generated using this role.
         * Accepts time suffixed strings (&#34;1h&#34;) or an integer number of seconds. Defaults to the system/engine default TTL time.
         * 
         * @return builder
         * 
         */
        public Builder ttl(String ttl) {
            return ttl(Output.of(ttl));
        }

        public SecretImpersonatedAccountState build() {
            return $;
        }
    }

}
