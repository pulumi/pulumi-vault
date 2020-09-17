// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

export function getEncode(args: GetEncodeArgs, opts?: pulumi.InvokeOptions): Promise<GetEncodeResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("vault:transform/getEncode:getEncode", {
        "batchInputs": args.batchInputs,
        "batchResults": args.batchResults,
        "encodedValue": args.encodedValue,
        "path": args.path,
        "roleName": args.roleName,
        "transformation": args.transformation,
        "tweak": args.tweak,
        "value": args.value,
    }, opts);
}

/**
 * A collection of arguments for invoking getEncode.
 */
export interface GetEncodeArgs {
    readonly batchInputs?: {[key: string]: any}[];
    readonly batchResults?: {[key: string]: any}[];
    readonly encodedValue?: string;
    readonly path: string;
    readonly roleName: string;
    readonly transformation?: string;
    readonly tweak?: string;
    readonly value?: string;
}

/**
 * A collection of values returned by getEncode.
 */
export interface GetEncodeResult {
    readonly batchInputs?: {[key: string]: any}[];
    readonly batchResults: {[key: string]: any}[];
    readonly encodedValue: string;
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