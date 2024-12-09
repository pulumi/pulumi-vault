// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Reads the Role of an Kubernetes from a Vault server. See the [Vault
 * documentation](https://www.vaultproject.io/api-docs/auth/kubernetes#read-config) for more
 * information.
 */
export function getAuthBackendConfig(args?: GetAuthBackendConfigArgs, opts?: pulumi.InvokeOptions): Promise<GetAuthBackendConfigResult> {
    args = args || {};
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("vault:kubernetes/getAuthBackendConfig:getAuthBackendConfig", {
        "backend": args.backend,
        "disableIssValidation": args.disableIssValidation,
        "disableLocalCaJwt": args.disableLocalCaJwt,
        "issuer": args.issuer,
        "kubernetesCaCert": args.kubernetesCaCert,
        "kubernetesHost": args.kubernetesHost,
        "namespace": args.namespace,
        "pemKeys": args.pemKeys,
        "useAnnotationsAsAliasMetadata": args.useAnnotationsAsAliasMetadata,
    }, opts);
}

/**
 * A collection of arguments for invoking getAuthBackendConfig.
 */
export interface GetAuthBackendConfigArgs {
    /**
     * The unique name for the Kubernetes backend the config to
     * retrieve Role attributes for resides in. Defaults to "kubernetes".
     */
    backend?: string;
    /**
     * (Optional) Disable JWT issuer validation. Allows to skip ISS validation. Requires Vault `v1.5.4+` or Vault auth kubernetes plugin `v0.7.1+`
     */
    disableIssValidation?: boolean;
    /**
     * (Optional) Disable defaulting to the local CA cert and service account JWT when running in a Kubernetes pod. Requires Vault `v1.5.4+` or Vault auth kubernetes plugin `v0.7.1+`
     */
    disableLocalCaJwt?: boolean;
    /**
     * Optional JWT issuer. If no issuer is specified, `kubernetes.io/serviceaccount` will be used as the default issuer.
     */
    issuer?: string;
    /**
     * PEM encoded CA cert for use by the TLS client used to talk with the Kubernetes API.
     */
    kubernetesCaCert?: string;
    /**
     * Host must be a host string, a host:port pair, or a URL to the base of the Kubernetes API server.
     */
    kubernetesHost?: string;
    /**
     * The namespace of the target resource.
     * The value should not contain leading or trailing forward slashes.
     * The `namespace` is always relative to the provider's configured namespace.
     * *Available only for Vault Enterprise*.
     */
    namespace?: string;
    /**
     * Optional list of PEM-formatted public keys or certificates used to verify the signatures of Kubernetes service account JWTs. If a certificate is given, its public key will be extracted. Not every installation of Kubernetes exposes these keys.
     */
    pemKeys?: string[];
    /**
     * (Optional) Use annotations from the client token's associated service account as alias metadata for the Vault entity. Requires Vault `v1.16+` or Vault auth kubernetes plugin `v0.18.0+`
     */
    useAnnotationsAsAliasMetadata?: boolean;
}

/**
 * A collection of values returned by getAuthBackendConfig.
 */
export interface GetAuthBackendConfigResult {
    readonly backend?: string;
    /**
     * (Optional) Disable JWT issuer validation. Allows to skip ISS validation. Requires Vault `v1.5.4+` or Vault auth kubernetes plugin `v0.7.1+`
     */
    readonly disableIssValidation: boolean;
    /**
     * (Optional) Disable defaulting to the local CA cert and service account JWT when running in a Kubernetes pod. Requires Vault `v1.5.4+` or Vault auth kubernetes plugin `v0.7.1+`
     */
    readonly disableLocalCaJwt: boolean;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * Optional JWT issuer. If no issuer is specified, `kubernetes.io/serviceaccount` will be used as the default issuer.
     */
    readonly issuer: string;
    /**
     * PEM encoded CA cert for use by the TLS client used to talk with the Kubernetes API.
     */
    readonly kubernetesCaCert: string;
    /**
     * Host must be a host string, a host:port pair, or a URL to the base of the Kubernetes API server.
     */
    readonly kubernetesHost: string;
    readonly namespace?: string;
    /**
     * Optional list of PEM-formatted public keys or certificates used to verify the signatures of Kubernetes service account JWTs. If a certificate is given, its public key will be extracted. Not every installation of Kubernetes exposes these keys.
     */
    readonly pemKeys: string[];
    /**
     * (Optional) Use annotations from the client token's associated service account as alias metadata for the Vault entity. Requires Vault `v1.16+` or Vault auth kubernetes plugin `v0.18.0+`
     */
    readonly useAnnotationsAsAliasMetadata: boolean;
}
/**
 * Reads the Role of an Kubernetes from a Vault server. See the [Vault
 * documentation](https://www.vaultproject.io/api-docs/auth/kubernetes#read-config) for more
 * information.
 */
export function getAuthBackendConfigOutput(args?: GetAuthBackendConfigOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetAuthBackendConfigResult> {
    args = args || {};
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("vault:kubernetes/getAuthBackendConfig:getAuthBackendConfig", {
        "backend": args.backend,
        "disableIssValidation": args.disableIssValidation,
        "disableLocalCaJwt": args.disableLocalCaJwt,
        "issuer": args.issuer,
        "kubernetesCaCert": args.kubernetesCaCert,
        "kubernetesHost": args.kubernetesHost,
        "namespace": args.namespace,
        "pemKeys": args.pemKeys,
        "useAnnotationsAsAliasMetadata": args.useAnnotationsAsAliasMetadata,
    }, opts);
}

/**
 * A collection of arguments for invoking getAuthBackendConfig.
 */
export interface GetAuthBackendConfigOutputArgs {
    /**
     * The unique name for the Kubernetes backend the config to
     * retrieve Role attributes for resides in. Defaults to "kubernetes".
     */
    backend?: pulumi.Input<string>;
    /**
     * (Optional) Disable JWT issuer validation. Allows to skip ISS validation. Requires Vault `v1.5.4+` or Vault auth kubernetes plugin `v0.7.1+`
     */
    disableIssValidation?: pulumi.Input<boolean>;
    /**
     * (Optional) Disable defaulting to the local CA cert and service account JWT when running in a Kubernetes pod. Requires Vault `v1.5.4+` or Vault auth kubernetes plugin `v0.7.1+`
     */
    disableLocalCaJwt?: pulumi.Input<boolean>;
    /**
     * Optional JWT issuer. If no issuer is specified, `kubernetes.io/serviceaccount` will be used as the default issuer.
     */
    issuer?: pulumi.Input<string>;
    /**
     * PEM encoded CA cert for use by the TLS client used to talk with the Kubernetes API.
     */
    kubernetesCaCert?: pulumi.Input<string>;
    /**
     * Host must be a host string, a host:port pair, or a URL to the base of the Kubernetes API server.
     */
    kubernetesHost?: pulumi.Input<string>;
    /**
     * The namespace of the target resource.
     * The value should not contain leading or trailing forward slashes.
     * The `namespace` is always relative to the provider's configured namespace.
     * *Available only for Vault Enterprise*.
     */
    namespace?: pulumi.Input<string>;
    /**
     * Optional list of PEM-formatted public keys or certificates used to verify the signatures of Kubernetes service account JWTs. If a certificate is given, its public key will be extracted. Not every installation of Kubernetes exposes these keys.
     */
    pemKeys?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * (Optional) Use annotations from the client token's associated service account as alias metadata for the Vault entity. Requires Vault `v1.16+` or Vault auth kubernetes plugin `v0.18.0+`
     */
    useAnnotationsAsAliasMetadata?: pulumi.Input<boolean>;
}
