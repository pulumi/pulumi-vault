// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vault from "@pulumi/vault";
 *
 * const config = new vault.NomadSecretBackend("config", {
 *     backend: "nomad",
 *     description: "test description",
 *     defaultLeaseTtlSeconds: 3600,
 *     maxLeaseTtlSeconds: 7200,
 *     address: "https://127.0.0.1:4646",
 *     token: "ae20ceaa-...",
 * });
 * const test = new vault.NomadSecretRole("test", {
 *     backend: config.backend,
 *     role: "test",
 *     type: "client",
 *     policies: ["readonly"],
 * });
 * const token = pulumi.all([config.backend, test.role]).apply(([backend, role]) => vault.getNomadAccessTokenOutput({
 *     backend: backend,
 *     role: role,
 * }));
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getNomadAccessToken(args: GetNomadAccessTokenArgs, opts?: pulumi.InvokeOptions): Promise<GetNomadAccessTokenResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("vault:index/getNomadAccessToken:getNomadAccessToken", {
        "backend": args.backend,
        "namespace": args.namespace,
        "role": args.role,
    }, opts);
}

/**
 * A collection of arguments for invoking getNomadAccessToken.
 */
export interface GetNomadAccessTokenArgs {
    /**
     * The path to the Nomad secret backend to
     * read credentials from, with no leading or trailing `/`s.
     */
    backend: string;
    /**
     * The namespace of the target resource.
     * The value should not contain leading or trailing forward slashes.
     * The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
     * *Available only for Vault Enterprise*.
     */
    namespace?: string;
    /**
     * The name of the Nomad secret backend role to generate
     * a token for, with no leading or trailing `/`s.
     */
    role: string;
}

/**
 * A collection of values returned by getNomadAccessToken.
 */
export interface GetNomadAccessTokenResult {
    /**
     * The public identifier for a specific token. It can be used 
     * to look up information about a token or to revoke a token.
     */
    readonly accessorId: string;
    readonly backend: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly namespace?: string;
    readonly role: string;
    /**
     * The token to be used when making requests to Nomad and should be kept private.
     */
    readonly secretId: string;
}
/**
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vault from "@pulumi/vault";
 *
 * const config = new vault.NomadSecretBackend("config", {
 *     backend: "nomad",
 *     description: "test description",
 *     defaultLeaseTtlSeconds: 3600,
 *     maxLeaseTtlSeconds: 7200,
 *     address: "https://127.0.0.1:4646",
 *     token: "ae20ceaa-...",
 * });
 * const test = new vault.NomadSecretRole("test", {
 *     backend: config.backend,
 *     role: "test",
 *     type: "client",
 *     policies: ["readonly"],
 * });
 * const token = pulumi.all([config.backend, test.role]).apply(([backend, role]) => vault.getNomadAccessTokenOutput({
 *     backend: backend,
 *     role: role,
 * }));
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getNomadAccessTokenOutput(args: GetNomadAccessTokenOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetNomadAccessTokenResult> {
    return pulumi.output(args).apply((a: any) => getNomadAccessToken(a, opts))
}

/**
 * A collection of arguments for invoking getNomadAccessToken.
 */
export interface GetNomadAccessTokenOutputArgs {
    /**
     * The path to the Nomad secret backend to
     * read credentials from, with no leading or trailing `/`s.
     */
    backend: pulumi.Input<string>;
    /**
     * The namespace of the target resource.
     * The value should not contain leading or trailing forward slashes.
     * The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
     * *Available only for Vault Enterprise*.
     */
    namespace?: pulumi.Input<string>;
    /**
     * The name of the Nomad secret backend role to generate
     * a token for, with no leading or trailing `/`s.
     */
    role: pulumi.Input<string>;
}
