// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vault.gcp.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.vault.gcp.inputs.AuthBackendCustomEndpointArgs;
import com.pulumi.vault.gcp.inputs.AuthBackendTuneArgs;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class AuthBackendState extends com.pulumi.resources.ResourceArgs {

    public static final AuthBackendState Empty = new AuthBackendState();

    /**
     * The mount accessor related to the auth mount. It is useful for integration with [Identity Secrets Engine](https://www.vaultproject.io/docs/secrets/identity/index.html).
     * 
     */
    @Import(name="accessor")
    private @Nullable Output<String> accessor;

    /**
     * @return The mount accessor related to the auth mount. It is useful for integration with [Identity Secrets Engine](https://www.vaultproject.io/docs/secrets/identity/index.html).
     * 
     */
    public Optional<Output<String>> accessor() {
        return Optional.ofNullable(this.accessor);
    }

    /**
     * The clients email associated with the credentials
     * 
     */
    @Import(name="clientEmail")
    private @Nullable Output<String> clientEmail;

    /**
     * @return The clients email associated with the credentials
     * 
     */
    public Optional<Output<String>> clientEmail() {
        return Optional.ofNullable(this.clientEmail);
    }

    /**
     * The Client ID of the credentials
     * 
     */
    @Import(name="clientId")
    private @Nullable Output<String> clientId;

    /**
     * @return The Client ID of the credentials
     * 
     */
    public Optional<Output<String>> clientId() {
        return Optional.ofNullable(this.clientId);
    }

    /**
     * A JSON string containing the contents of a GCP credentials file. If this value is empty, Vault will try to use Application Default Credentials from the machine on which the Vault server is running.
     * 
     */
    @Import(name="credentials")
    private @Nullable Output<String> credentials;

    /**
     * @return A JSON string containing the contents of a GCP credentials file. If this value is empty, Vault will try to use Application Default Credentials from the machine on which the Vault server is running.
     * 
     */
    public Optional<Output<String>> credentials() {
        return Optional.ofNullable(this.credentials);
    }

    /**
     * Specifies overrides to
     * [service endpoints](https://cloud.google.com/apis/design/glossary#api_service_endpoint)
     * used when making API requests. This allows specific requests made during authentication
     * to target alternative service endpoints for use in [Private Google Access](https://cloud.google.com/vpc/docs/configure-private-google-access)
     * environments. Requires Vault 1.11+.
     * 
     * Overrides are set at the subdomain level using the following keys:
     * 
     */
    @Import(name="customEndpoint")
    private @Nullable Output<AuthBackendCustomEndpointArgs> customEndpoint;

    /**
     * @return Specifies overrides to
     * [service endpoints](https://cloud.google.com/apis/design/glossary#api_service_endpoint)
     * used when making API requests. This allows specific requests made during authentication
     * to target alternative service endpoints for use in [Private Google Access](https://cloud.google.com/vpc/docs/configure-private-google-access)
     * environments. Requires Vault 1.11+.
     * 
     * Overrides are set at the subdomain level using the following keys:
     * 
     */
    public Optional<Output<AuthBackendCustomEndpointArgs>> customEndpoint() {
        return Optional.ofNullable(this.customEndpoint);
    }

    /**
     * A description of the auth method.
     * 
     */
    @Import(name="description")
    private @Nullable Output<String> description;

    /**
     * @return A description of the auth method.
     * 
     */
    public Optional<Output<String>> description() {
        return Optional.ofNullable(this.description);
    }

    /**
     * Cancels all upcoming rotations of the root credential until unset. Requires Vault Enterprise 1.19+.
     * 
     */
    @Import(name="disableAutomatedRotation")
    private @Nullable Output<Boolean> disableAutomatedRotation;

    /**
     * @return Cancels all upcoming rotations of the root credential until unset. Requires Vault Enterprise 1.19+.
     * 
     */
    public Optional<Output<Boolean>> disableAutomatedRotation() {
        return Optional.ofNullable(this.disableAutomatedRotation);
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
     * The audience claim value for plugin identity
     * tokens. Must match an allowed audience configured for the target [Workload Identity Pool](https://cloud.google.com/iam/docs/workload-identity-federation-with-other-providers#prepare).
     * Mutually exclusive with `credentials`.  Requires Vault 1.17+. *Available only for Vault Enterprise*.
     * 
     */
    @Import(name="identityTokenAudience")
    private @Nullable Output<String> identityTokenAudience;

    /**
     * @return The audience claim value for plugin identity
     * tokens. Must match an allowed audience configured for the target [Workload Identity Pool](https://cloud.google.com/iam/docs/workload-identity-federation-with-other-providers#prepare).
     * Mutually exclusive with `credentials`.  Requires Vault 1.17+. *Available only for Vault Enterprise*.
     * 
     */
    public Optional<Output<String>> identityTokenAudience() {
        return Optional.ofNullable(this.identityTokenAudience);
    }

    /**
     * The key to use for signing plugin identity
     * tokens. Requires Vault 1.17+. *Available only for Vault Enterprise*.
     * 
     */
    @Import(name="identityTokenKey")
    private @Nullable Output<String> identityTokenKey;

    /**
     * @return The key to use for signing plugin identity
     * tokens. Requires Vault 1.17+. *Available only for Vault Enterprise*.
     * 
     */
    public Optional<Output<String>> identityTokenKey() {
        return Optional.ofNullable(this.identityTokenKey);
    }

    /**
     * The TTL of generated tokens.
     * 
     */
    @Import(name="identityTokenTtl")
    private @Nullable Output<Integer> identityTokenTtl;

    /**
     * @return The TTL of generated tokens.
     * 
     */
    public Optional<Output<Integer>> identityTokenTtl() {
        return Optional.ofNullable(this.identityTokenTtl);
    }

    /**
     * Specifies if the auth method is local only.
     * 
     */
    @Import(name="local")
    private @Nullable Output<Boolean> local;

    /**
     * @return Specifies if the auth method is local only.
     * 
     */
    public Optional<Output<Boolean>> local() {
        return Optional.ofNullable(this.local);
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
     * The path to mount the auth method — this defaults to &#39;gcp&#39;.
     * 
     */
    @Import(name="path")
    private @Nullable Output<String> path;

    /**
     * @return The path to mount the auth method — this defaults to &#39;gcp&#39;.
     * 
     */
    public Optional<Output<String>> path() {
        return Optional.ofNullable(this.path);
    }

    /**
     * The ID of the private key from the credentials
     * 
     */
    @Import(name="privateKeyId")
    private @Nullable Output<String> privateKeyId;

    /**
     * @return The ID of the private key from the credentials
     * 
     */
    public Optional<Output<String>> privateKeyId() {
        return Optional.ofNullable(this.privateKeyId);
    }

    /**
     * The GCP Project ID
     * 
     */
    @Import(name="projectId")
    private @Nullable Output<String> projectId;

    /**
     * @return The GCP Project ID
     * 
     */
    public Optional<Output<String>> projectId() {
        return Optional.ofNullable(this.projectId);
    }

    /**
     * The amount of time in seconds Vault should wait before rotating the root credential.
     * A zero value tells Vault not to rotate the root credential. The minimum rotation period is 10 seconds. Requires Vault Enterprise 1.19+.
     * 
     */
    @Import(name="rotationPeriod")
    private @Nullable Output<Integer> rotationPeriod;

    /**
     * @return The amount of time in seconds Vault should wait before rotating the root credential.
     * A zero value tells Vault not to rotate the root credential. The minimum rotation period is 10 seconds. Requires Vault Enterprise 1.19+.
     * 
     */
    public Optional<Output<Integer>> rotationPeriod() {
        return Optional.ofNullable(this.rotationPeriod);
    }

    /**
     * The schedule, in [cron-style time format](https://en.wikipedia.org/wiki/Cron),
     * defining the schedule on which Vault should rotate the root token. Requires Vault Enterprise 1.19+.
     * 
     */
    @Import(name="rotationSchedule")
    private @Nullable Output<String> rotationSchedule;

    /**
     * @return The schedule, in [cron-style time format](https://en.wikipedia.org/wiki/Cron),
     * defining the schedule on which Vault should rotate the root token. Requires Vault Enterprise 1.19+.
     * 
     */
    public Optional<Output<String>> rotationSchedule() {
        return Optional.ofNullable(this.rotationSchedule);
    }

    /**
     * The maximum amount of time in seconds allowed to complete
     * a rotation when a scheduled token rotation occurs. The default rotation window is
     * unbound and the minimum allowable window is `3600`. Requires Vault Enterprise 1.19+.
     * 
     */
    @Import(name="rotationWindow")
    private @Nullable Output<Integer> rotationWindow;

    /**
     * @return The maximum amount of time in seconds allowed to complete
     * a rotation when a scheduled token rotation occurs. The default rotation window is
     * unbound and the minimum allowable window is `3600`. Requires Vault Enterprise 1.19+.
     * 
     */
    public Optional<Output<Integer>> rotationWindow() {
        return Optional.ofNullable(this.rotationWindow);
    }

    /**
     * Service Account to impersonate for plugin workload identity federation.
     * Required with `identity_token_audience`. Requires Vault 1.17+. *Available only for Vault Enterprise*.
     * 
     */
    @Import(name="serviceAccountEmail")
    private @Nullable Output<String> serviceAccountEmail;

    /**
     * @return Service Account to impersonate for plugin workload identity federation.
     * Required with `identity_token_audience`. Requires Vault 1.17+. *Available only for Vault Enterprise*.
     * 
     */
    public Optional<Output<String>> serviceAccountEmail() {
        return Optional.ofNullable(this.serviceAccountEmail);
    }

    /**
     * Extra configuration block. Structure is documented below.
     * 
     * The `tune` block is used to tune the auth backend:
     * 
     */
    @Import(name="tune")
    private @Nullable Output<AuthBackendTuneArgs> tune;

    /**
     * @return Extra configuration block. Structure is documented below.
     * 
     * The `tune` block is used to tune the auth backend:
     * 
     */
    public Optional<Output<AuthBackendTuneArgs>> tune() {
        return Optional.ofNullable(this.tune);
    }

    private AuthBackendState() {}

    private AuthBackendState(AuthBackendState $) {
        this.accessor = $.accessor;
        this.clientEmail = $.clientEmail;
        this.clientId = $.clientId;
        this.credentials = $.credentials;
        this.customEndpoint = $.customEndpoint;
        this.description = $.description;
        this.disableAutomatedRotation = $.disableAutomatedRotation;
        this.disableRemount = $.disableRemount;
        this.identityTokenAudience = $.identityTokenAudience;
        this.identityTokenKey = $.identityTokenKey;
        this.identityTokenTtl = $.identityTokenTtl;
        this.local = $.local;
        this.namespace = $.namespace;
        this.path = $.path;
        this.privateKeyId = $.privateKeyId;
        this.projectId = $.projectId;
        this.rotationPeriod = $.rotationPeriod;
        this.rotationSchedule = $.rotationSchedule;
        this.rotationWindow = $.rotationWindow;
        this.serviceAccountEmail = $.serviceAccountEmail;
        this.tune = $.tune;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(AuthBackendState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private AuthBackendState $;

        public Builder() {
            $ = new AuthBackendState();
        }

        public Builder(AuthBackendState defaults) {
            $ = new AuthBackendState(Objects.requireNonNull(defaults));
        }

        /**
         * @param accessor The mount accessor related to the auth mount. It is useful for integration with [Identity Secrets Engine](https://www.vaultproject.io/docs/secrets/identity/index.html).
         * 
         * @return builder
         * 
         */
        public Builder accessor(@Nullable Output<String> accessor) {
            $.accessor = accessor;
            return this;
        }

        /**
         * @param accessor The mount accessor related to the auth mount. It is useful for integration with [Identity Secrets Engine](https://www.vaultproject.io/docs/secrets/identity/index.html).
         * 
         * @return builder
         * 
         */
        public Builder accessor(String accessor) {
            return accessor(Output.of(accessor));
        }

        /**
         * @param clientEmail The clients email associated with the credentials
         * 
         * @return builder
         * 
         */
        public Builder clientEmail(@Nullable Output<String> clientEmail) {
            $.clientEmail = clientEmail;
            return this;
        }

        /**
         * @param clientEmail The clients email associated with the credentials
         * 
         * @return builder
         * 
         */
        public Builder clientEmail(String clientEmail) {
            return clientEmail(Output.of(clientEmail));
        }

        /**
         * @param clientId The Client ID of the credentials
         * 
         * @return builder
         * 
         */
        public Builder clientId(@Nullable Output<String> clientId) {
            $.clientId = clientId;
            return this;
        }

        /**
         * @param clientId The Client ID of the credentials
         * 
         * @return builder
         * 
         */
        public Builder clientId(String clientId) {
            return clientId(Output.of(clientId));
        }

        /**
         * @param credentials A JSON string containing the contents of a GCP credentials file. If this value is empty, Vault will try to use Application Default Credentials from the machine on which the Vault server is running.
         * 
         * @return builder
         * 
         */
        public Builder credentials(@Nullable Output<String> credentials) {
            $.credentials = credentials;
            return this;
        }

        /**
         * @param credentials A JSON string containing the contents of a GCP credentials file. If this value is empty, Vault will try to use Application Default Credentials from the machine on which the Vault server is running.
         * 
         * @return builder
         * 
         */
        public Builder credentials(String credentials) {
            return credentials(Output.of(credentials));
        }

        /**
         * @param customEndpoint Specifies overrides to
         * [service endpoints](https://cloud.google.com/apis/design/glossary#api_service_endpoint)
         * used when making API requests. This allows specific requests made during authentication
         * to target alternative service endpoints for use in [Private Google Access](https://cloud.google.com/vpc/docs/configure-private-google-access)
         * environments. Requires Vault 1.11+.
         * 
         * Overrides are set at the subdomain level using the following keys:
         * 
         * @return builder
         * 
         */
        public Builder customEndpoint(@Nullable Output<AuthBackendCustomEndpointArgs> customEndpoint) {
            $.customEndpoint = customEndpoint;
            return this;
        }

        /**
         * @param customEndpoint Specifies overrides to
         * [service endpoints](https://cloud.google.com/apis/design/glossary#api_service_endpoint)
         * used when making API requests. This allows specific requests made during authentication
         * to target alternative service endpoints for use in [Private Google Access](https://cloud.google.com/vpc/docs/configure-private-google-access)
         * environments. Requires Vault 1.11+.
         * 
         * Overrides are set at the subdomain level using the following keys:
         * 
         * @return builder
         * 
         */
        public Builder customEndpoint(AuthBackendCustomEndpointArgs customEndpoint) {
            return customEndpoint(Output.of(customEndpoint));
        }

        /**
         * @param description A description of the auth method.
         * 
         * @return builder
         * 
         */
        public Builder description(@Nullable Output<String> description) {
            $.description = description;
            return this;
        }

        /**
         * @param description A description of the auth method.
         * 
         * @return builder
         * 
         */
        public Builder description(String description) {
            return description(Output.of(description));
        }

        /**
         * @param disableAutomatedRotation Cancels all upcoming rotations of the root credential until unset. Requires Vault Enterprise 1.19+.
         * 
         * @return builder
         * 
         */
        public Builder disableAutomatedRotation(@Nullable Output<Boolean> disableAutomatedRotation) {
            $.disableAutomatedRotation = disableAutomatedRotation;
            return this;
        }

        /**
         * @param disableAutomatedRotation Cancels all upcoming rotations of the root credential until unset. Requires Vault Enterprise 1.19+.
         * 
         * @return builder
         * 
         */
        public Builder disableAutomatedRotation(Boolean disableAutomatedRotation) {
            return disableAutomatedRotation(Output.of(disableAutomatedRotation));
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
         * @param identityTokenAudience The audience claim value for plugin identity
         * tokens. Must match an allowed audience configured for the target [Workload Identity Pool](https://cloud.google.com/iam/docs/workload-identity-federation-with-other-providers#prepare).
         * Mutually exclusive with `credentials`.  Requires Vault 1.17+. *Available only for Vault Enterprise*.
         * 
         * @return builder
         * 
         */
        public Builder identityTokenAudience(@Nullable Output<String> identityTokenAudience) {
            $.identityTokenAudience = identityTokenAudience;
            return this;
        }

        /**
         * @param identityTokenAudience The audience claim value for plugin identity
         * tokens. Must match an allowed audience configured for the target [Workload Identity Pool](https://cloud.google.com/iam/docs/workload-identity-federation-with-other-providers#prepare).
         * Mutually exclusive with `credentials`.  Requires Vault 1.17+. *Available only for Vault Enterprise*.
         * 
         * @return builder
         * 
         */
        public Builder identityTokenAudience(String identityTokenAudience) {
            return identityTokenAudience(Output.of(identityTokenAudience));
        }

        /**
         * @param identityTokenKey The key to use for signing plugin identity
         * tokens. Requires Vault 1.17+. *Available only for Vault Enterprise*.
         * 
         * @return builder
         * 
         */
        public Builder identityTokenKey(@Nullable Output<String> identityTokenKey) {
            $.identityTokenKey = identityTokenKey;
            return this;
        }

        /**
         * @param identityTokenKey The key to use for signing plugin identity
         * tokens. Requires Vault 1.17+. *Available only for Vault Enterprise*.
         * 
         * @return builder
         * 
         */
        public Builder identityTokenKey(String identityTokenKey) {
            return identityTokenKey(Output.of(identityTokenKey));
        }

        /**
         * @param identityTokenTtl The TTL of generated tokens.
         * 
         * @return builder
         * 
         */
        public Builder identityTokenTtl(@Nullable Output<Integer> identityTokenTtl) {
            $.identityTokenTtl = identityTokenTtl;
            return this;
        }

        /**
         * @param identityTokenTtl The TTL of generated tokens.
         * 
         * @return builder
         * 
         */
        public Builder identityTokenTtl(Integer identityTokenTtl) {
            return identityTokenTtl(Output.of(identityTokenTtl));
        }

        /**
         * @param local Specifies if the auth method is local only.
         * 
         * @return builder
         * 
         */
        public Builder local(@Nullable Output<Boolean> local) {
            $.local = local;
            return this;
        }

        /**
         * @param local Specifies if the auth method is local only.
         * 
         * @return builder
         * 
         */
        public Builder local(Boolean local) {
            return local(Output.of(local));
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
         * @param path The path to mount the auth method — this defaults to &#39;gcp&#39;.
         * 
         * @return builder
         * 
         */
        public Builder path(@Nullable Output<String> path) {
            $.path = path;
            return this;
        }

        /**
         * @param path The path to mount the auth method — this defaults to &#39;gcp&#39;.
         * 
         * @return builder
         * 
         */
        public Builder path(String path) {
            return path(Output.of(path));
        }

        /**
         * @param privateKeyId The ID of the private key from the credentials
         * 
         * @return builder
         * 
         */
        public Builder privateKeyId(@Nullable Output<String> privateKeyId) {
            $.privateKeyId = privateKeyId;
            return this;
        }

        /**
         * @param privateKeyId The ID of the private key from the credentials
         * 
         * @return builder
         * 
         */
        public Builder privateKeyId(String privateKeyId) {
            return privateKeyId(Output.of(privateKeyId));
        }

        /**
         * @param projectId The GCP Project ID
         * 
         * @return builder
         * 
         */
        public Builder projectId(@Nullable Output<String> projectId) {
            $.projectId = projectId;
            return this;
        }

        /**
         * @param projectId The GCP Project ID
         * 
         * @return builder
         * 
         */
        public Builder projectId(String projectId) {
            return projectId(Output.of(projectId));
        }

        /**
         * @param rotationPeriod The amount of time in seconds Vault should wait before rotating the root credential.
         * A zero value tells Vault not to rotate the root credential. The minimum rotation period is 10 seconds. Requires Vault Enterprise 1.19+.
         * 
         * @return builder
         * 
         */
        public Builder rotationPeriod(@Nullable Output<Integer> rotationPeriod) {
            $.rotationPeriod = rotationPeriod;
            return this;
        }

        /**
         * @param rotationPeriod The amount of time in seconds Vault should wait before rotating the root credential.
         * A zero value tells Vault not to rotate the root credential. The minimum rotation period is 10 seconds. Requires Vault Enterprise 1.19+.
         * 
         * @return builder
         * 
         */
        public Builder rotationPeriod(Integer rotationPeriod) {
            return rotationPeriod(Output.of(rotationPeriod));
        }

        /**
         * @param rotationSchedule The schedule, in [cron-style time format](https://en.wikipedia.org/wiki/Cron),
         * defining the schedule on which Vault should rotate the root token. Requires Vault Enterprise 1.19+.
         * 
         * @return builder
         * 
         */
        public Builder rotationSchedule(@Nullable Output<String> rotationSchedule) {
            $.rotationSchedule = rotationSchedule;
            return this;
        }

        /**
         * @param rotationSchedule The schedule, in [cron-style time format](https://en.wikipedia.org/wiki/Cron),
         * defining the schedule on which Vault should rotate the root token. Requires Vault Enterprise 1.19+.
         * 
         * @return builder
         * 
         */
        public Builder rotationSchedule(String rotationSchedule) {
            return rotationSchedule(Output.of(rotationSchedule));
        }

        /**
         * @param rotationWindow The maximum amount of time in seconds allowed to complete
         * a rotation when a scheduled token rotation occurs. The default rotation window is
         * unbound and the minimum allowable window is `3600`. Requires Vault Enterprise 1.19+.
         * 
         * @return builder
         * 
         */
        public Builder rotationWindow(@Nullable Output<Integer> rotationWindow) {
            $.rotationWindow = rotationWindow;
            return this;
        }

        /**
         * @param rotationWindow The maximum amount of time in seconds allowed to complete
         * a rotation when a scheduled token rotation occurs. The default rotation window is
         * unbound and the minimum allowable window is `3600`. Requires Vault Enterprise 1.19+.
         * 
         * @return builder
         * 
         */
        public Builder rotationWindow(Integer rotationWindow) {
            return rotationWindow(Output.of(rotationWindow));
        }

        /**
         * @param serviceAccountEmail Service Account to impersonate for plugin workload identity federation.
         * Required with `identity_token_audience`. Requires Vault 1.17+. *Available only for Vault Enterprise*.
         * 
         * @return builder
         * 
         */
        public Builder serviceAccountEmail(@Nullable Output<String> serviceAccountEmail) {
            $.serviceAccountEmail = serviceAccountEmail;
            return this;
        }

        /**
         * @param serviceAccountEmail Service Account to impersonate for plugin workload identity federation.
         * Required with `identity_token_audience`. Requires Vault 1.17+. *Available only for Vault Enterprise*.
         * 
         * @return builder
         * 
         */
        public Builder serviceAccountEmail(String serviceAccountEmail) {
            return serviceAccountEmail(Output.of(serviceAccountEmail));
        }

        /**
         * @param tune Extra configuration block. Structure is documented below.
         * 
         * The `tune` block is used to tune the auth backend:
         * 
         * @return builder
         * 
         */
        public Builder tune(@Nullable Output<AuthBackendTuneArgs> tune) {
            $.tune = tune;
            return this;
        }

        /**
         * @param tune Extra configuration block. Structure is documented below.
         * 
         * The `tune` block is used to tune the auth backend:
         * 
         * @return builder
         * 
         */
        public Builder tune(AuthBackendTuneArgs tune) {
            return tune(Output.of(tune));
        }

        public AuthBackendState build() {
            return $;
        }
    }

}
