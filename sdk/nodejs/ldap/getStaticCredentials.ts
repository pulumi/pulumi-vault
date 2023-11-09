// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

export function getStaticCredentials(args: GetStaticCredentialsArgs, opts?: pulumi.InvokeOptions): Promise<GetStaticCredentialsResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("vault:ldap/getStaticCredentials:getStaticCredentials", {
        "mount": args.mount,
        "namespace": args.namespace,
        "roleName": args.roleName,
    }, opts);
}

/**
 * A collection of arguments for invoking getStaticCredentials.
 */
export interface GetStaticCredentialsArgs {
    mount: string;
    namespace?: string;
    roleName: string;
}

/**
 * A collection of values returned by getStaticCredentials.
 */
export interface GetStaticCredentialsResult {
    readonly dn: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly lastPassword: string;
    readonly lastVaultRotation: string;
    readonly mount: string;
    readonly namespace?: string;
    readonly password: string;
    readonly roleName: string;
    readonly rotationPeriod: number;
    readonly ttl: number;
    readonly username: string;
}
export function getStaticCredentialsOutput(args: GetStaticCredentialsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetStaticCredentialsResult> {
    return pulumi.output(args).apply((a: any) => getStaticCredentials(a, opts))
}

/**
 * A collection of arguments for invoking getStaticCredentials.
 */
export interface GetStaticCredentialsOutputArgs {
    mount: pulumi.Input<string>;
    namespace?: pulumi.Input<string>;
    roleName: pulumi.Input<string>;
}
