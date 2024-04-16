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
 * const app = new vault.identity.OidcClient("app", {
 *     name: "application",
 *     redirectUris: [
 *         "http://127.0.0.1:9200/v1/auth-methods/oidc:authenticate:callback",
 *         "http://127.0.0.1:8251/callback",
 *         "http://127.0.0.1:8080/callback",
 *     ],
 *     idTokenTtl: 2400,
 *     accessTokenTtl: 7200,
 * });
 * const creds = vault.identity.getOidcClientCredsOutput({
 *     name: app.name,
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getOidcClientCreds(args: GetOidcClientCredsArgs, opts?: pulumi.InvokeOptions): Promise<GetOidcClientCredsResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("vault:identity/getOidcClientCreds:getOidcClientCreds", {
        "name": args.name,
        "namespace": args.namespace,
    }, opts);
}

/**
 * A collection of arguments for invoking getOidcClientCreds.
 */
export interface GetOidcClientCredsArgs {
    /**
     * The name of the OIDC Client in Vault.
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
 * A collection of values returned by getOidcClientCreds.
 */
export interface GetOidcClientCredsResult {
    /**
     * The Client ID returned by Vault.
     */
    readonly clientId: string;
    /**
     * The Client Secret Key returned by Vault.
     * For public OpenID Clients `clientSecret` is set to an empty string `""`
     */
    readonly clientSecret: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
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
 * const app = new vault.identity.OidcClient("app", {
 *     name: "application",
 *     redirectUris: [
 *         "http://127.0.0.1:9200/v1/auth-methods/oidc:authenticate:callback",
 *         "http://127.0.0.1:8251/callback",
 *         "http://127.0.0.1:8080/callback",
 *     ],
 *     idTokenTtl: 2400,
 *     accessTokenTtl: 7200,
 * });
 * const creds = vault.identity.getOidcClientCredsOutput({
 *     name: app.name,
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getOidcClientCredsOutput(args: GetOidcClientCredsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetOidcClientCredsResult> {
    return pulumi.output(args).apply((a: any) => getOidcClientCreds(a, opts))
}

/**
 * A collection of arguments for invoking getOidcClientCreds.
 */
export interface GetOidcClientCredsOutputArgs {
    /**
     * The name of the OIDC Client in Vault.
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
