// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * This data source supports the "/transform/decode/{role_name}" Vault endpoint.
 *
 * It decodes the provided value using a named role.
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vault from "@pulumi/vault";
 *
 * const transform = new vault.Mount("transform", {
 *     path: "transform",
 *     type: "transform",
 * });
 * const ccn_fpe = new vault.transform.Transformation("ccn-fpe", {
 *     path: transform.path,
 *     type: "fpe",
 *     template: "builtin/creditcardnumber",
 *     tweakSource: "internal",
 *     allowedRoles: ["payments"],
 * });
 * const payments = new vault.transform.Role("payments", {
 *     path: ccn_fpe.path,
 *     transformations: ["ccn-fpe"],
 * });
 * const test = vault.transform.getDecodeOutput({
 *     path: payments.path,
 *     roleName: "payments",
 *     value: "9300-3376-4943-8903",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getDecode(args: GetDecodeArgs, opts?: pulumi.InvokeOptions): Promise<GetDecodeResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("vault:transform/getDecode:getDecode", {
        "batchInputs": args.batchInputs,
        "batchResults": args.batchResults,
        "decodedValue": args.decodedValue,
        "namespace": args.namespace,
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
    /**
     * Specifies a list of items to be decoded in a single batch. If this parameter is set, the top-level parameters 'value', 'transformation' and 'tweak' will be ignored. Each batch item within the list can specify these parameters instead.
     */
    batchInputs?: {[key: string]: any}[];
    /**
     * The result of decoding a batch.
     */
    batchResults?: {[key: string]: any}[];
    /**
     * The result of decoding a value.
     */
    decodedValue?: string;
    /**
     * The namespace of the target resource.
     * The value should not contain leading or trailing forward slashes.
     * The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
     * *Available only for Vault Enterprise*.
     */
    namespace?: string;
    /**
     * Path to where the back-end is mounted within Vault.
     */
    path: string;
    /**
     * The name of the role.
     */
    roleName: string;
    /**
     * The transformation to perform. If no value is provided and the role contains a single transformation, this value will be inferred from the role.
     */
    transformation?: string;
    /**
     * The tweak value to use. Only applicable for FPE transformations
     */
    tweak?: string;
    /**
     * The value in which to decode.
     */
    value?: string;
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
    readonly namespace?: string;
    readonly path: string;
    readonly roleName: string;
    readonly transformation?: string;
    readonly tweak?: string;
    readonly value?: string;
}
/**
 * This data source supports the "/transform/decode/{role_name}" Vault endpoint.
 *
 * It decodes the provided value using a named role.
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vault from "@pulumi/vault";
 *
 * const transform = new vault.Mount("transform", {
 *     path: "transform",
 *     type: "transform",
 * });
 * const ccn_fpe = new vault.transform.Transformation("ccn-fpe", {
 *     path: transform.path,
 *     type: "fpe",
 *     template: "builtin/creditcardnumber",
 *     tweakSource: "internal",
 *     allowedRoles: ["payments"],
 * });
 * const payments = new vault.transform.Role("payments", {
 *     path: ccn_fpe.path,
 *     transformations: ["ccn-fpe"],
 * });
 * const test = vault.transform.getDecodeOutput({
 *     path: payments.path,
 *     roleName: "payments",
 *     value: "9300-3376-4943-8903",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getDecodeOutput(args: GetDecodeOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetDecodeResult> {
    return pulumi.output(args).apply((a: any) => getDecode(a, opts))
}

/**
 * A collection of arguments for invoking getDecode.
 */
export interface GetDecodeOutputArgs {
    /**
     * Specifies a list of items to be decoded in a single batch. If this parameter is set, the top-level parameters 'value', 'transformation' and 'tweak' will be ignored. Each batch item within the list can specify these parameters instead.
     */
    batchInputs?: pulumi.Input<pulumi.Input<{[key: string]: any}>[]>;
    /**
     * The result of decoding a batch.
     */
    batchResults?: pulumi.Input<pulumi.Input<{[key: string]: any}>[]>;
    /**
     * The result of decoding a value.
     */
    decodedValue?: pulumi.Input<string>;
    /**
     * The namespace of the target resource.
     * The value should not contain leading or trailing forward slashes.
     * The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
     * *Available only for Vault Enterprise*.
     */
    namespace?: pulumi.Input<string>;
    /**
     * Path to where the back-end is mounted within Vault.
     */
    path: pulumi.Input<string>;
    /**
     * The name of the role.
     */
    roleName: pulumi.Input<string>;
    /**
     * The transformation to perform. If no value is provided and the role contains a single transformation, this value will be inferred from the role.
     */
    transformation?: pulumi.Input<string>;
    /**
     * The tweak value to use. Only applicable for FPE transformations
     */
    tweak?: pulumi.Input<string>;
    /**
     * The value in which to decode.
     */
    value?: pulumi.Input<string>;
}
