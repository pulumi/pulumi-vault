// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vault;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.core.internal.Codegen;
import com.pulumi.vault.inputs.ProviderAuthLoginArgs;
import com.pulumi.vault.inputs.ProviderAuthLoginAwsArgs;
import com.pulumi.vault.inputs.ProviderAuthLoginAzureArgs;
import com.pulumi.vault.inputs.ProviderAuthLoginCertArgs;
import com.pulumi.vault.inputs.ProviderAuthLoginGcpArgs;
import com.pulumi.vault.inputs.ProviderAuthLoginJwtArgs;
import com.pulumi.vault.inputs.ProviderAuthLoginKerberosArgs;
import com.pulumi.vault.inputs.ProviderAuthLoginOciArgs;
import com.pulumi.vault.inputs.ProviderAuthLoginOidcArgs;
import com.pulumi.vault.inputs.ProviderAuthLoginRadiusArgs;
import com.pulumi.vault.inputs.ProviderAuthLoginTokenFileArgs;
import com.pulumi.vault.inputs.ProviderAuthLoginUserpassArgs;
import com.pulumi.vault.inputs.ProviderClientAuthArgs;
import com.pulumi.vault.inputs.ProviderHeaderArgs;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ProviderArgs extends com.pulumi.resources.ResourceArgs {

    public static final ProviderArgs Empty = new ProviderArgs();

    @Import(name="addAddressToEnv")
    private @Nullable Output<String> addAddressToEnv;

    public Optional<Output<String>> addAddressToEnv() {
        return Optional.ofNullable(this.addAddressToEnv);
    }

    /**
     * URL of the root of the target Vault server.
     * 
     */
    @Import(name="address")
    private @Nullable Output<String> address;

    /**
     * @return URL of the root of the target Vault server.
     * 
     */
    public Optional<Output<String>> address() {
        return Optional.ofNullable(this.address);
    }

    /**
     * Login to vault with an existing auth method using auth/&lt;mount&gt;/login
     * 
     */
    @Import(name="authLogin", json=true)
    private @Nullable Output<ProviderAuthLoginArgs> authLogin;

    /**
     * @return Login to vault with an existing auth method using auth/&lt;mount&gt;/login
     * 
     */
    public Optional<Output<ProviderAuthLoginArgs>> authLogin() {
        return Optional.ofNullable(this.authLogin);
    }

    /**
     * Login to vault using the AWS method
     * 
     */
    @Import(name="authLoginAws", json=true)
    private @Nullable Output<ProviderAuthLoginAwsArgs> authLoginAws;

    /**
     * @return Login to vault using the AWS method
     * 
     */
    public Optional<Output<ProviderAuthLoginAwsArgs>> authLoginAws() {
        return Optional.ofNullable(this.authLoginAws);
    }

    /**
     * Login to vault using the azure method
     * 
     */
    @Import(name="authLoginAzure", json=true)
    private @Nullable Output<ProviderAuthLoginAzureArgs> authLoginAzure;

    /**
     * @return Login to vault using the azure method
     * 
     */
    public Optional<Output<ProviderAuthLoginAzureArgs>> authLoginAzure() {
        return Optional.ofNullable(this.authLoginAzure);
    }

    /**
     * Login to vault using the cert method
     * 
     */
    @Import(name="authLoginCert", json=true)
    private @Nullable Output<ProviderAuthLoginCertArgs> authLoginCert;

    /**
     * @return Login to vault using the cert method
     * 
     */
    public Optional<Output<ProviderAuthLoginCertArgs>> authLoginCert() {
        return Optional.ofNullable(this.authLoginCert);
    }

    /**
     * Login to vault using the gcp method
     * 
     */
    @Import(name="authLoginGcp", json=true)
    private @Nullable Output<ProviderAuthLoginGcpArgs> authLoginGcp;

    /**
     * @return Login to vault using the gcp method
     * 
     */
    public Optional<Output<ProviderAuthLoginGcpArgs>> authLoginGcp() {
        return Optional.ofNullable(this.authLoginGcp);
    }

    /**
     * Login to vault using the jwt method
     * 
     */
    @Import(name="authLoginJwt", json=true)
    private @Nullable Output<ProviderAuthLoginJwtArgs> authLoginJwt;

    /**
     * @return Login to vault using the jwt method
     * 
     */
    public Optional<Output<ProviderAuthLoginJwtArgs>> authLoginJwt() {
        return Optional.ofNullable(this.authLoginJwt);
    }

    /**
     * Login to vault using the kerberos method
     * 
     */
    @Import(name="authLoginKerberos", json=true)
    private @Nullable Output<ProviderAuthLoginKerberosArgs> authLoginKerberos;

    /**
     * @return Login to vault using the kerberos method
     * 
     */
    public Optional<Output<ProviderAuthLoginKerberosArgs>> authLoginKerberos() {
        return Optional.ofNullable(this.authLoginKerberos);
    }

    /**
     * Login to vault using the OCI method
     * 
     */
    @Import(name="authLoginOci", json=true)
    private @Nullable Output<ProviderAuthLoginOciArgs> authLoginOci;

    /**
     * @return Login to vault using the OCI method
     * 
     */
    public Optional<Output<ProviderAuthLoginOciArgs>> authLoginOci() {
        return Optional.ofNullable(this.authLoginOci);
    }

    /**
     * Login to vault using the oidc method
     * 
     */
    @Import(name="authLoginOidc", json=true)
    private @Nullable Output<ProviderAuthLoginOidcArgs> authLoginOidc;

    /**
     * @return Login to vault using the oidc method
     * 
     */
    public Optional<Output<ProviderAuthLoginOidcArgs>> authLoginOidc() {
        return Optional.ofNullable(this.authLoginOidc);
    }

    /**
     * Login to vault using the radius method
     * 
     */
    @Import(name="authLoginRadius", json=true)
    private @Nullable Output<ProviderAuthLoginRadiusArgs> authLoginRadius;

    /**
     * @return Login to vault using the radius method
     * 
     */
    public Optional<Output<ProviderAuthLoginRadiusArgs>> authLoginRadius() {
        return Optional.ofNullable(this.authLoginRadius);
    }

    /**
     * Login to vault using
     * 
     */
    @Import(name="authLoginTokenFile", json=true)
    private @Nullable Output<ProviderAuthLoginTokenFileArgs> authLoginTokenFile;

    /**
     * @return Login to vault using
     * 
     */
    public Optional<Output<ProviderAuthLoginTokenFileArgs>> authLoginTokenFile() {
        return Optional.ofNullable(this.authLoginTokenFile);
    }

    /**
     * Login to vault using the userpass method
     * 
     */
    @Import(name="authLoginUserpass", json=true)
    private @Nullable Output<ProviderAuthLoginUserpassArgs> authLoginUserpass;

    /**
     * @return Login to vault using the userpass method
     * 
     */
    public Optional<Output<ProviderAuthLoginUserpassArgs>> authLoginUserpass() {
        return Optional.ofNullable(this.authLoginUserpass);
    }

    /**
     * Path to directory containing CA certificate files to validate the server&#39;s certificate.
     * 
     */
    @Import(name="caCertDir")
    private @Nullable Output<String> caCertDir;

    /**
     * @return Path to directory containing CA certificate files to validate the server&#39;s certificate.
     * 
     */
    public Optional<Output<String>> caCertDir() {
        return Optional.ofNullable(this.caCertDir);
    }

    /**
     * Path to a CA certificate file to validate the server&#39;s certificate.
     * 
     */
    @Import(name="caCertFile")
    private @Nullable Output<String> caCertFile;

    /**
     * @return Path to a CA certificate file to validate the server&#39;s certificate.
     * 
     */
    public Optional<Output<String>> caCertFile() {
        return Optional.ofNullable(this.caCertFile);
    }

    /**
     * Client authentication credentials.
     * 
     */
    @Import(name="clientAuth", json=true)
    private @Nullable Output<ProviderClientAuthArgs> clientAuth;

    /**
     * @return Client authentication credentials.
     * 
     */
    public Optional<Output<ProviderClientAuthArgs>> clientAuth() {
        return Optional.ofNullable(this.clientAuth);
    }

    /**
     * The headers to send with each Vault request.
     * 
     */
    @Import(name="headers", json=true)
    private @Nullable Output<List<ProviderHeaderArgs>> headers;

    /**
     * @return The headers to send with each Vault request.
     * 
     */
    public Optional<Output<List<ProviderHeaderArgs>>> headers() {
        return Optional.ofNullable(this.headers);
    }

    /**
     * Maximum TTL for secret leases requested by this provider.
     * 
     */
    @Import(name="maxLeaseTtlSeconds", json=true)
    private @Nullable Output<Integer> maxLeaseTtlSeconds;

    /**
     * @return Maximum TTL for secret leases requested by this provider.
     * 
     */
    public Optional<Output<Integer>> maxLeaseTtlSeconds() {
        return Optional.ofNullable(this.maxLeaseTtlSeconds);
    }

    /**
     * Maximum number of retries when a 5xx error code is encountered.
     * 
     */
    @Import(name="maxRetries", json=true)
    private @Nullable Output<Integer> maxRetries;

    /**
     * @return Maximum number of retries when a 5xx error code is encountered.
     * 
     */
    public Optional<Output<Integer>> maxRetries() {
        return Optional.ofNullable(this.maxRetries);
    }

    /**
     * Maximum number of retries for Client Controlled Consistency related operations
     * 
     */
    @Import(name="maxRetriesCcc", json=true)
    private @Nullable Output<Integer> maxRetriesCcc;

    /**
     * @return Maximum number of retries for Client Controlled Consistency related operations
     * 
     */
    public Optional<Output<Integer>> maxRetriesCcc() {
        return Optional.ofNullable(this.maxRetriesCcc);
    }

    /**
     * The namespace to use. Available only for Vault Enterprise.
     * 
     */
    @Import(name="namespace")
    private @Nullable Output<String> namespace;

    /**
     * @return The namespace to use. Available only for Vault Enterprise.
     * 
     */
    public Optional<Output<String>> namespace() {
        return Optional.ofNullable(this.namespace);
    }

    /**
     * In the case where the Vault token is for a specific namespace and the provider namespace is not configured, use the
     * token namespace as the root namespace for all resources.
     * 
     */
    @Import(name="setNamespaceFromToken", json=true)
    private @Nullable Output<Boolean> setNamespaceFromToken;

    /**
     * @return In the case where the Vault token is for a specific namespace and the provider namespace is not configured, use the
     * token namespace as the root namespace for all resources.
     * 
     */
    public Optional<Output<Boolean>> setNamespaceFromToken() {
        return Optional.ofNullable(this.setNamespaceFromToken);
    }

    /**
     * Set this to true to prevent the creation of ephemeral child token used by this provider.
     * 
     */
    @Import(name="skipChildToken", json=true)
    private @Nullable Output<Boolean> skipChildToken;

    /**
     * @return Set this to true to prevent the creation of ephemeral child token used by this provider.
     * 
     */
    public Optional<Output<Boolean>> skipChildToken() {
        return Optional.ofNullable(this.skipChildToken);
    }

    /**
     * Skip the dynamic fetching of the Vault server version.
     * 
     */
    @Import(name="skipGetVaultVersion", json=true)
    private @Nullable Output<Boolean> skipGetVaultVersion;

    /**
     * @return Skip the dynamic fetching of the Vault server version.
     * 
     */
    public Optional<Output<Boolean>> skipGetVaultVersion() {
        return Optional.ofNullable(this.skipGetVaultVersion);
    }

    /**
     * Set this to true only if the target Vault server is an insecure development instance.
     * 
     */
    @Import(name="skipTlsVerify", json=true)
    private @Nullable Output<Boolean> skipTlsVerify;

    /**
     * @return Set this to true only if the target Vault server is an insecure development instance.
     * 
     */
    public Optional<Output<Boolean>> skipTlsVerify() {
        return Optional.ofNullable(this.skipTlsVerify);
    }

    /**
     * Name to use as the SNI host when connecting via TLS.
     * 
     */
    @Import(name="tlsServerName")
    private @Nullable Output<String> tlsServerName;

    /**
     * @return Name to use as the SNI host when connecting via TLS.
     * 
     */
    public Optional<Output<String>> tlsServerName() {
        return Optional.ofNullable(this.tlsServerName);
    }

    /**
     * Token to use to authenticate to Vault.
     * 
     */
    @Import(name="token")
    private @Nullable Output<String> token;

    /**
     * @return Token to use to authenticate to Vault.
     * 
     */
    public Optional<Output<String>> token() {
        return Optional.ofNullable(this.token);
    }

    /**
     * Token name to use for creating the Vault child token.
     * 
     */
    @Import(name="tokenName")
    private @Nullable Output<String> tokenName;

    /**
     * @return Token name to use for creating the Vault child token.
     * 
     */
    public Optional<Output<String>> tokenName() {
        return Optional.ofNullable(this.tokenName);
    }

    /**
     * Override the Vault server version, which is normally determined dynamically from the target Vault server
     * 
     */
    @Import(name="vaultVersionOverride")
    private @Nullable Output<String> vaultVersionOverride;

    /**
     * @return Override the Vault server version, which is normally determined dynamically from the target Vault server
     * 
     */
    public Optional<Output<String>> vaultVersionOverride() {
        return Optional.ofNullable(this.vaultVersionOverride);
    }

    private ProviderArgs() {}

    private ProviderArgs(ProviderArgs $) {
        this.addAddressToEnv = $.addAddressToEnv;
        this.address = $.address;
        this.authLogin = $.authLogin;
        this.authLoginAws = $.authLoginAws;
        this.authLoginAzure = $.authLoginAzure;
        this.authLoginCert = $.authLoginCert;
        this.authLoginGcp = $.authLoginGcp;
        this.authLoginJwt = $.authLoginJwt;
        this.authLoginKerberos = $.authLoginKerberos;
        this.authLoginOci = $.authLoginOci;
        this.authLoginOidc = $.authLoginOidc;
        this.authLoginRadius = $.authLoginRadius;
        this.authLoginTokenFile = $.authLoginTokenFile;
        this.authLoginUserpass = $.authLoginUserpass;
        this.caCertDir = $.caCertDir;
        this.caCertFile = $.caCertFile;
        this.clientAuth = $.clientAuth;
        this.headers = $.headers;
        this.maxLeaseTtlSeconds = $.maxLeaseTtlSeconds;
        this.maxRetries = $.maxRetries;
        this.maxRetriesCcc = $.maxRetriesCcc;
        this.namespace = $.namespace;
        this.setNamespaceFromToken = $.setNamespaceFromToken;
        this.skipChildToken = $.skipChildToken;
        this.skipGetVaultVersion = $.skipGetVaultVersion;
        this.skipTlsVerify = $.skipTlsVerify;
        this.tlsServerName = $.tlsServerName;
        this.token = $.token;
        this.tokenName = $.tokenName;
        this.vaultVersionOverride = $.vaultVersionOverride;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ProviderArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ProviderArgs $;

        public Builder() {
            $ = new ProviderArgs();
        }

        public Builder(ProviderArgs defaults) {
            $ = new ProviderArgs(Objects.requireNonNull(defaults));
        }

        public Builder addAddressToEnv(@Nullable Output<String> addAddressToEnv) {
            $.addAddressToEnv = addAddressToEnv;
            return this;
        }

        public Builder addAddressToEnv(String addAddressToEnv) {
            return addAddressToEnv(Output.of(addAddressToEnv));
        }

        /**
         * @param address URL of the root of the target Vault server.
         * 
         * @return builder
         * 
         */
        public Builder address(@Nullable Output<String> address) {
            $.address = address;
            return this;
        }

        /**
         * @param address URL of the root of the target Vault server.
         * 
         * @return builder
         * 
         */
        public Builder address(String address) {
            return address(Output.of(address));
        }

        /**
         * @param authLogin Login to vault with an existing auth method using auth/&lt;mount&gt;/login
         * 
         * @return builder
         * 
         */
        public Builder authLogin(@Nullable Output<ProviderAuthLoginArgs> authLogin) {
            $.authLogin = authLogin;
            return this;
        }

        /**
         * @param authLogin Login to vault with an existing auth method using auth/&lt;mount&gt;/login
         * 
         * @return builder
         * 
         */
        public Builder authLogin(ProviderAuthLoginArgs authLogin) {
            return authLogin(Output.of(authLogin));
        }

        /**
         * @param authLoginAws Login to vault using the AWS method
         * 
         * @return builder
         * 
         */
        public Builder authLoginAws(@Nullable Output<ProviderAuthLoginAwsArgs> authLoginAws) {
            $.authLoginAws = authLoginAws;
            return this;
        }

        /**
         * @param authLoginAws Login to vault using the AWS method
         * 
         * @return builder
         * 
         */
        public Builder authLoginAws(ProviderAuthLoginAwsArgs authLoginAws) {
            return authLoginAws(Output.of(authLoginAws));
        }

        /**
         * @param authLoginAzure Login to vault using the azure method
         * 
         * @return builder
         * 
         */
        public Builder authLoginAzure(@Nullable Output<ProviderAuthLoginAzureArgs> authLoginAzure) {
            $.authLoginAzure = authLoginAzure;
            return this;
        }

        /**
         * @param authLoginAzure Login to vault using the azure method
         * 
         * @return builder
         * 
         */
        public Builder authLoginAzure(ProviderAuthLoginAzureArgs authLoginAzure) {
            return authLoginAzure(Output.of(authLoginAzure));
        }

        /**
         * @param authLoginCert Login to vault using the cert method
         * 
         * @return builder
         * 
         */
        public Builder authLoginCert(@Nullable Output<ProviderAuthLoginCertArgs> authLoginCert) {
            $.authLoginCert = authLoginCert;
            return this;
        }

        /**
         * @param authLoginCert Login to vault using the cert method
         * 
         * @return builder
         * 
         */
        public Builder authLoginCert(ProviderAuthLoginCertArgs authLoginCert) {
            return authLoginCert(Output.of(authLoginCert));
        }

        /**
         * @param authLoginGcp Login to vault using the gcp method
         * 
         * @return builder
         * 
         */
        public Builder authLoginGcp(@Nullable Output<ProviderAuthLoginGcpArgs> authLoginGcp) {
            $.authLoginGcp = authLoginGcp;
            return this;
        }

        /**
         * @param authLoginGcp Login to vault using the gcp method
         * 
         * @return builder
         * 
         */
        public Builder authLoginGcp(ProviderAuthLoginGcpArgs authLoginGcp) {
            return authLoginGcp(Output.of(authLoginGcp));
        }

        /**
         * @param authLoginJwt Login to vault using the jwt method
         * 
         * @return builder
         * 
         */
        public Builder authLoginJwt(@Nullable Output<ProviderAuthLoginJwtArgs> authLoginJwt) {
            $.authLoginJwt = authLoginJwt;
            return this;
        }

        /**
         * @param authLoginJwt Login to vault using the jwt method
         * 
         * @return builder
         * 
         */
        public Builder authLoginJwt(ProviderAuthLoginJwtArgs authLoginJwt) {
            return authLoginJwt(Output.of(authLoginJwt));
        }

        /**
         * @param authLoginKerberos Login to vault using the kerberos method
         * 
         * @return builder
         * 
         */
        public Builder authLoginKerberos(@Nullable Output<ProviderAuthLoginKerberosArgs> authLoginKerberos) {
            $.authLoginKerberos = authLoginKerberos;
            return this;
        }

        /**
         * @param authLoginKerberos Login to vault using the kerberos method
         * 
         * @return builder
         * 
         */
        public Builder authLoginKerberos(ProviderAuthLoginKerberosArgs authLoginKerberos) {
            return authLoginKerberos(Output.of(authLoginKerberos));
        }

        /**
         * @param authLoginOci Login to vault using the OCI method
         * 
         * @return builder
         * 
         */
        public Builder authLoginOci(@Nullable Output<ProviderAuthLoginOciArgs> authLoginOci) {
            $.authLoginOci = authLoginOci;
            return this;
        }

        /**
         * @param authLoginOci Login to vault using the OCI method
         * 
         * @return builder
         * 
         */
        public Builder authLoginOci(ProviderAuthLoginOciArgs authLoginOci) {
            return authLoginOci(Output.of(authLoginOci));
        }

        /**
         * @param authLoginOidc Login to vault using the oidc method
         * 
         * @return builder
         * 
         */
        public Builder authLoginOidc(@Nullable Output<ProviderAuthLoginOidcArgs> authLoginOidc) {
            $.authLoginOidc = authLoginOidc;
            return this;
        }

        /**
         * @param authLoginOidc Login to vault using the oidc method
         * 
         * @return builder
         * 
         */
        public Builder authLoginOidc(ProviderAuthLoginOidcArgs authLoginOidc) {
            return authLoginOidc(Output.of(authLoginOidc));
        }

        /**
         * @param authLoginRadius Login to vault using the radius method
         * 
         * @return builder
         * 
         */
        public Builder authLoginRadius(@Nullable Output<ProviderAuthLoginRadiusArgs> authLoginRadius) {
            $.authLoginRadius = authLoginRadius;
            return this;
        }

        /**
         * @param authLoginRadius Login to vault using the radius method
         * 
         * @return builder
         * 
         */
        public Builder authLoginRadius(ProviderAuthLoginRadiusArgs authLoginRadius) {
            return authLoginRadius(Output.of(authLoginRadius));
        }

        /**
         * @param authLoginTokenFile Login to vault using
         * 
         * @return builder
         * 
         */
        public Builder authLoginTokenFile(@Nullable Output<ProviderAuthLoginTokenFileArgs> authLoginTokenFile) {
            $.authLoginTokenFile = authLoginTokenFile;
            return this;
        }

        /**
         * @param authLoginTokenFile Login to vault using
         * 
         * @return builder
         * 
         */
        public Builder authLoginTokenFile(ProviderAuthLoginTokenFileArgs authLoginTokenFile) {
            return authLoginTokenFile(Output.of(authLoginTokenFile));
        }

        /**
         * @param authLoginUserpass Login to vault using the userpass method
         * 
         * @return builder
         * 
         */
        public Builder authLoginUserpass(@Nullable Output<ProviderAuthLoginUserpassArgs> authLoginUserpass) {
            $.authLoginUserpass = authLoginUserpass;
            return this;
        }

        /**
         * @param authLoginUserpass Login to vault using the userpass method
         * 
         * @return builder
         * 
         */
        public Builder authLoginUserpass(ProviderAuthLoginUserpassArgs authLoginUserpass) {
            return authLoginUserpass(Output.of(authLoginUserpass));
        }

        /**
         * @param caCertDir Path to directory containing CA certificate files to validate the server&#39;s certificate.
         * 
         * @return builder
         * 
         */
        public Builder caCertDir(@Nullable Output<String> caCertDir) {
            $.caCertDir = caCertDir;
            return this;
        }

        /**
         * @param caCertDir Path to directory containing CA certificate files to validate the server&#39;s certificate.
         * 
         * @return builder
         * 
         */
        public Builder caCertDir(String caCertDir) {
            return caCertDir(Output.of(caCertDir));
        }

        /**
         * @param caCertFile Path to a CA certificate file to validate the server&#39;s certificate.
         * 
         * @return builder
         * 
         */
        public Builder caCertFile(@Nullable Output<String> caCertFile) {
            $.caCertFile = caCertFile;
            return this;
        }

        /**
         * @param caCertFile Path to a CA certificate file to validate the server&#39;s certificate.
         * 
         * @return builder
         * 
         */
        public Builder caCertFile(String caCertFile) {
            return caCertFile(Output.of(caCertFile));
        }

        /**
         * @param clientAuth Client authentication credentials.
         * 
         * @return builder
         * 
         */
        public Builder clientAuth(@Nullable Output<ProviderClientAuthArgs> clientAuth) {
            $.clientAuth = clientAuth;
            return this;
        }

        /**
         * @param clientAuth Client authentication credentials.
         * 
         * @return builder
         * 
         */
        public Builder clientAuth(ProviderClientAuthArgs clientAuth) {
            return clientAuth(Output.of(clientAuth));
        }

        /**
         * @param headers The headers to send with each Vault request.
         * 
         * @return builder
         * 
         */
        public Builder headers(@Nullable Output<List<ProviderHeaderArgs>> headers) {
            $.headers = headers;
            return this;
        }

        /**
         * @param headers The headers to send with each Vault request.
         * 
         * @return builder
         * 
         */
        public Builder headers(List<ProviderHeaderArgs> headers) {
            return headers(Output.of(headers));
        }

        /**
         * @param headers The headers to send with each Vault request.
         * 
         * @return builder
         * 
         */
        public Builder headers(ProviderHeaderArgs... headers) {
            return headers(List.of(headers));
        }

        /**
         * @param maxLeaseTtlSeconds Maximum TTL for secret leases requested by this provider.
         * 
         * @return builder
         * 
         */
        public Builder maxLeaseTtlSeconds(@Nullable Output<Integer> maxLeaseTtlSeconds) {
            $.maxLeaseTtlSeconds = maxLeaseTtlSeconds;
            return this;
        }

        /**
         * @param maxLeaseTtlSeconds Maximum TTL for secret leases requested by this provider.
         * 
         * @return builder
         * 
         */
        public Builder maxLeaseTtlSeconds(Integer maxLeaseTtlSeconds) {
            return maxLeaseTtlSeconds(Output.of(maxLeaseTtlSeconds));
        }

        /**
         * @param maxRetries Maximum number of retries when a 5xx error code is encountered.
         * 
         * @return builder
         * 
         */
        public Builder maxRetries(@Nullable Output<Integer> maxRetries) {
            $.maxRetries = maxRetries;
            return this;
        }

        /**
         * @param maxRetries Maximum number of retries when a 5xx error code is encountered.
         * 
         * @return builder
         * 
         */
        public Builder maxRetries(Integer maxRetries) {
            return maxRetries(Output.of(maxRetries));
        }

        /**
         * @param maxRetriesCcc Maximum number of retries for Client Controlled Consistency related operations
         * 
         * @return builder
         * 
         */
        public Builder maxRetriesCcc(@Nullable Output<Integer> maxRetriesCcc) {
            $.maxRetriesCcc = maxRetriesCcc;
            return this;
        }

        /**
         * @param maxRetriesCcc Maximum number of retries for Client Controlled Consistency related operations
         * 
         * @return builder
         * 
         */
        public Builder maxRetriesCcc(Integer maxRetriesCcc) {
            return maxRetriesCcc(Output.of(maxRetriesCcc));
        }

        /**
         * @param namespace The namespace to use. Available only for Vault Enterprise.
         * 
         * @return builder
         * 
         */
        public Builder namespace(@Nullable Output<String> namespace) {
            $.namespace = namespace;
            return this;
        }

        /**
         * @param namespace The namespace to use. Available only for Vault Enterprise.
         * 
         * @return builder
         * 
         */
        public Builder namespace(String namespace) {
            return namespace(Output.of(namespace));
        }

        /**
         * @param setNamespaceFromToken In the case where the Vault token is for a specific namespace and the provider namespace is not configured, use the
         * token namespace as the root namespace for all resources.
         * 
         * @return builder
         * 
         */
        public Builder setNamespaceFromToken(@Nullable Output<Boolean> setNamespaceFromToken) {
            $.setNamespaceFromToken = setNamespaceFromToken;
            return this;
        }

        /**
         * @param setNamespaceFromToken In the case where the Vault token is for a specific namespace and the provider namespace is not configured, use the
         * token namespace as the root namespace for all resources.
         * 
         * @return builder
         * 
         */
        public Builder setNamespaceFromToken(Boolean setNamespaceFromToken) {
            return setNamespaceFromToken(Output.of(setNamespaceFromToken));
        }

        /**
         * @param skipChildToken Set this to true to prevent the creation of ephemeral child token used by this provider.
         * 
         * @return builder
         * 
         */
        public Builder skipChildToken(@Nullable Output<Boolean> skipChildToken) {
            $.skipChildToken = skipChildToken;
            return this;
        }

        /**
         * @param skipChildToken Set this to true to prevent the creation of ephemeral child token used by this provider.
         * 
         * @return builder
         * 
         */
        public Builder skipChildToken(Boolean skipChildToken) {
            return skipChildToken(Output.of(skipChildToken));
        }

        /**
         * @param skipGetVaultVersion Skip the dynamic fetching of the Vault server version.
         * 
         * @return builder
         * 
         */
        public Builder skipGetVaultVersion(@Nullable Output<Boolean> skipGetVaultVersion) {
            $.skipGetVaultVersion = skipGetVaultVersion;
            return this;
        }

        /**
         * @param skipGetVaultVersion Skip the dynamic fetching of the Vault server version.
         * 
         * @return builder
         * 
         */
        public Builder skipGetVaultVersion(Boolean skipGetVaultVersion) {
            return skipGetVaultVersion(Output.of(skipGetVaultVersion));
        }

        /**
         * @param skipTlsVerify Set this to true only if the target Vault server is an insecure development instance.
         * 
         * @return builder
         * 
         */
        public Builder skipTlsVerify(@Nullable Output<Boolean> skipTlsVerify) {
            $.skipTlsVerify = skipTlsVerify;
            return this;
        }

        /**
         * @param skipTlsVerify Set this to true only if the target Vault server is an insecure development instance.
         * 
         * @return builder
         * 
         */
        public Builder skipTlsVerify(Boolean skipTlsVerify) {
            return skipTlsVerify(Output.of(skipTlsVerify));
        }

        /**
         * @param tlsServerName Name to use as the SNI host when connecting via TLS.
         * 
         * @return builder
         * 
         */
        public Builder tlsServerName(@Nullable Output<String> tlsServerName) {
            $.tlsServerName = tlsServerName;
            return this;
        }

        /**
         * @param tlsServerName Name to use as the SNI host when connecting via TLS.
         * 
         * @return builder
         * 
         */
        public Builder tlsServerName(String tlsServerName) {
            return tlsServerName(Output.of(tlsServerName));
        }

        /**
         * @param token Token to use to authenticate to Vault.
         * 
         * @return builder
         * 
         */
        public Builder token(@Nullable Output<String> token) {
            $.token = token;
            return this;
        }

        /**
         * @param token Token to use to authenticate to Vault.
         * 
         * @return builder
         * 
         */
        public Builder token(String token) {
            return token(Output.of(token));
        }

        /**
         * @param tokenName Token name to use for creating the Vault child token.
         * 
         * @return builder
         * 
         */
        public Builder tokenName(@Nullable Output<String> tokenName) {
            $.tokenName = tokenName;
            return this;
        }

        /**
         * @param tokenName Token name to use for creating the Vault child token.
         * 
         * @return builder
         * 
         */
        public Builder tokenName(String tokenName) {
            return tokenName(Output.of(tokenName));
        }

        /**
         * @param vaultVersionOverride Override the Vault server version, which is normally determined dynamically from the target Vault server
         * 
         * @return builder
         * 
         */
        public Builder vaultVersionOverride(@Nullable Output<String> vaultVersionOverride) {
            $.vaultVersionOverride = vaultVersionOverride;
            return this;
        }

        /**
         * @param vaultVersionOverride Override the Vault server version, which is normally determined dynamically from the target Vault server
         * 
         * @return builder
         * 
         */
        public Builder vaultVersionOverride(String vaultVersionOverride) {
            return vaultVersionOverride(Output.of(vaultVersionOverride));
        }

        public ProviderArgs build() {
            $.maxLeaseTtlSeconds = Codegen.integerProp("maxLeaseTtlSeconds").output().arg($.maxLeaseTtlSeconds).env("TERRAFORM_VAULT_MAX_TTL").def(1200).getNullable();
            $.maxRetries = Codegen.integerProp("maxRetries").output().arg($.maxRetries).env("VAULT_MAX_RETRIES").def(2).getNullable();
            $.skipTlsVerify = Codegen.booleanProp("skipTlsVerify").output().arg($.skipTlsVerify).env("VAULT_SKIP_VERIFY").getNullable();
            return $;
        }
    }

}
