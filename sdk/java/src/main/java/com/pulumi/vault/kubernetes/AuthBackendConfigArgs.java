// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vault.kubernetes;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class AuthBackendConfigArgs extends com.pulumi.resources.ResourceArgs {

    public static final AuthBackendConfigArgs Empty = new AuthBackendConfigArgs();

    /**
     * Unique name of the kubernetes backend to configure.
     * 
     */
    @Import(name="backend")
    private @Nullable Output<String> backend;

    /**
     * @return Unique name of the kubernetes backend to configure.
     * 
     */
    public Optional<Output<String>> backend() {
        return Optional.ofNullable(this.backend);
    }

    /**
     * Disable JWT issuer validation. Allows to skip ISS validation. Requires Vault `v1.5.4+` or Vault auth kubernetes plugin `v0.7.1+`
     * 
     */
    @Import(name="disableIssValidation")
    private @Nullable Output<Boolean> disableIssValidation;

    /**
     * @return Disable JWT issuer validation. Allows to skip ISS validation. Requires Vault `v1.5.4+` or Vault auth kubernetes plugin `v0.7.1+`
     * 
     */
    public Optional<Output<Boolean>> disableIssValidation() {
        return Optional.ofNullable(this.disableIssValidation);
    }

    /**
     * Disable defaulting to the local CA cert and service account JWT when running in a Kubernetes pod. Requires Vault `v1.5.4+` or Vault auth kubernetes plugin `v0.7.1+`
     * 
     */
    @Import(name="disableLocalCaJwt")
    private @Nullable Output<Boolean> disableLocalCaJwt;

    /**
     * @return Disable defaulting to the local CA cert and service account JWT when running in a Kubernetes pod. Requires Vault `v1.5.4+` or Vault auth kubernetes plugin `v0.7.1+`
     * 
     */
    public Optional<Output<Boolean>> disableLocalCaJwt() {
        return Optional.ofNullable(this.disableLocalCaJwt);
    }

    /**
     * JWT issuer. If no issuer is specified, `kubernetes.io/serviceaccount` will be used as the default issuer.
     * 
     */
    @Import(name="issuer")
    private @Nullable Output<String> issuer;

    /**
     * @return JWT issuer. If no issuer is specified, `kubernetes.io/serviceaccount` will be used as the default issuer.
     * 
     */
    public Optional<Output<String>> issuer() {
        return Optional.ofNullable(this.issuer);
    }

    /**
     * PEM encoded CA cert for use by the TLS client used to talk with the Kubernetes API.
     * 
     */
    @Import(name="kubernetesCaCert")
    private @Nullable Output<String> kubernetesCaCert;

    /**
     * @return PEM encoded CA cert for use by the TLS client used to talk with the Kubernetes API.
     * 
     */
    public Optional<Output<String>> kubernetesCaCert() {
        return Optional.ofNullable(this.kubernetesCaCert);
    }

    /**
     * Host must be a host string, a host:port pair, or a URL to the base of the Kubernetes API server.
     * 
     */
    @Import(name="kubernetesHost", required=true)
    private Output<String> kubernetesHost;

    /**
     * @return Host must be a host string, a host:port pair, or a URL to the base of the Kubernetes API server.
     * 
     */
    public Output<String> kubernetesHost() {
        return this.kubernetesHost;
    }

    /**
     * The namespace to provision the resource in.
     * The value should not contain leading or trailing forward slashes.
     * The `namespace` is always relative to the provider&#39;s configured namespace.
     * *Available only for Vault Enterprise*.
     * 
     */
    @Import(name="namespace")
    private @Nullable Output<String> namespace;

    /**
     * @return The namespace to provision the resource in.
     * The value should not contain leading or trailing forward slashes.
     * The `namespace` is always relative to the provider&#39;s configured namespace.
     * *Available only for Vault Enterprise*.
     * 
     */
    public Optional<Output<String>> namespace() {
        return Optional.ofNullable(this.namespace);
    }

    /**
     * List of PEM-formatted public keys or certificates used to verify the signatures of Kubernetes service account JWTs. If a certificate is given, its public key will be extracted. Not every installation of Kubernetes exposes these keys.
     * 
     */
    @Import(name="pemKeys")
    private @Nullable Output<List<String>> pemKeys;

    /**
     * @return List of PEM-formatted public keys or certificates used to verify the signatures of Kubernetes service account JWTs. If a certificate is given, its public key will be extracted. Not every installation of Kubernetes exposes these keys.
     * 
     */
    public Optional<Output<List<String>>> pemKeys() {
        return Optional.ofNullable(this.pemKeys);
    }

    /**
     * A service account JWT (or other token) used as a bearer token to access the TokenReview API to validate other JWTs during login. If not set the JWT used for login will be used to access the API.
     * 
     */
    @Import(name="tokenReviewerJwt")
    private @Nullable Output<String> tokenReviewerJwt;

    /**
     * @return A service account JWT (or other token) used as a bearer token to access the TokenReview API to validate other JWTs during login. If not set the JWT used for login will be used to access the API.
     * 
     */
    public Optional<Output<String>> tokenReviewerJwt() {
        return Optional.ofNullable(this.tokenReviewerJwt);
    }

    private AuthBackendConfigArgs() {}

    private AuthBackendConfigArgs(AuthBackendConfigArgs $) {
        this.backend = $.backend;
        this.disableIssValidation = $.disableIssValidation;
        this.disableLocalCaJwt = $.disableLocalCaJwt;
        this.issuer = $.issuer;
        this.kubernetesCaCert = $.kubernetesCaCert;
        this.kubernetesHost = $.kubernetesHost;
        this.namespace = $.namespace;
        this.pemKeys = $.pemKeys;
        this.tokenReviewerJwt = $.tokenReviewerJwt;
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
         * @param backend Unique name of the kubernetes backend to configure.
         * 
         * @return builder
         * 
         */
        public Builder backend(@Nullable Output<String> backend) {
            $.backend = backend;
            return this;
        }

        /**
         * @param backend Unique name of the kubernetes backend to configure.
         * 
         * @return builder
         * 
         */
        public Builder backend(String backend) {
            return backend(Output.of(backend));
        }

        /**
         * @param disableIssValidation Disable JWT issuer validation. Allows to skip ISS validation. Requires Vault `v1.5.4+` or Vault auth kubernetes plugin `v0.7.1+`
         * 
         * @return builder
         * 
         */
        public Builder disableIssValidation(@Nullable Output<Boolean> disableIssValidation) {
            $.disableIssValidation = disableIssValidation;
            return this;
        }

        /**
         * @param disableIssValidation Disable JWT issuer validation. Allows to skip ISS validation. Requires Vault `v1.5.4+` or Vault auth kubernetes plugin `v0.7.1+`
         * 
         * @return builder
         * 
         */
        public Builder disableIssValidation(Boolean disableIssValidation) {
            return disableIssValidation(Output.of(disableIssValidation));
        }

        /**
         * @param disableLocalCaJwt Disable defaulting to the local CA cert and service account JWT when running in a Kubernetes pod. Requires Vault `v1.5.4+` or Vault auth kubernetes plugin `v0.7.1+`
         * 
         * @return builder
         * 
         */
        public Builder disableLocalCaJwt(@Nullable Output<Boolean> disableLocalCaJwt) {
            $.disableLocalCaJwt = disableLocalCaJwt;
            return this;
        }

        /**
         * @param disableLocalCaJwt Disable defaulting to the local CA cert and service account JWT when running in a Kubernetes pod. Requires Vault `v1.5.4+` or Vault auth kubernetes plugin `v0.7.1+`
         * 
         * @return builder
         * 
         */
        public Builder disableLocalCaJwt(Boolean disableLocalCaJwt) {
            return disableLocalCaJwt(Output.of(disableLocalCaJwt));
        }

        /**
         * @param issuer JWT issuer. If no issuer is specified, `kubernetes.io/serviceaccount` will be used as the default issuer.
         * 
         * @return builder
         * 
         */
        public Builder issuer(@Nullable Output<String> issuer) {
            $.issuer = issuer;
            return this;
        }

        /**
         * @param issuer JWT issuer. If no issuer is specified, `kubernetes.io/serviceaccount` will be used as the default issuer.
         * 
         * @return builder
         * 
         */
        public Builder issuer(String issuer) {
            return issuer(Output.of(issuer));
        }

        /**
         * @param kubernetesCaCert PEM encoded CA cert for use by the TLS client used to talk with the Kubernetes API.
         * 
         * @return builder
         * 
         */
        public Builder kubernetesCaCert(@Nullable Output<String> kubernetesCaCert) {
            $.kubernetesCaCert = kubernetesCaCert;
            return this;
        }

        /**
         * @param kubernetesCaCert PEM encoded CA cert for use by the TLS client used to talk with the Kubernetes API.
         * 
         * @return builder
         * 
         */
        public Builder kubernetesCaCert(String kubernetesCaCert) {
            return kubernetesCaCert(Output.of(kubernetesCaCert));
        }

        /**
         * @param kubernetesHost Host must be a host string, a host:port pair, or a URL to the base of the Kubernetes API server.
         * 
         * @return builder
         * 
         */
        public Builder kubernetesHost(Output<String> kubernetesHost) {
            $.kubernetesHost = kubernetesHost;
            return this;
        }

        /**
         * @param kubernetesHost Host must be a host string, a host:port pair, or a URL to the base of the Kubernetes API server.
         * 
         * @return builder
         * 
         */
        public Builder kubernetesHost(String kubernetesHost) {
            return kubernetesHost(Output.of(kubernetesHost));
        }

        /**
         * @param namespace The namespace to provision the resource in.
         * The value should not contain leading or trailing forward slashes.
         * The `namespace` is always relative to the provider&#39;s configured namespace.
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
         * The `namespace` is always relative to the provider&#39;s configured namespace.
         * *Available only for Vault Enterprise*.
         * 
         * @return builder
         * 
         */
        public Builder namespace(String namespace) {
            return namespace(Output.of(namespace));
        }

        /**
         * @param pemKeys List of PEM-formatted public keys or certificates used to verify the signatures of Kubernetes service account JWTs. If a certificate is given, its public key will be extracted. Not every installation of Kubernetes exposes these keys.
         * 
         * @return builder
         * 
         */
        public Builder pemKeys(@Nullable Output<List<String>> pemKeys) {
            $.pemKeys = pemKeys;
            return this;
        }

        /**
         * @param pemKeys List of PEM-formatted public keys or certificates used to verify the signatures of Kubernetes service account JWTs. If a certificate is given, its public key will be extracted. Not every installation of Kubernetes exposes these keys.
         * 
         * @return builder
         * 
         */
        public Builder pemKeys(List<String> pemKeys) {
            return pemKeys(Output.of(pemKeys));
        }

        /**
         * @param pemKeys List of PEM-formatted public keys or certificates used to verify the signatures of Kubernetes service account JWTs. If a certificate is given, its public key will be extracted. Not every installation of Kubernetes exposes these keys.
         * 
         * @return builder
         * 
         */
        public Builder pemKeys(String... pemKeys) {
            return pemKeys(List.of(pemKeys));
        }

        /**
         * @param tokenReviewerJwt A service account JWT (or other token) used as a bearer token to access the TokenReview API to validate other JWTs during login. If not set the JWT used for login will be used to access the API.
         * 
         * @return builder
         * 
         */
        public Builder tokenReviewerJwt(@Nullable Output<String> tokenReviewerJwt) {
            $.tokenReviewerJwt = tokenReviewerJwt;
            return this;
        }

        /**
         * @param tokenReviewerJwt A service account JWT (or other token) used as a bearer token to access the TokenReview API to validate other JWTs during login. If not set the JWT used for login will be used to access the API.
         * 
         * @return builder
         * 
         */
        public Builder tokenReviewerJwt(String tokenReviewerJwt) {
            return tokenReviewerJwt(Output.of(tokenReviewerJwt));
        }

        public AuthBackendConfigArgs build() {
            if ($.kubernetesHost == null) {
                throw new MissingRequiredPropertyException("AuthBackendConfigArgs", "kubernetesHost");
            }
            return $;
        }
    }

}
