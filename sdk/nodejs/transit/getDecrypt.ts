// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * This is a data source which can be used to decrypt ciphertext using a Vault Transit key.
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vault from "@pulumi/vault";
 *
 * const test = vault.transit.getDecrypt({
 *     backend: "transit",
 *     ciphertext: "vault:v1:S3GtnJ5GUNCWV+/pdL9+g1Feu/nzAv+RlmTmE91Tu0rBkeIU8MEb2nSspC/1IQ==",
 *     key: "test",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
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
    /**
     * The path the transit secret backend is mounted at, with no leading or trailing `/`.
     */
    backend: string;
    /**
     * Ciphertext to be decoded.
     */
    ciphertext: string;
    /**
     * Context for key derivation. This is required if key derivation is enabled for this key.
     */
    context?: string;
    /**
     * Specifies the name of the transit key to decrypt against.
     */
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
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vault from "@pulumi/vault";
 *
 * const test = vault.transit.getDecrypt({
 *     backend: "transit",
 *     ciphertext: "vault:v1:S3GtnJ5GUNCWV+/pdL9+g1Feu/nzAv+RlmTmE91Tu0rBkeIU8MEb2nSspC/1IQ==",
 *     key: "test",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getDecryptOutput(args: GetDecryptOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetDecryptResult> {
    return pulumi.output(args).apply((a: any) => getDecrypt(a, opts))
}

/**
 * A collection of arguments for invoking getDecrypt.
 */
export interface GetDecryptOutputArgs {
    /**
     * The path the transit secret backend is mounted at, with no leading or trailing `/`.
     */
    backend: pulumi.Input<string>;
    /**
     * Ciphertext to be decoded.
     */
    ciphertext: pulumi.Input<string>;
    /**
     * Context for key derivation. This is required if key derivation is enabled for this key.
     */
    context?: pulumi.Input<string>;
    /**
     * Specifies the name of the transit key to decrypt against.
     */
    key: pulumi.Input<string>;
    namespace?: pulumi.Input<string>;
}
