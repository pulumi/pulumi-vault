// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vault.kubernetes.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetServiceAccountTokenArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetServiceAccountTokenArgs Empty = new GetServiceAccountTokenArgs();

    /**
     * The Kubernetes secret backend to generate service account
     * tokens from.
     * 
     */
    @Import(name="backend", required=true)
    private Output<String> backend;

    /**
     * @return The Kubernetes secret backend to generate service account
     * tokens from.
     * 
     */
    public Output<String> backend() {
        return this.backend;
    }

    /**
     * If true, generate a ClusterRoleBinding to grant
     * permissions across the whole cluster instead of within a namespace.
     * 
     */
    @Import(name="clusterRoleBinding")
    private @Nullable Output<Boolean> clusterRoleBinding;

    /**
     * @return If true, generate a ClusterRoleBinding to grant
     * permissions across the whole cluster instead of within a namespace.
     * 
     */
    public Optional<Output<Boolean>> clusterRoleBinding() {
        return Optional.ofNullable(this.clusterRoleBinding);
    }

    /**
     * The name of the Kubernetes namespace in which to
     * generate the credentials.
     * 
     */
    @Import(name="kubernetesNamespace", required=true)
    private Output<String> kubernetesNamespace;

    /**
     * @return The name of the Kubernetes namespace in which to
     * generate the credentials.
     * 
     */
    public Output<String> kubernetesNamespace() {
        return this.kubernetesNamespace;
    }

    /**
     * The namespace of the target resource.
     * The value should not contain leading or trailing forward slashes.
     * The `namespace` is always relative to the provider&#39;s configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
     * *Available only for Vault Enterprise*.
     * 
     */
    @Import(name="namespace")
    private @Nullable Output<String> namespace;

    /**
     * @return The namespace of the target resource.
     * The value should not contain leading or trailing forward slashes.
     * The `namespace` is always relative to the provider&#39;s configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
     * *Available only for Vault Enterprise*.
     * 
     */
    public Optional<Output<String>> namespace() {
        return Optional.ofNullable(this.namespace);
    }

    /**
     * The name of the Kubernetes secret backend role to generate service
     * account tokens from.
     * 
     */
    @Import(name="role", required=true)
    private Output<String> role;

    /**
     * @return The name of the Kubernetes secret backend role to generate service
     * account tokens from.
     * 
     */
    public Output<String> role() {
        return this.role;
    }

    /**
     * The TTL of the generated Kubernetes service account token, specified in
     * seconds or as a Go duration format string.
     * 
     */
    @Import(name="ttl")
    private @Nullable Output<String> ttl;

    /**
     * @return The TTL of the generated Kubernetes service account token, specified in
     * seconds or as a Go duration format string.
     * 
     */
    public Optional<Output<String>> ttl() {
        return Optional.ofNullable(this.ttl);
    }

    private GetServiceAccountTokenArgs() {}

    private GetServiceAccountTokenArgs(GetServiceAccountTokenArgs $) {
        this.backend = $.backend;
        this.clusterRoleBinding = $.clusterRoleBinding;
        this.kubernetesNamespace = $.kubernetesNamespace;
        this.namespace = $.namespace;
        this.role = $.role;
        this.ttl = $.ttl;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetServiceAccountTokenArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetServiceAccountTokenArgs $;

        public Builder() {
            $ = new GetServiceAccountTokenArgs();
        }

        public Builder(GetServiceAccountTokenArgs defaults) {
            $ = new GetServiceAccountTokenArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param backend The Kubernetes secret backend to generate service account
         * tokens from.
         * 
         * @return builder
         * 
         */
        public Builder backend(Output<String> backend) {
            $.backend = backend;
            return this;
        }

        /**
         * @param backend The Kubernetes secret backend to generate service account
         * tokens from.
         * 
         * @return builder
         * 
         */
        public Builder backend(String backend) {
            return backend(Output.of(backend));
        }

        /**
         * @param clusterRoleBinding If true, generate a ClusterRoleBinding to grant
         * permissions across the whole cluster instead of within a namespace.
         * 
         * @return builder
         * 
         */
        public Builder clusterRoleBinding(@Nullable Output<Boolean> clusterRoleBinding) {
            $.clusterRoleBinding = clusterRoleBinding;
            return this;
        }

        /**
         * @param clusterRoleBinding If true, generate a ClusterRoleBinding to grant
         * permissions across the whole cluster instead of within a namespace.
         * 
         * @return builder
         * 
         */
        public Builder clusterRoleBinding(Boolean clusterRoleBinding) {
            return clusterRoleBinding(Output.of(clusterRoleBinding));
        }

        /**
         * @param kubernetesNamespace The name of the Kubernetes namespace in which to
         * generate the credentials.
         * 
         * @return builder
         * 
         */
        public Builder kubernetesNamespace(Output<String> kubernetesNamespace) {
            $.kubernetesNamespace = kubernetesNamespace;
            return this;
        }

        /**
         * @param kubernetesNamespace The name of the Kubernetes namespace in which to
         * generate the credentials.
         * 
         * @return builder
         * 
         */
        public Builder kubernetesNamespace(String kubernetesNamespace) {
            return kubernetesNamespace(Output.of(kubernetesNamespace));
        }

        /**
         * @param namespace The namespace of the target resource.
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
         * @param namespace The namespace of the target resource.
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
         * @param role The name of the Kubernetes secret backend role to generate service
         * account tokens from.
         * 
         * @return builder
         * 
         */
        public Builder role(Output<String> role) {
            $.role = role;
            return this;
        }

        /**
         * @param role The name of the Kubernetes secret backend role to generate service
         * account tokens from.
         * 
         * @return builder
         * 
         */
        public Builder role(String role) {
            return role(Output.of(role));
        }

        /**
         * @param ttl The TTL of the generated Kubernetes service account token, specified in
         * seconds or as a Go duration format string.
         * 
         * @return builder
         * 
         */
        public Builder ttl(@Nullable Output<String> ttl) {
            $.ttl = ttl;
            return this;
        }

        /**
         * @param ttl The TTL of the generated Kubernetes service account token, specified in
         * seconds or as a Go duration format string.
         * 
         * @return builder
         * 
         */
        public Builder ttl(String ttl) {
            return ttl(Output.of(ttl));
        }

        public GetServiceAccountTokenArgs build() {
            if ($.backend == null) {
                throw new MissingRequiredPropertyException("GetServiceAccountTokenArgs", "backend");
            }
            if ($.kubernetesNamespace == null) {
                throw new MissingRequiredPropertyException("GetServiceAccountTokenArgs", "kubernetesNamespace");
            }
            if ($.role == null) {
                throw new MissingRequiredPropertyException("GetServiceAccountTokenArgs", "role");
            }
            return $;
        }
    }

}
