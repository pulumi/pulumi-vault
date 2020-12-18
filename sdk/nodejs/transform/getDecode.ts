// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

export function getDecode(args: GetDecodeArgs, opts?: pulumi.InvokeOptions): Promise<GetDecodeResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("vault:transform/getDecode:getDecode", {
        "batchInputs": args.batchInputs,
        "batchResults": args.batchResults,
        "decodedValue": args.decodedValue,
        "path": args.path,
        "roleName": args.roleName,
        "transformation": args.transformation,
        "tweak": args.tweak,
        "value": args.value,
    }, opts);
}

/**
 * A collection of arguments for invoking getDecode.
 */
export interface GetDecodeArgs {
    readonly batchInputs?: {[key: string]: any}[];
    readonly batchResults?: {[key: string]: any}[];
    readonly decodedValue?: string;
    readonly path: string;
    readonly roleName: string;
    readonly transformation?: string;
    readonly tweak?: string;
    readonly value?: string;
}

/**
 * A collection of values returned by getDecode.
 */
export interface GetDecodeResult {
    readonly batchInputs?: {[key: string]: any}[];
    readonly batchResults: {[key: string]: any}[];
    readonly decodedValue: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly path: string;
    readonly roleName: string;
    readonly transformation?: string;
    readonly tweak?: string;
    readonly value?: string;
}
