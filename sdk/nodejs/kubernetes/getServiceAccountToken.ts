// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as std from "@pulumi/std";
 * import * as vault from "@pulumi/vault";
 *
 * const config = new vault.kubernetes.SecretBackend("config", {
 *     path: "kubernetes",
 *     description: "kubernetes secrets engine description",
 *     kubernetesHost: "https://127.0.0.1:61233",
 *     kubernetesCaCert: std.file({
 *         input: "/path/to/cert",
 *     }).then(invoke => invoke.result),
 *     serviceAccountJwt: std.file({
 *         input: "/path/to/token",
 *     }).then(invoke => invoke.result),
 *     disableLocalCaJwt: false,
 * });
 * const role = new vault.kubernetes.SecretBackendRole("role", {
 *     backend: config.path,
 *     name: "service-account-name-role",
 *     allowedKubernetesNamespaces: ["*"],
 *     tokenMaxTtl: 43200,
 *     tokenDefaultTtl: 21600,
 *     serviceAccountName: "test-service-account-with-generated-token",
 *     extraLabels: {
 *         id: "abc123",
 *         name: "some_name",
 *     },
 *     extraAnnotations: {
 *         env: "development",
 *         location: "earth",
 *     },
 * });
 * const token = vault.kubernetes.getServiceAccountTokenOutput({
 *     backend: config.path,
 *     role: role.name,
 *     kubernetesNamespace: "test",
 *     clusterRoleBinding: false,
 *     ttl: "1h",
 * });
 * ```
 */
export function getServiceAccountToken(args: GetServiceAccountTokenArgs, opts?: pulumi.InvokeOptions): Promise<GetServiceAccountTokenResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("vault:kubernetes/getServiceAccountToken:getServiceAccountToken", {
        "backend": args.backend,
        "clusterRoleBinding": args.clusterRoleBinding,
        "kubernetesNamespace": args.kubernetesNamespace,
        "namespace": args.namespace,
        "role": args.role,
        "ttl": args.ttl,
    }, opts);
}

/**
 * A collection of arguments for invoking getServiceAccountToken.
 */
export interface GetServiceAccountTokenArgs {
    /**
     * The Kubernetes secret backend to generate service account 
     * tokens from.
     */
    backend: string;
    /**
     * If true, generate a ClusterRoleBinding to grant 
     * permissions across the whole cluster instead of within a namespace.
     */
    clusterRoleBinding?: boolean;
    /**
     * The name of the Kubernetes namespace in which to 
     * generate the credentials.
     */
    kubernetesNamespace: string;
    /**
     * The namespace of the target resource.
     * The value should not contain leading or trailing forward slashes.
     * The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
     * *Available only for Vault Enterprise*.
     */
    namespace?: string;
    /**
     * The name of the Kubernetes secret backend role to generate service 
     * account tokens from.
     */
    role: string;
    /**
     * The TTL of the generated Kubernetes service account token, specified in 
     * seconds or as a Go duration format string.
     */
    ttl?: string;
}

/**
 * A collection of values returned by getServiceAccountToken.
 */
export interface GetServiceAccountTokenResult {
    readonly backend: string;
    readonly clusterRoleBinding?: boolean;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly kubernetesNamespace: string;
    /**
     * The duration of the lease in seconds.
     */
    readonly leaseDuration: number;
    /**
     * The lease identifier assigned by Vault.
     */
    readonly leaseId: string;
    /**
     * True if the duration of this lease can be extended through renewal.
     */
    readonly leaseRenewable: boolean;
    readonly namespace?: string;
    readonly role: string;
    /**
     * The name of the service account associated with the token.
     */
    readonly serviceAccountName: string;
    /**
     * The Kubernetes namespace that the service account resides in.
     */
    readonly serviceAccountNamespace: string;
    /**
     * The Kubernetes service account token.
     */
    readonly serviceAccountToken: string;
    readonly ttl?: string;
}
/**
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as std from "@pulumi/std";
 * import * as vault from "@pulumi/vault";
 *
 * const config = new vault.kubernetes.SecretBackend("config", {
 *     path: "kubernetes",
 *     description: "kubernetes secrets engine description",
 *     kubernetesHost: "https://127.0.0.1:61233",
 *     kubernetesCaCert: std.file({
 *         input: "/path/to/cert",
 *     }).then(invoke => invoke.result),
 *     serviceAccountJwt: std.file({
 *         input: "/path/to/token",
 *     }).then(invoke => invoke.result),
 *     disableLocalCaJwt: false,
 * });
 * const role = new vault.kubernetes.SecretBackendRole("role", {
 *     backend: config.path,
 *     name: "service-account-name-role",
 *     allowedKubernetesNamespaces: ["*"],
 *     tokenMaxTtl: 43200,
 *     tokenDefaultTtl: 21600,
 *     serviceAccountName: "test-service-account-with-generated-token",
 *     extraLabels: {
 *         id: "abc123",
 *         name: "some_name",
 *     },
 *     extraAnnotations: {
 *         env: "development",
 *         location: "earth",
 *     },
 * });
 * const token = vault.kubernetes.getServiceAccountTokenOutput({
 *     backend: config.path,
 *     role: role.name,
 *     kubernetesNamespace: "test",
 *     clusterRoleBinding: false,
 *     ttl: "1h",
 * });
 * ```
 */
export function getServiceAccountTokenOutput(args: GetServiceAccountTokenOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetServiceAccountTokenResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("vault:kubernetes/getServiceAccountToken:getServiceAccountToken", {
        "backend": args.backend,
        "clusterRoleBinding": args.clusterRoleBinding,
        "kubernetesNamespace": args.kubernetesNamespace,
        "namespace": args.namespace,
        "role": args.role,
        "ttl": args.ttl,
    }, opts);
}

/**
 * A collection of arguments for invoking getServiceAccountToken.
 */
export interface GetServiceAccountTokenOutputArgs {
    /**
     * The Kubernetes secret backend to generate service account 
     * tokens from.
     */
    backend: pulumi.Input<string>;
    /**
     * If true, generate a ClusterRoleBinding to grant 
     * permissions across the whole cluster instead of within a namespace.
     */
    clusterRoleBinding?: pulumi.Input<boolean>;
    /**
     * The name of the Kubernetes namespace in which to 
     * generate the credentials.
     */
    kubernetesNamespace: pulumi.Input<string>;
    /**
     * The namespace of the target resource.
     * The value should not contain leading or trailing forward slashes.
     * The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
     * *Available only for Vault Enterprise*.
     */
    namespace?: pulumi.Input<string>;
    /**
     * The name of the Kubernetes secret backend role to generate service 
     * account tokens from.
     */
    role: pulumi.Input<string>;
    /**
     * The TTL of the generated Kubernetes service account token, specified in 
     * seconds or as a Go duration format string.
     */
    ttl?: pulumi.Input<string>;
}
