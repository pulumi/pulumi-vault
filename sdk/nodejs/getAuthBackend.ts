// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vault from "@pulumi/vault";
 *
 * const example = vault.getAuthBackend({
 *     path: "userpass",
 * });
 * ```
 */
export function getAuthBackend(args: GetAuthBackendArgs, opts?: pulumi.InvokeOptions): Promise<GetAuthBackendResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("vault:index/getAuthBackend:getAuthBackend", {
        "namespace": args.namespace,
        "path": args.path,
    }, opts);
}

/**
 * A collection of arguments for invoking getAuthBackend.
 */
export interface GetAuthBackendArgs {
    /**
     * The namespace of the target resource.
     * The value should not contain leading or trailing forward slashes.
     * The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
     * *Available only for Vault Enterprise*.
     */
    namespace?: string;
    /**
     * The auth backend mount point.
     */
    path: string;
}

/**
 * A collection of values returned by getAuthBackend.
 */
export interface GetAuthBackendResult {
    /**
     * The accessor for this auth method.
     */
    readonly accessor: string;
    /**
     * The default lease duration in seconds.
     */
    readonly defaultLeaseTtlSeconds: number;
    /**
     * A description of the auth method.
     */
    readonly description: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * Specifies whether to show this mount in the UI-specific listing endpoint.
     */
    readonly listingVisibility: string;
    /**
     * Specifies if the auth method is local only.
     */
    readonly local: boolean;
    /**
     * The maximum lease duration in seconds.
     */
    readonly maxLeaseTtlSeconds: number;
    readonly namespace?: string;
    readonly path: string;
    /**
     * The name of the auth method type.
     */
    readonly type: string;
}
/**
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vault from "@pulumi/vault";
 *
 * const example = vault.getAuthBackend({
 *     path: "userpass",
 * });
 * ```
 */
export function getAuthBackendOutput(args: GetAuthBackendOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetAuthBackendResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("vault:index/getAuthBackend:getAuthBackend", {
        "namespace": args.namespace,
        "path": args.path,
    }, opts);
}

/**
 * A collection of arguments for invoking getAuthBackend.
 */
export interface GetAuthBackendOutputArgs {
    /**
     * The namespace of the target resource.
     * The value should not contain leading or trailing forward slashes.
     * The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
     * *Available only for Vault Enterprise*.
     */
    namespace?: pulumi.Input<string>;
    /**
     * The auth backend mount point.
     */
    path: pulumi.Input<string>;
}
