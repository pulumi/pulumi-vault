// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vault.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class NomadSecretBackendState extends com.pulumi.resources.ResourceArgs {

    public static final NomadSecretBackendState Empty = new NomadSecretBackendState();

    /**
     * Specifies the address of the Nomad instance, provided
     * as &#34;protocol://host:port&#34; like &#34;http://127.0.0.1:4646&#34;.
     * 
     */
    @Import(name="address")
    private @Nullable Output<String> address;

    /**
     * @return Specifies the address of the Nomad instance, provided
     * as &#34;protocol://host:port&#34; like &#34;http://127.0.0.1:4646&#34;.
     * 
     */
    public Optional<Output<String>> address() {
        return Optional.ofNullable(this.address);
    }

    /**
     * The unique path this backend should be mounted at. Must
     * not begin or end with a `/`. Defaults to `nomad`.
     * 
     */
    @Import(name="backend")
    private @Nullable Output<String> backend;

    /**
     * @return The unique path this backend should be mounted at. Must
     * not begin or end with a `/`. Defaults to `nomad`.
     * 
     */
    public Optional<Output<String>> backend() {
        return Optional.ofNullable(this.backend);
    }

    /**
     * CA certificate to use when verifying the Nomad server certificate, must be
     * x509 PEM encoded.
     * 
     */
    @Import(name="caCert")
    private @Nullable Output<String> caCert;

    /**
     * @return CA certificate to use when verifying the Nomad server certificate, must be
     * x509 PEM encoded.
     * 
     */
    public Optional<Output<String>> caCert() {
        return Optional.ofNullable(this.caCert);
    }

    /**
     * Client certificate to provide to the Nomad server, must be x509 PEM encoded.
     * 
     */
    @Import(name="clientCert")
    private @Nullable Output<String> clientCert;

    /**
     * @return Client certificate to provide to the Nomad server, must be x509 PEM encoded.
     * 
     */
    public Optional<Output<String>> clientCert() {
        return Optional.ofNullable(this.clientCert);
    }

    /**
     * Client certificate key to provide to the Nomad server, must be x509 PEM encoded.
     * 
     */
    @Import(name="clientKey")
    private @Nullable Output<String> clientKey;

    /**
     * @return Client certificate key to provide to the Nomad server, must be x509 PEM encoded.
     * 
     */
    public Optional<Output<String>> clientKey() {
        return Optional.ofNullable(this.clientKey);
    }

    /**
     * Default lease duration for secrets in seconds.
     * 
     */
    @Import(name="defaultLeaseTtlSeconds")
    private @Nullable Output<Integer> defaultLeaseTtlSeconds;

    /**
     * @return Default lease duration for secrets in seconds.
     * 
     */
    public Optional<Output<Integer>> defaultLeaseTtlSeconds() {
        return Optional.ofNullable(this.defaultLeaseTtlSeconds);
    }

    /**
     * Human-friendly description of the mount for the Active Directory backend.
     * 
     */
    @Import(name="description")
    private @Nullable Output<String> description;

    /**
     * @return Human-friendly description of the mount for the Active Directory backend.
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
     * Mark the secrets engine as local-only. Local engines are not replicated or removed by
     * replication.Tolerance duration to use when checking the last rotation time.
     * 
     */
    @Import(name="local")
    private @Nullable Output<Boolean> local;

    /**
     * @return Mark the secrets engine as local-only. Local engines are not replicated or removed by
     * replication.Tolerance duration to use when checking the last rotation time.
     * 
     */
    public Optional<Output<Boolean>> local() {
        return Optional.ofNullable(this.local);
    }

    /**
     * Maximum possible lease duration for secrets in seconds.
     * 
     */
    @Import(name="maxLeaseTtlSeconds")
    private @Nullable Output<Integer> maxLeaseTtlSeconds;

    /**
     * @return Maximum possible lease duration for secrets in seconds.
     * 
     */
    public Optional<Output<Integer>> maxLeaseTtlSeconds() {
        return Optional.ofNullable(this.maxLeaseTtlSeconds);
    }

    /**
     * Specifies the maximum length to use for the name of the Nomad token
     * generated with Generate Credential. If omitted, 0 is used and ignored, defaulting to the max value allowed
     * by the Nomad version.
     * 
     */
    @Import(name="maxTokenNameLength")
    private @Nullable Output<Integer> maxTokenNameLength;

    /**
     * @return Specifies the maximum length to use for the name of the Nomad token
     * generated with Generate Credential. If omitted, 0 is used and ignored, defaulting to the max value allowed
     * by the Nomad version.
     * 
     */
    public Optional<Output<Integer>> maxTokenNameLength() {
        return Optional.ofNullable(this.maxTokenNameLength);
    }

    /**
     * Maximum possible lease duration for secrets in seconds.
     * 
     */
    @Import(name="maxTtl")
    private @Nullable Output<Integer> maxTtl;

    /**
     * @return Maximum possible lease duration for secrets in seconds.
     * 
     */
    public Optional<Output<Integer>> maxTtl() {
        return Optional.ofNullable(this.maxTtl);
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
     * Specifies the Nomad Management token to use.
     * 
     */
    @Import(name="token")
    private @Nullable Output<String> token;

    /**
     * @return Specifies the Nomad Management token to use.
     * 
     */
    public Optional<Output<String>> token() {
        return Optional.ofNullable(this.token);
    }

    /**
     * Specifies the ttl of the lease for the generated token.
     * 
     */
    @Import(name="ttl")
    private @Nullable Output<Integer> ttl;

    /**
     * @return Specifies the ttl of the lease for the generated token.
     * 
     */
    public Optional<Output<Integer>> ttl() {
        return Optional.ofNullable(this.ttl);
    }

    private NomadSecretBackendState() {}

    private NomadSecretBackendState(NomadSecretBackendState $) {
        this.address = $.address;
        this.backend = $.backend;
        this.caCert = $.caCert;
        this.clientCert = $.clientCert;
        this.clientKey = $.clientKey;
        this.defaultLeaseTtlSeconds = $.defaultLeaseTtlSeconds;
        this.description = $.description;
        this.disableRemount = $.disableRemount;
        this.local = $.local;
        this.maxLeaseTtlSeconds = $.maxLeaseTtlSeconds;
        this.maxTokenNameLength = $.maxTokenNameLength;
        this.maxTtl = $.maxTtl;
        this.namespace = $.namespace;
        this.token = $.token;
        this.ttl = $.ttl;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(NomadSecretBackendState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private NomadSecretBackendState $;

        public Builder() {
            $ = new NomadSecretBackendState();
        }

        public Builder(NomadSecretBackendState defaults) {
            $ = new NomadSecretBackendState(Objects.requireNonNull(defaults));
        }

        /**
         * @param address Specifies the address of the Nomad instance, provided
         * as &#34;protocol://host:port&#34; like &#34;http://127.0.0.1:4646&#34;.
         * 
         * @return builder
         * 
         */
        public Builder address(@Nullable Output<String> address) {
            $.address = address;
            return this;
        }

        /**
         * @param address Specifies the address of the Nomad instance, provided
         * as &#34;protocol://host:port&#34; like &#34;http://127.0.0.1:4646&#34;.
         * 
         * @return builder
         * 
         */
        public Builder address(String address) {
            return address(Output.of(address));
        }

        /**
         * @param backend The unique path this backend should be mounted at. Must
         * not begin or end with a `/`. Defaults to `nomad`.
         * 
         * @return builder
         * 
         */
        public Builder backend(@Nullable Output<String> backend) {
            $.backend = backend;
            return this;
        }

        /**
         * @param backend The unique path this backend should be mounted at. Must
         * not begin or end with a `/`. Defaults to `nomad`.
         * 
         * @return builder
         * 
         */
        public Builder backend(String backend) {
            return backend(Output.of(backend));
        }

        /**
         * @param caCert CA certificate to use when verifying the Nomad server certificate, must be
         * x509 PEM encoded.
         * 
         * @return builder
         * 
         */
        public Builder caCert(@Nullable Output<String> caCert) {
            $.caCert = caCert;
            return this;
        }

        /**
         * @param caCert CA certificate to use when verifying the Nomad server certificate, must be
         * x509 PEM encoded.
         * 
         * @return builder
         * 
         */
        public Builder caCert(String caCert) {
            return caCert(Output.of(caCert));
        }

        /**
         * @param clientCert Client certificate to provide to the Nomad server, must be x509 PEM encoded.
         * 
         * @return builder
         * 
         */
        public Builder clientCert(@Nullable Output<String> clientCert) {
            $.clientCert = clientCert;
            return this;
        }

        /**
         * @param clientCert Client certificate to provide to the Nomad server, must be x509 PEM encoded.
         * 
         * @return builder
         * 
         */
        public Builder clientCert(String clientCert) {
            return clientCert(Output.of(clientCert));
        }

        /**
         * @param clientKey Client certificate key to provide to the Nomad server, must be x509 PEM encoded.
         * 
         * @return builder
         * 
         */
        public Builder clientKey(@Nullable Output<String> clientKey) {
            $.clientKey = clientKey;
            return this;
        }

        /**
         * @param clientKey Client certificate key to provide to the Nomad server, must be x509 PEM encoded.
         * 
         * @return builder
         * 
         */
        public Builder clientKey(String clientKey) {
            return clientKey(Output.of(clientKey));
        }

        /**
         * @param defaultLeaseTtlSeconds Default lease duration for secrets in seconds.
         * 
         * @return builder
         * 
         */
        public Builder defaultLeaseTtlSeconds(@Nullable Output<Integer> defaultLeaseTtlSeconds) {
            $.defaultLeaseTtlSeconds = defaultLeaseTtlSeconds;
            return this;
        }

        /**
         * @param defaultLeaseTtlSeconds Default lease duration for secrets in seconds.
         * 
         * @return builder
         * 
         */
        public Builder defaultLeaseTtlSeconds(Integer defaultLeaseTtlSeconds) {
            return defaultLeaseTtlSeconds(Output.of(defaultLeaseTtlSeconds));
        }

        /**
         * @param description Human-friendly description of the mount for the Active Directory backend.
         * 
         * @return builder
         * 
         */
        public Builder description(@Nullable Output<String> description) {
            $.description = description;
            return this;
        }

        /**
         * @param description Human-friendly description of the mount for the Active Directory backend.
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
         * @param local Mark the secrets engine as local-only. Local engines are not replicated or removed by
         * replication.Tolerance duration to use when checking the last rotation time.
         * 
         * @return builder
         * 
         */
        public Builder local(@Nullable Output<Boolean> local) {
            $.local = local;
            return this;
        }

        /**
         * @param local Mark the secrets engine as local-only. Local engines are not replicated or removed by
         * replication.Tolerance duration to use when checking the last rotation time.
         * 
         * @return builder
         * 
         */
        public Builder local(Boolean local) {
            return local(Output.of(local));
        }

        /**
         * @param maxLeaseTtlSeconds Maximum possible lease duration for secrets in seconds.
         * 
         * @return builder
         * 
         */
        public Builder maxLeaseTtlSeconds(@Nullable Output<Integer> maxLeaseTtlSeconds) {
            $.maxLeaseTtlSeconds = maxLeaseTtlSeconds;
            return this;
        }

        /**
         * @param maxLeaseTtlSeconds Maximum possible lease duration for secrets in seconds.
         * 
         * @return builder
         * 
         */
        public Builder maxLeaseTtlSeconds(Integer maxLeaseTtlSeconds) {
            return maxLeaseTtlSeconds(Output.of(maxLeaseTtlSeconds));
        }

        /**
         * @param maxTokenNameLength Specifies the maximum length to use for the name of the Nomad token
         * generated with Generate Credential. If omitted, 0 is used and ignored, defaulting to the max value allowed
         * by the Nomad version.
         * 
         * @return builder
         * 
         */
        public Builder maxTokenNameLength(@Nullable Output<Integer> maxTokenNameLength) {
            $.maxTokenNameLength = maxTokenNameLength;
            return this;
        }

        /**
         * @param maxTokenNameLength Specifies the maximum length to use for the name of the Nomad token
         * generated with Generate Credential. If omitted, 0 is used and ignored, defaulting to the max value allowed
         * by the Nomad version.
         * 
         * @return builder
         * 
         */
        public Builder maxTokenNameLength(Integer maxTokenNameLength) {
            return maxTokenNameLength(Output.of(maxTokenNameLength));
        }

        /**
         * @param maxTtl Maximum possible lease duration for secrets in seconds.
         * 
         * @return builder
         * 
         */
        public Builder maxTtl(@Nullable Output<Integer> maxTtl) {
            $.maxTtl = maxTtl;
            return this;
        }

        /**
         * @param maxTtl Maximum possible lease duration for secrets in seconds.
         * 
         * @return builder
         * 
         */
        public Builder maxTtl(Integer maxTtl) {
            return maxTtl(Output.of(maxTtl));
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
         * @param token Specifies the Nomad Management token to use.
         * 
         * @return builder
         * 
         */
        public Builder token(@Nullable Output<String> token) {
            $.token = token;
            return this;
        }

        /**
         * @param token Specifies the Nomad Management token to use.
         * 
         * @return builder
         * 
         */
        public Builder token(String token) {
            return token(Output.of(token));
        }

        /**
         * @param ttl Specifies the ttl of the lease for the generated token.
         * 
         * @return builder
         * 
         */
        public Builder ttl(@Nullable Output<Integer> ttl) {
            $.ttl = ttl;
            return this;
        }

        /**
         * @param ttl Specifies the ttl of the lease for the generated token.
         * 
         * @return builder
         * 
         */
        public Builder ttl(Integer ttl) {
            return ttl(Output.of(ttl));
        }

        public NomadSecretBackendState build() {
            return $;
        }
    }

}
