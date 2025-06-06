// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vault from "@pulumi/vault";
 *
 * const kvv1 = new vault.Mount("kvv1", {
 *     path: "kvv1",
 *     type: "kv",
 *     options: {
 *         version: "1",
 *     },
 *     description: "KV Version 1 secret engine mount",
 * });
 * const awsSecret = new vault.kv.Secret("aws_secret", {
 *     path: pulumi.interpolate`${kvv1.path}/aws-secret`,
 *     dataJson: JSON.stringify({
 *         zip: "zap",
 *     }),
 * });
 * const azureSecret = new vault.kv.Secret("azure_secret", {
 *     path: pulumi.interpolate`${kvv1.path}/azure-secret`,
 *     dataJson: JSON.stringify({
 *         foo: "bar",
 *     }),
 * });
 * const secrets = vault.kv.getSecretsListOutput({
 *     path: kvv1.path,
 * });
 * ```
 *
 * ## Required Vault Capabilities
 *
 * Use of this resource requires the `read` capability on the given path.
 */
export function getSecretsList(args: GetSecretsListArgs, opts?: pulumi.InvokeOptions): Promise<GetSecretsListResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("vault:kv/getSecretsList:getSecretsList", {
        "namespace": args.namespace,
        "path": args.path,
    }, opts);
}

/**
 * A collection of arguments for invoking getSecretsList.
 */
export interface GetSecretsListArgs {
    /**
     * The namespace of the target resource.
     * The value should not contain leading or trailing forward slashes.
     * The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
     * *Available only for Vault Enterprise*.
     */
    namespace?: string;
    /**
     * Full KV-V1 path where secrets will be listed.
     */
    path: string;
}

/**
 * A collection of values returned by getSecretsList.
 */
export interface GetSecretsListResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * List of all secret names listed under the given path.
     */
    readonly names: string[];
    readonly namespace?: string;
    readonly path: string;
}
/**
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vault from "@pulumi/vault";
 *
 * const kvv1 = new vault.Mount("kvv1", {
 *     path: "kvv1",
 *     type: "kv",
 *     options: {
 *         version: "1",
 *     },
 *     description: "KV Version 1 secret engine mount",
 * });
 * const awsSecret = new vault.kv.Secret("aws_secret", {
 *     path: pulumi.interpolate`${kvv1.path}/aws-secret`,
 *     dataJson: JSON.stringify({
 *         zip: "zap",
 *     }),
 * });
 * const azureSecret = new vault.kv.Secret("azure_secret", {
 *     path: pulumi.interpolate`${kvv1.path}/azure-secret`,
 *     dataJson: JSON.stringify({
 *         foo: "bar",
 *     }),
 * });
 * const secrets = vault.kv.getSecretsListOutput({
 *     path: kvv1.path,
 * });
 * ```
 *
 * ## Required Vault Capabilities
 *
 * Use of this resource requires the `read` capability on the given path.
 */
export function getSecretsListOutput(args: GetSecretsListOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetSecretsListResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("vault:kv/getSecretsList:getSecretsList", {
        "namespace": args.namespace,
        "path": args.path,
    }, opts);
}

/**
 * A collection of arguments for invoking getSecretsList.
 */
export interface GetSecretsListOutputArgs {
    /**
     * The namespace of the target resource.
     * The value should not contain leading or trailing forward slashes.
     * The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
     * *Available only for Vault Enterprise*.
     */
    namespace?: pulumi.Input<string>;
    /**
     * Full KV-V1 path where secrets will be listed.
     */
    path: pulumi.Input<string>;
}
