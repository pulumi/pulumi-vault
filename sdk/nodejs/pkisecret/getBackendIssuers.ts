// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
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
 * const pki = new vault.Mount("pki", {
 *     path: "pki",
 *     type: "pki",
 *     description: "PKI secret engine mount",
 * });
 * const root = new vault.pkisecret.SecretBackendRootCert("root", {
 *     backend: pki.path,
 *     type: "internal",
 *     commonName: "example",
 *     ttl: "86400",
 *     issuerName: "example",
 * });
 * const test = vault.pkiSecret.getBackendIssuersOutput({
 *     backend: root.backend,
 * });
 * ```
 */
export function getBackendIssuers(args: GetBackendIssuersArgs, opts?: pulumi.InvokeOptions): Promise<GetBackendIssuersResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("vault:pkiSecret/getBackendIssuers:getBackendIssuers", {
        "backend": args.backend,
        "namespace": args.namespace,
    }, opts);
}

/**
 * A collection of arguments for invoking getBackendIssuers.
 */
export interface GetBackendIssuersArgs {
    /**
     * The path to the PKI secret backend to
     * read the issuers from, with no leading or trailing `/`s.
     */
    backend: string;
    /**
     * The namespace of the target resource.
     * The value should not contain leading or trailing forward slashes.
     * The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
     * *Available only for Vault Enterprise*.
     */
    namespace?: string;
}

/**
 * A collection of values returned by getBackendIssuers.
 */
export interface GetBackendIssuersResult {
    readonly backend: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * Map of issuer strings read from Vault.
     */
    readonly keyInfo: {[key: string]: any};
    /**
     * JSON-encoded issuer data read from Vault.
     */
    readonly keyInfoJson: string;
    /**
     * Keys used by issuers under the backend path.
     */
    readonly keys: string[];
    readonly namespace?: string;
}
/**
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vault from "@pulumi/vault";
 *
 * const pki = new vault.Mount("pki", {
 *     path: "pki",
 *     type: "pki",
 *     description: "PKI secret engine mount",
 * });
 * const root = new vault.pkisecret.SecretBackendRootCert("root", {
 *     backend: pki.path,
 *     type: "internal",
 *     commonName: "example",
 *     ttl: "86400",
 *     issuerName: "example",
 * });
 * const test = vault.pkiSecret.getBackendIssuersOutput({
 *     backend: root.backend,
 * });
 * ```
 */
export function getBackendIssuersOutput(args: GetBackendIssuersOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetBackendIssuersResult> {
    return pulumi.output(args).apply((a: any) => getBackendIssuers(a, opts))
}

/**
 * A collection of arguments for invoking getBackendIssuers.
 */
export interface GetBackendIssuersOutputArgs {
    /**
     * The path to the PKI secret backend to
     * read the issuers from, with no leading or trailing `/`s.
     */
    backend: pulumi.Input<string>;
    /**
     * The namespace of the target resource.
     * The value should not contain leading or trailing forward slashes.
     * The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
     * *Available only for Vault Enterprise*.
     */
    namespace?: pulumi.Input<string>;
}
