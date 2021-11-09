// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * This is a data source which can be used to encrypt plaintext using a Vault Transit key.
 */
export function getEncrypt(args: GetEncryptArgs, opts?: pulumi.InvokeOptions): Promise<GetEncryptResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("vault:transit/getEncrypt:getEncrypt", {
        "backend": args.backend,
        "context": args.context,
        "key": args.key,
        "keyVersion": args.keyVersion,
        "plaintext": args.plaintext,
    }, opts);
}

/**
 * A collection of arguments for invoking getEncrypt.
 */
export interface GetEncryptArgs {
    /**
     * The path the transit secret backend is mounted at, with no leading or trailing `/`.
     */
    backend: string;
    /**
     * Context for key derivation. This is required if key derivation is enabled for this key.
     */
    context?: string;
    /**
     * Specifies the name of the transit key to encrypt against.
     */
    key: string;
    /**
     * The version of the key to use for encryption. If not set, uses the latest version. Must be greater than or equal to the key's `minEncryptionVersion`, if set.
     */
    keyVersion?: number;
    /**
     * Plaintext to be encoded.
     */
    plaintext: string;
}

/**
 * A collection of values returned by getEncrypt.
 */
export interface GetEncryptResult {
    readonly backend: string;
    /**
     * Encrypted ciphertext returned from Vault
     */
    readonly ciphertext: string;
    readonly context?: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly key: string;
    readonly keyVersion?: number;
    readonly plaintext: string;
}

export function getEncryptOutput(args: GetEncryptOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetEncryptResult> {
    return pulumi.output(args).apply(a => getEncrypt(a, opts))
}

/**
 * A collection of arguments for invoking getEncrypt.
 */
export interface GetEncryptOutputArgs {
    /**
     * The path the transit secret backend is mounted at, with no leading or trailing `/`.
     */
    backend: pulumi.Input<string>;
    /**
     * Context for key derivation. This is required if key derivation is enabled for this key.
     */
    context?: pulumi.Input<string>;
    /**
     * Specifies the name of the transit key to encrypt against.
     */
    key: pulumi.Input<string>;
    /**
     * The version of the key to use for encryption. If not set, uses the latest version. Must be greater than or equal to the key's `minEncryptionVersion`, if set.
     */
    keyVersion?: pulumi.Input<number>;
    /**
     * Plaintext to be encoded.
     */
    plaintext: pulumi.Input<string>;
}
