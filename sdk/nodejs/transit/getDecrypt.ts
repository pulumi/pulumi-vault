// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * This is a data source which can be used to decrypt ciphertext using a Vault Transit key.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vault from "@pulumi/vault";
 *
 * const test = vault.transit.getDecrypt({
 *     backend: "transit",
 *     key: "test",
 *     ciphertext: "vault:v1:S3GtnJ5GUNCWV+/pdL9+g1Feu/nzAv+RlmTmE91Tu0rBkeIU8MEb2nSspC/1IQ==",
 * });
 * ```
 */
export function getDecrypt(args: GetDecryptArgs, opts?: pulumi.InvokeOptions): Promise<GetDecryptResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("vault:transit/getDecrypt:getDecrypt", {
        "backend": args.backend,
        "ciphertext": args.ciphertext,
        "context": args.context,
        "key": args.key,
        "namespace": args.namespace,
    }, opts);
}

/**
 * A collection of arguments for invoking getDecrypt.
 */
export interface GetDecryptArgs {
    backend: string;
    ciphertext: string;
    context?: string;
    key: string;
    namespace?: string;
}

/**
 * A collection of values returned by getDecrypt.
 */
export interface GetDecryptResult {
    readonly backend: string;
    readonly ciphertext: string;
    readonly context?: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly key: string;
    readonly namespace?: string;
    /**
     * Decrypted plaintext returned from Vault
     */
    readonly plaintext: string;
}
/**
 * This is a data source which can be used to decrypt ciphertext using a Vault Transit key.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vault from "@pulumi/vault";
 *
 * const test = vault.transit.getDecrypt({
 *     backend: "transit",
 *     key: "test",
 *     ciphertext: "vault:v1:S3GtnJ5GUNCWV+/pdL9+g1Feu/nzAv+RlmTmE91Tu0rBkeIU8MEb2nSspC/1IQ==",
 * });
 * ```
 */
export function getDecryptOutput(args: GetDecryptOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetDecryptResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("vault:transit/getDecrypt:getDecrypt", {
        "backend": args.backend,
        "ciphertext": args.ciphertext,
        "context": args.context,
        "key": args.key,
        "namespace": args.namespace,
    }, opts);
}

/**
 * A collection of arguments for invoking getDecrypt.
 */
export interface GetDecryptOutputArgs {
    backend: pulumi.Input<string>;
    ciphertext: pulumi.Input<string>;
    context?: pulumi.Input<string>;
    key: pulumi.Input<string>;
    namespace?: pulumi.Input<string>;
}
