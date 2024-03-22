// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vault from "@pulumi/vault";
 *
 * const key = new vault.identity.OidcKey("key", {
 *     allowedClientIds: ["*"],
 *     rotationPeriod: 3600,
 *     verificationTtl: 3600,
 * });
 * const app = new vault.identity.OidcClient("app", {
 *     key: key.name,
 *     redirectUris: [
 *         "http://127.0.0.1:9200/v1/auth-methods/oidc:authenticate:callback",
 *         "http://127.0.0.1:8251/callback",
 *         "http://127.0.0.1:8080/callback",
 *     ],
 *     idTokenTtl: 2400,
 *     accessTokenTtl: 7200,
 * });
 * const provider = new vault.identity.OidcProvider("provider", {allowedClientIds: [vault_identity_oidc_client.test.client_id]});
 * const publicKeys = vault.identity.getOidcPublicKeysOutput({
 *     name: provider.name,
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getOidcPublicKeys(args: GetOidcPublicKeysArgs, opts?: pulumi.InvokeOptions): Promise<GetOidcPublicKeysResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("vault:identity/getOidcPublicKeys:getOidcPublicKeys", {
        "name": args.name,
        "namespace": args.namespace,
    }, opts);
}

/**
 * A collection of arguments for invoking getOidcPublicKeys.
 */
export interface GetOidcPublicKeysArgs {
    /**
     * The name of the OIDC Provider in Vault.
     */
    name: string;
    /**
     * The namespace of the target resource.
     * The value should not contain leading or trailing forward slashes.
     * The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
     * *Available only for Vault Enterprise*.
     */
    namespace?: string;
}

/**
 * A collection of values returned by getOidcPublicKeys.
 */
export interface GetOidcPublicKeysResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * The public portion of keys for an OIDC provider. 
     * Clients can use them to validate the authenticity of an identity token.
     */
    readonly keys: {[key: string]: any}[];
    readonly name: string;
    readonly namespace?: string;
}
/**
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vault from "@pulumi/vault";
 *
 * const key = new vault.identity.OidcKey("key", {
 *     allowedClientIds: ["*"],
 *     rotationPeriod: 3600,
 *     verificationTtl: 3600,
 * });
 * const app = new vault.identity.OidcClient("app", {
 *     key: key.name,
 *     redirectUris: [
 *         "http://127.0.0.1:9200/v1/auth-methods/oidc:authenticate:callback",
 *         "http://127.0.0.1:8251/callback",
 *         "http://127.0.0.1:8080/callback",
 *     ],
 *     idTokenTtl: 2400,
 *     accessTokenTtl: 7200,
 * });
 * const provider = new vault.identity.OidcProvider("provider", {allowedClientIds: [vault_identity_oidc_client.test.client_id]});
 * const publicKeys = vault.identity.getOidcPublicKeysOutput({
 *     name: provider.name,
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getOidcPublicKeysOutput(args: GetOidcPublicKeysOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetOidcPublicKeysResult> {
    return pulumi.output(args).apply((a: any) => getOidcPublicKeys(a, opts))
}

/**
 * A collection of arguments for invoking getOidcPublicKeys.
 */
export interface GetOidcPublicKeysOutputArgs {
    /**
     * The name of the OIDC Provider in Vault.
     */
    name: pulumi.Input<string>;
    /**
     * The namespace of the target resource.
     * The value should not contain leading or trailing forward slashes.
     * The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
     * *Available only for Vault Enterprise*.
     */
    namespace?: pulumi.Input<string>;
}
