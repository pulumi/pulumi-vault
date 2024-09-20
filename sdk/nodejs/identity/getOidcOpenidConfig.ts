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
 * const key = new vault.identity.OidcKey("key", {
 *     name: "key",
 *     allowedClientIds: ["*"],
 *     rotationPeriod: 3600,
 *     verificationTtl: 3600,
 * });
 * const app = new vault.identity.OidcClient("app", {
 *     name: "application",
 *     key: key.name,
 *     redirectUris: [
 *         "http://127.0.0.1:9200/v1/auth-methods/oidc:authenticate:callback",
 *         "http://127.0.0.1:8251/callback",
 *         "http://127.0.0.1:8080/callback",
 *     ],
 *     idTokenTtl: 2400,
 *     accessTokenTtl: 7200,
 * });
 * const provider = new vault.identity.OidcProvider("provider", {
 *     name: "provider",
 *     allowedClientIds: [test.clientId],
 * });
 * const config = vault.identity.getOidcOpenidConfigOutput({
 *     name: provider.name,
 * });
 * ```
 */
export function getOidcOpenidConfig(args: GetOidcOpenidConfigArgs, opts?: pulumi.InvokeOptions): Promise<GetOidcOpenidConfigResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("vault:identity/getOidcOpenidConfig:getOidcOpenidConfig", {
        "name": args.name,
        "namespace": args.namespace,
    }, opts);
}

/**
 * A collection of arguments for invoking getOidcOpenidConfig.
 */
export interface GetOidcOpenidConfigArgs {
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
 * A collection of values returned by getOidcOpenidConfig.
 */
export interface GetOidcOpenidConfigResult {
    /**
     * The Authorization Endpoint for the provider.
     */
    readonly authorizationEndpoint: string;
    /**
     * The grant types supported by the provider.
     */
    readonly grantTypesSupporteds: string[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * The signing algorithms supported by 
     * the provider.
     */
    readonly idTokenSigningAlgValuesSupporteds: string[];
    /**
     * The URL of the issuer for the provider.
     */
    readonly issuer: string;
    /**
     * The well known keys URI for the provider.
     */
    readonly jwksUri: string;
    readonly name: string;
    readonly namespace?: string;
    /**
     * Specifies whether Request URI Parameter is 
     * supported by the provider.
     */
    readonly requestUriParameterSupported: boolean;
    /**
     * The response types supported by the provider.
     */
    readonly responseTypesSupporteds: string[];
    /**
     * The scopes supported by the provider.
     */
    readonly scopesSupporteds: string[];
    /**
     * The subject types supported by the provider.
     */
    readonly subjectTypesSupporteds: string[];
    /**
     * The Token Endpoint for the provider.
     */
    readonly tokenEndpoint: string;
    /**
     * The token endpoint auth methods supported by the provider.
     */
    readonly tokenEndpointAuthMethodsSupporteds: string[];
    /**
     * The User Info Endpoint for the provider
     */
    readonly userinfoEndpoint: string;
}
/**
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vault from "@pulumi/vault";
 *
 * const key = new vault.identity.OidcKey("key", {
 *     name: "key",
 *     allowedClientIds: ["*"],
 *     rotationPeriod: 3600,
 *     verificationTtl: 3600,
 * });
 * const app = new vault.identity.OidcClient("app", {
 *     name: "application",
 *     key: key.name,
 *     redirectUris: [
 *         "http://127.0.0.1:9200/v1/auth-methods/oidc:authenticate:callback",
 *         "http://127.0.0.1:8251/callback",
 *         "http://127.0.0.1:8080/callback",
 *     ],
 *     idTokenTtl: 2400,
 *     accessTokenTtl: 7200,
 * });
 * const provider = new vault.identity.OidcProvider("provider", {
 *     name: "provider",
 *     allowedClientIds: [test.clientId],
 * });
 * const config = vault.identity.getOidcOpenidConfigOutput({
 *     name: provider.name,
 * });
 * ```
 */
export function getOidcOpenidConfigOutput(args: GetOidcOpenidConfigOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetOidcOpenidConfigResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("vault:identity/getOidcOpenidConfig:getOidcOpenidConfig", {
        "name": args.name,
        "namespace": args.namespace,
    }, opts);
}

/**
 * A collection of arguments for invoking getOidcOpenidConfig.
 */
export interface GetOidcOpenidConfigOutputArgs {
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
