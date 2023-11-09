// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export function getAuthBackends(args?: GetAuthBackendsArgs, opts?: pulumi.InvokeOptions): Promise<GetAuthBackendsResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("vault:index/getAuthBackends:getAuthBackends", {
        "namespace": args.namespace,
        "type": args.type,
    }, opts);
}

/**
 * A collection of arguments for invoking getAuthBackends.
 */
export interface GetAuthBackendsArgs {
    /**
     * The namespace of the target resource.
     * The value should not contain leading or trailing forward slashes.
     * The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
     * *Available only for Vault Enterprise*.
     */
    namespace?: string;
    /**
     * The name of the auth method type. Allows filtering of backends returned by type.
     */
    type?: string;
}

/**
 * A collection of values returned by getAuthBackends.
 */
export interface GetAuthBackendsResult {
    /**
     * The accessor IDs for the auth methods.
     */
    readonly accessors: string[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly namespace?: string;
    /**
     * List of auth backend mount points.
     */
    readonly paths: string[];
    readonly type?: string;
}
export function getAuthBackendsOutput(args?: GetAuthBackendsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetAuthBackendsResult> {
    return pulumi.output(args).apply((a: any) => getAuthBackends(a, opts))
}

/**
 * A collection of arguments for invoking getAuthBackends.
 */
export interface GetAuthBackendsOutputArgs {
    /**
     * The namespace of the target resource.
     * The value should not contain leading or trailing forward slashes.
     * The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
     * *Available only for Vault Enterprise*.
     */
    namespace?: pulumi.Input<string>;
    /**
     * The name of the auth method type. Allows filtering of backends returned by type.
     */
    type?: pulumi.Input<string>;
}