// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vault.kubernetes.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class GetAuthBackendConfigResult {
    private @Nullable String backend;
    private Boolean disableIssValidation;
    private Boolean disableLocalCaJwt;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    /**
     * @return Optional JWT issuer. If no issuer is specified, `kubernetes.io/serviceaccount` will be used as the default issuer.
     * 
     */
    private String issuer;
    /**
     * @return PEM encoded CA cert for use by the TLS client used to talk with the Kubernetes API.
     * 
     */
    private String kubernetesCaCert;
    /**
     * @return Host must be a host string, a host:port pair, or a URL to the base of the Kubernetes API server.
     * 
     */
    private String kubernetesHost;
    private @Nullable String namespace;
    /**
     * @return Optional list of PEM-formatted public keys or certificates used to verify the signatures of Kubernetes service account JWTs. If a certificate is given, its public key will be extracted. Not every installation of Kubernetes exposes these keys.
     * 
     */
    private List<String> pemKeys;

    private GetAuthBackendConfigResult() {}
    public Optional<String> backend() {
        return Optional.ofNullable(this.backend);
    }
    public Boolean disableIssValidation() {
        return this.disableIssValidation;
    }
    public Boolean disableLocalCaJwt() {
        return this.disableLocalCaJwt;
    }
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @return Optional JWT issuer. If no issuer is specified, `kubernetes.io/serviceaccount` will be used as the default issuer.
     * 
     */
    public String issuer() {
        return this.issuer;
    }
    /**
     * @return PEM encoded CA cert for use by the TLS client used to talk with the Kubernetes API.
     * 
     */
    public String kubernetesCaCert() {
        return this.kubernetesCaCert;
    }
    /**
     * @return Host must be a host string, a host:port pair, or a URL to the base of the Kubernetes API server.
     * 
     */
    public String kubernetesHost() {
        return this.kubernetesHost;
    }
    public Optional<String> namespace() {
        return Optional.ofNullable(this.namespace);
    }
    /**
     * @return Optional list of PEM-formatted public keys or certificates used to verify the signatures of Kubernetes service account JWTs. If a certificate is given, its public key will be extracted. Not every installation of Kubernetes exposes these keys.
     * 
     */
    public List<String> pemKeys() {
        return this.pemKeys;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetAuthBackendConfigResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable String backend;
        private Boolean disableIssValidation;
        private Boolean disableLocalCaJwt;
        private String id;
        private String issuer;
        private String kubernetesCaCert;
        private String kubernetesHost;
        private @Nullable String namespace;
        private List<String> pemKeys;
        public Builder() {}
        public Builder(GetAuthBackendConfigResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.backend = defaults.backend;
    	      this.disableIssValidation = defaults.disableIssValidation;
    	      this.disableLocalCaJwt = defaults.disableLocalCaJwt;
    	      this.id = defaults.id;
    	      this.issuer = defaults.issuer;
    	      this.kubernetesCaCert = defaults.kubernetesCaCert;
    	      this.kubernetesHost = defaults.kubernetesHost;
    	      this.namespace = defaults.namespace;
    	      this.pemKeys = defaults.pemKeys;
        }

        @CustomType.Setter
        public Builder backend(@Nullable String backend) {

            this.backend = backend;
            return this;
        }
        @CustomType.Setter
        public Builder disableIssValidation(Boolean disableIssValidation) {
            if (disableIssValidation == null) {
              throw new MissingRequiredPropertyException("GetAuthBackendConfigResult", "disableIssValidation");
            }
            this.disableIssValidation = disableIssValidation;
            return this;
        }
        @CustomType.Setter
        public Builder disableLocalCaJwt(Boolean disableLocalCaJwt) {
            if (disableLocalCaJwt == null) {
              throw new MissingRequiredPropertyException("GetAuthBackendConfigResult", "disableLocalCaJwt");
            }
            this.disableLocalCaJwt = disableLocalCaJwt;
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            if (id == null) {
              throw new MissingRequiredPropertyException("GetAuthBackendConfigResult", "id");
            }
            this.id = id;
            return this;
        }
        @CustomType.Setter
        public Builder issuer(String issuer) {
            if (issuer == null) {
              throw new MissingRequiredPropertyException("GetAuthBackendConfigResult", "issuer");
            }
            this.issuer = issuer;
            return this;
        }
        @CustomType.Setter
        public Builder kubernetesCaCert(String kubernetesCaCert) {
            if (kubernetesCaCert == null) {
              throw new MissingRequiredPropertyException("GetAuthBackendConfigResult", "kubernetesCaCert");
            }
            this.kubernetesCaCert = kubernetesCaCert;
            return this;
        }
        @CustomType.Setter
        public Builder kubernetesHost(String kubernetesHost) {
            if (kubernetesHost == null) {
              throw new MissingRequiredPropertyException("GetAuthBackendConfigResult", "kubernetesHost");
            }
            this.kubernetesHost = kubernetesHost;
            return this;
        }
        @CustomType.Setter
        public Builder namespace(@Nullable String namespace) {

            this.namespace = namespace;
            return this;
        }
        @CustomType.Setter
        public Builder pemKeys(List<String> pemKeys) {
            if (pemKeys == null) {
              throw new MissingRequiredPropertyException("GetAuthBackendConfigResult", "pemKeys");
            }
            this.pemKeys = pemKeys;
            return this;
        }
        public Builder pemKeys(String... pemKeys) {
            return pemKeys(List.of(pemKeys));
        }
        public GetAuthBackendConfigResult build() {
            final var _resultValue = new GetAuthBackendConfigResult();
            _resultValue.backend = backend;
            _resultValue.disableIssValidation = disableIssValidation;
            _resultValue.disableLocalCaJwt = disableLocalCaJwt;
            _resultValue.id = id;
            _resultValue.issuer = issuer;
            _resultValue.kubernetesCaCert = kubernetesCaCert;
            _resultValue.kubernetesHost = kubernetesHost;
            _resultValue.namespace = namespace;
            _resultValue.pemKeys = pemKeys;
            return _resultValue;
        }
    }
}
