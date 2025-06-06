// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * This is a data source which can be used to sign an SSH public key
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vault from "@pulumi/vault";
 *
 * const test = vault.ssh.getSecretBackendSign({
 *     path: "ssh",
 *     publicKey: "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABgQDR6q4PTcuIkpdGEqaCaxnR8/REqlbSiEIKaRZkVSjiTXOaiSfUsy9cY2+7+oO9fLMUrhylImerjzEoagX1IjYvc9IeUBaRnfacN7QwUDfstgp2jknbg7rNX9j9nFxwltV/jYQPcRq8Ud0wn1nb4qixq+diM7+Up+xJOeaKxbpjEUJH5dcvaBB+Aa24tJpjOQxtFyQ6dUxlgJu0tcygZR92kKYCVjZDohlSED3i/Ak2KFwqCKx2IZWq9z1vMEgmRzv++4Qt1OsbpW8itiCyWn6lmV33eDCdjMrr9TEThQNnMinPrHdmVUnPZ/OomP+rLDRE9lQR16uaSvKhg5TWOFIXRPyEhX9arEATrE4KSWeQN2qgHOb6P24YqgEm1ZdHJq25q/nBBAa1x0tFMiWqZwOsGeJ9nTeOeyiqFKH5YRBo6DIy3ag3taFsfQSve6oqjnrudUd1hJ8/bNSz8amECfP0ULvAEAgpiurj3eCPc3OcXl4tAld9F6KwabEJV5eelcs= user@example.com",
 *     name: "test",
 *     validPrincipals: "my-user",
 * });
 * ```
 */
export function getSecretBackendSign(args: GetSecretBackendSignArgs, opts?: pulumi.InvokeOptions): Promise<GetSecretBackendSignResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("vault:ssh/getSecretBackendSign:getSecretBackendSign", {
        "certType": args.certType,
        "criticalOptions": args.criticalOptions,
        "extensions": args.extensions,
        "keyId": args.keyId,
        "name": args.name,
        "namespace": args.namespace,
        "path": args.path,
        "publicKey": args.publicKey,
        "ttl": args.ttl,
        "validPrincipals": args.validPrincipals,
    }, opts);
}

/**
 * A collection of arguments for invoking getSecretBackendSign.
 */
export interface GetSecretBackendSignArgs {
    /**
     * Specifies the type of certificate to be created; either "user" or "host".
     */
    certType?: string;
    /**
     * Specifies a map of the critical options that the certificate should be signed for. Defaults to none.
     */
    criticalOptions?: {[key: string]: string};
    /**
     * Specifies a map of the extensions that the certificate should be signed for. Defaults to none.
     */
    extensions?: {[key: string]: string};
    /**
     * Specifies the key id that the created certificate should have. If not specified, the display name of the token will be used.
     */
    keyId?: string;
    /**
     * Specifies the name of the role to sign.
     */
    name: string;
    namespace?: string;
    /**
     * Full path where SSH backend is mounted.
     */
    path: string;
    /**
     * Specifies the SSH public key that should be signed.
     */
    publicKey: string;
    /**
     * Specifies the Requested Time To Live. Cannot be greater than the role's maxTtl value. If not provided, the role's ttl value will be used. Note that the role values default to system values if not explicitly set.
     */
    ttl?: string;
    /**
     * Specifies valid principals, either usernames or hostnames, that the certificate should be signed for. Required unless the role has specified allowEmptyPrincipals or a value has been set for either the defaultUser or defaultUserTemplate role parameters.
     */
    validPrincipals?: string;
}

/**
 * A collection of values returned by getSecretBackendSign.
 */
export interface GetSecretBackendSignResult {
    readonly certType?: string;
    readonly criticalOptions?: {[key: string]: string};
    readonly extensions?: {[key: string]: string};
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly keyId?: string;
    readonly name: string;
    readonly namespace?: string;
    readonly path: string;
    readonly publicKey: string;
    /**
     * The serial number of the certificate returned from Vault
     */
    readonly serialNumber: string;
    /**
     * The signed certificate returned from Vault
     */
    readonly signedKey: string;
    readonly ttl?: string;
    readonly validPrincipals?: string;
}
/**
 * This is a data source which can be used to sign an SSH public key
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vault from "@pulumi/vault";
 *
 * const test = vault.ssh.getSecretBackendSign({
 *     path: "ssh",
 *     publicKey: "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABgQDR6q4PTcuIkpdGEqaCaxnR8/REqlbSiEIKaRZkVSjiTXOaiSfUsy9cY2+7+oO9fLMUrhylImerjzEoagX1IjYvc9IeUBaRnfacN7QwUDfstgp2jknbg7rNX9j9nFxwltV/jYQPcRq8Ud0wn1nb4qixq+diM7+Up+xJOeaKxbpjEUJH5dcvaBB+Aa24tJpjOQxtFyQ6dUxlgJu0tcygZR92kKYCVjZDohlSED3i/Ak2KFwqCKx2IZWq9z1vMEgmRzv++4Qt1OsbpW8itiCyWn6lmV33eDCdjMrr9TEThQNnMinPrHdmVUnPZ/OomP+rLDRE9lQR16uaSvKhg5TWOFIXRPyEhX9arEATrE4KSWeQN2qgHOb6P24YqgEm1ZdHJq25q/nBBAa1x0tFMiWqZwOsGeJ9nTeOeyiqFKH5YRBo6DIy3ag3taFsfQSve6oqjnrudUd1hJ8/bNSz8amECfP0ULvAEAgpiurj3eCPc3OcXl4tAld9F6KwabEJV5eelcs= user@example.com",
 *     name: "test",
 *     validPrincipals: "my-user",
 * });
 * ```
 */
export function getSecretBackendSignOutput(args: GetSecretBackendSignOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetSecretBackendSignResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("vault:ssh/getSecretBackendSign:getSecretBackendSign", {
        "certType": args.certType,
        "criticalOptions": args.criticalOptions,
        "extensions": args.extensions,
        "keyId": args.keyId,
        "name": args.name,
        "namespace": args.namespace,
        "path": args.path,
        "publicKey": args.publicKey,
        "ttl": args.ttl,
        "validPrincipals": args.validPrincipals,
    }, opts);
}

/**
 * A collection of arguments for invoking getSecretBackendSign.
 */
export interface GetSecretBackendSignOutputArgs {
    /**
     * Specifies the type of certificate to be created; either "user" or "host".
     */
    certType?: pulumi.Input<string>;
    /**
     * Specifies a map of the critical options that the certificate should be signed for. Defaults to none.
     */
    criticalOptions?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Specifies a map of the extensions that the certificate should be signed for. Defaults to none.
     */
    extensions?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Specifies the key id that the created certificate should have. If not specified, the display name of the token will be used.
     */
    keyId?: pulumi.Input<string>;
    /**
     * Specifies the name of the role to sign.
     */
    name: pulumi.Input<string>;
    namespace?: pulumi.Input<string>;
    /**
     * Full path where SSH backend is mounted.
     */
    path: pulumi.Input<string>;
    /**
     * Specifies the SSH public key that should be signed.
     */
    publicKey: pulumi.Input<string>;
    /**
     * Specifies the Requested Time To Live. Cannot be greater than the role's maxTtl value. If not provided, the role's ttl value will be used. Note that the role values default to system values if not explicitly set.
     */
    ttl?: pulumi.Input<string>;
    /**
     * Specifies valid principals, either usernames or hostnames, that the certificate should be signed for. Required unless the role has specified allowEmptyPrincipals or a value has been set for either the defaultUser or defaultUserTemplate role parameters.
     */
    validPrincipals?: pulumi.Input<string>;
}
