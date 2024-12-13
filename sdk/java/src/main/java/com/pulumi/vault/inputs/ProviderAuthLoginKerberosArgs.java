// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vault.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ProviderAuthLoginKerberosArgs extends com.pulumi.resources.ResourceArgs {

    public static final ProviderAuthLoginKerberosArgs Empty = new ProviderAuthLoginKerberosArgs();

    /**
     * Disable the Kerberos FAST negotiation.
     * 
     */
    @Import(name="disableFastNegotiation")
    private @Nullable Output<Boolean> disableFastNegotiation;

    /**
     * @return Disable the Kerberos FAST negotiation.
     * 
     */
    public Optional<Output<Boolean>> disableFastNegotiation() {
        return Optional.ofNullable(this.disableFastNegotiation);
    }

    /**
     * The Kerberos keytab file containing the entry of the login entity.
     * 
     */
    @Import(name="keytabPath")
    private @Nullable Output<String> keytabPath;

    /**
     * @return The Kerberos keytab file containing the entry of the login entity.
     * 
     */
    public Optional<Output<String>> keytabPath() {
        return Optional.ofNullable(this.keytabPath);
    }

    /**
     * A valid Kerberos configuration file e.g. /etc/krb5.conf.
     * 
     */
    @Import(name="krb5confPath")
    private @Nullable Output<String> krb5confPath;

    /**
     * @return A valid Kerberos configuration file e.g. /etc/krb5.conf.
     * 
     */
    public Optional<Output<String>> krb5confPath() {
        return Optional.ofNullable(this.krb5confPath);
    }

    /**
     * The path where the authentication engine is mounted.
     * 
     */
    @Import(name="mount")
    private @Nullable Output<String> mount;

    /**
     * @return The path where the authentication engine is mounted.
     * 
     */
    public Optional<Output<String>> mount() {
        return Optional.ofNullable(this.mount);
    }

    /**
     * The authentication engine&#39;s namespace. Conflicts with use_root_namespace
     * 
     */
    @Import(name="namespace")
    private @Nullable Output<String> namespace;

    /**
     * @return The authentication engine&#39;s namespace. Conflicts with use_root_namespace
     * 
     */
    public Optional<Output<String>> namespace() {
        return Optional.ofNullable(this.namespace);
    }

    /**
     * The Kerberos server&#39;s authoritative authentication domain
     * 
     */
    @Import(name="realm")
    private @Nullable Output<String> realm;

    /**
     * @return The Kerberos server&#39;s authoritative authentication domain
     * 
     */
    public Optional<Output<String>> realm() {
        return Optional.ofNullable(this.realm);
    }

    /**
     * Strip the host from the username found in the keytab.
     * 
     */
    @Import(name="removeInstanceName")
    private @Nullable Output<Boolean> removeInstanceName;

    /**
     * @return Strip the host from the username found in the keytab.
     * 
     */
    public Optional<Output<Boolean>> removeInstanceName() {
        return Optional.ofNullable(this.removeInstanceName);
    }

    /**
     * The service principle name.
     * 
     */
    @Import(name="service")
    private @Nullable Output<String> service;

    /**
     * @return The service principle name.
     * 
     */
    public Optional<Output<String>> service() {
        return Optional.ofNullable(this.service);
    }

    /**
     * Simple and Protected GSSAPI Negotiation Mechanism (SPNEGO) token
     * 
     */
    @Import(name="token")
    private @Nullable Output<String> token;

    /**
     * @return Simple and Protected GSSAPI Negotiation Mechanism (SPNEGO) token
     * 
     */
    public Optional<Output<String>> token() {
        return Optional.ofNullable(this.token);
    }

    /**
     * Authenticate to the root Vault namespace. Conflicts with namespace
     * 
     */
    @Import(name="useRootNamespace")
    private @Nullable Output<Boolean> useRootNamespace;

    /**
     * @return Authenticate to the root Vault namespace. Conflicts with namespace
     * 
     */
    public Optional<Output<Boolean>> useRootNamespace() {
        return Optional.ofNullable(this.useRootNamespace);
    }

    /**
     * The username to login into Kerberos with.
     * 
     */
    @Import(name="username")
    private @Nullable Output<String> username;

    /**
     * @return The username to login into Kerberos with.
     * 
     */
    public Optional<Output<String>> username() {
        return Optional.ofNullable(this.username);
    }

    private ProviderAuthLoginKerberosArgs() {}

    private ProviderAuthLoginKerberosArgs(ProviderAuthLoginKerberosArgs $) {
        this.disableFastNegotiation = $.disableFastNegotiation;
        this.keytabPath = $.keytabPath;
        this.krb5confPath = $.krb5confPath;
        this.mount = $.mount;
        this.namespace = $.namespace;
        this.realm = $.realm;
        this.removeInstanceName = $.removeInstanceName;
        this.service = $.service;
        this.token = $.token;
        this.useRootNamespace = $.useRootNamespace;
        this.username = $.username;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ProviderAuthLoginKerberosArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ProviderAuthLoginKerberosArgs $;

        public Builder() {
            $ = new ProviderAuthLoginKerberosArgs();
        }

        public Builder(ProviderAuthLoginKerberosArgs defaults) {
            $ = new ProviderAuthLoginKerberosArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param disableFastNegotiation Disable the Kerberos FAST negotiation.
         * 
         * @return builder
         * 
         */
        public Builder disableFastNegotiation(@Nullable Output<Boolean> disableFastNegotiation) {
            $.disableFastNegotiation = disableFastNegotiation;
            return this;
        }

        /**
         * @param disableFastNegotiation Disable the Kerberos FAST negotiation.
         * 
         * @return builder
         * 
         */
        public Builder disableFastNegotiation(Boolean disableFastNegotiation) {
            return disableFastNegotiation(Output.of(disableFastNegotiation));
        }

        /**
         * @param keytabPath The Kerberos keytab file containing the entry of the login entity.
         * 
         * @return builder
         * 
         */
        public Builder keytabPath(@Nullable Output<String> keytabPath) {
            $.keytabPath = keytabPath;
            return this;
        }

        /**
         * @param keytabPath The Kerberos keytab file containing the entry of the login entity.
         * 
         * @return builder
         * 
         */
        public Builder keytabPath(String keytabPath) {
            return keytabPath(Output.of(keytabPath));
        }

        /**
         * @param krb5confPath A valid Kerberos configuration file e.g. /etc/krb5.conf.
         * 
         * @return builder
         * 
         */
        public Builder krb5confPath(@Nullable Output<String> krb5confPath) {
            $.krb5confPath = krb5confPath;
            return this;
        }

        /**
         * @param krb5confPath A valid Kerberos configuration file e.g. /etc/krb5.conf.
         * 
         * @return builder
         * 
         */
        public Builder krb5confPath(String krb5confPath) {
            return krb5confPath(Output.of(krb5confPath));
        }

        /**
         * @param mount The path where the authentication engine is mounted.
         * 
         * @return builder
         * 
         */
        public Builder mount(@Nullable Output<String> mount) {
            $.mount = mount;
            return this;
        }

        /**
         * @param mount The path where the authentication engine is mounted.
         * 
         * @return builder
         * 
         */
        public Builder mount(String mount) {
            return mount(Output.of(mount));
        }

        /**
         * @param namespace The authentication engine&#39;s namespace. Conflicts with use_root_namespace
         * 
         * @return builder
         * 
         */
        public Builder namespace(@Nullable Output<String> namespace) {
            $.namespace = namespace;
            return this;
        }

        /**
         * @param namespace The authentication engine&#39;s namespace. Conflicts with use_root_namespace
         * 
         * @return builder
         * 
         */
        public Builder namespace(String namespace) {
            return namespace(Output.of(namespace));
        }

        /**
         * @param realm The Kerberos server&#39;s authoritative authentication domain
         * 
         * @return builder
         * 
         */
        public Builder realm(@Nullable Output<String> realm) {
            $.realm = realm;
            return this;
        }

        /**
         * @param realm The Kerberos server&#39;s authoritative authentication domain
         * 
         * @return builder
         * 
         */
        public Builder realm(String realm) {
            return realm(Output.of(realm));
        }

        /**
         * @param removeInstanceName Strip the host from the username found in the keytab.
         * 
         * @return builder
         * 
         */
        public Builder removeInstanceName(@Nullable Output<Boolean> removeInstanceName) {
            $.removeInstanceName = removeInstanceName;
            return this;
        }

        /**
         * @param removeInstanceName Strip the host from the username found in the keytab.
         * 
         * @return builder
         * 
         */
        public Builder removeInstanceName(Boolean removeInstanceName) {
            return removeInstanceName(Output.of(removeInstanceName));
        }

        /**
         * @param service The service principle name.
         * 
         * @return builder
         * 
         */
        public Builder service(@Nullable Output<String> service) {
            $.service = service;
            return this;
        }

        /**
         * @param service The service principle name.
         * 
         * @return builder
         * 
         */
        public Builder service(String service) {
            return service(Output.of(service));
        }

        /**
         * @param token Simple and Protected GSSAPI Negotiation Mechanism (SPNEGO) token
         * 
         * @return builder
         * 
         */
        public Builder token(@Nullable Output<String> token) {
            $.token = token;
            return this;
        }

        /**
         * @param token Simple and Protected GSSAPI Negotiation Mechanism (SPNEGO) token
         * 
         * @return builder
         * 
         */
        public Builder token(String token) {
            return token(Output.of(token));
        }

        /**
         * @param useRootNamespace Authenticate to the root Vault namespace. Conflicts with namespace
         * 
         * @return builder
         * 
         */
        public Builder useRootNamespace(@Nullable Output<Boolean> useRootNamespace) {
            $.useRootNamespace = useRootNamespace;
            return this;
        }

        /**
         * @param useRootNamespace Authenticate to the root Vault namespace. Conflicts with namespace
         * 
         * @return builder
         * 
         */
        public Builder useRootNamespace(Boolean useRootNamespace) {
            return useRootNamespace(Output.of(useRootNamespace));
        }

        /**
         * @param username The username to login into Kerberos with.
         * 
         * @return builder
         * 
         */
        public Builder username(@Nullable Output<String> username) {
            $.username = username;
            return this;
        }

        /**
         * @param username The username to login into Kerberos with.
         * 
         * @return builder
         * 
         */
        public Builder username(String username) {
            return username(Output.of(username));
        }

        public ProviderAuthLoginKerberosArgs build() {
            return $;
        }
    }

}